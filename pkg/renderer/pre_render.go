package renderer

import (
	"github.com/pkg/errors"
)

// Prerender runs the pre-rendering phase, with the following steps (if needed/applicable):
// - process file inclusions
// - wraps elements in a preamble
// - generated the ToC
func Prerender(ctx *Context) error {
	err := ProcessFileInclusions(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to pre-render document")
	}
	IncludePreamble(ctx)
	IncludeTableOfContents(ctx)
	return nil
}
