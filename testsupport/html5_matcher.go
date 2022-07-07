package testsupport

import (
	"bytes"
	"fmt"
	"io/ioutil"
	texttemplate "text/template"

	"github.com/google/go-cmp/cmp"
	. "github.com/onsi/ginkgo/v2"
	gomegatypes "github.com/onsi/gomega/types"
	"github.com/pkg/errors"
)

// ------------------------------------------------------
// Match HTML from string
// ------------------------------------------------------

// MatchHTML a custom matcher to verify that a document renders as the given template
// which will be processed with the given args
func MatchHTML(expected string) gomegatypes.GomegaMatcher {
	return &htmlMatcher{
		expected: expected,
	}
}

type htmlMatcher struct {
	expected string
	diffs    string
}

func (m *htmlMatcher) Match(actual interface{}) (success bool, err error) {
	if _, ok := actual.(string); !ok {
		return false, errors.Errorf("MatchHTML matcher expects a string (actual: %T)", actual)
	}
	if m.expected != actual {
		GinkgoT().Logf("actual HTML:\n'%s'", actual)
		GinkgoT().Logf("expected HTML:\n'%s'", m.expected)
		m.diffs = cmp.Diff(m.expected, actual.(string))
		return false, nil
	}
	return true, nil
}

func (m *htmlMatcher) FailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected HTML5 documents to match:\n%s", m.diffs)
}

func (m *htmlMatcher) NegatedFailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected HTML5 documents not to match:\n%s", m.diffs)
}

// ------------------------------------------------------
// Match HTML from file
// ------------------------------------------------------

// MatchHTMLFromFile a custom matcher to verify that a document renders
// as the content of the file with the given name
func MatchHTMLFromFile(filename string) gomegatypes.GomegaMatcher {
	return &htmlFileMatcher{
		filename: filename,
	}
}

type htmlFileMatcher struct {
	filename string
	diffs    string
}

func (m *htmlFileMatcher) Match(actual interface{}) (success bool, err error) {
	expected, err := ioutil.ReadFile(m.filename)
	if err != nil {
		return false, err
	}
	if _, ok := actual.(string); !ok {
		return false, errors.Errorf("MatchHTMLFromFile matcher expects a string (actual: %T)", actual)
	}
	expected = bytes.ReplaceAll(expected, []byte{'\r'}, []byte{})
	if string(expected) != actual {
		GinkgoT().Logf("actual HTML:\n'%s'", actual)
		GinkgoT().Logf("expected HTML:\n'%s'", string(expected))
		m.diffs = cmp.Diff(string(expected), actual)
		return false, nil
	}
	return true, nil
}

func (m *htmlFileMatcher) FailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected HTML5 documents to match:\n%s", m.diffs)
}

func (m *htmlFileMatcher) NegatedFailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected HTML5 documents not to match:\n%s", m.diffs)
}

// ------------------------------------------------------
// Match HTML from template
// ------------------------------------------------------

// MatchHTMLTemplate a custom matcher to verify that a document renders as the given template
// which will be processed with the given args
func MatchHTMLTemplate(expectedTmpl string, data interface{}) gomegatypes.GomegaMatcher {
	return &htmlTemplateMatcher{
		expectedTmpl: expectedTmpl,
		data:         data,
	}
}

type htmlTemplateMatcher struct {
	expected     string
	expectedTmpl string
	data         interface{}
	diffs        string
}

func (m *htmlTemplateMatcher) Match(actual interface{}) (success bool, err error) {
	if _, ok := actual.(string); !ok {
		return false, errors.Errorf("MatchHTMLTemplate matcher expects a string (actual: %T)", actual)
	}
	expectedTmpl, err := texttemplate.New("test").Parse(string(m.expectedTmpl))
	if err != nil {
		return false, err
	}
	out := new(bytes.Buffer)
	err = expectedTmpl.Execute(out, m.data)
	if err != nil {
		return false, err
	}
	m.expected = out.String()
	if m.expected != actual {
		m.diffs = cmp.Diff(m.expected, actual.(string))
		return false, nil
	}
	return true, nil
}

func (m *htmlTemplateMatcher) FailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected HTML5 documents to match:\n%s", m.diffs)
}

func (m *htmlTemplateMatcher) NegatedFailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected HTML5 documents not to match:\n%s", m.diffs)
}

// ------------------------------------------------------
// Match HTML from template file
// ------------------------------------------------------

// MatchHTMLTemplate a custom matcher to verify that a document renders as the given template
// which will be processed with the given args
func MatchHTMLTemplateFile(filename string, data interface{}) gomegatypes.GomegaMatcher {
	return &htmlTemplateFileMatcher{
		filename: filename,
		data:     data,
	}
}

type htmlTemplateFileMatcher struct {
	filename string
	data     interface{}
	diffs    string
}

func (m *htmlTemplateFileMatcher) Match(actual interface{}) (success bool, err error) {
	if _, ok := actual.(string); !ok {
		return false, errors.Errorf("MatchHTMLTemplate matcher expects a string (actual: %T)", actual)
	}

	expected, err := ioutil.ReadFile(m.filename)
	if err != nil {
		return false, err
	}
	expected = bytes.ReplaceAll(expected, []byte{'\r'}, []byte{})
	expectedTmpl, err := texttemplate.New("test").Parse(string(expected))
	if err != nil {
		return false, err
	}
	out := new(bytes.Buffer)
	err = expectedTmpl.Execute(out, m.data)
	if err != nil {
		return false, err
	}
	if out.String() != actual {
		m.diffs = cmp.Diff(out.String(), actual.(string))
		return false, nil
	}
	return true, nil
}

func (m *htmlTemplateFileMatcher) FailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected HTML5 documents to match:\n%s", m.diffs)
}

func (m *htmlTemplateFileMatcher) NegatedFailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected HTML5 documents not to match:\n%s", m.diffs)
}
