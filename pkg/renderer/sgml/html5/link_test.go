package html5_test

import (
	"time"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("links", func() {

	Context("bare URLs", func() {

		It("standalone URL with scheme", func() {
			source := `<https://example.com>`
			expected := `<div class="paragraph">
<p><a href="https://example.com" class="bare">https://example.com</a></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("URL with scheme in sentence", func() {
			source := `a link to <https://example.com>.`
			expected := `<div class="paragraph">
<p>a link to <a href="https://example.com" class="bare">https://example.com</a>.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("substituted URL with scheme", func() {
			source := `:example: https://example.com

a link to <{example}>.`
			expected := `<div class="paragraph">
<p>a link to <a href="https://example.com" class="bare">https://example.com</a>.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("substituted URL with scheme", func() {
			source := `:example: https://example.com

a link to <{example}>.`
			expected := `<div class="paragraph">
<p>a link to <a href="https://example.com" class="bare">https://example.com</a>.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("URL with query string", func() {
			source := `a link to https://example.com?foo=fighters&lang=en[].`
			expected := `<div class="paragraph">
<p>a link to <a href="https://example.com?foo=fighters&amp;lang=en" class="bare">https://example.com?foo=fighters&amp;lang=en</a>.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("substituted bare URL with query string", func() {
			source := `:example: https://example.com?foo=fighters&lang=en

a link to <{example}>.`
			expected := `<div class="paragraph">
<p>a link to <a href="https://example.com?foo=fighters&amp;lang=en" class="bare">https://example.com?foo=fighters&amp;lang=en</a>.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		Context("malformed", func() {

			It("should not parse URL without scheme", func() {
				source := `a link to <example.com>.`
				expected := `<div class="paragraph">
<p>a link to &lt;example.com&gt;.</p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("with invalid special character in URL", func() {
				source := `a link to https://example.com>[].`
				expected := `<div class="paragraph">
<p>a link to <a href="https://example.com" class="bare">https://example.com</a>&gt;[].</p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("with opening angle bracket", func() {
				source := `a link to <https://example.com[].`
				expected := `<div class="paragraph">
<p>a link to &lt;<a href="https://example.com" class="bare">https://example.com</a>.</p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})
		})
	})

	Context("email autolinks", func() {

		It("email in paragraph", func() {
			source := `write to contact@example.com.`
			expected := `<div class="paragraph">
<p>write to <a href="mailto:contact@example.com">contact@example.com</a>.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("external links", func() {

		It("without text", func() {

			source := "a link to https://example.com[]."
			expected := `<div class="paragraph">
<p>a link to <a href="https://example.com" class="bare">https://example.com</a>.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with trailing dot punctutation", func() {
			source := "a link to https://example.com."
			expected := `<div class="paragraph">
<p>a link to <a href="https://example.com" class="bare">https://example.com</a>.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with trailing question mark punctutation", func() {
			source := "a link to https://example.com?"
			expected := `<div class="paragraph">
<p>a link to <a href="https://example.com" class="bare">https://example.com</a>?</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with quoted text", func() {
			source := "https://example.com[_a_ *b* `c`]"
			expected := `<div class="paragraph">
<p><a href="https://example.com"><em>a</em> <strong>b</strong> <code>c</code></a></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with symbol in path and text", func() {
			source := "a link to https://example.com?foo=fighters&lang=en[a&b]."
			expected := `<div class="paragraph">
<p>a link to <a href="https://example.com?foo=fighters&amp;lang=en">a&amp;b</a>.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with unquoted text having comma", func() {
			source := "https://example.com[A, B, and C]"
			// here, `B` and `and C` are considered as other positional args,
			// not as part of the link text.
			expected := `<div class="paragraph">
<p><a href="https://example.com">A</a></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		// 		It("email link with unquoted text having comma", func() {
		// 			source := "mailto:foo@example.com[A, B, and C]"
		// 			expected := `<div class="paragraph">
		// <p><a href="mailto:foo@example.com?subject=B&amp;body=and+C">A</a></p>
		// </div>`
		// 			Expect(RenderHTML(source)).To(MatchHTML(expected))
		// 		})

		It("with quoted text having comma", func() {
			source := `mailto:foo@example.com["A, B, and C"]`
			expected := `<div class="paragraph">
<p><a href="mailto:foo@example.com">A, B, and C</a></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("inside a multiline paragraph", func() {
			source := `a https://website.com
and more text on the
next lines`
			expected := `<div class="paragraph">
<p>a <a href="https://website.com" class="bare">https://website.com</a>
and more text on the
next lines</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with text, target and role", func() {
			source := `a link to https://example.com[example,window=mytarget,role=myrole]`
			expected := `<div class="paragraph">
<p>a link to <a href="https://example.com" class="myrole" target="mytarget">example</a></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with target and role", func() {
			source := `a link to https://example.com[window=mytarget,role=myrole]`
			expected := `<div class="paragraph">
<p>a link to <a href="https://example.com" class="bare myrole" target="mytarget">https://example.com</a></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with text and custom target", func() {
			source := `a link to https://example.com[the doc,window=read-later]`
			expected := `<div class="paragraph">
<p>a link to <a href="https://example.com" target="read-later">the doc</a></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with text and custom target with noopener", func() {
			source := `a link to https://example.com[the doc,window=read-later,opts=noopener]`
			expected := `<div class="paragraph">
<p>a link to <a href="https://example.com" target="read-later" rel="noopener">the doc</a></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with text and explicit blank target", func() {
			source := `a link to https://example.com[the doc,window=_blank]`
			expected := `<div class="paragraph">
<p>a link to <a href="https://example.com" target="_blank" rel="noopener">the doc</a></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with text and blank target short-hand", func() {
			source := `a link to https://example.com[the doc^]` // the ^ character is used to define a `blank` target
			expected := `<div class="paragraph">
<p>a link to <a href="https://example.com" target="_blank" rel="noopener">the doc</a></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		Context("with document attribute substitutions", func() {

			It("with a document attribute substitution for the whole URL", func() {
				source := `:url: https://foo.bar
	
:url: https://foo2.bar
	
a link to {url}`
				expected := `<div class="paragraph">
<p>a link to <a href="https://foo2.bar" class="bare">https://foo2.bar</a></p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("with two document attribute substitutions only", func() {
				source := `:scheme: https
:path: foo.bar
	
a link to {scheme}://{path}`
				expected := `<div class="paragraph">
<p>a link to <a href="https://foo.bar" class="bare">https://foo.bar</a></p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("with two document attribute substitutions and a reset", func() {
				source := `:scheme: https
:path: foo.bar

:!path:
	
a link to {scheme}://{path}`
				expected := `<div class="paragraph">
<p>a link to <a href="https://{path}" class="bare">https://{path}</a></p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("with document attribute in section 0 title", func() {
				source := `= a title to {scheme}://{path} and https://foo.baz
:scheme: https
:path: foo.bar`
				expected := `a title to https://foo.bar and https://foo.baz`
				_, metadata, err := RenderHTMLWithMetadata(source)
				Expect(err).NotTo(HaveOccurred())
				Expect(metadata.Title).To(Equal(expected))
			})

			It("with document attribute in section 1 title", func() {
				source := `:scheme: https
:path: foo.bar
	
== a title to {scheme}://{path} and https://foo.baz`
				expected := `<div class="sect1">
<h2 id="_a_title_to_httpsfoo_bar_and_httpsfoo_baz">a title to <a href="https://foo.bar" class="bare">https://foo.bar</a> and <a href="https://foo.baz" class="bare">https://foo.baz</a></h2>
<div class="sectionbody">
</div>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("with two document attribute substitutions and a reset", func() {
				source := `:scheme: https
:path: foo.bar

:!path:

a link to {scheme}://{path} and https://foo.baz`
				expected := `<div class="paragraph">
<p>a link to <a href="https://{path}" class="bare">https://{path}</a> and <a href="https://foo.baz" class="bare">https://foo.baz</a></p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("with special characters", func() {
				source := `a link to https://example.com?a=1&b=2`
				expected := `<div class="paragraph">
<p>a link to <a href="https://example.com?a=1&amp;b=2" class="bare">https://example.com?a=1&amp;b=2</a></p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})
		})
	})

	Context("relative links", func() {

		It("relative link to doc without text", func() {
			source := "a link to link:foo.adoc[]."
			expected := `<div class="paragraph">
<p>a link to <a href="foo.adoc" class="bare">foo.adoc</a>.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("relative link to doc with text", func() {
			source := "a link to link:foo.adoc[foo doc]."
			expected := `<div class="paragraph">
<p>a link to <a href="foo.adoc">foo doc</a>.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("relative link with text having comma", func() {
			source := "a link to link:foo.adoc[A, B, and C]"
			expected := `<div class="paragraph">
<p>a link to <a href="foo.adoc">A</a></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("relative link with quoted text having comma", func() {
			// must wrap link text in quotes to retain it all,
			// otherwise, it's cut after the first comma
			// TODO: expect `target=b` and `role= 'and C'` attributes
			source := "a link to link:foo.adoc['A, B, and C']"
			expected := `<div class="paragraph">
<p>a link to <a href="foo.adoc">A, B, and C</a></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("relative link to external URL with text", func() {
			source := "a link to link:https://foo.bar[foo doc]."
			expected := `<div class="paragraph">
<p>a link to <a href="https://foo.bar">foo doc</a>.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("invalid relative link to doc", func() {
			source := "a link to link:foo.adoc."
			expected := `<div class="paragraph">
<p>a link to link:foo.adoc.</p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("relative link with quoted text", func() {
			source := "link:/[_a_ *b* `c`]"
			expected := `<div class="paragraph">
<p><a href="/"><em>a</em> <strong>b</strong> <code>c</code></a></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with text, target and role", func() {
			source := `a link to link:https://example.com[example,window=mytarget,role=myrole]`
			expected := `<div class="paragraph">
<p>a link to <a href="https://example.com" class="myrole" target="mytarget">example</a></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with target and role", func() {
			source := `a link to link:https://example.com[window=mytarget,role=myrole]`
			expected := `<div class="paragraph">
<p>a link to <a href="https://example.com" class="bare myrole" target="mytarget">https://example.com</a></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with text and custom target", func() {
			source := `a link to link:https://example.com[the doc,window=read-later]`
			expected := `<div class="paragraph">
<p>a link to <a href="https://example.com" target="read-later">the doc</a></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with text and custom target with noopener", func() {
			source := `a link to link:https://example.com[the doc,window=read-later,opts=noopener]`
			expected := `<div class="paragraph">
<p>a link to <a href="https://example.com" target="read-later" rel="noopener">the doc</a></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with text and explicit blank target", func() {
			source := `a link to link:https://example.com[the doc,window=_blank]`
			expected := `<div class="paragraph">
<p>a link to <a href="https://example.com" target="_blank" rel="noopener">the doc</a></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with text and blank target short-hand", func() {
			source := `a link to link:https://example.com[the doc^]` // the ^ character is used to define a `blank` target
			expected := `<div class="paragraph">
<p>a link to <a href="https://example.com" target="_blank" rel="noopener">the doc</a></p>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		Context("with document attribute substitutions", func() {

			It("with attribute in section 0 title", func() {
				source := `= a title to {scheme}://{path} and https://example.com
:scheme: https
:path: example.com`
				expectedTmpl := `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<meta name="generator" content="libasciidoc">
<title>a title to https://example.com and https://example.com</title>
</head>
<body class="article">
<div id="header">
<h1>a title to <a href="https://example.com" class="bare">https://example.com</a> and <a href="https://example.com" class="bare">https://example.com</a></h1>
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
				Expect(RenderHTML(source,
					configuration.WithHeaderFooter(true),
					configuration.WithLastUpdated(now),
				)).To(MatchHTMLTemplate(expectedTmpl,
					struct {
						LastUpdated string
					}{
						LastUpdated: now.Format(configuration.LastUpdatedFormat),
					}))
			})

			It("with attribute in section 1 title", func() {
				source := `
:scheme: https
:path: example.com

== a title to {scheme}://{path} and https://example.com
`
				expected := `<div class="sect1">
<h2 id="_a_title_to_httpsexample_com_and_httpsexample_com">a title to <a href="https://example.com" class="bare">https://example.com</a> and <a href="https://example.com" class="bare">https://example.com</a></h2>
<div class="sectionbody">
</div>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("relative link with two document attribute substitutions and a reset", func() {
				source := `:scheme: link
:path: foo.bar

:!path:

a link to {scheme}:{path}[] and https://foo.baz`
				expected := `<div class="paragraph">
<p>a link to <a href="{path}" class="bare">{path}</a> and <a href="https://foo.baz" class="bare">https://foo.baz</a></p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("links with underscores", func() {
				source := "link:a_[A] link:a_[A]"
				expected := `<div class="paragraph">
<p><a href="a_">A</a> <a href="a_">A</a></p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("links with line breaks in attributes", func() {
				source := `link:x[
title]`
				expected := `<div class="paragraph">
<p><a href="x">title</a></p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})
		})
	})
})
