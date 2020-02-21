package testsupport

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/parser"
)

// ParseDraftDocument parses the actual source with the options
func ParseDraftDocument(actual string, options ...interface{}) (interface{}, error) {
	r := strings.NewReader(actual)
	c := &drafDocumentParserConfig{
		preprocessing: true,
		filename:      "test.adoc",
	}
	for _, o := range options {
		if configure, ok := o.(BecomeDraftDocumentOption); ok {
			configure(c)
		} else if configure, ok := o.(FilenameOption); ok {
			configure(c)
		}
	}
	if !c.preprocessing {
		return parser.ParseReader(c.filename, r, parser.Entrypoint("AsciidocDocument"))
	}
	return parser.ParseDraftDocument(c.filename, r)
}

type drafDocumentParserConfig struct {
	filename      string
	preprocessing bool
}

func (c *drafDocumentParserConfig) setFilename(f string) {
	c.filename = f
}
