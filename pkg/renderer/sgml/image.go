package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderImageBlock(ctx *renderer.Context, img types.ImageBlock) (string, error) {
	result := &strings.Builder{}
	caption := &strings.Builder{}
	number := 0
	title := r.renderElementTitle(img.Attributes)

	// Matching asciidoctor behavior, we increment the counter if we have a title,
	// regardless if we have a number or not. This will probably lead to confusion
	// if mixed custom and stock captioning is used.
	if _, found := img.Attributes.GetAsString(types.AttrTitle); found {
		number = ctx.GetAndIncrementImageCounter()

		if s, ok := img.Attributes.GetAsString(types.AttrImageCaption); ok {
			caption.WriteString(s)
		} else {
			err := r.imageCaption.Execute(caption, struct {
				ImageNumber int
				Title       sanitized
			}{
				ImageNumber: number,
				Title:       title,
			})
			if err != nil {
				return "", errors.Wrap(err, "unable to format image caption")
			}
		}
	}
	err := r.blockImage.Execute(result, struct {
		ID          sanitized
		Title       sanitized
		ImageNumber int
		Caption     string
		Roles       sanitized
		Href        string
		Alt         string
		Width       string
		Height      string
		Path        string
	}{
		ID:          r.renderElementID(img.Attributes),
		Title:       title,
		ImageNumber: number,
		Caption:     caption.String(),
		Roles:       r.renderImageRoles(img.Attributes),
		Href:        img.Attributes.GetAsStringWithDefault(types.AttrInlineLink, ""),
		Alt:         img.Attributes.GetAsStringWithDefault(types.AttrImageAlt, ""),
		Width:       img.Attributes.GetAsStringWithDefault(types.AttrImageWidth, ""),
		Height:      img.Attributes.GetAsStringWithDefault(types.AttrImageHeight, ""),
		Path:        img.Location.String(),
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
		Roles  sanitized
		Title  sanitized
		Href   string
		Alt    string
		Width  string
		Height string
		Path   string
	}{
		Title:  r.renderElementTitle(img.Attributes),
		Roles:  r.renderImageRoles(img.Attributes),
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
