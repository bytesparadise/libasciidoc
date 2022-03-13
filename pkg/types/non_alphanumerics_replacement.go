package types

import (
	"strings"
	"unicode"
)

// ReplaceNonAlphanumerics replace all non alpha numeric characters with the given `replacement`
func ReplaceNonAlphanumerics(elements []interface{}, prefix, separator string) (string, error) {
	replacement, err := replaceNonAlphanumericsOnElements(elements, separator)
	if err != nil {
		return "", err
	}
	// avoid duplicate prefix
	if strings.HasPrefix(replacement, prefix) {
		return replacement, nil
	}
	return prefix + replacement, nil
}

func replaceNonAlphanumericsOnElements(elements []interface{}, separator string) (string, error) {
	result := &strings.Builder{}
	for i, element := range elements {
		switch e := element.(type) {
		case *QuotedText:
			r, err := replaceNonAlphanumericsOnElements(e.Elements, separator)
			if err != nil {
				return "", err
			}
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
	r = strings.ReplaceAll(r, separator+separator, separator)
	return r, nil
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
