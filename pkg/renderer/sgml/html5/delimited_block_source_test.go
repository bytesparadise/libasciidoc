package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" // nolint:golint
	. "github.com/onsi/gomega" // nolint:golintt
)

var _ = Describe("source blocks", func() {

	Context("as Asciidoc blocks", func() {

		It("with source attribute only", func() {
			source := `[source]
----
require 'sinatra'

get '/hi' do
  "Hello World!"
end
----`
			expected := `<div class="listingblock">
<div class="content">
<pre class="highlight"><code>require 'sinatra'

get '/hi' do
  "Hello World!"
end</code></pre>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with title, source and languages attributes", func() {
			source := `[source,ruby]
.Source block title
----
require 'sinatra'

get '/hi' do
  "Hello World!"
end

----`
			expected := `<div class="listingblock">
<div class="title">Source block title</div>
<div class="content">
<pre class="highlight"><code class="language-ruby" data-lang="ruby">require 'sinatra'

get '/hi' do
  "Hello World!"
end</code></pre>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with title, source and languages attributes and empty trailing line", func() {
			source := `[source,ruby]
.Source block title
----
require 'sinatra'

get '/hi' do
  "Hello World!"
end

----`
			expected := `<div class="listingblock">
<div class="title">Source block title</div>
<div class="content">
<pre class="highlight"><code class="language-ruby" data-lang="ruby">require 'sinatra'

get '/hi' do
  "Hello World!"
end</code></pre>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with title, source and unknown languages attributes", func() {
			source := `[source,brainfart]
.Source block title
----
int main(int argc, char **argv);
----`
			expected := `<div class="listingblock">
<div class="title">Source block title</div>
<div class="content">
<pre class="highlight"><code class="language-brainfart" data-lang="brainfart">int main(int argc, char **argv);</code></pre>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with id, title, source and languages attributes", func() {
			source := `[#id-for-source-block]
[source,ruby]
.app.rb
----
require 'sinatra'

get '/hi' do
  "Hello World!"
end
----`
			expected := `<div id="id-for-source-block" class="listingblock">
<div class="title">app.rb</div>
<div class="content">
<pre class="highlight"><code class="language-ruby" data-lang="ruby">require 'sinatra'

get '/hi' do
  "Hello World!"
end</code></pre>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with html content", func() {
			source := `[source]
----
<a>link</a>
----`
			expected := `<div class="listingblock">
<div class="content">
<pre class="highlight"><code>&lt;a&gt;link&lt;/a&gt;</code></pre>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with highlighter and callouts", func() {
			source := `:source-highlighter: chroma
[source, c]
----
#include <stdio.h>

printf("Hello world!\n"); // <1>
<a>link</a>
----
<1> A greeting
`
			expected := `<div class="listingblock">
<div class="content">
<pre class="chroma highlight"><code data-lang="c"><span class="tok-cp">#include</span> <span class="tok-cpf">&lt;stdio.h&gt;</span>

<span class="tok-n">printf</span><span class="tok-p">(</span><span class="tok-s">&#34;Hello world!</span><span class="tok-se">\n</span><span class="tok-s">&#34;</span><span class="tok-p">);</span> <span class="tok-o">//</span> <b class="conum">(1)</b>
<span class="tok-o">&lt;</span><span class="tok-n">a</span><span class="tok-o">&gt;</span><span class="tok-n">link</span><span class="tok-o">&lt;/</span><span class="tok-n">a</span><span class="tok-o">&gt;</span></code></pre>
</div>
</div>
<div class="colist arabic">
<ol>
<li>
<p>A greeting</p>
</li>
</ol>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with other content", func() {
			source := `----
  a<<b
----`
			expected := `<div class="listingblock">
<div class="content">
<pre>  a&lt;&lt;b</pre>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with callout and admonition block afterwards", func() {
			source := `[source]
----
const cookies = "cookies" <1>
----
<1> a constant

[NOTE]
====
a note
====`

			expected := `<div class="listingblock">
<div class="content">
<pre class="highlight"><code>const cookies = "cookies" <b class="conum">(1)</b></code></pre>
</div>
</div>
<div class="colist arabic">
<ol>
<li>
<p>a constant</p>
</li>
</ol>
</div>
<div class="admonitionblock note">
<table>
<tr>
<td class="icon">
<div class="title">Note</div>
</td>
<td class="content">
<div class="paragraph">
<p>a note</p>
</div>
</td>
</tr>
</table>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with nowrap option", func() {
			source := `[source%nowrap,go]
----
const Cookie = "cookie"
----`
			expected := `<div class="listingblock">
<div class="content">
<pre class="highlight nowrap"><code class="language-go" data-lang="go">const Cookie = "cookie"</code></pre>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		Context("with syntax highlighting", func() {

			It("with callouts and syntax highlighting", func() {
				source := `[source,java]
----
@QuarkusTest
public class GreetingResourceTest {

    @InjectMock
    @RestClient // <1>
    GreetingService greetingService;

    @Test
    public void testHelloEndpoint() {
        Mockito.when(greetingService.hello()).thenReturn("hello from mockito");

        given()
          .when().get("/hello")
          .then()
             .statusCode(200)
             .body(is("hello from mockito"));
    }

}
----
<1> We need to use the @RestClient CDI qualifier, since Quarkus creates the GreetingService bean with this qualifier.
`
				expected := `<div class="listingblock">
<div class="content">
<pre class="highlight"><code class="language-java" data-lang="java">@QuarkusTest
public class GreetingResourceTest {

    @InjectMock
    @RestClient // <b class="conum">(1)</b>
    GreetingService greetingService;

    @Test
    public void testHelloEndpoint() {
        Mockito.when(greetingService.hello()).thenReturn("hello from mockito");

        given()
          .when().get("/hello")
          .then()
             .statusCode(200)
             .body(is("hello from mockito"));
    }

}</code></pre>
</div>
</div>
<div class="colist arabic">
<ol>
<li>
<p>We need to use the @RestClient CDI qualifier, since Quarkus creates the GreetingService bean with this qualifier.</p>
</li>
</ol>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("should render source block with go syntax only", func() {
				source := `:source-highlighter: pygments
        
[source,go]
----
type Foo struct{
    Field string
}
----`
				expected := `<div class="listingblock">
<div class="content">
<pre class="pygments highlight"><code data-lang="go"><span class="tok-kd">type</span> <span class="tok-nx">Foo</span> <span class="tok-kd">struct</span><span class="tok-p">{</span>
    <span class="tok-nx">Field</span> <span class="tok-kt">string</span>
<span class="tok-p">}</span></code></pre>
</div>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("should render source block without highlighter when language is not set", func() {
				source := `:source-highlighter: pygments
        
[source]
----
type Foo struct{
    Field string
}
----`
				expected := `<div class="listingblock">
<div class="content">
<pre class="pygments highlight"><code>type Foo struct{
    Field string
}</code></pre>
</div>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("should render source block without highlighter when language is not set", func() {
				source := `:source-highlighter: pygments
        
[source]
----
type Foo struct{
    Field string
}
----`
				expected := `<div class="listingblock">
<div class="content">
<pre class="pygments highlight"><code>type Foo struct{
    Field string
}</code></pre>
</div>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("should render source block with go syntax and custom style", func() {
				source := `:source-highlighter: pygments
:pygments-style: manni

[source,go]
----
type Foo struct{
    Field string
}
----`
				expected := `<div class="listingblock">
<div class="content">
<pre class="pygments highlight"><code data-lang="go"><span class="tok-kd">type</span> <span class="tok-nx">Foo</span> <span class="tok-kd">struct</span><span class="tok-p">{</span>
    <span class="tok-nx">Field</span> <span class="tok-kt">string</span>
<span class="tok-p">}</span></code></pre>
</div>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("should render source block with go syntax, custom style and line numbers", func() {
				source := `:source-highlighter: pygments
:pygments-style: manni
:pygments-linenums-mode: inline

[source,go,linenums]
----
type Foo struct{
    Field string
}
----`
				expected := `<div class="listingblock">
<div class="content">
<pre class="pygments highlight"><code data-lang="go"><span class="tok-ln">1</span><span class="tok-kd">type</span> <span class="tok-nx">Foo</span> <span class="tok-kd">struct</span><span class="tok-p">{</span>
<span class="tok-ln">2</span>    <span class="tok-nx">Field</span> <span class="tok-kt">string</span>
<span class="tok-ln">3</span><span class="tok-p">}</span></code></pre>
</div>
</div>
` // the pygment.py sets the line number class to `tok-ln` but here we expect `tok-ln`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("should render source block with go syntax, custom style, inline css and line numbers", func() {
				source := `:source-highlighter: pygments
:pygments-style: manni
:pygments-css: style
:pygments-linenums-mode: inline

[source,go,linenums]
----
type Foo struct{
    Field string
}
----`
				expected := `<div class="listingblock">
<div class="content">
<pre class="pygments highlight"><code data-lang="go"><span style="margin-right:0.4em;padding:0 0.4em 0 0.4em;color:#7f7f7f">1</span><span style="color:#069;font-weight:bold">type</span> Foo <span style="color:#069;font-weight:bold">struct</span>{
<span style="margin-right:0.4em;padding:0 0.4em 0 0.4em;color:#7f7f7f">2</span>    Field <span style="color:#078;font-weight:bold">string</span>
<span style="margin-right:0.4em;padding:0 0.4em 0 0.4em;color:#7f7f7f">3</span>}</code></pre>
</div>
</div>
` // the pygment.py sets the line number class to `tok-ln` but here we expect `tok-ln`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("with nowrap option", func() {
				source := `:source-highlighter: pygments
:pygments-style: manni
:pygments-css: style
:pygments-linenums-mode: inline

[source%nowrap,go]
----
const Cookie = "cookie"
----`
				expected := `<div class="listingblock">
<div class="content">
<pre class="pygments highlight nowrap"><code data-lang="go"><span style="color:#069;font-weight:bold">const</span> Cookie = <span style="color:#c30">&#34;cookie&#34;</span></code></pre>
</div>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("with non-ascii content and syntax highlighting", func() {
				// source content is `…`, not `...`
				source := `:source-highlighter: chroma
:unicode: false

[source,c]
----
…
----`
				expected := `<div class="listingblock">
<div class="content">
<pre class="chroma highlight"><code data-lang="c"><span class="tok-err">…</span></code></pre>
</div>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("with symbol content and syntax highlighting", func() {
				// source content is `…`, not `...`
				source := `:source-highlighter: chroma
:unicode: false

[source,c]
----
Copyright (C)
----`
				expected := `<div class="listingblock">
<div class="content">
<pre class="chroma highlight"><code data-lang="c"><span class="tok-n">Copyright</span> <span class="tok-p">(</span><span class="tok-n">C</span><span class="tok-p">)</span></code></pre>
</div>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})
		})
	})

	Context("as Markdown block", func() {

		It("with HTML content", func() {
			source := ".title\n" +
				"```html\n" +
				"<!DOCTYPE html>\n" +
				"</html>\n" +
				"```"
			expected := `<div class="listingblock">
<div class="title">title</div>
<div class="content">
<pre class="highlight"><code class="language-html" data-lang="html">&lt;!DOCTYPE html&gt;
&lt;/html&gt;</code></pre>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with non-ascii content and syntax highlighting", func() {
			// source content is `…`, not `...`, but the `unicode:false` attribute should NOT affect the content of the listing/source block
			source :=
				":source-highlighter: chroma\n" +
					":unicode: false\n" +
					"\n" +
					"```c\n" +
					"…\n" +
					"```\n"
			expected := `<div class="listingblock">
<div class="content">
<pre class="chroma highlight"><code data-lang="c"><span class="tok-err">…</span></code></pre>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with symbol content and syntax highlighting", func() {
			// source content is `…`, not `...`, but the `unicode:false` attribute should NOT affect the content of the listing/source block
			source :=
				":source-highlighter: chroma\n" +
					":unicode: false\n" +
					"\n" +
					"```c\n" +
					"Copyright (C)\n" +
					"```\n"
			expected := `<div class="listingblock">
<div class="content">
<pre class="chroma highlight"><code data-lang="c"><span class="tok-n">Copyright</span> <span class="tok-p">(</span><span class="tok-n">C</span><span class="tok-p">)</span></code></pre>
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})
})
