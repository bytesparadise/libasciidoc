package testsupport

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

// ParseDocument parses the actual value into a Document
func ParseDocument(actual string, settings ...configuration.Setting) (types.Document, error) {
	r := strings.NewReader(actual)
	doc, err := parser.ParseDocument(r, configuration.NewConfiguration(settings...))
	if err == nil && logrus.IsLevelEnabled(logrus.DebugLevel) {
		log.Debug("final document:")
		spew.Dump(doc)
	}
	return doc, err
}
