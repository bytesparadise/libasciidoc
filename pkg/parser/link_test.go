package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("links - preflight", func() {

	Context("external links", func() {

		It("external link without text", func() {
			source := "a link to https://foo.bar"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "a link to "},
						types.InlineLink{
							Location: types.Location{
								types.StringElement{
									Content: "https://foo.bar",
								},
							},
							Attributes: types.ElementAttributes{},
						},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("external link with empty text", func() {
			source := "a link to https://foo.bar[]"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "a link to "},
						types.InlineLink{
							Location: types.Location{
								types.StringElement{
									Content: "https://foo.bar",
								},
							},
							Attributes: types.ElementAttributes{},
						},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("external link with text only", func() {
			source := "a link to mailto:foo@bar[the foo@bar email]."
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "a link to "},
						types.InlineLink{
							Location: types.Location{
								types.StringElement{
									Content: "mailto:foo@bar",
								},
							},
							Attributes: types.ElementAttributes{
								types.AttrInlineLinkText: types.InlineElements{
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
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("external link with text and extra attributes", func() {
			source := "a link to mailto:foo@bar[the foo@bar email, foo=bar]"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "a link to "},
						types.InlineLink{
							Location: types.Location{
								types.StringElement{
									Content: "mailto:foo@bar",
								},
							},
							Attributes: types.ElementAttributes{
								types.AttrInlineLinkText: types.InlineElements{
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
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("external link inside a multiline paragraph -  without attributes", func() {
			source := `a http://website.com
and more text on the
next lines`

			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{
							Content: "a ",
						},
						types.InlineLink{
							Attributes: types.ElementAttributes{},
							Location: types.Location{
								types.StringElement{
									Content: "http://website.com",
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
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("external link inside a multiline paragraph -  with attributes", func() {
			source := `a http://website.com[]
and more text on the
next lines`

			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{
							Content: "a ",
						},
						types.InlineLink{
							Attributes: types.ElementAttributes{},
							Location: types.Location{
								types.StringElement{
									Content: "http://website.com",
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
			Expect(source).To(EqualDocumentBlock(expected))
		})

		Context("text attribute with comma", func() {

			It("relative link only with text having comma", func() {
				source := `a link to http://website.com[A, B, and C]`
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "a link to "},
							types.InlineLink{
								Location: types.Location{
									types.StringElement{
										Content: "http://website.com",
									},
								},
								Attributes: types.ElementAttributes{
									types.AttrInlineLinkText: types.InlineElements{
										types.StringElement{
											Content: "A, B, and C",
										},
									},
								},
							},
						},
					},
				}
				Expect(source).To(EqualDocumentBlock(expected))
			})

			It("relative link only with doublequoted text having comma", func() {
				source := `a link to http://website.com["A, B, and C"]`
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "a link to "},
							types.InlineLink{
								Location: types.Location{
									types.StringElement{
										Content: "http://website.com",
									},
								},
								Attributes: types.ElementAttributes{
									types.AttrInlineLinkText: types.InlineElements{
										types.StringElement{
											Content: "A, B, and C",
										},
									},
								},
							},
						},
					},
				}
				Expect(source).To(EqualDocumentBlock(expected))
			})

			It("relative link with doublequoted text having comma and other attrs", func() {
				source := `a link to http://website.com["A, B, and C", role=foo]`
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "a link to "},
							types.InlineLink{
								Location: types.Location{
									types.StringElement{
										Content: "http://website.com",
									},
								},
								Attributes: types.ElementAttributes{
									types.AttrInlineLinkText: types.InlineElements{
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
				Expect(source).To(EqualDocumentBlock(expected))
			})

			It("relative link with text having comma and other attributes", func() {
				source := `a link to http://website.com[A, B, and C, role=foo]`
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "a link to "},
							types.InlineLink{
								Location: types.Location{
									types.StringElement{
										Content: "http://website.com",
									},
								},
								Attributes: types.ElementAttributes{
									types.AttrInlineLinkText: types.InlineElements{
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
				Expect(source).To(EqualDocumentBlock(expected))
			})
		})

	})

	Context("relative links", func() {

		It("relative link to doc without text", func() {
			source := "a link to link:foo.adoc[]"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "a link to "},
						types.InlineLink{
							Location: types.Location{
								types.StringElement{
									Content: "foo.adoc",
								},
							},
							Attributes: types.ElementAttributes{},
						},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("relative link to doc with text", func() {
			source := "a link to link:foo.adoc[foo doc]"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "a link to "},
						types.InlineLink{
							Location: types.Location{
								types.StringElement{
									Content: "foo.adoc",
								},
							},
							Attributes: types.ElementAttributes{
								types.AttrInlineLinkText: types.InlineElements{
									types.StringElement{
										Content: "foo doc",
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("relative link to external URL with text only", func() {
			source := "a link to link:https://foo.bar[foo doc]"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "a link to "},
						types.InlineLink{
							Location: types.Location{
								types.StringElement{
									Content: "https://foo.bar",
								},
							},
							Attributes: types.ElementAttributes{
								types.AttrInlineLinkText: types.InlineElements{
									types.StringElement{
										Content: "foo doc",
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("relative link to external URL with text and extra attributes", func() {
			source := "a link to link:https://foo.bar[foo doc, foo=bar]"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "a link to "},
						types.InlineLink{
							Location: types.Location{
								types.StringElement{
									Content: "https://foo.bar",
								},
							},
							Attributes: types.ElementAttributes{
								types.AttrInlineLinkText: types.InlineElements{
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
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("relative link to external URL with extra attributes only", func() {
			source := "a link to link:https://foo.bar[foo=bar]"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "a link to "},
						types.InlineLink{
							Location: types.Location{
								types.StringElement{
									Content: "https://foo.bar",
								},
							},
							Attributes: types.ElementAttributes{
								"foo": "bar",
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("invalid relative link to doc", func() {
			source := "a link to link:foo.adoc"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{
							Content: "a link to link:foo.adoc",
						},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("relative link with quoted text", func() {
			source := "link:/[a _a_ b *b* c `c`]"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.InlineLink{
							Location: types.Location{
								types.StringElement{
									Content: "/",
								},
							},
							Attributes: types.ElementAttributes{
								types.AttrInlineLinkText: types.InlineElements{
									types.StringElement{
										Content: "a ",
									},
									types.QuotedText{
										Kind: types.Italic,
										Elements: types.InlineElements{
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
										Elements: types.InlineElements{
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
										Elements: types.InlineElements{
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
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("relative link with all valid characters", func() {
			source := `a link to link:ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789~:/?#@!$&;=()*+,-_.%[as expected]`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "a link to "},
						types.InlineLink{
							Location: types.Location{
								types.StringElement{
									Content: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789~:/?#@!$&;=()*+,-_.%",
								},
							},
							Attributes: types.ElementAttributes{
								types.AttrInlineLinkText: types.InlineElements{
									types.StringElement{
										Content: "as expected",
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("relative link with encoded space", func() {
			source := `Test 1: link:/test/a b[with space]
Test 2: link:/test/a%20b[with encoded space]`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "Test 1: link:/test/a b[with space]"},
					},
					{
						types.StringElement{Content: "Test 2: "},
						types.InlineLink{
							Location: types.Location{
								types.StringElement{
									Content: "/test/a%20b",
								},
							},
							Attributes: types.ElementAttributes{
								types.AttrInlineLinkText: types.InlineElements{
									types.StringElement{
										Content: "with encoded space",
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		Context("text attribute with comma", func() {

			It("relative link only with text having comma", func() {
				source := `a link to link:https://foo.bar[A, B, and C]`
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "a link to "},
							types.InlineLink{
								Location: types.Location{
									types.StringElement{
										Content: "https://foo.bar",
									},
								},
								Attributes: types.ElementAttributes{
									types.AttrInlineLinkText: types.InlineElements{
										types.StringElement{
											Content: "A, B, and C",
										},
									},
								},
							},
						},
					},
				}
				Expect(source).To(EqualDocumentBlock(expected))
			})

			It("relative link only with doublequoted text having comma", func() {
				source := `a link to link:https://foo.bar["A, B, and C"]`
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "a link to "},
							types.InlineLink{
								Location: types.Location{
									types.StringElement{
										Content: "https://foo.bar",
									},
								},
								Attributes: types.ElementAttributes{
									types.AttrInlineLinkText: types.InlineElements{
										types.StringElement{
											Content: "A, B, and C",
										},
									},
								},
							},
						},
					},
				}
				Expect(source).To(EqualDocumentBlock(expected))
			})

			It("relative link with doublequoted text having comma and other attrs", func() {
				source := `a link to link:https://foo.bar["A, B, and C", role=foo]`
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "a link to "},
							types.InlineLink{
								Location: types.Location{
									types.StringElement{
										Content: "https://foo.bar",
									},
								},
								Attributes: types.ElementAttributes{
									types.AttrInlineLinkText: types.InlineElements{
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
				Expect(source).To(EqualDocumentBlock(expected))
			})

			It("relative link with text having comma and other attributes", func() {
				source := `a link to link:https://foo.bar[A, B, and C, role=foo]`
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "a link to "},
							types.InlineLink{
								Location: types.Location{
									types.StringElement{
										Content: "https://foo.bar",
									},
								},
								Attributes: types.ElementAttributes{
									types.AttrInlineLinkText: types.InlineElements{
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
				Expect(source).To(EqualDocumentBlock(expected))
			})
		})

	})

})
