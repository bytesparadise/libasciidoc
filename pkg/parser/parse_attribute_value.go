package parser

import (
	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func ParseAttributeValue(value string) ([]interface{}, error) {
	ctx := NewParseContext(configuration.NewConfiguration())
	return processSubstitutions(ctx, []interface{}{
		types.RawLine(value),
	}, headerSubstitutions(), Entrypoint("HeaderGroup"))
}
