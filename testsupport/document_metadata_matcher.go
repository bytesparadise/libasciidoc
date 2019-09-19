package testsupport

import (
	"context"
	"fmt"
	"reflect"

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
	expected interface{}
	actual   types.DocumentAttributes
}

func (m *metadataMatcher) Match(actual interface{}) (success bool, err error) {
	source, ok := actual.(types.Document)
	if !ok {
		return false, errors.Errorf("HaveMetadata matcher expects a Document (actual: %T)", actual)
	}
	ctx := renderer.Wrap(context.Background(), source)
	renderer.ProcessDocumentHeader(ctx)
	m.actual = ctx.Document.Attributes
	return reflect.DeepEqual(m.expected, m.actual), nil
}

func (m *metadataMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("expected document metadata to match:\n\texpected: '%v'\n\tactual:   '%v'", m.expected, m.actual)
}

func (m *metadataMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("expected document metadata not to match:\n\texpected: '%v'\n\tactual:   '%v'", m.expected, m.actual)
}
