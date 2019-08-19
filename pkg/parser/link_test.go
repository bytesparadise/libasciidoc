package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
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
			verifyDocumentBlock(expected, source)
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
			verifyDocumentBlock(expected, source)
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
			verifyDocumentBlock(expected, source)
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
			verifyDocumentBlock(expected, source)
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
			verifyDocumentBlock(expected, source)
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
			verifyDocumentBlock(expected, source)
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
			verifyDocumentBlock(expected, source)
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
			verifyDocumentBlock(expected, source)
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
			verifyDocumentBlock(expected, source)
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
			verifyDocumentBlock(expected, source)
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
			verifyDocumentBlock(expected, source)
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
			verifyDocumentBlock(expected, source)
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
			verifyDocumentBlock(expected, source)
		})

	})

})
