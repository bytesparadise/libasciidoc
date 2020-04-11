package testsupport

import (
	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/renderer/html5"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

// TableOfContents returns the Table of Contents for the given document
func TableOfContents(doc types.Document) (types.TableOfContents, error) {
	ctx := renderer.NewContext(doc, configuration.NewConfiguration())
	return html5.NewTableOfContents(ctx, doc)
}
