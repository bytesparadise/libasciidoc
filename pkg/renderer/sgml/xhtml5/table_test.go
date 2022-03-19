package xhtml5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2" // nolint:golint
	. "github.com/onsi/gomega"    // nolint:golintt
)

var _ = Describe("tables", func() {

	It("1-line table with 2 cells", func() {
		source := `|===
| *foo* foo  | _bar_  
|===`
		expected := `<table class="tableblock frame-all grid-all stretch">
<colgroup>
<col style="width: 50%;"/>
<col style="width: 50%;"/>
</colgroup>
<tbody>
<tr>
<td class="tableblock halign-left valign-top"><p class="tableblock"><strong>foo</strong> foo</p></td>
<td class="tableblock halign-left valign-top"><p class="tableblock"><em>bar</em></p></td>
</tr>
</tbody>
</table>
`
		Expect(RenderXHTML(source)).To(MatchHTML(expected))
	})

	It("1-line table with 3 cells", func() {
		source := `|===
| *foo* foo  | _bar_  | baz
|===`
		expected := `<table class="tableblock frame-all grid-all stretch">
<colgroup>
<col style="width: 33.3333%;"/>
<col style="width: 33.3333%;"/>
<col style="width: 33.3334%;"/>
</colgroup>
<tbody>
<tr>
<td class="tableblock halign-left valign-top"><p class="tableblock"><strong>foo</strong> foo</p></td>
<td class="tableblock halign-left valign-top"><p class="tableblock"><em>bar</em></p></td>
<td class="tableblock halign-left valign-top"><p class="tableblock">baz</p></td>
</tr>
</tbody>
</table>
`
		Expect(RenderXHTML(source)).To(MatchHTML(expected))
	})

	It("table with title, headers and 1 line per cell", func() {
		source := `.table title
|===
|Column header 1 |Column header 2

|Column 1, row 1
|Column 2, row 1

|Column 1, row 2
|Column 2, row 2
|===`
		expected := `<table class="tableblock frame-all grid-all stretch">
<caption class="title">Table 1. table title</caption>
<colgroup>
<col style="width: 50%;"/>
<col style="width: 50%;"/>
</colgroup>
<thead>
<tr>
<th class="tableblock halign-left valign-top">Column header 1</th>
<th class="tableblock halign-left valign-top">Column header 2</th>
</tr>
</thead>
<tbody>
<tr>
<td class="tableblock halign-left valign-top"><p class="tableblock">Column 1, row 1</p></td>
<td class="tableblock halign-left valign-top"><p class="tableblock">Column 2, row 1</p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-top"><p class="tableblock">Column 1, row 2</p></td>
<td class="tableblock halign-left valign-top"><p class="tableblock">Column 2, row 2</p></td>
</tr>
</tbody>
</table>
`
		Expect(RenderXHTML(source)).To(MatchHTML(expected))
	})

	It("table with title, custom caption", func() {
		source := `.table title
[caption="Example I. "]
|===
|Column header 1 |Column header 2

|Column 1, row 1
|Column 2, row 1

|Column 1, row 2
|Column 2, row 2
|===`
		expected := `<table class="tableblock frame-all grid-all stretch">
<caption class="title">Example I. table title</caption>
<colgroup>
<col style="width: 50%;"/>
<col style="width: 50%;"/>
</colgroup>
<thead>
<tr>
<th class="tableblock halign-left valign-top">Column header 1</th>
<th class="tableblock halign-left valign-top">Column header 2</th>
</tr>
</thead>
<tbody>
<tr>
<td class="tableblock halign-left valign-top"><p class="tableblock">Column 1, row 1</p></td>
<td class="tableblock halign-left valign-top"><p class="tableblock">Column 2, row 1</p></td>
</tr>
<tr>
<td class="tableblock halign-left valign-top"><p class="tableblock">Column 1, row 2</p></td>
<td class="tableblock halign-left valign-top"><p class="tableblock">Column 2, row 2</p></td>
</tr>
</tbody>
</table>
`
		Expect(RenderXHTML(source)).To(MatchHTML(expected))
	})

	It("empty table ", func() {
		source := `|===
|===`
		expected := `<table class="tableblock frame-all grid-all stretch">
</table>
`
		Expect(RenderXHTML(source)).To(MatchHTML(expected))
	})

	It("2 tables with 1 counter", func() {
		source := `|===
| foo | bar
|===

.Title 2
|===
| foo | bar
|===`
		expected := `<table class="tableblock frame-all grid-all stretch">
<colgroup>
<col style="width: 50%;"/>
<col style="width: 50%;"/>
</colgroup>
<tbody>
<tr>
<td class="tableblock halign-left valign-top"><p class="tableblock">foo</p></td>
<td class="tableblock halign-left valign-top"><p class="tableblock">bar</p></td>
</tr>
</tbody>
</table>
<table class="tableblock frame-all grid-all stretch">
<caption class="title">Table 1. Title 2</caption>
<colgroup>
<col style="width: 50%;"/>
<col style="width: 50%;"/>
</colgroup>
<tbody>
<tr>
<td class="tableblock halign-left valign-top"><p class="tableblock">foo</p></td>
<td class="tableblock halign-left valign-top"><p class="tableblock">bar</p></td>
</tr>
</tbody>
</table>
`
		Expect(RenderXHTML(source)).To(MatchHTML(expected))
	})

	It("2 tables with no caption label", func() {
		source := `:table-caption!:

.Title 1
|===
| foo | bar
|===

.Title 2
|===
| foo | bar
|===`
		expected := `<table class="tableblock frame-all grid-all stretch">
<caption class="title">Title 1</caption>
<colgroup>
<col style="width: 50%;"/>
<col style="width: 50%;"/>
</colgroup>
<tbody>
<tr>
<td class="tableblock halign-left valign-top"><p class="tableblock">foo</p></td>
<td class="tableblock halign-left valign-top"><p class="tableblock">bar</p></td>
</tr>
</tbody>
</table>
<table class="tableblock frame-all grid-all stretch">
<caption class="title">Title 2</caption>
<colgroup>
<col style="width: 50%;"/>
<col style="width: 50%;"/>
</colgroup>
<tbody>
<tr>
<td class="tableblock halign-left valign-top"><p class="tableblock">foo</p></td>
<td class="tableblock halign-left valign-top"><p class="tableblock">bar</p></td>
</tr>
</tbody>
</table>
`
		Expect(RenderXHTML(source)).To(MatchHTML(expected))
	})

	It("2 tables with custom caption label", func() {
		source := `:table-caption: Chart

.First
|===
| foo | bar
|===

.Second
|===
| foo | bar
|===`
		expected := `<table class="tableblock frame-all grid-all stretch">
<caption class="title">Chart 1. First</caption>
<colgroup>
<col style="width: 50%;"/>
<col style="width: 50%;"/>
</colgroup>
<tbody>
<tr>
<td class="tableblock halign-left valign-top"><p class="tableblock">foo</p></td>
<td class="tableblock halign-left valign-top"><p class="tableblock">bar</p></td>
</tr>
</tbody>
</table>
<table class="tableblock frame-all grid-all stretch">
<caption class="title">Chart 2. Second</caption>
<colgroup>
<col style="width: 50%;"/>
<col style="width: 50%;"/>
</colgroup>
<tbody>
<tr>
<td class="tableblock halign-left valign-top"><p class="tableblock">foo</p></td>
<td class="tableblock halign-left valign-top"><p class="tableblock">bar</p></td>
</tr>
</tbody>
</table>
`
		Expect(RenderXHTML(source)).To(MatchHTML(expected))
	})

	It("2 tables with 2 counters", func() {
		source := `.Title 1
|===
| foo | bar
|===

.Title 2
|===
| foo | bar
|===`
		expected := `<table class="tableblock frame-all grid-all stretch">
<caption class="title">Table 1. Title 1</caption>
<colgroup>
<col style="width: 50%;"/>
<col style="width: 50%;"/>
</colgroup>
<tbody>
<tr>
<td class="tableblock halign-left valign-top"><p class="tableblock">foo</p></td>
<td class="tableblock halign-left valign-top"><p class="tableblock">bar</p></td>
</tr>
</tbody>
</table>
<table class="tableblock frame-all grid-all stretch">
<caption class="title">Table 2. Title 2</caption>
<colgroup>
<col style="width: 50%;"/>
<col style="width: 50%;"/>
</colgroup>
<tbody>
<tr>
<td class="tableblock halign-left valign-top"><p class="tableblock">foo</p></td>
<td class="tableblock halign-left valign-top"><p class="tableblock">bar</p></td>
</tr>
</tbody>
</table>
`
		Expect(RenderXHTML(source)).To(MatchHTML(expected))
	})

	It("autowidth ", func() {
		source := "[%autowidth]\n|===\n|==="
		expected := `<table class="tableblock frame-all grid-all fit-content">
</table>
`
		Expect(RenderXHTML(source)).To(MatchHTML(expected))
	})

	It("fixed width (number)", func() {
		source := "[width=75]\n|===\n|==="
		expected := `<table class="tableblock frame-all grid-all" style="width: 75%;">
</table>
`
		Expect(RenderXHTML(source)).To(MatchHTML(expected))
	})

	It("fixed width (percent)", func() {
		source := "[width=75%]\n|===\n|==="
		expected := `<table class="tableblock frame-all grid-all" style="width: 75%;">
</table>
`
		Expect(RenderXHTML(source)).To(MatchHTML(expected))
	})

	It("fixed width (100 percent)", func() {
		source := "[width=100%]\n|===\n|==="
		expected := `<table class="tableblock frame-all grid-all stretch">
</table>
`
		Expect(RenderXHTML(source)).To(MatchHTML(expected))
	})

	It("fixed width (> 100 percent)", func() {
		source := "[width=205]\n|===\n|==="
		expected := `<table class="tableblock frame-all grid-all stretch">
</table>
`
		Expect(RenderXHTML(source)).To(MatchHTML(expected))
	})

	It("fixed width overrides fit", func() {
		source := "[%autowidth,width=25]\n|===\n|==="
		expected := `<table class="tableblock frame-all grid-all" style="width: 25%;">
</table>
`
		Expect(RenderXHTML(source)).To(MatchHTML(expected))
	})

	It("fixed width overrides fit (> 100 percent)", func() {
		source := "[%autowidth,width=205]\n|===\n|==="
		expected := `<table class="tableblock frame-all grid-all stretch">
</table>
`
		Expect(RenderXHTML(source)).To(MatchHTML(expected))
	})

	It("grid, frames, float, stripes", func() {
		source := "[%autowidth,grid=rows,frame=sides,stripes=hover,float=right]\n|===\n|==="
		expected := `<table class="tableblock frame-sides grid-rows stripes-hover fit-content right">
</table>
`
		Expect(RenderXHTML(source)).To(MatchHTML(expected))
	})

	It("table with cols relative widths", func() {
		source := "[cols=\"3,2,5\"]\n|===\n|one|two|three\n|==="
		expected := `<table class="tableblock frame-all grid-all stretch">
<colgroup>
<col style="width: 30%;"/>
<col style="width: 20%;"/>
<col style="width: 50%;"/>
</colgroup>
<tbody>
<tr>
<td class="tableblock halign-left valign-top"><p class="tableblock">one</p></td>
<td class="tableblock halign-left valign-top"><p class="tableblock">two</p></td>
<td class="tableblock halign-left valign-top"><p class="tableblock">three</p></td>
</tr>
</tbody>
</table>
`
		Expect(RenderXHTML(source)).To(MatchHTML(expected))
	})

	It("table with cols relative widths and header", func() {
		source := "[cols=\"3,2,5\"]\n|===\n|h1|h2|h3\n\n|one|two|three\n|==="
		expected := `<table class="tableblock frame-all grid-all stretch">
<colgroup>
<col style="width: 30%;"/>
<col style="width: 20%;"/>
<col style="width: 50%;"/>
</colgroup>
<thead>
<tr>
<th class="tableblock halign-left valign-top">h1</th>
<th class="tableblock halign-left valign-top">h2</th>
<th class="tableblock halign-left valign-top">h3</th>
</tr>
</thead>
<tbody>
<tr>
<td class="tableblock halign-left valign-top"><p class="tableblock">one</p></td>
<td class="tableblock halign-left valign-top"><p class="tableblock">two</p></td>
<td class="tableblock halign-left valign-top"><p class="tableblock">three</p></td>
</tr>
</tbody>
</table>
`
		Expect(RenderXHTML(source)).To(MatchHTML(expected))
	})

	It("autowidth overrides column widths", func() {
		source := "[%autowidth,cols=\"3,2,5\"]\n|===\n|h1|h2|h3\n\n|one|two|three\n|==="
		expected := `<table class="tableblock frame-all grid-all fit-content">
<colgroup>
<col/>
<col/>
<col/>
</colgroup>
<thead>
<tr>
<th class="tableblock halign-left valign-top">h1</th>
<th class="tableblock halign-left valign-top">h2</th>
<th class="tableblock halign-left valign-top">h3</th>
</tr>
</thead>
<tbody>
<tr>
<td class="tableblock halign-left valign-top"><p class="tableblock">one</p></td>
<td class="tableblock halign-left valign-top"><p class="tableblock">two</p></td>
<td class="tableblock halign-left valign-top"><p class="tableblock">three</p></td>
</tr>
</tbody>
</table>
`
		Expect(RenderXHTML(source)).To(MatchHTML(expected))
	})

	It("column auto-width", func() {
		source := "[cols=\"30,~,~\"]\n|===\n|h1|h2|h3\n\n|one|two|three\n|==="
		expected := `<table class="tableblock frame-all grid-all stretch">
<colgroup>
<col style="width: 30%;"/>
<col/>
<col/>
</colgroup>
<thead>
<tr>
<th class="tableblock halign-left valign-top">h1</th>
<th class="tableblock halign-left valign-top">h2</th>
<th class="tableblock halign-left valign-top">h3</th>
</tr>
</thead>
<tbody>
<tr>
<td class="tableblock halign-left valign-top"><p class="tableblock">one</p></td>
<td class="tableblock halign-left valign-top"><p class="tableblock">two</p></td>
<td class="tableblock halign-left valign-top"><p class="tableblock">three</p></td>
</tr>
</tbody>
</table>
`
		Expect(RenderXHTML(source)).To(MatchHTML(expected))
	})

	It("columns with repeat", func() {
		source := "[cols=\"3*10,2*~\"]\n|===\n|h1|h2|h3|h4|h5\n\n|one|two|three|four|five\n|==="
		expected := `<table class="tableblock frame-all grid-all stretch">
<colgroup>
<col style="width: 10%;"/>
<col style="width: 10%;"/>
<col style="width: 10%;"/>
<col/>
<col/>
</colgroup>
<thead>
<tr>
<th class="tableblock halign-left valign-top">h1</th>
<th class="tableblock halign-left valign-top">h2</th>
<th class="tableblock halign-left valign-top">h3</th>
<th class="tableblock halign-left valign-top">h4</th>
<th class="tableblock halign-left valign-top">h5</th>
</tr>
</thead>
<tbody>
<tr>
<td class="tableblock halign-left valign-top"><p class="tableblock">one</p></td>
<td class="tableblock halign-left valign-top"><p class="tableblock">two</p></td>
<td class="tableblock halign-left valign-top"><p class="tableblock">three</p></td>
<td class="tableblock halign-left valign-top"><p class="tableblock">four</p></td>
<td class="tableblock halign-left valign-top"><p class="tableblock">five</p></td>
</tr>
</tbody>
</table>
`
		Expect(RenderXHTML(source)).To(MatchHTML(expected))
	})

	It("columns with alignment changes", func() {
		source := "[cols=\"2*^.^,<,.>,>\"]\n|===\n|h1|h2|h3|h4|h5\n\n|one|two|three|four|five\n|==="
		expected := `<table class="tableblock frame-all grid-all stretch">
<colgroup>
<col style="width: 20%;"/>
<col style="width: 20%;"/>
<col style="width: 20%;"/>
<col style="width: 20%;"/>
<col style="width: 20%;"/>
</colgroup>
<thead>
<tr>
<th class="tableblock halign-center valign-middle">h1</th>
<th class="tableblock halign-center valign-middle">h2</th>
<th class="tableblock halign-left valign-top">h3</th>
<th class="tableblock halign-left valign-bottom">h4</th>
<th class="tableblock halign-right valign-top">h5</th>
</tr>
</thead>
<tbody>
<tr>
<td class="tableblock halign-center valign-middle"><p class="tableblock">one</p></td>
<td class="tableblock halign-center valign-middle"><p class="tableblock">two</p></td>
<td class="tableblock halign-left valign-top"><p class="tableblock">three</p></td>
<td class="tableblock halign-left valign-bottom"><p class="tableblock">four</p></td>
<td class="tableblock halign-right valign-top"><p class="tableblock">five</p></td>
</tr>
</tbody>
</table>
`
		Expect(RenderXHTML(source)).To(MatchHTML(expected))
	})

	It("with header option", func() {
		source := `[cols="3*^",options="header"]
|===
|Dir (X,Y,Z) |Num Cells |Size
|X |10 |0.1
|Y |5  |0.2
|Z |10 |0.1
|===`
		expected := `<table class="tableblock frame-all grid-all stretch">
<colgroup>
<col style="width: 33.3333%;"/>
<col style="width: 33.3333%;"/>
<col style="width: 33.3334%;"/>
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
`
		Expect(RenderXHTML(source)).To(MatchHTML(expected))
	})

	It("with header and footer options", func() {
		source := `[%header%footer,cols="2,2,1"] 
|===
|Column 1, header row
|Column 2, header row
|Column 3, header row

|Cell in column 1, row 2
|Cell in column 2, row 2
|Cell in column 3, row 2

|Column 1, footer row
|Column 2, footer row
|Column 3, footer row
|===`
		expected := `<table class="tableblock frame-all grid-all stretch">
<colgroup>
<col style="width: 40%;"/>
<col style="width: 40%;"/>
<col style="width: 20%;"/>
</colgroup>
<thead>
<tr>
<th class="tableblock halign-left valign-top">Column 1, header row</th>
<th class="tableblock halign-left valign-top">Column 2, header row</th>
<th class="tableblock halign-left valign-top">Column 3, header row</th>
</tr>
</thead>
<tbody>
<tr>
<td class="tableblock halign-left valign-top"><p class="tableblock">Cell in column 1, row 2</p></td>
<td class="tableblock halign-left valign-top"><p class="tableblock">Cell in column 2, row 2</p></td>
<td class="tableblock halign-left valign-top"><p class="tableblock">Cell in column 3, row 2</p></td>
</tr>
</tbody>
<tfoot>
<tr>
<td class="tableblock halign-left valign-top"><p class="tableblock">Column 1, footer row</p></td>
<td class="tableblock halign-left valign-top"><p class="tableblock">Column 2, footer row</p></td>
<td class="tableblock halign-left valign-top"><p class="tableblock">Column 3, footer row</p></td>
</tr>
</tfoot>
</table>
`
		Expect(RenderXHTML(source)).To(MatchHTML(expected))
	})

	It("with id and title", func() {
		source := `[#non-uniform-mesh]
.Non-Uniform Mesh Parameters
[cols="3*^",options="header"]
|===
|Dir (X,Y,Z) |Num Cells |Size
|X |10 |0.1
|Y |10 |0.1
|Y |5  |0.2
|Z |10 |0.1
|===`

		expected := `<table id="non-uniform-mesh" class="tableblock frame-all grid-all stretch">
<caption class="title">Table 1. Non-Uniform Mesh Parameters</caption>
<colgroup>
<col style="width: 33.3333%;"/>
<col style="width: 33.3333%;"/>
<col style="width: 33.3334%;"/>
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
`
		Expect(RenderXHTML(source)).To(MatchHTML(expected))
	})
	// TODO: Verify styles -- it's verified in the parser for now, but we still need to implement styles.
})
