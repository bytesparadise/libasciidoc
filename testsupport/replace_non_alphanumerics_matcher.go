package testsupport

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"

	gomegatypes "github.com/onsi/gomega/types"
	"github.com/pkg/errors"
)

// EqualWithoutNonAlphanumeric a custom matcher to verify that an inline content
func EqualWithoutNonAlphanumeric(expected string) gomegatypes.GomegaMatcher {
	return &nonalphanumericMatcher{
		expected: expected,
	}
}

type nonalphanumericMatcher struct {
	expected string
	actual   string
}

func (m *nonalphanumericMatcher) Match(actual interface{}) (success bool, err error) {
	content, ok := actual.(types.InlineElements)
	if !ok {
		return false, errors.Errorf("EqualWithoutNonAlphanumeric matcher expects an InlineElements (actual: %T)", actual)
	}
	m.actual, err = types.ReplaceNonAlphanumerics(content, "_")
	if err != nil {
		return false, err
	}
	return m.expected == m.actual, nil
}

func (m *nonalphanumericMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("expected non alphanumeric values to match:\n\texpected: '%v'\n\tactual:   '%v'", m.expected, m.actual)
}

func (m *nonalphanumericMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("expected non alphanumeric values not to match:\n\texpected: '%v'\n\tactual:   '%v'", m.expected, m.actual)
}
