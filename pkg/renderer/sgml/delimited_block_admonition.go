package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (r *sgmlRenderer) renderAdmonitionBlock(ctx *context, b *types.DelimitedBlock) (string, error) {
	kind, _ := b.Attributes.GetAsString(types.AttrStyle)
	kind = strings.ToLower(kind)
	icon, err := r.renderIcon(ctx, types.Icon{Class: strings.ToLower(kind), Attributes: b.Attributes}, true)
	if err != nil {
		return "", err
	}
	blocks := discardBlankLines(b.Elements)
	content, err := r.renderElements(ctx, blocks)
	if err != nil {
		return "", errors.Wrap(err, "unable to render admonition block content")
	}
	roles, err := r.renderElementRoles(ctx, b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render admonition block roles")
	}
	title, err := r.renderElementTitle(ctx, b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render admonition block title")
	}
	return r.execute(r.admonitionBlock, struct {
		Context *context
		ID      string
		Title   string
		Kind    string
		Roles   string
		Icon    string
		Content string
	}{
		Context: ctx,
		ID:      r.renderElementID(b.Attributes),
		Kind:    kind,
		Roles:   roles,
		Title:   title,
		Icon:    icon,
		Content: content,
	})
}

func (r *sgmlRenderer) renderAdmonitionParagraph(ctx *context, p *types.Paragraph) (string, error) {
	log.Debug("rendering admonition paragraph...")
	kind, found := p.Attributes.GetAsString(types.AttrStyle)
	if !found {
		return "", errors.Errorf("failed to render admonition paragraph with unknown kind: %T", p.Attributes[types.AttrStyle])
	}
	kind = strings.ToLower(kind)
	icon, err := r.renderIcon(ctx, types.Icon{Class: kind, Attributes: p.Attributes}, true)
	if err != nil {
		return "", err
	}
	content, err := r.renderParagraphElements(ctx, p)
	if err != nil {
		return "", err
	}
	roles, err := r.renderElementRoles(ctx, p.Attributes)
	if err != nil {
		return "", err
	}
	title, err := r.renderElementTitle(ctx, p.Attributes)
	if err != nil {
		return "", err
	}
	return r.execute(r.admonitionParagraph, struct {
		Context *context
		ID      string
		Title   string
		Roles   string
		Icon    string
		Kind    string
		Content string
	}{
		Context: ctx,
		ID:      r.renderElementID(p.Attributes),
		Title:   title,
		Kind:    kind,
		Roles:   roles,
		Icon:    icon,
		Content: content,
	})
}
