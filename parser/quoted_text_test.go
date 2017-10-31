package parser_test

import (
	"github.com/bytesparadise/libasciidoc/parser"
	"github.com/bytesparadise/libasciidoc/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Quoted Texts", func() {

	Context("Quoted text with single punctuation", func() {

		It("bold text of 1 word", func() {
			actualContent := "*hello*"
			expectedResult := &types.QuotedText{
				Kind: types.Bold,
				Elements: []types.InlineElement{
					&types.StringElement{Content: "hello"},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("bold text of 2 words", func() {
			actualContent := "*bold    content*"
			expectedResult := &types.QuotedText{
				Kind: types.Bold,
				Elements: []types.InlineElement{
					&types.StringElement{Content: "bold    content"},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("bold text of 3 words", func() {
			actualContent := "*some bold content*"
			expectedResult := &types.QuotedText{
				Kind: types.Bold,
				Elements: []types.InlineElement{
					&types.StringElement{Content: "some bold content"},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("italic text with3 words", func() {
			actualContent := "_some italic content_"
			expectedResult := &types.QuotedText{
				Kind: types.Italic,
				Elements: []types.InlineElement{
					&types.StringElement{Content: "some italic content"},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("monospace text with3 words", func() {
			actualContent := "`some monospace content`"
			expectedResult := &types.QuotedText{
				Kind: types.Monospace,
				Elements: []types.InlineElement{
					&types.StringElement{Content: "some monospace content"},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("bold text within italic text", func() {
			actualContent := "_some *bold* content_"
			expectedResult := &types.QuotedText{
				Kind: types.Italic,
				Elements: []types.InlineElement{
					&types.StringElement{Content: "some "},
					&types.QuotedText{
						Kind: types.Bold,
						Elements: []types.InlineElement{
							&types.StringElement{Content: "bold"},
						},
					},
					&types.StringElement{Content: " content"},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("monospace text within bold text within italic quote", func() {
			actualContent := "*some _italic and `monospaced content`_*"
			expectedResult := &types.QuotedText{
				Kind: types.Bold,
				Elements: []types.InlineElement{
					&types.StringElement{Content: "some "},
					&types.QuotedText{
						Kind: types.Italic,
						Elements: []types.InlineElement{
							&types.StringElement{Content: "italic and "},
							&types.QuotedText{
								Kind: types.Monospace,
								Elements: []types.InlineElement{
									&types.StringElement{Content: "monospaced content"},
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
			expectedResult := &types.QuotedText{
				Kind: types.Italic,
				Elements: []types.InlineElement{
					&types.StringElement{Content: "some "},
					&types.QuotedText{
						Kind: types.Italic,
						Elements: []types.InlineElement{
							&types.StringElement{Content: "very italic"},
						},
					},
					&types.StringElement{Content: " content"},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})
	})

	Context("Quoted text with double punctuation", func() {

		It("bold text of 1 word", func() {
			actualContent := "**hello**"
			expectedResult := &types.QuotedText{
				Kind: types.Bold,
				Elements: []types.InlineElement{
					&types.StringElement{Content: "hello"},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("italic text with3 words", func() {
			actualContent := "__some italic content__"
			expectedResult := &types.QuotedText{
				Kind: types.Italic,
				Elements: []types.InlineElement{
					&types.StringElement{Content: "some italic content"},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("monospace text with3 words", func() {
			actualContent := "``some monospace content``"
			expectedResult := &types.QuotedText{
				Kind: types.Monospace,
				Elements: []types.InlineElement{
					&types.StringElement{Content: "some monospace content"},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("bold text within italic text", func() {
			actualContent := "__some *bold* content__"
			expectedResult := &types.QuotedText{
				Kind: types.Italic,
				Elements: []types.InlineElement{
					&types.StringElement{Content: "some "},
					&types.QuotedText{
						Kind: types.Bold,
						Elements: []types.InlineElement{
							&types.StringElement{Content: "bold"},
						},
					},
					&types.StringElement{Content: " content"},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("QuotedText"))
		})

		It("monospace text within bold text within italic quote", func() {
			actualContent := "**some _italic and ``monospaced content``_**"
			expectedResult := &types.QuotedText{
				Kind: types.Bold,
				Elements: []types.InlineElement{
					&types.StringElement{Content: "some "},
					&types.QuotedText{
						Kind: types.Italic,
						Elements: []types.InlineElement{
							&types.StringElement{Content: "italic and "},
							&types.QuotedText{
								Kind: types.Monospace,
								Elements: []types.InlineElement{
									&types.StringElement{Content: "monospaced content"},
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
			expectedResult := &types.Paragraph{
				Lines: []*types.InlineContent{
					&types.InlineContent{
						Elements: []types.InlineElement{
							&types.StringElement{Content: "a paragraph with "},
							&types.QuotedText{
								Kind: types.Bold,
								Elements: []types.InlineElement{
									&types.StringElement{Content: "some bold content"},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Paragraph"))
		})

		It("inline with invalid bold text1", func() {
			actualContent := "a paragraph with *some bold content"
			expectedResult := &types.Paragraph{
				Lines: []*types.InlineContent{
					&types.InlineContent{
						Elements: []types.InlineElement{
							&types.StringElement{Content: "a paragraph with *some bold content"},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Paragraph"))
		})

		It("inline with invalid bold text2", func() {
			actualContent := "a paragraph with *some bold content *"
			expectedResult := &types.Paragraph{
				Lines: []*types.InlineContent{
					&types.InlineContent{
						Elements: []types.InlineElement{
							&types.StringElement{Content: "a paragraph with *some bold content *"},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Paragraph"))
		})

		It("inline with invalid bold text3", func() {
			actualContent := "a paragraph with * some bold content*"
			expectedResult := &types.Paragraph{
				Lines: []*types.InlineContent{
					&types.InlineContent{
						Elements: []types.InlineElement{
							&types.StringElement{Content: "a paragraph with * some bold content*"},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Paragraph"))
		})

		It("invalid italic text within bold text", func() {
			actualContent := "some *bold and _italic content _ together*."
			expectedResult := &types.Paragraph{
				Lines: []*types.InlineContent{
					&types.InlineContent{
						Elements: []types.InlineElement{
							&types.StringElement{Content: "some "},
							&types.QuotedText{
								Kind: types.Bold,
								Elements: []types.InlineElement{
									&types.StringElement{Content: "bold and _italic content _ together"},
								},
							},
							&types.StringElement{Content: "."},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Paragraph"))
		})

		It("italic text within invalid bold text", func() {
			actualContent := "some *bold and _italic content_ together *."
			expectedResult := &types.Paragraph{
				Lines: []*types.InlineContent{
					&types.InlineContent{
						Elements: []types.InlineElement{
							&types.StringElement{Content: "some *bold and "},
							&types.QuotedText{
								Kind: types.Italic,
								Elements: []types.InlineElement{
									&types.StringElement{Content: "italic content"},
								},
							},
							&types.StringElement{Content: " together *."},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Paragraph"))
		})
	})

	Context("Nested quoted text", func() {

		It("italic text within bold text", func() {
			actualContent := "some *bold and _italic content_ together*."
			expectedResult := &types.Paragraph{
				Lines: []*types.InlineContent{
					&types.InlineContent{
						Elements: []types.InlineElement{
							&types.StringElement{Content: "some "},
							&types.QuotedText{
								Kind: types.Bold,
								Elements: []types.InlineElement{
									&types.StringElement{Content: "bold and "},
									&types.QuotedText{
										Kind: types.Italic,
										Elements: []types.InlineElement{
											&types.StringElement{Content: "italic content"},
										},
									},
									&types.StringElement{Content: " together"},
								},
							},
							&types.StringElement{Content: "."},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Paragraph"))
		})

		It("simple-quote bold within simple-quote bold text", func() {
			// here we don't allow for bold text within bold text, to comply with the existing implementations (asciidoc and asciidoctor)
			actualContent := "*some *nested bold* content*"
			expectedResult := &types.InlineContent{
				Elements: []types.InlineElement{
					&types.QuotedText{
						Kind: types.Bold,
						Elements: []types.InlineElement{
							&types.StringElement{Content: "some "},
							&types.QuotedText{
								Kind: types.Bold,
								Elements: []types.InlineElement{
									&types.StringElement{Content: "nested bold"},
								},
							},
							&types.StringElement{Content: " content"},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineContent"))
		})

		It("double-quote bold within double-quote bold text", func() {
			// here we don't allow for bold text within bold text, to comply with the existing implementations (asciidoc and asciidoctor)
			actualContent := "**some **nested bold** content**"
			expectedResult := &types.InlineContent{
				Elements: []types.InlineElement{
					&types.QuotedText{
						Kind: types.Bold,
						Elements: []types.InlineElement{
							&types.StringElement{Content: "some "},
							&types.QuotedText{
								Kind: types.Bold,
								Elements: []types.InlineElement{
									&types.StringElement{Content: "nested bold"},
								},
							},
							&types.StringElement{Content: " content"},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineContent"))
		})

		It("simple-quote bold within double-quote bold text", func() {
			// here we don't allow for bold text within bold text, to comply with the existing implementations (asciidoc and asciidoctor)
			actualContent := "**some *nested bold* content**"
			expectedResult := &types.InlineContent{
				Elements: []types.InlineElement{
					&types.QuotedText{
						Kind: types.Bold,
						Elements: []types.InlineElement{
							&types.StringElement{Content: "some "},
							&types.QuotedText{
								Kind: types.Bold,
								Elements: []types.InlineElement{
									&types.StringElement{Content: "nested bold"},
								},
							},
							&types.StringElement{Content: " content"},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineContent"))
		})

		It("double-quote bold within simple-quote bold text", func() {
			// here we don't allow for bold text within bold text, to comply with the existing implementations (asciidoc and asciidoctor)
			actualContent := "*some **nested bold** content*"
			expectedResult := &types.InlineContent{
				Elements: []types.InlineElement{
					&types.QuotedText{
						Kind: types.Bold,
						Elements: []types.InlineElement{
							&types.StringElement{Content: "some "},
							&types.QuotedText{
								Kind: types.Bold,
								Elements: []types.InlineElement{
									&types.StringElement{Content: "nested bold"},
								},
							},
							&types.StringElement{Content: " content"},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineContent"))
		})

		It("simple-quote italic within simple-quote italic text", func() {
			// here we don't allow for italic text within italic text, to comply with the existing implementations (asciidoc and asciidoctor)
			actualContent := "_some _nested italic_ content_"
			expectedResult := &types.InlineContent{
				Elements: []types.InlineElement{
					&types.QuotedText{
						Kind: types.Italic,
						Elements: []types.InlineElement{
							&types.StringElement{Content: "some "},
							&types.QuotedText{
								Kind: types.Italic,
								Elements: []types.InlineElement{
									&types.StringElement{Content: "nested italic"},
								},
							},
							&types.StringElement{Content: " content"},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineContent"))
		})

		It("double-quote italic within double-quote italic text", func() {
			// here we don't allow for italic text within italic text, to comply with the existing implementations (asciidoc and asciidoctor)
			actualContent := "__some __nested italic__ content__"
			expectedResult := &types.InlineContent{
				Elements: []types.InlineElement{
					&types.QuotedText{
						Kind: types.Italic,
						Elements: []types.InlineElement{
							&types.StringElement{Content: "some "},
							&types.QuotedText{
								Kind: types.Italic,
								Elements: []types.InlineElement{
									&types.StringElement{Content: "nested italic"},
								},
							},
							&types.StringElement{Content: " content"},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineContent"))
		})

		It("simple-quote italic within double-quote italic text", func() {
			// here we allow for italic text within italic text, to comply with the existing implementations (asciidoc and asciidoctor)
			actualContent := "_some __nested italic__ content_"
			expectedResult := &types.InlineContent{
				Elements: []types.InlineElement{
					&types.QuotedText{
						Kind: types.Italic,
						Elements: []types.InlineElement{
							&types.StringElement{Content: "some "},
							&types.QuotedText{
								Kind: types.Italic,
								Elements: []types.InlineElement{
									&types.StringElement{Content: "nested italic"},
								},
							},
							&types.StringElement{Content: " content"},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineContent"))
		})

		It("double-quote italic within simple-quote italic text", func() {
			// here we allow for italic text within italic text, to comply with the existing implementations (asciidoc and asciidoctor)
			actualContent := "_some __nested italic__ content_"
			expectedResult := &types.InlineContent{
				Elements: []types.InlineElement{
					&types.QuotedText{
						Kind: types.Italic,
						Elements: []types.InlineElement{
							&types.StringElement{Content: "some "},
							&types.QuotedText{
								Kind: types.Italic,
								Elements: []types.InlineElement{
									&types.StringElement{Content: "nested italic"},
								},
							},
							&types.StringElement{Content: " content"},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineContent"))
		})

		It("simple-quote monospace within simple-quote monospace text", func() {
			// here we don't allow for monospace text within monospace text, to comply with the existing implementations (asciidoc and asciidoctor)
			actualContent := "`some `nested monospace` content`"
			expectedResult := &types.InlineContent{
				Elements: []types.InlineElement{
					&types.QuotedText{
						Kind: types.Monospace,
						Elements: []types.InlineElement{
							&types.StringElement{Content: "some "},
							&types.QuotedText{
								Kind: types.Monospace,
								Elements: []types.InlineElement{
									&types.StringElement{Content: "nested monospace"},
								},
							},
							&types.StringElement{Content: " content"},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineContent"))
		})

		It("double-quote monospace within double-quote monospace text", func() {
			// here we don't allow for monospace text within monospace text, to comply with the existing implementations (asciidoc and asciidoctor)
			actualContent := "``some ``nested monospace`` content``"
			expectedResult := &types.InlineContent{
				Elements: []types.InlineElement{
					&types.QuotedText{
						Kind: types.Monospace,
						Elements: []types.InlineElement{
							&types.StringElement{Content: "some "},
							&types.QuotedText{
								Kind: types.Monospace,
								Elements: []types.InlineElement{
									&types.StringElement{Content: "nested monospace"},
								},
							},
							&types.StringElement{Content: " content"},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineContent"))
		})

		It("simple-quote monospace within double-quote monospace text", func() {
			// here we allow for monospace text within monospace text, to comply with the existing implementations (asciidoc and asciidoctor)
			actualContent := "`some ``nested monospace`` content`"
			expectedResult := &types.InlineContent{
				Elements: []types.InlineElement{
					&types.QuotedText{
						Kind: types.Monospace,
						Elements: []types.InlineElement{
							&types.StringElement{Content: "some "},
							&types.QuotedText{
								Kind: types.Monospace,
								Elements: []types.InlineElement{
									&types.StringElement{Content: "nested monospace"},
								},
							},
							&types.StringElement{Content: " content"},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineContent"))
		})

		It("double-quote monospace within simple-quote monospace text", func() {
			// here we allow for monospace text within monospace text, to comply with the existing implementations (asciidoc and asciidoctor)
			actualContent := "`some ``nested monospace`` content`"
			expectedResult := &types.InlineContent{
				Elements: []types.InlineElement{
					&types.QuotedText{
						Kind: types.Monospace,
						Elements: []types.InlineElement{
							&types.StringElement{Content: "some "},
							&types.QuotedText{
								Kind: types.Monospace,
								Elements: []types.InlineElement{
									&types.StringElement{Content: "nested monospace"},
								},
							},
							&types.StringElement{Content: " content"},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineContent"))
		})

	})

	Context("Unbalanced quoted text", func() {

		It("unbalanced bold text - extra on left", func() {
			actualContent := "**some bold content*"
			expectedResult := &types.InlineContent{
				Elements: []types.InlineElement{
					&types.QuotedText{
						Kind: types.Bold,
						Elements: []types.InlineElement{
							&types.StringElement{Content: "*some bold content"},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineContent"))
		})

		It("unbalanced bold text - extra on right", func() {
			actualContent := "*some bold content**"
			expectedResult := &types.InlineContent{
				Elements: []types.InlineElement{
					&types.QuotedText{
						Kind: types.Bold,
						Elements: []types.InlineElement{
							&types.StringElement{Content: "some bold content"},
						},
					},
					&types.StringElement{Content: "*"},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineContent"))
		})

		It("unbalanced italic text - extra on left", func() {
			actualContent := "__some italic content_"
			expectedResult := &types.InlineContent{
				Elements: []types.InlineElement{
					&types.QuotedText{
						Kind: types.Italic,
						Elements: []types.InlineElement{
							&types.StringElement{Content: "_some italic content"},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineContent"))
		})

		It("unbalanced italic text - extra on right", func() {
			actualContent := "_some italic content__"
			expectedResult := &types.InlineContent{
				Elements: []types.InlineElement{
					&types.QuotedText{
						Kind: types.Italic,
						Elements: []types.InlineElement{
							&types.StringElement{Content: "some italic content"},
						},
					},
					&types.StringElement{Content: "_"},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineContent"))
		})

		It("unbalanced monospace text - extra on left", func() {
			actualContent := "``some monospace content`"
			expectedResult := &types.InlineContent{
				Elements: []types.InlineElement{
					&types.QuotedText{
						Kind: types.Monospace,
						Elements: []types.InlineElement{
							&types.StringElement{Content: "`some monospace content"},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineContent"))
		})

		It("unbalanced monospace text - extra on right", func() {
			actualContent := "`some monospace content``"
			expectedResult := &types.InlineContent{
				Elements: []types.InlineElement{
					&types.QuotedText{
						Kind: types.Monospace,
						Elements: []types.InlineElement{
							&types.StringElement{Content: "some monospace content"},
						},
					},
					&types.StringElement{Content: "`"},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("InlineContent"))
		})

		It("inline with unbalanced bold text", func() {
			actualContent := "a paragraph with *some bold content"
			expectedResult := &types.Paragraph{
				Lines: []*types.InlineContent{
					&types.InlineContent{
						Elements: []types.InlineElement{
							&types.StringElement{Content: "a paragraph with *some bold content"},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Paragraph"))
		})

	})

})
