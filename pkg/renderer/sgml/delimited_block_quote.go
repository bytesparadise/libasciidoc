package sgml

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (r *sgmlRenderer) renderQuoteBlock(ctx *context, b *types.DelimitedBlock) (string, error) {
	content, err := r.renderElements(ctx, b.Elements)
	if err != nil {
		return "", errors.Wrap(err, "unable to render quote block content")
	}
	roles, err := r.renderElementRoles(ctx, b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render quote block roles")
	}
	attribution := newAttribution(b)
	title, err := r.renderElementTitle(ctx, b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render quote block title")
	}
	return r.execute(r.quoteBlock, struct {
		Context     *context
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
}

func (r *sgmlRenderer) renderQuoteParagraph(ctx *context, p *types.Paragraph) (string, error) {
	log.Debug("rendering quote paragraph...")
	content, err := r.renderParagraphElements(ctx, p)
	if err != nil {
		return "", errors.Wrap(err, "unable to render quote paragraph lines")
	}
	attribution := newAttribution(p)
	title, err := r.renderElementTitle(ctx, p.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render callout list roles")
	}
	return r.execute(r.quoteParagraph, struct {
		Context     *context
		ID          string
		Title       string
		Attribution Attribution
		Content     string
	}{
		Context:     ctx,
		ID:          r.renderElementID(p.Attributes),
		Title:       title,
		Attribution: attribution,
		Content:     string(content),
	})
}
