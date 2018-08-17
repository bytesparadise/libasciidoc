package html5

import (
	"bytes"
	"fmt"
	"net/url"
	"path/filepath"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var blockImageTmpl texttemplate.Template
var inlineImageTmpl texttemplate.Template

// initializes the templates
func init() {
	blockImageTmpl = newTextTemplate("block image", `<div{{ if ne .ID "" }} id="{{ .ID }}"{{ end }} class="imageblock">
<div class="content">
{{ if ne .Href "" }}<a class="image" href="{{ .Href }}">{{ end }}<img src="{{ .Path }}" alt="{{ .Alt }}"{{ if .Width }} width="{{ .Width }}"{{ end }}{{ if .Height }} height="{{ .Height }}"{{ end }}>{{ if ne .Href "" }}</a>{{ end }}
</div>{{ if ne .Title "" }}
<div class="doctitle">{{ .Title }}</div>
{{ else }}
{{ end }}</div>`)
	inlineImageTmpl = newTextTemplate("inline image", `<span class="image"><img src="{{ .Path }}" alt="{{ .Alt }}"{{ if .Width }} width="{{ .Width }}"{{ end }}{{ if .Height }} height="{{ .Height }}"{{ end }}></span>`)
}

func renderBlockImage(ctx *renderer.Context, img types.BlockImage) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	var id, title, href string
	if i, ok := img.Attributes[types.AttrID].(string); ok {
		id = i
	}
	if t, ok := img.Attributes[types.AttrTitle].(string); ok {
		title = t
	}
	if l, ok := img.Attributes[types.AttrLink].(string); ok {
		href = l
	}
	err := blockImageTmpl.Execute(result, struct {
		ID     string
		Title  string
		Href   string
		Path   string
		Alt    string
		Width  string
		Height string
	}{
		ID:     id,
		Title:  title,
		Href:   href,
		Path:   getImageHref(ctx, img.Macro.Path),
		Alt:    img.Macro.Alt(),
		Width:  img.Macro.Width(),
		Height: img.Macro.Height(),
	})

	if err != nil {
		return nil, errors.Wrapf(err, "unable to render block image")
	}
	log.Debugf("rendered block image: %s", result.Bytes())
	return result.Bytes(), nil
}

func renderInlineImage(ctx *renderer.Context, img types.InlineImage) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	err := inlineImageTmpl.Execute(result, struct {
		ID     string
		Title  string
		Href   string
		Path   string
		Alt    string
		Width  string
		Height string
	}{
		Path:   getImageHref(ctx, img.Macro.Path),
		Alt:    img.Macro.Alt(),
		Width:  img.Macro.Width(),
		Height: img.Macro.Height(),
	})

	if err != nil {
		return nil, errors.Wrapf(err, "unable to render inline image")
	}
	log.Debugf("rendered inline image: %s", result.Bytes())
	return result.Bytes(), nil
}

// getImageLink returns the `href` value for the image. If the given location `l` is relative,
// then the context's `imagesdir` attribute is used (if it is set). If the location `l` is
// absolute, then it is returned as-is
//
func getImageHref(ctx *renderer.Context, l string) string {
	if _, err := url.ParseRequestURI(l); err == nil {
		// location is a valid URL, so return it as-is
		log.Debugf("location '%s' is an URL", l)
		return l
	}
	if filepath.IsAbs(l) {
		log.Debugf("location '%s' is an absolute path", l)
		return l
	}
	// use `imagesdir` attribute if it is set
	if imagesdir := ctx.GetImagesDir(); imagesdir != "" {
		log.Debugf("location '%s' is a relative path, adding '%s' as a prefix", l, imagesdir)
		return fmt.Sprintf("%s/%s", imagesdir, l)
	}
	// default
	log.Debugf("location '%s' is a relative path, but 'imagesdir' attribute was not set", l)
	return l

}
