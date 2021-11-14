package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("paragraphs", func() {

	Context("in final documents", func() {

		Context("thematic breaks", func() {

			It("form1 by itself", func() {
				source := "***"
				expected := &types.Document{
					Elements: []interface{}{
						&types.ThematicBreak{},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("form2 by itself", func() {
				source := "* * *"
				expected := &types.Document{
					Elements: []interface{}{
						&types.ThematicBreak{},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("form3 by itself", func() {
				source := "---"
				expected := &types.Document{
					Elements: []interface{}{
						&types.ThematicBreak{},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("form4 by itself", func() {
				source := "- - -"
				expected := &types.Document{
					Elements: []interface{}{
						&types.ThematicBreak{},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("form5 by itself", func() {
				source := "___"
				expected := &types.Document{
					Elements: []interface{}{
						&types.ThematicBreak{},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("form4 by itself", func() {
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
		})
	})
})
