package parser_test

import (
	"github.com/bytesparadise/libasciidoc/parser"
	"github.com/bytesparadise/libasciidoc/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("delimited blocks", func() {

	Context("fenced blocks", func() {

		It("fenced block with single line", func() {
			content := "some fenced code"
			actualContent := "```\n" + content + "\n```"
			expectedResult := types.DelimitedBlock{
				Kind:       types.FencedBlock,
				Attributes: map[string]interface{}{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: map[string]interface{}{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: content,
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("BlockElement"))
		})

		It("fenced block with no line", func() {
			actualContent := "```\n```"
			expectedResult := types.DelimitedBlock{
				Kind:       types.FencedBlock,
				Attributes: map[string]interface{}{},
				Elements:   []interface{}{},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("BlockElement"))
		})

		It("fenced block with multiple lines alone", func() {
			actualContent := "```\nsome fenced code\nwith an empty line\n\nin the middle\n```"
			expectedResult := types.DelimitedBlock{
				Kind:       types.FencedBlock,
				Attributes: map[string]interface{}{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: map[string]interface{}{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "some fenced code",
								},
							},
							{
								types.StringElement{
									Content: "with an empty line",
								},
							},
						},
					},
					types.BlankLine{},
					types.Paragraph{
						Attributes: map[string]interface{}{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "in the middle",
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("BlockElement"))
		})

		It("fenced block with multiple lines then a paragraph", func() {
			actualContent := "```\nsome fenced code\nwith an empty line\n\nin the middle\n```\nthen a normal paragraph."
			expectedResult := types.Document{
				Attributes:        map[string]interface{}{},
				ElementReferences: map[string]interface{}{},
				Elements: []interface{}{
					types.DelimitedBlock{
						Kind:       types.FencedBlock,
						Attributes: map[string]interface{}{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: map[string]interface{}{},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "some fenced code",
										},
									},
									{
										types.StringElement{
											Content: "with an empty line",
										},
									},
								},
							},
							types.BlankLine{},
							types.Paragraph{
								Attributes: map[string]interface{}{},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "in the middle",
										},
									},
								},
							},
						},
					},
					types.Paragraph{
						Attributes: map[string]interface{}{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "then a normal paragraph."},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("fenced block after a paragraph", func() {
			content := "some fenced code"
			actualContent := "a paragraph.\n```\n" + content + "\n```\n"
			expectedResult := types.Document{
				Attributes:        map[string]interface{}{},
				ElementReferences: map[string]interface{}{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: map[string]interface{}{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "a paragraph."},
							},
						},
					},
					types.DelimitedBlock{
						Kind:       types.FencedBlock,
						Attributes: map[string]interface{}{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: map[string]interface{}{},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: content,
										},
									},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("fenced block with unclosed delimiter", func() {
			actualContent := "```\nEnd of file here"
			expectedResult := types.DelimitedBlock{
				Kind:       types.FencedBlock,
				Attributes: map[string]interface{}{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: map[string]interface{}{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "End of file here",
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("BlockElement"))
		})
	})

	Context("listing blocks", func() {

		It("listing block with single line", func() {
			actualContent := `----
some listing code
----`
			expectedResult := types.DelimitedBlock{
				Kind:       types.ListingBlock,
				Attributes: map[string]interface{}{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: map[string]interface{}{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "some listing code",
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("BlockElement"))
		})

		It("listing block with no line", func() {
			actualContent := `----
----`
			expectedResult := types.DelimitedBlock{
				Kind:       types.ListingBlock,
				Attributes: map[string]interface{}{},
				Elements:   []interface{}{},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("BlockElement"))
		})

		It("listing block with multiple lines", func() {
			actualContent := `----
some listing code
with an empty line

in the middle
----`
			expectedResult := types.DelimitedBlock{
				Kind:       types.ListingBlock,
				Attributes: map[string]interface{}{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: map[string]interface{}{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "some listing code",
								},
							},
							{
								types.StringElement{
									Content: "with an empty line",
								},
							},
						},
					},
					types.BlankLine{},
					types.Paragraph{
						Attributes: map[string]interface{}{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "in the middle",
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("BlockElement"))
		})

		It("listing block with multiple lines then a paragraph", func() {
			actualContent := `---- 
some listing code
with an empty line

in the middle
----
then a normal paragraph.`
			expectedResult := types.Document{
				Attributes:        map[string]interface{}{},
				ElementReferences: map[string]interface{}{},
				Elements: []interface{}{
					types.DelimitedBlock{
						Kind:       types.ListingBlock,
						Attributes: map[string]interface{}{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: map[string]interface{}{},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "some listing code",
										},
									},
									{
										types.StringElement{
											Content: "with an empty line",
										},
									},
								},
							},
							types.BlankLine{},
							types.Paragraph{
								Attributes: map[string]interface{}{},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "in the middle",
										},
									},
								},
							},
						},
					},
					types.Paragraph{
						Attributes: map[string]interface{}{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "then a normal paragraph."},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("listing block after a paragraph", func() {
			actualContent := `a paragraph.
			
----
some listing code
----`
			expectedResult := types.Document{
				Attributes:        map[string]interface{}{},
				ElementReferences: map[string]interface{}{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: map[string]interface{}{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "a paragraph."},
							},
						},
					},
					types.DelimitedBlock{
						Kind:       types.ListingBlock,
						Attributes: map[string]interface{}{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: map[string]interface{}{},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "some listing code",
										},
									},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("listing block with unclosed delimiter", func() {
			actualContent := `----
End of file here.`
			expectedResult := types.DelimitedBlock{
				Kind:       types.ListingBlock,
				Attributes: map[string]interface{}{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: map[string]interface{}{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "End of file here.",
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("BlockElement"))
		})
	})

	Context("literal blocks with spaces indentation", func() {

		It("literal block from 1-line paragraph with single space", func() {
			actualContent := ` some literal content`
			expectedResult := types.LiteralBlock{
				Content: " some literal content",
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("BlockElement"))
		})

		It("literal block from paragraph with single space on first line", func() {
			actualContent := ` some literal content
on 2 lines.`
			expectedResult := types.LiteralBlock{
				Content: " some literal content\non 2 lines.",
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("BlockElement"))
		})

		It("mixing literal block and paragraph ", func() {
			actualContent := `   some literal content

a normal paragraph.`
			expectedResult := types.Document{
				Attributes:        map[string]interface{}{},
				ElementReferences: map[string]interface{}{},
				Elements: []interface{}{
					types.LiteralBlock{
						Content: "   some literal content",
					},
					types.Paragraph{
						Attributes: map[string]interface{}{},
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

		It("literal block from 1-line paragraph with delimiter", func() {
			actualContent := `....
some literal content
....
a normal paragraph.`
			expectedResult := types.Document{
				Attributes:        map[string]interface{}{},
				ElementReferences: map[string]interface{}{},
				Elements: []interface{}{
					types.LiteralBlock{
						Content: "some literal content",
					},
					types.Paragraph{
						Attributes: map[string]interface{}{},
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
				Attributes:        map[string]interface{}{},
				ElementReferences: map[string]interface{}{},
				Elements: []interface{}{
					types.LiteralBlock{
						Content: "some literal content",
					},
					types.Paragraph{
						Attributes: map[string]interface{}{},
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

	Context("example blocks", func() {

		It("example block with single line", func() {
			actualContent := `====
some listing code
====`
			expectedResult := types.DelimitedBlock{
				Kind:       types.ExampleBlock,
				Attributes: map[string]interface{}{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: map[string]interface{}{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "some listing code",
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("BlockElement"))
		})

		It("example block with multiple lines", func() {
			actualContent := `====
some listing code
with *bold content*

* and a list item
====`
			expectedResult := types.DelimitedBlock{
				Kind:       types.ExampleBlock,
				Attributes: map[string]interface{}{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: map[string]interface{}{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "some listing code",
								},
							},
							{
								types.StringElement{
									Content: "with ",
								},
								types.QuotedText{
									Kind: types.Bold,
									Elements: []interface{}{
										types.StringElement{
											Content: "bold content",
										},
									},
								},
							},
						},
					},
					types.BlankLine{},
					types.UnorderedList{
						Attributes: map[string]interface{}{},
						Items: []types.UnorderedListItem{
							{
								Level:       1,
								BulletStyle: types.OneAsterisk,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: map[string]interface{}{},
										Lines: []types.InlineElements{
											{
												types.StringElement{
													Content: "and a list item",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("BlockElement"))
		})

		It("example block with unclosed delimiter", func() {
			actualContent := `====
End of file here`
			expectedResult := types.DelimitedBlock{
				Kind:       types.ExampleBlock,
				Attributes: map[string]interface{}{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: map[string]interface{}{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "End of file here",
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("BlockElement"))
		})
	})

	Context("admonition blocks", func() {

		It("example block as admonition", func() {
			actualContent := `[NOTE]
====
foo
====`
			expectedResult := types.DelimitedBlock{
				Kind: types.ExampleBlock,
				Attributes: map[string]interface{}{
					types.AttrAdmonitionKind: types.Note,
				},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: map[string]interface{}{},
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
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("BlockElement"))

		})

		It("listing block as admonition", func() {
			actualContent := `[NOTE]
----
multiple

paragraphs
----
`
			expectedResult := types.Document{
				Attributes:        map[string]interface{}{},
				ElementReferences: map[string]interface{}{},
				Elements: []interface{}{
					types.DelimitedBlock{
						Kind: types.ListingBlock,
						Attributes: map[string]interface{}{
							types.AttrAdmonitionKind: types.Note,
						},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: map[string]interface{}{},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "multiple",
										},
									},
								},
							},
							types.BlankLine{},
							types.Paragraph{
								Attributes: map[string]interface{}{},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "paragraphs",
										},
									},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Document"))
		})

	})

})
