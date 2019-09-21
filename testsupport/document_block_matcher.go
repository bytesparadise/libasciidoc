package testsupport

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/parser"

	"github.com/onsi/gomega/types"
	"github.com/pkg/errors"
)

// EqualDocumentBlock a custom matcher to verify that a document block matches the expectation
func EqualDocumentBlock(expected interface{}) types.GomegaMatcher {
	return &documentBlockMatcher{
		expected: expected,
	}
}

type documentBlockMatcher struct {
	expected interface{}
	actual   interface{}
}

func (m *documentBlockMatcher) Match(actual interface{}) (success bool, err error) {
	content, ok := actual.(string)
	if !ok {
		return false, errors.Errorf("EqualDocumentBlock matcher expects a string (actual: %T)", actual)
	}
	r := strings.NewReader(content)
	opts := []parser.Option{parser.Entrypoint("DocumentBlock")}
	// if os.Getenv("DEBUG") == "true" {
	// 	opts = append(opts, parser.Debug(true))
	// }
	m.actual, err = parser.ParseReader("", r, opts...)
	if err != nil {
		return false, err
	}
	return reflect.DeepEqual(m.expected, m.actual), nil
}

func (m *documentBlockMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("expected document blocks to match:\n\texpected: '%v'\n\tactual'%v'", m.expected, m.actual)
}

func (m *documentBlockMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("expected document blocks not to match:\n\texpected: '%v'\n\tactual'%v'", m.expected, m.actual)
}
