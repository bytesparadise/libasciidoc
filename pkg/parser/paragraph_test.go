package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("paragraphs - preflight", func() {

	Context("paragraphs", func() {

		It("paragraph with 1 word", func() {
			source := "hello"
			expected := &types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						&types.StringElement{Content: "hello"},
					},
				},
			}
			verifyDocumentBlock(expected, source)
		})

		It("paragraph with few words and ending with spaces", func() {
			source := "a paragraph with some content  "
			expected := &types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						&types.StringElement{Content: "a paragraph with some content  "},
					},
				},
			}
			verifyDocumentBlock(expected, source)
		})

		It("paragraph with bold content and spaces", func() {
			source := "a paragraph with *some bold content*  "
			expected := &types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						&types.StringElement{Content: "a paragraph with "},
						&types.QuotedText{
							Kind: types.Bold,
							Elements: types.InlineElements{
								&types.StringElement{Content: "some bold content"},
							},
						},
						&types.StringElement{Content: "  "},
					},
				},
			}
			verifyDocumentBlock(expected, source)
		})

		It("paragraph with non-alphnum character before bold text", func() {
			source := "+*some bold content*"
			expected := &types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						&types.StringElement{Content: "+"},
						&types.QuotedText{
							Kind: types.Bold,
							Elements: types.InlineElements{
								&types.StringElement{Content: "some bold content"},
							},
						},
					},
				},
			}
			verifyDocumentBlock(expected, source)
		})

		It("paragraph with id and title", func() {
			source := `[#foo]
.a title
a paragraph`
			expected := &types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrID:       "foo",
					types.AttrCustomID: true,
					types.AttrTitle:    "a title",
				},
				Lines: []types.InlineElements{
					{
						&types.StringElement{Content: "a paragraph"},
					},
				},
			}
			verifyDocumentBlock(expected, source)
		})

		It("paragraph with words and dots on same line", func() {
			source := `foo. bar.`
			expected := &types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						&types.StringElement{Content: "foo. bar."},
					},
				},
			}

			verifyDocumentBlock(expected, source)
		})
		It("paragraph with words and dots on two lines", func() {
			source := `foo. 
bar.`
			expected := &types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						&types.StringElement{Content: "foo. "},
					},
					{
						&types.StringElement{Content: "bar."},
					},
				},
			}
			verifyDocumentBlock(expected, source)
		})
	})

	Context("paragraphs with line break", func() {

		It("with explicit line break", func() {
			source := `foo +
bar
baz`
			expected := &types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						&types.StringElement{Content: "foo"},
						&types.LineBreak{},
					},
					{
						&types.StringElement{Content: "bar"},
					},
					{
						&types.StringElement{Content: "baz"},
					},
				},
			}
			verifyDocumentBlock(expected, source)
		})

		It("with paragraph attribute", func() {
			source := `[%hardbreaks]
foo
bar
baz`
			expected := &types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrHardBreaks: nil,
				},
				Lines: []types.InlineElements{
					{
						&types.StringElement{Content: "foo"},
					},
					{
						&types.StringElement{Content: "bar"},
					},
					{
						&types.StringElement{Content: "baz"},
					},
				},
			}
			verifyDocumentBlock(expected, source)
		})

	})

	Context("admonition paragraphs", func() {

		It("note admonition paragraph", func() {
			source := `NOTE: this is a note.`
			expected := &types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrAdmonitionKind: types.Note,
				},
				Lines: []types.InlineElements{
					{
						&types.StringElement{
							Content: "this is a note.",
						},
					},
				},
			}
			verifyDocumentBlock(expected, source)
		})

		It("warning admonition paragraph", func() {
			source := `WARNING: this is a multiline
warning!`
			expected := &types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrAdmonitionKind: types.Warning,
				},
				Lines: []types.InlineElements{
					{
						&types.StringElement{
							Content: "this is a multiline",
						},
					},
					{
						&types.StringElement{
							Content: "warning!",
						},
					},
				},
			}
			verifyDocumentBlock(expected, source)
		})

		It("admonition note paragraph with id and title", func() {
			source := `[[foo]]
.bar
NOTE: this is a note.`
			expected := &types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrAdmonitionKind: types.Note,
					types.AttrID:             "foo",
					types.AttrCustomID:       true,
					types.AttrTitle:          "bar",
				},
				Lines: []types.InlineElements{
					{
						&types.StringElement{
							Content: "this is a note.",
						},
					},
				},
			}
			verifyDocumentBlock(expected, source)
		})

		It("caution admonition paragraph with single line", func() {
			source := `[CAUTION]
this is a caution!`
			expected := &types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrAdmonitionKind: types.Caution,
				},
				Lines: []types.InlineElements{
					{
						&types.StringElement{
							Content: "this is a caution!",
						},
					},
				},
			}
			verifyDocumentBlock(expected, source)
		})

		It("multiline caution admonition paragraph with title and id", func() {
			source := `[[foo]]
[CAUTION] 
.bar
this is a 
*caution*!`
			expected := &types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrAdmonitionKind: types.Caution,
					types.AttrID:             "foo",
					types.AttrCustomID:       true,
					types.AttrTitle:          "bar",
				},
				Lines: []types.InlineElements{
					{
						&types.StringElement{
							Content: "this is a ",
						},
					},
					{
						&types.QuotedText{
							Kind: types.Bold,
							Elements: types.InlineElements{
								&types.StringElement{
									Content: "caution",
								},
							},
						},
						&types.StringElement{
							Content: "!",
						},
					},
				},
			}
			verifyDocumentBlock(expected, source)
		})

		It("multiple admonition paragraphs", func() {
			source := `[NOTE]
No space after the [NOTE]!

[CAUTION]
And no space after [CAUTION] either.`
			expected := &types.PreflightDocument{
				Blocks: []interface{}{
					&types.Paragraph{
						Attributes: types.ElementAttributes{
							types.AttrAdmonitionKind: types.Note,
						},
						Lines: []types.InlineElements{
							{
								&types.StringElement{
									Content: "No space after the [NOTE]!",
								},
							},
						},
					},
					&types.BlankLine{},
					&types.Paragraph{
						Attributes: types.ElementAttributes{
							types.AttrAdmonitionKind: types.Caution,
						},
						Lines: []types.InlineElements{
							{
								&types.StringElement{
									Content: "And no space after [CAUTION] either.",
								},
							},
						},
					},
				},
			}
			verifyPreflight(expected, source)
		})
	})

	Context("verse paragraphs", func() {

		It("paragraph as a verse with author and title", func() {
			source := `[verse, john doe, verse title]
I am a verse paragraph.`
			expected := &types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Verse,
					types.AttrQuoteAuthor: "john doe",
					types.AttrQuoteTitle:  "verse title",
				},
				Lines: []types.InlineElements{
					{
						&types.StringElement{
							Content: "I am a verse paragraph.",
						},
					},
				},
			}
			verifyDocumentBlock(expected, source)
		})

		It("paragraph as a verse with author, title and other attributes", func() {
			source := `[[universal]]
[verse, john doe, verse title]
.universe
I am a verse paragraph.`
			expected := &types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Verse,
					types.AttrQuoteAuthor: "john doe",
					types.AttrQuoteTitle:  "verse title",
					types.AttrID:          "universal",
					types.AttrCustomID:    true,
					types.AttrTitle:       "universe",
				},
				Lines: []types.InlineElements{
					{
						&types.StringElement{
							Content: "I am a verse paragraph.",
						},
					},
				},
			}
			verifyDocumentBlock(expected, source)
		})

		It("paragraph as a verse with empty title", func() {
			source := `[verse, john doe, ]
I am a verse paragraph.`
			expected := &types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Verse,
					types.AttrQuoteAuthor: "john doe",
					types.AttrQuoteTitle:  "",
				},
				Lines: []types.InlineElements{
					{
						&types.StringElement{
							Content: "I am a verse paragraph.",
						},
					},
				},
			}
			verifyDocumentBlock(expected, source)
		})

		It("paragraph as a verse without title", func() {
			source := `[verse, john doe ]
I am a verse paragraph.`
			expected := &types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Verse,
					types.AttrQuoteAuthor: "john doe",
					types.AttrQuoteTitle:  "",
				},
				Lines: []types.InlineElements{
					{
						&types.StringElement{
							Content: "I am a verse paragraph.",
						},
					},
				},
			}
			verifyDocumentBlock(expected, source)
		})

		It("paragraph as a verse with empty author", func() {
			source := `[verse,  ]
I am a verse paragraph.`
			expected := &types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Verse,
					types.AttrQuoteAuthor: "",
					types.AttrQuoteTitle:  "",
				},
				Lines: []types.InlineElements{
					{
						&types.StringElement{
							Content: "I am a verse paragraph.",
						},
					},
				},
			}
			verifyDocumentBlock(expected, source)
		})

		It("paragraph as a verse without author", func() {
			source := `[verse]
I am a verse paragraph.`
			expected := &types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Verse,
					types.AttrQuoteAuthor: "",
					types.AttrQuoteTitle:  "",
				},
				Lines: []types.InlineElements{
					{
						&types.StringElement{
							Content: "I am a verse paragraph.",
						},
					},
				},
			}
			verifyDocumentBlock(expected, source)
		})

		It("image block as a verse", func() {
			source := `[verse, john doe, verse title]
image::foo.png[]`
			expected := &types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Verse,
					types.AttrQuoteAuthor: "john doe",
					types.AttrQuoteTitle:  "verse title",
				},
				Lines: []types.InlineElements{
					{
						&types.StringElement{
							Content: "image::foo.png[]",
						},
					},
				},
			}
			verifyDocumentBlock(expected, source)
		})
	})

	Context("quote paragraphs", func() {

		It("paragraph as a quote with author and title", func() {
			source := `[quote, john doe, quote title]
I am a quote paragraph.`
			expected := &types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Quote,
					types.AttrQuoteAuthor: "john doe",
					types.AttrQuoteTitle:  "quote title",
				},
				Lines: []types.InlineElements{
					{
						&types.StringElement{
							Content: "I am a quote paragraph.",
						},
					},
				},
			}
			verifyDocumentBlock(expected, source)
		})

		It("paragraph as a quote with author, title and other attributes", func() {
			source := `[[universal]]
[quote, john doe, quote title]
.universe
I am a quote paragraph.`
			expected := &types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Quote,
					types.AttrQuoteAuthor: "john doe",
					types.AttrQuoteTitle:  "quote title",
					types.AttrID:          "universal",
					types.AttrCustomID:    true,
					types.AttrTitle:       "universe",
				},
				Lines: []types.InlineElements{
					{
						&types.StringElement{
							Content: "I am a quote paragraph.",
						},
					},
				},
			}
			verifyDocumentBlock(expected, source)
		})

		It("paragraph as a quote with empty title", func() {
			source := `[quote, john doe, ]
I am a quote paragraph.`
			expected := &types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Quote,
					types.AttrQuoteAuthor: "john doe",
					types.AttrQuoteTitle:  "",
				},
				Lines: []types.InlineElements{
					{
						&types.StringElement{
							Content: "I am a quote paragraph.",
						},
					},
				},
			}
			verifyDocumentBlock(expected, source)
		})

		It("paragraph as a quote without title", func() {
			source := `[quote, john doe ]
I am a quote paragraph.`
			expected := &types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Quote,
					types.AttrQuoteAuthor: "john doe",
					types.AttrQuoteTitle:  "",
				},
				Lines: []types.InlineElements{
					{
						&types.StringElement{
							Content: "I am a quote paragraph.",
						},
					},
				},
			}
			verifyDocumentBlock(expected, source)
		})

		It("paragraph as a quote with empty author", func() {
			source := `[quote,  ]
I am a quote paragraph.`
			expected := &types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Quote,
					types.AttrQuoteAuthor: "",
					types.AttrQuoteTitle:  "",
				},
				Lines: []types.InlineElements{
					{
						&types.StringElement{
							Content: "I am a quote paragraph.",
						},
					},
				},
			}
			verifyDocumentBlock(expected, source)
		})

		It("paragraph as a quote without author", func() {
			source := `[quote]
I am a quote paragraph.`
			expected := &types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Quote,
					types.AttrQuoteAuthor: "",
					types.AttrQuoteTitle:  "",
				},
				Lines: []types.InlineElements{
					{
						&types.StringElement{
							Content: "I am a quote paragraph.",
						},
					},
				},
			}
			verifyDocumentBlock(expected, source)
		})

		It("inline image within a quote", func() {
			source := `[quote, john doe, quote title]
a foo image:foo.png[]`
			expected := &types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Quote,
					types.AttrQuoteAuthor: "john doe",
					types.AttrQuoteTitle:  "quote title",
				},
				Lines: []types.InlineElements{
					{
						&types.StringElement{
							Content: "a foo ",
						},
						&types.InlineImage{
							Attributes: types.ElementAttributes{
								types.AttrImageAlt: "foo",
							},
							Path: "foo.png",
						},
					},
				},
			}
			verifyDocumentBlock(expected, source)
		})

		It("image block is NOT a quote", func() {
			source := `[quote, john doe, quote title]
image::foo.png[]`
			expected := &types.ImageBlock{
				Attributes: types.ElementAttributes{
					types.AttrImageAlt: "foo",
					// quote attributes
					types.AttrKind:        types.Quote,
					types.AttrQuoteAuthor: "john doe",
					types.AttrQuoteTitle:  "quote title",
				},
				Path: "foo.png",
			}
			verifyDocumentBlock(expected, source) //, parser.Debug(true))
		})
	})
})
