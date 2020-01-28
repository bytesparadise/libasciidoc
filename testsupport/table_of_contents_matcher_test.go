package testsupport_test

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("metadata table of contents matcher", func() {

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
			Attributes:         types.DocumentAttributes{},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.Section{
					Attributes: types.ElementAttributes{
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
		matcher := testsupport.HaveTableOfContents(expected)
		// when
		result, err := matcher.Match(actual)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(BeTrue())
	})

	It("should not match", func() {
		// given
		actual := types.Document{
			Elements: []interface{}{},
		}
		matcher := testsupport.HaveTableOfContents(expected)
		// when
		result, err := matcher.Match(actual)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(BeFalse())
		// also verify the messages
		Expect(matcher.FailureMessage(actual)).To(Equal(fmt.Sprintf("expected table of contents to match:\n%s", compare(types.TableOfContents{Sections: []types.ToCSection{}}, expected))))
		Expect(matcher.NegatedFailureMessage(actual)).To(Equal(fmt.Sprintf("expected table of contents not to match:\n%s", compare(types.TableOfContents{Sections: []types.ToCSection{}}, expected))))
	})

	It("should return error when invalid type is input", func() {
		// given
		matcher := testsupport.HaveTableOfContents(types.TableOfContents{})
		// when
		result, err := matcher.Match(1) // not a string
		// then
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("HaveTableOfContents matcher expects a Document (actual: int)"))
		Expect(result).To(BeFalse())
	})
})
