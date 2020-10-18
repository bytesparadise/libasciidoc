package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderAdmonitionBlock(ctx *renderer.Context, b types.ExampleBlock) (string, error) {
	kind, _ := b.Attributes[types.AttrAdmonitionKind].(types.AdmonitionKind)
	icon, err := r.renderIcon(ctx, types.Icon{Class: string(kind), Attributes: b.Attributes}, true)
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
	err = r.admonitionBlock.Execute(result, struct {
		Context *renderer.Context
		ID      string
		Title   string
		Kind    types.AdmonitionKind
		Roles   string
		Icon    string
		Content string
	}{
		Context: ctx,
		ID:      r.renderElementID(b.Attributes),
		Kind:    kind,
		Roles:   roles,
		Title:   r.renderElementTitle(b.Attributes),
		Icon:    icon,
		Content: content,
	})
	return result.String(), err
}
