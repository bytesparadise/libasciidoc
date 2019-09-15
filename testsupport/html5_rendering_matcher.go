package testsupport

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/renderer/html5"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega/types"
	"github.com/pkg/errors"
	"github.com/sergi/go-diff/diffmatchpatch"
)

// RenderHTML5 a custom matcher to verify that a document renderes as the expectation
func RenderHTML5(expected string, opts ...renderer.Option) types.GomegaMatcher {
	return &html5RenderingMatcher{
		expected: expected,
		opts:     opts,
	}
}

type html5RenderingMatcher struct {
	expected string
	opts     []renderer.Option
}

func (m *html5RenderingMatcher) Match(actual interface{}) (success bool, err error) {
	content, ok := actual.(string)
	if !ok {
		return false, errors.Errorf("RenderHTML5 matcher expects a string (actual: %T)", actual)
	}
	r := strings.NewReader(content)
	doc, err := parser.ParseDocument("test.adoc", r)
	if err != nil {
		return false, err
	}
	buff := bytes.NewBuffer(nil)
	rendererCtx := renderer.Wrap(context.Background(), doc, m.opts...)
	// insert tables of contents, preamble and process file inclusions
	err = renderer.Prerender(rendererCtx)
	if err != nil {
		return false, err
	}
	_, err = html5.Render(rendererCtx, buff)
	if err != nil {
		return false, err
	}
	if strings.Contains(m.expected, "{{.LastUpdated}}") {
		m.expected = strings.Replace(m.expected, "{{.LastUpdated}}", rendererCtx.LastUpdated(), 1)
	}
	result := buff.String()
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(result, m.expected, true)
	GinkgoT().Log("%v", dmp.DiffPrettyText(diffs))
	return m.expected == result, nil
}

func (m *html5RenderingMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("expected HTML5 documents to match:\n\texpected: '%v'\n\tactual'%v'", m.expected, actual)
}

func (m *html5RenderingMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("expected HTML5 documents not to match:\n\texpected: '%v'\n\tactual'%v'", m.expected, actual)
}
