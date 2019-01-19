package parser_test

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("footnotes", func() {

	BeforeEach(func() {
		types.ResetFootnoteSequence()
	})

	Context("footnote macro", func() {

		It("footnote with single-line content", func() {
			footnoteContent := "some content"
			actualContent := fmt.Sprintf(`foo footnote:[%s]`, footnoteContent)
			footnote1 := types.Footnote{
				ID: 0,
				Elements: types.InlineElements{
					types.StringElement{
						Content: footnoteContent,
					},
				},
			}
			expectedResult := types.Document{
				Attributes:        types.DocumentAttributes{},
				ElementReferences: map[string]interface{}{},
				Footnotes: []types.Footnote{
					footnote1,
				},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "foo ",
								},
								footnote1,
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent) // need to get the whole document here
		})

		It("footnote with single-line rich content", func() {
			actualContent := `foo footnote:[some *rich* http://foo.com[content]]`
			footnote1 := types.Footnote{
				ID: 0,
				Elements: types.InlineElements{
					types.StringElement{
						Content: "some ",
					},
					types.QuotedText{
						Kind: types.Bold,
						Elements: types.InlineElements{
							types.StringElement{
								Content: "rich",
							},
						},
					},
					types.StringElement{
						Content: " ",
					},
					types.InlineLink{
						Attributes: types.ElementAttributes{
							types.AttrInlineLinkText: "content",
						},
						URL: "http://foo.com",
					},
				},
			}
			expectedResult := types.Document{
				Attributes:        types.DocumentAttributes{},
				ElementReferences: map[string]interface{}{},
				Footnotes: []types.Footnote{
					footnote1,
				},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "foo ",
								},
								footnote1,
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent) // need to get the whole document here
		})

		It("footnote in a paragraph", func() {
			actualContent := `This is another paragraph.footnote:[I am footnote text and will be displayed at the bottom of the article.]`
			footnote1 := types.Footnote{
				ID: 0,
				Elements: types.InlineElements{
					types.StringElement{
						Content: "I am footnote text and will be displayed at the bottom of the article.",
					},
				},
			}
			expectedResult := types.Document{
				Attributes:        types.DocumentAttributes{},
				ElementReferences: map[string]interface{}{},
				Footnotes: []types.Footnote{
					footnote1,
				},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "This is another paragraph.",
								},
								footnote1,
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent) // need to get the whole document here
		})

		// It("footnote with multi-line content", func() {
		// 	footnoteContent := `some
		// 	content`
		// 	actualContent := fmt.Sprintf("foo footnote:[%s]", footnoteContent)
		// 	footnote1 := types.Footnote{
		// 		Elements: types.InlineElements{
		// 			types.StringElement{
		// 				Content: footnoteContent,
		// 			},
		// 		},
		// 	}
		// 	expectedResult := types.Document{
		// 		Attributes:        types.DocumentAttributes{},
		// 		ElementReferences: map[string]interface{}{},
		// 		Footnotes: types.Footnotes{
		// 			footnote1,
		// 		},
		// 		FootnoteReferences: types.FootnoteReferences{},
		// 		Elements: []interface{}{
		// 			types.Paragraph{
		// 				Attributes: types.ElementAttributes{},
		// 				Lines: []types.InlineElements{
		// 					{
		// 						types.StringElement{
		// 							Content: "foo ",
		// 						},
		// 						footnote1,
		// 					},
		// 				},
		// 			},
		// 		},
		// 	}
		// 	verify(GinkgoT(), expectedResult, actualContent) // need to get the whole document here
		// })
	})

	Context("footnoteref macro", func() {

		It("footnoteref with single-line content", func() {
			footnoteRef := "ref"
			footnoteContent := "some content"
			actualContent := fmt.Sprintf(`foo footnoteref:[%[1]s,%[2]s] and footnoteref:[%[1]s] again`, footnoteRef, footnoteContent)
			footnote1 := types.Footnote{
				ID:  0,
				Ref: footnoteRef,
				Elements: types.InlineElements{
					types.StringElement{
						Content: footnoteContent,
					},
				},
			}
			footnote2 := types.Footnote{
				ID:       1,
				Ref:      footnoteRef,
				Elements: types.InlineElements{},
			}
			expectedResult := types.Document{
				Attributes:        types.DocumentAttributes{},
				ElementReferences: map[string]interface{}{},
				Footnotes: types.Footnotes{
					footnote1,
				},
				FootnoteReferences: types.FootnoteReferences{
					"ref": footnote1,
				},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "foo ",
								},
								footnote1,
								types.StringElement{
									Content: " and ",
								},
								footnote2,
								types.StringElement{
									Content: " again",
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("footnoteref with unknown reference", func() {
			footnoteRef1 := "ref"
			footnoteRef2 := "ref2"
			footnoteContent := "some content"
			actualContent := fmt.Sprintf(`foo footnoteref:[%[1]s,%[2]s] and footnoteref:[%[3]s] again`, footnoteRef1, footnoteContent, footnoteRef2)
			footnote1 := types.Footnote{
				ID:  0,
				Ref: footnoteRef1,
				Elements: types.InlineElements{
					types.StringElement{
						Content: footnoteContent,
					},
				},
			}
			footnote2 := types.Footnote{
				ID:       1,
				Ref:      footnoteRef2,
				Elements: types.InlineElements{},
			}
			expectedResult := types.Document{
				Attributes:        types.DocumentAttributes{},
				ElementReferences: map[string]interface{}{},
				Footnotes: types.Footnotes{
					footnote1,
				},
				FootnoteReferences: types.FootnoteReferences{
					"ref": footnote1,
				},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "foo ",
								},
								footnote1,
								types.StringElement{
									Content: " and ",
								},
								footnote2,
								types.StringElement{
									Content: " again",
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

	It("footnotes in document", func() {

		actualContent := `= title

a premable with a footnote:[foo]

== section 1 footnote:[bar]

a paragraph with another footnote:[baz]`
		footnote1 := types.Footnote{
			ID: 0,
			Elements: types.InlineElements{
				types.StringElement{
					Content: "foo",
				},
			},
		}
		footnote2 := types.Footnote{
			ID: 1,
			Elements: types.InlineElements{
				types.StringElement{
					Content: "bar",
				},
			},
		}
		footnote3 := types.Footnote{
			ID: 2,
			Elements: types.InlineElements{
				types.StringElement{
					Content: "baz",
				},
			},
		}
		docTitle := types.SectionTitle{
			Attributes: types.ElementAttributes{
				types.AttrID: "_title",
			},
			Elements: types.InlineElements{
				types.StringElement{
					Content: "title",
				},
			},
		}
		section1Title := types.SectionTitle{
			Attributes: types.ElementAttributes{
				types.AttrID: "_section_1",
			},
			Elements: types.InlineElements{
				types.StringElement{
					Content: "section 1 ",
				},
				footnote2,
			},
		}
		expectedResult := types.Document{
			Attributes: types.DocumentAttributes{
				"doctitle": docTitle,
			},
			ElementReferences: map[string]interface{}{
				"_section_1": section1Title,
			},
			Footnotes: types.Footnotes{
				footnote1,
				footnote2,
				footnote3,
			},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Preamble{
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "a premable with a ",
									},
									footnote1,
								},
							},
						},
						types.BlankLine{},
					},
				},
				types.Section{
					Level: 1,
					Title: section1Title,
					Elements: []interface{}{
						types.BlankLine{},
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "a paragraph with another ",
									},
									footnote3,
								},
							},
						},
					},
				},
			},
		}
		verify(GinkgoT(), expectedResult, actualContent) // need to get the whole document here
	})
})
