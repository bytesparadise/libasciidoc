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

// TODO: convert `ctx *parseContext` as a local variable instead of a func param
func ApplySubstitutions(ctx *ParseContext, done <-chan interface{}, fragmentStream <-chan types.DocumentFragment) chan types.DocumentFragment {
	processedFragmentStream := make(chan types.DocumentFragment, bufferSize)
	go func() {
		defer close(processedFragmentStream)
		for f := range fragmentStream {
			select {
			case processedFragmentStream <- applySubstitutions(ctx, f):
			case <-done:
				log.WithField("pipeline_stage", "apply_substitutions").Debug("received 'done' signal")
				return
			}
		}
		log.WithField("pipeline_stage", "apply_substitutions").Debug("done")
	}()
	return processedFragmentStream
}

func applySubstitutions(ctx *ParseContext, f types.DocumentFragment) types.DocumentFragment {
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.WithField("pipeline_stage", "apply_substitutions").Debugf("incoming fragment:\n%s", spew.Sdump(f))
	// }
	// if the fragment already contains an error, then send it as-is downstream
	if f.Error != nil {
		log.Debugf("skipping substitutions because of fragment with error: %v", f.Error)
		return f
	}
	// stats := &Stats{}
	opts := append(ctx.Opts,
		GlobalStore(types.AttrImagesDir, ctx.attributes.get(types.AttrImagesDir)),
		GlobalStore(usermacrosKey, ctx.userMacros),
	)

	elements := make([]interface{}, len(f.Elements))
	for i, element := range f.Elements {
		var err error
		if elements[i], err = applySubstitutionsOnElement(ctx, element, opts...); err != nil {
			return types.NewErrorFragment(f.Position, err)
		}
	}
	f.Elements = elements
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	// log.WithField("pipeline_stage", "apply_substitutions").Debugf("fragment with substitutions applied:\n%s", spew.Sdump(f))
	// 	log.WithField("pipeline_stage", "apply_substitutions").Debugf("stats:\n%s", PrettyPrintStats(stats))
	// }
	return f
}

func applySubstitutionsOnElement(ctx *ParseContext, element interface{}, opts ...Option) (interface{}, error) {
	switch e := element.(type) {
	case *types.FrontMatter:
		ctx.attributes.setAll(e.Attributes)
		return e, nil
	case *types.AttributeDeclaration:
		ctx.attributes.set(e.Name, e.Value)
		return e, nil
	case *types.AttributeReset:
		ctx.attributes.unset(e.Name)
		return e, nil
	case *types.DocumentHeader:
		if err := applySubstitutionsOnHeader(ctx, e, opts...); err != nil {
			return nil, err
		}
		return e, nil
	case types.BlockWithElements:
		if err := applySubstitutionsOnBlockWithElements(ctx, e, opts...); err != nil {
			return nil, err
		}
		return e, nil
	case types.BlockWithLocation:
		if err := applySubstitutionsOnBlockWithLocation(ctx, e, opts...); err != nil {
			return nil, err
		}
		return e, nil
	default:
		// log.WithField("pipeline_stage", "fragment_processing").Debugf("forwarding fragment content of type '%T' as-is", e)
		return element, nil
	}
}

// special case for document header: process the optional attribute declarations first, then the title
func applySubstitutionsOnHeader(ctx *ParseContext, header *types.DocumentHeader, opts ...Option) error {
	// TODO: duplicate from aggregate stage
	for _, elmt := range header.Elements {
		switch attr := elmt.(type) {
		case *types.AttributeDeclaration:
			ctx.attributes.set(attr.Name, attr.Value)
		case *types.AttributeReset:
			ctx.attributes.unset(attr.Name)
		}
	}
	title, err := newHeaderSubstitution().processElements(ctx, header.Title, opts...)
	if err != nil {
		return err
	}
	header.Title = title
	return nil
}

func applySubstitutionsOnBlockWithElements(ctx *ParseContext, block types.BlockWithElements, opts ...Option) error {
	if s := block.GetAttributes().GetAsStringWithDefault(types.AttrStyle, ""); s == types.Passthrough {
		log.Debugf("skipping substitutions on passthrough block of type '%T'", block)
		// simply merge the rawlines
		// TODO: use types.Merge?
		buf := &strings.Builder{}
		for i, l := range block.GetElements() {
			if l, ok := l.(types.RawLine); ok {
				buf.WriteString(string(l))
				if i < len(block.GetElements())-1 {
					buf.WriteString("\n")
				}
			}
		}
		return block.SetElements([]interface{}{
			&types.StringElement{
				Content: buf.String(),
			},
		})
	}
	log.Debugf("processing block with elements of type '%T'", block)
	if err := newElementAttributesSubstitution().processAttributes(ctx, block); err != nil {
		return err
	}
	// log.Debugf("applying substitutions on elements of block of type '%T'", block)
	s, err := newSubstitutions(block)
	if err != nil {
		return err
	}
	// also process extra stuff
	// TODO: move after call to `s.processBlockWithElements`?
	switch b := block.(type) {
	case *types.ListElements:
		for _, e := range b.Elements {
			if e, ok := e.(*types.LabeledListElement); ok {
				// process term of labeled list elements
				if e.Term, err = s.processElements(ctx, e.Term, opts...); err != nil {
					return err
				}
				e.Term = types.TrimLeft(e.Term, " ")
			}
		}
	case *types.DelimitedBlock:
		// process author of markdown-style quote block
		if b.Kind == types.MarkdownQuote {
			if elements, attribution := extractMarkdownQuoteAttribution(b.Elements); attribution != "" {
				b.Attributes = b.Attributes.Set(types.AttrQuoteAuthor, attribution)
				b.Elements = elements
			}
		}
	case *types.Table:
		if b.Header != nil {
			if err = s.processBlockWithElements(ctx, b.Header, opts...); err != nil {
				return err
			}
		}
		if b.Footer != nil {
			if err = s.processBlockWithElements(ctx, b.Footer, opts...); err != nil {
				return err
			}
		}
	}
	if err := s.processBlockWithElements(ctx, block, opts...); err != nil {
		return err
	}
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("processed block of type '%T' with elements: %s", block, spew.Sdump(block))
	// }
	return nil
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

func applySubstitutionsOnBlockWithLocation(ctx *ParseContext, block types.BlockWithLocation, opts ...Option) error {
	log.Debugf("processing block with attributes and location")
	if err := newElementAttributesSubstitution().processAttributes(ctx, block, opts...); err != nil {
		return err
	}
	// log.Debugf("applying substitutions on `location` of block of type '%T'", block)
	elements := block.GetLocation().Path
	elements, _, err := replaceAttributeRefsInElements(ctx, elements)
	if err != nil {
		return err
	}
	block.GetLocation().SetPath(elements)
	switch img := block.(type) {
	case *types.ImageBlock, *types.InlineImage:
		imagesdir := ctx.attributes.getAsStringWithDefault(types.AttrImagesDir, "")
		img.GetLocation().SetPathPrefix(imagesdir)
	}
	return nil
}

type substitutions []*substitution

func newSubstitutions(b types.BlockWithAttributes) (substitutions, error) {
	// TODO: introduce a `types.BlockWithSubstitution` interface?
	// note: would also be helpful for paragraphs with `[listing]` style.
	defaultSub, err := defaultSubstitution(b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to determine substitutions")
	}
	subs := strings.Split(b.GetAttributes().GetAsStringWithDefault(types.AttrSubstitutions, defaultSub), ",")
	result := make([]*substitution, 0, len(subs))
	// when dealing with incremental substitutions
	if allIncremental(subs) {
		d, err := newSubstitution(defaultSub)
		if err != nil {
			return nil, err
		}
		result = substitutions{d}
	}
	for _, sub := range subs {
		// log.Debugf("checking subs '%s'", sub)
		switch {
		case strings.HasSuffix(sub, "+"): // prepend
			s, err := newSubstitution(strings.TrimSuffix(sub, "+"))
			if err != nil {
				return nil, err
			}
			// log.Debugf("prepending subs '%s'", sub)
			result = append(substitutions{s}, result...)
		case strings.HasPrefix(sub, "+"): // append
			s, err := newSubstitution(strings.TrimPrefix(sub, "+"))
			if err != nil {
				return nil, err
			}
			// log.Debugf("appending subs '%s'", sub)
			result = append(result, s)
		case strings.HasPrefix(sub, "-"): // remove from all substitutions
			for _, s := range result {
				s.disable(substitutionKind(strings.TrimPrefix(sub, "-")))
			}
		default:
			s, err := newSubstitution(sub)
			if err != nil {
				return nil, err
			}
			result = append(result, s)
		}
	}

	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debug("substitutions to apply:")
	// 	for _, s := range result {
	// 		log.Debugf("%s: %+v", s.entrypoint, s.rules)
	// 	}
	// }
	return result, nil
}

// checks if all the given subs are incremental (ie, prefixed with `+|-` or suffixed with `-`)
func allIncremental(subs []string) bool {
	// init with first substitution
	allIncremental := true
	// check others
	for _, sub := range subs {
		if !isIncrementalSubstitution(sub) {
			return false
		}
	}
	return allIncremental
}

func isIncrementalSubstitution(sub string) bool {
	return strings.HasPrefix(sub, "+") ||
		strings.HasPrefix(sub, "-") ||
		strings.HasSuffix(sub, "+")
}

func (s substitutions) processBlockWithElements(ctx *ParseContext, block types.BlockWithElements, opts ...Option) error {
	elements, err := s.processElements(ctx, block.GetElements(), opts...)
	if err != nil {
		return err
	}
	if err := block.SetElements(elements); err != nil {
		return err
	}
	return nil
}

func (s substitutions) processElements(ctx *ParseContext, elements []interface{}, opts ...Option) ([]interface{}, error) {
	// skip if there's nothing to do
	// (and no need to return an empty slice, btw)
	if len(elements) == 0 {
		return nil, nil
	}
	for _, substitution := range s {
		var err error
		elements, err = substitution.processElements(ctx, elements, opts...)
		if err != nil {
			return nil, err
		}
	}
	return elements, nil
}

func defaultSubstitution(b interface{}) (string, error) {
	// log.Debugf("looking-up default substitution for block of type '%T'", b)
	switch b := b.(type) {
	case *types.DelimitedBlock:
		switch b.Kind {
		case types.Listing, types.Fenced, types.Literal:
			return "verbatim", nil
		case types.Example, types.Quote, types.Verse, types.Sidebar, types.MarkdownQuote:
			return "normal", nil
		case types.Comment, types.Passthrough:
			return "none", nil
		default:
			return "", fmt.Errorf("unsupported kind of delimited block: '%v'", b.Kind)
		}
	case *types.Paragraph:
		// if listing paragraph:
		switch b.GetAttributes().GetAsStringWithDefault(types.AttrStyle, "") {
		case types.Listing:
			return "verbatim", nil
		default:
			return "normal", nil
		}
	case *types.ListElements, types.ListElement, *types.QuotedText, *types.Table:
		return "normal", nil
	case *types.Section:
		return "header", nil
	default:
		return "", fmt.Errorf("unsupported kind of element: '%T'", b)
	}
}

func canHaveSubstitution(b interface{}) bool {
	switch b.(type) {
	case *types.DelimitedBlock, *types.Paragraph, *types.List, types.ListElement, *types.QuotedText, *types.Section:
		return true
	default:
		return false
	}
}

type substitutionRule string

const (
	AttributesGroup        substitutionRule = "AttributesGroup"
	ElementAttributesGroup substitutionRule = "ElementAttributesGroup"
	HeaderGroup            substitutionRule = "HeaderGroup"
	MacrosGroup            substitutionRule = "MacrosGroup"
	NoneGroup              substitutionRule = "NoneGroup"
	NormalGroup            substitutionRule = "NormalGroup"
	QuotesGroup            substitutionRule = "QuotesGroup"
	ReplacementsGroup      substitutionRule = "ReplacementsGroup"
	PostReplacementsGroup  substitutionRule = "PostReplacementsGroup"
	SpecialcharactersGroup substitutionRule = "SpecialCharactersGroup"
	VerbatimGroup          substitutionRule = "VerbatimGroup"
)

type substitution struct {
	entrypoint substitutionRule
	rules      map[substitutionKind]bool
	// hasAttributeSubstitutions bool // TODO: replace with a key in the parser's store
}

func newSubstitution(kind string) (*substitution, error) {
	switch kind {
	case "attributes":
		return newAttributesSubstitution(), nil
	case "element_attributes":
		return newElementAttributesSubstitution(), nil
	case "header":
		return newHeaderSubstitution(), nil
	case "macros":
		return newMacrosSubstitution(), nil
	case "normal":
		return newNormalSubstitution(), nil
	case "none":
		return newNoneSubstitution(), nil
	case "quotes":
		return newQuotesSubstitution(), nil
	case "replacements":
		return newReplacementsSubstitution(), nil
	case "post_replacements":
		return newPostReplacementsSubstitution(), nil
	case "specialchars":
		return newSpecialCharsSubstitution(), nil
	case "verbatim":
		return newVerbatimSubstitution(), nil
	default:
		return nil, fmt.Errorf("unsupported kind of substitution: '%v'", kind)
	}
}

func (s *substitution) clone() *substitution {
	rules := make(map[substitutionKind]bool, len(s.rules))
	for r, e := range s.rules {
		rules[r] = e
	}
	return &substitution{
		entrypoint: s.entrypoint,
		rules:      rules,
		// hasAttributeSubstitutions: false, // make sure it's set to false
	}
}
func newAttributesSubstitution() *substitution {
	return &substitution{
		entrypoint: AttributesGroup,
		rules: map[substitutionKind]bool{
			InlinePassthroughs: true,
			Attributes:         true,
		},
	}
}

func newElementAttributesSubstitution() *substitution {
	return &substitution{
		entrypoint: ElementAttributesGroup,
		rules: map[substitutionKind]bool{
			InlinePassthroughs: true,
			Attributes:         true,
			Quotes:             true,
			SpecialCharacters:  true, // TODO: is it needed?
		},
	}
}

func newNormalSubstitution() *substitution {
	return &substitution{
		entrypoint: NormalGroup,
		rules: map[substitutionKind]bool{
			InlinePassthroughs: true,
			SpecialCharacters:  true,
			Attributes:         true,
			Quotes:             true,
			Replacements:       true,
			Macros:             true,
			PostReplacements:   true,
		},
	}
}

func newHeaderSubstitution() *substitution {
	return &substitution{
		entrypoint: HeaderGroup,
		rules: map[substitutionKind]bool{
			InlinePassthroughs: true,
			SpecialCharacters:  true,
			Attributes:         true,
			Quotes:             true,
			Macros:             true,
			Replacements:       true,
		},
	}
}

func newMacrosSubstitution() *substitution {
	return &substitution{
		entrypoint: MacrosGroup,
		rules: map[substitutionKind]bool{
			Macros: true,
		},
	}
}

func newNoneSubstitution() *substitution {
	return &substitution{
		entrypoint: NoneGroup,
		rules:      map[substitutionKind]bool{},
	}
}

func newQuotesSubstitution() *substitution {
	return &substitution{
		entrypoint: QuotesGroup,
		rules: map[substitutionKind]bool{
			Quotes: true,
		},
	}
}

func newReplacementsSubstitution() *substitution {
	return &substitution{
		entrypoint: ReplacementsGroup,
		rules: map[substitutionKind]bool{
			Replacements: true,
		},
	}
}

func newPostReplacementsSubstitution() *substitution {
	return &substitution{
		entrypoint: PostReplacementsGroup,
		rules: map[substitutionKind]bool{
			PostReplacements: true,
		},
	}
}

func newSpecialCharsSubstitution() *substitution {
	return &substitution{
		entrypoint: SpecialcharactersGroup,
		rules: map[substitutionKind]bool{
			SpecialCharacters: true,
		},
	}
}

func newVerbatimSubstitution() *substitution {
	return &substitution{
		entrypoint: VerbatimGroup,
		rules: map[substitutionKind]bool{
			SpecialCharacters: true,
			Callouts:          true,
		},
	}
}

func newTableColumnsAttrSubstitution() *substitution {
	return &substitution{
		entrypoint: "TableColumnsAttribute",
	}
}

func (s *substitution) disable(kinds ...substitutionKind) {
	log.Debugf("disabling %v", kinds)
	rules := make(map[substitutionKind]bool, len(s.rules))
	// TODO: use a single loop to copy/remove
	for k, v := range s.rules {
		rules[k] = v
	}
	for _, k := range kinds {
		switch k {
		case "specialchars":
			delete(rules, SpecialCharacters)
		default:
			delete(rules, k)
		}
	}
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("remaining rules: %s", spew.Sdump(rules))
	// }
	s.rules = rules
}

func (s *substitution) hasEnablements() bool {
	return len(s.rules) > 0
}

func (s *substitution) processAttributes(ctx *ParseContext, element types.BlockWithAttributes, opts ...Option) error {
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.WithField("entrypoint", s.entrypoint).
	// 		Debugf("processing attributes of %s", spew.Sdump(element))
	// }
	// TODO: only parse element attributes if an attribute substitution occurred?
	attrs, err := s.parseAttributes(element.GetAttributes(), opts...)
	if err != nil {
		return err
	}
	element.SetAttributes(attrs)
	found, err := replaceAttributeRefsInAttributes(ctx, element)
	if err != nil {
		return err
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("attributes have substitions: %t", found)
	}
	if found {
		// clone substitution so we start again with `hasAttributeSubstitutions=false` without loosing this `s.hasAttributeSubstitutions` value (needed to re-process elements)
		return s.clone().processAttributes(ctx, element, opts...)
	}
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.WithField("entrypoint", s.entrypoint).Debugf("done processing attributes of %s", spew.Sdump(element))
	// }
	return nil
}

func (s *substitution) processElements(ctx *ParseContext, elements []interface{}, opts ...Option) ([]interface{}, error) {
	if log.IsLevelEnabled(log.DebugLevel) {
		// log.WithField("substitution_entrypoint", s.entrypoint).WithField("substitution_rules", s.rules).Debugf("processing %s", spew.Sdump(elements))
		log.
			WithField("substitution_entrypoint", s.entrypoint).
			WithField("substitution_rules", s.rules).
			// WithField("hasAttributeSubstitutions", s.hasAttributeSubstitutions).
			Debugf("processing elements")
	}
	elements, err := s.parseElements(ctx, elements, opts...)
	if err != nil {
		return nil, err
	}
	// log.WithField("group", step.group).Debug("attribute substitutions detected during parsing")
	// apply substitutions on elements
	elements, found, err := replaceAttributeRefsInElements(ctx, elements)
	if err != nil {
		return nil, err
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("content has attribute substitutions: %t", found)
	}
	if found {
		elements, _ = types.NewInlineElements(elements)
		// re-run the parser, skipping the `inline_passthrough` and `attribute` rules this time
		s.disable(InlinePassthroughs) // TODO: verify when elements contain an `InlinePassthroughs`
		if s.hasEnablements() {
			if elements, err = s.parseElements(ctx, elements, opts...); err != nil {
				return nil, err
			}
		}
	}
	return elements, nil
}

func (s *substitution) parseElements(ctx *ParseContext, elements []interface{}, opts ...Option) ([]interface{}, error) {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.WithField("entrypoint", s.entrypoint).Debug("parsing elements")
	}
	serialized, placeholders := serialize(elements)
	if len(serialized) == 0 {
		return nil, nil
	}
	result, err := s.parseContent(serialized, opts...)
	if err != nil {
		return nil, err
	}
	for _, element := range result {
		if e, ok := element.(types.BlockWithAttributes); ok {
			if len(e.GetAttributes()) == 0 {
				log.Debug("no attribute to parse")
				continue
			}
			// clone substitution, because we don't need all features (eg: inlinemacros)
			sa := s.clone()
			log.Debugf("disabling macros while parsing attributes")
			sa.disable(Macros) // TODO: disable more rules (only keep what's in ElementAttributesGroup)
			// skip call below if `sa` is empty.
			if len(e.GetAttributes()) == 0 {
				log.Debug("no rule enabled to parse attributes")
				continue
			}
			attrs, err := sa.parseAttributes(e.GetAttributes(), opts...)
			if err != nil {
				return nil, err
			}
			e.SetAttributes(attrs)
		}
	}

	// also, apply the substitutions on the placeholders, case by case
	for _, element := range placeholders.elements {
		if err := s.processPlaceHolderElements(ctx, element, opts...); err != nil {
			return nil, err
		}
	}
	return placeholders.restore(result)
}

func (s *substitution) processPlaceHolderElements(ctx *ParseContext, element interface{}, opts ...Option) error {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("processing placeholder element of type '%T'", element)
	}
	switch e := element.(type) {
	case types.BlockWithElements:
		// if element has its own substitution attribute, then use it
		if canHaveSubstitution(e) && e.GetAttributes().Has(types.AttrSubstitutions) {
			subs, err := newSubstitutions(e)
			if err != nil {
				return err
			}
			return subs.processBlockWithElements(ctx, e, opts...)
		}
		if err := s.processBlockWithElements(ctx, e, opts...); err != nil {
			return err
		}
	case types.BlockWithLocation:
		if err := s.processElementWithLocation(ctx, e, opts...); err != nil {
			return err
		}
	case *types.ListElementContinuation:
		var err error
		if e.Element, err = applySubstitutionsOnElement(ctx, e.Element); err != nil {
			return err
		}
	case []interface{}:
		for _, elmt := range e {
			if w, ok := elmt.(types.BlockWithElements); ok {
				if err := s.processPlaceHolderElements(ctx, w, opts...); err != nil {
					return err
				}
			}
		}
	default:
		log.Debugf("skipping substitutions on block of type '%T'", e)
	}
	return nil
}

func (s *substitution) processBlockWithElements(ctx *ParseContext, block types.BlockWithElements, opts ...Option) error {
	elements, err := s.processElements(ctx, block.GetElements(), opts...)
	if err != nil {
		return err
	}
	return block.SetElements(elements)
}

func (s *substitution) processElementWithLocation(ctx *ParseContext, element types.BlockWithLocation, opts ...Option) error {
	log.Debugf("processing element with location of type '%T'", element)
	// process attributes
	if err := s.processAttributes(ctx, element); err != nil {
		return err
	}
	// process path
	pathElements, err := s.processElements(ctx, element.GetLocation().Path, opts...)
	if err != nil {
		return err
	}
	element.GetLocation().SetPath(pathElements)
	return nil
}

func (s *substitution) parseAttributes(attributes types.Attributes, opts ...Option) (types.Attributes, error) {
	if !(s.entrypoint == NormalGroup || s.entrypoint == AttributesGroup || s.entrypoint == QuotesGroup || s.entrypoint == ElementAttributesGroup) { // TODO: include special_characters?
		// log.Debugf("no need to parse attributes for group '%s'", s.entrypoint)
		return attributes, nil
	}
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("parsing attributes with group '%s' on %s", s.entrypoint, spew.Sdump(attributes))
	// }
	for name, value := range attributes {
		switch value := value.(type) {
		case []interface{}:
			// TODO: skip serializing/parsing if there's a single, non-string item
			// Eg:
			// ([]interface {}) (len=1) {
			// 	(*types.AttributeSubstitution)({
			// 	 Name: (string) (len=12) "github-title"
			// 	})
			//    }
			serialized, placeholders := serialize(value)
			if len(serialized) == 0 {
				continue
			}
			values, err := s.parseContent(serialized, opts...)
			if err != nil {
				return nil, err
			}
			values, err = placeholders.restore(values)
			if err != nil {
				return nil, err
			}
			attributes[name] = values
		case string:
			if name == types.AttrCols {
				// parse with a specific rule
				values, err := newTableColumnsAttrSubstitution().parseContent([]byte(value), opts...)
				if err != nil {
					return nil, err
				}
				attributes[name] = values
				continue
			}
			values, err := s.parseContent([]byte(value), opts...)
			if err != nil {
				return nil, err
			}
			attributes[name] = types.Reduce(values)
		default:
			attributes[name] = value
		}
	}
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("parsed attributes with rule '%s': %s", s.entrypoint, spew.Sdump(attributes))
	// }
	return attributes, nil
}

func (s *substitution) parseContent(content []byte, opts ...Option) ([]interface{}, error) {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.WithField("entrypoint", s.entrypoint).WithField("substitution_rules", s.rules).Debugf("parsing '%s'", content)
	}
	result, err := Parse("", content, append(opts, Entrypoint(string(s.entrypoint)), GlobalStore(substitutionKey, s))...)
	if err != nil {
		return nil, err
	}
	r, ok := result.([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected type of content after parsing elements: '%T'", result)
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("parsed content:\n%s", spew.Sdump(r))
	}
	return r, nil
}

// replaces the AttributeSubstitution by their actual values.
// TODO: returns `true` if at least one AttributeSubstitution was found (whatever its replacement)?
func replaceAttributeRefsInAttributes(ctx *ParseContext, b types.BlockWithAttributes) (bool, error) {
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("replacing attribute refs in %s", spew.Sdump(b.GetAttributes()))
	// }
	found := false
	for key, value := range b.GetAttributes() {
		switch value := value.(type) {
		case []interface{}: // multi-value attributes
			value, f, err := replaceAttributeRefsInElements(ctx, value)
			if err != nil {
				return false, err
			}
			found = found || f
			// (a bit hack-ish): do not merge values when the attribute is `roles` or `options`
			switch key {
			case types.AttrRoles, types.AttrOptions:
				b.GetAttributes()[key] = value
			default:
				b.GetAttributes()[key] = types.Reduce(value)
			}
		default: // single-value attributes
			value, f, err := replaceAttributeRefsInElement(ctx, value)
			if err != nil {
				return false, err
			}
			found = found || f
			b.GetAttributes()[key] = types.Reduce(value)
		}
	}
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("replaced attribute refs: %s", spew.Sdump(b.GetAttributes()))
	// }
	return found, nil
}

// replaces the AttributeSubstitution or Counter substitution with its actual value, recursively if the given `element`
// is a slice
func replaceAttributeRefsInElement(ctx *ParseContext, element interface{}) (interface{}, bool, error) {
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("replacing attribute references in element of type '%T'", element)
	// }
	switch e := element.(type) {
	case []interface{}:
		return replaceAttributeRefsInElements(ctx, e)
	case *types.AttributeSubstitution:
		s, found, err := ctx.attributes.getAsString(e.Name)
		if err != nil {
			return nil, false, err
		} else if !found {
			log.Warnf("unable to find attribute '%s'", e.Name)
			return e, false, nil
		}
		return &types.StringElement{
			Content: s,
		}, true, nil
	case types.CounterSubstitution:
		return counterToStringElement(ctx, e)
	case types.BlockWithElements:
		// replace AttributeSubstitutions on attributes
		found := false
		f, err := replaceAttributeRefsInAttributes(ctx, e)
		if err != nil {
			return nil, false, err
		}
		found = found || f
		// replace AttributeSubstitutions on nested elements
		elements, f, err := replaceAttributeRefsInElements(ctx, e.GetElements())
		if err != nil {
			return nil, false, err
		}
		found = found || f
		// elements = types.Merge(elements)
		if err := e.SetElements(elements); err != nil {
			return nil, false, errors.Wrapf(err, "failed to replace attribute references in block of type '%T'", e)
		}
		return e, found, nil
	case types.BlockWithLocation:
		found := false
		// replace AttributeSubstitutions on attributes
		f, err := replaceAttributeRefsInAttributes(ctx, e)
		if err != nil {
			return nil, false, err
		}
		found = found || f
		// replace AttributeSubstitutions on embedded location
		f, err = replaceAttributeRefsInLocation(ctx, e)
		if err != nil {
			return nil, false, err
		}
		found = found || f
		return e, found, nil
	case types.BlockWithAttributes:
		// replace AttributeSubstitutions on attributes

		found, err := replaceAttributeRefsInAttributes(ctx, e)
		return e, found, err
	default:
		return e, false, nil
	}
}

func replaceAttributeRefsInLocation(ctx *ParseContext, b types.BlockWithLocation) (bool, error) {
	if b.GetLocation() == nil {
		return false, nil
	}
	path, found, err := replaceAttributeRefsInElements(ctx, b.GetLocation().Path)
	if err != nil {
		return false, err
	}
	b.GetLocation().Path = path
	return found, nil
}

func replaceAttributeRefsInElements(ctx *ParseContext, elements []interface{}) ([]interface{}, bool, error) {
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("replacing attribute refs in elements:\n%s", spew.Sdump(elements))
	// }
	result := make([]interface{}, len(elements)) // maximum capacity should exceed initial input
	found := false
	for i, element := range elements {
		element, f, err := replaceAttributeRefsInElement(ctx, element)
		if err != nil {
			return nil, false, err
		}
		found = found || f
		result[i] = element
	}
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("replaced attribute refs in elements:\n%s", spew.Sdump(result))
	// }
	return result, found, nil
}

func counterToStringElement(ctx *ParseContext, c types.CounterSubstitution) (*types.StringElement, bool, error) {
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
			return &types.StringElement{Content: ""}, true, nil
		}
		return &types.StringElement{
			Content: strconv.Itoa(counter),
		}, true, nil
	case rune:
		if increment {
			counter++
		}
		ctx.counters[c.Name] = counter
		if c.Hidden {
			// return empty string facilitates merging
			return &types.StringElement{Content: ""}, true, nil
		}
		return &types.StringElement{
			Content: string(counter),
		}, true, nil
	default:
		return nil, false, fmt.Errorf("unexpected type of counter value: '%T'", counter)
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
		// also check nested elements (eg, in QuotedText, etc.)
		// for each element, check *all* interfaces to see if there's a need to replace the placeholders
		if e, ok := e.(types.BlockWithElements); ok {
			elmts, err := p.restoreInBlockWithElements(e)
			if err != nil {
				return nil, err
			}
			elements[i] = elmts
		}
		if e, ok := e.(types.BlockWithAttributes); ok {
			e, err := p.restoreInBlockWithAttributes(e)
			if err != nil {
				return nil, err
			}
			elements[i] = e
		}
		if e, ok := e.(types.BlockWithLocation); ok {
			e, err := p.restoreInBlockWithLocation(e)
			if err != nil {
				return nil, err
			}
			elements[i] = e
		}
	}
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("restored elements:\n%v", spew.Sdump(elements))
	// }
	return elements, nil
}

func (p *placeholders) restoreInBlockWithElements(e types.BlockWithElements) (types.BlockWithElements, error) {
	elements := e.GetElements()
	result := make([]interface{}, len(elements))
	for i, e := range elements {
		switch e := e.(type) {
		case *types.ElementPlaceHolder:
			result[i] = p.elements[e.Ref]
		case []interface{}:
			e, err := p.restore(e)
			if err != nil {
				return nil, err
			}
			result[i] = e
		default:
			result[i] = e
		}
	}
	if err := e.SetElements(result); err != nil {
		return nil, err
	}
	return e, nil
}

func (p *placeholders) restoreInBlockWithLocation(e types.BlockWithLocation) (types.BlockWithLocation, error) {
	location := e.GetLocation()
	result := make([]interface{}, len(location.Path))
	for i, e := range location.Path {
		switch e := e.(type) {
		case *types.ElementPlaceHolder:
			result[i] = p.elements[e.Ref]
		case []interface{}:
			e, err := p.restore(e)
			if err != nil {
				return nil, err
			}
			result[i] = e
		default:
			result[i] = e
		}
	}
	e.GetLocation().Path = result
	return e, nil
}

func (p *placeholders) restoreInBlockWithAttributes(e types.BlockWithAttributes) (types.BlockWithAttributes, error) {
	attrs := e.GetAttributes()
	for key, value := range attrs {
		switch value := value.(type) {
		case *types.ElementPlaceHolder:
			attrs[key] = p.elements[value.Ref]
		case []interface{}:
			value, err := p.restore(value)
			if err != nil {
				return nil, err
			}
			attrs[key] = value
		}
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("restored placeholders in attributes:\n%s", spew.Sdump(attrs))
	}
	return e, nil
}

func serialize(content interface{}) ([]byte, *placeholders) {
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("serializing:\n%v", spew.Sdump(content))
	// }
	placeholders := newPlaceholders()
	result := bytes.NewBuffer(nil)
	switch content := content.(type) {
	case string: // for attributes with simple (string) values
		result.WriteString(content)
	case []interface{}: // for paragraph lines, attributes with complex values, etc.
		for i, element := range content {
			switch element := element.(type) {
			case types.RawContent:
				result.WriteString(string(element))
			case types.RawLine:
				result.WriteString(string(element))
				// add `\n` unless the next element is a single-line comment (otherwise we'll have to deal with an extra trailing `\n` afterwards)
				if i < len(content)-1 {
					if _, ok := content[i+1].(*types.SingleLineComment); !ok {
						result.WriteString("\n")
					}
				}
			case *types.SingleLineComment:
				// replace with placeholder
				p := placeholders.add(element)
				result.WriteString(p.String())
				// add `\n` unless the next element is a single-line comment
				if i < len(content)-1 {
					if _, ok := content[i+1].(*types.SingleLineComment); !ok {
						result.WriteString("\n")
					}
				}
			case *types.StringElement:
				result.WriteString(element.Content)
			default:
				// replace with placeholder
				p := placeholders.add(element)
				result.WriteString(p.String())
			}
		}
	}
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("serialized lines: '%s'\nplaceholders: %v", result.Bytes(), spew.Sdump(placeholders.elements))
	// }
	return result.Bytes(), placeholders
}

func (c *current) lookupCurrentSubstitution() (*substitution, bool) {
	s, found := c.globalStore[substitutionKey].(*substitution)
	return s, found
}

func (c *current) isSubstitutionEnabled(k substitutionKind) (bool, error) {
	s, found := c.lookupCurrentSubstitution()
	if !found {
		// log.Debugf("no current substitution, so assuming '%s' not enabled", k)
		return false, nil // TODO: should return `true`, at least for `attributes`?
	}
	enabled, found := s.rules[k]
	if !found {
		// log.Debugf("substitution '%s' not configured in '%s'", k, s.entrypoint)
		return false, nil
	}
	// log.Debugf("substitution '%s' enabled in rule '%s': %t", string(k), spew.Sdump(s.rules), enabled)
	return enabled, nil
}

type substitutionKind string

const (
	// substitutionKey the key in which substitutions are stored in the parser's GlobalStore
	substitutionKey string = "current_substitution"

	// Attributes the "attributes" substitution
	Attributes substitutionKind = "attributes"
	// Callouts the "callouts" substitution
	Callouts substitutionKind = "callouts"
	// InlinePassthroughs the "inline_passthrough" substitution
	InlinePassthroughs substitutionKind = "inline_passthrough" // nolint: gosec
	// Macros the "macros" substitution
	Macros substitutionKind = "macros"
	// None the "none" substitution
	None substitutionKind = "none"
	// PostReplacements the "post_replacements" substitution
	PostReplacements substitutionKind = "post_replacements"
	// Quotes the "quotes" substitution
	Quotes substitutionKind = "quotes"
	// Replacements the "replacements" substitution
	Replacements substitutionKind = "replacements"
	// SpecialCharacters the "specialchars" substitution
	SpecialCharacters substitutionKind = "specialchars"
)
