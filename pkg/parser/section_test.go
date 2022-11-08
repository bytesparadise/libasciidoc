package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("sections", func() {

	Context("in final documents", func() {

		Context("valid sections", func() {

			It("header only", func() {
				source := "= a header"
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Title: []interface{}{
								&types.StringElement{
									Content: "a header",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("header with trailing spaces", func() {
				source := "= a header  "
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Title: []interface{}{
								&types.StringElement{
									Content: "a header",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("section with trailing spaces", func() {
				source := "== a section  "
				expected := &types.Document{
					Elements: []interface{}{
						&types.Section{
							Level: 1,
							Attributes: types.Attributes{
								types.AttrID: "_a_section",
							},
							Title: []interface{}{
								&types.StringElement{
									Content: "a section",
								},
							},
						},
					},
					ElementReferences: types.ElementReferences{
						"_a_section": []interface{}{
							&types.StringElement{
								Content: "a section",
							},
						},
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_a_section",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("header and paragraph", func() {
				source := `= a header

and a paragraph`

				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Title: []interface{}{
								&types.StringElement{Content: "a header"},
							},
						},
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "and a paragraph"},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("section with link in title", func() {
				source := `== link to https://foo.bar
`
				section1aTitle := []interface{}{
					&types.StringElement{Content: "link to "},
					&types.InlineLink{
						Location: &types.Location{
							Scheme: "https://",
							Path:   "foo.bar",
						},
					},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_link_to_httpsfoo_bar",
							},
							Level: 1,
							Title: section1aTitle,
						},
					},
					ElementReferences: types.ElementReferences{
						"_link_to_httpsfoo_bar": section1aTitle,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_link_to_httpsfoo_bar",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("header, section 1 and paragraph with bold quote", func() {
				source := `= a header
				
== section 1

a paragraph with *bold content*`

				section1Title := []interface{}{
					&types.StringElement{Content: "section 1"},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Title: []interface{}{
								&types.StringElement{Content: "a header"},
							},
						},
						&types.Section{
							Level: 1,
							Title: section1Title,
							Attributes: types.Attributes{
								types.AttrID: "_section_1",
							},
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "a paragraph with "},
										&types.QuotedText{
											Kind: types.SingleQuoteBold,
											Elements: []interface{}{
												&types.StringElement{Content: "bold content"},
											},
										},
									},
								},
							},
						},
					},
					ElementReferences: types.ElementReferences{
						"_section_1": section1Title,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_section_1",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("two sections with level 0", func() {
				source := `= a first header

= a second header`
				otherDoctitle := []interface{}{
					&types.StringElement{Content: "a second header"},
				}

				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Title: []interface{}{
								&types.StringElement{Content: "a first header"},
							},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_a_second_header",
							},
							Level: 0,
							Title: otherDoctitle,
						},
					},
					ElementReferences: types.ElementReferences{
						"_a_second_header": otherDoctitle,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_a_second_header",
								Level: 0,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("section level 1 alone", func() {
				source := `== section 1`
				section1Title := []interface{}{
					&types.StringElement{Content: "section 1"},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_section_1",
							},
							Level: 1,
							Title: section1Title,
						},
					},
					ElementReferences: types.ElementReferences{
						"_section_1": section1Title,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_section_1",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("section level 1 with quoted text", func() {
				source := `==  *2 spaces and bold content*`
				sectionTitle := []interface{}{
					&types.QuotedText{
						Kind: types.SingleQuoteBold,
						Elements: []interface{}{
							&types.StringElement{Content: "2 spaces and bold content"},
						},
					},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_2_spaces_and_bold_content",
							},
							Level: 1,
							Title: sectionTitle,
						},
					},
					ElementReferences: types.ElementReferences{
						"_2_spaces_and_bold_content": sectionTitle,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_2_spaces_and_bold_content",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("section level 0 with nested section level 1", func() {
				source := `= a header

== section 1`
				section1Title := []interface{}{
					&types.StringElement{Content: "section 1"},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Title: []interface{}{
								&types.StringElement{Content: "a header"},
							},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_section_1",
							},
							Level: 1,
							Title: section1Title,
						},
					},
					ElementReferences: types.ElementReferences{
						"_section_1": section1Title,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_section_1",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("section level 0 with nested section level 2", func() {
				source := `= a header

=== section 2`
				section2Title := []interface{}{
					&types.StringElement{Content: "section 2"},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Title: []interface{}{
								&types.StringElement{Content: "a header"},
							},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_section_2",
							},
							Level: 2,
							Title: section2Title,
						},
					},
					ElementReferences: types.ElementReferences{
						"_section_2": section2Title,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_section_2",
								Level: 2,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("section level 1 with immediate paragraph", func() {
				source := `== a title
and a paragraph`
				section1Title := []interface{}{
					&types.StringElement{Content: "a title"},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_a_title",
							},
							Level: 1,
							Title: section1Title,
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "and a paragraph"},
									},
								},
							},
						},
					},
					ElementReferences: types.ElementReferences{
						"_a_title": section1Title,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_a_title",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("section level 1 with a paragraph separated by empty line", func() {
				source := `== a title
			
and a paragraph`
				section1Title := []interface{}{
					&types.StringElement{Content: "a title"},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_a_title",
							},
							Level: 1,
							Title: section1Title,
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "and a paragraph"},
									},
								},
							},
						},
					},
					ElementReferences: types.ElementReferences{
						"_a_title": section1Title,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_a_title",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("section level 1 with a paragraph separated by non-empty line", func() {
				source := "== a title\n    \nand a paragraph"
				section1Title := []interface{}{
					&types.StringElement{Content: "a title"},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_a_title",
							},
							Level: 1,
							Title: section1Title,
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "and a paragraph"},
									},
								},
							},
						},
					},
					ElementReferences: types.ElementReferences{
						"_a_title": section1Title,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_a_title",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("section levels 0, 1, 2, 1", func() {
				source := `= a header

== Section A
a paragraph

=== Section A.a
a paragraph

== Section B
a paragraph`
				sectionATitle := []interface{}{
					&types.StringElement{Content: "Section A"},
				}
				sectionAaTitle := []interface{}{
					&types.StringElement{Content: "Section A.a"},
				}
				sectionBTitle := []interface{}{
					&types.StringElement{Content: "Section B"},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Title: []interface{}{
								&types.StringElement{Content: "a header"},
							},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_Section_A",
							},
							Level: 1,
							Title: sectionATitle,
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "a paragraph"},
									},
								},
								&types.Section{
									Attributes: types.Attributes{
										types.AttrID: "_Section_A_a",
									},
									Level: 2,
									Title: sectionAaTitle,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{Content: "a paragraph"},
											},
										},
									},
								},
							},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_Section_B",
							},
							Level: 1,
							Title: sectionBTitle,
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "a paragraph"},
									},
								},
							},
						},
					},
					ElementReferences: types.ElementReferences{
						"_Section_A":   sectionATitle,
						"_Section_A_a": sectionAaTitle,
						"_Section_B":   sectionBTitle,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_Section_A",
								Level: 1,
								Children: []*types.ToCSection{
									{
										ID:    "_Section_A_a",
										Level: 2,
									},
								},
							},
							{
								ID:    "_Section_B",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("section levels 0, 1, 2, 2", func() {
				source := `= a header

== Section A
a paragraph

=== Section A.a
a paragraph

=== Section A.b
a paragraph`
				sectionATitle := []interface{}{
					&types.StringElement{Content: "Section A"},
				}
				sectionAaTitle := []interface{}{
					&types.StringElement{Content: "Section A.a"},
				}
				sectionBTitle := []interface{}{
					&types.StringElement{Content: "Section A.b"},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Title: []interface{}{
								&types.StringElement{Content: "a header"},
							},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_Section_A",
							},
							Level: 1,
							Title: sectionATitle,
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "a paragraph"},
									},
								},
								&types.Section{
									Attributes: types.Attributes{
										types.AttrID: "_Section_A_a",
									},
									Level: 2,
									Title: sectionAaTitle,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{Content: "a paragraph"},
											},
										},
									},
								},
								&types.Section{
									Attributes: types.Attributes{
										types.AttrID: "_Section_A_b",
									},
									Level: 2,
									Title: sectionBTitle,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{Content: "a paragraph"},
											},
										},
									},
								},
							},
						},
					},
					ElementReferences: types.ElementReferences{
						"_Section_A":   sectionATitle,
						"_Section_A_a": sectionAaTitle,
						"_Section_A_b": sectionBTitle,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_Section_A",
								Level: 1,
								Children: []*types.ToCSection{
									{
										ID:    "_Section_A_a",
										Level: 2,
									},
									{
										ID:    "_Section_A_b",
										Level: 2,
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("section levels 0, 2, 3, 3", func() {
				source := `= a header

=== Section A
a paragraph

==== Section A.a
a paragraph

==== Section A.b
a paragraph`
				sectionATitle := []interface{}{
					&types.StringElement{Content: "Section A"},
				}
				sectionAaTitle := []interface{}{
					&types.StringElement{Content: "Section A.a"},
				}
				sectionAbTitle := []interface{}{
					&types.StringElement{Content: "Section A.b"},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Title: []interface{}{
								&types.StringElement{Content: "a header"},
							},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_Section_A",
							},
							Level: 2,
							Title: sectionATitle,
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "a paragraph"},
									},
								},
								&types.Section{
									Attributes: types.Attributes{
										types.AttrID: "_Section_A_a",
									},
									Level: 3,
									Title: sectionAaTitle,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{Content: "a paragraph"},
											},
										},
									},
								},
								&types.Section{
									Attributes: types.Attributes{
										types.AttrID: "_Section_A_b",
									},
									Level: 3, // level is adjusted
									Title: sectionAbTitle,
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{Content: "a paragraph"},
											},
										},
									},
								},
							},
						},
					},
					ElementReferences: types.ElementReferences{
						"_Section_A":   sectionATitle,
						"_Section_A_a": sectionAaTitle,
						"_Section_A_b": sectionAbTitle,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_Section_A",
								Level: 2,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("section levels 0, 2, 2, 2", func() {
				source := `= a header

=== Section A
a paragraph

=== Section B
a paragraph

=== Section C
a paragraph`
				sectionATitle := []interface{}{
					&types.StringElement{Content: "Section A"},
				}
				sectionBTitle := []interface{}{
					&types.StringElement{Content: "Section B"},
				}
				sectionCTitle := []interface{}{
					&types.StringElement{Content: "Section C"},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Title: []interface{}{
								&types.StringElement{Content: "a header"},
							},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_Section_A",
							},
							Level: 2,
							Title: sectionATitle,
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "a paragraph"},
									},
								},
							},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_Section_B",
							},
							Level: 2,
							Title: sectionBTitle,
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "a paragraph"},
									},
								},
							},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_Section_C",
							},
							Level: 2,
							Title: sectionCTitle,
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "a paragraph"},
									},
								},
							},
						},
					},
					ElementReferences: types.ElementReferences{
						"_Section_A": sectionATitle,
						"_Section_B": sectionBTitle,
						"_Section_C": sectionCTitle,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_Section_A",
								Level: 2,
							},
							{
								ID:    "_Section_B",
								Level: 2,
							},
							{
								ID:    "_Section_C",
								Level: 2,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("single with custom block ID", func() {
				source := `[[custom_header]]
== a header`
				sectionTitle := []interface{}{
					&types.StringElement{Content: "a header"},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "custom_header",
							},
							Level: 1,
							Title: sectionTitle,
						},
					},
					ElementReferences: types.ElementReferences{
						"custom_header": sectionTitle,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "custom_header",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("single with custom inline ID", func() {
				source := `== a header [[custom_header]]`
				sectionTitle := []interface{}{
					&types.StringElement{
						Content: "a header",
					},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "custom_header",
							},
							Level: 1,
							Title: sectionTitle,
						},
					},
					ElementReferences: types.ElementReferences{
						"custom_header": sectionTitle,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "custom_header",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("single with attached inline anchor", func() {
				source := `== a header[[bookmark]]`
				sectionTitle := []interface{}{
					&types.StringElement{Content: "a header"},
					&types.InlineLink{
						Attributes: types.Attributes{
							types.AttrID: "bookmark",
						},
					},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_a_header",
							},
							Level: 1,
							Title: sectionTitle,
						},
					},
					ElementReferences: types.ElementReferences{
						"_a_header": sectionTitle,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_a_header",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("single with attached inline anchor and inline ID", func() {
				source := `== a header[[bookmark]] [[custom_header]]`
				sectionTitle := []interface{}{
					&types.StringElement{Content: "a header"},
					&types.InlineLink{
						Attributes: types.Attributes{
							types.AttrID: "bookmark",
						},
					},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "custom_header",
							},
							Level: 1,
							Title: sectionTitle,
						},
					},
					ElementReferences: types.ElementReferences{
						"custom_header": sectionTitle,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "custom_header",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("single with detached inline anchor and inline ID", func() {
				source := `== a header [[bookmark]] [[custom_header]]`
				sectionTitle := []interface{}{
					&types.StringElement{Content: "a header "},
					&types.InlineLink{
						Attributes: types.Attributes{
							types.AttrID: "bookmark",
						},
					},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "custom_header",
							},
							Level: 1,
							Title: sectionTitle,
						},
					},
					ElementReferences: types.ElementReferences{
						"custom_header": sectionTitle,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "custom_header",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("single line with apostrophe", func() {
				source := `== ...and we're back!`
				sectionTitle := []interface{}{
					&types.Symbol{
						Name: "...",
					},
					&types.StringElement{
						Content: "and we",
					},
					&types.Symbol{
						Name: "'",
					},
					&types.StringElement{
						Content: "re back!",
					},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_and_were_back",
							},
							Level: 1,
							Title: sectionTitle,
						},
					},
					ElementReferences: types.ElementReferences{
						"_and_were_back": sectionTitle,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_and_were_back",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("multiple sections with multiple inline custom IDs", func() {
				source := `[[custom_header]]
= a header

== Section F [[ignored]] [[foo]]

[[bar]]
== Section B
a paragraph`
				fooTitle := []interface{}{
					&types.StringElement{Content: "Section F "},
					&types.InlineLink{
						Attributes: types.Attributes{
							types.AttrID: "ignored",
						},
					},
				}
				barTitle := []interface{}{
					&types.StringElement{Content: "Section B"},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Title: []interface{}{
								&types.StringElement{Content: "a header"},
							},
							Attributes: types.Attributes{
								types.AttrID: "custom_header",
							},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "foo",
							},
							Level: 1,
							Title: fooTitle,
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "bar",
							},
							Level: 1,
							Title: barTitle,
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "a paragraph"},
									},
								},
							},
						},
					},
					ElementReferences: types.ElementReferences{
						"foo": fooTitle,
						"bar": barTitle,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "foo",
								Level: 1,
							},
							{
								ID:    "bar",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("sections with same title", func() {
				source := `== section 1

== section 1`
				section1aTitle := []interface{}{
					&types.StringElement{Content: "section 1"},
				}
				section1bTitle := []interface{}{
					&types.StringElement{Content: "section 1"},
				}

				expected := &types.Document{
					Elements: []interface{}{
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_section_1",
							},
							Level: 1,
							Title: section1aTitle,
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_section_1_2",
							},
							Level: 1,
							Title: section1bTitle,
						},
					},
					ElementReferences: types.ElementReferences{
						"_section_1":   section1aTitle,
						"_section_1_2": section1bTitle,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_section_1",
								Level: 1,
							},
							{
								ID:    "_section_1_2",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("section level 0 with nested section level 1 and custom ID prefix", func() {
				source := `= a header
:idprefix: custom_

== section 1`
				section1Title := []interface{}{
					&types.StringElement{Content: "section 1"},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Title: []interface{}{
								&types.StringElement{Content: "a header"},
							},
							Elements: []interface{}{
								&types.AttributeDeclaration{
									Name:  types.AttrIDPrefix,
									Value: "custom_",
								},
							},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "custom_section_1",
							},
							Level: 1,
							Title: section1Title,
						},
					},
					ElementReferences: types.ElementReferences{
						"custom_section_1": section1Title,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "custom_section_1",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("section level 0 with nested sections level 1 and custom ID prefixes - with idprefix as doc attribute", func() {
				source := `= a header
:idprefix: custom1a_

== section 1a

:idprefix: custom1b_

== section 1b`
				section1aTitle := []interface{}{
					&types.StringElement{Content: "section 1a"},
				}
				section1bTitle := []interface{}{
					&types.StringElement{Content: "section 1b"},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Title: []interface{}{
								&types.StringElement{Content: "a header"},
							},
							Elements: []interface{}{
								&types.AttributeDeclaration{
									Name:  "idprefix",
									Value: "custom1a_",
								},
							},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "custom1a_section_1a",
							},
							Level: 1,
							Title: section1aTitle,
							Elements: []interface{}{
								&types.AttributeDeclaration{
									Name:  "idprefix",
									Value: "custom1b_",
								},
							},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "custom1b_section_1b",
							},
							Level: 1,
							Title: section1bTitle,
						},
					},
					ElementReferences: types.ElementReferences{
						"custom1a_section_1a": section1aTitle,
						"custom1b_section_1b": section1bTitle,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "custom1a_section_1a",
								Level: 1,
							},
							{
								ID:    "custom1b_section_1b",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("section level 0 with nested sections level 1 and custom ID prefixes - without idprefix as doc attribute", func() {
				source := `= a header

:idprefix: custom1a_

== section 1a

:idprefix: custom1b_

== section 1b`
				section1aTitle := []interface{}{
					&types.StringElement{Content: "section 1a"},
				}
				section1bTitle := []interface{}{
					&types.StringElement{Content: "section 1b"},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Title: []interface{}{
								&types.StringElement{Content: "a header"},
							},
						},
						&types.AttributeDeclaration{ // not in the DocumentHeader because of the blankline in-between
							Name:  "idprefix",
							Value: "custom1a_",
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "custom1a_section_1a",
							},
							Level: 1,
							Title: section1aTitle,
							Elements: []interface{}{
								&types.AttributeDeclaration{
									Name:  "idprefix",
									Value: "custom1b_",
								},
							},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "custom1b_section_1b",
							},
							Level: 1,
							Title: section1bTitle,
						},
					},
					ElementReferences: types.ElementReferences{
						"custom1a_section_1a": section1aTitle,
						"custom1b_section_1b": section1bTitle,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "custom1a_section_1a",
								Level: 1,
							},
							{
								ID:    "custom1b_section_1b",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("header with preamble then section level 1", func() {
				source := `= A Title

a short preamble

== Section 1`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Title: []interface{}{
								&types.StringElement{Content: "A Title"},
							},
						},
						&types.Preamble{
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "a short preamble"},
									},
								},
							},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_Section_1",
							},
							Level: 1,
							Title: []interface{}{
								&types.StringElement{Content: "Section 1"},
							},
						},
					},
					ElementReferences: types.ElementReferences{
						"_Section_1": []interface{}{
							&types.StringElement{Content: "Section 1"},
						},
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_Section_1",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("front-matter on top of header", func() {
				source := `---
draft: true
---
				
= A Title

a short preamble

== Section 1`
				expected := &types.Document{
					Elements: []interface{}{
						&types.FrontMatter{
							Attributes: map[string]interface{}{
								"draft": true,
							},
						},
						&types.DocumentHeader{
							Title: []interface{}{
								&types.StringElement{Content: "A Title"},
							},
						},
						&types.Preamble{
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{Content: "a short preamble"},
									},
								},
							},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_Section_1",
							},
							Level: 1,
							Title: []interface{}{
								&types.StringElement{Content: "Section 1"},
							},
						},
					},
					ElementReferences: types.ElementReferences{
						"_Section_1": []interface{}{
							&types.StringElement{Content: "Section 1"},
						},
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_Section_1",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

		})

		Context("invalid sections", func() {

			It("header invalid - too many spaces", func() {
				source := "======= a header"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "======= a header"},
							},
						},
					}}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("header invalid - missing space", func() {
				source := "=a header"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "=a header"},
							},
						},
					}}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("header invalid - header space", func() {
				source := " = a header with a prefix space"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Attributes: types.Attributes{
								types.AttrStyle: types.LiteralParagraph,
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: " = a header with a prefix space", // spaces on first line of literal paragraphs are NOT trimmed by parser
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("header with invalid section1", func() {
				source := `= a header

 == section with prefix space`
				title := []interface{}{
					&types.StringElement{Content: "a header"},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Title: title,
						},
						&types.Paragraph{
							Attributes: types.Attributes{
								types.AttrStyle: types.LiteralParagraph,
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: " == section with prefix space", // spaces on first line of literal paragraphs are NOT trimmed by parser
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

		})

		Context("unsupported section syntax", func() {

			It("match unclosed example section", func() {
				source := `Document Title
==============
Doc Writer <thedoc@asciidoctor.org>`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "Document Title",
								},
							},
						},
						&types.DelimitedBlock{
							Kind: types.Example,
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											Content: "Doc Writer ",
										},
										&types.SpecialCharacter{
											Name: "<",
										},
										&types.InlineLink{
											Location: &types.Location{
												Scheme: "mailto:",
												Path:   "thedoc@asciidoctor.org",
											},
										},
										&types.SpecialCharacter{
											Name: ">",
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})
	})
})
