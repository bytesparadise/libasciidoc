package renderer

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	log "github.com/sirupsen/logrus"
)

// IncludeTableOfContents includes a Table Of Contents in the document
// if the `toc` attribute is present
func IncludeTableOfContents(ctx Context) Context {
	if t, found := ctx.Document.Attributes.GetAsString(types.AttrTableOfContents); found {
		ctx.Document = insertTableOfContents(ctx.Document, t)
	}
	return ctx
}

func insertTableOfContents(doc types.Document, location string) types.Document {
	log.Debugf("inserting a table of contents at location `%s`", location)
	// insert a TableOfContentsMacro element if `toc` value is:
	// - "auto" (or empty)
	// - "preamble"
	log.Debugf("inserting ToC macro with placement: '%s'", location)
	toc := types.TableOfContentsMacro{}
	switch location {
	case "", "auto":
		// insert TableOfContentsMacro at first position (in section0 if it exists)
		if header, ok := doc.Header(); ok {
			header.Elements = append([]interface{}{toc}, header.Elements...)
			doc.Elements[0] = header
		} else {
			doc.Elements = append([]interface{}{toc}, doc.Elements...)
		}
	case "preamble":
		// lookup preamble in elements (should be first)
		// insert TableOfContentsMacro just after preamble
		if header, ok := doc.Header(); ok {
			if preambleIndex, ok := lookupPreamble(header.Elements); ok {
				header.Elements = insert(header.Elements, toc, preambleIndex)
				doc.Elements[0] = header
			}
		} else if preambleIndex, ok := lookupPreamble(doc.Elements); ok {
			doc.Elements = insert(doc.Elements, toc, preambleIndex)
		}
	// case "macro":
	default:
		log.Warnf("invalid or unsupported value for 'toc' attribute: '%s'", location)
	}
	return doc
}

func lookupPreamble(elements []interface{}) (int, bool) {
	for i, e := range elements {
		if _, ok := e.(types.Preamble); ok {
			return i, true
		}
	}
	return -1, false
}
func insert(elements []interface{}, element interface{}, index int) []interface{} {
	remainingElements := make([]interface{}, len(elements)-(index+1))
	copy(remainingElements, elements[index+1:])
	result := append(elements[0:index+1], element)
	result = append(result, remainingElements...)
	return result
}
