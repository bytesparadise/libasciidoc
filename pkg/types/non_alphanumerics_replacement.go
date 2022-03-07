package types

import (
	"strings"
	"unicode"

	"github.com/pkg/errors"
)

// ReplaceNonAlphanumerics replace all non alpha numeric characters with the given `replacement`
func ReplaceNonAlphanumerics(elements []interface{}, prefix, separator string) (string, error) {
	replacement, err := replaceNonAlphanumericsOnElements(elements, separator)
	if err != nil {
		return "", err
	}
	// avoid double prefix
	if strings.HasPrefix(replacement, prefix) {
		return replacement, nil
	}
	return prefix + replacement, nil
}

func replaceNonAlphanumericsOnElements(elements []interface{}, separator string) (string, error) {
	result := &strings.Builder{}
	for _, element := range elements {
		switch e := element.(type) {
		case *QuotedString:
			r, err := replaceNonAlphanumericsOnElements(e.Elements, separator)
			if err != nil {
				return "", err
			}
			// buf.WriteString(separator)
			result.WriteString(r)
			result.WriteString(separator)
		case *QuotedText:
			r, err := replaceNonAlphanumericsOnElements(e.Elements, separator)
			if err != nil {
				return "", err
			}
			result.WriteString(r)
			result.WriteString(separator)
		case *StringElement:
			r, err := replaceNonAlphanumerics(e.Content, separator)
			if err != nil {
				return "", err
			}
			result.WriteString(r)
		case *Symbol:
			if e.Prefix != "" {
				result.WriteString(e.Prefix)
			}
		case *InlineLink:
			if e.Location != nil {
				r, err := replaceNonAlphanumerics(e.Location.Stringify(), separator)
				if err != nil {
					return "", err
				}
				result.WriteString(r)
				result.WriteString(separator)
			}
		case *Icon:
			s := e.Attributes.GetAsStringWithDefault(AttrImageAlt, e.Class)
			r, err := replaceNonAlphanumerics(s, separator)
			if err != nil {
				return "", err
			}
			result.WriteString(r)
			result.WriteString(separator)
		default:
			// other types are ignored
		}
	}
	return strings.TrimSuffix(result.String(), separator), nil
}

func replaceNonAlphanumerics(content, replacement string) (string, error) {
	buf := &strings.Builder{}
	lastCharIsSeparator := false

	// Drop the :// from links.
	content = strings.ReplaceAll(content, "://", "")

	for _, r := range strings.TrimLeft(content, " ") { // ignore header spaces
		switch {
		case unicode.Is(unicode.Letter, r) || unicode.Is(unicode.Number, r):
			_, err := buf.WriteString(strings.ToLower(string(r)))
			if err != nil {
				return "", errors.Wrapf(err, "error while normalizing String Element")
			}
			lastCharIsSeparator = false
		case !lastCharIsSeparator && (string(r) == " " || string(r) == "-" || string(r) == "."):
			_, err := buf.WriteString(replacement)
			if err != nil {
				return "", errors.Wrapf(err, "error while normalizing String Element")
			}
			lastCharIsSeparator = true
		}
	}
	// result := strings.TrimSuffix(buf.String(), replacement)
	result := buf.String()
	return result, nil
}
