package parser

import (
	"fmt"
	"sort"
	"strings"
)

func PrettyPrintStats(stats *Stats) string {
	rules := make([]string, 0, len(stats.ChoiceAltCnt))
	for r := range stats.ChoiceAltCnt {
		rules = append(rules, r)
	}
	sort.Strings(rules)
	buf := &strings.Builder{}
	for _, r := range rules {
		buf.WriteString(fmt.Sprintf("%s ", r))
		for choice, count := range stats.ChoiceAltCnt[r] {
			buf.WriteString(fmt.Sprintf("| %s->%dx ", choice, count))
		}
		buf.WriteString("\n")
	}
	buf.WriteString(fmt.Sprintf("---------------\nTotal: %d\n\n", stats.ExprCnt))
	return buf.String()
}
