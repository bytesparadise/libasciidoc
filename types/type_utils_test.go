package types

import (
	"github.com/stretchr/testify/require"

	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Normalizing String", func() {

	It("hello", func() {
		source := &InlineContent{
			Elements: []InlineElement{
				&StringElement{Content: "hello"},
			},
		}
		verify(GinkgoT(), "_hello", source)
	})

	It("héllo with an accent", func() {
		source := &InlineContent{
			Elements: []InlineElement{
				&StringElement{Content: "  héllo 1.2   and 3 Spaces"},
			},
		}
		verify(GinkgoT(), "_héllo_1_2_and_3_spaces", source)
	})

	It("a an accent and a swedish character", func() {
		source := &InlineContent{
			Elements: []InlineElement{
				&StringElement{Content: `A à ⌘`},
			},
		}
		verify(GinkgoT(), `_a_à`, source)
	})

	It("AŁA", func() {
		source := &InlineContent{
			Elements: []InlineElement{
				&StringElement{Content: `AŁA 0.1 ?`},
			},
		}
		verify(GinkgoT(), `_ała_0_1`, source)
	})

	It("it's  2 spaces, here !", func() {
		source := &InlineContent{
			Elements: []InlineElement{
				&StringElement{Content: `it's  2 spaces, here !`},
			},
		}
		verify(GinkgoT(), `_it_s_2_spaces_here`, source)
	})

	It("content with <strong> markup", func() {
		// == a section title, with *bold content*
		source := &InlineContent{
			Elements: []InlineElement{
				&StringElement{Content: "a section title, with"},
				&QuotedText{
					Kind: Bold,
					Elements: []InlineElement{
						&StringElement{Content: "bold content"},
					},
				},
			},
		}
		verify(GinkgoT(), `_a_section_title_with_strong_bold_content_strong`, source)
	})
})

func verify(t GinkgoTInterface, expected string, inlineContent *InlineContent) {
	t.Logf("Processing '%s'", inlineContent.String(0))
	result, err := ReplaceNonAlphanumerics(inlineContent, "_")
	require.Nil(t, err)
	t.Logf("Normalized result: '%s'", *result)
	assert.Equal(t, expected, *result)

}
