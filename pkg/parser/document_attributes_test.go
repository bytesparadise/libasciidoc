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

			title := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID:       "the_dangerous_and_thrilling_documentation_chronicles",
					types.AttrCustomID: false,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "The Dangerous and Thrilling Documentation Chronicles"},
				},
			}
			expectedResult := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"the_dangerous_and_thrilling_documentation_chronicles": title,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Level:      0,
						Title:      title,
						Attributes: types.ElementAttributes{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "This journey begins on a bleary Monday morning."},
									},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent)
		})

		Context("document authors", func() {

			Context("single author", func() {

				It("all author data", func() {
					actualContent := `= title
Kismet  Rainbow Chameleon  <kismet@asciidoctor.org>`
					expectedResult := types.Section{
						Level: 0,
						Title: types.SectionTitle{
							Attributes: types.ElementAttributes{
								types.AttrID:       "title",
								types.AttrCustomID: false,
							},
							Elements: types.InlineElements{
								types.StringElement{
									Content: "title",
								},
							},
						},
						Attributes: types.ElementAttributes{
							types.AttrAuthors: []types.DocumentAuthor{
								{
									FullName: "Kismet  Rainbow Chameleon  ",
									Email:    "kismet@asciidoctor.org",
								},
							},
						},
						Elements: []interface{}{},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Section0WithMetadata"))
				})

				It("lastname with underscores", func() {
					actualContent := `= title
Lazarus het_Draeke <lazarus@asciidoctor.org>`
					expectedResult := types.Section{
						Level: 0,
						Title: types.SectionTitle{
							Attributes: types.ElementAttributes{
								types.AttrID:       "title",
								types.AttrCustomID: false,
							},
							Elements: types.InlineElements{
								types.StringElement{
									Content: "title",
								},
							},
						},
						Attributes: types.ElementAttributes{
							types.AttrAuthors: []types.DocumentAuthor{
								{
									FullName: "Lazarus het_Draeke ",
									Email:    "lazarus@asciidoctor.org",
								},
							},
						},
						Elements: []interface{}{},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Section0WithMetadata"))
				})

				It("firstname and lastname only", func() {
					actualContent := `= title
Kismet Chameleon`
					expectedResult := types.Section{
						Level: 0,
						Title: types.SectionTitle{
							Attributes: types.ElementAttributes{
								types.AttrID:       "title",
								types.AttrCustomID: false,
							},
							Elements: types.InlineElements{
								types.StringElement{
									Content: "title",
								},
							},
						},
						Attributes: types.ElementAttributes{
							types.AttrAuthors: []types.DocumentAuthor{
								{
									FullName: "Kismet Chameleon",
									Email:    "",
								},
							},
						},
						Elements: []interface{}{},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Section0WithMetadata"))
				})

				It("firstname only", func() {
					actualContent := `= title
Chameleon`
					expectedResult := types.Section{
						Level: 0,
						Title: types.SectionTitle{
							Attributes: types.ElementAttributes{
								types.AttrID:       "title",
								types.AttrCustomID: false,
							},
							Elements: types.InlineElements{
								types.StringElement{
									Content: "title",
								},
							},
						},
						Attributes: types.ElementAttributes{
							types.AttrAuthors: []types.DocumentAuthor{
								{
									FullName: "Chameleon",
									Email:    "",
								},
							},
						},
						Elements: []interface{}{},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Section0WithMetadata"))
				})

				It("alternate author input", func() {
					actualContent := `= title
:author: Kismet Rainbow Chameleon` // `:"email":` is processed as a regular attribute
					expectedResult := types.Section{
						Level: 0,
						Title: types.SectionTitle{
							Attributes: types.ElementAttributes{
								types.AttrID:       "title",
								types.AttrCustomID: false,
							},
							Elements: types.InlineElements{
								types.StringElement{
									Content: "title",
								},
							},
						},
						Attributes: types.ElementAttributes{
							types.AttrAuthors: []types.DocumentAuthor{
								{
									FullName: "Kismet Rainbow Chameleon",
									Email:    "",
								},
							},
						},
						Elements: []interface{}{},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Section0WithMetadata"))
				})
			})

			Context("multiple authors", func() {

				It("2 authors only", func() {
					actualContent := `= title
Kismet  Rainbow Chameleon  <kismet@asciidoctor.org>; Lazarus het_Draeke <lazarus@asciidoctor.org>`
					expectedResult := types.Section{
						Level: 0,
						Title: types.SectionTitle{
							Attributes: types.ElementAttributes{
								types.AttrID:       "title",
								types.AttrCustomID: false,
							},
							Elements: types.InlineElements{
								types.StringElement{
									Content: "title",
								},
							},
						},
						Attributes: types.ElementAttributes{
							types.AttrAuthors: []types.DocumentAuthor{
								{
									FullName: "Kismet  Rainbow Chameleon  ",
									Email:    "kismet@asciidoctor.org",
								},
								{
									FullName: "Lazarus het_Draeke ",
									Email:    "lazarus@asciidoctor.org",
								},
							},
						},
						Elements: []interface{}{},
					}
					verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Section0WithMetadata"))
				})
			})
		})

		Context("document revision", func() {

			It("full document revision", func() {
				actualContent := `= title
				john doe
				v1.0, June 19, 2017: First incarnation`
				expectedResult := types.Section{
					Level: 0,
					Title: types.SectionTitle{
						Attributes: types.ElementAttributes{
							types.AttrID:       "title",
							types.AttrCustomID: false,
						},
						Elements: types.InlineElements{
							types.StringElement{
								Content: "title",
							},
						},
					},
					Attributes: types.ElementAttributes{
						types.AttrAuthors: []types.DocumentAuthor{
							{
								FullName: "john doe",
								Email:    "",
							},
						},
						types.AttrRevision: types.DocumentRevision{
							Revnumber: "1.0",
							Revdate:   "June 19, 2017",
							Revremark: "First incarnation",
						},
					},
					Elements: []interface{}{},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Section0WithMetadata"))
			})

			It("revision with revnumber and revdate only", func() {
				actualContent := `= title
				john doe
				v1.0, June 19, 2017`
				expectedResult := types.Section{
					Level: 0,
					Title: types.SectionTitle{
						Attributes: types.ElementAttributes{
							types.AttrID:       "title",
							types.AttrCustomID: false,
						},
						Elements: types.InlineElements{
							types.StringElement{
								Content: "title",
							},
						},
					},
					Attributes: types.ElementAttributes{
						types.AttrAuthors: []types.DocumentAuthor{
							{
								FullName: "john doe",
								Email:    "",
							},
						},
						types.AttrRevision: types.DocumentRevision{
							Revnumber: "1.0",
							Revdate:   "June 19, 2017",
							Revremark: "",
						},
					},
					Elements: []interface{}{},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Section0WithMetadata"))
			})

			It("revision with revnumber and revdate - with colon separator", func() {
				actualContent := `= title
				john doe
				1.0, June 19, 2017:`
				expectedResult := types.Section{
					Level: 0,
					Title: types.SectionTitle{
						Attributes: types.ElementAttributes{
							types.AttrID:       "title",
							types.AttrCustomID: false,
						},
						Elements: types.InlineElements{
							types.StringElement{
								Content: "title",
							},
						},
					},
					Attributes: types.ElementAttributes{
						types.AttrAuthors: []types.DocumentAuthor{
							{
								FullName: "john doe",
								Email:    "",
							},
						},
						types.AttrRevision: types.DocumentRevision{
							Revnumber: "1.0",
							Revdate:   "June 19, 2017",
							Revremark: "",
						},
					},
					Elements: []interface{}{},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Section0WithMetadata"))
			})
			It("revision with revnumber only - comma suffix", func() {
				actualContent := `= title
				john doe
				1.0,`
				expectedResult := types.Section{
					Level: 0,
					Title: types.SectionTitle{
						Attributes: types.ElementAttributes{
							types.AttrID:       "title",
							types.AttrCustomID: false,
						},
						Elements: types.InlineElements{
							types.StringElement{
								Content: "title",
							},
						},
					},
					Attributes: types.ElementAttributes{
						types.AttrAuthors: []types.DocumentAuthor{
							{
								FullName: "john doe",
								Email:    "",
							},
						},
						types.AttrRevision: types.DocumentRevision{
							Revnumber: "1.0",
							Revdate:   "",
							Revremark: "",
						},
					},
					Elements: []interface{}{},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Section0WithMetadata"))
			})

			It("revision with revdate as number - spaces and no prefix no suffix", func() {
				actualContent := `= title
				john doe
				1.0`
				expectedResult := types.Section{
					Level: 0,
					Title: types.SectionTitle{
						Attributes: types.ElementAttributes{
							types.AttrID:       "title",
							types.AttrCustomID: false,
						},
						Elements: types.InlineElements{
							types.StringElement{
								Content: "title",
							},
						},
					},
					Attributes: types.ElementAttributes{
						types.AttrAuthors: []types.DocumentAuthor{
							{
								FullName: "john doe",
								Email:    "",
							},
						},
						types.AttrRevision: types.DocumentRevision{
							Revnumber: "",
							Revdate:   "1.0",
							Revremark: "",
						},
					},
					Elements: []interface{}{},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Section0WithMetadata"))
			})

			It("revision with revdate as alphanum - spaces and no prefix no suffix", func() {
				actualContent := `= title
				john doe
				1.0a`
				expectedResult := types.Section{
					Level: 0,
					Title: types.SectionTitle{
						Attributes: types.ElementAttributes{
							types.AttrID:       "title",
							types.AttrCustomID: false,
						},
						Elements: types.InlineElements{
							types.StringElement{
								Content: "title",
							},
						},
					},
					Attributes: types.ElementAttributes{
						types.AttrAuthors: []types.DocumentAuthor{
							{
								FullName: "john doe",
								Email:    "",
							},
						},
						types.AttrRevision: types.DocumentRevision{
							Revnumber: "",
							Revdate:   "1.0a",
							Revremark: "",
						},
					},
					Elements: []interface{}{},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Section0WithMetadata"))
			})

			It("revision with revnumber only", func() {
				actualContent := `= title
				john doe
				v1.0:`
				expectedResult := types.Section{
					Level: 0,
					Title: types.SectionTitle{
						Attributes: types.ElementAttributes{
							types.AttrID:       "title",
							types.AttrCustomID: false,
						},
						Elements: types.InlineElements{
							types.StringElement{
								Content: "title",
							},
						},
					},
					Attributes: types.ElementAttributes{
						types.AttrAuthors: []types.DocumentAuthor{
							{
								FullName: "john doe",
								Email:    "",
							},
						},
						types.AttrRevision: types.DocumentRevision{
							Revnumber: "1.0",
							Revdate:   "",
							Revremark: "",
						},
					},
					Elements: []interface{}{},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Section0WithMetadata"))
			})

			It("revision with spaces and capital revnumber ", func() {
				actualContent := `= title
				john doe
				V1.0:`
				expectedResult := types.Section{
					Level: 0,
					Title: types.SectionTitle{
						Attributes: types.ElementAttributes{
							types.AttrID:       "title",
							types.AttrCustomID: false,
						},
						Elements: types.InlineElements{
							types.StringElement{
								Content: "title",
							},
						},
					},
					Attributes: types.ElementAttributes{
						types.AttrAuthors: []types.DocumentAuthor{
							{
								FullName: "john doe",
								Email:    "",
							},
						},
						types.AttrRevision: types.DocumentRevision{
							Revnumber: "1.0",
							Revdate:   "",
							Revremark: "",
						},
					},
					Elements: []interface{}{},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Section0WithMetadata"))
			})

			It("revision only - with comma separator", func() {
				actualContent := `= title
				john doe
				v1.0,`
				expectedResult := types.Section{
					Level: 0,
					Title: types.SectionTitle{
						Attributes: types.ElementAttributes{
							types.AttrID:       "title",
							types.AttrCustomID: false,
						},
						Elements: types.InlineElements{
							types.StringElement{
								Content: "title",
							},
						},
					},
					Attributes: types.ElementAttributes{
						types.AttrAuthors: []types.DocumentAuthor{
							{
								FullName: "john doe",
								Email:    "",
							},
						},
						types.AttrRevision: types.DocumentRevision{
							Revnumber: "1.0",
							Revdate:   "",
							Revremark: "",
						},
					},
					Elements: []interface{}{},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Section0WithMetadata"))
			})

			It("revision with revnumber plus comma and colon separators", func() {
				actualContent := `= title
				john doe
				v1.0,:`
				expectedResult := types.Section{
					Level: 0,
					Title: types.SectionTitle{
						Attributes: types.ElementAttributes{
							types.AttrID:       "title",
							types.AttrCustomID: false,
						},
						Elements: types.InlineElements{
							types.StringElement{
								Content: "title",
							},
						},
					},
					Attributes: types.ElementAttributes{
						types.AttrAuthors: []types.DocumentAuthor{
							{
								FullName: "john doe",
								Email:    "",
							},
						},
						types.AttrRevision: types.DocumentRevision{
							Revnumber: "1.0",
							Revdate:   "",
							Revremark: "",
						},
					},
					Elements: []interface{}{},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Section0WithMetadata"))
			})

			It("revision with revnumber plus colon separator", func() {
				actualContent := `= title
john doe
v1.0:`
				expectedResult := types.Section{
					Level: 0,
					Title: types.SectionTitle{
						Attributes: types.ElementAttributes{
							types.AttrID:       "title",
							types.AttrCustomID: false,
						},
						Elements: types.InlineElements{
							types.StringElement{
								Content: "title",
							},
						},
					},
					Attributes: types.ElementAttributes{
						types.AttrAuthors: []types.DocumentAuthor{
							{
								FullName: "john doe",
								Email:    "",
							},
						},
						types.AttrRevision: types.DocumentRevision{
							Revnumber: "1.0",
							Revdate:   "",
							Revremark: "",
						},
					},
					Elements: []interface{}{},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Section0WithMetadata"))
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
					Attributes:         types.DocumentAttributes{},
					ElementReferences:  types.ElementReferences{},
					Footnotes:          types.Footnotes{},
					FootnoteReferences: types.FootnoteReferences{},
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
:hardbreaks:
a paragraph`
				expectedResult := types.Document{
					Attributes:         types.DocumentAttributes{},
					ElementReferences:  types.ElementReferences{},
					Footnotes:          types.Footnotes{},
					FootnoteReferences: types.FootnoteReferences{},
					Elements: []interface{}{
						types.DocumentAttributeDeclaration{Name: "toc"},
						types.DocumentAttributeDeclaration{Name: "date", Value: "2017-01-01"},
						types.DocumentAttributeDeclaration{Name: "author", Value: "Xavier"},
						types.DocumentAttributeDeclaration{Name: types.DocumentAttrHardBreaks},
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
				verify(GinkgoT(), expectedResult, actualContent)
			})

			It("contiguous attributes and paragraph with blank line in-between", func() {
				actualContent := `:toc:
:date: 2017-01-01
:author: Xavier

a paragraph`
				expectedResult := types.Document{
					Attributes:         types.DocumentAttributes{},
					ElementReferences:  types.ElementReferences{},
					Footnotes:          types.Footnotes{},
					FootnoteReferences: types.FootnoteReferences{},
					Elements: []interface{}{
						types.DocumentAttributeDeclaration{Name: "toc"},
						types.DocumentAttributeDeclaration{Name: "date", Value: "2017-01-01"},
						types.DocumentAttributeDeclaration{Name: "author", Value: "Xavier"},
						types.BlankLine{},
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
				verify(GinkgoT(), expectedResult, actualContent)
			})

			It("splitted attributes and paragraph with blank line in-between", func() {
				actualContent := `:toc:
:date: 2017-01-01

:author: Xavier

:hardbreaks:

a paragraph`
				expectedResult := types.Document{
					Attributes:         types.DocumentAttributes{},
					ElementReferences:  types.ElementReferences{},
					Footnotes:          types.Footnotes{},
					FootnoteReferences: types.FootnoteReferences{},
					Elements: []interface{}{
						types.DocumentAttributeDeclaration{Name: "toc"},
						types.DocumentAttributeDeclaration{Name: "date", Value: "2017-01-01"},
						types.BlankLine{},
						types.DocumentAttributeDeclaration{Name: "author", Value: "Xavier"},
						types.BlankLine{},
						types.DocumentAttributeDeclaration{Name: "hardbreaks"},
						types.BlankLine{},
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
				verify(GinkgoT(), expectedResult, actualContent)
			})

			It("no header and attributes in body", func() {
				actualContent := `a paragraph
	
:toc:
:date: 2017-01-01
:author: Xavier`
				expectedResult := types.Document{
					Attributes:         types.DocumentAttributes{},
					ElementReferences:  types.ElementReferences{},
					Footnotes:          types.Footnotes{},
					FootnoteReferences: types.FootnoteReferences{},
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
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

		Context("document attribute substitutions", func() {

			It("paragraph with attribute substitution", func() {
				actualContent := `:author: Xavier
			
a paragraph written by {author}.`
				expectedResult := types.Document{
					Attributes:         types.DocumentAttributes{},
					ElementReferences:  types.ElementReferences{},
					Footnotes:          types.Footnotes{},
					FootnoteReferences: types.FootnoteReferences{},
					Elements: []interface{}{
						types.DocumentAttributeDeclaration{Name: "author", Value: "Xavier"},
						types.BlankLine{},
						types.Paragraph{
							Attributes: types.ElementAttributes{},
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
					Attributes:         types.DocumentAttributes{},
					ElementReferences:  types.ElementReferences{},
					Footnotes:          types.Footnotes{},
					FootnoteReferences: types.FootnoteReferences{},
					Elements: []interface{}{
						types.DocumentAttributeDeclaration{Name: "author", Value: "Xavier"},
						types.BlankLine{},
						types.DocumentAttributeReset{Name: "author1"},
						types.DocumentAttributeReset{Name: "author2"},
						types.Paragraph{
							Attributes: types.ElementAttributes{},
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
			title := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID:       "the_dangerous_and_thrilling_documentation_chronicles",
					types.AttrCustomID: false,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "The Dangerous and Thrilling Documentation Chronicles"},
				},
			}
			expectedResult := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"the_dangerous_and_thrilling_documentation_chronicles": title,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Level: 0,
						Title: title,
						Attributes: types.ElementAttributes{
							types.AttrAuthors: []types.DocumentAuthor{
								{
									FullName: "Kismet Rainbow Chameleon ",
									Email:    "kismet@asciidoctor.org",
								},
								{
									FullName: "Lazarus het_Draeke ",
									Email:    "lazarus@asciidoctor.org",
								},
							},
							types.AttrRevision: types.DocumentRevision{
								Revnumber: "1.0",
								Revdate:   "June 19, 2017",
								Revremark: "First incarnation",
							},
						},
						Elements: []interface{}{
							types.DocumentAttributeDeclaration{
								Name:  "toc",
								Value: "",
							},
							types.DocumentAttributeDeclaration{
								Name:  "keywords",
								Value: "documentation, team, obstacles, journey, victory",
							},
							types.BlankLine{},
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "This journey begins on a bleary Monday morning."},
									},
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

			title := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID:       "a_header",
					types.AttrCustomID: false,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "a header"},
				},
			}
			section1Title := types.SectionTitle{
				Attributes: types.ElementAttributes{
					types.AttrID:       "section_1",
					types.AttrCustomID: false,
				},
				Elements: types.InlineElements{
					types.StringElement{Content: "section 1"},
				},
			}
			expectedResult := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"a_header":  title,
					"section_1": section1Title,
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Level:      0,
						Title:      title,
						Attributes: types.ElementAttributes{},
						Elements: []interface{}{
							types.Section{
								Level:      1,
								Title:      section1Title,
								Attributes: types.ElementAttributes{},
								Elements: []interface{}{
									types.Paragraph{
										Attributes: types.ElementAttributes{},
										Lines: []types.InlineElements{
											{
												types.StringElement{Content: "a paragraph with "},
												types.QuotedText{
													Kind: types.Bold,
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
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
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
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
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
