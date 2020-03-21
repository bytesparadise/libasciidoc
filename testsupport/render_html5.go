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

// RenderHTML5Body renders the HTML body using the given source
func RenderHTML5Body(actual string, settings ...configuration.Setting) (string, error) {
	config := configuration.NewConfiguration(settings...)
	contentReader := strings.NewReader(actual)
	resultWriter := bytes.NewBuffer(nil)
	_, err := libasciidoc.ConvertToHTML(contentReader, resultWriter, config)
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
	for _, set := range options {
		set(&config)
	}
	contentReader := strings.NewReader(actual)
	resultWriter := bytes.NewBuffer(nil)
	metadata, err := libasciidoc.ConvertToHTML(contentReader, resultWriter, config)
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
	config := configuration.NewConfiguration(append(options, configuration.WithFilename(filename))...)
	stat, err := os.Stat(filename)
	if err != nil {
		return "", errors.Wrapf(err, "unable to get stats for file '%s'", filename)
	}
	config.LastUpdated = stat.ModTime()
	_, err = libasciidoc.ConvertFileToHTML(resultWriter, config)
	if err != nil {
		return "", err
	}
	// result := strings.Replace(resultWriter.String(), "{{.LastUpdated}}", stat.ModTime().Format(renderer.LastUpdatedFormat), 1)
	return resultWriter.String(), nil
}
