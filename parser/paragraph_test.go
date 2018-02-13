package parser_test

import (
	"github.com/bytesparadise/libasciidoc/parser"
	"github.com/bytesparadise/libasciidoc/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Paragraphs", func() {

	It("paragraph with 1 word", func() {
		actualContent := "hello"
		expectedDocument := &types.Paragraph{
			Lines: []*types.InlineContent{
				{
					Elements: []types.InlineElement{
						&types.StringElement{Content: "hello"},
					},
				},
			},
		}
		verify(GinkgoT(), expectedDocument, actualContent, parser.Entrypoint("Paragraph"))
	})

	It("paragraph with few words", func() {
		actualContent := "a paragraph with some content"
		expectedDocument := &types.Paragraph{
			Lines: []*types.InlineContent{
				{
					Elements: []types.InlineElement{
						&types.StringElement{Content: "a paragraph with some content"},
					},
				},
			},
		}
		verify(GinkgoT(), expectedDocument, actualContent, parser.Entrypoint("Paragraph"))
	})

	It("paragraph with bold content", func() {
		actualContent := "a paragraph with *some bold content*"
		expectedDocument := &types.Paragraph{
			Lines: []*types.InlineContent{
				{
					Elements: []types.InlineElement{
						&types.StringElement{Content: "a paragraph with "},
						&types.QuotedText{
							Kind: types.Bold,
							Elements: []types.InlineElement{
								&types.StringElement{Content: "some bold content"},
							},
						},
					},
				},
			},
		}
		verify(GinkgoT(), expectedDocument, actualContent, parser.Entrypoint("Paragraph"))
	})

	It("paragraph with id and title", func() {
		actualContent := `[#foo]
.a title
a paragraph`
		expectedDocument := &types.Paragraph{
			ID:    &types.ElementID{Value: "foo"},
			Title: &types.ElementTitle{Value: "a title"},
			Lines: []*types.InlineContent{
				{
					Elements: []types.InlineElement{
						&types.StringElement{Content: "a paragraph"},
					},
				},
			},
		}
		verify(GinkgoT(), expectedDocument, actualContent, parser.Entrypoint("Paragraph"))
	})
})
