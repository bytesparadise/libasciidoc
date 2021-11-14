package testsupport

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

// ParseDocumentFragments parses the actual source with the options
func ParseDocumentFragments(actual string, options ...interface{}) ([]types.DocumentFragment, error) {
	r := strings.NewReader(actual)
	c := &rawDocumentParserConfig{
		filename: "test.adoc",
	}
	ctx := parser.NewParseContext(configuration.NewConfiguration())
	for _, o := range options {
		switch set := o.(type) {
		case FilenameOption:
			set(c)
		case parser.Option:
			ctx.Opts = append(ctx.Opts, set)
		}
	}
	done := make(chan interface{})
	defer close(done)
	// ctx.Opts = append(ctx.Opts, parser.Debug(true))
	fragmentStream := parser.ParseFragments(ctx, r, done)
	result := []types.DocumentFragment{}
	for f := range fragmentStream {
		result = append(result, f)
	}
	return result, nil
}

type rawDocumentParserConfig struct {
	filename string
}

func (c *rawDocumentParserConfig) setFilename(f string) {
	c.filename = f
}
