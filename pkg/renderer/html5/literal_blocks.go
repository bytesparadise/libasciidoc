package html5

import (
	"bytes"
	"math"
	"strings"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var literalBlockTmpl texttemplate.Template

// initializes the templates
func init() {
	literalBlockTmpl = newTextTemplate("literal block", `{{ $ctx := .Context }}{{ with .Data }}<div {{ if .ID }}id="{{ .ID }}" {{ end }}class="literalblock">
{{ if .Title }}<div class="title">{{ .Title }}</div>
{{ end }}<div class="content">
<pre>{{ $lines := .Lines }}{{ range $index, $line := $lines}}{{ $line }}{{ includeNewline $ctx $index $lines }}{{ end }}</pre>
</div>
</div>{{ end }}`, texttemplate.FuncMap{
		"includeNewline": includeNewline,
	})
}

func renderLiteralBlock(ctx *renderer.Context, b types.LiteralBlock) ([]byte, error) {
	log.Debugf("rendering delimited block with content: %s", b.Lines)
	var lines []string
	if b.Attributes.GetAsString(types.AttrLiteralBlockType) == types.LiteralBlockWithSpacesOnFirstLine {
		// remove as many spaces as needed on each line
		lines = make([]string, len(b.Lines))
		spaceCount := float64(0)
		// first pass to detemine the minimum number of spaces to remove
		for i, line := range b.Lines {
			l := strings.TrimLeft(line, " ")
			if i == 0 {
				spaceCount = float64(len(line) - len(l))
			} else {
				spaceCount = math.Min(spaceCount, float64(len(line)-len(l)))
			}
		}
		// then remove the same number of spaces on each line
		spaces := strings.Repeat(" ", int(spaceCount))
		for i, line := range b.Lines {
			lines[i] = strings.TrimPrefix(line, spaces)
		}
	} else {
		// just use the lines as-is
		lines = b.Lines
	}
	result := bytes.NewBuffer(nil)
	err := literalBlockTmpl.Execute(result, ContextualPipeline{
		Context: ctx,
		Data: struct {
			ID    string
			Title string
			Lines []string
		}{
			ID:    b.Attributes.GetAsString(types.AttrID),
			Title: b.Attributes.GetAsString(types.AttrTitle),
			Lines: lines,
		}})
	if err != nil {
		return nil, errors.Wrapf(err, "unable to render delimited block")
	}
	return result.Bytes(), nil
}
