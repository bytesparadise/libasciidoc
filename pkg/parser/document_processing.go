package parser

import (
	"io"
	"reflect"
	"strconv"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"

	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// ParseDocument parses the content of the reader identitied by the filename
func ParseDocument(filename string, r io.Reader, opts ...Option) (types.Document, error) {
	draftDoc, err := ParseDraftDocument(filename, r, opts...)
	if err != nil {
		return types.Document{}, err
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debug("draft document")
		spew.Dump(draftDoc)
	}
	// apply document attribute substitutions and re-parse paragraphs that were affected
	blocks, err := applyDocumentAttributeSubstitutions(draftDoc.Blocks)
	if err != nil {
		return types.Document{}, err
	}
	// now, merge list items into proper lists
	blocks, err = rearrangeListItems(blocks, false)
	if err != nil {
		return types.Document{}, err
	}
	// now, rearrange elements in a hierarchical manner
	doc, err := rearrangeSections(blocks)
	if err != nil {
		return types.Document{}, err
	}
	// now, add front-matter attributes
	if len(draftDoc.FrontMatter.Content) > 0 {
		for k, v := range draftDoc.FrontMatter.Content {
			doc.Attributes[k] = v
		}
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debug("final document")
		spew.Dump(doc)
	}
	return doc, nil
}

// applyAttributeSubstitutions applies the document attribute substitutions
// and re-parse the paragraphs that were affected
func applyDocumentAttributeSubstitutions(blocks []interface{}) ([]interface{}, error) {
	// the document attributes, as they are resolved while processing the blocks
	attrs := make(map[string]string)
	result := make([]interface{}, 0, len(blocks)) // maximum capacity cannot exceed initial input
	for _, b := range blocks {
		switch b := b.(type) {
		case types.DocumentAttributeDeclaration:
			attrs[b.Name] = b.Value
		case types.DocumentAttributeReset:
			delete(attrs, b.Name)
		case types.Paragraph:
			for i, line := range b.Lines {
				if line, found := line.ApplyDocumentAttributeSubstitutions(attrs); found {
					// reparse the string elements, looking for links
					elements := make(types.InlineElements, 0, 2*len(line))
					for _, element := range line {
						switch element := element.(type) {
						case types.StringElement:
							r, err := parseInlineLinks(element)
							if err != nil {
								return []interface{}{}, errors.Wrap(err, "unable to process attribute substitutions")
							}
							elements = append(elements, r...)
						default:
							elements = append(elements, element)
						}
					}
					b.Lines[i] = elements
				} else {
					b.Lines[i] = line
				}
			}
		}
		result = append(result, b)
	}

	return result, nil
}

func parseInlineLinks(element types.StringElement) (types.InlineElements, error) {
	log.Debugf("parsing '%+v'", element.Content)
	elements, err := ParseReader("", strings.NewReader(element.Content), Entrypoint("InlineLinks"))
	if err != nil {
		return types.InlineElements{}, errors.Wrap(err, "error while parsing content for inline links")
	}

	log.Debugf("  giving '%+v'", elements)
	return elements.(types.InlineElements), nil
}

// rearrangeListItems moves the list items into lists, and nested lists if needed
func rearrangeListItems(blocks []interface{}, withinDelimitedBlock bool) ([]interface{}, error) {
	// log.Debugf("rearranging list items in %d blocks...", len(blocks))
	result := make([]interface{}, 0, len(blocks)) // maximum capacity cannot exceed initial input
	lists := []types.List{}                       // at each level (or depth), we have a list, whatever its type.
	blankline := false                            // track if the previous block was a blank line
	for _, block := range blocks {
		switch block := block.(type) {
		case types.DelimitedBlock:
			// process and replace the elements within this delimited block
			elements, err := rearrangeListItems(block.Elements, true)
			if err != nil {
				return nil, errors.Wrapf(err, "unable to rearrange list items in delimited block")
			}
			block.Elements = elements
			if len(lists) > 0 {
				switch list := lists[0].(type) { // just add the top-level list
				case *types.OrderedList:
					result = append(result, *list)
				case *types.UnorderedList:
					result = append(result, *list)
				case *types.LabeledList:
					result = append(result, *list)
				}
				// reset the list for further usage while processing the rest of the document
				lists = []types.List{}
			}
			result = append(result, block)
		case types.OrderedListItem, types.UnorderedListItem, types.LabeledListItem:
			// there's a special case: if the next list item has attributes and was preceded by a
			// blank line, then we need to start a new list
			if blankline && len(block.(types.DocumentElement).GetAttributes()) > 0 {
				if len(lists) > 0 {
					for _, list := range pruneLists(lists, 0) {
						result = append(result, unPtr(list))
					}
					// reset the list for further usage while processing the rest of the document
					lists = []types.List{}
				}
			}
			var err error
			lists, err = appendListItem(lists, block)
			if err != nil {
				return nil, errors.Wrapf(err, "unable to rearrange list items in delimited block")
			}
			blankline = false
		case types.ContinuedListItemElement:
			lists = appendContinuedListItemElement(lists, block)
			blankline = false
		case types.BlankLine:
			// blank lines are not part of the resulting Document sections (or top-level), but they are part of the delimited blocks
			// in some cases, they can also be used to split lists apart (when the next item has some attributes,
			// or if the next block is a comment)
			if withinDelimitedBlock && len(lists) == 0 { // only retain blank lines if within a delimited block, but not currently dealing with a list (or a set of nested lists)
				result = append(result, block)
			}
			blankline = true
		default:
			blankline = false
			// an block which is not a list item was found.
			// the first thing to do is to process the pending list items,
			// then only append this block to the result
			if len(lists) > 0 {
				log.Debugf("appending %d lists before processing element of type %T", len(lists), block)
				for _, list := range pruneLists(lists, 0) {
					result = append(result, unPtr(list))
				}
				// reset the list for further usage while processing the rest of the document
				lists = []types.List{}
			}
			result = append(result, block)
		}
	}
	// also when all is done, process the remaining pending list items
	if len(lists) > 0 {
		log.Debugf("processing the remaining %d lists...", len(lists))
		for _, list := range pruneLists(lists, 0) {
			result = append(result, unPtr(list))
		}
	}
	return result, nil
}

func unPtr(value interface{}) interface{} {
	v := reflect.ValueOf(value)
	k := v.Kind()
	if k == reflect.Ptr && v.Elem().IsValid() {
		return v.Elem().Interface()
	}
	return value
}

func appendListItem(lists []types.List, item interface{}) ([]types.List, error) {
	switch item := item.(type) {
	case types.OrderedListItem:
		return appendOrderedListItem(lists, &item)
	case types.UnorderedListItem:
		return appendUnorderedListItem(lists, &item)
	case types.LabeledListItem:
		return appendLabeledListItem(lists, item)
	}
	return lists, nil
}

func appendOrderedListItem(lists []types.List, item *types.OrderedListItem) ([]types.List, error) {
	maxLevel := 0
	log.Debugf("looking-up list for ordered list having items with level=%d and number style=%v", item.Level, item.NumberingStyle)
	for i, list := range lists {
		if list, ok := list.(*types.OrderedList); ok {
			// assume we can't have empty lists
			maxLevel++
			if list.Items[0].NumberingStyle == item.NumberingStyle {
				log.Debugf("found a matching ordered list at level %d", list.Items[0].Level)
				// prune items of "deeper/lower" level
				lists = pruneLists(lists, i)
				// apply the same level
				item.Level = list.Items[0].Level
				list.AddItem(*item)
				// also, prune the pointers to the remaining sublists (in case there is any...)
				return lists, nil
			}
		}
	}
	// force the current item level to (last seen level + 1)
	item.Level = maxLevel + 1
	// no match found: create a new list and if needed, adjust the level of the item
	log.Debugf("adding a new ordered list")
	list := types.NewOrderedList(item)
	// also, force the current item level to (last seen level + 1)
	item.Level = maxLevel + 1
	// also, attach this list to the one above, if it exists ;)
	// if len(lists) > 0 {
	// 	parentList := &(lists[len(lists)-1])
	// 	parentItem := (*parentList).LastItem()
	// 	parentItem.AddElement(list)
	// 	return append(lists, list), nil
	// }
	return append(lists, list), nil
}

func appendUnorderedListItem(lists []types.List, item *types.UnorderedListItem) ([]types.List, error) {
	maxLevel := 0
	log.Debugf("looking-up list for unordered list item with level=%d and bullet style=%v", item.Level, item.BulletStyle)
	for i, list := range lists {
		if list, ok := list.(*types.UnorderedList); ok {
			// assume we can't have empty lists
			maxLevel++
			if list.Items[0].BulletStyle == item.BulletStyle {
				log.Debugf("found a matching unordered list at level %d", list.Items[0].Level)
				// prune items of "deeper/lower" level
				lists = pruneLists(lists, i)
				// apply the same level
				item.Level = list.Items[0].Level
				list.AddItem(*item)
				return lists, nil
			}
		}
	}
	// no match found: create a new list and if needed, adjust the level of the item
	log.Debugf("adding a new unordered list")
	// also, force the current item level to (last seen level + 1)
	item.Level = maxLevel + 1
	// also, force the bullet-style based on the list on the level above (if it exists)
	if len(lists) > 0 {
		parentList := &(lists[len(lists)-1])
		parentItem := (*parentList).LastItem()
		// also, force the bullet style
		if parentItem, ok := parentItem.(*types.UnorderedListItem); ok {
			item.BulletStyle = item.BulletStyle.NextLevel(parentItem.BulletStyle)
		}
	}
	list := types.NewUnorderedList(item)
	return append(lists, list), nil
}

func appendLabeledListItem(lists []types.List, item types.LabeledListItem) ([]types.List, error) {
	maxLevel := 0
	log.Debugf("looking-up list for labeled list item with level=%d and term=%s", item.Level, item.Term)
	for i, list := range lists {
		log.Debugf("  comparing with list of type %T at level %d", list, i)
		if list, ok := list.(*types.LabeledList); ok {
			// assume we can't have empty lists
			maxLevel++
			log.Debugf("  comparing with list item level %d vs %d", list.Items[0].Level, item.Level)
			if list.Items[0].Level == item.Level {
				log.Debugf("found a matching labeled list")
				lists = pruneLists(lists, i)
				list.AddItem(item)
				log.Debugf("labeled list at level %d now has %d items", maxLevel, len(list.Items))
				return lists, nil
			}
		}
	}
	// no match found: create a new list and if needed, adjust the level of the item
	log.Debugf("adding a new labeled list")
	// also, force the current item level to (last seen level + 1)
	item.Level = maxLevel + 1
	list := types.NewLabeledList(item)
	return append(lists, list), nil
}

func appendContinuedListItemElement(lists []types.List, item types.ContinuedListItemElement) []types.List {
	lists = pruneLists(lists, len(lists)-1+item.Offset)
	log.Debugf("appending continued list item element with offset=%d (depth=%d)", item.Offset, len(lists))
	// lookup the list at which the item should be attached
	parentList := &(lists[len(lists)-1])
	parentItem := (*parentList).LastItem()
	parentItem.AddElement(item.Element)
	return lists
}

func pruneLists(lists []types.List, level int) []types.List {
	if level+1 < len(lists) {
		log.Debugf("pruning the list path from %d to %d level(s) deep", len(lists), level+1)
		// add the last list(s) as children of their parent, in reverse order,
		// because we copy the value, not the pointers
		for i := len(lists) - 1; i > level; i-- {
			log.Debugf("appending list at depth %d to the last item of the parent list...", (i + 1))
			parentList := &(lists[i-1])
			parentItem := (*parentList).LastItem()
			switch childList := lists[i].(type) {
			case *types.OrderedList:
				parentItem.AddElement(*childList)
			case *types.UnorderedList:
				parentItem.AddElement(*childList)
			case *types.LabeledList:
				parentItem.AddElement(*childList)
			}
		}
		// also, prune the pointers to the remaining sublists
		return lists[0 : level+1]
	}
	return lists
}

// rearrangeSections moves elements into section to obtain a hierarchical document instead of a flat thing
func rearrangeSections(blocks []interface{}) (types.Document, error) {

	// use same logic as with list items:
	// only append a child section to her parent section when
	// a sibling or higher level section is processed.

	log.Debugf("rearranging sections in %d blocks...", len(blocks))
	tle := make([]interface{}, 0, len(blocks)) // top-level elements
	sections := make([]types.Section, 0, 6)    // the path to the current section (eg: []{section-level0, section-level1, etc.})
	elementRefs := types.ElementReferences{}
	footnotes := types.Footnotes{}
	footnoteRefs := types.FootnoteReferences{}
	var previous *types.Section // the current "parent" section
	for _, element := range blocks {
		if e, ok := element.(types.Section); ok {
			// avoid duplicate IDs in sections
			referenceSection(e, elementRefs)
			if previous == nil { // set first parent
				log.Debugf("setting section with title %v as a top-level element", e.Title)
				sections = append(sections, e)
			} else if e.Level > previous.Level { // add new level
				log.Debugf("adding section with title %v as the first section at level %d", e.Title, e.Level)
				sections = append(sections, e)
			} else { // replace at the deepest level
				sections = pruneSections(sections, e.Level)
				if len(sections) > 0 && sections[0].Level == e.Level {
					log.Debugf("moving section with title %v as a new top-level element", e.Title)
					tle = append(tle, sections[0])
					sections = make([]types.Section, 0, 6)
				}
				log.Debugf("adding section with title %v as another section at level %d", e.Title, e.Level)
				sections = append(sections, e)
				// if len(sections) == 1 { // we have new top-level element
				// 	log.Debugf("setting section with title %v as secondary top-level", e.Title)
				// 	tle = append(tle, &e)
				// } else {
				// 	log.Debugf("adding section with title %v as child of section at level %d", e.Title, (len(sections) - 2))
				// 	sections[len(sections)-2].AddElement(e) // attach to parent
				// }
			}
			previous = &e // pointer to new current parent
		} else {
			if previous == nil {
				// log.Debugf("adding element of type %T as a top-level element", element)
				tle = append(tle, element)
			} else {
				parentSection := &(sections[len(sections)-1])
				// log.Debugf("adding element of type %T as a child of section with level %d", element, parentSection.Level)
				(*parentSection).AddElement(element)
			}
		}
		// also collect footnotes
		if e, ok := element.(types.FootnotesContainer); ok {
			// log.Debugf("collecting footnotes on element of type %T", element)
			f, fr, err := e.Footnotes()
			if err != nil {
				return types.Document{}, errors.Wrap(err, "unable to collect footnotes in document")
			}
			footnotes = append(footnotes, f...)
			for k, v := range fr {
				footnoteRefs[k] = v
			}
		}
	}
	// process the remaining sections
	sections = pruneSections(sections, 1)
	if len(sections) > 0 {
		tle = append(tle, sections[0])
	}

	return types.Document{
		Attributes:         types.DocumentAttributes{},
		Elements:           tle,
		ElementReferences:  elementRefs,
		Footnotes:          footnotes,
		FootnoteReferences: footnoteRefs,
	}, nil
}

func referenceSection(e types.Section, elementRefs types.ElementReferences) {
	id := e.Attributes.GetAsString(types.AttrID)
	for i := 1; ; i++ {
		var key string
		if i == 1 {
			key = id
		} else {
			key = id + "_" + strconv.Itoa(i)
		}
		if _, found := elementRefs[key]; !found {
			elementRefs[key] = e.Title
			// override the element id
			e.Attributes[types.AttrID] = key
			break
		}
	}
	elementRefs[e.Attributes.GetAsString(types.AttrID)] = e.Title
}

func pruneSections(sections []types.Section, level int) []types.Section {
	if len(sections) > 0 && level > 0 { // && level < len(sections) {
		log.Debugf("pruning the section path with %d level(s) of deep", len(sections))
		// add the last list(s) as children of their parent, in reverse order,
		// because we copy the value, not the pointers
		cut := len(sections)
		for i := len(sections) - 1; i > 0 && sections[i].Level >= level; i-- {
			parentSection := &(sections[i-1])
			log.Debugf("appending section at depth %d (%v) to the last element of the parent section (%v)", i, sections[i].Title, parentSection.Title)
			(*parentSection).AddElement(sections[i])
			cut = i
		}
		// also, prune the pointers to the remaining sublists
		sections := sections[0:cut]
		log.Debugf("sections list has now %d top-level elements", len(sections))
		return sections
	}
	return sections
}
