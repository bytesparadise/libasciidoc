package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("paragraphs - draft", func() {

	Context("paragraphs", func() {

		It("paragraph with 1 word", func() {
			source := "hello"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: [][]interface{}{
					{
						types.StringElement{Content: "hello"},
					},
				},
			}
			Expect(source).To(BecomeDocumentBlock(expected))
		})

		It("paragraph with few words and ending with spaces", func() {
			source := "a paragraph with some content  "
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: [][]interface{}{
					{
						types.StringElement{Content: "a paragraph with some content  "},
					},
				},
			}
			Expect(source).To(BecomeDocumentBlock(expected))
		})

		It("paragraph with bold content and spaces", func() {
			source := "a paragraph with *some bold content*  "
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: [][]interface{}{
					{
						types.StringElement{Content: "a paragraph with "},
						types.QuotedText{
							Kind: types.Bold,
							Elements: []interface{}{
								types.StringElement{Content: "some bold content"},
							},
						},
						types.StringElement{Content: "  "},
					},
				},
			}
			Expect(source).To(BecomeDocumentBlock(expected))
		})

		It("paragraph with non-alphnum character before bold text", func() {
			source := "+*some bold content*"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: [][]interface{}{
					{
						types.StringElement{Content: "+"},
						types.QuotedText{
							Kind: types.Bold,
							Elements: []interface{}{
								types.StringElement{Content: "some bold content"},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDocumentBlock(expected))
		})

		It("paragraph with id and title", func() {
			source := `[#foo]
.a title
a paragraph`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrID:       "foo",
					types.AttrCustomID: true,
					types.AttrTitle:    "a title",
				},
				Lines: [][]interface{}{
					{
						types.StringElement{Content: "a paragraph"},
					},
				},
			}
			Expect(source).To(BecomeDocumentBlock(expected))
		})

		It("paragraph with words and dots on same line", func() {
			source := `foo. bar.`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: [][]interface{}{
					{
						types.StringElement{Content: "foo. bar."},
					},
				},
			}

			Expect(source).To(BecomeDocumentBlock(expected))
		})
		It("paragraph with words and dots on two lines", func() {
			source := `foo. 
bar.`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: [][]interface{}{
					{
						types.StringElement{Content: "foo. "},
					},
					{
						types.StringElement{Content: "bar."},
					},
				},
			}
			Expect(source).To(BecomeDocumentBlock(expected))
		})

		It("paragraph with custom id prefix and title", func() {
			source := `:idprefix: bar_
			
.a title
a paragraph`
			expected := types.Document{
				Attributes: types.DocumentAttributes{
					types.AttrIDPrefix: "bar_",
				},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{
							types.AttrTitle: "a title", // there is no default ID. Only custom IDs
						},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "a paragraph"},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDocument(expected))
		})
	})

	Context("paragraphs with line break", func() {

		It("with explicit line break", func() {
			source := `foo +
bar
baz`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: [][]interface{}{
					{
						types.StringElement{Content: "foo"},
						types.LineBreak{},
					},
					{
						types.StringElement{Content: "bar"},
					},
					{
						types.StringElement{Content: "baz"},
					},
				},
			}
			Expect(source).To(BecomeDocumentBlock(expected))
		})

		It("with paragraph attribute", func() {
			source := `[%hardbreaks]
foo
bar
baz`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrHardBreaks: nil,
				},
				Lines: [][]interface{}{
					{
						types.StringElement{Content: "foo"},
					},
					{
						types.StringElement{Content: "bar"},
					},
					{
						types.StringElement{Content: "baz"},
					},
				},
			}
			Expect(source).To(BecomeDocumentBlock(expected))
		})

	})

	Context("admonition paragraphs", func() {

		It("note admonition paragraph", func() {
			source := `NOTE: this is a note.`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrAdmonitionKind: types.Note,
				},
				Lines: [][]interface{}{
					{
						types.StringElement{
							Content: "this is a note.",
						},
					},
				},
			}
			Expect(source).To(BecomeDocumentBlock(expected))
		})

		It("warning admonition paragraph", func() {
			source := `WARNING: this is a multiline
warning!`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrAdmonitionKind: types.Warning,
				},
				Lines: [][]interface{}{
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
			Expect(source).To(BecomeDocumentBlock(expected))
		})

		It("admonition note paragraph with id and title", func() {
			source := `[[foo]]
.bar
NOTE: this is a note.`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrAdmonitionKind: types.Note,
					types.AttrID:             "foo",
					types.AttrCustomID:       true,
					types.AttrTitle:          "bar",
				},
				Lines: [][]interface{}{
					{
						types.StringElement{
							Content: "this is a note.",
						},
					},
				},
			}
			Expect(source).To(BecomeDocumentBlock(expected))
		})

		It("caution admonition paragraph with single line", func() {
			source := `[CAUTION]
this is a caution!`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrAdmonitionKind: types.Caution,
				},
				Lines: [][]interface{}{
					{
						types.StringElement{
							Content: "this is a caution!",
						},
					},
				},
			}
			Expect(source).To(BecomeDocumentBlock(expected))
		})

		It("multiline caution admonition paragraph with title and id", func() {
			source := `[[foo]]
[CAUTION] 
.bar
this is a 
*caution*!`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrAdmonitionKind: types.Caution,
					types.AttrID:             "foo",
					types.AttrCustomID:       true,
					types.AttrTitle:          "bar",
				},
				Lines: [][]interface{}{
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
			Expect(source).To(BecomeDocumentBlock(expected))
		})

		It("multiple admonition paragraphs", func() {
			source := `[NOTE]
No space after the [NOTE]!

[CAUTION]
And no space after [CAUTION] either.`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{
							types.AttrAdmonitionKind: types.Note,
						},
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "No space after the [NOTE]!",
								},
							},
						},
					},
					types.BlankLine{},
					types.Paragraph{
						Attributes: types.ElementAttributes{
							types.AttrAdmonitionKind: types.Caution,
						},
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "And no space after [CAUTION] either.",
								},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})
	})

	Context("verse paragraphs", func() {

		It("paragraph as a verse with author and title", func() {
			source := `[verse, john doe, verse title]
I am a verse paragraph.`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Verse,
					types.AttrQuoteAuthor: "john doe",
					types.AttrQuoteTitle:  "verse title",
				},
				Lines: [][]interface{}{
					{
						types.StringElement{
							Content: "I am a verse paragraph.",
						},
					},
				},
			}
			Expect(source).To(BecomeDocumentBlock(expected))
		})

		It("paragraph as a verse with author, title and other attributes", func() {
			source := `[[universal]]
[verse, john doe, verse title]
.universe
I am a verse paragraph.`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Verse,
					types.AttrQuoteAuthor: "john doe",
					types.AttrQuoteTitle:  "verse title",
					types.AttrID:          "universal",
					types.AttrCustomID:    true,
					types.AttrTitle:       "universe",
				},
				Lines: [][]interface{}{
					{
						types.StringElement{
							Content: "I am a verse paragraph.",
						},
					},
				},
			}
			Expect(source).To(BecomeDocumentBlock(expected))
		})

		It("paragraph as a verse with empty title", func() {
			source := `[verse, john doe, ]
I am a verse paragraph.`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Verse,
					types.AttrQuoteAuthor: "john doe",
				},
				Lines: [][]interface{}{
					{
						types.StringElement{
							Content: "I am a verse paragraph.",
						},
					},
				},
			}
			Expect(source).To(BecomeDocumentBlock(expected))
		})

		It("paragraph as a verse without title", func() {
			source := `[verse, john doe ]
I am a verse paragraph.`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Verse,
					types.AttrQuoteAuthor: "john doe",
				},
				Lines: [][]interface{}{
					{
						types.StringElement{
							Content: "I am a verse paragraph.",
						},
					},
				},
			}
			Expect(source).To(BecomeDocumentBlock(expected))
		})

		It("paragraph as a verse with empty author", func() {
			source := `[verse,  ]
I am a verse paragraph.`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind: types.Verse,
				},
				Lines: [][]interface{}{
					{
						types.StringElement{
							Content: "I am a verse paragraph.",
						},
					},
				},
			}
			Expect(source).To(BecomeDocumentBlock(expected))
		})

		It("paragraph as a verse without author", func() {
			source := `[verse]
I am a verse paragraph.`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind: types.Verse,
				},
				Lines: [][]interface{}{
					{
						types.StringElement{
							Content: "I am a verse paragraph.",
						},
					},
				},
			}
			Expect(source).To(BecomeDocumentBlock(expected))
		})

		It("image block as a verse", func() {
			source := `[verse, john doe, verse title]
image::foo.png[]`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Verse,
					types.AttrQuoteAuthor: "john doe",
					types.AttrQuoteTitle:  "verse title",
				},
				Lines: [][]interface{}{
					{
						types.StringElement{
							Content: "image::foo.png[]",
						},
					},
				},
			}
			Expect(source).To(BecomeDocumentBlock(expected))
		})
	})

	Context("quote paragraphs", func() {

		It("paragraph as a quote with author and title", func() {
			source := `[quote, john doe, quote title]
I am a quote paragraph.`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Quote,
					types.AttrQuoteAuthor: "john doe",
					types.AttrQuoteTitle:  "quote title",
				},
				Lines: [][]interface{}{
					{
						types.StringElement{
							Content: "I am a quote paragraph.",
						},
					},
				},
			}
			Expect(source).To(BecomeDocumentBlock(expected))
		})

		It("paragraph as a quote with author, title and other attributes", func() {
			source := `[[universal]]
[quote, john doe, quote title]
.universe
I am a quote paragraph.`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Quote,
					types.AttrQuoteAuthor: "john doe",
					types.AttrQuoteTitle:  "quote title",
					types.AttrID:          "universal",
					types.AttrCustomID:    true,
					types.AttrTitle:       "universe",
				},
				Lines: [][]interface{}{
					{
						types.StringElement{
							Content: "I am a quote paragraph.",
						},
					},
				},
			}
			Expect(source).To(BecomeDocumentBlock(expected))
		})

		It("paragraph as a quote with empty title", func() {
			source := `[quote, john doe, ]
I am a quote paragraph.`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Quote,
					types.AttrQuoteAuthor: "john doe",
				},
				Lines: [][]interface{}{
					{
						types.StringElement{
							Content: "I am a quote paragraph.",
						},
					},
				},
			}
			Expect(source).To(BecomeDocumentBlock(expected))
		})

		It("paragraph as a quote without title", func() {
			source := `[quote, john doe ]
I am a quote paragraph.`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Quote,
					types.AttrQuoteAuthor: "john doe",
				},
				Lines: [][]interface{}{
					{
						types.StringElement{
							Content: "I am a quote paragraph.",
						},
					},
				},
			}
			Expect(source).To(BecomeDocumentBlock(expected))
		})

		It("paragraph as a quote with empty author", func() {
			source := `[quote,  ]
I am a quote paragraph.`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind: types.Quote,
				},
				Lines: [][]interface{}{
					{
						types.StringElement{
							Content: "I am a quote paragraph.",
						},
					},
				},
			}
			Expect(source).To(BecomeDocumentBlock(expected))
		})

		It("paragraph as a quote without author", func() {
			source := `[quote]
I am a quote paragraph.`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind: types.Quote,
				},
				Lines: [][]interface{}{
					{
						types.StringElement{
							Content: "I am a quote paragraph.",
						},
					},
				},
			}
			Expect(source).To(BecomeDocumentBlock(expected))
		})

		It("inline image within a quote", func() {
			source := `[quote, john doe, quote title]
a foo image:foo.png[]`
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Quote,
					types.AttrQuoteAuthor: "john doe",
					types.AttrQuoteTitle:  "quote title",
				},
				Lines: [][]interface{}{
					{
						types.StringElement{
							Content: "a foo ",
						},
						types.InlineImage{
							Attributes: types.ElementAttributes{},
							Location: types.Location{
								Elements: []interface{}{
									types.StringElement{
										Content: "foo.png",
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDocumentBlock(expected))
		})

		It("image block is NOT a quote", func() {
			source := `[quote, john doe, quote title]
image::foo.png[]`
			expected := types.ImageBlock{
				Attributes: types.ElementAttributes{

					// quote attributes
					types.AttrKind:        types.Quote,
					types.AttrQuoteAuthor: "john doe",
					types.AttrQuoteTitle:  "quote title",
				},
				Location: types.Location{
					Elements: []interface{}{
						types.StringElement{
							Content: "foo.png",
						},
					},
				},
			}
			Expect(source).To(BecomeDocumentBlock(expected)) //, parser.Debug(true))
		})
	})
})

var _ = Describe("paragraphs - final document", func() {

	It("paragraph with predefined attribute", func() {
		source := "hello {plus} world"
		expected := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{Content: "hello &#43; world"},
						},
					},
				},
			},
		}
		Expect(source).To(BecomeDocument(expected))
	})
})
