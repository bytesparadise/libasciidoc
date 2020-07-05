package sgml

import (
	"strconv"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderTable(ctx *renderer.Context, t types.Table) (string, error) {
	result := &strings.Builder{}
	caption := &strings.Builder{}

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

	header, err := r.renderTableHeader(ctx, t.Header, t.Columns)
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
		Columns     []types.TableColumn
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
		Columns:     t.Columns,
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

func (r *sgmlRenderer) renderTableHeader(ctx *renderer.Context, l types.TableLine, cols []types.TableColumn) (string, error) {
	result := &strings.Builder{}
	content := &strings.Builder{}
	col := 0
	for _, cell := range l.Cells {
		c, err := r.renderTableHeaderCell(ctx, cell, cols[col%len(cols)])
		col++
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

func (r *sgmlRenderer) renderTableHeaderCell(ctx *renderer.Context, cell []interface{}, col types.TableColumn) (string, error) {
	result := &strings.Builder{}
	content, err := r.renderInlineElements(ctx, cell)
	if err != nil {
		return "", errors.Wrap(err, "unable to render header cell")
	}
	err = r.tableHeaderCell.Execute(result, struct {
		Context *renderer.Context
		Content string
		Cell    []interface{}
		VAlign  string
		HAlign  string
	}{
		Context: ctx,
		Content: content,
		Cell:    cell,
		HAlign:  col.HAlign,
		VAlign:  col.VAlign,
	})
	return result.String(), err
}

func (r *sgmlRenderer) renderTableBody(ctx *renderer.Context, t types.Table) (string, error) {
	result := &strings.Builder{}
	content := &strings.Builder{}
	for _, row := range t.Lines {
		c, err := r.renderTableRow(ctx, row, t.Columns)
		if err != nil {
			return "", errors.Wrap(err, "unable to render header")
		}
		content.WriteString(c)
	}
	err := r.tableBody.Execute(result, struct {
		Context *renderer.Context
		Content string
		Rows    []types.TableLine
		Columns []types.TableColumn
	}{
		Context: ctx,
		Content: content.String(),
		Rows:    t.Lines,
		Columns: t.Columns,
	})
	return result.String(), err
}

func (r *sgmlRenderer) renderTableRow(ctx *renderer.Context, l types.TableLine, cols []types.TableColumn) (string, error) {
	result := &strings.Builder{}
	content := &strings.Builder{}
	col := 0
	for _, cell := range l.Cells {
		c, err := r.renderTableCell(ctx, cell, cols[col%len(cols)])
		col++
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

func (r *sgmlRenderer) renderTableCell(ctx *renderer.Context, cell []interface{}, col types.TableColumn) (string, error) {
	result := &strings.Builder{}
	content, err := r.renderInlineElements(ctx, cell)
	if err != nil {
		return "", errors.Wrap(err, "unable to render header cell")
	}
	err = r.tableCell.Execute(result, struct {
		Context *renderer.Context
		Content string
		Cell    []interface{}
		HAlign  string
		VAlign  string
	}{
		Context: ctx,
		Content: content,
		Cell:    cell,
		HAlign:  col.HAlign,
		VAlign:  col.VAlign,
	})
	return result.String(), err
}
