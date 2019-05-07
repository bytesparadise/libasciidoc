package renderer

import (
	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
)

// Prerender runs the pre-rendering phase, with the following steps (if needed/applicable):
// - process file inclusions
// - wraps elements in a preamble
// - generated the ToC
func Prerender(ctx *Context) error {
	IncludePreamble(ctx)
	IncludeTableOfContents(ctx)
	ProcessDocumentHeader(ctx)
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debug("pre-rendered document:")
		spew.Dump(ctx.Document)
	}
	return nil
}
