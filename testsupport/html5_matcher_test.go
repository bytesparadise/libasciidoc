package testsupport_test

import (
	"fmt"
	"strings"
	"time"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/testsupport"
	"github.com/sergi/go-diff/diffmatchpatch"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("html5 template matcher", func() {

	// given
	now := time.Now()
	tmpl := `<p>{{.LastUpdated}}</p>`
	matcher := testsupport.MatchHTMLTemplate(tmpl, now)

	It("should match", func() {
		// given
		actual := `<p>` + now.Format("2006-01-02 15:04:05 -0700") + `</p>`
		// when
		result, err := matcher.Match(actual)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(BeTrue())
	})

	It("should not match", func() {
		// given
		actual := `<p>cheesecake</p>`
		// when
		result, err := matcher.Match(actual)
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(BeFalse())
		expected := strings.Replace(tmpl, "{{.LastUpdated}}", now.Format(configuration.LastUpdatedFormat), 1)
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(actual, expected, true)
		Expect(matcher.FailureMessage(actual)).To(Equal(fmt.Sprintf("expected HTML5 documents to match:\n%s", dmp.DiffPrettyText(diffs))))
		Expect(matcher.NegatedFailureMessage(actual)).To(Equal(fmt.Sprintf("expected HTML5 documents not to match:\n%s", dmp.DiffPrettyText(diffs))))
	})

	It("should return error when invalid type is input", func() {
		// when
		result, err := matcher.Match(1)
		// then
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("MatchHTMLTemplate matcher expects a string (actual: int)"))
		Expect(result).To(BeFalse())
	})

})
