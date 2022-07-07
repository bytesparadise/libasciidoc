package html5

import (
	"io"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/renderer/sgml"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

// Render renders the document to the output, using the SGML renderer configured with the HTML5 templates
func Render(doc *types.Document, config *configuration.Configuration, output io.Writer) (types.Metadata, error) {
	return sgml.Render(doc, config, output, templates)
}
