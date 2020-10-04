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

// MatchRawDocument a custom matcher to verify that a document matches the given expectation
// Similar to the standard `Equal` matcher, but display a diff when the values don't match
func MatchRawDocument(expected types.RawDocument) gomegatypes.GomegaMatcher {
	return &rawDocumentMatcher{
		expected: expected,
	}
}

type rawDocumentMatcher struct {
	expected types.RawDocument
	diffs    string
}

func (m *rawDocumentMatcher) Match(actual interface{}) (success bool, err error) {
	if _, ok := actual.(types.RawDocument); !ok {
		return false, errors.Errorf("MatchRawDocument matcher expects a RawDocument (actual: %T)", actual)
	}
	if !reflect.DeepEqual(m.expected, actual) {
		if log.IsLevelEnabled(log.DebugLevel) {
			log.Debug("actual raw document:")
			spew.Fdump(log.StandardLogger().Out, actual)
			log.Debug("expected raw document:")
			spew.Fdump(log.StandardLogger().Out, m.expected)
		}
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(spew.Sdump(actual), spew.Sdump(m.expected), true)
		m.diffs = dmp.DiffPrettyText(diffs)
		return false, nil
	}
	return true, nil
}

func (m *rawDocumentMatcher) FailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected raw documents to match:\n%s", m.diffs)
}

func (m *rawDocumentMatcher) NegatedFailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected raw documents not to match:\n%s", m.diffs)
}
