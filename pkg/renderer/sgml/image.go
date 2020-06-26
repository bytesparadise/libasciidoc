package sgml

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderImageBlock(ctx *renderer.Context, img types.ImageBlock) (string, error) {
	result := &strings.Builder{}
	number := 0
	if _, found := img.Attributes.GetAsString(types.AttrTitle); found {
		number = ctx.GetAndIncrementImageCounter()
	}
	err := r.blockImage.Execute(result, struct {
		ID          sanitized
		Title       sanitized
		ImageNumber int
		Role        string
		Href        string
		Alt         string
		Width       string
		Height      string
		Path        string
	}{
		ID:          r.renderElementID(img.Attributes),
		Title:       r.renderElementTitle(img.Attributes),
		ImageNumber: number,
		Role:        img.Attributes.GetAsStringWithDefault(types.AttrRole, ""),
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
