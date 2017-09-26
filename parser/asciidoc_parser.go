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
			pos:  position{line: 16, col: 1, offset: 456},
			expr: &actionExpr{
				pos: position{line: 16, col: 13, offset: 468},
				run: (*parser).callonDocument1,
				expr: &seqExpr{
					pos: position{line: 16, col: 13, offset: 468},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 16, col: 13, offset: 468},
							label: "frontmatter",
							expr: &zeroOrOneExpr{
								pos: position{line: 16, col: 26, offset: 481},
								expr: &ruleRefExpr{
									pos:  position{line: 16, col: 26, offset: 481},
									name: "FrontMatter",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 16, col: 40, offset: 495},
							label: "blocks",
							expr: &zeroOrMoreExpr{
								pos: position{line: 16, col: 48, offset: 503},
								expr: &ruleRefExpr{
									pos:  position{line: 16, col: 48, offset: 503},
									name: "DocumentBlock",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 16, col: 64, offset: 519},
							name: "EOF",
						},
					},
				},
			},
		},
		{
			name: "DocumentBlock",
			pos:  position{line: 23, col: 1, offset: 705},
			expr: &actionExpr{
				pos: position{line: 23, col: 18, offset: 722},
				run: (*parser).callonDocumentBlock1,
				expr: &labeledExpr{
					pos:   position{line: 23, col: 18, offset: 722},
					label: "content",
					expr: &choiceExpr{
						pos: position{line: 23, col: 27, offset: 731},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 23, col: 27, offset: 731},
								name: "Section",
							},
							&ruleRefExpr{
								pos:  position{line: 23, col: 37, offset: 741},
								name: "StandaloneBlock",
							},
						},
					},
				},
			},
		},
		{
			name: "StandaloneBlock",
			pos:  position{line: 27, col: 1, offset: 806},
			expr: &choiceExpr{
				pos: position{line: 27, col: 20, offset: 825},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 27, col: 20, offset: 825},
						name: "DocumentAttributeDeclaration",
					},
					&ruleRefExpr{
						pos:  position{line: 27, col: 51, offset: 856},
						name: "DocumentAttributeReset",
					},
					&ruleRefExpr{
						pos:  position{line: 27, col: 76, offset: 881},
						name: "List",
					},
					&ruleRefExpr{
						pos:  position{line: 27, col: 83, offset: 888},
						name: "BlockImage",
					},
					&ruleRefExpr{
						pos:  position{line: 27, col: 96, offset: 901},
						name: "LiteralBlock",
					},
					&ruleRefExpr{
						pos:  position{line: 27, col: 111, offset: 916},
						name: "DelimitedBlock",
					},
					&ruleRefExpr{
						pos:  position{line: 27, col: 128, offset: 933},
						name: "Paragraph",
					},
					&ruleRefExpr{
						pos:  position{line: 27, col: 140, offset: 945},
						name: "ElementAttribute",
					},
					&ruleRefExpr{
						pos:  position{line: 27, col: 159, offset: 964},
						name: "BlankLine",
					},
				},
			},
		},
		{
			name: "FrontMatter",
			pos:  position{line: 32, col: 1, offset: 1127},
			expr: &ruleRefExpr{
				pos:  position{line: 32, col: 16, offset: 1142},
				name: "YamlFrontMatter",
			},
		},
		{
			name: "FrontMatter",
			pos:  position{line: 34, col: 1, offset: 1160},
			expr: &actionExpr{
				pos: position{line: 34, col: 16, offset: 1175},
				run: (*parser).callonFrontMatter1,
				expr: &seqExpr{
					pos: position{line: 34, col: 16, offset: 1175},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 34, col: 16, offset: 1175},
							name: "YamlFrontMatterToken",
						},
						&labeledExpr{
							pos:   position{line: 34, col: 37, offset: 1196},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 34, col: 45, offset: 1204},
								expr: &seqExpr{
									pos: position{line: 34, col: 46, offset: 1205},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 34, col: 46, offset: 1205},
											expr: &ruleRefExpr{
												pos:  position{line: 34, col: 47, offset: 1206},
												name: "YamlFrontMatterToken",
											},
										},
										&anyMatcher{
											line: 34, col: 68, offset: 1227,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 34, col: 72, offset: 1231},
							name: "YamlFrontMatterToken",
						},
					},
				},
			},
		},
		{
			name: "YamlFrontMatterToken",
			pos:  position{line: 38, col: 1, offset: 1318},
			expr: &seqExpr{
				pos: position{line: 38, col: 26, offset: 1343},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 38, col: 26, offset: 1343},
						val:        "---",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 38, col: 32, offset: 1349},
						name: "EOL",
					},
				},
			},
		},
		{
			name: "Section",
			pos:  position{line: 43, col: 1, offset: 1458},
			expr: &choiceExpr{
				pos: position{line: 43, col: 12, offset: 1469},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 43, col: 12, offset: 1469},
						name: "Section1",
					},
					&ruleRefExpr{
						pos:  position{line: 43, col: 23, offset: 1480},
						name: "Section2",
					},
					&ruleRefExpr{
						pos:  position{line: 43, col: 34, offset: 1491},
						name: "Section3",
					},
					&ruleRefExpr{
						pos:  position{line: 43, col: 45, offset: 1502},
						name: "Section4",
					},
					&ruleRefExpr{
						pos:  position{line: 43, col: 56, offset: 1513},
						name: "Section5",
					},
					&ruleRefExpr{
						pos:  position{line: 43, col: 67, offset: 1524},
						name: "Section6",
					},
				},
			},
		},
		{
			name: "Section1",
			pos:  position{line: 45, col: 1, offset: 1534},
			expr: &actionExpr{
				pos: position{line: 45, col: 13, offset: 1546},
				run: (*parser).callonSection11,
				expr: &seqExpr{
					pos: position{line: 45, col: 13, offset: 1546},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 45, col: 13, offset: 1546},
							label: "heading",
							expr: &ruleRefExpr{
								pos:  position{line: 45, col: 22, offset: 1555},
								name: "Heading1",
							},
						},
						&labeledExpr{
							pos:   position{line: 45, col: 32, offset: 1565},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 45, col: 42, offset: 1575},
								expr: &ruleRefExpr{
									pos:  position{line: 45, col: 42, offset: 1575},
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
			pos:  position{line: 49, col: 1, offset: 1676},
			expr: &actionExpr{
				pos: position{line: 49, col: 18, offset: 1693},
				run: (*parser).callonSection1Block1,
				expr: &labeledExpr{
					pos:   position{line: 49, col: 18, offset: 1693},
					label: "content",
					expr: &choiceExpr{
						pos: position{line: 49, col: 27, offset: 1702},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 49, col: 27, offset: 1702},
								name: "Section2",
							},
							&ruleRefExpr{
								pos:  position{line: 49, col: 38, offset: 1713},
								name: "Section3",
							},
							&ruleRefExpr{
								pos:  position{line: 49, col: 49, offset: 1724},
								name: "Section4",
							},
							&ruleRefExpr{
								pos:  position{line: 49, col: 60, offset: 1735},
								name: "Section5",
							},
							&ruleRefExpr{
								pos:  position{line: 49, col: 71, offset: 1746},
								name: "Section6",
							},
							&ruleRefExpr{
								pos:  position{line: 49, col: 82, offset: 1757},
								name: "StandaloneBlock",
							},
						},
					},
				},
			},
		},
		{
			name: "Section2",
			pos:  position{line: 53, col: 1, offset: 1822},
			expr: &actionExpr{
				pos: position{line: 53, col: 13, offset: 1834},
				run: (*parser).callonSection21,
				expr: &seqExpr{
					pos: position{line: 53, col: 13, offset: 1834},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 53, col: 13, offset: 1834},
							label: "heading",
							expr: &ruleRefExpr{
								pos:  position{line: 53, col: 22, offset: 1843},
								name: "Heading2",
							},
						},
						&labeledExpr{
							pos:   position{line: 53, col: 32, offset: 1853},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 53, col: 42, offset: 1863},
								expr: &ruleRefExpr{
									pos:  position{line: 53, col: 42, offset: 1863},
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
			pos:  position{line: 57, col: 1, offset: 1964},
			expr: &actionExpr{
				pos: position{line: 57, col: 18, offset: 1981},
				run: (*parser).callonSection2Block1,
				expr: &seqExpr{
					pos: position{line: 57, col: 18, offset: 1981},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 57, col: 18, offset: 1981},
							expr: &ruleRefExpr{
								pos:  position{line: 57, col: 19, offset: 1982},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 57, col: 28, offset: 1991},
							expr: &ruleRefExpr{
								pos:  position{line: 57, col: 29, offset: 1992},
								name: "Section2",
							},
						},
						&labeledExpr{
							pos:   position{line: 57, col: 38, offset: 2001},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 57, col: 47, offset: 2010},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 57, col: 47, offset: 2010},
										name: "Section3",
									},
									&ruleRefExpr{
										pos:  position{line: 57, col: 58, offset: 2021},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 57, col: 69, offset: 2032},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 57, col: 80, offset: 2043},
										name: "Section6",
									},
									&ruleRefExpr{
										pos:  position{line: 57, col: 91, offset: 2054},
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
			pos:  position{line: 61, col: 1, offset: 2119},
			expr: &actionExpr{
				pos: position{line: 61, col: 13, offset: 2131},
				run: (*parser).callonSection31,
				expr: &seqExpr{
					pos: position{line: 61, col: 13, offset: 2131},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 61, col: 13, offset: 2131},
							label: "heading",
							expr: &ruleRefExpr{
								pos:  position{line: 61, col: 22, offset: 2140},
								name: "Heading3",
							},
						},
						&labeledExpr{
							pos:   position{line: 61, col: 32, offset: 2150},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 61, col: 42, offset: 2160},
								expr: &ruleRefExpr{
									pos:  position{line: 61, col: 42, offset: 2160},
									name: "Section3Block",
								},
							},
						},
						&andExpr{
							pos: position{line: 61, col: 58, offset: 2176},
							expr: &zeroOrMoreExpr{
								pos: position{line: 61, col: 59, offset: 2177},
								expr: &ruleRefExpr{
									pos:  position{line: 61, col: 60, offset: 2178},
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
			pos:  position{line: 65, col: 1, offset: 2274},
			expr: &actionExpr{
				pos: position{line: 65, col: 18, offset: 2291},
				run: (*parser).callonSection3Block1,
				expr: &seqExpr{
					pos: position{line: 65, col: 18, offset: 2291},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 65, col: 18, offset: 2291},
							expr: &ruleRefExpr{
								pos:  position{line: 65, col: 19, offset: 2292},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 65, col: 28, offset: 2301},
							expr: &ruleRefExpr{
								pos:  position{line: 65, col: 29, offset: 2302},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 65, col: 38, offset: 2311},
							expr: &ruleRefExpr{
								pos:  position{line: 65, col: 39, offset: 2312},
								name: "Section3",
							},
						},
						&labeledExpr{
							pos:   position{line: 65, col: 48, offset: 2321},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 65, col: 57, offset: 2330},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 65, col: 57, offset: 2330},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 65, col: 68, offset: 2341},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 65, col: 79, offset: 2352},
										name: "Section6",
									},
									&ruleRefExpr{
										pos:  position{line: 65, col: 90, offset: 2363},
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
			pos:  position{line: 69, col: 1, offset: 2428},
			expr: &actionExpr{
				pos: position{line: 69, col: 13, offset: 2440},
				run: (*parser).callonSection41,
				expr: &seqExpr{
					pos: position{line: 69, col: 13, offset: 2440},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 69, col: 13, offset: 2440},
							label: "heading",
							expr: &ruleRefExpr{
								pos:  position{line: 69, col: 22, offset: 2449},
								name: "Heading4",
							},
						},
						&labeledExpr{
							pos:   position{line: 69, col: 32, offset: 2459},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 69, col: 42, offset: 2469},
								expr: &ruleRefExpr{
									pos:  position{line: 69, col: 42, offset: 2469},
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
			pos:  position{line: 73, col: 1, offset: 2570},
			expr: &actionExpr{
				pos: position{line: 73, col: 18, offset: 2587},
				run: (*parser).callonSection4Block1,
				expr: &seqExpr{
					pos: position{line: 73, col: 18, offset: 2587},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 73, col: 18, offset: 2587},
							expr: &ruleRefExpr{
								pos:  position{line: 73, col: 19, offset: 2588},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 73, col: 28, offset: 2597},
							expr: &ruleRefExpr{
								pos:  position{line: 73, col: 29, offset: 2598},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 73, col: 38, offset: 2607},
							expr: &ruleRefExpr{
								pos:  position{line: 73, col: 39, offset: 2608},
								name: "Section3",
							},
						},
						&notExpr{
							pos: position{line: 73, col: 48, offset: 2617},
							expr: &ruleRefExpr{
								pos:  position{line: 73, col: 49, offset: 2618},
								name: "Section4",
							},
						},
						&labeledExpr{
							pos:   position{line: 73, col: 58, offset: 2627},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 73, col: 67, offset: 2636},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 73, col: 67, offset: 2636},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 73, col: 78, offset: 2647},
										name: "Section6",
									},
									&ruleRefExpr{
										pos:  position{line: 73, col: 89, offset: 2658},
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
			pos:  position{line: 77, col: 1, offset: 2723},
			expr: &actionExpr{
				pos: position{line: 77, col: 13, offset: 2735},
				run: (*parser).callonSection51,
				expr: &seqExpr{
					pos: position{line: 77, col: 13, offset: 2735},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 77, col: 13, offset: 2735},
							label: "heading",
							expr: &ruleRefExpr{
								pos:  position{line: 77, col: 22, offset: 2744},
								name: "Heading5",
							},
						},
						&labeledExpr{
							pos:   position{line: 77, col: 32, offset: 2754},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 77, col: 42, offset: 2764},
								expr: &ruleRefExpr{
									pos:  position{line: 77, col: 42, offset: 2764},
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
			pos:  position{line: 81, col: 1, offset: 2865},
			expr: &actionExpr{
				pos: position{line: 81, col: 18, offset: 2882},
				run: (*parser).callonSection5Block1,
				expr: &seqExpr{
					pos: position{line: 81, col: 18, offset: 2882},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 81, col: 18, offset: 2882},
							expr: &ruleRefExpr{
								pos:  position{line: 81, col: 19, offset: 2883},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 81, col: 28, offset: 2892},
							expr: &ruleRefExpr{
								pos:  position{line: 81, col: 29, offset: 2893},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 81, col: 38, offset: 2902},
							expr: &ruleRefExpr{
								pos:  position{line: 81, col: 39, offset: 2903},
								name: "Section3",
							},
						},
						&notExpr{
							pos: position{line: 81, col: 48, offset: 2912},
							expr: &ruleRefExpr{
								pos:  position{line: 81, col: 49, offset: 2913},
								name: "Section4",
							},
						},
						&notExpr{
							pos: position{line: 81, col: 58, offset: 2922},
							expr: &ruleRefExpr{
								pos:  position{line: 81, col: 59, offset: 2923},
								name: "Section5",
							},
						},
						&labeledExpr{
							pos:   position{line: 81, col: 68, offset: 2932},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 81, col: 77, offset: 2941},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 81, col: 77, offset: 2941},
										name: "Section6",
									},
									&ruleRefExpr{
										pos:  position{line: 81, col: 88, offset: 2952},
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
			pos:  position{line: 85, col: 1, offset: 3017},
			expr: &actionExpr{
				pos: position{line: 85, col: 13, offset: 3029},
				run: (*parser).callonSection61,
				expr: &seqExpr{
					pos: position{line: 85, col: 13, offset: 3029},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 85, col: 13, offset: 3029},
							label: "heading",
							expr: &ruleRefExpr{
								pos:  position{line: 85, col: 22, offset: 3038},
								name: "Heading6",
							},
						},
						&labeledExpr{
							pos:   position{line: 85, col: 32, offset: 3048},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 85, col: 42, offset: 3058},
								expr: &ruleRefExpr{
									pos:  position{line: 85, col: 42, offset: 3058},
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
			pos:  position{line: 89, col: 1, offset: 3159},
			expr: &actionExpr{
				pos: position{line: 89, col: 18, offset: 3176},
				run: (*parser).callonSection6Block1,
				expr: &seqExpr{
					pos: position{line: 89, col: 18, offset: 3176},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 89, col: 18, offset: 3176},
							expr: &ruleRefExpr{
								pos:  position{line: 89, col: 19, offset: 3177},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 89, col: 28, offset: 3186},
							expr: &ruleRefExpr{
								pos:  position{line: 89, col: 29, offset: 3187},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 89, col: 38, offset: 3196},
							expr: &ruleRefExpr{
								pos:  position{line: 89, col: 39, offset: 3197},
								name: "Section3",
							},
						},
						&notExpr{
							pos: position{line: 89, col: 48, offset: 3206},
							expr: &ruleRefExpr{
								pos:  position{line: 89, col: 49, offset: 3207},
								name: "Section4",
							},
						},
						&notExpr{
							pos: position{line: 89, col: 58, offset: 3216},
							expr: &ruleRefExpr{
								pos:  position{line: 89, col: 59, offset: 3217},
								name: "Section5",
							},
						},
						&notExpr{
							pos: position{line: 89, col: 68, offset: 3226},
							expr: &ruleRefExpr{
								pos:  position{line: 89, col: 69, offset: 3227},
								name: "Section6",
							},
						},
						&labeledExpr{
							pos:   position{line: 89, col: 78, offset: 3236},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 89, col: 87, offset: 3245},
								name: "StandaloneBlock",
							},
						},
					},
				},
			},
		},
		{
			name: "Heading",
			pos:  position{line: 96, col: 1, offset: 3414},
			expr: &choiceExpr{
				pos: position{line: 96, col: 12, offset: 3425},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 96, col: 12, offset: 3425},
						name: "Heading1",
					},
					&ruleRefExpr{
						pos:  position{line: 96, col: 23, offset: 3436},
						name: "Heading2",
					},
					&ruleRefExpr{
						pos:  position{line: 96, col: 34, offset: 3447},
						name: "Heading3",
					},
					&ruleRefExpr{
						pos:  position{line: 96, col: 45, offset: 3458},
						name: "Heading4",
					},
					&ruleRefExpr{
						pos:  position{line: 96, col: 56, offset: 3469},
						name: "Heading5",
					},
					&ruleRefExpr{
						pos:  position{line: 96, col: 67, offset: 3480},
						name: "Heading6",
					},
				},
			},
		},
		{
			name: "Heading1",
			pos:  position{line: 98, col: 1, offset: 3490},
			expr: &actionExpr{
				pos: position{line: 98, col: 13, offset: 3502},
				run: (*parser).callonHeading11,
				expr: &seqExpr{
					pos: position{line: 98, col: 13, offset: 3502},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 98, col: 13, offset: 3502},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 98, col: 24, offset: 3513},
								expr: &ruleRefExpr{
									pos:  position{line: 98, col: 25, offset: 3514},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 98, col: 44, offset: 3533},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 98, col: 51, offset: 3540},
								val:        "=",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 98, col: 56, offset: 3545},
							expr: &ruleRefExpr{
								pos:  position{line: 98, col: 56, offset: 3545},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 98, col: 60, offset: 3549},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 98, col: 68, offset: 3557},
								name: "InlineContent",
							},
						},
						&choiceExpr{
							pos: position{line: 98, col: 83, offset: 3572},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 98, col: 83, offset: 3572},
									expr: &ruleRefExpr{
										pos:  position{line: 98, col: 83, offset: 3572},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 96, offset: 3585},
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
			pos:  position{line: 102, col: 1, offset: 3792},
			expr: &actionExpr{
				pos: position{line: 102, col: 13, offset: 3804},
				run: (*parser).callonHeading21,
				expr: &seqExpr{
					pos: position{line: 102, col: 13, offset: 3804},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 102, col: 13, offset: 3804},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 102, col: 24, offset: 3815},
								expr: &ruleRefExpr{
									pos:  position{line: 102, col: 25, offset: 3816},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 102, col: 44, offset: 3835},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 102, col: 51, offset: 3842},
								val:        "==",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 102, col: 57, offset: 3848},
							expr: &ruleRefExpr{
								pos:  position{line: 102, col: 57, offset: 3848},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 102, col: 61, offset: 3852},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 102, col: 69, offset: 3860},
								name: "InlineContent",
							},
						},
						&choiceExpr{
							pos: position{line: 102, col: 84, offset: 3875},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 102, col: 84, offset: 3875},
									expr: &ruleRefExpr{
										pos:  position{line: 102, col: 84, offset: 3875},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 102, col: 97, offset: 3888},
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
			pos:  position{line: 106, col: 1, offset: 3990},
			expr: &actionExpr{
				pos: position{line: 106, col: 13, offset: 4002},
				run: (*parser).callonHeading31,
				expr: &seqExpr{
					pos: position{line: 106, col: 13, offset: 4002},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 106, col: 13, offset: 4002},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 106, col: 24, offset: 4013},
								expr: &ruleRefExpr{
									pos:  position{line: 106, col: 25, offset: 4014},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 106, col: 44, offset: 4033},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 106, col: 51, offset: 4040},
								val:        "===",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 106, col: 58, offset: 4047},
							expr: &ruleRefExpr{
								pos:  position{line: 106, col: 58, offset: 4047},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 106, col: 62, offset: 4051},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 106, col: 70, offset: 4059},
								name: "InlineContent",
							},
						},
						&choiceExpr{
							pos: position{line: 106, col: 85, offset: 4074},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 106, col: 85, offset: 4074},
									expr: &ruleRefExpr{
										pos:  position{line: 106, col: 85, offset: 4074},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 106, col: 98, offset: 4087},
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
			pos:  position{line: 110, col: 1, offset: 4189},
			expr: &actionExpr{
				pos: position{line: 110, col: 13, offset: 4201},
				run: (*parser).callonHeading41,
				expr: &seqExpr{
					pos: position{line: 110, col: 13, offset: 4201},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 110, col: 13, offset: 4201},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 110, col: 24, offset: 4212},
								expr: &ruleRefExpr{
									pos:  position{line: 110, col: 25, offset: 4213},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 110, col: 44, offset: 4232},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 110, col: 51, offset: 4239},
								val:        "====",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 110, col: 59, offset: 4247},
							expr: &ruleRefExpr{
								pos:  position{line: 110, col: 59, offset: 4247},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 110, col: 63, offset: 4251},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 110, col: 71, offset: 4259},
								name: "InlineContent",
							},
						},
						&choiceExpr{
							pos: position{line: 110, col: 86, offset: 4274},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 110, col: 86, offset: 4274},
									expr: &ruleRefExpr{
										pos:  position{line: 110, col: 86, offset: 4274},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 110, col: 99, offset: 4287},
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
			pos:  position{line: 114, col: 1, offset: 4389},
			expr: &actionExpr{
				pos: position{line: 114, col: 13, offset: 4401},
				run: (*parser).callonHeading51,
				expr: &seqExpr{
					pos: position{line: 114, col: 13, offset: 4401},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 114, col: 13, offset: 4401},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 114, col: 24, offset: 4412},
								expr: &ruleRefExpr{
									pos:  position{line: 114, col: 25, offset: 4413},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 114, col: 44, offset: 4432},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 114, col: 51, offset: 4439},
								val:        "=====",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 114, col: 60, offset: 4448},
							expr: &ruleRefExpr{
								pos:  position{line: 114, col: 60, offset: 4448},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 114, col: 64, offset: 4452},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 114, col: 72, offset: 4460},
								name: "InlineContent",
							},
						},
						&choiceExpr{
							pos: position{line: 114, col: 87, offset: 4475},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 114, col: 87, offset: 4475},
									expr: &ruleRefExpr{
										pos:  position{line: 114, col: 87, offset: 4475},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 114, col: 100, offset: 4488},
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
			pos:  position{line: 118, col: 1, offset: 4590},
			expr: &actionExpr{
				pos: position{line: 118, col: 13, offset: 4602},
				run: (*parser).callonHeading61,
				expr: &seqExpr{
					pos: position{line: 118, col: 13, offset: 4602},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 118, col: 13, offset: 4602},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 118, col: 24, offset: 4613},
								expr: &ruleRefExpr{
									pos:  position{line: 118, col: 25, offset: 4614},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 118, col: 44, offset: 4633},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 118, col: 51, offset: 4640},
								val:        "======",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 118, col: 61, offset: 4650},
							expr: &ruleRefExpr{
								pos:  position{line: 118, col: 61, offset: 4650},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 118, col: 65, offset: 4654},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 118, col: 73, offset: 4662},
								name: "InlineContent",
							},
						},
						&choiceExpr{
							pos: position{line: 118, col: 88, offset: 4677},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 118, col: 88, offset: 4677},
									expr: &ruleRefExpr{
										pos:  position{line: 118, col: 88, offset: 4677},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 118, col: 101, offset: 4690},
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
			pos:  position{line: 126, col: 1, offset: 4908},
			expr: &choiceExpr{
				pos: position{line: 126, col: 33, offset: 4940},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 126, col: 33, offset: 4940},
						name: "DocumentAttributeDeclarationWithNameOnly",
					},
					&ruleRefExpr{
						pos:  position{line: 126, col: 76, offset: 4983},
						name: "DocumentAttributeDeclarationWithNameAndValue",
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclarationWithNameOnly",
			pos:  position{line: 128, col: 1, offset: 5030},
			expr: &actionExpr{
				pos: position{line: 128, col: 45, offset: 5074},
				run: (*parser).callonDocumentAttributeDeclarationWithNameOnly1,
				expr: &seqExpr{
					pos: position{line: 128, col: 45, offset: 5074},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 128, col: 45, offset: 5074},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 128, col: 49, offset: 5078},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 128, col: 55, offset: 5084},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 128, col: 70, offset: 5099},
							val:        ":",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 128, col: 74, offset: 5103},
							expr: &ruleRefExpr{
								pos:  position{line: 128, col: 74, offset: 5103},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 128, col: 78, offset: 5107},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclarationWithNameAndValue",
			pos:  position{line: 132, col: 1, offset: 5192},
			expr: &actionExpr{
				pos: position{line: 132, col: 49, offset: 5240},
				run: (*parser).callonDocumentAttributeDeclarationWithNameAndValue1,
				expr: &seqExpr{
					pos: position{line: 132, col: 49, offset: 5240},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 132, col: 49, offset: 5240},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 132, col: 53, offset: 5244},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 132, col: 59, offset: 5250},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 132, col: 74, offset: 5265},
							val:        ":",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 132, col: 78, offset: 5269},
							expr: &ruleRefExpr{
								pos:  position{line: 132, col: 78, offset: 5269},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 132, col: 82, offset: 5273},
							label: "value",
							expr: &zeroOrMoreExpr{
								pos: position{line: 132, col: 88, offset: 5279},
								expr: &seqExpr{
									pos: position{line: 132, col: 89, offset: 5280},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 132, col: 89, offset: 5280},
											expr: &ruleRefExpr{
												pos:  position{line: 132, col: 90, offset: 5281},
												name: "NEWLINE",
											},
										},
										&anyMatcher{
											line: 132, col: 98, offset: 5289,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 132, col: 102, offset: 5293},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeReset",
			pos:  position{line: 136, col: 1, offset: 5396},
			expr: &choiceExpr{
				pos: position{line: 136, col: 27, offset: 5422},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 136, col: 27, offset: 5422},
						name: "DocumentAttributeResetWithHeadingBangSymbol",
					},
					&ruleRefExpr{
						pos:  position{line: 136, col: 73, offset: 5468},
						name: "DocumentAttributeResetWithTrailingBangSymbol",
					},
				},
			},
		},
		{
			name: "DocumentAttributeResetWithHeadingBangSymbol",
			pos:  position{line: 138, col: 1, offset: 5514},
			expr: &actionExpr{
				pos: position{line: 138, col: 48, offset: 5561},
				run: (*parser).callonDocumentAttributeResetWithHeadingBangSymbol1,
				expr: &seqExpr{
					pos: position{line: 138, col: 48, offset: 5561},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 138, col: 48, offset: 5561},
							val:        ":!",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 138, col: 53, offset: 5566},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 138, col: 59, offset: 5572},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 138, col: 74, offset: 5587},
							val:        ":",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 138, col: 78, offset: 5591},
							expr: &ruleRefExpr{
								pos:  position{line: 138, col: 78, offset: 5591},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 138, col: 82, offset: 5595},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeResetWithTrailingBangSymbol",
			pos:  position{line: 142, col: 1, offset: 5669},
			expr: &actionExpr{
				pos: position{line: 142, col: 49, offset: 5717},
				run: (*parser).callonDocumentAttributeResetWithTrailingBangSymbol1,
				expr: &seqExpr{
					pos: position{line: 142, col: 49, offset: 5717},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 142, col: 49, offset: 5717},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 142, col: 53, offset: 5721},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 142, col: 59, offset: 5727},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 142, col: 74, offset: 5742},
							val:        "!:",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 142, col: 79, offset: 5747},
							expr: &ruleRefExpr{
								pos:  position{line: 142, col: 79, offset: 5747},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 142, col: 83, offset: 5751},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeSubstitution",
			pos:  position{line: 147, col: 1, offset: 5826},
			expr: &actionExpr{
				pos: position{line: 147, col: 34, offset: 5859},
				run: (*parser).callonDocumentAttributeSubstitution1,
				expr: &seqExpr{
					pos: position{line: 147, col: 34, offset: 5859},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 147, col: 34, offset: 5859},
							val:        "{",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 147, col: 38, offset: 5863},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 147, col: 44, offset: 5869},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 147, col: 59, offset: 5884},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "AttributeName",
			pos:  position{line: 154, col: 1, offset: 6138},
			expr: &seqExpr{
				pos: position{line: 154, col: 18, offset: 6155},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 154, col: 19, offset: 6156},
						alternatives: []interface{}{
							&charClassMatcher{
								pos:        position{line: 154, col: 19, offset: 6156},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
							&charClassMatcher{
								pos:        position{line: 154, col: 27, offset: 6164},
								val:        "[a-z]",
								ranges:     []rune{'a', 'z'},
								ignoreCase: false,
								inverted:   false,
							},
							&charClassMatcher{
								pos:        position{line: 154, col: 35, offset: 6172},
								val:        "[0-9]",
								ranges:     []rune{'0', '9'},
								ignoreCase: false,
								inverted:   false,
							},
							&litMatcher{
								pos:        position{line: 154, col: 43, offset: 6180},
								val:        "_",
								ignoreCase: false,
							},
						},
					},
					&zeroOrMoreExpr{
						pos: position{line: 154, col: 48, offset: 6185},
						expr: &choiceExpr{
							pos: position{line: 154, col: 49, offset: 6186},
							alternatives: []interface{}{
								&charClassMatcher{
									pos:        position{line: 154, col: 49, offset: 6186},
									val:        "[A-Z]",
									ranges:     []rune{'A', 'Z'},
									ignoreCase: false,
									inverted:   false,
								},
								&charClassMatcher{
									pos:        position{line: 154, col: 57, offset: 6194},
									val:        "[a-z]",
									ranges:     []rune{'a', 'z'},
									ignoreCase: false,
									inverted:   false,
								},
								&charClassMatcher{
									pos:        position{line: 154, col: 65, offset: 6202},
									val:        "[0-9]",
									ranges:     []rune{'0', '9'},
									ignoreCase: false,
									inverted:   false,
								},
								&litMatcher{
									pos:        position{line: 154, col: 73, offset: 6210},
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
			pos:  position{line: 159, col: 1, offset: 6323},
			expr: &actionExpr{
				pos: position{line: 159, col: 9, offset: 6331},
				run: (*parser).callonList1,
				expr: &seqExpr{
					pos: position{line: 159, col: 9, offset: 6331},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 159, col: 9, offset: 6331},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 159, col: 20, offset: 6342},
								expr: &ruleRefExpr{
									pos:  position{line: 159, col: 21, offset: 6343},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 161, col: 5, offset: 6435},
							label: "elements",
							expr: &oneOrMoreExpr{
								pos: position{line: 161, col: 14, offset: 6444},
								expr: &seqExpr{
									pos: position{line: 161, col: 15, offset: 6445},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 161, col: 15, offset: 6445},
											name: "ListItem",
										},
										&zeroOrOneExpr{
											pos: position{line: 161, col: 24, offset: 6454},
											expr: &ruleRefExpr{
												pos:  position{line: 161, col: 24, offset: 6454},
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
			pos:  position{line: 165, col: 1, offset: 6551},
			expr: &actionExpr{
				pos: position{line: 165, col: 13, offset: 6563},
				run: (*parser).callonListItem1,
				expr: &seqExpr{
					pos: position{line: 165, col: 13, offset: 6563},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 165, col: 13, offset: 6563},
							expr: &ruleRefExpr{
								pos:  position{line: 165, col: 13, offset: 6563},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 165, col: 17, offset: 6567},
							label: "level",
							expr: &choiceExpr{
								pos: position{line: 165, col: 24, offset: 6574},
								alternatives: []interface{}{
									&oneOrMoreExpr{
										pos: position{line: 165, col: 24, offset: 6574},
										expr: &litMatcher{
											pos:        position{line: 165, col: 24, offset: 6574},
											val:        "*",
											ignoreCase: false,
										},
									},
									&litMatcher{
										pos:        position{line: 165, col: 31, offset: 6581},
										val:        "-",
										ignoreCase: false,
									},
								},
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 165, col: 36, offset: 6586},
							expr: &ruleRefExpr{
								pos:  position{line: 165, col: 36, offset: 6586},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 165, col: 40, offset: 6590},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 165, col: 49, offset: 6599},
								name: "ListItemContent",
							},
						},
					},
				},
			},
		},
		{
			name: "ListItemContent",
			pos:  position{line: 169, col: 1, offset: 6696},
			expr: &actionExpr{
				pos: position{line: 169, col: 20, offset: 6715},
				run: (*parser).callonListItemContent1,
				expr: &labeledExpr{
					pos:   position{line: 169, col: 20, offset: 6715},
					label: "lines",
					expr: &oneOrMoreExpr{
						pos: position{line: 169, col: 26, offset: 6721},
						expr: &seqExpr{
							pos: position{line: 169, col: 27, offset: 6722},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 169, col: 27, offset: 6722},
									expr: &seqExpr{
										pos: position{line: 169, col: 29, offset: 6724},
										exprs: []interface{}{
											&zeroOrMoreExpr{
												pos: position{line: 169, col: 29, offset: 6724},
												expr: &ruleRefExpr{
													pos:  position{line: 169, col: 29, offset: 6724},
													name: "WS",
												},
											},
											&choiceExpr{
												pos: position{line: 169, col: 34, offset: 6729},
												alternatives: []interface{}{
													&oneOrMoreExpr{
														pos: position{line: 169, col: 34, offset: 6729},
														expr: &litMatcher{
															pos:        position{line: 169, col: 34, offset: 6729},
															val:        "*",
															ignoreCase: false,
														},
													},
													&litMatcher{
														pos:        position{line: 169, col: 41, offset: 6736},
														val:        "-",
														ignoreCase: false,
													},
												},
											},
											&oneOrMoreExpr{
												pos: position{line: 169, col: 46, offset: 6741},
												expr: &ruleRefExpr{
													pos:  position{line: 169, col: 46, offset: 6741},
													name: "WS",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 169, col: 51, offset: 6746},
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
			pos:  position{line: 176, col: 1, offset: 7016},
			expr: &actionExpr{
				pos: position{line: 176, col: 14, offset: 7029},
				run: (*parser).callonParagraph1,
				expr: &seqExpr{
					pos: position{line: 176, col: 14, offset: 7029},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 176, col: 14, offset: 7029},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 176, col: 25, offset: 7040},
								expr: &ruleRefExpr{
									pos:  position{line: 176, col: 26, offset: 7041},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 176, col: 45, offset: 7060},
							label: "lines",
							expr: &oneOrMoreExpr{
								pos: position{line: 176, col: 51, offset: 7066},
								expr: &ruleRefExpr{
									pos:  position{line: 176, col: 52, offset: 7067},
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
			pos:  position{line: 182, col: 1, offset: 7375},
			expr: &actionExpr{
				pos: position{line: 182, col: 18, offset: 7392},
				run: (*parser).callonInlineContent1,
				expr: &seqExpr{
					pos: position{line: 182, col: 18, offset: 7392},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 182, col: 18, offset: 7392},
							expr: &ruleRefExpr{
								pos:  position{line: 182, col: 19, offset: 7393},
								name: "FencedBlockDelimiter",
							},
						},
						&labeledExpr{
							pos:   position{line: 182, col: 40, offset: 7414},
							label: "elements",
							expr: &oneOrMoreExpr{
								pos: position{line: 182, col: 49, offset: 7423},
								expr: &seqExpr{
									pos: position{line: 182, col: 50, offset: 7424},
									exprs: []interface{}{
										&zeroOrMoreExpr{
											pos: position{line: 182, col: 50, offset: 7424},
											expr: &ruleRefExpr{
												pos:  position{line: 182, col: 50, offset: 7424},
												name: "WS",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 182, col: 54, offset: 7428},
											name: "InlineElement",
										},
										&zeroOrMoreExpr{
											pos: position{line: 182, col: 68, offset: 7442},
											expr: &ruleRefExpr{
												pos:  position{line: 182, col: 68, offset: 7442},
												name: "WS",
											},
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 182, col: 74, offset: 7448},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "InlineElement",
			pos:  position{line: 186, col: 1, offset: 7526},
			expr: &choiceExpr{
				pos: position{line: 186, col: 18, offset: 7543},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 186, col: 18, offset: 7543},
						name: "InlineImage",
					},
					&ruleRefExpr{
						pos:  position{line: 186, col: 32, offset: 7557},
						name: "QuotedText",
					},
					&ruleRefExpr{
						pos:  position{line: 186, col: 45, offset: 7570},
						name: "ExternalLink",
					},
					&ruleRefExpr{
						pos:  position{line: 186, col: 60, offset: 7585},
						name: "DocumentAttributeSubstitution",
					},
					&ruleRefExpr{
						pos:  position{line: 186, col: 92, offset: 7617},
						name: "Word",
					},
				},
			},
		},
		{
			name: "QuotedText",
			pos:  position{line: 191, col: 1, offset: 7760},
			expr: &choiceExpr{
				pos: position{line: 191, col: 15, offset: 7774},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 191, col: 15, offset: 7774},
						name: "BoldText",
					},
					&ruleRefExpr{
						pos:  position{line: 191, col: 26, offset: 7785},
						name: "ItalicText",
					},
					&ruleRefExpr{
						pos:  position{line: 191, col: 39, offset: 7798},
						name: "MonospaceText",
					},
				},
			},
		},
		{
			name: "BoldText",
			pos:  position{line: 193, col: 1, offset: 7813},
			expr: &actionExpr{
				pos: position{line: 193, col: 13, offset: 7825},
				run: (*parser).callonBoldText1,
				expr: &seqExpr{
					pos: position{line: 193, col: 13, offset: 7825},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 193, col: 13, offset: 7825},
							val:        "*",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 193, col: 17, offset: 7829},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 193, col: 26, offset: 7838},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 193, col: 45, offset: 7857},
							val:        "*",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ItalicText",
			pos:  position{line: 197, col: 1, offset: 7934},
			expr: &actionExpr{
				pos: position{line: 197, col: 15, offset: 7948},
				run: (*parser).callonItalicText1,
				expr: &seqExpr{
					pos: position{line: 197, col: 15, offset: 7948},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 197, col: 15, offset: 7948},
							val:        "_",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 197, col: 19, offset: 7952},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 197, col: 28, offset: 7961},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 197, col: 47, offset: 7980},
							val:        "_",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "MonospaceText",
			pos:  position{line: 201, col: 1, offset: 8059},
			expr: &actionExpr{
				pos: position{line: 201, col: 18, offset: 8076},
				run: (*parser).callonMonospaceText1,
				expr: &seqExpr{
					pos: position{line: 201, col: 18, offset: 8076},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 201, col: 18, offset: 8076},
							val:        "`",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 201, col: 22, offset: 8080},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 201, col: 31, offset: 8089},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 201, col: 50, offset: 8108},
							val:        "`",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "QuotedTextContent",
			pos:  position{line: 205, col: 1, offset: 8190},
			expr: &seqExpr{
				pos: position{line: 205, col: 22, offset: 8211},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 205, col: 22, offset: 8211},
						name: "QuotedTextContentElement",
					},
					&zeroOrMoreExpr{
						pos: position{line: 205, col: 47, offset: 8236},
						expr: &seqExpr{
							pos: position{line: 205, col: 48, offset: 8237},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 205, col: 48, offset: 8237},
									expr: &ruleRefExpr{
										pos:  position{line: 205, col: 48, offset: 8237},
										name: "WS",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 205, col: 52, offset: 8241},
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
			pos:  position{line: 207, col: 1, offset: 8269},
			expr: &choiceExpr{
				pos: position{line: 207, col: 29, offset: 8297},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 207, col: 29, offset: 8297},
						name: "QuotedText",
					},
					&ruleRefExpr{
						pos:  position{line: 207, col: 42, offset: 8310},
						name: "QuotedTextContentWord",
					},
					&ruleRefExpr{
						pos:  position{line: 207, col: 66, offset: 8334},
						name: "InvalidQuotedTextContentWord",
					},
				},
			},
		},
		{
			name: "QuotedTextContentWord",
			pos:  position{line: 209, col: 1, offset: 8364},
			expr: &oneOrMoreExpr{
				pos: position{line: 209, col: 26, offset: 8389},
				expr: &seqExpr{
					pos: position{line: 209, col: 27, offset: 8390},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 209, col: 27, offset: 8390},
							expr: &ruleRefExpr{
								pos:  position{line: 209, col: 28, offset: 8391},
								name: "NEWLINE",
							},
						},
						&notExpr{
							pos: position{line: 209, col: 36, offset: 8399},
							expr: &ruleRefExpr{
								pos:  position{line: 209, col: 37, offset: 8400},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 209, col: 40, offset: 8403},
							expr: &litMatcher{
								pos:        position{line: 209, col: 41, offset: 8404},
								val:        "*",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 209, col: 45, offset: 8408},
							expr: &litMatcher{
								pos:        position{line: 209, col: 46, offset: 8409},
								val:        "_",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 209, col: 50, offset: 8413},
							expr: &litMatcher{
								pos:        position{line: 209, col: 51, offset: 8414},
								val:        "`",
								ignoreCase: false,
							},
						},
						&anyMatcher{
							line: 209, col: 55, offset: 8418,
						},
					},
				},
			},
		},
		{
			name: "InvalidQuotedTextContentWord",
			pos:  position{line: 210, col: 1, offset: 8460},
			expr: &oneOrMoreExpr{
				pos: position{line: 210, col: 33, offset: 8492},
				expr: &seqExpr{
					pos: position{line: 210, col: 34, offset: 8493},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 210, col: 34, offset: 8493},
							expr: &ruleRefExpr{
								pos:  position{line: 210, col: 35, offset: 8494},
								name: "NEWLINE",
							},
						},
						&notExpr{
							pos: position{line: 210, col: 43, offset: 8502},
							expr: &ruleRefExpr{
								pos:  position{line: 210, col: 44, offset: 8503},
								name: "WS",
							},
						},
						&anyMatcher{
							line: 210, col: 48, offset: 8507,
						},
					},
				},
			},
		},
		{
			name: "ExternalLink",
			pos:  position{line: 215, col: 1, offset: 8724},
			expr: &actionExpr{
				pos: position{line: 215, col: 17, offset: 8740},
				run: (*parser).callonExternalLink1,
				expr: &seqExpr{
					pos: position{line: 215, col: 17, offset: 8740},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 215, col: 17, offset: 8740},
							label: "url",
							expr: &seqExpr{
								pos: position{line: 215, col: 22, offset: 8745},
								exprs: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 215, col: 22, offset: 8745},
										name: "URL_SCHEME",
									},
									&ruleRefExpr{
										pos:  position{line: 215, col: 33, offset: 8756},
										name: "URL",
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 215, col: 38, offset: 8761},
							label: "text",
							expr: &zeroOrOneExpr{
								pos: position{line: 215, col: 43, offset: 8766},
								expr: &seqExpr{
									pos: position{line: 215, col: 44, offset: 8767},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 215, col: 44, offset: 8767},
											val:        "[",
											ignoreCase: false,
										},
										&zeroOrMoreExpr{
											pos: position{line: 215, col: 48, offset: 8771},
											expr: &ruleRefExpr{
												pos:  position{line: 215, col: 49, offset: 8772},
												name: "URL_TEXT",
											},
										},
										&litMatcher{
											pos:        position{line: 215, col: 60, offset: 8783},
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
			pos:  position{line: 225, col: 1, offset: 9062},
			expr: &actionExpr{
				pos: position{line: 225, col: 15, offset: 9076},
				run: (*parser).callonBlockImage1,
				expr: &seqExpr{
					pos: position{line: 225, col: 15, offset: 9076},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 225, col: 15, offset: 9076},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 225, col: 26, offset: 9087},
								expr: &ruleRefExpr{
									pos:  position{line: 225, col: 27, offset: 9088},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 225, col: 46, offset: 9107},
							label: "image",
							expr: &ruleRefExpr{
								pos:  position{line: 225, col: 52, offset: 9113},
								name: "BlockImageMacro",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 225, col: 69, offset: 9130},
							expr: &ruleRefExpr{
								pos:  position{line: 225, col: 69, offset: 9130},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 225, col: 73, offset: 9134},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "BlockImageMacro",
			pos:  position{line: 230, col: 1, offset: 9303},
			expr: &actionExpr{
				pos: position{line: 230, col: 20, offset: 9322},
				run: (*parser).callonBlockImageMacro1,
				expr: &seqExpr{
					pos: position{line: 230, col: 20, offset: 9322},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 230, col: 20, offset: 9322},
							val:        "image::",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 230, col: 30, offset: 9332},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 230, col: 36, offset: 9338},
								name: "URL",
							},
						},
						&litMatcher{
							pos:        position{line: 230, col: 41, offset: 9343},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 230, col: 45, offset: 9347},
							label: "attributes",
							expr: &zeroOrOneExpr{
								pos: position{line: 230, col: 57, offset: 9359},
								expr: &ruleRefExpr{
									pos:  position{line: 230, col: 57, offset: 9359},
									name: "URL_TEXT",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 230, col: 68, offset: 9370},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "InlineImage",
			pos:  position{line: 234, col: 1, offset: 9445},
			expr: &actionExpr{
				pos: position{line: 234, col: 16, offset: 9460},
				run: (*parser).callonInlineImage1,
				expr: &labeledExpr{
					pos:   position{line: 234, col: 16, offset: 9460},
					label: "image",
					expr: &ruleRefExpr{
						pos:  position{line: 234, col: 22, offset: 9466},
						name: "InlineImageMacro",
					},
				},
			},
		},
		{
			name: "InlineImageMacro",
			pos:  position{line: 239, col: 1, offset: 9621},
			expr: &actionExpr{
				pos: position{line: 239, col: 21, offset: 9641},
				run: (*parser).callonInlineImageMacro1,
				expr: &seqExpr{
					pos: position{line: 239, col: 21, offset: 9641},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 239, col: 21, offset: 9641},
							val:        "image:",
							ignoreCase: false,
						},
						&notExpr{
							pos: position{line: 239, col: 30, offset: 9650},
							expr: &litMatcher{
								pos:        position{line: 239, col: 31, offset: 9651},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 239, col: 35, offset: 9655},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 239, col: 41, offset: 9661},
								name: "URL",
							},
						},
						&litMatcher{
							pos:        position{line: 239, col: 46, offset: 9666},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 239, col: 50, offset: 9670},
							label: "attributes",
							expr: &zeroOrOneExpr{
								pos: position{line: 239, col: 62, offset: 9682},
								expr: &ruleRefExpr{
									pos:  position{line: 239, col: 62, offset: 9682},
									name: "URL_TEXT",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 239, col: 73, offset: 9693},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "DelimitedBlock",
			pos:  position{line: 246, col: 1, offset: 10031},
			expr: &ruleRefExpr{
				pos:  position{line: 246, col: 19, offset: 10049},
				name: "FencedBlock",
			},
		},
		{
			name: "FencedBlock",
			pos:  position{line: 248, col: 1, offset: 10063},
			expr: &actionExpr{
				pos: position{line: 248, col: 16, offset: 10078},
				run: (*parser).callonFencedBlock1,
				expr: &seqExpr{
					pos: position{line: 248, col: 16, offset: 10078},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 248, col: 16, offset: 10078},
							name: "FencedBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 248, col: 37, offset: 10099},
							expr: &ruleRefExpr{
								pos:  position{line: 248, col: 37, offset: 10099},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 248, col: 41, offset: 10103},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 248, col: 49, offset: 10111},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 248, col: 58, offset: 10120},
								name: "FencedBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 248, col: 78, offset: 10140},
							name: "FencedBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 248, col: 99, offset: 10161},
							expr: &ruleRefExpr{
								pos:  position{line: 248, col: 99, offset: 10161},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 248, col: 103, offset: 10165},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "FencedBlockDelimiter",
			pos:  position{line: 252, col: 1, offset: 10253},
			expr: &litMatcher{
				pos:        position{line: 252, col: 25, offset: 10277},
				val:        "```",
				ignoreCase: false,
			},
		},
		{
			name: "FencedBlockContent",
			pos:  position{line: 254, col: 1, offset: 10284},
			expr: &labeledExpr{
				pos:   position{line: 254, col: 23, offset: 10306},
				label: "content",
				expr: &zeroOrMoreExpr{
					pos: position{line: 254, col: 31, offset: 10314},
					expr: &seqExpr{
						pos: position{line: 254, col: 32, offset: 10315},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 254, col: 32, offset: 10315},
								expr: &ruleRefExpr{
									pos:  position{line: 254, col: 33, offset: 10316},
									name: "FencedBlockDelimiter",
								},
							},
							&anyMatcher{
								line: 254, col: 54, offset: 10337,
							},
						},
					},
				},
			},
		},
		{
			name: "LiteralBlock",
			pos:  position{line: 259, col: 1, offset: 10610},
			expr: &choiceExpr{
				pos: position{line: 259, col: 17, offset: 10626},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 259, col: 17, offset: 10626},
						name: "ParagraphWithSpaces",
					},
					&ruleRefExpr{
						pos:  position{line: 259, col: 39, offset: 10648},
						name: "ParagraphWithLiteralBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 259, col: 76, offset: 10685},
						name: "ParagraphWithLiteralAttribute",
					},
				},
			},
		},
		{
			name: "ParagraphWithSpaces",
			pos:  position{line: 262, col: 1, offset: 10780},
			expr: &actionExpr{
				pos: position{line: 262, col: 24, offset: 10803},
				run: (*parser).callonParagraphWithSpaces1,
				expr: &seqExpr{
					pos: position{line: 262, col: 24, offset: 10803},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 262, col: 24, offset: 10803},
							label: "spaces",
							expr: &oneOrMoreExpr{
								pos: position{line: 262, col: 32, offset: 10811},
								expr: &ruleRefExpr{
									pos:  position{line: 262, col: 32, offset: 10811},
									name: "WS",
								},
							},
						},
						&notExpr{
							pos: position{line: 262, col: 37, offset: 10816},
							expr: &ruleRefExpr{
								pos:  position{line: 262, col: 38, offset: 10817},
								name: "NEWLINE",
							},
						},
						&labeledExpr{
							pos:   position{line: 262, col: 46, offset: 10825},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 262, col: 55, offset: 10834},
								name: "LiteralBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 262, col: 76, offset: 10855},
							name: "EndOfLiteralBlock",
						},
					},
				},
			},
		},
		{
			name: "LiteralBlockContent",
			pos:  position{line: 268, col: 1, offset: 11112},
			expr: &actionExpr{
				pos: position{line: 268, col: 24, offset: 11135},
				run: (*parser).callonLiteralBlockContent1,
				expr: &labeledExpr{
					pos:   position{line: 268, col: 24, offset: 11135},
					label: "content",
					expr: &oneOrMoreExpr{
						pos: position{line: 268, col: 32, offset: 11143},
						expr: &seqExpr{
							pos: position{line: 268, col: 33, offset: 11144},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 268, col: 33, offset: 11144},
									expr: &seqExpr{
										pos: position{line: 268, col: 35, offset: 11146},
										exprs: []interface{}{
											&ruleRefExpr{
												pos:  position{line: 268, col: 35, offset: 11146},
												name: "NEWLINE",
											},
											&ruleRefExpr{
												pos:  position{line: 268, col: 43, offset: 11154},
												name: "BlankLine",
											},
										},
									},
								},
								&anyMatcher{
									line: 268, col: 54, offset: 11165,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "EndOfLiteralBlock",
			pos:  position{line: 273, col: 1, offset: 11250},
			expr: &choiceExpr{
				pos: position{line: 273, col: 22, offset: 11271},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 273, col: 22, offset: 11271},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 273, col: 22, offset: 11271},
								name: "NEWLINE",
							},
							&ruleRefExpr{
								pos:  position{line: 273, col: 30, offset: 11279},
								name: "BlankLine",
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 273, col: 42, offset: 11291},
						name: "NEWLINE",
					},
					&ruleRefExpr{
						pos:  position{line: 273, col: 52, offset: 11301},
						name: "EOF",
					},
				},
			},
		},
		{
			name: "ParagraphWithLiteralBlockDelimiter",
			pos:  position{line: 276, col: 1, offset: 11361},
			expr: &actionExpr{
				pos: position{line: 276, col: 39, offset: 11399},
				run: (*parser).callonParagraphWithLiteralBlockDelimiter1,
				expr: &seqExpr{
					pos: position{line: 276, col: 39, offset: 11399},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 276, col: 39, offset: 11399},
							name: "LiteralBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 276, col: 61, offset: 11421},
							expr: &ruleRefExpr{
								pos:  position{line: 276, col: 61, offset: 11421},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 276, col: 65, offset: 11425},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 276, col: 73, offset: 11433},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 276, col: 81, offset: 11441},
								expr: &seqExpr{
									pos: position{line: 276, col: 82, offset: 11442},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 276, col: 82, offset: 11442},
											expr: &ruleRefExpr{
												pos:  position{line: 276, col: 83, offset: 11443},
												name: "LiteralBlockDelimiter",
											},
										},
										&anyMatcher{
											line: 276, col: 105, offset: 11465,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 276, col: 109, offset: 11469},
							name: "LiteralBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 276, col: 131, offset: 11491},
							expr: &ruleRefExpr{
								pos:  position{line: 276, col: 131, offset: 11491},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 276, col: 135, offset: 11495},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "LiteralBlockDelimiter",
			pos:  position{line: 280, col: 1, offset: 11579},
			expr: &litMatcher{
				pos:        position{line: 280, col: 26, offset: 11604},
				val:        "....",
				ignoreCase: false,
			},
		},
		{
			name: "ParagraphWithLiteralAttribute",
			pos:  position{line: 283, col: 1, offset: 11666},
			expr: &actionExpr{
				pos: position{line: 283, col: 34, offset: 11699},
				run: (*parser).callonParagraphWithLiteralAttribute1,
				expr: &seqExpr{
					pos: position{line: 283, col: 34, offset: 11699},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 283, col: 34, offset: 11699},
							val:        "[literal]",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 283, col: 46, offset: 11711},
							expr: &ruleRefExpr{
								pos:  position{line: 283, col: 46, offset: 11711},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 283, col: 50, offset: 11715},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 283, col: 58, offset: 11723},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 283, col: 67, offset: 11732},
								name: "LiteralBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 283, col: 88, offset: 11753},
							name: "EndOfLiteralBlock",
						},
					},
				},
			},
		},
		{
			name: "ElementAttribute",
			pos:  position{line: 290, col: 1, offset: 11965},
			expr: &labeledExpr{
				pos:   position{line: 290, col: 21, offset: 11985},
				label: "meta",
				expr: &choiceExpr{
					pos: position{line: 290, col: 27, offset: 11991},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 290, col: 27, offset: 11991},
							name: "ElementLink",
						},
						&ruleRefExpr{
							pos:  position{line: 290, col: 41, offset: 12005},
							name: "ElementID",
						},
						&ruleRefExpr{
							pos:  position{line: 290, col: 53, offset: 12017},
							name: "ElementTitle",
						},
					},
				},
			},
		},
		{
			name: "ElementLink",
			pos:  position{line: 293, col: 1, offset: 12088},
			expr: &actionExpr{
				pos: position{line: 293, col: 16, offset: 12103},
				run: (*parser).callonElementLink1,
				expr: &seqExpr{
					pos: position{line: 293, col: 16, offset: 12103},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 293, col: 16, offset: 12103},
							val:        "[",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 293, col: 20, offset: 12107},
							expr: &ruleRefExpr{
								pos:  position{line: 293, col: 20, offset: 12107},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 293, col: 24, offset: 12111},
							val:        "link",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 293, col: 31, offset: 12118},
							expr: &ruleRefExpr{
								pos:  position{line: 293, col: 31, offset: 12118},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 293, col: 35, offset: 12122},
							val:        "=",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 293, col: 39, offset: 12126},
							expr: &ruleRefExpr{
								pos:  position{line: 293, col: 39, offset: 12126},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 293, col: 43, offset: 12130},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 293, col: 48, offset: 12135},
								name: "URL",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 293, col: 52, offset: 12139},
							expr: &ruleRefExpr{
								pos:  position{line: 293, col: 52, offset: 12139},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 293, col: 56, offset: 12143},
							val:        "]",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 293, col: 60, offset: 12147},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ElementID",
			pos:  position{line: 298, col: 1, offset: 12257},
			expr: &actionExpr{
				pos: position{line: 298, col: 14, offset: 12270},
				run: (*parser).callonElementID1,
				expr: &seqExpr{
					pos: position{line: 298, col: 14, offset: 12270},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 298, col: 14, offset: 12270},
							val:        "[",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 298, col: 18, offset: 12274},
							expr: &ruleRefExpr{
								pos:  position{line: 298, col: 18, offset: 12274},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 298, col: 22, offset: 12278},
							val:        "#",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 298, col: 26, offset: 12282},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 298, col: 30, offset: 12286},
								name: "ID",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 298, col: 34, offset: 12290},
							expr: &ruleRefExpr{
								pos:  position{line: 298, col: 34, offset: 12290},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 298, col: 38, offset: 12294},
							val:        "]",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 298, col: 42, offset: 12298},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ElementTitle",
			pos:  position{line: 304, col: 1, offset: 12492},
			expr: &actionExpr{
				pos: position{line: 304, col: 17, offset: 12508},
				run: (*parser).callonElementTitle1,
				expr: &seqExpr{
					pos: position{line: 304, col: 17, offset: 12508},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 304, col: 17, offset: 12508},
							val:        ".",
							ignoreCase: false,
						},
						&notExpr{
							pos: position{line: 304, col: 21, offset: 12512},
							expr: &litMatcher{
								pos:        position{line: 304, col: 22, offset: 12513},
								val:        ".",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 304, col: 26, offset: 12517},
							expr: &ruleRefExpr{
								pos:  position{line: 304, col: 27, offset: 12518},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 304, col: 30, offset: 12521},
							label: "title",
							expr: &oneOrMoreExpr{
								pos: position{line: 304, col: 36, offset: 12527},
								expr: &seqExpr{
									pos: position{line: 304, col: 37, offset: 12528},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 304, col: 37, offset: 12528},
											expr: &ruleRefExpr{
												pos:  position{line: 304, col: 38, offset: 12529},
												name: "NEWLINE",
											},
										},
										&anyMatcher{
											line: 304, col: 46, offset: 12537,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 304, col: 50, offset: 12541},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "Word",
			pos:  position{line: 311, col: 1, offset: 12712},
			expr: &actionExpr{
				pos: position{line: 311, col: 9, offset: 12720},
				run: (*parser).callonWord1,
				expr: &oneOrMoreExpr{
					pos: position{line: 311, col: 9, offset: 12720},
					expr: &seqExpr{
						pos: position{line: 311, col: 10, offset: 12721},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 311, col: 10, offset: 12721},
								expr: &ruleRefExpr{
									pos:  position{line: 311, col: 11, offset: 12722},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 311, col: 19, offset: 12730},
								expr: &ruleRefExpr{
									pos:  position{line: 311, col: 20, offset: 12731},
									name: "WS",
								},
							},
							&anyMatcher{
								line: 311, col: 23, offset: 12734,
							},
						},
					},
				},
			},
		},
		{
			name: "BlankLine",
			pos:  position{line: 315, col: 1, offset: 12774},
			expr: &actionExpr{
				pos: position{line: 315, col: 14, offset: 12787},
				run: (*parser).callonBlankLine1,
				expr: &seqExpr{
					pos: position{line: 315, col: 14, offset: 12787},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 315, col: 14, offset: 12787},
							expr: &ruleRefExpr{
								pos:  position{line: 315, col: 15, offset: 12788},
								name: "EOF",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 315, col: 19, offset: 12792},
							expr: &ruleRefExpr{
								pos:  position{line: 315, col: 19, offset: 12792},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 315, col: 23, offset: 12796},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "URL",
			pos:  position{line: 319, col: 1, offset: 12837},
			expr: &actionExpr{
				pos: position{line: 319, col: 8, offset: 12844},
				run: (*parser).callonURL1,
				expr: &oneOrMoreExpr{
					pos: position{line: 319, col: 8, offset: 12844},
					expr: &seqExpr{
						pos: position{line: 319, col: 9, offset: 12845},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 319, col: 9, offset: 12845},
								expr: &ruleRefExpr{
									pos:  position{line: 319, col: 10, offset: 12846},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 319, col: 18, offset: 12854},
								expr: &ruleRefExpr{
									pos:  position{line: 319, col: 19, offset: 12855},
									name: "WS",
								},
							},
							&notExpr{
								pos: position{line: 319, col: 22, offset: 12858},
								expr: &litMatcher{
									pos:        position{line: 319, col: 23, offset: 12859},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 319, col: 27, offset: 12863},
								expr: &litMatcher{
									pos:        position{line: 319, col: 28, offset: 12864},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 319, col: 32, offset: 12868,
							},
						},
					},
				},
			},
		},
		{
			name: "ID",
			pos:  position{line: 323, col: 1, offset: 12908},
			expr: &actionExpr{
				pos: position{line: 323, col: 7, offset: 12914},
				run: (*parser).callonID1,
				expr: &oneOrMoreExpr{
					pos: position{line: 323, col: 7, offset: 12914},
					expr: &seqExpr{
						pos: position{line: 323, col: 8, offset: 12915},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 323, col: 8, offset: 12915},
								expr: &ruleRefExpr{
									pos:  position{line: 323, col: 9, offset: 12916},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 323, col: 17, offset: 12924},
								expr: &ruleRefExpr{
									pos:  position{line: 323, col: 18, offset: 12925},
									name: "WS",
								},
							},
							&notExpr{
								pos: position{line: 323, col: 21, offset: 12928},
								expr: &litMatcher{
									pos:        position{line: 323, col: 22, offset: 12929},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 323, col: 26, offset: 12933},
								expr: &litMatcher{
									pos:        position{line: 323, col: 27, offset: 12934},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 323, col: 31, offset: 12938,
							},
						},
					},
				},
			},
		},
		{
			name: "URL_TEXT",
			pos:  position{line: 327, col: 1, offset: 12978},
			expr: &actionExpr{
				pos: position{line: 327, col: 13, offset: 12990},
				run: (*parser).callonURL_TEXT1,
				expr: &oneOrMoreExpr{
					pos: position{line: 327, col: 13, offset: 12990},
					expr: &seqExpr{
						pos: position{line: 327, col: 14, offset: 12991},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 327, col: 14, offset: 12991},
								expr: &ruleRefExpr{
									pos:  position{line: 327, col: 15, offset: 12992},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 327, col: 23, offset: 13000},
								expr: &litMatcher{
									pos:        position{line: 327, col: 24, offset: 13001},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 327, col: 28, offset: 13005},
								expr: &litMatcher{
									pos:        position{line: 327, col: 29, offset: 13006},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 327, col: 33, offset: 13010,
							},
						},
					},
				},
			},
		},
		{
			name: "URL_SCHEME",
			pos:  position{line: 331, col: 1, offset: 13050},
			expr: &choiceExpr{
				pos: position{line: 331, col: 15, offset: 13064},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 331, col: 15, offset: 13064},
						val:        "http://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 331, col: 27, offset: 13076},
						val:        "https://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 331, col: 40, offset: 13089},
						val:        "ftp://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 331, col: 51, offset: 13100},
						val:        "irc://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 331, col: 62, offset: 13111},
						val:        "mailto:",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "DIGIT",
			pos:  position{line: 333, col: 1, offset: 13122},
			expr: &charClassMatcher{
				pos:        position{line: 333, col: 13, offset: 13134},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NEWLINE",
			pos:  position{line: 335, col: 1, offset: 13141},
			expr: &choiceExpr{
				pos: position{line: 335, col: 13, offset: 13153},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 335, col: 13, offset: 13153},
						val:        "\r\n",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 335, col: 22, offset: 13162},
						val:        "\r",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 335, col: 29, offset: 13169},
						val:        "\n",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "WS",
			pos:  position{line: 337, col: 1, offset: 13175},
			expr: &choiceExpr{
				pos: position{line: 337, col: 13, offset: 13187},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 337, col: 13, offset: 13187},
						val:        " ",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 337, col: 19, offset: 13193},
						run: (*parser).callonWS3,
						expr: &litMatcher{
							pos:        position{line: 337, col: 19, offset: 13193},
							val:        "\t",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 341, col: 1, offset: 13238},
			expr: &notExpr{
				pos: position{line: 341, col: 13, offset: 13250},
				expr: &anyMatcher{
					line: 341, col: 14, offset: 13251,
				},
			},
		},
		{
			name: "EOL",
			pos:  position{line: 343, col: 1, offset: 13254},
			expr: &choiceExpr{
				pos: position{line: 343, col: 13, offset: 13266},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 343, col: 13, offset: 13266},
						name: "NEWLINE",
					},
					&ruleRefExpr{
						pos:  position{line: 343, col: 23, offset: 13276},
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

func (c *current) onFencedBlock1(content interface{}) (interface{}, error) {
	return types.NewDelimitedBlock(types.FencedBlock, content.([]interface{}))
}

func (p *parser) callonFencedBlock1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFencedBlock1(stack["content"])
}

func (c *current) onParagraphWithSpaces1(spaces, content interface{}) (interface{}, error) {
	// fmt.Printf("matching LiteralBlock with raw content=`%v`\n", content)
	return types.NewLiteralBlock(spaces.([]interface{}), content.([]interface{}))
}

func (p *parser) callonParagraphWithSpaces1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onParagraphWithSpaces1(stack["spaces"], stack["content"])
}

func (c *current) onLiteralBlockContent1(content interface{}) (interface{}, error) {

	return content, nil
}

func (p *parser) callonLiteralBlockContent1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLiteralBlockContent1(stack["content"])
}

func (c *current) onParagraphWithLiteralBlockDelimiter1(content interface{}) (interface{}, error) {
	return types.NewLiteralBlock([]interface{}{}, content.([]interface{}))
}

func (p *parser) callonParagraphWithLiteralBlockDelimiter1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onParagraphWithLiteralBlockDelimiter1(stack["content"])
}

func (c *current) onParagraphWithLiteralAttribute1(content interface{}) (interface{}, error) {
	return types.NewLiteralBlock([]interface{}{}, content.([]interface{}))
}

func (p *parser) callonParagraphWithLiteralAttribute1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onParagraphWithLiteralAttribute1(stack["content"])
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
