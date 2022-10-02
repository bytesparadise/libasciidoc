package testsupport_test

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/bytesparadise/libasciidoc/testsupport"
	"github.com/google/go-cmp/cmp"

	"github.com/davecgh/go-spew/spew"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("inline elements matcher", func() {

	// given
	expected := []interface{}{
		&types.StringElement{
			Content: "a paragraph.",
		},
	}
	matcher := testsupport.MatchInlineElements(expected)

	It("should match", func() {
		// given
		actual := []interface{}{
			&types.StringElement{
				Content: "a paragraph.",
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
		actual := []interface{}{
			&types.StringElement{
				Content: "another paragraph.",
			},
		}
		// when
		result, err := matcher.Match(actual)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(BeFalse())
		diffs := cmp.Diff(spew.Sdump(expected), spew.Sdump(actual))
		Expect(matcher.FailureMessage(actual)).To(Equal(fmt.Sprintf("expected elements to match:\n%s", diffs)))
		Expect(matcher.NegatedFailureMessage(actual)).To(Equal(fmt.Sprintf("expected elements not to match:\n%s", diffs)))
	})

	It("should return error when invalid type is input", func() {
		// when
		result, err := matcher.Match(1)
		// then
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("MatchInlineElements matcher expects a '[]interface{}' (actual: int)"))
		Expect(result).To(BeFalse())
	})

})
