// Package signal provides SIGNAL TSV parsing helpers.
// The actual store is in the genomicindex package (unified SQLite).
package signal

import "strconv"

// Entry represents a single SIGNAL variant annotation.
type Entry struct {
	Pos       int64
	Ref       string
	Alt       string
	Gene      string
	CountAll  int     // n_impact: total carrier count
	FreqAll   float64 // f_impact: overall frequency
	Biallelic float64 // f_biallelic: biallelic frequency
}

// IndexColumns maps column names to their indices in the header.
func IndexColumns(header []string, names ...string) map[string]int {
	idx := make(map[string]int, len(names))
	for _, name := range names {
		idx[name] = -1
	}
	for i, col := range header {
		for _, name := range names {
			if col == name {
				idx[name] = i
			}
		}
	}
	return idx
}

// FormatFreq formats a frequency value for output.
func FormatFreq(v float64) string {
	if v == 0 {
		return "0"
	}
	return strconv.FormatFloat(v, 'g', 6, 64)
}
