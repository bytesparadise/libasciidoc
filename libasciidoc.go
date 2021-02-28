// Package libasciidoc is an open source Go library that converts Asciidoc
// content into HTML.
package libasciidoc

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/bytesparadise/libasciidoc/pkg/renderer/sgml/xhtml5"

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

// ConvertFile converts the content of the given filename into an output document.
// The conversion result is written in the given writer `output`, whereas the document metadata (title, etc.) (or an error if a problem occurred) is returned
// as the result of the function call.  The output format is determined by config.Backend (HTML5 default).
func ConvertFile(output io.Writer, config configuration.Configuration) (types.Metadata, error) {
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
	return Convert(file, output, config)
}

// Convert converts the content of the given reader `r` into a full output document, written in the given writer `output`.
// Returns an error if a problem occurred. The default will be HTML5, but depends on the config.BackEnd value.
func Convert(r io.Reader, output io.Writer, config configuration.Configuration) (types.Metadata, error) {

	var render func(*renderer.Context, types.Document, io.Writer) (types.Metadata, error)
	switch config.BackEnd {
	case "html", "html5", "":
		render = html5.Render
	case "xhtml", "xhtml5":
		render = xhtml5.Render
	default:
		return types.Metadata{}, fmt.Errorf("backend '%s' not supported", config.BackEnd)
	}

	start := time.Now()
	defer func() {
		duration := time.Since(start)
		log.Debugf("rendered the output in %v", duration)
	}()
	// log.Debugf("parsing the asciidoc source...")
	doc, err := parser.ParseDocument(r, config)
	if err != nil {
		return types.Metadata{}, err
	}
	// validate the document
	problems, err := validator.Validate(&doc)
	if err != nil {
		return types.Metadata{}, err
	}
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
	metadata, err := render(ctx, doc, output)
	if err != nil {
		return types.Metadata{}, err
	}
	// log.Debugf("Done processing document")
	return metadata, nil

}
