package types_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("section numbering", func() {

	It("should always number sections - explicitly in document header", func() {
		// given
		doc := &types.Document{
			Elements: []interface{}{
				&types.DocumentHeader{
					Elements: []interface{}{
						&types.AttributeDeclaration{
							Name: types.AttrSectionNumbering,
						},
					},
				},
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
		n, err := doc.SectionNumbers()

		// then
		Expect(err).NotTo(HaveOccurred())
		Expect(n["_section_1"]).To(Equal("1"))
		Expect(n["_section_1a"]).To(Equal("1.1"))
		Expect(n["_section_1b"]).To(Equal("1.2"))
	})

	It("should always number sections - explicitly in document body", func() {
		// given
		doc := &types.Document{
			Elements: []interface{}{
				&types.AttributeDeclaration{
					Name: types.AttrSectionNumbering,
				},
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
		n, err := doc.SectionNumbers()

		// then
		Expect(err).NotTo(HaveOccurred())
		Expect(n["_section_1"]).To(Equal("1"))
		Expect(n["_section_1a"]).To(Equal("1.1"))
		Expect(n["_section_1b"]).To(Equal("1.2"))
	})

	It("should never number sections - explicitly", func() {
		// given
		doc := &types.Document{
			Elements: []interface{}{
				&types.AttributeReset{
					Name: types.AttrSectionNumbering,
				},
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
		n, err := doc.SectionNumbers()

		// then
		Expect(err).NotTo(HaveOccurred())
		Expect(n).To(BeEmpty())
	})

	It("should never number sections - by default", func() {
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
		n, err := doc.SectionNumbers()

		// then
		Expect(err).NotTo(HaveOccurred())
		Expect(n).To(BeEmpty())
	})

	It("should number sections when enabled - case 1", func() {
		// given
		doc := &types.Document{
			Elements: []interface{}{
				&types.AttributeReset{
					Name: types.AttrSectionNumbering,
				},
				&types.Section{
					Attributes: types.Attributes{
						types.AttrID: "_disclaimer",
					},
				},
				&types.AttributeDeclaration{
					Name: types.AttrSectionNumbering,
				},
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
		n, err := doc.SectionNumbers()

		// then
		Expect(err).NotTo(HaveOccurred())
		Expect(n["_disclaimer"]).To(BeEmpty())
		Expect(n["_section_1"]).To(Equal("1"))
		Expect(n["_section_1a"]).To(Equal("1.1"))
		Expect(n["_section_1b"]).To(Equal("1.2"))
	})
})
