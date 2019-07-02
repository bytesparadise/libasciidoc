package parser_test

import (
	"io/ioutil"
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
		})
		timeout := 0.1
		if ci {
			timeout *= 10
		}
		Expect(runtime.Seconds()).Should(BeNumerically("<", timeout), "parsing shouldn't take too long (even on CI).")

	}, 1)

	Measure("bench parser on 'vert.x examples' doc", func(b Benchmarker) {
		f, err := os.Open("../../test/bench/vertx-examples.adoc")
		Expect(err).ShouldNot(HaveOccurred())
		defer func() {
			err := f.Close()
			Expect(err).ShouldNot(HaveOccurred())
		}()
		content, err := ioutil.ReadAll(f)
		Expect(err).ShouldNot(HaveOccurred())
		runtime := b.Time("runtime", func() {
			// given
			_, err := parser.Parse("vert.x samples", content)
			Expect(err).ShouldNot(HaveOccurred())
		})
		timeout := 0.2 * 50
		if ci {
			timeout *= 10
		}
		Expect(runtime.Seconds()).Should(BeNumerically("<", timeout), "parsing shouldn't take too long (even on CI).")

	}, 50)

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
