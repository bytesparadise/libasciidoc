package parser

import (
	"reflect"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// rearrangeListItems moves the list items into lists, and nested lists if needed
func rearrangeListItems(blocks []interface{}, withinDelimitedBlock bool) ([]interface{}, error) {
	// log.Debugf("rearranging list items in %d blocks...", len(blocks))
	result := make([]interface{}, 0, len(blocks)) // maximum capacity cannot exceed initial input
	lists := []types.List{}                       // at each level (or depth), we have a list, whatever its type.
	// track if the previous block was a blank line.
	// also, count the blanklines to determine the level of parent attachment when reaching a `ContinuedListItemElement`
	blanklineCount := 0
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
				case *types.CalloutList:
					result = append(result, *list)
				}
				// reset the list for further usage while processing the rest of the document
				lists = []types.List{}
			}
			result = append(result, block)
		case types.OrderedListItem, types.UnorderedListItem, types.LabeledListItem, types.CalloutListItem:
			// there's a special case: if the next list item has attributes and was preceded by a
			// blank line, then we need to start a new list
			if blanklineCount > 0 && len(block.(types.DocumentElement).GetAttributes()) > 0 {
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
			blanklineCount = 0
		case types.ContinuedListItemElement:
			block.Offset = blanklineCount
			lists = appendContinuedListItemElement(lists, block)
			blanklineCount = 0
		case types.BlankLine:
			// blank lines are not part of the resulting Document sections (or top-level), but they are part of the delimited blocks
			// in some cases, they can also be used to split lists apart (when the next item has some attributes,
			// or if the next block is a comment)
			if withinDelimitedBlock && len(lists) == 0 { // only retain blank lines if within a delimited block, but not currently dealing with a list (or a set of nested lists)
				result = append(result, block)
			}
			blanklineCount++
		default:
			blanklineCount = 0
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
	case types.CalloutListItem:
		return appendCalloutListItem(lists, item)
	}
	return lists, nil
}

func appendOrderedListItem(lists []types.List, item *types.OrderedListItem) ([]types.List, error) {
	maxLevel := 0
	log.Debugf("looking-up list for ordered list having items with level=%d and number style=%v", item.Level, item.Style)
	for i, list := range lists {
		if list, ok := list.(*types.OrderedList); ok {
			// assume we can't have empty lists
			maxLevel++
			if list.Items[0].Style == item.Style {
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
	return append(lists, list), nil
}

func appendCalloutListItem(lists []types.List, item types.CalloutListItem) ([]types.List, error) {
	for i, list := range lists {
		if list, ok := list.(*types.CalloutList); ok {
			// assume we can't have empty lists
			log.Debugf("found a matching callout list")
			// prune items of "deeper/lower" level
			lists = pruneLists(lists, i)
			// apply the same level
			list.AddItem(item)
			// also, prune the pointers to the remaining sublists (in case there is any...)
			return lists, nil
		}
	}
	// no match found: create a new list and if needed, adjust the level of the item
	log.Debugf("adding a new callout list")
	list := types.NewCalloutList(item)
	// also, force the current item level to (last seen level + 1)
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
	// first, let's parse the labeled list item term for more complex content
	if len(item.Term) == 1 {
		if term, ok := item.Term[0].(types.StringElement); ok {
			var err error
			item.Term, err = parseLabeledListItemTerm(term.Content)
			if err != nil {
				return nil, err
			}
		}
	}
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

// a labeled list item term may contain links, images, quoted text, footnotes, etc.
func parseLabeledListItemTerm(term string) ([]interface{}, error) {
	result := []interface{}{}
	elements, err := ParseReader("", strings.NewReader(term), Entrypoint("LabeledListItemTerm"))
	if err != nil {
		return []interface{}{}, errors.Wrap(err, "error while parsing content for inline links")
	}
	log.Debugf("parsed labeled list item term: '%+v'", elements)
	result = append(result, elements.([]interface{})...)
	return result, nil
}

func appendContinuedListItemElement(lists []types.List, item types.ContinuedListItemElement) []types.List {
	lists = pruneLists(lists, len(lists)-1-item.Offset)
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
