package sgml

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (r *sgmlRenderer) renderTable(ctx *renderer.Context, t types.Table) (string, error) {
	result := &strings.Builder{}
	caption := &strings.Builder{}

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
	number := 0
	title := r.renderElementTitle(t.Attributes)
	fit := "stretch"
	frame := t.Attributes.GetAsStringWithDefault(types.AttrFrame, "all")
	grid := t.Attributes.GetAsStringWithDefault(types.AttrGrid, "all")
	float := t.Attributes.GetAsStringWithDefault(types.AttrFloat, "")
	stripes := t.Attributes.GetAsStringWithDefault(types.AttrStripes, "")

	width, _ := strconv.Atoi(
		strings.TrimSuffix(t.Attributes.GetAsStringWithDefault(types.AttrWidth, ""), "%"))

	// These are derived from asciidoctor, and our rules here:
	// * Width can be a number or a percentage
	// * If width is >= 100, then it becomes "stretch" role, and we clear it
	// * If width is any other number (besides 0), we do not use the fitting role,
	//   and instead use an explicit style for the width.
	// * If width is unset, and %autowidth is set, then we use a fit-content role.
	// * If none of the above cases are true, we use stretch role (default)
	if t.Attributes.HasOption("autowidth") {
		fit = "fit-content"
	}
	if width >= 100 {
		width = 0
		fit = "stretch"
	} else if width > 0 {
		fit = ""
	}

	if t.Attributes.Has(types.AttrTitle) {
		number = ctx.GetAndIncrementTableCounter()
		if s, ok := t.Attributes.GetAsString(types.AttrCaption); ok {
			caption.WriteString(s)
		} else {
			err := r.tableCaption.Execute(caption, struct {
				TableNumber int
				Title       sanitized
			}{
				TableNumber: number,
				Title:       title,
			})
			if err != nil {
				return "", errors.Wrap(err, "unable to format table caption")
			}
		}
	}

	header, err := r.renderTableHeader(ctx, t.Header)
	if err != nil {
		return "", errors.Wrap(err, "failed to render table")
	}

	body, err := r.renderTableBody(ctx, t)
	if err != nil {
		return "", errors.Wrap(err, "failed to render table")
	}

	err = r.table.Execute(result, struct {
		Context     *renderer.Context
		Title       sanitized
		CellWidths  []string
		TableNumber int
		Caption     string
		Frame       string
		Grid        string
		Fit         string
		Float       string
		Stripes     string
		Width       int
		Roles       sanitized
		Header      string
		Body        string
	}{
		Context:     ctx,
		Title:       r.renderElementTitle(t.Attributes),
		CellWidths:  widths,
		TableNumber: number,
		Caption:     caption.String(),
		Roles:       r.renderElementRoles(t.Attributes),
		Frame:       frame,
		Grid:        grid,
		Fit:         fit,
		Float:       float,
		Stripes:     stripes,
		Width:       width,
		Header:      header,
		Body:        body,
	})
	if err != nil {
		return "", errors.Wrap(err, "failed to render table")
	}
	return result.String(), nil
}

func (r *sgmlRenderer) renderTableHeader(ctx *renderer.Context, l types.TableLine) (string, error) {
	result := &strings.Builder{}
	content := &strings.Builder{}
	for _, cell := range l.Cells {
		c, err := r.renderTableHeaderCell(ctx, cell)
		if err != nil {
			return "", errors.Wrap(err, "unable to render header")
		}
		content.WriteString(c)
	}
	err := r.tableHeader.Execute(result, struct {
		Context *renderer.Context
		Content string
		Cells   [][]interface{}
	}{
		Context: ctx,
		Content: content.String(),
		Cells:   l.Cells,
	})
	return result.String(), err
}

func (r *sgmlRenderer) renderTableHeaderCell(ctx *renderer.Context, cell []interface{}) (string, error) {
	result := &strings.Builder{}
	content, err := r.renderInlineElements(ctx, cell)
	if err != nil {
		return "", errors.Wrap(err, "unable to render header cell")
	}
	err = r.tableHeaderCell.Execute(result, struct {
		Context *renderer.Context
		Content string
		Cell    []interface{}
	}{
		Context: ctx,
		Content: content,
		Cell:    cell,
	})
	return result.String(), err
}

func (r *sgmlRenderer) renderTableBody(ctx *renderer.Context, t types.Table) (string, error) {
	result := &strings.Builder{}
	content := &strings.Builder{}
	for _, row := range t.Lines {
		c, err := r.renderTableRow(ctx, row)
		if err != nil {
			return "", errors.Wrap(err, "unable to render header")
		}
		content.WriteString(c)
	}
	err := r.tableBody.Execute(result, struct {
		Context *renderer.Context
		Content string
		Rows    []types.TableLine
	}{
		Context: ctx,
		Content: content.String(),
		Rows:    t.Lines,
	})
	return result.String(), err
}

func (r *sgmlRenderer) renderTableRow(ctx *renderer.Context, l types.TableLine) (string, error) {
	result := &strings.Builder{}
	content := &strings.Builder{}
	for _, cell := range l.Cells {
		c, err := r.renderTableCell(ctx, cell)
		if err != nil {
			return "", errors.Wrap(err, "unable to render header")
		}
		content.WriteString(c)
	}
	err := r.tableRow.Execute(result, struct {
		Context *renderer.Context
		Content string
		Cells   [][]interface{}
	}{
		Context: ctx,
		Content: content.String(),
		Cells:   l.Cells,
	})
	return result.String(), err
}

func (r *sgmlRenderer) renderTableCell(ctx *renderer.Context, cell []interface{}) (string, error) {
	result := &strings.Builder{}
	content, err := r.renderInlineElements(ctx, cell)
	if err != nil {
		return "", errors.Wrap(err, "unable to render header cell")
	}
	err = r.tableCell.Execute(result, struct {
		Context *renderer.Context
		Content string
		Cell    []interface{}
	}{
		Context: ctx,
		Content: content,
		Cell:    cell,
	})
	return result.String(), err
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
