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

// HaveTableOfContents a custom matcher to verify that a document matches the expectation
func HaveTableOfContents(expected types.Document) gomegatypes.GomegaMatcher {
	return &tocMatcher{
		expected: expected,
	}
}

type tocMatcher struct {
	expected types.Document
	actual   types.Document
}

func (m *tocMatcher) Match(actual interface{}) (success bool, err error) {
	doc, ok := actual.(types.Document)
	if !ok {
		return false, errors.Errorf("HaveTableOfContents matcher expects a Document (actual: %T)", actual)
	}
	ctx := renderer.Wrap(context.Background(), doc)
	renderer.IncludeTableOfContents(ctx)
	m.actual = ctx.Document
	return reflect.DeepEqual(m.expected, m.actual), nil
}

func (m *tocMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("expected document to match:\n\texpected: '%v'\n\tactual:   '%v'", m.expected, m.actual)
}

func (m *tocMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("expected document not to match:\n\texpected: '%v'\n\tactual:   '%v'", m.expected, m.actual)
}
