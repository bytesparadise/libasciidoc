package xhtml5_test

import (
	"time"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2" // nolint:golint
	. "github.com/onsi/gomega"    // nolint:golintt
)

var _ = Describe("document header", func() {

	It("header with quoted text", func() {
		source := `= The _Document_ *Title*`
		expectedTmpl := `<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml" lang="en">
<head>
<meta charset="UTF-8"/>
<meta http-equiv="X-UA-Compatible" content="IE=edge"/>
<meta name="viewport" content="width=device-width, initial-scale=1.0"/>
<meta name="generator" content="libasciidoc"/>
<link type="text/css" rel="stylesheet" href="/path/to/style.css"/>
<title>The Document Title</title>
</head>
<body class="article">
<div id="header">
<h1>The <em>Document</em> <strong>Title</strong></h1>
</div>
<div id="content">
</div>
<div id="footer">
<div id="footer-text">
Last updated {{ .LastUpdated }}
</div>
</div>
</body>
</html>
`
		now := time.Now()
		Expect(RenderXHTML(source, configuration.WithHeaderFooter(true),
			configuration.WithCSS([]string{"/path/to/style.css"}),
			configuration.WithLastUpdated(now),
		)).To(MatchHTMLTemplate(expectedTmpl,
			struct {
				LastUpdated string
			}{
				LastUpdated: now.Format(configuration.LastUpdatedFormat),
			}))
	})

	It("header with role", func() {
		source := `[.my_role]
= My Title`
		expectedTmpl := `<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml" lang="en">
<head>
<meta charset="UTF-8"/>
<meta http-equiv="X-UA-Compatible" content="IE=edge"/>
<meta name="viewport" content="width=device-width, initial-scale=1.0"/>
<meta name="generator" content="libasciidoc"/>
<link type="text/css" rel="stylesheet" href="/path/to/style.css"/>
<title>My Title</title>
</head>
<body class="article my_role">
<div id="header">
<h1>My Title</h1>
</div>
<div id="content">
</div>
<div id="footer">
<div id="footer-text">
Last updated {{ .LastUpdated }}
</div>
</div>
</body>
</html>
`
		now := time.Now()
		Expect(RenderXHTML(source, configuration.WithHeaderFooter(true),
			configuration.WithCSS([]string{"/path/to/style.css"}),
			configuration.WithLastUpdated(now),
		)).To(MatchHTMLTemplate(expectedTmpl,
			struct {
				LastUpdated string
			}{
				LastUpdated: now.Format(configuration.LastUpdatedFormat),
			}))
	})

	It("header with multple roles and id", func() {
		source := `[.role1#anchor.role2]
= My Title`
		expectedTmpl := `<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml" lang="en">
<head>
<meta charset="UTF-8"/>
<meta http-equiv="X-UA-Compatible" content="IE=edge"/>
<meta name="viewport" content="width=device-width, initial-scale=1.0"/>
<meta name="generator" content="libasciidoc"/>
<link type="text/css" rel="stylesheet" href="/path/to/style.css"/>
<title>My Title</title>
</head>
<body id="anchor" class="article role1 role2">
<div id="header">
<h1>My Title</h1>
</div>
<div id="content">
</div>
<div id="footer">
<div id="footer-text">
Last updated {{ .LastUpdated }}
</div>
</div>
</body>
</html>
`
		now := time.Now()
		Expect(RenderXHTML(source, configuration.WithHeaderFooter(true),
			configuration.WithCSS([]string{"/path/to/style.css"}),
			configuration.WithLastUpdated(now),
		)).To(MatchHTMLTemplate(expectedTmpl,
			struct {
				LastUpdated string
			}{
				LastUpdated: now.Format(configuration.LastUpdatedFormat),
			}))
	})

	It("should include adoc file without leveloffset from relative file", func() {
		source := "include::../../../../../test/includes/grandchild-include.adoc[]" // with filename `tmp/foo.adoc`, we are virtually in a subfolder
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
		Expect(RenderXHTML(source,
			configuration.WithFilename("tmp/foo.adoc"),
		)).To(MatchHTML(expected))
	})

	It("document with custom icon attributes", func() {
		// given
		attrs := map[string]interface{}{
			"icons":              "font",
			"source-highlighter": "pygments",
		}
		source := `[source]
----
foo
----

NOTE: a note`
		expected := `<div class="listingblock">
<div class="content">
<pre class="pygments highlight"><code>foo</code></pre>
</div>
</div>
<div class="admonitionblock note">
<table>
<tr>
<td class="icon">
<i class="fa icon-note" title="Note"></i>
</td>
<td class="content">
a note
</td>
</tr>
</table>
</div>
`
		Expect(RenderXHTML(source,
			configuration.WithAttributes(attrs),
		)).To(MatchHTML(expected))
	})

	It("document without custom icon attributes", func() {
		// given
		attrs := map[string]interface{}{}
		source := `[source]
----
foo
----

NOTE: a note`
		expected := `<div class="listingblock">
<div class="content">
<pre class="highlight"><code>foo</code></pre>
</div>
</div>
<div class="admonitionblock note">
<table>
<tr>
<td class="icon">
<div class="title">Note</div>
</td>
<td class="content">
a note
</td>
</tr>
</table>
</div>
`
		Expect(RenderXHTML(source,
			configuration.WithAttributes(attrs),
		)).To(MatchHTML(expected))
	})

	It("render manpage document with header and footer", func() {

		source := `= eve(1)
Andrew Stanton
v1.0.0

== Name

eve - analyzes an image to determine if it's a picture of a life form

== Synopsis

*eve* [_OPTION_]... _FILE_...

== Options

*-o, --out-file*=_OUT_FILE_::
Write result to file _OUT_FILE_.

*-c, --capture*::
Capture specimen if it's a picture of a life form.

== Exit status

*0*::
Success.
Image is a picture of a life form.

*1*::
Failure.
Image is not a picture of a life form.

== Resources

*Project web site:* http://eve.example.com

== Copying

Copyright (C) 2008 {author}. +
Free use of this software is granted under the terms of the MIT License.`

		expectedTmpl := `<!DOCTYPE html>
<html xmlns="http://www.w3.org/1999/xhtml" lang="en">
<head>
<meta charset="UTF-8"/>
<meta http-equiv="X-UA-Compatible" content="IE=edge"/>
<meta name="viewport" content="width=device-width, initial-scale=1.0"/>
<meta name="generator" content="libasciidoc"/>
<meta name="author" content="Andrew Stanton"/>
<link type="text/css" rel="stylesheet" href="/path/to/style.css"/>
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
<h2 id="_options">Options</h2>
<div class="sectionbody">
<div class="dlist">
<dl>
<dt class="hdlist1"><strong>-o, --out-file</strong>=<em>OUT_FILE</em></dt>
<dd>
<p>Write result to file <em>OUT_FILE</em>.</p>
</dd>
<dt class="hdlist1"><strong>-c, --capture</strong></dt>
<dd>
<p>Capture specimen if it&#8217;s a picture of a life form.</p>
</dd>
</dl>
</div>
</div>
</div>
<div class="sect1">
<h2 id="_exit_status">Exit status</h2>
<div class="sectionbody">
<div class="dlist">
<dl>
<dt class="hdlist1"><strong>0</strong></dt>
<dd>
<p>Success.
Image is a picture of a life form.</p>
</dd>
<dt class="hdlist1"><strong>1</strong></dt>
<dd>
<p>Failure.
Image is not a picture of a life form.</p>
</dd>
</dl>
</div>
</div>
</div>
<div class="sect1">
<h2 id="_resources">Resources</h2>
<div class="sectionbody">
<div class="paragraph">
<p><strong>Project web site:</strong> <a href="http://eve.example.com" class="bare">http://eve.example.com</a></p>
</div>
</div>
</div>
<div class="sect1">
<h2 id="_copying">Copying</h2>
<div class="sectionbody">
<div class="paragraph">
<p>Copyright &#169; 2008 Andrew Stanton.<br/>
Free use of this software is granted under the terms of the MIT License.</p>
</div>
</div>
</div>
</div>
<div id="footer">
<div id="footer-text">
Version 1.0.0<br/>
Last updated {{ .LastUpdated }}
</div>
</div>
</body>
</html>
`
		now := time.Now()
		Expect(RenderXHTML(source,
			configuration.WithAttributes(map[string]interface{}{
				types.AttrDocType: "manpage",
			}),
			configuration.WithCSS([]string{"/path/to/style.css"}),
			configuration.WithLastUpdated(now),
			configuration.WithHeaderFooter(true),
		)).To(MatchHTMLTemplate(expectedTmpl,
			struct {
				LastUpdated string
			}{
				LastUpdated: now.Format(configuration.LastUpdatedFormat),
			}))
	})

	It("render manpage document without header and footer", func() {

		source := `= eve(1)
Andrew Stanton
v1.0.0

== Name

eve - analyzes an image to determine if it's a picture of a life form

== Synopsis

*eve* [_OPTION_]... _FILE_...

== Options

*-o, --out-file*=_OUT_FILE_::
Write result to file _OUT_FILE_.

*-c, --capture*::
Capture specimen if it's a picture of a life form.

== Exit status

*0*::
Success.
Image is a picture of a life form.

*1*::
Failure.
Image is not a picture of a life form.

== Resources

*Project web site:* http://eve.example.com

== Copying

Copyright (C) 2008 {author}. +
Free use of this software is granted under the terms of the MIT License.`

		expectedTmpl := `<h2 id="_name">Name</h2>
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
<h2 id="_options">Options</h2>
<div class="sectionbody">
<div class="dlist">
<dl>
<dt class="hdlist1"><strong>-o, --out-file</strong>=<em>OUT_FILE</em></dt>
<dd>
<p>Write result to file <em>OUT_FILE</em>.</p>
</dd>
<dt class="hdlist1"><strong>-c, --capture</strong></dt>
<dd>
<p>Capture specimen if it&#8217;s a picture of a life form.</p>
</dd>
</dl>
</div>
</div>
</div>
<div class="sect1">
<h2 id="_exit_status">Exit status</h2>
<div class="sectionbody">
<div class="dlist">
<dl>
<dt class="hdlist1"><strong>0</strong></dt>
<dd>
<p>Success.
Image is a picture of a life form.</p>
</dd>
<dt class="hdlist1"><strong>1</strong></dt>
<dd>
<p>Failure.
Image is not a picture of a life form.</p>
</dd>
</dl>
</div>
</div>
</div>
<div class="sect1">
<h2 id="_resources">Resources</h2>
<div class="sectionbody">
<div class="paragraph">
<p><strong>Project web site:</strong> <a href="http://eve.example.com" class="bare">http://eve.example.com</a></p>
</div>
</div>
</div>
<div class="sect1">
<h2 id="_copying">Copying</h2>
<div class="sectionbody">
<div class="paragraph">
<p>Copyright &#169; 2008 Andrew Stanton.<br/>
Free use of this software is granted under the terms of the MIT License.</p>
</div>
</div>
</div>
`
		now := time.Now()
		Expect(RenderXHTML(source,
			configuration.WithAttributes(map[string]interface{}{
				types.AttrDocType: "manpage",
			}),
			configuration.WithCSS([]string{"/path/to/style.css"}),
			configuration.WithLastUpdated(now),
			configuration.WithHeaderFooter(false),
		)).To(MatchHTMLTemplate(expectedTmpl,
			struct {
				LastUpdated string
			}{
				LastUpdated: now.Format(configuration.LastUpdatedFormat),
			}))
	})
})
