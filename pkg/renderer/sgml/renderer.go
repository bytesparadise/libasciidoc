package sgml

import (
	"fmt"
	"io"
	"strings"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/davecgh/go-spew/spew"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// Renderer implements the backend render interface by using sgml.
type Renderer interface {

	// Render renders a document to the given output stream.
	Render(ctx *renderer.Context, doc *types.Document, output io.Writer) (types.Metadata, error)

	// SetFunction sets the named function.
	SetFunction(name string, fn interface{})

	// Templates returns the Templates used by this Renderer.
	// It cannot be altered on a given Renderer, since the old
	// templates may have already been parsed.
	Templates() Templates
}

// NewRenderer returns a new renderer
func NewRenderer(t Templates) Renderer {
	r := &sgmlRenderer{
		templates: t,
	}
	// Establish some default function handlers.
	r.functions = funcMap{
		"escape":              EscapeString,
		"trimRight":           trimRight,
		"trimLeft":            trimLeft,
		"trim":                trimBoth,
		"specialCharacter":    specialCharacter,
		"predefinedAttribute": predefinedAttribute,
		"halign":              halign,
		"valign":              valign,
	}
	return r
}

func trimLeft(s string) string {
	return strings.TrimLeft(s, " ")
}

func trimRight(s string) string {
	return strings.TrimRight(s, " ")
}

func trimBoth(s string) string {
	return strings.Trim(s, " ")
}

var specialCharacters = map[string]string{
	">": "&gt;",
	"<": "&lt;",
	"&": "&amp;",
}

func specialCharacter(c string) string {
	return specialCharacters[c]
}

var predefinedAttributes = map[string]string{
	"sp":             " ",
	"blank":          "",
	"empty":          "",
	"nbsp":           "\u00a0",
	"zwsp":           "\u200b",
	"wj":             "\u2060",
	"apos":           "&#39;",
	"quot":           "&#34;",
	"lsquo":          "\u2018",
	"rsquo":          "\u2019",
	"ldquo":          "\u201c",
	"rdquo":          "\u201d",
	"deg":            "\u00b0",
	"plus":           "&#43;",
	"brvbar":         "\u00a6",
	"vbar":           "|", // TODO: maybe convert this because of tables?
	"amp":            "&amp;",
	"lt":             "<",
	"gt":             ">",
	"startsb":        "[",
	"endsb":          "]",
	"caret":          "^",
	"asterisk":       "*",
	"tilde":          "~",
	"backslash":      `\`,
	"backtick":       "`",
	"two-colons":     "::",
	"two-semicolons": ";",
	"cpp":            "C++",
}

func predefinedAttribute(a string) string {
	// log.Debugf("predefined attribute '%s': '%s", a, predefinedAttributes[a])
	return predefinedAttributes[a]
}

func halign(v types.HAlign) string {
	switch v {
	case types.HAlignLeft:
		return "halign-left"
	case types.HAlignCenter:
		return "halign-center"
	case types.HAlignRight:
		return "halign-right"
	default:
		return string(v)
	}
}

func valign(v types.VAlign) string {
	switch v {
	case types.VAlignTop:
		return "valign-top"
	case types.VAlignMiddle:
		return "valign-middle"
	case types.VAlignBottom:
		return "valign-bottom"
	default:
		return string(v)
	}
}

func (r *sgmlRenderer) SetFunction(name string, fn interface{}) {
	r.functions[name] = fn
}

// Templates returns the Templates being used by this renderer.
// A copy is made, as we cannot change the original Templates
// due to it already being used.
func (r *sgmlRenderer) Templates() Templates {
	return r.templates
}

func (r *sgmlRenderer) newTemplate(name string, tmpl string, err error) (*textTemplate, error) {
	// NB: if the data is missing below, it will be an empty string.
	if err != nil {
		return nil, err
	}
	if len(tmpl) == 0 {
		return nil, fmt.Errorf("empty template for '%s'", name)
	}
	t := texttemplate.New(name)
	t.Funcs(r.functions)
	if t, err = t.Parse(tmpl); err != nil {
		log.Errorf("failed to initialize the '%s' template: %v", name, err)
		return nil, err
	}
	return t, nil
}

// Render renders the given document in HTML and writes the result in the given `writer`
func (r *sgmlRenderer) Render(ctx *renderer.Context, doc *types.Document, output io.Writer) (types.Metadata, error) {
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("rendering document of type '%s'\n%s", ctx.Attributes.GetAsStringWithDefault(types.AttrDocType, "article"), spew.Sdump(ctx.Attributes))
	// }

	// metadata to be returned to the caller
	var metadata types.Metadata
	// arguably this should be a time.Time for use in Go
	metadata.LastUpdated = ctx.Config.LastUpdated.Format(configuration.LastUpdatedFormat)
	err := r.prepareTemplates()
	if err != nil {
		return metadata, err
	}

	renderedTitle, exists, err := r.renderDocumentTitle(ctx, doc)
	if err != nil {
		return metadata, errors.Wrapf(err, "unable to render full document")
	}
	metadata.Title = string(renderedTitle) // retain an empty value if no title was defined in the document
	if !exists {
		renderedTitle = DefaultTitle
	}
	// process attribute declaration in the header
	if header := doc.Header(); header != nil {
		for _, e := range header.Elements {
			switch e := e.(type) {
			case *types.AttributeDeclaration:
				ctx.Attributes[e.Name] = e.Value
			case *types.AttributeReset:
				delete(ctx.Attributes, e.Name)
			}
		}
	}
	if ctx.Attributes.Has(types.AttrSectionNumbering) || ctx.Attributes.Has(types.AttrNumbered) {
		var err error
		if ctx.SectionNumbering, err = renderer.NewSectionNumbers(doc); err != nil {
			return metadata, errors.Wrapf(err, "unable to render full document")
		}
	} else {
		log.Debug("section numbering is not enabled")
	}

	// needs to be set before rendering the content elements
	if err := r.prerenderTableOfContents(ctx, doc.TableOfContents); err != nil {
		return metadata, errors.Wrapf(err, "unable to render full document")
	}
	metadata.TableOfContents = doc.TableOfContents
	renderedHeader, renderedContent, err := r.splitAndRender(ctx, doc)
	if err != nil {
		return metadata, errors.Wrapf(err, "unable to render full document")
	}
	roles, err := r.renderDocumentRoles(ctx, doc)
	if err != nil {
		return metadata, errors.Wrap(err, "unable to render fenced block content")
	}
	if ctx.Config.WrapInHTMLBodyElement {
		log.Debugf("Rendering full document...")
		err = r.article.Execute(output, struct {
			Doctype               string
			Generator             string
			Description           string
			Title                 string
			Authors               string
			Header                string
			ID                    string
			Roles                 string
			Content               string
			RevNumber             string
			LastUpdated           string
			CSS                   string
			IncludeHTMLBodyHeader bool
			IncludeHTMLBodyFooter bool
		}{
			Doctype:               ctx.Attributes.GetAsStringWithDefault(types.AttrDocType, "article"),
			Generator:             "libasciidoc", // TODO: externalize this value and include the lib version ?
			Description:           ctx.Attributes.GetAsStringWithDefault(types.AttrDescription, ""),
			Title:                 renderedTitle,
			Authors:               r.renderAuthors(ctx, doc),
			Header:                renderedHeader,
			Roles:                 roles,
			ID:                    r.renderDocumentID(doc),
			Content:               string(renderedContent), // nolint:gosec
			RevNumber:             ctx.Attributes.GetAsStringWithDefault("revnumber", ""),
			LastUpdated:           ctx.Config.LastUpdated.Format(configuration.LastUpdatedFormat),
			CSS:                   ctx.Config.CSS,
			IncludeHTMLBodyHeader: !ctx.Attributes.Has(types.AttrNoHeader),
			IncludeHTMLBodyFooter: !ctx.Attributes.Has(types.AttrNoFooter),
		})
		if err != nil {
			return metadata, errors.Wrapf(err, "unable to render full document")
		}
	} else {
		log.Debugf("Rendering document body...")
		_, err = output.Write([]byte(renderedContent))
		if err != nil {
			return metadata, errors.Wrapf(err, "unable to render full document")
		}
	}
	return metadata, err
}

// splitAndRender the document with the header elements on one side
// and all other elements (table of contents, with preamble, content) on the other side,
// then renders the header and other elements
func (r *sgmlRenderer) splitAndRender(ctx *renderer.Context, doc *types.Document) (string, string, error) {
	switch ctx.Attributes.GetAsStringWithDefault(types.AttrDocType, "article") {
	case "manpage":
		return r.splitAndRenderForManpage(ctx, doc)
	default:
		return r.splitAndRenderForArticle(ctx, doc)
	}
}

// splits the document with the title of the section 0 (if available) on one side
// and all other elements (table of contents, with preamble, content) on the other side
func (r *sgmlRenderer) splitAndRenderForArticle(ctx *renderer.Context, doc *types.Document) (string, string, error) {
	log.Debugf("rendering article (within HTML/Body: %t)", ctx.Config.WrapInHTMLBodyElement)

	renderedHeader, err := r.renderDocumentHeader(ctx, doc.Header())
	if err != nil {
		return "", "", err
	}
	renderedContent, err := r.renderDocumentBody(ctx, doc.BodyElements(), doc.TableOfContents, doc.Footnotes)
	if err != nil {
		return "", "", err
	}
	return renderedHeader, renderedContent, nil
}

// splits the document with the header elements on one side
// and the other elements (table of contents, with preamble, content) on the other side
func (r *sgmlRenderer) splitAndRenderForManpage(ctx *renderer.Context, doc *types.Document) (string, string, error) {
	log.Debugf("rendering manpage (within HTML/Body: %t)", ctx.Config.WrapInHTMLBodyElement)
	elements := doc.BodyElements()
	nameSection := elements[0].(*types.Section) // TODO: enforce
	if ctx.Config.WrapInHTMLBodyElement {
		renderedHeader, err := r.renderManpageHeader(ctx, doc.Header(), nameSection)
		if err != nil {
			return "", "", err
		}
		renderedContent, err := r.renderDocumentBody(ctx, elements[1:], doc.TableOfContents, doc.Footnotes)
		if err != nil {
			return "", "", err
		}
		return renderedHeader, renderedContent, nil
	}
	// in that case, we still want to display the name section
	renderedHeader, err := r.renderManpageHeader(ctx, nil, nameSection)
	if err != nil {
		return "", "", err
	}
	renderedContent, err := r.renderDocumentBody(ctx, elements[1:], doc.TableOfContents, doc.Footnotes)
	if err != nil {
		return "", "", err
	}
	result := &strings.Builder{}
	result.WriteString(renderedHeader)
	result.WriteString(renderedContent)
	return "", result.String(), nil
}

func (r *sgmlRenderer) renderDocumentRoles(ctx *renderer.Context, doc *types.Document) (string, error) {
	if header := doc.Header(); header != nil {
		return r.renderElementRoles(ctx, header.Attributes)
	}
	return "", nil
}

func (r *sgmlRenderer) renderDocumentID(doc *types.Document) string {
	if header := doc.Header(); header != nil {
		// if header.Attributes.Has(types.AttrCustomID) {
		// We only want to emit a document body ID if one was explicitly set
		return r.renderElementID(header.Attributes)
		// }
	}
	return ""
}

func (r *sgmlRenderer) renderAuthors(_ *renderer.Context, doc *types.Document) string { // TODO: pass header instead of doc
	if header := doc.Header(); header != nil {
		if a := header.Authors(); a != nil {
			authorStrs := make([]string, len(a))
			for i, author := range a {
				authorStrs[i] = author.FullName()
			}
			return strings.Join(authorStrs, "; ")

		}
	}
	return ""
}

// DefaultTitle the default title to render when the document has none
const DefaultTitle = "Untitled"

func (r *sgmlRenderer) renderDocumentTitle(ctx *renderer.Context, doc *types.Document) (string, bool, error) {
	if header := doc.Header(); header != nil {
		// TODO: This feels wrong.  The title should not need markup.
		title, err := r.renderPlainText(ctx, header.Title)
		if err != nil {
			return "", true, errors.Wrap(err, "unable to render document title")
		}
		return string(title), true, nil
	}
	return "", false, nil
}

func (r *sgmlRenderer) renderDocumentHeader(ctx *renderer.Context, header *types.DocumentHeader) (string, error) {
	if header == nil {
		return "", nil
	}
	renderedHeader, err := r.renderDocumentHeaderTitle(ctx, header)
	if err != nil {
		return "", err
	}
	documentDetails, err := r.renderDocumentDetails(ctx)
	if err != nil {
		return "", err
	}
	output := &strings.Builder{}
	err = r.articleHeader.Execute(output, struct {
		Header  string
		Details *string // TODO: convert to string (no need to be a pointer)
	}{
		Header:  renderedHeader,
		Details: documentDetails,
	})
	if err != nil {
		return "", err
	}
	return output.String(), nil
}

func (r *sgmlRenderer) renderDocumentHeaderTitle(ctx *renderer.Context, header *types.DocumentHeader) (string, error) {
	if header == nil {
		log.Debug("no header to render")
		return "", nil
	}
	return r.renderInlineElements(ctx, header.Title)
}

func (r *sgmlRenderer) renderManpageHeader(ctx *renderer.Context, header *types.DocumentHeader, nameSection *types.Section) (string, error) {
	renderedHeader, err := r.renderDocumentHeaderTitle(ctx, header)
	if err != nil {
		return "", err
	}
	renderedName, err := r.renderInlineElements(ctx, nameSection.Title)
	if err != nil {
		return "", err
	}
	description := nameSection.Elements[0].(*types.Paragraph) // TODO: type check
	if description.Attributes == nil {
		description.Attributes = types.Attributes{}
	}
	description.Attributes.Set(types.AttrStyle, "manpage")
	renderedContent, err := r.renderParagraph(ctx, description)
	if err != nil {
		return "", err
	}
	output := &strings.Builder{}
	err = r.manpageHeader.Execute(output, struct {
		Header    string
		Name      string
		Content   string
		IncludeH1 bool
	}{
		Header:    renderedHeader,
		Name:      renderedName,
		Content:   string(renderedContent), // nolint:gosec
		IncludeH1: len(renderedHeader) > 0,
	})
	if err != nil {
		return "", err
	}
	return output.String(), nil
}

// renderDocumentBody renders all document elements, including the footnotes,
// but not the HEAD and BODY containers
func (r *sgmlRenderer) renderDocumentBody(ctx *renderer.Context, source []interface{}, toc *types.TableOfContents, footnotes []*types.Footnote) (string, error) {
	var elements []interface{}
	if placement, found := ctx.Attributes[types.AttrTableOfContents]; found && toc != nil {
		switch placement {
		case "preamble":
			log.Debug("inserting ToC in preamble")
			// look-up the preamble in the source elements
			if p := lookupPreamble(source); p != nil {
				// insert ToC at the end of the preamble
				p.TableOfContents = toc
				elements = source
			} else {
				elements = source
			}
		case "", nil:
			log.Debug("inserting ToC as first element")
			elements = make([]interface{}, len(source)+1)
			elements[0] = toc
			copy(elements[1:], source)
		default:
			return "", fmt.Errorf("unsupported table of contents placement: '%s'", placement)
		}
	} else {
		log.Debug("not inserting ToC")
		elements = source
	}

	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("rendering elements:\n%s", spew.Sdump(elements))
	}
	buff := &strings.Builder{}
	renderedElements, err := r.renderElements(ctx, elements)
	if err != nil {
		return "", errors.Wrap(err, "failed to render document elements")
	}
	buff.WriteString(renderedElements)
	renderedFootnotes, err := r.renderFootnotes(ctx, footnotes)
	if err != nil {
		return "", errors.Wrap(err, "failed to render document elements")
	}
	buff.WriteString(renderedFootnotes)
	return buff.String(), nil
}

func lookupPreamble(elements []interface{}) *types.Preamble {
	for _, e := range elements {
		if p, ok := e.(*types.Preamble); ok {
			return p
		}
	}
	return nil
}
