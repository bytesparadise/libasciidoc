package testsupport

import (
	"fmt"
	"reflect"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/davecgh/go-spew/spew"
	gomegatypes "github.com/onsi/gomega/types"
	"github.com/pkg/errors"
	"github.com/sergi/go-diff/diffmatchpatch"
	log "github.com/sirupsen/logrus"
)

// MatchDocumentFragmentGroups a custom matcher to verify that a document matches the given expectation
// Similar to the standard `Equal` matcher, but display a diff when the values don't match
func MatchDocumentFragmentGroups(expected []types.DocumentFragment) gomegatypes.GomegaMatcher {
	return &documentFragmentGroupsMatcher{
		expected: expected,
	}
}

type documentFragmentGroupsMatcher struct {
	expected []types.DocumentFragment
	diffs    string
}

func (m *documentFragmentGroupsMatcher) Match(actual interface{}) (success bool, err error) {
	if _, ok := actual.([]types.DocumentFragment); !ok {
		return false, errors.Errorf("MatchDocumentFragmentGroups matcher expects an array of types.DocumentFragment (actual: %T)", actual)
	}
	if !reflect.DeepEqual(m.expected, actual) {
		if log.IsLevelEnabled(log.DebugLevel) {
			log.Debugf("actual document fragments:\n%s", spew.Sdump(actual))
			log.Debugf("expected document fragments:\n%s", spew.Sdump(m.expected))
		}
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(spew.Sdump(actual), spew.Sdump(m.expected), true)
		m.diffs = dmp.DiffPrettyText(diffs)
		return false, nil
	}
	return true, nil
}

func (m *documentFragmentGroupsMatcher) FailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected document fragments to match:\n%s", m.diffs)
}

func (m *documentFragmentGroupsMatcher) NegatedFailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected document fragments not to match:\n%s", m.diffs)
}
