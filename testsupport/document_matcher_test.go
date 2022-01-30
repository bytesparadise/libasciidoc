package testsupport_test

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/bytesparadise/libasciidoc/testsupport"
	"github.com/google/go-cmp/cmp"

	. "github.com/onsi/ginkgo" // nolint:golint
	. "github.com/onsi/gomega" // nolint:golintt
)

var _ = Describe("document matcher", func() {

	// given
	expected := &types.Document{
		Elements: []interface{}{
			&types.Paragraph{
				Elements: []interface{}{
					&types.StringElement{
						Content: "a paragraph.",
					},
				},
			},
		},
	}
	matcher := testsupport.MatchDocument(expected)

	It("should match", func() {
		// given
		actual := &types.Document{
			Elements: []interface{}{
				&types.Paragraph{
					Elements: []interface{}{
						&types.StringElement{
							Content: "a paragraph.",
						},
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
		actual := &types.Document{
			Elements: []interface{}{
				&types.Paragraph{
					Elements: []interface{}{
						&types.StringElement{
							Content: "another paragraph.",
						},
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
		Expect(matcher.FailureMessage(actual)).To(Equal(fmt.Sprintf("expected documents to match:\n%s", diffs)))
		Expect(matcher.NegatedFailureMessage(actual)).To(Equal(fmt.Sprintf("expected documents not to match:\n%s", diffs)))
	})

	It("should return error when invalid type is input", func() {
		// when
		result, err := matcher.Match(1)
		// then
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("MatchDocument matcher expects a Document (actual: int)"))
		Expect(result).To(BeFalse())
	})

})
