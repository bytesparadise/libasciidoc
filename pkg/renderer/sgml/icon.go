package sgml

import (
	"bytes"
	fmt "fmt"
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	"path"
	"strings"
)

func (r *sgmlRenderer) renderInlineIcon(ctx *renderer.Context, icon types.InlineIcon) ([]byte, error) {
	result := &bytes.Buffer{}

	iconStr, err := r.renderIcon(ctx, types.Icon{
		Class:      icon.Class,
		Attributes: icon.Attributes,
	}, false)
	if err != nil {
		return nil, err
	}
	err = r.inlineIcon.Execute(result, struct {
		Class  string
		Role   string
		Link   string
		Window string
		ID     string
		Icon   sanitized
	}{
		Class:  icon.Class,
		Icon:   iconStr,
		ID:     icon.Attributes.GetAsStringWithDefault(types.AttrID, ""),
		Link:   icon.Attributes.GetAsStringWithDefault(types.AttrInlineLink, ""),
		Window: icon.Attributes.GetAsStringWithDefault(types.AttrImageWindow, ""),
		Role:   icon.Attributes.GetAsStringWithDefault(types.AttrRole, ""),
	})

	if err != nil {
		return nil, errors.Wrapf(err, "unable to render inline image")
	}
	return result.Bytes(), nil
}

func (r *sgmlRenderer) renderIcon(ctx *renderer.Context, icon types.Icon, admonition bool) (sanitized, error) {
	icons := ctx.Attributes.GetAsStringWithDefault("icons", "text")
	var template *textTemplate
	switch icons {
	case "font":
		template = r.iconFont
	case "text":
		template = r.iconText
	case "image", "":
		template = r.iconImage
	default:
		return "", fmt.Errorf("unsupported icon type %s", icons)
	}
	title := ""
	if admonition {
		title = strings.Title(icon.Class)
	}
	s := &strings.Builder{}
	err := template.Execute(s, struct {
		Class      string
		Alt        string
		Title      string
		Link       string
		Window     string
		Size       string
		Rotate     string
		Flip       string
		Width      string
		Height     string
		Path       string
		Admonition bool
	}{
		Class:      icon.Class,
		Alt:        icon.Attributes.GetAsStringWithDefault(types.AttrImageAlt, strings.Title(icon.Class)),
		Title:      icon.Attributes.GetAsStringWithDefault(types.AttrImageTitle, title),
		Width:      icon.Attributes.GetAsStringWithDefault(types.AttrImageWidth, ""),
		Height:     icon.Attributes.GetAsStringWithDefault(types.AttrImageHeight, ""),
		Size:       icon.Attributes.GetAsStringWithDefault(types.AttrIconSize, ""),
		Rotate:     icon.Attributes.GetAsStringWithDefault(types.AttrIconRotate, ""),
		Flip:       icon.Attributes.GetAsStringWithDefault(types.AttrIconFlip, ""),
		Link:       icon.Attributes.GetAsStringWithDefault(types.AttrInlineLink, ""),
		Window:     icon.Attributes.GetAsStringWithDefault(types.AttrImageWindow, ""),
		Path:       renderIconPath(ctx, icon.Class),
		Admonition: admonition,
	})
	return sanitized(s.String()), err
}

func renderIconPath(ctx *renderer.Context, name string) string {
	// Icon files by default are in {imagesdir}/icons, where {imagesdir} defaults to "./images"
	// TODO: This should use the location resolver, but we need to be able to supply a different
	dir := ctx.Attributes.GetAsStringWithDefault("iconsdir",
		path.Join(ctx.Attributes.GetAsStringWithDefault("imagesdir", "./images"), "icons"))
	// TODO: perform attribute substitutions here!
	ext := ctx.Attributes.GetAsStringWithDefault("icontype", "png")

	return path.Join(dir, name+"."+ext)
}
