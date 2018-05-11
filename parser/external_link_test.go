package parser_test

import (
	"github.com/bytesparadise/libasciidoc/parser"
	"github.com/bytesparadise/libasciidoc/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("links", func() {

	Context("external links", func() {

		It("external link without label", func() {
			actualContent := "a link to https://foo.bar"
			expectedResult := types.InlineElements{
				types.StringElement{Content: "a link to "},
				types.Link{
					URL: "https://foo.bar",
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("external link with empty label", func() {
			actualContent := "a link to https://foo.bar[]"
			expectedResult := types.InlineElements{
				types.StringElement{Content: "a link to "},
				types.Link{
					URL:  "https://foo.bar",
					Text: "",
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("external link with text", func() {
			actualContent := "a link to mailto:foo@bar[the foo@bar email]"
			expectedResult := types.InlineElements{
				types.StringElement{Content: "a link to "},
				types.Link{
					URL:  "mailto:foo@bar",
					Text: "the foo@bar email",
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})
	})

	Context("relative links", func() {

		It("relative link to doc without text", func() {
			actualContent := "a link to link:foo.adoc[]"
			expectedResult := types.InlineElements{
				types.StringElement{Content: "a link to "},
				types.Link{
					URL: "foo.adoc",
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("relative link to doc with text", func() {
			actualContent := "a link to link:foo.adoc[foo doc]"
			expectedResult := types.InlineElements{
				types.StringElement{Content: "a link to "},
				types.Link{
					URL:  "foo.adoc",
					Text: "foo doc",
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("relative link to external URL with text", func() {
			actualContent := "a link to link:https://foo.bar[foo doc]"
			expectedResult := types.InlineElements{
				types.StringElement{Content: "a link to "},
				types.Link{
					URL:  "https://foo.bar",
					Text: "foo doc",
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("invalid relative link to doc", func() {
			actualContent := "a link to link:foo.adoc"
			expectedResult := types.InlineElements{
				types.StringElement{Content: "a link to link:foo.adoc"},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})
	})

})
