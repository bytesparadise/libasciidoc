package testsupport

import (
	"bytes"
	"os"
	"strings"

	"github.com/bytesparadise/libasciidoc"
	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// Render renders the HTML body using the given source
func Render(actual string, settings ...configuration.Setting) (string, error) {
	config := configuration.NewConfiguration(settings...)
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

// RenderHTML renders the HTML body using the given source
func RenderHTML(actual string, settings ...configuration.Setting) (string, error) {
	config := configuration.NewConfiguration(settings...)
	configuration.WithBackEnd("html5")(&config)
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

// RenderHTML5Title renders the HTML body using the given source
func RenderHTML5Title(actual string, options ...configuration.Setting) (string, error) {
	config := configuration.NewConfiguration()
	configuration.WithBackEnd("html5")(&config)
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

// RenderHTML5Document a custom matcher to verify that a block renders as expected
func RenderHTML5Document(filename string, options ...configuration.Setting) (string, error) {
	resultWriter := bytes.NewBuffer(nil)
	config := configuration.NewConfiguration(options...)
	configuration.WithFilename(filename)(&config)
	configuration.WithBackEnd("html5")(&config)
	stat, err := os.Stat(filename)
	if err != nil {
		return "", errors.Wrapf(err, "unable to get stats for file '%s'", filename)
	}
	config.LastUpdated = stat.ModTime()
	_, err = libasciidoc.ConvertFile(resultWriter, config)
	if err != nil {
		return "", err
	}
	// result := strings.Replace(resultWriter.String(), "{{.LastUpdated}}", stat.ModTime().Format(renderer.LastUpdatedFormat), 1)
	return resultWriter.String(), nil
}
