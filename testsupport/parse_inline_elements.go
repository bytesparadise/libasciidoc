package testsupport

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/parser"
)

// ParseInlineElements parses the actual source with the options
func ParseInlineElements(actual string, options ...parser.Option) (interface{}, error) {
	r := strings.NewReader(actual)
	return parser.ParseReader("", r, append(options, parser.Entrypoint("InlineElements"))...)
}
