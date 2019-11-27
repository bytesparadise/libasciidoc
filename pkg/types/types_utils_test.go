package types

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("convert to inline elements", func() {

	It("inline content without trailing spaces", func() {
		source := []interface{}{
			StringElement{Content: "hello"},
			StringElement{Content: "world"},
		}
		expected := InlineElements{
			StringElement{Content: "helloworld"},
		}
		// when
		result := MergeStringElements(source...)
		// then
		Expect(result).To(Equal(expected))
	})
	It("inline content with trailing spaces", func() {
		source := []interface{}{
			StringElement{Content: "hello, "},
			StringElement{Content: "world   "},
		}
		expected := InlineElements{
			StringElement{Content: "hello, world   "},
		}
		// when
		result := MergeStringElements(source...)
		// then
		Expect(result).To(Equal(expected))
	})
})

var _ = Describe("filter elements", func() {

	It("filter elements with all options", func() {
		// given
		source := []interface{}{
			BlankLine{},
			Preamble{},
			StringElement{
				Content: "foo",
			},
			Preamble{
				Elements: []interface{}{
					StringElement{
						Content: "bar",
					},
				},
			},
			[]interface{}{
				BlankLine{},
				StringElement{
					Content: "baz",
				},
			},
		}
		// when
		result := filterEmptyElements(source, filterBlankLine(), filterEmptyPreamble())
		// then
		expected := []interface{}{
			StringElement{
				Content: "foo",
			},
			Preamble{
				Elements: []interface{}{
					StringElement{
						Content: "bar",
					},
				},
			},
			StringElement{
				Content: "baz",
			},
		}
		Expect(result).To(Equal(expected))
	})
})
