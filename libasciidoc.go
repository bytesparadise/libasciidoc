package libasciidoc

import (
	"context"
	"encoding/json"
	"io"
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
	log.Infof("parsing the asciidoc source...")
	stats := parser.Stats{}
	start := time.Now()
	doc, err := parser.ParseFile(filename, parser.Memoize(true), parser.Statistics(&stats, "no match"))
	if err != nil {
		return nil, errors.Wrapf(err, "error while parsing the document")
	}
	duration := time.Since(start)
	log.Infof("parsed the asciidoc source in %v ", duration)
	b, err := json.MarshalIndent(stats.ChoiceAltCnt, "", "  ")
	if err != nil {
		log.Warnf("failed to produce stats ", err)
	}
	log.Infof("parsing stats: %s", string(b))
	return convertToHTML(ctx, doc, output, options...)
}

// ConvertToHTML converts the content of the given reader `r` into a full HTML document, written in the given writer `output`.
// Returns an error if a problem occurred
func ConvertToHTML(ctx context.Context, r io.Reader, output io.Writer, options ...renderer.Option) (map[string]interface{}, error) {
	start := time.Now()
	stats := parser.Stats{}
	doc, err := parser.ParseReader("", r, parser.Memoize(true), parser.Statistics(&stats, "no match"))
	if err != nil {
		return nil, errors.Wrapf(err, "error while parsing the document")
	}
	duration := time.Since(start)
	log.Infof("parsed the asciidoc source in %v ", duration)
	b, err := json.MarshalIndent(stats.ChoiceAltCnt, "", "  ")
	if err != nil {
		log.Warnf("failed to produce stats ", err)
	}
	log.Infof("parsing stats: %s", string(b))
	return convertToHTML(ctx, doc, output, options...)
}

func convertToHTML(ctx context.Context, doc interface{}, output io.Writer, options ...renderer.Option) (map[string]interface{}, error) {
	start := time.Now()
	metadata, err := htmlrenderer.Render(renderer.Wrap(ctx, doc.(types.Document), options...), output)
	if err != nil {
		return nil, errors.Wrapf(err, "error while rendering the document")
	}
	log.Debugf("Done processing document")
	duration := time.Since(start)
	log.Infof("rendered the HTML output in %v", duration)
	return metadata, nil
}
