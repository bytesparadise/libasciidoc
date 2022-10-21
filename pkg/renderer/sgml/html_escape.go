package sgml

import (
	"html"
	"strings"
)

// escapeString is like html5.Escape func except but bypasses
// a few replacements.  It is a bit more conservative.
func escapeString(s string) string {
	return htmlEscaper.Replace(s)
}

func unescapeString(s string) string {
	return html.UnescapeString(s)
}

var htmlEscaper = strings.NewReplacer(
	`&lt;`, "&lt;", // keep as-is (we do not want `&amp;lt;`)
	`&gt;`, "&gt;", // keep `&lgt;` as-is (we do not want `&amp;gt;`)
	`&amp;`, "&amp;", // keep `&amps` as-is (we do not want `&amp;amp;`) // TODO: still needed?
	`&#`, "&#", // assume this is for an character entity and this keep as-is
	// standard escape combinations
	`&`, "&amp;",
	`<`, "&lt;",
	`>`, "&gt;",
	`'`, "&#39;", // "&#39;" is shorter than "&apos;" and apos was not in HTML until HTML5.
	`"`, "&#34;", // "&#34;" is shorter than "&quot;".
)
