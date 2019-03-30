package renderer

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	log "github.com/sirupsen/logrus"
)

// IncludeTableOfContents includes a Table Of Contents in the document
// if the `toc` attribute is present
func IncludeTableOfContents(ctx *Context) {
	if location, ok := ctx.Document.Attributes.GetAsString(types.AttrTableOfContents); ok {
		ctx.Document.Elements = insertTableOfContents(ctx.Document.Elements, location)
	}
}
func insertTableOfContents(elements []interface{}, location string) []interface{} {
	log.Debugf("inserting a table of contents at location `%s`", location)
	result := []interface{}{}
	// insert a TableOfContentsMacro element if `toc` value is:
	// - "auto" (or empty)
	// - "preamble"
	log.Debugf("inserting ToC macro with placement: '%s'", location)
	switch location {
	case "", "auto":
		// insert TableOfContentsMacro at first position
		result = append([]interface{}{types.TableOfContentsMacro{}}, elements...)
	case "preamble":
		// lookup preamble in elements (should be first)
		preambleIndex := 0
		for i, e := range result {
			if _, ok := e.(types.Preamble); ok {
				preambleIndex = i
				break
			}
		}
		// insert TableOfContentsMacro just after preamble
		remainingElements := make([]interface{}, len(elements)-(preambleIndex+1))
		copy(remainingElements, elements[preambleIndex+1:])
		result = append(elements[0:preambleIndex+1], types.TableOfContentsMacro{})
		result = append(result, remainingElements...)
	case "macro":
	default:
		log.Warnf("invalid value for 'toc' attribute: '%s'", location)

	}

	return result
}
