package parser

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	log "github.com/sirupsen/logrus"
)

// IncludeTableOfContentsPlaceHolder includes a `TableOfContentsPlaceHolder` block in the document
// if the `toc` attribute is present
func includeTableOfContentsPlaceHolder(doc types.Document) types.Document {
	if t, found := doc.Attributes[types.AttrTableOfContents]; found {
		doc = doInsertTableOfContentsPlaceHolder(doc, t)
	}
	return doc
}

func doInsertTableOfContentsPlaceHolder(doc types.Document, location interface{}) types.Document {
	log.Debugf("inserting a table of contents at location `%s`", location)
	// insert a TableOfContentsPlaceHolder element if `toc` value is:
	// - "auto" (or `nil`)
	// - "preamble"
	log.Debugf("inserting ToC macro with placement: '%s'", location)
	toc := types.TableOfContentsPlaceHolder{}
	switch location {
	case "auto", nil:
		// insert TableOfContentsPlaceHolder at first position (in section '0' if it exists)
		if header, ok := doc.Header(); ok {
			header.Elements = append([]interface{}{toc}, header.Elements...)
			doc.Elements[0] = header
		} else {
			doc.Elements = append([]interface{}{toc}, doc.Elements...)
		}
	case "preamble":
		// lookup preamble in elements (should be first)
		// insert TableOfContentsPlaceHolder just after preamble
		if header, ok := doc.Header(); ok {
			if preambleIndex, ok := lookupPreamble(header.Elements); ok {
				header.Elements = insertAt(header.Elements, toc, preambleIndex)
				doc.Elements[0] = header
			}
		} else if preambleIndex, ok := lookupPreamble(doc.Elements); ok {
			doc.Elements = insertAt(doc.Elements, toc, preambleIndex)
		}
	// case "macro":
	default:
		log.Warnf("invalid or unsupported value for 'toc' attribute: '%s'", location)
	}
	return doc
}

// returns the index of the preamble if it was found in the given elements
func lookupPreamble(elements []interface{}) (int, bool) {
	for i, e := range elements {
		if _, ok := e.(types.Preamble); ok {
			return i, true
		}
	}
	return -1, false
}

// inserts the given element at the given index
func insertAt(elements []interface{}, element interface{}, index int) []interface{} {
	remainingElements := make([]interface{}, len(elements)-(index+1))
	copy(remainingElements, elements[index+1:])
	result := append(elements[0:index+1], element)
	result = append(result, remainingElements...)
	return result
}
