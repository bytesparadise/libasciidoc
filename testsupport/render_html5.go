package testsupport

import (
	"bytes"
	"os"
	"strings"

	"github.com/bytesparadise/libasciidoc"
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/pkg/errors"
)

// RenderHTML5Body renders the HTML body using the given source
func RenderHTML5Body(actual string, options ...interface{}) (string, error) {
	c := &html5RendererConfig{
		filename: "test.adoc",
		opts:     []renderer.Option{},
	}
	for _, o := range options {
		if configure, ok := o.(FilenameOption); ok {
			configure(c)
		} else if opt, ok := o.(renderer.Option); ok {
			c.opts = append(c.opts, opt)
		}
	}
	contentReader := strings.NewReader(actual)
	resultWriter := bytes.NewBuffer(nil)
	_, err := libasciidoc.ConvertToHTML(c.filename, contentReader, resultWriter, c.opts...)
	if err != nil {
		return "", err
	}
	// if strings.Contains(m.expected, "{{.LastUpdated}}") {
	// 	m.expected = strings.Replace(m.expected, "{{.LastUpdated}}", metadata.LastUpdated, 1)
	// }
	return resultWriter.String(), nil
}

// RenderHTML5Title renders the HTML body using the given source
func RenderHTML5Title(actual string, options ...interface{}) (string, error) {
	c := &html5RendererConfig{
		filename: "test.adoc",
		opts:     []renderer.Option{},
	}
	for _, o := range options {
		if configure, ok := o.(FilenameOption); ok {
			configure(c)
		} else if opt, ok := o.(renderer.Option); ok {
			c.opts = append(c.opts, opt)
		}
	}
	contentReader := strings.NewReader(actual)
	resultWriter := bytes.NewBuffer(nil)
	metadata, err := libasciidoc.ConvertToHTML(c.filename, contentReader, resultWriter, c.opts...)
	if err != nil {
		return "", err
	}
	// if strings.Contains(m.expected, "{{.LastUpdated}}") {
	// 	m.expected = strings.Replace(m.expected, "{{.LastUpdated}}", metadata.LastUpdated, 1)
	// }
	return metadata.Title, nil
}

// RenderHTML5Document a custom matcher to verify that a block renders as expected
func RenderHTML5Document(filename string, options ...renderer.Option) (string, error) {
	resultWriter := bytes.NewBuffer(nil)
	options = append(options, renderer.IncludeHeaderFooter(true))
	_, err := libasciidoc.ConvertFileToHTML(filename, resultWriter, options...)
	if err != nil {
		return "", err
	}
	stat, err := os.Stat(filename)
	if err != nil {
		return "", errors.Wrapf(err, "unable to get stats for file '%s'", filename)
	}
	result := strings.Replace(resultWriter.String(), "{{.LastUpdated}}", stat.ModTime().Format(renderer.LastUpdatedFormat), 1)
	return result, nil
}

type html5RendererConfig struct {
	filename string
	opts     []renderer.Option
}

func (c *html5RendererConfig) setFilename(f string) {
	c.filename = f
}
