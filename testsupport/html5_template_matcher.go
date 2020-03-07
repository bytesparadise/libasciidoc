package testsupport

import (
	"fmt"
	"strings"
	"time"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	gomegatypes "github.com/onsi/gomega/types"
	"github.com/pkg/errors"
)

// --------------------
// Render HTML5 Title
// --------------------

// MatchHTML5Template a custom matcher to verify that a document renders as the given template
// which will be processed with the given args
func MatchHTML5Template(expected string, lastUpdated time.Time) gomegatypes.GomegaMatcher {
	return &html5TemplateMatcher{
		expected:    expected,
		lastUpdated: lastUpdated,
	}
}

type html5TemplateMatcher struct {
	actual      string
	expected    string
	lastUpdated time.Time
}

func (m *html5TemplateMatcher) Match(actual interface{}) (success bool, err error) {
	if a, ok := actual.(string); ok {
		m.actual = a
	} else {
		return false, errors.Errorf("MatchHTML5Template matcher expects a string (actual: %T)", actual)
	}
	m.expected = strings.Replace(m.expected, "{{.LastUpdated}}", m.lastUpdated.Format(configuration.LastUpdatedFormat), 1)
	return m.expected == m.actual, nil
}

func (m *html5TemplateMatcher) FailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected HTML5 documents to match:\n\texpected: '%v'\n\tactual:   '%v'", m.expected, m.actual)
}

func (m *html5TemplateMatcher) NegatedFailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected HTML5 documents not to match:\n\texpected: '%v'\n\tactual:   '%v'", m.expected, m.actual)
}
