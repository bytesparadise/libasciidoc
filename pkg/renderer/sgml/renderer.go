package sgml

import (
	"fmt"
	"io"
	"strings"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// Renderer implements the backend render interface by using sgml.
type Renderer interface {

	// Render renders a document to the given output stream.
	Render(ctx *renderer.Context, doc types.Document, output io.Writer) (types.Metadata, error)

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
	log.Debugf("predefined attribute '%s': '%s", a, predefinedAttributes[a])
	return predefinedAttributes[a]
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
func (r *sgmlRenderer) Render(ctx *renderer.Context, doc types.Document, output io.Writer) (types.Metadata, error) {

	var md types.Metadata
	err := r.prepareTemplates()
	if err != nil {
		return md, err
	}
	if doc.Attributes.Has(types.AttrUnicode) {
		ctx.UseUnicode = true
	}
	renderedTitle, err := r.renderDocumentTitle(ctx, doc)
	if err != nil {
		return md, errors.Wrapf(err, "unable to render full document")
	}
	// needs to be set before rendering the content elements
	ctx.TableOfContents, err = r.newTableOfContents(ctx, doc)
	if err != nil {
		return md, errors.Wrapf(err, "unable to render full document")
	}
	renderedHeader, renderedContent, err := r.splitAndRender(ctx, doc)
	if err != nil {
		return md, errors.Wrapf(err, "unable to render full document")
	}
	roles, err := r.renderDocumentRoles(ctx, doc)
	if err != nil {
		return md, errors.Wrap(err, "unable to render fenced block content")
	}
	if ctx.Config.IncludeHeaderFooter {
		log.Debugf("Rendering full document...")
		err = r.article.Execute(output, struct {
			Generator     string
			Doctype       string
			Title         string
			Authors       string
			Header        string
			Role          string
			ID            string
			Roles         string
			Content       string
			RevNumber     string
			LastUpdated   string
			CSS           string
			IncludeHeader bool
			IncludeFooter bool
		}{
			Generator:     "libasciidoc", // TODO: externalize this value and include the lib version ?
			Doctype:       doc.Attributes.GetAsStringWithDefault(types.AttrDocType, "article"),
			Title:         renderedTitle,
			Authors:       r.renderAuthors(doc),
			Header:        renderedHeader,
			Roles:         roles,
			ID:            r.renderDocumentID(doc),
			Content:       string(renderedContent), //nolint: gosec
			RevNumber:     doc.Attributes.GetAsStringWithDefault("revnumber", ""),
			LastUpdated:   ctx.Config.LastUpdated.Format(configuration.LastUpdatedFormat),
			CSS:           ctx.Config.CSS,
			IncludeHeader: !doc.Attributes.Has(types.AttrNoHeader),
			IncludeFooter: !doc.Attributes.Has(types.AttrNoFooter),
		})
		if err != nil {
			return md, errors.Wrapf(err, "unable to render full document")
		}
	} else {
		_, err = output.Write([]byte(renderedContent))
		if err != nil {
			return md, errors.Wrapf(err, "unable to render full document")
		}
	}
	// generate the metadata to be returned to the caller
	md.Title = string(renderedTitle)
	// arguably this should be a time.Time for use in Go
	md.LastUpdated = ctx.Config.LastUpdated.Format(configuration.LastUpdatedFormat)
	md.TableOfContents = ctx.TableOfContents
	return md, err
}

// splitAndRender the document with the header elements on one side
// and all other elements (table of contents, with preamble, content) on the other side,
// then renders the header and other elements
func (r *sgmlRenderer) splitAndRender(ctx *renderer.Context, doc types.Document) (string, string, error) {
	switch doc.Attributes.GetAsStringWithDefault(types.AttrDocType, "article") {
	case "manpage":
		return r.splitAndRenderForManpage(ctx, doc)
	default:
		return r.splitAndRenderForArticle(ctx, doc)
	}
}

// splits the document with the title of the section 0 (if available) on one side
// and all other elements (table of contents, with preamble, content) on the other side
func (r *sgmlRenderer) splitAndRenderForArticle(ctx *renderer.Context, doc types.Document) (string, string, error) {
	if ctx.Config.IncludeHeaderFooter {
		if header, found := doc.Header(); found {
			renderedHeader, err := r.renderArticleHeader(ctx, header)
			if err != nil {
				return "", "", err
			}
			renderedContent, err := r.renderDocumentElements(ctx, header.Elements, doc.Footnotes)
			if err != nil {
				return "", "", err
			}
			return renderedHeader, renderedContent, nil
		}
	}
	renderedContent, err := r.renderDocumentElements(ctx, doc.Elements, doc.Footnotes)
	if err != nil {
		return "", "", err
	}
	return "", renderedContent, nil
}

// splits the document with the header elements on one side
// and the other elements (table of contents, with preamble, content) on the other side
func (r *sgmlRenderer) splitAndRenderForManpage(ctx *renderer.Context, doc types.Document) (string, string, error) {
	header, _ := doc.Header()
	nameSection := header.Elements[0].(types.Section)

	if ctx.Config.IncludeHeaderFooter {
		renderedHeader, err := r.renderManpageHeader(ctx, header, nameSection)
		if err != nil {
			return "", "", err
		}
		renderedContent, err := r.renderDocumentElements(ctx, header.Elements[1:], doc.Footnotes)
		if err != nil {
			return "", "", err
		}
		return renderedHeader, renderedContent, nil
	}
	// in that case, we still want to display the name section
	renderedHeader, err := r.renderManpageHeader(ctx, types.Section{}, nameSection)
	if err != nil {
		return "", "", err
	}
	renderedContent, err := r.renderDocumentElements(ctx, header.Elements[1:], doc.Footnotes)
	if err != nil {
		return "", "", err
	}
	result := &strings.Builder{}
	result.WriteString(renderedHeader)
	result.WriteString(renderedContent)
	return "", result.String(), nil
}

func (r *sgmlRenderer) renderDocumentRoles(ctx *Context, doc types.Document) (string, error) {
	if header, found := doc.Header(); found {
		return r.renderElementRoles(ctx, header.Attributes)
	}
	return "", nil
}

func (r *sgmlRenderer) renderDocumentID(doc types.Document) string {
	if header, found := doc.Header(); found {
		if header.Attributes.Has(types.AttrCustomID) {
			// We only want to emit a document body ID, if one was explicitly set
			return r.renderElementID(header.Attributes)
		}
	}
	return ""
}

func (r *sgmlRenderer) renderAuthors(doc types.Document) string {
	authors, found := doc.Authors()
	if !found {
		return ""
	}
	authorStrs := make([]string, len(authors))
	for i, author := range authors {
		authorStrs[i] = author.FullName
	}
	return strings.Join(authorStrs, "; ")
}

func (r *sgmlRenderer) renderDocumentTitle(ctx *renderer.Context, doc types.Document) (string, error) {
	if header, found := doc.Header(); found {
		// TODO: This feels wrong.  The title should not need markup.
		title, err := r.renderPlainText(ctx, header.Title)
		if err != nil {
			return "", errors.Wrap(err, "unable to render document title")
		}
		return string(title), nil
	}
	return "", nil
}

func (r *sgmlRenderer) renderArticleHeader(ctx *renderer.Context, header types.Section) (string, error) {
	renderedHeader, err := r.renderInlineElements(ctx, header.Title)
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

func (r *sgmlRenderer) renderManpageHeader(ctx *renderer.Context, header types.Section, nameSection types.Section) (string, error) {
	renderedHeader, err := r.renderInlineElements(ctx, header.Title)
	if err != nil {
		return "", err
	}
	renderedName, err := r.renderInlineElements(ctx, nameSection.Title)
	if err != nil {
		return "", err
	}
	description := nameSection.Elements[0].(types.Paragraph) // TODO: type check
	if description.Attributes == nil {
		description.Attributes = types.Attributes{}
	}
	description.Attributes.AddNonEmpty(types.AttrKind, "manpage")
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
		Content:   string(renderedContent), //nolint: gosec
		IncludeH1: len(renderedHeader) > 0,
	})
	if err != nil {
		return "", err
	}
	return output.String(), nil
}

// renderDocumentElements renders all document elements, including the footnotes,
// but not the HEAD and BODY containers
func (r *sgmlRenderer) renderDocumentElements(ctx *renderer.Context, source []interface{}, footnotes []types.Footnote) (string, error) {
	elements := []interface{}{}
	for i, e := range source {
		switch e := e.(type) {
		case types.Preamble:
			if !e.HasContent() { // why !HasContent ???
				// retain the preamble
				elements = append(elements, e)
				continue
			}
			// retain everything "as-is"
			elements = source
		case types.Section:
			if e.Level == 0 {
				// retain the section's elements...
				elements = append(elements, e.Elements)
				// ... and add the other elements (in case there's another section 0...)
				elements = append(elements, source[i+1:]...)
				continue
			}
			// retain everything "as-is"
			elements = source
		default:
			// retain everything "as-is"
			elements = source
		}
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
