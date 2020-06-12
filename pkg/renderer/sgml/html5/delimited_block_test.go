package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("delimited blocks", func() {

	Context("fenced blocks", func() {

		It("fenced block with multiple lines", func() {
			source := "```\n\nsome source code \n\nhere  \n\n\n\n```"
			expected := `<div class="listingblock">
<div class="content">
<pre class="highlight"><code>some source code

here</code></pre>
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("fenced block with id and title", func() {
			source := "[#id-for-fences]\n.fenced block title\n```\nsome source code\n\nhere\n\n\n\n```"
			expected := `<div id="id-for-fences" class="listingblock">
<div class="title">fenced block title</div>
<div class="content">
<pre class="highlight"><code>some source code

here</code></pre>
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("fenced block with external link inside", func() {
			source := "```" + "\n" +
				"a http://website.com" + "\n" +
				"and more text on the" + "\n" +
				"next lines" + "\n\n" +
				"```"
			expected := `<div class="listingblock">
<div class="content">
<pre class="highlight"><code>a http://website.com
and more text on the
next lines</code></pre>
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("listing blocks", func() {

		It("with multiple lines", func() {
			source := `----
some source code

here

----`
			expected := `<div class="listingblock">
<div class="content">
<pre>some source code

here</pre>
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with ID and title", func() {
			source := `[#id-for-listing-block]
.listing block title
----
some source code
----`
			expected := `<div id="id-for-listing-block" class="listingblock">
<div class="title">listing block title</div>
<div class="content">
<pre>some source code</pre>
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with ID and title and empty trailing line", func() {
			source := `[#id-for-listing-block]
.listing block title
----
some source code

----`
			expected := `<div id="id-for-listing-block" class="listingblock">
<div class="title">listing block title</div>
<div class="content">
<pre>some source code</pre>
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with html content", func() {
			source := `----
<a>link</a>
----`
			expected := `<div class="listingblock">
<div class="content">
<pre>&lt;a&gt;link&lt;/a&gt;</pre>
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with single callout", func() {
			source := `----
import <1>
----
<1> an import`
			expected := `<div class="listingblock">
<div class="content">
<pre>import <b class="conum">(1)</b></pre>
</div>
</div>
<div class="colist arabic">
<ol>
<li>
<p>an import</p>
</li>
</ol>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with multiple callouts and blankline between calloutitems", func() {
			source := `----
import <1>

func foo() {} <2>
----
<1> an import

<2> a func`
			expected := `<div class="listingblock">
<div class="content">
<pre>import <b class="conum">(1)</b>

func foo() {} <b class="conum">(2)</b></pre>
</div>
</div>
<div class="colist arabic">
<ol>
<li>
<p>an import</p>
</li>
<li>
<p>a func</p>
</li>
</ol>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with multiple callouts on same line", func() {
			source := `----
import <1> <2><3>

func foo() {} <4>
----
<1> an import
<2> a single import
<3> a single basic import
<4> a func`
			expected := `<div class="listingblock">
<div class="content">
<pre>import <b class="conum">(1)</b><b class="conum">(2)</b><b class="conum">(3)</b>

func foo() {} <b class="conum">(4)</b></pre>
</div>
</div>
<div class="colist arabic">
<ol>
<li>
<p>an import</p>
</li>
<li>
<p>a single import</p>
</li>
<li>
<p>a single basic import</p>
</li>
<li>
<p>a func</p>
</li>
</ol>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with invalid callout", func() {
			source := `----
import <a>
----
<a> an import`
			expected := `<div class="listingblock">
<div class="content">
<pre>import &lt;a&gt;</pre>
</div>
</div>
<div class="paragraph">
<p>&lt;a&gt; an import</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("source blocks", func() {

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
</div>`
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
</div>`
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
</div>`
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
</div>`
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
</div>`
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
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
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
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("example blocks", func() {

		It("example block with multiple elements - case 1", func() {
			source := `====
some listing code
with *bold content*

* and a list item

====`
			expected := `<div class="exampleblock">
<div class="content">
<div class="paragraph">
<p>some listing code
with <strong>bold content</strong></p>
</div>
<div class="ulist">
<ul>
<li>
<p>and a list item</p>
</li>
</ul>
</div>
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("example block with multiple elements - case 2", func() {
			source := `====
*bold content*

and more content
====`
			expected := `<div class="exampleblock">
<div class="content">
<div class="paragraph">
<p><strong>bold content</strong></p>
</div>
<div class="paragraph">
<p>and more content</p>
</div>
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("example block with multiple elements - case 3", func() {
			source := `====
*bold content*

and "more" content
====`
			expected := `<div class="exampleblock">
<div class="content">
<div class="paragraph">
<p><strong>bold content</strong></p>
</div>
<div class="paragraph">
<p>and "more" content</p>
</div>
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("example block with ID and title", func() {
			source := `[#id-for-example-block]
.example block title
====
foo

====`
			expected := `<div id="id-for-example-block" class="exampleblock">
<div class="title">Example 1. example block title</div>
<div class="content">
<div class="paragraph">
<p>foo</p>
</div>
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("admonition blocks", func() {

		It("admonition block with multiple elements alone", func() {
			source := `[NOTE]
====
some listing code
with *bold content*

* and a list item

====`
			expected := `<div class="admonitionblock note">
<table>
<tr>
<td class="icon">
<div class="title">Note</div>
</td>
<td class="content">
<div class="paragraph">
<p>some listing code
with <strong>bold content</strong></p>
</div>
<div class="ulist">
<ul>
<li>
<p>and a list item</p>
</li>
</ul>
</div>
</td>
</tr>
</table>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("admonition block with ID and title", func() {
			source := `[NOTE]
[#id-for-admonition-block]
.title for admonition block
====
some listing code
with *bold content*

* and a list item
====`
			expected := `<div id="id-for-admonition-block" class="admonitionblock note">
<table>
<tr>
<td class="icon">
<div class="title">Note</div>
</td>
<td class="content">
<div class="title">title for admonition block</div>
<div class="paragraph">
<p>some listing code
with <strong>bold content</strong></p>
</div>
<div class="ulist">
<ul>
<li>
<p>and a list item</p>
</li>
</ul>
</div>
</td>
</tr>
</table>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
		It("admonition block with ID, title and icon", func() {
			source := `:icons: font
			
[NOTE]
[#id-for-admonition-block]
.title for admonition block
====
some listing code
with *bold content*

* and a list item

====`
			expected := `<div id="id-for-admonition-block" class="admonitionblock note">
<table>
<tr>
<td class="icon">
<i class="fa icon-note" title="Note"></i>
</td>
<td class="content">
<div class="title">title for admonition block</div>
<div class="paragraph">
<p>some listing code
with <strong>bold content</strong></p>
</div>
<div class="ulist">
<ul>
<li>
<p>and a list item</p>
</li>
</ul>
</div>
</td>
</tr>
</table>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("admonition paragraph and admonition block with multiple elements", func() {
			source := `[CAUTION]                      
this is an admonition paragraph.
								
								
[NOTE]                         
.Title2                        
====                           
This is an admonition block
								
with another paragraph    
====      `
			expected := `<div class="admonitionblock caution">
<table>
<tr>
<td class="icon">
<div class="title">Caution</div>
</td>
<td class="content">
this is an admonition paragraph.
</td>
</tr>
</table>
</div>
<div class="admonitionblock note">
<table>
<tr>
<td class="icon">
<div class="title">Note</div>
</td>
<td class="content">
<div class="title">Title2</div>
<div class="paragraph">
<p>This is an admonition block</p>
</div>
<div class="paragraph">
<p>with another paragraph</p>
</div>
</td>
</tr>
</table>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("admonition paragraph with an icon", func() {
			source := `:icons: font

TIP: an admonition text on
2 lines.`
			expected := `<div class="admonitionblock tip">
<table>
<tr>
<td class="icon">
<i class="fa icon-tip" title="Tip"></i>
</td>
<td class="content">
an admonition text on
2 lines.
</td>
</tr>
</table>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("admonition paragraph with ID, title and icon", func() {
			source := `:icons: font

[#id-for-admonition-block]
.title for the admonition block
TIP: an admonition text on 1 line.
`
			expected := `<div id="id-for-admonition-block" class="admonitionblock tip">
<table>
<tr>
<td class="icon">
<i class="fa icon-tip" title="Tip"></i>
</td>
<td class="content">
<div class="title">title for the admonition block</div>
an admonition text on 1 line.
</td>
</tr>
</table>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("quote blocks", func() {

		It("single-line quote with author and title ", func() {
			source := `[quote, john doe, quote title]
____
some *quote* content

____`
			expected := `<div class="quoteblock">
<blockquote>
<div class="paragraph">
<p>some <strong>quote</strong> content</p>
</div>
</blockquote>
<div class="attribution">
&#8212; john doe<br>
<cite>quote title</cite>
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("single-line quote with author and title, and ID and title ", func() {
			source := `[#id-for-quote-block]
[quote, john doe, quote title]
.title for quote block
____
some *quote* content
____`
			expected := `<div id="id-for-quote-block" class="quoteblock">
<div class="title">title for quote block</div>
<blockquote>
<div class="paragraph">
<p>some <strong>quote</strong> content</p>
</div>
</blockquote>
<div class="attribution">
&#8212; john doe<br>
<cite>quote title</cite>
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("multi-line quote with author and title", func() {
			source := `[quote, john doe, quote title]
____

- some 
- quote 
- content

____`
			expected := `<div class="quoteblock">
<blockquote>
<div class="ulist">
<ul>
<li>
<p>some</p>
</li>
<li>
<p>quote</p>
</li>
<li>
<p>content</p>
</li>
</ul>
</div>
</blockquote>
<div class="attribution">
&#8212; john doe<br>
<cite>quote title</cite>
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("multi-line quote with author only and nested listing", func() {
			source := `[quote, john doe]
____
* some
----
* quote 
----
* content
____`
			expected := `<div class="quoteblock">
<blockquote>
<div class="ulist">
<ul>
<li>
<p>some</p>
</li>
</ul>
</div>
<div class="listingblock">
<div class="content">
<pre>* quote</pre>
</div>
</div>
<div class="ulist">
<ul>
<li>
<p>content</p>
</li>
</ul>
</div>
</blockquote>
<div class="attribution">
&#8212; john doe
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("single-line quote with title only", func() {
			source := `[quote, , quote title]
____
some quote content
____`
			expected := `<div class="quoteblock">
<blockquote>
<div class="paragraph">
<p>some quote content</p>
</div>
</blockquote>
<div class="attribution">
&#8212; quote title
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("multi-line quote without author and title", func() {
			source := `[quote]
____
lines 
	and tabs 
are preserved, but not trailing spaces   

____`

			expected := `<div class="quoteblock">
<blockquote>
<div class="paragraph">
<p>lines
	and tabs
are preserved, but not trailing spaces</p>
</div>
</blockquote>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("empty quote without author and title", func() {
			source := `[quote]
____
____`
			// asciidoctor will include an empty line in the `blockquote` element, I'm not sure why.
			expected := `<div class="quoteblock">
<blockquote>

</blockquote>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))

		})
	})

	Context("markdown-style quote blocks", func() {

		It("with single marker without author", func() {
			source := `> some text
on *multiple lines*`

			expected := `<div class="quoteblock">
<blockquote>
<div class="paragraph">
<p>some text
on <strong>multiple lines</strong></p>
</div>
</blockquote>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with marker on each line without author", func() {
			source := `> some text
> on *multiple lines*`

			expected := `<div class="quoteblock">
<blockquote>
<div class="paragraph">
<p>some text
on <strong>multiple lines</strong></p>
</div>
</blockquote>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with marker on each line with author", func() {
			source := `> some text
> on *multiple lines*
> -- John Doe`
			expected := `<div class="quoteblock">
<blockquote>
<div class="paragraph">
<p>some text
on <strong>multiple lines</strong></p>
</div>
</blockquote>
<div class="attribution">
&#8212; John Doe
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with marker on each line with author and title", func() {
			source := `.title
> some text
> on *multiple lines*
> -- John Doe`
			expected := `<div class="quoteblock">
<div class="title">title</div>
<blockquote>
<div class="paragraph">
<p>some text
on <strong>multiple lines</strong></p>
</div>
</blockquote>
<div class="attribution">
&#8212; John Doe
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with with author only", func() {
			source := `> -- John Doe`
			expected := `<div class="quoteblock">
<blockquote>

</blockquote>
<div class="attribution">
&#8212; John Doe
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("verse blocks", func() {

		It("single-line verse with author and title ", func() {
			source := `[verse, john doe, verse title]
____
some *verse* content

____`
			expected := `<div class="verseblock">
<pre class="content">some <strong>verse</strong> content</pre>
<div class="attribution">
&#8212; john doe<br>
<cite>verse title</cite>
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("single-line verse with author, id and title ", func() {
			source := `[verse, john doe, verse title]
[#id-for-verse-block]
.title for verse block
____
some *verse* content
____`
			expected := `<div id="id-for-verse-block" class="verseblock">
<div class="title">title for verse block</div>
<pre class="content">some <strong>verse</strong> content</pre>
<div class="attribution">
&#8212; john doe<br>
<cite>verse title</cite>
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("multi-line verse with author and title", func() {
			source := `[verse, john doe, verse title]
____
- some 
- verse 
- content

and more!

____`
			expected := `<div class="verseblock">
<pre class="content">- some
- verse
- content

and more!</pre>
<div class="attribution">
&#8212; john doe<br>
<cite>verse title</cite>
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("single-line verse with author only", func() {
			source := `[verse, john doe]
____
some verse content
____`
			expected := `<div class="verseblock">
<pre class="content">some verse content</pre>
<div class="attribution">
&#8212; john doe
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("single-line verse with title only", func() {
			source := `[verse, , verse title]
____
some verse content
____`
			expected := `<div class="verseblock">
<pre class="content">some verse content</pre>
<div class="attribution">
&#8212; verse title
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("multi-line verse without author and title", func() {
			source := `[verse]
____
lines 
	and tabs 
are preserved

____`

			expected := `<div class="verseblock">
<pre class="content">lines
	and tabs
are preserved</pre>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("empty verse without author and title", func() {
			source := `[verse]
____
____`
			expected := `<div class="verseblock">
<pre class="content"></pre>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))

		})
	})

	Context("sidebar blocks", func() {

		It("sidebar block with paragraph", func() {
			source := `****
some *verse* content

****`
			expected := `<div class="sidebarblock">
<div class="content">
<div class="paragraph">
<p>some <strong>verse</strong> content</p>
</div>
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("sidebar block with id, title, paragraph and sourcecode block", func() {
			source := `[#id-for-sidebar]
.title for sidebar
****
some *verse* content

----
foo
bar
----
****`
			expected := `<div id="id-for-sidebar" class="sidebarblock">
<div class="content">
<div class="title">title for sidebar</div>
<div class="paragraph">
<p>some <strong>verse</strong> content</p>
</div>
<div class="listingblock">
<div class="content">
<pre>foo
bar</pre>
</div>
</div>
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("syntax highlighting with pygments", func() {

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
</div>`
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
</div>`
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
</div>`
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
</div>`
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
</div>` // the pygment.py sets the line number class to `tok-ln` but here we expect `tok-ln`
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
</div>` // the pygment.py sets the line number class to `tok-ln` but here we expect `tok-ln`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	Context("passthrough blocks", func() {

		It("with title", func() {
			source := `.a title
++++
_foo_

*bar*
++++`
			expected := `_foo_

*bar*`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

	})

	Context("passthrough open block", func() {

		It("2-line paragraph followed by another paragraph", func() {
			source := `[pass]
_foo_
*bar*

another paragraph`
			expected := `_foo_
*bar*
<div class="paragraph">
<p>another paragraph</p>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})
})
