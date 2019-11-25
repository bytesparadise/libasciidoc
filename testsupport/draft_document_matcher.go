package testsupport

import (
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"

	gomegatypes "github.com/onsi/gomega/types"
	"github.com/pkg/errors"
)

// BecomeDraftDocument a custom matcher to verify that a draft document matches the expectation
func BecomeDraftDocument(expected types.DraftDocument, options ...interface{}) gomegatypes.GomegaMatcher {
	m := &draftDocumentMatcher{
		expected:      expected,
		preprocessing: true,
		filename:      "test.adoc",
	}
	for _, o := range options {
		if configure, ok := o.(BecomeDraftDocumentOption); ok {
			configure(m)
		} else if configure, ok := o.(FilenameOption); ok {
			configure(m)
		}
	}
	return m
}

type draftDocumentMatcher struct {
	filename      string
	preprocessing bool
	expected      interface{}
	actual        interface{}
	comparison    comparison
}

func (m *draftDocumentMatcher) setFilename(f string) {
	m.filename = f
}

func (m *draftDocumentMatcher) Match(actual interface{}) (success bool, err error) {
	content, ok := actual.(string)
	if !ok {
		return false, errors.Errorf("BecomeDocumentBlock matcher expects a string (actual: %T)", actual)
	}
	r := strings.NewReader(content)
	if !m.preprocessing {
		m.actual, err = parser.ParseReader(m.filename, r, parser.Entrypoint("AsciidocDocument"))
	} else {
		m.actual, err = parser.ParseDraftDocument(m.filename, r)
	}
	if err != nil {
		return false, err
	}
	m.comparison = compare(m.actual, m.expected)
	return m.comparison.diffs == "", nil
}

func (m *draftDocumentMatcher) FailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected draft documents to match:\n%s", m.comparison.diffs)
}

func (m *draftDocumentMatcher) NegatedFailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected draft documents not to match:\n%s", m.comparison.diffs)
}
