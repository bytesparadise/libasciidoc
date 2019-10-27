package testsupport

import (
	"reflect"

	"github.com/davecgh/go-spew/spew"
	. "github.com/onsi/ginkgo" // nolint: golint
	"github.com/sergi/go-diff/diffmatchpatch"
)

type comparison struct {
	actual   string
	expected string
	diffs    string
}

// compare compares the 'actual' vs 'expected' values and produces a 'comparison' report
func compare(actual interface{}, expected interface{}) comparison {
	c := comparison{
		actual:   spew.Sdump(actual),
		expected: spew.Sdump(expected),
	}
	if !reflect.DeepEqual(actual, expected) {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(c.actual, c.expected, true)
		c.diffs = dmp.DiffPrettyText(diffs)
	}
	GinkgoT().Logf("actual:\n%s", c.actual)
	GinkgoT().Logf("expected:\n%s", c.expected)
	GinkgoT().Logf("diff:\n%s", c.diffs)
	return c
}
