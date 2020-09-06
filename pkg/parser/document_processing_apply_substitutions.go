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

	blocks, err := applySubstitutions(rawDoc.Blocks, attrs, config, options...)
	if err != nil {
		return types.DraftDocument{}, err
	}
	// blocks, err = applyAttributeSubstitutions(blocks, attrs)
	// if err != nil {
	// 	return types.DraftDocument{}, err
	// }
	if len(blocks) == 0 {
		blocks = nil // avoid carrying empty slice
	}
	return types.DraftDocument{
		Attributes:  attrs.All(),
		FrontMatter: rawDoc.FrontMatter,
		Blocks:      blocks,
	}, nil
}

// ----------------------------------------------------------------------------
// Block substitutions
// ----------------------------------------------------------------------------

// applySubstitutions applies the substitutions on paragraphs and delimited blocks (including when in continued list elements)
func applySubstitutions(elements []interface{}, attrs types.AttributesWithOverrides, config configuration.Configuration, options ...Option) ([]interface{}, error) {
	if len(elements) == 0 {
		return nil, nil
	}
	result := []interface{}{}
	for _, e := range elements {
		switch e := e.(type) {
		case types.ContinuedListItemElement:
			r, err := applySubstitutions([]interface{}{e.Element}, attrs, config, options...)
			if err != nil {
				return nil, err
			}
			r[0], err = applyAttributeSubstitutionsOnElement(r[0], attrs)
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
			subs := delimitedBlockSubstitutions(e.Kind, config, options...)
			if err := applySubstitutionsOnDelimitedBlock(&e, subs); err != nil {
				return nil, err
			}
			r, err := applyAttributeSubstitutionsOnElement(e, attrs)
			if err != nil {
				return nil, err
			}
			result = append(result, r)
		default:
			// no support for element substitution here
			// so let's proceed with attribute substitutions
			e, err := applyAttributeSubstitutionsOnElement(e, attrs)
			if err != nil {
				return nil, err
			}
			// e = resolveLocationsOnElement(e, attrs)
			result = append(result, e)
		}
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debug("after all substitutions:")
		spew.Fdump(log.StandardLogger().Out, result)
	}
	return result, nil
}

func delimitedBlockSubstitutions(kind types.BlockKind, config configuration.Configuration, options ...Option) []blockSubstitution {
	switch kind {
	case types.Fenced, types.Listing, types.Literal, types.Source, types.Passthrough:
		// return the verbatim elements
		return []blockSubstitution{verbatimBlock(options...)}
	case types.Comment:
		return []blockSubstitution{none()}
	case types.Example, types.Quote, types.Sidebar:
		return []blockSubstitution{normalBlock(config, options...)}
	case types.Verse:
		return []blockSubstitution{verseBlock(config, options...)}
	case types.MarkdownQuote:
		return []blockSubstitution{markdownQuote(config, options...)}
	default:
		log.Errorf("unexpected kind of delimited block: '%s'. Will apply the 'none' substitution", kind)
		return []blockSubstitution{none()}
	}
}

// applySubstitutionsOnDelimitedBlock parses the given raw elements, depending on the given substitutions to apply
// May return the elements unchanged, or convert the elements to a source doc and parse with a custom entrypoint
func applySubstitutionsOnDelimitedBlock(b *types.DelimitedBlock, subs []blockSubstitution) error {
	log.Debug("applying delimited block substitutions")
	for _, sub := range subs {
		if err := sub(b); err != nil {
			return err
		}
	}
	return nil
}

type blockSubstitution func(b *types.DelimitedBlock) error

// performs all substitutions except for callouts
func normalBlock(config configuration.Configuration, options ...Option) blockSubstitution {
	return func(b *types.DelimitedBlock) error {
		log.Debugf("applying the 'normal' substitution on a delimited block content")
		raw, err := serializeBlock(b.Elements)
		if err != nil {
			return err
		}
		if b.Elements, err = parseContent(config.Filename, raw, append(options, Entrypoint("NormalBlockContentSubstitution"))...); err != nil {
			return err
		}
		// now, check if there are nested delimited blocks, in which case apply the same substitution recursively
		for i, e := range b.Elements {
			if d, ok := e.(types.DelimitedBlock); ok {
				subs := delimitedBlockSubstitutions(d.Kind, config, options...)
				if err := applySubstitutionsOnDelimitedBlock(&d, subs); err != nil {
					return err
				}
				b.Elements[i] = d // store back in the elements
			}
		}
		return err
	}
}

// performs all substitutions except for callouts and list items
func verseBlock(config configuration.Configuration, options ...Option) blockSubstitution {
	return func(b *types.DelimitedBlock) error {
		log.Debugf("applying the 'verse' substitution on a delimited block")
		raw, err := serializeBlock(b.Elements)
		if err != nil {
			return err
		}
		b.Elements, err = parseContent(config.Filename, raw, append(options, Entrypoint("VerseBlockContentSubstitution"))...)
		return err
	}
}

// replaces special characters and processes callouts
func verbatimBlock(options ...Option) blockSubstitution {
	return func(b *types.DelimitedBlock) error {
		log.Debugf("applying the 'verbatim' substitution on a delimited block")
		result := []interface{}{}
		for _, elmt := range b.Elements {
			switch elmt := elmt.(type) {
			case types.RawLine:
				elements, err := parseRawLine(elmt, append(options, Entrypoint("VerbatimContentSubstitution"))...)
				if err != nil {
					return errors.Wrapf(err, "failed to apply verbatim substitution on '%s'", elmt.Content)
				}
				result = append(result, elements...)
			default:
				result = append(result, elmt)
			}
		}
		b.Elements = result
		return nil
	}
}

func markdownQuote(config configuration.Configuration, options ...Option) blockSubstitution {
	return func(b *types.DelimitedBlock) error {
		log.Debugf("applying the 'normal' substitution on a markdown quote block")
		elements, author := extractQuoteBlockAttribution(b.Elements)
		if author != "" {
			if b.Attributes == nil {
				b.Attributes = types.Attributes{}
			}
			b.Attributes.Set(types.AttrQuoteAuthor, author)
		}
		raw, err := serializeBlock(elements)
		if err != nil {
			return err
		}
		b.Elements, err = parseContent(config.Filename, raw, append(options, Entrypoint("NormalBlockContentSubstitution"))...)
		return err
	}
}

func extractQuoteBlockAttribution(elements []interface{}) ([]interface{}, string) {
	log.Debug("extracting attribution on markdown block quote")
	// first, check if last line is an attribution (author)
	if len(elements) == 0 {
		return elements, ""
	}
	if l, ok := elements[len(elements)-1].(types.RawLine); ok {
		a, err := ParseReader("", strings.NewReader(l.Content), Entrypoint("MarkdownQuoteBlockAttribution"))
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

// disables substitutions
// returns the given content as-is (converting `RawLine` elements to `VerbatimLine` elements)
func none() blockSubstitution {
	return func(b *types.DelimitedBlock) error {
		for i, element := range b.Elements {
			switch e := element.(type) {
			case types.RawLine:
				b.Elements[i] = types.VerbatimLine{
					Content: e.Content,
				}
			}
		}
		return nil
	}
}

func parseRawLine(line types.RawLine, options ...Option) ([]interface{}, error) {
	result := []interface{}{}
	log.Debugf("parsing rawline '%s'", line.Content)
	e, err := ParseReader("", strings.NewReader(line.Content), options...)
	if err != nil {
		return nil, err
	}
	switch e := e.(type) {
	case []interface{}:
		result = append(result, e...)
	default:
		result = append(result, e)
	}
	log.Debugf("parsed elements: %v", result)
	return result, nil
}

func parseContent(filename string, content string, options ...Option) ([]interface{}, error) {
	// log.Debugf("parsing content '%s'", content)
	result, err := ParseReader(filename, strings.NewReader(content), options...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to parse '%s'", content)
	}
	if result, ok := result.([]interface{}); ok {
		// if log.IsLevelEnabled(log.DebugLevel) {
		// 	log.Debug("parsed content:")
		// 	spew.Fdump(log.StandardLogger().Out, types.Merge(result))
		// }
		return types.Merge(result), nil
	}
	return nil, fmt.Errorf("unexpected type of content: '%T'", result)
}

func serializeBlock(elements []interface{}) (string, error) {
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debug("serializing elements in a delimited block")
	// 	spew.Fdump(log.StandardLogger().Out, elements)
	// }
	buf := strings.Builder{}
	for i, elmt := range elements {
		if l, ok := elmt.(types.RawLine); ok {
			buf.WriteString(l.Content)
			if i < len(elements)-1 {
				buf.WriteString("\n")
			}
		} else {
			return "", fmt.Errorf("unexpected type of element while serializing the content of a delimited block: '%T'", elmt)
		}
	}
	log.Debugf("raw content: '%s'", buf.String())
	return buf.String(), nil
}

// ----------------------------------------------------------------------------
// Section substitutions
// ----------------------------------------------------------------------------

// applies the elements and attributes substitutions on the given section title.
func applySubstitutionsOnSection(s types.Section, attrs types.AttributesWithOverrides, options ...Option) (types.Section, error) {
	elements := s.Title
	subs := []elementsSubstitutionFunc{
		substituteInlinePassthroughFunc,
		substituteSpecialCharactersFunc,
		substituteQuotedTextsFunc,                 // done at the same time as the inline macros
		substituteAttributesFunc,                  // detect the replacements
		applyAttributeSubstitutionsOnElementsFunc, // apply the replacements
		substituteReplacementsFunc,
		substituteInlineMacrosFunc, // substituteQuotedTextAndInlineMacrosFunc,
		// resolveLocationsOnParagraphLines(attrs),
		substitutePostReplacementsFunc,
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

// applies the elements and attributes substitutions on the given image block.
func applySubstitutionsOnImageBlock(b types.ImageBlock, attrs types.AttributesWithOverrides, options ...Option) (types.ImageBlock, error) {
	elements := b.Location.Path
	subs := []elementsSubstitutionFunc{
		substituteAttributesFunc,                  // detect the replacements
		applyAttributeSubstitutionsOnElementsFunc, // apply the replacements
	}
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
	subs, err := paragraphSubstitutions(p.Attributes.GetAsStringWithDefault(types.AttrSubstitutions, "normal"))
	if err != nil {
		return types.Paragraph{}, err
	}
	elements := p.Lines
	// apply all the configured substitutions
	for _, sub := range subs {
		var err error
		if elements, err = sub(elements, attrs, options...); err != nil {
			return types.Paragraph{}, err
		}
	}
	p.Lines = splitLines(elements)
	// if log.IsLevelEnabled(log.DebugLevel) {
	// 	log.Debugf("paragraph after substitution:")
	// 	spew.Fdump(log.StandardLogger().Out, p)
	// }
	return p, nil
}

type elementsSubstitutionFunc func(lines []interface{}, attrs types.AttributesWithOverrides, options ...Option) ([]interface{}, error)

// paragraphSubstitutions returns the substitution funcs matching the given `subs` arg
// otherwise, returns a default substitution which will ultemately fail
func paragraphSubstitutions(subs string) ([]elementsSubstitutionFunc, error) {
	// log.Debugf("determining substitutions for '%s' on a paragraph", subs)
	funcs := []elementsSubstitutionFunc{}
	for _, s := range strings.Split(subs, ",") {
		switch s {
		case "specialcharacters", "specialchars":
			funcs = append(funcs, substituteSpecialCharactersFunc)
		case "quotes":
			funcs = append(funcs, substituteQuotedTextsFunc)
		case "attributes":
			funcs = append(funcs,
				substituteAttributesFunc,                  // detect the replacements
				applyAttributeSubstitutionsOnElementsFunc, // apply the replacements
			)
		case "macros":
			funcs = append(funcs,
				substituteInlineMacrosFunc,
			)
		case "replacements":
			funcs = append(funcs, substituteReplacementsFunc)
		case "post_replacements":
			funcs = append(funcs, substitutePostReplacementsFunc)
		case "normal":
			funcs = append(funcs,
				substituteInlinePassthroughFunc,
				substituteSpecialCharactersFunc,
				substituteQuotedTextsFunc,                 // done at the same time as the inline macros
				substituteAttributesFunc,                  // detect the replacements
				applyAttributeSubstitutionsOnElementsFunc, // apply the replacements
				substituteReplacementsFunc,
				substituteInlineMacrosFunc, // substituteQuotedTextAndInlineMacrosFunc,
				// resolveLocationsOnParagraphLines(attrs),
				substitutePostReplacementsFunc,
			)
		case "none":
			funcs = append(funcs, substituteNothingFunc)
		default:
			return nil, fmt.Errorf("unsupported substitution: '%s", s)
		}
	}
	return funcs, nil
}

var (
	substituteInlinePassthroughFunc = elementsSubstitution("InlinePassthroughSubstitution")
	substituteSpecialCharactersFunc = elementsSubstitution("SpecialCharactersSubstitution")
	substituteQuotedTextsFunc       = elementsSubstitutionWithPlaceholders("QuotedTextSubstitution")
	substituteAttributesFunc        = elementsSubstitutionWithPlaceholders("AttributesSubstitution") // TODO: include with applyAttributeSubstitutionsOnElementsFunc?
	substituteReplacementsFunc      = elementsSubstitutionWithPlaceholders("ReplacementsSubstitution")
	substituteInlineMacrosFunc      = elementsSubstitutionWithPlaceholders("InlineMacrosSubstitution") // elementsSubstitution("InlineMacrosSubstitution")
	substitutePostReplacementsFunc  = elementsSubstitutionWithPlaceholders("PostReplacementsSubstitution")
	substituteNothingFunc           = elementsSubstitution("NoneSubstitution")
)

func elementsSubstitution(ruleName string) elementsSubstitutionFunc {
	return func(elements []interface{}, _ types.AttributesWithOverrides, options ...Option) ([]interface{}, error) {
		log.Debugf("applying the '%s' substitution on elements", ruleName)
		result, err := parseElements(serializeElements(elements), append(options, Entrypoint(ruleName))...)
		if err != nil {
			return nil, err
		}
		if log.IsLevelEnabled(log.DebugLevel) {
			log.Debugf("applied '%s' substitution", ruleName)
			spew.Fdump(log.StandardLogger().Out, result)
		}
		return result, nil
	}
}

func elementsSubstitutionWithPlaceholders(ruleName string) elementsSubstitutionFunc {
	return func(elements []interface{}, attrs types.AttributesWithOverrides, options ...Option) ([]interface{}, error) {
		log.Debugf("applying the '%s' substitution on elements (imagesdir='%v')", ruleName, attrs.GetAsStringWithDefault("imagesdir", ""))
		elements, placeholders := serializeElementsWithPlaceHolders(elements)
		gb := GlobalStore("imagesdir", attrs.GetAsStringWithDefault("imagesdir", "")) // TODO:define a const for "imagesdir"
		options = append(options, Entrypoint(ruleName), gb)
		// process placeholder content (eg: quoted text may contain an inline link)
		for ref, placeholder := range placeholders {
			switch placeholder := placeholder.(type) {
			case types.QuotedString:
				var err error
				if placeholder.Elements, err = parseElements(placeholder.Elements, options...); err != nil {
					return nil, err
				}
				placeholders[ref] = placeholder
			case types.QuotedText:
				var err error
				if placeholder.Elements, err = parseElements(placeholder.Elements, options...); err != nil {
					return nil, err
				}
				placeholders[ref] = placeholder
			}
		}
		result := make([]interface{}, 0, len(elements))
		for _, element := range elements {
			switch element := element.(type) {
			case types.StringElement: // coming as-is from the Raw document
				elmts, err := parseContent("", element.Content, options...)
				if err != nil {
					return nil, err
				}
				elmts = restoreElements(elmts, placeholders)
				result = append(result, elmts...)
			default:
				result = append(result, element)
			}
		}
		if log.IsLevelEnabled(log.DebugLevel) {
			log.Debugf("applied '%s' substitution", ruleName)
			spew.Fdump(log.StandardLogger().Out, result)
		}
		return result, nil
	}
}

func parseElements(elements []interface{}, options ...Option) ([]interface{}, error) {
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

// replace the placeholders with their original element in the given elements
func restoreElements(elmts []interface{}, placeholders map[string]interface{}) []interface{} {
	// skip if there's nothing to restore
	if len(placeholders) == 0 {
		return elmts
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("restoring elements on")
		spew.Fdump(log.StandardLogger().Out, elmts)
	}
	for i, elmt := range elmts {
		switch elmt := elmt.(type) {
		case types.ElementPlaceHolder:
			elmts[i] = placeholders[elmt.Ref]
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
	return elmts
}

// replace the placeholders with their original element in the given attributes
func restoreAttributes(attrs types.Attributes, placeholders map[string]interface{}) types.Attributes {
	for key, value := range attrs {
		switch value := value.(type) {
		case types.ElementPlaceHolder:
			attrs[key] = placeholders[value.Ref]
		case types.ElementRole:
			attrs[key] = types.ElementRole(restoreElements(value, placeholders))
		case []interface{}:
			attrs[key] = restoreElements(value, placeholders)
		}
	}
	return attrs
}

var applyAttributeSubstitutionsOnElementsFunc = func(elements []interface{}, attrs types.AttributesWithOverrides, options ...Option) ([]interface{}, error) {
	for i, element := range elements {
		element, err := applyAttributeSubstitutionsOnElement(element, attrs)
		if err != nil {
			return nil, err
		}
		elements[i] = element
	}
	elements = types.Merge(elements)
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("applied attributes substitutions")
		spew.Fdump(log.StandardLogger().Out, elements)
	}
	return elements, nil
}

func splitLines(elements []interface{}) []interface{} {
	// after processing all the elements, we want to split them in separate lines again, to retain the initial input "form"
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
	return lines
}

func serializeElements(elements []interface{}) []interface{} {
	result := []interface{}{}
	for i, e := range elements {
		switch e := e.(type) {
		case []interface{}:
			result = append(result, e...)
			if i < len(elements)-1 {
				result = append(result, types.StringElement{
					Content: "\n",
				})
			}
		case types.RawLine:
			result = append(result, types.StringElement(e)) // converting
			if i < len(elements)-1 {
				result = append(result, types.StringElement{
					Content: "\n",
				})
			}
		case types.SingleLineComment:
			result = append(result, e)
		default:
			result = append(result, e)
		}
	}
	result = types.Merge(result)
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("serialized elements:")
		spew.Fdump(log.StandardLogger().Out, result)
	}
	return result
}

func serializeElementsWithPlaceHolders(elements []interface{}) ([]interface{}, map[string]interface{}) {
	result := []interface{}{}
	seq := 0
	placeholders := map[string]interface{}{}
	for i, e := range elements {
		switch e := e.(type) {
		case []interface{}:
			result = append(result, e...)
			if i < len(elements)-1 {
				result = append(result, types.StringElement{
					Content: "\n",
				})
			}
		case types.RawLine:
			result = append(result, types.StringElement(e)) // converting
			if i < len(elements)-1 {
				result = append(result, types.StringElement{
					Content: "\n",
				})
			}
		case types.StringElement:
			result = append(result, e)
		case types.SingleLineComment:
			result = append(result, e)
		default:
			// replace with placeholder
			seq++
			placeholders[strconv.Itoa(seq)] = e
			ph := types.ElementPlaceHolder{
				Ref: strconv.Itoa(seq),
			}
			result = append(result, types.StringElement{Content: ph.String()})
		}
	}
	result = types.Merge(result)
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("serialized elements:")
		spew.Fdump(log.StandardLogger().Out, result)
	}
	return result, placeholders
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
	log.Debugf("applying attribute substitution on element of type '%T'", element)
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
