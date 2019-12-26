package renderer_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("preambles", func() {

	sectionATitle := []interface{}{
		types.StringElement{Content: "Section A"},
	}

	sectionBTitle := []interface{}{
		types.StringElement{Content: "Section B"},
	}

	It("doc without sections", func() {
		source := types.Document{
			Attributes: types.DocumentAttributes{
				types.AttrTitle: "foo",
			},
			ElementReferences: types.ElementReferences{
				"section_a": sectionATitle,
				"section_b": sectionBTitle,
			},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{Content: "a short paragraph"},
						},
					},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{Content: "another short paragraph"},
						},
					},
				},
			},
		}
		expected := types.Document{
			Attributes: types.DocumentAttributes{
				types.AttrTitle: "foo",
			},
			ElementReferences: types.ElementReferences{
				"section_a": sectionATitle,
				"section_b": sectionBTitle,
			},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{Content: "a short paragraph"},
						},
					},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{Content: "another short paragraph"},
						},
					},
				},
			},
		}
		Expect(source).To(HavePreamble(expected))
	})

	It("doc with 1-paragraph preamble", func() {
		source := types.Document{
			Attributes: types.DocumentAttributes{
				types.AttrTitle: "foo",
			},
			ElementReferences: types.ElementReferences{
				"section_a": sectionATitle,
				"section_b": sectionBTitle,
			},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{Content: "a short paragraph"},
						},
					},
				},
				types.BlankLine{},
				types.Section{
					Level: 1,
					Title: sectionATitle,
					Attributes: types.ElementAttributes{
						types.AttrID:       "section_a",
						types.AttrCustomID: false,
					},

					Elements: []interface{}{},
				},
				types.Section{
					Level: 1,
					Title: sectionBTitle,
					Attributes: types.ElementAttributes{
						types.AttrID:       "section_b",
						types.AttrCustomID: false,
					},
					Elements: []interface{}{},
				},
			},
		}
		expected := types.Document{
			Attributes: types.DocumentAttributes{
				types.AttrTitle: "foo",
			},
			ElementReferences: types.ElementReferences{
				"section_a": sectionATitle,
				"section_b": sectionBTitle,
			},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Preamble{
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a short paragraph"},
								},
							},
						},
						types.BlankLine{},
					},
				},
				types.Section{
					Level: 1,
					Title: sectionATitle,
					Attributes: types.ElementAttributes{
						types.AttrID:       "section_a",
						types.AttrCustomID: false,
					},

					Elements: []interface{}{},
				},
				types.Section{
					Level: 1,
					Title: sectionBTitle,
					Attributes: types.ElementAttributes{
						types.AttrID:       "section_b",
						types.AttrCustomID: false,
					},
					Elements: []interface{}{},
				},
			},
		}
		Expect(source).To(HavePreamble(expected))
	})

	It("doc with 2-paragraph preamble", func() {
		source := types.Document{
			Attributes: types.DocumentAttributes{
				types.AttrTitle: "foo",
			},
			ElementReferences: types.ElementReferences{
				"section_a": sectionATitle,
				"section_b": sectionBTitle,
			},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{Content: "a short paragraph"},
						},
					},
				},
				types.BlankLine{},
				types.Paragraph{
					Attributes: types.ElementAttributes{},
					Lines: [][]interface{}{
						{
							types.StringElement{Content: "another short paragraph"},
						},
					},
				},
				types.BlankLine{},
				types.Section{
					Level: 1,
					Title: sectionATitle,
					Attributes: types.ElementAttributes{
						types.AttrID:       "section_a",
						types.AttrCustomID: false,
					},

					Elements: []interface{}{},
				},
				types.Section{
					Level: 1,
					Title: sectionBTitle,
					Attributes: types.ElementAttributes{
						types.AttrID:       "section_b",
						types.AttrCustomID: false,
					},
					Elements: []interface{}{},
				},
			},
		}
		expected := types.Document{
			Attributes: types.DocumentAttributes{
				types.AttrTitle: "foo",
			},
			ElementReferences: types.ElementReferences{
				"section_a": sectionATitle,
				"section_b": sectionBTitle,
			},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Preamble{
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "a short paragraph"},
								},
							},
						},
						types.BlankLine{},
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{Content: "another short paragraph"},
								},
							},
						},
						types.BlankLine{},
					},
				},
				types.Section{
					Level: 1,
					Title: sectionATitle,
					Attributes: types.ElementAttributes{
						types.AttrID:       "section_a",
						types.AttrCustomID: false,
					},

					Elements: []interface{}{},
				},
				types.Section{
					Level: 1,
					Title: sectionBTitle,
					Attributes: types.ElementAttributes{
						types.AttrID:       "section_b",
						types.AttrCustomID: false,
					},
					Elements: []interface{}{},
				},
			},
		}
		Expect(source).To(HavePreamble(expected))
	})

})
