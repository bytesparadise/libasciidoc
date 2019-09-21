package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("sections - preflight", func() {

	Context("valid sections", func() {

		It("header only", func() {
			source := "= a header"
			doctitle := types.InlineElements{
				types.StringElement{Content: "a header"},
			}
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "a_header",
							types.AttrCustomID: false,
						},
						Level:    0,
						Title:    doctitle,
						Elements: []interface{}{},
					},
				},
			}
			Expect(source).To(BecomePreflightDocument(expected))
		})

		It("header with many spaces around content", func() {
			source := "= a header   "
			doctitle := types.InlineElements{
				types.StringElement{Content: "a header   "},
			}
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "a_header",
							types.AttrCustomID: false,
						},
						Level:    0,
						Title:    doctitle,
						Elements: []interface{}{},
					},
				},
			}
			Expect(source).To(BecomePreflightDocument(expected))
		})

		It("header and paragraph", func() {
			source := `= a header

and a paragraph`

			doctitle := types.InlineElements{
				types.StringElement{Content: "a header"},
			}
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "a_header",
							types.AttrCustomID: false,
						},
						Level:    0,
						Title:    doctitle,
						Elements: []interface{}{},
					},
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
			}
			Expect(source).To(BecomePreflightDocument(expected))
		})

		It("two sections with level 0", func() {
			source := `= a first header

= a second header`
			doctitle := types.InlineElements{
				types.StringElement{Content: "a first header"},
			}
			otherDoctitle := types.InlineElements{
				types.StringElement{Content: "a second header"},
			}

			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "a_first_header",
							types.AttrCustomID: false,
						},
						Level:    0,
						Title:    doctitle,
						Elements: []interface{}{},
					},
					types.BlankLine{},
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "a_second_header",
							types.AttrCustomID: false,
						},
						Level:    0,
						Title:    otherDoctitle,
						Elements: []interface{}{},
					},
				},
			}
			Expect(source).To(BecomePreflightDocument(expected))
		})

		It("section level 1 alone", func() {
			source := `== section 1`
			section1Title := types.InlineElements{
				types.StringElement{Content: "section 1"},
			}
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "section_1",
							types.AttrCustomID: false,
						},
						Level:    1,
						Title:    section1Title,
						Elements: []interface{}{},
					},
				},
			}
			Expect(source).To(BecomePreflightDocument(expected))
		})

		It("section level 1 with quoted text", func() {
			source := `==  *2 spaces and bold content*`
			sectionTitle := types.InlineElements{
				types.QuotedText{
					Kind: types.Bold,
					Elements: types.InlineElements{
						types.StringElement{Content: "2 spaces and bold content"},
					},
				},
			}
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "2_spaces_and_bold_content",
							types.AttrCustomID: false,
						},
						Level:    1,
						Title:    sectionTitle,
						Elements: []interface{}{},
					},
				},
			}
			Expect(source).To(BecomePreflightDocument(expected))
		})

		It("section level 0 with nested section level 1", func() {
			source := `= a header

== section 1`
			doctitle := types.InlineElements{
				types.StringElement{Content: "a header"},
			}
			section1Title := types.InlineElements{
				types.StringElement{Content: "section 1"},
			}
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "a_header",
							types.AttrCustomID: false,
						},
						Level:    0,
						Title:    doctitle,
						Elements: []interface{}{},
					},
					types.BlankLine{},
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "section_1",
							types.AttrCustomID: false,
						},
						Level:    1,
						Title:    section1Title,
						Elements: []interface{}{},
					},
				},
			}
			Expect(source).To(BecomePreflightDocument(expected))
		})

		It("section level 0 with nested section level 2", func() {
			source := `= a header

=== section 2`
			doctitle := types.InlineElements{
				types.StringElement{Content: "a header"},
			}
			section2Title := types.InlineElements{
				types.StringElement{Content: "section 2"},
			}
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "a_header",
							types.AttrCustomID: false,
						},
						Level:    0,
						Title:    doctitle,
						Elements: []interface{}{},
					},
					types.BlankLine{},
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "section_2",
							types.AttrCustomID: false,
						},
						Level:    2,
						Title:    section2Title,
						Elements: []interface{}{},
					},
				},
			}
			Expect(source).To(BecomePreflightDocument(expected))
		})

		It("section level 1 with immediate paragraph", func() {
			source := `== a title
and a paragraph`
			section1Title := types.InlineElements{
				types.StringElement{Content: "a title"},
			}
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "a_title",
							types.AttrCustomID: false,
						},
						Level:    1,
						Title:    section1Title,
						Elements: []interface{}{},
					},
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
			Expect(source).To(BecomePreflightDocument(expected))
		})

		It("section level 1 with a paragraph separated by empty line", func() {
			source := `== a title
			
and a paragraph`
			section1Title := types.InlineElements{
				types.StringElement{Content: "a title"},
			}
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "a_title",
							types.AttrCustomID: false,
						},
						Level:    1,
						Title:    section1Title,
						Elements: []interface{}{},
					},
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
			}
			Expect(source).To(BecomePreflightDocument(expected))
		})

		It("section level 1 with a paragraph separated by non-empty line", func() {
			source := "== a title\n    \nand a paragraph"
			section1Title := types.InlineElements{
				types.StringElement{Content: "a title"},
			}
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "a_title",
							types.AttrCustomID: false,
						},
						Level:    1,
						Title:    section1Title,
						Elements: []interface{}{},
					},
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
			}
			Expect(source).To(BecomePreflightDocument(expected))
		})

		It("section levels 1, 2, 3, 2", func() {
			source := `= a header

== Section A
a paragraph

=== Section A.a
a paragraph

== Section B
a paragraph`
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "a_header",
							types.AttrCustomID: false,
						},
						Level: 0,
						Title: types.InlineElements{
							types.StringElement{Content: "a header"},
						},
						Elements: []interface{}{},
					},
					types.BlankLine{},
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "section_a",
							types.AttrCustomID: false,
						},
						Level: 1,
						Title: types.InlineElements{
							types.StringElement{Content: "Section A"},
						},
						Elements: []interface{}{},
					},
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
						Attributes: types.ElementAttributes{
							types.AttrID:       "section_a_a",
							types.AttrCustomID: false,
						},
						Level: 2,
						Title: types.InlineElements{
							types.StringElement{Content: "Section A.a"},
						},
						Elements: []interface{}{},
					},
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
						Attributes: types.ElementAttributes{
							types.AttrID:       "section_b",
							types.AttrCustomID: false,
						},
						Level: 1,
						Title: types.InlineElements{
							types.StringElement{Content: "Section B"},
						},
						Elements: []interface{}{},
					},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "a paragraph"},
							},
						},
					},
				},
			}
			Expect(source).To(BecomePreflightDocument(expected))
		})

		It("single section with custom IDs", func() {
			source := `[[custom_header]]
== a header`
			sectionTitle := types.InlineElements{
				types.StringElement{Content: "a header"},
			}
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "custom_header",
							types.AttrCustomID: true,
						},
						Level:    1,
						Title:    sectionTitle,
						Elements: []interface{}{},
					},
				},
			}
			Expect(source).To(BecomePreflightDocument(expected))
		})

		It("multiple sections with custom IDs", func() {
			source := `[[custom_header]]
= a header

== Section F [[ignored]] [[foo]]

[[bar]]
== Section B
a paragraph`
			doctitle := types.InlineElements{
				types.StringElement{Content: "a header"},
			}
			fooTitle := types.InlineElements{
				types.StringElement{Content: "Section F "},
			}
			barTitle := types.InlineElements{
				types.StringElement{Content: "Section B"},
			}
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "custom_header",
							types.AttrCustomID: true,
						},
						Level:    0,
						Title:    doctitle,
						Elements: []interface{}{},
					},
					types.BlankLine{},
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "foo",
							types.AttrCustomID: true,
						},
						Level:    1,
						Title:    fooTitle,
						Elements: []interface{}{},
					},
					types.BlankLine{},
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "bar",
							types.AttrCustomID: true,
						},
						Level:    1,
						Title:    barTitle,
						Elements: []interface{}{},
					},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "a paragraph"},
							},
						},
					},
				},
			}
			Expect(source).To(BecomePreflightDocument(expected))
		})

		It("sections with same title", func() {
			source := `== section 1

== section 1`
			section1aTitle := types.InlineElements{
				types.StringElement{Content: "section 1"},
			}
			section1bTitle := types.InlineElements{
				types.StringElement{Content: "section 1"},
			}
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "section_1",
							types.AttrCustomID: false,
						},
						Level:    1,
						Title:    section1aTitle,
						Elements: []interface{}{},
					},
					types.BlankLine{},
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "section_1", // duplicate ID will be processed afterwards
							types.AttrCustomID: false,
						},
						Level:    1,
						Title:    section1bTitle,
						Elements: []interface{}{},
					},
				},
			}
			Expect(source).To(BecomePreflightDocument(expected))
		})
	})

	Context("invalid sections", func() {
		It("header invalid - missing space", func() {
			source := "=a header"
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "=a header"},
							},
						},
					},
				}}
			Expect(source).To(BecomePreflightDocument(expected))
		})

		It("header invalid - header space", func() {
			source := " = a header with a prefix space"
			expected := types.PreflightDocument{
				Blocks: []interface{}{
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
			Expect(source).To(BecomePreflightDocument(expected))
		})

		It("header with invalid section1", func() {
			source := `= a header

   == section with prefix space`
			title := types.InlineElements{
				types.StringElement{Content: "a header"},
			}
			expected := types.PreflightDocument{
				Blocks: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "a_header",
							types.AttrCustomID: false,
						},
						Level:    0,
						Title:    title,
						Elements: []interface{}{},
					},
					types.BlankLine{},
					types.LiteralBlock{
						Attributes: types.ElementAttributes{
							types.AttrKind:             types.Literal,
							types.AttrLiteralBlockType: types.LiteralBlockWithSpacesOnFirstLine,
						},
						Lines: []string{
							"   == section with prefix space",
						},
					},
				},
			}
			Expect(source).To(BecomePreflightDocument(expected))
		})

	})

	Context("unsupported section syntax", func() {

		It("should not fail with underlined title", func() {
			source := `Document Title
==============
Doc Writer <thedoc@asciidoctor.org>`
			expected := types.PreflightDocument{
				Blocks: []interface{}{
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
			Expect(source).To(BecomePreflightDocument(expected))
		})
	})
})

var _ = Describe("sections - document", func() {

	Context("valid sections", func() {

		It("header only", func() {
			source := "= a header"
			doctitle := types.InlineElements{
				types.StringElement{Content: "a header"},
			}
			expected := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"a_header": doctitle,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "a_header",
							types.AttrCustomID: false,
						},
						Level:    0,
						Title:    doctitle,
						Elements: []interface{}{},
					},
				},
			}
			Expect(source).To(EqualDocument(expected))
		})

		It("header with many spaces around content", func() {
			source := "= a header   "
			doctitle := types.InlineElements{
				types.StringElement{Content: "a header   "},
			}
			expected := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"a_header": doctitle,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "a_header",
							types.AttrCustomID: false,
						},
						Level:    0,
						Title:    doctitle,
						Elements: []interface{}{},
					},
				},
			}
			Expect(source).To(EqualDocument(expected))
		})

		It("header and paragraph", func() {
			source := `= a header

and a paragraph`

			doctitle := types.InlineElements{
				types.StringElement{Content: "a header"},
			}
			expected := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"a_header": doctitle,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "a_header",
							types.AttrCustomID: false,
						},
						Level: 0,
						Title: doctitle,
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
			Expect(source).To(EqualDocument(expected))
		})

		It("two sections with level 0", func() {
			source := `= a first header

= a second header`
			doctitle := types.InlineElements{
				types.StringElement{Content: "a first header"},
			}
			otherDoctitle := types.InlineElements{
				types.StringElement{Content: "a second header"},
			}

			expected := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"a_first_header":  doctitle,
					"a_second_header": otherDoctitle,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "a_first_header",
							types.AttrCustomID: false,
						},
						Level:    0,
						Title:    doctitle,
						Elements: []interface{}{},
					},
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "a_second_header",
							types.AttrCustomID: false,
						},
						Level:    0,
						Title:    otherDoctitle,
						Elements: []interface{}{},
					},
				},
			}
			Expect(source).To(EqualDocument(expected))
		})

		It("section level 1 alone", func() {
			source := `== section 1`
			section1Title := types.InlineElements{
				types.StringElement{Content: "section 1"},
			}
			expected := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"section_1": section1Title,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "section_1",
							types.AttrCustomID: false,
						},
						Level:    1,
						Title:    section1Title,
						Elements: []interface{}{},
					},
				},
			}
			Expect(source).To(EqualDocument(expected))
		})

		It("section level 1 with quoted text", func() {
			source := `==  *2 spaces and bold content*`
			sectionTitle := types.InlineElements{
				types.QuotedText{
					Kind: types.Bold,
					Elements: types.InlineElements{
						types.StringElement{Content: "2 spaces and bold content"},
					},
				},
			}
			expected := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"2_spaces_and_bold_content": sectionTitle,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "2_spaces_and_bold_content",
							types.AttrCustomID: false,
						},
						Level:    1,
						Title:    sectionTitle,
						Elements: []interface{}{},
					},
				},
			}
			Expect(source).To(EqualDocument(expected))
		})

		It("section level 0 with nested section level 1", func() {
			source := `= a header

== section 1`
			doctitle := types.InlineElements{
				types.StringElement{Content: "a header"},
			}
			section1Title := types.InlineElements{
				types.StringElement{Content: "section 1"},
			}
			expected := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"a_header":  doctitle,
					"section_1": section1Title,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "a_header",
							types.AttrCustomID: false,
						},
						Level: 0,
						Title: doctitle,
						Elements: []interface{}{
							types.Section{
								Attributes: types.ElementAttributes{
									types.AttrID:       "section_1",
									types.AttrCustomID: false,
								},
								Level:    1,
								Title:    section1Title,
								Elements: []interface{}{},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocument(expected))
		})

		It("section level 0 with nested section level 2", func() {
			source := `= a header

=== section 2`
			doctitle := types.InlineElements{
				types.StringElement{Content: "a header"},
			}
			section2Title := types.InlineElements{
				types.StringElement{Content: "section 2"},
			}
			expected := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"a_header":  doctitle,
					"section_2": section2Title,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "a_header",
							types.AttrCustomID: false,
						},
						Level: 0,
						Title: doctitle,
						Elements: []interface{}{
							types.Section{
								Attributes: types.ElementAttributes{
									types.AttrID:       "section_2",
									types.AttrCustomID: false,
								},
								Level:    2,
								Title:    section2Title,
								Elements: []interface{}{},
							},
						},
					},
				},
			}
			Expect(source).To(EqualDocument(expected))
		})

		It("section level 1 with immediate paragraph", func() {
			source := `== a title
and a paragraph`
			section1Title := types.InlineElements{
				types.StringElement{Content: "a title"},
			}
			expected := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"a_title": section1Title,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "a_title",
							types.AttrCustomID: false,
						},
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
			Expect(source).To(EqualDocument(expected))
		})

		It("section level 1 with a paragraph separated by empty line", func() {
			source := `== a title
			
and a paragraph`
			section1Title := types.InlineElements{
				types.StringElement{Content: "a title"},
			}
			expected := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"a_title": section1Title,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "a_title",
							types.AttrCustomID: false,
						},
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
			Expect(source).To(EqualDocument(expected))
		})

		It("section level 1 with a paragraph separated by non-empty line", func() {
			source := "== a title\n    \nand a paragraph"
			section1Title := types.InlineElements{
				types.StringElement{Content: "a title"},
			}
			expected := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"a_title": section1Title,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "a_title",
							types.AttrCustomID: false,
						},
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
			Expect(source).To(EqualDocument(expected))
		})

		It("section levels 1, 2, 3, 2", func() {
			source := `= a header

== Section A
a paragraph

=== Section A.a
a paragraph

== Section B
a paragraph`
			doctitle := types.InlineElements{
				types.StringElement{Content: "a header"},
			}
			sectionATitle := types.InlineElements{
				types.StringElement{Content: "Section A"},
			}
			sectionAaTitle := types.InlineElements{
				types.StringElement{Content: "Section A.a"},
			}
			sectionBTitle := types.InlineElements{
				types.StringElement{Content: "Section B"},
			}
			expected := types.Document{
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
						Attributes: types.ElementAttributes{
							types.AttrID:       "a_header",
							types.AttrCustomID: false,
						},
						Level: 0,
						Title: doctitle,
						Elements: []interface{}{
							types.Section{
								Attributes: types.ElementAttributes{
									types.AttrID:       "section_a",
									types.AttrCustomID: false,
								},
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
									types.Section{
										Attributes: types.ElementAttributes{
											types.AttrID:       "section_a_a",
											types.AttrCustomID: false,
										},
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
										},
									},
								},
							},
							types.Section{
								Attributes: types.ElementAttributes{
									types.AttrID:       "section_b",
									types.AttrCustomID: false,
								},
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
					},
				},
			}
			Expect(source).To(EqualDocument(expected))
		})

		It("single section with custom IDs", func() {
			source := `[[custom_header]]
== a header`
			sectionTitle := types.InlineElements{
				types.StringElement{Content: "a header"},
			}
			expected := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"custom_header": sectionTitle,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "custom_header",
							types.AttrCustomID: true,
						},
						Level:    1,
						Title:    sectionTitle,
						Elements: []interface{}{},
					},
				},
			}
			Expect(source).To(EqualDocument(expected))
		})

		It("multiple sections with custom IDs", func() {
			source := `[[custom_header]]
= a header

== Section F [[ignored]] [[foo]]

[[bar]]
== Section B
a paragraph`
			doctitle := types.InlineElements{
				types.StringElement{Content: "a header"},
			}
			fooTitle := types.InlineElements{
				types.StringElement{Content: "Section F "},
			}
			barTitle := types.InlineElements{
				types.StringElement{Content: "Section B"},
			}
			expected := types.Document{
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
						Attributes: types.ElementAttributes{
							types.AttrID:       "custom_header",
							types.AttrCustomID: true,
						},
						Level: 0,
						Title: doctitle,
						Elements: []interface{}{
							types.Section{
								Attributes: types.ElementAttributes{
									types.AttrID:       "foo",
									types.AttrCustomID: true,
								},
								Level:    1,
								Title:    fooTitle,
								Elements: []interface{}{},
							},
							types.Section{
								Attributes: types.ElementAttributes{
									types.AttrID:       "bar",
									types.AttrCustomID: true,
								},
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
					},
				},
			}
			Expect(source).To(EqualDocument(expected))
		})

		It("sections with same title", func() {
			source := `== section 1

== section 1`
			section1aTitle := types.InlineElements{
				types.StringElement{Content: "section 1"},
			}
			section1bTitle := types.InlineElements{
				types.StringElement{Content: "section 1"},
			}

			expected := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"section_1":   section1aTitle,
					"section_1_2": section1bTitle,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "section_1",
							types.AttrCustomID: false,
						},
						Level:    1,
						Title:    section1aTitle,
						Elements: []interface{}{},
					},
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "section_1_2",
							types.AttrCustomID: false,
						},
						Level:    1,
						Title:    section1bTitle,
						Elements: []interface{}{},
					},
				},
			}
			Expect(source).To(EqualDocument(expected))
		})
	})

	Context("invalid sections", func() {

		It("header invalid - too many spaces", func() {
			source := "======= a header"
			expected := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "======= a header"},
							},
						},
					},
				}}
			Expect(source).To(EqualDocument(expected))
		})

		It("header invalid - missing space", func() {
			source := "=a header"
			expected := types.Document{
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
			Expect(source).To(EqualDocument(expected))
		})

		It("header invalid - header space", func() {
			source := " = a header with a prefix space"
			expected := types.Document{
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
			Expect(source).To(EqualDocument(expected))
		})

		It("header with invalid section1", func() {
			source := `= a header

 == section with prefix space`
			title := types.InlineElements{
				types.StringElement{Content: "a header"},
			}
			expected := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"a_header": title,
				}, Footnotes: types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID:       "a_header",
							types.AttrCustomID: false,
						},
						Level: 0,
						Title: title,
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
			Expect(source).To(EqualDocument(expected))
		})

	})

	Context("unsupported section syntax", func() {

		It("should not fail with underlined title", func() {
			source := `Document Title
==============
Doc Writer <thedoc@asciidoctor.org>`
			expected := types.Document{
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
			Expect(source).To(EqualDocument(expected))
		})
	})
})
