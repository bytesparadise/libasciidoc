package testsupport_test

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("html5 rendering assertions", func() {

	Context("HTML5 element matcher", func() {

		expected := `<div class="paragraph">
<p>hello, world!</p>
</div>`

		It("should match", func() {
			// given
			matcher := testsupport.RenderHTML5Element(expected)
			actual := "hello, world!"
			// when
			result, err := matcher.Match(actual)
			// then
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeTrue())
		})

		It("should not match", func() {
			// given
			matcher := testsupport.RenderHTML5Element(expected)
			actual := "foo"
			// when
			result, err := matcher.Match(actual)
			// then
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeFalse())
			// also verify the messages
			obtained := `<div class="paragraph">
<p>foo</p>
</div>`
			Expect(matcher.FailureMessage(actual)).To(Equal(fmt.Sprintf("expected HTML5 elements to match:\n\texpected: '%v'\n\tactual:   '%v'", expected, obtained)))
			Expect(matcher.NegatedFailureMessage(actual)).To(Equal(fmt.Sprintf("expected HTML5 elements not to match:\n\texpected: '%v'\n\tactual:   '%v'", expected, obtained)))
		})

		It("should return error when invalid type is input", func() {
			// given
			matcher := testsupport.RenderHTML5Element("")
			// when
			result, err := matcher.Match(1) // not a string
			// then
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("RenderHTML5Element matcher expects a string (actual: int)"))
			Expect(result).To(BeFalse())
		})
	})

	Context("HTML5 body matcher", func() {

		expected := `<div class="paragraph">
<p>hello, world!</p>
</div>`

		It("should match", func() {
			// given
			matcher := testsupport.RenderHTML5Body(expected)
			actual := "hello, world!"
			// when
			result, err := matcher.Match(actual)
			// then
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeTrue())
		})

		It("should not match", func() {
			// given
			matcher := testsupport.RenderHTML5Body(expected)
			actual := "foo"
			// when
			result, err := matcher.Match(actual)
			// then
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeFalse())
			// also verify messages
			obtained := `<div class="paragraph">
<p>foo</p>
</div>`
			Expect(matcher.FailureMessage(actual)).To(Equal(fmt.Sprintf("expected HTML5 bodies to match:\n\texpected: '%v'\n\tactual:   '%v'", expected, obtained)))
			Expect(matcher.NegatedFailureMessage(actual)).To(Equal(fmt.Sprintf("expected HTML5 bodies not to match:\n\texpected: '%v'\n\tactual:   '%v'", expected, obtained)))
		})

		It("should return error when invalid type is input", func() {
			// given
			matcher := testsupport.RenderHTML5Body("")
			// when
			result, err := matcher.Match(1) // not a string
			// then
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("RenderHTML5Body matcher expects a string (actual: int)"))
			Expect(result).To(BeFalse())
		})
	})

	Context("HTML5 title matcher", func() {

		expected := `hello, world!`

		It("should match strings", func() {
			// given
			matcher := testsupport.RenderHTML5Title(expected)
			actual := `= hello, world!`
			// when
			result, err := matcher.Match(actual)
			// then
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeTrue())
		})

		It("should match nils", func() {
			// given
			matcher := testsupport.RenderHTML5Title(nil)
			actual := `foo` // no title in this doc
			// when
			result, err := matcher.Match(actual)
			// then
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeTrue())
		})

		It("should not match strings", func() {
			// given
			matcher := testsupport.RenderHTML5Title(expected)
			actual := `foo` // no title in this doc
			// when
			result, err := matcher.Match(actual)
			// then
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("invalid type of title (<nil>)"))
			Expect(result).To(BeFalse())
			// also verify error message
			var obtained *string
			Expect(matcher.FailureMessage(actual)).To(Equal(fmt.Sprintf("expected HTML5 titles to match:\n\texpected: '%v'\n\tactual:   '%v'", expected, obtained)))
			Expect(matcher.NegatedFailureMessage(actual)).To(Equal(fmt.Sprintf("expected HTML5 titles not to match:\n\texpected: '%v'\n\tactual:   '%v'", expected, obtained)))
		})

		It("should return error when invalid type is input", func() {
			// given
			matcher := testsupport.RenderHTML5Title("")
			// when
			result, err := matcher.Match(1) // not a string
			// then
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("RenderHTML5Title matcher expects a string (actual: int)"))
			Expect(result).To(BeFalse())
		})
	})
})
