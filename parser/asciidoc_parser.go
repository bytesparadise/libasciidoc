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
								name: "InlineContent",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 56, col: 89, offset: 2054},
							expr: &ruleRefExpr{
								pos:  position{line: 56, col: 89, offset: 2054},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 56, col: 93, offset: 2058},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 56, col: 96, offset: 2061},
								expr: &ruleRefExpr{
									pos:  position{line: 56, col: 97, offset: 2062},
									name: "InlineElementID",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 56, col: 115, offset: 2080},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthors",
			pos:  position{line: 60, col: 1, offset: 2195},
			expr: &choiceExpr{
				pos: position{line: 60, col: 20, offset: 2214},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 60, col: 20, offset: 2214},
						name: "DocumentAuthorsInlineForm",
					},
					&ruleRefExpr{
						pos:  position{line: 60, col: 48, offset: 2242},
						name: "DocumentAuthorsAttributeForm",
					},
				},
			},
		},
		{
			name: "DocumentAuthorsInlineForm",
			pos:  position{line: 62, col: 1, offset: 2272},
			expr: &actionExpr{
				pos: position{line: 62, col: 30, offset: 2301},
				run: (*parser).callonDocumentAuthorsInlineForm1,
				expr: &seqExpr{
					pos: position{line: 62, col: 30, offset: 2301},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 62, col: 30, offset: 2301},
							expr: &ruleRefExpr{
								pos:  position{line: 62, col: 30, offset: 2301},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 62, col: 34, offset: 2305},
							expr: &litMatcher{
								pos:        position{line: 62, col: 35, offset: 2306},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 62, col: 39, offset: 2310},
							label: "authors",
							expr: &oneOrMoreExpr{
								pos: position{line: 62, col: 48, offset: 2319},
								expr: &ruleRefExpr{
									pos:  position{line: 62, col: 48, offset: 2319},
									name: "DocumentAuthor",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 62, col: 65, offset: 2336},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthorsAttributeForm",
			pos:  position{line: 66, col: 1, offset: 2406},
			expr: &actionExpr{
				pos: position{line: 66, col: 33, offset: 2438},
				run: (*parser).callonDocumentAuthorsAttributeForm1,
				expr: &seqExpr{
					pos: position{line: 66, col: 33, offset: 2438},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 66, col: 33, offset: 2438},
							expr: &ruleRefExpr{
								pos:  position{line: 66, col: 33, offset: 2438},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 66, col: 37, offset: 2442},
							val:        ":author:",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 66, col: 48, offset: 2453},
							label: "author",
							expr: &ruleRefExpr{
								pos:  position{line: 66, col: 56, offset: 2461},
								name: "DocumentAuthor",
							},
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthor",
			pos:  position{line: 70, col: 1, offset: 2552},
			expr: &actionExpr{
				pos: position{line: 70, col: 19, offset: 2570},
				run: (*parser).callonDocumentAuthor1,
				expr: &seqExpr{
					pos: position{line: 70, col: 19, offset: 2570},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 70, col: 19, offset: 2570},
							expr: &ruleRefExpr{
								pos:  position{line: 70, col: 19, offset: 2570},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 70, col: 23, offset: 2574},
							label: "namePart1",
							expr: &ruleRefExpr{
								pos:  position{line: 70, col: 34, offset: 2585},
								name: "DocumentAuthorNamePart",
							},
						},
						&labeledExpr{
							pos:   position{line: 70, col: 58, offset: 2609},
							label: "namePart2",
							expr: &zeroOrOneExpr{
								pos: position{line: 70, col: 68, offset: 2619},
								expr: &ruleRefExpr{
									pos:  position{line: 70, col: 69, offset: 2620},
									name: "DocumentAuthorNamePart",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 70, col: 94, offset: 2645},
							label: "namePart3",
							expr: &zeroOrOneExpr{
								pos: position{line: 70, col: 104, offset: 2655},
								expr: &ruleRefExpr{
									pos:  position{line: 70, col: 105, offset: 2656},
									name: "DocumentAuthorNamePart",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 70, col: 130, offset: 2681},
							label: "email",
							expr: &zeroOrOneExpr{
								pos: position{line: 70, col: 136, offset: 2687},
								expr: &ruleRefExpr{
									pos:  position{line: 70, col: 137, offset: 2688},
									name: "DocumentAuthorEmail",
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 70, col: 159, offset: 2710},
							expr: &ruleRefExpr{
								pos:  position{line: 70, col: 159, offset: 2710},
								name: "WS",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 70, col: 163, offset: 2714},
							expr: &litMatcher{
								pos:        position{line: 70, col: 163, offset: 2714},
								val:        ";",
								ignoreCase: false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 70, col: 168, offset: 2719},
							expr: &ruleRefExpr{
								pos:  position{line: 70, col: 168, offset: 2719},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthorNamePart",
			pos:  position{line: 75, col: 1, offset: 2884},
			expr: &seqExpr{
				pos: position{line: 75, col: 27, offset: 2910},
				exprs: []interface{}{
					&notExpr{
						pos: position{line: 75, col: 27, offset: 2910},
						expr: &litMatcher{
							pos:        position{line: 75, col: 28, offset: 2911},
							val:        "<",
							ignoreCase: false,
						},
					},
					&notExpr{
						pos: position{line: 75, col: 32, offset: 2915},
						expr: &litMatcher{
							pos:        position{line: 75, col: 33, offset: 2916},
							val:        ";",
							ignoreCase: false,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 37, offset: 2920},
						name: "Characters",
					},
					&zeroOrMoreExpr{
						pos: position{line: 75, col: 48, offset: 2931},
						expr: &ruleRefExpr{
							pos:  position{line: 75, col: 48, offset: 2931},
							name: "WS",
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthorEmail",
			pos:  position{line: 77, col: 1, offset: 2936},
			expr: &seqExpr{
				pos: position{line: 77, col: 24, offset: 2959},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 77, col: 24, offset: 2959},
						val:        "<",
						ignoreCase: false,
					},
					&labeledExpr{
						pos:   position{line: 77, col: 28, offset: 2963},
						label: "email",
						expr: &oneOrMoreExpr{
							pos: position{line: 77, col: 34, offset: 2969},
							expr: &seqExpr{
								pos: position{line: 77, col: 35, offset: 2970},
								exprs: []interface{}{
									&notExpr{
										pos: position{line: 77, col: 35, offset: 2970},
										expr: &litMatcher{
											pos:        position{line: 77, col: 36, offset: 2971},
											val:        ">",
											ignoreCase: false,
										},
									},
									&notExpr{
										pos: position{line: 77, col: 40, offset: 2975},
										expr: &ruleRefExpr{
											pos:  position{line: 77, col: 41, offset: 2976},
											name: "EOL",
										},
									},
									&anyMatcher{
										line: 77, col: 45, offset: 2980,
									},
								},
							},
						},
					},
					&litMatcher{
						pos:        position{line: 77, col: 49, offset: 2984},
						val:        ">",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "DocumentRevision",
			pos:  position{line: 81, col: 1, offset: 3120},
			expr: &actionExpr{
				pos: position{line: 81, col: 21, offset: 3140},
				run: (*parser).callonDocumentRevision1,
				expr: &seqExpr{
					pos: position{line: 81, col: 21, offset: 3140},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 81, col: 21, offset: 3140},
							expr: &ruleRefExpr{
								pos:  position{line: 81, col: 21, offset: 3140},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 81, col: 25, offset: 3144},
							expr: &litMatcher{
								pos:        position{line: 81, col: 26, offset: 3145},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 81, col: 30, offset: 3149},
							label: "revnumber",
							expr: &zeroOrOneExpr{
								pos: position{line: 81, col: 40, offset: 3159},
								expr: &ruleRefExpr{
									pos:  position{line: 81, col: 41, offset: 3160},
									name: "DocumentRevisionNumber",
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 81, col: 66, offset: 3185},
							expr: &litMatcher{
								pos:        position{line: 81, col: 66, offset: 3185},
								val:        ",",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 81, col: 71, offset: 3190},
							label: "revdate",
							expr: &zeroOrOneExpr{
								pos: position{line: 81, col: 79, offset: 3198},
								expr: &ruleRefExpr{
									pos:  position{line: 81, col: 80, offset: 3199},
									name: "DocumentRevisionDate",
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 81, col: 103, offset: 3222},
							expr: &litMatcher{
								pos:        position{line: 81, col: 103, offset: 3222},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 81, col: 108, offset: 3227},
							label: "revremark",
							expr: &zeroOrOneExpr{
								pos: position{line: 81, col: 118, offset: 3237},
								expr: &ruleRefExpr{
									pos:  position{line: 81, col: 119, offset: 3238},
									name: "DocumentRevisionRemark",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 81, col: 144, offset: 3263},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentRevisionNumber",
			pos:  position{line: 86, col: 1, offset: 3436},
			expr: &choiceExpr{
				pos: position{line: 86, col: 27, offset: 3462},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 86, col: 27, offset: 3462},
						exprs: []interface{}{
							&litMatcher{
								pos:        position{line: 86, col: 27, offset: 3462},
								val:        "v",
								ignoreCase: true,
							},
							&ruleRefExpr{
								pos:  position{line: 86, col: 32, offset: 3467},
								name: "DIGIT",
							},
							&zeroOrMoreExpr{
								pos: position{line: 86, col: 39, offset: 3474},
								expr: &seqExpr{
									pos: position{line: 86, col: 40, offset: 3475},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 86, col: 40, offset: 3475},
											expr: &ruleRefExpr{
												pos:  position{line: 86, col: 41, offset: 3476},
												name: "EOL",
											},
										},
										&notExpr{
											pos: position{line: 86, col: 45, offset: 3480},
											expr: &litMatcher{
												pos:        position{line: 86, col: 46, offset: 3481},
												val:        ",",
												ignoreCase: false,
											},
										},
										&notExpr{
											pos: position{line: 86, col: 50, offset: 3485},
											expr: &litMatcher{
												pos:        position{line: 86, col: 51, offset: 3486},
												val:        ":",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 86, col: 55, offset: 3490,
										},
									},
								},
							},
						},
					},
					&seqExpr{
						pos: position{line: 86, col: 61, offset: 3496},
						exprs: []interface{}{
							&zeroOrOneExpr{
								pos: position{line: 86, col: 61, offset: 3496},
								expr: &litMatcher{
									pos:        position{line: 86, col: 61, offset: 3496},
									val:        "v",
									ignoreCase: true,
								},
							},
							&ruleRefExpr{
								pos:  position{line: 86, col: 67, offset: 3502},
								name: "DIGIT",
							},
							&zeroOrMoreExpr{
								pos: position{line: 86, col: 74, offset: 3509},
								expr: &seqExpr{
									pos: position{line: 86, col: 75, offset: 3510},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 86, col: 75, offset: 3510},
											expr: &ruleRefExpr{
												pos:  position{line: 86, col: 76, offset: 3511},
												name: "EOL",
											},
										},
										&notExpr{
											pos: position{line: 86, col: 80, offset: 3515},
											expr: &litMatcher{
												pos:        position{line: 86, col: 81, offset: 3516},
												val:        ",",
												ignoreCase: false,
											},
										},
										&notExpr{
											pos: position{line: 86, col: 85, offset: 3520},
											expr: &litMatcher{
												pos:        position{line: 86, col: 86, offset: 3521},
												val:        ":",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 86, col: 90, offset: 3525,
										},
									},
								},
							},
							&zeroOrMoreExpr{
								pos: position{line: 86, col: 94, offset: 3529},
								expr: &ruleRefExpr{
									pos:  position{line: 86, col: 94, offset: 3529},
									name: "WS",
								},
							},
							&andExpr{
								pos: position{line: 86, col: 98, offset: 3533},
								expr: &litMatcher{
									pos:        position{line: 86, col: 99, offset: 3534},
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
			pos:  position{line: 87, col: 1, offset: 3538},
			expr: &zeroOrMoreExpr{
				pos: position{line: 87, col: 25, offset: 3562},
				expr: &seqExpr{
					pos: position{line: 87, col: 26, offset: 3563},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 87, col: 26, offset: 3563},
							expr: &ruleRefExpr{
								pos:  position{line: 87, col: 27, offset: 3564},
								name: "EOL",
							},
						},
						&notExpr{
							pos: position{line: 87, col: 31, offset: 3568},
							expr: &litMatcher{
								pos:        position{line: 87, col: 32, offset: 3569},
								val:        ":",
								ignoreCase: false,
							},
						},
						&anyMatcher{
							line: 87, col: 36, offset: 3573,
						},
					},
				},
			},
		},
		{
			name: "DocumentRevisionRemark",
			pos:  position{line: 88, col: 1, offset: 3578},
			expr: &zeroOrMoreExpr{
				pos: position{line: 88, col: 27, offset: 3604},
				expr: &seqExpr{
					pos: position{line: 88, col: 28, offset: 3605},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 88, col: 28, offset: 3605},
							expr: &ruleRefExpr{
								pos:  position{line: 88, col: 29, offset: 3606},
								name: "EOL",
							},
						},
						&anyMatcher{
							line: 88, col: 33, offset: 3610,
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclaration",
			pos:  position{line: 93, col: 1, offset: 3730},
			expr: &choiceExpr{
				pos: position{line: 93, col: 33, offset: 3762},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 93, col: 33, offset: 3762},
						name: "DocumentAttributeDeclarationWithNameOnly",
					},
					&ruleRefExpr{
						pos:  position{line: 93, col: 76, offset: 3805},
						name: "DocumentAttributeDeclarationWithNameAndValue",
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclarationWithNameOnly",
			pos:  position{line: 95, col: 1, offset: 3852},
			expr: &actionExpr{
				pos: position{line: 95, col: 45, offset: 3896},
				run: (*parser).callonDocumentAttributeDeclarationWithNameOnly1,
				expr: &seqExpr{
					pos: position{line: 95, col: 45, offset: 3896},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 95, col: 45, offset: 3896},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 95, col: 49, offset: 3900},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 95, col: 55, offset: 3906},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 95, col: 70, offset: 3921},
							val:        ":",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 95, col: 74, offset: 3925},
							expr: &ruleRefExpr{
								pos:  position{line: 95, col: 74, offset: 3925},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 95, col: 78, offset: 3929},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclarationWithNameAndValue",
			pos:  position{line: 99, col: 1, offset: 4014},
			expr: &actionExpr{
				pos: position{line: 99, col: 49, offset: 4062},
				run: (*parser).callonDocumentAttributeDeclarationWithNameAndValue1,
				expr: &seqExpr{
					pos: position{line: 99, col: 49, offset: 4062},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 99, col: 49, offset: 4062},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 99, col: 53, offset: 4066},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 99, col: 59, offset: 4072},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 99, col: 74, offset: 4087},
							val:        ":",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 99, col: 78, offset: 4091},
							expr: &ruleRefExpr{
								pos:  position{line: 99, col: 78, offset: 4091},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 99, col: 82, offset: 4095},
							label: "value",
							expr: &zeroOrMoreExpr{
								pos: position{line: 99, col: 88, offset: 4101},
								expr: &seqExpr{
									pos: position{line: 99, col: 89, offset: 4102},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 99, col: 89, offset: 4102},
											expr: &ruleRefExpr{
												pos:  position{line: 99, col: 90, offset: 4103},
												name: "NEWLINE",
											},
										},
										&anyMatcher{
											line: 99, col: 98, offset: 4111,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 99, col: 102, offset: 4115},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeReset",
			pos:  position{line: 103, col: 1, offset: 4218},
			expr: &choiceExpr{
				pos: position{line: 103, col: 27, offset: 4244},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 103, col: 27, offset: 4244},
						name: "DocumentAttributeResetWithSectionTitleBangSymbol",
					},
					&ruleRefExpr{
						pos:  position{line: 103, col: 78, offset: 4295},
						name: "DocumentAttributeResetWithTrailingBangSymbol",
					},
				},
			},
		},
		{
			name: "DocumentAttributeResetWithSectionTitleBangSymbol",
			pos:  position{line: 105, col: 1, offset: 4341},
			expr: &actionExpr{
				pos: position{line: 105, col: 53, offset: 4393},
				run: (*parser).callonDocumentAttributeResetWithSectionTitleBangSymbol1,
				expr: &seqExpr{
					pos: position{line: 105, col: 53, offset: 4393},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 105, col: 53, offset: 4393},
							val:        ":!",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 105, col: 58, offset: 4398},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 105, col: 64, offset: 4404},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 105, col: 79, offset: 4419},
							val:        ":",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 105, col: 83, offset: 4423},
							expr: &ruleRefExpr{
								pos:  position{line: 105, col: 83, offset: 4423},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 105, col: 87, offset: 4427},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeResetWithTrailingBangSymbol",
			pos:  position{line: 109, col: 1, offset: 4501},
			expr: &actionExpr{
				pos: position{line: 109, col: 49, offset: 4549},
				run: (*parser).callonDocumentAttributeResetWithTrailingBangSymbol1,
				expr: &seqExpr{
					pos: position{line: 109, col: 49, offset: 4549},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 109, col: 49, offset: 4549},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 109, col: 53, offset: 4553},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 109, col: 59, offset: 4559},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 109, col: 74, offset: 4574},
							val:        "!:",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 109, col: 79, offset: 4579},
							expr: &ruleRefExpr{
								pos:  position{line: 109, col: 79, offset: 4579},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 109, col: 83, offset: 4583},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeSubstitution",
			pos:  position{line: 113, col: 1, offset: 4657},
			expr: &actionExpr{
				pos: position{line: 113, col: 34, offset: 4690},
				run: (*parser).callonDocumentAttributeSubstitution1,
				expr: &seqExpr{
					pos: position{line: 113, col: 34, offset: 4690},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 113, col: 34, offset: 4690},
							val:        "{",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 113, col: 38, offset: 4694},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 113, col: 44, offset: 4700},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 113, col: 59, offset: 4715},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "AttributeName",
			pos:  position{line: 120, col: 1, offset: 4969},
			expr: &seqExpr{
				pos: position{line: 120, col: 18, offset: 4986},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 120, col: 19, offset: 4987},
						alternatives: []interface{}{
							&charClassMatcher{
								pos:        position{line: 120, col: 19, offset: 4987},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
							&charClassMatcher{
								pos:        position{line: 120, col: 27, offset: 4995},
								val:        "[a-z]",
								ranges:     []rune{'a', 'z'},
								ignoreCase: false,
								inverted:   false,
							},
							&charClassMatcher{
								pos:        position{line: 120, col: 35, offset: 5003},
								val:        "[0-9]",
								ranges:     []rune{'0', '9'},
								ignoreCase: false,
								inverted:   false,
							},
							&litMatcher{
								pos:        position{line: 120, col: 43, offset: 5011},
								val:        "_",
								ignoreCase: false,
							},
						},
					},
					&zeroOrMoreExpr{
						pos: position{line: 120, col: 48, offset: 5016},
						expr: &choiceExpr{
							pos: position{line: 120, col: 49, offset: 5017},
							alternatives: []interface{}{
								&charClassMatcher{
									pos:        position{line: 120, col: 49, offset: 5017},
									val:        "[A-Z]",
									ranges:     []rune{'A', 'Z'},
									ignoreCase: false,
									inverted:   false,
								},
								&charClassMatcher{
									pos:        position{line: 120, col: 57, offset: 5025},
									val:        "[a-z]",
									ranges:     []rune{'a', 'z'},
									ignoreCase: false,
									inverted:   false,
								},
								&charClassMatcher{
									pos:        position{line: 120, col: 65, offset: 5033},
									val:        "[0-9]",
									ranges:     []rune{'0', '9'},
									ignoreCase: false,
									inverted:   false,
								},
								&litMatcher{
									pos:        position{line: 120, col: 73, offset: 5041},
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
			pos:  position{line: 125, col: 1, offset: 5161},
			expr: &seqExpr{
				pos: position{line: 125, col: 25, offset: 5185},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 125, col: 25, offset: 5185},
						val:        "toc::[]",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 125, col: 35, offset: 5195},
						name: "NEWLINE",
					},
				},
			},
		},
		{
			name: "ElementAttribute",
			pos:  position{line: 130, col: 1, offset: 5318},
			expr: &actionExpr{
				pos: position{line: 130, col: 21, offset: 5338},
				run: (*parser).callonElementAttribute1,
				expr: &seqExpr{
					pos: position{line: 130, col: 21, offset: 5338},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 130, col: 21, offset: 5338},
							label: "attr",
							expr: &choiceExpr{
								pos: position{line: 130, col: 27, offset: 5344},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 130, col: 27, offset: 5344},
										name: "ElementID",
									},
									&ruleRefExpr{
										pos:  position{line: 130, col: 39, offset: 5356},
										name: "ElementTitle",
									},
									&ruleRefExpr{
										pos:  position{line: 130, col: 54, offset: 5371},
										name: "AdmonitionMarkerAttribute",
									},
									&ruleRefExpr{
										pos:  position{line: 130, col: 82, offset: 5399},
										name: "AttributeGroup",
									},
									&ruleRefExpr{
										pos:  position{line: 130, col: 99, offset: 5416},
										name: "InvalidElementAttribute",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 130, col: 124, offset: 5441},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ElementID",
			pos:  position{line: 134, col: 1, offset: 5532},
			expr: &choiceExpr{
				pos: position{line: 134, col: 14, offset: 5545},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 134, col: 14, offset: 5545},
						run: (*parser).callonElementID2,
						expr: &labeledExpr{
							pos:   position{line: 134, col: 14, offset: 5545},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 134, col: 18, offset: 5549},
								name: "InlineElementID",
							},
						},
					},
					&actionExpr{
						pos: position{line: 136, col: 5, offset: 5591},
						run: (*parser).callonElementID5,
						expr: &seqExpr{
							pos: position{line: 136, col: 5, offset: 5591},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 136, col: 5, offset: 5591},
									val:        "[#",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 136, col: 10, offset: 5596},
									label: "id",
									expr: &ruleRefExpr{
										pos:  position{line: 136, col: 14, offset: 5600},
										name: "ID",
									},
								},
								&litMatcher{
									pos:        position{line: 136, col: 18, offset: 5604},
									val:        "]",
									ignoreCase: false,
								},
								&zeroOrMoreExpr{
									pos: position{line: 136, col: 22, offset: 5608},
									expr: &ruleRefExpr{
										pos:  position{line: 136, col: 22, offset: 5608},
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
			pos:  position{line: 140, col: 1, offset: 5660},
			expr: &actionExpr{
				pos: position{line: 140, col: 20, offset: 5679},
				run: (*parser).callonInlineElementID1,
				expr: &seqExpr{
					pos: position{line: 140, col: 20, offset: 5679},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 140, col: 20, offset: 5679},
							val:        "[[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 140, col: 25, offset: 5684},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 140, col: 29, offset: 5688},
								name: "ID",
							},
						},
						&litMatcher{
							pos:        position{line: 140, col: 33, offset: 5692},
							val:        "]]",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 140, col: 38, offset: 5697},
							expr: &ruleRefExpr{
								pos:  position{line: 140, col: 38, offset: 5697},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "ElementTitle",
			pos:  position{line: 146, col: 1, offset: 5891},
			expr: &actionExpr{
				pos: position{line: 146, col: 17, offset: 5907},
				run: (*parser).callonElementTitle1,
				expr: &seqExpr{
					pos: position{line: 146, col: 17, offset: 5907},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 146, col: 17, offset: 5907},
							val:        ".",
							ignoreCase: false,
						},
						&notExpr{
							pos: position{line: 146, col: 21, offset: 5911},
							expr: &litMatcher{
								pos:        position{line: 146, col: 22, offset: 5912},
								val:        ".",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 146, col: 26, offset: 5916},
							expr: &ruleRefExpr{
								pos:  position{line: 146, col: 27, offset: 5917},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 146, col: 30, offset: 5920},
							label: "title",
							expr: &oneOrMoreExpr{
								pos: position{line: 146, col: 36, offset: 5926},
								expr: &seqExpr{
									pos: position{line: 146, col: 37, offset: 5927},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 146, col: 37, offset: 5927},
											expr: &ruleRefExpr{
												pos:  position{line: 146, col: 38, offset: 5928},
												name: "NEWLINE",
											},
										},
										&anyMatcher{
											line: 146, col: 46, offset: 5936,
										},
									},
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 146, col: 50, offset: 5940},
							expr: &ruleRefExpr{
								pos:  position{line: 146, col: 50, offset: 5940},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "AdmonitionMarkerAttribute",
			pos:  position{line: 151, col: 1, offset: 6085},
			expr: &actionExpr{
				pos: position{line: 151, col: 30, offset: 6114},
				run: (*parser).callonAdmonitionMarkerAttribute1,
				expr: &seqExpr{
					pos: position{line: 151, col: 30, offset: 6114},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 151, col: 30, offset: 6114},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 151, col: 34, offset: 6118},
							label: "k",
							expr: &ruleRefExpr{
								pos:  position{line: 151, col: 37, offset: 6121},
								name: "AdmonitionKind",
							},
						},
						&litMatcher{
							pos:        position{line: 151, col: 53, offset: 6137},
							val:        "]",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 151, col: 57, offset: 6141},
							expr: &ruleRefExpr{
								pos:  position{line: 151, col: 57, offset: 6141},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "AttributeGroup",
			pos:  position{line: 156, col: 1, offset: 6231},
			expr: &actionExpr{
				pos: position{line: 156, col: 19, offset: 6249},
				run: (*parser).callonAttributeGroup1,
				expr: &seqExpr{
					pos: position{line: 156, col: 19, offset: 6249},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 156, col: 19, offset: 6249},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 156, col: 23, offset: 6253},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 156, col: 34, offset: 6264},
								expr: &ruleRefExpr{
									pos:  position{line: 156, col: 35, offset: 6265},
									name: "GenericAttribute",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 156, col: 54, offset: 6284},
							val:        "]",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 156, col: 58, offset: 6288},
							expr: &ruleRefExpr{
								pos:  position{line: 156, col: 58, offset: 6288},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "GenericAttribute",
			pos:  position{line: 160, col: 1, offset: 6360},
			expr: &choiceExpr{
				pos: position{line: 160, col: 21, offset: 6380},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 160, col: 21, offset: 6380},
						run: (*parser).callonGenericAttribute2,
						expr: &seqExpr{
							pos: position{line: 160, col: 21, offset: 6380},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 160, col: 21, offset: 6380},
									label: "key",
									expr: &ruleRefExpr{
										pos:  position{line: 160, col: 26, offset: 6385},
										name: "AttributeKey",
									},
								},
								&litMatcher{
									pos:        position{line: 160, col: 40, offset: 6399},
									val:        "=",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 160, col: 44, offset: 6403},
									label: "value",
									expr: &ruleRefExpr{
										pos:  position{line: 160, col: 51, offset: 6410},
										name: "AttributeValue",
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 160, col: 67, offset: 6426},
									expr: &seqExpr{
										pos: position{line: 160, col: 68, offset: 6427},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 160, col: 68, offset: 6427},
												val:        ",",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 160, col: 72, offset: 6431},
												expr: &ruleRefExpr{
													pos:  position{line: 160, col: 72, offset: 6431},
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
						pos: position{line: 162, col: 5, offset: 6540},
						run: (*parser).callonGenericAttribute14,
						expr: &seqExpr{
							pos: position{line: 162, col: 5, offset: 6540},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 162, col: 5, offset: 6540},
									label: "key",
									expr: &ruleRefExpr{
										pos:  position{line: 162, col: 10, offset: 6545},
										name: "AttributeKey",
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 162, col: 24, offset: 6559},
									expr: &seqExpr{
										pos: position{line: 162, col: 25, offset: 6560},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 162, col: 25, offset: 6560},
												val:        ",",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 162, col: 29, offset: 6564},
												expr: &ruleRefExpr{
													pos:  position{line: 162, col: 29, offset: 6564},
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
			pos:  position{line: 166, col: 1, offset: 6658},
			expr: &actionExpr{
				pos: position{line: 166, col: 17, offset: 6674},
				run: (*parser).callonAttributeKey1,
				expr: &seqExpr{
					pos: position{line: 166, col: 17, offset: 6674},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 166, col: 17, offset: 6674},
							label: "key",
							expr: &oneOrMoreExpr{
								pos: position{line: 166, col: 22, offset: 6679},
								expr: &seqExpr{
									pos: position{line: 166, col: 23, offset: 6680},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 166, col: 23, offset: 6680},
											expr: &ruleRefExpr{
												pos:  position{line: 166, col: 24, offset: 6681},
												name: "WS",
											},
										},
										&notExpr{
											pos: position{line: 166, col: 27, offset: 6684},
											expr: &litMatcher{
												pos:        position{line: 166, col: 28, offset: 6685},
												val:        "=",
												ignoreCase: false,
											},
										},
										&notExpr{
											pos: position{line: 166, col: 32, offset: 6689},
											expr: &litMatcher{
												pos:        position{line: 166, col: 33, offset: 6690},
												val:        ",",
												ignoreCase: false,
											},
										},
										&notExpr{
											pos: position{line: 166, col: 37, offset: 6694},
											expr: &litMatcher{
												pos:        position{line: 166, col: 38, offset: 6695},
												val:        "]",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 166, col: 42, offset: 6699,
										},
									},
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 166, col: 46, offset: 6703},
							expr: &ruleRefExpr{
								pos:  position{line: 166, col: 46, offset: 6703},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "AttributeValue",
			pos:  position{line: 171, col: 1, offset: 6785},
			expr: &actionExpr{
				pos: position{line: 171, col: 19, offset: 6803},
				run: (*parser).callonAttributeValue1,
				expr: &seqExpr{
					pos: position{line: 171, col: 19, offset: 6803},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 171, col: 19, offset: 6803},
							expr: &ruleRefExpr{
								pos:  position{line: 171, col: 19, offset: 6803},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 171, col: 23, offset: 6807},
							label: "value",
							expr: &zeroOrMoreExpr{
								pos: position{line: 171, col: 29, offset: 6813},
								expr: &seqExpr{
									pos: position{line: 171, col: 30, offset: 6814},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 171, col: 30, offset: 6814},
											expr: &ruleRefExpr{
												pos:  position{line: 171, col: 31, offset: 6815},
												name: "WS",
											},
										},
										&notExpr{
											pos: position{line: 171, col: 34, offset: 6818},
											expr: &litMatcher{
												pos:        position{line: 171, col: 35, offset: 6819},
												val:        "=",
												ignoreCase: false,
											},
										},
										&notExpr{
											pos: position{line: 171, col: 39, offset: 6823},
											expr: &litMatcher{
												pos:        position{line: 171, col: 40, offset: 6824},
												val:        "]",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 171, col: 44, offset: 6828,
										},
									},
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 171, col: 48, offset: 6832},
							expr: &ruleRefExpr{
								pos:  position{line: 171, col: 48, offset: 6832},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "InvalidElementAttribute",
			pos:  position{line: 176, col: 1, offset: 6919},
			expr: &actionExpr{
				pos: position{line: 176, col: 28, offset: 6946},
				run: (*parser).callonInvalidElementAttribute1,
				expr: &seqExpr{
					pos: position{line: 176, col: 28, offset: 6946},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 176, col: 28, offset: 6946},
							val:        "[",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 176, col: 32, offset: 6950},
							expr: &ruleRefExpr{
								pos:  position{line: 176, col: 32, offset: 6950},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 176, col: 36, offset: 6954},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 176, col: 44, offset: 6962},
								expr: &seqExpr{
									pos: position{line: 176, col: 45, offset: 6963},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 176, col: 45, offset: 6963},
											expr: &litMatcher{
												pos:        position{line: 176, col: 46, offset: 6964},
												val:        "]",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 176, col: 50, offset: 6968,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 176, col: 54, offset: 6972},
							val:        "]",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 176, col: 58, offset: 6976},
							expr: &ruleRefExpr{
								pos:  position{line: 176, col: 58, offset: 6976},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "Section",
			pos:  position{line: 183, col: 1, offset: 7141},
			expr: &choiceExpr{
				pos: position{line: 183, col: 12, offset: 7152},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 183, col: 12, offset: 7152},
						name: "Section1",
					},
					&ruleRefExpr{
						pos:  position{line: 183, col: 23, offset: 7163},
						name: "Section2",
					},
					&ruleRefExpr{
						pos:  position{line: 183, col: 34, offset: 7174},
						name: "Section3",
					},
					&ruleRefExpr{
						pos:  position{line: 183, col: 45, offset: 7185},
						name: "Section4",
					},
					&ruleRefExpr{
						pos:  position{line: 183, col: 56, offset: 7196},
						name: "Section5",
					},
				},
			},
		},
		{
			name: "Section1",
			pos:  position{line: 186, col: 1, offset: 7207},
			expr: &actionExpr{
				pos: position{line: 186, col: 13, offset: 7219},
				run: (*parser).callonSection11,
				expr: &seqExpr{
					pos: position{line: 186, col: 13, offset: 7219},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 186, col: 13, offset: 7219},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 186, col: 21, offset: 7227},
								name: "Section1Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 186, col: 36, offset: 7242},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 186, col: 46, offset: 7252},
								expr: &ruleRefExpr{
									pos:  position{line: 186, col: 46, offset: 7252},
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
			pos:  position{line: 190, col: 1, offset: 7359},
			expr: &actionExpr{
				pos: position{line: 190, col: 18, offset: 7376},
				run: (*parser).callonSection1Block1,
				expr: &seqExpr{
					pos: position{line: 190, col: 18, offset: 7376},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 190, col: 18, offset: 7376},
							expr: &ruleRefExpr{
								pos:  position{line: 190, col: 19, offset: 7377},
								name: "Section1",
							},
						},
						&labeledExpr{
							pos:   position{line: 190, col: 28, offset: 7386},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 190, col: 37, offset: 7395},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 190, col: 37, offset: 7395},
										name: "Section2",
									},
									&ruleRefExpr{
										pos:  position{line: 190, col: 48, offset: 7406},
										name: "Section3",
									},
									&ruleRefExpr{
										pos:  position{line: 190, col: 59, offset: 7417},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 190, col: 70, offset: 7428},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 190, col: 81, offset: 7439},
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
			pos:  position{line: 194, col: 1, offset: 7501},
			expr: &actionExpr{
				pos: position{line: 194, col: 13, offset: 7513},
				run: (*parser).callonSection21,
				expr: &seqExpr{
					pos: position{line: 194, col: 13, offset: 7513},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 194, col: 13, offset: 7513},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 194, col: 21, offset: 7521},
								name: "Section2Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 194, col: 36, offset: 7536},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 194, col: 46, offset: 7546},
								expr: &ruleRefExpr{
									pos:  position{line: 194, col: 46, offset: 7546},
									name: "Section2Block",
								},
							},
						},
						&andExpr{
							pos: position{line: 194, col: 62, offset: 7562},
							expr: &zeroOrMoreExpr{
								pos: position{line: 194, col: 63, offset: 7563},
								expr: &ruleRefExpr{
									pos:  position{line: 194, col: 64, offset: 7564},
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
			pos:  position{line: 198, col: 1, offset: 7666},
			expr: &actionExpr{
				pos: position{line: 198, col: 18, offset: 7683},
				run: (*parser).callonSection2Block1,
				expr: &seqExpr{
					pos: position{line: 198, col: 18, offset: 7683},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 198, col: 18, offset: 7683},
							expr: &ruleRefExpr{
								pos:  position{line: 198, col: 19, offset: 7684},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 198, col: 28, offset: 7693},
							expr: &ruleRefExpr{
								pos:  position{line: 198, col: 29, offset: 7694},
								name: "Section2",
							},
						},
						&labeledExpr{
							pos:   position{line: 198, col: 38, offset: 7703},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 198, col: 47, offset: 7712},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 198, col: 47, offset: 7712},
										name: "Section3",
									},
									&ruleRefExpr{
										pos:  position{line: 198, col: 58, offset: 7723},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 198, col: 69, offset: 7734},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 198, col: 80, offset: 7745},
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
			pos:  position{line: 202, col: 1, offset: 7807},
			expr: &actionExpr{
				pos: position{line: 202, col: 13, offset: 7819},
				run: (*parser).callonSection31,
				expr: &seqExpr{
					pos: position{line: 202, col: 13, offset: 7819},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 202, col: 13, offset: 7819},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 202, col: 21, offset: 7827},
								name: "Section3Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 202, col: 36, offset: 7842},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 202, col: 46, offset: 7852},
								expr: &ruleRefExpr{
									pos:  position{line: 202, col: 46, offset: 7852},
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
			pos:  position{line: 206, col: 1, offset: 7959},
			expr: &actionExpr{
				pos: position{line: 206, col: 18, offset: 7976},
				run: (*parser).callonSection3Block1,
				expr: &seqExpr{
					pos: position{line: 206, col: 18, offset: 7976},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 206, col: 18, offset: 7976},
							expr: &ruleRefExpr{
								pos:  position{line: 206, col: 19, offset: 7977},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 206, col: 28, offset: 7986},
							expr: &ruleRefExpr{
								pos:  position{line: 206, col: 29, offset: 7987},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 206, col: 38, offset: 7996},
							expr: &ruleRefExpr{
								pos:  position{line: 206, col: 39, offset: 7997},
								name: "Section3",
							},
						},
						&labeledExpr{
							pos:   position{line: 206, col: 48, offset: 8006},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 206, col: 57, offset: 8015},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 206, col: 57, offset: 8015},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 206, col: 68, offset: 8026},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 206, col: 79, offset: 8037},
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
			pos:  position{line: 210, col: 1, offset: 8099},
			expr: &actionExpr{
				pos: position{line: 210, col: 13, offset: 8111},
				run: (*parser).callonSection41,
				expr: &seqExpr{
					pos: position{line: 210, col: 13, offset: 8111},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 210, col: 13, offset: 8111},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 210, col: 21, offset: 8119},
								name: "Section4Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 210, col: 36, offset: 8134},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 210, col: 46, offset: 8144},
								expr: &ruleRefExpr{
									pos:  position{line: 210, col: 46, offset: 8144},
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
			pos:  position{line: 214, col: 1, offset: 8251},
			expr: &actionExpr{
				pos: position{line: 214, col: 18, offset: 8268},
				run: (*parser).callonSection4Block1,
				expr: &seqExpr{
					pos: position{line: 214, col: 18, offset: 8268},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 214, col: 18, offset: 8268},
							expr: &ruleRefExpr{
								pos:  position{line: 214, col: 19, offset: 8269},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 214, col: 28, offset: 8278},
							expr: &ruleRefExpr{
								pos:  position{line: 214, col: 29, offset: 8279},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 214, col: 38, offset: 8288},
							expr: &ruleRefExpr{
								pos:  position{line: 214, col: 39, offset: 8289},
								name: "Section3",
							},
						},
						&notExpr{
							pos: position{line: 214, col: 48, offset: 8298},
							expr: &ruleRefExpr{
								pos:  position{line: 214, col: 49, offset: 8299},
								name: "Section4",
							},
						},
						&labeledExpr{
							pos:   position{line: 214, col: 58, offset: 8308},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 214, col: 67, offset: 8317},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 214, col: 67, offset: 8317},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 214, col: 78, offset: 8328},
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
			pos:  position{line: 218, col: 1, offset: 8390},
			expr: &actionExpr{
				pos: position{line: 218, col: 13, offset: 8402},
				run: (*parser).callonSection51,
				expr: &seqExpr{
					pos: position{line: 218, col: 13, offset: 8402},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 218, col: 13, offset: 8402},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 218, col: 21, offset: 8410},
								name: "Section5Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 218, col: 36, offset: 8425},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 218, col: 46, offset: 8435},
								expr: &ruleRefExpr{
									pos:  position{line: 218, col: 46, offset: 8435},
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
			pos:  position{line: 222, col: 1, offset: 8542},
			expr: &actionExpr{
				pos: position{line: 222, col: 18, offset: 8559},
				run: (*parser).callonSection5Block1,
				expr: &seqExpr{
					pos: position{line: 222, col: 18, offset: 8559},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 222, col: 18, offset: 8559},
							expr: &ruleRefExpr{
								pos:  position{line: 222, col: 19, offset: 8560},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 222, col: 28, offset: 8569},
							expr: &ruleRefExpr{
								pos:  position{line: 222, col: 29, offset: 8570},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 222, col: 38, offset: 8579},
							expr: &ruleRefExpr{
								pos:  position{line: 222, col: 39, offset: 8580},
								name: "Section3",
							},
						},
						&notExpr{
							pos: position{line: 222, col: 48, offset: 8589},
							expr: &ruleRefExpr{
								pos:  position{line: 222, col: 49, offset: 8590},
								name: "Section4",
							},
						},
						&notExpr{
							pos: position{line: 222, col: 58, offset: 8599},
							expr: &ruleRefExpr{
								pos:  position{line: 222, col: 59, offset: 8600},
								name: "Section5",
							},
						},
						&labeledExpr{
							pos:   position{line: 222, col: 68, offset: 8609},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 222, col: 77, offset: 8618},
								name: "BlockElement",
							},
						},
					},
				},
			},
		},
		{
			name: "SectionTitle",
			pos:  position{line: 230, col: 1, offset: 8791},
			expr: &choiceExpr{
				pos: position{line: 230, col: 17, offset: 8807},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 230, col: 17, offset: 8807},
						name: "Section1Title",
					},
					&ruleRefExpr{
						pos:  position{line: 230, col: 33, offset: 8823},
						name: "Section2Title",
					},
					&ruleRefExpr{
						pos:  position{line: 230, col: 49, offset: 8839},
						name: "Section3Title",
					},
					&ruleRefExpr{
						pos:  position{line: 230, col: 65, offset: 8855},
						name: "Section4Title",
					},
					&ruleRefExpr{
						pos:  position{line: 230, col: 81, offset: 8871},
						name: "Section5Title",
					},
				},
			},
		},
		{
			name: "Section1Title",
			pos:  position{line: 232, col: 1, offset: 8886},
			expr: &actionExpr{
				pos: position{line: 232, col: 18, offset: 8903},
				run: (*parser).callonSection1Title1,
				expr: &seqExpr{
					pos: position{line: 232, col: 18, offset: 8903},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 232, col: 18, offset: 8903},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 232, col: 29, offset: 8914},
								expr: &ruleRefExpr{
									pos:  position{line: 232, col: 30, offset: 8915},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 232, col: 49, offset: 8934},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 232, col: 56, offset: 8941},
								val:        "==",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 232, col: 62, offset: 8947},
							expr: &ruleRefExpr{
								pos:  position{line: 232, col: 62, offset: 8947},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 232, col: 66, offset: 8951},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 232, col: 75, offset: 8960},
								name: "InlineContent",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 232, col: 90, offset: 8975},
							expr: &ruleRefExpr{
								pos:  position{line: 232, col: 90, offset: 8975},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 232, col: 94, offset: 8979},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 232, col: 97, offset: 8982},
								expr: &ruleRefExpr{
									pos:  position{line: 232, col: 98, offset: 8983},
									name: "InlineElementID",
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 232, col: 116, offset: 9001},
							expr: &ruleRefExpr{
								pos:  position{line: 232, col: 116, offset: 9001},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 232, col: 120, offset: 9005},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 232, col: 125, offset: 9010},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 232, col: 125, offset: 9010},
									expr: &ruleRefExpr{
										pos:  position{line: 232, col: 125, offset: 9010},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 232, col: 138, offset: 9023},
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
			pos:  position{line: 236, col: 1, offset: 9138},
			expr: &actionExpr{
				pos: position{line: 236, col: 18, offset: 9155},
				run: (*parser).callonSection2Title1,
				expr: &seqExpr{
					pos: position{line: 236, col: 18, offset: 9155},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 236, col: 18, offset: 9155},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 236, col: 29, offset: 9166},
								expr: &ruleRefExpr{
									pos:  position{line: 236, col: 30, offset: 9167},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 236, col: 49, offset: 9186},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 236, col: 56, offset: 9193},
								val:        "===",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 236, col: 63, offset: 9200},
							expr: &ruleRefExpr{
								pos:  position{line: 236, col: 63, offset: 9200},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 236, col: 67, offset: 9204},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 236, col: 76, offset: 9213},
								name: "InlineContent",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 236, col: 91, offset: 9228},
							expr: &ruleRefExpr{
								pos:  position{line: 236, col: 91, offset: 9228},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 236, col: 95, offset: 9232},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 236, col: 98, offset: 9235},
								expr: &ruleRefExpr{
									pos:  position{line: 236, col: 99, offset: 9236},
									name: "InlineElementID",
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 236, col: 117, offset: 9254},
							expr: &ruleRefExpr{
								pos:  position{line: 236, col: 117, offset: 9254},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 236, col: 121, offset: 9258},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 236, col: 126, offset: 9263},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 236, col: 126, offset: 9263},
									expr: &ruleRefExpr{
										pos:  position{line: 236, col: 126, offset: 9263},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 236, col: 139, offset: 9276},
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
			pos:  position{line: 240, col: 1, offset: 9390},
			expr: &actionExpr{
				pos: position{line: 240, col: 18, offset: 9407},
				run: (*parser).callonSection3Title1,
				expr: &seqExpr{
					pos: position{line: 240, col: 18, offset: 9407},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 240, col: 18, offset: 9407},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 240, col: 29, offset: 9418},
								expr: &ruleRefExpr{
									pos:  position{line: 240, col: 30, offset: 9419},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 240, col: 49, offset: 9438},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 240, col: 56, offset: 9445},
								val:        "====",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 240, col: 64, offset: 9453},
							expr: &ruleRefExpr{
								pos:  position{line: 240, col: 64, offset: 9453},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 240, col: 68, offset: 9457},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 240, col: 77, offset: 9466},
								name: "InlineContent",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 240, col: 92, offset: 9481},
							expr: &ruleRefExpr{
								pos:  position{line: 240, col: 92, offset: 9481},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 240, col: 96, offset: 9485},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 240, col: 99, offset: 9488},
								expr: &ruleRefExpr{
									pos:  position{line: 240, col: 100, offset: 9489},
									name: "InlineElementID",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 240, col: 118, offset: 9507},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 240, col: 123, offset: 9512},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 240, col: 123, offset: 9512},
									expr: &ruleRefExpr{
										pos:  position{line: 240, col: 123, offset: 9512},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 240, col: 136, offset: 9525},
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
			pos:  position{line: 244, col: 1, offset: 9639},
			expr: &actionExpr{
				pos: position{line: 244, col: 18, offset: 9656},
				run: (*parser).callonSection4Title1,
				expr: &seqExpr{
					pos: position{line: 244, col: 18, offset: 9656},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 244, col: 18, offset: 9656},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 244, col: 29, offset: 9667},
								expr: &ruleRefExpr{
									pos:  position{line: 244, col: 30, offset: 9668},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 244, col: 49, offset: 9687},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 244, col: 56, offset: 9694},
								val:        "=====",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 244, col: 65, offset: 9703},
							expr: &ruleRefExpr{
								pos:  position{line: 244, col: 65, offset: 9703},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 244, col: 69, offset: 9707},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 244, col: 78, offset: 9716},
								name: "InlineContent",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 244, col: 93, offset: 9731},
							expr: &ruleRefExpr{
								pos:  position{line: 244, col: 93, offset: 9731},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 244, col: 97, offset: 9735},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 244, col: 100, offset: 9738},
								expr: &ruleRefExpr{
									pos:  position{line: 244, col: 101, offset: 9739},
									name: "InlineElementID",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 244, col: 119, offset: 9757},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 244, col: 124, offset: 9762},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 244, col: 124, offset: 9762},
									expr: &ruleRefExpr{
										pos:  position{line: 244, col: 124, offset: 9762},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 244, col: 137, offset: 9775},
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
			pos:  position{line: 248, col: 1, offset: 9889},
			expr: &actionExpr{
				pos: position{line: 248, col: 18, offset: 9906},
				run: (*parser).callonSection5Title1,
				expr: &seqExpr{
					pos: position{line: 248, col: 18, offset: 9906},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 248, col: 18, offset: 9906},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 248, col: 29, offset: 9917},
								expr: &ruleRefExpr{
									pos:  position{line: 248, col: 30, offset: 9918},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 248, col: 49, offset: 9937},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 248, col: 56, offset: 9944},
								val:        "======",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 248, col: 66, offset: 9954},
							expr: &ruleRefExpr{
								pos:  position{line: 248, col: 66, offset: 9954},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 248, col: 70, offset: 9958},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 248, col: 79, offset: 9967},
								name: "InlineContent",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 248, col: 94, offset: 9982},
							expr: &ruleRefExpr{
								pos:  position{line: 248, col: 94, offset: 9982},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 248, col: 98, offset: 9986},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 248, col: 101, offset: 9989},
								expr: &ruleRefExpr{
									pos:  position{line: 248, col: 102, offset: 9990},
									name: "InlineElementID",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 248, col: 120, offset: 10008},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 248, col: 125, offset: 10013},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 248, col: 125, offset: 10013},
									expr: &ruleRefExpr{
										pos:  position{line: 248, col: 125, offset: 10013},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 248, col: 138, offset: 10026},
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
			pos:  position{line: 255, col: 1, offset: 10241},
			expr: &actionExpr{
				pos: position{line: 255, col: 9, offset: 10249},
				run: (*parser).callonList1,
				expr: &seqExpr{
					pos: position{line: 255, col: 9, offset: 10249},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 255, col: 9, offset: 10249},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 255, col: 20, offset: 10260},
								expr: &ruleRefExpr{
									pos:  position{line: 255, col: 21, offset: 10261},
									name: "ListAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 257, col: 5, offset: 10350},
							label: "elements",
							expr: &ruleRefExpr{
								pos:  position{line: 257, col: 14, offset: 10359},
								name: "ListItems",
							},
						},
					},
				},
			},
		},
		{
			name: "ListItems",
			pos:  position{line: 261, col: 1, offset: 10453},
			expr: &oneOrMoreExpr{
				pos: position{line: 261, col: 14, offset: 10466},
				expr: &choiceExpr{
					pos: position{line: 261, col: 15, offset: 10467},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 261, col: 15, offset: 10467},
							name: "OrderedListItem",
						},
						&ruleRefExpr{
							pos:  position{line: 261, col: 33, offset: 10485},
							name: "UnorderedListItem",
						},
						&ruleRefExpr{
							pos:  position{line: 261, col: 53, offset: 10505},
							name: "LabeledListItem",
						},
					},
				},
			},
		},
		{
			name: "ListAttribute",
			pos:  position{line: 263, col: 1, offset: 10524},
			expr: &actionExpr{
				pos: position{line: 263, col: 18, offset: 10541},
				run: (*parser).callonListAttribute1,
				expr: &seqExpr{
					pos: position{line: 263, col: 18, offset: 10541},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 263, col: 18, offset: 10541},
							label: "attribute",
							expr: &choiceExpr{
								pos: position{line: 263, col: 29, offset: 10552},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 263, col: 29, offset: 10552},
										name: "HorizontalLayout",
									},
									&ruleRefExpr{
										pos:  position{line: 263, col: 48, offset: 10571},
										name: "ListID",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 263, col: 56, offset: 10579},
							name: "NEWLINE",
						},
					},
				},
			},
		},
		{
			name: "ListID",
			pos:  position{line: 267, col: 1, offset: 10618},
			expr: &actionExpr{
				pos: position{line: 267, col: 11, offset: 10628},
				run: (*parser).callonListID1,
				expr: &seqExpr{
					pos: position{line: 267, col: 11, offset: 10628},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 267, col: 11, offset: 10628},
							val:        "[#",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 267, col: 16, offset: 10633},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 267, col: 20, offset: 10637},
								name: "ID",
							},
						},
						&litMatcher{
							pos:        position{line: 267, col: 24, offset: 10641},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "HorizontalLayout",
			pos:  position{line: 271, col: 1, offset: 10707},
			expr: &actionExpr{
				pos: position{line: 271, col: 21, offset: 10727},
				run: (*parser).callonHorizontalLayout1,
				expr: &litMatcher{
					pos:        position{line: 271, col: 21, offset: 10727},
					val:        "[horizontal]",
					ignoreCase: false,
				},
			},
		},
		{
			name: "ListParagraph",
			pos:  position{line: 275, col: 1, offset: 10810},
			expr: &actionExpr{
				pos: position{line: 275, col: 19, offset: 10828},
				run: (*parser).callonListParagraph1,
				expr: &labeledExpr{
					pos:   position{line: 275, col: 19, offset: 10828},
					label: "lines",
					expr: &oneOrMoreExpr{
						pos: position{line: 275, col: 25, offset: 10834},
						expr: &seqExpr{
							pos: position{line: 276, col: 5, offset: 10840},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 276, col: 5, offset: 10840},
									expr: &ruleRefExpr{
										pos:  position{line: 276, col: 7, offset: 10842},
										name: "OrderedListItemPrefix",
									},
								},
								&notExpr{
									pos: position{line: 277, col: 5, offset: 10870},
									expr: &ruleRefExpr{
										pos:  position{line: 277, col: 7, offset: 10872},
										name: "UnorderedListItemPrefix",
									},
								},
								&notExpr{
									pos: position{line: 278, col: 5, offset: 10902},
									expr: &seqExpr{
										pos: position{line: 278, col: 7, offset: 10904},
										exprs: []interface{}{
											&ruleRefExpr{
												pos:  position{line: 278, col: 7, offset: 10904},
												name: "LabeledListItemTerm",
											},
											&ruleRefExpr{
												pos:  position{line: 278, col: 27, offset: 10924},
												name: "LabeledListItemSeparator",
											},
										},
									},
								},
								&notExpr{
									pos: position{line: 279, col: 5, offset: 10955},
									expr: &ruleRefExpr{
										pos:  position{line: 279, col: 7, offset: 10957},
										name: "ListItemContinuation",
									},
								},
								&notExpr{
									pos: position{line: 280, col: 5, offset: 10984},
									expr: &ruleRefExpr{
										pos:  position{line: 280, col: 7, offset: 10986},
										name: "ElementAttribute",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 281, col: 5, offset: 11008},
									name: "InlineContentWithTrailingSpaces",
								},
								&ruleRefExpr{
									pos:  position{line: 281, col: 37, offset: 11040},
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
			pos:  position{line: 285, col: 1, offset: 11109},
			expr: &actionExpr{
				pos: position{line: 285, col: 25, offset: 11133},
				run: (*parser).callonListItemContinuation1,
				expr: &seqExpr{
					pos: position{line: 285, col: 25, offset: 11133},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 285, col: 25, offset: 11133},
							val:        "+",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 285, col: 29, offset: 11137},
							expr: &ruleRefExpr{
								pos:  position{line: 285, col: 29, offset: 11137},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 285, col: 33, offset: 11141},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ContinuedBlockElement",
			pos:  position{line: 289, col: 1, offset: 11193},
			expr: &actionExpr{
				pos: position{line: 289, col: 26, offset: 11218},
				run: (*parser).callonContinuedBlockElement1,
				expr: &seqExpr{
					pos: position{line: 289, col: 26, offset: 11218},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 289, col: 26, offset: 11218},
							name: "ListItemContinuation",
						},
						&labeledExpr{
							pos:   position{line: 289, col: 47, offset: 11239},
							label: "element",
							expr: &ruleRefExpr{
								pos:  position{line: 289, col: 55, offset: 11247},
								name: "BlockElement",
							},
						},
					},
				},
			},
		},
		{
			name: "OrderedListItem",
			pos:  position{line: 296, col: 1, offset: 11403},
			expr: &actionExpr{
				pos: position{line: 296, col: 20, offset: 11422},
				run: (*parser).callonOrderedListItem1,
				expr: &seqExpr{
					pos: position{line: 296, col: 20, offset: 11422},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 296, col: 20, offset: 11422},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 296, col: 31, offset: 11433},
								expr: &ruleRefExpr{
									pos:  position{line: 296, col: 32, offset: 11434},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 296, col: 51, offset: 11453},
							label: "prefix",
							expr: &ruleRefExpr{
								pos:  position{line: 296, col: 59, offset: 11461},
								name: "OrderedListItemPrefix",
							},
						},
						&labeledExpr{
							pos:   position{line: 296, col: 82, offset: 11484},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 296, col: 91, offset: 11493},
								name: "OrderedListItemContent",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 296, col: 115, offset: 11517},
							expr: &ruleRefExpr{
								pos:  position{line: 296, col: 115, offset: 11517},
								name: "BlankLine",
							},
						},
					},
				},
			},
		},
		{
			name: "OrderedListItemPrefix",
			pos:  position{line: 300, col: 1, offset: 11665},
			expr: &choiceExpr{
				pos: position{line: 302, col: 1, offset: 11729},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 302, col: 1, offset: 11729},
						run: (*parser).callonOrderedListItemPrefix2,
						expr: &seqExpr{
							pos: position{line: 302, col: 1, offset: 11729},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 302, col: 1, offset: 11729},
									expr: &ruleRefExpr{
										pos:  position{line: 302, col: 1, offset: 11729},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 302, col: 5, offset: 11733},
									label: "style",
									expr: &litMatcher{
										pos:        position{line: 302, col: 12, offset: 11740},
										val:        ".",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 302, col: 17, offset: 11745},
									expr: &ruleRefExpr{
										pos:  position{line: 302, col: 17, offset: 11745},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 304, col: 5, offset: 11838},
						run: (*parser).callonOrderedListItemPrefix10,
						expr: &seqExpr{
							pos: position{line: 304, col: 5, offset: 11838},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 304, col: 5, offset: 11838},
									expr: &ruleRefExpr{
										pos:  position{line: 304, col: 5, offset: 11838},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 304, col: 9, offset: 11842},
									label: "style",
									expr: &litMatcher{
										pos:        position{line: 304, col: 16, offset: 11849},
										val:        "..",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 304, col: 22, offset: 11855},
									expr: &ruleRefExpr{
										pos:  position{line: 304, col: 22, offset: 11855},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 306, col: 5, offset: 11953},
						run: (*parser).callonOrderedListItemPrefix18,
						expr: &seqExpr{
							pos: position{line: 306, col: 5, offset: 11953},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 306, col: 5, offset: 11953},
									expr: &ruleRefExpr{
										pos:  position{line: 306, col: 5, offset: 11953},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 306, col: 9, offset: 11957},
									label: "style",
									expr: &litMatcher{
										pos:        position{line: 306, col: 16, offset: 11964},
										val:        "...",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 306, col: 23, offset: 11971},
									expr: &ruleRefExpr{
										pos:  position{line: 306, col: 23, offset: 11971},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 308, col: 5, offset: 12070},
						run: (*parser).callonOrderedListItemPrefix26,
						expr: &seqExpr{
							pos: position{line: 308, col: 5, offset: 12070},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 308, col: 5, offset: 12070},
									expr: &ruleRefExpr{
										pos:  position{line: 308, col: 5, offset: 12070},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 308, col: 9, offset: 12074},
									label: "style",
									expr: &litMatcher{
										pos:        position{line: 308, col: 16, offset: 12081},
										val:        "....",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 308, col: 24, offset: 12089},
									expr: &ruleRefExpr{
										pos:  position{line: 308, col: 24, offset: 12089},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 310, col: 5, offset: 12189},
						run: (*parser).callonOrderedListItemPrefix34,
						expr: &seqExpr{
							pos: position{line: 310, col: 5, offset: 12189},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 310, col: 5, offset: 12189},
									expr: &ruleRefExpr{
										pos:  position{line: 310, col: 5, offset: 12189},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 310, col: 9, offset: 12193},
									label: "style",
									expr: &litMatcher{
										pos:        position{line: 310, col: 16, offset: 12200},
										val:        ".....",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 310, col: 25, offset: 12209},
									expr: &ruleRefExpr{
										pos:  position{line: 310, col: 25, offset: 12209},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 313, col: 5, offset: 12332},
						run: (*parser).callonOrderedListItemPrefix42,
						expr: &seqExpr{
							pos: position{line: 313, col: 5, offset: 12332},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 313, col: 5, offset: 12332},
									expr: &ruleRefExpr{
										pos:  position{line: 313, col: 5, offset: 12332},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 313, col: 9, offset: 12336},
									label: "style",
									expr: &seqExpr{
										pos: position{line: 313, col: 16, offset: 12343},
										exprs: []interface{}{
											&oneOrMoreExpr{
												pos: position{line: 313, col: 16, offset: 12343},
												expr: &seqExpr{
													pos: position{line: 313, col: 17, offset: 12344},
													exprs: []interface{}{
														&notExpr{
															pos: position{line: 313, col: 17, offset: 12344},
															expr: &litMatcher{
																pos:        position{line: 313, col: 18, offset: 12345},
																val:        ".",
																ignoreCase: false,
															},
														},
														&notExpr{
															pos: position{line: 313, col: 22, offset: 12349},
															expr: &ruleRefExpr{
																pos:  position{line: 313, col: 23, offset: 12350},
																name: "WS",
															},
														},
														&notExpr{
															pos: position{line: 313, col: 26, offset: 12353},
															expr: &ruleRefExpr{
																pos:  position{line: 313, col: 27, offset: 12354},
																name: "NEWLINE",
															},
														},
														&charClassMatcher{
															pos:        position{line: 313, col: 35, offset: 12362},
															val:        "[0-9]",
															ranges:     []rune{'0', '9'},
															ignoreCase: false,
															inverted:   false,
														},
													},
												},
											},
											&litMatcher{
												pos:        position{line: 313, col: 43, offset: 12370},
												val:        ".",
												ignoreCase: false,
											},
										},
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 313, col: 48, offset: 12375},
									expr: &ruleRefExpr{
										pos:  position{line: 313, col: 48, offset: 12375},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 315, col: 5, offset: 12470},
						run: (*parser).callonOrderedListItemPrefix60,
						expr: &seqExpr{
							pos: position{line: 315, col: 5, offset: 12470},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 315, col: 5, offset: 12470},
									expr: &ruleRefExpr{
										pos:  position{line: 315, col: 5, offset: 12470},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 315, col: 9, offset: 12474},
									label: "style",
									expr: &seqExpr{
										pos: position{line: 315, col: 16, offset: 12481},
										exprs: []interface{}{
											&oneOrMoreExpr{
												pos: position{line: 315, col: 16, offset: 12481},
												expr: &seqExpr{
													pos: position{line: 315, col: 17, offset: 12482},
													exprs: []interface{}{
														&notExpr{
															pos: position{line: 315, col: 17, offset: 12482},
															expr: &litMatcher{
																pos:        position{line: 315, col: 18, offset: 12483},
																val:        ".",
																ignoreCase: false,
															},
														},
														&notExpr{
															pos: position{line: 315, col: 22, offset: 12487},
															expr: &ruleRefExpr{
																pos:  position{line: 315, col: 23, offset: 12488},
																name: "WS",
															},
														},
														&notExpr{
															pos: position{line: 315, col: 26, offset: 12491},
															expr: &ruleRefExpr{
																pos:  position{line: 315, col: 27, offset: 12492},
																name: "NEWLINE",
															},
														},
														&charClassMatcher{
															pos:        position{line: 315, col: 35, offset: 12500},
															val:        "[a-z]",
															ranges:     []rune{'a', 'z'},
															ignoreCase: false,
															inverted:   false,
														},
													},
												},
											},
											&litMatcher{
												pos:        position{line: 315, col: 43, offset: 12508},
												val:        ".",
												ignoreCase: false,
											},
										},
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 315, col: 48, offset: 12513},
									expr: &ruleRefExpr{
										pos:  position{line: 315, col: 48, offset: 12513},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 317, col: 5, offset: 12611},
						run: (*parser).callonOrderedListItemPrefix78,
						expr: &seqExpr{
							pos: position{line: 317, col: 5, offset: 12611},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 317, col: 5, offset: 12611},
									expr: &ruleRefExpr{
										pos:  position{line: 317, col: 5, offset: 12611},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 317, col: 9, offset: 12615},
									label: "style",
									expr: &seqExpr{
										pos: position{line: 317, col: 16, offset: 12622},
										exprs: []interface{}{
											&oneOrMoreExpr{
												pos: position{line: 317, col: 16, offset: 12622},
												expr: &seqExpr{
													pos: position{line: 317, col: 17, offset: 12623},
													exprs: []interface{}{
														&notExpr{
															pos: position{line: 317, col: 17, offset: 12623},
															expr: &litMatcher{
																pos:        position{line: 317, col: 18, offset: 12624},
																val:        ".",
																ignoreCase: false,
															},
														},
														&notExpr{
															pos: position{line: 317, col: 22, offset: 12628},
															expr: &ruleRefExpr{
																pos:  position{line: 317, col: 23, offset: 12629},
																name: "WS",
															},
														},
														&notExpr{
															pos: position{line: 317, col: 26, offset: 12632},
															expr: &ruleRefExpr{
																pos:  position{line: 317, col: 27, offset: 12633},
																name: "NEWLINE",
															},
														},
														&charClassMatcher{
															pos:        position{line: 317, col: 35, offset: 12641},
															val:        "[A-Z]",
															ranges:     []rune{'A', 'Z'},
															ignoreCase: false,
															inverted:   false,
														},
													},
												},
											},
											&litMatcher{
												pos:        position{line: 317, col: 43, offset: 12649},
												val:        ".",
												ignoreCase: false,
											},
										},
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 317, col: 48, offset: 12654},
									expr: &ruleRefExpr{
										pos:  position{line: 317, col: 48, offset: 12654},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 319, col: 5, offset: 12752},
						run: (*parser).callonOrderedListItemPrefix96,
						expr: &seqExpr{
							pos: position{line: 319, col: 5, offset: 12752},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 319, col: 5, offset: 12752},
									expr: &ruleRefExpr{
										pos:  position{line: 319, col: 5, offset: 12752},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 319, col: 9, offset: 12756},
									label: "style",
									expr: &seqExpr{
										pos: position{line: 319, col: 16, offset: 12763},
										exprs: []interface{}{
											&oneOrMoreExpr{
												pos: position{line: 319, col: 16, offset: 12763},
												expr: &seqExpr{
													pos: position{line: 319, col: 17, offset: 12764},
													exprs: []interface{}{
														&notExpr{
															pos: position{line: 319, col: 17, offset: 12764},
															expr: &litMatcher{
																pos:        position{line: 319, col: 18, offset: 12765},
																val:        ")",
																ignoreCase: false,
															},
														},
														&notExpr{
															pos: position{line: 319, col: 22, offset: 12769},
															expr: &ruleRefExpr{
																pos:  position{line: 319, col: 23, offset: 12770},
																name: "WS",
															},
														},
														&notExpr{
															pos: position{line: 319, col: 26, offset: 12773},
															expr: &ruleRefExpr{
																pos:  position{line: 319, col: 27, offset: 12774},
																name: "NEWLINE",
															},
														},
														&charClassMatcher{
															pos:        position{line: 319, col: 35, offset: 12782},
															val:        "[a-z]",
															ranges:     []rune{'a', 'z'},
															ignoreCase: false,
															inverted:   false,
														},
													},
												},
											},
											&litMatcher{
												pos:        position{line: 319, col: 43, offset: 12790},
												val:        ")",
												ignoreCase: false,
											},
										},
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 319, col: 48, offset: 12795},
									expr: &ruleRefExpr{
										pos:  position{line: 319, col: 48, offset: 12795},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 321, col: 5, offset: 12893},
						run: (*parser).callonOrderedListItemPrefix114,
						expr: &seqExpr{
							pos: position{line: 321, col: 5, offset: 12893},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 321, col: 5, offset: 12893},
									expr: &ruleRefExpr{
										pos:  position{line: 321, col: 5, offset: 12893},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 321, col: 9, offset: 12897},
									label: "style",
									expr: &seqExpr{
										pos: position{line: 321, col: 16, offset: 12904},
										exprs: []interface{}{
											&oneOrMoreExpr{
												pos: position{line: 321, col: 16, offset: 12904},
												expr: &seqExpr{
													pos: position{line: 321, col: 17, offset: 12905},
													exprs: []interface{}{
														&notExpr{
															pos: position{line: 321, col: 17, offset: 12905},
															expr: &litMatcher{
																pos:        position{line: 321, col: 18, offset: 12906},
																val:        ")",
																ignoreCase: false,
															},
														},
														&notExpr{
															pos: position{line: 321, col: 22, offset: 12910},
															expr: &ruleRefExpr{
																pos:  position{line: 321, col: 23, offset: 12911},
																name: "WS",
															},
														},
														&notExpr{
															pos: position{line: 321, col: 26, offset: 12914},
															expr: &ruleRefExpr{
																pos:  position{line: 321, col: 27, offset: 12915},
																name: "NEWLINE",
															},
														},
														&charClassMatcher{
															pos:        position{line: 321, col: 35, offset: 12923},
															val:        "[A-Z]",
															ranges:     []rune{'A', 'Z'},
															ignoreCase: false,
															inverted:   false,
														},
													},
												},
											},
											&litMatcher{
												pos:        position{line: 321, col: 43, offset: 12931},
												val:        ")",
												ignoreCase: false,
											},
										},
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 321, col: 48, offset: 12936},
									expr: &ruleRefExpr{
										pos:  position{line: 321, col: 48, offset: 12936},
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
			pos:  position{line: 325, col: 1, offset: 13034},
			expr: &actionExpr{
				pos: position{line: 325, col: 27, offset: 13060},
				run: (*parser).callonOrderedListItemContent1,
				expr: &labeledExpr{
					pos:   position{line: 325, col: 27, offset: 13060},
					label: "elements",
					expr: &seqExpr{
						pos: position{line: 325, col: 37, offset: 13070},
						exprs: []interface{}{
							&oneOrMoreExpr{
								pos: position{line: 325, col: 37, offset: 13070},
								expr: &ruleRefExpr{
									pos:  position{line: 325, col: 37, offset: 13070},
									name: "ListParagraph",
								},
							},
							&zeroOrMoreExpr{
								pos: position{line: 325, col: 52, offset: 13085},
								expr: &ruleRefExpr{
									pos:  position{line: 325, col: 52, offset: 13085},
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
			pos:  position{line: 332, col: 1, offset: 13411},
			expr: &actionExpr{
				pos: position{line: 332, col: 22, offset: 13432},
				run: (*parser).callonUnorderedListItem1,
				expr: &seqExpr{
					pos: position{line: 332, col: 22, offset: 13432},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 332, col: 22, offset: 13432},
							label: "prefix",
							expr: &ruleRefExpr{
								pos:  position{line: 332, col: 30, offset: 13440},
								name: "UnorderedListItemPrefix",
							},
						},
						&labeledExpr{
							pos:   position{line: 332, col: 55, offset: 13465},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 332, col: 64, offset: 13474},
								name: "UnorderedListItemContent",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 332, col: 90, offset: 13500},
							expr: &ruleRefExpr{
								pos:  position{line: 332, col: 90, offset: 13500},
								name: "BlankLine",
							},
						},
					},
				},
			},
		},
		{
			name: "UnorderedListItemPrefix",
			pos:  position{line: 336, col: 1, offset: 13624},
			expr: &choiceExpr{
				pos: position{line: 336, col: 28, offset: 13651},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 336, col: 28, offset: 13651},
						run: (*parser).callonUnorderedListItemPrefix2,
						expr: &seqExpr{
							pos: position{line: 336, col: 28, offset: 13651},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 336, col: 28, offset: 13651},
									expr: &ruleRefExpr{
										pos:  position{line: 336, col: 28, offset: 13651},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 336, col: 32, offset: 13655},
									label: "level",
									expr: &litMatcher{
										pos:        position{line: 336, col: 39, offset: 13662},
										val:        "*****",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 336, col: 48, offset: 13671},
									expr: &ruleRefExpr{
										pos:  position{line: 336, col: 48, offset: 13671},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 338, col: 5, offset: 13816},
						run: (*parser).callonUnorderedListItemPrefix10,
						expr: &seqExpr{
							pos: position{line: 338, col: 5, offset: 13816},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 338, col: 5, offset: 13816},
									expr: &ruleRefExpr{
										pos:  position{line: 338, col: 5, offset: 13816},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 338, col: 9, offset: 13820},
									label: "level",
									expr: &litMatcher{
										pos:        position{line: 338, col: 16, offset: 13827},
										val:        "****",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 338, col: 24, offset: 13835},
									expr: &ruleRefExpr{
										pos:  position{line: 338, col: 24, offset: 13835},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 340, col: 5, offset: 13980},
						run: (*parser).callonUnorderedListItemPrefix18,
						expr: &seqExpr{
							pos: position{line: 340, col: 5, offset: 13980},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 340, col: 5, offset: 13980},
									expr: &ruleRefExpr{
										pos:  position{line: 340, col: 5, offset: 13980},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 340, col: 9, offset: 13984},
									label: "level",
									expr: &litMatcher{
										pos:        position{line: 340, col: 16, offset: 13991},
										val:        "***",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 340, col: 23, offset: 13998},
									expr: &ruleRefExpr{
										pos:  position{line: 340, col: 23, offset: 13998},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 342, col: 5, offset: 14144},
						run: (*parser).callonUnorderedListItemPrefix26,
						expr: &seqExpr{
							pos: position{line: 342, col: 5, offset: 14144},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 342, col: 5, offset: 14144},
									expr: &ruleRefExpr{
										pos:  position{line: 342, col: 5, offset: 14144},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 342, col: 9, offset: 14148},
									label: "level",
									expr: &litMatcher{
										pos:        position{line: 342, col: 16, offset: 14155},
										val:        "**",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 342, col: 22, offset: 14161},
									expr: &ruleRefExpr{
										pos:  position{line: 342, col: 22, offset: 14161},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 344, col: 5, offset: 14305},
						run: (*parser).callonUnorderedListItemPrefix34,
						expr: &seqExpr{
							pos: position{line: 344, col: 5, offset: 14305},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 344, col: 5, offset: 14305},
									expr: &ruleRefExpr{
										pos:  position{line: 344, col: 5, offset: 14305},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 344, col: 9, offset: 14309},
									label: "level",
									expr: &litMatcher{
										pos:        position{line: 344, col: 16, offset: 14316},
										val:        "*",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 344, col: 21, offset: 14321},
									expr: &ruleRefExpr{
										pos:  position{line: 344, col: 21, offset: 14321},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 346, col: 5, offset: 14464},
						run: (*parser).callonUnorderedListItemPrefix42,
						expr: &seqExpr{
							pos: position{line: 346, col: 5, offset: 14464},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 346, col: 5, offset: 14464},
									expr: &ruleRefExpr{
										pos:  position{line: 346, col: 5, offset: 14464},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 346, col: 9, offset: 14468},
									label: "level",
									expr: &litMatcher{
										pos:        position{line: 346, col: 16, offset: 14475},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 346, col: 21, offset: 14480},
									expr: &ruleRefExpr{
										pos:  position{line: 346, col: 21, offset: 14480},
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
			pos:  position{line: 350, col: 1, offset: 14616},
			expr: &actionExpr{
				pos: position{line: 350, col: 29, offset: 14644},
				run: (*parser).callonUnorderedListItemContent1,
				expr: &labeledExpr{
					pos:   position{line: 350, col: 29, offset: 14644},
					label: "elements",
					expr: &seqExpr{
						pos: position{line: 350, col: 39, offset: 14654},
						exprs: []interface{}{
							&oneOrMoreExpr{
								pos: position{line: 350, col: 39, offset: 14654},
								expr: &ruleRefExpr{
									pos:  position{line: 350, col: 39, offset: 14654},
									name: "ListParagraph",
								},
							},
							&zeroOrMoreExpr{
								pos: position{line: 350, col: 54, offset: 14669},
								expr: &ruleRefExpr{
									pos:  position{line: 350, col: 54, offset: 14669},
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
			pos:  position{line: 357, col: 1, offset: 14993},
			expr: &choiceExpr{
				pos: position{line: 357, col: 20, offset: 15012},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 357, col: 20, offset: 15012},
						run: (*parser).callonLabeledListItem2,
						expr: &seqExpr{
							pos: position{line: 357, col: 20, offset: 15012},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 357, col: 20, offset: 15012},
									label: "term",
									expr: &ruleRefExpr{
										pos:  position{line: 357, col: 26, offset: 15018},
										name: "LabeledListItemTerm",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 357, col: 47, offset: 15039},
									name: "LabeledListItemSeparator",
								},
								&labeledExpr{
									pos:   position{line: 357, col: 72, offset: 15064},
									label: "description",
									expr: &ruleRefExpr{
										pos:  position{line: 357, col: 85, offset: 15077},
										name: "LabeledListItemDescription",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 359, col: 6, offset: 15204},
						run: (*parser).callonLabeledListItem9,
						expr: &seqExpr{
							pos: position{line: 359, col: 6, offset: 15204},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 359, col: 6, offset: 15204},
									label: "term",
									expr: &ruleRefExpr{
										pos:  position{line: 359, col: 12, offset: 15210},
										name: "LabeledListItemTerm",
									},
								},
								&litMatcher{
									pos:        position{line: 359, col: 33, offset: 15231},
									val:        "::",
									ignoreCase: false,
								},
								&zeroOrMoreExpr{
									pos: position{line: 359, col: 38, offset: 15236},
									expr: &ruleRefExpr{
										pos:  position{line: 359, col: 38, offset: 15236},
										name: "WS",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 359, col: 42, offset: 15240},
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
			pos:  position{line: 363, col: 1, offset: 15377},
			expr: &actionExpr{
				pos: position{line: 363, col: 24, offset: 15400},
				run: (*parser).callonLabeledListItemTerm1,
				expr: &labeledExpr{
					pos:   position{line: 363, col: 24, offset: 15400},
					label: "term",
					expr: &zeroOrMoreExpr{
						pos: position{line: 363, col: 29, offset: 15405},
						expr: &seqExpr{
							pos: position{line: 363, col: 30, offset: 15406},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 363, col: 30, offset: 15406},
									expr: &ruleRefExpr{
										pos:  position{line: 363, col: 31, offset: 15407},
										name: "NEWLINE",
									},
								},
								&notExpr{
									pos: position{line: 363, col: 39, offset: 15415},
									expr: &litMatcher{
										pos:        position{line: 363, col: 40, offset: 15416},
										val:        "::",
										ignoreCase: false,
									},
								},
								&anyMatcher{
									line: 363, col: 45, offset: 15421,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "LabeledListItemSeparator",
			pos:  position{line: 368, col: 1, offset: 15512},
			expr: &seqExpr{
				pos: position{line: 368, col: 30, offset: 15541},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 368, col: 30, offset: 15541},
						val:        "::",
						ignoreCase: false,
					},
					&oneOrMoreExpr{
						pos: position{line: 368, col: 35, offset: 15546},
						expr: &choiceExpr{
							pos: position{line: 368, col: 36, offset: 15547},
							alternatives: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 368, col: 36, offset: 15547},
									name: "WS",
								},
								&ruleRefExpr{
									pos:  position{line: 368, col: 41, offset: 15552},
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
			pos:  position{line: 370, col: 1, offset: 15563},
			expr: &actionExpr{
				pos: position{line: 370, col: 31, offset: 15593},
				run: (*parser).callonLabeledListItemDescription1,
				expr: &labeledExpr{
					pos:   position{line: 370, col: 31, offset: 15593},
					label: "elements",
					expr: &zeroOrMoreExpr{
						pos: position{line: 370, col: 40, offset: 15602},
						expr: &choiceExpr{
							pos: position{line: 370, col: 41, offset: 15603},
							alternatives: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 370, col: 41, offset: 15603},
									name: "ListParagraph",
								},
								&ruleRefExpr{
									pos:  position{line: 370, col: 57, offset: 15619},
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
			pos:  position{line: 378, col: 1, offset: 15926},
			expr: &choiceExpr{
				pos: position{line: 378, col: 19, offset: 15944},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 378, col: 19, offset: 15944},
						run: (*parser).callonAdmonitionKind2,
						expr: &litMatcher{
							pos:        position{line: 378, col: 19, offset: 15944},
							val:        "TIP",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 380, col: 5, offset: 15982},
						run: (*parser).callonAdmonitionKind4,
						expr: &litMatcher{
							pos:        position{line: 380, col: 5, offset: 15982},
							val:        "NOTE",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 382, col: 5, offset: 16022},
						run: (*parser).callonAdmonitionKind6,
						expr: &litMatcher{
							pos:        position{line: 382, col: 5, offset: 16022},
							val:        "IMPORTANT",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 384, col: 5, offset: 16072},
						run: (*parser).callonAdmonitionKind8,
						expr: &litMatcher{
							pos:        position{line: 384, col: 5, offset: 16072},
							val:        "WARNING",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 386, col: 5, offset: 16118},
						run: (*parser).callonAdmonitionKind10,
						expr: &litMatcher{
							pos:        position{line: 386, col: 5, offset: 16118},
							val:        "CAUTION",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Paragraph",
			pos:  position{line: 395, col: 1, offset: 16421},
			expr: &choiceExpr{
				pos: position{line: 395, col: 14, offset: 16434},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 395, col: 14, offset: 16434},
						run: (*parser).callonParagraph2,
						expr: &seqExpr{
							pos: position{line: 395, col: 14, offset: 16434},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 395, col: 14, offset: 16434},
									label: "attributes",
									expr: &zeroOrMoreExpr{
										pos: position{line: 395, col: 25, offset: 16445},
										expr: &ruleRefExpr{
											pos:  position{line: 395, col: 26, offset: 16446},
											name: "ElementAttribute",
										},
									},
								},
								&notExpr{
									pos: position{line: 395, col: 45, offset: 16465},
									expr: &seqExpr{
										pos: position{line: 395, col: 47, offset: 16467},
										exprs: []interface{}{
											&oneOrMoreExpr{
												pos: position{line: 395, col: 47, offset: 16467},
												expr: &litMatcher{
													pos:        position{line: 395, col: 47, offset: 16467},
													val:        "=",
													ignoreCase: false,
												},
											},
											&oneOrMoreExpr{
												pos: position{line: 395, col: 52, offset: 16472},
												expr: &ruleRefExpr{
													pos:  position{line: 395, col: 52, offset: 16472},
													name: "WS",
												},
											},
										},
									},
								},
								&labeledExpr{
									pos:   position{line: 395, col: 57, offset: 16477},
									label: "t",
									expr: &ruleRefExpr{
										pos:  position{line: 395, col: 60, offset: 16480},
										name: "AdmonitionKind",
									},
								},
								&litMatcher{
									pos:        position{line: 395, col: 76, offset: 16496},
									val:        ": ",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 395, col: 81, offset: 16501},
									label: "lines",
									expr: &oneOrMoreExpr{
										pos: position{line: 395, col: 87, offset: 16507},
										expr: &seqExpr{
											pos: position{line: 395, col: 88, offset: 16508},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 395, col: 88, offset: 16508},
													name: "InlineContentWithTrailingSpaces",
												},
												&ruleRefExpr{
													pos:  position{line: 395, col: 120, offset: 16540},
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
						pos: position{line: 397, col: 5, offset: 16692},
						run: (*parser).callonParagraph21,
						expr: &seqExpr{
							pos: position{line: 397, col: 5, offset: 16692},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 397, col: 5, offset: 16692},
									label: "attributes",
									expr: &zeroOrMoreExpr{
										pos: position{line: 397, col: 16, offset: 16703},
										expr: &ruleRefExpr{
											pos:  position{line: 397, col: 17, offset: 16704},
											name: "ElementAttribute",
										},
									},
								},
								&notExpr{
									pos: position{line: 397, col: 36, offset: 16723},
									expr: &seqExpr{
										pos: position{line: 397, col: 38, offset: 16725},
										exprs: []interface{}{
											&oneOrMoreExpr{
												pos: position{line: 397, col: 38, offset: 16725},
												expr: &litMatcher{
													pos:        position{line: 397, col: 38, offset: 16725},
													val:        "=",
													ignoreCase: false,
												},
											},
											&oneOrMoreExpr{
												pos: position{line: 397, col: 43, offset: 16730},
												expr: &ruleRefExpr{
													pos:  position{line: 397, col: 43, offset: 16730},
													name: "WS",
												},
											},
										},
									},
								},
								&labeledExpr{
									pos:   position{line: 397, col: 48, offset: 16735},
									label: "lines",
									expr: &oneOrMoreExpr{
										pos: position{line: 397, col: 54, offset: 16741},
										expr: &seqExpr{
											pos: position{line: 397, col: 55, offset: 16742},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 397, col: 55, offset: 16742},
													name: "InlineContentWithTrailingSpaces",
												},
												&ruleRefExpr{
													pos:  position{line: 397, col: 87, offset: 16774},
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
			name: "InlineContentWithTrailingSpaces",
			pos:  position{line: 403, col: 1, offset: 17085},
			expr: &actionExpr{
				pos: position{line: 403, col: 36, offset: 17120},
				run: (*parser).callonInlineContentWithTrailingSpaces1,
				expr: &seqExpr{
					pos: position{line: 403, col: 36, offset: 17120},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 403, col: 36, offset: 17120},
							expr: &ruleRefExpr{
								pos:  position{line: 403, col: 37, offset: 17121},
								name: "BlockDelimiter",
							},
						},
						&labeledExpr{
							pos:   position{line: 403, col: 52, offset: 17136},
							label: "elements",
							expr: &oneOrMoreExpr{
								pos: position{line: 403, col: 61, offset: 17145},
								expr: &seqExpr{
									pos: position{line: 403, col: 62, offset: 17146},
									exprs: []interface{}{
										&zeroOrMoreExpr{
											pos: position{line: 403, col: 62, offset: 17146},
											expr: &ruleRefExpr{
												pos:  position{line: 403, col: 62, offset: 17146},
												name: "WS",
											},
										},
										&notExpr{
											pos: position{line: 403, col: 66, offset: 17150},
											expr: &ruleRefExpr{
												pos:  position{line: 403, col: 67, offset: 17151},
												name: "InlineElementID",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 403, col: 83, offset: 17167},
											name: "InlineElement",
										},
										&zeroOrMoreExpr{
											pos: position{line: 403, col: 97, offset: 17181},
											expr: &ruleRefExpr{
												pos:  position{line: 403, col: 97, offset: 17181},
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
			name: "InlineContent",
			pos:  position{line: 407, col: 1, offset: 17314},
			expr: &actionExpr{
				pos: position{line: 407, col: 18, offset: 17331},
				run: (*parser).callonInlineContent1,
				expr: &seqExpr{
					pos: position{line: 407, col: 18, offset: 17331},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 407, col: 18, offset: 17331},
							expr: &ruleRefExpr{
								pos:  position{line: 407, col: 19, offset: 17332},
								name: "BlockDelimiter",
							},
						},
						&labeledExpr{
							pos:   position{line: 407, col: 34, offset: 17347},
							label: "elements",
							expr: &oneOrMoreExpr{
								pos: position{line: 407, col: 43, offset: 17356},
								expr: &seqExpr{
									pos: position{line: 407, col: 44, offset: 17357},
									exprs: []interface{}{
										&zeroOrMoreExpr{
											pos: position{line: 407, col: 44, offset: 17357},
											expr: &ruleRefExpr{
												pos:  position{line: 407, col: 44, offset: 17357},
												name: "WS",
											},
										},
										&notExpr{
											pos: position{line: 407, col: 48, offset: 17361},
											expr: &ruleRefExpr{
												pos:  position{line: 407, col: 49, offset: 17362},
												name: "InlineElementID",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 407, col: 65, offset: 17378},
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
			pos:  position{line: 411, col: 1, offset: 17500},
			expr: &choiceExpr{
				pos: position{line: 411, col: 18, offset: 17517},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 411, col: 18, offset: 17517},
						name: "CrossReference",
					},
					&ruleRefExpr{
						pos:  position{line: 411, col: 35, offset: 17534},
						name: "Passthrough",
					},
					&ruleRefExpr{
						pos:  position{line: 411, col: 49, offset: 17548},
						name: "InlineImage",
					},
					&ruleRefExpr{
						pos:  position{line: 411, col: 63, offset: 17562},
						name: "QuotedText",
					},
					&ruleRefExpr{
						pos:  position{line: 411, col: 76, offset: 17575},
						name: "Link",
					},
					&ruleRefExpr{
						pos:  position{line: 411, col: 83, offset: 17582},
						name: "DocumentAttributeSubstitution",
					},
					&ruleRefExpr{
						pos:  position{line: 411, col: 115, offset: 17614},
						name: "Characters",
					},
				},
			},
		},
		{
			name: "QuotedText",
			pos:  position{line: 416, col: 1, offset: 17865},
			expr: &choiceExpr{
				pos: position{line: 416, col: 15, offset: 17879},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 416, col: 15, offset: 17879},
						name: "BoldText",
					},
					&ruleRefExpr{
						pos:  position{line: 416, col: 26, offset: 17890},
						name: "ItalicText",
					},
					&ruleRefExpr{
						pos:  position{line: 416, col: 39, offset: 17903},
						name: "MonospaceText",
					},
					&ruleRefExpr{
						pos:  position{line: 417, col: 13, offset: 17931},
						name: "EscapedBoldText",
					},
					&ruleRefExpr{
						pos:  position{line: 417, col: 31, offset: 17949},
						name: "EscapedItalicText",
					},
					&ruleRefExpr{
						pos:  position{line: 417, col: 51, offset: 17969},
						name: "EscapedMonospaceText",
					},
				},
			},
		},
		{
			name: "BoldText",
			pos:  position{line: 419, col: 1, offset: 17991},
			expr: &choiceExpr{
				pos: position{line: 419, col: 13, offset: 18003},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 419, col: 13, offset: 18003},
						run: (*parser).callonBoldText2,
						expr: &seqExpr{
							pos: position{line: 419, col: 13, offset: 18003},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 419, col: 13, offset: 18003},
									expr: &litMatcher{
										pos:        position{line: 419, col: 14, offset: 18004},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 419, col: 19, offset: 18009},
									val:        "**",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 419, col: 24, offset: 18014},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 419, col: 33, offset: 18023},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 419, col: 52, offset: 18042},
									val:        "**",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 421, col: 5, offset: 18167},
						run: (*parser).callonBoldText10,
						expr: &seqExpr{
							pos: position{line: 421, col: 5, offset: 18167},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 421, col: 5, offset: 18167},
									expr: &litMatcher{
										pos:        position{line: 421, col: 6, offset: 18168},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 421, col: 11, offset: 18173},
									val:        "**",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 421, col: 16, offset: 18178},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 421, col: 25, offset: 18187},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 421, col: 44, offset: 18206},
									val:        "*",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 424, col: 5, offset: 18371},
						run: (*parser).callonBoldText18,
						expr: &seqExpr{
							pos: position{line: 424, col: 5, offset: 18371},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 424, col: 5, offset: 18371},
									expr: &litMatcher{
										pos:        position{line: 424, col: 6, offset: 18372},
										val:        "\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 424, col: 10, offset: 18376},
									val:        "*",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 424, col: 14, offset: 18380},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 424, col: 23, offset: 18389},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 424, col: 42, offset: 18408},
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
			pos:  position{line: 428, col: 1, offset: 18508},
			expr: &choiceExpr{
				pos: position{line: 428, col: 20, offset: 18527},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 428, col: 20, offset: 18527},
						run: (*parser).callonEscapedBoldText2,
						expr: &seqExpr{
							pos: position{line: 428, col: 20, offset: 18527},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 428, col: 20, offset: 18527},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 428, col: 33, offset: 18540},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 428, col: 33, offset: 18540},
												val:        "\\\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 428, col: 38, offset: 18545},
												expr: &litMatcher{
													pos:        position{line: 428, col: 38, offset: 18545},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 428, col: 44, offset: 18551},
									val:        "**",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 428, col: 49, offset: 18556},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 428, col: 58, offset: 18565},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 428, col: 77, offset: 18584},
									val:        "**",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 430, col: 5, offset: 18739},
						run: (*parser).callonEscapedBoldText13,
						expr: &seqExpr{
							pos: position{line: 430, col: 5, offset: 18739},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 430, col: 5, offset: 18739},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 430, col: 18, offset: 18752},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 430, col: 18, offset: 18752},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 430, col: 22, offset: 18756},
												expr: &litMatcher{
													pos:        position{line: 430, col: 22, offset: 18756},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 430, col: 28, offset: 18762},
									val:        "**",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 430, col: 33, offset: 18767},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 430, col: 42, offset: 18776},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 430, col: 61, offset: 18795},
									val:        "*",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 433, col: 5, offset: 18989},
						run: (*parser).callonEscapedBoldText24,
						expr: &seqExpr{
							pos: position{line: 433, col: 5, offset: 18989},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 433, col: 5, offset: 18989},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 433, col: 18, offset: 19002},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 433, col: 18, offset: 19002},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 433, col: 22, offset: 19006},
												expr: &litMatcher{
													pos:        position{line: 433, col: 22, offset: 19006},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 433, col: 28, offset: 19012},
									val:        "*",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 433, col: 32, offset: 19016},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 433, col: 41, offset: 19025},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 433, col: 60, offset: 19044},
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
			pos:  position{line: 437, col: 1, offset: 19196},
			expr: &choiceExpr{
				pos: position{line: 437, col: 15, offset: 19210},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 437, col: 15, offset: 19210},
						run: (*parser).callonItalicText2,
						expr: &seqExpr{
							pos: position{line: 437, col: 15, offset: 19210},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 437, col: 15, offset: 19210},
									expr: &litMatcher{
										pos:        position{line: 437, col: 16, offset: 19211},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 437, col: 21, offset: 19216},
									val:        "__",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 437, col: 26, offset: 19221},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 437, col: 35, offset: 19230},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 437, col: 54, offset: 19249},
									val:        "__",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 439, col: 5, offset: 19330},
						run: (*parser).callonItalicText10,
						expr: &seqExpr{
							pos: position{line: 439, col: 5, offset: 19330},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 439, col: 5, offset: 19330},
									expr: &litMatcher{
										pos:        position{line: 439, col: 6, offset: 19331},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 439, col: 11, offset: 19336},
									val:        "__",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 439, col: 16, offset: 19341},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 439, col: 25, offset: 19350},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 439, col: 44, offset: 19369},
									val:        "_",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 442, col: 5, offset: 19536},
						run: (*parser).callonItalicText18,
						expr: &seqExpr{
							pos: position{line: 442, col: 5, offset: 19536},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 442, col: 5, offset: 19536},
									expr: &litMatcher{
										pos:        position{line: 442, col: 6, offset: 19537},
										val:        "\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 442, col: 10, offset: 19541},
									val:        "_",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 442, col: 14, offset: 19545},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 442, col: 23, offset: 19554},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 442, col: 42, offset: 19573},
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
			pos:  position{line: 446, col: 1, offset: 19652},
			expr: &choiceExpr{
				pos: position{line: 446, col: 22, offset: 19673},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 446, col: 22, offset: 19673},
						run: (*parser).callonEscapedItalicText2,
						expr: &seqExpr{
							pos: position{line: 446, col: 22, offset: 19673},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 446, col: 22, offset: 19673},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 446, col: 35, offset: 19686},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 446, col: 35, offset: 19686},
												val:        "\\\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 446, col: 40, offset: 19691},
												expr: &litMatcher{
													pos:        position{line: 446, col: 40, offset: 19691},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 446, col: 46, offset: 19697},
									val:        "__",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 446, col: 51, offset: 19702},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 446, col: 60, offset: 19711},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 446, col: 79, offset: 19730},
									val:        "__",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 448, col: 5, offset: 19885},
						run: (*parser).callonEscapedItalicText13,
						expr: &seqExpr{
							pos: position{line: 448, col: 5, offset: 19885},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 448, col: 5, offset: 19885},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 448, col: 18, offset: 19898},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 448, col: 18, offset: 19898},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 448, col: 22, offset: 19902},
												expr: &litMatcher{
													pos:        position{line: 448, col: 22, offset: 19902},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 448, col: 28, offset: 19908},
									val:        "__",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 448, col: 33, offset: 19913},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 448, col: 42, offset: 19922},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 448, col: 61, offset: 19941},
									val:        "_",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 451, col: 5, offset: 20135},
						run: (*parser).callonEscapedItalicText24,
						expr: &seqExpr{
							pos: position{line: 451, col: 5, offset: 20135},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 451, col: 5, offset: 20135},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 451, col: 18, offset: 20148},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 451, col: 18, offset: 20148},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 451, col: 22, offset: 20152},
												expr: &litMatcher{
													pos:        position{line: 451, col: 22, offset: 20152},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 451, col: 28, offset: 20158},
									val:        "_",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 451, col: 32, offset: 20162},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 451, col: 41, offset: 20171},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 451, col: 60, offset: 20190},
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
			pos:  position{line: 455, col: 1, offset: 20342},
			expr: &choiceExpr{
				pos: position{line: 455, col: 18, offset: 20359},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 455, col: 18, offset: 20359},
						run: (*parser).callonMonospaceText2,
						expr: &seqExpr{
							pos: position{line: 455, col: 18, offset: 20359},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 455, col: 18, offset: 20359},
									expr: &litMatcher{
										pos:        position{line: 455, col: 19, offset: 20360},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 455, col: 24, offset: 20365},
									val:        "``",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 455, col: 29, offset: 20370},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 455, col: 38, offset: 20379},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 455, col: 57, offset: 20398},
									val:        "``",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 457, col: 5, offset: 20528},
						run: (*parser).callonMonospaceText10,
						expr: &seqExpr{
							pos: position{line: 457, col: 5, offset: 20528},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 457, col: 5, offset: 20528},
									expr: &litMatcher{
										pos:        position{line: 457, col: 6, offset: 20529},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 457, col: 11, offset: 20534},
									val:        "``",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 457, col: 16, offset: 20539},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 457, col: 25, offset: 20548},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 457, col: 44, offset: 20567},
									val:        "`",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 460, col: 5, offset: 20737},
						run: (*parser).callonMonospaceText18,
						expr: &seqExpr{
							pos: position{line: 460, col: 5, offset: 20737},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 460, col: 5, offset: 20737},
									expr: &litMatcher{
										pos:        position{line: 460, col: 6, offset: 20738},
										val:        "\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 460, col: 10, offset: 20742},
									val:        "`",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 460, col: 14, offset: 20746},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 460, col: 23, offset: 20755},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 460, col: 42, offset: 20774},
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
			pos:  position{line: 464, col: 1, offset: 20901},
			expr: &choiceExpr{
				pos: position{line: 464, col: 25, offset: 20925},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 464, col: 25, offset: 20925},
						run: (*parser).callonEscapedMonospaceText2,
						expr: &seqExpr{
							pos: position{line: 464, col: 25, offset: 20925},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 464, col: 25, offset: 20925},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 464, col: 38, offset: 20938},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 464, col: 38, offset: 20938},
												val:        "\\\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 464, col: 43, offset: 20943},
												expr: &litMatcher{
													pos:        position{line: 464, col: 43, offset: 20943},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 464, col: 49, offset: 20949},
									val:        "``",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 464, col: 54, offset: 20954},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 464, col: 63, offset: 20963},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 464, col: 82, offset: 20982},
									val:        "``",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 466, col: 5, offset: 21137},
						run: (*parser).callonEscapedMonospaceText13,
						expr: &seqExpr{
							pos: position{line: 466, col: 5, offset: 21137},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 466, col: 5, offset: 21137},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 466, col: 18, offset: 21150},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 466, col: 18, offset: 21150},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 466, col: 22, offset: 21154},
												expr: &litMatcher{
													pos:        position{line: 466, col: 22, offset: 21154},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 466, col: 28, offset: 21160},
									val:        "``",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 466, col: 33, offset: 21165},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 466, col: 42, offset: 21174},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 466, col: 61, offset: 21193},
									val:        "`",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 469, col: 5, offset: 21387},
						run: (*parser).callonEscapedMonospaceText24,
						expr: &seqExpr{
							pos: position{line: 469, col: 5, offset: 21387},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 469, col: 5, offset: 21387},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 469, col: 18, offset: 21400},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 469, col: 18, offset: 21400},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 469, col: 22, offset: 21404},
												expr: &litMatcher{
													pos:        position{line: 469, col: 22, offset: 21404},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 469, col: 28, offset: 21410},
									val:        "`",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 469, col: 32, offset: 21414},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 469, col: 41, offset: 21423},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 469, col: 60, offset: 21442},
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
			pos:  position{line: 473, col: 1, offset: 21594},
			expr: &seqExpr{
				pos: position{line: 473, col: 22, offset: 21615},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 473, col: 22, offset: 21615},
						name: "QuotedTextContentElement",
					},
					&zeroOrMoreExpr{
						pos: position{line: 473, col: 47, offset: 21640},
						expr: &seqExpr{
							pos: position{line: 473, col: 48, offset: 21641},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 473, col: 48, offset: 21641},
									expr: &ruleRefExpr{
										pos:  position{line: 473, col: 48, offset: 21641},
										name: "WS",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 473, col: 52, offset: 21645},
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
			pos:  position{line: 475, col: 1, offset: 21673},
			expr: &choiceExpr{
				pos: position{line: 475, col: 29, offset: 21701},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 475, col: 29, offset: 21701},
						name: "QuotedText",
					},
					&ruleRefExpr{
						pos:  position{line: 475, col: 42, offset: 21714},
						name: "QuotedTextCharacters",
					},
					&ruleRefExpr{
						pos:  position{line: 475, col: 65, offset: 21737},
						name: "CharactersWithQuotePunctuation",
					},
				},
			},
		},
		{
			name: "QuotedTextCharacters",
			pos:  position{line: 477, col: 1, offset: 21872},
			expr: &oneOrMoreExpr{
				pos: position{line: 477, col: 25, offset: 21896},
				expr: &seqExpr{
					pos: position{line: 477, col: 26, offset: 21897},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 477, col: 26, offset: 21897},
							expr: &ruleRefExpr{
								pos:  position{line: 477, col: 27, offset: 21898},
								name: "NEWLINE",
							},
						},
						&notExpr{
							pos: position{line: 477, col: 35, offset: 21906},
							expr: &ruleRefExpr{
								pos:  position{line: 477, col: 36, offset: 21907},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 477, col: 39, offset: 21910},
							expr: &litMatcher{
								pos:        position{line: 477, col: 40, offset: 21911},
								val:        "*",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 477, col: 44, offset: 21915},
							expr: &litMatcher{
								pos:        position{line: 477, col: 45, offset: 21916},
								val:        "_",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 477, col: 49, offset: 21920},
							expr: &litMatcher{
								pos:        position{line: 477, col: 50, offset: 21921},
								val:        "`",
								ignoreCase: false,
							},
						},
						&anyMatcher{
							line: 477, col: 54, offset: 21925,
						},
					},
				},
			},
		},
		{
			name: "CharactersWithQuotePunctuation",
			pos:  position{line: 479, col: 1, offset: 21968},
			expr: &actionExpr{
				pos: position{line: 479, col: 35, offset: 22002},
				run: (*parser).callonCharactersWithQuotePunctuation1,
				expr: &oneOrMoreExpr{
					pos: position{line: 479, col: 35, offset: 22002},
					expr: &seqExpr{
						pos: position{line: 479, col: 36, offset: 22003},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 479, col: 36, offset: 22003},
								expr: &ruleRefExpr{
									pos:  position{line: 479, col: 37, offset: 22004},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 479, col: 45, offset: 22012},
								expr: &ruleRefExpr{
									pos:  position{line: 479, col: 46, offset: 22013},
									name: "WS",
								},
							},
							&anyMatcher{
								line: 479, col: 50, offset: 22017,
							},
						},
					},
				},
			},
		},
		{
			name: "UnbalancedQuotePunctuation",
			pos:  position{line: 484, col: 1, offset: 22262},
			expr: &choiceExpr{
				pos: position{line: 484, col: 31, offset: 22292},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 484, col: 31, offset: 22292},
						val:        "*",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 484, col: 37, offset: 22298},
						val:        "_",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 484, col: 43, offset: 22304},
						val:        "`",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "Passthrough",
			pos:  position{line: 489, col: 1, offset: 22416},
			expr: &choiceExpr{
				pos: position{line: 489, col: 16, offset: 22431},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 489, col: 16, offset: 22431},
						name: "TriplePlusPassthrough",
					},
					&ruleRefExpr{
						pos:  position{line: 489, col: 40, offset: 22455},
						name: "SinglePlusPassthrough",
					},
					&ruleRefExpr{
						pos:  position{line: 489, col: 64, offset: 22479},
						name: "PassthroughMacro",
					},
				},
			},
		},
		{
			name: "SinglePlusPassthrough",
			pos:  position{line: 491, col: 1, offset: 22497},
			expr: &actionExpr{
				pos: position{line: 491, col: 26, offset: 22522},
				run: (*parser).callonSinglePlusPassthrough1,
				expr: &seqExpr{
					pos: position{line: 491, col: 26, offset: 22522},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 491, col: 26, offset: 22522},
							val:        "+",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 491, col: 30, offset: 22526},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 491, col: 38, offset: 22534},
								expr: &seqExpr{
									pos: position{line: 491, col: 39, offset: 22535},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 491, col: 39, offset: 22535},
											expr: &ruleRefExpr{
												pos:  position{line: 491, col: 40, offset: 22536},
												name: "NEWLINE",
											},
										},
										&notExpr{
											pos: position{line: 491, col: 48, offset: 22544},
											expr: &litMatcher{
												pos:        position{line: 491, col: 49, offset: 22545},
												val:        "+",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 491, col: 53, offset: 22549,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 491, col: 57, offset: 22553},
							val:        "+",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TriplePlusPassthrough",
			pos:  position{line: 495, col: 1, offset: 22648},
			expr: &actionExpr{
				pos: position{line: 495, col: 26, offset: 22673},
				run: (*parser).callonTriplePlusPassthrough1,
				expr: &seqExpr{
					pos: position{line: 495, col: 26, offset: 22673},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 495, col: 26, offset: 22673},
							val:        "+++",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 495, col: 32, offset: 22679},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 495, col: 40, offset: 22687},
								expr: &seqExpr{
									pos: position{line: 495, col: 41, offset: 22688},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 495, col: 41, offset: 22688},
											expr: &litMatcher{
												pos:        position{line: 495, col: 42, offset: 22689},
												val:        "+++",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 495, col: 48, offset: 22695,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 495, col: 52, offset: 22699},
							val:        "+++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "PassthroughMacro",
			pos:  position{line: 499, col: 1, offset: 22796},
			expr: &choiceExpr{
				pos: position{line: 499, col: 21, offset: 22816},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 499, col: 21, offset: 22816},
						run: (*parser).callonPassthroughMacro2,
						expr: &seqExpr{
							pos: position{line: 499, col: 21, offset: 22816},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 499, col: 21, offset: 22816},
									val:        "pass:[",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 499, col: 30, offset: 22825},
									label: "content",
									expr: &zeroOrMoreExpr{
										pos: position{line: 499, col: 38, offset: 22833},
										expr: &ruleRefExpr{
											pos:  position{line: 499, col: 39, offset: 22834},
											name: "PassthroughMacroCharacter",
										},
									},
								},
								&litMatcher{
									pos:        position{line: 499, col: 67, offset: 22862},
									val:        "]",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 501, col: 5, offset: 22953},
						run: (*parser).callonPassthroughMacro9,
						expr: &seqExpr{
							pos: position{line: 501, col: 5, offset: 22953},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 501, col: 5, offset: 22953},
									val:        "pass:q[",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 501, col: 15, offset: 22963},
									label: "content",
									expr: &zeroOrMoreExpr{
										pos: position{line: 501, col: 23, offset: 22971},
										expr: &choiceExpr{
											pos: position{line: 501, col: 24, offset: 22972},
											alternatives: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 501, col: 24, offset: 22972},
													name: "QuotedText",
												},
												&ruleRefExpr{
													pos:  position{line: 501, col: 37, offset: 22985},
													name: "PassthroughMacroCharacter",
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 501, col: 65, offset: 23013},
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
			pos:  position{line: 505, col: 1, offset: 23103},
			expr: &seqExpr{
				pos: position{line: 505, col: 31, offset: 23133},
				exprs: []interface{}{
					&notExpr{
						pos: position{line: 505, col: 31, offset: 23133},
						expr: &litMatcher{
							pos:        position{line: 505, col: 32, offset: 23134},
							val:        "]",
							ignoreCase: false,
						},
					},
					&anyMatcher{
						line: 505, col: 36, offset: 23138,
					},
				},
			},
		},
		{
			name: "CrossReference",
			pos:  position{line: 510, col: 1, offset: 23254},
			expr: &actionExpr{
				pos: position{line: 510, col: 19, offset: 23272},
				run: (*parser).callonCrossReference1,
				expr: &seqExpr{
					pos: position{line: 510, col: 19, offset: 23272},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 510, col: 19, offset: 23272},
							val:        "<<",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 510, col: 24, offset: 23277},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 510, col: 28, offset: 23281},
								name: "ID",
							},
						},
						&litMatcher{
							pos:        position{line: 510, col: 32, offset: 23285},
							val:        ">>",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Link",
			pos:  position{line: 517, col: 1, offset: 23444},
			expr: &choiceExpr{
				pos: position{line: 517, col: 9, offset: 23452},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 517, col: 9, offset: 23452},
						name: "RelativeLink",
					},
					&ruleRefExpr{
						pos:  position{line: 517, col: 24, offset: 23467},
						name: "ExternalLink",
					},
				},
			},
		},
		{
			name: "ExternalLink",
			pos:  position{line: 519, col: 1, offset: 23482},
			expr: &actionExpr{
				pos: position{line: 519, col: 17, offset: 23498},
				run: (*parser).callonExternalLink1,
				expr: &seqExpr{
					pos: position{line: 519, col: 17, offset: 23498},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 519, col: 17, offset: 23498},
							label: "url",
							expr: &seqExpr{
								pos: position{line: 519, col: 22, offset: 23503},
								exprs: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 519, col: 22, offset: 23503},
										name: "URL_SCHEME",
									},
									&ruleRefExpr{
										pos:  position{line: 519, col: 33, offset: 23514},
										name: "URL",
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 519, col: 38, offset: 23519},
							label: "text",
							expr: &zeroOrOneExpr{
								pos: position{line: 519, col: 43, offset: 23524},
								expr: &seqExpr{
									pos: position{line: 519, col: 44, offset: 23525},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 519, col: 44, offset: 23525},
											val:        "[",
											ignoreCase: false,
										},
										&zeroOrMoreExpr{
											pos: position{line: 519, col: 48, offset: 23529},
											expr: &ruleRefExpr{
												pos:  position{line: 519, col: 49, offset: 23530},
												name: "URL_TEXT",
											},
										},
										&litMatcher{
											pos:        position{line: 519, col: 60, offset: 23541},
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
			pos:  position{line: 526, col: 1, offset: 23702},
			expr: &actionExpr{
				pos: position{line: 526, col: 17, offset: 23718},
				run: (*parser).callonRelativeLink1,
				expr: &seqExpr{
					pos: position{line: 526, col: 17, offset: 23718},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 526, col: 17, offset: 23718},
							val:        "link:",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 526, col: 25, offset: 23726},
							label: "url",
							expr: &seqExpr{
								pos: position{line: 526, col: 30, offset: 23731},
								exprs: []interface{}{
									&zeroOrOneExpr{
										pos: position{line: 526, col: 30, offset: 23731},
										expr: &ruleRefExpr{
											pos:  position{line: 526, col: 30, offset: 23731},
											name: "URL_SCHEME",
										},
									},
									&ruleRefExpr{
										pos:  position{line: 526, col: 42, offset: 23743},
										name: "URL",
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 526, col: 47, offset: 23748},
							label: "text",
							expr: &seqExpr{
								pos: position{line: 526, col: 53, offset: 23754},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 526, col: 53, offset: 23754},
										val:        "[",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 526, col: 57, offset: 23758},
										expr: &ruleRefExpr{
											pos:  position{line: 526, col: 58, offset: 23759},
											name: "URL_TEXT",
										},
									},
									&litMatcher{
										pos:        position{line: 526, col: 69, offset: 23770},
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
			pos:  position{line: 536, col: 1, offset: 24032},
			expr: &actionExpr{
				pos: position{line: 536, col: 15, offset: 24046},
				run: (*parser).callonBlockImage1,
				expr: &seqExpr{
					pos: position{line: 536, col: 15, offset: 24046},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 536, col: 15, offset: 24046},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 536, col: 26, offset: 24057},
								expr: &ruleRefExpr{
									pos:  position{line: 536, col: 27, offset: 24058},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 536, col: 46, offset: 24077},
							label: "image",
							expr: &ruleRefExpr{
								pos:  position{line: 536, col: 52, offset: 24083},
								name: "BlockImageMacro",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 536, col: 69, offset: 24100},
							expr: &ruleRefExpr{
								pos:  position{line: 536, col: 69, offset: 24100},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 536, col: 73, offset: 24104},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "BlockImageMacro",
			pos:  position{line: 541, col: 1, offset: 24263},
			expr: &actionExpr{
				pos: position{line: 541, col: 20, offset: 24282},
				run: (*parser).callonBlockImageMacro1,
				expr: &seqExpr{
					pos: position{line: 541, col: 20, offset: 24282},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 541, col: 20, offset: 24282},
							val:        "image::",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 541, col: 30, offset: 24292},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 541, col: 36, offset: 24298},
								name: "URL",
							},
						},
						&litMatcher{
							pos:        position{line: 541, col: 41, offset: 24303},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 541, col: 45, offset: 24307},
							label: "attributes",
							expr: &zeroOrOneExpr{
								pos: position{line: 541, col: 57, offset: 24319},
								expr: &ruleRefExpr{
									pos:  position{line: 541, col: 57, offset: 24319},
									name: "URL_TEXT",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 541, col: 68, offset: 24330},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "InlineImage",
			pos:  position{line: 545, col: 1, offset: 24397},
			expr: &actionExpr{
				pos: position{line: 545, col: 16, offset: 24412},
				run: (*parser).callonInlineImage1,
				expr: &labeledExpr{
					pos:   position{line: 545, col: 16, offset: 24412},
					label: "image",
					expr: &ruleRefExpr{
						pos:  position{line: 545, col: 22, offset: 24418},
						name: "InlineImageMacro",
					},
				},
			},
		},
		{
			name: "InlineImageMacro",
			pos:  position{line: 550, col: 1, offset: 24563},
			expr: &actionExpr{
				pos: position{line: 550, col: 21, offset: 24583},
				run: (*parser).callonInlineImageMacro1,
				expr: &seqExpr{
					pos: position{line: 550, col: 21, offset: 24583},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 550, col: 21, offset: 24583},
							val:        "image:",
							ignoreCase: false,
						},
						&notExpr{
							pos: position{line: 550, col: 30, offset: 24592},
							expr: &litMatcher{
								pos:        position{line: 550, col: 31, offset: 24593},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 550, col: 35, offset: 24597},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 550, col: 41, offset: 24603},
								name: "URL",
							},
						},
						&litMatcher{
							pos:        position{line: 550, col: 46, offset: 24608},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 550, col: 50, offset: 24612},
							label: "attributes",
							expr: &zeroOrOneExpr{
								pos: position{line: 550, col: 62, offset: 24624},
								expr: &ruleRefExpr{
									pos:  position{line: 550, col: 62, offset: 24624},
									name: "URL_TEXT",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 550, col: 73, offset: 24635},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "DelimitedBlock",
			pos:  position{line: 557, col: 1, offset: 24965},
			expr: &choiceExpr{
				pos: position{line: 557, col: 19, offset: 24983},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 557, col: 19, offset: 24983},
						name: "FencedBlock",
					},
					&ruleRefExpr{
						pos:  position{line: 557, col: 33, offset: 24997},
						name: "ListingBlock",
					},
					&ruleRefExpr{
						pos:  position{line: 557, col: 48, offset: 25012},
						name: "ExampleBlock",
					},
				},
			},
		},
		{
			name: "BlockDelimiter",
			pos:  position{line: 559, col: 1, offset: 25026},
			expr: &choiceExpr{
				pos: position{line: 559, col: 19, offset: 25044},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 559, col: 19, offset: 25044},
						name: "LiteralBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 559, col: 43, offset: 25068},
						name: "FencedBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 559, col: 66, offset: 25091},
						name: "ListingBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 559, col: 90, offset: 25115},
						name: "ExampleBlockDelimiter",
					},
				},
			},
		},
		{
			name: "FencedBlockDelimiter",
			pos:  position{line: 561, col: 1, offset: 25138},
			expr: &litMatcher{
				pos:        position{line: 561, col: 25, offset: 25162},
				val:        "```",
				ignoreCase: false,
			},
		},
		{
			name: "FencedBlock",
			pos:  position{line: 563, col: 1, offset: 25169},
			expr: &actionExpr{
				pos: position{line: 563, col: 16, offset: 25184},
				run: (*parser).callonFencedBlock1,
				expr: &seqExpr{
					pos: position{line: 563, col: 16, offset: 25184},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 563, col: 16, offset: 25184},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 563, col: 27, offset: 25195},
								expr: &ruleRefExpr{
									pos:  position{line: 563, col: 28, offset: 25196},
									name: "ElementAttribute",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 563, col: 47, offset: 25215},
							name: "FencedBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 563, col: 68, offset: 25236},
							expr: &ruleRefExpr{
								pos:  position{line: 563, col: 68, offset: 25236},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 563, col: 72, offset: 25240},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 563, col: 80, offset: 25248},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 563, col: 88, offset: 25256},
								expr: &choiceExpr{
									pos: position{line: 563, col: 89, offset: 25257},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 563, col: 89, offset: 25257},
											name: "List",
										},
										&ruleRefExpr{
											pos:  position{line: 563, col: 96, offset: 25264},
											name: "Paragraph",
										},
										&ruleRefExpr{
											pos:  position{line: 563, col: 108, offset: 25276},
											name: "BlankLine",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 563, col: 120, offset: 25288},
							name: "FencedBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 563, col: 141, offset: 25309},
							expr: &ruleRefExpr{
								pos:  position{line: 563, col: 141, offset: 25309},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 563, col: 145, offset: 25313},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ListingBlockDelimiter",
			pos:  position{line: 567, col: 1, offset: 25429},
			expr: &litMatcher{
				pos:        position{line: 567, col: 26, offset: 25454},
				val:        "----",
				ignoreCase: false,
			},
		},
		{
			name: "ListingBlock",
			pos:  position{line: 569, col: 1, offset: 25462},
			expr: &actionExpr{
				pos: position{line: 569, col: 17, offset: 25478},
				run: (*parser).callonListingBlock1,
				expr: &seqExpr{
					pos: position{line: 569, col: 17, offset: 25478},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 569, col: 17, offset: 25478},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 569, col: 28, offset: 25489},
								expr: &ruleRefExpr{
									pos:  position{line: 569, col: 29, offset: 25490},
									name: "ElementAttribute",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 569, col: 48, offset: 25509},
							name: "ListingBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 569, col: 70, offset: 25531},
							expr: &ruleRefExpr{
								pos:  position{line: 569, col: 70, offset: 25531},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 569, col: 74, offset: 25535},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 569, col: 82, offset: 25543},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 569, col: 90, offset: 25551},
								expr: &choiceExpr{
									pos: position{line: 569, col: 91, offset: 25552},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 569, col: 91, offset: 25552},
											name: "List",
										},
										&ruleRefExpr{
											pos:  position{line: 569, col: 98, offset: 25559},
											name: "Paragraph",
										},
										&ruleRefExpr{
											pos:  position{line: 569, col: 110, offset: 25571},
											name: "BlankLine",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 569, col: 122, offset: 25583},
							name: "ListingBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 569, col: 144, offset: 25605},
							expr: &ruleRefExpr{
								pos:  position{line: 569, col: 144, offset: 25605},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 569, col: 148, offset: 25609},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ExampleBlockDelimiter",
			pos:  position{line: 573, col: 1, offset: 25726},
			expr: &litMatcher{
				pos:        position{line: 573, col: 26, offset: 25751},
				val:        "====",
				ignoreCase: false,
			},
		},
		{
			name: "ExampleBlock",
			pos:  position{line: 575, col: 1, offset: 25759},
			expr: &actionExpr{
				pos: position{line: 575, col: 17, offset: 25775},
				run: (*parser).callonExampleBlock1,
				expr: &seqExpr{
					pos: position{line: 575, col: 17, offset: 25775},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 575, col: 17, offset: 25775},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 575, col: 28, offset: 25786},
								expr: &ruleRefExpr{
									pos:  position{line: 575, col: 29, offset: 25787},
									name: "ElementAttribute",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 575, col: 48, offset: 25806},
							name: "ExampleBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 575, col: 70, offset: 25828},
							expr: &ruleRefExpr{
								pos:  position{line: 575, col: 70, offset: 25828},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 575, col: 74, offset: 25832},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 575, col: 82, offset: 25840},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 575, col: 90, offset: 25848},
								expr: &choiceExpr{
									pos: position{line: 575, col: 91, offset: 25849},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 575, col: 91, offset: 25849},
											name: "List",
										},
										&ruleRefExpr{
											pos:  position{line: 575, col: 98, offset: 25856},
											name: "Paragraph",
										},
										&ruleRefExpr{
											pos:  position{line: 575, col: 110, offset: 25868},
											name: "BlankLine",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 575, col: 123, offset: 25881},
							name: "ExampleBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 575, col: 145, offset: 25903},
							expr: &ruleRefExpr{
								pos:  position{line: 575, col: 145, offset: 25903},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 575, col: 149, offset: 25907},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "LiteralBlock",
			pos:  position{line: 582, col: 1, offset: 26291},
			expr: &choiceExpr{
				pos: position{line: 582, col: 17, offset: 26307},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 582, col: 17, offset: 26307},
						name: "ParagraphWithSpaces",
					},
					&ruleRefExpr{
						pos:  position{line: 582, col: 39, offset: 26329},
						name: "ParagraphWithLiteralBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 582, col: 76, offset: 26366},
						name: "ParagraphWithLiteralAttribute",
					},
				},
			},
		},
		{
			name: "ParagraphWithSpaces",
			pos:  position{line: 585, col: 1, offset: 26461},
			expr: &actionExpr{
				pos: position{line: 585, col: 24, offset: 26484},
				run: (*parser).callonParagraphWithSpaces1,
				expr: &seqExpr{
					pos: position{line: 585, col: 24, offset: 26484},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 585, col: 24, offset: 26484},
							label: "spaces",
							expr: &oneOrMoreExpr{
								pos: position{line: 585, col: 32, offset: 26492},
								expr: &ruleRefExpr{
									pos:  position{line: 585, col: 32, offset: 26492},
									name: "WS",
								},
							},
						},
						&notExpr{
							pos: position{line: 585, col: 37, offset: 26497},
							expr: &ruleRefExpr{
								pos:  position{line: 585, col: 38, offset: 26498},
								name: "NEWLINE",
							},
						},
						&labeledExpr{
							pos:   position{line: 585, col: 46, offset: 26506},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 585, col: 55, offset: 26515},
								name: "LiteralBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 585, col: 76, offset: 26536},
							name: "EndOfLiteralBlock",
						},
					},
				},
			},
		},
		{
			name: "LiteralBlockContent",
			pos:  position{line: 590, col: 1, offset: 26717},
			expr: &actionExpr{
				pos: position{line: 590, col: 24, offset: 26740},
				run: (*parser).callonLiteralBlockContent1,
				expr: &labeledExpr{
					pos:   position{line: 590, col: 24, offset: 26740},
					label: "content",
					expr: &oneOrMoreExpr{
						pos: position{line: 590, col: 32, offset: 26748},
						expr: &seqExpr{
							pos: position{line: 590, col: 33, offset: 26749},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 590, col: 33, offset: 26749},
									expr: &seqExpr{
										pos: position{line: 590, col: 35, offset: 26751},
										exprs: []interface{}{
											&ruleRefExpr{
												pos:  position{line: 590, col: 35, offset: 26751},
												name: "NEWLINE",
											},
											&ruleRefExpr{
												pos:  position{line: 590, col: 43, offset: 26759},
												name: "BlankLine",
											},
										},
									},
								},
								&anyMatcher{
									line: 590, col: 54, offset: 26770,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "EndOfLiteralBlock",
			pos:  position{line: 595, col: 1, offset: 26855},
			expr: &choiceExpr{
				pos: position{line: 595, col: 22, offset: 26876},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 595, col: 22, offset: 26876},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 595, col: 22, offset: 26876},
								name: "NEWLINE",
							},
							&ruleRefExpr{
								pos:  position{line: 595, col: 30, offset: 26884},
								name: "BlankLine",
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 595, col: 42, offset: 26896},
						name: "NEWLINE",
					},
					&ruleRefExpr{
						pos:  position{line: 595, col: 52, offset: 26906},
						name: "EOF",
					},
				},
			},
		},
		{
			name: "ParagraphWithLiteralBlockDelimiter",
			pos:  position{line: 598, col: 1, offset: 26966},
			expr: &actionExpr{
				pos: position{line: 598, col: 39, offset: 27004},
				run: (*parser).callonParagraphWithLiteralBlockDelimiter1,
				expr: &seqExpr{
					pos: position{line: 598, col: 39, offset: 27004},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 598, col: 39, offset: 27004},
							name: "LiteralBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 598, col: 61, offset: 27026},
							expr: &ruleRefExpr{
								pos:  position{line: 598, col: 61, offset: 27026},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 598, col: 65, offset: 27030},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 598, col: 73, offset: 27038},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 598, col: 81, offset: 27046},
								expr: &seqExpr{
									pos: position{line: 598, col: 82, offset: 27047},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 598, col: 82, offset: 27047},
											expr: &ruleRefExpr{
												pos:  position{line: 598, col: 83, offset: 27048},
												name: "LiteralBlockDelimiter",
											},
										},
										&anyMatcher{
											line: 598, col: 105, offset: 27070,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 598, col: 109, offset: 27074},
							name: "LiteralBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 598, col: 131, offset: 27096},
							expr: &ruleRefExpr{
								pos:  position{line: 598, col: 131, offset: 27096},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 598, col: 135, offset: 27100},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "LiteralBlockDelimiter",
			pos:  position{line: 602, col: 1, offset: 27184},
			expr: &litMatcher{
				pos:        position{line: 602, col: 26, offset: 27209},
				val:        "....",
				ignoreCase: false,
			},
		},
		{
			name: "ParagraphWithLiteralAttribute",
			pos:  position{line: 605, col: 1, offset: 27271},
			expr: &actionExpr{
				pos: position{line: 605, col: 34, offset: 27304},
				run: (*parser).callonParagraphWithLiteralAttribute1,
				expr: &seqExpr{
					pos: position{line: 605, col: 34, offset: 27304},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 605, col: 34, offset: 27304},
							val:        "[literal]",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 605, col: 46, offset: 27316},
							expr: &ruleRefExpr{
								pos:  position{line: 605, col: 46, offset: 27316},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 605, col: 50, offset: 27320},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 605, col: 58, offset: 27328},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 605, col: 67, offset: 27337},
								name: "LiteralBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 605, col: 88, offset: 27358},
							name: "EndOfLiteralBlock",
						},
					},
				},
			},
		},
		{
			name: "BlankLine",
			pos:  position{line: 612, col: 1, offset: 27561},
			expr: &actionExpr{
				pos: position{line: 612, col: 14, offset: 27574},
				run: (*parser).callonBlankLine1,
				expr: &seqExpr{
					pos: position{line: 612, col: 14, offset: 27574},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 612, col: 14, offset: 27574},
							expr: &ruleRefExpr{
								pos:  position{line: 612, col: 15, offset: 27575},
								name: "EOF",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 612, col: 19, offset: 27579},
							expr: &ruleRefExpr{
								pos:  position{line: 612, col: 19, offset: 27579},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 612, col: 23, offset: 27583},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "Characters",
			pos:  position{line: 619, col: 1, offset: 27730},
			expr: &actionExpr{
				pos: position{line: 619, col: 15, offset: 27744},
				run: (*parser).callonCharacters1,
				expr: &oneOrMoreExpr{
					pos: position{line: 619, col: 15, offset: 27744},
					expr: &seqExpr{
						pos: position{line: 619, col: 16, offset: 27745},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 619, col: 16, offset: 27745},
								expr: &ruleRefExpr{
									pos:  position{line: 619, col: 17, offset: 27746},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 619, col: 25, offset: 27754},
								expr: &ruleRefExpr{
									pos:  position{line: 619, col: 26, offset: 27755},
									name: "WS",
								},
							},
							&anyMatcher{
								line: 619, col: 29, offset: 27758,
							},
						},
					},
				},
			},
		},
		{
			name: "URL",
			pos:  position{line: 623, col: 1, offset: 27798},
			expr: &actionExpr{
				pos: position{line: 623, col: 8, offset: 27805},
				run: (*parser).callonURL1,
				expr: &oneOrMoreExpr{
					pos: position{line: 623, col: 8, offset: 27805},
					expr: &seqExpr{
						pos: position{line: 623, col: 9, offset: 27806},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 623, col: 9, offset: 27806},
								expr: &ruleRefExpr{
									pos:  position{line: 623, col: 10, offset: 27807},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 623, col: 18, offset: 27815},
								expr: &ruleRefExpr{
									pos:  position{line: 623, col: 19, offset: 27816},
									name: "WS",
								},
							},
							&notExpr{
								pos: position{line: 623, col: 22, offset: 27819},
								expr: &litMatcher{
									pos:        position{line: 623, col: 23, offset: 27820},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 623, col: 27, offset: 27824},
								expr: &litMatcher{
									pos:        position{line: 623, col: 28, offset: 27825},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 623, col: 32, offset: 27829,
							},
						},
					},
				},
			},
		},
		{
			name: "ID",
			pos:  position{line: 627, col: 1, offset: 27869},
			expr: &actionExpr{
				pos: position{line: 627, col: 7, offset: 27875},
				run: (*parser).callonID1,
				expr: &oneOrMoreExpr{
					pos: position{line: 627, col: 7, offset: 27875},
					expr: &seqExpr{
						pos: position{line: 627, col: 8, offset: 27876},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 627, col: 8, offset: 27876},
								expr: &ruleRefExpr{
									pos:  position{line: 627, col: 9, offset: 27877},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 627, col: 17, offset: 27885},
								expr: &ruleRefExpr{
									pos:  position{line: 627, col: 18, offset: 27886},
									name: "WS",
								},
							},
							&notExpr{
								pos: position{line: 627, col: 21, offset: 27889},
								expr: &litMatcher{
									pos:        position{line: 627, col: 22, offset: 27890},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 627, col: 26, offset: 27894},
								expr: &litMatcher{
									pos:        position{line: 627, col: 27, offset: 27895},
									val:        "]",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 627, col: 31, offset: 27899},
								expr: &litMatcher{
									pos:        position{line: 627, col: 32, offset: 27900},
									val:        "<<",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 627, col: 37, offset: 27905},
								expr: &litMatcher{
									pos:        position{line: 627, col: 38, offset: 27906},
									val:        ">>",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 627, col: 42, offset: 27910,
							},
						},
					},
				},
			},
		},
		{
			name: "URL_TEXT",
			pos:  position{line: 631, col: 1, offset: 27950},
			expr: &actionExpr{
				pos: position{line: 631, col: 13, offset: 27962},
				run: (*parser).callonURL_TEXT1,
				expr: &oneOrMoreExpr{
					pos: position{line: 631, col: 13, offset: 27962},
					expr: &seqExpr{
						pos: position{line: 631, col: 14, offset: 27963},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 631, col: 14, offset: 27963},
								expr: &ruleRefExpr{
									pos:  position{line: 631, col: 15, offset: 27964},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 631, col: 23, offset: 27972},
								expr: &litMatcher{
									pos:        position{line: 631, col: 24, offset: 27973},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 631, col: 28, offset: 27977},
								expr: &litMatcher{
									pos:        position{line: 631, col: 29, offset: 27978},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 631, col: 33, offset: 27982,
							},
						},
					},
				},
			},
		},
		{
			name: "URL_SCHEME",
			pos:  position{line: 635, col: 1, offset: 28022},
			expr: &choiceExpr{
				pos: position{line: 635, col: 15, offset: 28036},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 635, col: 15, offset: 28036},
						val:        "http://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 635, col: 27, offset: 28048},
						val:        "https://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 635, col: 40, offset: 28061},
						val:        "ftp://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 635, col: 51, offset: 28072},
						val:        "irc://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 635, col: 62, offset: 28083},
						val:        "mailto:",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "DIGIT",
			pos:  position{line: 637, col: 1, offset: 28094},
			expr: &charClassMatcher{
				pos:        position{line: 637, col: 10, offset: 28103},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NEWLINE",
			pos:  position{line: 639, col: 1, offset: 28110},
			expr: &choiceExpr{
				pos: position{line: 639, col: 12, offset: 28121},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 639, col: 12, offset: 28121},
						val:        "\r\n",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 639, col: 21, offset: 28130},
						val:        "\r",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 639, col: 28, offset: 28137},
						val:        "\n",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "WS",
			pos:  position{line: 641, col: 1, offset: 28143},
			expr: &choiceExpr{
				pos: position{line: 641, col: 7, offset: 28149},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 641, col: 7, offset: 28149},
						val:        " ",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 641, col: 13, offset: 28155},
						run: (*parser).callonWS3,
						expr: &litMatcher{
							pos:        position{line: 641, col: 13, offset: 28155},
							val:        "\t",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 645, col: 1, offset: 28200},
			expr: &notExpr{
				pos: position{line: 645, col: 8, offset: 28207},
				expr: &anyMatcher{
					line: 645, col: 9, offset: 28208,
				},
			},
		},
		{
			name: "EOL",
			pos:  position{line: 647, col: 1, offset: 28211},
			expr: &choiceExpr{
				pos: position{line: 647, col: 8, offset: 28218},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 647, col: 8, offset: 28218},
						name: "NEWLINE",
					},
					&ruleRefExpr{
						pos:  position{line: 647, col: 18, offset: 28228},
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

	return types.NewSectionTitle(content.(types.InlineContent), append(attributes.([]interface{}), id))
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
	return content.(types.DocElement), nil
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
	return content.(types.DocElement), nil
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
	return content.(types.DocElement), nil
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
	return content.(types.DocElement), nil
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
	return content.(types.DocElement), nil
}

func (p *parser) callonSection5Block1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection5Block1(stack["content"])
}

func (c *current) onSection1Title1(attributes, level, content, id interface{}) (interface{}, error) {

	return types.NewSectionTitle(content.(types.InlineContent), append(attributes.([]interface{}), id))
}

func (p *parser) callonSection1Title1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection1Title1(stack["attributes"], stack["level"], stack["content"], stack["id"])
}

func (c *current) onSection2Title1(attributes, level, content, id interface{}) (interface{}, error) {
	return types.NewSectionTitle(content.(types.InlineContent), append(attributes.([]interface{}), id))
}

func (p *parser) callonSection2Title1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection2Title1(stack["attributes"], stack["level"], stack["content"], stack["id"])
}

func (c *current) onSection3Title1(attributes, level, content, id interface{}) (interface{}, error) {
	return types.NewSectionTitle(content.(types.InlineContent), append(attributes.([]interface{}), id))
}

func (p *parser) callonSection3Title1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection3Title1(stack["attributes"], stack["level"], stack["content"], stack["id"])
}

func (c *current) onSection4Title1(attributes, level, content, id interface{}) (interface{}, error) {
	return types.NewSectionTitle(content.(types.InlineContent), append(attributes.([]interface{}), id))
}

func (p *parser) callonSection4Title1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection4Title1(stack["attributes"], stack["level"], stack["content"], stack["id"])
}

func (c *current) onSection5Title1(attributes, level, content, id interface{}) (interface{}, error) {
	return types.NewSectionTitle(content.(types.InlineContent), append(attributes.([]interface{}), id))
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
	return types.NewOrderedListItem(prefix.(types.OrderedListItemPrefix), content.([]types.DocElement), attributes.([]interface{}))
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
	return types.NewUnorderedListItem(prefix.(types.UnorderedListItemPrefix), content.([]types.DocElement))
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
	return types.NewLabeledListItem(term.([]interface{}), description.([]types.DocElement))
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

func (c *current) onInlineContentWithTrailingSpaces1(elements interface{}) (interface{}, error) {
	// includes heading and trailing spaces in the elements arg
	return types.NewInlineContent(elements.([]interface{}))
}

func (p *parser) callonInlineContentWithTrailingSpaces1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInlineContentWithTrailingSpaces1(stack["elements"])
}

func (c *current) onInlineContent1(elements interface{}) (interface{}, error) {
	// absorbs heading and trailing spaces
	return types.NewInlineContent(elements.([]interface{}))
}

func (p *parser) callonInlineContent1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInlineContent1(stack["elements"])
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
