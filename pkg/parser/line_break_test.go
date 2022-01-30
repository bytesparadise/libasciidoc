package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" // nolint:golint
	. "github.com/onsi/gomega" // nolint:golint
)

var _ = Describe("line breaks", func() {

	Context("in final documents", func() {

		It("simple case", func() {
			source := `since 2021 +`
			expected := &types.Document{
				Elements: []interface{}{
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "since 2021",
							},
							&types.LineBreak{},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("after punctuation", func() {
			source := `:author: Xavier
Copyright (C) 2021 {author}. +`
			expected := &types.Document{
				Elements: []interface{}{
					&types.AttributeDeclaration{
						Name:  "author",
						Value: "Xavier",
					},
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "Copyright © 2021 Xavier.",
							},
							&types.LineBreak{},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})
	})
})
