package html5

import (
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func generateID(ctx *renderer.Context, attrs types.ElementAttributes) string {
	id := attrs.GetAsString(types.AttrID)
	if id == "" {
		// ignore empty/unset ID
		return ""
	}
	if attrs.GetAsBool(types.AttrCustomID) {
		return id
	}
	// check if idprefix attribute is set, but only apply if ID attribute on element is not custom
	if idPrefix, ok := ctx.Document.Attributes.GetAsString(types.AttrIDPrefix); ok {
		return idPrefix + id
	}
	// default ID prefix is `_`
	return "_" + id
}
