package sgml

import (
	"bytes"
	"fmt"
	"math"
	"strconv"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (r *sgmlRenderer) renderTable(ctx *renderer.Context, t types.Table) ([]byte, error) {
	result := &bytes.Buffer{}
	// inspect first line to obtain cell width ratio
	widths := []string{}
	if len(t.Lines) > 0 {
		line := t.Lines[0]
		n := len(line.Cells)
		widths = make([]string, n)
		total := 0.0
		for i := 0; i < n-1; i++ {
			w := 100.0 / float64(n)
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
		title = fmt.Sprintf("Table %d. %s", ctx.GetAndIncrementTableCounter(), EscapeString(titleAttr))
	}
	err := r.table.Execute(result, ContextualPipeline{
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
		return strconv.Itoa(int(v))
	}
	for _, opt := range options {
		v = opt(v)
	}
	return fmt.Sprintf("%.4f", v)
}
