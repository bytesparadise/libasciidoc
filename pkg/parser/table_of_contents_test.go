package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("tables of contents", func() {

	Context("in final documents", func() {

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
				Sections: []*types.ToCSection{
					{
						ID:    "_section_a",
						Level: 1,
						Title: "Section A",
						Children: []*types.ToCSection{
							{
								ID:    "_section_a_a",
								Level: 2,
								Title: "Section A.a",
							},
							{
								ID:    "_section_a_b",
								Level: 2,
								Title: "Section A.b",
							},
						},
					},
					{
						ID:    "_section_b",
						Level: 1,
						Title: "Section B",
						Children: []*types.ToCSection{
							{
								ID:    "_section_b_a",
								Level: 2,
								Title: "Section B.a",
							},
						},
					},
					{
						ID:    "_section_c",
						Level: 1,
						Title: "Section C",
					},
				},
			}
			ctx := parser.NewParseContext(configuration.NewConfiguration())
			_, toc, err := parser.Aggregate(ctx, c)
			Expect(err).ToNot(HaveOccurred())
			Expect(toc).To(MatchTableOfContents(expectedToC))
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
				Sections: []*types.ToCSection{
					{
						ID:    "_section_a",
						Level: 1,
						Title: "Section A",
						Children: []*types.ToCSection{
							{
								ID:    "_section_a_a",
								Level: 2,
								Title: "Section A.a",
							},
							{
								ID:    "_section_a_b",
								Level: 2,
								Title: "Section A.b",
								Children: []*types.ToCSection{
									{
										ID:    "_section_that_shall_be_in_ToC",
										Level: 3,
										Title: "Section that shall be in ToC",
									},
								},
							},
						},
					},
					{
						ID:    "_section_b",
						Level: 1,
						Title: "Section B",
						Children: []*types.ToCSection{
							{
								ID:    "_section_b_a",
								Level: 2,
								Title: "Section B.a",
							},
						},
					},
					{
						ID:    "_section_c",
						Level: 1,
						Title: "Section C",
					},
				},
			}
			ctx := parser.NewParseContext(configuration.NewConfiguration())
			_, toc, err := parser.Aggregate(ctx, c)
			Expect(err).ToNot(HaveOccurred())
			Expect(toc).To(MatchTableOfContents(expectedToC))
		})
	})
})
