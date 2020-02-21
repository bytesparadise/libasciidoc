package testsupport

import (
	"bytes"
	"strings"
	"time"

	"github.com/bytesparadise/libasciidoc"
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

// DocumentMetadata processes the actual input into a document and returns its metadata
func DocumentMetadata(actual string, lastUpdated time.Time) (types.Metadata, error) {
	return libasciidoc.ConvertToHTML("", strings.NewReader(actual),
		bytes.NewBuffer(nil),
		renderer.IncludeHeaderFooter(false),
		renderer.LastUpdated(lastUpdated))
}
