package html5

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var crossReferenceTmpl template.Template

// initializes the templates
func init() {
	crossReferenceTmpl = newHTMLTemplate("cross reference", `<a href="#{{ .ID }}">{{ .Content }}</a>`)
}

func renderCrossReference(ctx *renderer.Context, xref types.CrossReference) ([]byte, error) {
	log.Debugf("rendering cross reference with ID: %s", xref.ID)
	result := bytes.NewBuffer(nil)
	renderedContentStr := fmt.Sprintf("[%s]", xref.ID)
	if target, found := ctx.Document.ElementReferences[xref.ID]; found {
		switch t := target.(type) {
		case types.SectionTitle:
			renderedContent, err := renderElement(ctx, t.Content)
			if err != nil {
				return nil, errors.Wrapf(err, "error while rendering sectionTitle content")
			}
			renderedContentStr = string(renderedContent)
		default:
			return nil, errors.Errorf("unable to process cross-reference to element of type %T", target)
		}
	}
	err := crossReferenceTmpl.Execute(result, struct {
		ID      string
		Content string
	}{
		ID:      xref.ID,
		Content: renderedContentStr,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render cross reference")
	}
	return result.Bytes(), nil
}
