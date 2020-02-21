package testsupport

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/parser"
)

// ParseDocumentBlock parses the actual source with the `DocumentBlock` entrypoint in the grammar
func ParseDocumentBlock(actual string) (interface{}, error) {
	r := strings.NewReader(actual)
	opts := []parser.Option{parser.Entrypoint("DocumentBlock")}
	// if os.Getenv("DEBUG") == "true" {
	// 	opts = append(opts, parser.Debug(true))
	// }
	return parser.ParseReader("", r, opts...)
}
