package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (r *sgmlRenderer) renderAdmonitionBlock(ctx *renderer.Context, b *types.DelimitedBlock) (string, error) {
	kind, _, err := b.Attributes.GetAsString(types.AttrStyle)
	if err != nil {
		return "", err
	}
	kind = strings.ToLower(kind)
	icon, err := r.renderIcon(ctx, types.Icon{Class: strings.ToLower(kind), Attributes: b.Attributes}, true)
	if err != nil {
		return "", err
	}
	result := &strings.Builder{}
	blocks := discardBlankLines(b.Elements)
	content, err := r.renderElements(ctx, blocks)
	if err != nil {
		return "", errors.Wrap(err, "unable to render admonition block content")
	}
	roles, err := r.renderElementRoles(ctx, b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render fenced block content")
	}
	title, err := r.renderElementTitle(ctx, b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render callout list roles")
	}

	err = r.admonitionBlock.Execute(result, struct {
		Context *renderer.Context
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
	return result.String(), err
}

func (r *sgmlRenderer) renderAdmonitionParagraph(ctx *renderer.Context, p *types.Paragraph) (string, error) {
	log.Debug("rendering admonition paragraph...")
	result := &strings.Builder{}
	kind, ok, err := p.Attributes.GetAsString(types.AttrStyle)
	if err != nil {
		return "", err
	}
	if !ok {
		return "", errors.Errorf("failed to render admonition with unknown kind: %T", p.Attributes[types.AttrStyle])
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
		return "", errors.Wrap(err, "unable to render fenced block content")
	}
	title, err := r.renderElementTitle(ctx, p.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render callout list roles")
	}

	err = r.admonitionParagraph.Execute(result, struct {
		Context *renderer.Context
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

	return result.String(), err
}
