package testsupport_test

import (
	"fmt"
	"strings"
	"time"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/testsupport"
	"github.com/google/go-cmp/cmp"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("html5 matcher", func() {

	It("should match", func() {
		// given
		actual := `<p>cheesecake</p>`
		matcher := testsupport.MatchHTML(`<p>cheesecake</p>`) // same content
		// when
		result, err := matcher.Match(actual)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(BeTrue())
	})

	It("should not match", func() {
		// given
		actual := `<p>cheesecake</p>`
		expected := `<p>chocolate</p>`
		matcher := testsupport.MatchHTML(expected) // not the same content
		// when
		result, err := matcher.Match(actual)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(BeFalse())
		diffs := cmp.Diff(expected, actual)
		Expect(matcher.FailureMessage(actual)).To(Equal(fmt.Sprintf("expected HTML5 documents to match:\n%s", diffs)))
		Expect(matcher.NegatedFailureMessage(actual)).To(Equal(fmt.Sprintf("expected HTML5 documents not to match:\n%s", diffs)))
	})

	It("should return error when invalid type is input", func() {
		// given
		matcher := testsupport.MatchHTML(`<p>cheesecake</p>`)
		// when
		result, err := matcher.Match(1)
		// then
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("MatchHTML matcher expects a string (actual: int)"))
		Expect(result).To(BeFalse())
	})

})

var _ = Describe("html5 template matcher", func() {

	// given
	now := time.Now()
	tmpl := `<p>{{ .LastUpdated }}</p>`

	It("should match", func() {
		// given
		matcher := testsupport.MatchHTMLTemplate(tmpl, struct {
			LastUpdated string
		}{
			LastUpdated: now.Format(configuration.LastUpdatedFormat),
		})
		actual := `<p>` + now.Format(configuration.LastUpdatedFormat) + `</p>`
		// when
		result, err := matcher.Match(actual)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(BeTrue())
	})

	It("should not match", func() {
		// given
		actual := `<p>cheesecake</p>`
		matcher := testsupport.MatchHTMLTemplate(tmpl, struct {
			LastUpdated string
		}{
			LastUpdated: now.Format(configuration.LastUpdatedFormat),
		})
		// when
		result, err := matcher.Match(actual)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(BeFalse())
		expected := strings.Replace(tmpl, "{{ .LastUpdated }}", now.Format(configuration.LastUpdatedFormat), 1)
		diffs := cmp.Diff(expected, actual)
		Expect(matcher.FailureMessage(actual)).To(Equal(fmt.Sprintf("expected HTML5 documents to match:\n%s", diffs)))
		Expect(matcher.NegatedFailureMessage(actual)).To(Equal(fmt.Sprintf("expected HTML5 documents not to match:\n%s", diffs)))
	})

	It("should return error when invalid type is input", func() {
		// given
		matcher := testsupport.MatchHTMLTemplate(tmpl, struct {
			LastUpdated string
		}{
			LastUpdated: now.Format(configuration.LastUpdatedFormat),
		})
		// when
		result, err := matcher.Match(1)
		// then
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("MatchHTMLTemplate matcher expects a string (actual: int)"))
		Expect(result).To(BeFalse())
	})

})
