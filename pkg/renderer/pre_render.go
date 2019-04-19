package renderer

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
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
	// TODO: IncludeAuthors: process author names as document atributes, based on the `Section0.Attributes[AttrAuthors]` (eg: 'firstname', firstname_2' etc.)
	// TODO: IncludeRevision: same logic with `Section0.Attributes[AttrRevision]`
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debug("pre-rendered document:")
		spew.Dump(ctx.Document)
	}
	return nil
}
