package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("thematic breaks", func() {

	Context("in final documents", func() {

		It("by itself", func() {
			source := "'''"
			expected := &types.Document{
				Elements: []interface{}{
					&types.ThematicBreak{},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("mk form1 by itself", func() {
			source := "***"
			expected := &types.Document{
				Elements: []interface{}{
					&types.ThematicBreak{},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("mk form2 by itself", func() {
			source := "* * *"
			expected := &types.Document{
				Elements: []interface{}{
					&types.ThematicBreak{},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("mk form3 by itself", func() {
			source := "---"
			expected := &types.Document{
				Elements: []interface{}{
					&types.ThematicBreak{},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("mk form4 by itself", func() {
			source := "- - -"
			expected := &types.Document{
				Elements: []interface{}{
					&types.ThematicBreak{},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("mk form5 by itself", func() {
			source := "___"
			expected := &types.Document{
				Elements: []interface{}{
					&types.ThematicBreak{},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("mk form4 by itself", func() {
			source := "_ _ _"
			expected := &types.Document{
				Elements: []interface{}{
					&types.ThematicBreak{},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("with leading text", func() {
			source := "text ***"
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{Content: "text ***"},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("with trailing spaces", func() {
			source := "'''   "
			expected := &types.Document{
				Elements: []interface{}{
					&types.ThematicBreak{},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		// NB: three asterisks gets confused with bullets if with trailing text
		It("with trailing text", func() {
			source := "* * * text"
			expected := &types.Document{
				Elements: []interface{}{
					&types.List{
						Kind: types.UnorderedListKind,
						Elements: []types.ListElement{
							&types.UnorderedListElement{
								BulletStyle: types.OneAsterisk,
								CheckStyle:  types.NoCheck,
								Elements: []interface{}{
									&types.Paragraph{
										Elements: []interface{}{
											&types.StringElement{Content: "* * text"},
										},
									},
								},
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		// does not break when embedded within a paragrap
		It("between 2 paragraphs", func() {
			source := `
some

'''

content`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{Content: "some"},
						},
					},
					&types.ThematicBreak{},
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{Content: "content"},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		// does not break when embedded within a paragrap
		It("within paragraph", func() {
			source := `
some
'''
content`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{Content: "some\n'''\ncontent"},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

	})
})
