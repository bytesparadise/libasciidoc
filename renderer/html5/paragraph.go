package html5

import (
	"bytes"
	"html/template"

	"github.com/bytesparadise/libasciidoc/renderer"
	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
)

var paragraphTmpl template.Template

// initializes the template
func init() {
	// TODO: use iterator and render func in the paragraph template
	paragraphTmpl = newHTMLTemplate("paragraph",
		`<div {{ if ne .ID "" }}id="{{.ID}}" {{ end }}class="paragraph">{{ if ne .Title "" }}
<div class="doctitle">{{.Title}}</div>{{ end }}
<p>{{.Lines}}</p>
</div>`)
}

func renderParagraph(ctx *renderer.Context, p types.Paragraph) ([]byte, error) {
	renderedLinesBuff := bytes.NewBuffer(nil)
	for i, line := range p.Lines {
		renderedLine, err := renderInlineContent(ctx, line)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render paragraph line")
		}
		renderedLinesBuff.Write(renderedLine)
		if i < len(p.Lines)-1 {
			renderedLinesBuff.WriteString("\n")
		}

	}
	// skip rendering if there's no content in the paragraph (eg: empty passthough)
	if renderedLinesBuff.Len() == 0 {
		return []byte{}, nil
	}
	result := bytes.NewBuffer(nil)
	err := paragraphTmpl.Execute(result, struct {
		ID    string
		Title string
		Lines template.HTML
	}{
		ID:    p.ID.Value,
		Title: p.Title.Value,
		Lines: template.HTML(renderedLinesBuff.String()), // here we must preserve the HTML tags
	})
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render paragraph")
	}
	// log.Debugf("rendered paragraph: %s", result.Bytes())
	return result.Bytes(), nil
}
