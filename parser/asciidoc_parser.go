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
						name: "TableOfContentsMacro",
					},
					&ruleRefExpr{
						pos:  position{line: 24, col: 99, offset: 842},
						name: "List",
					},
					&ruleRefExpr{
						pos:  position{line: 24, col: 106, offset: 849},
						name: "BlockImage",
					},
					&ruleRefExpr{
						pos:  position{line: 24, col: 119, offset: 862},
						name: "LiteralBlock",
					},
					&ruleRefExpr{
						pos:  position{line: 24, col: 134, offset: 877},
						name: "DelimitedBlock",
					},
					&ruleRefExpr{
						pos:  position{line: 24, col: 151, offset: 894},
						name: "Paragraph",
					},
					&seqExpr{
						pos: position{line: 24, col: 164, offset: 907},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 24, col: 164, offset: 907},
								name: "ElementAttribute",
							},
							&ruleRefExpr{
								pos:  position{line: 24, col: 181, offset: 924},
								name: "EOL",
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 24, col: 188, offset: 931},
						name: "BlankLine",
					},
				},
			},
		},
		{
			name: "Preamble",
			pos:  position{line: 26, col: 1, offset: 986},
			expr: &actionExpr{
				pos: position{line: 26, col: 13, offset: 998},
				run: (*parser).callonPreamble1,
				expr: &labeledExpr{
					pos:   position{line: 26, col: 13, offset: 998},
					label: "elements",
					expr: &zeroOrMoreExpr{
						pos: position{line: 26, col: 23, offset: 1008},
						expr: &ruleRefExpr{
							pos:  position{line: 26, col: 23, offset: 1008},
							name: "StandaloneBlock",
						},
					},
				},
			},
		},
		{
			name: "FrontMatter",
			pos:  position{line: 33, col: 1, offset: 1194},
			expr: &ruleRefExpr{
				pos:  position{line: 33, col: 16, offset: 1209},
				name: "YamlFrontMatter",
			},
		},
		{
			name: "FrontMatter",
			pos:  position{line: 35, col: 1, offset: 1227},
			expr: &actionExpr{
				pos: position{line: 35, col: 16, offset: 1242},
				run: (*parser).callonFrontMatter1,
				expr: &seqExpr{
					pos: position{line: 35, col: 16, offset: 1242},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 35, col: 16, offset: 1242},
							name: "YamlFrontMatterToken",
						},
						&labeledExpr{
							pos:   position{line: 35, col: 37, offset: 1263},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 35, col: 46, offset: 1272},
								name: "YamlFrontMatterContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 35, col: 70, offset: 1296},
							name: "YamlFrontMatterToken",
						},
					},
				},
			},
		},
		{
			name: "YamlFrontMatterToken",
			pos:  position{line: 39, col: 1, offset: 1376},
			expr: &seqExpr{
				pos: position{line: 39, col: 26, offset: 1401},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 39, col: 26, offset: 1401},
						val:        "---",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 39, col: 32, offset: 1407},
						name: "EOL",
					},
				},
			},
		},
		{
			name: "YamlFrontMatterContent",
			pos:  position{line: 41, col: 1, offset: 1412},
			expr: &actionExpr{
				pos: position{line: 41, col: 27, offset: 1438},
				run: (*parser).callonYamlFrontMatterContent1,
				expr: &zeroOrMoreExpr{
					pos: position{line: 41, col: 27, offset: 1438},
					expr: &seqExpr{
						pos: position{line: 41, col: 28, offset: 1439},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 41, col: 28, offset: 1439},
								expr: &ruleRefExpr{
									pos:  position{line: 41, col: 29, offset: 1440},
									name: "YamlFrontMatterToken",
								},
							},
							&anyMatcher{
								line: 41, col: 50, offset: 1461,
							},
						},
					},
				},
			},
		},
		{
			name: "DocumentHeader",
			pos:  position{line: 49, col: 1, offset: 1685},
			expr: &actionExpr{
				pos: position{line: 49, col: 19, offset: 1703},
				run: (*parser).callonDocumentHeader1,
				expr: &seqExpr{
					pos: position{line: 49, col: 19, offset: 1703},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 49, col: 19, offset: 1703},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 49, col: 27, offset: 1711},
								name: "DocumentTitle",
							},
						},
						&labeledExpr{
							pos:   position{line: 49, col: 42, offset: 1726},
							label: "authors",
							expr: &zeroOrOneExpr{
								pos: position{line: 49, col: 51, offset: 1735},
								expr: &ruleRefExpr{
									pos:  position{line: 49, col: 51, offset: 1735},
									name: "DocumentAuthors",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 49, col: 69, offset: 1753},
							label: "revision",
							expr: &zeroOrOneExpr{
								pos: position{line: 49, col: 79, offset: 1763},
								expr: &ruleRefExpr{
									pos:  position{line: 49, col: 79, offset: 1763},
									name: "DocumentRevision",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 49, col: 98, offset: 1782},
							label: "otherAttributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 49, col: 115, offset: 1799},
								expr: &ruleRefExpr{
									pos:  position{line: 49, col: 115, offset: 1799},
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
			pos:  position{line: 53, col: 1, offset: 1930},
			expr: &actionExpr{
				pos: position{line: 53, col: 18, offset: 1947},
				run: (*parser).callonDocumentTitle1,
				expr: &seqExpr{
					pos: position{line: 53, col: 18, offset: 1947},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 53, col: 18, offset: 1947},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 53, col: 29, offset: 1958},
								expr: &ruleRefExpr{
									pos:  position{line: 53, col: 30, offset: 1959},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 53, col: 49, offset: 1978},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 53, col: 56, offset: 1985},
								val:        "=",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 53, col: 61, offset: 1990},
							expr: &ruleRefExpr{
								pos:  position{line: 53, col: 61, offset: 1990},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 53, col: 65, offset: 1994},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 53, col: 73, offset: 2002},
								name: "InlineContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 53, col: 87, offset: 2016},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthors",
			pos:  position{line: 57, col: 1, offset: 2120},
			expr: &choiceExpr{
				pos: position{line: 57, col: 20, offset: 2139},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 57, col: 20, offset: 2139},
						name: "DocumentAuthorsInlineForm",
					},
					&ruleRefExpr{
						pos:  position{line: 57, col: 48, offset: 2167},
						name: "DocumentAuthorsAttributeForm",
					},
				},
			},
		},
		{
			name: "DocumentAuthorsInlineForm",
			pos:  position{line: 59, col: 1, offset: 2197},
			expr: &actionExpr{
				pos: position{line: 59, col: 30, offset: 2226},
				run: (*parser).callonDocumentAuthorsInlineForm1,
				expr: &seqExpr{
					pos: position{line: 59, col: 30, offset: 2226},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 59, col: 30, offset: 2226},
							expr: &ruleRefExpr{
								pos:  position{line: 59, col: 30, offset: 2226},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 59, col: 34, offset: 2230},
							expr: &litMatcher{
								pos:        position{line: 59, col: 35, offset: 2231},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 59, col: 39, offset: 2235},
							label: "authors",
							expr: &oneOrMoreExpr{
								pos: position{line: 59, col: 48, offset: 2244},
								expr: &ruleRefExpr{
									pos:  position{line: 59, col: 48, offset: 2244},
									name: "DocumentAuthor",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 59, col: 65, offset: 2261},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthorsAttributeForm",
			pos:  position{line: 63, col: 1, offset: 2331},
			expr: &actionExpr{
				pos: position{line: 63, col: 33, offset: 2363},
				run: (*parser).callonDocumentAuthorsAttributeForm1,
				expr: &seqExpr{
					pos: position{line: 63, col: 33, offset: 2363},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 63, col: 33, offset: 2363},
							expr: &ruleRefExpr{
								pos:  position{line: 63, col: 33, offset: 2363},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 63, col: 37, offset: 2367},
							val:        ":author:",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 63, col: 48, offset: 2378},
							label: "author",
							expr: &ruleRefExpr{
								pos:  position{line: 63, col: 56, offset: 2386},
								name: "DocumentAuthor",
							},
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthor",
			pos:  position{line: 67, col: 1, offset: 2479},
			expr: &actionExpr{
				pos: position{line: 67, col: 19, offset: 2497},
				run: (*parser).callonDocumentAuthor1,
				expr: &seqExpr{
					pos: position{line: 67, col: 19, offset: 2497},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 67, col: 19, offset: 2497},
							expr: &ruleRefExpr{
								pos:  position{line: 67, col: 19, offset: 2497},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 67, col: 23, offset: 2501},
							label: "namePart1",
							expr: &ruleRefExpr{
								pos:  position{line: 67, col: 34, offset: 2512},
								name: "DocumentAuthorNamePart",
							},
						},
						&labeledExpr{
							pos:   position{line: 67, col: 58, offset: 2536},
							label: "namePart2",
							expr: &zeroOrOneExpr{
								pos: position{line: 67, col: 68, offset: 2546},
								expr: &ruleRefExpr{
									pos:  position{line: 67, col: 69, offset: 2547},
									name: "DocumentAuthorNamePart",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 67, col: 94, offset: 2572},
							label: "namePart3",
							expr: &zeroOrOneExpr{
								pos: position{line: 67, col: 104, offset: 2582},
								expr: &ruleRefExpr{
									pos:  position{line: 67, col: 105, offset: 2583},
									name: "DocumentAuthorNamePart",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 67, col: 130, offset: 2608},
							label: "email",
							expr: &zeroOrOneExpr{
								pos: position{line: 67, col: 136, offset: 2614},
								expr: &ruleRefExpr{
									pos:  position{line: 67, col: 137, offset: 2615},
									name: "DocumentAuthorEmail",
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 67, col: 159, offset: 2637},
							expr: &ruleRefExpr{
								pos:  position{line: 67, col: 159, offset: 2637},
								name: "WS",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 67, col: 163, offset: 2641},
							expr: &litMatcher{
								pos:        position{line: 67, col: 163, offset: 2641},
								val:        ";",
								ignoreCase: false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 67, col: 168, offset: 2646},
							expr: &ruleRefExpr{
								pos:  position{line: 67, col: 168, offset: 2646},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthorNamePart",
			pos:  position{line: 72, col: 1, offset: 2811},
			expr: &seqExpr{
				pos: position{line: 72, col: 27, offset: 2837},
				exprs: []interface{}{
					&notExpr{
						pos: position{line: 72, col: 27, offset: 2837},
						expr: &litMatcher{
							pos:        position{line: 72, col: 28, offset: 2838},
							val:        "<",
							ignoreCase: false,
						},
					},
					&notExpr{
						pos: position{line: 72, col: 32, offset: 2842},
						expr: &litMatcher{
							pos:        position{line: 72, col: 33, offset: 2843},
							val:        ";",
							ignoreCase: false,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 72, col: 37, offset: 2847},
						name: "Characters",
					},
					&zeroOrMoreExpr{
						pos: position{line: 72, col: 48, offset: 2858},
						expr: &ruleRefExpr{
							pos:  position{line: 72, col: 48, offset: 2858},
							name: "WS",
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthorEmail",
			pos:  position{line: 74, col: 1, offset: 2863},
			expr: &seqExpr{
				pos: position{line: 74, col: 24, offset: 2886},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 74, col: 24, offset: 2886},
						val:        "<",
						ignoreCase: false,
					},
					&labeledExpr{
						pos:   position{line: 74, col: 28, offset: 2890},
						label: "email",
						expr: &oneOrMoreExpr{
							pos: position{line: 74, col: 34, offset: 2896},
							expr: &seqExpr{
								pos: position{line: 74, col: 35, offset: 2897},
								exprs: []interface{}{
									&notExpr{
										pos: position{line: 74, col: 35, offset: 2897},
										expr: &litMatcher{
											pos:        position{line: 74, col: 36, offset: 2898},
											val:        ">",
											ignoreCase: false,
										},
									},
									&notExpr{
										pos: position{line: 74, col: 40, offset: 2902},
										expr: &ruleRefExpr{
											pos:  position{line: 74, col: 41, offset: 2903},
											name: "EOL",
										},
									},
									&anyMatcher{
										line: 74, col: 45, offset: 2907,
									},
								},
							},
						},
					},
					&litMatcher{
						pos:        position{line: 74, col: 49, offset: 2911},
						val:        ">",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "DocumentRevision",
			pos:  position{line: 78, col: 1, offset: 3047},
			expr: &actionExpr{
				pos: position{line: 78, col: 21, offset: 3067},
				run: (*parser).callonDocumentRevision1,
				expr: &seqExpr{
					pos: position{line: 78, col: 21, offset: 3067},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 78, col: 21, offset: 3067},
							expr: &ruleRefExpr{
								pos:  position{line: 78, col: 21, offset: 3067},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 78, col: 25, offset: 3071},
							expr: &litMatcher{
								pos:        position{line: 78, col: 26, offset: 3072},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 78, col: 30, offset: 3076},
							label: "revnumber",
							expr: &zeroOrOneExpr{
								pos: position{line: 78, col: 40, offset: 3086},
								expr: &ruleRefExpr{
									pos:  position{line: 78, col: 41, offset: 3087},
									name: "DocumentRevisionNumber",
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 78, col: 66, offset: 3112},
							expr: &litMatcher{
								pos:        position{line: 78, col: 66, offset: 3112},
								val:        ",",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 78, col: 71, offset: 3117},
							label: "revdate",
							expr: &zeroOrOneExpr{
								pos: position{line: 78, col: 79, offset: 3125},
								expr: &ruleRefExpr{
									pos:  position{line: 78, col: 80, offset: 3126},
									name: "DocumentRevisionDate",
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 78, col: 103, offset: 3149},
							expr: &litMatcher{
								pos:        position{line: 78, col: 103, offset: 3149},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 78, col: 108, offset: 3154},
							label: "revremark",
							expr: &zeroOrOneExpr{
								pos: position{line: 78, col: 118, offset: 3164},
								expr: &ruleRefExpr{
									pos:  position{line: 78, col: 119, offset: 3165},
									name: "DocumentRevisionRemark",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 78, col: 144, offset: 3190},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentRevisionNumber",
			pos:  position{line: 83, col: 1, offset: 3363},
			expr: &choiceExpr{
				pos: position{line: 83, col: 27, offset: 3389},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 83, col: 27, offset: 3389},
						exprs: []interface{}{
							&litMatcher{
								pos:        position{line: 83, col: 27, offset: 3389},
								val:        "v",
								ignoreCase: true,
							},
							&ruleRefExpr{
								pos:  position{line: 83, col: 32, offset: 3394},
								name: "DIGIT",
							},
							&zeroOrMoreExpr{
								pos: position{line: 83, col: 39, offset: 3401},
								expr: &seqExpr{
									pos: position{line: 83, col: 40, offset: 3402},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 83, col: 40, offset: 3402},
											expr: &ruleRefExpr{
												pos:  position{line: 83, col: 41, offset: 3403},
												name: "EOL",
											},
										},
										&notExpr{
											pos: position{line: 83, col: 45, offset: 3407},
											expr: &litMatcher{
												pos:        position{line: 83, col: 46, offset: 3408},
												val:        ",",
												ignoreCase: false,
											},
										},
										&notExpr{
											pos: position{line: 83, col: 50, offset: 3412},
											expr: &litMatcher{
												pos:        position{line: 83, col: 51, offset: 3413},
												val:        ":",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 83, col: 55, offset: 3417,
										},
									},
								},
							},
						},
					},
					&seqExpr{
						pos: position{line: 83, col: 61, offset: 3423},
						exprs: []interface{}{
							&zeroOrOneExpr{
								pos: position{line: 83, col: 61, offset: 3423},
								expr: &litMatcher{
									pos:        position{line: 83, col: 61, offset: 3423},
									val:        "v",
									ignoreCase: true,
								},
							},
							&ruleRefExpr{
								pos:  position{line: 83, col: 67, offset: 3429},
								name: "DIGIT",
							},
							&zeroOrMoreExpr{
								pos: position{line: 83, col: 74, offset: 3436},
								expr: &seqExpr{
									pos: position{line: 83, col: 75, offset: 3437},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 83, col: 75, offset: 3437},
											expr: &ruleRefExpr{
												pos:  position{line: 83, col: 76, offset: 3438},
												name: "EOL",
											},
										},
										&notExpr{
											pos: position{line: 83, col: 80, offset: 3442},
											expr: &litMatcher{
												pos:        position{line: 83, col: 81, offset: 3443},
												val:        ",",
												ignoreCase: false,
											},
										},
										&notExpr{
											pos: position{line: 83, col: 85, offset: 3447},
											expr: &litMatcher{
												pos:        position{line: 83, col: 86, offset: 3448},
												val:        ":",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 83, col: 90, offset: 3452,
										},
									},
								},
							},
							&zeroOrMoreExpr{
								pos: position{line: 83, col: 94, offset: 3456},
								expr: &ruleRefExpr{
									pos:  position{line: 83, col: 94, offset: 3456},
									name: "WS",
								},
							},
							&andExpr{
								pos: position{line: 83, col: 98, offset: 3460},
								expr: &litMatcher{
									pos:        position{line: 83, col: 99, offset: 3461},
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
			pos:  position{line: 84, col: 1, offset: 3465},
			expr: &zeroOrMoreExpr{
				pos: position{line: 84, col: 25, offset: 3489},
				expr: &seqExpr{
					pos: position{line: 84, col: 26, offset: 3490},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 84, col: 26, offset: 3490},
							expr: &ruleRefExpr{
								pos:  position{line: 84, col: 27, offset: 3491},
								name: "EOL",
							},
						},
						&notExpr{
							pos: position{line: 84, col: 31, offset: 3495},
							expr: &litMatcher{
								pos:        position{line: 84, col: 32, offset: 3496},
								val:        ":",
								ignoreCase: false,
							},
						},
						&anyMatcher{
							line: 84, col: 36, offset: 3500,
						},
					},
				},
			},
		},
		{
			name: "DocumentRevisionRemark",
			pos:  position{line: 85, col: 1, offset: 3505},
			expr: &zeroOrMoreExpr{
				pos: position{line: 85, col: 27, offset: 3531},
				expr: &seqExpr{
					pos: position{line: 85, col: 28, offset: 3532},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 85, col: 28, offset: 3532},
							expr: &ruleRefExpr{
								pos:  position{line: 85, col: 29, offset: 3533},
								name: "EOL",
							},
						},
						&anyMatcher{
							line: 85, col: 33, offset: 3537,
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclaration",
			pos:  position{line: 90, col: 1, offset: 3657},
			expr: &choiceExpr{
				pos: position{line: 90, col: 33, offset: 3689},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 90, col: 33, offset: 3689},
						name: "DocumentAttributeDeclarationWithNameOnly",
					},
					&ruleRefExpr{
						pos:  position{line: 90, col: 76, offset: 3732},
						name: "DocumentAttributeDeclarationWithNameAndValue",
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclarationWithNameOnly",
			pos:  position{line: 92, col: 1, offset: 3779},
			expr: &actionExpr{
				pos: position{line: 92, col: 45, offset: 3823},
				run: (*parser).callonDocumentAttributeDeclarationWithNameOnly1,
				expr: &seqExpr{
					pos: position{line: 92, col: 45, offset: 3823},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 92, col: 45, offset: 3823},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 92, col: 49, offset: 3827},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 92, col: 55, offset: 3833},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 92, col: 70, offset: 3848},
							val:        ":",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 92, col: 74, offset: 3852},
							expr: &ruleRefExpr{
								pos:  position{line: 92, col: 74, offset: 3852},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 92, col: 78, offset: 3856},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclarationWithNameAndValue",
			pos:  position{line: 96, col: 1, offset: 3941},
			expr: &actionExpr{
				pos: position{line: 96, col: 49, offset: 3989},
				run: (*parser).callonDocumentAttributeDeclarationWithNameAndValue1,
				expr: &seqExpr{
					pos: position{line: 96, col: 49, offset: 3989},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 96, col: 49, offset: 3989},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 96, col: 53, offset: 3993},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 96, col: 59, offset: 3999},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 96, col: 74, offset: 4014},
							val:        ":",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 96, col: 78, offset: 4018},
							expr: &ruleRefExpr{
								pos:  position{line: 96, col: 78, offset: 4018},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 96, col: 82, offset: 4022},
							label: "value",
							expr: &zeroOrMoreExpr{
								pos: position{line: 96, col: 88, offset: 4028},
								expr: &seqExpr{
									pos: position{line: 96, col: 89, offset: 4029},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 96, col: 89, offset: 4029},
											expr: &ruleRefExpr{
												pos:  position{line: 96, col: 90, offset: 4030},
												name: "NEWLINE",
											},
										},
										&anyMatcher{
											line: 96, col: 98, offset: 4038,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 96, col: 102, offset: 4042},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeReset",
			pos:  position{line: 100, col: 1, offset: 4145},
			expr: &choiceExpr{
				pos: position{line: 100, col: 27, offset: 4171},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 100, col: 27, offset: 4171},
						name: "DocumentAttributeResetWithSectionTitleBangSymbol",
					},
					&ruleRefExpr{
						pos:  position{line: 100, col: 78, offset: 4222},
						name: "DocumentAttributeResetWithTrailingBangSymbol",
					},
				},
			},
		},
		{
			name: "DocumentAttributeResetWithSectionTitleBangSymbol",
			pos:  position{line: 102, col: 1, offset: 4268},
			expr: &actionExpr{
				pos: position{line: 102, col: 53, offset: 4320},
				run: (*parser).callonDocumentAttributeResetWithSectionTitleBangSymbol1,
				expr: &seqExpr{
					pos: position{line: 102, col: 53, offset: 4320},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 102, col: 53, offset: 4320},
							val:        ":!",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 102, col: 58, offset: 4325},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 102, col: 64, offset: 4331},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 102, col: 79, offset: 4346},
							val:        ":",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 102, col: 83, offset: 4350},
							expr: &ruleRefExpr{
								pos:  position{line: 102, col: 83, offset: 4350},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 102, col: 87, offset: 4354},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeResetWithTrailingBangSymbol",
			pos:  position{line: 106, col: 1, offset: 4428},
			expr: &actionExpr{
				pos: position{line: 106, col: 49, offset: 4476},
				run: (*parser).callonDocumentAttributeResetWithTrailingBangSymbol1,
				expr: &seqExpr{
					pos: position{line: 106, col: 49, offset: 4476},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 106, col: 49, offset: 4476},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 106, col: 53, offset: 4480},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 106, col: 59, offset: 4486},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 106, col: 74, offset: 4501},
							val:        "!:",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 106, col: 79, offset: 4506},
							expr: &ruleRefExpr{
								pos:  position{line: 106, col: 79, offset: 4506},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 106, col: 83, offset: 4510},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeSubstitution",
			pos:  position{line: 111, col: 1, offset: 4585},
			expr: &actionExpr{
				pos: position{line: 111, col: 34, offset: 4618},
				run: (*parser).callonDocumentAttributeSubstitution1,
				expr: &seqExpr{
					pos: position{line: 111, col: 34, offset: 4618},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 111, col: 34, offset: 4618},
							val:        "{",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 111, col: 38, offset: 4622},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 111, col: 44, offset: 4628},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 111, col: 59, offset: 4643},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "AttributeName",
			pos:  position{line: 118, col: 1, offset: 4897},
			expr: &seqExpr{
				pos: position{line: 118, col: 18, offset: 4914},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 118, col: 19, offset: 4915},
						alternatives: []interface{}{
							&charClassMatcher{
								pos:        position{line: 118, col: 19, offset: 4915},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
							&charClassMatcher{
								pos:        position{line: 118, col: 27, offset: 4923},
								val:        "[a-z]",
								ranges:     []rune{'a', 'z'},
								ignoreCase: false,
								inverted:   false,
							},
							&charClassMatcher{
								pos:        position{line: 118, col: 35, offset: 4931},
								val:        "[0-9]",
								ranges:     []rune{'0', '9'},
								ignoreCase: false,
								inverted:   false,
							},
							&litMatcher{
								pos:        position{line: 118, col: 43, offset: 4939},
								val:        "_",
								ignoreCase: false,
							},
						},
					},
					&zeroOrMoreExpr{
						pos: position{line: 118, col: 48, offset: 4944},
						expr: &choiceExpr{
							pos: position{line: 118, col: 49, offset: 4945},
							alternatives: []interface{}{
								&charClassMatcher{
									pos:        position{line: 118, col: 49, offset: 4945},
									val:        "[A-Z]",
									ranges:     []rune{'A', 'Z'},
									ignoreCase: false,
									inverted:   false,
								},
								&charClassMatcher{
									pos:        position{line: 118, col: 57, offset: 4953},
									val:        "[a-z]",
									ranges:     []rune{'a', 'z'},
									ignoreCase: false,
									inverted:   false,
								},
								&charClassMatcher{
									pos:        position{line: 118, col: 65, offset: 4961},
									val:        "[0-9]",
									ranges:     []rune{'0', '9'},
									ignoreCase: false,
									inverted:   false,
								},
								&litMatcher{
									pos:        position{line: 118, col: 73, offset: 4969},
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
			name: "TableOfContentsMacro",
			pos:  position{line: 123, col: 1, offset: 5089},
			expr: &seqExpr{
				pos: position{line: 123, col: 25, offset: 5113},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 123, col: 25, offset: 5113},
						val:        "toc::[]",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 123, col: 35, offset: 5123},
						name: "NEWLINE",
					},
				},
			},
		},
		{
			name: "Section",
			pos:  position{line: 128, col: 1, offset: 5236},
			expr: &choiceExpr{
				pos: position{line: 128, col: 12, offset: 5247},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 128, col: 12, offset: 5247},
						name: "Section1",
					},
					&ruleRefExpr{
						pos:  position{line: 128, col: 23, offset: 5258},
						name: "Section2",
					},
					&ruleRefExpr{
						pos:  position{line: 128, col: 34, offset: 5269},
						name: "Section3",
					},
					&ruleRefExpr{
						pos:  position{line: 128, col: 45, offset: 5280},
						name: "Section4",
					},
					&ruleRefExpr{
						pos:  position{line: 128, col: 56, offset: 5291},
						name: "Section5",
					},
				},
			},
		},
		{
			name: "Section1",
			pos:  position{line: 131, col: 1, offset: 5302},
			expr: &actionExpr{
				pos: position{line: 131, col: 13, offset: 5314},
				run: (*parser).callonSection11,
				expr: &seqExpr{
					pos: position{line: 131, col: 13, offset: 5314},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 131, col: 13, offset: 5314},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 131, col: 21, offset: 5322},
								name: "Section1Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 131, col: 36, offset: 5337},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 131, col: 46, offset: 5347},
								expr: &ruleRefExpr{
									pos:  position{line: 131, col: 46, offset: 5347},
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
			pos:  position{line: 135, col: 1, offset: 5455},
			expr: &actionExpr{
				pos: position{line: 135, col: 18, offset: 5472},
				run: (*parser).callonSection1Block1,
				expr: &seqExpr{
					pos: position{line: 135, col: 18, offset: 5472},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 135, col: 18, offset: 5472},
							expr: &ruleRefExpr{
								pos:  position{line: 135, col: 19, offset: 5473},
								name: "Section1",
							},
						},
						&labeledExpr{
							pos:   position{line: 135, col: 28, offset: 5482},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 135, col: 37, offset: 5491},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 135, col: 37, offset: 5491},
										name: "Section2",
									},
									&ruleRefExpr{
										pos:  position{line: 135, col: 48, offset: 5502},
										name: "Section3",
									},
									&ruleRefExpr{
										pos:  position{line: 135, col: 59, offset: 5513},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 135, col: 70, offset: 5524},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 135, col: 81, offset: 5535},
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
			pos:  position{line: 139, col: 1, offset: 5600},
			expr: &actionExpr{
				pos: position{line: 139, col: 13, offset: 5612},
				run: (*parser).callonSection21,
				expr: &seqExpr{
					pos: position{line: 139, col: 13, offset: 5612},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 139, col: 13, offset: 5612},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 139, col: 21, offset: 5620},
								name: "Section2Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 139, col: 36, offset: 5635},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 139, col: 46, offset: 5645},
								expr: &ruleRefExpr{
									pos:  position{line: 139, col: 46, offset: 5645},
									name: "Section2Block",
								},
							},
						},
						&andExpr{
							pos: position{line: 139, col: 62, offset: 5661},
							expr: &zeroOrMoreExpr{
								pos: position{line: 139, col: 63, offset: 5662},
								expr: &ruleRefExpr{
									pos:  position{line: 139, col: 64, offset: 5663},
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
			pos:  position{line: 143, col: 1, offset: 5766},
			expr: &actionExpr{
				pos: position{line: 143, col: 18, offset: 5783},
				run: (*parser).callonSection2Block1,
				expr: &seqExpr{
					pos: position{line: 143, col: 18, offset: 5783},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 143, col: 18, offset: 5783},
							expr: &ruleRefExpr{
								pos:  position{line: 143, col: 19, offset: 5784},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 143, col: 28, offset: 5793},
							expr: &ruleRefExpr{
								pos:  position{line: 143, col: 29, offset: 5794},
								name: "Section2",
							},
						},
						&labeledExpr{
							pos:   position{line: 143, col: 38, offset: 5803},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 143, col: 47, offset: 5812},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 143, col: 47, offset: 5812},
										name: "Section3",
									},
									&ruleRefExpr{
										pos:  position{line: 143, col: 58, offset: 5823},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 143, col: 69, offset: 5834},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 143, col: 80, offset: 5845},
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
			pos:  position{line: 147, col: 1, offset: 5910},
			expr: &actionExpr{
				pos: position{line: 147, col: 13, offset: 5922},
				run: (*parser).callonSection31,
				expr: &seqExpr{
					pos: position{line: 147, col: 13, offset: 5922},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 147, col: 13, offset: 5922},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 147, col: 21, offset: 5930},
								name: "Section3Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 147, col: 36, offset: 5945},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 147, col: 46, offset: 5955},
								expr: &ruleRefExpr{
									pos:  position{line: 147, col: 46, offset: 5955},
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
			pos:  position{line: 151, col: 1, offset: 6063},
			expr: &actionExpr{
				pos: position{line: 151, col: 18, offset: 6080},
				run: (*parser).callonSection3Block1,
				expr: &seqExpr{
					pos: position{line: 151, col: 18, offset: 6080},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 151, col: 18, offset: 6080},
							expr: &ruleRefExpr{
								pos:  position{line: 151, col: 19, offset: 6081},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 151, col: 28, offset: 6090},
							expr: &ruleRefExpr{
								pos:  position{line: 151, col: 29, offset: 6091},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 151, col: 38, offset: 6100},
							expr: &ruleRefExpr{
								pos:  position{line: 151, col: 39, offset: 6101},
								name: "Section3",
							},
						},
						&labeledExpr{
							pos:   position{line: 151, col: 48, offset: 6110},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 151, col: 57, offset: 6119},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 151, col: 57, offset: 6119},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 151, col: 68, offset: 6130},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 151, col: 79, offset: 6141},
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
			pos:  position{line: 155, col: 1, offset: 6206},
			expr: &actionExpr{
				pos: position{line: 155, col: 13, offset: 6218},
				run: (*parser).callonSection41,
				expr: &seqExpr{
					pos: position{line: 155, col: 13, offset: 6218},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 155, col: 13, offset: 6218},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 155, col: 21, offset: 6226},
								name: "Section4Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 155, col: 36, offset: 6241},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 155, col: 46, offset: 6251},
								expr: &ruleRefExpr{
									pos:  position{line: 155, col: 46, offset: 6251},
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
			pos:  position{line: 159, col: 1, offset: 6359},
			expr: &actionExpr{
				pos: position{line: 159, col: 18, offset: 6376},
				run: (*parser).callonSection4Block1,
				expr: &seqExpr{
					pos: position{line: 159, col: 18, offset: 6376},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 159, col: 18, offset: 6376},
							expr: &ruleRefExpr{
								pos:  position{line: 159, col: 19, offset: 6377},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 159, col: 28, offset: 6386},
							expr: &ruleRefExpr{
								pos:  position{line: 159, col: 29, offset: 6387},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 159, col: 38, offset: 6396},
							expr: &ruleRefExpr{
								pos:  position{line: 159, col: 39, offset: 6397},
								name: "Section3",
							},
						},
						&notExpr{
							pos: position{line: 159, col: 48, offset: 6406},
							expr: &ruleRefExpr{
								pos:  position{line: 159, col: 49, offset: 6407},
								name: "Section4",
							},
						},
						&labeledExpr{
							pos:   position{line: 159, col: 58, offset: 6416},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 159, col: 67, offset: 6425},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 159, col: 67, offset: 6425},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 159, col: 78, offset: 6436},
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
			pos:  position{line: 163, col: 1, offset: 6501},
			expr: &actionExpr{
				pos: position{line: 163, col: 13, offset: 6513},
				run: (*parser).callonSection51,
				expr: &seqExpr{
					pos: position{line: 163, col: 13, offset: 6513},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 163, col: 13, offset: 6513},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 163, col: 21, offset: 6521},
								name: "Section5Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 163, col: 36, offset: 6536},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 163, col: 46, offset: 6546},
								expr: &ruleRefExpr{
									pos:  position{line: 163, col: 46, offset: 6546},
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
			pos:  position{line: 167, col: 1, offset: 6654},
			expr: &actionExpr{
				pos: position{line: 167, col: 18, offset: 6671},
				run: (*parser).callonSection5Block1,
				expr: &seqExpr{
					pos: position{line: 167, col: 18, offset: 6671},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 167, col: 18, offset: 6671},
							expr: &ruleRefExpr{
								pos:  position{line: 167, col: 19, offset: 6672},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 167, col: 28, offset: 6681},
							expr: &ruleRefExpr{
								pos:  position{line: 167, col: 29, offset: 6682},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 167, col: 38, offset: 6691},
							expr: &ruleRefExpr{
								pos:  position{line: 167, col: 39, offset: 6692},
								name: "Section3",
							},
						},
						&notExpr{
							pos: position{line: 167, col: 48, offset: 6701},
							expr: &ruleRefExpr{
								pos:  position{line: 167, col: 49, offset: 6702},
								name: "Section4",
							},
						},
						&notExpr{
							pos: position{line: 167, col: 58, offset: 6711},
							expr: &ruleRefExpr{
								pos:  position{line: 167, col: 59, offset: 6712},
								name: "Section5",
							},
						},
						&labeledExpr{
							pos:   position{line: 167, col: 68, offset: 6721},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 167, col: 77, offset: 6730},
								name: "StandaloneBlock",
							},
						},
					},
				},
			},
		},
		{
			name: "SectionTitle",
			pos:  position{line: 175, col: 1, offset: 6906},
			expr: &choiceExpr{
				pos: position{line: 175, col: 17, offset: 6922},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 175, col: 17, offset: 6922},
						name: "Section1Title",
					},
					&ruleRefExpr{
						pos:  position{line: 175, col: 33, offset: 6938},
						name: "Section2Title",
					},
					&ruleRefExpr{
						pos:  position{line: 175, col: 49, offset: 6954},
						name: "Section3Title",
					},
					&ruleRefExpr{
						pos:  position{line: 175, col: 65, offset: 6970},
						name: "Section4Title",
					},
					&ruleRefExpr{
						pos:  position{line: 175, col: 81, offset: 6986},
						name: "Section5Title",
					},
				},
			},
		},
		{
			name: "Section1Title",
			pos:  position{line: 177, col: 1, offset: 7001},
			expr: &actionExpr{
				pos: position{line: 177, col: 18, offset: 7018},
				run: (*parser).callonSection1Title1,
				expr: &seqExpr{
					pos: position{line: 177, col: 18, offset: 7018},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 177, col: 18, offset: 7018},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 177, col: 29, offset: 7029},
								expr: &ruleRefExpr{
									pos:  position{line: 177, col: 30, offset: 7030},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 177, col: 49, offset: 7049},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 177, col: 56, offset: 7056},
								val:        "==",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 177, col: 62, offset: 7062},
							expr: &ruleRefExpr{
								pos:  position{line: 177, col: 62, offset: 7062},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 177, col: 66, offset: 7066},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 177, col: 74, offset: 7074},
								name: "InlineContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 177, col: 88, offset: 7088},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 177, col: 93, offset: 7093},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 177, col: 93, offset: 7093},
									expr: &ruleRefExpr{
										pos:  position{line: 177, col: 93, offset: 7093},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 177, col: 106, offset: 7106},
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
			pos:  position{line: 181, col: 1, offset: 7211},
			expr: &actionExpr{
				pos: position{line: 181, col: 18, offset: 7228},
				run: (*parser).callonSection2Title1,
				expr: &seqExpr{
					pos: position{line: 181, col: 18, offset: 7228},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 181, col: 18, offset: 7228},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 181, col: 29, offset: 7239},
								expr: &ruleRefExpr{
									pos:  position{line: 181, col: 30, offset: 7240},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 181, col: 49, offset: 7259},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 181, col: 56, offset: 7266},
								val:        "===",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 181, col: 63, offset: 7273},
							expr: &ruleRefExpr{
								pos:  position{line: 181, col: 63, offset: 7273},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 181, col: 67, offset: 7277},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 181, col: 75, offset: 7285},
								name: "InlineContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 181, col: 89, offset: 7299},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 181, col: 94, offset: 7304},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 181, col: 94, offset: 7304},
									expr: &ruleRefExpr{
										pos:  position{line: 181, col: 94, offset: 7304},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 181, col: 107, offset: 7317},
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
			pos:  position{line: 185, col: 1, offset: 7421},
			expr: &actionExpr{
				pos: position{line: 185, col: 18, offset: 7438},
				run: (*parser).callonSection3Title1,
				expr: &seqExpr{
					pos: position{line: 185, col: 18, offset: 7438},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 185, col: 18, offset: 7438},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 185, col: 29, offset: 7449},
								expr: &ruleRefExpr{
									pos:  position{line: 185, col: 30, offset: 7450},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 185, col: 49, offset: 7469},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 185, col: 56, offset: 7476},
								val:        "====",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 185, col: 64, offset: 7484},
							expr: &ruleRefExpr{
								pos:  position{line: 185, col: 64, offset: 7484},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 185, col: 68, offset: 7488},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 185, col: 76, offset: 7496},
								name: "InlineContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 185, col: 90, offset: 7510},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 185, col: 95, offset: 7515},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 185, col: 95, offset: 7515},
									expr: &ruleRefExpr{
										pos:  position{line: 185, col: 95, offset: 7515},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 185, col: 108, offset: 7528},
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
			pos:  position{line: 189, col: 1, offset: 7632},
			expr: &actionExpr{
				pos: position{line: 189, col: 18, offset: 7649},
				run: (*parser).callonSection4Title1,
				expr: &seqExpr{
					pos: position{line: 189, col: 18, offset: 7649},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 189, col: 18, offset: 7649},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 189, col: 29, offset: 7660},
								expr: &ruleRefExpr{
									pos:  position{line: 189, col: 30, offset: 7661},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 189, col: 49, offset: 7680},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 189, col: 56, offset: 7687},
								val:        "=====",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 189, col: 65, offset: 7696},
							expr: &ruleRefExpr{
								pos:  position{line: 189, col: 65, offset: 7696},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 189, col: 69, offset: 7700},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 189, col: 77, offset: 7708},
								name: "InlineContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 189, col: 91, offset: 7722},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 189, col: 96, offset: 7727},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 189, col: 96, offset: 7727},
									expr: &ruleRefExpr{
										pos:  position{line: 189, col: 96, offset: 7727},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 189, col: 109, offset: 7740},
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
			pos:  position{line: 193, col: 1, offset: 7844},
			expr: &actionExpr{
				pos: position{line: 193, col: 18, offset: 7861},
				run: (*parser).callonSection5Title1,
				expr: &seqExpr{
					pos: position{line: 193, col: 18, offset: 7861},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 193, col: 18, offset: 7861},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 193, col: 29, offset: 7872},
								expr: &ruleRefExpr{
									pos:  position{line: 193, col: 30, offset: 7873},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 193, col: 49, offset: 7892},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 193, col: 56, offset: 7899},
								val:        "======",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 193, col: 66, offset: 7909},
							expr: &ruleRefExpr{
								pos:  position{line: 193, col: 66, offset: 7909},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 193, col: 70, offset: 7913},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 193, col: 78, offset: 7921},
								name: "InlineContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 193, col: 92, offset: 7935},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 193, col: 97, offset: 7940},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 193, col: 97, offset: 7940},
									expr: &ruleRefExpr{
										pos:  position{line: 193, col: 97, offset: 7940},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 193, col: 110, offset: 7953},
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
			pos:  position{line: 200, col: 1, offset: 8163},
			expr: &actionExpr{
				pos: position{line: 200, col: 9, offset: 8171},
				run: (*parser).callonList1,
				expr: &seqExpr{
					pos: position{line: 200, col: 9, offset: 8171},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 200, col: 9, offset: 8171},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 200, col: 20, offset: 8182},
								expr: &ruleRefExpr{
									pos:  position{line: 200, col: 21, offset: 8183},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 202, col: 5, offset: 8275},
							label: "elements",
							expr: &oneOrMoreExpr{
								pos: position{line: 202, col: 14, offset: 8284},
								expr: &seqExpr{
									pos: position{line: 202, col: 15, offset: 8285},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 202, col: 15, offset: 8285},
											name: "ListItem",
										},
										&zeroOrOneExpr{
											pos: position{line: 202, col: 24, offset: 8294},
											expr: &ruleRefExpr{
												pos:  position{line: 202, col: 24, offset: 8294},
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
			pos:  position{line: 206, col: 1, offset: 8391},
			expr: &actionExpr{
				pos: position{line: 206, col: 13, offset: 8403},
				run: (*parser).callonListItem1,
				expr: &seqExpr{
					pos: position{line: 206, col: 13, offset: 8403},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 206, col: 13, offset: 8403},
							expr: &ruleRefExpr{
								pos:  position{line: 206, col: 13, offset: 8403},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 206, col: 17, offset: 8407},
							label: "level",
							expr: &choiceExpr{
								pos: position{line: 206, col: 24, offset: 8414},
								alternatives: []interface{}{
									&oneOrMoreExpr{
										pos: position{line: 206, col: 24, offset: 8414},
										expr: &litMatcher{
											pos:        position{line: 206, col: 24, offset: 8414},
											val:        "*",
											ignoreCase: false,
										},
									},
									&litMatcher{
										pos:        position{line: 206, col: 31, offset: 8421},
										val:        "-",
										ignoreCase: false,
									},
								},
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 206, col: 36, offset: 8426},
							expr: &ruleRefExpr{
								pos:  position{line: 206, col: 36, offset: 8426},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 206, col: 40, offset: 8430},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 206, col: 49, offset: 8439},
								name: "ListItemContent",
							},
						},
					},
				},
			},
		},
		{
			name: "ListItemContent",
			pos:  position{line: 210, col: 1, offset: 8536},
			expr: &actionExpr{
				pos: position{line: 210, col: 20, offset: 8555},
				run: (*parser).callonListItemContent1,
				expr: &labeledExpr{
					pos:   position{line: 210, col: 20, offset: 8555},
					label: "lines",
					expr: &oneOrMoreExpr{
						pos: position{line: 210, col: 26, offset: 8561},
						expr: &seqExpr{
							pos: position{line: 210, col: 27, offset: 8562},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 210, col: 27, offset: 8562},
									expr: &seqExpr{
										pos: position{line: 210, col: 29, offset: 8564},
										exprs: []interface{}{
											&zeroOrMoreExpr{
												pos: position{line: 210, col: 29, offset: 8564},
												expr: &ruleRefExpr{
													pos:  position{line: 210, col: 29, offset: 8564},
													name: "WS",
												},
											},
											&choiceExpr{
												pos: position{line: 210, col: 34, offset: 8569},
												alternatives: []interface{}{
													&oneOrMoreExpr{
														pos: position{line: 210, col: 34, offset: 8569},
														expr: &litMatcher{
															pos:        position{line: 210, col: 34, offset: 8569},
															val:        "*",
															ignoreCase: false,
														},
													},
													&litMatcher{
														pos:        position{line: 210, col: 41, offset: 8576},
														val:        "-",
														ignoreCase: false,
													},
												},
											},
											&oneOrMoreExpr{
												pos: position{line: 210, col: 46, offset: 8581},
												expr: &ruleRefExpr{
													pos:  position{line: 210, col: 46, offset: 8581},
													name: "WS",
												},
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 210, col: 51, offset: 8586},
									name: "InlineContent",
								},
								&ruleRefExpr{
									pos:  position{line: 210, col: 65, offset: 8600},
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
			pos:  position{line: 218, col: 1, offset: 8929},
			expr: &actionExpr{
				pos: position{line: 218, col: 14, offset: 8942},
				run: (*parser).callonParagraph1,
				expr: &seqExpr{
					pos: position{line: 218, col: 14, offset: 8942},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 218, col: 14, offset: 8942},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 218, col: 25, offset: 8953},
								expr: &ruleRefExpr{
									pos:  position{line: 218, col: 26, offset: 8954},
									name: "ElementAttribute",
								},
							},
						},
						&notExpr{
							pos: position{line: 218, col: 45, offset: 8973},
							expr: &seqExpr{
								pos: position{line: 218, col: 47, offset: 8975},
								exprs: []interface{}{
									&oneOrMoreExpr{
										pos: position{line: 218, col: 47, offset: 8975},
										expr: &litMatcher{
											pos:        position{line: 218, col: 47, offset: 8975},
											val:        "=",
											ignoreCase: false,
										},
									},
									&oneOrMoreExpr{
										pos: position{line: 218, col: 52, offset: 8980},
										expr: &ruleRefExpr{
											pos:  position{line: 218, col: 52, offset: 8980},
											name: "WS",
										},
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 218, col: 57, offset: 8985},
							label: "lines",
							expr: &oneOrMoreExpr{
								pos: position{line: 218, col: 63, offset: 8991},
								expr: &seqExpr{
									pos: position{line: 218, col: 64, offset: 8992},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 218, col: 64, offset: 8992},
											name: "InlineContent",
										},
										&ruleRefExpr{
											pos:  position{line: 218, col: 78, offset: 9006},
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
			pos:  position{line: 224, col: 1, offset: 9296},
			expr: &actionExpr{
				pos: position{line: 224, col: 18, offset: 9313},
				run: (*parser).callonInlineContent1,
				expr: &seqExpr{
					pos: position{line: 224, col: 18, offset: 9313},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 224, col: 18, offset: 9313},
							expr: &ruleRefExpr{
								pos:  position{line: 224, col: 19, offset: 9314},
								name: "BlockDelimiter",
							},
						},
						&labeledExpr{
							pos:   position{line: 224, col: 34, offset: 9329},
							label: "elements",
							expr: &oneOrMoreExpr{
								pos: position{line: 224, col: 43, offset: 9338},
								expr: &seqExpr{
									pos: position{line: 224, col: 44, offset: 9339},
									exprs: []interface{}{
										&zeroOrMoreExpr{
											pos: position{line: 224, col: 44, offset: 9339},
											expr: &ruleRefExpr{
												pos:  position{line: 224, col: 44, offset: 9339},
												name: "WS",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 224, col: 48, offset: 9343},
											name: "InlineElement",
										},
										&zeroOrMoreExpr{
											pos: position{line: 224, col: 62, offset: 9357},
											expr: &ruleRefExpr{
												pos:  position{line: 224, col: 62, offset: 9357},
												name: "WS",
											},
										},
									},
								},
							},
						},
						&andExpr{
							pos: position{line: 224, col: 68, offset: 9363},
							expr: &ruleRefExpr{
								pos:  position{line: 224, col: 69, offset: 9364},
								name: "EOL",
							},
						},
					},
				},
			},
		},
		{
			name: "InlineElement",
			pos:  position{line: 228, col: 1, offset: 9482},
			expr: &choiceExpr{
				pos: position{line: 228, col: 18, offset: 9499},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 228, col: 18, offset: 9499},
						name: "Passthrough",
					},
					&ruleRefExpr{
						pos:  position{line: 228, col: 32, offset: 9513},
						name: "InlineImage",
					},
					&ruleRefExpr{
						pos:  position{line: 228, col: 46, offset: 9527},
						name: "QuotedText",
					},
					&ruleRefExpr{
						pos:  position{line: 228, col: 59, offset: 9540},
						name: "ExternalLink",
					},
					&ruleRefExpr{
						pos:  position{line: 228, col: 74, offset: 9555},
						name: "DocumentAttributeSubstitution",
					},
					&ruleRefExpr{
						pos:  position{line: 228, col: 106, offset: 9587},
						name: "Characters",
					},
				},
			},
		},
		{
			name: "QuotedText",
			pos:  position{line: 233, col: 1, offset: 9838},
			expr: &choiceExpr{
				pos: position{line: 233, col: 15, offset: 9852},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 233, col: 15, offset: 9852},
						name: "BoldText",
					},
					&ruleRefExpr{
						pos:  position{line: 233, col: 26, offset: 9863},
						name: "ItalicText",
					},
					&ruleRefExpr{
						pos:  position{line: 233, col: 39, offset: 9876},
						name: "MonospaceText",
					},
					&ruleRefExpr{
						pos:  position{line: 234, col: 13, offset: 9904},
						name: "EscapedBoldText",
					},
					&ruleRefExpr{
						pos:  position{line: 234, col: 31, offset: 9922},
						name: "EscapedItalicText",
					},
					&ruleRefExpr{
						pos:  position{line: 234, col: 51, offset: 9942},
						name: "EscapedMonospaceText",
					},
				},
			},
		},
		{
			name: "BoldText",
			pos:  position{line: 236, col: 1, offset: 9964},
			expr: &choiceExpr{
				pos: position{line: 236, col: 13, offset: 9976},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 236, col: 13, offset: 9976},
						name: "BoldTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 236, col: 41, offset: 10004},
						name: "BoldTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 236, col: 73, offset: 10036},
						name: "BoldTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "BoldTextSimplePunctuation",
			pos:  position{line: 238, col: 1, offset: 10109},
			expr: &actionExpr{
				pos: position{line: 238, col: 30, offset: 10138},
				run: (*parser).callonBoldTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 238, col: 30, offset: 10138},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 238, col: 30, offset: 10138},
							expr: &litMatcher{
								pos:        position{line: 238, col: 31, offset: 10139},
								val:        "\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 238, col: 35, offset: 10143},
							val:        "*",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 238, col: 39, offset: 10147},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 238, col: 48, offset: 10156},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 238, col: 67, offset: 10175},
							val:        "*",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "BoldTextDoublePunctuation",
			pos:  position{line: 242, col: 1, offset: 10252},
			expr: &actionExpr{
				pos: position{line: 242, col: 30, offset: 10281},
				run: (*parser).callonBoldTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 242, col: 30, offset: 10281},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 242, col: 30, offset: 10281},
							expr: &litMatcher{
								pos:        position{line: 242, col: 31, offset: 10282},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 242, col: 36, offset: 10287},
							val:        "**",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 242, col: 41, offset: 10292},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 242, col: 50, offset: 10301},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 242, col: 69, offset: 10320},
							val:        "**",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "BoldTextUnbalancedPunctuation",
			pos:  position{line: 246, col: 1, offset: 10398},
			expr: &actionExpr{
				pos: position{line: 246, col: 34, offset: 10431},
				run: (*parser).callonBoldTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 246, col: 34, offset: 10431},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 246, col: 34, offset: 10431},
							expr: &litMatcher{
								pos:        position{line: 246, col: 35, offset: 10432},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 246, col: 40, offset: 10437},
							val:        "**",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 246, col: 45, offset: 10442},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 246, col: 54, offset: 10451},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 246, col: 73, offset: 10470},
							val:        "*",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedBoldText",
			pos:  position{line: 251, col: 1, offset: 10634},
			expr: &choiceExpr{
				pos: position{line: 251, col: 20, offset: 10653},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 251, col: 20, offset: 10653},
						name: "EscapedBoldTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 251, col: 55, offset: 10688},
						name: "EscapedBoldTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 251, col: 94, offset: 10727},
						name: "EscapedBoldTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "EscapedBoldTextSimplePunctuation",
			pos:  position{line: 253, col: 1, offset: 10807},
			expr: &actionExpr{
				pos: position{line: 253, col: 37, offset: 10843},
				run: (*parser).callonEscapedBoldTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 253, col: 37, offset: 10843},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 253, col: 37, offset: 10843},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 253, col: 50, offset: 10856},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 253, col: 50, offset: 10856},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 253, col: 54, offset: 10860},
										expr: &litMatcher{
											pos:        position{line: 253, col: 54, offset: 10860},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 253, col: 60, offset: 10866},
							val:        "*",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 253, col: 64, offset: 10870},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 253, col: 73, offset: 10879},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 253, col: 92, offset: 10898},
							val:        "*",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedBoldTextDoublePunctuation",
			pos:  position{line: 257, col: 1, offset: 11004},
			expr: &actionExpr{
				pos: position{line: 257, col: 37, offset: 11040},
				run: (*parser).callonEscapedBoldTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 257, col: 37, offset: 11040},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 257, col: 37, offset: 11040},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 257, col: 50, offset: 11053},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 257, col: 50, offset: 11053},
										val:        "\\\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 257, col: 55, offset: 11058},
										expr: &litMatcher{
											pos:        position{line: 257, col: 55, offset: 11058},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 257, col: 61, offset: 11064},
							val:        "**",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 257, col: 66, offset: 11069},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 257, col: 75, offset: 11078},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 257, col: 94, offset: 11097},
							val:        "**",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedBoldTextUnbalancedPunctuation",
			pos:  position{line: 261, col: 1, offset: 11205},
			expr: &actionExpr{
				pos: position{line: 261, col: 42, offset: 11246},
				run: (*parser).callonEscapedBoldTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 261, col: 42, offset: 11246},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 261, col: 42, offset: 11246},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 261, col: 55, offset: 11259},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 261, col: 55, offset: 11259},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 261, col: 59, offset: 11263},
										expr: &litMatcher{
											pos:        position{line: 261, col: 59, offset: 11263},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 261, col: 65, offset: 11269},
							val:        "**",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 261, col: 70, offset: 11274},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 261, col: 79, offset: 11283},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 261, col: 98, offset: 11302},
							val:        "*",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ItalicText",
			pos:  position{line: 266, col: 1, offset: 11495},
			expr: &choiceExpr{
				pos: position{line: 266, col: 15, offset: 11509},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 266, col: 15, offset: 11509},
						name: "ItalicTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 266, col: 45, offset: 11539},
						name: "ItalicTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 266, col: 79, offset: 11573},
						name: "ItalicTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "ItalicTextSimplePunctuation",
			pos:  position{line: 268, col: 1, offset: 11602},
			expr: &actionExpr{
				pos: position{line: 268, col: 32, offset: 11633},
				run: (*parser).callonItalicTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 268, col: 32, offset: 11633},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 268, col: 32, offset: 11633},
							expr: &litMatcher{
								pos:        position{line: 268, col: 33, offset: 11634},
								val:        "\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 268, col: 37, offset: 11638},
							val:        "_",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 268, col: 41, offset: 11642},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 268, col: 50, offset: 11651},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 268, col: 69, offset: 11670},
							val:        "_",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ItalicTextDoublePunctuation",
			pos:  position{line: 272, col: 1, offset: 11749},
			expr: &actionExpr{
				pos: position{line: 272, col: 32, offset: 11780},
				run: (*parser).callonItalicTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 272, col: 32, offset: 11780},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 272, col: 32, offset: 11780},
							expr: &litMatcher{
								pos:        position{line: 272, col: 33, offset: 11781},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 272, col: 38, offset: 11786},
							val:        "__",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 272, col: 43, offset: 11791},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 272, col: 52, offset: 11800},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 272, col: 71, offset: 11819},
							val:        "__",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ItalicTextUnbalancedPunctuation",
			pos:  position{line: 276, col: 1, offset: 11899},
			expr: &actionExpr{
				pos: position{line: 276, col: 36, offset: 11934},
				run: (*parser).callonItalicTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 276, col: 36, offset: 11934},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 276, col: 36, offset: 11934},
							expr: &litMatcher{
								pos:        position{line: 276, col: 37, offset: 11935},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 276, col: 42, offset: 11940},
							val:        "__",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 276, col: 47, offset: 11945},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 276, col: 56, offset: 11954},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 276, col: 75, offset: 11973},
							val:        "_",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedItalicText",
			pos:  position{line: 281, col: 1, offset: 12139},
			expr: &choiceExpr{
				pos: position{line: 281, col: 22, offset: 12160},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 281, col: 22, offset: 12160},
						name: "EscapedItalicTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 281, col: 59, offset: 12197},
						name: "EscapedItalicTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 281, col: 100, offset: 12238},
						name: "EscapedItalicTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "EscapedItalicTextSimplePunctuation",
			pos:  position{line: 283, col: 1, offset: 12320},
			expr: &actionExpr{
				pos: position{line: 283, col: 39, offset: 12358},
				run: (*parser).callonEscapedItalicTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 283, col: 39, offset: 12358},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 283, col: 39, offset: 12358},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 283, col: 52, offset: 12371},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 283, col: 52, offset: 12371},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 283, col: 56, offset: 12375},
										expr: &litMatcher{
											pos:        position{line: 283, col: 56, offset: 12375},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 283, col: 62, offset: 12381},
							val:        "_",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 283, col: 66, offset: 12385},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 283, col: 75, offset: 12394},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 283, col: 94, offset: 12413},
							val:        "_",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedItalicTextDoublePunctuation",
			pos:  position{line: 287, col: 1, offset: 12519},
			expr: &actionExpr{
				pos: position{line: 287, col: 39, offset: 12557},
				run: (*parser).callonEscapedItalicTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 287, col: 39, offset: 12557},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 287, col: 39, offset: 12557},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 287, col: 52, offset: 12570},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 287, col: 52, offset: 12570},
										val:        "\\\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 287, col: 57, offset: 12575},
										expr: &litMatcher{
											pos:        position{line: 287, col: 57, offset: 12575},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 287, col: 63, offset: 12581},
							val:        "__",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 287, col: 68, offset: 12586},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 287, col: 77, offset: 12595},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 287, col: 96, offset: 12614},
							val:        "__",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedItalicTextUnbalancedPunctuation",
			pos:  position{line: 291, col: 1, offset: 12722},
			expr: &actionExpr{
				pos: position{line: 291, col: 44, offset: 12765},
				run: (*parser).callonEscapedItalicTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 291, col: 44, offset: 12765},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 291, col: 44, offset: 12765},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 291, col: 57, offset: 12778},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 291, col: 57, offset: 12778},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 291, col: 61, offset: 12782},
										expr: &litMatcher{
											pos:        position{line: 291, col: 61, offset: 12782},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 291, col: 67, offset: 12788},
							val:        "__",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 291, col: 72, offset: 12793},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 291, col: 81, offset: 12802},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 291, col: 100, offset: 12821},
							val:        "_",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "MonospaceText",
			pos:  position{line: 296, col: 1, offset: 13014},
			expr: &choiceExpr{
				pos: position{line: 296, col: 18, offset: 13031},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 296, col: 18, offset: 13031},
						name: "MonospaceTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 296, col: 51, offset: 13064},
						name: "MonospaceTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 296, col: 88, offset: 13101},
						name: "MonospaceTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "MonospaceTextSimplePunctuation",
			pos:  position{line: 298, col: 1, offset: 13133},
			expr: &actionExpr{
				pos: position{line: 298, col: 35, offset: 13167},
				run: (*parser).callonMonospaceTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 298, col: 35, offset: 13167},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 298, col: 35, offset: 13167},
							expr: &litMatcher{
								pos:        position{line: 298, col: 36, offset: 13168},
								val:        "\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 298, col: 40, offset: 13172},
							val:        "`",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 298, col: 44, offset: 13176},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 298, col: 53, offset: 13185},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 298, col: 72, offset: 13204},
							val:        "`",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "MonospaceTextDoublePunctuation",
			pos:  position{line: 302, col: 1, offset: 13286},
			expr: &actionExpr{
				pos: position{line: 302, col: 35, offset: 13320},
				run: (*parser).callonMonospaceTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 302, col: 35, offset: 13320},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 302, col: 35, offset: 13320},
							expr: &litMatcher{
								pos:        position{line: 302, col: 36, offset: 13321},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 302, col: 41, offset: 13326},
							val:        "``",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 302, col: 46, offset: 13331},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 302, col: 55, offset: 13340},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 302, col: 74, offset: 13359},
							val:        "``",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "MonospaceTextUnbalancedPunctuation",
			pos:  position{line: 306, col: 1, offset: 13442},
			expr: &actionExpr{
				pos: position{line: 306, col: 39, offset: 13480},
				run: (*parser).callonMonospaceTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 306, col: 39, offset: 13480},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 306, col: 39, offset: 13480},
							expr: &litMatcher{
								pos:        position{line: 306, col: 40, offset: 13481},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 306, col: 45, offset: 13486},
							val:        "``",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 306, col: 50, offset: 13491},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 306, col: 59, offset: 13500},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 306, col: 78, offset: 13519},
							val:        "`",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedMonospaceText",
			pos:  position{line: 311, col: 1, offset: 13688},
			expr: &choiceExpr{
				pos: position{line: 311, col: 25, offset: 13712},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 311, col: 25, offset: 13712},
						name: "EscapedMonospaceTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 311, col: 65, offset: 13752},
						name: "EscapedMonospaceTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 311, col: 109, offset: 13796},
						name: "EscapedMonospaceTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "EscapedMonospaceTextSimplePunctuation",
			pos:  position{line: 313, col: 1, offset: 13881},
			expr: &actionExpr{
				pos: position{line: 313, col: 42, offset: 13922},
				run: (*parser).callonEscapedMonospaceTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 313, col: 42, offset: 13922},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 313, col: 42, offset: 13922},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 313, col: 55, offset: 13935},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 313, col: 55, offset: 13935},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 313, col: 59, offset: 13939},
										expr: &litMatcher{
											pos:        position{line: 313, col: 59, offset: 13939},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 313, col: 65, offset: 13945},
							val:        "`",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 313, col: 69, offset: 13949},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 313, col: 78, offset: 13958},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 313, col: 97, offset: 13977},
							val:        "`",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedMonospaceTextDoublePunctuation",
			pos:  position{line: 317, col: 1, offset: 14083},
			expr: &actionExpr{
				pos: position{line: 317, col: 42, offset: 14124},
				run: (*parser).callonEscapedMonospaceTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 317, col: 42, offset: 14124},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 317, col: 42, offset: 14124},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 317, col: 55, offset: 14137},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 317, col: 55, offset: 14137},
										val:        "\\\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 317, col: 60, offset: 14142},
										expr: &litMatcher{
											pos:        position{line: 317, col: 60, offset: 14142},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 317, col: 66, offset: 14148},
							val:        "``",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 317, col: 71, offset: 14153},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 317, col: 80, offset: 14162},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 317, col: 99, offset: 14181},
							val:        "``",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedMonospaceTextUnbalancedPunctuation",
			pos:  position{line: 321, col: 1, offset: 14289},
			expr: &actionExpr{
				pos: position{line: 321, col: 47, offset: 14335},
				run: (*parser).callonEscapedMonospaceTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 321, col: 47, offset: 14335},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 321, col: 47, offset: 14335},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 321, col: 60, offset: 14348},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 321, col: 60, offset: 14348},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 321, col: 64, offset: 14352},
										expr: &litMatcher{
											pos:        position{line: 321, col: 64, offset: 14352},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 321, col: 70, offset: 14358},
							val:        "``",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 321, col: 75, offset: 14363},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 321, col: 84, offset: 14372},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 321, col: 103, offset: 14391},
							val:        "`",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "QuotedTextContent",
			pos:  position{line: 326, col: 1, offset: 14584},
			expr: &seqExpr{
				pos: position{line: 326, col: 22, offset: 14605},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 326, col: 22, offset: 14605},
						name: "QuotedTextContentElement",
					},
					&zeroOrMoreExpr{
						pos: position{line: 326, col: 47, offset: 14630},
						expr: &seqExpr{
							pos: position{line: 326, col: 48, offset: 14631},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 326, col: 48, offset: 14631},
									expr: &ruleRefExpr{
										pos:  position{line: 326, col: 48, offset: 14631},
										name: "WS",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 326, col: 52, offset: 14635},
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
			pos:  position{line: 328, col: 1, offset: 14663},
			expr: &choiceExpr{
				pos: position{line: 328, col: 29, offset: 14691},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 328, col: 29, offset: 14691},
						name: "QuotedText",
					},
					&ruleRefExpr{
						pos:  position{line: 328, col: 42, offset: 14704},
						name: "QuotedTextCharacters",
					},
					&ruleRefExpr{
						pos:  position{line: 328, col: 65, offset: 14727},
						name: "CharactersWithQuotePunctuation",
					},
				},
			},
		},
		{
			name: "QuotedTextCharacters",
			pos:  position{line: 330, col: 1, offset: 14862},
			expr: &oneOrMoreExpr{
				pos: position{line: 330, col: 25, offset: 14886},
				expr: &seqExpr{
					pos: position{line: 330, col: 26, offset: 14887},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 330, col: 26, offset: 14887},
							expr: &ruleRefExpr{
								pos:  position{line: 330, col: 27, offset: 14888},
								name: "NEWLINE",
							},
						},
						&notExpr{
							pos: position{line: 330, col: 35, offset: 14896},
							expr: &ruleRefExpr{
								pos:  position{line: 330, col: 36, offset: 14897},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 330, col: 39, offset: 14900},
							expr: &litMatcher{
								pos:        position{line: 330, col: 40, offset: 14901},
								val:        "*",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 330, col: 44, offset: 14905},
							expr: &litMatcher{
								pos:        position{line: 330, col: 45, offset: 14906},
								val:        "_",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 330, col: 49, offset: 14910},
							expr: &litMatcher{
								pos:        position{line: 330, col: 50, offset: 14911},
								val:        "`",
								ignoreCase: false,
							},
						},
						&anyMatcher{
							line: 330, col: 54, offset: 14915,
						},
					},
				},
			},
		},
		{
			name: "CharactersWithQuotePunctuation",
			pos:  position{line: 331, col: 1, offset: 14957},
			expr: &actionExpr{
				pos: position{line: 331, col: 35, offset: 14991},
				run: (*parser).callonCharactersWithQuotePunctuation1,
				expr: &oneOrMoreExpr{
					pos: position{line: 331, col: 35, offset: 14991},
					expr: &seqExpr{
						pos: position{line: 331, col: 36, offset: 14992},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 331, col: 36, offset: 14992},
								expr: &ruleRefExpr{
									pos:  position{line: 331, col: 37, offset: 14993},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 331, col: 45, offset: 15001},
								expr: &ruleRefExpr{
									pos:  position{line: 331, col: 46, offset: 15002},
									name: "WS",
								},
							},
							&anyMatcher{
								line: 331, col: 50, offset: 15006,
							},
						},
					},
				},
			},
		},
		{
			name: "UnbalancedQuotePunctuation",
			pos:  position{line: 336, col: 1, offset: 15251},
			expr: &choiceExpr{
				pos: position{line: 336, col: 31, offset: 15281},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 336, col: 31, offset: 15281},
						val:        "*",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 336, col: 37, offset: 15287},
						val:        "_",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 336, col: 43, offset: 15293},
						val:        "`",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "Passthrough",
			pos:  position{line: 342, col: 1, offset: 15406},
			expr: &choiceExpr{
				pos: position{line: 342, col: 16, offset: 15421},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 342, col: 16, offset: 15421},
						name: "TriplePlusPassthrough",
					},
					&ruleRefExpr{
						pos:  position{line: 342, col: 40, offset: 15445},
						name: "SinglePlusPassthrough",
					},
					&ruleRefExpr{
						pos:  position{line: 342, col: 64, offset: 15469},
						name: "PassthroughMacro",
					},
				},
			},
		},
		{
			name: "SinglePlusPassthrough",
			pos:  position{line: 344, col: 1, offset: 15487},
			expr: &actionExpr{
				pos: position{line: 344, col: 26, offset: 15512},
				run: (*parser).callonSinglePlusPassthrough1,
				expr: &seqExpr{
					pos: position{line: 344, col: 26, offset: 15512},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 344, col: 26, offset: 15512},
							val:        "+",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 344, col: 30, offset: 15516},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 344, col: 38, offset: 15524},
								expr: &seqExpr{
									pos: position{line: 344, col: 39, offset: 15525},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 344, col: 39, offset: 15525},
											expr: &ruleRefExpr{
												pos:  position{line: 344, col: 40, offset: 15526},
												name: "NEWLINE",
											},
										},
										&notExpr{
											pos: position{line: 344, col: 48, offset: 15534},
											expr: &litMatcher{
												pos:        position{line: 344, col: 49, offset: 15535},
												val:        "+",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 344, col: 53, offset: 15539,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 344, col: 57, offset: 15543},
							val:        "+",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TriplePlusPassthrough",
			pos:  position{line: 348, col: 1, offset: 15638},
			expr: &actionExpr{
				pos: position{line: 348, col: 26, offset: 15663},
				run: (*parser).callonTriplePlusPassthrough1,
				expr: &seqExpr{
					pos: position{line: 348, col: 26, offset: 15663},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 348, col: 26, offset: 15663},
							val:        "+++",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 348, col: 32, offset: 15669},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 348, col: 40, offset: 15677},
								expr: &seqExpr{
									pos: position{line: 348, col: 41, offset: 15678},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 348, col: 41, offset: 15678},
											expr: &litMatcher{
												pos:        position{line: 348, col: 42, offset: 15679},
												val:        "+++",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 348, col: 48, offset: 15685,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 348, col: 52, offset: 15689},
							val:        "+++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "PassthroughMacro",
			pos:  position{line: 352, col: 1, offset: 15786},
			expr: &choiceExpr{
				pos: position{line: 352, col: 21, offset: 15806},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 352, col: 21, offset: 15806},
						name: "SimplePassthroughMacro",
					},
					&ruleRefExpr{
						pos:  position{line: 352, col: 46, offset: 15831},
						name: "PassthroughWithQuotedText",
					},
				},
			},
		},
		{
			name: "SimplePassthroughMacro",
			pos:  position{line: 354, col: 1, offset: 15858},
			expr: &actionExpr{
				pos: position{line: 354, col: 27, offset: 15884},
				run: (*parser).callonSimplePassthroughMacro1,
				expr: &seqExpr{
					pos: position{line: 354, col: 27, offset: 15884},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 354, col: 27, offset: 15884},
							val:        "pass:[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 354, col: 36, offset: 15893},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 354, col: 44, offset: 15901},
								expr: &ruleRefExpr{
									pos:  position{line: 354, col: 45, offset: 15902},
									name: "PassthroughMacroCharacter",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 354, col: 73, offset: 15930},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "PassthroughWithQuotedText",
			pos:  position{line: 358, col: 1, offset: 16020},
			expr: &actionExpr{
				pos: position{line: 358, col: 30, offset: 16049},
				run: (*parser).callonPassthroughWithQuotedText1,
				expr: &seqExpr{
					pos: position{line: 358, col: 30, offset: 16049},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 358, col: 30, offset: 16049},
							val:        "pass:q[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 358, col: 40, offset: 16059},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 358, col: 48, offset: 16067},
								expr: &choiceExpr{
									pos: position{line: 358, col: 49, offset: 16068},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 358, col: 49, offset: 16068},
											name: "QuotedText",
										},
										&ruleRefExpr{
											pos:  position{line: 358, col: 62, offset: 16081},
											name: "PassthroughMacroCharacter",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 358, col: 90, offset: 16109},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "PassthroughMacroCharacter",
			pos:  position{line: 362, col: 1, offset: 16199},
			expr: &seqExpr{
				pos: position{line: 362, col: 31, offset: 16229},
				exprs: []interface{}{
					&notExpr{
						pos: position{line: 362, col: 31, offset: 16229},
						expr: &litMatcher{
							pos:        position{line: 362, col: 32, offset: 16230},
							val:        "]",
							ignoreCase: false,
						},
					},
					&anyMatcher{
						line: 362, col: 36, offset: 16234,
					},
				},
			},
		},
		{
			name: "ExternalLink",
			pos:  position{line: 367, col: 1, offset: 16339},
			expr: &actionExpr{
				pos: position{line: 367, col: 17, offset: 16355},
				run: (*parser).callonExternalLink1,
				expr: &seqExpr{
					pos: position{line: 367, col: 17, offset: 16355},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 367, col: 17, offset: 16355},
							label: "url",
							expr: &seqExpr{
								pos: position{line: 367, col: 22, offset: 16360},
								exprs: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 367, col: 22, offset: 16360},
										name: "URL_SCHEME",
									},
									&ruleRefExpr{
										pos:  position{line: 367, col: 33, offset: 16371},
										name: "URL",
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 367, col: 38, offset: 16376},
							label: "text",
							expr: &zeroOrOneExpr{
								pos: position{line: 367, col: 43, offset: 16381},
								expr: &seqExpr{
									pos: position{line: 367, col: 44, offset: 16382},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 367, col: 44, offset: 16382},
											val:        "[",
											ignoreCase: false,
										},
										&zeroOrMoreExpr{
											pos: position{line: 367, col: 48, offset: 16386},
											expr: &ruleRefExpr{
												pos:  position{line: 367, col: 49, offset: 16387},
												name: "URL_TEXT",
											},
										},
										&litMatcher{
											pos:        position{line: 367, col: 60, offset: 16398},
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
			pos:  position{line: 377, col: 1, offset: 16677},
			expr: &actionExpr{
				pos: position{line: 377, col: 15, offset: 16691},
				run: (*parser).callonBlockImage1,
				expr: &seqExpr{
					pos: position{line: 377, col: 15, offset: 16691},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 377, col: 15, offset: 16691},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 377, col: 26, offset: 16702},
								expr: &ruleRefExpr{
									pos:  position{line: 377, col: 27, offset: 16703},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 377, col: 46, offset: 16722},
							label: "image",
							expr: &ruleRefExpr{
								pos:  position{line: 377, col: 52, offset: 16728},
								name: "BlockImageMacro",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 377, col: 69, offset: 16745},
							expr: &ruleRefExpr{
								pos:  position{line: 377, col: 69, offset: 16745},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 377, col: 73, offset: 16749},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "BlockImageMacro",
			pos:  position{line: 382, col: 1, offset: 16910},
			expr: &actionExpr{
				pos: position{line: 382, col: 20, offset: 16929},
				run: (*parser).callonBlockImageMacro1,
				expr: &seqExpr{
					pos: position{line: 382, col: 20, offset: 16929},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 382, col: 20, offset: 16929},
							val:        "image::",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 382, col: 30, offset: 16939},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 382, col: 36, offset: 16945},
								name: "URL",
							},
						},
						&litMatcher{
							pos:        position{line: 382, col: 41, offset: 16950},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 382, col: 45, offset: 16954},
							label: "attributes",
							expr: &zeroOrOneExpr{
								pos: position{line: 382, col: 57, offset: 16966},
								expr: &ruleRefExpr{
									pos:  position{line: 382, col: 57, offset: 16966},
									name: "URL_TEXT",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 382, col: 68, offset: 16977},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "InlineImage",
			pos:  position{line: 386, col: 1, offset: 17044},
			expr: &actionExpr{
				pos: position{line: 386, col: 16, offset: 17059},
				run: (*parser).callonInlineImage1,
				expr: &labeledExpr{
					pos:   position{line: 386, col: 16, offset: 17059},
					label: "image",
					expr: &ruleRefExpr{
						pos:  position{line: 386, col: 22, offset: 17065},
						name: "InlineImageMacro",
					},
				},
			},
		},
		{
			name: "InlineImageMacro",
			pos:  position{line: 391, col: 1, offset: 17212},
			expr: &actionExpr{
				pos: position{line: 391, col: 21, offset: 17232},
				run: (*parser).callonInlineImageMacro1,
				expr: &seqExpr{
					pos: position{line: 391, col: 21, offset: 17232},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 391, col: 21, offset: 17232},
							val:        "image:",
							ignoreCase: false,
						},
						&notExpr{
							pos: position{line: 391, col: 30, offset: 17241},
							expr: &litMatcher{
								pos:        position{line: 391, col: 31, offset: 17242},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 391, col: 35, offset: 17246},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 391, col: 41, offset: 17252},
								name: "URL",
							},
						},
						&litMatcher{
							pos:        position{line: 391, col: 46, offset: 17257},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 391, col: 50, offset: 17261},
							label: "attributes",
							expr: &zeroOrOneExpr{
								pos: position{line: 391, col: 62, offset: 17273},
								expr: &ruleRefExpr{
									pos:  position{line: 391, col: 62, offset: 17273},
									name: "URL_TEXT",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 391, col: 73, offset: 17284},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "DelimitedBlock",
			pos:  position{line: 398, col: 1, offset: 17614},
			expr: &choiceExpr{
				pos: position{line: 398, col: 19, offset: 17632},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 398, col: 19, offset: 17632},
						name: "FencedBlock",
					},
					&ruleRefExpr{
						pos:  position{line: 398, col: 33, offset: 17646},
						name: "ListingBlock",
					},
				},
			},
		},
		{
			name: "BlockDelimiter",
			pos:  position{line: 400, col: 1, offset: 17661},
			expr: &choiceExpr{
				pos: position{line: 400, col: 19, offset: 17679},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 400, col: 19, offset: 17679},
						name: "FencedBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 400, col: 42, offset: 17702},
						name: "ListingBlockDelimiter",
					},
				},
			},
		},
		{
			name: "FencedBlockDelimiter",
			pos:  position{line: 402, col: 1, offset: 17725},
			expr: &litMatcher{
				pos:        position{line: 402, col: 25, offset: 17749},
				val:        "```",
				ignoreCase: false,
			},
		},
		{
			name: "FencedBlock",
			pos:  position{line: 404, col: 1, offset: 17756},
			expr: &actionExpr{
				pos: position{line: 404, col: 16, offset: 17771},
				run: (*parser).callonFencedBlock1,
				expr: &seqExpr{
					pos: position{line: 404, col: 16, offset: 17771},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 404, col: 16, offset: 17771},
							name: "FencedBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 404, col: 37, offset: 17792},
							expr: &ruleRefExpr{
								pos:  position{line: 404, col: 37, offset: 17792},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 404, col: 41, offset: 17796},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 404, col: 49, offset: 17804},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 404, col: 58, offset: 17813},
								name: "FencedBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 404, col: 78, offset: 17833},
							name: "FencedBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 404, col: 99, offset: 17854},
							expr: &ruleRefExpr{
								pos:  position{line: 404, col: 99, offset: 17854},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 404, col: 103, offset: 17858},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "FencedBlockContent",
			pos:  position{line: 408, col: 1, offset: 17946},
			expr: &labeledExpr{
				pos:   position{line: 408, col: 23, offset: 17968},
				label: "content",
				expr: &zeroOrMoreExpr{
					pos: position{line: 408, col: 31, offset: 17976},
					expr: &seqExpr{
						pos: position{line: 408, col: 32, offset: 17977},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 408, col: 32, offset: 17977},
								expr: &ruleRefExpr{
									pos:  position{line: 408, col: 33, offset: 17978},
									name: "FencedBlockDelimiter",
								},
							},
							&anyMatcher{
								line: 408, col: 54, offset: 17999,
							},
						},
					},
				},
			},
		},
		{
			name: "ListingBlockDelimiter",
			pos:  position{line: 410, col: 1, offset: 18005},
			expr: &litMatcher{
				pos:        position{line: 410, col: 26, offset: 18030},
				val:        "----",
				ignoreCase: false,
			},
		},
		{
			name: "ListingBlock",
			pos:  position{line: 412, col: 1, offset: 18038},
			expr: &actionExpr{
				pos: position{line: 412, col: 17, offset: 18054},
				run: (*parser).callonListingBlock1,
				expr: &seqExpr{
					pos: position{line: 412, col: 17, offset: 18054},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 412, col: 17, offset: 18054},
							name: "ListingBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 412, col: 39, offset: 18076},
							expr: &ruleRefExpr{
								pos:  position{line: 412, col: 39, offset: 18076},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 412, col: 43, offset: 18080},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 412, col: 51, offset: 18088},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 412, col: 60, offset: 18097},
								name: "ListingBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 412, col: 81, offset: 18118},
							name: "ListingBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 412, col: 103, offset: 18140},
							expr: &ruleRefExpr{
								pos:  position{line: 412, col: 103, offset: 18140},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 412, col: 107, offset: 18144},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ListingBlockContent",
			pos:  position{line: 416, col: 1, offset: 18233},
			expr: &labeledExpr{
				pos:   position{line: 416, col: 24, offset: 18256},
				label: "content",
				expr: &zeroOrMoreExpr{
					pos: position{line: 416, col: 32, offset: 18264},
					expr: &seqExpr{
						pos: position{line: 416, col: 33, offset: 18265},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 416, col: 33, offset: 18265},
								expr: &ruleRefExpr{
									pos:  position{line: 416, col: 34, offset: 18266},
									name: "ListingBlockDelimiter",
								},
							},
							&anyMatcher{
								line: 416, col: 56, offset: 18288,
							},
						},
					},
				},
			},
		},
		{
			name: "LiteralBlock",
			pos:  position{line: 421, col: 1, offset: 18561},
			expr: &choiceExpr{
				pos: position{line: 421, col: 17, offset: 18577},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 421, col: 17, offset: 18577},
						name: "ParagraphWithSpaces",
					},
					&ruleRefExpr{
						pos:  position{line: 421, col: 39, offset: 18599},
						name: "ParagraphWithLiteralBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 421, col: 76, offset: 18636},
						name: "ParagraphWithLiteralAttribute",
					},
				},
			},
		},
		{
			name: "ParagraphWithSpaces",
			pos:  position{line: 424, col: 1, offset: 18731},
			expr: &actionExpr{
				pos: position{line: 424, col: 24, offset: 18754},
				run: (*parser).callonParagraphWithSpaces1,
				expr: &seqExpr{
					pos: position{line: 424, col: 24, offset: 18754},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 424, col: 24, offset: 18754},
							label: "spaces",
							expr: &oneOrMoreExpr{
								pos: position{line: 424, col: 32, offset: 18762},
								expr: &ruleRefExpr{
									pos:  position{line: 424, col: 32, offset: 18762},
									name: "WS",
								},
							},
						},
						&notExpr{
							pos: position{line: 424, col: 37, offset: 18767},
							expr: &ruleRefExpr{
								pos:  position{line: 424, col: 38, offset: 18768},
								name: "NEWLINE",
							},
						},
						&labeledExpr{
							pos:   position{line: 424, col: 46, offset: 18776},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 424, col: 55, offset: 18785},
								name: "LiteralBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 424, col: 76, offset: 18806},
							name: "EndOfLiteralBlock",
						},
					},
				},
			},
		},
		{
			name: "LiteralBlockContent",
			pos:  position{line: 429, col: 1, offset: 18987},
			expr: &actionExpr{
				pos: position{line: 429, col: 24, offset: 19010},
				run: (*parser).callonLiteralBlockContent1,
				expr: &labeledExpr{
					pos:   position{line: 429, col: 24, offset: 19010},
					label: "content",
					expr: &oneOrMoreExpr{
						pos: position{line: 429, col: 32, offset: 19018},
						expr: &seqExpr{
							pos: position{line: 429, col: 33, offset: 19019},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 429, col: 33, offset: 19019},
									expr: &seqExpr{
										pos: position{line: 429, col: 35, offset: 19021},
										exprs: []interface{}{
											&ruleRefExpr{
												pos:  position{line: 429, col: 35, offset: 19021},
												name: "NEWLINE",
											},
											&ruleRefExpr{
												pos:  position{line: 429, col: 43, offset: 19029},
												name: "BlankLine",
											},
										},
									},
								},
								&anyMatcher{
									line: 429, col: 54, offset: 19040,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "EndOfLiteralBlock",
			pos:  position{line: 434, col: 1, offset: 19125},
			expr: &choiceExpr{
				pos: position{line: 434, col: 22, offset: 19146},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 434, col: 22, offset: 19146},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 434, col: 22, offset: 19146},
								name: "NEWLINE",
							},
							&ruleRefExpr{
								pos:  position{line: 434, col: 30, offset: 19154},
								name: "BlankLine",
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 434, col: 42, offset: 19166},
						name: "NEWLINE",
					},
					&ruleRefExpr{
						pos:  position{line: 434, col: 52, offset: 19176},
						name: "EOF",
					},
				},
			},
		},
		{
			name: "ParagraphWithLiteralBlockDelimiter",
			pos:  position{line: 437, col: 1, offset: 19236},
			expr: &actionExpr{
				pos: position{line: 437, col: 39, offset: 19274},
				run: (*parser).callonParagraphWithLiteralBlockDelimiter1,
				expr: &seqExpr{
					pos: position{line: 437, col: 39, offset: 19274},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 437, col: 39, offset: 19274},
							name: "LiteralBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 437, col: 61, offset: 19296},
							expr: &ruleRefExpr{
								pos:  position{line: 437, col: 61, offset: 19296},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 437, col: 65, offset: 19300},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 437, col: 73, offset: 19308},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 437, col: 81, offset: 19316},
								expr: &seqExpr{
									pos: position{line: 437, col: 82, offset: 19317},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 437, col: 82, offset: 19317},
											expr: &ruleRefExpr{
												pos:  position{line: 437, col: 83, offset: 19318},
												name: "LiteralBlockDelimiter",
											},
										},
										&anyMatcher{
											line: 437, col: 105, offset: 19340,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 437, col: 109, offset: 19344},
							name: "LiteralBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 437, col: 131, offset: 19366},
							expr: &ruleRefExpr{
								pos:  position{line: 437, col: 131, offset: 19366},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 437, col: 135, offset: 19370},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "LiteralBlockDelimiter",
			pos:  position{line: 441, col: 1, offset: 19454},
			expr: &litMatcher{
				pos:        position{line: 441, col: 26, offset: 19479},
				val:        "....",
				ignoreCase: false,
			},
		},
		{
			name: "ParagraphWithLiteralAttribute",
			pos:  position{line: 444, col: 1, offset: 19541},
			expr: &actionExpr{
				pos: position{line: 444, col: 34, offset: 19574},
				run: (*parser).callonParagraphWithLiteralAttribute1,
				expr: &seqExpr{
					pos: position{line: 444, col: 34, offset: 19574},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 444, col: 34, offset: 19574},
							val:        "[literal]",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 444, col: 46, offset: 19586},
							expr: &ruleRefExpr{
								pos:  position{line: 444, col: 46, offset: 19586},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 444, col: 50, offset: 19590},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 444, col: 58, offset: 19598},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 444, col: 67, offset: 19607},
								name: "LiteralBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 444, col: 88, offset: 19628},
							name: "EndOfLiteralBlock",
						},
					},
				},
			},
		},
		{
			name: "ElementAttribute",
			pos:  position{line: 451, col: 1, offset: 19840},
			expr: &labeledExpr{
				pos:   position{line: 451, col: 21, offset: 19860},
				label: "meta",
				expr: &choiceExpr{
					pos: position{line: 451, col: 27, offset: 19866},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 451, col: 27, offset: 19866},
							name: "ElementLink",
						},
						&ruleRefExpr{
							pos:  position{line: 451, col: 41, offset: 19880},
							name: "ElementID",
						},
						&ruleRefExpr{
							pos:  position{line: 451, col: 53, offset: 19892},
							name: "ElementTitle",
						},
					},
				},
			},
		},
		{
			name: "ElementLink",
			pos:  position{line: 454, col: 1, offset: 19963},
			expr: &actionExpr{
				pos: position{line: 454, col: 16, offset: 19978},
				run: (*parser).callonElementLink1,
				expr: &seqExpr{
					pos: position{line: 454, col: 16, offset: 19978},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 454, col: 16, offset: 19978},
							val:        "[",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 454, col: 20, offset: 19982},
							expr: &ruleRefExpr{
								pos:  position{line: 454, col: 20, offset: 19982},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 454, col: 24, offset: 19986},
							val:        "link",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 454, col: 31, offset: 19993},
							expr: &ruleRefExpr{
								pos:  position{line: 454, col: 31, offset: 19993},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 454, col: 35, offset: 19997},
							val:        "=",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 454, col: 39, offset: 20001},
							expr: &ruleRefExpr{
								pos:  position{line: 454, col: 39, offset: 20001},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 454, col: 43, offset: 20005},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 454, col: 48, offset: 20010},
								name: "URL",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 454, col: 52, offset: 20014},
							expr: &ruleRefExpr{
								pos:  position{line: 454, col: 52, offset: 20014},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 454, col: 56, offset: 20018},
							val:        "]",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 454, col: 60, offset: 20022},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ElementID",
			pos:  position{line: 459, col: 1, offset: 20132},
			expr: &actionExpr{
				pos: position{line: 459, col: 14, offset: 20145},
				run: (*parser).callonElementID1,
				expr: &seqExpr{
					pos: position{line: 459, col: 14, offset: 20145},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 459, col: 14, offset: 20145},
							val:        "[",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 459, col: 18, offset: 20149},
							expr: &ruleRefExpr{
								pos:  position{line: 459, col: 18, offset: 20149},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 459, col: 22, offset: 20153},
							val:        "#",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 459, col: 26, offset: 20157},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 459, col: 30, offset: 20161},
								name: "ID",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 459, col: 34, offset: 20165},
							expr: &ruleRefExpr{
								pos:  position{line: 459, col: 34, offset: 20165},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 459, col: 38, offset: 20169},
							val:        "]",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 459, col: 42, offset: 20173},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ElementTitle",
			pos:  position{line: 465, col: 1, offset: 20367},
			expr: &actionExpr{
				pos: position{line: 465, col: 17, offset: 20383},
				run: (*parser).callonElementTitle1,
				expr: &seqExpr{
					pos: position{line: 465, col: 17, offset: 20383},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 465, col: 17, offset: 20383},
							val:        ".",
							ignoreCase: false,
						},
						&notExpr{
							pos: position{line: 465, col: 21, offset: 20387},
							expr: &litMatcher{
								pos:        position{line: 465, col: 22, offset: 20388},
								val:        ".",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 465, col: 26, offset: 20392},
							expr: &ruleRefExpr{
								pos:  position{line: 465, col: 27, offset: 20393},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 465, col: 30, offset: 20396},
							label: "title",
							expr: &oneOrMoreExpr{
								pos: position{line: 465, col: 36, offset: 20402},
								expr: &seqExpr{
									pos: position{line: 465, col: 37, offset: 20403},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 465, col: 37, offset: 20403},
											expr: &ruleRefExpr{
												pos:  position{line: 465, col: 38, offset: 20404},
												name: "NEWLINE",
											},
										},
										&anyMatcher{
											line: 465, col: 46, offset: 20412,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 465, col: 50, offset: 20416},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "BlankLine",
			pos:  position{line: 472, col: 1, offset: 20586},
			expr: &actionExpr{
				pos: position{line: 472, col: 14, offset: 20599},
				run: (*parser).callonBlankLine1,
				expr: &seqExpr{
					pos: position{line: 472, col: 14, offset: 20599},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 472, col: 14, offset: 20599},
							expr: &ruleRefExpr{
								pos:  position{line: 472, col: 15, offset: 20600},
								name: "EOF",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 472, col: 19, offset: 20604},
							expr: &ruleRefExpr{
								pos:  position{line: 472, col: 19, offset: 20604},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 472, col: 23, offset: 20608},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "Characters",
			pos:  position{line: 479, col: 1, offset: 20755},
			expr: &actionExpr{
				pos: position{line: 479, col: 15, offset: 20769},
				run: (*parser).callonCharacters1,
				expr: &oneOrMoreExpr{
					pos: position{line: 479, col: 15, offset: 20769},
					expr: &seqExpr{
						pos: position{line: 479, col: 16, offset: 20770},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 479, col: 16, offset: 20770},
								expr: &ruleRefExpr{
									pos:  position{line: 479, col: 17, offset: 20771},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 479, col: 25, offset: 20779},
								expr: &ruleRefExpr{
									pos:  position{line: 479, col: 26, offset: 20780},
									name: "WS",
								},
							},
							&anyMatcher{
								line: 479, col: 29, offset: 20783,
							},
						},
					},
				},
			},
		},
		{
			name: "URL",
			pos:  position{line: 484, col: 1, offset: 20824},
			expr: &actionExpr{
				pos: position{line: 484, col: 8, offset: 20831},
				run: (*parser).callonURL1,
				expr: &oneOrMoreExpr{
					pos: position{line: 484, col: 8, offset: 20831},
					expr: &seqExpr{
						pos: position{line: 484, col: 9, offset: 20832},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 484, col: 9, offset: 20832},
								expr: &ruleRefExpr{
									pos:  position{line: 484, col: 10, offset: 20833},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 484, col: 18, offset: 20841},
								expr: &ruleRefExpr{
									pos:  position{line: 484, col: 19, offset: 20842},
									name: "WS",
								},
							},
							&notExpr{
								pos: position{line: 484, col: 22, offset: 20845},
								expr: &litMatcher{
									pos:        position{line: 484, col: 23, offset: 20846},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 484, col: 27, offset: 20850},
								expr: &litMatcher{
									pos:        position{line: 484, col: 28, offset: 20851},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 484, col: 32, offset: 20855,
							},
						},
					},
				},
			},
		},
		{
			name: "ID",
			pos:  position{line: 488, col: 1, offset: 20895},
			expr: &actionExpr{
				pos: position{line: 488, col: 7, offset: 20901},
				run: (*parser).callonID1,
				expr: &oneOrMoreExpr{
					pos: position{line: 488, col: 7, offset: 20901},
					expr: &seqExpr{
						pos: position{line: 488, col: 8, offset: 20902},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 488, col: 8, offset: 20902},
								expr: &ruleRefExpr{
									pos:  position{line: 488, col: 9, offset: 20903},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 488, col: 17, offset: 20911},
								expr: &ruleRefExpr{
									pos:  position{line: 488, col: 18, offset: 20912},
									name: "WS",
								},
							},
							&notExpr{
								pos: position{line: 488, col: 21, offset: 20915},
								expr: &litMatcher{
									pos:        position{line: 488, col: 22, offset: 20916},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 488, col: 26, offset: 20920},
								expr: &litMatcher{
									pos:        position{line: 488, col: 27, offset: 20921},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 488, col: 31, offset: 20925,
							},
						},
					},
				},
			},
		},
		{
			name: "URL_TEXT",
			pos:  position{line: 492, col: 1, offset: 20965},
			expr: &actionExpr{
				pos: position{line: 492, col: 13, offset: 20977},
				run: (*parser).callonURL_TEXT1,
				expr: &oneOrMoreExpr{
					pos: position{line: 492, col: 13, offset: 20977},
					expr: &seqExpr{
						pos: position{line: 492, col: 14, offset: 20978},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 492, col: 14, offset: 20978},
								expr: &ruleRefExpr{
									pos:  position{line: 492, col: 15, offset: 20979},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 492, col: 23, offset: 20987},
								expr: &litMatcher{
									pos:        position{line: 492, col: 24, offset: 20988},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 492, col: 28, offset: 20992},
								expr: &litMatcher{
									pos:        position{line: 492, col: 29, offset: 20993},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 492, col: 33, offset: 20997,
							},
						},
					},
				},
			},
		},
		{
			name: "URL_SCHEME",
			pos:  position{line: 496, col: 1, offset: 21037},
			expr: &choiceExpr{
				pos: position{line: 496, col: 15, offset: 21051},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 496, col: 15, offset: 21051},
						val:        "http://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 496, col: 27, offset: 21063},
						val:        "https://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 496, col: 40, offset: 21076},
						val:        "ftp://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 496, col: 51, offset: 21087},
						val:        "irc://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 496, col: 62, offset: 21098},
						val:        "mailto:",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "DIGIT",
			pos:  position{line: 498, col: 1, offset: 21109},
			expr: &charClassMatcher{
				pos:        position{line: 498, col: 13, offset: 21121},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NEWLINE",
			pos:  position{line: 500, col: 1, offset: 21128},
			expr: &choiceExpr{
				pos: position{line: 500, col: 13, offset: 21140},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 500, col: 13, offset: 21140},
						val:        "\r\n",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 500, col: 22, offset: 21149},
						val:        "\r",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 500, col: 29, offset: 21156},
						val:        "\n",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "WS",
			pos:  position{line: 502, col: 1, offset: 21162},
			expr: &choiceExpr{
				pos: position{line: 502, col: 13, offset: 21174},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 502, col: 13, offset: 21174},
						val:        " ",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 502, col: 19, offset: 21180},
						run: (*parser).callonWS3,
						expr: &litMatcher{
							pos:        position{line: 502, col: 19, offset: 21180},
							val:        "\t",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 506, col: 1, offset: 21225},
			expr: &notExpr{
				pos: position{line: 506, col: 13, offset: 21237},
				expr: &anyMatcher{
					line: 506, col: 14, offset: 21238,
				},
			},
		},
		{
			name: "EOL",
			pos:  position{line: 508, col: 1, offset: 21241},
			expr: &choiceExpr{
				pos: position{line: 508, col: 13, offset: 21253},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 508, col: 13, offset: 21253},
						name: "NEWLINE",
					},
					&ruleRefExpr{
						pos:  position{line: 508, col: 23, offset: 21263},
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

func (c *current) onListingBlock1(content interface{}) (interface{}, error) {
	return types.NewDelimitedBlock(types.ListingBlock, content.([]interface{}))
}

func (p *parser) callonListingBlock1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onListingBlock1(stack["content"])
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
