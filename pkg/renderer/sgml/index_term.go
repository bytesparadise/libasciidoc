package sgml

import (
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func (r *sgmlRenderer) renderIndexTerm(ctx *renderer.Context, t types.IndexTerm) ([]byte, error) {
	return r.renderInlineElements(ctx, t.Term)
}

func (r *sgmlRenderer) renderConcealedIndexTerm(_ types.ConcealedIndexTerm) ([]byte, error) {
	return []byte{}, nil // do not render
}
