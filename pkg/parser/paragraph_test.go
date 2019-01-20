package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("paragraphs", func() {

	Context("paragraphs", func() {

		It("paragraph with 1 word", func() {
			actualContent := "hello"
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{},
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
				Attributes: types.ElementAttributes{},
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
				Attributes: types.ElementAttributes{},
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
				Attributes: types.ElementAttributes{
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

		It("paragraph with words and dots on same line", func() {
			actualContent := `foo. bar.`
			expectedResult := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "foo. bar."},
							},
						},
					},
				},
			}

			verify(GinkgoT(), expectedResult, actualContent)
		})
		It("paragraph with words and dots on two lines", func() {
			actualContent := `foo. 
bar.`
			expectedResult := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "foo. "},
							},
							{
								types.StringElement{Content: "bar."},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

	Context("paragraphs with line break", func() {

		It("with explicit line break", func() {
			actualContent := `foo +
bar
baz`
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
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
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("with paragraph attribute", func() {
			actualContent := `[%hardbreaks]
foo
bar
baz`
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrHardBreaks: nil,
				},
				Lines: []types.InlineElements{
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
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		// It("paragraph with InlineElementID", func() {
		// 	actualContent := `foo [[id]] bar`
		// 	expectedResult := types.Paragraph{
		// 		Attributes: types.ElementAttributes{},
		// 		Lines: []types.InlineElements{
		// 			{
		// 				types.StringElement{Content: "foo "},
		// 				types.StringElement{Content: " bar"},
		// 			},
		// 		},
		// 	}
		// 	verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		// })

	})

	Context("admonition paragraphs", func() {

		It("note admonition paragraph", func() {
			actualContent := `NOTE: this is a note.`
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{
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
				Attributes: types.ElementAttributes{
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
				Attributes: types.ElementAttributes{
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
				Attributes: types.ElementAttributes{
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
				Attributes: types.ElementAttributes{
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
				Attributes:         map[string]interface{}{},
				ElementReferences:  map[string]interface{}{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{
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
						Attributes: types.ElementAttributes{
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

		It("paragraph as a verse with author and title", func() {
			actualContent := `[verse, john doe, verse title]
I am a verse paragraph.`
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Verse,
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

		It("paragraph as a verse with author, title and other attributes", func() {
			actualContent := `[[universal]]
[verse, john doe, verse title]
.universe
I am a verse paragraph.`
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Verse,
					types.AttrQuoteAuthor: "john doe",
					types.AttrQuoteTitle:  "verse title",
					types.AttrID:          "universal",
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

		It("paragraph as a verse with empty title", func() {
			actualContent := `[verse, john doe, ]
I am a verse paragraph.`
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Verse,
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

		It("paragraph as a verse without title", func() {
			actualContent := `[verse, john doe ]
I am a verse paragraph.`
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Verse,
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

		It("paragraph as a verse with empty author", func() {
			actualContent := `[verse,  ]
I am a verse paragraph.`
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Verse,
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

		It("paragraph as a verse without author", func() {
			actualContent := `[verse]
I am a verse paragraph.`
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Verse,
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

		It("image block as a verse", func() {
			actualContent := `[verse, john doe, verse title]
image::foo.png[]`
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Verse,
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

	Context("quote paragraphs", func() {

		It("paragraph as a quote with author and title", func() {
			actualContent := `[quote, john doe, quote title]
I am a quote paragraph.`
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Quote,
					types.AttrQuoteAuthor: "john doe",
					types.AttrQuoteTitle:  "quote title",
				},
				Lines: []types.InlineElements{
					{
						types.StringElement{
							Content: "I am a quote paragraph.",
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("paragraph as a quote with author, title and other attributes", func() {
			actualContent := `[[universal]]
[quote, john doe, quote title]
.universe
I am a quote paragraph.`
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Quote,
					types.AttrQuoteAuthor: "john doe",
					types.AttrQuoteTitle:  "quote title",
					types.AttrID:          "universal",
					types.AttrTitle:       "universe",
				},
				Lines: []types.InlineElements{
					{
						types.StringElement{
							Content: "I am a quote paragraph.",
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("paragraph as a quote with empty title", func() {
			actualContent := `[quote, john doe, ]
I am a quote paragraph.`
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Quote,
					types.AttrQuoteAuthor: "john doe",
					types.AttrQuoteTitle:  "",
				},
				Lines: []types.InlineElements{
					{
						types.StringElement{
							Content: "I am a quote paragraph.",
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("paragraph as a quote without title", func() {
			actualContent := `[quote, john doe ]
I am a quote paragraph.`
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Quote,
					types.AttrQuoteAuthor: "john doe",
					types.AttrQuoteTitle:  "",
				},
				Lines: []types.InlineElements{
					{
						types.StringElement{
							Content: "I am a quote paragraph.",
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("paragraph as a quote with empty author", func() {
			actualContent := `[quote,  ]
I am a quote paragraph.`
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Quote,
					types.AttrQuoteAuthor: "",
					types.AttrQuoteTitle:  "",
				},
				Lines: []types.InlineElements{
					{
						types.StringElement{
							Content: "I am a quote paragraph.",
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("paragraph as a quote without author", func() {
			actualContent := `[quote]
I am a quote paragraph.`
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Quote,
					types.AttrQuoteAuthor: "",
					types.AttrQuoteTitle:  "",
				},
				Lines: []types.InlineElements{
					{
						types.StringElement{
							Content: "I am a quote paragraph.",
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("inline image within a quote", func() {
			actualContent := `[quote, john doe, quote title]
a foo image:foo.png[]`
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{
					types.AttrKind:        types.Quote,
					types.AttrQuoteAuthor: "john doe",
					types.AttrQuoteTitle:  "quote title",
				},
				Lines: []types.InlineElements{
					{
						types.StringElement{
							Content: "a foo ",
						},
						types.InlineImage{
							Attributes: types.ElementAttributes{
								types.AttrImageAlt:    "foo",
								types.AttrImageWidth:  "",
								types.AttrImageHeight: "",
							},
							Path: "foo.png",
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("image block is NOT a quote", func() {
			actualContent := `[quote, john doe, quote title]
image::foo.png[]`
			expectedResult := types.ImageBlock{
				Attributes: types.ElementAttributes{
					types.AttrImageAlt:    "foo",
					types.AttrImageWidth:  "",
					types.AttrImageHeight: "",
					// quote attributes
					types.AttrKind:        types.Quote,
					types.AttrQuoteAuthor: "john doe",
					types.AttrQuoteTitle:  "quote title",
				},
				Path: "foo.png",
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock")) //, parser.Debug(true))
		})
	})
})
