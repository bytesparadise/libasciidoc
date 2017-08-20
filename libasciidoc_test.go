package libasciidoc_test

import (
	"bytes"
	"strings"

	. "github.com/bytesparadise/libasciidoc"
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var _ = Describe("Rendering documents in HTML", func() {

	It("document with no section", func() {
		// main title alone is not rendered in the body
		content := "= a document"
		expected := ""
		verify(GinkgoT(), expected, content)
	})

	It("section levels 1 and 2", func() {
		content := `= a document title

== Section A

a paragraph with *bold content*`
		expected := `<div class="sect1">
<h2 id="_section_a">Section A</h2>
<div class="sectionbody">
<div class="paragraph">
<p>a paragraph with <strong>bold content</strong></p>
</div>
</div>
</div>`
		verify(GinkgoT(), expected, content)
	})

	It("section levels 1, 2 and 3", func() {
		content := `= a document title

== Section A

a paragraph with *bold content*

=== Section A.a

a paragraph`
		expected := `<div class="sect1">
<h2 id="_section_a">Section A</h2>
<div class="sectionbody">
<div class="paragraph">
<p>a paragraph with <strong>bold content</strong></p>
</div>
<div class="sect2">
<h3 id="_section_a_a">Section A.a</h3>
<div class="paragraph">
<p>a paragraph</p>
</div>
</div>
</div>
</div>`
		verify(GinkgoT(), expected, content)
	})

	It("section levels 1, 2, 3 and 2", func() {
		content := `= a document title

== Section A

a paragraph with *bold content*

=== Section A.a

a paragraph

== Section B

a paragraph with _italic content_`
		expected := `<div class="sect1">
<h2 id="_section_a">Section A</h2>
<div class="sectionbody">
<div class="paragraph">
<p>a paragraph with <strong>bold content</strong></p>
</div>
<div class="sect2">
<h3 id="_section_a_a">Section A.a</h3>
<div class="paragraph">
<p>a paragraph</p>
</div>
</div>
</div>
</div>
<div class="sect1">
<h2 id="_section_b">Section B</h2>
<div class="sectionbody">
<div class="paragraph">
<p>a paragraph with <em>italic content</em></p>
</div>
</div>
</div>`
		verify(GinkgoT(), expected, content)
	})
})

func verify(t GinkgoTInterface, expected, content string) {
	t.Logf("processing '%s'", content)
	contentReader := strings.NewReader(content)
	resultWriter := bytes.NewBuffer(make([]byte, 0))
	err := ConvertToHTML(contentReader, resultWriter)
	require.Nil(t, err, "Error found while parsing the document")
	t.Log("Done processing document")
	result := string(resultWriter.Bytes())
	t.Logf("** Actual output:\n`%s`\n", result)
	t.Logf("** Expected output:\n`%s`\n", expected)
	assert.Equal(t, expected, result)
}
