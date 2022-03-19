package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2" // nolint:golint
	. "github.com/onsi/gomega"    // nolint:golint
)

var _ = Describe("collect footnotes", func() {

	Context("in section titles", func() {

		var sectionWithFootnote1 *types.Section
		var sectionWithFootnoteRef1 *types.Section
		var sectionWithFootnote2 *types.Section
		var sectionWithFootnoteRef2 *types.Section
		var sectionWithoutFootnote *types.Section

		BeforeEach(func() {
			sectionWithFootnote1 = &types.Section{
				Title: []interface{}{
					&types.StringElement{
						Content: "cookies",
					},
					&types.Footnote{
						Ref: "", // without ref
						Elements: []interface{}{
							&types.StringElement{
								Content: "cookies",
							},
						},
					},
				},
			}
			sectionWithFootnoteRef1 = &types.Section{
				Title: []interface{}{
					&types.StringElement{
						Content: "cookies",
					},
					&types.FootnoteReference{
						ID: 1,
					},
				},
			}
			sectionWithFootnote2 = &types.Section{
				Title: []interface{}{
					&types.StringElement{
						Content: "pasta",
					},
					&types.Footnote{
						Ref: "pasta", // with ref
						Elements: []interface{}{
							&types.StringElement{
								Content: "pasta",
							},
						},
					},
				},
			}
			sectionWithFootnoteRef2 = &types.Section{
				Title: []interface{}{
					&types.StringElement{
						Content: "pasta",
					},
					&types.FootnoteReference{
						ID:  2,
						Ref: "pasta", // with ref
					},
				},
			}
			sectionWithoutFootnote = &types.Section{
				Title: []interface{}{
					&types.StringElement{
						Content: "coffee",
					},
				},
			}
		})

		It("no footnote", func() {
			// given
			c := make(chan types.DocumentFragment, 1)
			c <- types.DocumentFragment{
				Elements: []interface{}{
					sectionWithoutFootnote,
				},
			}
			close(c)
			footnotes := types.NewFootnotes()
			// when
			result := parser.CollectFootnotes(footnotes, make(<-chan interface{}), c)
			// then
			Expect(<-result).To(MatchDocumentFragment(types.DocumentFragment{
				Elements: []interface{}{
					sectionWithoutFootnote,
				},
			}))
			Expect(footnotes.Notes).To(BeEmpty())
		})

		It("single footnote", func() {
			// given
			c := make(chan types.DocumentFragment, 1)
			c <- types.DocumentFragment{
				Elements: []interface{}{
					sectionWithFootnote1,
					sectionWithoutFootnote,
				},
			}
			close(c)
			footnotes := types.NewFootnotes()
			// when
			result := parser.CollectFootnotes(footnotes, make(<-chan interface{}), c)
			// then
			Expect(<-result).To(MatchDocumentFragment(types.DocumentFragment{
				Elements: []interface{}{
					sectionWithFootnoteRef1,
					sectionWithoutFootnote,
				},
			}))
			Expect(footnotes.Notes).To(Equal([]*types.Footnote{
				{
					ID:  1,  // set
					Ref: "", // without ref
					Elements: []interface{}{
						&types.StringElement{
							Content: "cookies",
						},
					},
				},
			}))
		})

		It("multiple footnotes in same fragment", func() {
			// given
			c := make(chan types.DocumentFragment, 1)
			c <- types.DocumentFragment{
				Elements: []interface{}{
					sectionWithFootnote1,
					sectionWithoutFootnote,
					sectionWithFootnote2,
				},
			}
			close(c)
			footnotes := types.NewFootnotes()
			// when
			result := parser.CollectFootnotes(footnotes, make(<-chan interface{}), c)
			// then
			Expect(<-result).To(MatchDocumentFragment(types.DocumentFragment{
				Elements: []interface{}{
					sectionWithFootnoteRef1,
					sectionWithoutFootnote,
					sectionWithFootnoteRef2,
				},
			}))
			Expect(footnotes.Notes).To(Equal([]*types.Footnote{
				{
					ID:  1,  // set
					Ref: "", // without ref
					Elements: []interface{}{
						&types.StringElement{
							Content: "cookies",
						},
					},
				},
				{
					ID:  2,       // set
					Ref: "pasta", // with ref
					Elements: []interface{}{
						&types.StringElement{
							Content: "pasta",
						},
					},
				},
			}))
		})

		It("multiple footnotes in separate fragments", func() {
			// given
			c := make(chan types.DocumentFragment, 2)
			c <- types.DocumentFragment{
				Elements: []interface{}{
					sectionWithoutFootnote,
					sectionWithFootnote1,
					sectionWithoutFootnote,
				},
			}
			c <- types.DocumentFragment{
				Elements: []interface{}{
					sectionWithoutFootnote,
					sectionWithFootnote2,
					sectionWithoutFootnote,
				},
			}
			close(c)
			footnotes := types.NewFootnotes()
			// when
			result := parser.CollectFootnotes(footnotes, make(<-chan interface{}), c)
			// then
			Expect(<-result).To(MatchDocumentFragment(types.DocumentFragment{
				Elements: []interface{}{
					sectionWithoutFootnote,
					sectionWithFootnoteRef1,
					sectionWithoutFootnote,
				},
			}))
			Expect(<-result).To(MatchDocumentFragment(types.DocumentFragment{
				Elements: []interface{}{
					sectionWithoutFootnote,
					sectionWithFootnoteRef2,
					sectionWithoutFootnote,
				},
			}))
			Expect(footnotes.Notes).To(Equal([]*types.Footnote{
				{
					ID:  1,  // set
					Ref: "", // without ref
					Elements: []interface{}{
						&types.StringElement{
							Content: "cookies",
						},
					},
				},
				{
					ID:  2,       // set
					Ref: "pasta", // with ref
					Elements: []interface{}{
						&types.StringElement{
							Content: "pasta",
						},
					},
				},
			}))
		})
	})

	Context("in paragraphs", func() {

		var paragraphWithFootnote1 *types.Paragraph
		var paragraphWithFootnoteRef1 *types.Paragraph
		var paragraphWithFootnote2 *types.Paragraph
		var paragraphWithFootnoteRef2 *types.Paragraph
		var paragraphWithoutFootnote *types.Paragraph

		BeforeEach(func() {
			paragraphWithFootnote1 = &types.Paragraph{
				Elements: []interface{}{
					&types.StringElement{
						Content: "cookies",
					},
					&types.Footnote{
						Ref: "", // without ref
						Elements: []interface{}{
							&types.StringElement{
								Content: "cookies",
							},
						},
					},
				},
			}
			paragraphWithFootnoteRef1 = &types.Paragraph{
				Elements: []interface{}{
					&types.StringElement{
						Content: "cookies",
					},
					&types.FootnoteReference{
						ID: 1,
					},
				},
			}
			paragraphWithFootnote2 = &types.Paragraph{
				Elements: []interface{}{
					&types.StringElement{
						Content: "pasta",
					},
					&types.Footnote{
						Ref: "pasta", // with ref
						Elements: []interface{}{
							&types.StringElement{
								Content: "pasta",
							},
						},
					},
				},
			}
			paragraphWithFootnoteRef2 = &types.Paragraph{
				Elements: []interface{}{
					&types.StringElement{
						Content: "pasta",
					},
					&types.FootnoteReference{
						ID:  2,
						Ref: "pasta", // with ref
					},
				},
			}
			paragraphWithoutFootnote = &types.Paragraph{
				Elements: []interface{}{
					&types.StringElement{
						Content: "coffee",
					},
				},
			}
		})

		It("no footnote", func() {
			// given
			c := make(chan types.DocumentFragment, 1)
			c <- types.DocumentFragment{
				Elements: []interface{}{
					paragraphWithoutFootnote,
				},
			}
			close(c)
			footnotes := types.NewFootnotes()
			// when
			result := parser.CollectFootnotes(footnotes, make(<-chan interface{}), c)
			// then
			Expect(<-result).To(MatchDocumentFragment(types.DocumentFragment{
				Elements: []interface{}{
					paragraphWithoutFootnote,
				},
			}))
			Expect(footnotes.Notes).To(BeEmpty())
		})

		It("single footnote", func() {
			// given
			c := make(chan types.DocumentFragment, 1)
			c <- types.DocumentFragment{
				Elements: []interface{}{
					paragraphWithFootnote1,
					paragraphWithoutFootnote,
				},
			}
			close(c)
			footnotes := types.NewFootnotes()
			// when
			result := parser.CollectFootnotes(footnotes, make(<-chan interface{}), c)
			// then
			Expect(<-result).To(MatchDocumentFragment(types.DocumentFragment{
				Elements: []interface{}{
					paragraphWithFootnoteRef1,
					paragraphWithoutFootnote,
				},
			}))
			Expect(footnotes.Notes).To(Equal([]*types.Footnote{
				{
					ID:  1,  // set
					Ref: "", // without ref
					Elements: []interface{}{
						&types.StringElement{
							Content: "cookies",
						},
					},
				},
			}))
		})

		It("multiple footnotes in same fragment", func() {
			// given
			c := make(chan types.DocumentFragment, 1)
			c <- types.DocumentFragment{
				Elements: []interface{}{
					paragraphWithFootnote1,
					paragraphWithoutFootnote,
					paragraphWithFootnote2,
				},
			}
			close(c)
			footnotes := types.NewFootnotes()
			// when
			result := parser.CollectFootnotes(footnotes, make(<-chan interface{}), c)
			// then
			Expect(<-result).To(MatchDocumentFragment(types.DocumentFragment{
				Elements: []interface{}{
					paragraphWithFootnoteRef1,
					paragraphWithoutFootnote,
					paragraphWithFootnoteRef2,
				},
			}))
			Expect(footnotes.Notes).To(Equal([]*types.Footnote{
				{
					ID:  1,  // set
					Ref: "", // without ref
					Elements: []interface{}{
						&types.StringElement{
							Content: "cookies",
						},
					},
				},
				{
					ID:  2,       // set
					Ref: "pasta", // with ref
					Elements: []interface{}{
						&types.StringElement{
							Content: "pasta",
						},
					},
				},
			}))
		})

		It("multiple footnotes in separate fragments", func() {
			// given
			c := make(chan types.DocumentFragment, 2)
			c <- types.DocumentFragment{
				Elements: []interface{}{
					paragraphWithoutFootnote,
					paragraphWithFootnote1,
					paragraphWithoutFootnote,
				},
			}
			c <- types.DocumentFragment{
				Elements: []interface{}{
					paragraphWithoutFootnote,
					paragraphWithFootnote2,
					paragraphWithoutFootnote,
				},
			}
			close(c)
			footnotes := types.NewFootnotes()
			// when
			result := parser.CollectFootnotes(footnotes, make(<-chan interface{}), c)
			// then
			Expect(<-result).To(MatchDocumentFragment(types.DocumentFragment{
				Elements: []interface{}{
					paragraphWithoutFootnote,
					paragraphWithFootnoteRef1,
					paragraphWithoutFootnote,
				},
			}))
			Expect(<-result).To(MatchDocumentFragment(types.DocumentFragment{
				Elements: []interface{}{
					paragraphWithoutFootnote,
					paragraphWithFootnoteRef2,
					paragraphWithoutFootnote,
				},
			}))
			Expect(footnotes.Notes).To(Equal([]*types.Footnote{
				{
					ID:  1,  // set
					Ref: "", // without ref
					Elements: []interface{}{
						&types.StringElement{
							Content: "cookies",
						},
					},
				},
				{
					ID:  2,       // set
					Ref: "pasta", // with ref
					Elements: []interface{}{
						&types.StringElement{
							Content: "pasta",
						},
					},
				},
			}))
		})
	})

})
