package testsupport

import (
	"fmt"

	. "github.com/onsi/ginkgo" //nolint go-lint
	gomegatypes "github.com/onsi/gomega/types"
	"github.com/pkg/errors"
	"github.com/sergi/go-diff/diffmatchpatch"
)

// MatchHTML a custom matcher to verify that a document renders as the given template
// which will be processed with the given args
func MatchHTML(expected string) gomegatypes.GomegaMatcher {
	return &htmlMatcher{
		expected: expected,
	}
}

type htmlMatcher struct {
	expected string
	diffs    string
}

func (m *htmlMatcher) Match(actual interface{}) (success bool, err error) {
	if _, ok := actual.(string); !ok {
		return false, errors.Errorf("MatchHTML matcher expects a string (actual: %T)", actual)
	}
	if m.expected != actual {
		GinkgoT().Logf("actual HTML:\n%s", actual)
		GinkgoT().Logf("expected HTML:\n%s", m.expected)
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(actual.(string), m.expected, true)
		m.diffs = dmp.DiffPrettyText(diffs)
		return false, nil
	}
	return true, nil
}

func (m *htmlMatcher) FailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected HTML5 documents to match:\n%s", m.diffs)
}

func (m *htmlMatcher) NegatedFailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected HTML5 documents not to match:\n%s", m.diffs)
}
