package sgml

import (
	"path/filepath"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"

	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (r *sgmlRenderer) renderInternalCrossReference(ctx *renderer.Context, xref *types.InternalCrossReference) (string, error) {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("rendering cross reference with ID: %s", spew.Sdump(xref.ID))
	}
	result := &strings.Builder{}
	var label string
	xrefID, ok := xref.ID.(string)
	if !ok {
		return "", errors.Errorf("unable to process internal cross reference: invalid ID: '%v'", xref.ID)
	}
	if xrefLabel, ok := xref.Label.(string); ok {
		label = xrefLabel
	} else if target, found := ctx.ElementReferences[xrefID]; found {
		switch t := target.(type) {
		case string:
			label = t
		case []interface{}:
			// render as usual except for links as plain text (since the cross reference is already displayed as a link)
			buff := &strings.Builder{}
			for _, e := range t {
				switch e := e.(type) {
				case *types.InlineLink:
					renderedElement, err := r.renderPlainText(ctx, e)
					if err != nil {
						return "", err
					}
					buff.WriteString(renderedElement)
				default:
					renderedElement, err := r.renderElement(ctx, e)
					if err != nil {
						return "", err
					}
					buff.WriteString(renderedElement)
				}
			}
			label = buff.String()
		default:
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

func (r *sgmlRenderer) renderExternalCrossReference(ctx *renderer.Context, xref *types.ExternalCrossReference) (string, error) {
	// log.Debugf("rendering cross reference with ID: %s", xref.Location)
	result := &strings.Builder{}
	var label string
	var err error
	switch l := xref.Attributes[types.AttrXRefLabel].(type) {
	case string:
		label = l
	case []interface{}:
		if label, err = r.renderInlineElements(ctx, l); err != nil {
			return "", errors.Wrap(err, "unable to render external cross reference")
		}
	default:
		label = defaultXrefLabel(xref)
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

func defaultXrefLabel(xref *types.ExternalCrossReference) string {
	loc := xref.Location.ToDisplayString()
	ext := filepath.Ext(loc)
	if ext == "" {
		return "[" + loc + "]" // internal references are within brackets
	}
	return loc[:len(loc)-len(ext)] + ".html"
}

func getCrossReferenceLocation(xref *types.ExternalCrossReference) string {
	loc := xref.Location.ToDisplayString()
	ext := filepath.Ext(loc)
	if ext == "" { // internal reference
		return "#" + loc
	}
	return loc[:len(loc)-len(ext)] + ".html" // TODO output extension
}
