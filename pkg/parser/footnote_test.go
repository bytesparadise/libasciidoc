package parser_test

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("footnotes - document", func() {

	Context("final documents", func() {

		It("footnote with single-line content", func() {
			footnoteContent := "some content"
			source := fmt.Sprintf(`foo footnote:[%s]`, footnoteContent)
			expected := types.Document{
				Footnotes: []types.Footnote{
					{
						ID: 1,
						Elements: []interface{}{
							types.StringElement{
								Content: footnoteContent,
							},
						},
					},
				},
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "foo ",
								},
								types.FootnoteReference{
									ID: 1,
								},
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected)) // need to get the whole document here
		})

		It("footnote with single-line rich content", func() {
			source := `foo footnote:[some *rich* https://foo.com[content]]`
			footnote1 := types.Footnote{
				Elements: []interface{}{
					types.StringElement{
						Content: "some ",
					},
					types.QuotedText{
						Kind: types.Bold,
						Elements: []interface{}{
							types.StringElement{
								Content: "rich",
							},
						},
					},
					types.StringElement{
						Content: " ",
					},
					types.InlineLink{
						Attributes: types.Attributes{
							types.AttrInlineLinkText: "content",
						},
						Location: types.Location{
							Scheme: "https://",
							Path: []interface{}{
								types.StringElement{
									Content: "foo.com",
								},
							},
						},
					},
				},
			}
			expected := types.Document{
				Footnotes: []types.Footnote{
					{
						ID:       1,
						Elements: footnote1.Elements,
					},
				},
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "foo ",
								},
								types.FootnoteReference{
									ID: 1,
								},
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected)) // need to get the whole document here
		})

		It("footnote in a paragraph", func() {
			source := `This is another paragraph.footnote:[I am footnote text and will be displayed at the bottom of the article.]`
			expected := types.Document{
				Footnotes: []types.Footnote{
					{
						ID: 1,
						Elements: []interface{}{
							types.StringElement{
								Content: "I am footnote text and will be displayed at the bottom of the article.",
							},
						},
					},
				},
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "This is another paragraph.",
								},
								types.FootnoteReference{
									ID: 1,
								},
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected)) // need to get the whole document here
		})

		It("multiple footnotes including a reference", func() {
			source := `A statement.footnote:[a regular footnote.]
A bold statement!footnote:disclaimer[Opinions are my own.]

Another outrageous statement.footnote:disclaimer[]`

			expected := types.Document{
				Footnotes: []types.Footnote{
					{
						ID: 1,
						Elements: []interface{}{
							types.StringElement{
								Content: "a regular footnote.",
							},
						},
					},
					{
						ID:  2,
						Ref: "disclaimer",
						Elements: []interface{}{
							types.StringElement{
								Content: "Opinions are my own.",
							},
						},
					},
				},
				Elements: []interface{}{
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "A statement.",
								},
								types.FootnoteReference{
									ID: 1,
								},
							},
							{
								types.StringElement{
									Content: "A bold statement!",
								},
								types.FootnoteReference{
									ID:  2,
									Ref: "disclaimer",
								},
							},
						},
					},
					types.Paragraph{
						Lines: [][]interface{}{
							{
								types.StringElement{
									Content: "Another outrageous statement.",
								},
								types.FootnoteReference{
									ID:        2,
									Ref:       "disclaimer",
									Duplicate: true, // this FootnoteReference targets an already-existing footnote
								},
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("footnotes in document", func() {

			source := `= title
:idprefix: id_

a premable with a footnote:[foo]

== section 1 footnote:[bar]

a paragraph with another footnote.footnote:[baz]`

			docTitle := []interface{}{
				types.StringElement{
					Content: "title",
				},
			}
			section1Title := []interface{}{
				types.StringElement{
					Content: "section 1 ",
				},
				types.FootnoteReference{
					ID: 2,
				},
			}
			expected := types.Document{
				Attributes: types.Attributes{
					"idprefix": "id_",
				},
				ElementReferences: types.ElementReferences{
					"id_title":     docTitle,
					"id_section_1": section1Title,
				},
				Footnotes: []types.Footnote{
					{
						ID: 1,
						Elements: []interface{}{
							types.StringElement{
								Content: "foo",
							},
						},
					},
					{
						ID: 2,
						Elements: []interface{}{
							types.StringElement{
								Content: "bar",
							},
						},
					},
					{
						ID: 3,
						Elements: []interface{}{
							types.StringElement{
								Content: "baz",
							},
						},
					},
				},
				Elements: []interface{}{
					types.Section{
						Level: 0,
						Title: docTitle,
						Attributes: types.Attributes{
							types.AttrID: "id_title",
						},
						Elements: []interface{}{
							types.Preamble{ // preamble is inserted
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{
													Content: "a premable with a ",
												},
												types.FootnoteReference{
													ID: 1,
												},
											},
										},
									},
								},
							},
							types.Section{
								Attributes: types.Attributes{
									types.AttrID: "id_section_1",
								},
								Level: 1,
								Title: section1Title,
								Elements: []interface{}{
									types.Paragraph{
										Lines: [][]interface{}{
											{
												types.StringElement{
													Content: "a paragraph with another footnote.",
												},
												types.FootnoteReference{
													ID: 3,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected)) // need to get the whole document here
		})
	})
})
