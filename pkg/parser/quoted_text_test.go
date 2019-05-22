package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("quoted texts", func() {

	Context("quoted text with single punctuation", func() {

		It("bold text with 1 word", func() {
			actualContent := "*hello*"
			expectedResult := types.QuotedText{
				Kind: types.Bold,
				Elements: types.InlineElements{
					types.StringElement{Content: "hello"},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("bold text with 2 words", func() {
			actualContent := "*bold    content*"
			expectedResult := types.QuotedText{
				Kind: types.Bold,
				Elements: types.InlineElements{
					types.StringElement{Content: "bold    content"},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("bold text with 3 words", func() {
			actualContent := "*some bold content*"
			expectedResult := types.QuotedText{
				Kind: types.Bold,
				Elements: types.InlineElements{
					types.StringElement{Content: "some bold content"},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("italic text with 3 words in single quote", func() {
			actualContent := "_some italic content_"
			expectedResult := types.QuotedText{
				Kind: types.Italic,
				Elements: types.InlineElements{
					types.StringElement{Content: "some italic content"},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("monospace text with 3 words", func() {
			actualContent := "`some monospace content`"
			expectedResult := types.QuotedText{
				Kind: types.Monospace,
				Elements: types.InlineElements{
					types.StringElement{Content: "some monospace content"},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("invalid subscript text with 3 words", func() {
			actualContent := "~some subscript content~"
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "~some subscript content~"},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Paragraph"))
		})

		It("invalid superscript text with 3 words", func() {
			actualContent := "^some superscript content^"
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "^some superscript content^"},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Paragraph"))
		})

		It("bold text within italic text", func() {
			actualContent := "_some *bold* content_"
			expectedResult := types.QuotedText{
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
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("monospace text within bold text within italic quote", func() {
			actualContent := "*some _italic and `monospaced content`_*"
			expectedResult := types.QuotedText{
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
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("italic text within italic text", func() {
			actualContent := "_some _very italic_ content_"
			expectedResult := types.InlineElements{
				types.QuotedText{
					Kind: types.Italic,
					Elements: types.InlineElements{
						types.StringElement{Content: "some _very italic"},
					},
				},
				types.StringElement{Content: " content_"},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("subscript text attached", func() {
			actualContent := "O~2~ is a molecule"
			expectedResult := types.InlineElements{
				types.StringElement{Content: "O"},
				types.QuotedText{
					Kind: types.Subscript,
					Elements: types.InlineElements{
						types.StringElement{Content: "2"},
					},
				},
				types.StringElement{Content: " is a molecule"},
			}

			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("superscript text attached", func() {
			actualContent := "M^me^ White"
			expectedResult := types.InlineElements{
				types.StringElement{Content: "M"},
				types.QuotedText{
					Kind: types.Superscript,
					Elements: types.InlineElements{
						types.StringElement{Content: "me"},
					},
				},
				types.StringElement{Content: " White"},
			}

			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("invalid subscript text with 3 words", func() {
			actualContent := "~some subscript content~"
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "~some subscript content~"},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Paragraph"))
		})

	})

	Context("Quoted text with double punctuation", func() {

		It("bold text of 1 word in double quote", func() {
			actualContent := "**hello**"
			expectedResult := types.QuotedText{
				Kind: types.Bold,
				Elements: types.InlineElements{
					types.StringElement{Content: "hello"},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("italic text with 3 words in double quote", func() {
			actualContent := "__some italic content__"
			expectedResult := types.QuotedText{
				Kind: types.Italic,
				Elements: types.InlineElements{
					types.StringElement{Content: "some italic content"},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("monospace text with 3 words in double quote", func() {
			actualContent := "``some monospace content``"
			expectedResult := types.QuotedText{
				Kind: types.Monospace,
				Elements: types.InlineElements{
					types.StringElement{Content: "some monospace content"},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("superscript text within italic text", func() {
			actualContent := "__some ^superscript^ content__"
			expectedResult := types.QuotedText{
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
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("superscript text within italic text within bold quote", func() {
			actualContent := "**some _italic and ^superscriptcontent^_**"
			expectedResult := types.QuotedText{
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
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

	})

	Context("Quoted text inline", func() {

		It("inline content with bold text", func() {
			actualContent := "a paragraph with *some bold content*"
			expectedResult := types.InlineElements{
				types.StringElement{Content: "a paragraph with "},
				types.QuotedText{
					Kind: types.Bold,
					Elements: types.InlineElements{
						types.StringElement{Content: "some bold content"},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("inline content with invalid bold text - use case 1", func() {
			actualContent := "a paragraph with *some bold content"
			expectedResult := types.InlineElements{
				types.StringElement{Content: "a paragraph with *some bold content"},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("inline content with invalid bold text - use case 2", func() {
			actualContent := "a paragraph with *some bold content *"
			expectedResult := types.InlineElements{
				types.StringElement{Content: "a paragraph with *some bold content *"},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("inline content with invalid bold text - use case 3", func() {
			actualContent := "a paragraph with * some bold content*"
			expectedResult := types.InlineElements{
				types.StringElement{Content: "a paragraph with * some bold content*"},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("invalid italic text within bold text", func() {
			actualContent := "some *bold and _italic content _ together*."
			expectedResult := types.InlineElements{
				types.StringElement{Content: "some "},
				types.QuotedText{
					Kind: types.Bold,
					Elements: types.InlineElements{
						types.StringElement{Content: "bold and _italic content _ together"},
					},
				},
				types.StringElement{Content: "."},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("italic text within invalid bold text", func() {
			actualContent := "some *bold and _italic content_ together *."
			expectedResult := types.InlineElements{
				types.StringElement{Content: "some *bold and "},
				types.QuotedText{
					Kind: types.Italic,
					Elements: types.InlineElements{
						types.StringElement{Content: "italic content"},
					},
				},
				types.StringElement{Content: " together *."},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("inline content with invalid subscript text - use case 1", func() {
			actualContent := "a paragraph with ~some subscript content"
			expectedResult := types.InlineElements{
				types.StringElement{Content: "a paragraph with ~some subscript content"},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("inline content with invalid subscript text - use case 2", func() {
			actualContent := "a paragraph with ~some subscript content ~"
			expectedResult := types.InlineElements{
				types.StringElement{Content: "a paragraph with ~some subscript content ~"},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("inline content with invalid subscript text - use case 3", func() {
			actualContent := "a paragraph with ~ some subscript content~"
			expectedResult := types.InlineElements{
				types.StringElement{Content: "a paragraph with ~ some subscript content~"},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("inline content with invalid superscript text - use case 1", func() {
			actualContent := "a paragraph with ^some superscript content"
			expectedResult := types.InlineElements{
				types.StringElement{Content: "a paragraph with ^some superscript content"},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("inline content with invalid superscript text - use case 2", func() {
			actualContent := "a paragraph with ^some superscript content ^"
			expectedResult := types.InlineElements{
				types.StringElement{Content: "a paragraph with ^some superscript content ^"},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("inline content with invalid superscript text - use case 3", func() {
			actualContent := "a paragraph with ^ some superscript content^"
			expectedResult := types.InlineElements{
				types.StringElement{Content: "a paragraph with ^ some superscript content^"},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})
	})

	Context("nested quoted text", func() {

		It("italic text within bold text", func() {
			actualContent := "some *bold and _italic content_ together*."
			expectedResult := types.InlineElements{
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
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("single-quote bold within single-quote bold text", func() {
			// here we don't allow for bold text within bold text, to comply with the existing implementations (asciidoc and asciidoctor)
			actualContent := "*some *nested bold* content*"
			expectedResult := types.InlineElements{
				types.QuotedText{
					Kind: types.Bold,
					Elements: types.InlineElements{
						types.StringElement{Content: "some *nested bold"},
					},
				},
				types.StringElement{Content: " content*"},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("double-quote bold within double-quote bold text", func() {
			// here we don't allow for bold text within bold text, to comply with the existing implementations (asciidoc and asciidoctor)
			actualContent := "**some **nested bold** content**"
			expectedResult := types.InlineElements{
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
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("single-quote bold within double-quote bold text", func() {
			// here we don't allow for bold text within bold text, to comply with the existing implementations (asciidoc and asciidoctor)
			actualContent := "**some *nested bold* content**"
			expectedResult := types.InlineElements{
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
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("double-quote bold within single-quote bold text", func() {
			// here we don't allow for bold text within bold text, to comply with the existing implementations (asciidoc and asciidoctor)
			actualContent := "*some **nested bold** content*"
			expectedResult := types.InlineElements{
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
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("single-quote italic within single-quote italic text", func() {
			// here we don't allow for italic text within italic text, to comply with the existing implementations (asciidoc and asciidoctor)
			actualContent := "_some _nested italic_ content_"
			expectedResult := types.InlineElements{
				types.QuotedText{
					Kind: types.Italic,
					Elements: types.InlineElements{
						types.StringElement{Content: "some _nested italic"},
					},
				},
				types.StringElement{Content: " content_"},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("double-quote italic within double-quote italic text", func() {
			// here we don't allow for italic text within italic text, to comply with the existing implementations (asciidoc and asciidoctor)
			actualContent := "__some __nested italic__ content__"
			expectedResult := types.InlineElements{
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
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("single-quote italic within double-quote italic text", func() {
			// here we allow for italic text within italic text, to comply with the existing implementations (asciidoc and asciidoctor)
			actualContent := "_some __nested italic__ content_"
			expectedResult := types.InlineElements{
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
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("double-quote italic within single-quote italic text", func() {
			// here we allow for italic text within italic text, to comply with the existing implementations (asciidoc and asciidoctor)
			actualContent := "_some __nested italic__ content_"
			expectedResult := types.InlineElements{
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
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("single-quote monospace within single-quote monospace text", func() {
			// here we don't allow for monospace text within monospace text, to comply with the existing implementations (asciidoc and asciidoctor)
			actualContent := "`some `nested monospace` content`"
			expectedResult := types.InlineElements{
				types.QuotedText{
					Kind: types.Monospace,
					Elements: types.InlineElements{
						types.StringElement{Content: "some `nested monospace"},
					},
				},
				types.StringElement{Content: " content`"},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("double-quote monospace within double-quote monospace text", func() {
			// here we don't allow for monospace text within monospace text, to comply with the existing implementations (asciidoc and asciidoctor)
			actualContent := "``some ``nested monospace`` content``"
			expectedResult := types.InlineElements{
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
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("single-quote monospace within double-quote monospace text", func() {
			// here we allow for monospace text within monospace text, to comply with the existing implementations (asciidoc and asciidoctor)
			actualContent := "`some ``nested monospace`` content`"
			expectedResult := types.InlineElements{
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
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("double-quote monospace within single-quote monospace text", func() {
			// here we allow for monospace text within monospace text, to comply with the existing implementations (asciidoc and asciidoctor)
			actualContent := "`some ``nested monospace`` content`"
			expectedResult := types.InlineElements{
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
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("unbalanced bold in monospace - case 1", func() {
			actualContent := "`*a`"
			expectedResult := types.InlineElements{
				types.QuotedText{
					Kind: types.Monospace,
					Elements: types.InlineElements{
						types.StringElement{Content: "*a"},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("unbalanced bold in monospace - case 2", func() {
			actualContent := "`a*b`"
			expectedResult := types.InlineElements{
				types.QuotedText{
					Kind: types.Monospace,
					Elements: types.InlineElements{
						types.StringElement{Content: "a*b"},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("italic in monospace", func() {
			actualContent := "`_a_`"
			expectedResult := types.InlineElements{
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
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("unbalanced italic in monospace", func() {
			actualContent := "`a_b`"
			expectedResult := types.InlineElements{
				types.QuotedText{
					Kind: types.Monospace,
					Elements: types.InlineElements{
						types.StringElement{Content: "a_b"},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("unparsed bold in monospace", func() {
			actualContent := "`a*b*`"
			expectedResult := types.InlineElements{
				types.QuotedText{
					Kind: types.Monospace,
					Elements: types.InlineElements{
						types.StringElement{Content: "a*b*"},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("parsed subscript in monospace", func() {
			actualContent := "`a~b~`"
			expectedResult := types.InlineElements{
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
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("multiline in single quoted monospace - case 1", func() {
			actualContent := "`a\nb`"
			expectedResult := types.InlineElements{
				types.QuotedText{
					Kind: types.Monospace,
					Elements: types.InlineElements{
						types.StringElement{Content: "a\nb"},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("multiline in double quoted monospace - case 1", func() {
			actualContent := "`a\nb`"
			expectedResult := types.InlineElements{
				types.QuotedText{
					Kind: types.Monospace,
					Elements: types.InlineElements{
						types.StringElement{Content: "a\nb"},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("multiline in single quoted  monospace - case 2", func() {
			actualContent := "`a\n*b*`"
			expectedResult := types.InlineElements{
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
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("multiline in double quoted  monospace - case 2", func() {
			actualContent := "`a\n*b*`"
			expectedResult := types.InlineElements{
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
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("link in bold", func() {
			actualContent := "*a link:/[b]*"
			expectedResult := types.InlineElements{
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
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("image in bold", func() {
			actualContent := "*a image:foo.png[]*"
			expectedResult := types.InlineElements{
				types.QuotedText{
					Kind: types.Bold,
					Elements: types.InlineElements{
						types.StringElement{Content: "a "},
						types.InlineImage{
							Attributes: types.ElementAttributes{
								types.AttrImageAlt:    "foo",
								types.AttrImageHeight: "",
								types.AttrImageWidth:  "",
							},
							Path: "foo.png",
						},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("singleplus passthrough in bold", func() {
			actualContent := "*a +image:foo.png[]+*"
			expectedResult := types.InlineElements{
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
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("tripleplus passthrough in bold", func() {
			actualContent := "*a +++image:foo.png[]+++*"
			expectedResult := types.InlineElements{
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
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("link in italic", func() {
			actualContent := "_a link:/[b]_"
			expectedResult := types.InlineElements{
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
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("image in italic", func() {
			actualContent := "_a image:foo.png[]_"
			expectedResult := types.InlineElements{
				types.QuotedText{
					Kind: types.Italic,
					Elements: types.InlineElements{
						types.StringElement{Content: "a "},
						types.InlineImage{
							Attributes: types.ElementAttributes{
								types.AttrImageAlt:    "foo",
								types.AttrImageHeight: "",
								types.AttrImageWidth:  "",
							},
							Path: "foo.png",
						},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("singleplus passthrough in italic", func() {
			actualContent := "_a +image:foo.png[]+_"
			expectedResult := types.InlineElements{
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
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("tripleplus passthrough in italic", func() {
			actualContent := "_a +++image:foo.png[]+++_"
			expectedResult := types.InlineElements{
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
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("link in monospace", func() {
			actualContent := "`a link:/[b]`"
			expectedResult := types.InlineElements{
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
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("image in monospace", func() {
			actualContent := "`a image:foo.png[]`"
			expectedResult := types.InlineElements{
				types.QuotedText{
					Kind: types.Monospace,
					Elements: types.InlineElements{
						types.StringElement{Content: "a "},
						types.InlineImage{
							Attributes: types.ElementAttributes{
								types.AttrImageAlt:    "foo",
								types.AttrImageHeight: "",
								types.AttrImageWidth:  "",
							},
							Path: "foo.png",
						},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("singleplus passthrough in monospace", func() {
			actualContent := "`a +image:foo.png[]+`"
			expectedResult := types.InlineElements{
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
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("tripleplus passthrough in monospace", func() {
			actualContent := "`a +++image:foo.png[]+++`"
			expectedResult := types.InlineElements{
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
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

	})

	Context("unbalanced quoted text", func() {

		Context("unbalanced bold text", func() {

			It("unbalanced bold text - extra on left", func() {
				actualContent := "**some bold content*"
				expectedResult := types.InlineElements{
					types.QuotedText{
						Kind: types.Bold,
						Elements: types.InlineElements{
							types.StringElement{Content: "*some bold content"},
						},
					},
				}
				verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
			})

			It("unbalanced bold text - extra on right", func() {
				actualContent := "*some bold content**"
				expectedResult := types.InlineElements{
					types.QuotedText{
						Kind: types.Bold,
						Elements: types.InlineElements{
							types.StringElement{Content: "some bold content"},
						},
					},
					types.StringElement{Content: "*"},
				}
				verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
			})
		})

		Context("unbalanced italic text", func() {

			It("unbalanced italic text - extra on left", func() {
				actualContent := "__some italic content_"
				expectedResult := types.InlineElements{
					types.QuotedText{
						Kind: types.Italic,
						Elements: types.InlineElements{
							types.StringElement{Content: "_some italic content"},
						},
					},
				}
				verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
			})

			It("unbalanced italic text - extra on right", func() {
				actualContent := "_some italic content__"
				expectedResult := types.InlineElements{
					types.QuotedText{
						Kind: types.Italic,
						Elements: types.InlineElements{
							types.StringElement{Content: "some italic content"},
						},
					},
					types.StringElement{Content: "_"},
				}
				verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
			})
		})

		Context("unbalanced monospace text", func() {

			It("unbalanced monospace text - extra on left", func() {
				actualContent := "``some monospace content`"
				expectedResult := types.InlineElements{
					types.QuotedText{
						Kind: types.Monospace,
						Elements: types.InlineElements{
							types.StringElement{Content: "`some monospace content"},
						},
					},
				}
				verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
			})

			It("unbalanced monospace text - extra on right", func() {
				actualContent := "`some monospace content``"
				expectedResult := types.InlineElements{
					types.QuotedText{
						Kind: types.Monospace,
						Elements: types.InlineElements{
							types.StringElement{Content: "some monospace content"},
						},
					},
					types.StringElement{Content: "`"},
				}
				verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
			})
		})

		It("inline content with unbalanced bold text", func() {
			actualContent := "a paragraph with *some bold content"
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "a paragraph with *some bold content"},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

	})

	Context("prevented substitution", func() {

		Context("prevented bold text substitution", func() {

			Context("without nested quoted text", func() {

				It("escaped bold text with single backslash", func() {
					actualContent := `\*bold content*`
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "*bold content*"},
							},
						},
					}
					verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped bold text with multiple backslashes", func() {
					actualContent := `\\*bold content*`
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `\*bold content*`},
							},
						},
					}
					verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped bold text with double quote", func() {
					actualContent := `\\**bold content**`
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `**bold content**`},
							},
						},
					}
					verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped bold text with double quote and more backslashes", func() {
					actualContent := `\\\**bold content**`
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `\**bold content**`},
							},
						},
					}
					verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped bold text with unbalanced double quote", func() {
					actualContent := `\**bold content*`
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `**bold content*`},
							},
						},
					}
					verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped bold text with unbalanced double quote and more backslashes", func() {
					actualContent := `\\\**bold content*`
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `\\**bold content*`},
							},
						},
					}
					verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})
			})

			Context("with nested quoted text", func() {

				It("escaped bold text with nested italic text", func() {
					actualContent := `\*_italic content_*`
					expectedResult := types.Paragraph{
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
					verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped bold text with unbalanced double quote and nested italic test", func() {
					actualContent := `\**_italic content_*`
					expectedResult := types.Paragraph{
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
					verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped bold text with nested italic", func() {
					actualContent := `\*bold _and italic_ content*`
					expectedResult := types.Paragraph{
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
					verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})
			})

		})

		Context("prevented italic text substitution", func() {

			Context("without nested quoted text", func() {

				It("escaped italic text with single quote", func() {
					actualContent := `\_italic content_`
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "_italic content_"},
							},
						},
					}
					verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped italic text with single quote and more backslashes", func() {
					actualContent := `\\_italic content_`
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `\_italic content_`},
							},
						},
					}
					verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped italic text with double quote with 2 backslashes", func() {
					actualContent := `\\__italic content__`
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `__italic content__`},
							},
						},
					}
					verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped italic text with double quote with 3 backslashes", func() {
					actualContent := `\\\__italic content__`
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `\__italic content__`},
							},
						},
					}
					verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped italic text with unbalanced double quote", func() {
					actualContent := `\__italic content_`
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `__italic content_`},
							},
						},
					}
					verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped italic text with unbalanced double quote and more backslashes", func() {
					actualContent := `\\\__italic content_`
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `\\__italic content_`}, // only 1 backslash remove
							},
						},
					}
					verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})
			})

			Context("with nested quoted text", func() {

				It("escaped italic text with nested monospace text", func() {
					actualContent := `\` + "_`monospace content`_" // gives: \_`monospace content`_
					expectedResult := types.Paragraph{
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
					verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped italic text with unbalanced double quote and nested bold test", func() {
					actualContent := `\__*bold content*_`
					expectedResult := types.Paragraph{
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
					verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped italic text with nested bold text", func() {
					actualContent := `\_italic *and bold* content_`
					expectedResult := types.Paragraph{
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
					verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})
			})
		})

		Context("prevented monospace text substitution", func() {

			Context("without nested quoted text", func() {

				It("escaped monospace text with single quote", func() {
					actualContent := `\` + "`monospace content`"
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "`monospace content`"}, // backslash removed
							},
						},
					}
					verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped monospace text with single quote and more backslashes", func() {
					actualContent := `\\` + "`monospace content`"
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `\` + "`monospace content`"}, // only 1 backslash removed
							},
						},
					}
					verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped monospace text with double quote", func() {
					actualContent := `\\` + "`monospace content``"
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `\` + "`monospace content``"}, // 2 back slashes "consumed"
							},
						},
					}
					verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped monospace text with double quote and more backslashes", func() {
					actualContent := `\\\` + "``monospace content``" // 3 backslashes
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `\` + "``monospace content``"}, // 2 back slashes "consumed"
							},
						},
					}
					verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped monospace text with unbalanced double quote", func() {
					actualContent := `\` + "``monospace content`"
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "``monospace content`"},
							},
						},
					}
					verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped monospace text with unbalanced double quote and more backslashes", func() {
					actualContent := `\\\` + "``monospace content`" // 3 backslashes
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `\\` + "``monospace content`"}, // 2 backslashes removed
							},
						},
					}
					verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})
			})

			Context("with nested quoted text", func() {

				It("escaped monospace text with nested bold text", func() {
					actualContent := `\` + "`*bold content*`" // gives: \`*bold content*`
					expectedResult := types.Paragraph{
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
					verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped monospace text with unbalanced double backquote and nested bold test", func() {
					actualContent := `\` + "``*bold content*`"
					expectedResult := types.Paragraph{
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
					verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped monospace text with nested bold text", func() {
					actualContent := `\` + "`monospace *and bold* content`"
					expectedResult := types.Paragraph{
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
					verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})
			})
		})

		Context("prevented subscript text substitution", func() {

			Context("without nested quoted text", func() {

				It("escaped subscript text with single quote", func() {
					actualContent := `\~subscriptcontent~`
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "~subscriptcontent~"},
							},
						},
					}
					verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped subscript text with single quote and more backslashes", func() {
					actualContent := `\\~subscriptcontent~`
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `\~subscriptcontent~`}, // only 1 backslash removed
							},
						},
					}
					verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

			})

			Context("with nested quoted text", func() {

				It("escaped subscript text with nested bold text", func() {
					actualContent := `\~*boldcontent*~`
					expectedResult := types.Paragraph{
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
					verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped subscript text with nested bold text", func() {
					actualContent := `\~subscript *and bold* content~`
					expectedResult := types.Paragraph{
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
					verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})
			})
		})

		Context("prevented superscript text substitution", func() {

			Context("without nested quoted text", func() {

				It("escaped superscript text with single quote", func() {
					actualContent := `\^superscriptcontent^`
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "^superscriptcontent^"},
							},
						},
					}
					verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped superscript text with single quote and more backslashes", func() {
					actualContent := `\\^superscriptcontent^`
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `\^superscriptcontent^`}, // only 1 backslash removed
							},
						},
					}
					verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

			})

			Context("with nested quoted text", func() {

				It("escaped superscript text with nested bold text - case 1", func() {
					actualContent := `\^*bold content*^` // valid escaped superscript since it has no space within
					expectedResult := types.Paragraph{
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
					verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped superscript text with unbalanced double backquote and nested bold test", func() {
					actualContent := `\^*bold content*^`
					expectedResult := types.Paragraph{
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
					verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped superscript text with nested bold text - case 2", func() {
					actualContent := `\^superscript *and bold* content^` // invalid superscript text since it has spaces within
					expectedResult := types.Paragraph{
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
					verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})
			})
		})
	})
})
