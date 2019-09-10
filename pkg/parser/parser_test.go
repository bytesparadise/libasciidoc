package parser_test

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/davecgh/go-spew/spew"
	. "github.com/onsi/ginkgo"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func verifyDocument(expected interface{}, content string) {
	log.Debugf("processing: %s", content)
	r := strings.NewReader(content)
	preflightDoc, err := parser.ParseDocument("", r)
	require.NoError(GinkgoT(), err)
	GinkgoT().Logf("actual document: `%s`", spew.Sdump(preflightDoc))
	GinkgoT().Logf("expected document: `%s`", spew.Sdump(expected))
	assert.EqualValues(GinkgoT(), expected, preflightDoc)
}

func verifyPreflight(filename string, expected interface{}, content string) {
	log.Debugf("processing %s: %s", filename, content)
	r := strings.NewReader(content)
	preflightDoc, err := parser.ParsePreflightDocument(filename, r)
	require.NoError(GinkgoT(), err)
	GinkgoT().Logf("actual document: `%s`", spew.Sdump(preflightDoc))
	GinkgoT().Logf("expected document: `%s`", spew.Sdump(expected))
	assert.EqualValues(GinkgoT(), expected, preflightDoc)
}

func verifyPreflightWithoutPreprocessing(expected interface{}, content string) {
	log.Debugf("processing: %s", content)
	r := strings.NewReader(content)
	result, err := parser.ParseReader("", r, parser.Entrypoint("PreflightDocument"))
	require.NoError(GinkgoT(), err)
	GinkgoT().Logf("actual document: `%s`", spew.Sdump(result))
	GinkgoT().Logf("expected document: `%s`", spew.Sdump(expected))
	assert.EqualValues(GinkgoT(), expected, result)
}

func verifyDocumentBlock(expected interface{}, content string) {
	log.Debugf("processing: %s", content)
	r := strings.NewReader(content)
	opts := []parser.Option{parser.Entrypoint("DocumentBlock")}
	// if os.Getenv("DEBUG") == "true" {
	// 	opts = append(opts, parser.Debug(true))
	// }
	result, err := parser.ParseReader("", r, opts...)
	require.NoError(GinkgoT(), err)
	GinkgoT().Logf("actual document: `%s`", spew.Sdump(result))
	GinkgoT().Logf("expected document: `%s`", spew.Sdump(expected))
	assert.EqualValues(GinkgoT(), expected, result)
}
