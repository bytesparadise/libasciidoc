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
							label: "frontmatter",
							expr: &zeroOrOneExpr{
								pos: position{line: 12, col: 26, offset: 376},
								expr: &ruleRefExpr{
									pos:  position{line: 12, col: 26, offset: 376},
									name: "FrontMatter",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 12, col: 40, offset: 390},
							label: "blocks",
							expr: &zeroOrMoreExpr{
								pos: position{line: 12, col: 48, offset: 398},
								expr: &ruleRefExpr{
									pos:  position{line: 12, col: 48, offset: 398},
									name: "DocumentBlock",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 12, col: 64, offset: 414},
							name: "EOF",
						},
					},
				},
			},
		},
		{
			name: "DocumentBlock",
			pos:  position{line: 19, col: 1, offset: 600},
			expr: &actionExpr{
				pos: position{line: 19, col: 18, offset: 617},
				run: (*parser).callonDocumentBlock1,
				expr: &labeledExpr{
					pos:   position{line: 19, col: 18, offset: 617},
					label: "content",
					expr: &choiceExpr{
						pos: position{line: 19, col: 27, offset: 626},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 19, col: 27, offset: 626},
								name: "Section",
							},
							&ruleRefExpr{
								pos:  position{line: 19, col: 37, offset: 636},
								name: "StandaloneBlock",
							},
						},
					},
				},
			},
		},
		{
			name: "StandaloneBlock",
			pos:  position{line: 23, col: 1, offset: 701},
			expr: &choiceExpr{
				pos: position{line: 23, col: 20, offset: 720},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 23, col: 20, offset: 720},
						name: "DocumentAttributeDeclaration",
					},
					&ruleRefExpr{
						pos:  position{line: 23, col: 51, offset: 751},
						name: "DocumentAttributeReset",
					},
					&ruleRefExpr{
						pos:  position{line: 23, col: 76, offset: 776},
						name: "List",
					},
					&ruleRefExpr{
						pos:  position{line: 23, col: 83, offset: 783},
						name: "BlockImage",
					},
					&ruleRefExpr{
						pos:  position{line: 23, col: 96, offset: 796},
						name: "DelimitedBlock",
					},
					&ruleRefExpr{
						pos:  position{line: 23, col: 113, offset: 813},
						name: "Paragraph",
					},
					&ruleRefExpr{
						pos:  position{line: 23, col: 125, offset: 825},
						name: "ElementAttribute",
					},
					&ruleRefExpr{
						pos:  position{line: 23, col: 144, offset: 844},
						name: "BlankLine",
					},
				},
			},
		},
		{
			name: "FrontMatter",
			pos:  position{line: 28, col: 1, offset: 1007},
			expr: &ruleRefExpr{
				pos:  position{line: 28, col: 16, offset: 1022},
				name: "YamlFrontMatter",
			},
		},
		{
			name: "FrontMatter",
			pos:  position{line: 30, col: 1, offset: 1040},
			expr: &actionExpr{
				pos: position{line: 30, col: 16, offset: 1055},
				run: (*parser).callonFrontMatter1,
				expr: &seqExpr{
					pos: position{line: 30, col: 16, offset: 1055},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 30, col: 16, offset: 1055},
							name: "YamlFrontMatterToken",
						},
						&labeledExpr{
							pos:   position{line: 30, col: 37, offset: 1076},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 30, col: 45, offset: 1084},
								expr: &seqExpr{
									pos: position{line: 30, col: 46, offset: 1085},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 30, col: 46, offset: 1085},
											expr: &ruleRefExpr{
												pos:  position{line: 30, col: 47, offset: 1086},
												name: "YamlFrontMatterToken",
											},
										},
										&anyMatcher{
											line: 30, col: 68, offset: 1107,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 30, col: 72, offset: 1111},
							name: "YamlFrontMatterToken",
						},
					},
				},
			},
		},
		{
			name: "YamlFrontMatterToken",
			pos:  position{line: 34, col: 1, offset: 1198},
			expr: &seqExpr{
				pos: position{line: 34, col: 26, offset: 1223},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 34, col: 26, offset: 1223},
						val:        "---",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 34, col: 32, offset: 1229},
						name: "EOL",
					},
				},
			},
		},
		{
			name: "Section",
			pos:  position{line: 39, col: 1, offset: 1338},
			expr: &choiceExpr{
				pos: position{line: 39, col: 12, offset: 1349},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 39, col: 12, offset: 1349},
						name: "Section1",
					},
					&ruleRefExpr{
						pos:  position{line: 39, col: 23, offset: 1360},
						name: "Section2",
					},
					&ruleRefExpr{
						pos:  position{line: 39, col: 34, offset: 1371},
						name: "Section3",
					},
					&ruleRefExpr{
						pos:  position{line: 39, col: 45, offset: 1382},
						name: "Section4",
					},
					&ruleRefExpr{
						pos:  position{line: 39, col: 56, offset: 1393},
						name: "Section5",
					},
					&ruleRefExpr{
						pos:  position{line: 39, col: 67, offset: 1404},
						name: "Section6",
					},
				},
			},
		},
		{
			name: "Section1",
			pos:  position{line: 41, col: 1, offset: 1414},
			expr: &actionExpr{
				pos: position{line: 41, col: 13, offset: 1426},
				run: (*parser).callonSection11,
				expr: &seqExpr{
					pos: position{line: 41, col: 13, offset: 1426},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 41, col: 13, offset: 1426},
							label: "heading",
							expr: &ruleRefExpr{
								pos:  position{line: 41, col: 22, offset: 1435},
								name: "Heading1",
							},
						},
						&labeledExpr{
							pos:   position{line: 41, col: 32, offset: 1445},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 41, col: 42, offset: 1455},
								expr: &ruleRefExpr{
									pos:  position{line: 41, col: 42, offset: 1455},
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
			pos:  position{line: 45, col: 1, offset: 1556},
			expr: &actionExpr{
				pos: position{line: 45, col: 18, offset: 1573},
				run: (*parser).callonSection1Block1,
				expr: &labeledExpr{
					pos:   position{line: 45, col: 18, offset: 1573},
					label: "content",
					expr: &choiceExpr{
						pos: position{line: 45, col: 27, offset: 1582},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 45, col: 27, offset: 1582},
								name: "Section2",
							},
							&ruleRefExpr{
								pos:  position{line: 45, col: 38, offset: 1593},
								name: "Section3",
							},
							&ruleRefExpr{
								pos:  position{line: 45, col: 49, offset: 1604},
								name: "Section4",
							},
							&ruleRefExpr{
								pos:  position{line: 45, col: 60, offset: 1615},
								name: "Section5",
							},
							&ruleRefExpr{
								pos:  position{line: 45, col: 71, offset: 1626},
								name: "Section6",
							},
							&ruleRefExpr{
								pos:  position{line: 45, col: 82, offset: 1637},
								name: "StandaloneBlock",
							},
						},
					},
				},
			},
		},
		{
			name: "Section2",
			pos:  position{line: 49, col: 1, offset: 1702},
			expr: &actionExpr{
				pos: position{line: 49, col: 13, offset: 1714},
				run: (*parser).callonSection21,
				expr: &seqExpr{
					pos: position{line: 49, col: 13, offset: 1714},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 49, col: 13, offset: 1714},
							label: "heading",
							expr: &ruleRefExpr{
								pos:  position{line: 49, col: 22, offset: 1723},
								name: "Heading2",
							},
						},
						&labeledExpr{
							pos:   position{line: 49, col: 32, offset: 1733},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 49, col: 42, offset: 1743},
								expr: &ruleRefExpr{
									pos:  position{line: 49, col: 42, offset: 1743},
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
			pos:  position{line: 53, col: 1, offset: 1844},
			expr: &actionExpr{
				pos: position{line: 53, col: 18, offset: 1861},
				run: (*parser).callonSection2Block1,
				expr: &seqExpr{
					pos: position{line: 53, col: 18, offset: 1861},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 53, col: 18, offset: 1861},
							expr: &ruleRefExpr{
								pos:  position{line: 53, col: 19, offset: 1862},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 53, col: 28, offset: 1871},
							expr: &ruleRefExpr{
								pos:  position{line: 53, col: 29, offset: 1872},
								name: "Section2",
							},
						},
						&labeledExpr{
							pos:   position{line: 53, col: 38, offset: 1881},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 53, col: 47, offset: 1890},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 53, col: 47, offset: 1890},
										name: "Section3",
									},
									&ruleRefExpr{
										pos:  position{line: 53, col: 58, offset: 1901},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 53, col: 69, offset: 1912},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 53, col: 80, offset: 1923},
										name: "Section6",
									},
									&ruleRefExpr{
										pos:  position{line: 53, col: 91, offset: 1934},
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
			pos:  position{line: 57, col: 1, offset: 1999},
			expr: &actionExpr{
				pos: position{line: 57, col: 13, offset: 2011},
				run: (*parser).callonSection31,
				expr: &seqExpr{
					pos: position{line: 57, col: 13, offset: 2011},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 57, col: 13, offset: 2011},
							label: "heading",
							expr: &ruleRefExpr{
								pos:  position{line: 57, col: 22, offset: 2020},
								name: "Heading3",
							},
						},
						&labeledExpr{
							pos:   position{line: 57, col: 32, offset: 2030},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 57, col: 42, offset: 2040},
								expr: &ruleRefExpr{
									pos:  position{line: 57, col: 42, offset: 2040},
									name: "Section3Block",
								},
							},
						},
						&andExpr{
							pos: position{line: 57, col: 58, offset: 2056},
							expr: &zeroOrMoreExpr{
								pos: position{line: 57, col: 59, offset: 2057},
								expr: &ruleRefExpr{
									pos:  position{line: 57, col: 60, offset: 2058},
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
			pos:  position{line: 61, col: 1, offset: 2154},
			expr: &actionExpr{
				pos: position{line: 61, col: 18, offset: 2171},
				run: (*parser).callonSection3Block1,
				expr: &seqExpr{
					pos: position{line: 61, col: 18, offset: 2171},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 61, col: 18, offset: 2171},
							expr: &ruleRefExpr{
								pos:  position{line: 61, col: 19, offset: 2172},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 61, col: 28, offset: 2181},
							expr: &ruleRefExpr{
								pos:  position{line: 61, col: 29, offset: 2182},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 61, col: 38, offset: 2191},
							expr: &ruleRefExpr{
								pos:  position{line: 61, col: 39, offset: 2192},
								name: "Section3",
							},
						},
						&labeledExpr{
							pos:   position{line: 61, col: 48, offset: 2201},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 61, col: 57, offset: 2210},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 61, col: 57, offset: 2210},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 61, col: 68, offset: 2221},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 61, col: 79, offset: 2232},
										name: "Section6",
									},
									&ruleRefExpr{
										pos:  position{line: 61, col: 90, offset: 2243},
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
			pos:  position{line: 65, col: 1, offset: 2308},
			expr: &actionExpr{
				pos: position{line: 65, col: 13, offset: 2320},
				run: (*parser).callonSection41,
				expr: &seqExpr{
					pos: position{line: 65, col: 13, offset: 2320},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 65, col: 13, offset: 2320},
							label: "heading",
							expr: &ruleRefExpr{
								pos:  position{line: 65, col: 22, offset: 2329},
								name: "Heading4",
							},
						},
						&labeledExpr{
							pos:   position{line: 65, col: 32, offset: 2339},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 65, col: 42, offset: 2349},
								expr: &ruleRefExpr{
									pos:  position{line: 65, col: 42, offset: 2349},
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
			pos:  position{line: 69, col: 1, offset: 2450},
			expr: &actionExpr{
				pos: position{line: 69, col: 18, offset: 2467},
				run: (*parser).callonSection4Block1,
				expr: &seqExpr{
					pos: position{line: 69, col: 18, offset: 2467},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 69, col: 18, offset: 2467},
							expr: &ruleRefExpr{
								pos:  position{line: 69, col: 19, offset: 2468},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 69, col: 28, offset: 2477},
							expr: &ruleRefExpr{
								pos:  position{line: 69, col: 29, offset: 2478},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 69, col: 38, offset: 2487},
							expr: &ruleRefExpr{
								pos:  position{line: 69, col: 39, offset: 2488},
								name: "Section3",
							},
						},
						&notExpr{
							pos: position{line: 69, col: 48, offset: 2497},
							expr: &ruleRefExpr{
								pos:  position{line: 69, col: 49, offset: 2498},
								name: "Section4",
							},
						},
						&labeledExpr{
							pos:   position{line: 69, col: 58, offset: 2507},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 69, col: 67, offset: 2516},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 69, col: 67, offset: 2516},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 69, col: 78, offset: 2527},
										name: "Section6",
									},
									&ruleRefExpr{
										pos:  position{line: 69, col: 89, offset: 2538},
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
			pos:  position{line: 73, col: 1, offset: 2603},
			expr: &actionExpr{
				pos: position{line: 73, col: 13, offset: 2615},
				run: (*parser).callonSection51,
				expr: &seqExpr{
					pos: position{line: 73, col: 13, offset: 2615},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 73, col: 13, offset: 2615},
							label: "heading",
							expr: &ruleRefExpr{
								pos:  position{line: 73, col: 22, offset: 2624},
								name: "Heading5",
							},
						},
						&labeledExpr{
							pos:   position{line: 73, col: 32, offset: 2634},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 73, col: 42, offset: 2644},
								expr: &ruleRefExpr{
									pos:  position{line: 73, col: 42, offset: 2644},
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
			pos:  position{line: 77, col: 1, offset: 2745},
			expr: &actionExpr{
				pos: position{line: 77, col: 18, offset: 2762},
				run: (*parser).callonSection5Block1,
				expr: &seqExpr{
					pos: position{line: 77, col: 18, offset: 2762},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 77, col: 18, offset: 2762},
							expr: &ruleRefExpr{
								pos:  position{line: 77, col: 19, offset: 2763},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 77, col: 28, offset: 2772},
							expr: &ruleRefExpr{
								pos:  position{line: 77, col: 29, offset: 2773},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 77, col: 38, offset: 2782},
							expr: &ruleRefExpr{
								pos:  position{line: 77, col: 39, offset: 2783},
								name: "Section3",
							},
						},
						&notExpr{
							pos: position{line: 77, col: 48, offset: 2792},
							expr: &ruleRefExpr{
								pos:  position{line: 77, col: 49, offset: 2793},
								name: "Section4",
							},
						},
						&notExpr{
							pos: position{line: 77, col: 58, offset: 2802},
							expr: &ruleRefExpr{
								pos:  position{line: 77, col: 59, offset: 2803},
								name: "Section5",
							},
						},
						&labeledExpr{
							pos:   position{line: 77, col: 68, offset: 2812},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 77, col: 77, offset: 2821},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 77, col: 77, offset: 2821},
										name: "Section6",
									},
									&ruleRefExpr{
										pos:  position{line: 77, col: 88, offset: 2832},
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
			pos:  position{line: 81, col: 1, offset: 2897},
			expr: &actionExpr{
				pos: position{line: 81, col: 13, offset: 2909},
				run: (*parser).callonSection61,
				expr: &seqExpr{
					pos: position{line: 81, col: 13, offset: 2909},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 81, col: 13, offset: 2909},
							label: "heading",
							expr: &ruleRefExpr{
								pos:  position{line: 81, col: 22, offset: 2918},
								name: "Heading6",
							},
						},
						&labeledExpr{
							pos:   position{line: 81, col: 32, offset: 2928},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 81, col: 42, offset: 2938},
								expr: &ruleRefExpr{
									pos:  position{line: 81, col: 42, offset: 2938},
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
			pos:  position{line: 85, col: 1, offset: 3039},
			expr: &actionExpr{
				pos: position{line: 85, col: 18, offset: 3056},
				run: (*parser).callonSection6Block1,
				expr: &seqExpr{
					pos: position{line: 85, col: 18, offset: 3056},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 85, col: 18, offset: 3056},
							expr: &ruleRefExpr{
								pos:  position{line: 85, col: 19, offset: 3057},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 85, col: 28, offset: 3066},
							expr: &ruleRefExpr{
								pos:  position{line: 85, col: 29, offset: 3067},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 85, col: 38, offset: 3076},
							expr: &ruleRefExpr{
								pos:  position{line: 85, col: 39, offset: 3077},
								name: "Section3",
							},
						},
						&notExpr{
							pos: position{line: 85, col: 48, offset: 3086},
							expr: &ruleRefExpr{
								pos:  position{line: 85, col: 49, offset: 3087},
								name: "Section4",
							},
						},
						&notExpr{
							pos: position{line: 85, col: 58, offset: 3096},
							expr: &ruleRefExpr{
								pos:  position{line: 85, col: 59, offset: 3097},
								name: "Section5",
							},
						},
						&notExpr{
							pos: position{line: 85, col: 68, offset: 3106},
							expr: &ruleRefExpr{
								pos:  position{line: 85, col: 69, offset: 3107},
								name: "Section6",
							},
						},
						&labeledExpr{
							pos:   position{line: 85, col: 78, offset: 3116},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 85, col: 87, offset: 3125},
								name: "StandaloneBlock",
							},
						},
					},
				},
			},
		},
		{
			name: "Heading",
			pos:  position{line: 92, col: 1, offset: 3294},
			expr: &choiceExpr{
				pos: position{line: 92, col: 12, offset: 3305},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 92, col: 12, offset: 3305},
						name: "Heading1",
					},
					&ruleRefExpr{
						pos:  position{line: 92, col: 23, offset: 3316},
						name: "Heading2",
					},
					&ruleRefExpr{
						pos:  position{line: 92, col: 34, offset: 3327},
						name: "Heading3",
					},
					&ruleRefExpr{
						pos:  position{line: 92, col: 45, offset: 3338},
						name: "Heading4",
					},
					&ruleRefExpr{
						pos:  position{line: 92, col: 56, offset: 3349},
						name: "Heading5",
					},
					&ruleRefExpr{
						pos:  position{line: 92, col: 67, offset: 3360},
						name: "Heading6",
					},
				},
			},
		},
		{
			name: "Heading1",
			pos:  position{line: 94, col: 1, offset: 3370},
			expr: &actionExpr{
				pos: position{line: 94, col: 13, offset: 3382},
				run: (*parser).callonHeading11,
				expr: &seqExpr{
					pos: position{line: 94, col: 13, offset: 3382},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 94, col: 13, offset: 3382},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 94, col: 24, offset: 3393},
								expr: &ruleRefExpr{
									pos:  position{line: 94, col: 25, offset: 3394},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 94, col: 44, offset: 3413},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 94, col: 51, offset: 3420},
								val:        "=",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 94, col: 56, offset: 3425},
							expr: &ruleRefExpr{
								pos:  position{line: 94, col: 56, offset: 3425},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 94, col: 60, offset: 3429},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 94, col: 68, offset: 3437},
								name: "InlineContent",
							},
						},
						&choiceExpr{
							pos: position{line: 94, col: 83, offset: 3452},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 94, col: 83, offset: 3452},
									expr: &ruleRefExpr{
										pos:  position{line: 94, col: 83, offset: 3452},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 94, col: 96, offset: 3465},
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
			pos:  position{line: 98, col: 1, offset: 3672},
			expr: &actionExpr{
				pos: position{line: 98, col: 13, offset: 3684},
				run: (*parser).callonHeading21,
				expr: &seqExpr{
					pos: position{line: 98, col: 13, offset: 3684},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 98, col: 13, offset: 3684},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 98, col: 24, offset: 3695},
								expr: &ruleRefExpr{
									pos:  position{line: 98, col: 25, offset: 3696},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 98, col: 44, offset: 3715},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 98, col: 51, offset: 3722},
								val:        "==",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 98, col: 57, offset: 3728},
							expr: &ruleRefExpr{
								pos:  position{line: 98, col: 57, offset: 3728},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 98, col: 61, offset: 3732},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 98, col: 69, offset: 3740},
								name: "InlineContent",
							},
						},
						&choiceExpr{
							pos: position{line: 98, col: 84, offset: 3755},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 98, col: 84, offset: 3755},
									expr: &ruleRefExpr{
										pos:  position{line: 98, col: 84, offset: 3755},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 97, offset: 3768},
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
			pos:  position{line: 102, col: 1, offset: 3870},
			expr: &actionExpr{
				pos: position{line: 102, col: 13, offset: 3882},
				run: (*parser).callonHeading31,
				expr: &seqExpr{
					pos: position{line: 102, col: 13, offset: 3882},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 102, col: 13, offset: 3882},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 102, col: 24, offset: 3893},
								expr: &ruleRefExpr{
									pos:  position{line: 102, col: 25, offset: 3894},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 102, col: 44, offset: 3913},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 102, col: 51, offset: 3920},
								val:        "===",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 102, col: 58, offset: 3927},
							expr: &ruleRefExpr{
								pos:  position{line: 102, col: 58, offset: 3927},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 102, col: 62, offset: 3931},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 102, col: 70, offset: 3939},
								name: "InlineContent",
							},
						},
						&choiceExpr{
							pos: position{line: 102, col: 85, offset: 3954},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 102, col: 85, offset: 3954},
									expr: &ruleRefExpr{
										pos:  position{line: 102, col: 85, offset: 3954},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 102, col: 98, offset: 3967},
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
			pos:  position{line: 106, col: 1, offset: 4069},
			expr: &actionExpr{
				pos: position{line: 106, col: 13, offset: 4081},
				run: (*parser).callonHeading41,
				expr: &seqExpr{
					pos: position{line: 106, col: 13, offset: 4081},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 106, col: 13, offset: 4081},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 106, col: 24, offset: 4092},
								expr: &ruleRefExpr{
									pos:  position{line: 106, col: 25, offset: 4093},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 106, col: 44, offset: 4112},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 106, col: 51, offset: 4119},
								val:        "====",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 106, col: 59, offset: 4127},
							expr: &ruleRefExpr{
								pos:  position{line: 106, col: 59, offset: 4127},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 106, col: 63, offset: 4131},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 106, col: 71, offset: 4139},
								name: "InlineContent",
							},
						},
						&choiceExpr{
							pos: position{line: 106, col: 86, offset: 4154},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 106, col: 86, offset: 4154},
									expr: &ruleRefExpr{
										pos:  position{line: 106, col: 86, offset: 4154},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 106, col: 99, offset: 4167},
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
			pos:  position{line: 110, col: 1, offset: 4269},
			expr: &actionExpr{
				pos: position{line: 110, col: 13, offset: 4281},
				run: (*parser).callonHeading51,
				expr: &seqExpr{
					pos: position{line: 110, col: 13, offset: 4281},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 110, col: 13, offset: 4281},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 110, col: 24, offset: 4292},
								expr: &ruleRefExpr{
									pos:  position{line: 110, col: 25, offset: 4293},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 110, col: 44, offset: 4312},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 110, col: 51, offset: 4319},
								val:        "=====",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 110, col: 60, offset: 4328},
							expr: &ruleRefExpr{
								pos:  position{line: 110, col: 60, offset: 4328},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 110, col: 64, offset: 4332},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 110, col: 72, offset: 4340},
								name: "InlineContent",
							},
						},
						&choiceExpr{
							pos: position{line: 110, col: 87, offset: 4355},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 110, col: 87, offset: 4355},
									expr: &ruleRefExpr{
										pos:  position{line: 110, col: 87, offset: 4355},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 110, col: 100, offset: 4368},
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
			pos:  position{line: 114, col: 1, offset: 4470},
			expr: &actionExpr{
				pos: position{line: 114, col: 13, offset: 4482},
				run: (*parser).callonHeading61,
				expr: &seqExpr{
					pos: position{line: 114, col: 13, offset: 4482},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 114, col: 13, offset: 4482},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 114, col: 24, offset: 4493},
								expr: &ruleRefExpr{
									pos:  position{line: 114, col: 25, offset: 4494},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 114, col: 44, offset: 4513},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 114, col: 51, offset: 4520},
								val:        "======",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 114, col: 61, offset: 4530},
							expr: &ruleRefExpr{
								pos:  position{line: 114, col: 61, offset: 4530},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 114, col: 65, offset: 4534},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 114, col: 73, offset: 4542},
								name: "InlineContent",
							},
						},
						&choiceExpr{
							pos: position{line: 114, col: 88, offset: 4557},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 114, col: 88, offset: 4557},
									expr: &ruleRefExpr{
										pos:  position{line: 114, col: 88, offset: 4557},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 114, col: 101, offset: 4570},
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
			pos:  position{line: 122, col: 1, offset: 4788},
			expr: &choiceExpr{
				pos: position{line: 122, col: 33, offset: 4820},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 122, col: 33, offset: 4820},
						name: "DocumentAttributeDeclarationWithNameOnly",
					},
					&ruleRefExpr{
						pos:  position{line: 122, col: 76, offset: 4863},
						name: "DocumentAttributeDeclarationWithNameAndValue",
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclarationWithNameOnly",
			pos:  position{line: 124, col: 1, offset: 4910},
			expr: &actionExpr{
				pos: position{line: 124, col: 45, offset: 4954},
				run: (*parser).callonDocumentAttributeDeclarationWithNameOnly1,
				expr: &seqExpr{
					pos: position{line: 124, col: 45, offset: 4954},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 124, col: 45, offset: 4954},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 124, col: 49, offset: 4958},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 124, col: 55, offset: 4964},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 124, col: 70, offset: 4979},
							val:        ":",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 124, col: 74, offset: 4983},
							expr: &ruleRefExpr{
								pos:  position{line: 124, col: 74, offset: 4983},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 124, col: 78, offset: 4987},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclarationWithNameAndValue",
			pos:  position{line: 128, col: 1, offset: 5072},
			expr: &actionExpr{
				pos: position{line: 128, col: 49, offset: 5120},
				run: (*parser).callonDocumentAttributeDeclarationWithNameAndValue1,
				expr: &seqExpr{
					pos: position{line: 128, col: 49, offset: 5120},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 128, col: 49, offset: 5120},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 128, col: 53, offset: 5124},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 128, col: 59, offset: 5130},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 128, col: 74, offset: 5145},
							val:        ":",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 128, col: 78, offset: 5149},
							expr: &ruleRefExpr{
								pos:  position{line: 128, col: 78, offset: 5149},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 128, col: 82, offset: 5153},
							label: "value",
							expr: &zeroOrMoreExpr{
								pos: position{line: 128, col: 88, offset: 5159},
								expr: &seqExpr{
									pos: position{line: 128, col: 89, offset: 5160},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 128, col: 89, offset: 5160},
											expr: &ruleRefExpr{
												pos:  position{line: 128, col: 90, offset: 5161},
												name: "NEWLINE",
											},
										},
										&anyMatcher{
											line: 128, col: 98, offset: 5169,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 128, col: 102, offset: 5173},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeReset",
			pos:  position{line: 132, col: 1, offset: 5276},
			expr: &choiceExpr{
				pos: position{line: 132, col: 27, offset: 5302},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 132, col: 27, offset: 5302},
						name: "DocumentAttributeResetWithHeadingBangSymbol",
					},
					&ruleRefExpr{
						pos:  position{line: 132, col: 73, offset: 5348},
						name: "DocumentAttributeResetWithTrailingBangSymbol",
					},
				},
			},
		},
		{
			name: "DocumentAttributeResetWithHeadingBangSymbol",
			pos:  position{line: 134, col: 1, offset: 5394},
			expr: &actionExpr{
				pos: position{line: 134, col: 48, offset: 5441},
				run: (*parser).callonDocumentAttributeResetWithHeadingBangSymbol1,
				expr: &seqExpr{
					pos: position{line: 134, col: 48, offset: 5441},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 134, col: 48, offset: 5441},
							val:        ":!",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 134, col: 53, offset: 5446},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 134, col: 59, offset: 5452},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 134, col: 74, offset: 5467},
							val:        ":",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 134, col: 78, offset: 5471},
							expr: &ruleRefExpr{
								pos:  position{line: 134, col: 78, offset: 5471},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 134, col: 82, offset: 5475},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeResetWithTrailingBangSymbol",
			pos:  position{line: 138, col: 1, offset: 5549},
			expr: &actionExpr{
				pos: position{line: 138, col: 49, offset: 5597},
				run: (*parser).callonDocumentAttributeResetWithTrailingBangSymbol1,
				expr: &seqExpr{
					pos: position{line: 138, col: 49, offset: 5597},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 138, col: 49, offset: 5597},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 138, col: 53, offset: 5601},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 138, col: 59, offset: 5607},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 138, col: 74, offset: 5622},
							val:        "!:",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 138, col: 79, offset: 5627},
							expr: &ruleRefExpr{
								pos:  position{line: 138, col: 79, offset: 5627},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 138, col: 83, offset: 5631},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeSubstitution",
			pos:  position{line: 143, col: 1, offset: 5706},
			expr: &actionExpr{
				pos: position{line: 143, col: 34, offset: 5739},
				run: (*parser).callonDocumentAttributeSubstitution1,
				expr: &seqExpr{
					pos: position{line: 143, col: 34, offset: 5739},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 143, col: 34, offset: 5739},
							val:        "{",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 143, col: 38, offset: 5743},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 143, col: 44, offset: 5749},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 143, col: 59, offset: 5764},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "AttributeName",
			pos:  position{line: 150, col: 1, offset: 6018},
			expr: &seqExpr{
				pos: position{line: 150, col: 18, offset: 6035},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 150, col: 19, offset: 6036},
						alternatives: []interface{}{
							&charClassMatcher{
								pos:        position{line: 150, col: 19, offset: 6036},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
							&charClassMatcher{
								pos:        position{line: 150, col: 27, offset: 6044},
								val:        "[a-z]",
								ranges:     []rune{'a', 'z'},
								ignoreCase: false,
								inverted:   false,
							},
							&charClassMatcher{
								pos:        position{line: 150, col: 35, offset: 6052},
								val:        "[0-9]",
								ranges:     []rune{'0', '9'},
								ignoreCase: false,
								inverted:   false,
							},
							&litMatcher{
								pos:        position{line: 150, col: 43, offset: 6060},
								val:        "_",
								ignoreCase: false,
							},
						},
					},
					&zeroOrMoreExpr{
						pos: position{line: 150, col: 48, offset: 6065},
						expr: &choiceExpr{
							pos: position{line: 150, col: 49, offset: 6066},
							alternatives: []interface{}{
								&charClassMatcher{
									pos:        position{line: 150, col: 49, offset: 6066},
									val:        "[A-Z]",
									ranges:     []rune{'A', 'Z'},
									ignoreCase: false,
									inverted:   false,
								},
								&charClassMatcher{
									pos:        position{line: 150, col: 57, offset: 6074},
									val:        "[a-z]",
									ranges:     []rune{'a', 'z'},
									ignoreCase: false,
									inverted:   false,
								},
								&charClassMatcher{
									pos:        position{line: 150, col: 65, offset: 6082},
									val:        "[0-9]",
									ranges:     []rune{'0', '9'},
									ignoreCase: false,
									inverted:   false,
								},
								&litMatcher{
									pos:        position{line: 150, col: 73, offset: 6090},
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
			pos:  position{line: 155, col: 1, offset: 6203},
			expr: &actionExpr{
				pos: position{line: 155, col: 9, offset: 6211},
				run: (*parser).callonList1,
				expr: &seqExpr{
					pos: position{line: 155, col: 9, offset: 6211},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 155, col: 9, offset: 6211},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 155, col: 20, offset: 6222},
								expr: &ruleRefExpr{
									pos:  position{line: 155, col: 21, offset: 6223},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 157, col: 5, offset: 6315},
							label: "elements",
							expr: &oneOrMoreExpr{
								pos: position{line: 157, col: 14, offset: 6324},
								expr: &seqExpr{
									pos: position{line: 157, col: 15, offset: 6325},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 157, col: 15, offset: 6325},
											name: "ListItem",
										},
										&zeroOrOneExpr{
											pos: position{line: 157, col: 24, offset: 6334},
											expr: &ruleRefExpr{
												pos:  position{line: 157, col: 24, offset: 6334},
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
			pos:  position{line: 161, col: 1, offset: 6431},
			expr: &actionExpr{
				pos: position{line: 161, col: 13, offset: 6443},
				run: (*parser).callonListItem1,
				expr: &seqExpr{
					pos: position{line: 161, col: 13, offset: 6443},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 161, col: 13, offset: 6443},
							expr: &ruleRefExpr{
								pos:  position{line: 161, col: 13, offset: 6443},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 161, col: 17, offset: 6447},
							label: "level",
							expr: &choiceExpr{
								pos: position{line: 161, col: 24, offset: 6454},
								alternatives: []interface{}{
									&oneOrMoreExpr{
										pos: position{line: 161, col: 24, offset: 6454},
										expr: &litMatcher{
											pos:        position{line: 161, col: 24, offset: 6454},
											val:        "*",
											ignoreCase: false,
										},
									},
									&litMatcher{
										pos:        position{line: 161, col: 31, offset: 6461},
										val:        "-",
										ignoreCase: false,
									},
								},
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 161, col: 36, offset: 6466},
							expr: &ruleRefExpr{
								pos:  position{line: 161, col: 36, offset: 6466},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 161, col: 40, offset: 6470},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 161, col: 49, offset: 6479},
								name: "ListItemContent",
							},
						},
					},
				},
			},
		},
		{
			name: "ListItemContent",
			pos:  position{line: 165, col: 1, offset: 6576},
			expr: &actionExpr{
				pos: position{line: 165, col: 20, offset: 6595},
				run: (*parser).callonListItemContent1,
				expr: &labeledExpr{
					pos:   position{line: 165, col: 20, offset: 6595},
					label: "lines",
					expr: &oneOrMoreExpr{
						pos: position{line: 165, col: 26, offset: 6601},
						expr: &seqExpr{
							pos: position{line: 165, col: 27, offset: 6602},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 165, col: 27, offset: 6602},
									expr: &seqExpr{
										pos: position{line: 165, col: 29, offset: 6604},
										exprs: []interface{}{
											&zeroOrMoreExpr{
												pos: position{line: 165, col: 29, offset: 6604},
												expr: &ruleRefExpr{
													pos:  position{line: 165, col: 29, offset: 6604},
													name: "WS",
												},
											},
											&choiceExpr{
												pos: position{line: 165, col: 34, offset: 6609},
												alternatives: []interface{}{
													&oneOrMoreExpr{
														pos: position{line: 165, col: 34, offset: 6609},
														expr: &litMatcher{
															pos:        position{line: 165, col: 34, offset: 6609},
															val:        "*",
															ignoreCase: false,
														},
													},
													&litMatcher{
														pos:        position{line: 165, col: 41, offset: 6616},
														val:        "-",
														ignoreCase: false,
													},
												},
											},
											&oneOrMoreExpr{
												pos: position{line: 165, col: 46, offset: 6621},
												expr: &ruleRefExpr{
													pos:  position{line: 165, col: 46, offset: 6621},
													name: "WS",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 165, col: 51, offset: 6626},
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
			pos:  position{line: 172, col: 1, offset: 6896},
			expr: &actionExpr{
				pos: position{line: 172, col: 14, offset: 6909},
				run: (*parser).callonParagraph1,
				expr: &seqExpr{
					pos: position{line: 172, col: 14, offset: 6909},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 172, col: 14, offset: 6909},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 172, col: 25, offset: 6920},
								expr: &ruleRefExpr{
									pos:  position{line: 172, col: 26, offset: 6921},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 172, col: 45, offset: 6940},
							label: "lines",
							expr: &oneOrMoreExpr{
								pos: position{line: 172, col: 51, offset: 6946},
								expr: &ruleRefExpr{
									pos:  position{line: 172, col: 52, offset: 6947},
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
			pos:  position{line: 178, col: 1, offset: 7255},
			expr: &actionExpr{
				pos: position{line: 178, col: 18, offset: 7272},
				run: (*parser).callonInlineContent1,
				expr: &seqExpr{
					pos: position{line: 178, col: 18, offset: 7272},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 178, col: 18, offset: 7272},
							label: "elements",
							expr: &oneOrMoreExpr{
								pos: position{line: 178, col: 27, offset: 7281},
								expr: &seqExpr{
									pos: position{line: 178, col: 28, offset: 7282},
									exprs: []interface{}{
										&zeroOrMoreExpr{
											pos: position{line: 178, col: 28, offset: 7282},
											expr: &ruleRefExpr{
												pos:  position{line: 178, col: 28, offset: 7282},
												name: "WS",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 178, col: 32, offset: 7286},
											name: "InlineElement",
										},
										&zeroOrMoreExpr{
											pos: position{line: 178, col: 46, offset: 7300},
											expr: &ruleRefExpr{
												pos:  position{line: 178, col: 46, offset: 7300},
												name: "WS",
											},
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 178, col: 52, offset: 7306},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "InlineElement",
			pos:  position{line: 182, col: 1, offset: 7384},
			expr: &choiceExpr{
				pos: position{line: 182, col: 18, offset: 7401},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 182, col: 18, offset: 7401},
						name: "InlineImage",
					},
					&ruleRefExpr{
						pos:  position{line: 182, col: 32, offset: 7415},
						name: "QuotedText",
					},
					&ruleRefExpr{
						pos:  position{line: 182, col: 45, offset: 7428},
						name: "ExternalLink",
					},
					&ruleRefExpr{
						pos:  position{line: 182, col: 60, offset: 7443},
						name: "DocumentAttributeSubstitution",
					},
					&ruleRefExpr{
						pos:  position{line: 182, col: 92, offset: 7475},
						name: "Word",
					},
				},
			},
		},
		{
			name: "QuotedText",
			pos:  position{line: 187, col: 1, offset: 7618},
			expr: &choiceExpr{
				pos: position{line: 187, col: 15, offset: 7632},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 187, col: 15, offset: 7632},
						name: "BoldText",
					},
					&ruleRefExpr{
						pos:  position{line: 187, col: 26, offset: 7643},
						name: "ItalicText",
					},
					&ruleRefExpr{
						pos:  position{line: 187, col: 39, offset: 7656},
						name: "MonospaceText",
					},
				},
			},
		},
		{
			name: "BoldText",
			pos:  position{line: 189, col: 1, offset: 7671},
			expr: &actionExpr{
				pos: position{line: 189, col: 13, offset: 7683},
				run: (*parser).callonBoldText1,
				expr: &seqExpr{
					pos: position{line: 189, col: 13, offset: 7683},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 189, col: 13, offset: 7683},
							val:        "*",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 189, col: 17, offset: 7687},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 189, col: 26, offset: 7696},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 189, col: 45, offset: 7715},
							val:        "*",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ItalicText",
			pos:  position{line: 193, col: 1, offset: 7792},
			expr: &actionExpr{
				pos: position{line: 193, col: 15, offset: 7806},
				run: (*parser).callonItalicText1,
				expr: &seqExpr{
					pos: position{line: 193, col: 15, offset: 7806},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 193, col: 15, offset: 7806},
							val:        "_",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 193, col: 19, offset: 7810},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 193, col: 28, offset: 7819},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 193, col: 47, offset: 7838},
							val:        "_",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "MonospaceText",
			pos:  position{line: 197, col: 1, offset: 7917},
			expr: &actionExpr{
				pos: position{line: 197, col: 18, offset: 7934},
				run: (*parser).callonMonospaceText1,
				expr: &seqExpr{
					pos: position{line: 197, col: 18, offset: 7934},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 197, col: 18, offset: 7934},
							val:        "`",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 197, col: 22, offset: 7938},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 197, col: 31, offset: 7947},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 197, col: 50, offset: 7966},
							val:        "`",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "QuotedTextContent",
			pos:  position{line: 201, col: 1, offset: 8048},
			expr: &seqExpr{
				pos: position{line: 201, col: 22, offset: 8069},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 201, col: 22, offset: 8069},
						name: "QuotedTextContentElement",
					},
					&zeroOrMoreExpr{
						pos: position{line: 201, col: 47, offset: 8094},
						expr: &seqExpr{
							pos: position{line: 201, col: 48, offset: 8095},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 201, col: 48, offset: 8095},
									expr: &ruleRefExpr{
										pos:  position{line: 201, col: 48, offset: 8095},
										name: "WS",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 201, col: 52, offset: 8099},
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
			pos:  position{line: 203, col: 1, offset: 8127},
			expr: &choiceExpr{
				pos: position{line: 203, col: 29, offset: 8155},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 203, col: 29, offset: 8155},
						name: "QuotedText",
					},
					&ruleRefExpr{
						pos:  position{line: 203, col: 42, offset: 8168},
						name: "QuotedTextContentWord",
					},
					&ruleRefExpr{
						pos:  position{line: 203, col: 66, offset: 8192},
						name: "InvalidQuotedTextContentWord",
					},
				},
			},
		},
		{
			name: "QuotedTextContentWord",
			pos:  position{line: 205, col: 1, offset: 8222},
			expr: &oneOrMoreExpr{
				pos: position{line: 205, col: 26, offset: 8247},
				expr: &seqExpr{
					pos: position{line: 205, col: 27, offset: 8248},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 205, col: 27, offset: 8248},
							expr: &ruleRefExpr{
								pos:  position{line: 205, col: 28, offset: 8249},
								name: "NEWLINE",
							},
						},
						&notExpr{
							pos: position{line: 205, col: 36, offset: 8257},
							expr: &ruleRefExpr{
								pos:  position{line: 205, col: 37, offset: 8258},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 205, col: 40, offset: 8261},
							expr: &litMatcher{
								pos:        position{line: 205, col: 41, offset: 8262},
								val:        "*",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 205, col: 45, offset: 8266},
							expr: &litMatcher{
								pos:        position{line: 205, col: 46, offset: 8267},
								val:        "_",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 205, col: 50, offset: 8271},
							expr: &litMatcher{
								pos:        position{line: 205, col: 51, offset: 8272},
								val:        "`",
								ignoreCase: false,
							},
						},
						&anyMatcher{
							line: 205, col: 55, offset: 8276,
						},
					},
				},
			},
		},
		{
			name: "InvalidQuotedTextContentWord",
			pos:  position{line: 206, col: 1, offset: 8318},
			expr: &oneOrMoreExpr{
				pos: position{line: 206, col: 33, offset: 8350},
				expr: &seqExpr{
					pos: position{line: 206, col: 34, offset: 8351},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 206, col: 34, offset: 8351},
							expr: &ruleRefExpr{
								pos:  position{line: 206, col: 35, offset: 8352},
								name: "NEWLINE",
							},
						},
						&notExpr{
							pos: position{line: 206, col: 43, offset: 8360},
							expr: &ruleRefExpr{
								pos:  position{line: 206, col: 44, offset: 8361},
								name: "WS",
							},
						},
						&anyMatcher{
							line: 206, col: 48, offset: 8365,
						},
					},
				},
			},
		},
		{
			name: "ExternalLink",
			pos:  position{line: 211, col: 1, offset: 8582},
			expr: &actionExpr{
				pos: position{line: 211, col: 17, offset: 8598},
				run: (*parser).callonExternalLink1,
				expr: &seqExpr{
					pos: position{line: 211, col: 17, offset: 8598},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 211, col: 17, offset: 8598},
							label: "url",
							expr: &seqExpr{
								pos: position{line: 211, col: 22, offset: 8603},
								exprs: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 211, col: 22, offset: 8603},
										name: "URL_SCHEME",
									},
									&ruleRefExpr{
										pos:  position{line: 211, col: 33, offset: 8614},
										name: "URL",
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 211, col: 38, offset: 8619},
							label: "text",
							expr: &zeroOrOneExpr{
								pos: position{line: 211, col: 43, offset: 8624},
								expr: &seqExpr{
									pos: position{line: 211, col: 44, offset: 8625},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 211, col: 44, offset: 8625},
											val:        "[",
											ignoreCase: false,
										},
										&zeroOrMoreExpr{
											pos: position{line: 211, col: 48, offset: 8629},
											expr: &ruleRefExpr{
												pos:  position{line: 211, col: 49, offset: 8630},
												name: "URL_TEXT",
											},
										},
										&litMatcher{
											pos:        position{line: 211, col: 60, offset: 8641},
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
			pos:  position{line: 221, col: 1, offset: 8920},
			expr: &actionExpr{
				pos: position{line: 221, col: 15, offset: 8934},
				run: (*parser).callonBlockImage1,
				expr: &seqExpr{
					pos: position{line: 221, col: 15, offset: 8934},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 221, col: 15, offset: 8934},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 221, col: 26, offset: 8945},
								expr: &ruleRefExpr{
									pos:  position{line: 221, col: 27, offset: 8946},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 221, col: 46, offset: 8965},
							label: "image",
							expr: &ruleRefExpr{
								pos:  position{line: 221, col: 52, offset: 8971},
								name: "BlockImageMacro",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 221, col: 69, offset: 8988},
							expr: &ruleRefExpr{
								pos:  position{line: 221, col: 69, offset: 8988},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 221, col: 73, offset: 8992},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "BlockImageMacro",
			pos:  position{line: 226, col: 1, offset: 9161},
			expr: &actionExpr{
				pos: position{line: 226, col: 20, offset: 9180},
				run: (*parser).callonBlockImageMacro1,
				expr: &seqExpr{
					pos: position{line: 226, col: 20, offset: 9180},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 226, col: 20, offset: 9180},
							val:        "image::",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 226, col: 30, offset: 9190},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 226, col: 36, offset: 9196},
								name: "URL",
							},
						},
						&litMatcher{
							pos:        position{line: 226, col: 41, offset: 9201},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 226, col: 45, offset: 9205},
							label: "attributes",
							expr: &zeroOrOneExpr{
								pos: position{line: 226, col: 57, offset: 9217},
								expr: &ruleRefExpr{
									pos:  position{line: 226, col: 57, offset: 9217},
									name: "URL_TEXT",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 226, col: 68, offset: 9228},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "InlineImage",
			pos:  position{line: 230, col: 1, offset: 9303},
			expr: &actionExpr{
				pos: position{line: 230, col: 16, offset: 9318},
				run: (*parser).callonInlineImage1,
				expr: &labeledExpr{
					pos:   position{line: 230, col: 16, offset: 9318},
					label: "image",
					expr: &ruleRefExpr{
						pos:  position{line: 230, col: 22, offset: 9324},
						name: "InlineImageMacro",
					},
				},
			},
		},
		{
			name: "InlineImageMacro",
			pos:  position{line: 235, col: 1, offset: 9479},
			expr: &actionExpr{
				pos: position{line: 235, col: 21, offset: 9499},
				run: (*parser).callonInlineImageMacro1,
				expr: &seqExpr{
					pos: position{line: 235, col: 21, offset: 9499},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 235, col: 21, offset: 9499},
							val:        "image:",
							ignoreCase: false,
						},
						&notExpr{
							pos: position{line: 235, col: 30, offset: 9508},
							expr: &litMatcher{
								pos:        position{line: 235, col: 31, offset: 9509},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 235, col: 35, offset: 9513},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 235, col: 41, offset: 9519},
								name: "URL",
							},
						},
						&litMatcher{
							pos:        position{line: 235, col: 46, offset: 9524},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 235, col: 50, offset: 9528},
							label: "attributes",
							expr: &zeroOrOneExpr{
								pos: position{line: 235, col: 62, offset: 9540},
								expr: &ruleRefExpr{
									pos:  position{line: 235, col: 62, offset: 9540},
									name: "URL_TEXT",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 235, col: 73, offset: 9551},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "DelimitedBlock",
			pos:  position{line: 243, col: 1, offset: 9739},
			expr: &ruleRefExpr{
				pos:  position{line: 243, col: 19, offset: 9757},
				name: "SourceBlock",
			},
		},
		{
			name: "SourceBlock",
			pos:  position{line: 245, col: 1, offset: 9770},
			expr: &actionExpr{
				pos: position{line: 245, col: 16, offset: 9785},
				run: (*parser).callonSourceBlock1,
				expr: &seqExpr{
					pos: position{line: 245, col: 16, offset: 9785},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 245, col: 16, offset: 9785},
							name: "SourceBlockDelimiter",
						},
						&ruleRefExpr{
							pos:  position{line: 245, col: 37, offset: 9806},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 245, col: 45, offset: 9814},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 245, col: 53, offset: 9822},
								expr: &ruleRefExpr{
									pos:  position{line: 245, col: 54, offset: 9823},
									name: "SourceBlockLine",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 245, col: 73, offset: 9842},
							name: "SourceBlockDelimiter",
						},
					},
				},
			},
		},
		{
			name: "SourceBlockDelimiter",
			pos:  position{line: 249, col: 1, offset: 9947},
			expr: &litMatcher{
				pos:        position{line: 249, col: 25, offset: 9971},
				val:        "```",
				ignoreCase: false,
			},
		},
		{
			name: "SourceBlockLine",
			pos:  position{line: 251, col: 1, offset: 9978},
			expr: &seqExpr{
				pos: position{line: 251, col: 20, offset: 9997},
				exprs: []interface{}{
					&zeroOrMoreExpr{
						pos: position{line: 251, col: 20, offset: 9997},
						expr: &seqExpr{
							pos: position{line: 251, col: 21, offset: 9998},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 251, col: 21, offset: 9998},
									expr: &ruleRefExpr{
										pos:  position{line: 251, col: 22, offset: 9999},
										name: "EOL",
									},
								},
								&anyMatcher{
									line: 251, col: 26, offset: 10003,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 251, col: 30, offset: 10007},
						name: "NEWLINE",
					},
				},
			},
		},
		{
			name: "ElementAttribute",
			pos:  position{line: 256, col: 1, offset: 10125},
			expr: &labeledExpr{
				pos:   position{line: 256, col: 21, offset: 10145},
				label: "meta",
				expr: &choiceExpr{
					pos: position{line: 256, col: 27, offset: 10151},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 256, col: 27, offset: 10151},
							name: "ElementLink",
						},
						&ruleRefExpr{
							pos:  position{line: 256, col: 41, offset: 10165},
							name: "ElementID",
						},
						&ruleRefExpr{
							pos:  position{line: 256, col: 53, offset: 10177},
							name: "ElementTitle",
						},
					},
				},
			},
		},
		{
			name: "ElementLink",
			pos:  position{line: 259, col: 1, offset: 10248},
			expr: &actionExpr{
				pos: position{line: 259, col: 16, offset: 10263},
				run: (*parser).callonElementLink1,
				expr: &seqExpr{
					pos: position{line: 259, col: 16, offset: 10263},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 259, col: 16, offset: 10263},
							val:        "[",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 259, col: 20, offset: 10267},
							expr: &ruleRefExpr{
								pos:  position{line: 259, col: 20, offset: 10267},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 259, col: 24, offset: 10271},
							val:        "link",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 259, col: 31, offset: 10278},
							expr: &ruleRefExpr{
								pos:  position{line: 259, col: 31, offset: 10278},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 259, col: 35, offset: 10282},
							val:        "=",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 259, col: 39, offset: 10286},
							expr: &ruleRefExpr{
								pos:  position{line: 259, col: 39, offset: 10286},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 259, col: 43, offset: 10290},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 259, col: 48, offset: 10295},
								name: "URL",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 259, col: 52, offset: 10299},
							expr: &ruleRefExpr{
								pos:  position{line: 259, col: 52, offset: 10299},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 259, col: 56, offset: 10303},
							val:        "]",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 259, col: 60, offset: 10307},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ElementID",
			pos:  position{line: 264, col: 1, offset: 10417},
			expr: &actionExpr{
				pos: position{line: 264, col: 14, offset: 10430},
				run: (*parser).callonElementID1,
				expr: &seqExpr{
					pos: position{line: 264, col: 14, offset: 10430},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 264, col: 14, offset: 10430},
							val:        "[",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 264, col: 18, offset: 10434},
							expr: &ruleRefExpr{
								pos:  position{line: 264, col: 18, offset: 10434},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 264, col: 22, offset: 10438},
							val:        "#",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 264, col: 26, offset: 10442},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 264, col: 30, offset: 10446},
								name: "ID",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 264, col: 34, offset: 10450},
							expr: &ruleRefExpr{
								pos:  position{line: 264, col: 34, offset: 10450},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 264, col: 38, offset: 10454},
							val:        "]",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 264, col: 42, offset: 10458},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ElementTitle",
			pos:  position{line: 269, col: 1, offset: 10566},
			expr: &actionExpr{
				pos: position{line: 269, col: 17, offset: 10582},
				run: (*parser).callonElementTitle1,
				expr: &seqExpr{
					pos: position{line: 269, col: 17, offset: 10582},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 269, col: 17, offset: 10582},
							val:        ".",
							ignoreCase: false,
						},
						&notExpr{
							pos: position{line: 269, col: 21, offset: 10586},
							expr: &ruleRefExpr{
								pos:  position{line: 269, col: 22, offset: 10587},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 269, col: 25, offset: 10590},
							label: "title",
							expr: &oneOrMoreExpr{
								pos: position{line: 269, col: 31, offset: 10596},
								expr: &seqExpr{
									pos: position{line: 269, col: 32, offset: 10597},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 269, col: 32, offset: 10597},
											expr: &ruleRefExpr{
												pos:  position{line: 269, col: 33, offset: 10598},
												name: "NEWLINE",
											},
										},
										&anyMatcher{
											line: 269, col: 41, offset: 10606,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 269, col: 45, offset: 10610},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "Word",
			pos:  position{line: 276, col: 1, offset: 10781},
			expr: &actionExpr{
				pos: position{line: 276, col: 9, offset: 10789},
				run: (*parser).callonWord1,
				expr: &oneOrMoreExpr{
					pos: position{line: 276, col: 9, offset: 10789},
					expr: &seqExpr{
						pos: position{line: 276, col: 10, offset: 10790},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 276, col: 10, offset: 10790},
								expr: &ruleRefExpr{
									pos:  position{line: 276, col: 11, offset: 10791},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 276, col: 19, offset: 10799},
								expr: &ruleRefExpr{
									pos:  position{line: 276, col: 20, offset: 10800},
									name: "WS",
								},
							},
							&anyMatcher{
								line: 276, col: 23, offset: 10803,
							},
						},
					},
				},
			},
		},
		{
			name: "BlankLine",
			pos:  position{line: 280, col: 1, offset: 10843},
			expr: &actionExpr{
				pos: position{line: 280, col: 14, offset: 10856},
				run: (*parser).callonBlankLine1,
				expr: &seqExpr{
					pos: position{line: 280, col: 14, offset: 10856},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 280, col: 14, offset: 10856},
							expr: &ruleRefExpr{
								pos:  position{line: 280, col: 15, offset: 10857},
								name: "EOF",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 280, col: 19, offset: 10861},
							expr: &ruleRefExpr{
								pos:  position{line: 280, col: 19, offset: 10861},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 280, col: 23, offset: 10865},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "URL",
			pos:  position{line: 284, col: 1, offset: 10906},
			expr: &actionExpr{
				pos: position{line: 284, col: 8, offset: 10913},
				run: (*parser).callonURL1,
				expr: &oneOrMoreExpr{
					pos: position{line: 284, col: 8, offset: 10913},
					expr: &seqExpr{
						pos: position{line: 284, col: 9, offset: 10914},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 284, col: 9, offset: 10914},
								expr: &ruleRefExpr{
									pos:  position{line: 284, col: 10, offset: 10915},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 284, col: 18, offset: 10923},
								expr: &ruleRefExpr{
									pos:  position{line: 284, col: 19, offset: 10924},
									name: "WS",
								},
							},
							&notExpr{
								pos: position{line: 284, col: 22, offset: 10927},
								expr: &litMatcher{
									pos:        position{line: 284, col: 23, offset: 10928},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 284, col: 27, offset: 10932},
								expr: &litMatcher{
									pos:        position{line: 284, col: 28, offset: 10933},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 284, col: 32, offset: 10937,
							},
						},
					},
				},
			},
		},
		{
			name: "ID",
			pos:  position{line: 288, col: 1, offset: 10977},
			expr: &actionExpr{
				pos: position{line: 288, col: 7, offset: 10983},
				run: (*parser).callonID1,
				expr: &oneOrMoreExpr{
					pos: position{line: 288, col: 7, offset: 10983},
					expr: &seqExpr{
						pos: position{line: 288, col: 8, offset: 10984},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 288, col: 8, offset: 10984},
								expr: &ruleRefExpr{
									pos:  position{line: 288, col: 9, offset: 10985},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 288, col: 17, offset: 10993},
								expr: &ruleRefExpr{
									pos:  position{line: 288, col: 18, offset: 10994},
									name: "WS",
								},
							},
							&notExpr{
								pos: position{line: 288, col: 21, offset: 10997},
								expr: &litMatcher{
									pos:        position{line: 288, col: 22, offset: 10998},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 288, col: 26, offset: 11002},
								expr: &litMatcher{
									pos:        position{line: 288, col: 27, offset: 11003},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 288, col: 31, offset: 11007,
							},
						},
					},
				},
			},
		},
		{
			name: "URL_TEXT",
			pos:  position{line: 292, col: 1, offset: 11047},
			expr: &actionExpr{
				pos: position{line: 292, col: 13, offset: 11059},
				run: (*parser).callonURL_TEXT1,
				expr: &oneOrMoreExpr{
					pos: position{line: 292, col: 13, offset: 11059},
					expr: &seqExpr{
						pos: position{line: 292, col: 14, offset: 11060},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 292, col: 14, offset: 11060},
								expr: &ruleRefExpr{
									pos:  position{line: 292, col: 15, offset: 11061},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 292, col: 23, offset: 11069},
								expr: &litMatcher{
									pos:        position{line: 292, col: 24, offset: 11070},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 292, col: 28, offset: 11074},
								expr: &litMatcher{
									pos:        position{line: 292, col: 29, offset: 11075},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 292, col: 33, offset: 11079,
							},
						},
					},
				},
			},
		},
		{
			name: "URL_SCHEME",
			pos:  position{line: 296, col: 1, offset: 11119},
			expr: &choiceExpr{
				pos: position{line: 296, col: 15, offset: 11133},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 296, col: 15, offset: 11133},
						val:        "http://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 296, col: 27, offset: 11145},
						val:        "https://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 296, col: 40, offset: 11158},
						val:        "ftp://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 296, col: 51, offset: 11169},
						val:        "irc://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 296, col: 62, offset: 11180},
						val:        "mailto:",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "DIGIT",
			pos:  position{line: 298, col: 1, offset: 11191},
			expr: &charClassMatcher{
				pos:        position{line: 298, col: 13, offset: 11203},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NEWLINE",
			pos:  position{line: 300, col: 1, offset: 11210},
			expr: &choiceExpr{
				pos: position{line: 300, col: 13, offset: 11222},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 300, col: 13, offset: 11222},
						val:        "\r\n",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 300, col: 22, offset: 11231},
						val:        "\r",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 300, col: 29, offset: 11238},
						val:        "\n",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "WS",
			pos:  position{line: 302, col: 1, offset: 11244},
			expr: &choiceExpr{
				pos: position{line: 302, col: 13, offset: 11256},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 302, col: 13, offset: 11256},
						val:        " ",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 302, col: 19, offset: 11262},
						run: (*parser).callonWS3,
						expr: &litMatcher{
							pos:        position{line: 302, col: 19, offset: 11262},
							val:        "\t",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 306, col: 1, offset: 11307},
			expr: &notExpr{
				pos: position{line: 306, col: 13, offset: 11319},
				expr: &anyMatcher{
					line: 306, col: 14, offset: 11320,
				},
			},
		},
		{
			name: "EOL",
			pos:  position{line: 308, col: 1, offset: 11323},
			expr: &choiceExpr{
				pos: position{line: 308, col: 13, offset: 11335},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 308, col: 13, offset: 11335},
						name: "NEWLINE",
					},
					&ruleRefExpr{
						pos:  position{line: 308, col: 23, offset: 11345},
						name: "EOF",
					},
				},
			},
		},
	},
}

func (c *current) onDocument1(frontmatter, blocks interface{}) (interface{}, error) {
	if frontmatter != nil {
		return types.NewDocument(frontmatter.(*types.FrontMatter), blocks.([]interface{}))
	}
	return types.NewDocument(nil, blocks.([]interface{}))
}

func (p *parser) callonDocument1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocument1(stack["frontmatter"], stack["blocks"])
}

func (c *current) onDocumentBlock1(content interface{}) (interface{}, error) {
	return content.(types.DocElement), nil
}

func (p *parser) callonDocumentBlock1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocumentBlock1(stack["content"])
}

func (c *current) onFrontMatter1(content interface{}) (interface{}, error) {
	return types.NewYamlFrontMatter(content.([]interface{}))
}

func (p *parser) callonFrontMatter1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFrontMatter1(stack["content"])
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
