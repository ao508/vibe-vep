package signal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexColumns(t *testing.T) {
	header := []string{"Hugo_Symbol", "Chromosome", "Start_Position", "n_impact", "f_impact"}
	col := IndexColumns(header, "Hugo_Symbol", "Chromosome", "Start_Position", "Missing")
	assert.Equal(t, 0, col["Hugo_Symbol"])
	assert.Equal(t, 1, col["Chromosome"])
	assert.Equal(t, 2, col["Start_Position"])
	assert.Equal(t, -1, col["Missing"])
}

func TestFormatFreq(t *testing.T) {
	assert.Equal(t, "0", FormatFreq(0))
	assert.Equal(t, "0.000291511", FormatFreq(0.000291511))
	assert.Equal(t, "0.5", FormatFreq(0.5))
}
