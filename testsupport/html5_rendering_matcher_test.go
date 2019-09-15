package testsupport_test

import (
	"github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("html5 rendering assertions", func() {

	expected := `<div class="paragraph">
<p>hello, world!</p>
</div>`

	It("should match", func() {
		// given
		matcher := testsupport.RenderHTML5(expected)
		// when
		result, err := matcher.Match("hello, world!")
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(BeTrue())
	})

	It("should not match", func() {
		// given
		matcher := testsupport.EqualDocument(expected)
		// when
		result, err := matcher.Match("meh")
		// then
		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(BeFalse())
	})

})
