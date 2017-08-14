package parser_test

import (
	"github.com/bytesparadise/libasciidoc/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Parsing External Links", func() {

	It("external link", func() {
		actualContent := "a link to https://foo.bar"
		expectedDocument := &types.Document{
			Elements: []types.DocElement{
				&types.Paragraph{
					Lines: []*types.InlineContent{
						&types.InlineContent{
							Elements: []types.DocElement{
								&types.StringElement{Content: "a link to "},
								&types.ExternalLink{
									URL: "https://foo.bar",
								},
							},
						},
					},
				},
			},
		}
		verify(GinkgoT(), expectedDocument, actualContent)
	})

	It("external link with empty text", func() {
		actualContent := "a link to https://foo.bar[]"
		expectedDocument := &types.Document{
			Elements: []types.DocElement{
				&types.Paragraph{
					Lines: []*types.InlineContent{
						&types.InlineContent{
							Elements: []types.DocElement{
								&types.StringElement{Content: "a link to "},
								&types.ExternalLink{
									URL:  "https://foo.bar",
									Text: "",
								},
							},
						},
					},
				},
			},
		}
		verify(GinkgoT(), expectedDocument, actualContent)
	})

	It("external link with text", func() {
		actualContent := "a link to mailto:foo@bar[the foo@bar email]"
		expectedDocument := &types.Document{
			Elements: []types.DocElement{
				&types.Paragraph{
					Lines: []*types.InlineContent{
						&types.InlineContent{
							Elements: []types.DocElement{
								&types.StringElement{Content: "a link to "},
								&types.ExternalLink{
									URL:  "mailto:foo@bar",
									Text: "the foo@bar email",
								},
							},
						},
					},
				},
			},
		}
		verify(GinkgoT(), expectedDocument, actualContent)
	})
})
