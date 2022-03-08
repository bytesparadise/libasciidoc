package plugins

import (
  "fmt"

  "github.com/davecgh/go-spew/spew"
  "github.com/bytesparadise/libasciidoc/pkg/types"
)

type Plugin struct {
    name string
    apply func(*types.Document) (*types.Document)
}

var plugin_registry = [...]Plugin{Plugin{"diagram", DiagramApply}}

func Apply(doc *types.Document) (*types.Document) {
  for _, plugin := range plugin_registry {
    doc = plugin.apply(doc)
  }
  return doc
}

func DiagramApply(doc *types.Document) (*types.Document) {
  spew.Dump(doc)
  for _, element := range doc.Elements {
    fmt.Println("===")
    spew.Dump(element)
    switch elem := element.(type) {
    case types.WithElements:
        spew.Dump(elem)
        fmt.Println("Has elements")
    /*case *types.DelimitedBlock:
        if types.HasAttributeWithValue(elem.Attributes, "@positional-1", "plantuml") {
//          path := "test.jpg"
          path := []interface{}{"test.jpg"}
          location, _ := types.NewLocation("", path)
          doc.Elements[i], _ = types.NewImageBlock(location, elem.Attributes)
          spew.Dump(elem)
        } */
    }
  }
  return doc
}
