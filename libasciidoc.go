// Package libasciidoc is an open source Go library that converts Asciidoc
// content into HTML.
package libasciidoc

import (
	"io"
	"os"
	"strings"
	"time"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
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
func ConvertFile(output io.Writer, config *configuration.Configuration) (types.Metadata, error) {
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
func Convert(source io.Reader, output io.Writer, config *configuration.Configuration) (types.Metadata, error) {

	start := time.Now()
	defer func() {
		duration := time.Since(start)
		log.Debugf("rendered the output in %v", duration)
	}()
	p, err := parser.Preprocess(source, config)
	if err != nil {
		return types.Metadata{}, err
	}
	// log.Debugf("parsing the asciidoc source...")
	doc, err := parser.ParseDocument(strings.NewReader(p), config)
	if err != nil {
		return types.Metadata{}, err
	}
	// validate the document
	doctype := config.Attributes.GetAsStringWithDefault(types.AttrDocType, "article")
	problems, err := validator.Validate(doc, doctype)
	if err != nil {
		return types.Metadata{}, err
	}
	if len(problems) > 0 {
		// if any problem found, change the doctype to render the document as a regular article
		log.Warnf("changing doctype to 'article' because problems were found in the document: %v", problems)
		config.Attributes[types.AttrDocType] = "article" // switch to `article` rendering (in case it was a manpage with problems)
		for _, problem := range problems {
			switch problem.Severity {
			case validator.Error:
				log.Error(problem.Message)
			case validator.Warning:
				log.Warn(problem.Message)
			}
		}
	}
	// render
	metadata, err := renderer.Render(doc, config, output)
	if err != nil {
		return types.Metadata{}, err
	}
	// log.Debugf("Done processing document")
	return metadata, nil

}
