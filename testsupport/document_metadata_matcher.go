package testsupport

import (
	"context"
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"

	gomegatypes "github.com/onsi/gomega/types"
	"github.com/pkg/errors"
)

// HaveMetadata a custom matcher to verify that a document has the expected metadata
func HaveMetadata(expected interface{}) gomegatypes.GomegaMatcher {
	return &metadataMatcher{
		expected: expected,
	}
}

type metadataMatcher struct {
	expected   interface{}
	actual     interface{}
	comparison comparison
}

func (m *metadataMatcher) Match(actual interface{}) (success bool, err error) {
	source, ok := actual.(types.Document)
	if !ok {
		return false, errors.Errorf("HaveMetadata matcher expects a Document (actual: %T)", actual)
	}
	ctx := renderer.Wrap(context.Background(), source)
	renderer.ProcessDocumentHeader(ctx)
	m.actual = ctx.Document.Attributes
	m.comparison = compare(m.actual, m.expected)
	return m.comparison.diffs == "", nil
}

func (m *metadataMatcher) FailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected document metadata to match:\n%s", m.comparison.diffs)
}

func (m *metadataMatcher) NegatedFailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected document metadata not to match:\n%s", m.comparison.diffs)
}
