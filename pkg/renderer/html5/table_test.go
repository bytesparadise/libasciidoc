package html5_test

import . "github.com/onsi/ginkgo"

var _ = Describe("tables", func() {

	It("1-line table with 2 cells", func() {
		actualContent := `|===
| *foo* foo  | _bar_  
|===`
		expectedResult := `<table class="tableblock frame-all grid-all spread">
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
		verify(GinkgoT(), expectedResult, actualContent)
	})

	It("1-line table with 3 cells", func() {
		actualContent := `|===
| *foo* foo  | _bar_  | baz
|===`
		expectedResult := `<table class="tableblock frame-all grid-all spread">
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
		verify(GinkgoT(), expectedResult, actualContent)
	})

	It("table with title, headers and 1 line per cell", func() {
		actualContent := `.table title
|===
|Column heading 1 |Column heading 2

|Column 1, row 1
|Column 2, row 1

|Column 1, row 2
|Column 2, row 2
|===`
		expectedResult := `<table class="tableblock frame-all grid-all spread">
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
		verify(GinkgoT(), expectedResult, actualContent)
	})

	It("empty table ", func() {
		actualContent := `|===
|===`
		expectedResult := `<table class="tableblock frame-all grid-all spread">
</table>`
		verify(GinkgoT(), expectedResult, actualContent)
	})
})
