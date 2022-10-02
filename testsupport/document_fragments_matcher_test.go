package testsupport_test

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/bytesparadise/libasciidoc/testsupport"
	"github.com/google/go-cmp/cmp"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("document fragment matcher", func() {

	// given
	expected := []types.DocumentFragment{
		{
			Elements: []interface{}{
				&types.RawLine{
					Content: "a paragraph.",
				},
			},
		},
	}
	matcher := testsupport.MatchDocumentFragments(expected)

	It("should match", func() {
		// given
		actual := []types.DocumentFragment{
			{
				Elements: []interface{}{
					&types.RawLine{
						Content: "a paragraph.",
					},
				},
			},
		}
		// when
		result, err := matcher.Match(actual)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(BeTrue())
	})

	It("should not match", func() {
		// given
		actual := []types.DocumentFragment{
			{
				Elements: []interface{}{
					&types.RawLine{
						Content: "something else",
					},
				},
			},
		}
		// when
		result, err := matcher.Match(actual)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(BeFalse())
		diffs := cmp.Diff(expected, actual)
		Expect(matcher.FailureMessage(actual)).To(Equal(fmt.Sprintf("expected document fragments to match:\n%s", diffs)))
		Expect(matcher.NegatedFailureMessage(actual)).To(Equal(fmt.Sprintf("expected document fragments not to match:\n%s", diffs)))
	})

	It("should return error when invalid type is input", func() {
		// when
		result, err := matcher.Match(1)
		// then
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("MatchDocumentFragments matcher expects a '[]types.DocumentFragment' (actual: int)"))
		Expect(result).To(BeFalse())
	})

})
