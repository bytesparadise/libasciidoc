package parser

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strconv"
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
							label: "frontMatter",
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
							label: "documentHeader",
							expr: &zeroOrOneExpr{
								pos: position{line: 16, col: 56, offset: 511},
								expr: &ruleRefExpr{
									pos:  position{line: 16, col: 56, offset: 511},
									name: "DocumentHeader",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 16, col: 73, offset: 528},
							label: "blocks",
							expr: &ruleRefExpr{
								pos:  position{line: 16, col: 81, offset: 536},
								name: "DocumentBlocks",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 16, col: 97, offset: 552},
							name: "EOF",
						},
					},
				},
			},
		},
		{
			name: "DocumentBlocks",
			pos:  position{line: 20, col: 1, offset: 640},
			expr: &choiceExpr{
				pos: position{line: 20, col: 19, offset: 658},
				alternatives: []interface{}{
					&labeledExpr{
						pos:   position{line: 20, col: 19, offset: 658},
						label: "content",
						expr: &seqExpr{
							pos: position{line: 20, col: 28, offset: 667},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 20, col: 28, offset: 667},
									name: "Preamble",
								},
								&oneOrMoreExpr{
									pos: position{line: 20, col: 37, offset: 676},
									expr: &ruleRefExpr{
										pos:  position{line: 20, col: 37, offset: 676},
										name: "Section",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 20, col: 49, offset: 688},
						run: (*parser).callonDocumentBlocks7,
						expr: &labeledExpr{
							pos:   position{line: 20, col: 49, offset: 688},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 20, col: 58, offset: 697},
								expr: &ruleRefExpr{
									pos:  position{line: 20, col: 58, offset: 697},
									name: "StandaloneBlock",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "StandaloneBlock",
			pos:  position{line: 24, col: 1, offset: 744},
			expr: &choiceExpr{
				pos: position{line: 24, col: 20, offset: 763},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 24, col: 20, offset: 763},
						name: "DocumentAttributeDeclaration",
					},
					&ruleRefExpr{
						pos:  position{line: 24, col: 51, offset: 794},
						name: "DocumentAttributeReset",
					},
					&ruleRefExpr{
						pos:  position{line: 24, col: 76, offset: 819},
						name: "List",
					},
					&ruleRefExpr{
						pos:  position{line: 24, col: 83, offset: 826},
						name: "BlockImage",
					},
					&ruleRefExpr{
						pos:  position{line: 24, col: 96, offset: 839},
						name: "LiteralBlock",
					},
					&ruleRefExpr{
						pos:  position{line: 24, col: 111, offset: 854},
						name: "DelimitedBlock",
					},
					&ruleRefExpr{
						pos:  position{line: 24, col: 128, offset: 871},
						name: "Paragraph",
					},
					&seqExpr{
						pos: position{line: 24, col: 141, offset: 884},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 24, col: 141, offset: 884},
								name: "ElementAttribute",
							},
							&ruleRefExpr{
								pos:  position{line: 24, col: 158, offset: 901},
								name: "EOL",
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 24, col: 165, offset: 908},
						name: "BlankLine",
					},
				},
			},
		},
		{
			name: "Preamble",
			pos:  position{line: 26, col: 1, offset: 963},
			expr: &actionExpr{
				pos: position{line: 26, col: 13, offset: 975},
				run: (*parser).callonPreamble1,
				expr: &labeledExpr{
					pos:   position{line: 26, col: 13, offset: 975},
					label: "elements",
					expr: &zeroOrMoreExpr{
						pos: position{line: 26, col: 23, offset: 985},
						expr: &ruleRefExpr{
							pos:  position{line: 26, col: 23, offset: 985},
							name: "StandaloneBlock",
						},
					},
				},
			},
		},
		{
			name: "FrontMatter",
			pos:  position{line: 33, col: 1, offset: 1171},
			expr: &ruleRefExpr{
				pos:  position{line: 33, col: 16, offset: 1186},
				name: "YamlFrontMatter",
			},
		},
		{
			name: "FrontMatter",
			pos:  position{line: 35, col: 1, offset: 1204},
			expr: &actionExpr{
				pos: position{line: 35, col: 16, offset: 1219},
				run: (*parser).callonFrontMatter1,
				expr: &seqExpr{
					pos: position{line: 35, col: 16, offset: 1219},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 35, col: 16, offset: 1219},
							name: "YamlFrontMatterToken",
						},
						&labeledExpr{
							pos:   position{line: 35, col: 37, offset: 1240},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 35, col: 46, offset: 1249},
								name: "YamlFrontMatterContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 35, col: 70, offset: 1273},
							name: "YamlFrontMatterToken",
						},
					},
				},
			},
		},
		{
			name: "YamlFrontMatterToken",
			pos:  position{line: 39, col: 1, offset: 1353},
			expr: &seqExpr{
				pos: position{line: 39, col: 26, offset: 1378},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 39, col: 26, offset: 1378},
						val:        "---",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 39, col: 32, offset: 1384},
						name: "EOL",
					},
				},
			},
		},
		{
			name: "YamlFrontMatterContent",
			pos:  position{line: 41, col: 1, offset: 1389},
			expr: &actionExpr{
				pos: position{line: 41, col: 27, offset: 1415},
				run: (*parser).callonYamlFrontMatterContent1,
				expr: &zeroOrMoreExpr{
					pos: position{line: 41, col: 27, offset: 1415},
					expr: &seqExpr{
						pos: position{line: 41, col: 28, offset: 1416},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 41, col: 28, offset: 1416},
								expr: &ruleRefExpr{
									pos:  position{line: 41, col: 29, offset: 1417},
									name: "YamlFrontMatterToken",
								},
							},
							&anyMatcher{
								line: 41, col: 50, offset: 1438,
							},
						},
					},
				},
			},
		},
		{
			name: "DocumentHeader",
			pos:  position{line: 49, col: 1, offset: 1662},
			expr: &actionExpr{
				pos: position{line: 49, col: 19, offset: 1680},
				run: (*parser).callonDocumentHeader1,
				expr: &seqExpr{
					pos: position{line: 49, col: 19, offset: 1680},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 49, col: 19, offset: 1680},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 49, col: 27, offset: 1688},
								name: "DocumentTitle",
							},
						},
						&labeledExpr{
							pos:   position{line: 49, col: 42, offset: 1703},
							label: "authors",
							expr: &zeroOrOneExpr{
								pos: position{line: 49, col: 51, offset: 1712},
								expr: &ruleRefExpr{
									pos:  position{line: 49, col: 51, offset: 1712},
									name: "DocumentAuthors",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 49, col: 69, offset: 1730},
							label: "revision",
							expr: &zeroOrOneExpr{
								pos: position{line: 49, col: 79, offset: 1740},
								expr: &ruleRefExpr{
									pos:  position{line: 49, col: 79, offset: 1740},
									name: "DocumentRevision",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 49, col: 98, offset: 1759},
							label: "otherAttributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 49, col: 115, offset: 1776},
								expr: &ruleRefExpr{
									pos:  position{line: 49, col: 115, offset: 1776},
									name: "DocumentAttributeDeclaration",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "DocumentTitle",
			pos:  position{line: 53, col: 1, offset: 1907},
			expr: &actionExpr{
				pos: position{line: 53, col: 18, offset: 1924},
				run: (*parser).callonDocumentTitle1,
				expr: &seqExpr{
					pos: position{line: 53, col: 18, offset: 1924},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 53, col: 18, offset: 1924},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 53, col: 29, offset: 1935},
								expr: &ruleRefExpr{
									pos:  position{line: 53, col: 30, offset: 1936},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 53, col: 49, offset: 1955},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 53, col: 56, offset: 1962},
								val:        "=",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 53, col: 61, offset: 1967},
							expr: &ruleRefExpr{
								pos:  position{line: 53, col: 61, offset: 1967},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 53, col: 65, offset: 1971},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 53, col: 73, offset: 1979},
								name: "InlineContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 53, col: 87, offset: 1993},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthors",
			pos:  position{line: 57, col: 1, offset: 2097},
			expr: &choiceExpr{
				pos: position{line: 57, col: 20, offset: 2116},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 57, col: 20, offset: 2116},
						name: "DocumentAuthorsInlineForm",
					},
					&ruleRefExpr{
						pos:  position{line: 57, col: 48, offset: 2144},
						name: "DocumentAuthorsAttributeForm",
					},
				},
			},
		},
		{
			name: "DocumentAuthorsInlineForm",
			pos:  position{line: 59, col: 1, offset: 2174},
			expr: &actionExpr{
				pos: position{line: 59, col: 30, offset: 2203},
				run: (*parser).callonDocumentAuthorsInlineForm1,
				expr: &seqExpr{
					pos: position{line: 59, col: 30, offset: 2203},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 59, col: 30, offset: 2203},
							expr: &ruleRefExpr{
								pos:  position{line: 59, col: 30, offset: 2203},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 59, col: 34, offset: 2207},
							expr: &litMatcher{
								pos:        position{line: 59, col: 35, offset: 2208},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 59, col: 39, offset: 2212},
							label: "authors",
							expr: &oneOrMoreExpr{
								pos: position{line: 59, col: 48, offset: 2221},
								expr: &ruleRefExpr{
									pos:  position{line: 59, col: 48, offset: 2221},
									name: "DocumentAuthor",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 59, col: 65, offset: 2238},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthorsAttributeForm",
			pos:  position{line: 63, col: 1, offset: 2308},
			expr: &actionExpr{
				pos: position{line: 63, col: 33, offset: 2340},
				run: (*parser).callonDocumentAuthorsAttributeForm1,
				expr: &seqExpr{
					pos: position{line: 63, col: 33, offset: 2340},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 63, col: 33, offset: 2340},
							expr: &ruleRefExpr{
								pos:  position{line: 63, col: 33, offset: 2340},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 63, col: 37, offset: 2344},
							val:        ":author:",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 63, col: 48, offset: 2355},
							label: "author",
							expr: &ruleRefExpr{
								pos:  position{line: 63, col: 56, offset: 2363},
								name: "DocumentAuthor",
							},
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthor",
			pos:  position{line: 67, col: 1, offset: 2456},
			expr: &actionExpr{
				pos: position{line: 67, col: 19, offset: 2474},
				run: (*parser).callonDocumentAuthor1,
				expr: &seqExpr{
					pos: position{line: 67, col: 19, offset: 2474},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 67, col: 19, offset: 2474},
							expr: &ruleRefExpr{
								pos:  position{line: 67, col: 19, offset: 2474},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 67, col: 23, offset: 2478},
							label: "namePart1",
							expr: &ruleRefExpr{
								pos:  position{line: 67, col: 34, offset: 2489},
								name: "DocumentAuthorNamePart",
							},
						},
						&labeledExpr{
							pos:   position{line: 67, col: 58, offset: 2513},
							label: "namePart2",
							expr: &zeroOrOneExpr{
								pos: position{line: 67, col: 68, offset: 2523},
								expr: &ruleRefExpr{
									pos:  position{line: 67, col: 69, offset: 2524},
									name: "DocumentAuthorNamePart",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 67, col: 94, offset: 2549},
							label: "namePart3",
							expr: &zeroOrOneExpr{
								pos: position{line: 67, col: 104, offset: 2559},
								expr: &ruleRefExpr{
									pos:  position{line: 67, col: 105, offset: 2560},
									name: "DocumentAuthorNamePart",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 67, col: 130, offset: 2585},
							label: "email",
							expr: &zeroOrOneExpr{
								pos: position{line: 67, col: 136, offset: 2591},
								expr: &ruleRefExpr{
									pos:  position{line: 67, col: 137, offset: 2592},
									name: "DocumentAuthorEmail",
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 67, col: 159, offset: 2614},
							expr: &ruleRefExpr{
								pos:  position{line: 67, col: 159, offset: 2614},
								name: "WS",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 67, col: 163, offset: 2618},
							expr: &litMatcher{
								pos:        position{line: 67, col: 163, offset: 2618},
								val:        ";",
								ignoreCase: false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 67, col: 168, offset: 2623},
							expr: &ruleRefExpr{
								pos:  position{line: 67, col: 168, offset: 2623},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthorNamePart",
			pos:  position{line: 72, col: 1, offset: 2788},
			expr: &seqExpr{
				pos: position{line: 72, col: 27, offset: 2814},
				exprs: []interface{}{
					&notExpr{
						pos: position{line: 72, col: 27, offset: 2814},
						expr: &litMatcher{
							pos:        position{line: 72, col: 28, offset: 2815},
							val:        "<",
							ignoreCase: false,
						},
					},
					&notExpr{
						pos: position{line: 72, col: 32, offset: 2819},
						expr: &litMatcher{
							pos:        position{line: 72, col: 33, offset: 2820},
							val:        ";",
							ignoreCase: false,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 72, col: 37, offset: 2824},
						name: "Characters",
					},
					&zeroOrMoreExpr{
						pos: position{line: 72, col: 48, offset: 2835},
						expr: &ruleRefExpr{
							pos:  position{line: 72, col: 48, offset: 2835},
							name: "WS",
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthorEmail",
			pos:  position{line: 74, col: 1, offset: 2840},
			expr: &seqExpr{
				pos: position{line: 74, col: 24, offset: 2863},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 74, col: 24, offset: 2863},
						val:        "<",
						ignoreCase: false,
					},
					&labeledExpr{
						pos:   position{line: 74, col: 28, offset: 2867},
						label: "email",
						expr: &oneOrMoreExpr{
							pos: position{line: 74, col: 34, offset: 2873},
							expr: &seqExpr{
								pos: position{line: 74, col: 35, offset: 2874},
								exprs: []interface{}{
									&notExpr{
										pos: position{line: 74, col: 35, offset: 2874},
										expr: &litMatcher{
											pos:        position{line: 74, col: 36, offset: 2875},
											val:        ">",
											ignoreCase: false,
										},
									},
									&notExpr{
										pos: position{line: 74, col: 40, offset: 2879},
										expr: &ruleRefExpr{
											pos:  position{line: 74, col: 41, offset: 2880},
											name: "EOL",
										},
									},
									&anyMatcher{
										line: 74, col: 45, offset: 2884,
									},
								},
							},
						},
					},
					&litMatcher{
						pos:        position{line: 74, col: 49, offset: 2888},
						val:        ">",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "DocumentRevision",
			pos:  position{line: 78, col: 1, offset: 3024},
			expr: &actionExpr{
				pos: position{line: 78, col: 21, offset: 3044},
				run: (*parser).callonDocumentRevision1,
				expr: &seqExpr{
					pos: position{line: 78, col: 21, offset: 3044},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 78, col: 21, offset: 3044},
							expr: &ruleRefExpr{
								pos:  position{line: 78, col: 21, offset: 3044},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 78, col: 25, offset: 3048},
							expr: &litMatcher{
								pos:        position{line: 78, col: 26, offset: 3049},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 78, col: 30, offset: 3053},
							label: "revnumber",
							expr: &zeroOrOneExpr{
								pos: position{line: 78, col: 40, offset: 3063},
								expr: &ruleRefExpr{
									pos:  position{line: 78, col: 41, offset: 3064},
									name: "DocumentRevisionNumber",
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 78, col: 66, offset: 3089},
							expr: &litMatcher{
								pos:        position{line: 78, col: 66, offset: 3089},
								val:        ",",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 78, col: 71, offset: 3094},
							label: "revdate",
							expr: &zeroOrOneExpr{
								pos: position{line: 78, col: 79, offset: 3102},
								expr: &ruleRefExpr{
									pos:  position{line: 78, col: 80, offset: 3103},
									name: "DocumentRevisionDate",
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 78, col: 103, offset: 3126},
							expr: &litMatcher{
								pos:        position{line: 78, col: 103, offset: 3126},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 78, col: 108, offset: 3131},
							label: "revremark",
							expr: &zeroOrOneExpr{
								pos: position{line: 78, col: 118, offset: 3141},
								expr: &ruleRefExpr{
									pos:  position{line: 78, col: 119, offset: 3142},
									name: "DocumentRevisionRemark",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 78, col: 144, offset: 3167},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentRevisionNumber",
			pos:  position{line: 83, col: 1, offset: 3340},
			expr: &choiceExpr{
				pos: position{line: 83, col: 27, offset: 3366},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 83, col: 27, offset: 3366},
						exprs: []interface{}{
							&litMatcher{
								pos:        position{line: 83, col: 27, offset: 3366},
								val:        "v",
								ignoreCase: true,
							},
							&ruleRefExpr{
								pos:  position{line: 83, col: 32, offset: 3371},
								name: "DIGIT",
							},
							&zeroOrMoreExpr{
								pos: position{line: 83, col: 39, offset: 3378},
								expr: &seqExpr{
									pos: position{line: 83, col: 40, offset: 3379},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 83, col: 40, offset: 3379},
											expr: &ruleRefExpr{
												pos:  position{line: 83, col: 41, offset: 3380},
												name: "EOL",
											},
										},
										&notExpr{
											pos: position{line: 83, col: 45, offset: 3384},
											expr: &litMatcher{
												pos:        position{line: 83, col: 46, offset: 3385},
												val:        ",",
												ignoreCase: false,
											},
										},
										&notExpr{
											pos: position{line: 83, col: 50, offset: 3389},
											expr: &litMatcher{
												pos:        position{line: 83, col: 51, offset: 3390},
												val:        ":",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 83, col: 55, offset: 3394,
										},
									},
								},
							},
						},
					},
					&seqExpr{
						pos: position{line: 83, col: 61, offset: 3400},
						exprs: []interface{}{
							&zeroOrOneExpr{
								pos: position{line: 83, col: 61, offset: 3400},
								expr: &litMatcher{
									pos:        position{line: 83, col: 61, offset: 3400},
									val:        "v",
									ignoreCase: true,
								},
							},
							&ruleRefExpr{
								pos:  position{line: 83, col: 67, offset: 3406},
								name: "DIGIT",
							},
							&zeroOrMoreExpr{
								pos: position{line: 83, col: 74, offset: 3413},
								expr: &seqExpr{
									pos: position{line: 83, col: 75, offset: 3414},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 83, col: 75, offset: 3414},
											expr: &ruleRefExpr{
												pos:  position{line: 83, col: 76, offset: 3415},
												name: "EOL",
											},
										},
										&notExpr{
											pos: position{line: 83, col: 80, offset: 3419},
											expr: &litMatcher{
												pos:        position{line: 83, col: 81, offset: 3420},
												val:        ",",
												ignoreCase: false,
											},
										},
										&notExpr{
											pos: position{line: 83, col: 85, offset: 3424},
											expr: &litMatcher{
												pos:        position{line: 83, col: 86, offset: 3425},
												val:        ":",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 83, col: 90, offset: 3429,
										},
									},
								},
							},
							&zeroOrMoreExpr{
								pos: position{line: 83, col: 94, offset: 3433},
								expr: &ruleRefExpr{
									pos:  position{line: 83, col: 94, offset: 3433},
									name: "WS",
								},
							},
							&andExpr{
								pos: position{line: 83, col: 98, offset: 3437},
								expr: &litMatcher{
									pos:        position{line: 83, col: 99, offset: 3438},
									val:        ",",
									ignoreCase: false,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "DocumentRevisionDate",
			pos:  position{line: 84, col: 1, offset: 3442},
			expr: &zeroOrMoreExpr{
				pos: position{line: 84, col: 25, offset: 3466},
				expr: &seqExpr{
					pos: position{line: 84, col: 26, offset: 3467},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 84, col: 26, offset: 3467},
							expr: &ruleRefExpr{
								pos:  position{line: 84, col: 27, offset: 3468},
								name: "EOL",
							},
						},
						&notExpr{
							pos: position{line: 84, col: 31, offset: 3472},
							expr: &litMatcher{
								pos:        position{line: 84, col: 32, offset: 3473},
								val:        ":",
								ignoreCase: false,
							},
						},
						&anyMatcher{
							line: 84, col: 36, offset: 3477,
						},
					},
				},
			},
		},
		{
			name: "DocumentRevisionRemark",
			pos:  position{line: 85, col: 1, offset: 3482},
			expr: &zeroOrMoreExpr{
				pos: position{line: 85, col: 27, offset: 3508},
				expr: &seqExpr{
					pos: position{line: 85, col: 28, offset: 3509},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 85, col: 28, offset: 3509},
							expr: &ruleRefExpr{
								pos:  position{line: 85, col: 29, offset: 3510},
								name: "EOL",
							},
						},
						&anyMatcher{
							line: 85, col: 33, offset: 3514,
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclaration",
			pos:  position{line: 90, col: 1, offset: 3634},
			expr: &choiceExpr{
				pos: position{line: 90, col: 33, offset: 3666},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 90, col: 33, offset: 3666},
						name: "DocumentAttributeDeclarationWithNameOnly",
					},
					&ruleRefExpr{
						pos:  position{line: 90, col: 76, offset: 3709},
						name: "DocumentAttributeDeclarationWithNameAndValue",
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclarationWithNameOnly",
			pos:  position{line: 92, col: 1, offset: 3756},
			expr: &actionExpr{
				pos: position{line: 92, col: 45, offset: 3800},
				run: (*parser).callonDocumentAttributeDeclarationWithNameOnly1,
				expr: &seqExpr{
					pos: position{line: 92, col: 45, offset: 3800},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 92, col: 45, offset: 3800},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 92, col: 49, offset: 3804},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 92, col: 55, offset: 3810},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 92, col: 70, offset: 3825},
							val:        ":",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 92, col: 74, offset: 3829},
							expr: &ruleRefExpr{
								pos:  position{line: 92, col: 74, offset: 3829},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 92, col: 78, offset: 3833},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclarationWithNameAndValue",
			pos:  position{line: 96, col: 1, offset: 3918},
			expr: &actionExpr{
				pos: position{line: 96, col: 49, offset: 3966},
				run: (*parser).callonDocumentAttributeDeclarationWithNameAndValue1,
				expr: &seqExpr{
					pos: position{line: 96, col: 49, offset: 3966},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 96, col: 49, offset: 3966},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 96, col: 53, offset: 3970},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 96, col: 59, offset: 3976},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 96, col: 74, offset: 3991},
							val:        ":",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 96, col: 78, offset: 3995},
							expr: &ruleRefExpr{
								pos:  position{line: 96, col: 78, offset: 3995},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 96, col: 82, offset: 3999},
							label: "value",
							expr: &zeroOrMoreExpr{
								pos: position{line: 96, col: 88, offset: 4005},
								expr: &seqExpr{
									pos: position{line: 96, col: 89, offset: 4006},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 96, col: 89, offset: 4006},
											expr: &ruleRefExpr{
												pos:  position{line: 96, col: 90, offset: 4007},
												name: "NEWLINE",
											},
										},
										&anyMatcher{
											line: 96, col: 98, offset: 4015,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 96, col: 102, offset: 4019},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeReset",
			pos:  position{line: 100, col: 1, offset: 4122},
			expr: &choiceExpr{
				pos: position{line: 100, col: 27, offset: 4148},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 100, col: 27, offset: 4148},
						name: "DocumentAttributeResetWithSectionTitleBangSymbol",
					},
					&ruleRefExpr{
						pos:  position{line: 100, col: 78, offset: 4199},
						name: "DocumentAttributeResetWithTrailingBangSymbol",
					},
				},
			},
		},
		{
			name: "DocumentAttributeResetWithSectionTitleBangSymbol",
			pos:  position{line: 102, col: 1, offset: 4245},
			expr: &actionExpr{
				pos: position{line: 102, col: 53, offset: 4297},
				run: (*parser).callonDocumentAttributeResetWithSectionTitleBangSymbol1,
				expr: &seqExpr{
					pos: position{line: 102, col: 53, offset: 4297},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 102, col: 53, offset: 4297},
							val:        ":!",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 102, col: 58, offset: 4302},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 102, col: 64, offset: 4308},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 102, col: 79, offset: 4323},
							val:        ":",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 102, col: 83, offset: 4327},
							expr: &ruleRefExpr{
								pos:  position{line: 102, col: 83, offset: 4327},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 102, col: 87, offset: 4331},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeResetWithTrailingBangSymbol",
			pos:  position{line: 106, col: 1, offset: 4405},
			expr: &actionExpr{
				pos: position{line: 106, col: 49, offset: 4453},
				run: (*parser).callonDocumentAttributeResetWithTrailingBangSymbol1,
				expr: &seqExpr{
					pos: position{line: 106, col: 49, offset: 4453},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 106, col: 49, offset: 4453},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 106, col: 53, offset: 4457},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 106, col: 59, offset: 4463},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 106, col: 74, offset: 4478},
							val:        "!:",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 106, col: 79, offset: 4483},
							expr: &ruleRefExpr{
								pos:  position{line: 106, col: 79, offset: 4483},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 106, col: 83, offset: 4487},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeSubstitution",
			pos:  position{line: 111, col: 1, offset: 4562},
			expr: &actionExpr{
				pos: position{line: 111, col: 34, offset: 4595},
				run: (*parser).callonDocumentAttributeSubstitution1,
				expr: &seqExpr{
					pos: position{line: 111, col: 34, offset: 4595},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 111, col: 34, offset: 4595},
							val:        "{",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 111, col: 38, offset: 4599},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 111, col: 44, offset: 4605},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 111, col: 59, offset: 4620},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "AttributeName",
			pos:  position{line: 118, col: 1, offset: 4874},
			expr: &seqExpr{
				pos: position{line: 118, col: 18, offset: 4891},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 118, col: 19, offset: 4892},
						alternatives: []interface{}{
							&charClassMatcher{
								pos:        position{line: 118, col: 19, offset: 4892},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
							&charClassMatcher{
								pos:        position{line: 118, col: 27, offset: 4900},
								val:        "[a-z]",
								ranges:     []rune{'a', 'z'},
								ignoreCase: false,
								inverted:   false,
							},
							&charClassMatcher{
								pos:        position{line: 118, col: 35, offset: 4908},
								val:        "[0-9]",
								ranges:     []rune{'0', '9'},
								ignoreCase: false,
								inverted:   false,
							},
							&litMatcher{
								pos:        position{line: 118, col: 43, offset: 4916},
								val:        "_",
								ignoreCase: false,
							},
						},
					},
					&zeroOrMoreExpr{
						pos: position{line: 118, col: 48, offset: 4921},
						expr: &choiceExpr{
							pos: position{line: 118, col: 49, offset: 4922},
							alternatives: []interface{}{
								&charClassMatcher{
									pos:        position{line: 118, col: 49, offset: 4922},
									val:        "[A-Z]",
									ranges:     []rune{'A', 'Z'},
									ignoreCase: false,
									inverted:   false,
								},
								&charClassMatcher{
									pos:        position{line: 118, col: 57, offset: 4930},
									val:        "[a-z]",
									ranges:     []rune{'a', 'z'},
									ignoreCase: false,
									inverted:   false,
								},
								&charClassMatcher{
									pos:        position{line: 118, col: 65, offset: 4938},
									val:        "[0-9]",
									ranges:     []rune{'0', '9'},
									ignoreCase: false,
									inverted:   false,
								},
								&litMatcher{
									pos:        position{line: 118, col: 73, offset: 4946},
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
			name: "Section",
			pos:  position{line: 123, col: 1, offset: 5057},
			expr: &choiceExpr{
				pos: position{line: 123, col: 12, offset: 5068},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 123, col: 12, offset: 5068},
						name: "Section1",
					},
					&ruleRefExpr{
						pos:  position{line: 123, col: 23, offset: 5079},
						name: "Section2",
					},
					&ruleRefExpr{
						pos:  position{line: 123, col: 34, offset: 5090},
						name: "Section3",
					},
					&ruleRefExpr{
						pos:  position{line: 123, col: 45, offset: 5101},
						name: "Section4",
					},
					&ruleRefExpr{
						pos:  position{line: 123, col: 56, offset: 5112},
						name: "Section5",
					},
				},
			},
		},
		{
			name: "Section1",
			pos:  position{line: 126, col: 1, offset: 5123},
			expr: &actionExpr{
				pos: position{line: 126, col: 13, offset: 5135},
				run: (*parser).callonSection11,
				expr: &seqExpr{
					pos: position{line: 126, col: 13, offset: 5135},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 126, col: 13, offset: 5135},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 126, col: 21, offset: 5143},
								name: "Section1Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 126, col: 36, offset: 5158},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 126, col: 46, offset: 5168},
								expr: &ruleRefExpr{
									pos:  position{line: 126, col: 46, offset: 5168},
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
			pos:  position{line: 130, col: 1, offset: 5276},
			expr: &actionExpr{
				pos: position{line: 130, col: 18, offset: 5293},
				run: (*parser).callonSection1Block1,
				expr: &seqExpr{
					pos: position{line: 130, col: 18, offset: 5293},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 130, col: 18, offset: 5293},
							expr: &ruleRefExpr{
								pos:  position{line: 130, col: 19, offset: 5294},
								name: "Section1",
							},
						},
						&labeledExpr{
							pos:   position{line: 130, col: 28, offset: 5303},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 130, col: 37, offset: 5312},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 130, col: 37, offset: 5312},
										name: "Section2",
									},
									&ruleRefExpr{
										pos:  position{line: 130, col: 48, offset: 5323},
										name: "Section3",
									},
									&ruleRefExpr{
										pos:  position{line: 130, col: 59, offset: 5334},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 130, col: 70, offset: 5345},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 130, col: 81, offset: 5356},
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
			name: "Section2",
			pos:  position{line: 134, col: 1, offset: 5421},
			expr: &actionExpr{
				pos: position{line: 134, col: 13, offset: 5433},
				run: (*parser).callonSection21,
				expr: &seqExpr{
					pos: position{line: 134, col: 13, offset: 5433},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 134, col: 13, offset: 5433},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 134, col: 21, offset: 5441},
								name: "Section2Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 134, col: 36, offset: 5456},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 134, col: 46, offset: 5466},
								expr: &ruleRefExpr{
									pos:  position{line: 134, col: 46, offset: 5466},
									name: "Section2Block",
								},
							},
						},
						&andExpr{
							pos: position{line: 134, col: 62, offset: 5482},
							expr: &zeroOrMoreExpr{
								pos: position{line: 134, col: 63, offset: 5483},
								expr: &ruleRefExpr{
									pos:  position{line: 134, col: 64, offset: 5484},
									name: "Section2",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section2Block",
			pos:  position{line: 138, col: 1, offset: 5587},
			expr: &actionExpr{
				pos: position{line: 138, col: 18, offset: 5604},
				run: (*parser).callonSection2Block1,
				expr: &seqExpr{
					pos: position{line: 138, col: 18, offset: 5604},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 138, col: 18, offset: 5604},
							expr: &ruleRefExpr{
								pos:  position{line: 138, col: 19, offset: 5605},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 138, col: 28, offset: 5614},
							expr: &ruleRefExpr{
								pos:  position{line: 138, col: 29, offset: 5615},
								name: "Section2",
							},
						},
						&labeledExpr{
							pos:   position{line: 138, col: 38, offset: 5624},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 138, col: 47, offset: 5633},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 138, col: 47, offset: 5633},
										name: "Section3",
									},
									&ruleRefExpr{
										pos:  position{line: 138, col: 58, offset: 5644},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 138, col: 69, offset: 5655},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 138, col: 80, offset: 5666},
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
			pos:  position{line: 142, col: 1, offset: 5731},
			expr: &actionExpr{
				pos: position{line: 142, col: 13, offset: 5743},
				run: (*parser).callonSection31,
				expr: &seqExpr{
					pos: position{line: 142, col: 13, offset: 5743},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 142, col: 13, offset: 5743},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 142, col: 21, offset: 5751},
								name: "Section3Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 142, col: 36, offset: 5766},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 142, col: 46, offset: 5776},
								expr: &ruleRefExpr{
									pos:  position{line: 142, col: 46, offset: 5776},
									name: "Section3Block",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section3Block",
			pos:  position{line: 146, col: 1, offset: 5884},
			expr: &actionExpr{
				pos: position{line: 146, col: 18, offset: 5901},
				run: (*parser).callonSection3Block1,
				expr: &seqExpr{
					pos: position{line: 146, col: 18, offset: 5901},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 146, col: 18, offset: 5901},
							expr: &ruleRefExpr{
								pos:  position{line: 146, col: 19, offset: 5902},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 146, col: 28, offset: 5911},
							expr: &ruleRefExpr{
								pos:  position{line: 146, col: 29, offset: 5912},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 146, col: 38, offset: 5921},
							expr: &ruleRefExpr{
								pos:  position{line: 146, col: 39, offset: 5922},
								name: "Section3",
							},
						},
						&labeledExpr{
							pos:   position{line: 146, col: 48, offset: 5931},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 146, col: 57, offset: 5940},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 146, col: 57, offset: 5940},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 146, col: 68, offset: 5951},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 146, col: 79, offset: 5962},
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
			pos:  position{line: 150, col: 1, offset: 6027},
			expr: &actionExpr{
				pos: position{line: 150, col: 13, offset: 6039},
				run: (*parser).callonSection41,
				expr: &seqExpr{
					pos: position{line: 150, col: 13, offset: 6039},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 150, col: 13, offset: 6039},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 150, col: 21, offset: 6047},
								name: "Section4Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 150, col: 36, offset: 6062},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 150, col: 46, offset: 6072},
								expr: &ruleRefExpr{
									pos:  position{line: 150, col: 46, offset: 6072},
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
			pos:  position{line: 154, col: 1, offset: 6180},
			expr: &actionExpr{
				pos: position{line: 154, col: 18, offset: 6197},
				run: (*parser).callonSection4Block1,
				expr: &seqExpr{
					pos: position{line: 154, col: 18, offset: 6197},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 154, col: 18, offset: 6197},
							expr: &ruleRefExpr{
								pos:  position{line: 154, col: 19, offset: 6198},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 154, col: 28, offset: 6207},
							expr: &ruleRefExpr{
								pos:  position{line: 154, col: 29, offset: 6208},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 154, col: 38, offset: 6217},
							expr: &ruleRefExpr{
								pos:  position{line: 154, col: 39, offset: 6218},
								name: "Section3",
							},
						},
						&notExpr{
							pos: position{line: 154, col: 48, offset: 6227},
							expr: &ruleRefExpr{
								pos:  position{line: 154, col: 49, offset: 6228},
								name: "Section4",
							},
						},
						&labeledExpr{
							pos:   position{line: 154, col: 58, offset: 6237},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 154, col: 67, offset: 6246},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 154, col: 67, offset: 6246},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 154, col: 78, offset: 6257},
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
			pos:  position{line: 158, col: 1, offset: 6322},
			expr: &actionExpr{
				pos: position{line: 158, col: 13, offset: 6334},
				run: (*parser).callonSection51,
				expr: &seqExpr{
					pos: position{line: 158, col: 13, offset: 6334},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 158, col: 13, offset: 6334},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 158, col: 21, offset: 6342},
								name: "Section5Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 158, col: 36, offset: 6357},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 158, col: 46, offset: 6367},
								expr: &ruleRefExpr{
									pos:  position{line: 158, col: 46, offset: 6367},
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
			pos:  position{line: 162, col: 1, offset: 6475},
			expr: &actionExpr{
				pos: position{line: 162, col: 18, offset: 6492},
				run: (*parser).callonSection5Block1,
				expr: &seqExpr{
					pos: position{line: 162, col: 18, offset: 6492},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 162, col: 18, offset: 6492},
							expr: &ruleRefExpr{
								pos:  position{line: 162, col: 19, offset: 6493},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 162, col: 28, offset: 6502},
							expr: &ruleRefExpr{
								pos:  position{line: 162, col: 29, offset: 6503},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 162, col: 38, offset: 6512},
							expr: &ruleRefExpr{
								pos:  position{line: 162, col: 39, offset: 6513},
								name: "Section3",
							},
						},
						&notExpr{
							pos: position{line: 162, col: 48, offset: 6522},
							expr: &ruleRefExpr{
								pos:  position{line: 162, col: 49, offset: 6523},
								name: "Section4",
							},
						},
						&notExpr{
							pos: position{line: 162, col: 58, offset: 6532},
							expr: &ruleRefExpr{
								pos:  position{line: 162, col: 59, offset: 6533},
								name: "Section5",
							},
						},
						&labeledExpr{
							pos:   position{line: 162, col: 68, offset: 6542},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 162, col: 77, offset: 6551},
								name: "StandaloneBlock",
							},
						},
					},
				},
			},
		},
		{
			name: "SectionTitle",
			pos:  position{line: 170, col: 1, offset: 6727},
			expr: &choiceExpr{
				pos: position{line: 170, col: 17, offset: 6743},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 170, col: 17, offset: 6743},
						name: "Section1Title",
					},
					&ruleRefExpr{
						pos:  position{line: 170, col: 33, offset: 6759},
						name: "Section2Title",
					},
					&ruleRefExpr{
						pos:  position{line: 170, col: 49, offset: 6775},
						name: "Section3Title",
					},
					&ruleRefExpr{
						pos:  position{line: 170, col: 65, offset: 6791},
						name: "Section4Title",
					},
					&ruleRefExpr{
						pos:  position{line: 170, col: 81, offset: 6807},
						name: "Section5Title",
					},
				},
			},
		},
		{
			name: "Section1Title",
			pos:  position{line: 172, col: 1, offset: 6822},
			expr: &actionExpr{
				pos: position{line: 172, col: 18, offset: 6839},
				run: (*parser).callonSection1Title1,
				expr: &seqExpr{
					pos: position{line: 172, col: 18, offset: 6839},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 172, col: 18, offset: 6839},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 172, col: 29, offset: 6850},
								expr: &ruleRefExpr{
									pos:  position{line: 172, col: 30, offset: 6851},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 172, col: 49, offset: 6870},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 172, col: 56, offset: 6877},
								val:        "==",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 172, col: 62, offset: 6883},
							expr: &ruleRefExpr{
								pos:  position{line: 172, col: 62, offset: 6883},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 172, col: 66, offset: 6887},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 172, col: 74, offset: 6895},
								name: "InlineContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 172, col: 88, offset: 6909},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 172, col: 93, offset: 6914},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 172, col: 93, offset: 6914},
									expr: &ruleRefExpr{
										pos:  position{line: 172, col: 93, offset: 6914},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 172, col: 106, offset: 6927},
									name: "EOF",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section2Title",
			pos:  position{line: 176, col: 1, offset: 7032},
			expr: &actionExpr{
				pos: position{line: 176, col: 18, offset: 7049},
				run: (*parser).callonSection2Title1,
				expr: &seqExpr{
					pos: position{line: 176, col: 18, offset: 7049},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 176, col: 18, offset: 7049},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 176, col: 29, offset: 7060},
								expr: &ruleRefExpr{
									pos:  position{line: 176, col: 30, offset: 7061},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 176, col: 49, offset: 7080},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 176, col: 56, offset: 7087},
								val:        "===",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 176, col: 63, offset: 7094},
							expr: &ruleRefExpr{
								pos:  position{line: 176, col: 63, offset: 7094},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 176, col: 67, offset: 7098},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 176, col: 75, offset: 7106},
								name: "InlineContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 176, col: 89, offset: 7120},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 176, col: 94, offset: 7125},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 176, col: 94, offset: 7125},
									expr: &ruleRefExpr{
										pos:  position{line: 176, col: 94, offset: 7125},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 176, col: 107, offset: 7138},
									name: "EOF",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section3Title",
			pos:  position{line: 180, col: 1, offset: 7242},
			expr: &actionExpr{
				pos: position{line: 180, col: 18, offset: 7259},
				run: (*parser).callonSection3Title1,
				expr: &seqExpr{
					pos: position{line: 180, col: 18, offset: 7259},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 180, col: 18, offset: 7259},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 180, col: 29, offset: 7270},
								expr: &ruleRefExpr{
									pos:  position{line: 180, col: 30, offset: 7271},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 180, col: 49, offset: 7290},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 180, col: 56, offset: 7297},
								val:        "====",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 180, col: 64, offset: 7305},
							expr: &ruleRefExpr{
								pos:  position{line: 180, col: 64, offset: 7305},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 180, col: 68, offset: 7309},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 180, col: 76, offset: 7317},
								name: "InlineContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 180, col: 90, offset: 7331},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 180, col: 95, offset: 7336},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 180, col: 95, offset: 7336},
									expr: &ruleRefExpr{
										pos:  position{line: 180, col: 95, offset: 7336},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 180, col: 108, offset: 7349},
									name: "EOF",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section4Title",
			pos:  position{line: 184, col: 1, offset: 7453},
			expr: &actionExpr{
				pos: position{line: 184, col: 18, offset: 7470},
				run: (*parser).callonSection4Title1,
				expr: &seqExpr{
					pos: position{line: 184, col: 18, offset: 7470},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 184, col: 18, offset: 7470},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 184, col: 29, offset: 7481},
								expr: &ruleRefExpr{
									pos:  position{line: 184, col: 30, offset: 7482},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 184, col: 49, offset: 7501},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 184, col: 56, offset: 7508},
								val:        "=====",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 184, col: 65, offset: 7517},
							expr: &ruleRefExpr{
								pos:  position{line: 184, col: 65, offset: 7517},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 184, col: 69, offset: 7521},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 184, col: 77, offset: 7529},
								name: "InlineContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 184, col: 91, offset: 7543},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 184, col: 96, offset: 7548},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 184, col: 96, offset: 7548},
									expr: &ruleRefExpr{
										pos:  position{line: 184, col: 96, offset: 7548},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 184, col: 109, offset: 7561},
									name: "EOF",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section5Title",
			pos:  position{line: 188, col: 1, offset: 7665},
			expr: &actionExpr{
				pos: position{line: 188, col: 18, offset: 7682},
				run: (*parser).callonSection5Title1,
				expr: &seqExpr{
					pos: position{line: 188, col: 18, offset: 7682},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 188, col: 18, offset: 7682},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 188, col: 29, offset: 7693},
								expr: &ruleRefExpr{
									pos:  position{line: 188, col: 30, offset: 7694},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 188, col: 49, offset: 7713},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 188, col: 56, offset: 7720},
								val:        "======",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 188, col: 66, offset: 7730},
							expr: &ruleRefExpr{
								pos:  position{line: 188, col: 66, offset: 7730},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 188, col: 70, offset: 7734},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 188, col: 78, offset: 7742},
								name: "InlineContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 188, col: 92, offset: 7756},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 188, col: 97, offset: 7761},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 188, col: 97, offset: 7761},
									expr: &ruleRefExpr{
										pos:  position{line: 188, col: 97, offset: 7761},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 188, col: 110, offset: 7774},
									name: "EOF",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "List",
			pos:  position{line: 195, col: 1, offset: 7984},
			expr: &actionExpr{
				pos: position{line: 195, col: 9, offset: 7992},
				run: (*parser).callonList1,
				expr: &seqExpr{
					pos: position{line: 195, col: 9, offset: 7992},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 195, col: 9, offset: 7992},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 195, col: 20, offset: 8003},
								expr: &ruleRefExpr{
									pos:  position{line: 195, col: 21, offset: 8004},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 197, col: 5, offset: 8096},
							label: "elements",
							expr: &oneOrMoreExpr{
								pos: position{line: 197, col: 14, offset: 8105},
								expr: &seqExpr{
									pos: position{line: 197, col: 15, offset: 8106},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 197, col: 15, offset: 8106},
											name: "ListItem",
										},
										&zeroOrOneExpr{
											pos: position{line: 197, col: 24, offset: 8115},
											expr: &ruleRefExpr{
												pos:  position{line: 197, col: 24, offset: 8115},
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
			pos:  position{line: 201, col: 1, offset: 8212},
			expr: &actionExpr{
				pos: position{line: 201, col: 13, offset: 8224},
				run: (*parser).callonListItem1,
				expr: &seqExpr{
					pos: position{line: 201, col: 13, offset: 8224},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 201, col: 13, offset: 8224},
							expr: &ruleRefExpr{
								pos:  position{line: 201, col: 13, offset: 8224},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 201, col: 17, offset: 8228},
							label: "level",
							expr: &choiceExpr{
								pos: position{line: 201, col: 24, offset: 8235},
								alternatives: []interface{}{
									&oneOrMoreExpr{
										pos: position{line: 201, col: 24, offset: 8235},
										expr: &litMatcher{
											pos:        position{line: 201, col: 24, offset: 8235},
											val:        "*",
											ignoreCase: false,
										},
									},
									&litMatcher{
										pos:        position{line: 201, col: 31, offset: 8242},
										val:        "-",
										ignoreCase: false,
									},
								},
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 201, col: 36, offset: 8247},
							expr: &ruleRefExpr{
								pos:  position{line: 201, col: 36, offset: 8247},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 201, col: 40, offset: 8251},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 201, col: 49, offset: 8260},
								name: "ListItemContent",
							},
						},
					},
				},
			},
		},
		{
			name: "ListItemContent",
			pos:  position{line: 205, col: 1, offset: 8357},
			expr: &actionExpr{
				pos: position{line: 205, col: 20, offset: 8376},
				run: (*parser).callonListItemContent1,
				expr: &labeledExpr{
					pos:   position{line: 205, col: 20, offset: 8376},
					label: "lines",
					expr: &oneOrMoreExpr{
						pos: position{line: 205, col: 26, offset: 8382},
						expr: &seqExpr{
							pos: position{line: 205, col: 27, offset: 8383},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 205, col: 27, offset: 8383},
									expr: &seqExpr{
										pos: position{line: 205, col: 29, offset: 8385},
										exprs: []interface{}{
											&zeroOrMoreExpr{
												pos: position{line: 205, col: 29, offset: 8385},
												expr: &ruleRefExpr{
													pos:  position{line: 205, col: 29, offset: 8385},
													name: "WS",
												},
											},
											&choiceExpr{
												pos: position{line: 205, col: 34, offset: 8390},
												alternatives: []interface{}{
													&oneOrMoreExpr{
														pos: position{line: 205, col: 34, offset: 8390},
														expr: &litMatcher{
															pos:        position{line: 205, col: 34, offset: 8390},
															val:        "*",
															ignoreCase: false,
														},
													},
													&litMatcher{
														pos:        position{line: 205, col: 41, offset: 8397},
														val:        "-",
														ignoreCase: false,
													},
												},
											},
											&oneOrMoreExpr{
												pos: position{line: 205, col: 46, offset: 8402},
												expr: &ruleRefExpr{
													pos:  position{line: 205, col: 46, offset: 8402},
													name: "WS",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 205, col: 51, offset: 8407},
									name: "InlineContent",
								},
								&ruleRefExpr{
									pos:  position{line: 205, col: 65, offset: 8421},
									name: "EOL",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Paragraph",
			pos:  position{line: 213, col: 1, offset: 8750},
			expr: &actionExpr{
				pos: position{line: 213, col: 14, offset: 8763},
				run: (*parser).callonParagraph1,
				expr: &seqExpr{
					pos: position{line: 213, col: 14, offset: 8763},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 213, col: 14, offset: 8763},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 213, col: 25, offset: 8774},
								expr: &ruleRefExpr{
									pos:  position{line: 213, col: 26, offset: 8775},
									name: "ElementAttribute",
								},
							},
						},
						&notExpr{
							pos: position{line: 213, col: 45, offset: 8794},
							expr: &seqExpr{
								pos: position{line: 213, col: 47, offset: 8796},
								exprs: []interface{}{
									&oneOrMoreExpr{
										pos: position{line: 213, col: 47, offset: 8796},
										expr: &litMatcher{
											pos:        position{line: 213, col: 47, offset: 8796},
											val:        "=",
											ignoreCase: false,
										},
									},
									&oneOrMoreExpr{
										pos: position{line: 213, col: 52, offset: 8801},
										expr: &ruleRefExpr{
											pos:  position{line: 213, col: 52, offset: 8801},
											name: "WS",
										},
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 213, col: 57, offset: 8806},
							label: "lines",
							expr: &oneOrMoreExpr{
								pos: position{line: 213, col: 63, offset: 8812},
								expr: &seqExpr{
									pos: position{line: 213, col: 64, offset: 8813},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 213, col: 64, offset: 8813},
											name: "InlineContent",
										},
										&ruleRefExpr{
											pos:  position{line: 213, col: 78, offset: 8827},
											name: "EOL",
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
			name: "InlineContent",
			pos:  position{line: 219, col: 1, offset: 9117},
			expr: &actionExpr{
				pos: position{line: 219, col: 18, offset: 9134},
				run: (*parser).callonInlineContent1,
				expr: &seqExpr{
					pos: position{line: 219, col: 18, offset: 9134},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 219, col: 18, offset: 9134},
							expr: &ruleRefExpr{
								pos:  position{line: 219, col: 19, offset: 9135},
								name: "FencedBlockDelimiter",
							},
						},
						&labeledExpr{
							pos:   position{line: 219, col: 40, offset: 9156},
							label: "elements",
							expr: &oneOrMoreExpr{
								pos: position{line: 219, col: 49, offset: 9165},
								expr: &seqExpr{
									pos: position{line: 219, col: 50, offset: 9166},
									exprs: []interface{}{
										&zeroOrMoreExpr{
											pos: position{line: 219, col: 50, offset: 9166},
											expr: &ruleRefExpr{
												pos:  position{line: 219, col: 50, offset: 9166},
												name: "WS",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 219, col: 54, offset: 9170},
											name: "InlineElement",
										},
										&zeroOrMoreExpr{
											pos: position{line: 219, col: 68, offset: 9184},
											expr: &ruleRefExpr{
												pos:  position{line: 219, col: 68, offset: 9184},
												name: "WS",
											},
										},
									},
								},
							},
						},
						&andExpr{
							pos: position{line: 219, col: 74, offset: 9190},
							expr: &ruleRefExpr{
								pos:  position{line: 219, col: 75, offset: 9191},
								name: "EOL",
							},
						},
					},
				},
			},
		},
		{
			name: "InlineElement",
			pos:  position{line: 223, col: 1, offset: 9309},
			expr: &choiceExpr{
				pos: position{line: 223, col: 18, offset: 9326},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 223, col: 18, offset: 9326},
						name: "Passthrough",
					},
					&ruleRefExpr{
						pos:  position{line: 223, col: 32, offset: 9340},
						name: "InlineImage",
					},
					&ruleRefExpr{
						pos:  position{line: 223, col: 46, offset: 9354},
						name: "QuotedText",
					},
					&ruleRefExpr{
						pos:  position{line: 223, col: 59, offset: 9367},
						name: "ExternalLink",
					},
					&ruleRefExpr{
						pos:  position{line: 223, col: 74, offset: 9382},
						name: "DocumentAttributeSubstitution",
					},
					&ruleRefExpr{
						pos:  position{line: 223, col: 106, offset: 9414},
						name: "Characters",
					},
				},
			},
		},
		{
			name: "QuotedText",
			pos:  position{line: 228, col: 1, offset: 9665},
			expr: &choiceExpr{
				pos: position{line: 228, col: 15, offset: 9679},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 228, col: 15, offset: 9679},
						name: "BoldText",
					},
					&ruleRefExpr{
						pos:  position{line: 228, col: 26, offset: 9690},
						name: "ItalicText",
					},
					&ruleRefExpr{
						pos:  position{line: 228, col: 39, offset: 9703},
						name: "MonospaceText",
					},
					&ruleRefExpr{
						pos:  position{line: 229, col: 13, offset: 9731},
						name: "EscapedBoldText",
					},
					&ruleRefExpr{
						pos:  position{line: 229, col: 31, offset: 9749},
						name: "EscapedItalicText",
					},
					&ruleRefExpr{
						pos:  position{line: 229, col: 51, offset: 9769},
						name: "EscapedMonospaceText",
					},
				},
			},
		},
		{
			name: "BoldText",
			pos:  position{line: 231, col: 1, offset: 9791},
			expr: &choiceExpr{
				pos: position{line: 231, col: 13, offset: 9803},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 231, col: 13, offset: 9803},
						name: "BoldTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 231, col: 41, offset: 9831},
						name: "BoldTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 231, col: 73, offset: 9863},
						name: "BoldTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "BoldTextSimplePunctuation",
			pos:  position{line: 233, col: 1, offset: 9936},
			expr: &actionExpr{
				pos: position{line: 233, col: 30, offset: 9965},
				run: (*parser).callonBoldTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 233, col: 30, offset: 9965},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 233, col: 30, offset: 9965},
							expr: &litMatcher{
								pos:        position{line: 233, col: 31, offset: 9966},
								val:        "\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 233, col: 35, offset: 9970},
							val:        "*",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 233, col: 39, offset: 9974},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 233, col: 48, offset: 9983},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 233, col: 67, offset: 10002},
							val:        "*",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "BoldTextDoublePunctuation",
			pos:  position{line: 237, col: 1, offset: 10079},
			expr: &actionExpr{
				pos: position{line: 237, col: 30, offset: 10108},
				run: (*parser).callonBoldTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 237, col: 30, offset: 10108},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 237, col: 30, offset: 10108},
							expr: &litMatcher{
								pos:        position{line: 237, col: 31, offset: 10109},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 237, col: 36, offset: 10114},
							val:        "**",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 237, col: 41, offset: 10119},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 237, col: 50, offset: 10128},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 237, col: 69, offset: 10147},
							val:        "**",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "BoldTextUnbalancedPunctuation",
			pos:  position{line: 241, col: 1, offset: 10225},
			expr: &actionExpr{
				pos: position{line: 241, col: 34, offset: 10258},
				run: (*parser).callonBoldTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 241, col: 34, offset: 10258},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 241, col: 34, offset: 10258},
							expr: &litMatcher{
								pos:        position{line: 241, col: 35, offset: 10259},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 241, col: 40, offset: 10264},
							val:        "**",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 241, col: 45, offset: 10269},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 241, col: 54, offset: 10278},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 241, col: 73, offset: 10297},
							val:        "*",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedBoldText",
			pos:  position{line: 246, col: 1, offset: 10461},
			expr: &choiceExpr{
				pos: position{line: 246, col: 20, offset: 10480},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 246, col: 20, offset: 10480},
						name: "EscapedBoldTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 246, col: 55, offset: 10515},
						name: "EscapedBoldTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 246, col: 94, offset: 10554},
						name: "EscapedBoldTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "EscapedBoldTextSimplePunctuation",
			pos:  position{line: 248, col: 1, offset: 10634},
			expr: &actionExpr{
				pos: position{line: 248, col: 37, offset: 10670},
				run: (*parser).callonEscapedBoldTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 248, col: 37, offset: 10670},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 248, col: 37, offset: 10670},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 248, col: 50, offset: 10683},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 248, col: 50, offset: 10683},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 248, col: 54, offset: 10687},
										expr: &litMatcher{
											pos:        position{line: 248, col: 54, offset: 10687},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 248, col: 60, offset: 10693},
							val:        "*",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 248, col: 64, offset: 10697},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 248, col: 73, offset: 10706},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 248, col: 92, offset: 10725},
							val:        "*",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedBoldTextDoublePunctuation",
			pos:  position{line: 252, col: 1, offset: 10831},
			expr: &actionExpr{
				pos: position{line: 252, col: 37, offset: 10867},
				run: (*parser).callonEscapedBoldTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 252, col: 37, offset: 10867},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 252, col: 37, offset: 10867},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 252, col: 50, offset: 10880},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 252, col: 50, offset: 10880},
										val:        "\\\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 252, col: 55, offset: 10885},
										expr: &litMatcher{
											pos:        position{line: 252, col: 55, offset: 10885},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 252, col: 61, offset: 10891},
							val:        "**",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 252, col: 66, offset: 10896},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 252, col: 75, offset: 10905},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 252, col: 94, offset: 10924},
							val:        "**",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedBoldTextUnbalancedPunctuation",
			pos:  position{line: 256, col: 1, offset: 11032},
			expr: &actionExpr{
				pos: position{line: 256, col: 42, offset: 11073},
				run: (*parser).callonEscapedBoldTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 256, col: 42, offset: 11073},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 256, col: 42, offset: 11073},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 256, col: 55, offset: 11086},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 256, col: 55, offset: 11086},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 256, col: 59, offset: 11090},
										expr: &litMatcher{
											pos:        position{line: 256, col: 59, offset: 11090},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 256, col: 65, offset: 11096},
							val:        "**",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 256, col: 70, offset: 11101},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 256, col: 79, offset: 11110},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 256, col: 98, offset: 11129},
							val:        "*",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ItalicText",
			pos:  position{line: 261, col: 1, offset: 11322},
			expr: &choiceExpr{
				pos: position{line: 261, col: 15, offset: 11336},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 261, col: 15, offset: 11336},
						name: "ItalicTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 261, col: 45, offset: 11366},
						name: "ItalicTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 261, col: 79, offset: 11400},
						name: "ItalicTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "ItalicTextSimplePunctuation",
			pos:  position{line: 263, col: 1, offset: 11429},
			expr: &actionExpr{
				pos: position{line: 263, col: 32, offset: 11460},
				run: (*parser).callonItalicTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 263, col: 32, offset: 11460},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 263, col: 32, offset: 11460},
							expr: &litMatcher{
								pos:        position{line: 263, col: 33, offset: 11461},
								val:        "\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 263, col: 37, offset: 11465},
							val:        "_",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 263, col: 41, offset: 11469},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 263, col: 50, offset: 11478},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 263, col: 69, offset: 11497},
							val:        "_",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ItalicTextDoublePunctuation",
			pos:  position{line: 267, col: 1, offset: 11576},
			expr: &actionExpr{
				pos: position{line: 267, col: 32, offset: 11607},
				run: (*parser).callonItalicTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 267, col: 32, offset: 11607},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 267, col: 32, offset: 11607},
							expr: &litMatcher{
								pos:        position{line: 267, col: 33, offset: 11608},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 267, col: 38, offset: 11613},
							val:        "__",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 267, col: 43, offset: 11618},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 267, col: 52, offset: 11627},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 267, col: 71, offset: 11646},
							val:        "__",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ItalicTextUnbalancedPunctuation",
			pos:  position{line: 271, col: 1, offset: 11726},
			expr: &actionExpr{
				pos: position{line: 271, col: 36, offset: 11761},
				run: (*parser).callonItalicTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 271, col: 36, offset: 11761},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 271, col: 36, offset: 11761},
							expr: &litMatcher{
								pos:        position{line: 271, col: 37, offset: 11762},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 271, col: 42, offset: 11767},
							val:        "__",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 271, col: 47, offset: 11772},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 271, col: 56, offset: 11781},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 271, col: 75, offset: 11800},
							val:        "_",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedItalicText",
			pos:  position{line: 276, col: 1, offset: 11966},
			expr: &choiceExpr{
				pos: position{line: 276, col: 22, offset: 11987},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 276, col: 22, offset: 11987},
						name: "EscapedItalicTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 276, col: 59, offset: 12024},
						name: "EscapedItalicTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 276, col: 100, offset: 12065},
						name: "EscapedItalicTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "EscapedItalicTextSimplePunctuation",
			pos:  position{line: 278, col: 1, offset: 12147},
			expr: &actionExpr{
				pos: position{line: 278, col: 39, offset: 12185},
				run: (*parser).callonEscapedItalicTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 278, col: 39, offset: 12185},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 278, col: 39, offset: 12185},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 278, col: 52, offset: 12198},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 278, col: 52, offset: 12198},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 278, col: 56, offset: 12202},
										expr: &litMatcher{
											pos:        position{line: 278, col: 56, offset: 12202},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 278, col: 62, offset: 12208},
							val:        "_",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 278, col: 66, offset: 12212},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 278, col: 75, offset: 12221},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 278, col: 94, offset: 12240},
							val:        "_",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedItalicTextDoublePunctuation",
			pos:  position{line: 282, col: 1, offset: 12346},
			expr: &actionExpr{
				pos: position{line: 282, col: 39, offset: 12384},
				run: (*parser).callonEscapedItalicTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 282, col: 39, offset: 12384},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 282, col: 39, offset: 12384},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 282, col: 52, offset: 12397},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 282, col: 52, offset: 12397},
										val:        "\\\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 282, col: 57, offset: 12402},
										expr: &litMatcher{
											pos:        position{line: 282, col: 57, offset: 12402},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 282, col: 63, offset: 12408},
							val:        "__",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 282, col: 68, offset: 12413},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 282, col: 77, offset: 12422},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 282, col: 96, offset: 12441},
							val:        "__",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedItalicTextUnbalancedPunctuation",
			pos:  position{line: 286, col: 1, offset: 12549},
			expr: &actionExpr{
				pos: position{line: 286, col: 44, offset: 12592},
				run: (*parser).callonEscapedItalicTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 286, col: 44, offset: 12592},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 286, col: 44, offset: 12592},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 286, col: 57, offset: 12605},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 286, col: 57, offset: 12605},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 286, col: 61, offset: 12609},
										expr: &litMatcher{
											pos:        position{line: 286, col: 61, offset: 12609},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 286, col: 67, offset: 12615},
							val:        "__",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 286, col: 72, offset: 12620},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 286, col: 81, offset: 12629},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 286, col: 100, offset: 12648},
							val:        "_",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "MonospaceText",
			pos:  position{line: 291, col: 1, offset: 12841},
			expr: &choiceExpr{
				pos: position{line: 291, col: 18, offset: 12858},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 291, col: 18, offset: 12858},
						name: "MonospaceTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 291, col: 51, offset: 12891},
						name: "MonospaceTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 291, col: 88, offset: 12928},
						name: "MonospaceTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "MonospaceTextSimplePunctuation",
			pos:  position{line: 293, col: 1, offset: 12960},
			expr: &actionExpr{
				pos: position{line: 293, col: 35, offset: 12994},
				run: (*parser).callonMonospaceTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 293, col: 35, offset: 12994},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 293, col: 35, offset: 12994},
							expr: &litMatcher{
								pos:        position{line: 293, col: 36, offset: 12995},
								val:        "\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 293, col: 40, offset: 12999},
							val:        "`",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 293, col: 44, offset: 13003},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 293, col: 53, offset: 13012},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 293, col: 72, offset: 13031},
							val:        "`",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "MonospaceTextDoublePunctuation",
			pos:  position{line: 297, col: 1, offset: 13113},
			expr: &actionExpr{
				pos: position{line: 297, col: 35, offset: 13147},
				run: (*parser).callonMonospaceTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 297, col: 35, offset: 13147},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 297, col: 35, offset: 13147},
							expr: &litMatcher{
								pos:        position{line: 297, col: 36, offset: 13148},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 297, col: 41, offset: 13153},
							val:        "``",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 297, col: 46, offset: 13158},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 297, col: 55, offset: 13167},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 297, col: 74, offset: 13186},
							val:        "``",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "MonospaceTextUnbalancedPunctuation",
			pos:  position{line: 301, col: 1, offset: 13269},
			expr: &actionExpr{
				pos: position{line: 301, col: 39, offset: 13307},
				run: (*parser).callonMonospaceTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 301, col: 39, offset: 13307},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 301, col: 39, offset: 13307},
							expr: &litMatcher{
								pos:        position{line: 301, col: 40, offset: 13308},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 301, col: 45, offset: 13313},
							val:        "``",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 301, col: 50, offset: 13318},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 301, col: 59, offset: 13327},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 301, col: 78, offset: 13346},
							val:        "`",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedMonospaceText",
			pos:  position{line: 306, col: 1, offset: 13515},
			expr: &choiceExpr{
				pos: position{line: 306, col: 25, offset: 13539},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 306, col: 25, offset: 13539},
						name: "EscapedMonospaceTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 306, col: 65, offset: 13579},
						name: "EscapedMonospaceTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 306, col: 109, offset: 13623},
						name: "EscapedMonospaceTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "EscapedMonospaceTextSimplePunctuation",
			pos:  position{line: 308, col: 1, offset: 13708},
			expr: &actionExpr{
				pos: position{line: 308, col: 42, offset: 13749},
				run: (*parser).callonEscapedMonospaceTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 308, col: 42, offset: 13749},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 308, col: 42, offset: 13749},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 308, col: 55, offset: 13762},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 308, col: 55, offset: 13762},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 308, col: 59, offset: 13766},
										expr: &litMatcher{
											pos:        position{line: 308, col: 59, offset: 13766},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 308, col: 65, offset: 13772},
							val:        "`",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 308, col: 69, offset: 13776},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 308, col: 78, offset: 13785},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 308, col: 97, offset: 13804},
							val:        "`",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedMonospaceTextDoublePunctuation",
			pos:  position{line: 312, col: 1, offset: 13910},
			expr: &actionExpr{
				pos: position{line: 312, col: 42, offset: 13951},
				run: (*parser).callonEscapedMonospaceTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 312, col: 42, offset: 13951},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 312, col: 42, offset: 13951},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 312, col: 55, offset: 13964},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 312, col: 55, offset: 13964},
										val:        "\\\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 312, col: 60, offset: 13969},
										expr: &litMatcher{
											pos:        position{line: 312, col: 60, offset: 13969},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 312, col: 66, offset: 13975},
							val:        "``",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 312, col: 71, offset: 13980},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 312, col: 80, offset: 13989},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 312, col: 99, offset: 14008},
							val:        "``",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedMonospaceTextUnbalancedPunctuation",
			pos:  position{line: 316, col: 1, offset: 14116},
			expr: &actionExpr{
				pos: position{line: 316, col: 47, offset: 14162},
				run: (*parser).callonEscapedMonospaceTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 316, col: 47, offset: 14162},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 316, col: 47, offset: 14162},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 316, col: 60, offset: 14175},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 316, col: 60, offset: 14175},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 316, col: 64, offset: 14179},
										expr: &litMatcher{
											pos:        position{line: 316, col: 64, offset: 14179},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 316, col: 70, offset: 14185},
							val:        "``",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 316, col: 75, offset: 14190},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 316, col: 84, offset: 14199},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 316, col: 103, offset: 14218},
							val:        "`",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "QuotedTextContent",
			pos:  position{line: 321, col: 1, offset: 14411},
			expr: &seqExpr{
				pos: position{line: 321, col: 22, offset: 14432},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 321, col: 22, offset: 14432},
						name: "QuotedTextContentElement",
					},
					&zeroOrMoreExpr{
						pos: position{line: 321, col: 47, offset: 14457},
						expr: &seqExpr{
							pos: position{line: 321, col: 48, offset: 14458},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 321, col: 48, offset: 14458},
									expr: &ruleRefExpr{
										pos:  position{line: 321, col: 48, offset: 14458},
										name: "WS",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 321, col: 52, offset: 14462},
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
			pos:  position{line: 323, col: 1, offset: 14490},
			expr: &choiceExpr{
				pos: position{line: 323, col: 29, offset: 14518},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 323, col: 29, offset: 14518},
						name: "QuotedText",
					},
					&ruleRefExpr{
						pos:  position{line: 323, col: 42, offset: 14531},
						name: "QuotedTextCharacters",
					},
					&ruleRefExpr{
						pos:  position{line: 323, col: 65, offset: 14554},
						name: "CharactersWithQuotePunctuation",
					},
				},
			},
		},
		{
			name: "QuotedTextCharacters",
			pos:  position{line: 325, col: 1, offset: 14689},
			expr: &oneOrMoreExpr{
				pos: position{line: 325, col: 25, offset: 14713},
				expr: &seqExpr{
					pos: position{line: 325, col: 26, offset: 14714},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 325, col: 26, offset: 14714},
							expr: &ruleRefExpr{
								pos:  position{line: 325, col: 27, offset: 14715},
								name: "NEWLINE",
							},
						},
						&notExpr{
							pos: position{line: 325, col: 35, offset: 14723},
							expr: &ruleRefExpr{
								pos:  position{line: 325, col: 36, offset: 14724},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 325, col: 39, offset: 14727},
							expr: &litMatcher{
								pos:        position{line: 325, col: 40, offset: 14728},
								val:        "*",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 325, col: 44, offset: 14732},
							expr: &litMatcher{
								pos:        position{line: 325, col: 45, offset: 14733},
								val:        "_",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 325, col: 49, offset: 14737},
							expr: &litMatcher{
								pos:        position{line: 325, col: 50, offset: 14738},
								val:        "`",
								ignoreCase: false,
							},
						},
						&anyMatcher{
							line: 325, col: 54, offset: 14742,
						},
					},
				},
			},
		},
		{
			name: "CharactersWithQuotePunctuation",
			pos:  position{line: 326, col: 1, offset: 14784},
			expr: &actionExpr{
				pos: position{line: 326, col: 35, offset: 14818},
				run: (*parser).callonCharactersWithQuotePunctuation1,
				expr: &oneOrMoreExpr{
					pos: position{line: 326, col: 35, offset: 14818},
					expr: &seqExpr{
						pos: position{line: 326, col: 36, offset: 14819},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 326, col: 36, offset: 14819},
								expr: &ruleRefExpr{
									pos:  position{line: 326, col: 37, offset: 14820},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 326, col: 45, offset: 14828},
								expr: &ruleRefExpr{
									pos:  position{line: 326, col: 46, offset: 14829},
									name: "WS",
								},
							},
							&anyMatcher{
								line: 326, col: 50, offset: 14833,
							},
						},
					},
				},
			},
		},
		{
			name: "UnbalancedQuotePunctuation",
			pos:  position{line: 331, col: 1, offset: 15078},
			expr: &choiceExpr{
				pos: position{line: 331, col: 31, offset: 15108},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 331, col: 31, offset: 15108},
						val:        "*",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 331, col: 37, offset: 15114},
						val:        "_",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 331, col: 43, offset: 15120},
						val:        "`",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "Passthrough",
			pos:  position{line: 337, col: 1, offset: 15233},
			expr: &choiceExpr{
				pos: position{line: 337, col: 16, offset: 15248},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 337, col: 16, offset: 15248},
						name: "TriplePlusPassthrough",
					},
					&ruleRefExpr{
						pos:  position{line: 337, col: 40, offset: 15272},
						name: "SinglePlusPassthrough",
					},
					&ruleRefExpr{
						pos:  position{line: 337, col: 64, offset: 15296},
						name: "PassthroughMacro",
					},
				},
			},
		},
		{
			name: "SinglePlusPassthrough",
			pos:  position{line: 339, col: 1, offset: 15314},
			expr: &actionExpr{
				pos: position{line: 339, col: 26, offset: 15339},
				run: (*parser).callonSinglePlusPassthrough1,
				expr: &seqExpr{
					pos: position{line: 339, col: 26, offset: 15339},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 339, col: 26, offset: 15339},
							val:        "+",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 339, col: 30, offset: 15343},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 339, col: 38, offset: 15351},
								expr: &seqExpr{
									pos: position{line: 339, col: 39, offset: 15352},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 339, col: 39, offset: 15352},
											expr: &ruleRefExpr{
												pos:  position{line: 339, col: 40, offset: 15353},
												name: "NEWLINE",
											},
										},
										&notExpr{
											pos: position{line: 339, col: 48, offset: 15361},
											expr: &litMatcher{
												pos:        position{line: 339, col: 49, offset: 15362},
												val:        "+",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 339, col: 53, offset: 15366,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 339, col: 57, offset: 15370},
							val:        "+",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TriplePlusPassthrough",
			pos:  position{line: 343, col: 1, offset: 15465},
			expr: &actionExpr{
				pos: position{line: 343, col: 26, offset: 15490},
				run: (*parser).callonTriplePlusPassthrough1,
				expr: &seqExpr{
					pos: position{line: 343, col: 26, offset: 15490},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 343, col: 26, offset: 15490},
							val:        "+++",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 343, col: 32, offset: 15496},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 343, col: 40, offset: 15504},
								expr: &seqExpr{
									pos: position{line: 343, col: 41, offset: 15505},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 343, col: 41, offset: 15505},
											expr: &litMatcher{
												pos:        position{line: 343, col: 42, offset: 15506},
												val:        "+++",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 343, col: 48, offset: 15512,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 343, col: 52, offset: 15516},
							val:        "+++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "PassthroughMacro",
			pos:  position{line: 347, col: 1, offset: 15613},
			expr: &choiceExpr{
				pos: position{line: 347, col: 21, offset: 15633},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 347, col: 21, offset: 15633},
						name: "SimplePassthroughMacro",
					},
					&ruleRefExpr{
						pos:  position{line: 347, col: 46, offset: 15658},
						name: "PassthroughWithQuotedText",
					},
				},
			},
		},
		{
			name: "SimplePassthroughMacro",
			pos:  position{line: 349, col: 1, offset: 15685},
			expr: &actionExpr{
				pos: position{line: 349, col: 27, offset: 15711},
				run: (*parser).callonSimplePassthroughMacro1,
				expr: &seqExpr{
					pos: position{line: 349, col: 27, offset: 15711},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 349, col: 27, offset: 15711},
							val:        "pass:[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 349, col: 36, offset: 15720},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 349, col: 44, offset: 15728},
								expr: &ruleRefExpr{
									pos:  position{line: 349, col: 45, offset: 15729},
									name: "PassthroughMacroCharacter",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 349, col: 73, offset: 15757},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "PassthroughWithQuotedText",
			pos:  position{line: 353, col: 1, offset: 15847},
			expr: &actionExpr{
				pos: position{line: 353, col: 30, offset: 15876},
				run: (*parser).callonPassthroughWithQuotedText1,
				expr: &seqExpr{
					pos: position{line: 353, col: 30, offset: 15876},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 353, col: 30, offset: 15876},
							val:        "pass:q[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 353, col: 40, offset: 15886},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 353, col: 48, offset: 15894},
								expr: &choiceExpr{
									pos: position{line: 353, col: 49, offset: 15895},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 353, col: 49, offset: 15895},
											name: "QuotedText",
										},
										&ruleRefExpr{
											pos:  position{line: 353, col: 62, offset: 15908},
											name: "PassthroughMacroCharacter",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 353, col: 90, offset: 15936},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "PassthroughMacroCharacter",
			pos:  position{line: 357, col: 1, offset: 16026},
			expr: &seqExpr{
				pos: position{line: 357, col: 31, offset: 16056},
				exprs: []interface{}{
					&notExpr{
						pos: position{line: 357, col: 31, offset: 16056},
						expr: &litMatcher{
							pos:        position{line: 357, col: 32, offset: 16057},
							val:        "]",
							ignoreCase: false,
						},
					},
					&anyMatcher{
						line: 357, col: 36, offset: 16061,
					},
				},
			},
		},
		{
			name: "ExternalLink",
			pos:  position{line: 362, col: 1, offset: 16166},
			expr: &actionExpr{
				pos: position{line: 362, col: 17, offset: 16182},
				run: (*parser).callonExternalLink1,
				expr: &seqExpr{
					pos: position{line: 362, col: 17, offset: 16182},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 362, col: 17, offset: 16182},
							label: "url",
							expr: &seqExpr{
								pos: position{line: 362, col: 22, offset: 16187},
								exprs: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 362, col: 22, offset: 16187},
										name: "URL_SCHEME",
									},
									&ruleRefExpr{
										pos:  position{line: 362, col: 33, offset: 16198},
										name: "URL",
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 362, col: 38, offset: 16203},
							label: "text",
							expr: &zeroOrOneExpr{
								pos: position{line: 362, col: 43, offset: 16208},
								expr: &seqExpr{
									pos: position{line: 362, col: 44, offset: 16209},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 362, col: 44, offset: 16209},
											val:        "[",
											ignoreCase: false,
										},
										&zeroOrMoreExpr{
											pos: position{line: 362, col: 48, offset: 16213},
											expr: &ruleRefExpr{
												pos:  position{line: 362, col: 49, offset: 16214},
												name: "URL_TEXT",
											},
										},
										&litMatcher{
											pos:        position{line: 362, col: 60, offset: 16225},
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
			pos:  position{line: 372, col: 1, offset: 16504},
			expr: &actionExpr{
				pos: position{line: 372, col: 15, offset: 16518},
				run: (*parser).callonBlockImage1,
				expr: &seqExpr{
					pos: position{line: 372, col: 15, offset: 16518},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 372, col: 15, offset: 16518},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 372, col: 26, offset: 16529},
								expr: &ruleRefExpr{
									pos:  position{line: 372, col: 27, offset: 16530},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 372, col: 46, offset: 16549},
							label: "image",
							expr: &ruleRefExpr{
								pos:  position{line: 372, col: 52, offset: 16555},
								name: "BlockImageMacro",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 372, col: 69, offset: 16572},
							expr: &ruleRefExpr{
								pos:  position{line: 372, col: 69, offset: 16572},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 372, col: 73, offset: 16576},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "BlockImageMacro",
			pos:  position{line: 377, col: 1, offset: 16737},
			expr: &actionExpr{
				pos: position{line: 377, col: 20, offset: 16756},
				run: (*parser).callonBlockImageMacro1,
				expr: &seqExpr{
					pos: position{line: 377, col: 20, offset: 16756},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 377, col: 20, offset: 16756},
							val:        "image::",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 377, col: 30, offset: 16766},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 377, col: 36, offset: 16772},
								name: "URL",
							},
						},
						&litMatcher{
							pos:        position{line: 377, col: 41, offset: 16777},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 377, col: 45, offset: 16781},
							label: "attributes",
							expr: &zeroOrOneExpr{
								pos: position{line: 377, col: 57, offset: 16793},
								expr: &ruleRefExpr{
									pos:  position{line: 377, col: 57, offset: 16793},
									name: "URL_TEXT",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 377, col: 68, offset: 16804},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "InlineImage",
			pos:  position{line: 381, col: 1, offset: 16871},
			expr: &actionExpr{
				pos: position{line: 381, col: 16, offset: 16886},
				run: (*parser).callonInlineImage1,
				expr: &labeledExpr{
					pos:   position{line: 381, col: 16, offset: 16886},
					label: "image",
					expr: &ruleRefExpr{
						pos:  position{line: 381, col: 22, offset: 16892},
						name: "InlineImageMacro",
					},
				},
			},
		},
		{
			name: "InlineImageMacro",
			pos:  position{line: 386, col: 1, offset: 17039},
			expr: &actionExpr{
				pos: position{line: 386, col: 21, offset: 17059},
				run: (*parser).callonInlineImageMacro1,
				expr: &seqExpr{
					pos: position{line: 386, col: 21, offset: 17059},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 386, col: 21, offset: 17059},
							val:        "image:",
							ignoreCase: false,
						},
						&notExpr{
							pos: position{line: 386, col: 30, offset: 17068},
							expr: &litMatcher{
								pos:        position{line: 386, col: 31, offset: 17069},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 386, col: 35, offset: 17073},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 386, col: 41, offset: 17079},
								name: "URL",
							},
						},
						&litMatcher{
							pos:        position{line: 386, col: 46, offset: 17084},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 386, col: 50, offset: 17088},
							label: "attributes",
							expr: &zeroOrOneExpr{
								pos: position{line: 386, col: 62, offset: 17100},
								expr: &ruleRefExpr{
									pos:  position{line: 386, col: 62, offset: 17100},
									name: "URL_TEXT",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 386, col: 73, offset: 17111},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "DelimitedBlock",
			pos:  position{line: 393, col: 1, offset: 17441},
			expr: &ruleRefExpr{
				pos:  position{line: 393, col: 19, offset: 17459},
				name: "FencedBlock",
			},
		},
		{
			name: "FencedBlock",
			pos:  position{line: 395, col: 1, offset: 17473},
			expr: &actionExpr{
				pos: position{line: 395, col: 16, offset: 17488},
				run: (*parser).callonFencedBlock1,
				expr: &seqExpr{
					pos: position{line: 395, col: 16, offset: 17488},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 395, col: 16, offset: 17488},
							name: "FencedBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 395, col: 37, offset: 17509},
							expr: &ruleRefExpr{
								pos:  position{line: 395, col: 37, offset: 17509},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 395, col: 41, offset: 17513},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 395, col: 49, offset: 17521},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 395, col: 58, offset: 17530},
								name: "FencedBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 395, col: 78, offset: 17550},
							name: "FencedBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 395, col: 99, offset: 17571},
							expr: &ruleRefExpr{
								pos:  position{line: 395, col: 99, offset: 17571},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 395, col: 103, offset: 17575},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "FencedBlockDelimiter",
			pos:  position{line: 399, col: 1, offset: 17663},
			expr: &litMatcher{
				pos:        position{line: 399, col: 25, offset: 17687},
				val:        "```",
				ignoreCase: false,
			},
		},
		{
			name: "FencedBlockContent",
			pos:  position{line: 401, col: 1, offset: 17694},
			expr: &labeledExpr{
				pos:   position{line: 401, col: 23, offset: 17716},
				label: "content",
				expr: &zeroOrMoreExpr{
					pos: position{line: 401, col: 31, offset: 17724},
					expr: &seqExpr{
						pos: position{line: 401, col: 32, offset: 17725},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 401, col: 32, offset: 17725},
								expr: &ruleRefExpr{
									pos:  position{line: 401, col: 33, offset: 17726},
									name: "FencedBlockDelimiter",
								},
							},
							&anyMatcher{
								line: 401, col: 54, offset: 17747,
							},
						},
					},
				},
			},
		},
		{
			name: "LiteralBlock",
			pos:  position{line: 406, col: 1, offset: 18020},
			expr: &choiceExpr{
				pos: position{line: 406, col: 17, offset: 18036},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 406, col: 17, offset: 18036},
						name: "ParagraphWithSpaces",
					},
					&ruleRefExpr{
						pos:  position{line: 406, col: 39, offset: 18058},
						name: "ParagraphWithLiteralBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 406, col: 76, offset: 18095},
						name: "ParagraphWithLiteralAttribute",
					},
				},
			},
		},
		{
			name: "ParagraphWithSpaces",
			pos:  position{line: 409, col: 1, offset: 18190},
			expr: &actionExpr{
				pos: position{line: 409, col: 24, offset: 18213},
				run: (*parser).callonParagraphWithSpaces1,
				expr: &seqExpr{
					pos: position{line: 409, col: 24, offset: 18213},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 409, col: 24, offset: 18213},
							label: "spaces",
							expr: &oneOrMoreExpr{
								pos: position{line: 409, col: 32, offset: 18221},
								expr: &ruleRefExpr{
									pos:  position{line: 409, col: 32, offset: 18221},
									name: "WS",
								},
							},
						},
						&notExpr{
							pos: position{line: 409, col: 37, offset: 18226},
							expr: &ruleRefExpr{
								pos:  position{line: 409, col: 38, offset: 18227},
								name: "NEWLINE",
							},
						},
						&labeledExpr{
							pos:   position{line: 409, col: 46, offset: 18235},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 409, col: 55, offset: 18244},
								name: "LiteralBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 409, col: 76, offset: 18265},
							name: "EndOfLiteralBlock",
						},
					},
				},
			},
		},
		{
			name: "LiteralBlockContent",
			pos:  position{line: 414, col: 1, offset: 18446},
			expr: &actionExpr{
				pos: position{line: 414, col: 24, offset: 18469},
				run: (*parser).callonLiteralBlockContent1,
				expr: &labeledExpr{
					pos:   position{line: 414, col: 24, offset: 18469},
					label: "content",
					expr: &oneOrMoreExpr{
						pos: position{line: 414, col: 32, offset: 18477},
						expr: &seqExpr{
							pos: position{line: 414, col: 33, offset: 18478},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 414, col: 33, offset: 18478},
									expr: &seqExpr{
										pos: position{line: 414, col: 35, offset: 18480},
										exprs: []interface{}{
											&ruleRefExpr{
												pos:  position{line: 414, col: 35, offset: 18480},
												name: "NEWLINE",
											},
											&ruleRefExpr{
												pos:  position{line: 414, col: 43, offset: 18488},
												name: "BlankLine",
											},
										},
									},
								},
								&anyMatcher{
									line: 414, col: 54, offset: 18499,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "EndOfLiteralBlock",
			pos:  position{line: 419, col: 1, offset: 18584},
			expr: &choiceExpr{
				pos: position{line: 419, col: 22, offset: 18605},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 419, col: 22, offset: 18605},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 419, col: 22, offset: 18605},
								name: "NEWLINE",
							},
							&ruleRefExpr{
								pos:  position{line: 419, col: 30, offset: 18613},
								name: "BlankLine",
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 419, col: 42, offset: 18625},
						name: "NEWLINE",
					},
					&ruleRefExpr{
						pos:  position{line: 419, col: 52, offset: 18635},
						name: "EOF",
					},
				},
			},
		},
		{
			name: "ParagraphWithLiteralBlockDelimiter",
			pos:  position{line: 422, col: 1, offset: 18695},
			expr: &actionExpr{
				pos: position{line: 422, col: 39, offset: 18733},
				run: (*parser).callonParagraphWithLiteralBlockDelimiter1,
				expr: &seqExpr{
					pos: position{line: 422, col: 39, offset: 18733},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 422, col: 39, offset: 18733},
							name: "LiteralBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 422, col: 61, offset: 18755},
							expr: &ruleRefExpr{
								pos:  position{line: 422, col: 61, offset: 18755},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 422, col: 65, offset: 18759},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 422, col: 73, offset: 18767},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 422, col: 81, offset: 18775},
								expr: &seqExpr{
									pos: position{line: 422, col: 82, offset: 18776},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 422, col: 82, offset: 18776},
											expr: &ruleRefExpr{
												pos:  position{line: 422, col: 83, offset: 18777},
												name: "LiteralBlockDelimiter",
											},
										},
										&anyMatcher{
											line: 422, col: 105, offset: 18799,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 422, col: 109, offset: 18803},
							name: "LiteralBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 422, col: 131, offset: 18825},
							expr: &ruleRefExpr{
								pos:  position{line: 422, col: 131, offset: 18825},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 422, col: 135, offset: 18829},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "LiteralBlockDelimiter",
			pos:  position{line: 426, col: 1, offset: 18913},
			expr: &litMatcher{
				pos:        position{line: 426, col: 26, offset: 18938},
				val:        "....",
				ignoreCase: false,
			},
		},
		{
			name: "ParagraphWithLiteralAttribute",
			pos:  position{line: 429, col: 1, offset: 19000},
			expr: &actionExpr{
				pos: position{line: 429, col: 34, offset: 19033},
				run: (*parser).callonParagraphWithLiteralAttribute1,
				expr: &seqExpr{
					pos: position{line: 429, col: 34, offset: 19033},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 429, col: 34, offset: 19033},
							val:        "[literal]",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 429, col: 46, offset: 19045},
							expr: &ruleRefExpr{
								pos:  position{line: 429, col: 46, offset: 19045},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 429, col: 50, offset: 19049},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 429, col: 58, offset: 19057},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 429, col: 67, offset: 19066},
								name: "LiteralBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 429, col: 88, offset: 19087},
							name: "EndOfLiteralBlock",
						},
					},
				},
			},
		},
		{
			name: "ElementAttribute",
			pos:  position{line: 436, col: 1, offset: 19299},
			expr: &labeledExpr{
				pos:   position{line: 436, col: 21, offset: 19319},
				label: "meta",
				expr: &choiceExpr{
					pos: position{line: 436, col: 27, offset: 19325},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 436, col: 27, offset: 19325},
							name: "ElementLink",
						},
						&ruleRefExpr{
							pos:  position{line: 436, col: 41, offset: 19339},
							name: "ElementID",
						},
						&ruleRefExpr{
							pos:  position{line: 436, col: 53, offset: 19351},
							name: "ElementTitle",
						},
					},
				},
			},
		},
		{
			name: "ElementLink",
			pos:  position{line: 439, col: 1, offset: 19422},
			expr: &actionExpr{
				pos: position{line: 439, col: 16, offset: 19437},
				run: (*parser).callonElementLink1,
				expr: &seqExpr{
					pos: position{line: 439, col: 16, offset: 19437},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 439, col: 16, offset: 19437},
							val:        "[",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 439, col: 20, offset: 19441},
							expr: &ruleRefExpr{
								pos:  position{line: 439, col: 20, offset: 19441},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 439, col: 24, offset: 19445},
							val:        "link",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 439, col: 31, offset: 19452},
							expr: &ruleRefExpr{
								pos:  position{line: 439, col: 31, offset: 19452},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 439, col: 35, offset: 19456},
							val:        "=",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 439, col: 39, offset: 19460},
							expr: &ruleRefExpr{
								pos:  position{line: 439, col: 39, offset: 19460},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 439, col: 43, offset: 19464},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 439, col: 48, offset: 19469},
								name: "URL",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 439, col: 52, offset: 19473},
							expr: &ruleRefExpr{
								pos:  position{line: 439, col: 52, offset: 19473},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 439, col: 56, offset: 19477},
							val:        "]",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 439, col: 60, offset: 19481},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ElementID",
			pos:  position{line: 444, col: 1, offset: 19591},
			expr: &actionExpr{
				pos: position{line: 444, col: 14, offset: 19604},
				run: (*parser).callonElementID1,
				expr: &seqExpr{
					pos: position{line: 444, col: 14, offset: 19604},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 444, col: 14, offset: 19604},
							val:        "[",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 444, col: 18, offset: 19608},
							expr: &ruleRefExpr{
								pos:  position{line: 444, col: 18, offset: 19608},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 444, col: 22, offset: 19612},
							val:        "#",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 444, col: 26, offset: 19616},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 444, col: 30, offset: 19620},
								name: "ID",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 444, col: 34, offset: 19624},
							expr: &ruleRefExpr{
								pos:  position{line: 444, col: 34, offset: 19624},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 444, col: 38, offset: 19628},
							val:        "]",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 444, col: 42, offset: 19632},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ElementTitle",
			pos:  position{line: 450, col: 1, offset: 19826},
			expr: &actionExpr{
				pos: position{line: 450, col: 17, offset: 19842},
				run: (*parser).callonElementTitle1,
				expr: &seqExpr{
					pos: position{line: 450, col: 17, offset: 19842},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 450, col: 17, offset: 19842},
							val:        ".",
							ignoreCase: false,
						},
						&notExpr{
							pos: position{line: 450, col: 21, offset: 19846},
							expr: &litMatcher{
								pos:        position{line: 450, col: 22, offset: 19847},
								val:        ".",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 450, col: 26, offset: 19851},
							expr: &ruleRefExpr{
								pos:  position{line: 450, col: 27, offset: 19852},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 450, col: 30, offset: 19855},
							label: "title",
							expr: &oneOrMoreExpr{
								pos: position{line: 450, col: 36, offset: 19861},
								expr: &seqExpr{
									pos: position{line: 450, col: 37, offset: 19862},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 450, col: 37, offset: 19862},
											expr: &ruleRefExpr{
												pos:  position{line: 450, col: 38, offset: 19863},
												name: "NEWLINE",
											},
										},
										&anyMatcher{
											line: 450, col: 46, offset: 19871,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 450, col: 50, offset: 19875},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "BlankLine",
			pos:  position{line: 457, col: 1, offset: 20045},
			expr: &actionExpr{
				pos: position{line: 457, col: 14, offset: 20058},
				run: (*parser).callonBlankLine1,
				expr: &seqExpr{
					pos: position{line: 457, col: 14, offset: 20058},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 457, col: 14, offset: 20058},
							expr: &ruleRefExpr{
								pos:  position{line: 457, col: 15, offset: 20059},
								name: "EOF",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 457, col: 19, offset: 20063},
							expr: &ruleRefExpr{
								pos:  position{line: 457, col: 19, offset: 20063},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 457, col: 23, offset: 20067},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "Characters",
			pos:  position{line: 464, col: 1, offset: 20214},
			expr: &actionExpr{
				pos: position{line: 464, col: 15, offset: 20228},
				run: (*parser).callonCharacters1,
				expr: &oneOrMoreExpr{
					pos: position{line: 464, col: 15, offset: 20228},
					expr: &seqExpr{
						pos: position{line: 464, col: 16, offset: 20229},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 464, col: 16, offset: 20229},
								expr: &ruleRefExpr{
									pos:  position{line: 464, col: 17, offset: 20230},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 464, col: 25, offset: 20238},
								expr: &ruleRefExpr{
									pos:  position{line: 464, col: 26, offset: 20239},
									name: "WS",
								},
							},
							&anyMatcher{
								line: 464, col: 29, offset: 20242,
							},
						},
					},
				},
			},
		},
		{
			name: "URL",
			pos:  position{line: 469, col: 1, offset: 20283},
			expr: &actionExpr{
				pos: position{line: 469, col: 8, offset: 20290},
				run: (*parser).callonURL1,
				expr: &oneOrMoreExpr{
					pos: position{line: 469, col: 8, offset: 20290},
					expr: &seqExpr{
						pos: position{line: 469, col: 9, offset: 20291},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 469, col: 9, offset: 20291},
								expr: &ruleRefExpr{
									pos:  position{line: 469, col: 10, offset: 20292},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 469, col: 18, offset: 20300},
								expr: &ruleRefExpr{
									pos:  position{line: 469, col: 19, offset: 20301},
									name: "WS",
								},
							},
							&notExpr{
								pos: position{line: 469, col: 22, offset: 20304},
								expr: &litMatcher{
									pos:        position{line: 469, col: 23, offset: 20305},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 469, col: 27, offset: 20309},
								expr: &litMatcher{
									pos:        position{line: 469, col: 28, offset: 20310},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 469, col: 32, offset: 20314,
							},
						},
					},
				},
			},
		},
		{
			name: "ID",
			pos:  position{line: 473, col: 1, offset: 20354},
			expr: &actionExpr{
				pos: position{line: 473, col: 7, offset: 20360},
				run: (*parser).callonID1,
				expr: &oneOrMoreExpr{
					pos: position{line: 473, col: 7, offset: 20360},
					expr: &seqExpr{
						pos: position{line: 473, col: 8, offset: 20361},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 473, col: 8, offset: 20361},
								expr: &ruleRefExpr{
									pos:  position{line: 473, col: 9, offset: 20362},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 473, col: 17, offset: 20370},
								expr: &ruleRefExpr{
									pos:  position{line: 473, col: 18, offset: 20371},
									name: "WS",
								},
							},
							&notExpr{
								pos: position{line: 473, col: 21, offset: 20374},
								expr: &litMatcher{
									pos:        position{line: 473, col: 22, offset: 20375},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 473, col: 26, offset: 20379},
								expr: &litMatcher{
									pos:        position{line: 473, col: 27, offset: 20380},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 473, col: 31, offset: 20384,
							},
						},
					},
				},
			},
		},
		{
			name: "URL_TEXT",
			pos:  position{line: 477, col: 1, offset: 20424},
			expr: &actionExpr{
				pos: position{line: 477, col: 13, offset: 20436},
				run: (*parser).callonURL_TEXT1,
				expr: &oneOrMoreExpr{
					pos: position{line: 477, col: 13, offset: 20436},
					expr: &seqExpr{
						pos: position{line: 477, col: 14, offset: 20437},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 477, col: 14, offset: 20437},
								expr: &ruleRefExpr{
									pos:  position{line: 477, col: 15, offset: 20438},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 477, col: 23, offset: 20446},
								expr: &litMatcher{
									pos:        position{line: 477, col: 24, offset: 20447},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 477, col: 28, offset: 20451},
								expr: &litMatcher{
									pos:        position{line: 477, col: 29, offset: 20452},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 477, col: 33, offset: 20456,
							},
						},
					},
				},
			},
		},
		{
			name: "URL_SCHEME",
			pos:  position{line: 481, col: 1, offset: 20496},
			expr: &choiceExpr{
				pos: position{line: 481, col: 15, offset: 20510},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 481, col: 15, offset: 20510},
						val:        "http://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 481, col: 27, offset: 20522},
						val:        "https://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 481, col: 40, offset: 20535},
						val:        "ftp://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 481, col: 51, offset: 20546},
						val:        "irc://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 481, col: 62, offset: 20557},
						val:        "mailto:",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "DIGIT",
			pos:  position{line: 483, col: 1, offset: 20568},
			expr: &charClassMatcher{
				pos:        position{line: 483, col: 13, offset: 20580},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NEWLINE",
			pos:  position{line: 485, col: 1, offset: 20587},
			expr: &choiceExpr{
				pos: position{line: 485, col: 13, offset: 20599},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 485, col: 13, offset: 20599},
						val:        "\r\n",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 485, col: 22, offset: 20608},
						val:        "\r",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 485, col: 29, offset: 20615},
						val:        "\n",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "WS",
			pos:  position{line: 487, col: 1, offset: 20621},
			expr: &choiceExpr{
				pos: position{line: 487, col: 13, offset: 20633},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 487, col: 13, offset: 20633},
						val:        " ",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 487, col: 19, offset: 20639},
						run: (*parser).callonWS3,
						expr: &litMatcher{
							pos:        position{line: 487, col: 19, offset: 20639},
							val:        "\t",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 491, col: 1, offset: 20684},
			expr: &notExpr{
				pos: position{line: 491, col: 13, offset: 20696},
				expr: &anyMatcher{
					line: 491, col: 14, offset: 20697,
				},
			},
		},
		{
			name: "EOL",
			pos:  position{line: 493, col: 1, offset: 20700},
			expr: &choiceExpr{
				pos: position{line: 493, col: 13, offset: 20712},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 493, col: 13, offset: 20712},
						name: "NEWLINE",
					},
					&ruleRefExpr{
						pos:  position{line: 493, col: 23, offset: 20722},
						name: "EOF",
					},
				},
			},
		},
	},
}

func (c *current) onDocument1(frontMatter, documentHeader, blocks interface{}) (interface{}, error) {
	return types.NewDocument(frontMatter, documentHeader, blocks.([]interface{}))
}

func (p *parser) callonDocument1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocument1(stack["frontMatter"], stack["documentHeader"], stack["blocks"])
}

func (c *current) onDocumentBlocks7(content interface{}) (interface{}, error) {
	return content, nil
}

func (p *parser) callonDocumentBlocks7() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocumentBlocks7(stack["content"])
}

func (c *current) onPreamble1(elements interface{}) (interface{}, error) {
	return types.NewPreamble(elements.([]interface{}))
}

func (p *parser) callonPreamble1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onPreamble1(stack["elements"])
}

func (c *current) onFrontMatter1(content interface{}) (interface{}, error) {
	return types.NewYamlFrontMatter(content.(string))
}

func (p *parser) callonFrontMatter1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFrontMatter1(stack["content"])
}

func (c *current) onYamlFrontMatterContent1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonYamlFrontMatterContent1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onYamlFrontMatterContent1()
}

func (c *current) onDocumentHeader1(header, authors, revision, otherAttributes interface{}) (interface{}, error) {

	return types.NewDocumentHeader(header, authors, revision, otherAttributes.([]interface{}))
}

func (p *parser) callonDocumentHeader1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocumentHeader1(stack["header"], stack["authors"], stack["revision"], stack["otherAttributes"])
}

func (c *current) onDocumentTitle1(attributes, level, content interface{}) (interface{}, error) {

	return types.NewSectionTitle(content.(*types.InlineContent), attributes.([]interface{}))
}

func (p *parser) callonDocumentTitle1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocumentTitle1(stack["attributes"], stack["level"], stack["content"])
}

func (c *current) onDocumentAuthorsInlineForm1(authors interface{}) (interface{}, error) {
	return types.NewDocumentAuthors(authors.([]interface{}))
}

func (p *parser) callonDocumentAuthorsInlineForm1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocumentAuthorsInlineForm1(stack["authors"])
}

func (c *current) onDocumentAuthorsAttributeForm1(author interface{}) (interface{}, error) {
	return []*types.DocumentAuthor{author.(*types.DocumentAuthor)}, nil
}

func (p *parser) callonDocumentAuthorsAttributeForm1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocumentAuthorsAttributeForm1(stack["author"])
}

func (c *current) onDocumentAuthor1(namePart1, namePart2, namePart3, email interface{}) (interface{}, error) {
	return types.NewDocumentAuthor(namePart1, namePart2, namePart3, email)
}

func (p *parser) callonDocumentAuthor1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocumentAuthor1(stack["namePart1"], stack["namePart2"], stack["namePart3"], stack["email"])
}

func (c *current) onDocumentRevision1(revnumber, revdate, revremark interface{}) (interface{}, error) {
	return types.NewDocumentRevision(revnumber, revdate, revremark)
}

func (p *parser) callonDocumentRevision1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocumentRevision1(stack["revnumber"], stack["revdate"], stack["revremark"])
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

func (c *current) onDocumentAttributeResetWithSectionTitleBangSymbol1(name interface{}) (interface{}, error) {
	return types.NewDocumentAttributeReset(name.([]interface{}))
}

func (p *parser) callonDocumentAttributeResetWithSectionTitleBangSymbol1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocumentAttributeResetWithSectionTitleBangSymbol1(stack["name"])
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

func (c *current) onSection11(header, elements interface{}) (interface{}, error) {
	return types.NewSection(1, header.(*types.SectionTitle), elements.([]interface{}))
}

func (p *parser) callonSection11() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection11(stack["header"], stack["elements"])
}

func (c *current) onSection1Block1(content interface{}) (interface{}, error) {
	return content.(types.DocElement), nil
}

func (p *parser) callonSection1Block1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection1Block1(stack["content"])
}

func (c *current) onSection21(header, elements interface{}) (interface{}, error) {
	return types.NewSection(2, header.(*types.SectionTitle), elements.([]interface{}))
}

func (p *parser) callonSection21() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection21(stack["header"], stack["elements"])
}

func (c *current) onSection2Block1(content interface{}) (interface{}, error) {
	return content.(types.DocElement), nil
}

func (p *parser) callonSection2Block1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection2Block1(stack["content"])
}

func (c *current) onSection31(header, elements interface{}) (interface{}, error) {
	return types.NewSection(3, header.(*types.SectionTitle), elements.([]interface{}))
}

func (p *parser) callonSection31() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection31(stack["header"], stack["elements"])
}

func (c *current) onSection3Block1(content interface{}) (interface{}, error) {
	return content.(types.DocElement), nil
}

func (p *parser) callonSection3Block1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection3Block1(stack["content"])
}

func (c *current) onSection41(header, elements interface{}) (interface{}, error) {
	return types.NewSection(4, header.(*types.SectionTitle), elements.([]interface{}))
}

func (p *parser) callonSection41() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection41(stack["header"], stack["elements"])
}

func (c *current) onSection4Block1(content interface{}) (interface{}, error) {
	return content.(types.DocElement), nil
}

func (p *parser) callonSection4Block1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection4Block1(stack["content"])
}

func (c *current) onSection51(header, elements interface{}) (interface{}, error) {
	return types.NewSection(5, header.(*types.SectionTitle), elements.([]interface{}))
}

func (p *parser) callonSection51() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection51(stack["header"], stack["elements"])
}

func (c *current) onSection5Block1(content interface{}) (interface{}, error) {
	return content.(types.DocElement), nil
}

func (p *parser) callonSection5Block1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection5Block1(stack["content"])
}

func (c *current) onSection1Title1(attributes, level, content interface{}) (interface{}, error) {

	return types.NewSectionTitle(content.(*types.InlineContent), attributes.([]interface{}))
}

func (p *parser) callonSection1Title1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection1Title1(stack["attributes"], stack["level"], stack["content"])
}

func (c *current) onSection2Title1(attributes, level, content interface{}) (interface{}, error) {
	return types.NewSectionTitle(content.(*types.InlineContent), attributes.([]interface{}))
}

func (p *parser) callonSection2Title1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection2Title1(stack["attributes"], stack["level"], stack["content"])
}

func (c *current) onSection3Title1(attributes, level, content interface{}) (interface{}, error) {
	return types.NewSectionTitle(content.(*types.InlineContent), attributes.([]interface{}))
}

func (p *parser) callonSection3Title1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection3Title1(stack["attributes"], stack["level"], stack["content"])
}

func (c *current) onSection4Title1(attributes, level, content interface{}) (interface{}, error) {
	return types.NewSectionTitle(content.(*types.InlineContent), attributes.([]interface{}))
}

func (p *parser) callonSection4Title1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection4Title1(stack["attributes"], stack["level"], stack["content"])
}

func (c *current) onSection5Title1(attributes, level, content interface{}) (interface{}, error) {
	return types.NewSectionTitle(content.(*types.InlineContent), attributes.([]interface{}))
}

func (p *parser) callonSection5Title1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection5Title1(stack["attributes"], stack["level"], stack["content"])
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

	return types.NewListItemContent(lines.([]interface{}))
}

func (p *parser) callonListItemContent1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onListItemContent1(stack["lines"])
}

func (c *current) onParagraph1(attributes, lines interface{}) (interface{}, error) {
	return types.NewParagraph(lines.([]interface{}), attributes.([]interface{}))
}

func (p *parser) callonParagraph1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onParagraph1(stack["attributes"], stack["lines"])
}

func (c *current) onInlineContent1(elements interface{}) (interface{}, error) {
	// needs an "EOL" but does not consume it here.
	return types.NewInlineContent(elements.([]interface{}))
}

func (p *parser) callonInlineContent1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInlineContent1(stack["elements"])
}

func (c *current) onBoldTextSimplePunctuation1(content interface{}) (interface{}, error) {
	return types.NewQuotedText(types.Bold, content.([]interface{}))
}

func (p *parser) callonBoldTextSimplePunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBoldTextSimplePunctuation1(stack["content"])
}

func (c *current) onBoldTextDoublePunctuation1(content interface{}) (interface{}, error) {
	return types.NewQuotedText(types.Bold, content.([]interface{}))
}

func (p *parser) callonBoldTextDoublePunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBoldTextDoublePunctuation1(stack["content"])
}

func (c *current) onBoldTextUnbalancedPunctuation1(content interface{}) (interface{}, error) {
	// unbalanced `**` vs `*` punctuation
	result := append([]interface{}{"*"}, content.([]interface{}))
	return types.NewQuotedText(types.Bold, result)
}

func (p *parser) callonBoldTextUnbalancedPunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBoldTextUnbalancedPunctuation1(stack["content"])
}

func (c *current) onEscapedBoldTextSimplePunctuation1(backslashes, content interface{}) (interface{}, error) {
	return types.NewEscapedQuotedText(backslashes.([]interface{}), "*", content.([]interface{}))
}

func (p *parser) callonEscapedBoldTextSimplePunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEscapedBoldTextSimplePunctuation1(stack["backslashes"], stack["content"])
}

func (c *current) onEscapedBoldTextDoublePunctuation1(backslashes, content interface{}) (interface{}, error) {
	return types.NewEscapedQuotedText(backslashes.([]interface{}), "**", content.([]interface{}))
}

func (p *parser) callonEscapedBoldTextDoublePunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEscapedBoldTextDoublePunctuation1(stack["backslashes"], stack["content"])
}

func (c *current) onEscapedBoldTextUnbalancedPunctuation1(backslashes, content interface{}) (interface{}, error) {
	// unbalanced `**` vs `*` punctuation
	result := append([]interface{}{"*"}, content.([]interface{}))
	return types.NewEscapedQuotedText(backslashes.([]interface{}), "*", result)
}

func (p *parser) callonEscapedBoldTextUnbalancedPunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEscapedBoldTextUnbalancedPunctuation1(stack["backslashes"], stack["content"])
}

func (c *current) onItalicTextSimplePunctuation1(content interface{}) (interface{}, error) {
	return types.NewQuotedText(types.Italic, content.([]interface{}))
}

func (p *parser) callonItalicTextSimplePunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onItalicTextSimplePunctuation1(stack["content"])
}

func (c *current) onItalicTextDoublePunctuation1(content interface{}) (interface{}, error) {
	return types.NewQuotedText(types.Italic, content.([]interface{}))
}

func (p *parser) callonItalicTextDoublePunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onItalicTextDoublePunctuation1(stack["content"])
}

func (c *current) onItalicTextUnbalancedPunctuation1(content interface{}) (interface{}, error) {
	// unbalanced `__` vs `_` punctuation
	result := append([]interface{}{"_"}, content.([]interface{}))
	return types.NewQuotedText(types.Italic, result)
}

func (p *parser) callonItalicTextUnbalancedPunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onItalicTextUnbalancedPunctuation1(stack["content"])
}

func (c *current) onEscapedItalicTextSimplePunctuation1(backslashes, content interface{}) (interface{}, error) {
	return types.NewEscapedQuotedText(backslashes.([]interface{}), "_", content.([]interface{}))
}

func (p *parser) callonEscapedItalicTextSimplePunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEscapedItalicTextSimplePunctuation1(stack["backslashes"], stack["content"])
}

func (c *current) onEscapedItalicTextDoublePunctuation1(backslashes, content interface{}) (interface{}, error) {
	return types.NewEscapedQuotedText(backslashes.([]interface{}), "__", content.([]interface{}))
}

func (p *parser) callonEscapedItalicTextDoublePunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEscapedItalicTextDoublePunctuation1(stack["backslashes"], stack["content"])
}

func (c *current) onEscapedItalicTextUnbalancedPunctuation1(backslashes, content interface{}) (interface{}, error) {
	// unbalanced `__` vs `_` punctuation
	result := append([]interface{}{"_"}, content.([]interface{}))
	return types.NewEscapedQuotedText(backslashes.([]interface{}), "_", result)
}

func (p *parser) callonEscapedItalicTextUnbalancedPunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEscapedItalicTextUnbalancedPunctuation1(stack["backslashes"], stack["content"])
}

func (c *current) onMonospaceTextSimplePunctuation1(content interface{}) (interface{}, error) {
	return types.NewQuotedText(types.Monospace, content.([]interface{}))
}

func (p *parser) callonMonospaceTextSimplePunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMonospaceTextSimplePunctuation1(stack["content"])
}

func (c *current) onMonospaceTextDoublePunctuation1(content interface{}) (interface{}, error) {
	return types.NewQuotedText(types.Monospace, content.([]interface{}))
}

func (p *parser) callonMonospaceTextDoublePunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMonospaceTextDoublePunctuation1(stack["content"])
}

func (c *current) onMonospaceTextUnbalancedPunctuation1(content interface{}) (interface{}, error) {
	// unbalanced "``" vs "`" punctuation
	result := append([]interface{}{"`"}, content.([]interface{}))
	return types.NewQuotedText(types.Monospace, result)
}

func (p *parser) callonMonospaceTextUnbalancedPunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMonospaceTextUnbalancedPunctuation1(stack["content"])
}

func (c *current) onEscapedMonospaceTextSimplePunctuation1(backslashes, content interface{}) (interface{}, error) {
	return types.NewEscapedQuotedText(backslashes.([]interface{}), "`", content.([]interface{}))
}

func (p *parser) callonEscapedMonospaceTextSimplePunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEscapedMonospaceTextSimplePunctuation1(stack["backslashes"], stack["content"])
}

func (c *current) onEscapedMonospaceTextDoublePunctuation1(backslashes, content interface{}) (interface{}, error) {
	return types.NewEscapedQuotedText(backslashes.([]interface{}), "``", content.([]interface{}))
}

func (p *parser) callonEscapedMonospaceTextDoublePunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEscapedMonospaceTextDoublePunctuation1(stack["backslashes"], stack["content"])
}

func (c *current) onEscapedMonospaceTextUnbalancedPunctuation1(backslashes, content interface{}) (interface{}, error) {
	// unbalanced "``" vs "`" punctuation
	result := append([]interface{}{"`"}, content.([]interface{}))
	return types.NewEscapedQuotedText(backslashes.([]interface{}), "`", result)
}

func (p *parser) callonEscapedMonospaceTextUnbalancedPunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEscapedMonospaceTextUnbalancedPunctuation1(stack["backslashes"], stack["content"])
}

func (c *current) onCharactersWithQuotePunctuation1() (interface{}, error) {
	// can have "*", "_" or "`" within, maybe because the user inserted another quote, or made an error (extra or missing space, for example)
	return c.text, nil
}

func (p *parser) callonCharactersWithQuotePunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCharactersWithQuotePunctuation1()
}

func (c *current) onSinglePlusPassthrough1(content interface{}) (interface{}, error) {
	return types.NewPassthrough(types.SinglePlusPassthrough, content.([]interface{}))
}

func (p *parser) callonSinglePlusPassthrough1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSinglePlusPassthrough1(stack["content"])
}

func (c *current) onTriplePlusPassthrough1(content interface{}) (interface{}, error) {
	return types.NewPassthrough(types.TriplePlusPassthrough, content.([]interface{}))
}

func (p *parser) callonTriplePlusPassthrough1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onTriplePlusPassthrough1(stack["content"])
}

func (c *current) onSimplePassthroughMacro1(content interface{}) (interface{}, error) {
	return types.NewPassthrough(types.PassthroughMacro, content.([]interface{}))
}

func (p *parser) callonSimplePassthroughMacro1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSimplePassthroughMacro1(stack["content"])
}

func (c *current) onPassthroughWithQuotedText1(content interface{}) (interface{}, error) {
	return types.NewPassthrough(types.PassthroughMacro, content.([]interface{}))
}

func (p *parser) callonPassthroughWithQuotedText1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onPassthroughWithQuotedText1(stack["content"])
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
	return types.NewBlockImage(*image.(*types.ImageMacro), attributes.([]interface{}))
}

func (p *parser) callonBlockImage1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBlockImage1(stack["attributes"], stack["image"])
}

func (c *current) onBlockImageMacro1(path, attributes interface{}) (interface{}, error) {
	return types.NewImageMacro(path.(string), attributes)
}

func (p *parser) callonBlockImageMacro1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBlockImageMacro1(stack["path"], stack["attributes"])
}

func (c *current) onInlineImage1(image interface{}) (interface{}, error) {
	// here we can ignore the blank line in the returned element
	return types.NewInlineImage(*image.(*types.ImageMacro))
}

func (p *parser) callonInlineImage1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInlineImage1(stack["image"])
}

func (c *current) onInlineImageMacro1(path, attributes interface{}) (interface{}, error) {
	return types.NewImageMacro(path.(string), attributes)
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

func (c *current) onBlankLine1() (interface{}, error) {
	return types.NewBlankLine()
}

func (p *parser) callonBlankLine1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBlankLine1()
}

func (c *current) onCharacters1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonCharacters1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCharacters1()
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

	// errInvalidEntrypoint is returned when the specified entrypoint rule
	// does not exit.
	errInvalidEntrypoint = errors.New("invalid entrypoint")

	// errInvalidEncoding is returned when the source is not properly
	// utf8-encoded.
	errInvalidEncoding = errors.New("invalid encoding")

	// errMaxExprCnt is used to signal that the maximum number of
	// expressions have been parsed.
	errMaxExprCnt = errors.New("max number of expresssions parsed")
)

// Option is a function that can set an option on the parser. It returns
// the previous setting as an Option.
type Option func(*parser) Option

// MaxExpressions creates an Option to stop parsing after the provided
// number of expressions have been parsed, if the value is 0 then the parser will
// parse for as many steps as needed (possibly an infinite number).
//
// The default for maxExprCnt is 0.
func MaxExpressions(maxExprCnt uint64) Option {
	return func(p *parser) Option {
		oldMaxExprCnt := p.maxExprCnt
		p.maxExprCnt = maxExprCnt
		return MaxExpressions(oldMaxExprCnt)
	}
}

// Entrypoint creates an Option to set the rule name to use as entrypoint.
// The rule name must have been specified in the -alternate-entrypoints
// if generating the parser with the -optimize-grammar flag, otherwise
// it may have been optimized out. Passing an empty string sets the
// entrypoint to the first rule in the grammar.
//
// The default is to start parsing at the first rule in the grammar.
func Entrypoint(ruleName string) Option {
	return func(p *parser) Option {
		oldEntrypoint := p.entrypoint
		p.entrypoint = ruleName
		if ruleName == "" {
			p.entrypoint = g.rules[0].name
		}
		return Entrypoint(oldEntrypoint)
	}
}

// Statistics adds a user provided Stats struct to the parser to allow
// the user to process the results after the parsing has finished.
// Also the key for the "no match" counter is set.
//
// Example usage:
//
//     input := "input"
//     stats := Stats{}
//     _, err := Parse("input-file", []byte(input), Statistics(&stats, "no match"))
//     if err != nil {
//         log.Panicln(err)
//     }
//     b, err := json.MarshalIndent(stats.ChoiceAltCnt, "", "  ")
//     if err != nil {
//         log.Panicln(err)
//     }
//     fmt.Println(string(b))
//
func Statistics(stats *Stats, choiceNoMatch string) Option {
	return func(p *parser) Option {
		oldStats := p.Stats
		p.Stats = stats
		oldChoiceNoMatch := p.choiceNoMatch
		p.choiceNoMatch = choiceNoMatch
		if p.Stats.ChoiceAltCnt == nil {
			p.Stats.ChoiceAltCnt = make(map[string]map[string]int)
		}
		return Statistics(oldStats, oldChoiceNoMatch)
	}
}

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

// AllowInvalidUTF8 creates an Option to allow invalid UTF-8 bytes.
// Every invalid UTF-8 byte is treated as a utf8.RuneError (U+FFFD)
// by character class matchers and is matched by the any matcher.
// The returned matched value, c.text and c.offset are NOT affected.
//
// The default is false.
func AllowInvalidUTF8(b bool) Option {
	return func(p *parser) Option {
		old := p.allowInvalidUTF8
		p.allowInvalidUTF8 = b
		return AllowInvalidUTF8(old)
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

// InitState creates an Option to set a key to a certain value in
// the global "state" store.
func InitState(key string, value interface{}) Option {
	return func(p *parser) Option {
		old := p.cur.state[key]
		p.cur.state[key] = value
		return InitState(key, old)
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

	// state is a store for arbitrary key,value pairs that the user wants to be
	// tied to the backtracking of the parser.
	// This is always rolled back if a parsing rule fails.
	state storeDict

	// globalStore is a general store for the user to store arbitrary key-value
	// pairs that they need to manage and that they do not want tied to the
	// backtracking of the parser. This is only modified by the user and never
	// rolled back by the parser. It is always up to the user to keep this in a
	// consistent state.
	globalStore storeDict
}

type storeDict map[string]interface{}

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

type recoveryExpr struct {
	pos          position
	expr         interface{}
	recoverExpr  interface{}
	failureLabel []string
}

type seqExpr struct {
	pos   position
	exprs []interface{}
}

type throwExpr struct {
	pos   position
	label string
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

type stateCodeExpr struct {
	pos position
	run func(*parser) error
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
	stats := Stats{
		ChoiceAltCnt: make(map[string]map[string]int),
	}

	p := &parser{
		filename: filename,
		errs:     new(errList),
		data:     b,
		pt:       savepoint{position: position{line: 1}},
		recover:  true,
		cur: current{
			state:       make(storeDict),
			globalStore: make(storeDict),
		},
		maxFailPos:      position{col: 1, line: 1},
		maxFailExpected: make([]string, 0, 20),
		Stats:           &stats,
		// start rule is rule [0] unless an alternate entrypoint is specified
		entrypoint: g.rules[0].name,
		emptyState: make(storeDict),
	}
	p.setOptions(opts)

	if p.maxExprCnt == 0 {
		p.maxExprCnt = math.MaxUint64
	}

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

const choiceNoMatch = -1

// Stats stores some statistics, gathered during parsing
type Stats struct {
	// ExprCnt counts the number of expressions processed during parsing
	// This value is compared to the maximum number of expressions allowed
	// (set by the MaxExpressions option).
	ExprCnt uint64

	// ChoiceAltCnt is used to count for each ordered choice expression,
	// which alternative is used how may times.
	// These numbers allow to optimize the order of the ordered choice expression
	// to increase the performance of the parser
	//
	// The outer key of ChoiceAltCnt is composed of the name of the rule as well
	// as the line and the column of the ordered choice.
	// The inner key of ChoiceAltCnt is the number (one-based) of the matching alternative.
	// For each alternative the number of matches are counted. If an ordered choice does not
	// match, a special counter is incremented. The name of this counter is set with
	// the parser option Statistics.
	// For an alternative to be included in ChoiceAltCnt, it has to match at least once.
	ChoiceAltCnt map[string]map[string]int
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

	// parse fail
	maxFailPos            position
	maxFailExpected       []string
	maxFailInvertExpected bool

	// max number of expressions to be parsed
	maxExprCnt uint64
	// entrypoint for the parser
	entrypoint string

	allowInvalidUTF8 bool

	*Stats

	choiceNoMatch string
	// recovery expression stack, keeps track of the currently available recovery expression, these are traversed in reverse
	recoveryStack []map[string]interface{}

	// emptyState contains an empty storeDict, which is used to optimize cloneState if global "state" store is not used.
	emptyState storeDict
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

// push a recovery expression with its labels to the recoveryStack
func (p *parser) pushRecovery(labels []string, expr interface{}) {
	if cap(p.recoveryStack) == len(p.recoveryStack) {
		// create new empty slot in the stack
		p.recoveryStack = append(p.recoveryStack, nil)
	} else {
		// slice to 1 more
		p.recoveryStack = p.recoveryStack[:len(p.recoveryStack)+1]
	}

	m := make(map[string]interface{}, len(labels))
	for _, fl := range labels {
		m[fl] = expr
	}
	p.recoveryStack[len(p.recoveryStack)-1] = m
}

// pop a recovery expression from the recoveryStack
func (p *parser) popRecovery() {
	// GC that map
	p.recoveryStack[len(p.recoveryStack)-1] = nil

	p.recoveryStack = p.recoveryStack[:len(p.recoveryStack)-1]
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

	if rn == utf8.RuneError && n == 1 { // see utf8.DecodeRune
		if !p.allowInvalidUTF8 {
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

// Cloner is implemented by any value that has a Clone method, which returns a
// copy of the value. This is mainly used for types which are not passed by
// value (e.g map, slice, chan) or structs that contain such types.
//
// This is used in conjunction with the global state feature to create proper
// copies of the state to allow the parser to properly restore the state in
// the case of backtracking.
type Cloner interface {
	Clone() interface{}
}

// clone and return parser current state.
func (p *parser) cloneState() storeDict {
	if p.debug {
		defer p.out(p.in("cloneState"))
	}

	if len(p.cur.state) == 0 {
		if len(p.emptyState) > 0 {
			p.emptyState = make(storeDict)
		}
		return p.emptyState
	}

	state := make(storeDict, len(p.cur.state))
	for k, v := range p.cur.state {
		if c, ok := v.(Cloner); ok {
			state[k] = c.Clone()
		} else {
			state[k] = v
		}
	}
	return state
}

// restore parser current state to the state storeDict.
// every restoreState should applied only one time for every cloned state
func (p *parser) restoreState(state storeDict) {
	if p.debug {
		defer p.out(p.in("restoreState"))
	}
	p.cur.state = state
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

	startRule, ok := p.rules[p.entrypoint]
	if !ok {
		p.addErr(errInvalidEntrypoint)
		return nil, p.errs.err()
	}

	p.read() // advance to first rune
	val, ok = p.parseRule(startRule)
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

	p.ExprCnt++
	if p.ExprCnt > p.maxExprCnt {
		panic(errMaxExprCnt)
	}

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
	case *recoveryExpr:
		val, ok = p.parseRecoveryExpr(expr)
	case *ruleRefExpr:
		val, ok = p.parseRuleRefExpr(expr)
	case *seqExpr:
		val, ok = p.parseSeqExpr(expr)
	case *stateCodeExpr:
		val, ok = p.parseStateCodeExpr(expr)
	case *throwExpr:
		val, ok = p.parseThrowExpr(expr)
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
		state := p.cloneState()
		actVal, err := act.run(p)
		if err != nil {
			p.addErrAt(err, start.position, []string{})
		}
		p.restoreState(state)

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

	state := p.cloneState()

	ok, err := and.run(p)
	if err != nil {
		p.addErr(err)
	}
	p.restoreState(state)

	return nil, ok
}

func (p *parser) parseAndExpr(and *andExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndExpr"))
	}

	pt := p.pt
	state := p.cloneState()
	p.pushV()
	_, ok := p.parseExpr(and.expr)
	p.popV()
	p.restoreState(state)
	p.restore(pt)

	return nil, ok
}

func (p *parser) parseAnyMatcher(any *anyMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAnyMatcher"))
	}

	if p.pt.rn == utf8.RuneError && p.pt.w == 0 {
		// EOF - see utf8.DecodeRune
		p.failAt(false, p.pt.position, ".")
		return nil, false
	}
	start := p.pt
	p.read()
	p.failAt(true, start.position, ".")
	return p.sliceFrom(start), true
}

func (p *parser) parseCharClassMatcher(chr *charClassMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseCharClassMatcher"))
	}

	cur := p.pt.rn
	start := p.pt

	// can't match EOF
	if cur == utf8.RuneError && p.pt.w == 0 { // see utf8.DecodeRune
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

func (p *parser) incChoiceAltCnt(ch *choiceExpr, altI int) {
	choiceIdent := fmt.Sprintf("%s %d:%d", p.rstack[len(p.rstack)-1].name, ch.pos.line, ch.pos.col)
	m := p.ChoiceAltCnt[choiceIdent]
	if m == nil {
		m = make(map[string]int)
		p.ChoiceAltCnt[choiceIdent] = m
	}
	// We increment altI by 1, so the keys do not start at 0
	alt := strconv.Itoa(altI + 1)
	if altI == choiceNoMatch {
		alt = p.choiceNoMatch
	}
	m[alt]++
}

func (p *parser) parseChoiceExpr(ch *choiceExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseChoiceExpr"))
	}

	for altI, alt := range ch.alternatives {
		// dummy assignment to prevent compile error if optimized
		_ = altI

		state := p.cloneState()

		p.pushV()
		val, ok := p.parseExpr(alt)
		p.popV()
		if ok {
			p.incChoiceAltCnt(ch, altI)
			return val, ok
		}
		p.restoreState(state)
	}
	p.incChoiceAltCnt(ch, choiceNoMatch)
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

	state := p.cloneState()

	ok, err := not.run(p)
	if err != nil {
		p.addErr(err)
	}
	p.restoreState(state)

	return nil, !ok
}

func (p *parser) parseNotExpr(not *notExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotExpr"))
	}

	pt := p.pt
	state := p.cloneState()
	p.pushV()
	p.maxFailInvertExpected = !p.maxFailInvertExpected
	_, ok := p.parseExpr(not.expr)
	p.maxFailInvertExpected = !p.maxFailInvertExpected
	p.popV()
	p.restoreState(state)
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

func (p *parser) parseRecoveryExpr(recover *recoveryExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRecoveryExpr (" + strings.Join(recover.failureLabel, ",") + ")"))
	}

	p.pushRecovery(recover.failureLabel, recover.recoverExpr)
	val, ok := p.parseExpr(recover.expr)
	p.popRecovery()

	return val, ok
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
	state := p.cloneState()
	for _, expr := range seq.exprs {
		val, ok := p.parseExpr(expr)
		if !ok {
			p.restoreState(state)
			p.restore(pt)
			return nil, false
		}
		vals = append(vals, val)
	}
	return vals, true
}

func (p *parser) parseStateCodeExpr(state *stateCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseStateCodeExpr"))
	}

	err := state.run(p)
	if err != nil {
		p.addErr(err)
	}
	return nil, true
}

func (p *parser) parseThrowExpr(expr *throwExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseThrowExpr"))
	}

	for i := len(p.recoveryStack) - 1; i >= 0; i-- {
		if recoverExpr, ok := p.recoveryStack[i][expr.label]; ok {
			if val, ok := p.parseExpr(recoverExpr); ok {
				return val, ok
			}
		}
	}

	return nil, false
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
