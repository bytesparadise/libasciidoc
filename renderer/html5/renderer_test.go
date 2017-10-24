package html5_test

import (
	"bytes"
	"context"
	"strings"

	"github.com/bytesparadise/libasciidoc/parser"
	"github.com/bytesparadise/libasciidoc/renderer"
	"github.com/bytesparadise/libasciidoc/renderer/html5"
	"github.com/bytesparadise/libasciidoc/types"
	"github.com/davecgh/go-spew/spew"
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func verify(t GinkgoTInterface, expected, content string, options ...renderer.Option) {
	t.Logf("processing '%s'", content)
	reader := strings.NewReader(content)
	doc, err := parser.ParseReader("", reader)
	require.Nil(t, err, "Error found while parsing the document")
	actualDocument := doc.(*types.Document)
	t.Logf("actual document: `%s`", spew.Sdump(actualDocument))
	rendererCtx := renderer.Wrap(context.Background(), *actualDocument, options...)
	buff := bytes.NewBuffer(nil)
	err = html5.Render(rendererCtx, buff)
	t.Log("* Done processing document:")
	require.Nil(t, err)
	require.Empty(t, err)
	result := buff.String()
	expected = strings.Replace(expected, "\t", "", -1)
	if strings.Contains(expected, "{{.LastUpdated}}") {
		expected = strings.Replace(expected, "{{.LastUpdated}}", rendererCtx.LastUpdated(), 1)
	}
	t.Logf("** Actual output:\n`%s`\n", result)
	t.Logf("** Expected output:\n`%s`\n", expected) // remove tabs that can be inserted by VSCode while formatting the tests code
	assert.Equal(t, expected, result)
}

func singleLine(content string) string {
	return strings.Replace(content, "\n", "", -1)
}
