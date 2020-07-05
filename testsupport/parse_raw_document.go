package testsupport

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/parser"
)

// ParseRawDocument parses the actual source with the options
func ParseRawDocument(actual string, options ...interface{}) (interface{}, error) {
	r := strings.NewReader(actual)
	c := &rawDocumentParserConfig{
		fileInclusion: true,
		filename:      "test.adoc",
	}
	parserOptions := []parser.Option{}
	for _, o := range options {
		switch set := o.(type) {
		case RawDocumentParserOption:
			set(c)
		case FilenameOption:
			set(c)
		case parser.Option:
			parserOptions = append(parserOptions, set)
		}
	}
	if !c.fileInclusion {
		return parser.ParseReader(c.filename, r, append(parserOptions, parser.Entrypoint("AsciidocRawDocument"))...)
	}
	config := configuration.NewConfiguration(configuration.WithFilename(c.filename))
	return parser.ParseRawDocument(r, config, parserOptions...)
}

type rawDocumentParserConfig struct {
	fileInclusion bool
	filename      string
}

func (c *rawDocumentParserConfig) setFilename(f string) {
	c.filename = f
}
