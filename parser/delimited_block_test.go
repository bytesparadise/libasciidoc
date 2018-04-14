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
				Kind: types.FencedBlock,
				Elements: []types.DocElement{
					types.StringElement{
						Content: content,
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("BlockElement"))
		})

		It("fenced block with no line", func() {
			content := ""
			actualContent := "```\n" + content + "```"
			expectedResult := types.DelimitedBlock{
				Kind: types.FencedBlock,
				Elements: []types.DocElement{
					types.StringElement{
						Content: content,
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("BlockElement"))
		})

		It("fenced block with multiple lines", func() {
			content := "some fenced code\nwith an empty line\n\nin the middle"
			actualContent := "```\n" + content + "\n```"
			expectedResult := types.DelimitedBlock{
				Kind: types.FencedBlock,
				Elements: []types.DocElement{
					types.StringElement{
						Content: content,
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("BlockElement"))
		})

		It("fenced block with multiple lines then a paragraph", func() {
			content := "some fenced code\nwith an empty line\n\nin the middle"
			actualContent := "```\n" + content + "\n```\nthen a normal paragraph."
			expectedResult := types.Document{
				Attributes:        map[string]interface{}{},
				ElementReferences: map[string]interface{}{},
				Elements: []types.DocElement{
					types.DelimitedBlock{
						Kind: types.FencedBlock,
						Elements: []types.DocElement{
							types.StringElement{
								Content: content,
							},
						},
					},
					types.Paragraph{
						Lines: []types.InlineContent{
							{
								Elements: []types.InlineElement{
									types.StringElement{Content: "then a normal paragraph."},
								},
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
				Elements: []types.DocElement{
					types.Paragraph{
						Lines: []types.InlineContent{
							{
								Elements: []types.InlineElement{
									types.StringElement{Content: "a paragraph."},
								},
							},
						},
					},
					types.DelimitedBlock{
						Kind: types.FencedBlock,
						Elements: []types.DocElement{
							types.StringElement{
								Content: content,
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

	Context("listing blocks", func() {

		It("listing block with single line", func() {
			actualContent := `----
some listing code
----`
			expectedResult := types.DelimitedBlock{
				Kind: types.ListingBlock,
				Elements: []types.DocElement{
					types.StringElement{
						Content: "some listing code",
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("BlockElement"))
		})

		It("listing block with no line", func() {
			content := ""
			actualContent := "----\n" + content + "----"
			expectedResult := types.DelimitedBlock{
				Kind: types.ListingBlock,
				Elements: []types.DocElement{
					types.StringElement{
						Content: "",
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("BlockElement"))
		})

		It("listing block with multiple lines", func() {
			content := "some listing code\nwith an empty line\n\nin the middle"
			actualContent := "----\n" + content + "\n----"
			expectedResult := types.DelimitedBlock{
				Kind: types.ListingBlock,
				Elements: []types.DocElement{
					types.StringElement{
						Content: content,
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("BlockElement"))
		})

		It("listing block with multiple lines then a paragraph", func() {
			content := "some listing code\nwith an empty line\n\nin the middle"
			actualContent := "----\n" + content + "\n----\nthen a normal paragraph."
			expectedResult := types.Document{
				Attributes:        map[string]interface{}{},
				ElementReferences: map[string]interface{}{},
				Elements: []types.DocElement{
					types.DelimitedBlock{
						Kind: types.ListingBlock,
						Elements: []types.DocElement{
							types.StringElement{
								Content: content,
							},
						},
					},
					types.Paragraph{
						Lines: []types.InlineContent{
							{
								Elements: []types.InlineElement{
									types.StringElement{Content: "then a normal paragraph."},
								},
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
				Elements: []types.DocElement{
					types.Paragraph{
						Lines: []types.InlineContent{
							{
								Elements: []types.InlineElement{
									types.StringElement{Content: "a paragraph."},
								},
							},
						},
					},
					types.DelimitedBlock{
						Kind: types.ListingBlock,
						Elements: []types.DocElement{
							types.StringElement{
								Content: "some listing code",
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

	Context("Literal blocks with spaces indentation", func() {

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
				Elements: []types.DocElement{
					types.LiteralBlock{
						Content: "   some literal content",
					},
					types.Paragraph{
						Lines: []types.InlineContent{
							{
								Elements: []types.InlineElement{
									types.StringElement{Content: "a normal paragraph."},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

	Context("Literal blocks with block delimiter", func() {

		It("literal block from 1-line paragraph with delimiter", func() {
			actualContent := `....
some literal content
....
a normal paragraph.`
			expectedResult := types.Document{
				Attributes:        map[string]interface{}{},
				ElementReferences: map[string]interface{}{},
				Elements: []types.DocElement{
					types.LiteralBlock{
						Content: "some literal content",
					},
					types.Paragraph{
						Lines: []types.InlineContent{
							{
								Elements: []types.InlineElement{
									types.StringElement{Content: "a normal paragraph."},
								},
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
				Elements: []types.DocElement{
					types.LiteralBlock{
						Content: "some literal content",
					},
					types.Paragraph{
						Lines: []types.InlineContent{
							{
								Elements: []types.InlineElement{
									types.StringElement{Content: "a normal paragraph."},
								},
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
				Kind: types.ExampleBlock,
				Elements: []types.DocElement{
					types.Paragraph{
						Lines: []types.InlineContent{
							{
								Elements: []types.InlineElement{
									types.StringElement{
										Content: "some listing code",
									},
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
				Kind: types.ExampleBlock,
				Elements: []types.DocElement{
					types.Paragraph{
						Lines: []types.InlineContent{
							{
								Elements: []types.InlineElement{
									types.StringElement{
										Content: "some listing code",
									},
								},
							},
							{
								Elements: []types.InlineElement{
									types.StringElement{
										Content: "with ",
									},
									types.QuotedText{
										Kind: types.Bold,
										Elements: []types.InlineElement{
											types.StringElement{
												Content: "bold content",
											},
										},
									},
								},
							},
						},
					},
					types.UnorderedList{
						Attributes: map[string]interface{}{},
						Items: []types.UnorderedListItem{
							{
								Level:       1,
								BulletStyle: types.OneAsterisk,
								Elements: []types.DocElement{
									types.ListParagraph{
										Lines: []types.InlineContent{
											{
												Elements: []types.InlineElement{
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
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("BlockElement"))
		})
	})

})
