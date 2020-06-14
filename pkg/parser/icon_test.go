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
									types.InlineIcon{
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
									types.InlineIcon{
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
									types.InlineIcon{
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
									types.InlineIcon{
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
									types.InlineIcon{
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
									types.InlineIcon{
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
									types.InlineIcon{
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
									types.InlineIcon{
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
									types.InlineIcon{
										Class: "info",
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
