package testsupport

import (
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/renderer/html5"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

// TableOfContents returns the Table of Contents for the given document
func TableOfContents(actual types.Document) (types.TableOfContents, error) {
	ctx := renderer.NewContext(actual)
	return html5.NewTableOfContents(ctx)
}
