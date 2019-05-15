package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("user macros", func() {

	Context("user macros", func() {

		It("user block macro", func() {
			actualContent := "git::some/url.git[key1=value1,key2=value2]"
			expectedResult := types.UserMacro{
				Kind:  types.BlockMacro,
				Name:  "git",
				Value: "some/url.git",
				Attributes: types.ElementAttributes{
					"key1": "value1",
					"key2": "value2",
				},
				RawText: "git::some/url.git[key1=value1,key2=value2]",
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})

		It("inline user macro", func() {
			actualContent := "repository: git:some/url.git[key1=value1,key2=value2]"
			expectedResult := types.Paragraph{
				Attributes: types.ElementAttributes{},
				Lines: []types.InlineElements{
					{
						types.StringElement{
							Content: "repository: ",
						},
						types.UserMacro{
							Kind:  types.InlineMacro,
							Name:  "git",
							Value: "some/url.git",
							Attributes: types.ElementAttributes{
								"key1": "value1",
								"key2": "value2",
							},
							RawText: "git:some/url.git[key1=value1,key2=value2]",
						},
					},
				},
			}
			verifyWithPreprocessing(GinkgoT(), expectedResult, actualContent, parser.Entrypoint("DocumentBlock"))
		})
	})
})
