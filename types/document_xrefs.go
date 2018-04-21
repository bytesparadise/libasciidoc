package types

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// ElementReferences the element references in the document
type ElementReferences map[string]interface{}

// ElementReferencesCollector the visitor that traverses the whole document structure in search for elements with an ID
type ElementReferencesCollector struct {
	ElementReferences ElementReferences
}

// NewElementReferencesCollector initializes a new ElementReferencesCollector
func NewElementReferencesCollector() *ElementReferencesCollector {
	return &ElementReferencesCollector{
		ElementReferences: ElementReferences{},
	}
}

// BeforeVisit Implements Visitable#BeforeVisit()
func (c *ElementReferencesCollector) BeforeVisit(element Visitable) error {
	return nil
}

// Visit Implements Visitable#Visit()
func (c *ElementReferencesCollector) Visit(element Visitable) error {
	switch e := element.(type) {
	case Section:
		elementID := e.Title.Attributes[AttrID]
		if elementID, ok := elementID.(string); ok {
			log.Debugf("Adding element reference: %v", elementID)
			c.ElementReferences[elementID] = e.Title
		} else {
			return errors.Errorf("unexpected type of element id: %T", elementID)
		}
	}
	return nil
}

// AfterVisit Implements Visitable#AfterVisit()
func (c *ElementReferencesCollector) AfterVisit(element Visitable) error {
	return nil
}
