package testsupport

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

// ParseDocument parses the actual value into a Document
func ParseDocument(actual string, options ...interface{}) (*types.Document, error) {
	r := strings.NewReader(actual)
	allSettings := []configuration.Setting{configuration.WithFilename("test.adoc")}
	opts := []parser.Option{}
	for _, o := range options {
		switch o := o.(type) {
		case configuration.Setting:
			allSettings = append(allSettings, o)
		case parser.Option:
			opts = append(opts, o)
		default:
			return nil, errors.Errorf("unexpected type of option: '%T'", o)
		}
	}
	return parser.ParseDocument(r, configuration.NewConfiguration(allSettings...), opts...)
}
