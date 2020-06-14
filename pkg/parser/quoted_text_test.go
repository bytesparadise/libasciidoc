package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("quoted texts", func() {

	Context("draft document", func() {

		Context("quoted text with single punctuation", func() {

			It("bold text with 1 word", func() {
				source := "*hello*"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: "hello"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("bold text with 2 words", func() {
				source := "*bold    content*"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: "bold    content"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("bold text with 3 words", func() {
				source := "*some bold content*"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
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
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("italic text with 3 words in single quote", func() {
				source := "_some italic content_"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Italic,
										Elements: []interface{}{
											types.StringElement{Content: "some italic content"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("monospace text with 3 words", func() {
				source := "`some monospace content`"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "some monospace content"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("invalid subscript text with 3 words", func() {
				source := "~some subscript content~"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "~some subscript content~"},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("invalid superscript text with 3 words", func() {
				source := "^some superscript content^"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "^some superscript content^"},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("bold text within italic text", func() {
				source := "_some *bold* content_"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Italic,
										Elements: []interface{}{
											types.StringElement{Content: "some "},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{Content: "bold"},
												},
											},
											types.StringElement{Content: " content"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("monospace text within bold text within italic quote", func() {
				source := "*some _italic and `monospaced content`_*"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: "some "},
											types.QuotedText{
												Kind: types.Italic,
												Elements: []interface{}{
													types.StringElement{Content: "italic and "},
													types.QuotedText{
														Kind: types.Monospace,
														Elements: []interface{}{
															types.StringElement{Content: "monospaced content"},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("italic text within italic text", func() {
				source := "_some _very italic_ content_"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Italic,
										Elements: []interface{}{
											types.StringElement{Content: "some _very italic"},
										},
									},
									types.StringElement{Content: " content_"},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("bold delimiter text within bold text", func() {
				source := "*bold*content*"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: "bold*content"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("italic delimiter text within italic text", func() {
				source := "_italic_content_"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Italic,
										Elements: []interface{}{
											types.StringElement{Content: "italic_content"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("monospace delimiter text within monospace text", func() {
				source := "`monospace`content`"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "monospace`content"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("non-bold text then bold text", func() {
				source := "non*bold*content *bold content*"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "non*bold*content ",
									},
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: "bold content"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})
			It("non-italic text then italic text", func() {
				source := "non_italic_content _italic content_"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "non_italic_content ",
									},
									types.QuotedText{
										Kind: types.Italic,
										Elements: []interface{}{
											types.StringElement{Content: "italic content"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("non-monospace text then monospace text", func() {
				source := "non`monospace`content `monospace content`"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "non`monospace`content ",
									},
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "monospace content"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("subscript text attached", func() {
				source := "O~2~ is a molecule"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "O"},
									types.QuotedText{
										Kind: types.Subscript,
										Elements: []interface{}{
											types.StringElement{Content: "2"},
										},
									},
									types.StringElement{Content: " is a molecule"},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("superscript text attached", func() {
				source := "M^me^ White"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "M"},
									types.QuotedText{
										Kind: types.Superscript,
										Elements: []interface{}{
											types.StringElement{Content: "me"},
										},
									},
									types.StringElement{Content: " White"},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("invalid subscript text with 3 words", func() {
				source := "~some subscript content~"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "~some subscript content~"},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})
		})

		Context("Quoted text with double punctuation", func() {

			It("bold text of 1 word in double quote", func() {
				source := "**hello**"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: "hello"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})

			It("italic text with 3 words in double quote", func() {
				source := "__some italic content__"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Italic,
										Elements: []interface{}{
											types.StringElement{Content: "some italic content"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("monospace text with 3 words in double quote", func() {
				source := "``some monospace content``"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "some monospace content"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("superscript text within italic text", func() {
				source := "__some ^superscript^ content__"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Italic,
										Elements: []interface{}{
											types.StringElement{Content: "some "},
											types.QuotedText{
												Kind: types.Superscript,
												Elements: []interface{}{
													types.StringElement{Content: "superscript"},
												},
											},
											types.StringElement{Content: " content"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("superscript text within italic text within bold quote", func() {
				source := "**some _italic and ^superscriptcontent^_**"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: "some "},
											types.QuotedText{
												Kind: types.Italic,
												Elements: []interface{}{
													types.StringElement{Content: "italic and "},
													types.QuotedText{
														Kind: types.Superscript,
														Elements: []interface{}{
															types.StringElement{Content: "superscriptcontent"},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})
		})

		Context("Quoted text inline", func() {

			It("inline content with bold text", func() {
				source := "a paragraph with *some bold content*"
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
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("inline content with invalid bold text - use case 1", func() {
				source := "a paragraph with *some bold content"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a paragraph with *some bold content"},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("inline content with invalid bold text - use case 2", func() {
				source := "a paragraph with *some bold content *"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a paragraph with *some bold content *"},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("inline content with invalid bold text - use case 3", func() {
				source := "a paragraph with * some bold content*"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a paragraph with * some bold content*"},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("invalid italic text within bold text", func() {
				source := "some *bold and _italic content _ together*."
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{

									types.StringElement{Content: "some "},
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: "bold and _italic content _ together"},
										},
									},
									types.StringElement{Content: "."},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("italic text within invalid bold text", func() {
				source := "some *bold and _italic content_ together *."
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "some *bold and "},
									types.QuotedText{
										Kind: types.Italic,
										Elements: []interface{}{
											types.StringElement{Content: "italic content"},
										},
									},
									types.StringElement{Content: " together *."},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("inline content with invalid subscript text - use case 1", func() {
				source := "a paragraph with ~some subscript content"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a paragraph with ~some subscript content"},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("inline content with invalid subscript text - use case 2", func() {
				source := "a paragraph with ~some subscript content ~"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a paragraph with ~some subscript content ~"},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("inline content with invalid subscript text - use case 3", func() {
				source := "a paragraph with ~ some subscript content~"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a paragraph with ~ some subscript content~"},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("inline content with invalid superscript text - use case 1", func() {
				source := "a paragraph with ^some superscript content"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{

									types.StringElement{Content: "a paragraph with ^some superscript content"},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("inline content with invalid superscript text - use case 2", func() {
				source := "a paragraph with ^some superscript content ^"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{

									types.StringElement{Content: "a paragraph with ^some superscript content ^"},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("inline content with invalid superscript text - use case 3", func() {
				source := "a paragraph with ^ some superscript content^"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{

									types.StringElement{Content: "a paragraph with ^ some superscript content^"},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})
		})
		Context("attributes", func() {
			It("simple role italics", func() {
				source := "[myrole]_italics_"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Italic,
										Elements: []interface{}{
											types.StringElement{Content: "italics"},
										},
										Attributes: types.Attributes{
											types.AttrRole: "myrole",
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("simple role italics unconstrained", func() {
				source := "it[uncle]__al__ic"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "it",
									},
									types.QuotedText{
										Kind: types.Italic,
										Elements: []interface{}{
											types.StringElement{Content: "al"},
										},
										Attributes: types.Attributes{
											types.AttrRole: "uncle",
										},
									},
									types.StringElement{
										Content: "ic",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("simple role bold", func() {
				source := "[myrole]*bold*"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: "bold"},
										},
										Attributes: types.Attributes{
											types.AttrRole: "myrole",
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("simple role bold unconstrained", func() {
				source := "it[uncle]**al**ic"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "it",
									},
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: "al"},
										},
										Attributes: types.Attributes{
											types.AttrRole: "uncle",
										},
									},
									types.StringElement{
										Content: "ic",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("simple role mono", func() {
				source := "[myrole]`true`"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "true"},
										},
										Attributes: types.Attributes{
											types.AttrRole: "myrole",
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("simple role mono unconstrained", func() {
				source := "int[uncle]``eg``rate"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "int",
									},
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "eg"},
										},
										Attributes: types.Attributes{
											types.AttrRole: "uncle",
										},
									},
									types.StringElement{
										Content: "rate",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("role with comma truncates", func() {
				source := "[myrole,and nothing else]_italics_"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Italic,
										Elements: []interface{}{
											types.StringElement{Content: "italics"},
										},
										Attributes: types.Attributes{
											types.AttrRole: "myrole",
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("short-hand ID only", func() {
				source := "[#here]*bold*"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: "bold"},
										},
										Attributes: types.Attributes{
											types.AttrID:       "here",
											types.AttrCustomID: true,
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("short-hand role only", func() {
				source := "[.bob]**bold**"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: "bold"},
										},
										Attributes: types.Attributes{
											types.AttrRole: "bob",
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("short-hand multiple roles and id", func() {
				source := "[.r1#anchor.r2.r3]**bold**[#here.second.class]_text_"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: "bold"},
										},
										Attributes: types.Attributes{
											types.AttrRole:     []string{"r1", "r2", "r3"},
											types.AttrID:       "anchor",
											types.AttrCustomID: true,
										},
									},
									types.QuotedText{
										Kind: types.Italic,
										Elements: []interface{}{
											types.StringElement{Content: "text"},
										},
										Attributes: types.Attributes{
											types.AttrRole:     []string{"second", "class"},
											types.AttrID:       "here",
											types.AttrCustomID: true,
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("empty role", func() {
				source := "[]**bold**"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: "bold"},
										},
										Attributes: types.Attributes{
											types.AttrRole: "",
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			// This demonstrates that we cannot inject malicious data in these attributes.
			// The content is escaped by the renderer, not the parser.
			It("embedded garbage", func() {
				source := "[.<something \"wicked>]**bold**"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: "bold"},
										},
										Attributes: types.Attributes{
											types.AttrRole: "<something \"wicked>",
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

		})

		Context("nested quoted text", func() {

			It("italic text within bold text", func() {
				source := "some *bold and _italic content_ together*."
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "some "},
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: "bold and "},
											types.QuotedText{
												Kind: types.Italic,
												Elements: []interface{}{
													types.StringElement{Content: "italic content"},
												},
											},
											types.StringElement{Content: " together"},
										},
									},
									types.StringElement{Content: "."},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("single-quote bold within single-quote bold text", func() {
				// here we don't allow for bold text within bold text, to comply with the existing implementations (asciidoc and asciidoctor)
				source := "*some *nested bold* content*"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: "some *nested bold"},
										},
									},
									types.StringElement{Content: " content*"},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("double-quote bold within double-quote bold text", func() {
				// here we don't allow for bold text within bold text, to comply with the existing implementations (asciidoc and asciidoctor)
				source := "**some **nested bold** content**"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: "some "},
										},
									},
									types.StringElement{Content: "nested bold"},
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: " content"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("single-quote bold within double-quote bold text", func() {
				// here we don't allow for bold text within bold text, to comply with the existing implementations (asciidoc and asciidoctor)
				source := "**some *nested bold* content**"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: "some "},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{Content: "nested bold"},
												},
											},
											types.StringElement{Content: " content"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("double-quote bold within single-quote bold text", func() {
				// here we don't allow for bold text within bold text, to comply with the existing implementations (asciidoc and asciidoctor)
				source := "*some **nested bold** content*"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: "some "},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{Content: "nested bold"},
												},
											},
											types.StringElement{Content: " content"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("single-quote italic within single-quote italic text", func() {
				// here we don't allow for italic text within italic text, to comply with the existing implementations (asciidoc and asciidoctor)
				source := "_some _nested italic_ content_"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Italic,
										Elements: []interface{}{
											types.StringElement{Content: "some _nested italic"},
										},
									},
									types.StringElement{Content: " content_"},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("double-quote italic within double-quote italic text", func() {
				// here we don't allow for italic text within italic text, to comply with the existing implementations (asciidoc and asciidoctor)
				source := "__some __nested italic__ content__"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Italic,
										Elements: []interface{}{
											types.StringElement{Content: "some "},
										},
									},
									types.StringElement{Content: "nested italic"},
									types.QuotedText{
										Kind: types.Italic,
										Elements: []interface{}{
											types.StringElement{Content: " content"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("single-quote italic within double-quote italic text", func() {
				// here we allow for italic text within italic text, to comply with the existing implementations (asciidoc and asciidoctor)
				source := "_some __nested italic__ content_"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Italic,
										Elements: []interface{}{
											types.StringElement{Content: "some "},
											types.QuotedText{
												Kind: types.Italic,
												Elements: []interface{}{
													types.StringElement{Content: "nested italic"},
												},
											},
											types.StringElement{Content: " content"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("double-quote italic within single-quote italic text", func() {
				// here we allow for italic text within italic text, to comply with the existing implementations (asciidoc and asciidoctor)
				source := "_some __nested italic__ content_"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Italic,
										Elements: []interface{}{
											types.StringElement{Content: "some "},
											types.QuotedText{
												Kind: types.Italic,
												Elements: []interface{}{
													types.StringElement{Content: "nested italic"},
												},
											},
											types.StringElement{Content: " content"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("single-quote monospace within single-quote monospace text", func() {
				// here we don't allow for monospace text within monospace text, to comply with the existing implementations (asciidoc and asciidoctor)
				source := "`some `nested monospace` content`"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "some `nested monospace"},
										},
									},
									types.StringElement{Content: " content`"},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("double-quote monospace within double-quote monospace text", func() {
				// here we don't allow for monospace text within monospace text, to comply with the existing implementations (asciidoc and asciidoctor)
				source := "``some ``nested monospace`` content``"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "some "},
										},
									},
									types.StringElement{Content: "nested monospace"},
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: " content"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("single-quote monospace within double-quote monospace text", func() {
				// here we allow for monospace text within monospace text, to comply with the existing implementations (asciidoc and asciidoctor)
				source := "`some ``nested monospace`` content`"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "some "},
											types.QuotedText{
												Kind: types.Monospace,
												Elements: []interface{}{
													types.StringElement{Content: "nested monospace"},
												},
											},
											types.StringElement{Content: " content"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("double-quote monospace within single-quote monospace text", func() {
				// here we allow for monospace text within monospace text, to comply with the existing implementations (asciidoc and asciidoctor)
				source := "`some ``nested monospace`` content`"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "some "},
											types.QuotedText{
												Kind: types.Monospace,
												Elements: []interface{}{
													types.StringElement{Content: "nested monospace"},
												},
											},
											types.StringElement{Content: " content"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("unbalanced bold in monospace - case 1", func() {
				source := "`*a`"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "*a"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("unbalanced bold in monospace - case 2", func() {
				source := "`a*b`"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "a*b"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("italic in monospace", func() {
				source := "`_a_`"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.QuotedText{
												Kind: types.Italic,
												Elements: []interface{}{
													types.StringElement{Content: "a"},
												},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("unbalanced italic in monospace", func() {
				source := "`a_b`"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "a_b"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("unparsed bold in monospace", func() {
				source := "`a*b*`"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "a*b*"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("parsed subscript in monospace", func() {
				source := "`a~b~`"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "a"},
											types.QuotedText{
												Kind: types.Subscript,
												Elements: []interface{}{
													types.StringElement{Content: "b"},
												},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("multiline in single quoted monospace - case 1", func() {
				source := "`a\nb`"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "a\nb"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("multiline in double quoted monospace - case 1", func() {
				source := "`a\nb`"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "a\nb"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("multiline in single quoted  monospace - case 2", func() {
				source := "`a\n*b*`"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "a\n"},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{Content: "b"},
												},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("multiline in double quoted  monospace - case 2", func() {
				source := "`a\n*b*`"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "a\n"},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{Content: "b"},
												},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("link in bold", func() {
				source := "*a link:/[b]*"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: "a "},
											types.InlineLink{
												Attributes: types.Attributes{
													"positional-1": []interface{}{
														types.StringElement{
															Content: "b",
														},
													},
												},
												Location: types.Location{
													Path: []interface{}{
														types.StringElement{
															Content: "/",
														},
													},
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

			It("image in bold", func() {
				source := "*a image:foo.png[]*"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: "a "},
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
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("singleplus passthrough in bold", func() {
				source := "*a +image:foo.png[]+*"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: "a "},
											types.InlinePassthrough{
												Kind: types.SinglePlusPassthrough,
												Elements: []interface{}{
													types.StringElement{Content: "image:foo.png[]"},
												},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("tripleplus passthrough in bold", func() {
				source := "*a +++image:foo.png[]+++*"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: "a "},
											types.InlinePassthrough{
												Kind: types.TriplePlusPassthrough,
												Elements: []interface{}{
													types.StringElement{Content: "image:foo.png[]"},
												},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("link in italic", func() {
				source := "_a link:/[b]_"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Italic,
										Elements: []interface{}{
											types.StringElement{Content: "a "},
											types.InlineLink{
												Attributes: types.Attributes{
													"positional-1": []interface{}{
														types.StringElement{
															Content: "b",
														},
													},
												},
												Location: types.Location{
													Path: []interface{}{
														types.StringElement{
															Content: "/",
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("image in italic", func() {
				source := "_a image:foo.png[]_"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Italic,
										Elements: []interface{}{
											types.StringElement{Content: "a "},
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
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("singleplus passthrough in italic", func() {
				source := "_a +image:foo.png[]+_"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Italic,
										Elements: []interface{}{
											types.StringElement{Content: "a "},
											types.InlinePassthrough{
												Kind: types.SinglePlusPassthrough,
												Elements: []interface{}{
													types.StringElement{Content: "image:foo.png[]"},
												},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("tripleplus passthrough in italic", func() {
				source := "_a +++image:foo.png[]+++_"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Italic,
										Elements: []interface{}{
											types.StringElement{Content: "a "},
											types.InlinePassthrough{
												Kind: types.TriplePlusPassthrough,
												Elements: []interface{}{
													types.StringElement{Content: "image:foo.png[]"},
												},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("link in monospace", func() {
				source := "`a link:/[b]`"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "a "},
											types.InlineLink{
												Attributes: types.Attributes{
													"positional-1": []interface{}{
														types.StringElement{
															Content: "b",
														},
													},
												},
												Location: types.Location{
													Path: []interface{}{
														types.StringElement{
															Content: "/",
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("image in monospace", func() {
				source := "`a image:foo.png[]`"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "a "},
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
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("singleplus passthrough in monospace", func() {
				source := "`a +image:foo.png[]+`"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{

									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "a "},
											types.InlinePassthrough{
												Kind: types.SinglePlusPassthrough,
												Elements: []interface{}{
													types.StringElement{Content: "image:foo.png[]"},
												},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("tripleplus passthrough in monospace", func() {
				source := "`a +++image:foo.png[]+++`"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{

									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "a "},
											types.InlinePassthrough{
												Kind: types.TriplePlusPassthrough,
												Elements: []interface{}{
													types.StringElement{Content: "image:foo.png[]"},
												},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

		})

		Context("unbalanced quoted text", func() {

			Context("unbalanced bold text", func() {

				It("unbalanced bold text - extra on left", func() {
					source := "**some bold content*"
					expected := types.DraftDocument{
						Blocks: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.QuotedText{
											Kind: types.Bold,
											Elements: []interface{}{
												types.StringElement{Content: "*some bold content"},
											},
										},
									},
								},
							},
						},
					}
					Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
				})

				It("unbalanced bold text - extra on right", func() {
					source := "*some bold content**"
					expected := types.DraftDocument{
						Blocks: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{

										types.QuotedText{
											Kind: types.Bold,
											Elements: []interface{}{
												types.StringElement{Content: "some bold content"},
											},
										},
										types.StringElement{Content: "*"},
									},
								},
							},
						},
					}
					Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
				})
			})

			Context("unbalanced italic text", func() {

				It("unbalanced italic text - extra on left", func() {
					source := "__some italic content_"
					expected := types.DraftDocument{
						Blocks: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{

										types.QuotedText{
											Kind: types.Italic,
											Elements: []interface{}{
												types.StringElement{Content: "_some italic content"},
											},
										},
									},
								},
							},
						},
					}
					Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
				})

				It("unbalanced italic text - extra on right", func() {
					source := "_some italic content__"
					expected := types.DraftDocument{
						Blocks: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.QuotedText{
											Kind: types.Italic,
											Elements: []interface{}{
												types.StringElement{Content: "some italic content"},
											},
										},
										types.StringElement{Content: "_"},
									},
								},
							},
						},
					}
					Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
				})
			})

			Context("unbalanced monospace text", func() {

				It("unbalanced monospace text - extra on left", func() {
					source := "``some monospace content`"
					expected := types.DraftDocument{
						Blocks: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{

										types.QuotedText{
											Kind: types.Monospace,
											Elements: []interface{}{
												types.StringElement{Content: "`some monospace content"},
											},
										},
									},
								},
							},
						},
					}
					Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
				})

				It("unbalanced monospace text - extra on right", func() {
					source := "`some monospace content``"
					expected := types.DraftDocument{
						Blocks: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.QuotedText{
											Kind: types.Monospace,
											Elements: []interface{}{
												types.StringElement{Content: "some monospace content"},
											},
										},
										types.StringElement{Content: "`"},
									},
								},
							},
						},
					}
					Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
				})
			})

			It("inline content with unbalanced bold text", func() {
				source := "a paragraph with *some bold content"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a paragraph with *some bold content"},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

		})

		Context("prevented substitution", func() {

			Context("prevented bold text substitution", func() {

				Context("without nested quoted text", func() {

					It("escaped bold text with single backslash", func() {
						source := `\*bold content*`
						expected := types.DraftDocument{
							Blocks: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "*bold content*"},
										},
									},
								},
							},
						}
						Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
					})

					It("escaped bold text with multiple backslashes", func() {
						source := `\\*bold content*`
						expected := types.DraftDocument{
							Blocks: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: `\*bold content*`},
										},
									},
								},
							},
						}
						Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
					})

					It("escaped bold text with double quote", func() {
						source := `\\**bold content**`
						expected := types.DraftDocument{
							Blocks: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: `**bold content**`},
										},
									},
								},
							},
						}
						Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
					})

					It("escaped bold text with double quote and more backslashes", func() {
						source := `\\\**bold content**`
						expected := types.DraftDocument{
							Blocks: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: `\**bold content**`},
										},
									},
								},
							},
						}
						Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
					})

					It("escaped bold text with unbalanced double quote", func() {
						source := `\**bold content*`
						expected := types.DraftDocument{
							Blocks: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: `**bold content*`},
										},
									},
								},
							},
						}
						Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
					})

					It("escaped bold text with unbalanced double quote and more backslashes", func() {
						source := `\\\**bold content*`
						expected := types.DraftDocument{
							Blocks: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: `\\**bold content*`},
										},
									},
								},
							},
						}
						Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
					})
				})

				Context("with nested quoted text", func() {

					It("escaped bold text with nested italic text", func() {
						source := `\*_italic content_*`
						expected := types.DraftDocument{
							Blocks: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "*"},
											types.QuotedText{
												Kind: types.Italic,
												Elements: []interface{}{
													types.StringElement{Content: "italic content"},
												},
											},
											types.StringElement{Content: "*"},
										},
									},
								},
							},
						}
						Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
					})

					It("escaped bold text with unbalanced double quote and nested italic test", func() {
						source := `\**_italic content_*`
						expected := types.DraftDocument{
							Blocks: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "**"},
											types.QuotedText{
												Kind: types.Italic,
												Elements: []interface{}{
													types.StringElement{Content: "italic content"},
												},
											},
											types.StringElement{Content: "*"},
										},
									},
								},
							},
						}
						Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
					})

					It("escaped bold text with nested italic", func() {
						source := `\*bold _and italic_ content*`
						expected := types.DraftDocument{
							Blocks: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "*bold "},
											types.QuotedText{
												Kind: types.Italic,
												Elements: []interface{}{
													types.StringElement{Content: "and italic"},
												},
											},
											types.StringElement{Content: " content*"},
										},
									},
								},
							},
						}
						Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
					})
				})

			})

			Context("prevented italic text substitution", func() {

				Context("without nested quoted text", func() {

					It("escaped italic text with single quote", func() {
						source := `\_italic content_`
						expected := types.DraftDocument{
							Blocks: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "_italic content_"},
										},
									},
								},
							},
						}
						Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
					})

					It("escaped italic text with single quote and more backslashes", func() {
						source := `\\_italic content_`
						expected := types.DraftDocument{
							Blocks: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: `\_italic content_`},
										},
									},
								},
							},
						}
						Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
					})

					It("escaped italic text with double quote with 2 backslashes", func() {
						source := `\\__italic content__`
						expected := types.DraftDocument{
							Blocks: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: `__italic content__`},
										},
									},
								},
							},
						}
						Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
					})

					It("escaped italic text with double quote with 3 backslashes", func() {
						source := `\\\__italic content__`
						expected := types.DraftDocument{
							Blocks: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: `\__italic content__`},
										},
									},
								},
							},
						}
						Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
					})

					It("escaped italic text with unbalanced double quote", func() {
						source := `\__italic content_`
						expected := types.DraftDocument{
							Blocks: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: `__italic content_`},
										},
									},
								},
							},
						}
						Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
					})

					It("escaped italic text with unbalanced double quote and more backslashes", func() {
						source := `\\\__italic content_`
						expected := types.DraftDocument{
							Blocks: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: `\\__italic content_`}, // only 1 backslash remove
										},
									},
								},
							},
						}
						Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
					})
				})

				Context("with nested quoted text", func() {

					It("escaped italic text with nested monospace text", func() {
						source := `\` + "_`monospace content`_" // gives: \_`monospace content`_
						expected := types.DraftDocument{
							Blocks: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "_"},
											types.QuotedText{
												Kind: types.Monospace,
												Elements: []interface{}{
													types.StringElement{Content: "monospace content"},
												},
											},
											types.StringElement{Content: "_"},
										},
									},
								},
							},
						}
						Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
					})

					It("escaped italic text with unbalanced double quote and nested bold test", func() {
						source := `\__*bold content*_`
						expected := types.DraftDocument{
							Blocks: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "__"},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{Content: "bold content"},
												},
											},
											types.StringElement{Content: "_"},
										},
									},
								},
							},
						}
						Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
					})

					It("escaped italic text with nested bold text", func() {
						source := `\_italic *and bold* content_`
						expected := types.DraftDocument{
							Blocks: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "_italic "},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{Content: "and bold"},
												},
											},
											types.StringElement{Content: " content_"},
										},
									},
								},
							},
						}
						Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
					})
				})
			})

			Context("prevented monospace text substitution", func() {

				Context("without nested quoted text", func() {

					It("escaped monospace text with single quote", func() {
						source := `\` + "`monospace content`"
						expected := types.DraftDocument{
							Blocks: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "`monospace content`"}, // backslash removed
										},
									},
								},
							},
						}
						Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
					})

					It("escaped monospace text with single quote and more backslashes", func() {
						source := `\\` + "`monospace content`"
						expected := types.DraftDocument{
							Blocks: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: `\` + "`monospace content`"}, // only 1 backslash removed
										},
									},
								},
							},
						}
						Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
					})

					It("escaped monospace text with double quote", func() {
						source := `\\` + "`monospace content``"
						expected := types.DraftDocument{
							Blocks: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: `\` + "`monospace content``"}, // 2 back slashes "consumed"
										},
									},
								},
							},
						}
						Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
					})

					It("escaped monospace text with double quote and more backslashes", func() {
						source := `\\\` + "``monospace content``" // 3 backslashes
						expected := types.DraftDocument{
							Blocks: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: `\` + "``monospace content``"}, // 2 back slashes "consumed"
										},
									},
								},
							},
						}
						Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
					})

					It("escaped monospace text with unbalanced double quote", func() {
						source := `\` + "``monospace content`"
						expected := types.DraftDocument{
							Blocks: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "``monospace content`"},
										},
									},
								},
							},
						}
						Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
					})

					It("escaped monospace text with unbalanced double quote and more backslashes", func() {
						source := `\\\` + "``monospace content`" // 3 backslashes
						expected := types.DraftDocument{
							Blocks: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: `\\` + "``monospace content`"}, // 2 backslashes removed
										},
									},
								},
							},
						}
						Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
					})
				})

				Context("with nested quoted text", func() {

					It("escaped monospace text with nested bold text", func() {
						source := `\` + "`*bold content*`" // gives: \`*bold content*`
						expected := types.DraftDocument{
							Blocks: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "`"},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{Content: "bold content"},
												},
											},
											types.StringElement{Content: "`"},
										},
									},
								},
							},
						}
						Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
					})

					It("escaped monospace text with unbalanced double backquote and nested bold test", func() {
						source := `\` + "``*bold content*`"
						expected := types.DraftDocument{
							Blocks: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "``"},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{Content: "bold content"},
												},
											},
											types.StringElement{Content: "`"},
										},
									},
								},
							},
						}
						Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
					})

					It("escaped monospace text with nested bold text", func() {
						source := `\` + "`monospace *and bold* content`"
						expected := types.DraftDocument{
							Blocks: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "`monospace "},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{Content: "and bold"},
												},
											},
											types.StringElement{Content: " content`"},
										},
									},
								},
							},
						}
						Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
					})
				})
			})

			Context("prevented subscript text substitution", func() {

				Context("without nested quoted text", func() {

					It("escaped subscript text with single quote", func() {
						source := `\~subscriptcontent~`
						expected := types.DraftDocument{
							Blocks: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "~subscriptcontent~"},
										},
									},
								},
							},
						}
						Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
					})

					It("escaped subscript text with single quote and more backslashes", func() {
						source := `\\~subscriptcontent~`
						expected := types.DraftDocument{
							Blocks: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: `\~subscriptcontent~`}, // only 1 backslash removed
										},
									},
								},
							},
						}
						Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
					})

				})

				Context("with nested quoted text", func() {

					It("escaped subscript text with nested bold text", func() {
						source := `\~*boldcontent*~`
						expected := types.DraftDocument{
							Blocks: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "~"},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{Content: "boldcontent"},
												},
											},
											types.StringElement{Content: "~"},
										},
									},
								},
							},
						}
						Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
					})

					It("escaped subscript text with nested bold text", func() {
						source := `\~subscript *and bold* content~`
						expected := types.DraftDocument{
							Blocks: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: `\~subscript `},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{Content: "and bold"},
												},
											},
											types.StringElement{Content: " content~"},
										},
									},
								},
							},
						}
						Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
					})
				})
			})

			Context("prevented superscript text substitution", func() {

				Context("without nested quoted text", func() {

					It("escaped superscript text with single quote", func() {
						source := `\^superscriptcontent^`
						expected := types.DraftDocument{
							Blocks: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "^superscriptcontent^"},
										},
									},
								},
							},
						}
						Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
					})

					It("escaped superscript text with single quote and more backslashes", func() {
						source := `\\^superscriptcontent^`
						expected := types.DraftDocument{
							Blocks: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: `\^superscriptcontent^`}, // only 1 backslash removed
										},
									},
								},
							},
						}
						Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
					})

				})

				Context("with nested quoted text", func() {

					It("escaped superscript text with nested bold text - case 1", func() {
						source := `\^*bold content*^` // valid escaped superscript since it has no space within
						expected := types.DraftDocument{
							Blocks: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: `^`},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{Content: "bold content"},
												},
											},
											types.StringElement{Content: "^"},
										},
									},
								},
							},
						}
						Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
					})

					It("escaped superscript text with unbalanced double backquote and nested bold test", func() {
						source := `\^*bold content*^`
						expected := types.DraftDocument{
							Blocks: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "^"},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{Content: "bold content"},
												},
											},
											types.StringElement{Content: "^"},
										},
									},
								},
							},
						}
						Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
					})

					It("escaped superscript text with nested bold text - case 2", func() {
						source := `\^superscript *and bold* content^` // invalid superscript text since it has spaces within
						expected := types.DraftDocument{
							Blocks: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: `\^superscript `},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{Content: "and bold"},
												},
											},
											types.StringElement{Content: " content^"},
										},
									},
								},
							},
						}
						Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
					})
				})
			})
		})
	})

	Context("final document", func() {
		Context("quoted text with single punctuation", func() {

			It("bold text with 1 word", func() {
				source := "*hello*"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: "hello"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("bold text with 2 words", func() {
				source := "*bold    content*"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: "bold    content"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("bold text with 3 words", func() {
				source := "*some bold content*"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
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
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("italic text with 3 words in single quote", func() {
				source := "_some italic content_"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Italic,
										Elements: []interface{}{
											types.StringElement{Content: "some italic content"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("monospace text with 3 words", func() {
				source := "`some monospace content`"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "some monospace content"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("invalid subscript text with 3 words", func() {
				source := "~some subscript content~"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "~some subscript content~"},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("invalid superscript text with 3 words", func() {
				source := "^some superscript content^"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "^some superscript content^"},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("bold text within italic text", func() {
				source := "_some *bold* content_"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Italic,
										Elements: []interface{}{
											types.StringElement{Content: "some "},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{Content: "bold"},
												},
											},
											types.StringElement{Content: " content"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("monospace text within bold text within italic quote", func() {
				source := "*some _italic and `monospaced content`_*"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: "some "},
											types.QuotedText{
												Kind: types.Italic,
												Elements: []interface{}{
													types.StringElement{Content: "italic and "},
													types.QuotedText{
														Kind: types.Monospace,
														Elements: []interface{}{
															types.StringElement{Content: "monospaced content"},
														},
													},
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

			It("italic text within italic text", func() {
				source := "_some _very italic_ content_"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Italic,
										Elements: []interface{}{
											types.StringElement{Content: "some _very italic"},
										},
									},
									types.StringElement{Content: " content_"},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("subscript text attached", func() {
				source := "O~2~ is a molecule"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "O"},
									types.QuotedText{
										Kind: types.Subscript,
										Elements: []interface{}{
											types.StringElement{Content: "2"},
										},
									},
									types.StringElement{Content: " is a molecule"},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("superscript text attached", func() {
				source := "M^me^ White"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "M"},
									types.QuotedText{
										Kind: types.Superscript,
										Elements: []interface{}{
											types.StringElement{Content: "me"},
										},
									},
									types.StringElement{Content: " White"},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("invalid subscript text with 3 words", func() {
				source := "~some subscript content~"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "~some subscript content~"},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

		})

		Context("Quoted text with double punctuation", func() {

			It("bold text of 1 word in double quote", func() {
				source := "**hello**"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: "hello"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("italic text with 3 words in double quote", func() {
				source := "__some italic content__"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Italic,
										Elements: []interface{}{
											types.StringElement{Content: "some italic content"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("monospace text with 3 words in double quote", func() {
				source := "``some monospace content``"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "some monospace content"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("superscript text within italic text", func() {
				source := "__some ^superscript^ content__"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Italic,
										Elements: []interface{}{
											types.StringElement{Content: "some "},
											types.QuotedText{
												Kind: types.Superscript,
												Elements: []interface{}{
													types.StringElement{Content: "superscript"},
												},
											},
											types.StringElement{Content: " content"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("superscript text within italic text within bold quote", func() {
				source := "**some _italic and ^superscriptcontent^_**"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: "some "},
											types.QuotedText{
												Kind: types.Italic,
												Elements: []interface{}{
													types.StringElement{Content: "italic and "},
													types.QuotedText{
														Kind: types.Superscript,
														Elements: []interface{}{
															types.StringElement{Content: "superscriptcontent"},
														},
													},
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

		Context("Quoted text inline", func() {

			It("inline content with bold text", func() {
				source := "a paragraph with *some bold content*"
				expected := types.Document{
					Elements: []interface{}{
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
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("inline content with invalid bold text - use case 1", func() {
				source := "a paragraph with *some bold content"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a paragraph with *some bold content"},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("inline content with invalid bold text - use case 2", func() {
				source := "a paragraph with *some bold content *"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a paragraph with *some bold content *"},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("inline content with invalid bold text - use case 3", func() {
				source := "a paragraph with * some bold content*"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a paragraph with * some bold content*"},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("invalid italic text within bold text", func() {
				source := "some *bold and _italic content _ together*."
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{

									types.StringElement{Content: "some "},
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: "bold and _italic content _ together"},
										},
									},
									types.StringElement{Content: "."},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("italic text within invalid bold text", func() {
				source := "some *bold and _italic content_ together *."
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "some *bold and "},
									types.QuotedText{
										Kind: types.Italic,
										Elements: []interface{}{
											types.StringElement{Content: "italic content"},
										},
									},
									types.StringElement{Content: " together *."},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("inline content with invalid subscript text - use case 1", func() {
				source := "a paragraph with ~some subscript content"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a paragraph with ~some subscript content"},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("inline content with invalid subscript text - use case 2", func() {
				source := "a paragraph with ~some subscript content ~"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a paragraph with ~some subscript content ~"},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("inline content with invalid subscript text - use case 3", func() {
				source := "a paragraph with ~ some subscript content~"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a paragraph with ~ some subscript content~"},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("inline content with invalid superscript text - use case 1", func() {
				source := "a paragraph with ^some superscript content"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{

									types.StringElement{Content: "a paragraph with ^some superscript content"},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("inline content with invalid superscript text - use case 2", func() {
				source := "a paragraph with ^some superscript content ^"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{

									types.StringElement{Content: "a paragraph with ^some superscript content ^"},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("inline content with invalid superscript text - use case 3", func() {
				source := "a paragraph with ^ some superscript content^"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{

									types.StringElement{Content: "a paragraph with ^ some superscript content^"},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})

		Context("nested quoted text", func() {

			It("italic text within bold text", func() {
				source := "some *bold and _italic content_ together*."
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "some "},
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: "bold and "},
											types.QuotedText{
												Kind: types.Italic,
												Elements: []interface{}{
													types.StringElement{Content: "italic content"},
												},
											},
											types.StringElement{Content: " together"},
										},
									},
									types.StringElement{Content: "."},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("single-quote bold within single-quote bold text", func() {
				// here we don't allow for bold text within bold text, to comply with the existing implementations (asciidoc and asciidoctor)
				source := "*some *nested bold* content*"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: "some *nested bold"},
										},
									},
									types.StringElement{Content: " content*"},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("double-quote bold within double-quote bold text", func() {
				// here we don't allow for bold text within bold text, to comply with the existing implementations (asciidoc and asciidoctor)
				source := "**some **nested bold** content**"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: "some "},
										},
									},
									types.StringElement{Content: "nested bold"},
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: " content"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("single-quote bold within double-quote bold text", func() {
				// here we don't allow for bold text within bold text, to comply with the existing implementations (asciidoc and asciidoctor)
				source := "**some *nested bold* content**"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: "some "},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{Content: "nested bold"},
												},
											},
											types.StringElement{Content: " content"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("double-quote bold within single-quote bold text", func() {
				// here we don't allow for bold text within bold text, to comply with the existing implementations (asciidoc and asciidoctor)
				source := "*some **nested bold** content*"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: "some "},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{Content: "nested bold"},
												},
											},
											types.StringElement{Content: " content"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("single-quote italic within single-quote italic text", func() {
				// here we don't allow for italic text within italic text, to comply with the existing implementations (asciidoc and asciidoctor)
				source := "_some _nested italic_ content_"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Italic,
										Elements: []interface{}{
											types.StringElement{Content: "some _nested italic"},
										},
									},
									types.StringElement{Content: " content_"},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("double-quote italic within double-quote italic text", func() {
				// here we don't allow for italic text within italic text, to comply with the existing implementations (asciidoc and asciidoctor)
				source := "__some __nested italic__ content__"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Italic,
										Elements: []interface{}{
											types.StringElement{Content: "some "},
										},
									},
									types.StringElement{Content: "nested italic"},
									types.QuotedText{
										Kind: types.Italic,
										Elements: []interface{}{
											types.StringElement{Content: " content"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("single-quote italic within double-quote italic text", func() {
				// here we allow for italic text within italic text, to comply with the existing implementations (asciidoc and asciidoctor)
				source := "_some __nested italic__ content_"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Italic,
										Elements: []interface{}{
											types.StringElement{Content: "some "},
											types.QuotedText{
												Kind: types.Italic,
												Elements: []interface{}{
													types.StringElement{Content: "nested italic"},
												},
											},
											types.StringElement{Content: " content"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("double-quote italic within single-quote italic text", func() {
				// here we allow for italic text within italic text, to comply with the existing implementations (asciidoc and asciidoctor)
				source := "_some __nested italic__ content_"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Italic,
										Elements: []interface{}{
											types.StringElement{Content: "some "},
											types.QuotedText{
												Kind: types.Italic,
												Elements: []interface{}{
													types.StringElement{Content: "nested italic"},
												},
											},
											types.StringElement{Content: " content"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("single-quote monospace within single-quote monospace text", func() {
				// here we don't allow for monospace text within monospace text, to comply with the existing implementations (asciidoc and asciidoctor)
				source := "`some `nested monospace` content`"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "some `nested monospace"},
										},
									},
									types.StringElement{Content: " content`"},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("double-quote monospace within double-quote monospace text", func() {
				// here we don't allow for monospace text within monospace text, to comply with the existing implementations (asciidoc and asciidoctor)
				source := "``some ``nested monospace`` content``"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "some "},
										},
									},
									types.StringElement{Content: "nested monospace"},
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: " content"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("single-quote monospace within double-quote monospace text", func() {
				// here we allow for monospace text within monospace text, to comply with the existing implementations (asciidoc and asciidoctor)
				source := "`some ``nested monospace`` content`"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "some "},
											types.QuotedText{
												Kind: types.Monospace,
												Elements: []interface{}{
													types.StringElement{Content: "nested monospace"},
												},
											},
											types.StringElement{Content: " content"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("double-quote monospace within single-quote monospace text", func() {
				// here we allow for monospace text within monospace text, to comply with the existing implementations (asciidoc and asciidoctor)
				source := "`some ``nested monospace`` content`"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "some "},
											types.QuotedText{
												Kind: types.Monospace,
												Elements: []interface{}{
													types.StringElement{Content: "nested monospace"},
												},
											},
											types.StringElement{Content: " content"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("unbalanced bold in monospace - case 1", func() {
				source := "`*a`"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "*a"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("unbalanced bold in monospace - case 2", func() {
				source := "`a*b`"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "a*b"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("italic in monospace", func() {
				source := "`_a_`"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.QuotedText{
												Kind: types.Italic,
												Elements: []interface{}{
													types.StringElement{Content: "a"},
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

			It("unbalanced italic in monospace", func() {
				source := "`a_b`"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "a_b"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("unparsed bold in monospace", func() {
				source := "`a*b*`"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "a*b*"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("parsed subscript in monospace", func() {
				source := "`a~b~`"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "a"},
											types.QuotedText{
												Kind: types.Subscript,
												Elements: []interface{}{
													types.StringElement{Content: "b"},
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

			It("multiline in single quoted monospace - case 1", func() {
				source := "`a\nb`"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "a\nb"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("multiline in double quoted monospace - case 1", func() {
				source := "`a\nb`"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "a\nb"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("multiline in single quoted  monospace - case 2", func() {
				source := "`a\n*b*`"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "a\n"},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{Content: "b"},
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

			It("multiline in double quoted  monospace - case 2", func() {
				source := "`a\n*b*`"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "a\n"},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{Content: "b"},
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

			It("link in bold", func() {
				source := "*a link:/[b]*"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: "a "},
											types.InlineLink{
												Attributes: types.Attributes{
													"positional-1": []interface{}{
														types.StringElement{
															Content: "b",
														},
													},
												},
												Location: types.Location{
													Path: []interface{}{
														types.StringElement{
															Content: "/",
														},
													},
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

			It("image in bold", func() {
				source := "*a image:foo.png[]*"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: "a "},
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
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("singleplus passthrough in bold", func() {
				source := "*a +image:foo.png[]+*"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: "a "},
											types.InlinePassthrough{
												Kind: types.SinglePlusPassthrough,
												Elements: []interface{}{
													types.StringElement{Content: "image:foo.png[]"},
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

			It("tripleplus passthrough in bold", func() {
				source := "*a +++image:foo.png[]+++*"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: "a "},
											types.InlinePassthrough{
												Kind: types.TriplePlusPassthrough,
												Elements: []interface{}{
													types.StringElement{Content: "image:foo.png[]"},
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

			It("link in italic", func() {
				source := "_a link:/[b]_"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Italic,
										Elements: []interface{}{
											types.StringElement{Content: "a "},
											types.InlineLink{
												Attributes: types.Attributes{
													"positional-1": []interface{}{
														types.StringElement{
															Content: "b",
														},
													},
												},
												Location: types.Location{
													Path: []interface{}{
														types.StringElement{
															Content: "/",
														},
													},
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

			It("image in italic", func() {
				source := "_a image:foo.png[]_"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Italic,
										Elements: []interface{}{
											types.StringElement{Content: "a "},
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
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("singleplus passthrough in italic", func() {
				source := "_a +image:foo.png[]+_"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Italic,
										Elements: []interface{}{
											types.StringElement{Content: "a "},
											types.InlinePassthrough{
												Kind: types.SinglePlusPassthrough,
												Elements: []interface{}{
													types.StringElement{Content: "image:foo.png[]"},
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

			It("tripleplus passthrough in italic", func() {
				source := "_a +++image:foo.png[]+++_"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Italic,
										Elements: []interface{}{
											types.StringElement{Content: "a "},
											types.InlinePassthrough{
												Kind: types.TriplePlusPassthrough,
												Elements: []interface{}{
													types.StringElement{Content: "image:foo.png[]"},
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

			It("link in monospace", func() {
				source := "`a link:/[b]`"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "a "},
											types.InlineLink{
												Attributes: types.Attributes{
													"positional-1": []interface{}{
														types.StringElement{
															Content: "b",
														},
													},
												},
												Location: types.Location{
													Path: []interface{}{
														types.StringElement{
															Content: "/",
														},
													},
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

			It("image in monospace", func() {
				source := "`a image:foo.png[]`"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "a "},
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
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("singleplus passthrough in monospace", func() {
				source := "`a +image:foo.png[]+`"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{

									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "a "},
											types.InlinePassthrough{
												Kind: types.SinglePlusPassthrough,
												Elements: []interface{}{
													types.StringElement{Content: "image:foo.png[]"},
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

			It("tripleplus passthrough in monospace", func() {
				source := "`a +++image:foo.png[]+++`"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{

									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "a "},
											types.InlinePassthrough{
												Kind: types.TriplePlusPassthrough,
												Elements: []interface{}{
													types.StringElement{Content: "image:foo.png[]"},
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

		Context("unbalanced quoted text", func() {

			Context("unbalanced bold text", func() {

				It("unbalanced bold text - extra on left", func() {
					source := "**some bold content*"
					expected := types.Document{
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.QuotedText{
											Kind: types.Bold,
											Elements: []interface{}{
												types.StringElement{Content: "*some bold content"},
											},
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("unbalanced bold text - extra on right", func() {
					source := "*some bold content**"
					expected := types.Document{
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{

										types.QuotedText{
											Kind: types.Bold,
											Elements: []interface{}{
												types.StringElement{Content: "some bold content"},
											},
										},
										types.StringElement{Content: "*"},
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})
			})

			Context("unbalanced italic text", func() {

				It("unbalanced italic text - extra on left", func() {
					source := "__some italic content_"
					expected := types.Document{
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{

										types.QuotedText{
											Kind: types.Italic,
											Elements: []interface{}{
												types.StringElement{Content: "_some italic content"},
											},
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("unbalanced italic text - extra on right", func() {
					source := "_some italic content__"
					expected := types.Document{
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.QuotedText{
											Kind: types.Italic,
											Elements: []interface{}{
												types.StringElement{Content: "some italic content"},
											},
										},
										types.StringElement{Content: "_"},
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})
			})

			Context("unbalanced monospace text", func() {

				It("unbalanced monospace text - extra on left", func() {
					source := "``some monospace content`"
					expected := types.Document{
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{

										types.QuotedText{
											Kind: types.Monospace,
											Elements: []interface{}{
												types.StringElement{Content: "`some monospace content"},
											},
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("unbalanced monospace text - extra on right", func() {
					source := "`some monospace content``"
					expected := types.Document{
						Elements: []interface{}{
							types.Paragraph{
								Lines: [][]interface{}{
									{
										types.QuotedText{
											Kind: types.Monospace,
											Elements: []interface{}{
												types.StringElement{Content: "some monospace content"},
											},
										},
										types.StringElement{Content: "`"},
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})
			})

			It("inline content with unbalanced bold text", func() {
				source := "a paragraph with *some bold content"
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a paragraph with *some bold content"},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

		})

		Context("prevented substitution", func() {

			Context("prevented bold text substitution", func() {

				Context("without nested quoted text", func() {

					It("escaped bold text with single backslash", func() {
						source := `\*bold content*`
						expected := types.Document{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "*bold content*"},
										},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped bold text with multiple backslashes", func() {
						source := `\\*bold content*`
						expected := types.Document{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: `\*bold content*`},
										},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped bold text with double quote", func() {
						source := `\\**bold content**`
						expected := types.Document{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: `**bold content**`},
										},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped bold text with double quote and more backslashes", func() {
						source := `\\\**bold content**`
						expected := types.Document{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: `\**bold content**`},
										},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped bold text with unbalanced double quote", func() {
						source := `\**bold content*`
						expected := types.Document{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: `**bold content*`},
										},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped bold text with unbalanced double quote and more backslashes", func() {
						source := `\\\**bold content*`
						expected := types.Document{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: `\\**bold content*`},
										},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})
				})

				Context("with nested quoted text", func() {

					It("escaped bold text with nested italic text", func() {
						source := `\*_italic content_*`
						expected := types.Document{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "*"},
											types.QuotedText{
												Kind: types.Italic,
												Elements: []interface{}{
													types.StringElement{Content: "italic content"},
												},
											},
											types.StringElement{Content: "*"},
										},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped bold text with unbalanced double quote and nested italic test", func() {
						source := `\**_italic content_*`
						expected := types.Document{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "**"},
											types.QuotedText{
												Kind: types.Italic,
												Elements: []interface{}{
													types.StringElement{Content: "italic content"},
												},
											},
											types.StringElement{Content: "*"},
										},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped bold text with nested italic", func() {
						source := `\*bold _and italic_ content*`
						expected := types.Document{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "*bold "},
											types.QuotedText{
												Kind: types.Italic,
												Elements: []interface{}{
													types.StringElement{Content: "and italic"},
												},
											},
											types.StringElement{Content: " content*"},
										},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})
				})

			})

			Context("prevented italic text substitution", func() {

				Context("without nested quoted text", func() {

					It("escaped italic text with single quote", func() {
						source := `\_italic content_`
						expected := types.Document{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "_italic content_"},
										},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped italic text with single quote and more backslashes", func() {
						source := `\\_italic content_`
						expected := types.Document{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: `\_italic content_`},
										},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped italic text with double quote with 2 backslashes", func() {
						source := `\\__italic content__`
						expected := types.Document{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: `__italic content__`},
										},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped italic text with double quote with 3 backslashes", func() {
						source := `\\\__italic content__`
						expected := types.Document{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: `\__italic content__`},
										},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped italic text with unbalanced double quote", func() {
						source := `\__italic content_`
						expected := types.Document{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: `__italic content_`},
										},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped italic text with unbalanced double quote and more backslashes", func() {
						source := `\\\__italic content_`
						expected := types.Document{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: `\\__italic content_`}, // only 1 backslash remove
										},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})
				})

				Context("with nested quoted text", func() {

					It("escaped italic text with nested monospace text", func() {
						source := `\` + "_`monospace content`_" // gives: \_`monospace content`_
						expected := types.Document{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "_"},
											types.QuotedText{
												Kind: types.Monospace,
												Elements: []interface{}{
													types.StringElement{Content: "monospace content"},
												},
											},
											types.StringElement{Content: "_"},
										},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped italic text with unbalanced double quote and nested bold test", func() {
						source := `\__*bold content*_`
						expected := types.Document{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "__"},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{Content: "bold content"},
												},
											},
											types.StringElement{Content: "_"},
										},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped italic text with nested bold text", func() {
						source := `\_italic *and bold* content_`
						expected := types.Document{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "_italic "},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{Content: "and bold"},
												},
											},
											types.StringElement{Content: " content_"},
										},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})
				})
			})

			Context("prevented monospace text substitution", func() {

				Context("without nested quoted text", func() {

					It("escaped monospace text with single quote", func() {
						source := `\` + "`monospace content`"
						expected := types.Document{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "`monospace content`"}, // backslash removed
										},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped monospace text with single quote and more backslashes", func() {
						source := `\\` + "`monospace content`"
						expected := types.Document{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: `\` + "`monospace content`"}, // only 1 backslash removed
										},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped monospace text with double quote", func() {
						source := `\\` + "`monospace content``"
						expected := types.Document{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: `\` + "`monospace content``"}, // 2 back slashes "consumed"
										},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped monospace text with double quote and more backslashes", func() {
						source := `\\\` + "``monospace content``" // 3 backslashes
						expected := types.Document{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: `\` + "``monospace content``"}, // 2 back slashes "consumed"
										},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped monospace text with unbalanced double quote", func() {
						source := `\` + "``monospace content`"
						expected := types.Document{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "``monospace content`"},
										},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped monospace text with unbalanced double quote and more backslashes", func() {
						source := `\\\` + "``monospace content`" // 3 backslashes
						expected := types.Document{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: `\\` + "``monospace content`"}, // 2 backslashes removed
										},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})
				})

				Context("with nested quoted text", func() {

					It("escaped monospace text with nested bold text", func() {
						source := `\` + "`*bold content*`" // gives: \`*bold content*`
						expected := types.Document{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "`"},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{Content: "bold content"},
												},
											},
											types.StringElement{Content: "`"},
										},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped monospace text with unbalanced double backquote and nested bold test", func() {
						source := `\` + "``*bold content*`"
						expected := types.Document{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "``"},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{Content: "bold content"},
												},
											},
											types.StringElement{Content: "`"},
										},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped monospace text with nested bold text", func() {
						source := `\` + "`monospace *and bold* content`"
						expected := types.Document{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "`monospace "},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{Content: "and bold"},
												},
											},
											types.StringElement{Content: " content`"},
										},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})
				})
			})

			Context("prevented subscript text substitution", func() {

				Context("without nested quoted text", func() {

					It("escaped subscript text with single quote", func() {
						source := `\~subscriptcontent~`
						expected := types.Document{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "~subscriptcontent~"},
										},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped subscript text with single quote and more backslashes", func() {
						source := `\\~subscriptcontent~`
						expected := types.Document{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: `\~subscriptcontent~`}, // only 1 backslash removed
										},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

				})

				Context("with nested quoted text", func() {

					It("escaped subscript text with nested bold text", func() {
						source := `\~*boldcontent*~`
						expected := types.Document{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "~"},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{Content: "boldcontent"},
												},
											},
											types.StringElement{Content: "~"},
										},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped subscript text with nested bold text", func() {
						source := `\~subscript *and bold* content~`
						expected := types.Document{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: `\~subscript `},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{Content: "and bold"},
												},
											},
											types.StringElement{Content: " content~"},
										},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})
				})
			})

			Context("prevented superscript text substitution", func() {

				Context("without nested quoted text", func() {

					It("escaped superscript text with single quote", func() {
						source := `\^superscriptcontent^`
						expected := types.Document{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "^superscriptcontent^"},
										},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped superscript text with single quote and more backslashes", func() {
						source := `\\^superscriptcontent^`
						expected := types.Document{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: `\^superscriptcontent^`}, // only 1 backslash removed
										},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

				})

				Context("with nested quoted text", func() {

					It("escaped superscript text with nested bold text - case 1", func() {
						source := `\^*bold content*^` // valid escaped superscript since it has no space within
						expected := types.Document{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: `^`},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{Content: "bold content"},
												},
											},
											types.StringElement{Content: "^"},
										},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped superscript text with unbalanced double backquote and nested bold test", func() {
						source := `\^*bold content*^`
						expected := types.Document{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "^"},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{Content: "bold content"},
												},
											},
											types.StringElement{Content: "^"},
										},
									},
								},
							},
						}
						Expect(ParseDocument(source)).To(MatchDocument(expected))
					})

					It("escaped superscript text with nested bold text - case 2", func() {
						source := `\^superscript *and bold* content^` // invalid superscript text since it has spaces within
						expected := types.Document{
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: `\^superscript `},
											types.QuotedText{
												Kind: types.Bold,
												Elements: []interface{}{
													types.StringElement{Content: "and bold"},
												},
											},
											types.StringElement{Content: " content^"},
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
	})
})

var _ = Describe("quoted texts - final document", func() {

	It("image in bold", func() {
		source := "*a image:foo.png[]*"
		expected := types.Document{
			Elements: []interface{}{
				types.Paragraph{
					Lines: [][]interface{}{
						{
							types.QuotedText{
								Kind: types.Bold,
								Elements: []interface{}{
									types.StringElement{Content: "a "},
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
				},
			},
		}
		Expect(ParseDocument(source)).To(MatchDocument(expected))
	})

	It("image in italic", func() {
		source := "_a image:foo.png[]_"
		expected := types.Document{
			Elements: []interface{}{
				types.Paragraph{
					Lines: [][]interface{}{
						{
							types.QuotedText{
								Kind: types.Italic,
								Elements: []interface{}{
									types.StringElement{Content: "a "},
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
				},
			},
		}
		Expect(ParseDocument(source)).To(MatchDocument(expected))
	})

	It("image in monospace", func() {
		source := "`a image:foo.png[]`"
		expected := types.Document{
			Elements: []interface{}{
				types.Paragraph{
					Lines: [][]interface{}{
						{
							types.QuotedText{
								Kind: types.Monospace,
								Elements: []interface{}{
									types.StringElement{Content: "a "},
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
				},
			},
		}
		Expect(ParseDocument(source)).To(MatchDocument(expected))
	})
})
