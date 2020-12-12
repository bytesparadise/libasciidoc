package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("document processing", func() {

	Context("article docs", func() {

		It("should retain attributes passed in configuration", func() {
			source := `[source]
----
foo
----`
			expected := types.Document{
				Attributes: types.Attributes{
					types.AttrSyntaxHighlighter: "pygments",
				},
				Elements: []interface{}{
					types.ListingBlock{
						Attributes: types.Attributes{
							types.AttrStyle: types.Source,
						},
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "foo",
								},
							},
						},
					},
				},
			}
			Expect(ParseDocument(source, configuration.WithAttributes(map[string]string{
				types.AttrSyntaxHighlighter: "pygments",
			}))).To(Equal(expected))
		})

		It("should include toc and preamble", func() {
			source := `= A title
:toc:

Preamble comes here

== Section A

=== Section A.a

== Section B

== Section C`
			headerTitle := []interface{}{
				types.StringElement{
					Content: "A title",
				},
			}
			titleSectionA := []interface{}{
				types.StringElement{
					Content: "Section A",
				},
			}
			titleSectionAa := []interface{}{
				types.StringElement{
					Content: "Section A.a",
				},
			}
			titleSectionB := []interface{}{
				types.StringElement{
					Content: "Section B",
				},
			}
			titleSectionC := []interface{}{
				types.StringElement{
					Content: "Section C",
				},
			}
			expected := types.Document{
				Attributes: types.Attributes{
					types.AttrTableOfContents: nil,
				},
				Elements: []interface{}{
					types.Section{
						Level: 0,
						Attributes: types.Attributes{
							types.AttrID: "_a_title",
						},
						Title: headerTitle,
						Elements: []interface{}{
							types.TableOfContentsPlaceHolder{},
							types.Preamble{
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{
													Content: "Preamble comes here",
												},
											},
										},
									},
								},
							},
							types.Section{
								Level: 1,
								Attributes: types.Attributes{
									types.AttrID: "_section_a",
								},
								Title: titleSectionA,
								Elements: []interface{}{
									types.Section{
										Level: 2,
										Attributes: types.Attributes{
											types.AttrID: "_section_a_a",
										},
										Title:    titleSectionAa,
										Elements: []interface{}{},
									},
								},
							},
							types.Section{
								Level: 1,
								Attributes: types.Attributes{
									types.AttrID: "_section_b",
								},
								Title:    titleSectionB,
								Elements: []interface{}{},
							},
							types.Section{
								Level: 1,
								Attributes: types.Attributes{
									types.AttrID: "_section_c",
								},
								Title:    titleSectionC,
								Elements: []interface{}{},
							},
						},
					},
				},
				ElementReferences: types.ElementReferences{
					"_a_title":     headerTitle,
					"_section_a":   titleSectionA,
					"_section_a_a": titleSectionAa,
					"_section_b":   titleSectionB,
					"_section_c":   titleSectionC,
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
			headerSectionTitle := []interface{}{
				types.StringElement{
					Content: "eve(1)",
				},
			}
			nameSectionTitle := []interface{}{
				types.StringElement{
					Content: "Name",
				},
			}
			synopisSectionTitle := []interface{}{
				types.StringElement{
					Content: "Synopsis",
				},
			}

			expected := types.Document{
				Attributes: types.Attributes{
					types.AttrDocType: "manpage",
					types.AttrAuthors: []types.DocumentAuthor{
						{
							FullName: "Andrew Stanton",
						},
					},
					"firstname":      "Andrew",
					"lastname":       "Stanton",
					"author":         "Andrew Stanton",
					"authorinitials": "AS",
					types.AttrRevision: types.DocumentRevision{
						Revnumber: "1.0.0",
					},
					"revnumber": "1.0.0",
				},
				Elements: []interface{}{
					types.Section{
						Level: 0,
						Attributes: types.Attributes{
							types.AttrID: "_eve_1",
						},
						Title: headerSectionTitle,
						Elements: []interface{}{
							types.Section{
								Level: 1,
								Attributes: types.Attributes{
									types.AttrID: "_name",
								},
								Title: nameSectionTitle,
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{
													Content: "eve - analyzes an image to determine if it\u2019s a picture of a life form",
												},
											},
										},
									},
								},
							},
							types.Section{
								Level: 1,
								Attributes: types.Attributes{
									types.AttrID: "_synopsis",
								},
								Title:    synopisSectionTitle,
								Elements: []interface{}{},
							},
						},
					},
				},
				ElementReferences: types.ElementReferences{
					"_eve_1":    headerSectionTitle,
					"_name":     nameSectionTitle,
					"_synopsis": synopisSectionTitle,
				},
			}
			Expect(ParseDocument(source,
				configuration.WithAttributes(map[string]string{
					types.AttrDocType: "manpage",
				},
				))).To(MatchDocument(expected))
		})

	})
})
