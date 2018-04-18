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

// ConvertFileToHTML converts the content of the given filename into an HTML document.
// The conversion result is written in the given writer `output`, whereas the document metadata (title, etc.) (or an error if a problem occurred) is returned
// as the result of the function call.
func ConvertFileToHTML(ctx context.Context, filename string, output io.Writer, options ...renderer.Option) (map[string]interface{}, error) {
	doc, err := parser.ParseFile(filename)
	if err != nil {
		return nil, errors.Wrapf(err, "error while parsing the document")
	}
	return convertToHTML(ctx, doc, output, options...)
}

// ConvertToHTML converts the content of the given reader `r` into a full HTML document, written in the given writer `output`.
// Returns an error if a problem occurred
func ConvertToHTML(ctx context.Context, r io.Reader, output io.Writer, options ...renderer.Option) (map[string]interface{}, error) {
	doc, err := parser.ParseReader("", r)
	if err != nil {
		return nil, errors.Wrapf(err, "error while parsing the document")
	}
	return convertToHTML(ctx, doc, output, options...)
}

func convertToHTML(ctx context.Context, doc interface{}, output io.Writer, options ...renderer.Option) (map[string]interface{}, error) {
	metadata, err := htmlrenderer.Render(renderer.Wrap(ctx, doc.(types.Document), options...), output)
	if err != nil {
		return nil, errors.Wrapf(err, "error while rendering the document")
	}
	log.Debugf("Done processing document")
	return metadata, nil
}
