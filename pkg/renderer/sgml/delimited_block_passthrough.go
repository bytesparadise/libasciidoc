package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderPassthroughBlock(ctx *renderer.Context, b *types.DelimitedBlock) (string, error) {
	result := &strings.Builder{}
	// lines := discardEmptyLines(b.Lines)
	previousWithinDelimitedBlock := ctx.WithinDelimitedBlock
	defer func() {
		ctx.WithinDelimitedBlock = previousWithinDelimitedBlock
	}()
	ctx.WithinDelimitedBlock = true
	content, err := r.renderElements(ctx, b.Elements)
	if err != nil {
		return "", errors.Wrap(err, "unable to render passthrough")
	}
	roles, err := r.renderElementRoles(ctx, b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render fenced block content")
	}
	err = r.passthroughBlock.Execute(result, struct {
		Context *renderer.Context
		ID      string
		Roles   string
		Content string
	}{
		Context: ctx,
		ID:      r.renderElementID(b.Attributes),
		Roles:   roles,
		Content: content,
	})
	return result.String(), err
}
