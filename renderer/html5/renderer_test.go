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
	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func verify(t GinkgoTInterface, expected, content string, rendererOpts ...renderer.Option) {
	t.Logf("processing '%s'", content)
	reader := strings.NewReader(content)
	doc, err := parser.ParseReader("", reader)
	require.NoError(t, err, "Error found while parsing the document")
	t.Logf("actual document: `%s`", spew.Sdump(doc))
	buff := bytes.NewBuffer(nil)
	actualDocument := doc.(*types.Document)
	rendererCtx := renderer.Wrap(context.Background(), actualDocument, rendererOpts...)
	// if entrypoint := rendererCtx.Entrypoint(); entrypoint != nil {

	// }
	_, err = html5.Render(rendererCtx, buff)
	require.NoError(t, err)
	if strings.Contains(expected, "{{.LastUpdated}}") {
		expected = strings.Replace(expected, "{{.LastUpdated}}", rendererCtx.LastUpdated(), 1)
	}
	t.Log("* Done processing document:")
	result := buff.String()
	expected = strings.Replace(expected, "\t", "", -1)
	t.Logf("** Actual output:\n`%s`\n", result)
	t.Logf("** Expected output:\n`%s`\n", expected) // remove tabs that can be inserted by VSCode while formatting the tests code
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(result, expected, true)
	assert.Equal(t, expected, result, dmp.DiffPrettyText(diffs))
}

func singleLine(content string) string {
	return strings.Replace(content, "\n", "", -1)
}
