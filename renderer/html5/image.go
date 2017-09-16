package html5

import (
	"bytes"
	"html/template"

	asciidoc "github.com/bytesparadise/libasciidoc/context"
	"github.com/bytesparadise/libasciidoc/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var blockImageTmpl *template.Template
var inlineImageTmpl *template.Template

// initializes the templates
func init() {
	blockImageTmpl = newHTMLTemplate("block image", `<div{{if .ID }} id="{{.ID.Value}}"{{ end }} class="imageblock">
<div class="content">
{{if .Link}}<a class="image" href="{{.Link.Path}}">{{end}}<img src="{{.Macro.Path}}" alt="{{.Macro.Alt}}"{{if .Macro.Width}} width="{{.Macro.Width}}"{{end}}{{if .Macro.Height}} height="{{.Macro.Height}}"{{end}}>{{if .Link}}</a>{{end}}
</div>{{if .Title}}
<div class="title">{{.Title.Value}}</div>
{{else}}
{{end}}</div>`)
	inlineImageTmpl = newHTMLTemplate("inline image", `<span class="image"><img src="{{.Macro.Path}}" alt="{{.Macro.Alt}}"{{if .Macro.Width}} width="{{.Macro.Width}}"{{end}}{{if .Macro.Height}} height="{{.Macro.Height}}"{{end}}></span>`)
}

func renderBlockImage(ctx asciidoc.Context, img types.BlockImage) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	err := blockImageTmpl.Execute(result, img)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render block image")
	}
	log.Debugf("rendered block image: %s", result.Bytes())
	return result.Bytes(), nil
}

func renderInlineImage(ctx asciidoc.Context, img types.InlineImage) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	err := inlineImageTmpl.Execute(result, img)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render inline image")
	}
	log.Debugf("rendered inline image: %s", result.Bytes())
	return result.Bytes(), nil
}
