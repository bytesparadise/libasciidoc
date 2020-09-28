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
	caption := &strings.Builder{}
	number := 0
	title := r.renderElementTitle(img.Attributes)

	// Matching asciidoctor behavior, we increment the counter if we have a title,
	// regardless if we have a number or not. This will probably lead to confusion
	// if mixed custom and stock captioning is used.
	if _, found := img.Attributes.GetAsString(types.AttrTitle); found {
		c, ok := img.Attributes.GetAsString(types.AttrCaption)
		if !ok {
			c = ctx.Attributes.GetAsStringWithDefault(types.AttrFigureCaption, "Figure")
			if c != "" {
				// We always append the figure number, unless the caption is disabled.
				// This is for asciidoctor compatibility.
				c += " {counter:figure-number}. "
			}
		}
		// TODO: Replace this hack when we have attribute substitution fully working
		if strings.Contains(c, "{counter:figure-number}") {
			number = ctx.GetAndIncrementImageCounter()
			c = strings.ReplaceAll(c, "{counter:figure-number}", strconv.Itoa(number))
		}
		caption.WriteString(c)
	}
	roles, err := r.renderImageRoles(ctx, img.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render image roles")
	}
	err = r.blockImage.Execute(result, struct {
		ID          string
		Title       string
		ImageNumber int
		Caption     string
		Roles       string
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
		Roles:       roles,
		Href:        img.Attributes.GetAsStringWithDefault(types.AttrInlineLink, ""),
		Alt:         img.Attributes.GetAsStringWithDefault(types.AttrImageAlt, ""),
		Width:       img.Attributes.GetAsStringWithDefault(types.AttrWidth, ""),
		Height:      img.Attributes.GetAsStringWithDefault(types.AttrImageHeight, ""),
		Path:        img.Location.Stringify(),
	})

	if err != nil {
		return "", errors.Wrap(err, "unable to render block image")
	}
	// log.Debugf("rendered block image: %s", result.Bytes())
	return result.String(), nil
}

func (r *sgmlRenderer) renderInlineImage(ctx *Context, img types.InlineImage) (string, error) {
	result := &strings.Builder{}
	roles, err := r.renderImageRoles(ctx, img.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render image roles")
	}
	err = r.inlineImage.Execute(result, struct {
		Roles  string
		Title  string
		Href   string
		Alt    string
		Width  string
		Height string
		Path   string
	}{
		Title:  r.renderElementTitle(img.Attributes),
		Roles:  roles,
		Href:   img.Attributes.GetAsStringWithDefault(types.AttrInlineLink, ""),
		Alt:    img.Attributes.GetAsStringWithDefault(types.AttrImageAlt, ""),
		Width:  img.Attributes.GetAsStringWithDefault(types.AttrWidth, ""),
		Height: img.Attributes.GetAsStringWithDefault(types.AttrImageHeight, ""),
		Path:   img.Location.Stringify(),
	})

	if err != nil {
		return "", errors.Wrap(err, "unable to render inline image")
	}
	// log.Debugf("rendered inline image: %s", result.Bytes())
	return result.String(), nil
}
