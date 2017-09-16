package parser_test

import (
	"github.com/bytesparadise/libasciidoc/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Parsing Document Attributes", func() {

	Context("Valid document attributes", func() {

		It("valid attribute names", func() {

			actualContent := `:a:
:author: Xavier
:_author: Xavier
:Author: Xavier
:0Author: Xavier
:Auth0r: Xavier`
			expectedDocument := &types.Document{
				Attributes: &types.DocumentAttributes{},
				Elements: []types.DocElement{
					&types.DocumentAttributeDeclaration{Name: "a"},
					&types.DocumentAttributeDeclaration{Name: "author", Value: "Xavier"},
					&types.DocumentAttributeDeclaration{Name: "_author", Value: "Xavier"},
					&types.DocumentAttributeDeclaration{Name: "Author", Value: "Xavier"},
					&types.DocumentAttributeDeclaration{Name: "0Author", Value: "Xavier"},
					&types.DocumentAttributeDeclaration{Name: "Auth0r", Value: "Xavier"},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("heading section with attributes", func() {

			actualContent := `= a heading
:toc:
:date: 2017-01-01
:author: Xavier

a paragraph`
			expectedDocument := &types.Document{
				Attributes: &types.DocumentAttributes{
					"title": "a heading",
				},
				Elements: []types.DocElement{
					&types.Section{
						Heading: types.Heading{
							Level: 1,
							Content: &types.InlineContent{
								Elements: []types.InlineElement{
									&types.StringElement{Content: "a heading"},
								},
							},
							ID: &types.ElementID{
								Value: "_a_heading",
							},
						},
						Elements: []types.DocElement{
							&types.DocumentAttributeDeclaration{Name: "toc"},
							&types.DocumentAttributeDeclaration{Name: "date", Value: "2017-01-01"},
							&types.DocumentAttributeDeclaration{Name: "author", Value: "Xavier"},
							&types.Paragraph{
								Lines: []*types.InlineContent{
									&types.InlineContent{
										Elements: []types.InlineElement{
											&types.StringElement{Content: "a paragraph"},
										},
									},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("attributes and paragraph without blank line in-between", func() {

			actualContent := `:toc:
:date:  2017-01-01
:author: Xavier
a paragraph`
			expectedDocument := &types.Document{
				Attributes: &types.DocumentAttributes{},
				Elements: []types.DocElement{
					&types.DocumentAttributeDeclaration{Name: "toc"},
					&types.DocumentAttributeDeclaration{Name: "date", Value: "2017-01-01"},
					&types.DocumentAttributeDeclaration{Name: "author", Value: "Xavier"},
					&types.Paragraph{
						Lines: []*types.InlineContent{
							&types.InlineContent{
								Elements: []types.InlineElement{
									&types.StringElement{Content: "a paragraph"},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("contiguous attributes and paragraph with blank line in-between", func() {

			actualContent := `:toc:
:date: 2017-01-01
:author: Xavier

a paragraph`
			expectedDocument := &types.Document{
				Attributes: &types.DocumentAttributes{},
				Elements: []types.DocElement{
					&types.DocumentAttributeDeclaration{Name: "toc"},
					&types.DocumentAttributeDeclaration{Name: "date", Value: "2017-01-01"},
					&types.DocumentAttributeDeclaration{Name: "author", Value: "Xavier"},
					&types.Paragraph{
						Lines: []*types.InlineContent{
							&types.InlineContent{
								Elements: []types.InlineElement{
									&types.StringElement{Content: "a paragraph"},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("splitted attributes and paragraph with blank line in-between", func() {

			actualContent := `:toc:
:date: 2017-01-01

:author: Xavier

a paragraph`
			expectedDocument := &types.Document{
				Attributes: &types.DocumentAttributes{},
				Elements: []types.DocElement{
					&types.DocumentAttributeDeclaration{Name: "toc"},
					&types.DocumentAttributeDeclaration{Name: "date", Value: "2017-01-01"},
					&types.DocumentAttributeDeclaration{Name: "author", Value: "Xavier"},
					&types.Paragraph{
						Lines: []*types.InlineContent{
							&types.InlineContent{
								Elements: []types.InlineElement{
									&types.StringElement{Content: "a paragraph"},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("no heading and attributes in body", func() {

			actualContent := `a paragraph
		
:toc:
:date: 2017-01-01
:author: Xavier`
			expectedDocument := &types.Document{
				Attributes: &types.DocumentAttributes{},
				Elements: []types.DocElement{
					&types.Paragraph{
						Lines: []*types.InlineContent{
							&types.InlineContent{
								Elements: []types.InlineElement{
									&types.StringElement{Content: "a paragraph"},
								},
							},
						},
					},
					&types.DocumentAttributeDeclaration{Name: "toc"},
					&types.DocumentAttributeDeclaration{Name: "date", Value: "2017-01-01"},
					&types.DocumentAttributeDeclaration{Name: "author", Value: "Xavier"},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("paragraph with attribute substitution", func() {

			actualContent := `:author: Xavier
			
a paragraph written by {author}.`
			expectedDocument := &types.Document{
				Attributes: &types.DocumentAttributes{},
				Elements: []types.DocElement{
					&types.DocumentAttributeDeclaration{Name: "author", Value: "Xavier"},
					&types.Paragraph{
						Lines: []*types.InlineContent{
							&types.InlineContent{
								Elements: []types.InlineElement{
									&types.StringElement{Content: "a paragraph written by "},
									&types.DocumentAttributeSubstitution{Name: "author"},
									&types.StringElement{Content: "."},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("paragraph with attribute resets", func() {

			actualContent := `:author: Xavier
			
:!author1:
:author2!:
a paragraph written by {author}.`
			expectedDocument := &types.Document{
				Attributes: &types.DocumentAttributes{},
				Elements: []types.DocElement{
					&types.DocumentAttributeDeclaration{Name: "author", Value: "Xavier"},
					&types.DocumentAttributeReset{Name: "author1"},
					&types.DocumentAttributeReset{Name: "author2"},
					&types.Paragraph{
						Lines: []*types.InlineContent{
							&types.InlineContent{
								Elements: []types.InlineElement{
									&types.StringElement{Content: "a paragraph written by "},
									&types.DocumentAttributeSubstitution{Name: "author"},
									&types.StringElement{Content: "."},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})
	})

	Context("Invalid document attributes", func() {
		It("paragraph and without blank line in between", func() {

			actualContent := `a paragraph
:toc:
:date: 2017-01-01
:author: Xavier`
			expectedDocument := &types.Document{
				Attributes: &types.DocumentAttributes{},
				Elements: []types.DocElement{
					&types.Paragraph{
						Lines: []*types.InlineContent{
							&types.InlineContent{
								Elements: []types.InlineElement{
									&types.StringElement{Content: "a paragraph"},
								},
							},
							&types.InlineContent{
								Elements: []types.InlineElement{
									&types.StringElement{Content: ":toc:"},
								},
							},
							&types.InlineContent{
								Elements: []types.InlineElement{
									&types.StringElement{Content: ":date: 2017-01-01"},
								},
							},
							&types.InlineContent{
								Elements: []types.InlineElement{
									&types.StringElement{Content: ":author: Xavier"},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})

		It("invalid attribute names", func() {

			actualContent := `:@date: 2017-01-01
:{author}: Xavier`
			expectedDocument := &types.Document{
				Attributes: &types.DocumentAttributes{},
				Elements: []types.DocElement{
					&types.Paragraph{
						Lines: []*types.InlineContent{
							&types.InlineContent{
								Elements: []types.InlineElement{
									&types.StringElement{Content: ":@date: 2017-01-01"},
								},
							},
							&types.InlineContent{
								Elements: []types.InlineElement{
									&types.StringElement{Content: ":{author}: Xavier"},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})
	})
})
