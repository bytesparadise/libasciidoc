package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
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
						Content: types.InlineElements{
							types.StringElement{Content: "The Dangerous and Thrilling Documentation Chronicles"},
						},
					},
				},
				ElementReferences: map[string]interface{}{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: map[string]interface{}{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "This journey begins on a bleary Monday morning."},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})

		Context("document authors", func() {

			Context("single author", func() {

				It("all data", func() {
					actualContent := `= title
Kismet  Rainbow Chameleon  <kismet@asciidoctor.org>`
					expectedResult := types.DocumentHeader{
						Content: types.DocumentAttributes{
							"doctitle": types.SectionTitle{
								Attributes: map[string]interface{}{
									types.AttrID: "_title",
								},
								Content: types.InlineElements{
									types.StringElement{
										Content: "title",
									},
								},
							},
							"author":         "Kismet Rainbow Chameleon",
							"firstname":      "Kismet",
							"middlename":     "Rainbow",
							"lastname":       "Chameleon",
							"authorinitials": "KRC",
							"email":          "kismet@asciidoctor.org",
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentHeader"))
				})

				It("lastname with underscores", func() {
					actualContent := `= title
Lazarus het_Draeke <lazarus@asciidoctor.org>`
					expectedResult := types.DocumentHeader{
						Content: types.DocumentAttributes{
							"doctitle": types.SectionTitle{
								Attributes: map[string]interface{}{
									types.AttrID: "_title",
								},
								Content: types.InlineElements{
									types.StringElement{
										Content: "title",
									},
								},
							},
							"author":         "Lazarus het Draeke",
							"firstname":      "Lazarus",
							"lastname":       "het Draeke",
							"authorinitials": "Lh",
							"email":          "lazarus@asciidoctor.org",
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentHeader"))
				})

				It("firstname and lastname only", func() {
					actualContent := `= title
Kismet Chameleon`
					expectedResult := types.DocumentHeader{
						Content: types.DocumentAttributes{
							"doctitle": types.SectionTitle{
								Attributes: map[string]interface{}{
									types.AttrID: "_title",
								},
								Content: types.InlineElements{
									types.StringElement{
										Content: "title",
									},
								},
							},
							"author":         "Kismet Chameleon",
							"firstname":      "Kismet",
							"lastname":       "Chameleon",
							"authorinitials": "KC",
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentHeader"))
				})

				It("firstname only", func() {
					actualContent := `= title
Chameleon`
					expectedResult := types.DocumentHeader{
						Content: types.DocumentAttributes{
							"doctitle": types.SectionTitle{
								Attributes: map[string]interface{}{
									types.AttrID: "_title",
								},
								Content: types.InlineElements{
									types.StringElement{
										Content: "title",
									},
								},
							},
							"author":         "Chameleon",
							"firstname":      "Chameleon",
							"authorinitials": "C",
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentHeader"))
				})

				It("alternate author input", func() {
					actualContent := `= title
:author: Kismet Rainbow Chameleon` // `:"email":` is processed as a regular attribute
					expectedResult := types.DocumentHeader{
						Content: types.DocumentAttributes{
							"doctitle": types.SectionTitle{
								Attributes: map[string]interface{}{
									types.AttrID: "_title",
								},
								Content: types.InlineElements{
									types.StringElement{
										Content: "title",
									},
								},
							},
							"author":         "Kismet Rainbow Chameleon",
							"firstname":      "Kismet",
							"middlename":     "Rainbow",
							"lastname":       "Chameleon",
							"authorinitials": "KRC",
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentHeader"))
				})
			})

			Context("multiple authors", func() {
				It("2 authors only", func() {
					actualContent := `= title
Kismet  Rainbow Chameleon  <kismet@asciidoctor.org>; Lazarus het_Draeke <lazarus@asciidoctor.org>`
					expectedResult := types.DocumentHeader{
						Content: types.DocumentAttributes{
							"doctitle": types.SectionTitle{
								Attributes: map[string]interface{}{
									types.AttrID: "_title",
								},
								Content: types.InlineElements{
									types.StringElement{
										Content: "title",
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
						},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentHeader"))
				})
			})
		})

		Context("document revision", func() {

			It("full document revision", func() {
				actualContent := `= title
				john doe
				v1.0, June 19, 2017: First incarnation`
				expectedResult := types.DocumentHeader{
					Content: types.DocumentAttributes{
						"doctitle": types.SectionTitle{
							Attributes: map[string]interface{}{
								types.AttrID: "_title",
							},
							Content: types.InlineElements{
								types.StringElement{
									Content: "title",
								},
							},
						},
						"author":         "john doe",
						"authorinitials": "jd",
						"firstname":      "john",
						"lastname":       "doe",
						"revnumber":      "1.0",
						"revdate":        "June 19, 2017",
						"revremark":      "First incarnation",
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentHeader"))
			})

			It("revision with revnumber and revdate only", func() {
				actualContent := `= title
				john doe
				v1.0, June 19, 2017`
				expectedResult := types.DocumentHeader{
					Content: types.DocumentAttributes{
						"doctitle": types.SectionTitle{
							Attributes: map[string]interface{}{
								types.AttrID: "_title",
							},
							Content: types.InlineElements{
								types.StringElement{
									Content: "title",
								},
							},
						},
						"author":         "john doe",
						"authorinitials": "jd",
						"firstname":      "john",
						"lastname":       "doe",
						"revnumber":      "1.0",
						"revdate":        "June 19, 2017",
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentHeader"))
			})

			It("revision with revnumber and revdate - with colon separator", func() {
				actualContent := `= title
				john doe
				1.0, June 19, 2017:`
				expectedResult := types.DocumentHeader{
					Content: types.DocumentAttributes{
						"doctitle": types.SectionTitle{
							Attributes: map[string]interface{}{
								types.AttrID: "_title",
							},
							Content: types.InlineElements{
								types.StringElement{
									Content: "title",
								},
							},
						},
						"author":         "john doe",
						"authorinitials": "jd",
						"firstname":      "john",
						"lastname":       "doe",
						"revnumber":      "1.0",
						"revdate":        "June 19, 2017",
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentHeader"))
			})
			It("revision with revnumber only - comma suffix", func() {
				actualContent := `= title
				john doe
				1.0,`
				expectedResult := types.DocumentHeader{
					Content: types.DocumentAttributes{
						"doctitle": types.SectionTitle{
							Attributes: map[string]interface{}{
								types.AttrID: "_title",
							},
							Content: types.InlineElements{
								types.StringElement{
									Content: "title",
								},
							},
						},
						"author":         "john doe",
						"authorinitials": "jd",
						"firstname":      "john",
						"lastname":       "doe",
						"revnumber":      "1.0",
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentHeader"))
			})

			It("revision with revdate as number - spaces and no prefix no suffix", func() {
				actualContent := `= title
				john doe
				1.0`
				expectedResult := types.DocumentHeader{
					Content: types.DocumentAttributes{
						"doctitle": types.SectionTitle{
							Attributes: map[string]interface{}{
								types.AttrID: "_title",
							},
							Content: types.InlineElements{
								types.StringElement{
									Content: "title",
								},
							},
						},
						"author":         "john doe",
						"authorinitials": "jd",
						"firstname":      "john",
						"lastname":       "doe",
						"revdate":        "1.0",
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentHeader"))
			})

			It("revision with revdate as alphanum - spaces and no prefix no suffix", func() {
				actualContent := `= title
				john doe
				1.0a`
				expectedResult := types.DocumentHeader{
					Content: types.DocumentAttributes{
						"doctitle": types.SectionTitle{
							Attributes: map[string]interface{}{
								types.AttrID: "_title",
							},
							Content: types.InlineElements{
								types.StringElement{
									Content: "title",
								},
							},
						},
						"author":         "john doe",
						"authorinitials": "jd",
						"firstname":      "john",
						"lastname":       "doe",
						"revdate":        "1.0a",
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentHeader"))
			})

			It("revision with revnumber only", func() {
				actualContent := `= title
				john doe
				v1.0:`
				expectedResult := types.DocumentHeader{
					Content: types.DocumentAttributes{
						"doctitle": types.SectionTitle{
							Attributes: map[string]interface{}{
								types.AttrID: "_title",
							},
							Content: types.InlineElements{
								types.StringElement{
									Content: "title",
								},
							},
						},
						"author":         "john doe",
						"authorinitials": "jd",
						"firstname":      "john",
						"lastname":       "doe",
						"revnumber":      "1.0",
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentHeader"))
			})

			It("revision with spaces and capital revnumber ", func() {
				actualContent := `= title
				john doe
				V1.0:`
				expectedResult := types.DocumentHeader{
					Content: types.DocumentAttributes{
						"doctitle": types.SectionTitle{
							Attributes: map[string]interface{}{
								types.AttrID: "_title",
							},
							Content: types.InlineElements{
								types.StringElement{
									Content: "title",
								},
							},
						},
						"author":         "john doe",
						"authorinitials": "jd",
						"firstname":      "john",
						"lastname":       "doe",
						"revnumber":      "1.0",
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentHeader"))
			})

			It("revision only - with comma separator", func() {
				actualContent := `= title
				john doe
				v1.0,`
				expectedResult := types.DocumentHeader{
					Content: types.DocumentAttributes{
						"doctitle": types.SectionTitle{
							Attributes: map[string]interface{}{
								types.AttrID: "_title",
							},
							Content: types.InlineElements{
								types.StringElement{
									Content: "title",
								},
							},
						},
						"author":         "john doe",
						"authorinitials": "jd",
						"firstname":      "john",
						"lastname":       "doe",
						"revnumber":      "1.0",
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentHeader"))
			})

			It("revision with revnumber plus comma and colon separators", func() {
				actualContent := `= title
				john doe
				v1.0,:`
				expectedResult := types.DocumentHeader{
					Content: types.DocumentAttributes{
						"doctitle": types.SectionTitle{
							Attributes: map[string]interface{}{
								types.AttrID: "_title",
							},
							Content: types.InlineElements{
								types.StringElement{
									Content: "title",
								},
							},
						},
						"author":         "john doe",
						"authorinitials": "jd",
						"firstname":      "john",
						"lastname":       "doe",
						"revnumber":      "1.0",
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentHeader"))
			})

			It("revision with revnumber plus colon separator", func() {
				actualContent := `= title
john doe
v1.0:`
				expectedResult := types.DocumentHeader{
					Content: types.DocumentAttributes{
						"doctitle": types.SectionTitle{
							Attributes: map[string]interface{}{
								types.AttrID: "_title",
							},
							Content: types.InlineElements{
								types.StringElement{
									Content: "title",
								},
							},
						},
						"author":         "john doe",
						"authorinitials": "jd",
						"firstname":      "john",
						"lastname":       "doe",
						"revnumber":      "1.0",
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentHeader"))
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
					Elements: []interface{}{
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
					Elements: []interface{}{
						types.DocumentAttributeDeclaration{Name: "toc"},
						types.DocumentAttributeDeclaration{Name: "date", Value: "2017-01-01"},
						types.DocumentAttributeDeclaration{Name: "author", Value: "Xavier"},
						types.Paragraph{
							Attributes: map[string]interface{}{},
							Lines: []types.InlineElements{
								{
									types.StringElement{Content: "a paragraph"},
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
					Elements: []interface{}{
						types.DocumentAttributeDeclaration{Name: "toc"},
						types.DocumentAttributeDeclaration{Name: "date", Value: "2017-01-01"},
						types.DocumentAttributeDeclaration{Name: "author", Value: "Xavier"},
						types.BlankLine{},
						types.Paragraph{
							Attributes: map[string]interface{}{},
							Lines: []types.InlineElements{
								{
									types.StringElement{Content: "a paragraph"},
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
					Elements: []interface{}{
						types.DocumentAttributeDeclaration{Name: "toc"},
						types.DocumentAttributeDeclaration{Name: "date", Value: "2017-01-01"},
						types.BlankLine{},
						types.DocumentAttributeDeclaration{Name: "author", Value: "Xavier"},
						types.BlankLine{},
						types.Paragraph{
							Attributes: map[string]interface{}{},
							Lines: []types.InlineElements{
								{
									types.StringElement{Content: "a paragraph"},
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
					Elements: []interface{}{
						types.Paragraph{
							Attributes: map[string]interface{}{},
							Lines: []types.InlineElements{
								{
									types.StringElement{Content: "a paragraph"},
								},
							},
						},
						types.BlankLine{},
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
					Elements: []interface{}{
						types.DocumentAttributeDeclaration{Name: "author", Value: "Xavier"},
						types.BlankLine{},
						types.Paragraph{
							Attributes: map[string]interface{}{},
							Lines: []types.InlineElements{
								{
									types.StringElement{Content: "a paragraph written by "},
									types.DocumentAttributeSubstitution{Name: "author"},
									types.StringElement{Content: "."},
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
					Elements: []interface{}{
						types.DocumentAttributeDeclaration{Name: "author", Value: "Xavier"},
						types.BlankLine{},
						types.DocumentAttributeReset{Name: "author1"},
						types.DocumentAttributeReset{Name: "author2"},
						types.Paragraph{
							Attributes: map[string]interface{}{},
							Lines: []types.InlineElements{
								{
									types.StringElement{Content: "a paragraph written by "},
									types.DocumentAttributeSubstitution{Name: "author"},
									types.StringElement{Content: "."},
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
						Content: types.InlineElements{
							types.StringElement{Content: "The Dangerous and Thrilling Documentation Chronicles"},
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
				Elements: []interface{}{
					types.TableOfContentsMacro{},
					types.BlankLine{},
					types.Paragraph{
						Attributes: map[string]interface{}{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "This journey begins on a bleary Monday morning."},
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
						Content: types.InlineElements{
							types.StringElement{Content: "a header"},
						},
					},
				},
				ElementReferences: map[string]interface{}{
					"_section_1": types.SectionTitle{
						Attributes: map[string]interface{}{
							types.AttrID: "_section_1",
						},
						Content: types.InlineElements{
							types.StringElement{Content: "section 1"},
						},
					},
				},
				Elements: []interface{}{
					types.Section{
						Level: 1,
						Title: types.SectionTitle{
							Attributes: map[string]interface{}{
								types.AttrID: "_section_1",
							},
							Content: types.InlineElements{
								types.StringElement{Content: "section 1"},
							},
						},
						Elements: []interface{}{
							types.BlankLine{},
							types.Paragraph{
								Attributes: map[string]interface{}{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "a paragraph with "},
										types.QuotedText{Kind: types.Bold,
											Elements: types.InlineElements{
												types.StringElement{Content: "bold content"},
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
				Elements: []interface{}{
					types.Paragraph{
						Attributes: map[string]interface{}{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "a paragraph"},
							},
							{
								types.StringElement{Content: ":toc:"},
							},
							{
								types.StringElement{Content: ":date: 2017-01-01"},
							},
							{
								types.StringElement{Content: ":author: Xavier"},
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
				Elements: []interface{}{
					types.Paragraph{
						Attributes: map[string]interface{}{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: ":@date: 2017-01-01"},
							},
							{
								types.StringElement{Content: ":{author}: Xavier"},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})
})
