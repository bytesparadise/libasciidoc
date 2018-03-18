package parser_test

import (
	"github.com/bytesparadise/libasciidoc/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("sections", func() {

	Context("valid sections", func() {

		It("header only", func() {
			actualContent := "= a header"
			doctitle := types.SectionTitle{
				ID: types.ElementID{
					Value: "_a_header",
				},
				Content: types.InlineContent{
					Elements: []types.InlineElement{
						types.StringElement{Content: "a header"},
					},
				},
			}
			expectedResult := types.Document{
				Attributes: map[string]interface{}{
					"doctitle": doctitle,
				},
				ElementReferences: map[string]interface{}{},
				Elements:          []types.DocElement{},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("header and paragraph", func() {
			actualContent := `= a header

and a paragraph`

			doctitle := types.SectionTitle{
				ID: types.ElementID{
					Value: "_a_header",
				},
				Content: types.InlineContent{
					Elements: []types.InlineElement{
						types.StringElement{Content: "a header"},
					},
				},
			}
			expectedResult := types.Document{
				Attributes: map[string]interface{}{
					"doctitle": doctitle,
				},
				ElementReferences: map[string]interface{}{},
				Elements: []types.DocElement{
					types.Paragraph{
						Lines: []types.InlineContent{
							{
								Elements: []types.InlineElement{
									types.StringElement{Content: "and a paragraph"},
								},
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
				ID: types.ElementID{
					Value: "_section_1",
				},
				Content: types.InlineContent{
					Elements: []types.InlineElement{
						types.StringElement{Content: "section 1"},
					},
				},
			}
			expectedResult := types.Document{
				Attributes: map[string]interface{}{},
				ElementReferences: map[string]interface{}{
					"_section_1": section1Title,
				},
				Elements: []types.DocElement{
					types.Section{
						Level:    1,
						Title:    section1Title,
						Elements: []types.DocElement{},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("section level 1 with quoted text", func() {
			actualContent := `==  *2 spaces and bold content*`
			sectionTitle := types.SectionTitle{
				ID: types.ElementID{
					Value: "__strong_2_spaces_and_bold_content_strong",
				},
				Content: types.InlineContent{
					Elements: []types.InlineElement{
						types.QuotedText{
							Elements: []types.InlineElement{
								types.StringElement{Content: "2 spaces and bold content"},
							},
						},
					},
				},
			}
			expectedResult := types.Document{
				Attributes: map[string]interface{}{},
				ElementReferences: map[string]interface{}{
					"__strong_2_spaces_and_bold_content_strong": sectionTitle,
				},
				Elements: []types.DocElement{
					types.Section{
						Level:    1,
						Title:    sectionTitle,
						Elements: []types.DocElement{},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("section level 0 with nested section level 1", func() {
			actualContent := `= a header

== section 1`
			doctitle := types.SectionTitle{
				ID: types.ElementID{
					Value: "_a_header",
				},
				Content: types.InlineContent{
					Elements: []types.InlineElement{
						types.StringElement{Content: "a header"},
					},
				},
			}
			section1Title := types.SectionTitle{
				ID: types.ElementID{
					Value: "_section_1",
				},
				Content: types.InlineContent{
					Elements: []types.InlineElement{
						types.StringElement{Content: "section 1"},
					},
				},
			}
			expectedResult := types.Document{
				Attributes: map[string]interface{}{
					"doctitle": doctitle,
				},
				ElementReferences: map[string]interface{}{
					"_section_1": section1Title,
				},
				Elements: []types.DocElement{
					types.Section{
						Level:    1,
						Title:    section1Title,
						Elements: []types.DocElement{},
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
				ID: types.ElementID{
					Value: "_section_1",
				},
				Content: types.InlineContent{
					Elements: []types.InlineElement{
						types.StringElement{Content: "section 1"},
					},
				},
			}
			doctitle := types.SectionTitle{
				ID: types.ElementID{
					Value: "_a_header",
				},
				Content: types.InlineContent{
					Elements: []types.InlineElement{
						types.StringElement{Content: "a header"},
					},
				},
			}
			expectedResult := types.Document{
				Attributes: map[string]interface{}{
					"doctitle": doctitle,
				},
				ElementReferences: map[string]interface{}{
					"_section_1": section1Title,
				},
				Elements: []types.DocElement{
					types.Preamble{
						Elements: []types.DocElement{
							types.Paragraph{
								Lines: []types.InlineContent{
									{
										Elements: []types.InlineElement{
											types.StringElement{Content: "a short preamble"},
										},
									},
								},
							},
						},
					},
					types.Section{
						Level:    1,
						Title:    section1Title,
						Elements: []types.DocElement{},
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
				ID: types.ElementID{
					Value: "_a_header",
				},
				Content: types.InlineContent{
					Elements: []types.InlineElement{
						types.StringElement{Content: "a header"},
					},
				},
			}
			section2Title := types.SectionTitle{
				ID: types.ElementID{
					Value: "_section_2",
				},
				Content: types.InlineContent{
					Elements: []types.InlineElement{
						types.StringElement{Content: "section 2"},
					},
				},
			}
			expectedResult := types.Document{
				Attributes: map[string]interface{}{
					"doctitle": doctitle,
				},
				ElementReferences: map[string]interface{}{
					"_section_2": section2Title,
				},
				Elements: []types.DocElement{
					types.Section{
						Level:    2,
						Title:    section2Title,
						Elements: []types.DocElement{},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("section level 1 with immediate paragraph", func() {
			actualContent := `== a title
and a paragraph`
			section1Title := types.SectionTitle{
				ID: types.ElementID{
					Value: "_a_title",
				},
				Content: types.InlineContent{
					Elements: []types.InlineElement{
						types.StringElement{Content: "a title"},
					},
				},
			}
			expectedResult := types.Document{
				Attributes: map[string]interface{}{},
				ElementReferences: map[string]interface{}{
					"_a_title": section1Title,
				},
				Elements: []types.DocElement{
					types.Section{
						Level: 1,
						Title: section1Title,
						Elements: []types.DocElement{
							types.Paragraph{
								Lines: []types.InlineContent{
									{
										Elements: []types.InlineElement{
											types.StringElement{Content: "and a paragraph"},
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

		It("section level 1 with a paragraph separated by empty line", func() {
			actualContent := `== a title
			
and a paragraph`
			section1Title := types.SectionTitle{
				ID: types.ElementID{
					Value: "_a_title",
				},
				Content: types.InlineContent{
					Elements: []types.InlineElement{
						types.StringElement{Content: "a title"},
					},
				},
			}
			expectedResult := types.Document{
				Attributes: map[string]interface{}{},
				ElementReferences: map[string]interface{}{
					"_a_title": section1Title,
				},
				Elements: []types.DocElement{
					types.Section{
						Level: 1,
						Title: section1Title,
						Elements: []types.DocElement{
							types.Paragraph{
								Lines: []types.InlineContent{
									{
										Elements: []types.InlineElement{
											types.StringElement{Content: "and a paragraph"},
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

		It("section level 1 with a paragraph separated by non-empty line", func() {
			actualContent := "== a title\n    \nand a paragraph"
			section1Title := types.SectionTitle{
				ID: types.ElementID{
					Value: "_a_title",
				},
				Content: types.InlineContent{
					Elements: []types.InlineElement{
						types.StringElement{Content: "a title"},
					},
				},
			}
			expectedResult := types.Document{
				Attributes: map[string]interface{}{},
				ElementReferences: map[string]interface{}{
					"_a_title": section1Title,
				},
				Elements: []types.DocElement{
					types.Section{
						Level: 1,
						Title: section1Title,
						Elements: []types.DocElement{
							types.Paragraph{
								Lines: []types.InlineContent{
									{
										Elements: []types.InlineElement{
											types.StringElement{Content: "and a paragraph"},
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

		It("section levels 1, 2, 3, 2", func() {
			actualContent := `= a header

== Section A
a paragraph

=== Section A.a
a paragraph

== Section B
a paragraph`
			doctitle := types.SectionTitle{
				ID: types.ElementID{
					Value: "_a_header",
				},
				Content: types.InlineContent{
					Elements: []types.InlineElement{
						types.StringElement{Content: "a header"},
					},
				},
			}
			sectionATitle := types.SectionTitle{
				ID: types.ElementID{
					Value: "_section_a",
				},
				Content: types.InlineContent{
					Elements: []types.InlineElement{
						types.StringElement{Content: "Section A"},
					},
				},
			}
			sectionAaTitle := types.SectionTitle{
				ID: types.ElementID{
					Value: "_section_a_a",
				},
				Content: types.InlineContent{
					Elements: []types.InlineElement{
						types.StringElement{Content: "Section A.a"},
					},
				},
			}
			sectionBTitle := types.SectionTitle{
				ID: types.ElementID{
					Value: "_section_b",
				},
				Content: types.InlineContent{
					Elements: []types.InlineElement{
						types.StringElement{Content: "Section B"},
					},
				},
			}
			expectedResult := types.Document{
				Attributes: map[string]interface{}{
					"doctitle": doctitle,
				},
				ElementReferences: map[string]interface{}{
					"_section_a":   sectionATitle,
					"_section_a_a": sectionAaTitle,
					"_section_b":   sectionBTitle,
				},
				Elements: []types.DocElement{
					types.Section{
						Level: 1,
						Title: sectionATitle,
						Elements: []types.DocElement{
							types.Paragraph{
								Lines: []types.InlineContent{
									{
										Elements: []types.InlineElement{
											types.StringElement{Content: "a paragraph"},
										},
									},
								},
							},
							types.Section{
								Level: 2,
								Title: sectionAaTitle,
								Elements: []types.DocElement{
									types.Paragraph{
										Lines: []types.InlineContent{
											{
												Elements: []types.InlineElement{
													types.StringElement{Content: "a paragraph"},
												},
											},
										},
									},
								},
							},
						},
					},
					types.Section{
						Level: 1,
						Title: sectionBTitle,
						Elements: []types.DocElement{
							types.Paragraph{
								Lines: []types.InlineContent{
									{
										Elements: []types.InlineElement{
											types.StringElement{Content: "a paragraph"},
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

		It("section with IDs", func() {
			actualContent := `= a header

== Section F [[foo]]

[[bar]]
== Section B
a paragraph`
			doctitle := types.SectionTitle{
				ID: types.ElementID{
					Value: "_a_header",
				},
				Content: types.InlineContent{
					Elements: []types.InlineElement{
						types.StringElement{Content: "a header"},
					},
				},
			}
			fooTitle := types.SectionTitle{
				ID: types.ElementID{
					Value: "foo",
				},
				Content: types.InlineContent{
					Elements: []types.InlineElement{
						types.StringElement{Content: "Section F"},
					},
				},
			}
			barTitle := types.SectionTitle{
				ID: types.ElementID{
					Value: "bar",
				},
				Content: types.InlineContent{
					Elements: []types.InlineElement{
						types.StringElement{Content: "Section B"},
					},
				},
			}
			expectedResult := types.Document{
				Attributes: map[string]interface{}{
					"doctitle": doctitle,
				},
				ElementReferences: map[string]interface{}{
					"foo": fooTitle,
					"bar": barTitle,
				},
				Elements: []types.DocElement{
					types.Section{
						Level:    1,
						Title:    fooTitle,
						Elements: []types.DocElement{},
					},
					types.Section{
						Level: 1,
						Title: barTitle,
						Elements: []types.DocElement{
							types.Paragraph{
								Lines: []types.InlineContent{
									{
										Elements: []types.InlineElement{
											types.StringElement{Content: "a paragraph"},
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

	Context("invalid sections", func() {
		It("header invalid - missing space", func() {
			actualContent := "=a header"
			expectedResult := types.Document{
				Attributes:        map[string]interface{}{},
				ElementReferences: map[string]interface{}{},
				Elements: []types.DocElement{
					types.Paragraph{
						Lines: []types.InlineContent{
							{
								Elements: []types.InlineElement{
									types.StringElement{Content: "=a header"},
								},
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
				Elements: []types.DocElement{
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
				Attributes: map[string]interface{}{
					"doctitle": types.SectionTitle{
						ID: types.ElementID{
							Value: "_a_header",
						},
						Content: types.InlineContent{
							Elements: []types.InlineElement{
								types.StringElement{Content: "a header"},
							},
						},
					},
				},
				ElementReferences: map[string]interface{}{},
				Elements: []types.DocElement{
					types.LiteralBlock{
						Content: " == section 1",
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})

	})
})
