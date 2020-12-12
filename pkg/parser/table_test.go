package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("tables", func() {

	It("1-line table with 2 cells", func() {
		source := `|===
| *foo* foo  | _bar_  
|===
`
		expected := types.DraftDocument{
			Elements: []interface{}{
				types.Table{
					Columns: []types.TableColumn{
						{Width: "50", VAlign: "top", HAlign: "left"},
						{Width: "50", VAlign: "top", HAlign: "left"},
					},
					Lines: []types.TableLine{
						{
							Cells: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{
												Content: "foo",
											},
										},
									},
									types.StringElement{
										Content: " foo  ",
									},
								},
								{
									types.QuotedText{
										Kind: types.Italic,
										Elements: []interface{}{
											types.StringElement{
												Content: "bar",
											},
										},
									},
									types.StringElement{
										Content: "  ",
									},
								},
							},
						},
					},
				},
			},
		}
		Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
	})

	It("1-line table with 3 cells", func() {
		source := `|===
| *foo* foo  | _bar_  | baz
|===`
		expected := types.DraftDocument{
			Elements: []interface{}{
				types.Table{
					Columns: []types.TableColumn{
						{Width: "33.3333", VAlign: "top", HAlign: "left"},
						{Width: "33.3333", VAlign: "top", HAlign: "left"},
						{Width: "33.3334", VAlign: "top", HAlign: "left"},
					},
					Lines: []types.TableLine{
						{
							Cells: [][]interface{}{
								{
									types.QuotedText{
										Kind: types.Bold,
										Elements: []interface{}{
											types.StringElement{
												Content: "foo",
											},
										},
									},
									types.StringElement{
										Content: " foo  ",
									},
								},
								{
									types.QuotedText{
										Kind: types.Italic,
										Elements: []interface{}{
											types.StringElement{
												Content: "bar",
											},
										},
									},
									types.StringElement{
										Content: "  ",
									},
								},
								{
									types.StringElement{
										Content: "baz",
									},
								},
							},
						},
					},
				},
			},
		}
		Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
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
		expected := types.DraftDocument{
			Elements: []interface{}{
				types.Table{
					Attributes: types.Attributes{
						types.AttrTitle: "table title",
					},
					Columns: []types.TableColumn{
						{Width: "50", HAlign: "left", VAlign: "top"},
						{Width: "50", HAlign: "left", VAlign: "top"},
					},

					Header: types.TableLine{
						Cells: [][]interface{}{
							{
								types.StringElement{
									Content: "heading 1 ",
								},
							},
							{
								types.StringElement{
									Content: "heading 2",
								},
							},
						},
					},

					Lines: []types.TableLine{
						{
							Cells: [][]interface{}{
								{
									types.StringElement{
										Content: "row 1, column 1",
									},
								},
								{
									types.StringElement{
										Content: "row 1, column 2",
									},
								},
							},
						},
						{
							Cells: [][]interface{}{
								{
									types.StringElement{
										Content: "row 2, column 1",
									},
								},
								{
									types.StringElement{
										Content: "row 2, column 2",
									},
								},
							},
						},
					},
				},
			},
		}
		Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
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
		expected := types.DraftDocument{
			Elements: []interface{}{
				types.Table{
					Attributes: types.Attributes{
						types.AttrTitle:   "table title",
						types.AttrOptions: []interface{}{"autowidth"},
						types.AttrRoles:   []interface{}{"role1", "stretch"},
						types.AttrID:      "anchor",
					},
					Header: types.TableLine{
						Cells: [][]interface{}{
							{
								types.StringElement{
									Content: "heading 1 ",
								},
							},
							{
								types.StringElement{
									Content: "heading 2",
								},
							},
						},
					},
					Columns: []types.TableColumn{
						// autowidth clears width
						{HAlign: "left", VAlign: "top"},
						{HAlign: "left", VAlign: "top"},
					},
					Lines: []types.TableLine{
						{
							Cells: [][]interface{}{
								{
									types.StringElement{
										Content: "row 1, column 1",
									},
								},
								{
									types.StringElement{
										Content: "row 1, column 2",
									},
								},
							},
						},
						{
							Cells: [][]interface{}{
								{
									types.StringElement{
										Content: "row 2, column 1",
									},
								},
								{
									types.StringElement{
										Content: "row 2, column 2",
									},
								},
							},
						},
					},
				},
			},
		}
		Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
	})

	It("empty table ", func() {
		source := `|===
|===`
		expected := types.DraftDocument{
			Elements: []interface{}{
				types.Table{
					Columns: []types.TableColumn{},
					Lines:   []types.TableLine{},
				},
			},
		}
		Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
	})

	It("empty table with cols attr", func() {
		source := "[cols=\"3,2,5\"]\n|===\n|==="
		expected := types.DraftDocument{
			Elements: []interface{}{
				types.Table{
					Attributes: types.Attributes{
						types.AttrCols: "3,2,5",
					},
					Columns: []types.TableColumn{
						{Width: "30", HAlign: "left", VAlign: "top"},
						{Width: "20", HAlign: "left", VAlign: "top"},
						{Width: "50", HAlign: "left", VAlign: "top"},
					},
					Lines: []types.TableLine{},
				},
			},
		}
		Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
	})

	It("autowidth overrides column widths", func() {
		source := "[%autowidth,cols=\"3,2,5\"]\n|===\n|==="
		expected := types.DraftDocument{
			Elements: []interface{}{
				types.Table{
					Attributes: types.Attributes{
						types.AttrOptions: []interface{}{"autowidth"},
						types.AttrCols:    "3,2,5",
					},
					Columns: []types.TableColumn{
						{HAlign: "left", VAlign: "top"},
						{HAlign: "left", VAlign: "top"},
						{HAlign: "left", VAlign: "top"},
					},
					Lines: []types.TableLine{},
				},
			},
		}
		Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
	})

	It("column autowidth", func() {
		source := "[cols=\"30,~,~\"]\n|===\n|==="
		expected := types.DraftDocument{
			Elements: []interface{}{
				types.Table{
					Attributes: types.Attributes{
						types.AttrCols: "30,~,~",
					},
					Columns: []types.TableColumn{
						{Width: "30", HAlign: "left", VAlign: "top"},
						{HAlign: "left", VAlign: "top"},
						{HAlign: "left", VAlign: "top"},
					},
					Lines: []types.TableLine{},
				},
			},
		}
		Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
	})

	It("columns with repeat", func() {
		source := "[cols=\"3*10,2*~\"]\n|===\n|==="
		expected := types.DraftDocument{
			Elements: []interface{}{
				types.Table{
					Attributes: types.Attributes{
						types.AttrCols: "3*10,2*~",
					},
					Columns: []types.TableColumn{
						{Width: "10", HAlign: "left", VAlign: "top"},
						{Width: "10", HAlign: "left", VAlign: "top"},
						{Width: "10", HAlign: "left", VAlign: "top"},
						{HAlign: "left", VAlign: "top"},
						{HAlign: "left", VAlign: "top"},
					},
					Lines: []types.TableLine{},
				},
			},
		}
		Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
	})
	It("columns with alignment changes", func() {
		source := "[cols=\"2*^.^,<,.>\"]\n|===\n|==="
		expected := types.DraftDocument{Elements: []interface{}{
			types.Table{
				Attributes: types.Attributes{
					types.AttrCols: "2*^.^,<,.>",
				},
				Columns: []types.TableColumn{
					{Width: "25", HAlign: "center", VAlign: "middle"},
					{Width: "25", HAlign: "center", VAlign: "middle"},
					{Width: "25", HAlign: "left", VAlign: "top"},
					{Width: "25", HAlign: "left", VAlign: "bottom"},
				},
				Lines: []types.TableLine{},
			},
		},
		}
		Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
	})

	// TODO: This checks that we parse the styles -- we don't actually do anything with them further yet.
	It("columns with alignment changes and styles", func() {
		source := "[cols=\"2*^.^d,<e,.>s\"]\n|===\n|==="
		expected := types.DraftDocument{
			FrontMatter: types.FrontMatter{Content: nil},
			Elements: []interface{}{
				types.Table{
					Attributes: types.Attributes{
						types.AttrCols: "2*^.^d,<e,.>s",
					},
					Columns: []types.TableColumn{
						{Width: "25", HAlign: "center", VAlign: "middle"}, // "d" is aliased to ""
						{Width: "25", HAlign: "center", VAlign: "middle"},
						{Width: "25", HAlign: "left", VAlign: "top", Style: "e"},
						{Width: "25", HAlign: "left", VAlign: "bottom", Style: "s"},
					},
					Lines: []types.TableLine{},
				},
			},
		}
		Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
	})
})
