// Package clinvar provides ClinVar VCF parsing helpers.
// The actual store is in the genomicindex package (unified SQLite).
package clinvar

import (
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

// ParseVCFLine parses a single VCF data line into a ClinVar Entry.
// Returns the entry, normalized chromosome, and whether parsing succeeded.
func ParseVCFLine(line string) (Entry, string, bool) {
	// VCF: CHROM POS ID REF ALT QUAL FILTER INFO ...
	fields := strings.SplitN(line, "\t", 9)
	if len(fields) < 8 {
		return Entry{}, "", false
	}

	chrom := NormalizeChrom(fields[0])
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
		ClnSig:  ExtractInfo(info, "CLNSIG="),
		RevStat: ExtractInfo(info, "CLNREVSTAT="),
		ClnDN:   Truncate(ExtractInfo(info, "CLNDN="), 200),
	}

	// Skip entries without clinical significance
	if entry.ClnSig == "" {
		return Entry{}, "", false
	}

	return entry, chrom, true
}

// ExtractInfo extracts a value from a VCF INFO field.
func ExtractInfo(info, key string) string {
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

// NormalizeChrom removes "chr" prefix for consistent lookups.
func NormalizeChrom(chrom string) string {
	return strings.TrimPrefix(chrom, "chr")
}

// Truncate limits a string to maxLen bytes.
func Truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen]
}
