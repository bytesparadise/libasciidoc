package sgml

import (
	"bytes"
	"path/filepath"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (r *sgmlRenderer) renderInternalCrossReference(ctx *renderer.Context, xref types.InternalCrossReference) ([]byte, error) {
	log.Debugf("rendering cross reference with ID: %s", xref.ID)
	result := &bytes.Buffer{}
	var label string
	if xref.Label != "" {
		label = xref.Label
	} else if target, found := ctx.ElementReferences[xref.ID]; found {
		if t, ok := target.([]interface{}); ok {
			renderedContent, err := r.renderElement(ctx, t)
			if err != nil {
				return nil, errors.Wrapf(err, "error while rendering internal cross reference")
			}
			label = string(renderedContent)
		} else {
			return nil, errors.Errorf("unable to process internal cross reference to element of type %T", target)
		}
	} else {
		label = "[" + xref.ID + "]"
	}
	err := r.internalCrossReference.Execute(result, struct {
		Href  string
		Label string
	}{
		Href:  xref.ID,
		Label: label,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render internal cross reference")
	}
	return result.Bytes(), nil
}

func (r *sgmlRenderer) renderExternalCrossReference(ctx *renderer.Context, xref types.ExternalCrossReference) ([]byte, error) {
	log.Debugf("rendering cross reference with ID: %s", xref.Location)
	result := &bytes.Buffer{}
	label, err := r.renderInlineElements(ctx, xref.Label)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render external cross reference")
	}
	err = r.externalCrossReference.Execute(result, struct {
		Href  string
		Label string
	}{
		Href:  getCrossReferenceLocation(xref),
		Label: string(label),
	})
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render external cross reference")
	}
	return result.Bytes(), nil
}

func getCrossReferenceLocation(xref types.ExternalCrossReference) string {
	loc := xref.Location.String()
	ext := filepath.Ext(xref.Location.String())
	log.Debugf("ext of '%s': '%s'", loc, ext)
	return loc[:len(loc)-len(ext)] + ".html" // TODO output extension
}
