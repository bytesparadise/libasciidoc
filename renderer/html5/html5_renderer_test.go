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
	// simple quotes
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
	t.Run("italic content alone", func(t *testing.T) {
		// given
		content := "_italic content_"
		expected := `<div class="paragraph"><p><em>italic content</em></p></div>`
		verify(t, expected, content)
	})
	t.Run("italic content in sentence", func(t *testing.T) {
		// given
		content := "some _italic content_."
		expected := `<div class="paragraph"><p>some <em>italic content</em>.</p></div>`
		verify(t, expected, content)
	})
	t.Run("monospace content alone", func(t *testing.T) {
		// given
		content := "`monospace content`"
		expected := `<div class="paragraph"><p><code>monospace content</code></p></div>`
		verify(t, expected, content)
	})
	t.Run("monospace content in sentence", func(t *testing.T) {
		// given
		content := "some `monospace content`."
		expected := `<div class="paragraph"><p>some <code>monospace content</code>.</p></div>`
		verify(t, expected, content)
	})
	// nested quotes
	t.Run("italic content within bold quote in sentence", func(t *testing.T) {
		// given
		content := "some *bold and _italic content_* together."
		expected := `<div class="paragraph"><p>some <strong>bold and <em>italic content</em></strong> together.</p></div>`
		verify(t, expected, content)
	})
	t.Run("italic content within invalid bold quote in sentence", func(t *testing.T) {
		// given
		content := "some *bold and _italic content_ * together."
		expected := `<div class="paragraph"><p>some *bold and <em>italic content</em> * together.</p></div>`
		verify(t, expected, content)
	})
	t.Run("invalid italic content within bold quote in sentence", func(t *testing.T) {
		// given
		content := "some *bold and _italic content _ together*."
		expected := `<div class="paragraph"><p>some <strong>bold and _italic content _ together</strong>.</p></div>`
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
