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
			Elements: []interface{}{
				types.StringElement{
					Content: "a note",
				},
			},
		}
		footnote2 := types.Footnote{
			ID: 1,
			Elements: []interface{}{
				types.StringElement{
					Content: "another note",
				},
			},
		}
		content := []interface{}{
			types.StringElement{
				Content: "foo",
			},
			footnote1,
			types.StringElement{
				Content: "bar",
			},
			footnote2,
		}
		// when
		notes, refs, err := types.FindFootnotes(content)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(notes).To(HaveLen(2))
		Expect(notes[0]).To(Equal(footnote1))
		Expect(notes[1]).To(Equal(footnote2))
		Expect(refs).To(BeEmpty())
	})

	It("index footnotes with reference", func() {
		// given
		footnote1 := types.Footnote{
			Ref: "ref",
			Elements: []interface{}{
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
		content := []interface{}{
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
		// when
		notes, refs, err := types.FindFootnotes(content)
		// then
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(notes).To(HaveLen(1))
		Expect(notes[0]).To(Equal(footnote1))
		Expect(refs).To(HaveLen(1))
		Expect(refs["ref"]).To(Equal(footnote1))
	})

})
