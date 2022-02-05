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
	buf := &strings.Builder{}
	for _, element := range elements {
		switch element := element.(type) {
		case *QuotedString:
			r, err := replaceNonAlphanumericsOnElements(element.Elements, separator)
			if err != nil {
				return "", err
			}
			if buf.Len() > 0 {
				buf.WriteString(separator)
			}
			buf.WriteString(r)
		case *QuotedText:
			r, err := replaceNonAlphanumericsOnElements(element.Elements, separator)
			if err != nil {
				return "", err
			}
			if buf.Len() > 0 {
				buf.WriteString(separator)
			}
			buf.WriteString(r)
		case *StringElement:
			r, err := replaceNonAlphanumerics(element.Content, separator)
			if err != nil {
				return "", err
			}
			if buf.Len() > 0 {
				buf.WriteString(separator)
			}
			buf.WriteString(r)
		case *InlineLink:
			if element.Location != nil {
				r, err := replaceNonAlphanumerics(element.Location.Stringify(), separator)
				if err != nil {
					return "", err
				}
				if buf.Len() > 0 {
					buf.WriteString(separator)
				}
				buf.WriteString(r)
			}
		case *Icon:
			s := element.Attributes.GetAsStringWithDefault(AttrImageAlt, element.Class)
			r, err := replaceNonAlphanumerics(s, separator)
			if err != nil {
				return "", err
			}
			if buf.Len() > 0 {
				buf.WriteString(separator)
			}
			buf.WriteString(r)
		default:
			// other types are ignored
		}
	}

	// log.Debugf("normalized '%+v' to '%s'", elements, buf.String())
	return buf.String(), nil
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
	result := strings.TrimSuffix(buf.String(), replacement)
	return result, nil
}
