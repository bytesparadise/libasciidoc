package parser_test

import (
	"github.com/bytesparadise/libasciidoc/parser"
	"github.com/bytesparadise/libasciidoc/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("quoted texts", func() {

	Context("Quoted text with single punctuation", func() {

		It("bold text of 1 word", func() {
			actualContent := "*hello*"
			expectedResult := types.QuotedText{
				Kind: types.Bold,
				Elements: []interface{}{
					types.StringElement{Content: "hello"},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("bold text of 2 words", func() {
			actualContent := "*bold    content*"
			expectedResult := types.QuotedText{
				Kind: types.Bold,
				Elements: []interface{}{
					types.StringElement{Content: "bold    content"},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("bold text of 3 words", func() {
			actualContent := "*some bold content*"
			expectedResult := types.QuotedText{
				Kind: types.Bold,
				Elements: []interface{}{
					types.StringElement{Content: "some bold content"},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("italic text with3 words", func() {
			actualContent := "_some italic content_"
			expectedResult := types.QuotedText{
				Kind: types.Italic,
				Elements: []interface{}{
					types.StringElement{Content: "some italic content"},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("monospace text with3 words", func() {
			actualContent := "`some monospace content`"
			expectedResult := types.QuotedText{
				Kind: types.Monospace,
				Elements: []interface{}{
					types.StringElement{Content: "some monospace content"},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("bold text within italic text", func() {
			actualContent := "_some *bold* content_"
			expectedResult := types.QuotedText{
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
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("monospace text within bold text within italic quote", func() {
			actualContent := "*some _italic and `monospaced content`_*"
			expectedResult := types.QuotedText{
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
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("italic text within italic text", func() {
			actualContent := "_some _very italic_ content_"
			expectedResult := types.QuotedText{
				Kind: types.Italic,
				Elements: []interface{}{
					types.StringElement{Content: "some "},
					types.QuotedText{
						Kind: types.Italic,
						Elements: []interface{}{
							types.StringElement{Content: "very italic"},
						},
					},
					types.StringElement{Content: " content"},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})
	})

	Context("Quoted text with double punctuation", func() {

		It("bold text of 1 word", func() {
			actualContent := "**hello**"
			expectedResult := types.QuotedText{
				Kind: types.Bold,
				Elements: []interface{}{
					types.StringElement{Content: "hello"},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("italic text with3 words", func() {
			actualContent := "__some italic content__"
			expectedResult := types.QuotedText{
				Kind: types.Italic,
				Elements: []interface{}{
					types.StringElement{Content: "some italic content"},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("monospace text with3 words", func() {
			actualContent := "``some monospace content``"
			expectedResult := types.QuotedText{
				Kind: types.Monospace,
				Elements: []interface{}{
					types.StringElement{Content: "some monospace content"},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("bold text within italic text", func() {
			actualContent := "__some *bold* content__"
			expectedResult := types.QuotedText{
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
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("monospace text within bold text within italic quote", func() {
			actualContent := "**some _italic and ``monospaced content``_**"
			expectedResult := types.QuotedText{
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
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

	})

	Context("Quoted text inline", func() {

		It("inline with bold text", func() {
			actualContent := "a paragraph with *some bold content*"
			expectedResult := types.InlineElements{
				types.StringElement{Content: "a paragraph with "},
				types.QuotedText{
					Kind: types.Bold,
					Elements: []interface{}{
						types.StringElement{Content: "some bold content"},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("inline with invalid bold text1", func() {
			actualContent := "a paragraph with *some bold content"
			expectedResult := types.InlineElements{
				types.StringElement{Content: "a paragraph with *some bold content"},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("inline with invalid bold text2", func() {
			actualContent := "a paragraph with *some bold content *"
			expectedResult := types.InlineElements{
				types.StringElement{Content: "a paragraph with *some bold content *"},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("inline with invalid bold text3", func() {
			actualContent := "a paragraph with * some bold content*"
			expectedResult := types.InlineElements{
				types.StringElement{Content: "a paragraph with * some bold content*"},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("invalid italic text within bold text", func() {
			actualContent := "some *bold and _italic content _ together*."
			expectedResult := types.InlineElements{
				types.StringElement{Content: "some "},
				types.QuotedText{
					Kind: types.Bold,
					Elements: []interface{}{
						types.StringElement{Content: "bold and _italic content _ together"},
					},
				},
				types.StringElement{Content: "."},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("italic text within invalid bold text", func() {
			actualContent := "some *bold and _italic content_ together *."
			expectedResult := types.InlineElements{
				types.StringElement{Content: "some *bold and "},
				types.QuotedText{
					Kind: types.Italic,
					Elements: []interface{}{
						types.StringElement{Content: "italic content"},
					},
				},
				types.StringElement{Content: " together *."},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})
	})

	Context("Nested quoted text", func() {

		It("italic text within bold text", func() {
			actualContent := "some *bold and _italic content_ together*."
			expectedResult := types.InlineElements{
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
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("simple-quote bold within simple-quote bold text", func() {
			// here we don't allow for bold text within bold text, to comply with the existing implementations (asciidoc and asciidoctor)
			actualContent := "*some *nested bold* content*"
			expectedResult := types.InlineElements{
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
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("double-quote bold within double-quote bold text", func() {
			// here we don't allow for bold text within bold text, to comply with the existing implementations (asciidoc and asciidoctor)
			actualContent := "**some **nested bold** content**"
			expectedResult := types.InlineElements{
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
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("simple-quote bold within double-quote bold text", func() {
			// here we don't allow for bold text within bold text, to comply with the existing implementations (asciidoc and asciidoctor)
			actualContent := "**some *nested bold* content**"
			expectedResult := types.InlineElements{
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
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("double-quote bold within simple-quote bold text", func() {
			// here we don't allow for bold text within bold text, to comply with the existing implementations (asciidoc and asciidoctor)
			actualContent := "*some **nested bold** content*"
			expectedResult := types.InlineElements{
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
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("simple-quote italic within simple-quote italic text", func() {
			// here we don't allow for italic text within italic text, to comply with the existing implementations (asciidoc and asciidoctor)
			actualContent := "_some _nested italic_ content_"
			expectedResult := types.InlineElements{
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
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("double-quote italic within double-quote italic text", func() {
			// here we don't allow for italic text within italic text, to comply with the existing implementations (asciidoc and asciidoctor)
			actualContent := "__some __nested italic__ content__"
			expectedResult := types.InlineElements{
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
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("simple-quote italic within double-quote italic text", func() {
			// here we allow for italic text within italic text, to comply with the existing implementations (asciidoc and asciidoctor)
			actualContent := "_some __nested italic__ content_"
			expectedResult := types.InlineElements{
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
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("double-quote italic within simple-quote italic text", func() {
			// here we allow for italic text within italic text, to comply with the existing implementations (asciidoc and asciidoctor)
			actualContent := "_some __nested italic__ content_"
			expectedResult := types.InlineElements{
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
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("simple-quote monospace within simple-quote monospace text", func() {
			// here we don't allow for monospace text within monospace text, to comply with the existing implementations (asciidoc and asciidoctor)
			actualContent := "`some `nested monospace` content`"
			expectedResult := types.InlineElements{
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
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("double-quote monospace within double-quote monospace text", func() {
			// here we don't allow for monospace text within monospace text, to comply with the existing implementations (asciidoc and asciidoctor)
			actualContent := "``some ``nested monospace`` content``"
			expectedResult := types.InlineElements{
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
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("simple-quote monospace within double-quote monospace text", func() {
			// here we allow for monospace text within monospace text, to comply with the existing implementations (asciidoc and asciidoctor)
			actualContent := "`some ``nested monospace`` content`"
			expectedResult := types.InlineElements{
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
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("double-quote monospace within simple-quote monospace text", func() {
			// here we allow for monospace text within monospace text, to comply with the existing implementations (asciidoc and asciidoctor)
			actualContent := "`some ``nested monospace`` content`"
			expectedResult := types.InlineElements{
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
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

	})

	Context("Unbalanced quoted text", func() {

		Context("Unbalanced bold text", func() {

			It("unbalanced bold text - extra on left", func() {
				actualContent := "**some bold content*"
				expectedResult := types.InlineElements{
					types.QuotedText{
						Kind: types.Bold,
						Elements: []interface{}{
							types.StringElement{Content: "*some bold content"},
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
			})

			It("unbalanced bold text - extra on right", func() {
				actualContent := "*some bold content**"
				expectedResult := types.InlineElements{
					types.QuotedText{
						Kind: types.Bold,
						Elements: []interface{}{
							types.StringElement{Content: "some bold content"},
						},
					},
					types.StringElement{Content: "*"},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
			})
		})

		Context("unbalanced italic text", func() {

			It("unbalanced italic text - extra on left", func() {
				actualContent := "__some italic content_"
				expectedResult := types.InlineElements{
					types.QuotedText{
						Kind: types.Italic,
						Elements: []interface{}{
							types.StringElement{Content: "_some italic content"},
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
			})

			It("unbalanced italic text - extra on right", func() {
				actualContent := "_some italic content__"
				expectedResult := types.InlineElements{
					types.QuotedText{
						Kind: types.Italic,
						Elements: []interface{}{
							types.StringElement{Content: "some italic content"},
						},
					},
					types.StringElement{Content: "_"},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
			})
		})

		Context("unbalanced monospace text", func() {

			It("unbalanced monospace text - extra on left", func() {
				actualContent := "``some monospace content`"
				expectedResult := types.InlineElements{
					types.QuotedText{
						Kind: types.Monospace,
						Elements: []interface{}{
							types.StringElement{Content: "`some monospace content"},
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
			})

			It("unbalanced monospace text - extra on right", func() {
				actualContent := "`some monospace content``"
				expectedResult := types.InlineElements{
					types.QuotedText{
						Kind: types.Monospace,
						Elements: []interface{}{
							types.StringElement{Content: "some monospace content"},
						},
					},
					types.StringElement{Content: "`"},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
			})
		})

		It("inline with unbalanced bold text", func() {
			actualContent := "a paragraph with *some bold content"
			expectedResult := types.Paragraph{
				Attributes: map[string]interface{}{},
				Lines: []types.InlineElements{
					{
						types.StringElement{Content: "a paragraph with *some bold content"},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Paragraph"))
		})

	})

	Context("prevented substitution", func() {

		Context("prevented Bold text substitution", func() {

			It("escaped bold text with simple quote", func() {
				actualContent := `\*bold content*`
				expectedResult := types.Paragraph{
					Attributes: map[string]interface{}{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "*bold content*"},
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Paragraph"))
			})

			It("escaped bold text with simple quote and more backslashes", func() {
				actualContent := `\\*bold content*`
				expectedResult := types.Paragraph{
					Attributes: map[string]interface{}{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: `\*bold content*`},
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Paragraph"))
			})

			It("escaped bold text with double quote", func() {
				actualContent := `\\**bold content**`
				expectedResult := types.Paragraph{
					Attributes: map[string]interface{}{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: `**bold content**`},
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Paragraph"))
			})

			It("escaped bold text with double quote and more backslashes", func() {
				actualContent := `\\\**bold content**`
				expectedResult := types.Paragraph{
					Attributes: map[string]interface{}{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: `\**bold content**`},
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Paragraph"))
			})

			It("escaped bold text with unbalanced double quote", func() {
				actualContent := `\**bold content*`
				expectedResult := types.Paragraph{
					Attributes: map[string]interface{}{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: `**bold content*`},
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Paragraph"))
			})

			It("escaped bold text with unbalanced double quote and more backslashes", func() {
				actualContent := `\\\**bold content*`
				expectedResult := types.Paragraph{
					Attributes: map[string]interface{}{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: `\\**bold content*`},
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Paragraph"))
			})
		})

		Context("prevented Italic text substitution", func() {

			It("escaped italic text with simple quote", func() {
				actualContent := `\_italic content_`
				expectedResult := types.Paragraph{
					Attributes: map[string]interface{}{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "_italic content_"},
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Paragraph"))
			})

			It("escaped italic text with simple quote and more backslashes", func() {
				actualContent := `\\_italic content_`
				expectedResult := types.Paragraph{
					Attributes: map[string]interface{}{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: `\_italic content_`},
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Paragraph"))
			})

			It("escaped italic text with double quote", func() {
				actualContent := `\\__italic content__`
				expectedResult := types.Paragraph{
					Attributes: map[string]interface{}{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: `__italic content__`},
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Paragraph"))
			})

			It("escaped italic text with double quote and more backslashes", func() {
				actualContent := `\\\__italic content__`
				expectedResult := types.Paragraph{
					Attributes: map[string]interface{}{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: `\__italic content__`},
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Paragraph"))
			})

			It("escaped italic text with unbalanced double quote", func() {
				actualContent := `\__italic content_`
				expectedResult := types.Paragraph{
					Attributes: map[string]interface{}{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: `__italic content_`},
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Paragraph"))
			})

			It("escaped italic text with unbalanced double quote and more backslashes", func() {
				actualContent := `\\\__italic content_`
				expectedResult := types.Paragraph{
					Attributes: map[string]interface{}{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: `\\__italic content_`}, // only 1 backslash remove
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Paragraph"))
			})
		})

		Context("prevented Monospace text substitution", func() {

			It("escaped monospace text with simple quote", func() {
				actualContent := "\\`monospace content`"
				expectedResult := types.Paragraph{
					Attributes: map[string]interface{}{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "`monospace content`"},
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Paragraph"))
			})

			It("escaped monospace text with simple quote and more backslashes", func() {
				actualContent := "\\\\`monospace content`"
				expectedResult := types.Paragraph{
					Attributes: map[string]interface{}{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "\\`monospace content`"}, // only 1 backslash remove
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Paragraph"))
			})

			It("escaped monospace text with double quote", func() {
				actualContent := "\\\\``monospace content``"
				expectedResult := types.Paragraph{
					Attributes: map[string]interface{}{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "``monospace content``"},
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Paragraph"))
			})

			It("escaped monospace text with double quote and more backslashes", func() {
				actualContent := "\\\\\\``monospace content``" // 3 backslashes
				expectedResult := types.Paragraph{
					Attributes: map[string]interface{}{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "\\``monospace content``"},
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Paragraph"))
			})

			It("escaped monospace text with unbalanced double quote", func() {
				actualContent := "\\``monospace content`"
				expectedResult := types.Paragraph{
					Attributes: map[string]interface{}{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "``monospace content`"},
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Paragraph"))
			})

			It("escaped monospace text with unbalanced double quote and more backslashes", func() {
				actualContent := "\\\\\\``monospace content`" // 3 backslashes
				expectedResult := types.Paragraph{
					Attributes: map[string]interface{}{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "\\\\``monospace content`"}, // only 1 backslash remove
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Paragraph"))
			})
		})

		Context("include nested substitution", func() {
			It("escaped bold text with nested italic", func() {
				actualContent := `\*bold _and italic_ content*`
				expectedResult := types.Paragraph{
					Attributes: map[string]interface{}{},
					Lines: []types.InlineElements{
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
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Paragraph"))
			})
		})
	})

})
