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
</table>`
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
</table>`
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
</table>`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("empty table ", func() {
		source := `|===
|===`
		expected := `<table class="tableblock frame-all grid-all stretch">
</table>`
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
</table>`
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
</table>`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

})
