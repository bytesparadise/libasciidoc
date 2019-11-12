package testsupport_test

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("document assertions", func() {

	expected := types.Document{
		Attributes:         types.DocumentAttributes{},
		ElementReferences:  types.ElementReferences{},
		Footnotes:          types.Footnotes{},
		FootnoteReferences: types.FootnoteReferences{},
		Elements: []interface{}{
			types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{
							Content: "hello, world!",
						},
					},
				},
			},
		},
	}

	It("should match", func() {
		// given
		actual := "hello, world!"
		matcher := testsupport.BecomeDocument(expected)
		// when
		result, err := matcher.Match(actual)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(BeTrue())
	})

	It("should not match", func() {
		// given
		actual := "foo"
		matcher := testsupport.BecomeDocument(expected)
		// when
		result, err := matcher.Match(actual)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(BeFalse())
		// also verify the messages
		obtained := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{
								Content: "foo",
							},
						},
					},
				},
			},
		}
		Expect(matcher.FailureMessage(actual)).To(Equal(fmt.Sprintf("expected documents to match:\n%s", compare(obtained, expected))))
		Expect(matcher.NegatedFailureMessage(actual)).To(Equal(fmt.Sprintf("expected documents not to match:\n%s", compare(obtained, expected))))
	})

	It("should return error when invalid type is input", func() {
		// given
		actual := "foo"
		matcher := testsupport.BecomeDocument(types.Document{})
		_, err := matcher.Match(actual)
		Expect(err).ToNot(HaveOccurred())
		// when
		result, err := matcher.Match(1) // not a string
		// then
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("BecomeDocument matcher expects a string (actual: int)"))
		Expect(result).To(BeFalse())
	})
})
