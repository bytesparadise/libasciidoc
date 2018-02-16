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
										name: "ListItemContinuation",
									},
								},
								&notExpr{
									pos: position{line: 221, col: 50, offset: 8755},
									expr: &ruleRefExpr{
										pos:  position{line: 221, col: 52, offset: 8757},
										name: "UnorderedListItemPrefix",
									},
								},
								&notExpr{
									pos: position{line: 221, col: 77, offset: 8782},
									expr: &seqExpr{
										pos: position{line: 221, col: 79, offset: 8784},
										exprs: []interface{}{
											&ruleRefExpr{
												pos:  position{line: 221, col: 79, offset: 8784},
												name: "LabeledListItemTerm",
											},
											&ruleRefExpr{
												pos:  position{line: 221, col: 99, offset: 8804},
												name: "LabeledListItemSeparator",
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 221, col: 125, offset: 8830},
									name: "InlineContent",
								},
								&ruleRefExpr{
									pos:  position{line: 221, col: 139, offset: 8844},
									name: "EOL",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "ListItemContinuation",
			pos:  position{line: 225, col: 1, offset: 8913},
			expr: &actionExpr{
				pos: position{line: 225, col: 25, offset: 8937},
				run: (*parser).callonListItemContinuation1,
				expr: &seqExpr{
					pos: position{line: 225, col: 25, offset: 8937},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 225, col: 25, offset: 8937},
							val:        "+",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 225, col: 29, offset: 8941},
							expr: &ruleRefExpr{
								pos:  position{line: 225, col: 29, offset: 8941},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 225, col: 33, offset: 8945},
							name: "NEWLINE",
						},
					},
				},
			},
		},
		{
			name: "ContinuedBlockElement",
			pos:  position{line: 229, col: 1, offset: 9001},
			expr: &actionExpr{
				pos: position{line: 229, col: 26, offset: 9026},
				run: (*parser).callonContinuedBlockElement1,
				expr: &seqExpr{
					pos: position{line: 229, col: 26, offset: 9026},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 229, col: 26, offset: 9026},
							name: "ListItemContinuation",
						},
						&labeledExpr{
							pos:   position{line: 229, col: 47, offset: 9047},
							label: "element",
							expr: &ruleRefExpr{
								pos:  position{line: 229, col: 55, offset: 9055},
								name: "BlockElement",
							},
						},
					},
				},
			},
		},
		{
			name: "UnorderedListItem",
			pos:  position{line: 243, col: 1, offset: 9468},
			expr: &actionExpr{
				pos: position{line: 243, col: 22, offset: 9489},
				run: (*parser).callonUnorderedListItem1,
				expr: &seqExpr{
					pos: position{line: 243, col: 22, offset: 9489},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 243, col: 22, offset: 9489},
							label: "level",
							expr: &ruleRefExpr{
								pos:  position{line: 243, col: 29, offset: 9496},
								name: "UnorderedListItemPrefix",
							},
						},
						&labeledExpr{
							pos:   position{line: 243, col: 54, offset: 9521},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 243, col: 63, offset: 9530},
								name: "UnorderedListItemContent",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 243, col: 89, offset: 9556},
							expr: &ruleRefExpr{
								pos:  position{line: 243, col: 89, offset: 9556},
								name: "BlankLine",
							},
						},
					},
				},
			},
		},
		{
			name: "UnorderedListItemPrefix",
			pos:  position{line: 247, col: 1, offset: 9647},
			expr: &actionExpr{
				pos: position{line: 247, col: 28, offset: 9674},
				run: (*parser).callonUnorderedListItemPrefix1,
				expr: &seqExpr{
					pos: position{line: 247, col: 28, offset: 9674},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 247, col: 28, offset: 9674},
							expr: &ruleRefExpr{
								pos:  position{line: 247, col: 28, offset: 9674},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 247, col: 32, offset: 9678},
							label: "level",
							expr: &choiceExpr{
								pos: position{line: 247, col: 39, offset: 9685},
								alternatives: []interface{}{
									&oneOrMoreExpr{
										pos: position{line: 247, col: 39, offset: 9685},
										expr: &litMatcher{
											pos:        position{line: 247, col: 39, offset: 9685},
											val:        "*",
											ignoreCase: false,
										},
									},
									&litMatcher{
										pos:        position{line: 247, col: 46, offset: 9692},
										val:        "-",
										ignoreCase: false,
									},
								},
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 247, col: 51, offset: 9697},
							expr: &ruleRefExpr{
								pos:  position{line: 247, col: 51, offset: 9697},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "UnorderedListItemContent",
			pos:  position{line: 251, col: 1, offset: 9795},
			expr: &actionExpr{
				pos: position{line: 251, col: 29, offset: 9823},
				run: (*parser).callonUnorderedListItemContent1,
				expr: &labeledExpr{
					pos:   position{line: 251, col: 29, offset: 9823},
					label: "elements",
					expr: &seqExpr{
						pos: position{line: 251, col: 39, offset: 9833},
						exprs: []interface{}{
							&oneOrMoreExpr{
								pos: position{line: 251, col: 39, offset: 9833},
								expr: &ruleRefExpr{
									pos:  position{line: 251, col: 39, offset: 9833},
									name: "ListParagraph",
								},
							},
							&zeroOrMoreExpr{
								pos: position{line: 251, col: 54, offset: 9848},
								expr: &ruleRefExpr{
									pos:  position{line: 251, col: 54, offset: 9848},
									name: "ContinuedBlockElement",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "LabeledListItem",
			pos:  position{line: 264, col: 1, offset: 10343},
			expr: &choiceExpr{
				pos: position{line: 264, col: 20, offset: 10362},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 264, col: 20, offset: 10362},
						name: "LabeledListItemWithDescription",
					},
					&ruleRefExpr{
						pos:  position{line: 264, col: 53, offset: 10395},
						name: "LabeledListItemWithTermAlone",
					},
				},
			},
		},
		{
			name: "LabeledListItemWithTermAlone",
			pos:  position{line: 266, col: 1, offset: 10425},
			expr: &actionExpr{
				pos: position{line: 266, col: 33, offset: 10457},
				run: (*parser).callonLabeledListItemWithTermAlone1,
				expr: &seqExpr{
					pos: position{line: 266, col: 33, offset: 10457},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 266, col: 33, offset: 10457},
							label: "term",
							expr: &ruleRefExpr{
								pos:  position{line: 266, col: 39, offset: 10463},
								name: "LabeledListItemTerm",
							},
						},
						&litMatcher{
							pos:        position{line: 266, col: 61, offset: 10485},
							val:        "::",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 266, col: 66, offset: 10490},
							expr: &ruleRefExpr{
								pos:  position{line: 266, col: 66, offset: 10490},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 266, col: 70, offset: 10494},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "LabeledListItemTerm",
			pos:  position{line: 270, col: 1, offset: 10631},
			expr: &actionExpr{
				pos: position{line: 270, col: 24, offset: 10654},
				run: (*parser).callonLabeledListItemTerm1,
				expr: &labeledExpr{
					pos:   position{line: 270, col: 24, offset: 10654},
					label: "term",
					expr: &zeroOrMoreExpr{
						pos: position{line: 270, col: 29, offset: 10659},
						expr: &seqExpr{
							pos: position{line: 270, col: 30, offset: 10660},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 270, col: 30, offset: 10660},
									expr: &ruleRefExpr{
										pos:  position{line: 270, col: 31, offset: 10661},
										name: "NEWLINE",
									},
								},
								&notExpr{
									pos: position{line: 270, col: 39, offset: 10669},
									expr: &litMatcher{
										pos:        position{line: 270, col: 40, offset: 10670},
										val:        "::",
										ignoreCase: false,
									},
								},
								&anyMatcher{
									line: 270, col: 45, offset: 10675,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "LabeledListItemSeparator",
			pos:  position{line: 275, col: 1, offset: 10766},
			expr: &seqExpr{
				pos: position{line: 275, col: 30, offset: 10795},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 275, col: 30, offset: 10795},
						val:        "::",
						ignoreCase: false,
					},
					&oneOrMoreExpr{
						pos: position{line: 275, col: 35, offset: 10800},
						expr: &choiceExpr{
							pos: position{line: 275, col: 36, offset: 10801},
							alternatives: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 275, col: 36, offset: 10801},
									name: "WS",
								},
								&ruleRefExpr{
									pos:  position{line: 275, col: 41, offset: 10806},
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
			pos:  position{line: 277, col: 1, offset: 10817},
			expr: &actionExpr{
				pos: position{line: 277, col: 35, offset: 10851},
				run: (*parser).callonLabeledListItemWithDescription1,
				expr: &seqExpr{
					pos: position{line: 277, col: 35, offset: 10851},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 277, col: 35, offset: 10851},
							label: "term",
							expr: &ruleRefExpr{
								pos:  position{line: 277, col: 41, offset: 10857},
								name: "LabeledListItemTerm",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 277, col: 62, offset: 10878},
							name: "LabeledListItemSeparator",
						},
						&labeledExpr{
							pos:   position{line: 277, col: 87, offset: 10903},
							label: "description",
							expr: &ruleRefExpr{
								pos:  position{line: 277, col: 100, offset: 10916},
								name: "LabeledListItemDescription",
							},
						},
					},
				},
			},
		},
		{
			name: "LabeledListItemDescription",
			pos:  position{line: 281, col: 1, offset: 11041},
			expr: &actionExpr{
				pos: position{line: 281, col: 31, offset: 11071},
				run: (*parser).callonLabeledListItemDescription1,
				expr: &labeledExpr{
					pos:   position{line: 281, col: 31, offset: 11071},
					label: "elements",
					expr: &zeroOrMoreExpr{
						pos: position{line: 281, col: 40, offset: 11080},
						expr: &choiceExpr{
							pos: position{line: 281, col: 41, offset: 11081},
							alternatives: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 281, col: 41, offset: 11081},
									name: "ListParagraph",
								},
								&ruleRefExpr{
									pos:  position{line: 281, col: 57, offset: 11097},
									name: "ContinuedBlockElement",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Paragraph",
			pos:  position{line: 290, col: 1, offset: 11447},
			expr: &actionExpr{
				pos: position{line: 290, col: 14, offset: 11460},
				run: (*parser).callonParagraph1,
				expr: &seqExpr{
					pos: position{line: 290, col: 14, offset: 11460},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 290, col: 14, offset: 11460},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 290, col: 25, offset: 11471},
								expr: &ruleRefExpr{
									pos:  position{line: 290, col: 26, offset: 11472},
									name: "ElementAttribute",
								},
							},
						},
						&notExpr{
							pos: position{line: 290, col: 45, offset: 11491},
							expr: &seqExpr{
								pos: position{line: 290, col: 47, offset: 11493},
								exprs: []interface{}{
									&oneOrMoreExpr{
										pos: position{line: 290, col: 47, offset: 11493},
										expr: &litMatcher{
											pos:        position{line: 290, col: 47, offset: 11493},
											val:        "=",
											ignoreCase: false,
										},
									},
									&oneOrMoreExpr{
										pos: position{line: 290, col: 52, offset: 11498},
										expr: &ruleRefExpr{
											pos:  position{line: 290, col: 52, offset: 11498},
											name: "WS",
										},
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 290, col: 57, offset: 11503},
							label: "lines",
							expr: &oneOrMoreExpr{
								pos: position{line: 290, col: 63, offset: 11509},
								expr: &seqExpr{
									pos: position{line: 290, col: 64, offset: 11510},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 290, col: 64, offset: 11510},
											name: "InlineContent",
										},
										&ruleRefExpr{
											pos:  position{line: 290, col: 78, offset: 11524},
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
			pos:  position{line: 296, col: 1, offset: 11814},
			expr: &actionExpr{
				pos: position{line: 296, col: 18, offset: 11831},
				run: (*parser).callonInlineContent1,
				expr: &seqExpr{
					pos: position{line: 296, col: 18, offset: 11831},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 296, col: 18, offset: 11831},
							expr: &ruleRefExpr{
								pos:  position{line: 296, col: 19, offset: 11832},
								name: "BlockDelimiter",
							},
						},
						&labeledExpr{
							pos:   position{line: 296, col: 34, offset: 11847},
							label: "elements",
							expr: &oneOrMoreExpr{
								pos: position{line: 296, col: 43, offset: 11856},
								expr: &seqExpr{
									pos: position{line: 296, col: 44, offset: 11857},
									exprs: []interface{}{
										&zeroOrMoreExpr{
											pos: position{line: 296, col: 44, offset: 11857},
											expr: &ruleRefExpr{
												pos:  position{line: 296, col: 44, offset: 11857},
												name: "WS",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 296, col: 48, offset: 11861},
											name: "InlineElement",
										},
										&zeroOrMoreExpr{
											pos: position{line: 296, col: 62, offset: 11875},
											expr: &ruleRefExpr{
												pos:  position{line: 296, col: 62, offset: 11875},
												name: "WS",
											},
										},
									},
								},
							},
						},
						&andExpr{
							pos: position{line: 296, col: 68, offset: 11881},
							expr: &ruleRefExpr{
								pos:  position{line: 296, col: 69, offset: 11882},
								name: "EOL",
							},
						},
					},
				},
			},
		},
		{
			name: "InlineElement",
			pos:  position{line: 300, col: 1, offset: 12000},
			expr: &choiceExpr{
				pos: position{line: 300, col: 18, offset: 12017},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 300, col: 18, offset: 12017},
						name: "CrossReference",
					},
					&ruleRefExpr{
						pos:  position{line: 300, col: 35, offset: 12034},
						name: "Passthrough",
					},
					&ruleRefExpr{
						pos:  position{line: 300, col: 49, offset: 12048},
						name: "InlineImage",
					},
					&ruleRefExpr{
						pos:  position{line: 300, col: 63, offset: 12062},
						name: "QuotedText",
					},
					&ruleRefExpr{
						pos:  position{line: 300, col: 76, offset: 12075},
						name: "ExternalLink",
					},
					&ruleRefExpr{
						pos:  position{line: 300, col: 91, offset: 12090},
						name: "DocumentAttributeSubstitution",
					},
					&ruleRefExpr{
						pos:  position{line: 300, col: 123, offset: 12122},
						name: "Characters",
					},
				},
			},
		},
		{
			name: "QuotedText",
			pos:  position{line: 305, col: 1, offset: 12373},
			expr: &choiceExpr{
				pos: position{line: 305, col: 15, offset: 12387},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 305, col: 15, offset: 12387},
						name: "BoldText",
					},
					&ruleRefExpr{
						pos:  position{line: 305, col: 26, offset: 12398},
						name: "ItalicText",
					},
					&ruleRefExpr{
						pos:  position{line: 305, col: 39, offset: 12411},
						name: "MonospaceText",
					},
					&ruleRefExpr{
						pos:  position{line: 306, col: 13, offset: 12439},
						name: "EscapedBoldText",
					},
					&ruleRefExpr{
						pos:  position{line: 306, col: 31, offset: 12457},
						name: "EscapedItalicText",
					},
					&ruleRefExpr{
						pos:  position{line: 306, col: 51, offset: 12477},
						name: "EscapedMonospaceText",
					},
				},
			},
		},
		{
			name: "BoldText",
			pos:  position{line: 308, col: 1, offset: 12499},
			expr: &choiceExpr{
				pos: position{line: 308, col: 13, offset: 12511},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 308, col: 13, offset: 12511},
						name: "BoldTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 308, col: 41, offset: 12539},
						name: "BoldTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 308, col: 73, offset: 12571},
						name: "BoldTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "BoldTextSimplePunctuation",
			pos:  position{line: 310, col: 1, offset: 12644},
			expr: &actionExpr{
				pos: position{line: 310, col: 30, offset: 12673},
				run: (*parser).callonBoldTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 310, col: 30, offset: 12673},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 310, col: 30, offset: 12673},
							expr: &litMatcher{
								pos:        position{line: 310, col: 31, offset: 12674},
								val:        "\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 310, col: 35, offset: 12678},
							val:        "*",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 310, col: 39, offset: 12682},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 310, col: 48, offset: 12691},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 310, col: 67, offset: 12710},
							val:        "*",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "BoldTextDoublePunctuation",
			pos:  position{line: 314, col: 1, offset: 12787},
			expr: &actionExpr{
				pos: position{line: 314, col: 30, offset: 12816},
				run: (*parser).callonBoldTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 314, col: 30, offset: 12816},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 314, col: 30, offset: 12816},
							expr: &litMatcher{
								pos:        position{line: 314, col: 31, offset: 12817},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 314, col: 36, offset: 12822},
							val:        "**",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 314, col: 41, offset: 12827},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 314, col: 50, offset: 12836},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 314, col: 69, offset: 12855},
							val:        "**",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "BoldTextUnbalancedPunctuation",
			pos:  position{line: 318, col: 1, offset: 12933},
			expr: &actionExpr{
				pos: position{line: 318, col: 34, offset: 12966},
				run: (*parser).callonBoldTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 318, col: 34, offset: 12966},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 318, col: 34, offset: 12966},
							expr: &litMatcher{
								pos:        position{line: 318, col: 35, offset: 12967},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 318, col: 40, offset: 12972},
							val:        "**",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 318, col: 45, offset: 12977},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 318, col: 54, offset: 12986},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 318, col: 73, offset: 13005},
							val:        "*",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedBoldText",
			pos:  position{line: 323, col: 1, offset: 13169},
			expr: &choiceExpr{
				pos: position{line: 323, col: 20, offset: 13188},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 323, col: 20, offset: 13188},
						name: "EscapedBoldTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 323, col: 55, offset: 13223},
						name: "EscapedBoldTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 323, col: 94, offset: 13262},
						name: "EscapedBoldTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "EscapedBoldTextSimplePunctuation",
			pos:  position{line: 325, col: 1, offset: 13342},
			expr: &actionExpr{
				pos: position{line: 325, col: 37, offset: 13378},
				run: (*parser).callonEscapedBoldTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 325, col: 37, offset: 13378},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 325, col: 37, offset: 13378},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 325, col: 50, offset: 13391},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 325, col: 50, offset: 13391},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 325, col: 54, offset: 13395},
										expr: &litMatcher{
											pos:        position{line: 325, col: 54, offset: 13395},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 325, col: 60, offset: 13401},
							val:        "*",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 325, col: 64, offset: 13405},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 325, col: 73, offset: 13414},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 325, col: 92, offset: 13433},
							val:        "*",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedBoldTextDoublePunctuation",
			pos:  position{line: 329, col: 1, offset: 13539},
			expr: &actionExpr{
				pos: position{line: 329, col: 37, offset: 13575},
				run: (*parser).callonEscapedBoldTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 329, col: 37, offset: 13575},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 329, col: 37, offset: 13575},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 329, col: 50, offset: 13588},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 329, col: 50, offset: 13588},
										val:        "\\\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 329, col: 55, offset: 13593},
										expr: &litMatcher{
											pos:        position{line: 329, col: 55, offset: 13593},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 329, col: 61, offset: 13599},
							val:        "**",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 329, col: 66, offset: 13604},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 329, col: 75, offset: 13613},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 329, col: 94, offset: 13632},
							val:        "**",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedBoldTextUnbalancedPunctuation",
			pos:  position{line: 333, col: 1, offset: 13740},
			expr: &actionExpr{
				pos: position{line: 333, col: 42, offset: 13781},
				run: (*parser).callonEscapedBoldTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 333, col: 42, offset: 13781},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 333, col: 42, offset: 13781},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 333, col: 55, offset: 13794},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 333, col: 55, offset: 13794},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 333, col: 59, offset: 13798},
										expr: &litMatcher{
											pos:        position{line: 333, col: 59, offset: 13798},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 333, col: 65, offset: 13804},
							val:        "**",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 333, col: 70, offset: 13809},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 333, col: 79, offset: 13818},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 333, col: 98, offset: 13837},
							val:        "*",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ItalicText",
			pos:  position{line: 338, col: 1, offset: 14030},
			expr: &choiceExpr{
				pos: position{line: 338, col: 15, offset: 14044},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 338, col: 15, offset: 14044},
						name: "ItalicTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 338, col: 45, offset: 14074},
						name: "ItalicTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 338, col: 79, offset: 14108},
						name: "ItalicTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "ItalicTextSimplePunctuation",
			pos:  position{line: 340, col: 1, offset: 14137},
			expr: &actionExpr{
				pos: position{line: 340, col: 32, offset: 14168},
				run: (*parser).callonItalicTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 340, col: 32, offset: 14168},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 340, col: 32, offset: 14168},
							expr: &litMatcher{
								pos:        position{line: 340, col: 33, offset: 14169},
								val:        "\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 340, col: 37, offset: 14173},
							val:        "_",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 340, col: 41, offset: 14177},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 340, col: 50, offset: 14186},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 340, col: 69, offset: 14205},
							val:        "_",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ItalicTextDoublePunctuation",
			pos:  position{line: 344, col: 1, offset: 14284},
			expr: &actionExpr{
				pos: position{line: 344, col: 32, offset: 14315},
				run: (*parser).callonItalicTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 344, col: 32, offset: 14315},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 344, col: 32, offset: 14315},
							expr: &litMatcher{
								pos:        position{line: 344, col: 33, offset: 14316},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 344, col: 38, offset: 14321},
							val:        "__",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 344, col: 43, offset: 14326},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 344, col: 52, offset: 14335},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 344, col: 71, offset: 14354},
							val:        "__",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ItalicTextUnbalancedPunctuation",
			pos:  position{line: 348, col: 1, offset: 14434},
			expr: &actionExpr{
				pos: position{line: 348, col: 36, offset: 14469},
				run: (*parser).callonItalicTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 348, col: 36, offset: 14469},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 348, col: 36, offset: 14469},
							expr: &litMatcher{
								pos:        position{line: 348, col: 37, offset: 14470},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 348, col: 42, offset: 14475},
							val:        "__",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 348, col: 47, offset: 14480},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 348, col: 56, offset: 14489},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 348, col: 75, offset: 14508},
							val:        "_",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedItalicText",
			pos:  position{line: 353, col: 1, offset: 14674},
			expr: &choiceExpr{
				pos: position{line: 353, col: 22, offset: 14695},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 353, col: 22, offset: 14695},
						name: "EscapedItalicTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 353, col: 59, offset: 14732},
						name: "EscapedItalicTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 353, col: 100, offset: 14773},
						name: "EscapedItalicTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "EscapedItalicTextSimplePunctuation",
			pos:  position{line: 355, col: 1, offset: 14855},
			expr: &actionExpr{
				pos: position{line: 355, col: 39, offset: 14893},
				run: (*parser).callonEscapedItalicTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 355, col: 39, offset: 14893},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 355, col: 39, offset: 14893},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 355, col: 52, offset: 14906},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 355, col: 52, offset: 14906},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 355, col: 56, offset: 14910},
										expr: &litMatcher{
											pos:        position{line: 355, col: 56, offset: 14910},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 355, col: 62, offset: 14916},
							val:        "_",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 355, col: 66, offset: 14920},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 355, col: 75, offset: 14929},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 355, col: 94, offset: 14948},
							val:        "_",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedItalicTextDoublePunctuation",
			pos:  position{line: 359, col: 1, offset: 15054},
			expr: &actionExpr{
				pos: position{line: 359, col: 39, offset: 15092},
				run: (*parser).callonEscapedItalicTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 359, col: 39, offset: 15092},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 359, col: 39, offset: 15092},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 359, col: 52, offset: 15105},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 359, col: 52, offset: 15105},
										val:        "\\\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 359, col: 57, offset: 15110},
										expr: &litMatcher{
											pos:        position{line: 359, col: 57, offset: 15110},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 359, col: 63, offset: 15116},
							val:        "__",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 359, col: 68, offset: 15121},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 359, col: 77, offset: 15130},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 359, col: 96, offset: 15149},
							val:        "__",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedItalicTextUnbalancedPunctuation",
			pos:  position{line: 363, col: 1, offset: 15257},
			expr: &actionExpr{
				pos: position{line: 363, col: 44, offset: 15300},
				run: (*parser).callonEscapedItalicTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 363, col: 44, offset: 15300},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 363, col: 44, offset: 15300},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 363, col: 57, offset: 15313},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 363, col: 57, offset: 15313},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 363, col: 61, offset: 15317},
										expr: &litMatcher{
											pos:        position{line: 363, col: 61, offset: 15317},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 363, col: 67, offset: 15323},
							val:        "__",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 363, col: 72, offset: 15328},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 363, col: 81, offset: 15337},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 363, col: 100, offset: 15356},
							val:        "_",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "MonospaceText",
			pos:  position{line: 368, col: 1, offset: 15549},
			expr: &choiceExpr{
				pos: position{line: 368, col: 18, offset: 15566},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 368, col: 18, offset: 15566},
						name: "MonospaceTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 368, col: 51, offset: 15599},
						name: "MonospaceTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 368, col: 88, offset: 15636},
						name: "MonospaceTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "MonospaceTextSimplePunctuation",
			pos:  position{line: 370, col: 1, offset: 15668},
			expr: &actionExpr{
				pos: position{line: 370, col: 35, offset: 15702},
				run: (*parser).callonMonospaceTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 370, col: 35, offset: 15702},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 370, col: 35, offset: 15702},
							expr: &litMatcher{
								pos:        position{line: 370, col: 36, offset: 15703},
								val:        "\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 370, col: 40, offset: 15707},
							val:        "`",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 370, col: 44, offset: 15711},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 370, col: 53, offset: 15720},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 370, col: 72, offset: 15739},
							val:        "`",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "MonospaceTextDoublePunctuation",
			pos:  position{line: 374, col: 1, offset: 15821},
			expr: &actionExpr{
				pos: position{line: 374, col: 35, offset: 15855},
				run: (*parser).callonMonospaceTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 374, col: 35, offset: 15855},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 374, col: 35, offset: 15855},
							expr: &litMatcher{
								pos:        position{line: 374, col: 36, offset: 15856},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 374, col: 41, offset: 15861},
							val:        "``",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 374, col: 46, offset: 15866},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 374, col: 55, offset: 15875},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 374, col: 74, offset: 15894},
							val:        "``",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "MonospaceTextUnbalancedPunctuation",
			pos:  position{line: 378, col: 1, offset: 15977},
			expr: &actionExpr{
				pos: position{line: 378, col: 39, offset: 16015},
				run: (*parser).callonMonospaceTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 378, col: 39, offset: 16015},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 378, col: 39, offset: 16015},
							expr: &litMatcher{
								pos:        position{line: 378, col: 40, offset: 16016},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 378, col: 45, offset: 16021},
							val:        "``",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 378, col: 50, offset: 16026},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 378, col: 59, offset: 16035},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 378, col: 78, offset: 16054},
							val:        "`",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedMonospaceText",
			pos:  position{line: 383, col: 1, offset: 16223},
			expr: &choiceExpr{
				pos: position{line: 383, col: 25, offset: 16247},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 383, col: 25, offset: 16247},
						name: "EscapedMonospaceTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 383, col: 65, offset: 16287},
						name: "EscapedMonospaceTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 383, col: 109, offset: 16331},
						name: "EscapedMonospaceTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "EscapedMonospaceTextSimplePunctuation",
			pos:  position{line: 385, col: 1, offset: 16416},
			expr: &actionExpr{
				pos: position{line: 385, col: 42, offset: 16457},
				run: (*parser).callonEscapedMonospaceTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 385, col: 42, offset: 16457},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 385, col: 42, offset: 16457},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 385, col: 55, offset: 16470},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 385, col: 55, offset: 16470},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 385, col: 59, offset: 16474},
										expr: &litMatcher{
											pos:        position{line: 385, col: 59, offset: 16474},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 385, col: 65, offset: 16480},
							val:        "`",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 385, col: 69, offset: 16484},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 385, col: 78, offset: 16493},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 385, col: 97, offset: 16512},
							val:        "`",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedMonospaceTextDoublePunctuation",
			pos:  position{line: 389, col: 1, offset: 16618},
			expr: &actionExpr{
				pos: position{line: 389, col: 42, offset: 16659},
				run: (*parser).callonEscapedMonospaceTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 389, col: 42, offset: 16659},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 389, col: 42, offset: 16659},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 389, col: 55, offset: 16672},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 389, col: 55, offset: 16672},
										val:        "\\\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 389, col: 60, offset: 16677},
										expr: &litMatcher{
											pos:        position{line: 389, col: 60, offset: 16677},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 389, col: 66, offset: 16683},
							val:        "``",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 389, col: 71, offset: 16688},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 389, col: 80, offset: 16697},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 389, col: 99, offset: 16716},
							val:        "``",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedMonospaceTextUnbalancedPunctuation",
			pos:  position{line: 393, col: 1, offset: 16824},
			expr: &actionExpr{
				pos: position{line: 393, col: 47, offset: 16870},
				run: (*parser).callonEscapedMonospaceTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 393, col: 47, offset: 16870},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 393, col: 47, offset: 16870},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 393, col: 60, offset: 16883},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 393, col: 60, offset: 16883},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 393, col: 64, offset: 16887},
										expr: &litMatcher{
											pos:        position{line: 393, col: 64, offset: 16887},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 393, col: 70, offset: 16893},
							val:        "``",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 393, col: 75, offset: 16898},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 393, col: 84, offset: 16907},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 393, col: 103, offset: 16926},
							val:        "`",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "QuotedTextContent",
			pos:  position{line: 398, col: 1, offset: 17119},
			expr: &seqExpr{
				pos: position{line: 398, col: 22, offset: 17140},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 398, col: 22, offset: 17140},
						name: "QuotedTextContentElement",
					},
					&zeroOrMoreExpr{
						pos: position{line: 398, col: 47, offset: 17165},
						expr: &seqExpr{
							pos: position{line: 398, col: 48, offset: 17166},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 398, col: 48, offset: 17166},
									expr: &ruleRefExpr{
										pos:  position{line: 398, col: 48, offset: 17166},
										name: "WS",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 398, col: 52, offset: 17170},
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
			pos:  position{line: 400, col: 1, offset: 17198},
			expr: &choiceExpr{
				pos: position{line: 400, col: 29, offset: 17226},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 400, col: 29, offset: 17226},
						name: "QuotedText",
					},
					&ruleRefExpr{
						pos:  position{line: 400, col: 42, offset: 17239},
						name: "QuotedTextCharacters",
					},
					&ruleRefExpr{
						pos:  position{line: 400, col: 65, offset: 17262},
						name: "CharactersWithQuotePunctuation",
					},
				},
			},
		},
		{
			name: "QuotedTextCharacters",
			pos:  position{line: 402, col: 1, offset: 17397},
			expr: &oneOrMoreExpr{
				pos: position{line: 402, col: 25, offset: 17421},
				expr: &seqExpr{
					pos: position{line: 402, col: 26, offset: 17422},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 402, col: 26, offset: 17422},
							expr: &ruleRefExpr{
								pos:  position{line: 402, col: 27, offset: 17423},
								name: "NEWLINE",
							},
						},
						&notExpr{
							pos: position{line: 402, col: 35, offset: 17431},
							expr: &ruleRefExpr{
								pos:  position{line: 402, col: 36, offset: 17432},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 402, col: 39, offset: 17435},
							expr: &litMatcher{
								pos:        position{line: 402, col: 40, offset: 17436},
								val:        "*",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 402, col: 44, offset: 17440},
							expr: &litMatcher{
								pos:        position{line: 402, col: 45, offset: 17441},
								val:        "_",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 402, col: 49, offset: 17445},
							expr: &litMatcher{
								pos:        position{line: 402, col: 50, offset: 17446},
								val:        "`",
								ignoreCase: false,
							},
						},
						&anyMatcher{
							line: 402, col: 54, offset: 17450,
						},
					},
				},
			},
		},
		{
			name: "CharactersWithQuotePunctuation",
			pos:  position{line: 404, col: 1, offset: 17493},
			expr: &actionExpr{
				pos: position{line: 404, col: 35, offset: 17527},
				run: (*parser).callonCharactersWithQuotePunctuation1,
				expr: &oneOrMoreExpr{
					pos: position{line: 404, col: 35, offset: 17527},
					expr: &seqExpr{
						pos: position{line: 404, col: 36, offset: 17528},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 404, col: 36, offset: 17528},
								expr: &ruleRefExpr{
									pos:  position{line: 404, col: 37, offset: 17529},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 404, col: 45, offset: 17537},
								expr: &ruleRefExpr{
									pos:  position{line: 404, col: 46, offset: 17538},
									name: "WS",
								},
							},
							&anyMatcher{
								line: 404, col: 50, offset: 17542,
							},
						},
					},
				},
			},
		},
		{
			name: "UnbalancedQuotePunctuation",
			pos:  position{line: 409, col: 1, offset: 17787},
			expr: &choiceExpr{
				pos: position{line: 409, col: 31, offset: 17817},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 409, col: 31, offset: 17817},
						val:        "*",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 409, col: 37, offset: 17823},
						val:        "_",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 409, col: 43, offset: 17829},
						val:        "`",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "Passthrough",
			pos:  position{line: 414, col: 1, offset: 17941},
			expr: &choiceExpr{
				pos: position{line: 414, col: 16, offset: 17956},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 414, col: 16, offset: 17956},
						name: "TriplePlusPassthrough",
					},
					&ruleRefExpr{
						pos:  position{line: 414, col: 40, offset: 17980},
						name: "SinglePlusPassthrough",
					},
					&ruleRefExpr{
						pos:  position{line: 414, col: 64, offset: 18004},
						name: "PassthroughMacro",
					},
				},
			},
		},
		{
			name: "SinglePlusPassthrough",
			pos:  position{line: 416, col: 1, offset: 18022},
			expr: &actionExpr{
				pos: position{line: 416, col: 26, offset: 18047},
				run: (*parser).callonSinglePlusPassthrough1,
				expr: &seqExpr{
					pos: position{line: 416, col: 26, offset: 18047},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 416, col: 26, offset: 18047},
							val:        "+",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 416, col: 30, offset: 18051},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 416, col: 38, offset: 18059},
								expr: &seqExpr{
									pos: position{line: 416, col: 39, offset: 18060},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 416, col: 39, offset: 18060},
											expr: &ruleRefExpr{
												pos:  position{line: 416, col: 40, offset: 18061},
												name: "NEWLINE",
											},
										},
										&notExpr{
											pos: position{line: 416, col: 48, offset: 18069},
											expr: &litMatcher{
												pos:        position{line: 416, col: 49, offset: 18070},
												val:        "+",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 416, col: 53, offset: 18074,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 416, col: 57, offset: 18078},
							val:        "+",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TriplePlusPassthrough",
			pos:  position{line: 420, col: 1, offset: 18173},
			expr: &actionExpr{
				pos: position{line: 420, col: 26, offset: 18198},
				run: (*parser).callonTriplePlusPassthrough1,
				expr: &seqExpr{
					pos: position{line: 420, col: 26, offset: 18198},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 420, col: 26, offset: 18198},
							val:        "+++",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 420, col: 32, offset: 18204},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 420, col: 40, offset: 18212},
								expr: &seqExpr{
									pos: position{line: 420, col: 41, offset: 18213},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 420, col: 41, offset: 18213},
											expr: &litMatcher{
												pos:        position{line: 420, col: 42, offset: 18214},
												val:        "+++",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 420, col: 48, offset: 18220,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 420, col: 52, offset: 18224},
							val:        "+++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "PassthroughMacro",
			pos:  position{line: 424, col: 1, offset: 18321},
			expr: &choiceExpr{
				pos: position{line: 424, col: 21, offset: 18341},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 424, col: 21, offset: 18341},
						name: "SimplePassthroughMacro",
					},
					&ruleRefExpr{
						pos:  position{line: 424, col: 46, offset: 18366},
						name: "PassthroughWithQuotedText",
					},
				},
			},
		},
		{
			name: "SimplePassthroughMacro",
			pos:  position{line: 426, col: 1, offset: 18393},
			expr: &actionExpr{
				pos: position{line: 426, col: 27, offset: 18419},
				run: (*parser).callonSimplePassthroughMacro1,
				expr: &seqExpr{
					pos: position{line: 426, col: 27, offset: 18419},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 426, col: 27, offset: 18419},
							val:        "pass:[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 426, col: 36, offset: 18428},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 426, col: 44, offset: 18436},
								expr: &ruleRefExpr{
									pos:  position{line: 426, col: 45, offset: 18437},
									name: "PassthroughMacroCharacter",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 426, col: 73, offset: 18465},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "PassthroughWithQuotedText",
			pos:  position{line: 430, col: 1, offset: 18555},
			expr: &actionExpr{
				pos: position{line: 430, col: 30, offset: 18584},
				run: (*parser).callonPassthroughWithQuotedText1,
				expr: &seqExpr{
					pos: position{line: 430, col: 30, offset: 18584},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 430, col: 30, offset: 18584},
							val:        "pass:q[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 430, col: 40, offset: 18594},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 430, col: 48, offset: 18602},
								expr: &choiceExpr{
									pos: position{line: 430, col: 49, offset: 18603},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 430, col: 49, offset: 18603},
											name: "QuotedText",
										},
										&ruleRefExpr{
											pos:  position{line: 430, col: 62, offset: 18616},
											name: "PassthroughMacroCharacter",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 430, col: 90, offset: 18644},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "PassthroughMacroCharacter",
			pos:  position{line: 434, col: 1, offset: 18734},
			expr: &seqExpr{
				pos: position{line: 434, col: 31, offset: 18764},
				exprs: []interface{}{
					&notExpr{
						pos: position{line: 434, col: 31, offset: 18764},
						expr: &litMatcher{
							pos:        position{line: 434, col: 32, offset: 18765},
							val:        "]",
							ignoreCase: false,
						},
					},
					&anyMatcher{
						line: 434, col: 36, offset: 18769,
					},
				},
			},
		},
		{
			name: "CrossReference",
			pos:  position{line: 439, col: 1, offset: 18885},
			expr: &actionExpr{
				pos: position{line: 439, col: 19, offset: 18903},
				run: (*parser).callonCrossReference1,
				expr: &seqExpr{
					pos: position{line: 439, col: 19, offset: 18903},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 439, col: 19, offset: 18903},
							val:        "<<",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 439, col: 24, offset: 18908},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 439, col: 28, offset: 18912},
								name: "ID",
							},
						},
						&litMatcher{
							pos:        position{line: 439, col: 32, offset: 18916},
							val:        ">>",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ExternalLink",
			pos:  position{line: 446, col: 1, offset: 19075},
			expr: &actionExpr{
				pos: position{line: 446, col: 17, offset: 19091},
				run: (*parser).callonExternalLink1,
				expr: &seqExpr{
					pos: position{line: 446, col: 17, offset: 19091},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 446, col: 17, offset: 19091},
							label: "url",
							expr: &seqExpr{
								pos: position{line: 446, col: 22, offset: 19096},
								exprs: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 446, col: 22, offset: 19096},
										name: "URL_SCHEME",
									},
									&ruleRefExpr{
										pos:  position{line: 446, col: 33, offset: 19107},
										name: "URL",
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 446, col: 38, offset: 19112},
							label: "text",
							expr: &zeroOrOneExpr{
								pos: position{line: 446, col: 43, offset: 19117},
								expr: &seqExpr{
									pos: position{line: 446, col: 44, offset: 19118},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 446, col: 44, offset: 19118},
											val:        "[",
											ignoreCase: false,
										},
										&zeroOrMoreExpr{
											pos: position{line: 446, col: 48, offset: 19122},
											expr: &ruleRefExpr{
												pos:  position{line: 446, col: 49, offset: 19123},
												name: "URL_TEXT",
											},
										},
										&litMatcher{
											pos:        position{line: 446, col: 60, offset: 19134},
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
			pos:  position{line: 456, col: 1, offset: 19413},
			expr: &actionExpr{
				pos: position{line: 456, col: 15, offset: 19427},
				run: (*parser).callonBlockImage1,
				expr: &seqExpr{
					pos: position{line: 456, col: 15, offset: 19427},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 456, col: 15, offset: 19427},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 456, col: 26, offset: 19438},
								expr: &ruleRefExpr{
									pos:  position{line: 456, col: 27, offset: 19439},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 456, col: 46, offset: 19458},
							label: "image",
							expr: &ruleRefExpr{
								pos:  position{line: 456, col: 52, offset: 19464},
								name: "BlockImageMacro",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 456, col: 69, offset: 19481},
							expr: &ruleRefExpr{
								pos:  position{line: 456, col: 69, offset: 19481},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 456, col: 73, offset: 19485},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "BlockImageMacro",
			pos:  position{line: 461, col: 1, offset: 19646},
			expr: &actionExpr{
				pos: position{line: 461, col: 20, offset: 19665},
				run: (*parser).callonBlockImageMacro1,
				expr: &seqExpr{
					pos: position{line: 461, col: 20, offset: 19665},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 461, col: 20, offset: 19665},
							val:        "image::",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 461, col: 30, offset: 19675},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 461, col: 36, offset: 19681},
								name: "URL",
							},
						},
						&litMatcher{
							pos:        position{line: 461, col: 41, offset: 19686},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 461, col: 45, offset: 19690},
							label: "attributes",
							expr: &zeroOrOneExpr{
								pos: position{line: 461, col: 57, offset: 19702},
								expr: &ruleRefExpr{
									pos:  position{line: 461, col: 57, offset: 19702},
									name: "URL_TEXT",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 461, col: 68, offset: 19713},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "InlineImage",
			pos:  position{line: 465, col: 1, offset: 19780},
			expr: &actionExpr{
				pos: position{line: 465, col: 16, offset: 19795},
				run: (*parser).callonInlineImage1,
				expr: &labeledExpr{
					pos:   position{line: 465, col: 16, offset: 19795},
					label: "image",
					expr: &ruleRefExpr{
						pos:  position{line: 465, col: 22, offset: 19801},
						name: "InlineImageMacro",
					},
				},
			},
		},
		{
			name: "InlineImageMacro",
			pos:  position{line: 470, col: 1, offset: 19948},
			expr: &actionExpr{
				pos: position{line: 470, col: 21, offset: 19968},
				run: (*parser).callonInlineImageMacro1,
				expr: &seqExpr{
					pos: position{line: 470, col: 21, offset: 19968},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 470, col: 21, offset: 19968},
							val:        "image:",
							ignoreCase: false,
						},
						&notExpr{
							pos: position{line: 470, col: 30, offset: 19977},
							expr: &litMatcher{
								pos:        position{line: 470, col: 31, offset: 19978},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 470, col: 35, offset: 19982},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 470, col: 41, offset: 19988},
								name: "URL",
							},
						},
						&litMatcher{
							pos:        position{line: 470, col: 46, offset: 19993},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 470, col: 50, offset: 19997},
							label: "attributes",
							expr: &zeroOrOneExpr{
								pos: position{line: 470, col: 62, offset: 20009},
								expr: &ruleRefExpr{
									pos:  position{line: 470, col: 62, offset: 20009},
									name: "URL_TEXT",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 470, col: 73, offset: 20020},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "DelimitedBlock",
			pos:  position{line: 477, col: 1, offset: 20350},
			expr: &choiceExpr{
				pos: position{line: 477, col: 19, offset: 20368},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 477, col: 19, offset: 20368},
						name: "FencedBlock",
					},
					&ruleRefExpr{
						pos:  position{line: 477, col: 33, offset: 20382},
						name: "ListingBlock",
					},
				},
			},
		},
		{
			name: "BlockDelimiter",
			pos:  position{line: 479, col: 1, offset: 20397},
			expr: &choiceExpr{
				pos: position{line: 479, col: 19, offset: 20415},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 479, col: 19, offset: 20415},
						name: "FencedBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 479, col: 42, offset: 20438},
						name: "ListingBlockDelimiter",
					},
				},
			},
		},
		{
			name: "FencedBlockDelimiter",
			pos:  position{line: 481, col: 1, offset: 20461},
			expr: &litMatcher{
				pos:        position{line: 481, col: 25, offset: 20485},
				val:        "```",
				ignoreCase: false,
			},
		},
		{
			name: "FencedBlock",
			pos:  position{line: 483, col: 1, offset: 20492},
			expr: &actionExpr{
				pos: position{line: 483, col: 16, offset: 20507},
				run: (*parser).callonFencedBlock1,
				expr: &seqExpr{
					pos: position{line: 483, col: 16, offset: 20507},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 483, col: 16, offset: 20507},
							name: "FencedBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 483, col: 37, offset: 20528},
							expr: &ruleRefExpr{
								pos:  position{line: 483, col: 37, offset: 20528},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 483, col: 41, offset: 20532},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 483, col: 49, offset: 20540},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 483, col: 58, offset: 20549},
								name: "FencedBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 483, col: 78, offset: 20569},
							name: "FencedBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 483, col: 99, offset: 20590},
							expr: &ruleRefExpr{
								pos:  position{line: 483, col: 99, offset: 20590},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 483, col: 103, offset: 20594},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "FencedBlockContent",
			pos:  position{line: 487, col: 1, offset: 20682},
			expr: &labeledExpr{
				pos:   position{line: 487, col: 23, offset: 20704},
				label: "content",
				expr: &zeroOrMoreExpr{
					pos: position{line: 487, col: 31, offset: 20712},
					expr: &seqExpr{
						pos: position{line: 487, col: 32, offset: 20713},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 487, col: 32, offset: 20713},
								expr: &ruleRefExpr{
									pos:  position{line: 487, col: 33, offset: 20714},
									name: "FencedBlockDelimiter",
								},
							},
							&anyMatcher{
								line: 487, col: 54, offset: 20735,
							},
						},
					},
				},
			},
		},
		{
			name: "ListingBlockDelimiter",
			pos:  position{line: 489, col: 1, offset: 20741},
			expr: &litMatcher{
				pos:        position{line: 489, col: 26, offset: 20766},
				val:        "----",
				ignoreCase: false,
			},
		},
		{
			name: "ListingBlock",
			pos:  position{line: 491, col: 1, offset: 20774},
			expr: &actionExpr{
				pos: position{line: 491, col: 17, offset: 20790},
				run: (*parser).callonListingBlock1,
				expr: &seqExpr{
					pos: position{line: 491, col: 17, offset: 20790},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 491, col: 17, offset: 20790},
							name: "ListingBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 491, col: 39, offset: 20812},
							expr: &ruleRefExpr{
								pos:  position{line: 491, col: 39, offset: 20812},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 491, col: 43, offset: 20816},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 491, col: 51, offset: 20824},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 491, col: 60, offset: 20833},
								name: "ListingBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 491, col: 81, offset: 20854},
							name: "ListingBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 491, col: 103, offset: 20876},
							expr: &ruleRefExpr{
								pos:  position{line: 491, col: 103, offset: 20876},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 491, col: 107, offset: 20880},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ListingBlockContent",
			pos:  position{line: 495, col: 1, offset: 20969},
			expr: &labeledExpr{
				pos:   position{line: 495, col: 24, offset: 20992},
				label: "content",
				expr: &zeroOrMoreExpr{
					pos: position{line: 495, col: 32, offset: 21000},
					expr: &seqExpr{
						pos: position{line: 495, col: 33, offset: 21001},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 495, col: 33, offset: 21001},
								expr: &ruleRefExpr{
									pos:  position{line: 495, col: 34, offset: 21002},
									name: "ListingBlockDelimiter",
								},
							},
							&anyMatcher{
								line: 495, col: 56, offset: 21024,
							},
						},
					},
				},
			},
		},
		{
			name: "LiteralBlock",
			pos:  position{line: 500, col: 1, offset: 21297},
			expr: &choiceExpr{
				pos: position{line: 500, col: 17, offset: 21313},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 500, col: 17, offset: 21313},
						name: "ParagraphWithSpaces",
					},
					&ruleRefExpr{
						pos:  position{line: 500, col: 39, offset: 21335},
						name: "ParagraphWithLiteralBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 500, col: 76, offset: 21372},
						name: "ParagraphWithLiteralAttribute",
					},
				},
			},
		},
		{
			name: "ParagraphWithSpaces",
			pos:  position{line: 503, col: 1, offset: 21467},
			expr: &actionExpr{
				pos: position{line: 503, col: 24, offset: 21490},
				run: (*parser).callonParagraphWithSpaces1,
				expr: &seqExpr{
					pos: position{line: 503, col: 24, offset: 21490},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 503, col: 24, offset: 21490},
							label: "spaces",
							expr: &oneOrMoreExpr{
								pos: position{line: 503, col: 32, offset: 21498},
								expr: &ruleRefExpr{
									pos:  position{line: 503, col: 32, offset: 21498},
									name: "WS",
								},
							},
						},
						&notExpr{
							pos: position{line: 503, col: 37, offset: 21503},
							expr: &ruleRefExpr{
								pos:  position{line: 503, col: 38, offset: 21504},
								name: "NEWLINE",
							},
						},
						&labeledExpr{
							pos:   position{line: 503, col: 46, offset: 21512},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 503, col: 55, offset: 21521},
								name: "LiteralBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 503, col: 76, offset: 21542},
							name: "EndOfLiteralBlock",
						},
					},
				},
			},
		},
		{
			name: "LiteralBlockContent",
			pos:  position{line: 508, col: 1, offset: 21723},
			expr: &actionExpr{
				pos: position{line: 508, col: 24, offset: 21746},
				run: (*parser).callonLiteralBlockContent1,
				expr: &labeledExpr{
					pos:   position{line: 508, col: 24, offset: 21746},
					label: "content",
					expr: &oneOrMoreExpr{
						pos: position{line: 508, col: 32, offset: 21754},
						expr: &seqExpr{
							pos: position{line: 508, col: 33, offset: 21755},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 508, col: 33, offset: 21755},
									expr: &seqExpr{
										pos: position{line: 508, col: 35, offset: 21757},
										exprs: []interface{}{
											&ruleRefExpr{
												pos:  position{line: 508, col: 35, offset: 21757},
												name: "NEWLINE",
											},
											&ruleRefExpr{
												pos:  position{line: 508, col: 43, offset: 21765},
												name: "BlankLine",
											},
										},
									},
								},
								&anyMatcher{
									line: 508, col: 54, offset: 21776,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "EndOfLiteralBlock",
			pos:  position{line: 513, col: 1, offset: 21861},
			expr: &choiceExpr{
				pos: position{line: 513, col: 22, offset: 21882},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 513, col: 22, offset: 21882},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 513, col: 22, offset: 21882},
								name: "NEWLINE",
							},
							&ruleRefExpr{
								pos:  position{line: 513, col: 30, offset: 21890},
								name: "BlankLine",
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 513, col: 42, offset: 21902},
						name: "NEWLINE",
					},
					&ruleRefExpr{
						pos:  position{line: 513, col: 52, offset: 21912},
						name: "EOF",
					},
				},
			},
		},
		{
			name: "ParagraphWithLiteralBlockDelimiter",
			pos:  position{line: 516, col: 1, offset: 21972},
			expr: &actionExpr{
				pos: position{line: 516, col: 39, offset: 22010},
				run: (*parser).callonParagraphWithLiteralBlockDelimiter1,
				expr: &seqExpr{
					pos: position{line: 516, col: 39, offset: 22010},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 516, col: 39, offset: 22010},
							name: "LiteralBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 516, col: 61, offset: 22032},
							expr: &ruleRefExpr{
								pos:  position{line: 516, col: 61, offset: 22032},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 516, col: 65, offset: 22036},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 516, col: 73, offset: 22044},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 516, col: 81, offset: 22052},
								expr: &seqExpr{
									pos: position{line: 516, col: 82, offset: 22053},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 516, col: 82, offset: 22053},
											expr: &ruleRefExpr{
												pos:  position{line: 516, col: 83, offset: 22054},
												name: "LiteralBlockDelimiter",
											},
										},
										&anyMatcher{
											line: 516, col: 105, offset: 22076,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 516, col: 109, offset: 22080},
							name: "LiteralBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 516, col: 131, offset: 22102},
							expr: &ruleRefExpr{
								pos:  position{line: 516, col: 131, offset: 22102},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 516, col: 135, offset: 22106},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "LiteralBlockDelimiter",
			pos:  position{line: 520, col: 1, offset: 22190},
			expr: &litMatcher{
				pos:        position{line: 520, col: 26, offset: 22215},
				val:        "....",
				ignoreCase: false,
			},
		},
		{
			name: "ParagraphWithLiteralAttribute",
			pos:  position{line: 523, col: 1, offset: 22277},
			expr: &actionExpr{
				pos: position{line: 523, col: 34, offset: 22310},
				run: (*parser).callonParagraphWithLiteralAttribute1,
				expr: &seqExpr{
					pos: position{line: 523, col: 34, offset: 22310},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 523, col: 34, offset: 22310},
							val:        "[literal]",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 523, col: 46, offset: 22322},
							expr: &ruleRefExpr{
								pos:  position{line: 523, col: 46, offset: 22322},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 523, col: 50, offset: 22326},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 523, col: 58, offset: 22334},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 523, col: 67, offset: 22343},
								name: "LiteralBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 523, col: 88, offset: 22364},
							name: "EndOfLiteralBlock",
						},
					},
				},
			},
		},
		{
			name: "ElementAttribute",
			pos:  position{line: 530, col: 1, offset: 22576},
			expr: &choiceExpr{
				pos: position{line: 530, col: 21, offset: 22596},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 530, col: 21, offset: 22596},
						name: "ElementLink",
					},
					&ruleRefExpr{
						pos:  position{line: 530, col: 35, offset: 22610},
						name: "ElementID",
					},
					&ruleRefExpr{
						pos:  position{line: 530, col: 47, offset: 22622},
						name: "ElementTitle",
					},
					&ruleRefExpr{
						pos:  position{line: 530, col: 62, offset: 22637},
						name: "InvalidElementAttribute",
					},
				},
			},
		},
		{
			name: "ElementLink",
			pos:  position{line: 533, col: 1, offset: 22717},
			expr: &actionExpr{
				pos: position{line: 533, col: 16, offset: 22732},
				run: (*parser).callonElementLink1,
				expr: &seqExpr{
					pos: position{line: 533, col: 16, offset: 22732},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 533, col: 16, offset: 22732},
							val:        "[link=",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 533, col: 25, offset: 22741},
							expr: &ruleRefExpr{
								pos:  position{line: 533, col: 25, offset: 22741},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 533, col: 29, offset: 22745},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 533, col: 34, offset: 22750},
								name: "URL",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 533, col: 38, offset: 22754},
							expr: &ruleRefExpr{
								pos:  position{line: 533, col: 38, offset: 22754},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 533, col: 42, offset: 22758},
							val:        "]",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 533, col: 46, offset: 22762},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ElementID",
			pos:  position{line: 537, col: 1, offset: 22818},
			expr: &choiceExpr{
				pos: position{line: 537, col: 14, offset: 22831},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 537, col: 14, offset: 22831},
						name: "ElementIDNormal",
					},
					&ruleRefExpr{
						pos:  position{line: 537, col: 32, offset: 22849},
						name: "ElementIDShortHand",
					},
				},
			},
		},
		{
			name: "ElementIDNormal",
			pos:  position{line: 540, col: 1, offset: 22923},
			expr: &actionExpr{
				pos: position{line: 540, col: 20, offset: 22942},
				run: (*parser).callonElementIDNormal1,
				expr: &seqExpr{
					pos: position{line: 540, col: 20, offset: 22942},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 540, col: 20, offset: 22942},
							val:        "[[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 540, col: 25, offset: 22947},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 540, col: 29, offset: 22951},
								name: "ID",
							},
						},
						&litMatcher{
							pos:        position{line: 540, col: 33, offset: 22955},
							val:        "]]",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 540, col: 38, offset: 22960},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ElementIDShortHand",
			pos:  position{line: 543, col: 1, offset: 23011},
			expr: &actionExpr{
				pos: position{line: 543, col: 23, offset: 23033},
				run: (*parser).callonElementIDShortHand1,
				expr: &seqExpr{
					pos: position{line: 543, col: 23, offset: 23033},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 543, col: 23, offset: 23033},
							val:        "[#",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 543, col: 28, offset: 23038},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 543, col: 32, offset: 23042},
								name: "ID",
							},
						},
						&litMatcher{
							pos:        position{line: 543, col: 36, offset: 23046},
							val:        "]",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 543, col: 40, offset: 23050},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ElementTitle",
			pos:  position{line: 549, col: 1, offset: 23244},
			expr: &actionExpr{
				pos: position{line: 549, col: 17, offset: 23260},
				run: (*parser).callonElementTitle1,
				expr: &seqExpr{
					pos: position{line: 549, col: 17, offset: 23260},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 549, col: 17, offset: 23260},
							val:        ".",
							ignoreCase: false,
						},
						&notExpr{
							pos: position{line: 549, col: 21, offset: 23264},
							expr: &litMatcher{
								pos:        position{line: 549, col: 22, offset: 23265},
								val:        ".",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 549, col: 26, offset: 23269},
							expr: &ruleRefExpr{
								pos:  position{line: 549, col: 27, offset: 23270},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 549, col: 30, offset: 23273},
							label: "title",
							expr: &oneOrMoreExpr{
								pos: position{line: 549, col: 36, offset: 23279},
								expr: &seqExpr{
									pos: position{line: 549, col: 37, offset: 23280},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 549, col: 37, offset: 23280},
											expr: &ruleRefExpr{
												pos:  position{line: 549, col: 38, offset: 23281},
												name: "NEWLINE",
											},
										},
										&anyMatcher{
											line: 549, col: 46, offset: 23289,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 549, col: 50, offset: 23293},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "InvalidElementAttribute",
			pos:  position{line: 553, col: 1, offset: 23358},
			expr: &actionExpr{
				pos: position{line: 553, col: 28, offset: 23385},
				run: (*parser).callonInvalidElementAttribute1,
				expr: &seqExpr{
					pos: position{line: 553, col: 28, offset: 23385},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 553, col: 28, offset: 23385},
							val:        "[",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 553, col: 32, offset: 23389},
							expr: &ruleRefExpr{
								pos:  position{line: 553, col: 32, offset: 23389},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 553, col: 36, offset: 23393},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 553, col: 44, offset: 23401},
								expr: &seqExpr{
									pos: position{line: 553, col: 45, offset: 23402},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 553, col: 45, offset: 23402},
											expr: &litMatcher{
												pos:        position{line: 553, col: 46, offset: 23403},
												val:        "]",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 553, col: 50, offset: 23407,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 553, col: 54, offset: 23411},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "BlankLine",
			pos:  position{line: 560, col: 1, offset: 23577},
			expr: &actionExpr{
				pos: position{line: 560, col: 14, offset: 23590},
				run: (*parser).callonBlankLine1,
				expr: &seqExpr{
					pos: position{line: 560, col: 14, offset: 23590},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 560, col: 14, offset: 23590},
							expr: &ruleRefExpr{
								pos:  position{line: 560, col: 15, offset: 23591},
								name: "EOF",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 560, col: 19, offset: 23595},
							expr: &ruleRefExpr{
								pos:  position{line: 560, col: 19, offset: 23595},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 560, col: 23, offset: 23599},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "Characters",
			pos:  position{line: 567, col: 1, offset: 23746},
			expr: &actionExpr{
				pos: position{line: 567, col: 15, offset: 23760},
				run: (*parser).callonCharacters1,
				expr: &oneOrMoreExpr{
					pos: position{line: 567, col: 15, offset: 23760},
					expr: &seqExpr{
						pos: position{line: 567, col: 16, offset: 23761},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 567, col: 16, offset: 23761},
								expr: &ruleRefExpr{
									pos:  position{line: 567, col: 17, offset: 23762},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 567, col: 25, offset: 23770},
								expr: &ruleRefExpr{
									pos:  position{line: 567, col: 26, offset: 23771},
									name: "WS",
								},
							},
							&anyMatcher{
								line: 567, col: 29, offset: 23774,
							},
						},
					},
				},
			},
		},
		{
			name: "URL",
			pos:  position{line: 572, col: 1, offset: 23815},
			expr: &actionExpr{
				pos: position{line: 572, col: 8, offset: 23822},
				run: (*parser).callonURL1,
				expr: &oneOrMoreExpr{
					pos: position{line: 572, col: 8, offset: 23822},
					expr: &seqExpr{
						pos: position{line: 572, col: 9, offset: 23823},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 572, col: 9, offset: 23823},
								expr: &ruleRefExpr{
									pos:  position{line: 572, col: 10, offset: 23824},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 572, col: 18, offset: 23832},
								expr: &ruleRefExpr{
									pos:  position{line: 572, col: 19, offset: 23833},
									name: "WS",
								},
							},
							&notExpr{
								pos: position{line: 572, col: 22, offset: 23836},
								expr: &litMatcher{
									pos:        position{line: 572, col: 23, offset: 23837},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 572, col: 27, offset: 23841},
								expr: &litMatcher{
									pos:        position{line: 572, col: 28, offset: 23842},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 572, col: 32, offset: 23846,
							},
						},
					},
				},
			},
		},
		{
			name: "ID",
			pos:  position{line: 576, col: 1, offset: 23886},
			expr: &actionExpr{
				pos: position{line: 576, col: 7, offset: 23892},
				run: (*parser).callonID1,
				expr: &oneOrMoreExpr{
					pos: position{line: 576, col: 7, offset: 23892},
					expr: &seqExpr{
						pos: position{line: 576, col: 8, offset: 23893},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 576, col: 8, offset: 23893},
								expr: &ruleRefExpr{
									pos:  position{line: 576, col: 9, offset: 23894},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 576, col: 17, offset: 23902},
								expr: &ruleRefExpr{
									pos:  position{line: 576, col: 18, offset: 23903},
									name: "WS",
								},
							},
							&notExpr{
								pos: position{line: 576, col: 21, offset: 23906},
								expr: &litMatcher{
									pos:        position{line: 576, col: 22, offset: 23907},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 576, col: 26, offset: 23911},
								expr: &litMatcher{
									pos:        position{line: 576, col: 27, offset: 23912},
									val:        "]",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 576, col: 31, offset: 23916},
								expr: &litMatcher{
									pos:        position{line: 576, col: 32, offset: 23917},
									val:        "<<",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 576, col: 37, offset: 23922},
								expr: &litMatcher{
									pos:        position{line: 576, col: 38, offset: 23923},
									val:        ">>",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 576, col: 42, offset: 23927,
							},
						},
					},
				},
			},
		},
		{
			name: "URL_TEXT",
			pos:  position{line: 580, col: 1, offset: 23967},
			expr: &actionExpr{
				pos: position{line: 580, col: 13, offset: 23979},
				run: (*parser).callonURL_TEXT1,
				expr: &oneOrMoreExpr{
					pos: position{line: 580, col: 13, offset: 23979},
					expr: &seqExpr{
						pos: position{line: 580, col: 14, offset: 23980},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 580, col: 14, offset: 23980},
								expr: &ruleRefExpr{
									pos:  position{line: 580, col: 15, offset: 23981},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 580, col: 23, offset: 23989},
								expr: &litMatcher{
									pos:        position{line: 580, col: 24, offset: 23990},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 580, col: 28, offset: 23994},
								expr: &litMatcher{
									pos:        position{line: 580, col: 29, offset: 23995},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 580, col: 33, offset: 23999,
							},
						},
					},
				},
			},
		},
		{
			name: "URL_SCHEME",
			pos:  position{line: 584, col: 1, offset: 24039},
			expr: &choiceExpr{
				pos: position{line: 584, col: 15, offset: 24053},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 584, col: 15, offset: 24053},
						val:        "http://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 584, col: 27, offset: 24065},
						val:        "https://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 584, col: 40, offset: 24078},
						val:        "ftp://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 584, col: 51, offset: 24089},
						val:        "irc://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 584, col: 62, offset: 24100},
						val:        "mailto:",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "DIGIT",
			pos:  position{line: 586, col: 1, offset: 24111},
			expr: &charClassMatcher{
				pos:        position{line: 586, col: 10, offset: 24120},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NEWLINE",
			pos:  position{line: 588, col: 1, offset: 24127},
			expr: &choiceExpr{
				pos: position{line: 588, col: 12, offset: 24138},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 588, col: 12, offset: 24138},
						val:        "\r\n",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 588, col: 21, offset: 24147},
						val:        "\r",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 588, col: 28, offset: 24154},
						val:        "\n",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "WS",
			pos:  position{line: 590, col: 1, offset: 24160},
			expr: &choiceExpr{
				pos: position{line: 590, col: 7, offset: 24166},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 590, col: 7, offset: 24166},
						val:        " ",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 590, col: 13, offset: 24172},
						run: (*parser).callonWS3,
						expr: &litMatcher{
							pos:        position{line: 590, col: 13, offset: 24172},
							val:        "\t",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 594, col: 1, offset: 24217},
			expr: &notExpr{
				pos: position{line: 594, col: 8, offset: 24224},
				expr: &anyMatcher{
					line: 594, col: 9, offset: 24225,
				},
			},
		},
		{
			name: "EOL",
			pos:  position{line: 596, col: 1, offset: 24228},
			expr: &choiceExpr{
				pos: position{line: 596, col: 8, offset: 24235},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 596, col: 8, offset: 24235},
						name: "NEWLINE",
					},
					&ruleRefExpr{
						pos:  position{line: 596, col: 18, offset: 24245},
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

func (c *current) onListItemContinuation1() (interface{}, error) {
	return types.NewListItemContinuation()
}

func (p *parser) callonListItemContinuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onListItemContinuation1()
}

func (c *current) onContinuedBlockElement1(element interface{}) (interface{}, error) {
	return element, nil
}

func (p *parser) callonContinuedBlockElement1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onContinuedBlockElement1(stack["element"])
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
