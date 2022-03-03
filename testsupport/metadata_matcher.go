package testsupport

import (
	"fmt"
	"reflect"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/google/go-cmp/cmp"

	"github.com/davecgh/go-spew/spew"
	. "github.com/onsi/ginkgo" // nolint go-lint
	gomegatypes "github.com/onsi/gomega/types"
	"github.com/pkg/errors"
)

func MatchMetadata(expected types.Metadata) gomegatypes.GomegaMatcher {
	return &metadataMatcher{
		expected: expected,
	}
}

type metadataMatcher struct {
	expected types.Metadata
	diffs    string
}

func (m *metadataMatcher) Match(actual interface{}) (success bool, err error) {
	if _, ok := actual.(types.Metadata); !ok {
		return false, errors.Errorf("MatchMetadata matcher expects a 'types.Metadata' (actual: %T)", actual)
	}
	if !reflect.DeepEqual(m.expected, actual) {
		GinkgoT().Logf("actual HTML:\n'%s'", actual)
		GinkgoT().Logf("expected HTML:\n'%s'", m.expected)
		m.diffs = cmp.Diff(spew.Sdump(actual), spew.Sdump(m.expected))
		return false, nil
	}
	return true, nil
}

func (m *metadataMatcher) FailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected document metadata to match:\n%s", m.diffs)
}

func (m *metadataMatcher) NegatedFailureMessage(_ interface{}) (message string) {
	return fmt.Sprintf("expected document metadata not to match:\n%s", m.diffs)
}
