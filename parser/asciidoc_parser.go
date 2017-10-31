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
							expr: &zeroOrMoreExpr{
								pos: position{line: 35, col: 45, offset: 1248},
								expr: &seqExpr{
									pos: position{line: 35, col: 46, offset: 1249},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 35, col: 46, offset: 1249},
											expr: &ruleRefExpr{
												pos:  position{line: 35, col: 47, offset: 1250},
												name: "YamlFrontMatterToken",
											},
										},
										&anyMatcher{
											line: 35, col: 68, offset: 1271,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 35, col: 72, offset: 1275},
							name: "YamlFrontMatterToken",
						},
					},
				},
			},
		},
		{
			name: "YamlFrontMatterToken",
			pos:  position{line: 39, col: 1, offset: 1362},
			expr: &seqExpr{
				pos: position{line: 39, col: 26, offset: 1387},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 39, col: 26, offset: 1387},
						val:        "---",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 39, col: 32, offset: 1393},
						name: "EOL",
					},
				},
			},
		},
		{
			name: "DocumentHeader",
			pos:  position{line: 45, col: 1, offset: 1582},
			expr: &actionExpr{
				pos: position{line: 45, col: 19, offset: 1600},
				run: (*parser).callonDocumentHeader1,
				expr: &seqExpr{
					pos: position{line: 45, col: 19, offset: 1600},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 45, col: 19, offset: 1600},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 45, col: 27, offset: 1608},
								name: "DocumentTitle",
							},
						},
						&labeledExpr{
							pos:   position{line: 45, col: 42, offset: 1623},
							label: "authors",
							expr: &zeroOrOneExpr{
								pos: position{line: 45, col: 51, offset: 1632},
								expr: &ruleRefExpr{
									pos:  position{line: 45, col: 51, offset: 1632},
									name: "DocumentAuthors",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 45, col: 69, offset: 1650},
							label: "revision",
							expr: &zeroOrOneExpr{
								pos: position{line: 45, col: 79, offset: 1660},
								expr: &ruleRefExpr{
									pos:  position{line: 45, col: 79, offset: 1660},
									name: "DocumentRevision",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 45, col: 98, offset: 1679},
							label: "otherAttributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 45, col: 115, offset: 1696},
								expr: &ruleRefExpr{
									pos:  position{line: 45, col: 115, offset: 1696},
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
			pos:  position{line: 49, col: 1, offset: 1827},
			expr: &actionExpr{
				pos: position{line: 49, col: 18, offset: 1844},
				run: (*parser).callonDocumentTitle1,
				expr: &seqExpr{
					pos: position{line: 49, col: 18, offset: 1844},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 49, col: 18, offset: 1844},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 49, col: 29, offset: 1855},
								expr: &ruleRefExpr{
									pos:  position{line: 49, col: 30, offset: 1856},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 49, col: 49, offset: 1875},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 49, col: 56, offset: 1882},
								val:        "=",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 49, col: 61, offset: 1887},
							expr: &ruleRefExpr{
								pos:  position{line: 49, col: 61, offset: 1887},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 49, col: 65, offset: 1891},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 49, col: 73, offset: 1899},
								name: "InlineContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 49, col: 87, offset: 1913},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthors",
			pos:  position{line: 53, col: 1, offset: 2017},
			expr: &choiceExpr{
				pos: position{line: 53, col: 20, offset: 2036},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 53, col: 20, offset: 2036},
						name: "DocumentAuthorsInlineForm",
					},
					&ruleRefExpr{
						pos:  position{line: 53, col: 48, offset: 2064},
						name: "DocumentAuthorsAttributeForm",
					},
				},
			},
		},
		{
			name: "DocumentAuthorsInlineForm",
			pos:  position{line: 55, col: 1, offset: 2094},
			expr: &actionExpr{
				pos: position{line: 55, col: 30, offset: 2123},
				run: (*parser).callonDocumentAuthorsInlineForm1,
				expr: &seqExpr{
					pos: position{line: 55, col: 30, offset: 2123},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 55, col: 30, offset: 2123},
							expr: &ruleRefExpr{
								pos:  position{line: 55, col: 30, offset: 2123},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 55, col: 34, offset: 2127},
							expr: &litMatcher{
								pos:        position{line: 55, col: 35, offset: 2128},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 55, col: 39, offset: 2132},
							label: "authors",
							expr: &oneOrMoreExpr{
								pos: position{line: 55, col: 48, offset: 2141},
								expr: &ruleRefExpr{
									pos:  position{line: 55, col: 48, offset: 2141},
									name: "DocumentAuthor",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 55, col: 65, offset: 2158},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthorsAttributeForm",
			pos:  position{line: 59, col: 1, offset: 2228},
			expr: &actionExpr{
				pos: position{line: 59, col: 33, offset: 2260},
				run: (*parser).callonDocumentAuthorsAttributeForm1,
				expr: &seqExpr{
					pos: position{line: 59, col: 33, offset: 2260},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 59, col: 33, offset: 2260},
							expr: &ruleRefExpr{
								pos:  position{line: 59, col: 33, offset: 2260},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 59, col: 37, offset: 2264},
							val:        ":author:",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 59, col: 48, offset: 2275},
							label: "author",
							expr: &ruleRefExpr{
								pos:  position{line: 59, col: 56, offset: 2283},
								name: "DocumentAuthor",
							},
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthor",
			pos:  position{line: 63, col: 1, offset: 2376},
			expr: &actionExpr{
				pos: position{line: 63, col: 19, offset: 2394},
				run: (*parser).callonDocumentAuthor1,
				expr: &seqExpr{
					pos: position{line: 63, col: 19, offset: 2394},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 63, col: 19, offset: 2394},
							expr: &ruleRefExpr{
								pos:  position{line: 63, col: 19, offset: 2394},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 63, col: 23, offset: 2398},
							label: "namePart1",
							expr: &ruleRefExpr{
								pos:  position{line: 63, col: 34, offset: 2409},
								name: "DocumentAuthorNamePart",
							},
						},
						&labeledExpr{
							pos:   position{line: 63, col: 58, offset: 2433},
							label: "namePart2",
							expr: &zeroOrOneExpr{
								pos: position{line: 63, col: 68, offset: 2443},
								expr: &ruleRefExpr{
									pos:  position{line: 63, col: 69, offset: 2444},
									name: "DocumentAuthorNamePart",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 63, col: 94, offset: 2469},
							label: "namePart3",
							expr: &zeroOrOneExpr{
								pos: position{line: 63, col: 104, offset: 2479},
								expr: &ruleRefExpr{
									pos:  position{line: 63, col: 105, offset: 2480},
									name: "DocumentAuthorNamePart",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 63, col: 130, offset: 2505},
							label: "email",
							expr: &zeroOrOneExpr{
								pos: position{line: 63, col: 136, offset: 2511},
								expr: &ruleRefExpr{
									pos:  position{line: 63, col: 137, offset: 2512},
									name: "DocumentAuthorEmail",
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 63, col: 159, offset: 2534},
							expr: &ruleRefExpr{
								pos:  position{line: 63, col: 159, offset: 2534},
								name: "WS",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 63, col: 163, offset: 2538},
							expr: &litMatcher{
								pos:        position{line: 63, col: 163, offset: 2538},
								val:        ";",
								ignoreCase: false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 63, col: 168, offset: 2543},
							expr: &ruleRefExpr{
								pos:  position{line: 63, col: 168, offset: 2543},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthorNamePart",
			pos:  position{line: 68, col: 1, offset: 2708},
			expr: &seqExpr{
				pos: position{line: 68, col: 27, offset: 2734},
				exprs: []interface{}{
					&notExpr{
						pos: position{line: 68, col: 27, offset: 2734},
						expr: &litMatcher{
							pos:        position{line: 68, col: 28, offset: 2735},
							val:        "<",
							ignoreCase: false,
						},
					},
					&notExpr{
						pos: position{line: 68, col: 32, offset: 2739},
						expr: &litMatcher{
							pos:        position{line: 68, col: 33, offset: 2740},
							val:        ";",
							ignoreCase: false,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 68, col: 37, offset: 2744},
						name: "Word",
					},
					&zeroOrMoreExpr{
						pos: position{line: 68, col: 42, offset: 2749},
						expr: &ruleRefExpr{
							pos:  position{line: 68, col: 42, offset: 2749},
							name: "WS",
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthorEmail",
			pos:  position{line: 70, col: 1, offset: 2754},
			expr: &seqExpr{
				pos: position{line: 70, col: 24, offset: 2777},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 70, col: 24, offset: 2777},
						val:        "<",
						ignoreCase: false,
					},
					&labeledExpr{
						pos:   position{line: 70, col: 28, offset: 2781},
						label: "email",
						expr: &oneOrMoreExpr{
							pos: position{line: 70, col: 34, offset: 2787},
							expr: &seqExpr{
								pos: position{line: 70, col: 35, offset: 2788},
								exprs: []interface{}{
									&notExpr{
										pos: position{line: 70, col: 35, offset: 2788},
										expr: &litMatcher{
											pos:        position{line: 70, col: 36, offset: 2789},
											val:        ">",
											ignoreCase: false,
										},
									},
									&notExpr{
										pos: position{line: 70, col: 40, offset: 2793},
										expr: &ruleRefExpr{
											pos:  position{line: 70, col: 41, offset: 2794},
											name: "EOL",
										},
									},
									&anyMatcher{
										line: 70, col: 45, offset: 2798,
									},
								},
							},
						},
					},
					&litMatcher{
						pos:        position{line: 70, col: 49, offset: 2802},
						val:        ">",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "DocumentRevision",
			pos:  position{line: 74, col: 1, offset: 2938},
			expr: &actionExpr{
				pos: position{line: 74, col: 21, offset: 2958},
				run: (*parser).callonDocumentRevision1,
				expr: &seqExpr{
					pos: position{line: 74, col: 21, offset: 2958},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 74, col: 21, offset: 2958},
							expr: &ruleRefExpr{
								pos:  position{line: 74, col: 21, offset: 2958},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 74, col: 25, offset: 2962},
							expr: &litMatcher{
								pos:        position{line: 74, col: 26, offset: 2963},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 74, col: 30, offset: 2967},
							label: "revnumber",
							expr: &zeroOrOneExpr{
								pos: position{line: 74, col: 40, offset: 2977},
								expr: &ruleRefExpr{
									pos:  position{line: 74, col: 41, offset: 2978},
									name: "DocumentRevisionNumber",
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 74, col: 66, offset: 3003},
							expr: &litMatcher{
								pos:        position{line: 74, col: 66, offset: 3003},
								val:        ",",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 74, col: 71, offset: 3008},
							label: "revdate",
							expr: &zeroOrOneExpr{
								pos: position{line: 74, col: 79, offset: 3016},
								expr: &ruleRefExpr{
									pos:  position{line: 74, col: 80, offset: 3017},
									name: "DocumentRevisionDate",
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 74, col: 103, offset: 3040},
							expr: &litMatcher{
								pos:        position{line: 74, col: 103, offset: 3040},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 74, col: 108, offset: 3045},
							label: "revremark",
							expr: &zeroOrOneExpr{
								pos: position{line: 74, col: 118, offset: 3055},
								expr: &ruleRefExpr{
									pos:  position{line: 74, col: 119, offset: 3056},
									name: "DocumentRevisionRemark",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 74, col: 144, offset: 3081},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentRevisionNumber",
			pos:  position{line: 79, col: 1, offset: 3254},
			expr: &choiceExpr{
				pos: position{line: 79, col: 27, offset: 3280},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 79, col: 27, offset: 3280},
						exprs: []interface{}{
							&litMatcher{
								pos:        position{line: 79, col: 27, offset: 3280},
								val:        "v",
								ignoreCase: true,
							},
							&ruleRefExpr{
								pos:  position{line: 79, col: 32, offset: 3285},
								name: "DIGIT",
							},
							&zeroOrMoreExpr{
								pos: position{line: 79, col: 39, offset: 3292},
								expr: &seqExpr{
									pos: position{line: 79, col: 40, offset: 3293},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 79, col: 40, offset: 3293},
											expr: &ruleRefExpr{
												pos:  position{line: 79, col: 41, offset: 3294},
												name: "EOL",
											},
										},
										&notExpr{
											pos: position{line: 79, col: 45, offset: 3298},
											expr: &litMatcher{
												pos:        position{line: 79, col: 46, offset: 3299},
												val:        ",",
												ignoreCase: false,
											},
										},
										&notExpr{
											pos: position{line: 79, col: 50, offset: 3303},
											expr: &litMatcher{
												pos:        position{line: 79, col: 51, offset: 3304},
												val:        ":",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 79, col: 55, offset: 3308,
										},
									},
								},
							},
						},
					},
					&seqExpr{
						pos: position{line: 79, col: 61, offset: 3314},
						exprs: []interface{}{
							&zeroOrOneExpr{
								pos: position{line: 79, col: 61, offset: 3314},
								expr: &litMatcher{
									pos:        position{line: 79, col: 61, offset: 3314},
									val:        "v",
									ignoreCase: true,
								},
							},
							&ruleRefExpr{
								pos:  position{line: 79, col: 67, offset: 3320},
								name: "DIGIT",
							},
							&zeroOrMoreExpr{
								pos: position{line: 79, col: 74, offset: 3327},
								expr: &seqExpr{
									pos: position{line: 79, col: 75, offset: 3328},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 79, col: 75, offset: 3328},
											expr: &ruleRefExpr{
												pos:  position{line: 79, col: 76, offset: 3329},
												name: "EOL",
											},
										},
										&notExpr{
											pos: position{line: 79, col: 80, offset: 3333},
											expr: &litMatcher{
												pos:        position{line: 79, col: 81, offset: 3334},
												val:        ",",
												ignoreCase: false,
											},
										},
										&notExpr{
											pos: position{line: 79, col: 85, offset: 3338},
											expr: &litMatcher{
												pos:        position{line: 79, col: 86, offset: 3339},
												val:        ":",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 79, col: 90, offset: 3343,
										},
									},
								},
							},
							&zeroOrMoreExpr{
								pos: position{line: 79, col: 94, offset: 3347},
								expr: &ruleRefExpr{
									pos:  position{line: 79, col: 94, offset: 3347},
									name: "WS",
								},
							},
							&andExpr{
								pos: position{line: 79, col: 98, offset: 3351},
								expr: &litMatcher{
									pos:        position{line: 79, col: 99, offset: 3352},
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
			pos:  position{line: 80, col: 1, offset: 3356},
			expr: &zeroOrMoreExpr{
				pos: position{line: 80, col: 25, offset: 3380},
				expr: &seqExpr{
					pos: position{line: 80, col: 26, offset: 3381},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 80, col: 26, offset: 3381},
							expr: &ruleRefExpr{
								pos:  position{line: 80, col: 27, offset: 3382},
								name: "EOL",
							},
						},
						&notExpr{
							pos: position{line: 80, col: 31, offset: 3386},
							expr: &litMatcher{
								pos:        position{line: 80, col: 32, offset: 3387},
								val:        ":",
								ignoreCase: false,
							},
						},
						&anyMatcher{
							line: 80, col: 36, offset: 3391,
						},
					},
				},
			},
		},
		{
			name: "DocumentRevisionRemark",
			pos:  position{line: 81, col: 1, offset: 3396},
			expr: &zeroOrMoreExpr{
				pos: position{line: 81, col: 27, offset: 3422},
				expr: &seqExpr{
					pos: position{line: 81, col: 28, offset: 3423},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 81, col: 28, offset: 3423},
							expr: &ruleRefExpr{
								pos:  position{line: 81, col: 29, offset: 3424},
								name: "EOL",
							},
						},
						&anyMatcher{
							line: 81, col: 33, offset: 3428,
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclaration",
			pos:  position{line: 86, col: 1, offset: 3548},
			expr: &choiceExpr{
				pos: position{line: 86, col: 33, offset: 3580},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 86, col: 33, offset: 3580},
						name: "DocumentAttributeDeclarationWithNameOnly",
					},
					&ruleRefExpr{
						pos:  position{line: 86, col: 76, offset: 3623},
						name: "DocumentAttributeDeclarationWithNameAndValue",
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclarationWithNameOnly",
			pos:  position{line: 88, col: 1, offset: 3670},
			expr: &actionExpr{
				pos: position{line: 88, col: 45, offset: 3714},
				run: (*parser).callonDocumentAttributeDeclarationWithNameOnly1,
				expr: &seqExpr{
					pos: position{line: 88, col: 45, offset: 3714},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 88, col: 45, offset: 3714},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 88, col: 49, offset: 3718},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 88, col: 55, offset: 3724},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 88, col: 70, offset: 3739},
							val:        ":",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 88, col: 74, offset: 3743},
							expr: &ruleRefExpr{
								pos:  position{line: 88, col: 74, offset: 3743},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 88, col: 78, offset: 3747},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclarationWithNameAndValue",
			pos:  position{line: 92, col: 1, offset: 3832},
			expr: &actionExpr{
				pos: position{line: 92, col: 49, offset: 3880},
				run: (*parser).callonDocumentAttributeDeclarationWithNameAndValue1,
				expr: &seqExpr{
					pos: position{line: 92, col: 49, offset: 3880},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 92, col: 49, offset: 3880},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 92, col: 53, offset: 3884},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 92, col: 59, offset: 3890},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 92, col: 74, offset: 3905},
							val:        ":",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 92, col: 78, offset: 3909},
							expr: &ruleRefExpr{
								pos:  position{line: 92, col: 78, offset: 3909},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 92, col: 82, offset: 3913},
							label: "value",
							expr: &zeroOrMoreExpr{
								pos: position{line: 92, col: 88, offset: 3919},
								expr: &seqExpr{
									pos: position{line: 92, col: 89, offset: 3920},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 92, col: 89, offset: 3920},
											expr: &ruleRefExpr{
												pos:  position{line: 92, col: 90, offset: 3921},
												name: "NEWLINE",
											},
										},
										&anyMatcher{
											line: 92, col: 98, offset: 3929,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 92, col: 102, offset: 3933},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeReset",
			pos:  position{line: 96, col: 1, offset: 4036},
			expr: &choiceExpr{
				pos: position{line: 96, col: 27, offset: 4062},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 96, col: 27, offset: 4062},
						name: "DocumentAttributeResetWithSectionTitleBangSymbol",
					},
					&ruleRefExpr{
						pos:  position{line: 96, col: 78, offset: 4113},
						name: "DocumentAttributeResetWithTrailingBangSymbol",
					},
				},
			},
		},
		{
			name: "DocumentAttributeResetWithSectionTitleBangSymbol",
			pos:  position{line: 98, col: 1, offset: 4159},
			expr: &actionExpr{
				pos: position{line: 98, col: 53, offset: 4211},
				run: (*parser).callonDocumentAttributeResetWithSectionTitleBangSymbol1,
				expr: &seqExpr{
					pos: position{line: 98, col: 53, offset: 4211},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 98, col: 53, offset: 4211},
							val:        ":!",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 98, col: 58, offset: 4216},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 98, col: 64, offset: 4222},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 98, col: 79, offset: 4237},
							val:        ":",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 98, col: 83, offset: 4241},
							expr: &ruleRefExpr{
								pos:  position{line: 98, col: 83, offset: 4241},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 98, col: 87, offset: 4245},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeResetWithTrailingBangSymbol",
			pos:  position{line: 102, col: 1, offset: 4319},
			expr: &actionExpr{
				pos: position{line: 102, col: 49, offset: 4367},
				run: (*parser).callonDocumentAttributeResetWithTrailingBangSymbol1,
				expr: &seqExpr{
					pos: position{line: 102, col: 49, offset: 4367},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 102, col: 49, offset: 4367},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 102, col: 53, offset: 4371},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 102, col: 59, offset: 4377},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 102, col: 74, offset: 4392},
							val:        "!:",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 102, col: 79, offset: 4397},
							expr: &ruleRefExpr{
								pos:  position{line: 102, col: 79, offset: 4397},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 102, col: 83, offset: 4401},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeSubstitution",
			pos:  position{line: 107, col: 1, offset: 4476},
			expr: &actionExpr{
				pos: position{line: 107, col: 34, offset: 4509},
				run: (*parser).callonDocumentAttributeSubstitution1,
				expr: &seqExpr{
					pos: position{line: 107, col: 34, offset: 4509},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 107, col: 34, offset: 4509},
							val:        "{",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 107, col: 38, offset: 4513},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 107, col: 44, offset: 4519},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 107, col: 59, offset: 4534},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "AttributeName",
			pos:  position{line: 114, col: 1, offset: 4788},
			expr: &seqExpr{
				pos: position{line: 114, col: 18, offset: 4805},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 114, col: 19, offset: 4806},
						alternatives: []interface{}{
							&charClassMatcher{
								pos:        position{line: 114, col: 19, offset: 4806},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
							&charClassMatcher{
								pos:        position{line: 114, col: 27, offset: 4814},
								val:        "[a-z]",
								ranges:     []rune{'a', 'z'},
								ignoreCase: false,
								inverted:   false,
							},
							&charClassMatcher{
								pos:        position{line: 114, col: 35, offset: 4822},
								val:        "[0-9]",
								ranges:     []rune{'0', '9'},
								ignoreCase: false,
								inverted:   false,
							},
							&litMatcher{
								pos:        position{line: 114, col: 43, offset: 4830},
								val:        "_",
								ignoreCase: false,
							},
						},
					},
					&zeroOrMoreExpr{
						pos: position{line: 114, col: 48, offset: 4835},
						expr: &choiceExpr{
							pos: position{line: 114, col: 49, offset: 4836},
							alternatives: []interface{}{
								&charClassMatcher{
									pos:        position{line: 114, col: 49, offset: 4836},
									val:        "[A-Z]",
									ranges:     []rune{'A', 'Z'},
									ignoreCase: false,
									inverted:   false,
								},
								&charClassMatcher{
									pos:        position{line: 114, col: 57, offset: 4844},
									val:        "[a-z]",
									ranges:     []rune{'a', 'z'},
									ignoreCase: false,
									inverted:   false,
								},
								&charClassMatcher{
									pos:        position{line: 114, col: 65, offset: 4852},
									val:        "[0-9]",
									ranges:     []rune{'0', '9'},
									ignoreCase: false,
									inverted:   false,
								},
								&litMatcher{
									pos:        position{line: 114, col: 73, offset: 4860},
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
			pos:  position{line: 119, col: 1, offset: 4971},
			expr: &choiceExpr{
				pos: position{line: 119, col: 12, offset: 4982},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 119, col: 12, offset: 4982},
						name: "Section1",
					},
					&ruleRefExpr{
						pos:  position{line: 119, col: 23, offset: 4993},
						name: "Section2",
					},
					&ruleRefExpr{
						pos:  position{line: 119, col: 34, offset: 5004},
						name: "Section3",
					},
					&ruleRefExpr{
						pos:  position{line: 119, col: 45, offset: 5015},
						name: "Section4",
					},
					&ruleRefExpr{
						pos:  position{line: 119, col: 56, offset: 5026},
						name: "Section5",
					},
				},
			},
		},
		{
			name: "Section1",
			pos:  position{line: 122, col: 1, offset: 5037},
			expr: &actionExpr{
				pos: position{line: 122, col: 13, offset: 5049},
				run: (*parser).callonSection11,
				expr: &seqExpr{
					pos: position{line: 122, col: 13, offset: 5049},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 122, col: 13, offset: 5049},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 122, col: 21, offset: 5057},
								name: "Section1Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 122, col: 36, offset: 5072},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 122, col: 46, offset: 5082},
								expr: &ruleRefExpr{
									pos:  position{line: 122, col: 46, offset: 5082},
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
			pos:  position{line: 126, col: 1, offset: 5190},
			expr: &actionExpr{
				pos: position{line: 126, col: 18, offset: 5207},
				run: (*parser).callonSection1Block1,
				expr: &seqExpr{
					pos: position{line: 126, col: 18, offset: 5207},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 126, col: 18, offset: 5207},
							expr: &ruleRefExpr{
								pos:  position{line: 126, col: 19, offset: 5208},
								name: "Section1",
							},
						},
						&labeledExpr{
							pos:   position{line: 126, col: 28, offset: 5217},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 126, col: 37, offset: 5226},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 126, col: 37, offset: 5226},
										name: "Section2",
									},
									&ruleRefExpr{
										pos:  position{line: 126, col: 48, offset: 5237},
										name: "Section3",
									},
									&ruleRefExpr{
										pos:  position{line: 126, col: 59, offset: 5248},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 126, col: 70, offset: 5259},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 126, col: 81, offset: 5270},
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
			pos:  position{line: 130, col: 1, offset: 5335},
			expr: &actionExpr{
				pos: position{line: 130, col: 13, offset: 5347},
				run: (*parser).callonSection21,
				expr: &seqExpr{
					pos: position{line: 130, col: 13, offset: 5347},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 130, col: 13, offset: 5347},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 130, col: 21, offset: 5355},
								name: "Section2Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 130, col: 36, offset: 5370},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 130, col: 46, offset: 5380},
								expr: &ruleRefExpr{
									pos:  position{line: 130, col: 46, offset: 5380},
									name: "Section2Block",
								},
							},
						},
						&andExpr{
							pos: position{line: 130, col: 62, offset: 5396},
							expr: &zeroOrMoreExpr{
								pos: position{line: 130, col: 63, offset: 5397},
								expr: &ruleRefExpr{
									pos:  position{line: 130, col: 64, offset: 5398},
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
			pos:  position{line: 134, col: 1, offset: 5501},
			expr: &actionExpr{
				pos: position{line: 134, col: 18, offset: 5518},
				run: (*parser).callonSection2Block1,
				expr: &seqExpr{
					pos: position{line: 134, col: 18, offset: 5518},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 134, col: 18, offset: 5518},
							expr: &ruleRefExpr{
								pos:  position{line: 134, col: 19, offset: 5519},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 134, col: 28, offset: 5528},
							expr: &ruleRefExpr{
								pos:  position{line: 134, col: 29, offset: 5529},
								name: "Section2",
							},
						},
						&labeledExpr{
							pos:   position{line: 134, col: 38, offset: 5538},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 134, col: 47, offset: 5547},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 134, col: 47, offset: 5547},
										name: "Section3",
									},
									&ruleRefExpr{
										pos:  position{line: 134, col: 58, offset: 5558},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 134, col: 69, offset: 5569},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 134, col: 80, offset: 5580},
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
			pos:  position{line: 138, col: 1, offset: 5645},
			expr: &actionExpr{
				pos: position{line: 138, col: 13, offset: 5657},
				run: (*parser).callonSection31,
				expr: &seqExpr{
					pos: position{line: 138, col: 13, offset: 5657},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 138, col: 13, offset: 5657},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 138, col: 21, offset: 5665},
								name: "Section3Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 138, col: 36, offset: 5680},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 138, col: 46, offset: 5690},
								expr: &ruleRefExpr{
									pos:  position{line: 138, col: 46, offset: 5690},
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
			pos:  position{line: 142, col: 1, offset: 5798},
			expr: &actionExpr{
				pos: position{line: 142, col: 18, offset: 5815},
				run: (*parser).callonSection3Block1,
				expr: &seqExpr{
					pos: position{line: 142, col: 18, offset: 5815},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 142, col: 18, offset: 5815},
							expr: &ruleRefExpr{
								pos:  position{line: 142, col: 19, offset: 5816},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 142, col: 28, offset: 5825},
							expr: &ruleRefExpr{
								pos:  position{line: 142, col: 29, offset: 5826},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 142, col: 38, offset: 5835},
							expr: &ruleRefExpr{
								pos:  position{line: 142, col: 39, offset: 5836},
								name: "Section3",
							},
						},
						&labeledExpr{
							pos:   position{line: 142, col: 48, offset: 5845},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 142, col: 57, offset: 5854},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 142, col: 57, offset: 5854},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 142, col: 68, offset: 5865},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 142, col: 79, offset: 5876},
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
			pos:  position{line: 146, col: 1, offset: 5941},
			expr: &actionExpr{
				pos: position{line: 146, col: 13, offset: 5953},
				run: (*parser).callonSection41,
				expr: &seqExpr{
					pos: position{line: 146, col: 13, offset: 5953},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 146, col: 13, offset: 5953},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 146, col: 21, offset: 5961},
								name: "Section4Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 146, col: 36, offset: 5976},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 146, col: 46, offset: 5986},
								expr: &ruleRefExpr{
									pos:  position{line: 146, col: 46, offset: 5986},
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
			pos:  position{line: 150, col: 1, offset: 6094},
			expr: &actionExpr{
				pos: position{line: 150, col: 18, offset: 6111},
				run: (*parser).callonSection4Block1,
				expr: &seqExpr{
					pos: position{line: 150, col: 18, offset: 6111},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 150, col: 18, offset: 6111},
							expr: &ruleRefExpr{
								pos:  position{line: 150, col: 19, offset: 6112},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 150, col: 28, offset: 6121},
							expr: &ruleRefExpr{
								pos:  position{line: 150, col: 29, offset: 6122},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 150, col: 38, offset: 6131},
							expr: &ruleRefExpr{
								pos:  position{line: 150, col: 39, offset: 6132},
								name: "Section3",
							},
						},
						&notExpr{
							pos: position{line: 150, col: 48, offset: 6141},
							expr: &ruleRefExpr{
								pos:  position{line: 150, col: 49, offset: 6142},
								name: "Section4",
							},
						},
						&labeledExpr{
							pos:   position{line: 150, col: 58, offset: 6151},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 150, col: 67, offset: 6160},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 150, col: 67, offset: 6160},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 150, col: 78, offset: 6171},
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
			pos:  position{line: 154, col: 1, offset: 6236},
			expr: &actionExpr{
				pos: position{line: 154, col: 13, offset: 6248},
				run: (*parser).callonSection51,
				expr: &seqExpr{
					pos: position{line: 154, col: 13, offset: 6248},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 154, col: 13, offset: 6248},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 154, col: 21, offset: 6256},
								name: "Section5Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 154, col: 36, offset: 6271},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 154, col: 46, offset: 6281},
								expr: &ruleRefExpr{
									pos:  position{line: 154, col: 46, offset: 6281},
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
			pos:  position{line: 158, col: 1, offset: 6389},
			expr: &actionExpr{
				pos: position{line: 158, col: 18, offset: 6406},
				run: (*parser).callonSection5Block1,
				expr: &seqExpr{
					pos: position{line: 158, col: 18, offset: 6406},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 158, col: 18, offset: 6406},
							expr: &ruleRefExpr{
								pos:  position{line: 158, col: 19, offset: 6407},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 158, col: 28, offset: 6416},
							expr: &ruleRefExpr{
								pos:  position{line: 158, col: 29, offset: 6417},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 158, col: 38, offset: 6426},
							expr: &ruleRefExpr{
								pos:  position{line: 158, col: 39, offset: 6427},
								name: "Section3",
							},
						},
						&notExpr{
							pos: position{line: 158, col: 48, offset: 6436},
							expr: &ruleRefExpr{
								pos:  position{line: 158, col: 49, offset: 6437},
								name: "Section4",
							},
						},
						&notExpr{
							pos: position{line: 158, col: 58, offset: 6446},
							expr: &ruleRefExpr{
								pos:  position{line: 158, col: 59, offset: 6447},
								name: "Section5",
							},
						},
						&labeledExpr{
							pos:   position{line: 158, col: 68, offset: 6456},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 158, col: 77, offset: 6465},
								name: "StandaloneBlock",
							},
						},
					},
				},
			},
		},
		{
			name: "SectionTitle",
			pos:  position{line: 166, col: 1, offset: 6641},
			expr: &choiceExpr{
				pos: position{line: 166, col: 17, offset: 6657},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 166, col: 17, offset: 6657},
						name: "Section1Title",
					},
					&ruleRefExpr{
						pos:  position{line: 166, col: 33, offset: 6673},
						name: "Section2Title",
					},
					&ruleRefExpr{
						pos:  position{line: 166, col: 49, offset: 6689},
						name: "Section3Title",
					},
					&ruleRefExpr{
						pos:  position{line: 166, col: 65, offset: 6705},
						name: "Section4Title",
					},
					&ruleRefExpr{
						pos:  position{line: 166, col: 81, offset: 6721},
						name: "Section5Title",
					},
				},
			},
		},
		{
			name: "Section1Title",
			pos:  position{line: 168, col: 1, offset: 6736},
			expr: &actionExpr{
				pos: position{line: 168, col: 18, offset: 6753},
				run: (*parser).callonSection1Title1,
				expr: &seqExpr{
					pos: position{line: 168, col: 18, offset: 6753},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 168, col: 18, offset: 6753},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 168, col: 29, offset: 6764},
								expr: &ruleRefExpr{
									pos:  position{line: 168, col: 30, offset: 6765},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 168, col: 49, offset: 6784},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 168, col: 56, offset: 6791},
								val:        "==",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 168, col: 62, offset: 6797},
							expr: &ruleRefExpr{
								pos:  position{line: 168, col: 62, offset: 6797},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 168, col: 66, offset: 6801},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 168, col: 74, offset: 6809},
								name: "InlineContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 168, col: 88, offset: 6823},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 168, col: 93, offset: 6828},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 168, col: 93, offset: 6828},
									expr: &ruleRefExpr{
										pos:  position{line: 168, col: 93, offset: 6828},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 168, col: 106, offset: 6841},
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
			pos:  position{line: 172, col: 1, offset: 6946},
			expr: &actionExpr{
				pos: position{line: 172, col: 18, offset: 6963},
				run: (*parser).callonSection2Title1,
				expr: &seqExpr{
					pos: position{line: 172, col: 18, offset: 6963},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 172, col: 18, offset: 6963},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 172, col: 29, offset: 6974},
								expr: &ruleRefExpr{
									pos:  position{line: 172, col: 30, offset: 6975},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 172, col: 49, offset: 6994},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 172, col: 56, offset: 7001},
								val:        "===",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 172, col: 63, offset: 7008},
							expr: &ruleRefExpr{
								pos:  position{line: 172, col: 63, offset: 7008},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 172, col: 67, offset: 7012},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 172, col: 75, offset: 7020},
								name: "InlineContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 172, col: 89, offset: 7034},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 172, col: 94, offset: 7039},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 172, col: 94, offset: 7039},
									expr: &ruleRefExpr{
										pos:  position{line: 172, col: 94, offset: 7039},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 172, col: 107, offset: 7052},
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
			pos:  position{line: 176, col: 1, offset: 7156},
			expr: &actionExpr{
				pos: position{line: 176, col: 18, offset: 7173},
				run: (*parser).callonSection3Title1,
				expr: &seqExpr{
					pos: position{line: 176, col: 18, offset: 7173},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 176, col: 18, offset: 7173},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 176, col: 29, offset: 7184},
								expr: &ruleRefExpr{
									pos:  position{line: 176, col: 30, offset: 7185},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 176, col: 49, offset: 7204},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 176, col: 56, offset: 7211},
								val:        "====",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 176, col: 64, offset: 7219},
							expr: &ruleRefExpr{
								pos:  position{line: 176, col: 64, offset: 7219},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 176, col: 68, offset: 7223},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 176, col: 76, offset: 7231},
								name: "InlineContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 176, col: 90, offset: 7245},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 176, col: 95, offset: 7250},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 176, col: 95, offset: 7250},
									expr: &ruleRefExpr{
										pos:  position{line: 176, col: 95, offset: 7250},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 176, col: 108, offset: 7263},
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
			pos:  position{line: 180, col: 1, offset: 7367},
			expr: &actionExpr{
				pos: position{line: 180, col: 18, offset: 7384},
				run: (*parser).callonSection4Title1,
				expr: &seqExpr{
					pos: position{line: 180, col: 18, offset: 7384},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 180, col: 18, offset: 7384},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 180, col: 29, offset: 7395},
								expr: &ruleRefExpr{
									pos:  position{line: 180, col: 30, offset: 7396},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 180, col: 49, offset: 7415},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 180, col: 56, offset: 7422},
								val:        "=====",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 180, col: 65, offset: 7431},
							expr: &ruleRefExpr{
								pos:  position{line: 180, col: 65, offset: 7431},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 180, col: 69, offset: 7435},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 180, col: 77, offset: 7443},
								name: "InlineContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 180, col: 91, offset: 7457},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 180, col: 96, offset: 7462},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 180, col: 96, offset: 7462},
									expr: &ruleRefExpr{
										pos:  position{line: 180, col: 96, offset: 7462},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 180, col: 109, offset: 7475},
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
			pos:  position{line: 184, col: 1, offset: 7579},
			expr: &actionExpr{
				pos: position{line: 184, col: 18, offset: 7596},
				run: (*parser).callonSection5Title1,
				expr: &seqExpr{
					pos: position{line: 184, col: 18, offset: 7596},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 184, col: 18, offset: 7596},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 184, col: 29, offset: 7607},
								expr: &ruleRefExpr{
									pos:  position{line: 184, col: 30, offset: 7608},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 184, col: 49, offset: 7627},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 184, col: 56, offset: 7634},
								val:        "======",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 184, col: 66, offset: 7644},
							expr: &ruleRefExpr{
								pos:  position{line: 184, col: 66, offset: 7644},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 184, col: 70, offset: 7648},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 184, col: 78, offset: 7656},
								name: "InlineContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 184, col: 92, offset: 7670},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 184, col: 97, offset: 7675},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 184, col: 97, offset: 7675},
									expr: &ruleRefExpr{
										pos:  position{line: 184, col: 97, offset: 7675},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 184, col: 110, offset: 7688},
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
			pos:  position{line: 191, col: 1, offset: 7898},
			expr: &actionExpr{
				pos: position{line: 191, col: 9, offset: 7906},
				run: (*parser).callonList1,
				expr: &seqExpr{
					pos: position{line: 191, col: 9, offset: 7906},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 191, col: 9, offset: 7906},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 191, col: 20, offset: 7917},
								expr: &ruleRefExpr{
									pos:  position{line: 191, col: 21, offset: 7918},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 193, col: 5, offset: 8010},
							label: "elements",
							expr: &oneOrMoreExpr{
								pos: position{line: 193, col: 14, offset: 8019},
								expr: &seqExpr{
									pos: position{line: 193, col: 15, offset: 8020},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 193, col: 15, offset: 8020},
											name: "ListItem",
										},
										&zeroOrOneExpr{
											pos: position{line: 193, col: 24, offset: 8029},
											expr: &ruleRefExpr{
												pos:  position{line: 193, col: 24, offset: 8029},
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
			pos:  position{line: 197, col: 1, offset: 8126},
			expr: &actionExpr{
				pos: position{line: 197, col: 13, offset: 8138},
				run: (*parser).callonListItem1,
				expr: &seqExpr{
					pos: position{line: 197, col: 13, offset: 8138},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 197, col: 13, offset: 8138},
							expr: &ruleRefExpr{
								pos:  position{line: 197, col: 13, offset: 8138},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 197, col: 17, offset: 8142},
							label: "level",
							expr: &choiceExpr{
								pos: position{line: 197, col: 24, offset: 8149},
								alternatives: []interface{}{
									&oneOrMoreExpr{
										pos: position{line: 197, col: 24, offset: 8149},
										expr: &litMatcher{
											pos:        position{line: 197, col: 24, offset: 8149},
											val:        "*",
											ignoreCase: false,
										},
									},
									&litMatcher{
										pos:        position{line: 197, col: 31, offset: 8156},
										val:        "-",
										ignoreCase: false,
									},
								},
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 197, col: 36, offset: 8161},
							expr: &ruleRefExpr{
								pos:  position{line: 197, col: 36, offset: 8161},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 197, col: 40, offset: 8165},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 197, col: 49, offset: 8174},
								name: "ListItemContent",
							},
						},
					},
				},
			},
		},
		{
			name: "ListItemContent",
			pos:  position{line: 201, col: 1, offset: 8271},
			expr: &actionExpr{
				pos: position{line: 201, col: 20, offset: 8290},
				run: (*parser).callonListItemContent1,
				expr: &labeledExpr{
					pos:   position{line: 201, col: 20, offset: 8290},
					label: "lines",
					expr: &oneOrMoreExpr{
						pos: position{line: 201, col: 26, offset: 8296},
						expr: &seqExpr{
							pos: position{line: 201, col: 27, offset: 8297},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 201, col: 27, offset: 8297},
									expr: &seqExpr{
										pos: position{line: 201, col: 29, offset: 8299},
										exprs: []interface{}{
											&zeroOrMoreExpr{
												pos: position{line: 201, col: 29, offset: 8299},
												expr: &ruleRefExpr{
													pos:  position{line: 201, col: 29, offset: 8299},
													name: "WS",
												},
											},
											&choiceExpr{
												pos: position{line: 201, col: 34, offset: 8304},
												alternatives: []interface{}{
													&oneOrMoreExpr{
														pos: position{line: 201, col: 34, offset: 8304},
														expr: &litMatcher{
															pos:        position{line: 201, col: 34, offset: 8304},
															val:        "*",
															ignoreCase: false,
														},
													},
													&litMatcher{
														pos:        position{line: 201, col: 41, offset: 8311},
														val:        "-",
														ignoreCase: false,
													},
												},
											},
											&oneOrMoreExpr{
												pos: position{line: 201, col: 46, offset: 8316},
												expr: &ruleRefExpr{
													pos:  position{line: 201, col: 46, offset: 8316},
													name: "WS",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 201, col: 51, offset: 8321},
									name: "InlineContent",
								},
								&ruleRefExpr{
									pos:  position{line: 201, col: 65, offset: 8335},
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
			pos:  position{line: 209, col: 1, offset: 8672},
			expr: &actionExpr{
				pos: position{line: 209, col: 14, offset: 8685},
				run: (*parser).callonParagraph1,
				expr: &seqExpr{
					pos: position{line: 209, col: 14, offset: 8685},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 209, col: 14, offset: 8685},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 209, col: 25, offset: 8696},
								expr: &ruleRefExpr{
									pos:  position{line: 209, col: 26, offset: 8697},
									name: "ElementAttribute",
								},
							},
						},
						&notExpr{
							pos: position{line: 209, col: 45, offset: 8716},
							expr: &seqExpr{
								pos: position{line: 209, col: 47, offset: 8718},
								exprs: []interface{}{
									&oneOrMoreExpr{
										pos: position{line: 209, col: 47, offset: 8718},
										expr: &litMatcher{
											pos:        position{line: 209, col: 47, offset: 8718},
											val:        "=",
											ignoreCase: false,
										},
									},
									&oneOrMoreExpr{
										pos: position{line: 209, col: 52, offset: 8723},
										expr: &ruleRefExpr{
											pos:  position{line: 209, col: 52, offset: 8723},
											name: "WS",
										},
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 209, col: 57, offset: 8728},
							label: "lines",
							expr: &oneOrMoreExpr{
								pos: position{line: 209, col: 63, offset: 8734},
								expr: &seqExpr{
									pos: position{line: 209, col: 64, offset: 8735},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 209, col: 64, offset: 8735},
											name: "InlineContent",
										},
										&ruleRefExpr{
											pos:  position{line: 209, col: 78, offset: 8749},
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
			pos:  position{line: 215, col: 1, offset: 9047},
			expr: &actionExpr{
				pos: position{line: 215, col: 18, offset: 9064},
				run: (*parser).callonInlineContent1,
				expr: &seqExpr{
					pos: position{line: 215, col: 18, offset: 9064},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 215, col: 18, offset: 9064},
							expr: &ruleRefExpr{
								pos:  position{line: 215, col: 19, offset: 9065},
								name: "FencedBlockDelimiter",
							},
						},
						&labeledExpr{
							pos:   position{line: 215, col: 40, offset: 9086},
							label: "elements",
							expr: &oneOrMoreExpr{
								pos: position{line: 215, col: 49, offset: 9095},
								expr: &seqExpr{
									pos: position{line: 215, col: 50, offset: 9096},
									exprs: []interface{}{
										&zeroOrMoreExpr{
											pos: position{line: 215, col: 50, offset: 9096},
											expr: &ruleRefExpr{
												pos:  position{line: 215, col: 50, offset: 9096},
												name: "WS",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 215, col: 54, offset: 9100},
											name: "InlineElement",
										},
										&zeroOrMoreExpr{
											pos: position{line: 215, col: 68, offset: 9114},
											expr: &ruleRefExpr{
												pos:  position{line: 215, col: 68, offset: 9114},
												name: "WS",
											},
										},
									},
								},
							},
						},
						&andExpr{
							pos: position{line: 215, col: 74, offset: 9120},
							expr: &ruleRefExpr{
								pos:  position{line: 215, col: 75, offset: 9121},
								name: "EOL",
							},
						},
					},
				},
			},
		},
		{
			name: "InlineElement",
			pos:  position{line: 219, col: 1, offset: 9247},
			expr: &choiceExpr{
				pos: position{line: 219, col: 18, offset: 9264},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 219, col: 18, offset: 9264},
						name: "InlineImage",
					},
					&ruleRefExpr{
						pos:  position{line: 219, col: 32, offset: 9278},
						name: "QuotedText",
					},
					&ruleRefExpr{
						pos:  position{line: 219, col: 45, offset: 9291},
						name: "ExternalLink",
					},
					&ruleRefExpr{
						pos:  position{line: 219, col: 60, offset: 9306},
						name: "DocumentAttributeSubstitution",
					},
					&ruleRefExpr{
						pos:  position{line: 219, col: 92, offset: 9338},
						name: "Word",
					},
				},
			},
		},
		{
			name: "QuotedText",
			pos:  position{line: 224, col: 1, offset: 9481},
			expr: &choiceExpr{
				pos: position{line: 224, col: 15, offset: 9495},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 224, col: 15, offset: 9495},
						name: "BoldText",
					},
					&ruleRefExpr{
						pos:  position{line: 224, col: 26, offset: 9506},
						name: "ItalicText",
					},
					&ruleRefExpr{
						pos:  position{line: 224, col: 39, offset: 9519},
						name: "MonospaceText",
					},
				},
			},
		},
		{
			name: "BoldText",
			pos:  position{line: 226, col: 1, offset: 9534},
			expr: &choiceExpr{
				pos: position{line: 226, col: 13, offset: 9546},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 226, col: 13, offset: 9546},
						name: "BoldTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 226, col: 41, offset: 9574},
						name: "BoldTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 226, col: 73, offset: 9606},
						name: "BoldTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "BoldTextSimplePunctuation",
			pos:  position{line: 228, col: 1, offset: 9679},
			expr: &actionExpr{
				pos: position{line: 228, col: 30, offset: 9708},
				run: (*parser).callonBoldTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 228, col: 30, offset: 9708},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 228, col: 30, offset: 9708},
							val:        "*",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 228, col: 34, offset: 9712},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 228, col: 43, offset: 9721},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 228, col: 62, offset: 9740},
							val:        "*",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "BoldTextDoublePunctuation",
			pos:  position{line: 232, col: 1, offset: 9817},
			expr: &actionExpr{
				pos: position{line: 232, col: 30, offset: 9846},
				run: (*parser).callonBoldTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 232, col: 30, offset: 9846},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 232, col: 30, offset: 9846},
							val:        "**",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 232, col: 35, offset: 9851},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 232, col: 44, offset: 9860},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 232, col: 63, offset: 9879},
							val:        "**",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "BoldTextUnbalancedPunctuation",
			pos:  position{line: 236, col: 1, offset: 9957},
			expr: &actionExpr{
				pos: position{line: 236, col: 34, offset: 9990},
				run: (*parser).callonBoldTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 236, col: 34, offset: 9990},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 236, col: 34, offset: 9990},
							val:        "**",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 236, col: 39, offset: 9995},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 236, col: 48, offset: 10004},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 236, col: 67, offset: 10023},
							val:        "*",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ItalicText",
			pos:  position{line: 241, col: 1, offset: 10187},
			expr: &choiceExpr{
				pos: position{line: 241, col: 15, offset: 10201},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 241, col: 15, offset: 10201},
						name: "ItalicTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 241, col: 45, offset: 10231},
						name: "ItalicTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 241, col: 79, offset: 10265},
						name: "ItalicTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "ItalicTextSimplePunctuation",
			pos:  position{line: 243, col: 1, offset: 10294},
			expr: &actionExpr{
				pos: position{line: 243, col: 32, offset: 10325},
				run: (*parser).callonItalicTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 243, col: 32, offset: 10325},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 243, col: 32, offset: 10325},
							val:        "_",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 243, col: 36, offset: 10329},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 243, col: 45, offset: 10338},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 243, col: 64, offset: 10357},
							val:        "_",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ItalicTextDoublePunctuation",
			pos:  position{line: 247, col: 1, offset: 10436},
			expr: &actionExpr{
				pos: position{line: 247, col: 32, offset: 10467},
				run: (*parser).callonItalicTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 247, col: 32, offset: 10467},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 247, col: 32, offset: 10467},
							val:        "__",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 247, col: 37, offset: 10472},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 247, col: 46, offset: 10481},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 247, col: 65, offset: 10500},
							val:        "__",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ItalicTextUnbalancedPunctuation",
			pos:  position{line: 251, col: 1, offset: 10580},
			expr: &actionExpr{
				pos: position{line: 251, col: 36, offset: 10615},
				run: (*parser).callonItalicTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 251, col: 36, offset: 10615},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 251, col: 36, offset: 10615},
							val:        "__",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 251, col: 41, offset: 10620},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 251, col: 50, offset: 10629},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 251, col: 69, offset: 10648},
							val:        "_",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "MonospaceText",
			pos:  position{line: 256, col: 1, offset: 10814},
			expr: &choiceExpr{
				pos: position{line: 256, col: 18, offset: 10831},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 256, col: 18, offset: 10831},
						name: "MonospaceTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 256, col: 51, offset: 10864},
						name: "MonospaceTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 256, col: 88, offset: 10901},
						name: "MonospaceTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "MonospaceTextSimplePunctuation",
			pos:  position{line: 258, col: 1, offset: 10933},
			expr: &actionExpr{
				pos: position{line: 258, col: 35, offset: 10967},
				run: (*parser).callonMonospaceTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 258, col: 35, offset: 10967},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 258, col: 35, offset: 10967},
							val:        "`",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 258, col: 39, offset: 10971},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 258, col: 48, offset: 10980},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 258, col: 67, offset: 10999},
							val:        "`",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "MonospaceTextDoublePunctuation",
			pos:  position{line: 262, col: 1, offset: 11081},
			expr: &actionExpr{
				pos: position{line: 262, col: 35, offset: 11115},
				run: (*parser).callonMonospaceTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 262, col: 35, offset: 11115},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 262, col: 35, offset: 11115},
							val:        "``",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 262, col: 40, offset: 11120},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 262, col: 49, offset: 11129},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 262, col: 68, offset: 11148},
							val:        "``",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "MonospaceTextUnbalancedPunctuation",
			pos:  position{line: 266, col: 1, offset: 11231},
			expr: &actionExpr{
				pos: position{line: 266, col: 39, offset: 11269},
				run: (*parser).callonMonospaceTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 266, col: 39, offset: 11269},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 266, col: 39, offset: 11269},
							val:        "``",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 266, col: 44, offset: 11274},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 266, col: 53, offset: 11283},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 266, col: 72, offset: 11302},
							val:        "`",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "QuotedTextContent",
			pos:  position{line: 271, col: 1, offset: 11471},
			expr: &seqExpr{
				pos: position{line: 271, col: 22, offset: 11492},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 271, col: 22, offset: 11492},
						name: "QuotedTextContentElement",
					},
					&zeroOrMoreExpr{
						pos: position{line: 271, col: 47, offset: 11517},
						expr: &seqExpr{
							pos: position{line: 271, col: 48, offset: 11518},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 271, col: 48, offset: 11518},
									expr: &ruleRefExpr{
										pos:  position{line: 271, col: 48, offset: 11518},
										name: "WS",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 271, col: 52, offset: 11522},
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
			pos:  position{line: 273, col: 1, offset: 11550},
			expr: &choiceExpr{
				pos: position{line: 273, col: 29, offset: 11578},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 273, col: 29, offset: 11578},
						name: "QuotedText",
					},
					&ruleRefExpr{
						pos:  position{line: 273, col: 42, offset: 11591},
						name: "QuotedTextWord",
					},
					&ruleRefExpr{
						pos:  position{line: 273, col: 59, offset: 11608},
						name: "WordWithQuotePunctuation",
					},
				},
			},
		},
		{
			name: "QuotedTextWord",
			pos:  position{line: 275, col: 1, offset: 11737},
			expr: &oneOrMoreExpr{
				pos: position{line: 275, col: 19, offset: 11755},
				expr: &seqExpr{
					pos: position{line: 275, col: 20, offset: 11756},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 275, col: 20, offset: 11756},
							expr: &ruleRefExpr{
								pos:  position{line: 275, col: 21, offset: 11757},
								name: "NEWLINE",
							},
						},
						&notExpr{
							pos: position{line: 275, col: 29, offset: 11765},
							expr: &ruleRefExpr{
								pos:  position{line: 275, col: 30, offset: 11766},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 275, col: 33, offset: 11769},
							expr: &litMatcher{
								pos:        position{line: 275, col: 34, offset: 11770},
								val:        "*",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 275, col: 38, offset: 11774},
							expr: &litMatcher{
								pos:        position{line: 275, col: 39, offset: 11775},
								val:        "_",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 275, col: 43, offset: 11779},
							expr: &litMatcher{
								pos:        position{line: 275, col: 44, offset: 11780},
								val:        "`",
								ignoreCase: false,
							},
						},
						&anyMatcher{
							line: 275, col: 48, offset: 11784,
						},
					},
				},
			},
		},
		{
			name: "WordWithQuotePunctuation",
			pos:  position{line: 276, col: 1, offset: 11826},
			expr: &actionExpr{
				pos: position{line: 276, col: 29, offset: 11854},
				run: (*parser).callonWordWithQuotePunctuation1,
				expr: &oneOrMoreExpr{
					pos: position{line: 276, col: 29, offset: 11854},
					expr: &seqExpr{
						pos: position{line: 276, col: 30, offset: 11855},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 276, col: 30, offset: 11855},
								expr: &ruleRefExpr{
									pos:  position{line: 276, col: 31, offset: 11856},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 276, col: 39, offset: 11864},
								expr: &ruleRefExpr{
									pos:  position{line: 276, col: 40, offset: 11865},
									name: "WS",
								},
							},
							&anyMatcher{
								line: 276, col: 44, offset: 11869,
							},
						},
					},
				},
			},
		},
		{
			name: "UnbalancedQuotePunctuation",
			pos:  position{line: 281, col: 1, offset: 12114},
			expr: &choiceExpr{
				pos: position{line: 281, col: 31, offset: 12144},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 281, col: 31, offset: 12144},
						val:        "*",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 281, col: 37, offset: 12150},
						val:        "_",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 281, col: 43, offset: 12156},
						val:        "`",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "ExternalLink",
			pos:  position{line: 287, col: 1, offset: 12263},
			expr: &actionExpr{
				pos: position{line: 287, col: 17, offset: 12279},
				run: (*parser).callonExternalLink1,
				expr: &seqExpr{
					pos: position{line: 287, col: 17, offset: 12279},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 287, col: 17, offset: 12279},
							label: "url",
							expr: &seqExpr{
								pos: position{line: 287, col: 22, offset: 12284},
								exprs: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 287, col: 22, offset: 12284},
										name: "URL_SCHEME",
									},
									&ruleRefExpr{
										pos:  position{line: 287, col: 33, offset: 12295},
										name: "URL",
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 287, col: 38, offset: 12300},
							label: "text",
							expr: &zeroOrOneExpr{
								pos: position{line: 287, col: 43, offset: 12305},
								expr: &seqExpr{
									pos: position{line: 287, col: 44, offset: 12306},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 287, col: 44, offset: 12306},
											val:        "[",
											ignoreCase: false,
										},
										&zeroOrMoreExpr{
											pos: position{line: 287, col: 48, offset: 12310},
											expr: &ruleRefExpr{
												pos:  position{line: 287, col: 49, offset: 12311},
												name: "URL_TEXT",
											},
										},
										&litMatcher{
											pos:        position{line: 287, col: 60, offset: 12322},
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
			pos:  position{line: 297, col: 1, offset: 12601},
			expr: &actionExpr{
				pos: position{line: 297, col: 15, offset: 12615},
				run: (*parser).callonBlockImage1,
				expr: &seqExpr{
					pos: position{line: 297, col: 15, offset: 12615},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 297, col: 15, offset: 12615},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 297, col: 26, offset: 12626},
								expr: &ruleRefExpr{
									pos:  position{line: 297, col: 27, offset: 12627},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 297, col: 46, offset: 12646},
							label: "image",
							expr: &ruleRefExpr{
								pos:  position{line: 297, col: 52, offset: 12652},
								name: "BlockImageMacro",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 297, col: 69, offset: 12669},
							expr: &ruleRefExpr{
								pos:  position{line: 297, col: 69, offset: 12669},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 297, col: 73, offset: 12673},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "BlockImageMacro",
			pos:  position{line: 302, col: 1, offset: 12842},
			expr: &actionExpr{
				pos: position{line: 302, col: 20, offset: 12861},
				run: (*parser).callonBlockImageMacro1,
				expr: &seqExpr{
					pos: position{line: 302, col: 20, offset: 12861},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 302, col: 20, offset: 12861},
							val:        "image::",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 302, col: 30, offset: 12871},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 302, col: 36, offset: 12877},
								name: "URL",
							},
						},
						&litMatcher{
							pos:        position{line: 302, col: 41, offset: 12882},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 302, col: 45, offset: 12886},
							label: "attributes",
							expr: &zeroOrOneExpr{
								pos: position{line: 302, col: 57, offset: 12898},
								expr: &ruleRefExpr{
									pos:  position{line: 302, col: 57, offset: 12898},
									name: "URL_TEXT",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 302, col: 68, offset: 12909},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "InlineImage",
			pos:  position{line: 306, col: 1, offset: 12984},
			expr: &actionExpr{
				pos: position{line: 306, col: 16, offset: 12999},
				run: (*parser).callonInlineImage1,
				expr: &labeledExpr{
					pos:   position{line: 306, col: 16, offset: 12999},
					label: "image",
					expr: &ruleRefExpr{
						pos:  position{line: 306, col: 22, offset: 13005},
						name: "InlineImageMacro",
					},
				},
			},
		},
		{
			name: "InlineImageMacro",
			pos:  position{line: 311, col: 1, offset: 13160},
			expr: &actionExpr{
				pos: position{line: 311, col: 21, offset: 13180},
				run: (*parser).callonInlineImageMacro1,
				expr: &seqExpr{
					pos: position{line: 311, col: 21, offset: 13180},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 311, col: 21, offset: 13180},
							val:        "image:",
							ignoreCase: false,
						},
						&notExpr{
							pos: position{line: 311, col: 30, offset: 13189},
							expr: &litMatcher{
								pos:        position{line: 311, col: 31, offset: 13190},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 311, col: 35, offset: 13194},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 311, col: 41, offset: 13200},
								name: "URL",
							},
						},
						&litMatcher{
							pos:        position{line: 311, col: 46, offset: 13205},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 311, col: 50, offset: 13209},
							label: "attributes",
							expr: &zeroOrOneExpr{
								pos: position{line: 311, col: 62, offset: 13221},
								expr: &ruleRefExpr{
									pos:  position{line: 311, col: 62, offset: 13221},
									name: "URL_TEXT",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 311, col: 73, offset: 13232},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "DelimitedBlock",
			pos:  position{line: 318, col: 1, offset: 13570},
			expr: &ruleRefExpr{
				pos:  position{line: 318, col: 19, offset: 13588},
				name: "FencedBlock",
			},
		},
		{
			name: "FencedBlock",
			pos:  position{line: 320, col: 1, offset: 13602},
			expr: &actionExpr{
				pos: position{line: 320, col: 16, offset: 13617},
				run: (*parser).callonFencedBlock1,
				expr: &seqExpr{
					pos: position{line: 320, col: 16, offset: 13617},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 320, col: 16, offset: 13617},
							name: "FencedBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 320, col: 37, offset: 13638},
							expr: &ruleRefExpr{
								pos:  position{line: 320, col: 37, offset: 13638},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 320, col: 41, offset: 13642},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 320, col: 49, offset: 13650},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 320, col: 58, offset: 13659},
								name: "FencedBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 320, col: 78, offset: 13679},
							name: "FencedBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 320, col: 99, offset: 13700},
							expr: &ruleRefExpr{
								pos:  position{line: 320, col: 99, offset: 13700},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 320, col: 103, offset: 13704},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "FencedBlockDelimiter",
			pos:  position{line: 324, col: 1, offset: 13792},
			expr: &litMatcher{
				pos:        position{line: 324, col: 25, offset: 13816},
				val:        "```",
				ignoreCase: false,
			},
		},
		{
			name: "FencedBlockContent",
			pos:  position{line: 326, col: 1, offset: 13823},
			expr: &labeledExpr{
				pos:   position{line: 326, col: 23, offset: 13845},
				label: "content",
				expr: &zeroOrMoreExpr{
					pos: position{line: 326, col: 31, offset: 13853},
					expr: &seqExpr{
						pos: position{line: 326, col: 32, offset: 13854},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 326, col: 32, offset: 13854},
								expr: &ruleRefExpr{
									pos:  position{line: 326, col: 33, offset: 13855},
									name: "FencedBlockDelimiter",
								},
							},
							&anyMatcher{
								line: 326, col: 54, offset: 13876,
							},
						},
					},
				},
			},
		},
		{
			name: "LiteralBlock",
			pos:  position{line: 331, col: 1, offset: 14149},
			expr: &choiceExpr{
				pos: position{line: 331, col: 17, offset: 14165},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 331, col: 17, offset: 14165},
						name: "ParagraphWithSpaces",
					},
					&ruleRefExpr{
						pos:  position{line: 331, col: 39, offset: 14187},
						name: "ParagraphWithLiteralBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 331, col: 76, offset: 14224},
						name: "ParagraphWithLiteralAttribute",
					},
				},
			},
		},
		{
			name: "ParagraphWithSpaces",
			pos:  position{line: 334, col: 1, offset: 14319},
			expr: &actionExpr{
				pos: position{line: 334, col: 24, offset: 14342},
				run: (*parser).callonParagraphWithSpaces1,
				expr: &seqExpr{
					pos: position{line: 334, col: 24, offset: 14342},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 334, col: 24, offset: 14342},
							label: "spaces",
							expr: &oneOrMoreExpr{
								pos: position{line: 334, col: 32, offset: 14350},
								expr: &ruleRefExpr{
									pos:  position{line: 334, col: 32, offset: 14350},
									name: "WS",
								},
							},
						},
						&notExpr{
							pos: position{line: 334, col: 37, offset: 14355},
							expr: &ruleRefExpr{
								pos:  position{line: 334, col: 38, offset: 14356},
								name: "NEWLINE",
							},
						},
						&labeledExpr{
							pos:   position{line: 334, col: 46, offset: 14364},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 334, col: 55, offset: 14373},
								name: "LiteralBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 334, col: 76, offset: 14394},
							name: "EndOfLiteralBlock",
						},
					},
				},
			},
		},
		{
			name: "LiteralBlockContent",
			pos:  position{line: 339, col: 1, offset: 14575},
			expr: &actionExpr{
				pos: position{line: 339, col: 24, offset: 14598},
				run: (*parser).callonLiteralBlockContent1,
				expr: &labeledExpr{
					pos:   position{line: 339, col: 24, offset: 14598},
					label: "content",
					expr: &oneOrMoreExpr{
						pos: position{line: 339, col: 32, offset: 14606},
						expr: &seqExpr{
							pos: position{line: 339, col: 33, offset: 14607},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 339, col: 33, offset: 14607},
									expr: &seqExpr{
										pos: position{line: 339, col: 35, offset: 14609},
										exprs: []interface{}{
											&ruleRefExpr{
												pos:  position{line: 339, col: 35, offset: 14609},
												name: "NEWLINE",
											},
											&ruleRefExpr{
												pos:  position{line: 339, col: 43, offset: 14617},
												name: "BlankLine",
											},
										},
									},
								},
								&anyMatcher{
									line: 339, col: 54, offset: 14628,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "EndOfLiteralBlock",
			pos:  position{line: 344, col: 1, offset: 14713},
			expr: &choiceExpr{
				pos: position{line: 344, col: 22, offset: 14734},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 344, col: 22, offset: 14734},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 344, col: 22, offset: 14734},
								name: "NEWLINE",
							},
							&ruleRefExpr{
								pos:  position{line: 344, col: 30, offset: 14742},
								name: "BlankLine",
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 344, col: 42, offset: 14754},
						name: "NEWLINE",
					},
					&ruleRefExpr{
						pos:  position{line: 344, col: 52, offset: 14764},
						name: "EOF",
					},
				},
			},
		},
		{
			name: "ParagraphWithLiteralBlockDelimiter",
			pos:  position{line: 347, col: 1, offset: 14824},
			expr: &actionExpr{
				pos: position{line: 347, col: 39, offset: 14862},
				run: (*parser).callonParagraphWithLiteralBlockDelimiter1,
				expr: &seqExpr{
					pos: position{line: 347, col: 39, offset: 14862},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 347, col: 39, offset: 14862},
							name: "LiteralBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 347, col: 61, offset: 14884},
							expr: &ruleRefExpr{
								pos:  position{line: 347, col: 61, offset: 14884},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 347, col: 65, offset: 14888},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 347, col: 73, offset: 14896},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 347, col: 81, offset: 14904},
								expr: &seqExpr{
									pos: position{line: 347, col: 82, offset: 14905},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 347, col: 82, offset: 14905},
											expr: &ruleRefExpr{
												pos:  position{line: 347, col: 83, offset: 14906},
												name: "LiteralBlockDelimiter",
											},
										},
										&anyMatcher{
											line: 347, col: 105, offset: 14928,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 347, col: 109, offset: 14932},
							name: "LiteralBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 347, col: 131, offset: 14954},
							expr: &ruleRefExpr{
								pos:  position{line: 347, col: 131, offset: 14954},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 347, col: 135, offset: 14958},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "LiteralBlockDelimiter",
			pos:  position{line: 351, col: 1, offset: 15042},
			expr: &litMatcher{
				pos:        position{line: 351, col: 26, offset: 15067},
				val:        "....",
				ignoreCase: false,
			},
		},
		{
			name: "ParagraphWithLiteralAttribute",
			pos:  position{line: 354, col: 1, offset: 15129},
			expr: &actionExpr{
				pos: position{line: 354, col: 34, offset: 15162},
				run: (*parser).callonParagraphWithLiteralAttribute1,
				expr: &seqExpr{
					pos: position{line: 354, col: 34, offset: 15162},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 354, col: 34, offset: 15162},
							val:        "[literal]",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 354, col: 46, offset: 15174},
							expr: &ruleRefExpr{
								pos:  position{line: 354, col: 46, offset: 15174},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 354, col: 50, offset: 15178},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 354, col: 58, offset: 15186},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 354, col: 67, offset: 15195},
								name: "LiteralBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 354, col: 88, offset: 15216},
							name: "EndOfLiteralBlock",
						},
					},
				},
			},
		},
		{
			name: "ElementAttribute",
			pos:  position{line: 361, col: 1, offset: 15428},
			expr: &labeledExpr{
				pos:   position{line: 361, col: 21, offset: 15448},
				label: "meta",
				expr: &choiceExpr{
					pos: position{line: 361, col: 27, offset: 15454},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 361, col: 27, offset: 15454},
							name: "ElementLink",
						},
						&ruleRefExpr{
							pos:  position{line: 361, col: 41, offset: 15468},
							name: "ElementID",
						},
						&ruleRefExpr{
							pos:  position{line: 361, col: 53, offset: 15480},
							name: "ElementTitle",
						},
					},
				},
			},
		},
		{
			name: "ElementLink",
			pos:  position{line: 364, col: 1, offset: 15551},
			expr: &actionExpr{
				pos: position{line: 364, col: 16, offset: 15566},
				run: (*parser).callonElementLink1,
				expr: &seqExpr{
					pos: position{line: 364, col: 16, offset: 15566},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 364, col: 16, offset: 15566},
							val:        "[",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 364, col: 20, offset: 15570},
							expr: &ruleRefExpr{
								pos:  position{line: 364, col: 20, offset: 15570},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 364, col: 24, offset: 15574},
							val:        "link",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 364, col: 31, offset: 15581},
							expr: &ruleRefExpr{
								pos:  position{line: 364, col: 31, offset: 15581},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 364, col: 35, offset: 15585},
							val:        "=",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 364, col: 39, offset: 15589},
							expr: &ruleRefExpr{
								pos:  position{line: 364, col: 39, offset: 15589},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 364, col: 43, offset: 15593},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 364, col: 48, offset: 15598},
								name: "URL",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 364, col: 52, offset: 15602},
							expr: &ruleRefExpr{
								pos:  position{line: 364, col: 52, offset: 15602},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 364, col: 56, offset: 15606},
							val:        "]",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 364, col: 60, offset: 15610},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ElementID",
			pos:  position{line: 369, col: 1, offset: 15720},
			expr: &actionExpr{
				pos: position{line: 369, col: 14, offset: 15733},
				run: (*parser).callonElementID1,
				expr: &seqExpr{
					pos: position{line: 369, col: 14, offset: 15733},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 369, col: 14, offset: 15733},
							val:        "[",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 369, col: 18, offset: 15737},
							expr: &ruleRefExpr{
								pos:  position{line: 369, col: 18, offset: 15737},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 369, col: 22, offset: 15741},
							val:        "#",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 369, col: 26, offset: 15745},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 369, col: 30, offset: 15749},
								name: "ID",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 369, col: 34, offset: 15753},
							expr: &ruleRefExpr{
								pos:  position{line: 369, col: 34, offset: 15753},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 369, col: 38, offset: 15757},
							val:        "]",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 369, col: 42, offset: 15761},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ElementTitle",
			pos:  position{line: 375, col: 1, offset: 15955},
			expr: &actionExpr{
				pos: position{line: 375, col: 17, offset: 15971},
				run: (*parser).callonElementTitle1,
				expr: &seqExpr{
					pos: position{line: 375, col: 17, offset: 15971},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 375, col: 17, offset: 15971},
							val:        ".",
							ignoreCase: false,
						},
						&notExpr{
							pos: position{line: 375, col: 21, offset: 15975},
							expr: &litMatcher{
								pos:        position{line: 375, col: 22, offset: 15976},
								val:        ".",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 375, col: 26, offset: 15980},
							expr: &ruleRefExpr{
								pos:  position{line: 375, col: 27, offset: 15981},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 375, col: 30, offset: 15984},
							label: "title",
							expr: &oneOrMoreExpr{
								pos: position{line: 375, col: 36, offset: 15990},
								expr: &seqExpr{
									pos: position{line: 375, col: 37, offset: 15991},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 375, col: 37, offset: 15991},
											expr: &ruleRefExpr{
												pos:  position{line: 375, col: 38, offset: 15992},
												name: "NEWLINE",
											},
										},
										&anyMatcher{
											line: 375, col: 46, offset: 16000,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 375, col: 50, offset: 16004},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "Word",
			pos:  position{line: 382, col: 1, offset: 16175},
			expr: &actionExpr{
				pos: position{line: 382, col: 9, offset: 16183},
				run: (*parser).callonWord1,
				expr: &oneOrMoreExpr{
					pos: position{line: 382, col: 9, offset: 16183},
					expr: &seqExpr{
						pos: position{line: 382, col: 10, offset: 16184},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 382, col: 10, offset: 16184},
								expr: &ruleRefExpr{
									pos:  position{line: 382, col: 11, offset: 16185},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 382, col: 19, offset: 16193},
								expr: &ruleRefExpr{
									pos:  position{line: 382, col: 20, offset: 16194},
									name: "WS",
								},
							},
							&anyMatcher{
								line: 382, col: 23, offset: 16197,
							},
						},
					},
				},
			},
		},
		{
			name: "BlankLine",
			pos:  position{line: 386, col: 1, offset: 16237},
			expr: &actionExpr{
				pos: position{line: 386, col: 14, offset: 16250},
				run: (*parser).callonBlankLine1,
				expr: &seqExpr{
					pos: position{line: 386, col: 14, offset: 16250},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 386, col: 14, offset: 16250},
							expr: &ruleRefExpr{
								pos:  position{line: 386, col: 15, offset: 16251},
								name: "EOF",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 386, col: 19, offset: 16255},
							expr: &ruleRefExpr{
								pos:  position{line: 386, col: 19, offset: 16255},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 386, col: 23, offset: 16259},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "URL",
			pos:  position{line: 390, col: 1, offset: 16300},
			expr: &actionExpr{
				pos: position{line: 390, col: 8, offset: 16307},
				run: (*parser).callonURL1,
				expr: &oneOrMoreExpr{
					pos: position{line: 390, col: 8, offset: 16307},
					expr: &seqExpr{
						pos: position{line: 390, col: 9, offset: 16308},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 390, col: 9, offset: 16308},
								expr: &ruleRefExpr{
									pos:  position{line: 390, col: 10, offset: 16309},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 390, col: 18, offset: 16317},
								expr: &ruleRefExpr{
									pos:  position{line: 390, col: 19, offset: 16318},
									name: "WS",
								},
							},
							&notExpr{
								pos: position{line: 390, col: 22, offset: 16321},
								expr: &litMatcher{
									pos:        position{line: 390, col: 23, offset: 16322},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 390, col: 27, offset: 16326},
								expr: &litMatcher{
									pos:        position{line: 390, col: 28, offset: 16327},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 390, col: 32, offset: 16331,
							},
						},
					},
				},
			},
		},
		{
			name: "ID",
			pos:  position{line: 394, col: 1, offset: 16371},
			expr: &actionExpr{
				pos: position{line: 394, col: 7, offset: 16377},
				run: (*parser).callonID1,
				expr: &oneOrMoreExpr{
					pos: position{line: 394, col: 7, offset: 16377},
					expr: &seqExpr{
						pos: position{line: 394, col: 8, offset: 16378},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 394, col: 8, offset: 16378},
								expr: &ruleRefExpr{
									pos:  position{line: 394, col: 9, offset: 16379},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 394, col: 17, offset: 16387},
								expr: &ruleRefExpr{
									pos:  position{line: 394, col: 18, offset: 16388},
									name: "WS",
								},
							},
							&notExpr{
								pos: position{line: 394, col: 21, offset: 16391},
								expr: &litMatcher{
									pos:        position{line: 394, col: 22, offset: 16392},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 394, col: 26, offset: 16396},
								expr: &litMatcher{
									pos:        position{line: 394, col: 27, offset: 16397},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 394, col: 31, offset: 16401,
							},
						},
					},
				},
			},
		},
		{
			name: "URL_TEXT",
			pos:  position{line: 398, col: 1, offset: 16441},
			expr: &actionExpr{
				pos: position{line: 398, col: 13, offset: 16453},
				run: (*parser).callonURL_TEXT1,
				expr: &oneOrMoreExpr{
					pos: position{line: 398, col: 13, offset: 16453},
					expr: &seqExpr{
						pos: position{line: 398, col: 14, offset: 16454},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 398, col: 14, offset: 16454},
								expr: &ruleRefExpr{
									pos:  position{line: 398, col: 15, offset: 16455},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 398, col: 23, offset: 16463},
								expr: &litMatcher{
									pos:        position{line: 398, col: 24, offset: 16464},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 398, col: 28, offset: 16468},
								expr: &litMatcher{
									pos:        position{line: 398, col: 29, offset: 16469},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 398, col: 33, offset: 16473,
							},
						},
					},
				},
			},
		},
		{
			name: "URL_SCHEME",
			pos:  position{line: 402, col: 1, offset: 16513},
			expr: &choiceExpr{
				pos: position{line: 402, col: 15, offset: 16527},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 402, col: 15, offset: 16527},
						val:        "http://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 402, col: 27, offset: 16539},
						val:        "https://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 402, col: 40, offset: 16552},
						val:        "ftp://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 402, col: 51, offset: 16563},
						val:        "irc://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 402, col: 62, offset: 16574},
						val:        "mailto:",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "DIGIT",
			pos:  position{line: 404, col: 1, offset: 16585},
			expr: &charClassMatcher{
				pos:        position{line: 404, col: 13, offset: 16597},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NEWLINE",
			pos:  position{line: 406, col: 1, offset: 16604},
			expr: &choiceExpr{
				pos: position{line: 406, col: 13, offset: 16616},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 406, col: 13, offset: 16616},
						val:        "\r\n",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 406, col: 22, offset: 16625},
						val:        "\r",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 406, col: 29, offset: 16632},
						val:        "\n",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "WS",
			pos:  position{line: 408, col: 1, offset: 16638},
			expr: &choiceExpr{
				pos: position{line: 408, col: 13, offset: 16650},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 408, col: 13, offset: 16650},
						val:        " ",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 408, col: 19, offset: 16656},
						run: (*parser).callonWS3,
						expr: &litMatcher{
							pos:        position{line: 408, col: 19, offset: 16656},
							val:        "\t",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 412, col: 1, offset: 16701},
			expr: &notExpr{
				pos: position{line: 412, col: 13, offset: 16713},
				expr: &anyMatcher{
					line: 412, col: 14, offset: 16714,
				},
			},
		},
		{
			name: "EOL",
			pos:  position{line: 414, col: 1, offset: 16717},
			expr: &choiceExpr{
				pos: position{line: 414, col: 13, offset: 16729},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 414, col: 13, offset: 16729},
						name: "NEWLINE",
					},
					&ruleRefExpr{
						pos:  position{line: 414, col: 23, offset: 16739},
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
	return types.NewYamlFrontMatter(content.([]interface{}))
}

func (p *parser) callonFrontMatter1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFrontMatter1(stack["content"])
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
	// needs an "EOL" but does not consume it here.
	return types.NewInlineContent(c.text, elements.([]interface{}))
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
	// unbalanced `**` vs `*` punctuation
	result := append([]interface{}{"_"}, content.([]interface{}))
	return types.NewQuotedText(types.Italic, result)
}

func (p *parser) callonItalicTextUnbalancedPunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onItalicTextUnbalancedPunctuation1(stack["content"])
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
	// unbalanced `**` vs `*` punctuation
	result := append([]interface{}{"`"}, content.([]interface{}))
	return types.NewQuotedText(types.Monospace, result)
}

func (p *parser) callonMonospaceTextUnbalancedPunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMonospaceTextUnbalancedPunctuation1(stack["content"])
}

func (c *current) onWordWithQuotePunctuation1() (interface{}, error) {
	// can have "*", "_" or "`" within, maybe because the user inserted another quote, or made an error (extra or missing space, for example)
	return c.text, nil
}

func (p *parser) callonWordWithQuotePunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onWordWithQuotePunctuation1()
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
		p.addErr(errInvalidEncoding)
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

	if p.pt.rn != utf8.RuneError || p.pt.w > 1 { // see utf8.DecodeRune
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
