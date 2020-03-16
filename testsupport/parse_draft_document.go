package testsupport

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/parser"
)

// ParseDraftDocument parses the actual source with the options
func ParseDraftDocument(actual string, options ...interface{}) (interface{}, error) {
	r := strings.NewReader(actual)
	c := &drafDocumentParserConfig{
		preprocessing: true,
		filename:      "test.adoc",
	}
	parserOptions := []parser.Option{}
	for _, o := range options {
		switch set := o.(type) {
		case BecomeDraftDocumentOption:
			set(c)
		case FilenameOption:
			set(c)
		case parser.Option:
			parserOptions = append(parserOptions, set)
		}
	}

	if !c.preprocessing {
		return parser.ParseReader(c.filename, r, append(parserOptions, parser.Entrypoint("AsciidocDocument"))...)
	}
	config := configuration.NewConfiguration(configuration.WithFilename(c.filename))
	return parser.ParseDraftDocument(r, config, parserOptions...)
}

type drafDocumentParserConfig struct {
	filename      string
	preprocessing bool
}

func (c *drafDocumentParserConfig) setFilename(f string) {
	c.filename = f
}
