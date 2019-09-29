package testsupport

import (
	"context"
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"

	gomegatypes "github.com/onsi/gomega/types"
	"github.com/pkg/errors"
)

// HavePreamble a custom matcher to verify that a document matches the expectation
func HavePreamble(expected types.Document) gomegatypes.GomegaMatcher {
	return &preambleMatcher{
		expected: expected,
	}
}

type preambleMatcher struct {
	expected   types.Document
	actual     types.Document
	comparison comparison
}

func (m *preambleMatcher) Match(actual interface{}) (success bool, err error) {
	doc, ok := actual.(types.Document)
	if !ok {
		return false, errors.Errorf("HavePreamble matcher expects a Document (actual: %T)", actual)
	}
	ctx := renderer.Wrap(context.Background(), doc)
	renderer.IncludePreamble(ctx)
	m.actual = ctx.Document
	m.comparison = compare(m.actual, m.expected)
	return m.comparison.diffs == "", nil
}

func (m *preambleMatcher) FailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected documents to match:\n%s", m.comparison.diffs)
}

func (m *preambleMatcher) NegatedFailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected documents not to match:\n%s", m.comparison.diffs)
}
