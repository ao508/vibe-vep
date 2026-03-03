package output_test

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/inodb/vibe-vep/internal/annotate"
	"github.com/inodb/vibe-vep/internal/cache"
	"github.com/inodb/vibe-vep/internal/datasource/alphamissense"
	"github.com/inodb/vibe-vep/internal/datasource/clinvar"
	"github.com/inodb/vibe-vep/internal/datasource/hotspots"
	"github.com/inodb/vibe-vep/internal/datasource/signal"
	"github.com/inodb/vibe-vep/internal/maf"
)

// sourcesCtx holds all loaded annotation sources for benchmarking.
type sourcesCtx struct {
	amStore *alphamissense.Store
	amCount int64
	hsStore *hotspots.Store
	cvStore *clinvar.Store
	sigStore *signal.Store
}

// TestAnnotationSourcesBenchmark runs annotation source coverage and performance benchmarks
// against all TCGA MAF files and generates testdata/tcga/annotation_sources_report.md.
//
// Skipped with -short. Run with:
//
//	CGO_ENABLED=1 go test ./internal/output/ -run TestAnnotationSourcesBenchmark -v -count=1 -timeout 30m
func TestAnnotationSourcesBenchmark(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping annotation sources benchmark in short mode")
	}

	// Find TCGA MAF files.
	tcgaDir := findTCGADir(t)
	mafFiles, err := filepath.Glob(filepath.Join(tcgaDir, "*_data_mutations.txt"))
	if err != nil {
		t.Fatalf("glob MAF files: %v", err)
	}
	if len(mafFiles) == 0 {
		t.Skip("no TCGA MAF files found in", tcgaDir)
	}
	sort.Strings(mafFiles)

	// Load GENCODE cache.
	gtfPath, fastaPath, canonicalPath := findGENCODEFiles(t)
	c := cache.New()
	loader := cache.NewGENCODELoader(gtfPath, fastaPath)
	if canonicalPath != "" {
		overrides, err := cache.LoadCanonicalOverrides(canonicalPath)
		if err != nil {
			t.Logf("warning: could not load canonical overrides: %v", err)
		} else {
			loader.SetCanonicalOverrides(overrides)
		}
	}
	if err := loader.Load(c); err != nil {
		t.Fatalf("load GENCODE cache: %v", err)
	}
	t.Logf("loaded %d transcripts", c.TranscriptCount())

	var ctx sourcesCtx

	// Open AlphaMissense database.
	home, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("get home dir: %v", err)
	}
	dbPath := filepath.Join(home, ".vibe-vep", "grch38", "annotations.duckdb")
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		t.Skipf("AlphaMissense database not found at %s — run: vibe-vep prepare", dbPath)
	}
	ctx.amStore, err = alphamissense.Open(dbPath)
	if err != nil {
		t.Fatalf("open AlphaMissense: %v", err)
	}
	defer ctx.amStore.Close()
	if !ctx.amStore.Loaded() {
		t.Skip("AlphaMissense database is empty")
	}
	ctx.amCount, err = ctx.amStore.Count()
	if err != nil {
		t.Fatalf("count AlphaMissense rows: %v", err)
	}
	t.Logf("AlphaMissense variants in database: %d", ctx.amCount)

	// Load Hotspots.
	hotspotsPath := "/home/ino/vibe-vep/genome-nexus-importer/data/common_input/hotspots_v2_and_3d.txt"
	if _, err := os.Stat(hotspotsPath); err == nil {
		ctx.hsStore, err = hotspots.Load(hotspotsPath)
		if err != nil {
			t.Fatalf("load hotspots: %v", err)
		}
		t.Logf("loaded hotspots: %d genes, %d positions", ctx.hsStore.GeneCount(), ctx.hsStore.HotspotCount())
	} else {
		t.Logf("hotspots file not found, skipping")
	}

	// Load ClinVar.
	cvPath := filepath.Join(home, ".vibe-vep", "grch38", "clinvar.vcf.gz")
	if _, err := os.Stat(cvPath); err == nil {
		start := time.Now()
		ctx.cvStore, err = clinvar.Load(cvPath)
		if err != nil {
			t.Fatalf("load ClinVar: %v", err)
		}
		t.Logf("loaded ClinVar: %d variants in %s", ctx.cvStore.Count(), time.Since(start).Round(time.Millisecond))
	} else {
		t.Logf("ClinVar VCF not found, skipping")
	}

	// Load SIGNAL.
	sigPath := "/home/ino/vibe-vep/genome-nexus-importer/data/signal/input/signaldb_all_variants_frequencies.txt"
	if _, err := os.Stat(sigPath); err == nil {
		ctx.sigStore, err = signal.Load(sigPath)
		if err != nil {
			t.Fatalf("load SIGNAL: %v", err)
		}
		t.Logf("loaded SIGNAL: %d variants (note: GRCh37 coordinates)", ctx.sigStore.Count())
	} else {
		t.Logf("SIGNAL file not found, skipping")
	}

	var results []sourceStudyResult

	for _, mafFile := range mafFiles {
		name := studyName(mafFile)
		t.Run(name, func(t *testing.T) {
			r := benchmarkStudy(t, mafFile, c, &ctx)
			r.name = name
			results = append(results, r)
		})
	}

	// Write report.
	reportPath := filepath.Join(tcgaDir, "annotation_sources_report.md")
	writeAnnotationSourcesReport(t, reportPath, results, c.TranscriptCount(), &ctx)
}

// benchmarkStudy annotates a MAF file and benchmarks all annotation source lookups.
func benchmarkStudy(t *testing.T, mafFile string, c *cache.Cache, ctx *sourcesCtx) sourceStudyResult {
	t.Helper()

	parser, err := maf.NewParser(mafFile)
	if err != nil {
		t.Fatalf("open MAF: %v", err)
	}
	defer parser.Close()

	ann := annotate.NewAnnotator(c)

	var (
		totalVariants int
		missenseKeys  []alphamissense.LookupKey
		baseTime      time.Duration
		hsHits        int
		hsChecked     int
		hsTypeCounts  = make(map[string]int)
		cvHits        int
		cvSigCounts   = make(map[string]int)
		sigHits       int
	)
	seen := make(map[alphamissense.LookupKey]bool)

	for {
		v, _, err := parser.NextWithAnnotation()
		if err != nil {
			t.Fatalf("read variant: %v", err)
		}
		if v == nil {
			break
		}
		totalVariants++

		annStart := time.Now()
		vepAnns, err := ann.Annotate(v)
		baseTime += time.Since(annStart)
		if err != nil {
			continue
		}

		// Collect missense keys for AlphaMissense batch lookup.
		for _, a := range vepAnns {
			if !isMissenseConsequence(a.Consequence) {
				continue
			}
			chrom := v.NormalizeChrom()
			if len(chrom) > 0 && chrom[0] != 'c' {
				chrom = "chr" + chrom
			}
			k := alphamissense.LookupKey{Chrom: chrom, Pos: v.Pos, Ref: v.Ref, Alt: v.Alt}
			if !seen[k] {
				seen[k] = true
				missenseKeys = append(missenseKeys, k)
			}
			break
		}

		// Hotspot lookup.
		if ctx.hsStore != nil {
			for _, a := range vepAnns {
				if a.ProteinPosition > 0 && a.GeneName != "" {
					hsChecked++
					if h, ok := ctx.hsStore.Lookup(a.GeneName, a.ProteinPosition); ok {
						hsHits++
						hsTypeCounts[h.Type]++
					}
					break
				}
			}
		}

		// ClinVar lookup (variant-level).
		if ctx.cvStore != nil {
			chrom := v.NormalizeChrom()
			if e, ok := ctx.cvStore.Lookup(chrom, v.Pos, v.Ref, v.Alt); ok {
				cvHits++
				cvSigCounts[e.ClnSig]++
			}
		}

		// SIGNAL lookup (variant-level, note: GRCh37 coords so few hits on GRCh38 data).
		if ctx.sigStore != nil {
			chrom := v.NormalizeChrom()
			if _, ok := ctx.sigStore.Lookup(chrom, v.Pos, v.Ref, v.Alt); ok {
				sigHits++
			}
		}
	}

	missenseCount := len(missenseKeys)

	// AlphaMissense batch lookup.
	lookupStart := time.Now()
	amResults, err := ctx.amStore.BatchLookup(missenseKeys)
	amLookupTime := time.Since(lookupStart)
	if err != nil {
		t.Fatalf("batch lookup: %v", err)
	}

	amHits := len(amResults)
	classCounts := make(map[string]int)
	for _, r := range amResults {
		classCounts[r.Class]++
	}

	amCoverage := 0.0
	if missenseCount > 0 {
		amCoverage = float64(amHits) / float64(missenseCount) * 100
	}
	t.Logf("%d variants, %d missense, %d AM hits (%.1f%%), batch lookup %s",
		totalVariants, missenseCount, amHits, amCoverage, amLookupTime.Round(time.Millisecond))
	if ctx.hsStore != nil {
		t.Logf("  hotspots: %d checked, %d hits", hsChecked, hsHits)
	}
	if ctx.cvStore != nil {
		t.Logf("  ClinVar: %d hits out of %d variants", cvHits, totalVariants)
	}
	if ctx.sigStore != nil {
		t.Logf("  SIGNAL: %d hits out of %d variants (GRCh37 coords)", sigHits, totalVariants)
	}

	return sourceStudyResult{
		variants:     totalVariants,
		missense:     missenseCount,
		amHits:       amHits,
		classCounts:  classCounts,
		baseTime:     baseTime,
		amLookupTime: amLookupTime,
		hsChecked:    hsChecked,
		hsHits:       hsHits,
		hsTypeCounts: hsTypeCounts,
		cvHits:       cvHits,
		cvSigCounts:  cvSigCounts,
		sigHits:      sigHits,
	}
}

type sourceStudyResult struct {
	name         string
	variants     int
	missense     int
	amHits       int
	classCounts  map[string]int
	baseTime     time.Duration
	amLookupTime time.Duration
	hsChecked    int
	hsHits       int
	hsTypeCounts map[string]int
	cvHits       int
	cvSigCounts  map[string]int
	sigHits      int
}

// isMissenseConsequence returns true if the consequence includes missense_variant.
func isMissenseConsequence(consequence string) bool {
	for rest := consequence; rest != ""; {
		term := rest
		if i := strings.IndexByte(rest, ','); i >= 0 {
			term = rest[:i]
			rest = rest[i+1:]
		} else {
			rest = ""
		}
		if term == "missense_variant" {
			return true
		}
	}
	return false
}

func writeAnnotationSourcesReport(t *testing.T, path string, results []sourceStudyResult, transcriptCount int, ctx *sourcesCtx) {
	t.Helper()

	var sb strings.Builder
	sb.WriteString("# Annotation Sources Report\n\n")
	sb.WriteString(fmt.Sprintf("Generated: %s  \n", time.Now().UTC().Format("2006-01-02 15:04 UTC")))
	sb.WriteString(fmt.Sprintf("GENCODE transcripts: %d  \n", transcriptCount))
	sb.WriteString(fmt.Sprintf("AlphaMissense variants in database: %d  \n", ctx.amCount))
	if ctx.hsStore != nil {
		sb.WriteString(fmt.Sprintf("Cancer Hotspots: %d genes, %d positions  \n", ctx.hsStore.GeneCount(), ctx.hsStore.HotspotCount()))
	}
	if ctx.cvStore != nil {
		sb.WriteString(fmt.Sprintf("ClinVar variants: %d  \n", ctx.cvStore.Count()))
	}
	if ctx.sigStore != nil {
		sb.WriteString(fmt.Sprintf("SIGNAL variants: %d (GRCh37 coordinates)  \n", ctx.sigStore.Count()))
	}
	sb.WriteString(fmt.Sprintf("Workers: %d (GOMAXPROCS)\n\n", runtime.NumCPU()))

	// --- AlphaMissense Coverage ---
	sb.WriteString("## AlphaMissense Coverage\n\n")
	sb.WriteString("| Study | Variants | Missense | AM Hits | Coverage | likely_benign | ambiguous | likely_pathogenic |\n")
	sb.WriteString("|-------|----------|----------|---------|----------|---------------|-----------|-------------------|\n")

	var totVariants, totMissense, totAMHits int
	totClass := make(map[string]int)
	for _, r := range results {
		coverage := 0.0
		if r.missense > 0 {
			coverage = float64(r.amHits) / float64(r.missense) * 100
		}
		sb.WriteString(fmt.Sprintf("| %s | %d | %d | %d | %.1f%% | %d | %d | %d |\n",
			r.name, r.variants, r.missense, r.amHits, coverage,
			r.classCounts["likely_benign"],
			r.classCounts["ambiguous"],
			r.classCounts["likely_pathogenic"]))
		totVariants += r.variants
		totMissense += r.missense
		totAMHits += r.amHits
		for cls, n := range r.classCounts {
			totClass[cls] += n
		}
	}
	totCoverage := 0.0
	if totMissense > 0 {
		totCoverage = float64(totAMHits) / float64(totMissense) * 100
	}
	sb.WriteString(fmt.Sprintf("| **Total** | **%d** | **%d** | **%d** | **%.1f%%** | **%d** | **%d** | **%d** |\n\n",
		totVariants, totMissense, totAMHits, totCoverage,
		totClass["likely_benign"],
		totClass["ambiguous"],
		totClass["likely_pathogenic"]))

	// --- Cancer Hotspots Coverage ---
	if ctx.hsStore != nil {
		sb.WriteString("## Cancer Hotspots Coverage\n\n")
		sb.WriteString("| Study | Variants | Checked | Hotspot Hits | Hit Rate | single residue | in-frame indel | 3d | splice site |\n")
		sb.WriteString("|-------|----------|---------|--------------|----------|----------------|----------------|----|-------------|\n")

		var totChecked, totHsHits int
		totHsType := make(map[string]int)
		for _, r := range results {
			hitRate := 0.0
			if r.hsChecked > 0 {
				hitRate = float64(r.hsHits) / float64(r.hsChecked) * 100
			}
			sb.WriteString(fmt.Sprintf("| %s | %d | %d | %d | %.2f%% | %d | %d | %d | %d |\n",
				r.name, r.variants, r.hsChecked, r.hsHits, hitRate,
				r.hsTypeCounts["single residue"],
				r.hsTypeCounts["in-frame indel"],
				r.hsTypeCounts["3d"],
				r.hsTypeCounts["splice site"]))
			totChecked += r.hsChecked
			totHsHits += r.hsHits
			for typ, n := range r.hsTypeCounts {
				totHsType[typ] += n
			}
		}
		totHsRate := 0.0
		if totChecked > 0 {
			totHsRate = float64(totHsHits) / float64(totChecked) * 100
		}
		sb.WriteString(fmt.Sprintf("| **Total** | **%d** | **%d** | **%d** | **%.2f%%** | **%d** | **%d** | **%d** | **%d** |\n\n",
			totVariants, totChecked, totHsHits, totHsRate,
			totHsType["single residue"],
			totHsType["in-frame indel"],
			totHsType["3d"],
			totHsType["splice site"]))
	}

	// --- ClinVar Coverage ---
	if ctx.cvStore != nil {
		sb.WriteString("## ClinVar Coverage\n\n")
		sb.WriteString("| Study | Variants | ClinVar Hits | Hit Rate | Pathogenic | Likely_pathogenic | Uncertain | Benign | Likely_benign | Other |\n")
		sb.WriteString("|-------|----------|--------------|----------|------------|-------------------|-----------|--------|---------------|-------|\n")

		var totCVHits int
		totCVSig := make(map[string]int)
		for _, r := range results {
			hitRate := 0.0
			if r.variants > 0 {
				hitRate = float64(r.cvHits) / float64(r.variants) * 100
			}
			other := r.cvHits
			for _, key := range []string{"Pathogenic", "Likely_pathogenic", "Uncertain_significance", "Benign", "Likely_benign"} {
				other -= r.cvSigCounts[key]
			}
			sb.WriteString(fmt.Sprintf("| %s | %d | %d | %.2f%% | %d | %d | %d | %d | %d | %d |\n",
				r.name, r.variants, r.cvHits, hitRate,
				r.cvSigCounts["Pathogenic"],
				r.cvSigCounts["Likely_pathogenic"],
				r.cvSigCounts["Uncertain_significance"],
				r.cvSigCounts["Benign"],
				r.cvSigCounts["Likely_benign"],
				other))
			totCVHits += r.cvHits
			for sig, n := range r.cvSigCounts {
				totCVSig[sig] += n
			}
		}
		totCVRate := 0.0
		if totVariants > 0 {
			totCVRate = float64(totCVHits) / float64(totVariants) * 100
		}
		totOther := totCVHits
		for _, key := range []string{"Pathogenic", "Likely_pathogenic", "Uncertain_significance", "Benign", "Likely_benign"} {
			totOther -= totCVSig[key]
		}
		sb.WriteString(fmt.Sprintf("| **Total** | **%d** | **%d** | **%.2f%%** | **%d** | **%d** | **%d** | **%d** | **%d** | **%d** |\n\n",
			totVariants, totCVHits, totCVRate,
			totCVSig["Pathogenic"],
			totCVSig["Likely_pathogenic"],
			totCVSig["Uncertain_significance"],
			totCVSig["Benign"],
			totCVSig["Likely_benign"],
			totOther))
	}

	// --- SIGNAL Coverage ---
	if ctx.sigStore != nil {
		sb.WriteString("## SIGNAL Coverage\n\n")
		sb.WriteString("*Note: SIGNAL data uses GRCh37 coordinates. TCGA GDC MAFs use GRCh38, so few hits are expected.*\n\n")
		sb.WriteString("| Study | Variants | SIGNAL Hits | Hit Rate |\n")
		sb.WriteString("|-------|----------|-------------|----------|\n")

		var totSigHits int
		for _, r := range results {
			hitRate := 0.0
			if r.variants > 0 {
				hitRate = float64(r.sigHits) / float64(r.variants) * 100
			}
			sb.WriteString(fmt.Sprintf("| %s | %d | %d | %.2f%% |\n",
				r.name, r.variants, r.sigHits, hitRate))
			totSigHits += r.sigHits
		}
		totSigRate := 0.0
		if totVariants > 0 {
			totSigRate = float64(totSigHits) / float64(totVariants) * 100
		}
		sb.WriteString(fmt.Sprintf("| **Total** | **%d** | **%d** | **%.2f%%** |\n\n",
			totVariants, totSigHits, totSigRate))
	}

	// --- AlphaMissense Performance ---
	sb.WriteString("## AlphaMissense Performance\n\n")
	sb.WriteString("| Study | Variants | Base Time | AM Lookup Time | AM Overhead | Lookups/sec |\n")
	sb.WriteString("|-------|----------|-----------|----------------|-------------|-------------|\n")

	var totBaseTime, totAMLookupTime time.Duration
	for _, r := range results {
		overhead := 0.0
		if r.baseTime > 0 {
			overhead = float64(r.amLookupTime) / float64(r.baseTime) * 100
		}
		lookupsPerSec := 0.0
		if r.amLookupTime > 0 {
			lookupsPerSec = float64(r.missense) / r.amLookupTime.Seconds()
		}
		sb.WriteString(fmt.Sprintf("| %s | %d | %s | %s | %.1f%% | %.0f |\n",
			r.name, r.variants,
			r.baseTime.Round(time.Millisecond),
			r.amLookupTime.Round(time.Millisecond),
			overhead, lookupsPerSec))
		totBaseTime += r.baseTime
		totAMLookupTime += r.amLookupTime
	}
	totOverhead := 0.0
	if totBaseTime > 0 {
		totOverhead = float64(totAMLookupTime) / float64(totBaseTime) * 100
	}
	totLookupsPerSec := 0.0
	if totAMLookupTime > 0 {
		totLookupsPerSec = float64(totMissense) / totAMLookupTime.Seconds()
	}
	sb.WriteString(fmt.Sprintf("| **Total** | **%d** | **%s** | **%s** | **%.1f%%** | **%.0f** |\n",
		totVariants,
		totBaseTime.Round(time.Millisecond),
		totAMLookupTime.Round(time.Millisecond),
		totOverhead, totLookupsPerSec))

	if err := os.WriteFile(path, []byte(sb.String()), 0644); err != nil {
		t.Fatalf("write report: %v", err)
	}
	t.Logf("report written to %s", path)
}
