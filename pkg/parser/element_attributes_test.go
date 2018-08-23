package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("element attributes", func() {

	Context("element link", func() {

		Context("valid syntax", func() {
			It("element link alone", func() {
				actualContent := `[link=http://foo.bar]
a paragraph`
				expectedResult := types.Paragraph{
					Attributes: types.ElementAttributes{
						"link": "http://foo.bar",
					},
					Lines: []types.InlineElements{
						{
							types.StringElement{
								Content: "a paragraph",
							},
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})
			It("spaces in link", func() {
				actualContent := `[link= http://foo.bar  ]
a paragraph`
				expectedResult := types.Paragraph{
					Attributes: types.ElementAttributes{
						"link": "http://foo.bar",
					},
					Lines: []types.InlineElements{
						{
							types.StringElement{
								Content: "a paragraph",
							},
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})
		})

		Context("invalid syntax", func() {
			It("spaces before keyword", func() {
				actualContent := `[ link=http://foo.bar]
a paragraph`
				expectedResult := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{
								Content: "[ link=http://foo.bar]",
							},
						},
						{
							types.StringElement{
								Content: "a paragraph",
							},
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("unbalanced brackets", func() {
				actualContent := `[link=http://foo.bar
a paragraph`
				expectedResult := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{
								Content: "[link=http://foo.bar",
							},
						},
						{
							types.StringElement{
								Content: "a paragraph",
							},
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})
		})
	})

	Context("element id", func() {

		Context("valid syntax", func() {

			It("normal syntax", func() {
				actualContent := `[[img-foobar]]
a paragraph`
				expectedResult := types.Paragraph{
					Attributes: types.ElementAttributes{
						types.AttrID: "img-foobar",
					},
					Lines: []types.InlineElements{
						{
							types.StringElement{
								Content: "a paragraph",
							},
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("short-hand syntax", func() {
				actualContent := `[#img-foobar]
a paragraph`
				expectedResult := types.Paragraph{
					Attributes: types.ElementAttributes{
						types.AttrID: "img-foobar",
					},
					Lines: []types.InlineElements{
						{
							types.StringElement{
								Content: "a paragraph",
							},
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})
		})

		Context("invalid syntax", func() {

			It("extra spaces", func() {
				actualContent := `[ #img-foobar ]
a paragraph`
				expectedResult := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{
								Content: "[ #img-foobar ]",
							},
						},
						{
							types.StringElement{
								Content: "a paragraph",
							},
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("unbalanced brackets", func() {
				actualContent := `[#img-foobar
a paragraph`
				expectedResult := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{
								Content: "[#img-foobar",
							},
						},
						{
							types.StringElement{
								Content: "a paragraph",
							},
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})
		})
	})

	Context("element title", func() {

		Context("valid syntax", func() {

			It("valid element title", func() {
				actualContent := `.a title
a paragraph`
				expectedResult := types.Paragraph{
					Attributes: types.ElementAttributes{
						types.AttrTitle: "a title",
					},
					Lines: []types.InlineElements{
						{
							types.StringElement{
								Content: "a paragraph",
							},
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})
		})

		Context("invalid syntax", func() {
			It("extra space after dot", func() {

				actualContent := `. a title
a list item!`
				expectedResult := types.OrderedList{
					Attributes: types.ElementAttributes{},
					Items: []types.OrderedListItem{
						{
							Attributes:     map[string]interface{}{},
							Level:          1,
							Position:       1,
							NumberingStyle: types.Arabic,
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: []types.InlineElements{
										{
											types.StringElement{
												Content: "a title",
											},
										},
										{
											types.StringElement{
												Content: "a list item!",
											},
										},
									},
								},
							},
						},
					},
				}

				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("not a dot", func() {
				actualContent := `!a title
a paragraph`

				expectedResult := types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							types.StringElement{
								Content: "!a title",
							},
						},
						{
							types.StringElement{
								Content: "a paragraph",
							},
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})
		})
	})

	Context("element role", func() {

		Context("valid syntax", func() {

			It("shortcut role element", func() {
				actualContent := `[.a role]
a paragraph`
				expectedResult := types.Paragraph{
					Attributes: types.ElementAttributes{
						types.AttrRole: "a role",
					},
					Lines: []types.InlineElements{
						{
							types.StringElement{
								Content: "a paragraph",
							},
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})

			It("full role syntax", func() {
				actualContent := `[role=a role]
a paragraph`
				expectedResult := types.Paragraph{
					Attributes: types.ElementAttributes{
						types.AttrRole: "a role",
					},
					Lines: []types.InlineElements{
						{
							types.StringElement{
								Content: "a paragraph",
							},
						},
					},
				}
				verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
			})
		})
	})

})
