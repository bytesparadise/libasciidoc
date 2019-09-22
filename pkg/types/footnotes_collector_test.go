package types_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("footnotes collector", func() {

	It("index footnotes without reference", func() {
		// given
		footnote1 := types.Footnote{
			ID: 0,
			Elements: types.InlineElements{
				types.StringElement{
					Content: "a note",
				},
			},
		}
		footnote2 := types.Footnote{
			ID: 1,
			Elements: types.InlineElements{
				types.StringElement{
					Content: "another note",
				},
			},
		}
		content := types.InlineElements{
			types.StringElement{
				Content: "foo",
			},
			footnote1,
			types.StringElement{
				Content: "bar",
			},
			footnote2,
		}
		c := types.NewFootnotesCollector()
		// when
		err := content.AcceptVisitor(c)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(c.Footnotes).To(HaveLen(2))
		Expect(c.Footnotes[0]).To(Equal(footnote1))
		Expect(c.Footnotes[1]).To(Equal(footnote2))
		Expect(c.FootnoteReferences).To(BeEmpty())
	})

	It("index footnotes with reference", func() {
		// given
		footnote1 := types.Footnote{
			Ref: "ref",
			Elements: types.InlineElements{
				types.StringElement{
					Content: "a note",
				},
			},
		}
		footnote2 := types.Footnote{
			Ref: "ref",
		}
		footnote3 := types.Footnote{
			Ref: "ref",
		}
		content := types.InlineElements{
			types.StringElement{
				Content: "foo",
			},
			footnote1,
			types.StringElement{
				Content: "bar",
			},
			footnote2,
			footnote3,
		}
		c := types.NewFootnotesCollector()
		// when
		err := content.AcceptVisitor(c)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(c.Footnotes).To(HaveLen(1))
		Expect(c.Footnotes[0]).To(Equal(footnote1))
		Expect(c.FootnoteReferences).To(HaveLen(1))
		Expect(c.FootnoteReferences["ref"]).To(Equal(footnote1))
	})

})
