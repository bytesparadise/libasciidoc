package sgml

import (
	"bytes"
	"strconv"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderImageBlock(ctx *renderer.Context, img types.ImageBlock) ([]byte, error) {
	result := &bytes.Buffer{}
	title := ""
	if t, found := img.Attributes.GetAsString(types.AttrTitle); found {
		title = "Figure " + strconv.Itoa(ctx.GetAndIncrementImageCounter()) + ". " + EscapeString(t)
	}
	err := r.blockImage.Execute(result, struct {
		ID     string
		Title  string
		Role   string
		Href   string
		Alt    string
		Width  string
		Height string
		Path   string
	}{
		ID:     img.Attributes.GetAsStringWithDefault(types.AttrID, ""),
		Title:  title,
		Role:   img.Attributes.GetAsStringWithDefault(types.AttrRole, ""),
		Href:   img.Attributes.GetAsStringWithDefault(types.AttrInlineLink, ""),
		Alt:    img.Attributes.GetAsStringWithDefault(types.AttrImageAlt, ""),
		Width:  img.Attributes.GetAsStringWithDefault(types.AttrImageWidth, ""),
		Height: img.Attributes.GetAsStringWithDefault(types.AttrImageHeight, ""),
		Path:   img.Location.String(),
	})

	if err != nil {
		return nil, errors.Wrapf(err, "unable to render block image")
	}
	// log.Debugf("rendered block image: %s", result.Bytes())
	return result.Bytes(), nil
}

func (r *sgmlRenderer) renderInlineImage(img types.InlineImage) ([]byte, error) {
	result := &bytes.Buffer{}
	err := r.inlineImage.Execute(result, struct {
		Role   string
		Title  string
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
		return nil, errors.Wrapf(err, "unable to render inline image")
	}
	// log.Debugf("rendered inline image: %s", result.Bytes())
	return result.Bytes(), nil
}
