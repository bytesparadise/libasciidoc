package parser

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("insert preambles", func() {

	frontmatter := &types.FrontMatter{
		Attributes: map[string]interface{}{
			"draft": true,
		},
	}
	headerWithTitle := &types.DocumentHeader{
		Title: []interface{}{
			&types.StringElement{
				Content: "title",
			},
		},
		Elements: []interface{}{
			&types.AttributeDeclaration{
				Name:  "biscuits",
				Value: "cookies",
			},
		},
	}

	headerWithoutTitle := &types.DocumentHeader{
		Elements: []interface{}{
			&types.AttributeDeclaration{
				Name:  "biscuits",
				Value: "cookies",
			},
		},
	}
	sectionA := &types.Section{
		Level: 1,
		Title: []interface{}{
			&types.StringElement{Content: "Section A"},
		},
		Attributes: types.Attributes{
			types.AttrID: "_section_a",
		},
	}
	sectionB := &types.Section{
		Level: 1,
		Title: []interface{}{
			&types.StringElement{Content: "Section B"},
		},
		Attributes: types.Attributes{
			types.AttrID: "_section_b",
		},
	}
	paragraph := &types.Paragraph{
		Elements: []interface{}{
			&types.StringElement{
				Content: "a short paragraph",
			},
		},
	}
	anotherParagraph := &types.Paragraph{
		Elements: []interface{}{
			&types.StringElement{
				Content: "another short paragraph",
			},
		},
	}
	blankline := &types.BlankLine{}

	Context("no insertion", func() {

		It("should not insert when no sections", func() {
			// given
			doc := &types.Document{
				Elements: []interface{}{
					headerWithTitle,
					paragraph,
					blankline,
					paragraph,
				},
			}
			expected := &types.Document{
				Elements: []interface{}{
					headerWithTitle,
					paragraph,
					blankline,
					paragraph,
				},
			}
			// when
			insertPreamble(doc)
			// then
			Expect(doc).To(Equal(expected))
		})

		It("should not insert when no header", func() {
			// given
			doc := &types.Document{
				Elements: []interface{}{
					paragraph,
					sectionA,
					sectionB,
				},
			}
			expected := &types.Document{
				Elements: []interface{}{
					paragraph,
					sectionA,
					sectionB,
				},
			}
			// when
			insertPreamble(doc)
			// then
			Expect(doc).To(Equal(expected))
		})

		It("should not insert when header has no title", func() {
			// given
			doc := &types.Document{
				Elements: []interface{}{
					headerWithoutTitle,
					paragraph,
					sectionA,
					sectionB,
				},
			}
			expected := &types.Document{
				Elements: []interface{}{
					headerWithoutTitle,
					paragraph,
					sectionA,
					sectionB,
				},
			}
			// when
			insertPreamble(doc)
			// then
			Expect(doc).To(Equal(expected))
		})
	})

	It("should insert preamble with 1 paragraph and blankline", func() {
		// given
		doc := &types.Document{
			Elements: []interface{}{
				headerWithTitle,
				paragraph,
				blankline,
				sectionA,
				sectionB,
			},
		}
		expected := &types.Document{
			Elements: []interface{}{
				headerWithTitle,
				&types.Preamble{
					Elements: []interface{}{
						paragraph,
						blankline,
					},
				},
				sectionA,
				sectionB,
			},
		}
		// when
		insertPreamble(doc)
		// then
		Expect(doc).To(Equal(expected))
	})

	It("should insert preamble with 2 paragraphs and blanklines", func() {
		// given
		doc := &types.Document{
			Elements: []interface{}{
				headerWithTitle,
				paragraph,
				blankline,
				anotherParagraph,
				blankline,
				sectionA,
				sectionB,
			},
		}
		expected := &types.Document{
			Elements: []interface{}{
				headerWithTitle,
				&types.Preamble{
					Elements: []interface{}{
						paragraph,
						blankline,
						anotherParagraph,
						blankline,
					},
				},
				sectionA,
				sectionB,
			},
		}
		// when
		insertPreamble(doc)
		// then
		Expect(doc).To(Equal(expected))
	})

	It("should insert preamble with 1 paragraph and blankline when doc has frontmatter", func() {
		// given
		doc := &types.Document{
			Elements: []interface{}{
				frontmatter,
				headerWithTitle,
				paragraph,
				blankline,
				sectionA,
				sectionB,
			},
		}
		expected := &types.Document{
			Elements: []interface{}{
				frontmatter,
				headerWithTitle,
				&types.Preamble{
					Elements: []interface{}{
						paragraph,
						blankline,
					},
				},
				sectionA,
				sectionB,
			},
		}
		// when
		insertPreamble(doc)
		// then
		Expect(doc).To(Equal(expected))
	})

})
