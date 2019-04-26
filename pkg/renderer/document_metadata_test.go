package renderer_test

import (
	"context"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/davecgh/go-spew/spew"
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("document metadata", func() {

	Context("document authors", func() {

		It("should include no author", func() {
			actualContent := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Level:      0,
						Attributes: types.ElementAttributes{},
						Title:      types.SectionTitle{},
						Elements:   []interface{}{},
					},
				},
			}
			expectedContent := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Level:      0,
						Attributes: types.ElementAttributes{},
						Title:      types.SectionTitle{},
						Elements:   []interface{}{},
					},
				},
			}
			verifyDocumentMetadata(expectedContent, actualContent)
		})

		It("should include single author without middlename", func() {
			actualContent := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Level: 0,
						Attributes: types.ElementAttributes{
							types.AttrAuthors: []types.DocumentAuthor{
								{
									FullName: "Kismet Chameleon",
									Email:    "kismet@asciidoctor.org",
								},
							},
						},
						Title:    types.SectionTitle{},
						Elements: []interface{}{},
					},
				},
			}
			expectedContent := types.Document{
				Attributes: types.DocumentAttributes{
					"author":         "Kismet Chameleon",
					"firstname":      "Kismet",
					"lastname":       "Chameleon",
					"authorinitials": "KC",
					"email":          "kismet@asciidoctor.org",
				},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Level: 0,
						Attributes: types.ElementAttributes{
							types.AttrAuthors: []types.DocumentAuthor{
								{
									FullName: "Kismet Chameleon",
									Email:    "kismet@asciidoctor.org",
								},
							},
						},
						Title:    types.SectionTitle{},
						Elements: []interface{}{},
					},
				},
			}
			verifyDocumentMetadata(expectedContent, actualContent)
		})

		It("should include single author without middlename, last name and email", func() {
			actualContent := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Level: 0,
						Attributes: types.ElementAttributes{
							types.AttrAuthors: []types.DocumentAuthor{
								{
									FullName: "Kismet",
								},
							},
						},
						Title:    types.SectionTitle{},
						Elements: []interface{}{},
					},
				},
			}
			expectedContent := types.Document{
				Attributes: types.DocumentAttributes{
					"author":         "Kismet",
					"firstname":      "Kismet",
					"authorinitials": "K",
				},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Level: 0,
						Attributes: types.ElementAttributes{
							types.AttrAuthors: []types.DocumentAuthor{
								{
									FullName: "Kismet",
								},
							},
						},
						Title:    types.SectionTitle{},
						Elements: []interface{}{},
					},
				},
			}
			verifyDocumentMetadata(expectedContent, actualContent)
		})

		It("should include multiple authors", func() {
			actualContent := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Level: 0,
						Attributes: types.ElementAttributes{
							types.AttrAuthors: []types.DocumentAuthor{
								{
									FullName: "Kismet Rainbow Chameleon ",
									Email:    "kismet@asciidoctor.org",
								},
								{
									FullName: "Lazarus het_Draeke",
									Email:    "lazarus@asciidoctor.org",
								},
							},
						},
						Title:    types.SectionTitle{},
						Elements: []interface{}{},
					},
				},
			}
			expectedContent := types.Document{
				Attributes: types.DocumentAttributes{
					"author":           "Kismet Rainbow Chameleon",
					"firstname":        "Kismet",
					"middlename":       "Rainbow",
					"lastname":         "Chameleon",
					"authorinitials":   "KRC",
					"email":            "kismet@asciidoctor.org",
					"author_2":         "Lazarus het Draeke",
					"firstname_2":      "Lazarus",
					"lastname_2":       "het Draeke",
					"authorinitials_2": "Lh",
					"email_2":          "lazarus@asciidoctor.org",
				},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Level: 0,
						Attributes: types.ElementAttributes{
							types.AttrAuthors: []types.DocumentAuthor{
								{
									FullName: "Kismet Rainbow Chameleon ",
									Email:    "kismet@asciidoctor.org",
								},
								{
									FullName: "Lazarus het_Draeke",
									Email:    "lazarus@asciidoctor.org",
								},
							},
						},
						Title:    types.SectionTitle{},
						Elements: []interface{}{},
					},
				},
			}
			verifyDocumentMetadata(expectedContent, actualContent)
		})
	})

	Context("document revision", func() {

		It("should include full revision", func() {
			actualContent := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Level: 0,
						Attributes: types.ElementAttributes{
							types.AttrAuthors: []types.DocumentAuthor{
								{
									FullName: "Kismet",
								},
							},
							types.AttrRevision: types.DocumentRevision{
								Revnumber: "v1.0",
								Revdate:   "June 19, 2017",
								Revremark: "First incarnation",
							},
						},
						Title:    types.SectionTitle{},
						Elements: []interface{}{},
					},
				},
			}
			expectedContent := types.Document{
				Attributes: types.DocumentAttributes{
					"author":         "Kismet",
					"firstname":      "Kismet",
					"authorinitials": "K",
					"revnumber":      "v1.0",
					"revdate":        "June 19, 2017",
					"revremark":      "First incarnation",
				},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Level: 0,
						Attributes: types.ElementAttributes{
							types.AttrAuthors: []types.DocumentAuthor{
								{
									FullName: "Kismet",
								},
							},
							types.AttrRevision: types.DocumentRevision{
								Revnumber: "v1.0",
								Revdate:   "June 19, 2017",
								Revremark: "First incarnation",
							},
						},
						Title:    types.SectionTitle{},
						Elements: []interface{}{},
					},
				},
			}
			verifyDocumentMetadata(expectedContent, actualContent)
		})
	})

})

func verifyDocumentMetadata(expectedResult, actualContent types.Document) {
	ctx := renderer.Wrap(context.Background(), actualContent)
	renderer.ProcessDocumentHeader(ctx)
	GinkgoT().Logf("actual document: `%s`", spew.Sdump(ctx.Document))
	GinkgoT().Logf("expected document: `%s`", spew.Sdump(expectedResult))
	assert.EqualValues(GinkgoT(), expectedResult, ctx.Document)
}
