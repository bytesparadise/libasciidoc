package testsupport

import (
	"fmt"
	"strings"
	"time"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	gomegatypes "github.com/onsi/gomega/types"
	"github.com/pkg/errors"
	"github.com/sergi/go-diff/diffmatchpatch"
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
	diffs       string
}

func (m *html5TemplateMatcher) Match(actual interface{}) (success bool, err error) {
	if _, ok := actual.(string); !ok {
		return false, errors.Errorf("MatchHTML5Template matcher expects a string (actual: %T)", actual)
	}
	m.expected = strings.Replace(m.expected, "{{.LastUpdated}}", m.lastUpdated.Format(configuration.LastUpdatedFormat), 1)
	if m.expected != actual {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(actual.(string), m.expected, true)
		m.diffs = dmp.DiffPrettyText(diffs)
		return false, nil
	}
	return true, nil
}

func (m *html5TemplateMatcher) FailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected HTML5 documents to match:\n%s", m.diffs)
}

func (m *html5TemplateMatcher) NegatedFailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected HTML5 documents not to match:\n%s", m.diffs)
}
