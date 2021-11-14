package parser

// // processes the footnotes in the blocks, replaces them with `FootnoteReference`
// // and keep them in a separate `Footnotes`
// func processFootnotes(blocks []interface{}) ([]interface{}, []*types.Footnote) {
// 	logrus.Debug("processing footnotes...")
// 	footnotes := types.NewFootnotes()
// 	for _, block := range blocks {
// 		if c, ok := block.(types.WithFootnotes); ok {
// 			c.SubstituteFootnotes(footnotes)
// 		}
// 	}
// 	return blocks, footnotes.Notes()
// }
