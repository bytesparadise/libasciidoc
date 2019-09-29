package testsupport

import (
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/parser"

	"github.com/onsi/gomega/types"
	"github.com/pkg/errors"
)

// BecomePreflightDocument a custom matcher to verify that a preflight document matches the expectation
func BecomePreflightDocument(expected interface{}) types.GomegaMatcher {
	return &preflightDocumentMatcher{
		expected:      expected,
		preprocessing: true,
	}
}

// BecomePreflightDocumentWithoutPreprocessing a custom matcher to verify that a preflight document matches the expectation
func BecomePreflightDocumentWithoutPreprocessing(expected interface{}) types.GomegaMatcher {
	return &preflightDocumentMatcher{
		expected:      expected,
		preprocessing: false,
	}
}

type preflightDocumentMatcher struct {
	preprocessing bool
	expected      interface{}
	actual        interface{}
	comparison    comparison
}

func (m *preflightDocumentMatcher) Match(actual interface{}) (success bool, err error) {
	content, ok := actual.(string)
	if !ok {
		return false, errors.Errorf("EqualDocumentBlock matcher expects a string (actual: %T)", actual)
	}
	r := strings.NewReader(content)
	if !m.preprocessing {
		m.actual, err = parser.ParseReader("", r, parser.Entrypoint("PreflightDocument"))
	} else {
		m.actual, err = parser.ParsePreflightDocument("test.adoc", r)
	}
	if err != nil {
		return false, err
	}
	m.comparison = compare(m.actual, m.expected)
	return m.comparison.diffs == "", nil
}

func (m *preflightDocumentMatcher) FailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected preflight documents to match:\n%s", m.comparison.diffs)
}

func (m *preflightDocumentMatcher) NegatedFailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected preflight documents not to match:\n%s", m.comparison.diffs)
}
