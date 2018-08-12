package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
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
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("paragraph with few words and ending with spaces", func() {
			actualContent := "a paragraph with some content  "
			expectedResult := types.Paragraph{
				Attributes: map[string]interface{}{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "a paragraph with some content  "},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("paragraph with bold content and spaces", func() {
			actualContent := "a paragraph with *some bold content*  "
			expectedResult := types.Paragraph{
				Attributes: map[string]interface{}{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "a paragraph with "},
						types.QuotedText{
							Kind: types.Bold,
							Elements: types.InlineElements{
								types.StringElement{Content: "some bold content"},
							},
						},
						types.StringElement{Content: "  "},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
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
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
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
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
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
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
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
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
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
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
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
							Elements: types.InlineElements{
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
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
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
					types.BlankLine{},
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

	Context("verse paragraphs", func() {

		It("regular paragraph as a verse with author and title", func() {
			actualContent := `[verse, john doe, verse title]
I am a verse paragraph.`
			expectedResult := types.Paragraph{
				Attributes: map[string]interface{}{
					types.AttrBlockKind:   types.Verse,
					types.AttrQuoteAuthor: "john doe",
					types.AttrQuoteTitle:  "verse title",
				},
				Lines: []types.InlineElements{
					{
						types.StringElement{
							Content: "I am a verse paragraph.",
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("regular paragraph as a verse with author, title and other attributes", func() {
			actualContent := `[[universe]]
[verse, john doe, verse title]
.universe
I am a verse paragraph.`
			expectedResult := types.Paragraph{
				Attributes: map[string]interface{}{
					types.AttrBlockKind:   types.Verse,
					types.AttrQuoteAuthor: "john doe",
					types.AttrQuoteTitle:  "verse title",
					types.AttrID:          "universe",
					types.AttrTitle:       "universe",
				},
				Lines: []types.InlineElements{
					{
						types.StringElement{
							Content: "I am a verse paragraph.",
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("regular paragraph as a verse with empty title", func() {
			actualContent := `[verse, john doe, ]
I am a verse paragraph.`
			expectedResult := types.Paragraph{
				Attributes: map[string]interface{}{
					types.AttrBlockKind:   types.Verse,
					types.AttrQuoteAuthor: "john doe",
					types.AttrQuoteTitle:  "",
				},
				Lines: []types.InlineElements{
					{
						types.StringElement{
							Content: "I am a verse paragraph.",
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("regular paragraph as a verse without title", func() {
			actualContent := `[verse, john doe ]
I am a verse paragraph.`
			expectedResult := types.Paragraph{
				Attributes: map[string]interface{}{
					types.AttrBlockKind:   types.Verse,
					types.AttrQuoteAuthor: "john doe",
					types.AttrQuoteTitle:  "",
				},
				Lines: []types.InlineElements{
					{
						types.StringElement{
							Content: "I am a verse paragraph.",
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("regular paragraph as a verse with empty author", func() {
			actualContent := `[verse,  ]
I am a verse paragraph.`
			expectedResult := types.Paragraph{
				Attributes: map[string]interface{}{
					types.AttrBlockKind:   types.Verse,
					types.AttrQuoteAuthor: "",
					types.AttrQuoteTitle:  "",
				},
				Lines: []types.InlineElements{
					{
						types.StringElement{
							Content: "I am a verse paragraph.",
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("regular paragraph as a verse without author", func() {
			actualContent := `[verse]
I am a verse paragraph.`
			expectedResult := types.Paragraph{
				Attributes: map[string]interface{}{
					types.AttrBlockKind:   types.Verse,
					types.AttrQuoteAuthor: "",
					types.AttrQuoteTitle:  "",
				},
				Lines: []types.InlineElements{
					{
						types.StringElement{
							Content: "I am a verse paragraph.",
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Paragraph"))
		})

		It("image block as a verse", func() {
			actualContent := `[verse, john doe, verse title]
image::foo.png[]`
			expectedResult := types.Paragraph{
				Attributes: map[string]interface{}{
					types.AttrBlockKind:   types.Verse,
					types.AttrQuoteAuthor: "john doe",
					types.AttrQuoteTitle:  "verse title",
				},
				Lines: []types.InlineElements{
					{
						types.StringElement{
							Content: "image::foo.png[]",
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})
	})
})
