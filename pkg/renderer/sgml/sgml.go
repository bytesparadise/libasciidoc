package sgml

import (
	"fmt"
	"io"
	"strings"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/types"

	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func Render(doc *types.Document, config *configuration.Configuration, output io.Writer, tmpls Templates) (types.Metadata, error) {
	r := &sgmlRenderer{
		templates: tmpls,
		// Establish some default function handlers.
		functions: texttemplate.FuncMap{
			"escape":             escapeString,
			"halign":             halign,
			"valign":             valign,
			"lastInStrings":      lastInStrings,
			"trimLineFeedSuffix": trimLineFeedSuffix,
		},
	}
	ctx := newContext(doc, config)

	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("rendering document of type '%s'\n%s", ctx.Attributes.GetAsStringWithDefault(types.AttrDocType, "article"), spew.Sdump(ctx.Attributes))
	// }

	// metadata to be returned to the caller
	var metadata types.Metadata
	// arguably this should be a time.Time for use in Go
	metadata.LastUpdated = ctx.config.LastUpdated.Format(configuration.LastUpdatedFormat)
	renderedTitle, exists, err := r.renderDocumentTitle(ctx, doc)
	if err != nil {
		return metadata, errors.Wrapf(err, "unable to render full document")
	}
	metadata.Title = string(renderedTitle) // retain an empty value if no title was defined in the document
	if !exists {
		renderedTitle = DefaultTitle
	}
	// process attribute declaration in the header
	if header, _ := doc.Header(); header != nil {
		for _, e := range header.Elements {
			switch e := e.(type) {
			case *types.AttributeDeclaration:
				ctx.attributes[e.Name] = e.Value
			case *types.AttributeReset:
				delete(ctx.attributes, e.Name)
			}
		}
	}
	// also, process standalone attriute declaration before the first section
elements:
	for _, e := range doc.Elements {
		switch e := e.(type) {
		case *types.AttributeDeclaration:
			ctx.attributes[e.Name] = e.Value
		case *types.AttributeReset:
			delete(ctx.attributes, e.Name)
		default:
			break elements
		}
	}
	if ctx.sectionNumbering, err = doc.SectionNumbers(); err != nil {
		return metadata, errors.Wrapf(err, "unable to render full document")
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
	if ctx.config.WrapInHTMLBodyElement {
		log.Debugf("Rendering full document...")
		tmpl, err := r.article()
		if err != nil {
			return metadata, errors.Wrapf(err, "unable to render full document")
		}
		err = tmpl.Execute(output, struct {
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
			CSS                   []string
			IncludeHTMLBodyHeader bool
			IncludeHTMLBodyFooter bool
		}{
			Doctype:               ctx.attributes.GetAsStringWithDefault(types.AttrDocType, "article"),
			Generator:             "libasciidoc", // TODO: externalize this value and include the lib version ?
			Description:           ctx.attributes.GetAsStringWithDefault(types.AttrDescription, ""),
			Title:                 renderedTitle,
			Authors:               r.renderAuthors(ctx, doc),
			Header:                renderedHeader,
			Roles:                 roles,
			ID:                    r.renderDocumentID(doc),
			Content:               string(renderedContent),
			RevNumber:             ctx.attributes.GetAsStringWithDefault("revnumber", ""),
			LastUpdated:           ctx.config.LastUpdated.Format(configuration.LastUpdatedFormat),
			CSS:                   ctx.config.CSS,
			IncludeHTMLBodyHeader: !ctx.attributes.Has(types.AttrNoHeader),
			IncludeHTMLBodyFooter: !ctx.attributes.Has(types.AttrNoFooter),
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

var predefinedAttributes = map[string]string{
	"sp":             " ",
	"blank":          "",
	"empty":          "",
	"nbsp":           "&#160;",
	"zwsp":           "&#8203;",
	"wj":             "&#8288;",
	"apos":           "&#39;",
	"quot":           "&#34;",
	"lsquo":          "&#8216;",
	"rsquo":          "&#8217;",
	"ldquo":          "&#8220;",
	"rdquo":          "&#8221;",
	"deg":            "&#176;",
	"plus":           "&#43;",
	"brvbar":         "&#166;",
	"vbar":           "|",
	"amp":            "&",
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

func lastInStrings(slice []string, index int) bool {
	return index == len(slice)-1
}

func trimLineFeedSuffix(content string) string {
	return strings.TrimSuffix(content, "\n")
}

// splitAndRender the document with the header elements on one side
// and all other elements (table of contents, with preamble, content) on the other side,
// then renders the header and other elements
func (r *sgmlRenderer) splitAndRender(ctx *context, doc *types.Document) (string, string, error) {
	switch ctx.attributes.GetAsStringWithDefault(types.AttrDocType, "article") {
	case "manpage":
		return r.splitAndRenderForManpage(ctx, doc)
	default:
		return r.splitAndRenderForArticle(ctx, doc)
	}
}

// splits the document with the title of the section 0 (if available) on one side
// and all other elements (table of contents, with preamble, content) on the other side
func (r *sgmlRenderer) splitAndRenderForArticle(ctx *context, doc *types.Document) (string, string, error) {
	log.Debugf("rendering article (within HTML/Body: %t)", ctx.config.WrapInHTMLBodyElement)

	header, _ := doc.Header()
	renderedHeader, err := r.renderDocumentHeader(ctx, header)
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
func (r *sgmlRenderer) splitAndRenderForManpage(ctx *context, doc *types.Document) (string, string, error) {
	log.Debugf("rendering manpage (within HTML/Body: %t)", ctx.config.WrapInHTMLBodyElement)
	elements := doc.BodyElements()
	nameSection := elements[0].(*types.Section) // TODO: enforce
	if ctx.config.WrapInHTMLBodyElement {
		header, _ := doc.Header()
		renderedHeader, err := r.renderManpageHeader(ctx, header, nameSection)
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

func (r *sgmlRenderer) renderDocumentRoles(ctx *context, doc *types.Document) (string, error) {
	if header, _ := doc.Header(); header != nil {
		return r.renderElementRoles(ctx, header.Attributes)
	}
	return "", nil
}

func (r *sgmlRenderer) renderDocumentID(doc *types.Document) string {
	if header, _ := doc.Header(); header != nil {
		// if header.Attributes.Has(types.AttrCustomID) {
		// We only want to emit a document body ID if one was explicitly set
		return r.renderElementID(header.Attributes)
		// }
	}
	return ""
}

func (r *sgmlRenderer) renderAuthors(_ *context, doc *types.Document) string { // TODO: pass header instead of doc
	if header, _ := doc.Header(); header != nil {
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

func (r *sgmlRenderer) renderDocumentTitle(_ *context, doc *types.Document) (string, bool, error) {
	if header, _ := doc.Header(); header != nil && header.Title != nil {
		// TODO: This feels wrong.  The title should not need markup.
		title, err := RenderPlainText(header.Title)
		if err != nil {
			return "", true, errors.Wrap(err, "unable to render document title")
		}
		return string(title), true, nil
	}
	return "", false, nil
}

func (r *sgmlRenderer) renderDocumentHeader(ctx *context, header *types.DocumentHeader) (string, error) {
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
	return r.execute(r.articleHeader, struct {
		Header  string
		Details string
	}{
		Header:  renderedHeader,
		Details: documentDetails,
	})
}

func (r *sgmlRenderer) renderDocumentHeaderTitle(ctx *context, header *types.DocumentHeader) (string, error) {
	if header == nil || header.Title == nil {
		log.Debug("no header to render")
		return "", nil
	}
	return r.renderElements(ctx, header.Title)
}

func (r *sgmlRenderer) renderManpageHeader(ctx *context, header *types.DocumentHeader, nameSection *types.Section) (string, error) {
	renderedHeader, err := r.renderDocumentHeaderTitle(ctx, header)
	if err != nil {
		return "", err
	}
	renderedName, err := r.renderElements(ctx, nameSection.Title)
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
	tmpl, err := r.manpageHeader()
	if err != nil {
		return "", errors.Wrap(err, "unable to load manpage header template")
	}
	if err = tmpl.Execute(output, struct {
		Header    string
		Name      string
		Content   string
		IncludeH1 bool
	}{
		Header:    renderedHeader,
		Name:      renderedName,
		Content:   string(renderedContent),
		IncludeH1: len(renderedHeader) > 0,
	}); err != nil {
		return "", errors.Wrap(err, "unable to render manpage header")
	}
	return output.String(), nil
}

// renderDocumentBody renders all document elements, including the footnotes,
// but not the HEAD and BODY containers
func (r *sgmlRenderer) renderDocumentBody(ctx *context, source []interface{}, toc *types.TableOfContents, footnotes []*types.Footnote) (string, error) {
	var elements []interface{}
	if placement, found := ctx.attributes[types.AttrTableOfContents]; found && toc != nil {
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
