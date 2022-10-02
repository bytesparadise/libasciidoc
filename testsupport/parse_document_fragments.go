package testsupport

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

// ParseDocumentFragments parses the actual source with the options
func ParseDocumentFragments(actual string, options ...parser.Option) ([]types.DocumentFragment, error) {
	r := strings.NewReader(actual)
	ctx := parser.NewParseContext(configuration.NewConfiguration(), options...)
	done := make(chan interface{})
	defer close(done)
	// ctx.Opts = append(ctx.Opts, parser.Debug(true))
	fragmentStream := parser.ParseDocumentFragments(ctx, r, done)
	result := []types.DocumentFragment{}
	for f := range fragmentStream {
		result = append(result, f)
	}
	return result, nil
}
