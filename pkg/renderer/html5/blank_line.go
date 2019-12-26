package html5

import (
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"

	log "github.com/sirupsen/logrus"
)

func renderBlankLine(ctx *renderer.Context, l types.BlankLine) ([]byte, error) { //nolint:unparam
	if ctx.IncludeBlankLine() {
		log.Debug("rendering blankline")
		return []byte("\n\n"), nil
	}
	return []byte{}, nil
}

func renderLineBreak() ([]byte, error) {
	return []byte("<br>"), nil
}
