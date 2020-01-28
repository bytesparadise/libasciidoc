package testsupport_test

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("document table of contents assertions", func() {

	preamble := types.Preamble{
		Elements: []interface{}{
			types.BlankLine{},
			types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: [][]interface{}{
					{
						types.StringElement{Content: "A short preamble"},
					},
				},
			},
			types.BlankLine{},
		},
	}
	section := types.Section{
		Level:      1,
		Attributes: types.ElementAttributes{},
		Title: []interface{}{
			types.StringElement{Content: "section 1"},
		},
		Elements: []interface{}{},
	}
	tableOfContents := types.TableOfContentsPlaceHolder{}

	actual := types.Document{
		Attributes: types.DocumentAttributes{
			types.AttrTableOfContents: "",
		},
		ElementReferences:  types.ElementReferences{}, // can leave empty for this test
		Footnotes:          types.Footnotes{},
		FootnoteReferences: types.FootnoteReferences{},
		Elements: []interface{}{
			preamble,
			section,
		},
	}
	expected := types.Document{
		Attributes: types.DocumentAttributes{
			types.AttrTableOfContents: "",
		},
		ElementReferences:  types.ElementReferences{},
		Footnotes:          types.Footnotes{},
		FootnoteReferences: types.FootnoteReferences{},
		Elements: []interface{}{
			tableOfContents,
			preamble,
			section,
		},
	}

	It("should match", func() {
		// given
		matcher := testsupport.HaveTableOfContentsPlaceHolder(expected)
		// when
		result, err := matcher.Match(actual)
		// then
		GinkgoT().Log(matcher.FailureMessage(actual))
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(BeTrue())
	})

	It("should not match", func() {
		// given
		matcher := testsupport.HaveTableOfContentsPlaceHolder(expected)
		// when
		result, err := matcher.Match(types.Document{})
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(BeFalse())
		Expect(matcher.FailureMessage(actual)).To(Equal(fmt.Sprintf("expected documents to match:\n%s", compare(types.Document{}, expected))))
		Expect(matcher.NegatedFailureMessage(actual)).To(Equal(fmt.Sprintf("expected documents not to match:\n%s", compare(types.Document{}, expected))))
	})

	It("should return error when invalid type is input", func() {
		// given
		matcher := testsupport.HaveTableOfContentsPlaceHolder(types.Document{})
		// when
		result, err := matcher.Match(1) // not a doc
		// then
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("HaveTableOfContents matcher expects a Document (actual: int)"))
		Expect(result).To(BeFalse())
	})
})
