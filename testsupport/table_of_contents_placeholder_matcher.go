package testsupport

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"

	gomegatypes "github.com/onsi/gomega/types"
	"github.com/pkg/errors"
)

// HaveTableOfContentsPlaceHolder a custom matcher to verify that the document has the expected TableOfContentsPlaceHolder (at te expected location)
func HaveTableOfContentsPlaceHolder(expected types.Document) gomegatypes.GomegaMatcher {
	return &tocPlaceHolderMatcher{
		expected: expected,
	}
}

type tocPlaceHolderMatcher struct {
	expected   types.Document
	actual     types.Document
	comparison comparison
}

func (m *tocPlaceHolderMatcher) Match(actual interface{}) (success bool, err error) {
	doc, ok := actual.(types.Document)
	if !ok {
		return false, errors.Errorf("HaveTableOfContents matcher expects a Document (actual: %T)", actual)
	}
	ctx := renderer.NewContext(doc)
	ctx = renderer.IncludeTableOfContentsPlaceHolder(ctx)
	m.actual = ctx.Document
	m.comparison = compare(m.actual, m.expected)
	return m.comparison.diffs == "", nil
}

func (m *tocPlaceHolderMatcher) FailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected documents to match:\n%s", m.comparison.diffs)
}

func (m *tocPlaceHolderMatcher) NegatedFailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected documents not to match:\n%s", m.comparison.diffs)
}
