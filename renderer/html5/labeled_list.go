package html5

import (
	"bytes"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/renderer"
	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var defaultLabeledListTmpl texttemplate.Template
var horizontalLabeledListTmpl texttemplate.Template

// initializes the templates
func init() {
	defaultLabeledListTmpl = newTextTemplate("labeled list with default layout",
		`{{ $ctx := .Context }}{{ with .Data }}<div{{ if index .Attributes "ID" }} id="{{ index .Attributes "ID" }}"{{ end }} class="dlist">
<dl>
{{ $items := .Items }}{{ range $itemIndex, $item := $items }}<dt class="hdlist1">{{ $item.Term }}</dt>{{ if $item.Elements }}
<dd>
{{ $elements := $item.Elements }}{{ range $elementIndex, $element := $elements }}{{ renderElement $ctx $element | printf "%s" }}{{ if includeNewline $ctx $elementIndex $elements }}{{ print "\n" }}{{ end }}{{ end }}
</dd>{{ end }}
{{ end }}</dl>
</div>{{ end }}`,
		texttemplate.FuncMap{
			"renderElement":  renderElement,
			"wrap":           wrap,
			"includeNewline": includeNewline,
		})

	horizontalLabeledListTmpl = newTextTemplate("labeled list with horizontal layout",
		`{{ $ctx := .Context }}{{ with .Data }}<div{{ if index .Attributes "ID" }} id="{{ index .Attributes "ID" }}"{{ end }} class="hdlist">
<table>
<tr>
<td class="hdlist1">{{ $items := .Items }}{{ range $itemIndex, $item := $items }}
{{ $item.Term }}
{{ if $item.Elements }}</td>
<td class="hdlist2">
{{ $elements := $item.Elements }}{{ range $elementIndex, $element := $elements }}{{ renderElement $ctx $element | printf "%s" }}{{ if includeNewline $ctx $elementIndex $elements }}{{ print "\n" }}{{ end }}{{ end }}
{{ if includeNewline $ctx $itemIndex $items }}</td>
</tr>
<tr>
<td class="hdlist1">{{ else }}</td>{{ end }}{{ else }}<br>{{ end }}{{ end }}
</tr>
</table>
</div>{{ end }}`,
		texttemplate.FuncMap{
			"renderElement":  renderElement,
			"wrap":           wrap,
			"includeNewline": includeNewline,
		})

}

func renderLabeledList(ctx *renderer.Context, l types.LabeledList) ([]byte, error) {
	var tmpl texttemplate.Template
	if layout, ok := l.Attributes["layout"]; ok {
		switch layout {
		case "horizontal":
			tmpl = horizontalLabeledListTmpl
		default:
			return nil, errors.Errorf("unsupported labeled list layout: %s", layout)
		}
	} else {
		tmpl = defaultLabeledListTmpl
	}

	// make sure nested elements are aware of that their rendering occurs within a list
	ctx.SetWithinList(true)
	defer func() {
		ctx.SetWithinList(false)
	}()

	result := bytes.NewBuffer(nil)
	// here we must preserve the HTML tags
	err := tmpl.Execute(result, ContextualPipeline{
		Context: ctx,
		Data:    l,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render labeled list")
	}
	log.Debugf("rendered labeled list: %s", result.Bytes())
	return result.Bytes(), nil
}
