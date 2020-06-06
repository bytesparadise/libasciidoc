package testsupport_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("table of contents", func() {

	expected := types.TableOfContents{
		Sections: []types.ToCSection{
			{
				ID:       "title",
				Level:    1,
				Title:    "Title",
				Children: []types.ToCSection{},
			},
		},
	}

	It("should match", func() {
		// given
		actual := types.Document{
			ElementReferences: types.ElementReferences{},
			Footnotes:         []types.Footnote{},
			Elements: []interface{}{
				types.Section{
					Attributes: types.Attributes{
						types.AttrID: "title",
					},
					Level: 1,
					Title: []interface{}{
						types.StringElement{
							Content: "Title",
						},
					},
					Elements: []interface{}{},
				},
			},
		}
		// when
		result, err := testsupport.TableOfContents(actual)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(Equal(expected))
	})

	It("should not match", func() {
		// given
		actual := types.Document{
			Elements: []interface{}{},
		}
		// when
		result, err := testsupport.TableOfContents(actual)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).NotTo(Equal(expected))
	})

})
