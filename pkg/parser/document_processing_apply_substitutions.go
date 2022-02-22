package parser

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

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
	// stats := &Stats{}
	// opts := append(ctx.Opts,
	// GlobalStore(types.AttrImagesDir, ctx.attributes.get(types.AttrImagesDir)),
	// GlobalStore(usermacrosKey, ctx.userMacros),
	// )
	for i := range f.Elements {
		if err := applySubstitutionsOnElement(ctx, f.Elements[i], ctx.Opts...); err != nil {
			return types.NewErrorFragment(f.Position, err)
		}
	}
	return f
}

func applySubstitutionsOnElement(ctx *ParseContext, element interface{}, opts ...Option) error {
	switch b := element.(type) {
	case *types.FrontMatter:
		ctx.attributes.setAll(b.Attributes)
		return nil
	case *types.AttributeDeclaration:
		ctx.attributes.set(b.Name, b.Value)
		if b.Name == types.AttrExperimental {
			ctx.Opts = append(ctx.Opts, enableExperimentalMacros(true))
		}
		return nil
	case *types.AttributeReset:
		ctx.attributes.unset(b.Name)
		if b.Name == types.AttrExperimental {
			ctx.Opts = append(ctx.Opts, enableExperimentalMacros(false))
		}
		return nil
	case *types.DocumentHeader:
		return applySubstitutionsOnDocumentHeader(ctx, b, opts...)
	case *types.Section:
		return applySubstitutionsOnBlockWithTitle(ctx, b, opts...)
	case *types.Table:
		return applySubstitutionsOnTable(ctx, b, opts...)
	case *types.ListElementContinuation:
		return applySubstitutionsOnElement(ctx, b.Element, opts...)
	case types.WithElements:
		return applySubstitutionsOnBlockWithElements(ctx, b, opts...)
	case types.WithLocation:
		return applySubstitutionsOnBlockWithLocation(ctx, b, opts...)
	default:
		// do nothing
		return nil
	}
}

func applySubstitutionsOnDocumentHeader(ctx *ParseContext, b *types.DocumentHeader, opts ...Option) error {
	// process attribute declarations and resets defined in the header
	for _, elmt := range b.Elements {
		if err := applySubstitutionsOnElement(ctx, elmt, opts...); err != nil {
			return err
		}
	}
	if authors := b.Authors(); authors != nil {
		ctx.attributes.setAll(authors.Expand())
	}
	if revision := b.Revision(); revision != nil {
		ctx.attributes.setAll(revision.Expand())
	}
	return applySubstitutionsOnBlockWithTitle(ctx, b, opts...)
}

func applySubstitutionsOnBlockWithTitle(ctx *ParseContext, b types.WithTitle, opts ...Option) error {
	log.Debugf("processing element with title of type '%T' in 3 steps", b)
	if err := replaceAttributeRefsInBlockAttributes(ctx, b); err != nil {
		return err
	}
	opts = append(opts, Entrypoint("HeaderGroup"))
	// apply until Attribute substitution included
	// TODO: parse InlinePassthroughs and Attributes in the first pass (instead of dumb raw content)
	title, err := processSubstitutions(ctx, b.GetTitle(), headerSubstitutions(), opts...)
	if err != nil {
		return err
	}
	return b.SetTitle(title)
}

func applySubstitutionsOnTable(ctx *ParseContext, t *types.Table, opts ...Option) error {
	if err := replaceAttributeRefsInBlockAttributes(ctx, t); err != nil {
		return err
	}
	// also, deal with special `cols` attribute
	if cols, found := t.Attributes[types.AttrCols].(string); found {
		// parse with a specific rule
		values, err := Parse("", []byte(cols), append(opts, Entrypoint("TableColumnsAttribute"))...)
		if err != nil {
			return err
		}
		t.Attributes[types.AttrCols] = values
	}
	if t.Header != nil {
		for _, c := range t.Header.Cells {
			if err := applySubstitutionsOnBlockWithElements(ctx, c, opts...); err != nil {
				return err
			}
		}
	}
	for _, r := range t.Rows {
		for _, c := range r.Cells {
			if err := applySubstitutionsOnBlockWithElements(ctx, c, opts...); err != nil {
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
			if err := applySubstitutionsOnBlockWithElements(ctx, c, opts...); err != nil {
				return err
			}
		}
	}
	return nil
}

func applySubstitutionsOnBlockWithElements(ctx *ParseContext, b types.WithElements, opts ...Option) error {
	if err := replaceAttributeRefsInBlockAttributes(ctx, b); err != nil {
		return err
	}
	if s := b.GetAttributes().GetAsStringWithDefault(types.AttrStyle, ""); s == types.Passthrough {
		log.Debugf("skipping substitutions on passthrough block of type '%T'", b)
		content, _ := serialize(b.GetElements())
		return b.SetElements([]interface{}{
			&types.StringElement{
				Content: string(content),
			},
		})
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("applying substitutions on\n'%s'", spew.Sdump(b))
	}
	defer func() {
		if log.IsLevelEnabled(log.DebugLevel) {
			log.Debugf("applied substitutions:\n'%s'", spew.Sdump(b))
		}
	}()
	subs, err := newSubstitutions(b)
	if err != nil {
		return err
	}
	opts = append(opts, Entrypoint("Substitutions"))
	switch b := b.(type) {
	case *types.LabeledListElement:
		if b.Term, err = parseElements(b.Term, subs, opts...); err != nil {
			return err
		}
		for _, e := range b.GetElements() {
			if err := applySubstitutionsOnElement(ctx, e, opts...); err != nil {
				return err
			}
		}
		return nil
	case *types.DelimitedBlock:
		switch b.Kind {
		case types.Example, types.Quote, types.Sidebar:
			for _, e := range b.GetElements() {
				if err := applySubstitutionsOnElement(ctx, e, opts...); err != nil {
					return err
				}
			}
			return nil
		case types.MarkdownQuote:
			var attribution string
			if b.Elements, attribution = extractMarkdownQuoteAttribution(b.Elements); attribution != "" {
				b.Attributes = b.Attributes.Set(types.AttrQuoteAuthor, attribution)
			}
			b.Elements, err = processSubstitutions(ctx, b.Elements, subs, opts...)
			return err
		default:
			b.Elements, err = processSubstitutions(ctx, b.Elements, subs, opts...)
			return err
		}
	case *types.Paragraph:
		b.Elements, err = processSubstitutions(ctx, b.Elements, subs, opts...)
		return err
	case *types.TableCell:
		b.Elements, err = processSubstitutions(ctx, b.Elements, subs, opts...)
		return err
	default:
		for _, e := range b.GetElements() {
			if err := applySubstitutionsOnElement(ctx, e, opts...); err != nil {
				return err
			}
		}
		return nil
	}
}

func applySubstitutionsOnBlockWithLocation(ctx *ParseContext, b types.WithLocation, _ ...Option) error {
	_, _, err := replaceAttributeRefs(ctx, b)
	return err
}

func processSubstitutions(ctx *ParseContext, elements []interface{}, subs []string, opts ...Option) ([]interface{}, error) {
	// split the steps if attribute substitution is enabled
	for i, s := range subs {
		if s == Attributes {
			var err error
			// apply until Attribute substitution included
			if elements, err = parseElements(elements, subs, opts...); err != nil {
				return nil, err
			}
			// replace AttributeRefs found in previous step (inluding in inline element attributes)
			return replaceAttributeRefsAndReparse(ctx, elements, subs[i+1:], opts...)
		}
	}
	return parseElements(elements, subs, opts...)
}

// replaceAttributeRefsAndReparse recursively replaces the attribute refs, but only reparse the portions in which the replacements happened.
// for example: the location of an inline link, but not the whole paragraph.
func replaceAttributeRefsAndReparse(ctx *ParseContext, elements []interface{}, subs []string, opts ...Option) ([]interface{}, error) {
	found := false
	for i, e := range elements {
		switch e := e.(type) {
		case types.WithElements:
			// replace in attributes
			if err := replaceAttributeRefsInBlockAttributes(ctx, e); err != nil {
				return nil, err
			}
			// replace in elements
			if elements, err := replaceAttributeRefsAndReparse(ctx, e.GetElements(), subs, opts...); err != nil {
				return nil, err
			} else if err := e.SetElements(elements); err != nil {
				return nil, err
			}
		default:
			e, f, err := replaceAttributeRefs(ctx, e)
			if err != nil {
				return nil, err
			}
			found = found || f
			elements[i] = e
		}
	}
	if found {
		var err error
		elements, err = parseElements(elements, subs, opts...)
		if err != nil {
			return nil, err
		}
	}
	return elements, nil
}

// parseElements parse the elements, using placeholders for existing "structured" elements (ie, not RawLine or StringElements)
// Also, does not parse the content of the placeholders, but restores them at the end.
func parseElements(elements []interface{}, subs []string, opts ...Option) ([]interface{}, error) {
	serialized, placeholders := serialize(elements)
	if len(serialized) == 0 {
		return nil, nil
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("parsing '%s' with enabled substitutions %s", serialized, spew.Sdump(subs))
	}
	result, err := Parse("", serialized, append(opts, GlobalStore(enabledSubstitutions, subs))...)
	if err != nil {
		return nil, err
	}
	elmts, ok := result.([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected type of content after parsing elements: '%T'", result)
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("parsed content:\n%s", spew.Sdump(elmts))
	}
	return placeholders.restore(elmts)
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

func serialize(content []interface{}) ([]byte, *placeholders) {
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("serializing:\n%v", spew.Sdump(content))
	// }
	placeholders := newPlaceholders()
	result := bytes.NewBuffer(nil)
	for _, element := range content {
		switch element := element.(type) {
		case types.RawContent:
			result.WriteString(string(element))
		case types.RawLine:
			result.WriteString(string(element))
		case *types.SingleLineComment:
			// replace with placeholder
			p := placeholders.add(element)
			result.WriteString(p.String())
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
	return result.Bytes(), placeholders
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
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("restoring placeholders in\n%s", spew.Sdump(elements))
	// }
	// skip if there's nothing to restore
	if len(p.elements) == 0 {
		return elements, nil
	}
	for i, e := range elements {
		if e, ok := e.(*types.ElementPlaceHolder); ok {
			elements[i] = p.elements[e.Ref]
		}
	}
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("restored elements:\n%v", spew.Sdump(elements))
	// }
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
	log.Debugf("experimental enabled: %t", (found && enabled))
	return found && enabled
}

// ----------------------------------------------
// Substitutions
// ----------------------------------------------

func (c *current) lookupCurrentSubstitutions() ([]string, bool) {
	s, found := c.globalStore[enabledSubstitutions].([]string)
	return s, found
}

func (c *current) isSubstitutionEnabled(k string) bool {
	subs, found := c.lookupCurrentSubstitutions()
	if !found {
		// log.Debugf("substitutions not set in globalStore: assuming '%s' not enabled", k)
		return false // TODO: should return `true`, at least for `attributes`?
	}
	for _, s := range subs {
		if s == k {
			// log.Debugf("'%s' is enabled", k)
			return true
		}
	}
	// log.Debugf("'%s' is not enabled", k)
	return false
}

func newSubstitutions(b types.WithElements) ([]string, error) {
	// TODO: introduce a `types.BlockWithSubstitution` interface?
	// note: would also be helpful for paragraphs with `[listing]` style.
	defaultSubs, err := defaultSubstitutions(b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to determine substitutions")
	}
	// look-up the `subs` in the element's attributes
	attrSub, found, err := b.GetAttributes().GetAsString(types.AttrSubstitutions)
	if err != nil {
		return nil, err
	}
	if !found {
		return defaultSubs, nil
	}
	subs := strings.Split(attrSub, ",")
	var result []string
	// when dealing with incremental substitutions, use default sub as a baseline and append or prepend the incremental subs
	if allIncremental(subs) {
		result = defaultSubs
	} else {
		result = make([]string, 0, len(subs))
	}
	for _, sub := range subs {
		// log.Debugf("checking subs '%s'", sub)
		switch {
		case strings.HasSuffix(sub, "+"): // prepend
			s, err := substitutions(strings.TrimSuffix(sub, "+"))
			if err != nil {
				return nil, err
			}
			// log.Debugf("prepending subs '%s'", sub)
			result = append(s, result...)
		case strings.HasPrefix(sub, "+"): // append
			s, err := substitutions(strings.TrimPrefix(sub, "+"))
			if err != nil {
				return nil, err
			}
			// log.Debugf("appending subs '%s'", sub)
			result = append(result, s...)
		case strings.HasPrefix(sub, "-"): // remove from all substitutions
			s, err := substitutions(strings.TrimPrefix(sub, "-"))
			if err != nil {
				return nil, err
			}
		loop:
			for i := range s {
				for j := range result {
					if s[i] == result[j] {
						log.Debugf("removing '%s' at index %d", s[i], j)
						result = append(result[:j], result[j+1:]...) // remove at index `j`
						continue loop
					}
				}
			}
		default:
			s, err := substitutions(sub)
			if err != nil {
				return nil, err
			}
			result = append(result, s...)
		}
	}

	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("substitutions to apply on block of type '%T': %s", b, spew.Sdump(result))
	// }
	return result, nil
}

// checks if all the given subs are incremental (ie, prefixed with `+|-` or suffixed with `-`)
func allIncremental(subs []string) bool {
	for _, sub := range subs {
		if !isIncrementalSubstitution(sub) {
			return false
		}
	}
	return true
}

func isIncrementalSubstitution(sub string) bool {
	return strings.HasPrefix(sub, "+") ||
		strings.HasPrefix(sub, "-") ||
		strings.HasSuffix(sub, "+")
}

func normalSubstitutions() []string {
	return []string{
		InlinePassthroughs,
		Attributes,
		SpecialCharacters,
		Quotes,
		Replacements,
		Macros,
		PostReplacements,
	}
}

func headerSubstitutions() []string {
	return []string{
		InlinePassthroughs,
		Attributes,
		SpecialCharacters,
		Quotes,
		Macros,
		Replacements,
	}
}

func noneSubstitutions() []string {
	return []string{}
}

func verbatimSubstitutions() []string {
	return []string{
		Callouts,
		SpecialCharacters,
	}
}

func defaultSubstitutions(b types.WithElements) ([]string, error) {
	// log.Debugf("looking-up default substitution for block of type '%T'", b)
	switch b := b.(type) {
	case *types.DelimitedBlock:
		switch b.Kind {
		case types.Listing, types.Fenced, types.Literal:
			return verbatimSubstitutions(), nil
		case types.Example, types.Quote, types.Verse, types.Sidebar, types.MarkdownQuote:
			return normalSubstitutions(), nil
		case types.Comment, types.Passthrough:
			return noneSubstitutions(), nil
		default:
			return nil, fmt.Errorf("unsupported kind of delimited block: '%v'", b.Kind)
		}
	case *types.Paragraph:
		// if listing paragraph:
		switch b.GetAttributes().GetAsStringWithDefault(types.AttrStyle, "") {
		case types.Listing:
			return verbatimSubstitutions(), nil
		default:
			return normalSubstitutions(), nil
		}
	case *types.ListElements, types.ListElement, *types.QuotedText, *types.Table, *types.TableCell:
		return normalSubstitutions(), nil
	default:
		return nil, fmt.Errorf("unsupported kind of element: '%T'", b)
	}
}

func substitutions(s string) ([]string, error) {
	switch s {
	case "normal":
		return normalSubstitutions(), nil
	case "none":
		return noneSubstitutions(), nil
	case "verbatim":
		return verbatimSubstitutions(), nil
	case "attributes", "macros", "quotes", "replacements", "post_replacements", "callouts", "specialchars":
		return []string{
			s,
		}, nil
	default:
		// TODO: return `none` instead of `err` and log an error with the fragment position (use logger with fields?)
		return nil, fmt.Errorf("unsupported substitution: '%v'", s)
	}
}

const (
	// enabledSubstitutions the key in which enabled substitutions are stored in the parser's GlobalStore
	enabledSubstitutions string = "enabled_substitutions"

	// Attributes the "attributes" substitution
	Attributes string = "attributes"
	// Callouts the "callouts" substitution
	Callouts string = "callouts"
	// InlinePassthroughs the "inline_passthrough" substitution
	InlinePassthroughs string = "inline_passthrough" // nolint:gosec
	// Macros the "macros" substitution
	Macros string = "macros"
	// None the "none" substitution
	None string = "none"
	// PostReplacements the "post_replacements" substitution
	PostReplacements string = "post_replacements"
	// Quotes the "quotes" substitution
	Quotes string = "quotes"
	// Replacements the "replacements" substitution
	Replacements string = "replacements"
	// SpecialCharacters the "specialchars" substitution
	SpecialCharacters string = "specialchars"
)

func counterToStringElement(ctx *ParseContext, c *types.CounterSubstitution) (string, error) {
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

func replaceAttributeRefsInBlockAttributes(ctx *ParseContext, b types.WithAttributes) error {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("replacing attribute refs in attributes of block of type '%T':\n%s", b, spew.Sdump(b.GetAttributes()))
	}
	attrs := b.GetAttributes()
	for k, v := range attrs {
		switch v := v.(type) {
		case []interface{}:
			v, _, err := replaceAttributeRefsInSlice(ctx, v)
			if err != nil {
				return err
			}
			attrs[k] = types.Reduce(v)
		case types.Roles:
			roles := make(types.Roles, len(v))
			for i, r := range v {
				r, _, err := replaceAttributeRefs(ctx, r)
				if err != nil {
					return err
				}
				roles[i] = types.Reduce(r)
			}
			attrs[k] = roles
		case types.Options:
			options := make(types.Options, len(v))
			for i, r := range v {
				r, _, err := replaceAttributeRefs(ctx, r)
				if err != nil {
					return err
				}
				options[i] = types.Reduce(r)
			}
			attrs[k] = options
		default:
			// do nothing
		}
	}
	b.SetAttributes(attrs)
	return nil
}

func replaceAttributeRefsInSlice(ctx *ParseContext, elements []interface{}) ([]interface{}, bool, error) {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("replacing attribute refs in slice %s", spew.Sdump(elements))
	}
	found := false
	for i, e := range elements {
		v, f, err := replaceAttributeRefs(ctx, e)
		if err != nil {
			return nil, false, err
		}
		elements[i] = v
		found = found || f
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("replaced attribute refs: %s", spew.Sdump(elements))
	}
	return elements, found, nil // reduce elements to return a single `string` when it's possible
}

func replaceAttributeRefs(ctx *ParseContext, b interface{}) (interface{}, bool, error) {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("replacing attribute refs in %s", spew.Sdump(b))
	}
	switch b := b.(type) {
	case []interface{}:
		return replaceAttributeRefsInSlice(ctx, b)
	case types.WithElements:
		// replace in attributes
		if err := replaceAttributeRefsInBlockAttributes(ctx, b); err != nil {
			return nil, false, err
		}
		// replace in elements
		elements, found, err := replaceAttributeRefsInSlice(ctx, b.GetElements())
		if err != nil {
			return nil, false, err
		}
		if err := b.SetElements(elements); err != nil {
			return nil, false, err
		}
		return b, found, nil
	case types.WithLocation:
		// replace in attributes
		if err := replaceAttributeRefsInBlockAttributes(ctx, b); err != nil {
			return nil, false, err
		}
		// replace in location
		if b.GetLocation() == nil {
			// skip
			return b, false, nil
		}
		p, found, err := replaceAttributeRefs(ctx, b.GetLocation().Path)
		if err != nil {
			return nil, false, err
		}
		b.GetLocation().SetPath(p)
		return b, found, nil
	case *types.AttributeReference:
		e, err := replaceAttributeRef(ctx, b)
		if err != nil {
			return nil, false, err
		}
		return e, true, nil
	case *types.CounterSubstitution:
		s, err := counterToStringElement(ctx, b)
		if err != nil {
			return nil, false, err
		}
		return &types.StringElement{
			Content: s,
		}, true, nil
	default:
		// do nothing, keep as-is
		return b, false, nil
	}
}

func replaceAttributeRef(ctx *ParseContext, a *types.AttributeReference) (*types.StringElement, error) {
	s, found, err := ctx.attributes.getAsString(a.Name)
	if err != nil {
		return nil, err
	} else if !found {
		log.Warnf("unable to find attribute '%s'", a.Name)
		return &types.StringElement{
			Content: "{" + a.Name + "}",
		}, nil
	}
	return &types.StringElement{
		Content: s,
	}, nil
}
