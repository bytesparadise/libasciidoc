package parser_test

import (
	"github.com/bytesparadise/libasciidoc/parser"
	"github.com/bytesparadise/libasciidoc/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("document attributes", func() {

	Context("valid Document Header", func() {

		It("header alone", func() {
			actualContent := `= The Dangerous and Thrilling Documentation Chronicles
			
This journey begins on a bleary Monday morning.`
			expectedResult := types.Document{
				Attributes: map[string]interface{}{
					"doctitle": types.SectionTitle{
						Attributes: map[string]interface{}{
							types.AttrID: "_the_dangerous_and_thrilling_documentation_chronicles",
						},
						Content: types.InlineContent{
							Elements: []types.InlineElement{
								types.StringElement{Content: "The Dangerous and Thrilling Documentation Chronicles"},
							},
						},
					},
				},
				ElementReferences: map[string]interface{}{},
				Elements: []types.DocElement{
					types.Paragraph{
						Attributes: map[string]interface{}{},
						Lines: []types.InlineContent{
							{
								Elements: []types.InlineElement{
									types.StringElement{Content: "This journey begins on a bleary Monday morning."},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})

		Context("document Authors", func() {

			Context("single Author", func() {

				It("all data", func() {
					actualContent := `Kismet  Rainbow Chameleon  <kismet@asciidoctor.org>`
					expectedResult := []types.DocumentAuthor{
						{
							FullName:   "Kismet Rainbow Chameleon",
							FirstName:  "Kismet",
							MiddleName: "Rainbow",
							LastName:   "Chameleon",
							Initials:   "KRC",
							Email:      "kismet@asciidoctor.org",
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentAuthors"))
				})

				It("lastname with underscores", func() {
					actualContent := `Lazarus het_Draeke <lazarus@asciidoctor.org>`
					expectedResult := []types.DocumentAuthor{
						{
							FullName:  "Lazarus het Draeke",
							FirstName: "Lazarus",
							LastName:  "het Draeke",
							Initials:  "Lh",
							Email:     "lazarus@asciidoctor.org",
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentAuthors"))
				})

				It("firstname and lastname only", func() {
					actualContent := `Kismet Chameleon`
					expectedResult := []types.DocumentAuthor{
						{
							FullName:  "Kismet Chameleon",
							FirstName: "Kismet",
							LastName:  "Chameleon",
							Initials:  "KC",
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentAuthors"))
				})

				It("firstname only", func() {
					actualContent := `  Chameleon`
					expectedResult := []types.DocumentAuthor{
						{
							FullName:  "Chameleon",
							FirstName: "Chameleon",
							Initials:  "C",
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentAuthors"))
				})

				It("alternate author input", func() {
					actualContent := `:author: Kismet Rainbow Chameleon` // `:email:` is processed as a regular attribute
					expectedResult := []types.DocumentAuthor{
						{
							FullName:   "Kismet Rainbow Chameleon",
							FirstName:  "Kismet",
							MiddleName: "Rainbow",
							LastName:   "Chameleon",
							Initials:   "KRC",
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentAuthors"))
				})
			})

			Context("multiple authors", func() {
				It("2 authors only", func() {
					actualContent := `Kismet  Rainbow Chameleon  <kismet@asciidoctor.org>; Lazarus het_Draeke <lazarus@asciidoctor.org>`
					expectedResult := []types.DocumentAuthor{
						{
							FullName:   "Kismet Rainbow Chameleon",
							FirstName:  "Kismet",
							MiddleName: "Rainbow",
							LastName:   "Chameleon",
							Initials:   "KRC",
							Email:      "kismet@asciidoctor.org",
						},
						{
							FullName:  "Lazarus het Draeke",
							FirstName: "Lazarus",
							LastName:  "het Draeke",
							Initials:  "Lh",
							Email:     "lazarus@asciidoctor.org",
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentAuthors"))
				})
			})
		})

		Context("document Revision", func() {

			It("Full document revision", func() {
				actualContent := `v1.0, June 19, 2017: First incarnation`
				expectedResult := types.DocumentRevision{
					Revnumber: "1.0",
					Revdate:   "June 19, 2017",
					Revremark: "First incarnation",
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentRevision"))
			})

			It("revision with revnumber and revdate only", func() {
				actualContent := `v1.0, June 19, 2017`
				expectedResult := types.DocumentRevision{
					Revnumber: "1.0",
					Revdate:   "June 19, 2017",
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentRevision"))
			})

			It("revision with revnumber and revdate - with colon separator", func() {
				actualContent := `v1.0, June 19, 2017:`
				expectedResult := types.DocumentRevision{
					Revnumber: "1.0",
					Revdate:   "June 19, 2017",
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentRevision"))
			})
			It("revision with revnumber only - comma suffix", func() {
				actualContent := `1.0,`
				expectedResult := types.DocumentRevision{
					Revnumber: "1.0",
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentRevision"))
			})

			It("revision with revdate as number - spaces and no prefix no suffix", func() {
				actualContent := `1.0`
				expectedResult := types.DocumentRevision{
					Revdate: "1.0",
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentRevision"))
			})

			It("revision with revdate as alphanum - spaces and no prefix no suffix", func() {
				actualContent := `1.0a`
				expectedResult := types.DocumentRevision{
					Revdate: "1.0a",
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentRevision"))
			})

			It("revision with revnumber only", func() {
				actualContent := `v1.0:`
				expectedResult := types.DocumentRevision{
					Revnumber: "1.0",
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentRevision"))
			})

			It("revision with spaces and capital revnumber ", func() {
				actualContent := `V1.0:`
				expectedResult := types.DocumentRevision{
					Revnumber: "1.0",
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentRevision"))
			})

			It("revision only - with comma separator", func() {
				actualContent := `v1.0,`
				expectedResult := types.DocumentRevision{
					Revnumber: "1.0",
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentRevision"))
			})

			It("revision with revnumber plus comma and colon separators", func() {
				actualContent := `v1.0,:`
				expectedResult := types.DocumentRevision{
					Revnumber: "1.0",
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentRevision"))
			})

			It("revision with revnumber plus colon separator", func() {
				actualContent := `v1.0:`
				expectedResult := types.DocumentRevision{
					Revnumber: "1.0",
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentRevision"))
			})

		})

		Context("document Header Attributes", func() {

			It("valid attribute names", func() {
				actualContent := `:a:
:author: Xavier
:_author: Xavier
:Author: Xavier
:0Author: Xavier
:Auth0r: Xavier`
				expectedResult := types.Document{
					Attributes:        map[string]interface{}{},
					ElementReferences: map[string]interface{}{},
					Elements: []types.DocElement{
						types.DocumentAttributeDeclaration{Name: "a"},
						types.DocumentAttributeDeclaration{Name: "author", Value: "Xavier"},
						types.DocumentAttributeDeclaration{Name: "_author", Value: "Xavier"},
						types.DocumentAttributeDeclaration{Name: "Author", Value: "Xavier"},
						types.DocumentAttributeDeclaration{Name: "0Author", Value: "Xavier"},
						types.DocumentAttributeDeclaration{Name: "Auth0r", Value: "Xavier"},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent)
			})

			It("attributes and paragraph without blank line in-between", func() {
				actualContent := `:toc:
:date:  2017-01-01
:author: Xavier
a paragraph`
				expectedResult := types.Document{
					Attributes:        map[string]interface{}{},
					ElementReferences: map[string]interface{}{},
					Elements: []types.DocElement{
						types.DocumentAttributeDeclaration{Name: "toc"},
						types.DocumentAttributeDeclaration{Name: "date", Value: "2017-01-01"},
						types.DocumentAttributeDeclaration{Name: "author", Value: "Xavier"},
						types.Paragraph{
							Attributes: map[string]interface{}{},
							Lines: []types.InlineContent{
								{
									Elements: []types.InlineElement{
										types.StringElement{Content: "a paragraph"},
									},
								},
							},
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent)
			})

			It("contiguous attributes and paragraph with blank line in-between", func() {
				actualContent := `:toc:
:date: 2017-01-01
:author: Xavier

a paragraph`
				expectedResult := types.Document{
					Attributes:        map[string]interface{}{},
					ElementReferences: map[string]interface{}{},
					Elements: []types.DocElement{
						types.DocumentAttributeDeclaration{Name: "toc"},
						types.DocumentAttributeDeclaration{Name: "date", Value: "2017-01-01"},
						types.DocumentAttributeDeclaration{Name: "author", Value: "Xavier"},
						types.Paragraph{
							Attributes: map[string]interface{}{},
							Lines: []types.InlineContent{
								{
									Elements: []types.InlineElement{
										types.StringElement{Content: "a paragraph"},
									},
								},
							},
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent)
			})

			It("splitted attributes and paragraph with blank line in-between", func() {
				actualContent := `:toc:
:date: 2017-01-01

:author: Xavier

a paragraph`
				expectedResult := types.Document{
					Attributes:        map[string]interface{}{},
					ElementReferences: map[string]interface{}{},
					Elements: []types.DocElement{
						types.DocumentAttributeDeclaration{Name: "toc"},
						types.DocumentAttributeDeclaration{Name: "date", Value: "2017-01-01"},
						types.DocumentAttributeDeclaration{Name: "author", Value: "Xavier"},
						types.Paragraph{
							Attributes: map[string]interface{}{},
							Lines: []types.InlineContent{
								{
									Elements: []types.InlineElement{
										types.StringElement{Content: "a paragraph"},
									},
								},
							},
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent)
			})

			It("no header and attributes in body", func() {
				actualContent := `a paragraph
	
:toc:
:date: 2017-01-01
:author: Xavier`
				expectedResult := types.Document{
					Attributes:        map[string]interface{}{},
					ElementReferences: map[string]interface{}{},
					Elements: []types.DocElement{
						types.Paragraph{
							Attributes: map[string]interface{}{},
							Lines: []types.InlineContent{
								{
									Elements: []types.InlineElement{
										types.StringElement{Content: "a paragraph"},
									},
								},
							},
						},
						types.DocumentAttributeDeclaration{Name: "toc"},
						types.DocumentAttributeDeclaration{Name: "date", Value: "2017-01-01"},
						types.DocumentAttributeDeclaration{Name: "author", Value: "Xavier"},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent)
			})
		})

		Context("document Attribute Substitutions", func() {

			It("paragraph with attribute substitution", func() {
				actualContent := `:author: Xavier
			
a paragraph written by {author}.`
				expectedResult := types.Document{
					Attributes:        map[string]interface{}{},
					ElementReferences: map[string]interface{}{},
					Elements: []types.DocElement{
						types.DocumentAttributeDeclaration{Name: "author", Value: "Xavier"},
						types.Paragraph{
							Attributes: map[string]interface{}{},
							Lines: []types.InlineContent{
								{
									Elements: []types.InlineElement{
										types.StringElement{Content: "a paragraph written by "},
										types.DocumentAttributeSubstitution{Name: "author"},
										types.StringElement{Content: "."},
									},
								},
							},
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent)
			})

			It("paragraph with attribute resets", func() {
				actualContent := `:author: Xavier
							
:!author1:
:author2!:
a paragraph written by {author}.`
				expectedResult := types.Document{
					Attributes:        map[string]interface{}{},
					ElementReferences: map[string]interface{}{},
					Elements: []types.DocElement{
						types.DocumentAttributeDeclaration{Name: "author", Value: "Xavier"},
						types.DocumentAttributeReset{Name: "author1"},
						types.DocumentAttributeReset{Name: "author2"},
						types.Paragraph{
							Attributes: map[string]interface{}{},
							Lines: []types.InlineContent{
								{
									Elements: []types.InlineElement{
										types.StringElement{Content: "a paragraph written by "},
										types.DocumentAttributeSubstitution{Name: "author"},
										types.StringElement{Content: "."},
									},
								},
							},
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent)
			})
		})

		It("header with 2 authors, revision and attributes", func() {
			actualContent := `= The Dangerous and Thrilling Documentation Chronicles
Kismet Rainbow Chameleon <kismet@asciidoctor.org>; Lazarus het_Draeke <lazarus@asciidoctor.org>
v1.0, June 19, 2017: First incarnation
:toc:
:keywords: documentation, team, obstacles, journey, victory

This journey begins on a bleary Monday morning.`
			expectedResult := types.Document{
				Attributes: map[string]interface{}{
					"doctitle": types.SectionTitle{
						Attributes: map[string]interface{}{
							types.AttrID: "_the_dangerous_and_thrilling_documentation_chronicles",
						},
						Content: types.InlineContent{
							Elements: []types.InlineElement{
								types.StringElement{Content: "The Dangerous and Thrilling Documentation Chronicles"},
							},
						},
					},
					"author":           "Kismet Rainbow Chameleon",
					"firstname":        "Kismet",
					"middlename":       "Rainbow",
					"lastname":         "Chameleon",
					"authorinitials":   "KRC",
					"email":            "kismet@asciidoctor.org",
					"author_2":         "Lazarus het Draeke",
					"firstname_2":      "Lazarus",
					"lastname_2":       "het Draeke",
					"authorinitials_2": "Lh",
					"email_2":          "lazarus@asciidoctor.org",
					"revnumber":        "1.0",
					"revdate":          "June 19, 2017",
					"revremark":        "First incarnation",
					"keywords":         "documentation, team, obstacles, journey, victory",
					"toc":              "",
				},
				ElementReferences: map[string]interface{}{},
				Elements: []types.DocElement{
					types.TableOfContentsMacro{},
					types.Paragraph{
						Attributes: map[string]interface{}{},
						Lines: []types.InlineContent{
							{
								Elements: []types.InlineElement{
									types.StringElement{Content: "This journey begins on a bleary Monday morning."},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("header section inline with bold quote", func() {

			actualContent := `= a header
				
== section 1

a paragraph with *bold content*`
			expectedResult := types.Document{
				Attributes: map[string]interface{}{
					"doctitle": types.SectionTitle{
						Attributes: map[string]interface{}{
							types.AttrID: "_a_header",
						},
						Content: types.InlineContent{
							Elements: []types.InlineElement{
								types.StringElement{Content: "a header"},
							},
						},
					},
				},
				ElementReferences: map[string]interface{}{
					"_section_1": types.SectionTitle{
						Attributes: map[string]interface{}{
							types.AttrID: "_section_1",
						},
						Content: types.InlineContent{
							Elements: []types.InlineElement{
								types.StringElement{Content: "section 1"},
							},
						},
					},
				},
				Elements: []types.DocElement{
					types.Section{
						Level: 1,
						Title: types.SectionTitle{
							Attributes: map[string]interface{}{
								types.AttrID: "_section_1",
							},
							Content: types.InlineContent{
								Elements: []types.InlineElement{
									types.StringElement{Content: "section 1"},
								},
							},
						},
						Elements: []types.DocElement{
							types.Paragraph{
								Attributes: map[string]interface{}{},
								Lines: []types.InlineContent{
									types.InlineContent{
										Elements: []types.InlineElement{
											types.StringElement{Content: "a paragraph with "},
											types.QuotedText{Kind: types.Bold,
												Elements: []types.InlineElement{
													types.StringElement{Content: "bold content"},
												},
											},
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

	Context("invalid document attributes", func() {

		It("paragraph without blank line before attribute declarations", func() {
			actualContent := `a paragraph
:toc:
:date: 2017-01-01
:author: Xavier`
			expectedResult := types.Document{
				Attributes:        map[string]interface{}{},
				ElementReferences: map[string]interface{}{},
				Elements: []types.DocElement{
					types.Paragraph{
						Attributes: map[string]interface{}{},
						Lines: []types.InlineContent{
							{
								Elements: []types.InlineElement{
									types.StringElement{Content: "a paragraph"},
								},
							},
							{
								Elements: []types.InlineElement{
									types.StringElement{Content: ":toc:"},
								},
							},
							{
								Elements: []types.InlineElement{
									types.StringElement{Content: ":date: 2017-01-01"},
								},
							},
							{
								Elements: []types.InlineElement{
									types.StringElement{Content: ":author: Xavier"},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("invalid attribute names", func() {
			actualContent := `:@date: 2017-01-01
:{author}: Xavier`
			expectedResult := types.Document{
				Attributes:        map[string]interface{}{},
				ElementReferences: map[string]interface{}{},
				Elements: []types.DocElement{
					types.Paragraph{
						Attributes: map[string]interface{}{},
						Lines: []types.InlineContent{
							{
								Elements: []types.InlineElement{
									types.StringElement{Content: ":@date: 2017-01-01"},
								},
							},
							{
								Elements: []types.InlineElement{
									types.StringElement{Content: ":{author}: Xavier"},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})
})
