package testsupport_test

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/bytesparadise/libasciidoc/testsupport"

	"github.com/davecgh/go-spew/spew"
	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
	"github.com/sergi/go-diff/diffmatchpatch"
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
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(spew.Sdump(actual), spew.Sdump(expected), true)
		Expect(matcher.FailureMessage(actual)).To(Equal(fmt.Sprintf("expected elements to match:\n%s", dmp.DiffPrettyText(diffs))))
		Expect(matcher.NegatedFailureMessage(actual)).To(Equal(fmt.Sprintf("expected elements not to match:\n%s", dmp.DiffPrettyText(diffs))))
	})

	It("should return error when invalid type is input", func() {
		// when
		result, err := matcher.Match(1)
		// then
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("MatchInlineElements matcher expects a []interface{} (actual: int)"))
		Expect(result).To(BeFalse())
	})

})
