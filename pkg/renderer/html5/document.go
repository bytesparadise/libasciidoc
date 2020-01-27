package html5

import (
	"bytes"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// renderDocumentElements renders all document elements, including the footnotes,
// but not the HEAD and BODY containers
func renderDocumentElements(ctx renderer.Context) ([]byte, error) {
	elements := []interface{}{}
	for i, e := range ctx.Document.Elements {
		switch e := e.(type) {
		case types.Preamble:
			if !e.HasContent() {
				// retain the preamble
				elements = append(elements, e)
				continue
			}
			// retain everything "as-is"
			elements = ctx.Document.Elements
		case types.Section:
			if e.Level == 0 {
				// retain the section's elements...
				elements = append(elements, e.Elements)
				// ... and add the other elements
				elements = append(elements, ctx.Document.Elements[i+1:]...)
				continue
			}
			// retain everything "as-is"
			elements = ctx.Document.Elements
		default:
			// retain everything "as-is"
			elements = ctx.Document.Elements
		}
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debug("pre-rendered elements:")
		spew.Dump(elements)
	}
	// log.Debugf("rendered document with %d element(s)...", len(elements))
	buff := bytes.NewBuffer(nil)
	renderedElements, err := renderElements(ctx, elements)
	if err != nil {
		return []byte{}, errors.Wrapf(err, "failed to render document elements")
	}
	buff.Write(renderedElements)
	renderedFootnotes, err := renderFootnotes(ctx, ctx.Document.Footnotes)
	if err != nil {
		return []byte{}, errors.Wrapf(err, "failed to render document elements")
	}
	buff.Write(renderedFootnotes)

	return buff.Bytes(), nil
}

func renderDocumentTitle(ctx renderer.Context) ([]byte, error) {
	if documentTitle, hasTitle := ctx.Document.Title(); hasTitle {
		title, err := renderPlainText(ctx, documentTitle)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render document title")
		}
		return title, nil
	}
	return nil, nil
}

func renderDocumentHeader(ctx renderer.Context) ([]byte, error) {
	if documentTitle, hasTitle := ctx.Document.Title(); hasTitle {
		title, err := renderInlineElements(ctx, documentTitle)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to render document header")
		}
		return title, nil
	}
	return nil, nil
}
