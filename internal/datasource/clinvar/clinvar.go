// Package clinvar provides ClinVar clinical significance lookups.
// Data comes from the ClinVar VCF file (NCBI).
//
// Unlike smaller sources, ClinVar has 4M+ entries so gob caching would require
// too much memory (~800 MB). Instead, the gzipped VCF is parsed directly on
// each startup (~25s). This is acceptable since ClinVar is opt-in.
package clinvar

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Entry represents a single ClinVar variant annotation.
type Entry struct {
	Pos     int64
	Ref     string
	Alt     string
	ClnSig  string // Clinical significance (e.g., "Pathogenic")
	RevStat string // Review status
	ClnDN   string // Disease name(s), truncated to 200 chars
}

// Store holds ClinVar data as sorted slices per chromosome.
type Store struct {
	data map[string][]Entry // chromosome (no "chr" prefix) → sorted entries
}

// Load parses a ClinVar VCF file (gzipped or plain) into sorted in-memory slices.
func Load(path string) (*Store, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open clinvar file: %w", err)
	}
	defer f.Close()

	var scanner *bufio.Scanner
	if strings.HasSuffix(path, ".gz") {
		gz, err := gzip.NewReader(f)
		if err != nil {
			return nil, fmt.Errorf("open gzip reader: %w", err)
		}
		defer gz.Close()
		scanner = bufio.NewScanner(gz)
	} else {
		scanner = bufio.NewScanner(f)
	}
	scanner.Buffer(make([]byte, 4*1024*1024), 4*1024*1024)

	// String interning deduplicates identical strings (ClnSig has ~20 unique
	// values, RevStat ~10, Ref/Alt ~5 each) and detaches substrings from the
	// scanner's line buffer, preventing massive memory retention.
	intern := make(map[string]string)
	internStr := func(s string) string {
		if s == "" {
			return ""
		}
		if v, ok := intern[s]; ok {
			return v
		}
		// Make an independent copy to detach from scanner buffer.
		s = strings.Clone(s)
		intern[s] = s
		return s
	}

	data := make(map[string][]Entry)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || line[0] == '#' {
			continue
		}

		entry, chrom, ok := parseVCFLine(line)
		if !ok {
			continue
		}
		// Intern all string fields to reduce memory.
		entry.Ref = internStr(entry.Ref)
		entry.Alt = internStr(entry.Alt)
		entry.ClnSig = internStr(entry.ClnSig)
		entry.RevStat = internStr(entry.RevStat)
		entry.ClnDN = internStr(entry.ClnDN)
		data[chrom] = append(data[chrom], entry)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("read clinvar file: %w", err)
	}

	// Sort each chromosome's entries by position
	for chrom, entries := range data {
		sort.Slice(entries, func(i, j int) bool {
			if entries[i].Pos != entries[j].Pos {
				return entries[i].Pos < entries[j].Pos
			}
			if entries[i].Ref != entries[j].Ref {
				return entries[i].Ref < entries[j].Ref
			}
			return entries[i].Alt < entries[j].Alt
		})
		data[chrom] = entries
	}

	return &Store{data: data}, nil
}

// parseVCFLine parses a single VCF data line into a ClinVar Entry.
// Returns the entry, normalized chromosome, and whether parsing succeeded.
func parseVCFLine(line string) (Entry, string, bool) {
	// VCF: CHROM POS ID REF ALT QUAL FILTER INFO ...
	fields := strings.SplitN(line, "\t", 9)
	if len(fields) < 8 {
		return Entry{}, "", false
	}

	chrom := normalizeChrom(fields[0])
	pos, err := strconv.ParseInt(fields[1], 10, 64)
	if err != nil {
		return Entry{}, "", false
	}

	ref := fields[3]
	alt := fields[4]
	info := fields[7]

	// Handle multi-allelic: take first ALT only
	if idx := strings.IndexByte(alt, ','); idx >= 0 {
		alt = alt[:idx]
	}

	entry := Entry{
		Pos:     pos,
		Ref:     ref,
		Alt:     alt,
		ClnSig:  extractInfo(info, "CLNSIG="),
		RevStat: extractInfo(info, "CLNREVSTAT="),
		ClnDN:   truncate(extractInfo(info, "CLNDN="), 200),
	}

	// Skip entries without clinical significance
	if entry.ClnSig == "" {
		return Entry{}, "", false
	}

	return entry, chrom, true
}

// extractInfo extracts a value from a VCF INFO field.
func extractInfo(info, key string) string {
	idx := strings.Index(info, key)
	if idx < 0 {
		return ""
	}
	val := info[idx+len(key):]
	if end := strings.IndexByte(val, ';'); end >= 0 {
		val = val[:end]
	}
	return val
}

// Lookup finds ClinVar annotations for a specific variant.
func (s *Store) Lookup(chrom string, pos int64, ref, alt string) (Entry, bool) {
	chrom = normalizeChrom(chrom)
	entries := s.data[chrom]
	if len(entries) == 0 {
		return Entry{}, false
	}

	// Binary search for position
	i := sort.Search(len(entries), func(i int) bool {
		return entries[i].Pos >= pos
	})

	// Scan forward through entries at this position
	for ; i < len(entries) && entries[i].Pos == pos; i++ {
		if entries[i].Ref == ref && entries[i].Alt == alt {
			return entries[i], true
		}
	}
	return Entry{}, false
}

// Count returns the total number of ClinVar entries.
func (s *Store) Count() int {
	var n int
	for _, entries := range s.data {
		n += len(entries)
	}
	return n
}

// normalizeChrom removes "chr" prefix for consistent lookups.
func normalizeChrom(chrom string) string {
	return strings.TrimPrefix(chrom, "chr")
}

// truncate limits a string to maxLen bytes.
func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen]
}
