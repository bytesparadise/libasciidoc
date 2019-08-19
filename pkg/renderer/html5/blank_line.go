package html5

import (
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func renderBlankLine(ctx *renderer.Context, l types.BlankLine) ([]byte, error) { //nolint:unparam
	if ctx.IncludeBlankLine() {
		return []byte("\n"), nil
	}
	return make([]byte, 0), nil
}
