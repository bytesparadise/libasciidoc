package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderMarkdownQuoteBlock(ctx *renderer.Context, b *types.DelimitedBlock) (string, error) {
	result := &strings.Builder{}
	content, err := r.renderElements(ctx, b.Elements)
	if err != nil {
		return "", errors.Wrap(err, "unable to render example block content")
	}
	roles, err := r.renderElementRoles(ctx, b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render fenced block content")
	}
	attribution, err := newAttribution(b)
	if err != nil {
		return "", errors.Wrap(err, "unable to render fenced block content")
	}
	title, err := r.renderElementTitle(b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render callout list roles")
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
		Title:       title,
		Roles:       roles,
		Attribution: attribution,
		Content:     content,
	})
	return result.String(), err
}
