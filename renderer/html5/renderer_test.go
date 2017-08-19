package html5_test

import (
	"bytes"
	"context"
	"strings"

	"github.com/bytesparadise/libasciidoc/parser"
	. "github.com/bytesparadise/libasciidoc/renderer/html5"
	"github.com/bytesparadise/libasciidoc/types"
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func verify(t GinkgoTInterface, expected, content string) {
	t.Logf("processing '%s'", content)
	reader := strings.NewReader(content)
	doc, err := parser.ParseReader("", reader)
	require.Nil(t, err, "Error found while parsing the document")
	actualDocument := doc.(*types.Document)
	buff := bytes.NewBuffer(make([]byte, 0))
	err = Render(context.Background(), *actualDocument, buff)
	t.Log("Done processing document")
	require.Nil(t, err)
	require.Empty(t, err)
	result := string(buff.Bytes())
	t.Logf("** Actual output:\n%s\n", result)
	t.Logf("** Expected output:\n%s\n", expected)
	assert.Equal(t, expected, result)
}

func singleLine(content string) string {
	return strings.Replace(content, "\n", "", -1)
}
