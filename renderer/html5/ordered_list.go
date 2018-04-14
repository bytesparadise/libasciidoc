package html5

import (
	"bytes"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/renderer"
	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var orderedListTmpl texttemplate.Template

// initializes the templates
func init() {
	orderedListTmpl = newTextTemplate("ordered list",
		`{{ $ctx := .Context }}{{ with .Data }}{{ $items := .Items }}{{ $firstItem := index $items 0 }}<div{{ if index .Attributes "ID" }} id="{{ index .Attributes "ID" }}"{{ end }} class="olist {{ $firstItem.NumberingStyle }}">
<ol class="{{ $firstItem.NumberingStyle }}"{{ style $firstItem.NumberingStyle }}>
{{ range $itemIndex, $item := $items }}<li>
{{ $elements := $item.Elements }}{{ range $elementIndex, $element := $elements }}{{ renderElement $ctx $element | printf "%s" }}{{ if notLastItem $elementIndex $elements }}{{ print "\n" }}{{ end }}{{ end }}
</li>
{{ end }}</ol>
</div>{{ end }}`,
		texttemplate.FuncMap{
			"renderElement": renderElement,
			"wrap":          wrap,
			"notLastItem":   notLastItem,
			"style":         numberingType,
		})

}

func renderOrderedList(ctx *renderer.Context, l types.OrderedList) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	// here we must preserve the HTML tags
	err := orderedListTmpl.Execute(result, ContextualPipeline{
		Context: ctx,
		Data:    l,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render ordered list")
	}
	log.Debugf("rendered ordered list of items: %s", result.Bytes())
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
