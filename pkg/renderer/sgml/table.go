package sgml

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (r *sgmlRenderer) renderTable(ctx *context, t *types.Table) (string, error) {
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

	caption := &strings.Builder{}
	if t.Attributes.Has(types.AttrTitle) {
		c, found := t.Attributes.GetAsString(types.AttrCaption)
		if !found {
			c, found = ctx.attributes.GetAsString(types.AttrTableCaption)
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
	title, err := r.renderElementTitle(ctx, t.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render table title")
	}
	return r.execute(r.table, struct {
		Context     *context
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
}

func (r *sgmlRenderer) renderTableHeader(ctx *context, h *types.TableRow, cols []*types.TableColumn) (string, error) {
	if h == nil {
		return "", nil
	}
	content := &strings.Builder{}
	col := 0
	for _, cell := range h.Cells {
		c, err := r.renderTableHeaderCell(ctx, cell, cols[col%len(cols)])
		col++
		if err != nil {
			return "", errors.Wrap(err, "unable to render header")
		}
		content.WriteString(c)
	}
	return r.execute(r.tableHeader, struct {
		Content string
	}{
		Content: content.String(),
	})
}

func (r *sgmlRenderer) renderTableHeaderCell(ctx *context, c *types.TableCell, col *types.TableColumn) (string, error) {
	// assume that elements to render are within the first element of the cell, which should be a paragraph
	if len(c.Elements) == 1 {
		if p, ok := c.Elements[0].(*types.Paragraph); ok {
			content, err := r.renderInlineElements(ctx, p.Elements)
			if err != nil {
				return "", errors.Wrap(err, "unable to render header cell")
			}
			return r.execute(r.tableHeaderCell, struct {
				HAlign  types.HAlign
				VAlign  types.VAlign
				Content string
			}{
				HAlign:  col.HAlign,
				VAlign:  col.VAlign,
				Content: content,
			})
		}
	}
	return "", fmt.Errorf("invalid header content (expected a single paragraph)")
}

func (r *sgmlRenderer) renderTableFooter(ctx *context, f *types.TableRow, cols []*types.TableColumn) (string, error) {
	if f == nil {
		return "", nil
	}
	content := &strings.Builder{}
	col := 0
	for _, cell := range f.Cells {
		c, err := r.renderTableFooterCell(ctx, cell, cols[col%len(cols)])
		col++
		if err != nil {
			return "", errors.Wrap(err, "unable to render footer")
		}
		content.WriteString(c)
	}
	return r.execute(r.tableFooter, struct {
		Context *context
		Content string
		Cells   []*types.TableCell
	}{
		Context: ctx,
		Content: content.String(),
		Cells:   f.Cells,
	})
}

func (r *sgmlRenderer) renderTableFooterCell(ctx *context, c *types.TableCell, col *types.TableColumn) (string, error) {
	// assume that elements to render are within the first element of the cell, which should be a paragraph
	if len(c.Elements) == 1 {
		if p, ok := c.Elements[0].(*types.Paragraph); ok {
			content, err := r.renderInlineElements(ctx, p.Elements)
			if err != nil {
				return "", errors.Wrap(err, "unable to render footer cell")
			}
			return r.execute(r.tableFooterCell, struct {
				HAlign  types.HAlign
				VAlign  types.VAlign
				Content string
			}{
				HAlign:  col.HAlign,
				VAlign:  col.VAlign,
				Content: content,
			})
		}
	}
	return "", fmt.Errorf("invalid footer content (expected a single paragraph)")
}

func (r *sgmlRenderer) renderTableBody(ctx *context, rows []*types.TableRow, columns []*types.TableColumn) (string, error) {
	content := &strings.Builder{}
	for _, row := range rows {
		c, err := r.renderTableRow(ctx, row, columns)
		if err != nil {
			return "", errors.Wrap(err, "unable to render body")
		}
		content.WriteString(c)
	}
	return r.execute(r.tableBody, struct {
		Context *context
		Content string
		Rows    []*types.TableRow
		Columns []*types.TableColumn
	}{
		Context: ctx,
		Content: content.String(),
		Rows:    rows,
		Columns: columns,
	})
}

func (r *sgmlRenderer) renderTableRow(ctx *context, l *types.TableRow, cols []*types.TableColumn) (string, error) {
	content := &strings.Builder{}
	for i, cell := range l.Cells {
		c, err := r.renderTableCell(ctx, cell, cols[i])
		if err != nil {
			return "", errors.Wrap(err, "unable to render row")
		}
		content.WriteString(c)
	}
	return r.execute(r.tableRow, struct {
		Context *context
		Content string
		Cells   []*types.TableCell
	}{
		Context: ctx,
		Content: content.String(),
		Cells:   l.Cells,
	})
}

func (r *sgmlRenderer) renderTableCell(ctx *context, cell *types.TableCell, col *types.TableColumn) (string, error) {
	buff := &strings.Builder{}
	for _, element := range cell.Elements {
		renderedElement, err := r.renderTableCellBlock(ctx, element)
		if err != nil {
			return "", err
		}
		buff.WriteString(renderedElement)
	}
	tmpl := r.tableCell
	if col.Style == types.HeaderStyle {
		tmpl = r.tableHeaderCell
	}
	return r.execute(tmpl, struct {
		Context *context
		Content string
		Cell    *types.TableCell
		HAlign  types.HAlign
		VAlign  types.VAlign
	}{
		Context: ctx,
		Content: buff.String(),
		Cell:    cell,
		HAlign:  col.HAlign,
		VAlign:  col.VAlign,
	})
}

func (r *sgmlRenderer) renderTableCellBlock(ctx *context, element interface{}) (string, error) {
	switch e := element.(type) {
	case *types.Paragraph:
		log.Debug("rendering paragraph within table cell")
		content, err := r.renderElements(ctx, e.Elements)
		if err != nil {
			return "", errors.Wrap(err, "unable to render table cell paragraph content")
		}
		title, err := r.renderElementTitle(ctx, e.Attributes)
		if err != nil {
			return "", errors.Wrap(err, "unable to render table cell paragraph content")
		}
		result, err := r.execute(r.embeddedParagraph, struct {
			Context    *context
			ID         string // TODO: not used in template?
			Title      string // TODO: not used in template?
			CheckStyle string
			Class      string
			Content    string
		}{
			Context:    ctx,
			ID:         r.renderElementID(e.Attributes),
			Title:      title,
			Class:      "tableblock",
			CheckStyle: renderCheckStyle(e.Attributes[types.AttrCheckStyle]),
			Content:    string(content),
		})
		if err != nil {
			return "", errors.Wrap(err, "unable to render table cell paragraph content")
		}
		return strings.TrimSuffix(result, "\n"), nil
	default:
		// Note: Asciidoctor wraps the `<div class=imageblock>` elements within a `<div class="content">`, which we also do here for the sake of compatibility
		renderedElement, err := r.renderElement(ctx, e)
		if err != nil {
			return "", errors.Wrap(err, "unable to render table cell")
		}
		return r.execute(r.tableCellBlock, struct {
			Content string
		}{
			Content: renderedElement,
		})
	}
}
