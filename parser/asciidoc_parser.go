package parser

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/bytesparadise/libasciidoc/types"
)

// *****************************************************************************************
// This file is generated after its sibling `asciidoc-grammar.peg` file. DO NOT MODIFY !
// *****************************************************************************************

var g = &grammar{
	rules: []*rule{
		{
			name: "Document",
			pos:  position{line: 12, col: 1, offset: 351},
			expr: &actionExpr{
				pos: position{line: 12, col: 13, offset: 363},
				run: (*parser).callonDocument1,
				expr: &seqExpr{
					pos: position{line: 12, col: 13, offset: 363},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 12, col: 13, offset: 363},
							label: "blocks",
							expr: &zeroOrMoreExpr{
								pos: position{line: 12, col: 20, offset: 370},
								expr: &ruleRefExpr{
									pos:  position{line: 12, col: 20, offset: 370},
									name: "DocumentBlock",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 12, col: 35, offset: 385},
							name: "EOF",
						},
					},
				},
			},
		},
		{
			name: "DocumentBlock",
			pos:  position{line: 16, col: 1, offset: 444},
			expr: &actionExpr{
				pos: position{line: 16, col: 18, offset: 461},
				run: (*parser).callonDocumentBlock1,
				expr: &seqExpr{
					pos: position{line: 16, col: 18, offset: 461},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 16, col: 18, offset: 461},
							expr: &ruleRefExpr{
								pos:  position{line: 16, col: 19, offset: 462},
								name: "EOF",
							},
						},
						&labeledExpr{
							pos:   position{line: 16, col: 23, offset: 466},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 16, col: 32, offset: 475},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 16, col: 32, offset: 475},
										name: "Section",
									},
									&ruleRefExpr{
										pos:  position{line: 16, col: 42, offset: 485},
										name: "StandaloneBlock",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "StandaloneBlock",
			pos:  position{line: 20, col: 1, offset: 550},
			expr: &choiceExpr{
				pos: position{line: 20, col: 20, offset: 569},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 20, col: 20, offset: 569},
						name: "DocumentAttributeDeclaration",
					},
					&ruleRefExpr{
						pos:  position{line: 20, col: 51, offset: 600},
						name: "DocumentAttributeReset",
					},
					&ruleRefExpr{
						pos:  position{line: 20, col: 76, offset: 625},
						name: "List",
					},
					&ruleRefExpr{
						pos:  position{line: 20, col: 83, offset: 632},
						name: "BlockImage",
					},
					&ruleRefExpr{
						pos:  position{line: 20, col: 96, offset: 645},
						name: "DelimitedBlock",
					},
					&ruleRefExpr{
						pos:  position{line: 20, col: 113, offset: 662},
						name: "Paragraph",
					},
					&ruleRefExpr{
						pos:  position{line: 20, col: 125, offset: 674},
						name: "ElementAttribute",
					},
					&ruleRefExpr{
						pos:  position{line: 20, col: 144, offset: 693},
						name: "BlankLine",
					},
				},
			},
		},
		{
			name: "Section",
			pos:  position{line: 22, col: 1, offset: 748},
			expr: &choiceExpr{
				pos: position{line: 22, col: 12, offset: 759},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 22, col: 12, offset: 759},
						name: "Section1",
					},
					&ruleRefExpr{
						pos:  position{line: 22, col: 23, offset: 770},
						name: "Section2",
					},
					&ruleRefExpr{
						pos:  position{line: 22, col: 34, offset: 781},
						name: "Section3",
					},
					&ruleRefExpr{
						pos:  position{line: 22, col: 45, offset: 792},
						name: "Section4",
					},
					&ruleRefExpr{
						pos:  position{line: 22, col: 56, offset: 803},
						name: "Section5",
					},
					&ruleRefExpr{
						pos:  position{line: 22, col: 67, offset: 814},
						name: "Section6",
					},
				},
			},
		},
		{
			name: "Section1",
			pos:  position{line: 24, col: 1, offset: 824},
			expr: &actionExpr{
				pos: position{line: 24, col: 13, offset: 836},
				run: (*parser).callonSection11,
				expr: &seqExpr{
					pos: position{line: 24, col: 13, offset: 836},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 24, col: 13, offset: 836},
							label: "heading",
							expr: &ruleRefExpr{
								pos:  position{line: 24, col: 22, offset: 845},
								name: "Heading1",
							},
						},
						&labeledExpr{
							pos:   position{line: 24, col: 32, offset: 855},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 24, col: 42, offset: 865},
								expr: &ruleRefExpr{
									pos:  position{line: 24, col: 42, offset: 865},
									name: "Section1Block",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section1Block",
			pos:  position{line: 28, col: 1, offset: 966},
			expr: &actionExpr{
				pos: position{line: 28, col: 18, offset: 983},
				run: (*parser).callonSection1Block1,
				expr: &labeledExpr{
					pos:   position{line: 28, col: 18, offset: 983},
					label: "content",
					expr: &choiceExpr{
						pos: position{line: 28, col: 27, offset: 992},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 28, col: 27, offset: 992},
								name: "Section2",
							},
							&ruleRefExpr{
								pos:  position{line: 28, col: 38, offset: 1003},
								name: "Section3",
							},
							&ruleRefExpr{
								pos:  position{line: 28, col: 49, offset: 1014},
								name: "Section4",
							},
							&ruleRefExpr{
								pos:  position{line: 28, col: 60, offset: 1025},
								name: "Section5",
							},
							&ruleRefExpr{
								pos:  position{line: 28, col: 71, offset: 1036},
								name: "Section6",
							},
							&ruleRefExpr{
								pos:  position{line: 28, col: 82, offset: 1047},
								name: "StandaloneBlock",
							},
						},
					},
				},
			},
		},
		{
			name: "Section2",
			pos:  position{line: 32, col: 1, offset: 1112},
			expr: &actionExpr{
				pos: position{line: 32, col: 13, offset: 1124},
				run: (*parser).callonSection21,
				expr: &seqExpr{
					pos: position{line: 32, col: 13, offset: 1124},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 32, col: 13, offset: 1124},
							label: "heading",
							expr: &ruleRefExpr{
								pos:  position{line: 32, col: 22, offset: 1133},
								name: "Heading2",
							},
						},
						&labeledExpr{
							pos:   position{line: 32, col: 32, offset: 1143},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 32, col: 42, offset: 1153},
								expr: &ruleRefExpr{
									pos:  position{line: 32, col: 42, offset: 1153},
									name: "Section2Block",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section2Block",
			pos:  position{line: 36, col: 1, offset: 1254},
			expr: &actionExpr{
				pos: position{line: 36, col: 18, offset: 1271},
				run: (*parser).callonSection2Block1,
				expr: &seqExpr{
					pos: position{line: 36, col: 18, offset: 1271},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 36, col: 18, offset: 1271},
							expr: &ruleRefExpr{
								pos:  position{line: 36, col: 19, offset: 1272},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 36, col: 28, offset: 1281},
							expr: &ruleRefExpr{
								pos:  position{line: 36, col: 29, offset: 1282},
								name: "Section2",
							},
						},
						&labeledExpr{
							pos:   position{line: 36, col: 38, offset: 1291},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 36, col: 47, offset: 1300},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 36, col: 47, offset: 1300},
										name: "Section3",
									},
									&ruleRefExpr{
										pos:  position{line: 36, col: 58, offset: 1311},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 36, col: 69, offset: 1322},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 36, col: 80, offset: 1333},
										name: "Section6",
									},
									&ruleRefExpr{
										pos:  position{line: 36, col: 91, offset: 1344},
										name: "StandaloneBlock",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section3",
			pos:  position{line: 40, col: 1, offset: 1409},
			expr: &actionExpr{
				pos: position{line: 40, col: 13, offset: 1421},
				run: (*parser).callonSection31,
				expr: &seqExpr{
					pos: position{line: 40, col: 13, offset: 1421},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 40, col: 13, offset: 1421},
							label: "heading",
							expr: &ruleRefExpr{
								pos:  position{line: 40, col: 22, offset: 1430},
								name: "Heading3",
							},
						},
						&labeledExpr{
							pos:   position{line: 40, col: 32, offset: 1440},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 40, col: 42, offset: 1450},
								expr: &ruleRefExpr{
									pos:  position{line: 40, col: 42, offset: 1450},
									name: "Section3Block",
								},
							},
						},
						&andExpr{
							pos: position{line: 40, col: 58, offset: 1466},
							expr: &zeroOrMoreExpr{
								pos: position{line: 40, col: 59, offset: 1467},
								expr: &ruleRefExpr{
									pos:  position{line: 40, col: 60, offset: 1468},
									name: "Section3",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section3Block",
			pos:  position{line: 44, col: 1, offset: 1564},
			expr: &actionExpr{
				pos: position{line: 44, col: 18, offset: 1581},
				run: (*parser).callonSection3Block1,
				expr: &seqExpr{
					pos: position{line: 44, col: 18, offset: 1581},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 44, col: 18, offset: 1581},
							expr: &ruleRefExpr{
								pos:  position{line: 44, col: 19, offset: 1582},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 44, col: 28, offset: 1591},
							expr: &ruleRefExpr{
								pos:  position{line: 44, col: 29, offset: 1592},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 44, col: 38, offset: 1601},
							expr: &ruleRefExpr{
								pos:  position{line: 44, col: 39, offset: 1602},
								name: "Section3",
							},
						},
						&labeledExpr{
							pos:   position{line: 44, col: 48, offset: 1611},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 44, col: 57, offset: 1620},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 44, col: 57, offset: 1620},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 44, col: 68, offset: 1631},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 44, col: 79, offset: 1642},
										name: "Section6",
									},
									&ruleRefExpr{
										pos:  position{line: 44, col: 90, offset: 1653},
										name: "StandaloneBlock",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section4",
			pos:  position{line: 48, col: 1, offset: 1718},
			expr: &actionExpr{
				pos: position{line: 48, col: 13, offset: 1730},
				run: (*parser).callonSection41,
				expr: &seqExpr{
					pos: position{line: 48, col: 13, offset: 1730},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 48, col: 13, offset: 1730},
							label: "heading",
							expr: &ruleRefExpr{
								pos:  position{line: 48, col: 22, offset: 1739},
								name: "Heading4",
							},
						},
						&labeledExpr{
							pos:   position{line: 48, col: 32, offset: 1749},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 48, col: 42, offset: 1759},
								expr: &ruleRefExpr{
									pos:  position{line: 48, col: 42, offset: 1759},
									name: "Section4Block",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section4Block",
			pos:  position{line: 52, col: 1, offset: 1860},
			expr: &actionExpr{
				pos: position{line: 52, col: 18, offset: 1877},
				run: (*parser).callonSection4Block1,
				expr: &seqExpr{
					pos: position{line: 52, col: 18, offset: 1877},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 52, col: 18, offset: 1877},
							expr: &ruleRefExpr{
								pos:  position{line: 52, col: 19, offset: 1878},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 52, col: 28, offset: 1887},
							expr: &ruleRefExpr{
								pos:  position{line: 52, col: 29, offset: 1888},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 52, col: 38, offset: 1897},
							expr: &ruleRefExpr{
								pos:  position{line: 52, col: 39, offset: 1898},
								name: "Section3",
							},
						},
						&notExpr{
							pos: position{line: 52, col: 48, offset: 1907},
							expr: &ruleRefExpr{
								pos:  position{line: 52, col: 49, offset: 1908},
								name: "Section4",
							},
						},
						&labeledExpr{
							pos:   position{line: 52, col: 58, offset: 1917},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 52, col: 67, offset: 1926},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 52, col: 67, offset: 1926},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 52, col: 78, offset: 1937},
										name: "Section6",
									},
									&ruleRefExpr{
										pos:  position{line: 52, col: 89, offset: 1948},
										name: "StandaloneBlock",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section5",
			pos:  position{line: 56, col: 1, offset: 2013},
			expr: &actionExpr{
				pos: position{line: 56, col: 13, offset: 2025},
				run: (*parser).callonSection51,
				expr: &seqExpr{
					pos: position{line: 56, col: 13, offset: 2025},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 56, col: 13, offset: 2025},
							label: "heading",
							expr: &ruleRefExpr{
								pos:  position{line: 56, col: 22, offset: 2034},
								name: "Heading5",
							},
						},
						&labeledExpr{
							pos:   position{line: 56, col: 32, offset: 2044},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 56, col: 42, offset: 2054},
								expr: &ruleRefExpr{
									pos:  position{line: 56, col: 42, offset: 2054},
									name: "Section5Block",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section5Block",
			pos:  position{line: 60, col: 1, offset: 2155},
			expr: &actionExpr{
				pos: position{line: 60, col: 18, offset: 2172},
				run: (*parser).callonSection5Block1,
				expr: &seqExpr{
					pos: position{line: 60, col: 18, offset: 2172},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 60, col: 18, offset: 2172},
							expr: &ruleRefExpr{
								pos:  position{line: 60, col: 19, offset: 2173},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 60, col: 28, offset: 2182},
							expr: &ruleRefExpr{
								pos:  position{line: 60, col: 29, offset: 2183},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 60, col: 38, offset: 2192},
							expr: &ruleRefExpr{
								pos:  position{line: 60, col: 39, offset: 2193},
								name: "Section3",
							},
						},
						&notExpr{
							pos: position{line: 60, col: 48, offset: 2202},
							expr: &ruleRefExpr{
								pos:  position{line: 60, col: 49, offset: 2203},
								name: "Section4",
							},
						},
						&notExpr{
							pos: position{line: 60, col: 58, offset: 2212},
							expr: &ruleRefExpr{
								pos:  position{line: 60, col: 59, offset: 2213},
								name: "Section5",
							},
						},
						&labeledExpr{
							pos:   position{line: 60, col: 68, offset: 2222},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 60, col: 77, offset: 2231},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 60, col: 77, offset: 2231},
										name: "Section6",
									},
									&ruleRefExpr{
										pos:  position{line: 60, col: 88, offset: 2242},
										name: "StandaloneBlock",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section6",
			pos:  position{line: 64, col: 1, offset: 2307},
			expr: &actionExpr{
				pos: position{line: 64, col: 13, offset: 2319},
				run: (*parser).callonSection61,
				expr: &seqExpr{
					pos: position{line: 64, col: 13, offset: 2319},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 64, col: 13, offset: 2319},
							label: "heading",
							expr: &ruleRefExpr{
								pos:  position{line: 64, col: 22, offset: 2328},
								name: "Heading6",
							},
						},
						&labeledExpr{
							pos:   position{line: 64, col: 32, offset: 2338},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 64, col: 42, offset: 2348},
								expr: &ruleRefExpr{
									pos:  position{line: 64, col: 42, offset: 2348},
									name: "Section6Block",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section6Block",
			pos:  position{line: 68, col: 1, offset: 2449},
			expr: &actionExpr{
				pos: position{line: 68, col: 18, offset: 2466},
				run: (*parser).callonSection6Block1,
				expr: &seqExpr{
					pos: position{line: 68, col: 18, offset: 2466},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 68, col: 18, offset: 2466},
							expr: &ruleRefExpr{
								pos:  position{line: 68, col: 19, offset: 2467},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 68, col: 28, offset: 2476},
							expr: &ruleRefExpr{
								pos:  position{line: 68, col: 29, offset: 2477},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 68, col: 38, offset: 2486},
							expr: &ruleRefExpr{
								pos:  position{line: 68, col: 39, offset: 2487},
								name: "Section3",
							},
						},
						&notExpr{
							pos: position{line: 68, col: 48, offset: 2496},
							expr: &ruleRefExpr{
								pos:  position{line: 68, col: 49, offset: 2497},
								name: "Section4",
							},
						},
						&notExpr{
							pos: position{line: 68, col: 58, offset: 2506},
							expr: &ruleRefExpr{
								pos:  position{line: 68, col: 59, offset: 2507},
								name: "Section5",
							},
						},
						&notExpr{
							pos: position{line: 68, col: 68, offset: 2516},
							expr: &ruleRefExpr{
								pos:  position{line: 68, col: 69, offset: 2517},
								name: "Section6",
							},
						},
						&labeledExpr{
							pos:   position{line: 68, col: 78, offset: 2526},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 68, col: 87, offset: 2535},
								name: "StandaloneBlock",
							},
						},
					},
				},
			},
		},
		{
			name: "Heading",
			pos:  position{line: 76, col: 1, offset: 2705},
			expr: &choiceExpr{
				pos: position{line: 76, col: 12, offset: 2716},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 76, col: 12, offset: 2716},
						name: "Heading1",
					},
					&ruleRefExpr{
						pos:  position{line: 76, col: 23, offset: 2727},
						name: "Heading2",
					},
					&ruleRefExpr{
						pos:  position{line: 76, col: 34, offset: 2738},
						name: "Heading3",
					},
					&ruleRefExpr{
						pos:  position{line: 76, col: 45, offset: 2749},
						name: "Heading4",
					},
					&ruleRefExpr{
						pos:  position{line: 76, col: 56, offset: 2760},
						name: "Heading5",
					},
					&ruleRefExpr{
						pos:  position{line: 76, col: 67, offset: 2771},
						name: "Heading6",
					},
				},
			},
		},
		{
			name: "Heading1",
			pos:  position{line: 78, col: 1, offset: 2781},
			expr: &actionExpr{
				pos: position{line: 78, col: 13, offset: 2793},
				run: (*parser).callonHeading11,
				expr: &seqExpr{
					pos: position{line: 78, col: 13, offset: 2793},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 78, col: 13, offset: 2793},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 78, col: 24, offset: 2804},
								expr: &ruleRefExpr{
									pos:  position{line: 78, col: 25, offset: 2805},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 78, col: 44, offset: 2824},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 78, col: 51, offset: 2831},
								val:        "=",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 78, col: 56, offset: 2836},
							expr: &ruleRefExpr{
								pos:  position{line: 78, col: 56, offset: 2836},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 78, col: 60, offset: 2840},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 78, col: 68, offset: 2848},
								name: "InlineContent",
							},
						},
						&choiceExpr{
							pos: position{line: 78, col: 83, offset: 2863},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 78, col: 83, offset: 2863},
									expr: &ruleRefExpr{
										pos:  position{line: 78, col: 83, offset: 2863},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 78, col: 96, offset: 2876},
									name: "EOF",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Heading2",
			pos:  position{line: 82, col: 1, offset: 3083},
			expr: &actionExpr{
				pos: position{line: 82, col: 13, offset: 3095},
				run: (*parser).callonHeading21,
				expr: &seqExpr{
					pos: position{line: 82, col: 13, offset: 3095},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 82, col: 13, offset: 3095},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 82, col: 24, offset: 3106},
								expr: &ruleRefExpr{
									pos:  position{line: 82, col: 25, offset: 3107},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 82, col: 44, offset: 3126},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 82, col: 51, offset: 3133},
								val:        "==",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 82, col: 57, offset: 3139},
							expr: &ruleRefExpr{
								pos:  position{line: 82, col: 57, offset: 3139},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 82, col: 61, offset: 3143},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 82, col: 69, offset: 3151},
								name: "InlineContent",
							},
						},
						&choiceExpr{
							pos: position{line: 82, col: 84, offset: 3166},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 82, col: 84, offset: 3166},
									expr: &ruleRefExpr{
										pos:  position{line: 82, col: 84, offset: 3166},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 82, col: 97, offset: 3179},
									name: "EOF",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Heading3",
			pos:  position{line: 86, col: 1, offset: 3281},
			expr: &actionExpr{
				pos: position{line: 86, col: 13, offset: 3293},
				run: (*parser).callonHeading31,
				expr: &seqExpr{
					pos: position{line: 86, col: 13, offset: 3293},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 86, col: 13, offset: 3293},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 86, col: 24, offset: 3304},
								expr: &ruleRefExpr{
									pos:  position{line: 86, col: 25, offset: 3305},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 86, col: 44, offset: 3324},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 86, col: 51, offset: 3331},
								val:        "===",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 86, col: 58, offset: 3338},
							expr: &ruleRefExpr{
								pos:  position{line: 86, col: 58, offset: 3338},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 86, col: 62, offset: 3342},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 86, col: 70, offset: 3350},
								name: "InlineContent",
							},
						},
						&choiceExpr{
							pos: position{line: 86, col: 85, offset: 3365},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 86, col: 85, offset: 3365},
									expr: &ruleRefExpr{
										pos:  position{line: 86, col: 85, offset: 3365},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 86, col: 98, offset: 3378},
									name: "EOF",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Heading4",
			pos:  position{line: 90, col: 1, offset: 3480},
			expr: &actionExpr{
				pos: position{line: 90, col: 13, offset: 3492},
				run: (*parser).callonHeading41,
				expr: &seqExpr{
					pos: position{line: 90, col: 13, offset: 3492},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 90, col: 13, offset: 3492},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 90, col: 24, offset: 3503},
								expr: &ruleRefExpr{
									pos:  position{line: 90, col: 25, offset: 3504},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 90, col: 44, offset: 3523},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 90, col: 51, offset: 3530},
								val:        "====",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 90, col: 59, offset: 3538},
							expr: &ruleRefExpr{
								pos:  position{line: 90, col: 59, offset: 3538},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 90, col: 63, offset: 3542},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 90, col: 71, offset: 3550},
								name: "InlineContent",
							},
						},
						&choiceExpr{
							pos: position{line: 90, col: 86, offset: 3565},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 90, col: 86, offset: 3565},
									expr: &ruleRefExpr{
										pos:  position{line: 90, col: 86, offset: 3565},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 90, col: 99, offset: 3578},
									name: "EOF",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Heading5",
			pos:  position{line: 94, col: 1, offset: 3680},
			expr: &actionExpr{
				pos: position{line: 94, col: 13, offset: 3692},
				run: (*parser).callonHeading51,
				expr: &seqExpr{
					pos: position{line: 94, col: 13, offset: 3692},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 94, col: 13, offset: 3692},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 94, col: 24, offset: 3703},
								expr: &ruleRefExpr{
									pos:  position{line: 94, col: 25, offset: 3704},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 94, col: 44, offset: 3723},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 94, col: 51, offset: 3730},
								val:        "=====",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 94, col: 60, offset: 3739},
							expr: &ruleRefExpr{
								pos:  position{line: 94, col: 60, offset: 3739},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 94, col: 64, offset: 3743},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 94, col: 72, offset: 3751},
								name: "InlineContent",
							},
						},
						&choiceExpr{
							pos: position{line: 94, col: 87, offset: 3766},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 94, col: 87, offset: 3766},
									expr: &ruleRefExpr{
										pos:  position{line: 94, col: 87, offset: 3766},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 94, col: 100, offset: 3779},
									name: "EOF",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Heading6",
			pos:  position{line: 98, col: 1, offset: 3881},
			expr: &actionExpr{
				pos: position{line: 98, col: 13, offset: 3893},
				run: (*parser).callonHeading61,
				expr: &seqExpr{
					pos: position{line: 98, col: 13, offset: 3893},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 98, col: 13, offset: 3893},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 98, col: 24, offset: 3904},
								expr: &ruleRefExpr{
									pos:  position{line: 98, col: 25, offset: 3905},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 98, col: 44, offset: 3924},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 98, col: 51, offset: 3931},
								val:        "======",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 98, col: 61, offset: 3941},
							expr: &ruleRefExpr{
								pos:  position{line: 98, col: 61, offset: 3941},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 98, col: 65, offset: 3945},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 98, col: 73, offset: 3953},
								name: "InlineContent",
							},
						},
						&choiceExpr{
							pos: position{line: 98, col: 88, offset: 3968},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 98, col: 88, offset: 3968},
									expr: &ruleRefExpr{
										pos:  position{line: 98, col: 88, offset: 3968},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 101, offset: 3981},
									name: "EOF",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclaration",
			pos:  position{line: 106, col: 1, offset: 4199},
			expr: &choiceExpr{
				pos: position{line: 106, col: 33, offset: 4231},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 106, col: 33, offset: 4231},
						name: "DocumentAttributeDeclarationWithNameOnly",
					},
					&ruleRefExpr{
						pos:  position{line: 106, col: 76, offset: 4274},
						name: "DocumentAttributeDeclarationWithNameAndValue",
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclarationWithNameOnly",
			pos:  position{line: 108, col: 1, offset: 4321},
			expr: &actionExpr{
				pos: position{line: 108, col: 45, offset: 4365},
				run: (*parser).callonDocumentAttributeDeclarationWithNameOnly1,
				expr: &seqExpr{
					pos: position{line: 108, col: 45, offset: 4365},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 108, col: 45, offset: 4365},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 108, col: 49, offset: 4369},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 108, col: 55, offset: 4375},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 108, col: 70, offset: 4390},
							val:        ":",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 108, col: 74, offset: 4394},
							expr: &ruleRefExpr{
								pos:  position{line: 108, col: 74, offset: 4394},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 108, col: 78, offset: 4398},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclarationWithNameAndValue",
			pos:  position{line: 112, col: 1, offset: 4483},
			expr: &actionExpr{
				pos: position{line: 112, col: 49, offset: 4531},
				run: (*parser).callonDocumentAttributeDeclarationWithNameAndValue1,
				expr: &seqExpr{
					pos: position{line: 112, col: 49, offset: 4531},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 112, col: 49, offset: 4531},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 112, col: 53, offset: 4535},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 112, col: 59, offset: 4541},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 112, col: 74, offset: 4556},
							val:        ":",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 112, col: 78, offset: 4560},
							expr: &ruleRefExpr{
								pos:  position{line: 112, col: 78, offset: 4560},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 112, col: 82, offset: 4564},
							label: "value",
							expr: &zeroOrMoreExpr{
								pos: position{line: 112, col: 88, offset: 4570},
								expr: &seqExpr{
									pos: position{line: 112, col: 89, offset: 4571},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 112, col: 89, offset: 4571},
											expr: &ruleRefExpr{
												pos:  position{line: 112, col: 90, offset: 4572},
												name: "NEWLINE",
											},
										},
										&anyMatcher{
											line: 112, col: 98, offset: 4580,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 112, col: 102, offset: 4584},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeReset",
			pos:  position{line: 116, col: 1, offset: 4687},
			expr: &choiceExpr{
				pos: position{line: 116, col: 27, offset: 4713},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 116, col: 27, offset: 4713},
						name: "DocumentAttributeResetWithHeadingBangSymbol",
					},
					&ruleRefExpr{
						pos:  position{line: 116, col: 73, offset: 4759},
						name: "DocumentAttributeResetWithTrailingBangSymbol",
					},
				},
			},
		},
		{
			name: "DocumentAttributeResetWithHeadingBangSymbol",
			pos:  position{line: 118, col: 1, offset: 4805},
			expr: &actionExpr{
				pos: position{line: 118, col: 48, offset: 4852},
				run: (*parser).callonDocumentAttributeResetWithHeadingBangSymbol1,
				expr: &seqExpr{
					pos: position{line: 118, col: 48, offset: 4852},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 118, col: 48, offset: 4852},
							val:        ":!",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 118, col: 53, offset: 4857},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 118, col: 59, offset: 4863},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 118, col: 74, offset: 4878},
							val:        ":",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 118, col: 78, offset: 4882},
							expr: &ruleRefExpr{
								pos:  position{line: 118, col: 78, offset: 4882},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 118, col: 82, offset: 4886},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeResetWithTrailingBangSymbol",
			pos:  position{line: 122, col: 1, offset: 4960},
			expr: &actionExpr{
				pos: position{line: 122, col: 49, offset: 5008},
				run: (*parser).callonDocumentAttributeResetWithTrailingBangSymbol1,
				expr: &seqExpr{
					pos: position{line: 122, col: 49, offset: 5008},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 122, col: 49, offset: 5008},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 122, col: 53, offset: 5012},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 122, col: 59, offset: 5018},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 122, col: 74, offset: 5033},
							val:        "!:",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 122, col: 79, offset: 5038},
							expr: &ruleRefExpr{
								pos:  position{line: 122, col: 79, offset: 5038},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 122, col: 83, offset: 5042},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeSubstitution",
			pos:  position{line: 127, col: 1, offset: 5117},
			expr: &actionExpr{
				pos: position{line: 127, col: 34, offset: 5150},
				run: (*parser).callonDocumentAttributeSubstitution1,
				expr: &seqExpr{
					pos: position{line: 127, col: 34, offset: 5150},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 127, col: 34, offset: 5150},
							val:        "{",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 127, col: 38, offset: 5154},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 127, col: 44, offset: 5160},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 127, col: 59, offset: 5175},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "AttributeName",
			pos:  position{line: 134, col: 1, offset: 5429},
			expr: &seqExpr{
				pos: position{line: 134, col: 18, offset: 5446},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 134, col: 19, offset: 5447},
						alternatives: []interface{}{
							&charClassMatcher{
								pos:        position{line: 134, col: 19, offset: 5447},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
							&charClassMatcher{
								pos:        position{line: 134, col: 27, offset: 5455},
								val:        "[a-z]",
								ranges:     []rune{'a', 'z'},
								ignoreCase: false,
								inverted:   false,
							},
							&charClassMatcher{
								pos:        position{line: 134, col: 35, offset: 5463},
								val:        "[0-9]",
								ranges:     []rune{'0', '9'},
								ignoreCase: false,
								inverted:   false,
							},
							&litMatcher{
								pos:        position{line: 134, col: 43, offset: 5471},
								val:        "_",
								ignoreCase: false,
							},
						},
					},
					&zeroOrMoreExpr{
						pos: position{line: 134, col: 48, offset: 5476},
						expr: &choiceExpr{
							pos: position{line: 134, col: 49, offset: 5477},
							alternatives: []interface{}{
								&charClassMatcher{
									pos:        position{line: 134, col: 49, offset: 5477},
									val:        "[A-Z]",
									ranges:     []rune{'A', 'Z'},
									ignoreCase: false,
									inverted:   false,
								},
								&charClassMatcher{
									pos:        position{line: 134, col: 57, offset: 5485},
									val:        "[a-z]",
									ranges:     []rune{'a', 'z'},
									ignoreCase: false,
									inverted:   false,
								},
								&charClassMatcher{
									pos:        position{line: 134, col: 65, offset: 5493},
									val:        "[0-9]",
									ranges:     []rune{'0', '9'},
									ignoreCase: false,
									inverted:   false,
								},
								&litMatcher{
									pos:        position{line: 134, col: 73, offset: 5501},
									val:        "-",
									ignoreCase: false,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "List",
			pos:  position{line: 139, col: 1, offset: 5614},
			expr: &actionExpr{
				pos: position{line: 139, col: 9, offset: 5622},
				run: (*parser).callonList1,
				expr: &seqExpr{
					pos: position{line: 139, col: 9, offset: 5622},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 139, col: 9, offset: 5622},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 139, col: 20, offset: 5633},
								expr: &ruleRefExpr{
									pos:  position{line: 139, col: 21, offset: 5634},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 141, col: 5, offset: 5726},
							label: "elements",
							expr: &oneOrMoreExpr{
								pos: position{line: 141, col: 14, offset: 5735},
								expr: &seqExpr{
									pos: position{line: 141, col: 15, offset: 5736},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 141, col: 15, offset: 5736},
											name: "ListItem",
										},
										&zeroOrOneExpr{
											pos: position{line: 141, col: 24, offset: 5745},
											expr: &ruleRefExpr{
												pos:  position{line: 141, col: 24, offset: 5745},
												name: "BlankLine",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "ListItem",
			pos:  position{line: 145, col: 1, offset: 5842},
			expr: &actionExpr{
				pos: position{line: 145, col: 13, offset: 5854},
				run: (*parser).callonListItem1,
				expr: &seqExpr{
					pos: position{line: 145, col: 13, offset: 5854},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 145, col: 13, offset: 5854},
							expr: &ruleRefExpr{
								pos:  position{line: 145, col: 13, offset: 5854},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 145, col: 17, offset: 5858},
							label: "level",
							expr: &choiceExpr{
								pos: position{line: 145, col: 24, offset: 5865},
								alternatives: []interface{}{
									&oneOrMoreExpr{
										pos: position{line: 145, col: 24, offset: 5865},
										expr: &litMatcher{
											pos:        position{line: 145, col: 24, offset: 5865},
											val:        "*",
											ignoreCase: false,
										},
									},
									&litMatcher{
										pos:        position{line: 145, col: 31, offset: 5872},
										val:        "-",
										ignoreCase: false,
									},
								},
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 145, col: 36, offset: 5877},
							expr: &ruleRefExpr{
								pos:  position{line: 145, col: 36, offset: 5877},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 145, col: 40, offset: 5881},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 145, col: 49, offset: 5890},
								name: "ListItemContent",
							},
						},
					},
				},
			},
		},
		{
			name: "ListItemContent",
			pos:  position{line: 149, col: 1, offset: 5987},
			expr: &actionExpr{
				pos: position{line: 149, col: 20, offset: 6006},
				run: (*parser).callonListItemContent1,
				expr: &labeledExpr{
					pos:   position{line: 149, col: 20, offset: 6006},
					label: "lines",
					expr: &oneOrMoreExpr{
						pos: position{line: 149, col: 26, offset: 6012},
						expr: &seqExpr{
							pos: position{line: 149, col: 27, offset: 6013},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 149, col: 27, offset: 6013},
									expr: &seqExpr{
										pos: position{line: 149, col: 29, offset: 6015},
										exprs: []interface{}{
											&zeroOrMoreExpr{
												pos: position{line: 149, col: 29, offset: 6015},
												expr: &ruleRefExpr{
													pos:  position{line: 149, col: 29, offset: 6015},
													name: "WS",
												},
											},
											&choiceExpr{
												pos: position{line: 149, col: 34, offset: 6020},
												alternatives: []interface{}{
													&oneOrMoreExpr{
														pos: position{line: 149, col: 34, offset: 6020},
														expr: &litMatcher{
															pos:        position{line: 149, col: 34, offset: 6020},
															val:        "*",
															ignoreCase: false,
														},
													},
													&litMatcher{
														pos:        position{line: 149, col: 41, offset: 6027},
														val:        "-",
														ignoreCase: false,
													},
												},
											},
											&oneOrMoreExpr{
												pos: position{line: 149, col: 46, offset: 6032},
												expr: &ruleRefExpr{
													pos:  position{line: 149, col: 46, offset: 6032},
													name: "WS",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 149, col: 51, offset: 6037},
									name: "InlineContent",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Paragraph",
			pos:  position{line: 156, col: 1, offset: 6307},
			expr: &actionExpr{
				pos: position{line: 156, col: 14, offset: 6320},
				run: (*parser).callonParagraph1,
				expr: &seqExpr{
					pos: position{line: 156, col: 14, offset: 6320},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 156, col: 14, offset: 6320},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 156, col: 25, offset: 6331},
								expr: &ruleRefExpr{
									pos:  position{line: 156, col: 26, offset: 6332},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 156, col: 45, offset: 6351},
							label: "lines",
							expr: &oneOrMoreExpr{
								pos: position{line: 156, col: 51, offset: 6357},
								expr: &ruleRefExpr{
									pos:  position{line: 156, col: 52, offset: 6358},
									name: "InlineContent",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "InlineContent",
			pos:  position{line: 162, col: 1, offset: 6666},
			expr: &actionExpr{
				pos: position{line: 162, col: 18, offset: 6683},
				run: (*parser).callonInlineContent1,
				expr: &seqExpr{
					pos: position{line: 162, col: 18, offset: 6683},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 162, col: 18, offset: 6683},
							label: "elements",
							expr: &oneOrMoreExpr{
								pos: position{line: 162, col: 27, offset: 6692},
								expr: &seqExpr{
									pos: position{line: 162, col: 28, offset: 6693},
									exprs: []interface{}{
										&zeroOrMoreExpr{
											pos: position{line: 162, col: 28, offset: 6693},
											expr: &ruleRefExpr{
												pos:  position{line: 162, col: 28, offset: 6693},
												name: "WS",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 162, col: 32, offset: 6697},
											name: "InlineElement",
										},
										&zeroOrMoreExpr{
											pos: position{line: 162, col: 46, offset: 6711},
											expr: &ruleRefExpr{
												pos:  position{line: 162, col: 46, offset: 6711},
												name: "WS",
											},
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 162, col: 52, offset: 6717},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "InlineElement",
			pos:  position{line: 166, col: 1, offset: 6795},
			expr: &choiceExpr{
				pos: position{line: 166, col: 18, offset: 6812},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 166, col: 18, offset: 6812},
						name: "InlineImage",
					},
					&ruleRefExpr{
						pos:  position{line: 166, col: 32, offset: 6826},
						name: "QuotedText",
					},
					&ruleRefExpr{
						pos:  position{line: 166, col: 45, offset: 6839},
						name: "ExternalLink",
					},
					&ruleRefExpr{
						pos:  position{line: 166, col: 60, offset: 6854},
						name: "DocumentAttributeSubstitution",
					},
					&ruleRefExpr{
						pos:  position{line: 166, col: 92, offset: 6886},
						name: "Word",
					},
				},
			},
		},
		{
			name: "QuotedText",
			pos:  position{line: 171, col: 1, offset: 7029},
			expr: &choiceExpr{
				pos: position{line: 171, col: 15, offset: 7043},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 171, col: 15, offset: 7043},
						name: "BoldText",
					},
					&ruleRefExpr{
						pos:  position{line: 171, col: 26, offset: 7054},
						name: "ItalicText",
					},
					&ruleRefExpr{
						pos:  position{line: 171, col: 39, offset: 7067},
						name: "MonospaceText",
					},
				},
			},
		},
		{
			name: "BoldText",
			pos:  position{line: 173, col: 1, offset: 7082},
			expr: &actionExpr{
				pos: position{line: 173, col: 13, offset: 7094},
				run: (*parser).callonBoldText1,
				expr: &seqExpr{
					pos: position{line: 173, col: 13, offset: 7094},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 173, col: 13, offset: 7094},
							val:        "*",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 173, col: 17, offset: 7098},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 173, col: 26, offset: 7107},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 173, col: 45, offset: 7126},
							val:        "*",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ItalicText",
			pos:  position{line: 177, col: 1, offset: 7203},
			expr: &actionExpr{
				pos: position{line: 177, col: 15, offset: 7217},
				run: (*parser).callonItalicText1,
				expr: &seqExpr{
					pos: position{line: 177, col: 15, offset: 7217},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 177, col: 15, offset: 7217},
							val:        "_",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 177, col: 19, offset: 7221},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 177, col: 28, offset: 7230},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 177, col: 47, offset: 7249},
							val:        "_",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "MonospaceText",
			pos:  position{line: 181, col: 1, offset: 7328},
			expr: &actionExpr{
				pos: position{line: 181, col: 18, offset: 7345},
				run: (*parser).callonMonospaceText1,
				expr: &seqExpr{
					pos: position{line: 181, col: 18, offset: 7345},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 181, col: 18, offset: 7345},
							val:        "`",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 181, col: 22, offset: 7349},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 181, col: 31, offset: 7358},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 181, col: 50, offset: 7377},
							val:        "`",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "QuotedTextContent",
			pos:  position{line: 185, col: 1, offset: 7459},
			expr: &seqExpr{
				pos: position{line: 185, col: 22, offset: 7480},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 185, col: 22, offset: 7480},
						name: "QuotedTextContentElement",
					},
					&zeroOrMoreExpr{
						pos: position{line: 185, col: 47, offset: 7505},
						expr: &seqExpr{
							pos: position{line: 185, col: 48, offset: 7506},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 185, col: 48, offset: 7506},
									expr: &ruleRefExpr{
										pos:  position{line: 185, col: 48, offset: 7506},
										name: "WS",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 185, col: 52, offset: 7510},
									name: "QuotedTextContentElement",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "QuotedTextContentElement",
			pos:  position{line: 187, col: 1, offset: 7538},
			expr: &choiceExpr{
				pos: position{line: 187, col: 29, offset: 7566},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 187, col: 29, offset: 7566},
						name: "QuotedText",
					},
					&ruleRefExpr{
						pos:  position{line: 187, col: 42, offset: 7579},
						name: "QuotedTextContentWord",
					},
					&ruleRefExpr{
						pos:  position{line: 187, col: 66, offset: 7603},
						name: "InvalidQuotedTextContentWord",
					},
				},
			},
		},
		{
			name: "QuotedTextContentWord",
			pos:  position{line: 189, col: 1, offset: 7633},
			expr: &oneOrMoreExpr{
				pos: position{line: 189, col: 26, offset: 7658},
				expr: &seqExpr{
					pos: position{line: 189, col: 27, offset: 7659},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 189, col: 27, offset: 7659},
							expr: &ruleRefExpr{
								pos:  position{line: 189, col: 28, offset: 7660},
								name: "NEWLINE",
							},
						},
						&notExpr{
							pos: position{line: 189, col: 36, offset: 7668},
							expr: &ruleRefExpr{
								pos:  position{line: 189, col: 37, offset: 7669},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 189, col: 40, offset: 7672},
							expr: &litMatcher{
								pos:        position{line: 189, col: 41, offset: 7673},
								val:        "*",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 189, col: 45, offset: 7677},
							expr: &litMatcher{
								pos:        position{line: 189, col: 46, offset: 7678},
								val:        "_",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 189, col: 50, offset: 7682},
							expr: &litMatcher{
								pos:        position{line: 189, col: 51, offset: 7683},
								val:        "`",
								ignoreCase: false,
							},
						},
						&anyMatcher{
							line: 189, col: 55, offset: 7687,
						},
					},
				},
			},
		},
		{
			name: "InvalidQuotedTextContentWord",
			pos:  position{line: 190, col: 1, offset: 7729},
			expr: &oneOrMoreExpr{
				pos: position{line: 190, col: 33, offset: 7761},
				expr: &seqExpr{
					pos: position{line: 190, col: 34, offset: 7762},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 190, col: 34, offset: 7762},
							expr: &ruleRefExpr{
								pos:  position{line: 190, col: 35, offset: 7763},
								name: "NEWLINE",
							},
						},
						&notExpr{
							pos: position{line: 190, col: 43, offset: 7771},
							expr: &ruleRefExpr{
								pos:  position{line: 190, col: 44, offset: 7772},
								name: "WS",
							},
						},
						&anyMatcher{
							line: 190, col: 48, offset: 7776,
						},
					},
				},
			},
		},
		{
			name: "ExternalLink",
			pos:  position{line: 195, col: 1, offset: 7993},
			expr: &actionExpr{
				pos: position{line: 195, col: 17, offset: 8009},
				run: (*parser).callonExternalLink1,
				expr: &seqExpr{
					pos: position{line: 195, col: 17, offset: 8009},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 195, col: 17, offset: 8009},
							label: "url",
							expr: &seqExpr{
								pos: position{line: 195, col: 22, offset: 8014},
								exprs: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 195, col: 22, offset: 8014},
										name: "URL_SCHEME",
									},
									&ruleRefExpr{
										pos:  position{line: 195, col: 33, offset: 8025},
										name: "URL",
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 195, col: 38, offset: 8030},
							label: "text",
							expr: &zeroOrOneExpr{
								pos: position{line: 195, col: 43, offset: 8035},
								expr: &seqExpr{
									pos: position{line: 195, col: 44, offset: 8036},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 195, col: 44, offset: 8036},
											val:        "[",
											ignoreCase: false,
										},
										&zeroOrMoreExpr{
											pos: position{line: 195, col: 48, offset: 8040},
											expr: &ruleRefExpr{
												pos:  position{line: 195, col: 49, offset: 8041},
												name: "URL_TEXT",
											},
										},
										&litMatcher{
											pos:        position{line: 195, col: 60, offset: 8052},
											val:        "]",
											ignoreCase: false,
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "BlockImage",
			pos:  position{line: 205, col: 1, offset: 8331},
			expr: &actionExpr{
				pos: position{line: 205, col: 15, offset: 8345},
				run: (*parser).callonBlockImage1,
				expr: &seqExpr{
					pos: position{line: 205, col: 15, offset: 8345},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 205, col: 15, offset: 8345},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 205, col: 26, offset: 8356},
								expr: &ruleRefExpr{
									pos:  position{line: 205, col: 27, offset: 8357},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 205, col: 46, offset: 8376},
							label: "image",
							expr: &ruleRefExpr{
								pos:  position{line: 205, col: 52, offset: 8382},
								name: "BlockImageMacro",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 205, col: 69, offset: 8399},
							expr: &ruleRefExpr{
								pos:  position{line: 205, col: 69, offset: 8399},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 205, col: 73, offset: 8403},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "BlockImageMacro",
			pos:  position{line: 210, col: 1, offset: 8572},
			expr: &actionExpr{
				pos: position{line: 210, col: 20, offset: 8591},
				run: (*parser).callonBlockImageMacro1,
				expr: &seqExpr{
					pos: position{line: 210, col: 20, offset: 8591},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 210, col: 20, offset: 8591},
							val:        "image::",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 210, col: 30, offset: 8601},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 210, col: 36, offset: 8607},
								name: "URL",
							},
						},
						&litMatcher{
							pos:        position{line: 210, col: 41, offset: 8612},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 210, col: 45, offset: 8616},
							label: "attributes",
							expr: &zeroOrOneExpr{
								pos: position{line: 210, col: 57, offset: 8628},
								expr: &ruleRefExpr{
									pos:  position{line: 210, col: 57, offset: 8628},
									name: "URL_TEXT",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 210, col: 68, offset: 8639},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "InlineImage",
			pos:  position{line: 214, col: 1, offset: 8714},
			expr: &actionExpr{
				pos: position{line: 214, col: 16, offset: 8729},
				run: (*parser).callonInlineImage1,
				expr: &labeledExpr{
					pos:   position{line: 214, col: 16, offset: 8729},
					label: "image",
					expr: &ruleRefExpr{
						pos:  position{line: 214, col: 22, offset: 8735},
						name: "InlineImageMacro",
					},
				},
			},
		},
		{
			name: "InlineImageMacro",
			pos:  position{line: 219, col: 1, offset: 8890},
			expr: &actionExpr{
				pos: position{line: 219, col: 21, offset: 8910},
				run: (*parser).callonInlineImageMacro1,
				expr: &seqExpr{
					pos: position{line: 219, col: 21, offset: 8910},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 219, col: 21, offset: 8910},
							val:        "image:",
							ignoreCase: false,
						},
						&notExpr{
							pos: position{line: 219, col: 30, offset: 8919},
							expr: &litMatcher{
								pos:        position{line: 219, col: 31, offset: 8920},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 219, col: 35, offset: 8924},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 219, col: 41, offset: 8930},
								name: "URL",
							},
						},
						&litMatcher{
							pos:        position{line: 219, col: 46, offset: 8935},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 219, col: 50, offset: 8939},
							label: "attributes",
							expr: &zeroOrOneExpr{
								pos: position{line: 219, col: 62, offset: 8951},
								expr: &ruleRefExpr{
									pos:  position{line: 219, col: 62, offset: 8951},
									name: "URL_TEXT",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 219, col: 73, offset: 8962},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "DelimitedBlock",
			pos:  position{line: 227, col: 1, offset: 9150},
			expr: &ruleRefExpr{
				pos:  position{line: 227, col: 19, offset: 9168},
				name: "SourceBlock",
			},
		},
		{
			name: "SourceBlock",
			pos:  position{line: 229, col: 1, offset: 9181},
			expr: &actionExpr{
				pos: position{line: 229, col: 16, offset: 9196},
				run: (*parser).callonSourceBlock1,
				expr: &seqExpr{
					pos: position{line: 229, col: 16, offset: 9196},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 229, col: 16, offset: 9196},
							name: "SourceBlockDelimiter",
						},
						&ruleRefExpr{
							pos:  position{line: 229, col: 37, offset: 9217},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 229, col: 45, offset: 9225},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 229, col: 53, offset: 9233},
								expr: &ruleRefExpr{
									pos:  position{line: 229, col: 54, offset: 9234},
									name: "SourceBlockLine",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 229, col: 73, offset: 9253},
							name: "SourceBlockDelimiter",
						},
					},
				},
			},
		},
		{
			name: "SourceBlockDelimiter",
			pos:  position{line: 233, col: 1, offset: 9358},
			expr: &litMatcher{
				pos:        position{line: 233, col: 25, offset: 9382},
				val:        "```",
				ignoreCase: false,
			},
		},
		{
			name: "SourceBlockLine",
			pos:  position{line: 235, col: 1, offset: 9389},
			expr: &seqExpr{
				pos: position{line: 235, col: 20, offset: 9408},
				exprs: []interface{}{
					&zeroOrMoreExpr{
						pos: position{line: 235, col: 20, offset: 9408},
						expr: &seqExpr{
							pos: position{line: 235, col: 21, offset: 9409},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 235, col: 21, offset: 9409},
									expr: &ruleRefExpr{
										pos:  position{line: 235, col: 22, offset: 9410},
										name: "EOL",
									},
								},
								&anyMatcher{
									line: 235, col: 26, offset: 9414,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 235, col: 30, offset: 9418},
						name: "NEWLINE",
					},
				},
			},
		},
		{
			name: "ElementAttribute",
			pos:  position{line: 240, col: 1, offset: 9536},
			expr: &labeledExpr{
				pos:   position{line: 240, col: 21, offset: 9556},
				label: "meta",
				expr: &choiceExpr{
					pos: position{line: 240, col: 27, offset: 9562},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 240, col: 27, offset: 9562},
							name: "ElementLink",
						},
						&ruleRefExpr{
							pos:  position{line: 240, col: 41, offset: 9576},
							name: "ElementID",
						},
						&ruleRefExpr{
							pos:  position{line: 240, col: 53, offset: 9588},
							name: "ElementTitle",
						},
					},
				},
			},
		},
		{
			name: "ElementLink",
			pos:  position{line: 243, col: 1, offset: 9659},
			expr: &actionExpr{
				pos: position{line: 243, col: 16, offset: 9674},
				run: (*parser).callonElementLink1,
				expr: &seqExpr{
					pos: position{line: 243, col: 16, offset: 9674},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 243, col: 16, offset: 9674},
							val:        "[",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 243, col: 20, offset: 9678},
							expr: &ruleRefExpr{
								pos:  position{line: 243, col: 20, offset: 9678},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 243, col: 24, offset: 9682},
							val:        "link",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 243, col: 31, offset: 9689},
							expr: &ruleRefExpr{
								pos:  position{line: 243, col: 31, offset: 9689},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 243, col: 35, offset: 9693},
							val:        "=",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 243, col: 39, offset: 9697},
							expr: &ruleRefExpr{
								pos:  position{line: 243, col: 39, offset: 9697},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 243, col: 43, offset: 9701},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 243, col: 48, offset: 9706},
								name: "URL",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 243, col: 52, offset: 9710},
							expr: &ruleRefExpr{
								pos:  position{line: 243, col: 52, offset: 9710},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 243, col: 56, offset: 9714},
							val:        "]",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 243, col: 60, offset: 9718},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ElementID",
			pos:  position{line: 248, col: 1, offset: 9828},
			expr: &actionExpr{
				pos: position{line: 248, col: 14, offset: 9841},
				run: (*parser).callonElementID1,
				expr: &seqExpr{
					pos: position{line: 248, col: 14, offset: 9841},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 248, col: 14, offset: 9841},
							val:        "[",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 248, col: 18, offset: 9845},
							expr: &ruleRefExpr{
								pos:  position{line: 248, col: 18, offset: 9845},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 248, col: 22, offset: 9849},
							val:        "#",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 248, col: 26, offset: 9853},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 248, col: 30, offset: 9857},
								name: "ID",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 248, col: 34, offset: 9861},
							expr: &ruleRefExpr{
								pos:  position{line: 248, col: 34, offset: 9861},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 248, col: 38, offset: 9865},
							val:        "]",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 248, col: 42, offset: 9869},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ElementTitle",
			pos:  position{line: 253, col: 1, offset: 9977},
			expr: &actionExpr{
				pos: position{line: 253, col: 17, offset: 9993},
				run: (*parser).callonElementTitle1,
				expr: &seqExpr{
					pos: position{line: 253, col: 17, offset: 9993},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 253, col: 17, offset: 9993},
							val:        ".",
							ignoreCase: false,
						},
						&notExpr{
							pos: position{line: 253, col: 21, offset: 9997},
							expr: &ruleRefExpr{
								pos:  position{line: 253, col: 22, offset: 9998},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 253, col: 25, offset: 10001},
							label: "title",
							expr: &oneOrMoreExpr{
								pos: position{line: 253, col: 31, offset: 10007},
								expr: &seqExpr{
									pos: position{line: 253, col: 32, offset: 10008},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 253, col: 32, offset: 10008},
											expr: &ruleRefExpr{
												pos:  position{line: 253, col: 33, offset: 10009},
												name: "NEWLINE",
											},
										},
										&anyMatcher{
											line: 253, col: 41, offset: 10017,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 253, col: 45, offset: 10021},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "Word",
			pos:  position{line: 260, col: 1, offset: 10192},
			expr: &actionExpr{
				pos: position{line: 260, col: 9, offset: 10200},
				run: (*parser).callonWord1,
				expr: &oneOrMoreExpr{
					pos: position{line: 260, col: 9, offset: 10200},
					expr: &seqExpr{
						pos: position{line: 260, col: 10, offset: 10201},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 260, col: 10, offset: 10201},
								expr: &ruleRefExpr{
									pos:  position{line: 260, col: 11, offset: 10202},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 260, col: 19, offset: 10210},
								expr: &ruleRefExpr{
									pos:  position{line: 260, col: 20, offset: 10211},
									name: "WS",
								},
							},
							&anyMatcher{
								line: 260, col: 23, offset: 10214,
							},
						},
					},
				},
			},
		},
		{
			name: "BlankLine",
			pos:  position{line: 264, col: 1, offset: 10254},
			expr: &actionExpr{
				pos: position{line: 264, col: 14, offset: 10267},
				run: (*parser).callonBlankLine1,
				expr: &seqExpr{
					pos: position{line: 264, col: 14, offset: 10267},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 264, col: 14, offset: 10267},
							expr: &ruleRefExpr{
								pos:  position{line: 264, col: 15, offset: 10268},
								name: "EOF",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 264, col: 19, offset: 10272},
							expr: &ruleRefExpr{
								pos:  position{line: 264, col: 19, offset: 10272},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 264, col: 23, offset: 10276},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "URL",
			pos:  position{line: 268, col: 1, offset: 10317},
			expr: &actionExpr{
				pos: position{line: 268, col: 8, offset: 10324},
				run: (*parser).callonURL1,
				expr: &oneOrMoreExpr{
					pos: position{line: 268, col: 8, offset: 10324},
					expr: &seqExpr{
						pos: position{line: 268, col: 9, offset: 10325},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 268, col: 9, offset: 10325},
								expr: &ruleRefExpr{
									pos:  position{line: 268, col: 10, offset: 10326},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 268, col: 18, offset: 10334},
								expr: &ruleRefExpr{
									pos:  position{line: 268, col: 19, offset: 10335},
									name: "WS",
								},
							},
							&notExpr{
								pos: position{line: 268, col: 22, offset: 10338},
								expr: &litMatcher{
									pos:        position{line: 268, col: 23, offset: 10339},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 268, col: 27, offset: 10343},
								expr: &litMatcher{
									pos:        position{line: 268, col: 28, offset: 10344},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 268, col: 32, offset: 10348,
							},
						},
					},
				},
			},
		},
		{
			name: "ID",
			pos:  position{line: 272, col: 1, offset: 10388},
			expr: &actionExpr{
				pos: position{line: 272, col: 7, offset: 10394},
				run: (*parser).callonID1,
				expr: &oneOrMoreExpr{
					pos: position{line: 272, col: 7, offset: 10394},
					expr: &seqExpr{
						pos: position{line: 272, col: 8, offset: 10395},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 272, col: 8, offset: 10395},
								expr: &ruleRefExpr{
									pos:  position{line: 272, col: 9, offset: 10396},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 272, col: 17, offset: 10404},
								expr: &ruleRefExpr{
									pos:  position{line: 272, col: 18, offset: 10405},
									name: "WS",
								},
							},
							&notExpr{
								pos: position{line: 272, col: 21, offset: 10408},
								expr: &litMatcher{
									pos:        position{line: 272, col: 22, offset: 10409},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 272, col: 26, offset: 10413},
								expr: &litMatcher{
									pos:        position{line: 272, col: 27, offset: 10414},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 272, col: 31, offset: 10418,
							},
						},
					},
				},
			},
		},
		{
			name: "URL_TEXT",
			pos:  position{line: 276, col: 1, offset: 10458},
			expr: &actionExpr{
				pos: position{line: 276, col: 13, offset: 10470},
				run: (*parser).callonURL_TEXT1,
				expr: &oneOrMoreExpr{
					pos: position{line: 276, col: 13, offset: 10470},
					expr: &seqExpr{
						pos: position{line: 276, col: 14, offset: 10471},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 276, col: 14, offset: 10471},
								expr: &ruleRefExpr{
									pos:  position{line: 276, col: 15, offset: 10472},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 276, col: 23, offset: 10480},
								expr: &litMatcher{
									pos:        position{line: 276, col: 24, offset: 10481},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 276, col: 28, offset: 10485},
								expr: &litMatcher{
									pos:        position{line: 276, col: 29, offset: 10486},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 276, col: 33, offset: 10490,
							},
						},
					},
				},
			},
		},
		{
			name: "URL_SCHEME",
			pos:  position{line: 280, col: 1, offset: 10530},
			expr: &choiceExpr{
				pos: position{line: 280, col: 15, offset: 10544},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 280, col: 15, offset: 10544},
						val:        "http://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 280, col: 27, offset: 10556},
						val:        "https://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 280, col: 40, offset: 10569},
						val:        "ftp://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 280, col: 51, offset: 10580},
						val:        "irc://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 280, col: 62, offset: 10591},
						val:        "mailto:",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "DIGIT",
			pos:  position{line: 282, col: 1, offset: 10602},
			expr: &charClassMatcher{
				pos:        position{line: 282, col: 13, offset: 10614},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NEWLINE",
			pos:  position{line: 284, col: 1, offset: 10621},
			expr: &choiceExpr{
				pos: position{line: 284, col: 13, offset: 10633},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 284, col: 13, offset: 10633},
						val:        "\r\n",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 284, col: 22, offset: 10642},
						val:        "\r",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 284, col: 29, offset: 10649},
						val:        "\n",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "WS",
			pos:  position{line: 286, col: 1, offset: 10655},
			expr: &choiceExpr{
				pos: position{line: 286, col: 13, offset: 10667},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 286, col: 13, offset: 10667},
						val:        " ",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 286, col: 19, offset: 10673},
						run: (*parser).callonWS3,
						expr: &litMatcher{
							pos:        position{line: 286, col: 19, offset: 10673},
							val:        "\t",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 290, col: 1, offset: 10718},
			expr: &notExpr{
				pos: position{line: 290, col: 13, offset: 10730},
				expr: &anyMatcher{
					line: 290, col: 14, offset: 10731,
				},
			},
		},
		{
			name: "EOL",
			pos:  position{line: 292, col: 1, offset: 10734},
			expr: &choiceExpr{
				pos: position{line: 292, col: 13, offset: 10746},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 292, col: 13, offset: 10746},
						name: "NEWLINE",
					},
					&ruleRefExpr{
						pos:  position{line: 292, col: 23, offset: 10756},
						name: "EOF",
					},
				},
			},
		},
	},
}

func (c *current) onDocument1(blocks interface{}) (interface{}, error) {
	return types.NewDocument(blocks.([]interface{}))
}

func (p *parser) callonDocument1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocument1(stack["blocks"])
}

func (c *current) onDocumentBlock1(content interface{}) (interface{}, error) {
	return content.(types.DocElement), nil
}

func (p *parser) callonDocumentBlock1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocumentBlock1(stack["content"])
}

func (c *current) onSection11(heading, elements interface{}) (interface{}, error) {
	return types.NewSection(heading.(*types.Heading), elements.([]interface{}))
}

func (p *parser) callonSection11() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection11(stack["heading"], stack["elements"])
}

func (c *current) onSection1Block1(content interface{}) (interface{}, error) {
	return content.(types.DocElement), nil
}

func (p *parser) callonSection1Block1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection1Block1(stack["content"])
}

func (c *current) onSection21(heading, elements interface{}) (interface{}, error) {
	return types.NewSection(heading.(*types.Heading), elements.([]interface{}))
}

func (p *parser) callonSection21() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection21(stack["heading"], stack["elements"])
}

func (c *current) onSection2Block1(content interface{}) (interface{}, error) {
	return content.(types.DocElement), nil
}

func (p *parser) callonSection2Block1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection2Block1(stack["content"])
}

func (c *current) onSection31(heading, elements interface{}) (interface{}, error) {
	return types.NewSection(heading.(*types.Heading), elements.([]interface{}))
}

func (p *parser) callonSection31() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection31(stack["heading"], stack["elements"])
}

func (c *current) onSection3Block1(content interface{}) (interface{}, error) {
	return content.(types.DocElement), nil
}

func (p *parser) callonSection3Block1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection3Block1(stack["content"])
}

func (c *current) onSection41(heading, elements interface{}) (interface{}, error) {
	return types.NewSection(heading.(*types.Heading), elements.([]interface{}))
}

func (p *parser) callonSection41() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection41(stack["heading"], stack["elements"])
}

func (c *current) onSection4Block1(content interface{}) (interface{}, error) {
	return content.(types.DocElement), nil
}

func (p *parser) callonSection4Block1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection4Block1(stack["content"])
}

func (c *current) onSection51(heading, elements interface{}) (interface{}, error) {
	return types.NewSection(heading.(*types.Heading), elements.([]interface{}))
}

func (p *parser) callonSection51() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection51(stack["heading"], stack["elements"])
}

func (c *current) onSection5Block1(content interface{}) (interface{}, error) {
	return content.(types.DocElement), nil
}

func (p *parser) callonSection5Block1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection5Block1(stack["content"])
}

func (c *current) onSection61(heading, elements interface{}) (interface{}, error) {
	return types.NewSection(heading.(*types.Heading), elements.([]interface{}))
}

func (p *parser) callonSection61() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection61(stack["heading"], stack["elements"])
}

func (c *current) onSection6Block1(content interface{}) (interface{}, error) {
	return content.(types.DocElement), nil
}

func (p *parser) callonSection6Block1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection6Block1(stack["content"])
}

func (c *current) onHeading11(attributes, level, content interface{}) (interface{}, error) {
	//TODO: replace `(BlankLine? / EOF)` with `EOL` to allow for immediate attributes or any other content ?
	return types.NewHeading(1, content.(*types.InlineContent), attributes.([]interface{}))
}

func (p *parser) callonHeading11() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onHeading11(stack["attributes"], stack["level"], stack["content"])
}

func (c *current) onHeading21(attributes, level, content interface{}) (interface{}, error) {
	return types.NewHeading(2, content.(*types.InlineContent), attributes.([]interface{}))
}

func (p *parser) callonHeading21() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onHeading21(stack["attributes"], stack["level"], stack["content"])
}

func (c *current) onHeading31(attributes, level, content interface{}) (interface{}, error) {
	return types.NewHeading(3, content.(*types.InlineContent), attributes.([]interface{}))
}

func (p *parser) callonHeading31() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onHeading31(stack["attributes"], stack["level"], stack["content"])
}

func (c *current) onHeading41(attributes, level, content interface{}) (interface{}, error) {
	return types.NewHeading(4, content.(*types.InlineContent), attributes.([]interface{}))
}

func (p *parser) callonHeading41() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onHeading41(stack["attributes"], stack["level"], stack["content"])
}

func (c *current) onHeading51(attributes, level, content interface{}) (interface{}, error) {
	return types.NewHeading(5, content.(*types.InlineContent), attributes.([]interface{}))
}

func (p *parser) callonHeading51() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onHeading51(stack["attributes"], stack["level"], stack["content"])
}

func (c *current) onHeading61(attributes, level, content interface{}) (interface{}, error) {
	return types.NewHeading(6, content.(*types.InlineContent), attributes.([]interface{}))
}

func (p *parser) callonHeading61() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onHeading61(stack["attributes"], stack["level"], stack["content"])
}

func (c *current) onDocumentAttributeDeclarationWithNameOnly1(name interface{}) (interface{}, error) {
	return types.NewDocumentAttributeDeclaration(name.([]interface{}), nil)
}

func (p *parser) callonDocumentAttributeDeclarationWithNameOnly1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocumentAttributeDeclarationWithNameOnly1(stack["name"])
}

func (c *current) onDocumentAttributeDeclarationWithNameAndValue1(name, value interface{}) (interface{}, error) {
	return types.NewDocumentAttributeDeclaration(name.([]interface{}), value.([]interface{}))
}

func (p *parser) callonDocumentAttributeDeclarationWithNameAndValue1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocumentAttributeDeclarationWithNameAndValue1(stack["name"], stack["value"])
}

func (c *current) onDocumentAttributeResetWithHeadingBangSymbol1(name interface{}) (interface{}, error) {
	return types.NewDocumentAttributeReset(name.([]interface{}))
}

func (p *parser) callonDocumentAttributeResetWithHeadingBangSymbol1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocumentAttributeResetWithHeadingBangSymbol1(stack["name"])
}

func (c *current) onDocumentAttributeResetWithTrailingBangSymbol1(name interface{}) (interface{}, error) {
	return types.NewDocumentAttributeReset(name.([]interface{}))
}

func (p *parser) callonDocumentAttributeResetWithTrailingBangSymbol1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocumentAttributeResetWithTrailingBangSymbol1(stack["name"])
}

func (c *current) onDocumentAttributeSubstitution1(name interface{}) (interface{}, error) {
	return types.NewDocumentAttributeSubstitution(name.([]interface{}))
}

func (p *parser) callonDocumentAttributeSubstitution1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocumentAttributeSubstitution1(stack["name"])
}

func (c *current) onList1(attributes, elements interface{}) (interface{}, error) {
	return types.NewList(elements.([]interface{}), attributes.([]interface{}))
}

func (p *parser) callonList1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onList1(stack["attributes"], stack["elements"])
}

func (c *current) onListItem1(level, content interface{}) (interface{}, error) {
	return types.NewListItem(level, content.(*types.ListItemContent), nil)
}

func (p *parser) callonListItem1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onListItem1(stack["level"], stack["content"])
}

func (c *current) onListItemContent1(lines interface{}) (interface{}, error) {
	return types.NewListItemContent(c.text, lines.([]interface{}))
}

func (p *parser) callonListItemContent1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onListItemContent1(stack["lines"])
}

func (c *current) onParagraph1(attributes, lines interface{}) (interface{}, error) {
	return types.NewParagraph(c.text, lines.([]interface{}), attributes.([]interface{}))
}

func (p *parser) callonParagraph1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onParagraph1(stack["attributes"], stack["lines"])
}

func (c *current) onInlineContent1(elements interface{}) (interface{}, error) {
	return types.NewInlineContent(c.text, elements.([]interface{}))
}

func (p *parser) callonInlineContent1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInlineContent1(stack["elements"])
}

func (c *current) onBoldText1(content interface{}) (interface{}, error) {
	return types.NewQuotedText(types.Bold, content.([]interface{}))
}

func (p *parser) callonBoldText1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBoldText1(stack["content"])
}

func (c *current) onItalicText1(content interface{}) (interface{}, error) {
	return types.NewQuotedText(types.Italic, content.([]interface{}))
}

func (p *parser) callonItalicText1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onItalicText1(stack["content"])
}

func (c *current) onMonospaceText1(content interface{}) (interface{}, error) {
	return types.NewQuotedText(types.Monospace, content.([]interface{}))
}

func (p *parser) callonMonospaceText1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMonospaceText1(stack["content"])
}

func (c *current) onExternalLink1(url, text interface{}) (interface{}, error) {
	if text != nil {
		return types.NewExternalLink(url.([]interface{}), text.([]interface{}))
	}
	return types.NewExternalLink(url.([]interface{}), nil)
}

func (p *parser) callonExternalLink1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onExternalLink1(stack["url"], stack["text"])
}

func (c *current) onBlockImage1(attributes, image interface{}) (interface{}, error) {
	// here we can ignore the blank line in the returned element
	return types.NewBlockImage(c.text, *image.(*types.ImageMacro), attributes.([]interface{}))
}

func (p *parser) callonBlockImage1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBlockImage1(stack["attributes"], stack["image"])
}

func (c *current) onBlockImageMacro1(path, attributes interface{}) (interface{}, error) {
	return types.NewImageMacro(c.text, path.(string), attributes)
}

func (p *parser) callonBlockImageMacro1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBlockImageMacro1(stack["path"], stack["attributes"])
}

func (c *current) onInlineImage1(image interface{}) (interface{}, error) {
	// here we can ignore the blank line in the returned element
	return types.NewInlineImage(c.text, *image.(*types.ImageMacro))
}

func (p *parser) callonInlineImage1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInlineImage1(stack["image"])
}

func (c *current) onInlineImageMacro1(path, attributes interface{}) (interface{}, error) {
	return types.NewImageMacro(c.text, path.(string), attributes)
}

func (p *parser) callonInlineImageMacro1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInlineImageMacro1(stack["path"], stack["attributes"])
}

func (c *current) onSourceBlock1(content interface{}) (interface{}, error) {
	return types.NewDelimitedBlock(types.SourceBlock, content.([]interface{}))
}

func (p *parser) callonSourceBlock1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSourceBlock1(stack["content"])
}

func (c *current) onElementLink1(path interface{}) (interface{}, error) {
	return types.NewElementLink(path.(string))
}

func (p *parser) callonElementLink1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onElementLink1(stack["path"])
}

func (c *current) onElementID1(id interface{}) (interface{}, error) {
	return types.NewElementID(id.(string))
}

func (p *parser) callonElementID1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onElementID1(stack["id"])
}

func (c *current) onElementTitle1(title interface{}) (interface{}, error) {
	return types.NewElementTitle(title.([]interface{}))
}

func (p *parser) callonElementTitle1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onElementTitle1(stack["title"])
}

func (c *current) onWord1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonWord1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onWord1()
}

func (c *current) onBlankLine1() (interface{}, error) {
	return types.NewBlankLine()
}

func (p *parser) callonBlankLine1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBlankLine1()
}

func (c *current) onURL1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonURL1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onURL1()
}

func (c *current) onID1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonID1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onID1()
}

func (c *current) onURL_TEXT1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonURL_TEXT1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onURL_TEXT1()
}

func (c *current) onWS3() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonWS3() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onWS3()
}

var (
	// errNoRule is returned when the grammar to parse has no rule.
	errNoRule = errors.New("grammar has no rule")

	// errInvalidEncoding is returned when the source is not properly
	// utf8-encoded.
	errInvalidEncoding = errors.New("invalid encoding")
)

// Option is a function that can set an option on the parser. It returns
// the previous setting as an Option.
type Option func(*parser) Option

// Debug creates an Option to set the debug flag to b. When set to true,
// debugging information is printed to stdout while parsing.
//
// The default is false.
func Debug(b bool) Option {
	return func(p *parser) Option {
		old := p.debug
		p.debug = b
		return Debug(old)
	}
}

// Memoize creates an Option to set the memoize flag to b. When set to true,
// the parser will cache all results so each expression is evaluated only
// once. This guarantees linear parsing time even for pathological cases,
// at the expense of more memory and slower times for typical cases.
//
// The default is false.
func Memoize(b bool) Option {
	return func(p *parser) Option {
		old := p.memoize
		p.memoize = b
		return Memoize(old)
	}
}

// Recover creates an Option to set the recover flag to b. When set to
// true, this causes the parser to recover from panics and convert it
// to an error. Setting it to false can be useful while debugging to
// access the full stack trace.
//
// The default is true.
func Recover(b bool) Option {
	return func(p *parser) Option {
		old := p.recover
		p.recover = b
		return Recover(old)
	}
}

// GlobalStore creates an Option to set a key to a certain value in
// the globalStore.
func GlobalStore(key string, value interface{}) Option {
	return func(p *parser) Option {
		old := p.cur.globalStore[key]
		p.cur.globalStore[key] = value
		return GlobalStore(key, old)
	}
}

// ParseFile parses the file identified by filename.
func ParseFile(filename string, opts ...Option) (i interface{}, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := f.Close(); closeErr != nil {
			err = closeErr
		}
	}()
	return ParseReader(filename, f, opts...)
}

// ParseReader parses the data from r using filename as information in the
// error messages.
func ParseReader(filename string, r io.Reader, opts ...Option) (interface{}, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return Parse(filename, b, opts...)
}

// Parse parses the data from b using filename as information in the
// error messages.
func Parse(filename string, b []byte, opts ...Option) (interface{}, error) {
	return newParser(filename, b, opts...).parse(g)
}

// position records a position in the text.
type position struct {
	line, col, offset int
}

func (p position) String() string {
	return fmt.Sprintf("%d:%d [%d]", p.line, p.col, p.offset)
}

// savepoint stores all state required to go back to this point in the
// parser.
type savepoint struct {
	position
	rn rune
	w  int
}

type current struct {
	pos  position // start position of the match
	text []byte   // raw text of the match

	// the globalStore allows the parser to store arbitrary values
	globalStore map[string]interface{}
}

// the AST types...

type grammar struct {
	pos   position
	rules []*rule
}

type rule struct {
	pos         position
	name        string
	displayName string
	expr        interface{}
}

type choiceExpr struct {
	pos          position
	alternatives []interface{}
}

type actionExpr struct {
	pos  position
	expr interface{}
	run  func(*parser) (interface{}, error)
}

type seqExpr struct {
	pos   position
	exprs []interface{}
}

type labeledExpr struct {
	pos   position
	label string
	expr  interface{}
}

type expr struct {
	pos  position
	expr interface{}
}

type andExpr expr
type notExpr expr
type zeroOrOneExpr expr
type zeroOrMoreExpr expr
type oneOrMoreExpr expr

type ruleRefExpr struct {
	pos  position
	name string
}

type andCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type notCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type litMatcher struct {
	pos        position
	val        string
	ignoreCase bool
}

type charClassMatcher struct {
	pos             position
	val             string
	basicLatinChars [128]bool
	chars           []rune
	ranges          []rune
	classes         []*unicode.RangeTable
	ignoreCase      bool
	inverted        bool
}

type anyMatcher position

// errList cumulates the errors found by the parser.
type errList []error

func (e *errList) add(err error) {
	*e = append(*e, err)
}

func (e errList) err() error {
	if len(e) == 0 {
		return nil
	}
	e.dedupe()
	return e
}

func (e *errList) dedupe() {
	var cleaned []error
	set := make(map[string]bool)
	for _, err := range *e {
		if msg := err.Error(); !set[msg] {
			set[msg] = true
			cleaned = append(cleaned, err)
		}
	}
	*e = cleaned
}

func (e errList) Error() string {
	switch len(e) {
	case 0:
		return ""
	case 1:
		return e[0].Error()
	default:
		var buf bytes.Buffer

		for i, err := range e {
			if i > 0 {
				buf.WriteRune('\n')
			}
			buf.WriteString(err.Error())
		}
		return buf.String()
	}
}

// parserError wraps an error with a prefix indicating the rule in which
// the error occurred. The original error is stored in the Inner field.
type parserError struct {
	Inner    error
	pos      position
	prefix   string
	expected []string
}

// Error returns the error message.
func (p *parserError) Error() string {
	return p.prefix + ": " + p.Inner.Error()
}

// newParser creates a parser with the specified input source and options.
func newParser(filename string, b []byte, opts ...Option) *parser {
	p := &parser{
		filename: filename,
		errs:     new(errList),
		data:     b,
		pt:       savepoint{position: position{line: 1}},
		recover:  true,
		cur: current{
			globalStore: make(map[string]interface{}),
		},
		maxFailPos:      position{col: 1, line: 1},
		maxFailExpected: make([]string, 0, 20),
	}
	p.setOptions(opts)
	return p
}

// setOptions applies the options to the parser.
func (p *parser) setOptions(opts []Option) {
	for _, opt := range opts {
		opt(p)
	}
}

type resultTuple struct {
	v   interface{}
	b   bool
	end savepoint
}

type parser struct {
	filename string
	pt       savepoint
	cur      current

	data []byte
	errs *errList

	depth   int
	recover bool
	debug   bool

	memoize bool
	// memoization table for the packrat algorithm:
	// map[offset in source] map[expression or rule] {value, match}
	memo map[int]map[interface{}]resultTuple

	// rules table, maps the rule identifier to the rule node
	rules map[string]*rule
	// variables stack, map of label to value
	vstack []map[string]interface{}
	// rule stack, allows identification of the current rule in errors
	rstack []*rule

	// stats
	exprCnt int

	// parse fail
	maxFailPos            position
	maxFailExpected       []string
	maxFailInvertExpected bool
}

// push a variable set on the vstack.
func (p *parser) pushV() {
	if cap(p.vstack) == len(p.vstack) {
		// create new empty slot in the stack
		p.vstack = append(p.vstack, nil)
	} else {
		// slice to 1 more
		p.vstack = p.vstack[:len(p.vstack)+1]
	}

	// get the last args set
	m := p.vstack[len(p.vstack)-1]
	if m != nil && len(m) == 0 {
		// empty map, all good
		return
	}

	m = make(map[string]interface{})
	p.vstack[len(p.vstack)-1] = m
}

// pop a variable set from the vstack.
func (p *parser) popV() {
	// if the map is not empty, clear it
	m := p.vstack[len(p.vstack)-1]
	if len(m) > 0 {
		// GC that map
		p.vstack[len(p.vstack)-1] = nil
	}
	p.vstack = p.vstack[:len(p.vstack)-1]
}

func (p *parser) print(prefix, s string) string {
	if !p.debug {
		return s
	}

	fmt.Printf("%s %d:%d:%d: %s [%#U]\n",
		prefix, p.pt.line, p.pt.col, p.pt.offset, s, p.pt.rn)
	return s
}

func (p *parser) in(s string) string {
	p.depth++
	return p.print(strings.Repeat(" ", p.depth)+">", s)
}

func (p *parser) out(s string) string {
	p.depth--
	return p.print(strings.Repeat(" ", p.depth)+"<", s)
}

func (p *parser) addErr(err error) {
	p.addErrAt(err, p.pt.position, []string{})
}

func (p *parser) addErrAt(err error, pos position, expected []string) {
	var buf bytes.Buffer
	if p.filename != "" {
		buf.WriteString(p.filename)
	}
	if buf.Len() > 0 {
		buf.WriteString(":")
	}
	buf.WriteString(fmt.Sprintf("%d:%d (%d)", pos.line, pos.col, pos.offset))
	if len(p.rstack) > 0 {
		if buf.Len() > 0 {
			buf.WriteString(": ")
		}
		rule := p.rstack[len(p.rstack)-1]
		if rule.displayName != "" {
			buf.WriteString("rule " + rule.displayName)
		} else {
			buf.WriteString("rule " + rule.name)
		}
	}
	pe := &parserError{Inner: err, pos: pos, prefix: buf.String(), expected: expected}
	p.errs.add(pe)
}

func (p *parser) failAt(fail bool, pos position, want string) {
	// process fail if parsing fails and not inverted or parsing succeeds and invert is set
	if fail == p.maxFailInvertExpected {
		if pos.offset < p.maxFailPos.offset {
			return
		}

		if pos.offset > p.maxFailPos.offset {
			p.maxFailPos = pos
			p.maxFailExpected = p.maxFailExpected[:0]
		}

		if p.maxFailInvertExpected {
			want = "!" + want
		}
		p.maxFailExpected = append(p.maxFailExpected, want)
	}
}

// read advances the parser to the next rune.
func (p *parser) read() {
	p.pt.offset += p.pt.w
	rn, n := utf8.DecodeRune(p.data[p.pt.offset:])
	p.pt.rn = rn
	p.pt.w = n
	p.pt.col++
	if rn == '\n' {
		p.pt.line++
		p.pt.col = 0
	}

	if rn == utf8.RuneError {
		if n == 1 {
			p.addErr(errInvalidEncoding)
		}
	}
}

// restore parser position to the savepoint pt.
func (p *parser) restore(pt savepoint) {
	if p.debug {
		defer p.out(p.in("restore"))
	}
	if pt.offset == p.pt.offset {
		return
	}
	p.pt = pt
}

// get the slice of bytes from the savepoint start to the current position.
func (p *parser) sliceFrom(start savepoint) []byte {
	return p.data[start.position.offset:p.pt.position.offset]
}

func (p *parser) getMemoized(node interface{}) (resultTuple, bool) {
	if len(p.memo) == 0 {
		return resultTuple{}, false
	}
	m := p.memo[p.pt.offset]
	if len(m) == 0 {
		return resultTuple{}, false
	}
	res, ok := m[node]
	return res, ok
}

func (p *parser) setMemoized(pt savepoint, node interface{}, tuple resultTuple) {
	if p.memo == nil {
		p.memo = make(map[int]map[interface{}]resultTuple)
	}
	m := p.memo[pt.offset]
	if m == nil {
		m = make(map[interface{}]resultTuple)
		p.memo[pt.offset] = m
	}
	m[node] = tuple
}

func (p *parser) buildRulesTable(g *grammar) {
	p.rules = make(map[string]*rule, len(g.rules))
	for _, r := range g.rules {
		p.rules[r.name] = r
	}
}

func (p *parser) parse(g *grammar) (val interface{}, err error) {
	if len(g.rules) == 0 {
		p.addErr(errNoRule)
		return nil, p.errs.err()
	}

	// TODO : not super critical but this could be generated
	p.buildRulesTable(g)

	if p.recover {
		// panic can be used in action code to stop parsing immediately
		// and return the panic as an error.
		defer func() {
			if e := recover(); e != nil {
				if p.debug {
					defer p.out(p.in("panic handler"))
				}
				val = nil
				switch e := e.(type) {
				case error:
					p.addErr(e)
				default:
					p.addErr(fmt.Errorf("%v", e))
				}
				err = p.errs.err()
			}
		}()
	}

	// start rule is rule [0]
	p.read() // advance to first rune
	val, ok := p.parseRule(g.rules[0])
	if !ok {
		if len(*p.errs) == 0 {
			// If parsing fails, but no errors have been recorded, the expected values
			// for the farthest parser position are returned as error.
			maxFailExpectedMap := make(map[string]struct{}, len(p.maxFailExpected))
			for _, v := range p.maxFailExpected {
				maxFailExpectedMap[v] = struct{}{}
			}
			expected := make([]string, 0, len(maxFailExpectedMap))
			eof := false
			if _, ok := maxFailExpectedMap["!."]; ok {
				delete(maxFailExpectedMap, "!.")
				eof = true
			}
			for k := range maxFailExpectedMap {
				expected = append(expected, k)
			}
			sort.Strings(expected)
			if eof {
				expected = append(expected, "EOF")
			}
			p.addErrAt(errors.New("no match found, expected: "+listJoin(expected, ", ", "or")), p.maxFailPos, expected)
		}
		return nil, p.errs.err()
	}
	return val, p.errs.err()
}

func listJoin(list []string, sep string, lastSep string) string {
	switch len(list) {
	case 0:
		return ""
	case 1:
		return list[0]
	default:
		return fmt.Sprintf("%s %s %s", strings.Join(list[:len(list)-1], sep), lastSep, list[len(list)-1])
	}
}

func (p *parser) parseRule(rule *rule) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRule " + rule.name))
	}

	if p.memoize {
		res, ok := p.getMemoized(rule)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
	}

	start := p.pt
	p.rstack = append(p.rstack, rule)
	p.pushV()
	val, ok := p.parseExpr(rule.expr)
	p.popV()
	p.rstack = p.rstack[:len(p.rstack)-1]
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth)+"MATCH", string(p.sliceFrom(start)))
	}

	if p.memoize {
		p.setMemoized(start, rule, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseExpr(expr interface{}) (interface{}, bool) {
	var pt savepoint

	if p.memoize {
		res, ok := p.getMemoized(expr)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
		pt = p.pt
	}

	p.exprCnt++
	var val interface{}
	var ok bool
	switch expr := expr.(type) {
	case *actionExpr:
		val, ok = p.parseActionExpr(expr)
	case *andCodeExpr:
		val, ok = p.parseAndCodeExpr(expr)
	case *andExpr:
		val, ok = p.parseAndExpr(expr)
	case *anyMatcher:
		val, ok = p.parseAnyMatcher(expr)
	case *charClassMatcher:
		val, ok = p.parseCharClassMatcher(expr)
	case *choiceExpr:
		val, ok = p.parseChoiceExpr(expr)
	case *labeledExpr:
		val, ok = p.parseLabeledExpr(expr)
	case *litMatcher:
		val, ok = p.parseLitMatcher(expr)
	case *notCodeExpr:
		val, ok = p.parseNotCodeExpr(expr)
	case *notExpr:
		val, ok = p.parseNotExpr(expr)
	case *oneOrMoreExpr:
		val, ok = p.parseOneOrMoreExpr(expr)
	case *ruleRefExpr:
		val, ok = p.parseRuleRefExpr(expr)
	case *seqExpr:
		val, ok = p.parseSeqExpr(expr)
	case *zeroOrMoreExpr:
		val, ok = p.parseZeroOrMoreExpr(expr)
	case *zeroOrOneExpr:
		val, ok = p.parseZeroOrOneExpr(expr)
	default:
		panic(fmt.Sprintf("unknown expression type %T", expr))
	}
	if p.memoize {
		p.setMemoized(pt, expr, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseActionExpr(act *actionExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseActionExpr"))
	}

	start := p.pt
	val, ok := p.parseExpr(act.expr)
	if ok {
		p.cur.pos = start.position
		p.cur.text = p.sliceFrom(start)
		actVal, err := act.run(p)
		if err != nil {
			p.addErrAt(err, start.position, []string{})
		}
		val = actVal
	}
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth)+"MATCH", string(p.sliceFrom(start)))
	}
	return val, ok
}

func (p *parser) parseAndCodeExpr(and *andCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndCodeExpr"))
	}

	ok, err := and.run(p)
	if err != nil {
		p.addErr(err)
	}
	return nil, ok
}

func (p *parser) parseAndExpr(and *andExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndExpr"))
	}

	pt := p.pt
	p.pushV()
	_, ok := p.parseExpr(and.expr)
	p.popV()
	p.restore(pt)
	return nil, ok
}

func (p *parser) parseAnyMatcher(any *anyMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAnyMatcher"))
	}

	if p.pt.rn != utf8.RuneError {
		start := p.pt
		p.read()
		p.failAt(true, start.position, ".")
		return p.sliceFrom(start), true
	}
	p.failAt(false, p.pt.position, ".")
	return nil, false
}

func (p *parser) parseCharClassMatcher(chr *charClassMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseCharClassMatcher"))
	}

	cur := p.pt.rn
	start := p.pt

	// can't match EOF
	if cur == utf8.RuneError {
		p.failAt(false, start.position, chr.val)
		return nil, false
	}

	if chr.ignoreCase {
		cur = unicode.ToLower(cur)
	}

	// try to match in the list of available chars
	for _, rn := range chr.chars {
		if rn == cur {
			if chr.inverted {
				p.failAt(false, start.position, chr.val)
				return nil, false
			}
			p.read()
			p.failAt(true, start.position, chr.val)
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of ranges
	for i := 0; i < len(chr.ranges); i += 2 {
		if cur >= chr.ranges[i] && cur <= chr.ranges[i+1] {
			if chr.inverted {
				p.failAt(false, start.position, chr.val)
				return nil, false
			}
			p.read()
			p.failAt(true, start.position, chr.val)
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of Unicode classes
	for _, cl := range chr.classes {
		if unicode.Is(cl, cur) {
			if chr.inverted {
				p.failAt(false, start.position, chr.val)
				return nil, false
			}
			p.read()
			p.failAt(true, start.position, chr.val)
			return p.sliceFrom(start), true
		}
	}

	if chr.inverted {
		p.read()
		p.failAt(true, start.position, chr.val)
		return p.sliceFrom(start), true
	}
	p.failAt(false, start.position, chr.val)
	return nil, false
}

func (p *parser) parseChoiceExpr(ch *choiceExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseChoiceExpr"))
	}

	for _, alt := range ch.alternatives {
		p.pushV()
		val, ok := p.parseExpr(alt)
		p.popV()
		if ok {
			return val, ok
		}
	}
	return nil, false
}

func (p *parser) parseLabeledExpr(lab *labeledExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLabeledExpr"))
	}

	p.pushV()
	val, ok := p.parseExpr(lab.expr)
	p.popV()
	if ok && lab.label != "" {
		m := p.vstack[len(p.vstack)-1]
		m[lab.label] = val
	}
	return val, ok
}

func (p *parser) parseLitMatcher(lit *litMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLitMatcher"))
	}

	ignoreCase := ""
	if lit.ignoreCase {
		ignoreCase = "i"
	}
	val := fmt.Sprintf("%q%s", lit.val, ignoreCase)
	start := p.pt
	for _, want := range lit.val {
		cur := p.pt.rn
		if lit.ignoreCase {
			cur = unicode.ToLower(cur)
		}
		if cur != want {
			p.failAt(false, start.position, val)
			p.restore(start)
			return nil, false
		}
		p.read()
	}
	p.failAt(true, start.position, val)
	return p.sliceFrom(start), true
}

func (p *parser) parseNotCodeExpr(not *notCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotCodeExpr"))
	}

	ok, err := not.run(p)
	if err != nil {
		p.addErr(err)
	}
	return nil, !ok
}

func (p *parser) parseNotExpr(not *notExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotExpr"))
	}

	pt := p.pt
	p.pushV()
	p.maxFailInvertExpected = !p.maxFailInvertExpected
	_, ok := p.parseExpr(not.expr)
	p.maxFailInvertExpected = !p.maxFailInvertExpected
	p.popV()
	p.restore(pt)
	return nil, !ok
}

func (p *parser) parseOneOrMoreExpr(expr *oneOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseOneOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			if len(vals) == 0 {
				// did not match once, no match
				return nil, false
			}
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseRuleRefExpr(ref *ruleRefExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRuleRefExpr " + ref.name))
	}

	if ref.name == "" {
		panic(fmt.Sprintf("%s: invalid rule: missing name", ref.pos))
	}

	rule := p.rules[ref.name]
	if rule == nil {
		p.addErr(fmt.Errorf("undefined rule: %s", ref.name))
		return nil, false
	}
	return p.parseRule(rule)
}

func (p *parser) parseSeqExpr(seq *seqExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseSeqExpr"))
	}

	vals := make([]interface{}, 0, len(seq.exprs))

	pt := p.pt
	for _, expr := range seq.exprs {
		val, ok := p.parseExpr(expr)
		if !ok {
			p.restore(pt)
			return nil, false
		}
		vals = append(vals, val)
	}
	return vals, true
}

func (p *parser) parseZeroOrMoreExpr(expr *zeroOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseZeroOrOneExpr(expr *zeroOrOneExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrOneExpr"))
	}

	p.pushV()
	val, _ := p.parseExpr(expr.expr)
	p.popV()
	// whether it matched or not, consider it a match
	return val, true
}
