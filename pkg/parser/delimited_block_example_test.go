package parser_test

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" // nolint:golint
	. "github.com/onsi/gomega" // nolint:golint
)

var _ = Describe("example blocks", func() {

	Context("in final documents", func() {

		Context("as delimited blocks", func() {

			It("with single rich line", func() {
				source := `====
some *example* content
====`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Example,
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											Content: "some ",
										},
										&types.QuotedText{
											Kind: types.SingleQuoteBold,
											Elements: []interface{}{
												&types.StringElement{
													Content: "example",
												},
											},
										},
										&types.StringElement{
											Content: " content",
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
.standalone
====`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Example,
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with last line starting with a dot", func() {
				source := `====
some content

.standalone
====`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Example,
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											Content: "some content",
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with multiple lines", func() {
				source := `====
.title
some listing code
with *bold content*

* and a list item
====`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Example,
							Elements: []interface{}{
								&types.Paragraph{
									Attributes: types.Attributes{
										types.AttrTitle: "title",
									},
									Elements: []interface{}{
										&types.StringElement{
											Content: "some listing code\nwith ",
										},
										&types.QuotedText{
											Kind: types.SingleQuoteBold,
											Elements: []interface{}{
												&types.StringElement{
													Content: "bold content",
												},
											},
										},
									},
								},
								&types.List{
									Kind: types.UnorderedListKind,
									Elements: []types.ListElement{
										&types.UnorderedListElement{
											BulletStyle: types.OneAsterisk,
											CheckStyle:  types.NoCheck,
											Elements: []interface{}{
												&types.Paragraph{
													Elements: []interface{}{
														&types.StringElement{
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
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with unclosed delimiter", func() {
				source := `====
End of file here`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Example,
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											Content: "End of file here",
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
some content
====`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Example,
							Attributes: types.Attributes{
								types.AttrTitle: "example block title",
							},
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											Content: "some content",
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with starting delimiter only", func() {
				source := `====`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Example,
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			Context("with custom substitutions", func() {

				// in normal blocks, the substiution should be defined and applied on the elements
				// within the blocks
				// TODO: include character replacement (eg: `(C)`)
				source := `:github-url: https://github.com
			
====
[subs="$SUBS"]
a link to https://example.com[] <1>
and <more text> on the +
*next* lines with a link to {github-url}[]

====
`
				Context("explicit substitutions", func() {

					It("should apply the default substitution", func() {
						s := strings.ReplaceAll(source, "[subs=\"$SUBS\"]\n", "")
						expected := &types.Document{
							Elements: []interface{}{
								&types.AttributeDeclaration{
									Name:  "github-url",
									Value: "https://github.com",
								},
								&types.DelimitedBlock{
									Kind: types.Example,
									Elements: []interface{}{
										&types.Paragraph{
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
								&types.AttributeDeclaration{
									Name:  "github-url",
									Value: "https://github.com",
								},
								&types.DelimitedBlock{
									Kind: types.Example,
									Elements: []interface{}{
										&types.Paragraph{
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
								&types.AttributeDeclaration{
									Name:  "github-url",
									Value: "https://github.com",
								},
								&types.DelimitedBlock{
									Kind: types.Example,
									Elements: []interface{}{
										&types.Paragraph{
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
													Content: " lines with a link to {github-url}[]",
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
								&types.AttributeDeclaration{
									Name:  "github-url",
									Value: "https://github.com",
								},
								&types.DelimitedBlock{
									Kind: types.Example,
									Elements: []interface{}{
										&types.Paragraph{
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
													Content: " <1>\nand <more text> on the +\n*next* lines with a link to {github-url}[]",
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
								&types.AttributeDeclaration{
									Name:  "github-url",
									Value: "https://github.com",
								},
								&types.DelimitedBlock{
									Kind: types.Example,
									Elements: []interface{}{
										&types.Paragraph{
											Attributes: types.Attributes{
												types.AttrSubstitutions: "attributes",
											},
											Elements: []interface{}{
												&types.StringElement{
													Content: "a link to https://example.com[] <1>\nand <more text> on the +\n*next* lines with a link to https://github.com[]",
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
								&types.AttributeDeclaration{
									Name:  "github-url",
									Value: "https://github.com",
								},
								&types.DelimitedBlock{
									Kind: types.Example,
									Elements: []interface{}{
										&types.Paragraph{
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
								&types.AttributeDeclaration{
									Name:  "github-url",
									Value: "https://github.com",
								},
								&types.DelimitedBlock{
									Kind: types.Example,
									Elements: []interface{}{
										&types.Paragraph{
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
													Content: " on the +\n*next* lines with a link to {github-url}[]",
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
								&types.AttributeDeclaration{
									Name:  "github-url",
									Value: "https://github.com",
								},
								&types.DelimitedBlock{
									Kind: types.Example,
									Elements: []interface{}{
										&types.Paragraph{
											Attributes: types.Attributes{
												types.AttrSubstitutions: "replacements",
											},
											Elements: []interface{}{
												&types.StringElement{
													Content: "a link to https://example.com[] <1>\nand <more text> on the +\n*next* lines with a link to {github-url}[]",
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
								&types.AttributeDeclaration{
									Name:  "github-url",
									Value: "https://github.com",
								},
								&types.DelimitedBlock{
									Kind: types.Example,
									Elements: []interface{}{
										&types.Paragraph{
											Attributes: types.Attributes{
												types.AttrSubstitutions: "post_replacements",
											},
											Elements: []interface{}{
												&types.StringElement{
													Content: "a link to https://example.com[] <1>\nand <more text> on the",
												},
												&types.LineBreak{},
												&types.StringElement{
													Content: "\n*next* lines with a link to {github-url}[]",
												},
											},
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
								&types.AttributeDeclaration{
									Name:  "github-url",
									Value: "https://github.com",
								},
								&types.DelimitedBlock{
									Kind: types.Example,
									Elements: []interface{}{
										&types.Paragraph{
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
													Content: " lines with a link to {github-url}[]",
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
								&types.AttributeDeclaration{
									Name:  "github-url",
									Value: "https://github.com",
								},
								&types.DelimitedBlock{
									Kind: types.Example,
									Elements: []interface{}{
										&types.Paragraph{
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
													Content: " lines with a link to {github-url}[]",
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
						s := strings.ReplaceAll(source, "$SUBS", "none") // the `none` substitution applies to the *content of the elements* with the example block
						expected := &types.Document{
							Elements: []interface{}{
								&types.AttributeDeclaration{
									Name:  "github-url",
									Value: "https://github.com",
								},
								&types.DelimitedBlock{
									Kind: types.Example,
									Elements: []interface{}{
										&types.Paragraph{
											Attributes: types.Attributes{
												types.AttrSubstitutions: "none",
											},
											Elements: []interface{}{
												&types.StringElement{
													Content: "a link to https://example.com[] <1>\nand <more text> on the +\n*next* lines with a link to {github-url}[]",
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

		Context("as paragraph blocks", func() {

			It("with single rich line", func() {
				source := `[example]
some *example* content`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Attributes: types.Attributes{
								types.AttrStyle: types.Example,
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "some ",
								},
								&types.QuotedText{
									Kind: types.SingleQuoteBold,
									Elements: []interface{}{
										&types.StringElement{
											Content: "example",
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
				result, err := ParseDocument(source)
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(MatchDocument(expected))
			})
		})
	})
})
