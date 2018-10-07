package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("literal blocks", func() {

	Context("literal blocks with spaces indentation", func() {

		It("literal block from 1-line paragraph with single space", func() {
			actualContent := ` some literal content`
			expectedResult := types.LiteralBlock{
				Attributes: types.ElementAttributes{
					types.AttrKind:             types.Literal,
					types.AttrLiteralBlockType: types.LiteralBlockWithSpacesOnFirstLine,
				},
				Lines: []string{
					" some literal content",
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("literal block from paragraph with single space on first line", func() {
			actualContent := ` some literal content
on 3
lines.`
			expectedResult := types.LiteralBlock{
				Attributes: types.ElementAttributes{
					types.AttrKind:             types.Literal,
					types.AttrLiteralBlockType: types.LiteralBlockWithSpacesOnFirstLine,
				},
				Lines: []string{
					" some literal content",
					"on 3",
					"lines.",
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("mixing literal block with attributes followed by a paragraph ", func() {
			actualContent := `.title
[#ID]
  some literal content

a normal paragraph.`
			expectedResult := types.Document{
				Attributes:         map[string]interface{}{},
				ElementReferences:  map[string]interface{}{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.LiteralBlock{
						Attributes: types.ElementAttributes{
							types.AttrKind:             types.Literal,
							types.AttrLiteralBlockType: types.LiteralBlockWithSpacesOnFirstLine,
							types.AttrID:               "ID",
							types.AttrTitle:            "title",
						},
						Lines: []string{
							"  some literal content",
						},
					},
					types.BlankLine{},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "a normal paragraph."},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

	Context("literal blocks with block delimiter", func() {

		It("literal block with empty blank line", func() {

			actualContent := `....

some content
....`
			expectedResult := types.LiteralBlock{
				Attributes: types.ElementAttributes{
					types.AttrKind:             types.Literal,
					types.AttrLiteralBlockType: types.LiteralBlockWithDelimiter,
				},
				Lines: []string{
					"",
					"some content",
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("literal block with delimited and attributes followed by 1-line paragraph", func() {
			actualContent := `[#ID]
.title
....
some literal content
....
a normal paragraph.`
			expectedResult := types.Document{
				Attributes:         map[string]interface{}{},
				ElementReferences:  map[string]interface{}{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.LiteralBlock{
						Attributes: types.ElementAttributes{
							types.AttrKind:             types.Literal,
							types.AttrLiteralBlockType: types.LiteralBlockWithDelimiter,
							types.AttrID:               "ID",
							types.AttrTitle:            "title",
						},
						Lines: []string{
							"some literal content",
						},
					},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "a normal paragraph."},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})

	})

	Context("literal blocks with attribute", func() {

		It("literal block from 1-line paragraph with attribute", func() {
			actualContent := `[literal]   
some literal content

a normal paragraph.`
			expectedResult := types.Document{
				Attributes:         map[string]interface{}{},
				ElementReferences:  map[string]interface{}{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.LiteralBlock{
						Attributes: types.ElementAttributes{
							types.AttrKind:             types.Literal,
							types.AttrLiteralBlockType: types.LiteralBlockWithAttribute,
						},
						Lines: []string{
							"some literal content",
						},
					},
					types.BlankLine{},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "a normal paragraph."},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("literal block from 2-lines paragraph with attribute", func() {
			actualContent := `[#ID]
[literal]   
.title
some literal content
on two lines.

a normal paragraph.`
			expectedResult := types.Document{
				Attributes:         map[string]interface{}{},
				ElementReferences:  map[string]interface{}{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.LiteralBlock{
						Attributes: types.ElementAttributes{
							types.AttrKind:             types.Literal,
							types.AttrID:               "ID",
							types.AttrTitle:            "title",
							types.AttrLiteralBlockType: types.LiteralBlockWithAttribute,
						},
						Lines: []string{
							"some literal content",
							"on two lines.",
						},
					},
					types.BlankLine{},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "a normal paragraph."},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

})
