package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("comments", func() {

	Context("in final documents", func() {

		Context("single line comments", func() {

			It("alone", func() {
				source := `// A single-line comment.`
				expected := &types.Document{}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with prefixing spaces alone", func() {
				source := `  // A single-line comment.`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Attributes: types.Attributes{
								types.AttrStyle:            types.Literal,
								types.AttrLiteralBlockType: types.LiteralBlockWithSpacesOnFirstLine,
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "  // A single-line comment.",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with prefixing tabs alone", func() {
				source := "\t\t// A single-line comment."
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Attributes: types.Attributes{
								types.AttrStyle:            types.Literal,
								types.AttrLiteralBlockType: types.LiteralBlockWithSpacesOnFirstLine,
							},
							Elements: []interface{}{
								&types.StringElement{
									Content: "\t\t// A single-line comment.",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("at end of line", func() {
				source := `foo // A single-line comment.`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "foo // A single-line comment.",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("within a paragraph", func() {
				source := `a first line
// A single-line comment.
another line`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "a first line",
								},
								&types.StringElement{
									Content: "\nanother line",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			Context("invalid", func() {

				It("within a paragraph with tab", func() {
					source := `a first line
	// not a comment.
another line`
					expected := &types.Document{
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "a first line\n\t// not a comment.\nanother line",
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})
			})
		})

		Context("comment blocks", func() {

			It("alone", func() {
				source := `//// 
a *comment* block
with multiple lines
////`
				expected := &types.Document{}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with paragraphs around", func() {
				source := `a first paragraph
			
//// 
a *comment* block
with multiple lines
////
a second paragraph`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "a first paragraph",
								},
							},
						},
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "a second paragraph",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
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
				&types.StringElement{
					Content: "section 1",
				},
			}
			expected := &types.Document{
				ElementReferences: types.ElementReferences{
					"_section_1": section1Title,
				},
				Elements: []interface{}{
					&types.Section{
						Attributes: types.Attributes{
							types.AttrID: "_section_1",
						},
						Level: 1,
						Title: section1Title,
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "a first paragraph",
									},
								},
							},
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "a second paragraph",
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
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
			headerTitle := []interface{}{
				&types.StringElement{
					Content: "section 0",
				},
			}
			section1Title := []interface{}{
				&types.StringElement{
					Content: "section 1",
				},
			}
			expected := &types.Document{
				ElementReferences: types.ElementReferences{
					"_section_1": section1Title,
				},
				Elements: []interface{}{
					&types.DocumentHeader{
						Title: headerTitle,
					},
					&types.Section{
						Attributes: types.Attributes{
							types.AttrID: "_section_1",
						},
						Level: 1,
						Title: section1Title,
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "a first paragraph",
									},
								},
							},
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "a second paragraph",
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
