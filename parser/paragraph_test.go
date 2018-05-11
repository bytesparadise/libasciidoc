package parser_test

import (
	"github.com/bytesparadise/libasciidoc/parser"
	"github.com/bytesparadise/libasciidoc/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("paragraphs", func() {

	Context("regular paragraphs", func() {

		It("paragraph with 1 word", func() {
			actualContent := "hello"
			expectedResult := types.Paragraph{
				Attributes: map[string]interface{}{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "hello"},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("BlockElement"))
		})

		It("paragraph with few words", func() {
			actualContent := "a paragraph with some content"
			expectedResult := types.Paragraph{
				Attributes: map[string]interface{}{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "a paragraph with some content"},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("BlockElement"))
		})

		It("paragraph with bold content", func() {
			actualContent := "a paragraph with *some bold content*"
			expectedResult := types.Paragraph{
				Attributes: map[string]interface{}{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "a paragraph with "},
						types.QuotedText{
							Kind: types.Bold,
							Elements: []interface{}{
								types.StringElement{Content: "some bold content"},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("BlockElement"))
		})

		It("paragraph with id and title", func() {
			actualContent := `[#foo]
.a title
a paragraph`
			expectedResult := types.Paragraph{
				Attributes: map[string]interface{}{
					types.AttrID:    "foo",
					types.AttrTitle: "a title",
				},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "a paragraph"},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("BlockElement"))
		})
	})

	Context("admonition paragraphs", func() {
		It("note admonition paragraph", func() {
			actualContent := `NOTE: this is a note.`
			expectedResult := types.Paragraph{
				Attributes: map[string]interface{}{
					types.AttrAdmonitionKind: types.Note,
				},
				Lines: []types.InlineElements{
					{
						types.StringElement{
							Content: "this is a note.",
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("BlockElement"))
		})

		It("warning admonition paragraph", func() {
			actualContent := `WARNING: this is a multiline
warning!`
			expectedResult := types.Paragraph{
				Attributes: map[string]interface{}{
					types.AttrAdmonitionKind: types.Warning,
				},
				Lines: []types.InlineElements{
					{
						types.StringElement{
							Content: "this is a multiline",
						},
					},
					{
						types.StringElement{
							Content: "warning!",
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
			expectedResult := types.Paragraph{
				Attributes: map[string]interface{}{
					types.AttrAdmonitionKind: types.Note,
					types.AttrID:             "foo",
					types.AttrTitle:          "bar",
				},
				Lines: []types.InlineElements{
					{
						types.StringElement{
							Content: "this is a note.",
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("BlockElement"))
		})

		It("caution admonition paragraph with single line", func() {
			actualContent := `[CAUTION]
this is a caution!`
			expectedResult := types.Paragraph{
				Attributes: map[string]interface{}{
					types.AttrAdmonitionKind: types.Caution,
				},
				Lines: []types.InlineElements{
					{
						types.StringElement{
							Content: "this is a caution!",
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
			expectedResult := types.Paragraph{
				Attributes: map[string]interface{}{
					types.AttrAdmonitionKind: types.Caution,
					types.AttrID:             "foo",
					types.AttrTitle:          "bar",
				},
				Lines: []types.InlineElements{
					{
						types.StringElement{
							Content: "this is a ",
						},
					},
					{
						types.QuotedText{
							Kind: types.Bold,
							Elements: []interface{}{
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
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("BlockElement"))
		})

		It("multiple admonition paragraphs", func() {
			actualContent := `[NOTE]
No space after the [NOTE]!

[CAUTION]
And no space after [CAUTION] either.`
			expectedResult := types.Document{
				Attributes:        map[string]interface{}{},
				ElementReferences: map[string]interface{}{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: map[string]interface{}{
							types.AttrAdmonitionKind: types.Note,
						},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "No space after the [NOTE]!",
								},
							},
						},
					},
					types.Paragraph{
						Attributes: map[string]interface{}{
							types.AttrAdmonitionKind: types.Caution,
						},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "And no space after [CAUTION] either.",
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Document"))
		})
	})
})
