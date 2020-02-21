package testsupport

import (
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

// IncludePreamble returns a document with the Preamble included at the (expected) location
func IncludePreamble(actual types.Document) types.Document {
	ctx := renderer.NewContext(actual)
	ctx = renderer.IncludePreamble(ctx)
	return ctx.Document
}
