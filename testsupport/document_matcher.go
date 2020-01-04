package testsupport

import (
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"

	gomegatypes "github.com/onsi/gomega/types"
	"github.com/pkg/errors"
)

// BecomeDocument a custom matcher to verify that a document matches the expectation
func BecomeDocument(expected types.Document) gomegatypes.GomegaMatcher {
	return &documentMatcher{
		expected: expected,
	}
}

type documentMatcher struct {
	expected   types.Document
	actual     types.Document
	comparison comparison
}

func (m *documentMatcher) Match(actual interface{}) (success bool, err error) {
	content, ok := actual.(string)
	if !ok {
		return false, errors.Errorf("BecomeDocument matcher expects a string (actual: %T)", actual)
	}
	r := strings.NewReader(content)
	m.actual, err = parser.ParseDocument("", r) //, parser.Debug(true))
	if err != nil {
		return false, err
	}
	m.comparison = compare(m.actual, m.expected)
	return m.comparison.diffs == "", nil
}

func (m *documentMatcher) FailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected documents to match:\n%s", m.comparison.diffs)
}

func (m *documentMatcher) NegatedFailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected documents not to match:\n%s", m.comparison.diffs)
}
