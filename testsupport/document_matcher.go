package testsupport

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	gomegatypes "github.com/onsi/gomega/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// MatchDocument a custom matcher to verify that a document matches the given expectation
// Similar to the standard `Equal` matcher, but display a diff when the values don't match
func MatchDocument(expected *types.Document) gomegatypes.GomegaMatcher {
	return &documentMatcher{
		expected: expected,
	}
}

type documentMatcher struct {
	expected *types.Document
	diffs    string
}

var opts = []cmp.Option{cmpopts.IgnoreUnexported(
	types.List{},
	types.DelimitedBlock{},
	types.Footnotes{},
	types.TableOfContents{},
	types.AttributeDeclaration{},
	types.AttributeReference{},
	types.AttributeReset{},
	types.CounterSubstitution{},
	types.PredefinedAttribute{},
)}

func (m *documentMatcher) Match(actual interface{}) (success bool, err error) {
	if _, ok := actual.(*types.Document); !ok {
		return false, errors.Errorf("MatchDocument matcher expects a Document (actual: %T)", actual)
	}
	if diff := cmp.Diff(m.expected, actual, opts...); diff != "" {
		if log.IsLevelEnabled(log.DebugLevel) {
			log.Debugf("actual document:\n%s", spew.Sdump(actual))
			log.Debugf("expected document:\n%s", spew.Sdump(m.expected))
		}
		m.diffs = diff
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
