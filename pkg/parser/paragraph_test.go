package parser_test

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("paragraphs", func() {

	Context("draft document", func() {

		Context("regular paragraphs", func() {

			It("with explicit line break", func() {
				source := `foo +
bar
baz`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{Content: "foo"},
									types.LineBreak{},
								},
								[]interface{}{
									types.StringElement{Content: "bar"},
								},
								[]interface{}{
									types.StringElement{Content: "baz"},
								},
							},
						},
					},
				}
				result, err := ParseDraftDocument(source)
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(MatchDraftDocument(expected))
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
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{Content: "foo"},
								},
								[]interface{}{
									types.StringElement{Content: "bar"},
								},
								[]interface{}{
									types.StringElement{Content: "baz"},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})

			It("with paragraph title attribute", func() {
				source := `[title=My Title]
foo
baz`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.Attributes{
								types.AttrTitle: "My Title",
							},
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{Content: "foo"},
								},
								[]interface{}{
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
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{Content: "foo"},
								},
								[]interface{}{
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
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{Content: "foo"},
								},
								[]interface{}{
									types.StringElement{Content: "bar"},
								},
								[]interface{}{
									types.StringElement{Content: "baz"},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("not treat plusplus as line break", func() {
				source := `C++
foo`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{Content: "C++"},
								},
								[]interface{}{
									types.StringElement{Content: "foo"},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with counters", func() {
				source := `foo{counter:foo} bar{counter2:foo} baz{counter:foo} bob{counter:bob}`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{Content: "foo1 bar baz3 bob1"},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			Context("with substitutions", func() {

				// using the same input for all substitution tests
				source := `:github-url: https://github.com
					
[subs="$(SUBS)"]
a link to https://github.com[] <using the *inline link macro*>
another one using attribute substitution: {github-url}[]...
// a single-line comment`

				It("should apply the 'none' substitution", func() {
					s := strings.ReplaceAll(source, "$(SUBS)", "none")
					expected := types.DraftDocument{
						Attributes: types.Attributes{
							"github-url": "https://github.com",
						},
						Blocks: []interface{}{
							types.AttributeDeclaration{
								Name:  "github-url",
								Value: "https://github.com",
							},
							types.BlankLine{},
							types.Paragraph{
								Attributes: types.Attributes{
									types.AttrSubstitutions: "none",
								},
								Lines: []interface{}{
									[]interface{}{
										types.StringElement{Content: "a link to https://github.com[] <using the *inline link macro*>"},
									},
									[]interface{}{
										types.StringElement{Content: "another one using attribute substitution: {github-url}[]..."},
									},
									[]interface{}{
										types.SingleLineComment{
											Content: " a single-line comment",
										},
									},
								},
							},
						},
					}
					Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
				})

				It("should apply the 'quotes' substitution on multiple lines", func() {
					// quoted text is parsed but inline link macro is not
					s := strings.ReplaceAll(source, "$(SUBS)", "quotes")
					expected := types.DraftDocument{
						Attributes: types.Attributes{
							"github-url": "https://github.com",
						},
						Blocks: []interface{}{
							types.AttributeDeclaration{
								Name:  "github-url",
								Value: "https://github.com",
							},
							types.BlankLine{},
							types.Paragraph{
								Attributes: types.Attributes{
									types.AttrSubstitutions: "quotes",
								},
								Lines: []interface{}{
									[]interface{}{
										types.StringElement{
											Content: "a link to https://github.com[] <using the ",
										},
										types.QuotedText{
											Kind: types.Bold,
											Elements: []interface{}{
												types.StringElement{
													Content: "inline link macro",
												},
											},
										},
										types.StringElement{
											Content: ">",
										},
									},
									[]interface{}{
										types.StringElement{
											Content: "another one using attribute substitution: {github-url}[]...",
										},
									},
									[]interface{}{
										types.SingleLineComment{
											Content: " a single-line comment",
										},
									},
								},
							},
						},
					}
					Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
				})

				It("should apply the 'macros' substitution on multiple lines", func() {
					// quoted text is not parsed but inline link macro is
					s := strings.ReplaceAll(source, "$(SUBS)", "macros")
					expected := types.DraftDocument{
						Attributes: types.Attributes{
							"github-url": "https://github.com",
						},
						Blocks: []interface{}{
							types.AttributeDeclaration{
								Name:  "github-url",
								Value: "https://github.com",
							},
							types.BlankLine{},
							types.Paragraph{
								Attributes: types.Attributes{
									types.AttrSubstitutions: "macros",
								},
								Lines: []interface{}{
									[]interface{}{
										types.StringElement{
											Content: "a link to ",
										},
										types.InlineLink{
											Location: types.Location{
												Scheme: "https://",
												Path: []interface{}{
													types.StringElement{
														Content: "github.com",
													},
												},
											},
										},
										types.StringElement{
											Content: " <using the *inline link macro*>",
										},
									},
									[]interface{}{
										types.StringElement{
											Content: "another one using attribute substitution: {github-url}[]...",
										},
									},
									[]interface{}{
										types.SingleLineComment{
											Content: " a single-line comment",
										},
									},
								},
							},
						},
					}
					Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
				})

				It("should apply the 'attributes' substitution on multiple lines", func() {
					// quoted text is not parsed but inline link macro is
					s := strings.ReplaceAll(source, "$(SUBS)", "attributes")
					expected := types.DraftDocument{
						Attributes: types.Attributes{
							"github-url": "https://github.com",
						},
						Blocks: []interface{}{
							types.AttributeDeclaration{
								Name:  "github-url",
								Value: "https://github.com",
							},
							types.BlankLine{},
							types.Paragraph{
								Attributes: types.Attributes{
									types.AttrSubstitutions: "attributes",
								},
								Lines: []interface{}{
									[]interface{}{
										types.StringElement{Content: "a link to https://github.com[] <using the *inline link macro*>"},
									},
									[]interface{}{
										types.StringElement{Content: "another one using attribute substitution: "},
										types.InlineLink{
											Location: types.Location{
												Scheme: "https://",
												Path: []interface{}{
													types.StringElement{
														Content: "github.com",
													},
												},
											},
										},
										types.StringElement{Content: "..."},
									},
									[]interface{}{
										types.SingleLineComment{
											Content: " a single-line comment",
										},
									},
								},
							},
						},
					}
					Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
				})

				It("should apply the specialchars substitution on multiple lines", func() {
					// quoted text is not parsed but inline link macro is
					s := strings.ReplaceAll(source, "$(SUBS)", "specialchars")
					expected := types.DraftDocument{
						Attributes: types.Attributes{
							"github-url": "https://github.com",
						},
						Blocks: []interface{}{
							types.AttributeDeclaration{
								Name:  "github-url",
								Value: "https://github.com",
							},
							types.BlankLine{},
							types.Paragraph{
								Attributes: types.Attributes{
									types.AttrSubstitutions: "specialchars",
								},
								Lines: []interface{}{
									[]interface{}{
										types.StringElement{Content: "a link to https://github.com[] "},
										types.SpecialCharacter{Name: "<"},
										types.StringElement{Content: "using the *inline link macro*"},
										types.SpecialCharacter{Name: ">"},
									},
									[]interface{}{
										types.StringElement{Content: "another one using attribute substitution: {github-url}[]..."},
									},
									[]interface{}{
										types.SingleLineComment{
											Content: " a single-line comment",
										},
									},
								},
							},
						},
					}
					Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
				})

				It("should apply the replacements substitution on multiple lines", func() {
					// quoted text is not parsed but inline link macro is
					s := strings.ReplaceAll(source, "$(SUBS)", "replacements")
					expected := types.DraftDocument{
						Attributes: types.Attributes{
							"github-url": "https://github.com",
						},
						Blocks: []interface{}{
							types.AttributeDeclaration{
								Name:  "github-url",
								Value: "https://github.com",
							},
							types.BlankLine{},
							types.Paragraph{
								Attributes: types.Attributes{
									types.AttrSubstitutions: "replacements",
								},
								Lines: []interface{}{
									[]interface{}{
										types.StringElement{Content: "a link to https://github.com[] <using the *inline link macro*>"},
									},
									[]interface{}{
										types.StringElement{Content: "another one using attribute substitution: {github-url}[]\u2026\u200b"},
									},
									[]interface{}{
										types.SingleLineComment{
											Content: " a single-line comment",
										},
									},
								},
							},
						},
					}
					Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
				})

				It("should apply the 'quotes' and 'macros' substitutions", func() {
					// quoted text and inline link macro are both parsed
					s := strings.ReplaceAll(source, "$(SUBS)", "quotes,macros")
					expected := types.DraftDocument{
						Attributes: types.Attributes{
							"github-url": "https://github.com",
						},
						Blocks: []interface{}{
							types.AttributeDeclaration{
								Name:  "github-url",
								Value: "https://github.com",
							},
							types.BlankLine{},
							types.Paragraph{
								Attributes: types.Attributes{
									types.AttrSubstitutions: "quotes,macros",
								},
								Lines: []interface{}{
									[]interface{}{
										types.StringElement{
											Content: "a link to ",
										},
										types.InlineLink{
											Location: types.Location{
												Scheme: "https://",
												Path: []interface{}{
													types.StringElement{
														Content: "github.com",
													},
												},
											},
										},
										types.StringElement{
											Content: " <using the ",
										},
										types.QuotedText{
											Kind: types.Bold,
											Elements: []interface{}{
												types.StringElement{
													Content: "inline link macro",
												},
											},
										},
										types.StringElement{
											Content: ">",
										},
									},
									[]interface{}{
										types.StringElement{Content: "another one using attribute substitution: {github-url}[]..."},
									},
									[]interface{}{
										types.SingleLineComment{
											Content: " a single-line comment",
										},
									},
								},
							},
						},
					}
					Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
				})

				It("should apply the 'macros' and 'quotes' substitutions", func() {
					// quoted text and inline link macro are both parsed
					// (same as above, but with subs in reversed order)
					s := strings.ReplaceAll(source, "$(SUBS)", "macros,quotes")
					expected := types.DraftDocument{
						Attributes: types.Attributes{
							"github-url": "https://github.com",
						},
						Blocks: []interface{}{
							types.AttributeDeclaration{
								Name:  "github-url",
								Value: "https://github.com",
							},
							types.BlankLine{},
							types.Paragraph{
								Attributes: types.Attributes{
									types.AttrSubstitutions: "macros,quotes",
								},
								Lines: []interface{}{
									[]interface{}{
										types.StringElement{
											Content: "a link to ",
										},
										types.InlineLink{
											Location: types.Location{
												Scheme: "https://",
												Path: []interface{}{
													types.StringElement{
														Content: "github.com",
													},
												},
											},
										},
										types.StringElement{
											Content: " <using the ",
										},
										types.QuotedText{
											Kind: types.Bold,
											Elements: []interface{}{
												types.StringElement{
													Content: "inline link macro",
												},
											},
										},
										types.StringElement{
											Content: ">",
										},
									},
									[]interface{}{
										types.StringElement{Content: "another one using attribute substitution: {github-url}[]..."},
									},
									[]interface{}{
										types.SingleLineComment{
											Content: " a single-line comment",
										},
									},
								},
							},
						},
					}
					Expect(ParseDraftDocument(s)).To(MatchDraftDocument(expected))
				})
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
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{Content: "this is a note."},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
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
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{Content: "this is a multiline"},
								},
								[]interface{}{
									types.StringElement{Content: "warning!"},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
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
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{Content: "this is a note."},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
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
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{Content: "this is a caution!"},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
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
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{Content: "this is a "},
								},
								[]interface{}{
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
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
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
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{Content: "No space after the [NOTE]!"},
								},
							},
						},
						types.BlankLine{},
						types.Paragraph{
							Attributes: types.Attributes{
								types.AttrAdmonitionKind: types.Caution,
							},
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{Content: "And no space after [CAUTION] either."},
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
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{Content: "I am a verse paragraph."},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
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
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{Content: "I am a verse paragraph."},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
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
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{Content: "I am a verse paragraph."},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
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
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{Content: "I am a verse paragraph."},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
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
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{Content: "I am a verse paragraph."},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
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
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{Content: "I am a verse paragraph."},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("image block as a verse", func() {
				// assume that the author meant to use an image, so the `verse` attribute will be ignored during rendering
				source := `[verse, john doe, verse title]
image::foo.png[]`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.ImageBlock{
							Attributes: types.Attributes{
								types.AttrImageAlt:    "verse",
								types.AttrWidth:       "john doe",
								types.AttrImageHeight: "verse title",
							},
							Location: types.Location{
								Path: []interface{}{
									types.StringElement{Content: "foo.png"},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
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
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{Content: "I am a quote paragraph."},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
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
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{Content: "I am a quote paragraph."},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
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
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{Content: "I am a quote paragraph."},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
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
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{Content: "I am a quote paragraph."},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
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
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{Content: "I am a quote paragraph."},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
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
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{Content: "I am a quote paragraph."},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("image block is NOT a quote", func() {
				source := `[quote, john doe, quote title]
image::foo.png[]`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.ImageBlock{
							Location: types.Location{
								Scheme: "",
								Path:   []interface{}{types.StringElement{Content: "foo.png"}},
							},
							Attributes: types.Attributes{
								types.AttrImageAlt:    "quote",
								types.AttrWidth:       "john doe",
								types.AttrImageHeight: "quote title",
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
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
							Lines: []interface{}{
								[]interface{}{
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
							Lines: []interface{}{
								[]interface{}{
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
							Lines: []interface{}{
								[]interface{}{
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
							Lines: []interface{}{
								[]interface{}{
									types.PredefinedAttribute{
										Name: "blank",
									},
								},
							},
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
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{Content: "hello "},
									types.PredefinedAttribute{Name: "plus"},
									types.StringElement{Content: " world"},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			Context("with substitutions", func() {

				It("should apply the 'none' substitution", func() {
					source := `[subs="none"]
a *link* to https://github.com[] <here>`
					expected := types.Document{
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.Attributes{
									types.AttrSubstitutions: "none",
								},
								Lines: []interface{}{
									[]interface{}{
										types.StringElement{Content: "a *link* to https://github.com[] <here>"},
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("should apply the 'macros' substitution", func() {
					// quoted text is not parsed but inline link macro is
					source := `[subs="macros"]
a *link* to https://github.com[] <here>`
					expected := types.Document{
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.Attributes{
									types.AttrSubstitutions: "macros",
								},
								Lines: []interface{}{
									[]interface{}{
										types.StringElement{
											Content: "a *link* to ",
										},
										types.InlineLink{
											Location: types.Location{
												Scheme: "https://",
												Path: []interface{}{
													types.StringElement{
														Content: "github.com",
													},
												},
											},
										},
										types.StringElement{
											Content: " <here>",
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("should apply the 'attributes' substitution", func() {
					// quoted text is not parsed but inline link macro is
					source := `:github-url: https://github.com

[subs="attributes"]
a *link* to {github-url} <here>`
					expected := types.Document{
						Attributes: types.Attributes{
							"github-url": "https://github.com",
						},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.Attributes{
									types.AttrSubstitutions: "attributes",
								},
								Lines: []interface{}{
									[]interface{}{
										types.StringElement{
											Content: "a *link* to ",
										},
										types.InlineLink{ // converted into a link
											Location: types.Location{
												Scheme: "https://",
												Path: []interface{}{
													types.StringElement{
														Content: "github.com",
													},
												},
											},
										},
										types.StringElement{
											Content: " ",
										},
										types.SpecialCharacter{
											Name: "<",
										},
										types.StringElement{
											Content: "here",
										},
										types.SpecialCharacter{
											Name: ">",
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("should apply the 'macros' and 'quotes' substitutions", func() {
					// quoted text and inline link macro are both parsed
					// (same as above, but with subs in reversed order)
					source := `[subs="quotes,macros"]
a *link* to https://github.com[] <here>`
					expected := types.DraftDocument{
						Blocks: []interface{}{
							types.Paragraph{
								Attributes: types.Attributes{
									types.AttrSubstitutions: "quotes,macros",
								},
								Lines: []interface{}{
									[]interface{}{
										types.StringElement{
											Content: "a ",
										},
										types.QuotedText{
											Kind: types.Bold,
											Elements: []interface{}{
												types.StringElement{
													Content: "link",
												},
											},
										},
										types.StringElement{
											Content: " to ",
										},
										types.InlineLink{
											Location: types.Location{
												Scheme: "https://",
												Path: []interface{}{
													types.StringElement{
														Content: "github.com",
													},
												},
											},
										},
										types.StringElement{
											Content: " <here>",
										},
									},
								},
							},
						},
					}
					Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
				})
			})
		})

		Context("quote paragraphs", func() {

			It("inline image within a quote", func() {
				source := `[quote, john doe, quote title]
a foo image:foo.png[]`
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.Attributes{
								types.AttrKind:        types.Quote,
								types.AttrQuoteAuthor: "john doe",
								types.AttrQuoteTitle:  "quote title",
							},
							Lines: []interface{}{
								[]interface{}{
									types.StringElement{
										Content: "a foo ",
									},
									types.InlineImage{
										Attributes: types.Attributes{
											types.AttrImageAlt: "foo",
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
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})

	})
})
