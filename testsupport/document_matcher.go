package testsupport

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"

	gomegatypes "github.com/onsi/gomega/types"
	"github.com/pkg/errors"
)

// EqualDocument a custom matcher to verify that a document matches the expectation
func EqualDocument(expected types.Document) gomegatypes.GomegaMatcher {
	return &documentMatcher{
		expected: expected,
	}
} 

type documentMatcher struct {
	expected types.Document
	actual   types.Document
}

func (m *documentMatcher) Match(actual interface{}) (success bool, err error) {
	content, ok := actual.(string)
	if !ok {
		return false, errors.Errorf("EqualDocument matcher expects a string (actual: %T)", actual)
	}
	r := strings.NewReader(content)
	m.actual, err = parser.ParseDocument("", r)
	if err != nil {
		return false, err
	}
	return reflect.DeepEqual(m.expected, m.actual), nil
}

func (m *documentMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("expected documents to match:\n\texpected: '%v'\n\tactual:   '%v'", m.expected, m.actual)
}

func (m *documentMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("expected documents not to match:\n\texpected: '%v'\n\tactual:   '%v'", m.expected, m.actual)
}
