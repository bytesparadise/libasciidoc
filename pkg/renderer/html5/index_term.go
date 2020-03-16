package html5

import (
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func renderIndexTerm(ctx renderer.Context, t types.IndexTerm) ([]byte, error) {
	return renderElements(ctx, t.Term)
}
