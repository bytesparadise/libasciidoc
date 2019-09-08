package html5_test

import (
	"bytes"
	"context"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/renderer/html5"
	"github.com/davecgh/go-spew/spew"
	. "github.com/onsi/ginkgo"
	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func verify(filename, expected, content string, rendererOpts ...renderer.Option) {
	reader := strings.NewReader(content)
	doc, err := parser.ParseDocument(filename, reader)
	require.NoError(GinkgoT(), err, "Error found while parsing the document")
	GinkgoT().Logf("actual document: `%s`", spew.Sdump(doc))
	buff := bytes.NewBuffer(nil)
	rendererCtx := renderer.Wrap(context.Background(), doc, rendererOpts...)
	// insert tables of contents, preamble and process file inclusions
	err = renderer.Prerender(rendererCtx)
	require.NoError(GinkgoT(), err)
	_, err = html5.Render(rendererCtx, buff)
	require.NoError(GinkgoT(), err)
	if strings.Contains(expected, "{{.LastUpdated}}") {
		expected = strings.Replace(expected, "{{.LastUpdated}}", rendererCtx.LastUpdated(), 1)
	}
	GinkgoT().Log("* Done processing document:")
	result := buff.String()
	GinkgoT().Logf("** Actual output:\n`%s`\n", result)
	GinkgoT().Logf("** expected output:\n`%s`\n", expected) // remove tabs that can be inserted by VSCode while formatting the tests code
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(result, expected, true)
	assert.Equal(GinkgoT(), expected, result, dmp.DiffPrettyText(diffs))
}
