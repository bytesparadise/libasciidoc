package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("sections", func() {

	Context("valid sections", func() {

		It("header only", func() {
			actualContent := "= a header"
			doctitle := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID:       "a_header",
					types.AttrCustomID: false,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "a header"},
				},
			}
			expectedResult := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"a_header": doctitle,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Level:      0,
						Title:      doctitle,
						Attributes: types.ElementAttributes{},
						Elements:   []interface{}{},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent)
		})

		It("header with many spaces around content", func() {
			actualContent := "= a header   "
			doctitle := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID:       "a_header",
					types.AttrCustomID: false,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "a header   "},
				},
			}
			expectedResult := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"a_header": doctitle,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Level:      0,
						Title:      doctitle,
						Attributes: types.ElementAttributes{},
						Elements:   []interface{}{},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent)
		})

		It("header and paragraph", func() {
			actualContent := `= a header

and a paragraph`

			doctitle := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID:       "a_header",
					types.AttrCustomID: false,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "a header"},
				},
			}
			expectedResult := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"a_header": doctitle,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Level:      0,
						Title:      doctitle,
						Attributes: types.ElementAttributes{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "and a paragraph"},
									},
								},
							},
						},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent)
		})

		It("two sections with level 0", func() {
			actualContent := `= a first header

= a second header`
			doctitle := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID:       "a_first_header",
					types.AttrCustomID: false,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "a first header"},
				},
			}
			otherDoctitle := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID:       "a_second_header",
					types.AttrCustomID: false,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "a second header"},
				},
			}

			expectedResult := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"a_first_header":  doctitle,
					"a_second_header": otherDoctitle,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Level:      0,
						Title:      doctitle,
						Attributes: types.ElementAttributes{},
						Elements:   []interface{}{},
					},
					types.Section{
						Level:      0,
						Title:      otherDoctitle,
						Attributes: types.ElementAttributes{},
						Elements:   []interface{}{},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent)
		})

		It("section level 1 alone", func() {
			actualContent := `== section 1`
			section1Title := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID:       "section_1",
					types.AttrCustomID: false,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "section 1"},
				},
			}
			expectedResult := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"section_1": section1Title,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Level:      1,
						Title:      section1Title,
						Attributes: types.ElementAttributes{},
						Elements:   []interface{}{},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent)
		})

		It("section level 1 with quoted text", func() {
			actualContent := `==  *2 spaces and bold content*`
			sectionTitle := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID:       "2_spaces_and_bold_content",
					types.AttrCustomID: false,
				},
				Elements: types.InlineElements{
					types.QuotedText{
						Kind: types.Bold,
						Elements: types.InlineElements{
							types.StringElement{Content: "2 spaces and bold content"},
						},
					},
				},
			}
			expectedResult := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"2_spaces_and_bold_content": sectionTitle,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Level:      1,
						Title:      sectionTitle,
						Attributes: types.ElementAttributes{},
						Elements:   []interface{}{},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent)
		})

		It("section level 0 with nested section level 1", func() {
			actualContent := `= a header

== section 1`
			doctitle := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID:       "a_header",
					types.AttrCustomID: false,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "a header"},
				},
			}
			section1Title := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID:       "section_1",
					types.AttrCustomID: false,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "section 1"},
				},
			}
			expectedResult := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"a_header":  doctitle,
					"section_1": section1Title,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Level:      0,
						Title:      doctitle,
						Attributes: types.ElementAttributes{},
						Elements: []interface{}{
							types.Section{
								Level:      1,
								Title:      section1Title,
								Attributes: types.ElementAttributes{},
								Elements:   []interface{}{},
							},
						},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent)
		})

		It("section level 0 with nested section level 2", func() {
			actualContent := `= a header

=== section 2`
			doctitle := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID:       "a_header",
					types.AttrCustomID: false,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "a header"},
				},
			}
			section2Title := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID:       "section_2",
					types.AttrCustomID: false,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "section 2"},
				},
			}
			expectedResult := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"a_header":  doctitle,
					"section_2": section2Title,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Level:      0,
						Title:      doctitle,
						Attributes: types.ElementAttributes{},
						Elements: []interface{}{
							types.Section{
								Level:      2,
								Title:      section2Title,
								Attributes: types.ElementAttributes{},
								Elements:   []interface{}{},
							},
						},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent)
		})

		It("section level 1 with immediate paragraph", func() {
			actualContent := `== a title
and a paragraph`
			section1Title := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID:       "a_title",
					types.AttrCustomID: false,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "a title"},
				},
			}
			expectedResult := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"a_title": section1Title,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Level:      1,
						Title:      section1Title,
						Attributes: types.ElementAttributes{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "and a paragraph"},
									},
								},
							},
						},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent)
		})

		It("section level 1 with a paragraph separated by empty line", func() {
			actualContent := `== a title
			
and a paragraph`
			section1Title := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID:       "a_title",
					types.AttrCustomID: false,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "a title"},
				},
			}
			expectedResult := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"a_title": section1Title,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Level:      1,
						Title:      section1Title,
						Attributes: types.ElementAttributes{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "and a paragraph"},
									},
								},
							},
						},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent)
		})

		It("section level 1 with a paragraph separated by non-empty line", func() {
			actualContent := "== a title\n    \nand a paragraph"
			section1Title := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID:       "a_title",
					types.AttrCustomID: false,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "a title"},
				},
			}
			expectedResult := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"a_title": section1Title,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Level:      1,
						Title:      section1Title,
						Attributes: types.ElementAttributes{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "and a paragraph"},
									},
								},
							},
						},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent)
		})

		It("section levels 1, 2, 3, 2", func() {
			actualContent := `= a header

== Section A
a paragraph

=== Section A.a
a paragraph

== Section B
a paragraph`
			doctitle := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID:       "a_header",
					types.AttrCustomID: false,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "a header"},
				},
			}
			sectionATitle := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID:       "section_a",
					types.AttrCustomID: false,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "Section A"},
				},
			}
			sectionAaTitle := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID:       "section_a_a",
					types.AttrCustomID: false,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "Section A.a"},
				},
			}
			sectionBTitle := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID:       "section_b",
					types.AttrCustomID: false,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "Section B"},
				},
			}
			expectedResult := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"a_header":    doctitle,
					"section_a":   sectionATitle,
					"section_a_a": sectionAaTitle,
					"section_b":   sectionBTitle,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Level:      0,
						Title:      doctitle,
						Attributes: types.ElementAttributes{},
						Elements: []interface{}{
							types.Section{
								Level:      1,
								Title:      sectionATitle,
								Attributes: types.ElementAttributes{},
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: []types.InlineElements{
											{
												types.StringElement{Content: "a paragraph"},
											},
										},
									},
									types.BlankLine{},
									types.Section{
										Level:      2,
										Title:      sectionAaTitle,
										Attributes: types.ElementAttributes{},
										Elements: []interface{}{
											types.Paragraph{
												Attributes: types.ElementAttributes{},
												Lines: []types.InlineElements{
													{
														types.StringElement{Content: "a paragraph"},
													},
												},
											},
											types.BlankLine{},
										},
									},
								},
							},
							types.Section{
								Level:      1,
								Title:      sectionBTitle,
								Attributes: types.ElementAttributes{},
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: []types.InlineElements{
											{
												types.StringElement{Content: "a paragraph"},
											},
										},
									},
								},
							},
						},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent)
		})

		It("single section with custom IDs", func() {
			actualContent := `[[custom_header]]
== a header`
			sectionTitle := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID:       "custom_header",
					types.AttrCustomID: true,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "a header"},
				},
			}
			expectedResult := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"custom_header": sectionTitle,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Level:      1,
						Title:      sectionTitle,
						Attributes: types.ElementAttributes{},
						Elements:   []interface{}{},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent)
		})

		It("multiple sections with custom IDs", func() {
			actualContent := `[[custom_header]]
= a header

== Section F [[ignored]] [[foo]]

[[bar]]
== Section B
a paragraph`
			doctitle := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID:       "custom_header",
					types.AttrCustomID: true,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "a header"},
				},
			}
			fooTitle := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID:       "foo",
					types.AttrCustomID: true,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "Section F "},
				},
			}
			barTitle := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID:       "bar",
					types.AttrCustomID: true,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "Section B"},
				},
			}
			expectedResult := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"custom_header": doctitle,
					"foo":           fooTitle,
					"bar":           barTitle,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Level:      0,
						Title:      doctitle,
						Attributes: types.ElementAttributes{},
						Elements: []interface{}{
							types.Section{
								Level:      1,
								Title:      fooTitle,
								Attributes: types.ElementAttributes{},
								Elements:   []interface{}{},
							},
							types.Section{
								Level:      1,
								Title:      barTitle,
								Attributes: types.ElementAttributes{},
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: []types.InlineElements{
											{
												types.StringElement{Content: "a paragraph"},
											},
										},
									},
								},
							},
						},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent)
		})

		It("sections with same title", func() {
			actualContent := `== section 1

== section 1`
			section1aTitle := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID:       "section_1",
					types.AttrCustomID: false,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "section 1"},
				},
			}
			section1bTitle := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID:       "section_1_2",
					types.AttrCustomID: false,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "section 1"},
				},
			}

			expectedResult := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"section_1":   section1aTitle,
					"section_1_2": section1bTitle,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Level:      1,
						Title:      section1aTitle,
						Attributes: types.ElementAttributes{},
						Elements:   []interface{}{},
					},
					types.Section{
						Level:      1,
						Title:      section1bTitle,
						Attributes: types.ElementAttributes{},
						Elements:   []interface{}{},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent)
		})
	})

	Context("invalid sections", func() {
		It("header invalid - missing space", func() {
			actualContent := "=a header"
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
								types.StringElement{Content: "=a header"},
							},
						},
					},
				}}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent)
		})

		It("header invalid - header space", func() {
			actualContent := " = a header with a prefix space"
			expectedResult := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.LiteralBlock{
						Attributes: types.ElementAttributes{
							types.AttrKind:             types.Literal,
							types.AttrLiteralBlockType: types.LiteralBlockWithSpacesOnFirstLine,
						},
						Lines: []string{
							" = a header with a prefix space",
						},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent)
		})

		It("header with invalid section1", func() {
			actualContent := `= a header

 == section with prefix space`
			title := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID:       "a_header",
					types.AttrCustomID: false,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "a header"},
				},
			}
			expectedResult := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"a_header": title,
				}, Footnotes: types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Level:      0,
						Title:      title,
						Attributes: types.ElementAttributes{},
						Elements: []interface{}{
							types.LiteralBlock{
								Attributes: types.ElementAttributes{
									types.AttrKind:             types.Literal,
									types.AttrLiteralBlockType: types.LiteralBlockWithSpacesOnFirstLine,
								},
								Lines: []string{
									" == section with prefix space",
								},
							},
						},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent)
		})

	})

	Context("unsupported section syntax", func() {

		It("should not fail with underlined title", func() {
			actualContent := `Document Title
==============
Doc Writer <thedoc@asciidoctor.org>`
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
									Content: "Document Title",
								},
							},
							{
								types.StringElement{
									Content: "==============",
								},
							},
							{
								types.StringElement{
									Content: "Doc Writer <thedoc@asciidoctor.org>",
								},
							},
						},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent)
		})
	})
})
