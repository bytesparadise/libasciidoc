package testsupport_test

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/bytesparadise/libasciidoc/testsupport"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("document fragments matcher", func() {

	// given
	expected := &types.TableOfContents{
		Sections: []*types.ToCSection{
			{
				Title: "root",
			},
		},
	}
	matcher := testsupport.MatchTableOfContents(expected)

	It("should match", func() {
		// given
		actual := &types.TableOfContents{
			Sections: []*types.ToCSection{
				{
					Title: "root",
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
		actual := &types.TableOfContents{
			Sections: []*types.ToCSection{
				{
					Title: "something else",
				},
			},
		}
		// when
		result, err := matcher.Match(actual)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(BeFalse())
		diffs := cmp.Diff(expected, actual, cmpopts.IgnoreUnexported(types.TableOfContents{}))
		Expect(matcher.FailureMessage(actual)).To(Equal(fmt.Sprintf("expected table of contents to match:\n%s", diffs)))
		Expect(matcher.NegatedFailureMessage(actual)).To(Equal(fmt.Sprintf("expected table of contents not to match:\n%s", diffs)))
	})

	It("should return error when invalid type is input", func() {
		// when
		result, err := matcher.Match(1)
		// then
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("MatchDocumentFragment matcher expects a *types.TableOfContents (actual: int)"))
		Expect(result).To(BeFalse())
	})

})
