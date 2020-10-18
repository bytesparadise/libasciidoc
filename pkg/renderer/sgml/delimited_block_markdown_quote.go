package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderMarkdownQuoteBlock(ctx *renderer.Context, b types.MarkdownQuoteBlock) (string, error) {
	result := &strings.Builder{}
	content, err := r.renderLines(ctx, b.Lines)
	if err != nil {
		return "", errors.Wrap(err, "unable to render example block content")
	}
	roles, err := r.renderElementRoles(ctx, b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render fenced block content")
	}
	err = r.markdownQuoteBlock.Execute(result, struct {
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
		Attribution: markdownQuoteBlockAttribution(b),
		Content:     content,
	})
	return result.String(), err
}
