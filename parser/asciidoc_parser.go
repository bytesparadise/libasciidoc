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
			pos:  position{line: 19, col: 1, offset: 501},
			expr: &actionExpr{
				pos: position{line: 19, col: 13, offset: 513},
				run: (*parser).callonDocument1,
				expr: &seqExpr{
					pos: position{line: 19, col: 13, offset: 513},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 19, col: 13, offset: 513},
							label: "frontMatter",
							expr: &zeroOrOneExpr{
								pos: position{line: 19, col: 26, offset: 526},
								expr: &ruleRefExpr{
									pos:  position{line: 19, col: 26, offset: 526},
									name: "FrontMatter",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 19, col: 40, offset: 540},
							label: "documentHeader",
							expr: &zeroOrOneExpr{
								pos: position{line: 19, col: 56, offset: 556},
								expr: &ruleRefExpr{
									pos:  position{line: 19, col: 56, offset: 556},
									name: "DocumentHeader",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 19, col: 73, offset: 573},
							label: "blocks",
							expr: &ruleRefExpr{
								pos:  position{line: 19, col: 81, offset: 581},
								name: "DocumentBlocks",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 19, col: 97, offset: 597},
							name: "EOF",
						},
					},
				},
			},
		},
		{
			name: "DocumentBlocks",
			pos:  position{line: 23, col: 1, offset: 685},
			expr: &choiceExpr{
				pos: position{line: 23, col: 19, offset: 703},
				alternatives: []interface{}{
					&labeledExpr{
						pos:   position{line: 23, col: 19, offset: 703},
						label: "content",
						expr: &seqExpr{
							pos: position{line: 23, col: 28, offset: 712},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 23, col: 28, offset: 712},
									name: "Preamble",
								},
								&oneOrMoreExpr{
									pos: position{line: 23, col: 37, offset: 721},
									expr: &ruleRefExpr{
										pos:  position{line: 23, col: 37, offset: 721},
										name: "Section",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 23, col: 49, offset: 733},
						run: (*parser).callonDocumentBlocks7,
						expr: &labeledExpr{
							pos:   position{line: 23, col: 49, offset: 733},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 23, col: 58, offset: 742},
								expr: &ruleRefExpr{
									pos:  position{line: 23, col: 58, offset: 742},
									name: "BlockElement",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "BlockElement",
			pos:  position{line: 27, col: 1, offset: 786},
			expr: &choiceExpr{
				pos: position{line: 27, col: 17, offset: 802},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 27, col: 17, offset: 802},
						name: "DocumentAttributeDeclaration",
					},
					&ruleRefExpr{
						pos:  position{line: 27, col: 48, offset: 833},
						name: "DocumentAttributeReset",
					},
					&ruleRefExpr{
						pos:  position{line: 27, col: 73, offset: 858},
						name: "TableOfContentsMacro",
					},
					&ruleRefExpr{
						pos:  position{line: 27, col: 96, offset: 881},
						name: "BlockImage",
					},
					&ruleRefExpr{
						pos:  position{line: 27, col: 109, offset: 894},
						name: "List",
					},
					&ruleRefExpr{
						pos:  position{line: 27, col: 116, offset: 901},
						name: "LiteralBlock",
					},
					&ruleRefExpr{
						pos:  position{line: 27, col: 131, offset: 916},
						name: "DelimitedBlock",
					},
					&ruleRefExpr{
						pos:  position{line: 27, col: 148, offset: 933},
						name: "Paragraph",
					},
					&seqExpr{
						pos: position{line: 27, col: 161, offset: 946},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 27, col: 161, offset: 946},
								name: "ElementAttribute",
							},
							&ruleRefExpr{
								pos:  position{line: 27, col: 178, offset: 963},
								name: "EOL",
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 27, col: 185, offset: 970},
						name: "BlankLine",
					},
				},
			},
		},
		{
			name: "Preamble",
			pos:  position{line: 29, col: 1, offset: 1025},
			expr: &actionExpr{
				pos: position{line: 29, col: 13, offset: 1037},
				run: (*parser).callonPreamble1,
				expr: &labeledExpr{
					pos:   position{line: 29, col: 13, offset: 1037},
					label: "elements",
					expr: &zeroOrMoreExpr{
						pos: position{line: 29, col: 23, offset: 1047},
						expr: &ruleRefExpr{
							pos:  position{line: 29, col: 23, offset: 1047},
							name: "BlockElement",
						},
					},
				},
			},
		},
		{
			name: "FrontMatter",
			pos:  position{line: 36, col: 1, offset: 1230},
			expr: &ruleRefExpr{
				pos:  position{line: 36, col: 16, offset: 1245},
				name: "YamlFrontMatter",
			},
		},
		{
			name: "FrontMatter",
			pos:  position{line: 38, col: 1, offset: 1263},
			expr: &actionExpr{
				pos: position{line: 38, col: 16, offset: 1278},
				run: (*parser).callonFrontMatter1,
				expr: &seqExpr{
					pos: position{line: 38, col: 16, offset: 1278},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 38, col: 16, offset: 1278},
							name: "YamlFrontMatterToken",
						},
						&labeledExpr{
							pos:   position{line: 38, col: 37, offset: 1299},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 38, col: 46, offset: 1308},
								name: "YamlFrontMatterContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 38, col: 70, offset: 1332},
							name: "YamlFrontMatterToken",
						},
					},
				},
			},
		},
		{
			name: "YamlFrontMatterToken",
			pos:  position{line: 42, col: 1, offset: 1412},
			expr: &seqExpr{
				pos: position{line: 42, col: 26, offset: 1437},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 42, col: 26, offset: 1437},
						val:        "---",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 42, col: 32, offset: 1443},
						name: "EOL",
					},
				},
			},
		},
		{
			name: "YamlFrontMatterContent",
			pos:  position{line: 44, col: 1, offset: 1448},
			expr: &actionExpr{
				pos: position{line: 44, col: 27, offset: 1474},
				run: (*parser).callonYamlFrontMatterContent1,
				expr: &zeroOrMoreExpr{
					pos: position{line: 44, col: 27, offset: 1474},
					expr: &seqExpr{
						pos: position{line: 44, col: 28, offset: 1475},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 44, col: 28, offset: 1475},
								expr: &ruleRefExpr{
									pos:  position{line: 44, col: 29, offset: 1476},
									name: "YamlFrontMatterToken",
								},
							},
							&anyMatcher{
								line: 44, col: 50, offset: 1497,
							},
						},
					},
				},
			},
		},
		{
			name: "DocumentHeader",
			pos:  position{line: 52, col: 1, offset: 1721},
			expr: &actionExpr{
				pos: position{line: 52, col: 19, offset: 1739},
				run: (*parser).callonDocumentHeader1,
				expr: &seqExpr{
					pos: position{line: 52, col: 19, offset: 1739},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 52, col: 19, offset: 1739},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 52, col: 27, offset: 1747},
								name: "DocumentTitle",
							},
						},
						&labeledExpr{
							pos:   position{line: 52, col: 42, offset: 1762},
							label: "authors",
							expr: &zeroOrOneExpr{
								pos: position{line: 52, col: 51, offset: 1771},
								expr: &ruleRefExpr{
									pos:  position{line: 52, col: 51, offset: 1771},
									name: "DocumentAuthors",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 52, col: 69, offset: 1789},
							label: "revision",
							expr: &zeroOrOneExpr{
								pos: position{line: 52, col: 79, offset: 1799},
								expr: &ruleRefExpr{
									pos:  position{line: 52, col: 79, offset: 1799},
									name: "DocumentRevision",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 52, col: 98, offset: 1818},
							label: "otherAttributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 52, col: 115, offset: 1835},
								expr: &ruleRefExpr{
									pos:  position{line: 52, col: 115, offset: 1835},
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
			pos:  position{line: 56, col: 1, offset: 1966},
			expr: &actionExpr{
				pos: position{line: 56, col: 18, offset: 1983},
				run: (*parser).callonDocumentTitle1,
				expr: &seqExpr{
					pos: position{line: 56, col: 18, offset: 1983},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 56, col: 18, offset: 1983},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 56, col: 29, offset: 1994},
								expr: &ruleRefExpr{
									pos:  position{line: 56, col: 30, offset: 1995},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 56, col: 49, offset: 2014},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 56, col: 56, offset: 2021},
								val:        "=",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 56, col: 61, offset: 2026},
							expr: &ruleRefExpr{
								pos:  position{line: 56, col: 61, offset: 2026},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 56, col: 65, offset: 2030},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 56, col: 73, offset: 2038},
								name: "InlineContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 56, col: 87, offset: 2052},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthors",
			pos:  position{line: 60, col: 1, offset: 2156},
			expr: &choiceExpr{
				pos: position{line: 60, col: 20, offset: 2175},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 60, col: 20, offset: 2175},
						name: "DocumentAuthorsInlineForm",
					},
					&ruleRefExpr{
						pos:  position{line: 60, col: 48, offset: 2203},
						name: "DocumentAuthorsAttributeForm",
					},
				},
			},
		},
		{
			name: "DocumentAuthorsInlineForm",
			pos:  position{line: 62, col: 1, offset: 2233},
			expr: &actionExpr{
				pos: position{line: 62, col: 30, offset: 2262},
				run: (*parser).callonDocumentAuthorsInlineForm1,
				expr: &seqExpr{
					pos: position{line: 62, col: 30, offset: 2262},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 62, col: 30, offset: 2262},
							expr: &ruleRefExpr{
								pos:  position{line: 62, col: 30, offset: 2262},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 62, col: 34, offset: 2266},
							expr: &litMatcher{
								pos:        position{line: 62, col: 35, offset: 2267},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 62, col: 39, offset: 2271},
							label: "authors",
							expr: &oneOrMoreExpr{
								pos: position{line: 62, col: 48, offset: 2280},
								expr: &ruleRefExpr{
									pos:  position{line: 62, col: 48, offset: 2280},
									name: "DocumentAuthor",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 62, col: 65, offset: 2297},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthorsAttributeForm",
			pos:  position{line: 66, col: 1, offset: 2367},
			expr: &actionExpr{
				pos: position{line: 66, col: 33, offset: 2399},
				run: (*parser).callonDocumentAuthorsAttributeForm1,
				expr: &seqExpr{
					pos: position{line: 66, col: 33, offset: 2399},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 66, col: 33, offset: 2399},
							expr: &ruleRefExpr{
								pos:  position{line: 66, col: 33, offset: 2399},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 66, col: 37, offset: 2403},
							val:        ":author:",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 66, col: 48, offset: 2414},
							label: "author",
							expr: &ruleRefExpr{
								pos:  position{line: 66, col: 56, offset: 2422},
								name: "DocumentAuthor",
							},
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthor",
			pos:  position{line: 70, col: 1, offset: 2515},
			expr: &actionExpr{
				pos: position{line: 70, col: 19, offset: 2533},
				run: (*parser).callonDocumentAuthor1,
				expr: &seqExpr{
					pos: position{line: 70, col: 19, offset: 2533},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 70, col: 19, offset: 2533},
							expr: &ruleRefExpr{
								pos:  position{line: 70, col: 19, offset: 2533},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 70, col: 23, offset: 2537},
							label: "namePart1",
							expr: &ruleRefExpr{
								pos:  position{line: 70, col: 34, offset: 2548},
								name: "DocumentAuthorNamePart",
							},
						},
						&labeledExpr{
							pos:   position{line: 70, col: 58, offset: 2572},
							label: "namePart2",
							expr: &zeroOrOneExpr{
								pos: position{line: 70, col: 68, offset: 2582},
								expr: &ruleRefExpr{
									pos:  position{line: 70, col: 69, offset: 2583},
									name: "DocumentAuthorNamePart",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 70, col: 94, offset: 2608},
							label: "namePart3",
							expr: &zeroOrOneExpr{
								pos: position{line: 70, col: 104, offset: 2618},
								expr: &ruleRefExpr{
									pos:  position{line: 70, col: 105, offset: 2619},
									name: "DocumentAuthorNamePart",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 70, col: 130, offset: 2644},
							label: "email",
							expr: &zeroOrOneExpr{
								pos: position{line: 70, col: 136, offset: 2650},
								expr: &ruleRefExpr{
									pos:  position{line: 70, col: 137, offset: 2651},
									name: "DocumentAuthorEmail",
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 70, col: 159, offset: 2673},
							expr: &ruleRefExpr{
								pos:  position{line: 70, col: 159, offset: 2673},
								name: "WS",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 70, col: 163, offset: 2677},
							expr: &litMatcher{
								pos:        position{line: 70, col: 163, offset: 2677},
								val:        ";",
								ignoreCase: false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 70, col: 168, offset: 2682},
							expr: &ruleRefExpr{
								pos:  position{line: 70, col: 168, offset: 2682},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthorNamePart",
			pos:  position{line: 75, col: 1, offset: 2847},
			expr: &seqExpr{
				pos: position{line: 75, col: 27, offset: 2873},
				exprs: []interface{}{
					&notExpr{
						pos: position{line: 75, col: 27, offset: 2873},
						expr: &litMatcher{
							pos:        position{line: 75, col: 28, offset: 2874},
							val:        "<",
							ignoreCase: false,
						},
					},
					&notExpr{
						pos: position{line: 75, col: 32, offset: 2878},
						expr: &litMatcher{
							pos:        position{line: 75, col: 33, offset: 2879},
							val:        ";",
							ignoreCase: false,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 37, offset: 2883},
						name: "Characters",
					},
					&zeroOrMoreExpr{
						pos: position{line: 75, col: 48, offset: 2894},
						expr: &ruleRefExpr{
							pos:  position{line: 75, col: 48, offset: 2894},
							name: "WS",
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthorEmail",
			pos:  position{line: 77, col: 1, offset: 2899},
			expr: &seqExpr{
				pos: position{line: 77, col: 24, offset: 2922},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 77, col: 24, offset: 2922},
						val:        "<",
						ignoreCase: false,
					},
					&labeledExpr{
						pos:   position{line: 77, col: 28, offset: 2926},
						label: "email",
						expr: &oneOrMoreExpr{
							pos: position{line: 77, col: 34, offset: 2932},
							expr: &seqExpr{
								pos: position{line: 77, col: 35, offset: 2933},
								exprs: []interface{}{
									&notExpr{
										pos: position{line: 77, col: 35, offset: 2933},
										expr: &litMatcher{
											pos:        position{line: 77, col: 36, offset: 2934},
											val:        ">",
											ignoreCase: false,
										},
									},
									&notExpr{
										pos: position{line: 77, col: 40, offset: 2938},
										expr: &ruleRefExpr{
											pos:  position{line: 77, col: 41, offset: 2939},
											name: "EOL",
										},
									},
									&anyMatcher{
										line: 77, col: 45, offset: 2943,
									},
								},
							},
						},
					},
					&litMatcher{
						pos:        position{line: 77, col: 49, offset: 2947},
						val:        ">",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "DocumentRevision",
			pos:  position{line: 81, col: 1, offset: 3083},
			expr: &actionExpr{
				pos: position{line: 81, col: 21, offset: 3103},
				run: (*parser).callonDocumentRevision1,
				expr: &seqExpr{
					pos: position{line: 81, col: 21, offset: 3103},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 81, col: 21, offset: 3103},
							expr: &ruleRefExpr{
								pos:  position{line: 81, col: 21, offset: 3103},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 81, col: 25, offset: 3107},
							expr: &litMatcher{
								pos:        position{line: 81, col: 26, offset: 3108},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 81, col: 30, offset: 3112},
							label: "revnumber",
							expr: &zeroOrOneExpr{
								pos: position{line: 81, col: 40, offset: 3122},
								expr: &ruleRefExpr{
									pos:  position{line: 81, col: 41, offset: 3123},
									name: "DocumentRevisionNumber",
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 81, col: 66, offset: 3148},
							expr: &litMatcher{
								pos:        position{line: 81, col: 66, offset: 3148},
								val:        ",",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 81, col: 71, offset: 3153},
							label: "revdate",
							expr: &zeroOrOneExpr{
								pos: position{line: 81, col: 79, offset: 3161},
								expr: &ruleRefExpr{
									pos:  position{line: 81, col: 80, offset: 3162},
									name: "DocumentRevisionDate",
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 81, col: 103, offset: 3185},
							expr: &litMatcher{
								pos:        position{line: 81, col: 103, offset: 3185},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 81, col: 108, offset: 3190},
							label: "revremark",
							expr: &zeroOrOneExpr{
								pos: position{line: 81, col: 118, offset: 3200},
								expr: &ruleRefExpr{
									pos:  position{line: 81, col: 119, offset: 3201},
									name: "DocumentRevisionRemark",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 81, col: 144, offset: 3226},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentRevisionNumber",
			pos:  position{line: 86, col: 1, offset: 3399},
			expr: &choiceExpr{
				pos: position{line: 86, col: 27, offset: 3425},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 86, col: 27, offset: 3425},
						exprs: []interface{}{
							&litMatcher{
								pos:        position{line: 86, col: 27, offset: 3425},
								val:        "v",
								ignoreCase: true,
							},
							&ruleRefExpr{
								pos:  position{line: 86, col: 32, offset: 3430},
								name: "DIGIT",
							},
							&zeroOrMoreExpr{
								pos: position{line: 86, col: 39, offset: 3437},
								expr: &seqExpr{
									pos: position{line: 86, col: 40, offset: 3438},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 86, col: 40, offset: 3438},
											expr: &ruleRefExpr{
												pos:  position{line: 86, col: 41, offset: 3439},
												name: "EOL",
											},
										},
										&notExpr{
											pos: position{line: 86, col: 45, offset: 3443},
											expr: &litMatcher{
												pos:        position{line: 86, col: 46, offset: 3444},
												val:        ",",
												ignoreCase: false,
											},
										},
										&notExpr{
											pos: position{line: 86, col: 50, offset: 3448},
											expr: &litMatcher{
												pos:        position{line: 86, col: 51, offset: 3449},
												val:        ":",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 86, col: 55, offset: 3453,
										},
									},
								},
							},
						},
					},
					&seqExpr{
						pos: position{line: 86, col: 61, offset: 3459},
						exprs: []interface{}{
							&zeroOrOneExpr{
								pos: position{line: 86, col: 61, offset: 3459},
								expr: &litMatcher{
									pos:        position{line: 86, col: 61, offset: 3459},
									val:        "v",
									ignoreCase: true,
								},
							},
							&ruleRefExpr{
								pos:  position{line: 86, col: 67, offset: 3465},
								name: "DIGIT",
							},
							&zeroOrMoreExpr{
								pos: position{line: 86, col: 74, offset: 3472},
								expr: &seqExpr{
									pos: position{line: 86, col: 75, offset: 3473},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 86, col: 75, offset: 3473},
											expr: &ruleRefExpr{
												pos:  position{line: 86, col: 76, offset: 3474},
												name: "EOL",
											},
										},
										&notExpr{
											pos: position{line: 86, col: 80, offset: 3478},
											expr: &litMatcher{
												pos:        position{line: 86, col: 81, offset: 3479},
												val:        ",",
												ignoreCase: false,
											},
										},
										&notExpr{
											pos: position{line: 86, col: 85, offset: 3483},
											expr: &litMatcher{
												pos:        position{line: 86, col: 86, offset: 3484},
												val:        ":",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 86, col: 90, offset: 3488,
										},
									},
								},
							},
							&zeroOrMoreExpr{
								pos: position{line: 86, col: 94, offset: 3492},
								expr: &ruleRefExpr{
									pos:  position{line: 86, col: 94, offset: 3492},
									name: "WS",
								},
							},
							&andExpr{
								pos: position{line: 86, col: 98, offset: 3496},
								expr: &litMatcher{
									pos:        position{line: 86, col: 99, offset: 3497},
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
			pos:  position{line: 87, col: 1, offset: 3501},
			expr: &zeroOrMoreExpr{
				pos: position{line: 87, col: 25, offset: 3525},
				expr: &seqExpr{
					pos: position{line: 87, col: 26, offset: 3526},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 87, col: 26, offset: 3526},
							expr: &ruleRefExpr{
								pos:  position{line: 87, col: 27, offset: 3527},
								name: "EOL",
							},
						},
						&notExpr{
							pos: position{line: 87, col: 31, offset: 3531},
							expr: &litMatcher{
								pos:        position{line: 87, col: 32, offset: 3532},
								val:        ":",
								ignoreCase: false,
							},
						},
						&anyMatcher{
							line: 87, col: 36, offset: 3536,
						},
					},
				},
			},
		},
		{
			name: "DocumentRevisionRemark",
			pos:  position{line: 88, col: 1, offset: 3541},
			expr: &zeroOrMoreExpr{
				pos: position{line: 88, col: 27, offset: 3567},
				expr: &seqExpr{
					pos: position{line: 88, col: 28, offset: 3568},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 88, col: 28, offset: 3568},
							expr: &ruleRefExpr{
								pos:  position{line: 88, col: 29, offset: 3569},
								name: "EOL",
							},
						},
						&anyMatcher{
							line: 88, col: 33, offset: 3573,
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclaration",
			pos:  position{line: 93, col: 1, offset: 3693},
			expr: &choiceExpr{
				pos: position{line: 93, col: 33, offset: 3725},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 93, col: 33, offset: 3725},
						name: "DocumentAttributeDeclarationWithNameOnly",
					},
					&ruleRefExpr{
						pos:  position{line: 93, col: 76, offset: 3768},
						name: "DocumentAttributeDeclarationWithNameAndValue",
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclarationWithNameOnly",
			pos:  position{line: 95, col: 1, offset: 3815},
			expr: &actionExpr{
				pos: position{line: 95, col: 45, offset: 3859},
				run: (*parser).callonDocumentAttributeDeclarationWithNameOnly1,
				expr: &seqExpr{
					pos: position{line: 95, col: 45, offset: 3859},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 95, col: 45, offset: 3859},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 95, col: 49, offset: 3863},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 95, col: 55, offset: 3869},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 95, col: 70, offset: 3884},
							val:        ":",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 95, col: 74, offset: 3888},
							expr: &ruleRefExpr{
								pos:  position{line: 95, col: 74, offset: 3888},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 95, col: 78, offset: 3892},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclarationWithNameAndValue",
			pos:  position{line: 99, col: 1, offset: 3977},
			expr: &actionExpr{
				pos: position{line: 99, col: 49, offset: 4025},
				run: (*parser).callonDocumentAttributeDeclarationWithNameAndValue1,
				expr: &seqExpr{
					pos: position{line: 99, col: 49, offset: 4025},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 99, col: 49, offset: 4025},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 99, col: 53, offset: 4029},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 99, col: 59, offset: 4035},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 99, col: 74, offset: 4050},
							val:        ":",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 99, col: 78, offset: 4054},
							expr: &ruleRefExpr{
								pos:  position{line: 99, col: 78, offset: 4054},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 99, col: 82, offset: 4058},
							label: "value",
							expr: &zeroOrMoreExpr{
								pos: position{line: 99, col: 88, offset: 4064},
								expr: &seqExpr{
									pos: position{line: 99, col: 89, offset: 4065},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 99, col: 89, offset: 4065},
											expr: &ruleRefExpr{
												pos:  position{line: 99, col: 90, offset: 4066},
												name: "NEWLINE",
											},
										},
										&anyMatcher{
											line: 99, col: 98, offset: 4074,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 99, col: 102, offset: 4078},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeReset",
			pos:  position{line: 103, col: 1, offset: 4181},
			expr: &choiceExpr{
				pos: position{line: 103, col: 27, offset: 4207},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 103, col: 27, offset: 4207},
						name: "DocumentAttributeResetWithSectionTitleBangSymbol",
					},
					&ruleRefExpr{
						pos:  position{line: 103, col: 78, offset: 4258},
						name: "DocumentAttributeResetWithTrailingBangSymbol",
					},
				},
			},
		},
		{
			name: "DocumentAttributeResetWithSectionTitleBangSymbol",
			pos:  position{line: 105, col: 1, offset: 4304},
			expr: &actionExpr{
				pos: position{line: 105, col: 53, offset: 4356},
				run: (*parser).callonDocumentAttributeResetWithSectionTitleBangSymbol1,
				expr: &seqExpr{
					pos: position{line: 105, col: 53, offset: 4356},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 105, col: 53, offset: 4356},
							val:        ":!",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 105, col: 58, offset: 4361},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 105, col: 64, offset: 4367},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 105, col: 79, offset: 4382},
							val:        ":",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 105, col: 83, offset: 4386},
							expr: &ruleRefExpr{
								pos:  position{line: 105, col: 83, offset: 4386},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 105, col: 87, offset: 4390},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeResetWithTrailingBangSymbol",
			pos:  position{line: 109, col: 1, offset: 4464},
			expr: &actionExpr{
				pos: position{line: 109, col: 49, offset: 4512},
				run: (*parser).callonDocumentAttributeResetWithTrailingBangSymbol1,
				expr: &seqExpr{
					pos: position{line: 109, col: 49, offset: 4512},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 109, col: 49, offset: 4512},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 109, col: 53, offset: 4516},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 109, col: 59, offset: 4522},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 109, col: 74, offset: 4537},
							val:        "!:",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 109, col: 79, offset: 4542},
							expr: &ruleRefExpr{
								pos:  position{line: 109, col: 79, offset: 4542},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 109, col: 83, offset: 4546},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeSubstitution",
			pos:  position{line: 114, col: 1, offset: 4621},
			expr: &actionExpr{
				pos: position{line: 114, col: 34, offset: 4654},
				run: (*parser).callonDocumentAttributeSubstitution1,
				expr: &seqExpr{
					pos: position{line: 114, col: 34, offset: 4654},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 114, col: 34, offset: 4654},
							val:        "{",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 114, col: 38, offset: 4658},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 114, col: 44, offset: 4664},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 114, col: 59, offset: 4679},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "AttributeName",
			pos:  position{line: 121, col: 1, offset: 4933},
			expr: &seqExpr{
				pos: position{line: 121, col: 18, offset: 4950},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 121, col: 19, offset: 4951},
						alternatives: []interface{}{
							&charClassMatcher{
								pos:        position{line: 121, col: 19, offset: 4951},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
							&charClassMatcher{
								pos:        position{line: 121, col: 27, offset: 4959},
								val:        "[a-z]",
								ranges:     []rune{'a', 'z'},
								ignoreCase: false,
								inverted:   false,
							},
							&charClassMatcher{
								pos:        position{line: 121, col: 35, offset: 4967},
								val:        "[0-9]",
								ranges:     []rune{'0', '9'},
								ignoreCase: false,
								inverted:   false,
							},
							&litMatcher{
								pos:        position{line: 121, col: 43, offset: 4975},
								val:        "_",
								ignoreCase: false,
							},
						},
					},
					&zeroOrMoreExpr{
						pos: position{line: 121, col: 48, offset: 4980},
						expr: &choiceExpr{
							pos: position{line: 121, col: 49, offset: 4981},
							alternatives: []interface{}{
								&charClassMatcher{
									pos:        position{line: 121, col: 49, offset: 4981},
									val:        "[A-Z]",
									ranges:     []rune{'A', 'Z'},
									ignoreCase: false,
									inverted:   false,
								},
								&charClassMatcher{
									pos:        position{line: 121, col: 57, offset: 4989},
									val:        "[a-z]",
									ranges:     []rune{'a', 'z'},
									ignoreCase: false,
									inverted:   false,
								},
								&charClassMatcher{
									pos:        position{line: 121, col: 65, offset: 4997},
									val:        "[0-9]",
									ranges:     []rune{'0', '9'},
									ignoreCase: false,
									inverted:   false,
								},
								&litMatcher{
									pos:        position{line: 121, col: 73, offset: 5005},
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
			pos:  position{line: 126, col: 1, offset: 5125},
			expr: &seqExpr{
				pos: position{line: 126, col: 25, offset: 5149},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 126, col: 25, offset: 5149},
						val:        "toc::[]",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 126, col: 35, offset: 5159},
						name: "NEWLINE",
					},
				},
			},
		},
		{
			name: "Section",
			pos:  position{line: 131, col: 1, offset: 5272},
			expr: &choiceExpr{
				pos: position{line: 131, col: 12, offset: 5283},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 131, col: 12, offset: 5283},
						name: "Section1",
					},
					&ruleRefExpr{
						pos:  position{line: 131, col: 23, offset: 5294},
						name: "Section2",
					},
					&ruleRefExpr{
						pos:  position{line: 131, col: 34, offset: 5305},
						name: "Section3",
					},
					&ruleRefExpr{
						pos:  position{line: 131, col: 45, offset: 5316},
						name: "Section4",
					},
					&ruleRefExpr{
						pos:  position{line: 131, col: 56, offset: 5327},
						name: "Section5",
					},
				},
			},
		},
		{
			name: "Section1",
			pos:  position{line: 134, col: 1, offset: 5338},
			expr: &actionExpr{
				pos: position{line: 134, col: 13, offset: 5350},
				run: (*parser).callonSection11,
				expr: &seqExpr{
					pos: position{line: 134, col: 13, offset: 5350},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 134, col: 13, offset: 5350},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 134, col: 21, offset: 5358},
								name: "Section1Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 134, col: 36, offset: 5373},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 134, col: 46, offset: 5383},
								expr: &ruleRefExpr{
									pos:  position{line: 134, col: 46, offset: 5383},
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
			pos:  position{line: 138, col: 1, offset: 5491},
			expr: &actionExpr{
				pos: position{line: 138, col: 18, offset: 5508},
				run: (*parser).callonSection1Block1,
				expr: &seqExpr{
					pos: position{line: 138, col: 18, offset: 5508},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 138, col: 18, offset: 5508},
							expr: &ruleRefExpr{
								pos:  position{line: 138, col: 19, offset: 5509},
								name: "Section1",
							},
						},
						&labeledExpr{
							pos:   position{line: 138, col: 28, offset: 5518},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 138, col: 37, offset: 5527},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 138, col: 37, offset: 5527},
										name: "Section2",
									},
									&ruleRefExpr{
										pos:  position{line: 138, col: 48, offset: 5538},
										name: "Section3",
									},
									&ruleRefExpr{
										pos:  position{line: 138, col: 59, offset: 5549},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 138, col: 70, offset: 5560},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 138, col: 81, offset: 5571},
										name: "BlockElement",
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
			pos:  position{line: 142, col: 1, offset: 5633},
			expr: &actionExpr{
				pos: position{line: 142, col: 13, offset: 5645},
				run: (*parser).callonSection21,
				expr: &seqExpr{
					pos: position{line: 142, col: 13, offset: 5645},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 142, col: 13, offset: 5645},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 142, col: 21, offset: 5653},
								name: "Section2Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 142, col: 36, offset: 5668},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 142, col: 46, offset: 5678},
								expr: &ruleRefExpr{
									pos:  position{line: 142, col: 46, offset: 5678},
									name: "Section2Block",
								},
							},
						},
						&andExpr{
							pos: position{line: 142, col: 62, offset: 5694},
							expr: &zeroOrMoreExpr{
								pos: position{line: 142, col: 63, offset: 5695},
								expr: &ruleRefExpr{
									pos:  position{line: 142, col: 64, offset: 5696},
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
			pos:  position{line: 146, col: 1, offset: 5799},
			expr: &actionExpr{
				pos: position{line: 146, col: 18, offset: 5816},
				run: (*parser).callonSection2Block1,
				expr: &seqExpr{
					pos: position{line: 146, col: 18, offset: 5816},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 146, col: 18, offset: 5816},
							expr: &ruleRefExpr{
								pos:  position{line: 146, col: 19, offset: 5817},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 146, col: 28, offset: 5826},
							expr: &ruleRefExpr{
								pos:  position{line: 146, col: 29, offset: 5827},
								name: "Section2",
							},
						},
						&labeledExpr{
							pos:   position{line: 146, col: 38, offset: 5836},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 146, col: 47, offset: 5845},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 146, col: 47, offset: 5845},
										name: "Section3",
									},
									&ruleRefExpr{
										pos:  position{line: 146, col: 58, offset: 5856},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 146, col: 69, offset: 5867},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 146, col: 80, offset: 5878},
										name: "BlockElement",
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
			pos:  position{line: 150, col: 1, offset: 5940},
			expr: &actionExpr{
				pos: position{line: 150, col: 13, offset: 5952},
				run: (*parser).callonSection31,
				expr: &seqExpr{
					pos: position{line: 150, col: 13, offset: 5952},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 150, col: 13, offset: 5952},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 150, col: 21, offset: 5960},
								name: "Section3Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 150, col: 36, offset: 5975},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 150, col: 46, offset: 5985},
								expr: &ruleRefExpr{
									pos:  position{line: 150, col: 46, offset: 5985},
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
			pos:  position{line: 154, col: 1, offset: 6093},
			expr: &actionExpr{
				pos: position{line: 154, col: 18, offset: 6110},
				run: (*parser).callonSection3Block1,
				expr: &seqExpr{
					pos: position{line: 154, col: 18, offset: 6110},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 154, col: 18, offset: 6110},
							expr: &ruleRefExpr{
								pos:  position{line: 154, col: 19, offset: 6111},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 154, col: 28, offset: 6120},
							expr: &ruleRefExpr{
								pos:  position{line: 154, col: 29, offset: 6121},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 154, col: 38, offset: 6130},
							expr: &ruleRefExpr{
								pos:  position{line: 154, col: 39, offset: 6131},
								name: "Section3",
							},
						},
						&labeledExpr{
							pos:   position{line: 154, col: 48, offset: 6140},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 154, col: 57, offset: 6149},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 154, col: 57, offset: 6149},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 154, col: 68, offset: 6160},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 154, col: 79, offset: 6171},
										name: "BlockElement",
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
			pos:  position{line: 158, col: 1, offset: 6233},
			expr: &actionExpr{
				pos: position{line: 158, col: 13, offset: 6245},
				run: (*parser).callonSection41,
				expr: &seqExpr{
					pos: position{line: 158, col: 13, offset: 6245},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 158, col: 13, offset: 6245},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 158, col: 21, offset: 6253},
								name: "Section4Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 158, col: 36, offset: 6268},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 158, col: 46, offset: 6278},
								expr: &ruleRefExpr{
									pos:  position{line: 158, col: 46, offset: 6278},
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
			pos:  position{line: 162, col: 1, offset: 6386},
			expr: &actionExpr{
				pos: position{line: 162, col: 18, offset: 6403},
				run: (*parser).callonSection4Block1,
				expr: &seqExpr{
					pos: position{line: 162, col: 18, offset: 6403},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 162, col: 18, offset: 6403},
							expr: &ruleRefExpr{
								pos:  position{line: 162, col: 19, offset: 6404},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 162, col: 28, offset: 6413},
							expr: &ruleRefExpr{
								pos:  position{line: 162, col: 29, offset: 6414},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 162, col: 38, offset: 6423},
							expr: &ruleRefExpr{
								pos:  position{line: 162, col: 39, offset: 6424},
								name: "Section3",
							},
						},
						&notExpr{
							pos: position{line: 162, col: 48, offset: 6433},
							expr: &ruleRefExpr{
								pos:  position{line: 162, col: 49, offset: 6434},
								name: "Section4",
							},
						},
						&labeledExpr{
							pos:   position{line: 162, col: 58, offset: 6443},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 162, col: 67, offset: 6452},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 162, col: 67, offset: 6452},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 162, col: 78, offset: 6463},
										name: "BlockElement",
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
			pos:  position{line: 166, col: 1, offset: 6525},
			expr: &actionExpr{
				pos: position{line: 166, col: 13, offset: 6537},
				run: (*parser).callonSection51,
				expr: &seqExpr{
					pos: position{line: 166, col: 13, offset: 6537},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 166, col: 13, offset: 6537},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 166, col: 21, offset: 6545},
								name: "Section5Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 166, col: 36, offset: 6560},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 166, col: 46, offset: 6570},
								expr: &ruleRefExpr{
									pos:  position{line: 166, col: 46, offset: 6570},
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
			pos:  position{line: 170, col: 1, offset: 6678},
			expr: &actionExpr{
				pos: position{line: 170, col: 18, offset: 6695},
				run: (*parser).callonSection5Block1,
				expr: &seqExpr{
					pos: position{line: 170, col: 18, offset: 6695},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 170, col: 18, offset: 6695},
							expr: &ruleRefExpr{
								pos:  position{line: 170, col: 19, offset: 6696},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 170, col: 28, offset: 6705},
							expr: &ruleRefExpr{
								pos:  position{line: 170, col: 29, offset: 6706},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 170, col: 38, offset: 6715},
							expr: &ruleRefExpr{
								pos:  position{line: 170, col: 39, offset: 6716},
								name: "Section3",
							},
						},
						&notExpr{
							pos: position{line: 170, col: 48, offset: 6725},
							expr: &ruleRefExpr{
								pos:  position{line: 170, col: 49, offset: 6726},
								name: "Section4",
							},
						},
						&notExpr{
							pos: position{line: 170, col: 58, offset: 6735},
							expr: &ruleRefExpr{
								pos:  position{line: 170, col: 59, offset: 6736},
								name: "Section5",
							},
						},
						&labeledExpr{
							pos:   position{line: 170, col: 68, offset: 6745},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 170, col: 77, offset: 6754},
								name: "BlockElement",
							},
						},
					},
				},
			},
		},
		{
			name: "SectionTitle",
			pos:  position{line: 178, col: 1, offset: 6927},
			expr: &choiceExpr{
				pos: position{line: 178, col: 17, offset: 6943},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 178, col: 17, offset: 6943},
						name: "Section1Title",
					},
					&ruleRefExpr{
						pos:  position{line: 178, col: 33, offset: 6959},
						name: "Section2Title",
					},
					&ruleRefExpr{
						pos:  position{line: 178, col: 49, offset: 6975},
						name: "Section3Title",
					},
					&ruleRefExpr{
						pos:  position{line: 178, col: 65, offset: 6991},
						name: "Section4Title",
					},
					&ruleRefExpr{
						pos:  position{line: 178, col: 81, offset: 7007},
						name: "Section5Title",
					},
				},
			},
		},
		{
			name: "Section1Title",
			pos:  position{line: 180, col: 1, offset: 7022},
			expr: &actionExpr{
				pos: position{line: 180, col: 18, offset: 7039},
				run: (*parser).callonSection1Title1,
				expr: &seqExpr{
					pos: position{line: 180, col: 18, offset: 7039},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 180, col: 18, offset: 7039},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 180, col: 29, offset: 7050},
								expr: &ruleRefExpr{
									pos:  position{line: 180, col: 30, offset: 7051},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 180, col: 49, offset: 7070},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 180, col: 56, offset: 7077},
								val:        "==",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 180, col: 62, offset: 7083},
							expr: &ruleRefExpr{
								pos:  position{line: 180, col: 62, offset: 7083},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 180, col: 66, offset: 7087},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 180, col: 74, offset: 7095},
								name: "InlineContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 180, col: 88, offset: 7109},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 180, col: 93, offset: 7114},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 180, col: 93, offset: 7114},
									expr: &ruleRefExpr{
										pos:  position{line: 180, col: 93, offset: 7114},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 180, col: 106, offset: 7127},
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
			pos:  position{line: 184, col: 1, offset: 7232},
			expr: &actionExpr{
				pos: position{line: 184, col: 18, offset: 7249},
				run: (*parser).callonSection2Title1,
				expr: &seqExpr{
					pos: position{line: 184, col: 18, offset: 7249},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 184, col: 18, offset: 7249},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 184, col: 29, offset: 7260},
								expr: &ruleRefExpr{
									pos:  position{line: 184, col: 30, offset: 7261},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 184, col: 49, offset: 7280},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 184, col: 56, offset: 7287},
								val:        "===",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 184, col: 63, offset: 7294},
							expr: &ruleRefExpr{
								pos:  position{line: 184, col: 63, offset: 7294},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 184, col: 67, offset: 7298},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 184, col: 75, offset: 7306},
								name: "InlineContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 184, col: 89, offset: 7320},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 184, col: 94, offset: 7325},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 184, col: 94, offset: 7325},
									expr: &ruleRefExpr{
										pos:  position{line: 184, col: 94, offset: 7325},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 184, col: 107, offset: 7338},
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
			pos:  position{line: 188, col: 1, offset: 7442},
			expr: &actionExpr{
				pos: position{line: 188, col: 18, offset: 7459},
				run: (*parser).callonSection3Title1,
				expr: &seqExpr{
					pos: position{line: 188, col: 18, offset: 7459},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 188, col: 18, offset: 7459},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 188, col: 29, offset: 7470},
								expr: &ruleRefExpr{
									pos:  position{line: 188, col: 30, offset: 7471},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 188, col: 49, offset: 7490},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 188, col: 56, offset: 7497},
								val:        "====",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 188, col: 64, offset: 7505},
							expr: &ruleRefExpr{
								pos:  position{line: 188, col: 64, offset: 7505},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 188, col: 68, offset: 7509},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 188, col: 76, offset: 7517},
								name: "InlineContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 188, col: 90, offset: 7531},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 188, col: 95, offset: 7536},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 188, col: 95, offset: 7536},
									expr: &ruleRefExpr{
										pos:  position{line: 188, col: 95, offset: 7536},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 188, col: 108, offset: 7549},
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
			pos:  position{line: 192, col: 1, offset: 7653},
			expr: &actionExpr{
				pos: position{line: 192, col: 18, offset: 7670},
				run: (*parser).callonSection4Title1,
				expr: &seqExpr{
					pos: position{line: 192, col: 18, offset: 7670},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 192, col: 18, offset: 7670},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 192, col: 29, offset: 7681},
								expr: &ruleRefExpr{
									pos:  position{line: 192, col: 30, offset: 7682},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 192, col: 49, offset: 7701},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 192, col: 56, offset: 7708},
								val:        "=====",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 192, col: 65, offset: 7717},
							expr: &ruleRefExpr{
								pos:  position{line: 192, col: 65, offset: 7717},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 192, col: 69, offset: 7721},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 192, col: 77, offset: 7729},
								name: "InlineContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 192, col: 91, offset: 7743},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 192, col: 96, offset: 7748},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 192, col: 96, offset: 7748},
									expr: &ruleRefExpr{
										pos:  position{line: 192, col: 96, offset: 7748},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 192, col: 109, offset: 7761},
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
			pos:  position{line: 196, col: 1, offset: 7865},
			expr: &actionExpr{
				pos: position{line: 196, col: 18, offset: 7882},
				run: (*parser).callonSection5Title1,
				expr: &seqExpr{
					pos: position{line: 196, col: 18, offset: 7882},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 196, col: 18, offset: 7882},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 196, col: 29, offset: 7893},
								expr: &ruleRefExpr{
									pos:  position{line: 196, col: 30, offset: 7894},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 196, col: 49, offset: 7913},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 196, col: 56, offset: 7920},
								val:        "======",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 196, col: 66, offset: 7930},
							expr: &ruleRefExpr{
								pos:  position{line: 196, col: 66, offset: 7930},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 196, col: 70, offset: 7934},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 196, col: 78, offset: 7942},
								name: "InlineContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 196, col: 92, offset: 7956},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 196, col: 97, offset: 7961},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 196, col: 97, offset: 7961},
									expr: &ruleRefExpr{
										pos:  position{line: 196, col: 97, offset: 7961},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 196, col: 110, offset: 7974},
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
			pos:  position{line: 203, col: 1, offset: 8179},
			expr: &actionExpr{
				pos: position{line: 203, col: 9, offset: 8187},
				run: (*parser).callonList1,
				expr: &seqExpr{
					pos: position{line: 203, col: 9, offset: 8187},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 203, col: 9, offset: 8187},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 203, col: 20, offset: 8198},
								expr: &ruleRefExpr{
									pos:  position{line: 203, col: 21, offset: 8199},
									name: "ListAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 205, col: 5, offset: 8288},
							label: "elements",
							expr: &oneOrMoreExpr{
								pos: position{line: 205, col: 14, offset: 8297},
								expr: &choiceExpr{
									pos: position{line: 205, col: 15, offset: 8298},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 205, col: 15, offset: 8298},
											name: "UnorderedListItem",
										},
										&ruleRefExpr{
											pos:  position{line: 205, col: 35, offset: 8318},
											name: "LabeledListItem",
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
			name: "ListAttribute",
			pos:  position{line: 209, col: 1, offset: 8420},
			expr: &actionExpr{
				pos: position{line: 209, col: 18, offset: 8437},
				run: (*parser).callonListAttribute1,
				expr: &seqExpr{
					pos: position{line: 209, col: 18, offset: 8437},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 209, col: 18, offset: 8437},
							label: "attribute",
							expr: &choiceExpr{
								pos: position{line: 209, col: 29, offset: 8448},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 209, col: 29, offset: 8448},
										name: "HorizontalLayout",
									},
									&ruleRefExpr{
										pos:  position{line: 209, col: 48, offset: 8467},
										name: "ListID",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 209, col: 56, offset: 8475},
							name: "NEWLINE",
						},
					},
				},
			},
		},
		{
			name: "ListID",
			pos:  position{line: 213, col: 1, offset: 8514},
			expr: &actionExpr{
				pos: position{line: 213, col: 11, offset: 8524},
				run: (*parser).callonListID1,
				expr: &seqExpr{
					pos: position{line: 213, col: 11, offset: 8524},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 213, col: 11, offset: 8524},
							val:        "[#",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 213, col: 16, offset: 8529},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 213, col: 20, offset: 8533},
								name: "ID",
							},
						},
						&litMatcher{
							pos:        position{line: 213, col: 24, offset: 8537},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "HorizontalLayout",
			pos:  position{line: 217, col: 1, offset: 8603},
			expr: &actionExpr{
				pos: position{line: 217, col: 21, offset: 8623},
				run: (*parser).callonHorizontalLayout1,
				expr: &litMatcher{
					pos:        position{line: 217, col: 21, offset: 8623},
					val:        "[horizontal]",
					ignoreCase: false,
				},
			},
		},
		{
			name: "ListParagraph",
			pos:  position{line: 221, col: 1, offset: 8706},
			expr: &actionExpr{
				pos: position{line: 221, col: 19, offset: 8724},
				run: (*parser).callonListParagraph1,
				expr: &labeledExpr{
					pos:   position{line: 221, col: 19, offset: 8724},
					label: "lines",
					expr: &oneOrMoreExpr{
						pos: position{line: 221, col: 25, offset: 8730},
						expr: &seqExpr{
							pos: position{line: 221, col: 26, offset: 8731},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 221, col: 26, offset: 8731},
									expr: &ruleRefExpr{
										pos:  position{line: 221, col: 28, offset: 8733},
										name: "UnorderedListItemPrefix",
									},
								},
								&notExpr{
									pos: position{line: 221, col: 53, offset: 8758},
									expr: &seqExpr{
										pos: position{line: 221, col: 55, offset: 8760},
										exprs: []interface{}{
											&ruleRefExpr{
												pos:  position{line: 221, col: 55, offset: 8760},
												name: "LabeledListItemTerm",
											},
											&ruleRefExpr{
												pos:  position{line: 221, col: 75, offset: 8780},
												name: "LabeledListItemSeparator",
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 221, col: 101, offset: 8806},
									name: "InlineContent",
								},
								&ruleRefExpr{
									pos:  position{line: 221, col: 115, offset: 8820},
									name: "EOL",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "UnorderedListItem",
			pos:  position{line: 236, col: 1, offset: 9261},
			expr: &actionExpr{
				pos: position{line: 236, col: 22, offset: 9282},
				run: (*parser).callonUnorderedListItem1,
				expr: &seqExpr{
					pos: position{line: 236, col: 22, offset: 9282},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 236, col: 22, offset: 9282},
							label: "level",
							expr: &ruleRefExpr{
								pos:  position{line: 236, col: 29, offset: 9289},
								name: "UnorderedListItemPrefix",
							},
						},
						&labeledExpr{
							pos:   position{line: 236, col: 54, offset: 9314},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 236, col: 63, offset: 9323},
								name: "UnorderedListItemContent",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 236, col: 89, offset: 9349},
							expr: &ruleRefExpr{
								pos:  position{line: 236, col: 89, offset: 9349},
								name: "BlankLine",
							},
						},
					},
				},
			},
		},
		{
			name: "UnorderedListItemPrefix",
			pos:  position{line: 240, col: 1, offset: 9440},
			expr: &actionExpr{
				pos: position{line: 240, col: 28, offset: 9467},
				run: (*parser).callonUnorderedListItemPrefix1,
				expr: &seqExpr{
					pos: position{line: 240, col: 28, offset: 9467},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 240, col: 28, offset: 9467},
							expr: &ruleRefExpr{
								pos:  position{line: 240, col: 28, offset: 9467},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 240, col: 32, offset: 9471},
							label: "level",
							expr: &choiceExpr{
								pos: position{line: 240, col: 39, offset: 9478},
								alternatives: []interface{}{
									&oneOrMoreExpr{
										pos: position{line: 240, col: 39, offset: 9478},
										expr: &litMatcher{
											pos:        position{line: 240, col: 39, offset: 9478},
											val:        "*",
											ignoreCase: false,
										},
									},
									&litMatcher{
										pos:        position{line: 240, col: 46, offset: 9485},
										val:        "-",
										ignoreCase: false,
									},
								},
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 240, col: 51, offset: 9490},
							expr: &ruleRefExpr{
								pos:  position{line: 240, col: 51, offset: 9490},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "UnorderedListItemContent",
			pos:  position{line: 244, col: 1, offset: 9588},
			expr: &actionExpr{
				pos: position{line: 244, col: 29, offset: 9616},
				run: (*parser).callonUnorderedListItemContent1,
				expr: &labeledExpr{
					pos:   position{line: 244, col: 29, offset: 9616},
					label: "elements",
					expr: &oneOrMoreExpr{
						pos: position{line: 244, col: 38, offset: 9625},
						expr: &ruleRefExpr{
							pos:  position{line: 244, col: 39, offset: 9626},
							name: "ListParagraph",
						},
					},
				},
			},
		},
		{
			name: "LabeledListItem",
			pos:  position{line: 256, col: 1, offset: 10112},
			expr: &choiceExpr{
				pos: position{line: 256, col: 20, offset: 10131},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 256, col: 20, offset: 10131},
						name: "LabeledListItemWithDescription",
					},
					&ruleRefExpr{
						pos:  position{line: 256, col: 53, offset: 10164},
						name: "LabeledListItemWithTermAlone",
					},
				},
			},
		},
		{
			name: "LabeledListItemWithTermAlone",
			pos:  position{line: 258, col: 1, offset: 10194},
			expr: &actionExpr{
				pos: position{line: 258, col: 33, offset: 10226},
				run: (*parser).callonLabeledListItemWithTermAlone1,
				expr: &seqExpr{
					pos: position{line: 258, col: 33, offset: 10226},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 258, col: 33, offset: 10226},
							label: "term",
							expr: &ruleRefExpr{
								pos:  position{line: 258, col: 39, offset: 10232},
								name: "LabeledListItemTerm",
							},
						},
						&litMatcher{
							pos:        position{line: 258, col: 61, offset: 10254},
							val:        "::",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 258, col: 66, offset: 10259},
							expr: &ruleRefExpr{
								pos:  position{line: 258, col: 66, offset: 10259},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 258, col: 70, offset: 10263},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "LabeledListItemTerm",
			pos:  position{line: 262, col: 1, offset: 10400},
			expr: &actionExpr{
				pos: position{line: 262, col: 24, offset: 10423},
				run: (*parser).callonLabeledListItemTerm1,
				expr: &labeledExpr{
					pos:   position{line: 262, col: 24, offset: 10423},
					label: "term",
					expr: &zeroOrMoreExpr{
						pos: position{line: 262, col: 29, offset: 10428},
						expr: &seqExpr{
							pos: position{line: 262, col: 30, offset: 10429},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 262, col: 30, offset: 10429},
									expr: &ruleRefExpr{
										pos:  position{line: 262, col: 31, offset: 10430},
										name: "NEWLINE",
									},
								},
								&notExpr{
									pos: position{line: 262, col: 39, offset: 10438},
									expr: &litMatcher{
										pos:        position{line: 262, col: 40, offset: 10439},
										val:        "::",
										ignoreCase: false,
									},
								},
								&anyMatcher{
									line: 262, col: 45, offset: 10444,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "LabeledListItemSeparator",
			pos:  position{line: 267, col: 1, offset: 10535},
			expr: &seqExpr{
				pos: position{line: 267, col: 30, offset: 10564},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 267, col: 30, offset: 10564},
						val:        "::",
						ignoreCase: false,
					},
					&oneOrMoreExpr{
						pos: position{line: 267, col: 35, offset: 10569},
						expr: &choiceExpr{
							pos: position{line: 267, col: 36, offset: 10570},
							alternatives: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 267, col: 36, offset: 10570},
									name: "WS",
								},
								&ruleRefExpr{
									pos:  position{line: 267, col: 41, offset: 10575},
									name: "NEWLINE",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "LabeledListItemWithDescription",
			pos:  position{line: 269, col: 1, offset: 10586},
			expr: &actionExpr{
				pos: position{line: 269, col: 35, offset: 10620},
				run: (*parser).callonLabeledListItemWithDescription1,
				expr: &seqExpr{
					pos: position{line: 269, col: 35, offset: 10620},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 269, col: 35, offset: 10620},
							label: "term",
							expr: &ruleRefExpr{
								pos:  position{line: 269, col: 41, offset: 10626},
								name: "LabeledListItemTerm",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 269, col: 62, offset: 10647},
							name: "LabeledListItemSeparator",
						},
						&labeledExpr{
							pos:   position{line: 269, col: 87, offset: 10672},
							label: "description",
							expr: &ruleRefExpr{
								pos:  position{line: 269, col: 100, offset: 10685},
								name: "LabeledListItemDescription",
							},
						},
					},
				},
			},
		},
		{
			name: "LabeledListItemDescription",
			pos:  position{line: 273, col: 1, offset: 10810},
			expr: &actionExpr{
				pos: position{line: 273, col: 31, offset: 10840},
				run: (*parser).callonLabeledListItemDescription1,
				expr: &labeledExpr{
					pos:   position{line: 273, col: 31, offset: 10840},
					label: "elements",
					expr: &zeroOrMoreExpr{
						pos: position{line: 273, col: 40, offset: 10849},
						expr: &ruleRefExpr{
							pos:  position{line: 273, col: 41, offset: 10850},
							name: "ListParagraph",
						},
					},
				},
			},
		},
		{
			name: "Paragraph",
			pos:  position{line: 282, col: 1, offset: 11192},
			expr: &actionExpr{
				pos: position{line: 282, col: 14, offset: 11205},
				run: (*parser).callonParagraph1,
				expr: &seqExpr{
					pos: position{line: 282, col: 14, offset: 11205},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 282, col: 14, offset: 11205},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 282, col: 25, offset: 11216},
								expr: &ruleRefExpr{
									pos:  position{line: 282, col: 26, offset: 11217},
									name: "ElementAttribute",
								},
							},
						},
						&notExpr{
							pos: position{line: 282, col: 45, offset: 11236},
							expr: &seqExpr{
								pos: position{line: 282, col: 47, offset: 11238},
								exprs: []interface{}{
									&oneOrMoreExpr{
										pos: position{line: 282, col: 47, offset: 11238},
										expr: &litMatcher{
											pos:        position{line: 282, col: 47, offset: 11238},
											val:        "=",
											ignoreCase: false,
										},
									},
									&oneOrMoreExpr{
										pos: position{line: 282, col: 52, offset: 11243},
										expr: &ruleRefExpr{
											pos:  position{line: 282, col: 52, offset: 11243},
											name: "WS",
										},
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 282, col: 57, offset: 11248},
							label: "lines",
							expr: &oneOrMoreExpr{
								pos: position{line: 282, col: 63, offset: 11254},
								expr: &seqExpr{
									pos: position{line: 282, col: 64, offset: 11255},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 282, col: 64, offset: 11255},
											name: "InlineContent",
										},
										&ruleRefExpr{
											pos:  position{line: 282, col: 78, offset: 11269},
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
			pos:  position{line: 288, col: 1, offset: 11559},
			expr: &actionExpr{
				pos: position{line: 288, col: 18, offset: 11576},
				run: (*parser).callonInlineContent1,
				expr: &seqExpr{
					pos: position{line: 288, col: 18, offset: 11576},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 288, col: 18, offset: 11576},
							expr: &ruleRefExpr{
								pos:  position{line: 288, col: 19, offset: 11577},
								name: "BlockDelimiter",
							},
						},
						&labeledExpr{
							pos:   position{line: 288, col: 34, offset: 11592},
							label: "elements",
							expr: &oneOrMoreExpr{
								pos: position{line: 288, col: 43, offset: 11601},
								expr: &seqExpr{
									pos: position{line: 288, col: 44, offset: 11602},
									exprs: []interface{}{
										&zeroOrMoreExpr{
											pos: position{line: 288, col: 44, offset: 11602},
											expr: &ruleRefExpr{
												pos:  position{line: 288, col: 44, offset: 11602},
												name: "WS",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 288, col: 48, offset: 11606},
											name: "InlineElement",
										},
										&zeroOrMoreExpr{
											pos: position{line: 288, col: 62, offset: 11620},
											expr: &ruleRefExpr{
												pos:  position{line: 288, col: 62, offset: 11620},
												name: "WS",
											},
										},
									},
								},
							},
						},
						&andExpr{
							pos: position{line: 288, col: 68, offset: 11626},
							expr: &ruleRefExpr{
								pos:  position{line: 288, col: 69, offset: 11627},
								name: "EOL",
							},
						},
					},
				},
			},
		},
		{
			name: "InlineElement",
			pos:  position{line: 292, col: 1, offset: 11745},
			expr: &choiceExpr{
				pos: position{line: 292, col: 18, offset: 11762},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 292, col: 18, offset: 11762},
						name: "CrossReference",
					},
					&ruleRefExpr{
						pos:  position{line: 292, col: 35, offset: 11779},
						name: "Passthrough",
					},
					&ruleRefExpr{
						pos:  position{line: 292, col: 49, offset: 11793},
						name: "InlineImage",
					},
					&ruleRefExpr{
						pos:  position{line: 292, col: 63, offset: 11807},
						name: "QuotedText",
					},
					&ruleRefExpr{
						pos:  position{line: 292, col: 76, offset: 11820},
						name: "ExternalLink",
					},
					&ruleRefExpr{
						pos:  position{line: 292, col: 91, offset: 11835},
						name: "DocumentAttributeSubstitution",
					},
					&ruleRefExpr{
						pos:  position{line: 292, col: 123, offset: 11867},
						name: "Characters",
					},
				},
			},
		},
		{
			name: "QuotedText",
			pos:  position{line: 297, col: 1, offset: 12118},
			expr: &choiceExpr{
				pos: position{line: 297, col: 15, offset: 12132},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 297, col: 15, offset: 12132},
						name: "BoldText",
					},
					&ruleRefExpr{
						pos:  position{line: 297, col: 26, offset: 12143},
						name: "ItalicText",
					},
					&ruleRefExpr{
						pos:  position{line: 297, col: 39, offset: 12156},
						name: "MonospaceText",
					},
					&ruleRefExpr{
						pos:  position{line: 298, col: 13, offset: 12184},
						name: "EscapedBoldText",
					},
					&ruleRefExpr{
						pos:  position{line: 298, col: 31, offset: 12202},
						name: "EscapedItalicText",
					},
					&ruleRefExpr{
						pos:  position{line: 298, col: 51, offset: 12222},
						name: "EscapedMonospaceText",
					},
				},
			},
		},
		{
			name: "BoldText",
			pos:  position{line: 300, col: 1, offset: 12244},
			expr: &choiceExpr{
				pos: position{line: 300, col: 13, offset: 12256},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 300, col: 13, offset: 12256},
						name: "BoldTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 300, col: 41, offset: 12284},
						name: "BoldTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 300, col: 73, offset: 12316},
						name: "BoldTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "BoldTextSimplePunctuation",
			pos:  position{line: 302, col: 1, offset: 12389},
			expr: &actionExpr{
				pos: position{line: 302, col: 30, offset: 12418},
				run: (*parser).callonBoldTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 302, col: 30, offset: 12418},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 302, col: 30, offset: 12418},
							expr: &litMatcher{
								pos:        position{line: 302, col: 31, offset: 12419},
								val:        "\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 302, col: 35, offset: 12423},
							val:        "*",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 302, col: 39, offset: 12427},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 302, col: 48, offset: 12436},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 302, col: 67, offset: 12455},
							val:        "*",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "BoldTextDoublePunctuation",
			pos:  position{line: 306, col: 1, offset: 12532},
			expr: &actionExpr{
				pos: position{line: 306, col: 30, offset: 12561},
				run: (*parser).callonBoldTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 306, col: 30, offset: 12561},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 306, col: 30, offset: 12561},
							expr: &litMatcher{
								pos:        position{line: 306, col: 31, offset: 12562},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 306, col: 36, offset: 12567},
							val:        "**",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 306, col: 41, offset: 12572},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 306, col: 50, offset: 12581},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 306, col: 69, offset: 12600},
							val:        "**",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "BoldTextUnbalancedPunctuation",
			pos:  position{line: 310, col: 1, offset: 12678},
			expr: &actionExpr{
				pos: position{line: 310, col: 34, offset: 12711},
				run: (*parser).callonBoldTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 310, col: 34, offset: 12711},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 310, col: 34, offset: 12711},
							expr: &litMatcher{
								pos:        position{line: 310, col: 35, offset: 12712},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 310, col: 40, offset: 12717},
							val:        "**",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 310, col: 45, offset: 12722},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 310, col: 54, offset: 12731},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 310, col: 73, offset: 12750},
							val:        "*",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedBoldText",
			pos:  position{line: 315, col: 1, offset: 12914},
			expr: &choiceExpr{
				pos: position{line: 315, col: 20, offset: 12933},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 315, col: 20, offset: 12933},
						name: "EscapedBoldTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 315, col: 55, offset: 12968},
						name: "EscapedBoldTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 315, col: 94, offset: 13007},
						name: "EscapedBoldTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "EscapedBoldTextSimplePunctuation",
			pos:  position{line: 317, col: 1, offset: 13087},
			expr: &actionExpr{
				pos: position{line: 317, col: 37, offset: 13123},
				run: (*parser).callonEscapedBoldTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 317, col: 37, offset: 13123},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 317, col: 37, offset: 13123},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 317, col: 50, offset: 13136},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 317, col: 50, offset: 13136},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 317, col: 54, offset: 13140},
										expr: &litMatcher{
											pos:        position{line: 317, col: 54, offset: 13140},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 317, col: 60, offset: 13146},
							val:        "*",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 317, col: 64, offset: 13150},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 317, col: 73, offset: 13159},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 317, col: 92, offset: 13178},
							val:        "*",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedBoldTextDoublePunctuation",
			pos:  position{line: 321, col: 1, offset: 13284},
			expr: &actionExpr{
				pos: position{line: 321, col: 37, offset: 13320},
				run: (*parser).callonEscapedBoldTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 321, col: 37, offset: 13320},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 321, col: 37, offset: 13320},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 321, col: 50, offset: 13333},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 321, col: 50, offset: 13333},
										val:        "\\\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 321, col: 55, offset: 13338},
										expr: &litMatcher{
											pos:        position{line: 321, col: 55, offset: 13338},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 321, col: 61, offset: 13344},
							val:        "**",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 321, col: 66, offset: 13349},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 321, col: 75, offset: 13358},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 321, col: 94, offset: 13377},
							val:        "**",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedBoldTextUnbalancedPunctuation",
			pos:  position{line: 325, col: 1, offset: 13485},
			expr: &actionExpr{
				pos: position{line: 325, col: 42, offset: 13526},
				run: (*parser).callonEscapedBoldTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 325, col: 42, offset: 13526},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 325, col: 42, offset: 13526},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 325, col: 55, offset: 13539},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 325, col: 55, offset: 13539},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 325, col: 59, offset: 13543},
										expr: &litMatcher{
											pos:        position{line: 325, col: 59, offset: 13543},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 325, col: 65, offset: 13549},
							val:        "**",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 325, col: 70, offset: 13554},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 325, col: 79, offset: 13563},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 325, col: 98, offset: 13582},
							val:        "*",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ItalicText",
			pos:  position{line: 330, col: 1, offset: 13775},
			expr: &choiceExpr{
				pos: position{line: 330, col: 15, offset: 13789},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 330, col: 15, offset: 13789},
						name: "ItalicTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 330, col: 45, offset: 13819},
						name: "ItalicTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 330, col: 79, offset: 13853},
						name: "ItalicTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "ItalicTextSimplePunctuation",
			pos:  position{line: 332, col: 1, offset: 13882},
			expr: &actionExpr{
				pos: position{line: 332, col: 32, offset: 13913},
				run: (*parser).callonItalicTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 332, col: 32, offset: 13913},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 332, col: 32, offset: 13913},
							expr: &litMatcher{
								pos:        position{line: 332, col: 33, offset: 13914},
								val:        "\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 332, col: 37, offset: 13918},
							val:        "_",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 332, col: 41, offset: 13922},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 332, col: 50, offset: 13931},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 332, col: 69, offset: 13950},
							val:        "_",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ItalicTextDoublePunctuation",
			pos:  position{line: 336, col: 1, offset: 14029},
			expr: &actionExpr{
				pos: position{line: 336, col: 32, offset: 14060},
				run: (*parser).callonItalicTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 336, col: 32, offset: 14060},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 336, col: 32, offset: 14060},
							expr: &litMatcher{
								pos:        position{line: 336, col: 33, offset: 14061},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 336, col: 38, offset: 14066},
							val:        "__",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 336, col: 43, offset: 14071},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 336, col: 52, offset: 14080},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 336, col: 71, offset: 14099},
							val:        "__",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ItalicTextUnbalancedPunctuation",
			pos:  position{line: 340, col: 1, offset: 14179},
			expr: &actionExpr{
				pos: position{line: 340, col: 36, offset: 14214},
				run: (*parser).callonItalicTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 340, col: 36, offset: 14214},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 340, col: 36, offset: 14214},
							expr: &litMatcher{
								pos:        position{line: 340, col: 37, offset: 14215},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 340, col: 42, offset: 14220},
							val:        "__",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 340, col: 47, offset: 14225},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 340, col: 56, offset: 14234},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 340, col: 75, offset: 14253},
							val:        "_",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedItalicText",
			pos:  position{line: 345, col: 1, offset: 14419},
			expr: &choiceExpr{
				pos: position{line: 345, col: 22, offset: 14440},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 345, col: 22, offset: 14440},
						name: "EscapedItalicTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 345, col: 59, offset: 14477},
						name: "EscapedItalicTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 345, col: 100, offset: 14518},
						name: "EscapedItalicTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "EscapedItalicTextSimplePunctuation",
			pos:  position{line: 347, col: 1, offset: 14600},
			expr: &actionExpr{
				pos: position{line: 347, col: 39, offset: 14638},
				run: (*parser).callonEscapedItalicTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 347, col: 39, offset: 14638},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 347, col: 39, offset: 14638},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 347, col: 52, offset: 14651},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 347, col: 52, offset: 14651},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 347, col: 56, offset: 14655},
										expr: &litMatcher{
											pos:        position{line: 347, col: 56, offset: 14655},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 347, col: 62, offset: 14661},
							val:        "_",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 347, col: 66, offset: 14665},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 347, col: 75, offset: 14674},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 347, col: 94, offset: 14693},
							val:        "_",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedItalicTextDoublePunctuation",
			pos:  position{line: 351, col: 1, offset: 14799},
			expr: &actionExpr{
				pos: position{line: 351, col: 39, offset: 14837},
				run: (*parser).callonEscapedItalicTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 351, col: 39, offset: 14837},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 351, col: 39, offset: 14837},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 351, col: 52, offset: 14850},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 351, col: 52, offset: 14850},
										val:        "\\\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 351, col: 57, offset: 14855},
										expr: &litMatcher{
											pos:        position{line: 351, col: 57, offset: 14855},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 351, col: 63, offset: 14861},
							val:        "__",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 351, col: 68, offset: 14866},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 351, col: 77, offset: 14875},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 351, col: 96, offset: 14894},
							val:        "__",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedItalicTextUnbalancedPunctuation",
			pos:  position{line: 355, col: 1, offset: 15002},
			expr: &actionExpr{
				pos: position{line: 355, col: 44, offset: 15045},
				run: (*parser).callonEscapedItalicTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 355, col: 44, offset: 15045},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 355, col: 44, offset: 15045},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 355, col: 57, offset: 15058},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 355, col: 57, offset: 15058},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 355, col: 61, offset: 15062},
										expr: &litMatcher{
											pos:        position{line: 355, col: 61, offset: 15062},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 355, col: 67, offset: 15068},
							val:        "__",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 355, col: 72, offset: 15073},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 355, col: 81, offset: 15082},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 355, col: 100, offset: 15101},
							val:        "_",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "MonospaceText",
			pos:  position{line: 360, col: 1, offset: 15294},
			expr: &choiceExpr{
				pos: position{line: 360, col: 18, offset: 15311},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 360, col: 18, offset: 15311},
						name: "MonospaceTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 360, col: 51, offset: 15344},
						name: "MonospaceTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 360, col: 88, offset: 15381},
						name: "MonospaceTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "MonospaceTextSimplePunctuation",
			pos:  position{line: 362, col: 1, offset: 15413},
			expr: &actionExpr{
				pos: position{line: 362, col: 35, offset: 15447},
				run: (*parser).callonMonospaceTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 362, col: 35, offset: 15447},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 362, col: 35, offset: 15447},
							expr: &litMatcher{
								pos:        position{line: 362, col: 36, offset: 15448},
								val:        "\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 362, col: 40, offset: 15452},
							val:        "`",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 362, col: 44, offset: 15456},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 362, col: 53, offset: 15465},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 362, col: 72, offset: 15484},
							val:        "`",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "MonospaceTextDoublePunctuation",
			pos:  position{line: 366, col: 1, offset: 15566},
			expr: &actionExpr{
				pos: position{line: 366, col: 35, offset: 15600},
				run: (*parser).callonMonospaceTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 366, col: 35, offset: 15600},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 366, col: 35, offset: 15600},
							expr: &litMatcher{
								pos:        position{line: 366, col: 36, offset: 15601},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 366, col: 41, offset: 15606},
							val:        "``",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 366, col: 46, offset: 15611},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 366, col: 55, offset: 15620},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 366, col: 74, offset: 15639},
							val:        "``",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "MonospaceTextUnbalancedPunctuation",
			pos:  position{line: 370, col: 1, offset: 15722},
			expr: &actionExpr{
				pos: position{line: 370, col: 39, offset: 15760},
				run: (*parser).callonMonospaceTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 370, col: 39, offset: 15760},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 370, col: 39, offset: 15760},
							expr: &litMatcher{
								pos:        position{line: 370, col: 40, offset: 15761},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 370, col: 45, offset: 15766},
							val:        "``",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 370, col: 50, offset: 15771},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 370, col: 59, offset: 15780},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 370, col: 78, offset: 15799},
							val:        "`",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedMonospaceText",
			pos:  position{line: 375, col: 1, offset: 15968},
			expr: &choiceExpr{
				pos: position{line: 375, col: 25, offset: 15992},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 375, col: 25, offset: 15992},
						name: "EscapedMonospaceTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 375, col: 65, offset: 16032},
						name: "EscapedMonospaceTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 375, col: 109, offset: 16076},
						name: "EscapedMonospaceTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "EscapedMonospaceTextSimplePunctuation",
			pos:  position{line: 377, col: 1, offset: 16161},
			expr: &actionExpr{
				pos: position{line: 377, col: 42, offset: 16202},
				run: (*parser).callonEscapedMonospaceTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 377, col: 42, offset: 16202},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 377, col: 42, offset: 16202},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 377, col: 55, offset: 16215},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 377, col: 55, offset: 16215},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 377, col: 59, offset: 16219},
										expr: &litMatcher{
											pos:        position{line: 377, col: 59, offset: 16219},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 377, col: 65, offset: 16225},
							val:        "`",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 377, col: 69, offset: 16229},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 377, col: 78, offset: 16238},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 377, col: 97, offset: 16257},
							val:        "`",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedMonospaceTextDoublePunctuation",
			pos:  position{line: 381, col: 1, offset: 16363},
			expr: &actionExpr{
				pos: position{line: 381, col: 42, offset: 16404},
				run: (*parser).callonEscapedMonospaceTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 381, col: 42, offset: 16404},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 381, col: 42, offset: 16404},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 381, col: 55, offset: 16417},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 381, col: 55, offset: 16417},
										val:        "\\\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 381, col: 60, offset: 16422},
										expr: &litMatcher{
											pos:        position{line: 381, col: 60, offset: 16422},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 381, col: 66, offset: 16428},
							val:        "``",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 381, col: 71, offset: 16433},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 381, col: 80, offset: 16442},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 381, col: 99, offset: 16461},
							val:        "``",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedMonospaceTextUnbalancedPunctuation",
			pos:  position{line: 385, col: 1, offset: 16569},
			expr: &actionExpr{
				pos: position{line: 385, col: 47, offset: 16615},
				run: (*parser).callonEscapedMonospaceTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 385, col: 47, offset: 16615},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 385, col: 47, offset: 16615},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 385, col: 60, offset: 16628},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 385, col: 60, offset: 16628},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 385, col: 64, offset: 16632},
										expr: &litMatcher{
											pos:        position{line: 385, col: 64, offset: 16632},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 385, col: 70, offset: 16638},
							val:        "``",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 385, col: 75, offset: 16643},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 385, col: 84, offset: 16652},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 385, col: 103, offset: 16671},
							val:        "`",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "QuotedTextContent",
			pos:  position{line: 390, col: 1, offset: 16864},
			expr: &seqExpr{
				pos: position{line: 390, col: 22, offset: 16885},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 390, col: 22, offset: 16885},
						name: "QuotedTextContentElement",
					},
					&zeroOrMoreExpr{
						pos: position{line: 390, col: 47, offset: 16910},
						expr: &seqExpr{
							pos: position{line: 390, col: 48, offset: 16911},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 390, col: 48, offset: 16911},
									expr: &ruleRefExpr{
										pos:  position{line: 390, col: 48, offset: 16911},
										name: "WS",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 390, col: 52, offset: 16915},
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
			pos:  position{line: 392, col: 1, offset: 16943},
			expr: &choiceExpr{
				pos: position{line: 392, col: 29, offset: 16971},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 392, col: 29, offset: 16971},
						name: "QuotedText",
					},
					&ruleRefExpr{
						pos:  position{line: 392, col: 42, offset: 16984},
						name: "QuotedTextCharacters",
					},
					&ruleRefExpr{
						pos:  position{line: 392, col: 65, offset: 17007},
						name: "CharactersWithQuotePunctuation",
					},
				},
			},
		},
		{
			name: "QuotedTextCharacters",
			pos:  position{line: 394, col: 1, offset: 17142},
			expr: &oneOrMoreExpr{
				pos: position{line: 394, col: 25, offset: 17166},
				expr: &seqExpr{
					pos: position{line: 394, col: 26, offset: 17167},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 394, col: 26, offset: 17167},
							expr: &ruleRefExpr{
								pos:  position{line: 394, col: 27, offset: 17168},
								name: "NEWLINE",
							},
						},
						&notExpr{
							pos: position{line: 394, col: 35, offset: 17176},
							expr: &ruleRefExpr{
								pos:  position{line: 394, col: 36, offset: 17177},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 394, col: 39, offset: 17180},
							expr: &litMatcher{
								pos:        position{line: 394, col: 40, offset: 17181},
								val:        "*",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 394, col: 44, offset: 17185},
							expr: &litMatcher{
								pos:        position{line: 394, col: 45, offset: 17186},
								val:        "_",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 394, col: 49, offset: 17190},
							expr: &litMatcher{
								pos:        position{line: 394, col: 50, offset: 17191},
								val:        "`",
								ignoreCase: false,
							},
						},
						&anyMatcher{
							line: 394, col: 54, offset: 17195,
						},
					},
				},
			},
		},
		{
			name: "CharactersWithQuotePunctuation",
			pos:  position{line: 396, col: 1, offset: 17238},
			expr: &actionExpr{
				pos: position{line: 396, col: 35, offset: 17272},
				run: (*parser).callonCharactersWithQuotePunctuation1,
				expr: &oneOrMoreExpr{
					pos: position{line: 396, col: 35, offset: 17272},
					expr: &seqExpr{
						pos: position{line: 396, col: 36, offset: 17273},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 396, col: 36, offset: 17273},
								expr: &ruleRefExpr{
									pos:  position{line: 396, col: 37, offset: 17274},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 396, col: 45, offset: 17282},
								expr: &ruleRefExpr{
									pos:  position{line: 396, col: 46, offset: 17283},
									name: "WS",
								},
							},
							&anyMatcher{
								line: 396, col: 50, offset: 17287,
							},
						},
					},
				},
			},
		},
		{
			name: "UnbalancedQuotePunctuation",
			pos:  position{line: 401, col: 1, offset: 17532},
			expr: &choiceExpr{
				pos: position{line: 401, col: 31, offset: 17562},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 401, col: 31, offset: 17562},
						val:        "*",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 401, col: 37, offset: 17568},
						val:        "_",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 401, col: 43, offset: 17574},
						val:        "`",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "Passthrough",
			pos:  position{line: 406, col: 1, offset: 17686},
			expr: &choiceExpr{
				pos: position{line: 406, col: 16, offset: 17701},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 406, col: 16, offset: 17701},
						name: "TriplePlusPassthrough",
					},
					&ruleRefExpr{
						pos:  position{line: 406, col: 40, offset: 17725},
						name: "SinglePlusPassthrough",
					},
					&ruleRefExpr{
						pos:  position{line: 406, col: 64, offset: 17749},
						name: "PassthroughMacro",
					},
				},
			},
		},
		{
			name: "SinglePlusPassthrough",
			pos:  position{line: 408, col: 1, offset: 17767},
			expr: &actionExpr{
				pos: position{line: 408, col: 26, offset: 17792},
				run: (*parser).callonSinglePlusPassthrough1,
				expr: &seqExpr{
					pos: position{line: 408, col: 26, offset: 17792},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 408, col: 26, offset: 17792},
							val:        "+",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 408, col: 30, offset: 17796},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 408, col: 38, offset: 17804},
								expr: &seqExpr{
									pos: position{line: 408, col: 39, offset: 17805},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 408, col: 39, offset: 17805},
											expr: &ruleRefExpr{
												pos:  position{line: 408, col: 40, offset: 17806},
												name: "NEWLINE",
											},
										},
										&notExpr{
											pos: position{line: 408, col: 48, offset: 17814},
											expr: &litMatcher{
												pos:        position{line: 408, col: 49, offset: 17815},
												val:        "+",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 408, col: 53, offset: 17819,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 408, col: 57, offset: 17823},
							val:        "+",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TriplePlusPassthrough",
			pos:  position{line: 412, col: 1, offset: 17918},
			expr: &actionExpr{
				pos: position{line: 412, col: 26, offset: 17943},
				run: (*parser).callonTriplePlusPassthrough1,
				expr: &seqExpr{
					pos: position{line: 412, col: 26, offset: 17943},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 412, col: 26, offset: 17943},
							val:        "+++",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 412, col: 32, offset: 17949},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 412, col: 40, offset: 17957},
								expr: &seqExpr{
									pos: position{line: 412, col: 41, offset: 17958},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 412, col: 41, offset: 17958},
											expr: &litMatcher{
												pos:        position{line: 412, col: 42, offset: 17959},
												val:        "+++",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 412, col: 48, offset: 17965,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 412, col: 52, offset: 17969},
							val:        "+++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "PassthroughMacro",
			pos:  position{line: 416, col: 1, offset: 18066},
			expr: &choiceExpr{
				pos: position{line: 416, col: 21, offset: 18086},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 416, col: 21, offset: 18086},
						name: "SimplePassthroughMacro",
					},
					&ruleRefExpr{
						pos:  position{line: 416, col: 46, offset: 18111},
						name: "PassthroughWithQuotedText",
					},
				},
			},
		},
		{
			name: "SimplePassthroughMacro",
			pos:  position{line: 418, col: 1, offset: 18138},
			expr: &actionExpr{
				pos: position{line: 418, col: 27, offset: 18164},
				run: (*parser).callonSimplePassthroughMacro1,
				expr: &seqExpr{
					pos: position{line: 418, col: 27, offset: 18164},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 418, col: 27, offset: 18164},
							val:        "pass:[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 418, col: 36, offset: 18173},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 418, col: 44, offset: 18181},
								expr: &ruleRefExpr{
									pos:  position{line: 418, col: 45, offset: 18182},
									name: "PassthroughMacroCharacter",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 418, col: 73, offset: 18210},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "PassthroughWithQuotedText",
			pos:  position{line: 422, col: 1, offset: 18300},
			expr: &actionExpr{
				pos: position{line: 422, col: 30, offset: 18329},
				run: (*parser).callonPassthroughWithQuotedText1,
				expr: &seqExpr{
					pos: position{line: 422, col: 30, offset: 18329},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 422, col: 30, offset: 18329},
							val:        "pass:q[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 422, col: 40, offset: 18339},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 422, col: 48, offset: 18347},
								expr: &choiceExpr{
									pos: position{line: 422, col: 49, offset: 18348},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 422, col: 49, offset: 18348},
											name: "QuotedText",
										},
										&ruleRefExpr{
											pos:  position{line: 422, col: 62, offset: 18361},
											name: "PassthroughMacroCharacter",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 422, col: 90, offset: 18389},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "PassthroughMacroCharacter",
			pos:  position{line: 426, col: 1, offset: 18479},
			expr: &seqExpr{
				pos: position{line: 426, col: 31, offset: 18509},
				exprs: []interface{}{
					&notExpr{
						pos: position{line: 426, col: 31, offset: 18509},
						expr: &litMatcher{
							pos:        position{line: 426, col: 32, offset: 18510},
							val:        "]",
							ignoreCase: false,
						},
					},
					&anyMatcher{
						line: 426, col: 36, offset: 18514,
					},
				},
			},
		},
		{
			name: "CrossReference",
			pos:  position{line: 431, col: 1, offset: 18630},
			expr: &actionExpr{
				pos: position{line: 431, col: 19, offset: 18648},
				run: (*parser).callonCrossReference1,
				expr: &seqExpr{
					pos: position{line: 431, col: 19, offset: 18648},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 431, col: 19, offset: 18648},
							val:        "<<",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 431, col: 24, offset: 18653},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 431, col: 28, offset: 18657},
								name: "ID",
							},
						},
						&litMatcher{
							pos:        position{line: 431, col: 32, offset: 18661},
							val:        ">>",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ExternalLink",
			pos:  position{line: 438, col: 1, offset: 18820},
			expr: &actionExpr{
				pos: position{line: 438, col: 17, offset: 18836},
				run: (*parser).callonExternalLink1,
				expr: &seqExpr{
					pos: position{line: 438, col: 17, offset: 18836},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 438, col: 17, offset: 18836},
							label: "url",
							expr: &seqExpr{
								pos: position{line: 438, col: 22, offset: 18841},
								exprs: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 438, col: 22, offset: 18841},
										name: "URL_SCHEME",
									},
									&ruleRefExpr{
										pos:  position{line: 438, col: 33, offset: 18852},
										name: "URL",
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 438, col: 38, offset: 18857},
							label: "text",
							expr: &zeroOrOneExpr{
								pos: position{line: 438, col: 43, offset: 18862},
								expr: &seqExpr{
									pos: position{line: 438, col: 44, offset: 18863},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 438, col: 44, offset: 18863},
											val:        "[",
											ignoreCase: false,
										},
										&zeroOrMoreExpr{
											pos: position{line: 438, col: 48, offset: 18867},
											expr: &ruleRefExpr{
												pos:  position{line: 438, col: 49, offset: 18868},
												name: "URL_TEXT",
											},
										},
										&litMatcher{
											pos:        position{line: 438, col: 60, offset: 18879},
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
			pos:  position{line: 448, col: 1, offset: 19158},
			expr: &actionExpr{
				pos: position{line: 448, col: 15, offset: 19172},
				run: (*parser).callonBlockImage1,
				expr: &seqExpr{
					pos: position{line: 448, col: 15, offset: 19172},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 448, col: 15, offset: 19172},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 448, col: 26, offset: 19183},
								expr: &ruleRefExpr{
									pos:  position{line: 448, col: 27, offset: 19184},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 448, col: 46, offset: 19203},
							label: "image",
							expr: &ruleRefExpr{
								pos:  position{line: 448, col: 52, offset: 19209},
								name: "BlockImageMacro",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 448, col: 69, offset: 19226},
							expr: &ruleRefExpr{
								pos:  position{line: 448, col: 69, offset: 19226},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 448, col: 73, offset: 19230},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "BlockImageMacro",
			pos:  position{line: 453, col: 1, offset: 19391},
			expr: &actionExpr{
				pos: position{line: 453, col: 20, offset: 19410},
				run: (*parser).callonBlockImageMacro1,
				expr: &seqExpr{
					pos: position{line: 453, col: 20, offset: 19410},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 453, col: 20, offset: 19410},
							val:        "image::",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 453, col: 30, offset: 19420},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 453, col: 36, offset: 19426},
								name: "URL",
							},
						},
						&litMatcher{
							pos:        position{line: 453, col: 41, offset: 19431},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 453, col: 45, offset: 19435},
							label: "attributes",
							expr: &zeroOrOneExpr{
								pos: position{line: 453, col: 57, offset: 19447},
								expr: &ruleRefExpr{
									pos:  position{line: 453, col: 57, offset: 19447},
									name: "URL_TEXT",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 453, col: 68, offset: 19458},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "InlineImage",
			pos:  position{line: 457, col: 1, offset: 19525},
			expr: &actionExpr{
				pos: position{line: 457, col: 16, offset: 19540},
				run: (*parser).callonInlineImage1,
				expr: &labeledExpr{
					pos:   position{line: 457, col: 16, offset: 19540},
					label: "image",
					expr: &ruleRefExpr{
						pos:  position{line: 457, col: 22, offset: 19546},
						name: "InlineImageMacro",
					},
				},
			},
		},
		{
			name: "InlineImageMacro",
			pos:  position{line: 462, col: 1, offset: 19693},
			expr: &actionExpr{
				pos: position{line: 462, col: 21, offset: 19713},
				run: (*parser).callonInlineImageMacro1,
				expr: &seqExpr{
					pos: position{line: 462, col: 21, offset: 19713},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 462, col: 21, offset: 19713},
							val:        "image:",
							ignoreCase: false,
						},
						&notExpr{
							pos: position{line: 462, col: 30, offset: 19722},
							expr: &litMatcher{
								pos:        position{line: 462, col: 31, offset: 19723},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 462, col: 35, offset: 19727},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 462, col: 41, offset: 19733},
								name: "URL",
							},
						},
						&litMatcher{
							pos:        position{line: 462, col: 46, offset: 19738},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 462, col: 50, offset: 19742},
							label: "attributes",
							expr: &zeroOrOneExpr{
								pos: position{line: 462, col: 62, offset: 19754},
								expr: &ruleRefExpr{
									pos:  position{line: 462, col: 62, offset: 19754},
									name: "URL_TEXT",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 462, col: 73, offset: 19765},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "DelimitedBlock",
			pos:  position{line: 469, col: 1, offset: 20095},
			expr: &choiceExpr{
				pos: position{line: 469, col: 19, offset: 20113},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 469, col: 19, offset: 20113},
						name: "FencedBlock",
					},
					&ruleRefExpr{
						pos:  position{line: 469, col: 33, offset: 20127},
						name: "ListingBlock",
					},
				},
			},
		},
		{
			name: "BlockDelimiter",
			pos:  position{line: 471, col: 1, offset: 20142},
			expr: &choiceExpr{
				pos: position{line: 471, col: 19, offset: 20160},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 471, col: 19, offset: 20160},
						name: "FencedBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 471, col: 42, offset: 20183},
						name: "ListingBlockDelimiter",
					},
				},
			},
		},
		{
			name: "FencedBlockDelimiter",
			pos:  position{line: 473, col: 1, offset: 20206},
			expr: &litMatcher{
				pos:        position{line: 473, col: 25, offset: 20230},
				val:        "```",
				ignoreCase: false,
			},
		},
		{
			name: "FencedBlock",
			pos:  position{line: 475, col: 1, offset: 20237},
			expr: &actionExpr{
				pos: position{line: 475, col: 16, offset: 20252},
				run: (*parser).callonFencedBlock1,
				expr: &seqExpr{
					pos: position{line: 475, col: 16, offset: 20252},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 475, col: 16, offset: 20252},
							name: "FencedBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 475, col: 37, offset: 20273},
							expr: &ruleRefExpr{
								pos:  position{line: 475, col: 37, offset: 20273},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 475, col: 41, offset: 20277},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 475, col: 49, offset: 20285},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 475, col: 58, offset: 20294},
								name: "FencedBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 475, col: 78, offset: 20314},
							name: "FencedBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 475, col: 99, offset: 20335},
							expr: &ruleRefExpr{
								pos:  position{line: 475, col: 99, offset: 20335},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 475, col: 103, offset: 20339},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "FencedBlockContent",
			pos:  position{line: 479, col: 1, offset: 20427},
			expr: &labeledExpr{
				pos:   position{line: 479, col: 23, offset: 20449},
				label: "content",
				expr: &zeroOrMoreExpr{
					pos: position{line: 479, col: 31, offset: 20457},
					expr: &seqExpr{
						pos: position{line: 479, col: 32, offset: 20458},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 479, col: 32, offset: 20458},
								expr: &ruleRefExpr{
									pos:  position{line: 479, col: 33, offset: 20459},
									name: "FencedBlockDelimiter",
								},
							},
							&anyMatcher{
								line: 479, col: 54, offset: 20480,
							},
						},
					},
				},
			},
		},
		{
			name: "ListingBlockDelimiter",
			pos:  position{line: 481, col: 1, offset: 20486},
			expr: &litMatcher{
				pos:        position{line: 481, col: 26, offset: 20511},
				val:        "----",
				ignoreCase: false,
			},
		},
		{
			name: "ListingBlock",
			pos:  position{line: 483, col: 1, offset: 20519},
			expr: &actionExpr{
				pos: position{line: 483, col: 17, offset: 20535},
				run: (*parser).callonListingBlock1,
				expr: &seqExpr{
					pos: position{line: 483, col: 17, offset: 20535},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 483, col: 17, offset: 20535},
							name: "ListingBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 483, col: 39, offset: 20557},
							expr: &ruleRefExpr{
								pos:  position{line: 483, col: 39, offset: 20557},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 483, col: 43, offset: 20561},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 483, col: 51, offset: 20569},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 483, col: 60, offset: 20578},
								name: "ListingBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 483, col: 81, offset: 20599},
							name: "ListingBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 483, col: 103, offset: 20621},
							expr: &ruleRefExpr{
								pos:  position{line: 483, col: 103, offset: 20621},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 483, col: 107, offset: 20625},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ListingBlockContent",
			pos:  position{line: 487, col: 1, offset: 20714},
			expr: &labeledExpr{
				pos:   position{line: 487, col: 24, offset: 20737},
				label: "content",
				expr: &zeroOrMoreExpr{
					pos: position{line: 487, col: 32, offset: 20745},
					expr: &seqExpr{
						pos: position{line: 487, col: 33, offset: 20746},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 487, col: 33, offset: 20746},
								expr: &ruleRefExpr{
									pos:  position{line: 487, col: 34, offset: 20747},
									name: "ListingBlockDelimiter",
								},
							},
							&anyMatcher{
								line: 487, col: 56, offset: 20769,
							},
						},
					},
				},
			},
		},
		{
			name: "LiteralBlock",
			pos:  position{line: 492, col: 1, offset: 21042},
			expr: &choiceExpr{
				pos: position{line: 492, col: 17, offset: 21058},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 492, col: 17, offset: 21058},
						name: "ParagraphWithSpaces",
					},
					&ruleRefExpr{
						pos:  position{line: 492, col: 39, offset: 21080},
						name: "ParagraphWithLiteralBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 492, col: 76, offset: 21117},
						name: "ParagraphWithLiteralAttribute",
					},
				},
			},
		},
		{
			name: "ParagraphWithSpaces",
			pos:  position{line: 495, col: 1, offset: 21212},
			expr: &actionExpr{
				pos: position{line: 495, col: 24, offset: 21235},
				run: (*parser).callonParagraphWithSpaces1,
				expr: &seqExpr{
					pos: position{line: 495, col: 24, offset: 21235},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 495, col: 24, offset: 21235},
							label: "spaces",
							expr: &oneOrMoreExpr{
								pos: position{line: 495, col: 32, offset: 21243},
								expr: &ruleRefExpr{
									pos:  position{line: 495, col: 32, offset: 21243},
									name: "WS",
								},
							},
						},
						&notExpr{
							pos: position{line: 495, col: 37, offset: 21248},
							expr: &ruleRefExpr{
								pos:  position{line: 495, col: 38, offset: 21249},
								name: "NEWLINE",
							},
						},
						&labeledExpr{
							pos:   position{line: 495, col: 46, offset: 21257},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 495, col: 55, offset: 21266},
								name: "LiteralBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 495, col: 76, offset: 21287},
							name: "EndOfLiteralBlock",
						},
					},
				},
			},
		},
		{
			name: "LiteralBlockContent",
			pos:  position{line: 500, col: 1, offset: 21468},
			expr: &actionExpr{
				pos: position{line: 500, col: 24, offset: 21491},
				run: (*parser).callonLiteralBlockContent1,
				expr: &labeledExpr{
					pos:   position{line: 500, col: 24, offset: 21491},
					label: "content",
					expr: &oneOrMoreExpr{
						pos: position{line: 500, col: 32, offset: 21499},
						expr: &seqExpr{
							pos: position{line: 500, col: 33, offset: 21500},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 500, col: 33, offset: 21500},
									expr: &seqExpr{
										pos: position{line: 500, col: 35, offset: 21502},
										exprs: []interface{}{
											&ruleRefExpr{
												pos:  position{line: 500, col: 35, offset: 21502},
												name: "NEWLINE",
											},
											&ruleRefExpr{
												pos:  position{line: 500, col: 43, offset: 21510},
												name: "BlankLine",
											},
										},
									},
								},
								&anyMatcher{
									line: 500, col: 54, offset: 21521,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "EndOfLiteralBlock",
			pos:  position{line: 505, col: 1, offset: 21606},
			expr: &choiceExpr{
				pos: position{line: 505, col: 22, offset: 21627},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 505, col: 22, offset: 21627},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 505, col: 22, offset: 21627},
								name: "NEWLINE",
							},
							&ruleRefExpr{
								pos:  position{line: 505, col: 30, offset: 21635},
								name: "BlankLine",
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 505, col: 42, offset: 21647},
						name: "NEWLINE",
					},
					&ruleRefExpr{
						pos:  position{line: 505, col: 52, offset: 21657},
						name: "EOF",
					},
				},
			},
		},
		{
			name: "ParagraphWithLiteralBlockDelimiter",
			pos:  position{line: 508, col: 1, offset: 21717},
			expr: &actionExpr{
				pos: position{line: 508, col: 39, offset: 21755},
				run: (*parser).callonParagraphWithLiteralBlockDelimiter1,
				expr: &seqExpr{
					pos: position{line: 508, col: 39, offset: 21755},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 508, col: 39, offset: 21755},
							name: "LiteralBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 508, col: 61, offset: 21777},
							expr: &ruleRefExpr{
								pos:  position{line: 508, col: 61, offset: 21777},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 508, col: 65, offset: 21781},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 508, col: 73, offset: 21789},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 508, col: 81, offset: 21797},
								expr: &seqExpr{
									pos: position{line: 508, col: 82, offset: 21798},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 508, col: 82, offset: 21798},
											expr: &ruleRefExpr{
												pos:  position{line: 508, col: 83, offset: 21799},
												name: "LiteralBlockDelimiter",
											},
										},
										&anyMatcher{
											line: 508, col: 105, offset: 21821,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 508, col: 109, offset: 21825},
							name: "LiteralBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 508, col: 131, offset: 21847},
							expr: &ruleRefExpr{
								pos:  position{line: 508, col: 131, offset: 21847},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 508, col: 135, offset: 21851},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "LiteralBlockDelimiter",
			pos:  position{line: 512, col: 1, offset: 21935},
			expr: &litMatcher{
				pos:        position{line: 512, col: 26, offset: 21960},
				val:        "....",
				ignoreCase: false,
			},
		},
		{
			name: "ParagraphWithLiteralAttribute",
			pos:  position{line: 515, col: 1, offset: 22022},
			expr: &actionExpr{
				pos: position{line: 515, col: 34, offset: 22055},
				run: (*parser).callonParagraphWithLiteralAttribute1,
				expr: &seqExpr{
					pos: position{line: 515, col: 34, offset: 22055},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 515, col: 34, offset: 22055},
							val:        "[literal]",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 515, col: 46, offset: 22067},
							expr: &ruleRefExpr{
								pos:  position{line: 515, col: 46, offset: 22067},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 515, col: 50, offset: 22071},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 515, col: 58, offset: 22079},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 515, col: 67, offset: 22088},
								name: "LiteralBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 515, col: 88, offset: 22109},
							name: "EndOfLiteralBlock",
						},
					},
				},
			},
		},
		{
			name: "ElementAttribute",
			pos:  position{line: 522, col: 1, offset: 22321},
			expr: &choiceExpr{
				pos: position{line: 522, col: 21, offset: 22341},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 522, col: 21, offset: 22341},
						name: "ElementLink",
					},
					&ruleRefExpr{
						pos:  position{line: 522, col: 35, offset: 22355},
						name: "ElementID",
					},
					&ruleRefExpr{
						pos:  position{line: 522, col: 47, offset: 22367},
						name: "ElementTitle",
					},
					&ruleRefExpr{
						pos:  position{line: 522, col: 62, offset: 22382},
						name: "InvalidElementAttribute",
					},
				},
			},
		},
		{
			name: "ElementLink",
			pos:  position{line: 525, col: 1, offset: 22462},
			expr: &actionExpr{
				pos: position{line: 525, col: 16, offset: 22477},
				run: (*parser).callonElementLink1,
				expr: &seqExpr{
					pos: position{line: 525, col: 16, offset: 22477},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 525, col: 16, offset: 22477},
							val:        "[link=",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 525, col: 25, offset: 22486},
							expr: &ruleRefExpr{
								pos:  position{line: 525, col: 25, offset: 22486},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 525, col: 29, offset: 22490},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 525, col: 34, offset: 22495},
								name: "URL",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 525, col: 38, offset: 22499},
							expr: &ruleRefExpr{
								pos:  position{line: 525, col: 38, offset: 22499},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 525, col: 42, offset: 22503},
							val:        "]",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 525, col: 46, offset: 22507},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ElementID",
			pos:  position{line: 529, col: 1, offset: 22563},
			expr: &choiceExpr{
				pos: position{line: 529, col: 14, offset: 22576},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 529, col: 14, offset: 22576},
						name: "ElementIDNormal",
					},
					&ruleRefExpr{
						pos:  position{line: 529, col: 32, offset: 22594},
						name: "ElementIDShortHand",
					},
				},
			},
		},
		{
			name: "ElementIDNormal",
			pos:  position{line: 532, col: 1, offset: 22668},
			expr: &actionExpr{
				pos: position{line: 532, col: 20, offset: 22687},
				run: (*parser).callonElementIDNormal1,
				expr: &seqExpr{
					pos: position{line: 532, col: 20, offset: 22687},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 532, col: 20, offset: 22687},
							val:        "[[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 532, col: 25, offset: 22692},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 532, col: 29, offset: 22696},
								name: "ID",
							},
						},
						&litMatcher{
							pos:        position{line: 532, col: 33, offset: 22700},
							val:        "]]",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 532, col: 38, offset: 22705},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ElementIDShortHand",
			pos:  position{line: 535, col: 1, offset: 22756},
			expr: &actionExpr{
				pos: position{line: 535, col: 23, offset: 22778},
				run: (*parser).callonElementIDShortHand1,
				expr: &seqExpr{
					pos: position{line: 535, col: 23, offset: 22778},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 535, col: 23, offset: 22778},
							val:        "[#",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 535, col: 28, offset: 22783},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 535, col: 32, offset: 22787},
								name: "ID",
							},
						},
						&litMatcher{
							pos:        position{line: 535, col: 36, offset: 22791},
							val:        "]",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 535, col: 40, offset: 22795},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ElementTitle",
			pos:  position{line: 541, col: 1, offset: 22989},
			expr: &actionExpr{
				pos: position{line: 541, col: 17, offset: 23005},
				run: (*parser).callonElementTitle1,
				expr: &seqExpr{
					pos: position{line: 541, col: 17, offset: 23005},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 541, col: 17, offset: 23005},
							val:        ".",
							ignoreCase: false,
						},
						&notExpr{
							pos: position{line: 541, col: 21, offset: 23009},
							expr: &litMatcher{
								pos:        position{line: 541, col: 22, offset: 23010},
								val:        ".",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 541, col: 26, offset: 23014},
							expr: &ruleRefExpr{
								pos:  position{line: 541, col: 27, offset: 23015},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 541, col: 30, offset: 23018},
							label: "title",
							expr: &oneOrMoreExpr{
								pos: position{line: 541, col: 36, offset: 23024},
								expr: &seqExpr{
									pos: position{line: 541, col: 37, offset: 23025},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 541, col: 37, offset: 23025},
											expr: &ruleRefExpr{
												pos:  position{line: 541, col: 38, offset: 23026},
												name: "NEWLINE",
											},
										},
										&anyMatcher{
											line: 541, col: 46, offset: 23034,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 541, col: 50, offset: 23038},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "InvalidElementAttribute",
			pos:  position{line: 545, col: 1, offset: 23103},
			expr: &actionExpr{
				pos: position{line: 545, col: 28, offset: 23130},
				run: (*parser).callonInvalidElementAttribute1,
				expr: &seqExpr{
					pos: position{line: 545, col: 28, offset: 23130},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 545, col: 28, offset: 23130},
							val:        "[",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 545, col: 32, offset: 23134},
							expr: &ruleRefExpr{
								pos:  position{line: 545, col: 32, offset: 23134},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 545, col: 36, offset: 23138},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 545, col: 44, offset: 23146},
								expr: &seqExpr{
									pos: position{line: 545, col: 45, offset: 23147},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 545, col: 45, offset: 23147},
											expr: &litMatcher{
												pos:        position{line: 545, col: 46, offset: 23148},
												val:        "]",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 545, col: 50, offset: 23152,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 545, col: 54, offset: 23156},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "BlankLine",
			pos:  position{line: 552, col: 1, offset: 23322},
			expr: &actionExpr{
				pos: position{line: 552, col: 14, offset: 23335},
				run: (*parser).callonBlankLine1,
				expr: &seqExpr{
					pos: position{line: 552, col: 14, offset: 23335},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 552, col: 14, offset: 23335},
							expr: &ruleRefExpr{
								pos:  position{line: 552, col: 15, offset: 23336},
								name: "EOF",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 552, col: 19, offset: 23340},
							expr: &ruleRefExpr{
								pos:  position{line: 552, col: 19, offset: 23340},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 552, col: 23, offset: 23344},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "Characters",
			pos:  position{line: 559, col: 1, offset: 23491},
			expr: &actionExpr{
				pos: position{line: 559, col: 15, offset: 23505},
				run: (*parser).callonCharacters1,
				expr: &oneOrMoreExpr{
					pos: position{line: 559, col: 15, offset: 23505},
					expr: &seqExpr{
						pos: position{line: 559, col: 16, offset: 23506},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 559, col: 16, offset: 23506},
								expr: &ruleRefExpr{
									pos:  position{line: 559, col: 17, offset: 23507},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 559, col: 25, offset: 23515},
								expr: &ruleRefExpr{
									pos:  position{line: 559, col: 26, offset: 23516},
									name: "WS",
								},
							},
							&anyMatcher{
								line: 559, col: 29, offset: 23519,
							},
						},
					},
				},
			},
		},
		{
			name: "URL",
			pos:  position{line: 564, col: 1, offset: 23560},
			expr: &actionExpr{
				pos: position{line: 564, col: 8, offset: 23567},
				run: (*parser).callonURL1,
				expr: &oneOrMoreExpr{
					pos: position{line: 564, col: 8, offset: 23567},
					expr: &seqExpr{
						pos: position{line: 564, col: 9, offset: 23568},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 564, col: 9, offset: 23568},
								expr: &ruleRefExpr{
									pos:  position{line: 564, col: 10, offset: 23569},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 564, col: 18, offset: 23577},
								expr: &ruleRefExpr{
									pos:  position{line: 564, col: 19, offset: 23578},
									name: "WS",
								},
							},
							&notExpr{
								pos: position{line: 564, col: 22, offset: 23581},
								expr: &litMatcher{
									pos:        position{line: 564, col: 23, offset: 23582},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 564, col: 27, offset: 23586},
								expr: &litMatcher{
									pos:        position{line: 564, col: 28, offset: 23587},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 564, col: 32, offset: 23591,
							},
						},
					},
				},
			},
		},
		{
			name: "ID",
			pos:  position{line: 568, col: 1, offset: 23631},
			expr: &actionExpr{
				pos: position{line: 568, col: 7, offset: 23637},
				run: (*parser).callonID1,
				expr: &oneOrMoreExpr{
					pos: position{line: 568, col: 7, offset: 23637},
					expr: &seqExpr{
						pos: position{line: 568, col: 8, offset: 23638},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 568, col: 8, offset: 23638},
								expr: &ruleRefExpr{
									pos:  position{line: 568, col: 9, offset: 23639},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 568, col: 17, offset: 23647},
								expr: &ruleRefExpr{
									pos:  position{line: 568, col: 18, offset: 23648},
									name: "WS",
								},
							},
							&notExpr{
								pos: position{line: 568, col: 21, offset: 23651},
								expr: &litMatcher{
									pos:        position{line: 568, col: 22, offset: 23652},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 568, col: 26, offset: 23656},
								expr: &litMatcher{
									pos:        position{line: 568, col: 27, offset: 23657},
									val:        "]",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 568, col: 31, offset: 23661},
								expr: &litMatcher{
									pos:        position{line: 568, col: 32, offset: 23662},
									val:        "<<",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 568, col: 37, offset: 23667},
								expr: &litMatcher{
									pos:        position{line: 568, col: 38, offset: 23668},
									val:        ">>",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 568, col: 42, offset: 23672,
							},
						},
					},
				},
			},
		},
		{
			name: "URL_TEXT",
			pos:  position{line: 572, col: 1, offset: 23712},
			expr: &actionExpr{
				pos: position{line: 572, col: 13, offset: 23724},
				run: (*parser).callonURL_TEXT1,
				expr: &oneOrMoreExpr{
					pos: position{line: 572, col: 13, offset: 23724},
					expr: &seqExpr{
						pos: position{line: 572, col: 14, offset: 23725},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 572, col: 14, offset: 23725},
								expr: &ruleRefExpr{
									pos:  position{line: 572, col: 15, offset: 23726},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 572, col: 23, offset: 23734},
								expr: &litMatcher{
									pos:        position{line: 572, col: 24, offset: 23735},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 572, col: 28, offset: 23739},
								expr: &litMatcher{
									pos:        position{line: 572, col: 29, offset: 23740},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 572, col: 33, offset: 23744,
							},
						},
					},
				},
			},
		},
		{
			name: "URL_SCHEME",
			pos:  position{line: 576, col: 1, offset: 23784},
			expr: &choiceExpr{
				pos: position{line: 576, col: 15, offset: 23798},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 576, col: 15, offset: 23798},
						val:        "http://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 576, col: 27, offset: 23810},
						val:        "https://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 576, col: 40, offset: 23823},
						val:        "ftp://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 576, col: 51, offset: 23834},
						val:        "irc://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 576, col: 62, offset: 23845},
						val:        "mailto:",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "DIGIT",
			pos:  position{line: 578, col: 1, offset: 23856},
			expr: &charClassMatcher{
				pos:        position{line: 578, col: 10, offset: 23865},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NEWLINE",
			pos:  position{line: 580, col: 1, offset: 23872},
			expr: &choiceExpr{
				pos: position{line: 580, col: 12, offset: 23883},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 580, col: 12, offset: 23883},
						val:        "\r\n",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 580, col: 21, offset: 23892},
						val:        "\r",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 580, col: 28, offset: 23899},
						val:        "\n",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "WS",
			pos:  position{line: 582, col: 1, offset: 23905},
			expr: &choiceExpr{
				pos: position{line: 582, col: 7, offset: 23911},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 582, col: 7, offset: 23911},
						val:        " ",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 582, col: 13, offset: 23917},
						run: (*parser).callonWS3,
						expr: &litMatcher{
							pos:        position{line: 582, col: 13, offset: 23917},
							val:        "\t",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 586, col: 1, offset: 23962},
			expr: &notExpr{
				pos: position{line: 586, col: 8, offset: 23969},
				expr: &anyMatcher{
					line: 586, col: 9, offset: 23970,
				},
			},
		},
		{
			name: "EOL",
			pos:  position{line: 588, col: 1, offset: 23973},
			expr: &choiceExpr{
				pos: position{line: 588, col: 8, offset: 23980},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 588, col: 8, offset: 23980},
						name: "NEWLINE",
					},
					&ruleRefExpr{
						pos:  position{line: 588, col: 18, offset: 23990},
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

func (c *current) onListAttribute1(attribute interface{}) (interface{}, error) {
	return attribute, nil
}

func (p *parser) callonListAttribute1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onListAttribute1(stack["attribute"])
}

func (c *current) onListID1(id interface{}) (interface{}, error) {
	return map[string]interface{}{"ID": id.(string)}, nil
}

func (p *parser) callonListID1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onListID1(stack["id"])
}

func (c *current) onHorizontalLayout1() (interface{}, error) {
	return map[string]interface{}{"layout": "horizontal"}, nil
}

func (p *parser) callonHorizontalLayout1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onHorizontalLayout1()
}

func (c *current) onListParagraph1(lines interface{}) (interface{}, error) {
	return types.NewListParagraph(lines.([]interface{}))
}

func (p *parser) callonListParagraph1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onListParagraph1(stack["lines"])
}

func (c *current) onUnorderedListItem1(level, content interface{}) (interface{}, error) {
	return types.NewUnorderedListItem(level, content.([]types.DocElement))
}

func (p *parser) callonUnorderedListItem1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUnorderedListItem1(stack["level"], stack["content"])
}

func (c *current) onUnorderedListItemPrefix1(level interface{}) (interface{}, error) {
	// ignore whitespaces, only return the relevant "*"/"-" characters
	return level, nil
}

func (p *parser) callonUnorderedListItemPrefix1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUnorderedListItemPrefix1(stack["level"])
}

func (c *current) onUnorderedListItemContent1(elements interface{}) (interface{}, error) {
	// Another list or a literal paragraph immediately following a list item will be implicitly included in the list item
	return types.NewListItemContent(elements.([]interface{}))
}

func (p *parser) callonUnorderedListItemContent1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUnorderedListItemContent1(stack["elements"])
}

func (c *current) onLabeledListItemWithTermAlone1(term interface{}) (interface{}, error) {
	// here, WS is optional since there is no description afterwards
	return types.NewLabeledListItem(term.([]interface{}), nil)
}

func (p *parser) callonLabeledListItemWithTermAlone1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLabeledListItemWithTermAlone1(stack["term"])
}

func (c *current) onLabeledListItemTerm1(term interface{}) (interface{}, error) {
	return term, nil
}

func (p *parser) callonLabeledListItemTerm1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLabeledListItemTerm1(stack["term"])
}

func (c *current) onLabeledListItemWithDescription1(term, description interface{}) (interface{}, error) {
	return types.NewLabeledListItem(term.([]interface{}), description.([]types.DocElement))
}

func (p *parser) callonLabeledListItemWithDescription1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLabeledListItemWithDescription1(stack["term"], stack["description"])
}

func (c *current) onLabeledListItemDescription1(elements interface{}) (interface{}, error) {

	return types.NewListItemContent(elements.([]interface{}))
}

func (p *parser) callonLabeledListItemDescription1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLabeledListItemDescription1(stack["elements"])
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

func (c *current) onCrossReference1(id interface{}) (interface{}, error) {
	return types.NewCrossReference(id.(string))
}

func (p *parser) callonCrossReference1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCrossReference1(stack["id"])
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

func (c *current) onElementIDNormal1(id interface{}) (interface{}, error) {
	return types.NewElementID(id.(string))
}

func (p *parser) callonElementIDNormal1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onElementIDNormal1(stack["id"])
}

func (c *current) onElementIDShortHand1(id interface{}) (interface{}, error) {
	return types.NewElementID(id.(string))
}

func (p *parser) callonElementIDShortHand1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onElementIDShortHand1(stack["id"])
}

func (c *current) onElementTitle1(title interface{}) (interface{}, error) {
	return types.NewElementTitle(title.([]interface{}))
}

func (p *parser) callonElementTitle1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onElementTitle1(stack["title"])
}

func (c *current) onInvalidElementAttribute1(content interface{}) (interface{}, error) {
	return types.NewInvalidElementAttribute(c.text)
}

func (p *parser) callonInvalidElementAttribute1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInvalidElementAttribute1(stack["content"])
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
