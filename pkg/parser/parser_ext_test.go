package parser

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("block delimiter tracker", func() {

	It("should not be within delimited block", func() {
		// given
		t := newBlockDelimiterTracker()
		// then
		Expect(t.stack).To(BeEmpty())
	})

	It("should be within delimited block", func() {
		// given
		t := newBlockDelimiterTracker()
		// when
		t.track(types.Listing, 4) // entered block
		// then
		Expect(t.stack).NotTo(BeEmpty())
	})

	It("should still be within delimited block - case 1", func() {
		// given
		t := newBlockDelimiterTracker()
		// when
		t.track(types.Listing, 4) // entered block
		t.track(types.Comment, 4) // entered another block
		// then
		Expect(t.stack).NotTo(BeEmpty())
	})

	It("should still be within delimited block - case 2", func() {
		// given
		t := newBlockDelimiterTracker()
		// when
		t.track(types.Listing, 5) // entered first block
		t.track(types.Listing, 4) // entered second block
		// then
		Expect(t.stack).NotTo(BeEmpty())
	})

	It("should not be within delimited block anymore - case 1", func() {
		// given
		t := newBlockDelimiterTracker()
		// when
		t.track(types.Listing, 4) // entered block
		t.track(types.Listing, 4) // exited block
		// then
		Expect(t.stack).To(BeEmpty())
	})

	It("should not be within delimited block anymore - case 2", func() {
		// given
		t := newBlockDelimiterTracker()
		// when
		t.track(types.Listing, 4) // entered first block
		t.track(types.Comment, 4) // entered second block
		t.track(types.Comment, 4) // existed second block
		t.track(types.Listing, 4) // exited first block
		// then
		Expect(t.stack).To(BeEmpty())
	})

	It("should not be within delimited block anymore - case 3", func() {
		// given
		t := newBlockDelimiterTracker()
		// when
		t.track(types.Listing, 5) // entered first block
		t.track(types.Listing, 4) // entered second block
		t.track(types.Listing, 4) // exited second block
		t.track(types.Listing, 5) // exited first block
		// then
		Expect(t.stack).To(BeEmpty())
	})

	It("should not be within delimited block anymore - case 4", func() {
		// given
		t := newBlockDelimiterTracker()
		// when
		t.track(types.Listing, 4) // entered first block
		t.track(types.Listing, 5) // entered second block
		t.track(types.Listing, 5) // exited second block
		t.track(types.Listing, 4) // exited first block
		// then
		Expect(t.stack).To(BeEmpty())
	})

})
