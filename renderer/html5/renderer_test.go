package html5_test

import (
	"bytes"
	"context"
	"strings"

	asciidoc "github.com/bytesparadise/libasciidoc/context"
	"github.com/bytesparadise/libasciidoc/parser"
	. "github.com/bytesparadise/libasciidoc/renderer/html5"
	"github.com/bytesparadise/libasciidoc/types"
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func verify(t GinkgoTInterface, expected, content string) {
	expected = strings.Replace(expected, "\t", "", -1)
	t.Logf("processing '%s'", content)
	reader := strings.NewReader(content)
	doc, err := parser.ParseReader("", reader)
	require.Nil(t, err, "Error found while parsing the document")
	actualDocument := doc.(*types.Document)
	t.Logf("Actual document:\n%s", actualDocument.String(1))
	buff := bytes.NewBuffer(nil)
	ctx := asciidoc.Wrap(context.Background(), *actualDocument)
	err = Render(ctx, buff, nil)
	t.Log("Done processing document")
	require.Nil(t, err)
	require.Empty(t, err)
	result := string(buff.Bytes())
	t.Logf("** Actual output:\n`%s`\n", result)
	t.Logf("** Expected output:\n`%s`\n", expected) // remove tabs that can be inserted by VSCode while formatting the tests code
	assert.Equal(t, expected, result)
}

func singleLine(content string) string {
	return strings.Replace(content, "\n", "", -1)
}