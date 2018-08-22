package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("cross References", func() {

	Context("section reference", func() {

		It("xref with custom id", func() {
			actualContent := `[[thetitle]]
== a title

with some content linked to <<thetitle>>!`
			expectedResult := types.Document{
				Attributes: types.DocumentAttributes{},
				ElementReferences: map[string]interface{}{
					"thetitle": types.SectionTitle{
						Attributes: types.ElementAttributes{
							types.AttrID: "thetitle",
						},
						Content: types.InlineElements{
							types.StringElement{
								Content: "a title",
							},
						},
					},
				},
				Elements: []interface{}{
					types.Section{
						Level: 1,
						Title: types.SectionTitle{
							Attributes: types.ElementAttributes{
								types.AttrID: "thetitle",
							},
							Content: types.InlineElements{
								types.StringElement{
									Content: "a title",
								},
							},
						},
						Elements: []interface{}{
							types.BlankLine{},
							types.Paragraph{
								Attributes: types.ElementAttributes{},
								Lines: []types.InlineElements{
									{
										types.StringElement{Content: "with some content linked to "},
										types.CrossReference{ID: "thetitle"},
										types.StringElement{Content: "!"},
									},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("Document"))
		})
	})
})
