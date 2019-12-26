package renderer_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("document metadata", func() {

	Context("document authors", func() {

		It("should include no author", func() {
			source := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Section{
						Level:      0,
						Attributes: types.ElementAttributes{},
						Title:      []interface{}{},
						Elements:   []interface{}{},
					},
				},
			}
			expected := types.DocumentAttributes{}
			Expect(source).To(HaveMetadata(expected))
		})

		It("should include single author without middlename", func() {
			source := types.Document{
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
						Title:    []interface{}{},
						Elements: []interface{}{},
					},
				},
			}
			expected := types.DocumentAttributes{
				"author":         "Kismet Chameleon",
				"firstname":      "Kismet",
				"lastname":       "Chameleon",
				"authorinitials": "KC",
				"email":          "kismet@asciidoctor.org",
			}
			Expect(source).To(HaveMetadata(expected))
		})

		It("should include single author without middlename, last name and email", func() {
			source := types.Document{
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
						Title:    []interface{}{},
						Elements: []interface{}{},
					},
				},
			}
			expected := types.DocumentAttributes{
				"author":         "Kismet",
				"firstname":      "Kismet",
				"authorinitials": "K",
			}
			Expect(source).To(HaveMetadata(expected))
		})

		It("should include multiple authors", func() {
			source := types.Document{
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
						Title:    []interface{}{},
						Elements: []interface{}{},
					},
				},
			}
			expected := types.DocumentAttributes{
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
			}
			Expect(source).To(HaveMetadata(expected))
		})
	})

	Context("document revision", func() {

		It("should include full revision", func() {
			source := types.Document{
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
						Title:    []interface{}{},
						Elements: []interface{}{},
					},
				},
			}
			expected := types.DocumentAttributes{
				"author":         "Kismet",
				"firstname":      "Kismet",
				"authorinitials": "K",
				"revnumber":      "v1.0",
				"revdate":        "June 19, 2017",
				"revremark":      "First incarnation",
			}
			Expect(source).To(HaveMetadata(expected))
		})
	})

})
