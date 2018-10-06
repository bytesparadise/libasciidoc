package html5

import (
	"bytes"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var unorderedListTmpl texttemplate.Template

// initializes the templates
func init() {
	unorderedListTmpl = newTextTemplate("unordered list",
		`{{ $ctx := .Context }}{{ with .Data }}<div{{ if .ID }} id="{{ .ID }}"{{ end }} class="ulist{{ if .Role }} {{ .Role }}{{ end}}">
{{ if .Title }}<div class="title">{{ .Title }}</div>
{{ end }}<ul>
{{ $items := .Items }}{{ range $itemIndex, $item := $items }}<li>
{{ $elements := $item.Elements }}{{ renderElements $ctx $elements | printf "%s" }}
</li>
{{ end }}</ul>
</div>{{ end }}`,
		texttemplate.FuncMap{
			"renderElements": renderElements,
		})
}

func renderUnorderedList(ctx *renderer.Context, l types.UnorderedList) ([]byte, error) {
	// make sure nested elements are aware of that their rendering occurs within a list
	ctx.SetWithinList(true)
	defer func() {
		ctx.SetWithinList(false)
	}()

	result := bytes.NewBuffer(nil)
	// here we must preserve the HTML tags
	err := unorderedListTmpl.Execute(result, ContextualPipeline{
		Context: ctx,
		Data: struct {
			ID    string
			Title string
			Role  string
			Items []types.UnorderedListItem
		}{
			ID:    l.Attributes.GetAsString(types.AttrID),
			Title: l.Attributes.GetAsString(types.AttrTitle),
			Role:  l.Attributes.GetAsString(types.AttrRole),
			Items: l.Items,
		},
	})
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render unordered list")
	}
	log.Debugf("rendered unordered list of items: %s", result.Bytes())
	return result.Bytes(), nil
}
