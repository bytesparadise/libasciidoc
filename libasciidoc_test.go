package libasciidoc_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"
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
			Expect(source).To(RenderHTML5Body(expectedContent))
			Expect(source).To(RenderHTML5Title(nil))
		})

		It("document with no section", func() {
			// main title alone is not rendered in the body
			source := "= a document title"
			expectedTitle := "a document title"
			expectedContent := ""
			Expect(source).To(RenderHTML5Body(expectedContent))
			Expect(source).To(RenderHTML5Title(expectedTitle))
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
			Expect(source).To(RenderHTML5Body(expectedContent))
			Expect(source).To(RenderHTML5Title(expectedTitle))
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
			Expect(source).To(RenderHTML5Body(expectedContent))
			Expect(source).To(RenderHTML5Title(nil))
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
			Expect(source).To(RenderHTML5Body(expectedContent))
			Expect(source).To(RenderHTML5Title(expectedTitle))
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
			Expect(source).To(RenderHTML5Body(expectedContent))
			Expect(source).To(RenderHTML5Title(expectedTitle))
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
			Expect(source).To(RenderHTML5Document(expectedContent))
		})
	})

})
