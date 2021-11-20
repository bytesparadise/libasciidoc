// +build bench

package parser_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/bytesparadise/libasciidoc/pkg/parser"

	. "github.com/onsi/ginkgo"                  //nolint golint
	. "github.com/onsi/ginkgo/extensions/table" //nolint golint
	. "github.com/onsi/gomega"                  //nolint golint
)

const (
	doc1Paragraph = `= Lorem Ipsum
	
Lorem ipsum dolor sit amet, consetetur sadipscing elitr, 
sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, 
sed diam voluptua. 
At vero eos et accusam et justo duo dolores et ea rebum. 
Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. 
Lorem ipsum dolor sit amet, consetetur sadipscing elitr, 
sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, 
sed diam voluptua. 
At vero eos et accusam et justo duo dolores et ea rebum. 
Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.`

	doc10Sections = `=== foo
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
)

var _ = DescribeTable("basic stats",
	func(title, content string) {
		stats := parser.Stats{}
		_, err := parser.Parse(title, []byte(content), parser.Statistics(&stats, "no match")) // , parser.Debug(true))
		Expect(err).NotTo(HaveOccurred())
		fmt.Printf("%s\n", title)
		fmt.Printf("ExprCnt:      %d\n", stats.ExprCnt)
		result, _ := json.MarshalIndent(stats.ChoiceAltCnt, " ", " ")
		fmt.Printf("ChoiceAltCnt: \n%s\n", result)
	},
	Entry("parse a single line file", "1-line doc", doc1Paragraph),
	Entry("parse a 10-line file", "10-lines doc", doc10Sections),
)

var _ = Describe("real-world doc-based benchmarks", func() {

	Measure("parse the vert.x examples doc", func(b Benchmarker) {
		filename := "../../test/bench/vertx-examples.adoc"
		content, err := load(filename)
		Expect(err).NotTo(HaveOccurred())
		b.Time("runtime", func() {
			_, err := parser.Parse(filename, content)
			Expect(err).NotTo(HaveOccurred())
		})
	}, 10)

	Measure("parse the quarkus kafka streams doc", func(b Benchmarker) {
		filename := "../../test/bench/kafka-streams.adoc"
		content, err := load(filename)
		Expect(err).NotTo(HaveOccurred())
		b.Time("runtime", func() {
			_, err := parser.Parse(filename, content)
			Expect(err).NotTo(HaveOccurred())
		})
	}, 10)

})

func load(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		f.Close()
	}()
	return ioutil.ReadAll(f)
}
