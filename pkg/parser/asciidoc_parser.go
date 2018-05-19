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

	"github.com/bytesparadise/libasciidoc/pkg/types"
)

// *****************************************************************************************
// This file is generated after its sibling `asciidoc-grammar.peg` file. DO NOT MODIFY !
// *****************************************************************************************

var g = &grammar{
	rules: []*rule{
		{
			name: "Document",
			pos:  position{line: 18, col: 1, offset: 504},
			expr: &actionExpr{
				pos: position{line: 18, col: 13, offset: 516},
				run: (*parser).callonDocument1,
				expr: &seqExpr{
					pos: position{line: 18, col: 13, offset: 516},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 18, col: 13, offset: 516},
							label: "frontMatter",
							expr: &zeroOrOneExpr{
								pos: position{line: 18, col: 26, offset: 529},
								expr: &ruleRefExpr{
									pos:  position{line: 18, col: 26, offset: 529},
									name: "FrontMatter",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 18, col: 40, offset: 543},
							label: "documentHeader",
							expr: &zeroOrOneExpr{
								pos: position{line: 18, col: 56, offset: 559},
								expr: &ruleRefExpr{
									pos:  position{line: 18, col: 56, offset: 559},
									name: "DocumentHeader",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 18, col: 73, offset: 576},
							label: "blocks",
							expr: &ruleRefExpr{
								pos:  position{line: 18, col: 81, offset: 584},
								name: "DocumentBlocks",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 18, col: 97, offset: 600},
							name: "EOF",
						},
					},
				},
			},
		},
		{
			name: "DocumentBlocks",
			pos:  position{line: 22, col: 1, offset: 688},
			expr: &choiceExpr{
				pos: position{line: 22, col: 19, offset: 706},
				alternatives: []interface{}{
					&labeledExpr{
						pos:   position{line: 22, col: 19, offset: 706},
						label: "content",
						expr: &seqExpr{
							pos: position{line: 22, col: 28, offset: 715},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 22, col: 28, offset: 715},
									expr: &ruleRefExpr{
										pos:  position{line: 22, col: 28, offset: 715},
										name: "Preamble",
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 22, col: 38, offset: 725},
									expr: &ruleRefExpr{
										pos:  position{line: 22, col: 38, offset: 725},
										name: "Section",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 22, col: 50, offset: 737},
						run: (*parser).callonDocumentBlocks8,
						expr: &labeledExpr{
							pos:   position{line: 22, col: 50, offset: 737},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 22, col: 59, offset: 746},
								expr: &ruleRefExpr{
									pos:  position{line: 22, col: 59, offset: 746},
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
			pos:  position{line: 26, col: 1, offset: 791},
			expr: &choiceExpr{
				pos: position{line: 26, col: 17, offset: 807},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 26, col: 17, offset: 807},
						name: "DocumentAttributeDeclaration",
					},
					&ruleRefExpr{
						pos:  position{line: 26, col: 48, offset: 838},
						name: "DocumentAttributeReset",
					},
					&ruleRefExpr{
						pos:  position{line: 26, col: 73, offset: 863},
						name: "TableOfContentsMacro",
					},
					&ruleRefExpr{
						pos:  position{line: 26, col: 96, offset: 886},
						name: "BlockImage",
					},
					&ruleRefExpr{
						pos:  position{line: 26, col: 109, offset: 899},
						name: "List",
					},
					&ruleRefExpr{
						pos:  position{line: 26, col: 116, offset: 906},
						name: "LiteralBlock",
					},
					&ruleRefExpr{
						pos:  position{line: 26, col: 131, offset: 921},
						name: "DelimitedBlock",
					},
					&ruleRefExpr{
						pos:  position{line: 26, col: 148, offset: 938},
						name: "BlankLine",
					},
					&ruleRefExpr{
						pos:  position{line: 26, col: 160, offset: 950},
						name: "Paragraph",
					},
				},
			},
		},
		{
			name: "Preamble",
			pos:  position{line: 28, col: 1, offset: 1023},
			expr: &actionExpr{
				pos: position{line: 28, col: 13, offset: 1035},
				run: (*parser).callonPreamble1,
				expr: &labeledExpr{
					pos:   position{line: 28, col: 13, offset: 1035},
					label: "elements",
					expr: &oneOrMoreExpr{
						pos: position{line: 28, col: 23, offset: 1045},
						expr: &ruleRefExpr{
							pos:  position{line: 28, col: 23, offset: 1045},
							name: "BlockElement",
						},
					},
				},
			},
		},
		{
			name: "FrontMatter",
			pos:  position{line: 35, col: 1, offset: 1228},
			expr: &ruleRefExpr{
				pos:  position{line: 35, col: 16, offset: 1243},
				name: "YamlFrontMatter",
			},
		},
		{
			name: "FrontMatter",
			pos:  position{line: 37, col: 1, offset: 1261},
			expr: &actionExpr{
				pos: position{line: 37, col: 16, offset: 1276},
				run: (*parser).callonFrontMatter1,
				expr: &seqExpr{
					pos: position{line: 37, col: 16, offset: 1276},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 37, col: 16, offset: 1276},
							name: "YamlFrontMatterToken",
						},
						&labeledExpr{
							pos:   position{line: 37, col: 37, offset: 1297},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 37, col: 46, offset: 1306},
								name: "YamlFrontMatterContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 37, col: 70, offset: 1330},
							name: "YamlFrontMatterToken",
						},
					},
				},
			},
		},
		{
			name: "YamlFrontMatterToken",
			pos:  position{line: 41, col: 1, offset: 1410},
			expr: &seqExpr{
				pos: position{line: 41, col: 26, offset: 1435},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 41, col: 26, offset: 1435},
						val:        "---",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 41, col: 32, offset: 1441},
						name: "EOL",
					},
				},
			},
		},
		{
			name: "YamlFrontMatterContent",
			pos:  position{line: 43, col: 1, offset: 1446},
			expr: &actionExpr{
				pos: position{line: 43, col: 27, offset: 1472},
				run: (*parser).callonYamlFrontMatterContent1,
				expr: &zeroOrMoreExpr{
					pos: position{line: 43, col: 27, offset: 1472},
					expr: &seqExpr{
						pos: position{line: 43, col: 28, offset: 1473},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 43, col: 28, offset: 1473},
								expr: &ruleRefExpr{
									pos:  position{line: 43, col: 29, offset: 1474},
									name: "YamlFrontMatterToken",
								},
							},
							&anyMatcher{
								line: 43, col: 50, offset: 1495,
							},
						},
					},
				},
			},
		},
		{
			name: "DocumentHeader",
			pos:  position{line: 51, col: 1, offset: 1719},
			expr: &actionExpr{
				pos: position{line: 51, col: 19, offset: 1737},
				run: (*parser).callonDocumentHeader1,
				expr: &seqExpr{
					pos: position{line: 51, col: 19, offset: 1737},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 51, col: 19, offset: 1737},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 51, col: 27, offset: 1745},
								name: "DocumentTitle",
							},
						},
						&labeledExpr{
							pos:   position{line: 51, col: 42, offset: 1760},
							label: "authors",
							expr: &zeroOrOneExpr{
								pos: position{line: 51, col: 51, offset: 1769},
								expr: &ruleRefExpr{
									pos:  position{line: 51, col: 51, offset: 1769},
									name: "DocumentAuthors",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 51, col: 69, offset: 1787},
							label: "revision",
							expr: &zeroOrOneExpr{
								pos: position{line: 51, col: 79, offset: 1797},
								expr: &ruleRefExpr{
									pos:  position{line: 51, col: 79, offset: 1797},
									name: "DocumentRevision",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 51, col: 98, offset: 1816},
							label: "otherAttributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 51, col: 115, offset: 1833},
								expr: &ruleRefExpr{
									pos:  position{line: 51, col: 115, offset: 1833},
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
			pos:  position{line: 55, col: 1, offset: 1964},
			expr: &actionExpr{
				pos: position{line: 55, col: 18, offset: 1981},
				run: (*parser).callonDocumentTitle1,
				expr: &seqExpr{
					pos: position{line: 55, col: 18, offset: 1981},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 55, col: 18, offset: 1981},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 55, col: 29, offset: 1992},
								expr: &ruleRefExpr{
									pos:  position{line: 55, col: 30, offset: 1993},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 55, col: 49, offset: 2012},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 55, col: 56, offset: 2019},
								val:        "=",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 55, col: 61, offset: 2024},
							expr: &ruleRefExpr{
								pos:  position{line: 55, col: 61, offset: 2024},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 55, col: 65, offset: 2028},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 55, col: 74, offset: 2037},
								name: "InlineElements",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 55, col: 90, offset: 2053},
							expr: &ruleRefExpr{
								pos:  position{line: 55, col: 90, offset: 2053},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 55, col: 94, offset: 2057},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 55, col: 97, offset: 2060},
								expr: &ruleRefExpr{
									pos:  position{line: 55, col: 98, offset: 2061},
									name: "InlineElementID",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 55, col: 116, offset: 2079},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthors",
			pos:  position{line: 59, col: 1, offset: 2195},
			expr: &choiceExpr{
				pos: position{line: 59, col: 20, offset: 2214},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 59, col: 20, offset: 2214},
						name: "DocumentAuthorsInlineForm",
					},
					&ruleRefExpr{
						pos:  position{line: 59, col: 48, offset: 2242},
						name: "DocumentAuthorsAttributeForm",
					},
				},
			},
		},
		{
			name: "DocumentAuthorsInlineForm",
			pos:  position{line: 61, col: 1, offset: 2272},
			expr: &actionExpr{
				pos: position{line: 61, col: 30, offset: 2301},
				run: (*parser).callonDocumentAuthorsInlineForm1,
				expr: &seqExpr{
					pos: position{line: 61, col: 30, offset: 2301},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 61, col: 30, offset: 2301},
							expr: &ruleRefExpr{
								pos:  position{line: 61, col: 30, offset: 2301},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 61, col: 34, offset: 2305},
							expr: &litMatcher{
								pos:        position{line: 61, col: 35, offset: 2306},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 61, col: 39, offset: 2310},
							label: "authors",
							expr: &oneOrMoreExpr{
								pos: position{line: 61, col: 48, offset: 2319},
								expr: &ruleRefExpr{
									pos:  position{line: 61, col: 48, offset: 2319},
									name: "DocumentAuthor",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 61, col: 65, offset: 2336},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthorsAttributeForm",
			pos:  position{line: 65, col: 1, offset: 2406},
			expr: &actionExpr{
				pos: position{line: 65, col: 33, offset: 2438},
				run: (*parser).callonDocumentAuthorsAttributeForm1,
				expr: &seqExpr{
					pos: position{line: 65, col: 33, offset: 2438},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 65, col: 33, offset: 2438},
							expr: &ruleRefExpr{
								pos:  position{line: 65, col: 33, offset: 2438},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 65, col: 37, offset: 2442},
							val:        ":author:",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 65, col: 48, offset: 2453},
							label: "author",
							expr: &ruleRefExpr{
								pos:  position{line: 65, col: 56, offset: 2461},
								name: "DocumentAuthor",
							},
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthor",
			pos:  position{line: 69, col: 1, offset: 2552},
			expr: &actionExpr{
				pos: position{line: 69, col: 19, offset: 2570},
				run: (*parser).callonDocumentAuthor1,
				expr: &seqExpr{
					pos: position{line: 69, col: 19, offset: 2570},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 69, col: 19, offset: 2570},
							expr: &ruleRefExpr{
								pos:  position{line: 69, col: 19, offset: 2570},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 69, col: 23, offset: 2574},
							label: "namePart1",
							expr: &ruleRefExpr{
								pos:  position{line: 69, col: 34, offset: 2585},
								name: "DocumentAuthorNamePart",
							},
						},
						&labeledExpr{
							pos:   position{line: 69, col: 58, offset: 2609},
							label: "namePart2",
							expr: &zeroOrOneExpr{
								pos: position{line: 69, col: 68, offset: 2619},
								expr: &ruleRefExpr{
									pos:  position{line: 69, col: 69, offset: 2620},
									name: "DocumentAuthorNamePart",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 69, col: 94, offset: 2645},
							label: "namePart3",
							expr: &zeroOrOneExpr{
								pos: position{line: 69, col: 104, offset: 2655},
								expr: &ruleRefExpr{
									pos:  position{line: 69, col: 105, offset: 2656},
									name: "DocumentAuthorNamePart",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 69, col: 130, offset: 2681},
							label: "email",
							expr: &zeroOrOneExpr{
								pos: position{line: 69, col: 136, offset: 2687},
								expr: &ruleRefExpr{
									pos:  position{line: 69, col: 137, offset: 2688},
									name: "DocumentAuthorEmail",
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 69, col: 159, offset: 2710},
							expr: &ruleRefExpr{
								pos:  position{line: 69, col: 159, offset: 2710},
								name: "WS",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 69, col: 163, offset: 2714},
							expr: &litMatcher{
								pos:        position{line: 69, col: 163, offset: 2714},
								val:        ";",
								ignoreCase: false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 69, col: 168, offset: 2719},
							expr: &ruleRefExpr{
								pos:  position{line: 69, col: 168, offset: 2719},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthorNamePart",
			pos:  position{line: 74, col: 1, offset: 2884},
			expr: &seqExpr{
				pos: position{line: 74, col: 27, offset: 2910},
				exprs: []interface{}{
					&notExpr{
						pos: position{line: 74, col: 27, offset: 2910},
						expr: &litMatcher{
							pos:        position{line: 74, col: 28, offset: 2911},
							val:        "<",
							ignoreCase: false,
						},
					},
					&notExpr{
						pos: position{line: 74, col: 32, offset: 2915},
						expr: &litMatcher{
							pos:        position{line: 74, col: 33, offset: 2916},
							val:        ";",
							ignoreCase: false,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 74, col: 37, offset: 2920},
						name: "Word",
					},
					&zeroOrMoreExpr{
						pos: position{line: 74, col: 42, offset: 2925},
						expr: &ruleRefExpr{
							pos:  position{line: 74, col: 42, offset: 2925},
							name: "WS",
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthorEmail",
			pos:  position{line: 76, col: 1, offset: 2930},
			expr: &seqExpr{
				pos: position{line: 76, col: 24, offset: 2953},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 76, col: 24, offset: 2953},
						val:        "<",
						ignoreCase: false,
					},
					&labeledExpr{
						pos:   position{line: 76, col: 28, offset: 2957},
						label: "email",
						expr: &oneOrMoreExpr{
							pos: position{line: 76, col: 34, offset: 2963},
							expr: &seqExpr{
								pos: position{line: 76, col: 35, offset: 2964},
								exprs: []interface{}{
									&notExpr{
										pos: position{line: 76, col: 35, offset: 2964},
										expr: &litMatcher{
											pos:        position{line: 76, col: 36, offset: 2965},
											val:        ">",
											ignoreCase: false,
										},
									},
									&notExpr{
										pos: position{line: 76, col: 40, offset: 2969},
										expr: &ruleRefExpr{
											pos:  position{line: 76, col: 41, offset: 2970},
											name: "EOL",
										},
									},
									&anyMatcher{
										line: 76, col: 45, offset: 2974,
									},
								},
							},
						},
					},
					&litMatcher{
						pos:        position{line: 76, col: 49, offset: 2978},
						val:        ">",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "DocumentRevision",
			pos:  position{line: 80, col: 1, offset: 3114},
			expr: &actionExpr{
				pos: position{line: 80, col: 21, offset: 3134},
				run: (*parser).callonDocumentRevision1,
				expr: &seqExpr{
					pos: position{line: 80, col: 21, offset: 3134},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 80, col: 21, offset: 3134},
							expr: &ruleRefExpr{
								pos:  position{line: 80, col: 21, offset: 3134},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 80, col: 25, offset: 3138},
							expr: &litMatcher{
								pos:        position{line: 80, col: 26, offset: 3139},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 80, col: 30, offset: 3143},
							label: "revnumber",
							expr: &zeroOrOneExpr{
								pos: position{line: 80, col: 40, offset: 3153},
								expr: &ruleRefExpr{
									pos:  position{line: 80, col: 41, offset: 3154},
									name: "DocumentRevisionNumber",
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 80, col: 66, offset: 3179},
							expr: &litMatcher{
								pos:        position{line: 80, col: 66, offset: 3179},
								val:        ",",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 80, col: 71, offset: 3184},
							label: "revdate",
							expr: &zeroOrOneExpr{
								pos: position{line: 80, col: 79, offset: 3192},
								expr: &ruleRefExpr{
									pos:  position{line: 80, col: 80, offset: 3193},
									name: "DocumentRevisionDate",
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 80, col: 103, offset: 3216},
							expr: &litMatcher{
								pos:        position{line: 80, col: 103, offset: 3216},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 80, col: 108, offset: 3221},
							label: "revremark",
							expr: &zeroOrOneExpr{
								pos: position{line: 80, col: 118, offset: 3231},
								expr: &ruleRefExpr{
									pos:  position{line: 80, col: 119, offset: 3232},
									name: "DocumentRevisionRemark",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 80, col: 144, offset: 3257},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentRevisionNumber",
			pos:  position{line: 85, col: 1, offset: 3430},
			expr: &choiceExpr{
				pos: position{line: 85, col: 27, offset: 3456},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 85, col: 27, offset: 3456},
						exprs: []interface{}{
							&litMatcher{
								pos:        position{line: 85, col: 27, offset: 3456},
								val:        "v",
								ignoreCase: true,
							},
							&ruleRefExpr{
								pos:  position{line: 85, col: 32, offset: 3461},
								name: "DIGIT",
							},
							&zeroOrMoreExpr{
								pos: position{line: 85, col: 39, offset: 3468},
								expr: &seqExpr{
									pos: position{line: 85, col: 40, offset: 3469},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 85, col: 40, offset: 3469},
											expr: &ruleRefExpr{
												pos:  position{line: 85, col: 41, offset: 3470},
												name: "EOL",
											},
										},
										&notExpr{
											pos: position{line: 85, col: 45, offset: 3474},
											expr: &litMatcher{
												pos:        position{line: 85, col: 46, offset: 3475},
												val:        ",",
												ignoreCase: false,
											},
										},
										&notExpr{
											pos: position{line: 85, col: 50, offset: 3479},
											expr: &litMatcher{
												pos:        position{line: 85, col: 51, offset: 3480},
												val:        ":",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 85, col: 55, offset: 3484,
										},
									},
								},
							},
						},
					},
					&seqExpr{
						pos: position{line: 85, col: 61, offset: 3490},
						exprs: []interface{}{
							&zeroOrOneExpr{
								pos: position{line: 85, col: 61, offset: 3490},
								expr: &litMatcher{
									pos:        position{line: 85, col: 61, offset: 3490},
									val:        "v",
									ignoreCase: true,
								},
							},
							&ruleRefExpr{
								pos:  position{line: 85, col: 67, offset: 3496},
								name: "DIGIT",
							},
							&zeroOrMoreExpr{
								pos: position{line: 85, col: 74, offset: 3503},
								expr: &seqExpr{
									pos: position{line: 85, col: 75, offset: 3504},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 85, col: 75, offset: 3504},
											expr: &ruleRefExpr{
												pos:  position{line: 85, col: 76, offset: 3505},
												name: "EOL",
											},
										},
										&notExpr{
											pos: position{line: 85, col: 80, offset: 3509},
											expr: &litMatcher{
												pos:        position{line: 85, col: 81, offset: 3510},
												val:        ",",
												ignoreCase: false,
											},
										},
										&notExpr{
											pos: position{line: 85, col: 85, offset: 3514},
											expr: &litMatcher{
												pos:        position{line: 85, col: 86, offset: 3515},
												val:        ":",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 85, col: 90, offset: 3519,
										},
									},
								},
							},
							&zeroOrMoreExpr{
								pos: position{line: 85, col: 94, offset: 3523},
								expr: &ruleRefExpr{
									pos:  position{line: 85, col: 94, offset: 3523},
									name: "WS",
								},
							},
							&andExpr{
								pos: position{line: 85, col: 98, offset: 3527},
								expr: &litMatcher{
									pos:        position{line: 85, col: 99, offset: 3528},
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
			pos:  position{line: 86, col: 1, offset: 3532},
			expr: &zeroOrMoreExpr{
				pos: position{line: 86, col: 25, offset: 3556},
				expr: &seqExpr{
					pos: position{line: 86, col: 26, offset: 3557},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 86, col: 26, offset: 3557},
							expr: &ruleRefExpr{
								pos:  position{line: 86, col: 27, offset: 3558},
								name: "EOL",
							},
						},
						&notExpr{
							pos: position{line: 86, col: 31, offset: 3562},
							expr: &litMatcher{
								pos:        position{line: 86, col: 32, offset: 3563},
								val:        ":",
								ignoreCase: false,
							},
						},
						&anyMatcher{
							line: 86, col: 36, offset: 3567,
						},
					},
				},
			},
		},
		{
			name: "DocumentRevisionRemark",
			pos:  position{line: 87, col: 1, offset: 3572},
			expr: &zeroOrMoreExpr{
				pos: position{line: 87, col: 27, offset: 3598},
				expr: &seqExpr{
					pos: position{line: 87, col: 28, offset: 3599},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 87, col: 28, offset: 3599},
							expr: &ruleRefExpr{
								pos:  position{line: 87, col: 29, offset: 3600},
								name: "EOL",
							},
						},
						&anyMatcher{
							line: 87, col: 33, offset: 3604,
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclaration",
			pos:  position{line: 92, col: 1, offset: 3724},
			expr: &choiceExpr{
				pos: position{line: 92, col: 33, offset: 3756},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 92, col: 33, offset: 3756},
						name: "DocumentAttributeDeclarationWithNameOnly",
					},
					&ruleRefExpr{
						pos:  position{line: 92, col: 76, offset: 3799},
						name: "DocumentAttributeDeclarationWithNameAndValue",
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclarationWithNameOnly",
			pos:  position{line: 94, col: 1, offset: 3846},
			expr: &actionExpr{
				pos: position{line: 94, col: 45, offset: 3890},
				run: (*parser).callonDocumentAttributeDeclarationWithNameOnly1,
				expr: &seqExpr{
					pos: position{line: 94, col: 45, offset: 3890},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 94, col: 45, offset: 3890},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 94, col: 49, offset: 3894},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 94, col: 55, offset: 3900},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 94, col: 70, offset: 3915},
							val:        ":",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 94, col: 74, offset: 3919},
							expr: &ruleRefExpr{
								pos:  position{line: 94, col: 74, offset: 3919},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 94, col: 78, offset: 3923},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclarationWithNameAndValue",
			pos:  position{line: 98, col: 1, offset: 4008},
			expr: &actionExpr{
				pos: position{line: 98, col: 49, offset: 4056},
				run: (*parser).callonDocumentAttributeDeclarationWithNameAndValue1,
				expr: &seqExpr{
					pos: position{line: 98, col: 49, offset: 4056},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 98, col: 49, offset: 4056},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 98, col: 53, offset: 4060},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 98, col: 59, offset: 4066},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 98, col: 74, offset: 4081},
							val:        ":",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 98, col: 78, offset: 4085},
							expr: &ruleRefExpr{
								pos:  position{line: 98, col: 78, offset: 4085},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 98, col: 82, offset: 4089},
							label: "value",
							expr: &zeroOrMoreExpr{
								pos: position{line: 98, col: 88, offset: 4095},
								expr: &seqExpr{
									pos: position{line: 98, col: 89, offset: 4096},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 98, col: 89, offset: 4096},
											expr: &ruleRefExpr{
												pos:  position{line: 98, col: 90, offset: 4097},
												name: "NEWLINE",
											},
										},
										&anyMatcher{
											line: 98, col: 98, offset: 4105,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 98, col: 102, offset: 4109},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeReset",
			pos:  position{line: 102, col: 1, offset: 4212},
			expr: &choiceExpr{
				pos: position{line: 102, col: 27, offset: 4238},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 102, col: 27, offset: 4238},
						name: "DocumentAttributeResetWithSectionTitleBangSymbol",
					},
					&ruleRefExpr{
						pos:  position{line: 102, col: 78, offset: 4289},
						name: "DocumentAttributeResetWithTrailingBangSymbol",
					},
				},
			},
		},
		{
			name: "DocumentAttributeResetWithSectionTitleBangSymbol",
			pos:  position{line: 104, col: 1, offset: 4335},
			expr: &actionExpr{
				pos: position{line: 104, col: 53, offset: 4387},
				run: (*parser).callonDocumentAttributeResetWithSectionTitleBangSymbol1,
				expr: &seqExpr{
					pos: position{line: 104, col: 53, offset: 4387},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 104, col: 53, offset: 4387},
							val:        ":!",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 104, col: 58, offset: 4392},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 104, col: 64, offset: 4398},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 104, col: 79, offset: 4413},
							val:        ":",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 104, col: 83, offset: 4417},
							expr: &ruleRefExpr{
								pos:  position{line: 104, col: 83, offset: 4417},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 104, col: 87, offset: 4421},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeResetWithTrailingBangSymbol",
			pos:  position{line: 108, col: 1, offset: 4495},
			expr: &actionExpr{
				pos: position{line: 108, col: 49, offset: 4543},
				run: (*parser).callonDocumentAttributeResetWithTrailingBangSymbol1,
				expr: &seqExpr{
					pos: position{line: 108, col: 49, offset: 4543},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 108, col: 49, offset: 4543},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 108, col: 53, offset: 4547},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 108, col: 59, offset: 4553},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 108, col: 74, offset: 4568},
							val:        "!:",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 108, col: 79, offset: 4573},
							expr: &ruleRefExpr{
								pos:  position{line: 108, col: 79, offset: 4573},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 108, col: 83, offset: 4577},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeSubstitution",
			pos:  position{line: 112, col: 1, offset: 4651},
			expr: &actionExpr{
				pos: position{line: 112, col: 34, offset: 4684},
				run: (*parser).callonDocumentAttributeSubstitution1,
				expr: &seqExpr{
					pos: position{line: 112, col: 34, offset: 4684},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 112, col: 34, offset: 4684},
							val:        "{",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 112, col: 38, offset: 4688},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 112, col: 44, offset: 4694},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 112, col: 59, offset: 4709},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "AttributeName",
			pos:  position{line: 119, col: 1, offset: 4957},
			expr: &seqExpr{
				pos: position{line: 119, col: 18, offset: 4974},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 119, col: 19, offset: 4975},
						alternatives: []interface{}{
							&charClassMatcher{
								pos:        position{line: 119, col: 19, offset: 4975},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
							&charClassMatcher{
								pos:        position{line: 119, col: 27, offset: 4983},
								val:        "[a-z]",
								ranges:     []rune{'a', 'z'},
								ignoreCase: false,
								inverted:   false,
							},
							&charClassMatcher{
								pos:        position{line: 119, col: 35, offset: 4991},
								val:        "[0-9]",
								ranges:     []rune{'0', '9'},
								ignoreCase: false,
								inverted:   false,
							},
							&litMatcher{
								pos:        position{line: 119, col: 43, offset: 4999},
								val:        "_",
								ignoreCase: false,
							},
						},
					},
					&zeroOrMoreExpr{
						pos: position{line: 119, col: 48, offset: 5004},
						expr: &choiceExpr{
							pos: position{line: 119, col: 49, offset: 5005},
							alternatives: []interface{}{
								&charClassMatcher{
									pos:        position{line: 119, col: 49, offset: 5005},
									val:        "[A-Z]",
									ranges:     []rune{'A', 'Z'},
									ignoreCase: false,
									inverted:   false,
								},
								&charClassMatcher{
									pos:        position{line: 119, col: 57, offset: 5013},
									val:        "[a-z]",
									ranges:     []rune{'a', 'z'},
									ignoreCase: false,
									inverted:   false,
								},
								&charClassMatcher{
									pos:        position{line: 119, col: 65, offset: 5021},
									val:        "[0-9]",
									ranges:     []rune{'0', '9'},
									ignoreCase: false,
									inverted:   false,
								},
								&litMatcher{
									pos:        position{line: 119, col: 73, offset: 5029},
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
			pos:  position{line: 124, col: 1, offset: 5149},
			expr: &seqExpr{
				pos: position{line: 124, col: 25, offset: 5173},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 124, col: 25, offset: 5173},
						val:        "toc::[]",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 124, col: 35, offset: 5183},
						name: "NEWLINE",
					},
				},
			},
		},
		{
			name: "ElementAttribute",
			pos:  position{line: 129, col: 1, offset: 5306},
			expr: &actionExpr{
				pos: position{line: 129, col: 21, offset: 5326},
				run: (*parser).callonElementAttribute1,
				expr: &seqExpr{
					pos: position{line: 129, col: 21, offset: 5326},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 129, col: 21, offset: 5326},
							label: "attr",
							expr: &choiceExpr{
								pos: position{line: 129, col: 27, offset: 5332},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 129, col: 27, offset: 5332},
										name: "ElementID",
									},
									&ruleRefExpr{
										pos:  position{line: 129, col: 39, offset: 5344},
										name: "ElementTitle",
									},
									&ruleRefExpr{
										pos:  position{line: 129, col: 54, offset: 5359},
										name: "AdmonitionMarkerAttribute",
									},
									&ruleRefExpr{
										pos:  position{line: 129, col: 82, offset: 5387},
										name: "AttributeGroup",
									},
									&ruleRefExpr{
										pos:  position{line: 129, col: 99, offset: 5404},
										name: "InvalidElementAttribute",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 129, col: 124, offset: 5429},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ElementID",
			pos:  position{line: 133, col: 1, offset: 5520},
			expr: &choiceExpr{
				pos: position{line: 133, col: 14, offset: 5533},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 133, col: 14, offset: 5533},
						run: (*parser).callonElementID2,
						expr: &labeledExpr{
							pos:   position{line: 133, col: 14, offset: 5533},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 133, col: 18, offset: 5537},
								name: "InlineElementID",
							},
						},
					},
					&actionExpr{
						pos: position{line: 135, col: 5, offset: 5579},
						run: (*parser).callonElementID5,
						expr: &seqExpr{
							pos: position{line: 135, col: 5, offset: 5579},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 135, col: 5, offset: 5579},
									val:        "[#",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 135, col: 10, offset: 5584},
									label: "id",
									expr: &ruleRefExpr{
										pos:  position{line: 135, col: 14, offset: 5588},
										name: "ID",
									},
								},
								&litMatcher{
									pos:        position{line: 135, col: 18, offset: 5592},
									val:        "]",
									ignoreCase: false,
								},
								&zeroOrMoreExpr{
									pos: position{line: 135, col: 22, offset: 5596},
									expr: &ruleRefExpr{
										pos:  position{line: 135, col: 22, offset: 5596},
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
			pos:  position{line: 139, col: 1, offset: 5648},
			expr: &actionExpr{
				pos: position{line: 139, col: 20, offset: 5667},
				run: (*parser).callonInlineElementID1,
				expr: &seqExpr{
					pos: position{line: 139, col: 20, offset: 5667},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 139, col: 20, offset: 5667},
							val:        "[[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 139, col: 25, offset: 5672},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 139, col: 29, offset: 5676},
								name: "ID",
							},
						},
						&litMatcher{
							pos:        position{line: 139, col: 33, offset: 5680},
							val:        "]]",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 139, col: 38, offset: 5685},
							expr: &ruleRefExpr{
								pos:  position{line: 139, col: 38, offset: 5685},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "ElementTitle",
			pos:  position{line: 145, col: 1, offset: 5879},
			expr: &actionExpr{
				pos: position{line: 145, col: 17, offset: 5895},
				run: (*parser).callonElementTitle1,
				expr: &seqExpr{
					pos: position{line: 145, col: 17, offset: 5895},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 145, col: 17, offset: 5895},
							val:        ".",
							ignoreCase: false,
						},
						&notExpr{
							pos: position{line: 145, col: 21, offset: 5899},
							expr: &litMatcher{
								pos:        position{line: 145, col: 22, offset: 5900},
								val:        ".",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 145, col: 26, offset: 5904},
							expr: &ruleRefExpr{
								pos:  position{line: 145, col: 27, offset: 5905},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 145, col: 30, offset: 5908},
							label: "title",
							expr: &oneOrMoreExpr{
								pos: position{line: 145, col: 36, offset: 5914},
								expr: &seqExpr{
									pos: position{line: 145, col: 37, offset: 5915},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 145, col: 37, offset: 5915},
											expr: &ruleRefExpr{
												pos:  position{line: 145, col: 38, offset: 5916},
												name: "NEWLINE",
											},
										},
										&anyMatcher{
											line: 145, col: 46, offset: 5924,
										},
									},
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 145, col: 50, offset: 5928},
							expr: &ruleRefExpr{
								pos:  position{line: 145, col: 50, offset: 5928},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "AdmonitionMarkerAttribute",
			pos:  position{line: 150, col: 1, offset: 6073},
			expr: &actionExpr{
				pos: position{line: 150, col: 30, offset: 6102},
				run: (*parser).callonAdmonitionMarkerAttribute1,
				expr: &seqExpr{
					pos: position{line: 150, col: 30, offset: 6102},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 150, col: 30, offset: 6102},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 150, col: 34, offset: 6106},
							label: "k",
							expr: &ruleRefExpr{
								pos:  position{line: 150, col: 37, offset: 6109},
								name: "AdmonitionKind",
							},
						},
						&litMatcher{
							pos:        position{line: 150, col: 53, offset: 6125},
							val:        "]",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 150, col: 57, offset: 6129},
							expr: &ruleRefExpr{
								pos:  position{line: 150, col: 57, offset: 6129},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "AttributeGroup",
			pos:  position{line: 155, col: 1, offset: 6219},
			expr: &actionExpr{
				pos: position{line: 155, col: 19, offset: 6237},
				run: (*parser).callonAttributeGroup1,
				expr: &seqExpr{
					pos: position{line: 155, col: 19, offset: 6237},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 155, col: 19, offset: 6237},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 155, col: 23, offset: 6241},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 155, col: 34, offset: 6252},
								expr: &ruleRefExpr{
									pos:  position{line: 155, col: 35, offset: 6253},
									name: "GenericAttribute",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 155, col: 54, offset: 6272},
							val:        "]",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 155, col: 58, offset: 6276},
							expr: &ruleRefExpr{
								pos:  position{line: 155, col: 58, offset: 6276},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "GenericAttribute",
			pos:  position{line: 159, col: 1, offset: 6348},
			expr: &choiceExpr{
				pos: position{line: 159, col: 21, offset: 6368},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 159, col: 21, offset: 6368},
						run: (*parser).callonGenericAttribute2,
						expr: &seqExpr{
							pos: position{line: 159, col: 21, offset: 6368},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 159, col: 21, offset: 6368},
									label: "key",
									expr: &ruleRefExpr{
										pos:  position{line: 159, col: 26, offset: 6373},
										name: "AttributeKey",
									},
								},
								&litMatcher{
									pos:        position{line: 159, col: 40, offset: 6387},
									val:        "=",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 159, col: 44, offset: 6391},
									label: "value",
									expr: &ruleRefExpr{
										pos:  position{line: 159, col: 51, offset: 6398},
										name: "AttributeValue",
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 159, col: 67, offset: 6414},
									expr: &seqExpr{
										pos: position{line: 159, col: 68, offset: 6415},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 159, col: 68, offset: 6415},
												val:        ",",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 159, col: 72, offset: 6419},
												expr: &ruleRefExpr{
													pos:  position{line: 159, col: 72, offset: 6419},
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
						pos: position{line: 161, col: 5, offset: 6528},
						run: (*parser).callonGenericAttribute14,
						expr: &seqExpr{
							pos: position{line: 161, col: 5, offset: 6528},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 161, col: 5, offset: 6528},
									label: "key",
									expr: &ruleRefExpr{
										pos:  position{line: 161, col: 10, offset: 6533},
										name: "AttributeKey",
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 161, col: 24, offset: 6547},
									expr: &seqExpr{
										pos: position{line: 161, col: 25, offset: 6548},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 161, col: 25, offset: 6548},
												val:        ",",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 161, col: 29, offset: 6552},
												expr: &ruleRefExpr{
													pos:  position{line: 161, col: 29, offset: 6552},
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
			pos:  position{line: 165, col: 1, offset: 6646},
			expr: &actionExpr{
				pos: position{line: 165, col: 17, offset: 6662},
				run: (*parser).callonAttributeKey1,
				expr: &seqExpr{
					pos: position{line: 165, col: 17, offset: 6662},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 165, col: 17, offset: 6662},
							label: "key",
							expr: &oneOrMoreExpr{
								pos: position{line: 165, col: 22, offset: 6667},
								expr: &seqExpr{
									pos: position{line: 165, col: 23, offset: 6668},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 165, col: 23, offset: 6668},
											expr: &ruleRefExpr{
												pos:  position{line: 165, col: 24, offset: 6669},
												name: "WS",
											},
										},
										&notExpr{
											pos: position{line: 165, col: 27, offset: 6672},
											expr: &litMatcher{
												pos:        position{line: 165, col: 28, offset: 6673},
												val:        "=",
												ignoreCase: false,
											},
										},
										&notExpr{
											pos: position{line: 165, col: 32, offset: 6677},
											expr: &litMatcher{
												pos:        position{line: 165, col: 33, offset: 6678},
												val:        ",",
												ignoreCase: false,
											},
										},
										&notExpr{
											pos: position{line: 165, col: 37, offset: 6682},
											expr: &litMatcher{
												pos:        position{line: 165, col: 38, offset: 6683},
												val:        "]",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 165, col: 42, offset: 6687,
										},
									},
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 165, col: 46, offset: 6691},
							expr: &ruleRefExpr{
								pos:  position{line: 165, col: 46, offset: 6691},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "AttributeValue",
			pos:  position{line: 170, col: 1, offset: 6773},
			expr: &actionExpr{
				pos: position{line: 170, col: 19, offset: 6791},
				run: (*parser).callonAttributeValue1,
				expr: &seqExpr{
					pos: position{line: 170, col: 19, offset: 6791},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 170, col: 19, offset: 6791},
							expr: &ruleRefExpr{
								pos:  position{line: 170, col: 19, offset: 6791},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 170, col: 23, offset: 6795},
							label: "value",
							expr: &zeroOrMoreExpr{
								pos: position{line: 170, col: 29, offset: 6801},
								expr: &seqExpr{
									pos: position{line: 170, col: 30, offset: 6802},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 170, col: 30, offset: 6802},
											expr: &ruleRefExpr{
												pos:  position{line: 170, col: 31, offset: 6803},
												name: "WS",
											},
										},
										&notExpr{
											pos: position{line: 170, col: 34, offset: 6806},
											expr: &litMatcher{
												pos:        position{line: 170, col: 35, offset: 6807},
												val:        "=",
												ignoreCase: false,
											},
										},
										&notExpr{
											pos: position{line: 170, col: 39, offset: 6811},
											expr: &litMatcher{
												pos:        position{line: 170, col: 40, offset: 6812},
												val:        "]",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 170, col: 44, offset: 6816,
										},
									},
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 170, col: 48, offset: 6820},
							expr: &ruleRefExpr{
								pos:  position{line: 170, col: 48, offset: 6820},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "InvalidElementAttribute",
			pos:  position{line: 175, col: 1, offset: 6907},
			expr: &actionExpr{
				pos: position{line: 175, col: 28, offset: 6934},
				run: (*parser).callonInvalidElementAttribute1,
				expr: &seqExpr{
					pos: position{line: 175, col: 28, offset: 6934},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 175, col: 28, offset: 6934},
							val:        "[",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 175, col: 32, offset: 6938},
							expr: &ruleRefExpr{
								pos:  position{line: 175, col: 32, offset: 6938},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 175, col: 36, offset: 6942},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 175, col: 44, offset: 6950},
								expr: &seqExpr{
									pos: position{line: 175, col: 45, offset: 6951},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 175, col: 45, offset: 6951},
											expr: &litMatcher{
												pos:        position{line: 175, col: 46, offset: 6952},
												val:        "]",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 175, col: 50, offset: 6956,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 175, col: 54, offset: 6960},
							val:        "]",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 175, col: 58, offset: 6964},
							expr: &ruleRefExpr{
								pos:  position{line: 175, col: 58, offset: 6964},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "Section",
			pos:  position{line: 182, col: 1, offset: 7129},
			expr: &choiceExpr{
				pos: position{line: 182, col: 12, offset: 7140},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 182, col: 12, offset: 7140},
						name: "Section1",
					},
					&ruleRefExpr{
						pos:  position{line: 182, col: 23, offset: 7151},
						name: "Section2",
					},
					&ruleRefExpr{
						pos:  position{line: 182, col: 34, offset: 7162},
						name: "Section3",
					},
					&ruleRefExpr{
						pos:  position{line: 182, col: 45, offset: 7173},
						name: "Section4",
					},
					&ruleRefExpr{
						pos:  position{line: 182, col: 56, offset: 7184},
						name: "Section5",
					},
				},
			},
		},
		{
			name: "Section1",
			pos:  position{line: 185, col: 1, offset: 7195},
			expr: &actionExpr{
				pos: position{line: 185, col: 13, offset: 7207},
				run: (*parser).callonSection11,
				expr: &seqExpr{
					pos: position{line: 185, col: 13, offset: 7207},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 185, col: 13, offset: 7207},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 185, col: 21, offset: 7215},
								name: "Section1Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 185, col: 36, offset: 7230},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 185, col: 46, offset: 7240},
								expr: &ruleRefExpr{
									pos:  position{line: 185, col: 46, offset: 7240},
									name: "Section1Block",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section1Title",
			pos:  position{line: 189, col: 1, offset: 7347},
			expr: &actionExpr{
				pos: position{line: 189, col: 18, offset: 7364},
				run: (*parser).callonSection1Title1,
				expr: &seqExpr{
					pos: position{line: 189, col: 18, offset: 7364},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 189, col: 18, offset: 7364},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 189, col: 29, offset: 7375},
								expr: &ruleRefExpr{
									pos:  position{line: 189, col: 30, offset: 7376},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 189, col: 49, offset: 7395},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 189, col: 56, offset: 7402},
								val:        "==",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 189, col: 62, offset: 7408},
							expr: &ruleRefExpr{
								pos:  position{line: 189, col: 62, offset: 7408},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 189, col: 66, offset: 7412},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 189, col: 75, offset: 7421},
								name: "InlineElements",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 189, col: 91, offset: 7437},
							expr: &ruleRefExpr{
								pos:  position{line: 189, col: 91, offset: 7437},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 189, col: 95, offset: 7441},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 189, col: 98, offset: 7444},
								expr: &ruleRefExpr{
									pos:  position{line: 189, col: 99, offset: 7445},
									name: "InlineElementID",
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 189, col: 117, offset: 7463},
							expr: &ruleRefExpr{
								pos:  position{line: 189, col: 117, offset: 7463},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 189, col: 121, offset: 7467},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 189, col: 126, offset: 7472},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 189, col: 126, offset: 7472},
									expr: &ruleRefExpr{
										pos:  position{line: 189, col: 126, offset: 7472},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 189, col: 139, offset: 7485},
									name: "EOF",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section1Block",
			pos:  position{line: 193, col: 1, offset: 7601},
			expr: &actionExpr{
				pos: position{line: 193, col: 18, offset: 7618},
				run: (*parser).callonSection1Block1,
				expr: &seqExpr{
					pos: position{line: 193, col: 18, offset: 7618},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 193, col: 18, offset: 7618},
							expr: &ruleRefExpr{
								pos:  position{line: 193, col: 19, offset: 7619},
								name: "Section1",
							},
						},
						&labeledExpr{
							pos:   position{line: 193, col: 28, offset: 7628},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 193, col: 37, offset: 7637},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 193, col: 37, offset: 7637},
										name: "Section2",
									},
									&ruleRefExpr{
										pos:  position{line: 193, col: 48, offset: 7648},
										name: "Section3",
									},
									&ruleRefExpr{
										pos:  position{line: 193, col: 59, offset: 7659},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 193, col: 70, offset: 7670},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 193, col: 81, offset: 7681},
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
			pos:  position{line: 197, col: 1, offset: 7724},
			expr: &actionExpr{
				pos: position{line: 197, col: 13, offset: 7736},
				run: (*parser).callonSection21,
				expr: &seqExpr{
					pos: position{line: 197, col: 13, offset: 7736},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 197, col: 13, offset: 7736},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 197, col: 21, offset: 7744},
								name: "Section2Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 197, col: 36, offset: 7759},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 197, col: 46, offset: 7769},
								expr: &ruleRefExpr{
									pos:  position{line: 197, col: 46, offset: 7769},
									name: "Section2Block",
								},
							},
						},
						&andExpr{
							pos: position{line: 197, col: 62, offset: 7785},
							expr: &zeroOrMoreExpr{
								pos: position{line: 197, col: 63, offset: 7786},
								expr: &ruleRefExpr{
									pos:  position{line: 197, col: 64, offset: 7787},
									name: "Section2",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section2Title",
			pos:  position{line: 201, col: 1, offset: 7889},
			expr: &actionExpr{
				pos: position{line: 201, col: 18, offset: 7906},
				run: (*parser).callonSection2Title1,
				expr: &seqExpr{
					pos: position{line: 201, col: 18, offset: 7906},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 201, col: 18, offset: 7906},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 201, col: 29, offset: 7917},
								expr: &ruleRefExpr{
									pos:  position{line: 201, col: 30, offset: 7918},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 201, col: 49, offset: 7937},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 201, col: 56, offset: 7944},
								val:        "===",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 201, col: 63, offset: 7951},
							expr: &ruleRefExpr{
								pos:  position{line: 201, col: 63, offset: 7951},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 201, col: 67, offset: 7955},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 201, col: 76, offset: 7964},
								name: "InlineElements",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 201, col: 92, offset: 7980},
							expr: &ruleRefExpr{
								pos:  position{line: 201, col: 92, offset: 7980},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 201, col: 96, offset: 7984},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 201, col: 99, offset: 7987},
								expr: &ruleRefExpr{
									pos:  position{line: 201, col: 100, offset: 7988},
									name: "InlineElementID",
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 201, col: 118, offset: 8006},
							expr: &ruleRefExpr{
								pos:  position{line: 201, col: 118, offset: 8006},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 201, col: 122, offset: 8010},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 201, col: 127, offset: 8015},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 201, col: 127, offset: 8015},
									expr: &ruleRefExpr{
										pos:  position{line: 201, col: 127, offset: 8015},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 201, col: 140, offset: 8028},
									name: "EOF",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section2Block",
			pos:  position{line: 205, col: 1, offset: 8143},
			expr: &actionExpr{
				pos: position{line: 205, col: 18, offset: 8160},
				run: (*parser).callonSection2Block1,
				expr: &seqExpr{
					pos: position{line: 205, col: 18, offset: 8160},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 205, col: 18, offset: 8160},
							expr: &ruleRefExpr{
								pos:  position{line: 205, col: 19, offset: 8161},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 205, col: 28, offset: 8170},
							expr: &ruleRefExpr{
								pos:  position{line: 205, col: 29, offset: 8171},
								name: "Section2",
							},
						},
						&labeledExpr{
							pos:   position{line: 205, col: 38, offset: 8180},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 205, col: 47, offset: 8189},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 205, col: 47, offset: 8189},
										name: "Section3",
									},
									&ruleRefExpr{
										pos:  position{line: 205, col: 58, offset: 8200},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 205, col: 69, offset: 8211},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 205, col: 80, offset: 8222},
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
			pos:  position{line: 209, col: 1, offset: 8265},
			expr: &actionExpr{
				pos: position{line: 209, col: 13, offset: 8277},
				run: (*parser).callonSection31,
				expr: &seqExpr{
					pos: position{line: 209, col: 13, offset: 8277},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 209, col: 13, offset: 8277},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 209, col: 21, offset: 8285},
								name: "Section3Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 209, col: 36, offset: 8300},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 209, col: 46, offset: 8310},
								expr: &ruleRefExpr{
									pos:  position{line: 209, col: 46, offset: 8310},
									name: "Section3Block",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section3Title",
			pos:  position{line: 213, col: 1, offset: 8417},
			expr: &actionExpr{
				pos: position{line: 213, col: 18, offset: 8434},
				run: (*parser).callonSection3Title1,
				expr: &seqExpr{
					pos: position{line: 213, col: 18, offset: 8434},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 213, col: 18, offset: 8434},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 213, col: 29, offset: 8445},
								expr: &ruleRefExpr{
									pos:  position{line: 213, col: 30, offset: 8446},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 213, col: 49, offset: 8465},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 213, col: 56, offset: 8472},
								val:        "====",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 213, col: 64, offset: 8480},
							expr: &ruleRefExpr{
								pos:  position{line: 213, col: 64, offset: 8480},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 213, col: 68, offset: 8484},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 213, col: 77, offset: 8493},
								name: "InlineElements",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 213, col: 93, offset: 8509},
							expr: &ruleRefExpr{
								pos:  position{line: 213, col: 93, offset: 8509},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 213, col: 97, offset: 8513},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 213, col: 100, offset: 8516},
								expr: &ruleRefExpr{
									pos:  position{line: 213, col: 101, offset: 8517},
									name: "InlineElementID",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 213, col: 119, offset: 8535},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 213, col: 124, offset: 8540},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 213, col: 124, offset: 8540},
									expr: &ruleRefExpr{
										pos:  position{line: 213, col: 124, offset: 8540},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 213, col: 137, offset: 8553},
									name: "EOF",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section3Block",
			pos:  position{line: 217, col: 1, offset: 8668},
			expr: &actionExpr{
				pos: position{line: 217, col: 18, offset: 8685},
				run: (*parser).callonSection3Block1,
				expr: &seqExpr{
					pos: position{line: 217, col: 18, offset: 8685},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 217, col: 18, offset: 8685},
							expr: &ruleRefExpr{
								pos:  position{line: 217, col: 19, offset: 8686},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 217, col: 28, offset: 8695},
							expr: &ruleRefExpr{
								pos:  position{line: 217, col: 29, offset: 8696},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 217, col: 38, offset: 8705},
							expr: &ruleRefExpr{
								pos:  position{line: 217, col: 39, offset: 8706},
								name: "Section3",
							},
						},
						&labeledExpr{
							pos:   position{line: 217, col: 48, offset: 8715},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 217, col: 57, offset: 8724},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 217, col: 57, offset: 8724},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 217, col: 68, offset: 8735},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 217, col: 79, offset: 8746},
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
			pos:  position{line: 221, col: 1, offset: 8789},
			expr: &actionExpr{
				pos: position{line: 221, col: 13, offset: 8801},
				run: (*parser).callonSection41,
				expr: &seqExpr{
					pos: position{line: 221, col: 13, offset: 8801},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 221, col: 13, offset: 8801},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 221, col: 21, offset: 8809},
								name: "Section4Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 221, col: 36, offset: 8824},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 221, col: 46, offset: 8834},
								expr: &ruleRefExpr{
									pos:  position{line: 221, col: 46, offset: 8834},
									name: "Section4Block",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section4Title",
			pos:  position{line: 225, col: 1, offset: 8941},
			expr: &actionExpr{
				pos: position{line: 225, col: 18, offset: 8958},
				run: (*parser).callonSection4Title1,
				expr: &seqExpr{
					pos: position{line: 225, col: 18, offset: 8958},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 225, col: 18, offset: 8958},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 225, col: 29, offset: 8969},
								expr: &ruleRefExpr{
									pos:  position{line: 225, col: 30, offset: 8970},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 225, col: 49, offset: 8989},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 225, col: 56, offset: 8996},
								val:        "=====",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 225, col: 65, offset: 9005},
							expr: &ruleRefExpr{
								pos:  position{line: 225, col: 65, offset: 9005},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 225, col: 69, offset: 9009},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 225, col: 78, offset: 9018},
								name: "InlineElements",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 225, col: 94, offset: 9034},
							expr: &ruleRefExpr{
								pos:  position{line: 225, col: 94, offset: 9034},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 225, col: 98, offset: 9038},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 225, col: 101, offset: 9041},
								expr: &ruleRefExpr{
									pos:  position{line: 225, col: 102, offset: 9042},
									name: "InlineElementID",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 225, col: 120, offset: 9060},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 225, col: 125, offset: 9065},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 225, col: 125, offset: 9065},
									expr: &ruleRefExpr{
										pos:  position{line: 225, col: 125, offset: 9065},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 225, col: 138, offset: 9078},
									name: "EOF",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section4Block",
			pos:  position{line: 229, col: 1, offset: 9193},
			expr: &actionExpr{
				pos: position{line: 229, col: 18, offset: 9210},
				run: (*parser).callonSection4Block1,
				expr: &seqExpr{
					pos: position{line: 229, col: 18, offset: 9210},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 229, col: 18, offset: 9210},
							expr: &ruleRefExpr{
								pos:  position{line: 229, col: 19, offset: 9211},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 229, col: 28, offset: 9220},
							expr: &ruleRefExpr{
								pos:  position{line: 229, col: 29, offset: 9221},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 229, col: 38, offset: 9230},
							expr: &ruleRefExpr{
								pos:  position{line: 229, col: 39, offset: 9231},
								name: "Section3",
							},
						},
						&notExpr{
							pos: position{line: 229, col: 48, offset: 9240},
							expr: &ruleRefExpr{
								pos:  position{line: 229, col: 49, offset: 9241},
								name: "Section4",
							},
						},
						&labeledExpr{
							pos:   position{line: 229, col: 58, offset: 9250},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 229, col: 67, offset: 9259},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 229, col: 67, offset: 9259},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 229, col: 78, offset: 9270},
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
			pos:  position{line: 233, col: 1, offset: 9313},
			expr: &actionExpr{
				pos: position{line: 233, col: 13, offset: 9325},
				run: (*parser).callonSection51,
				expr: &seqExpr{
					pos: position{line: 233, col: 13, offset: 9325},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 233, col: 13, offset: 9325},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 233, col: 21, offset: 9333},
								name: "Section5Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 233, col: 36, offset: 9348},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 233, col: 46, offset: 9358},
								expr: &ruleRefExpr{
									pos:  position{line: 233, col: 46, offset: 9358},
									name: "Section5Block",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section5Title",
			pos:  position{line: 237, col: 1, offset: 9465},
			expr: &actionExpr{
				pos: position{line: 237, col: 18, offset: 9482},
				run: (*parser).callonSection5Title1,
				expr: &seqExpr{
					pos: position{line: 237, col: 18, offset: 9482},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 237, col: 18, offset: 9482},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 237, col: 29, offset: 9493},
								expr: &ruleRefExpr{
									pos:  position{line: 237, col: 30, offset: 9494},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 237, col: 49, offset: 9513},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 237, col: 56, offset: 9520},
								val:        "======",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 237, col: 66, offset: 9530},
							expr: &ruleRefExpr{
								pos:  position{line: 237, col: 66, offset: 9530},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 237, col: 70, offset: 9534},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 237, col: 79, offset: 9543},
								name: "InlineElements",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 237, col: 95, offset: 9559},
							expr: &ruleRefExpr{
								pos:  position{line: 237, col: 95, offset: 9559},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 237, col: 99, offset: 9563},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 237, col: 102, offset: 9566},
								expr: &ruleRefExpr{
									pos:  position{line: 237, col: 103, offset: 9567},
									name: "InlineElementID",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 237, col: 121, offset: 9585},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 237, col: 126, offset: 9590},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 237, col: 126, offset: 9590},
									expr: &ruleRefExpr{
										pos:  position{line: 237, col: 126, offset: 9590},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 237, col: 139, offset: 9603},
									name: "EOF",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Section5Block",
			pos:  position{line: 241, col: 1, offset: 9718},
			expr: &actionExpr{
				pos: position{line: 241, col: 18, offset: 9735},
				run: (*parser).callonSection5Block1,
				expr: &seqExpr{
					pos: position{line: 241, col: 18, offset: 9735},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 241, col: 18, offset: 9735},
							expr: &ruleRefExpr{
								pos:  position{line: 241, col: 19, offset: 9736},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 241, col: 28, offset: 9745},
							expr: &ruleRefExpr{
								pos:  position{line: 241, col: 29, offset: 9746},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 241, col: 38, offset: 9755},
							expr: &ruleRefExpr{
								pos:  position{line: 241, col: 39, offset: 9756},
								name: "Section3",
							},
						},
						&notExpr{
							pos: position{line: 241, col: 48, offset: 9765},
							expr: &ruleRefExpr{
								pos:  position{line: 241, col: 49, offset: 9766},
								name: "Section4",
							},
						},
						&notExpr{
							pos: position{line: 241, col: 58, offset: 9775},
							expr: &ruleRefExpr{
								pos:  position{line: 241, col: 59, offset: 9776},
								name: "Section5",
							},
						},
						&labeledExpr{
							pos:   position{line: 241, col: 68, offset: 9785},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 241, col: 77, offset: 9794},
								name: "BlockElement",
							},
						},
					},
				},
			},
		},
		{
			name: "List",
			pos:  position{line: 248, col: 1, offset: 9938},
			expr: &actionExpr{
				pos: position{line: 248, col: 9, offset: 9946},
				run: (*parser).callonList1,
				expr: &seqExpr{
					pos: position{line: 248, col: 9, offset: 9946},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 248, col: 9, offset: 9946},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 248, col: 20, offset: 9957},
								expr: &ruleRefExpr{
									pos:  position{line: 248, col: 21, offset: 9958},
									name: "ListAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 250, col: 5, offset: 10047},
							label: "elements",
							expr: &ruleRefExpr{
								pos:  position{line: 250, col: 14, offset: 10056},
								name: "ListItems",
							},
						},
					},
				},
			},
		},
		{
			name: "ListItems",
			pos:  position{line: 254, col: 1, offset: 10150},
			expr: &oneOrMoreExpr{
				pos: position{line: 254, col: 14, offset: 10163},
				expr: &choiceExpr{
					pos: position{line: 254, col: 15, offset: 10164},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 254, col: 15, offset: 10164},
							name: "OrderedListItem",
						},
						&ruleRefExpr{
							pos:  position{line: 254, col: 33, offset: 10182},
							name: "UnorderedListItem",
						},
						&ruleRefExpr{
							pos:  position{line: 254, col: 53, offset: 10202},
							name: "LabeledListItem",
						},
					},
				},
			},
		},
		{
			name: "ListAttribute",
			pos:  position{line: 256, col: 1, offset: 10221},
			expr: &actionExpr{
				pos: position{line: 256, col: 18, offset: 10238},
				run: (*parser).callonListAttribute1,
				expr: &seqExpr{
					pos: position{line: 256, col: 18, offset: 10238},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 256, col: 18, offset: 10238},
							label: "attribute",
							expr: &choiceExpr{
								pos: position{line: 256, col: 29, offset: 10249},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 256, col: 29, offset: 10249},
										name: "HorizontalLayout",
									},
									&ruleRefExpr{
										pos:  position{line: 256, col: 48, offset: 10268},
										name: "ListID",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 256, col: 56, offset: 10276},
							name: "NEWLINE",
						},
					},
				},
			},
		},
		{
			name: "ListID",
			pos:  position{line: 260, col: 1, offset: 10315},
			expr: &actionExpr{
				pos: position{line: 260, col: 11, offset: 10325},
				run: (*parser).callonListID1,
				expr: &seqExpr{
					pos: position{line: 260, col: 11, offset: 10325},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 260, col: 11, offset: 10325},
							val:        "[#",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 260, col: 16, offset: 10330},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 260, col: 20, offset: 10334},
								name: "ID",
							},
						},
						&litMatcher{
							pos:        position{line: 260, col: 24, offset: 10338},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "HorizontalLayout",
			pos:  position{line: 264, col: 1, offset: 10404},
			expr: &actionExpr{
				pos: position{line: 264, col: 21, offset: 10424},
				run: (*parser).callonHorizontalLayout1,
				expr: &litMatcher{
					pos:        position{line: 264, col: 21, offset: 10424},
					val:        "[horizontal]",
					ignoreCase: false,
				},
			},
		},
		{
			name: "InnerParagraph",
			pos:  position{line: 268, col: 1, offset: 10507},
			expr: &actionExpr{
				pos: position{line: 268, col: 20, offset: 10526},
				run: (*parser).callonInnerParagraph1,
				expr: &labeledExpr{
					pos:   position{line: 268, col: 20, offset: 10526},
					label: "lines",
					expr: &oneOrMoreExpr{
						pos: position{line: 268, col: 26, offset: 10532},
						expr: &seqExpr{
							pos: position{line: 269, col: 5, offset: 10538},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 269, col: 5, offset: 10538},
									expr: &ruleRefExpr{
										pos:  position{line: 269, col: 7, offset: 10540},
										name: "OrderedListItemPrefix",
									},
								},
								&notExpr{
									pos: position{line: 270, col: 5, offset: 10568},
									expr: &ruleRefExpr{
										pos:  position{line: 270, col: 7, offset: 10570},
										name: "UnorderedListItemPrefix",
									},
								},
								&notExpr{
									pos: position{line: 271, col: 5, offset: 10600},
									expr: &seqExpr{
										pos: position{line: 271, col: 7, offset: 10602},
										exprs: []interface{}{
											&ruleRefExpr{
												pos:  position{line: 271, col: 7, offset: 10602},
												name: "LabeledListItemTerm",
											},
											&ruleRefExpr{
												pos:  position{line: 271, col: 27, offset: 10622},
												name: "LabeledListItemSeparator",
											},
										},
									},
								},
								&notExpr{
									pos: position{line: 272, col: 5, offset: 10653},
									expr: &ruleRefExpr{
										pos:  position{line: 272, col: 7, offset: 10655},
										name: "ListItemContinuation",
									},
								},
								&notExpr{
									pos: position{line: 273, col: 5, offset: 10682},
									expr: &ruleRefExpr{
										pos:  position{line: 273, col: 7, offset: 10684},
										name: "ElementAttribute",
									},
								},
								&notExpr{
									pos: position{line: 274, col: 5, offset: 10706},
									expr: &ruleRefExpr{
										pos:  position{line: 274, col: 7, offset: 10708},
										name: "BlockDelimiter",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 275, col: 5, offset: 10728},
									name: "InlineElements",
								},
								&ruleRefExpr{
									pos:  position{line: 275, col: 20, offset: 10743},
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
			pos:  position{line: 279, col: 1, offset: 10813},
			expr: &actionExpr{
				pos: position{line: 279, col: 25, offset: 10837},
				run: (*parser).callonListItemContinuation1,
				expr: &seqExpr{
					pos: position{line: 279, col: 25, offset: 10837},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 279, col: 25, offset: 10837},
							val:        "+",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 279, col: 29, offset: 10841},
							expr: &ruleRefExpr{
								pos:  position{line: 279, col: 29, offset: 10841},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 279, col: 33, offset: 10845},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ContinuedBlockElement",
			pos:  position{line: 283, col: 1, offset: 10897},
			expr: &actionExpr{
				pos: position{line: 283, col: 26, offset: 10922},
				run: (*parser).callonContinuedBlockElement1,
				expr: &seqExpr{
					pos: position{line: 283, col: 26, offset: 10922},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 283, col: 26, offset: 10922},
							name: "ListItemContinuation",
						},
						&labeledExpr{
							pos:   position{line: 283, col: 47, offset: 10943},
							label: "element",
							expr: &ruleRefExpr{
								pos:  position{line: 283, col: 55, offset: 10951},
								name: "BlockElement",
							},
						},
					},
				},
			},
		},
		{
			name: "OrderedListItem",
			pos:  position{line: 290, col: 1, offset: 11107},
			expr: &actionExpr{
				pos: position{line: 290, col: 20, offset: 11126},
				run: (*parser).callonOrderedListItem1,
				expr: &seqExpr{
					pos: position{line: 290, col: 20, offset: 11126},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 290, col: 20, offset: 11126},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 290, col: 31, offset: 11137},
								expr: &ruleRefExpr{
									pos:  position{line: 290, col: 32, offset: 11138},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 290, col: 51, offset: 11157},
							label: "prefix",
							expr: &ruleRefExpr{
								pos:  position{line: 290, col: 59, offset: 11165},
								name: "OrderedListItemPrefix",
							},
						},
						&labeledExpr{
							pos:   position{line: 290, col: 82, offset: 11188},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 290, col: 91, offset: 11197},
								name: "OrderedListItemContent",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 290, col: 115, offset: 11221},
							expr: &ruleRefExpr{
								pos:  position{line: 290, col: 115, offset: 11221},
								name: "BlankLine",
							},
						},
					},
				},
			},
		},
		{
			name: "OrderedListItemPrefix",
			pos:  position{line: 294, col: 1, offset: 11364},
			expr: &choiceExpr{
				pos: position{line: 296, col: 1, offset: 11428},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 296, col: 1, offset: 11428},
						run: (*parser).callonOrderedListItemPrefix2,
						expr: &seqExpr{
							pos: position{line: 296, col: 1, offset: 11428},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 296, col: 1, offset: 11428},
									expr: &ruleRefExpr{
										pos:  position{line: 296, col: 1, offset: 11428},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 296, col: 5, offset: 11432},
									label: "style",
									expr: &litMatcher{
										pos:        position{line: 296, col: 12, offset: 11439},
										val:        ".",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 296, col: 17, offset: 11444},
									expr: &ruleRefExpr{
										pos:  position{line: 296, col: 17, offset: 11444},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 298, col: 5, offset: 11537},
						run: (*parser).callonOrderedListItemPrefix10,
						expr: &seqExpr{
							pos: position{line: 298, col: 5, offset: 11537},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 298, col: 5, offset: 11537},
									expr: &ruleRefExpr{
										pos:  position{line: 298, col: 5, offset: 11537},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 298, col: 9, offset: 11541},
									label: "style",
									expr: &litMatcher{
										pos:        position{line: 298, col: 16, offset: 11548},
										val:        "..",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 298, col: 22, offset: 11554},
									expr: &ruleRefExpr{
										pos:  position{line: 298, col: 22, offset: 11554},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 300, col: 5, offset: 11652},
						run: (*parser).callonOrderedListItemPrefix18,
						expr: &seqExpr{
							pos: position{line: 300, col: 5, offset: 11652},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 300, col: 5, offset: 11652},
									expr: &ruleRefExpr{
										pos:  position{line: 300, col: 5, offset: 11652},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 300, col: 9, offset: 11656},
									label: "style",
									expr: &litMatcher{
										pos:        position{line: 300, col: 16, offset: 11663},
										val:        "...",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 300, col: 23, offset: 11670},
									expr: &ruleRefExpr{
										pos:  position{line: 300, col: 23, offset: 11670},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 302, col: 5, offset: 11769},
						run: (*parser).callonOrderedListItemPrefix26,
						expr: &seqExpr{
							pos: position{line: 302, col: 5, offset: 11769},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 302, col: 5, offset: 11769},
									expr: &ruleRefExpr{
										pos:  position{line: 302, col: 5, offset: 11769},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 302, col: 9, offset: 11773},
									label: "style",
									expr: &litMatcher{
										pos:        position{line: 302, col: 16, offset: 11780},
										val:        "....",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 302, col: 24, offset: 11788},
									expr: &ruleRefExpr{
										pos:  position{line: 302, col: 24, offset: 11788},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 304, col: 5, offset: 11888},
						run: (*parser).callonOrderedListItemPrefix34,
						expr: &seqExpr{
							pos: position{line: 304, col: 5, offset: 11888},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 304, col: 5, offset: 11888},
									expr: &ruleRefExpr{
										pos:  position{line: 304, col: 5, offset: 11888},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 304, col: 9, offset: 11892},
									label: "style",
									expr: &litMatcher{
										pos:        position{line: 304, col: 16, offset: 11899},
										val:        ".....",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 304, col: 25, offset: 11908},
									expr: &ruleRefExpr{
										pos:  position{line: 304, col: 25, offset: 11908},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 307, col: 5, offset: 12031},
						run: (*parser).callonOrderedListItemPrefix42,
						expr: &seqExpr{
							pos: position{line: 307, col: 5, offset: 12031},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 307, col: 5, offset: 12031},
									expr: &ruleRefExpr{
										pos:  position{line: 307, col: 5, offset: 12031},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 307, col: 9, offset: 12035},
									label: "style",
									expr: &seqExpr{
										pos: position{line: 307, col: 16, offset: 12042},
										exprs: []interface{}{
											&oneOrMoreExpr{
												pos: position{line: 307, col: 16, offset: 12042},
												expr: &seqExpr{
													pos: position{line: 307, col: 17, offset: 12043},
													exprs: []interface{}{
														&notExpr{
															pos: position{line: 307, col: 17, offset: 12043},
															expr: &litMatcher{
																pos:        position{line: 307, col: 18, offset: 12044},
																val:        ".",
																ignoreCase: false,
															},
														},
														&notExpr{
															pos: position{line: 307, col: 22, offset: 12048},
															expr: &ruleRefExpr{
																pos:  position{line: 307, col: 23, offset: 12049},
																name: "WS",
															},
														},
														&notExpr{
															pos: position{line: 307, col: 26, offset: 12052},
															expr: &ruleRefExpr{
																pos:  position{line: 307, col: 27, offset: 12053},
																name: "NEWLINE",
															},
														},
														&charClassMatcher{
															pos:        position{line: 307, col: 35, offset: 12061},
															val:        "[0-9]",
															ranges:     []rune{'0', '9'},
															ignoreCase: false,
															inverted:   false,
														},
													},
												},
											},
											&litMatcher{
												pos:        position{line: 307, col: 43, offset: 12069},
												val:        ".",
												ignoreCase: false,
											},
										},
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 307, col: 48, offset: 12074},
									expr: &ruleRefExpr{
										pos:  position{line: 307, col: 48, offset: 12074},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 309, col: 5, offset: 12169},
						run: (*parser).callonOrderedListItemPrefix60,
						expr: &seqExpr{
							pos: position{line: 309, col: 5, offset: 12169},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 309, col: 5, offset: 12169},
									expr: &ruleRefExpr{
										pos:  position{line: 309, col: 5, offset: 12169},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 309, col: 9, offset: 12173},
									label: "style",
									expr: &seqExpr{
										pos: position{line: 309, col: 16, offset: 12180},
										exprs: []interface{}{
											&oneOrMoreExpr{
												pos: position{line: 309, col: 16, offset: 12180},
												expr: &seqExpr{
													pos: position{line: 309, col: 17, offset: 12181},
													exprs: []interface{}{
														&notExpr{
															pos: position{line: 309, col: 17, offset: 12181},
															expr: &litMatcher{
																pos:        position{line: 309, col: 18, offset: 12182},
																val:        ".",
																ignoreCase: false,
															},
														},
														&notExpr{
															pos: position{line: 309, col: 22, offset: 12186},
															expr: &ruleRefExpr{
																pos:  position{line: 309, col: 23, offset: 12187},
																name: "WS",
															},
														},
														&notExpr{
															pos: position{line: 309, col: 26, offset: 12190},
															expr: &ruleRefExpr{
																pos:  position{line: 309, col: 27, offset: 12191},
																name: "NEWLINE",
															},
														},
														&charClassMatcher{
															pos:        position{line: 309, col: 35, offset: 12199},
															val:        "[a-z]",
															ranges:     []rune{'a', 'z'},
															ignoreCase: false,
															inverted:   false,
														},
													},
												},
											},
											&litMatcher{
												pos:        position{line: 309, col: 43, offset: 12207},
												val:        ".",
												ignoreCase: false,
											},
										},
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 309, col: 48, offset: 12212},
									expr: &ruleRefExpr{
										pos:  position{line: 309, col: 48, offset: 12212},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 311, col: 5, offset: 12310},
						run: (*parser).callonOrderedListItemPrefix78,
						expr: &seqExpr{
							pos: position{line: 311, col: 5, offset: 12310},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 311, col: 5, offset: 12310},
									expr: &ruleRefExpr{
										pos:  position{line: 311, col: 5, offset: 12310},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 311, col: 9, offset: 12314},
									label: "style",
									expr: &seqExpr{
										pos: position{line: 311, col: 16, offset: 12321},
										exprs: []interface{}{
											&oneOrMoreExpr{
												pos: position{line: 311, col: 16, offset: 12321},
												expr: &seqExpr{
													pos: position{line: 311, col: 17, offset: 12322},
													exprs: []interface{}{
														&notExpr{
															pos: position{line: 311, col: 17, offset: 12322},
															expr: &litMatcher{
																pos:        position{line: 311, col: 18, offset: 12323},
																val:        ".",
																ignoreCase: false,
															},
														},
														&notExpr{
															pos: position{line: 311, col: 22, offset: 12327},
															expr: &ruleRefExpr{
																pos:  position{line: 311, col: 23, offset: 12328},
																name: "WS",
															},
														},
														&notExpr{
															pos: position{line: 311, col: 26, offset: 12331},
															expr: &ruleRefExpr{
																pos:  position{line: 311, col: 27, offset: 12332},
																name: "NEWLINE",
															},
														},
														&charClassMatcher{
															pos:        position{line: 311, col: 35, offset: 12340},
															val:        "[A-Z]",
															ranges:     []rune{'A', 'Z'},
															ignoreCase: false,
															inverted:   false,
														},
													},
												},
											},
											&litMatcher{
												pos:        position{line: 311, col: 43, offset: 12348},
												val:        ".",
												ignoreCase: false,
											},
										},
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 311, col: 48, offset: 12353},
									expr: &ruleRefExpr{
										pos:  position{line: 311, col: 48, offset: 12353},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 313, col: 5, offset: 12451},
						run: (*parser).callonOrderedListItemPrefix96,
						expr: &seqExpr{
							pos: position{line: 313, col: 5, offset: 12451},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 313, col: 5, offset: 12451},
									expr: &ruleRefExpr{
										pos:  position{line: 313, col: 5, offset: 12451},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 313, col: 9, offset: 12455},
									label: "style",
									expr: &seqExpr{
										pos: position{line: 313, col: 16, offset: 12462},
										exprs: []interface{}{
											&oneOrMoreExpr{
												pos: position{line: 313, col: 16, offset: 12462},
												expr: &seqExpr{
													pos: position{line: 313, col: 17, offset: 12463},
													exprs: []interface{}{
														&notExpr{
															pos: position{line: 313, col: 17, offset: 12463},
															expr: &litMatcher{
																pos:        position{line: 313, col: 18, offset: 12464},
																val:        ")",
																ignoreCase: false,
															},
														},
														&notExpr{
															pos: position{line: 313, col: 22, offset: 12468},
															expr: &ruleRefExpr{
																pos:  position{line: 313, col: 23, offset: 12469},
																name: "WS",
															},
														},
														&notExpr{
															pos: position{line: 313, col: 26, offset: 12472},
															expr: &ruleRefExpr{
																pos:  position{line: 313, col: 27, offset: 12473},
																name: "NEWLINE",
															},
														},
														&charClassMatcher{
															pos:        position{line: 313, col: 35, offset: 12481},
															val:        "[a-z]",
															ranges:     []rune{'a', 'z'},
															ignoreCase: false,
															inverted:   false,
														},
													},
												},
											},
											&litMatcher{
												pos:        position{line: 313, col: 43, offset: 12489},
												val:        ")",
												ignoreCase: false,
											},
										},
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 313, col: 48, offset: 12494},
									expr: &ruleRefExpr{
										pos:  position{line: 313, col: 48, offset: 12494},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 315, col: 5, offset: 12592},
						run: (*parser).callonOrderedListItemPrefix114,
						expr: &seqExpr{
							pos: position{line: 315, col: 5, offset: 12592},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 315, col: 5, offset: 12592},
									expr: &ruleRefExpr{
										pos:  position{line: 315, col: 5, offset: 12592},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 315, col: 9, offset: 12596},
									label: "style",
									expr: &seqExpr{
										pos: position{line: 315, col: 16, offset: 12603},
										exprs: []interface{}{
											&oneOrMoreExpr{
												pos: position{line: 315, col: 16, offset: 12603},
												expr: &seqExpr{
													pos: position{line: 315, col: 17, offset: 12604},
													exprs: []interface{}{
														&notExpr{
															pos: position{line: 315, col: 17, offset: 12604},
															expr: &litMatcher{
																pos:        position{line: 315, col: 18, offset: 12605},
																val:        ")",
																ignoreCase: false,
															},
														},
														&notExpr{
															pos: position{line: 315, col: 22, offset: 12609},
															expr: &ruleRefExpr{
																pos:  position{line: 315, col: 23, offset: 12610},
																name: "WS",
															},
														},
														&notExpr{
															pos: position{line: 315, col: 26, offset: 12613},
															expr: &ruleRefExpr{
																pos:  position{line: 315, col: 27, offset: 12614},
																name: "NEWLINE",
															},
														},
														&charClassMatcher{
															pos:        position{line: 315, col: 35, offset: 12622},
															val:        "[A-Z]",
															ranges:     []rune{'A', 'Z'},
															ignoreCase: false,
															inverted:   false,
														},
													},
												},
											},
											&litMatcher{
												pos:        position{line: 315, col: 43, offset: 12630},
												val:        ")",
												ignoreCase: false,
											},
										},
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 315, col: 48, offset: 12635},
									expr: &ruleRefExpr{
										pos:  position{line: 315, col: 48, offset: 12635},
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
			pos:  position{line: 319, col: 1, offset: 12733},
			expr: &actionExpr{
				pos: position{line: 319, col: 27, offset: 12759},
				run: (*parser).callonOrderedListItemContent1,
				expr: &labeledExpr{
					pos:   position{line: 319, col: 27, offset: 12759},
					label: "elements",
					expr: &seqExpr{
						pos: position{line: 319, col: 37, offset: 12769},
						exprs: []interface{}{
							&oneOrMoreExpr{
								pos: position{line: 319, col: 37, offset: 12769},
								expr: &ruleRefExpr{
									pos:  position{line: 319, col: 37, offset: 12769},
									name: "InnerParagraph",
								},
							},
							&zeroOrMoreExpr{
								pos: position{line: 319, col: 53, offset: 12785},
								expr: &ruleRefExpr{
									pos:  position{line: 319, col: 53, offset: 12785},
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
			pos:  position{line: 326, col: 1, offset: 13111},
			expr: &actionExpr{
				pos: position{line: 326, col: 22, offset: 13132},
				run: (*parser).callonUnorderedListItem1,
				expr: &seqExpr{
					pos: position{line: 326, col: 22, offset: 13132},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 326, col: 22, offset: 13132},
							label: "prefix",
							expr: &ruleRefExpr{
								pos:  position{line: 326, col: 30, offset: 13140},
								name: "UnorderedListItemPrefix",
							},
						},
						&labeledExpr{
							pos:   position{line: 326, col: 55, offset: 13165},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 326, col: 64, offset: 13174},
								name: "UnorderedListItemContent",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 326, col: 90, offset: 13200},
							expr: &ruleRefExpr{
								pos:  position{line: 326, col: 90, offset: 13200},
								name: "BlankLine",
							},
						},
					},
				},
			},
		},
		{
			name: "UnorderedListItemPrefix",
			pos:  position{line: 330, col: 1, offset: 13319},
			expr: &choiceExpr{
				pos: position{line: 330, col: 28, offset: 13346},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 330, col: 28, offset: 13346},
						run: (*parser).callonUnorderedListItemPrefix2,
						expr: &seqExpr{
							pos: position{line: 330, col: 28, offset: 13346},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 330, col: 28, offset: 13346},
									expr: &ruleRefExpr{
										pos:  position{line: 330, col: 28, offset: 13346},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 330, col: 32, offset: 13350},
									label: "level",
									expr: &litMatcher{
										pos:        position{line: 330, col: 39, offset: 13357},
										val:        "*****",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 330, col: 48, offset: 13366},
									expr: &ruleRefExpr{
										pos:  position{line: 330, col: 48, offset: 13366},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 332, col: 5, offset: 13505},
						run: (*parser).callonUnorderedListItemPrefix10,
						expr: &seqExpr{
							pos: position{line: 332, col: 5, offset: 13505},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 332, col: 5, offset: 13505},
									expr: &ruleRefExpr{
										pos:  position{line: 332, col: 5, offset: 13505},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 332, col: 9, offset: 13509},
									label: "level",
									expr: &litMatcher{
										pos:        position{line: 332, col: 16, offset: 13516},
										val:        "****",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 332, col: 24, offset: 13524},
									expr: &ruleRefExpr{
										pos:  position{line: 332, col: 24, offset: 13524},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 334, col: 5, offset: 13663},
						run: (*parser).callonUnorderedListItemPrefix18,
						expr: &seqExpr{
							pos: position{line: 334, col: 5, offset: 13663},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 334, col: 5, offset: 13663},
									expr: &ruleRefExpr{
										pos:  position{line: 334, col: 5, offset: 13663},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 334, col: 9, offset: 13667},
									label: "level",
									expr: &litMatcher{
										pos:        position{line: 334, col: 16, offset: 13674},
										val:        "***",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 334, col: 23, offset: 13681},
									expr: &ruleRefExpr{
										pos:  position{line: 334, col: 23, offset: 13681},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 336, col: 5, offset: 13821},
						run: (*parser).callonUnorderedListItemPrefix26,
						expr: &seqExpr{
							pos: position{line: 336, col: 5, offset: 13821},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 336, col: 5, offset: 13821},
									expr: &ruleRefExpr{
										pos:  position{line: 336, col: 5, offset: 13821},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 336, col: 9, offset: 13825},
									label: "level",
									expr: &litMatcher{
										pos:        position{line: 336, col: 16, offset: 13832},
										val:        "**",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 336, col: 22, offset: 13838},
									expr: &ruleRefExpr{
										pos:  position{line: 336, col: 22, offset: 13838},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 338, col: 5, offset: 13976},
						run: (*parser).callonUnorderedListItemPrefix34,
						expr: &seqExpr{
							pos: position{line: 338, col: 5, offset: 13976},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 338, col: 5, offset: 13976},
									expr: &ruleRefExpr{
										pos:  position{line: 338, col: 5, offset: 13976},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 338, col: 9, offset: 13980},
									label: "level",
									expr: &litMatcher{
										pos:        position{line: 338, col: 16, offset: 13987},
										val:        "*",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 338, col: 21, offset: 13992},
									expr: &ruleRefExpr{
										pos:  position{line: 338, col: 21, offset: 13992},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 340, col: 5, offset: 14129},
						run: (*parser).callonUnorderedListItemPrefix42,
						expr: &seqExpr{
							pos: position{line: 340, col: 5, offset: 14129},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 340, col: 5, offset: 14129},
									expr: &ruleRefExpr{
										pos:  position{line: 340, col: 5, offset: 14129},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 340, col: 9, offset: 14133},
									label: "level",
									expr: &litMatcher{
										pos:        position{line: 340, col: 16, offset: 14140},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 340, col: 21, offset: 14145},
									expr: &ruleRefExpr{
										pos:  position{line: 340, col: 21, offset: 14145},
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
			pos:  position{line: 344, col: 1, offset: 14275},
			expr: &actionExpr{
				pos: position{line: 344, col: 29, offset: 14303},
				run: (*parser).callonUnorderedListItemContent1,
				expr: &labeledExpr{
					pos:   position{line: 344, col: 29, offset: 14303},
					label: "elements",
					expr: &seqExpr{
						pos: position{line: 344, col: 39, offset: 14313},
						exprs: []interface{}{
							&oneOrMoreExpr{
								pos: position{line: 344, col: 39, offset: 14313},
								expr: &ruleRefExpr{
									pos:  position{line: 344, col: 39, offset: 14313},
									name: "InnerParagraph",
								},
							},
							&zeroOrMoreExpr{
								pos: position{line: 344, col: 55, offset: 14329},
								expr: &ruleRefExpr{
									pos:  position{line: 344, col: 55, offset: 14329},
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
			pos:  position{line: 351, col: 1, offset: 14653},
			expr: &choiceExpr{
				pos: position{line: 351, col: 20, offset: 14672},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 351, col: 20, offset: 14672},
						run: (*parser).callonLabeledListItem2,
						expr: &seqExpr{
							pos: position{line: 351, col: 20, offset: 14672},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 351, col: 20, offset: 14672},
									label: "term",
									expr: &ruleRefExpr{
										pos:  position{line: 351, col: 26, offset: 14678},
										name: "LabeledListItemTerm",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 351, col: 47, offset: 14699},
									name: "LabeledListItemSeparator",
								},
								&labeledExpr{
									pos:   position{line: 351, col: 72, offset: 14724},
									label: "description",
									expr: &ruleRefExpr{
										pos:  position{line: 351, col: 85, offset: 14737},
										name: "LabeledListItemDescription",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 353, col: 6, offset: 14859},
						run: (*parser).callonLabeledListItem9,
						expr: &seqExpr{
							pos: position{line: 353, col: 6, offset: 14859},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 353, col: 6, offset: 14859},
									label: "term",
									expr: &ruleRefExpr{
										pos:  position{line: 353, col: 12, offset: 14865},
										name: "LabeledListItemTerm",
									},
								},
								&litMatcher{
									pos:        position{line: 353, col: 33, offset: 14886},
									val:        "::",
									ignoreCase: false,
								},
								&zeroOrMoreExpr{
									pos: position{line: 353, col: 38, offset: 14891},
									expr: &ruleRefExpr{
										pos:  position{line: 353, col: 38, offset: 14891},
										name: "WS",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 353, col: 42, offset: 14895},
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
			pos:  position{line: 357, col: 1, offset: 15032},
			expr: &actionExpr{
				pos: position{line: 357, col: 24, offset: 15055},
				run: (*parser).callonLabeledListItemTerm1,
				expr: &labeledExpr{
					pos:   position{line: 357, col: 24, offset: 15055},
					label: "term",
					expr: &zeroOrMoreExpr{
						pos: position{line: 357, col: 29, offset: 15060},
						expr: &seqExpr{
							pos: position{line: 357, col: 30, offset: 15061},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 357, col: 30, offset: 15061},
									expr: &ruleRefExpr{
										pos:  position{line: 357, col: 31, offset: 15062},
										name: "NEWLINE",
									},
								},
								&notExpr{
									pos: position{line: 357, col: 39, offset: 15070},
									expr: &litMatcher{
										pos:        position{line: 357, col: 40, offset: 15071},
										val:        "::",
										ignoreCase: false,
									},
								},
								&anyMatcher{
									line: 357, col: 45, offset: 15076,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "LabeledListItemSeparator",
			pos:  position{line: 362, col: 1, offset: 15167},
			expr: &seqExpr{
				pos: position{line: 362, col: 30, offset: 15196},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 362, col: 30, offset: 15196},
						val:        "::",
						ignoreCase: false,
					},
					&oneOrMoreExpr{
						pos: position{line: 362, col: 35, offset: 15201},
						expr: &choiceExpr{
							pos: position{line: 362, col: 36, offset: 15202},
							alternatives: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 362, col: 36, offset: 15202},
									name: "WS",
								},
								&ruleRefExpr{
									pos:  position{line: 362, col: 41, offset: 15207},
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
			pos:  position{line: 364, col: 1, offset: 15218},
			expr: &actionExpr{
				pos: position{line: 364, col: 31, offset: 15248},
				run: (*parser).callonLabeledListItemDescription1,
				expr: &labeledExpr{
					pos:   position{line: 364, col: 31, offset: 15248},
					label: "elements",
					expr: &zeroOrMoreExpr{
						pos: position{line: 364, col: 40, offset: 15257},
						expr: &choiceExpr{
							pos: position{line: 364, col: 41, offset: 15258},
							alternatives: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 364, col: 41, offset: 15258},
									name: "InnerParagraph",
								},
								&ruleRefExpr{
									pos:  position{line: 364, col: 58, offset: 15275},
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
			pos:  position{line: 372, col: 1, offset: 15583},
			expr: &choiceExpr{
				pos: position{line: 372, col: 19, offset: 15601},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 372, col: 19, offset: 15601},
						run: (*parser).callonAdmonitionKind2,
						expr: &litMatcher{
							pos:        position{line: 372, col: 19, offset: 15601},
							val:        "TIP",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 374, col: 5, offset: 15639},
						run: (*parser).callonAdmonitionKind4,
						expr: &litMatcher{
							pos:        position{line: 374, col: 5, offset: 15639},
							val:        "NOTE",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 376, col: 5, offset: 15679},
						run: (*parser).callonAdmonitionKind6,
						expr: &litMatcher{
							pos:        position{line: 376, col: 5, offset: 15679},
							val:        "IMPORTANT",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 378, col: 5, offset: 15729},
						run: (*parser).callonAdmonitionKind8,
						expr: &litMatcher{
							pos:        position{line: 378, col: 5, offset: 15729},
							val:        "WARNING",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 380, col: 5, offset: 15775},
						run: (*parser).callonAdmonitionKind10,
						expr: &litMatcher{
							pos:        position{line: 380, col: 5, offset: 15775},
							val:        "CAUTION",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Paragraph",
			pos:  position{line: 389, col: 1, offset: 16078},
			expr: &choiceExpr{
				pos: position{line: 391, col: 5, offset: 16125},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 391, col: 5, offset: 16125},
						run: (*parser).callonParagraph2,
						expr: &seqExpr{
							pos: position{line: 391, col: 5, offset: 16125},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 391, col: 5, offset: 16125},
									label: "attributes",
									expr: &zeroOrMoreExpr{
										pos: position{line: 391, col: 16, offset: 16136},
										expr: &ruleRefExpr{
											pos:  position{line: 391, col: 17, offset: 16137},
											name: "ElementAttribute",
										},
									},
								},
								&notExpr{
									pos: position{line: 391, col: 36, offset: 16156},
									expr: &seqExpr{
										pos: position{line: 391, col: 38, offset: 16158},
										exprs: []interface{}{
											&oneOrMoreExpr{
												pos: position{line: 391, col: 38, offset: 16158},
												expr: &litMatcher{
													pos:        position{line: 391, col: 38, offset: 16158},
													val:        "=",
													ignoreCase: false,
												},
											},
											&oneOrMoreExpr{
												pos: position{line: 391, col: 43, offset: 16163},
												expr: &ruleRefExpr{
													pos:  position{line: 391, col: 43, offset: 16163},
													name: "WS",
												},
											},
											&notExpr{
												pos: position{line: 391, col: 47, offset: 16167},
												expr: &ruleRefExpr{
													pos:  position{line: 391, col: 48, offset: 16168},
													name: "NEWLINE",
												},
											},
										},
									},
								},
								&labeledExpr{
									pos:   position{line: 391, col: 57, offset: 16177},
									label: "t",
									expr: &ruleRefExpr{
										pos:  position{line: 391, col: 60, offset: 16180},
										name: "AdmonitionKind",
									},
								},
								&litMatcher{
									pos:        position{line: 391, col: 76, offset: 16196},
									val:        ": ",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 391, col: 81, offset: 16201},
									label: "lines",
									expr: &oneOrMoreExpr{
										pos: position{line: 391, col: 87, offset: 16207},
										expr: &seqExpr{
											pos: position{line: 391, col: 88, offset: 16208},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 391, col: 88, offset: 16208},
													name: "InlineElements",
												},
												&ruleRefExpr{
													pos:  position{line: 391, col: 103, offset: 16223},
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
						pos: position{line: 395, col: 5, offset: 16389},
						run: (*parser).callonParagraph23,
						expr: &seqExpr{
							pos: position{line: 395, col: 5, offset: 16389},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 395, col: 5, offset: 16389},
									label: "attributes",
									expr: &zeroOrMoreExpr{
										pos: position{line: 395, col: 16, offset: 16400},
										expr: &ruleRefExpr{
											pos:  position{line: 395, col: 17, offset: 16401},
											name: "ElementAttribute",
										},
									},
								},
								&notExpr{
									pos: position{line: 395, col: 36, offset: 16420},
									expr: &seqExpr{
										pos: position{line: 395, col: 38, offset: 16422},
										exprs: []interface{}{
											&oneOrMoreExpr{
												pos: position{line: 395, col: 38, offset: 16422},
												expr: &litMatcher{
													pos:        position{line: 395, col: 38, offset: 16422},
													val:        "=",
													ignoreCase: false,
												},
											},
											&oneOrMoreExpr{
												pos: position{line: 395, col: 43, offset: 16427},
												expr: &ruleRefExpr{
													pos:  position{line: 395, col: 43, offset: 16427},
													name: "WS",
												},
											},
											&notExpr{
												pos: position{line: 395, col: 47, offset: 16431},
												expr: &ruleRefExpr{
													pos:  position{line: 395, col: 48, offset: 16432},
													name: "NEWLINE",
												},
											},
										},
									},
								},
								&labeledExpr{
									pos:   position{line: 395, col: 57, offset: 16441},
									label: "lines",
									expr: &oneOrMoreExpr{
										pos: position{line: 395, col: 63, offset: 16447},
										expr: &seqExpr{
											pos: position{line: 395, col: 64, offset: 16448},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 395, col: 64, offset: 16448},
													name: "InlineElements",
												},
												&ruleRefExpr{
													pos:  position{line: 395, col: 79, offset: 16463},
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
			name: "InlineElements",
			pos:  position{line: 399, col: 1, offset: 16565},
			expr: &actionExpr{
				pos: position{line: 399, col: 19, offset: 16583},
				run: (*parser).callonInlineElements1,
				expr: &seqExpr{
					pos: position{line: 399, col: 19, offset: 16583},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 399, col: 19, offset: 16583},
							expr: &ruleRefExpr{
								pos:  position{line: 399, col: 20, offset: 16584},
								name: "BlockDelimiter",
							},
						},
						&labeledExpr{
							pos:   position{line: 399, col: 35, offset: 16599},
							label: "elements",
							expr: &oneOrMoreExpr{
								pos: position{line: 399, col: 44, offset: 16608},
								expr: &seqExpr{
									pos: position{line: 399, col: 45, offset: 16609},
									exprs: []interface{}{
										&zeroOrMoreExpr{
											pos: position{line: 399, col: 45, offset: 16609},
											expr: &ruleRefExpr{
												pos:  position{line: 399, col: 45, offset: 16609},
												name: "WS",
											},
										},
										&notExpr{
											pos: position{line: 399, col: 49, offset: 16613},
											expr: &ruleRefExpr{
												pos:  position{line: 399, col: 50, offset: 16614},
												name: "InlineElementID",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 399, col: 66, offset: 16630},
											name: "InlineElement",
										},
										&zeroOrMoreExpr{
											pos: position{line: 399, col: 80, offset: 16644},
											expr: &ruleRefExpr{
												pos:  position{line: 399, col: 80, offset: 16644},
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
			name: "InlineElement",
			pos:  position{line: 403, col: 1, offset: 16756},
			expr: &choiceExpr{
				pos: position{line: 403, col: 18, offset: 16773},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 403, col: 18, offset: 16773},
						name: "CrossReference",
					},
					&ruleRefExpr{
						pos:  position{line: 403, col: 35, offset: 16790},
						name: "Passthrough",
					},
					&ruleRefExpr{
						pos:  position{line: 403, col: 49, offset: 16804},
						name: "InlineImage",
					},
					&ruleRefExpr{
						pos:  position{line: 403, col: 63, offset: 16818},
						name: "QuotedText",
					},
					&ruleRefExpr{
						pos:  position{line: 403, col: 76, offset: 16831},
						name: "Link",
					},
					&ruleRefExpr{
						pos:  position{line: 403, col: 83, offset: 16838},
						name: "DocumentAttributeSubstitution",
					},
					&ruleRefExpr{
						pos:  position{line: 403, col: 115, offset: 16870},
						name: "Word",
					},
				},
			},
		},
		{
			name: "QuotedText",
			pos:  position{line: 408, col: 1, offset: 17115},
			expr: &choiceExpr{
				pos: position{line: 408, col: 15, offset: 17129},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 408, col: 15, offset: 17129},
						name: "BoldText",
					},
					&ruleRefExpr{
						pos:  position{line: 408, col: 26, offset: 17140},
						name: "ItalicText",
					},
					&ruleRefExpr{
						pos:  position{line: 408, col: 39, offset: 17153},
						name: "MonospaceText",
					},
					&ruleRefExpr{
						pos:  position{line: 409, col: 13, offset: 17181},
						name: "EscapedBoldText",
					},
					&ruleRefExpr{
						pos:  position{line: 409, col: 31, offset: 17199},
						name: "EscapedItalicText",
					},
					&ruleRefExpr{
						pos:  position{line: 409, col: 51, offset: 17219},
						name: "EscapedMonospaceText",
					},
				},
			},
		},
		{
			name: "BoldText",
			pos:  position{line: 411, col: 1, offset: 17241},
			expr: &choiceExpr{
				pos: position{line: 411, col: 13, offset: 17253},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 411, col: 13, offset: 17253},
						run: (*parser).callonBoldText2,
						expr: &seqExpr{
							pos: position{line: 411, col: 13, offset: 17253},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 411, col: 13, offset: 17253},
									expr: &litMatcher{
										pos:        position{line: 411, col: 14, offset: 17254},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 411, col: 19, offset: 17259},
									val:        "**",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 411, col: 24, offset: 17264},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 411, col: 33, offset: 17273},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 411, col: 52, offset: 17292},
									val:        "**",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 413, col: 5, offset: 17417},
						run: (*parser).callonBoldText10,
						expr: &seqExpr{
							pos: position{line: 413, col: 5, offset: 17417},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 413, col: 5, offset: 17417},
									expr: &litMatcher{
										pos:        position{line: 413, col: 6, offset: 17418},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 413, col: 11, offset: 17423},
									val:        "**",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 413, col: 16, offset: 17428},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 413, col: 25, offset: 17437},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 413, col: 44, offset: 17456},
									val:        "*",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 416, col: 5, offset: 17621},
						run: (*parser).callonBoldText18,
						expr: &seqExpr{
							pos: position{line: 416, col: 5, offset: 17621},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 416, col: 5, offset: 17621},
									expr: &litMatcher{
										pos:        position{line: 416, col: 6, offset: 17622},
										val:        "\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 416, col: 10, offset: 17626},
									val:        "*",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 416, col: 14, offset: 17630},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 416, col: 23, offset: 17639},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 416, col: 42, offset: 17658},
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
			pos:  position{line: 420, col: 1, offset: 17758},
			expr: &choiceExpr{
				pos: position{line: 420, col: 20, offset: 17777},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 420, col: 20, offset: 17777},
						run: (*parser).callonEscapedBoldText2,
						expr: &seqExpr{
							pos: position{line: 420, col: 20, offset: 17777},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 420, col: 20, offset: 17777},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 420, col: 33, offset: 17790},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 420, col: 33, offset: 17790},
												val:        "\\\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 420, col: 38, offset: 17795},
												expr: &litMatcher{
													pos:        position{line: 420, col: 38, offset: 17795},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 420, col: 44, offset: 17801},
									val:        "**",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 420, col: 49, offset: 17806},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 420, col: 58, offset: 17815},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 420, col: 77, offset: 17834},
									val:        "**",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 422, col: 5, offset: 17989},
						run: (*parser).callonEscapedBoldText13,
						expr: &seqExpr{
							pos: position{line: 422, col: 5, offset: 17989},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 422, col: 5, offset: 17989},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 422, col: 18, offset: 18002},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 422, col: 18, offset: 18002},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 422, col: 22, offset: 18006},
												expr: &litMatcher{
													pos:        position{line: 422, col: 22, offset: 18006},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 422, col: 28, offset: 18012},
									val:        "**",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 422, col: 33, offset: 18017},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 422, col: 42, offset: 18026},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 422, col: 61, offset: 18045},
									val:        "*",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 425, col: 5, offset: 18239},
						run: (*parser).callonEscapedBoldText24,
						expr: &seqExpr{
							pos: position{line: 425, col: 5, offset: 18239},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 425, col: 5, offset: 18239},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 425, col: 18, offset: 18252},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 425, col: 18, offset: 18252},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 425, col: 22, offset: 18256},
												expr: &litMatcher{
													pos:        position{line: 425, col: 22, offset: 18256},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 425, col: 28, offset: 18262},
									val:        "*",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 425, col: 32, offset: 18266},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 425, col: 41, offset: 18275},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 425, col: 60, offset: 18294},
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
			pos:  position{line: 429, col: 1, offset: 18446},
			expr: &choiceExpr{
				pos: position{line: 429, col: 15, offset: 18460},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 429, col: 15, offset: 18460},
						run: (*parser).callonItalicText2,
						expr: &seqExpr{
							pos: position{line: 429, col: 15, offset: 18460},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 429, col: 15, offset: 18460},
									expr: &litMatcher{
										pos:        position{line: 429, col: 16, offset: 18461},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 429, col: 21, offset: 18466},
									val:        "__",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 429, col: 26, offset: 18471},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 429, col: 35, offset: 18480},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 429, col: 54, offset: 18499},
									val:        "__",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 431, col: 5, offset: 18580},
						run: (*parser).callonItalicText10,
						expr: &seqExpr{
							pos: position{line: 431, col: 5, offset: 18580},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 431, col: 5, offset: 18580},
									expr: &litMatcher{
										pos:        position{line: 431, col: 6, offset: 18581},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 431, col: 11, offset: 18586},
									val:        "__",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 431, col: 16, offset: 18591},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 431, col: 25, offset: 18600},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 431, col: 44, offset: 18619},
									val:        "_",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 434, col: 5, offset: 18786},
						run: (*parser).callonItalicText18,
						expr: &seqExpr{
							pos: position{line: 434, col: 5, offset: 18786},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 434, col: 5, offset: 18786},
									expr: &litMatcher{
										pos:        position{line: 434, col: 6, offset: 18787},
										val:        "\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 434, col: 10, offset: 18791},
									val:        "_",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 434, col: 14, offset: 18795},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 434, col: 23, offset: 18804},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 434, col: 42, offset: 18823},
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
			pos:  position{line: 438, col: 1, offset: 18902},
			expr: &choiceExpr{
				pos: position{line: 438, col: 22, offset: 18923},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 438, col: 22, offset: 18923},
						run: (*parser).callonEscapedItalicText2,
						expr: &seqExpr{
							pos: position{line: 438, col: 22, offset: 18923},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 438, col: 22, offset: 18923},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 438, col: 35, offset: 18936},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 438, col: 35, offset: 18936},
												val:        "\\\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 438, col: 40, offset: 18941},
												expr: &litMatcher{
													pos:        position{line: 438, col: 40, offset: 18941},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 438, col: 46, offset: 18947},
									val:        "__",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 438, col: 51, offset: 18952},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 438, col: 60, offset: 18961},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 438, col: 79, offset: 18980},
									val:        "__",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 440, col: 5, offset: 19135},
						run: (*parser).callonEscapedItalicText13,
						expr: &seqExpr{
							pos: position{line: 440, col: 5, offset: 19135},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 440, col: 5, offset: 19135},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 440, col: 18, offset: 19148},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 440, col: 18, offset: 19148},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 440, col: 22, offset: 19152},
												expr: &litMatcher{
													pos:        position{line: 440, col: 22, offset: 19152},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 440, col: 28, offset: 19158},
									val:        "__",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 440, col: 33, offset: 19163},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 440, col: 42, offset: 19172},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 440, col: 61, offset: 19191},
									val:        "_",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 443, col: 5, offset: 19385},
						run: (*parser).callonEscapedItalicText24,
						expr: &seqExpr{
							pos: position{line: 443, col: 5, offset: 19385},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 443, col: 5, offset: 19385},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 443, col: 18, offset: 19398},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 443, col: 18, offset: 19398},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 443, col: 22, offset: 19402},
												expr: &litMatcher{
													pos:        position{line: 443, col: 22, offset: 19402},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 443, col: 28, offset: 19408},
									val:        "_",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 443, col: 32, offset: 19412},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 443, col: 41, offset: 19421},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 443, col: 60, offset: 19440},
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
			pos:  position{line: 447, col: 1, offset: 19592},
			expr: &choiceExpr{
				pos: position{line: 447, col: 18, offset: 19609},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 447, col: 18, offset: 19609},
						run: (*parser).callonMonospaceText2,
						expr: &seqExpr{
							pos: position{line: 447, col: 18, offset: 19609},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 447, col: 18, offset: 19609},
									expr: &litMatcher{
										pos:        position{line: 447, col: 19, offset: 19610},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 447, col: 24, offset: 19615},
									val:        "``",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 447, col: 29, offset: 19620},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 447, col: 38, offset: 19629},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 447, col: 57, offset: 19648},
									val:        "``",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 449, col: 5, offset: 19778},
						run: (*parser).callonMonospaceText10,
						expr: &seqExpr{
							pos: position{line: 449, col: 5, offset: 19778},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 449, col: 5, offset: 19778},
									expr: &litMatcher{
										pos:        position{line: 449, col: 6, offset: 19779},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 449, col: 11, offset: 19784},
									val:        "``",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 449, col: 16, offset: 19789},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 449, col: 25, offset: 19798},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 449, col: 44, offset: 19817},
									val:        "`",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 452, col: 5, offset: 19987},
						run: (*parser).callonMonospaceText18,
						expr: &seqExpr{
							pos: position{line: 452, col: 5, offset: 19987},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 452, col: 5, offset: 19987},
									expr: &litMatcher{
										pos:        position{line: 452, col: 6, offset: 19988},
										val:        "\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 452, col: 10, offset: 19992},
									val:        "`",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 452, col: 14, offset: 19996},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 452, col: 23, offset: 20005},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 452, col: 42, offset: 20024},
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
			pos:  position{line: 456, col: 1, offset: 20151},
			expr: &choiceExpr{
				pos: position{line: 456, col: 25, offset: 20175},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 456, col: 25, offset: 20175},
						run: (*parser).callonEscapedMonospaceText2,
						expr: &seqExpr{
							pos: position{line: 456, col: 25, offset: 20175},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 456, col: 25, offset: 20175},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 456, col: 38, offset: 20188},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 456, col: 38, offset: 20188},
												val:        "\\\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 456, col: 43, offset: 20193},
												expr: &litMatcher{
													pos:        position{line: 456, col: 43, offset: 20193},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 456, col: 49, offset: 20199},
									val:        "``",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 456, col: 54, offset: 20204},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 456, col: 63, offset: 20213},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 456, col: 82, offset: 20232},
									val:        "``",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 458, col: 5, offset: 20387},
						run: (*parser).callonEscapedMonospaceText13,
						expr: &seqExpr{
							pos: position{line: 458, col: 5, offset: 20387},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 458, col: 5, offset: 20387},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 458, col: 18, offset: 20400},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 458, col: 18, offset: 20400},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 458, col: 22, offset: 20404},
												expr: &litMatcher{
													pos:        position{line: 458, col: 22, offset: 20404},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 458, col: 28, offset: 20410},
									val:        "``",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 458, col: 33, offset: 20415},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 458, col: 42, offset: 20424},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 458, col: 61, offset: 20443},
									val:        "`",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 461, col: 5, offset: 20637},
						run: (*parser).callonEscapedMonospaceText24,
						expr: &seqExpr{
							pos: position{line: 461, col: 5, offset: 20637},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 461, col: 5, offset: 20637},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 461, col: 18, offset: 20650},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 461, col: 18, offset: 20650},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 461, col: 22, offset: 20654},
												expr: &litMatcher{
													pos:        position{line: 461, col: 22, offset: 20654},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 461, col: 28, offset: 20660},
									val:        "`",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 461, col: 32, offset: 20664},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 461, col: 41, offset: 20673},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 461, col: 60, offset: 20692},
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
			pos:  position{line: 465, col: 1, offset: 20844},
			expr: &seqExpr{
				pos: position{line: 465, col: 22, offset: 20865},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 465, col: 22, offset: 20865},
						name: "QuotedTextContentElement",
					},
					&zeroOrMoreExpr{
						pos: position{line: 465, col: 47, offset: 20890},
						expr: &seqExpr{
							pos: position{line: 465, col: 48, offset: 20891},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 465, col: 48, offset: 20891},
									expr: &ruleRefExpr{
										pos:  position{line: 465, col: 48, offset: 20891},
										name: "WS",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 465, col: 52, offset: 20895},
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
			pos:  position{line: 467, col: 1, offset: 20923},
			expr: &choiceExpr{
				pos: position{line: 467, col: 29, offset: 20951},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 467, col: 29, offset: 20951},
						name: "QuotedText",
					},
					&ruleRefExpr{
						pos:  position{line: 467, col: 42, offset: 20964},
						name: "QuotedTextWord",
					},
					&ruleRefExpr{
						pos:  position{line: 467, col: 59, offset: 20981},
						name: "WordWithQuotePunctuation",
					},
				},
			},
		},
		{
			name: "QuotedTextWord",
			pos:  position{line: 469, col: 1, offset: 21110},
			expr: &oneOrMoreExpr{
				pos: position{line: 469, col: 19, offset: 21128},
				expr: &seqExpr{
					pos: position{line: 469, col: 20, offset: 21129},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 469, col: 20, offset: 21129},
							expr: &ruleRefExpr{
								pos:  position{line: 469, col: 21, offset: 21130},
								name: "NEWLINE",
							},
						},
						&notExpr{
							pos: position{line: 469, col: 29, offset: 21138},
							expr: &ruleRefExpr{
								pos:  position{line: 469, col: 30, offset: 21139},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 469, col: 33, offset: 21142},
							expr: &litMatcher{
								pos:        position{line: 469, col: 34, offset: 21143},
								val:        "*",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 469, col: 38, offset: 21147},
							expr: &litMatcher{
								pos:        position{line: 469, col: 39, offset: 21148},
								val:        "_",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 469, col: 43, offset: 21152},
							expr: &litMatcher{
								pos:        position{line: 469, col: 44, offset: 21153},
								val:        "`",
								ignoreCase: false,
							},
						},
						&anyMatcher{
							line: 469, col: 48, offset: 21157,
						},
					},
				},
			},
		},
		{
			name: "WordWithQuotePunctuation",
			pos:  position{line: 471, col: 1, offset: 21200},
			expr: &actionExpr{
				pos: position{line: 471, col: 29, offset: 21228},
				run: (*parser).callonWordWithQuotePunctuation1,
				expr: &oneOrMoreExpr{
					pos: position{line: 471, col: 29, offset: 21228},
					expr: &seqExpr{
						pos: position{line: 471, col: 30, offset: 21229},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 471, col: 30, offset: 21229},
								expr: &ruleRefExpr{
									pos:  position{line: 471, col: 31, offset: 21230},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 471, col: 39, offset: 21238},
								expr: &ruleRefExpr{
									pos:  position{line: 471, col: 40, offset: 21239},
									name: "WS",
								},
							},
							&anyMatcher{
								line: 471, col: 44, offset: 21243,
							},
						},
					},
				},
			},
		},
		{
			name: "UnbalancedQuotePunctuation",
			pos:  position{line: 476, col: 1, offset: 21488},
			expr: &choiceExpr{
				pos: position{line: 476, col: 31, offset: 21518},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 476, col: 31, offset: 21518},
						val:        "*",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 476, col: 37, offset: 21524},
						val:        "_",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 476, col: 43, offset: 21530},
						val:        "`",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "Passthrough",
			pos:  position{line: 481, col: 1, offset: 21642},
			expr: &choiceExpr{
				pos: position{line: 481, col: 16, offset: 21657},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 481, col: 16, offset: 21657},
						name: "TriplePlusPassthrough",
					},
					&ruleRefExpr{
						pos:  position{line: 481, col: 40, offset: 21681},
						name: "SinglePlusPassthrough",
					},
					&ruleRefExpr{
						pos:  position{line: 481, col: 64, offset: 21705},
						name: "PassthroughMacro",
					},
				},
			},
		},
		{
			name: "SinglePlusPassthrough",
			pos:  position{line: 483, col: 1, offset: 21723},
			expr: &actionExpr{
				pos: position{line: 483, col: 26, offset: 21748},
				run: (*parser).callonSinglePlusPassthrough1,
				expr: &seqExpr{
					pos: position{line: 483, col: 26, offset: 21748},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 483, col: 26, offset: 21748},
							val:        "+",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 483, col: 30, offset: 21752},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 483, col: 38, offset: 21760},
								expr: &seqExpr{
									pos: position{line: 483, col: 39, offset: 21761},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 483, col: 39, offset: 21761},
											expr: &ruleRefExpr{
												pos:  position{line: 483, col: 40, offset: 21762},
												name: "NEWLINE",
											},
										},
										&notExpr{
											pos: position{line: 483, col: 48, offset: 21770},
											expr: &litMatcher{
												pos:        position{line: 483, col: 49, offset: 21771},
												val:        "+",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 483, col: 53, offset: 21775,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 483, col: 57, offset: 21779},
							val:        "+",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TriplePlusPassthrough",
			pos:  position{line: 487, col: 1, offset: 21874},
			expr: &actionExpr{
				pos: position{line: 487, col: 26, offset: 21899},
				run: (*parser).callonTriplePlusPassthrough1,
				expr: &seqExpr{
					pos: position{line: 487, col: 26, offset: 21899},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 487, col: 26, offset: 21899},
							val:        "+++",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 487, col: 32, offset: 21905},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 487, col: 40, offset: 21913},
								expr: &seqExpr{
									pos: position{line: 487, col: 41, offset: 21914},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 487, col: 41, offset: 21914},
											expr: &litMatcher{
												pos:        position{line: 487, col: 42, offset: 21915},
												val:        "+++",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 487, col: 48, offset: 21921,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 487, col: 52, offset: 21925},
							val:        "+++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "PassthroughMacro",
			pos:  position{line: 491, col: 1, offset: 22022},
			expr: &choiceExpr{
				pos: position{line: 491, col: 21, offset: 22042},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 491, col: 21, offset: 22042},
						run: (*parser).callonPassthroughMacro2,
						expr: &seqExpr{
							pos: position{line: 491, col: 21, offset: 22042},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 491, col: 21, offset: 22042},
									val:        "pass:[",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 491, col: 30, offset: 22051},
									label: "content",
									expr: &zeroOrMoreExpr{
										pos: position{line: 491, col: 38, offset: 22059},
										expr: &ruleRefExpr{
											pos:  position{line: 491, col: 39, offset: 22060},
											name: "PassthroughMacroCharacter",
										},
									},
								},
								&litMatcher{
									pos:        position{line: 491, col: 67, offset: 22088},
									val:        "]",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 493, col: 5, offset: 22179},
						run: (*parser).callonPassthroughMacro9,
						expr: &seqExpr{
							pos: position{line: 493, col: 5, offset: 22179},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 493, col: 5, offset: 22179},
									val:        "pass:q[",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 493, col: 15, offset: 22189},
									label: "content",
									expr: &zeroOrMoreExpr{
										pos: position{line: 493, col: 23, offset: 22197},
										expr: &choiceExpr{
											pos: position{line: 493, col: 24, offset: 22198},
											alternatives: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 493, col: 24, offset: 22198},
													name: "QuotedText",
												},
												&ruleRefExpr{
													pos:  position{line: 493, col: 37, offset: 22211},
													name: "PassthroughMacroCharacter",
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 493, col: 65, offset: 22239},
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
			pos:  position{line: 497, col: 1, offset: 22329},
			expr: &seqExpr{
				pos: position{line: 497, col: 31, offset: 22359},
				exprs: []interface{}{
					&notExpr{
						pos: position{line: 497, col: 31, offset: 22359},
						expr: &litMatcher{
							pos:        position{line: 497, col: 32, offset: 22360},
							val:        "]",
							ignoreCase: false,
						},
					},
					&anyMatcher{
						line: 497, col: 36, offset: 22364,
					},
				},
			},
		},
		{
			name: "CrossReference",
			pos:  position{line: 502, col: 1, offset: 22480},
			expr: &actionExpr{
				pos: position{line: 502, col: 19, offset: 22498},
				run: (*parser).callonCrossReference1,
				expr: &seqExpr{
					pos: position{line: 502, col: 19, offset: 22498},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 502, col: 19, offset: 22498},
							val:        "<<",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 502, col: 24, offset: 22503},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 502, col: 28, offset: 22507},
								name: "ID",
							},
						},
						&litMatcher{
							pos:        position{line: 502, col: 32, offset: 22511},
							val:        ">>",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Link",
			pos:  position{line: 509, col: 1, offset: 22670},
			expr: &choiceExpr{
				pos: position{line: 509, col: 9, offset: 22678},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 509, col: 9, offset: 22678},
						name: "RelativeLink",
					},
					&ruleRefExpr{
						pos:  position{line: 509, col: 24, offset: 22693},
						name: "ExternalLink",
					},
				},
			},
		},
		{
			name: "ExternalLink",
			pos:  position{line: 511, col: 1, offset: 22708},
			expr: &actionExpr{
				pos: position{line: 511, col: 17, offset: 22724},
				run: (*parser).callonExternalLink1,
				expr: &seqExpr{
					pos: position{line: 511, col: 17, offset: 22724},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 511, col: 17, offset: 22724},
							label: "url",
							expr: &seqExpr{
								pos: position{line: 511, col: 22, offset: 22729},
								exprs: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 511, col: 22, offset: 22729},
										name: "URL_SCHEME",
									},
									&ruleRefExpr{
										pos:  position{line: 511, col: 33, offset: 22740},
										name: "URL",
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 511, col: 38, offset: 22745},
							label: "text",
							expr: &zeroOrOneExpr{
								pos: position{line: 511, col: 43, offset: 22750},
								expr: &seqExpr{
									pos: position{line: 511, col: 44, offset: 22751},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 511, col: 44, offset: 22751},
											val:        "[",
											ignoreCase: false,
										},
										&zeroOrMoreExpr{
											pos: position{line: 511, col: 48, offset: 22755},
											expr: &ruleRefExpr{
												pos:  position{line: 511, col: 49, offset: 22756},
												name: "URL_TEXT",
											},
										},
										&litMatcher{
											pos:        position{line: 511, col: 60, offset: 22767},
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
			pos:  position{line: 518, col: 1, offset: 22928},
			expr: &actionExpr{
				pos: position{line: 518, col: 17, offset: 22944},
				run: (*parser).callonRelativeLink1,
				expr: &seqExpr{
					pos: position{line: 518, col: 17, offset: 22944},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 518, col: 17, offset: 22944},
							val:        "link:",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 518, col: 25, offset: 22952},
							label: "url",
							expr: &seqExpr{
								pos: position{line: 518, col: 30, offset: 22957},
								exprs: []interface{}{
									&zeroOrOneExpr{
										pos: position{line: 518, col: 30, offset: 22957},
										expr: &ruleRefExpr{
											pos:  position{line: 518, col: 30, offset: 22957},
											name: "URL_SCHEME",
										},
									},
									&ruleRefExpr{
										pos:  position{line: 518, col: 42, offset: 22969},
										name: "URL",
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 518, col: 47, offset: 22974},
							label: "text",
							expr: &seqExpr{
								pos: position{line: 518, col: 53, offset: 22980},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 518, col: 53, offset: 22980},
										val:        "[",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 518, col: 57, offset: 22984},
										expr: &ruleRefExpr{
											pos:  position{line: 518, col: 58, offset: 22985},
											name: "URL_TEXT",
										},
									},
									&litMatcher{
										pos:        position{line: 518, col: 69, offset: 22996},
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
			pos:  position{line: 528, col: 1, offset: 23258},
			expr: &actionExpr{
				pos: position{line: 528, col: 15, offset: 23272},
				run: (*parser).callonBlockImage1,
				expr: &seqExpr{
					pos: position{line: 528, col: 15, offset: 23272},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 528, col: 15, offset: 23272},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 528, col: 26, offset: 23283},
								expr: &ruleRefExpr{
									pos:  position{line: 528, col: 27, offset: 23284},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 528, col: 46, offset: 23303},
							label: "image",
							expr: &ruleRefExpr{
								pos:  position{line: 528, col: 52, offset: 23309},
								name: "BlockImageMacro",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 528, col: 69, offset: 23326},
							expr: &ruleRefExpr{
								pos:  position{line: 528, col: 69, offset: 23326},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 528, col: 73, offset: 23330},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "BlockImageMacro",
			pos:  position{line: 533, col: 1, offset: 23489},
			expr: &actionExpr{
				pos: position{line: 533, col: 20, offset: 23508},
				run: (*parser).callonBlockImageMacro1,
				expr: &seqExpr{
					pos: position{line: 533, col: 20, offset: 23508},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 533, col: 20, offset: 23508},
							val:        "image::",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 533, col: 30, offset: 23518},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 533, col: 36, offset: 23524},
								name: "URL",
							},
						},
						&litMatcher{
							pos:        position{line: 533, col: 41, offset: 23529},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 533, col: 45, offset: 23533},
							label: "attributes",
							expr: &zeroOrOneExpr{
								pos: position{line: 533, col: 57, offset: 23545},
								expr: &ruleRefExpr{
									pos:  position{line: 533, col: 57, offset: 23545},
									name: "URL_TEXT",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 533, col: 68, offset: 23556},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "InlineImage",
			pos:  position{line: 537, col: 1, offset: 23623},
			expr: &actionExpr{
				pos: position{line: 537, col: 16, offset: 23638},
				run: (*parser).callonInlineImage1,
				expr: &labeledExpr{
					pos:   position{line: 537, col: 16, offset: 23638},
					label: "image",
					expr: &ruleRefExpr{
						pos:  position{line: 537, col: 22, offset: 23644},
						name: "InlineImageMacro",
					},
				},
			},
		},
		{
			name: "InlineImageMacro",
			pos:  position{line: 542, col: 1, offset: 23789},
			expr: &actionExpr{
				pos: position{line: 542, col: 21, offset: 23809},
				run: (*parser).callonInlineImageMacro1,
				expr: &seqExpr{
					pos: position{line: 542, col: 21, offset: 23809},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 542, col: 21, offset: 23809},
							val:        "image:",
							ignoreCase: false,
						},
						&notExpr{
							pos: position{line: 542, col: 30, offset: 23818},
							expr: &litMatcher{
								pos:        position{line: 542, col: 31, offset: 23819},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 542, col: 35, offset: 23823},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 542, col: 41, offset: 23829},
								name: "URL",
							},
						},
						&litMatcher{
							pos:        position{line: 542, col: 46, offset: 23834},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 542, col: 50, offset: 23838},
							label: "attributes",
							expr: &zeroOrOneExpr{
								pos: position{line: 542, col: 62, offset: 23850},
								expr: &ruleRefExpr{
									pos:  position{line: 542, col: 62, offset: 23850},
									name: "URL_TEXT",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 542, col: 73, offset: 23861},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "DelimitedBlock",
			pos:  position{line: 549, col: 1, offset: 24191},
			expr: &choiceExpr{
				pos: position{line: 549, col: 19, offset: 24209},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 549, col: 19, offset: 24209},
						name: "FencedBlock",
					},
					&ruleRefExpr{
						pos:  position{line: 549, col: 33, offset: 24223},
						name: "ListingBlock",
					},
					&ruleRefExpr{
						pos:  position{line: 549, col: 48, offset: 24238},
						name: "ExampleBlock",
					},
				},
			},
		},
		{
			name: "BlockDelimiter",
			pos:  position{line: 551, col: 1, offset: 24252},
			expr: &choiceExpr{
				pos: position{line: 551, col: 19, offset: 24270},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 551, col: 19, offset: 24270},
						name: "LiteralBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 551, col: 43, offset: 24294},
						name: "FencedBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 551, col: 66, offset: 24317},
						name: "ListingBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 551, col: 90, offset: 24341},
						name: "ExampleBlockDelimiter",
					},
				},
			},
		},
		{
			name: "FencedBlockDelimiter",
			pos:  position{line: 553, col: 1, offset: 24364},
			expr: &litMatcher{
				pos:        position{line: 553, col: 25, offset: 24388},
				val:        "```",
				ignoreCase: false,
			},
		},
		{
			name: "FencedBlock",
			pos:  position{line: 555, col: 1, offset: 24395},
			expr: &actionExpr{
				pos: position{line: 555, col: 16, offset: 24410},
				run: (*parser).callonFencedBlock1,
				expr: &seqExpr{
					pos: position{line: 555, col: 16, offset: 24410},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 555, col: 16, offset: 24410},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 555, col: 27, offset: 24421},
								expr: &ruleRefExpr{
									pos:  position{line: 555, col: 28, offset: 24422},
									name: "ElementAttribute",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 555, col: 47, offset: 24441},
							name: "FencedBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 555, col: 68, offset: 24462},
							expr: &ruleRefExpr{
								pos:  position{line: 555, col: 68, offset: 24462},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 555, col: 72, offset: 24466},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 555, col: 80, offset: 24474},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 555, col: 88, offset: 24482},
								expr: &choiceExpr{
									pos: position{line: 555, col: 89, offset: 24483},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 555, col: 89, offset: 24483},
											name: "List",
										},
										&ruleRefExpr{
											pos:  position{line: 555, col: 96, offset: 24490},
											name: "InnerParagraph",
										},
										&ruleRefExpr{
											pos:  position{line: 555, col: 113, offset: 24507},
											name: "BlankLine",
										},
									},
								},
							},
						},
						&choiceExpr{
							pos: position{line: 555, col: 126, offset: 24520},
							alternatives: []interface{}{
								&seqExpr{
									pos: position{line: 555, col: 127, offset: 24521},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 555, col: 127, offset: 24521},
											name: "FencedBlockDelimiter",
										},
										&zeroOrMoreExpr{
											pos: position{line: 555, col: 148, offset: 24542},
											expr: &ruleRefExpr{
												pos:  position{line: 555, col: 148, offset: 24542},
												name: "WS",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 555, col: 152, offset: 24546},
											name: "EOL",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 555, col: 159, offset: 24553},
									name: "EOF",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "ListingBlockDelimiter",
			pos:  position{line: 559, col: 1, offset: 24670},
			expr: &litMatcher{
				pos:        position{line: 559, col: 26, offset: 24695},
				val:        "----",
				ignoreCase: false,
			},
		},
		{
			name: "ListingBlock",
			pos:  position{line: 561, col: 1, offset: 24703},
			expr: &actionExpr{
				pos: position{line: 561, col: 17, offset: 24719},
				run: (*parser).callonListingBlock1,
				expr: &seqExpr{
					pos: position{line: 561, col: 17, offset: 24719},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 561, col: 17, offset: 24719},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 561, col: 28, offset: 24730},
								expr: &ruleRefExpr{
									pos:  position{line: 561, col: 29, offset: 24731},
									name: "ElementAttribute",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 561, col: 48, offset: 24750},
							name: "ListingBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 561, col: 70, offset: 24772},
							expr: &ruleRefExpr{
								pos:  position{line: 561, col: 70, offset: 24772},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 561, col: 74, offset: 24776},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 561, col: 82, offset: 24784},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 561, col: 90, offset: 24792},
								expr: &choiceExpr{
									pos: position{line: 561, col: 91, offset: 24793},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 561, col: 91, offset: 24793},
											name: "List",
										},
										&ruleRefExpr{
											pos:  position{line: 561, col: 98, offset: 24800},
											name: "InnerParagraph",
										},
										&ruleRefExpr{
											pos:  position{line: 561, col: 115, offset: 24817},
											name: "BlankLine",
										},
									},
								},
							},
						},
						&choiceExpr{
							pos: position{line: 561, col: 128, offset: 24830},
							alternatives: []interface{}{
								&seqExpr{
									pos: position{line: 561, col: 129, offset: 24831},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 561, col: 129, offset: 24831},
											name: "ListingBlockDelimiter",
										},
										&zeroOrMoreExpr{
											pos: position{line: 561, col: 151, offset: 24853},
											expr: &ruleRefExpr{
												pos:  position{line: 561, col: 151, offset: 24853},
												name: "WS",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 561, col: 155, offset: 24857},
											name: "EOL",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 561, col: 162, offset: 24864},
									name: "EOF",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "ExampleBlockDelimiter",
			pos:  position{line: 565, col: 1, offset: 24982},
			expr: &litMatcher{
				pos:        position{line: 565, col: 26, offset: 25007},
				val:        "====",
				ignoreCase: false,
			},
		},
		{
			name: "ExampleBlock",
			pos:  position{line: 567, col: 1, offset: 25015},
			expr: &actionExpr{
				pos: position{line: 567, col: 17, offset: 25031},
				run: (*parser).callonExampleBlock1,
				expr: &seqExpr{
					pos: position{line: 567, col: 17, offset: 25031},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 567, col: 17, offset: 25031},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 567, col: 28, offset: 25042},
								expr: &ruleRefExpr{
									pos:  position{line: 567, col: 29, offset: 25043},
									name: "ElementAttribute",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 567, col: 48, offset: 25062},
							name: "ExampleBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 567, col: 70, offset: 25084},
							expr: &ruleRefExpr{
								pos:  position{line: 567, col: 70, offset: 25084},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 567, col: 74, offset: 25088},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 567, col: 82, offset: 25096},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 567, col: 90, offset: 25104},
								expr: &choiceExpr{
									pos: position{line: 567, col: 91, offset: 25105},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 567, col: 91, offset: 25105},
											name: "List",
										},
										&ruleRefExpr{
											pos:  position{line: 567, col: 98, offset: 25112},
											name: "InnerParagraph",
										},
										&ruleRefExpr{
											pos:  position{line: 567, col: 115, offset: 25129},
											name: "BlankLine",
										},
									},
								},
							},
						},
						&choiceExpr{
							pos: position{line: 567, col: 129, offset: 25143},
							alternatives: []interface{}{
								&seqExpr{
									pos: position{line: 567, col: 130, offset: 25144},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 567, col: 130, offset: 25144},
											name: "ExampleBlockDelimiter",
										},
										&zeroOrMoreExpr{
											pos: position{line: 567, col: 152, offset: 25166},
											expr: &ruleRefExpr{
												pos:  position{line: 567, col: 152, offset: 25166},
												name: "WS",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 567, col: 156, offset: 25170},
											name: "EOL",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 567, col: 163, offset: 25177},
									name: "EOF",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "LiteralBlock",
			pos:  position{line: 574, col: 1, offset: 25562},
			expr: &choiceExpr{
				pos: position{line: 574, col: 17, offset: 25578},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 574, col: 17, offset: 25578},
						name: "ParagraphWithSpaces",
					},
					&ruleRefExpr{
						pos:  position{line: 574, col: 39, offset: 25600},
						name: "ParagraphWithLiteralBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 574, col: 76, offset: 25637},
						name: "ParagraphWithLiteralAttribute",
					},
				},
			},
		},
		{
			name: "ParagraphWithSpaces",
			pos:  position{line: 577, col: 1, offset: 25732},
			expr: &actionExpr{
				pos: position{line: 577, col: 24, offset: 25755},
				run: (*parser).callonParagraphWithSpaces1,
				expr: &seqExpr{
					pos: position{line: 577, col: 24, offset: 25755},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 577, col: 24, offset: 25755},
							label: "spaces",
							expr: &oneOrMoreExpr{
								pos: position{line: 577, col: 32, offset: 25763},
								expr: &ruleRefExpr{
									pos:  position{line: 577, col: 32, offset: 25763},
									name: "WS",
								},
							},
						},
						&notExpr{
							pos: position{line: 577, col: 37, offset: 25768},
							expr: &ruleRefExpr{
								pos:  position{line: 577, col: 38, offset: 25769},
								name: "NEWLINE",
							},
						},
						&labeledExpr{
							pos:   position{line: 577, col: 46, offset: 25777},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 577, col: 55, offset: 25786},
								name: "LiteralBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 577, col: 76, offset: 25807},
							name: "EndOfLiteralBlock",
						},
					},
				},
			},
		},
		{
			name: "LiteralBlockContent",
			pos:  position{line: 582, col: 1, offset: 25988},
			expr: &actionExpr{
				pos: position{line: 582, col: 24, offset: 26011},
				run: (*parser).callonLiteralBlockContent1,
				expr: &labeledExpr{
					pos:   position{line: 582, col: 24, offset: 26011},
					label: "content",
					expr: &oneOrMoreExpr{
						pos: position{line: 582, col: 32, offset: 26019},
						expr: &seqExpr{
							pos: position{line: 582, col: 33, offset: 26020},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 582, col: 33, offset: 26020},
									expr: &seqExpr{
										pos: position{line: 582, col: 35, offset: 26022},
										exprs: []interface{}{
											&ruleRefExpr{
												pos:  position{line: 582, col: 35, offset: 26022},
												name: "NEWLINE",
											},
											&ruleRefExpr{
												pos:  position{line: 582, col: 43, offset: 26030},
												name: "BlankLine",
											},
										},
									},
								},
								&anyMatcher{
									line: 582, col: 54, offset: 26041,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "EndOfLiteralBlock",
			pos:  position{line: 587, col: 1, offset: 26126},
			expr: &choiceExpr{
				pos: position{line: 587, col: 22, offset: 26147},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 587, col: 22, offset: 26147},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 587, col: 22, offset: 26147},
								name: "NEWLINE",
							},
							&ruleRefExpr{
								pos:  position{line: 587, col: 30, offset: 26155},
								name: "BlankLine",
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 587, col: 42, offset: 26167},
						name: "NEWLINE",
					},
					&ruleRefExpr{
						pos:  position{line: 587, col: 52, offset: 26177},
						name: "EOF",
					},
				},
			},
		},
		{
			name: "ParagraphWithLiteralBlockDelimiter",
			pos:  position{line: 590, col: 1, offset: 26237},
			expr: &actionExpr{
				pos: position{line: 590, col: 39, offset: 26275},
				run: (*parser).callonParagraphWithLiteralBlockDelimiter1,
				expr: &seqExpr{
					pos: position{line: 590, col: 39, offset: 26275},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 590, col: 39, offset: 26275},
							name: "LiteralBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 590, col: 61, offset: 26297},
							expr: &ruleRefExpr{
								pos:  position{line: 590, col: 61, offset: 26297},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 590, col: 65, offset: 26301},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 590, col: 73, offset: 26309},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 590, col: 81, offset: 26317},
								expr: &seqExpr{
									pos: position{line: 590, col: 82, offset: 26318},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 590, col: 82, offset: 26318},
											expr: &ruleRefExpr{
												pos:  position{line: 590, col: 83, offset: 26319},
												name: "LiteralBlockDelimiter",
											},
										},
										&anyMatcher{
											line: 590, col: 105, offset: 26341,
										},
									},
								},
							},
						},
						&choiceExpr{
							pos: position{line: 590, col: 110, offset: 26346},
							alternatives: []interface{}{
								&seqExpr{
									pos: position{line: 590, col: 111, offset: 26347},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 590, col: 111, offset: 26347},
											name: "LiteralBlockDelimiter",
										},
										&zeroOrMoreExpr{
											pos: position{line: 590, col: 133, offset: 26369},
											expr: &ruleRefExpr{
												pos:  position{line: 590, col: 133, offset: 26369},
												name: "WS",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 590, col: 137, offset: 26373},
											name: "EOL",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 590, col: 144, offset: 26380},
									name: "EOF",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "LiteralBlockDelimiter",
			pos:  position{line: 594, col: 1, offset: 26465},
			expr: &litMatcher{
				pos:        position{line: 594, col: 26, offset: 26490},
				val:        "....",
				ignoreCase: false,
			},
		},
		{
			name: "ParagraphWithLiteralAttribute",
			pos:  position{line: 597, col: 1, offset: 26552},
			expr: &actionExpr{
				pos: position{line: 597, col: 34, offset: 26585},
				run: (*parser).callonParagraphWithLiteralAttribute1,
				expr: &seqExpr{
					pos: position{line: 597, col: 34, offset: 26585},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 597, col: 34, offset: 26585},
							val:        "[literal]",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 597, col: 46, offset: 26597},
							expr: &ruleRefExpr{
								pos:  position{line: 597, col: 46, offset: 26597},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 597, col: 50, offset: 26601},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 597, col: 58, offset: 26609},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 597, col: 67, offset: 26618},
								name: "LiteralBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 597, col: 88, offset: 26639},
							name: "EndOfLiteralBlock",
						},
					},
				},
			},
		},
		{
			name: "BlankLine",
			pos:  position{line: 604, col: 1, offset: 26842},
			expr: &actionExpr{
				pos: position{line: 604, col: 14, offset: 26855},
				run: (*parser).callonBlankLine1,
				expr: &seqExpr{
					pos: position{line: 604, col: 14, offset: 26855},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 604, col: 14, offset: 26855},
							expr: &ruleRefExpr{
								pos:  position{line: 604, col: 15, offset: 26856},
								name: "EOF",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 604, col: 19, offset: 26860},
							expr: &ruleRefExpr{
								pos:  position{line: 604, col: 19, offset: 26860},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 604, col: 23, offset: 26864},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "Word",
			pos:  position{line: 611, col: 1, offset: 27011},
			expr: &actionExpr{
				pos: position{line: 611, col: 9, offset: 27019},
				run: (*parser).callonWord1,
				expr: &oneOrMoreExpr{
					pos: position{line: 611, col: 9, offset: 27019},
					expr: &seqExpr{
						pos: position{line: 611, col: 10, offset: 27020},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 611, col: 10, offset: 27020},
								expr: &ruleRefExpr{
									pos:  position{line: 611, col: 11, offset: 27021},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 611, col: 19, offset: 27029},
								expr: &ruleRefExpr{
									pos:  position{line: 611, col: 20, offset: 27030},
									name: "WS",
								},
							},
							&anyMatcher{
								line: 611, col: 23, offset: 27033,
							},
						},
					},
				},
			},
		},
		{
			name: "URL",
			pos:  position{line: 615, col: 1, offset: 27073},
			expr: &actionExpr{
				pos: position{line: 615, col: 8, offset: 27080},
				run: (*parser).callonURL1,
				expr: &oneOrMoreExpr{
					pos: position{line: 615, col: 8, offset: 27080},
					expr: &seqExpr{
						pos: position{line: 615, col: 9, offset: 27081},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 615, col: 9, offset: 27081},
								expr: &ruleRefExpr{
									pos:  position{line: 615, col: 10, offset: 27082},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 615, col: 18, offset: 27090},
								expr: &ruleRefExpr{
									pos:  position{line: 615, col: 19, offset: 27091},
									name: "WS",
								},
							},
							&notExpr{
								pos: position{line: 615, col: 22, offset: 27094},
								expr: &litMatcher{
									pos:        position{line: 615, col: 23, offset: 27095},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 615, col: 27, offset: 27099},
								expr: &litMatcher{
									pos:        position{line: 615, col: 28, offset: 27100},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 615, col: 32, offset: 27104,
							},
						},
					},
				},
			},
		},
		{
			name: "ID",
			pos:  position{line: 619, col: 1, offset: 27144},
			expr: &actionExpr{
				pos: position{line: 619, col: 7, offset: 27150},
				run: (*parser).callonID1,
				expr: &oneOrMoreExpr{
					pos: position{line: 619, col: 7, offset: 27150},
					expr: &seqExpr{
						pos: position{line: 619, col: 8, offset: 27151},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 619, col: 8, offset: 27151},
								expr: &ruleRefExpr{
									pos:  position{line: 619, col: 9, offset: 27152},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 619, col: 17, offset: 27160},
								expr: &ruleRefExpr{
									pos:  position{line: 619, col: 18, offset: 27161},
									name: "WS",
								},
							},
							&notExpr{
								pos: position{line: 619, col: 21, offset: 27164},
								expr: &litMatcher{
									pos:        position{line: 619, col: 22, offset: 27165},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 619, col: 26, offset: 27169},
								expr: &litMatcher{
									pos:        position{line: 619, col: 27, offset: 27170},
									val:        "]",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 619, col: 31, offset: 27174},
								expr: &litMatcher{
									pos:        position{line: 619, col: 32, offset: 27175},
									val:        "<<",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 619, col: 37, offset: 27180},
								expr: &litMatcher{
									pos:        position{line: 619, col: 38, offset: 27181},
									val:        ">>",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 619, col: 42, offset: 27185,
							},
						},
					},
				},
			},
		},
		{
			name: "URL_TEXT",
			pos:  position{line: 623, col: 1, offset: 27225},
			expr: &actionExpr{
				pos: position{line: 623, col: 13, offset: 27237},
				run: (*parser).callonURL_TEXT1,
				expr: &oneOrMoreExpr{
					pos: position{line: 623, col: 13, offset: 27237},
					expr: &seqExpr{
						pos: position{line: 623, col: 14, offset: 27238},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 623, col: 14, offset: 27238},
								expr: &ruleRefExpr{
									pos:  position{line: 623, col: 15, offset: 27239},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 623, col: 23, offset: 27247},
								expr: &litMatcher{
									pos:        position{line: 623, col: 24, offset: 27248},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 623, col: 28, offset: 27252},
								expr: &litMatcher{
									pos:        position{line: 623, col: 29, offset: 27253},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 623, col: 33, offset: 27257,
							},
						},
					},
				},
			},
		},
		{
			name: "URL_SCHEME",
			pos:  position{line: 627, col: 1, offset: 27297},
			expr: &choiceExpr{
				pos: position{line: 627, col: 15, offset: 27311},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 627, col: 15, offset: 27311},
						val:        "http://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 627, col: 27, offset: 27323},
						val:        "https://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 627, col: 40, offset: 27336},
						val:        "ftp://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 627, col: 51, offset: 27347},
						val:        "irc://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 627, col: 62, offset: 27358},
						val:        "mailto:",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "DIGIT",
			pos:  position{line: 629, col: 1, offset: 27369},
			expr: &charClassMatcher{
				pos:        position{line: 629, col: 10, offset: 27378},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NEWLINE",
			pos:  position{line: 631, col: 1, offset: 27385},
			expr: &choiceExpr{
				pos: position{line: 631, col: 12, offset: 27396},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 631, col: 12, offset: 27396},
						val:        "\r\n",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 631, col: 21, offset: 27405},
						val:        "\r",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 631, col: 28, offset: 27412},
						val:        "\n",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "WS",
			pos:  position{line: 633, col: 1, offset: 27418},
			expr: &choiceExpr{
				pos: position{line: 633, col: 7, offset: 27424},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 633, col: 7, offset: 27424},
						val:        " ",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 633, col: 13, offset: 27430},
						run: (*parser).callonWS3,
						expr: &litMatcher{
							pos:        position{line: 633, col: 13, offset: 27430},
							val:        "\t",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 637, col: 1, offset: 27475},
			expr: &notExpr{
				pos: position{line: 637, col: 8, offset: 27482},
				expr: &anyMatcher{
					line: 637, col: 9, offset: 27483,
				},
			},
		},
		{
			name: "EOL",
			pos:  position{line: 639, col: 1, offset: 27486},
			expr: &choiceExpr{
				pos: position{line: 639, col: 8, offset: 27493},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 639, col: 8, offset: 27493},
						name: "NEWLINE",
					},
					&ruleRefExpr{
						pos:  position{line: 639, col: 18, offset: 27503},
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

func (c *current) onDocumentBlocks8(content interface{}) (interface{}, error) {

	return content, nil
}

func (p *parser) callonDocumentBlocks8() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDocumentBlocks8(stack["content"])
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

func (c *current) onSection1Title1(attributes, level, content, id interface{}) (interface{}, error) {

	return types.NewSectionTitle(content.(types.InlineElements), append(attributes.([]interface{}), id))
}

func (p *parser) callonSection1Title1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection1Title1(stack["attributes"], stack["level"], stack["content"], stack["id"])
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

func (c *current) onSection2Title1(attributes, level, content, id interface{}) (interface{}, error) {
	return types.NewSectionTitle(content.(types.InlineElements), append(attributes.([]interface{}), id))
}

func (p *parser) callonSection2Title1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection2Title1(stack["attributes"], stack["level"], stack["content"], stack["id"])
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

func (c *current) onSection3Title1(attributes, level, content, id interface{}) (interface{}, error) {
	return types.NewSectionTitle(content.(types.InlineElements), append(attributes.([]interface{}), id))
}

func (p *parser) callonSection3Title1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection3Title1(stack["attributes"], stack["level"], stack["content"], stack["id"])
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

func (c *current) onSection4Title1(attributes, level, content, id interface{}) (interface{}, error) {
	return types.NewSectionTitle(content.(types.InlineElements), append(attributes.([]interface{}), id))
}

func (p *parser) callonSection4Title1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection4Title1(stack["attributes"], stack["level"], stack["content"], stack["id"])
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

func (c *current) onSection5Title1(attributes, level, content, id interface{}) (interface{}, error) {
	return types.NewSectionTitle(content.(types.InlineElements), append(attributes.([]interface{}), id))
}

func (p *parser) callonSection5Title1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection5Title1(stack["attributes"], stack["level"], stack["content"], stack["id"])
}

func (c *current) onSection5Block1(content interface{}) (interface{}, error) {
	return content, nil
}

func (p *parser) callonSection5Block1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSection5Block1(stack["content"])
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

func (c *current) onInnerParagraph1(lines interface{}) (interface{}, error) {
	return types.NewParagraph(lines.([]interface{}), nil)
}

func (p *parser) callonInnerParagraph1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInnerParagraph1(stack["lines"])
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
	// ignore whitespaces, only return the relevant "*"/"-" Word
	return types.NewUnorderedListItemPrefix(types.FiveAsterisks, 5)
}

func (p *parser) callonUnorderedListItemPrefix2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUnorderedListItemPrefix2(stack["level"])
}

func (c *current) onUnorderedListItemPrefix10(level interface{}) (interface{}, error) {
	// ignore whitespaces, only return the relevant "*"/"-" Word
	return types.NewUnorderedListItemPrefix(types.FourAsterisks, 4)
}

func (p *parser) callonUnorderedListItemPrefix10() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUnorderedListItemPrefix10(stack["level"])
}

func (c *current) onUnorderedListItemPrefix18(level interface{}) (interface{}, error) {
	// ignore whitespaces, only return the relevant "*"/"-" Word
	return types.NewUnorderedListItemPrefix(types.ThreeAsterisks, 3)
}

func (p *parser) callonUnorderedListItemPrefix18() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUnorderedListItemPrefix18(stack["level"])
}

func (c *current) onUnorderedListItemPrefix26(level interface{}) (interface{}, error) {
	// ignore whitespaces, only return the relevant "*"/"-" Word
	return types.NewUnorderedListItemPrefix(types.TwoAsterisks, 2)
}

func (p *parser) callonUnorderedListItemPrefix26() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUnorderedListItemPrefix26(stack["level"])
}

func (c *current) onUnorderedListItemPrefix34(level interface{}) (interface{}, error) {
	// ignore whitespaces, only return the relevant "*"/"-" Word
	return types.NewUnorderedListItemPrefix(types.OneAsterisk, 1)
}

func (p *parser) callonUnorderedListItemPrefix34() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUnorderedListItemPrefix34(stack["level"])
}

func (c *current) onUnorderedListItemPrefix42(level interface{}) (interface{}, error) {
	// ignore whitespaces, only return the relevant "*"/"-" Word
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
	// TODO: replace with (InnerParagraph+ ContinuedBlockElement*) and use a single rule for all item contents ?
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

	return types.NewParagraph(lines.([]interface{}), append(attributes.([]interface{}), t.(types.AdmonitionKind)))

}

func (p *parser) callonParagraph2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onParagraph2(stack["attributes"], stack["t"], stack["lines"])
}

func (c *current) onParagraph23(attributes, lines interface{}) (interface{}, error) {

	return types.NewParagraph(lines.([]interface{}), attributes.([]interface{}))

}

func (p *parser) callonParagraph23() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onParagraph23(stack["attributes"], stack["lines"])
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

func (c *current) onWordWithQuotePunctuation1() (interface{}, error) {
	// can have "*", "_" or "`" within, maybe because the user inserted another quote, or made an error (extra or missing space, for example)
	return c.text, nil
}

func (p *parser) callonWordWithQuotePunctuation1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onWordWithQuotePunctuation1()
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

func (c *current) onWord1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonWord1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onWord1()
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
