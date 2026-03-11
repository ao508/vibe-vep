package annotate

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/inodb/vibe-vep/internal/cache"
	"github.com/inodb/vibe-vep/internal/vcf"
)

// ReverseMapProteinChange maps a protein change (e.g. KRAS G12C) back to
// genomic variant(s) using the canonical transcript's CDS sequence.
func ReverseMapProteinChange(c *cache.Cache, geneName string, refAA byte, protPos int64, altAA byte) ([]*vcf.Variant, error) {
	transcripts := c.FindTranscriptsByGene(geneName)
	if len(transcripts) == 0 {
		return nil, fmt.Errorf("gene %q not found in transcript cache", geneName)
	}

	// Find canonical protein-coding transcript
	var canonical *cache.Transcript
	for _, t := range transcripts {
		if t.IsCanonicalMSK && t.IsProteinCoding() {
			canonical = t
			break
		}
	}
	if canonical == nil {
		// Fall back to first protein-coding transcript
		for _, t := range transcripts {
			if t.IsProteinCoding() {
				canonical = t
				break
			}
		}
	}
	if canonical == nil {
		return nil, fmt.Errorf("no protein-coding transcript found for gene %q", geneName)
	}

	return reverseMapProtein(canonical, refAA, protPos, altAA)
}

// reverseMapProtein maps a protein change on a specific transcript to genomic variant(s).
func reverseMapProtein(t *cache.Transcript, refAA byte, protPos int64, altAA byte) ([]*vcf.Variant, error) {
	if len(t.CDSSequence) == 0 {
		return nil, fmt.Errorf("transcript %s has no CDS sequence", t.ID)
	}

	// Get reference codon
	refCodon := GetCodon(t.CDSSequence, protPos)
	if len(refCodon) != 3 {
		return nil, fmt.Errorf("protein position %d out of range for transcript %s", protPos, t.ID)
	}

	// Verify reference AA matches
	actualRefAA := TranslateCodon(refCodon)
	if actualRefAA != refAA {
		return nil, fmt.Errorf("reference amino acid mismatch at position %d: expected %c, got %c in transcript %s",
			protPos, refAA, actualRefAA, t.ID)
	}

	// CDS start position for this codon (1-based)
	cdsStart := (protPos-1)*3 + 1

	// Enumerate all single-base mutations of the codon that produce altAA
	var variants []*vcf.Variant
	for posInCodon := 0; posInCodon < 3; posInCodon++ {
		for _, base := range "ACGT" {
			if byte(base) == refCodon[posInCodon] {
				continue // skip ref base
			}
			mutCodon := MutateCodon(refCodon, posInCodon, byte(base))
			if TranslateCodon(mutCodon) != altAA {
				continue
			}

			cdsPos := cdsStart + int64(posInCodon)
			genomicPos := CDSToGenomic(cdsPos, t)
			if genomicPos == 0 {
				continue
			}

			refBase := string(refCodon[posInCodon])
			altBase := string(base)

			// Account for reverse strand
			if t.IsReverseStrand() {
				refBase = string(Complement(refCodon[posInCodon]))
				altBase = string(Complement(byte(base)))
			}

			variants = append(variants, &vcf.Variant{
				Chrom: t.Chrom,
				Pos:   genomicPos,
				Ref:   refBase,
				Alt:   altBase,
			})
		}
	}

	if len(variants) == 0 {
		return nil, fmt.Errorf("no single-base mutation of codon %s at position %d produces %c",
			refCodon, protPos, altAA)
	}

	return variants, nil
}

// reCDSChange parses a simple CDS substitution like "35G>T".
var reCDSChange = regexp.MustCompile(`^(\d+)([ACGT])>([ACGT])$`)

// reCDSDeletion parses a CDS deletion like "923del" or "100_102del".
var reCDSDeletion = regexp.MustCompile(`^(\d+)(?:_(\d+))?del$`)

// ReverseMapHGVSc maps an HGVSc notation (e.g. KRAS c.35G>T or KRAS c.34del)
// back to a genomic variant using the transcript's CDS-to-genomic mapping.
func ReverseMapHGVSc(c *cache.Cache, geneOrTranscript string, cdsChange string) ([]*vcf.Variant, error) {
	// Try substitution first
	if m := reCDSChange.FindStringSubmatch(cdsChange); m != nil {
		return reverseMapHGVScSubstitution(c, geneOrTranscript, m)
	}

	// Try deletion
	if m := reCDSDeletion.FindStringSubmatch(cdsChange); m != nil {
		return reverseMapHGVScDeletion(c, geneOrTranscript, m)
	}

	return nil, fmt.Errorf("unsupported CDS change notation %q (supported: substitutions like 35G>T, deletions like 923del or 100_102del)", cdsChange)
}

// findHGVScTranscript resolves a gene name or transcript ID to a transcript.
func findHGVScTranscript(c *cache.Cache, geneOrTranscript string) (*cache.Transcript, error) {
	if strings.HasPrefix(geneOrTranscript, "ENST") {
		transcript := c.GetTranscript(geneOrTranscript)
		if transcript == nil {
			transcripts := findTranscriptByPrefix(c, geneOrTranscript)
			if len(transcripts) > 0 {
				transcript = transcripts[0]
			}
		}
		if transcript == nil {
			return nil, fmt.Errorf("transcript %q not found", geneOrTranscript)
		}
		return transcript, nil
	}

	transcripts := c.FindTranscriptsByGene(geneOrTranscript)
	if len(transcripts) == 0 {
		return nil, fmt.Errorf("gene %q not found in transcript cache", geneOrTranscript)
	}
	for _, t := range transcripts {
		if t.IsCanonicalMSK && t.IsProteinCoding() {
			return t, nil
		}
	}
	for _, t := range transcripts {
		if t.IsProteinCoding() {
			return t, nil
		}
	}
	return nil, fmt.Errorf("no protein-coding transcript found for gene %q", geneOrTranscript)
}

// reverseMapHGVScSubstitution handles simple CDS substitutions like "35G>T".
func reverseMapHGVScSubstitution(c *cache.Cache, geneOrTranscript string, m []string) ([]*vcf.Variant, error) {
	cdsPos, _ := strconv.ParseInt(m[1], 10, 64)
	cdsRef := m[2]
	cdsAlt := m[3]

	transcript, err := findHGVScTranscript(c, geneOrTranscript)
	if err != nil {
		return nil, err
	}

	// Verify CDS ref base
	if len(transcript.CDSSequence) > 0 && cdsPos <= int64(len(transcript.CDSSequence)) {
		actualRef := string(transcript.CDSSequence[cdsPos-1])
		if actualRef != cdsRef {
			return nil, fmt.Errorf("CDS reference mismatch at position %d: expected %s, got %s in transcript %s",
				cdsPos, cdsRef, actualRef, transcript.ID)
		}
	}

	genomicPos := CDSToGenomic(cdsPos, transcript)
	if genomicPos == 0 {
		return nil, fmt.Errorf("CDS position %d could not be mapped to genomic coordinates in transcript %s",
			cdsPos, transcript.ID)
	}

	ref := cdsRef
	alt := cdsAlt
	if transcript.IsReverseStrand() {
		ref = string(Complement(cdsRef[0]))
		alt = string(Complement(cdsAlt[0]))
	}

	return []*vcf.Variant{{
		Chrom: transcript.Chrom,
		Pos:   genomicPos,
		Ref:   ref,
		Alt:   alt,
	}}, nil
}

// reverseMapHGVScDeletion handles CDS deletions like "923del" or "100_102del".
func reverseMapHGVScDeletion(c *cache.Cache, geneOrTranscript string, m []string) ([]*vcf.Variant, error) {
	cdsStart, _ := strconv.ParseInt(m[1], 10, 64)
	cdsEnd := cdsStart
	if m[2] != "" {
		cdsEnd, _ = strconv.ParseInt(m[2], 10, 64)
	}
	if cdsEnd < cdsStart {
		return nil, fmt.Errorf("invalid CDS deletion range: %d_%d", cdsStart, cdsEnd)
	}

	transcript, err := findHGVScTranscript(c, geneOrTranscript)
	if err != nil {
		return nil, err
	}

	// Extract deleted CDS bases
	if cdsEnd > int64(len(transcript.CDSSequence)) {
		return nil, fmt.Errorf("CDS position %d out of range for transcript %s (CDS length %d)",
			cdsEnd, transcript.ID, len(transcript.CDSSequence))
	}
	deletedBases := transcript.CDSSequence[cdsStart-1 : cdsEnd]

	// Map CDS positions to genomic
	genomicStart := CDSToGenomic(cdsStart, transcript)
	genomicEnd := CDSToGenomic(cdsEnd, transcript)
	if genomicStart == 0 || genomicEnd == 0 {
		return nil, fmt.Errorf("CDS positions %d-%d could not be mapped to genomic coordinates in transcript %s",
			cdsStart, cdsEnd, transcript.ID)
	}

	// Build VCF-convention deletion (padding base + deleted bases)
	if transcript.IsReverseStrand() {
		// Reverse strand: genomicStart > genomicEnd (reversed coords)
		// Deleted bases on genomic strand are reverse-complemented
		// Padding base is after the deletion in genomic coords (genomicEnd+1 is upstream)
		// but VCF convention uses the base before the deletion in genomic order
		// genomicEnd is the lower coordinate, so pad at genomicEnd-1
		padPos := genomicEnd - 1
		if padPos < 1 {
			return nil, fmt.Errorf("cannot add padding base before genomic position %d", genomicEnd)
		}
		// Get padding base from the genomic strand
		// On reverse strand, CDS pos before cdsStart maps to genomicStart+1
		// But we need the genomic base at padPos, which is genomicEnd-1
		// We can get it from CDS: the base after cdsEnd in CDS maps to genomicEnd-1 on reverse strand
		padCDSPos := cdsEnd + 1
		if padCDSPos > int64(len(transcript.CDSSequence)) {
			return nil, fmt.Errorf("cannot determine padding base: CDS position %d out of range", padCDSPos)
		}
		padBase := string(Complement(transcript.CDSSequence[padCDSPos-1]))

		// Deleted bases reverse-complemented for genomic strand
		genomicDeleted := ReverseComplement(string(deletedBases))

		return []*vcf.Variant{{
			Chrom: transcript.Chrom,
			Pos:   padPos,
			Ref:   padBase + genomicDeleted,
			Alt:   padBase,
		}}, nil
	}

	// Forward strand: padding base is at genomicStart-1
	padPos := genomicStart - 1
	if padPos < 1 {
		return nil, fmt.Errorf("cannot add padding base before genomic position %d", genomicStart)
	}
	// Get padding base from CDS (base before cdsStart)
	padCDSPos := cdsStart - 1
	if padCDSPos < 1 {
		return nil, fmt.Errorf("cannot determine padding base: CDS position %d out of range", padCDSPos)
	}
	padBase := string(transcript.CDSSequence[padCDSPos-1])

	return []*vcf.Variant{{
		Chrom: transcript.Chrom,
		Pos:   padPos,
		Ref:   padBase + string(deletedBases),
		Alt:   padBase,
	}}, nil
}

// findTranscriptByPrefix finds transcripts matching an ID prefix (without version).
func findTranscriptByPrefix(c *cache.Cache, prefix string) []*cache.Transcript {
	// Strip version suffix from prefix if present for matching
	base := prefix
	if idx := strings.IndexByte(prefix, '.'); idx >= 0 {
		base = prefix[:idx]
	}

	var matches []*cache.Transcript
	for _, chrom := range c.Chromosomes() {
		for _, t := range c.FindTranscriptsByChrom(chrom) {
			tid := t.ID
			if idx := strings.IndexByte(tid, '.'); idx >= 0 {
				tid = tid[:idx]
			}
			if tid == base {
				matches = append(matches, t)
			}
		}
	}
	return matches
}
