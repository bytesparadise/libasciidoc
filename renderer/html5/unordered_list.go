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
{{ $elements := $item.Elements }}{{ range $elementIndex, $element := $elements }}{{ renderElement $ctx $element | printf "%s" }}{{ if notLastItem $elementIndex $elements }}{{ print "\n" }}{{ end }}{{ end }}
</li>
{{ end }}</ul>
</div>{{ end }}`,
		texttemplate.FuncMap{
			"renderElement": renderElement,
			"wrap":          wrap,
			"notLastItem":   notLastItem,
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
		return nil, errors.Wrapf(err, "unable to render list of items")
	}
	log.Debugf("rendered unordered list of items: %s", result.Bytes())
	return result.Bytes(), nil
}

// func renderUnorderedListItem(ctx *renderer.Context, i types.UnorderedListItem) ([]byte, error) {
// 	renderedItemContent, err := renderUnorderedListItemContent(ctx, i.Elements)
// 	if err != nil {
// 		return nil, errors.Wrapf(err, "unable to render list item")
// 	}
// 	result := bytes.NewBuffer(nil)
// 	var renderedChildrenOutput *template.HTML
// 	if i.Children != nil {
// 		childrenOutput, err := renderUnorderedList(ctx, i.Children)
// 		if err != nil {
// 			return nil, errors.Wrapf(err, "unable to render list item")
// 		}
// 		htmlChildrenOutput := template.HTML(string(childrenOutput))
// 		renderedChildrenOutput = &htmlChildrenOutput
// 	}
// 	err = unorderedListItemTmpl.Execute(result, struct {
// 		Content  template.HTML
// 		Children *template.HTML
// 	}{
// 		Content:  template.HTML(string(renderedItemContent)),
// 		Children: renderedChildrenOutput,
// 	})
// 	if err != nil {
// 		return nil, errors.Wrapf(err, "unable to render list item")
// 	}
// 	log.Debugf("rendered item: %s", result.Bytes())
// 	return result.Bytes(), nil
// }

// func renderUnorderedListItemContent(ctx *renderer.Context, elements []types.DocElement) ([]byte, error) {
// 	renderedElementsBuff := bytes.NewBuffer(nil)
// 	for _, element := range elements {
// 		renderedElement, err := renderElement(ctx, element)
// 		if err != nil {
// 			return nil, errors.Wrapf(err, "failed to render list item content")
// 		}
// 		renderedElementsBuff.Write(renderedElement)
// 	}
// 	result := bytes.NewBuffer(nil)
// 	err := unorderedListItemContentTmpl.Execute(result, renderedElementsBuff.String())
// 	if err != nil {
// 		return nil, errors.Wrapf(err, "unable to render list item")
// 	}
// 	log.Debugf("rendered item content: %s", result.Bytes())
// 	return result.Bytes(), nil
// }
