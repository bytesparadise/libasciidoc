package testsupport_test

import (
	"reflect"

	"github.com/davecgh/go-spew/spew"
	"github.com/sergi/go-diff/diffmatchpatch"
)

// compare compares the 'actual' vs 'expected' values and produces a 'comparison' report
func compare(actual interface{}, expected interface{}) string {
	a := spew.Sdump(actual)
	e := spew.Sdump(expected)
	if reflect.DeepEqual(actual, expected) {
		return ""
	}
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(a, e, true)
	return dmp.DiffPrettyText(diffs)
}
