package renderer

import (
	"fmt"
	"io"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/renderer/sgml/html5"
	"github.com/bytesparadise/libasciidoc/pkg/renderer/sgml/xhtml5"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func Render(doc *types.Document, config *configuration.Configuration, output io.Writer) (types.Metadata, error) {
	switch config.BackEnd {
	case "html", "html5":
		return html5.Render(doc, config, output)
	case "xhtml", "xhtml5":
		return xhtml5.Render(doc, config, output)
	default:
		return types.Metadata{}, fmt.Errorf("backend '%s' not supported", config.BackEnd)
	}
}
