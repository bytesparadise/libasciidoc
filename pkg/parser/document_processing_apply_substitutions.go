package parser

// // ----------------------------------------------------------------------------
// // Substitutions
// // ----------------------------------------------------------------------------

// // ApplySubstitutions applies all the substitutions on delimited blocks, standalone paragraphs and paragraphs
// // in continued list items, and then attribute substitutions, and as a result returns a `DraftDocument`.
// func ApplySubstitutions(rawDoc types.DocumentFragments, config *configuration.Configuration) (types.DraftDocument, error) {
// 	// ctx := newSubstitutionContext(config)
// 	// // also, add all front-matter key/values
// 	// ctx.attributes.Add(rawDoc.FrontMatter.Content)
// 	// // also, add all AttributeDeclaration at the top of the document
// 	// ctx.attributes.Add(rawDoc.Attributes())

// 	// elements, err := applySubstitutions(ctx, rawDoc.Elements)
// 	// if err != nil {
// 	// 	return types.DraftDocument{}, err
// 	// }
// 	// if len(elements) == 0 {
// 	// 	elements = nil // avoid carrying empty slice
// 	// }
// 	return types.DraftDocument{
// 		// Attributes:  ctx.attributes.All(),
// 		// FrontMatter: rawDoc.FrontMatter,
// 		// Elements:    elements,
// 	}, nil
// }

// type substitutionContextDeprecated struct {
// 	attributes types.AttributesWithOverrides // TODO: replace with types.Attributes?
// 	config     configuration.Configuration
// }

// func newSubstitutionContextDeprecated(config *configuration.Configuration) substitutionContextDeprecated {
// 	return substitutionContextDeprecated{
// 		attributes: types.AttributesWithOverrides{
// 			Content:   types.Attributes{},
// 			Overrides: config.AttributeOverrides,
// 			Counters:  map[string]interface{}{},
// 		},
// 		config: config,
// 	}
// }

// // applySubstitutions applies the substitutions on paragraphs and delimited blocks (including when in continued list elements)
// func applySubstitutions(ctx substitutionContextDeprecated, elements []interface{}) ([]interface{}, error) {
// 	if len(elements) == 0 {
// 		return nil, nil
// 	}
// 	if log.IsLevelEnabled(log.DebugLevel) {
// 		log.Debug("applying substitutions on:")
// 		spew.Fdump(log.StandardLogger().Out, elements)
// 	}
// 	result := make([]interface{}, len(elements))
// 	for i, e := range elements {
// 		if a, ok := e.(types.WithAttributesToSubstitute); ok {
// 			log.Debugf("applying substitution on attributes of element of type '%T'", e)
// 			attrs, err := applyAttributeSubstitutionsOnAttributes(ctx, a.GetAttributes())
// 			if err != nil {
// 				return nil, err
// 			}
// 			e = a.ReplaceAttributes(attrs)
// 		}
// 		var err error
// 		switch e := e.(type) {
// 		case types.WithNestedElementSubstitution:
// 			log.Debugf("applying substitution on nested elements of element of type '%T'", e)
// 			subs, err := substitutionsFor(e)
// 			if err != nil {
// 				return nil, err
// 			}
// 			elements, err := applySubstitutionsOnElements(ctx, e.ElementsToSubstitute(), subs)
// 			if err != nil {
// 				return nil, err
// 			}
// 			result[i] = e.ReplaceElements(elements)
// 		case types.WithLineSubstitution:
// 			log.Debugf("applying substitution on lines of element of type '%T'", e)
// 			subs, err := substitutionsFor(e)
// 			if err != nil {
// 				return nil, err
// 			}
// 			elements, err := applySubstitutionsOnLines(ctx, e.LinesToSubstitute(), subs)
// 			if err != nil {
// 				return nil, err
// 			}
// 			result[i] = e.SubstituteLines(elements)
// 		case types.MarkdownQuoteBlock: // slightly different since there is an extraction for the author attributions
// 			e, err := applySubstitutionsOnMarkdownQuoteBlock(ctx, e)
// 			if err != nil {
// 				return nil, err
// 			}
// 			result[i] = e
// 		case types.ContinuedListItemElement:
// 			r, err := applySubstitutions(ctx, []interface{}{e.Element})
// 			if err != nil {
// 				return nil, err
// 			}
// 			e.Element = r[0]
// 			result[i] = e
// 		case types.ImageBlock:
// 			if e.Location, err = applySubstitutionsOnLocation(ctx, e.Location); err != nil {
// 				return nil, err
// 			}
// 			result[i] = e
// 		case types.Section:
// 			if e, err = applySubstitutionsOnSection(ctx, e); err != nil {
// 				return nil, err
// 			}
// 			result[i] = e
// 		default:
// 			// no support for element substitution here
// 			// so let's proceed with attribute substitutions
// 			if e, err = applyAttributeSubstitutionsOnElement(ctx, e); err != nil {
// 				return nil, err
// 			}
// 			result[i] = e
// 		}
// 	}
// 	if log.IsLevelEnabled(log.DebugLevel) {
// 		log.Debug("applied substitutions:")
// 		spew.Fdump(log.StandardLogger().Out, result)
// 	}
// 	return result, nil
// }

// // ----------------------------------------------------------------------------
// // Delimited Block substitutions
// // ----------------------------------------------------------------------------

// var substitutions = map[string]elementsSubstitution{
// 	"inline_passthrough": substituteInlinePassthrough,
// 	"callouts":           substituteCallouts,
// 	"specialcharacters":  substituteSpecialCharacters,
// 	"specialchars":       substituteSpecialCharacters,
// 	"quotes":             substituteQuotedTexts,
// 	"attributes":         substituteAttributes,
// 	"replacements":       substituteReplacements,
// 	"macros":             substituteInlineMacros,
// 	"post_replacements":  substitutePostReplacements,
// 	"none":               substituteNone,
// }

// func substitutionsFor(block types.WithCustomSubstitutions) ([]elementsSubstitution, error) {
// 	subs := make(funcs, 0, len(substitutions))
// 	subsitutionsToApply, err := block.SubstitutionsToApply()
// 	if err != nil {
// 		return nil, err
// 	}
// 	for _, s := range subsitutionsToApply {
// 		switch s {
// 		case "normal":
// 			subs = subs.append(
// 				"specialcharacters",
// 				"quotes",
// 				"attributes",
// 				"replacements",
// 				"macros",
// 				"post_replacements",
// 			)
// 		case "inline_passthrough", "callouts", "specialcharacters", "specialchars", "quotes", "attributes", "macros", "replacements", "post_replacements", "none":
// 			subs = subs.append(s)
// 		case "+callouts", "+specialcharacters", "+specialchars", "+quotes", "+attributes", "+macros", "+replacements", "+post_replacements", "+none":
// 			if len(subs) == 0 {
// 				subs = subs.append(block.DefaultSubstitutions()...)
// 			}
// 			subs = subs.append(strings.ReplaceAll(s, "+", ""))
// 		case "callouts+", "specialcharacters+", "specialchars+", "quotes+", "attributes+", "macros+", "replacements+", "post_replacements+", "none+":
// 			if len(subs) == 0 {
// 				subs = subs.append(block.DefaultSubstitutions()...)
// 			}
// 			subs = subs.prepend(strings.ReplaceAll(s, "+", ""))
// 		case "-callouts", "-specialcharacters", "-specialchars", "-quotes", "-attributes", "-macros", "-replacements", "-post_replacements", "-none":
// 			if len(subs) == 0 {
// 				subs = subs.append(block.DefaultSubstitutions()...)
// 			}
// 			subs = subs.remove(strings.ReplaceAll(s, "-", ""))
// 		default:
// 			return nil, fmt.Errorf("unsupported substitution: '%s", s)
// 		}
// 	}
// 	result := make([]elementsSubstitution, 0, len(subs))
// 	for _, s := range subs {
// 		if f, exists := substitutions[s]; exists {
// 			result = append(result, f)
// 		}
// 	}
// 	result = append(result, splitLines)
// 	return result, nil
// }

// type funcs []string

// func (f funcs) append(others ...string) funcs {
// 	return append(f, others...)
// }

// func (f funcs) prepend(other string) funcs {
// 	return append(funcs{other}, f...)
// }

// func (f funcs) remove(other string) funcs {
// 	for i, s := range f {
// 		if s == other {
// 			return append(f[:i], f[i+1:]...)
// 		}
// 	}
// 	// unchanged
// 	return f
// }

// func applySubstitutionsOnElements(ctx substitutionContextDeprecated, elements []interface{}, subs []elementsSubstitution) ([]interface{}, error) {
// 	// apply all the substitutions on elements that need to be processed
// 	for i, element := range elements {
// 		switch e := element.(type) {
// 		// if the block contains a block...
// 		case types.WithNestedElementSubstitution:
// 			lines, err := applySubstitutionsOnElements(ctx, e.ElementsToSubstitute(), subs)
// 			if err != nil {
// 				return nil, err
// 			}
// 			elements[i] = e.ReplaceElements(lines)
// 		case types.WithLineSubstitution:
// 			lines, err := applySubstitutionsOnLines(ctx, e.LinesToSubstitute(), subs)
// 			if err != nil {
// 				return nil, err
// 			}
// 			elements[i] = e.SubstituteLines(lines)
// 		default:
// 			// log.Debugf("nothing to substitute on element of type '%T'", element)
// 			// do nothing
// 		}
// 	}
// 	return elements, nil
// }

// func applySubstitutionsOnLines(ctx substitutionContextDeprecated, lines [][]interface{}, subs []elementsSubstitution) ([][]interface{}, error) {
// 	var err error
// 	for _, sub := range subs {
// 		if lines, err = sub(ctx, lines); err != nil {
// 			return nil, err
// 		}
// 	}
// 	return lines, nil
// }

// func applySubstitutionsOnMarkdownQuoteBlock(ctx substitutionContextDeprecated, b types.MarkdownQuoteBlock) (types.MarkdownQuoteBlock, error) {
// 	funcs := []elementsSubstitution{
// 		substituteInlinePassthrough,
// 		substituteSpecialCharacters,
// 		substituteQuotedTexts,
// 		substituteAttributes,
// 		substituteReplacements,
// 		substituteInlineMacros,
// 		substitutePostReplacements,
// 		splitLines}
// 	// attempt to extract the block attributions
// 	var author string
// 	if b.Lines, author = extractMarkdownQuoteAttribution(b.Lines); author != "" {
// 		if b.Attributes == nil {
// 			b.Attributes = types.Attributes{}
// 		}
// 		b.Attributes.Set(types.AttrQuoteAuthor, author)
// 	}
// 	if len(b.Lines) == 0 { // no more line to parse after extracting the author
// 		b.Lines = nil
// 		return b, nil
// 	}
// 	// apply all the substitutions
// 	var err error
// 	for _, sub := range funcs {
// 		if b.Lines, err = sub(ctx, b.Lines); err != nil {
// 			return types.MarkdownQuoteBlock{}, err
// 		}
// 	}
// 	return b, nil
// }

// func extractMarkdownQuoteAttribution(lines [][]interface{}) ([][]interface{}, string) {
// 	// first, check if last line is an attribution (author)
// 	if len(lines) == 0 {
// 		return lines, ""
// 	}
// 	if l, ok := lines[len(lines)-1][0].(types.StringElement); ok {
// 		a, err := ParseReader("", strings.NewReader(l.Content), Entrypoint("MarkdownQuoteAttribution"))
// 		// assume that the last line is not an author attribution if an error occurred
// 		if err != nil {
// 			return lines, ""
// 		}
// 		if a, ok := a.(string); ok {
// 			// log.Debugf("found attribution in markdown block: '%s'", a)
// 			return lines[:len(lines)-1], a
// 		}
// 	}
// 	return lines, ""
// }

// // ----------------------------------------------------------------------------
// // Section substitutions
// // ----------------------------------------------------------------------------

// // applies the elements and attributes substitutions on the given section title.
// func applySubstitutionsOnSection(ctx substitutionContextDeprecated, s types.Section) (types.Section, error) {
// 	elements := [][]interface{}{s.Title} // wrap to match the `elementsSubstitution` arg type
// 	subs := []elementsSubstitution{
// 		substituteInlinePassthrough,
// 		substituteSpecialCharacters,
// 		substituteQuotedTexts,
// 		substituteAttributes,
// 		substituteReplacements,
// 		substituteInlineMacros,
// 		substitutePostReplacements,
// 	}
// 	var err error
// 	for _, sub := range subs {
// 		if elements, err = sub(ctx, elements); err != nil {
// 			return types.Section{}, err
// 		}
// 	}
// 	s.Title = elements[0]
// 	if s, err = s.ResolveID(ctx.attributes); err != nil {
// 		return types.Section{}, err
// 	}
// 	// if log.IsLevelEnabled(log.DebugLevel) {
// 	// 	// log.Debugf("section after substitution:")
// 	// 	spew.Fdump(log.StandardLogger().Out, s)
// 	// }
// 	return s, nil
// }

// // ----------------------------------------------------------------------------
// // Location substitutions
// // ----------------------------------------------------------------------------

// // applies the elements and attributes substitutions on the given image block.
// func applySubstitutionsOnLocation(ctx substitutionContextDeprecated, l types.Location) (types.Location, error) {
// 	elements := [][]interface{}{l.Path} // wrap to match the `elementsSubstitution` arg type
// 	subs := []elementsSubstitution{substituteAttributes}
// 	var err error
// 	for _, sub := range subs {
// 		if elements, err = sub(ctx, elements); err != nil {
// 			return types.Location{}, err
// 		}
// 	}
// 	l.Path = elements[0]
// 	l = l.WithPathPrefix(ctx.attributes.GetAsStringWithDefault(types.AttrImagesDir, ""))
// 	return l, nil
// }

// // ----------------------------------------------------------------------------
// // Individual substitution funcs
// // ----------------------------------------------------------------------------

// // includes a call to `elementsSubstitution` with some post-processing on the result
// var substituteAttributes = func(ctx substitutionContextDeprecated, lines [][]interface{}) ([][]interface{}, error) {
// 	lines, err := newElementsSubstitution("AttributeSubs")(ctx, lines) // TODO: add a `substituteAttributes` var?
// 	if err != nil {
// 		return nil, err
// 	}
// 	for i, line := range lines {
// 		line, err := applyAttributeSubstitutionsOnElements(ctx, line)
// 		if err != nil {
// 			return nil, err
// 		}
// 		lines[i] = types.Merge(line)
// 	}
// 	// if log.IsLevelEnabled(log.DebugLevel) {
// 	// 	// log.Debugf("applied the 'attributes' substitution")
// 	// 	spew.Fdump(log.StandardLogger().Out, lines)
// 	// }
// 	return lines, nil
// }

// var (
// 	substituteInlinePassthrough = newElementsSubstitution("InlinePassthroughSubs")
// 	substituteSpecialCharacters = newElementsSubstitution("SpecialCharacterSubs")
// 	substituteQuotedTexts       = newElementsSubstitution("QuotedTextSubs")
// 	substituteReplacements      = newElementsSubstitution("ReplacementSubs")
// 	substituteInlineMacros      = newElementsSubstitution("InlineMacroSubs")
// 	substitutePostReplacements  = newElementsSubstitution("PostReplacementSubs")
// 	substituteNone              = newElementsSubstitution("NoneSubs")
// 	substituteCallouts          = newElementsSubstitution("CalloutSubs")
// )

// type elementsSubstitution func(ctx substitutionContextDeprecated, lines [][]interface{}) ([][]interface{}, error)

// func newElementsSubstitution(rule string) elementsSubstitution {
// 	return func(ctx substitutionContextDeprecated, lines [][]interface{}) ([][]interface{}, error) {
// 		log.Debugf("applying the '%s' rule on elements", rule)
// 		placeholders := &placeholders{
// 			seq:      0,
// 			elements: map[string]interface{}{},
// 		}
// 		s := serializeLines(lines, placeholders)
// 		imagesdirOption := GlobalStore(types.AttrImagesDir, ctx.attributes.GetAsStringWithDefault(types.AttrImagesDir, ""))
// 		usermacrosOptions := GlobalStore(usermacrosKey, ctx.config.Macros)
// 		// process placeholder content (eg: quoted text may contain an inline link)
// 		for ref, placeholder := range placeholders.elements {
// 			switch placeholder := placeholder.(type) { // TODO: create `PlaceHolder` interface?
// 			case types.QuotedString:
// 				var err error
// 				if placeholder.Elements, err = parserPlaceHolderElements(placeholder.Elements, imagesdirOption, usermacrosOptions, Entrypoint(rule)); err != nil {
// 					return nil, err
// 				}
// 				placeholders.elements[ref] = placeholder
// 			case types.QuotedText:
// 				var err error
// 				if placeholder.Elements, err = parserPlaceHolderElements(placeholder.Elements, imagesdirOption, usermacrosOptions, Entrypoint(rule)); err != nil {
// 					return nil, err
// 				}
// 				placeholders.elements[ref] = placeholder
// 			}
// 		}
// 		elmts, err := parseContent("", s, imagesdirOption, usermacrosOptions, Entrypoint(rule))
// 		if err != nil {
// 			return nil, err
// 		}
// 		elmts = restorePlaceholderElements(elmts, placeholders)
// 		if log.IsLevelEnabled(log.DebugLevel) {
// 			log.Debugf("applied the '%s' rule:", rule)
// 			spew.Fdump(log.StandardLogger().Out, [][]interface{}{elmts})
// 		}
// 		return [][]interface{}{elmts}, nil
// 	}
// }

// func parserPlaceHolderElements(elements []interface{}, opts ...Option) ([]interface{}, error) {
// 	result := make([]interface{}, 0, len(elements)) // default capacity (but may not be enough)
// 	for _, element := range elements {
// 		switch element := element.(type) {
// 		case *types.StringElement:
// 			elmts, err := parseContent("", element.Content, opts...)
// 			if err != nil {
// 				return nil, err
// 			}
// 			result = append(result, elmts...)
// 		default:
// 			result = append(result, element)
// 		}
// 	}
// 	return result, nil
// }

// func parseContent(filename string, content string, opts ...Option) ([]interface{}, error) {
// 	result, err := ParseReader(filename, strings.NewReader(content), opts...)
// 	if err != nil {
// 		return nil, errors.Wrapf(err, "unable to parse '%s'", content)
// 	}
// 	if result, ok := result.([]interface{}); ok {
// 		return types.Merge(result), nil
// 	}
// 	return []interface{}{result}, nil
// }

// // replace the placeholders with their original element in the given elements
// func restorePlaceholderElements(elements []interface{}, placeholders *placeholders) []interface{} {
// 	// skip if there's nothing to restore
// 	if len(placeholders.elements) == 0 {
// 		return elements
// 	}
// 	for i, e := range elements {
// 		//
// 		if e, ok := e.(types.ElementPlaceHolder); ok {
// 			elements[i] = placeholders.elements[e.Ref]
// 		}
// 		// for each element, check *all* interfaces to see if there's a need to replace the placeholders
// 		if e, ok := e.(types.WithPlaceholdersInElements); ok {
// 			elements[i] = e.RestoreElements(placeholders.elements)
// 		}
// 		if e, ok := e.(types.WithPlaceholdersInAttributes); ok {
// 			elements[i] = e.RestoreAttributes(placeholders.elements)
// 		}
// 		if e, ok := e.(types.WithPlaceholdersInLocation); ok {
// 			elements[i] = e.RestoreLocation(placeholders.elements)
// 		}
// 	}
// 	return elements
// }

// type placeholders struct {
// 	seq      int
// 	elements map[string]interface{}
// }

// func (p *placeholders) add(element interface{}) types.ElementPlaceHolder {
// 	p.seq++
// 	p.elements[strconv.Itoa(p.seq)] = element
// 	return types.ElementPlaceHolder{
// 		Ref: strconv.Itoa(p.seq),
// 	}

// }

// func serializeLines(lines [][]interface{}, placeholders *placeholders) string {
// 	result := strings.Builder{}
// 	for i, line := range lines {
// 		for _, e := range line {
// 			switch e := e.(type) {
// 			case *types.StringElement:
// 				result.WriteString(e.Content)
// 			case types.SingleLineComment:
// 				// replace with placeholder
// 				p := placeholders.add(e)
// 				result.WriteString(p.String())
// 			default:
// 				// replace with placeholder
// 				p := placeholders.add(e)
// 				result.WriteString(p.String())
// 			}
// 		}
// 		if i < len(lines)-1 {
// 			result.WriteString("\n")
// 		}
// 	}
// 	// if log.IsLevelEnabled(log.DebugLevel) {
// 	// 	log.Debug("serialized lines:")
// 	// 	spew.Fdump(log.StandardLogger().Out, result.String())
// 	// }
// 	return result.String()
// }

// func splitLines(_ substitutionContextDeprecated, lines [][]interface{}) ([][]interface{}, error) {
// 	result := make([][]interface{}, 0, len(lines))
// 	for _, line := range lines {
// 		pendingLine := []interface{}{}
// 		for _, element := range line {
// 			switch element := element.(type) {
// 			case *types.StringElement:
// 				// if content has line feeds, then split in multiple lines
// 				split := strings.Split(element.Content, "\n")
// 				for i, s := range split {
// 					if len(s) > 0 { // no need to append an empty StringElement
// 						pendingLine = append(pendingLine, types.StringElement{Content: s})
// 					}
// 					if i < len(split)-1 {
// 						result = append(result, pendingLine)
// 						pendingLine = []interface{}{} // reset for the next line
// 					}
// 				}
// 			default:
// 				pendingLine = append(pendingLine, element)
// 			}
// 		}
// 		// don't forget the last line (if applicable)
// 		result = append(result, pendingLine)
// 	}
// 	// if log.IsLevelEnabled(log.DebugLevel) {
// 	// 	log.Debug("splitted lines")
// 	// 	spew.Fdump(log.StandardLogger().Out, result)
// 	// }
// 	return result, nil
// }

// // ----------------------------------------------------------------------------
// // Attribute substitutions
// // ----------------------------------------------------------------------------

// func applyAttributeSubstitutionsOnElements(ctx substitutionContextDeprecated, elements []interface{}) ([]interface{}, error) {
// 	result := make([]interface{}, len(elements)) // maximum capacity should exceed initial input
// 	for i, element := range elements {
// 		e, err := applyAttributeSubstitutionsOnElement(ctx, element)
// 		if err != nil {
// 			return nil, err
// 		}
// 		result[i] = e
// 	}
// 	return result, nil
// }

// func applyAttributeSubstitutionsOnAttributes(ctx substitutionContextDeprecated, attributes types.Attributes) (types.Attributes, error) {
// 	for key, value := range attributes {
// 		switch key {
// 		case types.AttrRoles, types.AttrOptions: // multi-value attributes
// 			result := []interface{}{}
// 			if values, ok := value.([]interface{}); ok {
// 				for _, value := range values {
// 					switch value := value.(type) {
// 					case []interface{}:
// 						value, err := applyAttributeSubstitutionsOnElements(ctx, value)
// 						if err != nil {
// 							return nil, err
// 						}
// 						result = append(result, types.Reduce(value))
// 					default:
// 						result = append(result, value)
// 					}

// 				}
// 				attributes[key] = result
// 			}
// 		default: // single-value attributes
// 			if value, ok := value.([]interface{}); ok {
// 				value, err := applyAttributeSubstitutionsOnElements(ctx, value)
// 				if err != nil {
// 					return nil, err
// 				}
// 				attributes[key] = types.Reduce(value)
// 			}
// 		}
// 	}
// 	// if log.IsLevelEnabled(log.DebugLevel) {
// 	// 	log.Debug("applied substitutions on attributes")
// 	// 	spew.Fdump(log.StandardLogger().Out, attributes)
// 	// }
// 	return attributes, nil
// }

// func applyAttributeSubstitutionsOnLines(ctx substitutionContextDeprecated, lines [][]interface{}) ([][]interface{}, error) {
// 	for i, line := range lines {
// 		line, err := applyAttributeSubstitutionsOnElements(ctx, line)
// 		if err != nil {
// 			return nil, err
// 		}
// 		lines[i] = types.Merge(line)
// 	}
// 	return lines, nil
// }

// func applyAttributeSubstitutionsOnElement(ctx substitutionContextDeprecated, element interface{}) (interface{}, error) {
// 	var err error
// 	switch e := element.(type) {
// 	case types.AttributeReset:
// 		ctx.attributes.Set(e.Name, nil) // This allows us to test for a reset vs. undefined.
// 	case types.AttributeSubstitution:
// 		if value, ok := ctx.attributes.GetAsString(e.Name); ok {
// 			element = types.StringElement{
// 				Content: value,
// 			}
// 			break
// 		}
// 		log.Warnf("unable to find attribute '%s'", e.Name)
// 		element = types.StringElement{
// 			Content: "{" + e.Name + "}",
// 		}
// 	case types.CounterSubstitution:
// 		if element, err = applyCounterSubstitution(ctx, e); err != nil {
// 			return nil, err
// 		}
// 	case types.WithElementsToSubstitute:
// 		elmts, err := applyAttributeSubstitutionsOnElements(ctx, e.ElementsToSubstitute())
// 		if err != nil {
// 			return nil, err
// 		}
// 		element = e.ReplaceElements(types.Merge(elmts))
// 	case types.WithLineSubstitution:
// 		lines, err := applyAttributeSubstitutionsOnLines(ctx, e.LinesToSubstitute())
// 		if err != nil {
// 			return nil, err
// 		}
// 		element = e.SubstituteLines(lines)
// 	case types.ContinuedListItemElement:
// 		if e.Element, err = applyAttributeSubstitutionsOnElement(ctx, e.Element); err != nil {
// 			return nil, err
// 		}
// 	}
// 	// also, retain the attribute declaration value (if applicable)
// 	if e, ok := element.(types.AttributeDeclaration); ok {
// 		ctx.attributes.Set(e.Name, e.Value)
// 	}
// 	return element, nil
// }

// // applyCounterSubstitutions is called by applyAttributeSubstitutionsOnElement.  Unless there is an error with
// // the element (the counter is the wrong type, which should never occur), it will return a `StringElement, true`
// // (because we always either find the element, or allocate one), and `nil`.  On an error it will return `nil, false`,
// // and the error.  The extra boolean here is to fit the calling expectations of our caller.  This function was
// // factored out of a case from applyAttributeSubstitutionsOnElement in order to reduce the complexity of that
// // function, but otherwise it should have no callers.
// func applyCounterSubstitution(ctx substitutionContextDeprecated, c types.CounterSubstitution) (interface{}, error) {
// 	counter := ctx.attributes.Counters[c.Name]
// 	if counter == nil {
// 		counter = 0
// 	}
// 	increment := true
// 	if c.Value != nil {
// 		ctx.attributes.Counters[c.Name] = c.Value
// 		counter = c.Value
// 		increment = false
// 	}
// 	switch counter := counter.(type) {
// 	case int:
// 		if increment {
// 			counter++
// 		}
// 		ctx.attributes.Counters[c.Name] = counter
// 		if c.Hidden {
// 			// return empty string facilitates merging
// 			return &types.StringElement{Content: ""}, nil
// 		}
// 		return &types.StringElement{
// 			Content: strconv.Itoa(counter),
// 		}, nil
// 	case rune:
// 		if increment {
// 			counter++
// 		}
// 		ctx.attributes.Counters[c.Name] = counter
// 		if c.Hidden {
// 			// return empty string facilitates merging
// 			return &types.StringElement{Content: ""}, nil
// 		}
// 		return &types.StringElement{
// 			Content: string(counter),
// 		}, nil

// 	default:
// 		return nil, fmt.Errorf("invalid counter type %T", counter)
// 	}
// }
