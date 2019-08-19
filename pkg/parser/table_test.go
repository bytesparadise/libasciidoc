package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("tables", func() {

	It("1-line table with 2 cells", func() {
		source := `|===
| *foo* foo  | _bar_  
|===
`
		expected := types.Table{
			Attributes: types.ElementAttributes{},
			Lines: []types.TableLine{
				{
					Cells: []types.InlineElements{
						{
							types.QuotedText{
								Kind: types.Bold,
								Elements: types.InlineElements{
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
								Elements: types.InlineElements{
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
		}
		verifyDocumentBlock(expected, source)
	})

	It("1-line table with 3 cells", func() {
		source := `|===
| *foo* foo  | _bar_  | baz
|===`
		expected := types.Table{
			Attributes: types.ElementAttributes{},
			Lines: []types.TableLine{
				{
					Cells: []types.InlineElements{
						{
							types.QuotedText{
								Kind: types.Bold,
								Elements: types.InlineElements{
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
								Elements: types.InlineElements{
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
		}
		verifyDocumentBlock(expected, source)
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
		expected := types.Table{
			Attributes: types.ElementAttributes{
				types.AttrTitle: "table title",
			},
			Header: types.TableLine{
				Cells: []types.InlineElements{
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
					Cells: []types.InlineElements{
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
					Cells: []types.InlineElements{
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
		}
		verifyDocumentBlock(expected, source)
	})

	It("empty table ", func() {
		source := `|===
|===`
		expected := types.Table{
			Attributes: types.ElementAttributes{},
			Lines:      []types.TableLine{},
		}
		verifyDocumentBlock(expected, source)
	})
})
