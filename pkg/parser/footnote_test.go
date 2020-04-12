package parser_test

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("footnotes - draft", func() {

	BeforeEach(func() {
		types.ResetFootnoteSequence()
	})

	It("footnote with single-line content", func() {
		footnoteContent := "some content"
		source := fmt.Sprintf(`foo footnote:[%s]`, footnoteContent)
		footnote1 := types.Footnote{
			ID: 0,
			Elements: []interface{}{
				types.StringElement{
					Content: footnoteContent,
				},
			},
		}
		expected := types.DraftDocument{
			Blocks: []interface{}{
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "foo ",
							},
							footnote1,
						},
					},
				},
			},
		}
		Expect(ParseDraftDocument(source)).To(Equal(expected)) // need to get the whole document here
	})

	It("footnote with single-line rich content", func() {
		source := `foo footnote:[some *rich* https://foo.com[content]]`
		footnote1 := types.Footnote{
			ID: 0,
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
					Attributes: types.ElementAttributes{
						types.AttrInlineLinkText: []interface{}{
							types.StringElement{
								Content: "content",
							},
						},
					},
					Location: types.Location{
						Elements: []interface{}{
							types.StringElement{
								Content: "https://foo.com",
							},
						},
					},
				},
			},
		}
		expected := types.DraftDocument{
			Blocks: []interface{}{
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "foo ",
							},
							footnote1,
						},
					},
				},
			},
		}
		Expect(ParseDraftDocument(source)).To(Equal(expected)) // need to get the whole document here
	})

	It("footnote in a paragraph", func() {
		source := `This is another paragraph.footnote:[I am footnote text and will be displayed at the bottom of the article.]`
		expected := types.DraftDocument{
			Blocks: []interface{}{
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "This is another paragraph.",
							},
							types.Footnote{
								ID: 0,
								Elements: []interface{}{
									types.StringElement{
										Content: "I am footnote text and will be displayed at the bottom of the article.",
									},
								},
							},
						},
					},
				},
			},
		}
		Expect(ParseDraftDocument(source)).To(Equal(expected)) // need to get the whole document here
	})

	It("multiple footnotes including a reference", func() {
		source := `A statement.footnote:[a regular footnote.]
A bold statement!footnote:disclaimer[Opinions are my own.]

Another outrageous statement.footnote:disclaimer[]`
		expected := types.DraftDocument{
			Blocks: []interface{}{
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "A statement.",
							},
							types.Footnote{
								ID: 0,
								Elements: []interface{}{
									types.StringElement{
										Content: "a regular footnote.",
									},
								},
							},
						},
						{
							types.StringElement{
								Content: "A bold statement!",
							},
							types.Footnote{
								ID:  1,
								Ref: "disclaimer",
								Elements: []interface{}{
									types.StringElement{
										Content: "Opinions are my own.",
									},
								},
							},
						},
					},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "Another outrageous statement.",
							},
							types.Footnote{
								ID:       -1,
								Ref:      "disclaimer",
								Elements: []interface{}{},
							},
						},
					},
				},
			},
		}
		Expect(ParseDraftDocument(source)).To(Equal(expected))
	})

	It("footnotes in document", func() {

		source := `= title
:idprefix: id_

a premable with a footnote:[foo]

== section 1 footnote:[bar]

a paragraph with another footnote:[baz]`
		footnote1 := types.Footnote{
			ID: 0,
			Elements: []interface{}{
				types.StringElement{
					Content: "foo",
				},
			},
		}
		footnote2 := types.Footnote{
			ID: 1,
			Elements: []interface{}{
				types.StringElement{
					Content: "bar",
				},
			},
		}
		footnote3 := types.Footnote{
			ID: 2,
			Elements: []interface{}{
				types.StringElement{
					Content: "baz",
				},
			},
		}
		docTitle := []interface{}{
			types.StringElement{
				Content: "title",
			},
		}
		section1Title := []interface{}{
			types.StringElement{
				Content: "section 1 ",
			},
			footnote2,
		}
		expected := types.DraftDocument{
			Blocks: []interface{}{
				types.Section{
					Level:      0,
					Title:      docTitle,
					Attributes: types.ElementAttributes{},
					Elements:   []interface{}{},
				},
				types.DocumentAttributeDeclaration{
					Name:  "idprefix",
					Value: "id_",
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "a premable with a ",
							},
							footnote1,
						},
					},
				},
				types.BlankLine{},
				types.Section{
					Attributes: types.ElementAttributes{},
					Level:      1,
					Title:      section1Title,
					Elements:   []interface{}{},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "a paragraph with another ",
							},
							footnote3,
						},
					},
				},
			},
		}
		Expect(ParseDraftDocument(source)).To(Equal(expected)) // need to get the whole document here
	})
})

var _ = Describe("footnotes - document", func() {

	BeforeEach(func() {
		types.ResetFootnoteSequence()
	})

	It("footnote with single-line content", func() {
		footnoteContent := "some content"
		source := fmt.Sprintf(`foo footnote:[%s]`, footnoteContent)
		footnote1 := types.Footnote{
			ID: 0,
			Elements: []interface{}{
				types.StringElement{
					Content: footnoteContent,
				},
			},
		}
		expected := types.Document{
			Attributes:        types.DocumentAttributes{},
			ElementReferences: types.ElementReferences{},
			Footnotes: []types.Footnote{
				footnote1,
			},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "foo ",
							},
							footnote1,
						},
					},
				},
			},
		}
		Expect(ParseDocument(source)).To(Equal(expected)) // need to get the whole document here
	})

	It("footnote with single-line rich content", func() {
		source := `foo footnote:[some *rich* https://foo.com[content]]`
		footnote1 := types.Footnote{
			ID: 0,
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
					Attributes: types.ElementAttributes{
						types.AttrInlineLinkText: []interface{}{
							types.StringElement{
								Content: "content",
							},
						},
					},
					Location: types.Location{
						Elements: []interface{}{
							types.StringElement{
								Content: "https://foo.com",
							},
						},
					},
				},
			},
		}
		expected := types.Document{
			Attributes:        types.DocumentAttributes{},
			ElementReferences: types.ElementReferences{},
			Footnotes: []types.Footnote{
				footnote1,
			},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "foo ",
							},
							footnote1,
						},
					},
				},
			},
		}
		Expect(ParseDocument(source)).To(Equal(expected)) // need to get the whole document here
	})

	It("footnote in a paragraph", func() {
		source := `This is another paragraph.footnote:[I am footnote text and will be displayed at the bottom of the article.]`
		footnote1 := types.Footnote{
			ID: 0,
			Elements: []interface{}{
				types.StringElement{
					Content: "I am footnote text and will be displayed at the bottom of the article.",
				},
			},
		}
		expected := types.Document{
			Attributes:        types.DocumentAttributes{},
			ElementReferences: types.ElementReferences{},
			Footnotes: []types.Footnote{
				footnote1,
			},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "This is another paragraph.",
							},
							footnote1,
						},
					},
				},
			},
		}
		Expect(ParseDocument(source)).To(Equal(expected)) // need to get the whole document here
	})

	It("multiple footnotes including a reference", func() {
		source := `A statement.footnote:[a regular footnote.]
A bold statement!footnote:disclaimer[Opinions are my own.]

Another outrageous statement.footnote:disclaimer[]`

		footnote1 := types.Footnote{
			ID: 0,
			Elements: []interface{}{
				types.StringElement{
					Content: "a regular footnote.",
				},
			},
		}
		footnote2 := types.Footnote{
			ID:  1,
			Ref: "disclaimer",
			Elements: []interface{}{
				types.StringElement{
					Content: "Opinions are my own.",
				},
			},
		}
		footnote3 := types.Footnote{
			ID:       -1,
			Ref:      "disclaimer",
			Elements: []interface{}{},
		}
		expected := types.Document{
			Attributes:        types.DocumentAttributes{},
			ElementReferences: types.ElementReferences{},
			Footnotes: types.Footnotes{
				footnote1,
				footnote2,
				// footnote3, // not included since it has no element.
			},
			FootnoteReferences: types.FootnoteReferences{
				"disclaimer": footnote2,
			},
			Elements: []interface{}{
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "A statement.",
							},
							footnote1,
						},
						{
							types.StringElement{
								Content: "A bold statement!",
							},
							footnote2,
						},
					},
				},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{
								Content: "Another outrageous statement.",
							},
							footnote3,
						},
					},
				},
			},
		}
		Expect(ParseDocument(source)).To(Equal(expected))
	})

	It("footnotes in document", func() {

		source := `= title
:idprefix: id_

a premable with a footnote:[foo]

== section 1 footnote:[bar]

a paragraph with another footnote:[baz]`
		footnote1 := types.Footnote{
			ID: 0,
			Elements: []interface{}{
				types.StringElement{
					Content: "foo",
				},
			},
		}
		footnote2 := types.Footnote{
			ID: 1,
			Elements: []interface{}{
				types.StringElement{
					Content: "bar",
				},
			},
		}
		footnote3 := types.Footnote{
			ID: 2,
			Elements: []interface{}{
				types.StringElement{
					Content: "baz",
				},
			},
		}
		docTitle := []interface{}{
			types.StringElement{
				Content: "title",
			},
		}
		section1Title := []interface{}{
			types.StringElement{
				Content: "section 1 ",
			},
			footnote2,
		}
		expected := types.Document{
			Attributes: types.DocumentAttributes{
				"idprefix": "id_",
			},
			ElementReferences: types.ElementReferences{
				"id_title":     docTitle,
				"id_section_1": section1Title,
			},
			Footnotes: types.Footnotes{
				footnote1,
				footnote2,
				footnote3,
			},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Section{
					Level: 0,
					Title: docTitle,
					Attributes: types.ElementAttributes{
						types.AttrID: "id_title",
					},
					Elements: []interface{}{
						types.Preamble{ // preamble is inserted
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a premable with a ",
											},
											footnote1,
										},
									},
								},
							},
						},
						types.Section{
							Attributes: types.ElementAttributes{
								types.AttrID: "id_section_1",
							},
							Level: 1,
							Title: section1Title,
							Elements: []interface{}{
								types.Paragraph{
									Attributes: types.ElementAttributes{},
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a paragraph with another ",
											},
											footnote3,
										},
									},
								},
							},
						},
					},
				},
			},
		}
		Expect(ParseDocument(source)).To(Equal(expected)) // need to get the whole document here
	})
})
