package hotspots

import (
	"strings"

	"github.com/inodb/vibe-vep/internal/annotate"
	"github.com/inodb/vibe-vep/internal/vcf"
)

// Pre-built keys for Extra map (avoids string concatenation per annotation).
const (
	extraKeyHotspot = "hotspots.hotspot"
	extraKeyType    = "hotspots.type"
	extraKeyQValue  = "hotspots.qvalue"
)

// Source implements annotate.AnnotationSource for cancer hotspots.
type Source struct {
	store *Store
}

// NewSource creates an AnnotationSource backed by the given Store.
func NewSource(store *Store) *Source {
	return &Source{store: store}
}

func (s *Source) Name() string                   { return "hotspots" }
func (s *Source) Version() string                 { return "v2" }
func (s *Source) MatchLevel() annotate.MatchLevel { return annotate.MatchProteinPosition }
func (s *Source) Store() *Store                   { return s.store }

func (s *Source) Columns() []annotate.ColumnDef {
	return []annotate.ColumnDef{
		{Name: "hotspot", Description: "Y if position is a known cancer hotspot"},
		{Name: "type", Description: "Hotspot type: single residue, in-frame indel, 3d, splice"},
		{Name: "qvalue", Description: "Statistical significance (q-value)"},
	}
}

// Annotate marks annotations whose transcript+protein position is a known hotspot.
// Hotspot positions are defined per transcript, so the annotation's transcript must
// match the hotspot's transcript for the position to be meaningful.
func (s *Source) Annotate(_ *vcf.Variant, anns []*annotate.Annotation) {
	for _, ann := range anns {
		if ann.ProteinPosition == 0 || ann.TranscriptID == "" {
			continue
		}
		// Strip version suffix (e.g. "ENST00000311936.8" → "ENST00000311936")
		txID := ann.TranscriptID
		if i := strings.IndexByte(txID, '.'); i >= 0 {
			txID = txID[:i]
		}
		if h, ok := s.store.Lookup(txID, ann.ProteinPosition); ok {
			ann.SetExtraKey(extraKeyHotspot, "Y")
			if h.Type != "" {
				ann.SetExtraKey(extraKeyType, h.Type)
			}
			ann.SetExtraKey(extraKeyQValue, FormatQValue(h.QValue))
		}
	}
}
