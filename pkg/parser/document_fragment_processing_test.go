package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("document processing", func() {

	Context("article docs", func() {

		It("should retain attributes passed in configuration", func() {
			source := `[source]
----
foo
----`
			expected := &types.Document{
				Elements: []interface{}{
					&types.DelimitedBlock{
						Kind: types.Listing,
						Attributes: types.Attributes{
							types.AttrStyle: types.Source,
						},
						Elements: []interface{}{
							&types.StringElement{
								Content: "foo",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source, configuration.WithAttributes(map[string]interface{}{
				types.AttrSyntaxHighlighter: "pygments",
			}))).To(MatchDocument(expected))
		})

		It("should include toc and preamble", func() {
			source := `= A title
:toc:

Preamble comes here

== Section A

=== Section A.a

== Section B

== Section C`
			titleSectionA := []interface{}{
				&types.StringElement{
					Content: "Section A",
				},
			}
			titleSectionAa := []interface{}{
				&types.StringElement{
					Content: "Section A.a",
				},
			}
			titleSectionB := []interface{}{
				&types.StringElement{
					Content: "Section B",
				},
			}
			titleSectionC := []interface{}{
				&types.StringElement{
					Content: "Section C",
				},
			}
			expected := &types.Document{
				Elements: []interface{}{
					&types.DocumentHeader{
						Title: []interface{}{
							&types.StringElement{
								Content: "A title",
							},
						},
						Elements: []interface{}{
							&types.AttributeDeclaration{
								Name: "toc",
							},
						},
					},
					&types.Preamble{
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "Preamble comes here",
									},
								},
							},
						},
					},
					&types.Section{
						Level: 1,
						Attributes: types.Attributes{
							types.AttrID: "_Section_A",
						},
						Title: titleSectionA,
						Elements: []interface{}{
							&types.Section{
								Level: 2,
								Attributes: types.Attributes{
									types.AttrID: "_Section_A_a",
								},
								Title: titleSectionAa,
							},
						},
					},
					&types.Section{
						Level: 1,
						Attributes: types.Attributes{
							types.AttrID: "_Section_B",
						},
						Title: titleSectionB,
					},
					&types.Section{
						Level: 1,
						Attributes: types.Attributes{
							types.AttrID: "_Section_C",
						},
						Title: titleSectionC,
					},
				},
				ElementReferences: types.ElementReferences{
					"_Section_A":   titleSectionA,
					"_Section_A_a": titleSectionAa,
					"_Section_B":   titleSectionB,
					"_Section_C":   titleSectionC,
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
						{
							ID:    "_Section_C",
							Level: 1,
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})
	})

	Context("manpage docs", func() {

		It("manpage without preamble", func() {
			source := `= eve(1)
Andrew Stanton
v1.0.0

== Name

eve - analyzes an image to determine if it's a picture of a life form

== Synopsis
`
			nameSectionTitle := []interface{}{
				&types.StringElement{
					Content: "Name",
				},
			}
			synopisSectionTitle := []interface{}{
				&types.StringElement{
					Content: "Synopsis",
				},
			}

			expected := &types.Document{
				Elements: []interface{}{
					&types.DocumentHeader{
						Title: []interface{}{
							&types.StringElement{
								Content: "eve(1)",
							},
						},
						Elements: []interface{}{
							&types.AttributeDeclaration{
								Name: types.AttrAuthors,
								Value: types.DocumentAuthors{
									{
										DocumentAuthorFullName: &types.DocumentAuthorFullName{
											FirstName: "Andrew",
											LastName:  "Stanton",
										},
									},
								},
							},
							&types.AttributeDeclaration{
								Name: types.AttrRevision,
								Value: &types.DocumentRevision{
									Revnumber: "1.0.0",
								},
							},
						},
					},
					&types.Section{
						Level: 1,
						Attributes: types.Attributes{
							types.AttrID: "_Name",
						},
						Title: nameSectionTitle,
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "eve - analyzes an image to determine if it",
									},
									&types.Symbol{
										Name: "'",
									},
									&types.StringElement{
										Content: "s a picture of a life form",
									},
								},
							},
						},
					},
					&types.Section{
						Level: 1,
						Attributes: types.Attributes{
							types.AttrID: "_Synopsis",
						},
						Title: synopisSectionTitle,
					},
				},
				ElementReferences: types.ElementReferences{
					"_Name":     nameSectionTitle,
					"_Synopsis": synopisSectionTitle,
				},
				TableOfContents: &types.TableOfContents{
					MaxDepth: 2,
					Sections: []*types.ToCSection{
						{
							ID:    "_Name",
							Level: 1,
						},
						{
							ID:    "_Synopsis",
							Level: 1,
						},
					},
				},
			}
			Expect(ParseDocument(source,
				configuration.WithAttributes(map[string]interface{}{
					types.AttrDocType: "manpage",
				},
				))).To(MatchDocument(expected))
		})

	})
})
