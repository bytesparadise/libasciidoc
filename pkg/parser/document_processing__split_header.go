package parser

// func SplitHeader(done <-chan interface{}, fragmentStream <-chan types.DocumentFragment) <-chan types.DocumentFragment {
// 	resultStream := make(chan types.DocumentFragment, 1)
// 	go func() {
// 		defer close(resultStream)
// 		for fragment := range fragmentStream {
// 			for _, f := range splitHeader(fragment) {
// 				select {
// 				case <-done:
// 					log.WithField("pipeline_task", "split_header").Debug("received 'done' signal")
// 					return
// 				case resultStream <- f:
// 				}
// 			}
// 		}
// 		log.WithField("pipeline_task", "split_header").Debug("done")
// 	}()
// 	return resultStream
// }

// func splitHeader(f types.DocumentFragment) []types.DocumentFragment {
// 	if err := f.Error; err != nil {
// 		log.Debugf("skipping element splitting because of fragment with error: %v", f.Error)
// 		return []types.DocumentFragment{f}
// 	}
// 	result := make([]types.DocumentFragment, 0, len(f.Elements))
// 	for _, element := range f.Elements {
// 		switch e := element.(type) {
// 		case *types.DocumentHeader:
// 			result = append(result, types.NewDocumentFragment(e.Section))
// 			// if e.Authors != nil {
// 			// 	result = append(result, types.NewDocumentFragment(e.Authors))
// 			// }
// 			// if e.Revision != nil {
// 			// 	result = append(result, types.NewDocumentFragment(e.Revision))
// 			// }
// 			for _, a := range e.Attributes {
// 				result = append(result, types.NewDocumentFragment(a))
// 			}
// 		default:
// 			result = append(result, types.NewDocumentFragment(e))
// 		}
// 	}
// 	return result
// }
