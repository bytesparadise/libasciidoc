package parser

import (
	"time"

	"github.com/bytesparadise/libasciidoc/pkg/types"

	log "github.com/sirupsen/logrus"
)

// Aggregate pipeline task which organizes the sections in hierarchy, and
// keeps track of their references.
// Also, takes care of wrapping all blocks between header (section 0) and first child section
// into a `Preamble`
// Also, takes care of inserting the Table of Contents
// returns the whole document at once (or an error)
func Aggregate(ctx *ParseContext, fragmentStream <-chan types.DocumentFragment) (*types.Document, error) {
	doc, err := aggregate(ctx, fragmentStream)
	if err != nil {
		return nil, err
	}
	insertPreamble(doc)
	return doc, nil
}

func aggregate(ctx *ParseContext, fragmentStream <-chan types.DocumentFragment) (*types.Document, error) {
	attrs := ctx.attributes
	refs := types.ElementReferences{}
	doc := &types.Document{}
	// TODO: update `toc.MaxDepth` when `AttrTableOfContentsLevels` is declared afterwards
	toc := types.NewTableOfContents(attrs.getAsIntWithDefault(types.AttrTableOfContentsLevels, 2))

	a := &aggregator{
		doc,
	}
	for f := range fragmentStream {
		if f.Error != nil {
			log.WithField("start_offset", f.Position.Start).WithField("end_offset", f.Position.End).Error(f.Error)
			continue
		}
		start := time.Now()
		for _, element := range f.Elements {
			switch e := element.(type) {
			case *types.AttributeDeclaration:
				attrs.set(e.Name, e.Value)
				if e.Name == types.AttrTableOfContentsLevels {
					// TODO: raise a warning if value is invalid
					maxDepth := attrs.getAsIntWithDefault(types.AttrTableOfContentsLevels, 2)
					log.Debugf("setting ToC.MaxDepth to %d", maxDepth)
					toc.MaxDepth = maxDepth
				}
			case *types.FrontMatter:
				attrs.setAll(e.Attributes)
			case *types.DocumentHeader:
				for _, elmt := range e.Elements {
					switch attr := elmt.(type) {
					case *types.AttributeDeclaration:
						ctx.attributes.set(attr.Name, attr.Value)
						if attr.Name == types.AttrTableOfContentsLevels {
							// TODO: raise a warning if value is invalid
							maxDepth := attrs.getAsIntWithDefault(types.AttrTableOfContentsLevels, 2)
							log.Debugf("setting ToC.MaxDepth to %d", maxDepth)
							toc.MaxDepth = maxDepth
						}
					case *types.AttributeReset:
						ctx.attributes.unset(attr.Name)
					}
				}
				// do not add header to ToC
			case *types.AttributeReset:
				attrs.unset(e.Name)
			case *types.BlankLine, *types.SinglelineComment:
				// ignore
			case *types.Section:
				if err := e.ResolveID(attrs.allAttributes(), refs); err != nil {
					return nil, err
				}
				if toc != nil {
					toc.Add(e)
				}
			}

			// also, retain the element
			// yet, retain the element, in case we need it during rendering (eg: `figure-caption`, etc.)
			if err := a.append(element); err != nil {
				return nil, err
			}
			// also, check if the element has refs
			if e, ok := element.(types.Referencable); ok {
				e.Reference(refs)
			}
		}
		log.Debugf("time to aggregate fragment at %d: %d microseconds", f.Position.Start, time.Since(start).Microseconds())
	}
	if len(refs) > 0 {
		doc.ElementReferences = refs
		// also, resolve cross references (only needed if there are referenced elements)
		for _, e := range doc.Elements {
			if err := resolveCrossReferences(e, attrs); err != nil {
				return nil, err
			}
		}
	}
	if len(toc.Sections) > 0 {
		doc.TableOfContents = toc
	}
	log.WithField("pipeline_task", "aggregate").Debug("done")
	return doc, nil
}

func resolveCrossReferences(element interface{}, attrs *contextAttributes) error {
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("resolving cross references in element of type '%T'", element)
	}
	switch e := element.(type) {
	case types.WithElements:
		for _, elmt := range e.GetElements() {
			if err := resolveCrossReferences(elmt, attrs); err != nil {
				return err
			}
		}
	case *types.InternalCrossReference:
		if err := e.ResolveID(attrs.allAttributes()); err != nil {
			return err
		}
	}
	return nil
}

type aggregator []types.WithElementAddition

func (a *aggregator) append(e interface{}) error {
	switch e := e.(type) {
	case *types.Section:
		return a.appendSection(e)
	default:
		return a.appendElement(e)
	}
}

func (a *aggregator) appendElement(e interface{}) error {
	return (*a)[len(*a)-1].AddElement(e)
}

func (a *aggregator) appendSection(s *types.Section) error {
	// note: section levels start at 0, but first level is root (doc)
	if idx, found := a.indexOfParent(s); found {
		*a = (*a)[:idx+1] // trim to parent level
	}
	log.Debugf("adding section with level %d at position %d in levels", s.Level, len(*a))
	// append
	if err := (*a)[len(*a)-1].AddElement(s); err != nil {
		return err
	}
	*a = append(*a, s)
	return nil
}

// return the index of the parent element for the given section,
// taking account the given section's level, and also gaps in other
// sections (eg: `1,2,4` instead of `0,1,2`)
func (a *aggregator) indexOfParent(s *types.Section) (int, bool) {
	for i, e := range *a {
		if p, ok := e.(*types.Section); ok {
			if p.Level >= s.Level {
				log.Debugf("found parent at index %d for section with level %d", i-1, s.Level)
				return i - 1, true // return previous
			}
		}
	}
	//
	return -1, false
}

func insertPreamble(doc *types.Document) {
	preamble := newPreamble(doc)
	// if no element in the preamble, or if no section in the document,
	// or if all elements are AttributeDeclaration/AttributeReset and nothing else
	// then no preamble to insert
	if preamble == nil || !preamble.HasContent() {
		log.Debugf("no preamble to insert")
		return
	}
	// now, insert the preamble instead of the 'n' blocks that belong to the preamble
	// and copy the other items
	elements := make([]interface{}, len(doc.Elements)-len(preamble.Elements)+1)
	if frontmatter := doc.FrontMatter(); frontmatter != nil {
		elements[0] = frontmatter
	}
	if header, offset := doc.Header(); header != nil {
		log.Debug("inserting preamble after header")
		elements[0+offset] = header
		elements[1+offset] = preamble
		copy(elements[2+offset:], doc.Elements[1+len(preamble.Elements)+offset:])
	} else {
		log.Debug("inserting preamble at beginning of document")
		elements[0] = preamble
		copy(elements[1:], doc.Elements[len(preamble.Elements):])
	}
	doc.Elements = elements
}

func newPreamble(doc *types.Document) *types.Preamble {
	if header, _ := doc.Header(); header == nil || header.Title == nil {
		log.Debug("skipping preamble: no header in doc")
		return nil
	}
	preamble := &types.Preamble{
		Elements: make([]interface{}, 0, len(doc.Elements)),
	}
	for _, e := range doc.Elements {
		switch e.(type) {
		case *types.DocumentHeader, *types.FrontMatter:
			continue
		case *types.Section:
			return preamble
		default:
			preamble.Elements = append(preamble.Elements, e)
		}
	}
	return nil
}
