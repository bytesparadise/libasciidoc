package types_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("normalizing string", func() {

	It("hello", func() {
		source := types.InlineElements{
			types.StringElement{Content: "hello"},
		}
		Expect(source).To(EqualWithoutNonAlphanumeric("hello"))
	})

	It("héllo with an accent", func() {
		source := types.InlineElements{
			types.StringElement{Content: "  héllo 1.2   and 3 Spaces"},
		}
		Expect(source).To(EqualWithoutNonAlphanumeric("héllo_1_2_and_3_spaces"))
	})

	It("a an accent and a swedish character", func() {
		source := types.InlineElements{
			types.StringElement{Content: `A à ⌘`},
		}
		Expect(source).To(EqualWithoutNonAlphanumeric("a_à"))
	})

	It("AŁA", func() {
		source := types.InlineElements{
			types.StringElement{Content: `AŁA 0.1 ?`},
		}
		Expect(source).To(EqualWithoutNonAlphanumeric("ała_0_1"))
	})

	It("it's  2 spaces, here !", func() {
		source := types.InlineElements{
			types.StringElement{Content: `it's  2 spaces, here !`},
		}
		Expect(source).To(EqualWithoutNonAlphanumeric("it_s_2_spaces_here"))
	})

	It("content with <strong> markup", func() {
		// == a section title, with *bold content*
		source := types.InlineElements{
			types.StringElement{Content: "a section title, with"},
			types.QuotedText{
				Kind: types.Bold,
				Elements: []interface{}{
					types.StringElement{Content: "bold content"},
				},
			},
		}
		Expect(source).To(EqualWithoutNonAlphanumeric("a_section_title_with_bold_content"))
	})

	It("content with link", func() {
		// == a section title, with *bold content*
		source := types.InlineElements{
			types.StringElement{Content: "link to "},
			types.InlineLink{
				Attributes: types.ElementAttributes{},
				Location: types.Location{
					Elements: []interface{}{
						types.StringElement{
							Content: "https://foo.bar",
						},
					},
				},
			},
		}
		Expect(source).To(EqualWithoutNonAlphanumeric("link_to_https_foo_bar")) // asciidoctor will return `_link_to_https_foo_bar`
	})
})
