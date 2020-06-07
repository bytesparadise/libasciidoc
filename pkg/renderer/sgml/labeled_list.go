package sgml

import (
	"bytes"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderLabeledList(ctx *renderer.Context, l types.LabeledList) ([]byte, error) {
	tmpl, err := r.getLabeledListTmpl(l)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render labeled list")
	}

	result := &bytes.Buffer{}
	// here we must preserve the HTML tags
	err = tmpl.Execute(result, ContextualPipeline{
		Context: ctx,
		Data: struct {
			ID    string
			Title string
			Role  string
			Items []types.LabeledListItem
		}{
			ID:    r.renderElementID(l.Attributes),
			Title: r.renderElementTitle(l.Attributes),
			Role:  l.Attributes.GetAsStringWithDefault(types.AttrRole, ""),
			Items: l.Items,
		},
	})
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render labeled list")
	}
	// log.Debugf("rendered labeled list: %s", result.Bytes())
	return result.Bytes(), nil
}

func (r *sgmlRenderer) getLabeledListTmpl(l types.LabeledList) (*textTemplate, error) {
	if layout, ok := l.Attributes["layout"]; ok {
		switch layout {
		case "horizontal":
			return r.labeledListHorizontal, nil
		default:
			return nil, errors.Errorf("unsupported labeled list layout: %s", layout)
		}
	}
	if l.Attributes.Has(types.AttrQandA) {
		return r.qAndAList, nil
	}
	return r.labeledList, nil
}
