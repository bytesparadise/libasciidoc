package html5

import (
	"bytes"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var orderedListTmpl texttemplate.Template

// initializes the templates
func init() {
	orderedListTmpl = newTextTemplate("ordered list",
		`{{ $ctx := .Context }}{{ with .Data }}{{ $items := .Items }}{{ $firstItem := index $items 0 }}<div{{ if hasID .Attributes }} id="{{ getID .Attributes }}"{{ end }} class="olist {{ $firstItem.NumberingStyle }}">
<ol class="{{ $firstItem.NumberingStyle }}"{{ style $firstItem.NumberingStyle }}>
{{ range $itemIndex, $item := $items }}<li>
{{ $elements := $item.Elements }}{{ range $elementIndex, $element := $elements }}{{ renderElement $ctx $element | printf "%s" }}{{ if includeNewline $ctx $elementIndex $elements }}{{ print "\n" }}{{ end }}{{ end }}
</li>
{{ end }}</ol>
</div>{{ end }}`,
		texttemplate.FuncMap{
			"renderElement":  renderElement,
			"wrap":           wrap,
			"includeNewline": includeNewline,
			"style":          numberingType,
			"hasID":          hasID,
			"getID":          getID,
		})

}

func renderOrderedList(ctx *renderer.Context, l types.OrderedList) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	// make sure nested elements are aware of that their rendering occurs within a list
	ctx.SetWithinList(true)
	defer func() {
		ctx.SetWithinList(false)
	}()

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
