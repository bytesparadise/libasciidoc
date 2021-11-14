package testsupport_test

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/bytesparadise/libasciidoc/testsupport"

	"github.com/davecgh/go-spew/spew"
	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
	"github.com/sergi/go-diff/diffmatchpatch"
)

var _ = Describe("document fragments matcher", func() {

	// given
	expected := []types.DocumentFragment{
		{
			Elements: []interface{}{
				types.RawLine("a paragraph."),
			},
		},
	}
	matcher := testsupport.MatchDocumentFragments(expected)

	It("should match", func() {
		// given
		actual := []types.DocumentFragment{
			{
				Elements: []interface{}{
					types.RawLine("a paragraph."),
				},
			},
		}
		// when
		result, err := matcher.Match(actual)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(BeTrue())
	})

	It("should not match", func() {
		// given
		actual := []types.DocumentFragment{
			{
				Elements: []interface{}{
					types.RawLine("something else"),
				},
			},
		}
		// when
		result, err := matcher.Match(actual)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(BeFalse())
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(spew.Sdump(actual), spew.Sdump(expected), true)
		Expect(matcher.FailureMessage(actual)).To(Equal(fmt.Sprintf("expected document fragments to match:\n%s", dmp.DiffPrettyText(diffs))))
		Expect(matcher.NegatedFailureMessage(actual)).To(Equal(fmt.Sprintf("expected document fragments not to match:\n%s", dmp.DiffPrettyText(diffs))))
	})

	It("should return error when invalid type is input", func() {
		// when
		result, err := matcher.Match(1)
		// then
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("MatchDocumentFragments matcher expects an array of types.DocumentFragment (actual: int)"))
		Expect(result).To(BeFalse())
	})

})
