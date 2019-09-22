package testsupport

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/bytesparadise/libasciidoc"
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/renderer/html5"
	"github.com/bytesparadise/libasciidoc/pkg/types"

	. "github.com/onsi/ginkgo" // nolint: golint
	gomegatypes "github.com/onsi/gomega/types"
	"github.com/pkg/errors"
	"github.com/sergi/go-diff/diffmatchpatch"
)

// --------------------
// Render HTML5 Element
// --------------------

// RenderHTML5Element a custom matcher to verify that a block renders as the expectation
func RenderHTML5Element(expected string, opts ...renderer.Option) gomegatypes.GomegaMatcher {
	return &html5ElementMatcher{
		expected: expected,
		opts:     opts,
	}
}

type html5ElementMatcher struct {
	expected string
	actual   string
	opts     []renderer.Option
}

func (m *html5ElementMatcher) Match(actual interface{}) (success bool, err error) {
	content, ok := actual.(string)
	if !ok {
		return false, errors.Errorf("RenderHTML5Element matcher expects a string (actual: %T)", actual)
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
	m.actual = buff.String()
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(m.actual, m.expected, true)
	GinkgoT().Log("%v", dmp.DiffPrettyText(diffs))
	return m.expected == m.actual, nil
}

func (m *html5ElementMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("expected HTML5 elements to match:\n\texpected: '%v'\n\tactual:   '%v'", m.expected, m.actual)
}

func (m *html5ElementMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("expected HTML5 elements not to match:\n\texpected: '%v'\n\tactual:   '%v'", m.expected, m.actual)
}

// --------------------
// Render HTML5 Body
// --------------------

// RenderHTML5Body a custom matcher to verify that a block renders as the expectation
func RenderHTML5Body(expected string, opts ...renderer.Option) gomegatypes.GomegaMatcher {
	return &html5BodyMatcher{
		expected: expected,
		opts:     opts,
	}
}

type html5BodyMatcher struct {
	expected string
	actual   string
	opts     []renderer.Option
}

func (m *html5BodyMatcher) Match(actual interface{}) (success bool, err error) {
	content, ok := actual.(string)
	if !ok {
		return false, errors.Errorf("RenderHTML5Body matcher expects a string (actual: %T)", actual)
	}
	contentReader := strings.NewReader(content)
	resultWriter := bytes.NewBuffer(nil)
	_, err = libasciidoc.ConvertToHTML(context.Background(), contentReader, resultWriter, renderer.IncludeHeaderFooter(false))
	if err != nil {
		return false, err
	}
	m.actual = resultWriter.String()
	return m.expected == m.actual, nil
}

func (m *html5BodyMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("expected HTML5 bodies to match:\n\texpected: '%v'\n\tactual:   '%v'", m.expected, m.actual)
}

func (m *html5BodyMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("expected HTML5 bodies not to match:\n\texpected: '%v'\n\tactual:   '%v'", m.expected, m.actual)
}

// --------------------
// Render HTML5 Title
// --------------------

// RenderHTML5Title a custom matcher to verify that a block renders as the expectation
func RenderHTML5Title(expected interface{}, opts ...renderer.Option) gomegatypes.GomegaMatcher {
	return &html5TitleMatcher{
		expected: expected,
		opts:     opts,
	}
}

type html5TitleMatcher struct {
	expected interface{}
	actual   interface{}
	opts     []renderer.Option
}

func (m *html5TitleMatcher) Match(actual interface{}) (success bool, err error) {
	content, ok := actual.(string)
	if !ok {
		return false, errors.Errorf("RenderHTML5Title matcher expects a string (actual: %T)", actual)
	}
	contentReader := strings.NewReader(content)
	resultWriter := bytes.NewBuffer(nil)
	metadata, err := libasciidoc.ConvertToHTML(context.Background(), contentReader, resultWriter, renderer.IncludeHeaderFooter(false))
	if err != nil {
		return false, err
	}
	if metadata == nil {
		return false, errors.New("no metadata returned")
	}
	if m.expected == nil {
		actualTitle, found := metadata[types.AttrTitle]
		m.actual = actualTitle
		return !found, nil
	}

	actualTitle, ok := metadata[types.AttrTitle].(string)
	if !ok {
		return false, errors.Errorf("invalid type of title (%T)", metadata[types.AttrTitle])
	}
	m.actual = actualTitle
	return m.expected == m.actual, nil
}

func (m *html5TitleMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("expected HTML5 titles to match:\n\texpected: '%v'\n\tactual:   '%v'", m.expected, m.actual)
}

func (m *html5TitleMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("expected HTML5 titles not to match:\n\texpected: '%v'\n\tactual:   '%v'", m.expected, m.actual)
}

// ---------------------
// Render HTML5 Document
// ---------------------

// RenderHTML5Document a custom matcher to verify that a block renders as the expectation
func RenderHTML5Document(expected string, opts ...renderer.Option) gomegatypes.GomegaMatcher {
	return &html5DocumentMatcher{
		expected: expected,
		opts:     opts,
	}
}

type html5DocumentMatcher struct {
	expected string
	actual   string
	opts     []renderer.Option
}

func (m *html5DocumentMatcher) Match(actual interface{}) (success bool, err error) {
	content, ok := actual.(string)
	if !ok {
		return false, errors.Errorf("RenderHTML5Body matcher expects a string (actual: %T)", actual)
	}
	contentReader := strings.NewReader(content)
	resultWriter := bytes.NewBuffer(nil)
	lastUpdated := time.Now()
	_, err = libasciidoc.ConvertToHTML(context.Background(), contentReader, resultWriter, renderer.IncludeHeaderFooter(true))
	if err != nil {
		return false, err
	}
	m.expected = strings.Replace(m.expected, "{{.LastUpdated}}", lastUpdated.Format(renderer.LastUpdatedFormat), 1)
	m.actual = resultWriter.String()
	return m.expected == m.actual, nil
}

func (m *html5DocumentMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("expected HTML5 bodies to match:\n\texpected: '%v'\n\tactual:   '%v'", m.expected, m.actual)
}

func (m *html5DocumentMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("expected HTML5 bodies not to match:\n\texpected: '%v'\n\tactual:   '%v'", m.expected, m.actual)
}
