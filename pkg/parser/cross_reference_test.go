package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("cross-references - draft", func() {

	Context("section reference", func() {

		It("cross-reference with custom id alone", func() {
			source := `[[thetitle]]
== a title

with some content linked to <<thetitle>>!`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Section{
						Level: 1,
						Attributes: types.ElementAttributes{
							types.AttrID:       "thetitle",
							types.AttrCustomID: true,
						},
						Title: []interface{}{
							types.StringElement{
								Content: "a title",
							},
						},
						Elements: []interface{}{},
					},
					types.BlankLine{},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
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
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})

		It("cross-reference with custom id and label", func() {
			source := `[[thetitle]]
== a title

with some content linked to <<thetitle,a label to the title>>!`
			expected := types.DraftDocument{
				Blocks: []interface{}{
					types.Section{
						Level: 1,
						Attributes: types.ElementAttributes{
							types.AttrID:       "thetitle",
							types.AttrCustomID: true,
						},
						Title: []interface{}{
							types.StringElement{
								Content: "a title",
							},
						},
						Elements: []interface{}{},
					},
					types.BlankLine{},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
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
			}
			Expect(source).To(BecomeDraftDocument(expected))
		})
	})
})

var _ = Describe("cross-references - document", func() {

	Context("section reference", func() {

		It("cross-reference with custom id alone", func() {
			source := `[[thetitle]]
== a title

with some content linked to <<thetitle>>!`
			expected := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"thetitle": []interface{}{
						types.StringElement{
							Content: "a title",
						},
					},
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Level: 1,
						Attributes: types.ElementAttributes{
							types.AttrID:       "thetitle",
							types.AttrCustomID: true,
						},
						Title: []interface{}{
							types.StringElement{
								Content: "a title",
							},
						},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
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
			Expect(source).To(BecomeDocument(expected))
		})

		It("cross-reference with custom id and label", func() {
			source := `[[thetitle]]
== a title

with some content linked to <<thetitle,a label to the title>>!`
			expected := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"thetitle": []interface{}{
						types.StringElement{
							Content: "a title",
						},
					},
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Level: 1,
						Attributes: types.ElementAttributes{
							types.AttrID:       "thetitle",
							types.AttrCustomID: true,
						},
						Title: []interface{}{
							types.StringElement{
								Content: "a title",
							},
						},
						Elements: []interface{}{
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: [][]interface{}{
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
			Expect(source).To(BecomeDocument(expected))
		})
	})
})
