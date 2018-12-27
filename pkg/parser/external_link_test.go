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
							URL: "https://foo.bar",
							Attributes: types.ElementAttributes{
								"text": "",
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("external link with empty text", func() {
			actualContent := "a link to https://foo.bar[]"
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "a link to "},
						types.InlineLink{
							URL: "https://foo.bar",
							Attributes: types.ElementAttributes{
								"text": "",
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("external link with text", func() {
			actualContent := "a link to mailto:foo@bar[the foo@bar email]"
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "a link to "},
						types.InlineLink{
							URL: "mailto:foo@bar",
							Attributes: types.ElementAttributes{
								"text": "the foo@bar email",
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("external link with text and extra attributes", func() {
			actualContent := "a link to mailto:foo@bar[the foo@bar email, foo=bar]"
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "a link to "},
						types.InlineLink{
							URL: "mailto:foo@bar",
							Attributes: types.ElementAttributes{
								"text": "the foo@bar email",
								"foo":  "bar",
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("external link with extra attributes only", func() {
			actualContent := "a link to mailto:foo@bar[foo=bar]"
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "a link to "},
						types.InlineLink{
							URL: "mailto:foo@bar",
							Attributes: types.ElementAttributes{
								"text": "",
								"foo":  "bar",
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})
	})

	Context("relative links", func() {

		It("relative link to doc without text", func() {
			actualContent := "a link to link:foo.adoc[]"
			expectedResult := types.InlineElements{
				types.StringElement{Content: "a link to "},
				types.InlineLink{
					URL: "foo.adoc",
					Attributes: types.ElementAttributes{
						"text": "",
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("relative link to doc with text", func() {
			actualContent := "a link to link:foo.adoc[foo doc]"
			expectedResult := types.InlineElements{
				types.StringElement{Content: "a link to "},
				types.InlineLink{
					URL: "foo.adoc",
					Attributes: types.ElementAttributes{
						"text": "foo doc",
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("relative link to external URL with text", func() {
			actualContent := "a link to link:https://foo.bar[foo doc]"
			expectedResult := types.InlineElements{
				types.StringElement{Content: "a link to "},
				types.InlineLink{
					URL: "https://foo.bar",
					Attributes: types.ElementAttributes{
						"text": "foo doc",
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("relative link to external URL with text and extra attributes", func() {
			actualContent := "a link to link:https://foo.bar[foo doc, foo=bar]"
			expectedResult := types.InlineElements{
				types.StringElement{Content: "a link to "},
				types.InlineLink{
					URL: "https://foo.bar",
					Attributes: types.ElementAttributes{
						"text": "foo doc",
						"foo":  "bar",
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("relative link to external URL with extra attributes only", func() {
			actualContent := "a link to link:https://foo.bar[foo=bar]"
			expectedResult := types.InlineElements{
				types.StringElement{Content: "a link to "},
				types.InlineLink{
					URL: "https://foo.bar",
					Attributes: types.ElementAttributes{
						"text": "",
						"foo":  "bar",
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("invalid relative link to doc", func() {
			actualContent := "a link to link:foo.adoc"
			expectedResult := types.InlineElements{
				types.StringElement{
					Content: "a link to link:foo.adoc",
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})
	})

})
