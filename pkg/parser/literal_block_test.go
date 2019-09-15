package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("literal blocks - preflight", func() {

	Context("literal blocks with spaces indentation", func() {

		It("literal block from 1-line paragraph with single space", func() {
			source := ` some literal content`
			expected := types.LiteralBlock{
				Attributes: types.ElementAttributes{
					types.AttrKind:             types.Literal,
					types.AttrLiteralBlockType: types.LiteralBlockWithSpacesOnFirstLine,
				},
				Lines: []string{
					" some literal content",
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("literal block from paragraph with single space on first line", func() {
			source := ` some literal content
on 3
lines.`
			expected := types.LiteralBlock{
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
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("mixing literal block with attributes followed by a paragraph ", func() {
			source := `.title
[#ID]
  some literal content

a normal paragraph.`
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.LiteralBlock{
						Attributes: types.ElementAttributes{
							types.AttrKind:             types.Literal,
							types.AttrLiteralBlockType: types.LiteralBlockWithSpacesOnFirstLine,
							types.AttrID:               "ID",
							types.AttrCustomID:         true,
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
			Expect(source).To(EqualPreflightDocument(expected))
		})
	})

	Context("literal blocks with block delimiter", func() {

		It("literal block with empty blank line", func() {

			source := `....

some content
....`
			expected := types.LiteralBlock{
				Attributes: types.ElementAttributes{
					types.AttrKind:             types.Literal,
					types.AttrLiteralBlockType: types.LiteralBlockWithDelimiter,
				},
				Lines: []string{
					"",
					"some content",
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("literal block with delimited and attributes followed by 1-line paragraph", func() {
			source := `[#ID]
.title
....
some literal content
....
a normal paragraph.`
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.LiteralBlock{
						Attributes: types.ElementAttributes{
							types.AttrKind:             types.Literal,
							types.AttrLiteralBlockType: types.LiteralBlockWithDelimiter,
							types.AttrID:               "ID",
							types.AttrCustomID:         true,
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
			Expect(source).To(EqualPreflightDocument(expected))
		})
	})

	Context("literal blocks with attribute", func() {

		It("literal block from 1-line paragraph with attribute", func() {
			source := `[literal]   
some literal content

a normal paragraph.`
			expected := types.PreflightDocument{
				Blocks: []interface{}{
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
			Expect(source).To(EqualPreflightDocument(expected))
		})

		It("literal block from 2-lines paragraph with attribute", func() {
			source := `[#ID]
[literal]   
.title
some literal content
on two lines.

a normal paragraph.`
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.LiteralBlock{
						Attributes: types.ElementAttributes{
							types.AttrKind:             types.Literal,
							types.AttrID:               "ID",
							types.AttrCustomID:         true,
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
			Expect(source).To(EqualPreflightDocument(expected))
		})
	})

})
