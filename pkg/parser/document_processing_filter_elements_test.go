package parser

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
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

	// It("should remove document attribute declaration", func() {
	// 	actual := []interface{}{
	// 		&types.AttributeDeclaration{},
	// 		&types.Paragraph{
	// 			Elements: []interface{}{
	// 				&types.StringElement{},
	// 			},
	// 		},
	// 	}
	// 	expected := []interface{}{
	// 		&types.Paragraph{
	// 			Elements: []interface{}{
	// 				&types.StringElement{},
	// 			},
	// 		},
	// 	}
	// 	Expect(filterComments(actual, allMatchers...)).To(Equal(expected))
	// })

	// It("should remove document attribute substitution", func() {
	// 	actual := []interface{}{
	// 		&types.AttributeSubstitution{},
	// 		&types.Paragraph{
	// 			Elements: []interface{}{
	// 				&types.StringElement{},
	// 			},
	// 		},
	// 	}
	// 	expected := []interface{}{
	// 		&types.Paragraph{
	// 			Elements: []interface{}{
	// 				&types.StringElement{},
	// 			},
	// 		},
	// 	}
	// 	Expect(filterComments(actual, allMatchers...)).To(Equal(expected))
	// })

	// It("should remove document attribute reset", func() {
	// 	actual := []interface{}{
	// 		&types.AttributeReset{},
	// 		&types.Paragraph{
	// 			Elements: []interface{}{
	// 				&types.StringElement{},
	// 			},
	// 		},
	// 	}
	// 	expected := []interface{}{
	// 		&types.Paragraph{
	// 			Elements: []interface{}{
	// 				&types.StringElement{},
	// 			},
	// 		},
	// 	}
	// 	Expect(filterComments(actual, allMatchers...)).To(Equal(expected))
	// })

	It("should remove comment block", func() {
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

	It("should remove single line comment as a block", func() {
		actual := []interface{}{
			&types.SingleLineComment{},
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
					&types.SingleLineComment{},
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
					&types.SingleLineComment{},
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
							&types.SingleLineComment{},
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
