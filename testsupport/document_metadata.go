package testsupport

import (
	"bytes"
	"strings"
	"time"

	"github.com/bytesparadise/libasciidoc"
	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

// DocumentMetadata processes the actual input into a document and returns its metadata
func DocumentMetadata(actual string, lastUpdated time.Time) (types.Metadata, error) {
	return libasciidoc.Convert(strings.NewReader(actual),
		bytes.NewBuffer(nil),
		configuration.NewConfiguration(
			configuration.WithLastUpdated(lastUpdated),
			configuration.WithBackEnd("html5")))
}
