package testsupport

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/parser"

	"github.com/onsi/gomega/types"
	"github.com/pkg/errors"
)

// EqualPreflightDocument a custom matcher to verify that a preflight document matches the expectation
func EqualPreflightDocument(expected interface{}) types.GomegaMatcher {
	return &preflightDocumentMatcher{
		expected:      expected,
		preprocessing: true,
	}
}

// EqualPreflightDocumentWithoutPreprocessing a custom matcher to verify that a preflight document matches the expectation
func EqualPreflightDocumentWithoutPreprocessing(expected interface{}) types.GomegaMatcher {
	return &preflightDocumentMatcher{
		expected:      expected,
		preprocessing: false,
	}
}

type preflightDocumentMatcher struct {
	expected      interface{}
	preprocessing bool
}

func (m *preflightDocumentMatcher) Match(actual interface{}) (success bool, err error) {
	content, ok := actual.(string)
	if !ok {
		return false, errors.Errorf("EqualDocumentBlock matcher expects a string (actual: %T)", actual)
	}
	r := strings.NewReader(content)
	var doc interface{}
	if !m.preprocessing {
		doc, err = parser.ParseReader("", r, parser.Entrypoint("PreflightDocument"))

	} else {
		doc, err = parser.ParsePreflightDocument("test.adoc", r)
	}
	if err != nil {
		return false, err
	}
	return reflect.DeepEqual(m.expected, doc), nil
}

func (m *preflightDocumentMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("expected preflight documents to match:\n\texpected: '%v'\n\tactual'%v'", m.expected, actual)
}

func (m *preflightDocumentMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("expected preflight documents not to match:\n\texpected: '%v'\n\tactual'%v'", m.expected, actual)
}
