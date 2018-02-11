package html5

import (
	"bytes"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/renderer"
	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// var unorderedListTmpl *template.Template
// var unorderedNestedListTmpl *template.Template
// var unorderedListItemTmpl *template.Template
// var unorderedListItemContentTmpl *template.Template

const unorderedListTmpl = `{{ $ctx := .Context }}{{ with .Data }}<div{{ if index .Attributes "ID" }} id="{{ index .Attributes "ID" }}"{{ end }} class="ulist">
<ul>
{{ range .Items }}{{ template "items" wrap $ctx . }}{{ end }}</ul>
</div>{{ end }}`

const unorderedListItemTmpl = `{{ define "items" }}{{ $ctx := .Context }}{{ with .Data }}<li>
{{ $elements := .Elements }}{{ range $index, $element := $elements }}{{ renderElement $ctx $element | printf "%s" }}{{ if notLastItem $index $elements }}{{ print "\n" }}{{ end }}{{ end }}
</li>
{{ end }}{{ end }}`

// initializes the templates
func init() {
	// 	unorderedListTmpl = newHTMLTemplate("unordered list", `<div{{ if .ID }} id="{{.ID.Value}}"{{ end }} class="ulist">
	// <ul>
	// {{.Items}}
	// </ul>
	// </div>`)
	// 	unorderedListItemTmpl = newHTMLTemplate("unordered list item", `<li>
	// {{.Content}}{{ if .Children }}
	// {{.Children}}
	// </li>{{ else }}
	// </li>{{ end }}`)
	// 	unorderedListItemContentTmpl = newHTMLTemplate("unordered list item content", `<p>{{.}}</p>`)

}

func renderUnorderedList(ctx *renderer.Context, l *types.UnorderedList) ([]byte, error) {
	// TODO: move this to init
	t := texttemplate.New("unordered list")
	t.Funcs(texttemplate.FuncMap{
		"renderElement": renderElement,
		"wrap":          wrap,
		"notLastItem":   notLastItem,
	})
	var err error
	t, err = t.Parse(unorderedListItemTmpl)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to parse unordered list item template")
	}
	t, err = t.Parse(unorderedListTmpl)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to parse unordered list template")
	}

	// renderedElementsBuff := bytes.NewBuffer(nil)
	// for i, item := range l.Items {
	// 	renderedListItem, err := renderUnorderedListItem(ctx, *item)
	// 	if err != nil {
	// 		return nil, errors.Wrapf(err, "unable to render list of items")
	// 	}
	// 	renderedElementsBuff.Write(renderedListItem)
	// 	if i < len(l.Items)-1 {
	// 		renderedElementsBuff.WriteString("\n")
	// 	}
	// }

	// result := bytes.NewBuffer(nil)
	// // here we must preserve the HTML tags
	// err := unorderedListTmpl.Execute(result, struct {
	// 	ID    *types.ElementID
	// 	Items template.HTML
	// }{
	// 	ID:    l.ID,
	// 	Items: template.HTML(renderedElementsBuff.String()),
	// })
	// if err != nil {
	// 	return nil, errors.Wrapf(err, "unable to render l of items")
	// }
	result := bytes.NewBuffer(nil)
	// here we must preserve the HTML tags
	err = t.Execute(result, ContextualPipeline{
		Context: ctx,
		Data:    l,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render list of items")
	}
	log.Debugf("rendered list of items: %s", result.Bytes())
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
