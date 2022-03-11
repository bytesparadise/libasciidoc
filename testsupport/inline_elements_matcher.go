package testsupport

import (
	"fmt"
	"reflect"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/go-cmp/cmp"
	gomegatypes "github.com/onsi/gomega/types"
	"github.com/pkg/errors"
)

// MatchInlineElements a custom matcher to verify that a document matches the given expectation
// Similar to the standard `Equal` matcher, but display a diff when the values don't match
func MatchInlineElements(expected []interface{}) gomegatypes.GomegaMatcher {
	return &inlineElementsMatcher{
		expected: expected,
	}
}

type inlineElementsMatcher struct {
	expected []interface{}
	diffs    string
}

func (m *inlineElementsMatcher) Match(actual interface{}) (success bool, err error) {
	if _, ok := actual.([]interface{}); !ok {
		return false, errors.Errorf("MatchInlineElements matcher expects a []interface{} (actual: %T)", actual)
	}
	if !reflect.DeepEqual(m.expected, actual) {
		m.diffs = cmp.Diff(spew.Sdump(m.expected), spew.Sdump(actual))
		return false, nil
	}
	return true, nil
}

func (m *inlineElementsMatcher) FailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected elements to match:\n%s", m.diffs)
}

func (m *inlineElementsMatcher) NegatedFailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected elements not to match:\n%s", m.diffs)
}
