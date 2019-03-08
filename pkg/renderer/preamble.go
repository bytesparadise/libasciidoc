package renderer

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	log "github.com/sirupsen/logrus"
)

// IncludePreamble wraps all document elements before the first section in a `Preamble`,
// unless the document has no section. Returns a new document with the changes.
func IncludePreamble(doc types.Document) (types.Document, error) {
	doc.Elements = insertPreamble(doc.Elements)
	return doc, nil
}

func insertPreamble(blocks []interface{}) []interface{} {
	log.Debugf("generating preamble from %d blocks", len(blocks))
	preamble := types.NewEmptyPreamble()
	for _, block := range blocks {
		switch block.(type) {
		case types.Section:
			break
		default:
			preamble.Elements = append(preamble.Elements, block)
		}
	}
	// no element in the preamble, or no section in the document, so no preamble to generate
	if len(preamble.Elements) == 0 || len(preamble.Elements) == len(blocks) {
		log.Debugf("skipping preamble (%d vs %d)", len(preamble.Elements), len(blocks))
		return types.NilSafe(blocks)
	}
	// now, insert the preamble instead of the 'n' blocks that belong to the preamble
	// and copy the other items
	result := make([]interface{}, len(blocks)-len(preamble.Elements)+1)
	result[0] = preamble
	copy(result[1:], blocks[len(preamble.Elements):])
	log.Debugf("generated preamble with %d blocks", len(preamble.Elements))
	return result
}
