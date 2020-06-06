package html5

import (
	"bytes"
	"strconv"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

var blockImageTmpl texttemplate.Template
var inlineImageTmpl texttemplate.Template

// initializes the templates
func init() {
	blockImageTmpl = newTextTemplate("block image", `<div{{ if .ID }} id="{{ .ID }}"{{ end }} class="imageblock{{ if .Role }} {{ .Role }}{{ end }}">
<div class="content">
{{ if ne .Href "" }}<a class="image" href="{{ .Href }}">{{ end }}<img src="{{ .Path }}" alt="{{ .Alt }}"{{ if .Width }} width="{{ .Width }}"{{ end }}{{ if .Height }} height="{{ .Height }}"{{ end }}>{{ if ne .Href "" }}</a>{{ end }}
</div>{{ if .Title }}
<div class="title">{{ escape .Title }}</div>
{{ else }}
{{ end }}</div>`,
		texttemplate.FuncMap{
			"escape": EscapeString,
		})
	inlineImageTmpl = newTextTemplate("inline image", `<span class="image{{ if .Role }} {{ .Role }}{{ end }}"><img src="{{ .Path }}" alt="{{ .Alt }}"{{ if .Width }} width="{{ .Width }}"{{ end }}{{ if .Height }} height="{{ .Height }}"{{ end }}{{ if .Title }} title="{{ escape .Title }}"{{ end }}></span>`,
		texttemplate.FuncMap{
			"escape": EscapeString,
		})
}

func renderImageBlock(ctx renderer.Context, img types.ImageBlock) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	title := ""
	if t, found := img.Attributes.GetAsString(types.AttrTitle); found {
		title = "Figure " + strconv.Itoa(ctx.GetAndIncrementImageCounter()) + ". " + EscapeString(t)
	}
	err := blockImageTmpl.Execute(result, struct {
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

func renderInlineImage(img types.InlineImage) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	err := inlineImageTmpl.Execute(result, struct {
		Role   string
		Title  string
		Href   string
		Alt    string
		Width  string
		Height string
		Path   string
	}{
		Title:  renderElementTitle(img.Attributes),
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
