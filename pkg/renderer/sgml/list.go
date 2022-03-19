package sgml

import (
	"fmt"
	"io"
	"strings"
	text "text/template"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderList(ctx *renderer.Context, l *types.List) (string, error) {
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
func (r *sgmlRenderer) renderOrderedList(ctx *renderer.Context, l *types.List) (string, error) {
	result := &strings.Builder{}
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
	title, err := r.renderElementTitle(l.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render callout list roles")
	}
	err = r.orderedList.Execute(result, struct {
		Context   *renderer.Context
		ID        string
		Title     string
		Roles     string
		Style     string
		ListStyle string
		Start     string
		Content   string
		Reversed  bool
		// Elements  []types.ListElement
	}{
		ID:        r.renderElementID(l.Attributes),
		Title:     title,
		Roles:     roles,
		Style:     style,
		ListStyle: r.numberingType(style),
		Start:     l.Attributes.GetAsStringWithDefault(types.AttrStart, ""),
		Content:   string(content.String()),
		Reversed:  l.Attributes.HasOption("reversed"),
		// Elements:  l.Elements,
	})
	if err != nil {
		return "", errors.Wrap(err, "unable to render ordered list")
	}
	return result.String(), nil
}

func getNumberingStyle(l *types.List) (string, error) {
	if s, found, err := l.Attributes.GetAsString(types.AttrStyle); err != nil {
		return "", err
	} else if found {
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

func (r *sgmlRenderer) renderOrderedListElement(ctx *renderer.Context, w io.Writer, element *types.OrderedListElement) error {
	content, err := r.renderListElements(ctx, element.GetElements())
	if err != nil {
		return errors.Wrap(err, "unable to render unordered list element content")
	}
	return r.orderedListItem.Execute(w, struct {
		Context *renderer.Context
		Content string
	}{
		Context: ctx,
		Content: string(content),
	})
}

// -------------------------------------------------------
// Unordered Lists
// -------------------------------------------------------
func (r *sgmlRenderer) renderUnorderedList(ctx *renderer.Context, l *types.List) (string, error) {
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
	result := &strings.Builder{}
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
	title, err := r.renderElementTitle(l.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render callout list roles")
	}

	// here we must preserve the HTML tags
	err = r.unorderedList.Execute(result, struct {
		Context   *renderer.Context
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
		Content:   string(content.String()),
		Roles:     roles,
		Style:     r.renderElementStyle(l.Attributes),
	})
	if err != nil {
		return "", errors.Wrap(err, "unable to render unordered list")
	}
	return result.String(), nil
}
func (r *sgmlRenderer) renderUnorderedListElement(ctx *renderer.Context, w io.Writer, element types.ListElement) error {
	content, err := r.renderListElements(ctx, element.GetElements())
	if err != nil {
		return errors.Wrap(err, "unable to render unordered list element content")
	}
	return r.unorderedListItem.Execute(w, struct {
		Context *renderer.Context
		Content string
	}{
		Context: ctx,
		Content: string(content),
	})
}

// -------------------------------------------------------
// Labelled Lists
// -------------------------------------------------------
func (r *sgmlRenderer) renderLabeledList(ctx *renderer.Context, l *types.List) (string, error) {
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
	title, err := r.renderElementTitle(l.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render labeled list roles")
	}
	result := &strings.Builder{}
	// here we must preserve the HTML tags
	err = tmpl.Execute(result, struct {
		Context *renderer.Context
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
		Content: string(content.String()),
		Items:   l.Elements,
	})
	if err != nil {
		return "", errors.Wrap(err, "unable to render labeled list")
	}
	return result.String(), nil
}

func (r *sgmlRenderer) getLabeledListTmpl(l *types.List) (*text.Template, *text.Template, error) {
	if layout, ok := l.Attributes[types.AttrStyle]; ok {
		switch layout {
		case "qanda":
			return r.qAndAList, r.qAndAListItem, nil
		case "horizontal":
			return r.labeledListHorizontal, r.labeledListHorizontalItem, nil
		default:
			return nil, nil, errors.Errorf("unsupported labeled list layout: %s", layout)
		}
	}
	return r.labeledList, r.labeledListItem, nil
}

func (r *sgmlRenderer) renderLabeledListItem(ctx *renderer.Context, tmpl *text.Template, w io.Writer, continuation bool, element *types.LabeledListElement) (bool, error) {

	term, err := r.renderInlineElements(ctx, element.Term)
	if err != nil {
		return false, errors.Wrap(err, "unable to render labeled list term")
	}
	content, err := r.renderListElements(ctx, element.Elements)
	if err != nil {
		return false, errors.Wrap(err, "unable to render labeled list content")
	}
	err = tmpl.Execute(w, struct {
		Context      *renderer.Context
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
func (r *sgmlRenderer) renderCalloutList(ctx *renderer.Context, l *types.List) (string, error) {
	result := &strings.Builder{}
	content := &strings.Builder{}

	for _, element := range l.Elements {
		e, ok := element.(*types.CalloutListElement)
		if !ok {
			return "", errors.Errorf("unable to render callout list element of type '%T'", element)
		}
		if err := r.renderCalloutListItem(ctx, content, e); err != nil {
			return "", errors.Wrap(err, "unable to render callout list element")
		}
	}
	roles, err := r.renderElementRoles(ctx, l.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render callout list roles")
	}
	title, err := r.renderElementTitle(l.Attributes)
	if err != nil {
		return "", errors.Wrap(err, "unable to render callout list roles")
	}

	err = r.calloutList.Execute(result, struct {
		Context *renderer.Context
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
		Content: string(content.String()),
		Items:   l.Elements,
	})
	if err != nil {
		return "", errors.Wrap(err, "unable to render callout list")
	}
	return result.String(), nil
}

func (r *sgmlRenderer) renderCalloutListItem(ctx *renderer.Context, w io.Writer, element *types.CalloutListElement) error {
	content, err := r.renderListElements(ctx, element.Elements)
	if err != nil {
		return errors.Wrap(err, "unable to render callout list element content")
	}
	err = r.calloutListItem.Execute(w, struct {
		Context *renderer.Context
		Ref     int
		Content string
	}{
		Context: ctx,
		Ref:     element.Ref,
		Content: string(content),
	})
	if err != nil {
		return errors.Wrap(err, "unable to render callout list")
	}
	return nil
}
