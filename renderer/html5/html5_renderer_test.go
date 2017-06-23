package html5_test

import (
	"strings"
	"testing"

	"context"

	"github.com/bytesparadise/libasciidoc/parser"
	. "github.com/bytesparadise/libasciidoc/renderer/html5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHtml5Renderer(t *testing.T) {

	t.Run("bold content alone", func(t *testing.T) {
		// given
		content := "*bold content*"
		expected := `<div class="paragraph"><p><strong>bold content</strong></p></div>`
		verify(t, expected, content)
	})

	t.Run("bold content in sentence", func(t *testing.T) {
		// given
		content := "some *bold content*."
		expected := `<div class="paragraph"><p>some <strong>bold content</strong>.</p></div>`
		verify(t, expected, content)
	})
}

func verify(t *testing.T, expected, content string) {
	document, err := parser.ParseString(content)
	require.Nil(t, err)
	// when
	actual, errs := RenderToString(context.Background(), *document)
	// then
	require.Nil(t, errs)
	assert.Equal(t, singleLine(expected), singleLine(*actual))
}

func singleLine(content string) string {
	return strings.Replace(content, "\n", "", -1)
}
