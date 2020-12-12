package types

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// ------------------------------------------
// Tables
// ------------------------------------------

// TableColumn a table column
type TableColumn struct {
	widthVal float64 // internally used number, will be 0 for automatic, cleared post processing
	Width    string  // percentage or relative (0 for automatic)
	HAlign   string  // left, right, or center
	VAlign   string  // top, bottom, or middle
	Style    string  // Single character
}

var defaultColumn = TableColumn{
	widthVal: 1,
	HAlign:   "left",
	VAlign:   "top",
}

// Table the structure for the tables
type Table struct {
	Attributes Attributes
	Header     TableLine
	Columns    []TableColumn
	Lines      []TableLine
}

// parseNum like atoi, but stops on non-digit character (and only unsigned)
func (t Table) parseNum(c string) (int, string) {
	n := 0
	for c != "" && c[0] >= '0' && c[0] <= '9' {
		n *= 10
		n += int(c[0] - '0')
		c = c[1:]
	}
	return n, c
}

func (t Table) parseColumnsAttr() ([]TableColumn, error) {
	attr, ok := t.Attributes.GetAsString(AttrCols)
	if !ok {
		return nil, nil
	}
	var cols []TableColumn

	for _, c := range strings.Split(attr, ",") {
		c = strings.TrimSpace(c)
		n := 0
		repeat := 1
		n, c = t.parseNum(c)
		col := defaultColumn
		if c == "" {
			// If this was a number by itself, consider it a Width
			col.widthVal = float64(n)
			cols = append(cols, col)
			continue
		}

		if c[0] == '*' {
			repeat = n
			c = c[1:]
		} else if n != 0 {
			return nil, fmt.Errorf("bad column repeat (expected '*', got %s)", c)
		}

		// Horizontal alignment
		if c != "" {
			switch c[0] {
			case '<':
				col.HAlign = "left"
				c = c[1:]
			case '>':
				col.HAlign = "right"
				c = c[1:]
			case '^':
				col.HAlign = "center"
				c = c[1:]
			}
		}

		// Vertical alignment
		if len(c) >= 2 && c[0] == '.' {
			switch c[1] {
			case '<':
				col.VAlign = "top"
				c = c[2:]
			case '>':
				col.VAlign = "bottom"
				c = c[2:]
			case '^':
				col.VAlign = "middle"
				c = c[2:]
			default:
				return nil, errors.New("bad column vertical alignment")
			}
		}

		// Width
		if c != "" && c[0] == '~' {
			col.widthVal = 0 // Auto-width
			c = c[1:]
		} else if c != "" && c[0] >= '0' && c[0] <= '9' {
			n, c = t.parseNum(c)
			col.widthVal = float64(n)
		}

		// Style - must be the last item
		switch c {
		case "": // nothing left (no explicit style)
		case "a", "e", "h", "l", "m", "s", "v":
			col.Style = c
		case "d":
			col.Style = "" // leave default unset
		default:
			return nil, fmt.Errorf("bad column specification (%s unparsed)", c)
		}

		for repeat > 0 {
			cols = append(cols, col)
			repeat--
		}
	}

	return cols, nil
}

func (t Table) processColumnWidths() ([]TableColumn, error) {

	widths := make([]float64, len(t.Columns))
	cols := make([]TableColumn, 0, len(t.Columns))

	// Autowidth table uses full autowidth on all columns.
	if !t.Attributes.HasOption("autowidth") {

		percent := false
		total := 0.0
		for i, c := range t.Columns {
			if c.widthVal == 0 {
				percent = true
			}
			widths[i] = c.widthVal
			total += c.widthVal
		}

		if percent {
			// Zero or more fixed percentages.
			//  At least one column automatically expanding to fill remainder.
			if total > 100 {
				return nil, fmt.Errorf("total widths (%f) cannot exceed 100%%", total)
			}
		} else {
			// Relative widths, we have to calculate as percentages.
			used := 0.0
			for i, v := range widths {
				if i == len(widths)-1 {
					// Last column uses remainder -- addresses rounding errors.
					// (Also, faster, simpler math.)
					v = 100 - used
				} else {
					// This rounds to nearest .001 percent.  This allows us to
					// use %.6g format in templates without precision loss.
					v = v * 1000000 / total
					v = math.Round(v)
					v /= 10000
				}
				used += v
				widths[i] = v
			}
		}
	}

	// This is a little more complex because we use pass by value.
	for i, c := range t.Columns {
		if widths[i] != 0 {
			c.Width = strconv.FormatFloat(widths[i], 'g', 6, 64)
		}
		c.widthVal = 0
		cols = append(cols, c)
	}
	return cols, nil
}

// NewTable initializes a new table with the given lines and attributes
func NewTable(header interface{}, lines []interface{}, attributes interface{}) (Table, error) {
	t := Table{
		Attributes: toAttributes(attributes),
	}

	var err error
	if t.Columns, err = t.parseColumnsAttr(); err != nil {
		return Table{}, errors.Wrap(err, "failed to initialize a Table element")
	}

	if header, ok := header.(TableLine); ok {
		t.Header = header
		if t.Columns == nil {
			// columns determined by our cell count here
			for i := 0; i < len(header.Cells); i++ {
				t.Columns = append(t.Columns, defaultColumn)
			}
		}
	}
	// need to regroup columns of all lines, they dispatch on lines
	cells := make([][]interface{}, 0)
	for _, l := range lines {
		if l, ok := l.(TableLine); ok {
			// if no header line was set, inspect the first line to determine the number of columns per line
			if t.Columns == nil {
				for i := 0; i < len(l.Cells); i++ {
					t.Columns = append(t.Columns, defaultColumn)
				}
			}
			cells = append(cells, l.Cells...)
		}
	}

	// Calculate the actual widths now
	if t.Columns, err = t.processColumnWidths(); err != nil {
		return Table{}, errors.Wrap(err, "failed to initialize a Table element")
	}

	t.Lines = make([]TableLine, 0, len(cells))
	if len(lines) > 0 {
		log.Debugf("buffered %d columns for the table", len(cells))
		l := TableLine{
			Cells: make([][]interface{}, len(t.Columns)),
		}
		for i, c := range cells {
			log.Debugf("adding cell with content '%v' in table line at offset %d", c, i%len(t.Columns))
			l.Cells[i%len(t.Columns)] = c
			if (i+1)%len(t.Columns) == 0 { // switch to next line
				log.Debugf("adding line with content '%v' in table", l)
				t.Lines = append(t.Lines, l)
				l = TableLine{
					Cells: make([][]interface{}, len(t.Columns)),
				}
			}
		}
	}
	// log.Debugf("initialized a new table with %d line(s)", len(lines))
	return t, nil
}

// TableLine a table line is made of columns, each column being a group of []interface{} (to support quoted text, etc.)
type TableLine struct {
	Cells [][]interface{}
}

// NewTableLine initializes a new TableLine with the given columns
func NewTableLine(columns []interface{}) (TableLine, error) {
	c := make([][]interface{}, 0)
	for _, column := range columns {
		if e, ok := column.([]interface{}); ok {
			c = append(c, e)
		} else {
			return TableLine{}, errors.Errorf("unsupported element of type %T", column)
		}
	}
	// log.Debugf("initialized a new table line with %d columns", len(c))
	return TableLine{
		Cells: c,
	}, nil
}
