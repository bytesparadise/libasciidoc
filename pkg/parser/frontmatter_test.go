package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("front-matters - draft", func() {

	Context("yaml front-matter", func() {

		It("front-matter with simple attributes", func() {
			source := `---
title: a title
author: Xavier
---

first paragraph`
			expected := types.DraftDocument{
				FrontMatter: types.FrontMatter{
					Content: map[string]interface{}{
						"title":  "a title",
						"author": "Xavier",
					},
				},
				Blocks: []interface{}{
					types.BlankLine{},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "first paragraph"},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})

		It("empty front-matter", func() {
			source := `---
---

first paragraph`
			expected := types.DraftDocument{
				FrontMatter: types.FrontMatter{
					Content: map[string]interface{}{},
				},
				Blocks: []interface{}{
					types.BlankLine{},
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "first paragraph"},
							},
						},
					},
				},
			}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
		})
	})

})

var _ = Describe("front-matters", func() {

	Context("yaml front-matter", func() {

		It("front-matter with simple attributes", func() {
			source := `---
title: a title
author: Xavier
---

first paragraph`
			expected := types.Document{
				Attributes: types.DocumentAttributes{
					"title":  "a title", // TODO: convert `title` attribute from front-matter into `doctitle` here ?
					"author": "Xavier",
				},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "first paragraph"},
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("empty front-matter", func() {
			source := `---
---

first paragraph`
			expected := types.Document{
				Attributes:         types.DocumentAttributes{},
				ElementReferences:  types.ElementReferences{},
				Footnotes:          types.Footnotes{},
				FootnoteReferences: types.FootnoteReferences{},
				Elements: []interface{}{
					types.Paragraph{
						Attributes: types.ElementAttributes{},
						Lines: [][]interface{}{
							{
								types.StringElement{Content: "first paragraph"},
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})
	})

})
