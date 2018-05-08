package html5

import (
	"bytes"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/renderer"
	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var unorderedListTmpl texttemplate.Template

// initializes the templates
func init() {
	unorderedListTmpl = newTextTemplate("unordered list",
		`{{ $ctx := .Context }}{{ with .Data }}<div{{ if index .Attributes "ID" }} id="{{ index .Attributes "ID" }}"{{ end }} class="ulist">
<ul>
{{ $items := .Items }}{{ range $itemIndex, $item := $items }}<li>
{{ $elements := $item.Elements }}{{ range $elementIndex, $element := $elements }}{{ renderElement $ctx $element | printf "%s" }}{{ if includeNewline $ctx $elementIndex $elements }}{{ print "\n" }}{{ end }}{{ end }}
</li>
{{ end }}</ul>
</div>{{ end }}`,
		texttemplate.FuncMap{
			"renderElement":  renderElement,
			"wrap":           wrap,
			"includeNewline": includeNewline,
		})

}

func renderUnorderedList(ctx *renderer.Context, l types.UnorderedList) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	// here we must preserve the HTML tags
	err := unorderedListTmpl.Execute(result, ContextualPipeline{
		Context: ctx,
		Data:    l,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render unordered list")
	}
	log.Debugf("rendered unordered list of items: %s", result.Bytes())
	return result.Bytes(), nil
}
