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
								pos:  position{line: 56, col: 74, offset: 2039},
								name: "InlineElements",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 56, col: 90, offset: 2055},
							expr: &ruleRefExpr{
								pos:  position{line: 56, col: 90, offset: 2055},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 56, col: 94, offset: 2059},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 56, col: 97, offset: 2062},
								expr: &ruleRefExpr{
									pos:  position{line: 56, col: 98, offset: 2063},
									name: "InlineElementID",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 56, col: 116, offset: 2081},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthors",
			pos:  position{line: 60, col: 1, offset: 2197},
			expr: &choiceExpr{
				pos: position{line: 60, col: 20, offset: 2216},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 60, col: 20, offset: 2216},
						name: "DocumentAuthorsInlineForm",
					},
					&ruleRefExpr{
						pos:  position{line: 60, col: 48, offset: 2244},
						name: "DocumentAuthorsAttributeForm",
					},
				},
			},
		},
		{
			name: "DocumentAuthorsInlineForm",
			pos:  position{line: 62, col: 1, offset: 2274},
			expr: &actionExpr{
				pos: position{line: 62, col: 30, offset: 2303},
				run: (*parser).callonDocumentAuthorsInlineForm1,
				expr: &seqExpr{
					pos: position{line: 62, col: 30, offset: 2303},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 62, col: 30, offset: 2303},
							expr: &ruleRefExpr{
								pos:  position{line: 62, col: 30, offset: 2303},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 62, col: 34, offset: 2307},
							expr: &litMatcher{
								pos:        position{line: 62, col: 35, offset: 2308},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 62, col: 39, offset: 2312},
							label: "authors",
							expr: &oneOrMoreExpr{
								pos: position{line: 62, col: 48, offset: 2321},
								expr: &ruleRefExpr{
									pos:  position{line: 62, col: 48, offset: 2321},
									name: "DocumentAuthor",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 62, col: 65, offset: 2338},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthorsAttributeForm",
			pos:  position{line: 66, col: 1, offset: 2408},
			expr: &actionExpr{
				pos: position{line: 66, col: 33, offset: 2440},
				run: (*parser).callonDocumentAuthorsAttributeForm1,
				expr: &seqExpr{
					pos: position{line: 66, col: 33, offset: 2440},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 66, col: 33, offset: 2440},
							expr: &ruleRefExpr{
								pos:  position{line: 66, col: 33, offset: 2440},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 66, col: 37, offset: 2444},
							val:        ":author:",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 66, col: 48, offset: 2455},
							label: "author",
							expr: &ruleRefExpr{
								pos:  position{line: 66, col: 56, offset: 2463},
								name: "DocumentAuthor",
							},
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthor",
			pos:  position{line: 70, col: 1, offset: 2554},
			expr: &actionExpr{
				pos: position{line: 70, col: 19, offset: 2572},
				run: (*parser).callonDocumentAuthor1,
				expr: &seqExpr{
					pos: position{line: 70, col: 19, offset: 2572},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 70, col: 19, offset: 2572},
							expr: &ruleRefExpr{
								pos:  position{line: 70, col: 19, offset: 2572},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 70, col: 23, offset: 2576},
							label: "namePart1",
							expr: &ruleRefExpr{
								pos:  position{line: 70, col: 34, offset: 2587},
								name: "DocumentAuthorNamePart",
							},
						},
						&labeledExpr{
							pos:   position{line: 70, col: 58, offset: 2611},
							label: "namePart2",
							expr: &zeroOrOneExpr{
								pos: position{line: 70, col: 68, offset: 2621},
								expr: &ruleRefExpr{
									pos:  position{line: 70, col: 69, offset: 2622},
									name: "DocumentAuthorNamePart",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 70, col: 94, offset: 2647},
							label: "namePart3",
							expr: &zeroOrOneExpr{
								pos: position{line: 70, col: 104, offset: 2657},
								expr: &ruleRefExpr{
									pos:  position{line: 70, col: 105, offset: 2658},
									name: "DocumentAuthorNamePart",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 70, col: 130, offset: 2683},
							label: "email",
							expr: &zeroOrOneExpr{
								pos: position{line: 70, col: 136, offset: 2689},
								expr: &ruleRefExpr{
									pos:  position{line: 70, col: 137, offset: 2690},
									name: "DocumentAuthorEmail",
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 70, col: 159, offset: 2712},
							expr: &ruleRefExpr{
								pos:  position{line: 70, col: 159, offset: 2712},
								name: "WS",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 70, col: 163, offset: 2716},
							expr: &litMatcher{
								pos:        position{line: 70, col: 163, offset: 2716},
								val:        ";",
								ignoreCase: false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 70, col: 168, offset: 2721},
							expr: &ruleRefExpr{
								pos:  position{line: 70, col: 168, offset: 2721},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthorNamePart",
			pos:  position{line: 75, col: 1, offset: 2886},
			expr: &seqExpr{
				pos: position{line: 75, col: 27, offset: 2912},
				exprs: []interface{}{
					&notExpr{
						pos: position{line: 75, col: 27, offset: 2912},
						expr: &litMatcher{
							pos:        position{line: 75, col: 28, offset: 2913},
							val:        "<",
							ignoreCase: false,
						},
					},
					&notExpr{
						pos: position{line: 75, col: 32, offset: 2917},
						expr: &litMatcher{
							pos:        position{line: 75, col: 33, offset: 2918},
							val:        ";",
							ignoreCase: false,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 37, offset: 2922},
						name: "Characters",
					},
					&zeroOrMoreExpr{
						pos: position{line: 75, col: 48, offset: 2933},
						expr: &ruleRefExpr{
							pos:  position{line: 75, col: 48, offset: 2933},
							name: "WS",
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthorEmail",
			pos:  position{line: 77, col: 1, offset: 2938},
			expr: &seqExpr{
				pos: position{line: 77, col: 24, offset: 2961},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 77, col: 24, offset: 2961},
						val:        "<",
						ignoreCase: false,
					},
					&labeledExpr{
						pos:   position{line: 77, col: 28, offset: 2965},
						label: "email",
						expr: &oneOrMoreExpr{
							pos: position{line: 77, col: 34, offset: 2971},
							expr: &seqExpr{
								pos: position{line: 77, col: 35, offset: 2972},
								exprs: []interface{}{
									&notExpr{
										pos: position{line: 77, col: 35, offset: 2972},
										expr: &litMatcher{
											pos:        position{line: 77, col: 36, offset: 2973},
											val:        ">",
											ignoreCase: false,
										},
									},
									&notExpr{
										pos: position{line: 77, col: 40, offset: 2977},
										expr: &ruleRefExpr{
											pos:  position{line: 77, col: 41, offset: 2978},
											name: "EOL",
										},
									},
									&anyMatcher{
										line: 77, col: 45, offset: 2982,
									},
								},
							},
						},
					},
					&litMatcher{
						pos:        position{line: 77, col: 49, offset: 2986},
						val:        ">",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "DocumentRevision",
			pos:  position{line: 81, col: 1, offset: 3122},
			expr: &actionExpr{
				pos: position{line: 81, col: 21, offset: 3142},
				run: (*parser).callonDocumentRevision1,
				expr: &seqExpr{
					pos: position{line: 81, col: 21, offset: 3142},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 81, col: 21, offset: 3142},
							expr: &ruleRefExpr{
								pos:  position{line: 81, col: 21, offset: 3142},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 81, col: 25, offset: 3146},
							expr: &litMatcher{
								pos:        position{line: 81, col: 26, offset: 3147},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 81, col: 30, offset: 3151},
							label: "revnumber",
							expr: &zeroOrOneExpr{
								pos: position{line: 81, col: 40, offset: 3161},
								expr: &ruleRefExpr{
									pos:  position{line: 81, col: 41, offset: 3162},
									name: "DocumentRevisionNumber",
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 81, col: 66, offset: 3187},
							expr: &litMatcher{
								pos:        position{line: 81, col: 66, offset: 3187},
								val:        ",",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 81, col: 71, offset: 3192},
							label: "revdate",
							expr: &zeroOrOneExpr{
								pos: position{line: 81, col: 79, offset: 3200},
								expr: &ruleRefExpr{
									pos:  position{line: 81, col: 80, offset: 3201},
									name: "DocumentRevisionDate",
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 81, col: 103, offset: 3224},
							expr: &litMatcher{
								pos:        position{line: 81, col: 103, offset: 3224},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 81, col: 108, offset: 3229},
							label: "revremark",
							expr: &zeroOrOneExpr{
								pos: position{line: 81, col: 118, offset: 3239},
								expr: &ruleRefExpr{
									pos:  position{line: 81, col: 119, offset: 3240},
									name: "DocumentRevisionRemark",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 81, col: 144, offset: 3265},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentRevisionNumber",
			pos:  position{line: 86, col: 1, offset: 3438},
			expr: &choiceExpr{
				pos: position{line: 86, col: 27, offset: 3464},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 86, col: 27, offset: 3464},
						exprs: []interface{}{
							&litMatcher{
								pos:        position{line: 86, col: 27, offset: 3464},
								val:        "v",
								ignoreCase: true,
							},
							&ruleRefExpr{
								pos:  position{line: 86, col: 32, offset: 3469},
								name: "DIGIT",
							},
							&zeroOrMoreExpr{
								pos: position{line: 86, col: 39, offset: 3476},
								expr: &seqExpr{
									pos: position{line: 86, col: 40, offset: 3477},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 86, col: 40, offset: 3477},
											expr: &ruleRefExpr{
												pos:  position{line: 86, col: 41, offset: 3478},
												name: "EOL",
											},
										},
										&notExpr{
											pos: position{line: 86, col: 45, offset: 3482},
											expr: &litMatcher{
												pos:        position{line: 86, col: 46, offset: 3483},
												val:        ",",
												ignoreCase: false,
											},
										},
										&notExpr{
											pos: position{line: 86, col: 50, offset: 3487},
											expr: &litMatcher{
												pos:        position{line: 86, col: 51, offset: 3488},
												val:        ":",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 86, col: 55, offset: 3492,
										},
									},
								},
							},
						},
					},
					&seqExpr{
						pos: position{line: 86, col: 61, offset: 3498},
						exprs: []interface{}{
							&zeroOrOneExpr{
								pos: position{line: 86, col: 61, offset: 3498},
								expr: &litMatcher{
									pos:        position{line: 86, col: 61, offset: 3498},
									val:        "v",
									ignoreCase: true,
								},
							},
							&ruleRefExpr{
								pos:  position{line: 86, col: 67, offset: 3504},
								name: "DIGIT",
							},
							&zeroOrMoreExpr{
								pos: position{line: 86, col: 74, offset: 3511},
								expr: &seqExpr{
									pos: position{line: 86, col: 75, offset: 3512},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 86, col: 75, offset: 3512},
											expr: &ruleRefExpr{
												pos:  position{line: 86, col: 76, offset: 3513},
												name: "EOL",
											},
										},
										&notExpr{
											pos: position{line: 86, col: 80, offset: 3517},
											expr: &litMatcher{
												pos:        position{line: 86, col: 81, offset: 3518},
												val:        ",",
												ignoreCase: false,
											},
										},
										&notExpr{
											pos: position{line: 86, col: 85, offset: 3522},
											expr: &litMatcher{
												pos:        position{line: 86, col: 86, offset: 3523},
												val:        ":",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 86, col: 90, offset: 3527,
										},
									},
								},
							},
							&zeroOrMoreExpr{
								pos: position{line: 86, col: 94, offset: 3531},
								expr: &ruleRefExpr{
									pos:  position{line: 86, col: 94, offset: 3531},
									name: "WS",
								},
							},
							&andExpr{
								pos: position{line: 86, col: 98, offset: 3535},
								expr: &litMatcher{
									pos:        position{line: 86, col: 99, offset: 3536},
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
			pos:  position{line: 87, col: 1, offset: 3540},
			expr: &zeroOrMoreExpr{
				pos: position{line: 87, col: 25, offset: 3564},
				expr: &seqExpr{
					pos: position{line: 87, col: 26, offset: 3565},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 87, col: 26, offset: 3565},
							expr: &ruleRefExpr{
								pos:  position{line: 87, col: 27, offset: 3566},
								name: "EOL",
							},
						},
						&notExpr{
							pos: position{line: 87, col: 31, offset: 3570},
							expr: &litMatcher{
								pos:        position{line: 87, col: 32, offset: 3571},
								val:        ":",
								ignoreCase: false,
							},
						},
						&anyMatcher{
							line: 87, col: 36, offset: 3575,
						},
					},
				},
			},
		},
		{
			name: "DocumentRevisionRemark",
			pos:  position{line: 88, col: 1, offset: 3580},
			expr: &zeroOrMoreExpr{
				pos: position{line: 88, col: 27, offset: 3606},
				expr: &seqExpr{
					pos: position{line: 88, col: 28, offset: 3607},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 88, col: 28, offset: 3607},
							expr: &ruleRefExpr{
								pos:  position{line: 88, col: 29, offset: 3608},
								name: "EOL",
							},
						},
						&anyMatcher{
							line: 88, col: 33, offset: 3612,
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclaration",
			pos:  position{line: 93, col: 1, offset: 3732},
			expr: &choiceExpr{
				pos: position{line: 93, col: 33, offset: 3764},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 93, col: 33, offset: 3764},
						name: "DocumentAttributeDeclarationWithNameOnly",
					},
					&ruleRefExpr{
						pos:  position{line: 93, col: 76, offset: 3807},
						name: "DocumentAttributeDeclarationWithNameAndValue",
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclarationWithNameOnly",
			pos:  position{line: 95, col: 1, offset: 3854},
			expr: &actionExpr{
				pos: position{line: 95, col: 45, offset: 3898},
				run: (*parser).callonDocumentAttributeDeclarationWithNameOnly1,
				expr: &seqExpr{
					pos: position{line: 95, col: 45, offset: 3898},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 95, col: 45, offset: 3898},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 95, col: 49, offset: 3902},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 95, col: 55, offset: 3908},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 95, col: 70, offset: 3923},
							val:        ":",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 95, col: 74, offset: 3927},
							expr: &ruleRefExpr{
								pos:  position{line: 95, col: 74, offset: 3927},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 95, col: 78, offset: 3931},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclarationWithNameAndValue",
			pos:  position{line: 99, col: 1, offset: 4016},
			expr: &actionExpr{
				pos: position{line: 99, col: 49, offset: 4064},
				run: (*parser).callonDocumentAttributeDeclarationWithNameAndValue1,
				expr: &seqExpr{
					pos: position{line: 99, col: 49, offset: 4064},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 99, col: 49, offset: 4064},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 99, col: 53, offset: 4068},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 99, col: 59, offset: 4074},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 99, col: 74, offset: 4089},
							val:        ":",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 99, col: 78, offset: 4093},
							expr: &ruleRefExpr{
								pos:  position{line: 99, col: 78, offset: 4093},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 99, col: 82, offset: 4097},
							label: "value",
							expr: &zeroOrMoreExpr{
								pos: position{line: 99, col: 88, offset: 4103},
								expr: &seqExpr{
									pos: position{line: 99, col: 89, offset: 4104},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 99, col: 89, offset: 4104},
											expr: &ruleRefExpr{
												pos:  position{line: 99, col: 90, offset: 4105},
												name: "NEWLINE",
											},
										},
										&anyMatcher{
											line: 99, col: 98, offset: 4113,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 99, col: 102, offset: 4117},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeReset",
			pos:  position{line: 103, col: 1, offset: 4220},
			expr: &choiceExpr{
				pos: position{line: 103, col: 27, offset: 4246},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 103, col: 27, offset: 4246},
						name: "DocumentAttributeResetWithSectionTitleBangSymbol",
					},
					&ruleRefExpr{
						pos:  position{line: 103, col: 78, offset: 4297},
						name: "DocumentAttributeResetWithTrailingBangSymbol",
					},
				},
			},
		},
		{
			name: "DocumentAttributeResetWithSectionTitleBangSymbol",
			pos:  position{line: 105, col: 1, offset: 4343},
			expr: &actionExpr{
				pos: position{line: 105, col: 53, offset: 4395},
				run: (*parser).callonDocumentAttributeResetWithSectionTitleBangSymbol1,
				expr: &seqExpr{
					pos: position{line: 105, col: 53, offset: 4395},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 105, col: 53, offset: 4395},
							val:        ":!",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 105, col: 58, offset: 4400},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 105, col: 64, offset: 4406},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 105, col: 79, offset: 4421},
							val:        ":",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 105, col: 83, offset: 4425},
							expr: &ruleRefExpr{
								pos:  position{line: 105, col: 83, offset: 4425},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 105, col: 87, offset: 4429},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeResetWithTrailingBangSymbol",
			pos:  position{line: 109, col: 1, offset: 4503},
			expr: &actionExpr{
				pos: position{line: 109, col: 49, offset: 4551},
				run: (*parser).callonDocumentAttributeResetWithTrailingBangSymbol1,
				expr: &seqExpr{
					pos: position{line: 109, col: 49, offset: 4551},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 109, col: 49, offset: 4551},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 109, col: 53, offset: 4555},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 109, col: 59, offset: 4561},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 109, col: 74, offset: 4576},
							val:        "!:",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 109, col: 79, offset: 4581},
							expr: &ruleRefExpr{
								pos:  position{line: 109, col: 79, offset: 4581},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 109, col: 83, offset: 4585},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeSubstitution",
			pos:  position{line: 113, col: 1, offset: 4659},
			expr: &actionExpr{
				pos: position{line: 113, col: 34, offset: 4692},
				run: (*parser).callonDocumentAttributeSubstitution1,
				expr: &seqExpr{
					pos: position{line: 113, col: 34, offset: 4692},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 113, col: 34, offset: 4692},
							val:        "{",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 113, col: 38, offset: 4696},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 113, col: 44, offset: 4702},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 113, col: 59, offset: 4717},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "AttributeName",
			pos:  position{line: 120, col: 1, offset: 4971},
			expr: &seqExpr{
				pos: position{line: 120, col: 18, offset: 4988},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 120, col: 19, offset: 4989},
						alternatives: []interface{}{
							&charClassMatcher{
								pos:        position{line: 120, col: 19, offset: 4989},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
							&charClassMatcher{
								pos:        position{line: 120, col: 27, offset: 4997},
								val:        "[a-z]",
								ranges:     []rune{'a', 'z'},
								ignoreCase: false,
								inverted:   false,
							},
							&charClassMatcher{
								pos:        position{line: 120, col: 35, offset: 5005},
								val:        "[0-9]",
								ranges:     []rune{'0', '9'},
								ignoreCase: false,
								inverted:   false,
							},
							&litMatcher{
								pos:        position{line: 120, col: 43, offset: 5013},
								val:        "_",
								ignoreCase: false,
							},
						},
					},
					&zeroOrMoreExpr{
						pos: position{line: 120, col: 48, offset: 5018},
						expr: &choiceExpr{
							pos: position{line: 120, col: 49, offset: 5019},
							alternatives: []interface{}{
								&charClassMatcher{
									pos:        position{line: 120, col: 49, offset: 5019},
									val:        "[A-Z]",
									ranges:     []rune{'A', 'Z'},
									ignoreCase: false,
									inverted:   false,
								},
								&charClassMatcher{
									pos:        position{line: 120, col: 57, offset: 5027},
									val:        "[a-z]",
									ranges:     []rune{'a', 'z'},
									ignoreCase: false,
									inverted:   false,
								},
								&charClassMatcher{
									pos:        position{line: 120, col: 65, offset: 5035},
									val:        "[0-9]",
									ranges:     []rune{'0', '9'},
									ignoreCase: false,
									inverted:   false,
								},
								&litMatcher{
									pos:        position{line: 120, col: 73, offset: 5043},
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
			pos:  position{line: 125, col: 1, offset: 5163},
			expr: &seqExpr{
				pos: position{line: 125, col: 25, offset: 5187},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 125, col: 25, offset: 5187},
						val:        "toc::[]",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 125, col: 35, offset: 5197},
						name: "NEWLINE",
					},
				},
			},
		},
		{
			name: "ElementAttribute",
			pos:  position{line: 130, col: 1, offset: 5320},
			expr: &actionExpr{
				pos: position{line: 130, col: 21, offset: 5340},
				run: (*parser).callonElementAttribute1,
				expr: &seqExpr{
					pos: position{line: 130, col: 21, offset: 5340},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 130, col: 21, offset: 5340},
							label: "attr",
							expr: &choiceExpr{
								pos: position{line: 130, col: 27, offset: 5346},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 130, col: 27, offset: 5346},
										name: "ElementID",
									},
									&ruleRefExpr{
										pos:  position{line: 130, col: 39, offset: 5358},
										name: "ElementTitle",
									},
									&ruleRefExpr{
										pos:  position{line: 130, col: 54, offset: 5373},
										name: "AdmonitionMarkerAttribute",
									},
									&ruleRefExpr{
										pos:  position{line: 130, col: 82, offset: 5401},
										name: "AttributeGroup",
									},
									&ruleRefExpr{
										pos:  position{line: 130, col: 99, offset: 5418},
										name: "InvalidElementAttribute",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 130, col: 124, offset: 5443},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ElementID",
			pos:  position{line: 134, col: 1, offset: 5534},
			expr: &choiceExpr{
				pos: position{line: 134, col: 14, offset: 5547},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 134, col: 14, offset: 5547},
						run: (*parser).callonElementID2,
						expr: &labeledExpr{
							pos:   position{line: 134, col: 14, offset: 5547},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 134, col: 18, offset: 5551},
								name: "InlineElementID",
							},
						},
					},
					&actionExpr{
						pos: position{line: 136, col: 5, offset: 5593},
						run: (*parser).callonElementID5,
						expr: &seqExpr{
							pos: position{line: 136, col: 5, offset: 5593},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 136, col: 5, offset: 5593},
									val:        "[#",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 136, col: 10, offset: 5598},
									label: "id",
									expr: &ruleRefExpr{
										pos:  position{line: 136, col: 14, offset: 5602},
										name: "ID",
									},
								},
								&litMatcher{
									pos:        position{line: 136, col: 18, offset: 5606},
									val:        "]",
									ignoreCase: false,
								},
								&zeroOrMoreExpr{
									pos: position{line: 136, col: 22, offset: 5610},
									expr: &ruleRefExpr{
										pos:  position{line: 136, col: 22, offset: 5610},
										name: "WS",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "InlineElementID",
			pos:  position{line: 140, col: 1, offset: 5662},
			expr: &actionExpr{
				pos: position{line: 140, col: 20, offset: 5681},
				run: (*parser).callonInlineElementID1,
				expr: &seqExpr{
					pos: position{line: 140, col: 20, offset: 5681},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 140, col: 20, offset: 5681},
							val:        "[[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 140, col: 25, offset: 5686},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 140, col: 29, offset: 5690},
								name: "ID",
							},
						},
						&litMatcher{
							pos:        position{line: 140, col: 33, offset: 5694},
							val:        "]]",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 140, col: 38, offset: 5699},
							expr: &ruleRefExpr{
								pos:  position{line: 140, col: 38, offset: 5699},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "ElementTitle",
			pos:  position{line: 146, col: 1, offset: 5893},
			expr: &actionExpr{
				pos: position{line: 146, col: 17, offset: 5909},
				run: (*parser).callonElementTitle1,
				expr: &seqExpr{
					pos: position{line: 146, col: 17, offset: 5909},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 146, col: 17, offset: 5909},
							val:        ".",
							ignoreCase: false,
						},
						&notExpr{
							pos: position{line: 146, col: 21, offset: 5913},
							expr: &litMatcher{
								pos:        position{line: 146, col: 22, offset: 5914},
								val:        ".",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 146, col: 26, offset: 5918},
							expr: &ruleRefExpr{
								pos:  position{line: 146, col: 27, offset: 5919},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 146, col: 30, offset: 5922},
							label: "title",
							expr: &oneOrMoreExpr{
								pos: position{line: 146, col: 36, offset: 5928},
								expr: &seqExpr{
									pos: position{line: 146, col: 37, offset: 5929},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 146, col: 37, offset: 5929},
											expr: &ruleRefExpr{
												pos:  position{line: 146, col: 38, offset: 5930},
												name: "NEWLINE",
											},
										},
										&anyMatcher{
											line: 146, col: 46, offset: 5938,
										},
									},
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 146, col: 50, offset: 5942},
							expr: &ruleRefExpr{
								pos:  position{line: 146, col: 50, offset: 5942},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "AdmonitionMarkerAttribute",
			pos:  position{line: 151, col: 1, offset: 6087},
			expr: &actionExpr{
				pos: position{line: 151, col: 30, offset: 6116},
				run: (*parser).callonAdmonitionMarkerAttribute1,
				expr: &seqExpr{
					pos: position{line: 151, col: 30, offset: 6116},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 151, col: 30, offset: 6116},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 151, col: 34, offset: 6120},
							label: "k",
							expr: &ruleRefExpr{
								pos:  position{line: 151, col: 37, offset: 6123},
								name: "AdmonitionKind",
							},
						},
						&litMatcher{
							pos:        position{line: 151, col: 53, offset: 6139},
							val:        "]",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 151, col: 57, offset: 6143},
							expr: &ruleRefExpr{
								pos:  position{line: 151, col: 57, offset: 6143},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "AttributeGroup",
			pos:  position{line: 156, col: 1, offset: 6233},
			expr: &actionExpr{
				pos: position{line: 156, col: 19, offset: 6251},
				run: (*parser).callonAttributeGroup1,
				expr: &seqExpr{
					pos: position{line: 156, col: 19, offset: 6251},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 156, col: 19, offset: 6251},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 156, col: 23, offset: 6255},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 156, col: 34, offset: 6266},
								expr: &ruleRefExpr{
									pos:  position{line: 156, col: 35, offset: 6267},
									name: "GenericAttribute",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 156, col: 54, offset: 6286},
							val:        "]",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 156, col: 58, offset: 6290},
							expr: &ruleRefExpr{
								pos:  position{line: 156, col: 58, offset: 6290},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "GenericAttribute",
			pos:  position{line: 160, col: 1, offset: 6362},
			expr: &choiceExpr{
				pos: position{line: 160, col: 21, offset: 6382},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 160, col: 21, offset: 6382},
						run: (*parser).callonGenericAttribute2,
						expr: &seqExpr{
							pos: position{line: 160, col: 21, offset: 6382},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 160, col: 21, offset: 6382},
									label: "key",
									expr: &ruleRefExpr{
										pos:  position{line: 160, col: 26, offset: 6387},
										name: "AttributeKey",
									},
								},
								&litMatcher{
									pos:        position{line: 160, col: 40, offset: 6401},
									val:        "=",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 160, col: 44, offset: 6405},
									label: "value",
									expr: &ruleRefExpr{
										pos:  position{line: 160, col: 51, offset: 6412},
										name: "AttributeValue",
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 160, col: 67, offset: 6428},
									expr: &seqExpr{
										pos: position{line: 160, col: 68, offset: 6429},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 160, col: 68, offset: 6429},
												val:        ",",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 160, col: 72, offset: 6433},
												expr: &ruleRefExpr{
													pos:  position{line: 160, col: 72, offset: 6433},
													name: "WS",
												},
											},
										},
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 162, col: 5, offset: 6542},
						run: (*parser).callonGenericAttribute14,
						expr: &seqExpr{
							pos: position{line: 162, col: 5, offset: 6542},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 162, col: 5, offset: 6542},
									label: "key",
									expr: &ruleRefExpr{
										pos:  position{line: 162, col: 10, offset: 6547},
										name: "AttributeKey",
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 162, col: 24, offset: 6561},
									expr: &seqExpr{
										pos: position{line: 162, col: 25, offset: 6562},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 162, col: 25, offset: 6562},
												val:        ",",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 162, col: 29, offset: 6566},
												expr: &ruleRefExpr{
													pos:  position{line: 162, col: 29, offset: 6566},
													name: "WS",
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
		},
		{
			name: "AttributeKey",
			pos:  position{line: 166, col: 1, offset: 6660},
			expr: &actionExpr{
				pos: position{line: 166, col: 17, offset: 6676},
				run: (*parser).callonAttributeKey1,
				expr: &seqExpr{
					pos: position{line: 166, col: 17, offset: 6676},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 166, col: 17, offset: 6676},
							label: "key",
							expr: &oneOrMoreExpr{
								pos: position{line: 166, col: 22, offset: 6681},
								expr: &seqExpr{
									pos: position{line: 166, col: 23, offset: 6682},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 166, col: 23, offset: 6682},
											expr: &ruleRefExpr{
												pos:  position{line: 166, col: 24, offset: 6683},
												name: "WS",
											},
										},
										&notExpr{
											pos: position{line: 166, col: 27, offset: 6686},
											expr: &litMatcher{
												pos:        position{line: 166, col: 28, offset: 6687},
												val:        "=",
												ignoreCase: false,
											},
										},
										&notExpr{
											pos: position{line: 166, col: 32, offset: 6691},
											expr: &litMatcher{
												pos:        position{line: 166, col: 33, offset: 6692},
												val:        ",",
												ignoreCase: false,
											},
										},
										&notExpr{
											pos: position{line: 166, col: 37, offset: 6696},
											expr: &litMatcher{
												pos:        position{line: 166, col: 38, offset: 6697},
												val:        "]",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 166, col: 42, offset: 6701,
										},
									},
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 166, col: 46, offset: 6705},
							expr: &ruleRefExpr{
								pos:  position{line: 166, col: 46, offset: 6705},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "AttributeValue",
			pos:  position{line: 171, col: 1, offset: 6787},
			expr: &actionExpr{
				pos: position{line: 171, col: 19, offset: 6805},
				run: (*parser).callonAttributeValue1,
				expr: &seqExpr{
					pos: position{line: 171, col: 19, offset: 6805},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 171, col: 19, offset: 6805},
							expr: &ruleRefExpr{
								pos:  position{line: 171, col: 19, offset: 6805},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 171, col: 23, offset: 6809},
							label: "value",
							expr: &zeroOrMoreExpr{
								pos: position{line: 171, col: 29, offset: 6815},
								expr: &seqExpr{
									pos: position{line: 171, col: 30, offset: 6816},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 171, col: 30, offset: 6816},
											expr: &ruleRefExpr{
												pos:  position{line: 171, col: 31, offset: 6817},
												name: "WS",
											},
										},
										&notExpr{
											pos: position{line: 171, col: 34, offset: 6820},
											expr: &litMatcher{
												pos:        position{line: 171, col: 35, offset: 6821},
												val:        "=",
												ignoreCase: false,
											},
										},
										&notExpr{
											pos: position{line: 171, col: 39, offset: 6825},
											expr: &litMatcher{
												pos:        position{line: 171, col: 40, offset: 6826},
												val:        "]",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 171, col: 44, offset: 6830,
										},
									},
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 171, col: 48, offset: 6834},
							expr: &ruleRefExpr{
								pos:  position{line: 171, col: 48, offset: 6834},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "InvalidElementAttribute",
			pos:  position{line: 176, col: 1, offset: 6921},
			expr: &actionExpr{
				pos: position{line: 176, col: 28, offset: 6948},
				run: (*parser).callonInvalidElementAttribute1,
				expr: &seqExpr{
					pos: position{line: 176, col: 28, offset: 6948},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 176, col: 28, offset: 6948},
							val:        "[",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 176, col: 32, offset: 6952},
							expr: &ruleRefExpr{
								pos:  position{line: 176, col: 32, offset: 6952},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 176, col: 36, offset: 6956},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 176, col: 44, offset: 6964},
								expr: &seqExpr{
									pos: position{line: 176, col: 45, offset: 6965},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 176, col: 45, offset: 6965},
											expr: &litMatcher{
												pos:        position{line: 176, col: 46, offset: 6966},
												val:        "]",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 176, col: 50, offset: 6970,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 176, col: 54, offset: 6974},
							val:        "]",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 176, col: 58, offset: 6978},
							expr: &ruleRefExpr{
								pos:  position{line: 176, col: 58, offset: 6978},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "Section",
			pos:  position{line: 183, col: 1, offset: 7143},
			expr: &choiceExpr{
				pos: position{line: 183, col: 12, offset: 7154},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 183, col: 12, offset: 7154},
						name: "Section1",
					},
					&ruleRefExpr{
						pos:  position{line: 183, col: 23, offset: 7165},
						name: "Section2",
					},
					&ruleRefExpr{
						pos:  position{line: 183, col: 34, offset: 7176},
						name: "Section3",
					},
					&ruleRefExpr{
						pos:  position{line: 183, col: 45, offset: 7187},
						name: "Section4",
					},
					&ruleRefExpr{
						pos:  position{line: 183, col: 56, offset: 7198},
						name: "Section5",
					},
				},
			},
		},
		{
			name: "Section1",
			pos:  position{line: 186, col: 1, offset: 7209},
			expr: &actionExpr{
				pos: position{line: 186, col: 13, offset: 7221},
				run: (*parser).callonSection11,
				expr: &seqExpr{
					pos: position{line: 186, col: 13, offset: 7221},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 186, col: 13, offset: 7221},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 186, col: 21, offset: 7229},
								name: "Section1Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 186, col: 36, offset: 7244},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 186, col: 46, offset: 7254},
								expr: &ruleRefExpr{
									pos:  position{line: 186, col: 46, offset: 7254},
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
			pos:  position{line: 190, col: 1, offset: 7361},
			expr: &actionExpr{
				pos: position{line: 190, col: 18, offset: 7378},
				run: (*parser).callonSection1Block1,
				expr: &seqExpr{
					pos: position{line: 190, col: 18, offset: 7378},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 190, col: 18, offset: 7378},
							expr: &ruleRefExpr{
								pos:  position{line: 190, col: 19, offset: 7379},
								name: "Section1",
							},
						},
						&labeledExpr{
							pos:   position{line: 190, col: 28, offset: 7388},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 190, col: 37, offset: 7397},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 190, col: 37, offset: 7397},
										name: "Section2",
									},
									&ruleRefExpr{
										pos:  position{line: 190, col: 48, offset: 7408},
										name: "Section3",
									},
									&ruleRefExpr{
										pos:  position{line: 190, col: 59, offset: 7419},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 190, col: 70, offset: 7430},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 190, col: 81, offset: 7441},
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
			pos:  position{line: 194, col: 1, offset: 7484},
			expr: &actionExpr{
				pos: position{line: 194, col: 13, offset: 7496},
				run: (*parser).callonSection21,
				expr: &seqExpr{
					pos: position{line: 194, col: 13, offset: 7496},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 194, col: 13, offset: 7496},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 194, col: 21, offset: 7504},
								name: "Section2Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 194, col: 36, offset: 7519},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 194, col: 46, offset: 7529},
								expr: &ruleRefExpr{
									pos:  position{line: 194, col: 46, offset: 7529},
									name: "Section2Block",
								},
							},
						},
						&andExpr{
							pos: position{line: 194, col: 62, offset: 7545},
							expr: &zeroOrMoreExpr{
								pos: position{line: 194, col: 63, offset: 7546},
								expr: &ruleRefExpr{
									pos:  position{line: 194, col: 64, offset: 7547},
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
			pos:  position{line: 198, col: 1, offset: 7649},
			expr: &actionExpr{
				pos: position{line: 198, col: 18, offset: 7666},
				run: (*parser).callonSection2Block1,
				expr: &seqExpr{
					pos: position{line: 198, col: 18, offset: 7666},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 198, col: 18, offset: 7666},
							expr: &ruleRefExpr{
								pos:  position{line: 198, col: 19, offset: 7667},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 198, col: 28, offset: 7676},
							expr: &ruleRefExpr{
								pos:  position{line: 198, col: 29, offset: 7677},
								name: "Section2",
							},
						},
						&labeledExpr{
							pos:   position{line: 198, col: 38, offset: 7686},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 198, col: 47, offset: 7695},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 198, col: 47, offset: 7695},
										name: "Section3",
									},
									&ruleRefExpr{
										pos:  position{line: 198, col: 58, offset: 7706},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 198, col: 69, offset: 7717},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 198, col: 80, offset: 7728},
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
			pos:  position{line: 202, col: 1, offset: 7771},
			expr: &actionExpr{
				pos: position{line: 202, col: 13, offset: 7783},
				run: (*parser).callonSection31,
				expr: &seqExpr{
					pos: position{line: 202, col: 13, offset: 7783},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 202, col: 13, offset: 7783},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 202, col: 21, offset: 7791},
								name: "Section3Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 202, col: 36, offset: 7806},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 202, col: 46, offset: 7816},
								expr: &ruleRefExpr{
									pos:  position{line: 202, col: 46, offset: 7816},
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
			pos:  position{line: 206, col: 1, offset: 7923},
			expr: &actionExpr{
				pos: position{line: 206, col: 18, offset: 7940},
				run: (*parser).callonSection3Block1,
				expr: &seqExpr{
					pos: position{line: 206, col: 18, offset: 7940},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 206, col: 18, offset: 7940},
							expr: &ruleRefExpr{
								pos:  position{line: 206, col: 19, offset: 7941},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 206, col: 28, offset: 7950},
							expr: &ruleRefExpr{
								pos:  position{line: 206, col: 29, offset: 7951},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 206, col: 38, offset: 7960},
							expr: &ruleRefExpr{
								pos:  position{line: 206, col: 39, offset: 7961},
								name: "Section3",
							},
						},
						&labeledExpr{
							pos:   position{line: 206, col: 48, offset: 7970},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 206, col: 57, offset: 7979},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 206, col: 57, offset: 7979},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 206, col: 68, offset: 7990},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 206, col: 79, offset: 8001},
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
			pos:  position{line: 210, col: 1, offset: 8044},
			expr: &actionExpr{
				pos: position{line: 210, col: 13, offset: 8056},
				run: (*parser).callonSection41,
				expr: &seqExpr{
					pos: position{line: 210, col: 13, offset: 8056},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 210, col: 13, offset: 8056},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 210, col: 21, offset: 8064},
								name: "Section4Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 210, col: 36, offset: 8079},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 210, col: 46, offset: 8089},
								expr: &ruleRefExpr{
									pos:  position{line: 210, col: 46, offset: 8089},
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
			pos:  position{line: 214, col: 1, offset: 8196},
			expr: &actionExpr{
				pos: position{line: 214, col: 18, offset: 8213},
				run: (*parser).callonSection4Block1,
				expr: &seqExpr{
					pos: position{line: 214, col: 18, offset: 8213},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 214, col: 18, offset: 8213},
							expr: &ruleRefExpr{
								pos:  position{line: 214, col: 19, offset: 8214},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 214, col: 28, offset: 8223},
							expr: &ruleRefExpr{
								pos:  position{line: 214, col: 29, offset: 8224},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 214, col: 38, offset: 8233},
							expr: &ruleRefExpr{
								pos:  position{line: 214, col: 39, offset: 8234},
								name: "Section3",
							},
						},
						&notExpr{
							pos: position{line: 214, col: 48, offset: 8243},
							expr: &ruleRefExpr{
								pos:  position{line: 214, col: 49, offset: 8244},
								name: "Section4",
							},
						},
						&labeledExpr{
							pos:   position{line: 214, col: 58, offset: 8253},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 214, col: 67, offset: 8262},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 214, col: 67, offset: 8262},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 214, col: 78, offset: 8273},
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
			pos:  position{line: 218, col: 1, offset: 8316},
			expr: &actionExpr{
				pos: position{line: 218, col: 13, offset: 8328},
				run: (*parser).callonSection51,
				expr: &seqExpr{
					pos: position{line: 218, col: 13, offset: 8328},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 218, col: 13, offset: 8328},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 218, col: 21, offset: 8336},
								name: "Section5Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 218, col: 36, offset: 8351},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 218, col: 46, offset: 8361},
								expr: &ruleRefExpr{
									pos:  position{line: 218, col: 46, offset: 8361},
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
			pos:  position{line: 222, col: 1, offset: 8468},
			expr: &actionExpr{
				pos: position{line: 222, col: 18, offset: 8485},
				run: (*parser).callonSection5Block1,
				expr: &seqExpr{
					pos: position{line: 222, col: 18, offset: 8485},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 222, col: 18, offset: 8485},
							expr: &ruleRefExpr{
								pos:  position{line: 222, col: 19, offset: 8486},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 222, col: 28, offset: 8495},
							expr: &ruleRefExpr{
								pos:  position{line: 222, col: 29, offset: 8496},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 222, col: 38, offset: 8505},
							expr: &ruleRefExpr{
								pos:  position{line: 222, col: 39, offset: 8506},
								name: "Section3",
							},
						},
						&notExpr{
							pos: position{line: 222, col: 48, offset: 8515},
							expr: &ruleRefExpr{
								pos:  position{line: 222, col: 49, offset: 8516},
								name: "Section4",
							},
						},
						&notExpr{
							pos: position{line: 222, col: 58, offset: 8525},
							expr: &ruleRefExpr{
								pos:  position{line: 222, col: 59, offset: 8526},
								name: "Section5",
							},
						},
						&labeledExpr{
							pos:   position{line: 222, col: 68, offset: 8535},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 222, col: 77, offset: 8544},
								name: "BlockElement",
							},
						},
					},
				},
			},
		},
		{
			name: "SectionTitle",
			pos:  position{line: 230, col: 1, offset: 8698},
			expr: &choiceExpr{
				pos: position{line: 230, col: 17, offset: 8714},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 230, col: 17, offset: 8714},
						name: "Section1Title",
					},
					&ruleRefExpr{
						pos:  position{line: 230, col: 33, offset: 8730},
						name: "Section2Title",
					},
					&ruleRefExpr{
						pos:  position{line: 230, col: 49, offset: 8746},
						name: "Section3Title",
					},
					&ruleRefExpr{
						pos:  position{line: 230, col: 65, offset: 8762},
						name: "Section4Title",
					},
					&ruleRefExpr{
						pos:  position{line: 230, col: 81, offset: 8778},
						name: "Section5Title",
					},
				},
			},
		},
		{
			name: "Section1Title",
			pos:  position{line: 232, col: 1, offset: 8793},
			expr: &actionExpr{
				pos: position{line: 232, col: 18, offset: 8810},
				run: (*parser).callonSection1Title1,
				expr: &seqExpr{
					pos: position{line: 232, col: 18, offset: 8810},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 232, col: 18, offset: 8810},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 232, col: 29, offset: 8821},
								expr: &ruleRefExpr{
									pos:  position{line: 232, col: 30, offset: 8822},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 232, col: 49, offset: 8841},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 232, col: 56, offset: 8848},
								val:        "==",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 232, col: 62, offset: 8854},
							expr: &ruleRefExpr{
								pos:  position{line: 232, col: 62, offset: 8854},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 232, col: 66, offset: 8858},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 232, col: 75, offset: 8867},
								name: "InlineElements",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 232, col: 91, offset: 8883},
							expr: &ruleRefExpr{
								pos:  position{line: 232, col: 91, offset: 8883},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 232, col: 95, offset: 8887},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 232, col: 98, offset: 8890},
								expr: &ruleRefExpr{
									pos:  position{line: 232, col: 99, offset: 8891},
									name: "InlineElementID",
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 232, col: 117, offset: 8909},
							expr: &ruleRefExpr{
								pos:  position{line: 232, col: 117, offset: 8909},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 232, col: 121, offset: 8913},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 232, col: 126, offset: 8918},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 232, col: 126, offset: 8918},
									expr: &ruleRefExpr{
										pos:  position{line: 232, col: 126, offset: 8918},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 232, col: 139, offset: 8931},
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
			pos:  position{line: 236, col: 1, offset: 9047},
			expr: &actionExpr{
				pos: position{line: 236, col: 18, offset: 9064},
				run: (*parser).callonSection2Title1,
				expr: &seqExpr{
					pos: position{line: 236, col: 18, offset: 9064},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 236, col: 18, offset: 9064},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 236, col: 29, offset: 9075},
								expr: &ruleRefExpr{
									pos:  position{line: 236, col: 30, offset: 9076},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 236, col: 49, offset: 9095},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 236, col: 56, offset: 9102},
								val:        "===",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 236, col: 63, offset: 9109},
							expr: &ruleRefExpr{
								pos:  position{line: 236, col: 63, offset: 9109},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 236, col: 67, offset: 9113},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 236, col: 76, offset: 9122},
								name: "InlineElements",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 236, col: 92, offset: 9138},
							expr: &ruleRefExpr{
								pos:  position{line: 236, col: 92, offset: 9138},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 236, col: 96, offset: 9142},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 236, col: 99, offset: 9145},
								expr: &ruleRefExpr{
									pos:  position{line: 236, col: 100, offset: 9146},
									name: "InlineElementID",
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 236, col: 118, offset: 9164},
							expr: &ruleRefExpr{
								pos:  position{line: 236, col: 118, offset: 9164},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 236, col: 122, offset: 9168},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 236, col: 127, offset: 9173},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 236, col: 127, offset: 9173},
									expr: &ruleRefExpr{
										pos:  position{line: 236, col: 127, offset: 9173},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 236, col: 140, offset: 9186},
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
			pos:  position{line: 240, col: 1, offset: 9301},
			expr: &actionExpr{
				pos: position{line: 240, col: 18, offset: 9318},
				run: (*parser).callonSection3Title1,
				expr: &seqExpr{
					pos: position{line: 240, col: 18, offset: 9318},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 240, col: 18, offset: 9318},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 240, col: 29, offset: 9329},
								expr: &ruleRefExpr{
									pos:  position{line: 240, col: 30, offset: 9330},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 240, col: 49, offset: 9349},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 240, col: 56, offset: 9356},
								val:        "====",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 240, col: 64, offset: 9364},
							expr: &ruleRefExpr{
								pos:  position{line: 240, col: 64, offset: 9364},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 240, col: 68, offset: 9368},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 240, col: 77, offset: 9377},
								name: "InlineElements",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 240, col: 93, offset: 9393},
							expr: &ruleRefExpr{
								pos:  position{line: 240, col: 93, offset: 9393},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 240, col: 97, offset: 9397},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 240, col: 100, offset: 9400},
								expr: &ruleRefExpr{
									pos:  position{line: 240, col: 101, offset: 9401},
									name: "InlineElementID",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 240, col: 119, offset: 9419},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 240, col: 124, offset: 9424},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 240, col: 124, offset: 9424},
									expr: &ruleRefExpr{
										pos:  position{line: 240, col: 124, offset: 9424},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 240, col: 137, offset: 9437},
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
			pos:  position{line: 244, col: 1, offset: 9552},
			expr: &actionExpr{
				pos: position{line: 244, col: 18, offset: 9569},
				run: (*parser).callonSection4Title1,
				expr: &seqExpr{
					pos: position{line: 244, col: 18, offset: 9569},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 244, col: 18, offset: 9569},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 244, col: 29, offset: 9580},
								expr: &ruleRefExpr{
									pos:  position{line: 244, col: 30, offset: 9581},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 244, col: 49, offset: 9600},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 244, col: 56, offset: 9607},
								val:        "=====",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 244, col: 65, offset: 9616},
							expr: &ruleRefExpr{
								pos:  position{line: 244, col: 65, offset: 9616},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 244, col: 69, offset: 9620},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 244, col: 78, offset: 9629},
								name: "InlineElements",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 244, col: 94, offset: 9645},
							expr: &ruleRefExpr{
								pos:  position{line: 244, col: 94, offset: 9645},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 244, col: 98, offset: 9649},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 244, col: 101, offset: 9652},
								expr: &ruleRefExpr{
									pos:  position{line: 244, col: 102, offset: 9653},
									name: "InlineElementID",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 244, col: 120, offset: 9671},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 244, col: 125, offset: 9676},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 244, col: 125, offset: 9676},
									expr: &ruleRefExpr{
										pos:  position{line: 244, col: 125, offset: 9676},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 244, col: 138, offset: 9689},
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
			pos:  position{line: 248, col: 1, offset: 9804},
			expr: &actionExpr{
				pos: position{line: 248, col: 18, offset: 9821},
				run: (*parser).callonSection5Title1,
				expr: &seqExpr{
					pos: position{line: 248, col: 18, offset: 9821},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 248, col: 18, offset: 9821},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 248, col: 29, offset: 9832},
								expr: &ruleRefExpr{
									pos:  position{line: 248, col: 30, offset: 9833},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 248, col: 49, offset: 9852},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 248, col: 56, offset: 9859},
								val:        "======",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 248, col: 66, offset: 9869},
							expr: &ruleRefExpr{
								pos:  position{line: 248, col: 66, offset: 9869},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 248, col: 70, offset: 9873},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 248, col: 79, offset: 9882},
								name: "InlineElements",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 248, col: 95, offset: 9898},
							expr: &ruleRefExpr{
								pos:  position{line: 248, col: 95, offset: 9898},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 248, col: 99, offset: 9902},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 248, col: 102, offset: 9905},
								expr: &ruleRefExpr{
									pos:  position{line: 248, col: 103, offset: 9906},
									name: "InlineElementID",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 248, col: 121, offset: 9924},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 248, col: 126, offset: 9929},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 248, col: 126, offset: 9929},
									expr: &ruleRefExpr{
										pos:  position{line: 248, col: 126, offset: 9929},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 248, col: 139, offset: 9942},
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
			pos:  position{line: 255, col: 1, offset: 10158},
			expr: &actionExpr{
				pos: position{line: 255, col: 9, offset: 10166},
				run: (*parser).callonList1,
				expr: &seqExpr{
					pos: position{line: 255, col: 9, offset: 10166},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 255, col: 9, offset: 10166},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 255, col: 20, offset: 10177},
								expr: &ruleRefExpr{
									pos:  position{line: 255, col: 21, offset: 10178},
									name: "ListAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 257, col: 5, offset: 10267},
							label: "elements",
							expr: &ruleRefExpr{
								pos:  position{line: 257, col: 14, offset: 10276},
								name: "ListItems",
							},
						},
					},
				},
			},
		},
		{
			name: "ListItems",
			pos:  position{line: 261, col: 1, offset: 10370},
			expr: &oneOrMoreExpr{
				pos: position{line: 261, col: 14, offset: 10383},
				expr: &choiceExpr{
					pos: position{line: 261, col: 15, offset: 10384},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 261, col: 15, offset: 10384},
							name: "OrderedListItem",
						},
						&ruleRefExpr{
							pos:  position{line: 261, col: 33, offset: 10402},
							name: "UnorderedListItem",
						},
						&ruleRefExpr{
							pos:  position{line: 261, col: 53, offset: 10422},
							name: "LabeledListItem",
						},
					},
				},
			},
		},
		{
			name: "ListAttribute",
			pos:  position{line: 263, col: 1, offset: 10441},
			expr: &actionExpr{
				pos: position{line: 263, col: 18, offset: 10458},
				run: (*parser).callonListAttribute1,
				expr: &seqExpr{
					pos: position{line: 263, col: 18, offset: 10458},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 263, col: 18, offset: 10458},
							label: "attribute",
							expr: &choiceExpr{
								pos: position{line: 263, col: 29, offset: 10469},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 263, col: 29, offset: 10469},
										name: "HorizontalLayout",
									},
									&ruleRefExpr{
										pos:  position{line: 263, col: 48, offset: 10488},
										name: "ListID",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 263, col: 56, offset: 10496},
							name: "NEWLINE",
						},
					},
				},
			},
		},
		{
			name: "ListID",
			pos:  position{line: 267, col: 1, offset: 10535},
			expr: &actionExpr{
				pos: position{line: 267, col: 11, offset: 10545},
				run: (*parser).callonListID1,
				expr: &seqExpr{
					pos: position{line: 267, col: 11, offset: 10545},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 267, col: 11, offset: 10545},
							val:        "[#",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 267, col: 16, offset: 10550},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 267, col: 20, offset: 10554},
								name: "ID",
							},
						},
						&litMatcher{
							pos:        position{line: 267, col: 24, offset: 10558},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "HorizontalLayout",
			pos:  position{line: 271, col: 1, offset: 10624},
			expr: &actionExpr{
				pos: position{line: 271, col: 21, offset: 10644},
				run: (*parser).callonHorizontalLayout1,
				expr: &litMatcher{
					pos:        position{line: 271, col: 21, offset: 10644},
					val:        "[horizontal]",
					ignoreCase: false,
				},
			},
		},
		{
			name: "ListParagraph",
			pos:  position{line: 275, col: 1, offset: 10727},
			expr: &actionExpr{
				pos: position{line: 275, col: 19, offset: 10745},
				run: (*parser).callonListParagraph1,
				expr: &labeledExpr{
					pos:   position{line: 275, col: 19, offset: 10745},
					label: "lines",
					expr: &oneOrMoreExpr{
						pos: position{line: 275, col: 25, offset: 10751},
						expr: &seqExpr{
							pos: position{line: 276, col: 5, offset: 10757},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 276, col: 5, offset: 10757},
									expr: &ruleRefExpr{
										pos:  position{line: 276, col: 7, offset: 10759},
										name: "OrderedListItemPrefix",
									},
								},
								&notExpr{
									pos: position{line: 277, col: 5, offset: 10787},
									expr: &ruleRefExpr{
										pos:  position{line: 277, col: 7, offset: 10789},
										name: "UnorderedListItemPrefix",
									},
								},
								&notExpr{
									pos: position{line: 278, col: 5, offset: 10819},
									expr: &seqExpr{
										pos: position{line: 278, col: 7, offset: 10821},
										exprs: []interface{}{
											&ruleRefExpr{
												pos:  position{line: 278, col: 7, offset: 10821},
												name: "LabeledListItemTerm",
											},
											&ruleRefExpr{
												pos:  position{line: 278, col: 27, offset: 10841},
												name: "LabeledListItemSeparator",
											},
										},
									},
								},
								&notExpr{
									pos: position{line: 279, col: 5, offset: 10872},
									expr: &ruleRefExpr{
										pos:  position{line: 279, col: 7, offset: 10874},
										name: "ListItemContinuation",
									},
								},
								&notExpr{
									pos: position{line: 280, col: 5, offset: 10901},
									expr: &ruleRefExpr{
										pos:  position{line: 280, col: 7, offset: 10903},
										name: "ElementAttribute",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 281, col: 5, offset: 10925},
									name: "InlineElementsWithTrailingSpaces",
								},
								&ruleRefExpr{
									pos:  position{line: 281, col: 38, offset: 10958},
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
			pos:  position{line: 285, col: 1, offset: 11027},
			expr: &actionExpr{
				pos: position{line: 285, col: 25, offset: 11051},
				run: (*parser).callonListItemContinuation1,
				expr: &seqExpr{
					pos: position{line: 285, col: 25, offset: 11051},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 285, col: 25, offset: 11051},
							val:        "+",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 285, col: 29, offset: 11055},
							expr: &ruleRefExpr{
								pos:  position{line: 285, col: 29, offset: 11055},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 285, col: 33, offset: 11059},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ContinuedBlockElement",
			pos:  position{line: 289, col: 1, offset: 11111},
			expr: &actionExpr{
				pos: position{line: 289, col: 26, offset: 11136},
				run: (*parser).callonContinuedBlockElement1,
				expr: &seqExpr{
					pos: position{line: 289, col: 26, offset: 11136},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 289, col: 26, offset: 11136},
							name: "ListItemContinuation",
						},
						&labeledExpr{
							pos:   position{line: 289, col: 47, offset: 11157},
							label: "element",
							expr: &ruleRefExpr{
								pos:  position{line: 289, col: 55, offset: 11165},
								name: "BlockElement",
							},
						},
					},
				},
			},
		},
		{
			name: "OrderedListItem",
			pos:  position{line: 296, col: 1, offset: 11321},
			expr: &actionExpr{
				pos: position{line: 296, col: 20, offset: 11340},
				run: (*parser).callonOrderedListItem1,
				expr: &seqExpr{
					pos: position{line: 296, col: 20, offset: 11340},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 296, col: 20, offset: 11340},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 296, col: 31, offset: 11351},
								expr: &ruleRefExpr{
									pos:  position{line: 296, col: 32, offset: 11352},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 296, col: 51, offset: 11371},
							label: "prefix",
							expr: &ruleRefExpr{
								pos:  position{line: 296, col: 59, offset: 11379},
								name: "OrderedListItemPrefix",
							},
						},
						&labeledExpr{
							pos:   position{line: 296, col: 82, offset: 11402},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 296, col: 91, offset: 11411},
								name: "OrderedListItemContent",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 296, col: 115, offset: 11435},
							expr: &ruleRefExpr{
								pos:  position{line: 296, col: 115, offset: 11435},
								name: "BlankLine",
							},
						},
					},
				},
			},
		},
		{
			name: "OrderedListItemPrefix",
			pos:  position{line: 300, col: 1, offset: 11578},
			expr: &choiceExpr{
				pos: position{line: 302, col: 1, offset: 11642},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 302, col: 1, offset: 11642},
						run: (*parser).callonOrderedListItemPrefix2,
						expr: &seqExpr{
							pos: position{line: 302, col: 1, offset: 11642},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 302, col: 1, offset: 11642},
									expr: &ruleRefExpr{
										pos:  position{line: 302, col: 1, offset: 11642},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 302, col: 5, offset: 11646},
									label: "style",
									expr: &litMatcher{
										pos:        position{line: 302, col: 12, offset: 11653},
										val:        ".",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 302, col: 17, offset: 11658},
									expr: &ruleRefExpr{
										pos:  position{line: 302, col: 17, offset: 11658},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 304, col: 5, offset: 11751},
						run: (*parser).callonOrderedListItemPrefix10,
						expr: &seqExpr{
							pos: position{line: 304, col: 5, offset: 11751},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 304, col: 5, offset: 11751},
									expr: &ruleRefExpr{
										pos:  position{line: 304, col: 5, offset: 11751},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 304, col: 9, offset: 11755},
									label: "style",
									expr: &litMatcher{
										pos:        position{line: 304, col: 16, offset: 11762},
										val:        "..",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 304, col: 22, offset: 11768},
									expr: &ruleRefExpr{
										pos:  position{line: 304, col: 22, offset: 11768},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 306, col: 5, offset: 11866},
						run: (*parser).callonOrderedListItemPrefix18,
						expr: &seqExpr{
							pos: position{line: 306, col: 5, offset: 11866},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 306, col: 5, offset: 11866},
									expr: &ruleRefExpr{
										pos:  position{line: 306, col: 5, offset: 11866},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 306, col: 9, offset: 11870},
									label: "style",
									expr: &litMatcher{
										pos:        position{line: 306, col: 16, offset: 11877},
										val:        "...",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 306, col: 23, offset: 11884},
									expr: &ruleRefExpr{
										pos:  position{line: 306, col: 23, offset: 11884},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 308, col: 5, offset: 11983},
						run: (*parser).callonOrderedListItemPrefix26,
						expr: &seqExpr{
							pos: position{line: 308, col: 5, offset: 11983},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 308, col: 5, offset: 11983},
									expr: &ruleRefExpr{
										pos:  position{line: 308, col: 5, offset: 11983},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 308, col: 9, offset: 11987},
									label: "style",
									expr: &litMatcher{
										pos:        position{line: 308, col: 16, offset: 11994},
										val:        "....",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 308, col: 24, offset: 12002},
									expr: &ruleRefExpr{
										pos:  position{line: 308, col: 24, offset: 12002},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 310, col: 5, offset: 12102},
						run: (*parser).callonOrderedListItemPrefix34,
						expr: &seqExpr{
							pos: position{line: 310, col: 5, offset: 12102},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 310, col: 5, offset: 12102},
									expr: &ruleRefExpr{
										pos:  position{line: 310, col: 5, offset: 12102},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 310, col: 9, offset: 12106},
									label: "style",
									expr: &litMatcher{
										pos:        position{line: 310, col: 16, offset: 12113},
										val:        ".....",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 310, col: 25, offset: 12122},
									expr: &ruleRefExpr{
										pos:  position{line: 310, col: 25, offset: 12122},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 313, col: 5, offset: 12245},
						run: (*parser).callonOrderedListItemPrefix42,
						expr: &seqExpr{
							pos: position{line: 313, col: 5, offset: 12245},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 313, col: 5, offset: 12245},
									expr: &ruleRefExpr{
										pos:  position{line: 313, col: 5, offset: 12245},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 313, col: 9, offset: 12249},
									label: "style",
									expr: &seqExpr{
										pos: position{line: 313, col: 16, offset: 12256},
										exprs: []interface{}{
											&oneOrMoreExpr{
												pos: position{line: 313, col: 16, offset: 12256},
												expr: &seqExpr{
													pos: position{line: 313, col: 17, offset: 12257},
													exprs: []interface{}{
														&notExpr{
															pos: position{line: 313, col: 17, offset: 12257},
															expr: &litMatcher{
																pos:        position{line: 313, col: 18, offset: 12258},
																val:        ".",
																ignoreCase: false,
															},
														},
														&notExpr{
															pos: position{line: 313, col: 22, offset: 12262},
															expr: &ruleRefExpr{
																pos:  position{line: 313, col: 23, offset: 12263},
																name: "WS",
															},
														},
														&notExpr{
															pos: position{line: 313, col: 26, offset: 12266},
															expr: &ruleRefExpr{
																pos:  position{line: 313, col: 27, offset: 12267},
																name: "NEWLINE",
															},
														},
														&charClassMatcher{
															pos:        position{line: 313, col: 35, offset: 12275},
															val:        "[0-9]",
															ranges:     []rune{'0', '9'},
															ignoreCase: false,
															inverted:   false,
														},
													},
												},
											},
											&litMatcher{
												pos:        position{line: 313, col: 43, offset: 12283},
												val:        ".",
												ignoreCase: false,
											},
										},
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 313, col: 48, offset: 12288},
									expr: &ruleRefExpr{
										pos:  position{line: 313, col: 48, offset: 12288},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 315, col: 5, offset: 12383},
						run: (*parser).callonOrderedListItemPrefix60,
						expr: &seqExpr{
							pos: position{line: 315, col: 5, offset: 12383},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 315, col: 5, offset: 12383},
									expr: &ruleRefExpr{
										pos:  position{line: 315, col: 5, offset: 12383},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 315, col: 9, offset: 12387},
									label: "style",
									expr: &seqExpr{
										pos: position{line: 315, col: 16, offset: 12394},
										exprs: []interface{}{
											&oneOrMoreExpr{
												pos: position{line: 315, col: 16, offset: 12394},
												expr: &seqExpr{
													pos: position{line: 315, col: 17, offset: 12395},
													exprs: []interface{}{
														&notExpr{
															pos: position{line: 315, col: 17, offset: 12395},
															expr: &litMatcher{
																pos:        position{line: 315, col: 18, offset: 12396},
																val:        ".",
																ignoreCase: false,
															},
														},
														&notExpr{
															pos: position{line: 315, col: 22, offset: 12400},
															expr: &ruleRefExpr{
																pos:  position{line: 315, col: 23, offset: 12401},
																name: "WS",
															},
														},
														&notExpr{
															pos: position{line: 315, col: 26, offset: 12404},
															expr: &ruleRefExpr{
																pos:  position{line: 315, col: 27, offset: 12405},
																name: "NEWLINE",
															},
														},
														&charClassMatcher{
															pos:        position{line: 315, col: 35, offset: 12413},
															val:        "[a-z]",
															ranges:     []rune{'a', 'z'},
															ignoreCase: false,
															inverted:   false,
														},
													},
												},
											},
											&litMatcher{
												pos:        position{line: 315, col: 43, offset: 12421},
												val:        ".",
												ignoreCase: false,
											},
										},
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 315, col: 48, offset: 12426},
									expr: &ruleRefExpr{
										pos:  position{line: 315, col: 48, offset: 12426},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 317, col: 5, offset: 12524},
						run: (*parser).callonOrderedListItemPrefix78,
						expr: &seqExpr{
							pos: position{line: 317, col: 5, offset: 12524},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 317, col: 5, offset: 12524},
									expr: &ruleRefExpr{
										pos:  position{line: 317, col: 5, offset: 12524},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 317, col: 9, offset: 12528},
									label: "style",
									expr: &seqExpr{
										pos: position{line: 317, col: 16, offset: 12535},
										exprs: []interface{}{
											&oneOrMoreExpr{
												pos: position{line: 317, col: 16, offset: 12535},
												expr: &seqExpr{
													pos: position{line: 317, col: 17, offset: 12536},
													exprs: []interface{}{
														&notExpr{
															pos: position{line: 317, col: 17, offset: 12536},
															expr: &litMatcher{
																pos:        position{line: 317, col: 18, offset: 12537},
																val:        ".",
																ignoreCase: false,
															},
														},
														&notExpr{
															pos: position{line: 317, col: 22, offset: 12541},
															expr: &ruleRefExpr{
																pos:  position{line: 317, col: 23, offset: 12542},
																name: "WS",
															},
														},
														&notExpr{
															pos: position{line: 317, col: 26, offset: 12545},
															expr: &ruleRefExpr{
																pos:  position{line: 317, col: 27, offset: 12546},
																name: "NEWLINE",
															},
														},
														&charClassMatcher{
															pos:        position{line: 317, col: 35, offset: 12554},
															val:        "[A-Z]",
															ranges:     []rune{'A', 'Z'},
															ignoreCase: false,
															inverted:   false,
														},
													},
												},
											},
											&litMatcher{
												pos:        position{line: 317, col: 43, offset: 12562},
												val:        ".",
												ignoreCase: false,
											},
										},
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 317, col: 48, offset: 12567},
									expr: &ruleRefExpr{
										pos:  position{line: 317, col: 48, offset: 12567},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 319, col: 5, offset: 12665},
						run: (*parser).callonOrderedListItemPrefix96,
						expr: &seqExpr{
							pos: position{line: 319, col: 5, offset: 12665},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 319, col: 5, offset: 12665},
									expr: &ruleRefExpr{
										pos:  position{line: 319, col: 5, offset: 12665},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 319, col: 9, offset: 12669},
									label: "style",
									expr: &seqExpr{
										pos: position{line: 319, col: 16, offset: 12676},
										exprs: []interface{}{
											&oneOrMoreExpr{
												pos: position{line: 319, col: 16, offset: 12676},
												expr: &seqExpr{
													pos: position{line: 319, col: 17, offset: 12677},
													exprs: []interface{}{
														&notExpr{
															pos: position{line: 319, col: 17, offset: 12677},
															expr: &litMatcher{
																pos:        position{line: 319, col: 18, offset: 12678},
																val:        ")",
																ignoreCase: false,
															},
														},
														&notExpr{
															pos: position{line: 319, col: 22, offset: 12682},
															expr: &ruleRefExpr{
																pos:  position{line: 319, col: 23, offset: 12683},
																name: "WS",
															},
														},
														&notExpr{
															pos: position{line: 319, col: 26, offset: 12686},
															expr: &ruleRefExpr{
																pos:  position{line: 319, col: 27, offset: 12687},
																name: "NEWLINE",
															},
														},
														&charClassMatcher{
															pos:        position{line: 319, col: 35, offset: 12695},
															val:        "[a-z]",
															ranges:     []rune{'a', 'z'},
															ignoreCase: false,
															inverted:   false,
														},
													},
												},
											},
											&litMatcher{
												pos:        position{line: 319, col: 43, offset: 12703},
												val:        ")",
												ignoreCase: false,
											},
										},
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 319, col: 48, offset: 12708},
									expr: &ruleRefExpr{
										pos:  position{line: 319, col: 48, offset: 12708},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 321, col: 5, offset: 12806},
						run: (*parser).callonOrderedListItemPrefix114,
						expr: &seqExpr{
							pos: position{line: 321, col: 5, offset: 12806},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 321, col: 5, offset: 12806},
									expr: &ruleRefExpr{
										pos:  position{line: 321, col: 5, offset: 12806},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 321, col: 9, offset: 12810},
									label: "style",
									expr: &seqExpr{
										pos: position{line: 321, col: 16, offset: 12817},
										exprs: []interface{}{
											&oneOrMoreExpr{
												pos: position{line: 321, col: 16, offset: 12817},
												expr: &seqExpr{
													pos: position{line: 321, col: 17, offset: 12818},
													exprs: []interface{}{
														&notExpr{
															pos: position{line: 321, col: 17, offset: 12818},
															expr: &litMatcher{
																pos:        position{line: 321, col: 18, offset: 12819},
																val:        ")",
																ignoreCase: false,
															},
														},
														&notExpr{
															pos: position{line: 321, col: 22, offset: 12823},
															expr: &ruleRefExpr{
																pos:  position{line: 321, col: 23, offset: 12824},
																name: "WS",
															},
														},
														&notExpr{
															pos: position{line: 321, col: 26, offset: 12827},
															expr: &ruleRefExpr{
																pos:  position{line: 321, col: 27, offset: 12828},
																name: "NEWLINE",
															},
														},
														&charClassMatcher{
															pos:        position{line: 321, col: 35, offset: 12836},
															val:        "[A-Z]",
															ranges:     []rune{'A', 'Z'},
															ignoreCase: false,
															inverted:   false,
														},
													},
												},
											},
											&litMatcher{
												pos:        position{line: 321, col: 43, offset: 12844},
												val:        ")",
												ignoreCase: false,
											},
										},
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 321, col: 48, offset: 12849},
									expr: &ruleRefExpr{
										pos:  position{line: 321, col: 48, offset: 12849},
										name: "WS",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "OrderedListItemContent",
			pos:  position{line: 325, col: 1, offset: 12947},
			expr: &actionExpr{
				pos: position{line: 325, col: 27, offset: 12973},
				run: (*parser).callonOrderedListItemContent1,
				expr: &labeledExpr{
					pos:   position{line: 325, col: 27, offset: 12973},
					label: "elements",
					expr: &seqExpr{
						pos: position{line: 325, col: 37, offset: 12983},
						exprs: []interface{}{
							&oneOrMoreExpr{
								pos: position{line: 325, col: 37, offset: 12983},
								expr: &ruleRefExpr{
									pos:  position{line: 325, col: 37, offset: 12983},
									name: "ListParagraph",
								},
							},
							&zeroOrMoreExpr{
								pos: position{line: 325, col: 52, offset: 12998},
								expr: &ruleRefExpr{
									pos:  position{line: 325, col: 52, offset: 12998},
									name: "ContinuedBlockElement",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "UnorderedListItem",
			pos:  position{line: 332, col: 1, offset: 13324},
			expr: &actionExpr{
				pos: position{line: 332, col: 22, offset: 13345},
				run: (*parser).callonUnorderedListItem1,
				expr: &seqExpr{
					pos: position{line: 332, col: 22, offset: 13345},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 332, col: 22, offset: 13345},
							label: "prefix",
							expr: &ruleRefExpr{
								pos:  position{line: 332, col: 30, offset: 13353},
								name: "UnorderedListItemPrefix",
							},
						},
						&labeledExpr{
							pos:   position{line: 332, col: 55, offset: 13378},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 332, col: 64, offset: 13387},
								name: "UnorderedListItemContent",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 332, col: 90, offset: 13413},
							expr: &ruleRefExpr{
								pos:  position{line: 332, col: 90, offset: 13413},
								name: "BlankLine",
							},
						},
					},
				},
			},
		},
		{
			name: "UnorderedListItemPrefix",
			pos:  position{line: 336, col: 1, offset: 13532},
			expr: &choiceExpr{
				pos: position{line: 336, col: 28, offset: 13559},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 336, col: 28, offset: 13559},
						run: (*parser).callonUnorderedListItemPrefix2,
						expr: &seqExpr{
							pos: position{line: 336, col: 28, offset: 13559},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 336, col: 28, offset: 13559},
									expr: &ruleRefExpr{
										pos:  position{line: 336, col: 28, offset: 13559},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 336, col: 32, offset: 13563},
									label: "level",
									expr: &litMatcher{
										pos:        position{line: 336, col: 39, offset: 13570},
										val:        "*****",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 336, col: 48, offset: 13579},
									expr: &ruleRefExpr{
										pos:  position{line: 336, col: 48, offset: 13579},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 338, col: 5, offset: 13724},
						run: (*parser).callonUnorderedListItemPrefix10,
						expr: &seqExpr{
							pos: position{line: 338, col: 5, offset: 13724},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 338, col: 5, offset: 13724},
									expr: &ruleRefExpr{
										pos:  position{line: 338, col: 5, offset: 13724},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 338, col: 9, offset: 13728},
									label: "level",
									expr: &litMatcher{
										pos:        position{line: 338, col: 16, offset: 13735},
										val:        "****",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 338, col: 24, offset: 13743},
									expr: &ruleRefExpr{
										pos:  position{line: 338, col: 24, offset: 13743},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 340, col: 5, offset: 13888},
						run: (*parser).callonUnorderedListItemPrefix18,
						expr: &seqExpr{
							pos: position{line: 340, col: 5, offset: 13888},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 340, col: 5, offset: 13888},
									expr: &ruleRefExpr{
										pos:  position{line: 340, col: 5, offset: 13888},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 340, col: 9, offset: 13892},
									label: "level",
									expr: &litMatcher{
										pos:        position{line: 340, col: 16, offset: 13899},
										val:        "***",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 340, col: 23, offset: 13906},
									expr: &ruleRefExpr{
										pos:  position{line: 340, col: 23, offset: 13906},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 342, col: 5, offset: 14052},
						run: (*parser).callonUnorderedListItemPrefix26,
						expr: &seqExpr{
							pos: position{line: 342, col: 5, offset: 14052},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 342, col: 5, offset: 14052},
									expr: &ruleRefExpr{
										pos:  position{line: 342, col: 5, offset: 14052},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 342, col: 9, offset: 14056},
									label: "level",
									expr: &litMatcher{
										pos:        position{line: 342, col: 16, offset: 14063},
										val:        "**",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 342, col: 22, offset: 14069},
									expr: &ruleRefExpr{
										pos:  position{line: 342, col: 22, offset: 14069},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 344, col: 5, offset: 14213},
						run: (*parser).callonUnorderedListItemPrefix34,
						expr: &seqExpr{
							pos: position{line: 344, col: 5, offset: 14213},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 344, col: 5, offset: 14213},
									expr: &ruleRefExpr{
										pos:  position{line: 344, col: 5, offset: 14213},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 344, col: 9, offset: 14217},
									label: "level",
									expr: &litMatcher{
										pos:        position{line: 344, col: 16, offset: 14224},
										val:        "*",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 344, col: 21, offset: 14229},
									expr: &ruleRefExpr{
										pos:  position{line: 344, col: 21, offset: 14229},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 346, col: 5, offset: 14372},
						run: (*parser).callonUnorderedListItemPrefix42,
						expr: &seqExpr{
							pos: position{line: 346, col: 5, offset: 14372},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 346, col: 5, offset: 14372},
									expr: &ruleRefExpr{
										pos:  position{line: 346, col: 5, offset: 14372},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 346, col: 9, offset: 14376},
									label: "level",
									expr: &litMatcher{
										pos:        position{line: 346, col: 16, offset: 14383},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 346, col: 21, offset: 14388},
									expr: &ruleRefExpr{
										pos:  position{line: 346, col: 21, offset: 14388},
										name: "WS",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "UnorderedListItemContent",
			pos:  position{line: 350, col: 1, offset: 14524},
			expr: &actionExpr{
				pos: position{line: 350, col: 29, offset: 14552},
				run: (*parser).callonUnorderedListItemContent1,
				expr: &labeledExpr{
					pos:   position{line: 350, col: 29, offset: 14552},
					label: "elements",
					expr: &seqExpr{
						pos: position{line: 350, col: 39, offset: 14562},
						exprs: []interface{}{
							&oneOrMoreExpr{
								pos: position{line: 350, col: 39, offset: 14562},
								expr: &ruleRefExpr{
									pos:  position{line: 350, col: 39, offset: 14562},
									name: "ListParagraph",
								},
							},
							&zeroOrMoreExpr{
								pos: position{line: 350, col: 54, offset: 14577},
								expr: &ruleRefExpr{
									pos:  position{line: 350, col: 54, offset: 14577},
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
			pos:  position{line: 357, col: 1, offset: 14901},
			expr: &choiceExpr{
				pos: position{line: 357, col: 20, offset: 14920},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 357, col: 20, offset: 14920},
						run: (*parser).callonLabeledListItem2,
						expr: &seqExpr{
							pos: position{line: 357, col: 20, offset: 14920},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 357, col: 20, offset: 14920},
									label: "term",
									expr: &ruleRefExpr{
										pos:  position{line: 357, col: 26, offset: 14926},
										name: "LabeledListItemTerm",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 357, col: 47, offset: 14947},
									name: "LabeledListItemSeparator",
								},
								&labeledExpr{
									pos:   position{line: 357, col: 72, offset: 14972},
									label: "description",
									expr: &ruleRefExpr{
										pos:  position{line: 357, col: 85, offset: 14985},
										name: "LabeledListItemDescription",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 359, col: 6, offset: 15107},
						run: (*parser).callonLabeledListItem9,
						expr: &seqExpr{
							pos: position{line: 359, col: 6, offset: 15107},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 359, col: 6, offset: 15107},
									label: "term",
									expr: &ruleRefExpr{
										pos:  position{line: 359, col: 12, offset: 15113},
										name: "LabeledListItemTerm",
									},
								},
								&litMatcher{
									pos:        position{line: 359, col: 33, offset: 15134},
									val:        "::",
									ignoreCase: false,
								},
								&zeroOrMoreExpr{
									pos: position{line: 359, col: 38, offset: 15139},
									expr: &ruleRefExpr{
										pos:  position{line: 359, col: 38, offset: 15139},
										name: "WS",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 359, col: 42, offset: 15143},
									name: "EOL",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "LabeledListItemTerm",
			pos:  position{line: 363, col: 1, offset: 15280},
			expr: &actionExpr{
				pos: position{line: 363, col: 24, offset: 15303},
				run: (*parser).callonLabeledListItemTerm1,
				expr: &labeledExpr{
					pos:   position{line: 363, col: 24, offset: 15303},
					label: "term",
					expr: &zeroOrMoreExpr{
						pos: position{line: 363, col: 29, offset: 15308},
						expr: &seqExpr{
							pos: position{line: 363, col: 30, offset: 15309},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 363, col: 30, offset: 15309},
									expr: &ruleRefExpr{
										pos:  position{line: 363, col: 31, offset: 15310},
										name: "NEWLINE",
									},
								},
								&notExpr{
									pos: position{line: 363, col: 39, offset: 15318},
									expr: &litMatcher{
										pos:        position{line: 363, col: 40, offset: 15319},
										val:        "::",
										ignoreCase: false,
									},
								},
								&anyMatcher{
									line: 363, col: 45, offset: 15324,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "LabeledListItemSeparator",
			pos:  position{line: 368, col: 1, offset: 15415},
			expr: &seqExpr{
				pos: position{line: 368, col: 30, offset: 15444},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 368, col: 30, offset: 15444},
						val:        "::",
						ignoreCase: false,
					},
					&oneOrMoreExpr{
						pos: position{line: 368, col: 35, offset: 15449},
						expr: &choiceExpr{
							pos: position{line: 368, col: 36, offset: 15450},
							alternatives: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 368, col: 36, offset: 15450},
									name: "WS",
								},
								&ruleRefExpr{
									pos:  position{line: 368, col: 41, offset: 15455},
									name: "NEWLINE",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "LabeledListItemDescription",
			pos:  position{line: 370, col: 1, offset: 15466},
			expr: &actionExpr{
				pos: position{line: 370, col: 31, offset: 15496},
				run: (*parser).callonLabeledListItemDescription1,
				expr: &labeledExpr{
					pos:   position{line: 370, col: 31, offset: 15496},
					label: "elements",
					expr: &zeroOrMoreExpr{
						pos: position{line: 370, col: 40, offset: 15505},
						expr: &choiceExpr{
							pos: position{line: 370, col: 41, offset: 15506},
							alternatives: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 370, col: 41, offset: 15506},
									name: "ListParagraph",
								},
								&ruleRefExpr{
									pos:  position{line: 370, col: 57, offset: 15522},
									name: "ContinuedBlockElement",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "AdmonitionKind",
			pos:  position{line: 378, col: 1, offset: 15829},
			expr: &choiceExpr{
				pos: position{line: 378, col: 19, offset: 15847},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 378, col: 19, offset: 15847},
						run: (*parser).callonAdmonitionKind2,
						expr: &litMatcher{
							pos:        position{line: 378, col: 19, offset: 15847},
							val:        "TIP",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 380, col: 5, offset: 15885},
						run: (*parser).callonAdmonitionKind4,
						expr: &litMatcher{
							pos:        position{line: 380, col: 5, offset: 15885},
							val:        "NOTE",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 382, col: 5, offset: 15925},
						run: (*parser).callonAdmonitionKind6,
						expr: &litMatcher{
							pos:        position{line: 382, col: 5, offset: 15925},
							val:        "IMPORTANT",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 384, col: 5, offset: 15975},
						run: (*parser).callonAdmonitionKind8,
						expr: &litMatcher{
							pos:        position{line: 384, col: 5, offset: 15975},
							val:        "WARNING",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 386, col: 5, offset: 16021},
						run: (*parser).callonAdmonitionKind10,
						expr: &litMatcher{
							pos:        position{line: 386, col: 5, offset: 16021},
							val:        "CAUTION",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Paragraph",
			pos:  position{line: 395, col: 1, offset: 16324},
			expr: &choiceExpr{
				pos: position{line: 395, col: 14, offset: 16337},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 395, col: 14, offset: 16337},
						run: (*parser).callonParagraph2,
						expr: &seqExpr{
							pos: position{line: 395, col: 14, offset: 16337},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 395, col: 14, offset: 16337},
									label: "attributes",
									expr: &zeroOrMoreExpr{
										pos: position{line: 395, col: 25, offset: 16348},
										expr: &ruleRefExpr{
											pos:  position{line: 395, col: 26, offset: 16349},
											name: "ElementAttribute",
										},
									},
								},
								&notExpr{
									pos: position{line: 395, col: 45, offset: 16368},
									expr: &seqExpr{
										pos: position{line: 395, col: 47, offset: 16370},
										exprs: []interface{}{
											&oneOrMoreExpr{
												pos: position{line: 395, col: 47, offset: 16370},
												expr: &litMatcher{
													pos:        position{line: 395, col: 47, offset: 16370},
													val:        "=",
													ignoreCase: false,
												},
											},
											&oneOrMoreExpr{
												pos: position{line: 395, col: 52, offset: 16375},
												expr: &ruleRefExpr{
													pos:  position{line: 395, col: 52, offset: 16375},
													name: "WS",
												},
											},
										},
									},
								},
								&labeledExpr{
									pos:   position{line: 395, col: 57, offset: 16380},
									label: "t",
									expr: &ruleRefExpr{
										pos:  position{line: 395, col: 60, offset: 16383},
										name: "AdmonitionKind",
									},
								},
								&litMatcher{
									pos:        position{line: 395, col: 76, offset: 16399},
									val:        ": ",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 395, col: 81, offset: 16404},
									label: "lines",
									expr: &oneOrMoreExpr{
										pos: position{line: 395, col: 87, offset: 16410},
										expr: &seqExpr{
											pos: position{line: 395, col: 88, offset: 16411},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 395, col: 88, offset: 16411},
													name: "InlineElementsWithTrailingSpaces",
												},
												&ruleRefExpr{
													pos:  position{line: 395, col: 121, offset: 16444},
													name: "EOL",
												},
											},
										},
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 397, col: 5, offset: 16596},
						run: (*parser).callonParagraph21,
						expr: &seqExpr{
							pos: position{line: 397, col: 5, offset: 16596},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 397, col: 5, offset: 16596},
									label: "attributes",
									expr: &zeroOrMoreExpr{
										pos: position{line: 397, col: 16, offset: 16607},
										expr: &ruleRefExpr{
											pos:  position{line: 397, col: 17, offset: 16608},
											name: "ElementAttribute",
										},
									},
								},
								&notExpr{
									pos: position{line: 397, col: 36, offset: 16627},
									expr: &seqExpr{
										pos: position{line: 397, col: 38, offset: 16629},
										exprs: []interface{}{
											&oneOrMoreExpr{
												pos: position{line: 397, col: 38, offset: 16629},
												expr: &litMatcher{
													pos:        position{line: 397, col: 38, offset: 16629},
													val:        "=",
													ignoreCase: false,
												},
											},
											&oneOrMoreExpr{
												pos: position{line: 397, col: 43, offset: 16634},
												expr: &ruleRefExpr{
													pos:  position{line: 397, col: 43, offset: 16634},
													name: "WS",
												},
											},
										},
									},
								},
								&labeledExpr{
									pos:   position{line: 397, col: 48, offset: 16639},
									label: "lines",
									expr: &oneOrMoreExpr{
										pos: position{line: 397, col: 54, offset: 16645},
										expr: &seqExpr{
											pos: position{line: 397, col: 55, offset: 16646},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 397, col: 55, offset: 16646},
													name: "InlineElementsWithTrailingSpaces",
												},
												&ruleRefExpr{
													pos:  position{line: 397, col: 88, offset: 16679},
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
			},
		},
		{
			name: "InlineElementsWithTrailingSpaces",
			pos:  position{line: 403, col: 1, offset: 16990},
			expr: &actionExpr{
				pos: position{line: 403, col: 37, offset: 17026},
				run: (*parser).callonInlineElementsWithTrailingSpaces1,
				expr: &seqExpr{
					pos: position{line: 403, col: 37, offset: 17026},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 403, col: 37, offset: 17026},
							expr: &ruleRefExpr{
								pos:  position{line: 403, col: 38, offset: 17027},
								name: "BlockDelimiter",
							},
						},
						&labeledExpr{
							pos:   position{line: 403, col: 53, offset: 17042},
							label: "elements",
							expr: &oneOrMoreExpr{
								pos: position{line: 403, col: 62, offset: 17051},
								expr: &seqExpr{
									pos: position{line: 403, col: 63, offset: 17052},
									exprs: []interface{}{
										&zeroOrMoreExpr{
											pos: position{line: 403, col: 63, offset: 17052},
											expr: &ruleRefExpr{
												pos:  position{line: 403, col: 63, offset: 17052},
												name: "WS",
											},
										},
										&notExpr{
											pos: position{line: 403, col: 67, offset: 17056},
											expr: &ruleRefExpr{
												pos:  position{line: 403, col: 68, offset: 17057},
												name: "InlineElementID",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 403, col: 84, offset: 17073},
											name: "InlineElement",
										},
										&zeroOrMoreExpr{
											pos: position{line: 403, col: 98, offset: 17087},
											expr: &ruleRefExpr{
												pos:  position{line: 403, col: 98, offset: 17087},
												name: "WS",
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
			name: "InlineElements",
			pos:  position{line: 407, col: 1, offset: 17221},
			expr: &actionExpr{
				pos: position{line: 407, col: 19, offset: 17239},
				run: (*parser).callonInlineElements1,
				expr: &seqExpr{
					pos: position{line: 407, col: 19, offset: 17239},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 407, col: 19, offset: 17239},
							expr: &ruleRefExpr{
								pos:  position{line: 407, col: 20, offset: 17240},
								name: "BlockDelimiter",
							},
						},
						&labeledExpr{
							pos:   position{line: 407, col: 35, offset: 17255},
							label: "elements",
							expr: &oneOrMoreExpr{
								pos: position{line: 407, col: 44, offset: 17264},
								expr: &seqExpr{
									pos: position{line: 407, col: 45, offset: 17265},
									exprs: []interface{}{
										&zeroOrMoreExpr{
											pos: position{line: 407, col: 45, offset: 17265},
											expr: &ruleRefExpr{
												pos:  position{line: 407, col: 45, offset: 17265},
												name: "WS",
											},
										},
										&notExpr{
											pos: position{line: 407, col: 49, offset: 17269},
											expr: &ruleRefExpr{
												pos:  position{line: 407, col: 50, offset: 17270},
												name: "InlineElementID",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 407, col: 66, offset: 17286},
											name: "InlineElement",
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
			name: "InlineElement",
			pos:  position{line: 411, col: 1, offset: 17409},
			expr: &choiceExpr{
				pos: position{line: 411, col: 18, offset: 17426},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 411, col: 18, offset: 17426},
						name: "CrossReference",
					},
					&ruleRefExpr{
						pos:  position{line: 411, col: 35, offset: 17443},
						name: "Passthrough",
					},
					&ruleRefExpr{
						pos:  position{line: 411, col: 49, offset: 17457},
						name: "InlineImage",
					},
					&ruleRefExpr{
						pos:  position{line: 411, col: 63, offset: 17471},
						name: "QuotedText",
					},
					&ruleRefExpr{
						pos:  position{line: 411, col: 76, offset: 17484},
						name: "Link",
					},
					&ruleRefExpr{
						pos:  position{line: 411, col: 83, offset: 17491},
						name: "DocumentAttributeSubstitution",
					},
					&ruleRefExpr{
						pos:  position{line: 411, col: 115, offset: 17523},
						name: "Characters",
					},
				},
			},
		},
		{
			name: "QuotedText",
			pos:  position{line: 416, col: 1, offset: 17774},
			expr: &choiceExpr{
				pos: position{line: 416, col: 15, offset: 17788},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 416, col: 15, offset: 17788},
						name: "BoldText",
					},
					&ruleRefExpr{
						pos:  position{line: 416, col: 26, offset: 17799},
						name: "ItalicText",
					},
					&ruleRefExpr{
						pos:  position{line: 416, col: 39, offset: 17812},
						name: "MonospaceText",
					},
					&ruleRefExpr{
						pos:  position{line: 417, col: 13, offset: 17840},
						name: "EscapedBoldText",
					},
					&ruleRefExpr{
						pos:  position{line: 417, col: 31, offset: 17858},
						name: "EscapedItalicText",
					},
					&ruleRefExpr{
						pos:  position{line: 417, col: 51, offset: 17878},
						name: "EscapedMonospaceText",
					},
				},
			},
		},
		{
			name: "BoldText",
			pos:  position{line: 419, col: 1, offset: 17900},
			expr: &choiceExpr{
				pos: position{line: 419, col: 13, offset: 17912},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 419, col: 13, offset: 17912},
						run: (*parser).callonBoldText2,
						expr: &seqExpr{
							pos: position{line: 419, col: 13, offset: 17912},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 419, col: 13, offset: 17912},
									expr: &litMatcher{
										pos:        position{line: 419, col: 14, offset: 17913},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 419, col: 19, offset: 17918},
									val:        "**",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 419, col: 24, offset: 17923},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 419, col: 33, offset: 17932},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 419, col: 52, offset: 17951},
									val:        "**",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 421, col: 5, offset: 18076},
						run: (*parser).callonBoldText10,
						expr: &seqExpr{
							pos: position{line: 421, col: 5, offset: 18076},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 421, col: 5, offset: 18076},
									expr: &litMatcher{
										pos:        position{line: 421, col: 6, offset: 18077},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 421, col: 11, offset: 18082},
									val:        "**",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 421, col: 16, offset: 18087},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 421, col: 25, offset: 18096},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 421, col: 44, offset: 18115},
									val:        "*",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 424, col: 5, offset: 18280},
						run: (*parser).callonBoldText18,
						expr: &seqExpr{
							pos: position{line: 424, col: 5, offset: 18280},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 424, col: 5, offset: 18280},
									expr: &litMatcher{
										pos:        position{line: 424, col: 6, offset: 18281},
										val:        "\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 424, col: 10, offset: 18285},
									val:        "*",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 424, col: 14, offset: 18289},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 424, col: 23, offset: 18298},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 424, col: 42, offset: 18317},
									val:        "*",
									ignoreCase: false,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "EscapedBoldText",
			pos:  position{line: 428, col: 1, offset: 18417},
			expr: &choiceExpr{
				pos: position{line: 428, col: 20, offset: 18436},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 428, col: 20, offset: 18436},
						run: (*parser).callonEscapedBoldText2,
						expr: &seqExpr{
							pos: position{line: 428, col: 20, offset: 18436},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 428, col: 20, offset: 18436},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 428, col: 33, offset: 18449},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 428, col: 33, offset: 18449},
												val:        "\\\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 428, col: 38, offset: 18454},
												expr: &litMatcher{
													pos:        position{line: 428, col: 38, offset: 18454},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 428, col: 44, offset: 18460},
									val:        "**",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 428, col: 49, offset: 18465},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 428, col: 58, offset: 18474},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 428, col: 77, offset: 18493},
									val:        "**",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 430, col: 5, offset: 18648},
						run: (*parser).callonEscapedBoldText13,
						expr: &seqExpr{
							pos: position{line: 430, col: 5, offset: 18648},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 430, col: 5, offset: 18648},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 430, col: 18, offset: 18661},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 430, col: 18, offset: 18661},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 430, col: 22, offset: 18665},
												expr: &litMatcher{
													pos:        position{line: 430, col: 22, offset: 18665},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 430, col: 28, offset: 18671},
									val:        "**",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 430, col: 33, offset: 18676},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 430, col: 42, offset: 18685},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 430, col: 61, offset: 18704},
									val:        "*",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 433, col: 5, offset: 18898},
						run: (*parser).callonEscapedBoldText24,
						expr: &seqExpr{
							pos: position{line: 433, col: 5, offset: 18898},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 433, col: 5, offset: 18898},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 433, col: 18, offset: 18911},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 433, col: 18, offset: 18911},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 433, col: 22, offset: 18915},
												expr: &litMatcher{
													pos:        position{line: 433, col: 22, offset: 18915},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 433, col: 28, offset: 18921},
									val:        "*",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 433, col: 32, offset: 18925},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 433, col: 41, offset: 18934},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 433, col: 60, offset: 18953},
									val:        "*",
									ignoreCase: false,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "ItalicText",
			pos:  position{line: 437, col: 1, offset: 19105},
			expr: &choiceExpr{
				pos: position{line: 437, col: 15, offset: 19119},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 437, col: 15, offset: 19119},
						run: (*parser).callonItalicText2,
						expr: &seqExpr{
							pos: position{line: 437, col: 15, offset: 19119},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 437, col: 15, offset: 19119},
									expr: &litMatcher{
										pos:        position{line: 437, col: 16, offset: 19120},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 437, col: 21, offset: 19125},
									val:        "__",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 437, col: 26, offset: 19130},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 437, col: 35, offset: 19139},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 437, col: 54, offset: 19158},
									val:        "__",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 439, col: 5, offset: 19239},
						run: (*parser).callonItalicText10,
						expr: &seqExpr{
							pos: position{line: 439, col: 5, offset: 19239},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 439, col: 5, offset: 19239},
									expr: &litMatcher{
										pos:        position{line: 439, col: 6, offset: 19240},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 439, col: 11, offset: 19245},
									val:        "__",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 439, col: 16, offset: 19250},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 439, col: 25, offset: 19259},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 439, col: 44, offset: 19278},
									val:        "_",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 442, col: 5, offset: 19445},
						run: (*parser).callonItalicText18,
						expr: &seqExpr{
							pos: position{line: 442, col: 5, offset: 19445},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 442, col: 5, offset: 19445},
									expr: &litMatcher{
										pos:        position{line: 442, col: 6, offset: 19446},
										val:        "\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 442, col: 10, offset: 19450},
									val:        "_",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 442, col: 14, offset: 19454},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 442, col: 23, offset: 19463},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 442, col: 42, offset: 19482},
									val:        "_",
									ignoreCase: false,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "EscapedItalicText",
			pos:  position{line: 446, col: 1, offset: 19561},
			expr: &choiceExpr{
				pos: position{line: 446, col: 22, offset: 19582},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 446, col: 22, offset: 19582},
						run: (*parser).callonEscapedItalicText2,
						expr: &seqExpr{
							pos: position{line: 446, col: 22, offset: 19582},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 446, col: 22, offset: 19582},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 446, col: 35, offset: 19595},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 446, col: 35, offset: 19595},
												val:        "\\\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 446, col: 40, offset: 19600},
												expr: &litMatcher{
													pos:        position{line: 446, col: 40, offset: 19600},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 446, col: 46, offset: 19606},
									val:        "__",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 446, col: 51, offset: 19611},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 446, col: 60, offset: 19620},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 446, col: 79, offset: 19639},
									val:        "__",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 448, col: 5, offset: 19794},
						run: (*parser).callonEscapedItalicText13,
						expr: &seqExpr{
							pos: position{line: 448, col: 5, offset: 19794},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 448, col: 5, offset: 19794},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 448, col: 18, offset: 19807},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 448, col: 18, offset: 19807},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 448, col: 22, offset: 19811},
												expr: &litMatcher{
													pos:        position{line: 448, col: 22, offset: 19811},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 448, col: 28, offset: 19817},
									val:        "__",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 448, col: 33, offset: 19822},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 448, col: 42, offset: 19831},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 448, col: 61, offset: 19850},
									val:        "_",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 451, col: 5, offset: 20044},
						run: (*parser).callonEscapedItalicText24,
						expr: &seqExpr{
							pos: position{line: 451, col: 5, offset: 20044},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 451, col: 5, offset: 20044},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 451, col: 18, offset: 20057},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 451, col: 18, offset: 20057},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 451, col: 22, offset: 20061},
												expr: &litMatcher{
													pos:        position{line: 451, col: 22, offset: 20061},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 451, col: 28, offset: 20067},
									val:        "_",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 451, col: 32, offset: 20071},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 451, col: 41, offset: 20080},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 451, col: 60, offset: 20099},
									val:        "_",
									ignoreCase: false,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "MonospaceText",
			pos:  position{line: 455, col: 1, offset: 20251},
			expr: &choiceExpr{
				pos: position{line: 455, col: 18, offset: 20268},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 455, col: 18, offset: 20268},
						run: (*parser).callonMonospaceText2,
						expr: &seqExpr{
							pos: position{line: 455, col: 18, offset: 20268},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 455, col: 18, offset: 20268},
									expr: &litMatcher{
										pos:        position{line: 455, col: 19, offset: 20269},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 455, col: 24, offset: 20274},
									val:        "``",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 455, col: 29, offset: 20279},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 455, col: 38, offset: 20288},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 455, col: 57, offset: 20307},
									val:        "``",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 457, col: 5, offset: 20437},
						run: (*parser).callonMonospaceText10,
						expr: &seqExpr{
							pos: position{line: 457, col: 5, offset: 20437},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 457, col: 5, offset: 20437},
									expr: &litMatcher{
										pos:        position{line: 457, col: 6, offset: 20438},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 457, col: 11, offset: 20443},
									val:        "``",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 457, col: 16, offset: 20448},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 457, col: 25, offset: 20457},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 457, col: 44, offset: 20476},
									val:        "`",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 460, col: 5, offset: 20646},
						run: (*parser).callonMonospaceText18,
						expr: &seqExpr{
							pos: position{line: 460, col: 5, offset: 20646},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 460, col: 5, offset: 20646},
									expr: &litMatcher{
										pos:        position{line: 460, col: 6, offset: 20647},
										val:        "\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 460, col: 10, offset: 20651},
									val:        "`",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 460, col: 14, offset: 20655},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 460, col: 23, offset: 20664},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 460, col: 42, offset: 20683},
									val:        "`",
									ignoreCase: false,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "EscapedMonospaceText",
			pos:  position{line: 464, col: 1, offset: 20810},
			expr: &choiceExpr{
				pos: position{line: 464, col: 25, offset: 20834},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 464, col: 25, offset: 20834},
						run: (*parser).callonEscapedMonospaceText2,
						expr: &seqExpr{
							pos: position{line: 464, col: 25, offset: 20834},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 464, col: 25, offset: 20834},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 464, col: 38, offset: 20847},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 464, col: 38, offset: 20847},
												val:        "\\\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 464, col: 43, offset: 20852},
												expr: &litMatcher{
													pos:        position{line: 464, col: 43, offset: 20852},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 464, col: 49, offset: 20858},
									val:        "``",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 464, col: 54, offset: 20863},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 464, col: 63, offset: 20872},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 464, col: 82, offset: 20891},
									val:        "``",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 466, col: 5, offset: 21046},
						run: (*parser).callonEscapedMonospaceText13,
						expr: &seqExpr{
							pos: position{line: 466, col: 5, offset: 21046},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 466, col: 5, offset: 21046},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 466, col: 18, offset: 21059},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 466, col: 18, offset: 21059},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 466, col: 22, offset: 21063},
												expr: &litMatcher{
													pos:        position{line: 466, col: 22, offset: 21063},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 466, col: 28, offset: 21069},
									val:        "``",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 466, col: 33, offset: 21074},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 466, col: 42, offset: 21083},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 466, col: 61, offset: 21102},
									val:        "`",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 469, col: 5, offset: 21296},
						run: (*parser).callonEscapedMonospaceText24,
						expr: &seqExpr{
							pos: position{line: 469, col: 5, offset: 21296},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 469, col: 5, offset: 21296},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 469, col: 18, offset: 21309},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 469, col: 18, offset: 21309},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 469, col: 22, offset: 21313},
												expr: &litMatcher{
													pos:        position{line: 469, col: 22, offset: 21313},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 469, col: 28, offset: 21319},
									val:        "`",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 469, col: 32, offset: 21323},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 469, col: 41, offset: 21332},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 469, col: 60, offset: 21351},
									val:        "`",
									ignoreCase: false,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "QuotedTextContent",
			pos:  position{line: 473, col: 1, offset: 21503},
			expr: &seqExpr{
				pos: position{line: 473, col: 22, offset: 21524},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 473, col: 22, offset: 21524},
						name: "QuotedTextContentElement",
					},
					&zeroOrMoreExpr{
						pos: position{line: 473, col: 47, offset: 21549},
						expr: &seqExpr{
							pos: position{line: 473, col: 48, offset: 21550},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 473, col: 48, offset: 21550},
									expr: &ruleRefExpr{
										pos:  position{line: 473, col: 48, offset: 21550},
										name: "WS",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 473, col: 52, offset: 21554},
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
			pos:  position{line: 475, col: 1, offset: 21582},
			expr: &choiceExpr{
				pos: position{line: 475, col: 29, offset: 21610},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 475, col: 29, offset: 21610},
						name: "QuotedText",
					},
					&ruleRefExpr{
						pos:  position{line: 475, col: 42, offset: 21623},
						name: "QuotedTextCharacters",
					},
					&ruleRefExpr{
						pos:  position{line: 475, col: 65, offset: 21646},
						name: "CharactersWithQuotePunctuation",
					},
				},
			},
		},
		{
			name: "QuotedTextCharacters",
			pos:  position{line: 477, col: 1, offset: 21781},
			expr: &oneOrMoreExpr{
				pos: position{line: 477, col: 25, offset: 21805},
				expr: &seqExpr{
					pos: position{line: 477, col: 26, offset: 21806},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 477, col: 26, offset: 21806},
							expr: &ruleRefExpr{
								pos:  position{line: 477, col: 27, offset: 21807},
								name: "NEWLINE",
							},
						},
						&notExpr{
							pos: position{line: 477, col: 35, offset: 21815},
							expr: &ruleRefExpr{
								pos:  position{line: 477, col: 36, offset: 21816},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 477, col: 39, offset: 21819},
							expr: &litMatcher{
								pos:        position{line: 477, col: 40, offset: 21820},
								val:        "*",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 477, col: 44, offset: 21824},
							expr: &litMatcher{
								pos:        position{line: 477, col: 45, offset: 21825},
								val:        "_",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 477, col: 49, offset: 21829},
							expr: &litMatcher{
								pos:        position{line: 477, col: 50, offset: 21830},
								val:        "`",
								ignoreCase: false,
							},
						},
						&anyMatcher{
							line: 477, col: 54, offset: 21834,
						},
					},
				},
			},
		},
		{
			name: "CharactersWithQuotePunctuation",
			pos:  position{line: 479, col: 1, offset: 21877},
			expr: &actionExpr{
				pos: position{line: 479, col: 35, offset: 21911},
				run: (*parser).callonCharactersWithQuotePunctuation1,
				expr: &oneOrMoreExpr{
					pos: position{line: 479, col: 35, offset: 21911},
					expr: &seqExpr{
						pos: position{line: 479, col: 36, offset: 21912},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 479, col: 36, offset: 21912},
								expr: &ruleRefExpr{
									pos:  position{line: 479, col: 37, offset: 21913},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 479, col: 45, offset: 21921},
								expr: &ruleRefExpr{
									pos:  position{line: 479, col: 46, offset: 21922},
									name: "WS",
								},
							},
							&anyMatcher{
								line: 479, col: 50, offset: 21926,
							},
						},
					},
				},
			},
		},
		{
			name: "UnbalancedQuotePunctuation",
			pos:  position{line: 484, col: 1, offset: 22171},
			expr: &choiceExpr{
				pos: position{line: 484, col: 31, offset: 22201},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 484, col: 31, offset: 22201},
						val:        "*",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 484, col: 37, offset: 22207},
						val:        "_",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 484, col: 43, offset: 22213},
						val:        "`",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "Passthrough",
			pos:  position{line: 489, col: 1, offset: 22325},
			expr: &choiceExpr{
				pos: position{line: 489, col: 16, offset: 22340},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 489, col: 16, offset: 22340},
						name: "TriplePlusPassthrough",
					},
					&ruleRefExpr{
						pos:  position{line: 489, col: 40, offset: 22364},
						name: "SinglePlusPassthrough",
					},
					&ruleRefExpr{
						pos:  position{line: 489, col: 64, offset: 22388},
						name: "PassthroughMacro",
					},
				},
			},
		},
		{
			name: "SinglePlusPassthrough",
			pos:  position{line: 491, col: 1, offset: 22406},
			expr: &actionExpr{
				pos: position{line: 491, col: 26, offset: 22431},
				run: (*parser).callonSinglePlusPassthrough1,
				expr: &seqExpr{
					pos: position{line: 491, col: 26, offset: 22431},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 491, col: 26, offset: 22431},
							val:        "+",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 491, col: 30, offset: 22435},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 491, col: 38, offset: 22443},
								expr: &seqExpr{
									pos: position{line: 491, col: 39, offset: 22444},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 491, col: 39, offset: 22444},
											expr: &ruleRefExpr{
												pos:  position{line: 491, col: 40, offset: 22445},
												name: "NEWLINE",
											},
										},
										&notExpr{
											pos: position{line: 491, col: 48, offset: 22453},
											expr: &litMatcher{
												pos:        position{line: 491, col: 49, offset: 22454},
												val:        "+",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 491, col: 53, offset: 22458,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 491, col: 57, offset: 22462},
							val:        "+",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TriplePlusPassthrough",
			pos:  position{line: 495, col: 1, offset: 22557},
			expr: &actionExpr{
				pos: position{line: 495, col: 26, offset: 22582},
				run: (*parser).callonTriplePlusPassthrough1,
				expr: &seqExpr{
					pos: position{line: 495, col: 26, offset: 22582},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 495, col: 26, offset: 22582},
							val:        "+++",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 495, col: 32, offset: 22588},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 495, col: 40, offset: 22596},
								expr: &seqExpr{
									pos: position{line: 495, col: 41, offset: 22597},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 495, col: 41, offset: 22597},
											expr: &litMatcher{
												pos:        position{line: 495, col: 42, offset: 22598},
												val:        "+++",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 495, col: 48, offset: 22604,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 495, col: 52, offset: 22608},
							val:        "+++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "PassthroughMacro",
			pos:  position{line: 499, col: 1, offset: 22705},
			expr: &choiceExpr{
				pos: position{line: 499, col: 21, offset: 22725},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 499, col: 21, offset: 22725},
						run: (*parser).callonPassthroughMacro2,
						expr: &seqExpr{
							pos: position{line: 499, col: 21, offset: 22725},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 499, col: 21, offset: 22725},
									val:        "pass:[",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 499, col: 30, offset: 22734},
									label: "content",
									expr: &zeroOrMoreExpr{
										pos: position{line: 499, col: 38, offset: 22742},
										expr: &ruleRefExpr{
											pos:  position{line: 499, col: 39, offset: 22743},
											name: "PassthroughMacroCharacter",
										},
									},
								},
								&litMatcher{
									pos:        position{line: 499, col: 67, offset: 22771},
									val:        "]",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 501, col: 5, offset: 22862},
						run: (*parser).callonPassthroughMacro9,
						expr: &seqExpr{
							pos: position{line: 501, col: 5, offset: 22862},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 501, col: 5, offset: 22862},
									val:        "pass:q[",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 501, col: 15, offset: 22872},
									label: "content",
									expr: &zeroOrMoreExpr{
										pos: position{line: 501, col: 23, offset: 22880},
										expr: &choiceExpr{
											pos: position{line: 501, col: 24, offset: 22881},
											alternatives: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 501, col: 24, offset: 22881},
													name: "QuotedText",
												},
												&ruleRefExpr{
													pos:  position{line: 501, col: 37, offset: 22894},
													name: "PassthroughMacroCharacter",
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 501, col: 65, offset: 22922},
									val:        "]",
									ignoreCase: false,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "PassthroughMacroCharacter",
			pos:  position{line: 505, col: 1, offset: 23012},
			expr: &seqExpr{
				pos: position{line: 505, col: 31, offset: 23042},
				exprs: []interface{}{
					&notExpr{
						pos: position{line: 505, col: 31, offset: 23042},
						expr: &litMatcher{
							pos:        position{line: 505, col: 32, offset: 23043},
							val:        "]",
							ignoreCase: false,
						},
					},
					&anyMatcher{
						line: 505, col: 36, offset: 23047,
					},
				},
			},
		},
		{
			name: "CrossReference",
			pos:  position{line: 510, col: 1, offset: 23163},
			expr: &actionExpr{
				pos: position{line: 510, col: 19, offset: 23181},
				run: (*parser).callonCrossReference1,
				expr: &seqExpr{
					pos: position{line: 510, col: 19, offset: 23181},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 510, col: 19, offset: 23181},
							val:        "<<",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 510, col: 24, offset: 23186},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 510, col: 28, offset: 23190},
								name: "ID",
							},
						},
						&litMatcher{
							pos:        position{line: 510, col: 32, offset: 23194},
							val:        ">>",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Link",
			pos:  position{line: 517, col: 1, offset: 23353},
			expr: &choiceExpr{
				pos: position{line: 517, col: 9, offset: 23361},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 517, col: 9, offset: 23361},
						name: "RelativeLink",
					},
					&ruleRefExpr{
						pos:  position{line: 517, col: 24, offset: 23376},
						name: "ExternalLink",
					},
				},
			},
		},
		{
			name: "ExternalLink",
			pos:  position{line: 519, col: 1, offset: 23391},
			expr: &actionExpr{
				pos: position{line: 519, col: 17, offset: 23407},
				run: (*parser).callonExternalLink1,
				expr: &seqExpr{
					pos: position{line: 519, col: 17, offset: 23407},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 519, col: 17, offset: 23407},
							label: "url",
							expr: &seqExpr{
								pos: position{line: 519, col: 22, offset: 23412},
								exprs: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 519, col: 22, offset: 23412},
										name: "URL_SCHEME",
									},
									&ruleRefExpr{
										pos:  position{line: 519, col: 33, offset: 23423},
										name: "URL",
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 519, col: 38, offset: 23428},
							label: "text",
							expr: &zeroOrOneExpr{
								pos: position{line: 519, col: 43, offset: 23433},
								expr: &seqExpr{
									pos: position{line: 519, col: 44, offset: 23434},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 519, col: 44, offset: 23434},
											val:        "[",
											ignoreCase: false,
										},
										&zeroOrMoreExpr{
											pos: position{line: 519, col: 48, offset: 23438},
											expr: &ruleRefExpr{
												pos:  position{line: 519, col: 49, offset: 23439},
												name: "URL_TEXT",
											},
										},
										&litMatcher{
											pos:        position{line: 519, col: 60, offset: 23450},
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
			name: "RelativeLink",
			pos:  position{line: 526, col: 1, offset: 23611},
			expr: &actionExpr{
				pos: position{line: 526, col: 17, offset: 23627},
				run: (*parser).callonRelativeLink1,
				expr: &seqExpr{
					pos: position{line: 526, col: 17, offset: 23627},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 526, col: 17, offset: 23627},
							val:        "link:",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 526, col: 25, offset: 23635},
							label: "url",
							expr: &seqExpr{
								pos: position{line: 526, col: 30, offset: 23640},
								exprs: []interface{}{
									&zeroOrOneExpr{
										pos: position{line: 526, col: 30, offset: 23640},
										expr: &ruleRefExpr{
											pos:  position{line: 526, col: 30, offset: 23640},
											name: "URL_SCHEME",
										},
									},
									&ruleRefExpr{
										pos:  position{line: 526, col: 42, offset: 23652},
										name: "URL",
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 526, col: 47, offset: 23657},
							label: "text",
							expr: &seqExpr{
								pos: position{line: 526, col: 53, offset: 23663},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 526, col: 53, offset: 23663},
										val:        "[",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 526, col: 57, offset: 23667},
										expr: &ruleRefExpr{
											pos:  position{line: 526, col: 58, offset: 23668},
											name: "URL_TEXT",
										},
									},
									&litMatcher{
										pos:        position{line: 526, col: 69, offset: 23679},
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
		{
			name: "BlockImage",
			pos:  position{line: 536, col: 1, offset: 23941},
			expr: &actionExpr{
				pos: position{line: 536, col: 15, offset: 23955},
				run: (*parser).callonBlockImage1,
				expr: &seqExpr{
					pos: position{line: 536, col: 15, offset: 23955},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 536, col: 15, offset: 23955},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 536, col: 26, offset: 23966},
								expr: &ruleRefExpr{
									pos:  position{line: 536, col: 27, offset: 23967},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 536, col: 46, offset: 23986},
							label: "image",
							expr: &ruleRefExpr{
								pos:  position{line: 536, col: 52, offset: 23992},
								name: "BlockImageMacro",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 536, col: 69, offset: 24009},
							expr: &ruleRefExpr{
								pos:  position{line: 536, col: 69, offset: 24009},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 536, col: 73, offset: 24013},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "BlockImageMacro",
			pos:  position{line: 541, col: 1, offset: 24172},
			expr: &actionExpr{
				pos: position{line: 541, col: 20, offset: 24191},
				run: (*parser).callonBlockImageMacro1,
				expr: &seqExpr{
					pos: position{line: 541, col: 20, offset: 24191},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 541, col: 20, offset: 24191},
							val:        "image::",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 541, col: 30, offset: 24201},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 541, col: 36, offset: 24207},
								name: "URL",
							},
						},
						&litMatcher{
							pos:        position{line: 541, col: 41, offset: 24212},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 541, col: 45, offset: 24216},
							label: "attributes",
							expr: &zeroOrOneExpr{
								pos: position{line: 541, col: 57, offset: 24228},
								expr: &ruleRefExpr{
									pos:  position{line: 541, col: 57, offset: 24228},
									name: "URL_TEXT",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 541, col: 68, offset: 24239},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "InlineImage",
			pos:  position{line: 545, col: 1, offset: 24306},
			expr: &actionExpr{
				pos: position{line: 545, col: 16, offset: 24321},
				run: (*parser).callonInlineImage1,
				expr: &labeledExpr{
					pos:   position{line: 545, col: 16, offset: 24321},
					label: "image",
					expr: &ruleRefExpr{
						pos:  position{line: 545, col: 22, offset: 24327},
						name: "InlineImageMacro",
					},
				},
			},
		},
		{
			name: "InlineImageMacro",
			pos:  position{line: 550, col: 1, offset: 24472},
			expr: &actionExpr{
				pos: position{line: 550, col: 21, offset: 24492},
				run: (*parser).callonInlineImageMacro1,
				expr: &seqExpr{
					pos: position{line: 550, col: 21, offset: 24492},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 550, col: 21, offset: 24492},
							val:        "image:",
							ignoreCase: false,
						},
						&notExpr{
							pos: position{line: 550, col: 30, offset: 24501},
							expr: &litMatcher{
								pos:        position{line: 550, col: 31, offset: 24502},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 550, col: 35, offset: 24506},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 550, col: 41, offset: 24512},
								name: "URL",
							},
						},
						&litMatcher{
							pos:        position{line: 550, col: 46, offset: 24517},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 550, col: 50, offset: 24521},
							label: "attributes",
							expr: &zeroOrOneExpr{
								pos: position{line: 550, col: 62, offset: 24533},
								expr: &ruleRefExpr{
									pos:  position{line: 550, col: 62, offset: 24533},
									name: "URL_TEXT",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 550, col: 73, offset: 24544},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "DelimitedBlock",
			pos:  position{line: 557, col: 1, offset: 24874},
			expr: &choiceExpr{
				pos: position{line: 557, col: 19, offset: 24892},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 557, col: 19, offset: 24892},
						name: "FencedBlock",
					},
					&ruleRefExpr{
						pos:  position{line: 557, col: 33, offset: 24906},
						name: "ListingBlock",
					},
					&ruleRefExpr{
						pos:  position{line: 557, col: 48, offset: 24921},
						name: "ExampleBlock",
					},
				},
			},
		},
		{
			name: "BlockDelimiter",
			pos:  position{line: 559, col: 1, offset: 24935},
			expr: &choiceExpr{
				pos: position{line: 559, col: 19, offset: 24953},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 559, col: 19, offset: 24953},
						name: "LiteralBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 559, col: 43, offset: 24977},
						name: "FencedBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 559, col: 66, offset: 25000},
						name: "ListingBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 559, col: 90, offset: 25024},
						name: "ExampleBlockDelimiter",
					},
				},
			},
		},
		{
			name: "FencedBlockDelimiter",
			pos:  position{line: 561, col: 1, offset: 25047},
			expr: &litMatcher{
				pos:        position{line: 561, col: 25, offset: 25071},
				val:        "```",
				ignoreCase: false,
			},
		},
		{
			name: "FencedBlock",
			pos:  position{line: 563, col: 1, offset: 25078},
			expr: &actionExpr{
				pos: position{line: 563, col: 16, offset: 25093},
				run: (*parser).callonFencedBlock1,
				expr: &seqExpr{
					pos: position{line: 563, col: 16, offset: 25093},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 563, col: 16, offset: 25093},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 563, col: 27, offset: 25104},
								expr: &ruleRefExpr{
									pos:  position{line: 563, col: 28, offset: 25105},
									name: "ElementAttribute",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 563, col: 47, offset: 25124},
							name: "FencedBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 563, col: 68, offset: 25145},
							expr: &ruleRefExpr{
								pos:  position{line: 563, col: 68, offset: 25145},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 563, col: 72, offset: 25149},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 563, col: 80, offset: 25157},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 563, col: 88, offset: 25165},
								expr: &choiceExpr{
									pos: position{line: 563, col: 89, offset: 25166},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 563, col: 89, offset: 25166},
											name: "List",
										},
										&ruleRefExpr{
											pos:  position{line: 563, col: 96, offset: 25173},
											name: "Paragraph",
										},
										&ruleRefExpr{
											pos:  position{line: 563, col: 108, offset: 25185},
											name: "BlankLine",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 563, col: 120, offset: 25197},
							name: "FencedBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 563, col: 141, offset: 25218},
							expr: &ruleRefExpr{
								pos:  position{line: 563, col: 141, offset: 25218},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 563, col: 145, offset: 25222},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ListingBlockDelimiter",
			pos:  position{line: 567, col: 1, offset: 25338},
			expr: &litMatcher{
				pos:        position{line: 567, col: 26, offset: 25363},
				val:        "----",
				ignoreCase: false,
			},
		},
		{
			name: "ListingBlock",
			pos:  position{line: 569, col: 1, offset: 25371},
			expr: &actionExpr{
				pos: position{line: 569, col: 17, offset: 25387},
				run: (*parser).callonListingBlock1,
				expr: &seqExpr{
					pos: position{line: 569, col: 17, offset: 25387},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 569, col: 17, offset: 25387},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 569, col: 28, offset: 25398},
								expr: &ruleRefExpr{
									pos:  position{line: 569, col: 29, offset: 25399},
									name: "ElementAttribute",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 569, col: 48, offset: 25418},
							name: "ListingBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 569, col: 70, offset: 25440},
							expr: &ruleRefExpr{
								pos:  position{line: 569, col: 70, offset: 25440},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 569, col: 74, offset: 25444},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 569, col: 82, offset: 25452},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 569, col: 90, offset: 25460},
								expr: &choiceExpr{
									pos: position{line: 569, col: 91, offset: 25461},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 569, col: 91, offset: 25461},
											name: "List",
										},
										&ruleRefExpr{
											pos:  position{line: 569, col: 98, offset: 25468},
											name: "Paragraph",
										},
										&ruleRefExpr{
											pos:  position{line: 569, col: 110, offset: 25480},
											name: "BlankLine",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 569, col: 122, offset: 25492},
							name: "ListingBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 569, col: 144, offset: 25514},
							expr: &ruleRefExpr{
								pos:  position{line: 569, col: 144, offset: 25514},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 569, col: 148, offset: 25518},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ExampleBlockDelimiter",
			pos:  position{line: 573, col: 1, offset: 25635},
			expr: &litMatcher{
				pos:        position{line: 573, col: 26, offset: 25660},
				val:        "====",
				ignoreCase: false,
			},
		},
		{
			name: "ExampleBlock",
			pos:  position{line: 575, col: 1, offset: 25668},
			expr: &actionExpr{
				pos: position{line: 575, col: 17, offset: 25684},
				run: (*parser).callonExampleBlock1,
				expr: &seqExpr{
					pos: position{line: 575, col: 17, offset: 25684},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 575, col: 17, offset: 25684},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 575, col: 28, offset: 25695},
								expr: &ruleRefExpr{
									pos:  position{line: 575, col: 29, offset: 25696},
									name: "ElementAttribute",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 575, col: 48, offset: 25715},
							name: "ExampleBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 575, col: 70, offset: 25737},
							expr: &ruleRefExpr{
								pos:  position{line: 575, col: 70, offset: 25737},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 575, col: 74, offset: 25741},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 575, col: 82, offset: 25749},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 575, col: 90, offset: 25757},
								expr: &choiceExpr{
									pos: position{line: 575, col: 91, offset: 25758},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 575, col: 91, offset: 25758},
											name: "List",
										},
										&ruleRefExpr{
											pos:  position{line: 575, col: 98, offset: 25765},
											name: "Paragraph",
										},
										&ruleRefExpr{
											pos:  position{line: 575, col: 110, offset: 25777},
											name: "BlankLine",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 575, col: 123, offset: 25790},
							name: "ExampleBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 575, col: 145, offset: 25812},
							expr: &ruleRefExpr{
								pos:  position{line: 575, col: 145, offset: 25812},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 575, col: 149, offset: 25816},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "LiteralBlock",
			pos:  position{line: 582, col: 1, offset: 26200},
			expr: &choiceExpr{
				pos: position{line: 582, col: 17, offset: 26216},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 582, col: 17, offset: 26216},
						name: "ParagraphWithSpaces",
					},
					&ruleRefExpr{
						pos:  position{line: 582, col: 39, offset: 26238},
						name: "ParagraphWithLiteralBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 582, col: 76, offset: 26275},
						name: "ParagraphWithLiteralAttribute",
					},
				},
			},
		},
		{
			name: "ParagraphWithSpaces",
			pos:  position{line: 585, col: 1, offset: 26370},
			expr: &actionExpr{
				pos: position{line: 585, col: 24, offset: 26393},
				run: (*parser).callonParagraphWithSpaces1,
				expr: &seqExpr{
					pos: position{line: 585, col: 24, offset: 26393},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 585, col: 24, offset: 26393},
							label: "spaces",
							expr: &oneOrMoreExpr{
								pos: position{line: 585, col: 32, offset: 26401},
								expr: &ruleRefExpr{
									pos:  position{line: 585, col: 32, offset: 26401},
									name: "WS",
								},
							},
						},
						&notExpr{
							pos: position{line: 585, col: 37, offset: 26406},
							expr: &ruleRefExpr{
								pos:  position{line: 585, col: 38, offset: 26407},
								name: "NEWLINE",
							},
						},
						&labeledExpr{
							pos:   position{line: 585, col: 46, offset: 26415},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 585, col: 55, offset: 26424},
								name: "LiteralBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 585, col: 76, offset: 26445},
							name: "EndOfLiteralBlock",
						},
					},
				},
			},
		},
		{
			name: "LiteralBlockContent",
			pos:  position{line: 590, col: 1, offset: 26626},
			expr: &actionExpr{
				pos: position{line: 590, col: 24, offset: 26649},
				run: (*parser).callonLiteralBlockContent1,
				expr: &labeledExpr{
					pos:   position{line: 590, col: 24, offset: 26649},
					label: "content",
					expr: &oneOrMoreExpr{
						pos: position{line: 590, col: 32, offset: 26657},
						expr: &seqExpr{
							pos: position{line: 590, col: 33, offset: 26658},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 590, col: 33, offset: 26658},
									expr: &seqExpr{
										pos: position{line: 590, col: 35, offset: 26660},
										exprs: []interface{}{
											&ruleRefExpr{
												pos:  position{line: 590, col: 35, offset: 26660},
												name: "NEWLINE",
											},
											&ruleRefExpr{
												pos:  position{line: 590, col: 43, offset: 26668},
												name: "BlankLine",
											},
										},
									},
								},
								&anyMatcher{
									line: 590, col: 54, offset: 26679,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "EndOfLiteralBlock",
			pos:  position{line: 595, col: 1, offset: 26764},
			expr: &choiceExpr{
				pos: position{line: 595, col: 22, offset: 26785},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 595, col: 22, offset: 26785},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 595, col: 22, offset: 26785},
								name: "NEWLINE",
							},
							&ruleRefExpr{
								pos:  position{line: 595, col: 30, offset: 26793},
								name: "BlankLine",
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 595, col: 42, offset: 26805},
						name: "NEWLINE",
					},
					&ruleRefExpr{
						pos:  position{line: 595, col: 52, offset: 26815},
						name: "EOF",
					},
				},
			},
		},
		{
			name: "ParagraphWithLiteralBlockDelimiter",
			pos:  position{line: 598, col: 1, offset: 26875},
			expr: &actionExpr{
				pos: position{line: 598, col: 39, offset: 26913},
				run: (*parser).callonParagraphWithLiteralBlockDelimiter1,
				expr: &seqExpr{
					pos: position{line: 598, col: 39, offset: 26913},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 598, col: 39, offset: 26913},
							name: "LiteralBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 598, col: 61, offset: 26935},
							expr: &ruleRefExpr{
								pos:  position{line: 598, col: 61, offset: 26935},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 598, col: 65, offset: 26939},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 598, col: 73, offset: 26947},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 598, col: 81, offset: 26955},
								expr: &seqExpr{
									pos: position{line: 598, col: 82, offset: 26956},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 598, col: 82, offset: 26956},
											expr: &ruleRefExpr{
												pos:  position{line: 598, col: 83, offset: 26957},
												name: "LiteralBlockDelimiter",
											},
										},
										&anyMatcher{
											line: 598, col: 105, offset: 26979,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 598, col: 109, offset: 26983},
							name: "LiteralBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 598, col: 131, offset: 27005},
							expr: &ruleRefExpr{
								pos:  position{line: 598, col: 131, offset: 27005},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 598, col: 135, offset: 27009},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "LiteralBlockDelimiter",
			pos:  position{line: 602, col: 1, offset: 27093},
			expr: &litMatcher{
				pos:        position{line: 602, col: 26, offset: 27118},
				val:        "....",
				ignoreCase: false,
			},
		},
		{
			name: "ParagraphWithLiteralAttribute",
			pos:  position{line: 605, col: 1, offset: 27180},
			expr: &actionExpr{
				pos: position{line: 605, col: 34, offset: 27213},
				run: (*parser).callonParagraphWithLiteralAttribute1,
				expr: &seqExpr{
					pos: position{line: 605, col: 34, offset: 27213},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 605, col: 34, offset: 27213},
							val:        "[literal]",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 605, col: 46, offset: 27225},
							expr: &ruleRefExpr{
								pos:  position{line: 605, col: 46, offset: 27225},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 605, col: 50, offset: 27229},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 605, col: 58, offset: 27237},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 605, col: 67, offset: 27246},
								name: "LiteralBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 605, col: 88, offset: 27267},
							name: "EndOfLiteralBlock",
						},
					},
				},
			},
		},
		{
			name: "BlankLine",
			pos:  position{line: 612, col: 1, offset: 27470},
			expr: &actionExpr{
				pos: position{line: 612, col: 14, offset: 27483},
				run: (*parser).callonBlankLine1,
				expr: &seqExpr{
					pos: position{line: 612, col: 14, offset: 27483},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 612, col: 14, offset: 27483},
							expr: &ruleRefExpr{
								pos:  position{line: 612, col: 15, offset: 27484},
								name: "EOF",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 612, col: 19, offset: 27488},
							expr: &ruleRefExpr{
								pos:  position{line: 612, col: 19, offset: 27488},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 612, col: 23, offset: 27492},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "Characters",
			pos:  position{line: 619, col: 1, offset: 27639},
			expr: &actionExpr{
				pos: position{line: 619, col: 15, offset: 27653},
				run: (*parser).callonCharacters1,
				expr: &oneOrMoreExpr{
					pos: position{line: 619, col: 15, offset: 27653},
					expr: &seqExpr{
						pos: position{line: 619, col: 16, offset: 27654},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 619, col: 16, offset: 27654},
								expr: &ruleRefExpr{
									pos:  position{line: 619, col: 17, offset: 27655},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 619, col: 25, offset: 27663},
								expr: &ruleRefExpr{
									pos:  position{line: 619, col: 26, offset: 27664},
									name: "WS",
								},
							},
							&anyMatcher{
								line: 619, col: 29, offset: 27667,
							},
						},
					},
				},
			},
		},
		{
			name: "URL",
			pos:  position{line: 623, col: 1, offset: 27707},
			expr: &actionExpr{
				pos: position{line: 623, col: 8, offset: 27714},
				run: (*parser).callonURL1,
				expr: &oneOrMoreExpr{
					pos: position{line: 623, col: 8, offset: 27714},
					expr: &seqExpr{
						pos: position{line: 623, col: 9, offset: 27715},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 623, col: 9, offset: 27715},
								expr: &ruleRefExpr{
									pos:  position{line: 623, col: 10, offset: 27716},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 623, col: 18, offset: 27724},
								expr: &ruleRefExpr{
									pos:  position{line: 623, col: 19, offset: 27725},
									name: "WS",
								},
							},
							&notExpr{
								pos: position{line: 623, col: 22, offset: 27728},
								expr: &litMatcher{
									pos:        position{line: 623, col: 23, offset: 27729},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 623, col: 27, offset: 27733},
								expr: &litMatcher{
									pos:        position{line: 623, col: 28, offset: 27734},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 623, col: 32, offset: 27738,
							},
						},
					},
				},
			},
		},
		{
			name: "ID",
			pos:  position{line: 627, col: 1, offset: 27778},
			expr: &actionExpr{
				pos: position{line: 627, col: 7, offset: 27784},
				run: (*parser).callonID1,
				expr: &oneOrMoreExpr{
					pos: position{line: 627, col: 7, offset: 27784},
					expr: &seqExpr{
						pos: position{line: 627, col: 8, offset: 27785},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 627, col: 8, offset: 27785},
								expr: &ruleRefExpr{
									pos:  position{line: 627, col: 9, offset: 27786},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 627, col: 17, offset: 27794},
								expr: &ruleRefExpr{
									pos:  position{line: 627, col: 18, offset: 27795},
									name: "WS",
								},
							},
							&notExpr{
								pos: position{line: 627, col: 21, offset: 27798},
								expr: &litMatcher{
									pos:        position{line: 627, col: 22, offset: 27799},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 627, col: 26, offset: 27803},
								expr: &litMatcher{
									pos:        position{line: 627, col: 27, offset: 27804},
									val:        "]",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 627, col: 31, offset: 27808},
								expr: &litMatcher{
									pos:        position{line: 627, col: 32, offset: 27809},
									val:        "<<",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 627, col: 37, offset: 27814},
								expr: &litMatcher{
									pos:        position{line: 627, col: 38, offset: 27815},
									val:        ">>",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 627, col: 42, offset: 27819,
							},
						},
					},
				},
			},
		},
		{
			name: "URL_TEXT",
			pos:  position{line: 631, col: 1, offset: 27859},
			expr: &actionExpr{
				pos: position{line: 631, col: 13, offset: 27871},
				run: (*parser).callonURL_TEXT1,
				expr: &oneOrMoreExpr{
					pos: position{line: 631, col: 13, offset: 27871},
					expr: &seqExpr{
						pos: position{line: 631, col: 14, offset: 27872},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 631, col: 14, offset: 27872},
								expr: &ruleRefExpr{
									pos:  position{line: 631, col: 15, offset: 27873},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 631, col: 23, offset: 27881},
								expr: &litMatcher{
									pos:        position{line: 631, col: 24, offset: 27882},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 631, col: 28, offset: 27886},
								expr: &litMatcher{
									pos:        position{line: 631, col: 29, offset: 27887},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 631, col: 33, offset: 27891,
							},
						},
					},
				},
			},
		},
		{
			name: "URL_SCHEME",
			pos:  position{line: 635, col: 1, offset: 27931},
			expr: &choiceExpr{
				pos: position{line: 635, col: 15, offset: 27945},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 635, col: 15, offset: 27945},
						val:        "http://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 635, col: 27, offset: 27957},
						val:        "https://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 635, col: 40, offset: 27970},
						val:        "ftp://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 635, col: 51, offset: 27981},
						val:        "irc://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 635, col: 62, offset: 27992},
						val:        "mailto:",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "DIGIT",
			pos:  position{line: 637, col: 1, offset: 28003},
			expr: &charClassMatcher{
				pos:        position{line: 637, col: 10, offset: 28012},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NEWLINE",
			pos:  position{line: 639, col: 1, offset: 28019},
			expr: &choiceExpr{
				pos: position{line: 639, col: 12, offset: 28030},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 639, col: 12, offset: 28030},
						val:        "\r\n",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 639, col: 21, offset: 28039},
						val:        "\r",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 639, col: 28, offset: 28046},
						val:        "\n",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "WS",
			pos:  position{line: 641, col: 1, offset: 28052},
			expr: &choiceExpr{
				pos: position{line: 641, col: 7, offset: 28058},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 641, col: 7, offset: 28058},
						val:        " ",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 641, col: 13, offset: 28064},
						run: (*parser).callonWS3,
						expr: &litMatcher{
							pos:        position{line: 641, col: 13, offset: 28064},
							val:        "\t",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 645, col: 1, offset: 28109},
			expr: &notExpr{
				pos: position{line: 645, col: 8, offset: 28116},
				expr: &anyMatcher{
					line: 645, col: 9, offset: 28117,
				},
			},
		},
		{
			name: "EOL",
			pos:  position{line: 647, col: 1, offset: 28120},
			expr: &choiceExpr{
				pos: position{line: 647, col: 8, offset: 28127},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 647, col: 8, offset: 28127},
						name: "NEWLINE",
					},
					&ruleRefExpr{
						pos:  position{line: 647, col: 18, offset: 28137},
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

func (c *current) onDocumentTitle1(attributes, level, content, id interface{}) (interface{}, error) {

	return types.NewSectionTitle(content.(types.InlineElements), append(attributes.([]interface{}), id))
}

func (p *parser) callonDocumentTitle1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocumentTitle1(stack["attributes"], stack["level"], stack["content"], stack["id"])
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
	return []types.DocumentAuthor{author.(types.DocumentAuthor)}, nil
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

func (c *current) onElementAttribute1(attr interface{}) (interface{}, error) {
	return attr, nil // avoid returning something like `[]interface{}{attr, EOL}`
}

func (p *parser) callonElementAttribute1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onElementAttribute1(stack["attr"])
}

func (c *current) onElementID2(id interface{}) (interface{}, error) {
	return id, nil
}

func (p *parser) callonElementID2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onElementID2(stack["id"])
}

func (c *current) onElementID5(id interface{}) (interface{}, error) {
	return types.NewElementID(id.(string))
}

func (p *parser) callonElementID5() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onElementID5(stack["id"])
}

func (c *current) onInlineElementID1(id interface{}) (interface{}, error) {
	return types.NewElementID(id.(string))
}

func (p *parser) callonInlineElementID1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInlineElementID1(stack["id"])
}

func (c *current) onElementTitle1(title interface{}) (interface{}, error) {
	return types.NewElementTitle(title.([]interface{}))
}

func (p *parser) callonElementTitle1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onElementTitle1(stack["title"])
}

func (c *current) onAdmonitionMarkerAttribute1(k interface{}) (interface{}, error) {
	return k, nil
}

func (p *parser) callonAdmonitionMarkerAttribute1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAdmonitionMarkerAttribute1(stack["k"])
}

func (c *current) onAttributeGroup1(attributes interface{}) (interface{}, error) {
	return types.NewAttributeGroup(attributes.([]interface{}))
}

func (p *parser) callonAttributeGroup1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAttributeGroup1(stack["attributes"])
}

func (c *current) onGenericAttribute2(key, value interface{}) (interface{}, error) {
	// value is set
	return types.NewGenericAttribute(key.([]interface{}), value.([]interface{}))
}

func (p *parser) callonGenericAttribute2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onGenericAttribute2(stack["key"], stack["value"])
}

func (c *current) onGenericAttribute14(key interface{}) (interface{}, error) {
	// value is not set
	return types.NewGenericAttribute(key.([]interface{}), nil)
}

func (p *parser) callonGenericAttribute14() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onGenericAttribute14(stack["key"])
}

func (c *current) onAttributeKey1(key interface{}) (interface{}, error) {
	// fmt.Printf("found attribute key: %v\n", key)
	return key, nil
}

func (p *parser) callonAttributeKey1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAttributeKey1(stack["key"])
}

func (c *current) onAttributeValue1(value interface{}) (interface{}, error) {
	// fmt.Printf("found attribute value: %v\n", value)
	return value, nil
}

func (p *parser) callonAttributeValue1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAttributeValue1(stack["value"])
}

func (c *current) onInvalidElementAttribute1(content interface{}) (interface{}, error) {
	return types.NewInvalidElementAttribute(c.text)
}

func (p *parser) callonInvalidElementAttribute1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInvalidElementAttribute1(stack["content"])
}

func (c *current) onSection11(header, elements interface{}) (interface{}, error) {
	return types.NewSection(1, header.(types.SectionTitle), elements.([]interface{}))
}

func (p *parser) callonSection11() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection11(stack["header"], stack["elements"])
}

func (c *current) onSection1Block1(content interface{}) (interface{}, error) {
	return content, nil
}

func (p *parser) callonSection1Block1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection1Block1(stack["content"])
}

func (c *current) onSection21(header, elements interface{}) (interface{}, error) {
	return types.NewSection(2, header.(types.SectionTitle), elements.([]interface{}))
}

func (p *parser) callonSection21() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection21(stack["header"], stack["elements"])
}

func (c *current) onSection2Block1(content interface{}) (interface{}, error) {
	return content, nil
}

func (p *parser) callonSection2Block1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection2Block1(stack["content"])
}

func (c *current) onSection31(header, elements interface{}) (interface{}, error) {
	return types.NewSection(3, header.(types.SectionTitle), elements.([]interface{}))
}

func (p *parser) callonSection31() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection31(stack["header"], stack["elements"])
}

func (c *current) onSection3Block1(content interface{}) (interface{}, error) {
	return content, nil
}

func (p *parser) callonSection3Block1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection3Block1(stack["content"])
}

func (c *current) onSection41(header, elements interface{}) (interface{}, error) {
	return types.NewSection(4, header.(types.SectionTitle), elements.([]interface{}))
}

func (p *parser) callonSection41() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection41(stack["header"], stack["elements"])
}

func (c *current) onSection4Block1(content interface{}) (interface{}, error) {
	return content, nil
}

func (p *parser) callonSection4Block1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection4Block1(stack["content"])
}

func (c *current) onSection51(header, elements interface{}) (interface{}, error) {
	return types.NewSection(5, header.(types.SectionTitle), elements.([]interface{}))
}

func (p *parser) callonSection51() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection51(stack["header"], stack["elements"])
}

func (c *current) onSection5Block1(content interface{}) (interface{}, error) {
	return content, nil
}

func (p *parser) callonSection5Block1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection5Block1(stack["content"])
}

func (c *current) onSection1Title1(attributes, level, content, id interface{}) (interface{}, error) {

	return types.NewSectionTitle(content.(types.InlineElements), append(attributes.([]interface{}), id))
}

func (p *parser) callonSection1Title1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection1Title1(stack["attributes"], stack["level"], stack["content"], stack["id"])
}

func (c *current) onSection2Title1(attributes, level, content, id interface{}) (interface{}, error) {
	return types.NewSectionTitle(content.(types.InlineElements), append(attributes.([]interface{}), id))
}

func (p *parser) callonSection2Title1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection2Title1(stack["attributes"], stack["level"], stack["content"], stack["id"])
}

func (c *current) onSection3Title1(attributes, level, content, id interface{}) (interface{}, error) {
	return types.NewSectionTitle(content.(types.InlineElements), append(attributes.([]interface{}), id))
}

func (p *parser) callonSection3Title1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection3Title1(stack["attributes"], stack["level"], stack["content"], stack["id"])
}

func (c *current) onSection4Title1(attributes, level, content, id interface{}) (interface{}, error) {
	return types.NewSectionTitle(content.(types.InlineElements), append(attributes.([]interface{}), id))
}

func (p *parser) callonSection4Title1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection4Title1(stack["attributes"], stack["level"], stack["content"], stack["id"])
}

func (c *current) onSection5Title1(attributes, level, content, id interface{}) (interface{}, error) {
	return types.NewSectionTitle(content.(types.InlineElements), append(attributes.([]interface{}), id))
}

func (p *parser) callonSection5Title1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection5Title1(stack["attributes"], stack["level"], stack["content"], stack["id"])
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

func (c *current) onOrderedListItem1(attributes, prefix, content interface{}) (interface{}, error) {
	return types.NewOrderedListItem(prefix.(types.OrderedListItemPrefix), content.([]interface{}), attributes.([]interface{}))
}

func (p *parser) callonOrderedListItem1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOrderedListItem1(stack["attributes"], stack["prefix"], stack["content"])
}

func (c *current) onOrderedListItemPrefix2(style interface{}) (interface{}, error) {
	// numbering style: "."
	return types.NewOrderedListItemPrefix(types.Arabic, 1)
}

func (p *parser) callonOrderedListItemPrefix2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOrderedListItemPrefix2(stack["style"])
}

func (c *current) onOrderedListItemPrefix10(style interface{}) (interface{}, error) {
	// numbering style: ".."
	return types.NewOrderedListItemPrefix(types.LowerAlpha, 2)
}

func (p *parser) callonOrderedListItemPrefix10() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOrderedListItemPrefix10(stack["style"])
}

func (c *current) onOrderedListItemPrefix18(style interface{}) (interface{}, error) {
	// numbering style: "..."
	return types.NewOrderedListItemPrefix(types.LowerRoman, 3)
}

func (p *parser) callonOrderedListItemPrefix18() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOrderedListItemPrefix18(stack["style"])
}

func (c *current) onOrderedListItemPrefix26(style interface{}) (interface{}, error) {
	// numbering style: "...."
	return types.NewOrderedListItemPrefix(types.UpperAlpha, 4)
}

func (p *parser) callonOrderedListItemPrefix26() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOrderedListItemPrefix26(stack["style"])
}

func (c *current) onOrderedListItemPrefix34(style interface{}) (interface{}, error) {
	// numbering style: "....."
	return types.NewOrderedListItemPrefix(types.UpperRoman, 5)
	// explicit numbering
}

func (p *parser) callonOrderedListItemPrefix34() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOrderedListItemPrefix34(stack["style"])
}

func (c *current) onOrderedListItemPrefix42(style interface{}) (interface{}, error) {
	// numbering style: "1."
	return types.NewOrderedListItemPrefix(types.Arabic, 1)
}

func (p *parser) callonOrderedListItemPrefix42() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOrderedListItemPrefix42(stack["style"])
}

func (c *current) onOrderedListItemPrefix60(style interface{}) (interface{}, error) {
	// numbering style: "a."
	return types.NewOrderedListItemPrefix(types.LowerAlpha, 1)
}

func (p *parser) callonOrderedListItemPrefix60() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOrderedListItemPrefix60(stack["style"])
}

func (c *current) onOrderedListItemPrefix78(style interface{}) (interface{}, error) {
	// numbering style: "A."
	return types.NewOrderedListItemPrefix(types.UpperAlpha, 1)
}

func (p *parser) callonOrderedListItemPrefix78() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOrderedListItemPrefix78(stack["style"])
}

func (c *current) onOrderedListItemPrefix96(style interface{}) (interface{}, error) {
	// numbering style: "i)"
	return types.NewOrderedListItemPrefix(types.LowerRoman, 1)
}

func (p *parser) callonOrderedListItemPrefix96() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOrderedListItemPrefix96(stack["style"])
}

func (c *current) onOrderedListItemPrefix114(style interface{}) (interface{}, error) {
	// numbering style: "I)"
	return types.NewOrderedListItemPrefix(types.UpperRoman, 1)
}

func (p *parser) callonOrderedListItemPrefix114() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOrderedListItemPrefix114(stack["style"])
}

func (c *current) onOrderedListItemContent1(elements interface{}) (interface{}, error) {
	// Another list or a literal paragraph immediately following a list item will be implicitly included in the list item
	return types.NewListItemContent(elements.([]interface{}))
}

func (p *parser) callonOrderedListItemContent1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOrderedListItemContent1(stack["elements"])
}

func (c *current) onUnorderedListItem1(prefix, content interface{}) (interface{}, error) {
	return types.NewUnorderedListItem(prefix.(types.UnorderedListItemPrefix), content.([]interface{}))
}

func (p *parser) callonUnorderedListItem1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUnorderedListItem1(stack["prefix"], stack["content"])
}

func (c *current) onUnorderedListItemPrefix2(level interface{}) (interface{}, error) {
	// ignore whitespaces, only return the relevant "*"/"-" characters
	return types.NewUnorderedListItemPrefix(types.FiveAsterisks, 5)
}

func (p *parser) callonUnorderedListItemPrefix2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUnorderedListItemPrefix2(stack["level"])
}

func (c *current) onUnorderedListItemPrefix10(level interface{}) (interface{}, error) {
	// ignore whitespaces, only return the relevant "*"/"-" characters
	return types.NewUnorderedListItemPrefix(types.FourAsterisks, 4)
}

func (p *parser) callonUnorderedListItemPrefix10() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUnorderedListItemPrefix10(stack["level"])
}

func (c *current) onUnorderedListItemPrefix18(level interface{}) (interface{}, error) {
	// ignore whitespaces, only return the relevant "*"/"-" characters
	return types.NewUnorderedListItemPrefix(types.ThreeAsterisks, 3)
}

func (p *parser) callonUnorderedListItemPrefix18() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUnorderedListItemPrefix18(stack["level"])
}

func (c *current) onUnorderedListItemPrefix26(level interface{}) (interface{}, error) {
	// ignore whitespaces, only return the relevant "*"/"-" characters
	return types.NewUnorderedListItemPrefix(types.TwoAsterisks, 2)
}

func (p *parser) callonUnorderedListItemPrefix26() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUnorderedListItemPrefix26(stack["level"])
}

func (c *current) onUnorderedListItemPrefix34(level interface{}) (interface{}, error) {
	// ignore whitespaces, only return the relevant "*"/"-" characters
	return types.NewUnorderedListItemPrefix(types.OneAsterisk, 1)
}

func (p *parser) callonUnorderedListItemPrefix34() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUnorderedListItemPrefix34(stack["level"])
}

func (c *current) onUnorderedListItemPrefix42(level interface{}) (interface{}, error) {
	// ignore whitespaces, only return the relevant "*"/"-" characters
	return types.NewUnorderedListItemPrefix(types.Dash, 1)
}

func (p *parser) callonUnorderedListItemPrefix42() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUnorderedListItemPrefix42(stack["level"])
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

func (c *current) onLabeledListItem2(term, description interface{}) (interface{}, error) {
	return types.NewLabeledListItem(term.([]interface{}), description.([]interface{}))
}

func (p *parser) callonLabeledListItem2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLabeledListItem2(stack["term"], stack["description"])
}

func (c *current) onLabeledListItem9(term interface{}) (interface{}, error) {
	// here, WS is optional since there is no description afterwards
	return types.NewLabeledListItem(term.([]interface{}), nil)
}

func (p *parser) callonLabeledListItem9() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLabeledListItem9(stack["term"])
}

func (c *current) onLabeledListItemTerm1(term interface{}) (interface{}, error) {
	return term, nil
}

func (p *parser) callonLabeledListItemTerm1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLabeledListItemTerm1(stack["term"])
}

func (c *current) onLabeledListItemDescription1(elements interface{}) (interface{}, error) {
	// TODO: replace with (ListParagraph+ ContinuedBlockElement*) and use a single rule for all item contents ?
	return types.NewListItemContent(elements.([]interface{}))
}

func (p *parser) callonLabeledListItemDescription1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLabeledListItemDescription1(stack["elements"])
}

func (c *current) onAdmonitionKind2() (interface{}, error) {
	return types.Tip, nil
}

func (p *parser) callonAdmonitionKind2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAdmonitionKind2()
}

func (c *current) onAdmonitionKind4() (interface{}, error) {
	return types.Note, nil
}

func (p *parser) callonAdmonitionKind4() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAdmonitionKind4()
}

func (c *current) onAdmonitionKind6() (interface{}, error) {
	return types.Important, nil
}

func (p *parser) callonAdmonitionKind6() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAdmonitionKind6()
}

func (c *current) onAdmonitionKind8() (interface{}, error) {
	return types.Warning, nil
}

func (p *parser) callonAdmonitionKind8() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAdmonitionKind8()
}

func (c *current) onAdmonitionKind10() (interface{}, error) {
	return types.Caution, nil
}

func (p *parser) callonAdmonitionKind10() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAdmonitionKind10()
}

func (c *current) onParagraph2(attributes, t, lines interface{}) (interface{}, error) {
	// admonition paragraph
	return types.NewParagraph(lines.([]interface{}), append(attributes.([]interface{}), t.(types.AdmonitionKind)))
}

func (p *parser) callonParagraph2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onParagraph2(stack["attributes"], stack["t"], stack["lines"])
}

func (c *current) onParagraph21(attributes, lines interface{}) (interface{}, error) {
	// regular paragraph
	return types.NewParagraph(lines.([]interface{}), attributes.([]interface{}))
}

func (p *parser) callonParagraph21() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onParagraph21(stack["attributes"], stack["lines"])
}

func (c *current) onInlineElementsWithTrailingSpaces1(elements interface{}) (interface{}, error) {
	// includes heading and trailing spaces in the elements arg
	return types.NewInlineElements(elements.([]interface{}))
}

func (p *parser) callonInlineElementsWithTrailingSpaces1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInlineElementsWithTrailingSpaces1(stack["elements"])
}

func (c *current) onInlineElements1(elements interface{}) (interface{}, error) {
	// absorbs heading and trailing spaces
	return types.NewInlineElements(elements.([]interface{}))
}

func (p *parser) callonInlineElements1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInlineElements1(stack["elements"])
}

func (c *current) onBoldText2(content interface{}) (interface{}, error) {
	// double punctuation must be evaluated first
	return types.NewQuotedText(types.Bold, content.([]interface{}))
}

func (p *parser) callonBoldText2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBoldText2(stack["content"])
}

func (c *current) onBoldText10(content interface{}) (interface{}, error) {
	// unbalanced `**` vs `*` punctuation
	result := append([]interface{}{"*"}, content.([]interface{}))
	return types.NewQuotedText(types.Bold, result)
}

func (p *parser) callonBoldText10() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBoldText10(stack["content"])
}

func (c *current) onBoldText18(content interface{}) (interface{}, error) {
	// single punctuation
	return types.NewQuotedText(types.Bold, content.([]interface{}))
}

func (p *parser) callonBoldText18() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBoldText18(stack["content"])
}

func (c *current) onEscapedBoldText2(backslashes, content interface{}) (interface{}, error) {
	// double punctuation must be evaluated first
	return types.NewEscapedQuotedText(backslashes.([]interface{}), "**", content.([]interface{}))
}

func (p *parser) callonEscapedBoldText2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEscapedBoldText2(stack["backslashes"], stack["content"])
}

func (c *current) onEscapedBoldText13(backslashes, content interface{}) (interface{}, error) {
	// unbalanced `**` vs `*` punctuation
	result := append([]interface{}{"*"}, content.([]interface{}))
	return types.NewEscapedQuotedText(backslashes.([]interface{}), "*", result)
}

func (p *parser) callonEscapedBoldText13() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEscapedBoldText13(stack["backslashes"], stack["content"])
}

func (c *current) onEscapedBoldText24(backslashes, content interface{}) (interface{}, error) {
	// simple punctuation must be evaluated last
	return types.NewEscapedQuotedText(backslashes.([]interface{}), "*", content.([]interface{}))
}

func (p *parser) callonEscapedBoldText24() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEscapedBoldText24(stack["backslashes"], stack["content"])
}

func (c *current) onItalicText2(content interface{}) (interface{}, error) {
	return types.NewQuotedText(types.Italic, content.([]interface{}))
}

func (p *parser) callonItalicText2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onItalicText2(stack["content"])
}

func (c *current) onItalicText10(content interface{}) (interface{}, error) {
	// unbalanced `__` vs `_` punctuation
	result := append([]interface{}{"_"}, content.([]interface{}))
	return types.NewQuotedText(types.Italic, result)
}

func (p *parser) callonItalicText10() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onItalicText10(stack["content"])
}

func (c *current) onItalicText18(content interface{}) (interface{}, error) {
	return types.NewQuotedText(types.Italic, content.([]interface{}))
}

func (p *parser) callonItalicText18() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onItalicText18(stack["content"])
}

func (c *current) onEscapedItalicText2(backslashes, content interface{}) (interface{}, error) {
	// double punctuation must be evaluated first
	return types.NewEscapedQuotedText(backslashes.([]interface{}), "__", content.([]interface{}))
}

func (p *parser) callonEscapedItalicText2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEscapedItalicText2(stack["backslashes"], stack["content"])
}

func (c *current) onEscapedItalicText13(backslashes, content interface{}) (interface{}, error) {
	// unbalanced `__` vs `_` punctuation
	result := append([]interface{}{"_"}, content.([]interface{}))
	return types.NewEscapedQuotedText(backslashes.([]interface{}), "_", result)
}

func (p *parser) callonEscapedItalicText13() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEscapedItalicText13(stack["backslashes"], stack["content"])
}

func (c *current) onEscapedItalicText24(backslashes, content interface{}) (interface{}, error) {
	// simple punctuation must be evaluated last
	return types.NewEscapedQuotedText(backslashes.([]interface{}), "_", content.([]interface{}))
}

func (p *parser) callonEscapedItalicText24() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEscapedItalicText24(stack["backslashes"], stack["content"])
}

func (c *current) onMonospaceText2(content interface{}) (interface{}, error) {
	// double punctuation must be evaluated first
	return types.NewQuotedText(types.Monospace, content.([]interface{}))
}

func (p *parser) callonMonospaceText2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMonospaceText2(stack["content"])
}

func (c *current) onMonospaceText10(content interface{}) (interface{}, error) {
	// unbalanced "``" vs "`" punctuation
	result := append([]interface{}{"`"}, content.([]interface{}))
	return types.NewQuotedText(types.Monospace, result)
}

func (p *parser) callonMonospaceText10() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMonospaceText10(stack["content"])
}

func (c *current) onMonospaceText18(content interface{}) (interface{}, error) {
	// simple punctuation must be evaluated last
	return types.NewQuotedText(types.Monospace, content.([]interface{}))
}

func (p *parser) callonMonospaceText18() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onMonospaceText18(stack["content"])
}

func (c *current) onEscapedMonospaceText2(backslashes, content interface{}) (interface{}, error) {
	// double punctuation must be evaluated first
	return types.NewEscapedQuotedText(backslashes.([]interface{}), "``", content.([]interface{}))
}

func (p *parser) callonEscapedMonospaceText2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEscapedMonospaceText2(stack["backslashes"], stack["content"])
}

func (c *current) onEscapedMonospaceText13(backslashes, content interface{}) (interface{}, error) {
	// unbalanced "``" vs "`" punctuation
	result := append([]interface{}{"`"}, content.([]interface{}))
	return types.NewEscapedQuotedText(backslashes.([]interface{}), "`", result)
}

func (p *parser) callonEscapedMonospaceText13() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEscapedMonospaceText13(stack["backslashes"], stack["content"])
}

func (c *current) onEscapedMonospaceText24(backslashes, content interface{}) (interface{}, error) {
	// simple punctuation must be evaluated last
	return types.NewEscapedQuotedText(backslashes.([]interface{}), "`", content.([]interface{}))
}

func (p *parser) callonEscapedMonospaceText24() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEscapedMonospaceText24(stack["backslashes"], stack["content"])
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

func (c *current) onPassthroughMacro2(content interface{}) (interface{}, error) {
	return types.NewPassthrough(types.PassthroughMacro, content.([]interface{}))
}

func (p *parser) callonPassthroughMacro2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onPassthroughMacro2(stack["content"])
}

func (c *current) onPassthroughMacro9(content interface{}) (interface{}, error) {
	return types.NewPassthrough(types.PassthroughMacro, content.([]interface{}))
}

func (p *parser) callonPassthroughMacro9() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onPassthroughMacro9(stack["content"])
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
		return types.NewLink(url.([]interface{}), text.([]interface{}))
	}
	return types.NewLink(url.([]interface{}), nil)
}

func (p *parser) callonExternalLink1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onExternalLink1(stack["url"], stack["text"])
}

func (c *current) onRelativeLink1(url, text interface{}) (interface{}, error) {
	if text != nil {
		return types.NewLink(url.([]interface{}), text.([]interface{}))
	}
	return types.NewLink(url.([]interface{}), nil)
}

func (p *parser) callonRelativeLink1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onRelativeLink1(stack["url"], stack["text"])
}

func (c *current) onBlockImage1(attributes, image interface{}) (interface{}, error) {
	// here we can ignore the blank line in the returned element
	return types.NewBlockImage(image.(types.ImageMacro), attributes.([]interface{}))
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
	return types.NewInlineImage(image.(types.ImageMacro))
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

func (c *current) onFencedBlock1(attributes, content interface{}) (interface{}, error) {
	return types.NewDelimitedBlock(types.FencedBlock, content.([]interface{}), attributes.([]interface{}))
}

func (p *parser) callonFencedBlock1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFencedBlock1(stack["attributes"], stack["content"])
}

func (c *current) onListingBlock1(attributes, content interface{}) (interface{}, error) {
	return types.NewDelimitedBlock(types.ListingBlock, content.([]interface{}), attributes.([]interface{}))
}

func (p *parser) callonListingBlock1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onListingBlock1(stack["attributes"], stack["content"])
}

func (c *current) onExampleBlock1(attributes, content interface{}) (interface{}, error) {
	return types.NewDelimitedBlock(types.ExampleBlock, content.([]interface{}), attributes.([]interface{}))
}

func (p *parser) callonExampleBlock1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onExampleBlock1(stack["attributes"], stack["content"])
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
