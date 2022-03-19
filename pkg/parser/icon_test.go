package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"
	. "github.com/onsi/ginkgo/v2" // nolint:golint
	. "github.com/onsi/gomega"    // nolint:golint
)

var _ = Describe("icons", func() {

	Context("inline icons", func() {

		Context("inline elements", func() {

			It("inline icon with empty alt only", func() {
				source := "icon:tip[]"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.Icon{
									Class: "tip",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("inline icon with empty alt and trailing spaces", func() {
				source := "icon:note[]  \t\t  "
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								// suffix spaces are trimmed on each line
								&types.Icon{
									Class: "note",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("inline icon with empty alt surrounded by text", func() {
				source := "beware icon:caution[] of tigers"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "beware ",
								},
								&types.Icon{
									Class: "caution",
								},
								&types.StringElement{
									Content: " of tigers",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("inline icon with size alone", func() {
				source := "icon:caution[2x]"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.Icon{
									Class:      "caution",
									Attributes: types.Attributes{types.AttrIconSize: "2x"},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("inline icon with other attribute (title)", func() {
				source := "icon:caution[title=\"bogus\"]"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.Icon{
									Class:      "caution",
									Attributes: types.Attributes{types.AttrTitle: "bogus"},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("inline icon with anchor attribute", func() {
				source := "icon:caution[id=anchor]"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.Icon{
									Class: "caution",
									Attributes: types.Attributes{
										types.AttrID: "anchor",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("inline icon with multiple other attributes", func() {
				source := `icon:caution[id=anchor,title="White Fang"]`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.Icon{
									Class: "caution",
									Attributes: types.Attributes{
										types.AttrID:    "anchor",
										types.AttrTitle: "White Fang",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("inline icon with size and multiple other attributes", func() {
				source := "icon:caution[fw,id=anchor,title=\"White Fang\"]"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.Icon{
									Class: "caution",
									Attributes: types.Attributes{
										types.AttrID:       "anchor",
										types.AttrTitle:    "White Fang",
										types.AttrIconSize: "fw",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("inline icon with space after colon", func() {
				source := "here is my icon: icon:info[]"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "here is my icon: ",
								},
								&types.Icon{
									Class: "info",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("inline icon in title works", func() {
				source := `== a icon:note[] from me`
				title := []interface{}{
					&types.StringElement{
						Content: "a ",
					},
					&types.Icon{
						Class: "note",
					},
					&types.StringElement{
						Content: " from me",
					},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_a_note_from_me",
							},
							Level: 1,
							Title: title,
						},
					},
					ElementReferences: types.ElementReferences{
						"_a_note_from_me": title,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_a_note_from_me",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("inline icon at title start", func() {
				source := `= icon:warning[] or what icon:note[] to do`
				title := []interface{}{
					&types.Icon{Class: "warning"},
					&types.StringElement{Content: " or what "},
					&types.Icon{Class: "note"},
					&types.StringElement{Content: " to do"},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Title: title,
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			// Note that the parsing that occurs here does not include the re-parse of the list item term.
			// That is done in second pass.
			It("inline icon as labeled list item description", func() {
				source := `discount:: icon:tags[alt="Discount"] Cheap cheap!
item 2:: two`
				expected := &types.Document{
					Elements: []interface{}{
						&types.List{
							Kind: types.LabeledListKind,
							Elements: []types.ListElement{
								&types.LabeledListElement{
									Style: "::",
									Term: []interface{}{
										&types.StringElement{Content: "discount"},
									},
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.Icon{Class: "tags", Attributes: types.Attributes{types.AttrImageAlt: "Discount"}},
												&types.StringElement{Content: " Cheap cheap!"},
											},
										},
									},
								},
								&types.LabeledListElement{
									Style: "::",
									Term: []interface{}{
										&types.StringElement{Content: "item 2"},
									},
									Elements: []interface{}{
										&types.Paragraph{
											Elements: []interface{}{
												&types.StringElement{Content: "two"},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))

			})

			It("inline icon in quoted text", func() {
				source := `an _italicized icon:warning[] message_`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "an ",
								},
								&types.QuotedText{
									Kind: types.SingleQuoteItalic,
									Elements: []interface{}{
										&types.StringElement{Content: "italicized "},
										&types.Icon{Class: "warning"},
										&types.StringElement{Content: " message"},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
			It("inline icon in marked text", func() {
				source := `#marked icon:warning[] message#`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.QuotedText{
									Kind: types.SingleQuoteMarked,
									Elements: []interface{}{
										&types.StringElement{Content: "marked "},
										&types.Icon{Class: "warning"},
										&types.StringElement{Content: " message"},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
			It("inline icon in bold text", func() {
				source := `in *bold icon:warning[] message*`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "in "},
								&types.QuotedText{
									Kind: types.SingleQuoteBold,
									Elements: []interface{}{
										&types.StringElement{Content: "bold "},
										&types.Icon{Class: "warning"},
										&types.StringElement{Content: " message"},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
			It("inline icon in monospace text", func() {
				source := "in `monospace icon:warning[] message`"
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "in "},
								&types.QuotedText{
									Kind: types.SingleQuoteMonospace,
									Elements: []interface{}{
										&types.StringElement{Content: "monospace "},
										&types.Icon{Class: "warning"},
										&types.StringElement{Content: " message"},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})
	})
})
