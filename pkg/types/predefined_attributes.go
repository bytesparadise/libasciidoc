package types

// Predefined the predefined document attributes, mainly for special characters
var Predefined map[string]string

func init() {
	Predefined = map[string]string{
		"sp":             " ",
		"blank":          "",
		"empty":          "",
		"nbsp":           "\u00a0",
		"zwsp":           "\u200b",
		"wj":             "\u2060",
		"apos":           "&#39;",
		"quot":           "&#34;",
		"lsquo":          "\u2018",
		"rsquo":          "\u2019",
		"ldquo":          "\u201c",
		"rdquo":          "\u201d",
		"deg":            "\u00b0",
		"plus":           "&#43;", // leave this to prevent passthrough decode?
		"brvbar":         "\u00a6",
		"vbar":           "|", // TODO: maybe convert this because of tables?
		"amp":            "&amp;",
		"lt":             "&lt;",
		"gt":             "&gt;",
		"startsb":        "[",
		"endsb":          "]",
		"caret":          "^",
		"asterisk":       "*",
		"tilde":          "~",
		"backslash":      `\`,
		"backtick":       "`",
		"two-colons":     "::",
		"two-semicolons": ";",
		"cpp":            "C++",
	}
}
