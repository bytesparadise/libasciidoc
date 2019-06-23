package types

import (
	"github.com/davecgh/go-spew/spew"
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var _ = Describe("convert to inline elements", func() {

	It("inline content without trailing spaces", func() {
		source := []interface{}{
			&StringElement{Content: "hello"},
			&StringElement{Content: "world"},
		}
		expected := InlineElements{
			&StringElement{Content: "helloworld"},
		}
		// when
		result := mergeElements(source...)
		// then
		assert.Equal(GinkgoT(), expected, result)
	})
	It("inline content with trailing spaces", func() {
		source := []interface{}{
			&StringElement{Content: "hello, "},
			&StringElement{Content: "world   "},
		}
		expected := InlineElements{
			&StringElement{Content: "hello, world   "},
		}
		// when
		result := mergeElements(source...)
		// then
		assert.Equal(GinkgoT(), expected, result)
	})
})

var _ = Describe("normalizing string", func() {

	It("hello", func() {
		source := InlineElements{
			&StringElement{Content: "hello"},
		}
		verify(GinkgoT(), "hello", source)
	})

	It("héllo with an accent", func() {
		source := InlineElements{
			&StringElement{Content: "  héllo 1.2   and 3 Spaces"},
		}
		verify(GinkgoT(), "héllo_1_2_and_3_spaces", source)
	})

	It("a an accent and a swedish character", func() {
		source := InlineElements{
			&StringElement{Content: `A à ⌘`},
		}
		verify(GinkgoT(), `a_à`, source)
	})

	It("AŁA", func() {
		source := InlineElements{
			&StringElement{Content: `AŁA 0.1 ?`},
		}
		verify(GinkgoT(), `ała_0_1`, source)
	})

	It("it's  2 spaces, here !", func() {
		source := InlineElements{
			&StringElement{Content: `it's  2 spaces, here !`},
		}
		verify(GinkgoT(), `it_s_2_spaces_here`, source)
	})

	It("content with <strong> markup", func() {
		// == a section title, with *bold content*
		source := InlineElements{
			&StringElement{Content: "a section title, with"},
			QuotedText{
				Kind: Bold,
				Elements: []interface{}{
					&StringElement{Content: "bold content"},
				},
			},
		}
		verify(GinkgoT(), `a_section_title_with_bold_content`, source)
	})
})

func verify(t GinkgoTInterface, expected string, inlineContent InlineElements) {
	t.Logf("Processing '%s'", spew.Sprint(inlineContent))
	result, err := replaceNonAlphanumerics(inlineContent, "_")
	require.Nil(t, err)
	t.Logf("Normalized result: '%s'", result)
	assert.Equal(t, expected, result)

}

var _ = Describe("filter elements", func() {

	It("filter elements with all options", func() {
		// given
		source := []interface{}{
			BlankLine{},
			Preamble{},
			&StringElement{
				Content: "foo",
			},
			Preamble{
				Elements: []interface{}{
					&StringElement{
						Content: "bar",
					},
				},
			},
			[]interface{}{
				BlankLine{},
				&StringElement{
					Content: "baz",
				},
			},
		}
		// when
		result := filterEmptyElements(source, filterBlankLine(), filterEmptyPreamble())
		// then
		expected := []interface{}{
			&StringElement{
				Content: "foo",
			},
			Preamble{
				Elements: []interface{}{
					&StringElement{
						Content: "bar",
					},
				},
			},
			&StringElement{
				Content: "baz",
			},
		}
		assert.Equal(GinkgoT(), expected, result)
	})
})
