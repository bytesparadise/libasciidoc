package testsupport_test

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("document block assertions", func() {

	expected := types.Paragraph{
		Attributes: types.ElementAttributes{},
		Lines: []types.InlineElements{
			{
				types.StringElement{
					Content: "hello, world!",
				},
			},
		},
	}

	It("should match", func() {
		// given
		matcher := testsupport.EqualDocumentBlock(expected)
		actual := "hello, world!"
		// when
		result, err := matcher.Match(actual)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(BeTrue())
	})

	It("should not match", func() {
		// given
		matcher := testsupport.EqualDocumentBlock(expected)
		actual := "foo"
		// when
		result, err := matcher.Match(actual)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(BeFalse())
		// also verify the messages
		obtained := types.Paragraph{
			Attributes: types.ElementAttributes{},
			Lines: []types.InlineElements{
				{
					types.StringElement{
						Content: "foo",
					},
				},
			},
		}
		Expect(matcher.FailureMessage(actual)).To(Equal(fmt.Sprintf("expected document blocks to match:\n%s", compare(obtained, expected))))
		Expect(matcher.NegatedFailureMessage(actual)).To(Equal(fmt.Sprintf("expected document blocks not to match:\n%s", compare(obtained, expected))))
	})

	It("should return error when invalid type is input", func() {
		// given
		matcher := testsupport.EqualDocumentBlock(expected)
		// when
		result, err := matcher.Match(1) // not a string
		// then
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("EqualDocumentBlock matcher expects a string (actual: int)"))
		Expect(result).To(BeFalse())
	})
})
