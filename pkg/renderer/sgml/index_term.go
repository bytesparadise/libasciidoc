package sgml

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func (r *sgmlRenderer) renderIndexTerm(ctx *context, t *types.IndexTerm) (string, error) {
	return r.renderInlineElements(ctx, t.Term)
}

func (r *sgmlRenderer) renderConcealedIndexTerm(_ *types.ConcealedIndexTerm) (string, error) {
	return "", nil // do not render in SGML
}
