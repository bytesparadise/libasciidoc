package parser_test

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("paragraphs", func() {

	Context("in raw documents", func() {

		Context("regular paragraphs", func() {

			It("with basic content", func() {
				source := `cookie
chocolate
pasta`
				expected := []types.DocumentFragment{
					{
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									types.RawLine("cookie"),
									types.RawLine("chocolate"),
									types.RawLine("pasta"),
								},
							},
						},
					},
				}
				Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
			})

			It("with hardbreaks attribute", func() {
				source := `[%hardbreaks]
cookie
chocolate
pasta`
				expected := []types.DocumentFragment{
					{
						Elements: []interface{}{
							&types.Paragraph{
								Attributes: types.Attributes{
									types.AttrOptions: []interface{}{"hardbreaks"},
								},
								Elements: []interface{}{
									types.RawLine("cookie"),
									types.RawLine("chocolate"),
									types.RawLine("pasta"),
								},
							},
						},
					},
				}
				Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
			})

			It("with title attribute", func() {
				source := `[title=my title]
cookie
pasta`
				expected := []types.DocumentFragment{
					{
						Elements: []interface{}{
							&types.Paragraph{
								Attributes: types.Attributes{
									types.AttrTitle: "my title",
								},
								Elements: []interface{}{
									types.RawLine("cookie"),
									types.RawLine("pasta"),
								},
							},
						},
					},
				}
				Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
			})

			It("with custom title attribute - explicit and unquoted", func() {
				source := `:title: cookies
				
[title=my {title}]
cookie
pasta`
				expected := []types.DocumentFragment{
					{
						Elements: []interface{}{
							&types.AttributeDeclaration{
								Name:  "title",
								Value: "cookies",
							},
						},
					},
					{
						Elements: []interface{}{
							&types.BlankLine{},
						},
					},
					{
						Elements: []interface{}{
							&types.Paragraph{
								Attributes: types.Attributes{
									types.AttrTitle: []interface{}{
										&types.StringElement{
											Content: "my ",
										},
										&types.AttributeSubstitution{
											Name: "title",
										},
									},
								},
								Elements: []interface{}{
									types.RawLine("cookie"),
									types.RawLine("pasta"),
								},
							},
						},
					},
				}
				Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
			})

			It("with multiple attributes and blanklines in-between", func() {
				// attributes with blanklines in-between are dropped :/
				source := `[%hardbreaks.role1.role2]

[#anchor]

cookie
pasta`
				expected := []types.DocumentFragment{
					{
						Elements: []interface{}{
							&types.BlankLine{},
						},
					},
					{
						Elements: []interface{}{
							&types.BlankLine{},
						},
					},
					{
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									types.RawLine("cookie"),
									types.RawLine("pasta"),
								},
							},
						},
					},
				}
				Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
			})

			It("with block attributes splitting 2 paragraphs", func() {
				source := `a paragraph
[.left.text-center]
another paragraph with an image image:cookie.jpg[cookie]
`
				expected := []types.DocumentFragment{
					{
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									types.RawLine("a paragraph"),
								},
							},
						},
					},
					{
						Elements: []interface{}{
							&types.Paragraph{
								Attributes: types.Attributes{
									types.AttrRoles: []interface{}{
										"left",
										"text-center",
									},
								},
								Elements: []interface{}{
									types.RawLine("another paragraph with an image image:cookie.jpg[cookie]"),
								},
							},
						},
					},
				}
				Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
			})

			It("with block attributes splitting paragraph and block image", func() {
				source := `a paragraph
[.left.text-center]
image::cookie.jpg[cookie]
`
				expected := []types.DocumentFragment{
					{
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									types.RawLine("a paragraph"),
								},
							},
						},
					},
					{
						Elements: []interface{}{
							&types.ImageBlock{
								Attributes: types.Attributes{
									types.AttrRoles: []interface{}{
										"left",
										"text-center",
									},
									types.AttrImageAlt: "cookie",
								},
								Location: &types.Location{
									Path: []interface{}{
										&types.StringElement{
											Content: "cookie.jpg",
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
			})

			Context("with custom substitutions", func() {

				// using the same input for all substitution tests
				source := `:github-url: https://github.com
:github-title: GitHub

[subs="$SUBS"]
links to {github-title}: https://github.com[{github-title}] and *<https://github.com[{github-title}]>*
and another one using attribute substitution: {github-url}[{github-title}]...
// a single-line comment.`

				It("should read multiple lines", func() {
					s := strings.ReplaceAll(source, "$SUBS", "normal")
					expected := []types.DocumentFragment{
						{
							Elements: []interface{}{
								&types.AttributeDeclaration{
									Name:  "github-url",
									Value: "https://github.com",
								},
							},
						},
						{
							Elements: []interface{}{
								&types.AttributeDeclaration{
									Name:  "github-title",
									Value: "GitHub",
								},
							},
						},
						{
							Elements: []interface{}{
								&types.BlankLine{},
							},
						},
						{
							Elements: []interface{}{
								&types.Paragraph{
									Attributes: types.Attributes{
										types.AttrSubstitutions: "normal",
									},
									Elements: []interface{}{
										types.RawLine("links to {github-title}: https://github.com[{github-title}] and *<https://github.com[{github-title}]>*"),
										types.RawLine("and another one using attribute substitution: {github-url}[{github-title}]..."),
										&types.SingleLineComment{
											Content: " a single-line comment.",
										},
									},
								},
							},
						},
					}
					Expect(ParseDocumentFragments(s)).To(MatchDocumentFragments(expected))
				})

			})
		})
	})

	Context("in final documents", func() {

		Context("regular paragraphs", func() {

			It("3 with basic content", func() {
				source := `cookie

chocolate

pasta`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "cookie",
								},
							},
						},
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "chocolate",
								},
							},
						},
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "pasta",
								},
							},
						},
					},
				}
				result, err := ParseDocument(source)
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(MatchDocument(expected))
			})

			It("with title attribute", func() {
				source := `[title=my title]
cookie
pasta`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Attributes: types.Attributes{
								types.AttrTitle: "my title",
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "cookie\npasta",
								},
							},
						},
					},
				}
				result, err := ParseDocument(source)
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(MatchDocument(expected))
			})

			It("with custom title attribute - explicit and unquoted", func() {
				source := `:title: cookies
				
[title=my {title}]
cookie
pasta`
				expected := &types.Document{
					// Attributes: types.Attributes{
					// 	"title": "cookies",
					// },
					Elements: []interface{}{
						&types.AttributeDeclaration{
							Name:  "title",
							Value: "cookies",
						},
						&types.Paragraph{
							Attributes: types.Attributes{
								types.AttrTitle: "my cookies",
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "cookie\npasta",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with custom title attribute - explicit and single quoted", func() {
				source := `:title: cookies
				
[title='my {title}']
cookie
pasta`
				expected := &types.Document{
					// Attributes: types.Attributes{
					// 	"title": "cookies",
					// },
					Elements: []interface{}{
						&types.AttributeDeclaration{
							Name:  "title",
							Value: "cookies",
						},
						&types.Paragraph{
							Attributes: types.Attributes{
								types.AttrTitle: "my cookies",
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "cookie\npasta",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with custom title attribute - explicit and double quoted", func() {
				source := `:title: cookies
				
[title="my {title}"]
cookie
pasta`
				expected := &types.Document{
					// Attributes: types.Attributes{
					// 	"title": "cookies",
					// },
					Elements: []interface{}{
						&types.AttributeDeclaration{
							Name:  "title",
							Value: "cookies",
						},
						&types.Paragraph{
							Attributes: types.Attributes{
								types.AttrTitle: "my cookies",
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "cookie\npasta",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with custom title attribute - implicit", func() {
				source := `:title: cookies
				
.my {title}
cookie
pasta`
				expected := &types.Document{
					// Attributes: types.Attributes{
					// 	"title": "cookies",
					// },
					Elements: []interface{}{
						&types.AttributeDeclaration{
							Name:  "title",
							Value: "cookies",
						},
						&types.Paragraph{
							Attributes: types.Attributes{
								types.AttrTitle: "my cookies",
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "cookie\npasta",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with multiple attributes without blanklines in-between", func() {
				source := `[%hardbreaks.role1.role2]
[#anchor]
cookie
pasta`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Attributes: types.Attributes{
								types.AttrID:      "anchor",
								types.AttrRoles:   []interface{}{"role1", "role2"},
								types.AttrOptions: []interface{}{"hardbreaks"},
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "cookie\npasta",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with multiple attributes and blanklines in-between", func() {
				// attributes are not to paragraph because of blanklines
				source := `[%hardbreaks.role1.role2]

[#anchor]

cookie
pasta`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							// Attributes: types.Attributes{
							// 	types.AttrID:      "anchor",
							// 	types.AttrRoles:   []interface{}{string("role1"), string("role2")},
							// 	types.AttrOptions: []interface{}{string("hardbreaks")},
							// },
							Elements: []interface{}{
								&types.StringElement{
									Content: "cookie\npasta",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with paragraph roles and attribute - case 1", func() {
				source := `[.role1%hardbreaks.role2]
cookie
pasta`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Attributes: types.Attributes{
								types.AttrOptions: []interface{}{"hardbreaks"},
								types.AttrRoles:   []interface{}{"role1", "role2"},
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "cookie\npasta",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with paragraph roles - case 2", func() {
				source := `[.role1%hardbreaks]
[.role2]
cookie
pasta`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Attributes: types.Attributes{
								types.AttrOptions: []interface{}{"hardbreaks"},
								types.AttrRoles:   []interface{}{"role1", "role2"},
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "cookie\npasta",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("not treat plusplus as line break", func() {
				source := `C++
cookie`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "C++\ncookie",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with block attributes splitting 2 paragraphs", func() {
				source := `a paragraph
[.left.text-center]
another paragraph with an image image:cookie.jpg[cookie]
`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "a paragraph",
								},
							},
						},
						&types.Paragraph{
							Attributes: types.Attributes{
								types.AttrRoles: []interface{}{
									"left",
									"text-center",
								},
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "another paragraph with an image ",
								},
								&types.InlineImage{
									Attributes: types.Attributes{
										types.AttrImageAlt: "cookie",
									},
									Location: &types.Location{
										Path: []interface{}{
											&types.StringElement{
												Content: "cookie.jpg",
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

			It("with block attributes splitting paragraph and block image", func() {
				source := `a paragraph
[.left.text-center]
image::cookie.jpg[cookie]
`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "a paragraph",
								},
							},
						},
						&types.ImageBlock{
							Attributes: types.Attributes{
								types.AttrRoles: []interface{}{
									"left",
									"text-center",
								},
								types.AttrImageAlt: "cookie",
							},
							Location: &types.Location{
								Path: []interface{}{
									&types.StringElement{
										Content: "cookie.jpg",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			Context("with counters", func() {

				It("default", func() {
					source := `cookie{counter:cookie} chocolate{counter2:cookie} pasta{counter:cookie} bob{counter:bob}`
					expected := &types.Document{
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "cookie1 chocolate pasta3 bob1",
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("with numeric start", func() {
					source := `cookie{counter:cookie:2} chocolate{counter2:cookie} pasta{counter:cookie} bob{counter:bob:10}`
					expected := &types.Document{
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "cookie2 chocolate pasta4 bob10",
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("with alphanumeric start", func() {
					source := `cookie{counter:cookie:b} chocolate{counter2:cookie} pasta{counter:cookie} bob{counter:bob:z}`
					expected := &types.Document{
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "cookieb chocolate pastad bobz",
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})
			})

			It("with custom id prefix and title", func() {
				source := `:idprefix: bar_
			
.a title
a paragraph`
				expected := &types.Document{
					// Attributes: types.Attributes{
					// 	types.AttrIDPrefix: "bar_",
					// },
					Elements: []interface{}{
						&types.AttributeDeclaration{
							Name:  "idprefix",
							Value: "bar_",
						},
						&types.Paragraph{
							Attributes: types.Attributes{
								types.AttrTitle: "a title", // there is no default ID. Only custom IDs
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "a paragraph",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("empty paragraph", func() {
				source := `{blank}`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.PredefinedAttribute{
									Name: "blank",
								},
							},
						},
					},
				}
				result, err := ParseDocument(source)
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(MatchDocument(expected))
			})

			It("paragraph with predefined attribute", func() {
				source := "hello {plus} world"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "hello ",
								},
								&types.PredefinedAttribute{Name: "plus"},
								&types.StringElement{
									Content: " world",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("2 paragraphs with list item continuation", func() {
				source := `hello
+
world`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "hello",
								},
							},
						},
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "+\nworld",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("2 paragraphs with list item continuation after blankline", func() {
				source := `hello
+
world`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "hello",
								},
							},
						},
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "+\nworld",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			Context("with custom substitutions", func() {

				// using the same input for all substitution tests
				source := `:github-url: https://github.com
:github-title: GitHub

[subs="$SUBS"]
links to {github-title}: https://github.com[{github-title}] and *<https://github.com[_{github-title}_]>*
and another one using attribute substitution: {github-url}[{github-title}]...
// a single-line comment.`

				It("should apply the 'default' substitution", func() {
					// quoted text is parsed but inline link macro is not
					s := strings.ReplaceAll(source, "[subs=\"$SUBS\"]\n", "")
					expected := &types.Document{
						// Attributes: types.Attributes{
						// 	"github-url":   "https://github.com",
						// 	"github-title": "GitHub",
						// },
						Elements: []interface{}{
							&types.AttributeDeclaration{
								Name:  "github-url",
								Value: "https://github.com",
							},
							&types.AttributeDeclaration{
								Name:  "github-title",
								Value: "GitHub",
							},
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "links to GitHub: ",
									},
									&types.InlineLink{
										Attributes: types.Attributes{
											types.AttrInlineLinkText: "GitHub",
										},
										Location: &types.Location{
											Scheme: "https://",
											Path: []interface{}{
												&types.StringElement{
													Content: "github.com",
												},
											},
										},
									},
									&types.StringElement{
										Content: " and ",
									},
									&types.QuotedText{
										Kind: types.SingleQuoteBold,
										Elements: []interface{}{
											&types.SpecialCharacter{
												Name: "<",
											},
											&types.InlineLink{
												Attributes: types.Attributes{
													types.AttrInlineLinkText: []interface{}{
														&types.QuotedText{
															Kind: types.SingleQuoteItalic,
															Elements: []interface{}{
																&types.StringElement{
																	Content: "GitHub",
																},
															},
														},
													},
												},
												Location: &types.Location{
													Scheme: "https://",
													Path: []interface{}{
														&types.StringElement{
															Content: "github.com",
														},
													},
												},
											},
											&types.SpecialCharacter{
												Name: ">",
											},
										},
									},
									&types.StringElement{
										Content: "\nand another one using attribute substitution: ",
									},
									&types.InlineLink{
										Attributes: types.Attributes{
											types.AttrInlineLinkText: "GitHub",
										},
										Location: &types.Location{
											Scheme: "https://",
											Path: []interface{}{
												&types.StringElement{
													Content: "github.com",
												},
											},
										},
									},
									&types.StringElement{
										Content: "\u2026\u200b", // symbol for ellipsis, applied by the 'replacements' substitution
									},
								},
							},
						},
					}
					Expect(ParseDocument(s)).To(MatchDocument(expected))
				})

				It("should apply the 'normal' substitution", func() {
					// quoted text is parsed but inline link macro is not
					s := strings.ReplaceAll(source, "$SUBS", "normal")
					expected := &types.Document{
						// Attributes: types.Attributes{
						// 	"github-url":   "https://github.com",
						// 	"github-title": "GitHub",
						// },
						Elements: []interface{}{
							&types.AttributeDeclaration{
								Name:  "github-url",
								Value: "https://github.com",
							},
							&types.AttributeDeclaration{
								Name:  "github-title",
								Value: "GitHub",
							},
							&types.Paragraph{
								Attributes: types.Attributes{
									types.AttrSubstitutions: "normal",
								},
								Elements: []interface{}{
									&types.StringElement{
										Content: "links to GitHub: ",
									},
									&types.InlineLink{
										Attributes: types.Attributes{
											types.AttrInlineLinkText: "GitHub",
										},
										Location: &types.Location{
											Scheme: "https://",
											Path: []interface{}{
												&types.StringElement{
													Content: "github.com",
												},
											},
										},
									},
									&types.StringElement{
										Content: " and ",
									},
									&types.QuotedText{
										Kind: types.SingleQuoteBold,
										Elements: []interface{}{
											&types.SpecialCharacter{
												Name: "<",
											},
											&types.InlineLink{
												Attributes: types.Attributes{
													types.AttrInlineLinkText: []interface{}{
														&types.QuotedText{
															Kind: types.SingleQuoteItalic,
															Elements: []interface{}{
																&types.StringElement{ //
																	Content: "GitHub",
																},
															},
														},
													},
												},
												Location: &types.Location{
													Scheme: "https://",
													Path: []interface{}{
														&types.StringElement{
															Content: "github.com",
														},
													},
												},
											},
											&types.SpecialCharacter{
												Name: ">",
											},
										},
									},
									&types.StringElement{
										Content: "\nand another one using attribute substitution: ",
									},
									&types.InlineLink{
										Attributes: types.Attributes{
											types.AttrInlineLinkText: "GitHub",
										},
										Location: &types.Location{
											Scheme: "https://",
											Path: []interface{}{
												&types.StringElement{
													Content: "github.com",
												},
											},
										},
									},
									&types.StringElement{
										Content: "\u2026\u200b", // symbol for ellipsis, applied by the 'replacements' substitution
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
						// Attributes: types.Attributes{
						// 	"github-url":   "https://github.com",
						// 	"github-title": "GitHub",
						// },
						Elements: []interface{}{
							&types.AttributeDeclaration{
								Name:  "github-url",
								Value: "https://github.com",
							},
							&types.AttributeDeclaration{
								Name:  "github-title",
								Value: "GitHub",
							},
							&types.Paragraph{
								Attributes: types.Attributes{
									types.AttrSubstitutions: "none",
								},
								Elements: []interface{}{
									&types.StringElement{
										Content: "links to {github-title}: https://github.com[{github-title}] and *<https://github.com[_{github-title}_]>*" +
											"\nand another one using attribute substitution: {github-url}[{github-title}]...",
									},
								},
							},
						},
					}
					Expect(ParseDocument(s)).To(MatchDocument(expected))
				})

				It("should apply the 'quotes' substitution", func() {
					// quoted text is parsed but inline link macros are not
					s := strings.ReplaceAll(source, "$SUBS", "quotes")
					expected := &types.Document{
						// Attributes: types.Attributes{
						// 	"github-url":   "https://github.com",
						// 	"github-title": "GitHub",
						// },
						Elements: []interface{}{
							&types.AttributeDeclaration{
								Name:  "github-url",
								Value: "https://github.com",
							},
							&types.AttributeDeclaration{
								Name:  "github-title",
								Value: "GitHub",
							},
							&types.Paragraph{
								Attributes: types.Attributes{
									types.AttrSubstitutions: "quotes",
								},
								Elements: []interface{}{
									&types.StringElement{
										Content: "links to {github-title}: https://github.com[{github-title}] and ",
									},
									&types.QuotedText{
										Kind: types.SingleQuoteBold,
										Elements: []interface{}{
											&types.StringElement{
												Content: "<https://github.com[",
											},
											&types.QuotedText{
												Kind: types.SingleQuoteItalic,
												Elements: []interface{}{
													&types.StringElement{
														Content: "{github-title}",
													},
												},
											},
											&types.StringElement{
												Content: "]>",
											},
										},
									},
									&types.StringElement{
										Content: "\nand another one using attribute substitution: {github-url}[{github-title}]...",
									},
								},
							},
						},
					}
					Expect(ParseDocument(s)).To(MatchDocument(expected))
				})

				It("should apply the 'macros' substitution", func() {
					// quoted text is not parsed but inline link macro is
					s := strings.ReplaceAll(source, "$SUBS", "macros")
					expected := &types.Document{
						// Attributes: types.Attributes{
						// 	"github-url":   "https://github.com",
						// 	"github-title": "GitHub",
						// },
						Elements: []interface{}{
							&types.AttributeDeclaration{
								Name:  "github-url",
								Value: "https://github.com",
							},
							&types.AttributeDeclaration{
								Name:  "github-title",
								Value: "GitHub",
							},
							&types.Paragraph{
								Attributes: types.Attributes{
									types.AttrSubstitutions: "macros",
								},
								Elements: []interface{}{
									&types.StringElement{
										Content: "links to {github-title}: ",
									},
									&types.InlineLink{
										Attributes: types.Attributes{
											types.AttrInlineLinkText: "{github-title}",
										},
										Location: &types.Location{
											Scheme: "https://",
											Path: []interface{}{
												&types.StringElement{
													Content: "github.com",
												},
											},
										},
									},
									&types.StringElement{
										Content: " and *<",
									},
									&types.InlineLink{
										Attributes: types.Attributes{
											types.AttrInlineLinkText: "_{github-title}_",
										},
										Location: &types.Location{
											Scheme: "https://",
											Path: []interface{}{
												&types.StringElement{
													Content: "github.com",
												},
											},
										},
									},
									&types.StringElement{
										Content: ">*\nand another one using attribute substitution: {github-url}[{github-title}]...",
									},
								},
							},
						},
					}
					Expect(ParseDocument(s)).To(MatchDocument(expected))
				})

				It("should apply the 'attributes' substitution", func() {
					// quoted text is not parsed but inline link macro is
					s := strings.ReplaceAll(source, "$SUBS", "attributes")
					expected := &types.Document{
						// Attributes: types.Attributes{
						// 	"github-url":   "https://github.com",
						// 	"github-title": "GitHub",
						// },
						Elements: []interface{}{
							&types.AttributeDeclaration{
								Name:  "github-url",
								Value: "https://github.com",
							},
							&types.AttributeDeclaration{
								Name:  "github-title",
								Value: "GitHub",
							},
							&types.Paragraph{
								Attributes: types.Attributes{
									types.AttrSubstitutions: "attributes",
								},
								Elements: []interface{}{
									&types.StringElement{
										Content: "links to GitHub: https://github.com[GitHub] and *<https://github.com[_GitHub_]>*" +
											"\nand another one using attribute substitution: https://github.com[GitHub]...",
									},
								},
							},
						},
					}
					Expect(ParseDocument(s)).To(MatchDocument(expected))
				})

				It("should apply the 'specialchars' substitution", func() {
					// quoted text is not parsed but inline link macro is
					s := strings.ReplaceAll(source, "$SUBS", "specialchars")
					expected := &types.Document{
						// Attributes: types.Attributes{
						// 	"github-url":   "https://github.com",
						// 	"github-title": "GitHub",
						// },
						Elements: []interface{}{
							&types.AttributeDeclaration{
								Name:  "github-url",
								Value: "https://github.com",
							},
							&types.AttributeDeclaration{
								Name:  "github-title",
								Value: "GitHub",
							},
							&types.Paragraph{
								Attributes: types.Attributes{
									types.AttrSubstitutions: "specialchars",
								},
								Elements: []interface{}{
									&types.StringElement{
										Content: "links to {github-title}: https://github.com[{github-title}] and *",
									},
									&types.SpecialCharacter{
										Name: "<",
									},
									&types.StringElement{
										Content: "https://github.com[_{github-title}_]",
									},
									&types.SpecialCharacter{
										Name: ">",
									},
									&types.StringElement{
										Content: "*" +
											"\nand another one using attribute substitution: {github-url}[{github-title}]...",
									},
								},
							},
						},
					}
					Expect(ParseDocument(s)).To(MatchDocument(expected))
				})

				It("should apply the 'replacements' substitution", func() {
					// quoted text is not parsed but inline link macro is
					s := strings.ReplaceAll(source, "$SUBS", "replacements")
					expected := &types.Document{
						// Attributes: types.Attributes{
						// 	"github-url":   "https://github.com",
						// 	"github-title": "GitHub",
						// },
						Elements: []interface{}{
							&types.AttributeDeclaration{
								Name:  "github-url",
								Value: "https://github.com",
							},
							&types.AttributeDeclaration{
								Name:  "github-title",
								Value: "GitHub",
							},
							&types.Paragraph{
								Attributes: types.Attributes{
									types.AttrSubstitutions: "replacements",
								},
								Elements: []interface{}{
									&types.StringElement{
										Content: "links to {github-title}: https://github.com[{github-title}] and *<https://github.com[_{github-title}_]>*" +
											"\nand another one using attribute substitution: {github-url}[{github-title}]\u2026\u200b",
									},
								},
							},
						},
					}
					Expect(ParseDocument(s)).To(MatchDocument(expected))
				})

				It("should apply the 'quotes,macros' substitutions", func() {
					// quoted texts and macros are both parsed at the root of the paragraph content
					// but and macros within quotes are, too
					// Note: Asciidoctor 2.0.12 does not parse the macros within the quotes, so the
					// 2nd link of the first line is not detected
					s := strings.ReplaceAll(source, "$SUBS", "quotes,macros")
					expected := &types.Document{
						// Attributes: types.Attributes{
						// 	"github-url":   "https://github.com",
						// 	"github-title": "GitHub",
						// },
						Elements: []interface{}{
							&types.AttributeDeclaration{
								Name:  "github-url",
								Value: "https://github.com",
							},
							&types.AttributeDeclaration{
								Name:  "github-title",
								Value: "GitHub",
							},
							&types.Paragraph{
								Attributes: types.Attributes{
									types.AttrSubstitutions: "quotes,macros",
								},
								Elements: []interface{}{
									&types.StringElement{
										Content: "links to {github-title}: ",
									},
									&types.InlineLink{
										Attributes: types.Attributes{
											types.AttrInlineLinkText: "{github-title}",
										},
										Location: &types.Location{
											Scheme: "https://",
											Path: []interface{}{
												&types.StringElement{
													Content: "github.com",
												},
											},
										},
									},
									&types.StringElement{
										Content: " and ",
									},
									&types.QuotedText{
										Kind: types.SingleQuoteBold,
										Elements: []interface{}{
											&types.StringElement{
												Content: "<",
											},
											&types.InlineLink{
												Attributes: types.Attributes{
													types.AttrInlineLinkText: []interface{}{
														&types.QuotedText{
															Kind: types.SingleQuoteItalic,
															Elements: []interface{}{
																&types.StringElement{
																	Content: "{github-title}",
																},
															},
														},
													},
												},
												Location: &types.Location{
													Scheme: "https://",
													Path: []interface{}{
														&types.StringElement{
															Content: "github.com",
														},
													},
												},
											},
											&types.StringElement{
												Content: ">",
											},
										},
									},
									&types.StringElement{
										Content: "\nand another one using attribute substitution: {github-url}[{github-title}]...",
									},
								},
							},
						},
					}
					Expect(ParseDocument(s)).To(MatchDocument(expected))
				})

				It("should apply the 'macros,quotes' substitutions", func() {
					// quoted text and inline link macro are both parsed
					// (same as above, but with subs in reversed order)
					s := strings.ReplaceAll(source, "$SUBS", "macros,quotes")
					expected := &types.Document{
						// Attributes: types.Attributes{
						// 	"github-url":   "https://github.com",
						// 	"github-title": "GitHub",
						// },
						Elements: []interface{}{
							&types.AttributeDeclaration{
								Name:  "github-url",
								Value: "https://github.com",
							},
							&types.AttributeDeclaration{
								Name:  "github-title",
								Value: "GitHub",
							},
							&types.Paragraph{
								Attributes: types.Attributes{
									types.AttrSubstitutions: "macros,quotes",
								},
								Elements: []interface{}{
									&types.StringElement{
										Content: "links to {github-title}: ",
									},
									&types.InlineLink{
										Attributes: types.Attributes{
											types.AttrInlineLinkText: "{github-title}",
										},
										Location: &types.Location{
											Scheme: "https://",
											Path: []interface{}{
												&types.StringElement{
													Content: "github.com",
												},
											},
										},
									},
									&types.StringElement{
										Content: " and ",
									},
									&types.QuotedText{
										Kind: types.SingleQuoteBold,
										Elements: []interface{}{
											&types.StringElement{
												Content: "<",
											},
											&types.InlineLink{
												Attributes: types.Attributes{
													types.AttrInlineLinkText: []interface{}{
														&types.QuotedText{
															Kind: types.SingleQuoteItalic,
															Elements: []interface{}{
																&types.StringElement{
																	Content: "{github-title}",
																},
															},
														},
													},
												},
												Location: &types.Location{
													Scheme: "https://",
													Path: []interface{}{
														&types.StringElement{
															Content: "github.com",
														},
													},
												},
											},
											&types.StringElement{
												Content: ">",
											},
										},
									},
									&types.StringElement{
										Content: "\nand another one using attribute substitution: {github-url}[{github-title}]...",
									},
								},
							},
						},
					}
					Expect(ParseDocument(s)).To(MatchDocument(expected))
				})

				It("should apply the 'attributes,macros' substitution", func() {
					// inline links are fully parsed
					s := strings.ReplaceAll(source, "$SUBS", "attributes,macros")
					expected := &types.Document{
						// Attributes: types.Attributes{
						// 	"github-url":   "https://github.com",
						// 	"github-title": "GitHub",
						// },
						Elements: []interface{}{
							&types.AttributeDeclaration{
								Name:  "github-url",
								Value: "https://github.com",
							},
							&types.AttributeDeclaration{
								Name:  "github-title",
								Value: "GitHub",
							},
							&types.Paragraph{
								Attributes: types.Attributes{
									types.AttrSubstitutions: "attributes,macros",
								},
								Elements: []interface{}{
									&types.StringElement{
										Content: "links to GitHub: ",
									},
									&types.InlineLink{
										Attributes: types.Attributes{
											types.AttrInlineLinkText: "GitHub",
										},
										Location: &types.Location{
											Scheme: "https://",
											Path: []interface{}{
												&types.StringElement{
													Content: "github.com",
												},
											},
										},
									},
									&types.StringElement{
										Content: " and *<",
									},
									&types.InlineLink{
										Attributes: types.Attributes{
											types.AttrInlineLinkText: "_GitHub_",
										},
										Location: &types.Location{
											Scheme: "https://",
											Path: []interface{}{
												&types.StringElement{
													Content: "github.com",
												},
											},
										},
									},
									&types.StringElement{
										Content: ">*\nand another one using attribute substitution: ",
									},
									&types.InlineLink{
										Attributes: types.Attributes{
											types.AttrInlineLinkText: "GitHub",
										},
										Location: &types.Location{
											Scheme: "https://",
											Path: []interface{}{
												&types.StringElement{
													Content: "github.com",
												},
											},
										},
									},
									&types.StringElement{
										Content: "...", // left as-is
									},
								},
							},
						},
					}
					Expect(ParseDocument(s)).To(MatchDocument(expected))
				})

				It("should apply the 'macros,attributes' substitution", func() {
					// inline links with URL coming from attribute susbtitutions are left as-is
					// however, inline link 'text' attribute coming from attribute susbtitutions are replaced
					s := strings.ReplaceAll(source, "$SUBS", "macros,attributes")
					expected := &types.Document{
						// Attributes: types.Attributes{
						// 	"github-url":   "https://github.com",
						// 	"github-title": "GitHub",
						// },
						Elements: []interface{}{
							&types.AttributeDeclaration{
								Name:  "github-url",
								Value: "https://github.com",
							},
							&types.AttributeDeclaration{
								Name:  "github-title",
								Value: "GitHub",
							},
							&types.Paragraph{
								Attributes: types.Attributes{
									types.AttrSubstitutions: "macros,attributes",
								},
								Elements: []interface{}{
									&types.StringElement{
										Content: "links to GitHub: ",
									},
									&types.InlineLink{
										Attributes: types.Attributes{
											types.AttrInlineLinkText: "GitHub",
										},
										Location: &types.Location{
											Scheme: "https://",
											Path: []interface{}{
												&types.StringElement{
													Content: "github.com",
												},
											},
										},
									},
									&types.StringElement{
										Content: " and *<",
									},
									&types.InlineLink{
										Attributes: types.Attributes{
											types.AttrInlineLinkText: "_GitHub_",
										},
										Location: &types.Location{
											Scheme: "https://",
											Path: []interface{}{
												&types.StringElement{
													Content: "github.com",
												},
											},
										},
									},
									&types.StringElement{
										Content: ">*\nand another one using attribute substitution: https://github.com[GitHub]...",
									},
								},
							},
						},
					}
					Expect(ParseDocument(s)).To(MatchDocument(expected))
				})
			})
		})

		Context("admonition paragraphs", func() {

			It("note admonition paragraph", func() {
				source := `NOTE: this is a note.`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Attributes: types.Attributes{
								types.AttrStyle: types.Note,
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "this is a note.",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("warning admonition paragraph", func() {
				source := `WARNING: this is a multiline
warning!`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Attributes: types.Attributes{
								types.AttrStyle: types.Warning,
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "this is a multiline\nwarning!",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("admonition note paragraph with id and title", func() {
				source := `[[cookie]]
.chocolate
NOTE: this is a note.`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Attributes: types.Attributes{
								types.AttrStyle: types.Note,
								types.AttrID:    "cookie",
								types.AttrTitle: "chocolate",
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "this is a note.",
								},
							},
						},
					},
					ElementReferences: types.ElementReferences{
						"cookie": "chocolate",
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("caution admonition paragraph with single line", func() {
				source := `[CAUTION]
this is a caution!`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Attributes: types.Attributes{
								types.AttrStyle: types.Caution,
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "this is a caution!",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("multiline caution admonition paragraph with title and id", func() {
				source := `[[cookie]]
[CAUTION] 
.chocolate
this is a 
*caution*!`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Attributes: types.Attributes{
								types.AttrStyle: types.Caution,
								types.AttrID:    "cookie",
								types.AttrTitle: "chocolate",
							},
							Elements: []interface{}{
								// suffix spaces are trimmed on each line
								&types.StringElement{
									Content: "this is a\n",
								},
								&types.QuotedText{
									Kind: types.SingleQuoteBold,
									Elements: []interface{}{
										&types.StringElement{
											Content: "caution",
										},
									},
								},
								&types.StringElement{
									Content: "!",
								},
							},
						},
					},
					ElementReferences: types.ElementReferences{
						"cookie": "chocolate",
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("multiple admonition paragraphs", func() {
				source := `[NOTE]
No space after the [NOTE]!

[CAUTION]
And no space after [CAUTION] either.`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Attributes: types.Attributes{
								types.AttrStyle: types.Note,
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "No space after the [NOTE]!",
								},
							},
						},
						&types.Paragraph{
							Attributes: types.Attributes{
								types.AttrStyle: types.Caution,
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "And no space after [CAUTION] either.",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("not an admonition paragraph", func() {
				source := `cookie
NOTE: a note`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "cookie\nNOTE: a note",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})

		Context("quote paragraphs", func() {

			It("inline image within a quote", func() {
				source := `[quote, john doe, quote title]
a cookie image:cookie.png[]`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Attributes: types.Attributes{
								types.AttrStyle:       types.Quote,
								types.AttrQuoteAuthor: "john doe",
								types.AttrQuoteTitle:  "quote title",
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "a cookie ",
								},
								&types.InlineImage{
									Location: &types.Location{
										Path: []interface{}{
											&types.StringElement{
												Content: "cookie.png",
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

		Context("verse paragraphs", func() {

			It("with author and title", func() {
				source := `[verse, john doe, verse title]
I am a verse paragraph.`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Attributes: types.Attributes{
								types.AttrStyle:       types.Verse,
								types.AttrQuoteAuthor: "john doe",
								types.AttrQuoteTitle:  "verse title",
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "I am a verse paragraph.",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with author, title and other attributes", func() {
				source := `[[universal]]
[verse, john doe, verse title]
.universe
I am a verse paragraph.`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Attributes: types.Attributes{
								types.AttrStyle:       types.Verse,
								types.AttrQuoteAuthor: "john doe",
								types.AttrQuoteTitle:  "verse title",
								types.AttrID:          "universal",
								types.AttrTitle:       "universe",
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "I am a verse paragraph.",
								},
							},
						},
					},
					ElementReferences: types.ElementReferences{
						"universal": "universe",
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with empty title", func() {
				source := `[verse, john doe, ]
I am a verse paragraph.`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Attributes: types.Attributes{
								types.AttrStyle:       types.Verse,
								types.AttrQuoteAuthor: "john doe",
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "I am a verse paragraph.",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("without title", func() {
				source := `[verse, john doe ]
I am a verse paragraph.`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Attributes: types.Attributes{
								types.AttrStyle:       types.Verse,
								types.AttrQuoteAuthor: "john doe",
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "I am a verse paragraph.",
								},
							},
						},
					},
				}
				doc, err := ParseDocument(source)
				Expect(err).NotTo(HaveOccurred())
				Expect(doc).To(MatchDocument(expected))
			})

			It("with empty author", func() {
				source := `[verse,  ]
I am a verse paragraph.`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Attributes: types.Attributes{
								types.AttrStyle: types.Verse,
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "I am a verse paragraph.",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("without author", func() {
				source := `[verse]
I am a verse paragraph.`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Attributes: types.Attributes{
								types.AttrStyle: types.Verse,
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "I am a verse paragraph.",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			// 			It("image block as a verse", func() {
			// 				// assume that the author meant to use an image, so the `verse` attribute will be ignored during rendering
			// 				source := `[verse, john doe, verse title]
			// image::cookie.png[]`
			// 				expected := &types.Document{
			// 					Elements: []interface{}{
			// 						&types.Paragraph{
			// 							Attributes: types.Attributes{
			// 								types.AttrStyle:       types.Verse,
			// 								types.AttrQuoteAuthor: "john doe",
			// 								types.AttrQuoteTitle:  "verse title",
			// 							},
			// 							Elements: []interface{}{
			// 								&types.StringElement{Content: "image::cookie.png[]",},
			// 							},
			// 						},
			// 					},
			// 				}
			// 				Expect(ParseDocument(source)).To(MatchDocument(expected))
			// 			})
		})

	})
})
