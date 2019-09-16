package testsupport

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"

	gomegatypes "github.com/onsi/gomega/types"
	"github.com/pkg/errors"
)

// ------------------------
// Equal document
// ------------------------

// EqualDocument a custom matcher to verify that a document matches the expectation
func EqualDocument(expected interface{}) gomegatypes.GomegaMatcher {
	return &documentMatcher{
		expected: expected,
	}
}

type documentMatcher struct {
	expected interface{}
	actual   types.Document
}

func (m *documentMatcher) Match(actual interface{}) (success bool, err error) {
	content, ok := actual.(string)
	if !ok {
		return false, errors.Errorf("EqualDocument matcher expects a string (actual: %T)", actual)
	}
	r := strings.NewReader(content)
	m.actual, err = parser.ParseDocument("", r)
	if err != nil {
		return false, err
	}
	return reflect.DeepEqual(m.expected, m.actual), nil
}

func (m *documentMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("expected documents to match:\n\texpected: '%v'\n\tactual: '%v'", m.expected, m.actual)
}

func (m *documentMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("expected documents not to match:\n\texpected: '%v'\n\tactual: '%v'", m.expected, m.actual)
}

// ----------------------
// Document metadata
// ----------------------

// HaveMetadata a custom matcher to verify that a document has the expected metadata
func HaveMetadata(expected interface{}) gomegatypes.GomegaMatcher {
	return &metadataMatcher{
		expected: expected,
	}
}

type metadataMatcher struct {
	expected interface{}
	actual   types.DocumentAttributes
}

func (m *metadataMatcher) Match(actual interface{}) (success bool, err error) {
	source, ok := actual.(types.Document)
	if !ok {
		return false, errors.Errorf("HaveMetadata matcher expects a Document (actual: %T)", actual)
	}
	ctx := renderer.Wrap(context.Background(), source)
	renderer.ProcessDocumentHeader(ctx)
	m.actual = ctx.Document.Attributes
	return reflect.DeepEqual(m.expected, m.actual), nil
}

func (m *metadataMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("expected document metadata to match:\n\texpected: '%v'\n\tactual: '%v'", m.expected, m.actual)
}

func (m *metadataMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("expected document metadata not to match:\n\texpected: '%v'\n\tactual: '%v'", m.expected, m.actual)
}
