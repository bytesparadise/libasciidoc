package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("cross-references", func() {

	Context("section reference", func() {

		It("cross-reference with custom id", func() {
			actualContent := `[[thetitle]]
== a title

with some content linked to <<thetitle>>!`
			expectedResult := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"thetitle": types.SectionTitle{
						Attributes: types.ElementAttributes{
							types.AttrID:       "thetitle",
							types.AttrCustomID: true,
						},
						Elements: types.InlineElements{
							types.StringElement{
								Content: "a title",
							},
						},
					},
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Level: 1,
						Title: types.SectionTitle{
							Attributes: types.ElementAttributes{
								types.AttrID:       "thetitle",
								types.AttrCustomID: true,
							},
							Elements: types.InlineElements{
								types.StringElement{
									Content: "a title",
								},
							},
						},
						Attributes: types.ElementAttributes{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "with some content linked to ",
										},
										types.CrossReference{
											ID:    "thetitle",
											Label: "",
										},
										types.StringElement{
											Content: "!",
										},
									},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Document"))
		})

		It("cross-reference with custom id and label", func() {
			actualContent := `[[thetitle]]
== a title

with some content linked to <<thetitle,a label to the title>>!`
			expectedResult := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"thetitle": types.SectionTitle{
						Attributes: types.ElementAttributes{
							types.AttrID:       "thetitle",
							types.AttrCustomID: true,
						},
						Elements: types.InlineElements{
							types.StringElement{
								Content: "a title",
							},
						},
					},
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Level: 1,
						Title: types.SectionTitle{
							Attributes: types.ElementAttributes{
								types.AttrID:       "thetitle",
								types.AttrCustomID: true,
							},
							Elements: types.InlineElements{
								types.StringElement{
									Content: "a title",
								},
							},
						},
						Attributes: types.ElementAttributes{},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{
											Content: "with some content linked to ",
										},
										types.CrossReference{
											ID:    "thetitle",
											Label: "a label to the title",
										},
										types.StringElement{
											Content: "!",
										},
									},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Document"))
		})
	})
})
