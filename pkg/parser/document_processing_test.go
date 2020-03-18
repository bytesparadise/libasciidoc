package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("document processing", func() {

	It("should retain attributes passed in configuration", func() {
		source := `[source]
----
foo
----`
		expected := types.Document{
			Attributes: types.DocumentAttributes{
				types.AttrSyntaxHighlighter: "pygments",
			},
			ElementReferences:  types.ElementReferences{},
			Footnotes:          types.Footnotes{},
			FootnoteReferences: types.FootnoteReferences{},
			Elements: []interface{}{
				types.DelimitedBlock{
					Attributes: types.ElementAttributes{
						types.AttrKind: types.Source,
					},
					Kind: types.Source,
					Elements: []interface{}{
						types.Paragraph{
							Attributes: types.ElementAttributes{},
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "foo",
									},
								},
							},
						},
					},
				},
			},
		}
		Expect(ParseDocument(source, configuration.WithAttributes(map[string]string{
			types.AttrSyntaxHighlighter: "pygments",
		}))).To(Equal(expected))
	})
})
