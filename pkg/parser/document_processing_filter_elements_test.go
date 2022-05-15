package parser

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("element filters", func() {

	It("should retain content in a delimited block and paragraph", func() {
		actual := []interface{}{
			&types.DelimitedBlock{
				Kind: types.Listing,
				Elements: []interface{}{
					&types.StringElement{},
				},
			},
			&types.BlankLine{},
			&types.Paragraph{
				Elements: []interface{}{
					&types.StringElement{},
				},
			},
		}
		expected := []interface{}{
			&types.DelimitedBlock{
				Kind: types.Listing,
				Elements: []interface{}{
					&types.StringElement{},
				},
			},
			&types.Paragraph{
				Elements: []interface{}{
					&types.StringElement{},
				},
			},
		}
		Expect(doFilterOut(actual, allMatchers...)).To(Equal(expected))
	})

	It("should not retain blanklines in a delimited blocks", func() {
		actual := []interface{}{
			&types.DelimitedBlock{
				Kind: types.Listing,
				Elements: []interface{}{
					&types.StringElement{},
					&types.BlankLine{},
					&types.StringElement{},
				},
			},
		}
		expected := []interface{}{
			&types.DelimitedBlock{
				Kind: types.Listing,
				Elements: []interface{}{
					&types.StringElement{},
					&types.StringElement{},
				},
			},
		}
		Expect(doFilterOut(actual, allMatchers...)).To(Equal(expected))
	})

	It("should remove comment block at root", func() {
		actual := []interface{}{
			&types.DelimitedBlock{
				Kind: types.Comment,
			},
			&types.Paragraph{
				Elements: []interface{}{
					&types.StringElement{},
				},
			},
		}
		expected := []interface{}{
			&types.Paragraph{
				Elements: []interface{}{
					&types.StringElement{},
				},
			},
		}
		Expect(doFilterOut(actual, allMatchers...)).To(Equal(expected))
	})

	It("should remove comment blocks in document header", func() {
		actual := []interface{}{
			&types.DocumentHeader{
				Elements: []interface{}{
					&types.DelimitedBlock{
						Kind: types.Comment,
					},
					&types.AttributeDeclaration{
						Name:  "cookie",
						Value: "yummy",
					},
					&types.DelimitedBlock{
						Kind: types.Comment,
					},
				},
			},
		}
		expected := []interface{}{
			&types.DocumentHeader{
				Elements: []interface{}{
					&types.AttributeDeclaration{
						Name:  "cookie",
						Value: "yummy",
					},
				},
			},
		}
		Expect(doFilterOut(actual, allMatchers...)).To(Equal(expected))
	})

	It("should remove single line comments in document header", func() {
		actual := []interface{}{
			&types.DocumentHeader{
				Elements: []interface{}{
					&types.SinglelineComment{},
					&types.AttributeDeclaration{
						Name:  "cookie",
						Value: "yummy",
					},
					&types.SinglelineComment{},
				},
			},
		}
		expected := []interface{}{
			&types.DocumentHeader{
				Elements: []interface{}{
					&types.AttributeDeclaration{
						Name:  "cookie",
						Value: "yummy",
					},
				},
			},
		}
		Expect(doFilterOut(actual, allMatchers...)).To(Equal(expected))
	})

	It("should remove single line comment as a block", func() {
		actual := []interface{}{
			&types.SinglelineComment{},
			&types.Paragraph{
				Elements: []interface{}{
					&types.StringElement{},
				},
			},
		}
		expected := []interface{}{
			&types.Paragraph{
				Elements: []interface{}{
					&types.StringElement{},
				},
			},
		}
		Expect(doFilterOut(actual, allMatchers...)).To(Equal(expected))
	})

	It("should remove single line comment in a paragraph", func() {
		actual := []interface{}{
			&types.Paragraph{
				Elements: []interface{}{
					&types.StringElement{},
					&types.SinglelineComment{},
				},
			},
		}
		expected := []interface{}{
			&types.Paragraph{
				Elements: []interface{}{
					&types.StringElement{},
				},
			},
		}
		Expect(doFilterOut(actual, allMatchers...)).To(Equal(expected))
	})

	It("should retain paragraph with single line comment only", func() {
		actual := []interface{}{
			&types.Paragraph{
				Elements: []interface{}{
					&types.SinglelineComment{},
				},
			},
		}
		expected := []interface{}{
			&types.Paragraph{},
		}
		Expect(doFilterOut(actual, allMatchers...)).To(Equal(expected))
	})

	It("should retain paragraph with empty content", func() {
		actual := []interface{}{
			&types.Paragraph{},
		}
		expected := []interface{}{
			&types.Paragraph{},
		}
		Expect(doFilterOut(actual, allMatchers...)).To(Equal(expected))
	})

	It("should remove single line comment in an ordered list item", func() {
		actual := []interface{}{
			&types.List{
				Kind: types.OrderedListKind,
				Elements: []types.ListElement{
					&types.OrderedListElement{
						Elements: []interface{}{
							&types.StringElement{},
							&types.SinglelineComment{},
						},
					},
				},
			},
		}
		expected := []interface{}{
			&types.List{
				Kind: types.OrderedListKind,
				Elements: []types.ListElement{
					&types.OrderedListElement{
						Elements: []interface{}{
							&types.StringElement{},
						},
					},
				},
			},
		}
		Expect(doFilterOut(actual, allMatchers...)).To(Equal(expected))
	})
})
