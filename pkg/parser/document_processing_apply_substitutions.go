package parser

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// ----------------------------------------------------------------------------
// Substitutions
// ----------------------------------------------------------------------------

// ApplySubstitutions applies all the substitutions on delimited blocks, standalone paragraphs and paragraphs
// in continued list items, and then attribute substitutions, and as a result returns a `DraftDocument`.
func ApplySubstitutions(rawDoc types.RawDocument, config configuration.Configuration, options ...Option) (types.DraftDocument, error) {
	attrs := types.AttributesWithOverrides{
		Content:   types.Attributes{},
		Overrides: config.AttributeOverrides,
		Counters:  map[string]interface{}{},
	}
	// also, add all front-matter key/values
	attrs.Add(rawDoc.FrontMatter.Content)
	// also, add all AttributeDeclaration at the top of the document
	attrs.Add(rawDoc.Attributes())

	blocks, err := applySubstitutions(rawDoc.Blocks, attrs, options...)
	if err != nil {
		return types.DraftDocument{}, err
	}
	if len(blocks) == 0 {
		blocks = nil // avoid carrying empty slice
	}
	return types.DraftDocument{
		Attributes:  attrs.All(),
		FrontMatter: rawDoc.FrontMatter,
		Blocks:      blocks,
	}, nil
}

// applySubstitutions applies the substitutions on paragraphs and delimited blocks (including when in continued list elements)
func applySubstitutions(elements []interface{}, attrs types.AttributesWithOverrides, options ...Option) ([]interface{}, error) {
	if len(elements) == 0 {
		return nil, nil
	}
	result := []interface{}{}
	for _, e := range elements {
		switch e := e.(type) {
		case types.ContinuedListItemElement:
			r, err := applySubstitutions([]interface{}{e.Element}, attrs, options...)
			if err != nil {
				return nil, err
			}
			e.Element = r[0]
			result = append(result, e)
		case types.Paragraph:
			p, err := applySubstitutionsOnParagraph(e, attrs, options...)
			if err != nil {
				return nil, err
			}
			result = append(result, p)
		case types.ImageBlock:
			i, err := applySubstitutionsOnImageBlock(e, attrs, options...)
			if err != nil {
				return nil, err
			}
			result = append(result, i)
		case types.Section:
			s, err := applySubstitutionsOnSection(e, attrs, options...)
			if err != nil {
				return nil, err
			}
			result = append(result, s)
		case types.DelimitedBlock:
			b, err := applySubstitutionsOnDelimitedBlock(e, attrs, options...)
			if err != nil {
				return nil, err
			}
			result = append(result, b)
		default:
			// no support for element substitution here
			// so let's proceed with attribute substitutions
			e, err := applyAttributeSubstitutionsOnElement(e, attrs)
			if err != nil {
				return nil, err
			}
			result = append(result, e)
		}
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debug("after all substitutions:")
		spew.Fdump(log.StandardLogger().Out, result)
	}
	return result, nil
}

// ----------------------------------------------------------------------------
// Delimited Block substitutions
// ----------------------------------------------------------------------------

// applySubstitutionsOnDelimitedBlock parses the given raw elements, depending on the given substitutions to apply
// May return the elements unchanged, or convert the elements to a source doc and parse with a custom entrypoint
func applySubstitutionsOnDelimitedBlock(b types.DelimitedBlock, attrs types.AttributesWithOverrides, options ...Option) (types.DelimitedBlock, error) {
	switch b.Kind {
	case types.Example, types.Quote, types.Sidebar:
		return applyNormalBlockSubstitutions(b, attrs, options...)
	case types.Fenced, types.Listing, types.Literal, types.Source, types.Passthrough:
		return applyVerbatimBlockSubstitutions(b, attrs, options...)
	case types.Verse:
		return applyVerseBlockSubstitutions(b, attrs, options...)
	case types.MarkdownQuote:
		return applyMarkdownQuoteBlockSubstitutions(b, attrs, options...)
	case types.Comment:
		return applyCommentBlockSubstitutions(b, attrs, options...)
	}
	return b, fmt.Errorf("unsupported block type: '%s", string(b.Kind))
}

func applyNormalBlockSubstitutions(b types.DelimitedBlock, attrs types.AttributesWithOverrides, options ...Option) (types.DelimitedBlock, error) {
	funcs := []elementsSubstitution{}
	subs, _ := b.Attributes.GetAsString(types.AttrSubstitutions)
	for _, s := range strings.Split(subs, ",") {
		switch s {
		case "", "normal":
			funcs = append(funcs,
				substituteInlinePassthrough,
				substituteSpecialCharacters,
				substituteQuotedTexts,
				substituteAttributes,
				substituteReplacements,
				substituteInlineMacros,
				substitutePostReplacements)
		case "specialcharacters", "specialchars":
			funcs = append(funcs, substituteSpecialCharacters)
		case "quotes":
			funcs = append(funcs, substituteQuotedTexts)
		case "attributes":
			funcs = append(funcs, substituteAttributes)
		case "macros":
			funcs = append(funcs, substituteInlineMacros)
		case "replacements":
			funcs = append(funcs, substituteReplacements)
		case "post_replacements":
			funcs = append(funcs, substitutePostReplacements)
		case "none":
			funcs = append(funcs, substituteNone)
		default:
			return types.DelimitedBlock{}, fmt.Errorf("unsupported substitution: '%s", s)
		}
	}
	funcs = append(funcs, splitLines)

	// first, parse the raw lines to extract "sub blocks"
	blocks, err := substituteNormalBlocks(b.Elements, attrs, options...)
	if err != nil {
		return types.DelimitedBlock{}, err
	}
	// apply all the substitutions on blocks that need to be processed
	for i, block := range blocks {
		switch block := block.(type) {
		case types.Paragraph:
			for _, sub := range funcs {
				if block.Lines, err = sub(block.Lines, attrs, options...); err != nil {
					return types.DelimitedBlock{}, err
				}
			}
			blocks[i] = block
		case types.DelimitedBlock:
			for _, sub := range funcs {
				if block.Elements, err = sub(block.Elements, attrs, options...); err != nil {
					return types.DelimitedBlock{}, err
				}
			}
			blocks[i] = block
		default:
			// do nothing
		}
	}
	b.Elements = blocks
	return b, nil
}

func applyVerbatimBlockSubstitutions(b types.DelimitedBlock, attrs types.AttributesWithOverrides, options ...Option) (types.DelimitedBlock, error) {
	funcs := []elementsSubstitution{}
	subs, _ := b.Attributes.GetAsString(types.AttrSubstitutions)
	for _, s := range strings.Split(subs, ",") {
		switch s {
		case "":
			funcs = append(funcs, substituteCallouts, substituteSpecialCharacters)
		case "normal":
			funcs = append(funcs,
				substituteInlinePassthrough,
				substituteSpecialCharacters,
				substituteQuotedTexts,
				substituteAttributes,
				substituteReplacements,
				substituteInlineMacros,
				substitutePostReplacements)
		case "specialcharacters", "specialchars":
			funcs = append(funcs, substituteSpecialCharacters)
		case "quotes":
			funcs = append(funcs, substituteQuotedTexts)
		case "attributes":
			funcs = append(funcs, substituteAttributes)
		case "macros":
			funcs = append(funcs, substituteInlineMacros)
		case "replacements":
			funcs = append(funcs, substituteReplacements)
		case "post_replacements":
			funcs = append(funcs, substitutePostReplacements)
		case "none":
			funcs = append(funcs, substituteNone)
		default:
			return types.DelimitedBlock{}, fmt.Errorf("unsupported substitution: '%s", s)
		}
	}
	funcs = append(funcs, splitLines)
	// apply all the substitutions
	var err error
	for _, sub := range funcs {
		if b.Elements, err = sub(b.Elements, attrs, options...); err != nil {
			return types.DelimitedBlock{}, err
		}
	}
	return b, nil
}

func applyVerseBlockSubstitutions(b types.DelimitedBlock, attrs types.AttributesWithOverrides, options ...Option) (types.DelimitedBlock, error) {
	funcs := []elementsSubstitution{}
	subs, _ := b.Attributes.GetAsString(types.AttrSubstitutions)
	for _, s := range strings.Split(subs, ",") {
		switch s {
		case "", "normal":
			funcs = append(funcs,
				substituteInlinePassthrough,
				substituteSpecialCharacters,
				substituteQuotedTexts,
				substituteAttributes,
				substituteReplacements,
				substituteVerseMacros,
				substitutePostReplacements,
			)
		case "specialcharacters", "specialchars":
			funcs = append(funcs, substituteSpecialCharacters)
		case "quotes":
			funcs = append(funcs, substituteQuotedTexts)
		case "attributes":
			funcs = append(funcs, substituteAttributes)
		case "macros":
			funcs = append(funcs, substituteVerseMacros)
		case "replacements":
			funcs = append(funcs, substituteReplacements)
		case "post_replacements":
			funcs = append(funcs, substitutePostReplacements)
		case "none":
			funcs = append(funcs, substituteNone)
		default:
			return types.DelimitedBlock{}, fmt.Errorf("unsupported substitution: '%s", s)
		}
	}
	funcs = append(funcs, splitLines)
	// apply all the substitutions
	var err error
	for _, sub := range funcs {
		if b.Elements, err = sub(b.Elements, attrs, options...); err != nil {
			return types.DelimitedBlock{}, err
		}
	}
	return b, nil
}

func applyMarkdownQuoteBlockSubstitutions(b types.DelimitedBlock, attrs types.AttributesWithOverrides, options ...Option) (types.DelimitedBlock, error) {
	funcs := []elementsSubstitution{
		substituteInlinePassthrough,
		substituteSpecialCharacters,
		substituteQuotedTexts,
		substituteAttributes,
		substituteReplacements,
		substituteMarkdownQuoteMacros,
		substitutePostReplacements}
	// attempt to extract the block attributions
	var author string
	if b.Elements, author = extractMarkdownQuoteAttribution(b.Elements); author != "" {
		if b.Attributes == nil {
			b.Attributes = types.Attributes{}
		}
		b.Attributes.Set(types.AttrQuoteAuthor, author)
	}
	if len(b.Elements) == 0 { // no more line to parse after extracting the author
		return b, nil
	}
	// apply all the substitutions
	var err error
	for _, sub := range funcs {
		if b.Elements, err = sub(b.Elements, attrs, options...); err != nil {
			return types.DelimitedBlock{}, err
		}
	}
	return b, nil
}

func applyCommentBlockSubstitutions(b types.DelimitedBlock, attrs types.AttributesWithOverrides, options ...Option) (types.DelimitedBlock, error) {
	funcs := []elementsSubstitution{substituteNone, splitLines}
	var err error
	for _, sub := range funcs {
		if b.Elements, err = sub(b.Elements, attrs, options...); err != nil {
			return types.DelimitedBlock{}, err
		}
	}
	return b, nil
}

func extractMarkdownQuoteAttribution(elements []interface{}) ([]interface{}, string) {
	log.Debug("extracting attribution on markdown block quote")
	// first, check if last line is an attribution (author)
	if len(elements) == 0 {
		return elements, ""
	}
	if l, ok := elements[len(elements)-1].(types.RawLine); ok {
		a, err := ParseReader("", strings.NewReader(l.Content), Entrypoint("MarkdownQuoteAttribution"))
		// assume that the last line is not an author attribution if an error occurred
		if err != nil {
			return elements, ""
		}
		if a, ok := a.(string); ok {
			log.Debugf("found attribution in markdown block: '%s'", a)
			return elements[:len(elements)-1], a
		}
	}
	return elements, ""
}

// ----------------------------------------------------------------------------
// Section substitutions
// ----------------------------------------------------------------------------

// applies the elements and attributes substitutions on the given section title.
func applySubstitutionsOnSection(s types.Section, attrs types.AttributesWithOverrides, options ...Option) (types.Section, error) {
	elements := s.Title
	subs := []elementsSubstitution{
		substituteInlinePassthrough,
		substituteSpecialCharacters,
		substituteQuotedTexts,
		substituteAttributes,
		substituteReplacements,
		substituteInlineMacros,
		substitutePostReplacements,
	}
	var err error
	for _, sub := range subs {
		if elements, err = sub(elements, attrs, options...); err != nil {
			return types.Section{}, err
		}
	}
	s.Title = elements
	if s, err = s.ResolveID(attrs); err != nil {
		return types.Section{}, err
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("section title after substitution:")
		spew.Fdump(log.StandardLogger().Out, s.Title)
	}
	return s, nil
}

// ----------------------------------------------------------------------------
// Image Block substitutions
// ----------------------------------------------------------------------------

// applies the elements and attributes substitutions on the given image block.
func applySubstitutionsOnImageBlock(b types.ImageBlock, attrs types.AttributesWithOverrides, options ...Option) (types.ImageBlock, error) {
	elements := b.Location.Path
	subs := []elementsSubstitution{substituteAttributes}
	var err error
	for _, sub := range subs {
		if elements, err = sub(elements, attrs, options...); err != nil {
			return types.ImageBlock{}, err
		}
	}
	b.Location.Path = elements
	b.Location = b.Location.WithPathPrefix(attrs.GetAsStringWithDefault("imagesdir", ""))
	if !b.Attributes.Has(types.AttrImageAlt) {
		alt := filepath.Base(b.Location.Stringify())
		ext := filepath.Ext(alt)
		alt = alt[0 : len(alt)-len(ext)]
		b.Attributes = b.Attributes.Set(types.AttrImageAlt, alt)
	}

	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("image block after substitution:")
		spew.Fdump(log.StandardLogger().Out, b)
	}
	return b, nil
}

// ----------------------------------------------------------------------------
// Paragraph substitutions
// ----------------------------------------------------------------------------

// applies the elements and attributes substitutions on the given paragraph.
// Attributes substitution is triggered only if there is no specific substitution or if the `attributes` substitution is explicitly set.
func applySubstitutionsOnParagraph(p types.Paragraph, attrs types.AttributesWithOverrides, options ...Option) (types.Paragraph, error) {
	subs, err := paragraphSubstitutions(p)
	if err != nil {
		return types.Paragraph{}, err
	}
	// apply all the configured substitutions
	for _, sub := range subs {
		var err error
		if p.Lines, err = sub(p.Lines, attrs, options...); err != nil {
			return types.Paragraph{}, err
		}
	}
	return p, nil
}

// paragraphSubstitutions returns the substitution funcs to apply on the given paragraph `p`
// otherwise, returns a default substitution which will ultemately fail
func paragraphSubstitutions(p types.Paragraph) ([]elementsSubstitution, error) {
	subs, _ := p.Attributes.GetAsString(types.AttrSubstitutions)
	// log.Debugf("determining substitutions for '%s' on a paragraph", subs)
	funcs := []elementsSubstitution{}
	for _, s := range strings.Split(subs, ",") {
		switch s {
		case "specialcharacters", "specialchars":
			funcs = append(funcs, substituteSpecialCharacters)
		case "quotes":
			funcs = append(funcs, substituteQuotedTexts)
		case "attributes":
			funcs = append(funcs, substituteAttributes)
		case "macros":
			funcs = append(funcs, substituteInlineMacros)
		case "replacements":
			funcs = append(funcs, substituteReplacements)
		case "post_replacements":
			funcs = append(funcs, substitutePostReplacements)
		case "", "normal":
			funcs = append(funcs,
				substituteInlinePassthrough,
				substituteSpecialCharacters,
				substituteQuotedTexts,
				substituteAttributes,
				substituteReplacements,
				substituteInlineMacros,
				substitutePostReplacements,
			)
		case "none":
			funcs = append(funcs, substituteNone)
		default:
			return nil, fmt.Errorf("unsupported substitution: '%s", s)
		}
	}
	funcs = append(funcs, splitLines)
	return funcs, nil
}

// ----------------------------------------------------------------------------
// Individual substitution funcs
// ----------------------------------------------------------------------------

// includes a call to `elementsSubstitution` with some post-processing on the result
var substituteAttributes = func(elements []interface{}, attrs types.AttributesWithOverrides, options ...Option) ([]interface{}, error) {
	elements, err := newElementsSubstitution("AttributeSubs", "AttributeSubs")(elements, attrs, options...)
	if err != nil {
		return nil, err
	}
	for i, element := range elements {
		element, err := applyAttributeSubstitutionsOnElement(element, attrs)
		if err != nil {
			return nil, err
		}
		elements[i] = element
	}
	elements = types.Merge(elements)
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("applied the 'attributes' substitution")
		spew.Fdump(log.StandardLogger().Out, elements)
	}
	return elements, nil
}

var (
	substituteInlinePassthrough   = newElementsSubstitution("InlinePassthroughSubs", "InlinePassthroughSubs")
	substituteSpecialCharacters   = newElementsSubstitution("SpecialCharacterSubs", "SpecialCharacterSubs")
	substituteQuotedTexts         = newElementsSubstitution("QuotedTextSubs", "QuotedTextSubs")
	substituteReplacements        = newElementsSubstitution("ReplacementSubs", "ReplacementSubs")
	substituteInlineMacros        = newElementsSubstitution("InlineMacroSubs", "InlineMacroSubs")
	substituteNormalBlocks        = newElementsSubstitution("NormalBlocks", "NormalBlocks")
	substituteVerseMacros         = newElementsSubstitution("VerseMacroSubs", "VerseMacroSubs")
	substituteMarkdownQuoteMacros = newElementsSubstitution("MarkdownQuoteMacroSubs", "MarkdownQuoteLine")
	substitutePostReplacements    = newElementsSubstitution("PostReplacementSubs", "PostReplacementSubs")
	substituteNone                = newElementsSubstitution("NoneSubs", "NoneSubs") // TODO: no need for placeholder support here?
	substituteCallouts            = newElementsSubstitution("CalloutSubs", "CalloutSubs")
)

type elementsSubstitution func(elements []interface{}, attrs types.AttributesWithOverrides, options ...Option) ([]interface{}, error)

func newElementsSubstitution(contentRuleName, placeholderRuleName string) elementsSubstitution {
	return func(elements []interface{}, attrs types.AttributesWithOverrides, options ...Option) ([]interface{}, error) {
		log.Debugf("applying the '%s' rule on elements", contentRuleName)
		placeholders := newPlaceHolders()
		s := serializeElementsWithPlaceHolders(elements, placeholders)
		options = append(options, GlobalStore("imagesdir", attrs.GetAsStringWithDefault("imagesdir", ""))) // TODO: define a const for "imagesdir"
		// process placeholder content (eg: quoted text may contain an inline link)
		for ref, placeholder := range placeholders.elements {
			switch placeholder := placeholder.(type) { // TODO: create `PlaceHolder` interface?
			case types.QuotedString:
				var err error
				if placeholder.Elements, err = parserPlaceHolderElements(placeholder.Elements, append(options, Entrypoint(placeholderRuleName))...); err != nil {
					return nil, err
				}
				placeholders.elements[ref] = placeholder
			case types.QuotedText:
				var err error
				if placeholder.Elements, err = parserPlaceHolderElements(placeholder.Elements, append(options, Entrypoint(placeholderRuleName))...); err != nil {
					return nil, err
				}
				placeholders.elements[ref] = placeholder
			}
		}
		result := make([]interface{}, 0, len(elements))
		elmts, err := parseContent("", s, append(options, Entrypoint(contentRuleName))...)
		if err != nil {
			return nil, err
		}
		elmts = restoreElements(elmts, placeholders)
		result = append(result, elmts...)
		if log.IsLevelEnabled(log.DebugLevel) {
			log.Debugf("applied the '%s' rule:", contentRuleName)
			spew.Fdump(log.StandardLogger().Out, result)
		}
		return result, nil
	}
}

func parserPlaceHolderElements(elements []interface{}, options ...Option) ([]interface{}, error) {
	result := make([]interface{}, 0, len(elements)) // default capacity (but may not be enough)
	for _, element := range elements {
		switch element := element.(type) {
		case types.StringElement:
			elmts, err := parseContent("", element.Content, options...)
			if err != nil {
				return nil, err
			}
			result = append(result, elmts...)
		default:
			result = append(result, element)
		}
	}
	return result, nil
}

func parseContent(filename string, content string, options ...Option) ([]interface{}, error) {
	// log.Debugf("parsing content '%s'", content)
	result, err := ParseReader(filename, strings.NewReader(content), options...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to parse '%s'", content)
	}
	if result, ok := result.([]interface{}); ok {
		return types.Merge(result), nil
	}
	return []interface{}{result}, nil
}

// replace the placeholders with their original element in the given elements
func restoreElements(elmts []interface{}, placeholders *placeholders) []interface{} {
	// skip if there's nothing to restore
	if len(placeholders.elements) == 0 {
		return elmts
	}
	for i, elmt := range elmts {
		switch elmt := elmt.(type) {
		case types.ElementPlaceHolder:
			elmts[i] = placeholders.elements[elmt.Ref]
		case types.Paragraph:
			elmt.Lines = restoreElements(elmt.Lines, placeholders)
			elmts[i] = elmt
		case types.InlineLink: // TODO: use an interface and implement the `restoreElements` func on these types, instead
			elmt.Location.Path = restoreElements(elmt.Location.Path, placeholders)
			elmt.Attributes = restoreAttributes(elmt.Attributes, placeholders)
			elmts[i] = elmt
		case types.QuotedText:
			elmt.Elements = restoreElements(elmt.Elements, placeholders)
			elmt.Attributes = restoreAttributes(elmt.Attributes, placeholders)
			elmts[i] = elmt
		case types.QuotedString:
			elmt.Elements = restoreElements(elmt.Elements, placeholders)
			elmts[i] = elmt
		case types.IndexTerm:
			elmt.Term = restoreElements(elmt.Term, placeholders)
			elmts[i] = elmt
		case types.ExternalCrossReference:
			elmt.Label = restoreElements(elmt.Label, placeholders)
			elmts[i] = elmt
		case types.Footnote:
			elmt.Elements = restoreElements(elmt.Elements, placeholders)
			elmts[i] = elmt
		case types.ElementRole:
			elmts[i] = types.ElementRole(restoreElements(elmt, placeholders))
		case []interface{}:
			elmts[i] = restoreElements(elmt, placeholders)
		default:
			// do nothing, keep elmt as-is
		}
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("restored elements")
		spew.Fdump(log.StandardLogger().Out, elmts)
	}
	return elmts
}

// replace the placeholders with their original element in the given attributes
func restoreAttributes(attrs types.Attributes, placeholders *placeholders) types.Attributes {
	for key, value := range attrs {
		switch value := value.(type) {
		case types.ElementPlaceHolder:
			attrs[key] = placeholders.elements[value.Ref]
		case types.ElementRole:
			attrs[key] = types.ElementRole(restoreElements(value, placeholders))
		case []interface{}:
			attrs[key] = restoreElements(value, placeholders)
		}
	}
	return attrs
}

func splitLines(elements []interface{}, _ types.AttributesWithOverrides, _ ...Option) ([]interface{}, error) {
	// after processing all the elements, we want to split them in separate lines again, to retain the initial input format
	lines := []interface{}{}
	line := []interface{}{}
	for _, element := range types.Merge(elements) {
		switch element := element.(type) {
		case types.StringElement:
			// if content has line breaks, then split in multiple lines
			if split := strings.Split(element.Content, "\n"); len(split) > 1 {
				for i, s := range split {
					if len(s) > 0 { // no need to insert empty StringElements
						line = append(line, types.StringElement{Content: s})
					}
					if i < len(split)-1 {
						lines = append(lines, line)
						line = []interface{}{} // reset for the next line, except for the last item
					}
				}
			} else {
				line = append(line, element)
			}
		case types.SingleLineComment: // single-line comments are on their own lines
			if len(line) > 0 {
				lines = append(lines, line)
			}
			lines = append(lines, []interface{}{element})
			line = []interface{}{} // reset for the next line
		default:
			line = append(line, element)
		}
	}
	if len(line) > 0 { // don't forget the last line (if applicable)
		lines = append(lines, line)
	}
	return lines, nil
}

type placeholders struct {
	seq      int
	elements map[string]interface{}
}

func newPlaceHolders() *placeholders {
	return &placeholders{
		seq:      0,
		elements: map[string]interface{}{},
	}
}
func (p *placeholders) add(element interface{}) types.ElementPlaceHolder {
	p.seq++
	p.elements[strconv.Itoa(p.seq)] = element
	return types.ElementPlaceHolder{
		Ref: strconv.Itoa(p.seq),
	}

}

func serializeElementsWithPlaceHolders(elements []interface{}, placeholders *placeholders) string {
	result := strings.Builder{}
	for i, e := range elements {
		// log.Debugf("serializing element of type '%T'", e)
		switch e := e.(type) {
		case []interface{}:
			r := serializeElementsWithPlaceHolders(e, placeholders)
			result.WriteString(r)
			if i < len(elements)-1 {
				result.WriteString("\n")
			}
		case types.RawLine:
			result.WriteString(e.Content)
			if i < len(elements)-1 {
				result.WriteString("\n")
			}
		case types.StringElement:
			result.WriteString(e.Content)
		default:
			// replace with placeholder
			p := placeholders.add(e)
			result.WriteString(p.String())
		}
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		spew.Fdump(log.StandardLogger().Out, result.String())
	}
	return result.String()
}

// ----------------------------------------------------------------------------
// Attribute substitutions
// ----------------------------------------------------------------------------

func applyAttributeSubstitutionsOnElements(elements []interface{}, attrs types.AttributesWithOverrides) ([]interface{}, error) {
	result := make([]interface{}, 0, len(elements)) // maximum capacity should exceed initial input
	for _, element := range elements {
		e, err := applyAttributeSubstitutionsOnElement(element, attrs)
		if err != nil {
			return nil, err
		}
		result = append(result, e)
	}
	result = types.Merge(result)
	return result, nil
}

func applyAttributeSubstitutionsOnElement(element interface{}, attrs types.AttributesWithOverrides) (interface{}, error) {
	// log.Debugf("applying attribute substitution on element of type '%T'", element)
	var err error
	switch e := element.(type) {
	case types.Paragraph:
		for i, l := range e.Lines {
			l, err := applyAttributeSubstitutionsOnElement(l, attrs)
			if err != nil {
				return nil, err
			}
			e.Lines[i] = l
		}
		return e, nil
	case []interface{}:
		return applyAttributeSubstitutionsOnElements(e, attrs)
	case types.AttributeDeclaration:
		attrs.Set(e.Name, e.Value)
		return e, nil
	case types.AttributeReset:
		attrs.Set(e.Name, nil) // This allows us to test for a reset vs. undefined.
		return e, nil
	case types.AttributeSubstitution:
		if value, ok := attrs.GetAsString(e.Name); ok {
			return types.StringElement{
				Content: value,
			}, nil
		}
		log.Warnf("unable to find attribute '%s'", e.Name)
		return types.StringElement{
			Content: "{" + e.Name + "}",
		}, nil
	case types.CounterSubstitution:
		return applyCounterSubstitution(e, attrs)
	case types.ImageBlock:
		e.Location.Path, err = applyAttributeSubstitutionsOnElements(e.Location.Path, attrs)
		return e, err
	case types.InlineImage:
		e.Location.Path, err = applyAttributeSubstitutionsOnElements(e.Location.Path, attrs)
		return e, err
	case types.InlineLink:
		e.Location.Path, err = applyAttributeSubstitutionsOnElements(e.Location.Path, attrs)
		return e, err
	case types.ExternalCrossReference:
		e.Location.Path, err = applyAttributeSubstitutionsOnElements(e.Location.Path, attrs)
		return e, err
	case types.Section:
		title, err := applyAttributeSubstitutionsOnElements(e.Title, attrs)
		if err != nil {
			return nil, err
		}
		e.Title = title
		return e.ResolveID(attrs)
	case types.OrderedListItem:
		e.Elements, err = applyAttributeSubstitutionsOnElements(e.Elements, attrs)
		return e, err
	case types.UnorderedListItem:
		e.Elements, err = applyAttributeSubstitutionsOnElements(e.Elements, attrs)
		return e, err
	case types.LabeledListItem:
		e.Elements, err = applyAttributeSubstitutionsOnElements(e.Elements, attrs)
		return e, err
	case types.QuotedText:
		e.Elements, err = applyAttributeSubstitutionsOnElements(e.Elements, attrs)
		return e, err
	case types.ContinuedListItemElement:
		e.Element, err = applyAttributeSubstitutionsOnElement(e.Element, attrs)
		return e, err
	case types.DelimitedBlock:
		e.Elements, err = applyAttributeSubstitutionsOnElements(e.Elements, attrs)
		return e, err
	default:
		return e, nil
	}
}

// applyCounterSubstitutions is called by applyAttributeSubstitutionsOnElement.  Unless there is an error with
// the element (the counter is the wrong type, which should never occur), it will return a `StringElement, true`
// (because we always either find the element, or allocate one), and `nil`.  On an error it will return `nil, false`,
// and the error.  The extra boolean here is to fit the calling expectations of our caller.  This function was
// factored out of a case from applyAttributeSubstitutionsOnElement in order to reduce the complexity of that
// function, but otherwise it should have no callers.
func applyCounterSubstitution(c types.CounterSubstitution, attrs types.AttributesWithOverrides) (interface{}, error) {
	log.Debugf("applying counter substitution for '%s'", c.Name)
	counter := attrs.Counters[c.Name]
	if counter == nil {
		counter = 0
	}
	increment := true
	if c.Value != nil {
		attrs.Counters[c.Name] = c.Value
		counter = c.Value
		increment = false
	}
	switch counter := counter.(type) {
	case int:
		if increment {
			counter++
		}
		attrs.Counters[c.Name] = counter
		if c.Hidden {
			// return empty string facilitates merging
			return types.StringElement{Content: ""}, nil
		}
		return types.StringElement{
			Content: strconv.Itoa(counter),
		}, nil
	case rune:
		if increment {
			counter++
		}
		attrs.Counters[c.Name] = counter
		if c.Hidden {
			// return empty string facilitates merging
			return types.StringElement{Content: ""}, nil
		}
		return types.StringElement{
			Content: string(counter),
		}, nil

	default:
		return nil, fmt.Errorf("invalid counter type %T", counter)
	}
}
