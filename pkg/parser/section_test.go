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
					types.AttrID: "_a_header",
				},
				Content: types.InlineElements{
					types.StringElement{Content: "a header"},
				},
			}
			expectedResult := types.Document{
				Attributes: types.DocumentAttributes{
					"doctitle": doctitle,
				},
				ElementReferences: map[string]interface{}{},
				Elements:          []interface{}{},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("header with many spaces around content", func() {
			actualContent := "= a header   "
			doctitle := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID: "_a_header",
				},
				Content: types.InlineElements{
					types.StringElement{Content: "a header   "},
				},
			}
			expectedResult := types.Document{
				Attributes: types.DocumentAttributes{
					"doctitle": doctitle,
				},
				ElementReferences: map[string]interface{}{},
				Elements:          []interface{}{},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("header and paragraph", func() {
			actualContent := `= a header

and a paragraph`

			doctitle := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID: "_a_header",
				},
				Content: types.InlineElements{
					types.StringElement{Content: "a header"},
				},
			}
			expectedResult := types.Document{
				Attributes: types.DocumentAttributes{
					"doctitle": doctitle,
				},
				ElementReferences: map[string]interface{}{},
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
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("section level 1 alone", func() {
			actualContent := `== section 1`
			section1Title := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID: "_section_1",
				},
				Content: types.InlineElements{
					types.StringElement{Content: "section 1"},
				},
			}
			expectedResult := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: map[string]interface{}{
					"_section_1": section1Title,
				},
				Elements: []interface{}{
					types.Section{
						Level:    1,
						Title:    section1Title,
						Elements: []interface{}{},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("section level 1 with quoted text", func() {
			actualContent := `==  *2 spaces and bold content*`
			sectionTitle := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID: "_2_spaces_and_bold_content",
				},
				Content: types.InlineElements{
					types.QuotedText{
						Elements: types.InlineElements{
							types.StringElement{Content: "2 spaces and bold content"},
						},
					},
				},
			}
			expectedResult := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: map[string]interface{}{
					"_2_spaces_and_bold_content": sectionTitle,
				},
				Elements: []interface{}{
					types.Section{
						Level:    1,
						Title:    sectionTitle,
						Elements: []interface{}{},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("section level 0 with nested section level 1", func() {
			actualContent := `= a header

== section 1`
			doctitle := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID: "_a_header",
				},
				Content: types.InlineElements{
					types.StringElement{Content: "a header"},
				},
			}
			section1Title := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID: "_section_1",
				},
				Content: types.InlineElements{
					types.StringElement{Content: "section 1"},
				},
			}
			expectedResult := types.Document{
				Attributes: types.DocumentAttributes{
					"doctitle": doctitle,
				},
				ElementReferences: map[string]interface{}{
					"_section_1": section1Title,
				},
				Elements: []interface{}{
					types.Section{
						Level:    1,
						Title:    section1Title,
						Elements: []interface{}{},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("section level 0 with preamble and section level 1", func() {
			actualContent := `= a header

a short preamble

== section 1`
			section1Title := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID: "_section_1",
				},
				Content: types.InlineElements{
					types.StringElement{Content: "section 1"},
				},
			}
			doctitle := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID: "_a_header",
				},
				Content: types.InlineElements{
					types.StringElement{Content: "a header"},
				},
			}
			expectedResult := types.Document{
				Attributes: types.DocumentAttributes{
					"doctitle": doctitle,
				},
				ElementReferences: map[string]interface{}{
					"_section_1": section1Title,
				},
				Elements: []interface{}{
					types.Preamble{
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "a short preamble"},
									},
								},
							},
							types.BlankLine{},
						},
					},
					types.Section{
						Level:    1,
						Title:    section1Title,
						Elements: []interface{}{},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("section level 0 with nested section level 2", func() {
			actualContent := "= a header\n" +
				"\n" +
				"=== section 2"
			doctitle := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID: "_a_header",
				},
				Content: types.InlineElements{
					types.StringElement{Content: "a header"},
				},
			}
			section2Title := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID: "_section_2",
				},
				Content: types.InlineElements{
					types.StringElement{Content: "section 2"},
				},
			}
			expectedResult := types.Document{
				Attributes: types.DocumentAttributes{
					"doctitle": doctitle,
				},
				ElementReferences: map[string]interface{}{
					"_section_2": section2Title,
				},
				Elements: []interface{}{
					types.Section{
						Level:    2,
						Title:    section2Title,
						Elements: []interface{}{},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("section level 1 with immediate paragraph", func() {
			actualContent := `== a title
and a paragraph`
			section1Title := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID: "_a_title",
				},
				Content: types.InlineElements{
					types.StringElement{Content: "a title"},
				},
			}
			expectedResult := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: map[string]interface{}{
					"_a_title": section1Title,
				},
				Elements: []interface{}{
					types.Section{
						Level: 1,
						Title: section1Title,
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
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("section level 1 with a paragraph separated by empty line", func() {
			actualContent := `== a title
			
and a paragraph`
			section1Title := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID: "_a_title",
				},
				Content: types.InlineElements{
					types.StringElement{Content: "a title"},
				},
			}
			expectedResult := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: map[string]interface{}{
					"_a_title": section1Title,
				},
				Elements: []interface{}{
					types.Section{
						Level: 1,
						Title: section1Title,
						Elements: []interface{}{
							types.BlankLine{},
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
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("section level 1 with a paragraph separated by non-empty line", func() {
			actualContent := "== a title\n    \nand a paragraph"
			section1Title := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID: "_a_title",
				},
				Content: types.InlineElements{
					types.StringElement{Content: "a title"},
				},
			}
			expectedResult := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: map[string]interface{}{
					"_a_title": section1Title,
				},
				Elements: []interface{}{
					types.Section{
						Level: 1,
						Title: section1Title,
						Elements: []interface{}{
							types.BlankLine{},
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
			verify(GinkgoT(), expectedResult, actualContent)
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
					types.AttrID: "_a_header",
				},
				Content: types.InlineElements{
					types.StringElement{Content: "a header"},
				},
			}
			sectionATitle := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID: "_section_a",
				},
				Content: types.InlineElements{
					types.StringElement{Content: "Section A"},
				},
			}
			sectionAaTitle := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID: "_section_a_a",
				},
				Content: types.InlineElements{
					types.StringElement{Content: "Section A.a"},
				},
			}
			sectionBTitle := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID: "_section_b",
				},
				Content: types.InlineElements{
					types.StringElement{Content: "Section B"},
				},
			}
			expectedResult := types.Document{
				Attributes: types.DocumentAttributes{
					"doctitle": doctitle,
				},
				ElementReferences: map[string]interface{}{
					"_section_a":   sectionATitle,
					"_section_a_a": sectionAaTitle,
					"_section_b":   sectionBTitle,
				},
				Elements: []interface{}{
					types.Section{
						Level: 1,
						Title: sectionATitle,
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
								Level: 2,
								Title: sectionAaTitle,
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
						Level: 1,
						Title: sectionBTitle,
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
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("single section with custom IDs", func() {
			actualContent := `[[custom_header]]
== a header`
			sectionTitle := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID: "custom_header",
				},
				Content: types.InlineElements{
					types.StringElement{Content: "a header"},
				},
			}
			expectedResult := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: map[string]interface{}{
					"custom_header": sectionTitle,
				},
				Elements: []interface{}{
					types.Section{
						Level:    1,
						Title:    sectionTitle,
						Elements: []interface{}{},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("multiple sections with custom IDs", func() {
			actualContent := `[[custom_header]]
= a header

== Section F [[foo]]

[[bar]]
== Section B
a paragraph`
			doctitle := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID: "custom_header",
				},
				Content: types.InlineElements{
					types.StringElement{Content: "a header"},
				},
			}
			fooTitle := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID: "foo",
				},
				Content: types.InlineElements{
					types.StringElement{Content: "Section F "},
				},
			}
			barTitle := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID: "bar",
				},
				Content: types.InlineElements{
					types.StringElement{Content: "Section B"},
				},
			}
			expectedResult := types.Document{
				Attributes: types.DocumentAttributes{
					"doctitle": doctitle,
				},
				ElementReferences: map[string]interface{}{
					"foo": fooTitle,
					"bar": barTitle,
				},
				Elements: []interface{}{
					types.Section{
						Level: 1,
						Title: fooTitle,
						Elements: []interface{}{
							types.BlankLine{},
						},
					},
					types.Section{
						Level: 1,
						Title: barTitle,
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
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

	Context("invalid sections", func() {
		It("header invalid - missing space", func() {
			actualContent := "=a header"
			expectedResult := types.Document{
				Attributes:        map[string]interface{}{},
				ElementReferences: map[string]interface{}{},
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
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("header invalid - header space", func() {
			actualContent := " = a header"
			expectedResult := types.Document{
				Attributes:        map[string]interface{}{},
				ElementReferences: map[string]interface{}{},
				Elements: []interface{}{
					types.LiteralBlock{
						Content: " = a header",
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("header with invalid section1", func() {
			actualContent := "= a header\n" +
				"\n" +
				" == section 1"
			expectedResult := types.Document{
				Attributes: types.DocumentAttributes{
					"doctitle": types.SectionTitle{
						Attributes: types.ElementAttributes{
							types.AttrID: "_a_header",
						},
						Content: types.InlineElements{
							types.StringElement{Content: "a header"},
						},
					},
				},
				ElementReferences: map[string]interface{}{},
				Elements: []interface{}{
					types.LiteralBlock{
						Content: " == section 1",
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("multiple sections level 0", func() {
			actualContent := `= header 1

foo

= header 2
bar`
			header2Title := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID: "_header_2",
				},
				Content: types.InlineElements{
					types.StringElement{Content: "header 2"},
				},
			}
			expectedResult := types.Document{
				Attributes: types.DocumentAttributes{
					"doctitle": types.SectionTitle{
						Attributes: types.ElementAttributes{
							types.AttrID: "_header_1",
						},
						Content: types.InlineElements{
							types.StringElement{Content: "header 1"},
						},
					},
				},
				ElementReferences: map[string]interface{}{
					"_header_2": header2Title,
				},

				Elements: []interface{}{
					types.Preamble{
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "foo",
										},
									},
								},
							},
							types.BlankLine{},
						},
					},
					types.Section{
						Level: 0,
						Title: header2Title,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "bar",
										},
									},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})
})
