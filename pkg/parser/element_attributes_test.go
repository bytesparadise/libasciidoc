package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("element attributes - preflight", func() {

	Context("element link", func() {

		Context("valid syntax", func() {
			It("element link alone", func() {
				source := `[link=http://foo.bar]
a paragraph`
				expected := &types.Paragraph{
					Attributes: types.ElementAttributes{
						"link": "http://foo.bar",
					},
					Lines: []types.InlineElements{
						{
							&types.StringElement{
								Content: "a paragraph",
							},
						},
					},
				}
				verifyDocumentBlock(expected, source)
			})
			It("spaces in link", func() {
				source := `[link= http://foo.bar  ]
a paragraph`
				expected := &types.Paragraph{
					Attributes: types.ElementAttributes{
						"link": "http://foo.bar",
					},
					Lines: []types.InlineElements{
						{
							&types.StringElement{
								Content: "a paragraph",
							},
						},
					},
				}
				verifyDocumentBlock(expected, source)
			})
		})

		Context("invalid syntax", func() {
			It("spaces before keyword", func() {
				source := `[ link=http://foo.bar]
a paragraph`
				expected := &types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							&types.StringElement{
								Content: "[ link=http://foo.bar]",
							},
						},
						{
							&types.StringElement{
								Content: "a paragraph",
							},
						},
					},
				}
				verifyDocumentBlock(expected, source)
			})

			It("unbalanced brackets", func() {
				source := `[link=http://foo.bar
a paragraph`
				expected := &types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							&types.StringElement{
								Content: "[link=http://foo.bar",
							},
						},
						{
							&types.StringElement{
								Content: "a paragraph",
							},
						},
					},
				}
				verifyDocumentBlock(expected, source)
			})
		})
	})

	Context("element id", func() {

		Context("valid syntax", func() {

			It("normal syntax", func() {
				source := `[[img-foobar]]
a paragraph`
				expected := &types.Paragraph{
					Attributes: types.ElementAttributes{
						types.AttrID:       "img-foobar",
						types.AttrCustomID: true,
					},
					Lines: []types.InlineElements{
						{
							&types.StringElement{
								Content: "a paragraph",
							},
						},
					},
				}
				verifyDocumentBlock(expected, source)
			})

			It("short-hand syntax", func() {
				source := `[#img-foobar]
a paragraph`
				expected := &types.Paragraph{
					Attributes: types.ElementAttributes{
						types.AttrID:       "img-foobar",
						types.AttrCustomID: true,
					},
					Lines: []types.InlineElements{
						{
							&types.StringElement{
								Content: "a paragraph",
							},
						},
					},
				}
				verifyDocumentBlock(expected, source)
			})
		})

		Context("invalid syntax", func() {

			It("extra spaces", func() {
				source := `[ #img-foobar ]
a paragraph`
				expected := &types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							&types.StringElement{
								Content: "[ #img-foobar ]",
							},
						},
						{
							&types.StringElement{
								Content: "a paragraph",
							},
						},
					},
				}
				verifyDocumentBlock(expected, source)
			})

			It("unbalanced brackets", func() {
				source := `[#img-foobar
a paragraph`
				expected := &types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							&types.StringElement{
								Content: "[#img-foobar",
							},
						},
						{
							&types.StringElement{
								Content: "a paragraph",
							},
						},
					},
				}
				verifyDocumentBlock(expected, source)
			})
		})
	})

	Context("element title", func() {

		Context("valid syntax", func() {

			It("valid element title", func() {
				source := `.a title
a paragraph`
				expected := &types.Paragraph{
					Attributes: types.ElementAttributes{
						types.AttrTitle: "a title",
					},
					Lines: []types.InlineElements{
						{
							&types.StringElement{
								Content: "a paragraph",
							},
						},
					},
				}
				verifyDocumentBlock(expected, source)
			})
		})

		Context("invalid syntax", func() {

			It("extra space after dot", func() {
				source := `. a title
a list item!`
				expected := &types.OrderedListItem{
					Attributes:     map[string]interface{}{},
					Level:          1,
					Position:       1,
					NumberingStyle: types.Arabic,
					Elements: []interface{}{
						&types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: []types.InlineElements{
								{
									&types.StringElement{
										Content: "a title",
									},
								},
								{
									&types.StringElement{
										Content: "a list item!",
									},
								},
							},
						},
					},
				}
				verifyDocumentBlock(expected, source)
			})

			It("not a dot", func() {
				source := `!a title
a paragraph`

				expected := &types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: []types.InlineElements{
						{
							&types.StringElement{
								Content: "!a title",
							},
						},
						{
							&types.StringElement{
								Content: "a paragraph",
							},
						},
					},
				}
				verifyDocumentBlock(expected, source)
			})
		})
	})

	Context("element role", func() {

		Context("valid syntax", func() {

			It("shortcut role element", func() {
				source := `[.a role]
a paragraph`
				expected := &types.Paragraph{
					Attributes: types.ElementAttributes{
						types.AttrRole: "a role",
					},
					Lines: []types.InlineElements{
						{
							&types.StringElement{
								Content: "a paragraph",
							},
						},
					},
				}
				verifyDocumentBlock(expected, source)
			})

			It("full role syntax", func() {
				source := `[role=a role]
a paragraph`
				expected := &types.Paragraph{
					Attributes: types.ElementAttributes{
						types.AttrRole: "a role",
					},
					Lines: []types.InlineElements{
						{
							&types.StringElement{
								Content: "a paragraph",
							},
						},
					},
				}
				verifyDocumentBlock(expected, source)
			})
		})
	})

})
