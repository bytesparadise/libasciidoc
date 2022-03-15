package main

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/bytesparadise/libasciidoc/pkg/plugins"
)

var PreRender plugins.PreRenderFunc = func(doc *types.Document) (*types.Document, error) {
	return doc, nil
}
