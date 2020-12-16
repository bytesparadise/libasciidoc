package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderListingBlock(ctx *renderer.Context, b types.ListingBlock) (string, error) {
	if k, found := b.Attributes[types.AttrStyle]; found && k == types.Source {
		return r.renderSourceBlock(ctx, b)
	}
	previousWithinDelimitedBlock := ctx.WithinDelimitedBlock
	defer func() {
		ctx.WithinDelimitedBlock = previousWithinDelimitedBlock
	}()
	ctx.WithinDelimitedBlock = true
	result := &strings.Builder{}
	lines := discardEmptyLines(b.Lines)
	content, err := r.renderLines(ctx, lines)
	if err != nil {
		return "", errors.Wrap(err, "unable to render listing block content")
	}
	roles, err := r.renderElementRoles(ctx, b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render listing block roles")
	}

	err = r.listingBlock.Execute(result, struct {
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

func (r *sgmlRenderer) renderListingParagraph(ctx *renderer.Context, p types.Paragraph) (string, error) {
	result := &strings.Builder{}
	content, err := r.renderLines(ctx, p.Lines)
	if err != nil {
		return "", errors.Wrap(err, "unable to render listing block content")
	}
	roles, err := r.renderElementRoles(ctx, p.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render listing block roles")
	}

	err = r.listingBlock.Execute(result, struct {
		Context *renderer.Context
		ID      string
		Title   string
		Roles   string
		Content string
	}{
		Context: ctx,
		ID:      r.renderElementID(p.Attributes),
		Title:   r.renderElementTitle(p.Attributes),
		Roles:   roles,
		Content: content,
	})
	return result.String(), err
}
