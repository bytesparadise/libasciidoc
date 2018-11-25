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
				Attributes: types.ElementAttributes{
					types.AttrKind: types.Bold,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "hello"},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("bold text with 2 words", func() {
			actualContent := "*bold    content*"
			expectedResult := types.QuotedText{
				Attributes: types.ElementAttributes{
					types.AttrKind: types.Bold,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "bold    content"},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("bold text with 3 words", func() {
			actualContent := "*some bold content*"
			expectedResult := types.QuotedText{
				Attributes: types.ElementAttributes{
					types.AttrKind: types.Bold,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "some bold content"},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("italic text with 3 words", func() {
			actualContent := "_some italic content_"
			expectedResult := types.QuotedText{
				Attributes: types.ElementAttributes{
					types.AttrKind: types.Italic,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "some italic content"},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("monospace text with 3 words", func() {
			actualContent := "`some monospace content`"
			expectedResult := types.QuotedText{
				Attributes: types.ElementAttributes{
					types.AttrKind: types.Monospace,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "some monospace content"},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("subscript text with 3 words", func() {
			actualContent := "~some subscript content~"
			expectedResult := types.QuotedText{
				Attributes: types.ElementAttributes{
					types.AttrKind: types.Subscript,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "some subscript content"},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("superscript text with 3 words", func() {
			actualContent := "^some superscript content^"
			expectedResult := types.QuotedText{
				Attributes: types.ElementAttributes{
					types.AttrKind: types.Superscript,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "some superscript content"},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("bold text within italic text", func() {
			actualContent := "_some *bold* content_"
			expectedResult := types.QuotedText{
				Attributes: types.ElementAttributes{
					types.AttrKind: types.Italic,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "some "},
					types.QuotedText{
						Attributes: types.ElementAttributes{
							types.AttrKind: types.Bold,
						},
						Elements: types.InlineElements{
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
				Attributes: types.ElementAttributes{
					types.AttrKind: types.Bold,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "some "},
					types.QuotedText{
						Attributes: types.ElementAttributes{
							types.AttrKind: types.Italic,
						},
						Elements: types.InlineElements{
							types.StringElement{Content: "italic and "},
							types.QuotedText{
								Attributes: types.ElementAttributes{
									types.AttrKind: types.Monospace,
								},
								Elements: types.InlineElements{
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
				Attributes: types.ElementAttributes{
					types.AttrKind: types.Italic,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "some "},
					types.QuotedText{
						Attributes: types.ElementAttributes{
							types.AttrKind: types.Italic,
						},
						Elements: types.InlineElements{
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
				Attributes: types.ElementAttributes{
					types.AttrKind: types.Bold,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "hello"},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("italic text with 3 words", func() {
			actualContent := "__some italic content__"
			expectedResult := types.QuotedText{
				Attributes: types.ElementAttributes{
					types.AttrKind: types.Italic,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "some italic content"},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("monospace text with 3 words", func() {
			actualContent := "``some monospace content``"
			expectedResult := types.QuotedText{
				Attributes: types.ElementAttributes{
					types.AttrKind: types.Monospace,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "some monospace content"},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("subscript text with 3 words", func() {
			actualContent := "~~some subscript content~~"
			expectedResult := types.QuotedText{
				Attributes: types.ElementAttributes{
					types.AttrKind: types.Subscript,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "some subscript content"},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("superscript text attached", func() {
			actualContent := "O~2~ is a molecule"
			expectedResult := types.InlineElements{
				types.StringElement{Content: "O"},
				types.QuotedText{
					Attributes: types.ElementAttributes{
						types.AttrKind: types.Subscript,
					},
					Elements: types.InlineElements{
						types.StringElement{Content: "2"},
					},
				},
				types.StringElement{Content: " is a molecule"},
			}

			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("superscript text with 3 words", func() {
			actualContent := "^^some superscript content^^"
			expectedResult := types.QuotedText{
				Attributes: types.ElementAttributes{
					types.AttrKind: types.Superscript,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "some superscript content"},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("superscript text attached", func() {
			actualContent := "M^me^ White"
			expectedResult := types.InlineElements{
				types.StringElement{Content: "M"},
				types.QuotedText{
					Attributes: types.ElementAttributes{
						types.AttrKind: types.Superscript,
					},
					Elements: types.InlineElements{
						types.StringElement{Content: "me"},
					},
				},
				types.StringElement{Content: " White"},
			}

			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("superscript text within italic text", func() {
			actualContent := "__some ^superscript^ content__"
			expectedResult := types.QuotedText{
				Attributes: types.ElementAttributes{
					types.AttrKind: types.Italic,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "some "},
					types.QuotedText{
						Attributes: types.ElementAttributes{
							types.AttrKind: types.Superscript,
						},
						Elements: types.InlineElements{
							types.StringElement{Content: "superscript"},
						},
					},
					types.StringElement{Content: " content"},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("superscript text within italic text within bold quote", func() {
			actualContent := "**some _italic and ^^superscript content^^_**"
			expectedResult := types.QuotedText{
				Attributes: types.ElementAttributes{
					types.AttrKind: types.Bold,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "some "},
					types.QuotedText{
						Attributes: types.ElementAttributes{
							types.AttrKind: types.Italic,
						},
						Elements: types.InlineElements{
							types.StringElement{Content: "italic and "},
							types.QuotedText{
								Attributes: types.ElementAttributes{
									types.AttrKind: types.Superscript,
								},
								Elements: types.InlineElements{
									types.StringElement{Content: "superscript content"},
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

		It("inline content with bold text", func() {
			actualContent := "a paragraph with *some bold content*"
			expectedResult := types.InlineElements{
				types.StringElement{Content: "a paragraph with "},
				types.QuotedText{
					Attributes: types.ElementAttributes{
						types.AttrKind: types.Bold,
					},
					Elements: types.InlineElements{
						types.StringElement{Content: "some bold content"},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("inline content with invalid bold text - use case 1", func() {
			actualContent := "a paragraph with *some bold content"
			expectedResult := types.InlineElements{
				types.StringElement{Content: "a paragraph with *some bold content"},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("inline content with invalid bold text - use case 2", func() {
			actualContent := "a paragraph with *some bold content *"
			expectedResult := types.InlineElements{
				types.StringElement{Content: "a paragraph with *some bold content *"},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("inline content with invalid bold text - use case 3", func() {
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
					Attributes: types.ElementAttributes{
						types.AttrKind: types.Bold,
					},
					Elements: types.InlineElements{
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
					Attributes: types.ElementAttributes{
						types.AttrKind: types.Italic,
					},
					Elements: types.InlineElements{
						types.StringElement{Content: "italic content"},
					},
				},
				types.StringElement{Content: " together *."},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("inline content with invalid subscript text - use case 1", func() {
			actualContent := "a paragraph with ~some subscript content"
			expectedResult := types.InlineElements{
				types.StringElement{Content: "a paragraph with ~some subscript content"},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("inline content with invalid subscript text - use case 2", func() {
			actualContent := "a paragraph with ~some subscript content ~"
			expectedResult := types.InlineElements{
				types.StringElement{Content: "a paragraph with ~some subscript content ~"},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("inline content with invalid subscript text - use case 3", func() {
			actualContent := "a paragraph with ~ some subscript content~"
			expectedResult := types.InlineElements{
				types.StringElement{Content: "a paragraph with ~ some subscript content~"},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("inline content with invalid superscript text - use case 1", func() {
			actualContent := "a paragraph with ^some superscript content"
			expectedResult := types.InlineElements{
				types.StringElement{Content: "a paragraph with ^some superscript content"},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("inline content with invalid superscript text - use case 2", func() {
			actualContent := "a paragraph with ^some superscript content ^"
			expectedResult := types.InlineElements{
				types.StringElement{Content: "a paragraph with ^some superscript content ^"},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})

		It("inline content with invalid superscript text - use case 3", func() {
			actualContent := "a paragraph with ^ some superscript content^"
			expectedResult := types.InlineElements{
				types.StringElement{Content: "a paragraph with ^ some superscript content^"},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
		})
	})

	Context("nested quoted text", func() {

		It("italic text within bold text", func() {
			actualContent := "some *bold and _italic content_ together*."
			expectedResult := types.InlineElements{
				types.StringElement{Content: "some "},
				types.QuotedText{
					Attributes: types.ElementAttributes{
						types.AttrKind: types.Bold,
					},
					Elements: types.InlineElements{
						types.StringElement{Content: "bold and "},
						types.QuotedText{
							Attributes: types.ElementAttributes{
								types.AttrKind: types.Italic,
							},
							Elements: types.InlineElements{
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
					Attributes: types.ElementAttributes{
						types.AttrKind: types.Bold,
					},
					Elements: types.InlineElements{
						types.StringElement{Content: "some "},
						types.QuotedText{
							Attributes: types.ElementAttributes{
								types.AttrKind: types.Bold,
							},
							Elements: types.InlineElements{
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
					Attributes: types.ElementAttributes{
						types.AttrKind: types.Bold,
					},
					Elements: types.InlineElements{
						types.StringElement{Content: "some "},
						types.QuotedText{
							Attributes: types.ElementAttributes{
								types.AttrKind: types.Bold,
							},
							Elements: types.InlineElements{
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
					Attributes: types.ElementAttributes{
						types.AttrKind: types.Bold,
					},
					Elements: types.InlineElements{
						types.StringElement{Content: "some "},
						types.QuotedText{
							Attributes: types.ElementAttributes{
								types.AttrKind: types.Bold,
							},
							Elements: types.InlineElements{
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
					Attributes: types.ElementAttributes{
						types.AttrKind: types.Bold,
					},
					Elements: types.InlineElements{
						types.StringElement{Content: "some "},
						types.QuotedText{
							Attributes: types.ElementAttributes{
								types.AttrKind: types.Bold,
							},
							Elements: types.InlineElements{
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
					Attributes: types.ElementAttributes{
						types.AttrKind: types.Italic,
					},
					Elements: types.InlineElements{
						types.StringElement{Content: "some "},
						types.QuotedText{
							Attributes: types.ElementAttributes{
								types.AttrKind: types.Italic,
							},
							Elements: types.InlineElements{
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
					Attributes: types.ElementAttributes{
						types.AttrKind: types.Italic,
					},
					Elements: types.InlineElements{
						types.StringElement{Content: "some "},
						types.QuotedText{
							Attributes: types.ElementAttributes{
								types.AttrKind: types.Italic,
							},
							Elements: types.InlineElements{
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
					Attributes: types.ElementAttributes{
						types.AttrKind: types.Italic,
					},
					Elements: types.InlineElements{
						types.StringElement{Content: "some "},
						types.QuotedText{
							Attributes: types.ElementAttributes{
								types.AttrKind: types.Italic,
							},
							Elements: types.InlineElements{
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
					Attributes: types.ElementAttributes{
						types.AttrKind: types.Italic,
					},
					Elements: types.InlineElements{
						types.StringElement{Content: "some "},
						types.QuotedText{
							Attributes: types.ElementAttributes{
								types.AttrKind: types.Italic,
							},
							Elements: types.InlineElements{
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
					Attributes: types.ElementAttributes{
						types.AttrKind: types.Monospace,
					},
					Elements: types.InlineElements{
						types.StringElement{Content: "some "},
						types.QuotedText{
							Attributes: types.ElementAttributes{
								types.AttrKind: types.Monospace,
							},
							Elements: types.InlineElements{
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
					Attributes: types.ElementAttributes{
						types.AttrKind: types.Monospace,
					},
					Elements: types.InlineElements{
						types.StringElement{Content: "some "},
						types.QuotedText{
							Attributes: types.ElementAttributes{
								types.AttrKind: types.Monospace,
							},
							Elements: types.InlineElements{
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
					Attributes: types.ElementAttributes{
						types.AttrKind: types.Monospace,
					},
					Elements: types.InlineElements{
						types.StringElement{Content: "some "},
						types.QuotedText{
							Attributes: types.ElementAttributes{
								types.AttrKind: types.Monospace,
							},
							Elements: types.InlineElements{
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
					Attributes: types.ElementAttributes{
						types.AttrKind: types.Monospace,
					},
					Elements: types.InlineElements{
						types.StringElement{Content: "some "},
						types.QuotedText{
							Attributes: types.ElementAttributes{
								types.AttrKind: types.Monospace,
							},
							Elements: types.InlineElements{
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
						Attributes: types.ElementAttributes{
							types.AttrKind: types.Bold,
						},
						Elements: types.InlineElements{
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
						Attributes: types.ElementAttributes{
							types.AttrKind: types.Bold,
						},
						Elements: types.InlineElements{
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
						Attributes: types.ElementAttributes{
							types.AttrKind: types.Italic,
						},
						Elements: types.InlineElements{
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
						Attributes: types.ElementAttributes{
							types.AttrKind: types.Italic,
						},
						Elements: types.InlineElements{
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
						Attributes: types.ElementAttributes{
							types.AttrKind: types.Monospace,
						},
						Elements: types.InlineElements{
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
						Attributes: types.ElementAttributes{
							types.AttrKind: types.Monospace,
						},
						Elements: types.InlineElements{
							types.StringElement{Content: "some monospace content"},
						},
					},
					types.StringElement{Content: "`"},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineElements"))
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
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

	})

	Context("prevented substitution", func() {

		Context("prevented bold text substitution", func() {

			Context("without nested quoted text", func() {

				It("escaped bold text with simple quote", func() {
					actualContent := `\*bold content*`
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "*bold content*"},
							},
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped bold text with simple quote and more backslashes", func() {
					actualContent := `\\*bold content*`
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `\*bold content*`},
							},
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
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
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
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
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
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
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
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
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
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
									Attributes: types.ElementAttributes{
										types.AttrKind: types.Italic,
									},
									Elements: types.InlineElements{
										types.StringElement{Content: "italic content"},
									},
								},
								types.StringElement{Content: "*"},
							},
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped bold text with unbalanced double quote and nested italic test", func() {
					actualContent := `\**_italic content_*`
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "**"},
								types.QuotedText{
									Attributes: types.ElementAttributes{
										types.AttrKind: types.Italic,
									},
									Elements: types.InlineElements{
										types.StringElement{Content: "italic content"},
									},
								},
								types.StringElement{Content: "*"},
							},
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped bold text with nested italic", func() {
					actualContent := `\*bold _and italic_ content*`
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "*bold "},
								types.QuotedText{
									Attributes: types.ElementAttributes{
										types.AttrKind: types.Italic,
									},
									Elements: types.InlineElements{
										types.StringElement{Content: "and italic"},
									},
								},
								types.StringElement{Content: " content*"},
							},
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})
			})

		})

		Context("prevented italic text substitution", func() {

			Context("without nested quoted text", func() {

				It("escaped italic text with simple quote", func() {
					actualContent := `\_italic content_`
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "_italic content_"},
							},
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped italic text with simple quote and more backslashes", func() {
					actualContent := `\\_italic content_`
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `\_italic content_`},
							},
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped italic text with double quote", func() {
					actualContent := `\\__italic content__`
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `__italic content__`},
							},
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped italic text with double quote and more backslashes", func() {
					actualContent := `\\\__italic content__`
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `\__italic content__`},
							},
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
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
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
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
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
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
									Attributes: types.ElementAttributes{
										types.AttrKind: types.Monospace,
									},
									Elements: types.InlineElements{
										types.StringElement{Content: "monospace content"},
									},
								},
								types.StringElement{Content: "_"},
							},
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped italic text with unbalanced double quote and nested bold test", func() {
					actualContent := `\__*bold content*_`
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "__"},
								types.QuotedText{
									Attributes: types.ElementAttributes{
										types.AttrKind: types.Bold,
									},
									Elements: types.InlineElements{
										types.StringElement{Content: "bold content"},
									},
								},
								types.StringElement{Content: "_"},
							},
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped italic text with nested bold text", func() {
					actualContent := `\_italic *and bold* content_`
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "_italic "},
								types.QuotedText{
									Attributes: types.ElementAttributes{
										types.AttrKind: types.Bold,
									},
									Elements: types.InlineElements{
										types.StringElement{Content: "and bold"},
									},
								},
								types.StringElement{Content: " content_"},
							},
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})
			})
		})

		Context("prevented monospace text substitution", func() {

			Context("without nested quoted text", func() {

				It("escaped monospace text with simple quote", func() {
					actualContent := `\` + "`monospace content`"
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "`monospace content`"}, // backslash removed
							},
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped monospace text with simple quote and more backslashes", func() {
					actualContent := `\\` + "`monospace content`"
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `\` + "`monospace content`"}, // only 1 backslash removed
							},
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
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
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
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
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
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
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
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
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
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
									Attributes: types.ElementAttributes{
										types.AttrKind: types.Bold,
									},
									Elements: types.InlineElements{
										types.StringElement{Content: "bold content"},
									},
								},
								types.StringElement{Content: "`"},
							},
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped monospace text with unbalanced double backquote and nested bold test", func() {
					actualContent := `\` + "``*bold content*`"
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "``"},
								types.QuotedText{
									Attributes: types.ElementAttributes{
										types.AttrKind: types.Bold,
									},
									Elements: types.InlineElements{
										types.StringElement{Content: "bold content"},
									},
								},
								types.StringElement{Content: "`"},
							},
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped monospace text with nested bold text", func() {
					actualContent := `\` + "`monospace *and bold* content`"
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "`monospace "},
								types.QuotedText{
									Attributes: types.ElementAttributes{
										types.AttrKind: types.Bold,
									},
									Elements: types.InlineElements{
										types.StringElement{Content: "and bold"},
									},
								},
								types.StringElement{Content: " content`"},
							},
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})
			})
		})

		Context("prevented subscript text substitution", func() {

			Context("without nested quoted text", func() {

				It("escaped subscript text with simple quote", func() {
					actualContent := `\~subscript content~`
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "~subscript content~"},
							},
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped subscript text with simple quote and more backslashes", func() {
					actualContent := `\\~subscript content~`
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `\~subscript content~`}, // only 1 backslash removed
							},
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped subscript text with double quote", func() {
					actualContent := `\\~subscript content~~`
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `\~subscript content~~`},
							},
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped subscript text with double quote and more backslashes", func() {
					actualContent := `\\\~~subscript content~~` // 3 backslashes
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `\~~subscript content~~`}, // 2 backslashes removed
							},
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped subscript text with unbalanced double quote", func() {
					actualContent := `\~~subscript content~`
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "~~subscript content~"},
							},
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped subscript text with unbalanced double quote and more backslashes", func() {
					actualContent := `\\\~~subscript content~` // 3 backslashes
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `\\~~subscript content~`}, // 2 backslashes removed
							},
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})
			})

			Context("with nested quoted text", func() {

				It("escaped subscript text with nested bold text", func() {
					actualContent := `\~*bold content*~`
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "~"},
								types.QuotedText{
									Attributes: types.ElementAttributes{
										types.AttrKind: types.Bold,
									},
									Elements: types.InlineElements{
										types.StringElement{Content: "bold content"},
									},
								},
								types.StringElement{Content: "~"},
							},
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped subscript text with unbalanced double backquote and nested bold test", func() {
					actualContent := `\~~*bold content*~`
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "~~"},
								types.QuotedText{
									Attributes: types.ElementAttributes{
										types.AttrKind: types.Bold,
									},
									Elements: types.InlineElements{
										types.StringElement{Content: "bold content"},
									},
								},
								types.StringElement{Content: "~"},
							},
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped subscript text with nested bold text", func() {
					actualContent := `\~subscript *and bold* content~`
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "~subscript "},
								types.QuotedText{
									Attributes: types.ElementAttributes{
										types.AttrKind: types.Bold,
									},
									Elements: types.InlineElements{
										types.StringElement{Content: "and bold"},
									},
								},
								types.StringElement{Content: " content~"},
							},
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})
			})
		})

		Context("prevented superscript text substitution", func() {

			Context("without nested quoted text", func() {

				It("escaped superscript text with simple quote", func() {
					actualContent := `\^superscript content^`
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "^superscript content^"},
							},
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped superscript text with simple quote and more backslashes", func() {
					actualContent := `\\^superscript content^`
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `\^superscript content^`}, // only 1 backslash removed
							},
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped superscript text with double quote", func() {
					actualContent := `\\^^superscript content^^`
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `^^superscript content^^`}, // 2 backslashes removed
							},
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped superscript text with double quote and more backslashes", func() {
					actualContent := `\\\` + "^^superscript content^^" // 3 backslashes
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `\^^superscript content^^`}, // 2 backslashes removed
							},
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped superscript text with unbalanced double quote", func() {
					actualContent := `\^^superscript content^`
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "^^superscript content^"},
							},
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped superscript text with unbalanced double quote and more backslashes", func() {
					actualContent := `\\\^^superscript content^` // 3 backslashes
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: `\\^^superscript content^`}, // only 1 backslash removed
							},
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})
			})

			Context("with nested quoted text", func() {

				It("escaped superscript text with nested bold text", func() {
					actualContent := `\^*bold content*^`
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "^"},
								types.QuotedText{
									Attributes: types.ElementAttributes{
										types.AttrKind: types.Bold,
									},
									Elements: types.InlineElements{
										types.StringElement{Content: "bold content"},
									},
								},
								types.StringElement{Content: "^"},
							},
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped superscript text with unbalanced double backquote and nested bold test", func() {
					actualContent := `\^^*bold content*^`
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "^^"},
								types.QuotedText{
									Attributes: types.ElementAttributes{
										types.AttrKind: types.Bold,
									},
									Elements: types.InlineElements{
										types.StringElement{Content: "bold content"},
									},
								},
								types.StringElement{Content: "^"},
							},
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})

				It("escaped superscript text with nested bold text", func() {
					actualContent := `\^superscript *and bold* content^`
					expectedResult := types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "^superscript "},
								types.QuotedText{
									Attributes: types.ElementAttributes{
										types.AttrKind: types.Bold,
									},
									Elements: types.InlineElements{
										types.StringElement{Content: "and bold"},
									},
								},
								types.StringElement{Content: " content^"},
							},
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
				})
			})
		})
	})
})
