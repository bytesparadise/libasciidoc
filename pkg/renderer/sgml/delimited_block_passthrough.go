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
		return "", errors.Wrap(err, "unable to render passthrough block content")
	}
	roles, err := r.renderElementRoles(ctx, b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render passthrough block roles")
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
		Content: strings.Trim(content, "\n"),
	})
	return result.String(), err
}
