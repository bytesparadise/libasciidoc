package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (r *sgmlRenderer) renderLiteralBlock(ctx *renderer.Context, b *types.DelimitedBlock) (string, error) {
	log.Debugf("rendering literal block")
	content, err := r.renderElements(ctx, b.Elements)
	if err != nil {
		return "", errors.Wrap(err, "unable to render literal block content")
	}
	roles, err := r.renderElementRoles(ctx, b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render literal block roles")
	}
	title, err := r.renderElementTitle(ctx, b.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render literal block title")
	}

	result := &strings.Builder{}
	err = r.literalBlock.Execute(result, struct {
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
	if err != nil {
		return "", errors.Wrap(err, "unable to render literal block")
	}
	return result.String(), nil
}
