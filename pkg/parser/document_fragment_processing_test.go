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
							types.AttrID: "_section_a",
						},
						Title: titleSectionA,
						Elements: []interface{}{
							&types.Section{
								Level: 2,
								Attributes: types.Attributes{
									types.AttrID: "_section_a_a",
								},
								Title: titleSectionAa,
							},
						},
					},
					&types.Section{
						Level: 1,
						Attributes: types.Attributes{
							types.AttrID: "_section_b",
						},
						Title: titleSectionB,
					},
					&types.Section{
						Level: 1,
						Attributes: types.Attributes{
							types.AttrID: "_section_c",
						},
						Title: titleSectionC,
					},
				},
				ElementReferences: types.ElementReferences{
					"_section_a":   titleSectionA,
					"_section_a_a": titleSectionAa,
					"_section_b":   titleSectionB,
					"_section_c":   titleSectionC,
				},
				TableOfContents: &types.TableOfContents{
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
							},
						},
						{
							ID:    "_section_b",
							Level: 1,
						},
						{
							ID:    "_section_c",
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
							types.AttrID: "_name",
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
							types.AttrID: "_synopsis",
						},
						Title: synopisSectionTitle,
					},
				},
				ElementReferences: types.ElementReferences{
					"_name":     nameSectionTitle,
					"_synopsis": synopisSectionTitle,
				},
				TableOfContents: &types.TableOfContents{
					MaxDepth: 2,
					Sections: []*types.ToCSection{
						{
							ID:    "_name",
							Level: 1,
						},
						{
							ID:    "_synopsis",
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
