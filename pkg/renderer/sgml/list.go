package sgml

import (
	"fmt"
	"io"
	"strings"
	texttemplate "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderList(ctx *context, l *types.List) (string, error) {
	switch l.Kind {
	case types.OrderedListKind:
		return r.renderOrderedList(ctx, l)
	case types.UnorderedListKind:
		return r.renderUnorderedList(ctx, l)
	case types.LabeledListKind:
		return r.renderLabeledList(ctx, l)
	case types.CalloutListKind:
		return r.renderCalloutList(ctx, l)
	default:
		return "", fmt.Errorf("unable to render list of kind '%s'", l.Kind)
	}
}

// -------------------------------------------------------
// Ordered Lists
// -------------------------------------------------------
func (r *sgmlRenderer) renderOrderedList(ctx *context, l *types.List) (string, error) {
	content := &strings.Builder{}

	for _, element := range l.Elements {
		e, ok := element.(*types.OrderedListElement)
		if !ok {
			return "", errors.Errorf("unable to render ordered list element of type '%T'", element)
		}
		if err := r.renderOrderedListElement(ctx, content, e); err != nil {
			return "", errors.Wrap(err, "unable to render ordered list")
		}
	}
	roles, err := r.renderElementRoles(ctx, l.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render ordered list roles")
	}
	style, err := getNumberingStyle(l)
	if err != nil {
		return "", errors.Wrap(err, "unable to render ordered list roles")
	}
	title, err := r.renderElementTitle(ctx, l.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render callout list roles")
	}
	return r.execute(r.orderedList, struct {
		Context   *context
		ID        string
		Title     string
		Roles     string
		Style     string
		ListStyle string
		Start     string
		Content   string
		Reversed  bool
	}{
		ID:        r.renderElementID(l.Attributes),
		Title:     title,
		Roles:     roles,
		Style:     style,
		ListStyle: r.numberingType(style),
		Start:     l.Attributes.GetAsStringWithDefault(types.AttrStart, ""),
		Content:   content.String(),
		Reversed:  l.Attributes.HasOption("reversed"),
	})
}

func getNumberingStyle(l *types.List) (string, error) {
	if s, found := l.Attributes.GetAsString(types.AttrStyle); found {
		return s, nil
	}
	e, ok := l.Elements[0].(*types.OrderedListElement)
	if !ok {
		return "", errors.Errorf("unable to render ordered list style based on element of type '%T'", l.Elements[0])
	}
	return e.Style, nil
}

// this numbering style is only really relevant to HTML
func (r *sgmlRenderer) numberingType(style string) string {
	switch style {
	case types.LowerAlpha:
		return `a`
	case types.UpperAlpha:
		return `A`
	case types.LowerRoman:
		return `i`
	case types.UpperRoman:
		return `I`
	default:
		return ""
	}
}

func (r *sgmlRenderer) renderOrderedListElement(ctx *context, w io.Writer, element *types.OrderedListElement) error {
	content, err := r.renderListElements(ctx, element.GetElements())
	if err != nil {
		return errors.Wrap(err, "unable to render ordered list element content")
	}
	tmpl, err := r.orderedListElement()
	if err != nil {
		return errors.Wrap(err, "unable to load ordered list element template")
	}
	if err = tmpl.Execute(w, struct {
		Context *context
		Content string
	}{
		Context: ctx,
		Content: string(content),
	}); err != nil {
		return errors.Wrap(err, "unable to render ordered list element")
	}
	return nil
}

// -------------------------------------------------------
// Unordered Lists
// -------------------------------------------------------
func (r *sgmlRenderer) renderUnorderedList(ctx *context, l *types.List) (string, error) {
	// make sure nested elements are aware of that their rendering occurs within a list
	checkList := false
	if len(l.Elements) > 0 {
		e, ok := l.Elements[0].(*types.UnorderedListElement)
		if !ok {
			return "", errors.Errorf("unable to render unordered list element of type '%T'", l.Elements[0])
		}
		if e.CheckStyle != types.NoCheck {
			checkList = true
		}
	}
	content := &strings.Builder{}

	for _, element := range l.Elements {
		if err := r.renderUnorderedListElement(ctx, content, element); err != nil {
			return "", errors.Wrap(err, "unable to render unordered list")
		}
	}
	roles, err := r.renderElementRoles(ctx, l.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render unordered list roles")
	}
	title, err := r.renderElementTitle(ctx, l.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render callout list roles")
	}
	return r.execute(r.unorderedList, struct {
		Context   *context
		ID        string
		Title     string
		Roles     string
		Style     string
		Checklist bool
		Items     []types.ListElement
		Content   string
	}{
		Context:   ctx,
		ID:        r.renderElementID(l.Attributes),
		Title:     title,
		Checklist: checkList,
		Items:     l.Elements,
		Content:   content.String(),
		Roles:     roles,
		Style:     r.renderElementStyle(l.Attributes),
	})
}
func (r *sgmlRenderer) renderUnorderedListElement(ctx *context, w io.Writer, element types.ListElement) error {
	content, err := r.renderListElements(ctx, element.GetElements())
	if err != nil {
		return errors.Wrap(err, "unable to render unordered list element content")
	}
	tmpl, err := r.unorderedListElement()
	if err != nil {
		return errors.Wrap(err, "unable to load unordered list element template")
	}
	if err := tmpl.Execute(w, struct {
		Context *context
		Content string
	}{
		Context: ctx,
		Content: string(content),
	}); err != nil {
		return errors.Wrap(err, "unable to render unordered list element")
	}
	return nil
}

// -------------------------------------------------------
// Labelled Lists
// -------------------------------------------------------
func (r *sgmlRenderer) renderLabeledList(ctx *context, l *types.List) (string, error) {
	tmpl, itemTmpl, err := r.getLabeledListTmpl(l)
	if err != nil {
		return "", errors.Wrap(err, "unable to render labeled list")
	}

	content := &strings.Builder{}
	cont := false
	for _, element := range l.Elements {
		e, ok := element.(*types.LabeledListElement)
		if !ok {
			return "", errors.Errorf("unable to render labeled list element of type '%T'", element)
		}
		if cont, err = r.renderLabeledListItem(ctx, itemTmpl, content, cont, e); err != nil {
			return "", errors.Wrap(err, "unable to render labeled list")
		}
	}
	roles, err := r.renderElementRoles(ctx, l.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render labeled list roles")
	}
	title, err := r.renderElementTitle(ctx, l.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render labeled list roles")
	}
	result := &strings.Builder{}
	if err := tmpl.Execute(result, struct {
		Context *context
		ID      string
		Title   string
		Roles   string
		Content string
		Items   []types.ListElement
	}{
		Context: ctx,
		ID:      r.renderElementID(l.Attributes),
		Title:   title,
		Roles:   roles,
		Content: content.String(),
		Items:   l.Elements,
	}); err != nil {
		return "", errors.Wrap(err, "unable to render labeled list")
	}
	return result.String(), nil
}

func (r *sgmlRenderer) getLabeledListTmpl(l *types.List) (*texttemplate.Template, *texttemplate.Template, error) {
	if layout, ok := l.Attributes[types.AttrStyle]; ok {
		switch layout {
		case "qanda":
			listTmpl, err := r.qAndAList()
			if err != nil {
				return nil, nil, errors.Wrap(err, "unable to load q&A list template")
			}
			listElementTmpl, err := r.qAndAListElement()
			if err != nil {
				return nil, nil, errors.Wrap(err, "unable to load q&A list element template")
			}
			return listTmpl, listElementTmpl, nil
		case "horizontal":
			listTmpl, err := r.labeledListHorizontal()
			if err != nil {
				return nil, nil, errors.Wrap(err, "unable to load horizontal list template")
			}
			listElementTmpl, err := r.labeledListHorizontalElement()
			if err != nil {
				return nil, nil, errors.Wrap(err, "unable to load horizontal list element template")
			}
			return listTmpl, listElementTmpl, nil
		default:
			return nil, nil, errors.Errorf("unsupported labeled list layout: %s", layout)
		}
	}
	listTmpl, err := r.labeledList()
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to load labeled list template")
	}
	listElementTmpl, err := r.labeledListElement()
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to load labeld list element template")
	}
	return listTmpl, listElementTmpl, nil
}

func (r *sgmlRenderer) renderLabeledListItem(ctx *context, tmpl *texttemplate.Template, w io.Writer, continuation bool, element *types.LabeledListElement) (bool, error) {

	term, err := r.renderInlineElements(ctx, element.Term)
	if err != nil {
		return false, errors.Wrap(err, "unable to render labeled list term")
	}
	content, err := r.renderListElements(ctx, element.Elements)
	if err != nil {
		return false, errors.Wrap(err, "unable to render labeled list content")
	}
	err = tmpl.Execute(w, struct {
		Context      *context
		Term         string
		Content      string
		Continuation bool
	}{
		Context:      ctx,
		Term:         string(term),
		Continuation: continuation,
		Content:      string(content),
	})
	return content == "", err
}

// -------------------------------------------------------
// Callout Lists
// -------------------------------------------------------
func (r *sgmlRenderer) renderCalloutList(ctx *context, l *types.List) (string, error) {
	content := &strings.Builder{}
	for _, element := range l.Elements {
		e, ok := element.(*types.CalloutListElement)
		if !ok {
			return "", errors.Errorf("unable to render callout list element of type '%T'", element)
		}
		rendererElement, err := r.renderCalloutListElement(ctx, e)
		if err != nil {
			return "", errors.Wrap(err, "unable to render callout list element")
		}
		content.WriteString(rendererElement)
	}
	roles, err := r.renderElementRoles(ctx, l.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render callout list roles")
	}
	title, err := r.renderElementTitle(ctx, l.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render callout list roles")
	}
	return r.execute(r.calloutList, struct {
		Context *context
		ID      string
		Title   string
		Roles   string
		Content string
		Items   []types.ListElement
	}{
		Context: ctx,
		ID:      r.renderElementID(l.Attributes),
		Title:   title,
		Roles:   roles,
		Content: content.String(),
		Items:   l.Elements,
	})
}

func (r *sgmlRenderer) renderCalloutListElement(ctx *context, element *types.CalloutListElement) (string, error) {
	content, err := r.renderListElements(ctx, element.Elements)
	if err != nil {
		return "", errors.Wrap(err, "unable to render callout list element content")
	}
	return r.execute(r.calloutListElement, struct {
		Context *context
		Ref     int
		Content string
	}{
		Context: ctx,
		Ref:     element.Ref,
		Content: string(content),
	})
}
