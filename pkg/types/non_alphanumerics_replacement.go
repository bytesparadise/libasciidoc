package types

import (
	"strings"
	"unicode"
)

// ReplaceNonAlphanumerics replace all non alpha numeric characters with the given `replacement`
func ReplaceNonAlphanumerics(elements []interface{}, prefix, separator string) string {
	replacement := replaceNonAlphanumericsOnElements(elements, separator)
	// avoid duplicate prefix
	if strings.HasPrefix(replacement, prefix) {
		return replacement
	}
	return prefix + replacement
}

func replaceNonAlphanumericsOnElements(elements []interface{}, separator string) string {
	result := &strings.Builder{}
	for i, element := range elements {
		switch e := element.(type) {
		case *QuotedText:
			r := replaceNonAlphanumericsOnElements(e.Elements, separator)
			result.WriteString(r)
			result.WriteString(separator)
		case *StringElement:
			content := e.Content
			if i == 0 {
				// trim heading spaces only if this StringElement is in first position
				content = strings.TrimLeft(e.Content, " ")
			}
			r := replaceNonAlphanumerics(content, separator)
			result.WriteString(r)
		case *Symbol:
			if e.Prefix != "" {
				result.WriteString(e.Prefix)
			}
		case *InlineLink:
			if e.Location != nil {
				r := replaceNonAlphanumerics(e.Location.Stringify(), separator)
				result.WriteString(r)
				result.WriteString(separator)
			}
		case *Icon:
			s := e.Attributes.GetAsStringWithDefault(AttrImageAlt, e.Class)
			r := replaceNonAlphanumerics(s, separator)
			result.WriteString(r)
			result.WriteString(separator)
		default:
			// other types are ignored
		}
	}
	r := strings.TrimSuffix(result.String(), separator)
	// avoid duplicate separators
	return strings.ReplaceAll(r, separator+separator, separator)
}

func replaceNonAlphanumerics(content, replacement string) string {
	buf := &strings.Builder{}
	lastCharIsSeparator := false

	// Drop the :// from links.
	content = strings.ReplaceAll(content, "://", "")

	for _, r := range content {
		switch {
		case unicode.Is(unicode.Letter, r) || unicode.Is(unicode.Number, r):
			buf.WriteString(strings.ToLower(string(r)))
			lastCharIsSeparator = false
		case !lastCharIsSeparator && (string(r) == " " || string(r) == "-" || string(r) == "."):
			buf.WriteString(replacement)
			lastCharIsSeparator = true
		}
	}
	// result := strings.TrimSuffix(buf.String(), replacement)
	result := buf.String()
	return result
}
