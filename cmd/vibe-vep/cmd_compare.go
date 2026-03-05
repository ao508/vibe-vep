package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"go.uber.org/zap"

	"github.com/inodb/vibe-vep/internal/annotate"
	"github.com/inodb/vibe-vep/internal/maf"
	"github.com/inodb/vibe-vep/internal/output"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newCompareCmd(verbose *bool) *cobra.Command {
	var (
		assembly string
		columns  string
		all      bool
	)

	cmd := &cobra.Command{
		Use:   "compare <maf-file>",
		Short: "Compare MAF annotations against VEP predictions",
		Long:  "Compare MAF annotations against VEP predictions with categorized mismatch analysis.",
		Example: `  vibe-vep compare data_mutations.txt
  vibe-vep compare --columns consequence,hgvsp data_mutations.txt
  vibe-vep compare --all data_mutations.txt`,
		Args: cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return viper.BindPFlags(cmd.Flags())
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			logger, err := newLogger(*verbose)
			if err != nil {
				return fmt.Errorf("creating logger: %w", err)
			}
			defer logger.Sync()

			// Parse columns
			colMap := make(map[string]bool)
			for _, c := range strings.Split(viper.GetString("columns"), ",") {
				c = strings.TrimSpace(strings.ToLower(c))
				if c != "" {
					colMap[c] = true
				}
			}

			return runCompare(logger, args[0],
				viper.GetString("assembly"),
				colMap,
				viper.GetBool("all"),
				viper.GetBool("no-cache"),
				viper.GetBool("clear-cache"),
			)
		},
	}

	cmd.Flags().StringVar(&assembly, "assembly", "GRCh38", "Genome assembly: GRCh37 or GRCh38")
	cmd.Flags().StringVar(&columns, "columns", "consequence,hgvsp,hgvsc", "Columns to compare (comma-separated)")
	cmd.Flags().BoolVar(&all, "all", false, "Show all rows, not just non-matches")
	addCacheFlags(cmd)

	return cmd
}

func runCompare(logger *zap.Logger, inputPath, assembly string, columns map[string]bool, showAll bool, noCache, clearCache bool) error {
	parser, err := maf.NewParser(inputPath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("%w (check that the file path is correct)", err)
		}
		return err
	}
	defer parser.Close()

	// Load transcript cache
	cr, err := loadCache(logger, assembly, noCache, clearCache)
	if err != nil {
		return err
	}
	if cr.store != nil {
		defer cr.store.Close()
	}
	defer cr.closeSources()

	ann := annotate.NewAnnotator(cr.cache)
	ann.SetLogger(logger)

	cmpWriter := output.NewCompareWriter(os.Stdout, columns, showAll)
	if err := cmpWriter.WriteHeader(); err != nil {
		return fmt.Errorf("writing header: %w", err)
	}

	// Parse variants in a goroutine, send to worker pool.
	items := make(chan annotate.WorkItem, 2*runtime.NumCPU())
	var parseErr error
	go func() {
		defer close(items)
		seq := 0
		for {
			v, mafAnn, err := parser.NextWithAnnotation()
			if err != nil {
				parseErr = fmt.Errorf("reading variant: %w", err)
				return
			}
			if v == nil {
				return
			}
			items <- annotate.WorkItem{Seq: seq, Variant: v, Extra: mafAnn}
			seq++
		}
	}()

	results := ann.ParallelAnnotate(items, 0)

	progress := func(n int) {
		logger.Info("progress", zap.Int("variants_processed", n))
	}

	if err := annotate.OrderedCollectWithProgress(results, 2*time.Second, progress, func(r annotate.WorkResult) error {
		mafAnn := r.Extra.(*maf.MAFAnnotation)
		if r.Err != nil {
			logger.Warn("failed to annotate variant",
				zap.String("chrom", r.Variant.Chrom),
				zap.Int64("pos", r.Variant.Pos),
				zap.Error(r.Err))
			return nil
		}
		return cmpWriter.WriteComparison(r.Variant, mafAnn, r.Anns)
	}); err != nil {
		return err
	}

	if parseErr != nil {
		return parseErr
	}

	if err := cmpWriter.Flush(); err != nil {
		return fmt.Errorf("flushing output: %w", err)
	}

	cmpWriter.WriteSummary(os.Stderr)

	return nil
}
