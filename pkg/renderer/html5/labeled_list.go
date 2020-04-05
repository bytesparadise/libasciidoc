package html5

import (
	"bytes"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

var defaultLabeledListTmpl texttemplate.Template
var horizontalLabeledListTmpl texttemplate.Template
var qandaLabeledListTmpl texttemplate.Template

// initializes the templates
func init() {
	defaultLabeledListTmpl = newTextTemplate("labeled list with default layout",
		`{{ $ctx := .Context }}{{ with .Data }}<div{{ if .ID }} id="{{ .ID }}"{{ end }} class="dlist{{ if .Role }} {{ .Role }}{{ end }}">
{{ if .Title }}<div class="title">{{ escape .Title }}</div>
{{ end }}<dl>
{{ $items := .Items }}{{ range $itemIndex, $item := $items }}<dt class="hdlist1">{{ renderInlineElements $ctx $item.Term | printf "%s" }}</dt>{{ if $item.Elements }}
<dd>
{{ renderElements $ctx $item.Elements | printf "%s" }}
</dd>{{ end }}
{{ end }}</dl>
</div>{{ end }}`,
		texttemplate.FuncMap{
			"renderInlineElements": renderInlineElements,
			"renderElements":       renderListElements,
			"escape":               EscapeString,
		})

	horizontalLabeledListTmpl = newTextTemplate("labeled list with horizontal layout",
		`{{ $ctx := .Context }}{{ with .Data }}<div{{ if .ID }} id="{{ .ID }}"{{ end }} class="hdlist{{ if .Role }} {{ .Role }}{{ end }}">
{{ if .Title }}<div class="title">{{ escape .Title }}</div>
{{ end }}<table>
<tr>
<td class="hdlist1">{{ $items := .Items }}{{ range $itemIndex, $item := $items }}
{{ renderInlineElements $ctx $item.Term | printf "%s" }}
{{ if $item.Elements }}</td>
<td class="hdlist2">
{{ renderElements $ctx $item.Elements | printf "%s" }}
{{ if includeNewline $ctx $itemIndex $items }}</td>
</tr>
<tr>
<td class="hdlist1">{{ else }}</td>{{ end }}{{ else }}<br>{{ end }}{{ end }}
</tr>
</table>
</div>{{ end }}`,
		texttemplate.FuncMap{
			"renderInlineElements": renderInlineElements,
			"renderElements":       renderListElements,
			"includeNewline":       includeNewline,
			"escape":               EscapeString,
		})

	qandaLabeledListTmpl = newTextTemplate("qanda labeled list",
		`{{ $ctx := .Context }}{{ with .Data }}<div{{ if .ID }} id="{{ .ID }}"{{ end }} class="qlist qanda">
{{ if .Title }}<div class="title">{{ escape .Title }}</div>
{{ end }}<ol>
{{ $items := .Items }}{{ range $itemIndex, $item := $items }}<li>
<p><em>{{ renderInlineElements $ctx $item.Term | printf "%s" }}</em></p>
{{ if $item.Elements }}{{ renderElements $ctx $item.Elements | printf "%s" }}{{ end }}
</li>
{{ end }}</ol>
</div>{{ end }}`,
		texttemplate.FuncMap{
			"renderInlineElements": renderInlineElements,
			"renderElements":       renderListElements,
			"escape":               EscapeString,
		})

}

func renderLabeledList(ctx renderer.Context, l types.LabeledList) ([]byte, error) {
	var tmpl texttemplate.Template
	tmpl, err := getLabeledListTmpl(l)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render labeled list")
	}

	result := bytes.NewBuffer(nil)
	// here we must preserve the HTML tags
	err = tmpl.Execute(result, ContextualPipeline{
		Context: ctx,
		Data: struct {
			ID    string
			Title string
			Role  string
			Items []types.LabeledListItem
		}{
			ID:    renderElementID(l.Attributes),
			Title: renderElementTitle(l.Attributes),
			Role:  l.Attributes.GetAsString(types.AttrRole),
			Items: l.Items,
		},
	})
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render labeled list")
	}
	// log.Debugf("rendered labeled list: %s", result.Bytes())
	return result.Bytes(), nil
}

func getLabeledListTmpl(l types.LabeledList) (texttemplate.Template, error) {
	if layout, ok := l.Attributes["layout"]; ok {
		switch layout {
		case "horizontal":
			return horizontalLabeledListTmpl, nil
		default:
			return texttemplate.Template{}, errors.Errorf("unsupported labeled list layout: %s", layout)
		}
	}
	if l.Attributes.Has(types.AttrQandA) {
		return qandaLabeledListTmpl, nil
	}
	return defaultLabeledListTmpl, nil
}
