package libasciidoc_test

import (
	"bytes"
	"context"
	"strings"
	"time"

	. "github.com/bytesparadise/libasciidoc"
	"github.com/bytesparadise/libasciidoc/renderer"
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var _ = Describe("Rendering documents in HTML", func() {

	Context("Document Body", func() {

		It("empty document", func() {
			// main title alone is not rendered in the body
			source := ""
			expectedContent := ""
			verifyDocumentBody(GinkgoT(), nil, expectedContent, source)
		})

		It("document with no section", func() {
			// main title alone is not rendered in the body
			source := "= a document title"
			expectedTitle := "a document title"
			expectedContent := ""
			verifyDocumentBody(GinkgoT(), &expectedTitle, expectedContent, source)
		})

		It("section levels 0 and 1", func() {
			source := `= a document title

== Section A

a paragraph with *bold content*`
			expectedTitle := "a document title"
			expectedContent := `<div class="sect1">
<h2 id="_section_a">Section A</h2>
<div class="sectionbody">
<div class="paragraph">
<p>a paragraph with <strong>bold content</strong></p>
</div>
</div>
</div>`
			verifyDocumentBody(GinkgoT(), &expectedTitle, expectedContent, source)
		})

		It("section level 1 with a paragraph", func() {
			source := `== Section A

a paragraph with *bold content*`
			expectedContent := `<div class="sect1">
<h2 id="_section_a">Section A</h2>
<div class="sectionbody">
<div class="paragraph">
<p>a paragraph with <strong>bold content</strong></p>
</div>
</div>
</div>`
			verifyDocumentBody(GinkgoT(), nil, expectedContent, source)
		})

		It("section levels 0, 1 and 3", func() {
			source := `= a document title

== Section A

a paragraph with *bold content*

==== Section A.a.a

a paragraph`
			expectedTitle := "a document title"
			expectedContent := `<div class="sect1">
<h2 id="_section_a">Section A</h2>
<div class="sectionbody">
<div class="paragraph">
<p>a paragraph with <strong>bold content</strong></p>
</div>
<div class="sect3">
<h4 id="_section_a_a_a">Section A.a.a</h4>
<div class="paragraph">
<p>a paragraph</p>
</div>
</div>
</div>
</div>`
			verifyDocumentBody(GinkgoT(), &expectedTitle, expectedContent, source)
		})

		It("section levels 1, 2, 3 and 2", func() {
			source := `= a document title

== Section A

a paragraph with *bold content*

=== Section A.a

a paragraph

== Section B

a paragraph with _italic content_`
			expectedTitle := "a document title"
			expectedContent := `<div class="sect1">
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
			verifyDocumentBody(GinkgoT(), &expectedTitle, expectedContent, source)
		})
	})

	Context("Complete Document ", func() {

		It("section levels 0 and 5", func() {
			source := `= a document title

====== Section A

a paragraph with *bold content*`
			expectedContent := `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<!--[if IE]><meta http-equiv="X-UA-Compatible" content="IE=edge"><![endif]-->
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<meta name="generator" content="libasciidoc">
<title>a document title</title>
<body class="article">
<div id="header">
<h1>a document title</h1>
</div>
<div id="content">
<div class="sect5">
<h6 id="_section_a">Section A</h6>
<div class="paragraph">
<p>a paragraph with <strong>bold content</strong></p>
</div>
</div>
</div>
<div id="footer">
<div id="footer-text">
Last updated {{.LastUpdated}}
</div>
</div>
</body>
</html>`
			verifyCompleteDocument(GinkgoT(), expectedContent, source)
		})
	})

})

func verifyDocumentBody(t GinkgoTInterface, expectedTitle *string, expectedContent, source string) {
	t.Logf("processing '%s'", source)
	sourceReader := strings.NewReader(source)
	resultWriter := bytes.NewBuffer(nil)
	metadata, err := ConvertToHTMLBody(context.Background(), sourceReader, resultWriter)
	require.Nil(t, err, "Error found while parsing the document")
	require.NotNil(t, metadata)
	t.Log("Done processing document")
	result := string(resultWriter.Bytes())
	t.Logf("** Actual output:\n`%s`\n", result)
	t.Logf("** expectedContent output:\n`%s`\n", expectedContent)
	assert.Equal(t, expectedContent, result)
	actualTitle := metadata.GetTitle()
	if actualTitle == nil {
		assert.Nil(t, actualTitle)
	} else {
		assert.Equal(t, *expectedTitle, *actualTitle)
	}
}

func verifyCompleteDocument(t GinkgoTInterface, expectedContent, source string) {
	t.Logf("processing '%s'", source)
	sourceReader := strings.NewReader(source)
	resultWriter := bytes.NewBuffer(nil)
	lastUpdated := time.Now()
	err := ConvertToHTML(context.Background(), sourceReader, resultWriter, renderer.LastUpdated(lastUpdated))
	require.Nil(t, err, "Error found while parsing the document")
	t.Log("Done processing document")
	result := resultWriter.String()
	t.Logf("** Actual output:\n`%s`\n", result)
	require.Nil(t, err)
	expectedContent = strings.Replace(expectedContent, "{{.LastUpdated}}", lastUpdated.Format(renderer.LastUpdatedFormat), 1)
	t.Logf("** expectedContent output:\n`%s`\n", expectedContent)
	assert.Equal(t, expectedContent, result)
}
