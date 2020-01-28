package testsupport_test

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/bytesparadise/libasciidoc"
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("document metadata assertions", func() {

	lastUpdated := time.Now()
	expected := types.Metadata{
		LastUpdated: lastUpdated.Format(renderer.LastUpdatedFormat),
		TableOfContents: types.TableOfContents{
			Sections: []types.ToCSection{
				{
					ID:       "_section_1",
					Level:    1,
					Title:    "Section 1",
					Children: []types.ToCSection{},
				},
			},
		},
	}

	It("should match", func() {
		// given
		matcher := testsupport.HaveMetadata(expected, lastUpdated)
		actual := `== Section 1`
		// when
		result, err := matcher.Match(actual)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(BeTrue())
	})

	It("should not match", func() {
		// given
		matcher := testsupport.HaveMetadata(expected, lastUpdated)
		actual := `foo`
		// when
		result, err := matcher.Match(actual)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(BeFalse())
		// also verify messages
		obtained, err := libasciidoc.ConvertToHTML("", strings.NewReader(actual), bytes.NewBuffer(nil), renderer.IncludeHeaderFooter(false), renderer.LastUpdated(lastUpdated))
		Expect(err).ToNot(HaveOccurred())
		GinkgoT().Logf(matcher.FailureMessage(result))
		GinkgoT().Logf(fmt.Sprintf("expected metadata to match:\n%s", compare(obtained, expected)))
		Expect(matcher.FailureMessage(actual)).To(Equal(fmt.Sprintf("expected metadata to match:\n%s", compare(obtained, expected))))
		Expect(matcher.NegatedFailureMessage(actual)).To(Equal(fmt.Sprintf("expected metadata not to match:\n%s", compare(obtained, expected))))
	})

	It("should return error when invalid type is input", func() {
		// given
		matcher := testsupport.HaveMetadata(types.Metadata{}, lastUpdated)
		// when
		result, err := matcher.Match(1) // not a doc
		// then
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("HaveMetadata matcher expects a string (actual: int)"))
		Expect(result).To(BeFalse())
	})
})
