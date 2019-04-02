package types

import (
	"bytes"
	"strings"
	"unicode"

	"github.com/pkg/errors"
)

// replaceNonAlphanumerics replace all non alpha numeric characters with the given `replacement`
func replaceNonAlphanumerics(source InlineElements, replacement string) (string, error) {
	v := newReplaceNonAlphanumericsVisitor(replacement)
	err := source.AcceptVisitor(v)
	if err != nil {
		return "", err
	}
	return v.normalizedContent(), nil
}

//replaceNonAlphanumericsVisitor a visitor that builds a string representation of the visited elements,
// in which all non-alphanumeric characters have been replaced with a "_"
type replaceNonAlphanumericsVisitor struct {
	buf         bytes.Buffer
	replacement string
}

var _ Visitor = &replaceNonAlphanumericsVisitor{}

// newReplaceNonAlphanumericsVisitor returns a new replaceNonAlphanumericsVisitor
func newReplaceNonAlphanumericsVisitor(replacement string) *replaceNonAlphanumericsVisitor {
	buf := bytes.NewBuffer(nil)
	return &replaceNonAlphanumericsVisitor{
		buf:         *buf,
		replacement: replacement,
	}
}

// Visit method called when an element is visited
func (v *replaceNonAlphanumericsVisitor) Visit(element Visitable) error {
	// log.Debugf("visiting element of type '%T'", element)
	if element, ok := element.(StringElement); ok {
		if v.buf.Len() > 0 {
			v.buf.WriteString("_")
		}
		normalized, err := v.normalize(element.Content)
		if err != nil {
			return errors.Wrapf(err, "error while normalizing String Element")
		}
		v.buf.WriteString(normalized)
		return nil
	}
	// other types are ignored
	return nil
}

// normalize returns the normalized content
func (v *replaceNonAlphanumericsVisitor) normalize(source string) (string, error) {
	buf := bytes.NewBuffer(nil)
	lastCharIsSpace := false
	for _, r := range strings.TrimLeft(source, " ") { // ignore header spaces
		if unicode.Is(unicode.Letter, r) || unicode.Is(unicode.Number, r) {
			_, err := buf.WriteString(strings.ToLower(string(r)))
			if err != nil {
				return "", errors.Wrapf(err, "unable to normalize value")
			}
			lastCharIsSpace = false
		} else if !lastCharIsSpace && (unicode.Is(unicode.Space, r) || unicode.Is(unicode.Punct, r)) {
			_, err := buf.WriteString(v.replacement)
			if err != nil {
				return "", errors.Wrapf(err, "unable to normalize value")
			}
			lastCharIsSpace = true
		}
	}
	return strings.TrimSuffix(buf.String(), v.replacement), nil
}

// NormalizedContent returns the normalized content
func (v *replaceNonAlphanumericsVisitor) normalizedContent() string {
	return v.buf.String()
}
