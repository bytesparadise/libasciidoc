package renderer_test

import (
	"context"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/davecgh/go-spew/spew"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("file inclusions", func() {

	It("should include adoc file with section 0 at root level without offset", func() {
		actualContent := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "a first paragraph"},
						},
					},
				},
				types.FileInclusion{
					Attributes: types.ElementAttributes{},
					Path:       "html5/includes/chapter-a.adoc",
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "a second paragraph"},
						},
					},
				},
			},
		}
		expectedResult := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "a first paragraph"},
						},
					},
				},
				types.Section{
					Level: 0,
					Title: types.SectionTitle{
						Attributes: types.ElementAttributes{
							types.AttrID:       "chapter_a",
							types.AttrCustomID: false,
						},
						Elements: []interface{}{
							types.StringElement{
								Content: "Chapter A",
							},
						},
					},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{Content: "content"},
								},
							},
						},
					},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "a second paragraph"},
						},
					},
				},
			},
		}
		verifyFileInclusions(expectedResult, actualContent)
	})

	It("should include adoc file with section 0 at root level with valid offset", func() {
		actualContent := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "a first paragraph"},
						},
					},
				},
				types.FileInclusion{
					Attributes: types.ElementAttributes{
						types.AttrLevelOffset: "+1",
					},
					Path: "html5/includes/chapter-a.adoc",
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "a second paragraph"},
						},
					},
				},
			},
		}
		expectedResult := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "a first paragraph"},
						},
					},
				},
				types.Section{
					Level: 1,
					Title: types.SectionTitle{
						Attributes: types.ElementAttributes{
							types.AttrID:       "chapter_a",
							types.AttrCustomID: false,
						},
						Elements: []interface{}{
							types.StringElement{
								Content: "Chapter A",
							},
						},
					},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{Content: "content"},
								},
							},
						},
					},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "a second paragraph"},
						},
					},
				},
			},
		}
		verifyFileInclusions(expectedResult, actualContent)
	})

	It("should include adoc file with section 0 within existin section with valid offset", func() {
		actualContent := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Section{
					Level: 1,
					Title: types.SectionTitle{
						Attributes: types.ElementAttributes{
							types.AttrID:       "chapter_a",
							types.AttrCustomID: false,
						},
						Elements: []interface{}{
							types.StringElement{
								Content: "Chapter A",
							},
						},
					},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{Content: "a first paragraph"},
								},
							},
						},
						types.FileInclusion{
							Attributes: types.ElementAttributes{
								types.AttrLevelOffset: "+2",
							},
							Path: "html5/includes/chapter-a.adoc",
						},
						types.BlankLine{},
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{Content: "a second paragraph"},
								},
							},
						},
					},
				},
			},
		}
		expectedResult := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Section{
					Level: 1,
					Title: types.SectionTitle{
						Attributes: types.ElementAttributes{
							types.AttrID:       "chapter_a",
							types.AttrCustomID: false,
						},
						Elements: []interface{}{
							types.StringElement{
								Content: "Chapter A",
							},
						},
					},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{Content: "a first paragraph"},
								},
							},
						},
						types.Section{
							Level: 2,
							Title: types.SectionTitle{
								Attributes: types.ElementAttributes{
									types.AttrID:       "chapter_a",
									types.AttrCustomID: false,
								},
								Elements: []interface{}{
									types.StringElement{
										Content: "Chapter A",
									},
								},
							},
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: []types.InlineElements{
										{
											types.StringElement{Content: "content"},
										},
									},
								},
							},
						},
						types.BlankLine{},
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{Content: "a second paragraph"},
								},
							},
						},
					},
				},
			},
		}
		verifyFileInclusions(expectedResult, actualContent)
	})

	It("should include adoc file with 2 paragraphs at root level without offset", func() {
		actualContent := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "a first paragraph"},
						},
					},
				},
				types.FileInclusion{
					Path: "html5/includes/grandchild-include.adoc",
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "a second paragraph"},
						},
					},
				},
			},
		}
		expectedResult := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "a first paragraph"},
						},
					},
				},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "first line of grandchild"},
						},
					},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "last line of grandchild"},
						},
					},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "a second paragraph"},
						},
					},
				},
			},
		}
		verifyFileInclusions(expectedResult, actualContent)
	})

	It("should include unparsed adoc file in delimited block", func() {
		actualContent := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "a first paragraph"},
						},
					},
				},
				types.DelimitedBlock{
					Kind:       types.Source,
					Attributes: types.ElementAttributes{},
					Elements: []interface{}{
						types.FileInclusion{
							Attributes: types.ElementAttributes{},
							Path:       "html5/includes/chapter-a.adoc",
						},
					},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "a second paragraph"},
						},
					},
				},
			},
		}
		expectedResult := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "a first paragraph"},
						},
					},
				},
				types.DelimitedBlock{
					Kind:       types.Source,
					Attributes: types.ElementAttributes{},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{Content: "= Chapter A"},
								},
							},
						},
						types.BlankLine{},
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{Content: "content"},
								},
							},
						},
					},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "a second paragraph"},
						},
					},
				},
			},
		}
		verifyFileInclusions(expectedResult, actualContent)
	})

	It("should include unparsed adoc file with line range in delimited block", func() {
		actualContent := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "a first paragraph"},
						},
					},
				},
				types.DelimitedBlock{
					Kind:       types.Source,
					Attributes: types.ElementAttributes{},
					Elements: []interface{}{
						types.FileInclusion{
							Attributes: types.ElementAttributes{
								types.AttrLineRanges: types.LineRanges{
									{Start: 3, End: 3},
								},
							},
							Path: "html5/includes/chapter-a.adoc",
						},
					},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "a second paragraph"},
						},
					},
				},
			},
		}
		expectedResult := types.Document{
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "a first paragraph"},
						},
					},
				},
				types.DelimitedBlock{
					Kind:       types.Source,
					Attributes: types.ElementAttributes{},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{Content: "content"},
								},
							},
						},
					},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{Content: "a second paragraph"},
						},
					},
				},
			},
		}
		verifyFileInclusions(expectedResult, actualContent)
	})

	Context("lines to include with ranges", func() {

		It("include with a single line range", func() {
			actualContent := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.FileInclusion{
						Attributes: types.ElementAttributes{
							types.AttrLineRanges: types.LineRanges{
								types.LineRange{Start: 1, End: 2}, // include blankline to get 2 paragraphs
								types.LineRange{Start: 5, End: -1},
							},
						},
						Path: "html5/includes/hello_world.go",
					},
				},
			}
			expectedResult := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "package includes",
								},
							},
						},
					},
					types.BlankLine{},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{
									Content: "func helloworld() {",
								},
							},
							{
								types.StringElement{
									Content: `	fmt.Println("hello, world!")`,
								},
							},
							{
								types.StringElement{
									Content: "}",
								},
							},
						},
					},
				},
			}
			verifyFileInclusions(expectedResult, actualContent)
		})

	})
})

func verifyFileInclusions(expectedResult, actualContent types.Document) {
	ctx := renderer.Wrap(context.Background(), actualContent)
	err := renderer.ProcessFileInclusions(ctx)
	Expect(err).ShouldNot(HaveOccurred())
	GinkgoT().Logf("actual document: `%s`", spew.Sdump(ctx.Document))
	GinkgoT().Logf("expected document: `%s`", spew.Sdump(expectedResult))
	assert.EqualValues(GinkgoT(), expectedResult, ctx.Document)
}

var _ = Describe("sections level offset", func() {

	It("should apply level offset without section 0", func() {
		section1Title := types.SectionTitle{
			Attributes: types.ElementAttributes{
				types.AttrID:       "section_1",
				types.AttrCustomID: false,
			},
			Elements: types.InlineElements{
				types.StringElement{
					Content: "section 1 title",
				},
			},
		}
		section2Title := types.SectionTitle{
			Attributes: types.ElementAttributes{
				types.AttrID:       "section_2",
				types.AttrCustomID: false,
			},
			Elements: types.InlineElements{
				types.StringElement{
					Content: "section 2 title",
				},
			},
		}

		actualContent := types.Document{
			Attributes: types.DocumentAttributes{},
			ElementReferences: types.ElementReferences{
				"section_1": section1Title,
				"section_2": section2Title,
			},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Section{
					Level: 1,
					Title: section1Title,
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "a paragraph...",
									},
								},
							},
						},
						types.Section{
							Level:    2,
							Title:    section2Title,
							Elements: []interface{}{},
						},
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "another paragraph...",
									},
								},
							},
						},
					},
				},
			},
		}

		expectedResult := types.Document{
			Attributes: types.DocumentAttributes{},
			ElementReferences: types.ElementReferences{
				"section_1": section1Title,
				"section_2": section2Title,
			},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Section{
					Level: 2,
					Title: section1Title,
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "a paragraph...",
									},
								},
							},
						},
						types.Section{
							Level:    3,
							Title:    section2Title,
							Elements: []interface{}{},
						},
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "another paragraph...",
									},
								},
							},
						},
					},
				},
			},
		}

		verifyLevelOffset(expectedResult, actualContent, "+1")
	})

	It("should apply level offset with section 0", func() {
		docTitle := types.SectionTitle{
			Attributes: types.ElementAttributes{
				types.AttrID:       "title",
				types.AttrCustomID: false,
			},
			Elements: types.InlineElements{
				types.StringElement{
					Content: "title",
				},
			},
		}
		section1Title := types.SectionTitle{
			Attributes: types.ElementAttributes{
				types.AttrID:       "section_1",
				types.AttrCustomID: false,
			},
			Elements: types.InlineElements{
				types.StringElement{
					Content: "section 1 title",
				},
			},
		}

		actualContent := types.Document{
			Attributes: types.DocumentAttributes{
				types.AttrTitle: docTitle,
				"idprefix":      "id_",
			},
			ElementReferences: types.ElementReferences{
				"section_1": section1Title,
			},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Section{
					Level:    1,
					Title:    section1Title,
					Elements: []interface{}{},
				},
			},
		}

		expectedResult := types.Document{
			Attributes: types.DocumentAttributes{
				types.AttrTitle: docTitle,
				"idprefix":      "id_",
			},
			ElementReferences: types.ElementReferences{
				"section_1": section1Title,
			},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Section{
					Level: 1,
					Title: docTitle,
					Elements: []interface{}{
						types.Section{
							Level:    2,
							Title:    section1Title,
							Elements: []interface{}{},
						},
					},
				},
			},
		}

		verifyLevelOffset(expectedResult, actualContent, "+1")
	})

	It("should not change elements when empty level offset", func() {
		section1Title := types.SectionTitle{
			Attributes: types.ElementAttributes{
				types.AttrID:       "section_1",
				types.AttrCustomID: false,
			},
			Elements: types.InlineElements{
				types.StringElement{
					Content: "section 1 title",
				},
			},
		}
		section2Title := types.SectionTitle{
			Attributes: types.ElementAttributes{
				types.AttrID:       "section_2",
				types.AttrCustomID: false,
			},
			Elements: types.InlineElements{
				types.StringElement{
					Content: "section 2 title",
				},
			},
		}

		actualContent := types.Document{
			Attributes: types.DocumentAttributes{},
			ElementReferences: types.ElementReferences{
				"section_1": section1Title,
				"section_2": section2Title,
			},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Section{
					Level: 1,
					Title: section1Title,
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "a paragraph...",
									},
								},
							},
						},
						types.Section{
							Level:    2,
							Title:    section2Title,
							Elements: []interface{}{},
						},
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "another paragraph...",
									},
								},
							},
						},
					},
				},
			},
		}
		expectedResult := types.Document{
			Attributes: types.DocumentAttributes{},
			ElementReferences: types.ElementReferences{
				"section_1": section1Title,
				"section_2": section2Title,
			},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Section{
					Level: 1,
					Title: section1Title,
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "a paragraph...",
									},
								},
							},
						},
						types.Section{
							Level:    2,
							Title:    section2Title,
							Elements: []interface{}{},
						},
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									types.StringElement{
										Content: "another paragraph...",
									},
								},
							},
						},
					},
				},
			},
		}
		verifyLevelOffset(expectedResult, actualContent, "")
	})

})

func verifyLevelOffset(expectedResult, actualContent types.Document, levelOffset string) {
	result, err := renderer.ApplyLevelOffset(actualContent, levelOffset)
	Expect(err).ShouldNot(HaveOccurred())
	GinkgoT().Logf("actual document: `%s`", spew.Sdump(result))
	GinkgoT().Logf("expected document: `%s`", spew.Sdump(expectedResult))
	assert.EqualValues(GinkgoT(), expectedResult, result)
}
