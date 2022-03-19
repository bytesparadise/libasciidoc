package testsupport_test

import (
	"fmt"
	"time"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/bytesparadise/libasciidoc/testsupport"
	"github.com/davecgh/go-spew/spew"
	"github.com/google/go-cmp/cmp"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("metadata matcher", func() {

	now := time.Now()

	It("should match", func() {
		// given
		actual := types.Metadata{
			Title:       "cheesecake",
			LastUpdated: now.Format(configuration.LastUpdatedFormat),
		}
		matcher := testsupport.MatchMetadata(types.Metadata{
			Title:       "cheesecake",
			LastUpdated: now.Format(configuration.LastUpdatedFormat),
		}) // same content
		// when
		result, err := matcher.Match(actual)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(BeTrue())
	})

	It("should not match", func() {
		// given
		actual := types.Metadata{
			Title:       "cheesecake",
			LastUpdated: now.Format(configuration.LastUpdatedFormat),
		}
		expected := types.Metadata{
			Title:       "chocolate",
			LastUpdated: now.Format(configuration.LastUpdatedFormat),
		} // not the same content
		matcher := testsupport.MatchMetadata(expected)
		// when
		result, err := matcher.Match(actual)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(BeFalse())
		diffs := cmp.Diff(spew.Sdump(expected), spew.Sdump(actual))
		Expect(matcher.FailureMessage(actual)).To(Equal(fmt.Sprintf("expected document metadata to match:\n%s", diffs)))
		Expect(matcher.NegatedFailureMessage(actual)).To(Equal(fmt.Sprintf("expected document metadata not to match:\n%s", diffs)))
	})

	It("should return error when invalid type is input", func() {
		// given
		matcher := testsupport.MatchMetadata(types.Metadata{
			Title:       "cheesecake",
			LastUpdated: now.Format(configuration.LastUpdatedFormat),
		})
		// when
		result, err := matcher.Match(1)
		// then
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("MatchMetadata matcher expects a 'types.Metadata' (actual: int)"))
		Expect(result).To(BeFalse())
	})
})
