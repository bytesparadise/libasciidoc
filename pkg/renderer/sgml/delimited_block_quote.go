package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderQuoteBlock(ctx *renderer.Context, b types.QuoteBlock) (string, error) {
	result := &strings.Builder{}
	content, err := r.renderElements(ctx, b.Elements)
	if err != nil {
		return "", errors.Wrap(err, "unable to render example block content")
	}
	roles, err := r.renderElementRoles(ctx, b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render fenced block content")
	}
	err = r.quoteBlock.Execute(result, struct {
		Context     *renderer.Context
		ID          string
		Title       string
		Roles       string
		Attribution Attribution
		Content     string
	}{
		Context:     ctx,
		ID:          r.renderElementID(b.Attributes),
		Title:       r.renderElementTitle(b.Attributes),
		Roles:       roles,
		Attribution: quoteBlockAttribution(b),
		Content:     content,
	})
	return result.String(), err
}
