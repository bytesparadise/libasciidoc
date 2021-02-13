package parser

import (
	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("stack", func() {

	It("should push in empty stack", func() {
		// given
		s := newStack()
		// when
		s.push("cookie")
		// then
		Expect(s.size()).To(Equal(1))
	})

	It("should push in non-empty stack", func() {
		// given
		s := newStack()
		s.push("cookie")
		Expect(s.size()).To(Equal(1))
		// when
		s.push("chocolate")
		// then
		Expect(s.size()).To(Equal(2))
	})

	It("should pop from empty stack", func() {
		// given
		s := newStack()
		// when
		r := s.pop()
		// then
		Expect(r).To(BeNil())
	})

	It("should pop from non-empty stack", func() {
		// given
		s := newStack()
		s.push("cookie")
		s.push("chocolate")
		Expect(s.size()).To(Equal(2))
		// when
		r := s.pop()
		// then
		Expect(r).To(Equal("chocolate"))
		Expect(s.size()).To(Equal(1))
		// when
		r = s.pop()
		// then
		Expect(r).To(Equal("cookie"))
		Expect(s.size()).To(Equal(0))
		// when
		r = s.pop()
		// then
		Expect(r).To(BeNil())
		Expect(s.size()).To(Equal(0))
	})

	It("should get from empty stack", func() {
		// given
		s := newStack()
		// when
		r := s.get()
		// then
		Expect(r).To(BeNil())
	})

	It("should get from non-empty stack", func() {
		// given
		s := newStack()
		s.push("cookie")
		s.push("chocolate")
		Expect(s.size()).To(Equal(2))
		// when
		r := s.get()
		// then
		Expect(r).To(Equal("chocolate"))
		Expect(s.size()).To(Equal(2))
	})
})
