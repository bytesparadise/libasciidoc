package html5

import (
	"bytes"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

var orderedListTmpl texttemplate.Template

// initializes the templates
func init() {
	orderedListTmpl = newTextTemplate("ordered list",
		`{{ $ctx := .Context }}{{ with .Data }}{{ $items := .Items }}<div{{ if .ID }} id="{{ .ID }}"{{ end }} class="olist {{ .NumberingStyle }}{{ if .Role }} {{ .Role }}{{ end}}">
{{ if .Title }}<div class="title">{{ escape .Title }}</div>
{{ end }}<ol class="{{ .NumberingStyle }}"{{ style .NumberingStyle }}{{ if .Start }} start="{{ .Start }}"{{ end }}>
{{ range $itemIndex, $item := $items }}<li>
{{ renderElements $ctx $item.Elements | printf "%s" }}
</li>
{{ end }}</ol>
</div>{{ end }}`,
		texttemplate.FuncMap{
			"renderElements": renderListElements,
			"style":          numberingType,
			"escape":         EscapeString,
		})

}

func renderOrderedList(ctx *renderer.Context, l types.OrderedList) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	err := orderedListTmpl.Execute(result, ContextualPipeline{
		Context: ctx,
		Data: struct {
			ID             string
			Title          string
			Role           string
			NumberingStyle string
			Start          string
			Items          []types.OrderedListItem
		}{
			generateID(ctx, l.Attributes),
			l.Attributes.GetAsString(types.AttrTitle),
			l.Attributes.GetAsString(types.AttrRole),
			getNumberingStyle(l),
			l.Attributes.GetAsString(types.AttrStart),
			l.Items,
		},
	})
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render ordered list")
	}
	return result.Bytes(), nil
}

func getNumberingStyle(l types.OrderedList) string {
	if s := l.Attributes.GetAsString(types.AttrNumberingStyle); s != "" {
		return s
	}
	return string(l.Items[0].NumberingStyle)
}

func numberingType(style string) string {
	switch style {
	case string(types.LowerAlpha):
		return ` type="a"`
	case string(types.UpperAlpha):
		return ` type="A"`
	case string(types.LowerRoman):
		return ` type="i"`
	case string(types.UpperRoman):
		return ` type="I"`
	default:
		return ""
	}
}
