package testsupport

import (
	"bytes"
	"strings"

	"github.com/bytesparadise/libasciidoc"
	"github.com/bytesparadise/libasciidoc/pkg/configuration"
)

// MetadataTitle returns the title entry from the document metadata
func MetadataTitle(actual string, options ...configuration.Setting) (string, error) {
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
