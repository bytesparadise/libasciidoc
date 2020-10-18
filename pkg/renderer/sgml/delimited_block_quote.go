package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
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

func (r *sgmlRenderer) renderQuoteParagraph(ctx *renderer.Context, p types.Paragraph) (string, error) {
	log.Debug("rendering quote paragraph...")
	result := &strings.Builder{}

	content, err := r.renderLines(ctx, p.Lines)
	if err != nil {
		return "", errors.Wrap(err, "unable to render quote paragraph lines")
	}
	err = r.quoteParagraph.Execute(result, struct {
		Context     *renderer.Context
		ID          string
		Title       string
		Attribution Attribution
		Content     string
		Lines       [][]interface{}
	}{
		Context:     ctx,
		ID:          r.renderElementID(p.Attributes),
		Title:       r.renderElementTitle(p.Attributes),
		Attribution: paragraphAttribution(p),
		Content:     string(content),
		Lines:       p.Lines,
	})

	return result.String(), err
}
