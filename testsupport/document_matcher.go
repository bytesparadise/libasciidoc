package testsupport

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/parser"

	"github.com/onsi/gomega/types"
	"github.com/pkg/errors"
)

// EqualDocument a custom matcher to verify that a document matches the expectation
func EqualDocument(expected interface{}) types.GomegaMatcher {
	return &documentMatcher{
		expected: expected,
	}
}

type documentMatcher struct {
	expected interface{}
}

func (m *documentMatcher) Match(actual interface{}) (success bool, err error) {
	content, ok := actual.(string)
	if !ok {
		return false, errors.Errorf("EqualDocument matcher expects a string (actual: %T)", actual)
	}
	r := strings.NewReader(content)
	doc, err := parser.ParseDocument("", r)
	if err != nil {
		return false, err
	}
	return reflect.DeepEqual(m.expected, doc), nil
}

func (m *documentMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("expected documents to match:\n\texpected: '%v'\n\tactual'%v'", m.expected, actual)
}

func (m *documentMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("expected documents not to match:\n\texpected: '%v'\n\tactual'%v'", m.expected, actual)
}
