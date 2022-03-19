package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("links", func() {

	Context("in final documents", func() {

		Context("bare URLs", func() {

			It("should parse standalone URL with scheme ", func() {
				source := `<https://example.com>`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.InlineLink{
									Location: &types.Location{
										Scheme: "https://",
										Path:   "example.com",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("should parse URL with scheme in sentence", func() {
				source := `a link to <https://example.com>.`
				expected := &types.Document{
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
									Content: ".",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("should parse substituted URL with scheme", func() {
				source := `:example: https://example.com

a link to <{example}>.`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Elements: []interface{}{
								&types.AttributeDeclaration{
									Name:  "example",
									Value: "https://example.com",
								},
							},
						},
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
									Content: ".",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			Context("malformed", func() {

				It("should not parse URL without scheme", func() {
					source := `a link to <example.com>.`
					expected := &types.Document{
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "a link to ",
									},
									&types.SpecialCharacter{
										Name: "<",
									},
									&types.StringElement{
										Content: "example.com",
									},
									&types.SpecialCharacter{
										Name: ">",
									},
									&types.StringElement{
										Content: ".",
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("should parse with special character in URL", func() {
					source := `a link to https://example.com>[].`
					expected := &types.Document{
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "a link to ",
									},
									&types.InlineLink{
										Location: &types.Location{
											Scheme: "https://",
											Path: []interface{}{
												&types.StringElement{
													Content: "example.com",
												},
												&types.SpecialCharacter{
													Name: ">",
												},
											},
										},
									},
									&types.StringElement{
										Content: ".",
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("should parse with opening angle bracket", func() {
					source := `a link to <https://example.com[].`
					expected := &types.Document{
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "a link to ",
									},
									&types.SpecialCharacter{
										Name: "<",
									},
									&types.InlineLink{
										Location: &types.Location{
											Scheme: "https://",
											Path:   "example.com",
										},
									},
									&types.StringElement{
										Content: ".",
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})
			})
		})

		Context("external links", func() {

			It("without text", func() {
				source := "a link to https://example.com"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "a link to "},
								&types.InlineLink{
									Location: &types.Location{
										Scheme: "https://",
										Path:   "example.com",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with trailing dot punctutation", func() {
				source := "a link to https://example.com."
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "a link to "},
								&types.InlineLink{
									Location: &types.Location{
										Scheme: "https://",
										Path:   "example.com",
									},
								},
								&types.StringElement{Content: "."},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with trailing question mark punctutation", func() {
				source := "a link to https://example.com?"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "a link to "},
								&types.InlineLink{
									Location: &types.Location{
										Scheme: "https://",
										Path:   "example.com",
									},
								},
								&types.StringElement{Content: "?"},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with empty text", func() {
				source := "a link to https://example.com[]"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "a link to "},
								&types.InlineLink{
									Location: &types.Location{
										Scheme: "https://",
										Path:   "example.com",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with text only", func() {
				source := "a link to mailto:foo@bar[the foo@bar email]."
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "a link to "},
								&types.InlineLink{
									Location: &types.Location{
										Scheme: "mailto:",
										Path:   "foo@bar",
									},
									Attributes: types.Attributes{
										types.AttrInlineLinkText: "the foo@bar email",
									},
								},
								&types.StringElement{
									Content: ".",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with text and extra attributes", func() {
				source := "a link to mailto:foo@bar[the foo@bar email, foo=bar]"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "a link to "},
								&types.InlineLink{
									Location: &types.Location{
										Scheme: "mailto:",
										Path:   "foo@bar",
									},
									Attributes: types.Attributes{
										types.AttrInlineLinkText: "the foo@bar email",
										"foo":                    "bar",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("inside a multiline paragraph -  without attributes", func() {
				source := `a http://website.com
and more text on the
next lines`

				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "a ",
								},
								&types.InlineLink{
									Location: &types.Location{
										Scheme: "http://",
										Path:   "website.com",
									},
								},
								&types.StringElement{
									Content: "\nand more text on the\nnext lines",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("inside a multiline paragraph -  with attributes", func() {
				source := `a http://website.com[]
and more text on the
next lines`

				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "a ",
								},
								&types.InlineLink{
									Location: &types.Location{
										Scheme: "http://",
										Path:   "website.com",
									},
								},
								&types.StringElement{
									Content: "\nand more text on the\nnext lines",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with more text afterwards", func() {
				source := `a link to https://example.com and more text`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "a link to "},
								&types.InlineLink{
									Location: &types.Location{
										Scheme: "https://",
										Path:   "example.com",
									},
								},
								&types.StringElement{Content: " and more text"},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			Context("text attribute with comma", func() {

				It("only with text having comma", func() {
					source := `a link to http://website.com[A, B, and C]`
					expected := &types.Document{
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{Content: "a link to "},
									&types.InlineLink{
										Location: &types.Location{
											Scheme: "http://",
											Path:   "website.com",
										},
										Attributes: types.Attributes{
											types.AttrInlineLinkText: "A",
											types.AttrPositional2:    "B",
											types.AttrPositional3:    "and C",
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("only with doublequoted text having comma", func() {
					source := `a link to http://website.com["A, B, and C"]`
					expected := &types.Document{
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{Content: "a link to "},
									&types.InlineLink{
										Location: &types.Location{
											Scheme: "http://",
											Path:   "website.com",
										},
										Attributes: types.Attributes{
											types.AttrInlineLinkText: "A, B, and C",
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("with doublequoted text having comma and other attrs", func() {
					source := `a link to http://website.com["A, B, and C", role=foo]`
					expected := &types.Document{
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{Content: "a link to "},
									&types.InlineLink{
										Location: &types.Location{
											Scheme: "http://",
											Path:   "website.com",
										},
										Attributes: types.Attributes{
											types.AttrInlineLinkText: "A, B, and C",
											types.AttrRoles:          types.Roles{"foo"},
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("with text having comma and other attributes", func() {
					source := `a link to http://website.com[A, B, and C, role=foo]`
					expected := &types.Document{
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{Content: "a link to "},
									&types.InlineLink{
										Location: &types.Location{
											Scheme: "http://",
											Path:   "website.com",
										},
										Attributes: types.Attributes{
											types.AttrInlineLinkText: "A",
											types.AttrPositional2:    "B",
											types.AttrPositional3:    "and C",
											types.AttrRoles:          types.Roles{"foo"},
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})
			})

			It("with special characters", func() {
				source := "a link to https://foo*_.com"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "a link to "},
								&types.InlineLink{
									Location: &types.Location{
										Scheme: "https://",
										Path:   "foo*_.com",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with quoted text without attributes", func() {
				source := "a link to https://example.com[_a_ *b* `c`]"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "a link to ",
								},
								&types.InlineLink{
									Attributes: types.Attributes{
										types.AttrInlineLinkText: []interface{}{
											&types.QuotedText{
												Kind: types.SingleQuoteItalic,
												Elements: []interface{}{
													&types.StringElement{
														Content: "a",
													},
												},
											},
											&types.StringElement{
												Content: " ",
											},
											&types.QuotedText{
												Kind: types.SingleQuoteBold,
												Elements: []interface{}{
													&types.StringElement{
														Content: "b",
													},
												},
											},
											&types.StringElement{
												Content: " ",
											},
											&types.QuotedText{
												Kind: types.SingleQuoteMonospace,
												Elements: []interface{}{
													&types.StringElement{
														Content: "c",
													},
												},
											},
										},
									},
									Location: &types.Location{
										Scheme: "https://",
										Path:   "example.com",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with quoted text with attributes", func() {
				source := "a link to https://example.com[[.myrole1]_a_]"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "a link to "},
								&types.InlineLink{
									Attributes: types.Attributes{
										types.AttrInlineLinkText: []interface{}{
											&types.QuotedText{
												Kind: types.SingleQuoteItalic,
												Attributes: types.Attributes{
													types.AttrRoles: types.Roles{"myrole1"},
												},
												Elements: []interface{}{
													&types.StringElement{
														Content: "a",
													},
												},
											},
										},
									},
									Location: &types.Location{
										Scheme: "https://",
										Path:   "example.com",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with quoted texts with attributes", func() {
				source := "a link to https://example.com[[.myrole1]_a_ [.myrole2]*b* [.myrole3]`c`]"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "a link to "},
								&types.InlineLink{
									Attributes: types.Attributes{
										types.AttrInlineLinkText: []interface{}{
											&types.QuotedText{
												Kind: types.SingleQuoteItalic,
												Attributes: types.Attributes{
													types.AttrRoles: types.Roles{"myrole1"},
												},
												Elements: []interface{}{
													&types.StringElement{
														Content: "a",
													},
												},
											},
											&types.StringElement{
												Content: " ",
											},
											&types.QuotedText{
												Kind: types.SingleQuoteBold,
												Attributes: types.Attributes{
													types.AttrRoles: types.Roles{"myrole2"},
												},
												Elements: []interface{}{
													&types.StringElement{
														Content: "b",
													},
												},
											},
											&types.StringElement{
												Content: " ",
											},
											&types.QuotedText{
												Kind: types.SingleQuoteMonospace,
												Attributes: types.Attributes{
													types.AttrRoles: types.Roles{"myrole3"},
												},
												Elements: []interface{}{
													&types.StringElement{
														Content: "c",
													},
												},
											},
										},
									},
									Location: &types.Location{
										Scheme: "https://",
										Path:   "example.com",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("in bold text", func() {
				source := `a link to *https://example.com[]*`

				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "a link to "},
								&types.QuotedText{
									Kind: types.SingleQuoteBold,
									Elements: []interface{}{
										&types.InlineLink{
											Location: &types.Location{
												Scheme: "https://",
												Path:   "example.com",
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

			It("with special characters", func() {
				source := "a link to https://foo*_.com"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "a link to "},
								&types.InlineLink{
									Location: &types.Location{
										Scheme: "https://",
										Path:   "foo*_.com",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("in bold text", func() {
				source := `a link to *https://example.com[]*`

				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "a link to "},
								&types.QuotedText{
									Kind: types.SingleQuoteBold,
									Elements: []interface{}{
										&types.InlineLink{
											Location: &types.Location{
												Scheme: "https://",
												Path:   "example.com",
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

			It("in italic text", func() {
				source := `a link to _https://example.com[]_`

				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "a link to "},
								&types.QuotedText{
									Kind: types.SingleQuoteItalic,
									Elements: []interface{}{
										&types.InlineLink{
											Location: &types.Location{
												Scheme: "https://",
												Path:   "example.com",
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

			Context("with document attribute substitutions", func() {

				It("with a document attribute substitution for the whole URL", func() {
					source := `:url: https://example.com
:url: https://foo2.bar
	
a link to {url}`

					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name:  "url",
										Value: "https://example.com",
									},
									&types.AttributeDeclaration{
										Name:  "url",
										Value: "https://foo2.bar",
									},
								},
							},
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "a link to ",
									},
									&types.InlineLink{
										Location: &types.Location{
											Scheme: "https://",
											Path:   "foo2.bar",
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("with two document attribute substitutions only", func() {
					source := `:scheme: https
:path: example.com
	
a link to {scheme}://{path} and https://foo.com`

					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name:  "scheme",
										Value: "https",
									},
									&types.AttributeDeclaration{
										Name:  "path",
										Value: "example.com",
									},
								},
							},
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{Content: "a link to "},
									&types.InlineLink{
										Location: &types.Location{
											Scheme: "https://",
											Path:   "example.com",
										},
									},
									&types.StringElement{Content: " and "},
									&types.InlineLink{
										Location: &types.Location{
											Scheme: "https://",
											Path:   "foo.com",
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("with two document attribute substitutions in bold text", func() {
					source := `:scheme: https
:path: example.com
	
a link to *{scheme}://{path}[] and https://foo.com[]*`

					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name:  "scheme",
										Value: "https",
									},
									&types.AttributeDeclaration{
										Name:  "path",
										Value: "example.com",
									},
								},
							},
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "a link to ",
									},
									&types.QuotedText{
										Kind: types.SingleQuoteBold,
										Elements: []interface{}{
											&types.InlineLink{
												Location: &types.Location{
													Scheme: "https://",
													Path:   "example.com",
												},
											},
											&types.StringElement{
												Content: " and ",
											},
											&types.InlineLink{
												Location: &types.Location{
													Scheme: "https://",
													Path:   "foo.com",
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

				It("with two document attribute substitutions and a reset", func() {
					source := `:scheme: https
:path: example.com
	
:!path:
	
a link to {scheme}://{path} and https://foo.com`

					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name:  "scheme",
										Value: "https",
									},
									&types.AttributeDeclaration{
										Name:  "path",
										Value: "example.com",
									},
									&types.AttributeReset{
										Name: "path",
									},
								},
							},
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{Content: "a link to "},
									&types.InlineLink{
										Location: &types.Location{
											Scheme: "https://",
											Path:   "{path}", // substitution failed at during parsing
										},
									},
									&types.StringElement{Content: " and "},
									&types.InlineLink{
										Location: &types.Location{
											Scheme: "https://",
											Path:   "foo.com",
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("with document attribute in section 0 title", func() {
					source := `= a title to {scheme}://{path} and https://foo.com
:scheme: https
:path: example.com`

					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: []interface{}{
									&types.StringElement{Content: "a title to "},
									&types.InlineLink{
										Location: &types.Location{
											Scheme: "https://",
											Path:   "example.com",
										},
									},
									&types.StringElement{Content: " and "},
									&types.InlineLink{
										Location: &types.Location{
											Scheme: "https://",
											Path:   "foo.com",
										},
									},
								},
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name:  "scheme",
										Value: "https",
									},
									&types.AttributeDeclaration{
										Name:  "path",
										Value: "example.com",
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("with document attribute in section 1 title", func() {
					source := `:scheme: https
:path: example.com
	
== a title to {scheme}://{path} and https://foo.com`

					title := []interface{}{
						&types.StringElement{
							Content: "a title to ",
						},
						&types.InlineLink{
							Location: &types.Location{
								Scheme: "https://",
								Path:   "example.com",
							},
						},
						&types.StringElement{Content: " and "},
						&types.InlineLink{
							Location: &types.Location{
								Scheme: "https://",
								Path:   "foo.com",
							},
						},
					}
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name:  "scheme",
										Value: "https",
									},
									&types.AttributeDeclaration{
										Name:  "path",
										Value: "example.com",
									},
								},
							},
							&types.Section{
								Level: 1,
								Attributes: types.Attributes{
									types.AttrID: "_a_title_to_httpsexample_com_and_httpsfoo_com",
								},
								Title: title,
							},
						},
						ElementReferences: types.ElementReferences{
							"_a_title_to_httpsexample_com_and_httpsfoo_com": title,
						},
						TableOfContents: &types.TableOfContents{
							MaxDepth: 2,
							Sections: []*types.ToCSection{
								{
									ID:    "_a_title_to_httpsexample_com_and_httpsfoo_com",
									Level: 1,
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})
			})
		})

		Context("relative links", func() {

			It("to doc without text", func() {
				source := "a link to link:foo.adoc[]"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "a link to "},
								&types.InlineLink{
									Location: &types.Location{
										Scheme: "",
										Path:   "foo.adoc",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("to doc with text", func() {
				source := "a link to link:foo.adoc[foo doc]"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "a link to "},
								&types.InlineLink{
									Location: &types.Location{
										Scheme: "",
										Path:   "foo.adoc",
									},
									Attributes: types.Attributes{
										types.AttrInlineLinkText: "foo doc",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("to external URL with text only", func() {
				source := "a link to link:https://example.com[foo doc]"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "a link to "},
								&types.InlineLink{
									Location: &types.Location{
										Scheme: "https://",
										Path:   "example.com",
									},
									Attributes: types.Attributes{
										types.AttrInlineLinkText: "foo doc",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("to external URL with text and extra attributes", func() {
				source := "a link to link:https://example.com[foo doc, foo=bar]"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "a link to "},
								&types.InlineLink{
									Location: &types.Location{
										Scheme: "https://",
										Path:   "example.com",
									},
									Attributes: types.Attributes{
										types.AttrInlineLinkText: "foo doc",
										"foo":                    "bar",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("to external URL with extra attributes only", func() {
				source := "a link to link:https://example.com[foo=bar]"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "a link to "},
								&types.InlineLink{
									Location: &types.Location{
										Scheme: "https://",
										Path:   "example.com",
									},
									Attributes: types.Attributes{
										"foo": "bar",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("invalid syntax", func() {
				source := "a link to link:foo.adoc"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "a link to link:foo.adoc",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with quoted text attribute", func() {
				source := "link:/[a _a_ b *b* c `c`]"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.InlineLink{
									Location: &types.Location{
										Scheme: "",
										Path:   "/",
									},
									Attributes: types.Attributes{
										types.AttrInlineLinkText: []interface{}{
											&types.StringElement{
												Content: "a ",
											},
											&types.QuotedText{
												Kind: types.SingleQuoteItalic,
												Elements: []interface{}{
													&types.StringElement{
														Content: "a",
													},
												},
											},
											&types.StringElement{
												Content: " b ",
											},
											&types.QuotedText{
												Kind: types.SingleQuoteBold,
												Elements: []interface{}{
													&types.StringElement{
														Content: "b",
													},
												},
											},
											&types.StringElement{
												Content: " c ",
											},
											&types.QuotedText{
												Kind: types.SingleQuoteMonospace,
												Elements: []interface{}{
													&types.StringElement{
														Content: "c",
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

			It("with all valid characters", func() {
				source := `a link to link:ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789~:/?#@!$&;=()*+,-_.%[as expected]`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "a link to "},
								&types.InlineLink{
									Location: &types.Location{
										Scheme: "",
										Path:   "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789~:/?#@!$&;=()*+,-_.%",
									},
									Attributes: types.Attributes{
										types.AttrInlineLinkText: "as expected",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with encoded space", func() {
				source := `Test 1: link:/test/a b[with space]
Test 2: link:/test/a%20b[with encoded space]`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "Test 1: link:/test/a b[with space]\nTest 2: ",
								},
								&types.InlineLink{
									Location: &types.Location{
										Scheme: "",
										Path:   "/test/a%20b",
									},
									Attributes: types.Attributes{
										types.AttrInlineLinkText: "with encoded space",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with two document attribute substitutions and a reset", func() {
				source := `
:scheme: link
:path: example.com

:!path:

a link to {scheme}:{path}[] and https://foo.com`

				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Elements: []interface{}{
								&types.AttributeDeclaration{
									Name:  "scheme",
									Value: "link",
								},
								&types.AttributeDeclaration{
									Name:  "path",
									Value: "example.com",
								},
								&types.AttributeReset{
									Name: "path",
								},
							},
						},
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "a link to ",
								},
								&types.InlineLink{
									Location: &types.Location{
										Path: "{path}", // substitution failed
									},
								},
								&types.StringElement{
									Content: " and ",
								},
								&types.InlineLink{
									Location: &types.Location{
										Scheme: "https://",
										Path:   "foo.com",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("within quoted text", func() {
				source := "*link:foo[]*"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteBold,
									Elements: []interface{}{
										&types.InlineLink{
											Location: &types.Location{
												Path: "foo",
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

			It("with underscores", func() {
				source := "link:a_[A] link:a_[A]"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.InlineLink{
									Attributes: types.Attributes{
										types.AttrInlineLinkText: "A",
									},
									Location: &types.Location{
										Path: "a_",
									},
								},
								&types.StringElement{
									Content: " ",
								},
								&types.InlineLink{
									Attributes: types.Attributes{
										types.AttrInlineLinkText: "A",
									},
									Location: &types.Location{
										Path: "a_",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with line breaks in attributes", func() {
				source := `link:x[
title]`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.InlineLink{
									Attributes: types.Attributes{
										types.AttrInlineLinkText: "title",
									},
									Location: &types.Location{
										Path: "x",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			Context("text attribute with comma", func() {

				It("with text having comma", func() {
					source := `a link to link:https://example.com[A, B, and C]`
					expected := &types.Document{
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{Content: "a link to "},
									&types.InlineLink{
										Location: &types.Location{
											Scheme: "https://",
											Path:   "example.com",
										},
										Attributes: types.Attributes{
											types.AttrInlineLinkText: "A",
											types.AttrPositional2:    "B",
											types.AttrPositional3:    "and C",
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("with doublequoted text having comma", func() {
					source := `a link to link:https://example.com["A, B, and C"]`
					expected := &types.Document{
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{Content: "a link to "},
									&types.InlineLink{
										Location: &types.Location{
											Scheme: "https://",
											Path:   "example.com",
										},
										Attributes: types.Attributes{
											types.AttrInlineLinkText: "A, B, and C",
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("with doublequoted text having comma and other attrs", func() {
					source := `a link to link:https://example.com["A, B, and C", role=foo]`
					expected := &types.Document{
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{Content: "a link to "},
									&types.InlineLink{
										Location: &types.Location{
											Scheme: "https://",
											Path:   "example.com",
										},
										Attributes: types.Attributes{
											types.AttrInlineLinkText: "A, B, and C",
											types.AttrRoles:          types.Roles{"foo"},
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("with text having comma and other attributes", func() {
					source := `a link to link:https://example.com[A, B, and C, role=foo]`
					expected := &types.Document{
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{Content: "a link to "},
									&types.InlineLink{
										Location: &types.Location{
											Scheme: "https://",
											Path:   "example.com",
										},
										Attributes: types.Attributes{
											types.AttrInlineLinkText: "A",
											types.AttrPositional2:    "B",
											types.AttrPositional3:    "and C",
											types.AttrRoles:          types.Roles{"foo"},
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

		Context("inline anchors", func() {

			It("opening a paragraph", func() {
				source := `[[bookmark-a]]Inline anchors make arbitrary content referenceable.`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.InlineLink{
									Attributes: types.Attributes{
										types.AttrID: "bookmark-a",
									},
								},
								&types.StringElement{
									Content: "Inline anchors make arbitrary content referenceable.",
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
