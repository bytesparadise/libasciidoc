package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("inline menus", func() {

	Context("in final documents", func() {

		It("with main path", func() {
			source := `:experimental:
 
Select menu:File[].`
			expected := &types.Document{
				Elements: []interface{}{
					&types.AttributeDeclaration{
						Name: "experimental",
					},
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "Select ",
							},
							&types.InlineMenu{
								Path: []string{
									"File",
								},
							},
							&types.StringElement{
								Content: ".",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("with single sub path", func() {
			source := `:experimental:
 
Select menu:File[Save].`
			expected := &types.Document{
				Elements: []interface{}{
					&types.AttributeDeclaration{
						Name: "experimental",
					},
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "Select ",
							},
							&types.InlineMenu{
								Path: []string{
									"File",
									"Save",
								},
							},
							&types.StringElement{
								Content: ".",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("with multiple sub paths", func() {
			source := `:experimental:
 
Select menu:File[Zoom > Reset].`
			expected := &types.Document{
				Elements: []interface{}{
					&types.AttributeDeclaration{
						Name: "experimental",
					},
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "Select ",
							},
							&types.InlineMenu{
								Path: []string{
									"File",
									"Zoom",
									"Reset",
								},
							},
							&types.StringElement{
								Content: ".",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("when experimental is not enabled", func() {
			source := `Select menu:File[Zoom > Reset].`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "Select menu:File[Zoom ",
							},
							&types.SpecialCharacter{
								Name: ">",
							},
							&types.StringElement{
								Content: " Reset].",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})
	})
})
