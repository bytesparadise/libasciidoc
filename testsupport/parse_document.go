package testsupport

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

// ParseDocument parses the actual value into a Document
func ParseDocument(actual string) (types.Document, error) {
	r := strings.NewReader(actual)
	return parser.ParseDocument(r, configuration.NewConfiguration()) //, parser.Debug(true))
}
