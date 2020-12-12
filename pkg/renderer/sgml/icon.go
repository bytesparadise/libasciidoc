package sgml

import (
	fmt "fmt"
	"path"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderInlineIcon(ctx *renderer.Context, icon types.Icon) (string, error) {
	result := &strings.Builder{}

	iconStr, err := r.renderIcon(ctx, types.Icon{
		Class:      icon.Class,
		Attributes: icon.Attributes,
	}, false)
	if err != nil {
		return "", err
	}
	err = r.inlineIcon.Execute(result, struct {
		Class  string
		Role   string
		Link   string
		Window string
		ID     string
		Icon   string
	}{
		Class:  icon.Class,
		Icon:   iconStr,
		ID:     r.renderElementID(icon.Attributes),
		Link:   icon.Attributes.GetAsStringWithDefault(types.AttrInlineLink, ""),
		Window: icon.Attributes.GetAsStringWithDefault(types.AttrImageWindow, ""),
		Role:   icon.Attributes.GetAsStringWithDefault(types.AttrRoles, ""),
	})

	if err != nil {
		return "", errors.Wrap(err, "unable to render inline image")
	}
	return result.String(), nil
}

var defaultIconClasses = map[string]string{
	types.AttrCautionCaption:   "Caution",
	types.AttrImportantCaption: "Important",
	types.AttrNoteCaption:      "Note",
	types.AttrTipCaption:       "Tip",
	types.AttrWarningCaption:   "Warning",
}

func (r *sgmlRenderer) renderIcon(ctx *renderer.Context, icon types.Icon, admonition bool) (string, error) {
	icons := ctx.Attributes.GetAsStringWithDefault("icons", "text")
	var template *textTemplate
	font := false
	switch icons {
	case "font":
		font = true
		template = r.iconFont
	case "text":
		template = r.iconText
	case "image", "":
		template = r.iconImage
	default:
		return "", fmt.Errorf("unsupported icon type %s", icons)
	}
	title := ""
	alt := icon.Class

	// TODO: This is rather inconsistent, and done for CSS compatibility.  We should
	// expand the templates, and eliminate this code in the future.
	if admonition {
		// Admonition uses title on block instead of the icon, and the alt text is
		// taken from the caption.  However, in admonitions using the font, the alt
		// is used as the title element instead.  Go figure.
		alt = ctx.Attributes.GetAsStringWithDefault(icon.Class+"-caption", defaultIconClasses[icon.Class+"-caption"])
		alt = icon.Attributes.GetAsStringWithDefault(types.AttrCaption, alt)
		if font {
			title = alt
			alt = ""
		}
	} else {
		// Inline icons use the alt attribute, and may optionally carry a title.
		// The alt is the icon class name unless overridden.  They don't use the caption at all.
		alt = icon.Attributes.GetAsStringWithDefault(types.AttrImageAlt, alt)
		title = icon.Attributes.GetAsStringWithDefault(types.AttrTitle, "")
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
		Alt:        alt,
		Title:      title,
		Width:      icon.Attributes.GetAsStringWithDefault(types.AttrWidth, ""),
		Height:     icon.Attributes.GetAsStringWithDefault(types.AttrHeight, ""),
		Size:       icon.Attributes.GetAsStringWithDefault(types.AttrIconSize, ""),
		Rotate:     icon.Attributes.GetAsStringWithDefault(types.AttrIconRotate, ""),
		Flip:       icon.Attributes.GetAsStringWithDefault(types.AttrIconFlip, ""),
		Link:       icon.Attributes.GetAsStringWithDefault(types.AttrInlineLink, ""),
		Window:     icon.Attributes.GetAsStringWithDefault(types.AttrImageWindow, ""),
		Path:       renderIconPath(ctx, icon.Class),
		Admonition: admonition,
	})
	return string(s.String()), err
}

func renderIconPath(ctx *renderer.Context, name string) string {
	// Icon files by default are in {imagesdir}/icons, where {imagesdir} defaults to "./images"
	dir := ctx.Attributes.GetAsStringWithDefault("iconsdir",
		path.Join(ctx.Attributes.GetAsStringWithDefault("imagesdir", "./images"), "icons"))
	// TODO: perform attribute substitutions here!
	ext := ctx.Attributes.GetAsStringWithDefault("icontype", "png")

	return path.Join(dir, name+"."+ext)
}
