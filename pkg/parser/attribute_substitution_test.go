package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("attribute substitutions", func() {

	Context("in final documents", func() {

		It("paragraph with attribute reference", func() {
			source := `:author: Xavier

a paragraph written by {author}.`
			expected := &types.Document{
				Elements: []interface{}{
					&types.DocumentHeader{
						Elements: []interface{}{
							&types.AttributeDeclaration{
								Name:  "author",
								Value: "Xavier",
							},
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

		It("paragraph with attribute resets", func() {
			source := `:author: Xavier
				
:!author1:
:author2!:
a paragraph written by {author}.`
			expected := &types.Document{
				Elements: []interface{}{
					&types.DocumentHeader{
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
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("paragraph with attribute reference from front-matter", func() {
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

		It("link with attribute reference with hyphen", func() {
			source := `:download-version: 1.0.0

a link to https://example.com/version/v{download-version}[here]`

			expected := &types.Document{
				Elements: []interface{}{
					&types.DocumentHeader{
						Elements: []interface{}{
							&types.AttributeDeclaration{
								Name:  "download-version",
								Value: "1.0.0",
							},
						},
					},
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "a link to ",
							},
							&types.InlineLink{
								Attributes: types.Attributes{
									types.AttrInlineLinkText: "here",
								},
								Location: &types.Location{
									Scheme: "https://",
									Path:   "example.com/version/v1.0.0",
								},
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("link with unknown attribute reference in URL", func() {
			source := `:version: 1.0.0

a link to https://example.com/version/v{unknown}[here]`

			expected := &types.Document{
				Elements: []interface{}{
					&types.DocumentHeader{
						Elements: []interface{}{
							&types.AttributeDeclaration{
								Name:  "version",
								Value: "1.0.0",
							},
						},
					},
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "a link to ",
							},
							&types.InlineLink{
								Attributes: types.Attributes{
									types.AttrInlineLinkText: "here",
								},
								Location: &types.Location{
									Scheme: "https://",
									Path:   "example.com/version/v{unknown}",
								},
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})
	})
})
