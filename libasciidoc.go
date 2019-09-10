// Package libasciidoc is an open source Go library that converts Asciidoc
// content into HTML.
package libasciidoc

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	htmlrenderer "github.com/bytesparadise/libasciidoc/pkg/renderer/html5"
	"github.com/bytesparadise/libasciidoc/pkg/types"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var (
	// BuildCommit lastest build commit (set by Makefile)
	BuildCommit = ""
	// BuildTag if the `BuildCommit` matches a tag
	BuildTag = ""
	// BuildTime set by build script (set by Makefile)
	BuildTime = ""
)

// ConvertFileToHTML converts the content of the given filename into an HTML document.
// The conversion result is written in the given writer `output`, whereas the document metadata (title, etc.) (or an error if a problem occurred) is returned
// as the result of the function call.
func ConvertFileToHTML(ctx context.Context, filename string, output io.Writer, options ...renderer.Option) (map[string]interface{}, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.Wrapf(err, "error opening %s", filename)
	}
	defer file.Close()
	return ConvertToHTML(ctx, file, output, options...)
}

// ConvertToHTML converts the content of the given reader `r` into a full HTML document, written in the given writer `output`.
// Returns an error if a problem occurred
func ConvertToHTML(ctx context.Context, r io.Reader, output io.Writer, options ...renderer.Option) (map[string]interface{}, error) {
	log.Debugf("parsing the asciidoc source...")
	doc, err := parser.ParseDocument("", r)
	if err != nil {
		return nil, errors.Wrapf(err, "error while parsing the document")
	}
	return convertToHTML(ctx, doc, output, options...)
}

func convertToHTML(ctx context.Context, doc types.Document, output io.Writer, options ...renderer.Option) (map[string]interface{}, error) {
	start := time.Now()
	rendererCtx := renderer.Wrap(ctx, doc, options...)
	// insert tables of contents, preamble and process file inclusions
	err := renderer.Prerender(rendererCtx)
	if err != nil {
		return nil, errors.Wrapf(err, "error while rendering the document")
	}
	metadata, err := htmlrenderer.Render(rendererCtx, output)
	if err != nil {
		return nil, errors.Wrapf(err, "error while rendering the document")
	}
	log.Debugf("Done processing document")
	duration := time.Since(start)
	log.Debugf("rendered the HTML output in %v", duration)
	return metadata, nil
}
