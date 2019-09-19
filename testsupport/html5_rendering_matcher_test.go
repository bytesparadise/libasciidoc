package testsupport_test

import (
	"fmt"

	"github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("html5 rendering assertions", func() {

	Context("HTML5 element matcher", func() {

		actual := "hello, world!"
		expected := `<div class="paragraph">
<p>hello, world!</p>
</div>`

		It("should match", func() {
			// given
			matcher := testsupport.RenderHTML5Element(expected)
			// when
			result, err := matcher.Match(actual)
			// then
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeTrue())
		})

		It("should not match", func() {
			// given
			matcher := testsupport.RenderHTML5Element(expected)
			// when
			result, err := matcher.Match("foo")
			// then
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeFalse())
		})

		Context("messages", func() {

			It("failure message", func() {
				// given
				matcher := testsupport.RenderHTML5Element(expected)
				_, err := matcher.Match(actual)
				Expect(err).ToNot(HaveOccurred())
				// when
				msg := matcher.FailureMessage(actual)
				// then
				Expect(msg).To(Equal(fmt.Sprintf("expected HTML5 elements to match:\n\texpected: '%v'\n\tactual:   '%v'", expected, expected)))
			})

			It("negated failure message", func() {
				// given
				matcher := testsupport.RenderHTML5Element(expected)
				_, err := matcher.Match(actual)
				Expect(err).ToNot(HaveOccurred())
				// when
				msg := matcher.NegatedFailureMessage(actual)
				// then
				Expect(msg).To(Equal(fmt.Sprintf("expected HTML5 elements not to match:\n\texpected: '%v'\n\tactual:   '%v'", expected, expected)))

			})
		})

		Context("failures", func() {

			It("should return error when invalid type is input", func() {
				// given
				matcher := testsupport.RenderHTML5Element("")
				_, err := matcher.Match(actual)
				Expect(err).ToNot(HaveOccurred())
				// when
				result, err := matcher.Match(1) // not a string
				// then
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("RenderHTML5Element matcher expects a string (actual: int)"))
				Expect(result).To(BeFalse())
			})
		})
	})

	Context("HTML5 body matcher", func() {

		actual := "hello, world!"
		expected := `<div class="paragraph">
<p>hello, world!</p>
</div>`

		It("should match", func() {
			// given
			matcher := testsupport.RenderHTML5Body(expected)
			// when
			result, err := matcher.Match(actual)
			// then
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeTrue())
		})

		It("should not match", func() {
			// given
			matcher := testsupport.RenderHTML5Body(expected)
			// when
			result, err := matcher.Match("foo")
			// then
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeFalse())
		})

		Context("messages", func() {

			It("failure message", func() {
				// given
				matcher := testsupport.RenderHTML5Body(expected)
				_, err := matcher.Match(actual)
				Expect(err).ToNot(HaveOccurred())
				// when
				msg := matcher.FailureMessage(actual)
				// then
				Expect(msg).To(Equal(fmt.Sprintf("expected HTML5 bodies to match:\n\texpected: '%v'\n\tactual:   '%v'", expected, expected)))
			})

			It("negated failure message", func() {
				// given
				matcher := testsupport.RenderHTML5Body(expected)
				_, err := matcher.Match(actual)
				Expect(err).ToNot(HaveOccurred())
				// when
				msg := matcher.NegatedFailureMessage(actual)
				// then
				Expect(msg).To(Equal(fmt.Sprintf("expected HTML5 bodies not to match:\n\texpected: '%v'\n\tactual:   '%v'", expected, expected)))

			})
		})

		Context("failures", func() {

			It("should return error when invalid type is input", func() {
				// given
				matcher := testsupport.RenderHTML5Body("")
				_, err := matcher.Match(actual)
				Expect(err).ToNot(HaveOccurred())
				// when
				result, err := matcher.Match(1) // not a string
				// then
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("RenderHTML5Body matcher expects a string (actual: int)"))
				Expect(result).To(BeFalse())
			})
		})
	})

	Context("HTML5 title matcher", func() {

		actual := `= hello, world!`
		expected := `hello, world!`

		It("should match strings", func() {
			// given
			matcher := testsupport.RenderHTML5Title(expected)
			// when
			result, err := matcher.Match(actual)
			// then
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeTrue())
		})

		It("should not match strings", func() {
			// given
			matcher := testsupport.RenderHTML5Title(expected)
			// when
			result, err := matcher.Match("foo")
			// then
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("invalid type of title (<nil>)"))
			Expect(result).To(BeFalse())
		})

		It("should not match nils", func() {
			// given
			matcher := testsupport.RenderHTML5Title(nil)
			// when
			result, err := matcher.Match("foo") // no title in this doc
			// then
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(BeTrue())
		})

		Context("messages", func() {

			It("failure message", func() {
				// given
				matcher := testsupport.RenderHTML5Title(expected)
				_, err := matcher.Match(actual)
				Expect(err).ToNot(HaveOccurred())
				// when
				msg := matcher.FailureMessage(actual)
				// then
				Expect(msg).To(Equal(fmt.Sprintf("expected HTML5 titles to match:\n\texpected: '%v'\n\tactual:   '%v'", expected, expected)))
			})

			It("negated failure message", func() {
				// given
				matcher := testsupport.RenderHTML5Title(expected)
				_, err := matcher.Match(actual)
				Expect(err).ToNot(HaveOccurred())
				// when
				msg := matcher.NegatedFailureMessage(actual)
				// then
				Expect(msg).To(Equal(fmt.Sprintf("expected HTML5 titles not to match:\n\texpected: '%v'\n\tactual:   '%v'", expected, expected)))

			})
		})

		Context("failures", func() {

			It("should return error when invalid type is input", func() {
				// given
				matcher := testsupport.RenderHTML5Title("")
				_, err := matcher.Match(actual)
				Expect(err).ToNot(HaveOccurred())
				// when
				result, err := matcher.Match(1) // not a string
				// then
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("RenderHTML5Title matcher expects a string (actual: int)"))
				Expect(result).To(BeFalse())
			})
		})
	})
})
