package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("paragraphs", func() {

	Context("draft document", func() {

		Context("default paragraphs", func() {

			It("paragraph with 1 word", func() {
				source := "hello"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "hello"},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})

			It("paragraph with few words and ending with spaces", func() {
				source := "a paragraph with some content  "
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a paragraph with some content  "},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})

			It("paragraph with bold content and spaces", func() {
				source := "a paragraph with *some bold content*  "
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
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
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})

			It("paragraph with non-alphanum character before bold text", func() {
				source := "+*some bold content*"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
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
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})

			It("paragraph with id and title", func() {
				source := `[#foo]
.a title
a paragraph`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.Attributes{
								types.AttrID:       "foo",
								types.AttrCustomID: true,
								types.AttrTitle:    "a title",
							},
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a paragraph"},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})

			It("paragraph with words and dots on same line", func() {
				source := `foo. bar.`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "foo. bar."},
								},
							},
						},
					},
				}

				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})

			It("paragraph with words and dots on two lines", func() {
				source := `foo. 
bar.`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
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
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})

		})

		Context("paragraphs with line break", func() {

			It("with explicit line break", func() {
				source := `foo +
bar
baz`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
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
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})

			It("with paragraph attribute", func() {
				source := `[%hardbreaks]
foo
bar
baz`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.Attributes{
								types.AttrOptions: map[string]bool{"hardbreaks": true},
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
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})

			It("with paragraph multiple attribute", func() {
				source := `[%hardbreaks.role1.role2]
[#anchor]
foo
baz`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.Attributes{
								types.AttrCustomID: true,
								types.AttrID:       "anchor",
								types.AttrRole:     []string{"role1", "role2"},
								types.AttrOptions:  map[string]bool{"hardbreaks": true},
							},
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "foo"},
								},
								{
									types.StringElement{Content: "baz"},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})

			It("with paragraph roles and attribute", func() {
				source := `[.role1%hardbreaks.role2]
foo
bar
baz`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.Attributes{
								types.AttrOptions: map[string]bool{"hardbreaks": true},
								types.AttrRole:    []string{"role1", "role2"},
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
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})

			It("not treat plusplus as line break", func() {
				source := `C++
foo`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "C++"},
								},
								{
									types.StringElement{Content: "foo"},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})
		})

		Context("admonition paragraphs", func() {

			It("note admonition paragraph", func() {
				source := `NOTE: this is a note.`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.Attributes{
								types.AttrAdmonitionKind: types.Note,
							},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "this is a note.",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})

			It("warning admonition paragraph", func() {
				source := `WARNING: this is a multiline
warning!`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.Attributes{
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
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})

			It("admonition note paragraph with id and title", func() {
				source := `[[foo]]
.bar
NOTE: this is a note.`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.Attributes{
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
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})

			It("caution admonition paragraph with single line", func() {
				source := `[CAUTION]
this is a caution!`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.Attributes{
								types.AttrAdmonitionKind: types.Caution,
							},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "this is a caution!",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})

			It("multiline caution admonition paragraph with title and id", func() {
				source := `[[foo]]
[CAUTION] 
.bar
this is a 
*caution*!`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.Attributes{
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
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})

			It("multiple admonition paragraphs", func() {
				source := `[NOTE]
No space after the [NOTE]!

[CAUTION]
And no space after [CAUTION] either.`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.Attributes{
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
							Attributes: types.Attributes{
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
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})
		})

		Context("verse paragraphs", func() {

			It("paragraph as a verse with author and title", func() {
				source := `[verse, john doe, verse title]
I am a verse paragraph.`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.Attributes{
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
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})

			It("paragraph as a verse with author, title and other attributes", func() {
				source := `[[universal]]
[verse, john doe, verse title]
.universe
I am a verse paragraph.`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.Attributes{
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
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})

			It("paragraph as a verse with empty title", func() {
				source := `[verse, john doe, ]
I am a verse paragraph.`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.Attributes{
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
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})

			It("paragraph as a verse without title", func() {
				source := `[verse, john doe ]
I am a verse paragraph.`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.Attributes{
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
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})

			It("paragraph as a verse with empty author", func() {
				source := `[verse,  ]
I am a verse paragraph.`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.Attributes{
								types.AttrKind: types.Verse,
							},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "I am a verse paragraph.",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})

			It("paragraph as a verse without author", func() {
				source := `[verse]
I am a verse paragraph.`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.Attributes{
								types.AttrKind: types.Verse,
							},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "I am a verse paragraph.",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})

			It("image block as a verse", func() {
				source := `[verse, john doe, verse title]
image::foo.png[]`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.Attributes{
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
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})
		})

		Context("quote paragraphs", func() {

			It("paragraph as a quote with author and title", func() {
				source := `[quote, john doe, quote title]
I am a quote paragraph.`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.Attributes{
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
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})

			It("paragraph as a quote with author, title and other attributes", func() {
				source := `[[universal]]
[quote, john doe, quote title]
.universe
I am a quote paragraph.`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.Attributes{
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
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})

			It("paragraph as a quote with empty title", func() {
				source := `[quote, john doe, ]
I am a quote paragraph.`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.Attributes{
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
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})

			It("paragraph as a quote without title", func() {
				source := `[quote, john doe ]
I am a quote paragraph.`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.Attributes{
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
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})

			It("paragraph as a quote with empty author", func() {
				source := `[quote,  ]
I am a quote paragraph.`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.Attributes{
								types.AttrKind: types.Quote,
							},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "I am a quote paragraph.",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})

			It("paragraph as a quote without author", func() {
				source := `[quote]
I am a quote paragraph.`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.Attributes{
								types.AttrKind: types.Quote,
							},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "I am a quote paragraph.",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})

			It("inline image within a quote", func() {
				source := `[quote, john doe, quote title]
a foo image:foo.png[]`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.Attributes{
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
										Location: types.Location{
											Path: []interface{}{
												types.StringElement{
													Content: "foo.png",
												},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})

			It("image block is NOT a quote", func() {
				source := `[quote, john doe, quote title]
image::foo.png[]`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.ImageBlock{
							Attributes: types.Attributes{

								// quote attributes
								types.AttrKind:        types.Quote,
								types.AttrQuoteAuthor: "john doe",
								types.AttrQuoteTitle:  "quote title",
							},
							Location: types.Location{
								Path: []interface{}{
									types.StringElement{
										Content: "foo.png",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})
		})

		Context("thematic breaks", func() {
			It("thematic break form1 by itself", func() {
				source := "***"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.ThematicBreak{},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})
			It("thematic break form2 by itself", func() {
				source := "* * *"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.ThematicBreak{},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})
			It("thematic break form3 by itself", func() {
				source := "---"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.ThematicBreak{},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})
			It("thematic break form4 by itself", func() {
				source := "- - -"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.ThematicBreak{},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})
			It("thematic break form5 by itself", func() {
				source := "___"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.ThematicBreak{},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})
			It("thematic break form4 by itself", func() {
				source := "_ _ _"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.ThematicBreak{},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})
			It("thematic break with leading text", func() {
				source := "text ***"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "text ***"},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})

			// NB: three asterisks gets confused with bullets if with trailing text
			It("thematic break with trailing text", func() {
				source := "* * * text"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "* * * text"},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})

		})
	})

	Context("final document", func() {

		Context("default paragraph", func() {

			It("paragraph with custom id prefix and title", func() {
				source := `:idprefix: bar_
			
.a title
a paragraph`
				expected := types.Document{
					Attributes: types.Attributes{
						types.AttrIDPrefix: "bar_",
					},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.Attributes{
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
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("empty paragraph", func() {
				source := `{blank}`
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("paragraph with predefined attribute", func() {
				source := "hello {plus} world"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "hello &#43; world"},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})
	})
})
