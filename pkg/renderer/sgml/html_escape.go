package sgml

import (
	"strings"
)

// EscapeString is like html5.Escape func except but bypasses
// a few replacements.  It is a bit more conservative.
func EscapeString(s string) string {
	return htmlEscaper.Replace(s)
}

var htmlEscaper = strings.NewReplacer(
	`&lt;`, "&lt;", // keep as-is (we do not want `&amp;lt;`)
	`&gt;`, "&gt;", // keep `&lgt;` as-is (we do not want `&amp;gt;`)
	`&amp;`, "&amp;", // keep `&amps` as-is (we do not want `&amp;amp;`)
	`&#`, "&#", // assume this is for an character entity and this keep as-is
	// standard escape combinations
	`&`, "&amp;",
	`<`, "&lt;",
	`>`, "&gt;",
	// TODO: These two should be substituted as well.  The elements here could wind up in attributes.
	// `'`, "&#39;", // "&#39;" is shorter than "&apos;" and apos was not in HTML until HTML5.
	// `"`, "&#34;", // "&#34;" is shorter than "&quot;".
)
