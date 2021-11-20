package types

import (
	"strings"
	"unicode"

	"github.com/pkg/errors"
)

// ReplaceNonAlphanumerics replace all non alpha numeric characters with the given `replacement`
func ReplaceNonAlphanumerics(elements []interface{}, replacement string) (string, error) {
	buf := &strings.Builder{}
	for _, element := range elements {
		switch element := element.(type) {
		case *QuotedText:
			r, err := ReplaceNonAlphanumerics(element.Elements, replacement)
			if err != nil {
				return "", err
			}
			if buf.Len() > 0 {
				buf.WriteString(replacement)
			}
			buf.WriteString(r)
		case *StringElement:
			r, err := replaceNonAlphanumerics(element.Content, replacement)
			if err != nil {
				return "", err
			}
			if buf.Len() > 0 {
				buf.WriteString(replacement)
			}
			buf.WriteString(r)
		case *InlineLink:
			r, err := replaceNonAlphanumerics(element.Location.Stringify(), replacement)
			if err != nil {
				return "", err
			}
			if buf.Len() > 0 {
				buf.WriteString(replacement)
			}
			buf.WriteString(r)
		case *Icon:
			s := element.Attributes.GetAsStringWithDefault(AttrImageAlt, element.Class)
			r, err := replaceNonAlphanumerics(s, replacement)
			if err != nil {
				return "", err
			}
			if buf.Len() > 0 {
				buf.WriteString(replacement)
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
	lastCharIsSpace := false

	// Drop the :// from links.
	content = strings.ReplaceAll(content, "://", "")

	for _, r := range strings.TrimLeft(content, " ") { // ignore header spaces
		if unicode.Is(unicode.Letter, r) || unicode.Is(unicode.Number, r) {
			_, err := buf.WriteString(strings.ToLower(string(r)))
			if err != nil {
				return "", errors.Wrapf(err, "error while normalizing String Element")
			}
			lastCharIsSpace = false
		} else if !lastCharIsSpace && (unicode.Is(unicode.Space, r) || unicode.Is(unicode.Punct, r)) {
			_, err := buf.WriteString(replacement)
			if err != nil {
				return "", errors.Wrapf(err, "error while normalizing String Element")
			}
			lastCharIsSpace = true
		}
	}
	result := strings.TrimSuffix(buf.String(), replacement)
	// log.Debugf("normalized '%s' to '%s'", content, result)
	return result, nil
}
