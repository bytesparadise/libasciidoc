package testsupport

import (
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/parser"

	"github.com/onsi/gomega/types"
	"github.com/pkg/errors"
)

// BecomePreflightDocument a custom matcher to verify that a preflight document matches the expectation
func BecomePreflightDocument(expected interface{}, options ...interface{}) types.GomegaMatcher {
	m := &preflightDocumentMatcher{
		expected:      expected,
		preprocessing: true,
		filename:      "test.adoc",
	}
	for _, o := range options {
		if configure, ok := o.(BecomePreflightDocumentOption); ok {
			configure(m)
		} else if configure, ok := o.(FilenameOption); ok {
			configure(m)
		}
	}
	return m
}

type preflightDocumentMatcher struct {
	filename      string
	preprocessing bool
	expected      interface{}
	actual        interface{}
	comparison    comparison
}

func (m *preflightDocumentMatcher) setFilename(f string) {
	m.filename = f
}

func (m *preflightDocumentMatcher) Match(actual interface{}) (success bool, err error) {
	content, ok := actual.(string)
	if !ok {
		return false, errors.Errorf("EqualDocumentBlock matcher expects a string (actual: %T)", actual)
	}
	r := strings.NewReader(content)
	if !m.preprocessing {
		m.actual, err = parser.ParseReader(m.filename, r, parser.Entrypoint("PreflightAsciidocDocument"))
	} else {
		m.actual, err = parser.ParsePreflightDocument(m.filename, r)
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
