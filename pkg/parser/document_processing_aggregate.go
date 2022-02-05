package parser

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"

	log "github.com/sirupsen/logrus"
)

// Aggregate pipeline task which organizes the sections in hierarchy, and
// keeps track of their references.
// Also, takes care of wrapping all blocks between header (section 0) and first child section
// into a `Preamble`
// Also, takes care of inserting the Table of Contents
// returns the whole document at once (or an error)
func Aggregate(ctx *ParseContext, fragmentStream <-chan types.DocumentFragment) (*types.Document, *types.TableOfContents, error) {
	doc, toc, err := aggregate(ctx, fragmentStream)
	if err != nil {
		return nil, nil, err
	}
	insertPreamble(doc)
	insertTableOfContents(ctx, doc, toc)
	return doc, toc, nil
}

func aggregate(ctx *ParseContext, fragmentStream <-chan types.DocumentFragment) (*types.Document, *types.TableOfContents, error) {
	attrs := ctx.attributes
	refs := types.ElementReferences{}
	root := &types.Document{}

	lvls := &levels{
		root,
	}
	var toc *types.TableOfContents
	for f := range fragmentStream {
		if f.Error != nil {
			log.WithField("start_offset", f.Position.Start).WithField("end_offset", f.Position.End).Error(f.Error)
			continue
		}
		for _, element := range f.Elements {
			switch e := element.(type) {
			case *types.AttributeDeclaration:
				attrs.set(e.Name, e.Value)
				if e.Name == types.AttrTableOfContents {
					toc = types.NewTableOfContents(attrs.getAsIntWithDefault(types.AttrTableOfContentsLevels, 2))
				}
				// yet, retain the element, in case we need it during rendering (eg: `figure-caption`, etc.)
				if err := lvls.appendElement(e); err != nil {
					return nil, nil, err
				}
			case *types.FrontMatter:
				attrs.setAll(e.Attributes)
				if err := lvls.appendElement(e); err != nil {
					return nil, nil, err
				}
			case *types.DocumentHeader:
				// TODO: is it needed in this pipeline stage?
				for _, elmt := range e.Elements {
					switch attr := elmt.(type) {
					case *types.AttributeDeclaration:
						ctx.attributes.set(attr.Name, attr.Value)
					case *types.AttributeReset:
						ctx.attributes.unset(attr.Name)
					}
				}
				if attrs.has(types.AttrTableOfContents) {
					toc = types.NewTableOfContents(attrs.getAsIntWithDefault(types.AttrTableOfContentsLevels, 2))
				}
				if err := lvls.appendElement(e); err != nil {
					return nil, nil, err
				}
				// do not add header to ToC
			case *types.AttributeReset:
				attrs.unset(e.Name)
				// yet, retain the element, in case we need it during rendering (eg: `figure-caption`, etc.)
				if err := lvls.appendElement(e); err != nil {
					return nil, nil, err
				}
			case *types.BlankLine, *types.SingleLineComment:
				// ignore
			case *types.Section:
				if err := e.ResolveID(attrs.allAttributes(), refs); err != nil {
					return nil, nil, err
				}
				if err := lvls.appendSection(e); err != nil {
					return nil, nil, err
				}
				if toc != nil {
					toc.Add(e)
				}
			default:
				if err := lvls.appendElement(e); err != nil {
					return nil, nil, err
				}
			}
			// also, check if the element has refs
			if e, ok := element.(types.Referencable); ok {
				e.Reference(refs)
			}
		}
	}

	log.WithField("pipeline_task", "aggregate").Debug("done")
	if len(refs) > 0 {
		root.ElementReferences = refs
	}
	return root, toc, nil
}

type levels []types.WithElementAddition

func (l *levels) appendSection(s *types.Section) error {
	// note: section levels start at 0, but first level is root (doc)
	if idx, found := l.indexOfParent(s); found {
		*l = (*l)[:idx+1] // trim to parent level
	}
	log.Debugf("adding section with level %d at position %d in levels", s.Level, len(*l))
	// append
	if err := (*l)[len(*l)-1].AddElement(s); err != nil {
		return err
	}
	*l = append(*l, s)
	return nil
}

// return the index of the parent element for the given section,
// taking account the given section's level, and also gaps in other
// sections (eg: `1,2,4` instead of `0,1,2`)
func (l *levels) indexOfParent(s *types.Section) (int, bool) {
	for i, e := range *l {
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

func (l *levels) appendElement(e interface{}) error {
	return (*l)[len(*l)-1].AddElement(e)
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
	if header := doc.Header(); header != nil {
		log.Debug("inserting preamble after header")
		elements[0] = header
		elements[1] = preamble
		copy(elements[2:], doc.Elements[1+len(preamble.Elements):])
	} else {
		log.Debug("inserting preamble at beginning of document")
		elements[0] = preamble
		copy(elements[1:], doc.Elements[len(preamble.Elements):])
	}
	doc.Elements = elements
}

func newPreamble(doc *types.Document) *types.Preamble {
	if doc.Header() == nil {
		log.Debug("skipping preamble: no header in doc")
		return nil
	}
	preamble := &types.Preamble{
		Elements: make([]interface{}, 0, len(doc.Elements)),
	}
	for _, e := range doc.Elements {
		switch e.(type) {
		case *types.DocumentHeader:
			continue
		case *types.Section:
			return preamble
		default:
			preamble.Elements = append(preamble.Elements, e)
		}
	}
	return nil
}

func insertTableOfContents(ctx *ParseContext, doc *types.Document, toc *types.TableOfContents) {
	if toc == nil {
		log.Debug("no table of contents to insert")
		return
	}
	if ctx.attributes.has(types.AttrTableOfContents) {
		doc.TableOfContents = toc
	}
}
