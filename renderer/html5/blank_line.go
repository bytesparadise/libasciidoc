package html5

import (
	"github.com/bytesparadise/libasciidoc/renderer"
	"github.com/bytesparadise/libasciidoc/types"
	log "github.com/sirupsen/logrus"
)

func renderBlankLine(ctx *renderer.Context, p types.BlankLine) ([]byte, error) {
	if ctx.RenderBlankLine() {
		log.Debugf("rendering blankline")
		return []byte("\n"), nil
	}
	return make([]byte, 0), nil
}
