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

func verifyWithPreprocessing(t GinkgoTInterface, expectedResult interface{}, content string, options ...parser.Option) {
	log.Debugf("processing: %s", content)
	r := strings.NewReader(content)
	allOptions := append(options)
	preparsedDoc, err := parser.PreparseDocument("", r, allOptions...)
	require.NoError(t, err)
	result, err := parser.Parse("", preparsedDoc, allOptions...)
	require.NoError(t, err)
	t.Logf("actual document: `%s`", spew.Sdump(result))
	t.Logf("expected document: `%s`", spew.Sdump(expectedResult))
	assert.EqualValues(t, expectedResult, result)
}

func verifyWithoutPreprocessing(t GinkgoTInterface, expectedResult interface{}, content string, options ...parser.Option) {
	log.Debugf("processing: %s", content)
	r := strings.NewReader(content)
	allOptions := append(options)
	result, err := parser.ParseReader("", r, allOptions...)
	require.NoError(t, err)
	t.Logf("actual document: `%s`", spew.Sdump(result))
	t.Logf("expected document: `%s`", spew.Sdump(expectedResult))
	assert.EqualValues(t, expectedResult, result)
}

func verifyError(t GinkgoTInterface, content string, options ...parser.Option) {
	log.Debugf("processing: %s", content)
	reader := strings.NewReader(content)
	allOptions := append(options, parser.Recover(false))
	_, err := parser.ParseDocument("", reader, allOptions...)
	require.Error(t, err)
}
