// +build stats

package parser_test

import (
	"encoding/json"
	"fmt"

	"github.com/bytesparadise/libasciidoc/pkg/parser"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" // nolint:golint
	. "github.com/onsi/gomega" // nolint:golint
)

var _ = Describe("paragraphs", func() {

	It("multiline paragraph with rich content", func() {
		source := `:fds-version: 6.7.4
	
PyroSim is a graphical user interface for the https://github.com/firemodels/fds/releases/tag/FDS{fds-version}/[Fire Dynamics Simulator (FDS) version {fds-version}].
FDS is closely integrated into PyroSim.
FDS models can predict smoke, temperature, carbon monoxide, and other substances during fires.
The results of these simulations are used to ensure the safety of buildings before construction, evaluate safety options of existing buildings, reconstruct fires for post-accident investigation, and assist in firefighter training.
`
		expected := &types.Document{
			Elements: []interface{}{
				&types.AttributeDeclaration{
					Name:  "fds-version",
					Value: string("6.7.4"),
				},
				&types.Paragraph{
					Elements: []interface{}{
						&types.StringElement{
							Content: "PyroSim is a graphical user interface for the ",
						},
						// https://github.com/firemodels/fds/releases/tag/FDS{fds-version}/[Fire Dynamics Simulator (FDS) version {fds-version}]
						&types.InlineLink{
							Attributes: types.Attributes{
								types.AttrInlineLinkText: "Fire Dynamics Simulator (FDS) version 6.7.4",
							},
							Location: &types.Location{
								Scheme: "https://",
								Path:   "github.com/firemodels/fds/releases/tag/FDS6.7.4/",
							},
						},
						&types.StringElement{
							Content: ".\nFDS is closely integrated into PyroSim.\nFDS models can predict smoke, temperature, carbon monoxide, and other substances during fires.\nThe results of these simulations are used to ensure the safety of buildings before construction, evaluate safety options of existing buildings, reconstruct fires for post-accident investigation, and assist in firefighter training.",
						},
					},
				},
			},
		}
		stats := parser.Stats{}
		Expect(ParseDocument(source, parser.Debug(true), parser.Statistics(&stats, "no match"))).To(MatchDocument(expected))
		fmt.Printf("ExprCnt:      %d\n", stats.ExprCnt)
		result, _ := json.MarshalIndent(stats.ChoiceAltCnt, " ", " ")
		fmt.Printf("ChoiceAltCnt: \n%s\n", result)
	})
})
