package parser_test

import (
	"os"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/parser"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("parser benchmark", func() {

	Measure("bench parser on 10 lines", func(b Benchmarker) {
		runtime := b.Time("runtime", func() {
			// given
			actualContent := `=== foo
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
			stats, err := parseReader(actualContent)
			Expect(err).ShouldNot(HaveOccurred())
			b.RecordValue("expr count", float64(stats.ExprCnt))
		})

		Expect(runtime.Seconds()).Should(BeNumerically("<", 0.5), "parsing shouldn't take too long (even on CI).")

	}, 10)

	Measure("bench parser on 2 lines", func(b Benchmarker) {
		runtime := b.Time("runtime", func() {
			// given
			actualContent := `=== foo1
bar1

=== foo2
bar2`
			stats, err := parseReader(actualContent)
			Expect(err).ShouldNot(HaveOccurred())
			b.RecordValue("expr count", float64(stats.ExprCnt))
		})

		Expect(runtime.Seconds()).Should(BeNumerically("<", 0.1), "parsing shouldn't take too long (even on CI).")

	}, 10)

	Measure("bench parser on 1 line", func(b Benchmarker) {
		runtime := b.Time("runtime", func() {
			// given
			actualContent := `=== foo1
bar1`
			stats, err := parseReader(actualContent)
			Expect(err).ShouldNot(HaveOccurred())
			b.RecordValue("expr count", float64(stats.ExprCnt))
		})

		Expect(runtime.Seconds()).Should(BeNumerically("<", 0.1), "parsing shouldn't take too long (even on CI).")

	}, 1)

})

func parseReader(content string) (parser.Stats, error) {
	reader := strings.NewReader(content)
	stats := parser.Stats{}
	allOptions := make([]parser.Option, 0)
	allOptions = append(allOptions, parser.AllowInvalidUTF8(false), parser.Statistics(&stats, "no match"))
	if os.Getenv("DEBUG") == "true" {
		allOptions = append(allOptions, parser.Debug(true))
	}
	_, err := parser.ParseReader("", reader, allOptions...) //, Debug(true))
	return stats, err
}
