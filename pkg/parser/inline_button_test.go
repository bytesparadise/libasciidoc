package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("inline buttons", func() {

	Context("in final documents", func() {

		It("when experimental is enabled", func() {
			source := `:experimental:
 
Click on btn:[OK].`
			expected := &types.Document{
				Elements: []interface{}{
					&types.DocumentHeader{
						Elements: []interface{}{
							&types.AttributeDeclaration{
								Name: "experimental",
							},
						},
					},
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "Click on ",
							},
							&types.InlineButton{
								Attributes: types.Attributes{
									types.AttrButtonLabel: "OK",
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
			source := `Click on btn:[OK].`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "Click on btn:[OK].",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})
	})
})
