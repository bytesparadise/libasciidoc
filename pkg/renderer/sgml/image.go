package sgml

import (
	"net/url"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderImageBlock(ctx *renderer.Context, img types.ImageBlock) (string, error) {
	result := &strings.Builder{}
	title, err := r.renderElementTitle(img.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render image")
	}

	// Matching asciidoctor behavior, we increment the counter if we have a title,
	// regardless if we have a number or not. This will probably lead to confusion
	// if mixed custom and stock captioning is used.
	caption := &strings.Builder{}
	number := 0
	if title != "" {
		c, found, err := img.Attributes.GetAsString(types.AttrCaption)
		if err != nil {
			return "", errors.Wrap(err, "unable to render image")
		} else if !found {
			if c = ctx.Attributes.GetAsStringWithDefault(types.AttrFigureCaption, "Figure"); c != "" {
				// We always append the figure number, unless the caption is disabled.
				// This is for asciidoctor compatibility.
				c += " {counter:figure-number}. "
			}
		}
		if strings.Contains(c, "{counter:figure-number}") {
			// TODO: Replace this hack when we have attribute substitution fully working
			number = ctx.GetAndIncrementImageCounter()
			c = strings.ReplaceAll(c, "{counter:figure-number}", strconv.Itoa(number))
		}
		caption.WriteString(c)
	}
	roles, err := r.renderImageRoles(ctx, img.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render image")
	}
	path := img.Location.Stringify()
	alt, err := r.renderImageAlt(img.Attributes, path)
	if err != nil {
		return "", errors.Wrap(err, "unable to render image")
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
		Alt:         alt,
		Width:       img.Attributes.GetAsStringWithDefault(types.AttrWidth, ""),
		Height:      img.Attributes.GetAsStringWithDefault(types.AttrHeight, ""),
		Path:        path,
	})

	if err != nil {
		return "", errors.Wrap(err, "unable to render image")
	}
	return result.String(), nil
}

func (r *sgmlRenderer) renderInlineImage(ctx *Context, img types.InlineImage) (string, error) {
	result := &strings.Builder{}
	roles, err := r.renderImageRoles(ctx, img.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render image")
	}
	path := img.Location.Stringify()
	alt, err := r.renderImageAlt(img.Attributes, path)
	if err != nil {
		return "", errors.Wrap(err, "unable to render image")
	}
	title, err := r.renderElementTitle(img.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render callout list roles")
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
		Title:  title,
		Roles:  roles,
		Href:   img.Attributes.GetAsStringWithDefault(types.AttrInlineLink, ""),
		Alt:    alt,
		Width:  img.Attributes.GetAsStringWithDefault(types.AttrWidth, ""),
		Height: img.Attributes.GetAsStringWithDefault(types.AttrHeight, ""),
		Path:   path,
	})

	if err != nil {
		return "", errors.Wrap(err, "unable to render inline image")
	}
	// log.Debugf("rendered inline image: %s", result.Bytes())
	return result.String(), nil
}

func (r *sgmlRenderer) renderImageAlt(attrs types.Attributes, path string) (string, error) {
	if alt, found, err := attrs.GetAsString(types.AttrImageAlt); err != nil {
		return "", errors.Wrap(err, "unable to render image")
	} else if found {
		return alt, nil
	}
	u, err := url.Parse(path)
	if err != nil {
		return "", errors.Wrap(err, "unable to render image")
	}
	// return base path without its extension
	return strings.TrimSuffix(filepath.Base(u.Path), filepath.Ext(u.Path)), nil
}
