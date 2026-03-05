package clinvar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseVCFLine(t *testing.T) {
	line := "7\t140753336\t846933\tT\tA\t.\t.\tALLELEID=826271;CLNDN=Melanoma;CLNSIG=Pathogenic;CLNREVSTAT=reviewed_by_expert_panel"
	entry, chrom, ok := ParseVCFLine(line)
	assert.True(t, ok)
	assert.Equal(t, "7", chrom)
	assert.Equal(t, int64(140753336), entry.Pos)
	assert.Equal(t, "T", entry.Ref)
	assert.Equal(t, "A", entry.Alt)
	assert.Equal(t, "Pathogenic", entry.ClnSig)
	assert.Equal(t, "reviewed_by_expert_panel", entry.RevStat)
	assert.Equal(t, "Melanoma", entry.ClnDN)
}

func TestParseVCFLineNoCLNSIG(t *testing.T) {
	line := "1\t100\t.\tA\tG\t.\t.\tALLELEID=99"
	_, _, ok := ParseVCFLine(line)
	assert.False(t, ok, "should skip entries without CLNSIG")
}

func TestParseVCFLineMultiAllelic(t *testing.T) {
	line := "1\t100\t.\tA\tG,T\t.\t.\tALLELEID=99;CLNSIG=Benign;CLNREVSTAT=no_assertion"
	entry, _, ok := ParseVCFLine(line)
	assert.True(t, ok)
	assert.Equal(t, "G", entry.Alt, "should take first ALT only")
}

func TestExtractInfo(t *testing.T) {
	info := "ALLELEID=826271;CLNDN=Melanoma;CLNSIG=Pathogenic;CLNREVSTAT=reviewed_by_expert_panel"
	assert.Equal(t, "Pathogenic", ExtractInfo(info, "CLNSIG="))
	assert.Equal(t, "Melanoma", ExtractInfo(info, "CLNDN="))
	assert.Equal(t, "reviewed_by_expert_panel", ExtractInfo(info, "CLNREVSTAT="))
	assert.Equal(t, "", ExtractInfo(info, "MISSING="))
}

func TestNormalizeChrom(t *testing.T) {
	assert.Equal(t, "7", NormalizeChrom("chr7"))
	assert.Equal(t, "7", NormalizeChrom("7"))
	assert.Equal(t, "X", NormalizeChrom("chrX"))
}

func TestTruncate(t *testing.T) {
	assert.Equal(t, "abc", Truncate("abc", 5))
	assert.Equal(t, "ab", Truncate("abcde", 2))
}
