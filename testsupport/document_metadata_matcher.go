package testsupport

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/bytesparadise/libasciidoc"
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"

	gomegatypes "github.com/onsi/gomega/types"
	"github.com/pkg/errors"
)

// HaveMetadata a custom matcher to verify that the given metadata are returned from a given document
func HaveMetadata(expected types.Metadata, lastUpdated time.Time) gomegatypes.GomegaMatcher {
	return &metadataMatcher{
		expected:    expected,
		lastUpdated: lastUpdated,
	}
}

type metadataMatcher struct {
	expected    types.Metadata
	lastUpdated time.Time
	actual      types.Metadata
	comparison  comparison
}

func (m *metadataMatcher) Match(actual interface{}) (success bool, err error) {
	content, ok := actual.(string)
	if !ok {
		return false, errors.Errorf("HaveMetadata matcher expects a string (actual: %T)", actual)
	}
	metadata, err := libasciidoc.ConvertToHTML("", strings.NewReader(content), bytes.NewBuffer(nil), renderer.IncludeHeaderFooter(false), renderer.LastUpdated(m.lastUpdated))
	if err != nil {
		return false, err
	}
	m.actual = metadata
	m.comparison = compare(m.actual, m.expected)
	return m.comparison.diffs == "", nil
}

func (m *metadataMatcher) FailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected metadata to match:\n%s", m.comparison.diffs)
}

func (m *metadataMatcher) NegatedFailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected metadata not to match:\n%s", m.comparison.diffs)
}
