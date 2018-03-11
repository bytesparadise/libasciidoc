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
						name: "Admonition",
					},
					&ruleRefExpr{
						pos:  position{line: 27, col: 161, offset: 946},
						name: "Paragraph",
					},
					&seqExpr{
						pos: position{line: 27, col: 174, offset: 959},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 27, col: 174, offset: 959},
								name: "ElementAttribute",
							},
							&ruleRefExpr{
								pos:  position{line: 27, col: 191, offset: 976},
								name: "EOL",
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 27, col: 198, offset: 983},
						name: "BlankLine",
					},
				},
			},
		},
		{
			name: "Preamble",
			pos:  position{line: 29, col: 1, offset: 1038},
			expr: &actionExpr{
				pos: position{line: 29, col: 13, offset: 1050},
				run: (*parser).callonPreamble1,
				expr: &labeledExpr{
					pos:   position{line: 29, col: 13, offset: 1050},
					label: "elements",
					expr: &zeroOrMoreExpr{
						pos: position{line: 29, col: 23, offset: 1060},
						expr: &ruleRefExpr{
							pos:  position{line: 29, col: 23, offset: 1060},
							name: "BlockElement",
						},
					},
				},
			},
		},
		{
			name: "FrontMatter",
			pos:  position{line: 36, col: 1, offset: 1243},
			expr: &ruleRefExpr{
				pos:  position{line: 36, col: 16, offset: 1258},
				name: "YamlFrontMatter",
			},
		},
		{
			name: "FrontMatter",
			pos:  position{line: 38, col: 1, offset: 1276},
			expr: &actionExpr{
				pos: position{line: 38, col: 16, offset: 1291},
				run: (*parser).callonFrontMatter1,
				expr: &seqExpr{
					pos: position{line: 38, col: 16, offset: 1291},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 38, col: 16, offset: 1291},
							name: "YamlFrontMatterToken",
						},
						&labeledExpr{
							pos:   position{line: 38, col: 37, offset: 1312},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 38, col: 46, offset: 1321},
								name: "YamlFrontMatterContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 38, col: 70, offset: 1345},
							name: "YamlFrontMatterToken",
						},
					},
				},
			},
		},
		{
			name: "YamlFrontMatterToken",
			pos:  position{line: 42, col: 1, offset: 1425},
			expr: &seqExpr{
				pos: position{line: 42, col: 26, offset: 1450},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 42, col: 26, offset: 1450},
						val:        "---",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 42, col: 32, offset: 1456},
						name: "EOL",
					},
				},
			},
		},
		{
			name: "YamlFrontMatterContent",
			pos:  position{line: 44, col: 1, offset: 1461},
			expr: &actionExpr{
				pos: position{line: 44, col: 27, offset: 1487},
				run: (*parser).callonYamlFrontMatterContent1,
				expr: &zeroOrMoreExpr{
					pos: position{line: 44, col: 27, offset: 1487},
					expr: &seqExpr{
						pos: position{line: 44, col: 28, offset: 1488},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 44, col: 28, offset: 1488},
								expr: &ruleRefExpr{
									pos:  position{line: 44, col: 29, offset: 1489},
									name: "YamlFrontMatterToken",
								},
							},
							&anyMatcher{
								line: 44, col: 50, offset: 1510,
							},
						},
					},
				},
			},
		},
		{
			name: "DocumentHeader",
			pos:  position{line: 52, col: 1, offset: 1734},
			expr: &actionExpr{
				pos: position{line: 52, col: 19, offset: 1752},
				run: (*parser).callonDocumentHeader1,
				expr: &seqExpr{
					pos: position{line: 52, col: 19, offset: 1752},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 52, col: 19, offset: 1752},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 52, col: 27, offset: 1760},
								name: "DocumentTitle",
							},
						},
						&labeledExpr{
							pos:   position{line: 52, col: 42, offset: 1775},
							label: "authors",
							expr: &zeroOrOneExpr{
								pos: position{line: 52, col: 51, offset: 1784},
								expr: &ruleRefExpr{
									pos:  position{line: 52, col: 51, offset: 1784},
									name: "DocumentAuthors",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 52, col: 69, offset: 1802},
							label: "revision",
							expr: &zeroOrOneExpr{
								pos: position{line: 52, col: 79, offset: 1812},
								expr: &ruleRefExpr{
									pos:  position{line: 52, col: 79, offset: 1812},
									name: "DocumentRevision",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 52, col: 98, offset: 1831},
							label: "otherAttributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 52, col: 115, offset: 1848},
								expr: &ruleRefExpr{
									pos:  position{line: 52, col: 115, offset: 1848},
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
			pos:  position{line: 56, col: 1, offset: 1979},
			expr: &actionExpr{
				pos: position{line: 56, col: 18, offset: 1996},
				run: (*parser).callonDocumentTitle1,
				expr: &seqExpr{
					pos: position{line: 56, col: 18, offset: 1996},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 56, col: 18, offset: 1996},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 56, col: 29, offset: 2007},
								expr: &ruleRefExpr{
									pos:  position{line: 56, col: 30, offset: 2008},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 56, col: 49, offset: 2027},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 56, col: 56, offset: 2034},
								val:        "=",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 56, col: 61, offset: 2039},
							expr: &ruleRefExpr{
								pos:  position{line: 56, col: 61, offset: 2039},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 56, col: 65, offset: 2043},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 56, col: 74, offset: 2052},
								name: "InlineContent",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 56, col: 89, offset: 2067},
							expr: &ruleRefExpr{
								pos:  position{line: 56, col: 89, offset: 2067},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 56, col: 93, offset: 2071},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 56, col: 96, offset: 2074},
								expr: &ruleRefExpr{
									pos:  position{line: 56, col: 97, offset: 2075},
									name: "InlineElementID",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 56, col: 115, offset: 2093},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthors",
			pos:  position{line: 60, col: 1, offset: 2209},
			expr: &choiceExpr{
				pos: position{line: 60, col: 20, offset: 2228},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 60, col: 20, offset: 2228},
						name: "DocumentAuthorsInlineForm",
					},
					&ruleRefExpr{
						pos:  position{line: 60, col: 48, offset: 2256},
						name: "DocumentAuthorsAttributeForm",
					},
				},
			},
		},
		{
			name: "DocumentAuthorsInlineForm",
			pos:  position{line: 62, col: 1, offset: 2286},
			expr: &actionExpr{
				pos: position{line: 62, col: 30, offset: 2315},
				run: (*parser).callonDocumentAuthorsInlineForm1,
				expr: &seqExpr{
					pos: position{line: 62, col: 30, offset: 2315},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 62, col: 30, offset: 2315},
							expr: &ruleRefExpr{
								pos:  position{line: 62, col: 30, offset: 2315},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 62, col: 34, offset: 2319},
							expr: &litMatcher{
								pos:        position{line: 62, col: 35, offset: 2320},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 62, col: 39, offset: 2324},
							label: "authors",
							expr: &oneOrMoreExpr{
								pos: position{line: 62, col: 48, offset: 2333},
								expr: &ruleRefExpr{
									pos:  position{line: 62, col: 48, offset: 2333},
									name: "DocumentAuthor",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 62, col: 65, offset: 2350},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthorsAttributeForm",
			pos:  position{line: 66, col: 1, offset: 2420},
			expr: &actionExpr{
				pos: position{line: 66, col: 33, offset: 2452},
				run: (*parser).callonDocumentAuthorsAttributeForm1,
				expr: &seqExpr{
					pos: position{line: 66, col: 33, offset: 2452},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 66, col: 33, offset: 2452},
							expr: &ruleRefExpr{
								pos:  position{line: 66, col: 33, offset: 2452},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 66, col: 37, offset: 2456},
							val:        ":author:",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 66, col: 48, offset: 2467},
							label: "author",
							expr: &ruleRefExpr{
								pos:  position{line: 66, col: 56, offset: 2475},
								name: "DocumentAuthor",
							},
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthor",
			pos:  position{line: 70, col: 1, offset: 2568},
			expr: &actionExpr{
				pos: position{line: 70, col: 19, offset: 2586},
				run: (*parser).callonDocumentAuthor1,
				expr: &seqExpr{
					pos: position{line: 70, col: 19, offset: 2586},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 70, col: 19, offset: 2586},
							expr: &ruleRefExpr{
								pos:  position{line: 70, col: 19, offset: 2586},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 70, col: 23, offset: 2590},
							label: "namePart1",
							expr: &ruleRefExpr{
								pos:  position{line: 70, col: 34, offset: 2601},
								name: "DocumentAuthorNamePart",
							},
						},
						&labeledExpr{
							pos:   position{line: 70, col: 58, offset: 2625},
							label: "namePart2",
							expr: &zeroOrOneExpr{
								pos: position{line: 70, col: 68, offset: 2635},
								expr: &ruleRefExpr{
									pos:  position{line: 70, col: 69, offset: 2636},
									name: "DocumentAuthorNamePart",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 70, col: 94, offset: 2661},
							label: "namePart3",
							expr: &zeroOrOneExpr{
								pos: position{line: 70, col: 104, offset: 2671},
								expr: &ruleRefExpr{
									pos:  position{line: 70, col: 105, offset: 2672},
									name: "DocumentAuthorNamePart",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 70, col: 130, offset: 2697},
							label: "email",
							expr: &zeroOrOneExpr{
								pos: position{line: 70, col: 136, offset: 2703},
								expr: &ruleRefExpr{
									pos:  position{line: 70, col: 137, offset: 2704},
									name: "DocumentAuthorEmail",
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 70, col: 159, offset: 2726},
							expr: &ruleRefExpr{
								pos:  position{line: 70, col: 159, offset: 2726},
								name: "WS",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 70, col: 163, offset: 2730},
							expr: &litMatcher{
								pos:        position{line: 70, col: 163, offset: 2730},
								val:        ";",
								ignoreCase: false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 70, col: 168, offset: 2735},
							expr: &ruleRefExpr{
								pos:  position{line: 70, col: 168, offset: 2735},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthorNamePart",
			pos:  position{line: 75, col: 1, offset: 2900},
			expr: &seqExpr{
				pos: position{line: 75, col: 27, offset: 2926},
				exprs: []interface{}{
					&notExpr{
						pos: position{line: 75, col: 27, offset: 2926},
						expr: &litMatcher{
							pos:        position{line: 75, col: 28, offset: 2927},
							val:        "<",
							ignoreCase: false,
						},
					},
					&notExpr{
						pos: position{line: 75, col: 32, offset: 2931},
						expr: &litMatcher{
							pos:        position{line: 75, col: 33, offset: 2932},
							val:        ";",
							ignoreCase: false,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 37, offset: 2936},
						name: "Characters",
					},
					&zeroOrMoreExpr{
						pos: position{line: 75, col: 48, offset: 2947},
						expr: &ruleRefExpr{
							pos:  position{line: 75, col: 48, offset: 2947},
							name: "WS",
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthorEmail",
			pos:  position{line: 77, col: 1, offset: 2952},
			expr: &seqExpr{
				pos: position{line: 77, col: 24, offset: 2975},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 77, col: 24, offset: 2975},
						val:        "<",
						ignoreCase: false,
					},
					&labeledExpr{
						pos:   position{line: 77, col: 28, offset: 2979},
						label: "email",
						expr: &oneOrMoreExpr{
							pos: position{line: 77, col: 34, offset: 2985},
							expr: &seqExpr{
								pos: position{line: 77, col: 35, offset: 2986},
								exprs: []interface{}{
									&notExpr{
										pos: position{line: 77, col: 35, offset: 2986},
										expr: &litMatcher{
											pos:        position{line: 77, col: 36, offset: 2987},
											val:        ">",
											ignoreCase: false,
										},
									},
									&notExpr{
										pos: position{line: 77, col: 40, offset: 2991},
										expr: &ruleRefExpr{
											pos:  position{line: 77, col: 41, offset: 2992},
											name: "EOL",
										},
									},
									&anyMatcher{
										line: 77, col: 45, offset: 2996,
									},
								},
							},
						},
					},
					&litMatcher{
						pos:        position{line: 77, col: 49, offset: 3000},
						val:        ">",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "DocumentRevision",
			pos:  position{line: 81, col: 1, offset: 3136},
			expr: &actionExpr{
				pos: position{line: 81, col: 21, offset: 3156},
				run: (*parser).callonDocumentRevision1,
				expr: &seqExpr{
					pos: position{line: 81, col: 21, offset: 3156},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 81, col: 21, offset: 3156},
							expr: &ruleRefExpr{
								pos:  position{line: 81, col: 21, offset: 3156},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 81, col: 25, offset: 3160},
							expr: &litMatcher{
								pos:        position{line: 81, col: 26, offset: 3161},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 81, col: 30, offset: 3165},
							label: "revnumber",
							expr: &zeroOrOneExpr{
								pos: position{line: 81, col: 40, offset: 3175},
								expr: &ruleRefExpr{
									pos:  position{line: 81, col: 41, offset: 3176},
									name: "DocumentRevisionNumber",
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 81, col: 66, offset: 3201},
							expr: &litMatcher{
								pos:        position{line: 81, col: 66, offset: 3201},
								val:        ",",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 81, col: 71, offset: 3206},
							label: "revdate",
							expr: &zeroOrOneExpr{
								pos: position{line: 81, col: 79, offset: 3214},
								expr: &ruleRefExpr{
									pos:  position{line: 81, col: 80, offset: 3215},
									name: "DocumentRevisionDate",
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 81, col: 103, offset: 3238},
							expr: &litMatcher{
								pos:        position{line: 81, col: 103, offset: 3238},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 81, col: 108, offset: 3243},
							label: "revremark",
							expr: &zeroOrOneExpr{
								pos: position{line: 81, col: 118, offset: 3253},
								expr: &ruleRefExpr{
									pos:  position{line: 81, col: 119, offset: 3254},
									name: "DocumentRevisionRemark",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 81, col: 144, offset: 3279},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentRevisionNumber",
			pos:  position{line: 86, col: 1, offset: 3452},
			expr: &choiceExpr{
				pos: position{line: 86, col: 27, offset: 3478},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 86, col: 27, offset: 3478},
						exprs: []interface{}{
							&litMatcher{
								pos:        position{line: 86, col: 27, offset: 3478},
								val:        "v",
								ignoreCase: true,
							},
							&ruleRefExpr{
								pos:  position{line: 86, col: 32, offset: 3483},
								name: "DIGIT",
							},
							&zeroOrMoreExpr{
								pos: position{line: 86, col: 39, offset: 3490},
								expr: &seqExpr{
									pos: position{line: 86, col: 40, offset: 3491},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 86, col: 40, offset: 3491},
											expr: &ruleRefExpr{
												pos:  position{line: 86, col: 41, offset: 3492},
												name: "EOL",
											},
										},
										&notExpr{
											pos: position{line: 86, col: 45, offset: 3496},
											expr: &litMatcher{
												pos:        position{line: 86, col: 46, offset: 3497},
												val:        ",",
												ignoreCase: false,
											},
										},
										&notExpr{
											pos: position{line: 86, col: 50, offset: 3501},
											expr: &litMatcher{
												pos:        position{line: 86, col: 51, offset: 3502},
												val:        ":",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 86, col: 55, offset: 3506,
										},
									},
								},
							},
						},
					},
					&seqExpr{
						pos: position{line: 86, col: 61, offset: 3512},
						exprs: []interface{}{
							&zeroOrOneExpr{
								pos: position{line: 86, col: 61, offset: 3512},
								expr: &litMatcher{
									pos:        position{line: 86, col: 61, offset: 3512},
									val:        "v",
									ignoreCase: true,
								},
							},
							&ruleRefExpr{
								pos:  position{line: 86, col: 67, offset: 3518},
								name: "DIGIT",
							},
							&zeroOrMoreExpr{
								pos: position{line: 86, col: 74, offset: 3525},
								expr: &seqExpr{
									pos: position{line: 86, col: 75, offset: 3526},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 86, col: 75, offset: 3526},
											expr: &ruleRefExpr{
												pos:  position{line: 86, col: 76, offset: 3527},
												name: "EOL",
											},
										},
										&notExpr{
											pos: position{line: 86, col: 80, offset: 3531},
											expr: &litMatcher{
												pos:        position{line: 86, col: 81, offset: 3532},
												val:        ",",
												ignoreCase: false,
											},
										},
										&notExpr{
											pos: position{line: 86, col: 85, offset: 3536},
											expr: &litMatcher{
												pos:        position{line: 86, col: 86, offset: 3537},
												val:        ":",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 86, col: 90, offset: 3541,
										},
									},
								},
							},
							&zeroOrMoreExpr{
								pos: position{line: 86, col: 94, offset: 3545},
								expr: &ruleRefExpr{
									pos:  position{line: 86, col: 94, offset: 3545},
									name: "WS",
								},
							},
							&andExpr{
								pos: position{line: 86, col: 98, offset: 3549},
								expr: &litMatcher{
									pos:        position{line: 86, col: 99, offset: 3550},
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
			pos:  position{line: 87, col: 1, offset: 3554},
			expr: &zeroOrMoreExpr{
				pos: position{line: 87, col: 25, offset: 3578},
				expr: &seqExpr{
					pos: position{line: 87, col: 26, offset: 3579},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 87, col: 26, offset: 3579},
							expr: &ruleRefExpr{
								pos:  position{line: 87, col: 27, offset: 3580},
								name: "EOL",
							},
						},
						&notExpr{
							pos: position{line: 87, col: 31, offset: 3584},
							expr: &litMatcher{
								pos:        position{line: 87, col: 32, offset: 3585},
								val:        ":",
								ignoreCase: false,
							},
						},
						&anyMatcher{
							line: 87, col: 36, offset: 3589,
						},
					},
				},
			},
		},
		{
			name: "DocumentRevisionRemark",
			pos:  position{line: 88, col: 1, offset: 3594},
			expr: &zeroOrMoreExpr{
				pos: position{line: 88, col: 27, offset: 3620},
				expr: &seqExpr{
					pos: position{line: 88, col: 28, offset: 3621},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 88, col: 28, offset: 3621},
							expr: &ruleRefExpr{
								pos:  position{line: 88, col: 29, offset: 3622},
								name: "EOL",
							},
						},
						&anyMatcher{
							line: 88, col: 33, offset: 3626,
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclaration",
			pos:  position{line: 93, col: 1, offset: 3746},
			expr: &choiceExpr{
				pos: position{line: 93, col: 33, offset: 3778},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 93, col: 33, offset: 3778},
						name: "DocumentAttributeDeclarationWithNameOnly",
					},
					&ruleRefExpr{
						pos:  position{line: 93, col: 76, offset: 3821},
						name: "DocumentAttributeDeclarationWithNameAndValue",
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclarationWithNameOnly",
			pos:  position{line: 95, col: 1, offset: 3868},
			expr: &actionExpr{
				pos: position{line: 95, col: 45, offset: 3912},
				run: (*parser).callonDocumentAttributeDeclarationWithNameOnly1,
				expr: &seqExpr{
					pos: position{line: 95, col: 45, offset: 3912},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 95, col: 45, offset: 3912},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 95, col: 49, offset: 3916},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 95, col: 55, offset: 3922},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 95, col: 70, offset: 3937},
							val:        ":",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 95, col: 74, offset: 3941},
							expr: &ruleRefExpr{
								pos:  position{line: 95, col: 74, offset: 3941},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 95, col: 78, offset: 3945},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclarationWithNameAndValue",
			pos:  position{line: 99, col: 1, offset: 4030},
			expr: &actionExpr{
				pos: position{line: 99, col: 49, offset: 4078},
				run: (*parser).callonDocumentAttributeDeclarationWithNameAndValue1,
				expr: &seqExpr{
					pos: position{line: 99, col: 49, offset: 4078},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 99, col: 49, offset: 4078},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 99, col: 53, offset: 4082},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 99, col: 59, offset: 4088},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 99, col: 74, offset: 4103},
							val:        ":",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 99, col: 78, offset: 4107},
							expr: &ruleRefExpr{
								pos:  position{line: 99, col: 78, offset: 4107},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 99, col: 82, offset: 4111},
							label: "value",
							expr: &zeroOrMoreExpr{
								pos: position{line: 99, col: 88, offset: 4117},
								expr: &seqExpr{
									pos: position{line: 99, col: 89, offset: 4118},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 99, col: 89, offset: 4118},
											expr: &ruleRefExpr{
												pos:  position{line: 99, col: 90, offset: 4119},
												name: "NEWLINE",
											},
										},
										&anyMatcher{
											line: 99, col: 98, offset: 4127,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 99, col: 102, offset: 4131},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeReset",
			pos:  position{line: 103, col: 1, offset: 4234},
			expr: &choiceExpr{
				pos: position{line: 103, col: 27, offset: 4260},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 103, col: 27, offset: 4260},
						name: "DocumentAttributeResetWithSectionTitleBangSymbol",
					},
					&ruleRefExpr{
						pos:  position{line: 103, col: 78, offset: 4311},
						name: "DocumentAttributeResetWithTrailingBangSymbol",
					},
				},
			},
		},
		{
			name: "DocumentAttributeResetWithSectionTitleBangSymbol",
			pos:  position{line: 105, col: 1, offset: 4357},
			expr: &actionExpr{
				pos: position{line: 105, col: 53, offset: 4409},
				run: (*parser).callonDocumentAttributeResetWithSectionTitleBangSymbol1,
				expr: &seqExpr{
					pos: position{line: 105, col: 53, offset: 4409},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 105, col: 53, offset: 4409},
							val:        ":!",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 105, col: 58, offset: 4414},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 105, col: 64, offset: 4420},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 105, col: 79, offset: 4435},
							val:        ":",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 105, col: 83, offset: 4439},
							expr: &ruleRefExpr{
								pos:  position{line: 105, col: 83, offset: 4439},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 105, col: 87, offset: 4443},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeResetWithTrailingBangSymbol",
			pos:  position{line: 109, col: 1, offset: 4517},
			expr: &actionExpr{
				pos: position{line: 109, col: 49, offset: 4565},
				run: (*parser).callonDocumentAttributeResetWithTrailingBangSymbol1,
				expr: &seqExpr{
					pos: position{line: 109, col: 49, offset: 4565},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 109, col: 49, offset: 4565},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 109, col: 53, offset: 4569},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 109, col: 59, offset: 4575},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 109, col: 74, offset: 4590},
							val:        "!:",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 109, col: 79, offset: 4595},
							expr: &ruleRefExpr{
								pos:  position{line: 109, col: 79, offset: 4595},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 109, col: 83, offset: 4599},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeSubstitution",
			pos:  position{line: 113, col: 1, offset: 4673},
			expr: &actionExpr{
				pos: position{line: 113, col: 34, offset: 4706},
				run: (*parser).callonDocumentAttributeSubstitution1,
				expr: &seqExpr{
					pos: position{line: 113, col: 34, offset: 4706},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 113, col: 34, offset: 4706},
							val:        "{",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 113, col: 38, offset: 4710},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 113, col: 44, offset: 4716},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 113, col: 59, offset: 4731},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "AttributeName",
			pos:  position{line: 120, col: 1, offset: 4985},
			expr: &seqExpr{
				pos: position{line: 120, col: 18, offset: 5002},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 120, col: 19, offset: 5003},
						alternatives: []interface{}{
							&charClassMatcher{
								pos:        position{line: 120, col: 19, offset: 5003},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
							&charClassMatcher{
								pos:        position{line: 120, col: 27, offset: 5011},
								val:        "[a-z]",
								ranges:     []rune{'a', 'z'},
								ignoreCase: false,
								inverted:   false,
							},
							&charClassMatcher{
								pos:        position{line: 120, col: 35, offset: 5019},
								val:        "[0-9]",
								ranges:     []rune{'0', '9'},
								ignoreCase: false,
								inverted:   false,
							},
							&litMatcher{
								pos:        position{line: 120, col: 43, offset: 5027},
								val:        "_",
								ignoreCase: false,
							},
						},
					},
					&zeroOrMoreExpr{
						pos: position{line: 120, col: 48, offset: 5032},
						expr: &choiceExpr{
							pos: position{line: 120, col: 49, offset: 5033},
							alternatives: []interface{}{
								&charClassMatcher{
									pos:        position{line: 120, col: 49, offset: 5033},
									val:        "[A-Z]",
									ranges:     []rune{'A', 'Z'},
									ignoreCase: false,
									inverted:   false,
								},
								&charClassMatcher{
									pos:        position{line: 120, col: 57, offset: 5041},
									val:        "[a-z]",
									ranges:     []rune{'a', 'z'},
									ignoreCase: false,
									inverted:   false,
								},
								&charClassMatcher{
									pos:        position{line: 120, col: 65, offset: 5049},
									val:        "[0-9]",
									ranges:     []rune{'0', '9'},
									ignoreCase: false,
									inverted:   false,
								},
								&litMatcher{
									pos:        position{line: 120, col: 73, offset: 5057},
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
			pos:  position{line: 125, col: 1, offset: 5177},
			expr: &seqExpr{
				pos: position{line: 125, col: 25, offset: 5201},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 125, col: 25, offset: 5201},
						val:        "toc::[]",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 125, col: 35, offset: 5211},
						name: "NEWLINE",
					},
				},
			},
		},
		{
			name: "Section",
			pos:  position{line: 130, col: 1, offset: 5324},
			expr: &choiceExpr{
				pos: position{line: 130, col: 12, offset: 5335},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 130, col: 12, offset: 5335},
						name: "Section1",
					},
					&ruleRefExpr{
						pos:  position{line: 130, col: 23, offset: 5346},
						name: "Section2",
					},
					&ruleRefExpr{
						pos:  position{line: 130, col: 34, offset: 5357},
						name: "Section3",
					},
					&ruleRefExpr{
						pos:  position{line: 130, col: 45, offset: 5368},
						name: "Section4",
					},
					&ruleRefExpr{
						pos:  position{line: 130, col: 56, offset: 5379},
						name: "Section5",
					},
				},
			},
		},
		{
			name: "Section1",
			pos:  position{line: 133, col: 1, offset: 5390},
			expr: &actionExpr{
				pos: position{line: 133, col: 13, offset: 5402},
				run: (*parser).callonSection11,
				expr: &seqExpr{
					pos: position{line: 133, col: 13, offset: 5402},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 133, col: 13, offset: 5402},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 133, col: 21, offset: 5410},
								name: "Section1Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 133, col: 36, offset: 5425},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 133, col: 46, offset: 5435},
								expr: &ruleRefExpr{
									pos:  position{line: 133, col: 46, offset: 5435},
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
			pos:  position{line: 137, col: 1, offset: 5543},
			expr: &actionExpr{
				pos: position{line: 137, col: 18, offset: 5560},
				run: (*parser).callonSection1Block1,
				expr: &seqExpr{
					pos: position{line: 137, col: 18, offset: 5560},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 137, col: 18, offset: 5560},
							expr: &ruleRefExpr{
								pos:  position{line: 137, col: 19, offset: 5561},
								name: "Section1",
							},
						},
						&labeledExpr{
							pos:   position{line: 137, col: 28, offset: 5570},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 137, col: 37, offset: 5579},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 137, col: 37, offset: 5579},
										name: "Section2",
									},
									&ruleRefExpr{
										pos:  position{line: 137, col: 48, offset: 5590},
										name: "Section3",
									},
									&ruleRefExpr{
										pos:  position{line: 137, col: 59, offset: 5601},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 137, col: 70, offset: 5612},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 137, col: 81, offset: 5623},
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
			pos:  position{line: 141, col: 1, offset: 5685},
			expr: &actionExpr{
				pos: position{line: 141, col: 13, offset: 5697},
				run: (*parser).callonSection21,
				expr: &seqExpr{
					pos: position{line: 141, col: 13, offset: 5697},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 141, col: 13, offset: 5697},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 141, col: 21, offset: 5705},
								name: "Section2Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 141, col: 36, offset: 5720},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 141, col: 46, offset: 5730},
								expr: &ruleRefExpr{
									pos:  position{line: 141, col: 46, offset: 5730},
									name: "Section2Block",
								},
							},
						},
						&andExpr{
							pos: position{line: 141, col: 62, offset: 5746},
							expr: &zeroOrMoreExpr{
								pos: position{line: 141, col: 63, offset: 5747},
								expr: &ruleRefExpr{
									pos:  position{line: 141, col: 64, offset: 5748},
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
			pos:  position{line: 145, col: 1, offset: 5851},
			expr: &actionExpr{
				pos: position{line: 145, col: 18, offset: 5868},
				run: (*parser).callonSection2Block1,
				expr: &seqExpr{
					pos: position{line: 145, col: 18, offset: 5868},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 145, col: 18, offset: 5868},
							expr: &ruleRefExpr{
								pos:  position{line: 145, col: 19, offset: 5869},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 145, col: 28, offset: 5878},
							expr: &ruleRefExpr{
								pos:  position{line: 145, col: 29, offset: 5879},
								name: "Section2",
							},
						},
						&labeledExpr{
							pos:   position{line: 145, col: 38, offset: 5888},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 145, col: 47, offset: 5897},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 145, col: 47, offset: 5897},
										name: "Section3",
									},
									&ruleRefExpr{
										pos:  position{line: 145, col: 58, offset: 5908},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 145, col: 69, offset: 5919},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 145, col: 80, offset: 5930},
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
			pos:  position{line: 149, col: 1, offset: 5992},
			expr: &actionExpr{
				pos: position{line: 149, col: 13, offset: 6004},
				run: (*parser).callonSection31,
				expr: &seqExpr{
					pos: position{line: 149, col: 13, offset: 6004},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 149, col: 13, offset: 6004},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 149, col: 21, offset: 6012},
								name: "Section3Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 149, col: 36, offset: 6027},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 149, col: 46, offset: 6037},
								expr: &ruleRefExpr{
									pos:  position{line: 149, col: 46, offset: 6037},
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
			pos:  position{line: 153, col: 1, offset: 6145},
			expr: &actionExpr{
				pos: position{line: 153, col: 18, offset: 6162},
				run: (*parser).callonSection3Block1,
				expr: &seqExpr{
					pos: position{line: 153, col: 18, offset: 6162},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 153, col: 18, offset: 6162},
							expr: &ruleRefExpr{
								pos:  position{line: 153, col: 19, offset: 6163},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 153, col: 28, offset: 6172},
							expr: &ruleRefExpr{
								pos:  position{line: 153, col: 29, offset: 6173},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 153, col: 38, offset: 6182},
							expr: &ruleRefExpr{
								pos:  position{line: 153, col: 39, offset: 6183},
								name: "Section3",
							},
						},
						&labeledExpr{
							pos:   position{line: 153, col: 48, offset: 6192},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 153, col: 57, offset: 6201},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 153, col: 57, offset: 6201},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 153, col: 68, offset: 6212},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 153, col: 79, offset: 6223},
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
			pos:  position{line: 157, col: 1, offset: 6285},
			expr: &actionExpr{
				pos: position{line: 157, col: 13, offset: 6297},
				run: (*parser).callonSection41,
				expr: &seqExpr{
					pos: position{line: 157, col: 13, offset: 6297},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 157, col: 13, offset: 6297},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 157, col: 21, offset: 6305},
								name: "Section4Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 157, col: 36, offset: 6320},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 157, col: 46, offset: 6330},
								expr: &ruleRefExpr{
									pos:  position{line: 157, col: 46, offset: 6330},
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
			pos:  position{line: 161, col: 1, offset: 6438},
			expr: &actionExpr{
				pos: position{line: 161, col: 18, offset: 6455},
				run: (*parser).callonSection4Block1,
				expr: &seqExpr{
					pos: position{line: 161, col: 18, offset: 6455},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 161, col: 18, offset: 6455},
							expr: &ruleRefExpr{
								pos:  position{line: 161, col: 19, offset: 6456},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 161, col: 28, offset: 6465},
							expr: &ruleRefExpr{
								pos:  position{line: 161, col: 29, offset: 6466},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 161, col: 38, offset: 6475},
							expr: &ruleRefExpr{
								pos:  position{line: 161, col: 39, offset: 6476},
								name: "Section3",
							},
						},
						&notExpr{
							pos: position{line: 161, col: 48, offset: 6485},
							expr: &ruleRefExpr{
								pos:  position{line: 161, col: 49, offset: 6486},
								name: "Section4",
							},
						},
						&labeledExpr{
							pos:   position{line: 161, col: 58, offset: 6495},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 161, col: 67, offset: 6504},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 161, col: 67, offset: 6504},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 161, col: 78, offset: 6515},
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
			pos:  position{line: 165, col: 1, offset: 6577},
			expr: &actionExpr{
				pos: position{line: 165, col: 13, offset: 6589},
				run: (*parser).callonSection51,
				expr: &seqExpr{
					pos: position{line: 165, col: 13, offset: 6589},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 165, col: 13, offset: 6589},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 165, col: 21, offset: 6597},
								name: "Section5Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 165, col: 36, offset: 6612},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 165, col: 46, offset: 6622},
								expr: &ruleRefExpr{
									pos:  position{line: 165, col: 46, offset: 6622},
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
			pos:  position{line: 169, col: 1, offset: 6730},
			expr: &actionExpr{
				pos: position{line: 169, col: 18, offset: 6747},
				run: (*parser).callonSection5Block1,
				expr: &seqExpr{
					pos: position{line: 169, col: 18, offset: 6747},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 169, col: 18, offset: 6747},
							expr: &ruleRefExpr{
								pos:  position{line: 169, col: 19, offset: 6748},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 169, col: 28, offset: 6757},
							expr: &ruleRefExpr{
								pos:  position{line: 169, col: 29, offset: 6758},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 169, col: 38, offset: 6767},
							expr: &ruleRefExpr{
								pos:  position{line: 169, col: 39, offset: 6768},
								name: "Section3",
							},
						},
						&notExpr{
							pos: position{line: 169, col: 48, offset: 6777},
							expr: &ruleRefExpr{
								pos:  position{line: 169, col: 49, offset: 6778},
								name: "Section4",
							},
						},
						&notExpr{
							pos: position{line: 169, col: 58, offset: 6787},
							expr: &ruleRefExpr{
								pos:  position{line: 169, col: 59, offset: 6788},
								name: "Section5",
							},
						},
						&labeledExpr{
							pos:   position{line: 169, col: 68, offset: 6797},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 169, col: 77, offset: 6806},
								name: "BlockElement",
							},
						},
					},
				},
			},
		},
		{
			name: "SectionTitle",
			pos:  position{line: 177, col: 1, offset: 6979},
			expr: &choiceExpr{
				pos: position{line: 177, col: 17, offset: 6995},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 177, col: 17, offset: 6995},
						name: "Section1Title",
					},
					&ruleRefExpr{
						pos:  position{line: 177, col: 33, offset: 7011},
						name: "Section2Title",
					},
					&ruleRefExpr{
						pos:  position{line: 177, col: 49, offset: 7027},
						name: "Section3Title",
					},
					&ruleRefExpr{
						pos:  position{line: 177, col: 65, offset: 7043},
						name: "Section4Title",
					},
					&ruleRefExpr{
						pos:  position{line: 177, col: 81, offset: 7059},
						name: "Section5Title",
					},
				},
			},
		},
		{
			name: "Section1Title",
			pos:  position{line: 179, col: 1, offset: 7074},
			expr: &actionExpr{
				pos: position{line: 179, col: 18, offset: 7091},
				run: (*parser).callonSection1Title1,
				expr: &seqExpr{
					pos: position{line: 179, col: 18, offset: 7091},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 179, col: 18, offset: 7091},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 179, col: 29, offset: 7102},
								expr: &ruleRefExpr{
									pos:  position{line: 179, col: 30, offset: 7103},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 179, col: 49, offset: 7122},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 179, col: 56, offset: 7129},
								val:        "==",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 179, col: 62, offset: 7135},
							expr: &ruleRefExpr{
								pos:  position{line: 179, col: 62, offset: 7135},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 179, col: 66, offset: 7139},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 179, col: 75, offset: 7148},
								name: "InlineContent",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 179, col: 90, offset: 7163},
							expr: &ruleRefExpr{
								pos:  position{line: 179, col: 90, offset: 7163},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 179, col: 94, offset: 7167},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 179, col: 97, offset: 7170},
								expr: &ruleRefExpr{
									pos:  position{line: 179, col: 98, offset: 7171},
									name: "InlineElementID",
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 179, col: 116, offset: 7189},
							expr: &ruleRefExpr{
								pos:  position{line: 179, col: 116, offset: 7189},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 179, col: 120, offset: 7193},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 179, col: 125, offset: 7198},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 179, col: 125, offset: 7198},
									expr: &ruleRefExpr{
										pos:  position{line: 179, col: 125, offset: 7198},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 179, col: 138, offset: 7211},
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
			pos:  position{line: 183, col: 1, offset: 7327},
			expr: &actionExpr{
				pos: position{line: 183, col: 18, offset: 7344},
				run: (*parser).callonSection2Title1,
				expr: &seqExpr{
					pos: position{line: 183, col: 18, offset: 7344},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 183, col: 18, offset: 7344},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 183, col: 29, offset: 7355},
								expr: &ruleRefExpr{
									pos:  position{line: 183, col: 30, offset: 7356},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 183, col: 49, offset: 7375},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 183, col: 56, offset: 7382},
								val:        "===",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 183, col: 63, offset: 7389},
							expr: &ruleRefExpr{
								pos:  position{line: 183, col: 63, offset: 7389},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 183, col: 67, offset: 7393},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 183, col: 76, offset: 7402},
								name: "InlineContent",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 183, col: 91, offset: 7417},
							expr: &ruleRefExpr{
								pos:  position{line: 183, col: 91, offset: 7417},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 183, col: 95, offset: 7421},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 183, col: 98, offset: 7424},
								expr: &ruleRefExpr{
									pos:  position{line: 183, col: 99, offset: 7425},
									name: "InlineElementID",
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 183, col: 117, offset: 7443},
							expr: &ruleRefExpr{
								pos:  position{line: 183, col: 117, offset: 7443},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 183, col: 121, offset: 7447},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 183, col: 126, offset: 7452},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 183, col: 126, offset: 7452},
									expr: &ruleRefExpr{
										pos:  position{line: 183, col: 126, offset: 7452},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 183, col: 139, offset: 7465},
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
			pos:  position{line: 187, col: 1, offset: 7580},
			expr: &actionExpr{
				pos: position{line: 187, col: 18, offset: 7597},
				run: (*parser).callonSection3Title1,
				expr: &seqExpr{
					pos: position{line: 187, col: 18, offset: 7597},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 187, col: 18, offset: 7597},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 187, col: 29, offset: 7608},
								expr: &ruleRefExpr{
									pos:  position{line: 187, col: 30, offset: 7609},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 187, col: 49, offset: 7628},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 187, col: 56, offset: 7635},
								val:        "====",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 187, col: 64, offset: 7643},
							expr: &ruleRefExpr{
								pos:  position{line: 187, col: 64, offset: 7643},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 187, col: 68, offset: 7647},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 187, col: 77, offset: 7656},
								name: "InlineContent",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 187, col: 92, offset: 7671},
							expr: &ruleRefExpr{
								pos:  position{line: 187, col: 92, offset: 7671},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 187, col: 96, offset: 7675},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 187, col: 99, offset: 7678},
								expr: &ruleRefExpr{
									pos:  position{line: 187, col: 100, offset: 7679},
									name: "InlineElementID",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 187, col: 118, offset: 7697},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 187, col: 123, offset: 7702},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 187, col: 123, offset: 7702},
									expr: &ruleRefExpr{
										pos:  position{line: 187, col: 123, offset: 7702},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 187, col: 136, offset: 7715},
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
			pos:  position{line: 191, col: 1, offset: 7830},
			expr: &actionExpr{
				pos: position{line: 191, col: 18, offset: 7847},
				run: (*parser).callonSection4Title1,
				expr: &seqExpr{
					pos: position{line: 191, col: 18, offset: 7847},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 191, col: 18, offset: 7847},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 191, col: 29, offset: 7858},
								expr: &ruleRefExpr{
									pos:  position{line: 191, col: 30, offset: 7859},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 191, col: 49, offset: 7878},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 191, col: 56, offset: 7885},
								val:        "=====",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 191, col: 65, offset: 7894},
							expr: &ruleRefExpr{
								pos:  position{line: 191, col: 65, offset: 7894},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 191, col: 69, offset: 7898},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 191, col: 78, offset: 7907},
								name: "InlineContent",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 191, col: 93, offset: 7922},
							expr: &ruleRefExpr{
								pos:  position{line: 191, col: 93, offset: 7922},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 191, col: 97, offset: 7926},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 191, col: 100, offset: 7929},
								expr: &ruleRefExpr{
									pos:  position{line: 191, col: 101, offset: 7930},
									name: "InlineElementID",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 191, col: 119, offset: 7948},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 191, col: 124, offset: 7953},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 191, col: 124, offset: 7953},
									expr: &ruleRefExpr{
										pos:  position{line: 191, col: 124, offset: 7953},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 191, col: 137, offset: 7966},
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
			pos:  position{line: 195, col: 1, offset: 8081},
			expr: &actionExpr{
				pos: position{line: 195, col: 18, offset: 8098},
				run: (*parser).callonSection5Title1,
				expr: &seqExpr{
					pos: position{line: 195, col: 18, offset: 8098},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 195, col: 18, offset: 8098},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 195, col: 29, offset: 8109},
								expr: &ruleRefExpr{
									pos:  position{line: 195, col: 30, offset: 8110},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 195, col: 49, offset: 8129},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 195, col: 56, offset: 8136},
								val:        "======",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 195, col: 66, offset: 8146},
							expr: &ruleRefExpr{
								pos:  position{line: 195, col: 66, offset: 8146},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 195, col: 70, offset: 8150},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 195, col: 79, offset: 8159},
								name: "InlineContent",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 195, col: 94, offset: 8174},
							expr: &ruleRefExpr{
								pos:  position{line: 195, col: 94, offset: 8174},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 195, col: 98, offset: 8178},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 195, col: 101, offset: 8181},
								expr: &ruleRefExpr{
									pos:  position{line: 195, col: 102, offset: 8182},
									name: "InlineElementID",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 195, col: 120, offset: 8200},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 195, col: 125, offset: 8205},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 195, col: 125, offset: 8205},
									expr: &ruleRefExpr{
										pos:  position{line: 195, col: 125, offset: 8205},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 195, col: 138, offset: 8218},
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
			pos:  position{line: 202, col: 1, offset: 8434},
			expr: &actionExpr{
				pos: position{line: 202, col: 9, offset: 8442},
				run: (*parser).callonList1,
				expr: &seqExpr{
					pos: position{line: 202, col: 9, offset: 8442},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 202, col: 9, offset: 8442},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 202, col: 20, offset: 8453},
								expr: &ruleRefExpr{
									pos:  position{line: 202, col: 21, offset: 8454},
									name: "ListAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 204, col: 5, offset: 8543},
							label: "elements",
							expr: &oneOrMoreExpr{
								pos: position{line: 204, col: 14, offset: 8552},
								expr: &choiceExpr{
									pos: position{line: 204, col: 15, offset: 8553},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 204, col: 15, offset: 8553},
											name: "UnorderedListItem",
										},
										&ruleRefExpr{
											pos:  position{line: 204, col: 35, offset: 8573},
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
			pos:  position{line: 208, col: 1, offset: 8675},
			expr: &actionExpr{
				pos: position{line: 208, col: 18, offset: 8692},
				run: (*parser).callonListAttribute1,
				expr: &seqExpr{
					pos: position{line: 208, col: 18, offset: 8692},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 208, col: 18, offset: 8692},
							label: "attribute",
							expr: &choiceExpr{
								pos: position{line: 208, col: 29, offset: 8703},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 208, col: 29, offset: 8703},
										name: "HorizontalLayout",
									},
									&ruleRefExpr{
										pos:  position{line: 208, col: 48, offset: 8722},
										name: "ListID",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 208, col: 56, offset: 8730},
							name: "NEWLINE",
						},
					},
				},
			},
		},
		{
			name: "ListID",
			pos:  position{line: 212, col: 1, offset: 8769},
			expr: &actionExpr{
				pos: position{line: 212, col: 11, offset: 8779},
				run: (*parser).callonListID1,
				expr: &seqExpr{
					pos: position{line: 212, col: 11, offset: 8779},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 212, col: 11, offset: 8779},
							val:        "[#",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 212, col: 16, offset: 8784},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 212, col: 20, offset: 8788},
								name: "ID",
							},
						},
						&litMatcher{
							pos:        position{line: 212, col: 24, offset: 8792},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "HorizontalLayout",
			pos:  position{line: 216, col: 1, offset: 8858},
			expr: &actionExpr{
				pos: position{line: 216, col: 21, offset: 8878},
				run: (*parser).callonHorizontalLayout1,
				expr: &litMatcher{
					pos:        position{line: 216, col: 21, offset: 8878},
					val:        "[horizontal]",
					ignoreCase: false,
				},
			},
		},
		{
			name: "ListParagraph",
			pos:  position{line: 220, col: 1, offset: 8961},
			expr: &actionExpr{
				pos: position{line: 220, col: 19, offset: 8979},
				run: (*parser).callonListParagraph1,
				expr: &labeledExpr{
					pos:   position{line: 220, col: 19, offset: 8979},
					label: "lines",
					expr: &oneOrMoreExpr{
						pos: position{line: 220, col: 25, offset: 8985},
						expr: &seqExpr{
							pos: position{line: 220, col: 26, offset: 8986},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 220, col: 26, offset: 8986},
									expr: &ruleRefExpr{
										pos:  position{line: 220, col: 28, offset: 8988},
										name: "ListItemContinuation",
									},
								},
								&notExpr{
									pos: position{line: 220, col: 50, offset: 9010},
									expr: &ruleRefExpr{
										pos:  position{line: 220, col: 52, offset: 9012},
										name: "UnorderedListItemPrefix",
									},
								},
								&notExpr{
									pos: position{line: 220, col: 77, offset: 9037},
									expr: &seqExpr{
										pos: position{line: 220, col: 79, offset: 9039},
										exprs: []interface{}{
											&ruleRefExpr{
												pos:  position{line: 220, col: 79, offset: 9039},
												name: "LabeledListItemTerm",
											},
											&ruleRefExpr{
												pos:  position{line: 220, col: 99, offset: 9059},
												name: "LabeledListItemSeparator",
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 220, col: 125, offset: 9085},
									name: "InlineContentWithTrailingSpaces",
								},
								&ruleRefExpr{
									pos:  position{line: 220, col: 157, offset: 9117},
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
			pos:  position{line: 224, col: 1, offset: 9186},
			expr: &actionExpr{
				pos: position{line: 224, col: 25, offset: 9210},
				run: (*parser).callonListItemContinuation1,
				expr: &seqExpr{
					pos: position{line: 224, col: 25, offset: 9210},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 224, col: 25, offset: 9210},
							val:        "+",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 224, col: 29, offset: 9214},
							expr: &ruleRefExpr{
								pos:  position{line: 224, col: 29, offset: 9214},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 224, col: 33, offset: 9218},
							name: "NEWLINE",
						},
					},
				},
			},
		},
		{
			name: "ContinuedBlockElement",
			pos:  position{line: 228, col: 1, offset: 9274},
			expr: &actionExpr{
				pos: position{line: 228, col: 26, offset: 9299},
				run: (*parser).callonContinuedBlockElement1,
				expr: &seqExpr{
					pos: position{line: 228, col: 26, offset: 9299},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 228, col: 26, offset: 9299},
							name: "ListItemContinuation",
						},
						&labeledExpr{
							pos:   position{line: 228, col: 47, offset: 9320},
							label: "element",
							expr: &ruleRefExpr{
								pos:  position{line: 228, col: 55, offset: 9328},
								name: "BlockElement",
							},
						},
					},
				},
			},
		},
		{
			name: "UnorderedListItem",
			pos:  position{line: 235, col: 1, offset: 9481},
			expr: &actionExpr{
				pos: position{line: 235, col: 22, offset: 9502},
				run: (*parser).callonUnorderedListItem1,
				expr: &seqExpr{
					pos: position{line: 235, col: 22, offset: 9502},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 235, col: 22, offset: 9502},
							label: "level",
							expr: &ruleRefExpr{
								pos:  position{line: 235, col: 29, offset: 9509},
								name: "UnorderedListItemPrefix",
							},
						},
						&labeledExpr{
							pos:   position{line: 235, col: 54, offset: 9534},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 235, col: 63, offset: 9543},
								name: "UnorderedListItemContent",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 235, col: 89, offset: 9569},
							expr: &ruleRefExpr{
								pos:  position{line: 235, col: 89, offset: 9569},
								name: "BlankLine",
							},
						},
					},
				},
			},
		},
		{
			name: "UnorderedListItemPrefix",
			pos:  position{line: 239, col: 1, offset: 9660},
			expr: &actionExpr{
				pos: position{line: 239, col: 28, offset: 9687},
				run: (*parser).callonUnorderedListItemPrefix1,
				expr: &seqExpr{
					pos: position{line: 239, col: 28, offset: 9687},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 239, col: 28, offset: 9687},
							expr: &ruleRefExpr{
								pos:  position{line: 239, col: 28, offset: 9687},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 239, col: 32, offset: 9691},
							label: "level",
							expr: &choiceExpr{
								pos: position{line: 239, col: 39, offset: 9698},
								alternatives: []interface{}{
									&oneOrMoreExpr{
										pos: position{line: 239, col: 39, offset: 9698},
										expr: &litMatcher{
											pos:        position{line: 239, col: 39, offset: 9698},
											val:        "*",
											ignoreCase: false,
										},
									},
									&litMatcher{
										pos:        position{line: 239, col: 46, offset: 9705},
										val:        "-",
										ignoreCase: false,
									},
								},
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 239, col: 51, offset: 9710},
							expr: &ruleRefExpr{
								pos:  position{line: 239, col: 51, offset: 9710},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "UnorderedListItemContent",
			pos:  position{line: 243, col: 1, offset: 9808},
			expr: &actionExpr{
				pos: position{line: 243, col: 29, offset: 9836},
				run: (*parser).callonUnorderedListItemContent1,
				expr: &labeledExpr{
					pos:   position{line: 243, col: 29, offset: 9836},
					label: "elements",
					expr: &seqExpr{
						pos: position{line: 243, col: 39, offset: 9846},
						exprs: []interface{}{
							&oneOrMoreExpr{
								pos: position{line: 243, col: 39, offset: 9846},
								expr: &ruleRefExpr{
									pos:  position{line: 243, col: 39, offset: 9846},
									name: "ListParagraph",
								},
							},
							&zeroOrMoreExpr{
								pos: position{line: 243, col: 54, offset: 9861},
								expr: &ruleRefExpr{
									pos:  position{line: 243, col: 54, offset: 9861},
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
			pos:  position{line: 250, col: 1, offset: 10180},
			expr: &choiceExpr{
				pos: position{line: 250, col: 20, offset: 10199},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 250, col: 20, offset: 10199},
						run: (*parser).callonLabeledListItem2,
						expr: &seqExpr{
							pos: position{line: 250, col: 20, offset: 10199},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 250, col: 20, offset: 10199},
									label: "term",
									expr: &ruleRefExpr{
										pos:  position{line: 250, col: 26, offset: 10205},
										name: "LabeledListItemTerm",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 250, col: 47, offset: 10226},
									name: "LabeledListItemSeparator",
								},
								&labeledExpr{
									pos:   position{line: 250, col: 72, offset: 10251},
									label: "description",
									expr: &ruleRefExpr{
										pos:  position{line: 250, col: 85, offset: 10264},
										name: "LabeledListItemDescription",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 252, col: 6, offset: 10391},
						run: (*parser).callonLabeledListItem9,
						expr: &seqExpr{
							pos: position{line: 252, col: 6, offset: 10391},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 252, col: 6, offset: 10391},
									label: "term",
									expr: &ruleRefExpr{
										pos:  position{line: 252, col: 12, offset: 10397},
										name: "LabeledListItemTerm",
									},
								},
								&litMatcher{
									pos:        position{line: 252, col: 34, offset: 10419},
									val:        "::",
									ignoreCase: false,
								},
								&zeroOrMoreExpr{
									pos: position{line: 252, col: 39, offset: 10424},
									expr: &ruleRefExpr{
										pos:  position{line: 252, col: 39, offset: 10424},
										name: "WS",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 252, col: 43, offset: 10428},
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
			pos:  position{line: 256, col: 1, offset: 10565},
			expr: &actionExpr{
				pos: position{line: 256, col: 24, offset: 10588},
				run: (*parser).callonLabeledListItemTerm1,
				expr: &labeledExpr{
					pos:   position{line: 256, col: 24, offset: 10588},
					label: "term",
					expr: &zeroOrMoreExpr{
						pos: position{line: 256, col: 29, offset: 10593},
						expr: &seqExpr{
							pos: position{line: 256, col: 30, offset: 10594},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 256, col: 30, offset: 10594},
									expr: &ruleRefExpr{
										pos:  position{line: 256, col: 31, offset: 10595},
										name: "NEWLINE",
									},
								},
								&notExpr{
									pos: position{line: 256, col: 39, offset: 10603},
									expr: &litMatcher{
										pos:        position{line: 256, col: 40, offset: 10604},
										val:        "::",
										ignoreCase: false,
									},
								},
								&anyMatcher{
									line: 256, col: 45, offset: 10609,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "LabeledListItemSeparator",
			pos:  position{line: 261, col: 1, offset: 10700},
			expr: &seqExpr{
				pos: position{line: 261, col: 30, offset: 10729},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 261, col: 30, offset: 10729},
						val:        "::",
						ignoreCase: false,
					},
					&oneOrMoreExpr{
						pos: position{line: 261, col: 35, offset: 10734},
						expr: &choiceExpr{
							pos: position{line: 261, col: 36, offset: 10735},
							alternatives: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 261, col: 36, offset: 10735},
									name: "WS",
								},
								&ruleRefExpr{
									pos:  position{line: 261, col: 41, offset: 10740},
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
			pos:  position{line: 263, col: 1, offset: 10751},
			expr: &actionExpr{
				pos: position{line: 263, col: 31, offset: 10781},
				run: (*parser).callonLabeledListItemDescription1,
				expr: &labeledExpr{
					pos:   position{line: 263, col: 31, offset: 10781},
					label: "elements",
					expr: &zeroOrMoreExpr{
						pos: position{line: 263, col: 40, offset: 10790},
						expr: &choiceExpr{
							pos: position{line: 263, col: 41, offset: 10791},
							alternatives: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 263, col: 41, offset: 10791},
									name: "ListParagraph",
								},
								&ruleRefExpr{
									pos:  position{line: 263, col: 57, offset: 10807},
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
			pos:  position{line: 272, col: 1, offset: 11157},
			expr: &actionExpr{
				pos: position{line: 272, col: 14, offset: 11170},
				run: (*parser).callonParagraph1,
				expr: &seqExpr{
					pos: position{line: 272, col: 14, offset: 11170},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 272, col: 14, offset: 11170},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 272, col: 25, offset: 11181},
								expr: &ruleRefExpr{
									pos:  position{line: 272, col: 26, offset: 11182},
									name: "ElementAttribute",
								},
							},
						},
						&notExpr{
							pos: position{line: 272, col: 45, offset: 11201},
							expr: &seqExpr{
								pos: position{line: 272, col: 47, offset: 11203},
								exprs: []interface{}{
									&oneOrMoreExpr{
										pos: position{line: 272, col: 47, offset: 11203},
										expr: &litMatcher{
											pos:        position{line: 272, col: 47, offset: 11203},
											val:        "=",
											ignoreCase: false,
										},
									},
									&oneOrMoreExpr{
										pos: position{line: 272, col: 52, offset: 11208},
										expr: &ruleRefExpr{
											pos:  position{line: 272, col: 52, offset: 11208},
											name: "WS",
										},
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 272, col: 57, offset: 11213},
							label: "lines",
							expr: &oneOrMoreExpr{
								pos: position{line: 272, col: 63, offset: 11219},
								expr: &seqExpr{
									pos: position{line: 272, col: 64, offset: 11220},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 272, col: 64, offset: 11220},
											name: "InlineContentWithTrailingSpaces",
										},
										&ruleRefExpr{
											pos:  position{line: 272, col: 96, offset: 11252},
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
			pos:  position{line: 278, col: 1, offset: 11542},
			expr: &actionExpr{
				pos: position{line: 278, col: 36, offset: 11577},
				run: (*parser).callonInlineContentWithTrailingSpaces1,
				expr: &seqExpr{
					pos: position{line: 278, col: 36, offset: 11577},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 278, col: 36, offset: 11577},
							expr: &ruleRefExpr{
								pos:  position{line: 278, col: 37, offset: 11578},
								name: "BlockDelimiter",
							},
						},
						&labeledExpr{
							pos:   position{line: 278, col: 52, offset: 11593},
							label: "elements",
							expr: &oneOrMoreExpr{
								pos: position{line: 278, col: 61, offset: 11602},
								expr: &seqExpr{
									pos: position{line: 278, col: 62, offset: 11603},
									exprs: []interface{}{
										&zeroOrMoreExpr{
											pos: position{line: 278, col: 62, offset: 11603},
											expr: &ruleRefExpr{
												pos:  position{line: 278, col: 62, offset: 11603},
												name: "WS",
											},
										},
										&notExpr{
											pos: position{line: 278, col: 66, offset: 11607},
											expr: &ruleRefExpr{
												pos:  position{line: 278, col: 67, offset: 11608},
												name: "InlineElementID",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 278, col: 83, offset: 11624},
											name: "InlineElement",
										},
										&zeroOrMoreExpr{
											pos: position{line: 278, col: 97, offset: 11638},
											expr: &ruleRefExpr{
												pos:  position{line: 278, col: 97, offset: 11638},
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
			pos:  position{line: 282, col: 1, offset: 11750},
			expr: &actionExpr{
				pos: position{line: 282, col: 18, offset: 11767},
				run: (*parser).callonInlineContent1,
				expr: &seqExpr{
					pos: position{line: 282, col: 18, offset: 11767},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 282, col: 18, offset: 11767},
							expr: &ruleRefExpr{
								pos:  position{line: 282, col: 19, offset: 11768},
								name: "BlockDelimiter",
							},
						},
						&labeledExpr{
							pos:   position{line: 282, col: 34, offset: 11783},
							label: "elements",
							expr: &oneOrMoreExpr{
								pos: position{line: 282, col: 43, offset: 11792},
								expr: &seqExpr{
									pos: position{line: 282, col: 44, offset: 11793},
									exprs: []interface{}{
										&zeroOrMoreExpr{
											pos: position{line: 282, col: 44, offset: 11793},
											expr: &ruleRefExpr{
												pos:  position{line: 282, col: 44, offset: 11793},
												name: "WS",
											},
										},
										&notExpr{
											pos: position{line: 282, col: 48, offset: 11797},
											expr: &ruleRefExpr{
												pos:  position{line: 282, col: 49, offset: 11798},
												name: "InlineElementID",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 282, col: 65, offset: 11814},
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
			pos:  position{line: 286, col: 1, offset: 11936},
			expr: &choiceExpr{
				pos: position{line: 286, col: 19, offset: 11954},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 286, col: 19, offset: 11954},
						name: "CrossReference",
					},
					&ruleRefExpr{
						pos:  position{line: 286, col: 36, offset: 11971},
						name: "Passthrough",
					},
					&ruleRefExpr{
						pos:  position{line: 286, col: 50, offset: 11985},
						name: "InlineImage",
					},
					&ruleRefExpr{
						pos:  position{line: 286, col: 64, offset: 11999},
						name: "QuotedText",
					},
					&ruleRefExpr{
						pos:  position{line: 286, col: 77, offset: 12012},
						name: "Link",
					},
					&ruleRefExpr{
						pos:  position{line: 286, col: 84, offset: 12019},
						name: "DocumentAttributeSubstitution",
					},
					&ruleRefExpr{
						pos:  position{line: 286, col: 116, offset: 12051},
						name: "Characters",
					},
				},
			},
		},
		{
			name: "Admonition",
			pos:  position{line: 293, col: 1, offset: 12323},
			expr: &choiceExpr{
				pos: position{line: 293, col: 15, offset: 12337},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 293, col: 15, offset: 12337},
						run: (*parser).callonAdmonition2,
						expr: &seqExpr{
							pos: position{line: 293, col: 15, offset: 12337},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 293, col: 15, offset: 12337},
									label: "attributes",
									expr: &zeroOrMoreExpr{
										pos: position{line: 293, col: 26, offset: 12348},
										expr: &ruleRefExpr{
											pos:  position{line: 293, col: 27, offset: 12349},
											name: "ElementAttribute",
										},
									},
								},
								&notExpr{
									pos: position{line: 293, col: 46, offset: 12368},
									expr: &seqExpr{
										pos: position{line: 293, col: 48, offset: 12370},
										exprs: []interface{}{
											&oneOrMoreExpr{
												pos: position{line: 293, col: 48, offset: 12370},
												expr: &litMatcher{
													pos:        position{line: 293, col: 48, offset: 12370},
													val:        "=",
													ignoreCase: false,
												},
											},
											&oneOrMoreExpr{
												pos: position{line: 293, col: 53, offset: 12375},
												expr: &ruleRefExpr{
													pos:  position{line: 293, col: 53, offset: 12375},
													name: "WS",
												},
											},
										},
									},
								},
								&labeledExpr{
									pos:   position{line: 293, col: 58, offset: 12380},
									label: "t",
									expr: &ruleRefExpr{
										pos:  position{line: 293, col: 61, offset: 12383},
										name: "AdmonitionKind",
									},
								},
								&litMatcher{
									pos:        position{line: 293, col: 77, offset: 12399},
									val:        ": ",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 293, col: 82, offset: 12404},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 293, col: 91, offset: 12413},
										name: "AdmonitionParagraph",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 296, col: 1, offset: 12554},
						run: (*parser).callonAdmonition18,
						expr: &seqExpr{
							pos: position{line: 296, col: 1, offset: 12554},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 296, col: 1, offset: 12554},
									label: "attributes",
									expr: &zeroOrMoreExpr{
										pos: position{line: 296, col: 12, offset: 12565},
										expr: &ruleRefExpr{
											pos:  position{line: 296, col: 13, offset: 12566},
											name: "ElementAttribute",
										},
									},
								},
								&litMatcher{
									pos:        position{line: 296, col: 32, offset: 12585},
									val:        "[",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 296, col: 36, offset: 12589},
									label: "t",
									expr: &ruleRefExpr{
										pos:  position{line: 296, col: 39, offset: 12592},
										name: "AdmonitionKind",
									},
								},
								&litMatcher{
									pos:        position{line: 296, col: 55, offset: 12608},
									val:        "]",
									ignoreCase: false,
								},
								&zeroOrMoreExpr{
									pos: position{line: 296, col: 59, offset: 12612},
									expr: &ruleRefExpr{
										pos:  position{line: 296, col: 59, offset: 12612},
										name: "WS",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 296, col: 63, offset: 12616},
									name: "NEWLINE",
								},
								&labeledExpr{
									pos:   position{line: 296, col: 71, offset: 12624},
									label: "otherAttributes",
									expr: &zeroOrMoreExpr{
										pos: position{line: 296, col: 87, offset: 12640},
										expr: &ruleRefExpr{
											pos:  position{line: 296, col: 88, offset: 12641},
											name: "ElementAttribute",
										},
									},
								},
								&labeledExpr{
									pos:   position{line: 296, col: 107, offset: 12660},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 296, col: 116, offset: 12669},
										name: "AdmonitionParagraph",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "AdmonitionParagraph",
			pos:  position{line: 300, col: 1, offset: 12848},
			expr: &actionExpr{
				pos: position{line: 300, col: 24, offset: 12871},
				run: (*parser).callonAdmonitionParagraph1,
				expr: &labeledExpr{
					pos:   position{line: 300, col: 24, offset: 12871},
					label: "lines",
					expr: &oneOrMoreExpr{
						pos: position{line: 300, col: 30, offset: 12877},
						expr: &seqExpr{
							pos: position{line: 300, col: 31, offset: 12878},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 300, col: 31, offset: 12878},
									name: "InlineContentWithTrailingSpaces",
								},
								&ruleRefExpr{
									pos:  position{line: 300, col: 63, offset: 12910},
									name: "EOL",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "AdmonitionKind",
			pos:  position{line: 304, col: 1, offset: 12984},
			expr: &choiceExpr{
				pos: position{line: 304, col: 19, offset: 13002},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 304, col: 19, offset: 13002},
						run: (*parser).callonAdmonitionKind2,
						expr: &litMatcher{
							pos:        position{line: 304, col: 19, offset: 13002},
							val:        "TIP",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 306, col: 5, offset: 13040},
						run: (*parser).callonAdmonitionKind4,
						expr: &litMatcher{
							pos:        position{line: 306, col: 5, offset: 13040},
							val:        "NOTE",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 308, col: 5, offset: 13080},
						run: (*parser).callonAdmonitionKind6,
						expr: &litMatcher{
							pos:        position{line: 308, col: 5, offset: 13080},
							val:        "IMPORTANT",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 310, col: 5, offset: 13130},
						run: (*parser).callonAdmonitionKind8,
						expr: &litMatcher{
							pos:        position{line: 310, col: 5, offset: 13130},
							val:        "WARNING",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 312, col: 5, offset: 13176},
						run: (*parser).callonAdmonitionKind10,
						expr: &litMatcher{
							pos:        position{line: 312, col: 5, offset: 13176},
							val:        "CAUTION",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "QuotedText",
			pos:  position{line: 319, col: 1, offset: 13460},
			expr: &choiceExpr{
				pos: position{line: 319, col: 15, offset: 13474},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 319, col: 15, offset: 13474},
						name: "BoldText",
					},
					&ruleRefExpr{
						pos:  position{line: 319, col: 26, offset: 13485},
						name: "ItalicText",
					},
					&ruleRefExpr{
						pos:  position{line: 319, col: 39, offset: 13498},
						name: "MonospaceText",
					},
					&ruleRefExpr{
						pos:  position{line: 320, col: 13, offset: 13526},
						name: "EscapedBoldText",
					},
					&ruleRefExpr{
						pos:  position{line: 320, col: 31, offset: 13544},
						name: "EscapedItalicText",
					},
					&ruleRefExpr{
						pos:  position{line: 320, col: 51, offset: 13564},
						name: "EscapedMonospaceText",
					},
				},
			},
		},
		{
			name: "BoldText",
			pos:  position{line: 322, col: 1, offset: 13586},
			expr: &choiceExpr{
				pos: position{line: 322, col: 13, offset: 13598},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 322, col: 13, offset: 13598},
						run: (*parser).callonBoldText2,
						expr: &seqExpr{
							pos: position{line: 322, col: 13, offset: 13598},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 322, col: 13, offset: 13598},
									expr: &litMatcher{
										pos:        position{line: 322, col: 14, offset: 13599},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 322, col: 19, offset: 13604},
									val:        "**",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 322, col: 24, offset: 13609},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 322, col: 33, offset: 13618},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 322, col: 52, offset: 13637},
									val:        "**",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 324, col: 5, offset: 13762},
						run: (*parser).callonBoldText10,
						expr: &seqExpr{
							pos: position{line: 324, col: 5, offset: 13762},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 324, col: 5, offset: 13762},
									expr: &litMatcher{
										pos:        position{line: 324, col: 6, offset: 13763},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 324, col: 11, offset: 13768},
									val:        "**",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 324, col: 16, offset: 13773},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 324, col: 25, offset: 13782},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 324, col: 44, offset: 13801},
									val:        "*",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 327, col: 5, offset: 13966},
						run: (*parser).callonBoldText18,
						expr: &seqExpr{
							pos: position{line: 327, col: 5, offset: 13966},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 327, col: 5, offset: 13966},
									expr: &litMatcher{
										pos:        position{line: 327, col: 6, offset: 13967},
										val:        "\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 327, col: 10, offset: 13971},
									val:        "*",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 327, col: 14, offset: 13975},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 327, col: 23, offset: 13984},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 327, col: 42, offset: 14003},
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
			pos:  position{line: 331, col: 1, offset: 14103},
			expr: &choiceExpr{
				pos: position{line: 331, col: 20, offset: 14122},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 331, col: 20, offset: 14122},
						run: (*parser).callonEscapedBoldText2,
						expr: &seqExpr{
							pos: position{line: 331, col: 20, offset: 14122},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 331, col: 20, offset: 14122},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 331, col: 33, offset: 14135},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 331, col: 33, offset: 14135},
												val:        "\\\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 331, col: 38, offset: 14140},
												expr: &litMatcher{
													pos:        position{line: 331, col: 38, offset: 14140},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 331, col: 44, offset: 14146},
									val:        "**",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 331, col: 49, offset: 14151},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 331, col: 58, offset: 14160},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 331, col: 77, offset: 14179},
									val:        "**",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 333, col: 5, offset: 14334},
						run: (*parser).callonEscapedBoldText13,
						expr: &seqExpr{
							pos: position{line: 333, col: 5, offset: 14334},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 333, col: 5, offset: 14334},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 333, col: 18, offset: 14347},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 333, col: 18, offset: 14347},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 333, col: 22, offset: 14351},
												expr: &litMatcher{
													pos:        position{line: 333, col: 22, offset: 14351},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 333, col: 28, offset: 14357},
									val:        "**",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 333, col: 33, offset: 14362},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 333, col: 42, offset: 14371},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 333, col: 61, offset: 14390},
									val:        "*",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 336, col: 5, offset: 14584},
						run: (*parser).callonEscapedBoldText24,
						expr: &seqExpr{
							pos: position{line: 336, col: 5, offset: 14584},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 336, col: 5, offset: 14584},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 336, col: 18, offset: 14597},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 336, col: 18, offset: 14597},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 336, col: 22, offset: 14601},
												expr: &litMatcher{
													pos:        position{line: 336, col: 22, offset: 14601},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 336, col: 28, offset: 14607},
									val:        "*",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 336, col: 32, offset: 14611},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 336, col: 41, offset: 14620},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 336, col: 60, offset: 14639},
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
			pos:  position{line: 340, col: 1, offset: 14791},
			expr: &choiceExpr{
				pos: position{line: 340, col: 15, offset: 14805},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 340, col: 15, offset: 14805},
						run: (*parser).callonItalicText2,
						expr: &seqExpr{
							pos: position{line: 340, col: 15, offset: 14805},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 340, col: 15, offset: 14805},
									expr: &litMatcher{
										pos:        position{line: 340, col: 16, offset: 14806},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 340, col: 21, offset: 14811},
									val:        "__",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 340, col: 26, offset: 14816},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 340, col: 35, offset: 14825},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 340, col: 54, offset: 14844},
									val:        "__",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 342, col: 5, offset: 14925},
						run: (*parser).callonItalicText10,
						expr: &seqExpr{
							pos: position{line: 342, col: 5, offset: 14925},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 342, col: 5, offset: 14925},
									expr: &litMatcher{
										pos:        position{line: 342, col: 6, offset: 14926},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 342, col: 11, offset: 14931},
									val:        "__",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 342, col: 16, offset: 14936},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 342, col: 25, offset: 14945},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 342, col: 44, offset: 14964},
									val:        "_",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 345, col: 5, offset: 15131},
						run: (*parser).callonItalicText18,
						expr: &seqExpr{
							pos: position{line: 345, col: 5, offset: 15131},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 345, col: 5, offset: 15131},
									expr: &litMatcher{
										pos:        position{line: 345, col: 6, offset: 15132},
										val:        "\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 345, col: 10, offset: 15136},
									val:        "_",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 345, col: 14, offset: 15140},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 345, col: 23, offset: 15149},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 345, col: 42, offset: 15168},
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
			pos:  position{line: 349, col: 1, offset: 15247},
			expr: &choiceExpr{
				pos: position{line: 349, col: 22, offset: 15268},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 349, col: 22, offset: 15268},
						run: (*parser).callonEscapedItalicText2,
						expr: &seqExpr{
							pos: position{line: 349, col: 22, offset: 15268},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 349, col: 22, offset: 15268},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 349, col: 35, offset: 15281},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 349, col: 35, offset: 15281},
												val:        "\\\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 349, col: 40, offset: 15286},
												expr: &litMatcher{
													pos:        position{line: 349, col: 40, offset: 15286},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 349, col: 46, offset: 15292},
									val:        "__",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 349, col: 51, offset: 15297},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 349, col: 60, offset: 15306},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 349, col: 79, offset: 15325},
									val:        "__",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 351, col: 5, offset: 15480},
						run: (*parser).callonEscapedItalicText13,
						expr: &seqExpr{
							pos: position{line: 351, col: 5, offset: 15480},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 351, col: 5, offset: 15480},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 351, col: 18, offset: 15493},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 351, col: 18, offset: 15493},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 351, col: 22, offset: 15497},
												expr: &litMatcher{
													pos:        position{line: 351, col: 22, offset: 15497},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 351, col: 28, offset: 15503},
									val:        "__",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 351, col: 33, offset: 15508},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 351, col: 42, offset: 15517},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 351, col: 61, offset: 15536},
									val:        "_",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 354, col: 5, offset: 15730},
						run: (*parser).callonEscapedItalicText24,
						expr: &seqExpr{
							pos: position{line: 354, col: 5, offset: 15730},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 354, col: 5, offset: 15730},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 354, col: 18, offset: 15743},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 354, col: 18, offset: 15743},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 354, col: 22, offset: 15747},
												expr: &litMatcher{
													pos:        position{line: 354, col: 22, offset: 15747},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 354, col: 28, offset: 15753},
									val:        "_",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 354, col: 32, offset: 15757},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 354, col: 41, offset: 15766},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 354, col: 60, offset: 15785},
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
			pos:  position{line: 358, col: 1, offset: 15937},
			expr: &choiceExpr{
				pos: position{line: 358, col: 18, offset: 15954},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 358, col: 18, offset: 15954},
						run: (*parser).callonMonospaceText2,
						expr: &seqExpr{
							pos: position{line: 358, col: 18, offset: 15954},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 358, col: 18, offset: 15954},
									expr: &litMatcher{
										pos:        position{line: 358, col: 19, offset: 15955},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 358, col: 24, offset: 15960},
									val:        "``",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 358, col: 29, offset: 15965},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 358, col: 38, offset: 15974},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 358, col: 57, offset: 15993},
									val:        "``",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 360, col: 5, offset: 16123},
						run: (*parser).callonMonospaceText10,
						expr: &seqExpr{
							pos: position{line: 360, col: 5, offset: 16123},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 360, col: 5, offset: 16123},
									expr: &litMatcher{
										pos:        position{line: 360, col: 6, offset: 16124},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 360, col: 11, offset: 16129},
									val:        "``",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 360, col: 16, offset: 16134},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 360, col: 25, offset: 16143},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 360, col: 44, offset: 16162},
									val:        "`",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 363, col: 5, offset: 16332},
						run: (*parser).callonMonospaceText18,
						expr: &seqExpr{
							pos: position{line: 363, col: 5, offset: 16332},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 363, col: 5, offset: 16332},
									expr: &litMatcher{
										pos:        position{line: 363, col: 6, offset: 16333},
										val:        "\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 363, col: 10, offset: 16337},
									val:        "`",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 363, col: 14, offset: 16341},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 363, col: 23, offset: 16350},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 363, col: 42, offset: 16369},
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
			pos:  position{line: 367, col: 1, offset: 16496},
			expr: &choiceExpr{
				pos: position{line: 367, col: 25, offset: 16520},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 367, col: 25, offset: 16520},
						run: (*parser).callonEscapedMonospaceText2,
						expr: &seqExpr{
							pos: position{line: 367, col: 25, offset: 16520},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 367, col: 25, offset: 16520},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 367, col: 38, offset: 16533},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 367, col: 38, offset: 16533},
												val:        "\\\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 367, col: 43, offset: 16538},
												expr: &litMatcher{
													pos:        position{line: 367, col: 43, offset: 16538},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 367, col: 49, offset: 16544},
									val:        "``",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 367, col: 54, offset: 16549},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 367, col: 63, offset: 16558},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 367, col: 82, offset: 16577},
									val:        "``",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 369, col: 5, offset: 16732},
						run: (*parser).callonEscapedMonospaceText13,
						expr: &seqExpr{
							pos: position{line: 369, col: 5, offset: 16732},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 369, col: 5, offset: 16732},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 369, col: 18, offset: 16745},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 369, col: 18, offset: 16745},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 369, col: 22, offset: 16749},
												expr: &litMatcher{
													pos:        position{line: 369, col: 22, offset: 16749},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 369, col: 28, offset: 16755},
									val:        "``",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 369, col: 33, offset: 16760},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 369, col: 42, offset: 16769},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 369, col: 61, offset: 16788},
									val:        "`",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 372, col: 5, offset: 16982},
						run: (*parser).callonEscapedMonospaceText24,
						expr: &seqExpr{
							pos: position{line: 372, col: 5, offset: 16982},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 372, col: 5, offset: 16982},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 372, col: 18, offset: 16995},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 372, col: 18, offset: 16995},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 372, col: 22, offset: 16999},
												expr: &litMatcher{
													pos:        position{line: 372, col: 22, offset: 16999},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 372, col: 28, offset: 17005},
									val:        "`",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 372, col: 32, offset: 17009},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 372, col: 41, offset: 17018},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 372, col: 60, offset: 17037},
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
			pos:  position{line: 376, col: 1, offset: 17189},
			expr: &seqExpr{
				pos: position{line: 376, col: 22, offset: 17210},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 376, col: 22, offset: 17210},
						name: "QuotedTextContentElement",
					},
					&zeroOrMoreExpr{
						pos: position{line: 376, col: 47, offset: 17235},
						expr: &seqExpr{
							pos: position{line: 376, col: 48, offset: 17236},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 376, col: 48, offset: 17236},
									expr: &ruleRefExpr{
										pos:  position{line: 376, col: 48, offset: 17236},
										name: "WS",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 376, col: 52, offset: 17240},
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
			pos:  position{line: 378, col: 1, offset: 17268},
			expr: &choiceExpr{
				pos: position{line: 378, col: 29, offset: 17296},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 378, col: 29, offset: 17296},
						name: "QuotedText",
					},
					&ruleRefExpr{
						pos:  position{line: 378, col: 42, offset: 17309},
						name: "QuotedTextCharacters",
					},
					&ruleRefExpr{
						pos:  position{line: 378, col: 65, offset: 17332},
						name: "CharactersWithQuotePunctuation",
					},
				},
			},
		},
		{
			name: "QuotedTextCharacters",
			pos:  position{line: 380, col: 1, offset: 17467},
			expr: &oneOrMoreExpr{
				pos: position{line: 380, col: 25, offset: 17491},
				expr: &seqExpr{
					pos: position{line: 380, col: 26, offset: 17492},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 380, col: 26, offset: 17492},
							expr: &ruleRefExpr{
								pos:  position{line: 380, col: 27, offset: 17493},
								name: "NEWLINE",
							},
						},
						&notExpr{
							pos: position{line: 380, col: 35, offset: 17501},
							expr: &ruleRefExpr{
								pos:  position{line: 380, col: 36, offset: 17502},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 380, col: 39, offset: 17505},
							expr: &litMatcher{
								pos:        position{line: 380, col: 40, offset: 17506},
								val:        "*",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 380, col: 44, offset: 17510},
							expr: &litMatcher{
								pos:        position{line: 380, col: 45, offset: 17511},
								val:        "_",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 380, col: 49, offset: 17515},
							expr: &litMatcher{
								pos:        position{line: 380, col: 50, offset: 17516},
								val:        "`",
								ignoreCase: false,
							},
						},
						&anyMatcher{
							line: 380, col: 54, offset: 17520,
						},
					},
				},
			},
		},
		{
			name: "CharactersWithQuotePunctuation",
			pos:  position{line: 382, col: 1, offset: 17563},
			expr: &actionExpr{
				pos: position{line: 382, col: 35, offset: 17597},
				run: (*parser).callonCharactersWithQuotePunctuation1,
				expr: &oneOrMoreExpr{
					pos: position{line: 382, col: 35, offset: 17597},
					expr: &seqExpr{
						pos: position{line: 382, col: 36, offset: 17598},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 382, col: 36, offset: 17598},
								expr: &ruleRefExpr{
									pos:  position{line: 382, col: 37, offset: 17599},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 382, col: 45, offset: 17607},
								expr: &ruleRefExpr{
									pos:  position{line: 382, col: 46, offset: 17608},
									name: "WS",
								},
							},
							&anyMatcher{
								line: 382, col: 50, offset: 17612,
							},
						},
					},
				},
			},
		},
		{
			name: "UnbalancedQuotePunctuation",
			pos:  position{line: 387, col: 1, offset: 17857},
			expr: &choiceExpr{
				pos: position{line: 387, col: 31, offset: 17887},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 387, col: 31, offset: 17887},
						val:        "*",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 387, col: 37, offset: 17893},
						val:        "_",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 387, col: 43, offset: 17899},
						val:        "`",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "Passthrough",
			pos:  position{line: 392, col: 1, offset: 18011},
			expr: &choiceExpr{
				pos: position{line: 392, col: 16, offset: 18026},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 392, col: 16, offset: 18026},
						name: "TriplePlusPassthrough",
					},
					&ruleRefExpr{
						pos:  position{line: 392, col: 40, offset: 18050},
						name: "SinglePlusPassthrough",
					},
					&ruleRefExpr{
						pos:  position{line: 392, col: 64, offset: 18074},
						name: "PassthroughMacro",
					},
				},
			},
		},
		{
			name: "SinglePlusPassthrough",
			pos:  position{line: 394, col: 1, offset: 18092},
			expr: &actionExpr{
				pos: position{line: 394, col: 26, offset: 18117},
				run: (*parser).callonSinglePlusPassthrough1,
				expr: &seqExpr{
					pos: position{line: 394, col: 26, offset: 18117},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 394, col: 26, offset: 18117},
							val:        "+",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 394, col: 30, offset: 18121},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 394, col: 38, offset: 18129},
								expr: &seqExpr{
									pos: position{line: 394, col: 39, offset: 18130},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 394, col: 39, offset: 18130},
											expr: &ruleRefExpr{
												pos:  position{line: 394, col: 40, offset: 18131},
												name: "NEWLINE",
											},
										},
										&notExpr{
											pos: position{line: 394, col: 48, offset: 18139},
											expr: &litMatcher{
												pos:        position{line: 394, col: 49, offset: 18140},
												val:        "+",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 394, col: 53, offset: 18144,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 394, col: 57, offset: 18148},
							val:        "+",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TriplePlusPassthrough",
			pos:  position{line: 398, col: 1, offset: 18243},
			expr: &actionExpr{
				pos: position{line: 398, col: 26, offset: 18268},
				run: (*parser).callonTriplePlusPassthrough1,
				expr: &seqExpr{
					pos: position{line: 398, col: 26, offset: 18268},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 398, col: 26, offset: 18268},
							val:        "+++",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 398, col: 32, offset: 18274},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 398, col: 40, offset: 18282},
								expr: &seqExpr{
									pos: position{line: 398, col: 41, offset: 18283},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 398, col: 41, offset: 18283},
											expr: &litMatcher{
												pos:        position{line: 398, col: 42, offset: 18284},
												val:        "+++",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 398, col: 48, offset: 18290,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 398, col: 52, offset: 18294},
							val:        "+++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "PassthroughMacro",
			pos:  position{line: 402, col: 1, offset: 18391},
			expr: &choiceExpr{
				pos: position{line: 402, col: 21, offset: 18411},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 402, col: 21, offset: 18411},
						run: (*parser).callonPassthroughMacro2,
						expr: &seqExpr{
							pos: position{line: 402, col: 21, offset: 18411},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 402, col: 21, offset: 18411},
									val:        "pass:[",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 402, col: 30, offset: 18420},
									label: "content",
									expr: &zeroOrMoreExpr{
										pos: position{line: 402, col: 38, offset: 18428},
										expr: &ruleRefExpr{
											pos:  position{line: 402, col: 39, offset: 18429},
											name: "PassthroughMacroCharacter",
										},
									},
								},
								&litMatcher{
									pos:        position{line: 402, col: 67, offset: 18457},
									val:        "]",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 404, col: 5, offset: 18548},
						run: (*parser).callonPassthroughMacro9,
						expr: &seqExpr{
							pos: position{line: 404, col: 5, offset: 18548},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 404, col: 5, offset: 18548},
									val:        "pass:q[",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 404, col: 15, offset: 18558},
									label: "content",
									expr: &zeroOrMoreExpr{
										pos: position{line: 404, col: 23, offset: 18566},
										expr: &choiceExpr{
											pos: position{line: 404, col: 24, offset: 18567},
											alternatives: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 404, col: 24, offset: 18567},
													name: "QuotedText",
												},
												&ruleRefExpr{
													pos:  position{line: 404, col: 37, offset: 18580},
													name: "PassthroughMacroCharacter",
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 404, col: 65, offset: 18608},
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
			pos:  position{line: 408, col: 1, offset: 18698},
			expr: &seqExpr{
				pos: position{line: 408, col: 31, offset: 18728},
				exprs: []interface{}{
					&notExpr{
						pos: position{line: 408, col: 31, offset: 18728},
						expr: &litMatcher{
							pos:        position{line: 408, col: 32, offset: 18729},
							val:        "]",
							ignoreCase: false,
						},
					},
					&anyMatcher{
						line: 408, col: 36, offset: 18733,
					},
				},
			},
		},
		{
			name: "CrossReference",
			pos:  position{line: 413, col: 1, offset: 18849},
			expr: &actionExpr{
				pos: position{line: 413, col: 19, offset: 18867},
				run: (*parser).callonCrossReference1,
				expr: &seqExpr{
					pos: position{line: 413, col: 19, offset: 18867},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 413, col: 19, offset: 18867},
							val:        "<<",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 413, col: 24, offset: 18872},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 413, col: 28, offset: 18876},
								name: "ID",
							},
						},
						&litMatcher{
							pos:        position{line: 413, col: 32, offset: 18880},
							val:        ">>",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Link",
			pos:  position{line: 420, col: 1, offset: 19039},
			expr: &choiceExpr{
				pos: position{line: 420, col: 9, offset: 19047},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 420, col: 9, offset: 19047},
						name: "RelativeLink",
					},
					&ruleRefExpr{
						pos:  position{line: 420, col: 24, offset: 19062},
						name: "ExternalLink",
					},
				},
			},
		},
		{
			name: "ExternalLink",
			pos:  position{line: 422, col: 1, offset: 19077},
			expr: &actionExpr{
				pos: position{line: 422, col: 17, offset: 19093},
				run: (*parser).callonExternalLink1,
				expr: &seqExpr{
					pos: position{line: 422, col: 17, offset: 19093},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 422, col: 17, offset: 19093},
							label: "url",
							expr: &seqExpr{
								pos: position{line: 422, col: 22, offset: 19098},
								exprs: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 422, col: 22, offset: 19098},
										name: "URL_SCHEME",
									},
									&ruleRefExpr{
										pos:  position{line: 422, col: 33, offset: 19109},
										name: "URL",
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 422, col: 38, offset: 19114},
							label: "text",
							expr: &zeroOrOneExpr{
								pos: position{line: 422, col: 43, offset: 19119},
								expr: &seqExpr{
									pos: position{line: 422, col: 44, offset: 19120},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 422, col: 44, offset: 19120},
											val:        "[",
											ignoreCase: false,
										},
										&zeroOrMoreExpr{
											pos: position{line: 422, col: 48, offset: 19124},
											expr: &ruleRefExpr{
												pos:  position{line: 422, col: 49, offset: 19125},
												name: "URL_TEXT",
											},
										},
										&litMatcher{
											pos:        position{line: 422, col: 60, offset: 19136},
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
			pos:  position{line: 429, col: 1, offset: 19297},
			expr: &actionExpr{
				pos: position{line: 429, col: 17, offset: 19313},
				run: (*parser).callonRelativeLink1,
				expr: &seqExpr{
					pos: position{line: 429, col: 17, offset: 19313},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 429, col: 17, offset: 19313},
							val:        "link:",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 429, col: 25, offset: 19321},
							label: "url",
							expr: &seqExpr{
								pos: position{line: 429, col: 30, offset: 19326},
								exprs: []interface{}{
									&zeroOrOneExpr{
										pos: position{line: 429, col: 30, offset: 19326},
										expr: &ruleRefExpr{
											pos:  position{line: 429, col: 30, offset: 19326},
											name: "URL_SCHEME",
										},
									},
									&ruleRefExpr{
										pos:  position{line: 429, col: 42, offset: 19338},
										name: "URL",
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 429, col: 47, offset: 19343},
							label: "text",
							expr: &seqExpr{
								pos: position{line: 429, col: 53, offset: 19349},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 429, col: 53, offset: 19349},
										val:        "[",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 429, col: 57, offset: 19353},
										expr: &ruleRefExpr{
											pos:  position{line: 429, col: 58, offset: 19354},
											name: "URL_TEXT",
										},
									},
									&litMatcher{
										pos:        position{line: 429, col: 69, offset: 19365},
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
			pos:  position{line: 439, col: 1, offset: 19627},
			expr: &actionExpr{
				pos: position{line: 439, col: 15, offset: 19641},
				run: (*parser).callonBlockImage1,
				expr: &seqExpr{
					pos: position{line: 439, col: 15, offset: 19641},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 439, col: 15, offset: 19641},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 439, col: 26, offset: 19652},
								expr: &ruleRefExpr{
									pos:  position{line: 439, col: 27, offset: 19653},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 439, col: 46, offset: 19672},
							label: "image",
							expr: &ruleRefExpr{
								pos:  position{line: 439, col: 52, offset: 19678},
								name: "BlockImageMacro",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 439, col: 69, offset: 19695},
							expr: &ruleRefExpr{
								pos:  position{line: 439, col: 69, offset: 19695},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 439, col: 73, offset: 19699},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "BlockImageMacro",
			pos:  position{line: 444, col: 1, offset: 19860},
			expr: &actionExpr{
				pos: position{line: 444, col: 20, offset: 19879},
				run: (*parser).callonBlockImageMacro1,
				expr: &seqExpr{
					pos: position{line: 444, col: 20, offset: 19879},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 444, col: 20, offset: 19879},
							val:        "image::",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 444, col: 30, offset: 19889},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 444, col: 36, offset: 19895},
								name: "URL",
							},
						},
						&litMatcher{
							pos:        position{line: 444, col: 41, offset: 19900},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 444, col: 45, offset: 19904},
							label: "attributes",
							expr: &zeroOrOneExpr{
								pos: position{line: 444, col: 57, offset: 19916},
								expr: &ruleRefExpr{
									pos:  position{line: 444, col: 57, offset: 19916},
									name: "URL_TEXT",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 444, col: 68, offset: 19927},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "InlineImage",
			pos:  position{line: 448, col: 1, offset: 19994},
			expr: &actionExpr{
				pos: position{line: 448, col: 16, offset: 20009},
				run: (*parser).callonInlineImage1,
				expr: &labeledExpr{
					pos:   position{line: 448, col: 16, offset: 20009},
					label: "image",
					expr: &ruleRefExpr{
						pos:  position{line: 448, col: 22, offset: 20015},
						name: "InlineImageMacro",
					},
				},
			},
		},
		{
			name: "InlineImageMacro",
			pos:  position{line: 453, col: 1, offset: 20162},
			expr: &actionExpr{
				pos: position{line: 453, col: 21, offset: 20182},
				run: (*parser).callonInlineImageMacro1,
				expr: &seqExpr{
					pos: position{line: 453, col: 21, offset: 20182},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 453, col: 21, offset: 20182},
							val:        "image:",
							ignoreCase: false,
						},
						&notExpr{
							pos: position{line: 453, col: 30, offset: 20191},
							expr: &litMatcher{
								pos:        position{line: 453, col: 31, offset: 20192},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 453, col: 35, offset: 20196},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 453, col: 41, offset: 20202},
								name: "URL",
							},
						},
						&litMatcher{
							pos:        position{line: 453, col: 46, offset: 20207},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 453, col: 50, offset: 20211},
							label: "attributes",
							expr: &zeroOrOneExpr{
								pos: position{line: 453, col: 62, offset: 20223},
								expr: &ruleRefExpr{
									pos:  position{line: 453, col: 62, offset: 20223},
									name: "URL_TEXT",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 453, col: 73, offset: 20234},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "DelimitedBlock",
			pos:  position{line: 460, col: 1, offset: 20564},
			expr: &choiceExpr{
				pos: position{line: 460, col: 19, offset: 20582},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 460, col: 19, offset: 20582},
						name: "FencedBlock",
					},
					&ruleRefExpr{
						pos:  position{line: 460, col: 33, offset: 20596},
						name: "ListingBlock",
					},
					&ruleRefExpr{
						pos:  position{line: 460, col: 48, offset: 20611},
						name: "ExampleBlock",
					},
				},
			},
		},
		{
			name: "BlockDelimiter",
			pos:  position{line: 462, col: 1, offset: 20625},
			expr: &choiceExpr{
				pos: position{line: 462, col: 19, offset: 20643},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 462, col: 19, offset: 20643},
						name: "LiteralBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 462, col: 43, offset: 20667},
						name: "FencedBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 462, col: 66, offset: 20690},
						name: "ListingBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 462, col: 90, offset: 20714},
						name: "ExampleBlockDelimiter",
					},
				},
			},
		},
		{
			name: "FencedBlockDelimiter",
			pos:  position{line: 464, col: 1, offset: 20737},
			expr: &litMatcher{
				pos:        position{line: 464, col: 25, offset: 20761},
				val:        "```",
				ignoreCase: false,
			},
		},
		{
			name: "FencedBlock",
			pos:  position{line: 466, col: 1, offset: 20768},
			expr: &actionExpr{
				pos: position{line: 466, col: 16, offset: 20783},
				run: (*parser).callonFencedBlock1,
				expr: &seqExpr{
					pos: position{line: 466, col: 16, offset: 20783},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 466, col: 16, offset: 20783},
							name: "FencedBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 466, col: 37, offset: 20804},
							expr: &ruleRefExpr{
								pos:  position{line: 466, col: 37, offset: 20804},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 466, col: 41, offset: 20808},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 466, col: 49, offset: 20816},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 466, col: 57, offset: 20824},
								expr: &seqExpr{
									pos: position{line: 466, col: 58, offset: 20825},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 466, col: 58, offset: 20825},
											expr: &ruleRefExpr{
												pos:  position{line: 466, col: 59, offset: 20826},
												name: "FencedBlockDelimiter",
											},
										},
										&anyMatcher{
											line: 466, col: 80, offset: 20847,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 466, col: 84, offset: 20851},
							name: "FencedBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 466, col: 105, offset: 20872},
							expr: &ruleRefExpr{
								pos:  position{line: 466, col: 105, offset: 20872},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 466, col: 109, offset: 20876},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ListingBlockDelimiter",
			pos:  position{line: 470, col: 1, offset: 20969},
			expr: &litMatcher{
				pos:        position{line: 470, col: 26, offset: 20994},
				val:        "----",
				ignoreCase: false,
			},
		},
		{
			name: "ListingBlock",
			pos:  position{line: 472, col: 1, offset: 21002},
			expr: &actionExpr{
				pos: position{line: 472, col: 17, offset: 21018},
				run: (*parser).callonListingBlock1,
				expr: &seqExpr{
					pos: position{line: 472, col: 17, offset: 21018},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 472, col: 17, offset: 21018},
							name: "ListingBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 472, col: 39, offset: 21040},
							expr: &ruleRefExpr{
								pos:  position{line: 472, col: 39, offset: 21040},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 472, col: 43, offset: 21044},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 472, col: 51, offset: 21052},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 472, col: 59, offset: 21060},
								expr: &seqExpr{
									pos: position{line: 472, col: 60, offset: 21061},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 472, col: 60, offset: 21061},
											expr: &ruleRefExpr{
												pos:  position{line: 472, col: 61, offset: 21062},
												name: "ListingBlockDelimiter",
											},
										},
										&anyMatcher{
											line: 472, col: 83, offset: 21084,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 472, col: 87, offset: 21088},
							name: "ListingBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 472, col: 109, offset: 21110},
							expr: &ruleRefExpr{
								pos:  position{line: 472, col: 109, offset: 21110},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 472, col: 113, offset: 21114},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ExampleBlockDelimiter",
			pos:  position{line: 476, col: 1, offset: 21208},
			expr: &litMatcher{
				pos:        position{line: 476, col: 26, offset: 21233},
				val:        "====",
				ignoreCase: false,
			},
		},
		{
			name: "ExampleBlock",
			pos:  position{line: 478, col: 1, offset: 21241},
			expr: &actionExpr{
				pos: position{line: 478, col: 17, offset: 21257},
				run: (*parser).callonExampleBlock1,
				expr: &seqExpr{
					pos: position{line: 478, col: 17, offset: 21257},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 478, col: 17, offset: 21257},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 478, col: 28, offset: 21268},
								expr: &ruleRefExpr{
									pos:  position{line: 478, col: 29, offset: 21269},
									name: "ElementAttribute",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 478, col: 48, offset: 21288},
							name: "ExampleBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 478, col: 70, offset: 21310},
							expr: &ruleRefExpr{
								pos:  position{line: 478, col: 70, offset: 21310},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 478, col: 74, offset: 21314},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 478, col: 82, offset: 21322},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 478, col: 90, offset: 21330},
								expr: &choiceExpr{
									pos: position{line: 478, col: 91, offset: 21331},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 478, col: 91, offset: 21331},
											name: "List",
										},
										&ruleRefExpr{
											pos:  position{line: 478, col: 98, offset: 21338},
											name: "Paragraph",
										},
										&ruleRefExpr{
											pos:  position{line: 478, col: 110, offset: 21350},
											name: "BlankLine",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 478, col: 123, offset: 21363},
							name: "ExampleBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 478, col: 145, offset: 21385},
							expr: &ruleRefExpr{
								pos:  position{line: 478, col: 145, offset: 21385},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 478, col: 149, offset: 21389},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "LiteralBlock",
			pos:  position{line: 485, col: 1, offset: 21773},
			expr: &choiceExpr{
				pos: position{line: 485, col: 17, offset: 21789},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 485, col: 17, offset: 21789},
						name: "ParagraphWithSpaces",
					},
					&ruleRefExpr{
						pos:  position{line: 485, col: 39, offset: 21811},
						name: "ParagraphWithLiteralBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 485, col: 76, offset: 21848},
						name: "ParagraphWithLiteralAttribute",
					},
				},
			},
		},
		{
			name: "ParagraphWithSpaces",
			pos:  position{line: 488, col: 1, offset: 21943},
			expr: &actionExpr{
				pos: position{line: 488, col: 24, offset: 21966},
				run: (*parser).callonParagraphWithSpaces1,
				expr: &seqExpr{
					pos: position{line: 488, col: 24, offset: 21966},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 488, col: 24, offset: 21966},
							label: "spaces",
							expr: &oneOrMoreExpr{
								pos: position{line: 488, col: 32, offset: 21974},
								expr: &ruleRefExpr{
									pos:  position{line: 488, col: 32, offset: 21974},
									name: "WS",
								},
							},
						},
						&notExpr{
							pos: position{line: 488, col: 37, offset: 21979},
							expr: &ruleRefExpr{
								pos:  position{line: 488, col: 38, offset: 21980},
								name: "NEWLINE",
							},
						},
						&labeledExpr{
							pos:   position{line: 488, col: 46, offset: 21988},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 488, col: 55, offset: 21997},
								name: "LiteralBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 488, col: 76, offset: 22018},
							name: "EndOfLiteralBlock",
						},
					},
				},
			},
		},
		{
			name: "LiteralBlockContent",
			pos:  position{line: 493, col: 1, offset: 22199},
			expr: &actionExpr{
				pos: position{line: 493, col: 24, offset: 22222},
				run: (*parser).callonLiteralBlockContent1,
				expr: &labeledExpr{
					pos:   position{line: 493, col: 24, offset: 22222},
					label: "content",
					expr: &oneOrMoreExpr{
						pos: position{line: 493, col: 32, offset: 22230},
						expr: &seqExpr{
							pos: position{line: 493, col: 33, offset: 22231},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 493, col: 33, offset: 22231},
									expr: &seqExpr{
										pos: position{line: 493, col: 35, offset: 22233},
										exprs: []interface{}{
											&ruleRefExpr{
												pos:  position{line: 493, col: 35, offset: 22233},
												name: "NEWLINE",
											},
											&ruleRefExpr{
												pos:  position{line: 493, col: 43, offset: 22241},
												name: "BlankLine",
											},
										},
									},
								},
								&anyMatcher{
									line: 493, col: 54, offset: 22252,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "EndOfLiteralBlock",
			pos:  position{line: 498, col: 1, offset: 22337},
			expr: &choiceExpr{
				pos: position{line: 498, col: 22, offset: 22358},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 498, col: 22, offset: 22358},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 498, col: 22, offset: 22358},
								name: "NEWLINE",
							},
							&ruleRefExpr{
								pos:  position{line: 498, col: 30, offset: 22366},
								name: "BlankLine",
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 498, col: 42, offset: 22378},
						name: "NEWLINE",
					},
					&ruleRefExpr{
						pos:  position{line: 498, col: 52, offset: 22388},
						name: "EOF",
					},
				},
			},
		},
		{
			name: "ParagraphWithLiteralBlockDelimiter",
			pos:  position{line: 501, col: 1, offset: 22448},
			expr: &actionExpr{
				pos: position{line: 501, col: 39, offset: 22486},
				run: (*parser).callonParagraphWithLiteralBlockDelimiter1,
				expr: &seqExpr{
					pos: position{line: 501, col: 39, offset: 22486},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 501, col: 39, offset: 22486},
							name: "LiteralBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 501, col: 61, offset: 22508},
							expr: &ruleRefExpr{
								pos:  position{line: 501, col: 61, offset: 22508},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 501, col: 65, offset: 22512},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 501, col: 73, offset: 22520},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 501, col: 81, offset: 22528},
								expr: &seqExpr{
									pos: position{line: 501, col: 82, offset: 22529},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 501, col: 82, offset: 22529},
											expr: &ruleRefExpr{
												pos:  position{line: 501, col: 83, offset: 22530},
												name: "LiteralBlockDelimiter",
											},
										},
										&anyMatcher{
											line: 501, col: 105, offset: 22552,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 501, col: 109, offset: 22556},
							name: "LiteralBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 501, col: 131, offset: 22578},
							expr: &ruleRefExpr{
								pos:  position{line: 501, col: 131, offset: 22578},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 501, col: 135, offset: 22582},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "LiteralBlockDelimiter",
			pos:  position{line: 505, col: 1, offset: 22666},
			expr: &litMatcher{
				pos:        position{line: 505, col: 26, offset: 22691},
				val:        "....",
				ignoreCase: false,
			},
		},
		{
			name: "ParagraphWithLiteralAttribute",
			pos:  position{line: 508, col: 1, offset: 22753},
			expr: &actionExpr{
				pos: position{line: 508, col: 34, offset: 22786},
				run: (*parser).callonParagraphWithLiteralAttribute1,
				expr: &seqExpr{
					pos: position{line: 508, col: 34, offset: 22786},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 508, col: 34, offset: 22786},
							val:        "[literal]",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 508, col: 46, offset: 22798},
							expr: &ruleRefExpr{
								pos:  position{line: 508, col: 46, offset: 22798},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 508, col: 50, offset: 22802},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 508, col: 58, offset: 22810},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 508, col: 67, offset: 22819},
								name: "LiteralBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 508, col: 88, offset: 22840},
							name: "EndOfLiteralBlock",
						},
					},
				},
			},
		},
		{
			name: "ElementAttribute",
			pos:  position{line: 515, col: 1, offset: 23052},
			expr: &choiceExpr{
				pos: position{line: 515, col: 21, offset: 23072},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 515, col: 21, offset: 23072},
						name: "ElementLink",
					},
					&ruleRefExpr{
						pos:  position{line: 515, col: 35, offset: 23086},
						name: "ElementID",
					},
					&ruleRefExpr{
						pos:  position{line: 515, col: 47, offset: 23098},
						name: "ElementTitle",
					},
					&ruleRefExpr{
						pos:  position{line: 515, col: 62, offset: 23113},
						name: "InvalidElementAttribute",
					},
				},
			},
		},
		{
			name: "ElementLink",
			pos:  position{line: 518, col: 1, offset: 23193},
			expr: &actionExpr{
				pos: position{line: 518, col: 16, offset: 23208},
				run: (*parser).callonElementLink1,
				expr: &seqExpr{
					pos: position{line: 518, col: 16, offset: 23208},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 518, col: 16, offset: 23208},
							val:        "[link=",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 518, col: 25, offset: 23217},
							expr: &ruleRefExpr{
								pos:  position{line: 518, col: 25, offset: 23217},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 518, col: 29, offset: 23221},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 518, col: 34, offset: 23226},
								name: "URL",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 518, col: 38, offset: 23230},
							expr: &ruleRefExpr{
								pos:  position{line: 518, col: 38, offset: 23230},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 518, col: 42, offset: 23234},
							val:        "]",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 518, col: 46, offset: 23238},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ElementID",
			pos:  position{line: 522, col: 1, offset: 23294},
			expr: &choiceExpr{
				pos: position{line: 522, col: 14, offset: 23307},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 522, col: 14, offset: 23307},
						run: (*parser).callonElementID2,
						expr: &seqExpr{
							pos: position{line: 522, col: 14, offset: 23307},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 522, col: 14, offset: 23307},
									label: "id",
									expr: &ruleRefExpr{
										pos:  position{line: 522, col: 18, offset: 23311},
										name: "InlineElementID",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 522, col: 35, offset: 23328},
									name: "EOL",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 524, col: 5, offset: 23357},
						run: (*parser).callonElementID7,
						expr: &seqExpr{
							pos: position{line: 524, col: 5, offset: 23357},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 524, col: 5, offset: 23357},
									val:        "[#",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 524, col: 10, offset: 23362},
									label: "id",
									expr: &ruleRefExpr{
										pos:  position{line: 524, col: 14, offset: 23366},
										name: "ID",
									},
								},
								&litMatcher{
									pos:        position{line: 524, col: 18, offset: 23370},
									val:        "]",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 524, col: 22, offset: 23374},
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
			pos:  position{line: 528, col: 1, offset: 23426},
			expr: &actionExpr{
				pos: position{line: 528, col: 20, offset: 23445},
				run: (*parser).callonInlineElementID1,
				expr: &seqExpr{
					pos: position{line: 528, col: 20, offset: 23445},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 528, col: 20, offset: 23445},
							val:        "[[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 528, col: 25, offset: 23450},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 528, col: 29, offset: 23454},
								name: "ID",
							},
						},
						&litMatcher{
							pos:        position{line: 528, col: 33, offset: 23458},
							val:        "]]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ElementTitle",
			pos:  position{line: 534, col: 1, offset: 23653},
			expr: &actionExpr{
				pos: position{line: 534, col: 17, offset: 23669},
				run: (*parser).callonElementTitle1,
				expr: &seqExpr{
					pos: position{line: 534, col: 17, offset: 23669},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 534, col: 17, offset: 23669},
							val:        ".",
							ignoreCase: false,
						},
						&notExpr{
							pos: position{line: 534, col: 21, offset: 23673},
							expr: &litMatcher{
								pos:        position{line: 534, col: 22, offset: 23674},
								val:        ".",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 534, col: 26, offset: 23678},
							expr: &ruleRefExpr{
								pos:  position{line: 534, col: 27, offset: 23679},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 534, col: 30, offset: 23682},
							label: "title",
							expr: &oneOrMoreExpr{
								pos: position{line: 534, col: 36, offset: 23688},
								expr: &seqExpr{
									pos: position{line: 534, col: 37, offset: 23689},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 534, col: 37, offset: 23689},
											expr: &ruleRefExpr{
												pos:  position{line: 534, col: 38, offset: 23690},
												name: "NEWLINE",
											},
										},
										&anyMatcher{
											line: 534, col: 46, offset: 23698,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 534, col: 50, offset: 23702},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "InvalidElementAttribute",
			pos:  position{line: 538, col: 1, offset: 23767},
			expr: &actionExpr{
				pos: position{line: 538, col: 28, offset: 23794},
				run: (*parser).callonInvalidElementAttribute1,
				expr: &seqExpr{
					pos: position{line: 538, col: 28, offset: 23794},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 538, col: 28, offset: 23794},
							val:        "[",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 538, col: 32, offset: 23798},
							expr: &ruleRefExpr{
								pos:  position{line: 538, col: 32, offset: 23798},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 538, col: 36, offset: 23802},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 538, col: 44, offset: 23810},
								expr: &seqExpr{
									pos: position{line: 538, col: 45, offset: 23811},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 538, col: 45, offset: 23811},
											expr: &litMatcher{
												pos:        position{line: 538, col: 46, offset: 23812},
												val:        "]",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 538, col: 50, offset: 23816,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 538, col: 54, offset: 23820},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "BlankLine",
			pos:  position{line: 545, col: 1, offset: 23986},
			expr: &actionExpr{
				pos: position{line: 545, col: 14, offset: 23999},
				run: (*parser).callonBlankLine1,
				expr: &seqExpr{
					pos: position{line: 545, col: 14, offset: 23999},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 545, col: 14, offset: 23999},
							expr: &ruleRefExpr{
								pos:  position{line: 545, col: 15, offset: 24000},
								name: "EOF",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 545, col: 19, offset: 24004},
							expr: &ruleRefExpr{
								pos:  position{line: 545, col: 19, offset: 24004},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 545, col: 23, offset: 24008},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "Characters",
			pos:  position{line: 552, col: 1, offset: 24155},
			expr: &actionExpr{
				pos: position{line: 552, col: 15, offset: 24169},
				run: (*parser).callonCharacters1,
				expr: &oneOrMoreExpr{
					pos: position{line: 552, col: 15, offset: 24169},
					expr: &seqExpr{
						pos: position{line: 552, col: 16, offset: 24170},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 552, col: 16, offset: 24170},
								expr: &ruleRefExpr{
									pos:  position{line: 552, col: 17, offset: 24171},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 552, col: 25, offset: 24179},
								expr: &ruleRefExpr{
									pos:  position{line: 552, col: 26, offset: 24180},
									name: "WS",
								},
							},
							&anyMatcher{
								line: 552, col: 29, offset: 24183,
							},
						},
					},
				},
			},
		},
		{
			name: "URL",
			pos:  position{line: 556, col: 1, offset: 24223},
			expr: &actionExpr{
				pos: position{line: 556, col: 8, offset: 24230},
				run: (*parser).callonURL1,
				expr: &oneOrMoreExpr{
					pos: position{line: 556, col: 8, offset: 24230},
					expr: &seqExpr{
						pos: position{line: 556, col: 9, offset: 24231},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 556, col: 9, offset: 24231},
								expr: &ruleRefExpr{
									pos:  position{line: 556, col: 10, offset: 24232},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 556, col: 18, offset: 24240},
								expr: &ruleRefExpr{
									pos:  position{line: 556, col: 19, offset: 24241},
									name: "WS",
								},
							},
							&notExpr{
								pos: position{line: 556, col: 22, offset: 24244},
								expr: &litMatcher{
									pos:        position{line: 556, col: 23, offset: 24245},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 556, col: 27, offset: 24249},
								expr: &litMatcher{
									pos:        position{line: 556, col: 28, offset: 24250},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 556, col: 32, offset: 24254,
							},
						},
					},
				},
			},
		},
		{
			name: "ID",
			pos:  position{line: 560, col: 1, offset: 24294},
			expr: &actionExpr{
				pos: position{line: 560, col: 7, offset: 24300},
				run: (*parser).callonID1,
				expr: &oneOrMoreExpr{
					pos: position{line: 560, col: 7, offset: 24300},
					expr: &seqExpr{
						pos: position{line: 560, col: 8, offset: 24301},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 560, col: 8, offset: 24301},
								expr: &ruleRefExpr{
									pos:  position{line: 560, col: 9, offset: 24302},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 560, col: 17, offset: 24310},
								expr: &ruleRefExpr{
									pos:  position{line: 560, col: 18, offset: 24311},
									name: "WS",
								},
							},
							&notExpr{
								pos: position{line: 560, col: 21, offset: 24314},
								expr: &litMatcher{
									pos:        position{line: 560, col: 22, offset: 24315},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 560, col: 26, offset: 24319},
								expr: &litMatcher{
									pos:        position{line: 560, col: 27, offset: 24320},
									val:        "]",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 560, col: 31, offset: 24324},
								expr: &litMatcher{
									pos:        position{line: 560, col: 32, offset: 24325},
									val:        "<<",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 560, col: 37, offset: 24330},
								expr: &litMatcher{
									pos:        position{line: 560, col: 38, offset: 24331},
									val:        ">>",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 560, col: 42, offset: 24335,
							},
						},
					},
				},
			},
		},
		{
			name: "URL_TEXT",
			pos:  position{line: 564, col: 1, offset: 24375},
			expr: &actionExpr{
				pos: position{line: 564, col: 13, offset: 24387},
				run: (*parser).callonURL_TEXT1,
				expr: &oneOrMoreExpr{
					pos: position{line: 564, col: 13, offset: 24387},
					expr: &seqExpr{
						pos: position{line: 564, col: 14, offset: 24388},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 564, col: 14, offset: 24388},
								expr: &ruleRefExpr{
									pos:  position{line: 564, col: 15, offset: 24389},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 564, col: 23, offset: 24397},
								expr: &litMatcher{
									pos:        position{line: 564, col: 24, offset: 24398},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 564, col: 28, offset: 24402},
								expr: &litMatcher{
									pos:        position{line: 564, col: 29, offset: 24403},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 564, col: 33, offset: 24407,
							},
						},
					},
				},
			},
		},
		{
			name: "URL_SCHEME",
			pos:  position{line: 568, col: 1, offset: 24447},
			expr: &choiceExpr{
				pos: position{line: 568, col: 15, offset: 24461},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 568, col: 15, offset: 24461},
						val:        "http://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 568, col: 27, offset: 24473},
						val:        "https://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 568, col: 40, offset: 24486},
						val:        "ftp://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 568, col: 51, offset: 24497},
						val:        "irc://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 568, col: 62, offset: 24508},
						val:        "mailto:",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "DIGIT",
			pos:  position{line: 570, col: 1, offset: 24519},
			expr: &charClassMatcher{
				pos:        position{line: 570, col: 10, offset: 24528},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NEWLINE",
			pos:  position{line: 572, col: 1, offset: 24535},
			expr: &choiceExpr{
				pos: position{line: 572, col: 12, offset: 24546},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 572, col: 12, offset: 24546},
						val:        "\r\n",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 572, col: 21, offset: 24555},
						val:        "\r",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 572, col: 28, offset: 24562},
						val:        "\n",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "WS",
			pos:  position{line: 574, col: 1, offset: 24568},
			expr: &choiceExpr{
				pos: position{line: 574, col: 7, offset: 24574},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 574, col: 7, offset: 24574},
						val:        " ",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 574, col: 13, offset: 24580},
						run: (*parser).callonWS3,
						expr: &litMatcher{
							pos:        position{line: 574, col: 13, offset: 24580},
							val:        "\t",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 578, col: 1, offset: 24625},
			expr: &notExpr{
				pos: position{line: 578, col: 8, offset: 24632},
				expr: &anyMatcher{
					line: 578, col: 9, offset: 24633,
				},
			},
		},
		{
			name: "EOL",
			pos:  position{line: 580, col: 1, offset: 24636},
			expr: &choiceExpr{
				pos: position{line: 580, col: 8, offset: 24643},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 580, col: 8, offset: 24643},
						name: "NEWLINE",
					},
					&ruleRefExpr{
						pos:  position{line: 580, col: 18, offset: 24653},
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

func (c *current) onAdmonition2(attributes, t, content interface{}) (interface{}, error) {
	// paragraph style
	return types.NewAdmonition(t.(types.AdmonitionKind), content, attributes.([]interface{}))
}

func (p *parser) callonAdmonition2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAdmonition2(stack["attributes"], stack["t"], stack["content"])
}

func (c *current) onAdmonition18(attributes, t, otherAttributes, content interface{}) (interface{}, error) {
	// block style
	return types.NewAdmonition(t.(types.AdmonitionKind), content, append(attributes.([]interface{}), otherAttributes.([]interface{})...))
}

func (p *parser) callonAdmonition18() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAdmonition18(stack["attributes"], stack["t"], stack["otherAttributes"], stack["content"])
}

func (c *current) onAdmonitionParagraph1(lines interface{}) (interface{}, error) {
	return types.NewAdmonitionParagraph(lines.([]interface{}))
}

func (p *parser) callonAdmonitionParagraph1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAdmonitionParagraph1(stack["lines"])
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
	return types.NewDelimitedBlock(types.FencedBlock, content.([]interface{}), nil)
}

func (p *parser) callonFencedBlock1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFencedBlock1(stack["content"])
}

func (c *current) onListingBlock1(content interface{}) (interface{}, error) {
	return types.NewDelimitedBlock(types.ListingBlock, content.([]interface{}), nil)
}

func (p *parser) callonListingBlock1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onListingBlock1(stack["content"])
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
