package sgml

import (
	"strconv"
	"strings"
	"text/template"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderTable(ctx *renderer.Context, t *types.Table) (string, error) {
	result := &strings.Builder{}
	caption := &strings.Builder{}

	number := 0
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
		c, found, err := t.Attributes.GetAsString(types.AttrCaption)
		if err != nil {
			return "", err
		} else if !found {
			c, found, err = ctx.Attributes.GetAsString(types.AttrTableCaption)
			if err != nil {
				return "", errors.Wrap(err, "unable to render table")
			}
			if found && c != "" {
				// We always append the figure number, unless the caption is disabled.
				// This is for asciidoctor compatibility.
				c += " {counter:table-number}. "
			}
		}
		// TODO: This is a very primitive and incomplete replacement of the counter attribute only.
		// This should be removed when attribute values are allowed to contain attributes.
		// Also this expansion should be limited to just singly quoted strings in the Attribute list,
		// or the default.  Ultimately this should all be done long before it gets into the renderer.
		if strings.Contains(c, "{counter:table-number}") {
			number = ctx.GetAndIncrementTableCounter()
			c = strings.ReplaceAll(c, "{counter:table-number}", strconv.Itoa(number))
		}
		caption.WriteString(c)
	}
	columns, err := t.Columns()
	if err != nil {
		return "", errors.Wrap(err, "failed to render table")
	}
	header, err := r.renderTableHeader(ctx, t.Header, columns)
	if err != nil {
		return "", errors.Wrap(err, "failed to render table")
	}
	footer, err := r.renderTableFooter(ctx, t.Footer, columns)
	if err != nil {
		return "", errors.Wrap(err, "failed to render table")
	}
	body, err := r.renderTableBody(ctx, t.Rows, columns)
	if err != nil {
		return "", errors.Wrap(err, "failed to render table")
	}
	roles, err := r.renderElementRoles(ctx, t.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render table roles")
	}
	title, err := r.renderElementTitle(t.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render table title")
	}
	err = r.table.Execute(result, struct {
		Context     *renderer.Context
		ID          string
		Title       string
		Columns     []*types.TableColumn
		TableNumber int
		Caption     string
		Frame       string
		Grid        string
		Fit         string
		Float       string
		Stripes     string
		Width       int
		Roles       string
		Header      string
		Body        string
		Footer      string
	}{
		Context:     ctx,
		ID:          r.renderElementID(t.Attributes),
		Title:       title,
		Columns:     columns,
		TableNumber: number,
		Caption:     caption.String(),
		Roles:       roles,
		Frame:       frame,
		Grid:        grid,
		Fit:         fit,
		Float:       float,
		Stripes:     stripes,
		Width:       width,
		Header:      header,
		Body:        body,
		Footer:      footer,
	})
	if err != nil {
		return "", errors.Wrap(err, "failed to render table")
	}
	return result.String(), nil
}

func (r *sgmlRenderer) renderTableHeader(ctx *renderer.Context, h *types.TableRow, cols []*types.TableColumn) (string, error) {
	if h == nil {
		return "", nil
	}
	result := &strings.Builder{}
	content := &strings.Builder{}
	col := 0
	for _, cell := range h.Cells {
		c, err := r.renderTableCell(ctx, r.tableHeaderCell, cell, cols[col%len(cols)])
		col++
		if err != nil {
			return "", errors.Wrap(err, "unable to render header")
		}
		content.WriteString(c)
	}
	err := r.tableHeader.Execute(result, struct {
		Context *renderer.Context
		Content string
		Cells   []*types.TableCell
	}{
		Context: ctx,
		Content: content.String(),
		Cells:   h.Cells,
	})
	return result.String(), err
}

func (r *sgmlRenderer) renderTableFooter(ctx *renderer.Context, f *types.TableRow, cols []*types.TableColumn) (string, error) {
	if f == nil {
		return "", nil
	}
	result := &strings.Builder{}
	content := &strings.Builder{}
	col := 0
	for _, cell := range f.Cells {
		c, err := r.renderTableCell(ctx, r.tableFooterCell, cell, cols[col%len(cols)])
		col++
		if err != nil {
			return "", errors.Wrap(err, "unable to render header")
		}
		content.WriteString(c)
	}
	err := r.tableFooter.Execute(result, struct {
		Context *renderer.Context
		Content string
		Cells   []*types.TableCell
	}{
		Context: ctx,
		Content: content.String(),
		Cells:   f.Cells,
	})
	return result.String(), err
}

func (r *sgmlRenderer) renderTableBody(ctx *renderer.Context, rows []*types.TableRow, columns []*types.TableColumn) (string, error) {
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debug("rendering table body")
	// 	log.Debugf("columns:\n%s", spew.Sdump(columns))
	// 	log.Debugf("rows:\n%s", spew.Sdump(rows))
	// }
	result := &strings.Builder{}
	content := &strings.Builder{}
	for _, row := range rows {
		c, err := r.renderTableRow(ctx, row, columns)
		if err != nil {
			return "", errors.Wrap(err, "unable to render body")
		}
		content.WriteString(c)
	}
	err := r.tableBody.Execute(result, struct {
		Context *renderer.Context
		Content string
		Rows    []*types.TableRow
		Columns []*types.TableColumn
	}{
		Context: ctx,
		Content: content.String(),
		Rows:    rows,
		Columns: columns,
	})
	return result.String(), err
}

func (r *sgmlRenderer) renderTableRow(ctx *renderer.Context, l *types.TableRow, cols []*types.TableColumn) (string, error) {
	result := &strings.Builder{}
	content := &strings.Builder{}
	for i, cell := range l.Cells {
		c, err := r.renderTableCell(ctx, r.tableCell, cell, cols[i])
		if err != nil {
			return "", errors.Wrap(err, "unable to render row")
		}
		content.WriteString(c)
	}
	err := r.tableRow.Execute(result, struct {
		Context *renderer.Context
		Content string
		Cells   []*types.TableCell
	}{
		Context: ctx,
		Content: content.String(),
		Cells:   l.Cells,
	})
	return result.String(), err
}

func (r *sgmlRenderer) renderTableCell(ctx *renderer.Context, tmpl *template.Template, cell *types.TableCell, col *types.TableColumn) (string, error) {
	result := &strings.Builder{}
	content, err := r.renderInlineElements(ctx, cell.Elements)
	if err != nil {
		return "", errors.Wrap(err, "unable to render cell")
	}
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("rendering cell with content '%s' and def %s", content, spew.Sdump(col))
	// }
	err = tmpl.Execute(result, struct {
		Context *renderer.Context
		Content string
		Cell    *types.TableCell
		HAlign  types.HAlign
		VAlign  types.VAlign
	}{
		Context: ctx,
		Content: content,
		Cell:    cell,
		HAlign:  col.HAlign,
		VAlign:  col.VAlign,
	})
	return result.String(), err
}
