package html5_test

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"
	"os"

	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/renderer/html5"
	"github.com/davecgh/go-spew/spew"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	log "github.com/sirupsen/logrus"
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

func verifyConsoleOutput(console Readable, errorMsg string) {
	GinkgoT().Logf(console.String())
	out := make(map[string]interface{})
	err := json.Unmarshal(console.Bytes(), &out)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(out["level"]).Should(Equal("error"))
	Expect(out["msg"]).Should(Equal(errorMsg))
}

func configureLogger() (Readable, func()) {
	fmtr := log.StandardLogger().Formatter

	buf := bytes.NewBuffer(nil)
	log.SetOutput(buf)
	log.SetFormatter(&log.JSONFormatter{
		DisableTimestamp: true,
	})
	return buf, func() {
		log.SetOutput(os.Stdout)
		log.SetFormatter(fmtr)
	}
}

type Readable interface {
	Bytes() []byte
	String() string
}