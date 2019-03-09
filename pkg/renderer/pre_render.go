package renderer

import (
	"github.com/pkg/errors"
)

// Prerender runs the pre-rendering phase, with the following steps (if needed/applicable):
// - process file inclusions
// - wraps elements in a preamble
// - generated the ToC
func Prerender(ctx *Context) error {
	doc, err := ProcessFileInclusions(ctx.Document)
	if err != nil {
		return errors.Wrap(err, "failed to pre-render document")
	}
	doc, err = IncludePreamble(doc)
	if err != nil {
		return errors.Wrap(err, "failed to pre-render document")
	}
	doc, err = IncludeTableOfContents(doc)
	if err != nil {
		return errors.Wrap(err, "failed to pre-render document")
	}
	ctx.Document = doc
	return nil
}
