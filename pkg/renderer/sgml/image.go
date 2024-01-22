package sgml

import (
	"encoding/base64"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func (r *sgmlRenderer) renderImageBlock(ctx *context, img *types.ImageBlock) (string, error) {
	title, err := r.renderElementTitle(ctx, img.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render image")
	}

	// Matching asciidoctor behavior, we increment the counter if we have a title,
	// regardless if we have a number or not. This will probably lead to confusion
	// if mixed custom and stock captioning is used.
	caption := &strings.Builder{}
	number := 0
	if title != "" {
		c, found := img.Attributes.GetAsString(types.AttrCaption)
		if !found {
			c, found = ctx.attributes.GetAsString(types.AttrFigureCaption)
			if found && c != "" {
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
	src := r.getImageSrc(ctx, img.Location)
	alt, err := r.renderImageAlt(img.Attributes, img.Location.ToString())
	if err != nil {
		return "", errors.Wrap(err, "unable to render image")
	}
	return r.execute(r.blockImage, struct {
		ID          string
		Src         string
		Title       string
		ImageNumber int
		Caption     string
		Roles       string
		Href        string
		Alt         string
		Width       string
		Height      string
	}{
		ID:          r.renderElementID(img.Attributes),
		Src:         src,
		Title:       title,
		ImageNumber: number,
		Caption:     caption.String(),
		Roles:       roles,
		Href:        img.Attributes.GetAsStringWithDefault(types.AttrInlineLink, ""),
		Alt:         alt,
		Width:       img.Attributes.GetAsStringWithDefault(types.AttrWidth, ""),
		Height:      img.Attributes.GetAsStringWithDefault(types.AttrHeight, ""),
	})
}

func (r *sgmlRenderer) renderInlineImage(ctx *context, img *types.InlineImage) (string, error) {
	roles, err := r.renderImageRoles(ctx, img.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render inline image")
	}
	href := img.Attributes.GetAsStringWithDefault(types.AttrInlineLink, "")
	src := r.getImageSrc(ctx, img.Location)
	alt, err := r.renderImageAlt(img.Attributes, src)
	if err != nil {
		return "", errors.Wrap(err, "unable to render inline image")
	}
	title, err := r.renderElementTitle(ctx, img.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render inline image roles")
	}
	return r.execute(r.inlineImage, struct {
		Src    string
		Roles  string
		Title  string
		Href   string
		Alt    string
		Width  string
		Height string
	}{
		Src:    src,
		Title:  title,
		Roles:  roles,
		Href:   href,
		Alt:    alt,
		Width:  img.Attributes.GetAsStringWithDefault(types.AttrWidth, ""),
		Height: img.Attributes.GetAsStringWithDefault(types.AttrHeight, ""),
	})
}

func (r *sgmlRenderer) getImageSrc(ctx *context, location *types.Location) string {
	if imagesdir, found := ctx.attributes.GetAsString(types.AttrImagesDir); found {
		location.SetPathPrefix(imagesdir)
	}
	src := location.ToString()

	// if Data URI is enables, then include the content of the file in the `src` attribute of the `<img>` tag
	if !ctx.attributes.Has("data-uri") {
		return src
	}
	dir := filepath.Dir(ctx.config.Filename)
	src = filepath.Join(dir, src)
	result := "data:image/" + strings.TrimPrefix(filepath.Ext(src), ".") + ";base64,"
	data, err := os.ReadFile(src)
	if err != nil {
		log.Warnf("image to embed not found or not readable: %s", src)
		return result
	}
	result += base64.StdEncoding.EncodeToString(data)
	return result
}

func (r *sgmlRenderer) renderImageAlt(attrs types.Attributes, path string) (string, error) {
	if alt, found := attrs.GetAsString(types.AttrImageAlt); found {
		return alt, nil
	}
	u, err := url.Parse(path)
	if err != nil {
		return "", errors.Wrap(err, "unable to render image")
	}
	// return base path without its extension
	result := strings.TrimSuffix(filepath.Base(u.Path), filepath.Ext(u.Path))
	// replace separators
	result = strings.ReplaceAll(result, "-", " ")
	result = strings.ReplaceAll(result, "_", " ")
	return result, nil
}
