package renderer

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	log "github.com/sirupsen/logrus"
)

// IncludeTableOfContents includes a Table Of Contents in the document
// if the `toc` attribute is present
func IncludeTableOfContents(ctx *Context) {
	if d, found := types.SearchAttributeDeclaration(ctx.Document.Elements, types.AttrTableOfContents); found {
		insertTableOfContents(&ctx.Document, d.Value)
	}
}

func insertTableOfContents(doc *types.Document, location string) {
	log.Debugf("inserting a table of contents at location `%s`", location)
	// insert a TableOfContentsMacro element if `toc` value is:
	// - "auto" (or empty)
	// - "preamble"
	log.Debugf("inserting ToC macro with placement: '%s'", location)
	switch location {
	case "", "auto":
		// insert TableOfContentsMacro at first position (in section0 if it exists)
		if s, ok := doc.Elements[0].(types.Section); ok && s.Level == 0 {
			s.Elements = append([]interface{}{types.TableOfContentsMacro{}}, s.Elements...)
			doc.Elements[0] = s
		} else {
			doc.Elements = append([]interface{}{types.TableOfContentsMacro{}}, doc.Elements...)
		}
	case "preamble":
		// lookup preamble in elements (should be first)
		preambleIndex := 0
		for i, e := range doc.Elements {
			if _, ok := e.(types.Preamble); ok {
				preambleIndex = i
				break
			}
		}
		// insert TableOfContentsMacro just after preamble
		remainingElements := make([]interface{}, len(doc.Elements)-(preambleIndex+1))
		copy(remainingElements, doc.Elements[preambleIndex+1:])
		doc.Elements = append(doc.Elements[0:preambleIndex+1], types.TableOfContentsMacro{})
		doc.Elements = append(doc.Elements, remainingElements...)
	case "macro":
	default:
		log.Warnf("invalid value for 'toc' attribute: '%s'", location)

	}
}
