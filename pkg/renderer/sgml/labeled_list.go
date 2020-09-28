package sgml

import (
	"io"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderLabeledList(ctx *renderer.Context, l types.LabeledList) (string, error) {
	tmpl, itemTmpl, err := r.getLabeledListTmpl(l)
	if err != nil {
		return "", errors.Wrap(err, "unable to render labeled list")
	}

	content := &strings.Builder{}
	cont := false
	for _, item := range l.Items {
		if cont, err = r.renderLabeledListItem(ctx, itemTmpl, content, cont, item); err != nil {
			return "", errors.Wrap(err, "unable to render unordered list")
		}
	}
	roles, err := r.renderElementRoles(ctx, l.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render labeled list roles")
	}
	result := &strings.Builder{}
	// here we must preserve the HTML tags
	err = tmpl.Execute(result, struct {
		Context *renderer.Context
		ID      string
		Title   string
		Roles   string
		Content string
		Items   []types.LabeledListItem
	}{
		Context: ctx,
		ID:      r.renderElementID(l.Attributes),
		Title:   r.renderElementTitle(l.Attributes),
		Roles:   roles,
		Content: string(content.String()),
		Items:   l.Items,
	})
	if err != nil {
		return "", errors.Wrap(err, "unable to render labeled list")
	}
	// log.Debugf("rendered labeled list: %s", result.Bytes())
	return result.String(), nil
}

func (r *sgmlRenderer) getLabeledListTmpl(l types.LabeledList) (*textTemplate, *textTemplate, error) {
	if layout, ok := l.Attributes[types.AttrStyle]; ok {
		switch layout {
		case "qanda":
			return r.qAndAList, r.qAndAListItem, nil
		case "horizontal":
			return r.labeledListHorizontal, r.labeledListHorizontalItem, nil
		default:
			return nil, nil, errors.Errorf("unsupported labeled list layout: %s", layout)
		}
	}
	return r.labeledList, r.labeledListItem, nil
}

func (r *sgmlRenderer) renderLabeledListItem(ctx *renderer.Context, tmpl *textTemplate, w io.Writer, continuation bool, item types.LabeledListItem) (bool, error) {

	term, err := r.renderInlineElements(ctx, item.Term)
	if err != nil {
		return false, errors.Wrap(err, "unable to render labeled list term")
	}
	content, err := r.renderListElements(ctx, item.Elements)
	if err != nil {
		return false, errors.Wrap(err, "unable to render labeled list content")
	}
	err = tmpl.Execute(w, struct {
		Context      *renderer.Context
		Term         string
		Content      string
		Continuation bool
	}{
		Context:      ctx,
		Term:         string(term),
		Continuation: continuation,
		Content:      string(content),
	})
	return content == "", err
}
