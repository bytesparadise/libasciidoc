package libasciidoc

import (
	"context"
	"io"

	"github.com/bytesparadise/libasciidoc/parser"
	"github.com/bytesparadise/libasciidoc/renderer"
	htmlrenderer "github.com/bytesparadise/libasciidoc/renderer/html5"
	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// ConvertToHTMLBody converts the content of the given reader `r` into an set of <DIV> elements for an HTML/BODY document.
// The conversion result is written in the given writer `w`, whereas the document metadata (title, etc.) (or an error if a problem occurred) is returned
// as the result of the function call.
func ConvertToHTMLBody(r io.Reader, w io.Writer) (*types.DocumentAttributes, error) {
	doc, err := parser.ParseReader("", r)
	if err != nil {
		return nil, errors.Wrapf(err, "error while parsing the document")
	}
	document := doc.(*types.Document)
	options := renderer.Options{}
	options[renderer.IncludeHeaderFooter] = false // force value
	err = htmlrenderer.Render(context.Background(), *document, w, options)
	if err != nil {
		return nil, errors.Wrapf(err, "error while rendering the document")
	}
	log.Debugf("Done processing document")
	return document.Attributes, nil
}

// ConvertToHTML converts the content of the given reader `r` into a full HTML document, written in the given writer `w`.
// Returns an error if a problem occurred
func ConvertToHTML(r io.Reader, w io.Writer, options renderer.Options) error {

	doc, err := parser.ParseReader("", r)
	if err != nil {
		return errors.Wrapf(err, "error while parsing the document")
	}
	document := doc.(*types.Document)
	options[renderer.IncludeHeaderFooter] = true // force value
	err = htmlrenderer.Render(context.Background(), *document, w, options)
	if err != nil {
		return errors.Wrapf(err, "error while rendering the document")
	}
	log.Debugf("Done processing document")
	return nil
}
