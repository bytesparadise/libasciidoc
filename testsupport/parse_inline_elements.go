package testsupport

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
)

// ParseInlineElements parses the actual source with the options
func ParseInlineElements(actual string, options ...parser.Option) (interface{}, error) {
	r := strings.NewReader(actual)
	result, err := parser.ParseReader("", r, append(options, parser.Entrypoint("InlineElements"))...)
	if err == nil && log.IsLevelEnabled(log.DebugLevel) {
		spew.Fdump(log.StandardLogger().Out, result)
	}
	return result, err
}
