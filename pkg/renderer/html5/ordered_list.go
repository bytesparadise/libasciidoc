package html5

import (
	"bytes"
	"html"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

var orderedListTmpl texttemplate.Template

// initializes the templates
func init() {
	orderedListTmpl = newTextTemplate("ordered list",
		`{{ $ctx := .Context }}{{ with .Data }}{{ $items := .Items }}{{ $firstItem := index $items 0 }}<div{{ if .ID }} id="{{ .ID }}"{{ end }} class="olist {{ $firstItem.NumberingStyle }}{{ if .Role }} {{ .Role }}{{ end}}">
{{ if .Title }}<div class="title">{{ escape .Title }}</div>
{{ end }}<ol class="{{ $firstItem.NumberingStyle }}"{{ style $firstItem.NumberingStyle }}{{ if .Start }} start="{{ .Start }}"{{ end }}>
{{ range $itemIndex, $item := $items }}<li>
{{ renderElements $ctx $item.Elements | printf "%s" }}
</li>
{{ end }}</ol>
</div>{{ end }}`,
		texttemplate.FuncMap{
			"renderElements": renderListElements,
			"style":          numberingType,
			"escape":         html.EscapeString,
		})

}

func renderOrderedList(ctx *renderer.Context, l types.OrderedList) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	err := orderedListTmpl.Execute(result, ContextualPipeline{
		Context: ctx,
		Data: struct {
			ID    string
			Title string
			Role  string
			Start string
			Items []types.OrderedListItem
		}{
			generateID(ctx, l.Attributes),
			l.Attributes.GetAsString(types.AttrTitle),
			l.Attributes.GetAsString(types.AttrRole),
			l.Attributes.GetAsString(types.AttrStart),
			l.Items,
		},
	})
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render ordered list")
	}
	// log.Debugf("rendered ordered list of items: %s", result.Bytes())
	return result.Bytes(), nil
}

func numberingType(s types.NumberingStyle) string {
	switch s {
	case types.LowerAlpha:
		return ` type="a"`
	case types.UpperAlpha:
		return ` type="A"`
	case types.LowerRoman:
		return ` type="i"`
	case types.UpperRoman:
		return ` type="I"`
	default:
		return ""
	}
}
