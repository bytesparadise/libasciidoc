package sgml

import (
	"path/filepath"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (r *sgmlRenderer) renderInternalCrossReference(ctx *renderer.Context, xref types.InternalCrossReference) (string, error) {
	log.Debugf("rendering cross reference with ID: %s", xref.ID)
	result := &strings.Builder{}
	var label string
	xrefID, ok := xref.ID.(string)
	if !ok {
		return "", errors.Errorf("unable to process internal cross reference: invalid ID: '%v'", xref.ID)
	}
	if xrefLabel, ok := xref.Label.(string); ok {
		label = xrefLabel
	} else if target, found := ctx.ElementReferences[xrefID]; found {
		if t, ok := target.([]interface{}); ok {
			renderedContent, err := r.renderElement(ctx, t)
			if err != nil {
				return "", errors.Wrap(err, "error while rendering internal cross reference")
			}
			label = renderedContent
		} else {
			return "", errors.Errorf("unable to process internal cross reference to element of type %T", target)
		}
	} else {
		label = "[" + xrefID + "]"
	}
	err := r.internalCrossReference.Execute(result, struct {
		Href  string
		Label string
	}{
		Href:  xrefID,
		Label: label,
	})
	if err != nil {
		return "", errors.Wrapf(err, "unable to render internal cross reference")
	}
	return result.String(), nil
}

func (r *sgmlRenderer) renderExternalCrossReference(ctx *renderer.Context, xref types.ExternalCrossReference) (string, error) {
	log.Debugf("rendering cross reference with ID: %s", xref.Location)
	result := &strings.Builder{}
	var label string
	var err error
	switch l := xref.Label.(type) {
	case string:
		label = l
	case []interface{}:
		if label, err = r.renderInlineElements(ctx, l); err != nil {
			return "", errors.Wrap(err, "unable to render external cross reference")
		}
	default:
		return "", errors.Errorf("unable to render external cross reference label of type '%T'", xref.Label)
	}
	err = r.externalCrossReference.Execute(result, struct {
		Href  string
		Label string
	}{
		Href:  getCrossReferenceLocation(xref),
		Label: label,
	})
	if err != nil {
		return "", errors.Wrap(err, "unable to render external cross reference")
	}
	return result.String(), nil
}

func getCrossReferenceLocation(xref types.ExternalCrossReference) string {
	loc := xref.Location.Stringify()
	ext := filepath.Ext(xref.Location.Stringify())
	log.Debugf("ext of '%s': '%s'", loc, ext)
	return loc[:len(loc)-len(ext)] + ".html" // TODO output extension
}
