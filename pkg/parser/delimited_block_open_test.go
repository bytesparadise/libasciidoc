package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("open blocks", func() {

	Context("in final documents", func() {

		Context("without masquerade", func() {

			It("with single paragraph", func() {
				source := `--
some content
--`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Open,
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											Content: `some content`,
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with table", func() {
				source := `[#block-id]
.Block Title
--
[cols="2*^"]
|===
a|
[#id]
.A title
image::image.png[]
a|
[#another-id]
.Another title
image::another-image.png[]
|===
--`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Open,
							Attributes: types.Attributes{
								types.AttrID:    "block-id",
								types.AttrTitle: "Block Title",
							},
							Elements: []interface{}{
								&types.Table{
									Attributes: types.Attributes{
										types.AttrCols: []interface{}{
											&types.TableColumn{
												Multiplier: 2,
												HAlign:     types.HAlignCenter,
												VAlign:     types.VAlignTop,
												Weight:     1,
											},
										},
									},
									Rows: []*types.TableRow{
										{
											Cells: []*types.TableCell{
												{
													Format: "a",
													Elements: []interface{}{
														&types.ImageBlock{
															Attributes: types.Attributes{
																types.AttrID:    "id",
																types.AttrTitle: "A title",
															},
															Location: &types.Location{
																Path: "image.png",
															},
														},
													},
												},
												{
													Format: "a",
													Elements: []interface{}{
														&types.ImageBlock{
															Attributes: types.Attributes{
																types.AttrID:    "another-id",
																types.AttrTitle: "Another title",
															},
															Location: &types.Location{
																Path: "another-image.png",
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
					ElementReferences: types.ElementReferences{
						"block-id": "Block Title",
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})
	})
})
