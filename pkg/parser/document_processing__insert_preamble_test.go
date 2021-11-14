package parser

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("insert preambles", func() {

	header := &types.DocumentHeader{
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
					header,
					paragraph,
					blankline,
					paragraph,
				},
			}
			expected := &types.Document{
				Elements: []interface{}{
					header,
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
	})

	It("should insert preamble with 1 paragraph and blankline", func() {
		// given
		doc := &types.Document{
			Elements: []interface{}{
				header,
				paragraph,
				blankline,
				sectionA,
				sectionB,
			},
		}
		expected := &types.Document{
			Elements: []interface{}{
				header,
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
				header,
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
				header,
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

})
