package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" // nolint:golint
	. "github.com/onsi/gomega" // nolint:golintt
)

var _ = Describe("blank lines", func() {

	Context("in raw documents", func() {

		It("blank line between 2 paragraphs", func() {
			source := `first paragraph
 
second paragraph`
			expected := []types.DocumentFragment{
				{
					Position: types.Position{
						Start: 0,
						End:   16,
					},
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								types.RawLine("first paragraph"),
							},
						},
					},
				},
				{
					Position: types.Position{
						Start: 16,
						End:   18,
					},
					Elements: []interface{}{
						&types.BlankLine{},
					},
				},
				{
					Position: types.Position{
						Start: 18,
						End:   34,
					},
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								types.RawLine("second paragraph"),
							},
						},
					},
				},
			}
			Expect(ParseDocumentFragments(source)).To(MatchDocumentFragmentGroups(expected))
		})

		It("blank line with spaces and tabs between 2 paragraphs and after second paragraph", func() {
			source := `first paragraph
		 

		
second paragraph
`
			expected := []types.DocumentFragment{
				{
					Position: types.Position{
						Start: 0,
						End:   16,
					},
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								types.RawLine("first paragraph"),
							},
						},
					},
				},
				{
					Position: types.Position{
						Start: 16,
						End:   20,
					},
					Elements: []interface{}{
						&types.BlankLine{},
					},
				},
				{
					Position: types.Position{
						Start: 20,
						End:   21,
					},
					Elements: []interface{}{
						&types.BlankLine{},
					},
				},
				{
					Position: types.Position{
						Start: 21,
						End:   24,
					},
					Elements: []interface{}{
						&types.BlankLine{},
					},
				},
				{
					Position: types.Position{
						Start: 24,
						End:   41,
					},
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								types.RawLine("second paragraph"),
							},
						},
					},
				},
			}
			Expect(ParseDocumentFragments(source)).To(MatchDocumentFragmentGroups(expected))
		})

		It("blank line with attributes", func() {
			source := `.ignored
 
`
			expected := []types.DocumentFragment{
				{
					Position: types.Position{
						Start: 0,
						End:   11,
					},
					Elements: []interface{}{
						&types.BlankLine{},
					},
				},
			}
			Expect(ParseDocumentFragments(source)).To(MatchDocumentFragmentGroups(expected))
		})
	})

	Context("in final documents", func() {

		It("blank line between 2 paragraphs", func() {
			source := `first paragraph
 
second paragraph`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{Content: "first paragraph"},
						},
					},
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{Content: "second paragraph"},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})
		It("blank line with spaces and tabs between 2 paragraphs and after second paragraph", func() {
			source := `first paragraph
		 

		
second paragraph
`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{Content: "first paragraph"},
						},
					},
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{Content: "second paragraph"},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})
	})
})
