package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("documents", func() {

	Context("draft documents", func() {

		It("should parse empty document", func() {
			source := ``
			expected := types.DraftDocument{}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("should parse header without empty first line", func() {
			source := `= My title
Garrett D'Amore
1.0, July 4, 2020
`
			expected := types.DraftDocument{
				Attributes: types.Attributes{
					"revnumber":      "1.0",
					"revdate":        "July 4, 2020",
					"author":         "Garrett D'Amore",
					"authorinitials": "GD",
					"authors": []types.DocumentAuthor{
						{
							FullName: "Garrett D'Amore",
							Email:    "",
						},
					},
					"firstname": "Garrett",
					"lastname":  "D'Amore",
					"revision": types.DocumentRevision{
						Revnumber: "1.0",
						Revdate:   "July 4, 2020",
						Revremark: "",
					},
				},
				Elements: []interface{}{
					types.Section{
						Level: 0,
						Attributes: types.Attributes{
							"id": "_my_title",
						},
						Title: []interface{}{
							types.StringElement{
								Content: "My title",
							},
						},
						Elements: []interface{}{},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))

		})

		It("should parse header with empty first line", func() {
			source := `
= My title
Garrett D'Amore
1.0, July 4, 2020`
			expected := types.DraftDocument{
				Attributes: types.Attributes{
					"revnumber":      "1.0",
					"revdate":        "July 4, 2020",
					"author":         "Garrett D'Amore",
					"authorinitials": "GD",
					"authors": []types.DocumentAuthor{
						{
							FullName: "Garrett D'Amore",
							Email:    "",
						},
					},
					"firstname": "Garrett",
					"lastname":  "D'Amore",
					"revision": types.DocumentRevision{
						Revnumber: "1.0",
						Revdate:   "July 4, 2020",
						Revremark: "",
					},
				},
				Elements: []interface{}{
					types.Section{
						Level: 0,
						Attributes: types.Attributes{
							"id": "_my_title",
						},
						Title: []interface{}{
							types.StringElement{
								Content: "My title",
							},
						},
						Elements: []interface{}{},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))

		})
	})

	Context("final documents", func() {

		It("should parse empty document", func() {
			source := ``
			expected := types.Document{
				Elements: []interface{}{},
			}
			Expect(ParseDocument(source)).To(Equal(expected))
		})
	})
})
