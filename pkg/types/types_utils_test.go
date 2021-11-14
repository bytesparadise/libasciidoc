package types_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"                  //nolint golint
	. "github.com/onsi/ginkgo/extensions/table" //nolint golint
	. "github.com/onsi/gomega"                  //nolint golint
)

var _ = Describe("convert to inline elements", func() {

	It("inline content without trailing spaces", func() {
		source := []interface{}{
			&types.StringElement{Content: "hello"},
			&types.StringElement{Content: "world"},
		}
		expected := []interface{}{
			&types.StringElement{Content: "helloworld"},
		}
		Expect(types.Merge(source...)).To(Equal(expected))
	})

	It("inline content with trailing spaces", func() {
		source := []interface{}{
			&types.StringElement{Content: "hello, "},
			&types.StringElement{Content: "world   "},
		}
		expected := []interface{}{
			&types.StringElement{Content: "hello, world   "},
		}
		Expect(types.Merge(source...)).To(Equal(expected))
	})
})

// var _ = DescribeTable("TrimLeft",

// 	func(source, expected []interface{}) {

// 	},
// 	Entry("empty slice",
// 		[]interface{}{},
// 		[]interface{}{}),

// 	Entry("valid slice",
// 		[]interface{}{
// 			&types.StringElement{
// 				Content: "  cookies",
// 			},
// 			&types.StringElement{
// 				Content: "  pasta",
// 			},
// 		},
// 		[]interface{}{
// 			&types.StringElement{
// 				Content: "cookies", // trimmed
// 			},
// 			&types.StringElement{
// 				Content: "  pasta",
// 			},
// 		}),

// 	Entry("noop slice",
// 		[]interface{}{
// 			&types.SpecialCharacter{
// 				Name: ">",
// 			},
// 			&types.StringElement{
// 				Content: "  cookies",
// 			},
// 			&types.StringElement{
// 				Content: "  pasta",
// 			},
// 		},
// 		[]interface{}{
// 			&types.SpecialCharacter{
// 				Name: ">",
// 			},
// 			&types.StringElement{
// 				Content: "  cookies", // not trimmed
// 			},
// 			&types.StringElement{
// 				Content: "  pasta",
// 			},
// 		}),
// )

var _ = DescribeTable("split elements per line",
	func(elements []interface{}, expected [][]interface{}) {
		result := types.SplitElementsPerLine(elements)
		Expect(result).To(Equal(expected))

	},
	Entry("empty elements",
		[]interface{}{},
		[][]interface{}{}),

	Entry("single line",
		[]interface{}{
			&types.StringElement{
				Content: "cookie",
			},
			&types.Callout{
				Ref: 1,
			},
		},
		[][]interface{}{
			{
				&types.StringElement{
					Content: "cookie",
				},
				&types.Callout{
					Ref: 1,
				},
			},
		}),

	Entry("2 lines without callouts",
		[]interface{}{
			&types.StringElement{
				Content: "cookie",
			},
			&types.Callout{
				Ref: 1,
			},
			&types.StringElement{
				Content: "\npasta",
			},
			&types.Callout{
				Ref: 2,
			},
		},
		[][]interface{}{
			{
				&types.StringElement{
					Content: "cookie",
				},
				&types.Callout{
					Ref: 1,
				},
			},
			{
				&types.StringElement{
					Content: "pasta",
				},
				&types.Callout{
					Ref: 2,
				},
			},
		}),

	Entry("3 lines without callouts",
		[]interface{}{
			&types.StringElement{
				Content: "cookie\npasta\nchocolate",
			},
		},
		[][]interface{}{
			{
				&types.StringElement{
					Content: "cookie",
				},
			},
			{
				&types.StringElement{
					Content: "pasta",
				},
			},
			{
				&types.StringElement{
					Content: "chocolate",
				},
			},
		}),

	Entry("3 lines without callouts",
		[]interface{}{
			&types.StringElement{
				Content: "cookie",
			},
			&types.Callout{
				Ref: 1,
			},
			&types.StringElement{
				Content: "here\npasta",
			},
			&types.Callout{
				Ref: 2,
			},
			&types.StringElement{
				Content: "also\nchocolate",
			},
			&types.Callout{
				Ref: 3,
			},
		},
		[][]interface{}{
			{
				&types.StringElement{
					Content: "cookie",
				},
				&types.Callout{
					Ref: 1,
				},
				&types.StringElement{
					Content: "here",
				},
			},
			{
				&types.StringElement{
					Content: "pasta",
				},
				&types.Callout{
					Ref: 2,
				},
				&types.StringElement{
					Content: "also",
				},
			},
			{
				&types.StringElement{
					Content: "chocolate",
				},
				&types.Callout{
					Ref: 3,
				},
			},
		}),
)

var _ = DescribeTable("insert element in slice",

	func(elements []interface{}, element interface{}, index int, expected []interface{}) {
		result := types.InsertAt(elements, element, index)
		Expect(result).To(Equal(expected))
	},

	Entry("empty elements",
		[]interface{}{},
		&types.TableOfContents{},
		0,
		[]interface{}{
			&types.TableOfContents{},
		}),

	Entry("insert after preamble elements",
		[]interface{}{
			&types.Section{},
			&types.Preamble{},
			&types.Section{},
		},
		&types.TableOfContents{},
		2,
		[]interface{}{
			&types.Section{},
			&types.Preamble{},
			&types.TableOfContents{},
			&types.Section{},
		}),
)
