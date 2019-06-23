package parser

import (
	"io"
	"strconv"

	"github.com/davecgh/go-spew/spew"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// ParseDocument parses the content of the reader identitied by the filename
func ParseDocument(filename string, r io.Reader, opts ...Option) (*types.Document, error) {
	preflightDoc, err := ParsePreflightDocument(filename, r, opts...)
	if err != nil {
		return nil, err
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debug("preflight document")
		spew.Dump(preflightDoc)
	}
	// now, merge list items into proper lists
	blocks, err := rearrangeListItems(preflightDoc.Blocks, false)
	if err != nil {
		return nil, err
	}
	// now, rearrange elements in a hierarchical manner
	doc, err := rearrangeSections(blocks)
	if err != nil {
		return nil, err
	}
	// now, add front-matter attributes
	if preflightDoc.FrontMatter != nil {
		for k, v := range preflightDoc.FrontMatter.Content {
			doc.Attributes[k] = v
		}
	}
	return doc, nil
}

// rearrangeListItems moves the list items into lists, and nested lists if needed
func rearrangeListItems(blocks []interface{}, withinDelimitedBlock bool) ([]interface{}, error) {
	result := make([]interface{}, 0, len(blocks)) // maximum capacity cannot exceed initial input
	lists := []types.List{}                       // at each level (or depth), we have a list, whatever its type.
	blankline := false                            // track if the previous block was a blank line
	for _, block := range blocks {
		switch block := block.(type) {
		case *types.DelimitedBlock:
			// process and replace the elements within this delimited block
			elements, err := rearrangeListItems(block.Elements, true)
			if err != nil {
				return nil, errors.Wrapf(err, "unable to rearrange list items in delimited block")
			}
			block.Elements = elements
			if len(lists) > 0 {
				result = append(result, lists[0]) // just add the top-level list
				// reset the list for further usage while processing the rest of the document
				lists = []types.List{}
			}
			result = append(result, block)
		case types.ListItem:
			// there's a special case: if the next list item has attributes and was preceeded by a
			// blank line, then we need to start a new list
			if blankline && len(block.GetAttributes()) > 0 {
				if len(lists) > 0 {
					result = append(result, lists[0]) // just add the top-level list
					// reset the list for further usage while processing the rest of the document
					lists = []types.List{}
				}
			}
			lists = appendListItem(lists, block)
			blankline = false
		case *types.ContinuedListItemElement:
			lists = appendContinuedListItemElement(lists, block)
			blankline = false
		case *types.BlankLine:
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
				result = append(result, lists[0]) // just add the top-level list
				// reset the list for further usage while processing the rest of the document
				lists = []types.List{}
			}
			result = append(result, block)
		}
	}
	// also when all is done, process the remaining pending list items
	if len(lists) > 0 {
		result = append(result, lists[0]) // just add the top-level list
	}
	return result, nil
}

func appendListItem(lists []types.List, item types.ListItem) []types.List {
	switch item := item.(type) {
	case *types.OrderedListItem:
		return appendOrderedListItem(lists, item)
	case *types.UnorderedListItem:
		return appendUnorderedListItem(lists, item)
	case *types.LabeledListItem:
		return appendLabeledListItem(lists, item)
	}
	return lists
}

func appendOrderedListItem(lists []types.List, item *types.OrderedListItem) []types.List {
	maxLevel := 0
	log.Debugf("looking-up list for ordered list item with level=%d and number style=%v", item.Level, item.NumberingStyle)
	for i, list := range lists {
		if list, ok := list.(*types.OrderedList); ok {
			// assume we can't have empty lists
			maxLevel++
			if list.Items[0].NumberingStyle == item.NumberingStyle {
				log.Debugf("found a matching ordered list")
				if lastItem, ok := list.LastItem().(*types.OrderedListItem); ok {
					item.Position = lastItem.Position + 1
				}
				list.AddItem(item)
				// apply the same level
				item.Level = list.Items[0].Level
				// also, prune the pointers to the remaining sublists
				return prune(lists, i)
			}
		}
	}
	// no match found: create a new list and if needed, adjust the level of the item
	log.Debugf("adding a new ordered list")
	list := types.NewOrderedList(item)
	// also, force the current item level to (last seen level + 1)
	item.Level = maxLevel + 1
	// also, attach this list to the one above, if it exists ;)
	if len(lists) > 0 {
		parentList := lists[len(lists)-1]
		parentItem := parentList.LastItem()
		parentItem.AddElement(list)
	}
	return append(lists, list)
}

func appendUnorderedListItem(lists []types.List, item *types.UnorderedListItem) []types.List {
	maxLevel := 0
	log.Debugf("looking-up list for unordered list item with level=%d and bullet style=%v", item.Level, item.BulletStyle)
	for i, list := range lists {
		if list, ok := list.(*types.UnorderedList); ok {
			// assume we can't have empty lists
			maxLevel++
			if list.Items[0].BulletStyle == item.BulletStyle {
				log.Debugf("found a matching unordered list")
				list.AddItem(item)
				// apply the same level
				item.Level = list.Items[0].Level
				return prune(lists, i)
			}
		}
	}
	// no match found: create a new list and if needed, adjust the level of the item
	log.Debugf("adding a new unordered list")
	list := types.NewUnorderedList(item)
	// also, force the current item level to (last seen level + 1)
	item.Level = maxLevel + 1
	// also, move this first item's attributes to the parent list level
	// also, attach this list to the one above, if it exists ;)
	if len(lists) > 0 {
		parentList := lists[len(lists)-1]
		parentItem := parentList.LastItem()
		parentItem.AddElement(list)
		// also, force the bullet style
		if parentItem, ok := parentItem.(*types.UnorderedListItem); ok {
			item.BulletStyle = item.BulletStyle.NextLevel(parentItem.BulletStyle)
		}
	}
	return append(lists, list)
}

func appendLabeledListItem(lists []types.List, item *types.LabeledListItem) []types.List {
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
				list.AddItem(item)
				log.Debugf("labeled list at level %d now has %d items", maxLevel, len(list.Items))
				return prune(lists, i)
			}
		}
	}
	// no match found: create a new list and if needed, adjust the level of the item
	log.Debugf("adding a new labeled list")
	list := types.NewLabeledList(item)
	// also, force the current item level to (last seen level + 1)
	item.Level = maxLevel + 1
	// also, attach this list to the one above, if it exists ;)
	if len(lists) > 0 {
		parentList := lists[len(lists)-1]
		parentItem := parentList.LastItem()
		parentItem.AddElement(list)
	}
	return append(lists, list)
}

func prune(lists []types.List, level int) []types.List {
	// also, prune the pointers to the remaining sublists
	if level+1 < len(lists) {
		log.Debugf("pruning the list path from %d to %d levels", len(lists), level+1)
		return lists[0 : level+1]
	}
	return lists
}

func appendContinuedListItemElement(lists []types.List, item *types.ContinuedListItemElement) []types.List {
	// lookup the list at which the item should be attached
	parentList := lists[len(lists)-1+item.Offset]
	parentItem := parentList.LastItem()
	parentItem.AddElement(item.Element)
	return lists
}

// rearrangeSections moves elements into section to obtain a hierarchical document instead of a flat thing
func rearrangeSections(blocks []interface{}) (*types.Document, error) {
	elements := make([]interface{}, 0, len(blocks)) // allocate enough room
	references := types.ElementReferences{}
	footnotes := types.Footnotes{}
	footnoteRefs := types.FootnoteReferences{}
	sections := make([]*types.Section, 0, 6) // the path to the current section (eg: []{section-level0, section-level1, etc.})
	var parent *types.Section                // the current "parent" section
	for _, element := range blocks {
		if e, ok := element.(*types.Section); ok {
			// avoid duplicate IDs in sections
			id := e.Attributes.GetAsString(types.AttrID)
			for i := 1; ; i++ {
				var key string
				if i == 1 {
					key = id
				} else {
					key = id + "_" + strconv.Itoa(i)
				}
				if _, found := references[key]; !found {
					references[key] = e.Title
					// override the element id
					e.Attributes[types.AttrID] = key
					break
				}
			}
			references[e.Attributes.GetAsString(types.AttrID)] = e.Title
			if parent == nil { // set first parent
				sections = append(sections, e)
				elements = append(elements, e)
			} else if e.Level == parent.Level { // replace at the deepest level
				sections[len(sections)-1] = e
				if len(sections) == 1 { // we have new top-level element
					elements = append(elements, e)
				} else {
					sections[len(sections)-2].Elements = append(sections[len(sections)-2].Elements, e) // attach to parent
				}
			} else if e.Level > parent.Level { // add new level
				sections = append(sections, e)
				parent.Elements = append(parent.Elements, e)
			} else { // remove all levels below current section and set this section instead
				for i := 0; i < len(sections); i++ {
					if sections[i].Level >= e.Level {
						sections = sections[0 : i+1]
						sections[i] = e
						if i == 0 { // in case we have new top-level element
							elements = append(elements, element)
						} else {
							sections[i-1].Elements = append(sections[i-1].Elements, e)
						}
						break
					}
				}
			}
			parent = e // pointer to new current parent
		} else {
			if parent == nil {
				elements = append(elements, element)
			} else {
				parent.Elements = append(parent.Elements, element)
			}
		}
		// also collect footnotes
		if e, ok := element.(types.FootnotesContainer); ok {
			f, fr, err := e.Footnotes()
			if err != nil {
				return nil, errors.Wrap(err, "unable to collect footnotes in document")
			}
			footnotes = append(footnotes, f...)
			for k, v := range fr {
				footnoteRefs[k] = v
			}
		}
	}
	return &types.Document{
		Attributes:         types.DocumentAttributes{},
		Footnotes:          footnotes,
		FootnoteReferences: footnoteRefs,
		ElementReferences:  references,
		Elements:           elements,
	}, nil
}
