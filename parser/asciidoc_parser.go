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
			pos:  position{line: 60, col: 1, offset: 2196},
			expr: &choiceExpr{
				pos: position{line: 60, col: 20, offset: 2215},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 60, col: 20, offset: 2215},
						name: "DocumentAuthorsInlineForm",
					},
					&ruleRefExpr{
						pos:  position{line: 60, col: 48, offset: 2243},
						name: "DocumentAuthorsAttributeForm",
					},
				},
			},
		},
		{
			name: "DocumentAuthorsInlineForm",
			pos:  position{line: 62, col: 1, offset: 2273},
			expr: &actionExpr{
				pos: position{line: 62, col: 30, offset: 2302},
				run: (*parser).callonDocumentAuthorsInlineForm1,
				expr: &seqExpr{
					pos: position{line: 62, col: 30, offset: 2302},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 62, col: 30, offset: 2302},
							expr: &ruleRefExpr{
								pos:  position{line: 62, col: 30, offset: 2302},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 62, col: 34, offset: 2306},
							expr: &litMatcher{
								pos:        position{line: 62, col: 35, offset: 2307},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 62, col: 39, offset: 2311},
							label: "authors",
							expr: &oneOrMoreExpr{
								pos: position{line: 62, col: 48, offset: 2320},
								expr: &ruleRefExpr{
									pos:  position{line: 62, col: 48, offset: 2320},
									name: "DocumentAuthor",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 62, col: 65, offset: 2337},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthorsAttributeForm",
			pos:  position{line: 66, col: 1, offset: 2407},
			expr: &actionExpr{
				pos: position{line: 66, col: 33, offset: 2439},
				run: (*parser).callonDocumentAuthorsAttributeForm1,
				expr: &seqExpr{
					pos: position{line: 66, col: 33, offset: 2439},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 66, col: 33, offset: 2439},
							expr: &ruleRefExpr{
								pos:  position{line: 66, col: 33, offset: 2439},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 66, col: 37, offset: 2443},
							val:        ":author:",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 66, col: 48, offset: 2454},
							label: "author",
							expr: &ruleRefExpr{
								pos:  position{line: 66, col: 56, offset: 2462},
								name: "DocumentAuthor",
							},
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthor",
			pos:  position{line: 70, col: 1, offset: 2555},
			expr: &actionExpr{
				pos: position{line: 70, col: 19, offset: 2573},
				run: (*parser).callonDocumentAuthor1,
				expr: &seqExpr{
					pos: position{line: 70, col: 19, offset: 2573},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 70, col: 19, offset: 2573},
							expr: &ruleRefExpr{
								pos:  position{line: 70, col: 19, offset: 2573},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 70, col: 23, offset: 2577},
							label: "namePart1",
							expr: &ruleRefExpr{
								pos:  position{line: 70, col: 34, offset: 2588},
								name: "DocumentAuthorNamePart",
							},
						},
						&labeledExpr{
							pos:   position{line: 70, col: 58, offset: 2612},
							label: "namePart2",
							expr: &zeroOrOneExpr{
								pos: position{line: 70, col: 68, offset: 2622},
								expr: &ruleRefExpr{
									pos:  position{line: 70, col: 69, offset: 2623},
									name: "DocumentAuthorNamePart",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 70, col: 94, offset: 2648},
							label: "namePart3",
							expr: &zeroOrOneExpr{
								pos: position{line: 70, col: 104, offset: 2658},
								expr: &ruleRefExpr{
									pos:  position{line: 70, col: 105, offset: 2659},
									name: "DocumentAuthorNamePart",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 70, col: 130, offset: 2684},
							label: "email",
							expr: &zeroOrOneExpr{
								pos: position{line: 70, col: 136, offset: 2690},
								expr: &ruleRefExpr{
									pos:  position{line: 70, col: 137, offset: 2691},
									name: "DocumentAuthorEmail",
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 70, col: 159, offset: 2713},
							expr: &ruleRefExpr{
								pos:  position{line: 70, col: 159, offset: 2713},
								name: "WS",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 70, col: 163, offset: 2717},
							expr: &litMatcher{
								pos:        position{line: 70, col: 163, offset: 2717},
								val:        ";",
								ignoreCase: false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 70, col: 168, offset: 2722},
							expr: &ruleRefExpr{
								pos:  position{line: 70, col: 168, offset: 2722},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthorNamePart",
			pos:  position{line: 75, col: 1, offset: 2887},
			expr: &seqExpr{
				pos: position{line: 75, col: 27, offset: 2913},
				exprs: []interface{}{
					&notExpr{
						pos: position{line: 75, col: 27, offset: 2913},
						expr: &litMatcher{
							pos:        position{line: 75, col: 28, offset: 2914},
							val:        "<",
							ignoreCase: false,
						},
					},
					&notExpr{
						pos: position{line: 75, col: 32, offset: 2918},
						expr: &litMatcher{
							pos:        position{line: 75, col: 33, offset: 2919},
							val:        ";",
							ignoreCase: false,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 37, offset: 2923},
						name: "Characters",
					},
					&zeroOrMoreExpr{
						pos: position{line: 75, col: 48, offset: 2934},
						expr: &ruleRefExpr{
							pos:  position{line: 75, col: 48, offset: 2934},
							name: "WS",
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthorEmail",
			pos:  position{line: 77, col: 1, offset: 2939},
			expr: &seqExpr{
				pos: position{line: 77, col: 24, offset: 2962},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 77, col: 24, offset: 2962},
						val:        "<",
						ignoreCase: false,
					},
					&labeledExpr{
						pos:   position{line: 77, col: 28, offset: 2966},
						label: "email",
						expr: &oneOrMoreExpr{
							pos: position{line: 77, col: 34, offset: 2972},
							expr: &seqExpr{
								pos: position{line: 77, col: 35, offset: 2973},
								exprs: []interface{}{
									&notExpr{
										pos: position{line: 77, col: 35, offset: 2973},
										expr: &litMatcher{
											pos:        position{line: 77, col: 36, offset: 2974},
											val:        ">",
											ignoreCase: false,
										},
									},
									&notExpr{
										pos: position{line: 77, col: 40, offset: 2978},
										expr: &ruleRefExpr{
											pos:  position{line: 77, col: 41, offset: 2979},
											name: "EOL",
										},
									},
									&anyMatcher{
										line: 77, col: 45, offset: 2983,
									},
								},
							},
						},
					},
					&litMatcher{
						pos:        position{line: 77, col: 49, offset: 2987},
						val:        ">",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "DocumentRevision",
			pos:  position{line: 81, col: 1, offset: 3123},
			expr: &actionExpr{
				pos: position{line: 81, col: 21, offset: 3143},
				run: (*parser).callonDocumentRevision1,
				expr: &seqExpr{
					pos: position{line: 81, col: 21, offset: 3143},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 81, col: 21, offset: 3143},
							expr: &ruleRefExpr{
								pos:  position{line: 81, col: 21, offset: 3143},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 81, col: 25, offset: 3147},
							expr: &litMatcher{
								pos:        position{line: 81, col: 26, offset: 3148},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 81, col: 30, offset: 3152},
							label: "revnumber",
							expr: &zeroOrOneExpr{
								pos: position{line: 81, col: 40, offset: 3162},
								expr: &ruleRefExpr{
									pos:  position{line: 81, col: 41, offset: 3163},
									name: "DocumentRevisionNumber",
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 81, col: 66, offset: 3188},
							expr: &litMatcher{
								pos:        position{line: 81, col: 66, offset: 3188},
								val:        ",",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 81, col: 71, offset: 3193},
							label: "revdate",
							expr: &zeroOrOneExpr{
								pos: position{line: 81, col: 79, offset: 3201},
								expr: &ruleRefExpr{
									pos:  position{line: 81, col: 80, offset: 3202},
									name: "DocumentRevisionDate",
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 81, col: 103, offset: 3225},
							expr: &litMatcher{
								pos:        position{line: 81, col: 103, offset: 3225},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 81, col: 108, offset: 3230},
							label: "revremark",
							expr: &zeroOrOneExpr{
								pos: position{line: 81, col: 118, offset: 3240},
								expr: &ruleRefExpr{
									pos:  position{line: 81, col: 119, offset: 3241},
									name: "DocumentRevisionRemark",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 81, col: 144, offset: 3266},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentRevisionNumber",
			pos:  position{line: 86, col: 1, offset: 3439},
			expr: &choiceExpr{
				pos: position{line: 86, col: 27, offset: 3465},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 86, col: 27, offset: 3465},
						exprs: []interface{}{
							&litMatcher{
								pos:        position{line: 86, col: 27, offset: 3465},
								val:        "v",
								ignoreCase: true,
							},
							&ruleRefExpr{
								pos:  position{line: 86, col: 32, offset: 3470},
								name: "DIGIT",
							},
							&zeroOrMoreExpr{
								pos: position{line: 86, col: 39, offset: 3477},
								expr: &seqExpr{
									pos: position{line: 86, col: 40, offset: 3478},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 86, col: 40, offset: 3478},
											expr: &ruleRefExpr{
												pos:  position{line: 86, col: 41, offset: 3479},
												name: "EOL",
											},
										},
										&notExpr{
											pos: position{line: 86, col: 45, offset: 3483},
											expr: &litMatcher{
												pos:        position{line: 86, col: 46, offset: 3484},
												val:        ",",
												ignoreCase: false,
											},
										},
										&notExpr{
											pos: position{line: 86, col: 50, offset: 3488},
											expr: &litMatcher{
												pos:        position{line: 86, col: 51, offset: 3489},
												val:        ":",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 86, col: 55, offset: 3493,
										},
									},
								},
							},
						},
					},
					&seqExpr{
						pos: position{line: 86, col: 61, offset: 3499},
						exprs: []interface{}{
							&zeroOrOneExpr{
								pos: position{line: 86, col: 61, offset: 3499},
								expr: &litMatcher{
									pos:        position{line: 86, col: 61, offset: 3499},
									val:        "v",
									ignoreCase: true,
								},
							},
							&ruleRefExpr{
								pos:  position{line: 86, col: 67, offset: 3505},
								name: "DIGIT",
							},
							&zeroOrMoreExpr{
								pos: position{line: 86, col: 74, offset: 3512},
								expr: &seqExpr{
									pos: position{line: 86, col: 75, offset: 3513},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 86, col: 75, offset: 3513},
											expr: &ruleRefExpr{
												pos:  position{line: 86, col: 76, offset: 3514},
												name: "EOL",
											},
										},
										&notExpr{
											pos: position{line: 86, col: 80, offset: 3518},
											expr: &litMatcher{
												pos:        position{line: 86, col: 81, offset: 3519},
												val:        ",",
												ignoreCase: false,
											},
										},
										&notExpr{
											pos: position{line: 86, col: 85, offset: 3523},
											expr: &litMatcher{
												pos:        position{line: 86, col: 86, offset: 3524},
												val:        ":",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 86, col: 90, offset: 3528,
										},
									},
								},
							},
							&zeroOrMoreExpr{
								pos: position{line: 86, col: 94, offset: 3532},
								expr: &ruleRefExpr{
									pos:  position{line: 86, col: 94, offset: 3532},
									name: "WS",
								},
							},
							&andExpr{
								pos: position{line: 86, col: 98, offset: 3536},
								expr: &litMatcher{
									pos:        position{line: 86, col: 99, offset: 3537},
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
			pos:  position{line: 87, col: 1, offset: 3541},
			expr: &zeroOrMoreExpr{
				pos: position{line: 87, col: 25, offset: 3565},
				expr: &seqExpr{
					pos: position{line: 87, col: 26, offset: 3566},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 87, col: 26, offset: 3566},
							expr: &ruleRefExpr{
								pos:  position{line: 87, col: 27, offset: 3567},
								name: "EOL",
							},
						},
						&notExpr{
							pos: position{line: 87, col: 31, offset: 3571},
							expr: &litMatcher{
								pos:        position{line: 87, col: 32, offset: 3572},
								val:        ":",
								ignoreCase: false,
							},
						},
						&anyMatcher{
							line: 87, col: 36, offset: 3576,
						},
					},
				},
			},
		},
		{
			name: "DocumentRevisionRemark",
			pos:  position{line: 88, col: 1, offset: 3581},
			expr: &zeroOrMoreExpr{
				pos: position{line: 88, col: 27, offset: 3607},
				expr: &seqExpr{
					pos: position{line: 88, col: 28, offset: 3608},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 88, col: 28, offset: 3608},
							expr: &ruleRefExpr{
								pos:  position{line: 88, col: 29, offset: 3609},
								name: "EOL",
							},
						},
						&anyMatcher{
							line: 88, col: 33, offset: 3613,
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclaration",
			pos:  position{line: 93, col: 1, offset: 3733},
			expr: &choiceExpr{
				pos: position{line: 93, col: 33, offset: 3765},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 93, col: 33, offset: 3765},
						name: "DocumentAttributeDeclarationWithNameOnly",
					},
					&ruleRefExpr{
						pos:  position{line: 93, col: 76, offset: 3808},
						name: "DocumentAttributeDeclarationWithNameAndValue",
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclarationWithNameOnly",
			pos:  position{line: 95, col: 1, offset: 3855},
			expr: &actionExpr{
				pos: position{line: 95, col: 45, offset: 3899},
				run: (*parser).callonDocumentAttributeDeclarationWithNameOnly1,
				expr: &seqExpr{
					pos: position{line: 95, col: 45, offset: 3899},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 95, col: 45, offset: 3899},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 95, col: 49, offset: 3903},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 95, col: 55, offset: 3909},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 95, col: 70, offset: 3924},
							val:        ":",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 95, col: 74, offset: 3928},
							expr: &ruleRefExpr{
								pos:  position{line: 95, col: 74, offset: 3928},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 95, col: 78, offset: 3932},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclarationWithNameAndValue",
			pos:  position{line: 99, col: 1, offset: 4017},
			expr: &actionExpr{
				pos: position{line: 99, col: 49, offset: 4065},
				run: (*parser).callonDocumentAttributeDeclarationWithNameAndValue1,
				expr: &seqExpr{
					pos: position{line: 99, col: 49, offset: 4065},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 99, col: 49, offset: 4065},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 99, col: 53, offset: 4069},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 99, col: 59, offset: 4075},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 99, col: 74, offset: 4090},
							val:        ":",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 99, col: 78, offset: 4094},
							expr: &ruleRefExpr{
								pos:  position{line: 99, col: 78, offset: 4094},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 99, col: 82, offset: 4098},
							label: "value",
							expr: &zeroOrMoreExpr{
								pos: position{line: 99, col: 88, offset: 4104},
								expr: &seqExpr{
									pos: position{line: 99, col: 89, offset: 4105},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 99, col: 89, offset: 4105},
											expr: &ruleRefExpr{
												pos:  position{line: 99, col: 90, offset: 4106},
												name: "NEWLINE",
											},
										},
										&anyMatcher{
											line: 99, col: 98, offset: 4114,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 99, col: 102, offset: 4118},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeReset",
			pos:  position{line: 103, col: 1, offset: 4221},
			expr: &choiceExpr{
				pos: position{line: 103, col: 27, offset: 4247},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 103, col: 27, offset: 4247},
						name: "DocumentAttributeResetWithSectionTitleBangSymbol",
					},
					&ruleRefExpr{
						pos:  position{line: 103, col: 78, offset: 4298},
						name: "DocumentAttributeResetWithTrailingBangSymbol",
					},
				},
			},
		},
		{
			name: "DocumentAttributeResetWithSectionTitleBangSymbol",
			pos:  position{line: 105, col: 1, offset: 4344},
			expr: &actionExpr{
				pos: position{line: 105, col: 53, offset: 4396},
				run: (*parser).callonDocumentAttributeResetWithSectionTitleBangSymbol1,
				expr: &seqExpr{
					pos: position{line: 105, col: 53, offset: 4396},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 105, col: 53, offset: 4396},
							val:        ":!",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 105, col: 58, offset: 4401},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 105, col: 64, offset: 4407},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 105, col: 79, offset: 4422},
							val:        ":",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 105, col: 83, offset: 4426},
							expr: &ruleRefExpr{
								pos:  position{line: 105, col: 83, offset: 4426},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 105, col: 87, offset: 4430},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeResetWithTrailingBangSymbol",
			pos:  position{line: 109, col: 1, offset: 4504},
			expr: &actionExpr{
				pos: position{line: 109, col: 49, offset: 4552},
				run: (*parser).callonDocumentAttributeResetWithTrailingBangSymbol1,
				expr: &seqExpr{
					pos: position{line: 109, col: 49, offset: 4552},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 109, col: 49, offset: 4552},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 109, col: 53, offset: 4556},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 109, col: 59, offset: 4562},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 109, col: 74, offset: 4577},
							val:        "!:",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 109, col: 79, offset: 4582},
							expr: &ruleRefExpr{
								pos:  position{line: 109, col: 79, offset: 4582},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 109, col: 83, offset: 4586},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeSubstitution",
			pos:  position{line: 114, col: 1, offset: 4661},
			expr: &actionExpr{
				pos: position{line: 114, col: 34, offset: 4694},
				run: (*parser).callonDocumentAttributeSubstitution1,
				expr: &seqExpr{
					pos: position{line: 114, col: 34, offset: 4694},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 114, col: 34, offset: 4694},
							val:        "{",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 114, col: 38, offset: 4698},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 114, col: 44, offset: 4704},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 114, col: 59, offset: 4719},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "AttributeName",
			pos:  position{line: 121, col: 1, offset: 4973},
			expr: &seqExpr{
				pos: position{line: 121, col: 18, offset: 4990},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 121, col: 19, offset: 4991},
						alternatives: []interface{}{
							&charClassMatcher{
								pos:        position{line: 121, col: 19, offset: 4991},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
							&charClassMatcher{
								pos:        position{line: 121, col: 27, offset: 4999},
								val:        "[a-z]",
								ranges:     []rune{'a', 'z'},
								ignoreCase: false,
								inverted:   false,
							},
							&charClassMatcher{
								pos:        position{line: 121, col: 35, offset: 5007},
								val:        "[0-9]",
								ranges:     []rune{'0', '9'},
								ignoreCase: false,
								inverted:   false,
							},
							&litMatcher{
								pos:        position{line: 121, col: 43, offset: 5015},
								val:        "_",
								ignoreCase: false,
							},
						},
					},
					&zeroOrMoreExpr{
						pos: position{line: 121, col: 48, offset: 5020},
						expr: &choiceExpr{
							pos: position{line: 121, col: 49, offset: 5021},
							alternatives: []interface{}{
								&charClassMatcher{
									pos:        position{line: 121, col: 49, offset: 5021},
									val:        "[A-Z]",
									ranges:     []rune{'A', 'Z'},
									ignoreCase: false,
									inverted:   false,
								},
								&charClassMatcher{
									pos:        position{line: 121, col: 57, offset: 5029},
									val:        "[a-z]",
									ranges:     []rune{'a', 'z'},
									ignoreCase: false,
									inverted:   false,
								},
								&charClassMatcher{
									pos:        position{line: 121, col: 65, offset: 5037},
									val:        "[0-9]",
									ranges:     []rune{'0', '9'},
									ignoreCase: false,
									inverted:   false,
								},
								&litMatcher{
									pos:        position{line: 121, col: 73, offset: 5045},
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
			pos:  position{line: 126, col: 1, offset: 5165},
			expr: &seqExpr{
				pos: position{line: 126, col: 25, offset: 5189},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 126, col: 25, offset: 5189},
						val:        "toc::[]",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 126, col: 35, offset: 5199},
						name: "NEWLINE",
					},
				},
			},
		},
		{
			name: "Section",
			pos:  position{line: 131, col: 1, offset: 5312},
			expr: &choiceExpr{
				pos: position{line: 131, col: 12, offset: 5323},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 131, col: 12, offset: 5323},
						name: "Section1",
					},
					&ruleRefExpr{
						pos:  position{line: 131, col: 23, offset: 5334},
						name: "Section2",
					},
					&ruleRefExpr{
						pos:  position{line: 131, col: 34, offset: 5345},
						name: "Section3",
					},
					&ruleRefExpr{
						pos:  position{line: 131, col: 45, offset: 5356},
						name: "Section4",
					},
					&ruleRefExpr{
						pos:  position{line: 131, col: 56, offset: 5367},
						name: "Section5",
					},
				},
			},
		},
		{
			name: "Section1",
			pos:  position{line: 134, col: 1, offset: 5378},
			expr: &actionExpr{
				pos: position{line: 134, col: 13, offset: 5390},
				run: (*parser).callonSection11,
				expr: &seqExpr{
					pos: position{line: 134, col: 13, offset: 5390},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 134, col: 13, offset: 5390},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 134, col: 21, offset: 5398},
								name: "Section1Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 134, col: 36, offset: 5413},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 134, col: 46, offset: 5423},
								expr: &ruleRefExpr{
									pos:  position{line: 134, col: 46, offset: 5423},
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
			pos:  position{line: 138, col: 1, offset: 5531},
			expr: &actionExpr{
				pos: position{line: 138, col: 18, offset: 5548},
				run: (*parser).callonSection1Block1,
				expr: &seqExpr{
					pos: position{line: 138, col: 18, offset: 5548},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 138, col: 18, offset: 5548},
							expr: &ruleRefExpr{
								pos:  position{line: 138, col: 19, offset: 5549},
								name: "Section1",
							},
						},
						&labeledExpr{
							pos:   position{line: 138, col: 28, offset: 5558},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 138, col: 37, offset: 5567},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 138, col: 37, offset: 5567},
										name: "Section2",
									},
									&ruleRefExpr{
										pos:  position{line: 138, col: 48, offset: 5578},
										name: "Section3",
									},
									&ruleRefExpr{
										pos:  position{line: 138, col: 59, offset: 5589},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 138, col: 70, offset: 5600},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 138, col: 81, offset: 5611},
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
			pos:  position{line: 142, col: 1, offset: 5673},
			expr: &actionExpr{
				pos: position{line: 142, col: 13, offset: 5685},
				run: (*parser).callonSection21,
				expr: &seqExpr{
					pos: position{line: 142, col: 13, offset: 5685},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 142, col: 13, offset: 5685},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 142, col: 21, offset: 5693},
								name: "Section2Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 142, col: 36, offset: 5708},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 142, col: 46, offset: 5718},
								expr: &ruleRefExpr{
									pos:  position{line: 142, col: 46, offset: 5718},
									name: "Section2Block",
								},
							},
						},
						&andExpr{
							pos: position{line: 142, col: 62, offset: 5734},
							expr: &zeroOrMoreExpr{
								pos: position{line: 142, col: 63, offset: 5735},
								expr: &ruleRefExpr{
									pos:  position{line: 142, col: 64, offset: 5736},
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
			pos:  position{line: 146, col: 1, offset: 5839},
			expr: &actionExpr{
				pos: position{line: 146, col: 18, offset: 5856},
				run: (*parser).callonSection2Block1,
				expr: &seqExpr{
					pos: position{line: 146, col: 18, offset: 5856},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 146, col: 18, offset: 5856},
							expr: &ruleRefExpr{
								pos:  position{line: 146, col: 19, offset: 5857},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 146, col: 28, offset: 5866},
							expr: &ruleRefExpr{
								pos:  position{line: 146, col: 29, offset: 5867},
								name: "Section2",
							},
						},
						&labeledExpr{
							pos:   position{line: 146, col: 38, offset: 5876},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 146, col: 47, offset: 5885},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 146, col: 47, offset: 5885},
										name: "Section3",
									},
									&ruleRefExpr{
										pos:  position{line: 146, col: 58, offset: 5896},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 146, col: 69, offset: 5907},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 146, col: 80, offset: 5918},
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
			pos:  position{line: 150, col: 1, offset: 5980},
			expr: &actionExpr{
				pos: position{line: 150, col: 13, offset: 5992},
				run: (*parser).callonSection31,
				expr: &seqExpr{
					pos: position{line: 150, col: 13, offset: 5992},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 150, col: 13, offset: 5992},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 150, col: 21, offset: 6000},
								name: "Section3Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 150, col: 36, offset: 6015},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 150, col: 46, offset: 6025},
								expr: &ruleRefExpr{
									pos:  position{line: 150, col: 46, offset: 6025},
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
			pos:  position{line: 154, col: 1, offset: 6133},
			expr: &actionExpr{
				pos: position{line: 154, col: 18, offset: 6150},
				run: (*parser).callonSection3Block1,
				expr: &seqExpr{
					pos: position{line: 154, col: 18, offset: 6150},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 154, col: 18, offset: 6150},
							expr: &ruleRefExpr{
								pos:  position{line: 154, col: 19, offset: 6151},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 154, col: 28, offset: 6160},
							expr: &ruleRefExpr{
								pos:  position{line: 154, col: 29, offset: 6161},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 154, col: 38, offset: 6170},
							expr: &ruleRefExpr{
								pos:  position{line: 154, col: 39, offset: 6171},
								name: "Section3",
							},
						},
						&labeledExpr{
							pos:   position{line: 154, col: 48, offset: 6180},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 154, col: 57, offset: 6189},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 154, col: 57, offset: 6189},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 154, col: 68, offset: 6200},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 154, col: 79, offset: 6211},
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
			pos:  position{line: 158, col: 1, offset: 6273},
			expr: &actionExpr{
				pos: position{line: 158, col: 13, offset: 6285},
				run: (*parser).callonSection41,
				expr: &seqExpr{
					pos: position{line: 158, col: 13, offset: 6285},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 158, col: 13, offset: 6285},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 158, col: 21, offset: 6293},
								name: "Section4Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 158, col: 36, offset: 6308},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 158, col: 46, offset: 6318},
								expr: &ruleRefExpr{
									pos:  position{line: 158, col: 46, offset: 6318},
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
			pos:  position{line: 162, col: 1, offset: 6426},
			expr: &actionExpr{
				pos: position{line: 162, col: 18, offset: 6443},
				run: (*parser).callonSection4Block1,
				expr: &seqExpr{
					pos: position{line: 162, col: 18, offset: 6443},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 162, col: 18, offset: 6443},
							expr: &ruleRefExpr{
								pos:  position{line: 162, col: 19, offset: 6444},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 162, col: 28, offset: 6453},
							expr: &ruleRefExpr{
								pos:  position{line: 162, col: 29, offset: 6454},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 162, col: 38, offset: 6463},
							expr: &ruleRefExpr{
								pos:  position{line: 162, col: 39, offset: 6464},
								name: "Section3",
							},
						},
						&notExpr{
							pos: position{line: 162, col: 48, offset: 6473},
							expr: &ruleRefExpr{
								pos:  position{line: 162, col: 49, offset: 6474},
								name: "Section4",
							},
						},
						&labeledExpr{
							pos:   position{line: 162, col: 58, offset: 6483},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 162, col: 67, offset: 6492},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 162, col: 67, offset: 6492},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 162, col: 78, offset: 6503},
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
			pos:  position{line: 166, col: 1, offset: 6565},
			expr: &actionExpr{
				pos: position{line: 166, col: 13, offset: 6577},
				run: (*parser).callonSection51,
				expr: &seqExpr{
					pos: position{line: 166, col: 13, offset: 6577},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 166, col: 13, offset: 6577},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 166, col: 21, offset: 6585},
								name: "Section5Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 166, col: 36, offset: 6600},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 166, col: 46, offset: 6610},
								expr: &ruleRefExpr{
									pos:  position{line: 166, col: 46, offset: 6610},
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
			pos:  position{line: 170, col: 1, offset: 6718},
			expr: &actionExpr{
				pos: position{line: 170, col: 18, offset: 6735},
				run: (*parser).callonSection5Block1,
				expr: &seqExpr{
					pos: position{line: 170, col: 18, offset: 6735},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 170, col: 18, offset: 6735},
							expr: &ruleRefExpr{
								pos:  position{line: 170, col: 19, offset: 6736},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 170, col: 28, offset: 6745},
							expr: &ruleRefExpr{
								pos:  position{line: 170, col: 29, offset: 6746},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 170, col: 38, offset: 6755},
							expr: &ruleRefExpr{
								pos:  position{line: 170, col: 39, offset: 6756},
								name: "Section3",
							},
						},
						&notExpr{
							pos: position{line: 170, col: 48, offset: 6765},
							expr: &ruleRefExpr{
								pos:  position{line: 170, col: 49, offset: 6766},
								name: "Section4",
							},
						},
						&notExpr{
							pos: position{line: 170, col: 58, offset: 6775},
							expr: &ruleRefExpr{
								pos:  position{line: 170, col: 59, offset: 6776},
								name: "Section5",
							},
						},
						&labeledExpr{
							pos:   position{line: 170, col: 68, offset: 6785},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 170, col: 77, offset: 6794},
								name: "BlockElement",
							},
						},
					},
				},
			},
		},
		{
			name: "SectionTitle",
			pos:  position{line: 178, col: 1, offset: 6967},
			expr: &choiceExpr{
				pos: position{line: 178, col: 17, offset: 6983},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 178, col: 17, offset: 6983},
						name: "Section1Title",
					},
					&ruleRefExpr{
						pos:  position{line: 178, col: 33, offset: 6999},
						name: "Section2Title",
					},
					&ruleRefExpr{
						pos:  position{line: 178, col: 49, offset: 7015},
						name: "Section3Title",
					},
					&ruleRefExpr{
						pos:  position{line: 178, col: 65, offset: 7031},
						name: "Section4Title",
					},
					&ruleRefExpr{
						pos:  position{line: 178, col: 81, offset: 7047},
						name: "Section5Title",
					},
				},
			},
		},
		{
			name: "Section1Title",
			pos:  position{line: 180, col: 1, offset: 7062},
			expr: &actionExpr{
				pos: position{line: 180, col: 18, offset: 7079},
				run: (*parser).callonSection1Title1,
				expr: &seqExpr{
					pos: position{line: 180, col: 18, offset: 7079},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 180, col: 18, offset: 7079},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 180, col: 29, offset: 7090},
								expr: &ruleRefExpr{
									pos:  position{line: 180, col: 30, offset: 7091},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 180, col: 49, offset: 7110},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 180, col: 56, offset: 7117},
								val:        "==",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 180, col: 62, offset: 7123},
							expr: &ruleRefExpr{
								pos:  position{line: 180, col: 62, offset: 7123},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 180, col: 66, offset: 7127},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 180, col: 75, offset: 7136},
								name: "InlineContent",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 180, col: 90, offset: 7151},
							expr: &ruleRefExpr{
								pos:  position{line: 180, col: 90, offset: 7151},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 180, col: 94, offset: 7155},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 180, col: 97, offset: 7158},
								expr: &ruleRefExpr{
									pos:  position{line: 180, col: 98, offset: 7159},
									name: "InlineElementID",
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 180, col: 116, offset: 7177},
							expr: &ruleRefExpr{
								pos:  position{line: 180, col: 116, offset: 7177},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 180, col: 120, offset: 7181},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 180, col: 125, offset: 7186},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 180, col: 125, offset: 7186},
									expr: &ruleRefExpr{
										pos:  position{line: 180, col: 125, offset: 7186},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 180, col: 138, offset: 7199},
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
			pos:  position{line: 184, col: 1, offset: 7315},
			expr: &actionExpr{
				pos: position{line: 184, col: 18, offset: 7332},
				run: (*parser).callonSection2Title1,
				expr: &seqExpr{
					pos: position{line: 184, col: 18, offset: 7332},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 184, col: 18, offset: 7332},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 184, col: 29, offset: 7343},
								expr: &ruleRefExpr{
									pos:  position{line: 184, col: 30, offset: 7344},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 184, col: 49, offset: 7363},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 184, col: 56, offset: 7370},
								val:        "===",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 184, col: 63, offset: 7377},
							expr: &ruleRefExpr{
								pos:  position{line: 184, col: 63, offset: 7377},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 184, col: 67, offset: 7381},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 184, col: 76, offset: 7390},
								name: "InlineContent",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 184, col: 91, offset: 7405},
							expr: &ruleRefExpr{
								pos:  position{line: 184, col: 91, offset: 7405},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 184, col: 95, offset: 7409},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 184, col: 98, offset: 7412},
								expr: &ruleRefExpr{
									pos:  position{line: 184, col: 99, offset: 7413},
									name: "InlineElementID",
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 184, col: 117, offset: 7431},
							expr: &ruleRefExpr{
								pos:  position{line: 184, col: 117, offset: 7431},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 184, col: 121, offset: 7435},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 184, col: 126, offset: 7440},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 184, col: 126, offset: 7440},
									expr: &ruleRefExpr{
										pos:  position{line: 184, col: 126, offset: 7440},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 184, col: 139, offset: 7453},
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
			pos:  position{line: 189, col: 1, offset: 7604},
			expr: &actionExpr{
				pos: position{line: 189, col: 18, offset: 7621},
				run: (*parser).callonSection3Title1,
				expr: &seqExpr{
					pos: position{line: 189, col: 18, offset: 7621},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 189, col: 18, offset: 7621},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 189, col: 29, offset: 7632},
								expr: &ruleRefExpr{
									pos:  position{line: 189, col: 30, offset: 7633},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 189, col: 49, offset: 7652},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 189, col: 56, offset: 7659},
								val:        "====",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 189, col: 64, offset: 7667},
							expr: &ruleRefExpr{
								pos:  position{line: 189, col: 64, offset: 7667},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 189, col: 68, offset: 7671},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 189, col: 77, offset: 7680},
								name: "InlineContent",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 189, col: 92, offset: 7695},
							expr: &ruleRefExpr{
								pos:  position{line: 189, col: 92, offset: 7695},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 189, col: 96, offset: 7699},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 189, col: 99, offset: 7702},
								expr: &ruleRefExpr{
									pos:  position{line: 189, col: 100, offset: 7703},
									name: "InlineElementID",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 189, col: 118, offset: 7721},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 189, col: 123, offset: 7726},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 189, col: 123, offset: 7726},
									expr: &ruleRefExpr{
										pos:  position{line: 189, col: 123, offset: 7726},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 189, col: 136, offset: 7739},
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
			pos:  position{line: 193, col: 1, offset: 7854},
			expr: &actionExpr{
				pos: position{line: 193, col: 18, offset: 7871},
				run: (*parser).callonSection4Title1,
				expr: &seqExpr{
					pos: position{line: 193, col: 18, offset: 7871},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 193, col: 18, offset: 7871},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 193, col: 29, offset: 7882},
								expr: &ruleRefExpr{
									pos:  position{line: 193, col: 30, offset: 7883},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 193, col: 49, offset: 7902},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 193, col: 56, offset: 7909},
								val:        "=====",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 193, col: 65, offset: 7918},
							expr: &ruleRefExpr{
								pos:  position{line: 193, col: 65, offset: 7918},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 193, col: 69, offset: 7922},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 193, col: 78, offset: 7931},
								name: "InlineContent",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 193, col: 93, offset: 7946},
							expr: &ruleRefExpr{
								pos:  position{line: 193, col: 93, offset: 7946},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 193, col: 97, offset: 7950},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 193, col: 100, offset: 7953},
								expr: &ruleRefExpr{
									pos:  position{line: 193, col: 101, offset: 7954},
									name: "InlineElementID",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 193, col: 119, offset: 7972},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 193, col: 124, offset: 7977},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 193, col: 124, offset: 7977},
									expr: &ruleRefExpr{
										pos:  position{line: 193, col: 124, offset: 7977},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 193, col: 137, offset: 7990},
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
			pos:  position{line: 197, col: 1, offset: 8105},
			expr: &actionExpr{
				pos: position{line: 197, col: 18, offset: 8122},
				run: (*parser).callonSection5Title1,
				expr: &seqExpr{
					pos: position{line: 197, col: 18, offset: 8122},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 197, col: 18, offset: 8122},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 197, col: 29, offset: 8133},
								expr: &ruleRefExpr{
									pos:  position{line: 197, col: 30, offset: 8134},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 197, col: 49, offset: 8153},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 197, col: 56, offset: 8160},
								val:        "======",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 197, col: 66, offset: 8170},
							expr: &ruleRefExpr{
								pos:  position{line: 197, col: 66, offset: 8170},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 197, col: 70, offset: 8174},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 197, col: 79, offset: 8183},
								name: "InlineContent",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 197, col: 94, offset: 8198},
							expr: &ruleRefExpr{
								pos:  position{line: 197, col: 94, offset: 8198},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 197, col: 98, offset: 8202},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 197, col: 101, offset: 8205},
								expr: &ruleRefExpr{
									pos:  position{line: 197, col: 102, offset: 8206},
									name: "InlineElementID",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 197, col: 120, offset: 8224},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 197, col: 125, offset: 8229},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 197, col: 125, offset: 8229},
									expr: &ruleRefExpr{
										pos:  position{line: 197, col: 125, offset: 8229},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 197, col: 138, offset: 8242},
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
			pos:  position{line: 204, col: 1, offset: 8458},
			expr: &actionExpr{
				pos: position{line: 204, col: 9, offset: 8466},
				run: (*parser).callonList1,
				expr: &seqExpr{
					pos: position{line: 204, col: 9, offset: 8466},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 204, col: 9, offset: 8466},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 204, col: 20, offset: 8477},
								expr: &ruleRefExpr{
									pos:  position{line: 204, col: 21, offset: 8478},
									name: "ListAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 206, col: 5, offset: 8567},
							label: "elements",
							expr: &oneOrMoreExpr{
								pos: position{line: 206, col: 14, offset: 8576},
								expr: &choiceExpr{
									pos: position{line: 206, col: 15, offset: 8577},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 206, col: 15, offset: 8577},
											name: "UnorderedListItem",
										},
										&ruleRefExpr{
											pos:  position{line: 206, col: 35, offset: 8597},
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
			pos:  position{line: 210, col: 1, offset: 8699},
			expr: &actionExpr{
				pos: position{line: 210, col: 18, offset: 8716},
				run: (*parser).callonListAttribute1,
				expr: &seqExpr{
					pos: position{line: 210, col: 18, offset: 8716},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 210, col: 18, offset: 8716},
							label: "attribute",
							expr: &choiceExpr{
								pos: position{line: 210, col: 29, offset: 8727},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 210, col: 29, offset: 8727},
										name: "HorizontalLayout",
									},
									&ruleRefExpr{
										pos:  position{line: 210, col: 48, offset: 8746},
										name: "ListID",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 210, col: 56, offset: 8754},
							name: "NEWLINE",
						},
					},
				},
			},
		},
		{
			name: "ListID",
			pos:  position{line: 214, col: 1, offset: 8793},
			expr: &actionExpr{
				pos: position{line: 214, col: 11, offset: 8803},
				run: (*parser).callonListID1,
				expr: &seqExpr{
					pos: position{line: 214, col: 11, offset: 8803},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 214, col: 11, offset: 8803},
							val:        "[#",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 214, col: 16, offset: 8808},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 214, col: 20, offset: 8812},
								name: "ID",
							},
						},
						&litMatcher{
							pos:        position{line: 214, col: 24, offset: 8816},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "HorizontalLayout",
			pos:  position{line: 218, col: 1, offset: 8882},
			expr: &actionExpr{
				pos: position{line: 218, col: 21, offset: 8902},
				run: (*parser).callonHorizontalLayout1,
				expr: &litMatcher{
					pos:        position{line: 218, col: 21, offset: 8902},
					val:        "[horizontal]",
					ignoreCase: false,
				},
			},
		},
		{
			name: "ListParagraph",
			pos:  position{line: 222, col: 1, offset: 8985},
			expr: &actionExpr{
				pos: position{line: 222, col: 19, offset: 9003},
				run: (*parser).callonListParagraph1,
				expr: &labeledExpr{
					pos:   position{line: 222, col: 19, offset: 9003},
					label: "lines",
					expr: &oneOrMoreExpr{
						pos: position{line: 222, col: 25, offset: 9009},
						expr: &seqExpr{
							pos: position{line: 222, col: 26, offset: 9010},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 222, col: 26, offset: 9010},
									expr: &ruleRefExpr{
										pos:  position{line: 222, col: 28, offset: 9012},
										name: "ListItemContinuation",
									},
								},
								&notExpr{
									pos: position{line: 222, col: 50, offset: 9034},
									expr: &ruleRefExpr{
										pos:  position{line: 222, col: 52, offset: 9036},
										name: "UnorderedListItemPrefix",
									},
								},
								&notExpr{
									pos: position{line: 222, col: 77, offset: 9061},
									expr: &seqExpr{
										pos: position{line: 222, col: 79, offset: 9063},
										exprs: []interface{}{
											&ruleRefExpr{
												pos:  position{line: 222, col: 79, offset: 9063},
												name: "LabeledListItemTerm",
											},
											&ruleRefExpr{
												pos:  position{line: 222, col: 99, offset: 9083},
												name: "LabeledListItemSeparator",
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 222, col: 125, offset: 9109},
									name: "InlineContentWithTrailingSpaces",
								},
								&ruleRefExpr{
									pos:  position{line: 222, col: 157, offset: 9141},
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
			pos:  position{line: 226, col: 1, offset: 9210},
			expr: &actionExpr{
				pos: position{line: 226, col: 25, offset: 9234},
				run: (*parser).callonListItemContinuation1,
				expr: &seqExpr{
					pos: position{line: 226, col: 25, offset: 9234},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 226, col: 25, offset: 9234},
							val:        "+",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 226, col: 29, offset: 9238},
							expr: &ruleRefExpr{
								pos:  position{line: 226, col: 29, offset: 9238},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 226, col: 33, offset: 9242},
							name: "NEWLINE",
						},
					},
				},
			},
		},
		{
			name: "ContinuedBlockElement",
			pos:  position{line: 230, col: 1, offset: 9298},
			expr: &actionExpr{
				pos: position{line: 230, col: 26, offset: 9323},
				run: (*parser).callonContinuedBlockElement1,
				expr: &seqExpr{
					pos: position{line: 230, col: 26, offset: 9323},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 230, col: 26, offset: 9323},
							name: "ListItemContinuation",
						},
						&labeledExpr{
							pos:   position{line: 230, col: 47, offset: 9344},
							label: "element",
							expr: &ruleRefExpr{
								pos:  position{line: 230, col: 55, offset: 9352},
								name: "BlockElement",
							},
						},
					},
				},
			},
		},
		{
			name: "UnorderedListItem",
			pos:  position{line: 237, col: 1, offset: 9505},
			expr: &actionExpr{
				pos: position{line: 237, col: 22, offset: 9526},
				run: (*parser).callonUnorderedListItem1,
				expr: &seqExpr{
					pos: position{line: 237, col: 22, offset: 9526},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 237, col: 22, offset: 9526},
							label: "level",
							expr: &ruleRefExpr{
								pos:  position{line: 237, col: 29, offset: 9533},
								name: "UnorderedListItemPrefix",
							},
						},
						&labeledExpr{
							pos:   position{line: 237, col: 54, offset: 9558},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 237, col: 63, offset: 9567},
								name: "UnorderedListItemContent",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 237, col: 89, offset: 9593},
							expr: &ruleRefExpr{
								pos:  position{line: 237, col: 89, offset: 9593},
								name: "BlankLine",
							},
						},
					},
				},
			},
		},
		{
			name: "UnorderedListItemPrefix",
			pos:  position{line: 241, col: 1, offset: 9684},
			expr: &actionExpr{
				pos: position{line: 241, col: 28, offset: 9711},
				run: (*parser).callonUnorderedListItemPrefix1,
				expr: &seqExpr{
					pos: position{line: 241, col: 28, offset: 9711},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 241, col: 28, offset: 9711},
							expr: &ruleRefExpr{
								pos:  position{line: 241, col: 28, offset: 9711},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 241, col: 32, offset: 9715},
							label: "level",
							expr: &choiceExpr{
								pos: position{line: 241, col: 39, offset: 9722},
								alternatives: []interface{}{
									&oneOrMoreExpr{
										pos: position{line: 241, col: 39, offset: 9722},
										expr: &litMatcher{
											pos:        position{line: 241, col: 39, offset: 9722},
											val:        "*",
											ignoreCase: false,
										},
									},
									&litMatcher{
										pos:        position{line: 241, col: 46, offset: 9729},
										val:        "-",
										ignoreCase: false,
									},
								},
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 241, col: 51, offset: 9734},
							expr: &ruleRefExpr{
								pos:  position{line: 241, col: 51, offset: 9734},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "UnorderedListItemContent",
			pos:  position{line: 245, col: 1, offset: 9832},
			expr: &actionExpr{
				pos: position{line: 245, col: 29, offset: 9860},
				run: (*parser).callonUnorderedListItemContent1,
				expr: &labeledExpr{
					pos:   position{line: 245, col: 29, offset: 9860},
					label: "elements",
					expr: &seqExpr{
						pos: position{line: 245, col: 39, offset: 9870},
						exprs: []interface{}{
							&oneOrMoreExpr{
								pos: position{line: 245, col: 39, offset: 9870},
								expr: &ruleRefExpr{
									pos:  position{line: 245, col: 39, offset: 9870},
									name: "ListParagraph",
								},
							},
							&zeroOrMoreExpr{
								pos: position{line: 245, col: 54, offset: 9885},
								expr: &ruleRefExpr{
									pos:  position{line: 245, col: 54, offset: 9885},
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
			pos:  position{line: 254, col: 1, offset: 10206},
			expr: &choiceExpr{
				pos: position{line: 254, col: 20, offset: 10225},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 254, col: 20, offset: 10225},
						name: "LabeledListItemWithDescription",
					},
					&ruleRefExpr{
						pos:  position{line: 254, col: 53, offset: 10258},
						name: "LabeledListItemWithTermAlone",
					},
				},
			},
		},
		{
			name: "LabeledListItemWithTermAlone",
			pos:  position{line: 256, col: 1, offset: 10288},
			expr: &actionExpr{
				pos: position{line: 256, col: 33, offset: 10320},
				run: (*parser).callonLabeledListItemWithTermAlone1,
				expr: &seqExpr{
					pos: position{line: 256, col: 33, offset: 10320},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 256, col: 33, offset: 10320},
							label: "term",
							expr: &ruleRefExpr{
								pos:  position{line: 256, col: 39, offset: 10326},
								name: "LabeledListItemTerm",
							},
						},
						&litMatcher{
							pos:        position{line: 256, col: 61, offset: 10348},
							val:        "::",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 256, col: 66, offset: 10353},
							expr: &ruleRefExpr{
								pos:  position{line: 256, col: 66, offset: 10353},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 256, col: 70, offset: 10357},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "LabeledListItemTerm",
			pos:  position{line: 260, col: 1, offset: 10494},
			expr: &actionExpr{
				pos: position{line: 260, col: 24, offset: 10517},
				run: (*parser).callonLabeledListItemTerm1,
				expr: &labeledExpr{
					pos:   position{line: 260, col: 24, offset: 10517},
					label: "term",
					expr: &zeroOrMoreExpr{
						pos: position{line: 260, col: 29, offset: 10522},
						expr: &seqExpr{
							pos: position{line: 260, col: 30, offset: 10523},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 260, col: 30, offset: 10523},
									expr: &ruleRefExpr{
										pos:  position{line: 260, col: 31, offset: 10524},
										name: "NEWLINE",
									},
								},
								&notExpr{
									pos: position{line: 260, col: 39, offset: 10532},
									expr: &litMatcher{
										pos:        position{line: 260, col: 40, offset: 10533},
										val:        "::",
										ignoreCase: false,
									},
								},
								&anyMatcher{
									line: 260, col: 45, offset: 10538,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "LabeledListItemSeparator",
			pos:  position{line: 265, col: 1, offset: 10629},
			expr: &seqExpr{
				pos: position{line: 265, col: 30, offset: 10658},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 265, col: 30, offset: 10658},
						val:        "::",
						ignoreCase: false,
					},
					&oneOrMoreExpr{
						pos: position{line: 265, col: 35, offset: 10663},
						expr: &choiceExpr{
							pos: position{line: 265, col: 36, offset: 10664},
							alternatives: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 265, col: 36, offset: 10664},
									name: "WS",
								},
								&ruleRefExpr{
									pos:  position{line: 265, col: 41, offset: 10669},
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
			pos:  position{line: 267, col: 1, offset: 10680},
			expr: &actionExpr{
				pos: position{line: 267, col: 35, offset: 10714},
				run: (*parser).callonLabeledListItemWithDescription1,
				expr: &seqExpr{
					pos: position{line: 267, col: 35, offset: 10714},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 267, col: 35, offset: 10714},
							label: "term",
							expr: &ruleRefExpr{
								pos:  position{line: 267, col: 41, offset: 10720},
								name: "LabeledListItemTerm",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 267, col: 62, offset: 10741},
							name: "LabeledListItemSeparator",
						},
						&labeledExpr{
							pos:   position{line: 267, col: 87, offset: 10766},
							label: "description",
							expr: &ruleRefExpr{
								pos:  position{line: 267, col: 100, offset: 10779},
								name: "LabeledListItemDescription",
							},
						},
					},
				},
			},
		},
		{
			name: "LabeledListItemDescription",
			pos:  position{line: 271, col: 1, offset: 10904},
			expr: &actionExpr{
				pos: position{line: 271, col: 31, offset: 10934},
				run: (*parser).callonLabeledListItemDescription1,
				expr: &labeledExpr{
					pos:   position{line: 271, col: 31, offset: 10934},
					label: "elements",
					expr: &zeroOrMoreExpr{
						pos: position{line: 271, col: 40, offset: 10943},
						expr: &choiceExpr{
							pos: position{line: 271, col: 41, offset: 10944},
							alternatives: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 271, col: 41, offset: 10944},
									name: "ListParagraph",
								},
								&ruleRefExpr{
									pos:  position{line: 271, col: 57, offset: 10960},
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
			pos:  position{line: 280, col: 1, offset: 11310},
			expr: &actionExpr{
				pos: position{line: 280, col: 14, offset: 11323},
				run: (*parser).callonParagraph1,
				expr: &seqExpr{
					pos: position{line: 280, col: 14, offset: 11323},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 280, col: 14, offset: 11323},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 280, col: 25, offset: 11334},
								expr: &ruleRefExpr{
									pos:  position{line: 280, col: 26, offset: 11335},
									name: "ElementAttribute",
								},
							},
						},
						&notExpr{
							pos: position{line: 280, col: 45, offset: 11354},
							expr: &seqExpr{
								pos: position{line: 280, col: 47, offset: 11356},
								exprs: []interface{}{
									&oneOrMoreExpr{
										pos: position{line: 280, col: 47, offset: 11356},
										expr: &litMatcher{
											pos:        position{line: 280, col: 47, offset: 11356},
											val:        "=",
											ignoreCase: false,
										},
									},
									&oneOrMoreExpr{
										pos: position{line: 280, col: 52, offset: 11361},
										expr: &ruleRefExpr{
											pos:  position{line: 280, col: 52, offset: 11361},
											name: "WS",
										},
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 280, col: 57, offset: 11366},
							label: "lines",
							expr: &oneOrMoreExpr{
								pos: position{line: 280, col: 63, offset: 11372},
								expr: &seqExpr{
									pos: position{line: 280, col: 64, offset: 11373},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 280, col: 64, offset: 11373},
											name: "InlineContentWithTrailingSpaces",
										},
										&ruleRefExpr{
											pos:  position{line: 280, col: 96, offset: 11405},
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
			name: "InlineContentWithTrailingSpaces",
			pos:  position{line: 286, col: 1, offset: 11695},
			expr: &actionExpr{
				pos: position{line: 286, col: 36, offset: 11730},
				run: (*parser).callonInlineContentWithTrailingSpaces1,
				expr: &seqExpr{
					pos: position{line: 286, col: 36, offset: 11730},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 286, col: 36, offset: 11730},
							expr: &ruleRefExpr{
								pos:  position{line: 286, col: 37, offset: 11731},
								name: "BlockDelimiter",
							},
						},
						&labeledExpr{
							pos:   position{line: 286, col: 52, offset: 11746},
							label: "elements",
							expr: &oneOrMoreExpr{
								pos: position{line: 286, col: 61, offset: 11755},
								expr: &seqExpr{
									pos: position{line: 286, col: 62, offset: 11756},
									exprs: []interface{}{
										&zeroOrMoreExpr{
											pos: position{line: 286, col: 62, offset: 11756},
											expr: &ruleRefExpr{
												pos:  position{line: 286, col: 62, offset: 11756},
												name: "WS",
											},
										},
										&notExpr{
											pos: position{line: 286, col: 66, offset: 11760},
											expr: &ruleRefExpr{
												pos:  position{line: 286, col: 67, offset: 11761},
												name: "InlineElementID",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 286, col: 83, offset: 11777},
											name: "InlineElement",
										},
										&zeroOrMoreExpr{
											pos: position{line: 286, col: 97, offset: 11791},
											expr: &ruleRefExpr{
												pos:  position{line: 286, col: 97, offset: 11791},
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
			pos:  position{line: 290, col: 1, offset: 11903},
			expr: &actionExpr{
				pos: position{line: 290, col: 18, offset: 11920},
				run: (*parser).callonInlineContent1,
				expr: &seqExpr{
					pos: position{line: 290, col: 18, offset: 11920},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 290, col: 18, offset: 11920},
							expr: &ruleRefExpr{
								pos:  position{line: 290, col: 19, offset: 11921},
								name: "BlockDelimiter",
							},
						},
						&labeledExpr{
							pos:   position{line: 290, col: 34, offset: 11936},
							label: "elements",
							expr: &oneOrMoreExpr{
								pos: position{line: 290, col: 43, offset: 11945},
								expr: &seqExpr{
									pos: position{line: 290, col: 44, offset: 11946},
									exprs: []interface{}{
										&zeroOrMoreExpr{
											pos: position{line: 290, col: 44, offset: 11946},
											expr: &ruleRefExpr{
												pos:  position{line: 290, col: 44, offset: 11946},
												name: "WS",
											},
										},
										&notExpr{
											pos: position{line: 290, col: 48, offset: 11950},
											expr: &ruleRefExpr{
												pos:  position{line: 290, col: 49, offset: 11951},
												name: "InlineElementID",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 290, col: 65, offset: 11967},
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
			pos:  position{line: 294, col: 1, offset: 12089},
			expr: &choiceExpr{
				pos: position{line: 294, col: 19, offset: 12107},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 294, col: 19, offset: 12107},
						name: "CrossReference",
					},
					&ruleRefExpr{
						pos:  position{line: 294, col: 36, offset: 12124},
						name: "Passthrough",
					},
					&ruleRefExpr{
						pos:  position{line: 294, col: 50, offset: 12138},
						name: "InlineImage",
					},
					&ruleRefExpr{
						pos:  position{line: 294, col: 64, offset: 12152},
						name: "QuotedText",
					},
					&ruleRefExpr{
						pos:  position{line: 294, col: 77, offset: 12165},
						name: "Link",
					},
					&ruleRefExpr{
						pos:  position{line: 294, col: 84, offset: 12172},
						name: "DocumentAttributeSubstitution",
					},
					&ruleRefExpr{
						pos:  position{line: 294, col: 116, offset: 12204},
						name: "Characters",
					},
				},
			},
		},
		{
			name: "QuotedText",
			pos:  position{line: 299, col: 1, offset: 12456},
			expr: &choiceExpr{
				pos: position{line: 299, col: 15, offset: 12470},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 299, col: 15, offset: 12470},
						name: "BoldText",
					},
					&ruleRefExpr{
						pos:  position{line: 299, col: 26, offset: 12481},
						name: "ItalicText",
					},
					&ruleRefExpr{
						pos:  position{line: 299, col: 39, offset: 12494},
						name: "MonospaceText",
					},
					&ruleRefExpr{
						pos:  position{line: 300, col: 13, offset: 12522},
						name: "EscapedBoldText",
					},
					&ruleRefExpr{
						pos:  position{line: 300, col: 31, offset: 12540},
						name: "EscapedItalicText",
					},
					&ruleRefExpr{
						pos:  position{line: 300, col: 51, offset: 12560},
						name: "EscapedMonospaceText",
					},
				},
			},
		},
		{
			name: "BoldText",
			pos:  position{line: 302, col: 1, offset: 12582},
			expr: &choiceExpr{
				pos: position{line: 302, col: 13, offset: 12594},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 302, col: 13, offset: 12594},
						name: "BoldTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 302, col: 41, offset: 12622},
						name: "BoldTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 302, col: 73, offset: 12654},
						name: "BoldTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "BoldTextSimplePunctuation",
			pos:  position{line: 304, col: 1, offset: 12727},
			expr: &actionExpr{
				pos: position{line: 304, col: 30, offset: 12756},
				run: (*parser).callonBoldTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 304, col: 30, offset: 12756},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 304, col: 30, offset: 12756},
							expr: &litMatcher{
								pos:        position{line: 304, col: 31, offset: 12757},
								val:        "\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 304, col: 35, offset: 12761},
							val:        "*",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 304, col: 39, offset: 12765},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 304, col: 48, offset: 12774},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 304, col: 67, offset: 12793},
							val:        "*",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "BoldTextDoublePunctuation",
			pos:  position{line: 308, col: 1, offset: 12870},
			expr: &actionExpr{
				pos: position{line: 308, col: 30, offset: 12899},
				run: (*parser).callonBoldTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 308, col: 30, offset: 12899},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 308, col: 30, offset: 12899},
							expr: &litMatcher{
								pos:        position{line: 308, col: 31, offset: 12900},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 308, col: 36, offset: 12905},
							val:        "**",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 308, col: 41, offset: 12910},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 308, col: 50, offset: 12919},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 308, col: 69, offset: 12938},
							val:        "**",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "BoldTextUnbalancedPunctuation",
			pos:  position{line: 312, col: 1, offset: 13016},
			expr: &actionExpr{
				pos: position{line: 312, col: 34, offset: 13049},
				run: (*parser).callonBoldTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 312, col: 34, offset: 13049},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 312, col: 34, offset: 13049},
							expr: &litMatcher{
								pos:        position{line: 312, col: 35, offset: 13050},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 312, col: 40, offset: 13055},
							val:        "**",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 312, col: 45, offset: 13060},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 312, col: 54, offset: 13069},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 312, col: 73, offset: 13088},
							val:        "*",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedBoldText",
			pos:  position{line: 317, col: 1, offset: 13252},
			expr: &choiceExpr{
				pos: position{line: 317, col: 20, offset: 13271},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 317, col: 20, offset: 13271},
						name: "EscapedBoldTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 317, col: 55, offset: 13306},
						name: "EscapedBoldTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 317, col: 94, offset: 13345},
						name: "EscapedBoldTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "EscapedBoldTextSimplePunctuation",
			pos:  position{line: 319, col: 1, offset: 13425},
			expr: &actionExpr{
				pos: position{line: 319, col: 37, offset: 13461},
				run: (*parser).callonEscapedBoldTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 319, col: 37, offset: 13461},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 319, col: 37, offset: 13461},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 319, col: 50, offset: 13474},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 319, col: 50, offset: 13474},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 319, col: 54, offset: 13478},
										expr: &litMatcher{
											pos:        position{line: 319, col: 54, offset: 13478},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 319, col: 60, offset: 13484},
							val:        "*",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 319, col: 64, offset: 13488},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 319, col: 73, offset: 13497},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 319, col: 92, offset: 13516},
							val:        "*",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedBoldTextDoublePunctuation",
			pos:  position{line: 323, col: 1, offset: 13622},
			expr: &actionExpr{
				pos: position{line: 323, col: 37, offset: 13658},
				run: (*parser).callonEscapedBoldTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 323, col: 37, offset: 13658},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 323, col: 37, offset: 13658},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 323, col: 50, offset: 13671},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 323, col: 50, offset: 13671},
										val:        "\\\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 323, col: 55, offset: 13676},
										expr: &litMatcher{
											pos:        position{line: 323, col: 55, offset: 13676},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 323, col: 61, offset: 13682},
							val:        "**",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 323, col: 66, offset: 13687},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 323, col: 75, offset: 13696},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 323, col: 94, offset: 13715},
							val:        "**",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedBoldTextUnbalancedPunctuation",
			pos:  position{line: 327, col: 1, offset: 13823},
			expr: &actionExpr{
				pos: position{line: 327, col: 42, offset: 13864},
				run: (*parser).callonEscapedBoldTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 327, col: 42, offset: 13864},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 327, col: 42, offset: 13864},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 327, col: 55, offset: 13877},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 327, col: 55, offset: 13877},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 327, col: 59, offset: 13881},
										expr: &litMatcher{
											pos:        position{line: 327, col: 59, offset: 13881},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 327, col: 65, offset: 13887},
							val:        "**",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 327, col: 70, offset: 13892},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 327, col: 79, offset: 13901},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 327, col: 98, offset: 13920},
							val:        "*",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ItalicText",
			pos:  position{line: 332, col: 1, offset: 14113},
			expr: &choiceExpr{
				pos: position{line: 332, col: 15, offset: 14127},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 332, col: 15, offset: 14127},
						name: "ItalicTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 332, col: 45, offset: 14157},
						name: "ItalicTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 332, col: 79, offset: 14191},
						name: "ItalicTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "ItalicTextSimplePunctuation",
			pos:  position{line: 334, col: 1, offset: 14220},
			expr: &actionExpr{
				pos: position{line: 334, col: 32, offset: 14251},
				run: (*parser).callonItalicTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 334, col: 32, offset: 14251},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 334, col: 32, offset: 14251},
							expr: &litMatcher{
								pos:        position{line: 334, col: 33, offset: 14252},
								val:        "\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 334, col: 37, offset: 14256},
							val:        "_",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 334, col: 41, offset: 14260},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 334, col: 50, offset: 14269},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 334, col: 69, offset: 14288},
							val:        "_",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ItalicTextDoublePunctuation",
			pos:  position{line: 338, col: 1, offset: 14367},
			expr: &actionExpr{
				pos: position{line: 338, col: 32, offset: 14398},
				run: (*parser).callonItalicTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 338, col: 32, offset: 14398},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 338, col: 32, offset: 14398},
							expr: &litMatcher{
								pos:        position{line: 338, col: 33, offset: 14399},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 338, col: 38, offset: 14404},
							val:        "__",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 338, col: 43, offset: 14409},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 338, col: 52, offset: 14418},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 338, col: 71, offset: 14437},
							val:        "__",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ItalicTextUnbalancedPunctuation",
			pos:  position{line: 342, col: 1, offset: 14517},
			expr: &actionExpr{
				pos: position{line: 342, col: 36, offset: 14552},
				run: (*parser).callonItalicTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 342, col: 36, offset: 14552},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 342, col: 36, offset: 14552},
							expr: &litMatcher{
								pos:        position{line: 342, col: 37, offset: 14553},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 342, col: 42, offset: 14558},
							val:        "__",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 342, col: 47, offset: 14563},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 342, col: 56, offset: 14572},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 342, col: 75, offset: 14591},
							val:        "_",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedItalicText",
			pos:  position{line: 347, col: 1, offset: 14757},
			expr: &choiceExpr{
				pos: position{line: 347, col: 22, offset: 14778},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 347, col: 22, offset: 14778},
						name: "EscapedItalicTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 347, col: 59, offset: 14815},
						name: "EscapedItalicTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 347, col: 100, offset: 14856},
						name: "EscapedItalicTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "EscapedItalicTextSimplePunctuation",
			pos:  position{line: 349, col: 1, offset: 14938},
			expr: &actionExpr{
				pos: position{line: 349, col: 39, offset: 14976},
				run: (*parser).callonEscapedItalicTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 349, col: 39, offset: 14976},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 349, col: 39, offset: 14976},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 349, col: 52, offset: 14989},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 349, col: 52, offset: 14989},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 349, col: 56, offset: 14993},
										expr: &litMatcher{
											pos:        position{line: 349, col: 56, offset: 14993},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 349, col: 62, offset: 14999},
							val:        "_",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 349, col: 66, offset: 15003},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 349, col: 75, offset: 15012},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 349, col: 94, offset: 15031},
							val:        "_",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedItalicTextDoublePunctuation",
			pos:  position{line: 353, col: 1, offset: 15137},
			expr: &actionExpr{
				pos: position{line: 353, col: 39, offset: 15175},
				run: (*parser).callonEscapedItalicTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 353, col: 39, offset: 15175},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 353, col: 39, offset: 15175},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 353, col: 52, offset: 15188},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 353, col: 52, offset: 15188},
										val:        "\\\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 353, col: 57, offset: 15193},
										expr: &litMatcher{
											pos:        position{line: 353, col: 57, offset: 15193},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 353, col: 63, offset: 15199},
							val:        "__",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 353, col: 68, offset: 15204},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 353, col: 77, offset: 15213},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 353, col: 96, offset: 15232},
							val:        "__",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedItalicTextUnbalancedPunctuation",
			pos:  position{line: 357, col: 1, offset: 15340},
			expr: &actionExpr{
				pos: position{line: 357, col: 44, offset: 15383},
				run: (*parser).callonEscapedItalicTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 357, col: 44, offset: 15383},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 357, col: 44, offset: 15383},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 357, col: 57, offset: 15396},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 357, col: 57, offset: 15396},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 357, col: 61, offset: 15400},
										expr: &litMatcher{
											pos:        position{line: 357, col: 61, offset: 15400},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 357, col: 67, offset: 15406},
							val:        "__",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 357, col: 72, offset: 15411},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 357, col: 81, offset: 15420},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 357, col: 100, offset: 15439},
							val:        "_",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "MonospaceText",
			pos:  position{line: 362, col: 1, offset: 15632},
			expr: &choiceExpr{
				pos: position{line: 362, col: 18, offset: 15649},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 362, col: 18, offset: 15649},
						name: "MonospaceTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 362, col: 51, offset: 15682},
						name: "MonospaceTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 362, col: 88, offset: 15719},
						name: "MonospaceTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "MonospaceTextSimplePunctuation",
			pos:  position{line: 364, col: 1, offset: 15751},
			expr: &actionExpr{
				pos: position{line: 364, col: 35, offset: 15785},
				run: (*parser).callonMonospaceTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 364, col: 35, offset: 15785},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 364, col: 35, offset: 15785},
							expr: &litMatcher{
								pos:        position{line: 364, col: 36, offset: 15786},
								val:        "\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 364, col: 40, offset: 15790},
							val:        "`",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 364, col: 44, offset: 15794},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 364, col: 53, offset: 15803},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 364, col: 72, offset: 15822},
							val:        "`",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "MonospaceTextDoublePunctuation",
			pos:  position{line: 368, col: 1, offset: 15904},
			expr: &actionExpr{
				pos: position{line: 368, col: 35, offset: 15938},
				run: (*parser).callonMonospaceTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 368, col: 35, offset: 15938},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 368, col: 35, offset: 15938},
							expr: &litMatcher{
								pos:        position{line: 368, col: 36, offset: 15939},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 368, col: 41, offset: 15944},
							val:        "``",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 368, col: 46, offset: 15949},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 368, col: 55, offset: 15958},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 368, col: 74, offset: 15977},
							val:        "``",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "MonospaceTextUnbalancedPunctuation",
			pos:  position{line: 372, col: 1, offset: 16060},
			expr: &actionExpr{
				pos: position{line: 372, col: 39, offset: 16098},
				run: (*parser).callonMonospaceTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 372, col: 39, offset: 16098},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 372, col: 39, offset: 16098},
							expr: &litMatcher{
								pos:        position{line: 372, col: 40, offset: 16099},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 372, col: 45, offset: 16104},
							val:        "``",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 372, col: 50, offset: 16109},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 372, col: 59, offset: 16118},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 372, col: 78, offset: 16137},
							val:        "`",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedMonospaceText",
			pos:  position{line: 377, col: 1, offset: 16306},
			expr: &choiceExpr{
				pos: position{line: 377, col: 25, offset: 16330},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 377, col: 25, offset: 16330},
						name: "EscapedMonospaceTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 377, col: 65, offset: 16370},
						name: "EscapedMonospaceTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 377, col: 109, offset: 16414},
						name: "EscapedMonospaceTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "EscapedMonospaceTextSimplePunctuation",
			pos:  position{line: 379, col: 1, offset: 16499},
			expr: &actionExpr{
				pos: position{line: 379, col: 42, offset: 16540},
				run: (*parser).callonEscapedMonospaceTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 379, col: 42, offset: 16540},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 379, col: 42, offset: 16540},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 379, col: 55, offset: 16553},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 379, col: 55, offset: 16553},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 379, col: 59, offset: 16557},
										expr: &litMatcher{
											pos:        position{line: 379, col: 59, offset: 16557},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 379, col: 65, offset: 16563},
							val:        "`",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 379, col: 69, offset: 16567},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 379, col: 78, offset: 16576},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 379, col: 97, offset: 16595},
							val:        "`",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedMonospaceTextDoublePunctuation",
			pos:  position{line: 383, col: 1, offset: 16701},
			expr: &actionExpr{
				pos: position{line: 383, col: 42, offset: 16742},
				run: (*parser).callonEscapedMonospaceTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 383, col: 42, offset: 16742},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 383, col: 42, offset: 16742},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 383, col: 55, offset: 16755},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 383, col: 55, offset: 16755},
										val:        "\\\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 383, col: 60, offset: 16760},
										expr: &litMatcher{
											pos:        position{line: 383, col: 60, offset: 16760},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 383, col: 66, offset: 16766},
							val:        "``",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 383, col: 71, offset: 16771},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 383, col: 80, offset: 16780},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 383, col: 99, offset: 16799},
							val:        "``",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedMonospaceTextUnbalancedPunctuation",
			pos:  position{line: 387, col: 1, offset: 16907},
			expr: &actionExpr{
				pos: position{line: 387, col: 47, offset: 16953},
				run: (*parser).callonEscapedMonospaceTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 387, col: 47, offset: 16953},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 387, col: 47, offset: 16953},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 387, col: 60, offset: 16966},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 387, col: 60, offset: 16966},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 387, col: 64, offset: 16970},
										expr: &litMatcher{
											pos:        position{line: 387, col: 64, offset: 16970},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 387, col: 70, offset: 16976},
							val:        "``",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 387, col: 75, offset: 16981},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 387, col: 84, offset: 16990},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 387, col: 103, offset: 17009},
							val:        "`",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "QuotedTextContent",
			pos:  position{line: 392, col: 1, offset: 17202},
			expr: &seqExpr{
				pos: position{line: 392, col: 22, offset: 17223},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 392, col: 22, offset: 17223},
						name: "QuotedTextContentElement",
					},
					&zeroOrMoreExpr{
						pos: position{line: 392, col: 47, offset: 17248},
						expr: &seqExpr{
							pos: position{line: 392, col: 48, offset: 17249},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 392, col: 48, offset: 17249},
									expr: &ruleRefExpr{
										pos:  position{line: 392, col: 48, offset: 17249},
										name: "WS",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 392, col: 52, offset: 17253},
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
			pos:  position{line: 394, col: 1, offset: 17281},
			expr: &choiceExpr{
				pos: position{line: 394, col: 29, offset: 17309},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 394, col: 29, offset: 17309},
						name: "QuotedText",
					},
					&ruleRefExpr{
						pos:  position{line: 394, col: 42, offset: 17322},
						name: "QuotedTextCharacters",
					},
					&ruleRefExpr{
						pos:  position{line: 394, col: 65, offset: 17345},
						name: "CharactersWithQuotePunctuation",
					},
				},
			},
		},
		{
			name: "QuotedTextCharacters",
			pos:  position{line: 396, col: 1, offset: 17480},
			expr: &oneOrMoreExpr{
				pos: position{line: 396, col: 25, offset: 17504},
				expr: &seqExpr{
					pos: position{line: 396, col: 26, offset: 17505},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 396, col: 26, offset: 17505},
							expr: &ruleRefExpr{
								pos:  position{line: 396, col: 27, offset: 17506},
								name: "NEWLINE",
							},
						},
						&notExpr{
							pos: position{line: 396, col: 35, offset: 17514},
							expr: &ruleRefExpr{
								pos:  position{line: 396, col: 36, offset: 17515},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 396, col: 39, offset: 17518},
							expr: &litMatcher{
								pos:        position{line: 396, col: 40, offset: 17519},
								val:        "*",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 396, col: 44, offset: 17523},
							expr: &litMatcher{
								pos:        position{line: 396, col: 45, offset: 17524},
								val:        "_",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 396, col: 49, offset: 17528},
							expr: &litMatcher{
								pos:        position{line: 396, col: 50, offset: 17529},
								val:        "`",
								ignoreCase: false,
							},
						},
						&anyMatcher{
							line: 396, col: 54, offset: 17533,
						},
					},
				},
			},
		},
		{
			name: "CharactersWithQuotePunctuation",
			pos:  position{line: 398, col: 1, offset: 17576},
			expr: &actionExpr{
				pos: position{line: 398, col: 35, offset: 17610},
				run: (*parser).callonCharactersWithQuotePunctuation1,
				expr: &oneOrMoreExpr{
					pos: position{line: 398, col: 35, offset: 17610},
					expr: &seqExpr{
						pos: position{line: 398, col: 36, offset: 17611},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 398, col: 36, offset: 17611},
								expr: &ruleRefExpr{
									pos:  position{line: 398, col: 37, offset: 17612},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 398, col: 45, offset: 17620},
								expr: &ruleRefExpr{
									pos:  position{line: 398, col: 46, offset: 17621},
									name: "WS",
								},
							},
							&anyMatcher{
								line: 398, col: 50, offset: 17625,
							},
						},
					},
				},
			},
		},
		{
			name: "UnbalancedQuotePunctuation",
			pos:  position{line: 403, col: 1, offset: 17870},
			expr: &choiceExpr{
				pos: position{line: 403, col: 31, offset: 17900},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 403, col: 31, offset: 17900},
						val:        "*",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 403, col: 37, offset: 17906},
						val:        "_",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 403, col: 43, offset: 17912},
						val:        "`",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "Passthrough",
			pos:  position{line: 408, col: 1, offset: 18024},
			expr: &choiceExpr{
				pos: position{line: 408, col: 16, offset: 18039},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 408, col: 16, offset: 18039},
						name: "TriplePlusPassthrough",
					},
					&ruleRefExpr{
						pos:  position{line: 408, col: 40, offset: 18063},
						name: "SinglePlusPassthrough",
					},
					&ruleRefExpr{
						pos:  position{line: 408, col: 64, offset: 18087},
						name: "PassthroughMacro",
					},
				},
			},
		},
		{
			name: "SinglePlusPassthrough",
			pos:  position{line: 410, col: 1, offset: 18105},
			expr: &actionExpr{
				pos: position{line: 410, col: 26, offset: 18130},
				run: (*parser).callonSinglePlusPassthrough1,
				expr: &seqExpr{
					pos: position{line: 410, col: 26, offset: 18130},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 410, col: 26, offset: 18130},
							val:        "+",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 410, col: 30, offset: 18134},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 410, col: 38, offset: 18142},
								expr: &seqExpr{
									pos: position{line: 410, col: 39, offset: 18143},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 410, col: 39, offset: 18143},
											expr: &ruleRefExpr{
												pos:  position{line: 410, col: 40, offset: 18144},
												name: "NEWLINE",
											},
										},
										&notExpr{
											pos: position{line: 410, col: 48, offset: 18152},
											expr: &litMatcher{
												pos:        position{line: 410, col: 49, offset: 18153},
												val:        "+",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 410, col: 53, offset: 18157,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 410, col: 57, offset: 18161},
							val:        "+",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TriplePlusPassthrough",
			pos:  position{line: 414, col: 1, offset: 18256},
			expr: &actionExpr{
				pos: position{line: 414, col: 26, offset: 18281},
				run: (*parser).callonTriplePlusPassthrough1,
				expr: &seqExpr{
					pos: position{line: 414, col: 26, offset: 18281},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 414, col: 26, offset: 18281},
							val:        "+++",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 414, col: 32, offset: 18287},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 414, col: 40, offset: 18295},
								expr: &seqExpr{
									pos: position{line: 414, col: 41, offset: 18296},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 414, col: 41, offset: 18296},
											expr: &litMatcher{
												pos:        position{line: 414, col: 42, offset: 18297},
												val:        "+++",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 414, col: 48, offset: 18303,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 414, col: 52, offset: 18307},
							val:        "+++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "PassthroughMacro",
			pos:  position{line: 418, col: 1, offset: 18404},
			expr: &choiceExpr{
				pos: position{line: 418, col: 21, offset: 18424},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 418, col: 21, offset: 18424},
						name: "SimplePassthroughMacro",
					},
					&ruleRefExpr{
						pos:  position{line: 418, col: 46, offset: 18449},
						name: "PassthroughWithQuotedText",
					},
				},
			},
		},
		{
			name: "SimplePassthroughMacro",
			pos:  position{line: 420, col: 1, offset: 18476},
			expr: &actionExpr{
				pos: position{line: 420, col: 27, offset: 18502},
				run: (*parser).callonSimplePassthroughMacro1,
				expr: &seqExpr{
					pos: position{line: 420, col: 27, offset: 18502},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 420, col: 27, offset: 18502},
							val:        "pass:[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 420, col: 36, offset: 18511},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 420, col: 44, offset: 18519},
								expr: &ruleRefExpr{
									pos:  position{line: 420, col: 45, offset: 18520},
									name: "PassthroughMacroCharacter",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 420, col: 73, offset: 18548},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "PassthroughWithQuotedText",
			pos:  position{line: 424, col: 1, offset: 18638},
			expr: &actionExpr{
				pos: position{line: 424, col: 30, offset: 18667},
				run: (*parser).callonPassthroughWithQuotedText1,
				expr: &seqExpr{
					pos: position{line: 424, col: 30, offset: 18667},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 424, col: 30, offset: 18667},
							val:        "pass:q[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 424, col: 40, offset: 18677},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 424, col: 48, offset: 18685},
								expr: &choiceExpr{
									pos: position{line: 424, col: 49, offset: 18686},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 424, col: 49, offset: 18686},
											name: "QuotedText",
										},
										&ruleRefExpr{
											pos:  position{line: 424, col: 62, offset: 18699},
											name: "PassthroughMacroCharacter",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 424, col: 90, offset: 18727},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "PassthroughMacroCharacter",
			pos:  position{line: 428, col: 1, offset: 18817},
			expr: &seqExpr{
				pos: position{line: 428, col: 31, offset: 18847},
				exprs: []interface{}{
					&notExpr{
						pos: position{line: 428, col: 31, offset: 18847},
						expr: &litMatcher{
							pos:        position{line: 428, col: 32, offset: 18848},
							val:        "]",
							ignoreCase: false,
						},
					},
					&anyMatcher{
						line: 428, col: 36, offset: 18852,
					},
				},
			},
		},
		{
			name: "CrossReference",
			pos:  position{line: 433, col: 1, offset: 18968},
			expr: &actionExpr{
				pos: position{line: 433, col: 19, offset: 18986},
				run: (*parser).callonCrossReference1,
				expr: &seqExpr{
					pos: position{line: 433, col: 19, offset: 18986},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 433, col: 19, offset: 18986},
							val:        "<<",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 433, col: 24, offset: 18991},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 433, col: 28, offset: 18995},
								name: "ID",
							},
						},
						&litMatcher{
							pos:        position{line: 433, col: 32, offset: 18999},
							val:        ">>",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Link",
			pos:  position{line: 440, col: 1, offset: 19158},
			expr: &choiceExpr{
				pos: position{line: 440, col: 9, offset: 19166},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 440, col: 9, offset: 19166},
						name: "RelativeLink",
					},
					&ruleRefExpr{
						pos:  position{line: 440, col: 24, offset: 19181},
						name: "ExternalLink",
					},
				},
			},
		},
		{
			name: "ExternalLink",
			pos:  position{line: 442, col: 1, offset: 19196},
			expr: &actionExpr{
				pos: position{line: 442, col: 17, offset: 19212},
				run: (*parser).callonExternalLink1,
				expr: &seqExpr{
					pos: position{line: 442, col: 17, offset: 19212},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 442, col: 17, offset: 19212},
							label: "url",
							expr: &seqExpr{
								pos: position{line: 442, col: 22, offset: 19217},
								exprs: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 442, col: 22, offset: 19217},
										name: "URL_SCHEME",
									},
									&ruleRefExpr{
										pos:  position{line: 442, col: 33, offset: 19228},
										name: "URL",
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 442, col: 38, offset: 19233},
							label: "text",
							expr: &zeroOrOneExpr{
								pos: position{line: 442, col: 43, offset: 19238},
								expr: &seqExpr{
									pos: position{line: 442, col: 44, offset: 19239},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 442, col: 44, offset: 19239},
											val:        "[",
											ignoreCase: false,
										},
										&zeroOrMoreExpr{
											pos: position{line: 442, col: 48, offset: 19243},
											expr: &ruleRefExpr{
												pos:  position{line: 442, col: 49, offset: 19244},
												name: "URL_TEXT",
											},
										},
										&litMatcher{
											pos:        position{line: 442, col: 60, offset: 19255},
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
			pos:  position{line: 449, col: 1, offset: 19416},
			expr: &actionExpr{
				pos: position{line: 449, col: 17, offset: 19432},
				run: (*parser).callonRelativeLink1,
				expr: &seqExpr{
					pos: position{line: 449, col: 17, offset: 19432},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 449, col: 17, offset: 19432},
							val:        "link:",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 449, col: 25, offset: 19440},
							label: "url",
							expr: &seqExpr{
								pos: position{line: 449, col: 30, offset: 19445},
								exprs: []interface{}{
									&zeroOrOneExpr{
										pos: position{line: 449, col: 30, offset: 19445},
										expr: &ruleRefExpr{
											pos:  position{line: 449, col: 30, offset: 19445},
											name: "URL_SCHEME",
										},
									},
									&ruleRefExpr{
										pos:  position{line: 449, col: 42, offset: 19457},
										name: "URL",
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 449, col: 47, offset: 19462},
							label: "text",
							expr: &seqExpr{
								pos: position{line: 449, col: 53, offset: 19468},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 449, col: 53, offset: 19468},
										val:        "[",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 449, col: 57, offset: 19472},
										expr: &ruleRefExpr{
											pos:  position{line: 449, col: 58, offset: 19473},
											name: "URL_TEXT",
										},
									},
									&litMatcher{
										pos:        position{line: 449, col: 69, offset: 19484},
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
			pos:  position{line: 459, col: 1, offset: 19746},
			expr: &actionExpr{
				pos: position{line: 459, col: 15, offset: 19760},
				run: (*parser).callonBlockImage1,
				expr: &seqExpr{
					pos: position{line: 459, col: 15, offset: 19760},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 459, col: 15, offset: 19760},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 459, col: 26, offset: 19771},
								expr: &ruleRefExpr{
									pos:  position{line: 459, col: 27, offset: 19772},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 459, col: 46, offset: 19791},
							label: "image",
							expr: &ruleRefExpr{
								pos:  position{line: 459, col: 52, offset: 19797},
								name: "BlockImageMacro",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 459, col: 69, offset: 19814},
							expr: &ruleRefExpr{
								pos:  position{line: 459, col: 69, offset: 19814},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 459, col: 73, offset: 19818},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "BlockImageMacro",
			pos:  position{line: 464, col: 1, offset: 19979},
			expr: &actionExpr{
				pos: position{line: 464, col: 20, offset: 19998},
				run: (*parser).callonBlockImageMacro1,
				expr: &seqExpr{
					pos: position{line: 464, col: 20, offset: 19998},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 464, col: 20, offset: 19998},
							val:        "image::",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 464, col: 30, offset: 20008},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 464, col: 36, offset: 20014},
								name: "URL",
							},
						},
						&litMatcher{
							pos:        position{line: 464, col: 41, offset: 20019},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 464, col: 45, offset: 20023},
							label: "attributes",
							expr: &zeroOrOneExpr{
								pos: position{line: 464, col: 57, offset: 20035},
								expr: &ruleRefExpr{
									pos:  position{line: 464, col: 57, offset: 20035},
									name: "URL_TEXT",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 464, col: 68, offset: 20046},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "InlineImage",
			pos:  position{line: 468, col: 1, offset: 20113},
			expr: &actionExpr{
				pos: position{line: 468, col: 16, offset: 20128},
				run: (*parser).callonInlineImage1,
				expr: &labeledExpr{
					pos:   position{line: 468, col: 16, offset: 20128},
					label: "image",
					expr: &ruleRefExpr{
						pos:  position{line: 468, col: 22, offset: 20134},
						name: "InlineImageMacro",
					},
				},
			},
		},
		{
			name: "InlineImageMacro",
			pos:  position{line: 473, col: 1, offset: 20281},
			expr: &actionExpr{
				pos: position{line: 473, col: 21, offset: 20301},
				run: (*parser).callonInlineImageMacro1,
				expr: &seqExpr{
					pos: position{line: 473, col: 21, offset: 20301},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 473, col: 21, offset: 20301},
							val:        "image:",
							ignoreCase: false,
						},
						&notExpr{
							pos: position{line: 473, col: 30, offset: 20310},
							expr: &litMatcher{
								pos:        position{line: 473, col: 31, offset: 20311},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 473, col: 35, offset: 20315},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 473, col: 41, offset: 20321},
								name: "URL",
							},
						},
						&litMatcher{
							pos:        position{line: 473, col: 46, offset: 20326},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 473, col: 50, offset: 20330},
							label: "attributes",
							expr: &zeroOrOneExpr{
								pos: position{line: 473, col: 62, offset: 20342},
								expr: &ruleRefExpr{
									pos:  position{line: 473, col: 62, offset: 20342},
									name: "URL_TEXT",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 473, col: 73, offset: 20353},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "DelimitedBlock",
			pos:  position{line: 480, col: 1, offset: 20683},
			expr: &choiceExpr{
				pos: position{line: 480, col: 19, offset: 20701},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 480, col: 19, offset: 20701},
						name: "FencedBlock",
					},
					&ruleRefExpr{
						pos:  position{line: 480, col: 33, offset: 20715},
						name: "ListingBlock",
					},
				},
			},
		},
		{
			name: "BlockDelimiter",
			pos:  position{line: 482, col: 1, offset: 20730},
			expr: &choiceExpr{
				pos: position{line: 482, col: 19, offset: 20748},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 482, col: 19, offset: 20748},
						name: "FencedBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 482, col: 42, offset: 20771},
						name: "ListingBlockDelimiter",
					},
				},
			},
		},
		{
			name: "FencedBlockDelimiter",
			pos:  position{line: 484, col: 1, offset: 20794},
			expr: &litMatcher{
				pos:        position{line: 484, col: 25, offset: 20818},
				val:        "```",
				ignoreCase: false,
			},
		},
		{
			name: "FencedBlock",
			pos:  position{line: 486, col: 1, offset: 20825},
			expr: &actionExpr{
				pos: position{line: 486, col: 16, offset: 20840},
				run: (*parser).callonFencedBlock1,
				expr: &seqExpr{
					pos: position{line: 486, col: 16, offset: 20840},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 486, col: 16, offset: 20840},
							name: "FencedBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 486, col: 37, offset: 20861},
							expr: &ruleRefExpr{
								pos:  position{line: 486, col: 37, offset: 20861},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 486, col: 41, offset: 20865},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 486, col: 49, offset: 20873},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 486, col: 58, offset: 20882},
								name: "FencedBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 486, col: 78, offset: 20902},
							name: "FencedBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 486, col: 99, offset: 20923},
							expr: &ruleRefExpr{
								pos:  position{line: 486, col: 99, offset: 20923},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 486, col: 103, offset: 20927},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "FencedBlockContent",
			pos:  position{line: 490, col: 1, offset: 21015},
			expr: &labeledExpr{
				pos:   position{line: 490, col: 23, offset: 21037},
				label: "content",
				expr: &zeroOrMoreExpr{
					pos: position{line: 490, col: 31, offset: 21045},
					expr: &seqExpr{
						pos: position{line: 490, col: 32, offset: 21046},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 490, col: 32, offset: 21046},
								expr: &ruleRefExpr{
									pos:  position{line: 490, col: 33, offset: 21047},
									name: "FencedBlockDelimiter",
								},
							},
							&anyMatcher{
								line: 490, col: 54, offset: 21068,
							},
						},
					},
				},
			},
		},
		{
			name: "ListingBlockDelimiter",
			pos:  position{line: 492, col: 1, offset: 21074},
			expr: &litMatcher{
				pos:        position{line: 492, col: 26, offset: 21099},
				val:        "----",
				ignoreCase: false,
			},
		},
		{
			name: "ListingBlock",
			pos:  position{line: 494, col: 1, offset: 21107},
			expr: &actionExpr{
				pos: position{line: 494, col: 17, offset: 21123},
				run: (*parser).callonListingBlock1,
				expr: &seqExpr{
					pos: position{line: 494, col: 17, offset: 21123},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 494, col: 17, offset: 21123},
							name: "ListingBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 494, col: 39, offset: 21145},
							expr: &ruleRefExpr{
								pos:  position{line: 494, col: 39, offset: 21145},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 494, col: 43, offset: 21149},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 494, col: 51, offset: 21157},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 494, col: 60, offset: 21166},
								name: "ListingBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 494, col: 81, offset: 21187},
							name: "ListingBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 494, col: 103, offset: 21209},
							expr: &ruleRefExpr{
								pos:  position{line: 494, col: 103, offset: 21209},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 494, col: 107, offset: 21213},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ListingBlockContent",
			pos:  position{line: 498, col: 1, offset: 21302},
			expr: &labeledExpr{
				pos:   position{line: 498, col: 24, offset: 21325},
				label: "content",
				expr: &zeroOrMoreExpr{
					pos: position{line: 498, col: 32, offset: 21333},
					expr: &seqExpr{
						pos: position{line: 498, col: 33, offset: 21334},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 498, col: 33, offset: 21334},
								expr: &ruleRefExpr{
									pos:  position{line: 498, col: 34, offset: 21335},
									name: "ListingBlockDelimiter",
								},
							},
							&anyMatcher{
								line: 498, col: 56, offset: 21357,
							},
						},
					},
				},
			},
		},
		{
			name: "LiteralBlock",
			pos:  position{line: 503, col: 1, offset: 21630},
			expr: &choiceExpr{
				pos: position{line: 503, col: 17, offset: 21646},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 503, col: 17, offset: 21646},
						name: "ParagraphWithSpaces",
					},
					&ruleRefExpr{
						pos:  position{line: 503, col: 39, offset: 21668},
						name: "ParagraphWithLiteralBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 503, col: 76, offset: 21705},
						name: "ParagraphWithLiteralAttribute",
					},
				},
			},
		},
		{
			name: "ParagraphWithSpaces",
			pos:  position{line: 506, col: 1, offset: 21800},
			expr: &actionExpr{
				pos: position{line: 506, col: 24, offset: 21823},
				run: (*parser).callonParagraphWithSpaces1,
				expr: &seqExpr{
					pos: position{line: 506, col: 24, offset: 21823},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 506, col: 24, offset: 21823},
							label: "spaces",
							expr: &oneOrMoreExpr{
								pos: position{line: 506, col: 32, offset: 21831},
								expr: &ruleRefExpr{
									pos:  position{line: 506, col: 32, offset: 21831},
									name: "WS",
								},
							},
						},
						&notExpr{
							pos: position{line: 506, col: 37, offset: 21836},
							expr: &ruleRefExpr{
								pos:  position{line: 506, col: 38, offset: 21837},
								name: "NEWLINE",
							},
						},
						&labeledExpr{
							pos:   position{line: 506, col: 46, offset: 21845},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 506, col: 55, offset: 21854},
								name: "LiteralBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 506, col: 76, offset: 21875},
							name: "EndOfLiteralBlock",
						},
					},
				},
			},
		},
		{
			name: "LiteralBlockContent",
			pos:  position{line: 511, col: 1, offset: 22056},
			expr: &actionExpr{
				pos: position{line: 511, col: 24, offset: 22079},
				run: (*parser).callonLiteralBlockContent1,
				expr: &labeledExpr{
					pos:   position{line: 511, col: 24, offset: 22079},
					label: "content",
					expr: &oneOrMoreExpr{
						pos: position{line: 511, col: 32, offset: 22087},
						expr: &seqExpr{
							pos: position{line: 511, col: 33, offset: 22088},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 511, col: 33, offset: 22088},
									expr: &seqExpr{
										pos: position{line: 511, col: 35, offset: 22090},
										exprs: []interface{}{
											&ruleRefExpr{
												pos:  position{line: 511, col: 35, offset: 22090},
												name: "NEWLINE",
											},
											&ruleRefExpr{
												pos:  position{line: 511, col: 43, offset: 22098},
												name: "BlankLine",
											},
										},
									},
								},
								&anyMatcher{
									line: 511, col: 54, offset: 22109,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "EndOfLiteralBlock",
			pos:  position{line: 516, col: 1, offset: 22194},
			expr: &choiceExpr{
				pos: position{line: 516, col: 22, offset: 22215},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 516, col: 22, offset: 22215},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 516, col: 22, offset: 22215},
								name: "NEWLINE",
							},
							&ruleRefExpr{
								pos:  position{line: 516, col: 30, offset: 22223},
								name: "BlankLine",
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 516, col: 42, offset: 22235},
						name: "NEWLINE",
					},
					&ruleRefExpr{
						pos:  position{line: 516, col: 52, offset: 22245},
						name: "EOF",
					},
				},
			},
		},
		{
			name: "ParagraphWithLiteralBlockDelimiter",
			pos:  position{line: 519, col: 1, offset: 22305},
			expr: &actionExpr{
				pos: position{line: 519, col: 39, offset: 22343},
				run: (*parser).callonParagraphWithLiteralBlockDelimiter1,
				expr: &seqExpr{
					pos: position{line: 519, col: 39, offset: 22343},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 519, col: 39, offset: 22343},
							name: "LiteralBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 519, col: 61, offset: 22365},
							expr: &ruleRefExpr{
								pos:  position{line: 519, col: 61, offset: 22365},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 519, col: 65, offset: 22369},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 519, col: 73, offset: 22377},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 519, col: 81, offset: 22385},
								expr: &seqExpr{
									pos: position{line: 519, col: 82, offset: 22386},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 519, col: 82, offset: 22386},
											expr: &ruleRefExpr{
												pos:  position{line: 519, col: 83, offset: 22387},
												name: "LiteralBlockDelimiter",
											},
										},
										&anyMatcher{
											line: 519, col: 105, offset: 22409,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 519, col: 109, offset: 22413},
							name: "LiteralBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 519, col: 131, offset: 22435},
							expr: &ruleRefExpr{
								pos:  position{line: 519, col: 131, offset: 22435},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 519, col: 135, offset: 22439},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "LiteralBlockDelimiter",
			pos:  position{line: 523, col: 1, offset: 22523},
			expr: &litMatcher{
				pos:        position{line: 523, col: 26, offset: 22548},
				val:        "....",
				ignoreCase: false,
			},
		},
		{
			name: "ParagraphWithLiteralAttribute",
			pos:  position{line: 526, col: 1, offset: 22610},
			expr: &actionExpr{
				pos: position{line: 526, col: 34, offset: 22643},
				run: (*parser).callonParagraphWithLiteralAttribute1,
				expr: &seqExpr{
					pos: position{line: 526, col: 34, offset: 22643},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 526, col: 34, offset: 22643},
							val:        "[literal]",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 526, col: 46, offset: 22655},
							expr: &ruleRefExpr{
								pos:  position{line: 526, col: 46, offset: 22655},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 526, col: 50, offset: 22659},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 526, col: 58, offset: 22667},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 526, col: 67, offset: 22676},
								name: "LiteralBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 526, col: 88, offset: 22697},
							name: "EndOfLiteralBlock",
						},
					},
				},
			},
		},
		{
			name: "ElementAttribute",
			pos:  position{line: 533, col: 1, offset: 22909},
			expr: &choiceExpr{
				pos: position{line: 533, col: 21, offset: 22929},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 533, col: 21, offset: 22929},
						name: "ElementLink",
					},
					&ruleRefExpr{
						pos:  position{line: 533, col: 35, offset: 22943},
						name: "ElementID",
					},
					&ruleRefExpr{
						pos:  position{line: 533, col: 47, offset: 22955},
						name: "ElementTitle",
					},
					&ruleRefExpr{
						pos:  position{line: 533, col: 62, offset: 22970},
						name: "InvalidElementAttribute",
					},
				},
			},
		},
		{
			name: "ElementLink",
			pos:  position{line: 536, col: 1, offset: 23050},
			expr: &actionExpr{
				pos: position{line: 536, col: 16, offset: 23065},
				run: (*parser).callonElementLink1,
				expr: &seqExpr{
					pos: position{line: 536, col: 16, offset: 23065},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 536, col: 16, offset: 23065},
							val:        "[link=",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 536, col: 25, offset: 23074},
							expr: &ruleRefExpr{
								pos:  position{line: 536, col: 25, offset: 23074},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 536, col: 29, offset: 23078},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 536, col: 34, offset: 23083},
								name: "URL",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 536, col: 38, offset: 23087},
							expr: &ruleRefExpr{
								pos:  position{line: 536, col: 38, offset: 23087},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 536, col: 42, offset: 23091},
							val:        "]",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 536, col: 46, offset: 23095},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ElementID",
			pos:  position{line: 540, col: 1, offset: 23151},
			expr: &choiceExpr{
				pos: position{line: 540, col: 14, offset: 23164},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 540, col: 14, offset: 23164},
						run: (*parser).callonElementID2,
						expr: &seqExpr{
							pos: position{line: 540, col: 14, offset: 23164},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 540, col: 14, offset: 23164},
									label: "id",
									expr: &ruleRefExpr{
										pos:  position{line: 540, col: 18, offset: 23168},
										name: "InlineElementID",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 540, col: 35, offset: 23185},
									name: "EOL",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 542, col: 5, offset: 23214},
						run: (*parser).callonElementID7,
						expr: &seqExpr{
							pos: position{line: 542, col: 5, offset: 23214},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 542, col: 5, offset: 23214},
									val:        "[#",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 542, col: 10, offset: 23219},
									label: "id",
									expr: &ruleRefExpr{
										pos:  position{line: 542, col: 14, offset: 23223},
										name: "ID",
									},
								},
								&litMatcher{
									pos:        position{line: 542, col: 18, offset: 23227},
									val:        "]",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 542, col: 22, offset: 23231},
									name: "EOL",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "InlineElementID",
			pos:  position{line: 546, col: 1, offset: 23283},
			expr: &actionExpr{
				pos: position{line: 546, col: 20, offset: 23302},
				run: (*parser).callonInlineElementID1,
				expr: &seqExpr{
					pos: position{line: 546, col: 20, offset: 23302},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 546, col: 20, offset: 23302},
							val:        "[[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 546, col: 25, offset: 23307},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 546, col: 29, offset: 23311},
								name: "ID",
							},
						},
						&litMatcher{
							pos:        position{line: 546, col: 33, offset: 23315},
							val:        "]]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ElementTitle",
			pos:  position{line: 552, col: 1, offset: 23510},
			expr: &actionExpr{
				pos: position{line: 552, col: 17, offset: 23526},
				run: (*parser).callonElementTitle1,
				expr: &seqExpr{
					pos: position{line: 552, col: 17, offset: 23526},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 552, col: 17, offset: 23526},
							val:        ".",
							ignoreCase: false,
						},
						&notExpr{
							pos: position{line: 552, col: 21, offset: 23530},
							expr: &litMatcher{
								pos:        position{line: 552, col: 22, offset: 23531},
								val:        ".",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 552, col: 26, offset: 23535},
							expr: &ruleRefExpr{
								pos:  position{line: 552, col: 27, offset: 23536},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 552, col: 30, offset: 23539},
							label: "title",
							expr: &oneOrMoreExpr{
								pos: position{line: 552, col: 36, offset: 23545},
								expr: &seqExpr{
									pos: position{line: 552, col: 37, offset: 23546},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 552, col: 37, offset: 23546},
											expr: &ruleRefExpr{
												pos:  position{line: 552, col: 38, offset: 23547},
												name: "NEWLINE",
											},
										},
										&anyMatcher{
											line: 552, col: 46, offset: 23555,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 552, col: 50, offset: 23559},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "InvalidElementAttribute",
			pos:  position{line: 556, col: 1, offset: 23624},
			expr: &actionExpr{
				pos: position{line: 556, col: 28, offset: 23651},
				run: (*parser).callonInvalidElementAttribute1,
				expr: &seqExpr{
					pos: position{line: 556, col: 28, offset: 23651},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 556, col: 28, offset: 23651},
							val:        "[",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 556, col: 32, offset: 23655},
							expr: &ruleRefExpr{
								pos:  position{line: 556, col: 32, offset: 23655},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 556, col: 36, offset: 23659},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 556, col: 44, offset: 23667},
								expr: &seqExpr{
									pos: position{line: 556, col: 45, offset: 23668},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 556, col: 45, offset: 23668},
											expr: &litMatcher{
												pos:        position{line: 556, col: 46, offset: 23669},
												val:        "]",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 556, col: 50, offset: 23673,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 556, col: 54, offset: 23677},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "BlankLine",
			pos:  position{line: 563, col: 1, offset: 23843},
			expr: &actionExpr{
				pos: position{line: 563, col: 14, offset: 23856},
				run: (*parser).callonBlankLine1,
				expr: &seqExpr{
					pos: position{line: 563, col: 14, offset: 23856},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 563, col: 14, offset: 23856},
							expr: &ruleRefExpr{
								pos:  position{line: 563, col: 15, offset: 23857},
								name: "EOF",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 563, col: 19, offset: 23861},
							expr: &ruleRefExpr{
								pos:  position{line: 563, col: 19, offset: 23861},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 563, col: 23, offset: 23865},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "Characters",
			pos:  position{line: 570, col: 1, offset: 24012},
			expr: &actionExpr{
				pos: position{line: 570, col: 15, offset: 24026},
				run: (*parser).callonCharacters1,
				expr: &oneOrMoreExpr{
					pos: position{line: 570, col: 15, offset: 24026},
					expr: &seqExpr{
						pos: position{line: 570, col: 16, offset: 24027},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 570, col: 16, offset: 24027},
								expr: &ruleRefExpr{
									pos:  position{line: 570, col: 17, offset: 24028},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 570, col: 25, offset: 24036},
								expr: &ruleRefExpr{
									pos:  position{line: 570, col: 26, offset: 24037},
									name: "WS",
								},
							},
							&anyMatcher{
								line: 570, col: 29, offset: 24040,
							},
						},
					},
				},
			},
		},
		{
			name: "URL",
			pos:  position{line: 575, col: 1, offset: 24081},
			expr: &actionExpr{
				pos: position{line: 575, col: 8, offset: 24088},
				run: (*parser).callonURL1,
				expr: &oneOrMoreExpr{
					pos: position{line: 575, col: 8, offset: 24088},
					expr: &seqExpr{
						pos: position{line: 575, col: 9, offset: 24089},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 575, col: 9, offset: 24089},
								expr: &ruleRefExpr{
									pos:  position{line: 575, col: 10, offset: 24090},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 575, col: 18, offset: 24098},
								expr: &ruleRefExpr{
									pos:  position{line: 575, col: 19, offset: 24099},
									name: "WS",
								},
							},
							&notExpr{
								pos: position{line: 575, col: 22, offset: 24102},
								expr: &litMatcher{
									pos:        position{line: 575, col: 23, offset: 24103},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 575, col: 27, offset: 24107},
								expr: &litMatcher{
									pos:        position{line: 575, col: 28, offset: 24108},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 575, col: 32, offset: 24112,
							},
						},
					},
				},
			},
		},
		{
			name: "ID",
			pos:  position{line: 579, col: 1, offset: 24152},
			expr: &actionExpr{
				pos: position{line: 579, col: 7, offset: 24158},
				run: (*parser).callonID1,
				expr: &oneOrMoreExpr{
					pos: position{line: 579, col: 7, offset: 24158},
					expr: &seqExpr{
						pos: position{line: 579, col: 8, offset: 24159},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 579, col: 8, offset: 24159},
								expr: &ruleRefExpr{
									pos:  position{line: 579, col: 9, offset: 24160},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 579, col: 17, offset: 24168},
								expr: &ruleRefExpr{
									pos:  position{line: 579, col: 18, offset: 24169},
									name: "WS",
								},
							},
							&notExpr{
								pos: position{line: 579, col: 21, offset: 24172},
								expr: &litMatcher{
									pos:        position{line: 579, col: 22, offset: 24173},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 579, col: 26, offset: 24177},
								expr: &litMatcher{
									pos:        position{line: 579, col: 27, offset: 24178},
									val:        "]",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 579, col: 31, offset: 24182},
								expr: &litMatcher{
									pos:        position{line: 579, col: 32, offset: 24183},
									val:        "<<",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 579, col: 37, offset: 24188},
								expr: &litMatcher{
									pos:        position{line: 579, col: 38, offset: 24189},
									val:        ">>",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 579, col: 42, offset: 24193,
							},
						},
					},
				},
			},
		},
		{
			name: "URL_TEXT",
			pos:  position{line: 583, col: 1, offset: 24233},
			expr: &actionExpr{
				pos: position{line: 583, col: 13, offset: 24245},
				run: (*parser).callonURL_TEXT1,
				expr: &oneOrMoreExpr{
					pos: position{line: 583, col: 13, offset: 24245},
					expr: &seqExpr{
						pos: position{line: 583, col: 14, offset: 24246},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 583, col: 14, offset: 24246},
								expr: &ruleRefExpr{
									pos:  position{line: 583, col: 15, offset: 24247},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 583, col: 23, offset: 24255},
								expr: &litMatcher{
									pos:        position{line: 583, col: 24, offset: 24256},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 583, col: 28, offset: 24260},
								expr: &litMatcher{
									pos:        position{line: 583, col: 29, offset: 24261},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 583, col: 33, offset: 24265,
							},
						},
					},
				},
			},
		},
		{
			name: "URL_SCHEME",
			pos:  position{line: 587, col: 1, offset: 24305},
			expr: &choiceExpr{
				pos: position{line: 587, col: 15, offset: 24319},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 587, col: 15, offset: 24319},
						val:        "http://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 587, col: 27, offset: 24331},
						val:        "https://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 587, col: 40, offset: 24344},
						val:        "ftp://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 587, col: 51, offset: 24355},
						val:        "irc://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 587, col: 62, offset: 24366},
						val:        "mailto:",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "DIGIT",
			pos:  position{line: 589, col: 1, offset: 24377},
			expr: &charClassMatcher{
				pos:        position{line: 589, col: 10, offset: 24386},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NEWLINE",
			pos:  position{line: 591, col: 1, offset: 24393},
			expr: &choiceExpr{
				pos: position{line: 591, col: 12, offset: 24404},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 591, col: 12, offset: 24404},
						val:        "\r\n",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 591, col: 21, offset: 24413},
						val:        "\r",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 591, col: 28, offset: 24420},
						val:        "\n",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "WS",
			pos:  position{line: 593, col: 1, offset: 24426},
			expr: &choiceExpr{
				pos: position{line: 593, col: 7, offset: 24432},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 593, col: 7, offset: 24432},
						val:        " ",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 593, col: 13, offset: 24438},
						run: (*parser).callonWS3,
						expr: &litMatcher{
							pos:        position{line: 593, col: 13, offset: 24438},
							val:        "\t",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 597, col: 1, offset: 24483},
			expr: &notExpr{
				pos: position{line: 597, col: 8, offset: 24490},
				expr: &anyMatcher{
					line: 597, col: 9, offset: 24491,
				},
			},
		},
		{
			name: "EOL",
			pos:  position{line: 599, col: 1, offset: 24494},
			expr: &choiceExpr{
				pos: position{line: 599, col: 8, offset: 24501},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 599, col: 8, offset: 24501},
						name: "NEWLINE",
					},
					&ruleRefExpr{
						pos:  position{line: 599, col: 18, offset: 24511},
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

	return types.NewSectionTitle(content.(*types.InlineContent), append(attributes.([]interface{}), id))
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

func (c *current) onSection1Title1(attributes, level, content, id interface{}) (interface{}, error) {

	return types.NewSectionTitle(content.(*types.InlineContent), append(attributes.([]interface{}), id))
}

func (p *parser) callonSection1Title1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection1Title1(stack["attributes"], stack["level"], stack["content"], stack["id"])
}

func (c *current) onSection2Title1(attributes, level, content, id interface{}) (interface{}, error) {
	fmt.Println("New Section 2...")
	return types.NewSectionTitle(content.(*types.InlineContent), append(attributes.([]interface{}), id))
}

func (p *parser) callonSection2Title1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection2Title1(stack["attributes"], stack["level"], stack["content"], stack["id"])
}

func (c *current) onSection3Title1(attributes, level, content, id interface{}) (interface{}, error) {
	return types.NewSectionTitle(content.(*types.InlineContent), append(attributes.([]interface{}), id))
}

func (p *parser) callonSection3Title1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection3Title1(stack["attributes"], stack["level"], stack["content"], stack["id"])
}

func (c *current) onSection4Title1(attributes, level, content, id interface{}) (interface{}, error) {
	return types.NewSectionTitle(content.(*types.InlineContent), append(attributes.([]interface{}), id))
}

func (p *parser) callonSection4Title1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection4Title1(stack["attributes"], stack["level"], stack["content"], stack["id"])
}

func (c *current) onSection5Title1(attributes, level, content, id interface{}) (interface{}, error) {
	return types.NewSectionTitle(content.(*types.InlineContent), append(attributes.([]interface{}), id))
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

func (c *current) onInlineContentWithTrailingSpaces1(elements interface{}) (interface{}, error) {
	// absorbs heading and trailing spaces
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

func (c *current) onElementID2(id interface{}) (interface{}, error) {
	return id, nil
}

func (p *parser) callonElementID2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onElementID2(stack["id"])
}

func (c *current) onElementID7(id interface{}) (interface{}, error) {
	return types.NewElementID(id.(string))
}

func (p *parser) callonElementID7() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onElementID7(stack["id"])
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
