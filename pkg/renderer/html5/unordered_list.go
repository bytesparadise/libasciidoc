package html5

import (
	"bytes"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

var unorderedListTmpl texttemplate.Template

// initializes the templates
func init() {
	unorderedListTmpl = newTextTemplate("unordered list",
		`{{ $ctx := .Context }}{{ with .Data }}<div{{ if .ID }} id="{{ .ID }}"{{ end }} class="ulist{{ if .Checklist }} checklist{{ end }}{{ if .Role }} {{ .Role }}{{ end}}">
{{ if .Title }}<div class="title">{{ escape .Title }}</div>
{{ end }}<ul{{ if .Checklist }} class="checklist"{{ end }}>
{{ $items := .Items }}{{ range $itemIndex, $item := $items }}<li>
{{ $elements := $item.Elements }}{{ renderElements $ctx $elements | printf "%s" }}
</li>
{{ end }}</ul>
</div>{{ end }}`,
		texttemplate.FuncMap{
			"renderElements": renderListElements,
			"escape":         EscapeString,
		})
}

func renderUnorderedList(ctx *renderer.Context, l types.UnorderedList) ([]byte, error) {
	// make sure nested elements are aware of that their rendering occurs within a list
	checkList := false
	if len(l.Items) > 0 {
		if l.Items[0].CheckStyle != types.NoCheck {
			checkList = true
		}
	}
	result := bytes.NewBuffer(nil)
	// here we must preserve the HTML tags
	err := unorderedListTmpl.Execute(result, ContextualPipeline{
		Context: ctx,
		Data: struct {
			ID        string
			Title     string
			Role      string
			Checklist bool
			Items     []types.UnorderedListItem
		}{
			ID:        renderElementID(l.Attributes),
			Title:     renderTitle(l.Attributes),
			Role:      l.Attributes.GetAsString(types.AttrRole),
			Checklist: checkList,
			Items:     l.Items,
		},
	})
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render unordered list")
	}
	// log.Debugf("rendered unordered list of items: %s", result.Bytes())
	return result.Bytes(), nil
}
