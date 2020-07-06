package libasciidoc_test

import (
	"os"
	"time"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
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

	lastUpdated := time.Now()

	Context("article", func() {

		Context("document body", func() {

			It("empty document", func() {
				// main title alone is not rendered in the body
				source := ""
				expectedContent := ""
				Expect(RenderHTML(source)).To(Equal(expectedContent))
				Expect(RenderHTML5Title(source)).To(Equal(""))
			})

			It("document with no section", func() {
				// main title alone is not rendered in the body
				source := "= a document title"
				expectedTitle := "a document title"
				expectedContent := ""
				Expect(RenderHTML(source)).To(Equal(expectedContent))
				Expect(RenderHTML5Title(source)).To(Equal(expectedTitle))
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
</div>
`
				Expect(RenderHTML(source)).To(Equal(expectedContent))
				Expect(RenderHTML5Title(source)).To(Equal(expectedTitle))
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
</div>
`
				Expect(RenderHTML(source)).To(Equal(expectedContent))
				Expect(RenderHTML5Title(source)).To(Equal(""))
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
</div>
`
				Expect(RenderHTML(source)).To(Equal(expectedContent))
				Expect(RenderHTML5Title(source)).To(Equal(expectedTitle))
				Expect(DocumentMetadata(source, lastUpdated)).To(Equal(types.Metadata{
					Title:       "a document title",
					LastUpdated: lastUpdated.Format(configuration.LastUpdatedFormat),
					TableOfContents: types.TableOfContents{
						Sections: []types.ToCSection{
							{
								ID:    "_section_a",
								Level: 1,
								Title: "Section A",
								Children: []types.ToCSection{
									{
										ID:       "_section_a_a_a",
										Level:    3,
										Title:    "Section A.a.a",
										Children: []types.ToCSection{},
									},
								},
							},
						},
					},
				}))
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
</div>
`
				Expect(RenderHTML(source)).To(Equal(expectedContent))
				Expect(RenderHTML5Title(source)).To(Equal(expectedTitle))
				Expect(DocumentMetadata(source, lastUpdated)).To(Equal(types.Metadata{
					Title:       "a document title",
					LastUpdated: lastUpdated.Format(configuration.LastUpdatedFormat),
					TableOfContents: types.TableOfContents{
						Sections: []types.ToCSection{
							{
								ID:    "_section_a",
								Level: 1,
								Title: "Section A",
								Children: []types.ToCSection{
									{
										ID:       "_section_a_a",
										Level:    2,
										Title:    "Section A.a",
										Children: []types.ToCSection{},
									},
								},
							},
							{
								ID:       "_section_b",
								Level:    1,
								Title:    "Section B",
								Children: []types.ToCSection{},
							},
						},
					},
				}))
			})

			It("should include adoc file without leveloffset from local file", func() {
				source := "include::test/includes/grandchild-include.adoc[]"
				expected := `<div class="sect1">
<h2 id="_grandchild_title">grandchild title</h2>
<div class="sectionbody">
<div class="paragraph">
<p>first line of grandchild</p>
</div>
<div class="paragraph">
<p>last line of grandchild</p>
</div>
</div>
</div>
`
				Expect(RenderHTML(source, configuration.WithFilename("test.adoc"))).To(Equal(expected))
				Expect(DocumentMetadata(source, lastUpdated)).To(Equal(types.Metadata{
					Title:       "",
					LastUpdated: lastUpdated.Format(configuration.LastUpdatedFormat),
					TableOfContents: types.TableOfContents{
						Sections: []types.ToCSection{
							{
								ID:       "_grandchild_title",
								Level:    1,
								Title:    "grandchild title",
								Children: []types.ToCSection{},
							},
						},
					},
				}))
			})
		})

		Context("complete Document ", func() {

			It("using existing file", func() {
				expectedContent := `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<meta name="generator" content="libasciidoc">
<link type="text/css" rel="stylesheet" href="path/to/style.css">
<title>Chapter A</title>
</head>
<body class="article">
<div id="header">
<h1>Chapter A</h1>
</div>
<div id="content">
<div class="paragraph">
<p>content</p>
</div>
</div>
<div id="footer">
<div id="footer-text">
Last updated {{.LastUpdated}}
</div>
</div>
</body>
</html>`
				filename := "test/includes/chapter-a.adoc"
				stat, err := os.Stat(filename)
				Expect(err).NotTo(HaveOccurred())
				Expect(RenderHTML5Document(filename, configuration.WithCSS("path/to/style.css"), configuration.WithHeaderFooter(true))).To(MatchHTMLTemplate(expectedContent, stat.ModTime()))
			})
		})
	})

	Context("manpage", func() {

		Context("document body", func() {

			It("should render valid manpage", func() {
				source := `= eve(1)
Andrew Stanton
v1.0.0

== Name

eve - analyzes an image to determine if it's a picture of a life form

== Synopsis

*eve* [_OPTION_]... _FILE_...

== Copying

Copyright (C) 2008 {author}. +
Free use of this software is granted under the terms of the MIT License.`

				expectedContent := `<h2 id="_name">Name</h2>
<div class="sectionbody">
<p>eve - analyzes an image to determine if it&#8217;s a picture of a life form</p>
</div>
<div class="sect1">
<h2 id="_synopsis">Synopsis</h2>
<div class="sectionbody">
<div class="paragraph">
<p><strong>eve</strong> [<em>OPTION</em>]&#8230;&#8203; <em>FILE</em>&#8230;&#8203;</p>
</div>
</div>
</div>
<div class="sect1">
<h2 id="_copying">Copying</h2>
<div class="sectionbody">
<div class="paragraph">
<p>Copyright &#169; 2008 Andrew Stanton.<br>
Free use of this software is granted under the terms of the MIT License.</p>
</div>
</div>
</div>
`
				Expect(RenderHTML(source, configuration.WithAttribute(types.AttrDocType, "manpage"))).To(Equal(expectedContent))
			})
		})

		Context("full document", func() {

			It("should render valid manpage", func() {
				source := `= eve(1)
Andrew Stanton
v1.0.0

== Name

eve - analyzes an image to determine if it's a picture of a life form

== Synopsis

*eve* [_OPTION_]... _FILE_...

== Copying

Copyright (C) 2008 {author}. +
Free use of this software is granted under the terms of the MIT License.`

				expectedContent := `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<meta name="generator" content="libasciidoc">
<meta name="author" content="Andrew Stanton">
<link type="text/css" rel="stylesheet" href="path/to/style.css">
<title>eve(1)</title>
</head>
<body class="manpage">
<div id="header">
<h1>eve(1) Manual Page</h1>
<h2 id="_name">Name</h2>
<div class="sectionbody">
<p>eve - analyzes an image to determine if it&#8217;s a picture of a life form</p>
</div>
</div>
<div id="content">
<div class="sect1">
<h2 id="_synopsis">Synopsis</h2>
<div class="sectionbody">
<div class="paragraph">
<p><strong>eve</strong> [<em>OPTION</em>]&#8230;&#8203; <em>FILE</em>&#8230;&#8203;</p>
</div>
</div>
</div>
<div class="sect1">
<h2 id="_copying">Copying</h2>
<div class="sectionbody">
<div class="paragraph">
<p>Copyright &#169; 2008 Andrew Stanton.<br>
Free use of this software is granted under the terms of the MIT License.</p>
</div>
</div>
</div>
</div>
<div id="footer">
<div id="footer-text">
Version 1.0.0<br>
Last updated {{.LastUpdated}}
</div>
</div>
</body>
</html>`
				Expect(RenderHTML(source,
					configuration.WithAttribute(types.AttrDocType, "manpage"),
					configuration.WithLastUpdated(lastUpdated),
					configuration.WithCSS("path/to/style.css"),
					configuration.WithHeaderFooter(true))).To(MatchHTMLTemplate(expectedContent, lastUpdated))
			})

			It("should render invalid manpage as article", func() {
				source := `= eve(1)
Andrew Stanton
v1.0.0

== Foo

eve - analyzes an image to determine if it's a picture of a life form

== Synopsis

*eve* [_OPTION_]... _FILE_...

== Copying

Copyright (C) 2008 {author}. +
Free use of this software is granted under the terms of the MIT License.`

				expectedContent := `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<meta name="generator" content="libasciidoc">
<meta name="author" content="Andrew Stanton">
<link type="text/css" rel="stylesheet" href="path/to/style.css">
<title>eve(1)</title>
</head>
<body class="article">
<div id="header">
<h1>eve(1)</h1>
<div class="details">
<span id="author" class="author">Andrew Stanton</span><br>
<span id="revnumber">version 1.0.0</span>
</div>
</div>
<div id="content">
<div class="sect1">
<h2 id="_foo">Foo</h2>
<div class="sectionbody">
<div class="paragraph">
<p>eve - analyzes an image to determine if it&#8217;s a picture of a life form</p>
</div>
</div>
</div>
<div class="sect1">
<h2 id="_synopsis">Synopsis</h2>
<div class="sectionbody">
<div class="paragraph">
<p><strong>eve</strong> [<em>OPTION</em>]&#8230;&#8203; <em>FILE</em>&#8230;&#8203;</p>
</div>
</div>
</div>
<div class="sect1">
<h2 id="_copying">Copying</h2>
<div class="sectionbody">
<div class="paragraph">
<p>Copyright &#169; 2008 Andrew Stanton.<br>
Free use of this software is granted under the terms of the MIT License.</p>
</div>
</div>
</div>
</div>
<div id="footer">
<div id="footer-text">
Version 1.0.0<br>
Last updated {{.LastUpdated}}
</div>
</div>
</body>
</html>`
				Expect(RenderHTML(source,
					configuration.WithAttribute(types.AttrDocType, "manpage"),
					configuration.WithLastUpdated(lastUpdated),
					configuration.WithCSS("path/to/style.css"),
					configuration.WithHeaderFooter(true))).To(MatchHTMLTemplate(expectedContent, lastUpdated))
			})

			It("should render html", func() {
				source := `= Story

Our story begins.`
				expectedContent := `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<meta name="generator" content="libasciidoc">
<title>Story</title>
</head>
<body class="article">
<div id="header">
<h1>Story</h1>
</div>
<div id="content">
<div class="paragraph">
<p>Our story begins.</p>
</div>
</div>
<div id="footer">
<div id="footer-text">
Last updated {{.LastUpdated}}
</div>
</div>
</body>
</html>`
				Expect(Render(source,
					configuration.WithBackEnd("html5"),
					configuration.WithLastUpdated(lastUpdated),
					configuration.WithHeaderFooter(true))).To(MatchHTMLTemplate(expectedContent, lastUpdated))
			})

			It("should render xhtml", func() {
				source := `= Story

Our story begins.`
				expectedContent := `<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml" lang="en">
<head>
<meta charset="UTF-8"/>
<meta http-equiv="X-UA-Compatible" content="IE=edge"/>
<meta name="viewport" content="width=device-width, initial-scale=1.0"/>
<meta name="generator" content="libasciidoc"/>
<title>Story</title>
</head>
<body class="article">
<div id="header">
<h1>Story</h1>
</div>
<div id="content">
<div class="paragraph">
<p>Our story begins.</p>
</div>
</div>
<div id="footer">
<div id="footer-text">
Last updated {{.LastUpdated}}
</div>
</div>
</body>
</html>`
				Expect(Render(source,
					configuration.WithBackEnd("xhtml5"),
					configuration.WithLastUpdated(lastUpdated),
					configuration.WithHeaderFooter(true))).To(MatchHTMLTemplate(expectedContent, lastUpdated))
			})

			It("should fail given bogus backend", func() {
				source := `= Story

Our story begins.`
				doc, err := Render(source,
					configuration.WithBackEnd("wordperfect"),
					configuration.WithLastUpdated(lastUpdated),
					configuration.WithHeaderFooter(true))
				Expect(doc).To(BeEmpty())
				Expect(err).To(MatchError("backend 'wordperfect' not supported"))
			})

		})

	})
})
