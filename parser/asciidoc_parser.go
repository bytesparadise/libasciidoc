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
			pos:  position{line: 188, col: 1, offset: 7568},
			expr: &actionExpr{
				pos: position{line: 188, col: 18, offset: 7585},
				run: (*parser).callonSection3Title1,
				expr: &seqExpr{
					pos: position{line: 188, col: 18, offset: 7585},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 188, col: 18, offset: 7585},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 188, col: 29, offset: 7596},
								expr: &ruleRefExpr{
									pos:  position{line: 188, col: 30, offset: 7597},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 188, col: 49, offset: 7616},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 188, col: 56, offset: 7623},
								val:        "====",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 188, col: 64, offset: 7631},
							expr: &ruleRefExpr{
								pos:  position{line: 188, col: 64, offset: 7631},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 188, col: 68, offset: 7635},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 188, col: 77, offset: 7644},
								name: "InlineContent",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 188, col: 92, offset: 7659},
							expr: &ruleRefExpr{
								pos:  position{line: 188, col: 92, offset: 7659},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 188, col: 96, offset: 7663},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 188, col: 99, offset: 7666},
								expr: &ruleRefExpr{
									pos:  position{line: 188, col: 100, offset: 7667},
									name: "InlineElementID",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 188, col: 118, offset: 7685},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 188, col: 123, offset: 7690},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 188, col: 123, offset: 7690},
									expr: &ruleRefExpr{
										pos:  position{line: 188, col: 123, offset: 7690},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 188, col: 136, offset: 7703},
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
			pos:  position{line: 192, col: 1, offset: 7818},
			expr: &actionExpr{
				pos: position{line: 192, col: 18, offset: 7835},
				run: (*parser).callonSection4Title1,
				expr: &seqExpr{
					pos: position{line: 192, col: 18, offset: 7835},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 192, col: 18, offset: 7835},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 192, col: 29, offset: 7846},
								expr: &ruleRefExpr{
									pos:  position{line: 192, col: 30, offset: 7847},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 192, col: 49, offset: 7866},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 192, col: 56, offset: 7873},
								val:        "=====",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 192, col: 65, offset: 7882},
							expr: &ruleRefExpr{
								pos:  position{line: 192, col: 65, offset: 7882},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 192, col: 69, offset: 7886},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 192, col: 78, offset: 7895},
								name: "InlineContent",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 192, col: 93, offset: 7910},
							expr: &ruleRefExpr{
								pos:  position{line: 192, col: 93, offset: 7910},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 192, col: 97, offset: 7914},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 192, col: 100, offset: 7917},
								expr: &ruleRefExpr{
									pos:  position{line: 192, col: 101, offset: 7918},
									name: "InlineElementID",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 192, col: 119, offset: 7936},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 192, col: 124, offset: 7941},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 192, col: 124, offset: 7941},
									expr: &ruleRefExpr{
										pos:  position{line: 192, col: 124, offset: 7941},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 192, col: 137, offset: 7954},
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
			pos:  position{line: 196, col: 1, offset: 8069},
			expr: &actionExpr{
				pos: position{line: 196, col: 18, offset: 8086},
				run: (*parser).callonSection5Title1,
				expr: &seqExpr{
					pos: position{line: 196, col: 18, offset: 8086},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 196, col: 18, offset: 8086},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 196, col: 29, offset: 8097},
								expr: &ruleRefExpr{
									pos:  position{line: 196, col: 30, offset: 8098},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 196, col: 49, offset: 8117},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 196, col: 56, offset: 8124},
								val:        "======",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 196, col: 66, offset: 8134},
							expr: &ruleRefExpr{
								pos:  position{line: 196, col: 66, offset: 8134},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 196, col: 70, offset: 8138},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 196, col: 79, offset: 8147},
								name: "InlineContent",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 196, col: 94, offset: 8162},
							expr: &ruleRefExpr{
								pos:  position{line: 196, col: 94, offset: 8162},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 196, col: 98, offset: 8166},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 196, col: 101, offset: 8169},
								expr: &ruleRefExpr{
									pos:  position{line: 196, col: 102, offset: 8170},
									name: "InlineElementID",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 196, col: 120, offset: 8188},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 196, col: 125, offset: 8193},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 196, col: 125, offset: 8193},
									expr: &ruleRefExpr{
										pos:  position{line: 196, col: 125, offset: 8193},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 196, col: 138, offset: 8206},
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
			pos:  position{line: 203, col: 1, offset: 8422},
			expr: &actionExpr{
				pos: position{line: 203, col: 9, offset: 8430},
				run: (*parser).callonList1,
				expr: &seqExpr{
					pos: position{line: 203, col: 9, offset: 8430},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 203, col: 9, offset: 8430},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 203, col: 20, offset: 8441},
								expr: &ruleRefExpr{
									pos:  position{line: 203, col: 21, offset: 8442},
									name: "ListAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 205, col: 5, offset: 8531},
							label: "elements",
							expr: &oneOrMoreExpr{
								pos: position{line: 205, col: 14, offset: 8540},
								expr: &choiceExpr{
									pos: position{line: 205, col: 15, offset: 8541},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 205, col: 15, offset: 8541},
											name: "UnorderedListItem",
										},
										&ruleRefExpr{
											pos:  position{line: 205, col: 35, offset: 8561},
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
			pos:  position{line: 209, col: 1, offset: 8663},
			expr: &actionExpr{
				pos: position{line: 209, col: 18, offset: 8680},
				run: (*parser).callonListAttribute1,
				expr: &seqExpr{
					pos: position{line: 209, col: 18, offset: 8680},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 209, col: 18, offset: 8680},
							label: "attribute",
							expr: &choiceExpr{
								pos: position{line: 209, col: 29, offset: 8691},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 209, col: 29, offset: 8691},
										name: "HorizontalLayout",
									},
									&ruleRefExpr{
										pos:  position{line: 209, col: 48, offset: 8710},
										name: "ListID",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 209, col: 56, offset: 8718},
							name: "NEWLINE",
						},
					},
				},
			},
		},
		{
			name: "ListID",
			pos:  position{line: 213, col: 1, offset: 8757},
			expr: &actionExpr{
				pos: position{line: 213, col: 11, offset: 8767},
				run: (*parser).callonListID1,
				expr: &seqExpr{
					pos: position{line: 213, col: 11, offset: 8767},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 213, col: 11, offset: 8767},
							val:        "[#",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 213, col: 16, offset: 8772},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 213, col: 20, offset: 8776},
								name: "ID",
							},
						},
						&litMatcher{
							pos:        position{line: 213, col: 24, offset: 8780},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "HorizontalLayout",
			pos:  position{line: 217, col: 1, offset: 8846},
			expr: &actionExpr{
				pos: position{line: 217, col: 21, offset: 8866},
				run: (*parser).callonHorizontalLayout1,
				expr: &litMatcher{
					pos:        position{line: 217, col: 21, offset: 8866},
					val:        "[horizontal]",
					ignoreCase: false,
				},
			},
		},
		{
			name: "ListParagraph",
			pos:  position{line: 221, col: 1, offset: 8949},
			expr: &actionExpr{
				pos: position{line: 221, col: 19, offset: 8967},
				run: (*parser).callonListParagraph1,
				expr: &labeledExpr{
					pos:   position{line: 221, col: 19, offset: 8967},
					label: "lines",
					expr: &oneOrMoreExpr{
						pos: position{line: 221, col: 25, offset: 8973},
						expr: &seqExpr{
							pos: position{line: 221, col: 26, offset: 8974},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 221, col: 26, offset: 8974},
									expr: &ruleRefExpr{
										pos:  position{line: 221, col: 28, offset: 8976},
										name: "ListItemContinuation",
									},
								},
								&notExpr{
									pos: position{line: 221, col: 50, offset: 8998},
									expr: &ruleRefExpr{
										pos:  position{line: 221, col: 52, offset: 9000},
										name: "UnorderedListItemPrefix",
									},
								},
								&notExpr{
									pos: position{line: 221, col: 77, offset: 9025},
									expr: &seqExpr{
										pos: position{line: 221, col: 79, offset: 9027},
										exprs: []interface{}{
											&ruleRefExpr{
												pos:  position{line: 221, col: 79, offset: 9027},
												name: "LabeledListItemTerm",
											},
											&ruleRefExpr{
												pos:  position{line: 221, col: 99, offset: 9047},
												name: "LabeledListItemSeparator",
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 221, col: 125, offset: 9073},
									name: "InlineContentWithTrailingSpaces",
								},
								&ruleRefExpr{
									pos:  position{line: 221, col: 157, offset: 9105},
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
			pos:  position{line: 225, col: 1, offset: 9174},
			expr: &actionExpr{
				pos: position{line: 225, col: 25, offset: 9198},
				run: (*parser).callonListItemContinuation1,
				expr: &seqExpr{
					pos: position{line: 225, col: 25, offset: 9198},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 225, col: 25, offset: 9198},
							val:        "+",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 225, col: 29, offset: 9202},
							expr: &ruleRefExpr{
								pos:  position{line: 225, col: 29, offset: 9202},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 225, col: 33, offset: 9206},
							name: "NEWLINE",
						},
					},
				},
			},
		},
		{
			name: "ContinuedBlockElement",
			pos:  position{line: 229, col: 1, offset: 9262},
			expr: &actionExpr{
				pos: position{line: 229, col: 26, offset: 9287},
				run: (*parser).callonContinuedBlockElement1,
				expr: &seqExpr{
					pos: position{line: 229, col: 26, offset: 9287},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 229, col: 26, offset: 9287},
							name: "ListItemContinuation",
						},
						&labeledExpr{
							pos:   position{line: 229, col: 47, offset: 9308},
							label: "element",
							expr: &ruleRefExpr{
								pos:  position{line: 229, col: 55, offset: 9316},
								name: "BlockElement",
							},
						},
					},
				},
			},
		},
		{
			name: "UnorderedListItem",
			pos:  position{line: 236, col: 1, offset: 9469},
			expr: &actionExpr{
				pos: position{line: 236, col: 22, offset: 9490},
				run: (*parser).callonUnorderedListItem1,
				expr: &seqExpr{
					pos: position{line: 236, col: 22, offset: 9490},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 236, col: 22, offset: 9490},
							label: "level",
							expr: &ruleRefExpr{
								pos:  position{line: 236, col: 29, offset: 9497},
								name: "UnorderedListItemPrefix",
							},
						},
						&labeledExpr{
							pos:   position{line: 236, col: 54, offset: 9522},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 236, col: 63, offset: 9531},
								name: "UnorderedListItemContent",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 236, col: 89, offset: 9557},
							expr: &ruleRefExpr{
								pos:  position{line: 236, col: 89, offset: 9557},
								name: "BlankLine",
							},
						},
					},
				},
			},
		},
		{
			name: "UnorderedListItemPrefix",
			pos:  position{line: 240, col: 1, offset: 9648},
			expr: &actionExpr{
				pos: position{line: 240, col: 28, offset: 9675},
				run: (*parser).callonUnorderedListItemPrefix1,
				expr: &seqExpr{
					pos: position{line: 240, col: 28, offset: 9675},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 240, col: 28, offset: 9675},
							expr: &ruleRefExpr{
								pos:  position{line: 240, col: 28, offset: 9675},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 240, col: 32, offset: 9679},
							label: "level",
							expr: &choiceExpr{
								pos: position{line: 240, col: 39, offset: 9686},
								alternatives: []interface{}{
									&oneOrMoreExpr{
										pos: position{line: 240, col: 39, offset: 9686},
										expr: &litMatcher{
											pos:        position{line: 240, col: 39, offset: 9686},
											val:        "*",
											ignoreCase: false,
										},
									},
									&litMatcher{
										pos:        position{line: 240, col: 46, offset: 9693},
										val:        "-",
										ignoreCase: false,
									},
								},
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 240, col: 51, offset: 9698},
							expr: &ruleRefExpr{
								pos:  position{line: 240, col: 51, offset: 9698},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "UnorderedListItemContent",
			pos:  position{line: 244, col: 1, offset: 9796},
			expr: &actionExpr{
				pos: position{line: 244, col: 29, offset: 9824},
				run: (*parser).callonUnorderedListItemContent1,
				expr: &labeledExpr{
					pos:   position{line: 244, col: 29, offset: 9824},
					label: "elements",
					expr: &seqExpr{
						pos: position{line: 244, col: 39, offset: 9834},
						exprs: []interface{}{
							&oneOrMoreExpr{
								pos: position{line: 244, col: 39, offset: 9834},
								expr: &ruleRefExpr{
									pos:  position{line: 244, col: 39, offset: 9834},
									name: "ListParagraph",
								},
							},
							&zeroOrMoreExpr{
								pos: position{line: 244, col: 54, offset: 9849},
								expr: &ruleRefExpr{
									pos:  position{line: 244, col: 54, offset: 9849},
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
			pos:  position{line: 253, col: 1, offset: 10170},
			expr: &choiceExpr{
				pos: position{line: 253, col: 20, offset: 10189},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 253, col: 20, offset: 10189},
						name: "LabeledListItemWithDescription",
					},
					&ruleRefExpr{
						pos:  position{line: 253, col: 53, offset: 10222},
						name: "LabeledListItemWithTermAlone",
					},
				},
			},
		},
		{
			name: "LabeledListItemWithTermAlone",
			pos:  position{line: 255, col: 1, offset: 10252},
			expr: &actionExpr{
				pos: position{line: 255, col: 33, offset: 10284},
				run: (*parser).callonLabeledListItemWithTermAlone1,
				expr: &seqExpr{
					pos: position{line: 255, col: 33, offset: 10284},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 255, col: 33, offset: 10284},
							label: "term",
							expr: &ruleRefExpr{
								pos:  position{line: 255, col: 39, offset: 10290},
								name: "LabeledListItemTerm",
							},
						},
						&litMatcher{
							pos:        position{line: 255, col: 61, offset: 10312},
							val:        "::",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 255, col: 66, offset: 10317},
							expr: &ruleRefExpr{
								pos:  position{line: 255, col: 66, offset: 10317},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 255, col: 70, offset: 10321},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "LabeledListItemTerm",
			pos:  position{line: 259, col: 1, offset: 10458},
			expr: &actionExpr{
				pos: position{line: 259, col: 24, offset: 10481},
				run: (*parser).callonLabeledListItemTerm1,
				expr: &labeledExpr{
					pos:   position{line: 259, col: 24, offset: 10481},
					label: "term",
					expr: &zeroOrMoreExpr{
						pos: position{line: 259, col: 29, offset: 10486},
						expr: &seqExpr{
							pos: position{line: 259, col: 30, offset: 10487},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 259, col: 30, offset: 10487},
									expr: &ruleRefExpr{
										pos:  position{line: 259, col: 31, offset: 10488},
										name: "NEWLINE",
									},
								},
								&notExpr{
									pos: position{line: 259, col: 39, offset: 10496},
									expr: &litMatcher{
										pos:        position{line: 259, col: 40, offset: 10497},
										val:        "::",
										ignoreCase: false,
									},
								},
								&anyMatcher{
									line: 259, col: 45, offset: 10502,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "LabeledListItemSeparator",
			pos:  position{line: 264, col: 1, offset: 10593},
			expr: &seqExpr{
				pos: position{line: 264, col: 30, offset: 10622},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 264, col: 30, offset: 10622},
						val:        "::",
						ignoreCase: false,
					},
					&oneOrMoreExpr{
						pos: position{line: 264, col: 35, offset: 10627},
						expr: &choiceExpr{
							pos: position{line: 264, col: 36, offset: 10628},
							alternatives: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 264, col: 36, offset: 10628},
									name: "WS",
								},
								&ruleRefExpr{
									pos:  position{line: 264, col: 41, offset: 10633},
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
			pos:  position{line: 266, col: 1, offset: 10644},
			expr: &actionExpr{
				pos: position{line: 266, col: 35, offset: 10678},
				run: (*parser).callonLabeledListItemWithDescription1,
				expr: &seqExpr{
					pos: position{line: 266, col: 35, offset: 10678},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 266, col: 35, offset: 10678},
							label: "term",
							expr: &ruleRefExpr{
								pos:  position{line: 266, col: 41, offset: 10684},
								name: "LabeledListItemTerm",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 266, col: 62, offset: 10705},
							name: "LabeledListItemSeparator",
						},
						&labeledExpr{
							pos:   position{line: 266, col: 87, offset: 10730},
							label: "description",
							expr: &ruleRefExpr{
								pos:  position{line: 266, col: 100, offset: 10743},
								name: "LabeledListItemDescription",
							},
						},
					},
				},
			},
		},
		{
			name: "LabeledListItemDescription",
			pos:  position{line: 270, col: 1, offset: 10868},
			expr: &actionExpr{
				pos: position{line: 270, col: 31, offset: 10898},
				run: (*parser).callonLabeledListItemDescription1,
				expr: &labeledExpr{
					pos:   position{line: 270, col: 31, offset: 10898},
					label: "elements",
					expr: &zeroOrMoreExpr{
						pos: position{line: 270, col: 40, offset: 10907},
						expr: &choiceExpr{
							pos: position{line: 270, col: 41, offset: 10908},
							alternatives: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 270, col: 41, offset: 10908},
									name: "ListParagraph",
								},
								&ruleRefExpr{
									pos:  position{line: 270, col: 57, offset: 10924},
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
			pos:  position{line: 279, col: 1, offset: 11274},
			expr: &actionExpr{
				pos: position{line: 279, col: 14, offset: 11287},
				run: (*parser).callonParagraph1,
				expr: &seqExpr{
					pos: position{line: 279, col: 14, offset: 11287},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 279, col: 14, offset: 11287},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 279, col: 25, offset: 11298},
								expr: &ruleRefExpr{
									pos:  position{line: 279, col: 26, offset: 11299},
									name: "ElementAttribute",
								},
							},
						},
						&notExpr{
							pos: position{line: 279, col: 45, offset: 11318},
							expr: &seqExpr{
								pos: position{line: 279, col: 47, offset: 11320},
								exprs: []interface{}{
									&oneOrMoreExpr{
										pos: position{line: 279, col: 47, offset: 11320},
										expr: &litMatcher{
											pos:        position{line: 279, col: 47, offset: 11320},
											val:        "=",
											ignoreCase: false,
										},
									},
									&oneOrMoreExpr{
										pos: position{line: 279, col: 52, offset: 11325},
										expr: &ruleRefExpr{
											pos:  position{line: 279, col: 52, offset: 11325},
											name: "WS",
										},
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 279, col: 57, offset: 11330},
							label: "lines",
							expr: &oneOrMoreExpr{
								pos: position{line: 279, col: 63, offset: 11336},
								expr: &seqExpr{
									pos: position{line: 279, col: 64, offset: 11337},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 279, col: 64, offset: 11337},
											name: "InlineContentWithTrailingSpaces",
										},
										&ruleRefExpr{
											pos:  position{line: 279, col: 96, offset: 11369},
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
			pos:  position{line: 285, col: 1, offset: 11659},
			expr: &actionExpr{
				pos: position{line: 285, col: 36, offset: 11694},
				run: (*parser).callonInlineContentWithTrailingSpaces1,
				expr: &seqExpr{
					pos: position{line: 285, col: 36, offset: 11694},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 285, col: 36, offset: 11694},
							expr: &ruleRefExpr{
								pos:  position{line: 285, col: 37, offset: 11695},
								name: "BlockDelimiter",
							},
						},
						&labeledExpr{
							pos:   position{line: 285, col: 52, offset: 11710},
							label: "elements",
							expr: &oneOrMoreExpr{
								pos: position{line: 285, col: 61, offset: 11719},
								expr: &seqExpr{
									pos: position{line: 285, col: 62, offset: 11720},
									exprs: []interface{}{
										&zeroOrMoreExpr{
											pos: position{line: 285, col: 62, offset: 11720},
											expr: &ruleRefExpr{
												pos:  position{line: 285, col: 62, offset: 11720},
												name: "WS",
											},
										},
										&notExpr{
											pos: position{line: 285, col: 66, offset: 11724},
											expr: &ruleRefExpr{
												pos:  position{line: 285, col: 67, offset: 11725},
												name: "InlineElementID",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 285, col: 83, offset: 11741},
											name: "InlineElement",
										},
										&zeroOrMoreExpr{
											pos: position{line: 285, col: 97, offset: 11755},
											expr: &ruleRefExpr{
												pos:  position{line: 285, col: 97, offset: 11755},
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
			pos:  position{line: 289, col: 1, offset: 11867},
			expr: &actionExpr{
				pos: position{line: 289, col: 18, offset: 11884},
				run: (*parser).callonInlineContent1,
				expr: &seqExpr{
					pos: position{line: 289, col: 18, offset: 11884},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 289, col: 18, offset: 11884},
							expr: &ruleRefExpr{
								pos:  position{line: 289, col: 19, offset: 11885},
								name: "BlockDelimiter",
							},
						},
						&labeledExpr{
							pos:   position{line: 289, col: 34, offset: 11900},
							label: "elements",
							expr: &oneOrMoreExpr{
								pos: position{line: 289, col: 43, offset: 11909},
								expr: &seqExpr{
									pos: position{line: 289, col: 44, offset: 11910},
									exprs: []interface{}{
										&zeroOrMoreExpr{
											pos: position{line: 289, col: 44, offset: 11910},
											expr: &ruleRefExpr{
												pos:  position{line: 289, col: 44, offset: 11910},
												name: "WS",
											},
										},
										&notExpr{
											pos: position{line: 289, col: 48, offset: 11914},
											expr: &ruleRefExpr{
												pos:  position{line: 289, col: 49, offset: 11915},
												name: "InlineElementID",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 289, col: 65, offset: 11931},
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
			pos:  position{line: 293, col: 1, offset: 12053},
			expr: &choiceExpr{
				pos: position{line: 293, col: 19, offset: 12071},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 293, col: 19, offset: 12071},
						name: "CrossReference",
					},
					&ruleRefExpr{
						pos:  position{line: 293, col: 36, offset: 12088},
						name: "Passthrough",
					},
					&ruleRefExpr{
						pos:  position{line: 293, col: 50, offset: 12102},
						name: "InlineImage",
					},
					&ruleRefExpr{
						pos:  position{line: 293, col: 64, offset: 12116},
						name: "QuotedText",
					},
					&ruleRefExpr{
						pos:  position{line: 293, col: 77, offset: 12129},
						name: "Link",
					},
					&ruleRefExpr{
						pos:  position{line: 293, col: 84, offset: 12136},
						name: "DocumentAttributeSubstitution",
					},
					&ruleRefExpr{
						pos:  position{line: 293, col: 116, offset: 12168},
						name: "Characters",
					},
				},
			},
		},
		{
			name: "QuotedText",
			pos:  position{line: 298, col: 1, offset: 12420},
			expr: &choiceExpr{
				pos: position{line: 298, col: 15, offset: 12434},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 298, col: 15, offset: 12434},
						name: "BoldText",
					},
					&ruleRefExpr{
						pos:  position{line: 298, col: 26, offset: 12445},
						name: "ItalicText",
					},
					&ruleRefExpr{
						pos:  position{line: 298, col: 39, offset: 12458},
						name: "MonospaceText",
					},
					&ruleRefExpr{
						pos:  position{line: 299, col: 13, offset: 12486},
						name: "EscapedBoldText",
					},
					&ruleRefExpr{
						pos:  position{line: 299, col: 31, offset: 12504},
						name: "EscapedItalicText",
					},
					&ruleRefExpr{
						pos:  position{line: 299, col: 51, offset: 12524},
						name: "EscapedMonospaceText",
					},
				},
			},
		},
		{
			name: "BoldText",
			pos:  position{line: 301, col: 1, offset: 12546},
			expr: &choiceExpr{
				pos: position{line: 301, col: 13, offset: 12558},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 301, col: 13, offset: 12558},
						name: "BoldTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 301, col: 41, offset: 12586},
						name: "BoldTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 301, col: 73, offset: 12618},
						name: "BoldTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "BoldTextSimplePunctuation",
			pos:  position{line: 303, col: 1, offset: 12691},
			expr: &actionExpr{
				pos: position{line: 303, col: 30, offset: 12720},
				run: (*parser).callonBoldTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 303, col: 30, offset: 12720},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 303, col: 30, offset: 12720},
							expr: &litMatcher{
								pos:        position{line: 303, col: 31, offset: 12721},
								val:        "\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 303, col: 35, offset: 12725},
							val:        "*",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 303, col: 39, offset: 12729},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 303, col: 48, offset: 12738},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 303, col: 67, offset: 12757},
							val:        "*",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "BoldTextDoublePunctuation",
			pos:  position{line: 307, col: 1, offset: 12834},
			expr: &actionExpr{
				pos: position{line: 307, col: 30, offset: 12863},
				run: (*parser).callonBoldTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 307, col: 30, offset: 12863},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 307, col: 30, offset: 12863},
							expr: &litMatcher{
								pos:        position{line: 307, col: 31, offset: 12864},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 307, col: 36, offset: 12869},
							val:        "**",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 307, col: 41, offset: 12874},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 307, col: 50, offset: 12883},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 307, col: 69, offset: 12902},
							val:        "**",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "BoldTextUnbalancedPunctuation",
			pos:  position{line: 311, col: 1, offset: 12980},
			expr: &actionExpr{
				pos: position{line: 311, col: 34, offset: 13013},
				run: (*parser).callonBoldTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 311, col: 34, offset: 13013},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 311, col: 34, offset: 13013},
							expr: &litMatcher{
								pos:        position{line: 311, col: 35, offset: 13014},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 311, col: 40, offset: 13019},
							val:        "**",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 311, col: 45, offset: 13024},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 311, col: 54, offset: 13033},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 311, col: 73, offset: 13052},
							val:        "*",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedBoldText",
			pos:  position{line: 316, col: 1, offset: 13216},
			expr: &choiceExpr{
				pos: position{line: 316, col: 20, offset: 13235},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 316, col: 20, offset: 13235},
						name: "EscapedBoldTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 316, col: 55, offset: 13270},
						name: "EscapedBoldTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 316, col: 94, offset: 13309},
						name: "EscapedBoldTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "EscapedBoldTextSimplePunctuation",
			pos:  position{line: 318, col: 1, offset: 13389},
			expr: &actionExpr{
				pos: position{line: 318, col: 37, offset: 13425},
				run: (*parser).callonEscapedBoldTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 318, col: 37, offset: 13425},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 318, col: 37, offset: 13425},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 318, col: 50, offset: 13438},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 318, col: 50, offset: 13438},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 318, col: 54, offset: 13442},
										expr: &litMatcher{
											pos:        position{line: 318, col: 54, offset: 13442},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 318, col: 60, offset: 13448},
							val:        "*",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 318, col: 64, offset: 13452},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 318, col: 73, offset: 13461},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 318, col: 92, offset: 13480},
							val:        "*",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedBoldTextDoublePunctuation",
			pos:  position{line: 322, col: 1, offset: 13586},
			expr: &actionExpr{
				pos: position{line: 322, col: 37, offset: 13622},
				run: (*parser).callonEscapedBoldTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 322, col: 37, offset: 13622},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 322, col: 37, offset: 13622},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 322, col: 50, offset: 13635},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 322, col: 50, offset: 13635},
										val:        "\\\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 322, col: 55, offset: 13640},
										expr: &litMatcher{
											pos:        position{line: 322, col: 55, offset: 13640},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 322, col: 61, offset: 13646},
							val:        "**",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 322, col: 66, offset: 13651},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 322, col: 75, offset: 13660},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 322, col: 94, offset: 13679},
							val:        "**",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedBoldTextUnbalancedPunctuation",
			pos:  position{line: 326, col: 1, offset: 13787},
			expr: &actionExpr{
				pos: position{line: 326, col: 42, offset: 13828},
				run: (*parser).callonEscapedBoldTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 326, col: 42, offset: 13828},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 326, col: 42, offset: 13828},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 326, col: 55, offset: 13841},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 326, col: 55, offset: 13841},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 326, col: 59, offset: 13845},
										expr: &litMatcher{
											pos:        position{line: 326, col: 59, offset: 13845},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 326, col: 65, offset: 13851},
							val:        "**",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 326, col: 70, offset: 13856},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 326, col: 79, offset: 13865},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 326, col: 98, offset: 13884},
							val:        "*",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ItalicText",
			pos:  position{line: 331, col: 1, offset: 14077},
			expr: &choiceExpr{
				pos: position{line: 331, col: 15, offset: 14091},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 331, col: 15, offset: 14091},
						name: "ItalicTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 331, col: 45, offset: 14121},
						name: "ItalicTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 331, col: 79, offset: 14155},
						name: "ItalicTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "ItalicTextSimplePunctuation",
			pos:  position{line: 333, col: 1, offset: 14184},
			expr: &actionExpr{
				pos: position{line: 333, col: 32, offset: 14215},
				run: (*parser).callonItalicTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 333, col: 32, offset: 14215},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 333, col: 32, offset: 14215},
							expr: &litMatcher{
								pos:        position{line: 333, col: 33, offset: 14216},
								val:        "\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 333, col: 37, offset: 14220},
							val:        "_",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 333, col: 41, offset: 14224},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 333, col: 50, offset: 14233},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 333, col: 69, offset: 14252},
							val:        "_",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ItalicTextDoublePunctuation",
			pos:  position{line: 337, col: 1, offset: 14331},
			expr: &actionExpr{
				pos: position{line: 337, col: 32, offset: 14362},
				run: (*parser).callonItalicTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 337, col: 32, offset: 14362},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 337, col: 32, offset: 14362},
							expr: &litMatcher{
								pos:        position{line: 337, col: 33, offset: 14363},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 337, col: 38, offset: 14368},
							val:        "__",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 337, col: 43, offset: 14373},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 337, col: 52, offset: 14382},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 337, col: 71, offset: 14401},
							val:        "__",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ItalicTextUnbalancedPunctuation",
			pos:  position{line: 341, col: 1, offset: 14481},
			expr: &actionExpr{
				pos: position{line: 341, col: 36, offset: 14516},
				run: (*parser).callonItalicTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 341, col: 36, offset: 14516},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 341, col: 36, offset: 14516},
							expr: &litMatcher{
								pos:        position{line: 341, col: 37, offset: 14517},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 341, col: 42, offset: 14522},
							val:        "__",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 341, col: 47, offset: 14527},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 341, col: 56, offset: 14536},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 341, col: 75, offset: 14555},
							val:        "_",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedItalicText",
			pos:  position{line: 346, col: 1, offset: 14721},
			expr: &choiceExpr{
				pos: position{line: 346, col: 22, offset: 14742},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 346, col: 22, offset: 14742},
						name: "EscapedItalicTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 346, col: 59, offset: 14779},
						name: "EscapedItalicTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 346, col: 100, offset: 14820},
						name: "EscapedItalicTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "EscapedItalicTextSimplePunctuation",
			pos:  position{line: 348, col: 1, offset: 14902},
			expr: &actionExpr{
				pos: position{line: 348, col: 39, offset: 14940},
				run: (*parser).callonEscapedItalicTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 348, col: 39, offset: 14940},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 348, col: 39, offset: 14940},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 348, col: 52, offset: 14953},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 348, col: 52, offset: 14953},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 348, col: 56, offset: 14957},
										expr: &litMatcher{
											pos:        position{line: 348, col: 56, offset: 14957},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 348, col: 62, offset: 14963},
							val:        "_",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 348, col: 66, offset: 14967},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 348, col: 75, offset: 14976},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 348, col: 94, offset: 14995},
							val:        "_",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedItalicTextDoublePunctuation",
			pos:  position{line: 352, col: 1, offset: 15101},
			expr: &actionExpr{
				pos: position{line: 352, col: 39, offset: 15139},
				run: (*parser).callonEscapedItalicTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 352, col: 39, offset: 15139},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 352, col: 39, offset: 15139},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 352, col: 52, offset: 15152},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 352, col: 52, offset: 15152},
										val:        "\\\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 352, col: 57, offset: 15157},
										expr: &litMatcher{
											pos:        position{line: 352, col: 57, offset: 15157},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 352, col: 63, offset: 15163},
							val:        "__",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 352, col: 68, offset: 15168},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 352, col: 77, offset: 15177},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 352, col: 96, offset: 15196},
							val:        "__",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedItalicTextUnbalancedPunctuation",
			pos:  position{line: 356, col: 1, offset: 15304},
			expr: &actionExpr{
				pos: position{line: 356, col: 44, offset: 15347},
				run: (*parser).callonEscapedItalicTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 356, col: 44, offset: 15347},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 356, col: 44, offset: 15347},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 356, col: 57, offset: 15360},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 356, col: 57, offset: 15360},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 356, col: 61, offset: 15364},
										expr: &litMatcher{
											pos:        position{line: 356, col: 61, offset: 15364},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 356, col: 67, offset: 15370},
							val:        "__",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 356, col: 72, offset: 15375},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 356, col: 81, offset: 15384},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 356, col: 100, offset: 15403},
							val:        "_",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "MonospaceText",
			pos:  position{line: 361, col: 1, offset: 15596},
			expr: &choiceExpr{
				pos: position{line: 361, col: 18, offset: 15613},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 361, col: 18, offset: 15613},
						name: "MonospaceTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 361, col: 51, offset: 15646},
						name: "MonospaceTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 361, col: 88, offset: 15683},
						name: "MonospaceTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "MonospaceTextSimplePunctuation",
			pos:  position{line: 363, col: 1, offset: 15715},
			expr: &actionExpr{
				pos: position{line: 363, col: 35, offset: 15749},
				run: (*parser).callonMonospaceTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 363, col: 35, offset: 15749},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 363, col: 35, offset: 15749},
							expr: &litMatcher{
								pos:        position{line: 363, col: 36, offset: 15750},
								val:        "\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 363, col: 40, offset: 15754},
							val:        "`",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 363, col: 44, offset: 15758},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 363, col: 53, offset: 15767},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 363, col: 72, offset: 15786},
							val:        "`",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "MonospaceTextDoublePunctuation",
			pos:  position{line: 367, col: 1, offset: 15868},
			expr: &actionExpr{
				pos: position{line: 367, col: 35, offset: 15902},
				run: (*parser).callonMonospaceTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 367, col: 35, offset: 15902},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 367, col: 35, offset: 15902},
							expr: &litMatcher{
								pos:        position{line: 367, col: 36, offset: 15903},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 367, col: 41, offset: 15908},
							val:        "``",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 367, col: 46, offset: 15913},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 367, col: 55, offset: 15922},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 367, col: 74, offset: 15941},
							val:        "``",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "MonospaceTextUnbalancedPunctuation",
			pos:  position{line: 371, col: 1, offset: 16024},
			expr: &actionExpr{
				pos: position{line: 371, col: 39, offset: 16062},
				run: (*parser).callonMonospaceTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 371, col: 39, offset: 16062},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 371, col: 39, offset: 16062},
							expr: &litMatcher{
								pos:        position{line: 371, col: 40, offset: 16063},
								val:        "\\\\",
								ignoreCase: false,
							},
						},
						&litMatcher{
							pos:        position{line: 371, col: 45, offset: 16068},
							val:        "``",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 371, col: 50, offset: 16073},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 371, col: 59, offset: 16082},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 371, col: 78, offset: 16101},
							val:        "`",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedMonospaceText",
			pos:  position{line: 376, col: 1, offset: 16270},
			expr: &choiceExpr{
				pos: position{line: 376, col: 25, offset: 16294},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 376, col: 25, offset: 16294},
						name: "EscapedMonospaceTextDoublePunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 376, col: 65, offset: 16334},
						name: "EscapedMonospaceTextUnbalancedPunctuation",
					},
					&ruleRefExpr{
						pos:  position{line: 376, col: 109, offset: 16378},
						name: "EscapedMonospaceTextSimplePunctuation",
					},
				},
			},
		},
		{
			name: "EscapedMonospaceTextSimplePunctuation",
			pos:  position{line: 378, col: 1, offset: 16463},
			expr: &actionExpr{
				pos: position{line: 378, col: 42, offset: 16504},
				run: (*parser).callonEscapedMonospaceTextSimplePunctuation1,
				expr: &seqExpr{
					pos: position{line: 378, col: 42, offset: 16504},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 378, col: 42, offset: 16504},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 378, col: 55, offset: 16517},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 378, col: 55, offset: 16517},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 378, col: 59, offset: 16521},
										expr: &litMatcher{
											pos:        position{line: 378, col: 59, offset: 16521},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 378, col: 65, offset: 16527},
							val:        "`",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 378, col: 69, offset: 16531},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 378, col: 78, offset: 16540},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 378, col: 97, offset: 16559},
							val:        "`",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedMonospaceTextDoublePunctuation",
			pos:  position{line: 382, col: 1, offset: 16665},
			expr: &actionExpr{
				pos: position{line: 382, col: 42, offset: 16706},
				run: (*parser).callonEscapedMonospaceTextDoublePunctuation1,
				expr: &seqExpr{
					pos: position{line: 382, col: 42, offset: 16706},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 382, col: 42, offset: 16706},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 382, col: 55, offset: 16719},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 382, col: 55, offset: 16719},
										val:        "\\\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 382, col: 60, offset: 16724},
										expr: &litMatcher{
											pos:        position{line: 382, col: 60, offset: 16724},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 382, col: 66, offset: 16730},
							val:        "``",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 382, col: 71, offset: 16735},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 382, col: 80, offset: 16744},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 382, col: 99, offset: 16763},
							val:        "``",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EscapedMonospaceTextUnbalancedPunctuation",
			pos:  position{line: 386, col: 1, offset: 16871},
			expr: &actionExpr{
				pos: position{line: 386, col: 47, offset: 16917},
				run: (*parser).callonEscapedMonospaceTextUnbalancedPunctuation1,
				expr: &seqExpr{
					pos: position{line: 386, col: 47, offset: 16917},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 386, col: 47, offset: 16917},
							label: "backslashes",
							expr: &seqExpr{
								pos: position{line: 386, col: 60, offset: 16930},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 386, col: 60, offset: 16930},
										val:        "\\",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 386, col: 64, offset: 16934},
										expr: &litMatcher{
											pos:        position{line: 386, col: 64, offset: 16934},
											val:        "\\",
											ignoreCase: false,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 386, col: 70, offset: 16940},
							val:        "``",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 386, col: 75, offset: 16945},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 386, col: 84, offset: 16954},
								name: "QuotedTextContent",
							},
						},
						&litMatcher{
							pos:        position{line: 386, col: 103, offset: 16973},
							val:        "`",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "QuotedTextContent",
			pos:  position{line: 391, col: 1, offset: 17166},
			expr: &seqExpr{
				pos: position{line: 391, col: 22, offset: 17187},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 391, col: 22, offset: 17187},
						name: "QuotedTextContentElement",
					},
					&zeroOrMoreExpr{
						pos: position{line: 391, col: 47, offset: 17212},
						expr: &seqExpr{
							pos: position{line: 391, col: 48, offset: 17213},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 391, col: 48, offset: 17213},
									expr: &ruleRefExpr{
										pos:  position{line: 391, col: 48, offset: 17213},
										name: "WS",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 391, col: 52, offset: 17217},
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
			pos:  position{line: 393, col: 1, offset: 17245},
			expr: &choiceExpr{
				pos: position{line: 393, col: 29, offset: 17273},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 393, col: 29, offset: 17273},
						name: "QuotedText",
					},
					&ruleRefExpr{
						pos:  position{line: 393, col: 42, offset: 17286},
						name: "QuotedTextCharacters",
					},
					&ruleRefExpr{
						pos:  position{line: 393, col: 65, offset: 17309},
						name: "CharactersWithQuotePunctuation",
					},
				},
			},
		},
		{
			name: "QuotedTextCharacters",
			pos:  position{line: 395, col: 1, offset: 17444},
			expr: &oneOrMoreExpr{
				pos: position{line: 395, col: 25, offset: 17468},
				expr: &seqExpr{
					pos: position{line: 395, col: 26, offset: 17469},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 395, col: 26, offset: 17469},
							expr: &ruleRefExpr{
								pos:  position{line: 395, col: 27, offset: 17470},
								name: "NEWLINE",
							},
						},
						&notExpr{
							pos: position{line: 395, col: 35, offset: 17478},
							expr: &ruleRefExpr{
								pos:  position{line: 395, col: 36, offset: 17479},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 395, col: 39, offset: 17482},
							expr: &litMatcher{
								pos:        position{line: 395, col: 40, offset: 17483},
								val:        "*",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 395, col: 44, offset: 17487},
							expr: &litMatcher{
								pos:        position{line: 395, col: 45, offset: 17488},
								val:        "_",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 395, col: 49, offset: 17492},
							expr: &litMatcher{
								pos:        position{line: 395, col: 50, offset: 17493},
								val:        "`",
								ignoreCase: false,
							},
						},
						&anyMatcher{
							line: 395, col: 54, offset: 17497,
						},
					},
				},
			},
		},
		{
			name: "CharactersWithQuotePunctuation",
			pos:  position{line: 397, col: 1, offset: 17540},
			expr: &actionExpr{
				pos: position{line: 397, col: 35, offset: 17574},
				run: (*parser).callonCharactersWithQuotePunctuation1,
				expr: &oneOrMoreExpr{
					pos: position{line: 397, col: 35, offset: 17574},
					expr: &seqExpr{
						pos: position{line: 397, col: 36, offset: 17575},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 397, col: 36, offset: 17575},
								expr: &ruleRefExpr{
									pos:  position{line: 397, col: 37, offset: 17576},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 397, col: 45, offset: 17584},
								expr: &ruleRefExpr{
									pos:  position{line: 397, col: 46, offset: 17585},
									name: "WS",
								},
							},
							&anyMatcher{
								line: 397, col: 50, offset: 17589,
							},
						},
					},
				},
			},
		},
		{
			name: "UnbalancedQuotePunctuation",
			pos:  position{line: 402, col: 1, offset: 17834},
			expr: &choiceExpr{
				pos: position{line: 402, col: 31, offset: 17864},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 402, col: 31, offset: 17864},
						val:        "*",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 402, col: 37, offset: 17870},
						val:        "_",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 402, col: 43, offset: 17876},
						val:        "`",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "Passthrough",
			pos:  position{line: 407, col: 1, offset: 17988},
			expr: &choiceExpr{
				pos: position{line: 407, col: 16, offset: 18003},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 407, col: 16, offset: 18003},
						name: "TriplePlusPassthrough",
					},
					&ruleRefExpr{
						pos:  position{line: 407, col: 40, offset: 18027},
						name: "SinglePlusPassthrough",
					},
					&ruleRefExpr{
						pos:  position{line: 407, col: 64, offset: 18051},
						name: "PassthroughMacro",
					},
				},
			},
		},
		{
			name: "SinglePlusPassthrough",
			pos:  position{line: 409, col: 1, offset: 18069},
			expr: &actionExpr{
				pos: position{line: 409, col: 26, offset: 18094},
				run: (*parser).callonSinglePlusPassthrough1,
				expr: &seqExpr{
					pos: position{line: 409, col: 26, offset: 18094},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 409, col: 26, offset: 18094},
							val:        "+",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 409, col: 30, offset: 18098},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 409, col: 38, offset: 18106},
								expr: &seqExpr{
									pos: position{line: 409, col: 39, offset: 18107},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 409, col: 39, offset: 18107},
											expr: &ruleRefExpr{
												pos:  position{line: 409, col: 40, offset: 18108},
												name: "NEWLINE",
											},
										},
										&notExpr{
											pos: position{line: 409, col: 48, offset: 18116},
											expr: &litMatcher{
												pos:        position{line: 409, col: 49, offset: 18117},
												val:        "+",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 409, col: 53, offset: 18121,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 409, col: 57, offset: 18125},
							val:        "+",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TriplePlusPassthrough",
			pos:  position{line: 413, col: 1, offset: 18220},
			expr: &actionExpr{
				pos: position{line: 413, col: 26, offset: 18245},
				run: (*parser).callonTriplePlusPassthrough1,
				expr: &seqExpr{
					pos: position{line: 413, col: 26, offset: 18245},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 413, col: 26, offset: 18245},
							val:        "+++",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 413, col: 32, offset: 18251},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 413, col: 40, offset: 18259},
								expr: &seqExpr{
									pos: position{line: 413, col: 41, offset: 18260},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 413, col: 41, offset: 18260},
											expr: &litMatcher{
												pos:        position{line: 413, col: 42, offset: 18261},
												val:        "+++",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 413, col: 48, offset: 18267,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 413, col: 52, offset: 18271},
							val:        "+++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "PassthroughMacro",
			pos:  position{line: 417, col: 1, offset: 18368},
			expr: &choiceExpr{
				pos: position{line: 417, col: 21, offset: 18388},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 417, col: 21, offset: 18388},
						name: "SimplePassthroughMacro",
					},
					&ruleRefExpr{
						pos:  position{line: 417, col: 46, offset: 18413},
						name: "PassthroughWithQuotedText",
					},
				},
			},
		},
		{
			name: "SimplePassthroughMacro",
			pos:  position{line: 419, col: 1, offset: 18440},
			expr: &actionExpr{
				pos: position{line: 419, col: 27, offset: 18466},
				run: (*parser).callonSimplePassthroughMacro1,
				expr: &seqExpr{
					pos: position{line: 419, col: 27, offset: 18466},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 419, col: 27, offset: 18466},
							val:        "pass:[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 419, col: 36, offset: 18475},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 419, col: 44, offset: 18483},
								expr: &ruleRefExpr{
									pos:  position{line: 419, col: 45, offset: 18484},
									name: "PassthroughMacroCharacter",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 419, col: 73, offset: 18512},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "PassthroughWithQuotedText",
			pos:  position{line: 423, col: 1, offset: 18602},
			expr: &actionExpr{
				pos: position{line: 423, col: 30, offset: 18631},
				run: (*parser).callonPassthroughWithQuotedText1,
				expr: &seqExpr{
					pos: position{line: 423, col: 30, offset: 18631},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 423, col: 30, offset: 18631},
							val:        "pass:q[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 423, col: 40, offset: 18641},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 423, col: 48, offset: 18649},
								expr: &choiceExpr{
									pos: position{line: 423, col: 49, offset: 18650},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 423, col: 49, offset: 18650},
											name: "QuotedText",
										},
										&ruleRefExpr{
											pos:  position{line: 423, col: 62, offset: 18663},
											name: "PassthroughMacroCharacter",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 423, col: 90, offset: 18691},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "PassthroughMacroCharacter",
			pos:  position{line: 427, col: 1, offset: 18781},
			expr: &seqExpr{
				pos: position{line: 427, col: 31, offset: 18811},
				exprs: []interface{}{
					&notExpr{
						pos: position{line: 427, col: 31, offset: 18811},
						expr: &litMatcher{
							pos:        position{line: 427, col: 32, offset: 18812},
							val:        "]",
							ignoreCase: false,
						},
					},
					&anyMatcher{
						line: 427, col: 36, offset: 18816,
					},
				},
			},
		},
		{
			name: "CrossReference",
			pos:  position{line: 432, col: 1, offset: 18932},
			expr: &actionExpr{
				pos: position{line: 432, col: 19, offset: 18950},
				run: (*parser).callonCrossReference1,
				expr: &seqExpr{
					pos: position{line: 432, col: 19, offset: 18950},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 432, col: 19, offset: 18950},
							val:        "<<",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 432, col: 24, offset: 18955},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 432, col: 28, offset: 18959},
								name: "ID",
							},
						},
						&litMatcher{
							pos:        position{line: 432, col: 32, offset: 18963},
							val:        ">>",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Link",
			pos:  position{line: 439, col: 1, offset: 19122},
			expr: &choiceExpr{
				pos: position{line: 439, col: 9, offset: 19130},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 439, col: 9, offset: 19130},
						name: "RelativeLink",
					},
					&ruleRefExpr{
						pos:  position{line: 439, col: 24, offset: 19145},
						name: "ExternalLink",
					},
				},
			},
		},
		{
			name: "ExternalLink",
			pos:  position{line: 441, col: 1, offset: 19160},
			expr: &actionExpr{
				pos: position{line: 441, col: 17, offset: 19176},
				run: (*parser).callonExternalLink1,
				expr: &seqExpr{
					pos: position{line: 441, col: 17, offset: 19176},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 441, col: 17, offset: 19176},
							label: "url",
							expr: &seqExpr{
								pos: position{line: 441, col: 22, offset: 19181},
								exprs: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 441, col: 22, offset: 19181},
										name: "URL_SCHEME",
									},
									&ruleRefExpr{
										pos:  position{line: 441, col: 33, offset: 19192},
										name: "URL",
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 441, col: 38, offset: 19197},
							label: "text",
							expr: &zeroOrOneExpr{
								pos: position{line: 441, col: 43, offset: 19202},
								expr: &seqExpr{
									pos: position{line: 441, col: 44, offset: 19203},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 441, col: 44, offset: 19203},
											val:        "[",
											ignoreCase: false,
										},
										&zeroOrMoreExpr{
											pos: position{line: 441, col: 48, offset: 19207},
											expr: &ruleRefExpr{
												pos:  position{line: 441, col: 49, offset: 19208},
												name: "URL_TEXT",
											},
										},
										&litMatcher{
											pos:        position{line: 441, col: 60, offset: 19219},
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
			pos:  position{line: 448, col: 1, offset: 19380},
			expr: &actionExpr{
				pos: position{line: 448, col: 17, offset: 19396},
				run: (*parser).callonRelativeLink1,
				expr: &seqExpr{
					pos: position{line: 448, col: 17, offset: 19396},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 448, col: 17, offset: 19396},
							val:        "link:",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 448, col: 25, offset: 19404},
							label: "url",
							expr: &seqExpr{
								pos: position{line: 448, col: 30, offset: 19409},
								exprs: []interface{}{
									&zeroOrOneExpr{
										pos: position{line: 448, col: 30, offset: 19409},
										expr: &ruleRefExpr{
											pos:  position{line: 448, col: 30, offset: 19409},
											name: "URL_SCHEME",
										},
									},
									&ruleRefExpr{
										pos:  position{line: 448, col: 42, offset: 19421},
										name: "URL",
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 448, col: 47, offset: 19426},
							label: "text",
							expr: &seqExpr{
								pos: position{line: 448, col: 53, offset: 19432},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 448, col: 53, offset: 19432},
										val:        "[",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 448, col: 57, offset: 19436},
										expr: &ruleRefExpr{
											pos:  position{line: 448, col: 58, offset: 19437},
											name: "URL_TEXT",
										},
									},
									&litMatcher{
										pos:        position{line: 448, col: 69, offset: 19448},
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
			pos:  position{line: 458, col: 1, offset: 19710},
			expr: &actionExpr{
				pos: position{line: 458, col: 15, offset: 19724},
				run: (*parser).callonBlockImage1,
				expr: &seqExpr{
					pos: position{line: 458, col: 15, offset: 19724},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 458, col: 15, offset: 19724},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 458, col: 26, offset: 19735},
								expr: &ruleRefExpr{
									pos:  position{line: 458, col: 27, offset: 19736},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 458, col: 46, offset: 19755},
							label: "image",
							expr: &ruleRefExpr{
								pos:  position{line: 458, col: 52, offset: 19761},
								name: "BlockImageMacro",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 458, col: 69, offset: 19778},
							expr: &ruleRefExpr{
								pos:  position{line: 458, col: 69, offset: 19778},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 458, col: 73, offset: 19782},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "BlockImageMacro",
			pos:  position{line: 463, col: 1, offset: 19943},
			expr: &actionExpr{
				pos: position{line: 463, col: 20, offset: 19962},
				run: (*parser).callonBlockImageMacro1,
				expr: &seqExpr{
					pos: position{line: 463, col: 20, offset: 19962},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 463, col: 20, offset: 19962},
							val:        "image::",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 463, col: 30, offset: 19972},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 463, col: 36, offset: 19978},
								name: "URL",
							},
						},
						&litMatcher{
							pos:        position{line: 463, col: 41, offset: 19983},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 463, col: 45, offset: 19987},
							label: "attributes",
							expr: &zeroOrOneExpr{
								pos: position{line: 463, col: 57, offset: 19999},
								expr: &ruleRefExpr{
									pos:  position{line: 463, col: 57, offset: 19999},
									name: "URL_TEXT",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 463, col: 68, offset: 20010},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "InlineImage",
			pos:  position{line: 467, col: 1, offset: 20077},
			expr: &actionExpr{
				pos: position{line: 467, col: 16, offset: 20092},
				run: (*parser).callonInlineImage1,
				expr: &labeledExpr{
					pos:   position{line: 467, col: 16, offset: 20092},
					label: "image",
					expr: &ruleRefExpr{
						pos:  position{line: 467, col: 22, offset: 20098},
						name: "InlineImageMacro",
					},
				},
			},
		},
		{
			name: "InlineImageMacro",
			pos:  position{line: 472, col: 1, offset: 20245},
			expr: &actionExpr{
				pos: position{line: 472, col: 21, offset: 20265},
				run: (*parser).callonInlineImageMacro1,
				expr: &seqExpr{
					pos: position{line: 472, col: 21, offset: 20265},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 472, col: 21, offset: 20265},
							val:        "image:",
							ignoreCase: false,
						},
						&notExpr{
							pos: position{line: 472, col: 30, offset: 20274},
							expr: &litMatcher{
								pos:        position{line: 472, col: 31, offset: 20275},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 472, col: 35, offset: 20279},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 472, col: 41, offset: 20285},
								name: "URL",
							},
						},
						&litMatcher{
							pos:        position{line: 472, col: 46, offset: 20290},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 472, col: 50, offset: 20294},
							label: "attributes",
							expr: &zeroOrOneExpr{
								pos: position{line: 472, col: 62, offset: 20306},
								expr: &ruleRefExpr{
									pos:  position{line: 472, col: 62, offset: 20306},
									name: "URL_TEXT",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 472, col: 73, offset: 20317},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "DelimitedBlock",
			pos:  position{line: 479, col: 1, offset: 20647},
			expr: &choiceExpr{
				pos: position{line: 479, col: 19, offset: 20665},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 479, col: 19, offset: 20665},
						name: "FencedBlock",
					},
					&ruleRefExpr{
						pos:  position{line: 479, col: 33, offset: 20679},
						name: "ListingBlock",
					},
				},
			},
		},
		{
			name: "BlockDelimiter",
			pos:  position{line: 481, col: 1, offset: 20694},
			expr: &choiceExpr{
				pos: position{line: 481, col: 19, offset: 20712},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 481, col: 19, offset: 20712},
						name: "FencedBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 481, col: 42, offset: 20735},
						name: "ListingBlockDelimiter",
					},
				},
			},
		},
		{
			name: "FencedBlockDelimiter",
			pos:  position{line: 483, col: 1, offset: 20758},
			expr: &litMatcher{
				pos:        position{line: 483, col: 25, offset: 20782},
				val:        "```",
				ignoreCase: false,
			},
		},
		{
			name: "FencedBlock",
			pos:  position{line: 485, col: 1, offset: 20789},
			expr: &actionExpr{
				pos: position{line: 485, col: 16, offset: 20804},
				run: (*parser).callonFencedBlock1,
				expr: &seqExpr{
					pos: position{line: 485, col: 16, offset: 20804},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 485, col: 16, offset: 20804},
							name: "FencedBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 485, col: 37, offset: 20825},
							expr: &ruleRefExpr{
								pos:  position{line: 485, col: 37, offset: 20825},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 485, col: 41, offset: 20829},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 485, col: 49, offset: 20837},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 485, col: 58, offset: 20846},
								name: "FencedBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 485, col: 78, offset: 20866},
							name: "FencedBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 485, col: 99, offset: 20887},
							expr: &ruleRefExpr{
								pos:  position{line: 485, col: 99, offset: 20887},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 485, col: 103, offset: 20891},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "FencedBlockContent",
			pos:  position{line: 489, col: 1, offset: 20979},
			expr: &labeledExpr{
				pos:   position{line: 489, col: 23, offset: 21001},
				label: "content",
				expr: &zeroOrMoreExpr{
					pos: position{line: 489, col: 31, offset: 21009},
					expr: &seqExpr{
						pos: position{line: 489, col: 32, offset: 21010},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 489, col: 32, offset: 21010},
								expr: &ruleRefExpr{
									pos:  position{line: 489, col: 33, offset: 21011},
									name: "FencedBlockDelimiter",
								},
							},
							&anyMatcher{
								line: 489, col: 54, offset: 21032,
							},
						},
					},
				},
			},
		},
		{
			name: "ListingBlockDelimiter",
			pos:  position{line: 491, col: 1, offset: 21038},
			expr: &litMatcher{
				pos:        position{line: 491, col: 26, offset: 21063},
				val:        "----",
				ignoreCase: false,
			},
		},
		{
			name: "ListingBlock",
			pos:  position{line: 493, col: 1, offset: 21071},
			expr: &actionExpr{
				pos: position{line: 493, col: 17, offset: 21087},
				run: (*parser).callonListingBlock1,
				expr: &seqExpr{
					pos: position{line: 493, col: 17, offset: 21087},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 493, col: 17, offset: 21087},
							name: "ListingBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 493, col: 39, offset: 21109},
							expr: &ruleRefExpr{
								pos:  position{line: 493, col: 39, offset: 21109},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 493, col: 43, offset: 21113},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 493, col: 51, offset: 21121},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 493, col: 60, offset: 21130},
								name: "ListingBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 493, col: 81, offset: 21151},
							name: "ListingBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 493, col: 103, offset: 21173},
							expr: &ruleRefExpr{
								pos:  position{line: 493, col: 103, offset: 21173},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 493, col: 107, offset: 21177},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ListingBlockContent",
			pos:  position{line: 497, col: 1, offset: 21266},
			expr: &labeledExpr{
				pos:   position{line: 497, col: 24, offset: 21289},
				label: "content",
				expr: &zeroOrMoreExpr{
					pos: position{line: 497, col: 32, offset: 21297},
					expr: &seqExpr{
						pos: position{line: 497, col: 33, offset: 21298},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 497, col: 33, offset: 21298},
								expr: &ruleRefExpr{
									pos:  position{line: 497, col: 34, offset: 21299},
									name: "ListingBlockDelimiter",
								},
							},
							&anyMatcher{
								line: 497, col: 56, offset: 21321,
							},
						},
					},
				},
			},
		},
		{
			name: "LiteralBlock",
			pos:  position{line: 502, col: 1, offset: 21594},
			expr: &choiceExpr{
				pos: position{line: 502, col: 17, offset: 21610},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 502, col: 17, offset: 21610},
						name: "ParagraphWithSpaces",
					},
					&ruleRefExpr{
						pos:  position{line: 502, col: 39, offset: 21632},
						name: "ParagraphWithLiteralBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 502, col: 76, offset: 21669},
						name: "ParagraphWithLiteralAttribute",
					},
				},
			},
		},
		{
			name: "ParagraphWithSpaces",
			pos:  position{line: 505, col: 1, offset: 21764},
			expr: &actionExpr{
				pos: position{line: 505, col: 24, offset: 21787},
				run: (*parser).callonParagraphWithSpaces1,
				expr: &seqExpr{
					pos: position{line: 505, col: 24, offset: 21787},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 505, col: 24, offset: 21787},
							label: "spaces",
							expr: &oneOrMoreExpr{
								pos: position{line: 505, col: 32, offset: 21795},
								expr: &ruleRefExpr{
									pos:  position{line: 505, col: 32, offset: 21795},
									name: "WS",
								},
							},
						},
						&notExpr{
							pos: position{line: 505, col: 37, offset: 21800},
							expr: &ruleRefExpr{
								pos:  position{line: 505, col: 38, offset: 21801},
								name: "NEWLINE",
							},
						},
						&labeledExpr{
							pos:   position{line: 505, col: 46, offset: 21809},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 505, col: 55, offset: 21818},
								name: "LiteralBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 505, col: 76, offset: 21839},
							name: "EndOfLiteralBlock",
						},
					},
				},
			},
		},
		{
			name: "LiteralBlockContent",
			pos:  position{line: 510, col: 1, offset: 22020},
			expr: &actionExpr{
				pos: position{line: 510, col: 24, offset: 22043},
				run: (*parser).callonLiteralBlockContent1,
				expr: &labeledExpr{
					pos:   position{line: 510, col: 24, offset: 22043},
					label: "content",
					expr: &oneOrMoreExpr{
						pos: position{line: 510, col: 32, offset: 22051},
						expr: &seqExpr{
							pos: position{line: 510, col: 33, offset: 22052},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 510, col: 33, offset: 22052},
									expr: &seqExpr{
										pos: position{line: 510, col: 35, offset: 22054},
										exprs: []interface{}{
											&ruleRefExpr{
												pos:  position{line: 510, col: 35, offset: 22054},
												name: "NEWLINE",
											},
											&ruleRefExpr{
												pos:  position{line: 510, col: 43, offset: 22062},
												name: "BlankLine",
											},
										},
									},
								},
								&anyMatcher{
									line: 510, col: 54, offset: 22073,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "EndOfLiteralBlock",
			pos:  position{line: 515, col: 1, offset: 22158},
			expr: &choiceExpr{
				pos: position{line: 515, col: 22, offset: 22179},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 515, col: 22, offset: 22179},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 515, col: 22, offset: 22179},
								name: "NEWLINE",
							},
							&ruleRefExpr{
								pos:  position{line: 515, col: 30, offset: 22187},
								name: "BlankLine",
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 515, col: 42, offset: 22199},
						name: "NEWLINE",
					},
					&ruleRefExpr{
						pos:  position{line: 515, col: 52, offset: 22209},
						name: "EOF",
					},
				},
			},
		},
		{
			name: "ParagraphWithLiteralBlockDelimiter",
			pos:  position{line: 518, col: 1, offset: 22269},
			expr: &actionExpr{
				pos: position{line: 518, col: 39, offset: 22307},
				run: (*parser).callonParagraphWithLiteralBlockDelimiter1,
				expr: &seqExpr{
					pos: position{line: 518, col: 39, offset: 22307},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 518, col: 39, offset: 22307},
							name: "LiteralBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 518, col: 61, offset: 22329},
							expr: &ruleRefExpr{
								pos:  position{line: 518, col: 61, offset: 22329},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 518, col: 65, offset: 22333},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 518, col: 73, offset: 22341},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 518, col: 81, offset: 22349},
								expr: &seqExpr{
									pos: position{line: 518, col: 82, offset: 22350},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 518, col: 82, offset: 22350},
											expr: &ruleRefExpr{
												pos:  position{line: 518, col: 83, offset: 22351},
												name: "LiteralBlockDelimiter",
											},
										},
										&anyMatcher{
											line: 518, col: 105, offset: 22373,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 518, col: 109, offset: 22377},
							name: "LiteralBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 518, col: 131, offset: 22399},
							expr: &ruleRefExpr{
								pos:  position{line: 518, col: 131, offset: 22399},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 518, col: 135, offset: 22403},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "LiteralBlockDelimiter",
			pos:  position{line: 522, col: 1, offset: 22487},
			expr: &litMatcher{
				pos:        position{line: 522, col: 26, offset: 22512},
				val:        "....",
				ignoreCase: false,
			},
		},
		{
			name: "ParagraphWithLiteralAttribute",
			pos:  position{line: 525, col: 1, offset: 22574},
			expr: &actionExpr{
				pos: position{line: 525, col: 34, offset: 22607},
				run: (*parser).callonParagraphWithLiteralAttribute1,
				expr: &seqExpr{
					pos: position{line: 525, col: 34, offset: 22607},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 525, col: 34, offset: 22607},
							val:        "[literal]",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 525, col: 46, offset: 22619},
							expr: &ruleRefExpr{
								pos:  position{line: 525, col: 46, offset: 22619},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 525, col: 50, offset: 22623},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 525, col: 58, offset: 22631},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 525, col: 67, offset: 22640},
								name: "LiteralBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 525, col: 88, offset: 22661},
							name: "EndOfLiteralBlock",
						},
					},
				},
			},
		},
		{
			name: "ElementAttribute",
			pos:  position{line: 532, col: 1, offset: 22873},
			expr: &choiceExpr{
				pos: position{line: 532, col: 21, offset: 22893},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 532, col: 21, offset: 22893},
						name: "ElementLink",
					},
					&ruleRefExpr{
						pos:  position{line: 532, col: 35, offset: 22907},
						name: "ElementID",
					},
					&ruleRefExpr{
						pos:  position{line: 532, col: 47, offset: 22919},
						name: "ElementTitle",
					},
					&ruleRefExpr{
						pos:  position{line: 532, col: 62, offset: 22934},
						name: "InvalidElementAttribute",
					},
				},
			},
		},
		{
			name: "ElementLink",
			pos:  position{line: 535, col: 1, offset: 23014},
			expr: &actionExpr{
				pos: position{line: 535, col: 16, offset: 23029},
				run: (*parser).callonElementLink1,
				expr: &seqExpr{
					pos: position{line: 535, col: 16, offset: 23029},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 535, col: 16, offset: 23029},
							val:        "[link=",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 535, col: 25, offset: 23038},
							expr: &ruleRefExpr{
								pos:  position{line: 535, col: 25, offset: 23038},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 535, col: 29, offset: 23042},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 535, col: 34, offset: 23047},
								name: "URL",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 535, col: 38, offset: 23051},
							expr: &ruleRefExpr{
								pos:  position{line: 535, col: 38, offset: 23051},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 535, col: 42, offset: 23055},
							val:        "]",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 535, col: 46, offset: 23059},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ElementID",
			pos:  position{line: 539, col: 1, offset: 23115},
			expr: &choiceExpr{
				pos: position{line: 539, col: 14, offset: 23128},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 539, col: 14, offset: 23128},
						run: (*parser).callonElementID2,
						expr: &seqExpr{
							pos: position{line: 539, col: 14, offset: 23128},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 539, col: 14, offset: 23128},
									label: "id",
									expr: &ruleRefExpr{
										pos:  position{line: 539, col: 18, offset: 23132},
										name: "InlineElementID",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 539, col: 35, offset: 23149},
									name: "EOL",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 541, col: 5, offset: 23178},
						run: (*parser).callonElementID7,
						expr: &seqExpr{
							pos: position{line: 541, col: 5, offset: 23178},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 541, col: 5, offset: 23178},
									val:        "[#",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 541, col: 10, offset: 23183},
									label: "id",
									expr: &ruleRefExpr{
										pos:  position{line: 541, col: 14, offset: 23187},
										name: "ID",
									},
								},
								&litMatcher{
									pos:        position{line: 541, col: 18, offset: 23191},
									val:        "]",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 541, col: 22, offset: 23195},
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
			pos:  position{line: 545, col: 1, offset: 23247},
			expr: &actionExpr{
				pos: position{line: 545, col: 20, offset: 23266},
				run: (*parser).callonInlineElementID1,
				expr: &seqExpr{
					pos: position{line: 545, col: 20, offset: 23266},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 545, col: 20, offset: 23266},
							val:        "[[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 545, col: 25, offset: 23271},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 545, col: 29, offset: 23275},
								name: "ID",
							},
						},
						&litMatcher{
							pos:        position{line: 545, col: 33, offset: 23279},
							val:        "]]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ElementTitle",
			pos:  position{line: 551, col: 1, offset: 23474},
			expr: &actionExpr{
				pos: position{line: 551, col: 17, offset: 23490},
				run: (*parser).callonElementTitle1,
				expr: &seqExpr{
					pos: position{line: 551, col: 17, offset: 23490},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 551, col: 17, offset: 23490},
							val:        ".",
							ignoreCase: false,
						},
						&notExpr{
							pos: position{line: 551, col: 21, offset: 23494},
							expr: &litMatcher{
								pos:        position{line: 551, col: 22, offset: 23495},
								val:        ".",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 551, col: 26, offset: 23499},
							expr: &ruleRefExpr{
								pos:  position{line: 551, col: 27, offset: 23500},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 551, col: 30, offset: 23503},
							label: "title",
							expr: &oneOrMoreExpr{
								pos: position{line: 551, col: 36, offset: 23509},
								expr: &seqExpr{
									pos: position{line: 551, col: 37, offset: 23510},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 551, col: 37, offset: 23510},
											expr: &ruleRefExpr{
												pos:  position{line: 551, col: 38, offset: 23511},
												name: "NEWLINE",
											},
										},
										&anyMatcher{
											line: 551, col: 46, offset: 23519,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 551, col: 50, offset: 23523},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "InvalidElementAttribute",
			pos:  position{line: 555, col: 1, offset: 23588},
			expr: &actionExpr{
				pos: position{line: 555, col: 28, offset: 23615},
				run: (*parser).callonInvalidElementAttribute1,
				expr: &seqExpr{
					pos: position{line: 555, col: 28, offset: 23615},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 555, col: 28, offset: 23615},
							val:        "[",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 555, col: 32, offset: 23619},
							expr: &ruleRefExpr{
								pos:  position{line: 555, col: 32, offset: 23619},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 555, col: 36, offset: 23623},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 555, col: 44, offset: 23631},
								expr: &seqExpr{
									pos: position{line: 555, col: 45, offset: 23632},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 555, col: 45, offset: 23632},
											expr: &litMatcher{
												pos:        position{line: 555, col: 46, offset: 23633},
												val:        "]",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 555, col: 50, offset: 23637,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 555, col: 54, offset: 23641},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "BlankLine",
			pos:  position{line: 562, col: 1, offset: 23807},
			expr: &actionExpr{
				pos: position{line: 562, col: 14, offset: 23820},
				run: (*parser).callonBlankLine1,
				expr: &seqExpr{
					pos: position{line: 562, col: 14, offset: 23820},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 562, col: 14, offset: 23820},
							expr: &ruleRefExpr{
								pos:  position{line: 562, col: 15, offset: 23821},
								name: "EOF",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 562, col: 19, offset: 23825},
							expr: &ruleRefExpr{
								pos:  position{line: 562, col: 19, offset: 23825},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 562, col: 23, offset: 23829},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "Characters",
			pos:  position{line: 569, col: 1, offset: 23976},
			expr: &actionExpr{
				pos: position{line: 569, col: 15, offset: 23990},
				run: (*parser).callonCharacters1,
				expr: &oneOrMoreExpr{
					pos: position{line: 569, col: 15, offset: 23990},
					expr: &seqExpr{
						pos: position{line: 569, col: 16, offset: 23991},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 569, col: 16, offset: 23991},
								expr: &ruleRefExpr{
									pos:  position{line: 569, col: 17, offset: 23992},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 569, col: 25, offset: 24000},
								expr: &ruleRefExpr{
									pos:  position{line: 569, col: 26, offset: 24001},
									name: "WS",
								},
							},
							&anyMatcher{
								line: 569, col: 29, offset: 24004,
							},
						},
					},
				},
			},
		},
		{
			name: "URL",
			pos:  position{line: 574, col: 1, offset: 24045},
			expr: &actionExpr{
				pos: position{line: 574, col: 8, offset: 24052},
				run: (*parser).callonURL1,
				expr: &oneOrMoreExpr{
					pos: position{line: 574, col: 8, offset: 24052},
					expr: &seqExpr{
						pos: position{line: 574, col: 9, offset: 24053},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 574, col: 9, offset: 24053},
								expr: &ruleRefExpr{
									pos:  position{line: 574, col: 10, offset: 24054},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 574, col: 18, offset: 24062},
								expr: &ruleRefExpr{
									pos:  position{line: 574, col: 19, offset: 24063},
									name: "WS",
								},
							},
							&notExpr{
								pos: position{line: 574, col: 22, offset: 24066},
								expr: &litMatcher{
									pos:        position{line: 574, col: 23, offset: 24067},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 574, col: 27, offset: 24071},
								expr: &litMatcher{
									pos:        position{line: 574, col: 28, offset: 24072},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 574, col: 32, offset: 24076,
							},
						},
					},
				},
			},
		},
		{
			name: "ID",
			pos:  position{line: 578, col: 1, offset: 24116},
			expr: &actionExpr{
				pos: position{line: 578, col: 7, offset: 24122},
				run: (*parser).callonID1,
				expr: &oneOrMoreExpr{
					pos: position{line: 578, col: 7, offset: 24122},
					expr: &seqExpr{
						pos: position{line: 578, col: 8, offset: 24123},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 578, col: 8, offset: 24123},
								expr: &ruleRefExpr{
									pos:  position{line: 578, col: 9, offset: 24124},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 578, col: 17, offset: 24132},
								expr: &ruleRefExpr{
									pos:  position{line: 578, col: 18, offset: 24133},
									name: "WS",
								},
							},
							&notExpr{
								pos: position{line: 578, col: 21, offset: 24136},
								expr: &litMatcher{
									pos:        position{line: 578, col: 22, offset: 24137},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 578, col: 26, offset: 24141},
								expr: &litMatcher{
									pos:        position{line: 578, col: 27, offset: 24142},
									val:        "]",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 578, col: 31, offset: 24146},
								expr: &litMatcher{
									pos:        position{line: 578, col: 32, offset: 24147},
									val:        "<<",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 578, col: 37, offset: 24152},
								expr: &litMatcher{
									pos:        position{line: 578, col: 38, offset: 24153},
									val:        ">>",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 578, col: 42, offset: 24157,
							},
						},
					},
				},
			},
		},
		{
			name: "URL_TEXT",
			pos:  position{line: 582, col: 1, offset: 24197},
			expr: &actionExpr{
				pos: position{line: 582, col: 13, offset: 24209},
				run: (*parser).callonURL_TEXT1,
				expr: &oneOrMoreExpr{
					pos: position{line: 582, col: 13, offset: 24209},
					expr: &seqExpr{
						pos: position{line: 582, col: 14, offset: 24210},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 582, col: 14, offset: 24210},
								expr: &ruleRefExpr{
									pos:  position{line: 582, col: 15, offset: 24211},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 582, col: 23, offset: 24219},
								expr: &litMatcher{
									pos:        position{line: 582, col: 24, offset: 24220},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 582, col: 28, offset: 24224},
								expr: &litMatcher{
									pos:        position{line: 582, col: 29, offset: 24225},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 582, col: 33, offset: 24229,
							},
						},
					},
				},
			},
		},
		{
			name: "URL_SCHEME",
			pos:  position{line: 586, col: 1, offset: 24269},
			expr: &choiceExpr{
				pos: position{line: 586, col: 15, offset: 24283},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 586, col: 15, offset: 24283},
						val:        "http://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 586, col: 27, offset: 24295},
						val:        "https://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 586, col: 40, offset: 24308},
						val:        "ftp://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 586, col: 51, offset: 24319},
						val:        "irc://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 586, col: 62, offset: 24330},
						val:        "mailto:",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "DIGIT",
			pos:  position{line: 588, col: 1, offset: 24341},
			expr: &charClassMatcher{
				pos:        position{line: 588, col: 10, offset: 24350},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NEWLINE",
			pos:  position{line: 590, col: 1, offset: 24357},
			expr: &choiceExpr{
				pos: position{line: 590, col: 12, offset: 24368},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 590, col: 12, offset: 24368},
						val:        "\r\n",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 590, col: 21, offset: 24377},
						val:        "\r",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 590, col: 28, offset: 24384},
						val:        "\n",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "WS",
			pos:  position{line: 592, col: 1, offset: 24390},
			expr: &choiceExpr{
				pos: position{line: 592, col: 7, offset: 24396},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 592, col: 7, offset: 24396},
						val:        " ",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 592, col: 13, offset: 24402},
						run: (*parser).callonWS3,
						expr: &litMatcher{
							pos:        position{line: 592, col: 13, offset: 24402},
							val:        "\t",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 596, col: 1, offset: 24447},
			expr: &notExpr{
				pos: position{line: 596, col: 8, offset: 24454},
				expr: &anyMatcher{
					line: 596, col: 9, offset: 24455,
				},
			},
		},
		{
			name: "EOL",
			pos:  position{line: 598, col: 1, offset: 24458},
			expr: &choiceExpr{
				pos: position{line: 598, col: 8, offset: 24465},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 598, col: 8, offset: 24465},
						name: "NEWLINE",
					},
					&ruleRefExpr{
						pos:  position{line: 598, col: 18, offset: 24475},
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
