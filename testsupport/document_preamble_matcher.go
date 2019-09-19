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

// HavePreamble a custom matcher to verify that a document matches the expectation
func HavePreamble(expected types.Document) gomegatypes.GomegaMatcher {
	return &preambleMatcher{
		expected: expected,
	}
}

type preambleMatcher struct {
	expected types.Document
	actual   types.Document
}

func (m *preambleMatcher) Match(actual interface{}) (success bool, err error) {
	doc, ok := actual.(types.Document)
	if !ok {
		return false, errors.Errorf("HavePreamble matcher expects a Document (actual: %T)", actual)
	}
	ctx := renderer.Wrap(context.Background(), doc)
	renderer.IncludePreamble(ctx)
	m.actual = ctx.Document
	return reflect.DeepEqual(m.expected, m.actual), nil
}

func (m *preambleMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("expected document to match:\n\texpected: '%v'\n\tactual:   '%v'", m.expected, m.actual)
}

func (m *preambleMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("expected document not to match:\n\texpected: '%v'\n\tactual:   '%v'", m.expected, m.actual)
}
