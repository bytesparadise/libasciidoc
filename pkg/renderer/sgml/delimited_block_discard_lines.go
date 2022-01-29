package sgml

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
)

func discardBlankLines(lines []interface{}) []interface{} {
	// discard blank elements at the end
	// log.Debugf("discarding blank lines on %d elements...", len(lines))
	filteredLines := make([]interface{}, len(lines))
	copy(filteredLines, lines)
	// leading empty lines
	for {
		if len(filteredLines) == 0 {
			break
		}
		if _, ok := filteredLines[0].(types.BlankLine); ok {
			// remove last element of the slice since it's a blank line
			filteredLines = filteredLines[:len(filteredLines)-1]
		} else {
			break
		}
	}
	// trailing empty lines
	for {
		if len(filteredLines) == 0 {
			break
		}
		if _, ok := filteredLines[len(filteredLines)-1].(types.BlankLine); ok {
			// remove last element of the slice since it's a blank line
			filteredLines = filteredLines[:len(filteredLines)-1]
		} else {
			break
		}
	}
	return filteredLines
}
