package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("documents", func() {

	Context("draft document", func() {

		It("empty document", func() {
			source := ``
			expected := types.DraftDocument{}
			Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
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
