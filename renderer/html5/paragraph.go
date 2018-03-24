package html5

import (
	"bytes"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/renderer"
	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
)

var paragraphTmpl texttemplate.Template

// initializes the template
func init() {
	paragraphTmpl = newTextTemplate("paragraph",
		`{{ $ctx := .Context }}{{ with .Data }}{{ $renderedElements := renderElements $ctx .Lines | printf "%s"  }}{{ if ne $renderedElements "" }}<div {{ if ne .ID "" }}id="{{ .ID }}" {{ end }}class="paragraph">{{ if ne .Title "" }}
<div class="doctitle">{{ .Title }}</div>{{ end }}
<p>{{ $renderedElements }}</p>
</div>{{ end }}{{ end }}`,
		texttemplate.FuncMap{
			"renderElements": renderInlineContents,
			"notLastItem":    notLastItem,
		})
}

func renderParagraph(ctx *renderer.Context, p types.Paragraph) ([]byte, error) {
	if len(p.Lines) == 0 {
		return make([]byte, 0), nil
	}
	result := bytes.NewBuffer(nil)
	err := paragraphTmpl.Execute(result, ContextualPipeline{
		Context: ctx,
		Data: struct {
			ID    string
			Title string
			Lines []types.InlineContent
		}{
			ID:    p.ID.Value,
			Title: p.Title.Value,
			Lines: p.Lines,
		},
	})
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render paragraph")
	}
	return result.Bytes(), nil
}
