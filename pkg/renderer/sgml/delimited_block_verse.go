package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (r *sgmlRenderer) renderVerseBlock(ctx *context, b *types.DelimitedBlock) (string, error) {
	roles, err := r.renderElementRoles(ctx, b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render verser block roles")
	}
	previousWithinDelimitedBlock := ctx.withinDelimitedBlock
	defer func() {
		ctx.withinDelimitedBlock = previousWithinDelimitedBlock
	}()
	ctx.withinDelimitedBlock = true
	content, err := r.renderElements(ctx, b.Elements)
	if err != nil {
		return "", errors.Wrap(err, "unable to render verse block content")
	}
	attribution := newAttribution(b)
	title, err := r.renderElementTitle(ctx, b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render verse block title")
	}
	return r.execute(r.verseBlock, struct {
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
		Content:     strings.Trim(string(content), "\n"),
	})
}

func (r *sgmlRenderer) renderVerseParagraph(ctx *context, p *types.Paragraph) (string, error) {
	log.Debug("rendering verse paragraph...")
	content, err := RenderParagraphElements(p)
	if err != nil {
		return "", errors.Wrap(err, "unable to render verse paragraph lines")
	}
	attribution := newAttribution(p)
	title, err := r.renderElementTitle(ctx, p.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render callout list roles")
	}
	return r.execute(r.verseParagraph, struct {
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
