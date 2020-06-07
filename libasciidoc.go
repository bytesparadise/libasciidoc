// Package libasciidoc is an open source Go library that converts Asciidoc
// content into HTML.
package libasciidoc

import (
	"io"
	"os"
	"time"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/renderer/sgml/html5"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/bytesparadise/libasciidoc/pkg/validator"
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
func ConvertFileToHTML(output io.Writer, config configuration.Configuration) (types.Metadata, error) {
	file, err := os.Open(config.Filename)
	if err != nil {
		return types.Metadata{}, errors.Wrapf(err, "error opening %s", config.Filename)
	}
	defer file.Close()
	// use the file mtime as the `last updated` value
	stat, err := os.Stat(config.Filename)
	if err != nil {
		return types.Metadata{}, errors.Wrapf(err, "error opening %s", config.Filename)
	}
	config.LastUpdated = stat.ModTime()
	return ConvertToHTML(file, output, config)
}

// ConvertToHTML converts the content of the given reader `r` into a full HTML document, written in the given writer `output`.
// Returns an error if a problem occurred
func ConvertToHTML(r io.Reader, output io.Writer, config configuration.Configuration) (types.Metadata, error) {
	start := time.Now()
	defer func() {
		duration := time.Since(start)
		log.Debugf("rendered the HTML output in %v", duration)
	}()
	log.Debugf("parsing the asciidoc source...")
	doc, err := parser.ParseDocument(r, config) //, parser.Debug(true))
	if err != nil {
		return types.Metadata{}, err
	}
	// validate the document
	problems := validator.Validate(&doc)
	for _, problem := range problems {
		switch problem.Severity {
		case validator.Error:
			log.Error(problem.Message)
		case validator.Warning:
			log.Warn(problem.Message)
		}
	}
	// render
	ctx := renderer.NewContext(doc, config)
	metadata, err := html5.Render(ctx, doc, output)
	if err != nil {
		return types.Metadata{}, err
	}
	log.Debugf("Done processing document")
	return metadata, nil
}
