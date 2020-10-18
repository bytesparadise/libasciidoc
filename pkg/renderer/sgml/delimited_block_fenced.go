package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderFencedBlock(ctx *renderer.Context, b types.FencedBlock) (string, error) {
	previousWithinDelimitedBlock := ctx.WithinDelimitedBlock
	previousIncludeBlankLine := ctx.IncludeBlankLine
	defer func() {
		ctx.WithinDelimitedBlock = previousWithinDelimitedBlock
		ctx.IncludeBlankLine = previousIncludeBlankLine
	}()
	ctx.WithinDelimitedBlock = true
	ctx.IncludeBlankLine = true
	result := &strings.Builder{}
	lines := discardEmptyLines(b.Lines)
	content, err := r.renderLines(ctx, lines)
	if err != nil {
		return "", errors.Wrap(err, "unable to render fenced block content")
	}
	roles, err := r.renderElementRoles(ctx, b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render fenced block roles")
	}

	err = r.fencedBlock.Execute(result, struct {
		Context *renderer.Context
		ID      string
		Title   string
		Roles   string
		Content string
	}{
		Context: ctx,
		ID:      r.renderElementID(b.Attributes),
		Title:   r.renderElementTitle(b.Attributes),
		Roles:   roles,
		Content: content,
	})
	return result.String(), err
}
