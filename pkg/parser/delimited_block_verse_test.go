package parser_test

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2" // nolint:golint
	. "github.com/onsi/gomega"    // nolint:golint
)

var _ = Describe("verse blocks", func() {

	Context("in final documents", func() {

		Context("as delimited blocks", func() {

			It("single line verse with author and title", func() {
				source := `[verse, john doe, verse title]
____
some *verse* content
____`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Verse,
							Attributes: types.Attributes{
								types.AttrStyle:       types.Verse,
								types.AttrQuoteAuthor: "john doe",
								types.AttrQuoteTitle:  "verse title",
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "some ",
								},
								&types.QuotedText{
									Kind: types.SingleQuoteBold,
									Elements: []interface{}{
										&types.StringElement{
											Content: "verse",
										},
									},
								},
								&types.StringElement{
									Content: " content",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("single line verse with author and title and empty lines around", func() {
				source := `[verse, john doe, verse title]
____

some *verse* content

____`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Verse,
							Attributes: types.Attributes{
								types.AttrStyle:       types.Verse,
								types.AttrQuoteAuthor: "john doe",
								types.AttrQuoteTitle:  "verse title",
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "\nsome ",
								},
								&types.QuotedText{
									Kind: types.SingleQuoteBold,
									Elements: []interface{}{
										&types.StringElement{
											Content: "verse",
										},
									},
								},
								&types.StringElement{
									Content: " content\n",
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
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Verse,
							Attributes: types.Attributes{
								types.AttrStyle:       types.Verse,
								types.AttrQuoteAuthor: "john doe",
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "- some\n- verse\n- content", // suffix spaces are trimmed
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
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Verse,
							Attributes: types.Attributes{
								types.AttrStyle:      types.Verse,
								types.AttrQuoteTitle: "verse title",
							},
							Elements: []interface{}{
								&types.StringElement{
									// suffix spaces are trimmed
									Content: "some verse content",
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
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Verse,
							Attributes: types.Attributes{
								types.AttrStyle: types.Verse,
							},
							Elements: []interface{}{
								&types.StringElement{
									// suffix spaces are trimmed on each line
									Content: "* some\n----\n* verse\n----\n* content",
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
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Verse,
							Attributes: types.Attributes{
								types.AttrStyle: types.Verse,
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "* foo\n\n\n\t* bar",
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
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Verse,
							Attributes: types.Attributes{
								types.AttrStyle: types.Verse,
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
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Verse,
							Attributes: types.Attributes{
								types.AttrStyle: types.Verse,
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "foo",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
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
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name:  "github-url",
										Value: "https://github.com",
									},
								},
							},
							&types.DelimitedBlock{
								Kind: types.Verse,
								Attributes: types.Attributes{
									types.AttrStyle:       types.Verse,
									types.AttrQuoteAuthor: "john doe",
									types.AttrQuoteTitle:  "verse title",
								},
								Elements: []interface{}{
									&types.StringElement{
										Content: "a link to ",
									},
									&types.InlineLink{
										Location: &types.Location{
											Scheme: "https://",
											Path:   "example.com",
										},
									},
									&types.StringElement{
										Content: " ",
									},
									&types.SpecialCharacter{ // callout is not detected with the `normal` susbtitution
										Name: "<",
									},
									&types.StringElement{
										Content: "1",
									},
									&types.SpecialCharacter{
										Name: ">",
									},
									&types.StringElement{
										Content: "\nand ",
									},
									&types.SpecialCharacter{
										Name: "<",
									},
									&types.StringElement{
										Content: "more text",
									},
									&types.SpecialCharacter{
										Name: ">",
									},
									&types.StringElement{
										Content: " on the",
									},
									&types.LineBreak{},
									&types.StringElement{
										Content: "\n",
									},
									&types.QuotedText{
										Kind: types.SingleQuoteBold,
										Elements: []interface{}{
											&types.StringElement{
												Content: "next",
											},
										},
									},
									&types.StringElement{
										Content: " lines with a link to ",
									},
									&types.InlineLink{
										Location: &types.Location{
											Scheme: "https://",
											Path:   "github.com",
										},
									},
									&types.StringElement{
										Content: "\n\n* not a list item",
									},
								},
							},
							&types.List{
								Kind: types.CalloutListKind,
								Elements: []types.ListElement{
									&types.CalloutListElement{
										Ref: 1,
										Elements: []interface{}{
											&types.Paragraph{
												Elements: []interface{}{
													&types.StringElement{
														Content: "a callout",
													},
												},
											},
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(s)).To(MatchDocument(expected))
				})

				It("should apply the 'normal' substitution", func() {
					s := strings.ReplaceAll(source, "$SUBS", "normal")
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name:  "github-url",
										Value: "https://github.com",
									},
								},
							},
							&types.DelimitedBlock{
								Kind: types.Verse,
								Attributes: types.Attributes{
									types.AttrStyle:         types.Verse,
									types.AttrQuoteAuthor:   "john doe",
									types.AttrQuoteTitle:    "verse title",
									types.AttrSubstitutions: "normal",
								},
								Elements: []interface{}{
									&types.StringElement{
										Content: "a link to ",
									},
									&types.InlineLink{
										Location: &types.Location{
											Scheme: "https://",
											Path:   "example.com",
										},
									},
									&types.StringElement{
										Content: " ",
									},
									&types.SpecialCharacter{ // callout is not detected with the `normal` susbtitution
										Name: "<",
									},
									&types.StringElement{
										Content: "1",
									},
									&types.SpecialCharacter{
										Name: ">",
									},
									&types.StringElement{
										Content: "\nand ",
									},
									&types.SpecialCharacter{
										Name: "<",
									},
									&types.StringElement{
										Content: "more text",
									},
									&types.SpecialCharacter{
										Name: ">",
									},
									&types.StringElement{
										Content: " on the",
									},
									&types.LineBreak{},
									&types.StringElement{
										Content: "\n",
									},
									&types.QuotedText{
										Kind: types.SingleQuoteBold,
										Elements: []interface{}{
											&types.StringElement{
												Content: "next",
											},
										},
									},
									&types.StringElement{
										Content: " lines with a link to ",
									},
									&types.InlineLink{
										Location: &types.Location{
											Scheme: "https://",
											Path:   "github.com",
										},
									},
									&types.StringElement{
										Content: "\n\n* not a list item",
									},
								},
							},
							&types.List{
								Kind: types.CalloutListKind,
								Elements: []types.ListElement{
									&types.CalloutListElement{
										Ref: 1,
										Elements: []interface{}{
											&types.Paragraph{
												Elements: []interface{}{
													&types.StringElement{
														Content: "a callout",
													},
												},
											},
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(s)).To(MatchDocument(expected))
				})

				It("should apply the 'quotes' substitution", func() {
					s := strings.ReplaceAll(source, "$SUBS", "quotes")
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name:  "github-url",
										Value: "https://github.com",
									},
								},
							},
							&types.DelimitedBlock{
								Kind: types.Verse,
								Attributes: types.Attributes{
									types.AttrStyle:         types.Verse,
									types.AttrQuoteAuthor:   "john doe",
									types.AttrQuoteTitle:    "verse title",
									types.AttrSubstitutions: "quotes",
								},
								Elements: []interface{}{
									&types.StringElement{
										Content: "a link to https://example.com[] <1>\nand <more text> on the +\n",
									},
									&types.QuotedText{
										Kind: types.SingleQuoteBold,
										Elements: []interface{}{
											&types.StringElement{
												Content: "next",
											},
										},
									},
									&types.StringElement{
										Content: " lines with a link to {github-url}[]\n\n* not a list item",
									},
								},
							},
							&types.List{
								Kind: types.CalloutListKind,
								Elements: []types.ListElement{
									&types.CalloutListElement{
										Ref: 1,
										Elements: []interface{}{
											&types.Paragraph{
												Elements: []interface{}{
													&types.StringElement{
														Content: "a callout",
													},
												},
											},
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(s)).To(MatchDocument(expected))
				})

				It("should apply the 'macros' substitution", func() {
					s := strings.ReplaceAll(source, "$SUBS", "macros")
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name:  "github-url",
										Value: "https://github.com",
									},
								},
							},
							&types.DelimitedBlock{
								Kind: types.Verse,
								Attributes: types.Attributes{
									types.AttrStyle:         types.Verse,
									types.AttrQuoteAuthor:   "john doe",
									types.AttrQuoteTitle:    "verse title",
									types.AttrSubstitutions: "macros",
								},
								Elements: []interface{}{
									&types.StringElement{
										Content: "a link to ",
									},
									&types.InlineLink{
										Location: &types.Location{
											Scheme: "https://",
											Path:   "example.com",
										},
									},
									&types.StringElement{
										Content: " <1>\nand <more text> on the +\n*next* lines with a link to {github-url}[]\n\n* not a list item",
									},
								},
							},
							&types.List{
								Kind: types.CalloutListKind,
								Elements: []types.ListElement{
									&types.CalloutListElement{
										Ref: 1,
										Elements: []interface{}{
											&types.Paragraph{
												Elements: []interface{}{
													&types.StringElement{
														Content: "a callout",
													},
												},
											},
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(s)).To(MatchDocument(expected))
				})

				It("should apply the 'attributes' substitution", func() {
					s := strings.ReplaceAll(source, "$SUBS", "attributes")
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name:  "github-url",
										Value: "https://github.com",
									},
								},
							},
							&types.DelimitedBlock{
								Kind: types.Verse,
								Attributes: types.Attributes{
									types.AttrStyle:         types.Verse,
									types.AttrQuoteAuthor:   "john doe",
									types.AttrQuoteTitle:    "verse title",
									types.AttrSubstitutions: "attributes",
								},
								Elements: []interface{}{
									&types.StringElement{
										Content: "a link to https://example.com[] <1>\nand <more text> on the +\n*next* lines with a link to https://github.com[]\n\n* not a list item",
									},
								},
							},
							&types.List{
								Kind: types.CalloutListKind,
								Elements: []types.ListElement{
									&types.CalloutListElement{
										Ref: 1,
										Elements: []interface{}{
											&types.Paragraph{
												Elements: []interface{}{
													&types.StringElement{
														Content: "a callout",
													},
												},
											},
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(s)).To(MatchDocument(expected))
				})

				It("should apply the 'attributes,macros' substitution", func() {
					s := strings.ReplaceAll(source, "$SUBS", "attributes,macros")
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name:  "github-url",
										Value: "https://github.com",
									},
								},
							},
							&types.DelimitedBlock{
								Kind: types.Verse,
								Attributes: types.Attributes{
									types.AttrStyle:         types.Verse,
									types.AttrQuoteAuthor:   "john doe",
									types.AttrQuoteTitle:    "verse title",
									types.AttrSubstitutions: "attributes,macros",
								},
								Elements: []interface{}{
									&types.StringElement{
										Content: "a link to ",
									},
									&types.InlineLink{
										Location: &types.Location{
											Scheme: "https://",
											Path:   "example.com",
										},
									},
									&types.StringElement{
										Content: " <1>\nand <more text> on the +\n*next* lines with a link to ",
									},
									&types.InlineLink{
										Location: &types.Location{
											Scheme: "https://",
											Path:   "github.com",
										},
									},
									&types.StringElement{
										Content: "\n\n* not a list item",
									},
								},
							},
							&types.List{
								Kind: types.CalloutListKind,
								Elements: []types.ListElement{
									&types.CalloutListElement{
										Ref: 1,
										Elements: []interface{}{
											&types.Paragraph{
												Elements: []interface{}{
													&types.StringElement{
														Content: "a callout",
													},
												},
											},
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(s)).To(MatchDocument(expected))
				})

				It("should apply the 'specialchars' substitution", func() {
					s := strings.ReplaceAll(source, "$SUBS", "specialchars")
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name:  "github-url",
										Value: "https://github.com",
									},
								},
							},
							&types.DelimitedBlock{
								Kind: types.Verse,
								Attributes: types.Attributes{
									types.AttrStyle:         types.Verse,
									types.AttrQuoteAuthor:   "john doe",
									types.AttrQuoteTitle:    "verse title",
									types.AttrSubstitutions: "specialchars",
								},
								Elements: []interface{}{
									&types.StringElement{
										Content: "a link to https://example.com[] ",
									},
									&types.SpecialCharacter{
										Name: "<",
									},
									&types.StringElement{
										Content: "1",
									},
									&types.SpecialCharacter{
										Name: ">",
									},
									&types.StringElement{
										Content: "\nand ",
									},
									&types.SpecialCharacter{
										Name: "<",
									},
									&types.StringElement{
										Content: "more text",
									},
									&types.SpecialCharacter{
										Name: ">",
									},
									&types.StringElement{
										Content: " on the +\n*next* lines with a link to {github-url}[]\n\n* not a list item",
									},
								},
							},
							&types.List{
								Kind: types.CalloutListKind,
								Elements: []types.ListElement{
									&types.CalloutListElement{
										Ref: 1,
										Elements: []interface{}{
											&types.Paragraph{
												Elements: []interface{}{
													&types.StringElement{
														Content: "a callout",
													},
												},
											},
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(s)).To(MatchDocument(expected))
				})

				It("should apply the 'replacements' substitution", func() {
					s := strings.ReplaceAll(source, "$SUBS", "replacements")
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name:  "github-url",
										Value: "https://github.com",
									},
								},
							},
							&types.DelimitedBlock{
								Kind: types.Verse,
								Attributes: types.Attributes{
									types.AttrStyle:         types.Verse,
									types.AttrQuoteAuthor:   "john doe",
									types.AttrQuoteTitle:    "verse title",
									types.AttrSubstitutions: "replacements",
								},
								Elements: []interface{}{
									&types.StringElement{
										Content: "a link to https://example.com[] <1>\nand <more text> on the +\n*next* lines with a link to {github-url}[]\n\n* not a list item",
									},
								},
							},
							&types.List{
								Kind: types.CalloutListKind,
								Elements: []types.ListElement{
									&types.CalloutListElement{
										Ref: 1,
										Elements: []interface{}{
											&types.Paragraph{
												Elements: []interface{}{
													&types.StringElement{
														Content: "a callout",
													},
												},
											},
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(s)).To(MatchDocument(expected))
				})

				It("should apply the 'post_replacements' substitution", func() {
					s := strings.ReplaceAll(source, "$SUBS", "post_replacements")
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name:  "github-url",
										Value: "https://github.com",
									},
								},
							},
							&types.DelimitedBlock{
								Kind: types.Verse,
								Attributes: types.Attributes{
									types.AttrStyle:         types.Verse,
									types.AttrQuoteAuthor:   "john doe",
									types.AttrQuoteTitle:    "verse title",
									types.AttrSubstitutions: "post_replacements",
								},
								Elements: []interface{}{
									&types.StringElement{
										Content: "a link to https://example.com[] <1>\nand <more text> on the",
									},
									&types.LineBreak{},
									&types.StringElement{
										Content: "\n*next* lines with a link to {github-url}[]\n\n* not a list item",
									},
								},
							},
							&types.List{
								Kind: types.CalloutListKind,
								Elements: []types.ListElement{
									&types.CalloutListElement{
										Ref: 1,
										Elements: []interface{}{
											&types.Paragraph{
												Elements: []interface{}{
													&types.StringElement{
														Content: "a callout",
													},
												},
											},
										},
									},
								},
							},
						},
					}
					doc, err := ParseDocument(s)
					Expect(err).NotTo(HaveOccurred())
					Expect(doc).To(MatchDocument(expected))
				})

				It("should apply the 'quotes,macros' substitution", func() {
					s := strings.ReplaceAll(source, "$SUBS", "quotes,macros")
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name:  "github-url",
										Value: "https://github.com",
									},
								},
							},
							&types.DelimitedBlock{
								Kind: types.Verse,
								Attributes: types.Attributes{
									types.AttrStyle:         types.Verse,
									types.AttrQuoteAuthor:   "john doe",
									types.AttrQuoteTitle:    "verse title",
									types.AttrSubstitutions: "quotes,macros",
								},
								Elements: []interface{}{
									&types.StringElement{
										Content: "a link to ",
									},
									&types.InlineLink{
										Location: &types.Location{
											Scheme: "https://",
											Path:   "example.com",
										},
									},
									&types.StringElement{
										Content: " <1>\nand <more text> on the +\n",
									},
									&types.QuotedText{
										Kind: types.SingleQuoteBold,
										Elements: []interface{}{
											&types.StringElement{
												Content: "next",
											},
										},
									},
									&types.StringElement{
										Content: " lines with a link to {github-url}[]\n\n* not a list item",
									},
								},
							},
							&types.List{
								Kind: types.CalloutListKind,
								Elements: []types.ListElement{
									&types.CalloutListElement{
										Ref: 1,
										Elements: []interface{}{
											&types.Paragraph{
												Elements: []interface{}{
													&types.StringElement{
														Content: "a callout",
													},
												},
											},
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(s)).To(MatchDocument(expected))
				})

				It("should apply the 'macros,quotes' substitution", func() {
					s := strings.ReplaceAll(source, "$SUBS", "macros,quotes")
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name:  "github-url",
										Value: "https://github.com",
									},
								},
							},
							&types.DelimitedBlock{
								Kind: types.Verse,
								Attributes: types.Attributes{
									types.AttrStyle:         types.Verse,
									types.AttrQuoteAuthor:   "john doe",
									types.AttrQuoteTitle:    "verse title",
									types.AttrSubstitutions: "macros,quotes",
								},
								Elements: []interface{}{
									&types.StringElement{
										Content: "a link to ",
									},
									&types.InlineLink{
										Location: &types.Location{
											Scheme: "https://",
											Path:   "example.com",
										},
									},
									&types.StringElement{
										Content: " <1>\nand <more text> on the +\n",
									},
									&types.QuotedText{
										Kind: types.SingleQuoteBold,
										Elements: []interface{}{
											&types.StringElement{
												Content: "next",
											},
										},
									},
									&types.StringElement{
										Content: " lines with a link to {github-url}[]\n\n* not a list item",
									},
								},
							},
							&types.List{
								Kind: types.CalloutListKind,
								Elements: []types.ListElement{
									&types.CalloutListElement{
										Ref: 1,
										Elements: []interface{}{
											&types.Paragraph{
												Elements: []interface{}{
													&types.StringElement{
														Content: "a callout",
													},
												},
											},
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(s)).To(MatchDocument(expected))
				})

				It("should apply the 'none' substitution", func() {
					s := strings.ReplaceAll(source, "$SUBS", "none")
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name:  "github-url",
										Value: "https://github.com",
									},
								},
							},
							&types.DelimitedBlock{
								Kind: types.Verse,
								Attributes: types.Attributes{
									types.AttrStyle:         types.Verse,
									types.AttrQuoteAuthor:   "john doe",
									types.AttrQuoteTitle:    "verse title",
									types.AttrSubstitutions: "none",
								},
								Elements: []interface{}{
									&types.StringElement{
										Content: "a link to https://example.com[] <1>\nand <more text> on the +\n*next* lines with a link to {github-url}[]\n\n* not a list item",
									},
								},
							},
							&types.List{
								Kind: types.CalloutListKind,
								Elements: []types.ListElement{
									&types.CalloutListElement{
										Ref: 1,
										Elements: []interface{}{
											&types.Paragraph{
												Elements: []interface{}{
													&types.StringElement{
														Content: "a callout",
													},
												},
											},
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(s)).To(MatchDocument(expected))
				})
			})

			Context("with variable delimiter length", func() {

				It("with 5 chars", func() {
					source := `[verse]
_____
some *verse* content
_____`
					expected := &types.Document{
						Elements: []interface{}{
							&types.DelimitedBlock{
								Kind: types.Verse,
								Attributes: types.Attributes{
									types.AttrStyle: types.Verse,
								},
								Elements: []interface{}{
									&types.StringElement{
										Content: "some ",
									},
									&types.QuotedText{
										Kind: types.SingleQuoteBold,
										Elements: []interface{}{
											&types.StringElement{
												Content: "verse",
											},
										},
									},
									&types.StringElement{
										Content: " content",
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("with 5 chars with nested with 4 chars", func() {
					// this is an edge case: the inner delimiters are treated as 3 nested italic quoted texts (single+double+single)
					source := `[verse]
_____
[verse]
____
some *verse* content
____
_____`
					expected := &types.Document{
						Elements: []interface{}{
							&types.DelimitedBlock{
								Kind: types.Verse,
								Attributes: types.Attributes{
									types.AttrStyle: types.Verse,
								},
								Elements: []interface{}{
									&types.StringElement{
										Content: "[verse]\n",
									},
									&types.QuotedText{
										Kind: types.SingleQuoteItalic,
										Elements: []interface{}{
											&types.QuotedText{
												Kind: types.DoubleQuoteItalic,
												Elements: []interface{}{
													&types.QuotedText{
														Kind: types.SingleQuoteItalic,
														Elements: []interface{}{
															&types.StringElement{
																Content: "\nsome ",
															},
															&types.QuotedText{
																Kind: types.SingleQuoteBold,
																Elements: []interface{}{
																	&types.StringElement{
																		Content: "verse",
																	},
																},
															},
															&types.StringElement{
																Content: " content\n",
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
			})
		})
	})

})
