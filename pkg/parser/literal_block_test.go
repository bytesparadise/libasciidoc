package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("literal blocks", func() {

	Context("draft document", func() {

		Context("literal blocks with spaces indentation", func() {

			It("literal block from 1-line paragraph with single space", func() {
				source := ` some literal content`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.LiteralBlock{
							Attributes: types.Attributes{
								types.AttrKind:             types.Literal,
								types.AttrLiteralBlockType: types.LiteralBlockWithSpacesOnFirstLine,
							},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: " some literal content",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("literal block from paragraph with single space on first line", func() {
				source := ` some literal content
on 3
lines.`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.LiteralBlock{
							Attributes: types.Attributes{
								types.AttrKind:             types.Literal,
								types.AttrLiteralBlockType: types.LiteralBlockWithSpacesOnFirstLine,
							},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: " some literal content",
									},
								},
								{
									types.StringElement{
										Content: "on 3",
									},
								},
								{
									types.StringElement{
										Content: "lines.",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("mixing literal block with attributes followed by a paragraph ", func() {
				source := `.title
[#ID]
  some literal content

a normal paragraph.`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.LiteralBlock{
							Attributes: types.Attributes{
								types.AttrKind:             types.Literal,
								types.AttrLiteralBlockType: types.LiteralBlockWithSpacesOnFirstLine,
								types.AttrID:               "ID",
								types.AttrCustomID:         true,
								types.AttrTitle:            "title",
							},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "  some literal content",
									},
								},
							},
						},
						types.BlankLine{},
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a normal paragraph."},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})
		})

		Context("literal blocks with block delimiter", func() {

			It("literal block with empty blank line", func() {

				source := `....

some content
....`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.LiteralBlock{
							Attributes: types.Attributes{
								types.AttrKind:             types.Literal,
								types.AttrLiteralBlockType: types.LiteralBlockWithDelimiter,
							},
							Lines: [][]interface{}{
								{},
								{
									types.StringElement{
										Content: "some content",
									},
								},
							},
						},
					},
				}
				result, err := ParseDraftDocument(source)
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(MatchDraftDocument(expected))
			})

			It("literal block with delimited and attributes followed by 1-line paragraph", func() {
				source := `[#ID]
.title
....
some literal content
....
a normal paragraph.`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.LiteralBlock{
							Attributes: types.Attributes{
								types.AttrKind:             types.Literal,
								types.AttrLiteralBlockType: types.LiteralBlockWithDelimiter,
								types.AttrID:               "ID",
								types.AttrCustomID:         true,
								types.AttrTitle:            "title",
							},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "some literal content",
									},
								},
							},
						},
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a normal paragraph."},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})
		})

		Context("literal blocks with attribute", func() {

			It("literal block from 1-line paragraph with attribute", func() {
				source := `[literal]   
some literal content

a normal paragraph.`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.LiteralBlock{
							Attributes: types.Attributes{
								types.AttrKind:             types.Literal,
								types.AttrLiteralBlockType: types.LiteralBlockWithAttribute,
							},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "some literal content",
									},
								},
							},
						},
						types.BlankLine{},
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a normal paragraph."},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("literal block from 2-lines paragraph with attribute", func() {
				source := `[#ID]
[literal]   
.title
some literal content
on two lines.

a normal paragraph.`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.LiteralBlock{
							Attributes: types.Attributes{
								types.AttrKind:             types.Literal,
								types.AttrID:               "ID",
								types.AttrCustomID:         true,
								types.AttrTitle:            "title",
								types.AttrLiteralBlockType: types.LiteralBlockWithAttribute,
							},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "some literal content",
									},
								},
								{
									types.StringElement{
										Content: "on two lines.",
									},
								},
							},
						},
						types.BlankLine{},
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a normal paragraph."},
								},
							},
						},
					},
				}
				result, err := ParseDraftDocument(source)
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(MatchDraftDocument(expected))
			})
		})
	})
})
