package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderListingBlock(ctx *renderer.Context, b *types.DelimitedBlock) (string, error) {
	if k, found := b.Attributes[types.AttrStyle]; found && k == types.Source {
		return r.renderSourceBlock(ctx, b)
	}
	previousWithinDelimitedBlock := ctx.WithinDelimitedBlock
	defer func() {
		ctx.WithinDelimitedBlock = previousWithinDelimitedBlock
	}()
	ctx.WithinDelimitedBlock = true
	result := &strings.Builder{}
	content, err := r.renderElements(ctx, b.Elements)
	if err != nil {
		return "", errors.Wrap(err, "unable to render listing block content")
	}
	roles, err := r.renderElementRoles(ctx, b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render listing block roles")
	}
	title, err := r.renderElementTitle(ctx, b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render listing block title")
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
		Title:   title,
		Roles:   roles,
		Content: strings.Trim(content, "\n"),
	})
	return result.String(), err
}

func (r *sgmlRenderer) renderListingParagraph(ctx *renderer.Context, p *types.Paragraph) (string, error) {
	result := &strings.Builder{}
	content, err := r.renderElements(ctx, p.Elements)
	if err != nil {
		return "", errors.Wrap(err, "unable to render listing block content")
	}
	roles, err := r.renderElementRoles(ctx, p.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render listing block roles")
	}
	title, err := r.renderElementTitle(ctx, p.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render callout list roles")
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
		Title:   title,
		Roles:   roles,
		Content: content,
	})
	return result.String(), err
}
