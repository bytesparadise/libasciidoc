package types_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("normalizing string", func() {

	It("hello", func() {
		source := []interface{}{
			&types.StringElement{Content: "hello"},
		}
		Expect(types.ReplaceNonAlphanumerics(source, "_")).To(Equal("hello"))
	})

	It("héllo with an accent", func() {
		source := []interface{}{
			&types.StringElement{Content: "  héllo 1.2   and 3 Spaces"},
		}
		Expect(types.ReplaceNonAlphanumerics(source, "_")).To(Equal("héllo_1_2_and_3_spaces"))
	})

	It("a an accent and a swedish character", func() {
		source := []interface{}{
			&types.StringElement{Content: `A à ⌘`},
		}
		Expect(types.ReplaceNonAlphanumerics(source, "_")).To(Equal("a_à"))
	})

	It("AŁA", func() {
		source := []interface{}{
			&types.StringElement{Content: `AŁA 0.1 ?`},
		}
		Expect(types.ReplaceNonAlphanumerics(source, "_")).To(Equal("ała_0_1"))
	})

	It("it's  2 spaces, here !", func() {
		source := []interface{}{
			&types.StringElement{Content: `it's  2 spaces, here !`},
		}
		Expect(types.ReplaceNonAlphanumerics(source, "_")).To(Equal("it_s_2_spaces_here"))
	})

	It("content with <strong> markup", func() {
		// == a section title, with *bold content*
		source := []interface{}{
			&types.StringElement{Content: "a section title, with"},
			&types.QuotedText{
				Kind: types.SingleQuoteBold,
				Elements: []interface{}{
					&types.StringElement{Content: "bold content"},
				},
			},
		}
		Expect(types.ReplaceNonAlphanumerics(source, "_")).To(Equal("a_section_title_with_bold_content"))
	})

	It("content with link", func() {
		// == a section title, with *bold content*
		source := []interface{}{
			&types.StringElement{Content: "link to "},
			&types.InlineLink{
				Attributes: types.Attributes{},
				Location: &types.Location{
					Scheme: "https://",
					Path: []interface{}{
						&types.StringElement{
							Content: "foo.bar",
						},
					},
				},
			},
		}
		Expect(types.ReplaceNonAlphanumerics(source, "_")).To(Equal("link_to_httpsfoo_bar")) // asciidoctor will return `_link_to_httpsfoo_bar`
	})
})
