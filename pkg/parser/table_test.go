package parser_test

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"
	log "github.com/sirupsen/logrus"

	. "github.com/onsi/ginkgo"                  //nolint golint
	. "github.com/onsi/ginkgo/extensions/table" //nolint golint
	. "github.com/onsi/gomega"                  //nolint golint
)

var _ = Describe("tables", func() {

	Context("in final documents", func() {

		It("1-line table with 2 cells", func() {
			source := `|===
| *cookie* cookie  | _pasta_  
|===
`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Table{
						// Columns: []*types.TableColumn{
						// 	{Width: "50", VAlign: "top", HAlign: "left"},
						// 	{Width: "50", VAlign: "top", HAlign: "left"},
						// },
						Rows: []*types.TableRow{
							{
								Cells: []*types.TableCell{
									{
										Elements: []interface{}{
											&types.QuotedText{
												Kind: types.SingleQuoteBold,
												Elements: []interface{}{
													&types.StringElement{
														Content: "cookie",
													},
												},
											},
											&types.StringElement{
												Content: " cookie  ",
											},
										},
									},
									{
										Elements: []interface{}{
											&types.QuotedText{
												Kind: types.SingleQuoteItalic,
												Elements: []interface{}{
													&types.StringElement{
														Content: "pasta",
													},
												},
											},
											&types.StringElement{
												Content: "  ",
											},
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

		It("1-line table with 3 cells", func() {
			source := `|===
| *cookie* cookie  | _pasta_  | chocolate
|===`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Table{
						// Columns: []*types.TableColumn{
						// 	{Width: "33.3333", VAlign: "top", HAlign: "left"},
						// 	{Width: "33.3333", VAlign: "top", HAlign: "left"},
						// 	{Width: "33.3334", VAlign: "top", HAlign: "left"},
						// },
						Rows: []*types.TableRow{
							{
								Cells: []*types.TableCell{
									{
										Elements: []interface{}{
											&types.QuotedText{
												Kind: types.SingleQuoteBold,
												Elements: []interface{}{
													&types.StringElement{
														Content: "cookie",
													},
												},
											},
											&types.StringElement{
												Content: " cookie  ",
											},
										},
									},
									{
										Elements: []interface{}{
											&types.QuotedText{
												Kind: types.SingleQuoteItalic,
												Elements: []interface{}{
													&types.StringElement{
														Content: "pasta",
													},
												},
											},
											&types.StringElement{
												Content: "  ",
											},
										},
									},
									{
										Elements: []interface{}{
											&types.StringElement{
												Content: "chocolate",
											},
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

		It("2-line table with 3 cells", func() {
			source := `|===
| some cookies | some chocolate | some pasta
| more cookies | more chocolate | more pasta
|===`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Table{
						// Columns: []*types.TableColumn{
						// 	{Width: "33.3333", VAlign: "top", HAlign: "left"},
						// 	{Width: "33.3333", VAlign: "top", HAlign: "left"},
						// 	{Width: "33.3334", VAlign: "top", HAlign: "left"},
						// },
						Rows: []*types.TableRow{
							{
								Cells: []*types.TableCell{
									{
										Elements: []interface{}{
											&types.StringElement{
												Content: "some cookies ",
											},
										},
									},
									{
										Elements: []interface{}{
											&types.StringElement{
												Content: "some chocolate ",
											},
										},
									},
									{
										Elements: []interface{}{
											&types.StringElement{
												Content: "some pasta",
											},
										},
									},
								},
							},
							{
								Cells: []*types.TableCell{
									{
										Elements: []interface{}{
											&types.StringElement{
												Content: "more cookies ",
											},
										},
									},
									{
										Elements: []interface{}{
											&types.StringElement{
												Content: "more chocolate ",
											},
										},
									},
									{
										Elements: []interface{}{
											&types.StringElement{
												Content: "more pasta",
											},
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

		It("table with title, headers and 1 line per cell", func() {
			source := `.table title
|===
|heading 1 |heading 2

|row 1, column 1
|row 1, column 2

|row 2, column 1
|row 2, column 2
|===`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Table{
						Attributes: types.Attributes{
							types.AttrTitle: "table title",
						},
						// Columns: []*types.TableColumn{
						// 	{Width: "50", HAlign: "left", VAlign: "top"},
						// 	{Width: "50", HAlign: "left", VAlign: "top"},
						// },

						Header: &types.TableRow{
							Cells: []*types.TableCell{
								{
									Elements: []interface{}{
										&types.StringElement{
											Content: "heading 1 ",
										},
									},
								},
								{
									Elements: []interface{}{
										&types.StringElement{
											Content: "heading 2",
										},
									},
								},
							},
						},

						Rows: []*types.TableRow{
							{
								Cells: []*types.TableCell{
									{
										Elements: []interface{}{
											&types.StringElement{
												Content: "row 1, column 1",
											},
										},
									},
									{
										Elements: []interface{}{
											&types.StringElement{
												Content: "row 1, column 2",
											},
										},
									},
								},
							},
							{
								Cells: []*types.TableCell{
									{
										Elements: []interface{}{
											&types.StringElement{
												Content: "row 2, column 1",
											},
										},
									},
									{
										Elements: []interface{}{
											&types.StringElement{
												Content: "row 2, column 2",
											},
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

		It("table with title, headers, id and multiple roles, stretch", func() {
			source := `.table title
[#anchor.role1%autowidth.stretch]
|===
|heading 1 |heading 2

|row 1, column 1
|row 1, column 2

|row 2, column 1
|row 2, column 2
|===`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Table{
						Attributes: types.Attributes{
							types.AttrTitle:   "table title",
							types.AttrOptions: []interface{}{"autowidth"},
							types.AttrRoles:   []interface{}{"role1", "stretch"},
							types.AttrID:      "anchor",
						},
						Header: &types.TableRow{
							Cells: []*types.TableCell{
								{
									Elements: []interface{}{
										&types.StringElement{
											Content: "heading 1 ",
										},
									},
								},
								{
									Elements: []interface{}{
										&types.StringElement{
											Content: "heading 2",
										},
									},
								},
							},
						},
						// Columns: []*types.TableColumn{
						// 	// autowidth clears width
						// 	{HAlign: "left", VAlign: "top"},
						// 	{HAlign: "left", VAlign: "top"},
						// },
						Rows: []*types.TableRow{
							{
								Cells: []*types.TableCell{
									{
										Elements: []interface{}{
											&types.StringElement{
												Content: "row 1, column 1",
											},
										},
									},
									{
										Elements: []interface{}{
											&types.StringElement{
												Content: "row 1, column 2",
											},
										},
									},
								},
							},
							{
								Cells: []*types.TableCell{
									{
										Elements: []interface{}{
											&types.StringElement{
												Content: "row 2, column 1",
											},
										},
									},
									{
										Elements: []interface{}{
											&types.StringElement{
												Content: "row 2, column 2",
											},
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

		It("empty table ", func() {
			source := `|===
|===`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Table{
						// Columns: []*types.TableColumn{},
						// Lines:   []*types.TableLine{},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("with cols attribute", func() {
			source := `[cols="2*^.^,<,.>"]
|===
|===`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Table{
						Attributes: types.Attributes{
							types.AttrCols: []interface{}{
								&types.TableColumn{
									Multiplier: 2,
									HAlign:     types.HAlignCenter,
									VAlign:     types.VAlignMiddle,
									Weight:     1,
								},
								&types.TableColumn{
									Multiplier: 1,
									HAlign:     types.HAlignLeft,
									VAlign:     types.VAlignTop,
									Weight:     1,
								},
								&types.TableColumn{
									Multiplier: 1,
									HAlign:     types.HAlignLeft,
									VAlign:     types.VAlignBottom,
									Weight:     1,
								},
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("columns as document attribute", func() {
			source := `:cols: 2*^.^d,<e,.>s
			
[cols={cols}]
|===
|===`
			expected := &types.Document{
				// Attributes: types.Attributes{
				// 	"cols": "2*^.^d,<e,.>s",
				// },
				Elements: []interface{}{
					&types.AttributeDeclaration{
						Name:  "cols",
						Value: "2*^.^d,<e,.>s",
					},
					&types.Table{
						Attributes: types.Attributes{
							types.AttrCols: []interface{}{
								&types.TableColumn{
									Multiplier: 2,
									HAlign:     types.HAlignCenter,
									VAlign:     types.VAlignMiddle,
									Style:      types.DefaultStyle,
									Weight:     1,
								},
								&types.TableColumn{
									Multiplier: 1,
									HAlign:     types.HAlignLeft,
									VAlign:     types.VAlignTop,
									Style:      types.EmphasisStyle,
									Weight:     1,
								},
								&types.TableColumn{
									Multiplier: 1,
									HAlign:     types.HAlignLeft,
									VAlign:     types.VAlignBottom,
									Style:      types.StrongStyle,
									Weight:     1,
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

var _ = Describe("table cols", func() {

	DescribeTable("valid",
		func(source string, expected []*types.TableColumn) {
			// given
			log.Debugf("processing '%s'", source)
			content := strings.NewReader(source)
			// when parsing only (ie, no substitution applied)
			result, err := parser.ParseReader("", content, parser.Entrypoint("TableColumnsAttribute"))
			// then
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(BeAssignableToTypeOf([]interface{}{}))
			cols := result.([]interface{})
			// now, set the attribute in the table and call the `Columns()` method
			t := &types.Table{
				Attributes: types.Attributes{
					types.AttrCols: result,
				},
				Rows: []*types.TableRow{{}},
			}
			t.Rows[0].Cells = make([]*types.TableCell, len(cols))
			for i := range cols {
				t.Rows[0].Cells[i] = &types.TableCell{}
			}
			Expect(t.Columns()).To(Equal(expected))
		},

		Entry(`1`, `1`,
			[]*types.TableColumn{
				{
					Multiplier: 1,
					HAlign:     types.HAlignLeft,
					VAlign:     types.VAlignTop,
					Weight:     1,
					Width:      "100",
				},
			}),
		Entry(`20,~,~`, `20,~,~`,
			[]*types.TableColumn{
				{
					Multiplier: 1,
					HAlign:     types.HAlignLeft,
					VAlign:     types.VAlignTop,
					Weight:     20,
					Width:      "20",
				},
				{
					Multiplier: 1,
					HAlign:     types.HAlignLeft,
					VAlign:     types.VAlignTop,
					Autowidth:  true,
					Width:      "",
				},
				{
					Multiplier: 1,
					HAlign:     types.HAlignLeft,
					VAlign:     types.VAlignTop,
					Autowidth:  true,
					Width:      "",
				},
			}),

		Entry(`<,>`, `<,>`,
			[]*types.TableColumn{
				{
					Multiplier: 1,
					HAlign:     types.HAlignLeft,
					VAlign:     types.VAlignTop,
					Weight:     1,
					Width:      "50",
				},
				{
					Multiplier: 1,
					HAlign:     types.HAlignRight,
					VAlign:     types.VAlignTop,
					Weight:     1,
					Width:      "50",
				},
			}),
		Entry(`.<,.>`, `.<,.>`,
			[]*types.TableColumn{
				{
					Multiplier: 1,
					HAlign:     types.HAlignLeft,
					VAlign:     types.VAlignTop,
					Weight:     1,
					Width:      "50",
				},
				{
					Multiplier: 1,
					HAlign:     types.HAlignLeft,
					VAlign:     types.VAlignBottom,
					Weight:     1,
					Width:      "50",
				},
			}),
		Entry(`<.<,>.>`, `<.<,>.>`,
			[]*types.TableColumn{
				{
					Multiplier: 1,
					HAlign:     types.HAlignLeft,
					VAlign:     types.VAlignTop,
					Weight:     1,
					Width:      "50",
				},
				{
					Multiplier: 1,
					HAlign:     types.HAlignRight,
					VAlign:     types.VAlignBottom,
					Weight:     1,
					Width:      "50",
				},
			}),
		Entry(`<.<1,>.>2`, `<.<1,>.>2`,
			[]*types.TableColumn{
				{
					Multiplier: 1,
					HAlign:     types.HAlignLeft,
					VAlign:     types.VAlignTop,
					Weight:     1,
					Width:      "33.3333",
				},
				{
					Multiplier: 1,
					HAlign:     types.HAlignRight,
					VAlign:     types.VAlignBottom,
					Weight:     2,
					Width:      "66.6667",
				},
			}),
		Entry(`2*<.<1,1*>.>2`, `2*<.<1,1*>.>2`,
			[]*types.TableColumn{
				{
					Multiplier: 2,
					HAlign:     types.HAlignLeft,
					VAlign:     types.VAlignTop,
					Weight:     1,
					Width:      "25",
				},
				{
					Multiplier: 2,
					HAlign:     types.HAlignLeft,
					VAlign:     types.VAlignTop,
					Weight:     1,
					Width:      "25",
				},
				{
					Multiplier: 1,
					HAlign:     types.HAlignRight,
					VAlign:     types.VAlignBottom,
					Weight:     2,
					Width:      "50",
				},
			}),
		// with style
		Entry(`2*^.^d,<e,.>s`, `2*^.^d,<e,.>s`,
			[]*types.TableColumn{
				{
					Multiplier: 2,
					HAlign:     types.HAlignCenter,
					VAlign:     types.VAlignMiddle,
					Style:      types.DefaultStyle,
					Weight:     1,
					Width:      "25",
				},
				{
					Multiplier: 2,
					HAlign:     types.HAlignCenter,
					VAlign:     types.VAlignMiddle,
					Style:      types.DefaultStyle,
					Weight:     1,
					Width:      "25",
				},
				{
					Multiplier: 1,
					HAlign:     types.HAlignLeft,
					VAlign:     types.VAlignTop,
					Style:      types.EmphasisStyle,
					Weight:     1,
					Width:      "25",
				},
				{
					Multiplier: 1,
					HAlign:     types.HAlignLeft,
					VAlign:     types.VAlignBottom,
					Style:      types.StrongStyle,
					Weight:     1,
					Width:      "25",
				},
			}),
	)

	DescribeTable("invalid",
		func(source string) {
			// given
			log.Debugf("processing '%s'", source)
			content := strings.NewReader(source)
			// when parsing only (ie, no substitution applied)
			_, err := parser.ParseReader("", content, parser.Entrypoint("TableColumnsAttribute"))
			// then
			Expect(err).To(HaveOccurred())
		},

		// unknown case: should return an error
		Entry(`invalid`, `invalid`),
	)
})
