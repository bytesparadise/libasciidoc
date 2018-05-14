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
			pos:  position{line: 18, col: 1, offset: 500},
			expr: &actionExpr{
				pos: position{line: 18, col: 13, offset: 512},
				run: (*parser).callonDocument1,
				expr: &seqExpr{
					pos: position{line: 18, col: 13, offset: 512},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 18, col: 13, offset: 512},
							label: "frontMatter",
							expr: &zeroOrOneExpr{
								pos: position{line: 18, col: 26, offset: 525},
								expr: &ruleRefExpr{
									pos:  position{line: 18, col: 26, offset: 525},
									name: "FrontMatter",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 18, col: 40, offset: 539},
							label: "documentHeader",
							expr: &zeroOrOneExpr{
								pos: position{line: 18, col: 56, offset: 555},
								expr: &ruleRefExpr{
									pos:  position{line: 18, col: 56, offset: 555},
									name: "DocumentHeader",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 18, col: 73, offset: 572},
							label: "blocks",
							expr: &ruleRefExpr{
								pos:  position{line: 18, col: 81, offset: 580},
								name: "DocumentBlocks",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 18, col: 97, offset: 596},
							name: "EOF",
						},
					},
				},
			},
		},
		{
			name: "DocumentBlocks",
			pos:  position{line: 22, col: 1, offset: 684},
			expr: &choiceExpr{
				pos: position{line: 22, col: 19, offset: 702},
				alternatives: []interface{}{
					&labeledExpr{
						pos:   position{line: 22, col: 19, offset: 702},
						label: "content",
						expr: &seqExpr{
							pos: position{line: 22, col: 28, offset: 711},
							exprs: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 22, col: 28, offset: 711},
									expr: &ruleRefExpr{
										pos:  position{line: 22, col: 28, offset: 711},
										name: "Preamble",
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 22, col: 38, offset: 721},
									expr: &ruleRefExpr{
										pos:  position{line: 22, col: 38, offset: 721},
										name: "Section",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 22, col: 50, offset: 733},
						run: (*parser).callonDocumentBlocks8,
						expr: &labeledExpr{
							pos:   position{line: 22, col: 50, offset: 733},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 22, col: 59, offset: 742},
								expr: &ruleRefExpr{
									pos:  position{line: 22, col: 59, offset: 742},
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
			pos:  position{line: 26, col: 1, offset: 787},
			expr: &choiceExpr{
				pos: position{line: 26, col: 17, offset: 803},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 26, col: 17, offset: 803},
						name: "DocumentAttributeDeclaration",
					},
					&ruleRefExpr{
						pos:  position{line: 26, col: 48, offset: 834},
						name: "DocumentAttributeReset",
					},
					&ruleRefExpr{
						pos:  position{line: 26, col: 73, offset: 859},
						name: "TableOfContentsMacro",
					},
					&ruleRefExpr{
						pos:  position{line: 26, col: 96, offset: 882},
						name: "BlockImage",
					},
					&ruleRefExpr{
						pos:  position{line: 26, col: 109, offset: 895},
						name: "List",
					},
					&ruleRefExpr{
						pos:  position{line: 26, col: 116, offset: 902},
						name: "LiteralBlock",
					},
					&ruleRefExpr{
						pos:  position{line: 26, col: 131, offset: 917},
						name: "DelimitedBlock",
					},
					&ruleRefExpr{
						pos:  position{line: 26, col: 148, offset: 934},
						name: "BlankLine",
					},
					&ruleRefExpr{
						pos:  position{line: 26, col: 160, offset: 946},
						name: "Paragraph",
					},
				},
			},
		},
		{
			name: "Preamble",
			pos:  position{line: 28, col: 1, offset: 1019},
			expr: &actionExpr{
				pos: position{line: 28, col: 13, offset: 1031},
				run: (*parser).callonPreamble1,
				expr: &labeledExpr{
					pos:   position{line: 28, col: 13, offset: 1031},
					label: "elements",
					expr: &oneOrMoreExpr{
						pos: position{line: 28, col: 23, offset: 1041},
						expr: &ruleRefExpr{
							pos:  position{line: 28, col: 23, offset: 1041},
							name: "BlockElement",
						},
					},
				},
			},
		},
		{
			name: "FrontMatter",
			pos:  position{line: 35, col: 1, offset: 1224},
			expr: &ruleRefExpr{
				pos:  position{line: 35, col: 16, offset: 1239},
				name: "YamlFrontMatter",
			},
		},
		{
			name: "FrontMatter",
			pos:  position{line: 37, col: 1, offset: 1257},
			expr: &actionExpr{
				pos: position{line: 37, col: 16, offset: 1272},
				run: (*parser).callonFrontMatter1,
				expr: &seqExpr{
					pos: position{line: 37, col: 16, offset: 1272},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 37, col: 16, offset: 1272},
							name: "YamlFrontMatterToken",
						},
						&labeledExpr{
							pos:   position{line: 37, col: 37, offset: 1293},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 37, col: 46, offset: 1302},
								name: "YamlFrontMatterContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 37, col: 70, offset: 1326},
							name: "YamlFrontMatterToken",
						},
					},
				},
			},
		},
		{
			name: "YamlFrontMatterToken",
			pos:  position{line: 41, col: 1, offset: 1406},
			expr: &seqExpr{
				pos: position{line: 41, col: 26, offset: 1431},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 41, col: 26, offset: 1431},
						val:        "---",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 41, col: 32, offset: 1437},
						name: "EOL",
					},
				},
			},
		},
		{
			name: "YamlFrontMatterContent",
			pos:  position{line: 43, col: 1, offset: 1442},
			expr: &actionExpr{
				pos: position{line: 43, col: 27, offset: 1468},
				run: (*parser).callonYamlFrontMatterContent1,
				expr: &zeroOrMoreExpr{
					pos: position{line: 43, col: 27, offset: 1468},
					expr: &seqExpr{
						pos: position{line: 43, col: 28, offset: 1469},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 43, col: 28, offset: 1469},
								expr: &ruleRefExpr{
									pos:  position{line: 43, col: 29, offset: 1470},
									name: "YamlFrontMatterToken",
								},
							},
							&anyMatcher{
								line: 43, col: 50, offset: 1491,
							},
						},
					},
				},
			},
		},
		{
			name: "DocumentHeader",
			pos:  position{line: 51, col: 1, offset: 1715},
			expr: &actionExpr{
				pos: position{line: 51, col: 19, offset: 1733},
				run: (*parser).callonDocumentHeader1,
				expr: &seqExpr{
					pos: position{line: 51, col: 19, offset: 1733},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 51, col: 19, offset: 1733},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 51, col: 27, offset: 1741},
								name: "DocumentTitle",
							},
						},
						&labeledExpr{
							pos:   position{line: 51, col: 42, offset: 1756},
							label: "authors",
							expr: &zeroOrOneExpr{
								pos: position{line: 51, col: 51, offset: 1765},
								expr: &ruleRefExpr{
									pos:  position{line: 51, col: 51, offset: 1765},
									name: "DocumentAuthors",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 51, col: 69, offset: 1783},
							label: "revision",
							expr: &zeroOrOneExpr{
								pos: position{line: 51, col: 79, offset: 1793},
								expr: &ruleRefExpr{
									pos:  position{line: 51, col: 79, offset: 1793},
									name: "DocumentRevision",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 51, col: 98, offset: 1812},
							label: "otherAttributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 51, col: 115, offset: 1829},
								expr: &ruleRefExpr{
									pos:  position{line: 51, col: 115, offset: 1829},
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
			pos:  position{line: 55, col: 1, offset: 1960},
			expr: &actionExpr{
				pos: position{line: 55, col: 18, offset: 1977},
				run: (*parser).callonDocumentTitle1,
				expr: &seqExpr{
					pos: position{line: 55, col: 18, offset: 1977},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 55, col: 18, offset: 1977},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 55, col: 29, offset: 1988},
								expr: &ruleRefExpr{
									pos:  position{line: 55, col: 30, offset: 1989},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 55, col: 49, offset: 2008},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 55, col: 56, offset: 2015},
								val:        "=",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 55, col: 61, offset: 2020},
							expr: &ruleRefExpr{
								pos:  position{line: 55, col: 61, offset: 2020},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 55, col: 65, offset: 2024},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 55, col: 74, offset: 2033},
								name: "InlineElements",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 55, col: 90, offset: 2049},
							expr: &ruleRefExpr{
								pos:  position{line: 55, col: 90, offset: 2049},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 55, col: 94, offset: 2053},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 55, col: 97, offset: 2056},
								expr: &ruleRefExpr{
									pos:  position{line: 55, col: 98, offset: 2057},
									name: "InlineElementID",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 55, col: 116, offset: 2075},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthors",
			pos:  position{line: 59, col: 1, offset: 2191},
			expr: &choiceExpr{
				pos: position{line: 59, col: 20, offset: 2210},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 59, col: 20, offset: 2210},
						name: "DocumentAuthorsInlineForm",
					},
					&ruleRefExpr{
						pos:  position{line: 59, col: 48, offset: 2238},
						name: "DocumentAuthorsAttributeForm",
					},
				},
			},
		},
		{
			name: "DocumentAuthorsInlineForm",
			pos:  position{line: 61, col: 1, offset: 2268},
			expr: &actionExpr{
				pos: position{line: 61, col: 30, offset: 2297},
				run: (*parser).callonDocumentAuthorsInlineForm1,
				expr: &seqExpr{
					pos: position{line: 61, col: 30, offset: 2297},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 61, col: 30, offset: 2297},
							expr: &ruleRefExpr{
								pos:  position{line: 61, col: 30, offset: 2297},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 61, col: 34, offset: 2301},
							expr: &litMatcher{
								pos:        position{line: 61, col: 35, offset: 2302},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 61, col: 39, offset: 2306},
							label: "authors",
							expr: &oneOrMoreExpr{
								pos: position{line: 61, col: 48, offset: 2315},
								expr: &ruleRefExpr{
									pos:  position{line: 61, col: 48, offset: 2315},
									name: "DocumentAuthor",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 61, col: 65, offset: 2332},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthorsAttributeForm",
			pos:  position{line: 65, col: 1, offset: 2402},
			expr: &actionExpr{
				pos: position{line: 65, col: 33, offset: 2434},
				run: (*parser).callonDocumentAuthorsAttributeForm1,
				expr: &seqExpr{
					pos: position{line: 65, col: 33, offset: 2434},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 65, col: 33, offset: 2434},
							expr: &ruleRefExpr{
								pos:  position{line: 65, col: 33, offset: 2434},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 65, col: 37, offset: 2438},
							val:        ":author:",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 65, col: 48, offset: 2449},
							label: "author",
							expr: &ruleRefExpr{
								pos:  position{line: 65, col: 56, offset: 2457},
								name: "DocumentAuthor",
							},
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthor",
			pos:  position{line: 69, col: 1, offset: 2548},
			expr: &actionExpr{
				pos: position{line: 69, col: 19, offset: 2566},
				run: (*parser).callonDocumentAuthor1,
				expr: &seqExpr{
					pos: position{line: 69, col: 19, offset: 2566},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 69, col: 19, offset: 2566},
							expr: &ruleRefExpr{
								pos:  position{line: 69, col: 19, offset: 2566},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 69, col: 23, offset: 2570},
							label: "namePart1",
							expr: &ruleRefExpr{
								pos:  position{line: 69, col: 34, offset: 2581},
								name: "DocumentAuthorNamePart",
							},
						},
						&labeledExpr{
							pos:   position{line: 69, col: 58, offset: 2605},
							label: "namePart2",
							expr: &zeroOrOneExpr{
								pos: position{line: 69, col: 68, offset: 2615},
								expr: &ruleRefExpr{
									pos:  position{line: 69, col: 69, offset: 2616},
									name: "DocumentAuthorNamePart",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 69, col: 94, offset: 2641},
							label: "namePart3",
							expr: &zeroOrOneExpr{
								pos: position{line: 69, col: 104, offset: 2651},
								expr: &ruleRefExpr{
									pos:  position{line: 69, col: 105, offset: 2652},
									name: "DocumentAuthorNamePart",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 69, col: 130, offset: 2677},
							label: "email",
							expr: &zeroOrOneExpr{
								pos: position{line: 69, col: 136, offset: 2683},
								expr: &ruleRefExpr{
									pos:  position{line: 69, col: 137, offset: 2684},
									name: "DocumentAuthorEmail",
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 69, col: 159, offset: 2706},
							expr: &ruleRefExpr{
								pos:  position{line: 69, col: 159, offset: 2706},
								name: "WS",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 69, col: 163, offset: 2710},
							expr: &litMatcher{
								pos:        position{line: 69, col: 163, offset: 2710},
								val:        ";",
								ignoreCase: false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 69, col: 168, offset: 2715},
							expr: &ruleRefExpr{
								pos:  position{line: 69, col: 168, offset: 2715},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthorNamePart",
			pos:  position{line: 74, col: 1, offset: 2880},
			expr: &seqExpr{
				pos: position{line: 74, col: 27, offset: 2906},
				exprs: []interface{}{
					&notExpr{
						pos: position{line: 74, col: 27, offset: 2906},
						expr: &litMatcher{
							pos:        position{line: 74, col: 28, offset: 2907},
							val:        "<",
							ignoreCase: false,
						},
					},
					&notExpr{
						pos: position{line: 74, col: 32, offset: 2911},
						expr: &litMatcher{
							pos:        position{line: 74, col: 33, offset: 2912},
							val:        ";",
							ignoreCase: false,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 74, col: 37, offset: 2916},
						name: "Word",
					},
					&zeroOrMoreExpr{
						pos: position{line: 74, col: 42, offset: 2921},
						expr: &ruleRefExpr{
							pos:  position{line: 74, col: 42, offset: 2921},
							name: "WS",
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthorEmail",
			pos:  position{line: 76, col: 1, offset: 2926},
			expr: &seqExpr{
				pos: position{line: 76, col: 24, offset: 2949},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 76, col: 24, offset: 2949},
						val:        "<",
						ignoreCase: false,
					},
					&labeledExpr{
						pos:   position{line: 76, col: 28, offset: 2953},
						label: "email",
						expr: &oneOrMoreExpr{
							pos: position{line: 76, col: 34, offset: 2959},
							expr: &seqExpr{
								pos: position{line: 76, col: 35, offset: 2960},
								exprs: []interface{}{
									&notExpr{
										pos: position{line: 76, col: 35, offset: 2960},
										expr: &litMatcher{
											pos:        position{line: 76, col: 36, offset: 2961},
											val:        ">",
											ignoreCase: false,
										},
									},
									&notExpr{
										pos: position{line: 76, col: 40, offset: 2965},
										expr: &ruleRefExpr{
											pos:  position{line: 76, col: 41, offset: 2966},
											name: "EOL",
										},
									},
									&anyMatcher{
										line: 76, col: 45, offset: 2970,
									},
								},
							},
						},
					},
					&litMatcher{
						pos:        position{line: 76, col: 49, offset: 2974},
						val:        ">",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "DocumentRevision",
			pos:  position{line: 80, col: 1, offset: 3110},
			expr: &actionExpr{
				pos: position{line: 80, col: 21, offset: 3130},
				run: (*parser).callonDocumentRevision1,
				expr: &seqExpr{
					pos: position{line: 80, col: 21, offset: 3130},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 80, col: 21, offset: 3130},
							expr: &ruleRefExpr{
								pos:  position{line: 80, col: 21, offset: 3130},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 80, col: 25, offset: 3134},
							expr: &litMatcher{
								pos:        position{line: 80, col: 26, offset: 3135},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 80, col: 30, offset: 3139},
							label: "revnumber",
							expr: &zeroOrOneExpr{
								pos: position{line: 80, col: 40, offset: 3149},
								expr: &ruleRefExpr{
									pos:  position{line: 80, col: 41, offset: 3150},
									name: "DocumentRevisionNumber",
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 80, col: 66, offset: 3175},
							expr: &litMatcher{
								pos:        position{line: 80, col: 66, offset: 3175},
								val:        ",",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 80, col: 71, offset: 3180},
							label: "revdate",
							expr: &zeroOrOneExpr{
								pos: position{line: 80, col: 79, offset: 3188},
								expr: &ruleRefExpr{
									pos:  position{line: 80, col: 80, offset: 3189},
									name: "DocumentRevisionDate",
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 80, col: 103, offset: 3212},
							expr: &litMatcher{
								pos:        position{line: 80, col: 103, offset: 3212},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 80, col: 108, offset: 3217},
							label: "revremark",
							expr: &zeroOrOneExpr{
								pos: position{line: 80, col: 118, offset: 3227},
								expr: &ruleRefExpr{
									pos:  position{line: 80, col: 119, offset: 3228},
									name: "DocumentRevisionRemark",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 80, col: 144, offset: 3253},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentRevisionNumber",
			pos:  position{line: 85, col: 1, offset: 3426},
			expr: &choiceExpr{
				pos: position{line: 85, col: 27, offset: 3452},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 85, col: 27, offset: 3452},
						exprs: []interface{}{
							&litMatcher{
								pos:        position{line: 85, col: 27, offset: 3452},
								val:        "v",
								ignoreCase: true,
							},
							&ruleRefExpr{
								pos:  position{line: 85, col: 32, offset: 3457},
								name: "DIGIT",
							},
							&zeroOrMoreExpr{
								pos: position{line: 85, col: 39, offset: 3464},
								expr: &seqExpr{
									pos: position{line: 85, col: 40, offset: 3465},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 85, col: 40, offset: 3465},
											expr: &ruleRefExpr{
												pos:  position{line: 85, col: 41, offset: 3466},
												name: "EOL",
											},
										},
										&notExpr{
											pos: position{line: 85, col: 45, offset: 3470},
											expr: &litMatcher{
												pos:        position{line: 85, col: 46, offset: 3471},
												val:        ",",
												ignoreCase: false,
											},
										},
										&notExpr{
											pos: position{line: 85, col: 50, offset: 3475},
											expr: &litMatcher{
												pos:        position{line: 85, col: 51, offset: 3476},
												val:        ":",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 85, col: 55, offset: 3480,
										},
									},
								},
							},
						},
					},
					&seqExpr{
						pos: position{line: 85, col: 61, offset: 3486},
						exprs: []interface{}{
							&zeroOrOneExpr{
								pos: position{line: 85, col: 61, offset: 3486},
								expr: &litMatcher{
									pos:        position{line: 85, col: 61, offset: 3486},
									val:        "v",
									ignoreCase: true,
								},
							},
							&ruleRefExpr{
								pos:  position{line: 85, col: 67, offset: 3492},
								name: "DIGIT",
							},
							&zeroOrMoreExpr{
								pos: position{line: 85, col: 74, offset: 3499},
								expr: &seqExpr{
									pos: position{line: 85, col: 75, offset: 3500},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 85, col: 75, offset: 3500},
											expr: &ruleRefExpr{
												pos:  position{line: 85, col: 76, offset: 3501},
												name: "EOL",
											},
										},
										&notExpr{
											pos: position{line: 85, col: 80, offset: 3505},
											expr: &litMatcher{
												pos:        position{line: 85, col: 81, offset: 3506},
												val:        ",",
												ignoreCase: false,
											},
										},
										&notExpr{
											pos: position{line: 85, col: 85, offset: 3510},
											expr: &litMatcher{
												pos:        position{line: 85, col: 86, offset: 3511},
												val:        ":",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 85, col: 90, offset: 3515,
										},
									},
								},
							},
							&zeroOrMoreExpr{
								pos: position{line: 85, col: 94, offset: 3519},
								expr: &ruleRefExpr{
									pos:  position{line: 85, col: 94, offset: 3519},
									name: "WS",
								},
							},
							&andExpr{
								pos: position{line: 85, col: 98, offset: 3523},
								expr: &litMatcher{
									pos:        position{line: 85, col: 99, offset: 3524},
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
			pos:  position{line: 86, col: 1, offset: 3528},
			expr: &zeroOrMoreExpr{
				pos: position{line: 86, col: 25, offset: 3552},
				expr: &seqExpr{
					pos: position{line: 86, col: 26, offset: 3553},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 86, col: 26, offset: 3553},
							expr: &ruleRefExpr{
								pos:  position{line: 86, col: 27, offset: 3554},
								name: "EOL",
							},
						},
						&notExpr{
							pos: position{line: 86, col: 31, offset: 3558},
							expr: &litMatcher{
								pos:        position{line: 86, col: 32, offset: 3559},
								val:        ":",
								ignoreCase: false,
							},
						},
						&anyMatcher{
							line: 86, col: 36, offset: 3563,
						},
					},
				},
			},
		},
		{
			name: "DocumentRevisionRemark",
			pos:  position{line: 87, col: 1, offset: 3568},
			expr: &zeroOrMoreExpr{
				pos: position{line: 87, col: 27, offset: 3594},
				expr: &seqExpr{
					pos: position{line: 87, col: 28, offset: 3595},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 87, col: 28, offset: 3595},
							expr: &ruleRefExpr{
								pos:  position{line: 87, col: 29, offset: 3596},
								name: "EOL",
							},
						},
						&anyMatcher{
							line: 87, col: 33, offset: 3600,
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclaration",
			pos:  position{line: 92, col: 1, offset: 3720},
			expr: &choiceExpr{
				pos: position{line: 92, col: 33, offset: 3752},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 92, col: 33, offset: 3752},
						name: "DocumentAttributeDeclarationWithNameOnly",
					},
					&ruleRefExpr{
						pos:  position{line: 92, col: 76, offset: 3795},
						name: "DocumentAttributeDeclarationWithNameAndValue",
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclarationWithNameOnly",
			pos:  position{line: 94, col: 1, offset: 3842},
			expr: &actionExpr{
				pos: position{line: 94, col: 45, offset: 3886},
				run: (*parser).callonDocumentAttributeDeclarationWithNameOnly1,
				expr: &seqExpr{
					pos: position{line: 94, col: 45, offset: 3886},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 94, col: 45, offset: 3886},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 94, col: 49, offset: 3890},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 94, col: 55, offset: 3896},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 94, col: 70, offset: 3911},
							val:        ":",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 94, col: 74, offset: 3915},
							expr: &ruleRefExpr{
								pos:  position{line: 94, col: 74, offset: 3915},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 94, col: 78, offset: 3919},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclarationWithNameAndValue",
			pos:  position{line: 98, col: 1, offset: 4004},
			expr: &actionExpr{
				pos: position{line: 98, col: 49, offset: 4052},
				run: (*parser).callonDocumentAttributeDeclarationWithNameAndValue1,
				expr: &seqExpr{
					pos: position{line: 98, col: 49, offset: 4052},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 98, col: 49, offset: 4052},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 98, col: 53, offset: 4056},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 98, col: 59, offset: 4062},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 98, col: 74, offset: 4077},
							val:        ":",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 98, col: 78, offset: 4081},
							expr: &ruleRefExpr{
								pos:  position{line: 98, col: 78, offset: 4081},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 98, col: 82, offset: 4085},
							label: "value",
							expr: &zeroOrMoreExpr{
								pos: position{line: 98, col: 88, offset: 4091},
								expr: &seqExpr{
									pos: position{line: 98, col: 89, offset: 4092},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 98, col: 89, offset: 4092},
											expr: &ruleRefExpr{
												pos:  position{line: 98, col: 90, offset: 4093},
												name: "NEWLINE",
											},
										},
										&anyMatcher{
											line: 98, col: 98, offset: 4101,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 98, col: 102, offset: 4105},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeReset",
			pos:  position{line: 102, col: 1, offset: 4208},
			expr: &choiceExpr{
				pos: position{line: 102, col: 27, offset: 4234},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 102, col: 27, offset: 4234},
						name: "DocumentAttributeResetWithSectionTitleBangSymbol",
					},
					&ruleRefExpr{
						pos:  position{line: 102, col: 78, offset: 4285},
						name: "DocumentAttributeResetWithTrailingBangSymbol",
					},
				},
			},
		},
		{
			name: "DocumentAttributeResetWithSectionTitleBangSymbol",
			pos:  position{line: 104, col: 1, offset: 4331},
			expr: &actionExpr{
				pos: position{line: 104, col: 53, offset: 4383},
				run: (*parser).callonDocumentAttributeResetWithSectionTitleBangSymbol1,
				expr: &seqExpr{
					pos: position{line: 104, col: 53, offset: 4383},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 104, col: 53, offset: 4383},
							val:        ":!",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 104, col: 58, offset: 4388},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 104, col: 64, offset: 4394},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 104, col: 79, offset: 4409},
							val:        ":",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 104, col: 83, offset: 4413},
							expr: &ruleRefExpr{
								pos:  position{line: 104, col: 83, offset: 4413},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 104, col: 87, offset: 4417},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeResetWithTrailingBangSymbol",
			pos:  position{line: 108, col: 1, offset: 4491},
			expr: &actionExpr{
				pos: position{line: 108, col: 49, offset: 4539},
				run: (*parser).callonDocumentAttributeResetWithTrailingBangSymbol1,
				expr: &seqExpr{
					pos: position{line: 108, col: 49, offset: 4539},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 108, col: 49, offset: 4539},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 108, col: 53, offset: 4543},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 108, col: 59, offset: 4549},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 108, col: 74, offset: 4564},
							val:        "!:",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 108, col: 79, offset: 4569},
							expr: &ruleRefExpr{
								pos:  position{line: 108, col: 79, offset: 4569},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 108, col: 83, offset: 4573},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeSubstitution",
			pos:  position{line: 112, col: 1, offset: 4647},
			expr: &actionExpr{
				pos: position{line: 112, col: 34, offset: 4680},
				run: (*parser).callonDocumentAttributeSubstitution1,
				expr: &seqExpr{
					pos: position{line: 112, col: 34, offset: 4680},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 112, col: 34, offset: 4680},
							val:        "{",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 112, col: 38, offset: 4684},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 112, col: 44, offset: 4690},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 112, col: 59, offset: 4705},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "AttributeName",
			pos:  position{line: 119, col: 1, offset: 4953},
			expr: &seqExpr{
				pos: position{line: 119, col: 18, offset: 4970},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 119, col: 19, offset: 4971},
						alternatives: []interface{}{
							&charClassMatcher{
								pos:        position{line: 119, col: 19, offset: 4971},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
							&charClassMatcher{
								pos:        position{line: 119, col: 27, offset: 4979},
								val:        "[a-z]",
								ranges:     []rune{'a', 'z'},
								ignoreCase: false,
								inverted:   false,
							},
							&charClassMatcher{
								pos:        position{line: 119, col: 35, offset: 4987},
								val:        "[0-9]",
								ranges:     []rune{'0', '9'},
								ignoreCase: false,
								inverted:   false,
							},
							&litMatcher{
								pos:        position{line: 119, col: 43, offset: 4995},
								val:        "_",
								ignoreCase: false,
							},
						},
					},
					&zeroOrMoreExpr{
						pos: position{line: 119, col: 48, offset: 5000},
						expr: &choiceExpr{
							pos: position{line: 119, col: 49, offset: 5001},
							alternatives: []interface{}{
								&charClassMatcher{
									pos:        position{line: 119, col: 49, offset: 5001},
									val:        "[A-Z]",
									ranges:     []rune{'A', 'Z'},
									ignoreCase: false,
									inverted:   false,
								},
								&charClassMatcher{
									pos:        position{line: 119, col: 57, offset: 5009},
									val:        "[a-z]",
									ranges:     []rune{'a', 'z'},
									ignoreCase: false,
									inverted:   false,
								},
								&charClassMatcher{
									pos:        position{line: 119, col: 65, offset: 5017},
									val:        "[0-9]",
									ranges:     []rune{'0', '9'},
									ignoreCase: false,
									inverted:   false,
								},
								&litMatcher{
									pos:        position{line: 119, col: 73, offset: 5025},
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
			pos:  position{line: 124, col: 1, offset: 5145},
			expr: &seqExpr{
				pos: position{line: 124, col: 25, offset: 5169},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 124, col: 25, offset: 5169},
						val:        "toc::[]",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 124, col: 35, offset: 5179},
						name: "NEWLINE",
					},
				},
			},
		},
		{
			name: "ElementAttribute",
			pos:  position{line: 129, col: 1, offset: 5302},
			expr: &actionExpr{
				pos: position{line: 129, col: 21, offset: 5322},
				run: (*parser).callonElementAttribute1,
				expr: &seqExpr{
					pos: position{line: 129, col: 21, offset: 5322},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 129, col: 21, offset: 5322},
							label: "attr",
							expr: &choiceExpr{
								pos: position{line: 129, col: 27, offset: 5328},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 129, col: 27, offset: 5328},
										name: "ElementID",
									},
									&ruleRefExpr{
										pos:  position{line: 129, col: 39, offset: 5340},
										name: "ElementTitle",
									},
									&ruleRefExpr{
										pos:  position{line: 129, col: 54, offset: 5355},
										name: "AdmonitionMarkerAttribute",
									},
									&ruleRefExpr{
										pos:  position{line: 129, col: 82, offset: 5383},
										name: "AttributeGroup",
									},
									&ruleRefExpr{
										pos:  position{line: 129, col: 99, offset: 5400},
										name: "InvalidElementAttribute",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 129, col: 124, offset: 5425},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ElementID",
			pos:  position{line: 133, col: 1, offset: 5516},
			expr: &choiceExpr{
				pos: position{line: 133, col: 14, offset: 5529},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 133, col: 14, offset: 5529},
						run: (*parser).callonElementID2,
						expr: &labeledExpr{
							pos:   position{line: 133, col: 14, offset: 5529},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 133, col: 18, offset: 5533},
								name: "InlineElementID",
							},
						},
					},
					&actionExpr{
						pos: position{line: 135, col: 5, offset: 5575},
						run: (*parser).callonElementID5,
						expr: &seqExpr{
							pos: position{line: 135, col: 5, offset: 5575},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 135, col: 5, offset: 5575},
									val:        "[#",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 135, col: 10, offset: 5580},
									label: "id",
									expr: &ruleRefExpr{
										pos:  position{line: 135, col: 14, offset: 5584},
										name: "ID",
									},
								},
								&litMatcher{
									pos:        position{line: 135, col: 18, offset: 5588},
									val:        "]",
									ignoreCase: false,
								},
								&zeroOrMoreExpr{
									pos: position{line: 135, col: 22, offset: 5592},
									expr: &ruleRefExpr{
										pos:  position{line: 135, col: 22, offset: 5592},
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
			pos:  position{line: 139, col: 1, offset: 5644},
			expr: &actionExpr{
				pos: position{line: 139, col: 20, offset: 5663},
				run: (*parser).callonInlineElementID1,
				expr: &seqExpr{
					pos: position{line: 139, col: 20, offset: 5663},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 139, col: 20, offset: 5663},
							val:        "[[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 139, col: 25, offset: 5668},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 139, col: 29, offset: 5672},
								name: "ID",
							},
						},
						&litMatcher{
							pos:        position{line: 139, col: 33, offset: 5676},
							val:        "]]",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 139, col: 38, offset: 5681},
							expr: &ruleRefExpr{
								pos:  position{line: 139, col: 38, offset: 5681},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "ElementTitle",
			pos:  position{line: 145, col: 1, offset: 5875},
			expr: &actionExpr{
				pos: position{line: 145, col: 17, offset: 5891},
				run: (*parser).callonElementTitle1,
				expr: &seqExpr{
					pos: position{line: 145, col: 17, offset: 5891},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 145, col: 17, offset: 5891},
							val:        ".",
							ignoreCase: false,
						},
						&notExpr{
							pos: position{line: 145, col: 21, offset: 5895},
							expr: &litMatcher{
								pos:        position{line: 145, col: 22, offset: 5896},
								val:        ".",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 145, col: 26, offset: 5900},
							expr: &ruleRefExpr{
								pos:  position{line: 145, col: 27, offset: 5901},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 145, col: 30, offset: 5904},
							label: "title",
							expr: &oneOrMoreExpr{
								pos: position{line: 145, col: 36, offset: 5910},
								expr: &seqExpr{
									pos: position{line: 145, col: 37, offset: 5911},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 145, col: 37, offset: 5911},
											expr: &ruleRefExpr{
												pos:  position{line: 145, col: 38, offset: 5912},
												name: "NEWLINE",
											},
										},
										&anyMatcher{
											line: 145, col: 46, offset: 5920,
										},
									},
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 145, col: 50, offset: 5924},
							expr: &ruleRefExpr{
								pos:  position{line: 145, col: 50, offset: 5924},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "AdmonitionMarkerAttribute",
			pos:  position{line: 150, col: 1, offset: 6069},
			expr: &actionExpr{
				pos: position{line: 150, col: 30, offset: 6098},
				run: (*parser).callonAdmonitionMarkerAttribute1,
				expr: &seqExpr{
					pos: position{line: 150, col: 30, offset: 6098},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 150, col: 30, offset: 6098},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 150, col: 34, offset: 6102},
							label: "k",
							expr: &ruleRefExpr{
								pos:  position{line: 150, col: 37, offset: 6105},
								name: "AdmonitionKind",
							},
						},
						&litMatcher{
							pos:        position{line: 150, col: 53, offset: 6121},
							val:        "]",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 150, col: 57, offset: 6125},
							expr: &ruleRefExpr{
								pos:  position{line: 150, col: 57, offset: 6125},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "AttributeGroup",
			pos:  position{line: 155, col: 1, offset: 6215},
			expr: &actionExpr{
				pos: position{line: 155, col: 19, offset: 6233},
				run: (*parser).callonAttributeGroup1,
				expr: &seqExpr{
					pos: position{line: 155, col: 19, offset: 6233},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 155, col: 19, offset: 6233},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 155, col: 23, offset: 6237},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 155, col: 34, offset: 6248},
								expr: &ruleRefExpr{
									pos:  position{line: 155, col: 35, offset: 6249},
									name: "GenericAttribute",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 155, col: 54, offset: 6268},
							val:        "]",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 155, col: 58, offset: 6272},
							expr: &ruleRefExpr{
								pos:  position{line: 155, col: 58, offset: 6272},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "GenericAttribute",
			pos:  position{line: 159, col: 1, offset: 6344},
			expr: &choiceExpr{
				pos: position{line: 159, col: 21, offset: 6364},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 159, col: 21, offset: 6364},
						run: (*parser).callonGenericAttribute2,
						expr: &seqExpr{
							pos: position{line: 159, col: 21, offset: 6364},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 159, col: 21, offset: 6364},
									label: "key",
									expr: &ruleRefExpr{
										pos:  position{line: 159, col: 26, offset: 6369},
										name: "AttributeKey",
									},
								},
								&litMatcher{
									pos:        position{line: 159, col: 40, offset: 6383},
									val:        "=",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 159, col: 44, offset: 6387},
									label: "value",
									expr: &ruleRefExpr{
										pos:  position{line: 159, col: 51, offset: 6394},
										name: "AttributeValue",
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 159, col: 67, offset: 6410},
									expr: &seqExpr{
										pos: position{line: 159, col: 68, offset: 6411},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 159, col: 68, offset: 6411},
												val:        ",",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 159, col: 72, offset: 6415},
												expr: &ruleRefExpr{
													pos:  position{line: 159, col: 72, offset: 6415},
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
						pos: position{line: 161, col: 5, offset: 6524},
						run: (*parser).callonGenericAttribute14,
						expr: &seqExpr{
							pos: position{line: 161, col: 5, offset: 6524},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 161, col: 5, offset: 6524},
									label: "key",
									expr: &ruleRefExpr{
										pos:  position{line: 161, col: 10, offset: 6529},
										name: "AttributeKey",
									},
								},
								&zeroOrOneExpr{
									pos: position{line: 161, col: 24, offset: 6543},
									expr: &seqExpr{
										pos: position{line: 161, col: 25, offset: 6544},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 161, col: 25, offset: 6544},
												val:        ",",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 161, col: 29, offset: 6548},
												expr: &ruleRefExpr{
													pos:  position{line: 161, col: 29, offset: 6548},
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
			pos:  position{line: 165, col: 1, offset: 6642},
			expr: &actionExpr{
				pos: position{line: 165, col: 17, offset: 6658},
				run: (*parser).callonAttributeKey1,
				expr: &seqExpr{
					pos: position{line: 165, col: 17, offset: 6658},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 165, col: 17, offset: 6658},
							label: "key",
							expr: &oneOrMoreExpr{
								pos: position{line: 165, col: 22, offset: 6663},
								expr: &seqExpr{
									pos: position{line: 165, col: 23, offset: 6664},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 165, col: 23, offset: 6664},
											expr: &ruleRefExpr{
												pos:  position{line: 165, col: 24, offset: 6665},
												name: "WS",
											},
										},
										&notExpr{
											pos: position{line: 165, col: 27, offset: 6668},
											expr: &litMatcher{
												pos:        position{line: 165, col: 28, offset: 6669},
												val:        "=",
												ignoreCase: false,
											},
										},
										&notExpr{
											pos: position{line: 165, col: 32, offset: 6673},
											expr: &litMatcher{
												pos:        position{line: 165, col: 33, offset: 6674},
												val:        ",",
												ignoreCase: false,
											},
										},
										&notExpr{
											pos: position{line: 165, col: 37, offset: 6678},
											expr: &litMatcher{
												pos:        position{line: 165, col: 38, offset: 6679},
												val:        "]",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 165, col: 42, offset: 6683,
										},
									},
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 165, col: 46, offset: 6687},
							expr: &ruleRefExpr{
								pos:  position{line: 165, col: 46, offset: 6687},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "AttributeValue",
			pos:  position{line: 170, col: 1, offset: 6769},
			expr: &actionExpr{
				pos: position{line: 170, col: 19, offset: 6787},
				run: (*parser).callonAttributeValue1,
				expr: &seqExpr{
					pos: position{line: 170, col: 19, offset: 6787},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 170, col: 19, offset: 6787},
							expr: &ruleRefExpr{
								pos:  position{line: 170, col: 19, offset: 6787},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 170, col: 23, offset: 6791},
							label: "value",
							expr: &zeroOrMoreExpr{
								pos: position{line: 170, col: 29, offset: 6797},
								expr: &seqExpr{
									pos: position{line: 170, col: 30, offset: 6798},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 170, col: 30, offset: 6798},
											expr: &ruleRefExpr{
												pos:  position{line: 170, col: 31, offset: 6799},
												name: "WS",
											},
										},
										&notExpr{
											pos: position{line: 170, col: 34, offset: 6802},
											expr: &litMatcher{
												pos:        position{line: 170, col: 35, offset: 6803},
												val:        "=",
												ignoreCase: false,
											},
										},
										&notExpr{
											pos: position{line: 170, col: 39, offset: 6807},
											expr: &litMatcher{
												pos:        position{line: 170, col: 40, offset: 6808},
												val:        "]",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 170, col: 44, offset: 6812,
										},
									},
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 170, col: 48, offset: 6816},
							expr: &ruleRefExpr{
								pos:  position{line: 170, col: 48, offset: 6816},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "InvalidElementAttribute",
			pos:  position{line: 175, col: 1, offset: 6903},
			expr: &actionExpr{
				pos: position{line: 175, col: 28, offset: 6930},
				run: (*parser).callonInvalidElementAttribute1,
				expr: &seqExpr{
					pos: position{line: 175, col: 28, offset: 6930},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 175, col: 28, offset: 6930},
							val:        "[",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 175, col: 32, offset: 6934},
							expr: &ruleRefExpr{
								pos:  position{line: 175, col: 32, offset: 6934},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 175, col: 36, offset: 6938},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 175, col: 44, offset: 6946},
								expr: &seqExpr{
									pos: position{line: 175, col: 45, offset: 6947},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 175, col: 45, offset: 6947},
											expr: &litMatcher{
												pos:        position{line: 175, col: 46, offset: 6948},
												val:        "]",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 175, col: 50, offset: 6952,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 175, col: 54, offset: 6956},
							val:        "]",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 175, col: 58, offset: 6960},
							expr: &ruleRefExpr{
								pos:  position{line: 175, col: 58, offset: 6960},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "Section",
			pos:  position{line: 182, col: 1, offset: 7125},
			expr: &choiceExpr{
				pos: position{line: 182, col: 12, offset: 7136},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 182, col: 12, offset: 7136},
						name: "Section1",
					},
					&ruleRefExpr{
						pos:  position{line: 182, col: 23, offset: 7147},
						name: "Section2",
					},
					&ruleRefExpr{
						pos:  position{line: 182, col: 34, offset: 7158},
						name: "Section3",
					},
					&ruleRefExpr{
						pos:  position{line: 182, col: 45, offset: 7169},
						name: "Section4",
					},
					&ruleRefExpr{
						pos:  position{line: 182, col: 56, offset: 7180},
						name: "Section5",
					},
				},
			},
		},
		{
			name: "Section1",
			pos:  position{line: 185, col: 1, offset: 7191},
			expr: &actionExpr{
				pos: position{line: 185, col: 13, offset: 7203},
				run: (*parser).callonSection11,
				expr: &seqExpr{
					pos: position{line: 185, col: 13, offset: 7203},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 185, col: 13, offset: 7203},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 185, col: 21, offset: 7211},
								name: "Section1Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 185, col: 36, offset: 7226},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 185, col: 46, offset: 7236},
								expr: &ruleRefExpr{
									pos:  position{line: 185, col: 46, offset: 7236},
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
			pos:  position{line: 189, col: 1, offset: 7343},
			expr: &actionExpr{
				pos: position{line: 189, col: 18, offset: 7360},
				run: (*parser).callonSection1Title1,
				expr: &seqExpr{
					pos: position{line: 189, col: 18, offset: 7360},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 189, col: 18, offset: 7360},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 189, col: 29, offset: 7371},
								expr: &ruleRefExpr{
									pos:  position{line: 189, col: 30, offset: 7372},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 189, col: 49, offset: 7391},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 189, col: 56, offset: 7398},
								val:        "==",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 189, col: 62, offset: 7404},
							expr: &ruleRefExpr{
								pos:  position{line: 189, col: 62, offset: 7404},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 189, col: 66, offset: 7408},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 189, col: 75, offset: 7417},
								name: "InlineElements",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 189, col: 91, offset: 7433},
							expr: &ruleRefExpr{
								pos:  position{line: 189, col: 91, offset: 7433},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 189, col: 95, offset: 7437},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 189, col: 98, offset: 7440},
								expr: &ruleRefExpr{
									pos:  position{line: 189, col: 99, offset: 7441},
									name: "InlineElementID",
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 189, col: 117, offset: 7459},
							expr: &ruleRefExpr{
								pos:  position{line: 189, col: 117, offset: 7459},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 189, col: 121, offset: 7463},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 189, col: 126, offset: 7468},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 189, col: 126, offset: 7468},
									expr: &ruleRefExpr{
										pos:  position{line: 189, col: 126, offset: 7468},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 189, col: 139, offset: 7481},
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
			pos:  position{line: 193, col: 1, offset: 7597},
			expr: &actionExpr{
				pos: position{line: 193, col: 18, offset: 7614},
				run: (*parser).callonSection1Block1,
				expr: &seqExpr{
					pos: position{line: 193, col: 18, offset: 7614},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 193, col: 18, offset: 7614},
							expr: &ruleRefExpr{
								pos:  position{line: 193, col: 19, offset: 7615},
								name: "Section1",
							},
						},
						&labeledExpr{
							pos:   position{line: 193, col: 28, offset: 7624},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 193, col: 37, offset: 7633},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 193, col: 37, offset: 7633},
										name: "Section2",
									},
									&ruleRefExpr{
										pos:  position{line: 193, col: 48, offset: 7644},
										name: "Section3",
									},
									&ruleRefExpr{
										pos:  position{line: 193, col: 59, offset: 7655},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 193, col: 70, offset: 7666},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 193, col: 81, offset: 7677},
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
			pos:  position{line: 197, col: 1, offset: 7720},
			expr: &actionExpr{
				pos: position{line: 197, col: 13, offset: 7732},
				run: (*parser).callonSection21,
				expr: &seqExpr{
					pos: position{line: 197, col: 13, offset: 7732},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 197, col: 13, offset: 7732},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 197, col: 21, offset: 7740},
								name: "Section2Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 197, col: 36, offset: 7755},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 197, col: 46, offset: 7765},
								expr: &ruleRefExpr{
									pos:  position{line: 197, col: 46, offset: 7765},
									name: "Section2Block",
								},
							},
						},
						&andExpr{
							pos: position{line: 197, col: 62, offset: 7781},
							expr: &zeroOrMoreExpr{
								pos: position{line: 197, col: 63, offset: 7782},
								expr: &ruleRefExpr{
									pos:  position{line: 197, col: 64, offset: 7783},
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
			pos:  position{line: 201, col: 1, offset: 7885},
			expr: &actionExpr{
				pos: position{line: 201, col: 18, offset: 7902},
				run: (*parser).callonSection2Title1,
				expr: &seqExpr{
					pos: position{line: 201, col: 18, offset: 7902},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 201, col: 18, offset: 7902},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 201, col: 29, offset: 7913},
								expr: &ruleRefExpr{
									pos:  position{line: 201, col: 30, offset: 7914},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 201, col: 49, offset: 7933},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 201, col: 56, offset: 7940},
								val:        "===",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 201, col: 63, offset: 7947},
							expr: &ruleRefExpr{
								pos:  position{line: 201, col: 63, offset: 7947},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 201, col: 67, offset: 7951},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 201, col: 76, offset: 7960},
								name: "InlineElements",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 201, col: 92, offset: 7976},
							expr: &ruleRefExpr{
								pos:  position{line: 201, col: 92, offset: 7976},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 201, col: 96, offset: 7980},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 201, col: 99, offset: 7983},
								expr: &ruleRefExpr{
									pos:  position{line: 201, col: 100, offset: 7984},
									name: "InlineElementID",
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 201, col: 118, offset: 8002},
							expr: &ruleRefExpr{
								pos:  position{line: 201, col: 118, offset: 8002},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 201, col: 122, offset: 8006},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 201, col: 127, offset: 8011},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 201, col: 127, offset: 8011},
									expr: &ruleRefExpr{
										pos:  position{line: 201, col: 127, offset: 8011},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 201, col: 140, offset: 8024},
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
			pos:  position{line: 205, col: 1, offset: 8139},
			expr: &actionExpr{
				pos: position{line: 205, col: 18, offset: 8156},
				run: (*parser).callonSection2Block1,
				expr: &seqExpr{
					pos: position{line: 205, col: 18, offset: 8156},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 205, col: 18, offset: 8156},
							expr: &ruleRefExpr{
								pos:  position{line: 205, col: 19, offset: 8157},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 205, col: 28, offset: 8166},
							expr: &ruleRefExpr{
								pos:  position{line: 205, col: 29, offset: 8167},
								name: "Section2",
							},
						},
						&labeledExpr{
							pos:   position{line: 205, col: 38, offset: 8176},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 205, col: 47, offset: 8185},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 205, col: 47, offset: 8185},
										name: "Section3",
									},
									&ruleRefExpr{
										pos:  position{line: 205, col: 58, offset: 8196},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 205, col: 69, offset: 8207},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 205, col: 80, offset: 8218},
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
			pos:  position{line: 209, col: 1, offset: 8261},
			expr: &actionExpr{
				pos: position{line: 209, col: 13, offset: 8273},
				run: (*parser).callonSection31,
				expr: &seqExpr{
					pos: position{line: 209, col: 13, offset: 8273},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 209, col: 13, offset: 8273},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 209, col: 21, offset: 8281},
								name: "Section3Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 209, col: 36, offset: 8296},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 209, col: 46, offset: 8306},
								expr: &ruleRefExpr{
									pos:  position{line: 209, col: 46, offset: 8306},
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
			pos:  position{line: 213, col: 1, offset: 8413},
			expr: &actionExpr{
				pos: position{line: 213, col: 18, offset: 8430},
				run: (*parser).callonSection3Title1,
				expr: &seqExpr{
					pos: position{line: 213, col: 18, offset: 8430},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 213, col: 18, offset: 8430},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 213, col: 29, offset: 8441},
								expr: &ruleRefExpr{
									pos:  position{line: 213, col: 30, offset: 8442},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 213, col: 49, offset: 8461},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 213, col: 56, offset: 8468},
								val:        "====",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 213, col: 64, offset: 8476},
							expr: &ruleRefExpr{
								pos:  position{line: 213, col: 64, offset: 8476},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 213, col: 68, offset: 8480},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 213, col: 77, offset: 8489},
								name: "InlineElements",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 213, col: 93, offset: 8505},
							expr: &ruleRefExpr{
								pos:  position{line: 213, col: 93, offset: 8505},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 213, col: 97, offset: 8509},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 213, col: 100, offset: 8512},
								expr: &ruleRefExpr{
									pos:  position{line: 213, col: 101, offset: 8513},
									name: "InlineElementID",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 213, col: 119, offset: 8531},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 213, col: 124, offset: 8536},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 213, col: 124, offset: 8536},
									expr: &ruleRefExpr{
										pos:  position{line: 213, col: 124, offset: 8536},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 213, col: 137, offset: 8549},
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
			pos:  position{line: 217, col: 1, offset: 8664},
			expr: &actionExpr{
				pos: position{line: 217, col: 18, offset: 8681},
				run: (*parser).callonSection3Block1,
				expr: &seqExpr{
					pos: position{line: 217, col: 18, offset: 8681},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 217, col: 18, offset: 8681},
							expr: &ruleRefExpr{
								pos:  position{line: 217, col: 19, offset: 8682},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 217, col: 28, offset: 8691},
							expr: &ruleRefExpr{
								pos:  position{line: 217, col: 29, offset: 8692},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 217, col: 38, offset: 8701},
							expr: &ruleRefExpr{
								pos:  position{line: 217, col: 39, offset: 8702},
								name: "Section3",
							},
						},
						&labeledExpr{
							pos:   position{line: 217, col: 48, offset: 8711},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 217, col: 57, offset: 8720},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 217, col: 57, offset: 8720},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 217, col: 68, offset: 8731},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 217, col: 79, offset: 8742},
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
			pos:  position{line: 221, col: 1, offset: 8785},
			expr: &actionExpr{
				pos: position{line: 221, col: 13, offset: 8797},
				run: (*parser).callonSection41,
				expr: &seqExpr{
					pos: position{line: 221, col: 13, offset: 8797},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 221, col: 13, offset: 8797},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 221, col: 21, offset: 8805},
								name: "Section4Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 221, col: 36, offset: 8820},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 221, col: 46, offset: 8830},
								expr: &ruleRefExpr{
									pos:  position{line: 221, col: 46, offset: 8830},
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
			pos:  position{line: 225, col: 1, offset: 8937},
			expr: &actionExpr{
				pos: position{line: 225, col: 18, offset: 8954},
				run: (*parser).callonSection4Title1,
				expr: &seqExpr{
					pos: position{line: 225, col: 18, offset: 8954},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 225, col: 18, offset: 8954},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 225, col: 29, offset: 8965},
								expr: &ruleRefExpr{
									pos:  position{line: 225, col: 30, offset: 8966},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 225, col: 49, offset: 8985},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 225, col: 56, offset: 8992},
								val:        "=====",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 225, col: 65, offset: 9001},
							expr: &ruleRefExpr{
								pos:  position{line: 225, col: 65, offset: 9001},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 225, col: 69, offset: 9005},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 225, col: 78, offset: 9014},
								name: "InlineElements",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 225, col: 94, offset: 9030},
							expr: &ruleRefExpr{
								pos:  position{line: 225, col: 94, offset: 9030},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 225, col: 98, offset: 9034},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 225, col: 101, offset: 9037},
								expr: &ruleRefExpr{
									pos:  position{line: 225, col: 102, offset: 9038},
									name: "InlineElementID",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 225, col: 120, offset: 9056},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 225, col: 125, offset: 9061},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 225, col: 125, offset: 9061},
									expr: &ruleRefExpr{
										pos:  position{line: 225, col: 125, offset: 9061},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 225, col: 138, offset: 9074},
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
			pos:  position{line: 229, col: 1, offset: 9189},
			expr: &actionExpr{
				pos: position{line: 229, col: 18, offset: 9206},
				run: (*parser).callonSection4Block1,
				expr: &seqExpr{
					pos: position{line: 229, col: 18, offset: 9206},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 229, col: 18, offset: 9206},
							expr: &ruleRefExpr{
								pos:  position{line: 229, col: 19, offset: 9207},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 229, col: 28, offset: 9216},
							expr: &ruleRefExpr{
								pos:  position{line: 229, col: 29, offset: 9217},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 229, col: 38, offset: 9226},
							expr: &ruleRefExpr{
								pos:  position{line: 229, col: 39, offset: 9227},
								name: "Section3",
							},
						},
						&notExpr{
							pos: position{line: 229, col: 48, offset: 9236},
							expr: &ruleRefExpr{
								pos:  position{line: 229, col: 49, offset: 9237},
								name: "Section4",
							},
						},
						&labeledExpr{
							pos:   position{line: 229, col: 58, offset: 9246},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 229, col: 67, offset: 9255},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 229, col: 67, offset: 9255},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 229, col: 78, offset: 9266},
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
			pos:  position{line: 233, col: 1, offset: 9309},
			expr: &actionExpr{
				pos: position{line: 233, col: 13, offset: 9321},
				run: (*parser).callonSection51,
				expr: &seqExpr{
					pos: position{line: 233, col: 13, offset: 9321},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 233, col: 13, offset: 9321},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 233, col: 21, offset: 9329},
								name: "Section5Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 233, col: 36, offset: 9344},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 233, col: 46, offset: 9354},
								expr: &ruleRefExpr{
									pos:  position{line: 233, col: 46, offset: 9354},
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
			pos:  position{line: 237, col: 1, offset: 9461},
			expr: &actionExpr{
				pos: position{line: 237, col: 18, offset: 9478},
				run: (*parser).callonSection5Title1,
				expr: &seqExpr{
					pos: position{line: 237, col: 18, offset: 9478},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 237, col: 18, offset: 9478},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 237, col: 29, offset: 9489},
								expr: &ruleRefExpr{
									pos:  position{line: 237, col: 30, offset: 9490},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 237, col: 49, offset: 9509},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 237, col: 56, offset: 9516},
								val:        "======",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 237, col: 66, offset: 9526},
							expr: &ruleRefExpr{
								pos:  position{line: 237, col: 66, offset: 9526},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 237, col: 70, offset: 9530},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 237, col: 79, offset: 9539},
								name: "InlineElements",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 237, col: 95, offset: 9555},
							expr: &ruleRefExpr{
								pos:  position{line: 237, col: 95, offset: 9555},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 237, col: 99, offset: 9559},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 237, col: 102, offset: 9562},
								expr: &ruleRefExpr{
									pos:  position{line: 237, col: 103, offset: 9563},
									name: "InlineElementID",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 237, col: 121, offset: 9581},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 237, col: 126, offset: 9586},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 237, col: 126, offset: 9586},
									expr: &ruleRefExpr{
										pos:  position{line: 237, col: 126, offset: 9586},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 237, col: 139, offset: 9599},
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
			pos:  position{line: 241, col: 1, offset: 9714},
			expr: &actionExpr{
				pos: position{line: 241, col: 18, offset: 9731},
				run: (*parser).callonSection5Block1,
				expr: &seqExpr{
					pos: position{line: 241, col: 18, offset: 9731},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 241, col: 18, offset: 9731},
							expr: &ruleRefExpr{
								pos:  position{line: 241, col: 19, offset: 9732},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 241, col: 28, offset: 9741},
							expr: &ruleRefExpr{
								pos:  position{line: 241, col: 29, offset: 9742},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 241, col: 38, offset: 9751},
							expr: &ruleRefExpr{
								pos:  position{line: 241, col: 39, offset: 9752},
								name: "Section3",
							},
						},
						&notExpr{
							pos: position{line: 241, col: 48, offset: 9761},
							expr: &ruleRefExpr{
								pos:  position{line: 241, col: 49, offset: 9762},
								name: "Section4",
							},
						},
						&notExpr{
							pos: position{line: 241, col: 58, offset: 9771},
							expr: &ruleRefExpr{
								pos:  position{line: 241, col: 59, offset: 9772},
								name: "Section5",
							},
						},
						&labeledExpr{
							pos:   position{line: 241, col: 68, offset: 9781},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 241, col: 77, offset: 9790},
								name: "BlockElement",
							},
						},
					},
				},
			},
		},
		{
			name: "List",
			pos:  position{line: 248, col: 1, offset: 9934},
			expr: &actionExpr{
				pos: position{line: 248, col: 9, offset: 9942},
				run: (*parser).callonList1,
				expr: &seqExpr{
					pos: position{line: 248, col: 9, offset: 9942},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 248, col: 9, offset: 9942},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 248, col: 20, offset: 9953},
								expr: &ruleRefExpr{
									pos:  position{line: 248, col: 21, offset: 9954},
									name: "ListAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 250, col: 5, offset: 10043},
							label: "elements",
							expr: &ruleRefExpr{
								pos:  position{line: 250, col: 14, offset: 10052},
								name: "ListItems",
							},
						},
					},
				},
			},
		},
		{
			name: "ListItems",
			pos:  position{line: 254, col: 1, offset: 10146},
			expr: &oneOrMoreExpr{
				pos: position{line: 254, col: 14, offset: 10159},
				expr: &choiceExpr{
					pos: position{line: 254, col: 15, offset: 10160},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 254, col: 15, offset: 10160},
							name: "OrderedListItem",
						},
						&ruleRefExpr{
							pos:  position{line: 254, col: 33, offset: 10178},
							name: "UnorderedListItem",
						},
						&ruleRefExpr{
							pos:  position{line: 254, col: 53, offset: 10198},
							name: "LabeledListItem",
						},
					},
				},
			},
		},
		{
			name: "ListAttribute",
			pos:  position{line: 256, col: 1, offset: 10217},
			expr: &actionExpr{
				pos: position{line: 256, col: 18, offset: 10234},
				run: (*parser).callonListAttribute1,
				expr: &seqExpr{
					pos: position{line: 256, col: 18, offset: 10234},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 256, col: 18, offset: 10234},
							label: "attribute",
							expr: &choiceExpr{
								pos: position{line: 256, col: 29, offset: 10245},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 256, col: 29, offset: 10245},
										name: "HorizontalLayout",
									},
									&ruleRefExpr{
										pos:  position{line: 256, col: 48, offset: 10264},
										name: "ListID",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 256, col: 56, offset: 10272},
							name: "NEWLINE",
						},
					},
				},
			},
		},
		{
			name: "ListID",
			pos:  position{line: 260, col: 1, offset: 10311},
			expr: &actionExpr{
				pos: position{line: 260, col: 11, offset: 10321},
				run: (*parser).callonListID1,
				expr: &seqExpr{
					pos: position{line: 260, col: 11, offset: 10321},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 260, col: 11, offset: 10321},
							val:        "[#",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 260, col: 16, offset: 10326},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 260, col: 20, offset: 10330},
								name: "ID",
							},
						},
						&litMatcher{
							pos:        position{line: 260, col: 24, offset: 10334},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "HorizontalLayout",
			pos:  position{line: 264, col: 1, offset: 10400},
			expr: &actionExpr{
				pos: position{line: 264, col: 21, offset: 10420},
				run: (*parser).callonHorizontalLayout1,
				expr: &litMatcher{
					pos:        position{line: 264, col: 21, offset: 10420},
					val:        "[horizontal]",
					ignoreCase: false,
				},
			},
		},
		{
			name: "InnerParagraph",
			pos:  position{line: 268, col: 1, offset: 10503},
			expr: &actionExpr{
				pos: position{line: 268, col: 20, offset: 10522},
				run: (*parser).callonInnerParagraph1,
				expr: &labeledExpr{
					pos:   position{line: 268, col: 20, offset: 10522},
					label: "lines",
					expr: &oneOrMoreExpr{
						pos: position{line: 268, col: 26, offset: 10528},
						expr: &seqExpr{
							pos: position{line: 269, col: 5, offset: 10534},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 269, col: 5, offset: 10534},
									expr: &ruleRefExpr{
										pos:  position{line: 269, col: 7, offset: 10536},
										name: "OrderedListItemPrefix",
									},
								},
								&notExpr{
									pos: position{line: 270, col: 5, offset: 10564},
									expr: &ruleRefExpr{
										pos:  position{line: 270, col: 7, offset: 10566},
										name: "UnorderedListItemPrefix",
									},
								},
								&notExpr{
									pos: position{line: 271, col: 5, offset: 10596},
									expr: &seqExpr{
										pos: position{line: 271, col: 7, offset: 10598},
										exprs: []interface{}{
											&ruleRefExpr{
												pos:  position{line: 271, col: 7, offset: 10598},
												name: "LabeledListItemTerm",
											},
											&ruleRefExpr{
												pos:  position{line: 271, col: 27, offset: 10618},
												name: "LabeledListItemSeparator",
											},
										},
									},
								},
								&notExpr{
									pos: position{line: 272, col: 5, offset: 10649},
									expr: &ruleRefExpr{
										pos:  position{line: 272, col: 7, offset: 10651},
										name: "ListItemContinuation",
									},
								},
								&notExpr{
									pos: position{line: 273, col: 5, offset: 10678},
									expr: &ruleRefExpr{
										pos:  position{line: 273, col: 7, offset: 10680},
										name: "ElementAttribute",
									},
								},
								&notExpr{
									pos: position{line: 274, col: 5, offset: 10702},
									expr: &ruleRefExpr{
										pos:  position{line: 274, col: 7, offset: 10704},
										name: "BlockDelimiter",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 275, col: 5, offset: 10724},
									name: "InlineElements",
								},
								&ruleRefExpr{
									pos:  position{line: 275, col: 20, offset: 10739},
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
			pos:  position{line: 279, col: 1, offset: 10809},
			expr: &actionExpr{
				pos: position{line: 279, col: 25, offset: 10833},
				run: (*parser).callonListItemContinuation1,
				expr: &seqExpr{
					pos: position{line: 279, col: 25, offset: 10833},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 279, col: 25, offset: 10833},
							val:        "+",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 279, col: 29, offset: 10837},
							expr: &ruleRefExpr{
								pos:  position{line: 279, col: 29, offset: 10837},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 279, col: 33, offset: 10841},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ContinuedBlockElement",
			pos:  position{line: 283, col: 1, offset: 10893},
			expr: &actionExpr{
				pos: position{line: 283, col: 26, offset: 10918},
				run: (*parser).callonContinuedBlockElement1,
				expr: &seqExpr{
					pos: position{line: 283, col: 26, offset: 10918},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 283, col: 26, offset: 10918},
							name: "ListItemContinuation",
						},
						&labeledExpr{
							pos:   position{line: 283, col: 47, offset: 10939},
							label: "element",
							expr: &ruleRefExpr{
								pos:  position{line: 283, col: 55, offset: 10947},
								name: "BlockElement",
							},
						},
					},
				},
			},
		},
		{
			name: "OrderedListItem",
			pos:  position{line: 290, col: 1, offset: 11103},
			expr: &actionExpr{
				pos: position{line: 290, col: 20, offset: 11122},
				run: (*parser).callonOrderedListItem1,
				expr: &seqExpr{
					pos: position{line: 290, col: 20, offset: 11122},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 290, col: 20, offset: 11122},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 290, col: 31, offset: 11133},
								expr: &ruleRefExpr{
									pos:  position{line: 290, col: 32, offset: 11134},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 290, col: 51, offset: 11153},
							label: "prefix",
							expr: &ruleRefExpr{
								pos:  position{line: 290, col: 59, offset: 11161},
								name: "OrderedListItemPrefix",
							},
						},
						&labeledExpr{
							pos:   position{line: 290, col: 82, offset: 11184},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 290, col: 91, offset: 11193},
								name: "OrderedListItemContent",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 290, col: 115, offset: 11217},
							expr: &ruleRefExpr{
								pos:  position{line: 290, col: 115, offset: 11217},
								name: "BlankLine",
							},
						},
					},
				},
			},
		},
		{
			name: "OrderedListItemPrefix",
			pos:  position{line: 294, col: 1, offset: 11360},
			expr: &choiceExpr{
				pos: position{line: 296, col: 1, offset: 11424},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 296, col: 1, offset: 11424},
						run: (*parser).callonOrderedListItemPrefix2,
						expr: &seqExpr{
							pos: position{line: 296, col: 1, offset: 11424},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 296, col: 1, offset: 11424},
									expr: &ruleRefExpr{
										pos:  position{line: 296, col: 1, offset: 11424},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 296, col: 5, offset: 11428},
									label: "style",
									expr: &litMatcher{
										pos:        position{line: 296, col: 12, offset: 11435},
										val:        ".",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 296, col: 17, offset: 11440},
									expr: &ruleRefExpr{
										pos:  position{line: 296, col: 17, offset: 11440},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 298, col: 5, offset: 11533},
						run: (*parser).callonOrderedListItemPrefix10,
						expr: &seqExpr{
							pos: position{line: 298, col: 5, offset: 11533},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 298, col: 5, offset: 11533},
									expr: &ruleRefExpr{
										pos:  position{line: 298, col: 5, offset: 11533},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 298, col: 9, offset: 11537},
									label: "style",
									expr: &litMatcher{
										pos:        position{line: 298, col: 16, offset: 11544},
										val:        "..",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 298, col: 22, offset: 11550},
									expr: &ruleRefExpr{
										pos:  position{line: 298, col: 22, offset: 11550},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 300, col: 5, offset: 11648},
						run: (*parser).callonOrderedListItemPrefix18,
						expr: &seqExpr{
							pos: position{line: 300, col: 5, offset: 11648},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 300, col: 5, offset: 11648},
									expr: &ruleRefExpr{
										pos:  position{line: 300, col: 5, offset: 11648},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 300, col: 9, offset: 11652},
									label: "style",
									expr: &litMatcher{
										pos:        position{line: 300, col: 16, offset: 11659},
										val:        "...",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 300, col: 23, offset: 11666},
									expr: &ruleRefExpr{
										pos:  position{line: 300, col: 23, offset: 11666},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 302, col: 5, offset: 11765},
						run: (*parser).callonOrderedListItemPrefix26,
						expr: &seqExpr{
							pos: position{line: 302, col: 5, offset: 11765},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 302, col: 5, offset: 11765},
									expr: &ruleRefExpr{
										pos:  position{line: 302, col: 5, offset: 11765},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 302, col: 9, offset: 11769},
									label: "style",
									expr: &litMatcher{
										pos:        position{line: 302, col: 16, offset: 11776},
										val:        "....",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 302, col: 24, offset: 11784},
									expr: &ruleRefExpr{
										pos:  position{line: 302, col: 24, offset: 11784},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 304, col: 5, offset: 11884},
						run: (*parser).callonOrderedListItemPrefix34,
						expr: &seqExpr{
							pos: position{line: 304, col: 5, offset: 11884},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 304, col: 5, offset: 11884},
									expr: &ruleRefExpr{
										pos:  position{line: 304, col: 5, offset: 11884},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 304, col: 9, offset: 11888},
									label: "style",
									expr: &litMatcher{
										pos:        position{line: 304, col: 16, offset: 11895},
										val:        ".....",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 304, col: 25, offset: 11904},
									expr: &ruleRefExpr{
										pos:  position{line: 304, col: 25, offset: 11904},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 307, col: 5, offset: 12027},
						run: (*parser).callonOrderedListItemPrefix42,
						expr: &seqExpr{
							pos: position{line: 307, col: 5, offset: 12027},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 307, col: 5, offset: 12027},
									expr: &ruleRefExpr{
										pos:  position{line: 307, col: 5, offset: 12027},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 307, col: 9, offset: 12031},
									label: "style",
									expr: &seqExpr{
										pos: position{line: 307, col: 16, offset: 12038},
										exprs: []interface{}{
											&oneOrMoreExpr{
												pos: position{line: 307, col: 16, offset: 12038},
												expr: &seqExpr{
													pos: position{line: 307, col: 17, offset: 12039},
													exprs: []interface{}{
														&notExpr{
															pos: position{line: 307, col: 17, offset: 12039},
															expr: &litMatcher{
																pos:        position{line: 307, col: 18, offset: 12040},
																val:        ".",
																ignoreCase: false,
															},
														},
														&notExpr{
															pos: position{line: 307, col: 22, offset: 12044},
															expr: &ruleRefExpr{
																pos:  position{line: 307, col: 23, offset: 12045},
																name: "WS",
															},
														},
														&notExpr{
															pos: position{line: 307, col: 26, offset: 12048},
															expr: &ruleRefExpr{
																pos:  position{line: 307, col: 27, offset: 12049},
																name: "NEWLINE",
															},
														},
														&charClassMatcher{
															pos:        position{line: 307, col: 35, offset: 12057},
															val:        "[0-9]",
															ranges:     []rune{'0', '9'},
															ignoreCase: false,
															inverted:   false,
														},
													},
												},
											},
											&litMatcher{
												pos:        position{line: 307, col: 43, offset: 12065},
												val:        ".",
												ignoreCase: false,
											},
										},
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 307, col: 48, offset: 12070},
									expr: &ruleRefExpr{
										pos:  position{line: 307, col: 48, offset: 12070},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 309, col: 5, offset: 12165},
						run: (*parser).callonOrderedListItemPrefix60,
						expr: &seqExpr{
							pos: position{line: 309, col: 5, offset: 12165},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 309, col: 5, offset: 12165},
									expr: &ruleRefExpr{
										pos:  position{line: 309, col: 5, offset: 12165},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 309, col: 9, offset: 12169},
									label: "style",
									expr: &seqExpr{
										pos: position{line: 309, col: 16, offset: 12176},
										exprs: []interface{}{
											&oneOrMoreExpr{
												pos: position{line: 309, col: 16, offset: 12176},
												expr: &seqExpr{
													pos: position{line: 309, col: 17, offset: 12177},
													exprs: []interface{}{
														&notExpr{
															pos: position{line: 309, col: 17, offset: 12177},
															expr: &litMatcher{
																pos:        position{line: 309, col: 18, offset: 12178},
																val:        ".",
																ignoreCase: false,
															},
														},
														&notExpr{
															pos: position{line: 309, col: 22, offset: 12182},
															expr: &ruleRefExpr{
																pos:  position{line: 309, col: 23, offset: 12183},
																name: "WS",
															},
														},
														&notExpr{
															pos: position{line: 309, col: 26, offset: 12186},
															expr: &ruleRefExpr{
																pos:  position{line: 309, col: 27, offset: 12187},
																name: "NEWLINE",
															},
														},
														&charClassMatcher{
															pos:        position{line: 309, col: 35, offset: 12195},
															val:        "[a-z]",
															ranges:     []rune{'a', 'z'},
															ignoreCase: false,
															inverted:   false,
														},
													},
												},
											},
											&litMatcher{
												pos:        position{line: 309, col: 43, offset: 12203},
												val:        ".",
												ignoreCase: false,
											},
										},
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 309, col: 48, offset: 12208},
									expr: &ruleRefExpr{
										pos:  position{line: 309, col: 48, offset: 12208},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 311, col: 5, offset: 12306},
						run: (*parser).callonOrderedListItemPrefix78,
						expr: &seqExpr{
							pos: position{line: 311, col: 5, offset: 12306},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 311, col: 5, offset: 12306},
									expr: &ruleRefExpr{
										pos:  position{line: 311, col: 5, offset: 12306},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 311, col: 9, offset: 12310},
									label: "style",
									expr: &seqExpr{
										pos: position{line: 311, col: 16, offset: 12317},
										exprs: []interface{}{
											&oneOrMoreExpr{
												pos: position{line: 311, col: 16, offset: 12317},
												expr: &seqExpr{
													pos: position{line: 311, col: 17, offset: 12318},
													exprs: []interface{}{
														&notExpr{
															pos: position{line: 311, col: 17, offset: 12318},
															expr: &litMatcher{
																pos:        position{line: 311, col: 18, offset: 12319},
																val:        ".",
																ignoreCase: false,
															},
														},
														&notExpr{
															pos: position{line: 311, col: 22, offset: 12323},
															expr: &ruleRefExpr{
																pos:  position{line: 311, col: 23, offset: 12324},
																name: "WS",
															},
														},
														&notExpr{
															pos: position{line: 311, col: 26, offset: 12327},
															expr: &ruleRefExpr{
																pos:  position{line: 311, col: 27, offset: 12328},
																name: "NEWLINE",
															},
														},
														&charClassMatcher{
															pos:        position{line: 311, col: 35, offset: 12336},
															val:        "[A-Z]",
															ranges:     []rune{'A', 'Z'},
															ignoreCase: false,
															inverted:   false,
														},
													},
												},
											},
											&litMatcher{
												pos:        position{line: 311, col: 43, offset: 12344},
												val:        ".",
												ignoreCase: false,
											},
										},
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 311, col: 48, offset: 12349},
									expr: &ruleRefExpr{
										pos:  position{line: 311, col: 48, offset: 12349},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 313, col: 5, offset: 12447},
						run: (*parser).callonOrderedListItemPrefix96,
						expr: &seqExpr{
							pos: position{line: 313, col: 5, offset: 12447},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 313, col: 5, offset: 12447},
									expr: &ruleRefExpr{
										pos:  position{line: 313, col: 5, offset: 12447},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 313, col: 9, offset: 12451},
									label: "style",
									expr: &seqExpr{
										pos: position{line: 313, col: 16, offset: 12458},
										exprs: []interface{}{
											&oneOrMoreExpr{
												pos: position{line: 313, col: 16, offset: 12458},
												expr: &seqExpr{
													pos: position{line: 313, col: 17, offset: 12459},
													exprs: []interface{}{
														&notExpr{
															pos: position{line: 313, col: 17, offset: 12459},
															expr: &litMatcher{
																pos:        position{line: 313, col: 18, offset: 12460},
																val:        ")",
																ignoreCase: false,
															},
														},
														&notExpr{
															pos: position{line: 313, col: 22, offset: 12464},
															expr: &ruleRefExpr{
																pos:  position{line: 313, col: 23, offset: 12465},
																name: "WS",
															},
														},
														&notExpr{
															pos: position{line: 313, col: 26, offset: 12468},
															expr: &ruleRefExpr{
																pos:  position{line: 313, col: 27, offset: 12469},
																name: "NEWLINE",
															},
														},
														&charClassMatcher{
															pos:        position{line: 313, col: 35, offset: 12477},
															val:        "[a-z]",
															ranges:     []rune{'a', 'z'},
															ignoreCase: false,
															inverted:   false,
														},
													},
												},
											},
											&litMatcher{
												pos:        position{line: 313, col: 43, offset: 12485},
												val:        ")",
												ignoreCase: false,
											},
										},
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 313, col: 48, offset: 12490},
									expr: &ruleRefExpr{
										pos:  position{line: 313, col: 48, offset: 12490},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 315, col: 5, offset: 12588},
						run: (*parser).callonOrderedListItemPrefix114,
						expr: &seqExpr{
							pos: position{line: 315, col: 5, offset: 12588},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 315, col: 5, offset: 12588},
									expr: &ruleRefExpr{
										pos:  position{line: 315, col: 5, offset: 12588},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 315, col: 9, offset: 12592},
									label: "style",
									expr: &seqExpr{
										pos: position{line: 315, col: 16, offset: 12599},
										exprs: []interface{}{
											&oneOrMoreExpr{
												pos: position{line: 315, col: 16, offset: 12599},
												expr: &seqExpr{
													pos: position{line: 315, col: 17, offset: 12600},
													exprs: []interface{}{
														&notExpr{
															pos: position{line: 315, col: 17, offset: 12600},
															expr: &litMatcher{
																pos:        position{line: 315, col: 18, offset: 12601},
																val:        ")",
																ignoreCase: false,
															},
														},
														&notExpr{
															pos: position{line: 315, col: 22, offset: 12605},
															expr: &ruleRefExpr{
																pos:  position{line: 315, col: 23, offset: 12606},
																name: "WS",
															},
														},
														&notExpr{
															pos: position{line: 315, col: 26, offset: 12609},
															expr: &ruleRefExpr{
																pos:  position{line: 315, col: 27, offset: 12610},
																name: "NEWLINE",
															},
														},
														&charClassMatcher{
															pos:        position{line: 315, col: 35, offset: 12618},
															val:        "[A-Z]",
															ranges:     []rune{'A', 'Z'},
															ignoreCase: false,
															inverted:   false,
														},
													},
												},
											},
											&litMatcher{
												pos:        position{line: 315, col: 43, offset: 12626},
												val:        ")",
												ignoreCase: false,
											},
										},
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 315, col: 48, offset: 12631},
									expr: &ruleRefExpr{
										pos:  position{line: 315, col: 48, offset: 12631},
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
			pos:  position{line: 319, col: 1, offset: 12729},
			expr: &actionExpr{
				pos: position{line: 319, col: 27, offset: 12755},
				run: (*parser).callonOrderedListItemContent1,
				expr: &labeledExpr{
					pos:   position{line: 319, col: 27, offset: 12755},
					label: "elements",
					expr: &seqExpr{
						pos: position{line: 319, col: 37, offset: 12765},
						exprs: []interface{}{
							&oneOrMoreExpr{
								pos: position{line: 319, col: 37, offset: 12765},
								expr: &ruleRefExpr{
									pos:  position{line: 319, col: 37, offset: 12765},
									name: "InnerParagraph",
								},
							},
							&zeroOrMoreExpr{
								pos: position{line: 319, col: 53, offset: 12781},
								expr: &ruleRefExpr{
									pos:  position{line: 319, col: 53, offset: 12781},
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
			pos:  position{line: 326, col: 1, offset: 13107},
			expr: &actionExpr{
				pos: position{line: 326, col: 22, offset: 13128},
				run: (*parser).callonUnorderedListItem1,
				expr: &seqExpr{
					pos: position{line: 326, col: 22, offset: 13128},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 326, col: 22, offset: 13128},
							label: "prefix",
							expr: &ruleRefExpr{
								pos:  position{line: 326, col: 30, offset: 13136},
								name: "UnorderedListItemPrefix",
							},
						},
						&labeledExpr{
							pos:   position{line: 326, col: 55, offset: 13161},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 326, col: 64, offset: 13170},
								name: "UnorderedListItemContent",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 326, col: 90, offset: 13196},
							expr: &ruleRefExpr{
								pos:  position{line: 326, col: 90, offset: 13196},
								name: "BlankLine",
							},
						},
					},
				},
			},
		},
		{
			name: "UnorderedListItemPrefix",
			pos:  position{line: 330, col: 1, offset: 13315},
			expr: &choiceExpr{
				pos: position{line: 330, col: 28, offset: 13342},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 330, col: 28, offset: 13342},
						run: (*parser).callonUnorderedListItemPrefix2,
						expr: &seqExpr{
							pos: position{line: 330, col: 28, offset: 13342},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 330, col: 28, offset: 13342},
									expr: &ruleRefExpr{
										pos:  position{line: 330, col: 28, offset: 13342},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 330, col: 32, offset: 13346},
									label: "level",
									expr: &litMatcher{
										pos:        position{line: 330, col: 39, offset: 13353},
										val:        "*****",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 330, col: 48, offset: 13362},
									expr: &ruleRefExpr{
										pos:  position{line: 330, col: 48, offset: 13362},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 332, col: 5, offset: 13501},
						run: (*parser).callonUnorderedListItemPrefix10,
						expr: &seqExpr{
							pos: position{line: 332, col: 5, offset: 13501},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 332, col: 5, offset: 13501},
									expr: &ruleRefExpr{
										pos:  position{line: 332, col: 5, offset: 13501},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 332, col: 9, offset: 13505},
									label: "level",
									expr: &litMatcher{
										pos:        position{line: 332, col: 16, offset: 13512},
										val:        "****",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 332, col: 24, offset: 13520},
									expr: &ruleRefExpr{
										pos:  position{line: 332, col: 24, offset: 13520},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 334, col: 5, offset: 13659},
						run: (*parser).callonUnorderedListItemPrefix18,
						expr: &seqExpr{
							pos: position{line: 334, col: 5, offset: 13659},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 334, col: 5, offset: 13659},
									expr: &ruleRefExpr{
										pos:  position{line: 334, col: 5, offset: 13659},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 334, col: 9, offset: 13663},
									label: "level",
									expr: &litMatcher{
										pos:        position{line: 334, col: 16, offset: 13670},
										val:        "***",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 334, col: 23, offset: 13677},
									expr: &ruleRefExpr{
										pos:  position{line: 334, col: 23, offset: 13677},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 336, col: 5, offset: 13817},
						run: (*parser).callonUnorderedListItemPrefix26,
						expr: &seqExpr{
							pos: position{line: 336, col: 5, offset: 13817},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 336, col: 5, offset: 13817},
									expr: &ruleRefExpr{
										pos:  position{line: 336, col: 5, offset: 13817},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 336, col: 9, offset: 13821},
									label: "level",
									expr: &litMatcher{
										pos:        position{line: 336, col: 16, offset: 13828},
										val:        "**",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 336, col: 22, offset: 13834},
									expr: &ruleRefExpr{
										pos:  position{line: 336, col: 22, offset: 13834},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 338, col: 5, offset: 13972},
						run: (*parser).callonUnorderedListItemPrefix34,
						expr: &seqExpr{
							pos: position{line: 338, col: 5, offset: 13972},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 338, col: 5, offset: 13972},
									expr: &ruleRefExpr{
										pos:  position{line: 338, col: 5, offset: 13972},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 338, col: 9, offset: 13976},
									label: "level",
									expr: &litMatcher{
										pos:        position{line: 338, col: 16, offset: 13983},
										val:        "*",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 338, col: 21, offset: 13988},
									expr: &ruleRefExpr{
										pos:  position{line: 338, col: 21, offset: 13988},
										name: "WS",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 340, col: 5, offset: 14125},
						run: (*parser).callonUnorderedListItemPrefix42,
						expr: &seqExpr{
							pos: position{line: 340, col: 5, offset: 14125},
							exprs: []interface{}{
								&zeroOrMoreExpr{
									pos: position{line: 340, col: 5, offset: 14125},
									expr: &ruleRefExpr{
										pos:  position{line: 340, col: 5, offset: 14125},
										name: "WS",
									},
								},
								&labeledExpr{
									pos:   position{line: 340, col: 9, offset: 14129},
									label: "level",
									expr: &litMatcher{
										pos:        position{line: 340, col: 16, offset: 14136},
										val:        "-",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 340, col: 21, offset: 14141},
									expr: &ruleRefExpr{
										pos:  position{line: 340, col: 21, offset: 14141},
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
			pos:  position{line: 344, col: 1, offset: 14271},
			expr: &actionExpr{
				pos: position{line: 344, col: 29, offset: 14299},
				run: (*parser).callonUnorderedListItemContent1,
				expr: &labeledExpr{
					pos:   position{line: 344, col: 29, offset: 14299},
					label: "elements",
					expr: &seqExpr{
						pos: position{line: 344, col: 39, offset: 14309},
						exprs: []interface{}{
							&oneOrMoreExpr{
								pos: position{line: 344, col: 39, offset: 14309},
								expr: &ruleRefExpr{
									pos:  position{line: 344, col: 39, offset: 14309},
									name: "InnerParagraph",
								},
							},
							&zeroOrMoreExpr{
								pos: position{line: 344, col: 55, offset: 14325},
								expr: &ruleRefExpr{
									pos:  position{line: 344, col: 55, offset: 14325},
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
			pos:  position{line: 351, col: 1, offset: 14649},
			expr: &choiceExpr{
				pos: position{line: 351, col: 20, offset: 14668},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 351, col: 20, offset: 14668},
						run: (*parser).callonLabeledListItem2,
						expr: &seqExpr{
							pos: position{line: 351, col: 20, offset: 14668},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 351, col: 20, offset: 14668},
									label: "term",
									expr: &ruleRefExpr{
										pos:  position{line: 351, col: 26, offset: 14674},
										name: "LabeledListItemTerm",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 351, col: 47, offset: 14695},
									name: "LabeledListItemSeparator",
								},
								&labeledExpr{
									pos:   position{line: 351, col: 72, offset: 14720},
									label: "description",
									expr: &ruleRefExpr{
										pos:  position{line: 351, col: 85, offset: 14733},
										name: "LabeledListItemDescription",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 353, col: 6, offset: 14855},
						run: (*parser).callonLabeledListItem9,
						expr: &seqExpr{
							pos: position{line: 353, col: 6, offset: 14855},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 353, col: 6, offset: 14855},
									label: "term",
									expr: &ruleRefExpr{
										pos:  position{line: 353, col: 12, offset: 14861},
										name: "LabeledListItemTerm",
									},
								},
								&litMatcher{
									pos:        position{line: 353, col: 33, offset: 14882},
									val:        "::",
									ignoreCase: false,
								},
								&zeroOrMoreExpr{
									pos: position{line: 353, col: 38, offset: 14887},
									expr: &ruleRefExpr{
										pos:  position{line: 353, col: 38, offset: 14887},
										name: "WS",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 353, col: 42, offset: 14891},
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
			pos:  position{line: 357, col: 1, offset: 15028},
			expr: &actionExpr{
				pos: position{line: 357, col: 24, offset: 15051},
				run: (*parser).callonLabeledListItemTerm1,
				expr: &labeledExpr{
					pos:   position{line: 357, col: 24, offset: 15051},
					label: "term",
					expr: &zeroOrMoreExpr{
						pos: position{line: 357, col: 29, offset: 15056},
						expr: &seqExpr{
							pos: position{line: 357, col: 30, offset: 15057},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 357, col: 30, offset: 15057},
									expr: &ruleRefExpr{
										pos:  position{line: 357, col: 31, offset: 15058},
										name: "NEWLINE",
									},
								},
								&notExpr{
									pos: position{line: 357, col: 39, offset: 15066},
									expr: &litMatcher{
										pos:        position{line: 357, col: 40, offset: 15067},
										val:        "::",
										ignoreCase: false,
									},
								},
								&anyMatcher{
									line: 357, col: 45, offset: 15072,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "LabeledListItemSeparator",
			pos:  position{line: 362, col: 1, offset: 15163},
			expr: &seqExpr{
				pos: position{line: 362, col: 30, offset: 15192},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 362, col: 30, offset: 15192},
						val:        "::",
						ignoreCase: false,
					},
					&oneOrMoreExpr{
						pos: position{line: 362, col: 35, offset: 15197},
						expr: &choiceExpr{
							pos: position{line: 362, col: 36, offset: 15198},
							alternatives: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 362, col: 36, offset: 15198},
									name: "WS",
								},
								&ruleRefExpr{
									pos:  position{line: 362, col: 41, offset: 15203},
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
			pos:  position{line: 364, col: 1, offset: 15214},
			expr: &actionExpr{
				pos: position{line: 364, col: 31, offset: 15244},
				run: (*parser).callonLabeledListItemDescription1,
				expr: &labeledExpr{
					pos:   position{line: 364, col: 31, offset: 15244},
					label: "elements",
					expr: &zeroOrMoreExpr{
						pos: position{line: 364, col: 40, offset: 15253},
						expr: &choiceExpr{
							pos: position{line: 364, col: 41, offset: 15254},
							alternatives: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 364, col: 41, offset: 15254},
									name: "InnerParagraph",
								},
								&ruleRefExpr{
									pos:  position{line: 364, col: 58, offset: 15271},
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
			pos:  position{line: 372, col: 1, offset: 15579},
			expr: &choiceExpr{
				pos: position{line: 372, col: 19, offset: 15597},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 372, col: 19, offset: 15597},
						run: (*parser).callonAdmonitionKind2,
						expr: &litMatcher{
							pos:        position{line: 372, col: 19, offset: 15597},
							val:        "TIP",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 374, col: 5, offset: 15635},
						run: (*parser).callonAdmonitionKind4,
						expr: &litMatcher{
							pos:        position{line: 374, col: 5, offset: 15635},
							val:        "NOTE",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 376, col: 5, offset: 15675},
						run: (*parser).callonAdmonitionKind6,
						expr: &litMatcher{
							pos:        position{line: 376, col: 5, offset: 15675},
							val:        "IMPORTANT",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 378, col: 5, offset: 15725},
						run: (*parser).callonAdmonitionKind8,
						expr: &litMatcher{
							pos:        position{line: 378, col: 5, offset: 15725},
							val:        "WARNING",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 380, col: 5, offset: 15771},
						run: (*parser).callonAdmonitionKind10,
						expr: &litMatcher{
							pos:        position{line: 380, col: 5, offset: 15771},
							val:        "CAUTION",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Paragraph",
			pos:  position{line: 389, col: 1, offset: 16074},
			expr: &choiceExpr{
				pos: position{line: 391, col: 5, offset: 16121},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 391, col: 5, offset: 16121},
						run: (*parser).callonParagraph2,
						expr: &seqExpr{
							pos: position{line: 391, col: 5, offset: 16121},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 391, col: 5, offset: 16121},
									label: "attributes",
									expr: &zeroOrMoreExpr{
										pos: position{line: 391, col: 16, offset: 16132},
										expr: &ruleRefExpr{
											pos:  position{line: 391, col: 17, offset: 16133},
											name: "ElementAttribute",
										},
									},
								},
								&notExpr{
									pos: position{line: 391, col: 36, offset: 16152},
									expr: &seqExpr{
										pos: position{line: 391, col: 38, offset: 16154},
										exprs: []interface{}{
											&oneOrMoreExpr{
												pos: position{line: 391, col: 38, offset: 16154},
												expr: &litMatcher{
													pos:        position{line: 391, col: 38, offset: 16154},
													val:        "=",
													ignoreCase: false,
												},
											},
											&oneOrMoreExpr{
												pos: position{line: 391, col: 43, offset: 16159},
												expr: &ruleRefExpr{
													pos:  position{line: 391, col: 43, offset: 16159},
													name: "WS",
												},
											},
											&notExpr{
												pos: position{line: 391, col: 47, offset: 16163},
												expr: &ruleRefExpr{
													pos:  position{line: 391, col: 48, offset: 16164},
													name: "NEWLINE",
												},
											},
										},
									},
								},
								&labeledExpr{
									pos:   position{line: 391, col: 57, offset: 16173},
									label: "t",
									expr: &ruleRefExpr{
										pos:  position{line: 391, col: 60, offset: 16176},
										name: "AdmonitionKind",
									},
								},
								&litMatcher{
									pos:        position{line: 391, col: 76, offset: 16192},
									val:        ": ",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 391, col: 81, offset: 16197},
									label: "lines",
									expr: &oneOrMoreExpr{
										pos: position{line: 391, col: 87, offset: 16203},
										expr: &seqExpr{
											pos: position{line: 391, col: 88, offset: 16204},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 391, col: 88, offset: 16204},
													name: "InlineElements",
												},
												&ruleRefExpr{
													pos:  position{line: 391, col: 103, offset: 16219},
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
						pos: position{line: 395, col: 5, offset: 16385},
						run: (*parser).callonParagraph23,
						expr: &seqExpr{
							pos: position{line: 395, col: 5, offset: 16385},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 395, col: 5, offset: 16385},
									label: "attributes",
									expr: &zeroOrMoreExpr{
										pos: position{line: 395, col: 16, offset: 16396},
										expr: &ruleRefExpr{
											pos:  position{line: 395, col: 17, offset: 16397},
											name: "ElementAttribute",
										},
									},
								},
								&notExpr{
									pos: position{line: 395, col: 36, offset: 16416},
									expr: &seqExpr{
										pos: position{line: 395, col: 38, offset: 16418},
										exprs: []interface{}{
											&oneOrMoreExpr{
												pos: position{line: 395, col: 38, offset: 16418},
												expr: &litMatcher{
													pos:        position{line: 395, col: 38, offset: 16418},
													val:        "=",
													ignoreCase: false,
												},
											},
											&oneOrMoreExpr{
												pos: position{line: 395, col: 43, offset: 16423},
												expr: &ruleRefExpr{
													pos:  position{line: 395, col: 43, offset: 16423},
													name: "WS",
												},
											},
											&notExpr{
												pos: position{line: 395, col: 47, offset: 16427},
												expr: &ruleRefExpr{
													pos:  position{line: 395, col: 48, offset: 16428},
													name: "NEWLINE",
												},
											},
										},
									},
								},
								&labeledExpr{
									pos:   position{line: 395, col: 57, offset: 16437},
									label: "lines",
									expr: &oneOrMoreExpr{
										pos: position{line: 395, col: 63, offset: 16443},
										expr: &seqExpr{
											pos: position{line: 395, col: 64, offset: 16444},
											exprs: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 395, col: 64, offset: 16444},
													name: "InlineElements",
												},
												&ruleRefExpr{
													pos:  position{line: 395, col: 79, offset: 16459},
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
			pos:  position{line: 399, col: 1, offset: 16561},
			expr: &actionExpr{
				pos: position{line: 399, col: 19, offset: 16579},
				run: (*parser).callonInlineElements1,
				expr: &seqExpr{
					pos: position{line: 399, col: 19, offset: 16579},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 399, col: 19, offset: 16579},
							expr: &ruleRefExpr{
								pos:  position{line: 399, col: 20, offset: 16580},
								name: "BlockDelimiter",
							},
						},
						&labeledExpr{
							pos:   position{line: 399, col: 35, offset: 16595},
							label: "elements",
							expr: &oneOrMoreExpr{
								pos: position{line: 399, col: 44, offset: 16604},
								expr: &seqExpr{
									pos: position{line: 399, col: 45, offset: 16605},
									exprs: []interface{}{
										&zeroOrMoreExpr{
											pos: position{line: 399, col: 45, offset: 16605},
											expr: &ruleRefExpr{
												pos:  position{line: 399, col: 45, offset: 16605},
												name: "WS",
											},
										},
										&notExpr{
											pos: position{line: 399, col: 49, offset: 16609},
											expr: &ruleRefExpr{
												pos:  position{line: 399, col: 50, offset: 16610},
												name: "InlineElementID",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 399, col: 66, offset: 16626},
											name: "InlineElement",
										},
										&zeroOrMoreExpr{
											pos: position{line: 399, col: 80, offset: 16640},
											expr: &ruleRefExpr{
												pos:  position{line: 399, col: 80, offset: 16640},
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
			pos:  position{line: 403, col: 1, offset: 16752},
			expr: &choiceExpr{
				pos: position{line: 403, col: 18, offset: 16769},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 403, col: 18, offset: 16769},
						name: "CrossReference",
					},
					&ruleRefExpr{
						pos:  position{line: 403, col: 35, offset: 16786},
						name: "Passthrough",
					},
					&ruleRefExpr{
						pos:  position{line: 403, col: 49, offset: 16800},
						name: "InlineImage",
					},
					&ruleRefExpr{
						pos:  position{line: 403, col: 63, offset: 16814},
						name: "QuotedText",
					},
					&ruleRefExpr{
						pos:  position{line: 403, col: 76, offset: 16827},
						name: "Link",
					},
					&ruleRefExpr{
						pos:  position{line: 403, col: 83, offset: 16834},
						name: "DocumentAttributeSubstitution",
					},
					&ruleRefExpr{
						pos:  position{line: 403, col: 115, offset: 16866},
						name: "Word",
					},
				},
			},
		},
		{
			name: "QuotedText",
			pos:  position{line: 408, col: 1, offset: 17111},
			expr: &choiceExpr{
				pos: position{line: 408, col: 15, offset: 17125},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 408, col: 15, offset: 17125},
						name: "BoldText",
					},
					&ruleRefExpr{
						pos:  position{line: 408, col: 26, offset: 17136},
						name: "ItalicText",
					},
					&ruleRefExpr{
						pos:  position{line: 408, col: 39, offset: 17149},
						name: "MonospaceText",
					},
					&ruleRefExpr{
						pos:  position{line: 409, col: 13, offset: 17177},
						name: "EscapedBoldText",
					},
					&ruleRefExpr{
						pos:  position{line: 409, col: 31, offset: 17195},
						name: "EscapedItalicText",
					},
					&ruleRefExpr{
						pos:  position{line: 409, col: 51, offset: 17215},
						name: "EscapedMonospaceText",
					},
				},
			},
		},
		{
			name: "BoldText",
			pos:  position{line: 411, col: 1, offset: 17237},
			expr: &choiceExpr{
				pos: position{line: 411, col: 13, offset: 17249},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 411, col: 13, offset: 17249},
						run: (*parser).callonBoldText2,
						expr: &seqExpr{
							pos: position{line: 411, col: 13, offset: 17249},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 411, col: 13, offset: 17249},
									expr: &litMatcher{
										pos:        position{line: 411, col: 14, offset: 17250},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 411, col: 19, offset: 17255},
									val:        "**",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 411, col: 24, offset: 17260},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 411, col: 33, offset: 17269},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 411, col: 52, offset: 17288},
									val:        "**",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 413, col: 5, offset: 17413},
						run: (*parser).callonBoldText10,
						expr: &seqExpr{
							pos: position{line: 413, col: 5, offset: 17413},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 413, col: 5, offset: 17413},
									expr: &litMatcher{
										pos:        position{line: 413, col: 6, offset: 17414},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 413, col: 11, offset: 17419},
									val:        "**",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 413, col: 16, offset: 17424},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 413, col: 25, offset: 17433},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 413, col: 44, offset: 17452},
									val:        "*",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 416, col: 5, offset: 17617},
						run: (*parser).callonBoldText18,
						expr: &seqExpr{
							pos: position{line: 416, col: 5, offset: 17617},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 416, col: 5, offset: 17617},
									expr: &litMatcher{
										pos:        position{line: 416, col: 6, offset: 17618},
										val:        "\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 416, col: 10, offset: 17622},
									val:        "*",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 416, col: 14, offset: 17626},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 416, col: 23, offset: 17635},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 416, col: 42, offset: 17654},
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
			pos:  position{line: 420, col: 1, offset: 17754},
			expr: &choiceExpr{
				pos: position{line: 420, col: 20, offset: 17773},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 420, col: 20, offset: 17773},
						run: (*parser).callonEscapedBoldText2,
						expr: &seqExpr{
							pos: position{line: 420, col: 20, offset: 17773},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 420, col: 20, offset: 17773},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 420, col: 33, offset: 17786},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 420, col: 33, offset: 17786},
												val:        "\\\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 420, col: 38, offset: 17791},
												expr: &litMatcher{
													pos:        position{line: 420, col: 38, offset: 17791},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 420, col: 44, offset: 17797},
									val:        "**",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 420, col: 49, offset: 17802},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 420, col: 58, offset: 17811},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 420, col: 77, offset: 17830},
									val:        "**",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 422, col: 5, offset: 17985},
						run: (*parser).callonEscapedBoldText13,
						expr: &seqExpr{
							pos: position{line: 422, col: 5, offset: 17985},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 422, col: 5, offset: 17985},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 422, col: 18, offset: 17998},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 422, col: 18, offset: 17998},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 422, col: 22, offset: 18002},
												expr: &litMatcher{
													pos:        position{line: 422, col: 22, offset: 18002},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 422, col: 28, offset: 18008},
									val:        "**",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 422, col: 33, offset: 18013},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 422, col: 42, offset: 18022},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 422, col: 61, offset: 18041},
									val:        "*",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 425, col: 5, offset: 18235},
						run: (*parser).callonEscapedBoldText24,
						expr: &seqExpr{
							pos: position{line: 425, col: 5, offset: 18235},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 425, col: 5, offset: 18235},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 425, col: 18, offset: 18248},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 425, col: 18, offset: 18248},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 425, col: 22, offset: 18252},
												expr: &litMatcher{
													pos:        position{line: 425, col: 22, offset: 18252},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 425, col: 28, offset: 18258},
									val:        "*",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 425, col: 32, offset: 18262},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 425, col: 41, offset: 18271},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 425, col: 60, offset: 18290},
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
			pos:  position{line: 429, col: 1, offset: 18442},
			expr: &choiceExpr{
				pos: position{line: 429, col: 15, offset: 18456},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 429, col: 15, offset: 18456},
						run: (*parser).callonItalicText2,
						expr: &seqExpr{
							pos: position{line: 429, col: 15, offset: 18456},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 429, col: 15, offset: 18456},
									expr: &litMatcher{
										pos:        position{line: 429, col: 16, offset: 18457},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 429, col: 21, offset: 18462},
									val:        "__",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 429, col: 26, offset: 18467},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 429, col: 35, offset: 18476},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 429, col: 54, offset: 18495},
									val:        "__",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 431, col: 5, offset: 18576},
						run: (*parser).callonItalicText10,
						expr: &seqExpr{
							pos: position{line: 431, col: 5, offset: 18576},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 431, col: 5, offset: 18576},
									expr: &litMatcher{
										pos:        position{line: 431, col: 6, offset: 18577},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 431, col: 11, offset: 18582},
									val:        "__",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 431, col: 16, offset: 18587},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 431, col: 25, offset: 18596},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 431, col: 44, offset: 18615},
									val:        "_",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 434, col: 5, offset: 18782},
						run: (*parser).callonItalicText18,
						expr: &seqExpr{
							pos: position{line: 434, col: 5, offset: 18782},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 434, col: 5, offset: 18782},
									expr: &litMatcher{
										pos:        position{line: 434, col: 6, offset: 18783},
										val:        "\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 434, col: 10, offset: 18787},
									val:        "_",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 434, col: 14, offset: 18791},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 434, col: 23, offset: 18800},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 434, col: 42, offset: 18819},
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
			pos:  position{line: 438, col: 1, offset: 18898},
			expr: &choiceExpr{
				pos: position{line: 438, col: 22, offset: 18919},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 438, col: 22, offset: 18919},
						run: (*parser).callonEscapedItalicText2,
						expr: &seqExpr{
							pos: position{line: 438, col: 22, offset: 18919},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 438, col: 22, offset: 18919},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 438, col: 35, offset: 18932},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 438, col: 35, offset: 18932},
												val:        "\\\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 438, col: 40, offset: 18937},
												expr: &litMatcher{
													pos:        position{line: 438, col: 40, offset: 18937},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 438, col: 46, offset: 18943},
									val:        "__",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 438, col: 51, offset: 18948},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 438, col: 60, offset: 18957},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 438, col: 79, offset: 18976},
									val:        "__",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 440, col: 5, offset: 19131},
						run: (*parser).callonEscapedItalicText13,
						expr: &seqExpr{
							pos: position{line: 440, col: 5, offset: 19131},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 440, col: 5, offset: 19131},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 440, col: 18, offset: 19144},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 440, col: 18, offset: 19144},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 440, col: 22, offset: 19148},
												expr: &litMatcher{
													pos:        position{line: 440, col: 22, offset: 19148},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 440, col: 28, offset: 19154},
									val:        "__",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 440, col: 33, offset: 19159},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 440, col: 42, offset: 19168},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 440, col: 61, offset: 19187},
									val:        "_",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 443, col: 5, offset: 19381},
						run: (*parser).callonEscapedItalicText24,
						expr: &seqExpr{
							pos: position{line: 443, col: 5, offset: 19381},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 443, col: 5, offset: 19381},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 443, col: 18, offset: 19394},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 443, col: 18, offset: 19394},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 443, col: 22, offset: 19398},
												expr: &litMatcher{
													pos:        position{line: 443, col: 22, offset: 19398},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 443, col: 28, offset: 19404},
									val:        "_",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 443, col: 32, offset: 19408},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 443, col: 41, offset: 19417},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 443, col: 60, offset: 19436},
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
			pos:  position{line: 447, col: 1, offset: 19588},
			expr: &choiceExpr{
				pos: position{line: 447, col: 18, offset: 19605},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 447, col: 18, offset: 19605},
						run: (*parser).callonMonospaceText2,
						expr: &seqExpr{
							pos: position{line: 447, col: 18, offset: 19605},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 447, col: 18, offset: 19605},
									expr: &litMatcher{
										pos:        position{line: 447, col: 19, offset: 19606},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 447, col: 24, offset: 19611},
									val:        "``",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 447, col: 29, offset: 19616},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 447, col: 38, offset: 19625},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 447, col: 57, offset: 19644},
									val:        "``",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 449, col: 5, offset: 19774},
						run: (*parser).callonMonospaceText10,
						expr: &seqExpr{
							pos: position{line: 449, col: 5, offset: 19774},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 449, col: 5, offset: 19774},
									expr: &litMatcher{
										pos:        position{line: 449, col: 6, offset: 19775},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 449, col: 11, offset: 19780},
									val:        "``",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 449, col: 16, offset: 19785},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 449, col: 25, offset: 19794},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 449, col: 44, offset: 19813},
									val:        "`",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 452, col: 5, offset: 19983},
						run: (*parser).callonMonospaceText18,
						expr: &seqExpr{
							pos: position{line: 452, col: 5, offset: 19983},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 452, col: 5, offset: 19983},
									expr: &litMatcher{
										pos:        position{line: 452, col: 6, offset: 19984},
										val:        "\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 452, col: 10, offset: 19988},
									val:        "`",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 452, col: 14, offset: 19992},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 452, col: 23, offset: 20001},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 452, col: 42, offset: 20020},
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
			pos:  position{line: 456, col: 1, offset: 20147},
			expr: &choiceExpr{
				pos: position{line: 456, col: 25, offset: 20171},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 456, col: 25, offset: 20171},
						run: (*parser).callonEscapedMonospaceText2,
						expr: &seqExpr{
							pos: position{line: 456, col: 25, offset: 20171},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 456, col: 25, offset: 20171},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 456, col: 38, offset: 20184},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 456, col: 38, offset: 20184},
												val:        "\\\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 456, col: 43, offset: 20189},
												expr: &litMatcher{
													pos:        position{line: 456, col: 43, offset: 20189},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 456, col: 49, offset: 20195},
									val:        "``",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 456, col: 54, offset: 20200},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 456, col: 63, offset: 20209},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 456, col: 82, offset: 20228},
									val:        "``",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 458, col: 5, offset: 20383},
						run: (*parser).callonEscapedMonospaceText13,
						expr: &seqExpr{
							pos: position{line: 458, col: 5, offset: 20383},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 458, col: 5, offset: 20383},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 458, col: 18, offset: 20396},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 458, col: 18, offset: 20396},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 458, col: 22, offset: 20400},
												expr: &litMatcher{
													pos:        position{line: 458, col: 22, offset: 20400},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 458, col: 28, offset: 20406},
									val:        "``",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 458, col: 33, offset: 20411},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 458, col: 42, offset: 20420},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 458, col: 61, offset: 20439},
									val:        "`",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 461, col: 5, offset: 20633},
						run: (*parser).callonEscapedMonospaceText24,
						expr: &seqExpr{
							pos: position{line: 461, col: 5, offset: 20633},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 461, col: 5, offset: 20633},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 461, col: 18, offset: 20646},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 461, col: 18, offset: 20646},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 461, col: 22, offset: 20650},
												expr: &litMatcher{
													pos:        position{line: 461, col: 22, offset: 20650},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 461, col: 28, offset: 20656},
									val:        "`",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 461, col: 32, offset: 20660},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 461, col: 41, offset: 20669},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 461, col: 60, offset: 20688},
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
			pos:  position{line: 465, col: 1, offset: 20840},
			expr: &seqExpr{
				pos: position{line: 465, col: 22, offset: 20861},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 465, col: 22, offset: 20861},
						name: "QuotedTextContentElement",
					},
					&zeroOrMoreExpr{
						pos: position{line: 465, col: 47, offset: 20886},
						expr: &seqExpr{
							pos: position{line: 465, col: 48, offset: 20887},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 465, col: 48, offset: 20887},
									expr: &ruleRefExpr{
										pos:  position{line: 465, col: 48, offset: 20887},
										name: "WS",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 465, col: 52, offset: 20891},
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
			pos:  position{line: 467, col: 1, offset: 20919},
			expr: &choiceExpr{
				pos: position{line: 467, col: 29, offset: 20947},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 467, col: 29, offset: 20947},
						name: "QuotedText",
					},
					&ruleRefExpr{
						pos:  position{line: 467, col: 42, offset: 20960},
						name: "QuotedTextWord",
					},
					&ruleRefExpr{
						pos:  position{line: 467, col: 59, offset: 20977},
						name: "WordWithQuotePunctuation",
					},
				},
			},
		},
		{
			name: "QuotedTextWord",
			pos:  position{line: 469, col: 1, offset: 21106},
			expr: &oneOrMoreExpr{
				pos: position{line: 469, col: 19, offset: 21124},
				expr: &seqExpr{
					pos: position{line: 469, col: 20, offset: 21125},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 469, col: 20, offset: 21125},
							expr: &ruleRefExpr{
								pos:  position{line: 469, col: 21, offset: 21126},
								name: "NEWLINE",
							},
						},
						&notExpr{
							pos: position{line: 469, col: 29, offset: 21134},
							expr: &ruleRefExpr{
								pos:  position{line: 469, col: 30, offset: 21135},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 469, col: 33, offset: 21138},
							expr: &litMatcher{
								pos:        position{line: 469, col: 34, offset: 21139},
								val:        "*",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 469, col: 38, offset: 21143},
							expr: &litMatcher{
								pos:        position{line: 469, col: 39, offset: 21144},
								val:        "_",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 469, col: 43, offset: 21148},
							expr: &litMatcher{
								pos:        position{line: 469, col: 44, offset: 21149},
								val:        "`",
								ignoreCase: false,
							},
						},
						&anyMatcher{
							line: 469, col: 48, offset: 21153,
						},
					},
				},
			},
		},
		{
			name: "WordWithQuotePunctuation",
			pos:  position{line: 471, col: 1, offset: 21196},
			expr: &actionExpr{
				pos: position{line: 471, col: 29, offset: 21224},
				run: (*parser).callonWordWithQuotePunctuation1,
				expr: &oneOrMoreExpr{
					pos: position{line: 471, col: 29, offset: 21224},
					expr: &seqExpr{
						pos: position{line: 471, col: 30, offset: 21225},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 471, col: 30, offset: 21225},
								expr: &ruleRefExpr{
									pos:  position{line: 471, col: 31, offset: 21226},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 471, col: 39, offset: 21234},
								expr: &ruleRefExpr{
									pos:  position{line: 471, col: 40, offset: 21235},
									name: "WS",
								},
							},
							&anyMatcher{
								line: 471, col: 44, offset: 21239,
							},
						},
					},
				},
			},
		},
		{
			name: "UnbalancedQuotePunctuation",
			pos:  position{line: 476, col: 1, offset: 21484},
			expr: &choiceExpr{
				pos: position{line: 476, col: 31, offset: 21514},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 476, col: 31, offset: 21514},
						val:        "*",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 476, col: 37, offset: 21520},
						val:        "_",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 476, col: 43, offset: 21526},
						val:        "`",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "Passthrough",
			pos:  position{line: 481, col: 1, offset: 21638},
			expr: &choiceExpr{
				pos: position{line: 481, col: 16, offset: 21653},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 481, col: 16, offset: 21653},
						name: "TriplePlusPassthrough",
					},
					&ruleRefExpr{
						pos:  position{line: 481, col: 40, offset: 21677},
						name: "SinglePlusPassthrough",
					},
					&ruleRefExpr{
						pos:  position{line: 481, col: 64, offset: 21701},
						name: "PassthroughMacro",
					},
				},
			},
		},
		{
			name: "SinglePlusPassthrough",
			pos:  position{line: 483, col: 1, offset: 21719},
			expr: &actionExpr{
				pos: position{line: 483, col: 26, offset: 21744},
				run: (*parser).callonSinglePlusPassthrough1,
				expr: &seqExpr{
					pos: position{line: 483, col: 26, offset: 21744},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 483, col: 26, offset: 21744},
							val:        "+",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 483, col: 30, offset: 21748},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 483, col: 38, offset: 21756},
								expr: &seqExpr{
									pos: position{line: 483, col: 39, offset: 21757},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 483, col: 39, offset: 21757},
											expr: &ruleRefExpr{
												pos:  position{line: 483, col: 40, offset: 21758},
												name: "NEWLINE",
											},
										},
										&notExpr{
											pos: position{line: 483, col: 48, offset: 21766},
											expr: &litMatcher{
												pos:        position{line: 483, col: 49, offset: 21767},
												val:        "+",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 483, col: 53, offset: 21771,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 483, col: 57, offset: 21775},
							val:        "+",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TriplePlusPassthrough",
			pos:  position{line: 487, col: 1, offset: 21870},
			expr: &actionExpr{
				pos: position{line: 487, col: 26, offset: 21895},
				run: (*parser).callonTriplePlusPassthrough1,
				expr: &seqExpr{
					pos: position{line: 487, col: 26, offset: 21895},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 487, col: 26, offset: 21895},
							val:        "+++",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 487, col: 32, offset: 21901},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 487, col: 40, offset: 21909},
								expr: &seqExpr{
									pos: position{line: 487, col: 41, offset: 21910},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 487, col: 41, offset: 21910},
											expr: &litMatcher{
												pos:        position{line: 487, col: 42, offset: 21911},
												val:        "+++",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 487, col: 48, offset: 21917,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 487, col: 52, offset: 21921},
							val:        "+++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "PassthroughMacro",
			pos:  position{line: 491, col: 1, offset: 22018},
			expr: &choiceExpr{
				pos: position{line: 491, col: 21, offset: 22038},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 491, col: 21, offset: 22038},
						run: (*parser).callonPassthroughMacro2,
						expr: &seqExpr{
							pos: position{line: 491, col: 21, offset: 22038},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 491, col: 21, offset: 22038},
									val:        "pass:[",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 491, col: 30, offset: 22047},
									label: "content",
									expr: &zeroOrMoreExpr{
										pos: position{line: 491, col: 38, offset: 22055},
										expr: &ruleRefExpr{
											pos:  position{line: 491, col: 39, offset: 22056},
											name: "PassthroughMacroCharacter",
										},
									},
								},
								&litMatcher{
									pos:        position{line: 491, col: 67, offset: 22084},
									val:        "]",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 493, col: 5, offset: 22175},
						run: (*parser).callonPassthroughMacro9,
						expr: &seqExpr{
							pos: position{line: 493, col: 5, offset: 22175},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 493, col: 5, offset: 22175},
									val:        "pass:q[",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 493, col: 15, offset: 22185},
									label: "content",
									expr: &zeroOrMoreExpr{
										pos: position{line: 493, col: 23, offset: 22193},
										expr: &choiceExpr{
											pos: position{line: 493, col: 24, offset: 22194},
											alternatives: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 493, col: 24, offset: 22194},
													name: "QuotedText",
												},
												&ruleRefExpr{
													pos:  position{line: 493, col: 37, offset: 22207},
													name: "PassthroughMacroCharacter",
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 493, col: 65, offset: 22235},
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
			pos:  position{line: 497, col: 1, offset: 22325},
			expr: &seqExpr{
				pos: position{line: 497, col: 31, offset: 22355},
				exprs: []interface{}{
					&notExpr{
						pos: position{line: 497, col: 31, offset: 22355},
						expr: &litMatcher{
							pos:        position{line: 497, col: 32, offset: 22356},
							val:        "]",
							ignoreCase: false,
						},
					},
					&anyMatcher{
						line: 497, col: 36, offset: 22360,
					},
				},
			},
		},
		{
			name: "CrossReference",
			pos:  position{line: 502, col: 1, offset: 22476},
			expr: &actionExpr{
				pos: position{line: 502, col: 19, offset: 22494},
				run: (*parser).callonCrossReference1,
				expr: &seqExpr{
					pos: position{line: 502, col: 19, offset: 22494},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 502, col: 19, offset: 22494},
							val:        "<<",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 502, col: 24, offset: 22499},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 502, col: 28, offset: 22503},
								name: "ID",
							},
						},
						&litMatcher{
							pos:        position{line: 502, col: 32, offset: 22507},
							val:        ">>",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Link",
			pos:  position{line: 509, col: 1, offset: 22666},
			expr: &choiceExpr{
				pos: position{line: 509, col: 9, offset: 22674},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 509, col: 9, offset: 22674},
						name: "RelativeLink",
					},
					&ruleRefExpr{
						pos:  position{line: 509, col: 24, offset: 22689},
						name: "ExternalLink",
					},
				},
			},
		},
		{
			name: "ExternalLink",
			pos:  position{line: 511, col: 1, offset: 22704},
			expr: &actionExpr{
				pos: position{line: 511, col: 17, offset: 22720},
				run: (*parser).callonExternalLink1,
				expr: &seqExpr{
					pos: position{line: 511, col: 17, offset: 22720},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 511, col: 17, offset: 22720},
							label: "url",
							expr: &seqExpr{
								pos: position{line: 511, col: 22, offset: 22725},
								exprs: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 511, col: 22, offset: 22725},
										name: "URL_SCHEME",
									},
									&ruleRefExpr{
										pos:  position{line: 511, col: 33, offset: 22736},
										name: "URL",
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 511, col: 38, offset: 22741},
							label: "text",
							expr: &zeroOrOneExpr{
								pos: position{line: 511, col: 43, offset: 22746},
								expr: &seqExpr{
									pos: position{line: 511, col: 44, offset: 22747},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 511, col: 44, offset: 22747},
											val:        "[",
											ignoreCase: false,
										},
										&zeroOrMoreExpr{
											pos: position{line: 511, col: 48, offset: 22751},
											expr: &ruleRefExpr{
												pos:  position{line: 511, col: 49, offset: 22752},
												name: "URL_TEXT",
											},
										},
										&litMatcher{
											pos:        position{line: 511, col: 60, offset: 22763},
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
			pos:  position{line: 518, col: 1, offset: 22924},
			expr: &actionExpr{
				pos: position{line: 518, col: 17, offset: 22940},
				run: (*parser).callonRelativeLink1,
				expr: &seqExpr{
					pos: position{line: 518, col: 17, offset: 22940},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 518, col: 17, offset: 22940},
							val:        "link:",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 518, col: 25, offset: 22948},
							label: "url",
							expr: &seqExpr{
								pos: position{line: 518, col: 30, offset: 22953},
								exprs: []interface{}{
									&zeroOrOneExpr{
										pos: position{line: 518, col: 30, offset: 22953},
										expr: &ruleRefExpr{
											pos:  position{line: 518, col: 30, offset: 22953},
											name: "URL_SCHEME",
										},
									},
									&ruleRefExpr{
										pos:  position{line: 518, col: 42, offset: 22965},
										name: "URL",
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 518, col: 47, offset: 22970},
							label: "text",
							expr: &seqExpr{
								pos: position{line: 518, col: 53, offset: 22976},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 518, col: 53, offset: 22976},
										val:        "[",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 518, col: 57, offset: 22980},
										expr: &ruleRefExpr{
											pos:  position{line: 518, col: 58, offset: 22981},
											name: "URL_TEXT",
										},
									},
									&litMatcher{
										pos:        position{line: 518, col: 69, offset: 22992},
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
			pos:  position{line: 528, col: 1, offset: 23254},
			expr: &actionExpr{
				pos: position{line: 528, col: 15, offset: 23268},
				run: (*parser).callonBlockImage1,
				expr: &seqExpr{
					pos: position{line: 528, col: 15, offset: 23268},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 528, col: 15, offset: 23268},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 528, col: 26, offset: 23279},
								expr: &ruleRefExpr{
									pos:  position{line: 528, col: 27, offset: 23280},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 528, col: 46, offset: 23299},
							label: "image",
							expr: &ruleRefExpr{
								pos:  position{line: 528, col: 52, offset: 23305},
								name: "BlockImageMacro",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 528, col: 69, offset: 23322},
							expr: &ruleRefExpr{
								pos:  position{line: 528, col: 69, offset: 23322},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 528, col: 73, offset: 23326},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "BlockImageMacro",
			pos:  position{line: 533, col: 1, offset: 23485},
			expr: &actionExpr{
				pos: position{line: 533, col: 20, offset: 23504},
				run: (*parser).callonBlockImageMacro1,
				expr: &seqExpr{
					pos: position{line: 533, col: 20, offset: 23504},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 533, col: 20, offset: 23504},
							val:        "image::",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 533, col: 30, offset: 23514},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 533, col: 36, offset: 23520},
								name: "URL",
							},
						},
						&litMatcher{
							pos:        position{line: 533, col: 41, offset: 23525},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 533, col: 45, offset: 23529},
							label: "attributes",
							expr: &zeroOrOneExpr{
								pos: position{line: 533, col: 57, offset: 23541},
								expr: &ruleRefExpr{
									pos:  position{line: 533, col: 57, offset: 23541},
									name: "URL_TEXT",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 533, col: 68, offset: 23552},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "InlineImage",
			pos:  position{line: 537, col: 1, offset: 23619},
			expr: &actionExpr{
				pos: position{line: 537, col: 16, offset: 23634},
				run: (*parser).callonInlineImage1,
				expr: &labeledExpr{
					pos:   position{line: 537, col: 16, offset: 23634},
					label: "image",
					expr: &ruleRefExpr{
						pos:  position{line: 537, col: 22, offset: 23640},
						name: "InlineImageMacro",
					},
				},
			},
		},
		{
			name: "InlineImageMacro",
			pos:  position{line: 542, col: 1, offset: 23785},
			expr: &actionExpr{
				pos: position{line: 542, col: 21, offset: 23805},
				run: (*parser).callonInlineImageMacro1,
				expr: &seqExpr{
					pos: position{line: 542, col: 21, offset: 23805},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 542, col: 21, offset: 23805},
							val:        "image:",
							ignoreCase: false,
						},
						&notExpr{
							pos: position{line: 542, col: 30, offset: 23814},
							expr: &litMatcher{
								pos:        position{line: 542, col: 31, offset: 23815},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 542, col: 35, offset: 23819},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 542, col: 41, offset: 23825},
								name: "URL",
							},
						},
						&litMatcher{
							pos:        position{line: 542, col: 46, offset: 23830},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 542, col: 50, offset: 23834},
							label: "attributes",
							expr: &zeroOrOneExpr{
								pos: position{line: 542, col: 62, offset: 23846},
								expr: &ruleRefExpr{
									pos:  position{line: 542, col: 62, offset: 23846},
									name: "URL_TEXT",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 542, col: 73, offset: 23857},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "DelimitedBlock",
			pos:  position{line: 549, col: 1, offset: 24187},
			expr: &choiceExpr{
				pos: position{line: 549, col: 19, offset: 24205},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 549, col: 19, offset: 24205},
						name: "FencedBlock",
					},
					&ruleRefExpr{
						pos:  position{line: 549, col: 33, offset: 24219},
						name: "ListingBlock",
					},
					&ruleRefExpr{
						pos:  position{line: 549, col: 48, offset: 24234},
						name: "ExampleBlock",
					},
				},
			},
		},
		{
			name: "BlockDelimiter",
			pos:  position{line: 551, col: 1, offset: 24248},
			expr: &choiceExpr{
				pos: position{line: 551, col: 19, offset: 24266},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 551, col: 19, offset: 24266},
						name: "LiteralBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 551, col: 43, offset: 24290},
						name: "FencedBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 551, col: 66, offset: 24313},
						name: "ListingBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 551, col: 90, offset: 24337},
						name: "ExampleBlockDelimiter",
					},
				},
			},
		},
		{
			name: "FencedBlockDelimiter",
			pos:  position{line: 553, col: 1, offset: 24360},
			expr: &litMatcher{
				pos:        position{line: 553, col: 25, offset: 24384},
				val:        "```",
				ignoreCase: false,
			},
		},
		{
			name: "FencedBlock",
			pos:  position{line: 555, col: 1, offset: 24391},
			expr: &actionExpr{
				pos: position{line: 555, col: 16, offset: 24406},
				run: (*parser).callonFencedBlock1,
				expr: &seqExpr{
					pos: position{line: 555, col: 16, offset: 24406},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 555, col: 16, offset: 24406},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 555, col: 27, offset: 24417},
								expr: &ruleRefExpr{
									pos:  position{line: 555, col: 28, offset: 24418},
									name: "ElementAttribute",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 555, col: 47, offset: 24437},
							name: "FencedBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 555, col: 68, offset: 24458},
							expr: &ruleRefExpr{
								pos:  position{line: 555, col: 68, offset: 24458},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 555, col: 72, offset: 24462},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 555, col: 80, offset: 24470},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 555, col: 88, offset: 24478},
								expr: &choiceExpr{
									pos: position{line: 555, col: 89, offset: 24479},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 555, col: 89, offset: 24479},
											name: "List",
										},
										&ruleRefExpr{
											pos:  position{line: 555, col: 96, offset: 24486},
											name: "InnerParagraph",
										},
										&ruleRefExpr{
											pos:  position{line: 555, col: 113, offset: 24503},
											name: "BlankLine",
										},
									},
								},
							},
						},
						&choiceExpr{
							pos: position{line: 555, col: 126, offset: 24516},
							alternatives: []interface{}{
								&seqExpr{
									pos: position{line: 555, col: 127, offset: 24517},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 555, col: 127, offset: 24517},
											name: "FencedBlockDelimiter",
										},
										&zeroOrMoreExpr{
											pos: position{line: 555, col: 148, offset: 24538},
											expr: &ruleRefExpr{
												pos:  position{line: 555, col: 148, offset: 24538},
												name: "WS",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 555, col: 152, offset: 24542},
											name: "EOL",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 555, col: 159, offset: 24549},
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
			pos:  position{line: 559, col: 1, offset: 24666},
			expr: &litMatcher{
				pos:        position{line: 559, col: 26, offset: 24691},
				val:        "----",
				ignoreCase: false,
			},
		},
		{
			name: "ListingBlock",
			pos:  position{line: 561, col: 1, offset: 24699},
			expr: &actionExpr{
				pos: position{line: 561, col: 17, offset: 24715},
				run: (*parser).callonListingBlock1,
				expr: &seqExpr{
					pos: position{line: 561, col: 17, offset: 24715},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 561, col: 17, offset: 24715},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 561, col: 28, offset: 24726},
								expr: &ruleRefExpr{
									pos:  position{line: 561, col: 29, offset: 24727},
									name: "ElementAttribute",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 561, col: 48, offset: 24746},
							name: "ListingBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 561, col: 70, offset: 24768},
							expr: &ruleRefExpr{
								pos:  position{line: 561, col: 70, offset: 24768},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 561, col: 74, offset: 24772},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 561, col: 82, offset: 24780},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 561, col: 90, offset: 24788},
								expr: &choiceExpr{
									pos: position{line: 561, col: 91, offset: 24789},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 561, col: 91, offset: 24789},
											name: "List",
										},
										&ruleRefExpr{
											pos:  position{line: 561, col: 98, offset: 24796},
											name: "InnerParagraph",
										},
										&ruleRefExpr{
											pos:  position{line: 561, col: 115, offset: 24813},
											name: "BlankLine",
										},
									},
								},
							},
						},
						&choiceExpr{
							pos: position{line: 561, col: 128, offset: 24826},
							alternatives: []interface{}{
								&seqExpr{
									pos: position{line: 561, col: 129, offset: 24827},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 561, col: 129, offset: 24827},
											name: "ListingBlockDelimiter",
										},
										&zeroOrMoreExpr{
											pos: position{line: 561, col: 151, offset: 24849},
											expr: &ruleRefExpr{
												pos:  position{line: 561, col: 151, offset: 24849},
												name: "WS",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 561, col: 155, offset: 24853},
											name: "EOL",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 561, col: 162, offset: 24860},
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
			pos:  position{line: 565, col: 1, offset: 24978},
			expr: &litMatcher{
				pos:        position{line: 565, col: 26, offset: 25003},
				val:        "====",
				ignoreCase: false,
			},
		},
		{
			name: "ExampleBlock",
			pos:  position{line: 567, col: 1, offset: 25011},
			expr: &actionExpr{
				pos: position{line: 567, col: 17, offset: 25027},
				run: (*parser).callonExampleBlock1,
				expr: &seqExpr{
					pos: position{line: 567, col: 17, offset: 25027},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 567, col: 17, offset: 25027},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 567, col: 28, offset: 25038},
								expr: &ruleRefExpr{
									pos:  position{line: 567, col: 29, offset: 25039},
									name: "ElementAttribute",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 567, col: 48, offset: 25058},
							name: "ExampleBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 567, col: 70, offset: 25080},
							expr: &ruleRefExpr{
								pos:  position{line: 567, col: 70, offset: 25080},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 567, col: 74, offset: 25084},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 567, col: 82, offset: 25092},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 567, col: 90, offset: 25100},
								expr: &choiceExpr{
									pos: position{line: 567, col: 91, offset: 25101},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 567, col: 91, offset: 25101},
											name: "List",
										},
										&ruleRefExpr{
											pos:  position{line: 567, col: 98, offset: 25108},
											name: "InnerParagraph",
										},
										&ruleRefExpr{
											pos:  position{line: 567, col: 115, offset: 25125},
											name: "BlankLine",
										},
									},
								},
							},
						},
						&choiceExpr{
							pos: position{line: 567, col: 129, offset: 25139},
							alternatives: []interface{}{
								&seqExpr{
									pos: position{line: 567, col: 130, offset: 25140},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 567, col: 130, offset: 25140},
											name: "ExampleBlockDelimiter",
										},
										&zeroOrMoreExpr{
											pos: position{line: 567, col: 152, offset: 25162},
											expr: &ruleRefExpr{
												pos:  position{line: 567, col: 152, offset: 25162},
												name: "WS",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 567, col: 156, offset: 25166},
											name: "EOL",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 567, col: 163, offset: 25173},
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
			pos:  position{line: 574, col: 1, offset: 25558},
			expr: &choiceExpr{
				pos: position{line: 574, col: 17, offset: 25574},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 574, col: 17, offset: 25574},
						name: "ParagraphWithSpaces",
					},
					&ruleRefExpr{
						pos:  position{line: 574, col: 39, offset: 25596},
						name: "ParagraphWithLiteralBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 574, col: 76, offset: 25633},
						name: "ParagraphWithLiteralAttribute",
					},
				},
			},
		},
		{
			name: "ParagraphWithSpaces",
			pos:  position{line: 577, col: 1, offset: 25728},
			expr: &actionExpr{
				pos: position{line: 577, col: 24, offset: 25751},
				run: (*parser).callonParagraphWithSpaces1,
				expr: &seqExpr{
					pos: position{line: 577, col: 24, offset: 25751},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 577, col: 24, offset: 25751},
							label: "spaces",
							expr: &oneOrMoreExpr{
								pos: position{line: 577, col: 32, offset: 25759},
								expr: &ruleRefExpr{
									pos:  position{line: 577, col: 32, offset: 25759},
									name: "WS",
								},
							},
						},
						&notExpr{
							pos: position{line: 577, col: 37, offset: 25764},
							expr: &ruleRefExpr{
								pos:  position{line: 577, col: 38, offset: 25765},
								name: "NEWLINE",
							},
						},
						&labeledExpr{
							pos:   position{line: 577, col: 46, offset: 25773},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 577, col: 55, offset: 25782},
								name: "LiteralBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 577, col: 76, offset: 25803},
							name: "EndOfLiteralBlock",
						},
					},
				},
			},
		},
		{
			name: "LiteralBlockContent",
			pos:  position{line: 582, col: 1, offset: 25984},
			expr: &actionExpr{
				pos: position{line: 582, col: 24, offset: 26007},
				run: (*parser).callonLiteralBlockContent1,
				expr: &labeledExpr{
					pos:   position{line: 582, col: 24, offset: 26007},
					label: "content",
					expr: &oneOrMoreExpr{
						pos: position{line: 582, col: 32, offset: 26015},
						expr: &seqExpr{
							pos: position{line: 582, col: 33, offset: 26016},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 582, col: 33, offset: 26016},
									expr: &seqExpr{
										pos: position{line: 582, col: 35, offset: 26018},
										exprs: []interface{}{
											&ruleRefExpr{
												pos:  position{line: 582, col: 35, offset: 26018},
												name: "NEWLINE",
											},
											&ruleRefExpr{
												pos:  position{line: 582, col: 43, offset: 26026},
												name: "BlankLine",
											},
										},
									},
								},
								&anyMatcher{
									line: 582, col: 54, offset: 26037,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "EndOfLiteralBlock",
			pos:  position{line: 587, col: 1, offset: 26122},
			expr: &choiceExpr{
				pos: position{line: 587, col: 22, offset: 26143},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 587, col: 22, offset: 26143},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 587, col: 22, offset: 26143},
								name: "NEWLINE",
							},
							&ruleRefExpr{
								pos:  position{line: 587, col: 30, offset: 26151},
								name: "BlankLine",
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 587, col: 42, offset: 26163},
						name: "NEWLINE",
					},
					&ruleRefExpr{
						pos:  position{line: 587, col: 52, offset: 26173},
						name: "EOF",
					},
				},
			},
		},
		{
			name: "ParagraphWithLiteralBlockDelimiter",
			pos:  position{line: 590, col: 1, offset: 26233},
			expr: &actionExpr{
				pos: position{line: 590, col: 39, offset: 26271},
				run: (*parser).callonParagraphWithLiteralBlockDelimiter1,
				expr: &seqExpr{
					pos: position{line: 590, col: 39, offset: 26271},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 590, col: 39, offset: 26271},
							name: "LiteralBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 590, col: 61, offset: 26293},
							expr: &ruleRefExpr{
								pos:  position{line: 590, col: 61, offset: 26293},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 590, col: 65, offset: 26297},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 590, col: 73, offset: 26305},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 590, col: 81, offset: 26313},
								expr: &seqExpr{
									pos: position{line: 590, col: 82, offset: 26314},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 590, col: 82, offset: 26314},
											expr: &ruleRefExpr{
												pos:  position{line: 590, col: 83, offset: 26315},
												name: "LiteralBlockDelimiter",
											},
										},
										&anyMatcher{
											line: 590, col: 105, offset: 26337,
										},
									},
								},
							},
						},
						&choiceExpr{
							pos: position{line: 590, col: 110, offset: 26342},
							alternatives: []interface{}{
								&seqExpr{
									pos: position{line: 590, col: 111, offset: 26343},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 590, col: 111, offset: 26343},
											name: "LiteralBlockDelimiter",
										},
										&zeroOrMoreExpr{
											pos: position{line: 590, col: 133, offset: 26365},
											expr: &ruleRefExpr{
												pos:  position{line: 590, col: 133, offset: 26365},
												name: "WS",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 590, col: 137, offset: 26369},
											name: "EOL",
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 590, col: 144, offset: 26376},
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
			pos:  position{line: 594, col: 1, offset: 26461},
			expr: &litMatcher{
				pos:        position{line: 594, col: 26, offset: 26486},
				val:        "....",
				ignoreCase: false,
			},
		},
		{
			name: "ParagraphWithLiteralAttribute",
			pos:  position{line: 597, col: 1, offset: 26548},
			expr: &actionExpr{
				pos: position{line: 597, col: 34, offset: 26581},
				run: (*parser).callonParagraphWithLiteralAttribute1,
				expr: &seqExpr{
					pos: position{line: 597, col: 34, offset: 26581},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 597, col: 34, offset: 26581},
							val:        "[literal]",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 597, col: 46, offset: 26593},
							expr: &ruleRefExpr{
								pos:  position{line: 597, col: 46, offset: 26593},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 597, col: 50, offset: 26597},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 597, col: 58, offset: 26605},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 597, col: 67, offset: 26614},
								name: "LiteralBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 597, col: 88, offset: 26635},
							name: "EndOfLiteralBlock",
						},
					},
				},
			},
		},
		{
			name: "BlankLine",
			pos:  position{line: 604, col: 1, offset: 26838},
			expr: &actionExpr{
				pos: position{line: 604, col: 14, offset: 26851},
				run: (*parser).callonBlankLine1,
				expr: &seqExpr{
					pos: position{line: 604, col: 14, offset: 26851},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 604, col: 14, offset: 26851},
							expr: &ruleRefExpr{
								pos:  position{line: 604, col: 15, offset: 26852},
								name: "EOF",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 604, col: 19, offset: 26856},
							expr: &ruleRefExpr{
								pos:  position{line: 604, col: 19, offset: 26856},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 604, col: 23, offset: 26860},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "Word",
			pos:  position{line: 611, col: 1, offset: 27007},
			expr: &actionExpr{
				pos: position{line: 611, col: 9, offset: 27015},
				run: (*parser).callonWord1,
				expr: &oneOrMoreExpr{
					pos: position{line: 611, col: 9, offset: 27015},
					expr: &seqExpr{
						pos: position{line: 611, col: 10, offset: 27016},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 611, col: 10, offset: 27016},
								expr: &ruleRefExpr{
									pos:  position{line: 611, col: 11, offset: 27017},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 611, col: 19, offset: 27025},
								expr: &ruleRefExpr{
									pos:  position{line: 611, col: 20, offset: 27026},
									name: "WS",
								},
							},
							&anyMatcher{
								line: 611, col: 23, offset: 27029,
							},
						},
					},
				},
			},
		},
		{
			name: "URL",
			pos:  position{line: 615, col: 1, offset: 27069},
			expr: &actionExpr{
				pos: position{line: 615, col: 8, offset: 27076},
				run: (*parser).callonURL1,
				expr: &oneOrMoreExpr{
					pos: position{line: 615, col: 8, offset: 27076},
					expr: &seqExpr{
						pos: position{line: 615, col: 9, offset: 27077},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 615, col: 9, offset: 27077},
								expr: &ruleRefExpr{
									pos:  position{line: 615, col: 10, offset: 27078},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 615, col: 18, offset: 27086},
								expr: &ruleRefExpr{
									pos:  position{line: 615, col: 19, offset: 27087},
									name: "WS",
								},
							},
							&notExpr{
								pos: position{line: 615, col: 22, offset: 27090},
								expr: &litMatcher{
									pos:        position{line: 615, col: 23, offset: 27091},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 615, col: 27, offset: 27095},
								expr: &litMatcher{
									pos:        position{line: 615, col: 28, offset: 27096},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 615, col: 32, offset: 27100,
							},
						},
					},
				},
			},
		},
		{
			name: "ID",
			pos:  position{line: 619, col: 1, offset: 27140},
			expr: &actionExpr{
				pos: position{line: 619, col: 7, offset: 27146},
				run: (*parser).callonID1,
				expr: &oneOrMoreExpr{
					pos: position{line: 619, col: 7, offset: 27146},
					expr: &seqExpr{
						pos: position{line: 619, col: 8, offset: 27147},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 619, col: 8, offset: 27147},
								expr: &ruleRefExpr{
									pos:  position{line: 619, col: 9, offset: 27148},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 619, col: 17, offset: 27156},
								expr: &ruleRefExpr{
									pos:  position{line: 619, col: 18, offset: 27157},
									name: "WS",
								},
							},
							&notExpr{
								pos: position{line: 619, col: 21, offset: 27160},
								expr: &litMatcher{
									pos:        position{line: 619, col: 22, offset: 27161},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 619, col: 26, offset: 27165},
								expr: &litMatcher{
									pos:        position{line: 619, col: 27, offset: 27166},
									val:        "]",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 619, col: 31, offset: 27170},
								expr: &litMatcher{
									pos:        position{line: 619, col: 32, offset: 27171},
									val:        "<<",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 619, col: 37, offset: 27176},
								expr: &litMatcher{
									pos:        position{line: 619, col: 38, offset: 27177},
									val:        ">>",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 619, col: 42, offset: 27181,
							},
						},
					},
				},
			},
		},
		{
			name: "URL_TEXT",
			pos:  position{line: 623, col: 1, offset: 27221},
			expr: &actionExpr{
				pos: position{line: 623, col: 13, offset: 27233},
				run: (*parser).callonURL_TEXT1,
				expr: &oneOrMoreExpr{
					pos: position{line: 623, col: 13, offset: 27233},
					expr: &seqExpr{
						pos: position{line: 623, col: 14, offset: 27234},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 623, col: 14, offset: 27234},
								expr: &ruleRefExpr{
									pos:  position{line: 623, col: 15, offset: 27235},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 623, col: 23, offset: 27243},
								expr: &litMatcher{
									pos:        position{line: 623, col: 24, offset: 27244},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 623, col: 28, offset: 27248},
								expr: &litMatcher{
									pos:        position{line: 623, col: 29, offset: 27249},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 623, col: 33, offset: 27253,
							},
						},
					},
				},
			},
		},
		{
			name: "URL_SCHEME",
			pos:  position{line: 627, col: 1, offset: 27293},
			expr: &choiceExpr{
				pos: position{line: 627, col: 15, offset: 27307},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 627, col: 15, offset: 27307},
						val:        "http://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 627, col: 27, offset: 27319},
						val:        "https://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 627, col: 40, offset: 27332},
						val:        "ftp://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 627, col: 51, offset: 27343},
						val:        "irc://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 627, col: 62, offset: 27354},
						val:        "mailto:",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "DIGIT",
			pos:  position{line: 629, col: 1, offset: 27365},
			expr: &charClassMatcher{
				pos:        position{line: 629, col: 10, offset: 27374},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NEWLINE",
			pos:  position{line: 631, col: 1, offset: 27381},
			expr: &choiceExpr{
				pos: position{line: 631, col: 12, offset: 27392},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 631, col: 12, offset: 27392},
						val:        "\r\n",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 631, col: 21, offset: 27401},
						val:        "\r",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 631, col: 28, offset: 27408},
						val:        "\n",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "WS",
			pos:  position{line: 633, col: 1, offset: 27414},
			expr: &choiceExpr{
				pos: position{line: 633, col: 7, offset: 27420},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 633, col: 7, offset: 27420},
						val:        " ",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 633, col: 13, offset: 27426},
						run: (*parser).callonWS3,
						expr: &litMatcher{
							pos:        position{line: 633, col: 13, offset: 27426},
							val:        "\t",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 637, col: 1, offset: 27471},
			expr: &notExpr{
				pos: position{line: 637, col: 8, offset: 27478},
				expr: &anyMatcher{
					line: 637, col: 9, offset: 27479,
				},
			},
		},
		{
			name: "EOL",
			pos:  position{line: 639, col: 1, offset: 27482},
			expr: &choiceExpr{
				pos: position{line: 639, col: 8, offset: 27489},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 639, col: 8, offset: 27489},
						name: "NEWLINE",
					},
					&ruleRefExpr{
						pos:  position{line: 639, col: 18, offset: 27499},
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
