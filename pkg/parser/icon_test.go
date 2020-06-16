package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"
	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("icons", func() {

	Context("inline icons", func() {

		Context("draft document", func() {

			It("inline icon with empty alt only", func() {
				source := "icon:tip[]"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.Icon{
										Class: "tip",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("inline icon with empty alt and trailing spaces", func() {
				source := "icon:note[]  \t\t  "
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.Icon{
										Class: "note",
									},
									types.StringElement{
										Content: "  \t\t  ",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("inline icon with empty alt surrounded by text", func() {
				source := "beware icon:caution[] of tigers"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "beware ",
									},
									types.Icon{
										Class: "caution",
									},
									types.StringElement{
										Content: " of tigers",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("inline icon with size alone", func() {
				source := "icon:caution[2x]"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.Icon{
										Class:      "caution",
										Attributes: types.Attributes{types.AttrIconSize: "2x"},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("inline icon with other attribute (title)", func() {
				source := "icon:caution[title=\"bogus\"]"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.Icon{
										Class:      "caution",
										Attributes: types.Attributes{types.AttrImageTitle: "bogus"},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("inline icon with anchor attribute", func() {
				source := "icon:caution[id=anchor]"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.Icon{
										Class: "caution",
										Attributes: types.Attributes{
											types.AttrID:       "anchor",
											types.AttrCustomID: true,
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("inline icon with multiple other attributes", func() {
				source := "icon:caution[id=anchor,title=\"White Fang\"]"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.Icon{
										Class: "caution",
										Attributes: types.Attributes{
											types.AttrID:         "anchor",
											types.AttrCustomID:   true,
											types.AttrImageTitle: "White Fang",
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("inline icon with size and multiple other attributes", func() {
				source := "icon:caution[fw,id=anchor,title=\"White Fang\"]"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.Icon{
										Class: "caution",
										Attributes: types.Attributes{
											types.AttrID:         "anchor",
											types.AttrCustomID:   true,
											types.AttrImageTitle: "White Fang",
											types.AttrIconSize:   "fw",
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("inline icon with space after colon", func() {
				source := "here is my icon: icon:info[]"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "here is my icon: ",
									},
									types.Icon{
										Class: "info",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("inline icon in title works", func() {
				source := `== a icon:note[] from me`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Section{
							Level: 1,
							Title: []interface{}{
								types.StringElement{
									Content: "a ",
								},
								types.Icon{
									Class: "note",
								},
								types.StringElement{
									Content: " from me",
								},
							},
							Elements: []interface{}{},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("inline icon at title start", func() {
				source := `= icon:warning[] or what icon:note[] to do`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Section{
							Level: 0,
							Title: []interface{}{
								types.Icon{Class: "warning"},
								types.StringElement{Content: " or what "},
								types.Icon{Class: "note"},
								types.StringElement{Content: " to do"},
							},
							Elements: []interface{}{},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			// NB: The existing grammar for labeled list items does not support any markup
			// in the term text.
			It("inline icon as labeled list item description", func() {
				source := `discount:: icon:tags[alt="Discount"] Cheap cheap!
item 2:: two`
				expected := types.DraftDocument{
					Blocks: []interface{}{

						types.LabeledListItem{
							Level: 1,
							Term: []interface{}{
								types.StringElement{Content: "discount"},
							},
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.Icon{Class: "tags", Attributes: types.Attributes{types.AttrImageAlt: "Discount"}},
											types.StringElement{Content: " Cheap cheap!"},
										},
									},
								},
							},
						},
						types.LabeledListItem{
							Level: 1,
							Term: []interface{}{
								types.StringElement{Content: "item 2"},
							},
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{Content: "two"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))

			})

			It("inline icon in quoted text", func() {
				source := `an _italicized icon:warning[] message_`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "an "},
									types.QuotedText{
										Kind: types.Italic,
										Elements: []interface{}{
											types.StringElement{Content: "italicized "},
											types.Icon{Class: "warning"},
											types.StringElement{Content: " message"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})
			It("inline icon in marked text", func() {
				source := `#marked icon:warning[] message#`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Marked,
										Elements: []interface{}{
											types.StringElement{Content: "marked "},
											types.Icon{Class: "warning"},
											types.StringElement{Content: " message"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})
			It("inline icon in bold text", func() {
				source := `in *bold icon:warning[] message*`
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "in "},
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{Content: "bold "},
											types.Icon{Class: "warning"},
											types.StringElement{Content: " message"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})
			It("inline icon in monospace text", func() {
				source := "in `monospace icon:warning[] message`"
				expected := types.DraftDocument{
					Blocks: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "in "},
									types.QuotedText{
										Kind: types.Monospace,
										Elements: []interface{}{
											types.StringElement{Content: "monospace "},
											types.Icon{Class: "warning"},
											types.StringElement{Content: " message"},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})
		})
	})
})
