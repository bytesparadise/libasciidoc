package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("links", func() {

	Context("draft document", func() {

		Context("external links", func() {

			It("external link without text", func() {
				source := "a link to https://foo.bar"
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{Content: "a link to "},
							types.InlineLink{
								Location: types.Location{
									Elements: []interface{}{
										types.StringElement{
											Content: "https://foo.bar",
										},
									},
								},
								Attributes: types.ElementAttributes{},
							},
						},
					},
				}
				Expect(ParseDocumentBlock(source)).To(Equal(expected))
			})

			It("external link with empty text", func() {
				source := "a link to https://foo.bar[]"
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{Content: "a link to "},
							types.InlineLink{
								Location: types.Location{
									Elements: []interface{}{
										types.StringElement{
											Content: "https://foo.bar",
										},
									},
								},
								Attributes: types.ElementAttributes{},
							},
						},
					},
				}
				Expect(ParseDocumentBlock(source)).To(Equal(expected))
			})

			It("external link with text only", func() {
				source := "a link to mailto:foo@bar[the foo@bar email]."
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{Content: "a link to "},
							types.InlineLink{
								Location: types.Location{
									Elements: []interface{}{
										types.StringElement{
											Content: "mailto:foo@bar",
										},
									},
								},
								Attributes: types.ElementAttributes{
									types.AttrInlineLinkText: []interface{}{
										types.StringElement{
											Content: "the foo@bar email",
										},
									},
								},
							},
							types.StringElement{Content: "."},
						},
					},
				}
				Expect(ParseDocumentBlock(source)).To(Equal(expected))
			})

			It("external link with text and extra attributes", func() {
				source := "a link to mailto:foo@bar[the foo@bar email, foo=bar]"
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{Content: "a link to "},
							types.InlineLink{
								Location: types.Location{
									Elements: []interface{}{
										types.StringElement{
											Content: "mailto:foo@bar",
										},
									},
								},
								Attributes: types.ElementAttributes{
									types.AttrInlineLinkText: []interface{}{
										types.StringElement{
											Content: "the foo@bar email",
										},
									},
									"foo": "bar",
								},
							},
						},
					},
				}
				Expect(ParseDocumentBlock(source)).To(Equal(expected))
			})

			It("external link inside a multiline paragraph -  without attributes", func() {
				source := `a http://website.com
and more text on the
next lines`

				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "a ",
							},
							types.InlineLink{
								Attributes: types.ElementAttributes{},
								Location: types.Location{
									Elements: []interface{}{
										types.StringElement{
											Content: "http://website.com",
										},
									},
								},
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
				}
				Expect(ParseDocumentBlock(source)).To(Equal(expected))
			})

			It("external link inside a multiline paragraph -  with attributes", func() {
				source := `a http://website.com[]
and more text on the
next lines`

				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "a ",
							},
							types.InlineLink{
								Attributes: types.ElementAttributes{},
								Location: types.Location{
									Elements: []interface{}{
										types.StringElement{
											Content: "http://website.com",
										},
									},
								},
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
				}
				Expect(ParseDocumentBlock(source)).To(Equal(expected))
			})

			It("external link with more text afterwards", func() {
				source := `a link to https://foo.bar and more text`
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{Content: "a link to "},
							types.InlineLink{
								Location: types.Location{
									Elements: []interface{}{
										types.StringElement{
											Content: "https://foo.bar",
										},
									},
								},
								Attributes: types.ElementAttributes{},
							},
							types.StringElement{Content: " and more text"},
						},
					},
				}
				Expect(ParseDocumentBlock(source)).To(Equal(expected))
			})

			Context("text attribute with comma", func() {

				It("external link only with text having comma", func() {
					source := `a link to http://website.com[A, B, and C]`
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "a link to "},
								types.InlineLink{
									Location: types.Location{
										Elements: []interface{}{
											types.StringElement{
												Content: "http://website.com",
											},
										},
									},
									Attributes: types.ElementAttributes{
										types.AttrInlineLinkText: []interface{}{
											types.StringElement{
												Content: "A, B, and C",
											},
										},
									},
								},
							},
						},
					}
					Expect(ParseDocumentBlock(source)).To(Equal(expected))
				})

				It("external link only with doublequoted text having comma", func() {
					source := `a link to http://website.com["A, B, and C"]`
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "a link to "},
								types.InlineLink{
									Location: types.Location{
										Elements: []interface{}{
											types.StringElement{
												Content: "http://website.com",
											},
										},
									},
									Attributes: types.ElementAttributes{
										types.AttrInlineLinkText: []interface{}{
											types.StringElement{
												Content: "A, B, and C",
											},
										},
									},
								},
							},
						},
					}
					Expect(ParseDocumentBlock(source)).To(Equal(expected))
				})

				It("external link with doublequoted text having comma and other attrs", func() {
					source := `a link to http://website.com["A, B, and C", role=foo]`
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "a link to "},
								types.InlineLink{
									Location: types.Location{
										Elements: []interface{}{
											types.StringElement{
												Content: "http://website.com",
											},
										},
									},
									Attributes: types.ElementAttributes{
										types.AttrInlineLinkText: []interface{}{
											types.StringElement{
												Content: "A, B, and C",
											},
										},
										"role": "foo",
									},
								},
							},
						},
					}
					Expect(ParseDocumentBlock(source)).To(Equal(expected))
				})

				It("external link with text having comma and other attributes", func() {
					source := `a link to http://website.com[A, B, and C, role=foo]`
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "a link to "},
								types.InlineLink{
									Location: types.Location{
										Elements: []interface{}{
											types.StringElement{
												Content: "http://website.com",
											},
										},
									},
									Attributes: types.ElementAttributes{
										types.AttrInlineLinkText: []interface{}{
											types.StringElement{
												Content: "A",
											},
										},
										"B":     nil,
										"and C": nil,
										"role":  "foo",
									},
								},
							},
						},
					}
					Expect(ParseDocumentBlock(source)).To(Equal(expected))
				})
			})

			It("external link with special characters", func() {
				source := "a link to https://foo*_.com"

				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a link to "},
									types.InlineLink{
										Attributes: types.ElementAttributes{},
										Location: types.Location{
											Elements: []interface{}{
												types.StringElement{
													Content: "https://foo*_.com",
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

			It("external link in bold text", func() {
				source := `a link to *https://foo.com[]*`

				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a link to "},
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.InlineLink{
												Attributes: types.ElementAttributes{},
												Location: types.Location{
													Elements: []interface{}{
														types.StringElement{
															Content: "https://foo.com",
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
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

		})

		Context("relative links", func() {

			It("relative link to doc without text", func() {
				source := "a link to link:foo.adoc[]"
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{Content: "a link to "},
							types.InlineLink{
								Location: types.Location{
									Elements: []interface{}{
										types.StringElement{
											Content: "foo.adoc",
										},
									},
								},
								Attributes: types.ElementAttributes{},
							},
						},
					},
				}
				Expect(ParseDocumentBlock(source)).To(Equal(expected))
			})

			It("relative link to doc with text", func() {
				source := "a link to link:foo.adoc[foo doc]"
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{Content: "a link to "},
							types.InlineLink{
								Location: types.Location{
									Elements: []interface{}{
										types.StringElement{
											Content: "foo.adoc",
										},
									},
								},
								Attributes: types.ElementAttributes{
									types.AttrInlineLinkText: []interface{}{
										types.StringElement{
											Content: "foo doc",
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocumentBlock(source)).To(Equal(expected))
			})

			It("relative link to external URL with text only", func() {
				source := "a link to link:https://foo.bar[foo doc]"
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{Content: "a link to "},
							types.InlineLink{
								Location: types.Location{
									Elements: []interface{}{
										types.StringElement{
											Content: "https://foo.bar",
										},
									},
								},
								Attributes: types.ElementAttributes{
									types.AttrInlineLinkText: []interface{}{
										types.StringElement{
											Content: "foo doc",
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocumentBlock(source)).To(Equal(expected))
			})

			It("relative link to external URL with text and extra attributes", func() {
				source := "a link to link:https://foo.bar[foo doc, foo=bar]"
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{Content: "a link to "},
							types.InlineLink{
								Location: types.Location{
									Elements: []interface{}{
										types.StringElement{
											Content: "https://foo.bar",
										},
									},
								},
								Attributes: types.ElementAttributes{
									types.AttrInlineLinkText: []interface{}{
										types.StringElement{
											Content: "foo doc",
										},
									},
									"foo": "bar",
								},
							},
						},
					},
				}
				Expect(ParseDocumentBlock(source)).To(Equal(expected))
			})

			It("relative link to external URL with extra attributes only", func() {
				source := "a link to link:https://foo.bar[foo=bar]"
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{Content: "a link to "},
							types.InlineLink{
								Location: types.Location{
									Elements: []interface{}{
										types.StringElement{
											Content: "https://foo.bar",
										},
									},
								},
								Attributes: types.ElementAttributes{
									"foo": "bar",
								},
							},
						},
					},
				}
				Expect(ParseDocumentBlock(source)).To(Equal(expected))
			})

			It("invalid relative link to doc", func() {
				source := "a link to link:foo.adoc"
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "a link to link:foo.adoc",
							},
						},
					},
				}
				Expect(ParseDocumentBlock(source)).To(Equal(expected))
			})

			It("relative link with quoted text attribute", func() {
				source := "link:/[a _a_ b *b* c `c`]"
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.InlineLink{
								Location: types.Location{
									Elements: []interface{}{
										types.StringElement{
											Content: "/",
										},
									},
								},
								Attributes: types.ElementAttributes{
									types.AttrInlineLinkText: []interface{}{
										types.StringElement{
											Content: "a ",
										},
										types.QuotedText{
											Kind: types.Italic,
											Elements: []interface{}{
												types.StringElement{
													Content: "a",
												},
											},
										},
										types.StringElement{
											Content: " b ",
										},
										types.QuotedText{
											Kind: types.Bold,
											Elements: []interface{}{
												types.StringElement{
													Content: "b",
												},
											},
										},
										types.StringElement{
											Content: " c ",
										},
										types.QuotedText{
											Kind: types.Monospace,
											Elements: []interface{}{
												types.StringElement{
													Content: "c",
												},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocumentBlock(source)).To(Equal(expected))
			})

			It("relative link within quoted text", func() {
				source := "*link:foo[]*"
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.QuotedText{
								Kind: types.Bold,
								Elements: []interface{}{
									types.InlineLink{
										Location: types.Location{
											Elements: []interface{}{
												types.StringElement{
													Content: "foo",
												},
											},
										},
										Attributes: types.ElementAttributes{},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocumentBlock(source)).To(Equal(expected))
			})

			It("relative link with all valid characters", func() {
				source := `a link to link:ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789~:/?#@!$&;=()*+,-_.%[as expected]`
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{Content: "a link to "},
							types.InlineLink{
								Location: types.Location{
									Elements: []interface{}{
										types.StringElement{
											Content: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789~:/?#@!$&;=()*+,-_.%",
										},
									},
								},
								Attributes: types.ElementAttributes{
									types.AttrInlineLinkText: []interface{}{
										types.StringElement{
											Content: "as expected",
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocumentBlock(source)).To(Equal(expected))
			})

			It("relative link with encoded space", func() {
				source := `Test 1: link:/test/a b[with space]
Test 2: link:/test/a%20b[with encoded space]`
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{Content: "Test 1: link:/test/a b[with space]"},
						},
						{
							types.StringElement{Content: "Test 2: "},
							types.InlineLink{
								Location: types.Location{
									Elements: []interface{}{
										types.StringElement{
											Content: "/test/a%20b",
										},
									},
								},
								Attributes: types.ElementAttributes{
									types.AttrInlineLinkText: []interface{}{
										types.StringElement{
											Content: "with encoded space",
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocumentBlock(source)).To(Equal(expected))
			})

			It("relative link with two document attribute substitutions and a reset", func() {
				source := `
:scheme: link
:path: foo.bar

:!path:

a link to {scheme}:{path}[] and https://foo.baz`

				expected := types.Document{
					Attributes: types.DocumentAttributes{
						"scheme": "link",
					},
					ElementReferences: types.ElementReferences{},
					Footnotes:         []types.Footnote{},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a link to "},
									types.InlineLink{
										Attributes: types.ElementAttributes{},
										Location: types.Location{
											Elements: []interface{}{
												types.StringElement{
													Content: "{path}",
												},
											},
										},
									},
									types.StringElement{Content: " and "},
									types.InlineLink{
										Attributes: types.ElementAttributes{},
										Location: types.Location{
											Elements: []interface{}{
												types.StringElement{
													Content: "https://foo.baz",
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

			Context("text attribute with comma", func() {

				It("relative link only with text having comma", func() {
					source := `a link to link:https://foo.bar[A, B, and C]`
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "a link to "},
								types.InlineLink{
									Location: types.Location{
										Elements: []interface{}{
											types.StringElement{
												Content: "https://foo.bar",
											},
										},
									},
									Attributes: types.ElementAttributes{
										types.AttrInlineLinkText: []interface{}{
											types.StringElement{
												Content: "A, B, and C",
											},
										},
									},
								},
							},
						},
					}
					Expect(ParseDocumentBlock(source)).To(Equal(expected))
				})

				It("relative link only with doublequoted text having comma", func() {
					source := `a link to link:https://foo.bar["A, B, and C"]`
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "a link to "},
								types.InlineLink{
									Location: types.Location{
										Elements: []interface{}{
											types.StringElement{
												Content: "https://foo.bar",
											},
										},
									},
									Attributes: types.ElementAttributes{
										types.AttrInlineLinkText: []interface{}{
											types.StringElement{
												Content: "A, B, and C",
											},
										},
									},
								},
							},
						},
					}
					Expect(ParseDocumentBlock(source)).To(Equal(expected))
				})

				It("relative link with doublequoted text having comma and other attrs", func() {
					source := `a link to link:https://foo.bar["A, B, and C", role=foo]`
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "a link to "},
								types.InlineLink{
									Location: types.Location{
										Elements: []interface{}{
											types.StringElement{
												Content: "https://foo.bar",
											},
										},
									},
									Attributes: types.ElementAttributes{
										types.AttrInlineLinkText: []interface{}{
											types.StringElement{
												Content: "A, B, and C",
											},
										},
										"role": "foo",
									},
								},
							},
						},
					}
					Expect(ParseDocumentBlock(source)).To(Equal(expected))
				})

				It("relative link with text having comma and other attributes", func() {
					source := `a link to link:https://foo.bar[A, B, and C, role=foo]`
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "a link to "},
								types.InlineLink{
									Location: types.Location{
										Elements: []interface{}{
											types.StringElement{
												Content: "https://foo.bar",
											},
										},
									},
									Attributes: types.ElementAttributes{
										types.AttrInlineLinkText: []interface{}{
											types.StringElement{
												Content: "A",
											},
										},
										"B":     nil,
										"and C": nil,
										"role":  "foo",
									},
								},
							},
						},
					}
					Expect(ParseDocumentBlock(source)).To(Equal(expected))
				})

			})

		})

		Context("with document attribute substitutions", func() {

			It("external link with a document attribute substitution for the whole URL", func() {
				source := `:url: https://foo.bar

a link to {url}`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DocumentAttributeDeclaration{
							Name:  "url",
							Value: "https://foo.bar",
						},
						types.BlankLine{},
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a link to "},
									types.DocumentAttributeSubstitution{
										Name: "url",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("external link with a document attribute substitution for the whole URL", func() {
				source := `:url: https://foo.bar

a link to {url}`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DocumentAttributeDeclaration{
							Name:  "url",
							Value: "https://foo.bar",
						},
						types.BlankLine{},
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a link to "},
									types.DocumentAttributeSubstitution{
										Name: "url",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("external link with two document attribute substitutions", func() {
				source := `:scheme: https
:path: foo.bar

a link to {scheme}://{path}`

				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DocumentAttributeDeclaration{
							Name:  "scheme",
							Value: "https",
						},
						types.DocumentAttributeDeclaration{
							Name:  "path",
							Value: "foo.bar",
						},
						types.BlankLine{},
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a link to ",
									},
									types.DocumentAttributeSubstitution{
										Name: "scheme",
									},
									types.StringElement{
										Content: "://",
									},
									types.DocumentAttributeSubstitution{
										Name: "path",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("external link with two document attribute substitutions in bold text", func() {
				source := `
:scheme: https
:path: foo.bar

a link to *{scheme}://{path}[] and https://foo.baz[]*`

				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.BlankLine{},
						types.DocumentAttributeDeclaration{
							Name:  "scheme",
							Value: "https",
						},
						types.DocumentAttributeDeclaration{
							Name:  "path",
							Value: "foo.bar",
						},
						types.BlankLine{},
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a link to "},
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.DocumentAttributeSubstitution{
												Name: "scheme",
											},
											types.StringElement{
												Content: "://",
											},
											types.DocumentAttributeSubstitution{
												Name: "path",
											},
											types.StringElement{Content: "[] and "},
											types.InlineLink{
												Attributes: types.ElementAttributes{},
												Location: types.Location{
													Elements: []interface{}{
														types.StringElement{
															Content: "https://foo.baz",
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
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})
		})
	})

	Context("final document", func() {

		It("relative link within quoted text", func() {
			source := "*link:foo[]*"
			expected := types.Document{
				Attributes:        types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{},
				Footnotes:         []types.Footnote{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.QuotedText{
									Kind: types.Bold,
									Elements: []interface{}{
										types.InlineLink{
											Location: types.Location{
												Elements: []interface{}{
													types.StringElement{
														Content: "foo",
													},
												},
											},
											Attributes: types.ElementAttributes{},
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

		It("external link with special characters", func() {
			source := "a link to https://foo*_.com"
			expected := types.Document{
				Attributes:        types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{},
				Footnotes:         []types.Footnote{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "a link to "},
								types.InlineLink{
									Attributes: types.ElementAttributes{},
									Location: types.Location{
										Elements: []interface{}{
											types.StringElement{
												Content: "https://foo*_.com",
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

		It("external link in bold text", func() {
			source := `a link to *https://foo.com[]*`

			expected := types.Document{
				Attributes:        types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{},
				Footnotes:         []types.Footnote{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "a link to "},
								types.QuotedText{
									Kind: types.Bold,
									Elements: []interface{}{
										types.InlineLink{
											Attributes: types.ElementAttributes{},
											Location: types.Location{
												Elements: []interface{}{
													types.StringElement{
														Content: "https://foo.com",
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

		It("external link in italic text", func() {
			source := `a link to _https://foo.com[]_`

			expected := types.Document{
				Attributes:        types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{},
				Footnotes:         []types.Footnote{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "a link to "},
								types.QuotedText{
									Kind: types.Italic,
									Elements: []interface{}{
										types.InlineLink{
											Attributes: types.ElementAttributes{},
											Location: types.Location{
												Elements: []interface{}{
													types.StringElement{
														Content: "https://foo.com",
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

		Context("with document attribute substitutions", func() {

			It("external link with a document attribute substitution for the whole URL", func() {
				source := `
:url: https://foo.bar
:url: https://foo2.bar

a link to {url}`

				expected := types.Document{
					Attributes: types.DocumentAttributes{
						"url": "https://foo2.bar", // overriden by second declaration
					},
					ElementReferences: types.ElementReferences{},
					Footnotes:         []types.Footnote{},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a link to "},
									types.InlineLink{
										Location: types.Location{
											Elements: []interface{}{
												types.StringElement{
													Content: "https://foo2.bar",
												},
											},
										},
										Attributes: types.ElementAttributes{},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("external link with two document attribute substitutions only", func() {
				source := `
:scheme: https
:path: foo.bar

a link to {scheme}://{path} and https://foo.baz`

				expected := types.Document{
					Attributes: types.DocumentAttributes{
						"scheme": "https",
						"path":   "foo.bar",
					},
					ElementReferences: types.ElementReferences{},
					Footnotes:         []types.Footnote{},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a link to "},
									types.InlineLink{
										Location: types.Location{
											Elements: []interface{}{
												types.StringElement{
													Content: "https://foo.bar",
												},
											},
										},
										Attributes: types.ElementAttributes{},
									},
									types.StringElement{Content: " and "},
									types.InlineLink{
										Attributes: types.ElementAttributes{},
										Location: types.Location{
											Elements: []interface{}{
												types.StringElement{
													Content: "https://foo.baz",
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

			It("external link with two document attribute substitutions in bold text", func() {
				source := `
:scheme: https
:path: foo.bar

a link to *{scheme}://{path}[] and https://foo.baz[]*`

				expected := types.Document{
					Attributes: types.DocumentAttributes{
						"scheme": "https",
						"path":   "foo.bar",
					},
					ElementReferences: types.ElementReferences{},
					Footnotes:         []types.Footnote{},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a link to "},
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.InlineLink{
												Location: types.Location{
													Elements: []interface{}{
														types.StringElement{
															Content: "https://foo.bar",
														},
													},
												},
												Attributes: types.ElementAttributes{},
											},
											types.StringElement{Content: " and "},
											types.InlineLink{
												Attributes: types.ElementAttributes{},
												Location: types.Location{
													Elements: []interface{}{
														types.StringElement{
															Content: "https://foo.baz",
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

			It("external link with two document attribute substitutions and a reset", func() {
				source := `
:scheme: https
:path: foo.bar

:!path:

a link to {scheme}://{path} and https://foo.baz`

				expected := types.Document{
					Attributes: types.DocumentAttributes{
						"scheme": "https",
					},
					ElementReferences: types.ElementReferences{},
					Footnotes:         []types.Footnote{},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a link to "},
									types.InlineLink{
										Attributes: types.ElementAttributes{},
										Location: types.Location{
											Elements: []interface{}{
												types.StringElement{
													Content: "https://{path}",
												},
											},
										},
									},
									types.StringElement{Content: " and "},
									types.InlineLink{
										Attributes: types.ElementAttributes{},
										Location: types.Location{
											Elements: []interface{}{
												types.StringElement{
													Content: "https://foo.baz",
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

			It("external link with document attribute in section 0 title", func() {
				source := `= a title to {scheme}://{path} and https://foo.baz
:scheme: https
:path: foo.bar`

				title := []interface{}{
					types.StringElement{Content: "a title to "},
					types.InlineLink{
						Attributes: types.ElementAttributes{},
						Location: types.Location{
							Elements: []interface{}{
								types.StringElement{
									Content: "https://foo.bar",
								},
							},
						},
					},
					types.StringElement{Content: " and "},
					types.InlineLink{
						Attributes: types.ElementAttributes{},
						Location: types.Location{
							Elements: []interface{}{
								types.StringElement{
									Content: "https://foo.baz",
								},
							},
						},
					},
				}
				expected := types.Document{
					Attributes: types.DocumentAttributes{
						"scheme": "https",
						"path":   "foo.bar",
					},
					ElementReferences: types.ElementReferences{
						"_a_title_to_https_foo_bar_and_https_foo_baz": title,
					},
					Footnotes: []types.Footnote{},
					Elements: []interface{}{
						types.Section{
							Level: 0,
							Attributes: types.ElementAttributes{
								types.AttrID: "_a_title_to_https_foo_bar_and_https_foo_baz",
							},
							Title:    title,
							Elements: []interface{}{},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("external link with document attribute in section 1 title", func() {
				source := `:scheme: https
:path: foo.bar

== a title to {scheme}://{path} and https://foo.baz`

				title := []interface{}{
					types.StringElement{Content: "a title to "},
					types.InlineLink{
						Attributes: types.ElementAttributes{},
						Location: types.Location{
							Elements: []interface{}{
								types.StringElement{
									Content: "https://foo.bar",
								},
							},
						},
					},
					types.StringElement{Content: " and "},
					types.InlineLink{
						Attributes: types.ElementAttributes{},
						Location: types.Location{
							Elements: []interface{}{
								types.StringElement{
									Content: "https://foo.baz",
								},
							},
						},
					},
				}
				expected := types.Document{
					Attributes: types.DocumentAttributes{
						"scheme": "https",
						"path":   "foo.bar",
					},
					ElementReferences: types.ElementReferences{
						"_a_title_to_https_foo_bar_and_https_foo_baz": title,
					},
					Footnotes: []types.Footnote{},
					Elements: []interface{}{
						types.Section{
							Level: 1,
							Attributes: types.ElementAttributes{
								types.AttrID: "_a_title_to_https_foo_bar_and_https_foo_baz",
							},
							Title:    title,
							Elements: []interface{}{},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})

	})
})
