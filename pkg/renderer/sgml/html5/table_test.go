package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("tables", func() {

	It("1-line table with 2 cells", func() {
		source := `|===
| *foo* foo  | _bar_  
|===`
		expected := `<table class="tableblock frame-all grid-all stretch">
<colgroup>
<col style="width: 50%;">
<col style="width: 50%;">
</colgroup>
<tbody>
<tr>
<td class="tableblock halign-left valign-top"><p class="tableblock"><strong>foo</strong> foo</p></td>
<td class="tableblock halign-left valign-top"><p class="tableblock"><em>bar</em></p></td>
</tr>
</tbody>
</table>
`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("1-line table with 3 cells", func() {
		source := `|===
| *foo* foo  | _bar_  | baz
|===`
		expected := `<table class="tableblock frame-all grid-all stretch">
<colgroup>
<col style="width: 33.3333%;">
<col style="width: 33.3333%;">
<col style="width: 33.3334%;">
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
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("table with title, headers and 1 line per cell", func() {
		source := `.table title
|===
|Column heading 1 |Column heading 2

|Column 1, row 1
|Column 2, row 1

|Column 1, row 2
|Column 2, row 2
|===`
		expected := `<table class="tableblock frame-all grid-all stretch">
<caption class="title">Table 1. table title</caption>
<colgroup>
<col style="width: 50%;">
<col style="width: 50%;">
</colgroup>
<thead>
<tr>
<th class="tableblock halign-left valign-top">Column heading 1</th>
<th class="tableblock halign-left valign-top">Column heading 2</th>
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
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("table with title, custom caption", func() {
		source := `.table title
[caption="Example I. "]
|===
|Column heading 1 |Column heading 2

|Column 1, row 1
|Column 2, row 1

|Column 1, row 2
|Column 2, row 2
|===`
		expected := `<table class="tableblock frame-all grid-all stretch">
<caption class="title">Example I. table title</caption>
<colgroup>
<col style="width: 50%;">
<col style="width: 50%;">
</colgroup>
<thead>
<tr>
<th class="tableblock halign-left valign-top">Column heading 1</th>
<th class="tableblock halign-left valign-top">Column heading 2</th>
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
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("empty table ", func() {
		source := `|===
|===`
		expected := `<table class="tableblock frame-all grid-all stretch">
</table>
`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
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
<col style="width: 50%;">
<col style="width: 50%;">
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
<col style="width: 50%;">
<col style="width: 50%;">
</colgroup>
<tbody>
<tr>
<td class="tableblock halign-left valign-top"><p class="tableblock">foo</p></td>
<td class="tableblock halign-left valign-top"><p class="tableblock">bar</p></td>
</tr>
</tbody>
</table>
`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
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
<col style="width: 50%;">
<col style="width: 50%;">
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
<col style="width: 50%;">
<col style="width: 50%;">
</colgroup>
<tbody>
<tr>
<td class="tableblock halign-left valign-top"><p class="tableblock">foo</p></td>
<td class="tableblock halign-left valign-top"><p class="tableblock">bar</p></td>
</tr>
</tbody>
</table>
`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("autowidth ", func() {
		source := "[%autowidth]\n|===\n|==="
		expected := `<table class="tableblock frame-all grid-all fit-content">
</table>
`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("fixed width (number)", func() {
		source := "[width=75]\n|===\n|==="
		expected := `<table class="tableblock frame-all grid-all" style="width: 75%;">
</table>
`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("fixed width (percent)", func() {
		source := "[width=75%]\n|===\n|==="
		expected := `<table class="tableblock frame-all grid-all" style="width: 75%;">
</table>
`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("fixed width (100 percent)", func() {
		source := "[width=100%]\n|===\n|==="
		expected := `<table class="tableblock frame-all grid-all stretch">
</table>
`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("fixed width (> 100 percent)", func() {
		source := "[width=205]\n|===\n|==="
		expected := `<table class="tableblock frame-all grid-all stretch">
</table>
`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("fixed width overrides fit", func() {
		source := "[%autowidth,width=25]\n|===\n|==="
		expected := `<table class="tableblock frame-all grid-all" style="width: 25%;">
</table>
`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("fixed width overrides fit (> 100 percent)", func() {
		source := "[%autowidth,width=205]\n|===\n|==="
		expected := `<table class="tableblock frame-all grid-all stretch">
</table>
`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("grid, frames, float, stripes", func() {
		source := "[%autowidth,grid=rows,frame=sides,stripes=hover,float=right]\n|===\n|==="
		expected := `<table class="tableblock frame-sides grid-rows stripes-hover fit-content right">
</table>
`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("table with cols relative widths", func() {
		source := "[cols=\"3,2,5\"]\n|===\n|one|two|three\n|==="
		expected := `<table class="tableblock frame-all grid-all stretch">
<colgroup>
<col style="width: 30%;">
<col style="width: 20%;">
<col style="width: 50%;">
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
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("table with cols relative widths and header", func() {
		source := "[cols=\"3,2,5\"]\n|===\n|h1|h2|h3\n\n|one|two|three\n|==="
		expected := `<table class="tableblock frame-all grid-all stretch">
<colgroup>
<col style="width: 30%;">
<col style="width: 20%;">
<col style="width: 50%;">
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
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("autowidth overrides column widths", func() {
		source := "[%autowidth,cols=\"3,2,5\"]\n|===\n|h1|h2|h3\n\n|one|two|three\n|==="
		expected := `<table class="tableblock frame-all grid-all fit-content">
<colgroup>
<col>
<col>
<col>
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
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("column auto-width", func() {
		source := "[cols=\"30,~,~\"]\n|===\n|h1|h2|h3\n\n|one|two|three\n|==="
		expected := `<table class="tableblock frame-all grid-all stretch">
<colgroup>
<col style="width: 30%;">
<col>
<col>
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
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("columns with repeat", func() {
		source := "[cols=\"3*10,2*~\"]\n|===\n|h1|h2|h3|h4|h5\n\n|one|two|three|four|five\n|==="
		expected := `<table class="tableblock frame-all grid-all stretch">
<colgroup>
<col style="width: 10%;">
<col style="width: 10%;">
<col style="width: 10%;">
<col>
<col>
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
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("columns with alignment changes", func() {
		// source := "[cols=\"2*^.^,<,.>,>\"]\n|===\n|==="

		source := "[cols=\"2*^.^,<,.>,>\"]\n|===\n|h1|h2|h3|h4|h5\n\n|one|two|three|four|five\n|==="
		expected := `<table class="tableblock frame-all grid-all stretch">
<colgroup>
<col style="width: 20%;">
<col style="width: 20%;">
<col style="width: 20%;">
<col style="width: 20%;">
<col style="width: 20%;">
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
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	// TODO: Verify styles -- it's verified in the parser for now, but we still need to implement styles.
})
