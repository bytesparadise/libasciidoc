package parser_test

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("footnotes - preflight", func() {

	BeforeEach(func() {
		types.ResetFootnoteSequence()
	})

	Context("footnote macro", func() {

		It("footnote with single-line content", func() {
			footnoteContent := "some content"
			source := fmt.Sprintf(`foo footnote:[%s]`, footnoteContent)
			footnote1 := types.Footnote{
				ID: 0,
				Elements: types.InlineElements{
					&types.StringElement{
						Content: footnoteContent,
					},
				},
			}
			expected := &types.PreflightDocument{
				Blocks: []interface{}{
					&types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								&types.StringElement{
									Content: "foo ",
								},
								&footnote1,
							},
						},
					},
				},
			}
			verifyPreflight(expected, source) // need to get the whole document here
		})

		It("footnote with single-line rich content", func() {
			source := `foo footnote:[some *rich* https://foo.com[content]]`
			footnote1 := types.Footnote{
				ID: 0,
				Elements: types.InlineElements{
					&types.StringElement{
						Content: "some ",
					},
					&types.QuotedText{
						Kind: types.Bold,
						Elements: types.InlineElements{
							&types.StringElement{
								Content: "rich",
							},
						},
					},
					&types.StringElement{
						Content: " ",
					},
					&types.InlineLink{
						Attributes: types.ElementAttributes{
							types.AttrInlineLinkText: types.InlineElements{
								&types.StringElement{
									Content: "content",
								},
							},
						},
						Location: types.Location{
							&types.StringElement{
								Content: "https://foo.com",
							},
						},
					},
				},
			}
			expected := &types.PreflightDocument{
				Blocks: []interface{}{
					&types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								&types.StringElement{
									Content: "foo ",
								},
								&footnote1,
							},
						},
					},
				},
			}
			verifyPreflight(expected, source) // need to get the whole document here
		})

		It("footnote in a paragraph", func() {
			source := `This is another paragraph.footnote:[I am footnote text and will be displayed at the bottom of the article.]`
			expected := &types.PreflightDocument{
				Blocks: []interface{}{
					&types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								&types.StringElement{
									Content: "This is another paragraph.",
								},
								&types.Footnote{
									ID: 0,
									Elements: types.InlineElements{
										&types.StringElement{
											Content: "I am footnote text and will be displayed at the bottom of the article.",
										},
									},
								},
							},
						},
					},
				},
			}
			verifyPreflight(expected, source) // need to get the whole document here
		})

	})

	Context("footnoteref macro", func() {

		It("footnoteref with single-line content", func() {
			footnoteRef := "ref"
			footnoteContent := "some content"
			source := fmt.Sprintf(`foo footnoteref:[%[1]s,%[2]s] and footnoteref:[%[1]s] again`, footnoteRef, footnoteContent)
			footnote1 := types.Footnote{
				ID:  0,
				Ref: footnoteRef,
				Elements: types.InlineElements{
					&types.StringElement{
						Content: footnoteContent,
					},
				},
			}
			footnote2 := types.Footnote{
				ID:       1,
				Ref:      footnoteRef,
				Elements: types.InlineElements{},
			}
			expected := &types.PreflightDocument{
				Blocks: []interface{}{
					&types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								&types.StringElement{
									Content: "foo ",
								},
								&footnote1,
								&types.StringElement{
									Content: " and ",
								},
								&footnote2,
								&types.StringElement{
									Content: " again",
								},
							},
						},
					},
				},
			}
			verifyPreflight(expected, source)
		})

		It("footnoteref with unknown reference", func() {
			footnoteRef1 := "ref"
			footnoteRef2 := "ref2"
			footnoteContent := "some content"
			source := fmt.Sprintf(`foo footnoteref:[%[1]s,%[2]s] and footnoteref:[%[3]s] again`, footnoteRef1, footnoteContent, footnoteRef2)
			footnote1 := types.Footnote{
				ID:  0,
				Ref: footnoteRef1,
				Elements: types.InlineElements{
					&types.StringElement{
						Content: footnoteContent,
					},
				},
			}
			footnote2 := types.Footnote{
				ID:       1,
				Ref:      footnoteRef2,
				Elements: types.InlineElements{},
			}
			expected := &types.PreflightDocument{
				Blocks: []interface{}{
					&types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								&types.StringElement{
									Content: "foo ",
								},
								&footnote1,
								&types.StringElement{
									Content: " and ",
								},
								&footnote2,
								&types.StringElement{
									Content: " again",
								},
							},
						},
					},
				},
			}
			verifyPreflight(expected, source)
		})
	})

	It("footnotes in document", func() {

		source := `= title
:idprefix: id_

a premable with a footnote:[foo]

== section 1 footnote:[bar]

a paragraph with another footnote:[baz]`
		footnote1 := types.Footnote{
			ID: 0,
			Elements: types.InlineElements{
				&types.StringElement{
					Content: "foo",
				},
			},
		}
		footnote2 := types.Footnote{
			ID: 1,
			Elements: types.InlineElements{
				&types.StringElement{
					Content: "bar",
				},
			},
		}
		footnote3 := types.Footnote{
			ID: 2,
			Elements: types.InlineElements{
				&types.StringElement{
					Content: "baz",
				},
			},
		}
		docTitle := types.InlineElements{
			&types.StringElement{
				Content: "title",
			},
		}
		section1Title := types.InlineElements{
			&types.StringElement{
				Content: "section 1 ",
			},
			&footnote2,
		}
		expected := &types.PreflightDocument{
			Blocks: []interface{}{
				&types.Section{
					Level: 0,
					Title: docTitle,
					Attributes: types.ElementAttributes{
						types.AttrID:       "title",
						types.AttrCustomID: false,
					},

					Elements: []interface{}{},
				},
				&types.DocumentAttributeDeclaration{
					Name:  "idprefix",
					Value: "id_",
				},
				&types.BlankLine{},
				&types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							&types.StringElement{
								Content: "a premable with a ",
							},
							&footnote1,
						},
					},
				},
				&types.BlankLine{},
				&types.Section{
					Attributes: types.ElementAttributes{
						types.AttrID:       "section_1",
						types.AttrCustomID: false,
					},
					Level:    1,
					Title:    section1Title,
					Elements: []interface{}{},
				},
				&types.BlankLine{},
				&types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							&types.StringElement{
								Content: "a paragraph with another ",
							},
							&footnote3,
						},
					},
				},
			},
		}
		verifyPreflight(expected, source) // need to get the whole document here
	})
})

var _ = Describe("footnotes - document", func() {

	BeforeEach(func() {
		types.ResetFootnoteSequence()
	})

	Context("footnote macro", func() {

		It("footnote with single-line content", func() {
			footnoteContent := "some content"
			source := fmt.Sprintf(`foo footnote:[%s]`, footnoteContent)
			footnote1 := types.Footnote{
				ID: 0,
				Elements: types.InlineElements{
					&types.StringElement{
						Content: footnoteContent,
					},
				},
			}
			expected := &types.Document{
				Attributes:        types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{},
				Footnotes: []*types.Footnote{
					&footnote1,
				},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					&types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								&types.StringElement{
									Content: "foo ",
								},
								&footnote1,
							},
						},
					},
				},
			}
			verifyDocument(expected, source) // need to get the whole document here
		})

		It("footnote with single-line rich content", func() {
			source := `foo footnote:[some *rich* https://foo.com[content]]`
			footnote1 := types.Footnote{
				ID: 0,
				Elements: types.InlineElements{
					&types.StringElement{
						Content: "some ",
					},
					&types.QuotedText{
						Kind: types.Bold,
						Elements: types.InlineElements{
							&types.StringElement{
								Content: "rich",
							},
						},
					},
					&types.StringElement{
						Content: " ",
					},
					&types.InlineLink{
						Attributes: types.ElementAttributes{
							types.AttrInlineLinkText: types.InlineElements{
								&types.StringElement{
									Content: "content",
								},
							},
						},
						Location: types.Location{
							&types.StringElement{
								Content: "https://foo.com",
							},
						},
					},
				},
			}
			expected := &types.Document{
				Attributes:        types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{},
				Footnotes: []*types.Footnote{
					&footnote1,
				},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					&types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								&types.StringElement{
									Content: "foo ",
								},
								&footnote1,
							},
						},
					},
				},
			}
			verifyDocument(expected, source) // need to get the whole document here
		})

		It("footnote in a paragraph", func() {
			source := `This is another paragraph.footnote:[I am footnote text and will be displayed at the bottom of the article.]`
			footnote1 := types.Footnote{
				ID: 0,
				Elements: types.InlineElements{
					&types.StringElement{
						Content: "I am footnote text and will be displayed at the bottom of the article.",
					},
				},
			}
			expected := &types.Document{
				Attributes:        types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{},
				Footnotes: []*types.Footnote{
					&footnote1,
				},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					&types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								&types.StringElement{
									Content: "This is another paragraph.",
								},
								&footnote1,
							},
						},
					},
				},
			}
			verifyDocument(expected, source) // need to get the whole document here
		})

	})

	Context("footnoteref macro", func() {

		It("footnoteref with single-line content", func() {
			footnoteRef := "ref"
			footnoteContent := "some content"
			source := fmt.Sprintf(`foo footnoteref:[%[1]s,%[2]s] and footnoteref:[%[1]s] again`, footnoteRef, footnoteContent)
			footnote1 := types.Footnote{
				ID:  0,
				Ref: footnoteRef,
				Elements: types.InlineElements{
					&types.StringElement{
						Content: footnoteContent,
					},
				},
			}
			footnote2 := types.Footnote{
				ID:       1,
				Ref:      footnoteRef,
				Elements: types.InlineElements{},
			}
			expected := &types.Document{
				Attributes:        types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{},
				Footnotes: types.Footnotes{
					&footnote1,
				},
				FootnoteReferences: types.FootnoteReferences{
					"ref": &footnote1,
				},
				Elements: []interface{}{
					&types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								&types.StringElement{
									Content: "foo ",
								},
								&footnote1,
								&types.StringElement{
									Content: " and ",
								},
								&footnote2,
								&types.StringElement{
									Content: " again",
								},
							},
						},
					},
				},
			}
			verifyDocument(expected, source)
		})

		It("footnoteref with unknown reference", func() {
			footnoteRef1 := "ref"
			footnoteRef2 := "ref2"
			footnoteContent := "some content"
			source := fmt.Sprintf(`foo footnoteref:[%[1]s,%[2]s] and footnoteref:[%[3]s] again`, footnoteRef1, footnoteContent, footnoteRef2)
			footnote1 := types.Footnote{
				ID:  0,
				Ref: footnoteRef1,
				Elements: types.InlineElements{
					&types.StringElement{
						Content: footnoteContent,
					},
				},
			}
			footnote2 := types.Footnote{
				ID:       1,
				Ref:      footnoteRef2,
				Elements: types.InlineElements{},
			}
			expected := &types.Document{
				Attributes:        types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{},
				Footnotes: types.Footnotes{
					&footnote1,
				},
				FootnoteReferences: types.FootnoteReferences{
					"ref": &footnote1,
				},
				Elements: []interface{}{
					&types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								&types.StringElement{
									Content: "foo ",
								},
								&footnote1,
								&types.StringElement{
									Content: " and ",
								},
								&footnote2,
								&types.StringElement{
									Content: " again",
								},
							},
						},
					},
				},
			}
			verifyDocument(expected, source)
		})
	})

	It("footnotes in document", func() {

		source := `= title
:idprefix: id_

a premable with a footnote:[foo]

== section 1 footnote:[bar]

a paragraph with another footnote:[baz]`
		footnote1 := types.Footnote{
			ID: 0,
			Elements: types.InlineElements{
				&types.StringElement{
					Content: "foo",
				},
			},
		}
		footnote2 := types.Footnote{
			ID: 1,
			Elements: types.InlineElements{
				&types.StringElement{
					Content: "bar",
				},
			},
		}
		footnote3 := types.Footnote{
			ID: 2,
			Elements: types.InlineElements{
				&types.StringElement{
					Content: "baz",
				},
			},
		}
		docTitle := types.InlineElements{
			&types.StringElement{
				Content: "title",
			},
		}
		section1Title := types.InlineElements{
			&types.StringElement{
				Content: "section 1 ",
			},
			&footnote2,
		}
		expected := &types.Document{
			Attributes: types.DocumentAttributes{},
			ElementReferences: types.ElementReferences{
				"title":     docTitle,
				"section_1": section1Title,
			},
			Footnotes: types.Footnotes{
				&footnote1,
				&footnote2,
				&footnote3,
			},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				&types.Section{
					Level: 0,
					Title: docTitle,
					Attributes: types.ElementAttributes{
						types.AttrID:       "title",
						types.AttrCustomID: false,
					},
					Elements: []interface{}{
						&types.DocumentAttributeDeclaration{
							Name:  "idprefix",
							Value: "id_",
						},
						&types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									&types.StringElement{
										Content: "a premable with a ",
									},
									&footnote1,
								},
							},
						},
						&types.Section{
							Attributes: types.ElementAttributes{
								types.AttrID:       "section_1",
								types.AttrCustomID: false,
							},
							Level: 1,
							Title: section1Title,
							Elements: []interface{}{
								&types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: []types.InlineElements{
										{
											&types.StringElement{
												Content: "a paragraph with another ",
											},
											&footnote3,
										},
									},
								},
							},
						},
					},
				},
			},
		}
		verifyDocument(expected, source) // need to get the whole document here
	})
})
