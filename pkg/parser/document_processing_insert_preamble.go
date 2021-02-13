package parser

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

// includePreamble wraps all document elements before the first section in a `Preamble`,
// unless the document has no section. Returns a new document with the changes.
func includePreamble(doc types.Document) types.Document {
	if header, ok := doc.Header(); ok {
		header.Elements = doInsertPreamble(header.Elements)
		doc.Elements[0] = header // need to update the header in the parent doc as we don't use pointers here.
	} else {
		doc.Elements = doInsertPreamble(doc.Elements)
	}
	return doc
}

func doInsertPreamble(blocks []interface{}) []interface{} {
	preamble := types.Preamble{
		Elements: make([]interface{}, 0, len(blocks)),
	}
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
		return blocks
	}
	// now, insert the preamble instead of the 'n' blocks that belong to the preamble
	// and copy the other items
	result := make([]interface{}, len(blocks)-len(preamble.Elements)+1)
	result[0] = preamble
	copy(result[1:], blocks[len(preamble.Elements):])
	return result
}
