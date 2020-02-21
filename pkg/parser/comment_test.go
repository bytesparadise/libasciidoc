package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("comments", func() {

	Context("draft document", func() {

		Context("single line comments", func() {

			It("single line comment alone", func() {
				source := `// A single-line comment.`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.SingleLineComment{
							Content: " A single-line comment.",
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})

			It("single line comment with prefixing spaces alone", func() {
				source := `  // A single-line comment.`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.SingleLineComment{
							Content: " A single-line comment.",
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})

			It("single line comment with prefixing tabs alone", func() {
				source := "\t\t// A single-line comment."
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.SingleLineComment{
							Content: " A single-line comment.",
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})

			It("single line comment at end of line", func() {
				source := `foo // A single-line comment.`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "foo // A single-line comment."},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})

			It("single line comment within a paragraph", func() {
				source := `a first line
// A single-line comment.
another line`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a first line"},
								},
								{
									types.SingleLineComment{Content: " A single-line comment."},
								},
								{
									types.StringElement{Content: "another line"},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})

			It("single line comment within a paragraph with tab", func() {
				source := `a first line
	// A single-line comment.
another line`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a first line"},
								},
								{
									types.SingleLineComment{Content: " A single-line comment."},
								},
								{
									types.StringElement{Content: "another line"},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})
		})

		Context("comment blocks", func() {

			It("comment block alone", func() {
				source := `//// 
a *comment* block
with multiple lines
////`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.DelimitedBlock{
							Attributes: types.ElementAttributes{},
							Kind:       types.Comment,
							Elements: []interface{}{
								types.StringElement{
									Content: "a *comment* block",
								},
								types.StringElement{
									Content: "with multiple lines",
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})

			It("comment block with paragraphs around", func() {
				source := `a first paragraph
//// 
a *comment* block
with multiple lines
////
a second paragraph`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a first paragraph"},
								},
							},
						},
						types.DelimitedBlock{
							Attributes: types.ElementAttributes{},
							Kind:       types.Comment,
							Elements: []interface{}{
								types.StringElement{
									Content: "a *comment* block",
								},
								types.StringElement{
									Content: "with multiple lines",
								},
							},
						},
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a second paragraph"},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(Equal(expected))
			})
		})
	})

	Context("final document", func() {

		Context("single line comments", func() {

			It("single line comment alone", func() {
				source := `// A single-line comment.`
				expected := types.Document{
					Attributes:         types.DocumentAttributes{},
					ElementReferences:  types.ElementReferences{},
					Footnotes:          types.Footnotes{},
					FootnoteReferences: types.FootnoteReferences{},
					Elements:           []interface{}{},
				}
				Expect(ParseDocument(source)).To(Equal(expected))
			})

			It("single line comment with prefixing spaces alone", func() {
				source := `  // A single-line comment.`
				expected := types.Document{
					Attributes:         types.DocumentAttributes{},
					ElementReferences:  types.ElementReferences{},
					Footnotes:          types.Footnotes{},
					FootnoteReferences: types.FootnoteReferences{},
					Elements:           []interface{}{},
				}
				Expect(ParseDocument(source)).To(Equal(expected))
			})

			It("single line comment with prefixing tabs alone", func() {
				source := "\t\t// A single-line comment."
				expected := types.Document{
					Attributes:         types.DocumentAttributes{},
					ElementReferences:  types.ElementReferences{},
					Footnotes:          types.Footnotes{},
					FootnoteReferences: types.FootnoteReferences{},
					Elements:           []interface{}{},
				}
				Expect(ParseDocument(source)).To(Equal(expected))
			})

			It("single line comment at end of line", func() {
				source := `foo // A single-line comment.`
				expected := types.Document{
					Attributes:         types.DocumentAttributes{},
					ElementReferences:  types.ElementReferences{},
					Footnotes:          types.Footnotes{},
					FootnoteReferences: types.FootnoteReferences{},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "foo // A single-line comment."},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(Equal(expected))
			})

			It("single line comment within a paragraph", func() {
				source := `a first line
// A single-line comment.
another line`
				expected := types.Document{
					Attributes:         types.DocumentAttributes{},
					ElementReferences:  types.ElementReferences{},
					Footnotes:          types.Footnotes{},
					FootnoteReferences: types.FootnoteReferences{},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a first line"},
								},
								{
									types.StringElement{Content: "another line"},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(Equal(expected))
			})

			It("single line comment within a paragraph with tab", func() {
				source := `a first line
	// A single-line comment.
another line`
				expected := types.Document{
					Attributes:         types.DocumentAttributes{},
					ElementReferences:  types.ElementReferences{},
					Footnotes:          types.Footnotes{},
					FootnoteReferences: types.FootnoteReferences{},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a first line"},
								},
								{
									types.StringElement{Content: "another line"},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(Equal(expected))
			})
		})

		Context("comment blocks", func() {

			It("comment block alone", func() {
				source := `//// 
a *comment* block
with multiple lines
////`
				expected := types.Document{
					Attributes:         types.DocumentAttributes{},
					ElementReferences:  types.ElementReferences{},
					Footnotes:          types.Footnotes{},
					FootnoteReferences: types.FootnoteReferences{},
					Elements:           []interface{}{},
				}
				Expect(ParseDocument(source)).To(Equal(expected))
			})

			It("comment block with paragraphs around", func() {
				source := `a first paragraph
//// 
a *comment* block
with multiple lines
////
a second paragraph`
				expected := types.Document{
					Attributes:         types.DocumentAttributes{},
					ElementReferences:  types.ElementReferences{},
					Footnotes:          types.Footnotes{},
					FootnoteReferences: types.FootnoteReferences{},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a first paragraph"},
								},
							},
						},
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a second paragraph"},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(Equal(expected))
			})
		})

		It("comment in section", func() {
			source := `== section 1

a first paragraph
//// 
a *comment* block
with multiple lines
////
a second paragraph`
			section1Title := []interface{}{
				types.StringElement{
					Content: "section 1",
				},
			}
			expected := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"_section_1": section1Title,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID: "_section_1",
						},
						Level: 1,
						Title: section1Title,
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "a first paragraph"},
									},
								},
							},
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
									{
										types.StringElement{Content: "a second paragraph"},
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(Equal(expected))
		})

		It("comment in preamble", func() {
			source := `= section 0

//// 
a *comment* block
with multiple lines
////

== section 1

a first paragraph

a second paragraph`
			section0Title := []interface{}{
				types.StringElement{
					Content: "section 0",
				},
			}
			section1Title := []interface{}{
				types.StringElement{
					Content: "section 1",
				},
			}
			expected := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"_section_0": section0Title,
					"_section_1": section1Title,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Attributes: types.ElementAttributes{
							types.AttrID: "_section_0",
						},
						Level: 0,
						Title: section0Title,
						Elements: []interface{}{
							types.Section{
								Attributes: types.ElementAttributes{
									types.AttrID: "_section_1",
								},
								Level: 1,
								Title: section1Title,
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: [][]interface{}{
											{
												types.StringElement{Content: "a first paragraph"},
											},
										},
									},
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: [][]interface{}{
											{
												types.StringElement{Content: "a second paragraph"},
											},
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(Equal(expected))
		})
	})

})
