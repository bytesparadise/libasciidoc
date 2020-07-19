package testsupport

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/parser"

	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
)

// ParseDocumentBlock parses the actual source with the `DocumentRawBlock` entrypoint in the grammar
func ParseDocumentBlock(actual string, opts ...parser.Option) (interface{}, error) {
	r := strings.NewReader(actual)
	result, err := parser.ParseReader("", r, append(opts, parser.Entrypoint("DocumentRawBlock"))...)
	if err == nil && log.IsLevelEnabled(log.DebugLevel) {
		spew.Fdump(log.StandardLogger().Out, result)
	}
	return result, err
}
