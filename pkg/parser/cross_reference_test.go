package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("cross-references - preflight", func() {

	Context("section reference", func() {

		It("cross-reference with custom id alone", func() {
			source := `[[thetitle]]
== a title

with some content linked to <<thetitle>>!`
			expected := &types.PreflightDocument{
				Blocks: []interface{}{
					&types.Section{
						Level: 1,
						Attributes: types.ElementAttributes{
							types.AttrID:       "thetitle",
							types.AttrCustomID: true,
						},
						Title: types.InlineElements{
							&types.StringElement{
								Content: "a title",
							},
						},
						Elements: []interface{}{},
					},
					&types.BlankLine{},
					&types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								&types.StringElement{
									Content: "with some content linked to ",
								},
								&types.CrossReference{
									ID:    "thetitle",
									Label: "",
								},
								&types.StringElement{
									Content: "!",
								},
							},
						},
					},
				},
			}
			verifyPreflight(expected, source)
		})

		It("cross-reference with custom id and label", func() {
			source := `[[thetitle]]
== a title

with some content linked to <<thetitle,a label to the title>>!`
			expected := &types.PreflightDocument{
				Blocks: []interface{}{
					&types.Section{
						Level: 1,
						Attributes: types.ElementAttributes{
							types.AttrID:       "thetitle",
							types.AttrCustomID: true,
						},
						Title: types.InlineElements{
							&types.StringElement{
								Content: "a title",
							},
						},
						Elements: []interface{}{},
					},
					&types.BlankLine{},
					&types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: []types.InlineElements{
							{
								&types.StringElement{
									Content: "with some content linked to ",
								},
								&types.CrossReference{
									ID:    "thetitle",
									Label: "a label to the title",
								},
								&types.StringElement{
									Content: "!",
								},
							},
						},
					},
				},
			}
			verifyPreflight(expected, source)
		})
	})
})

var _ = Describe("cross-references - document", func() {

	Context("section reference", func() {

		It("cross-reference with custom id alone", func() {
			source := `[[thetitle]]
== a title

with some content linked to <<thetitle>>!`
			expected := &types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"thetitle": types.InlineElements{
						&types.StringElement{
							Content: "a title",
						},
					},
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					&types.Section{
						Level: 1,
						Attributes: types.ElementAttributes{
							types.AttrID:       "thetitle",
							types.AttrCustomID: true,
						},
						Title: types.InlineElements{
							&types.StringElement{
								Content: "a title",
							},
						},
						Elements: []interface{}{
							&types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										&types.StringElement{
											Content: "with some content linked to ",
										},
										&types.CrossReference{
											ID:    "thetitle",
											Label: "",
										},
										&types.StringElement{
											Content: "!",
										},
									},
								},
							},
						},
					},
				},
			}
			verifyDocument(expected, source)
		})

		It("cross-reference with custom id and label", func() {
			source := `[[thetitle]]
== a title

with some content linked to <<thetitle,a label to the title>>!`
			expected := &types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: types.ElementReferences{
					"thetitle": types.InlineElements{
						&types.StringElement{
							Content: "a title",
						},
					},
				},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					&types.Section{
						Level: 1,
						Attributes: types.ElementAttributes{
							types.AttrID:       "thetitle",
							types.AttrCustomID: true,
						},
						Title: types.InlineElements{
							&types.StringElement{
								Content: "a title",
							},
						},
						Elements: []interface{}{
							&types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										&types.StringElement{
											Content: "with some content linked to ",
										},
										&types.CrossReference{
											ID:    "thetitle",
											Label: "a label to the title",
										},
										&types.StringElement{
											Content: "!",
										},
									},
								},
							},
						},
					},
				},
			}
			verifyDocument(expected, source)
		})
	})
})
