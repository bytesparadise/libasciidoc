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
func ConvertToHTMLBody(ctx context.Context, r io.Reader, w io.Writer) (map[string]interface{}, error) {
	doc, err := parser.ParseReader("", r)
	if err != nil {
		return nil, errors.Wrapf(err, "error while parsing the document")
	}
	document := doc.(*types.Document)
	options := []renderer.Option{renderer.IncludeHeaderFooter(false)}
	metadata, err := htmlrenderer.Render(renderer.Wrap(ctx, *document, options...), w)
	if err != nil {
		return nil, errors.Wrapf(err, "error while rendering the document")
	}
	log.Debugf("Done processing document")
	return metadata, nil
}

// ConvertToHTML converts the content of the given reader `r` into a full HTML document, written in the given writer `w`.
// Returns an error if a problem occurred
func ConvertToHTML(ctx context.Context, r io.Reader, w io.Writer, options ...renderer.Option) (map[string]interface{}, error) {
	doc, err := parser.ParseReader("", r)
	if err != nil {
		return nil, errors.Wrapf(err, "error while parsing the document")
	}
	document := doc.(*types.Document)
	// force/override value
	options = append(options, renderer.IncludeHeaderFooter(true))
	metadata, err := htmlrenderer.Render(renderer.Wrap(ctx, *document, options...), w)
	if err != nil {
		return nil, errors.Wrapf(err, "error while rendering the document")
	}
	log.Debugf("Done processing document")
	return metadata, nil
}
