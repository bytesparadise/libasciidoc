package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("inline elements", func() {

	It("bold text without parenthesis", func() {
		source := "*some bold content*"
		expected := types.Paragraph{
			Attributes: types.ElementAttributes{},
			Lines: []types.InlineElements{
				{
					types.QuotedText{
						Kind: types.Bold,
						Elements: types.InlineElements{
							types.StringElement{Content: "some bold content"},
						},
					},
				},
			},
		}
		Expect(source).To(EqualDocumentBlock(expected))
	})

	It("bold text within parenthesis", func() {
		source := "(*some bold content*)"
		expected := types.Paragraph{
			Attributes: types.ElementAttributes{},
			Lines: []types.InlineElements{
				{
					types.StringElement{Content: "("},
					types.QuotedText{
						Kind: types.Bold,
						Elements: types.InlineElements{
							types.StringElement{Content: "some bold content"},
						},
					},
					types.StringElement{Content: ")"},
				},
			},
		}
		Expect(source).To(EqualDocumentBlock(expected))
	})

	It("bold text within words", func() {
		source := "some*bold*content"
		expected := types.Paragraph{
			Attributes: types.ElementAttributes{},
			Lines: []types.InlineElements{
				{
					types.StringElement{Content: "some*bold*content"},
				},
			},
		}
		Expect(source).To(EqualDocumentBlock(expected))
	})

	It("invalid bold portion of text", func() {
		source := "*foo*bar"
		expected := types.Paragraph{
			Attributes: types.ElementAttributes{},
			Lines: []types.InlineElements{
				{
					types.StringElement{Content: "*foo*bar"},
				},
			},
		}
		Expect(source).To(EqualDocumentBlock(expected))
	})

	It("valid bold portion of text", func() {
		source := "**foo**bar"
		expected := types.Paragraph{
			Attributes: types.ElementAttributes{},
			Lines: []types.InlineElements{
				{
					types.QuotedText{
						Kind: types.Bold,
						Elements: types.InlineElements{
							types.StringElement{Content: "foo"},
						},
					},
					types.StringElement{Content: "bar"},
				},
			},
		}
		Expect(source).To(EqualDocumentBlock(expected))
	})

	It("latin characters", func() {
		source := "à bientôt"
		expected := types.Paragraph{
			Attributes: types.ElementAttributes{},
			Lines: []types.InlineElements{
				{
					types.StringElement{Content: "à bientôt"},
				},
			},
		}
		Expect(source).To(EqualDocumentBlock(expected))
	})
})
