package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderSidebarBlock(ctx *renderer.Context, b *types.DelimitedBlock) (string, error) {
	result := &strings.Builder{}

	blocks := discardBlankLines(b.Elements)
	content, err := r.renderElements(ctx, blocks)
	if err != nil {
		return "", errors.Wrap(err, "unable to render sidebar block content")
	}
	roles, err := r.renderElementRoles(ctx, b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render sidebar block roles")
	}
	title, err := r.renderElementTitle(ctx, b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render sidebar block title")
	}

	err = r.sidebarBlock.Execute(result, struct {
		Context *renderer.Context
		ID      string
		Title   string
		Roles   string
		Content string
	}{
		Context: ctx,
		ID:      r.renderElementID(b.Attributes),
		Title:   title,
		Roles:   roles,
		Content: content,
	})
	return result.String(), err
}
