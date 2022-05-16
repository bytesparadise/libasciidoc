package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (r *sgmlRenderer) renderQuoteBlock(ctx *renderer.Context, b *types.DelimitedBlock) (string, error) {
	result := &strings.Builder{}
	content, err := r.renderElements(ctx, b.Elements)
	if err != nil {
		return "", errors.Wrap(err, "unable to render quote block content")
	}
	roles, err := r.renderElementRoles(ctx, b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render quote block roles")
	}
	attribution, err := newAttribution(b)
	if err != nil {
		return "", errors.Wrap(err, "unable to render quote block atribution")
	}
	title, err := r.renderElementTitle(ctx, b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render quote block title")
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
		Title:       title,
		Roles:       roles,
		Attribution: attribution,
		Content:     content,
	})
	return result.String(), err
}

func (r *sgmlRenderer) renderQuoteParagraph(ctx *renderer.Context, p *types.Paragraph) (string, error) {
	log.Debug("rendering quote paragraph...")
	result := &strings.Builder{}

	content, err := r.renderParagraphElements(ctx, p)
	if err != nil {
		return "", errors.Wrap(err, "unable to render quote paragraph lines")
	}
	attribution, err := newAttribution(p)
	if err != nil {
		return "", errors.Wrap(err, "unable to render quote paragraph lines")
	}
	title, err := r.renderElementTitle(ctx, p.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render callout list roles")
	}

	err = r.quoteParagraph.Execute(result, struct {
		Context     *renderer.Context
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

	return result.String(), err
}
