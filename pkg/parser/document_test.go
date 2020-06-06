package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("documents", func() {

	Context("draft document", func() {

		It("empty docunment", func() {
			source := ``
			expected := types.DraftDocument{
				Blocks: []interface{}{},
			}
			Expect(ParseDraftDocument(source)).To(Equal(expected))
		})
	})

	Context("final document", func() {

		It("empty document", func() {
			source := ``
			expected := types.Document{
				Elements: []interface{}{},
			}
			Expect(ParseDocument(source)).To(Equal(expected))
		})
	})
})
