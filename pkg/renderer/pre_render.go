package renderer

import (
	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
)

// Prerender runs the pre-rendering phase, with the following steps (if needed/applicable):
// - wraps elements in a preamble
// - generates the ToC
// - processes the document headers (added in the document attributes)
func Prerender(ctx Context) error {
	ctx = IncludePreamble(ctx)
	ctx = IncludeTableOfContentsPlaceHolder(ctx)
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debug("pre-rendered document:")
		spew.Dump(ctx.Document)
	}
	return nil
}
