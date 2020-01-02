package parser

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("block filters", func() {

	It("should remove blank line", func() {
		actual := []interface{}{
			types.BlankLine{},
			types.Paragraph{
				Lines: [][]interface{}{
					{
						types.StringElement{},
					},
				},
			},
		}
		expected := []interface{}{
			types.Paragraph{
				Lines: [][]interface{}{
					{
						types.StringElement{},
					},
				},
			},
		}
		Expect(filter(actual, allMatchers...)).To(Equal(expected))
	})

	It("should retain blank line in a delimited block", func() {
		actual := []interface{}{
			types.DelimitedBlock{
				Kind: types.Fenced,
				Elements: []interface{}{
					types.BlankLine{},
				},
			},
			types.Paragraph{
				Lines: [][]interface{}{
					{
						types.StringElement{},
					},
				},
			},
		}
		expected := []interface{}{
			types.DelimitedBlock{
				Kind: types.Fenced,
				Elements: []interface{}{
					types.BlankLine{},
				},
			},
			types.Paragraph{
				Lines: [][]interface{}{
					{
						types.StringElement{},
					},
				},
			},
		}
		Expect(filter(actual, allMatchers...)).To(Equal(expected))
	})

	It("should remove document attribute declaration", func() {
		actual := []interface{}{
			types.DocumentAttributeDeclaration{},
			types.Paragraph{
				Lines: [][]interface{}{
					{
						types.StringElement{},
					},
				},
			},
		}
		expected := []interface{}{
			types.Paragraph{
				Lines: [][]interface{}{
					{
						types.StringElement{},
					},
				},
			},
		}
		Expect(filter(actual, allMatchers...)).To(Equal(expected))
	})

	It("should remove document attribute substitution", func() {
		actual := []interface{}{
			types.DocumentAttributeSubstitution{},
			types.Paragraph{
				Lines: [][]interface{}{
					{
						types.StringElement{},
					},
				},
			},
		}
		expected := []interface{}{
			types.Paragraph{
				Lines: [][]interface{}{
					{
						types.StringElement{},
					},
				},
			},
		}
		Expect(filter(actual, allMatchers...)).To(Equal(expected))
	})

	It("should remove document attribute reset", func() {
		actual := []interface{}{
			types.DocumentAttributeReset{},
			types.Paragraph{
				Lines: [][]interface{}{
					{
						types.StringElement{},
					},
				},
			},
		}
		expected := []interface{}{
			types.Paragraph{
				Lines: [][]interface{}{
					{
						types.StringElement{},
					},
				},
			},
		}
		Expect(filter(actual, allMatchers...)).To(Equal(expected))
	})

	It("should remove comment block", func() {
		actual := []interface{}{
			types.DelimitedBlock{
				Kind: types.Comment,
			},
			types.Paragraph{
				Lines: [][]interface{}{
					{
						types.StringElement{},
					},
				},
			},
		}
		expected := []interface{}{
			types.Paragraph{
				Lines: [][]interface{}{
					{
						types.StringElement{},
					},
				},
			},
		}
		Expect(filter(actual, allMatchers...)).To(Equal(expected))
	})

	It("should remove single line comment as a block", func() {
		actual := []interface{}{
			types.SingleLineComment{},
			types.Paragraph{
				Lines: [][]interface{}{
					{
						types.StringElement{},
					},
				},
			},
		}
		expected := []interface{}{
			types.Paragraph{
				Lines: [][]interface{}{
					{
						types.StringElement{},
					},
				},
			},
		}
		Expect(filter(actual, allMatchers...)).To(Equal(expected))
	})

	It("should remove single line comment in a paragraph", func() {
		actual := []interface{}{
			types.Paragraph{
				Lines: [][]interface{}{
					{
						types.StringElement{},
						types.SingleLineComment{},
					},
				},
			},
		}
		expected := []interface{}{
			types.Paragraph{
				Lines: [][]interface{}{
					{
						types.StringElement{},
					},
				},
			},
		}
		Expect(filter(actual, allMatchers...)).To(Equal(expected))
	})

	It("should retain paragraph with single line comment only", func() {
		actual := []interface{}{
			types.Paragraph{
				Lines: [][]interface{}{
					{
						types.SingleLineComment{},
					},
				},
			},
		}
		expected := []interface{}{
			types.Paragraph{
				Lines: [][]interface{}{},
			},
		}
		Expect(filter(actual, allMatchers...)).To(Equal(expected))
	})

	It("should retain paragraph with empty content", func() {
		actual := []interface{}{
			types.Paragraph{
				Lines: [][]interface{}{},
			},
		}
		expected := []interface{}{
			types.Paragraph{
				Lines: [][]interface{}{},
			},
		}
		Expect(filter(actual, allMatchers...)).To(Equal(expected))
	})

	It("should remove single line comment in an ordered list item", func() {
		actual := []interface{}{
			types.OrderedList{
				Items: []types.OrderedListItem{
					{
						Elements: []interface{}{
							types.StringElement{},
							types.SingleLineComment{},
						},
					},
				},
			},
		}
		expected := []interface{}{
			types.OrderedList{
				Items: []types.OrderedListItem{
					{
						Elements: []interface{}{
							types.StringElement{},
						},
					},
				},
			},
		}
		Expect(filter(actual, allMatchers...)).To(Equal(expected))
	})

	It("should remove single line comment in an ordered list item", func() {
		actual := []interface{}{
			types.UnorderedList{
				Items: []types.UnorderedListItem{
					{
						Elements: []interface{}{
							types.StringElement{},
							types.SingleLineComment{},
						},
					},
				},
			},
		}
		expected := []interface{}{
			types.UnorderedList{
				Items: []types.UnorderedListItem{
					{
						Elements: []interface{}{
							types.StringElement{},
						},
					},
				},
			},
		}
		Expect(filter(actual, allMatchers...)).To(Equal(expected))
	})

	It("should remove single line comment in an labeled list item", func() {
		actual := []interface{}{
			types.LabeledList{
				Items: []types.LabeledListItem{
					{
						Elements: []interface{}{
							types.StringElement{},
							types.SingleLineComment{},
						},
					},
				},
			},
		}
		expected := []interface{}{
			types.LabeledList{
				Items: []types.LabeledListItem{
					{
						Elements: []interface{}{
							types.StringElement{},
						},
					},
				},
			},
		}
		Expect(filter(actual, allMatchers...)).To(Equal(expected))
	})

})
