package parser_test

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/davecgh/go-spew/spew"
	. "github.com/onsi/ginkgo"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var _ = Describe("document preprocessing", func() {

	Context("preserve attribute declarations", func() {

		It("should return raw lines with relative level offset", func() {
			actualContent := `:assets: ./includes
some text 

and here, too`
			expectedContent := `:assets: ./includes
some text 

and here, too`
			verifyPreprocessing(expectedContent, actualContent)
		})
	})

	Context("simple file inclusions", func() {

		It("should return raw lines without file inclusion", func() {
			actualContent := `some text 

and here, too`
			expectedContent := `some text 

and here, too`
			verifyPreprocessing(expectedContent, actualContent)
		})

		It("should return raw lines with simple file inclusion", func() {
			actualContent := `some text 

include::./includes/chapter-a.adoc[]

and here, too`
			expectedContent := `some text 

= Chapter A

content

and here, too`
			verifyPreprocessing(expectedContent, actualContent)
		})
	})

	Context("file inclusions in delimited blocks", func() {

		It("should return raw lines with relative level offset", func() {
			actualContent := `____
include::includes/chapter-a.adoc[]
____`
			expectedContent := `____
= Chapter A

content
____`
			verifyPreprocessing(expectedContent, actualContent)
		})

	})
	Context("file inclusions with level offset", func() {

		It("should return raw lines with relative level offset", func() {
			actualContent := `some text 

include::./includes/chapter-a.adoc[leveloffset=+1]

and here, too`
			expectedContent := `some text 

== Chapter A

content

and here, too`
			verifyPreprocessing(expectedContent, actualContent)
		})
	})

	Context("file inclusions with recursion", func() {

		It("should return raw lines nested inclusions", func() {
			actualContent := `some text 

include::./includes/parent-include.adoc[]

and here, too`
			// note: there will be a single empty line after the inclusion
			expectedContent := `some text 

first line of parent

first line of child

first line of grandchild

last line of grandchild

last line of child

last line of parent

and here, too`
			verifyPreprocessing(expectedContent, actualContent)
		})
	})

	Context("file inclusions with line ranges", func() {

		It("should return raw lines with simple file inclusion with single range", func() {
			actualContent := `some text 

include::./includes/chapter-a.adoc[lines=2..4]

and here, too`
			// note: there will be a single empty line after the inclusion
			expectedContent := `some text 


content

and here, too`
			verifyPreprocessing(expectedContent, actualContent)
		})

		It("should return raw lines with simple file inclusion with multiple ranges", func() {
			actualContent := `some text 

include::./includes/chapter-a.adoc[lines=1;3..4]

and here, too`
			// note: there will be a single empty line after the inclusion
			expectedContent := `some text 

= Chapter A
content

and here, too`
			verifyPreprocessing(expectedContent, actualContent)
		})

	})

})

func verifyPreprocessing(expected, actual string) {
	log.Debugf("processing: %s", actual)
	reader := strings.NewReader(actual)
	result, err := parser.PreparseDocument("", reader)
	if err != nil {
		log.WithError(err).Errorf("Error found while parsing the document (%T)", err)
	}
	require.NoError(GinkgoT(), err)
	GinkgoT().Logf("actual document: `%s`", spew.Sdump(string(result)))
	GinkgoT().Logf("expected document: `%s`", spew.Sdump(expected))
	assert.EqualValues(GinkgoT(), expected, string(result))
}
