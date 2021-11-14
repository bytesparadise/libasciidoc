package testsupport

import (
	"bytes"
	"strings"

	"github.com/bytesparadise/libasciidoc"
	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	log "github.com/sirupsen/logrus"
)

// RenderHTML renders the HTML body using the given source
func RenderHTML(actual string, settings ...configuration.Setting) (string, error) {
	config := configuration.NewConfiguration(settings...)
	configuration.WithBackEnd("html5")(config)
	contentReader := strings.NewReader(actual)
	resultWriter := bytes.NewBuffer(nil)
	_, err := libasciidoc.Convert(contentReader, resultWriter, config)
	if err != nil {
		log.Error(err)
		return "", err
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debug(resultWriter.String())
	}
	return resultWriter.String(), nil
}
