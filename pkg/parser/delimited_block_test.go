package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("delimited blocks", func() {

	Context("draft document", func() {

		Context("fenced block", func() {

			It("with single line", func() {
				content := "some fenced code"
				source := "```\n" + content + "\n" + "```"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Kind: types.Fenced,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: content,
								},
							},
						},
					},
				}
				result, err := ParseDraftDocument(source)
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(MatchDraftDocument(expected))
			})

			It("with no line", func() {
				source := "```\n```"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Kind:     types.Fenced,
							Elements: []interface{}{},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with multiple lines alone", func() {
				source := "```\nsome fenced code\nwith an empty line\n\nin the middle\n```"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Kind: types.Fenced,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "some fenced code",
								},
								types.VerbatimLine{
									Content: "with an empty line",
								},
								types.VerbatimLine{},
								types.VerbatimLine{
									Content: "in the middle",
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with multiple lines then a paragraph", func() {
				source := "```\nsome fenced code\nwith an empty line\n\nin the middle\n```\nthen a normal paragraph."
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Kind: types.Fenced,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "some fenced code",
								},
								types.VerbatimLine{
									Content: "with an empty line",
								},
								types.VerbatimLine{
									Content: "",
								},
								types.VerbatimLine{
									Content: "in the middle",
								},
							},
						},
						types.Paragraph{
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{
										Content: "then a normal paragraph.",
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

			It("after a paragraph", func() {
				content := "some fenced code"
				source := "a paragraph.\n\n```\n" + content + "\n" + "```\n"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{
										Content: "a paragraph.",
									},
								},
							},
						},
						types.BlankLine{},
						types.DelimitedBlock{
							Kind: types.Fenced,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: content,
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with unclosed delimiter", func() {
				source := "```\nEnd of file here"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Kind: types.Fenced,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "End of file here",
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with external link inside - without attributes", func() {
				source := "```" + "\n" +
					"a http://website.com\n" +
					"and more text on the\n" +
					"next lines\n" +
					"```"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Kind: types.Fenced,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "a http://website.com",
								},
								types.VerbatimLine{
									Content: "and more text on the",
								},
								types.VerbatimLine{
									Content: "next lines",
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with external link inside - with attributes", func() {
				source := "```" + "\n" +
					"a http://website.com[]" + "\n" +
					"and more text on the" + "\n" +
					"next lines" + "\n" +
					"```"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Kind: types.Fenced,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "a http://website.com[]",
								},
								types.VerbatimLine{
									Content: "and more text on the",
								},
								types.VerbatimLine{
									Content: "next lines",
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with unrendered list", func() {
				source := "```\n" +
					"* some \n" +
					"* listing \n" +
					"* content \n```"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Kind: types.Fenced,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "* some ",
								},
								types.VerbatimLine{
									Content: "* listing ",
								},
								types.VerbatimLine{
									Content: "* content ",
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})
		})

		Context("listing block", func() {

			It("with single line", func() {
				source := `----
some listing code
----`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "some listing code",
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with no line", func() {
				source := `----
----`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Kind:     types.Listing,
							Elements: []interface{}{},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with multiple lines alone", func() {
				source := `----
some listing code
with an empty line

in the middle
----`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "some listing code",
								},
								types.VerbatimLine{
									Content: "with an empty line",
								},
								types.VerbatimLine{
									Content: "",
								},
								types.VerbatimLine{
									Content: "in the middle",
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with unrendered list", func() {
				source := `----
* some 
* listing 
* content
----`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "* some ",
								},
								types.VerbatimLine{
									Content: "* listing ",
								},
								types.VerbatimLine{
									Content: "* content",
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with multiple lines then a paragraph", func() {
				source := `---- 
some listing code
with an empty line

in the middle
----
then a normal paragraph.`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "some listing code",
								},
								types.VerbatimLine{
									Content: "with an empty line",
								},
								types.VerbatimLine{
									Content: "",
								},
								types.VerbatimLine{
									Content: "in the middle",
								},
							},
						},
						types.Paragraph{
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{
										Content: "then a normal paragraph.",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("after a paragraph", func() {
				source := `a paragraph.

----
some listing code
----`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{
										Content: "a paragraph.",
									},
								},
							},
						},
						types.BlankLine{}, // blankline is required between paragraph and the next block
						types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "some listing code",
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with unclosed delimiter", func() {
				source := `----
End of file here.`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "End of file here.",
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with single callout", func() {
				source := `----
import <1>
----
<1> an import`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "import ",
									Callouts: []types.Callout{
										{
											Ref: 1,
										},
									},
								},
							},
						},
						types.CalloutListItem{
							Ref: 1,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "an import",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with multiple callouts on different lines", func() {
				source := `----
import <1>

func foo() {} <2>
----
<1> an import
<2> a func`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "import ",
									Callouts: []types.Callout{
										{
											Ref: 1,
										},
									},
								},
								types.VerbatimLine{
									Content: "",
								},
								types.VerbatimLine{
									Content: "func foo() {} ",
									Callouts: []types.Callout{
										{
											Ref: 2,
										},
									},
								},
							},
						},
						types.CalloutListItem{
							Ref: 1,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "an import",
											},
										},
									},
								},
							},
						},
						types.CalloutListItem{
							Ref: 2,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "a func",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with multiple callouts on same line", func() {
				source := `----
import <1> <2><3>

func foo() {} <4>
----
<1> an import
<2> a single import
<3> a single basic import
<4> a func`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "import ",
									Callouts: []types.Callout{
										{
											Ref: 1,
										},
										{
											Ref: 2,
										},
										{
											Ref: 3,
										},
									},
								},
								types.VerbatimLine{
									Content: "",
								},
								types.VerbatimLine{
									Content: "func foo() {} ",
									Callouts: []types.Callout{
										{
											Ref: 4,
										},
									},
								},
							},
						},
						types.CalloutListItem{
							Ref: 1,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "an import",
											},
										},
									},
								},
							},
						},
						types.CalloutListItem{
							Ref: 2,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "a single import",
											},
										},
									},
								},
							},
						},
						types.CalloutListItem{
							Ref: 3,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "a single basic import",
											},
										},
									},
								},
							},
						},
						types.CalloutListItem{
							Ref: 4,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "a func",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with invalid callout", func() {
				source := `----
import <a>
----
<a> an import`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "import <a>",
								},
							},
						},
						types.Paragraph{
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{
										Content: "<a> an import",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})
		})

		Context("example block", func() {

			It("with single line", func() {
				source := `====
some listing code
====`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Kind: types.Example,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
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
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with single line starting with a dot", func() {
				source := `====
.foo
====`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Kind: types.Example,
							Elements: []interface{}{
								types.Attributes{
									types.AttrTitle: "foo",
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with multiple lines", func() {
				source := `====
.foo
some listing code
with *bold content*

* and a list item
====`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Kind: types.Example,
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.Attributes{
										types.AttrTitle: "foo",
									},
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "some listing code",
											},
										},
										[]interface{}{
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
								types.UnorderedListItem{
									Level:       1,
									BulletStyle: types.OneAsterisk,
									CheckStyle:  types.NoCheck,
									Elements: []interface{}{
										types.Paragraph{
											Lines: []interface{}{
												[]interface{}{
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
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with unclosed delimiter", func() {
				source := `====
End of doc here`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Kind: types.Example,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "End of doc here",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with title", func() {
				source := `.example block title
====
foo
====`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrTitle: "example block title",
							},
							Kind: types.Example,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "foo",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("example block starting delimiter only", func() {
				source := `====`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Kind:     types.Example,
							Elements: []interface{}{},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})
		})

		Context("admonition block", func() {

			It("example block as admonition", func() {
				source := `[NOTE]
====
foo
====`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrAdmonitionKind: types.Note,
							},
							Kind: types.Example,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "foo",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))

			})

			It("as admonition", func() {
				source := `[NOTE]
----
multiple

paragraphs
----
`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrAdmonitionKind: types.Note,
							},
							Kind: types.Listing,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "multiple",
								},
								types.VerbatimLine{},
								types.VerbatimLine{
									Content: "paragraphs",
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})
		})

		Context("quote block", func() {

			It("single-line quote block with author and title", func() {
				source := `[quote, john doe, quote title]
____
some *quote* content
____`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrKind:        types.Quote,
								types.AttrQuoteAuthor: "john doe",
								types.AttrQuoteTitle:  "quote title",
							},
							Kind: types.Quote,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "some ",
											},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
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
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("multi-line quote with author only", func() {
				source := `[quote, john doe,   ]
____
- some 
- quote 
- content 
____
`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrKind:        types.Quote,
								types.AttrQuoteAuthor: "john doe",
							},
							Kind: types.Quote,
							Elements: []interface{}{
								types.UnorderedListItem{
									Level:       1,
									BulletStyle: types.Dash,
									CheckStyle:  types.NoCheck,
									Elements: []interface{}{
										types.Paragraph{
											Lines: []interface{}{
												[]interface{}{
													types.StringElement{
														Content: "some ",
													},
												},
											},
										},
									},
								},
								types.UnorderedListItem{
									Level:       1,
									BulletStyle: types.Dash,
									CheckStyle:  types.NoCheck,
									Elements: []interface{}{
										types.Paragraph{
											Lines: []interface{}{
												[]interface{}{
													types.StringElement{
														Content: "quote ",
													},
												},
											},
										},
									},
								},
								types.UnorderedListItem{
									Level:       1,
									BulletStyle: types.Dash,
									CheckStyle:  types.NoCheck,
									Elements: []interface{}{
										types.Paragraph{
											Lines: []interface{}{
												[]interface{}{
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
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("single-line quote with title only", func() {
				source := `[quote, ,quote title]
____
some quote content 
____
`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrKind:       types.Quote,
								types.AttrQuoteTitle: "quote title",
							},
							Kind: types.Quote,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "some quote content ",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("multi-line quote with rendered list and without author and title", func() {
				source := `[quote]
____
* some


* quote 


* content
____`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Quote,
							},
							Kind: types.Quote,
							Elements: []interface{}{
								types.UnorderedListItem{
									Level:       1,
									BulletStyle: types.OneAsterisk,
									CheckStyle:  types.NoCheck,
									Elements: []interface{}{
										types.Paragraph{
											Lines: []interface{}{
												[]interface{}{
													types.StringElement{
														Content: "some",
													},
												},
											},
										},
									},
								},
								types.BlankLine{},
								types.BlankLine{},
								types.UnorderedListItem{
									Level:       1,
									BulletStyle: types.OneAsterisk,
									CheckStyle:  types.NoCheck,
									Elements: []interface{}{
										types.Paragraph{
											Lines: []interface{}{
												[]interface{}{
													types.StringElement{
														Content: "quote ",
													},
												},
											},
										},
									},
								},
								types.BlankLine{},
								types.BlankLine{},
								types.UnorderedListItem{
									Level:       1,
									BulletStyle: types.OneAsterisk,
									CheckStyle:  types.NoCheck,
									Elements: []interface{}{
										types.Paragraph{
											Lines: []interface{}{
												[]interface{}{
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
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("empty quote without author and title", func() {
				source := `[quote]
____
____`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Quote,
							},
							Kind:     types.Quote,
							Elements: []interface{}{},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("unclosed quote without author and title", func() {
				source := `[quote]
____
foo
`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Quote,
							},
							Kind: types.Quote,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "foo",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})
		})

		Context("markdown-style quote block", func() {

			It("with single marker without author", func() {
				source := `> some text
on *multiple lines*`

				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Kind: types.MarkdownQuote,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "some text",
											},
										},
										[]interface{}{
											types.StringElement{
												Content: "on ",
											},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{
														Content: "multiple lines",
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
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with marker on each line without author", func() {
				source := `> some text
> on *multiple lines*`

				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Kind: types.MarkdownQuote,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "some text",
											},
										},
										[]interface{}{
											types.StringElement{
												Content: "on ",
											},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{
														Content: "multiple lines",
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
				result, err := ParseDraftDocument(source)
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(MatchDraftDocument(expected))
			})

			It("with marker on each line with author only", func() {
				source := `> some text
> on *multiple lines*
> -- John Doe`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrQuoteAuthor: "John Doe",
							},
							Kind: types.MarkdownQuote,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "some text",
											},
										},
										[]interface{}{
											types.StringElement{
												Content: "on ",
											},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{
														Content: "multiple lines",
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
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with marker on each line with author and title", func() {
				source := `.title
> some text
> on *multiple lines*
> -- John Doe`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrTitle:       "title",
								types.AttrQuoteAuthor: "John Doe",
							},
							Kind: types.MarkdownQuote,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "some text",
											},
										},
										[]interface{}{
											types.StringElement{
												Content: "on ",
											},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{
														Content: "multiple lines",
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
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with with author only", func() {
				source := `> -- John Doe`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrQuoteAuthor: "John Doe",
							},
							Kind:     types.MarkdownQuote,
							Elements: []interface{}{},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})
		})

		Context("verse block", func() {

			It("single line verse with author and title", func() {
				source := `[verse, john doe, verse title]
____
some *verse* content
____`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrKind:        types.Verse,
								types.AttrQuoteAuthor: "john doe",
								types.AttrQuoteTitle:  "verse title",
							},
							Kind: types.Verse,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "some ",
											},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
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
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("multi-line verse with unrendered list and author only", func() {
				source := `[verse, john doe,   ]
____
- some 
- verse 
- content 
____
`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrKind:        types.Verse,
								types.AttrQuoteAuthor: "john doe",
							},
							Kind: types.Verse,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "- some ",
											},
										},
										[]interface{}{
											types.StringElement{
												Content: "- verse ",
											},
										},
										[]interface{}{
											types.StringElement{
												Content: "- content ",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("multi-line verse with title only", func() {
				source := `[verse, ,verse title]
____
some verse content 
____
`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrKind:       types.Verse,
								types.AttrQuoteTitle: "verse title",
							},
							Kind: types.Verse,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "some verse content ",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("multi-line verse with unrendered lists and block without author and title", func() {
				source := `[verse]
____
* some
----
* verse 
----
* content
____`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Verse,
							},
							Kind: types.Verse,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "* some",
											},
										},
										[]interface{}{
											types.StringElement{
												Content: "----",
											},
										},
										[]interface{}{
											types.StringElement{
												Content: "* verse ",
											},
										},
										[]interface{}{
											types.StringElement{
												Content: "----",
											},
										},
										[]interface{}{
											types.StringElement{
												Content: "* content",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("multi-line verse with unrendered list without author and title", func() {
				source := `[verse]
____
* foo


	* bar
____`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Verse,
							},
							Kind: types.Verse,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "* foo",
											},
										},
									},
								},
								types.BlankLine{},
								types.BlankLine{},
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "\t* bar",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("empty verse without author and title", func() {
				source := `[verse]
____
____`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Verse,
							},
							Kind:     types.Verse,
							Elements: []interface{}{},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("unclosed verse without author and title", func() {
				source := `[verse]
____
foo
`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Verse,
							},
							Kind: types.Verse,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "foo",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})
		})

		Context("source block", func() {

			sourceCode := []interface{}{
				types.VerbatimLine{
					Content: "package foo",
				},
				types.VerbatimLine{},
				types.VerbatimLine{
					Content: "// Foo",
				},
				types.VerbatimLine{
					Content: "type Foo struct{",
				},
				types.VerbatimLine{
					Content: "    Bar string",
				},
				types.VerbatimLine{
					Content: "}",
				},
			}

			It("with source attribute only", func() {
				source := `[source]
----
package foo

// Foo
type Foo struct{
    Bar string
}
----`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Source,
							},
							Kind:     types.Source,
							Elements: sourceCode,
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with source attribute and comma", func() {
				source := `[source,]
----
package foo

// Foo
type Foo struct{
    Bar string
}
----`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Source,
							},
							Kind:     types.Source,
							Elements: sourceCode,
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with title, source and language attributes", func() {
				source := `[source,go]
.foo.go
----
package foo

// Foo
type Foo struct{
    Bar string
}
----`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrKind:     types.Source,
								types.AttrLanguage: "go",
								types.AttrTitle:    "foo.go",
							},
							Kind:     types.Source,
							Elements: sourceCode,
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with id, title, source and language and other attributes", func() {
				source := `[#id-for-source-block]
[source,go,linenums]
.foo.go
----
package foo

// Foo
type Foo struct{
    Bar string
}
----`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrKind:     types.Source,
								types.AttrLanguage: "go",
								types.AttrID:       "id-for-source-block",
								types.AttrCustomID: true,
								types.AttrTitle:    "foo.go",
								types.AttrLineNums: nil,
							},
							Kind:     types.Source,
							Elements: sourceCode,
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})
		})

		Context("sidebar block", func() {

			It("with paragraph", func() {
				source := `****
some *bold* content
****`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Kind: types.Sidebar,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "some ",
											},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{
														Content: "bold",
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
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with title, paragraph and sourcecode block", func() {
				source := `.a title
****
some *bold* content

----
foo
bar
----
****`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrTitle: "a title",
							},
							Kind: types.Sidebar,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "some ",
											},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{
														Content: "bold",
													},
												},
											},
											types.StringElement{
												Content: " content",
											},
										},
									},
								},
								types.BlankLine{},
								types.DelimitedBlock{
									Kind: types.Listing,
									Elements: []interface{}{
										types.VerbatimLine{
											Content: "foo",
										},
										types.VerbatimLine{
											Content: "bar",
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})
		})

		Context("passthrough block", func() {

			It("with title", func() {
				source := `.a title
++++
_foo_

*bar*
++++`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrTitle: "a title",
							},
							Kind: types.Passthrough,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "_foo_",
								},
								types.VerbatimLine{
									Content: "",
								},
								types.VerbatimLine{
									Content: "*bar*",
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})
		})

		Context("passthrough open block", func() {

			It("2-line paragraph followed by another paragraph", func() {
				source := `[pass]
_foo_
*bar*

another paragraph`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Passthrough,
							},
							Kind: types.Passthrough,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "_foo_",
								},
								types.VerbatimLine{
									Content: "*bar*",
								},
							},
						},
						types.BlankLine{},
						types.Paragraph{
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{
										Content: "another paragraph",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})
		})
	})

	Context("final document", func() {

		Context("fenced block", func() {

			It("with single line", func() {
				content := "some fenced code"
				source := "```\n" + content + "\n" + "```"
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Kind: types.Fenced,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: content,
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with no line", func() {
				source := "```\n```"
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Kind:     types.Fenced,
							Elements: []interface{}{},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with multiple lines alone", func() {
				source := "```\nsome fenced code\nwith an empty line\n\nin the middle\n```"
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Kind: types.Fenced,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "some fenced code",
								},
								types.VerbatimLine{
									Content: "with an empty line",
								},
								types.VerbatimLine{
									Content: "",
								},
								types.VerbatimLine{
									Content: "in the middle",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with multiple lines then a paragraph", func() {
				source := "```\nsome fenced code\nwith an empty line\n\nin the middle\n```\nthen a normal paragraph."
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Kind: types.Fenced,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "some fenced code",
								},
								types.VerbatimLine{
									Content: "with an empty line",
								},
								types.VerbatimLine{
									Content: "",
								},
								types.VerbatimLine{
									Content: "in the middle",
								},
							},
						},
						types.Paragraph{
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{Content: "then a normal paragraph."},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("after a paragraph", func() {
				content := "some fenced code"
				source := "a paragraph.\n\n```\n" + content + "\n" + "```\n"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{Content: "a paragraph."},
								},
							},
						},
						types.DelimitedBlock{
							Kind: types.Fenced,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: content,
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with unclosed delimiter", func() {
				source := "```\nEnd of file here"
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Kind: types.Fenced,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "End of file here",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with external link inside - without attributes", func() {
				source := "```\n" +
					"a http://website.com\n" +
					"and more text on the\n" +
					"next lines\n" +
					"```"
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Kind: types.Fenced,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "a http://website.com",
								},
								types.VerbatimLine{
									Content: "and more text on the",
								},
								types.VerbatimLine{
									Content: "next lines",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with external link inside - with attributes", func() {
				source := "```" + "\n" +
					"a http://website.com[]" + "\n" +
					"and more text on the" + "\n" +
					"next lines" + "\n" +
					"```"
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Kind: types.Fenced,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "a http://website.com[]",
								},
								types.VerbatimLine{
									Content: "and more text on the",
								},
								types.VerbatimLine{
									Content: "next lines",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with unrendered list", func() {
				source := "```\n" +
					"* some \n" +
					"* listing \n" +
					"* content \n```"
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Kind: types.Fenced,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "* some ",
								},
								types.VerbatimLine{
									Content: "* listing ",
								},
								types.VerbatimLine{
									Content: "* content ",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

		})

		Context("listing block", func() {

			It("with single line", func() {
				source := `----
some listing code
----`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "some listing code",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with no line", func() {
				source := `----
----`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Kind:     types.Listing,
							Elements: []interface{}{},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with multiple lines alone", func() {
				source := `----
some listing code
with an empty line

in the middle
----`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "some listing code",
								},
								types.VerbatimLine{
									Content: "with an empty line",
								},
								types.VerbatimLine{
									Content: "",
								},
								types.VerbatimLine{
									Content: "in the middle",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with unrendered list", func() {
				source := `----
* some 
* listing 
* content
----`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "* some ",
								},
								types.VerbatimLine{
									Content: "* listing ",
								},
								types.VerbatimLine{
									Content: "* content",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with multiple lines then a paragraph", func() {
				source := `---- 
some listing code
with an empty line

in the middle
----
then a normal paragraph.`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "some listing code",
								},
								types.VerbatimLine{
									Content: "with an empty line",
								},
								types.VerbatimLine{
									Content: "",
								},
								types.VerbatimLine{
									Content: "in the middle",
								},
							},
						},
						types.Paragraph{
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{Content: "then a normal paragraph."},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("after a paragraph", func() {
				source := `a paragraph.

----
some listing code
----`
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{Content: "a paragraph."},
								},
							},
						},
						types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "some listing code",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with unclosed delimiter", func() {
				source := `----
End of file here.`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "End of file here.",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with single callout", func() {
				source := `----
import <1>
----
<1> an import`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "import ",
									Callouts: []types.Callout{
										{
											Ref: 1,
										},
									},
								},
							},
						},
						types.CalloutList{
							Items: []types.CalloutListItem{
								{
									Ref: 1,
									Elements: []interface{}{
										types.Paragraph{
											Lines: []interface{}{
												[]interface{}{
													types.StringElement{
														Content: "an import",
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
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with multiple callouts on different lines", func() {
				source := `----
import <1>

func foo() {} <2>
----
<1> an import
<2> a func`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "import ",
									Callouts: []types.Callout{
										{
											Ref: 1,
										},
									},
								},
								types.VerbatimLine{
									Content: "",
								},
								types.VerbatimLine{
									Content: "func foo() {} ",
									Callouts: []types.Callout{
										{
											Ref: 2,
										},
									},
								},
							},
						},
						types.CalloutList{
							Items: []types.CalloutListItem{
								{
									Ref: 1,
									Elements: []interface{}{
										types.Paragraph{
											Lines: []interface{}{
												[]interface{}{
													types.StringElement{
														Content: "an import",
													},
												},
											},
										},
									},
								},
								{
									Ref: 2,
									Elements: []interface{}{
										types.Paragraph{
											Lines: []interface{}{
												[]interface{}{
													types.StringElement{
														Content: "a func",
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
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with multiple callouts on same line", func() {
				source := `----
import <1> <2><3>

func foo() {} <4>
----
<1> an import
<2> a single import
<3> a single basic import
<4> a func`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "import ",
									Callouts: []types.Callout{
										{
											Ref: 1,
										},
										{
											Ref: 2,
										},
										{
											Ref: 3,
										},
									},
								},
								types.VerbatimLine{
									Content: "",
								},
								types.VerbatimLine{
									Content: "func foo() {} ",
									Callouts: []types.Callout{
										{
											Ref: 4,
										},
									},
								},
							},
						},
						types.CalloutList{
							Items: []types.CalloutListItem{
								{
									Ref: 1,
									Elements: []interface{}{
										types.Paragraph{
											Lines: []interface{}{
												[]interface{}{
													types.StringElement{
														Content: "an import",
													},
												},
											},
										},
									},
								},
								{
									Ref: 2,
									Elements: []interface{}{
										types.Paragraph{
											Lines: []interface{}{
												[]interface{}{
													types.StringElement{
														Content: "a single import",
													},
												},
											},
										},
									},
								},
								{
									Ref: 3,
									Elements: []interface{}{
										types.Paragraph{
											Lines: []interface{}{
												[]interface{}{
													types.StringElement{
														Content: "a single basic import",
													},
												},
											},
										},
									},
								},
								{
									Ref: 4,
									Elements: []interface{}{
										types.Paragraph{
											Lines: []interface{}{
												[]interface{}{
													types.StringElement{
														Content: "a func",
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
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with invalid callout", func() {
				source := `----
import <a>
----
<a> an import`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "import <a>",
								},
							},
						},
						types.Paragraph{
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{
										Content: "<a> an import",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})

		Context("example block", func() {

			It("with single line", func() {
				source := `====
some listing code
====`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Kind: types.Example,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
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
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with single line starting with a dot", func() {
				source := `====
.foo
====`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Kind:     types.Example,
							Elements: []interface{}{},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with multiple lines", func() {
				source := `====
.foo
some listing code
with *bold content*

* and a list item
====`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Kind: types.Example,
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.Attributes{
										types.AttrTitle: "foo",
									},
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "some listing code",
											},
										},
										[]interface{}{
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
									Items: []types.UnorderedListItem{
										{
											Level:       1,
											BulletStyle: types.OneAsterisk,
											CheckStyle:  types.NoCheck,
											Elements: []interface{}{
												types.Paragraph{
													Lines: []interface{}{
														[]interface{}{
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
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with unclosed delimiter", func() {
				source := `====
End of file here`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Kind: types.Example,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "End of file here",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with title", func() {
				source := `.example block title
====
foo
====`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrTitle: "example block title",
							},
							Kind: types.Example,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "foo",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("example block starting delimiter only", func() {
				source := `====`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Kind:     types.Example,
							Elements: []interface{}{},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})

		Context("admonition block", func() {

			It("example block as admonition", func() {
				source := `[NOTE]
====
foo
====`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrAdmonitionKind: types.Note,
							},
							Kind: types.Example,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "foo",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))

			})

			It("example block as admonition with multiple lines", func() {
				source := `[NOTE]
====
multiple

paragraphs
====
`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrAdmonitionKind: types.Note,
							},
							Kind: types.Example,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "multiple",
											},
										},
									},
								},
								types.BlankLine{},
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
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
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})

		Context("quote block", func() {

			It("single-line quote block with author and title", func() {
				source := `[quote, john doe, quote title]
____
some *quote* content
____`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrKind:        types.Quote,
								types.AttrQuoteAuthor: "john doe",
								types.AttrQuoteTitle:  "quote title",
							},
							Kind: types.Quote,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "some ",
											},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
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
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("multi-line quote with author only", func() {
				source := `[quote, john doe,   ]
____
- some 
- quote 
- content 
____
`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrKind:        types.Quote,
								types.AttrQuoteAuthor: "john doe",
							},
							Kind: types.Quote,
							Elements: []interface{}{
								types.UnorderedList{
									Items: []types.UnorderedListItem{
										{
											Level:       1,
											BulletStyle: types.Dash,
											CheckStyle:  types.NoCheck,
											Elements: []interface{}{
												types.Paragraph{
													Lines: []interface{}{
														[]interface{}{
															types.StringElement{
																Content: "some ",
															},
														},
													},
												},
											},
										},
										{
											Level:       1,
											BulletStyle: types.Dash,
											CheckStyle:  types.NoCheck,
											Elements: []interface{}{
												types.Paragraph{
													Lines: []interface{}{
														[]interface{}{
															types.StringElement{
																Content: "quote ",
															},
														},
													},
												},
											},
										},
										{
											Level:       1,
											BulletStyle: types.Dash,
											CheckStyle:  types.NoCheck,
											Elements: []interface{}{
												types.Paragraph{
													Lines: []interface{}{
														[]interface{}{
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
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("single-line quote with title only", func() {
				source := `[quote, ,quote title]
____
some quote content 
____
`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrKind:       types.Quote,
								types.AttrQuoteTitle: "quote title",
							},
							Kind: types.Quote,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "some quote content ",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("multi-line quote with rendered lists and block and without author and title", func() {
				source := `[quote]
____
* some
----
* quote 
----
* content
____`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Quote,
							},
							Kind: types.Quote,
							Elements: []interface{}{
								types.UnorderedList{
									Items: []types.UnorderedListItem{
										{
											Level:       1,
											BulletStyle: types.OneAsterisk,
											CheckStyle:  types.NoCheck,
											Elements: []interface{}{
												types.Paragraph{
													Lines: []interface{}{
														[]interface{}{
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
									Kind: types.Listing,
									Elements: []interface{}{
										types.VerbatimLine{
											Content: "* quote ",
										},
									},
								},
								types.UnorderedList{
									Items: []types.UnorderedListItem{
										{
											Level:       1,
											BulletStyle: types.OneAsterisk,
											CheckStyle:  types.NoCheck,
											Elements: []interface{}{
												types.Paragraph{
													Lines: []interface{}{
														[]interface{}{
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
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("multi-line quote with rendered list and without author and title", func() {
				source := `[quote]
____
* some


* quote 


* content
____`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Quote,
							},
							Kind: types.Quote,
							Elements: []interface{}{
								types.UnorderedList{
									Items: []types.UnorderedListItem{
										{
											Level:       1,
											BulletStyle: types.OneAsterisk,
											CheckStyle:  types.NoCheck,
											Elements: []interface{}{
												types.Paragraph{
													Lines: []interface{}{
														[]interface{}{
															types.StringElement{
																Content: "some",
															},
														},
													},
												},
											},
										},
										{
											Level:       1,
											BulletStyle: types.OneAsterisk,
											CheckStyle:  types.NoCheck,
											Elements: []interface{}{
												types.Paragraph{
													Lines: []interface{}{
														[]interface{}{
															types.StringElement{
																Content: "quote ",
															},
														},
													},
												},
											},
										},
										{
											Level:       1,
											BulletStyle: types.OneAsterisk,
											CheckStyle:  types.NoCheck,
											Elements: []interface{}{
												types.Paragraph{
													Lines: []interface{}{
														[]interface{}{
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
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("empty quote without author and title", func() {
				source := `[quote]
____
____`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Quote,
							},
							Kind:     types.Quote,
							Elements: []interface{}{},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("unclosed quote without author and title", func() {
				source := `[quote]
____
foo
`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Quote,
							},
							Kind: types.Quote,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "foo",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})

		Context("verse block", func() {

			It("single line verse with author and title", func() {
				source := `[verse, john doe, verse title]
____
some *verse* content
____`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrKind:        types.Verse,
								types.AttrQuoteAuthor: "john doe",
								types.AttrQuoteTitle:  "verse title",
							},
							Kind: types.Verse,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "some ",
											},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
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
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("multi-line verse with unrendered list and author only", func() {
				source := `[verse, john doe,   ]
____
- some 
- verse 
- content 
____
`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrKind:        types.Verse,
								types.AttrQuoteAuthor: "john doe",
							},
							Kind: types.Verse,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "- some ",
											},
										},
										[]interface{}{
											types.StringElement{
												Content: "- verse ",
											},
										},
										[]interface{}{
											types.StringElement{
												Content: "- content ",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("multi-line verse with title only", func() {
				source := `[verse, ,verse title]
____
some verse content 
____
`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrKind:       types.Verse,
								types.AttrQuoteTitle: "verse title",
							},
							Kind: types.Verse,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "some verse content ",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("multi-line verse with unrendered lists and block without author and title", func() {
				source := `[verse]
____
* some
----
* verse 
----
* content
____`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Verse,
							},
							Kind: types.Verse,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "* some",
											},
										},
										[]interface{}{
											types.StringElement{
												Content: "----",
											},
										},
										[]interface{}{
											types.StringElement{
												Content: "* verse ",
											},
										},
										[]interface{}{
											types.StringElement{
												Content: "----",
											},
										},
										[]interface{}{
											types.StringElement{
												Content: "* content",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("multi-line verse with unrendered list without author and title", func() {
				source := `[verse]
____
* foo


	* bar
____`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Verse,
							},
							Kind: types.Verse,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "* foo",
											},
										},
									},
								},
								types.BlankLine{},
								types.BlankLine{},
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "\t* bar",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("empty verse without author and title", func() {
				source := `[verse]
____
____`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Verse,
							},
							Kind:     types.Verse,
							Elements: []interface{}{},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("unclosed verse without author and title", func() {
				source := `[verse]
____
foo
`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Verse,
							},
							Kind: types.Verse,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "foo",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})

		Context("source block", func() {

			It("with source attribute only", func() {
				source := `[source]
----
require 'sinatra'

get '/hi' do
  "Hello World!"
end
----`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Source,
							},
							Kind: types.Source,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "require 'sinatra'",
								},
								types.VerbatimLine{},
								types.VerbatimLine{
									Content: "get '/hi' do",
								},
								types.VerbatimLine{
									Content: "  \"Hello World!\"",
								},
								types.VerbatimLine{
									Content: "end",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with title, source and languages attributes", func() {
				source := `[source,ruby]
.Source block title
----
require 'sinatra'

get '/hi' do
  "Hello World!"
end
----`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrKind:     types.Source,
								types.AttrLanguage: "ruby",
								types.AttrTitle:    "Source block title",
							},
							Kind: types.Source,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "require 'sinatra'",
								},
								types.VerbatimLine{},
								types.VerbatimLine{
									Content: "get '/hi' do",
								},
								types.VerbatimLine{
									Content: "  \"Hello World!\"",
								},
								types.VerbatimLine{
									Content: "end",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with id, title, source and languages attributes", func() {
				source := `[#id-for-source-block]
[source,ruby]
.app.rb
----
require 'sinatra'

get '/hi' do
  "Hello World!"
end
----`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrKind:     types.Source,
								types.AttrLanguage: "ruby",
								types.AttrID:       "id-for-source-block",
								types.AttrCustomID: true,
								types.AttrTitle:    "app.rb",
							},
							Kind: types.Source,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "require 'sinatra'",
								},
								types.VerbatimLine{},
								types.VerbatimLine{
									Content: "get '/hi' do",
								},
								types.VerbatimLine{
									Content: "  \"Hello World!\"",
								},
								types.VerbatimLine{
									Content: "end",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})

		Context("sidebar block", func() {

			It("with paragraph", func() {
				source := `****
some *verse* content
****`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Kind: types.Sidebar,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "some ",
											},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
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
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with title, paragraph and sourcecode block", func() {
				source := `.a title
****
some *verse* content

----
foo
bar
----
****`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrTitle: "a title",
							},
							Kind: types.Sidebar,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "some ",
											},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
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
								types.BlankLine{}, // blankline is required between paragraph and the next block
								types.DelimitedBlock{
									Kind: types.Listing,
									Elements: []interface{}{
										types.VerbatimLine{
											Content: "foo",
										},
										types.VerbatimLine{
											Content: "bar",
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})

		Context("passthrough block", func() {

			It("with title", func() {
				source := `.a title
++++
_foo_

*bar*
++++`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrTitle: "a title",
							},
							Kind: types.Passthrough,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "_foo_",
								},
								types.VerbatimLine{
									Content: "",
								},
								types.VerbatimLine{
									Content: "*bar*",
								},
							},
						},
					},
				}
				result, err := ParseDocument(source)
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(MatchDocument(expected))
			})
		})

		Context("passthrough open block", func() {

			It("2-line paragraph followed by another paragraph", func() {
				source := `[pass]
_foo_
*bar*

another paragraph`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Passthrough,
							},
							Kind: types.Passthrough,
							Elements: []interface{}{
								types.VerbatimLine{
									Content: "_foo_",
								},
								types.VerbatimLine{
									Content: "*bar*",
								},
							},
						},
						types.Paragraph{
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{
										Content: "another paragraph",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})

		Context("markdown-style quote block", func() {

			It("with single marker without author", func() {
				source := `> some text
on *multiple lines*`

				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Kind: types.MarkdownQuote,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "some text",
											},
										},
										[]interface{}{
											types.StringElement{
												Content: "on ",
											},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{
														Content: "multiple lines",
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
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with marker on each line without author", func() {
				source := `> some text
> on *multiple lines*`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Kind: types.MarkdownQuote,
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "some text",
											},
										},
										[]interface{}{
											types.StringElement{
												Content: "on ",
											},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{
														Content: "multiple lines",
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
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with marker on each line with author only", func() {
				source := `> some text
> on *multiple lines*
> -- John Doe`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Kind: types.MarkdownQuote,
							Attributes: types.Attributes{
								types.AttrQuoteAuthor: "John Doe",
							},
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "some text",
											},
										},
										[]interface{}{
											types.StringElement{
												Content: "on ",
											},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{
														Content: "multiple lines",
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
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with marker on each line with author and title", func() {
				source := `.title
> some text
> on *multiple lines*
> -- John Doe`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Kind: types.MarkdownQuote,
							Attributes: types.Attributes{
								types.AttrTitle:       "title",
								types.AttrQuoteAuthor: "John Doe",
							},
							Elements: []interface{}{
								types.Paragraph{
									Lines: []interface{}{
										[]interface{}{
											types.StringElement{
												Content: "some text",
											},
										},
										[]interface{}{
											types.StringElement{
												Content: "on ",
											},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{
														Content: "multiple lines",
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
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with with author only", func() {
				source := `> -- John Doe`
				expected := types.Document{
					Elements: []interface{}{
						types.DelimitedBlock{
							Kind: types.MarkdownQuote,
							Attributes: types.Attributes{
								types.AttrQuoteAuthor: "John Doe",
							},
							Elements: []interface{}{},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})
	})
})
