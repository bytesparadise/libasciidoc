package testsupport

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	log "github.com/sirupsen/logrus"

	"github.com/pkg/errors"
)

func PreparseDocument(source string, options ...interface{}) (string, error) {
	settings := []configuration.Setting{
		configuration.WithFilename("test.adoc"),
	}
	opts := []parser.Option{}
	for _, o := range options {
		switch o := o.(type) {
		case configuration.Setting:
			settings = append(settings, o)
		case parser.Option:
			opts = append(opts, o)
		default:
			return "", errors.Errorf("unexpected type of option: '%T'", o)
		}
	}
	result, err := parser.Preprocess(strings.NewReader(source), configuration.NewConfiguration(settings...), opts...)
	if log.IsLevelEnabled(log.DebugLevel) && err == nil {
		log.Debugf("preparsed document:\n'%s'", result)
	}
	return result, err

}
