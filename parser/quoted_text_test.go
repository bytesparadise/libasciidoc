package parser_test

import (
	"github.com/bytesparadise/libasciidoc/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Quoted Texts", func() {

	It("bold text of 1 word", func() {
		actualContent := "*hello*"
		expectedDocument := &types.Document{
			Attributes: &types.DocumentAttributes{},
			Elements: []types.DocElement{
				&types.Paragraph{
					Lines: []*types.InlineContent{
						&types.InlineContent{
							Elements: []types.InlineElement{
								&types.QuotedText{
									Kind: types.Bold,
									Elements: []types.InlineElement{
										&types.StringElement{Content: "hello"},
									},
								},
							},
						},
					},
				},
			},
		}
		verify(GinkgoT(), expectedDocument, actualContent)
	})

	It("bold text of 2 words", func() {
		actualContent := "*bold    content*"
		expectedDocument := &types.Document{
			Attributes: &types.DocumentAttributes{},
			Elements: []types.DocElement{
				&types.Paragraph{
					Lines: []*types.InlineContent{
						&types.InlineContent{
							Elements: []types.InlineElement{
								&types.QuotedText{
									Kind: types.Bold,
									Elements: []types.InlineElement{
										&types.StringElement{Content: "bold    content"},
									},
								},
							},
						},
					},
				},
			},
		}
		verify(GinkgoT(), expectedDocument, actualContent)
	})

	It("bold text of 3 words", func() {
		actualContent := "*some bold content*"
		expectedDocument := &types.Document{
			Attributes: &types.DocumentAttributes{},
			Elements: []types.DocElement{
				&types.Paragraph{
					Lines: []*types.InlineContent{
						&types.InlineContent{
							Elements: []types.InlineElement{
								&types.QuotedText{
									Kind: types.Bold,
									Elements: []types.InlineElement{
										&types.StringElement{Content: "some bold content"},
									},
								},
							},
						},
					},
				},
			},
		}
		verify(GinkgoT(), expectedDocument, actualContent)
	})

	It("inline with bold text", func() {
		actualContent := "a paragraph with *some bold content*"
		expectedDocument := &types.Document{
			Attributes: &types.DocumentAttributes{},
			Elements: []types.DocElement{
				&types.Paragraph{
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
				},
			},
		}
		verify(GinkgoT(), expectedDocument, actualContent)
	})

	It("inline with invalid bold text1", func() {
		actualContent := "a paragraph with *some bold content"
		expectedDocument := &types.Document{
			Attributes: &types.DocumentAttributes{},
			Elements: []types.DocElement{
				&types.Paragraph{
					Lines: []*types.InlineContent{
						&types.InlineContent{
							Elements: []types.InlineElement{
								&types.StringElement{Content: "a paragraph with *some bold content"},
							},
						},
					},
				},
			},
		}
		verify(GinkgoT(), expectedDocument, actualContent)
	})

	It("inline with invalid bold text2", func() {
		actualContent := "a paragraph with *some bold content *"
		expectedDocument := &types.Document{
			Attributes: &types.DocumentAttributes{},
			Elements: []types.DocElement{
				&types.Paragraph{
					Lines: []*types.InlineContent{
						&types.InlineContent{
							Elements: []types.InlineElement{
								&types.StringElement{Content: "a paragraph with *some bold content *"},
							},
						},
					},
				},
			},
		}
		verify(GinkgoT(), expectedDocument, actualContent)
	})

	It("inline with invalid bold text3", func() {
		actualContent := "a paragraph with * some bold content*"
		expectedDocument := &types.Document{
			Attributes: &types.DocumentAttributes{},
			Elements: []types.DocElement{
				&types.Paragraph{
					Lines: []*types.InlineContent{
						&types.InlineContent{
							Elements: []types.InlineElement{
								&types.StringElement{Content: "a paragraph with * some bold content*"},
							},
						},
					},
				},
			},
		}
		verify(GinkgoT(), expectedDocument, actualContent)
	})

	It("italic text with3 words", func() {
		actualContent := "_some italic content_"
		expectedDocument := &types.Document{
			Attributes: &types.DocumentAttributes{},
			Elements: []types.DocElement{
				&types.Paragraph{
					Lines: []*types.InlineContent{
						&types.InlineContent{
							Elements: []types.InlineElement{
								&types.QuotedText{
									Kind: types.Italic,
									Elements: []types.InlineElement{
										&types.StringElement{Content: "some italic content"},
									},
								},
							},
						},
					},
				},
			},
		}
		verify(GinkgoT(), expectedDocument, actualContent)
	})

	It("monospace text with3 words", func() {
		actualContent := "`some monospace content`"
		expectedDocument := &types.Document{
			Attributes: &types.DocumentAttributes{},
			Elements: []types.DocElement{
				&types.Paragraph{
					Lines: []*types.InlineContent{
						&types.InlineContent{
							Elements: []types.InlineElement{
								&types.QuotedText{
									Kind: types.Monospace,
									Elements: []types.InlineElement{
										&types.StringElement{Content: "some monospace content"},
									},
								},
							},
						},
					},
				},
			},
		}
		verify(GinkgoT(), expectedDocument, actualContent)
	})

	It("italic text within bold text", func() {
		actualContent := "some *bold and _italic content_ together*."
		expectedDocument := &types.Document{
			Attributes: &types.DocumentAttributes{},
			Elements: []types.DocElement{
				&types.Paragraph{
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
				},
			},
		}
		verify(GinkgoT(), expectedDocument, actualContent)
	})

	It("invalid italic text within bold text", func() {
		actualContent := "some *bold and _italic content _ together*."
		expectedDocument := &types.Document{
			Attributes: &types.DocumentAttributes{},
			Elements: []types.DocElement{
				&types.Paragraph{
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
				},
			},
		}
		verify(GinkgoT(), expectedDocument, actualContent)
	})

	It("italic text within invalid bold text", func() {
		actualContent := "some *bold and _italic content_ together *."
		expectedDocument := &types.Document{
			Attributes: &types.DocumentAttributes{},
			Elements: []types.DocElement{
				&types.Paragraph{
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
				},
			},
		}
		verify(GinkgoT(), expectedDocument, actualContent)
	})

	It("bold text within italic text", func() {
		actualContent := "_some *bold* content_"
		expectedDocument := &types.Document{
			Attributes: &types.DocumentAttributes{},
			Elements: []types.DocElement{
				&types.Paragraph{
					Lines: []*types.InlineContent{
						&types.InlineContent{
							Elements: []types.InlineElement{
								&types.QuotedText{
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
								},
							},
						},
					},
				},
			},
		}
		verify(GinkgoT(), expectedDocument, actualContent)
	})

	It("monospace text within bold text within italic quote", func() {
		actualContent := "*some _italic and `monospaced content`_*"
		expectedDocument := &types.Document{
			Attributes: &types.DocumentAttributes{},
			Elements: []types.DocElement{
				&types.Paragraph{
					Lines: []*types.InlineContent{
						&types.InlineContent{
							Elements: []types.InlineElement{
								&types.QuotedText{
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
								},
							},
						},
					},
				},
			},
		}
		verify(GinkgoT(), expectedDocument, actualContent)
	})

	It("italic text within italic text", func() {
		actualContent := "_some _very italic_ content_"
		expectedDocument := &types.Document{
			Attributes: &types.DocumentAttributes{},
			Elements: []types.DocElement{
				&types.Paragraph{
					Lines: []*types.InlineContent{
						&types.InlineContent{
							Elements: []types.InlineElement{
								&types.QuotedText{
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
								},
							},
						},
					},
				},
			},
		}
		verify(GinkgoT(), expectedDocument, actualContent)
	})

})
