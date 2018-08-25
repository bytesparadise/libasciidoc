package html5

import (
	"bytes"
	"fmt"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var crossReferenceTmpl texttemplate.Template

// initializes the templates
func init() {
	crossReferenceTmpl = newTextTemplate("cross reference", `<a href="#{{ .ID }}">{{ .Label }}</a>`)
}

func renderCrossReference(ctx *renderer.Context, xref types.CrossReference) ([]byte, error) {
	log.Debugf("rendering cross reference with ID: %s", xref.ID)
	result := bytes.NewBuffer(nil)
	var label string
	if xref.Label != "" {
		label = xref.Label
	} else if target, found := ctx.Document.ElementReferences[xref.ID]; found {
		switch t := target.(type) {
		case types.SectionTitle:
			renderedContent, err := renderElement(ctx, t.Content)
			if err != nil {
				return nil, errors.Wrapf(err, "error while rendering sectionTitle content")
			}
			label = string(renderedContent)
		default:
			return nil, errors.Errorf("unable to process cross-reference to element of type %T", target)
		}
	} else {
		label = fmt.Sprintf("[%s]", xref.ID)
	}
	err := crossReferenceTmpl.Execute(result, struct {
		ID    string
		Label string
	}{
		ID:    xref.ID,
		Label: label,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render cross reference")
	}
	return result.Bytes(), nil
}
