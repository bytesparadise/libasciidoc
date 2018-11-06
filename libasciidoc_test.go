package libasciidoc_test

import (
	"bytes"
	"context"
	"strings"
	"time"

	. "github.com/bytesparadise/libasciidoc"
	_ "github.com/bytesparadise/libasciidoc/pkg/log"
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	. "github.com/onsi/ginkgo"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var _ = Describe("documents", func() {

	var level log.Level

	BeforeEach(func() {
		// turn down the logger to `warn` to avoid the noise
		level = log.GetLevel()
		log.SetLevel(log.WarnLevel)
	})

	AfterEach(func() {
		// restore the logger level
		log.SetLevel(level)
	})

	Context("document Body", func() {

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

	Context("complete Document ", func() {

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
</head>
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

func verifyDocumentBody(t GinkgoTInterface, expectedRenderedTitle *string, expectedContent, source string) {
	t.Logf("processing '%s'", source)
	sourceReader := strings.NewReader(source)
	resultWriter := bytes.NewBuffer(nil)
	metadata, err := ConvertToHTML(context.Background(), sourceReader, resultWriter, renderer.IncludeHeaderFooter(false))
	require.Nil(t, err, "Error found while parsing the document")
	require.NotNil(t, metadata)
	t.Log("Done processing document")
	result := resultWriter.String()
	t.Logf("** Actual output:\n`%s`\n", result)
	t.Logf("** expectedContent output:\n`%s`\n", expectedContent)
	assert.Equal(t, expectedContent, result)
	actualTitle := metadata["doctitle"]
	if expectedRenderedTitle == nil {
		assert.Nil(t, actualTitle)
	} else {
		t.Logf("Actual title: %v", actualTitle)
		t.Logf("Expected title: %v", *expectedRenderedTitle)
		assert.Equal(t, *expectedRenderedTitle, actualTitle)
	}
}

func verifyCompleteDocument(t GinkgoTInterface, expectedContent, source string) {
	t.Logf("processing '%s'", source)
	sourceReader := strings.NewReader(source)
	resultWriter := bytes.NewBuffer(nil)
	lastUpdated := time.Now()
	_, err := ConvertToHTML(context.Background(), sourceReader, resultWriter, renderer.IncludeHeaderFooter(true), renderer.LastUpdated(lastUpdated))
	require.Nil(t, err, "Error found while parsing the document")
	t.Log("Done processing document")
	result := resultWriter.String()
	t.Logf("** Actual output:\n`%s`\n", result)
	require.Nil(t, err)
	expectedContent = strings.Replace(expectedContent, "{{.LastUpdated}}", lastUpdated.Format(renderer.LastUpdatedFormat), 1)
	t.Logf("** expectedContent output:\n`%s`\n", expectedContent)
	assert.Equal(t, expectedContent, result)
}
