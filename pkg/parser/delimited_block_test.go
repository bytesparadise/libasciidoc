package parser_test

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("delimited blocks", func() {

	Context("draft document", func() {

		Context("example block", func() {

			It("with single plaintext line", func() {
				source := `====
some listing code
====`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.ExampleBlock{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
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

				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with single line starting with a dot", func() {
				source := `====
.foo
====`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.ExampleBlock{
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

			It("with rich lines", func() {
				source := `====
.foo
some listing *bold code*
with _italic content_

* and a list item
====`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.ExampleBlock{
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.Attributes{
										types.AttrTitle: "foo",
									},
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "some listing ",
											},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{
														Content: "bold code",
													},
												},
											},
										},
										{
											types.StringElement{
												Content: "with ",
											},
											types.QuotedText{
												Kind: types.Italic,
												Elements: []interface{}{
													types.StringElement{
														Content: "italic content",
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
											Lines: [][]interface{}{
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
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with unclosed delimiter", func() {
				source := `====
End of doc here`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.ExampleBlock{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
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
					Elements: []interface{}{
						types.ExampleBlock{
							Attributes: types.Attributes{
								types.AttrTitle: "example block title",
							},
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
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
					Elements: []interface{}{
						types.ExampleBlock{},
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
					Elements: []interface{}{
						types.QuoteBlock{
							Attributes: types.Attributes{
								types.AttrKind:        types.Quote,
								types.AttrQuoteAuthor: "john doe",
								types.AttrQuoteTitle:  "quote title",
							},
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
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
				result, err := ParseDraftDocument(source)
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(MatchDraftDocument(expected))
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
					Elements: []interface{}{
						types.QuoteBlock{
							Attributes: types.Attributes{
								types.AttrKind:        types.Quote,
								types.AttrQuoteAuthor: "john doe",
							},
							Elements: []interface{}{
								types.UnorderedListItem{
									Level:       1,
									BulletStyle: types.Dash,
									CheckStyle:  types.NoCheck,
									Elements: []interface{}{
										types.Paragraph{
											Lines: [][]interface{}{
												{
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
											Lines: [][]interface{}{
												{
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
											Lines: [][]interface{}{
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
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("single-line quote with title only", func() {
				source := `[quote, ,quote title]
____
some quote content 
____
`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.QuoteBlock{
							Attributes: types.Attributes{
								types.AttrKind:       types.Quote,
								types.AttrQuoteTitle: "quote title",
							},
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
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
					Elements: []interface{}{
						types.QuoteBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Quote,
							},
							Elements: []interface{}{
								types.UnorderedListItem{
									Level:       1,
									BulletStyle: types.OneAsterisk,
									CheckStyle:  types.NoCheck,
									Elements: []interface{}{
										types.Paragraph{
											Lines: [][]interface{}{
												{
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
											Lines: [][]interface{}{
												{
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
											Lines: [][]interface{}{
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
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("empty quote without author and title", func() {
				source := `[quote]
____
____`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.QuoteBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Quote,
							},
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
					Elements: []interface{}{
						types.QuoteBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Quote,
							},
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
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

		Context("sidebar block", func() {

			It("with paragraph", func() {
				source := `****
some *bold* content
****`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.SidebarBlock{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
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
					Elements: []interface{}{
						types.SidebarBlock{
							Attributes: types.Attributes{
								types.AttrTitle: "a title",
							},
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
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
								types.ListingBlock{
									Lines: [][]interface{}{
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
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})
		})

		Context("with custom substitutions", func() {

			// testing custom substitutions on example blocks only, as
			// other verbatim blocks (fenced, literal, source, passthrough)
			// share the same implementation

			// also, see https://asciidoctor.org/docs/user-manual/#incremental-substitutions
			// "When you set the subs attribute on a block, you automatically remove all of its default substitutions.
			// For example, if you set subs on a literal block, and assign it a value of attributes,
			// only attributes are substituted."

			source := `:github-url: https://github.com
				
[subs="$SUBS"]
====
a link to https://example.com[] <1>
and <more text> on the +
*next* lines with a link to {github-url}[]

* a list item
====

<1> a callout
`

			It("should apply the default substitution", func() {
				s := strings.ReplaceAll(source, "[subs=\"$SUBS\"]\n", "")
				expected := types.DraftDocument{
					Attributes: types.Attributes{
						"github-url": "https://github.com",
					},
					Elements: []interface{}{
						types.AttributeDeclaration{
							Name:  "github-url",
							Value: "https://github.com",
						},
						types.BlankLine{},
						types.ExampleBlock{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a link to ",
											},
											types.InlineLink{
												Location: types.Location{
													Scheme: "https://",
													Path: []interface{}{
														types.StringElement{
															Content: "example.com",
														},
													},
												},
											},
											types.StringElement{
												Content: " ",
											},
											types.SpecialCharacter{
												Name: "<",
											},
											types.StringElement{
												Content: "1",
											},
											types.SpecialCharacter{
												Name: ">",
											},
										},
										{
											types.StringElement{
												Content: "and ",
											},
											types.SpecialCharacter{
												Name: "<",
											},
											types.StringElement{
												Content: "more text",
											},
											types.SpecialCharacter{
												Name: ">",
											},
											types.StringElement{
												Content: " on the",
											},
											types.LineBreak{},
										},
										{
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{
														Content: "next",
													},
												},
											},
											types.StringElement{
												Content: " lines with a link to ",
											},
											types.InlineLink{
												Location: types.Location{
													Scheme: "https://",
													Path: []interface{}{
														types.StringElement{
															Content: "github.com",
														},
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
											Lines: [][]interface{}{
												{
													types.StringElement{
														Content: "a list item",
													},
												},
											},
										},
									},
								},
							},
						},
						types.BlankLine{},
						types.CalloutListItem{
							Ref: 1,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a callout",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
			})

			It("should apply the 'normal' substitution", func() {
				s := strings.ReplaceAll(source, "$SUBS", "normal")
				expected := types.DraftDocument{
					Attributes: types.Attributes{
						"github-url": "https://github.com",
					},
					Elements: []interface{}{
						types.AttributeDeclaration{
							Name:  "github-url",
							Value: "https://github.com",
						},
						types.BlankLine{},
						types.ExampleBlock{
							Attributes: types.Attributes{
								types.AttrSubstitutions: "normal",
							},
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a link to ",
											},
											types.InlineLink{
												Location: types.Location{
													Scheme: "https://",
													Path: []interface{}{
														types.StringElement{
															Content: "example.com",
														},
													},
												},
											},
											types.StringElement{
												Content: " ",
											},
											types.SpecialCharacter{ // callout is not detected with the `normal` susbtitution
												Name: "<",
											},
											types.StringElement{
												Content: "1",
											},
											types.SpecialCharacter{
												Name: ">",
											},
										},
										{
											types.StringElement{
												Content: "and ",
											},
											types.SpecialCharacter{
												Name: "<",
											},
											types.StringElement{
												Content: "more text",
											},
											types.SpecialCharacter{
												Name: ">",
											},
											types.StringElement{
												Content: " on the",
											},
											types.LineBreak{},
										},
										{
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{
														Content: "next",
													},
												},
											},
											types.StringElement{
												Content: " lines with a link to ",
											},
											types.InlineLink{
												Location: types.Location{
													Scheme: "https://",
													Path: []interface{}{
														types.StringElement{
															Content: "github.com",
														},
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
											Lines: [][]interface{}{
												{
													types.StringElement{
														Content: "a list item",
													},
												},
											},
										},
									},
								},
							},
						},
						types.BlankLine{},
						types.CalloutListItem{
							Ref: 1,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a callout",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
			})

			It("should apply the 'quotes' substitution", func() {
				s := strings.ReplaceAll(source, "$SUBS", "quotes")
				expected := types.DraftDocument{
					Attributes: types.Attributes{
						"github-url": "https://github.com",
					},
					Elements: []interface{}{
						types.AttributeDeclaration{
							Name:  "github-url",
							Value: "https://github.com",
						},
						types.BlankLine{},
						types.ExampleBlock{
							Attributes: types.Attributes{
								types.AttrSubstitutions: "quotes",
							},
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a link to https://example.com[] <1>",
											},
										},
										{
											types.StringElement{
												Content: "and <more text> on the +",
											},
										},
										{
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{
														Content: "next",
													},
												},
											},
											types.StringElement{
												Content: " lines with a link to {github-url}[]",
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
											Lines: [][]interface{}{
												{
													types.StringElement{
														Content: "a list item",
													},
												},
											},
										},
									},
								},
							},
						},
						types.BlankLine{},
						types.CalloutListItem{
							Ref: 1,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a callout",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
			})

			It("should apply the 'macros' substitution", func() {
				s := strings.ReplaceAll(source, "$SUBS", "macros")
				expected := types.DraftDocument{
					Attributes: types.Attributes{
						"github-url": "https://github.com",
					},
					Elements: []interface{}{
						types.AttributeDeclaration{
							Name:  "github-url",
							Value: "https://github.com",
						},
						types.BlankLine{},
						types.ExampleBlock{
							Attributes: types.Attributes{
								types.AttrSubstitutions: "macros",
							},
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a link to ",
											},
											types.InlineLink{
												Location: types.Location{
													Scheme: "https://",
													Path: []interface{}{
														types.StringElement{
															Content: "example.com",
														},
													},
												},
											},
											types.StringElement{
												Content: " <1>",
											},
										},
										{
											types.StringElement{
												Content: "and <more text> on the +",
											},
										},
										{
											types.StringElement{
												Content: "*next* lines with a link to {github-url}[]",
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
											Lines: [][]interface{}{
												{
													types.StringElement{
														Content: "a list item",
													},
												},
											},
										},
									},
								},
							},
						},
						types.BlankLine{},
						types.CalloutListItem{
							Ref: 1,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a callout",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
			})

			It("should apply the 'attributes' substitution", func() {
				s := strings.ReplaceAll(source, "$SUBS", "attributes")
				expected := types.DraftDocument{
					Attributes: types.Attributes{
						"github-url": "https://github.com",
					},
					Elements: []interface{}{
						types.AttributeDeclaration{
							Name:  "github-url",
							Value: "https://github.com",
						},
						types.BlankLine{},
						types.ExampleBlock{
							Attributes: types.Attributes{
								types.AttrSubstitutions: "attributes",
							},
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a link to https://example.com[] <1>",
											},
										},
										{
											types.StringElement{
												Content: "and <more text> on the +",
											},
										},
										{
											types.StringElement{
												Content: "*next* lines with a link to https://github.com[]",
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
											Lines: [][]interface{}{
												{
													types.StringElement{
														Content: "a list item",
													},
												},
											},
										},
									},
								},
							},
						},
						types.BlankLine{},
						types.CalloutListItem{
							Ref: 1,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a callout",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
			})

			It("should apply the 'attributes,macros' substitution", func() {
				s := strings.ReplaceAll(source, "$SUBS", "attributes,macros")
				expected := types.DraftDocument{
					Attributes: types.Attributes{
						"github-url": "https://github.com",
					},
					Elements: []interface{}{
						types.AttributeDeclaration{
							Name:  "github-url",
							Value: "https://github.com",
						},
						types.BlankLine{},
						types.ExampleBlock{
							Attributes: types.Attributes{
								types.AttrSubstitutions: "attributes,macros",
							},
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a link to ",
											},
											types.InlineLink{
												Location: types.Location{
													Scheme: "https://",
													Path: []interface{}{
														types.StringElement{
															Content: "example.com",
														},
													},
												},
											},
											types.StringElement{
												Content: " <1>",
											},
										},
										{
											types.StringElement{
												Content: "and <more text> on the +",
											},
										},
										{
											types.StringElement{
												Content: "*next* lines with a link to ",
											},
											types.InlineLink{
												Location: types.Location{
													Scheme: "https://",
													Path: []interface{}{
														types.StringElement{
															Content: "github.com",
														},
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
											Lines: [][]interface{}{
												{
													types.StringElement{
														Content: "a list item",
													},
												},
											},
										},
									},
								},
							},
						},
						types.BlankLine{},
						types.CalloutListItem{
							Ref: 1,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a callout",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
			})

			It("should apply the 'specialchars' substitution", func() {
				s := strings.ReplaceAll(source, "$SUBS", "specialchars")
				expected := types.DraftDocument{
					Attributes: types.Attributes{
						"github-url": "https://github.com",
					},
					Elements: []interface{}{
						types.AttributeDeclaration{
							Name:  "github-url",
							Value: "https://github.com",
						},
						types.BlankLine{},
						types.ExampleBlock{
							Attributes: types.Attributes{
								types.AttrSubstitutions: "specialchars",
							},
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a link to https://example.com[] ",
											},
											types.SpecialCharacter{
												Name: "<",
											},
											types.StringElement{
												Content: "1",
											},
											types.SpecialCharacter{
												Name: ">",
											},
										},
										{
											types.StringElement{
												Content: "and ",
											},
											types.SpecialCharacter{
												Name: "<",
											},
											types.StringElement{
												Content: "more text",
											},
											types.SpecialCharacter{
												Name: ">",
											},
											types.StringElement{
												Content: " on the +",
											},
										},
										{
											types.StringElement{
												Content: "*next* lines with a link to {github-url}[]",
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
											Lines: [][]interface{}{
												{
													types.StringElement{
														Content: "a list item",
													},
												},
											},
										},
									},
								},
							},
						},
						types.BlankLine{},
						types.CalloutListItem{
							Ref: 1,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a callout",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
			})

			It("should apply the 'replacements' substitution", func() {
				s := strings.ReplaceAll(source, "$SUBS", "replacements")
				expected := types.DraftDocument{
					Attributes: types.Attributes{
						"github-url": "https://github.com",
					},
					Elements: []interface{}{
						types.AttributeDeclaration{
							Name:  "github-url",
							Value: "https://github.com",
						},
						types.BlankLine{},
						types.ExampleBlock{
							Attributes: types.Attributes{
								types.AttrSubstitutions: "replacements",
							},
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a link to https://example.com[] <1>",
											},
										},
										{
											types.StringElement{
												Content: "and <more text> on the +",
											},
										},
										{
											types.StringElement{
												Content: "*next* lines with a link to {github-url}[]",
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
											Lines: [][]interface{}{
												{
													types.StringElement{
														Content: "a list item",
													},
												},
											},
										},
									},
								},
							},
						},
						types.BlankLine{},
						types.CalloutListItem{
							Ref: 1,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a callout",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
			})

			It("should apply the 'post_replacements' substitution", func() {
				s := strings.ReplaceAll(source, "$SUBS", "post_replacements")
				expected := types.DraftDocument{
					Attributes: types.Attributes{
						"github-url": "https://github.com",
					},
					Elements: []interface{}{
						types.AttributeDeclaration{
							Name:  "github-url",
							Value: "https://github.com",
						},
						types.BlankLine{},
						types.ExampleBlock{
							Attributes: types.Attributes{
								types.AttrSubstitutions: "post_replacements",
							},
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a link to https://example.com[] <1>",
											},
										},
										{
											types.StringElement{
												Content: "and <more text> on the",
											},
											types.LineBreak{},
										},
										{
											types.StringElement{
												Content: "*next* lines with a link to {github-url}[]",
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
											Lines: [][]interface{}{
												{
													types.StringElement{
														Content: "a list item",
													},
												},
											},
										},
									},
								},
							},
						},
						types.BlankLine{},
						types.CalloutListItem{
							Ref: 1,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a callout",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
			})

			It("should apply the 'quotes,macros' substitution", func() {
				s := strings.ReplaceAll(source, "$SUBS", "quotes,macros")
				expected := types.DraftDocument{
					Attributes: types.Attributes{
						"github-url": "https://github.com",
					},
					Elements: []interface{}{
						types.AttributeDeclaration{
							Name:  "github-url",
							Value: "https://github.com",
						},
						types.BlankLine{},
						types.ExampleBlock{
							Attributes: types.Attributes{
								types.AttrSubstitutions: "quotes,macros",
							},
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a link to ",
											},
											types.InlineLink{
												Location: types.Location{
													Scheme: "https://",
													Path: []interface{}{
														types.StringElement{
															Content: "example.com",
														},
													},
												},
											},
											types.StringElement{
												Content: " <1>",
											},
										},
										{
											types.StringElement{
												Content: "and <more text> on the +",
											},
										},
										{
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{
														Content: "next",
													},
												},
											},
											types.StringElement{
												Content: " lines with a link to {github-url}[]",
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
											Lines: [][]interface{}{
												{
													types.StringElement{
														Content: "a list item",
													},
												},
											},
										},
									},
								},
							},
						},
						types.BlankLine{},
						types.CalloutListItem{
							Ref: 1,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a callout",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
			})

			It("should apply the 'macros,quotes' substitution", func() {
				s := strings.ReplaceAll(source, "$SUBS", "macros,quotes")
				expected := types.DraftDocument{
					Attributes: types.Attributes{
						"github-url": "https://github.com",
					},
					Elements: []interface{}{
						types.AttributeDeclaration{
							Name:  "github-url",
							Value: "https://github.com",
						},
						types.BlankLine{},
						types.ExampleBlock{
							Attributes: types.Attributes{
								types.AttrSubstitutions: "macros,quotes",
							},
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a link to ",
											},
											types.InlineLink{
												Location: types.Location{
													Scheme: "https://",
													Path: []interface{}{
														types.StringElement{
															Content: "example.com",
														},
													},
												},
											},
											types.StringElement{
												Content: " <1>",
											},
										},
										{
											types.StringElement{
												Content: "and <more text> on the +",
											},
										},
										{
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{
														Content: "next",
													},
												},
											},
											types.StringElement{
												Content: " lines with a link to {github-url}[]",
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
											Lines: [][]interface{}{
												{
													types.StringElement{
														Content: "a list item",
													},
												},
											},
										},
									},
								},
							},
						},
						types.BlankLine{},
						types.CalloutListItem{
							Ref: 1,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a callout",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
			})

			It("should apply the 'none' substitution", func() {
				s := strings.ReplaceAll(source, "$SUBS", "none")
				expected := types.DraftDocument{
					Attributes: types.Attributes{
						"github-url": "https://github.com",
					},
					Elements: []interface{}{
						types.AttributeDeclaration{
							Name:  "github-url",
							Value: "https://github.com",
						},
						types.BlankLine{},
						types.ExampleBlock{
							Attributes: types.Attributes{
								types.AttrSubstitutions: "none",
							},
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a link to https://example.com[] <1>",
											},
										},
										{
											types.StringElement{
												Content: "and <more text> on the +",
											},
										},
										{
											types.StringElement{
												Content: "*next* lines with a link to {github-url}[]",
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
											Lines: [][]interface{}{
												{
													types.StringElement{
														Content: "a list item",
													},
												},
											},
										},
									},
								},
							},
						},
						types.BlankLine{},
						types.CalloutListItem{
							Ref: 1,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a callout",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
			})
		})

		Context("fenced block", func() {

			It("with single line", func() {
				content := "some fenced code"
				source := "```\n" + content + "\n" + "```"
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.FencedBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: content,
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

			It("with special characters line", func() {
				source := "```\n<some fenced code>\n" + "```"
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.FencedBlock{
							Lines: [][]interface{}{
								{
									types.SpecialCharacter{
										Name: "<",
									},
									types.StringElement{
										Content: "some fenced code",
									},
									types.SpecialCharacter{
										Name: ">",
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

			It("with no line", func() {
				source := "```\n```"
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.FencedBlock{
							Lines: [][]interface{}{
								{},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with multiple lines alone", func() {
				source := "```\nsome fenced code\nwith an empty line\n\nin the middle\n```"
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.FencedBlock{
							Lines: [][]interface{}{
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
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with multiple lines then a paragraph", func() {
				source := "```\nsome fenced code\nwith an empty line\n\nin the middle\n```\nthen a normal paragraph."
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.FencedBlock{
							Lines: [][]interface{}{
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
								{}, // empty line
								{
									types.StringElement{
										Content: "in the middle",
									},
								},
							},
						},
						types.Paragraph{
							Lines: [][]interface{}{
								{
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
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a paragraph.",
									},
								},
							},
						},
						types.BlankLine{},
						types.FencedBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: content,
									},
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
					Elements: []interface{}{
						types.FencedBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "End of file here",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with external link inside - without attributes", func() {
				source := "```" + "\n" +
					"a https://example.com\n" +
					"and more text on the\n" +
					"next lines\n" +
					"```"
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.FencedBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a https://example.com",
									},
								},
								{
									types.StringElement{
										Content: "and more text on the",
									},
								},
								{
									types.StringElement{
										Content: "next lines",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with external link inside - with attributes", func() {
				source := "```" + "\n" +
					"a https://example.com[]" + "\n" +
					"and more text on the" + "\n" +
					"next lines" + "\n" +
					"```"
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.FencedBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a https://example.com[]",
									},
								},
								{
									types.StringElement{
										Content: "and more text on the",
									},
								},
								{
									types.StringElement{
										Content: "next lines",
									},
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
					Elements: []interface{}{
						types.FencedBlock{
							Lines: [][]interface{}{
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
										Content: "* content ",
									},
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
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "some listing code",
									},
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
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{},
							},
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
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
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
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with unrendered list", func() {
				source := `----
* some 
* listing 
* content
----`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
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
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
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
						types.Paragraph{
							Lines: [][]interface{}{
								{
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
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a paragraph.",
									},
								},
							},
						},
						types.BlankLine{}, // blankline is required between paragraph and the next block
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "some listing code",
									},
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
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "End of file here.",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with single callout", func() {
				source := `----
<import> <1>
----
<1> an import`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.SpecialCharacter{
										Name: "<",
									},
									types.StringElement{
										Content: "import",
									},
									types.SpecialCharacter{
										Name: ">",
									},
									types.StringElement{
										Content: " ",
									},
									types.Callout{
										Ref: 1,
									},
								},
							},
						},
						types.CalloutListItem{
							Ref: 1,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
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
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "import ",
									},
									types.Callout{
										Ref: 1,
									},
								},
								{},
								{
									types.StringElement{
										Content: "func foo() {} ",
									},
									types.Callout{
										Ref: 2,
									},
								},
							},
						},
						types.CalloutListItem{
							Ref: 1,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
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
									Lines: [][]interface{}{
										{
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
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "import ",
									},
									types.Callout{
										Ref: 1,
									},
									types.Callout{
										Ref: 2,
									},
									types.Callout{
										Ref: 3,
									},
								},
								{},
								{
									types.StringElement{
										Content: "func foo() {} ",
									},
									types.Callout{
										Ref: 4,
									},
								},
							},
						},
						types.CalloutListItem{
							Ref: 1,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
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
									Lines: [][]interface{}{
										{
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
									Lines: [][]interface{}{
										{
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
									Lines: [][]interface{}{
										{
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
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "import ",
									},
									types.SpecialCharacter{
										Name: "<",
									},
									types.StringElement{
										Content: "a",
									},
									types.SpecialCharacter{
										Name: ">",
									},
								},
							},
						},
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.SpecialCharacter{
										Name: "<",
									},
									types.StringElement{
										Content: "a",
									},
									types.SpecialCharacter{
										Name: ">",
									},
									types.StringElement{
										Content: " an import",
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

			sourceCode := [][]interface{}{
				{
					types.StringElement{
						Content: "package foo",
					},
				},
				{},
				{
					types.StringElement{
						Content: "// Foo",
					},
				},
				{
					types.StringElement{
						Content: "type Foo struct{",
					},
				},
				{
					types.StringElement{
						Content: "    Bar string",
					},
				},
				{
					types.StringElement{
						Content: "}",
					},
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
					Elements: []interface{}{
						types.ListingBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Source,
							},
							Lines: sourceCode,
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
					Elements: []interface{}{
						types.ListingBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Source,
							},
							Lines: sourceCode,
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
					Elements: []interface{}{
						types.ListingBlock{
							Attributes: types.Attributes{
								types.AttrKind:     types.Source,
								types.AttrLanguage: "go",
								types.AttrTitle:    "foo.go",
							},
							Lines: sourceCode,
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
					Elements: []interface{}{
						types.ListingBlock{
							Attributes: types.Attributes{
								types.AttrKind:     types.Source,
								types.AttrLanguage: "go",
								types.AttrID:       "id-for-source-block",
								types.AttrCustomID: true,
								types.AttrTitle:    "foo.go",
								types.AttrLineNums: nil,
							},
							Lines: sourceCode,
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
					Elements: []interface{}{
						types.PassthroughBlock{
							Attributes: types.Attributes{
								types.AttrTitle: "a title",
							},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "_foo_",
									},
								},
								{},
								{
									types.StringElement{
										Content: "*bar*",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with special characters", func() {
				source := `++++
<input>

<input>
++++`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.PassthroughBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "<input>",
									},
								},
								{},
								{
									types.StringElement{
										Content: "<input>",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with inline link", func() {
				source := `++++
http://example.com[]
++++`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.PassthroughBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "http://example.com[]",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with inline pass", func() {
				source := `++++
pass:[foo]
++++`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.PassthroughBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "pass:[foo]",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with quoted text", func() {
				source := `++++
*foo*
++++`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.PassthroughBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "*foo*",
									},
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
					Elements: []interface{}{
						types.PassthroughBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Passthrough,
							},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "_foo_",
									},
								},
								{
									types.StringElement{
										Content: "*bar*",
									},
								},
							},
						},
						types.BlankLine{},
						types.Paragraph{
							Lines: [][]interface{}{
								{
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

		Context("with custom substitutions", func() {

			// testing custom substitutions on listing blocks only, as
			// other verbatim blocks (fenced, literal, source, passthrough)
			// share the same implementation

			source := `:github-url: https://github.com

[subs="$SUBS"]
----
a link to https://example.com[] <1>
and <more text> on the +
*next* lines with a link to {github-url}[]

* not a list item
----

<1> a callout
`

			It("should apply the default substitution", func() {
				s := strings.ReplaceAll(source, "[subs=\"$SUBS\"]\n", "") // remove the 'subs' attribute
				expected := types.DraftDocument{
					Attributes: types.Attributes{
						"github-url": "https://github.com",
					},
					Elements: []interface{}{
						types.AttributeDeclaration{
							Name:  "github-url",
							Value: "https://github.com",
						},
						types.BlankLine{},
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a link to https://example.com[] ",
									},
									types.Callout{
										Ref: 1,
									},
								},
								{
									types.StringElement{
										Content: "and ",
									},
									types.SpecialCharacter{
										Name: "<",
									},
									types.StringElement{
										Content: "more text",
									},
									types.SpecialCharacter{
										Name: ">",
									},
									types.StringElement{
										Content: " on the +",
									},
								},
								{
									types.StringElement{
										Content: "*next* lines with a link to {github-url}[]",
									},
								},
								{},
								{
									types.StringElement{
										Content: "* not a list item",
									},
								},
							},
						},
						types.BlankLine{},
						types.CalloutListItem{
							Ref: 1,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a callout",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
			})

			It("should apply the 'normal' substitution", func() {
				s := strings.ReplaceAll(source, "$SUBS", "normal")
				expected := types.DraftDocument{
					Attributes: types.Attributes{
						"github-url": "https://github.com",
					},
					Elements: []interface{}{
						types.AttributeDeclaration{
							Name:  "github-url",
							Value: "https://github.com",
						},
						types.BlankLine{},
						types.ListingBlock{
							Attributes: types.Attributes{
								types.AttrSubstitutions: "normal",
							},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a link to ",
									},
									types.InlineLink{
										Location: types.Location{
											Scheme: "https://",
											Path: []interface{}{
												types.StringElement{
													Content: "example.com",
												},
											},
										},
									},
									types.StringElement{
										Content: " ",
									},
									types.SpecialCharacter{ // callout is not detected with the `normal` susbtitution
										Name: "<",
									},
									types.StringElement{
										Content: "1",
									},
									types.SpecialCharacter{
										Name: ">",
									},
								},
								{
									types.StringElement{
										Content: "and ",
									},
									types.SpecialCharacter{
										Name: "<",
									},
									types.StringElement{
										Content: "more text",
									},
									types.SpecialCharacter{
										Name: ">",
									},
									types.StringElement{
										Content: " on the",
									},
									types.LineBreak{},
								},
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{
												Content: "next",
											},
										},
									},
									types.StringElement{
										Content: " lines with a link to ",
									},
									types.InlineLink{
										Location: types.Location{
											Scheme: "https://",
											Path: []interface{}{
												types.StringElement{
													Content: "github.com",
												},
											},
										},
									},
								},
								{},
								{
									types.StringElement{
										Content: "* not a list item",
									},
								},
							},
						},
						types.BlankLine{},
						types.CalloutListItem{
							Ref: 1,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a callout",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
			})

			It("should apply the 'quotes' substitution", func() {
				s := strings.ReplaceAll(source, "$SUBS", "quotes")
				expected := types.DraftDocument{
					Attributes: types.Attributes{
						"github-url": "https://github.com",
					},
					Elements: []interface{}{
						types.AttributeDeclaration{
							Name:  "github-url",
							Value: "https://github.com",
						},
						types.BlankLine{},
						types.ListingBlock{
							Attributes: types.Attributes{
								types.AttrSubstitutions: "quotes",
							},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a link to https://example.com[] <1>",
									},
								},
								{
									types.StringElement{
										Content: "and <more text> on the +",
									},
								},
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{
												Content: "next",
											},
										},
									},
									types.StringElement{
										Content: " lines with a link to {github-url}[]",
									},
								},
								{},
								{
									types.StringElement{
										Content: "* not a list item",
									},
								},
							},
						},
						types.BlankLine{},
						types.CalloutListItem{
							Ref: 1,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a callout",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
			})

			It("should apply the 'macros' substitution", func() {
				s := strings.ReplaceAll(source, "$SUBS", "macros")
				expected := types.DraftDocument{
					Attributes: types.Attributes{
						"github-url": "https://github.com",
					},
					Elements: []interface{}{
						types.AttributeDeclaration{
							Name:  "github-url",
							Value: "https://github.com",
						},
						types.BlankLine{},
						types.ListingBlock{
							Attributes: types.Attributes{
								types.AttrSubstitutions: "macros",
							},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a link to ",
									},
									types.InlineLink{
										Location: types.Location{
											Scheme: "https://",
											Path: []interface{}{
												types.StringElement{
													Content: "example.com",
												},
											},
										},
									},
									types.StringElement{
										Content: " <1>",
									},
								},
								{
									types.StringElement{
										Content: "and <more text> on the +",
									},
								},
								{
									types.StringElement{
										Content: "*next* lines with a link to {github-url}[]",
									},
								},
								{},
								{
									types.StringElement{
										Content: "* not a list item",
									},
								},
							},
						},
						types.BlankLine{},
						types.CalloutListItem{
							Ref: 1,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a callout",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
			})

			It("should apply the 'attributes' substitution", func() {
				s := strings.ReplaceAll(source, "$SUBS", "attributes")
				expected := types.DraftDocument{
					Attributes: types.Attributes{
						"github-url": "https://github.com",
					},
					Elements: []interface{}{
						types.AttributeDeclaration{
							Name:  "github-url",
							Value: "https://github.com",
						},
						types.BlankLine{},
						types.ListingBlock{
							Attributes: types.Attributes{
								types.AttrSubstitutions: "attributes",
							},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a link to https://example.com[] <1>",
									},
								},
								{
									types.StringElement{
										Content: "and <more text> on the +",
									},
								},
								{
									types.StringElement{
										Content: "*next* lines with a link to https://github.com[]",
									},
								},
								{},
								{
									types.StringElement{
										Content: "* not a list item",
									},
								},
							},
						},
						types.BlankLine{},
						types.CalloutListItem{
							Ref: 1,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a callout",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
			})

			It("should apply the 'attributes,macros' substitution", func() {
				s := strings.ReplaceAll(source, "$SUBS", "attributes,macros")
				expected := types.DraftDocument{
					Attributes: types.Attributes{
						"github-url": "https://github.com",
					},
					Elements: []interface{}{
						types.AttributeDeclaration{
							Name:  "github-url",
							Value: "https://github.com",
						},
						types.BlankLine{},
						types.ListingBlock{
							Attributes: types.Attributes{
								types.AttrSubstitutions: "attributes,macros",
							},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a link to ",
									},
									types.InlineLink{
										Location: types.Location{
											Scheme: "https://",
											Path: []interface{}{
												types.StringElement{
													Content: "example.com",
												},
											},
										},
									},
									types.StringElement{
										Content: " <1>",
									},
								},
								{
									types.StringElement{
										Content: "and <more text> on the +",
									},
								},
								{
									types.StringElement{
										Content: "*next* lines with a link to ",
									},
									types.InlineLink{
										Location: types.Location{
											Scheme: "https://",
											Path: []interface{}{
												types.StringElement{
													Content: "github.com",
												},
											},
										},
									},
								},
								{},
								{
									types.StringElement{
										Content: "* not a list item",
									},
								},
							},
						},
						types.BlankLine{},
						types.CalloutListItem{
							Ref: 1,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a callout",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
			})

			It("should apply the 'specialchars' substitution", func() {
				s := strings.ReplaceAll(source, "$SUBS", "specialchars")
				expected := types.DraftDocument{
					Attributes: types.Attributes{
						"github-url": "https://github.com",
					},
					Elements: []interface{}{
						types.AttributeDeclaration{
							Name:  "github-url",
							Value: "https://github.com",
						},
						types.BlankLine{},
						types.ListingBlock{
							Attributes: types.Attributes{
								types.AttrSubstitutions: "specialchars",
							},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a link to https://example.com[] ",
									},
									types.SpecialCharacter{
										Name: "<",
									},
									types.StringElement{
										Content: "1",
									},
									types.SpecialCharacter{
										Name: ">",
									},
								},
								{
									types.StringElement{
										Content: "and ",
									},
									types.SpecialCharacter{
										Name: "<",
									},
									types.StringElement{
										Content: "more text",
									},
									types.SpecialCharacter{
										Name: ">",
									},
									types.StringElement{
										Content: " on the +",
									},
								},
								{
									types.StringElement{
										Content: "*next* lines with a link to {github-url}[]",
									},
								},
								{},
								{
									types.StringElement{
										Content: "* not a list item",
									},
								},
							},
						},
						types.BlankLine{},
						types.CalloutListItem{
							Ref: 1,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a callout",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
			})

			It("should apply the 'replacements' substitution", func() {
				s := strings.ReplaceAll(source, "$SUBS", "replacements")
				expected := types.DraftDocument{
					Attributes: types.Attributes{
						"github-url": "https://github.com",
					},
					Elements: []interface{}{
						types.AttributeDeclaration{
							Name:  "github-url",
							Value: "https://github.com",
						},
						types.BlankLine{},
						types.ListingBlock{
							Attributes: types.Attributes{
								types.AttrSubstitutions: "replacements",
							},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a link to https://example.com[] <1>",
									},
								},
								{
									types.StringElement{
										Content: "and <more text> on the +",
									},
								},
								{
									types.StringElement{
										Content: "*next* lines with a link to {github-url}[]",
									},
								},
								{},
								{
									types.StringElement{
										Content: "* not a list item",
									},
								},
							},
						},
						types.BlankLine{},
						types.CalloutListItem{
							Ref: 1,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a callout",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
			})

			It("should apply the 'post_replacements' substitution", func() {
				s := strings.ReplaceAll(source, "$SUBS", "post_replacements")
				expected := types.DraftDocument{
					Attributes: types.Attributes{
						"github-url": "https://github.com",
					},
					Elements: []interface{}{
						types.AttributeDeclaration{
							Name:  "github-url",
							Value: "https://github.com",
						},
						types.BlankLine{},
						types.ListingBlock{
							Attributes: types.Attributes{
								types.AttrSubstitutions: "post_replacements",
							},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a link to https://example.com[] <1>",
									},
								},
								{
									types.StringElement{
										Content: "and <more text> on the",
									},
									types.LineBreak{},
								},
								{
									types.StringElement{
										Content: "*next* lines with a link to {github-url}[]",
									},
								},
								{},
								{
									types.StringElement{
										Content: "* not a list item",
									},
								},
							},
						},
						types.BlankLine{},
						types.CalloutListItem{
							Ref: 1,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a callout",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
			})

			It("should apply the 'quotes,macros' substitution", func() {
				s := strings.ReplaceAll(source, "$SUBS", "quotes,macros")
				expected := types.DraftDocument{
					Attributes: types.Attributes{
						"github-url": "https://github.com",
					},
					Elements: []interface{}{
						types.AttributeDeclaration{
							Name:  "github-url",
							Value: "https://github.com",
						},
						types.BlankLine{},
						types.ListingBlock{
							Attributes: types.Attributes{
								types.AttrSubstitutions: "quotes,macros",
							},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a link to ",
									},
									types.InlineLink{
										Location: types.Location{
											Scheme: "https://",
											Path: []interface{}{
												types.StringElement{
													Content: "example.com",
												},
											},
										},
									},
									types.StringElement{
										Content: " <1>",
									},
								},
								{
									types.StringElement{
										Content: "and <more text> on the +",
									},
								},
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{
												Content: "next",
											},
										},
									},
									types.StringElement{
										Content: " lines with a link to {github-url}[]",
									},
								},
								{},
								{
									types.StringElement{
										Content: "* not a list item",
									},
								},
							},
						},
						types.BlankLine{},
						types.CalloutListItem{
							Ref: 1,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a callout",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
			})

			It("should apply the 'macros,quotes' substitution", func() {
				s := strings.ReplaceAll(source, "$SUBS", "macros,quotes")
				expected := types.DraftDocument{
					Attributes: types.Attributes{
						"github-url": "https://github.com",
					},
					Elements: []interface{}{
						types.AttributeDeclaration{
							Name:  "github-url",
							Value: "https://github.com",
						},
						types.BlankLine{},
						types.ListingBlock{
							Attributes: types.Attributes{
								types.AttrSubstitutions: "macros,quotes",
							},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a link to ",
									},
									types.InlineLink{
										Location: types.Location{
											Scheme: "https://",
											Path: []interface{}{
												types.StringElement{
													Content: "example.com",
												},
											},
										},
									},
									types.StringElement{
										Content: " <1>",
									},
								},
								{
									types.StringElement{
										Content: "and <more text> on the +",
									},
								},
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{
												Content: "next",
											},
										},
									},
									types.StringElement{
										Content: " lines with a link to {github-url}[]",
									},
								},
								{},
								{
									types.StringElement{
										Content: "* not a list item",
									},
								},
							},
						},
						types.BlankLine{},
						types.CalloutListItem{
							Ref: 1,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a callout",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
			})

			It("should apply the 'none' substitution", func() {
				s := strings.ReplaceAll(source, "$SUBS", "none")
				expected := types.DraftDocument{
					Attributes: types.Attributes{
						"github-url": "https://github.com",
					},
					Elements: []interface{}{
						types.AttributeDeclaration{
							Name:  "github-url",
							Value: "https://github.com",
						},
						types.BlankLine{},
						types.ListingBlock{
							Attributes: types.Attributes{
								types.AttrSubstitutions: "none",
							},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a link to https://example.com[] <1>",
									},
								},
								{
									types.StringElement{
										Content: "and <more text> on the +",
									},
								},
								{
									types.StringElement{
										Content: "*next* lines with a link to {github-url}[]",
									},
								},
								{},
								{
									types.StringElement{
										Content: "* not a list item",
									},
								},
							},
						},
						types.BlankLine{},
						types.CalloutListItem{
							Ref: 1,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a callout",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
			})

			It("should apply the 'quotes' substitutions on a passthrough block", func() {
				source := `[subs=quotes]
.a title
++++
_foo_

*bar*
++++`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.PassthroughBlock{
							Attributes: types.Attributes{
								types.AttrSubstitutions: "quotes",
								types.AttrTitle:         "a title",
							},
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Italic,
										Elements: []interface{}{
											types.StringElement{
												Content: "foo",
											},
										},
									},
								},
								{},
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
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
					Elements: []interface{}{
						types.ExampleBlock{
							Attributes: types.Attributes{
								types.AttrAdmonitionKind: types.Note,
							},
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
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
					Elements: []interface{}{
						types.ListingBlock{
							Attributes: types.Attributes{
								types.AttrAdmonitionKind: types.Note,
							},
							Lines: [][]interface{}{
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
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})
		})

		Context("markdown-style quote block", func() {

			It("with single marker without author", func() {
				source := `> some text
on *multiple lines*
with a link to https://example.com[]`

				expected := types.DraftDocument{
					Elements: []interface{}{
						types.MarkdownQuoteBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "some text",
									},
								},
								{
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
								{
									types.StringElement{
										Content: "with a link to ",
									},
									types.InlineLink{
										Location: types.Location{
											Scheme: "https://",
											Path: []interface{}{
												types.StringElement{
													Content: "example.com",
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

			It("with marker on each line without author", func() {
				source := `> some text
> on *multiple lines*`

				expected := types.DraftDocument{
					Elements: []interface{}{
						types.MarkdownQuoteBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "some text",
									},
								},
								{
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
					Elements: []interface{}{
						types.MarkdownQuoteBlock{
							Attributes: types.Attributes{
								types.AttrQuoteAuthor: "John Doe",
							},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "some text",
									},
								},
								{
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
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with marker on each line with author and title", func() {
				source := `.title
> some text
> on *multiple lines*
> -- John Doe`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.MarkdownQuoteBlock{
							Attributes: types.Attributes{
								types.AttrTitle:       "title",
								types.AttrQuoteAuthor: "John Doe",
							},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "some text",
									},
								},
								{
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
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with with author only", func() {
				source := `> -- John Doe`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.MarkdownQuoteBlock{
							Attributes: types.Attributes{
								types.AttrQuoteAuthor: "John Doe",
							},
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
					Elements: []interface{}{
						types.VerseBlock{
							Attributes: types.Attributes{
								types.AttrKind:        types.Verse,
								types.AttrQuoteAuthor: "john doe",
								types.AttrQuoteTitle:  "verse title",
							},
							Lines: [][]interface{}{
								{
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
				}
				result, err := ParseDraftDocument(source)
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(MatchDraftDocument(expected))
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
					Elements: []interface{}{
						types.VerseBlock{
							Attributes: types.Attributes{
								types.AttrKind:        types.Verse,
								types.AttrQuoteAuthor: "john doe",
							},
							Lines: [][]interface{}{
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
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("multi-line verse with title only", func() {
				source := `[verse, ,verse title]
____
some verse content 
____
`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.VerseBlock{
							Attributes: types.Attributes{
								types.AttrKind:       types.Verse,
								types.AttrQuoteTitle: "verse title",
							},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "some verse content ",
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
					Elements: []interface{}{
						types.VerseBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Verse,
							},
							Lines: [][]interface{}{
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
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("multi-line verse with unrendered list without author and title", func() {
				source := `[verse]
____
* foo


	* bar
____`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.VerseBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Verse,
							},
							Lines: [][]interface{}{
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
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("empty verse without author and title", func() {
				source := `[verse]
____
____`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.VerseBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Verse,
							},
							Lines: [][]interface{}{
								{},
							},
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
					Elements: []interface{}{
						types.VerseBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Verse,
							},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "foo",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			Context("with custom substitutions", func() {

				source := `:github-url: https://github.com
				
[subs="$SUBS"]
[verse, john doe, verse title]
____
a link to https://example.com[] <1>
and <more text> on the +
*next* lines with a link to {github-url}[]

* not a list item
____

<1> a callout
`

				It("should apply the default substitution", func() {
					s := strings.ReplaceAll(source, "[subs=\"$SUBS\"]\n", "")
					expected := types.DraftDocument{
						Attributes: types.Attributes{
							"github-url": "https://github.com",
						},
						Elements: []interface{}{
							types.AttributeDeclaration{
								Name:  "github-url",
								Value: "https://github.com",
							},
							types.BlankLine{},
							types.VerseBlock{
								Attributes: types.Attributes{
									types.AttrKind:        types.Verse,
									types.AttrQuoteAuthor: "john doe",
									types.AttrQuoteTitle:  "verse title",
								},
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "a link to ",
										},
										types.InlineLink{
											Location: types.Location{
												Scheme: "https://",
												Path: []interface{}{
													types.StringElement{
														Content: "example.com",
													},
												},
											},
										},
										types.StringElement{
											Content: " ",
										},
										types.SpecialCharacter{ // callout is not detected with the `normal` susbtitution
											Name: "<",
										},
										types.StringElement{
											Content: "1",
										},
										types.SpecialCharacter{
											Name: ">",
										},
									},
									{
										types.StringElement{
											Content: "and ",
										},
										types.SpecialCharacter{
											Name: "<",
										},
										types.StringElement{
											Content: "more text",
										},
										types.SpecialCharacter{
											Name: ">",
										},
										types.StringElement{
											Content: " on the",
										},
										types.LineBreak{},
									},
									{
										types.QuotedText{
											Kind: types.Bold,
											Elements: []interface{}{
												types.StringElement{
													Content: "next",
												},
											},
										},
										types.StringElement{
											Content: " lines with a link to ",
										},
										types.InlineLink{
											Location: types.Location{
												Scheme: "https://",
												Path: []interface{}{
													types.StringElement{
														Content: "github.com",
													},
												},
											},
										},
									},
									{},
									{
										types.StringElement{
											Content: "* not a list item",
										},
									},
								},
							},
							types.BlankLine{},
							types.CalloutListItem{
								Ref: 1,
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{
													Content: "a callout",
												},
											},
										},
									},
								},
							},
						},
					}
					Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
				})

				It("should apply the 'normal' substitution", func() {
					s := strings.ReplaceAll(source, "$SUBS", "normal")
					expected := types.DraftDocument{
						Attributes: types.Attributes{
							"github-url": "https://github.com",
						},
						Elements: []interface{}{
							types.AttributeDeclaration{
								Name:  "github-url",
								Value: "https://github.com",
							},
							types.BlankLine{},
							types.VerseBlock{
								Attributes: types.Attributes{
									types.AttrKind:          types.Verse,
									types.AttrQuoteAuthor:   "john doe",
									types.AttrQuoteTitle:    "verse title",
									types.AttrSubstitutions: "normal",
								},
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "a link to ",
										},
										types.InlineLink{
											Location: types.Location{
												Scheme: "https://",
												Path: []interface{}{
													types.StringElement{
														Content: "example.com",
													},
												},
											},
										},
										types.StringElement{
											Content: " ",
										},
										types.SpecialCharacter{ // callout is not detected with the `normal` susbtitution
											Name: "<",
										},
										types.StringElement{
											Content: "1",
										},
										types.SpecialCharacter{
											Name: ">",
										},
									},
									{
										types.StringElement{
											Content: "and ",
										},
										types.SpecialCharacter{
											Name: "<",
										},
										types.StringElement{
											Content: "more text",
										},
										types.SpecialCharacter{
											Name: ">",
										},
										types.StringElement{
											Content: " on the",
										},
										types.LineBreak{},
									},
									{
										types.QuotedText{
											Kind: types.Bold,
											Elements: []interface{}{
												types.StringElement{
													Content: "next",
												},
											},
										},
										types.StringElement{
											Content: " lines with a link to ",
										},
										types.InlineLink{
											Location: types.Location{
												Scheme: "https://",
												Path: []interface{}{
													types.StringElement{
														Content: "github.com",
													},
												},
											},
										},
									},
									{},
									{
										types.StringElement{
											Content: "* not a list item",
										},
									},
								},
							},
							types.BlankLine{},
							types.CalloutListItem{
								Ref: 1,
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{
													Content: "a callout",
												},
											},
										},
									},
								},
							},
						},
					}
					Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
				})

				It("should apply the 'quotes' substitution", func() {
					s := strings.ReplaceAll(source, "$SUBS", "quotes")
					expected := types.DraftDocument{
						Attributes: types.Attributes{
							"github-url": "https://github.com",
						},
						Elements: []interface{}{
							types.AttributeDeclaration{
								Name:  "github-url",
								Value: "https://github.com",
							},
							types.BlankLine{},
							types.VerseBlock{
								Attributes: types.Attributes{
									types.AttrKind:          types.Verse,
									types.AttrQuoteAuthor:   "john doe",
									types.AttrQuoteTitle:    "verse title",
									types.AttrSubstitutions: "quotes",
								},
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "a link to https://example.com[] <1>",
										},
									},
									{
										types.StringElement{
											Content: "and <more text> on the +",
										},
									},
									{
										types.QuotedText{
											Kind: types.Bold,
											Elements: []interface{}{
												types.StringElement{
													Content: "next",
												},
											},
										},
										types.StringElement{
											Content: " lines with a link to {github-url}[]",
										},
									},
									{},
									{
										types.StringElement{
											Content: "* not a list item",
										},
									},
								},
							},
							types.BlankLine{},
							types.CalloutListItem{
								Ref: 1,
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{
													Content: "a callout",
												},
											},
										},
									},
								},
							},
						},
					}
					Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
				})

				It("should apply the 'macros' substitution", func() {
					s := strings.ReplaceAll(source, "$SUBS", "macros")
					expected := types.DraftDocument{
						Attributes: types.Attributes{
							"github-url": "https://github.com",
						},
						Elements: []interface{}{
							types.AttributeDeclaration{
								Name:  "github-url",
								Value: "https://github.com",
							},
							types.BlankLine{},
							types.VerseBlock{
								Attributes: types.Attributes{
									types.AttrKind:          types.Verse,
									types.AttrQuoteAuthor:   "john doe",
									types.AttrQuoteTitle:    "verse title",
									types.AttrSubstitutions: "macros",
								},
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "a link to ",
										},
										types.InlineLink{
											Location: types.Location{
												Scheme: "https://",
												Path: []interface{}{
													types.StringElement{
														Content: "example.com",
													},
												},
											},
										},
										types.StringElement{
											Content: " <1>",
										},
									},
									{
										types.StringElement{
											Content: "and <more text> on the +",
										},
									},
									{
										types.StringElement{
											Content: "*next* lines with a link to {github-url}[]",
										},
									},
									{},
									{
										types.StringElement{
											Content: "* not a list item",
										},
									},
								},
							},
							types.BlankLine{},
							types.CalloutListItem{
								Ref: 1,
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{
													Content: "a callout",
												},
											},
										},
									},
								},
							},
						},
					}
					Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
				})

				It("should apply the 'attributes' substitution", func() {
					s := strings.ReplaceAll(source, "$SUBS", "attributes")
					expected := types.DraftDocument{
						Attributes: types.Attributes{
							"github-url": "https://github.com",
						},
						Elements: []interface{}{
							types.AttributeDeclaration{
								Name:  "github-url",
								Value: "https://github.com",
							},
							types.BlankLine{},
							types.VerseBlock{
								Attributes: types.Attributes{
									types.AttrKind:          types.Verse,
									types.AttrQuoteAuthor:   "john doe",
									types.AttrQuoteTitle:    "verse title",
									types.AttrSubstitutions: "attributes",
								},
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "a link to https://example.com[] <1>",
										},
									},
									{
										types.StringElement{
											Content: "and <more text> on the +",
										},
									},
									{
										types.StringElement{
											Content: "*next* lines with a link to https://github.com[]",
										},
									},
									{},
									{
										types.StringElement{
											Content: "* not a list item",
										},
									},
								},
							},
							types.BlankLine{},
							types.CalloutListItem{
								Ref: 1,
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{
													Content: "a callout",
												},
											},
										},
									},
								},
							},
						},
					}
					Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
				})

				It("should apply the 'attributes,macros' substitution", func() {
					s := strings.ReplaceAll(source, "$SUBS", "attributes,macros")
					expected := types.DraftDocument{
						Attributes: types.Attributes{
							"github-url": "https://github.com",
						},
						Elements: []interface{}{
							types.AttributeDeclaration{
								Name:  "github-url",
								Value: "https://github.com",
							},
							types.BlankLine{},
							types.VerseBlock{
								Attributes: types.Attributes{
									types.AttrKind:          types.Verse,
									types.AttrQuoteAuthor:   "john doe",
									types.AttrQuoteTitle:    "verse title",
									types.AttrSubstitutions: "attributes,macros",
								},
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "a link to ",
										},
										types.InlineLink{
											Location: types.Location{
												Scheme: "https://",
												Path: []interface{}{
													types.StringElement{
														Content: "example.com",
													},
												},
											},
										},
										types.StringElement{
											Content: " <1>",
										},
									},
									{
										types.StringElement{
											Content: "and <more text> on the +",
										},
									},
									{
										types.StringElement{
											Content: "*next* lines with a link to ",
										},
										types.InlineLink{
											Location: types.Location{
												Scheme: "https://",
												Path: []interface{}{
													types.StringElement{
														Content: "github.com",
													},
												},
											},
										},
									},
									{},
									{
										types.StringElement{
											Content: "* not a list item",
										},
									},
								},
							},
							types.BlankLine{},
							types.CalloutListItem{
								Ref: 1,
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{
													Content: "a callout",
												},
											},
										},
									},
								},
							},
						},
					}
					Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
				})

				It("should apply the 'specialchars' substitution", func() {
					s := strings.ReplaceAll(source, "$SUBS", "specialchars")
					expected := types.DraftDocument{
						Attributes: types.Attributes{
							"github-url": "https://github.com",
						},
						Elements: []interface{}{
							types.AttributeDeclaration{
								Name:  "github-url",
								Value: "https://github.com",
							},
							types.BlankLine{},
							types.VerseBlock{
								Attributes: types.Attributes{
									types.AttrKind:          types.Verse,
									types.AttrQuoteAuthor:   "john doe",
									types.AttrQuoteTitle:    "verse title",
									types.AttrSubstitutions: "specialchars",
								},
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "a link to https://example.com[] ",
										},
										types.SpecialCharacter{
											Name: "<",
										},
										types.StringElement{
											Content: "1",
										},
										types.SpecialCharacter{
											Name: ">",
										},
									},
									{
										types.StringElement{
											Content: "and ",
										},
										types.SpecialCharacter{
											Name: "<",
										},
										types.StringElement{
											Content: "more text",
										},
										types.SpecialCharacter{
											Name: ">",
										},
										types.StringElement{
											Content: " on the +",
										},
									},
									{
										types.StringElement{
											Content: "*next* lines with a link to {github-url}[]",
										},
									},
									{},
									{
										types.StringElement{
											Content: "* not a list item",
										},
									},
								},
							},
							types.BlankLine{},
							types.CalloutListItem{
								Ref: 1,
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{
													Content: "a callout",
												},
											},
										},
									},
								},
							},
						},
					}
					Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
				})

				It("should apply the 'replacements' substitution", func() {
					s := strings.ReplaceAll(source, "$SUBS", "replacements")
					expected := types.DraftDocument{
						Attributes: types.Attributes{
							"github-url": "https://github.com",
						},
						Elements: []interface{}{
							types.AttributeDeclaration{
								Name:  "github-url",
								Value: "https://github.com",
							},
							types.BlankLine{},
							types.VerseBlock{
								Attributes: types.Attributes{
									types.AttrKind:          types.Verse,
									types.AttrQuoteAuthor:   "john doe",
									types.AttrQuoteTitle:    "verse title",
									types.AttrSubstitutions: "replacements",
								},
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "a link to https://example.com[] <1>",
										},
									},
									{
										types.StringElement{
											Content: "and <more text> on the +",
										},
									},
									{
										types.StringElement{
											Content: "*next* lines with a link to {github-url}[]",
										},
									},
									{},
									{
										types.StringElement{
											Content: "* not a list item",
										},
									},
								},
							},
							types.BlankLine{},
							types.CalloutListItem{
								Ref: 1,
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{
													Content: "a callout",
												},
											},
										},
									},
								},
							},
						},
					}
					Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
				})

				It("should apply the 'post_replacements' substitution", func() {
					s := strings.ReplaceAll(source, "$SUBS", "post_replacements")
					expected := types.DraftDocument{
						Attributes: types.Attributes{
							"github-url": "https://github.com",
						},
						Elements: []interface{}{
							types.AttributeDeclaration{
								Name:  "github-url",
								Value: "https://github.com",
							},
							types.BlankLine{},
							types.VerseBlock{
								Attributes: types.Attributes{
									types.AttrKind:          types.Verse,
									types.AttrQuoteAuthor:   "john doe",
									types.AttrQuoteTitle:    "verse title",
									types.AttrSubstitutions: "post_replacements",
								},
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "a link to https://example.com[] <1>",
										},
									},
									{
										types.StringElement{
											Content: "and <more text> on the",
										},
										types.LineBreak{},
									},
									{
										types.StringElement{
											Content: "*next* lines with a link to {github-url}[]",
										},
									},
									{},
									{
										types.StringElement{
											Content: "* not a list item",
										},
									},
								},
							},
							types.BlankLine{},
							types.CalloutListItem{
								Ref: 1,
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{
													Content: "a callout",
												},
											},
										},
									},
								},
							},
						},
					}
					Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
				})

				It("should apply the 'quotes,macros' substitution", func() {
					s := strings.ReplaceAll(source, "$SUBS", "quotes,macros")
					expected := types.DraftDocument{
						Attributes: types.Attributes{
							"github-url": "https://github.com",
						},
						Elements: []interface{}{
							types.AttributeDeclaration{
								Name:  "github-url",
								Value: "https://github.com",
							},
							types.BlankLine{},
							types.VerseBlock{
								Attributes: types.Attributes{
									types.AttrKind:          types.Verse,
									types.AttrQuoteAuthor:   "john doe",
									types.AttrQuoteTitle:    "verse title",
									types.AttrSubstitutions: "quotes,macros",
								},
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "a link to ",
										},
										types.InlineLink{
											Location: types.Location{
												Scheme: "https://",
												Path: []interface{}{
													types.StringElement{
														Content: "example.com",
													},
												},
											},
										},
										types.StringElement{
											Content: " <1>",
										},
									},
									{
										types.StringElement{
											Content: "and <more text> on the +",
										},
									},
									{
										types.QuotedText{
											Kind: types.Bold,
											Elements: []interface{}{
												types.StringElement{
													Content: "next",
												},
											},
										},
										types.StringElement{
											Content: " lines with a link to {github-url}[]",
										},
									},
									{},
									{
										types.StringElement{
											Content: "* not a list item",
										},
									},
								},
							},
							types.BlankLine{},
							types.CalloutListItem{
								Ref: 1,
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{
													Content: "a callout",
												},
											},
										},
									},
								},
							},
						},
					}
					Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
				})

				It("should apply the 'macros,quotes' substitution", func() {
					s := strings.ReplaceAll(source, "$SUBS", "macros,quotes")
					expected := types.DraftDocument{
						Attributes: types.Attributes{
							"github-url": "https://github.com",
						},
						Elements: []interface{}{
							types.AttributeDeclaration{
								Name:  "github-url",
								Value: "https://github.com",
							},
							types.BlankLine{},
							types.VerseBlock{
								Attributes: types.Attributes{
									types.AttrKind:          types.Verse,
									types.AttrQuoteAuthor:   "john doe",
									types.AttrQuoteTitle:    "verse title",
									types.AttrSubstitutions: "macros,quotes",
								},
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "a link to ",
										},
										types.InlineLink{
											Location: types.Location{
												Scheme: "https://",
												Path: []interface{}{
													types.StringElement{
														Content: "example.com",
													},
												},
											},
										},
										types.StringElement{
											Content: " <1>",
										},
									},
									{
										types.StringElement{
											Content: "and <more text> on the +",
										},
									},
									{
										types.QuotedText{
											Kind: types.Bold,
											Elements: []interface{}{
												types.StringElement{
													Content: "next",
												},
											},
										},
										types.StringElement{
											Content: " lines with a link to {github-url}[]",
										},
									},
									{},
									{
										types.StringElement{
											Content: "* not a list item",
										},
									},
								},
							},
							types.BlankLine{},
							types.CalloutListItem{
								Ref: 1,
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{
													Content: "a callout",
												},
											},
										},
									},
								},
							},
						},
					}
					Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
				})

				It("should apply the 'none' substitution", func() {
					s := strings.ReplaceAll(source, "$SUBS", "none")
					expected := types.DraftDocument{
						Attributes: types.Attributes{
							"github-url": "https://github.com",
						},
						Elements: []interface{}{
							types.AttributeDeclaration{
								Name:  "github-url",
								Value: "https://github.com",
							},
							types.BlankLine{},
							types.VerseBlock{
								Attributes: types.Attributes{
									types.AttrKind:          types.Verse,
									types.AttrQuoteAuthor:   "john doe",
									types.AttrQuoteTitle:    "verse title",
									types.AttrSubstitutions: "none",
								},
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "a link to https://example.com[] <1>",
										},
									},
									{
										types.StringElement{
											Content: "and <more text> on the +",
										},
									},
									{
										types.StringElement{
											Content: "*next* lines with a link to {github-url}[]",
										},
									},
									{},
									{
										types.StringElement{
											Content: "* not a list item",
										},
									},
								},
							},
							types.BlankLine{},
							types.CalloutListItem{
								Ref: 1,
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{
													Content: "a callout",
												},
											},
										},
									},
								},
							},
						},
					}
					Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
				})
			})
		})

	})

	Context("final document", func() {

		Context("example block", func() {

			It("with single line", func() {
				source := `====
some listing code
====`
				expected := types.Document{
					Elements: []interface{}{
						types.ExampleBlock{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
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
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with single line starting with a dot", func() {
				source := `====
.foo
====`
				expected := types.Document{
					Elements: []interface{}{
						types.ExampleBlock{
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
						types.ExampleBlock{
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.Attributes{
										types.AttrTitle: "foo",
									},
									Lines: [][]interface{}{
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
									Items: []types.UnorderedListItem{
										{
											Level:       1,
											BulletStyle: types.OneAsterisk,
											CheckStyle:  types.NoCheck,
											Elements: []interface{}{
												types.Paragraph{
													Lines: [][]interface{}{
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
						types.ExampleBlock{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
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
						types.ExampleBlock{
							Attributes: types.Attributes{
								types.AttrTitle: "example block title",
							},
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
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
						types.ExampleBlock{
							Elements: []interface{}{},
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
						types.QuoteBlock{
							Attributes: types.Attributes{
								types.AttrKind:        types.Quote,
								types.AttrQuoteAuthor: "john doe",
								types.AttrQuoteTitle:  "quote title",
							},
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
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
						types.QuoteBlock{
							Attributes: types.Attributes{
								types.AttrKind:        types.Quote,
								types.AttrQuoteAuthor: "john doe",
							},
							Elements: []interface{}{
								types.UnorderedList{
									Items: []types.UnorderedListItem{
										{
											Level:       1,
											BulletStyle: types.Dash,
											CheckStyle:  types.NoCheck,
											Elements: []interface{}{
												types.Paragraph{
													Lines: [][]interface{}{
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
											Level:       1,
											BulletStyle: types.Dash,
											CheckStyle:  types.NoCheck,
											Elements: []interface{}{
												types.Paragraph{
													Lines: [][]interface{}{
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
											Level:       1,
											BulletStyle: types.Dash,
											CheckStyle:  types.NoCheck,
											Elements: []interface{}{
												types.Paragraph{
													Lines: [][]interface{}{
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
						types.QuoteBlock{
							Attributes: types.Attributes{
								types.AttrKind:       types.Quote,
								types.AttrQuoteTitle: "quote title",
							},
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
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

			It("with single line starting with a dot", func() {
				source := `[quote]
____
.foo
____`
				expected := types.Document{
					Elements: []interface{}{
						types.QuoteBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Quote,
							},
							Elements: []interface{}{},
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
						types.QuoteBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Quote,
							},
							Elements: []interface{}{
								types.UnorderedList{
									Items: []types.UnorderedListItem{
										{
											Level:       1,
											BulletStyle: types.OneAsterisk,
											CheckStyle:  types.NoCheck,
											Elements: []interface{}{
												types.Paragraph{
													Lines: [][]interface{}{
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
								types.ListingBlock{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "* quote ",
											},
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
													Lines: [][]interface{}{
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
						types.QuoteBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Quote,
							},
							Elements: []interface{}{
								types.UnorderedList{
									Items: []types.UnorderedListItem{
										{
											Level:       1,
											BulletStyle: types.OneAsterisk,
											CheckStyle:  types.NoCheck,
											Elements: []interface{}{
												types.Paragraph{
													Lines: [][]interface{}{
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
											Level:       1,
											BulletStyle: types.OneAsterisk,
											CheckStyle:  types.NoCheck,
											Elements: []interface{}{
												types.Paragraph{
													Lines: [][]interface{}{
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
											Level:       1,
											BulletStyle: types.OneAsterisk,
											CheckStyle:  types.NoCheck,
											Elements: []interface{}{
												types.Paragraph{
													Lines: [][]interface{}{
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
						types.QuoteBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Quote,
							},
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
						types.QuoteBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Quote,
							},
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
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

		Context("sidebar block", func() {

			It("with paragraph", func() {
				source := `****
some *verse* content
****`
				expected := types.Document{
					Elements: []interface{}{
						types.SidebarBlock{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
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
						types.SidebarBlock{
							Attributes: types.Attributes{
								types.AttrTitle: "a title",
							},
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
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
								types.ListingBlock{
									Lines: [][]interface{}{
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
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with single line starting with a dot", func() {
				source := `
****
.foo
****`
				expected := types.Document{
					Elements: []interface{}{
						types.SidebarBlock{
							Elements: []interface{}{},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

		})

		Context("fenced block", func() {

			It("with single line", func() {
				content := "some fenced code"
				source := "```\n" + content + "\n" + "```"
				expected := types.Document{
					Elements: []interface{}{
						types.FencedBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: content,
									},
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
						types.FencedBlock{
							Lines: [][]interface{}{
								{},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with multiple lines alone", func() {
				source := "```\nsome fenced code\nwith an empty line\n\nin the middle\n```"
				expected := types.Document{
					Elements: []interface{}{
						types.FencedBlock{
							Lines: [][]interface{}{
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
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with multiple lines then a paragraph", func() {
				source := "```\nsome fenced code\nwith an empty line\n\nin the middle\n```\nthen a normal paragraph."
				expected := types.Document{
					Elements: []interface{}{
						types.FencedBlock{
							Lines: [][]interface{}{
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
								{},
								{
									types.StringElement{
										Content: "in the middle",
									},
								},
							},
						},
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "then a normal paragraph."},
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
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a paragraph.",
									},
								},
							},
						},
						types.FencedBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: content,
									},
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
						types.FencedBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "End of file here",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with external link inside - without attributes", func() {
				source := "```\n" +
					"a https://example.com\n" +
					"and more text on the\n" +
					"next lines\n" +
					"```"
				expected := types.Document{
					Elements: []interface{}{
						types.FencedBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a https://example.com",
									},
								},
								{
									types.StringElement{
										Content: "and more text on the",
									},
								},
								{
									types.StringElement{
										Content: "next lines",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with external link inside - with attributes", func() {
				source := "```" + "\n" +
					"a https://example.com[]" + "\n" +
					"and more text on the" + "\n" +
					"next lines" + "\n" +
					"```"
				expected := types.Document{
					Elements: []interface{}{
						types.FencedBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a https://example.com[]",
									},
								},
								{
									types.StringElement{
										Content: "and more text on the",
									},
								},
								{
									types.StringElement{
										Content: "next lines",
									},
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
						types.FencedBlock{
							Lines: [][]interface{}{
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
										Content: "* content ",
									},
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
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "some listing code",
									},
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
						types.ListingBlock{
							Lines: [][]interface{}{
								{},
							},
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
						types.ListingBlock{
							Lines: [][]interface{}{
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
						types.ListingBlock{
							Lines: [][]interface{}{
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
						types.ListingBlock{
							Lines: [][]interface{}{
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
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "then a normal paragraph."},
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
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a paragraph.",
									},
								},
							},
						},
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "some listing code",
									},
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
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "End of file here.",
									},
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
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "import ",
									},
									types.Callout{
										Ref: 1,
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
											Lines: [][]interface{}{
												{
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
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "import ",
									},
									types.Callout{
										Ref: 1,
									},
								},
								{},
								{
									types.StringElement{
										Content: "func foo() {} ",
									},
									types.Callout{
										Ref: 2,
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
											Lines: [][]interface{}{
												{
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
											Lines: [][]interface{}{
												{
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
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "import ",
									},
									types.Callout{
										Ref: 1,
									},
									types.Callout{
										Ref: 2,
									},
									types.Callout{
										Ref: 3,
									},
								},
								{},
								{
									types.StringElement{
										Content: "func foo() {} ",
									},
									types.Callout{
										Ref: 4,
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
											Lines: [][]interface{}{
												{
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
											Lines: [][]interface{}{
												{
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
											Lines: [][]interface{}{
												{
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
											Lines: [][]interface{}{
												{
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
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "import ",
									},
									types.SpecialCharacter{
										Name: "<",
									},
									types.StringElement{
										Content: "a",
									},
									types.SpecialCharacter{
										Name: ">",
									},
								},
							},
						},
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.SpecialCharacter{
										Name: "<",
									},
									types.StringElement{
										Content: "a",
									},
									types.SpecialCharacter{
										Name: ">",
									},
									types.StringElement{
										Content: " an import",
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
						types.ListingBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Source,
							},
							Lines: [][]interface{}{
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
						types.ListingBlock{
							Attributes: types.Attributes{
								types.AttrKind:     types.Source,
								types.AttrLanguage: "ruby",
								types.AttrTitle:    "Source block title",
							},
							Lines: [][]interface{}{
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
						types.ListingBlock{
							Attributes: types.Attributes{
								types.AttrKind:     types.Source,
								types.AttrLanguage: "ruby",
								types.AttrID:       "id-for-source-block",
								types.AttrCustomID: true,
								types.AttrTitle:    "app.rb",
							},
							Lines: [][]interface{}{
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
						types.PassthroughBlock{
							Attributes: types.Attributes{
								types.AttrTitle: "a title",
							},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "_foo_",
									},
								},
								{},
								{
									types.StringElement{
										Content: "*bar*",
									},
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
						types.PassthroughBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Passthrough,
							},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "_foo_",
									},
								},
								{
									types.StringElement{
										Content: "*bar*",
									},
								},
							},
						},
						types.Paragraph{
							Lines: [][]interface{}{
								{
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

		Context("admonition block", func() {

			It("example block as admonition", func() {
				source := `[NOTE]
====
foo
====`
				expected := types.Document{
					Elements: []interface{}{
						types.ExampleBlock{
							Attributes: types.Attributes{
								types.AttrAdmonitionKind: types.Note,
							},
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
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
						types.ExampleBlock{
							Attributes: types.Attributes{
								types.AttrAdmonitionKind: types.Note,
							},
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "multiple",
											},
										},
									},
								},
								types.BlankLine{},
								types.Paragraph{
									Lines: [][]interface{}{
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
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})

		Context("markdown-style quote block", func() {

			It("with single marker without author", func() {
				source := `> some text
on *multiple lines*`

				expected := types.Document{
					Elements: []interface{}{
						types.MarkdownQuoteBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "some text",
									},
								},
								{
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
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with marker on each line without author", func() {
				source := `> some text
> on *multiple lines*`
				expected := types.Document{
					Elements: []interface{}{
						types.MarkdownQuoteBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "some text",
									},
								},
								{
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
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with marker on each line with author only", func() {
				source := `> some text
> on *multiple lines*
> -- John Doe`
				expected := types.Document{
					Elements: []interface{}{
						types.MarkdownQuoteBlock{
							Attributes: types.Attributes{
								types.AttrQuoteAuthor: "John Doe",
							},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "some text",
									},
								},
								{
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
						types.MarkdownQuoteBlock{
							Attributes: types.Attributes{
								types.AttrTitle:       "title",
								types.AttrQuoteAuthor: "John Doe",
							},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "some text",
									},
								},
								{
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
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with with author only", func() {
				source := `> -- John Doe`
				expected := types.Document{
					Elements: []interface{}{
						types.MarkdownQuoteBlock{
							Attributes: types.Attributes{
								types.AttrQuoteAuthor: "John Doe",
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
						types.VerseBlock{
							Attributes: types.Attributes{
								types.AttrKind:        types.Verse,
								types.AttrQuoteAuthor: "john doe",
								types.AttrQuoteTitle:  "verse title",
							},
							Lines: [][]interface{}{
								{
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
						types.VerseBlock{
							Attributes: types.Attributes{
								types.AttrKind:        types.Verse,
								types.AttrQuoteAuthor: "john doe",
							},
							Lines: [][]interface{}{
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
						types.VerseBlock{
							Attributes: types.Attributes{
								types.AttrKind:       types.Verse,
								types.AttrQuoteTitle: "verse title",
							},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "some verse content ",
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
						types.VerseBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Verse,
							},
							Lines: [][]interface{}{
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
						types.VerseBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Verse,
							},
							Lines: [][]interface{}{
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
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("empty verse without author and title", func() {
				source := `[verse]
____
____`
				expected := types.Document{
					Elements: []interface{}{
						types.VerseBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Verse,
							},
							Lines: [][]interface{}{
								{},
							},
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
						types.VerseBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Verse,
							},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "foo",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})
	})
})
