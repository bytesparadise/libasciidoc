package parser_test

import (
	"github.com/bytesparadise/libasciidoc/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Parsing Document Attributes", func() {

	Context("Valid document attributes", func() {

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
							&types.DocumentAttribute{Name: "toc"},
							&types.DocumentAttribute{Name: "date", Value: "2017-01-01"},
							&types.DocumentAttribute{Name: "author", Value: "Xavier"},
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
					&types.DocumentAttribute{Name: "toc"},
					&types.DocumentAttribute{Name: "date", Value: "2017-01-01"},
					&types.DocumentAttribute{Name: "author", Value: "Xavier"},
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
					&types.DocumentAttribute{Name: "toc"},
					&types.DocumentAttribute{Name: "date", Value: "2017-01-01"},
					&types.DocumentAttribute{Name: "author", Value: "Xavier"},
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
					&types.DocumentAttribute{Name: "toc"},
					&types.DocumentAttribute{Name: "date", Value: "2017-01-01"},
					&types.DocumentAttribute{Name: "author", Value: "Xavier"},
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
					&types.DocumentAttribute{Name: "toc"},
					&types.DocumentAttribute{Name: "date", Value: "2017-01-01"},
					&types.DocumentAttribute{Name: "author", Value: "Xavier"},
				},
			}
			verify(GinkgoT(), expectedDocument, actualContent)
		})
	})

	Context("Valid document attributes", func() {
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
	})
})
