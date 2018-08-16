package html5

import (
	"bytes"
	"fmt"
	"math"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var tableTmpl texttemplate.Template

func init() {
	tableTmpl = newTextTemplate("table", `{{ $ctx := .Context }}{{ with .Data }}<table class="tableblock frame-all grid-all stretch">{{ if .Lines }}
{{ if .Title }}<caption class="title">{{ .Title }}</caption>
{{ end }}<colgroup>
{{ $cellWidths := .CellWidths }}{{ range $index, $width := $cellWidths }}<col style="width: {{ $width }}%;">{{ includeNewline $ctx $index $cellWidths }}{{ end }}
</colgroup>
{{ if .Header.Cells }}<thead>
<tr>
{{ $headerCells := .Header.Cells }}{{ range $index, $cell := $headerCells }}<th class="tableblock halign-left valign-top">{{ renderElement $ctx $cell }}</th>{{ includeNewline $ctx $index $headerCells }}{{ end }}
</tr>
</thead>
{{ end }}<tbody>
{{ range $indexLine, $line := .Lines }}<tr>
{{ range $indexCells, $cell := $line.Cells }}<td class="tableblock halign-left valign-top"><p class="tableblock">{{ renderElement $ctx $cell }}</p></td>{{ includeNewline $ctx $indexCells $line.Cells }}{{ end }}
</tr>
{{ end }}</tbody>{{ end }}
</table>{{ end }}`,
		texttemplate.FuncMap{
			"renderElement":  renderElementAsString,
			"includeNewline": includeNewline,
		})
}

func renderTable(ctx *renderer.Context, t types.Table) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	// inspect first line to obtain cell width ratio
	widths := []string{}
	if len(t.Lines) > 0 {
		line := t.Lines[0]
		n := len(line.Cells)
		widths = make([]string, n)
		total := float64(0.0)
		for i := 0; i < n-1; i++ {
			w := float64(100.0) / float64(n)
			widths[i] = formatColumnWidth(w)
			total += w
		}
		// last width
		// int values don't need 4 decimals precision

		widths[n-1] = formatColumnWidth(100-total, lastColumn()) // make sure the last width as the upper rounded value
		log.Debugf("current total width: %v -> %v", total, widths[n-1])
	}
	var title string
	if titleAttr, ok := t.Attributes[types.AttrTitle].(string); ok {
		c := ctx.GetAndIncrementTableCounter()
		title = fmt.Sprintf("Table %d. %s", c, titleAttr)
	}
	err := tableTmpl.Execute(result, ContextualPipeline{
		Context: ctx,
		Data: struct {
			Title      string
			CellWidths []string
			Header     types.TableLine
			Lines      []types.TableLine
		}{
			Title:      title,
			CellWidths: widths,
			Header:     t.Header,
			Lines:      t.Lines,
		},
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to render table")
	}
	return result.Bytes(), nil
}

type formatColumnWidthOption func(float64) float64

func lastColumn() formatColumnWidthOption {
	return func(v float64) float64 {
		return v + 0.00005
	}
}

func formatColumnWidth(v float64, options ...formatColumnWidthOption) string {
	if v == math.Trunc(v) {
		// whole numbers don't need 4 decimals
		return fmt.Sprintf("%d", int(v))
	}
	for _, opt := range options {
		v = opt(v)
	}
	return fmt.Sprintf("%.4f", v)
}
