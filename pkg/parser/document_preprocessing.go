package parser

import (
	"io"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
)

// Preprocess reads line by line to look-up and process file inclusions
func Preprocess(source io.Reader, config *configuration.Configuration, opts ...Option) (string, error) {
	ctx := NewParseContext(config, opts...) // each pipeline step will have its own clone of `ctx`
	return processFileInclusions(ctx, source)
}
