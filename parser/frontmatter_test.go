package parser_test

import (
	"github.com/bytesparadise/libasciidoc/types"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Front-Matter", func() {
	Context("YAML Front-matter", func() {

		It("front-matter with simple attributes", func() {
			actualDocument := `---
title: a title
author: Xavier
---

first paragraph`
			expectedDocument := &types.Document{
				Attributes: &types.DocumentAttributes{
					"title":  "a title",
					"author": "Xavier",
				},
				Elements: []types.DocElement{
					&types.Paragraph{
						Lines: []*types.InlineContent{
							&types.InlineContent{
								Elements: []types.InlineElement{
									&types.StringElement{Content: "first paragraph"},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualDocument)
		})

		It("empty front-matter", func() {
			actualDocument := `---
---

first paragraph`
			expectedDocument := &types.Document{
				Attributes: &types.DocumentAttributes{},
				Elements: []types.DocElement{
					&types.Paragraph{
						Lines: []*types.InlineContent{
							&types.InlineContent{
								Elements: []types.InlineElement{
									&types.StringElement{Content: "first paragraph"},
								},
							},
						},
					},
				},
			}
			verify(GinkgoT(), expectedDocument, actualDocument)
		})
	})

})
