package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
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
		verifyDocumentBlock(expected, source)
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
		verifyDocumentBlock(expected, source)
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
		verifyDocumentBlock(expected, source)
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
		verifyDocumentBlock(expected, source)
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
		verifyDocumentBlock(expected, source)
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
		verifyDocumentBlock(expected, source)
	})
})
