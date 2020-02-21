package testsupport_test

import (
	"fmt"
	"time"

	"github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("html5 template matcher", func() {

	// given
	now := time.Now()
	tmpl := `<p>{{.LastUpdated}}</p>`
	matcher := testsupport.MatchHTML5Template(tmpl, now)

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
		Expect(matcher.FailureMessage(actual)).To(Equal(fmt.Sprintf("expected HTML5 documents to match:\n\texpected: '%v'\n\tactual:   '%v'", `<p>`+now.Format("2006-01-02 15:04:05 -0700")+`</p>`, actual)))
		Expect(matcher.NegatedFailureMessage(actual)).To(Equal(fmt.Sprintf("expected HTML5 documents not to match:\n\texpected: '%v'\n\tactual:   '%v'", `<p>`+now.Format("2006-01-02 15:04:05 -0700")+`</p>`, actual)))
	})

	It("should return error when invalid type is input", func() {
		// when
		result, err := matcher.Match(1)
		// then
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("MatchHTML5Template matcher expects a string (actual: int)"))
		Expect(result).To(BeFalse())
	})

})
