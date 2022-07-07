package sgml

import (
	fmt "fmt"
	"path"
	"strings"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderInlineIcon(ctx *context, icon *types.Icon) (string, error) {
	iconStr, err := r.renderIcon(ctx, types.Icon{
		Class:      icon.Class,
		Attributes: icon.Attributes,
	}, false)
	if err != nil {
		return "", err
	}
	return r.execute(r.inlineIcon, struct {
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
}

var defaultIconClasses = map[string]string{
	types.AttrCautionCaption:   "Caution",
	types.AttrImportantCaption: "Important",
	types.AttrNoteCaption:      "Note",
	types.AttrTipCaption:       "Tip",
	types.AttrWarningCaption:   "Warning",
}

func (r *sgmlRenderer) renderIcon(ctx *context, icon types.Icon, admonition bool) (string, error) {
	icons := ctx.attributes.GetAsStringWithDefault("icons", "text")
	var tmpl *texttemplate.Template
	var err error
	font := false
	switch icons {
	case "font":
		font = true
		tmpl, err = r.iconFont()
	case "text":
		tmpl, err = r.iconText()
	case "image", "":
		tmpl, err = r.iconImage()
	default:
		return "", fmt.Errorf("unsupported icon type %s", icons)
	}
	if err != nil {
		return "", errors.Wrap(err, "unable to load icon template")
	}
	title := ""
	alt := icon.Class

	// TODO: This is rather inconsistent, and done for CSS compatibility.  We should
	// expand the templates, and eliminate this code in the future.
	if admonition {
		// Admonition uses title on block instead of the icon, and the alt text is
		// taken from the caption.  However, in admonitions using the font, the alt
		// is used as the title element instead.  Go figure.
		alt = ctx.attributes.GetAsStringWithDefault(icon.Class+"-caption", defaultIconClasses[icon.Class+"-caption"])
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
	if err := tmpl.Execute(s, struct {
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
		Src        string
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
		Src:        renderIconPath(ctx, icon.Class),
		Admonition: admonition,
	}); err != nil {
		return "", errors.Wrap(err, "unable to render icon")
	}
	return string(s.String()), nil
}

func renderIconPath(ctx *context, name string) string {
	// Icon files by default are in {imagesdir}/icons, where {imagesdir} defaults to "./images"
	dir := ctx.attributes.GetAsStringWithDefault("iconsdir",
		path.Join(ctx.attributes.GetAsStringWithDefault(types.AttrImagesDir, "./images"), "icons"))
	// TODO: perform attribute substitutions here!
	ext := ctx.attributes.GetAsStringWithDefault("icontype", "png")

	return path.Join(dir, name+"."+ext)
}
