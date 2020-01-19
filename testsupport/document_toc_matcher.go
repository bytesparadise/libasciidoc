package testsupport

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"

	gomegatypes "github.com/onsi/gomega/types"
	"github.com/pkg/errors"
)

// HaveTableOfContents a custom matcher to verify that a document matches the expectation
func HaveTableOfContents(expected types.Document) gomegatypes.GomegaMatcher {
	return &tocMatcher{
		expected: expected,
	}
}

type tocMatcher struct {
	expected   types.Document
	actual     types.Document
	comparison comparison
}

func (m *tocMatcher) Match(actual interface{}) (success bool, err error) {
	doc, ok := actual.(types.Document)
	if !ok {
		return false, errors.Errorf("HaveTableOfContents matcher expects a Document (actual: %T)", actual)
	}
	ctx := renderer.NewContext(doc)
	ctx = renderer.IncludeTableOfContents(ctx)
	m.actual = ctx.Document
	m.comparison = compare(m.actual, m.expected)
	return m.comparison.diffs == "", nil
}

func (m *tocMatcher) FailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected documents to match:\n%s", m.comparison.diffs)
}

func (m *tocMatcher) NegatedFailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected documents not to match:\n%s", m.comparison.diffs)
}
