package testsupport

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// ParseDocument parses the actual value into a Document
func ParseDocument(actual string, options ...interface{}) (*types.Document, error) {
	allSettings := []configuration.Setting{
		configuration.WithFilename("test.adoc"),
	}
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
	c := configuration.NewConfiguration(allSettings...)
	p, err := parser.Preprocess(strings.NewReader(actual), c, opts...)
	if err != nil {
		return nil, err
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("preparsed document:\n'%s'", p)
	}
	return parser.ParseDocument(strings.NewReader(p), c, opts...)
}
