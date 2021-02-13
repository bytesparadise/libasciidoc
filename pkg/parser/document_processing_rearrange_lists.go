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
	log.Debug("rearranging list items...")
	a := &listArranger{
		blocks:               make([]interface{}, 0, len(blocks)),
		lists:                make([]types.List, 0, len(blocks)),
		blanklineCounter:     0,
		withinDelimitedBlock: withinDelimitedBlock,
	}
	var err error
	for _, block := range blocks {
		switch block := block.(type) {
		case types.ExampleBlock:
			if block.Elements, err = rearrangeListItems(block.Elements, true); err != nil {
				return nil, err
			}
			a.appendBlock(block)
		case types.QuoteBlock:
			if block.Elements, err = rearrangeListItems(block.Elements, true); err != nil {
				return nil, err
			}
			a.appendBlock(block)
		case types.SidebarBlock:
			if block.Elements, err = rearrangeListItems(block.Elements, true); err != nil {
				return nil, err
			}
			a.appendBlock(block)
		case types.OrderedListItem, types.UnorderedListItem, types.LabeledListItem, types.CalloutListItem:
			// there's a special case: if the next list item has attributes and was preceded by a
			// blank line, then we need to start a new list
			if a.blanklineCounter > 0 && len(block.(types.DocumentElement).GetAttributes()) > 0 {
				a.appendPendingLists()
			}
			if err = a.appendListItem(block); err != nil {
				return nil, errors.Wrapf(err, "unable to rearrange list items in delimited block")
			}
		case types.ContinuedListItemElement:
			block.Offset = a.blanklineCounter
			a.appendContinuedListItemElement(block)
		case types.BlankLine:
			a.appendBlankline(block)
		default:
			a.appendBlock(block)
		}
	}
	// also when all is done, process the remaining pending list items
	a.appendPendingLists()
	return a.blocks, nil
}

type listArranger struct {
	blocks               []interface{}
	lists                []types.List
	blanklineCounter     int
	withinDelimitedBlock bool
}

// an block which is not a list item was found.
// the first thing to do is to process the pending list items,
// then only append this block to the result
func (a *listArranger) appendBlock(block interface{}) {
	if len(a.lists) > 0 {
		a.pruneLists(0)
		for _, list := range a.lists {
			a.blocks = append(a.blocks, unPtr(list))
		}
		// reset the list for further usage while processing the rest of the document
		a.lists = []types.List{}
	}
	a.blocks = append(a.blocks, block)
}

func (a *listArranger) appendList(list types.List) {
	a.lists = append(a.lists, list)
}

func (a *listArranger) appendBlankline(l types.BlankLine) {
	// blank lines are not part of the resulting Document sections (or top-level), but they are part of the delimited blocks
	// in some cases, they can also be used to split lists apart (when the next item has some attributes,
	// or if the next block is a comment)
	if a.withinDelimitedBlock && len(a.lists) == 0 { // only retain blank lines if within a delimited block, but not currently dealing with a list (or a set of nested lists)
		a.appendBlock(l)
	}
	a.blanklineCounter++
}

func (a *listArranger) appendContinuedListItemElement(item types.ContinuedListItemElement) {
	item.Offset = a.blanklineCounter
	a.pruneLists(len(a.lists) - 1 - item.Offset)
	// log.Debugf("appending continued list item element with offset=%d (depth=%d)", item.Offset, len(a.lists))
	// lookup the list at which the item should be attached
	parentList := &(a.lists[len(a.lists)-1])
	parentItem := (*parentList).LastItem()
	parentItem.AddElement(item.Element)
	a.blanklineCounter = 0
}

func (a *listArranger) appendListItem(item interface{}) error {
	a.blanklineCounter = 0
	switch item := item.(type) {
	case types.OrderedListItem:
		return a.appendOrderedListItem(&item)
	case types.UnorderedListItem:
		return a.appendUnorderedListItem(&item)
	case types.LabeledListItem:
		return a.appendLabeledListItem(item)
	case types.CalloutListItem:
		return a.appendCalloutListItem(item)
	}
	return nil
}

func (a *listArranger) appendPendingLists() {
	if len(a.lists) > 0 {
		// log.Debugf("processing the remaining %d lists...", len(a.lists))
		a.pruneLists(0)
		for _, list := range a.lists {
			a.blocks = append(a.blocks, unPtr(list))
		}
		a.lists = []types.List{}
	}
}

func (a *listArranger) appendOrderedListItem(item *types.OrderedListItem) error {
	maxLevel := 0
	// log.Debugf("looking-up list for ordered list having items with level=%d and number style=%v", item.Level, item.Style)
	for i, list := range a.lists {
		if list, ok := list.(*types.OrderedList); ok {
			// assume we can't have empty lists
			maxLevel++
			if list.Items[0].Style == item.Style {
				// log.Debugf("found a matching ordered list at level %d", list.Items[0].Level)
				// prune items of "deeper/lower" level
				a.pruneLists(i)
				// apply the same level
				item.Level = list.Items[0].Level
				list.AddItem(*item)
				// also, prune the pointers to the remaining sublists (in case there is any...)
				return nil
			}
		}
	}
	// no match found: create a new list and if needed, adjust the level of the item
	// force the current item level to (last seen level + 1)
	item.Level = maxLevel + 1
	// log.Debugf("adding a new ordered list")
	a.appendList(types.NewOrderedList(item))
	return nil
}

func (a *listArranger) appendCalloutListItem(item types.CalloutListItem) error {
	for i, list := range a.lists {
		if list, ok := list.(*types.CalloutList); ok {
			// assume we can't have empty lists
			// log.Debugf("found a matching callout list")
			// prune items of "deeper/lower" level
			a.pruneLists(i)
			// apply the same level
			list.AddItem(item)
			// also, prune the pointers to the remaining sublists (in case there is any...)
			return nil
		}
	}
	// no match found: create a new list and if needed, adjust the level of the item
	// log.Debugf("adding a new callout list")
	a.appendList(types.NewCalloutList(item))
	return nil
}

func (a *listArranger) appendUnorderedListItem(item *types.UnorderedListItem) error {
	maxLevel := 0
	// log.Debugf("looking-up list for unordered list item with level=%d and bullet style=%v", item.Level, item.BulletStyle)
	for i, list := range a.lists {
		if list, ok := list.(*types.UnorderedList); ok {
			// assume we can't have empty lists
			maxLevel++
			if list.Items[0].BulletStyle == item.BulletStyle {
				// log.Debugf("found a matching unordered list at level %d", list.Items[0].Level)
				// prune items of "deeper/lower" level
				a.pruneLists(i)
				// apply the same level
				item.Level = list.Items[0].Level
				list.AddItem(*item)
				return nil
			}
		}
	}
	// no match found: create a new list and if needed, adjust the level of the item
	// log.Debugf("adding a new unordered list")
	// also, force the current item level to (last seen level + 1)
	item.Level = maxLevel + 1
	// also, force the bullet-style based on the list on the level above (if it exists)
	if len(a.lists) > 0 {
		parentList := &(a.lists[len(a.lists)-1])
		parentItem := (*parentList).LastItem()
		// also, force the bullet style
		if parentItem, ok := parentItem.(*types.UnorderedListItem); ok {
			item.BulletStyle = item.BulletStyle.NextLevel(parentItem.BulletStyle)
		}
	}
	a.appendList(types.NewUnorderedList(*item))
	return nil
}

func (a *listArranger) appendLabeledListItem(item types.LabeledListItem) error {
	// first, let's parse the labeled list item term for more complex content
	if len(item.Term) == 1 {
		if term, ok := item.Term[0].(types.StringElement); ok {
			var err error
			item.Term, err = parseLabeledListItemTerm(term.Content)
			if err != nil {
				return err
			}
		}
	}
	maxLevel := 0
	// log.Debugf("looking-up list for labeled list item with level=%d and term=%s", item.Level, item.Term)
	for i, list := range a.lists {
		// log.Debugf("  comparing with list of type %T at level %d", list, i)
		if list, ok := list.(*types.LabeledList); ok {
			// assume we can't have empty lists
			maxLevel++
			// log.Debugf("  comparing with list item level %d vs %d", list.Items[0].Level, item.Level)
			if list.Items[0].Level == item.Level {
				// log.Debugf("found a matching labeled list")
				a.pruneLists(i)
				list.AddItem(item)
				// log.Debugf("labeled list at level %d now has %d items", maxLevel, len(list.Items))
				return nil
			}
		}
	}
	// no match found: create a new list and if needed, adjust the level of the item
	// log.Debugf("adding a new labeled list")
	// also, force the current item level to (last seen level + 1)
	item.Level = maxLevel + 1
	a.appendList(types.NewLabeledList(item))
	return nil
}

// a labeled list item term may contain links, images, quoted text, footnotes, etc.
func parseLabeledListItemTerm(term string) ([]interface{}, error) {
	// result := []interface{}{}
	elements, err := ParseReader("", strings.NewReader(term), Entrypoint("LabeledListItemTerm"))
	if err != nil {
		return []interface{}{}, errors.Wrap(err, "error while parsing content for inline links")
	}
	return elements.([]interface{}), nil
}

func (a *listArranger) pruneLists(level int) {
	if level+1 < len(a.lists) {
		// log.Debugf("pruning the list path from %d to %d level(s) deep", len(a.lists), level+1)
		// add the last list(s) as children of their parent, in reverse order,
		// because we copy the value, not the pointers
		for i := len(a.lists) - 1; i > level; i-- {
			// log.Debugf("appending list at depth %d to the last item of the parent list...", (i + 1))
			parentList := &(a.lists[i-1])
			parentItem := (*parentList).LastItem()
			switch childList := a.lists[i].(type) {
			case *types.OrderedList:
				parentItem.AddElement(*childList)
			case *types.UnorderedList:
				parentItem.AddElement(*childList)
			case *types.LabeledList:
				parentItem.AddElement(*childList)
			}
		}
		// also, prune the pointers to the remaining sublists
		a.lists = a.lists[0 : level+1]
	}
}

func unPtr(value interface{}) interface{} {
	v := reflect.ValueOf(value)
	k := v.Kind()
	if k == reflect.Ptr && v.Elem().IsValid() {
		return v.Elem().Interface()
	}
	return value
}
