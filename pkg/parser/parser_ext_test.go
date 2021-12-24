package parser

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("block delimiter tracker", func() {

	It("should not be within delimited block", func() {
		// given
		t := newBlockDelimiterTracker()
		// then
		Expect(t.withinDelimitedBlock()).To(BeFalse())
	})

	It("should be within delimited block", func() {
		// given
		t := newBlockDelimiterTracker()
		// when
		t.push(types.BlockDelimiterKind(types.Listing)) // entered block
		// then
		Expect(t.withinDelimitedBlock()).To(BeTrue())
	})

	It("should still be within delimited block - case 1", func() {
		// given
		t := newBlockDelimiterTracker()
		// when
		t.push(types.BlockDelimiterKind(types.Listing)) // entered block
		t.push(types.BlockDelimiterKind(types.Comment)) // entered another block
		// then
		Expect(t.withinDelimitedBlock()).To(BeTrue())
	})

	It("should still be within delimited block - case 1", func() {
		// given
		t := newBlockDelimiterTracker()
		// when
		t.push(types.BlockDelimiterKind(types.Listing)) // entered first block
		t.push(types.BlockDelimiterKind(types.Comment)) // entered second block
		// then
		Expect(t.withinDelimitedBlock()).To(BeTrue())
	})

	It("should not be within delimited block anymore - case 1", func() {
		// given
		t := newBlockDelimiterTracker()
		// when
		t.push(types.BlockDelimiterKind(types.Listing)) // entered block
		t.push(types.BlockDelimiterKind(types.Listing)) // exited block
		// then
		Expect(t.withinDelimitedBlock()).To(BeFalse())
	})

	It("should not be within delimited block anymore - case 2", func() {
		// given
		t := newBlockDelimiterTracker()
		// when
		t.push(types.BlockDelimiterKind(types.Listing)) // entered first block
		t.push(types.BlockDelimiterKind(types.Comment)) // entered second block
		t.push(types.BlockDelimiterKind(types.Comment)) // existed second block
		t.push(types.BlockDelimiterKind(types.Listing)) // exited first block
		// then
		Expect(t.withinDelimitedBlock()).To(BeFalse())
	})

})
