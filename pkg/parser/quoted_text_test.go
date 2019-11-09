package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("quoted texts - draft", func() {

	Context("quoted text with single punctuation", func() {

		It("bold text with 1 word", func() {
			source := "*hello*"
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.QuotedText{
									Kind: types.Bold,
									Elements: types.InlineElements{
										types.StringElement{Content: "hello"},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})

		It("bold text with 2 words", func() {
			source := "*bold    content*"
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.QuotedText{
									Kind: types.Bold,
									Elements: types.InlineElements{
										types.StringElement{Content: "bold    content"},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})

		It("bold text with 3 words", func() {
			source := "*some bold content*"
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.QuotedText{
									Kind: types.Bold,
									Elements: types.InlineElements{
										types.StringElement{Content: "some bold content"},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})

		It("italic text with 3 words in single quote", func() {
			source := "_some italic content_"
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.QuotedText{
									Kind: types.Italic,
									Elements: types.InlineElements{
										types.StringElement{Content: "some italic content"},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})

		It("monospace text with 3 words", func() {
			source := "`some monospace content`"
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.QuotedText{
									Kind: types.Monospace,
									Elements: types.InlineElements{
										types.StringElement{Content: "some monospace content"},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})

		It("invalid subscript text with 3 words", func() {
			source := "~some subscript content~"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "~some subscript content~"},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("invalid superscript text with 3 words", func() {
			source := "^some superscript content^"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "^some superscript content^"},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("bold text within italic text", func() {
			source := "_some *bold* content_"
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.QuotedText{
									Kind: types.Italic,
									Elements: types.InlineElements{
										types.StringElement{Content: "some "},
										types.QuotedText{
											Kind: types.Bold,
											Elements: types.InlineElements{
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
			Expect(source).To(BecomeDraftDocument(expected))
		})

		It("monospace text within bold text within italic quote", func() {
			source := "*some _italic and `monospaced content`_*"
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.QuotedText{
									Kind: types.Bold,
									Elements: types.InlineElements{
										types.StringElement{Content: "some "},
										types.QuotedText{
											Kind: types.Italic,
											Elements: types.InlineElements{
												types.StringElement{Content: "italic and "},
												types.QuotedText{
													Kind: types.Monospace,
													Elements: types.InlineElements{
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
			Expect(source).To(BecomeDraftDocument(expected))
		})

		It("italic text within italic text", func() {
			source := "_some _very italic_ content_"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.QuotedText{
							Kind: types.Italic,
							Elements: types.InlineElements{
								types.StringElement{Content: "some _very italic"},
							},
						},
						types.StringElement{Content: " content_"},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("subscript text attached", func() {
			source := "O~2~ is a molecule"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "O"},
						types.QuotedText{
							Kind: types.Subscript,
							Elements: types.InlineElements{
								types.StringElement{Content: "2"},
							},
						},
						types.StringElement{Content: " is a molecule"},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("superscript text attached", func() {
			source := "M^me^ White"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "M"},
						types.QuotedText{
							Kind: types.Superscript,
							Elements: types.InlineElements{
								types.StringElement{Content: "me"},
							},
						},
						types.StringElement{Content: " White"},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("invalid subscript text with 3 words", func() {
			source := "~some subscript content~"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "~some subscript content~"},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

	})

	Context("Quoted text with double punctuation", func() {

		It("bold text of 1 word in double quote", func() {
			source := "**hello**"
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.QuotedText{
									Kind: types.Bold,
									Elements: types.InlineElements{
										types.StringElement{Content: "hello"},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})

		It("italic text with 3 words in double quote", func() {
			source := "__some italic content__"
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.QuotedText{
									Kind: types.Italic,
									Elements: types.InlineElements{
										types.StringElement{Content: "some italic content"},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})

		It("monospace text with 3 words in double quote", func() {
			source := "``some monospace content``"
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.QuotedText{
									Kind: types.Monospace,
									Elements: types.InlineElements{
										types.StringElement{Content: "some monospace content"},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})

		It("superscript text within italic text", func() {
			source := "__some ^superscript^ content__"
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.QuotedText{
									Kind: types.Italic,
									Elements: types.InlineElements{
										types.StringElement{Content: "some "},
										types.QuotedText{
											Kind: types.Superscript,
											Elements: types.InlineElements{
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
			Expect(source).To(BecomeDraftDocument(expected))
		})

		It("superscript text within italic text within bold quote", func() {
			source := "**some _italic and ^superscriptcontent^_**"
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.QuotedText{
									Kind: types.Bold,
									Elements: types.InlineElements{
										types.StringElement{Content: "some "},
										types.QuotedText{
											Kind: types.Italic,
											Elements: types.InlineElements{
												types.StringElement{Content: "italic and "},
												types.QuotedText{
													Kind: types.Superscript,
													Elements: types.InlineElements{
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
			Expect(source).To(BecomeDraftDocument(expected))
		})
	})

	Context("Quoted text inline", func() {

		It("inline content with bold text", func() {
			source := "a paragraph with *some bold content*"
			expected := types.Paragraph{
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
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("inline content with invalid bold text - use case 1", func() {
			source := "a paragraph with *some bold content"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "a paragraph with *some bold content"},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("inline content with invalid bold text - use case 2", func() {
			source := "a paragraph with *some bold content *"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "a paragraph with *some bold content *"},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("inline content with invalid bold text - use case 3", func() {
			source := "a paragraph with * some bold content*"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "a paragraph with * some bold content*"},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("invalid italic text within bold text", func() {
			source := "some *bold and _italic content _ together*."
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{

						types.StringElement{Content: "some "},
						types.QuotedText{
							Kind: types.Bold,
							Elements: types.InlineElements{
								types.StringElement{Content: "bold and _italic content _ together"},
							},
						},
						types.StringElement{Content: "."},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("italic text within invalid bold text", func() {
			source := "some *bold and _italic content_ together *."
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "some *bold and "},
						types.QuotedText{
							Kind: types.Italic,
							Elements: types.InlineElements{
								types.StringElement{Content: "italic content"},
							},
						},
						types.StringElement{Content: " together *."},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("inline content with invalid subscript text - use case 1", func() {
			source := "a paragraph with ~some subscript content"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "a paragraph with ~some subscript content"},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("inline content with invalid subscript text - use case 2", func() {
			source := "a paragraph with ~some subscript content ~"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "a paragraph with ~some subscript content ~"},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("inline content with invalid subscript text - use case 3", func() {
			source := "a paragraph with ~ some subscript content~"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "a paragraph with ~ some subscript content~"},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("inline content with invalid superscript text - use case 1", func() {
			source := "a paragraph with ^some superscript content"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{

						types.StringElement{Content: "a paragraph with ^some superscript content"},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("inline content with invalid superscript text - use case 2", func() {
			source := "a paragraph with ^some superscript content ^"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{

						types.StringElement{Content: "a paragraph with ^some superscript content ^"},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("inline content with invalid superscript text - use case 3", func() {
			source := "a paragraph with ^ some superscript content^"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{

						types.StringElement{Content: "a paragraph with ^ some superscript content^"},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})
	})

	Context("nested quoted text", func() {

		It("italic text within bold text", func() {
			source := "some *bold and _italic content_ together*."
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "some "},
						types.QuotedText{
							Kind: types.Bold,
							Elements: types.InlineElements{
								types.StringElement{Content: "bold and "},
								types.QuotedText{
									Kind: types.Italic,
									Elements: types.InlineElements{
										types.StringElement{Content: "italic content"},
									},
								},
								types.StringElement{Content: " together"},
							},
						},
						types.StringElement{Content: "."},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("single-quote bold within single-quote bold text", func() {
			// here we don't allow for bold text within bold text, to comply with the existing implementations (asciidoc and asciidoctor)
			source := "*some *nested bold* content*"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.QuotedText{
							Kind: types.Bold,
							Elements: types.InlineElements{
								types.StringElement{Content: "some *nested bold"},
							},
						},
						types.StringElement{Content: " content*"},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("double-quote bold within double-quote bold text", func() {
			// here we don't allow for bold text within bold text, to comply with the existing implementations (asciidoc and asciidoctor)
			source := "**some **nested bold** content**"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.QuotedText{
							Kind: types.Bold,
							Elements: types.InlineElements{
								types.StringElement{Content: "some "},
							},
						},
						types.StringElement{Content: "nested bold"},
						types.QuotedText{
							Kind: types.Bold,
							Elements: types.InlineElements{
								types.StringElement{Content: " content"},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("single-quote bold within double-quote bold text", func() {
			// here we don't allow for bold text within bold text, to comply with the existing implementations (asciidoc and asciidoctor)
			source := "**some *nested bold* content**"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.QuotedText{
							Kind: types.Bold,
							Elements: types.InlineElements{
								types.StringElement{Content: "some "},
								types.QuotedText{
									Kind: types.Bold,
									Elements: types.InlineElements{
										types.StringElement{Content: "nested bold"},
									},
								},
								types.StringElement{Content: " content"},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("double-quote bold within single-quote bold text", func() {
			// here we don't allow for bold text within bold text, to comply with the existing implementations (asciidoc and asciidoctor)
			source := "*some **nested bold** content*"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.QuotedText{
							Kind: types.Bold,
							Elements: types.InlineElements{
								types.StringElement{Content: "some "},
								types.QuotedText{
									Kind: types.Bold,
									Elements: types.InlineElements{
										types.StringElement{Content: "nested bold"},
									},
								},
								types.StringElement{Content: " content"},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("single-quote italic within single-quote italic text", func() {
			// here we don't allow for italic text within italic text, to comply with the existing implementations (asciidoc and asciidoctor)
			source := "_some _nested italic_ content_"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.QuotedText{
							Kind: types.Italic,
							Elements: types.InlineElements{
								types.StringElement{Content: "some _nested italic"},
							},
						},
						types.StringElement{Content: " content_"},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("double-quote italic within double-quote italic text", func() {
			// here we don't allow for italic text within italic text, to comply with the existing implementations (asciidoc and asciidoctor)
			source := "__some __nested italic__ content__"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.QuotedText{
							Kind: types.Italic,
							Elements: types.InlineElements{
								types.StringElement{Content: "some "},
							},
						},
						types.StringElement{Content: "nested italic"},
						types.QuotedText{
							Kind: types.Italic,
							Elements: types.InlineElements{
								types.StringElement{Content: " content"},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("single-quote italic within double-quote italic text", func() {
			// here we allow for italic text within italic text, to comply with the existing implementations (asciidoc and asciidoctor)
			source := "_some __nested italic__ content_"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.QuotedText{
							Kind: types.Italic,
							Elements: types.InlineElements{
								types.StringElement{Content: "some "},
								types.QuotedText{
									Kind: types.Italic,
									Elements: types.InlineElements{
										types.StringElement{Content: "nested italic"},
									},
								},
								types.StringElement{Content: " content"},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("double-quote italic within single-quote italic text", func() {
			// here we allow for italic text within italic text, to comply with the existing implementations (asciidoc and asciidoctor)
			source := "_some __nested italic__ content_"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.QuotedText{
							Kind: types.Italic,
							Elements: types.InlineElements{
								types.StringElement{Content: "some "},
								types.QuotedText{
									Kind: types.Italic,
									Elements: types.InlineElements{
										types.StringElement{Content: "nested italic"},
									},
								},
								types.StringElement{Content: " content"},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("single-quote monospace within single-quote monospace text", func() {
			// here we don't allow for monospace text within monospace text, to comply with the existing implementations (asciidoc and asciidoctor)
			source := "`some `nested monospace` content`"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.QuotedText{
							Kind: types.Monospace,
							Elements: types.InlineElements{
								types.StringElement{Content: "some `nested monospace"},
							},
						},
						types.StringElement{Content: " content`"},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("double-quote monospace within double-quote monospace text", func() {
			// here we don't allow for monospace text within monospace text, to comply with the existing implementations (asciidoc and asciidoctor)
			source := "``some ``nested monospace`` content``"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.QuotedText{
							Kind: types.Monospace,
							Elements: types.InlineElements{
								types.StringElement{Content: "some "},
							},
						},
						types.StringElement{Content: "nested monospace"},
						types.QuotedText{
							Kind: types.Monospace,
							Elements: types.InlineElements{
								types.StringElement{Content: " content"},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("single-quote monospace within double-quote monospace text", func() {
			// here we allow for monospace text within monospace text, to comply with the existing implementations (asciidoc and asciidoctor)
			source := "`some ``nested monospace`` content`"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.QuotedText{
							Kind: types.Monospace,
							Elements: types.InlineElements{
								types.StringElement{Content: "some "},
								types.QuotedText{
									Kind: types.Monospace,
									Elements: types.InlineElements{
										types.StringElement{Content: "nested monospace"},
									},
								},
								types.StringElement{Content: " content"},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("double-quote monospace within single-quote monospace text", func() {
			// here we allow for monospace text within monospace text, to comply with the existing implementations (asciidoc and asciidoctor)
			source := "`some ``nested monospace`` content`"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.QuotedText{
							Kind: types.Monospace,
							Elements: types.InlineElements{
								types.StringElement{Content: "some "},
								types.QuotedText{
									Kind: types.Monospace,
									Elements: types.InlineElements{
										types.StringElement{Content: "nested monospace"},
									},
								},
								types.StringElement{Content: " content"},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("unbalanced bold in monospace - case 1", func() {
			source := "`*a`"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.QuotedText{
							Kind: types.Monospace,
							Elements: types.InlineElements{
								types.StringElement{Content: "*a"},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("unbalanced bold in monospace - case 2", func() {
			source := "`a*b`"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.QuotedText{
							Kind: types.Monospace,
							Elements: types.InlineElements{
								types.StringElement{Content: "a*b"},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("italic in monospace", func() {
			source := "`_a_`"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.QuotedText{
							Kind: types.Monospace,
							Elements: types.InlineElements{
								types.QuotedText{
									Kind: types.Italic,
									Elements: types.InlineElements{
										types.StringElement{Content: "a"},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("unbalanced italic in monospace", func() {
			source := "`a_b`"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.QuotedText{
							Kind: types.Monospace,
							Elements: types.InlineElements{
								types.StringElement{Content: "a_b"},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("unparsed bold in monospace", func() {
			source := "`a*b*`"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.QuotedText{
							Kind: types.Monospace,
							Elements: types.InlineElements{
								types.StringElement{Content: "a*b*"},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("parsed subscript in monospace", func() {
			source := "`a~b~`"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.QuotedText{
							Kind: types.Monospace,
							Elements: types.InlineElements{
								types.StringElement{Content: "a"},
								types.QuotedText{
									Kind: types.Subscript,
									Elements: types.InlineElements{
										types.StringElement{Content: "b"},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("multiline in single quoted monospace - case 1", func() {
			source := "`a\nb`"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.QuotedText{
							Kind: types.Monospace,
							Elements: types.InlineElements{
								types.StringElement{Content: "a\nb"},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("multiline in double quoted monospace - case 1", func() {
			source := "`a\nb`"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.QuotedText{
							Kind: types.Monospace,
							Elements: types.InlineElements{
								types.StringElement{Content: "a\nb"},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("multiline in single quoted  monospace - case 2", func() {
			source := "`a\n*b*`"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.QuotedText{
							Kind: types.Monospace,
							Elements: types.InlineElements{
								types.StringElement{Content: "a\n"},
								types.QuotedText{
									Kind: types.Bold,
									Elements: types.InlineElements{
										types.StringElement{Content: "b"},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("multiline in double quoted  monospace - case 2", func() {
			source := "`a\n*b*`"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.QuotedText{
							Kind: types.Monospace,
							Elements: types.InlineElements{
								types.StringElement{Content: "a\n"},
								types.QuotedText{
									Kind: types.Bold,
									Elements: types.InlineElements{
										types.StringElement{Content: "b"},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("link in bold", func() {
			source := "*a link:/[b]*"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.QuotedText{
							Kind: types.Bold,
							Elements: types.InlineElements{
								types.StringElement{Content: "a "},
								types.InlineLink{
									Attributes: types.ElementAttributes{
										types.AttrInlineLinkText: types.InlineElements{
											types.StringElement{
												Content: "b",
											},
										},
									},
									Location: types.Location{
										types.StringElement{
											Content: "/",
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("image in bold", func() {
			source := "*a image:foo.png[]*"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.QuotedText{
							Kind: types.Bold,
							Elements: types.InlineElements{
								types.StringElement{Content: "a "},
								types.InlineImage{
									Attributes: types.ElementAttributes{
										types.AttrImageAlt: "foo",
									},
									Path: "foo.png",
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("singleplus passthrough in bold", func() {
			source := "*a +image:foo.png[]+*"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.QuotedText{
							Kind: types.Bold,
							Elements: types.InlineElements{
								types.StringElement{Content: "a "},
								types.Passthrough{
									Kind: types.SinglePlusPassthrough,
									Elements: types.InlineElements{
										types.StringElement{Content: "image:foo.png[]"},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("tripleplus passthrough in bold", func() {
			source := "*a +++image:foo.png[]+++*"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.QuotedText{
							Kind: types.Bold,
							Elements: types.InlineElements{
								types.StringElement{Content: "a "},
								types.Passthrough{
									Kind: types.TriplePlusPassthrough,
									Elements: types.InlineElements{
										types.StringElement{Content: "image:foo.png[]"},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("link in italic", func() {
			source := "_a link:/[b]_"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.QuotedText{
							Kind: types.Italic,
							Elements: types.InlineElements{
								types.StringElement{Content: "a "},
								types.InlineLink{
									Attributes: types.ElementAttributes{
										types.AttrInlineLinkText: types.InlineElements{
											types.StringElement{
												Content: "b",
											},
										},
									},
									Location: types.Location{
										types.StringElement{
											Content: "/",
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("image in italic", func() {
			source := "_a image:foo.png[]_"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.QuotedText{
							Kind: types.Italic,
							Elements: types.InlineElements{
								types.StringElement{Content: "a "},
								types.InlineImage{
									Attributes: types.ElementAttributes{
										types.AttrImageAlt: "foo",
									},
									Path: "foo.png",
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("singleplus passthrough in italic", func() {
			source := "_a +image:foo.png[]+_"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.QuotedText{
							Kind: types.Italic,
							Elements: types.InlineElements{
								types.StringElement{Content: "a "},
								types.Passthrough{
									Kind: types.SinglePlusPassthrough,
									Elements: types.InlineElements{
										types.StringElement{Content: "image:foo.png[]"},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("tripleplus passthrough in italic", func() {
			source := "_a +++image:foo.png[]+++_"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.QuotedText{
							Kind: types.Italic,
							Elements: types.InlineElements{
								types.StringElement{Content: "a "},
								types.Passthrough{
									Kind: types.TriplePlusPassthrough,
									Elements: types.InlineElements{
										types.StringElement{Content: "image:foo.png[]"},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("link in monospace", func() {
			source := "`a link:/[b]`"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.QuotedText{
							Kind: types.Monospace,
							Elements: types.InlineElements{
								types.StringElement{Content: "a "},
								types.InlineLink{
									Attributes: types.ElementAttributes{
										types.AttrInlineLinkText: types.InlineElements{
											types.StringElement{
												Content: "b",
											},
										},
									},
									Location: types.Location{
										types.StringElement{
											Content: "/",
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("image in monospace", func() {
			source := "`a image:foo.png[]`"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.QuotedText{
							Kind: types.Monospace,
							Elements: types.InlineElements{
								types.StringElement{Content: "a "},
								types.InlineImage{
									Attributes: types.ElementAttributes{
										types.AttrImageAlt: "foo",
									},
									Path: "foo.png",
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("singleplus passthrough in monospace", func() {
			source := "`a +image:foo.png[]+`"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{

						types.QuotedText{
							Kind: types.Monospace,
							Elements: types.InlineElements{
								types.StringElement{Content: "a "},
								types.Passthrough{
									Kind: types.SinglePlusPassthrough,
									Elements: types.InlineElements{
										types.StringElement{Content: "image:foo.png[]"},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

		It("tripleplus passthrough in monospace", func() {
			source := "`a +++image:foo.png[]+++`"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{

						types.QuotedText{
							Kind: types.Monospace,
							Elements: types.InlineElements{
								types.StringElement{Content: "a "},
								types.Passthrough{
									Kind: types.TriplePlusPassthrough,
									Elements: types.InlineElements{
										types.StringElement{Content: "image:foo.png[]"},
									},
								},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

	})

	Context("unbalanced quoted text", func() {

		Context("unbalanced bold text", func() {

			It("unbalanced bold text - extra on left", func() {
				source := "**some bold content*"
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{

							types.QuotedText{
								Kind: types.Bold,
								Elements: types.InlineElements{
									types.StringElement{Content: "*some bold content"},
								},
							},
						},
					},
				}
				Expect(source).To(EqualDocumentBlock(expected))
			})

			It("unbalanced bold text - extra on right", func() {
				source := "*some bold content**"
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{

							types.QuotedText{
								Kind: types.Bold,
								Elements: types.InlineElements{
									types.StringElement{Content: "some bold content"},
								},
							},
							types.StringElement{Content: "*"},
						},
					},
				}
				Expect(source).To(EqualDocumentBlock(expected))
			})
		})

		Context("unbalanced italic text", func() {

			It("unbalanced italic text - extra on left", func() {
				source := "__some italic content_"
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{

							types.QuotedText{
								Kind: types.Italic,
								Elements: types.InlineElements{
									types.StringElement{Content: "_some italic content"},
								},
							},
						},
					},
				}
				Expect(source).To(EqualDocumentBlock(expected))
			})

			It("unbalanced italic text - extra on right", func() {
				source := "_some italic content__"
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.QuotedText{
								Kind: types.Italic,
								Elements: types.InlineElements{
									types.StringElement{Content: "some italic content"},
								},
							},
							types.StringElement{Content: "_"},
						},
					},
				}
				Expect(source).To(EqualDocumentBlock(expected))
			})
		})

		Context("unbalanced monospace text", func() {

			It("unbalanced monospace text - extra on left", func() {
				source := "``some monospace content`"
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{

							types.QuotedText{
								Kind: types.Monospace,
								Elements: types.InlineElements{
									types.StringElement{Content: "`some monospace content"},
								},
							},
						},
					},
				}
				Expect(source).To(EqualDocumentBlock(expected))
			})

			It("unbalanced monospace text - extra on right", func() {
				source := "`some monospace content``"
				expected := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.QuotedText{
								Kind: types.Monospace,
								Elements: types.InlineElements{
									types.StringElement{Content: "some monospace content"},
								},
							},
							types.StringElement{Content: "`"},
						},
					},
				}
				Expect(source).To(EqualDocumentBlock(expected))
			})
		})

		It("inline content with unbalanced bold text", func() {
			source := "a paragraph with *some bold content"
			expected := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "a paragraph with *some bold content"},
					},
				},
			}
			Expect(source).To(EqualDocumentBlock(expected))
		})

	})

	Context("prevented substitution", func() {

		Context("prevented bold text substitution", func() {

			Context("without nested quoted text", func() {

				It("escaped bold text with single backslash", func() {
					source := `\*bold content*`
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "*bold content*"},
							},
						},
					}
					Expect(source).To(EqualDocumentBlock(expected))
				})

				It("escaped bold text with multiple backslashes", func() {
					source := `\\*bold content*`
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `\*bold content*`},
							},
						},
					}
					Expect(source).To(EqualDocumentBlock(expected))
				})

				It("escaped bold text with double quote", func() {
					source := `\\**bold content**`
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `**bold content**`},
							},
						},
					}
					Expect(source).To(EqualDocumentBlock(expected))
				})

				It("escaped bold text with double quote and more backslashes", func() {
					source := `\\\**bold content**`
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `\**bold content**`},
							},
						},
					}
					Expect(source).To(EqualDocumentBlock(expected))
				})

				It("escaped bold text with unbalanced double quote", func() {
					source := `\**bold content*`
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `**bold content*`},
							},
						},
					}
					Expect(source).To(EqualDocumentBlock(expected))
				})

				It("escaped bold text with unbalanced double quote and more backslashes", func() {
					source := `\\\**bold content*`
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `\\**bold content*`},
							},
						},
					}
					Expect(source).To(EqualDocumentBlock(expected))
				})
			})

			Context("with nested quoted text", func() {

				It("escaped bold text with nested italic text", func() {
					source := `\*_italic content_*`
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "*"},
								types.QuotedText{
									Kind: types.Italic,
									Elements: types.InlineElements{
										types.StringElement{Content: "italic content"},
									},
								},
								types.StringElement{Content: "*"},
							},
						},
					}
					Expect(source).To(EqualDocumentBlock(expected))
				})

				It("escaped bold text with unbalanced double quote and nested italic test", func() {
					source := `\**_italic content_*`
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "**"},
								types.QuotedText{
									Kind: types.Italic,
									Elements: types.InlineElements{
										types.StringElement{Content: "italic content"},
									},
								},
								types.StringElement{Content: "*"},
							},
						},
					}
					Expect(source).To(EqualDocumentBlock(expected))
				})

				It("escaped bold text with nested italic", func() {
					source := `\*bold _and italic_ content*`
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "*bold "},
								types.QuotedText{
									Kind: types.Italic,
									Elements: types.InlineElements{
										types.StringElement{Content: "and italic"},
									},
								},
								types.StringElement{Content: " content*"},
							},
						},
					}
					Expect(source).To(EqualDocumentBlock(expected))
				})
			})

		})

		Context("prevented italic text substitution", func() {

			Context("without nested quoted text", func() {

				It("escaped italic text with single quote", func() {
					source := `\_italic content_`
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "_italic content_"},
							},
						},
					}
					Expect(source).To(EqualDocumentBlock(expected))
				})

				It("escaped italic text with single quote and more backslashes", func() {
					source := `\\_italic content_`
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `\_italic content_`},
							},
						},
					}
					Expect(source).To(EqualDocumentBlock(expected))
				})

				It("escaped italic text with double quote with 2 backslashes", func() {
					source := `\\__italic content__`
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `__italic content__`},
							},
						},
					}
					Expect(source).To(EqualDocumentBlock(expected))
				})

				It("escaped italic text with double quote with 3 backslashes", func() {
					source := `\\\__italic content__`
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `\__italic content__`},
							},
						},
					}
					Expect(source).To(EqualDocumentBlock(expected))
				})

				It("escaped italic text with unbalanced double quote", func() {
					source := `\__italic content_`
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `__italic content_`},
							},
						},
					}
					Expect(source).To(EqualDocumentBlock(expected))
				})

				It("escaped italic text with unbalanced double quote and more backslashes", func() {
					source := `\\\__italic content_`
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `\\__italic content_`}, // only 1 backslash remove
							},
						},
					}
					Expect(source).To(EqualDocumentBlock(expected))
				})
			})

			Context("with nested quoted text", func() {

				It("escaped italic text with nested monospace text", func() {
					source := `\` + "_`monospace content`_" // gives: \_`monospace content`_
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "_"},
								types.QuotedText{
									Kind: types.Monospace,
									Elements: types.InlineElements{
										types.StringElement{Content: "monospace content"},
									},
								},
								types.StringElement{Content: "_"},
							},
						},
					}
					Expect(source).To(EqualDocumentBlock(expected))
				})

				It("escaped italic text with unbalanced double quote and nested bold test", func() {
					source := `\__*bold content*_`
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "__"},
								types.QuotedText{
									Kind: types.Bold,
									Elements: types.InlineElements{
										types.StringElement{Content: "bold content"},
									},
								},
								types.StringElement{Content: "_"},
							},
						},
					}
					Expect(source).To(EqualDocumentBlock(expected))
				})

				It("escaped italic text with nested bold text", func() {
					source := `\_italic *and bold* content_`
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "_italic "},
								types.QuotedText{
									Kind: types.Bold,
									Elements: types.InlineElements{
										types.StringElement{Content: "and bold"},
									},
								},
								types.StringElement{Content: " content_"},
							},
						},
					}
					Expect(source).To(EqualDocumentBlock(expected))
				})
			})
		})

		Context("prevented monospace text substitution", func() {

			Context("without nested quoted text", func() {

				It("escaped monospace text with single quote", func() {
					source := `\` + "`monospace content`"
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "`monospace content`"}, // backslash removed
							},
						},
					}
					Expect(source).To(EqualDocumentBlock(expected))
				})

				It("escaped monospace text with single quote and more backslashes", func() {
					source := `\\` + "`monospace content`"
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `\` + "`monospace content`"}, // only 1 backslash removed
							},
						},
					}
					Expect(source).To(EqualDocumentBlock(expected))
				})

				It("escaped monospace text with double quote", func() {
					source := `\\` + "`monospace content``"
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `\` + "`monospace content``"}, // 2 back slashes "consumed"
							},
						},
					}
					Expect(source).To(EqualDocumentBlock(expected))
				})

				It("escaped monospace text with double quote and more backslashes", func() {
					source := `\\\` + "``monospace content``" // 3 backslashes
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `\` + "``monospace content``"}, // 2 back slashes "consumed"
							},
						},
					}
					Expect(source).To(EqualDocumentBlock(expected))
				})

				It("escaped monospace text with unbalanced double quote", func() {
					source := `\` + "``monospace content`"
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "``monospace content`"},
							},
						},
					}
					Expect(source).To(EqualDocumentBlock(expected))
				})

				It("escaped monospace text with unbalanced double quote and more backslashes", func() {
					source := `\\\` + "``monospace content`" // 3 backslashes
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `\\` + "``monospace content`"}, // 2 backslashes removed
							},
						},
					}
					Expect(source).To(EqualDocumentBlock(expected))
				})
			})

			Context("with nested quoted text", func() {

				It("escaped monospace text with nested bold text", func() {
					source := `\` + "`*bold content*`" // gives: \`*bold content*`
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "`"},
								types.QuotedText{
									Kind: types.Bold,
									Elements: types.InlineElements{
										types.StringElement{Content: "bold content"},
									},
								},
								types.StringElement{Content: "`"},
							},
						},
					}
					Expect(source).To(EqualDocumentBlock(expected))
				})

				It("escaped monospace text with unbalanced double backquote and nested bold test", func() {
					source := `\` + "``*bold content*`"
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "``"},
								types.QuotedText{
									Kind: types.Bold,
									Elements: types.InlineElements{
										types.StringElement{Content: "bold content"},
									},
								},
								types.StringElement{Content: "`"},
							},
						},
					}
					Expect(source).To(EqualDocumentBlock(expected))
				})

				It("escaped monospace text with nested bold text", func() {
					source := `\` + "`monospace *and bold* content`"
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "`monospace "},
								types.QuotedText{
									Kind: types.Bold,
									Elements: types.InlineElements{
										types.StringElement{Content: "and bold"},
									},
								},
								types.StringElement{Content: " content`"},
							},
						},
					}
					Expect(source).To(EqualDocumentBlock(expected))
				})
			})
		})

		Context("prevented subscript text substitution", func() {

			Context("without nested quoted text", func() {

				It("escaped subscript text with single quote", func() {
					source := `\~subscriptcontent~`
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "~subscriptcontent~"},
							},
						},
					}
					Expect(source).To(EqualDocumentBlock(expected))
				})

				It("escaped subscript text with single quote and more backslashes", func() {
					source := `\\~subscriptcontent~`
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `\~subscriptcontent~`}, // only 1 backslash removed
							},
						},
					}
					Expect(source).To(EqualDocumentBlock(expected))
				})

			})

			Context("with nested quoted text", func() {

				It("escaped subscript text with nested bold text", func() {
					source := `\~*boldcontent*~`
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "~"},
								types.QuotedText{
									Kind: types.Bold,
									Elements: types.InlineElements{
										types.StringElement{Content: "boldcontent"},
									},
								},
								types.StringElement{Content: "~"},
							},
						},
					}
					Expect(source).To(EqualDocumentBlock(expected))
				})

				It("escaped subscript text with nested bold text", func() {
					source := `\~subscript *and bold* content~`
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `\~subscript `},
								types.QuotedText{
									Kind: types.Bold,
									Elements: types.InlineElements{
										types.StringElement{Content: "and bold"},
									},
								},
								types.StringElement{Content: " content~"},
							},
						},
					}
					Expect(source).To(EqualDocumentBlock(expected))
				})
			})
		})

		Context("prevented superscript text substitution", func() {

			Context("without nested quoted text", func() {

				It("escaped superscript text with single quote", func() {
					source := `\^superscriptcontent^`
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "^superscriptcontent^"},
							},
						},
					}
					Expect(source).To(EqualDocumentBlock(expected))
				})

				It("escaped superscript text with single quote and more backslashes", func() {
					source := `\\^superscriptcontent^`
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `\^superscriptcontent^`}, // only 1 backslash removed
							},
						},
					}
					Expect(source).To(EqualDocumentBlock(expected))
				})

			})

			Context("with nested quoted text", func() {

				It("escaped superscript text with nested bold text - case 1", func() {
					source := `\^*bold content*^` // valid escaped superscript since it has no space within
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `^`},
								types.QuotedText{
									Kind: types.Bold,
									Elements: types.InlineElements{
										types.StringElement{Content: "bold content"},
									},
								},
								types.StringElement{Content: "^"},
							},
						},
					}
					Expect(source).To(EqualDocumentBlock(expected))
				})

				It("escaped superscript text with unbalanced double backquote and nested bold test", func() {
					source := `\^*bold content*^`
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "^"},
								types.QuotedText{
									Kind: types.Bold,
									Elements: types.InlineElements{
										types.StringElement{Content: "bold content"},
									},
								},
								types.StringElement{Content: "^"},
							},
						},
					}
					Expect(source).To(EqualDocumentBlock(expected))
				})

				It("escaped superscript text with nested bold text - case 2", func() {
					source := `\^superscript *and bold* content^` // invalid superscript text since it has spaces within
					expected := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `\^superscript `},
								types.QuotedText{
									Kind: types.Bold,
									Elements: types.InlineElements{
										types.StringElement{Content: "and bold"},
									},
								},
								types.StringElement{Content: " content^"},
							},
						},
					}
					Expect(source).To(EqualDocumentBlock(expected))
				})
			})
		})
	})
})
