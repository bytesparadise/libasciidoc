package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("documents", func() {

	Context("raw documents", func() {

		It("should parse empty document", func() {
			source := ``
			expected := []types.DocumentFragment{}
			Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))
		})

		It("should parse header without empty first line", func() {
			source := `= My title
Garrett D'Amore
1.0, July 4, 2020
`
			expected := []types.DocumentFragment{
				{
					Position: types.Position{
						Start: 0,
						End:   45,
					},
					Elements: []interface{}{
						&types.DocumentHeader{
							Title: []interface{}{
								&types.StringElement{
									Content: "My title",
								},
							},
							Elements: []interface{}{
								&types.AttributeDeclaration{
									Name: types.AttrAuthors,
									Value: types.DocumentAuthors{
										{
											DocumentAuthorFullName: &types.DocumentAuthorFullName{
												FirstName: "Garrett",
												LastName:  "D'Amore",
											},
										},
									},
								},
								&types.AttributeDeclaration{
									Name: types.AttrRevision,
									Value: &types.DocumentRevision{
										Revnumber: "1.0",
										Revdate:   "July 4, 2020",
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))

		})

		It("should parse header with empty first line", func() {
			source := `
= My title
Garrett D'Amore
1.0, July 4, 2020`
			expected := []types.DocumentFragment{
				{
					Position: types.Position{
						Start: 0,
						End:   1,
					},
					Elements: []interface{}{
						&types.BlankLine{},
					},
				},
				{
					Position: types.Position{
						Start: 1,
						End:   45,
					},
					Elements: []interface{}{
						&types.DocumentHeader{
							Title: []interface{}{
								&types.StringElement{
									Content: "My title",
								},
							},
							Elements: []interface{}{
								&types.AttributeDeclaration{
									Name: types.AttrAuthors,
									Value: types.DocumentAuthors{
										{
											DocumentAuthorFullName: &types.DocumentAuthorFullName{
												FirstName: "Garrett",
												LastName:  "D'Amore",
											},
										},
									},
								},
								&types.AttributeDeclaration{
									Name: types.AttrRevision,
									Value: &types.DocumentRevision{
										Revnumber: "1.0",
										Revdate:   "July 4, 2020",
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDocumentFragments(source)).To(MatchDocumentFragments(expected))

		})
	})

	Context("in final documents", func() {

		It("should parse empty document", func() {
			source := ``
			expected := &types.Document{}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should parse basic document", func() {
			source := `== Lorem Ipsum
			
Lorem ipsum dolor sit amet, consetetur sadipscing elitr, 
sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, 
sed diam voluptua. 
At vero eos et accusam et justo duo dolores et ea rebum. 
Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. 
Lorem ipsum dolor sit amet, consetetur sadipscing elitr, 
sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, 
sed diam voluptua. 
At vero eos et accusam et justo duo dolores et ea rebum. 
Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit *amet*.`

			expected := &types.Document{
				Elements: []interface{}{
					&types.Section{
						Level: 1,
						Attributes: types.Attributes{
							types.AttrID: "_Lorem_Ipsum",
						},
						Title: []interface{}{
							&types.StringElement{
								Content: "Lorem Ipsum",
							},
						},
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										// suffix spaces are trimmed on each line
										Content: `Lorem ipsum dolor sit amet, consetetur sadipscing elitr,
sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat,
sed diam voluptua.
At vero eos et accusam et justo duo dolores et ea rebum.
Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.
Lorem ipsum dolor sit amet, consetetur sadipscing elitr,
sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat,
sed diam voluptua.
At vero eos et accusam et justo duo dolores et ea rebum.
Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit `,
									},
									&types.QuotedText{
										Kind: types.SingleQuoteBold,
										Elements: []interface{}{
											&types.StringElement{
												Content: "amet",
											},
										},
									},
									&types.StringElement{
										Content: ".",
									},
								},
							},
						},
					},
				},
				ElementReferences: types.ElementReferences{
					"_Lorem_Ipsum": []interface{}{
						&types.StringElement{Content: "Lorem Ipsum"},
					},
				},
				TableOfContents: &types.TableOfContents{
					MaxDepth: 2,
					Sections: []*types.ToCSection{
						{
							ID:    "_Lorem_Ipsum",
							Level: 1,
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})
	})
})
