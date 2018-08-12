package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("tables", func() {

	It("1-line table with 2 cells", func() {
		actualContent := `|===
| *foo* foo  | _bar_  
|===
`
		expectedResult := types.Table{
			Attributes: map[string]interface{}{},
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
		verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
	})

	It("1-line table with 3 cells", func() {
		actualContent := `|===
| *foo* foo  | _bar_  | baz
|===`
		expectedResult := types.Table{
			Attributes: map[string]interface{}{},
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
		verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
	})

	It("table with title, headers and 1 line per cell", func() {
		actualContent := `.table title
|===
|heading 1 |heading 2

|row 1, column 1
|row 1, column 2

|row 2, column 1
|row 2, column 2
|===`
		expectedResult := types.Table{
			Attributes: map[string]interface{}{
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
				// types.BlankLine{},
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
		verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
	})

	It("empty table ", func() {
		actualContent := `|===
|===`
		expectedResult := types.Table{
			Attributes: map[string]interface{}{},
			Lines:      []types.TableLine{},
		}
		verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
	})
})
