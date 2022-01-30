package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" // nolint:golint
	. "github.com/onsi/gomega" // nolint:golint
)

var _ = Describe("front-matters", func() {

	Context("in final documents", func() {

		Context("yaml front-matter", func() {

			It("with simple attributes", func() {
				source := `---
title: a title
author: Xavier
---

first paragraph`
				expected := &types.Document{
					Elements: []interface{}{
						&types.FrontMatter{
							Attributes: types.Attributes{
								"title":  "a title", // TODO: convert `title` attribute from front-matter into `doctitle` here ?
								"author": "Xavier",
							},
						},
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "first paragraph"},
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
				expected := &types.Document{
					Elements: []interface{}{
						&types.FrontMatter{},
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{Content: "first paragraph"},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})
	})
})
