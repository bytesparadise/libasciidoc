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

// MatchDraftDocument a custom matcher to verify that a document matches the given expectation
// Similar to the standard `Equal` matcher, but display a diff when the values don't match
func MatchDraftDocument(expected types.DraftDocument) gomegatypes.GomegaMatcher {
	return &draftDocumentMatcher{
		expected: expected,
	}
}

type draftDocumentMatcher struct {
	expected types.DraftDocument
	diffs    string
}

func (m *draftDocumentMatcher) Match(actual interface{}) (success bool, err error) {
	if _, ok := actual.(types.DraftDocument); !ok {
		return false, errors.Errorf("MatchDraftDocument matcher expects a DraftDocument (actual: %T)", actual)
	}
	if !reflect.DeepEqual(m.expected, actual) {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(spew.Sdump(actual), spew.Sdump(m.expected), true)
		m.diffs = dmp.DiffPrettyText(diffs)
		return false, nil
	}
	return true, nil
}

func (m *draftDocumentMatcher) FailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected draft documents to match:\n%s", m.diffs)
}

func (m *draftDocumentMatcher) NegatedFailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected draft documents not to match:\n%s", m.diffs)
}
