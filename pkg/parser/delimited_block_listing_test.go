package parser_test

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"
	log "github.com/sirupsen/logrus"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("listing blocks", func() {

	Context("in final documents", func() {

		Context("as delimited blocks", func() {

			It("with single rich line", func() {
				source := `----
some *listing* code
----`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								&types.StringElement{
									Content: "some *listing* code",
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
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Listing,
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
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								&types.StringElement{
									Content: "some listing code\nwith an empty line\n\nin the middle",
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
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								&types.StringElement{
									Content: "* some\n* listing\n* content", // suffix spaces are trimmed
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
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								&types.StringElement{
									Content: "some listing code\nwith an empty line\n\nin the middle",
								},
							},
						},
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "then a normal paragraph."},
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
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "a paragraph.",
								},
							},
						},
						&types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								&types.StringElement{
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
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								&types.StringElement{
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
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								&types.StringElement{
									Content: "import ",
								},
								&types.Callout{
									Ref: 1,
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
													Content: "an import",
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
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								&types.StringElement{
									Content: "import ",
								},
								&types.Callout{
									Ref: 1,
								},
								&types.StringElement{
									Content: "\n\nfunc foo() {} ",
								},
								&types.Callout{
									Ref: 2,
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
													Content: "an import",
												},
											},
										},
									},
								},
								&types.CalloutListElement{
									Ref: 2,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{
													Content: "a func",
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
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								&types.StringElement{
									Content: "import ",
								},
								&types.Callout{
									Ref: 1,
								},
								&types.Callout{
									Ref: 2,
								},
								&types.Callout{
									Ref: 3,
								},
								&types.StringElement{
									Content: "\n\nfunc foo() {} ",
								},
								&types.Callout{
									Ref: 4,
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
													Content: "an import",
												},
											},
										},
									},
								},
								&types.CalloutListElement{
									Ref: 2,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{
													Content: "a single import",
												},
											},
										},
									},
								},
								&types.CalloutListElement{
									Ref: 3,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{
													Content: "a single basic import",
												},
											},
										},
									},
								},
								&types.CalloutListElement{
									Ref: 4,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{
													Content: "a func",
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
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								&types.StringElement{
									Content: "import ",
								},
								&types.SpecialCharacter{
									Name: "<",
								},
								&types.StringElement{
									Content: "a",
								},
								&types.SpecialCharacter{
									Name: ">",
								},
							},
						},
						&types.Paragraph{
							Elements: []interface{}{
								&types.SpecialCharacter{
									Name: "<",
								},
								&types.StringElement{
									Content: "a",
								},
								&types.SpecialCharacter{
									Name: ">",
								},
								&types.StringElement{
									Content: " an import",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with quoted text and link in title", func() {
				source := `.a *link* to https://github.com[GitHub]
----
content
----`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Listing,
							Attributes: types.Attributes{
								types.AttrTitle: []interface{}{
									&types.StringElement{
										Content: "a ",
									},
									&types.QuotedText{
										Kind: types.SingleQuoteBold,
										Elements: []interface{}{
											&types.StringElement{
												Content: "link",
											},
										},
									},
									&types.StringElement{
										Content: " to ",
									},
									&types.InlineLink{
										Attributes: types.Attributes{
											types.AttrInlineLinkText: "GitHub",
										},
										Location: &types.Location{
											Scheme: "https://",
											Path:   "github.com",
										},
									},
								},
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "content",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
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
`
				It("should apply the default substitution", func() {
					s := strings.ReplaceAll(source, "[subs=\"$SUBS\"]\n", "") // remove the 'subs' attribute
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
								Kind: types.Listing,
								Elements: []interface{}{
									&types.StringElement{
										Content: "a link to https://example.com[] ",
									},
									&types.Callout{
										Ref: 1,
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
						},
					}
					Expect(ParseDocument(s)).To(MatchDocument(expected))
				})

				It("should apply the 'verbatim' substitution", func() {
					s := strings.ReplaceAll(source, "$SUBS", "verbatim")
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
								Kind: types.Listing,
								Attributes: types.Attributes{
									types.AttrSubstitutions: "verbatim",
								},
								Elements: []interface{}{
									&types.StringElement{
										Content: "a link to https://example.com[] ",
									},
									&types.Callout{
										Ref: 1,
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
								Kind: types.Listing,
								Attributes: types.Attributes{
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
								Kind: types.Listing,
								Attributes: types.Attributes{
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
								Kind: types.Listing,
								Attributes: types.Attributes{
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
								Kind: types.Listing,
								Attributes: types.Attributes{
									types.AttrSubstitutions: "attributes",
								},
								Elements: []interface{}{
									&types.StringElement{
										Content: "a link to https://example.com[] <1>\nand <more text> on the +\n*next* lines with a link to https://github.com[]\n\n* not a list item",
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
								Kind: types.Listing,
								Attributes: types.Attributes{
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
								Kind: types.Listing,
								Attributes: types.Attributes{
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
								Kind: types.Listing,
								Attributes: types.Attributes{
									types.AttrSubstitutions: "replacements",
								},
								Elements: []interface{}{
									&types.StringElement{
										Content: "a link to https://example.com[] <1>\nand <more text> on the +\n*next* lines with a link to {github-url}[]\n\n* not a list item",
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
								Kind: types.Listing,
								Attributes: types.Attributes{
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
						},
					}
					Expect(ParseDocument(s)).To(MatchDocument(expected))
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
								Kind: types.Listing,
								Attributes: types.Attributes{
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
								Kind: types.Listing,
								Attributes: types.Attributes{
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
								Kind: types.Listing,
								Attributes: types.Attributes{
									types.AttrSubstitutions: "none",
								},
								Elements: []interface{}{
									&types.StringElement{
										Content: "a link to https://example.com[] <1>\nand <more text> on the +\n*next* lines with a link to {github-url}[]\n\n* not a list item",
									},
								},
							},
						},
					}
					Expect(ParseDocument(s)).To(MatchDocument(expected))
				})

				It("should apply the 'quotes+' substitution", func() {
					s := strings.ReplaceAll(source, "$SUBS", "quotes+") // same as `quotes,"default"`
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
								Kind: types.Listing,
								Attributes: types.Attributes{
									types.AttrSubstitutions: "quotes+",
								},
								Elements: []interface{}{
									&types.StringElement{
										Content: "a link to https://example.com[] ",
									},
									&types.Callout{
										Ref: 1,
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
										Content: " on the +\n",
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
						},
					}
					Expect(ParseDocument(s)).To(MatchDocument(expected))
				})

				It("should apply the '-callouts' substitution", func() {
					s := strings.ReplaceAll(source, "$SUBS", "-callouts")
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
								Kind: types.Listing,
								Attributes: types.Attributes{
									types.AttrSubstitutions: "-callouts",
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
						},
					}
					Expect(ParseDocument(s)).To(MatchDocument(expected))
				})

				It("should apply the '+quotes' substitution", func() {
					s := strings.ReplaceAll(source, "$SUBS", "+quotes") // default + quotes
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
								Kind: types.Listing,
								Attributes: types.Attributes{
									types.AttrSubstitutions: "+quotes",
								},
								Elements: []interface{}{
									&types.StringElement{
										Content: "a link to https://example.com[] ",
									},
									&types.Callout{
										Ref: 1,
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
										Content: " on the +\n",
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
						},
					}
					Expect(ParseDocument(s)).To(MatchDocument(expected))
				})

				It("should apply the '-quotes' substitution", func() {
					s := strings.ReplaceAll(source, "$SUBS", "-quotes") // default - quotes
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
								Kind: types.Listing,
								Attributes: types.Attributes{
									types.AttrSubstitutions: "-quotes",
								},
								Elements: []interface{}{
									&types.StringElement{
										Content: "a link to https://example.com[] ",
									},
									&types.Callout{
										Ref: 1,
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
						},
					}
					Expect(ParseDocument(s)).To(MatchDocument(expected))
				})

				It("should fail when substitution is unknown", func() {
					logs, reset := ConfigureLogger(log.ErrorLevel)
					defer reset()
					s := strings.ReplaceAll(source, "$SUBS", "unknown")
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name:  "github-url",
										Value: string("https://github.com"),
									},
								},
							},
						},
					}
					Expect(ParseDocument(s)).To(MatchDocument(expected))
					Expect(logs).To(ContainJSONLogWithOffset(log.ErrorLevel, 33, 182, "unsupported substitution: 'unknown'"))
				})
			})

			Context("with variable delimiter length", func() {

				It("with 5 chars", func() {
					source := `-----
some *listing* content
-----`
					expected := &types.Document{
						Elements: []interface{}{
							&types.DelimitedBlock{
								Kind: types.Listing,
								Elements: []interface{}{
									&types.StringElement{
										Content: "some *listing* content",
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("with 5 chars with nested with 4 chars", func() {
					source := `-----
----
some *listing* content
----
-----`
					expected := &types.Document{
						Elements: []interface{}{
							&types.DelimitedBlock{
								Kind: types.Listing,
								Elements: []interface{}{
									&types.StringElement{
										Content: "----\nsome *listing* content\n----",
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})
			})
		})

		Context("as paragraph blocks", func() {

			It("with single rich line", func() {
				source := `[listing]
some *listing* content`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Attributes: types.Attributes{
								types.AttrStyle: types.Listing,
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "some *listing* content", // no quote substitution
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			Context("with custom substitutions", func() {
				source := `:github-url: https://github.com
		
[listing]
[subs="$SUBS"]
a link to https://example.com[] <1>
and <more text> on the +
*next* lines with a link to {github-url}[]

<1> a callout`

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
							&types.Paragraph{
								Attributes: types.Attributes{
									types.AttrStyle: types.Listing,
								},
								Elements: []interface{}{
									&types.StringElement{
										Content: "a link to https://example.com[] ",
									},
									&types.Callout{
										Ref: 1,
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
										Content: " on the +\n*next* lines with a link to {github-url}[]",
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
							&types.Paragraph{
								Attributes: types.Attributes{
									types.AttrStyle:         types.Listing,
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
										Content: " lines with a link to {github-url}[]",
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

				It("should apply the '+quotes' substitution", func() {
					s := strings.ReplaceAll(source, "$SUBS", "+quotes") // ie, default + quotes
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
							&types.Paragraph{
								Attributes: types.Attributes{
									types.AttrStyle:         types.Listing,
									types.AttrSubstitutions: "+quotes",
								},
								Elements: []interface{}{
									&types.StringElement{
										Content: "a link to https://example.com[] ",
									},
									&types.Callout{
										Ref: 1,
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
										Content: " on the +\n",
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
										Content: " lines with a link to {github-url}[]",
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

				It("should apply the 'macros,+quotes,-quotes' substitution", func() {
					s := strings.ReplaceAll(source, "$SUBS", "macros,+quotes,-quotes")
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
							&types.Paragraph{
								Attributes: types.Attributes{
									types.AttrStyle:         types.Listing,
									types.AttrSubstitutions: "macros,+quotes,-quotes", // ie, "macros" only
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
										Content: " <1>\nand <more text> on the +\n*next* lines with a link to {github-url}[]",
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
		})
	})
})
