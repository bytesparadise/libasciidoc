package html5

import (
	"bytes"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

var calloutListTmpl texttemplate.Template

// initializes the templates
func init() {

	calloutListTmpl = newTextTemplate("ordered list",
		`{{ $ctx := .Context }}{{ with .Data }}{{ $items := .Items }}<div{{ if .ID }} id="{{ .ID }}"{{ end }} class="colist arabic{{ if .Role }} {{ .Role }}{{ end}}">
{{ if .Title }}<div class="title">{{ escape .Title }}</div>
{{ end }}<ol>
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

func renderCalloutList(ctx renderer.Context, l types.CalloutList) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	err := calloutListTmpl.Execute(result, ContextualPipeline{
		Context: ctx,
		Data: struct {
			ID    string
			Title string
			Role  string
			Items []types.CalloutListItem
		}{
			ID:    renderElementID(l.Attributes),
			Title: l.Attributes.GetAsStringWithDefault(types.AttrTitle, ""),
			Role:  l.Attributes.GetAsStringWithDefault(types.AttrRole, ""),
			Items: l.Items,
		},
	})
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render callout list")
	}
	return result.Bytes(), nil
}
