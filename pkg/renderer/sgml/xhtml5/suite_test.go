package xhtml5_test

import (
	"bytes"
	"github.com/bytesparadise/libasciidoc"
	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
	log "github.com/sirupsen/logrus"
	"strings"

	"testing"

	_ "github.com/bytesparadise/libasciidoc/testsupport"
)

func RenderXHTML(actual string, settings ...configuration.Setting) (string, error) {
	config := configuration.NewConfiguration(settings...)
	configuration.WithBackEnd("xhtml5")(&config)

	contentReader := strings.NewReader(actual)
	resultWriter := bytes.NewBuffer(nil)
	_, err := libasciidoc.Convert(contentReader, resultWriter, config)
	if err != nil {
		return "", err
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debug(resultWriter.String())
	}
	return resultWriter.String(), nil
}

// RenderXHTML5Title renders the HTML body using the given source
func RenderXHTML5Title(actual string, options ...configuration.Setting) (string, error) {
	config := configuration.NewConfiguration(configuration.WithBackEnd("xhtml5"))
	for _, set := range options {
		set(&config)
	}
	contentReader := strings.NewReader(actual)
	resultWriter := bytes.NewBuffer(nil)
	metadata, err := libasciidoc.Convert(contentReader, resultWriter, config)
	if err != nil {
		return "", err
	}
	// if strings.Contains(m.expected, "{{.LastUpdated}}") {
	// 	m.expected = strings.Replace(m.expected, "{{.LastUpdated}}", metadata.LastUpdated, 1)
	// }
	return metadata.Title, nil
}

func TestXHtml5(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "XHtml5 Suite")
}
