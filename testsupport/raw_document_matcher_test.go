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

var _ = Describe("raw document matcher", func() {

	// given
	expected := types.RawDocument{
		Blocks: []interface{}{
			types.Paragraph{
				Lines: []interface{}{
					types.RawLine{
						Content: "a paragraph.",
					},
				},
			},
		},
	}
	matcher := testsupport.MatchRawDocument(expected)

	It("should match", func() {
		// given
		actual := types.RawDocument{
			Blocks: []interface{}{
				types.Paragraph{
					Lines: []interface{}{
						types.RawLine{
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
		actual := types.RawDocument{
			Blocks: []interface{}{
				types.Paragraph{
					Lines: []interface{}{
						types.RawLine{
							Content: "another paragraph.", // different content
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
		Expect(matcher.FailureMessage(actual)).To(Equal(fmt.Sprintf("expected raw documents to match:\n%s", dmp.DiffPrettyText(diffs))))
		Expect(matcher.NegatedFailureMessage(actual)).To(Equal(fmt.Sprintf("expected raw documents not to match:\n%s", dmp.DiffPrettyText(diffs))))
	})

	It("should return error when invalid type is input", func() {
		// when
		result, err := matcher.Match(1)
		// then
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("MatchRawDocument matcher expects a RawDocument (actual: int)"))
		Expect(result).To(BeFalse())
	})

})
