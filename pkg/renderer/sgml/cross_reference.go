package sgml

import (
	"path/filepath"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"

	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (r *sgmlRenderer) renderInternalCrossReference(ctx *context, xref *types.InternalCrossReference) (string, error) {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("rendering cross reference with ID: %s", spew.Sdump(xref.ID))
	}
	var label string
	xrefID, ok := xref.ID.(string)
	if !ok {
		return "", errors.Errorf("unable to process internal cross reference: invalid ID: '%v'", xref.ID)
	}
	if xrefLabel, ok := xref.Label.(string); ok {
		label = xrefLabel
	} else if target, found := ctx.elementReferences[xrefID]; found {
		switch t := target.(type) {
		case string:
			label = t
		case []interface{}:
			// render as usual except for links as plain text (since the cross reference is already displayed as a link)
			buff := &strings.Builder{}
			for _, e := range t {
				switch e := e.(type) {
				case *types.InlineLink:
					renderedElement, err := RenderPlainText(e)
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
	return r.execute(r.internalCrossReference, struct {
		Href  string
		Label string
	}{
		Href:  xrefID,
		Label: label,
	})
}

func (r *sgmlRenderer) renderExternalCrossReference(ctx *context, xref *types.ExternalCrossReference) (string, error) {
	// log.Debugf("rendering cross reference with ID: %s", xref.Location)
	var label string
	var err error
	switch l := xref.Attributes[types.AttrXRefLabel].(type) {
	case string:
		label = l
	case []interface{}:
		if label, err = r.renderInlineElements(ctx, l); err != nil {
			return "", err
		}
	default:
		label = defaultXrefLabel(xref)
	}
	return r.execute(r.externalCrossReference, struct {
		Href  string
		Label string
	}{
		Href:  getCrossReferenceLocation(xref),
		Label: label,
	})
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
