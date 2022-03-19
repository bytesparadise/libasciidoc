package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2" // nolint:golint
	. "github.com/onsi/gomega"    // nolint:golintt
)

var _ = Describe("ordered lists", func() {

	It("with implicit numbering style on a single line", func() {
		source := `. item on a single line`
		expected := `<div class="olist arabic">
<ol class="arabic">
<li>
<p>item on a single line</p>
</li>
</ol>
</div>
`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("with implicit numbering style on multiple lines with leading tabs", func() {
		// leading tabs should be trimmed
		source := `. item 
	on 
	multiple 
	lines
`
		expected := `<div class="olist arabic">
<ol class="arabic">
<li>
<p>item
on
multiple
lines</p>
</li>
</ol>
</div>
`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("with title and role", func() {
		source := `.title
[#myid]
[.myrole]
. item 1`
		expected := `<div id="myid" class="olist arabic myrole">
<div class="title">title</div>
<ol class="arabic">
<li>
<p>item 1</p>
</li>
</ol>
</div>
`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("with explicit start only", func() {
		source := `[start=5]
. item`
		expected := `<div class="olist arabic">
<ol class="arabic" start="5">
<li>
<p>item</p>
</li>
</ol>
</div>
`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("with explicit quoted numbering and start", func() {
		source := `["lowerroman", start="5"]
. item`
		expected := `<div class="olist lowerroman">
<ol class="lowerroman" type="i" start="5">
<li>
<p>item</p>
</li>
</ol>
</div>
`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("item reversed with explicit quoted numbering and start", func() {
		source := `[lowerroman%reversed, start="5"]
. item 1
. item 2`
		expected := `<div class="olist lowerroman">
<ol class="lowerroman" type="i" start="5" reversed>
<li>
<p>item 1</p>
</li>
<li>
<p>item 2</p>
</li>
</ol>
</div>
`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	Context("with list element continuation", func() {

		It("case 1", func() {
			source := `. item 1
+
foo`
			expected := `<div class="olist arabic">
<ol class="arabic">
<li>
<p>item 1</p>
<div class="paragraph">
<p>foo</p>
</div>
</li>
</ol>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("case 2", func() {
			source := `. item 1
+
----
foo
----`
			expected := `<div class="olist arabic">
<ol class="arabic">
<li>
<p>item 1</p>
<div class="listingblock">
<div class="content">
<pre>foo</pre>
</div>
</div>
</li>
</ol>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("case 3", func() {
			source := `. cookie
+
image::cookie.png[]
+
. chocolate
+
image::chocolate.png[]`
			expected := `<div class="olist arabic">
<ol class="arabic">
<li>
<p>cookie</p>
<div class="imageblock">
<div class="content">
<img src="cookie.png" alt="cookie">
</div>
</div>
</li>
<li>
<p>chocolate</p>
<div class="imageblock">
<div class="content">
<img src="chocolate.png" alt="chocolate">
</div>
</div>
</li>
</ol>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("case 4", func() {
			source := `. In the table, enter the data shown in <<non-uniform-mesh>>
+
[#non-uniform-mesh]
.Non-Uniform Mesh Parameters
[cols="3*^",options="header"]
|===
|Dir (X,Y,Z) |Num Cells |Size
|X |10 |0.1
|Y |10 |0.1
|Y |5  |0.2
|Z |10 |0.1
|===
+
. Click *OK*`
			expected := `<div class="olist arabic">
<ol class="arabic">
<li>
<p>In the table, enter the data shown in <a href="#non-uniform-mesh">Non-Uniform Mesh Parameters</a></p>
<table id="non-uniform-mesh" class="tableblock frame-all grid-all stretch">
<caption class="title">Table 1. Non-Uniform Mesh Parameters</caption>
<colgroup>
<col style="width: 33.3333%;">
<col style="width: 33.3333%;">
<col style="width: 33.3334%;">
</colgroup>
<thead>
<tr>
<th class="tableblock halign-center valign-top">Dir (X,Y,Z)</th>
<th class="tableblock halign-center valign-top">Num Cells</th>
<th class="tableblock halign-center valign-top">Size</th>
</tr>
</thead>
<tbody>
<tr>
<td class="tableblock halign-center valign-top"><p class="tableblock">X</p></td>
<td class="tableblock halign-center valign-top"><p class="tableblock">10</p></td>
<td class="tableblock halign-center valign-top"><p class="tableblock">0.1</p></td>
</tr>
<tr>
<td class="tableblock halign-center valign-top"><p class="tableblock">Y</p></td>
<td class="tableblock halign-center valign-top"><p class="tableblock">10</p></td>
<td class="tableblock halign-center valign-top"><p class="tableblock">0.1</p></td>
</tr>
<tr>
<td class="tableblock halign-center valign-top"><p class="tableblock">Y</p></td>
<td class="tableblock halign-center valign-top"><p class="tableblock">5</p></td>
<td class="tableblock halign-center valign-top"><p class="tableblock">0.2</p></td>
</tr>
<tr>
<td class="tableblock halign-center valign-top"><p class="tableblock">Z</p></td>
<td class="tableblock halign-center valign-top"><p class="tableblock">10</p></td>
<td class="tableblock halign-center valign-top"><p class="tableblock">0.1</p></td>
</tr>
</tbody>
</table>
</li>
<li>
<p>Click <strong>OK</strong></p>
</li>
</ol>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})

	It("with unnumbered items", func() {
		source := `. item 1
		.. item 1.1
		... item 1.1.1
		... item 1.1.2
		.. item 1.2
		. item 2
		.. item 2.1`
		expected := `<div class="olist arabic">
<ol class="arabic">
<li>
<p>item 1</p>
<div class="olist loweralpha">
<ol class="loweralpha" type="a">
<li>
<p>item 1.1</p>
<div class="olist lowerroman">
<ol class="lowerroman" type="i">
<li>
<p>item 1.1.1</p>
</li>
<li>
<p>item 1.1.2</p>
</li>
</ol>
</div>
</li>
<li>
<p>item 1.2</p>
</li>
</ol>
</div>
</li>
<li>
<p>item 2</p>
<div class="olist loweralpha">
<ol class="loweralpha" type="a">
<li>
<p>item 2.1</p>
</li>
</ol>
</div>
</li>
</ol>
</div>
`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("mixed with unordered list - simple case", func() {
		source := `. Linux
* Fedora
* Ubuntu
* Slackware
. BSD
* FreeBSD
* NetBSD`
		expected := `<div class="olist arabic">
<ol class="arabic">
<li>
<p>Linux</p>
<div class="ulist">
<ul>
<li>
<p>Fedora</p>
</li>
<li>
<p>Ubuntu</p>
</li>
<li>
<p>Slackware</p>
</li>
</ul>
</div>
</li>
<li>
<p>BSD</p>
<div class="ulist">
<ul>
<li>
<p>FreeBSD</p>
</li>
<li>
<p>NetBSD</p>
</li>
</ul>
</div>
</li>
</ol>
</div>
`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("mixed with unordered list - complex case", func() {
		source := `- unordered 1
1. ordered 1.1
	a. ordered 1.1.a
	b. ordered 1.1.b
	c. ordered 1.1.c
2. ordered 1.2
	i)  ordered 1.2.i
	ii) ordered 1.2.ii
3. ordered 1.3
4. ordered 1.4
- unordered 2
* unordered 2.1
** unordered 2.1.1
with some
extra lines.
** unordered 2.1.2
* unordered 2.2
- unordered 3
. ordered 3.1
. ordered 3.2
[upperroman]
	.. ordered 3.2.I
	.. ordered 3.2.II
. ordered 3.3`
		expected := `<div class="ulist">
<ul>
<li>
<p>unordered 1</p>
<div class="olist arabic">
<ol class="arabic">
<li>
<p>ordered 1.1</p>
<div class="olist loweralpha">
<ol class="loweralpha" type="a">
<li>
<p>ordered 1.1.a</p>
</li>
<li>
<p>ordered 1.1.b</p>
</li>
<li>
<p>ordered 1.1.c</p>
</li>
</ol>
</div>
</li>
<li>
<p>ordered 1.2</p>
<div class="olist lowerroman">
<ol class="lowerroman" type="i">
<li>
<p>ordered 1.2.i</p>
</li>
<li>
<p>ordered 1.2.ii</p>
</li>
</ol>
</div>
</li>
<li>
<p>ordered 1.3</p>
</li>
<li>
<p>ordered 1.4</p>
</li>
</ol>
</div>
</li>
<li>
<p>unordered 2</p>
<div class="ulist">
<ul>
<li>
<p>unordered 2.1</p>
<div class="ulist">
<ul>
<li>
<p>unordered 2.1.1
with some
extra lines.</p>
</li>
<li>
<p>unordered 2.1.2</p>
</li>
</ul>
</div>
</li>
<li>
<p>unordered 2.2</p>
</li>
</ul>
</div>
</li>
<li>
<p>unordered 3</p>
<div class="olist arabic">
<ol class="arabic">
<li>
<p>ordered 3.1</p>
</li>
<li>
<p>ordered 3.2</p>
<div class="olist upperroman">
<ol class="upperroman" type="I">
<li>
<p>ordered 3.2.I</p>
</li>
<li>
<p>ordered 3.2.II</p>
</li>
</ol>
</div>
</li>
<li>
<p>ordered 3.3</p>
</li>
</ol>
</div>
</li>
</ul>
</div>
`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("all kinds of lists - complex case 3", func() {
		source := `* foo
1. bar
a. foo
2. baz
* foo2
- bar2`
		expected := `<div class="ulist">
<ul>
<li>
<p>foo</p>
<div class="olist arabic">
<ol class="arabic">
<li>
<p>bar</p>
<div class="olist loweralpha">
<ol class="loweralpha" type="a">
<li>
<p>foo</p>
</li>
</ol>
</div>
</li>
<li>
<p>baz</p>
</li>
</ol>
</div>
</li>
<li>
<p>foo2</p>
<div class="ulist">
<ul>
<li>
<p>bar2</p>
</li>
</ul>
</div>
</li>
</ul>
</div>
`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("drop principal text in list item", func() {
		source := `. {blank}
+
----
print("one")
----
. {blank}
+
----
print("one")
----`
		expected := `<div class="olist arabic">
<ol class="arabic">
<li>
<p></p>
<div class="listingblock">
<div class="content">
<pre>print("one")</pre>
</div>
</div>
</li>
<li>
<p></p>
<div class="listingblock">
<div class="content">
<pre>print("one")</pre>
</div>
</div>
</li>
</ol>
</div>
`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	Context("attach to ordered list item ancestor", func() {

		It("attach to grandparent ordered list item", func() {
			source := `. grandparent list item
.. parent list item
... child list item


+
paragraph attached to grandparent list item`
			expected := `<div class="olist arabic">
<ol class="arabic">
<li>
<p>grandparent list item</p>
<div class="olist loweralpha">
<ol class="loweralpha" type="a">
<li>
<p>parent list item</p>
<div class="olist lowerroman">
<ol class="lowerroman" type="i">
<li>
<p>child list item</p>
</li>
</ol>
</div>
</li>
</ol>
</div>
<div class="paragraph">
<p>paragraph attached to grandparent list item</p>
</div>
</li>
</ol>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("attach to parent ordered list item", func() {
			source := `. grandparent list item
.. parent list item
... child list item

+
paragraph attached to parent list item`
			expected := `<div class="olist arabic">
<ol class="arabic">
<li>
<p>grandparent list item</p>
<div class="olist loweralpha">
<ol class="loweralpha" type="a">
<li>
<p>parent list item</p>
<div class="olist lowerroman">
<ol class="lowerroman" type="i">
<li>
<p>child list item</p>
</li>
</ol>
</div>
<div class="paragraph">
<p>paragraph attached to parent list item</p>
</div>
</li>
</ol>
</div>
</li>
</ol>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("attach to child ordered list item", func() {
			source := `. grandparent list item
.. parent list item
... child list item
+
paragraph attached to child list item`
			expected := `<div class="olist arabic">
<ol class="arabic">
<li>
<p>grandparent list item</p>
<div class="olist loweralpha">
<ol class="loweralpha" type="a">
<li>
<p>parent list item</p>
<div class="olist lowerroman">
<ol class="lowerroman" type="i">
<li>
<p>child list item</p>
<div class="paragraph">
<p>paragraph attached to child list item</p>
</div>
</li>
</ol>
</div>
</li>
</ol>
</div>
</li>
</ol>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})
})
