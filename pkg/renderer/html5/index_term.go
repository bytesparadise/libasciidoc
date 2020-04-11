package html5

import (
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func renderIndexTerm(ctx renderer.Context, t types.IndexTerm) ([]byte, error) {
	return renderInlineElements(ctx, t.Term)
}

func renderConcealedIndexTerm(_ renderer.Context, _ types.ConcealedIndexTerm) ([]byte, error) {
	return []byte{}, nil // do not render
}
