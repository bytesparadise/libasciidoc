package testsupport

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/go-cmp/cmp"
	gomegatypes "github.com/onsi/gomega/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// MatchDocumentFragment a custom matcher to verify that a document matches the given expectation
// Similar to the standard `Equal` matcher, but display a diff when the values don't match
func MatchDocumentFragment(expected types.DocumentFragment) gomegatypes.GomegaMatcher {
	return &documentFragmentMatcher{
		expected: expected,
	}
}

type documentFragmentMatcher struct {
	expected types.DocumentFragment
	diffs    string
}

func (m *documentFragmentMatcher) Match(actual interface{}) (success bool, err error) {
	if _, ok := actual.(types.DocumentFragment); !ok {
		return false, errors.Errorf("MatchDocumentFragment matcher expects a 'types.DocumentFragment' (actual: %T)", actual)
	}
	if diff := cmp.Diff(m.expected, actual, opts...); diff != "" {
		if log.IsLevelEnabled(log.DebugLevel) {
			log.Debugf("actual document fragment:\n%s", spew.Sdump(actual))
			log.Debugf("expected document fragment:\n%s", spew.Sdump(m.expected))
		}
		m.diffs = diff
		return false, nil
	}
	return true, nil
}

func (m *documentFragmentMatcher) FailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected document fragments to match:\n%s", m.diffs)
}

func (m *documentFragmentMatcher) NegatedFailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected document fragments not to match:\n%s", m.diffs)
}
