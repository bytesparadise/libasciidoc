package testsupport

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"

	gomegatypes "github.com/onsi/gomega/types"
	"github.com/pkg/errors"
)

// HaveAttributes a custom matcher to verify that a document has the expected metadata
func HaveAttributes(expected interface{}) gomegatypes.GomegaMatcher {
	return &documentAttributesMatcher{
		expected: expected,
	}
}

type documentAttributesMatcher struct {
	expected   interface{}
	actual     interface{}
	comparison comparison
}

func (m *documentAttributesMatcher) Match(actual interface{}) (success bool, err error) {
	source, ok := actual.(types.Document)
	if !ok {
		return false, errors.Errorf("HaveAttributes matcher expects a Document (actual: %T)", actual)
	}
	ctx := renderer.NewContext(source)
	renderer.ProcessDocumentHeader(ctx)
	m.actual = ctx.Document.Attributes
	m.comparison = compare(m.actual, m.expected)
	return m.comparison.diffs == "", nil
}

func (m *documentAttributesMatcher) FailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected document metadata to match:\n%s", m.comparison.diffs)
}

func (m *documentAttributesMatcher) NegatedFailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected document metadata not to match:\n%s", m.comparison.diffs)
}
