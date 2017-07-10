package parser_test

import (
	"github.com/bytesparadise/libasciidoc/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Testing with Ginkgo", func() {
	It("delimited source block with single line", func() {

		content := "some source code"
		actualContent := "```\n" + content + "\n```"
		expectedDocument := &types.Document{
			Elements: []types.DocElement{
				&types.DelimitedBlock{
					Kind:    types.SourceBlock,
					Content: content,
				},
			},
		}
		compare(GinkgoT(), expectedDocument, actualContent)
	})
	It("delimited source block with multiple lines", func() {

		content := "some source code\nwith an empty line\n\nin the middle"
		actualContent := "```\n" + content + "\n```"
		expectedDocument := &types.Document{
			Elements: []types.DocElement{
				&types.DelimitedBlock{
					Kind:    types.SourceBlock,
					Content: content,
				},
			},
		}
		compare(GinkgoT(), expectedDocument, actualContent)
	})
	It("delimited source block with no line", func() {

		content := ""
		actualContent := "```\n" + content + "```"
		expectedDocument := &types.Document{
			Elements: []types.DocElement{
				&types.DelimitedBlock{
					Kind:    types.SourceBlock,
					Content: content,
				},
			},
		}
		compare(GinkgoT(), expectedDocument, actualContent)
	})
})
