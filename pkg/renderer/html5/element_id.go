package html5

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	log "github.com/sirupsen/logrus"
)

func generateID(ctx *renderer.Context, attrs types.ElementAttributes) string {
	id := attrs.GetAsString(types.AttrID)
	log.Debugf("default ID: %s", id)
	if id == "" {
		// ignore empty/unset ID
		return ""
	}
	if attrs.GetAsBool(types.AttrCustomID) {
		log.Debugf("has custom ID")
		return id
	}
	// check if idprefix attribute is set, but only apply if ID attribute on element is not custom
	if ctx.Document.Attributes.Has(types.AttrIDPrefix) {
		log.Debugf("has ID prefix")
		return fmt.Sprintf("%s%s", ctx.Document.Attributes.GetAsString(types.AttrIDPrefix), id)
	}
	// default ID prefix is `_`
	log.Debugf("default ID prefix")
	return fmt.Sprintf("_%s", id)
}
