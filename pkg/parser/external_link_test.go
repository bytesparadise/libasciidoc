package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("links", func() {

	Context("external links", func() {

		It("external link without text", func() {
			actualContent := "a link to https://foo.bar"
			expectedResult := types.Paragraph{
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
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("external link with empty text", func() {
			actualContent := "a link to https://foo.bar[]"
			expectedResult := types.Paragraph{
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
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("external link with text", func() {
			actualContent := "a link to mailto:foo@bar[the foo@bar email]"
			expectedResult := types.Paragraph{
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
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("external link with text and extra attributes", func() {
			actualContent := "a link to mailto:foo@bar[the foo@bar email, foo=bar]"
			expectedResult := types.Paragraph{
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
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

	})

	Context("relative links", func() {

		It("relative link to doc without text", func() {
			actualContent := "a link to link:foo.adoc[]"
			expectedResult := types.InlineElements{
				types.StringElement{Content: "a link to "},
				types.InlineLink{
					Location: types.Location{
						types.StringElement{
							Content: "foo.adoc",
						},
					},
					Attributes: types.ElementAttributes{},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("relative link to doc with text", func() {
			actualContent := "a link to link:foo.adoc[foo doc]"
			expectedResult := types.InlineElements{
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
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("relative link to external URL with text", func() {
			actualContent := "a link to link:https://foo.bar[foo doc]"
			expectedResult := types.InlineElements{
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
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("relative link to external URL with text and extra attributes", func() {
			actualContent := "a link to link:https://foo.bar[foo doc, foo=bar]"
			expectedResult := types.InlineElements{
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
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("relative link to external URL with extra attributes only", func() {
			actualContent := "a link to link:https://foo.bar[foo=bar]"
			expectedResult := types.InlineElements{
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
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("invalid relative link to doc", func() {
			actualContent := "a link to link:foo.adoc"
			expectedResult := types.InlineElements{
				types.StringElement{
					Content: "a link to link:foo.adoc",
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("relative link with quoted text", func() {
			actualContent := "link:/[_a_ *b* `c`]"
			expectedResult := types.InlineLink{
				Location: types.Location{
					types.StringElement{
						Content: "/",
					},
				},
				Attributes: types.ElementAttributes{
					types.AttrInlineLinkText: types.InlineElements{
						types.QuotedText{
							Kind: types.Italic,
							Elements: types.InlineElements{
								types.StringElement{
									Content: "a",
								},
							},
						},
						types.StringElement{
							Content: " ",
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
							Content: " ",
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
			}
			// verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("RelativeLink"), parser.Debug(true))
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("RelativeLink"))
		})

	})

})
