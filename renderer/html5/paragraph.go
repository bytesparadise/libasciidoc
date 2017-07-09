package html5

import (
	"html/template"

	"bytes"

	"context"

	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var paragraphTmpl *template.Template

// initializes the template
func init() {
	paragraphTmpl = newTemplate("paragraph",
		`<div class="paragraph">
<p>{{.}}</p>
</div>`)
}

func renderParagraph(ctx context.Context, paragraph types.Paragraph) ([]byte, error) {
	renderedElementsBuff := bytes.NewBuffer(make([]byte, 0))
	for _, line := range paragraph.Lines {
		renderedElement, err := renderInlineContent(ctx, *line)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render paragraph element")
		}
		renderedElementsBuff.Write(renderedElement)
	}
	result := bytes.NewBuffer(make([]byte, 0))
	// here we must preserve the HTML tags
	err := paragraphTmpl.Execute(result, template.HTML(renderedElementsBuff.String()))
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render inline content")
	}
	log.Debugf("rendered paragraph: %s", result.Bytes())
	return result.Bytes(), nil
}
