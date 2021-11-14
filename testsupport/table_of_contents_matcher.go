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

// MatchTableOfContents a custom matcher to verify that a TableOfContents matches the given expectation
// Similar to the standard `Equal` matcher, but display a diff when the values don't match
func MatchTableOfContents(expected *types.TableOfContents) gomegatypes.GomegaMatcher {
	return &tableOfContentsMatcher{
		expected: expected,
	}
}

type tableOfContentsMatcher struct {
	expected *types.TableOfContents
	diffs    string
}

func (m *tableOfContentsMatcher) Match(actual interface{}) (success bool, err error) {
	if _, ok := actual.(*types.TableOfContents); !ok {
		return false, errors.Errorf("MatchDocumentFragment matcher expects a *types.TableOfContents (actual: %T)", actual)
	}
	if diff := cmp.Diff(m.expected, actual, opts...); diff != "" {
		if log.IsLevelEnabled(log.DebugLevel) {
			log.Debugf("actual table of contents:\n%s", spew.Sdump(actual))
			log.Debugf("expected table of contents:\n%s", spew.Sdump(m.expected))
		}
		m.diffs = diff
		return false, nil
	}
	return true, nil
}

func (m *tableOfContentsMatcher) FailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected table of contents to match:\n%s", m.diffs)
}

func (m *tableOfContentsMatcher) NegatedFailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected table of contents not to match:\n%s", m.diffs)
}
