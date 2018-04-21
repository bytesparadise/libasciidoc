package parser_test

import (
	"github.com/bytesparadise/libasciidoc/parser"
	"github.com/bytesparadise/libasciidoc/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("admonitions", func() {

	Context("admonition paragraphs", func() {
		It("note admonition paragraph", func() {
			actualContent := `NOTE: this is a note.`
			expectedResult := types.Admonition{
				Kind:       types.Note,
				Attributes: map[string]interface{}{},
				Content: types.AdmonitionParagraph{
					Lines: []types.InlineContent{
						{
							Elements: []types.InlineElement{
								types.StringElement{
									Content: "this is a note.",
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("BlockElement"))
		})

		It("warning admonition paragraph", func() {
			actualContent := `WARNING: this is a multiline
warning!`
			expectedResult := types.Admonition{
				Kind:       types.Warning,
				Attributes: map[string]interface{}{},
				Content: types.AdmonitionParagraph{
					Lines: []types.InlineContent{
						{
							Elements: []types.InlineElement{
								types.StringElement{
									Content: "this is a multiline",
								},
							},
						},
						{
							Elements: []types.InlineElement{
								types.StringElement{
									Content: "warning!",
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("BlockElement"))
		})

		It("admonition note paragraph with id and title", func() {
			actualContent := `[[foo]]
.bar
NOTE: this is a note.`
			expectedResult := types.Admonition{
				Attributes: map[string]interface{}{
					types.AttrID:    "foo",
					types.AttrTitle: "bar",
				},
				Kind: types.Note,
				Content: types.AdmonitionParagraph{
					Lines: []types.InlineContent{
						{
							Elements: []types.InlineElement{
								types.StringElement{
									Content: "this is a note.",
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("BlockElement"))
		})
	})

	Context("admonition paragraphs", func() {
		It("caution admonition paragraph", func() {
			actualContent := `[CAUTION] 
this is a caution!`
			expectedResult := types.Admonition{
				Kind:       types.Caution,
				Attributes: map[string]interface{}{},
				Content: types.AdmonitionParagraph{
					Lines: []types.InlineContent{
						{
							Elements: []types.InlineElement{
								types.StringElement{
									Content: "this is a caution!",
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("BlockElement"))
		})

		It("multiline caution admonition paragraph with title and id", func() {
			actualContent := `[[foo]]
[CAUTION] 
.bar
this is a 
*caution*!`
			expectedResult := types.Admonition{
				Attributes: map[string]interface{}{
					types.AttrID:    "foo",
					types.AttrTitle: "bar",
				},
				Kind: types.Caution,
				Content: types.AdmonitionParagraph{
					Lines: []types.InlineContent{
						{
							Elements: []types.InlineElement{
								types.StringElement{
									Content: "this is a ",
								},
							},
						},
						{
							Elements: []types.InlineElement{
								types.QuotedText{
									Kind: types.Bold,
									Elements: []types.InlineElement{
										types.StringElement{
											Content: "caution",
										},
									},
								},
								types.StringElement{
									Content: "!",
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("BlockElement"))
		})
	})

})
