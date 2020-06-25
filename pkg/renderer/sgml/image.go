package sgml

import (
	"strconv"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderImageBlock(ctx *renderer.Context, img types.ImageBlock) (string, error) {
	result := &strings.Builder{}
	title := sanitized("")
	if t, found := img.Attributes.GetAsString(types.AttrTitle); found {
		// TODO: This should be moved to the template
		title = sanitized("Figure " + strconv.Itoa(ctx.GetAndIncrementImageCounter()) + ". " + EscapeString(t))
	}
	err := r.blockImage.Execute(result, struct {
		ID     sanitized
		Title  sanitized
		Role   string
		Href   string
		Alt    string
		Width  string
		Height string
		Path   string
	}{
		ID:     r.renderElementID(img.Attributes),
		Title:  title,
		Role:   img.Attributes.GetAsStringWithDefault(types.AttrRole, ""),
		Href:   img.Attributes.GetAsStringWithDefault(types.AttrInlineLink, ""),
		Alt:    img.Attributes.GetAsStringWithDefault(types.AttrImageAlt, ""),
		Width:  img.Attributes.GetAsStringWithDefault(types.AttrImageWidth, ""),
		Height: img.Attributes.GetAsStringWithDefault(types.AttrImageHeight, ""),
		Path:   img.Location.String(),
	})

	if err != nil {
		return "", errors.Wrap(err, "unable to render block image")
	}
	// log.Debugf("rendered block image: %s", result.Bytes())
	return result.String(), nil
}

func (r *sgmlRenderer) renderInlineImage(img types.InlineImage) (string, error) {
	result := &strings.Builder{}
	err := r.inlineImage.Execute(result, struct {
		Role   string
		Title  sanitized
		Href   string
		Alt    string
		Width  string
		Height string
		Path   string
	}{
		Title:  r.renderElementTitle(img.Attributes),
		Role:   img.Attributes.GetAsStringWithDefault(types.AttrRole, ""),
		Alt:    img.Attributes.GetAsStringWithDefault(types.AttrImageAlt, ""),
		Width:  img.Attributes.GetAsStringWithDefault(types.AttrImageWidth, ""),
		Height: img.Attributes.GetAsStringWithDefault(types.AttrImageHeight, ""),
		Path:   img.Location.String(),
	})

	if err != nil {
		return "", errors.Wrap(err, "unable to render inline image")
	}
	// log.Debugf("rendered inline image: %s", result.Bytes())
	return result.String(), nil
}
