package sgml

import (
	text "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
)

// These type aliases provide local names for names in other packages,
// thereby helping minimize collisions based on conflicting package
// names.  It also reduces the imports we have to use everywhere else.

type Context = renderer.Context

type textTemplate = text.Template
type funcMap = text.FuncMap
