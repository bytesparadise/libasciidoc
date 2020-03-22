package types

// Predefined the predefined document attributes, mainly for special characters
var Predefined map[string]string

func init() {
	Predefined = map[string]string{
		"sp":             " ",
		"blank":          "",
		"empty":          "",
		"nbsp":           "&#160;",
		"zwsp":           "&#8203;",
		"wj":             "&#8288;",
		"apos":           "&#39;",
		"quot":           "&#34;",
		"lsquo":          "&#8216;",
		"rsquo":          "&#8217;",
		"ldquo":          "&#8220;",
		"rdquo":          "&#8221;",
		"deg":            "&#176;",
		"plus":           "&#43;",
		"brvbar":         "&#166;",
		"vbar":           "|",
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
