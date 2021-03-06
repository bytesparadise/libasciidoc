package sgml

import (
	"io"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderCalloutList(ctx *renderer.Context, l types.CalloutList) (string, error) {
	result := &strings.Builder{}
	content := &strings.Builder{}

	for _, item := range l.Items {

		err := r.renderCalloutListItem(ctx, content, item)
		if err != nil {
			return "", errors.Wrap(err, "unable to render callout list item")
		}
	}
	roles, err := r.renderElementRoles(ctx, l.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render callout list roles")
	}
	title, err := r.renderElementTitle(l.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render callout list roles")
	}

	err = r.calloutList.Execute(result, struct {
		Context *renderer.Context
		ID      string
		Title   string
		Roles   string
		Content string
		Items   []types.CalloutListItem
	}{
		Context: ctx,
		ID:      r.renderElementID(l.Attributes),
		Title:   title,
		Roles:   roles,
		Content: string(content.String()),
		Items:   l.Items,
	})
	if err != nil {
		return "", errors.Wrap(err, "unable to render callout list")
	}
	return result.String(), nil
}

func (r *sgmlRenderer) renderCalloutListItem(ctx *renderer.Context, w io.Writer, item types.CalloutListItem) error {

	content, err := r.renderListElements(ctx, item.Elements)
	if err != nil {
		return errors.Wrap(err, "unable to render callout list item content")
	}
	err = r.calloutListItem.Execute(w, struct {
		Context *renderer.Context
		Ref     int
		Content string
	}{
		Context: ctx,
		Ref:     item.Ref,
		Content: string(content),
	})
	if err != nil {
		return errors.Wrap(err, "unable to render callout list")
	}
	return nil
}
