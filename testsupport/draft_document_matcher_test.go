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

var _ = Describe("draft document matcher", func() {

	// given
	expected := types.DraftDocument{
		Blocks: []interface{}{
			types.Paragraph{
				Lines: [][]interface{}{
					{
						types.StringElement{
							Content: "a paragraph.",
						},
					},
				},
			},
		},
	}
	matcher := testsupport.MatchDraftDocument(expected)

	It("should match", func() {
		// given
		actual := types.DraftDocument{
			Blocks: []interface{}{
				types.Paragraph{
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "a paragraph.",
							},
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
		actual := types.DraftDocument{
			Blocks: []interface{}{
				types.Paragraph{
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "another paragraph.",
							},
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
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(spew.Sdump(actual), spew.Sdump(expected), true)
		Expect(matcher.FailureMessage(actual)).To(Equal(fmt.Sprintf("expected draft documents to match:\n%s", dmp.DiffPrettyText(diffs))))
		Expect(matcher.NegatedFailureMessage(actual)).To(Equal(fmt.Sprintf("expected draft documents not to match:\n%s", dmp.DiffPrettyText(diffs))))
	})

	It("should return error when invalid type is input", func() {
		// when
		result, err := matcher.Match(1)
		// then
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("MatchDraftDocument matcher expects a DraftDocument (actual: int)"))
		Expect(result).To(BeFalse())
	})

})
