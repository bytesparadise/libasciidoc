package sgml

import (
	"bytes"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderStringElement(_ *renderer.Context, str types.StringElement) ([]byte, error) {
	buf := &bytes.Buffer{}
	err := r.stringElement.Execute(buf, str.Content)
	if err != nil {
		return []byte{}, errors.Wrapf(err, "unable to render string")
	}

	// NB: For all SGML flavors we are aware of, the numeric entities from
	// Unicode are supported.  We generally avoid named entities.
	result := convert(buf.String(), ellipsis, copyright, trademark, registered)
	return []byte(result), nil
}

func ellipsis(source string) string {
	return strings.Replace(source, "...", "&#8230;&#8203;", -1)
}

func copyright(source string) string {
	return strings.Replace(source, "(C)", "&#169;", -1)
}

func trademark(source string) string {
	return strings.Replace(source, "(TM)", "&#153;", -1)
}

func registered(source string) string {
	return strings.Replace(source, "(R)", "&#174;", -1)
}

type converter func(string) string

func convert(source string, converters ...converter) string {
	result := source
	for _, convert := range converters {
		result = convert(result)
	}
	return result
}
