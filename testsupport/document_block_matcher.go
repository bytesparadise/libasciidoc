package testsupport

import (
	"fmt"
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
	expected   interface{}
	actual     interface{}
	comparison comparison
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
	m.comparison = compare(m.actual, m.expected)
	return m.comparison.diffs == "", nil
}

func (m *documentBlockMatcher) FailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected document blocks to match:\n%s", m.comparison.diffs)
}

func (m *documentBlockMatcher) NegatedFailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected document blocks not to match:\n%s", m.comparison.diffs)
}
