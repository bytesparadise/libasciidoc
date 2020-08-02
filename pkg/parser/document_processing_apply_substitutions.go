package parser

import (
	"fmt"
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
	log.Debug("apply block substitutions")
	if len(elements) == 0 {
		return nil, nil
	}
	result := []interface{}{}
	for _, e := range elements {
		switch e := e.(type) {
		case types.Paragraph:
			p, err := applyParagraphSubstitutions(e, attrs, options...)
			if err != nil {
				return nil, err
			}
			result = append(result, p)
		case types.DelimitedBlock:
			subs := delimitedBlockSubstitutions(e.Kind, config, options...)
			if err := applyDelimitedBlockSubstitutions(&e, subs); err != nil {
				return nil, err
			}
			r, err := applyAttributeSubstitutions(e, attrs)
			if err != nil {
				return nil, err
			}
			result = append(result, r)
		case types.ContinuedListItemElement:
			r, err := applySubstitutions([]interface{}{e.Element}, attrs, config, options...)
			if err != nil {
				return nil, err
			}
			r[0], err = applyAttributeSubstitutions(r[0], attrs)
			if err != nil {
				return nil, err
			}
			e.Element = r[0]
			result = append(result, e)
		default:
			// no support for element substitution here
			// so let's proceed with attribute substitutions
			e, err := applyAttributeSubstitutions(e, attrs)
			if err != nil {
				return nil, err
			}
			result = append(result, e)
		}
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debug("after substitutions:")
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

// applyDelimitedBlockSubstitutions parses the given raw elements, depending on the given substitutions to apply
// May return the elements unchanged, or convert the elements to a source doc and parse with a custom entrypoint
func applyDelimitedBlockSubstitutions(b *types.DelimitedBlock, subs []blockSubstitution) error {
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
				if err := applyDelimitedBlockSubstitutions(&d, subs); err != nil {
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
	log.Debugf("parsing '%s'", line.Content)
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
	log.Debugf("parsing '%s'", content)
	result, err := ParseReader(filename, strings.NewReader(content), options...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to parse '%s'", content)
	}
	if result, ok := result.([]interface{}); ok {
		if log.IsLevelEnabled(log.DebugLevel) {
			log.Debug("parsed content:")
			spew.Fdump(log.StandardLogger().Out, types.Merge(result))
		}
		return types.Merge(result), nil
	}
	return nil, fmt.Errorf("unexpected type of content: '%T'", result)
}

func serializeBlock(elements []interface{}) (string, error) {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debug("serializing elements in a delimited block")
		spew.Fdump(log.StandardLogger().Out, elements)
	}
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
// Paragraph substitutions
// ----------------------------------------------------------------------------

// applies the elements and attributes substitutions on the given paragraph.
// Attributes substitution is triggered only if there is no specific substitution or if the `attributes` substitution is explicitly set.
func applyParagraphSubstitutions(p types.Paragraph, attrs types.AttributesWithOverrides, options ...Option) (types.Paragraph, error) {
	subs := p.Attributes.GetAsStringWithDefault(types.AttrSubstitutions, "normal")
	// apply all the configured substitutions
	for _, sub := range strings.Split(subs, ",") {
		lines, err := substitution(sub)(p.Lines, options...)
		if err != nil {
			return types.Paragraph{}, err
		}
		p.Lines = lines
		switch sub {
		case "normal", "attributes":
			if p.Lines, err = applyAttributeSubstitutionsOnLines(p.Lines, attrs); err != nil {
				return types.Paragraph{}, err
			}
		}
	}
	return p, nil
}

type paragraphSubstitutionFunc func(lines []interface{}, options ...Option) ([]interface{}, error)

// substitution returns the substitution func matching the given name
// otherwise, returns a default substitution which will ultemately fail
func substitution(name string) paragraphSubstitutionFunc {
	log.Debugf("applying the '%s' paragraph substitution", name)
	switch name {
	case "normal":
		return paragraphSubstitution("NormalParagraphContentSubstitution")
	case "quotes":
		return paragraphSubstitution("QuotedTextSubstitution")
	case "macros":
		return paragraphSubstitution("InlineMacrosSubstitution")
	case "specialcharacters", "specialchars":
		return paragraphSubstitution("SpecialCharactersSubstitution")
	case "attributes":
		return paragraphSubstitution("AttributesSubstitution")
	case "replacements":
		return paragraphSubstitution("ReplacementsSubstitution")
	case "none":
		return paragraphSubstitution("NoneSubstitution")
	default:
		return func(lines []interface{}, options ...Option) ([]interface{}, error) {
			return nil, fmt.Errorf("unsupported substitution: '%s", name)
		}
	}
}

func paragraphSubstitution(entrypoint string) paragraphSubstitutionFunc {
	return func(lines []interface{}, options ...Option) ([]interface{}, error) {
		elements := []interface{}{}
		for _, element := range serializeParagraphLines(lines) {
			switch element := element.(type) {
			case types.StringElement: // coming as-is from the Raw document
				elmts, err := parseContent("", element.Content, Entrypoint(entrypoint))
				if err != nil {
					return nil, err
				}
				elements = append(elements, elmts...)
			default:
				elements = append(elements, element)
			}
		}
		result := splitLines(elements)
		if log.IsLevelEnabled(log.DebugLevel) {
			log.Debugf("paragraph lines after substitution:")
			spew.Fdump(log.StandardLogger().Out, result)
		}
		return result, nil
	}
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

// // substitutes textual and character reference replacements
// var replacements = func(elements []interface{}, config configuration.Configuration, options ...Option) ([]interface{}, error) {
// 	return elements, nil
// }

// // replaces the line break character (+)
// var postReplacements = func(elements []interface{}, config configuration.Configuration, options ...Option) ([]interface{}, error) {
// 	return elements, nil
// }

func serializeParagraphLines(elements []interface{}) []interface{} {
	result := []interface{}{}
	for i, e := range elements {
		switch e := e.(type) {
		case types.RawLine:
			result = append(result, types.StringElement(e)) // converting
			if i < len(elements)-1 {
				result = append(result, types.StringElement{
					Content: "\n",
				})
			}
		case types.SingleLineComment:
			result = append(result, e)
		case []interface{}:
			result = append(result, e...)
			if i < len(elements)-1 {
				result = append(result, types.StringElement{
					Content: "\n",
				})
			}
		}
	}
	result = types.Merge(result)
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("serialized paragraph:")
		spew.Fdump(log.StandardLogger().Out, result)
	}
	return result
}

// ----------------------------------------------------------------------------
// Attribute substitutions
// ----------------------------------------------------------------------------

// applyAttributeSubstitutions applies the document attribute substitutions
// and re-parses the content if they were affected (ie, a substitution occurred)
func applyAttributeSubstitutions(element interface{}, attrs types.AttributesWithOverrides) (interface{}, error) {
	result, _, err := applyAttributeSubstitutionsOnElement(element, attrs)
	return result, err

}

func applyAttributeSubstitutionsOnElement(element interface{}, attrs types.AttributesWithOverrides) (interface{}, bool, error) {
	log.Debugf("applying attribute substitutions on  element of type '%T'", element)
	switch e := element.(type) {
	case types.AttributeDeclaration:
		attrs.Set(e.Name, e.Value)
		return e, false, nil
	case types.AttributeReset:
		attrs.Set(e.Name, nil) // This allows us to test for a reset vs. undefined.
		return e, false, nil
	case types.AttributeSubstitution:
		if value, ok := attrs.GetAsString(e.Name); ok {
			return types.StringElement{
				Content: value,
			}, true, nil
		}
		log.Warnf("unable to find attribute '%s'", e.Name)
		return types.StringElement{
			Content: "{" + e.Name + "}",
		}, false, nil
	case types.CounterSubstitution:
		return applyCounterSubstitution(e, attrs)
	case types.ImageBlock:
		return e.ResolveLocation(attrs), false, nil
	case types.InlineImage:
		return e.ResolveLocation(attrs), false, nil
	case types.ExternalCrossReference:
		return e.ResolveLocation(attrs), false, nil
	case types.Section:
		title, _, err := applyAttributeSubstitutionsOnElements(e.Title, attrs)
		if err != nil {
			return nil, false, err
		}
		e.Title = title
		e, err = e.ResolveID(attrs)
		if err != nil {
			return nil, false, err
		}
		return e, false, nil
	case types.OrderedListItem:
		elmts, _, err := applyAttributeSubstitutionsOnElements(e.Elements, attrs)
		if err != nil {
			return nil, false, err
		}
		e.Elements = elmts
		return e, false, nil
	case types.UnorderedListItem:
		elmts, _, err := applyAttributeSubstitutionsOnElements(e.Elements, attrs)
		if err != nil {
			return nil, false, err
		}
		e.Elements = elmts
		return e, false, nil
	case types.LabeledListItem:
		elmts, _, err := applyAttributeSubstitutionsOnElements(e.Elements, attrs)
		if err != nil {
			return nil, false, err
		}
		e.Elements = elmts
		return e, false, nil
	case types.QuotedText:
		elmts, _, err := applyAttributeSubstitutionsOnElements(e.Elements, attrs)
		if err != nil {
			return nil, false, err
		}
		e.Elements = elmts
		return e, false, nil
	case types.ContinuedListItemElement:
		elmt, applied, err := applyAttributeSubstitutionsOnElement(e.Element, attrs)
		if err != nil {
			return nil, false, err
		}
		e.Element = elmt
		return e, applied, nil
	case types.DelimitedBlock:
		elmts, _, err := applyAttributeSubstitutionsOnElements(e.Elements, attrs)
		if err != nil {
			return nil, false, err
		}
		e.Elements = elmts
		return e, false, nil
	case types.Paragraph:
		for i, l := range e.Lines {
			if l, ok := l.([]interface{}); ok {
				l, _, err := applyAttributeSubstitutionsOnElement(l, attrs)
				if err != nil {
					return nil, false, err
				}
				e.Lines[i] = l
			}
		}
		return e, false, nil
	case []interface{}:
		return applyAttributeSubstitutionsOnElements(e, attrs)
	default:
		return e, false, nil
	}
}

func applyAttributeSubstitutionsOnElements(elements []interface{}, attrs types.AttributesWithOverrides) ([]interface{}, bool, error) {
	result := make([]interface{}, 0, len(elements)) // maximum capacity should exceed initial input
	applied := false
	for _, element := range elements {
		e, a, err := applyAttributeSubstitutionsOnElement(element, attrs)
		if err != nil {
			return nil, false, err
		}
		result = append(result, e)
		applied = applied || a
	}
	result = types.Merge(result)
	if !applied {
		return result, false, nil
	}
	result, err := parseInlineLinks(result)
	return result, true, err
}

func applyAttributeSubstitutionsOnLines(lines []interface{}, attrs types.AttributesWithOverrides) ([]interface{}, error) {
	for i, line := range lines {
		if line, ok := line.([]interface{}); ok {
			line, _, err := applyAttributeSubstitutionsOnElements(line, attrs)
			if err != nil {
				return nil, err
			}
			lines[i] = line
		}
	}
	return lines, nil
}

// applyCounterSubstitutions is called by applyAttributeSubstitutionsOnElement.  Unless there is an error with
// the element (the counter is the wrong type, which should never occur), it will return a `StringElement, true`
// (because we always either find the element, or allocate one), and `nil`.  On an error it will return `nil, false`,
// and the error.  The extra boolean here is to fit the calling expectations of our caller.  This function was
// factored out of a case from applyAttributeSubstitutionsOnElement in order to reduce the complexity of that
// function, but otherwise it should have no callers.
func applyCounterSubstitution(c types.CounterSubstitution, attrs types.AttributesWithOverrides) (interface{}, bool, error) {
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
			return types.StringElement{Content: ""}, true, nil
		}
		return types.StringElement{
			Content: strconv.Itoa(counter),
		}, true, nil
	case rune:
		if increment {
			counter++
		}
		attrs.Counters[c.Name] = counter
		if c.Hidden {
			// return empty string facilitates merging
			return types.StringElement{Content: ""}, true, nil
		}
		return types.StringElement{
			Content: string(counter),
		}, true, nil

	default:
		return nil, false, fmt.Errorf("invalid counter type %T", counter)
	}

}

// if a document attribute substitution happened, we need to parse the string element in search
// for a potentially new link. Eg `{url}` giving `https://foo.com`
func parseInlineLinks(elements []interface{}) ([]interface{}, error) {
	result := []interface{}{}
	for _, element := range elements {
		switch element := element.(type) {
		case types.StringElement:
			log.Debugf("looking for links in line element of type %[1]T (%[1]v)", element)
			elements, err := ParseReader("", strings.NewReader(element.Content), Entrypoint("InlineLinks"))
			if err != nil {
				return []interface{}{}, errors.Wrap(err, "error while parsing content for inline links")
			}
			log.Debugf("  found: %+v", elements)
			result = append(result, elements.([]interface{})...)
		default:
			result = append(result, element)
		}
	}
	return result, nil
}
