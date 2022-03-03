package libasciidoc_test

import (
	"os"
	"strings"
	"time"

	"github.com/bytesparadise/libasciidoc"
	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" // nolint:golint
	. "github.com/onsi/gomega" // nolint:golintt
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

	Context("article documents", func() {

		Context("with body output only", func() {

			It("should render empty document", func() {
				// main title alone is not rendered in the body
				source := ""
				expected := ""
				output, metadata, err := RenderHTMLWithMetadata(source)
				Expect(err).NotTo(HaveOccurred())
				Expect(output).To(MatchHTML(expected))
				Expect(metadata.Title).To(BeEmpty())
			})

			It("should render demo.adoc", func() {
				// given
				info, err := os.Stat("test/compat/demo.adoc")
				Expect(err).NotTo(HaveOccurred())

				// when
				output, metadata, err := RenderHTMLFromFile("test/compat/demo.adoc",
					configuration.WithAttribute("libasciidoc-version", "0.7.0"),
					configuration.WithCSS("path/to/style.css"))

				// then
				Expect(err).NotTo(HaveOccurred())
				Expect(output).To(MatchHTMLFromFile("test/compat/demo.html"))
				Expect(metadata).To(MatchMetadata(types.Metadata{
					Title:       "Libasciidoc Demo",
					LastUpdated: info.ModTime().Format(configuration.LastUpdatedFormat),
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:     "first_steps_with_asciidoc",
								Level:  1,
								Title:  "First Steps with AsciiDoc",
								Number: "1",
								Children: []*types.ToCSection{
									{
										ID:     "lists_upon_lists",
										Level:  2,
										Title:  "Lists Upon Lists",
										Number: "1.1",
									},
								},
							},
							{
								ID:     "and_were_back",
								Level:  1,
								Title:  "&#8230;&#8203;and we&#8217;re back!",
								Number: "2",
								Children: []*types.ToCSection{
									{
										ID:     "block_quotes_and_smart_ones",
										Level:  2,
										Title:  "Block Quotes and &#8220;Smart&#8221; Ones",
										Number: "2.1",
									},
								},
							},
							{
								ID:     "literally",
								Level:  1,
								Title:  "Getting Literal",
								Number: "3",
							},
							{
								ID:    "wrapup",
								Level: 1, Title: "Wrap-up",
								Number: "4",
							},
						},
					},
				},
				))
			})
		})

		Context("with full output", func() {

			It("should render demo.adoc", func() {
				filename := "test/compat/demo.adoc"
				stat, err := os.Stat(filename)
				Expect(err).NotTo(HaveOccurred())

				out := &strings.Builder{}
				_, err = libasciidoc.ConvertFile(out,
					configuration.NewConfiguration(
						configuration.WithFilename(filename),
						configuration.WithAttribute("libasciidoc-version", "0.7.0"),
						configuration.WithCSS("path/to/style.css"),
						configuration.WithHeaderFooter(true)))
				Expect(err).NotTo(HaveOccurred())
				Expect(out.String()).To(MatchHTMLTemplateFile(string("test/compat/demo-full.tmpl.html"),
					struct {
						LastUpdated string
						Version     string
						CSS         string
					}{
						LastUpdated: stat.ModTime().Format(configuration.LastUpdatedFormat),
						Version:     "0.7.0",
						CSS:         "path/to/style.css",
					}))
			})

			It("not fail on article.adoc with html5 backend", func() {
				out := &strings.Builder{}
				_, err := libasciidoc.ConvertFile(out,
					configuration.NewConfiguration(
						configuration.WithFilename("test/compat/article.adoc"),
						configuration.WithBackEnd("html5"),
						configuration.WithCSS("path/to/style.css"),
						configuration.WithHeaderFooter(true)))
				Expect(err).NotTo(HaveOccurred())
			})

			It("not fail on article.adoc with xhtml5 backend", func() {
				out := &strings.Builder{}
				_, err := libasciidoc.ConvertFile(out,
					configuration.NewConfiguration(
						configuration.WithFilename("test/compat/article.adoc"),
						configuration.WithBackEnd("xhtml5"),
						configuration.WithCSS("path/to/style.css"),
						configuration.WithHeaderFooter(true)))
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})

	Context("manpage docs", func() {

		Context("with body only", func() {

			It("should render valid manpage", func() {
				_, reset := ConfigureLogger(log.WarnLevel)
				defer reset()
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

				expected := `<h2 id="_name">Name</h2>
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
				out := &strings.Builder{}
				_, err := libasciidoc.Convert(
					strings.NewReader(source),
					out,
					configuration.NewConfiguration(
						configuration.WithAttribute(types.AttrDocType, "manpage"),
					))
				Expect(err).NotTo(HaveOccurred())
				Expect(out.String()).To(MatchHTML(expected))
			})
		})

		Context("full", func() {

			It("should render valid manpage", func() {
				_, reset := ConfigureLogger(log.WarnLevel)
				defer reset()
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

				expectedTmpl := `<!DOCTYPE html>
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
Last updated {{ .LastUpdated }}
</div>
</div>
</body>
</html>
`
				out := &strings.Builder{}
				_, err := libasciidoc.Convert(
					strings.NewReader(source),
					out,
					configuration.NewConfiguration(
						configuration.WithAttribute(types.AttrDocType, "manpage"),
						configuration.WithLastUpdated(lastUpdated),
						configuration.WithCSS("path/to/style.css"),
						configuration.WithHeaderFooter(true),
					))
				Expect(err).NotTo(HaveOccurred())
				Expect(out.String()).To(MatchHTMLTemplate(expectedTmpl,
					struct {
						LastUpdated string
					}{
						LastUpdated: lastUpdated.Format(configuration.LastUpdatedFormat),
					}))
			})

			It("should render invalid manpage as article", func() {
				logs, reset := ConfigureLogger(log.WarnLevel)
				defer reset()
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

				expectedTmpl := `<!DOCTYPE html>
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
Last updated {{ .LastUpdated }}
</div>
</div>
</body>
</html>
`
				out := &strings.Builder{}
				_, err := libasciidoc.Convert(
					strings.NewReader(source),
					out,
					configuration.NewConfiguration(
						configuration.WithAttribute(types.AttrDocType, "manpage"),
						configuration.WithLastUpdated(lastUpdated),
						configuration.WithCSS("path/to/style.css"),
						configuration.WithHeaderFooter(true),
					))
				Expect(err).NotTo(HaveOccurred())
				Expect(out.String()).To(MatchHTMLTemplate(expectedTmpl,
					struct {
						LastUpdated string
					}{
						LastUpdated: lastUpdated.Format(configuration.LastUpdatedFormat),
					}))
				Expect(logs).To(ContainJSONLog(log.WarnLevel, "changing doctype to 'article' because problems were found in the document"))
				Expect(logs).To(ContainJSONLog(log.ErrorLevel, "manpage document is missing the 'Name' section"))
			})

			It("should render html", func() {
				source := `= Story

Our story begins.`
				expectedTmpl := `<!DOCTYPE html>
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
Last updated {{ .LastUpdated }}
</div>
</div>
</body>
</html>
`
				out := &strings.Builder{}
				_, err := libasciidoc.Convert(
					strings.NewReader(source),
					out,
					configuration.NewConfiguration(
						configuration.WithLastUpdated(lastUpdated),
						configuration.WithBackEnd("html5"),
						configuration.WithHeaderFooter(true),
					))
				Expect(err).NotTo(HaveOccurred())
				Expect(out.String()).To(MatchHTMLTemplate(expectedTmpl,
					struct {
						LastUpdated string
					}{
						LastUpdated: lastUpdated.Format(configuration.LastUpdatedFormat),
					}))
			})

			It("should render with xhtml5 backend", func() {
				source := `= Story

Our story begins.`
				expectedTmpl := `<!DOCTYPE html>
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
Last updated {{ .LastUpdated }}
</div>
</div>
</body>
</html>
`
				out := &strings.Builder{}
				_, err := libasciidoc.Convert(
					strings.NewReader(source),
					out,
					configuration.NewConfiguration(
						configuration.WithLastUpdated(lastUpdated),
						configuration.WithBackEnd("xhtml5"),
						configuration.WithHeaderFooter(true),
					))
				Expect(err).NotTo(HaveOccurred())
				Expect(out.String()).To(MatchHTMLTemplate(expectedTmpl,
					struct {
						LastUpdated string
					}{
						LastUpdated: lastUpdated.Format(configuration.LastUpdatedFormat),
					}))
			})

			It("should fail given bogus backend", func() {
				_, reset := ConfigureLogger(log.WarnLevel)
				defer reset()
				source := `= Story

Our story begins.`
				out := &strings.Builder{}
				_, err := libasciidoc.Convert(
					strings.NewReader(source),
					out,
					configuration.NewConfiguration(
						configuration.WithBackEnd("wordperfect"),
						configuration.WithLastUpdated(lastUpdated),
						configuration.WithHeaderFooter(true),
					))
				Expect(err).To(MatchError("backend 'wordperfect' not supported"))
			})
		})
	})

})
