package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("tables of contents", func() {

	Context("in document fragments", func() {

		It("with default level", func() {
			/*
				= A title
				:toc:
				== Section A
				=== Section A.a
				=== Section A.b
				==== Section that shall not be in ToC
				== Section B
				=== Section B.a
				== Section C
			*/
			c := make(chan types.DocumentFragment, 10)
			c <- types.DocumentFragment{
				Elements: []interface{}{
					&types.DocumentHeader{
						Title: []interface{}{
							&types.StringElement{
								Content: "A title",
							},
						},
						Elements: []interface{}{
							&types.AttributeDeclaration{
								Name: types.AttrTableOfContents,
							},
						},
					},
				},
			}
			c <- types.DocumentFragment{
				Elements: []interface{}{
					&types.Section{
						Level: 1,
						Attributes: types.Attributes{
							types.AttrID: "_section_a",
						},
						Title: []interface{}{
							&types.StringElement{
								Content: "Section A",
							},
						},
					},
				},
			}
			c <- types.DocumentFragment{
				Elements: []interface{}{
					&types.Section{
						Level: 2,
						Attributes: types.Attributes{
							types.AttrID: "_section_a_a",
						},
						Title: []interface{}{
							&types.StringElement{
								Content: "Section A.a",
							},
						},
					},
				},
			}
			c <- types.DocumentFragment{
				Elements: []interface{}{
					&types.Section{
						Level: 2,
						Attributes: types.Attributes{
							types.AttrID: "_section_a_b",
						},
						Title: []interface{}{
							&types.StringElement{
								Content: "Section A.b",
							},
						},
					},
				},
			}
			c <- types.DocumentFragment{
				Elements: []interface{}{
					&types.Section{
						Level: 3,
						Attributes: types.Attributes{
							types.AttrID: "_section_that_shall_not_be_in_ToC",
						},
						Title: []interface{}{
							&types.StringElement{
								Content: "Section that shall not be in ToC",
							},
						},
					},
				},
			}
			c <- types.DocumentFragment{
				Elements: []interface{}{
					&types.Section{
						Level: 1,
						Attributes: types.Attributes{
							types.AttrID: "_section_b",
						},
						Title: []interface{}{
							&types.StringElement{
								Content: "Section B",
							},
						},
					},
				},
			}
			c <- types.DocumentFragment{
				Elements: []interface{}{
					&types.Section{
						Level: 2,
						Attributes: types.Attributes{
							types.AttrID: "_section_b_a",
						},
						Title: []interface{}{
							&types.StringElement{
								Content: "Section B.a",
							},
						},
					},
				},
			}
			c <- types.DocumentFragment{
				Elements: []interface{}{
					&types.Section{
						Level: 1,
						Attributes: types.Attributes{
							types.AttrID: "_section_c",
						},
						Title: []interface{}{
							&types.StringElement{
								Content: "Section C",
							},
						},
					},
				},
			}
			close(c)

			expectedToC := &types.TableOfContents{
				MaxDepth: 2,
				Sections: []*types.ToCSection{
					{
						ID:    "_section_a",
						Level: 1,
						Children: []*types.ToCSection{
							{
								ID:    "_section_a_a",
								Level: 2,
							},
							{
								ID:    "_section_a_b",
								Level: 2,
							},
						},
					},
					{
						ID:    "_section_b",
						Level: 1,
						Children: []*types.ToCSection{
							{
								ID:    "_section_b_a",
								Level: 2,
							},
						},
					},
					{
						ID:    "_section_c",
						Level: 1,
					},
				},
			}
			ctx := parser.NewParseContext(configuration.NewConfiguration())
			doc, err := parser.Aggregate(ctx, c)
			Expect(err).ToNot(HaveOccurred())
			Expect(doc.TableOfContents).To(MatchTableOfContents(expectedToC))
		})

		It("with custom level", func() {
			/*
				= A title
				:toc:
				:toclevels: 3

				== Section A
				=== Section A.a
				=== Section A.b
				==== Section that shall be in ToC
				== Section B
				=== Section B.a
				== Section C
			*/
			c := make(chan types.DocumentFragment, 10)
			c <- types.DocumentFragment{
				Elements: []interface{}{
					&types.DocumentHeader{
						Title: []interface{}{
							&types.StringElement{
								Content: "A title",
							},
						},
						Elements: []interface{}{
							&types.AttributeDeclaration{
								Name: types.AttrTableOfContents,
							},
							&types.AttributeDeclaration{
								Name:  types.AttrTableOfContentsLevels,
								Value: "3",
							},
						},
					},
				},
			}
			c <- types.DocumentFragment{
				Elements: []interface{}{
					&types.Section{
						Level: 1,
						Attributes: types.Attributes{
							types.AttrID: "_section_a",
						},
						Title: []interface{}{
							&types.StringElement{
								Content: "Section A",
							},
						},
					},
				},
			}
			c <- types.DocumentFragment{
				Elements: []interface{}{
					&types.Section{
						Level: 2,
						Attributes: types.Attributes{
							types.AttrID: "_section_a_a",
						},
						Title: []interface{}{
							&types.StringElement{
								Content: "Section A.a",
							},
						},
					},
				},
			}
			c <- types.DocumentFragment{
				Elements: []interface{}{
					&types.Section{
						Level: 2,
						Attributes: types.Attributes{
							types.AttrID: "_section_a_b",
						},
						Title: []interface{}{
							&types.StringElement{
								Content: "Section A.b",
							},
						},
					},
				},
			}
			c <- types.DocumentFragment{
				Elements: []interface{}{
					&types.Section{
						Level: 3,
						Attributes: types.Attributes{
							types.AttrID: "_section_that_shall_be_in_ToC",
						},
						Title: []interface{}{
							&types.StringElement{
								Content: "Section that shall be in ToC",
							},
						},
					},
				},
			}
			c <- types.DocumentFragment{
				Elements: []interface{}{
					&types.Section{
						Level: 1,
						Attributes: types.Attributes{
							types.AttrID: "_section_b",
						},
						Title: []interface{}{
							&types.StringElement{
								Content: "Section B",
							},
						},
					},
				},
			}
			c <- types.DocumentFragment{
				Elements: []interface{}{
					&types.Section{
						Level: 2,
						Attributes: types.Attributes{
							types.AttrID: "_section_b_a",
						},
						Title: []interface{}{
							&types.StringElement{
								Content: "Section B.a",
							},
						},
					},
				},
			}
			c <- types.DocumentFragment{
				Elements: []interface{}{
					&types.Section{
						Level: 1,
						Attributes: types.Attributes{
							types.AttrID: "_section_c",
						},
						Title: []interface{}{
							&types.StringElement{
								Content: "Section C",
							},
						},
					},
				},
			}
			close(c)

			expectedToC := &types.TableOfContents{
				MaxDepth: 3,
				Sections: []*types.ToCSection{
					{
						ID:    "_section_a",
						Level: 1,
						Children: []*types.ToCSection{
							{
								ID:    "_section_a_a",
								Level: 2,
							},
							{
								ID:    "_section_a_b",
								Level: 2,
								Children: []*types.ToCSection{
									{
										ID:    "_section_that_shall_be_in_ToC",
										Level: 3,
									},
								},
							},
						},
					},
					{
						ID:    "_section_b",
						Level: 1,
						Children: []*types.ToCSection{
							{
								ID:    "_section_b_a",
								Level: 2,
							},
						},
					},
					{
						ID:    "_section_c",
						Level: 1,
					},
				},
			}
			ctx := parser.NewParseContext(configuration.NewConfiguration())
			doc, err := parser.Aggregate(ctx, c)
			Expect(err).ToNot(HaveOccurred())
			Expect(doc.TableOfContents).To(MatchTableOfContents(expectedToC))
		})

	})

	Context("in final documents", func() {

		// same titles for all tests in this context
		section1Title := []interface{}{
			&types.StringElement{
				Content: "Section ",
			},
			&types.QuotedText{
				Kind: types.SingleQuoteBold,
				Elements: []interface{}{
					&types.StringElement{
						Content: "1",
					},
				},
			},
		}
		section2Title := []interface{}{
			&types.StringElement{
				Content: "Section 2",
			},
		}

		It("without comments in document header", func() {
			source := `= Title
:toc: preamble

a preamble 

== Section *1*

== Section 2`
			expected := &types.Document{
				Elements: []interface{}{
					&types.DocumentHeader{
						Title: []interface{}{
							&types.StringElement{
								Content: "Title",
							},
						},
						Elements: []interface{}{
							&types.AttributeDeclaration{
								Name:  types.AttrTableOfContents,
								Value: "preamble",
							},
						},
					},
					&types.Preamble{
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "a preamble",
									},
								},
							},
						},
					},
					&types.Section{
						Level: 1,
						Attributes: types.Attributes{
							types.AttrID: "_Section_1",
						},
						Title: section1Title,
					},
					&types.Section{
						Level: 1,
						Attributes: types.Attributes{
							types.AttrID: "_Section_2",
						},
						Title: section2Title,
					},
				},
				ElementReferences: types.ElementReferences{
					"_Section_1": section1Title,
					"_Section_2": section2Title,
				},
				TableOfContents: &types.TableOfContents{
					MaxDepth: 2,
					Sections: []*types.ToCSection{
						{
							ID:    "_Section_1",
							Level: 1,
						},
						{
							ID:    "_Section_2",
							Level: 1,
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("with single line comments in document header", func() {
			source := `= Title
// a comment
// another comment
:toc: preamble
// and once more

a preamble 

== Section *1*

== Section 2`
			expected := &types.Document{
				Elements: []interface{}{
					&types.DocumentHeader{
						Title: []interface{}{
							&types.StringElement{
								Content: "Title",
							},
						},
						Elements: []interface{}{
							// single comments are filtered out
							&types.AttributeDeclaration{
								Name:  types.AttrTableOfContents,
								Value: "preamble",
							},
							// single comment is filtered out
						},
					},
					&types.Preamble{
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "a preamble",
									},
								},
							},
						},
					},
					&types.Section{
						Level: 1,
						Attributes: types.Attributes{
							types.AttrID: "_Section_1",
						},
						Title: section1Title,
					},
					&types.Section{
						Level: 1,
						Attributes: types.Attributes{
							types.AttrID: "_Section_2",
						},
						Title: section2Title,
					},
				},
				ElementReferences: types.ElementReferences{
					"_Section_1": section1Title,
					"_Section_2": section2Title,
				},
				TableOfContents: &types.TableOfContents{
					MaxDepth: 2,
					Sections: []*types.ToCSection{
						{
							ID:    "_Section_1",
							Level: 1,
						},
						{
							ID:    "_Section_2",
							Level: 1,
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("with comment blocks in document header", func() {
			source := `= Title
////
a 
comment 
block
////
:toc: preamble
////
another 
comment 
block
////

a preamble 

== Section *1*

== Section 2`
			expected := &types.Document{
				Elements: []interface{}{
					&types.DocumentHeader{
						Title: []interface{}{
							&types.StringElement{
								Content: "Title",
							},
						},
						Elements: []interface{}{
							// comment block is filtered out
							&types.AttributeDeclaration{
								Name:  types.AttrTableOfContents,
								Value: "preamble",
							},
							// comment block is filtered out
						},
					},
					&types.Preamble{
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "a preamble",
									},
								},
							},
						},
					},
					&types.Section{
						Level: 1,
						Attributes: types.Attributes{
							types.AttrID: "_Section_1",
						},
						Title: section1Title,
					},
					&types.Section{
						Level: 1,
						Attributes: types.Attributes{
							types.AttrID: "_Section_2",
						},
						Title: section2Title,
					},
				},
				ElementReferences: types.ElementReferences{
					"_Section_1": section1Title,
					"_Section_2": section2Title,
				},
				TableOfContents: &types.TableOfContents{
					MaxDepth: 2,
					Sections: []*types.ToCSection{
						{
							ID:    "_Section_1",
							Level: 1,
						},
						{
							ID:    "_Section_2",
							Level: 1,
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should render with custom title without passthrough macro", func() {
			source := `= Title
:toc:
:toc-title: <h3>Table of Contents</h3>

== Section *1*

== Section 2
`
			expected := &types.Document{
				Elements: []interface{}{
					&types.DocumentHeader{
						Title: []interface{}{
							&types.StringElement{
								Content: "Title",
							},
						},
						Elements: []interface{}{
							&types.AttributeDeclaration{
								Name: types.AttrTableOfContents,
							},
							&types.AttributeDeclaration{
								Name: types.AttrTableOfContentsTitle,
								Value: []interface{}{
									&types.SpecialCharacter{
										Name: "<",
									},
									&types.StringElement{
										Content: "h3",
									},
									&types.SpecialCharacter{
										Name: ">",
									},
									&types.StringElement{
										Content: "Table of Contents",
									},
									&types.SpecialCharacter{
										Name: "<",
									},
									&types.StringElement{
										Content: "/h3",
									},
									&types.SpecialCharacter{
										Name: ">",
									},
								},
							},
						},
					},
					&types.Section{
						Level: 1,
						Attributes: types.Attributes{
							types.AttrID: "_Section_1",
						},
						Title: section1Title,
					},
					&types.Section{
						Level: 1,
						Attributes: types.Attributes{
							types.AttrID: "_Section_2",
						},
						Title: section2Title,
					},
				},
				ElementReferences: types.ElementReferences{
					"_Section_1": section1Title,
					"_Section_2": section2Title,
				},
				TableOfContents: &types.TableOfContents{
					MaxDepth: 2,
					Sections: []*types.ToCSection{
						{
							ID:    "_Section_1",
							Level: 1,
						},
						{
							ID:    "_Section_2",
							Level: 1,
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should render with custom title with passthrough macro", func() {
			source := `= Title
:toc:
:toc-title: pass:[<h3>Table of Contents</h3>]

== Section *1*

== Section 2
`
			expected := &types.Document{
				Elements: []interface{}{
					&types.DocumentHeader{
						Title: []interface{}{
							&types.StringElement{
								Content: "Title",
							},
						},
						Elements: []interface{}{
							&types.AttributeDeclaration{
								Name: types.AttrTableOfContents,
							},
							&types.AttributeDeclaration{
								Name: types.AttrTableOfContentsTitle,
								Value: []interface{}{
									&types.InlinePassthrough{
										Kind: types.PassthroughMacro,
										Elements: []interface{}{
											&types.StringElement{
												Content: "<h3>Table of Contents</h3>",
											},
										},
									},
								},
							},
						},
					},
					&types.Section{
						Level: 1,
						Attributes: types.Attributes{
							types.AttrID: "_Section_1",
						},
						Title: section1Title,
					},
					&types.Section{
						Level: 1,
						Attributes: types.Attributes{
							types.AttrID: "_Section_2",
						},
						Title: section2Title,
					},
				},
				ElementReferences: types.ElementReferences{
					"_Section_1": section1Title,
					"_Section_2": section2Title,
				},
				TableOfContents: &types.TableOfContents{
					MaxDepth: 2,
					Sections: []*types.ToCSection{
						{
							ID:    "_Section_1",
							Level: 1,
						},
						{
							ID:    "_Section_2",
							Level: 1,
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})
	})
})
