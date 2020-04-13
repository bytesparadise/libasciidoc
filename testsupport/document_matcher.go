package testsupport

import (
	"fmt"
	"reflect"

	"github.com/bytesparadise/libasciidoc/pkg/types"

	"github.com/davecgh/go-spew/spew"
	gomegatypes "github.com/onsi/gomega/types"
	"github.com/pkg/errors"
	"github.com/sergi/go-diff/diffmatchpatch"
)

// MatchDocument a custom matcher to verify that a document matches the given expectation
// Similar to the standard `Equal` matcher, but display a diff when the values don't match
func MatchDocument(expected types.Document) gomegatypes.GomegaMatcher {
	return &documentMatcher{
		expected: expected,
	}
}

type documentMatcher struct {
	expected types.Document
	diffs    string
}

func (m *documentMatcher) Match(actual interface{}) (success bool, err error) {
	if _, ok := actual.(types.Document); !ok {
		return false, errors.Errorf("MatchDocument matcher expects a Document (actual: %T)", actual)
	}
	if !reflect.DeepEqual(m.expected, actual) {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(spew.Sdump(actual), spew.Sdump(m.expected), true)
		m.diffs = dmp.DiffPrettyText(diffs)
		return false, nil
	}
	return true, nil
}

func (m *documentMatcher) FailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected documents to match:\n%s", m.diffs)
}

func (m *documentMatcher) NegatedFailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected documents not to match:\n%s", m.diffs)
}
