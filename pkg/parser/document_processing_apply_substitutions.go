package parser

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bytesparadise/libasciidoc/pkg/renderer/sgml"
	"github.com/bytesparadise/libasciidoc/pkg/types"

	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
)

// ApplySubstitutions parses the "inline content" of the incomgin fragment elements
// (eg, paragraph content to convert rawlines into slices of StringElement, QuotedText, InlineLinks, etc.),
// while also applying the required substitutions (default or custom)
func ApplySubstitutions(ctx *ParseContext, done <-chan interface{}, fragmentStream <-chan types.DocumentFragment) chan types.DocumentFragment {
	processedFragmentStream := make(chan types.DocumentFragment, bufferSize)
	go func() {
		defer close(processedFragmentStream)
		for f := range fragmentStream {
			select {
			case processedFragmentStream <- applySubstitutionsOnFragment(ctx, f):
			case <-done:
				log.WithField("pipeline_stage", "apply_substitutions").Debug("received 'done' signal")
				return
			}
		}
		log.WithField("pipeline_stage", "apply_substitutions").Debug("done")
	}()
	return processedFragmentStream
}

func applySubstitutionsOnFragment(ctx *ParseContext, f types.DocumentFragment) types.DocumentFragment {
	if f.Error != nil {
		log.Debugf("skipping substitutions because of fragment with error: %v", f.Error)
		return f
	}
	start := time.Now()
	if err := applySubstitutionsOnElements(ctx, f.Elements); err != nil {
		return types.NewErrorFragment(f.Position, err)
	}
	log.Debugf("time to apply substitutions on fragment at %d: %d microseconds", f.Position.Start, time.Since(start).Microseconds())
	return f
}

func applySubstitutionsOnElements(ctx *ParseContext, elements []interface{}, opts ...Option) error {
	for _, e := range elements {
		if err := applySubstitutionsOnElement(ctx, e, opts...); err != nil {
			return err
		}
	}
	return nil
}

func applySubstitutionsOnElement(ctx *ParseContext, element interface{}, opts ...Option) error {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("applying substitutions on element of type '%T'", element)
	}
	switch b := element.(type) {
	case *types.FrontMatter:
		ctx.attributes.setAll(b.Attributes)
		return nil
	case *types.DocumentHeader:
		return applySubstitutionsOnDocumentHeader(ctx, b, opts...)
	case *types.AttributeDeclaration:
		return applySubstitutionsOnAttributeDeclaration(ctx, b, opts...)
	case *types.AttributeReset:
		return applySubstitutionsOnAttributeReset(ctx, b, opts...)
	case *types.Section:
		return applySubstitutionsOnWithTitle(ctx, b, opts...)
	case *types.Table:
		return applySubstitutionsOnTable(ctx, b, opts...)
	case *types.ListContinuation:
		return applySubstitutionsOnElement(ctx, b.Element, opts...)
	case types.WithElements:
		return applySubstitutionsOnWithElements(ctx, b, opts...)
	case types.WithLocation:
		return applySubstitutionsOnWithLocation(ctx, b, opts...)
	default:
		// do nothing
		return nil
	}
}

func applySubstitutionsOnDocumentHeader(ctx *ParseContext, b *types.DocumentHeader, opts ...Option) error {
	// process attribute declarations and resets defined in the header *before* the title itself
	if err := applySubstitutionsOnElements(ctx, b.Elements, opts...); err != nil {
		return err
	}
	if authors := b.Authors(); authors != nil {
		ctx.attributes.setAll(authors.Expand())
	}
	if revision := b.Revision(); revision != nil {
		ctx.attributes.setAll(revision.Expand())
	}
	return applySubstitutionsOnWithTitle(ctx, b, opts...)
}

// TODO: move that to "parse fragment" phase, ie, merge `AttributeDeclarationValueGroup` rule into `AttributeDeclarationValue`?
func applySubstitutionsOnAttributeDeclaration(ctx *ParseContext, b *types.AttributeDeclaration, _ ...Option) error {
	v, err := replaceAttributeRefsInValue(ctx, b.Value)
	if err != nil {
		return err
	}
	b.Value = v
	ctx.attributes.set(b.Name, b.Value)
	if b.Name == types.AttrExperimental {
		ctx.opts = append(ctx.opts, enableExperimentalMacros(true)) // TODO: add `ctx.ExperimentalMacrosEnabled()` instead?
	}
	return nil
}

func applySubstitutionsOnAttributeReset(ctx *ParseContext, b *types.AttributeReset, _ ...Option) error {
	ctx.attributes.unset(b.Name)
	if b.Name == types.AttrExperimental {
		ctx.opts = append(ctx.opts, enableExperimentalMacros(false))
	}
	return nil
}

func applySubstitutionsOnWithTitle(ctx *ParseContext, b types.WithTitle, _ ...Option) error {
	log.Debugf("processing element with title of type '%T'", b)
	// attributes
	if err := replaceAttributeRefsInAttributeValues(ctx, b.GetAttributes()); err != nil {
		return err
	}
	if err := reparseAttributes(ctx, b); err != nil {
		return err
	}
	phase1, phase2 := headerSubstitutions().split()
	if phase1.contains(AttributeRefs) {
		title, err := replaceAttributeRefsInElementsAndReparse(ctx, b.GetTitle(), phase2)
		if err != nil {
			return err
		}
		b.SetTitle(title)
	}
	return nil
}

func applySubstitutionsOnTable(ctx *ParseContext, t *types.Table, opts ...Option) error {
	// attributes
	if err := replaceAttributeRefsInAttributeValues(ctx, t.GetAttributes()); err != nil {
		return err
	}
	if err := reparseAttributes(ctx, t); err != nil {
		return err
	}

	// rows and cells
	if cols, ok := t.Attributes[types.AttrCols].([]interface{}); ok {
		t.SetColumnDefinitions(cols)
	}
	if t.Header != nil {
		for _, c := range t.Header.Cells {
			// assume elements to process were wrapped in a paragraph at parse time
			if len(c.Elements) == 1 {
				if p, ok := c.Elements[0].(*types.Paragraph); ok {
					if err := applySubstitutionsOnWithElements(ctx, p, opts...); err != nil {
						return err
					}
				}
			}
		}
	}
	for _, r := range t.Rows {
		for _, c := range r.Cells {
			if err := applySubstitutionsOnWithElements(ctx, c, opts...); err != nil {
				return err
			}
		}
	}
	// do not retain rows if empty
	if len(t.Rows) == 0 {
		t.Rows = nil
	}
	if t.Footer != nil {
		for _, c := range t.Footer.Cells {
			if err := applySubstitutionsOnWithElements(ctx, c, opts...); err != nil {
				return err
			}
		}
	}
	return nil
}

func applySubstitutionsOnWithLocation(ctx *ParseContext, b types.WithLocation, opts ...Option) error {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("applying substitution on WithLocation of type '%T'", b)
	}
	// attributes
	if err := replaceAttributeRefsInAttributeValues(ctx, b.GetAttributes()); err != nil {
		return err
	}
	if err := reparseAttributes(ctx, b, opts...); err != nil {
		return err
	}

	// location
	if b.GetLocation() != nil {
		if p, ok := b.GetLocation().Path.([]interface{}); ok {
			p, err := replaceAttributeRefsInElements(ctx, p) // no need to reparse the location's path
			if err != nil {
				return err
			}
			b.GetLocation().Path = types.Reduce(p)
		}
	}
	return nil
}

func applySubstitutionsOnWithElements(ctx *ParseContext, b types.WithElements, opts ...Option) error {
	// attributes
	if err := replaceAttributeRefsInAttributeValues(ctx, b.GetAttributes()); err != nil {
		return err
	}
	if err := reparseAttributes(ctx, b, opts...); err != nil {
		return err
	}
	subs, err := newSubstitutions(b)
	if err != nil {
		return err
	}
	opts = append(opts, Entrypoint("NormalGroup")) // TODO: move this into NewParseContext ?
	switch b := b.(type) {
	case *types.LabeledListElement:
		if b.Term, err = applySubstitutionsOnSlice(ctx, b.Term, subs, opts...); err != nil {
			return err
		}
		return applySubstitutionsOnElements(ctx, b.Elements, opts...)
	case *types.DelimitedBlock:
		switch b.Kind {
		case types.Example, types.Quote, types.Sidebar, types.Open: // TODO: add a func on *types.DelimitedBlock to avoid checking the exact same kinds in multiple places
			return applySubstitutionsOnElements(ctx, b.Elements, opts...)
		case types.MarkdownQuote:
			var attribution string
			if b.Elements, attribution = extractMarkdownQuoteAttribution(b.Elements); attribution != "" {
				b.Attributes = b.Attributes.Set(types.AttrQuoteAuthor, attribution)
			}
			b.Elements, err = applySubstitutionsOnSlice(ctx, b.Elements, subs, opts...)
			return err
		default:
			b.Elements, err = applySubstitutionsOnSlice(ctx, b.Elements, subs, opts...)
			return err
		}
	case *types.Paragraph:
		b.Elements, err = applySubstitutionsOnSlice(ctx, b.Elements, subs, opts...)
		// reparse attributes in inline elements :/
		for _, e := range b.Elements {
			if e, ok := e.(types.WithAttributes); ok {
				if err := reparseInlineAttributes(ctx, e, subs, opts...); err != nil {
					return err
				}
			}
		}
		return err
	default:
		return applySubstitutionsOnElements(ctx, b.GetElements(), opts...)
	}
}

func reparseAttributes(ctx *ParseContext, element types.WithAttributes, opts ...Option) error {
	// in some specific cases, we need to reparse values to support links and quoted texts
	attrs := element.GetAttributes()
	for k, v := range attrs {
		switch k {
		case types.AttrTitle, types.AttrXRefLabel:
			v, err := parseWithSubstitutions(v, attributeSubstitutions(), append(append(ctx.opts, Entrypoint("AttributeStructuredValue")), opts...)...)
			if err != nil {
				return err
			}
			attrs[k] = types.Reduce(v)
		case types.AttrInlineLinkText:
			subs := attributeSubstitutions()
			// same as above, but do not allow for inline macros (eg: links)
			if err := subs.remove(Macros); err != nil {
				return err
			}
			v, err := parseWithSubstitutions(v, subs, append(append(ctx.opts, Entrypoint("AttributeStructuredValue")), opts...)...)
			if err != nil {
				return err
			}
			attrs[k] = types.Reduce(v)
		case types.AttrCols:
			s, err := serializePlainText(v) // TODO: call sgml.RenderPlainText?
			if err != nil {
				return err
			}
			v, err := parseWithSubstitutions(s, attributeSubstitutions(), append(append(ctx.opts, Entrypoint("TableColumnsAttribute")), opts...)...)
			if err != nil {
				return err
			}
			attrs[k] = types.Reduce(v)
		default:
			attrs[k] = v
		}
	}
	return nil
}

func reparseInlineAttributes(ctx *ParseContext, element types.WithAttributes, subs *substitutions, opts ...Option) error {
	// in some specific cases, we need to reparse values to support links and quoted texts
	attrs := element.GetAttributes()
	for k, v := range attrs {
		switch k {
		case types.AttrTitle, types.AttrXRefLabel:
			v, err := parseWithSubstitutions(v, subs, append(append(ctx.opts, opts...), Entrypoint("AttributeStructuredValue"))...)
			if err != nil {
				return err
			}
			attrs[k] = types.Reduce(v)
		case types.AttrInlineLinkText:
			// same as above, but do not allow for inline macros (eg: links)
			if err := subs.remove(Macros); err != nil {
				return err
			}
			v, err := parseWithSubstitutions(v, subs, append(append(ctx.opts, opts...), Entrypoint("AttributeStructuredValue"))...)
			if err != nil {
				return err
			}
			attrs[k] = types.Reduce(v)
		default:
			attrs[k] = v
		}
	}
	// also, recursively perform the same changes on nested elements with attributes
	switch e := element.(type) {
	case types.WithElements:
		for _, e := range e.GetElements() {
			if e, ok := e.(types.WithAttributes); ok {
				if err := reparseInlineAttributes(ctx, e, subs, opts...); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func replaceAttributeRefsInAttributeValues(ctx *ParseContext, attrs types.Attributes) error {
	var err error
	for k, v := range attrs {
		if attrs[k], err = replaceAttributeRefsInValue(ctx, v); err != nil {
			return err
		}
	}
	return nil
}

func replaceAttributeRefsInValue(ctx *ParseContext, value interface{}) (interface{}, error) {
	var err error
	switch value := value.(type) {
	case []interface{}:
		return replaceAttributeRefsInSlicedValue(ctx, value)
	case types.Roles:
		for i, r := range value {
			if value[i], err = replaceAttributeRefsInValue(ctx, r); err != nil {
				return nil, err
			}
		}
		return value, nil
	case types.Options:
		for i, r := range value {
			if value[i], err = replaceAttributeRefsInValue(ctx, r); err != nil {
				return nil, err
			}
		}
		return value, nil
	default:
		return value, nil
	}
}

func replaceAttributeRefsInSlicedValue(ctx *ParseContext, elements []interface{}) (interface{}, error) {
	result := make([]interface{}, 0, len(elements))
	for _, e := range elements {
		switch e := e.(type) {
		case *types.StringElement:
			result = append(result, e.Content)
		case *types.AttributeReference:
			v, _, err := valueForAttributeRef(ctx, e)
			if err != nil {
				return nil, err
			}
			result = append(result, v)
		case []interface{}: // entries in Roles and Options
			v, err := replaceAttributeRefsInSlicedValue(ctx, e)
			if err != nil {
				return nil, err
			}
			result = append(result, v)
		default:
			result = append(result, e) // unchanged
		}
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("replaced attribute refs: %s", spew.Sdump(result))
	}
	return types.Reduce(result), nil
}

func applySubstitutionsOnSlice(ctx *ParseContext, elements []interface{}, subs *substitutions, opts ...Option) ([]interface{}, error) {
	if len(elements) == 0 {
		// nothing to do
		return nil, nil
	}
	// TODO: no need to parse with `inline_passthrough,attributes` substitutions if the elements do not contain `{` or `pass:` ?
	phase1, phase2 := subs.split()
	// if len(plan) == 0 {
	// 	return nil, fmt.Errorf("unable to apply substitutions: empty plan")
	// }
	// parse
	var err error
	elements, err = parseWithSubstitutions(elements, phase1, append(ctx.opts, opts...)...)
	if err != nil {
		return nil, err
	}
	// replace attribute references if applicable
	if phase1.contains(AttributeRefs) {
		if elements, err = replaceAttributeRefsInElementsAndReparse(ctx, elements, phase2); err != nil {
			return nil, err
		}
	}
	return elements, nil
}

func extractMarkdownQuoteAttribution(elements []interface{}) ([]interface{}, string) {
	// first, check if last line is an attribution (author)
	if len(elements) == 0 {
		return elements, ""
	}
	log.Debugf("attempting to extract markdown-style quote block author")
	if l, ok := elements[len(elements)-1].(types.RawLine); ok {
		a, err := ParseReader("", strings.NewReader(string(l)), Entrypoint("MarkdownQuoteAttribution"))
		// assume that the last line is not an author attribution if an error occurred
		if err != nil {
			log.Debugf("failed to extract markdown-style quote block author: %v", err)
			return elements, ""
		}
		log.Debugf("found attribution in markdown block: '%[1]v' (%[1]T)", a)
		if a, ok := a.(string); ok {
			return elements[:len(elements)-1], a
		}
	}
	return elements, ""
}

func replaceAttributeRefsInElementsAndReparse(ctx *ParseContext, elements []interface{}, subs *substitutions, opts ...Option) ([]interface{}, error) {
	replaced := false
	for i, e := range elements {
		switch e := e.(type) {
		case *types.AttributeReference:
			v, r, err := valueForAttributeRef(ctx, e)
			if err != nil {
				return nil, err
			}
			switch v := v.(type) {
			case string:
				elements[i] = &types.StringElement{
					Content: v,
				}
			default:
				elements[i] = v
			}
			replaced = replaced || r
		case *types.CounterSubstitution:
			v, err := valueForCounter(ctx, e)
			if err != nil {
				return nil, err
			}
			elements[i] = &types.StringElement{
				Content: v,
			}
		case types.WithElements: // if `subs=macros,attributes`, then replace within inline macros
			// in attributes
			if err := replaceAttributeRefsInAttributeValues(ctx, e.GetAttributes()); err != nil {
				return nil, err
			}
			// in elements
			elmts, err := replaceAttributeRefsInElementsAndReparse(ctx, e.GetElements(), subs, opts...)
			if err != nil {
				return nil, err
			}
			if err := e.SetElements(elmts); err != nil {
				return nil, err
			}
		case types.WithLocation: // if `subs=macros,attributes`, then replace within inline macros
			// in attributes
			if err := replaceAttributeRefsInAttributeValues(ctx, e.GetAttributes()); err != nil {
				return nil, err
			}
			// in location
			if e.GetLocation() != nil {
				if p, ok := e.GetLocation().Path.([]interface{}); ok {
					p, err := replaceAttributeRefsInElements(ctx, p) // no need to reparse the location's path
					if err != nil {
						return nil, err
					}
					e.GetLocation().Path = types.Reduce(p)
				}
			}
		}
	}
	// if attribute ref was found, reparse content (if subs)
	if replaced && subs != nil {
		if log.IsLevelEnabled(log.DebugLevel) {
			log.Debugf("reparsing (phase2) %s", spew.Sdump(elements))
		}
		return parseWithSubstitutions(elements, subs, append(opts, Entrypoint("NormalGroup"))...)
	}
	return elements, nil
}

func replaceAttributeRefsInElements(ctx *ParseContext, elements []interface{}) ([]interface{}, error) {
	for i, e := range elements {
		switch e := e.(type) {
		case *types.AttributeReference:
			v, _, err := valueForAttributeRef(ctx, e)
			if err != nil {
				return nil, err
			}
			switch v := v.(type) {
			case string:
				elements[i] = &types.StringElement{
					Content: v,
				}
			default:
				elements[i] = v
			}
		case *types.CounterSubstitution:
			v, err := valueForCounter(ctx, e)
			if err != nil {
				return nil, err
			}
			elements[i] = &types.StringElement{
				Content: v,
			}
		default:
			// do nothing, keep as-is
		}
	}
	return elements, nil
}

func valueForAttributeRef(ctx *ParseContext, a *types.AttributeReference) (interface{}, bool, error) {
	v, found := ctx.attributes.get(a.Name)
	if !found {
		log.Warnf("unable to find entry for attribute with key '%s' in context", a.Name)
		return "{" + a.Name + "}", false, nil
	}
	switch v := v.(type) {
	case []interface{}:
		s, err := sgml.RenderPlainText(v)
		return s, true, err
	case *types.DocumentAuthorFullName:
		return v.FullName(), true, nil
	default:
		return v, true, nil
	}
}

func valueForCounter(ctx *ParseContext, c *types.CounterSubstitution) (string, error) {
	counter := ctx.counters[c.Name]
	if counter == nil {
		counter = 0
	}
	increment := true
	if c.Value != nil {
		ctx.counters[c.Name] = c.Value
		counter = c.Value
		increment = false
	}
	switch counter := counter.(type) {
	case int:
		if increment {
			counter++
		}
		ctx.counters[c.Name] = counter
		if c.Hidden {
			// return empty string facilitates merging
			return "", nil
		}
		return strconv.Itoa(counter), nil
	case rune:
		if increment {
			counter++
		}
		ctx.counters[c.Name] = counter
		if c.Hidden {
			// return empty string facilitates merging
			return "", nil
		}
		return string(counter), nil
	default:
		return "", fmt.Errorf("unexpected type of counter value: '%T'", counter)
	}
}

// parseElementsWithSubstitutions parse the elements, using placeholders for existing "structured" elements (ie, not RawLine or StringElements)
// Also, does not parse the content of the placeholders, but restores them at the end.
func parseWithSubstitutions(content interface{}, subs *substitutions, opts ...Option) ([]interface{}, error) {
	serialized, placeholders, err := serialize(content)
	if err != nil {
		return nil, err
	}
	if len(serialized) == 0 {
		return nil, nil
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("parsing '%s' with '%s' substitutions", serialized, subs.toString())
	}
	elements, err := parseContent(serialized, append(opts, GlobalStore(enabledSubstitutionsKey, subs))...)
	if err != nil {
		return nil, err
	}
	elements, err = placeholders.restore(elements)
	if err != nil {
		return nil, err
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("parsed content:\n%s", spew.Sdump(elements))
	}
	return elements, nil
}

func serialize(content interface{}) ([]byte, *placeholders, error) {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("serializing:\n%v", spew.Sdump(content))
	}
	placeholders := newPlaceholders()
	switch content := content.(type) {
	case string:
		return []byte(content), placeholders, nil
	case []interface{}:
		result := bytes.NewBuffer(nil)
		for _, element := range content {
			switch element := element.(type) {
			case types.RawLine:
				result.WriteString(string(element))
			case string:
				result.WriteString(string(element))
			case *types.StringElement:
				result.WriteString(element.Content)
			case *types.SpecialCharacter:
				result.WriteString(element.Name) // reserialize, so we can detect bare URLs (eg: `a link to <{base_url}>`)
			default:
				// replace with placeholder
				p := placeholders.add(element)
				result.WriteString(p.String())
			}
		}
		// if log.IsLevelEnabled(log.DebugLevel) {
		// 	log.Debugf("serialized lines: '%s'\nplaceholders: %v", result.Bytes(), spew.Sdump(placeholders.elements))
		// }
		return result.Bytes(), placeholders, nil
	default:
		return nil, nil, fmt.Errorf("unexpected type of content to serialize: %T", content)
	}
}

func serializePlainText(content interface{}) (string, error) {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("serializing plaintext:\n%v", spew.Sdump(content))
	}
	switch content := content.(type) {
	case string:
		return content, nil
	case []interface{}:
		result := bytes.NewBuffer(nil)
		for _, element := range content {
			switch element := element.(type) {
			case []interface{}:
				s, err := serializePlainText(element)
				if err != nil {
					return "", err
				}
				result.WriteString(s)
			case types.RawLine:
				result.WriteString(string(element))
			case *types.StringElement:
				result.WriteString(element.Content)
			case *types.SpecialCharacter:
				result.WriteString(element.Name) // reserialize, so we can detect bare URLs (eg: `a link to <{base_url}>`)
			case *types.InlinePassthrough:
				c, err := serializePlainText(element.Elements)
				if err != nil {
					return "", err
				}
				result.WriteString(c) // reserialize, so we can detect bare URLs (eg: `a link to <{base_url}>`)
			default:
				return "", fmt.Errorf("unexpected type of content to serialize as plain text: %T", element)
			}
		}
		if log.IsLevelEnabled(log.DebugLevel) {
			log.Debugf("serialized plaintext: '%s'", result.Bytes())
		}
		return result.String(), nil
	default:
		return "", fmt.Errorf("unexpected type of content to serialize as plain text: %T", content)
	}
}

type placeholders struct {
	seq      int
	elements map[string]interface{}
}

func newPlaceholders() *placeholders {
	return &placeholders{
		seq:      0,
		elements: map[string]interface{}{},
	}
}

func (p *placeholders) add(element interface{}) *types.ElementPlaceHolder {
	p.seq++
	p.elements[strconv.Itoa(p.seq)] = element
	return &types.ElementPlaceHolder{
		Ref: strconv.Itoa(p.seq),
	}

}

// replace the placeholders with their original element in the given elements
func (p *placeholders) restore(elements []interface{}) ([]interface{}, error) {
	for i, e := range elements {
		switch e := e.(type) {
		case *types.ElementPlaceHolder:
			elements[i] = p.elements[e.Ref]
		case types.WithElements:
			elmts, err := p.restore(e.GetElements())
			if err != nil {
				return nil, err
			}
			if err := e.SetElements(elmts); err != nil {
				return nil, err
			}
		}
	}
	return elements, nil
}

// ----------------------------------------------
// Support for experimental macros at parse time
// ----------------------------------------------
const experimentalMacrosKey string = "experimental_macros"

// sets the `experimental_macros` flag to `true` or `false` in the global store, so it can be checked by the `isExperimentalEnabled` method
func enableExperimentalMacros(enabled bool) Option {
	return GlobalStore(experimentalMacrosKey, enabled)
}

// checks if the `experimental` doc attribute was set (no value is expected, but we set a flag to handle the case where the attribute was reset)
func (c *current) isExperimentalEnabled() bool {
	enabled, found := c.globalStore[experimentalMacrosKey].(bool)
	// log.Debugf("experimental enabled: %t", (found && enabled))
	return found && enabled
}
