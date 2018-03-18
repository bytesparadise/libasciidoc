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
			pos:  position{line: 60, col: 1, offset: 2208},
			expr: &choiceExpr{
				pos: position{line: 60, col: 20, offset: 2227},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 60, col: 20, offset: 2227},
						name: "DocumentAuthorsInlineForm",
					},
					&ruleRefExpr{
						pos:  position{line: 60, col: 48, offset: 2255},
						name: "DocumentAuthorsAttributeForm",
					},
				},
			},
		},
		{
			name: "DocumentAuthorsInlineForm",
			pos:  position{line: 62, col: 1, offset: 2285},
			expr: &actionExpr{
				pos: position{line: 62, col: 30, offset: 2314},
				run: (*parser).callonDocumentAuthorsInlineForm1,
				expr: &seqExpr{
					pos: position{line: 62, col: 30, offset: 2314},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 62, col: 30, offset: 2314},
							expr: &ruleRefExpr{
								pos:  position{line: 62, col: 30, offset: 2314},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 62, col: 34, offset: 2318},
							expr: &litMatcher{
								pos:        position{line: 62, col: 35, offset: 2319},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 62, col: 39, offset: 2323},
							label: "authors",
							expr: &oneOrMoreExpr{
								pos: position{line: 62, col: 48, offset: 2332},
								expr: &ruleRefExpr{
									pos:  position{line: 62, col: 48, offset: 2332},
									name: "DocumentAuthor",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 62, col: 65, offset: 2349},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthorsAttributeForm",
			pos:  position{line: 66, col: 1, offset: 2419},
			expr: &actionExpr{
				pos: position{line: 66, col: 33, offset: 2451},
				run: (*parser).callonDocumentAuthorsAttributeForm1,
				expr: &seqExpr{
					pos: position{line: 66, col: 33, offset: 2451},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 66, col: 33, offset: 2451},
							expr: &ruleRefExpr{
								pos:  position{line: 66, col: 33, offset: 2451},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 66, col: 37, offset: 2455},
							val:        ":author:",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 66, col: 48, offset: 2466},
							label: "author",
							expr: &ruleRefExpr{
								pos:  position{line: 66, col: 56, offset: 2474},
								name: "DocumentAuthor",
							},
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthor",
			pos:  position{line: 70, col: 1, offset: 2565},
			expr: &actionExpr{
				pos: position{line: 70, col: 19, offset: 2583},
				run: (*parser).callonDocumentAuthor1,
				expr: &seqExpr{
					pos: position{line: 70, col: 19, offset: 2583},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 70, col: 19, offset: 2583},
							expr: &ruleRefExpr{
								pos:  position{line: 70, col: 19, offset: 2583},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 70, col: 23, offset: 2587},
							label: "namePart1",
							expr: &ruleRefExpr{
								pos:  position{line: 70, col: 34, offset: 2598},
								name: "DocumentAuthorNamePart",
							},
						},
						&labeledExpr{
							pos:   position{line: 70, col: 58, offset: 2622},
							label: "namePart2",
							expr: &zeroOrOneExpr{
								pos: position{line: 70, col: 68, offset: 2632},
								expr: &ruleRefExpr{
									pos:  position{line: 70, col: 69, offset: 2633},
									name: "DocumentAuthorNamePart",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 70, col: 94, offset: 2658},
							label: "namePart3",
							expr: &zeroOrOneExpr{
								pos: position{line: 70, col: 104, offset: 2668},
								expr: &ruleRefExpr{
									pos:  position{line: 70, col: 105, offset: 2669},
									name: "DocumentAuthorNamePart",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 70, col: 130, offset: 2694},
							label: "email",
							expr: &zeroOrOneExpr{
								pos: position{line: 70, col: 136, offset: 2700},
								expr: &ruleRefExpr{
									pos:  position{line: 70, col: 137, offset: 2701},
									name: "DocumentAuthorEmail",
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 70, col: 159, offset: 2723},
							expr: &ruleRefExpr{
								pos:  position{line: 70, col: 159, offset: 2723},
								name: "WS",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 70, col: 163, offset: 2727},
							expr: &litMatcher{
								pos:        position{line: 70, col: 163, offset: 2727},
								val:        ";",
								ignoreCase: false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 70, col: 168, offset: 2732},
							expr: &ruleRefExpr{
								pos:  position{line: 70, col: 168, offset: 2732},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthorNamePart",
			pos:  position{line: 75, col: 1, offset: 2897},
			expr: &seqExpr{
				pos: position{line: 75, col: 27, offset: 2923},
				exprs: []interface{}{
					&notExpr{
						pos: position{line: 75, col: 27, offset: 2923},
						expr: &litMatcher{
							pos:        position{line: 75, col: 28, offset: 2924},
							val:        "<",
							ignoreCase: false,
						},
					},
					&notExpr{
						pos: position{line: 75, col: 32, offset: 2928},
						expr: &litMatcher{
							pos:        position{line: 75, col: 33, offset: 2929},
							val:        ";",
							ignoreCase: false,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 37, offset: 2933},
						name: "Characters",
					},
					&zeroOrMoreExpr{
						pos: position{line: 75, col: 48, offset: 2944},
						expr: &ruleRefExpr{
							pos:  position{line: 75, col: 48, offset: 2944},
							name: "WS",
						},
					},
				},
			},
		},
		{
			name: "DocumentAuthorEmail",
			pos:  position{line: 77, col: 1, offset: 2949},
			expr: &seqExpr{
				pos: position{line: 77, col: 24, offset: 2972},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 77, col: 24, offset: 2972},
						val:        "<",
						ignoreCase: false,
					},
					&labeledExpr{
						pos:   position{line: 77, col: 28, offset: 2976},
						label: "email",
						expr: &oneOrMoreExpr{
							pos: position{line: 77, col: 34, offset: 2982},
							expr: &seqExpr{
								pos: position{line: 77, col: 35, offset: 2983},
								exprs: []interface{}{
									&notExpr{
										pos: position{line: 77, col: 35, offset: 2983},
										expr: &litMatcher{
											pos:        position{line: 77, col: 36, offset: 2984},
											val:        ">",
											ignoreCase: false,
										},
									},
									&notExpr{
										pos: position{line: 77, col: 40, offset: 2988},
										expr: &ruleRefExpr{
											pos:  position{line: 77, col: 41, offset: 2989},
											name: "EOL",
										},
									},
									&anyMatcher{
										line: 77, col: 45, offset: 2993,
									},
								},
							},
						},
					},
					&litMatcher{
						pos:        position{line: 77, col: 49, offset: 2997},
						val:        ">",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "DocumentRevision",
			pos:  position{line: 81, col: 1, offset: 3133},
			expr: &actionExpr{
				pos: position{line: 81, col: 21, offset: 3153},
				run: (*parser).callonDocumentRevision1,
				expr: &seqExpr{
					pos: position{line: 81, col: 21, offset: 3153},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 81, col: 21, offset: 3153},
							expr: &ruleRefExpr{
								pos:  position{line: 81, col: 21, offset: 3153},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 81, col: 25, offset: 3157},
							expr: &litMatcher{
								pos:        position{line: 81, col: 26, offset: 3158},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 81, col: 30, offset: 3162},
							label: "revnumber",
							expr: &zeroOrOneExpr{
								pos: position{line: 81, col: 40, offset: 3172},
								expr: &ruleRefExpr{
									pos:  position{line: 81, col: 41, offset: 3173},
									name: "DocumentRevisionNumber",
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 81, col: 66, offset: 3198},
							expr: &litMatcher{
								pos:        position{line: 81, col: 66, offset: 3198},
								val:        ",",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 81, col: 71, offset: 3203},
							label: "revdate",
							expr: &zeroOrOneExpr{
								pos: position{line: 81, col: 79, offset: 3211},
								expr: &ruleRefExpr{
									pos:  position{line: 81, col: 80, offset: 3212},
									name: "DocumentRevisionDate",
								},
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 81, col: 103, offset: 3235},
							expr: &litMatcher{
								pos:        position{line: 81, col: 103, offset: 3235},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 81, col: 108, offset: 3240},
							label: "revremark",
							expr: &zeroOrOneExpr{
								pos: position{line: 81, col: 118, offset: 3250},
								expr: &ruleRefExpr{
									pos:  position{line: 81, col: 119, offset: 3251},
									name: "DocumentRevisionRemark",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 81, col: 144, offset: 3276},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentRevisionNumber",
			pos:  position{line: 86, col: 1, offset: 3449},
			expr: &choiceExpr{
				pos: position{line: 86, col: 27, offset: 3475},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 86, col: 27, offset: 3475},
						exprs: []interface{}{
							&litMatcher{
								pos:        position{line: 86, col: 27, offset: 3475},
								val:        "v",
								ignoreCase: true,
							},
							&ruleRefExpr{
								pos:  position{line: 86, col: 32, offset: 3480},
								name: "DIGIT",
							},
							&zeroOrMoreExpr{
								pos: position{line: 86, col: 39, offset: 3487},
								expr: &seqExpr{
									pos: position{line: 86, col: 40, offset: 3488},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 86, col: 40, offset: 3488},
											expr: &ruleRefExpr{
												pos:  position{line: 86, col: 41, offset: 3489},
												name: "EOL",
											},
										},
										&notExpr{
											pos: position{line: 86, col: 45, offset: 3493},
											expr: &litMatcher{
												pos:        position{line: 86, col: 46, offset: 3494},
												val:        ",",
												ignoreCase: false,
											},
										},
										&notExpr{
											pos: position{line: 86, col: 50, offset: 3498},
											expr: &litMatcher{
												pos:        position{line: 86, col: 51, offset: 3499},
												val:        ":",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 86, col: 55, offset: 3503,
										},
									},
								},
							},
						},
					},
					&seqExpr{
						pos: position{line: 86, col: 61, offset: 3509},
						exprs: []interface{}{
							&zeroOrOneExpr{
								pos: position{line: 86, col: 61, offset: 3509},
								expr: &litMatcher{
									pos:        position{line: 86, col: 61, offset: 3509},
									val:        "v",
									ignoreCase: true,
								},
							},
							&ruleRefExpr{
								pos:  position{line: 86, col: 67, offset: 3515},
								name: "DIGIT",
							},
							&zeroOrMoreExpr{
								pos: position{line: 86, col: 74, offset: 3522},
								expr: &seqExpr{
									pos: position{line: 86, col: 75, offset: 3523},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 86, col: 75, offset: 3523},
											expr: &ruleRefExpr{
												pos:  position{line: 86, col: 76, offset: 3524},
												name: "EOL",
											},
										},
										&notExpr{
											pos: position{line: 86, col: 80, offset: 3528},
											expr: &litMatcher{
												pos:        position{line: 86, col: 81, offset: 3529},
												val:        ",",
												ignoreCase: false,
											},
										},
										&notExpr{
											pos: position{line: 86, col: 85, offset: 3533},
											expr: &litMatcher{
												pos:        position{line: 86, col: 86, offset: 3534},
												val:        ":",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 86, col: 90, offset: 3538,
										},
									},
								},
							},
							&zeroOrMoreExpr{
								pos: position{line: 86, col: 94, offset: 3542},
								expr: &ruleRefExpr{
									pos:  position{line: 86, col: 94, offset: 3542},
									name: "WS",
								},
							},
							&andExpr{
								pos: position{line: 86, col: 98, offset: 3546},
								expr: &litMatcher{
									pos:        position{line: 86, col: 99, offset: 3547},
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
			pos:  position{line: 87, col: 1, offset: 3551},
			expr: &zeroOrMoreExpr{
				pos: position{line: 87, col: 25, offset: 3575},
				expr: &seqExpr{
					pos: position{line: 87, col: 26, offset: 3576},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 87, col: 26, offset: 3576},
							expr: &ruleRefExpr{
								pos:  position{line: 87, col: 27, offset: 3577},
								name: "EOL",
							},
						},
						&notExpr{
							pos: position{line: 87, col: 31, offset: 3581},
							expr: &litMatcher{
								pos:        position{line: 87, col: 32, offset: 3582},
								val:        ":",
								ignoreCase: false,
							},
						},
						&anyMatcher{
							line: 87, col: 36, offset: 3586,
						},
					},
				},
			},
		},
		{
			name: "DocumentRevisionRemark",
			pos:  position{line: 88, col: 1, offset: 3591},
			expr: &zeroOrMoreExpr{
				pos: position{line: 88, col: 27, offset: 3617},
				expr: &seqExpr{
					pos: position{line: 88, col: 28, offset: 3618},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 88, col: 28, offset: 3618},
							expr: &ruleRefExpr{
								pos:  position{line: 88, col: 29, offset: 3619},
								name: "EOL",
							},
						},
						&anyMatcher{
							line: 88, col: 33, offset: 3623,
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclaration",
			pos:  position{line: 93, col: 1, offset: 3743},
			expr: &choiceExpr{
				pos: position{line: 93, col: 33, offset: 3775},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 93, col: 33, offset: 3775},
						name: "DocumentAttributeDeclarationWithNameOnly",
					},
					&ruleRefExpr{
						pos:  position{line: 93, col: 76, offset: 3818},
						name: "DocumentAttributeDeclarationWithNameAndValue",
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclarationWithNameOnly",
			pos:  position{line: 95, col: 1, offset: 3865},
			expr: &actionExpr{
				pos: position{line: 95, col: 45, offset: 3909},
				run: (*parser).callonDocumentAttributeDeclarationWithNameOnly1,
				expr: &seqExpr{
					pos: position{line: 95, col: 45, offset: 3909},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 95, col: 45, offset: 3909},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 95, col: 49, offset: 3913},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 95, col: 55, offset: 3919},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 95, col: 70, offset: 3934},
							val:        ":",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 95, col: 74, offset: 3938},
							expr: &ruleRefExpr{
								pos:  position{line: 95, col: 74, offset: 3938},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 95, col: 78, offset: 3942},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeDeclarationWithNameAndValue",
			pos:  position{line: 99, col: 1, offset: 4027},
			expr: &actionExpr{
				pos: position{line: 99, col: 49, offset: 4075},
				run: (*parser).callonDocumentAttributeDeclarationWithNameAndValue1,
				expr: &seqExpr{
					pos: position{line: 99, col: 49, offset: 4075},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 99, col: 49, offset: 4075},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 99, col: 53, offset: 4079},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 99, col: 59, offset: 4085},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 99, col: 74, offset: 4100},
							val:        ":",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 99, col: 78, offset: 4104},
							expr: &ruleRefExpr{
								pos:  position{line: 99, col: 78, offset: 4104},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 99, col: 82, offset: 4108},
							label: "value",
							expr: &zeroOrMoreExpr{
								pos: position{line: 99, col: 88, offset: 4114},
								expr: &seqExpr{
									pos: position{line: 99, col: 89, offset: 4115},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 99, col: 89, offset: 4115},
											expr: &ruleRefExpr{
												pos:  position{line: 99, col: 90, offset: 4116},
												name: "NEWLINE",
											},
										},
										&anyMatcher{
											line: 99, col: 98, offset: 4124,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 99, col: 102, offset: 4128},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeReset",
			pos:  position{line: 103, col: 1, offset: 4231},
			expr: &choiceExpr{
				pos: position{line: 103, col: 27, offset: 4257},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 103, col: 27, offset: 4257},
						name: "DocumentAttributeResetWithSectionTitleBangSymbol",
					},
					&ruleRefExpr{
						pos:  position{line: 103, col: 78, offset: 4308},
						name: "DocumentAttributeResetWithTrailingBangSymbol",
					},
				},
			},
		},
		{
			name: "DocumentAttributeResetWithSectionTitleBangSymbol",
			pos:  position{line: 105, col: 1, offset: 4354},
			expr: &actionExpr{
				pos: position{line: 105, col: 53, offset: 4406},
				run: (*parser).callonDocumentAttributeResetWithSectionTitleBangSymbol1,
				expr: &seqExpr{
					pos: position{line: 105, col: 53, offset: 4406},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 105, col: 53, offset: 4406},
							val:        ":!",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 105, col: 58, offset: 4411},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 105, col: 64, offset: 4417},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 105, col: 79, offset: 4432},
							val:        ":",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 105, col: 83, offset: 4436},
							expr: &ruleRefExpr{
								pos:  position{line: 105, col: 83, offset: 4436},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 105, col: 87, offset: 4440},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeResetWithTrailingBangSymbol",
			pos:  position{line: 109, col: 1, offset: 4514},
			expr: &actionExpr{
				pos: position{line: 109, col: 49, offset: 4562},
				run: (*parser).callonDocumentAttributeResetWithTrailingBangSymbol1,
				expr: &seqExpr{
					pos: position{line: 109, col: 49, offset: 4562},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 109, col: 49, offset: 4562},
							val:        ":",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 109, col: 53, offset: 4566},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 109, col: 59, offset: 4572},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 109, col: 74, offset: 4587},
							val:        "!:",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 109, col: 79, offset: 4592},
							expr: &ruleRefExpr{
								pos:  position{line: 109, col: 79, offset: 4592},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 109, col: 83, offset: 4596},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "DocumentAttributeSubstitution",
			pos:  position{line: 113, col: 1, offset: 4670},
			expr: &actionExpr{
				pos: position{line: 113, col: 34, offset: 4703},
				run: (*parser).callonDocumentAttributeSubstitution1,
				expr: &seqExpr{
					pos: position{line: 113, col: 34, offset: 4703},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 113, col: 34, offset: 4703},
							val:        "{",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 113, col: 38, offset: 4707},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 113, col: 44, offset: 4713},
								name: "AttributeName",
							},
						},
						&litMatcher{
							pos:        position{line: 113, col: 59, offset: 4728},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "AttributeName",
			pos:  position{line: 120, col: 1, offset: 4982},
			expr: &seqExpr{
				pos: position{line: 120, col: 18, offset: 4999},
				exprs: []interface{}{
					&choiceExpr{
						pos: position{line: 120, col: 19, offset: 5000},
						alternatives: []interface{}{
							&charClassMatcher{
								pos:        position{line: 120, col: 19, offset: 5000},
								val:        "[A-Z]",
								ranges:     []rune{'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
							&charClassMatcher{
								pos:        position{line: 120, col: 27, offset: 5008},
								val:        "[a-z]",
								ranges:     []rune{'a', 'z'},
								ignoreCase: false,
								inverted:   false,
							},
							&charClassMatcher{
								pos:        position{line: 120, col: 35, offset: 5016},
								val:        "[0-9]",
								ranges:     []rune{'0', '9'},
								ignoreCase: false,
								inverted:   false,
							},
							&litMatcher{
								pos:        position{line: 120, col: 43, offset: 5024},
								val:        "_",
								ignoreCase: false,
							},
						},
					},
					&zeroOrMoreExpr{
						pos: position{line: 120, col: 48, offset: 5029},
						expr: &choiceExpr{
							pos: position{line: 120, col: 49, offset: 5030},
							alternatives: []interface{}{
								&charClassMatcher{
									pos:        position{line: 120, col: 49, offset: 5030},
									val:        "[A-Z]",
									ranges:     []rune{'A', 'Z'},
									ignoreCase: false,
									inverted:   false,
								},
								&charClassMatcher{
									pos:        position{line: 120, col: 57, offset: 5038},
									val:        "[a-z]",
									ranges:     []rune{'a', 'z'},
									ignoreCase: false,
									inverted:   false,
								},
								&charClassMatcher{
									pos:        position{line: 120, col: 65, offset: 5046},
									val:        "[0-9]",
									ranges:     []rune{'0', '9'},
									ignoreCase: false,
									inverted:   false,
								},
								&litMatcher{
									pos:        position{line: 120, col: 73, offset: 5054},
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
			pos:  position{line: 125, col: 1, offset: 5174},
			expr: &seqExpr{
				pos: position{line: 125, col: 25, offset: 5198},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 125, col: 25, offset: 5198},
						val:        "toc::[]",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 125, col: 35, offset: 5208},
						name: "NEWLINE",
					},
				},
			},
		},
		{
			name: "Section",
			pos:  position{line: 130, col: 1, offset: 5321},
			expr: &choiceExpr{
				pos: position{line: 130, col: 12, offset: 5332},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 130, col: 12, offset: 5332},
						name: "Section1",
					},
					&ruleRefExpr{
						pos:  position{line: 130, col: 23, offset: 5343},
						name: "Section2",
					},
					&ruleRefExpr{
						pos:  position{line: 130, col: 34, offset: 5354},
						name: "Section3",
					},
					&ruleRefExpr{
						pos:  position{line: 130, col: 45, offset: 5365},
						name: "Section4",
					},
					&ruleRefExpr{
						pos:  position{line: 130, col: 56, offset: 5376},
						name: "Section5",
					},
				},
			},
		},
		{
			name: "Section1",
			pos:  position{line: 133, col: 1, offset: 5387},
			expr: &actionExpr{
				pos: position{line: 133, col: 13, offset: 5399},
				run: (*parser).callonSection11,
				expr: &seqExpr{
					pos: position{line: 133, col: 13, offset: 5399},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 133, col: 13, offset: 5399},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 133, col: 21, offset: 5407},
								name: "Section1Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 133, col: 36, offset: 5422},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 133, col: 46, offset: 5432},
								expr: &ruleRefExpr{
									pos:  position{line: 133, col: 46, offset: 5432},
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
			pos:  position{line: 137, col: 1, offset: 5539},
			expr: &actionExpr{
				pos: position{line: 137, col: 18, offset: 5556},
				run: (*parser).callonSection1Block1,
				expr: &seqExpr{
					pos: position{line: 137, col: 18, offset: 5556},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 137, col: 18, offset: 5556},
							expr: &ruleRefExpr{
								pos:  position{line: 137, col: 19, offset: 5557},
								name: "Section1",
							},
						},
						&labeledExpr{
							pos:   position{line: 137, col: 28, offset: 5566},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 137, col: 37, offset: 5575},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 137, col: 37, offset: 5575},
										name: "Section2",
									},
									&ruleRefExpr{
										pos:  position{line: 137, col: 48, offset: 5586},
										name: "Section3",
									},
									&ruleRefExpr{
										pos:  position{line: 137, col: 59, offset: 5597},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 137, col: 70, offset: 5608},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 137, col: 81, offset: 5619},
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
			pos:  position{line: 141, col: 1, offset: 5681},
			expr: &actionExpr{
				pos: position{line: 141, col: 13, offset: 5693},
				run: (*parser).callonSection21,
				expr: &seqExpr{
					pos: position{line: 141, col: 13, offset: 5693},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 141, col: 13, offset: 5693},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 141, col: 21, offset: 5701},
								name: "Section2Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 141, col: 36, offset: 5716},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 141, col: 46, offset: 5726},
								expr: &ruleRefExpr{
									pos:  position{line: 141, col: 46, offset: 5726},
									name: "Section2Block",
								},
							},
						},
						&andExpr{
							pos: position{line: 141, col: 62, offset: 5742},
							expr: &zeroOrMoreExpr{
								pos: position{line: 141, col: 63, offset: 5743},
								expr: &ruleRefExpr{
									pos:  position{line: 141, col: 64, offset: 5744},
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
			pos:  position{line: 145, col: 1, offset: 5846},
			expr: &actionExpr{
				pos: position{line: 145, col: 18, offset: 5863},
				run: (*parser).callonSection2Block1,
				expr: &seqExpr{
					pos: position{line: 145, col: 18, offset: 5863},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 145, col: 18, offset: 5863},
							expr: &ruleRefExpr{
								pos:  position{line: 145, col: 19, offset: 5864},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 145, col: 28, offset: 5873},
							expr: &ruleRefExpr{
								pos:  position{line: 145, col: 29, offset: 5874},
								name: "Section2",
							},
						},
						&labeledExpr{
							pos:   position{line: 145, col: 38, offset: 5883},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 145, col: 47, offset: 5892},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 145, col: 47, offset: 5892},
										name: "Section3",
									},
									&ruleRefExpr{
										pos:  position{line: 145, col: 58, offset: 5903},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 145, col: 69, offset: 5914},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 145, col: 80, offset: 5925},
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
			pos:  position{line: 149, col: 1, offset: 5987},
			expr: &actionExpr{
				pos: position{line: 149, col: 13, offset: 5999},
				run: (*parser).callonSection31,
				expr: &seqExpr{
					pos: position{line: 149, col: 13, offset: 5999},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 149, col: 13, offset: 5999},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 149, col: 21, offset: 6007},
								name: "Section3Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 149, col: 36, offset: 6022},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 149, col: 46, offset: 6032},
								expr: &ruleRefExpr{
									pos:  position{line: 149, col: 46, offset: 6032},
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
			pos:  position{line: 153, col: 1, offset: 6139},
			expr: &actionExpr{
				pos: position{line: 153, col: 18, offset: 6156},
				run: (*parser).callonSection3Block1,
				expr: &seqExpr{
					pos: position{line: 153, col: 18, offset: 6156},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 153, col: 18, offset: 6156},
							expr: &ruleRefExpr{
								pos:  position{line: 153, col: 19, offset: 6157},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 153, col: 28, offset: 6166},
							expr: &ruleRefExpr{
								pos:  position{line: 153, col: 29, offset: 6167},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 153, col: 38, offset: 6176},
							expr: &ruleRefExpr{
								pos:  position{line: 153, col: 39, offset: 6177},
								name: "Section3",
							},
						},
						&labeledExpr{
							pos:   position{line: 153, col: 48, offset: 6186},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 153, col: 57, offset: 6195},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 153, col: 57, offset: 6195},
										name: "Section4",
									},
									&ruleRefExpr{
										pos:  position{line: 153, col: 68, offset: 6206},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 153, col: 79, offset: 6217},
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
			pos:  position{line: 157, col: 1, offset: 6279},
			expr: &actionExpr{
				pos: position{line: 157, col: 13, offset: 6291},
				run: (*parser).callonSection41,
				expr: &seqExpr{
					pos: position{line: 157, col: 13, offset: 6291},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 157, col: 13, offset: 6291},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 157, col: 21, offset: 6299},
								name: "Section4Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 157, col: 36, offset: 6314},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 157, col: 46, offset: 6324},
								expr: &ruleRefExpr{
									pos:  position{line: 157, col: 46, offset: 6324},
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
			pos:  position{line: 161, col: 1, offset: 6431},
			expr: &actionExpr{
				pos: position{line: 161, col: 18, offset: 6448},
				run: (*parser).callonSection4Block1,
				expr: &seqExpr{
					pos: position{line: 161, col: 18, offset: 6448},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 161, col: 18, offset: 6448},
							expr: &ruleRefExpr{
								pos:  position{line: 161, col: 19, offset: 6449},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 161, col: 28, offset: 6458},
							expr: &ruleRefExpr{
								pos:  position{line: 161, col: 29, offset: 6459},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 161, col: 38, offset: 6468},
							expr: &ruleRefExpr{
								pos:  position{line: 161, col: 39, offset: 6469},
								name: "Section3",
							},
						},
						&notExpr{
							pos: position{line: 161, col: 48, offset: 6478},
							expr: &ruleRefExpr{
								pos:  position{line: 161, col: 49, offset: 6479},
								name: "Section4",
							},
						},
						&labeledExpr{
							pos:   position{line: 161, col: 58, offset: 6488},
							label: "content",
							expr: &choiceExpr{
								pos: position{line: 161, col: 67, offset: 6497},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 161, col: 67, offset: 6497},
										name: "Section5",
									},
									&ruleRefExpr{
										pos:  position{line: 161, col: 78, offset: 6508},
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
			pos:  position{line: 165, col: 1, offset: 6570},
			expr: &actionExpr{
				pos: position{line: 165, col: 13, offset: 6582},
				run: (*parser).callonSection51,
				expr: &seqExpr{
					pos: position{line: 165, col: 13, offset: 6582},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 165, col: 13, offset: 6582},
							label: "header",
							expr: &ruleRefExpr{
								pos:  position{line: 165, col: 21, offset: 6590},
								name: "Section5Title",
							},
						},
						&labeledExpr{
							pos:   position{line: 165, col: 36, offset: 6605},
							label: "elements",
							expr: &zeroOrMoreExpr{
								pos: position{line: 165, col: 46, offset: 6615},
								expr: &ruleRefExpr{
									pos:  position{line: 165, col: 46, offset: 6615},
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
			pos:  position{line: 169, col: 1, offset: 6722},
			expr: &actionExpr{
				pos: position{line: 169, col: 18, offset: 6739},
				run: (*parser).callonSection5Block1,
				expr: &seqExpr{
					pos: position{line: 169, col: 18, offset: 6739},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 169, col: 18, offset: 6739},
							expr: &ruleRefExpr{
								pos:  position{line: 169, col: 19, offset: 6740},
								name: "Section1",
							},
						},
						&notExpr{
							pos: position{line: 169, col: 28, offset: 6749},
							expr: &ruleRefExpr{
								pos:  position{line: 169, col: 29, offset: 6750},
								name: "Section2",
							},
						},
						&notExpr{
							pos: position{line: 169, col: 38, offset: 6759},
							expr: &ruleRefExpr{
								pos:  position{line: 169, col: 39, offset: 6760},
								name: "Section3",
							},
						},
						&notExpr{
							pos: position{line: 169, col: 48, offset: 6769},
							expr: &ruleRefExpr{
								pos:  position{line: 169, col: 49, offset: 6770},
								name: "Section4",
							},
						},
						&notExpr{
							pos: position{line: 169, col: 58, offset: 6779},
							expr: &ruleRefExpr{
								pos:  position{line: 169, col: 59, offset: 6780},
								name: "Section5",
							},
						},
						&labeledExpr{
							pos:   position{line: 169, col: 68, offset: 6789},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 169, col: 77, offset: 6798},
								name: "BlockElement",
							},
						},
					},
				},
			},
		},
		{
			name: "SectionTitle",
			pos:  position{line: 177, col: 1, offset: 6971},
			expr: &choiceExpr{
				pos: position{line: 177, col: 17, offset: 6987},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 177, col: 17, offset: 6987},
						name: "Section1Title",
					},
					&ruleRefExpr{
						pos:  position{line: 177, col: 33, offset: 7003},
						name: "Section2Title",
					},
					&ruleRefExpr{
						pos:  position{line: 177, col: 49, offset: 7019},
						name: "Section3Title",
					},
					&ruleRefExpr{
						pos:  position{line: 177, col: 65, offset: 7035},
						name: "Section4Title",
					},
					&ruleRefExpr{
						pos:  position{line: 177, col: 81, offset: 7051},
						name: "Section5Title",
					},
				},
			},
		},
		{
			name: "Section1Title",
			pos:  position{line: 179, col: 1, offset: 7066},
			expr: &actionExpr{
				pos: position{line: 179, col: 18, offset: 7083},
				run: (*parser).callonSection1Title1,
				expr: &seqExpr{
					pos: position{line: 179, col: 18, offset: 7083},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 179, col: 18, offset: 7083},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 179, col: 29, offset: 7094},
								expr: &ruleRefExpr{
									pos:  position{line: 179, col: 30, offset: 7095},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 179, col: 49, offset: 7114},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 179, col: 56, offset: 7121},
								val:        "==",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 179, col: 62, offset: 7127},
							expr: &ruleRefExpr{
								pos:  position{line: 179, col: 62, offset: 7127},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 179, col: 66, offset: 7131},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 179, col: 75, offset: 7140},
								name: "InlineContent",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 179, col: 90, offset: 7155},
							expr: &ruleRefExpr{
								pos:  position{line: 179, col: 90, offset: 7155},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 179, col: 94, offset: 7159},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 179, col: 97, offset: 7162},
								expr: &ruleRefExpr{
									pos:  position{line: 179, col: 98, offset: 7163},
									name: "InlineElementID",
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 179, col: 116, offset: 7181},
							expr: &ruleRefExpr{
								pos:  position{line: 179, col: 116, offset: 7181},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 179, col: 120, offset: 7185},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 179, col: 125, offset: 7190},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 179, col: 125, offset: 7190},
									expr: &ruleRefExpr{
										pos:  position{line: 179, col: 125, offset: 7190},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 179, col: 138, offset: 7203},
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
			pos:  position{line: 183, col: 1, offset: 7318},
			expr: &actionExpr{
				pos: position{line: 183, col: 18, offset: 7335},
				run: (*parser).callonSection2Title1,
				expr: &seqExpr{
					pos: position{line: 183, col: 18, offset: 7335},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 183, col: 18, offset: 7335},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 183, col: 29, offset: 7346},
								expr: &ruleRefExpr{
									pos:  position{line: 183, col: 30, offset: 7347},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 183, col: 49, offset: 7366},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 183, col: 56, offset: 7373},
								val:        "===",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 183, col: 63, offset: 7380},
							expr: &ruleRefExpr{
								pos:  position{line: 183, col: 63, offset: 7380},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 183, col: 67, offset: 7384},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 183, col: 76, offset: 7393},
								name: "InlineContent",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 183, col: 91, offset: 7408},
							expr: &ruleRefExpr{
								pos:  position{line: 183, col: 91, offset: 7408},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 183, col: 95, offset: 7412},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 183, col: 98, offset: 7415},
								expr: &ruleRefExpr{
									pos:  position{line: 183, col: 99, offset: 7416},
									name: "InlineElementID",
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 183, col: 117, offset: 7434},
							expr: &ruleRefExpr{
								pos:  position{line: 183, col: 117, offset: 7434},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 183, col: 121, offset: 7438},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 183, col: 126, offset: 7443},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 183, col: 126, offset: 7443},
									expr: &ruleRefExpr{
										pos:  position{line: 183, col: 126, offset: 7443},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 183, col: 139, offset: 7456},
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
			pos:  position{line: 187, col: 1, offset: 7570},
			expr: &actionExpr{
				pos: position{line: 187, col: 18, offset: 7587},
				run: (*parser).callonSection3Title1,
				expr: &seqExpr{
					pos: position{line: 187, col: 18, offset: 7587},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 187, col: 18, offset: 7587},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 187, col: 29, offset: 7598},
								expr: &ruleRefExpr{
									pos:  position{line: 187, col: 30, offset: 7599},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 187, col: 49, offset: 7618},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 187, col: 56, offset: 7625},
								val:        "====",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 187, col: 64, offset: 7633},
							expr: &ruleRefExpr{
								pos:  position{line: 187, col: 64, offset: 7633},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 187, col: 68, offset: 7637},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 187, col: 77, offset: 7646},
								name: "InlineContent",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 187, col: 92, offset: 7661},
							expr: &ruleRefExpr{
								pos:  position{line: 187, col: 92, offset: 7661},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 187, col: 96, offset: 7665},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 187, col: 99, offset: 7668},
								expr: &ruleRefExpr{
									pos:  position{line: 187, col: 100, offset: 7669},
									name: "InlineElementID",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 187, col: 118, offset: 7687},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 187, col: 123, offset: 7692},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 187, col: 123, offset: 7692},
									expr: &ruleRefExpr{
										pos:  position{line: 187, col: 123, offset: 7692},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 187, col: 136, offset: 7705},
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
			pos:  position{line: 191, col: 1, offset: 7819},
			expr: &actionExpr{
				pos: position{line: 191, col: 18, offset: 7836},
				run: (*parser).callonSection4Title1,
				expr: &seqExpr{
					pos: position{line: 191, col: 18, offset: 7836},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 191, col: 18, offset: 7836},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 191, col: 29, offset: 7847},
								expr: &ruleRefExpr{
									pos:  position{line: 191, col: 30, offset: 7848},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 191, col: 49, offset: 7867},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 191, col: 56, offset: 7874},
								val:        "=====",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 191, col: 65, offset: 7883},
							expr: &ruleRefExpr{
								pos:  position{line: 191, col: 65, offset: 7883},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 191, col: 69, offset: 7887},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 191, col: 78, offset: 7896},
								name: "InlineContent",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 191, col: 93, offset: 7911},
							expr: &ruleRefExpr{
								pos:  position{line: 191, col: 93, offset: 7911},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 191, col: 97, offset: 7915},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 191, col: 100, offset: 7918},
								expr: &ruleRefExpr{
									pos:  position{line: 191, col: 101, offset: 7919},
									name: "InlineElementID",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 191, col: 119, offset: 7937},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 191, col: 124, offset: 7942},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 191, col: 124, offset: 7942},
									expr: &ruleRefExpr{
										pos:  position{line: 191, col: 124, offset: 7942},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 191, col: 137, offset: 7955},
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
			pos:  position{line: 195, col: 1, offset: 8069},
			expr: &actionExpr{
				pos: position{line: 195, col: 18, offset: 8086},
				run: (*parser).callonSection5Title1,
				expr: &seqExpr{
					pos: position{line: 195, col: 18, offset: 8086},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 195, col: 18, offset: 8086},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 195, col: 29, offset: 8097},
								expr: &ruleRefExpr{
									pos:  position{line: 195, col: 30, offset: 8098},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 195, col: 49, offset: 8117},
							label: "level",
							expr: &litMatcher{
								pos:        position{line: 195, col: 56, offset: 8124},
								val:        "======",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 195, col: 66, offset: 8134},
							expr: &ruleRefExpr{
								pos:  position{line: 195, col: 66, offset: 8134},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 195, col: 70, offset: 8138},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 195, col: 79, offset: 8147},
								name: "InlineContent",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 195, col: 94, offset: 8162},
							expr: &ruleRefExpr{
								pos:  position{line: 195, col: 94, offset: 8162},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 195, col: 98, offset: 8166},
							label: "id",
							expr: &zeroOrOneExpr{
								pos: position{line: 195, col: 101, offset: 8169},
								expr: &ruleRefExpr{
									pos:  position{line: 195, col: 102, offset: 8170},
									name: "InlineElementID",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 195, col: 120, offset: 8188},
							name: "EOL",
						},
						&choiceExpr{
							pos: position{line: 195, col: 125, offset: 8193},
							alternatives: []interface{}{
								&zeroOrOneExpr{
									pos: position{line: 195, col: 125, offset: 8193},
									expr: &ruleRefExpr{
										pos:  position{line: 195, col: 125, offset: 8193},
										name: "BlankLine",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 195, col: 138, offset: 8206},
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
			pos:  position{line: 202, col: 1, offset: 8421},
			expr: &actionExpr{
				pos: position{line: 202, col: 9, offset: 8429},
				run: (*parser).callonList1,
				expr: &seqExpr{
					pos: position{line: 202, col: 9, offset: 8429},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 202, col: 9, offset: 8429},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 202, col: 20, offset: 8440},
								expr: &ruleRefExpr{
									pos:  position{line: 202, col: 21, offset: 8441},
									name: "ListAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 204, col: 5, offset: 8530},
							label: "elements",
							expr: &oneOrMoreExpr{
								pos: position{line: 204, col: 14, offset: 8539},
								expr: &choiceExpr{
									pos: position{line: 204, col: 15, offset: 8540},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 204, col: 15, offset: 8540},
											name: "UnorderedListItem",
										},
										&ruleRefExpr{
											pos:  position{line: 204, col: 35, offset: 8560},
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
			pos:  position{line: 208, col: 1, offset: 8662},
			expr: &actionExpr{
				pos: position{line: 208, col: 18, offset: 8679},
				run: (*parser).callonListAttribute1,
				expr: &seqExpr{
					pos: position{line: 208, col: 18, offset: 8679},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 208, col: 18, offset: 8679},
							label: "attribute",
							expr: &choiceExpr{
								pos: position{line: 208, col: 29, offset: 8690},
								alternatives: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 208, col: 29, offset: 8690},
										name: "HorizontalLayout",
									},
									&ruleRefExpr{
										pos:  position{line: 208, col: 48, offset: 8709},
										name: "ListID",
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 208, col: 56, offset: 8717},
							name: "NEWLINE",
						},
					},
				},
			},
		},
		{
			name: "ListID",
			pos:  position{line: 212, col: 1, offset: 8756},
			expr: &actionExpr{
				pos: position{line: 212, col: 11, offset: 8766},
				run: (*parser).callonListID1,
				expr: &seqExpr{
					pos: position{line: 212, col: 11, offset: 8766},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 212, col: 11, offset: 8766},
							val:        "[#",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 212, col: 16, offset: 8771},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 212, col: 20, offset: 8775},
								name: "ID",
							},
						},
						&litMatcher{
							pos:        position{line: 212, col: 24, offset: 8779},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "HorizontalLayout",
			pos:  position{line: 216, col: 1, offset: 8845},
			expr: &actionExpr{
				pos: position{line: 216, col: 21, offset: 8865},
				run: (*parser).callonHorizontalLayout1,
				expr: &litMatcher{
					pos:        position{line: 216, col: 21, offset: 8865},
					val:        "[horizontal]",
					ignoreCase: false,
				},
			},
		},
		{
			name: "ListParagraph",
			pos:  position{line: 220, col: 1, offset: 8948},
			expr: &actionExpr{
				pos: position{line: 220, col: 19, offset: 8966},
				run: (*parser).callonListParagraph1,
				expr: &labeledExpr{
					pos:   position{line: 220, col: 19, offset: 8966},
					label: "lines",
					expr: &oneOrMoreExpr{
						pos: position{line: 220, col: 25, offset: 8972},
						expr: &seqExpr{
							pos: position{line: 220, col: 26, offset: 8973},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 220, col: 26, offset: 8973},
									expr: &ruleRefExpr{
										pos:  position{line: 220, col: 28, offset: 8975},
										name: "ListItemContinuation",
									},
								},
								&notExpr{
									pos: position{line: 220, col: 50, offset: 8997},
									expr: &ruleRefExpr{
										pos:  position{line: 220, col: 52, offset: 8999},
										name: "UnorderedListItemPrefix",
									},
								},
								&notExpr{
									pos: position{line: 220, col: 77, offset: 9024},
									expr: &seqExpr{
										pos: position{line: 220, col: 79, offset: 9026},
										exprs: []interface{}{
											&ruleRefExpr{
												pos:  position{line: 220, col: 79, offset: 9026},
												name: "LabeledListItemTerm",
											},
											&ruleRefExpr{
												pos:  position{line: 220, col: 99, offset: 9046},
												name: "LabeledListItemSeparator",
											},
										},
									},
								},
								&ruleRefExpr{
									pos:  position{line: 220, col: 125, offset: 9072},
									name: "InlineContentWithTrailingSpaces",
								},
								&ruleRefExpr{
									pos:  position{line: 220, col: 157, offset: 9104},
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
			pos:  position{line: 224, col: 1, offset: 9173},
			expr: &actionExpr{
				pos: position{line: 224, col: 25, offset: 9197},
				run: (*parser).callonListItemContinuation1,
				expr: &seqExpr{
					pos: position{line: 224, col: 25, offset: 9197},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 224, col: 25, offset: 9197},
							val:        "+",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 224, col: 29, offset: 9201},
							expr: &ruleRefExpr{
								pos:  position{line: 224, col: 29, offset: 9201},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 224, col: 33, offset: 9205},
							name: "NEWLINE",
						},
					},
				},
			},
		},
		{
			name: "ContinuedBlockElement",
			pos:  position{line: 228, col: 1, offset: 9261},
			expr: &actionExpr{
				pos: position{line: 228, col: 26, offset: 9286},
				run: (*parser).callonContinuedBlockElement1,
				expr: &seqExpr{
					pos: position{line: 228, col: 26, offset: 9286},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 228, col: 26, offset: 9286},
							name: "ListItemContinuation",
						},
						&labeledExpr{
							pos:   position{line: 228, col: 47, offset: 9307},
							label: "element",
							expr: &ruleRefExpr{
								pos:  position{line: 228, col: 55, offset: 9315},
								name: "BlockElement",
							},
						},
					},
				},
			},
		},
		{
			name: "UnorderedListItem",
			pos:  position{line: 235, col: 1, offset: 9468},
			expr: &actionExpr{
				pos: position{line: 235, col: 22, offset: 9489},
				run: (*parser).callonUnorderedListItem1,
				expr: &seqExpr{
					pos: position{line: 235, col: 22, offset: 9489},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 235, col: 22, offset: 9489},
							label: "level",
							expr: &ruleRefExpr{
								pos:  position{line: 235, col: 29, offset: 9496},
								name: "UnorderedListItemPrefix",
							},
						},
						&labeledExpr{
							pos:   position{line: 235, col: 54, offset: 9521},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 235, col: 63, offset: 9530},
								name: "UnorderedListItemContent",
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 235, col: 89, offset: 9556},
							expr: &ruleRefExpr{
								pos:  position{line: 235, col: 89, offset: 9556},
								name: "BlankLine",
							},
						},
					},
				},
			},
		},
		{
			name: "UnorderedListItemPrefix",
			pos:  position{line: 239, col: 1, offset: 9647},
			expr: &actionExpr{
				pos: position{line: 239, col: 28, offset: 9674},
				run: (*parser).callonUnorderedListItemPrefix1,
				expr: &seqExpr{
					pos: position{line: 239, col: 28, offset: 9674},
					exprs: []interface{}{
						&zeroOrMoreExpr{
							pos: position{line: 239, col: 28, offset: 9674},
							expr: &ruleRefExpr{
								pos:  position{line: 239, col: 28, offset: 9674},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 239, col: 32, offset: 9678},
							label: "level",
							expr: &choiceExpr{
								pos: position{line: 239, col: 39, offset: 9685},
								alternatives: []interface{}{
									&oneOrMoreExpr{
										pos: position{line: 239, col: 39, offset: 9685},
										expr: &litMatcher{
											pos:        position{line: 239, col: 39, offset: 9685},
											val:        "*",
											ignoreCase: false,
										},
									},
									&litMatcher{
										pos:        position{line: 239, col: 46, offset: 9692},
										val:        "-",
										ignoreCase: false,
									},
								},
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 239, col: 51, offset: 9697},
							expr: &ruleRefExpr{
								pos:  position{line: 239, col: 51, offset: 9697},
								name: "WS",
							},
						},
					},
				},
			},
		},
		{
			name: "UnorderedListItemContent",
			pos:  position{line: 243, col: 1, offset: 9795},
			expr: &actionExpr{
				pos: position{line: 243, col: 29, offset: 9823},
				run: (*parser).callonUnorderedListItemContent1,
				expr: &labeledExpr{
					pos:   position{line: 243, col: 29, offset: 9823},
					label: "elements",
					expr: &seqExpr{
						pos: position{line: 243, col: 39, offset: 9833},
						exprs: []interface{}{
							&oneOrMoreExpr{
								pos: position{line: 243, col: 39, offset: 9833},
								expr: &ruleRefExpr{
									pos:  position{line: 243, col: 39, offset: 9833},
									name: "ListParagraph",
								},
							},
							&zeroOrMoreExpr{
								pos: position{line: 243, col: 54, offset: 9848},
								expr: &ruleRefExpr{
									pos:  position{line: 243, col: 54, offset: 9848},
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
			pos:  position{line: 250, col: 1, offset: 10167},
			expr: &choiceExpr{
				pos: position{line: 250, col: 20, offset: 10186},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 250, col: 20, offset: 10186},
						run: (*parser).callonLabeledListItem2,
						expr: &seqExpr{
							pos: position{line: 250, col: 20, offset: 10186},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 250, col: 20, offset: 10186},
									label: "term",
									expr: &ruleRefExpr{
										pos:  position{line: 250, col: 26, offset: 10192},
										name: "LabeledListItemTerm",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 250, col: 47, offset: 10213},
									name: "LabeledListItemSeparator",
								},
								&labeledExpr{
									pos:   position{line: 250, col: 72, offset: 10238},
									label: "description",
									expr: &ruleRefExpr{
										pos:  position{line: 250, col: 85, offset: 10251},
										name: "LabeledListItemDescription",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 252, col: 6, offset: 10378},
						run: (*parser).callonLabeledListItem9,
						expr: &seqExpr{
							pos: position{line: 252, col: 6, offset: 10378},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 252, col: 6, offset: 10378},
									label: "term",
									expr: &ruleRefExpr{
										pos:  position{line: 252, col: 12, offset: 10384},
										name: "LabeledListItemTerm",
									},
								},
								&litMatcher{
									pos:        position{line: 252, col: 34, offset: 10406},
									val:        "::",
									ignoreCase: false,
								},
								&zeroOrMoreExpr{
									pos: position{line: 252, col: 39, offset: 10411},
									expr: &ruleRefExpr{
										pos:  position{line: 252, col: 39, offset: 10411},
										name: "WS",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 252, col: 43, offset: 10415},
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
			pos:  position{line: 256, col: 1, offset: 10552},
			expr: &actionExpr{
				pos: position{line: 256, col: 24, offset: 10575},
				run: (*parser).callonLabeledListItemTerm1,
				expr: &labeledExpr{
					pos:   position{line: 256, col: 24, offset: 10575},
					label: "term",
					expr: &zeroOrMoreExpr{
						pos: position{line: 256, col: 29, offset: 10580},
						expr: &seqExpr{
							pos: position{line: 256, col: 30, offset: 10581},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 256, col: 30, offset: 10581},
									expr: &ruleRefExpr{
										pos:  position{line: 256, col: 31, offset: 10582},
										name: "NEWLINE",
									},
								},
								&notExpr{
									pos: position{line: 256, col: 39, offset: 10590},
									expr: &litMatcher{
										pos:        position{line: 256, col: 40, offset: 10591},
										val:        "::",
										ignoreCase: false,
									},
								},
								&anyMatcher{
									line: 256, col: 45, offset: 10596,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "LabeledListItemSeparator",
			pos:  position{line: 261, col: 1, offset: 10687},
			expr: &seqExpr{
				pos: position{line: 261, col: 30, offset: 10716},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 261, col: 30, offset: 10716},
						val:        "::",
						ignoreCase: false,
					},
					&oneOrMoreExpr{
						pos: position{line: 261, col: 35, offset: 10721},
						expr: &choiceExpr{
							pos: position{line: 261, col: 36, offset: 10722},
							alternatives: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 261, col: 36, offset: 10722},
									name: "WS",
								},
								&ruleRefExpr{
									pos:  position{line: 261, col: 41, offset: 10727},
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
			pos:  position{line: 263, col: 1, offset: 10738},
			expr: &actionExpr{
				pos: position{line: 263, col: 31, offset: 10768},
				run: (*parser).callonLabeledListItemDescription1,
				expr: &labeledExpr{
					pos:   position{line: 263, col: 31, offset: 10768},
					label: "elements",
					expr: &zeroOrMoreExpr{
						pos: position{line: 263, col: 40, offset: 10777},
						expr: &choiceExpr{
							pos: position{line: 263, col: 41, offset: 10778},
							alternatives: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 263, col: 41, offset: 10778},
									name: "ListParagraph",
								},
								&ruleRefExpr{
									pos:  position{line: 263, col: 57, offset: 10794},
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
			pos:  position{line: 272, col: 1, offset: 11144},
			expr: &actionExpr{
				pos: position{line: 272, col: 14, offset: 11157},
				run: (*parser).callonParagraph1,
				expr: &seqExpr{
					pos: position{line: 272, col: 14, offset: 11157},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 272, col: 14, offset: 11157},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 272, col: 25, offset: 11168},
								expr: &ruleRefExpr{
									pos:  position{line: 272, col: 26, offset: 11169},
									name: "ElementAttribute",
								},
							},
						},
						&notExpr{
							pos: position{line: 272, col: 45, offset: 11188},
							expr: &seqExpr{
								pos: position{line: 272, col: 47, offset: 11190},
								exprs: []interface{}{
									&oneOrMoreExpr{
										pos: position{line: 272, col: 47, offset: 11190},
										expr: &litMatcher{
											pos:        position{line: 272, col: 47, offset: 11190},
											val:        "=",
											ignoreCase: false,
										},
									},
									&oneOrMoreExpr{
										pos: position{line: 272, col: 52, offset: 11195},
										expr: &ruleRefExpr{
											pos:  position{line: 272, col: 52, offset: 11195},
											name: "WS",
										},
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 272, col: 57, offset: 11200},
							label: "lines",
							expr: &oneOrMoreExpr{
								pos: position{line: 272, col: 63, offset: 11206},
								expr: &seqExpr{
									pos: position{line: 272, col: 64, offset: 11207},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 272, col: 64, offset: 11207},
											name: "InlineContentWithTrailingSpaces",
										},
										&ruleRefExpr{
											pos:  position{line: 272, col: 96, offset: 11239},
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
			pos:  position{line: 278, col: 1, offset: 11529},
			expr: &actionExpr{
				pos: position{line: 278, col: 36, offset: 11564},
				run: (*parser).callonInlineContentWithTrailingSpaces1,
				expr: &seqExpr{
					pos: position{line: 278, col: 36, offset: 11564},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 278, col: 36, offset: 11564},
							expr: &ruleRefExpr{
								pos:  position{line: 278, col: 37, offset: 11565},
								name: "BlockDelimiter",
							},
						},
						&labeledExpr{
							pos:   position{line: 278, col: 52, offset: 11580},
							label: "elements",
							expr: &oneOrMoreExpr{
								pos: position{line: 278, col: 61, offset: 11589},
								expr: &seqExpr{
									pos: position{line: 278, col: 62, offset: 11590},
									exprs: []interface{}{
										&zeroOrMoreExpr{
											pos: position{line: 278, col: 62, offset: 11590},
											expr: &ruleRefExpr{
												pos:  position{line: 278, col: 62, offset: 11590},
												name: "WS",
											},
										},
										&notExpr{
											pos: position{line: 278, col: 66, offset: 11594},
											expr: &ruleRefExpr{
												pos:  position{line: 278, col: 67, offset: 11595},
												name: "InlineElementID",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 278, col: 83, offset: 11611},
											name: "InlineElement",
										},
										&zeroOrMoreExpr{
											pos: position{line: 278, col: 97, offset: 11625},
											expr: &ruleRefExpr{
												pos:  position{line: 278, col: 97, offset: 11625},
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
			pos:  position{line: 282, col: 1, offset: 11737},
			expr: &actionExpr{
				pos: position{line: 282, col: 18, offset: 11754},
				run: (*parser).callonInlineContent1,
				expr: &seqExpr{
					pos: position{line: 282, col: 18, offset: 11754},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 282, col: 18, offset: 11754},
							expr: &ruleRefExpr{
								pos:  position{line: 282, col: 19, offset: 11755},
								name: "BlockDelimiter",
							},
						},
						&labeledExpr{
							pos:   position{line: 282, col: 34, offset: 11770},
							label: "elements",
							expr: &oneOrMoreExpr{
								pos: position{line: 282, col: 43, offset: 11779},
								expr: &seqExpr{
									pos: position{line: 282, col: 44, offset: 11780},
									exprs: []interface{}{
										&zeroOrMoreExpr{
											pos: position{line: 282, col: 44, offset: 11780},
											expr: &ruleRefExpr{
												pos:  position{line: 282, col: 44, offset: 11780},
												name: "WS",
											},
										},
										&notExpr{
											pos: position{line: 282, col: 48, offset: 11784},
											expr: &ruleRefExpr{
												pos:  position{line: 282, col: 49, offset: 11785},
												name: "InlineElementID",
											},
										},
										&ruleRefExpr{
											pos:  position{line: 282, col: 65, offset: 11801},
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
			pos:  position{line: 286, col: 1, offset: 11923},
			expr: &choiceExpr{
				pos: position{line: 286, col: 19, offset: 11941},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 286, col: 19, offset: 11941},
						name: "CrossReference",
					},
					&ruleRefExpr{
						pos:  position{line: 286, col: 36, offset: 11958},
						name: "Passthrough",
					},
					&ruleRefExpr{
						pos:  position{line: 286, col: 50, offset: 11972},
						name: "InlineImage",
					},
					&ruleRefExpr{
						pos:  position{line: 286, col: 64, offset: 11986},
						name: "QuotedText",
					},
					&ruleRefExpr{
						pos:  position{line: 286, col: 77, offset: 11999},
						name: "Link",
					},
					&ruleRefExpr{
						pos:  position{line: 286, col: 84, offset: 12006},
						name: "DocumentAttributeSubstitution",
					},
					&ruleRefExpr{
						pos:  position{line: 286, col: 116, offset: 12038},
						name: "Characters",
					},
				},
			},
		},
		{
			name: "Admonition",
			pos:  position{line: 293, col: 1, offset: 12310},
			expr: &choiceExpr{
				pos: position{line: 293, col: 15, offset: 12324},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 293, col: 15, offset: 12324},
						run: (*parser).callonAdmonition2,
						expr: &seqExpr{
							pos: position{line: 293, col: 15, offset: 12324},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 293, col: 15, offset: 12324},
									label: "attributes",
									expr: &zeroOrMoreExpr{
										pos: position{line: 293, col: 26, offset: 12335},
										expr: &ruleRefExpr{
											pos:  position{line: 293, col: 27, offset: 12336},
											name: "ElementAttribute",
										},
									},
								},
								&notExpr{
									pos: position{line: 293, col: 46, offset: 12355},
									expr: &seqExpr{
										pos: position{line: 293, col: 48, offset: 12357},
										exprs: []interface{}{
											&oneOrMoreExpr{
												pos: position{line: 293, col: 48, offset: 12357},
												expr: &litMatcher{
													pos:        position{line: 293, col: 48, offset: 12357},
													val:        "=",
													ignoreCase: false,
												},
											},
											&oneOrMoreExpr{
												pos: position{line: 293, col: 53, offset: 12362},
												expr: &ruleRefExpr{
													pos:  position{line: 293, col: 53, offset: 12362},
													name: "WS",
												},
											},
										},
									},
								},
								&labeledExpr{
									pos:   position{line: 293, col: 58, offset: 12367},
									label: "t",
									expr: &ruleRefExpr{
										pos:  position{line: 293, col: 61, offset: 12370},
										name: "AdmonitionKind",
									},
								},
								&litMatcher{
									pos:        position{line: 293, col: 77, offset: 12386},
									val:        ": ",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 293, col: 82, offset: 12391},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 293, col: 91, offset: 12400},
										name: "AdmonitionParagraph",
									},
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 296, col: 1, offset: 12541},
						run: (*parser).callonAdmonition18,
						expr: &seqExpr{
							pos: position{line: 296, col: 1, offset: 12541},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 296, col: 1, offset: 12541},
									label: "attributes",
									expr: &zeroOrMoreExpr{
										pos: position{line: 296, col: 12, offset: 12552},
										expr: &ruleRefExpr{
											pos:  position{line: 296, col: 13, offset: 12553},
											name: "ElementAttribute",
										},
									},
								},
								&litMatcher{
									pos:        position{line: 296, col: 32, offset: 12572},
									val:        "[",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 296, col: 36, offset: 12576},
									label: "t",
									expr: &ruleRefExpr{
										pos:  position{line: 296, col: 39, offset: 12579},
										name: "AdmonitionKind",
									},
								},
								&litMatcher{
									pos:        position{line: 296, col: 55, offset: 12595},
									val:        "]",
									ignoreCase: false,
								},
								&zeroOrMoreExpr{
									pos: position{line: 296, col: 59, offset: 12599},
									expr: &ruleRefExpr{
										pos:  position{line: 296, col: 59, offset: 12599},
										name: "WS",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 296, col: 63, offset: 12603},
									name: "NEWLINE",
								},
								&labeledExpr{
									pos:   position{line: 296, col: 71, offset: 12611},
									label: "otherAttributes",
									expr: &zeroOrMoreExpr{
										pos: position{line: 296, col: 87, offset: 12627},
										expr: &ruleRefExpr{
											pos:  position{line: 296, col: 88, offset: 12628},
											name: "ElementAttribute",
										},
									},
								},
								&labeledExpr{
									pos:   position{line: 296, col: 107, offset: 12647},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 296, col: 116, offset: 12656},
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
			pos:  position{line: 300, col: 1, offset: 12835},
			expr: &actionExpr{
				pos: position{line: 300, col: 24, offset: 12858},
				run: (*parser).callonAdmonitionParagraph1,
				expr: &labeledExpr{
					pos:   position{line: 300, col: 24, offset: 12858},
					label: "lines",
					expr: &oneOrMoreExpr{
						pos: position{line: 300, col: 30, offset: 12864},
						expr: &seqExpr{
							pos: position{line: 300, col: 31, offset: 12865},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 300, col: 31, offset: 12865},
									name: "InlineContentWithTrailingSpaces",
								},
								&ruleRefExpr{
									pos:  position{line: 300, col: 63, offset: 12897},
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
			pos:  position{line: 304, col: 1, offset: 12971},
			expr: &choiceExpr{
				pos: position{line: 304, col: 19, offset: 12989},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 304, col: 19, offset: 12989},
						run: (*parser).callonAdmonitionKind2,
						expr: &litMatcher{
							pos:        position{line: 304, col: 19, offset: 12989},
							val:        "TIP",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 306, col: 5, offset: 13027},
						run: (*parser).callonAdmonitionKind4,
						expr: &litMatcher{
							pos:        position{line: 306, col: 5, offset: 13027},
							val:        "NOTE",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 308, col: 5, offset: 13067},
						run: (*parser).callonAdmonitionKind6,
						expr: &litMatcher{
							pos:        position{line: 308, col: 5, offset: 13067},
							val:        "IMPORTANT",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 310, col: 5, offset: 13117},
						run: (*parser).callonAdmonitionKind8,
						expr: &litMatcher{
							pos:        position{line: 310, col: 5, offset: 13117},
							val:        "WARNING",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 312, col: 5, offset: 13163},
						run: (*parser).callonAdmonitionKind10,
						expr: &litMatcher{
							pos:        position{line: 312, col: 5, offset: 13163},
							val:        "CAUTION",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "QuotedText",
			pos:  position{line: 319, col: 1, offset: 13447},
			expr: &choiceExpr{
				pos: position{line: 319, col: 15, offset: 13461},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 319, col: 15, offset: 13461},
						name: "BoldText",
					},
					&ruleRefExpr{
						pos:  position{line: 319, col: 26, offset: 13472},
						name: "ItalicText",
					},
					&ruleRefExpr{
						pos:  position{line: 319, col: 39, offset: 13485},
						name: "MonospaceText",
					},
					&ruleRefExpr{
						pos:  position{line: 320, col: 13, offset: 13513},
						name: "EscapedBoldText",
					},
					&ruleRefExpr{
						pos:  position{line: 320, col: 31, offset: 13531},
						name: "EscapedItalicText",
					},
					&ruleRefExpr{
						pos:  position{line: 320, col: 51, offset: 13551},
						name: "EscapedMonospaceText",
					},
				},
			},
		},
		{
			name: "BoldText",
			pos:  position{line: 322, col: 1, offset: 13573},
			expr: &choiceExpr{
				pos: position{line: 322, col: 13, offset: 13585},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 322, col: 13, offset: 13585},
						run: (*parser).callonBoldText2,
						expr: &seqExpr{
							pos: position{line: 322, col: 13, offset: 13585},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 322, col: 13, offset: 13585},
									expr: &litMatcher{
										pos:        position{line: 322, col: 14, offset: 13586},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 322, col: 19, offset: 13591},
									val:        "**",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 322, col: 24, offset: 13596},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 322, col: 33, offset: 13605},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 322, col: 52, offset: 13624},
									val:        "**",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 324, col: 5, offset: 13749},
						run: (*parser).callonBoldText10,
						expr: &seqExpr{
							pos: position{line: 324, col: 5, offset: 13749},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 324, col: 5, offset: 13749},
									expr: &litMatcher{
										pos:        position{line: 324, col: 6, offset: 13750},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 324, col: 11, offset: 13755},
									val:        "**",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 324, col: 16, offset: 13760},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 324, col: 25, offset: 13769},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 324, col: 44, offset: 13788},
									val:        "*",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 327, col: 5, offset: 13953},
						run: (*parser).callonBoldText18,
						expr: &seqExpr{
							pos: position{line: 327, col: 5, offset: 13953},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 327, col: 5, offset: 13953},
									expr: &litMatcher{
										pos:        position{line: 327, col: 6, offset: 13954},
										val:        "\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 327, col: 10, offset: 13958},
									val:        "*",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 327, col: 14, offset: 13962},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 327, col: 23, offset: 13971},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 327, col: 42, offset: 13990},
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
			pos:  position{line: 331, col: 1, offset: 14090},
			expr: &choiceExpr{
				pos: position{line: 331, col: 20, offset: 14109},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 331, col: 20, offset: 14109},
						run: (*parser).callonEscapedBoldText2,
						expr: &seqExpr{
							pos: position{line: 331, col: 20, offset: 14109},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 331, col: 20, offset: 14109},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 331, col: 33, offset: 14122},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 331, col: 33, offset: 14122},
												val:        "\\\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 331, col: 38, offset: 14127},
												expr: &litMatcher{
													pos:        position{line: 331, col: 38, offset: 14127},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 331, col: 44, offset: 14133},
									val:        "**",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 331, col: 49, offset: 14138},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 331, col: 58, offset: 14147},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 331, col: 77, offset: 14166},
									val:        "**",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 333, col: 5, offset: 14321},
						run: (*parser).callonEscapedBoldText13,
						expr: &seqExpr{
							pos: position{line: 333, col: 5, offset: 14321},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 333, col: 5, offset: 14321},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 333, col: 18, offset: 14334},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 333, col: 18, offset: 14334},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 333, col: 22, offset: 14338},
												expr: &litMatcher{
													pos:        position{line: 333, col: 22, offset: 14338},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 333, col: 28, offset: 14344},
									val:        "**",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 333, col: 33, offset: 14349},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 333, col: 42, offset: 14358},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 333, col: 61, offset: 14377},
									val:        "*",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 336, col: 5, offset: 14571},
						run: (*parser).callonEscapedBoldText24,
						expr: &seqExpr{
							pos: position{line: 336, col: 5, offset: 14571},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 336, col: 5, offset: 14571},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 336, col: 18, offset: 14584},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 336, col: 18, offset: 14584},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 336, col: 22, offset: 14588},
												expr: &litMatcher{
													pos:        position{line: 336, col: 22, offset: 14588},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 336, col: 28, offset: 14594},
									val:        "*",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 336, col: 32, offset: 14598},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 336, col: 41, offset: 14607},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 336, col: 60, offset: 14626},
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
			pos:  position{line: 340, col: 1, offset: 14778},
			expr: &choiceExpr{
				pos: position{line: 340, col: 15, offset: 14792},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 340, col: 15, offset: 14792},
						run: (*parser).callonItalicText2,
						expr: &seqExpr{
							pos: position{line: 340, col: 15, offset: 14792},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 340, col: 15, offset: 14792},
									expr: &litMatcher{
										pos:        position{line: 340, col: 16, offset: 14793},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 340, col: 21, offset: 14798},
									val:        "__",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 340, col: 26, offset: 14803},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 340, col: 35, offset: 14812},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 340, col: 54, offset: 14831},
									val:        "__",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 342, col: 5, offset: 14912},
						run: (*parser).callonItalicText10,
						expr: &seqExpr{
							pos: position{line: 342, col: 5, offset: 14912},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 342, col: 5, offset: 14912},
									expr: &litMatcher{
										pos:        position{line: 342, col: 6, offset: 14913},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 342, col: 11, offset: 14918},
									val:        "__",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 342, col: 16, offset: 14923},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 342, col: 25, offset: 14932},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 342, col: 44, offset: 14951},
									val:        "_",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 345, col: 5, offset: 15118},
						run: (*parser).callonItalicText18,
						expr: &seqExpr{
							pos: position{line: 345, col: 5, offset: 15118},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 345, col: 5, offset: 15118},
									expr: &litMatcher{
										pos:        position{line: 345, col: 6, offset: 15119},
										val:        "\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 345, col: 10, offset: 15123},
									val:        "_",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 345, col: 14, offset: 15127},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 345, col: 23, offset: 15136},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 345, col: 42, offset: 15155},
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
			pos:  position{line: 349, col: 1, offset: 15234},
			expr: &choiceExpr{
				pos: position{line: 349, col: 22, offset: 15255},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 349, col: 22, offset: 15255},
						run: (*parser).callonEscapedItalicText2,
						expr: &seqExpr{
							pos: position{line: 349, col: 22, offset: 15255},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 349, col: 22, offset: 15255},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 349, col: 35, offset: 15268},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 349, col: 35, offset: 15268},
												val:        "\\\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 349, col: 40, offset: 15273},
												expr: &litMatcher{
													pos:        position{line: 349, col: 40, offset: 15273},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 349, col: 46, offset: 15279},
									val:        "__",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 349, col: 51, offset: 15284},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 349, col: 60, offset: 15293},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 349, col: 79, offset: 15312},
									val:        "__",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 351, col: 5, offset: 15467},
						run: (*parser).callonEscapedItalicText13,
						expr: &seqExpr{
							pos: position{line: 351, col: 5, offset: 15467},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 351, col: 5, offset: 15467},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 351, col: 18, offset: 15480},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 351, col: 18, offset: 15480},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 351, col: 22, offset: 15484},
												expr: &litMatcher{
													pos:        position{line: 351, col: 22, offset: 15484},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 351, col: 28, offset: 15490},
									val:        "__",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 351, col: 33, offset: 15495},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 351, col: 42, offset: 15504},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 351, col: 61, offset: 15523},
									val:        "_",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 354, col: 5, offset: 15717},
						run: (*parser).callonEscapedItalicText24,
						expr: &seqExpr{
							pos: position{line: 354, col: 5, offset: 15717},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 354, col: 5, offset: 15717},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 354, col: 18, offset: 15730},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 354, col: 18, offset: 15730},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 354, col: 22, offset: 15734},
												expr: &litMatcher{
													pos:        position{line: 354, col: 22, offset: 15734},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 354, col: 28, offset: 15740},
									val:        "_",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 354, col: 32, offset: 15744},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 354, col: 41, offset: 15753},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 354, col: 60, offset: 15772},
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
			pos:  position{line: 358, col: 1, offset: 15924},
			expr: &choiceExpr{
				pos: position{line: 358, col: 18, offset: 15941},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 358, col: 18, offset: 15941},
						run: (*parser).callonMonospaceText2,
						expr: &seqExpr{
							pos: position{line: 358, col: 18, offset: 15941},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 358, col: 18, offset: 15941},
									expr: &litMatcher{
										pos:        position{line: 358, col: 19, offset: 15942},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 358, col: 24, offset: 15947},
									val:        "``",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 358, col: 29, offset: 15952},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 358, col: 38, offset: 15961},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 358, col: 57, offset: 15980},
									val:        "``",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 360, col: 5, offset: 16110},
						run: (*parser).callonMonospaceText10,
						expr: &seqExpr{
							pos: position{line: 360, col: 5, offset: 16110},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 360, col: 5, offset: 16110},
									expr: &litMatcher{
										pos:        position{line: 360, col: 6, offset: 16111},
										val:        "\\\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 360, col: 11, offset: 16116},
									val:        "``",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 360, col: 16, offset: 16121},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 360, col: 25, offset: 16130},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 360, col: 44, offset: 16149},
									val:        "`",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 363, col: 5, offset: 16319},
						run: (*parser).callonMonospaceText18,
						expr: &seqExpr{
							pos: position{line: 363, col: 5, offset: 16319},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 363, col: 5, offset: 16319},
									expr: &litMatcher{
										pos:        position{line: 363, col: 6, offset: 16320},
										val:        "\\",
										ignoreCase: false,
									},
								},
								&litMatcher{
									pos:        position{line: 363, col: 10, offset: 16324},
									val:        "`",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 363, col: 14, offset: 16328},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 363, col: 23, offset: 16337},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 363, col: 42, offset: 16356},
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
			pos:  position{line: 367, col: 1, offset: 16483},
			expr: &choiceExpr{
				pos: position{line: 367, col: 25, offset: 16507},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 367, col: 25, offset: 16507},
						run: (*parser).callonEscapedMonospaceText2,
						expr: &seqExpr{
							pos: position{line: 367, col: 25, offset: 16507},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 367, col: 25, offset: 16507},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 367, col: 38, offset: 16520},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 367, col: 38, offset: 16520},
												val:        "\\\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 367, col: 43, offset: 16525},
												expr: &litMatcher{
													pos:        position{line: 367, col: 43, offset: 16525},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 367, col: 49, offset: 16531},
									val:        "``",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 367, col: 54, offset: 16536},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 367, col: 63, offset: 16545},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 367, col: 82, offset: 16564},
									val:        "``",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 369, col: 5, offset: 16719},
						run: (*parser).callonEscapedMonospaceText13,
						expr: &seqExpr{
							pos: position{line: 369, col: 5, offset: 16719},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 369, col: 5, offset: 16719},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 369, col: 18, offset: 16732},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 369, col: 18, offset: 16732},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 369, col: 22, offset: 16736},
												expr: &litMatcher{
													pos:        position{line: 369, col: 22, offset: 16736},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 369, col: 28, offset: 16742},
									val:        "``",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 369, col: 33, offset: 16747},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 369, col: 42, offset: 16756},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 369, col: 61, offset: 16775},
									val:        "`",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 372, col: 5, offset: 16969},
						run: (*parser).callonEscapedMonospaceText24,
						expr: &seqExpr{
							pos: position{line: 372, col: 5, offset: 16969},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 372, col: 5, offset: 16969},
									label: "backslashes",
									expr: &seqExpr{
										pos: position{line: 372, col: 18, offset: 16982},
										exprs: []interface{}{
											&litMatcher{
												pos:        position{line: 372, col: 18, offset: 16982},
												val:        "\\",
												ignoreCase: false,
											},
											&zeroOrMoreExpr{
												pos: position{line: 372, col: 22, offset: 16986},
												expr: &litMatcher{
													pos:        position{line: 372, col: 22, offset: 16986},
													val:        "\\",
													ignoreCase: false,
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 372, col: 28, offset: 16992},
									val:        "`",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 372, col: 32, offset: 16996},
									label: "content",
									expr: &ruleRefExpr{
										pos:  position{line: 372, col: 41, offset: 17005},
										name: "QuotedTextContent",
									},
								},
								&litMatcher{
									pos:        position{line: 372, col: 60, offset: 17024},
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
			pos:  position{line: 376, col: 1, offset: 17176},
			expr: &seqExpr{
				pos: position{line: 376, col: 22, offset: 17197},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 376, col: 22, offset: 17197},
						name: "QuotedTextContentElement",
					},
					&zeroOrMoreExpr{
						pos: position{line: 376, col: 47, offset: 17222},
						expr: &seqExpr{
							pos: position{line: 376, col: 48, offset: 17223},
							exprs: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 376, col: 48, offset: 17223},
									expr: &ruleRefExpr{
										pos:  position{line: 376, col: 48, offset: 17223},
										name: "WS",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 376, col: 52, offset: 17227},
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
			pos:  position{line: 378, col: 1, offset: 17255},
			expr: &choiceExpr{
				pos: position{line: 378, col: 29, offset: 17283},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 378, col: 29, offset: 17283},
						name: "QuotedText",
					},
					&ruleRefExpr{
						pos:  position{line: 378, col: 42, offset: 17296},
						name: "QuotedTextCharacters",
					},
					&ruleRefExpr{
						pos:  position{line: 378, col: 65, offset: 17319},
						name: "CharactersWithQuotePunctuation",
					},
				},
			},
		},
		{
			name: "QuotedTextCharacters",
			pos:  position{line: 380, col: 1, offset: 17454},
			expr: &oneOrMoreExpr{
				pos: position{line: 380, col: 25, offset: 17478},
				expr: &seqExpr{
					pos: position{line: 380, col: 26, offset: 17479},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 380, col: 26, offset: 17479},
							expr: &ruleRefExpr{
								pos:  position{line: 380, col: 27, offset: 17480},
								name: "NEWLINE",
							},
						},
						&notExpr{
							pos: position{line: 380, col: 35, offset: 17488},
							expr: &ruleRefExpr{
								pos:  position{line: 380, col: 36, offset: 17489},
								name: "WS",
							},
						},
						&notExpr{
							pos: position{line: 380, col: 39, offset: 17492},
							expr: &litMatcher{
								pos:        position{line: 380, col: 40, offset: 17493},
								val:        "*",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 380, col: 44, offset: 17497},
							expr: &litMatcher{
								pos:        position{line: 380, col: 45, offset: 17498},
								val:        "_",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 380, col: 49, offset: 17502},
							expr: &litMatcher{
								pos:        position{line: 380, col: 50, offset: 17503},
								val:        "`",
								ignoreCase: false,
							},
						},
						&anyMatcher{
							line: 380, col: 54, offset: 17507,
						},
					},
				},
			},
		},
		{
			name: "CharactersWithQuotePunctuation",
			pos:  position{line: 382, col: 1, offset: 17550},
			expr: &actionExpr{
				pos: position{line: 382, col: 35, offset: 17584},
				run: (*parser).callonCharactersWithQuotePunctuation1,
				expr: &oneOrMoreExpr{
					pos: position{line: 382, col: 35, offset: 17584},
					expr: &seqExpr{
						pos: position{line: 382, col: 36, offset: 17585},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 382, col: 36, offset: 17585},
								expr: &ruleRefExpr{
									pos:  position{line: 382, col: 37, offset: 17586},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 382, col: 45, offset: 17594},
								expr: &ruleRefExpr{
									pos:  position{line: 382, col: 46, offset: 17595},
									name: "WS",
								},
							},
							&anyMatcher{
								line: 382, col: 50, offset: 17599,
							},
						},
					},
				},
			},
		},
		{
			name: "UnbalancedQuotePunctuation",
			pos:  position{line: 387, col: 1, offset: 17844},
			expr: &choiceExpr{
				pos: position{line: 387, col: 31, offset: 17874},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 387, col: 31, offset: 17874},
						val:        "*",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 387, col: 37, offset: 17880},
						val:        "_",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 387, col: 43, offset: 17886},
						val:        "`",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "Passthrough",
			pos:  position{line: 392, col: 1, offset: 17998},
			expr: &choiceExpr{
				pos: position{line: 392, col: 16, offset: 18013},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 392, col: 16, offset: 18013},
						name: "TriplePlusPassthrough",
					},
					&ruleRefExpr{
						pos:  position{line: 392, col: 40, offset: 18037},
						name: "SinglePlusPassthrough",
					},
					&ruleRefExpr{
						pos:  position{line: 392, col: 64, offset: 18061},
						name: "PassthroughMacro",
					},
				},
			},
		},
		{
			name: "SinglePlusPassthrough",
			pos:  position{line: 394, col: 1, offset: 18079},
			expr: &actionExpr{
				pos: position{line: 394, col: 26, offset: 18104},
				run: (*parser).callonSinglePlusPassthrough1,
				expr: &seqExpr{
					pos: position{line: 394, col: 26, offset: 18104},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 394, col: 26, offset: 18104},
							val:        "+",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 394, col: 30, offset: 18108},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 394, col: 38, offset: 18116},
								expr: &seqExpr{
									pos: position{line: 394, col: 39, offset: 18117},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 394, col: 39, offset: 18117},
											expr: &ruleRefExpr{
												pos:  position{line: 394, col: 40, offset: 18118},
												name: "NEWLINE",
											},
										},
										&notExpr{
											pos: position{line: 394, col: 48, offset: 18126},
											expr: &litMatcher{
												pos:        position{line: 394, col: 49, offset: 18127},
												val:        "+",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 394, col: 53, offset: 18131,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 394, col: 57, offset: 18135},
							val:        "+",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TriplePlusPassthrough",
			pos:  position{line: 398, col: 1, offset: 18230},
			expr: &actionExpr{
				pos: position{line: 398, col: 26, offset: 18255},
				run: (*parser).callonTriplePlusPassthrough1,
				expr: &seqExpr{
					pos: position{line: 398, col: 26, offset: 18255},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 398, col: 26, offset: 18255},
							val:        "+++",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 398, col: 32, offset: 18261},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 398, col: 40, offset: 18269},
								expr: &seqExpr{
									pos: position{line: 398, col: 41, offset: 18270},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 398, col: 41, offset: 18270},
											expr: &litMatcher{
												pos:        position{line: 398, col: 42, offset: 18271},
												val:        "+++",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 398, col: 48, offset: 18277,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 398, col: 52, offset: 18281},
							val:        "+++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "PassthroughMacro",
			pos:  position{line: 402, col: 1, offset: 18378},
			expr: &choiceExpr{
				pos: position{line: 402, col: 21, offset: 18398},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 402, col: 21, offset: 18398},
						run: (*parser).callonPassthroughMacro2,
						expr: &seqExpr{
							pos: position{line: 402, col: 21, offset: 18398},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 402, col: 21, offset: 18398},
									val:        "pass:[",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 402, col: 30, offset: 18407},
									label: "content",
									expr: &zeroOrMoreExpr{
										pos: position{line: 402, col: 38, offset: 18415},
										expr: &ruleRefExpr{
											pos:  position{line: 402, col: 39, offset: 18416},
											name: "PassthroughMacroCharacter",
										},
									},
								},
								&litMatcher{
									pos:        position{line: 402, col: 67, offset: 18444},
									val:        "]",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 404, col: 5, offset: 18535},
						run: (*parser).callonPassthroughMacro9,
						expr: &seqExpr{
							pos: position{line: 404, col: 5, offset: 18535},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 404, col: 5, offset: 18535},
									val:        "pass:q[",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 404, col: 15, offset: 18545},
									label: "content",
									expr: &zeroOrMoreExpr{
										pos: position{line: 404, col: 23, offset: 18553},
										expr: &choiceExpr{
											pos: position{line: 404, col: 24, offset: 18554},
											alternatives: []interface{}{
												&ruleRefExpr{
													pos:  position{line: 404, col: 24, offset: 18554},
													name: "QuotedText",
												},
												&ruleRefExpr{
													pos:  position{line: 404, col: 37, offset: 18567},
													name: "PassthroughMacroCharacter",
												},
											},
										},
									},
								},
								&litMatcher{
									pos:        position{line: 404, col: 65, offset: 18595},
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
			pos:  position{line: 408, col: 1, offset: 18685},
			expr: &seqExpr{
				pos: position{line: 408, col: 31, offset: 18715},
				exprs: []interface{}{
					&notExpr{
						pos: position{line: 408, col: 31, offset: 18715},
						expr: &litMatcher{
							pos:        position{line: 408, col: 32, offset: 18716},
							val:        "]",
							ignoreCase: false,
						},
					},
					&anyMatcher{
						line: 408, col: 36, offset: 18720,
					},
				},
			},
		},
		{
			name: "CrossReference",
			pos:  position{line: 413, col: 1, offset: 18836},
			expr: &actionExpr{
				pos: position{line: 413, col: 19, offset: 18854},
				run: (*parser).callonCrossReference1,
				expr: &seqExpr{
					pos: position{line: 413, col: 19, offset: 18854},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 413, col: 19, offset: 18854},
							val:        "<<",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 413, col: 24, offset: 18859},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 413, col: 28, offset: 18863},
								name: "ID",
							},
						},
						&litMatcher{
							pos:        position{line: 413, col: 32, offset: 18867},
							val:        ">>",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Link",
			pos:  position{line: 420, col: 1, offset: 19026},
			expr: &choiceExpr{
				pos: position{line: 420, col: 9, offset: 19034},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 420, col: 9, offset: 19034},
						name: "RelativeLink",
					},
					&ruleRefExpr{
						pos:  position{line: 420, col: 24, offset: 19049},
						name: "ExternalLink",
					},
				},
			},
		},
		{
			name: "ExternalLink",
			pos:  position{line: 422, col: 1, offset: 19064},
			expr: &actionExpr{
				pos: position{line: 422, col: 17, offset: 19080},
				run: (*parser).callonExternalLink1,
				expr: &seqExpr{
					pos: position{line: 422, col: 17, offset: 19080},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 422, col: 17, offset: 19080},
							label: "url",
							expr: &seqExpr{
								pos: position{line: 422, col: 22, offset: 19085},
								exprs: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 422, col: 22, offset: 19085},
										name: "URL_SCHEME",
									},
									&ruleRefExpr{
										pos:  position{line: 422, col: 33, offset: 19096},
										name: "URL",
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 422, col: 38, offset: 19101},
							label: "text",
							expr: &zeroOrOneExpr{
								pos: position{line: 422, col: 43, offset: 19106},
								expr: &seqExpr{
									pos: position{line: 422, col: 44, offset: 19107},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 422, col: 44, offset: 19107},
											val:        "[",
											ignoreCase: false,
										},
										&zeroOrMoreExpr{
											pos: position{line: 422, col: 48, offset: 19111},
											expr: &ruleRefExpr{
												pos:  position{line: 422, col: 49, offset: 19112},
												name: "URL_TEXT",
											},
										},
										&litMatcher{
											pos:        position{line: 422, col: 60, offset: 19123},
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
			pos:  position{line: 429, col: 1, offset: 19284},
			expr: &actionExpr{
				pos: position{line: 429, col: 17, offset: 19300},
				run: (*parser).callonRelativeLink1,
				expr: &seqExpr{
					pos: position{line: 429, col: 17, offset: 19300},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 429, col: 17, offset: 19300},
							val:        "link:",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 429, col: 25, offset: 19308},
							label: "url",
							expr: &seqExpr{
								pos: position{line: 429, col: 30, offset: 19313},
								exprs: []interface{}{
									&zeroOrOneExpr{
										pos: position{line: 429, col: 30, offset: 19313},
										expr: &ruleRefExpr{
											pos:  position{line: 429, col: 30, offset: 19313},
											name: "URL_SCHEME",
										},
									},
									&ruleRefExpr{
										pos:  position{line: 429, col: 42, offset: 19325},
										name: "URL",
									},
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 429, col: 47, offset: 19330},
							label: "text",
							expr: &seqExpr{
								pos: position{line: 429, col: 53, offset: 19336},
								exprs: []interface{}{
									&litMatcher{
										pos:        position{line: 429, col: 53, offset: 19336},
										val:        "[",
										ignoreCase: false,
									},
									&zeroOrMoreExpr{
										pos: position{line: 429, col: 57, offset: 19340},
										expr: &ruleRefExpr{
											pos:  position{line: 429, col: 58, offset: 19341},
											name: "URL_TEXT",
										},
									},
									&litMatcher{
										pos:        position{line: 429, col: 69, offset: 19352},
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
			pos:  position{line: 439, col: 1, offset: 19614},
			expr: &actionExpr{
				pos: position{line: 439, col: 15, offset: 19628},
				run: (*parser).callonBlockImage1,
				expr: &seqExpr{
					pos: position{line: 439, col: 15, offset: 19628},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 439, col: 15, offset: 19628},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 439, col: 26, offset: 19639},
								expr: &ruleRefExpr{
									pos:  position{line: 439, col: 27, offset: 19640},
									name: "ElementAttribute",
								},
							},
						},
						&labeledExpr{
							pos:   position{line: 439, col: 46, offset: 19659},
							label: "image",
							expr: &ruleRefExpr{
								pos:  position{line: 439, col: 52, offset: 19665},
								name: "BlockImageMacro",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 439, col: 69, offset: 19682},
							expr: &ruleRefExpr{
								pos:  position{line: 439, col: 69, offset: 19682},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 439, col: 73, offset: 19686},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "BlockImageMacro",
			pos:  position{line: 444, col: 1, offset: 19845},
			expr: &actionExpr{
				pos: position{line: 444, col: 20, offset: 19864},
				run: (*parser).callonBlockImageMacro1,
				expr: &seqExpr{
					pos: position{line: 444, col: 20, offset: 19864},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 444, col: 20, offset: 19864},
							val:        "image::",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 444, col: 30, offset: 19874},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 444, col: 36, offset: 19880},
								name: "URL",
							},
						},
						&litMatcher{
							pos:        position{line: 444, col: 41, offset: 19885},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 444, col: 45, offset: 19889},
							label: "attributes",
							expr: &zeroOrOneExpr{
								pos: position{line: 444, col: 57, offset: 19901},
								expr: &ruleRefExpr{
									pos:  position{line: 444, col: 57, offset: 19901},
									name: "URL_TEXT",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 444, col: 68, offset: 19912},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "InlineImage",
			pos:  position{line: 448, col: 1, offset: 19979},
			expr: &actionExpr{
				pos: position{line: 448, col: 16, offset: 19994},
				run: (*parser).callonInlineImage1,
				expr: &labeledExpr{
					pos:   position{line: 448, col: 16, offset: 19994},
					label: "image",
					expr: &ruleRefExpr{
						pos:  position{line: 448, col: 22, offset: 20000},
						name: "InlineImageMacro",
					},
				},
			},
		},
		{
			name: "InlineImageMacro",
			pos:  position{line: 453, col: 1, offset: 20145},
			expr: &actionExpr{
				pos: position{line: 453, col: 21, offset: 20165},
				run: (*parser).callonInlineImageMacro1,
				expr: &seqExpr{
					pos: position{line: 453, col: 21, offset: 20165},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 453, col: 21, offset: 20165},
							val:        "image:",
							ignoreCase: false,
						},
						&notExpr{
							pos: position{line: 453, col: 30, offset: 20174},
							expr: &litMatcher{
								pos:        position{line: 453, col: 31, offset: 20175},
								val:        ":",
								ignoreCase: false,
							},
						},
						&labeledExpr{
							pos:   position{line: 453, col: 35, offset: 20179},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 453, col: 41, offset: 20185},
								name: "URL",
							},
						},
						&litMatcher{
							pos:        position{line: 453, col: 46, offset: 20190},
							val:        "[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 453, col: 50, offset: 20194},
							label: "attributes",
							expr: &zeroOrOneExpr{
								pos: position{line: 453, col: 62, offset: 20206},
								expr: &ruleRefExpr{
									pos:  position{line: 453, col: 62, offset: 20206},
									name: "URL_TEXT",
								},
							},
						},
						&litMatcher{
							pos:        position{line: 453, col: 73, offset: 20217},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "DelimitedBlock",
			pos:  position{line: 460, col: 1, offset: 20547},
			expr: &choiceExpr{
				pos: position{line: 460, col: 19, offset: 20565},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 460, col: 19, offset: 20565},
						name: "FencedBlock",
					},
					&ruleRefExpr{
						pos:  position{line: 460, col: 33, offset: 20579},
						name: "ListingBlock",
					},
					&ruleRefExpr{
						pos:  position{line: 460, col: 48, offset: 20594},
						name: "ExampleBlock",
					},
				},
			},
		},
		{
			name: "BlockDelimiter",
			pos:  position{line: 462, col: 1, offset: 20608},
			expr: &choiceExpr{
				pos: position{line: 462, col: 19, offset: 20626},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 462, col: 19, offset: 20626},
						name: "LiteralBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 462, col: 43, offset: 20650},
						name: "FencedBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 462, col: 66, offset: 20673},
						name: "ListingBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 462, col: 90, offset: 20697},
						name: "ExampleBlockDelimiter",
					},
				},
			},
		},
		{
			name: "FencedBlockDelimiter",
			pos:  position{line: 464, col: 1, offset: 20720},
			expr: &litMatcher{
				pos:        position{line: 464, col: 25, offset: 20744},
				val:        "```",
				ignoreCase: false,
			},
		},
		{
			name: "FencedBlock",
			pos:  position{line: 466, col: 1, offset: 20751},
			expr: &actionExpr{
				pos: position{line: 466, col: 16, offset: 20766},
				run: (*parser).callonFencedBlock1,
				expr: &seqExpr{
					pos: position{line: 466, col: 16, offset: 20766},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 466, col: 16, offset: 20766},
							name: "FencedBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 466, col: 37, offset: 20787},
							expr: &ruleRefExpr{
								pos:  position{line: 466, col: 37, offset: 20787},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 466, col: 41, offset: 20791},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 466, col: 49, offset: 20799},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 466, col: 57, offset: 20807},
								expr: &seqExpr{
									pos: position{line: 466, col: 58, offset: 20808},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 466, col: 58, offset: 20808},
											expr: &ruleRefExpr{
												pos:  position{line: 466, col: 59, offset: 20809},
												name: "FencedBlockDelimiter",
											},
										},
										&anyMatcher{
											line: 466, col: 80, offset: 20830,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 466, col: 84, offset: 20834},
							name: "FencedBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 466, col: 105, offset: 20855},
							expr: &ruleRefExpr{
								pos:  position{line: 466, col: 105, offset: 20855},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 466, col: 109, offset: 20859},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ListingBlockDelimiter",
			pos:  position{line: 470, col: 1, offset: 20952},
			expr: &litMatcher{
				pos:        position{line: 470, col: 26, offset: 20977},
				val:        "----",
				ignoreCase: false,
			},
		},
		{
			name: "ListingBlock",
			pos:  position{line: 472, col: 1, offset: 20985},
			expr: &actionExpr{
				pos: position{line: 472, col: 17, offset: 21001},
				run: (*parser).callonListingBlock1,
				expr: &seqExpr{
					pos: position{line: 472, col: 17, offset: 21001},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 472, col: 17, offset: 21001},
							name: "ListingBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 472, col: 39, offset: 21023},
							expr: &ruleRefExpr{
								pos:  position{line: 472, col: 39, offset: 21023},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 472, col: 43, offset: 21027},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 472, col: 51, offset: 21035},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 472, col: 59, offset: 21043},
								expr: &seqExpr{
									pos: position{line: 472, col: 60, offset: 21044},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 472, col: 60, offset: 21044},
											expr: &ruleRefExpr{
												pos:  position{line: 472, col: 61, offset: 21045},
												name: "ListingBlockDelimiter",
											},
										},
										&anyMatcher{
											line: 472, col: 83, offset: 21067,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 472, col: 87, offset: 21071},
							name: "ListingBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 472, col: 109, offset: 21093},
							expr: &ruleRefExpr{
								pos:  position{line: 472, col: 109, offset: 21093},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 472, col: 113, offset: 21097},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ExampleBlockDelimiter",
			pos:  position{line: 476, col: 1, offset: 21191},
			expr: &litMatcher{
				pos:        position{line: 476, col: 26, offset: 21216},
				val:        "====",
				ignoreCase: false,
			},
		},
		{
			name: "ExampleBlock",
			pos:  position{line: 478, col: 1, offset: 21224},
			expr: &actionExpr{
				pos: position{line: 478, col: 17, offset: 21240},
				run: (*parser).callonExampleBlock1,
				expr: &seqExpr{
					pos: position{line: 478, col: 17, offset: 21240},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 478, col: 17, offset: 21240},
							label: "attributes",
							expr: &zeroOrMoreExpr{
								pos: position{line: 478, col: 28, offset: 21251},
								expr: &ruleRefExpr{
									pos:  position{line: 478, col: 29, offset: 21252},
									name: "ElementAttribute",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 478, col: 48, offset: 21271},
							name: "ExampleBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 478, col: 70, offset: 21293},
							expr: &ruleRefExpr{
								pos:  position{line: 478, col: 70, offset: 21293},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 478, col: 74, offset: 21297},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 478, col: 82, offset: 21305},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 478, col: 90, offset: 21313},
								expr: &choiceExpr{
									pos: position{line: 478, col: 91, offset: 21314},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 478, col: 91, offset: 21314},
											name: "List",
										},
										&ruleRefExpr{
											pos:  position{line: 478, col: 98, offset: 21321},
											name: "Paragraph",
										},
										&ruleRefExpr{
											pos:  position{line: 478, col: 110, offset: 21333},
											name: "BlankLine",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 478, col: 123, offset: 21346},
							name: "ExampleBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 478, col: 145, offset: 21368},
							expr: &ruleRefExpr{
								pos:  position{line: 478, col: 145, offset: 21368},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 478, col: 149, offset: 21372},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "LiteralBlock",
			pos:  position{line: 485, col: 1, offset: 21756},
			expr: &choiceExpr{
				pos: position{line: 485, col: 17, offset: 21772},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 485, col: 17, offset: 21772},
						name: "ParagraphWithSpaces",
					},
					&ruleRefExpr{
						pos:  position{line: 485, col: 39, offset: 21794},
						name: "ParagraphWithLiteralBlockDelimiter",
					},
					&ruleRefExpr{
						pos:  position{line: 485, col: 76, offset: 21831},
						name: "ParagraphWithLiteralAttribute",
					},
				},
			},
		},
		{
			name: "ParagraphWithSpaces",
			pos:  position{line: 488, col: 1, offset: 21926},
			expr: &actionExpr{
				pos: position{line: 488, col: 24, offset: 21949},
				run: (*parser).callonParagraphWithSpaces1,
				expr: &seqExpr{
					pos: position{line: 488, col: 24, offset: 21949},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 488, col: 24, offset: 21949},
							label: "spaces",
							expr: &oneOrMoreExpr{
								pos: position{line: 488, col: 32, offset: 21957},
								expr: &ruleRefExpr{
									pos:  position{line: 488, col: 32, offset: 21957},
									name: "WS",
								},
							},
						},
						&notExpr{
							pos: position{line: 488, col: 37, offset: 21962},
							expr: &ruleRefExpr{
								pos:  position{line: 488, col: 38, offset: 21963},
								name: "NEWLINE",
							},
						},
						&labeledExpr{
							pos:   position{line: 488, col: 46, offset: 21971},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 488, col: 55, offset: 21980},
								name: "LiteralBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 488, col: 76, offset: 22001},
							name: "EndOfLiteralBlock",
						},
					},
				},
			},
		},
		{
			name: "LiteralBlockContent",
			pos:  position{line: 493, col: 1, offset: 22182},
			expr: &actionExpr{
				pos: position{line: 493, col: 24, offset: 22205},
				run: (*parser).callonLiteralBlockContent1,
				expr: &labeledExpr{
					pos:   position{line: 493, col: 24, offset: 22205},
					label: "content",
					expr: &oneOrMoreExpr{
						pos: position{line: 493, col: 32, offset: 22213},
						expr: &seqExpr{
							pos: position{line: 493, col: 33, offset: 22214},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 493, col: 33, offset: 22214},
									expr: &seqExpr{
										pos: position{line: 493, col: 35, offset: 22216},
										exprs: []interface{}{
											&ruleRefExpr{
												pos:  position{line: 493, col: 35, offset: 22216},
												name: "NEWLINE",
											},
											&ruleRefExpr{
												pos:  position{line: 493, col: 43, offset: 22224},
												name: "BlankLine",
											},
										},
									},
								},
								&anyMatcher{
									line: 493, col: 54, offset: 22235,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "EndOfLiteralBlock",
			pos:  position{line: 498, col: 1, offset: 22320},
			expr: &choiceExpr{
				pos: position{line: 498, col: 22, offset: 22341},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 498, col: 22, offset: 22341},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 498, col: 22, offset: 22341},
								name: "NEWLINE",
							},
							&ruleRefExpr{
								pos:  position{line: 498, col: 30, offset: 22349},
								name: "BlankLine",
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 498, col: 42, offset: 22361},
						name: "NEWLINE",
					},
					&ruleRefExpr{
						pos:  position{line: 498, col: 52, offset: 22371},
						name: "EOF",
					},
				},
			},
		},
		{
			name: "ParagraphWithLiteralBlockDelimiter",
			pos:  position{line: 501, col: 1, offset: 22431},
			expr: &actionExpr{
				pos: position{line: 501, col: 39, offset: 22469},
				run: (*parser).callonParagraphWithLiteralBlockDelimiter1,
				expr: &seqExpr{
					pos: position{line: 501, col: 39, offset: 22469},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 501, col: 39, offset: 22469},
							name: "LiteralBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 501, col: 61, offset: 22491},
							expr: &ruleRefExpr{
								pos:  position{line: 501, col: 61, offset: 22491},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 501, col: 65, offset: 22495},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 501, col: 73, offset: 22503},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 501, col: 81, offset: 22511},
								expr: &seqExpr{
									pos: position{line: 501, col: 82, offset: 22512},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 501, col: 82, offset: 22512},
											expr: &ruleRefExpr{
												pos:  position{line: 501, col: 83, offset: 22513},
												name: "LiteralBlockDelimiter",
											},
										},
										&anyMatcher{
											line: 501, col: 105, offset: 22535,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 501, col: 109, offset: 22539},
							name: "LiteralBlockDelimiter",
						},
						&zeroOrMoreExpr{
							pos: position{line: 501, col: 131, offset: 22561},
							expr: &ruleRefExpr{
								pos:  position{line: 501, col: 131, offset: 22561},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 501, col: 135, offset: 22565},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "LiteralBlockDelimiter",
			pos:  position{line: 505, col: 1, offset: 22649},
			expr: &litMatcher{
				pos:        position{line: 505, col: 26, offset: 22674},
				val:        "....",
				ignoreCase: false,
			},
		},
		{
			name: "ParagraphWithLiteralAttribute",
			pos:  position{line: 508, col: 1, offset: 22736},
			expr: &actionExpr{
				pos: position{line: 508, col: 34, offset: 22769},
				run: (*parser).callonParagraphWithLiteralAttribute1,
				expr: &seqExpr{
					pos: position{line: 508, col: 34, offset: 22769},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 508, col: 34, offset: 22769},
							val:        "[literal]",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 508, col: 46, offset: 22781},
							expr: &ruleRefExpr{
								pos:  position{line: 508, col: 46, offset: 22781},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 508, col: 50, offset: 22785},
							name: "NEWLINE",
						},
						&labeledExpr{
							pos:   position{line: 508, col: 58, offset: 22793},
							label: "content",
							expr: &ruleRefExpr{
								pos:  position{line: 508, col: 67, offset: 22802},
								name: "LiteralBlockContent",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 508, col: 88, offset: 22823},
							name: "EndOfLiteralBlock",
						},
					},
				},
			},
		},
		{
			name: "ElementAttribute",
			pos:  position{line: 515, col: 1, offset: 23035},
			expr: &choiceExpr{
				pos: position{line: 515, col: 21, offset: 23055},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 515, col: 21, offset: 23055},
						name: "ElementLink",
					},
					&ruleRefExpr{
						pos:  position{line: 515, col: 35, offset: 23069},
						name: "ElementID",
					},
					&ruleRefExpr{
						pos:  position{line: 515, col: 47, offset: 23081},
						name: "ElementTitle",
					},
					&ruleRefExpr{
						pos:  position{line: 515, col: 62, offset: 23096},
						name: "InvalidElementAttribute",
					},
				},
			},
		},
		{
			name: "ElementLink",
			pos:  position{line: 518, col: 1, offset: 23176},
			expr: &actionExpr{
				pos: position{line: 518, col: 16, offset: 23191},
				run: (*parser).callonElementLink1,
				expr: &seqExpr{
					pos: position{line: 518, col: 16, offset: 23191},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 518, col: 16, offset: 23191},
							val:        "[link=",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 518, col: 25, offset: 23200},
							expr: &ruleRefExpr{
								pos:  position{line: 518, col: 25, offset: 23200},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 518, col: 29, offset: 23204},
							label: "path",
							expr: &ruleRefExpr{
								pos:  position{line: 518, col: 34, offset: 23209},
								name: "URL",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 518, col: 38, offset: 23213},
							expr: &ruleRefExpr{
								pos:  position{line: 518, col: 38, offset: 23213},
								name: "WS",
							},
						},
						&litMatcher{
							pos:        position{line: 518, col: 42, offset: 23217},
							val:        "]",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 518, col: 46, offset: 23221},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "ElementID",
			pos:  position{line: 522, col: 1, offset: 23277},
			expr: &choiceExpr{
				pos: position{line: 522, col: 14, offset: 23290},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 522, col: 14, offset: 23290},
						run: (*parser).callonElementID2,
						expr: &seqExpr{
							pos: position{line: 522, col: 14, offset: 23290},
							exprs: []interface{}{
								&labeledExpr{
									pos:   position{line: 522, col: 14, offset: 23290},
									label: "id",
									expr: &ruleRefExpr{
										pos:  position{line: 522, col: 18, offset: 23294},
										name: "InlineElementID",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 522, col: 35, offset: 23311},
									name: "EOL",
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 524, col: 5, offset: 23340},
						run: (*parser).callonElementID7,
						expr: &seqExpr{
							pos: position{line: 524, col: 5, offset: 23340},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 524, col: 5, offset: 23340},
									val:        "[#",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 524, col: 10, offset: 23345},
									label: "id",
									expr: &ruleRefExpr{
										pos:  position{line: 524, col: 14, offset: 23349},
										name: "ID",
									},
								},
								&litMatcher{
									pos:        position{line: 524, col: 18, offset: 23353},
									val:        "]",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 524, col: 22, offset: 23357},
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
			pos:  position{line: 528, col: 1, offset: 23409},
			expr: &actionExpr{
				pos: position{line: 528, col: 20, offset: 23428},
				run: (*parser).callonInlineElementID1,
				expr: &seqExpr{
					pos: position{line: 528, col: 20, offset: 23428},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 528, col: 20, offset: 23428},
							val:        "[[",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 528, col: 25, offset: 23433},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 528, col: 29, offset: 23437},
								name: "ID",
							},
						},
						&litMatcher{
							pos:        position{line: 528, col: 33, offset: 23441},
							val:        "]]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "ElementTitle",
			pos:  position{line: 534, col: 1, offset: 23636},
			expr: &actionExpr{
				pos: position{line: 534, col: 17, offset: 23652},
				run: (*parser).callonElementTitle1,
				expr: &seqExpr{
					pos: position{line: 534, col: 17, offset: 23652},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 534, col: 17, offset: 23652},
							val:        ".",
							ignoreCase: false,
						},
						&notExpr{
							pos: position{line: 534, col: 21, offset: 23656},
							expr: &litMatcher{
								pos:        position{line: 534, col: 22, offset: 23657},
								val:        ".",
								ignoreCase: false,
							},
						},
						&notExpr{
							pos: position{line: 534, col: 26, offset: 23661},
							expr: &ruleRefExpr{
								pos:  position{line: 534, col: 27, offset: 23662},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 534, col: 30, offset: 23665},
							label: "title",
							expr: &oneOrMoreExpr{
								pos: position{line: 534, col: 36, offset: 23671},
								expr: &seqExpr{
									pos: position{line: 534, col: 37, offset: 23672},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 534, col: 37, offset: 23672},
											expr: &ruleRefExpr{
												pos:  position{line: 534, col: 38, offset: 23673},
												name: "NEWLINE",
											},
										},
										&anyMatcher{
											line: 534, col: 46, offset: 23681,
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 534, col: 50, offset: 23685},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "InvalidElementAttribute",
			pos:  position{line: 538, col: 1, offset: 23750},
			expr: &actionExpr{
				pos: position{line: 538, col: 28, offset: 23777},
				run: (*parser).callonInvalidElementAttribute1,
				expr: &seqExpr{
					pos: position{line: 538, col: 28, offset: 23777},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 538, col: 28, offset: 23777},
							val:        "[",
							ignoreCase: false,
						},
						&oneOrMoreExpr{
							pos: position{line: 538, col: 32, offset: 23781},
							expr: &ruleRefExpr{
								pos:  position{line: 538, col: 32, offset: 23781},
								name: "WS",
							},
						},
						&labeledExpr{
							pos:   position{line: 538, col: 36, offset: 23785},
							label: "content",
							expr: &zeroOrMoreExpr{
								pos: position{line: 538, col: 44, offset: 23793},
								expr: &seqExpr{
									pos: position{line: 538, col: 45, offset: 23794},
									exprs: []interface{}{
										&notExpr{
											pos: position{line: 538, col: 45, offset: 23794},
											expr: &litMatcher{
												pos:        position{line: 538, col: 46, offset: 23795},
												val:        "]",
												ignoreCase: false,
											},
										},
										&anyMatcher{
											line: 538, col: 50, offset: 23799,
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 538, col: 54, offset: 23803},
							val:        "]",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "BlankLine",
			pos:  position{line: 545, col: 1, offset: 23969},
			expr: &actionExpr{
				pos: position{line: 545, col: 14, offset: 23982},
				run: (*parser).callonBlankLine1,
				expr: &seqExpr{
					pos: position{line: 545, col: 14, offset: 23982},
					exprs: []interface{}{
						&notExpr{
							pos: position{line: 545, col: 14, offset: 23982},
							expr: &ruleRefExpr{
								pos:  position{line: 545, col: 15, offset: 23983},
								name: "EOF",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 545, col: 19, offset: 23987},
							expr: &ruleRefExpr{
								pos:  position{line: 545, col: 19, offset: 23987},
								name: "WS",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 545, col: 23, offset: 23991},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "Characters",
			pos:  position{line: 552, col: 1, offset: 24138},
			expr: &actionExpr{
				pos: position{line: 552, col: 15, offset: 24152},
				run: (*parser).callonCharacters1,
				expr: &oneOrMoreExpr{
					pos: position{line: 552, col: 15, offset: 24152},
					expr: &seqExpr{
						pos: position{line: 552, col: 16, offset: 24153},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 552, col: 16, offset: 24153},
								expr: &ruleRefExpr{
									pos:  position{line: 552, col: 17, offset: 24154},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 552, col: 25, offset: 24162},
								expr: &ruleRefExpr{
									pos:  position{line: 552, col: 26, offset: 24163},
									name: "WS",
								},
							},
							&anyMatcher{
								line: 552, col: 29, offset: 24166,
							},
						},
					},
				},
			},
		},
		{
			name: "URL",
			pos:  position{line: 556, col: 1, offset: 24206},
			expr: &actionExpr{
				pos: position{line: 556, col: 8, offset: 24213},
				run: (*parser).callonURL1,
				expr: &oneOrMoreExpr{
					pos: position{line: 556, col: 8, offset: 24213},
					expr: &seqExpr{
						pos: position{line: 556, col: 9, offset: 24214},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 556, col: 9, offset: 24214},
								expr: &ruleRefExpr{
									pos:  position{line: 556, col: 10, offset: 24215},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 556, col: 18, offset: 24223},
								expr: &ruleRefExpr{
									pos:  position{line: 556, col: 19, offset: 24224},
									name: "WS",
								},
							},
							&notExpr{
								pos: position{line: 556, col: 22, offset: 24227},
								expr: &litMatcher{
									pos:        position{line: 556, col: 23, offset: 24228},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 556, col: 27, offset: 24232},
								expr: &litMatcher{
									pos:        position{line: 556, col: 28, offset: 24233},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 556, col: 32, offset: 24237,
							},
						},
					},
				},
			},
		},
		{
			name: "ID",
			pos:  position{line: 560, col: 1, offset: 24277},
			expr: &actionExpr{
				pos: position{line: 560, col: 7, offset: 24283},
				run: (*parser).callonID1,
				expr: &oneOrMoreExpr{
					pos: position{line: 560, col: 7, offset: 24283},
					expr: &seqExpr{
						pos: position{line: 560, col: 8, offset: 24284},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 560, col: 8, offset: 24284},
								expr: &ruleRefExpr{
									pos:  position{line: 560, col: 9, offset: 24285},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 560, col: 17, offset: 24293},
								expr: &ruleRefExpr{
									pos:  position{line: 560, col: 18, offset: 24294},
									name: "WS",
								},
							},
							&notExpr{
								pos: position{line: 560, col: 21, offset: 24297},
								expr: &litMatcher{
									pos:        position{line: 560, col: 22, offset: 24298},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 560, col: 26, offset: 24302},
								expr: &litMatcher{
									pos:        position{line: 560, col: 27, offset: 24303},
									val:        "]",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 560, col: 31, offset: 24307},
								expr: &litMatcher{
									pos:        position{line: 560, col: 32, offset: 24308},
									val:        "<<",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 560, col: 37, offset: 24313},
								expr: &litMatcher{
									pos:        position{line: 560, col: 38, offset: 24314},
									val:        ">>",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 560, col: 42, offset: 24318,
							},
						},
					},
				},
			},
		},
		{
			name: "URL_TEXT",
			pos:  position{line: 564, col: 1, offset: 24358},
			expr: &actionExpr{
				pos: position{line: 564, col: 13, offset: 24370},
				run: (*parser).callonURL_TEXT1,
				expr: &oneOrMoreExpr{
					pos: position{line: 564, col: 13, offset: 24370},
					expr: &seqExpr{
						pos: position{line: 564, col: 14, offset: 24371},
						exprs: []interface{}{
							&notExpr{
								pos: position{line: 564, col: 14, offset: 24371},
								expr: &ruleRefExpr{
									pos:  position{line: 564, col: 15, offset: 24372},
									name: "NEWLINE",
								},
							},
							&notExpr{
								pos: position{line: 564, col: 23, offset: 24380},
								expr: &litMatcher{
									pos:        position{line: 564, col: 24, offset: 24381},
									val:        "[",
									ignoreCase: false,
								},
							},
							&notExpr{
								pos: position{line: 564, col: 28, offset: 24385},
								expr: &litMatcher{
									pos:        position{line: 564, col: 29, offset: 24386},
									val:        "]",
									ignoreCase: false,
								},
							},
							&anyMatcher{
								line: 564, col: 33, offset: 24390,
							},
						},
					},
				},
			},
		},
		{
			name: "URL_SCHEME",
			pos:  position{line: 568, col: 1, offset: 24430},
			expr: &choiceExpr{
				pos: position{line: 568, col: 15, offset: 24444},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 568, col: 15, offset: 24444},
						val:        "http://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 568, col: 27, offset: 24456},
						val:        "https://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 568, col: 40, offset: 24469},
						val:        "ftp://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 568, col: 51, offset: 24480},
						val:        "irc://",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 568, col: 62, offset: 24491},
						val:        "mailto:",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "DIGIT",
			pos:  position{line: 570, col: 1, offset: 24502},
			expr: &charClassMatcher{
				pos:        position{line: 570, col: 10, offset: 24511},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "NEWLINE",
			pos:  position{line: 572, col: 1, offset: 24518},
			expr: &choiceExpr{
				pos: position{line: 572, col: 12, offset: 24529},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 572, col: 12, offset: 24529},
						val:        "\r\n",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 572, col: 21, offset: 24538},
						val:        "\r",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 572, col: 28, offset: 24545},
						val:        "\n",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "WS",
			pos:  position{line: 574, col: 1, offset: 24551},
			expr: &choiceExpr{
				pos: position{line: 574, col: 7, offset: 24557},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 574, col: 7, offset: 24557},
						val:        " ",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 574, col: 13, offset: 24563},
						run: (*parser).callonWS3,
						expr: &litMatcher{
							pos:        position{line: 574, col: 13, offset: 24563},
							val:        "\t",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "EOF",
			pos:  position{line: 578, col: 1, offset: 24608},
			expr: &notExpr{
				pos: position{line: 578, col: 8, offset: 24615},
				expr: &anyMatcher{
					line: 578, col: 9, offset: 24616,
				},
			},
		},
		{
			name: "EOL",
			pos:  position{line: 580, col: 1, offset: 24619},
			expr: &choiceExpr{
				pos: position{line: 580, col: 8, offset: 24626},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 580, col: 8, offset: 24626},
						name: "NEWLINE",
					},
					&ruleRefExpr{
						pos:  position{line: 580, col: 18, offset: 24636},
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
