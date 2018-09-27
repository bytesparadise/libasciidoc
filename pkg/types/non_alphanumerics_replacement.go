package types

import (
	"bytes"
	"strings"
	"unicode"

	"github.com/pkg/errors"
)

//NormalizationFunc a function that is used to normalize a string.
type NormalizationFunc func(string) ([]byte, error)

// newReplaceNonAlphanumericsFunc replaces all non alphanumerical characters and remove (accents)
// in the given 'source' with the given 'replacement'.
func newReplaceNonAlphanumericsFunc(replacement string) NormalizationFunc {
	return func(source string) ([]byte, error) {
		buf := bytes.NewBuffer(nil)
		lastCharIsSpace := false
		for _, r := range strings.TrimLeft(source, " ") { // ignore header spaces
			if unicode.Is(unicode.Letter, r) || unicode.Is(unicode.Number, r) {
				_, err := buf.WriteString(strings.ToLower(string(r)))
				if err != nil {
					return nil, errors.Wrapf(err, "unable to normalize value")
				}
				lastCharIsSpace = false
			} else if !lastCharIsSpace && (unicode.Is(unicode.Space, r) || unicode.Is(unicode.Punct, r)) {
				_, err := buf.WriteString(replacement)
				if err != nil {
					return nil, errors.Wrapf(err, "unable to normalize value")
				}
				lastCharIsSpace = true
			}
		}
		result := strings.TrimSuffix(buf.String(), replacement)
		return []byte(result), nil
	}
}

// replaceNonAlphanumerics replace all non alpha numeric characters with the given `replacement`
func replaceNonAlphanumerics(source InlineElements, replacement string) (string, error) {
	v := newreplaceNonAlphanumericsVisitor()
	err := source.Accept(&v)
	if err != nil {
		return "", err
	}
	return v.normalizedContent(), nil
}

//replaceNonAlphanumericsVisitor a visitor that builds a string representation of the visited elements,
// in which all non-alphanumeric characters have been replaced with a "_"
type replaceNonAlphanumericsVisitor struct {
	buf       bytes.Buffer
	normalize NormalizationFunc
}

// newreplaceNonAlphanumericsVisitor returns a new replaceNonAlphanumericsVisitor
func newreplaceNonAlphanumericsVisitor() replaceNonAlphanumericsVisitor {
	buf := bytes.NewBuffer(nil)
	return replaceNonAlphanumericsVisitor{
		buf:       *buf,
		normalize: newReplaceNonAlphanumericsFunc("_"),
	}
}

// Visit method called when an element is visited
func (v *replaceNonAlphanumericsVisitor) Visit(element Visitable) error {
	// log.Debugf("visiting element of type '%T'", element)
	if element, ok := element.(StringElement); ok {
		v.buf.WriteString("_")
		normalized, err := v.normalize(element.Content)
		if err != nil {
			return errors.Wrapf(err, "error while normalizing String Element")
		}
		v.buf.Write(normalized)
	}
	// other types are ignored
	return nil
}

// NormalizedContent returns the normalized content
func (v *replaceNonAlphanumericsVisitor) normalizedContent() string {
	return v.buf.String()
}
