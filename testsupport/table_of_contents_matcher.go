package testsupport

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/renderer/html5"
	"github.com/bytesparadise/libasciidoc/pkg/types"

	gomegatypes "github.com/onsi/gomega/types"
	"github.com/pkg/errors"
)

// HaveTableOfContents a custom matcher to verify that an inline content
func HaveTableOfContents(expected types.TableOfContents) gomegatypes.GomegaMatcher {
	return &tableOfContentsMatcher{
		expected: expected,
	}
}

type tableOfContentsMatcher struct {
	expected   types.TableOfContents
	actual     types.TableOfContents
	comparison comparison
}

func (m *tableOfContentsMatcher) Match(actual interface{}) (success bool, err error) {
	doc, ok := actual.(types.Document)
	if !ok {
		return false, errors.Errorf("HaveTableOfContents matcher expects a Document (actual: %T)", actual)
	}
	m.actual, err = html5.NewTableOfContents(renderer.NewContext(doc))
	if err != nil {
		return false, err
	}
	m.comparison = compare(m.actual, m.expected)
	return m.comparison.diffs == "", nil
}

func (m *tableOfContentsMatcher) FailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected table of contents to match:\n%s", m.comparison.diffs)
}

func (m *tableOfContentsMatcher) NegatedFailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected table of contents not to match:\n%s", m.comparison.diffs)
}
