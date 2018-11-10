// +build pending

package test_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
)

var _ = Describe("pending", func() {

	// verifies that all files in the `supported` subfolder match their sibling golden file
	DescribeTable("supported", compare, entries("fixtures/pending/*.adoc")...)

})
