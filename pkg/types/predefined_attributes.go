package types

// PredefinedAttributes the predefined document attributes
// May be converted into HTML entities
var predefinedAttributes = []string{
	"sp",
	"blank",
	"empty",
	"nbsp",
	"zwsp",
	"wj",
	"apos",
	"quot",
	"lsquo",
	"rsquo",
	"ldquo",
	"rdquo",
	"deg",
	"plus",
	"brvbar",
	"vbar",
	"amp",
	"lt",
	"gt",
	"startsb",
	"endsb",
	"caret",
	"asterisk",
	"tilde",
	"backslash",
	"backtick",
	"two-colons",
	"two-semicolons",
	"cpp",
}

func isPrefedinedAttribute(a string) bool {
	for _, v := range predefinedAttributes {
		if v == a {
			return true
		}
	}
	return false
}
