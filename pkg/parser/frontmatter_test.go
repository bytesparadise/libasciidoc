package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("front-matter", func() {
	Context("yaml front-matter", func() {

		It("front-matter with simple attributes", func() {
			actualDocument := `---
title: a title
author: Xavier
---

first paragraph`
			expectedResult := types.Document{
				Attributes: map[string]interface{}{
					"title":  "a title", // TODO: convert `title` attribute from front-matter into `doctitle` here ?
					"author": "Xavier",
				},
				ElementReferences: map[string]interface{}{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: map[string]interface{}{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "first paragraph"},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualDocument)
		})

		It("empty front-matter", func() {
			actualDocument := `---
---

first paragraph`
			expectedResult := types.Document{
				Attributes:        map[string]interface{}{},
				ElementReferences: map[string]interface{}{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: map[string]interface{}{},
						Lines: []types.InlineElements{
							{
								types.StringElement{Content: "first paragraph"},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedResult, actualDocument)
		})
	})

})
