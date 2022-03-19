package renderer_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("section numbering", func() {

	It("should number sections", func() {
		// given
		doc := &types.Document{
			Elements: []interface{}{
				&types.Section{
					Attributes: types.Attributes{
						types.AttrID: "_section_1",
					},
					Elements: []interface{}{
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_section_1a",
							},
							Elements: []interface{}{},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_section_1b",
							},
							Elements: []interface{}{},
						},
					},
				},
			},
		}
		// when
		n, err := renderer.NewSectionNumbers(doc)

		// then
		Expect(err).NotTo(HaveOccurred())
		Expect(n["_section_1"]).To(Equal("1"))
		Expect(n["_section_1a"]).To(Equal("1.1"))
		Expect(n["_section_1b"]).To(Equal("1.2"))
	})
})
