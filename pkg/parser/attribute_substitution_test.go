package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("attribute substitutions", func() {

	Context("in final documents", func() {

		It("paragraph with attribute substitution", func() {
			source := `:author: Xavier

a paragraph written by {author}.`
			expected := &types.Document{
				Elements: []interface{}{
					&types.AttributeDeclaration{
						Name:  "author",
						Value: "Xavier",
					},
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "a paragraph written by Xavier.",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("paragraph with attribute resets", func() {
			source := `:author: Xavier
				
:!author1:
:author2!:
a paragraph written by {author}.`
			expected := &types.Document{
				Elements: []interface{}{
					&types.AttributeDeclaration{
						Name:  "author",
						Value: "Xavier",
					},
					&types.AttributeReset{
						Name: "author1",
					},
					&types.AttributeReset{
						Name: "author2",
					},
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "a paragraph written by Xavier.",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("header with 2 authors, revision and attributes", func() {
			source := `= Document Title
John Foo Doe <johndoe@example.com>; Jane the_Doe <jane@example.com>
v1.0, March 29, 2020: Updated revision
:toc:
:keywords: documentation, team, obstacles, journey, victory

This journey continues`
			expected := &types.Document{
				Elements: []interface{}{
					&types.DocumentHeader{
						Title: []interface{}{
							&types.StringElement{Content: "Document Title"},
						},
						Elements: []interface{}{
							&types.AttributeDeclaration{
								Name: types.AttrAuthors,
								Value: types.DocumentAuthors{
									{
										DocumentAuthorFullName: &types.DocumentAuthorFullName{
											FirstName:  "John",
											MiddleName: "Foo",
											LastName:   "Doe",
										},
										Email: "johndoe@example.com",
									},
									{
										DocumentAuthorFullName: &types.DocumentAuthorFullName{
											FirstName: "Jane",
											LastName:  "the Doe",
										},
										Email: "jane@example.com",
									},
								},
							},
							&types.AttributeDeclaration{
								Name: types.AttrRevision,
								Value: &types.DocumentRevision{
									Revnumber: "1.0",
									Revdate:   "March 29, 2020",
									Revremark: "Updated revision",
								},
							},
							&types.AttributeDeclaration{
								Name: types.AttrTableOfContents,
							},
							&types.AttributeDeclaration{
								Name:  "keywords",
								Value: "documentation, team, obstacles, journey, victory",
							},
						},
					},
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "This journey continues",
							},
						},
					},
				},
				TableOfContents: &types.TableOfContents{}, // TODO: should we include a ToC when it's empty?
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("paragraph with attribute substitution from front-matter", func() {
			source := `---
author: Xavier
---

a paragraph written by {author}.`
			expected := &types.Document{
				Elements: []interface{}{
					&types.FrontMatter{
						Attributes: types.Attributes{
							"author": "Xavier",
						},
					},
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "a paragraph written by Xavier.",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})
	})
})
