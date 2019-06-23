package parser_test

import (
	"os"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/parser"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("parser benchmark", func() {

	ci := os.Getenv("CI") != ""

	Measure("bench parser on 10 lines", func(b Benchmarker) {
		runtime := b.Time("runtime", func() {
			// given
			source := `=== foo
bar

=== foo
bar

=== foo
bar

=== foo
bar

=== foo
bar

=== foo
bar

=== foo
bar

=== foo
bar

=== foo
bar

=== foo
bar`
			err := parseReader(source)
			Expect(err).ShouldNot(HaveOccurred())
			// b.RecordValue("expr count", float64(stats.ExprCnt))
		})
		timeout := 0.5
		if ci {
			timeout *= 10
		}
		Expect(runtime.Seconds()).Should(BeNumerically("<", timeout), "parsing shouldn't take too long (even on CI).")

	}, 10)

	Measure("bench parser on 1 line", func(b Benchmarker) {
		runtime := b.Time("runtime", func() {
			// given
			source := `=== foo1
bar1`
			err := parseReader(source)
			Expect(err).ShouldNot(HaveOccurred())
			// b.RecordValue("expr count", float64(stats.ExprCnt))
		})
		timeout := 0.1
		if ci {
			timeout *= 10
		}
		Expect(runtime.Seconds()).Should(BeNumerically("<", timeout), "parsing shouldn't take too long (even on CI).")

	}, 1)

	Measure("bench parser on basic doc", func(b Benchmarker) {
		runtime := b.Time("runtime", func() {
			// given
			source := `= Introduction to AsciiDoc
Doc Writer <doc@example.com>

A preface about https://asciidoc.org[AsciiDoc].

== First Section

* item 1
* item 2

[source,ruby]
puts "Hello, World!"
`
			err := parseReader(source)
			Expect(err).ShouldNot(HaveOccurred())
			// b.RecordValue("expr count", float64(stats.ExprCnt))
		})
		timeout := 0.1
		if ci {
			timeout *= 10
		}
		Expect(runtime.Seconds()).Should(BeNumerically("<", timeout), "parsing shouldn't take too long (even on CI).")

	}, 1)

})

func parseReader(content string) error {
	reader := strings.NewReader(content)
	// stats := parser.Stats{}
	// opts := []parser.Option{parser.AllowInvalidUTF8(false), parser.Statistics(&stats, "no match")}
	opts := []parser.Option{parser.AllowInvalidUTF8(false)}
	// if os.Getenv("DEBUG") == "true" {
	// 	opts = append(opts, parser.Debug(true))
	// }
	_, err := parser.ParseDocument("", reader, opts...) //, Debug(true))
	return err
}
