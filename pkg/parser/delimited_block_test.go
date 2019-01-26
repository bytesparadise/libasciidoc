package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("delimited blocks", func() {

	Context("fenced blocks", func() {

		It("fenced block with single line", func() {
			content := "some fenced code"
			actualContent := "```\n" + content + "\n```"
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{},
				Kind:       types.Fenced,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
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
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("fenced block with no line", func() {
			actualContent := "```\n```"
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{},
				Kind:       types.Fenced,
				Elements:   []interface{}{},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("fenced block with multiple lines alone", func() {
			actualContent := "```\nsome fenced code\nwith an empty line\n\nin the middle\n```"
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{},
				Kind:       types.Fenced,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
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
						Attributes: types.ElementAttributes{},
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
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("fenced block with multiple lines then a paragraph", func() {
			actualContent := "```\nsome fenced code\nwith an empty line\n\nin the middle\n```\nthen a normal paragraph."
			expectedResult := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						Kind:       types.Fenced,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
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
								Attributes: types.ElementAttributes{},
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
						Attributes: types.ElementAttributes{},
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
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "a paragraph."},
							},
						},
					},
					types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						Kind:       types.Fenced,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
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
				Attributes: types.ElementAttributes{},
				Kind:       types.Fenced,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
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
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})
	})

	Context("listing blocks", func() {

		It("listing block with single line", func() {
			actualContent := `----
some listing code
----`
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{},
				Kind:       types.Listing,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
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
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("listing block with no line", func() {
			actualContent := `----
----`
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{},
				Kind:       types.Listing,
				Elements:   []interface{}{},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("listing block with multiple lines alone", func() {
			actualContent := `----
some listing code
with an empty line

in the middle
----`
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{},
				Kind:       types.Listing,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
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
							{},
							{
								types.StringElement{
									Content: "in the middle",
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("listing block with unrendered list", func() {
			actualContent := `----
* some 
* listing 
* content
----`
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{},
				Kind:       types.Listing,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "* some ",
								},
							},
							{
								types.StringElement{
									Content: "* listing ",
								},
							},
							{
								types.StringElement{
									Content: "* content",
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("listing block with multiple lines then a paragraph", func() {
			actualContent := `---- 
some listing code
with an empty line

in the middle
----
then a normal paragraph.`
			expectedResult := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						Kind:       types.Listing,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
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
									{},
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
						Attributes: types.ElementAttributes{},
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

		It("listing block just after a paragraph", func() {
			actualContent := `a paragraph.
----
some listing code
----`
			expectedResult := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "a paragraph."},
							},
						},
					},
					types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						Kind:       types.Listing,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
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
				Attributes: types.ElementAttributes{},
				Kind:       types.Listing,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
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
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})
	})

	Context("example blocks", func() {

		It("example block with single line", func() {
			actualContent := `====
some listing code
====`
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{},
				Kind:       types.Example,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
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
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("example block with single line starting with a dot", func() {
			actualContent := `====
.foo
====`
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{},
				Kind:       types.Example,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: ".foo",
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("example block with multiple lines", func() {
			actualContent := `====
.foo
some listing code
with *bold content*

* and a list item
====`
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{},
				Kind:       types.Example,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: ".foo",
								},
							},
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
									Elements: types.InlineElements{
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
						Attributes: types.ElementAttributes{},
						Items: []types.UnorderedListItem{
							{
								Attributes:  types.ElementAttributes{},
								Level:       1,
								BulletStyle: types.OneAsterisk,
								CheckStyle:  types.NoCheck,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
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
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("example block with unclosed delimiter", func() {
			actualContent := `====
End of file here`
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{},
				Kind:       types.Example,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
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
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("example block with title", func() {
			actualContent := `.example block title
====
foo
====`
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{
					types.AttrTitle: "example block title",
				},
				Kind: types.Example,
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
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("example block starting delimiter only", func() {
			actualContent := `====`
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{},
				Kind:       types.Example,
				Elements:   []interface{}{},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})
	})

	Context("admonition blocks", func() {

		It("example block as admonition", func() {
			actualContent := `[NOTE]
====
foo
====`
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{
					types.AttrAdmonitionKind: types.Note,
				},
				Kind: types.Example,
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
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))

		})

		It("listing block as admonition", func() {
			actualContent := `[NOTE]
----
multiple

paragraphs
----
`
			expectedResult := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.DelimitedBlock{
						Attributes: types.ElementAttributes{
							types.AttrAdmonitionKind: types.Note,
						},
						Kind: types.Listing,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "multiple",
										},
									},
									{},
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

	Context("quote blocks", func() {

		It("single-line quote block with author and title", func() {
			actualContent := `[quote, john doe, quote title]
____
some *quote* content
____`
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Quote,
					types.AttrQuoteAuthor: "john doe",
					types.AttrQuoteTitle:  "quote title",
				},
				Kind: types.Quote,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "some ",
								},
								types.QuotedText{
									Kind: types.Bold,
									Elements: types.InlineElements{
										types.StringElement{
											Content: "quote",
										},
									},
								},
								types.StringElement{
									Content: " content",
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("multi-line quote with author only", func() {
			actualContent := `[quote, john doe,   ]
____
- some 
- quote 
- content 
____
`
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Quote,
					types.AttrQuoteAuthor: "john doe",
					types.AttrQuoteTitle:  "",
				},
				Kind: types.Quote,
				Elements: []interface{}{
					types.UnorderedList{
						Attributes: types.ElementAttributes{},
						Items: []types.UnorderedListItem{
							{
								Attributes:  types.ElementAttributes{},
								Level:       1,
								BulletStyle: types.Dash,
								CheckStyle:  types.NoCheck,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: []types.InlineElements{
											{
												types.StringElement{
													Content: "some ",
												},
											},
										},
									},
								},
							},
							{
								Attributes:  types.ElementAttributes{},
								Level:       1,
								BulletStyle: types.Dash,
								CheckStyle:  types.NoCheck,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: []types.InlineElements{
											{
												types.StringElement{
													Content: "quote ",
												},
											},
										},
									},
								},
							},
							{
								Attributes:  types.ElementAttributes{},
								Level:       1,
								BulletStyle: types.Dash,
								CheckStyle:  types.NoCheck,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: []types.InlineElements{
											{
												types.StringElement{
													Content: "content ",
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
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("single-line quote with title only", func() {
			actualContent := `[quote, ,quote title]
____
some quote content 
____
`
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Quote,
					types.AttrQuoteAuthor: "",
					types.AttrQuoteTitle:  "quote title",
				},
				Kind: types.Quote,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "some quote content ",
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("multi-line quote with rendered lists and block and without author and title", func() {
			actualContent := `[quote]
____
* some
----
* quote 
----
* content
____`
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Quote,
					types.AttrQuoteAuthor: "",
					types.AttrQuoteTitle:  "",
				},
				Kind: types.Quote,
				Elements: []interface{}{
					types.UnorderedList{
						Attributes: types.ElementAttributes{},
						Items: []types.UnorderedListItem{
							{
								Attributes:  types.ElementAttributes{},
								Level:       1,
								BulletStyle: types.OneAsterisk,
								CheckStyle:  types.NoCheck,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: []types.InlineElements{
											{
												types.StringElement{
													Content: "some",
												},
											},
										},
									},
								},
							},
						},
					},
					types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						Kind:       types.Listing,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "* quote ",
										},
									},
								},
							},
						},
					},
					types.UnorderedList{
						Attributes: types.ElementAttributes{},
						Items: []types.UnorderedListItem{
							{
								Attributes:  types.ElementAttributes{},
								Level:       1,
								BulletStyle: types.OneAsterisk,
								CheckStyle:  types.NoCheck,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: []types.InlineElements{
											{
												types.StringElement{
													Content: "content",
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
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("multi-line quote with rendered list and without author and title", func() {
			actualContent := `[quote]
____
* some


* quote 


* content
____`
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Quote,
					types.AttrQuoteAuthor: "",
					types.AttrQuoteTitle:  "",
				},
				Kind: types.Quote,
				Elements: []interface{}{
					types.UnorderedList{
						Attributes: types.ElementAttributes{},
						Items: []types.UnorderedListItem{
							{
								Attributes:  types.ElementAttributes{},
								Level:       1,
								BulletStyle: types.OneAsterisk,
								CheckStyle:  types.NoCheck,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: []types.InlineElements{
											{
												types.StringElement{
													Content: "some",
												},
											},
										},
									},
								},
							},
							{
								Attributes:  types.ElementAttributes{},
								Level:       1,
								BulletStyle: types.OneAsterisk,
								CheckStyle:  types.NoCheck,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: []types.InlineElements{
											{
												types.StringElement{
													Content: "quote ",
												},
											},
										},
									},
								},
							},
							{
								Attributes:  types.ElementAttributes{},
								Level:       1,
								BulletStyle: types.OneAsterisk,
								CheckStyle:  types.NoCheck,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: []types.InlineElements{
											{
												types.StringElement{
													Content: "content",
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
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("empty quote without author and title", func() {
			actualContent := `[quote]
____
____`
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Quote,
					types.AttrQuoteAuthor: "",
					types.AttrQuoteTitle:  "",
				},
				Kind:     types.Quote,
				Elements: []interface{}{},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("unclosed quote without author and title", func() {
			actualContent := `[quote]
____
foo
`
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Quote,
					types.AttrQuoteAuthor: "",
					types.AttrQuoteTitle:  "",
				},
				Kind: types.Quote,
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
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})
	})

	Context("verse blocks", func() {

		It("single line verse with author and title", func() {
			actualContent := `[verse, john doe, verse title]
____
some *verse* content
____`
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Verse,
					types.AttrQuoteAuthor: "john doe",
					types.AttrQuoteTitle:  "verse title",
				},
				Kind: types.Verse,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "some ",
								},
								types.QuotedText{
									Kind: types.Bold,
									Elements: types.InlineElements{
										types.StringElement{
											Content: "verse",
										},
									},
								},
								types.StringElement{
									Content: " content",
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("multi-line verse with unrendered list author only", func() {
			actualContent := `[verse, john doe,   ]
____
- some 
- verse 
- content 
____
`
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Verse,
					types.AttrQuoteAuthor: "john doe",
					types.AttrQuoteTitle:  "",
				},
				Kind: types.Verse,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "- some ",
								},
							},
							{
								types.StringElement{
									Content: "- verse ",
								},
							},
							{
								types.StringElement{
									Content: "- content ",
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("multi-line verse with title only", func() {
			actualContent := `[verse, ,verse title]
____
some verse content 
____
`
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Verse,
					types.AttrQuoteAuthor: "",
					types.AttrQuoteTitle:  "verse title",
				},
				Kind: types.Verse,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "some verse content ",
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("multi-line verse with unrendered lists and block without author and title", func() {
			actualContent := `[verse]
____
* some
----
* verse 
----
* content
____`
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Verse,
					types.AttrQuoteAuthor: "",
					types.AttrQuoteTitle:  "",
				},
				Kind: types.Verse,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "* some",
								},
							},
							{
								types.StringElement{
									Content: "----",
								},
							},
							{
								types.StringElement{
									Content: "* verse ",
								},
							},
							{
								types.StringElement{
									Content: "----",
								},
							},
							{
								types.StringElement{
									Content: "* content",
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("multi-line verse with unrendered list without author and title", func() {
			actualContent := `[verse]
____
* foo


	* bar
____`
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Verse,
					types.AttrQuoteAuthor: "",
					types.AttrQuoteTitle:  "",
				},
				Kind: types.Verse,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "* foo",
								},
							},
							{},
							{},
							{
								types.StringElement{
									Content: "\t* bar",
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("empty verse without author and title", func() {
			actualContent := `[verse]
____
____`
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Verse,
					types.AttrQuoteAuthor: "",
					types.AttrQuoteTitle:  "",
				},
				Kind:     types.Verse,
				Elements: []interface{}{},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("unclosed verse without author and title", func() {
			actualContent := `[verse]
____
foo
`
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Verse,
					types.AttrQuoteAuthor: "",
					types.AttrQuoteTitle:  "",
				},
				Kind: types.Verse,
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
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})
	})

	Context("source blocks", func() {

		It("with source attribute only", func() {
			actualContent := `[source]
----
require 'sinatra'

get '/hi' do
  "Hello World!"
end
----`
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{
					types.AttrKind:     types.Source,
					types.AttrLanguage: "",
				},
				Kind: types.Source,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "require 'sinatra'",
								},
							},
							{},
							{
								types.StringElement{
									Content: "get '/hi' do",
								},
							},
							{
								types.StringElement{
									Content: "  \"Hello World!\"",
								},
							},
							{
								types.StringElement{
									Content: "end",
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("with title, source and languages attributes", func() {
			actualContent := `[source,ruby]
.Source block title
----
require 'sinatra'

get '/hi' do
  "Hello World!"
end
----`
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{
					types.AttrKind:     types.Source,
					types.AttrLanguage: "ruby",
					types.AttrTitle:    "Source block title",
				},
				Kind: types.Source,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "require 'sinatra'",
								},
							},
							{},
							{
								types.StringElement{
									Content: "get '/hi' do",
								},
							},
							{
								types.StringElement{
									Content: "  \"Hello World!\"",
								},
							},
							{
								types.StringElement{
									Content: "end",
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("with id, title, source and languages attributes", func() {
			actualContent := `[#id-for-source-block]
[source,ruby]
.app.rb
----
require 'sinatra'

get '/hi' do
  "Hello World!"
end
----`
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{
					types.AttrKind:     types.Source,
					types.AttrLanguage: "ruby",
					types.AttrID:       "id-for-source-block",
					types.AttrTitle:    "app.rb",
				},
				Kind: types.Source,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "require 'sinatra'",
								},
							},
							{},
							{
								types.StringElement{
									Content: "get '/hi' do",
								},
							},
							{
								types.StringElement{
									Content: "  \"Hello World!\"",
								},
							},
							{
								types.StringElement{
									Content: "end",
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})
	})

	Context("sidebar blocks", func() {

		It("sidebar block with paragraph", func() {
			actualContent := `****
some *verse* content
****`
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{},
				Kind:       types.Sidebar,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "some ",
								},
								types.QuotedText{
									Kind: types.Bold,
									Elements: types.InlineElements{
										types.StringElement{
											Content: "verse",
										},
									},
								},
								types.StringElement{
									Content: " content",
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("sidebar block with title, paragraph and sourcecode block", func() {
			actualContent := `.a title
****
some *verse* content
----
foo
bar
----
****`
			expectedResult := types.DelimitedBlock{
				Attributes: types.ElementAttributes{
					types.AttrTitle: "a title",
				},
				Kind: types.Sidebar,
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "some ",
								},
								types.QuotedText{
									Kind: types.Bold,
									Elements: types.InlineElements{
										types.StringElement{
											Content: "verse",
										},
									},
								},
								types.StringElement{
									Content: " content",
								},
							},
						},
					},
					types.DelimitedBlock{
						Attributes: types.ElementAttributes{},
						Kind:       types.Listing,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "foo",
										},
									},
									{
										types.StringElement{
											Content: "bar",
										},
									},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})
	})

})
