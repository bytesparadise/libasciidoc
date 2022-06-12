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
						types.AttrID: "_getting_started",
					},
					Elements: []interface{}{
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_introduction",
							},
							Elements: []interface{}{},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_download_and_install",
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
		Expect(n["_getting_started"]).To(Equal("1"))
		Expect(n["_introduction"]).To(Equal("1.1"))
		Expect(n["_download_and_install"]).To(Equal("1.2"))
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
						types.AttrID: "_getting_started",
					},
					Elements: []interface{}{
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_introduction",
							},
							Elements: []interface{}{},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_download_and_install",
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
		Expect(n["_getting_started"]).To(Equal("1"))
		Expect(n["_introduction"]).To(Equal("1.1"))
		Expect(n["_download_and_install"]).To(Equal("1.2"))
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
						types.AttrID: "_getting_started",
					},
					Elements: []interface{}{
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_introduction",
							},
							Elements: []interface{}{},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_download_and_install",
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
						types.AttrID: "_getting_started",
					},
					Elements: []interface{}{
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_introduction",
							},
							Elements: []interface{}{},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_download_and_install",
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
						types.AttrID: "_getting_started",
					},
					Elements: []interface{}{
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_introduction",
							},
							Elements: []interface{}{},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_download_and_install",
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
		Expect(n["_getting_started"]).To(Equal("1"))
		Expect(n["_introduction"]).To(Equal("1.1"))
		Expect(n["_download_and_install"]).To(Equal("1.2"))
	})

	It("should number sections when enabled - case 2", func() {
		// given
		doc := &types.Document{
			Elements: []interface{}{
				&types.FrontMatter{
					Attributes: map[string]interface{}{
						"draft": true,
					},
				},
				&types.DocumentHeader{
					Title: []interface{}{},
				},
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
						types.AttrID: "_getting_started",
					},
					Elements: []interface{}{
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_introduction",
							},
							Elements: []interface{}{},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_download_and_install",
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
		Expect(n["_getting_started"]).To(Equal("1"))
		Expect(n["_introduction"]).To(Equal("1.1"))
		Expect(n["_download_and_install"]).To(Equal("1.2"))
	})
})
