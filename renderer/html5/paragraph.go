package html5

import (
	"bytes"
	"html/template"

	"github.com/bytesparadise/libasciidoc/renderer"
	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
)

var paragraphTmpl *template.Template

// initializes the template
func init() {
	paragraphTmpl = newHTMLTemplate("paragraph",
		`<div {{ if .ID }}id="{{.ID.Value}}" {{ end }}class="paragraph">{{ if .Title}}
<div class="doctitle">{{.Title.Value}}</div>{{ end }}
<p>{{.Lines}}</p>
</div>`)
}

func renderParagraph(ctx *renderer.Context, paragraph types.Paragraph) ([]byte, error) {
	renderedLinesBuff := bytes.NewBuffer(nil)
	for i, line := range paragraph.Lines {
		renderedLine, err := renderInlineContent(ctx, *line)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render paragraph line")
		}
		renderedLinesBuff.Write(renderedLine)
		if i < len(paragraph.Lines)-1 {
			renderedLinesBuff.WriteString("\n")
		}

	}
	// skip rendering if there's no content in the paragraph (eg: empty passthough)
	if renderedLinesBuff.Len() == 0 {
		return []byte{}, nil
	}
	result := bytes.NewBuffer(nil)
	err := paragraphTmpl.Execute(result, struct {
		ID    *types.ElementID
		Title *types.ElementTitle
		Lines template.HTML
	}{
		ID:    paragraph.ID,
		Title: paragraph.Title,
		Lines: template.HTML(renderedLinesBuff.String()), // here we must preserve the HTML tags
	})
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render paragraph")
	}
	// log.Debugf("rendered paragraph: %s", result.Bytes())
	return result.Bytes(), nil
}
