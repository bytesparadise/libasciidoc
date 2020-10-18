package parser_test

import (
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("listing blocks", func() {

	Context("draft documents", func() {

		Context("delimited blocks", func() {

			It("with single line", func() {
				source := `----
some listing code
----`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "some listing code",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with no line", func() {
				source := `----
----`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with multiple lines alone", func() {
				source := `----
some listing code
with an empty line

in the middle
----`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "some listing code",
									},
								},
								{
									types.StringElement{
										Content: "with an empty line",
									},
								},
								{},
								{
									types.StringElement{
										Content: "in the middle",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with unrendered list", func() {
				source := `----
* some 
* listing 
* content
----`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "* some ",
									},
								},
								{
									types.StringElement{
										Content: "* listing ",
									},
								},
								{
									types.StringElement{
										Content: "* content",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with multiple lines then a paragraph", func() {
				source := `---- 
some listing code
with an empty line

in the middle
----
then a normal paragraph.`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "some listing code",
									},
								},
								{
									types.StringElement{
										Content: "with an empty line",
									},
								},
								{},
								{
									types.StringElement{
										Content: "in the middle",
									},
								},
							},
						},
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "then a normal paragraph.",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("after a paragraph", func() {
				source := `a paragraph.

----
some listing code
----`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a paragraph.",
									},
								},
							},
						},
						types.BlankLine{}, // blankline is required between paragraph and the next block
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "some listing code",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with unclosed delimiter", func() {
				source := `----
End of file here.`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "End of file here.",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with single callout", func() {
				source := `----
<import> <1>
----
<1> an import`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.SpecialCharacter{
										Name: "<",
									},
									types.StringElement{
										Content: "import",
									},
									types.SpecialCharacter{
										Name: ">",
									},
									types.StringElement{
										Content: " ",
									},
									types.Callout{
										Ref: 1,
									},
								},
							},
						},
						types.CalloutListItem{
							Ref: 1,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "an import",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with multiple callouts on different lines", func() {
				source := `----
import <1>

func foo() {} <2>
----
<1> an import
<2> a func`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "import ",
									},
									types.Callout{
										Ref: 1,
									},
								},
								{},
								{
									types.StringElement{
										Content: "func foo() {} ",
									},
									types.Callout{
										Ref: 2,
									},
								},
							},
						},
						types.CalloutListItem{
							Ref: 1,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "an import",
											},
										},
									},
								},
							},
						},
						types.CalloutListItem{
							Ref: 2,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a func",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with multiple callouts on same line", func() {
				source := `----
import <1> <2><3>

func foo() {} <4>
----
<1> an import
<2> a single import
<3> a single basic import
<4> a func`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "import ",
									},
									types.Callout{
										Ref: 1,
									},
									types.Callout{
										Ref: 2,
									},
									types.Callout{
										Ref: 3,
									},
								},
								{},
								{
									types.StringElement{
										Content: "func foo() {} ",
									},
									types.Callout{
										Ref: 4,
									},
								},
							},
						},
						types.CalloutListItem{
							Ref: 1,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "an import",
											},
										},
									},
								},
							},
						},
						types.CalloutListItem{
							Ref: 2,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a single import",
											},
										},
									},
								},
							},
						},
						types.CalloutListItem{
							Ref: 3,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a single basic import",
											},
										},
									},
								},
							},
						},
						types.CalloutListItem{
							Ref: 4,
							Elements: []interface{}{
								types.Paragraph{
									Lines: [][]interface{}{
										{
											types.StringElement{
												Content: "a func",
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with invalid callout", func() {
				source := `----
import <a>
----
<a> an import`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "import ",
									},
									types.SpecialCharacter{
										Name: "<",
									},
									types.StringElement{
										Content: "a",
									},
									types.SpecialCharacter{
										Name: ">",
									},
								},
							},
						},
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.SpecialCharacter{
										Name: "<",
									},
									types.StringElement{
										Content: "a",
									},
									types.SpecialCharacter{
										Name: ">",
									},
									types.StringElement{
										Content: " an import",
									},
								},
							},
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			sourceCode := [][]interface{}{
				{
					types.StringElement{
						Content: "package foo",
					},
				},
				{},
				{
					types.StringElement{
						Content: "// Foo",
					},
				},
				{
					types.StringElement{
						Content: "type Foo struct{",
					},
				},
				{
					types.StringElement{
						Content: "    Bar string",
					},
				},
				{
					types.StringElement{
						Content: "}",
					},
				},
			}

			It("with source attribute only", func() {
				source := `[source]
----
package foo

// Foo
type Foo struct{
    Bar string
}
----`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.ListingBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Source,
							},
							Lines: sourceCode,
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with source attribute and comma", func() {
				source := `[source,]
----
package foo

// Foo
type Foo struct{
    Bar string
}
----`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.ListingBlock{
							Attributes: types.Attributes{
								types.AttrKind: types.Source,
							},
							Lines: sourceCode,
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with title, source and language attributes", func() {
				source := `[source,go]
.foo.go
----
package foo

// Foo
type Foo struct{
    Bar string
}
----`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.ListingBlock{
							Attributes: types.Attributes{
								types.AttrKind:     types.Source,
								types.AttrLanguage: "go",
								types.AttrTitle:    "foo.go",
							},
							Lines: sourceCode,
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})

			It("with id, title, source and language and other attributes", func() {
				source := `[#id-for-source-block]
[source,go,linenums]
.foo.go
----
package foo

// Foo
type Foo struct{
    Bar string
}
----`
				expected := types.DraftDocument{
					Elements: []interface{}{
						types.ListingBlock{
							Attributes: types.Attributes{
								types.AttrKind:     types.Source,
								types.AttrLanguage: "go",
								types.AttrID:       "id-for-source-block",
								types.AttrCustomID: true,
								types.AttrTitle:    "foo.go",
								types.AttrLineNums: nil,
							},
							Lines: sourceCode,
						},
					},
				}
				Expect(ParseDraftDocument(source)).To(MatchDraftDocument(expected))
			})
		})
	})

	Context("final documents", func() {

		Context("delimited blocks", func() {

			It("with single line", func() {
				source := `----
some listing code
----`
				expected := types.Document{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "some listing code",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with no line", func() {
				source := `----
----`
				expected := types.Document{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with multiple lines alone", func() {
				source := `----
some listing code
with an empty line

in the middle
----`
				expected := types.Document{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "some listing code",
									},
								},
								{
									types.StringElement{
										Content: "with an empty line",
									},
								},
								{},
								{
									types.StringElement{
										Content: "in the middle",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with unrendered list", func() {
				source := `----
* some 
* listing 
* content
----`
				expected := types.Document{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "* some ",
									},
								},
								{
									types.StringElement{
										Content: "* listing ",
									},
								},
								{
									types.StringElement{
										Content: "* content",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with multiple lines then a paragraph", func() {
				source := `---- 
some listing code
with an empty line

in the middle
----
then a normal paragraph.`
				expected := types.Document{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "some listing code",
									},
								},
								{
									types.StringElement{
										Content: "with an empty line",
									},
								},
								{},
								{
									types.StringElement{
										Content: "in the middle",
									},
								},
							},
						},
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "then a normal paragraph."},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("after a paragraph", func() {
				source := `a paragraph.
	
----
some listing code
----`
				expected := types.Document{
					Elements: []interface{}{
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "a paragraph.",
									},
								},
							},
						},
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "some listing code",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with unclosed delimiter", func() {
				source := `----
End of file here.`
				expected := types.Document{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "End of file here.",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with single callout", func() {
				source := `----
import <1>
----
<1> an import`
				expected := types.Document{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "import ",
									},
									types.Callout{
										Ref: 1,
									},
								},
							},
						},
						types.CalloutList{
							Items: []types.CalloutListItem{
								{
									Ref: 1,
									Elements: []interface{}{
										types.Paragraph{
											Lines: [][]interface{}{
												{
													types.StringElement{
														Content: "an import",
													},
												},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with multiple callouts on different lines", func() {
				source := `----
import <1>

func foo() {} <2>
----
<1> an import
<2> a func`
				expected := types.Document{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "import ",
									},
									types.Callout{
										Ref: 1,
									},
								},
								{},
								{
									types.StringElement{
										Content: "func foo() {} ",
									},
									types.Callout{
										Ref: 2,
									},
								},
							},
						},
						types.CalloutList{
							Items: []types.CalloutListItem{
								{
									Ref: 1,
									Elements: []interface{}{
										types.Paragraph{
											Lines: [][]interface{}{
												{
													types.StringElement{
														Content: "an import",
													},
												},
											},
										},
									},
								},
								{
									Ref: 2,
									Elements: []interface{}{
										types.Paragraph{
											Lines: [][]interface{}{
												{
													types.StringElement{
														Content: "a func",
													},
												},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with multiple callouts on same line", func() {
				source := `----
import <1> <2><3>

func foo() {} <4>
----
<1> an import
<2> a single import
<3> a single basic import
<4> a func`
				expected := types.Document{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "import ",
									},
									types.Callout{
										Ref: 1,
									},
									types.Callout{
										Ref: 2,
									},
									types.Callout{
										Ref: 3,
									},
								},
								{},
								{
									types.StringElement{
										Content: "func foo() {} ",
									},
									types.Callout{
										Ref: 4,
									},
								},
							},
						},
						types.CalloutList{
							Items: []types.CalloutListItem{
								{
									Ref: 1,
									Elements: []interface{}{
										types.Paragraph{
											Lines: [][]interface{}{
												{
													types.StringElement{
														Content: "an import",
													},
												},
											},
										},
									},
								},
								{
									Ref: 2,
									Elements: []interface{}{
										types.Paragraph{
											Lines: [][]interface{}{
												{
													types.StringElement{
														Content: "a single import",
													},
												},
											},
										},
									},
								},
								{
									Ref: 3,
									Elements: []interface{}{
										types.Paragraph{
											Lines: [][]interface{}{
												{
													types.StringElement{
														Content: "a single basic import",
													},
												},
											},
										},
									},
								},
								{
									Ref: 4,
									Elements: []interface{}{
										types.Paragraph{
											Lines: [][]interface{}{
												{
													types.StringElement{
														Content: "a func",
													},
												},
											},
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("with invalid callout", func() {
				source := `----
import <a>
----
<a> an import`
				expected := types.Document{
					Elements: []interface{}{
						types.ListingBlock{
							Lines: [][]interface{}{
								{
									types.StringElement{
										Content: "import ",
									},
									types.SpecialCharacter{
										Name: "<",
									},
									types.StringElement{
										Content: "a",
									},
									types.SpecialCharacter{
										Name: ">",
									},
								},
							},
						},
						types.Paragraph{
							Lines: [][]interface{}{
								{
									types.SpecialCharacter{
										Name: "<",
									},
									types.StringElement{
										Content: "a",
									},
									types.SpecialCharacter{
										Name: ">",
									},
									types.StringElement{
										Content: " an import",
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			Context("source block", func() {

				It("with source attribute only", func() {
					source := `[source]
----
require 'sinatra'

get '/hi' do
  "Hello World!"
end
----`
					expected := types.Document{
						Elements: []interface{}{
							types.ListingBlock{
								Attributes: types.Attributes{
									types.AttrKind: types.Source,
								},
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "require 'sinatra'",
										},
									},
									{},
									{
										types.StringElement{
											Content: "get '/hi' do",
										},
									},
									{
										types.StringElement{
											Content: "  \"Hello World!\"",
										},
									},
									{
										types.StringElement{
											Content: "end",
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("with title, source and languages attributes", func() {
					source := `[source,ruby]
.Source block title
----
require 'sinatra'

get '/hi' do
  "Hello World!"
end
----`
					expected := types.Document{
						Elements: []interface{}{
							types.ListingBlock{
								Attributes: types.Attributes{
									types.AttrKind:     types.Source,
									types.AttrLanguage: "ruby",
									types.AttrTitle:    "Source block title",
								},
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "require 'sinatra'",
										},
									},
									{},
									{
										types.StringElement{
											Content: "get '/hi' do",
										},
									},
									{
										types.StringElement{
											Content: "  \"Hello World!\"",
										},
									},
									{
										types.StringElement{
											Content: "end",
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("with id, title, source and languages attributes", func() {
					source := `[#id-for-source-block]
[source,ruby]
.app.rb
----
require 'sinatra'

get '/hi' do
  "Hello World!"
end
----`
					expected := types.Document{
						Elements: []interface{}{
							types.ListingBlock{
								Attributes: types.Attributes{
									types.AttrKind:     types.Source,
									types.AttrLanguage: "ruby",
									types.AttrID:       "id-for-source-block",
									types.AttrCustomID: true,
									types.AttrTitle:    "app.rb",
								},
								Lines: [][]interface{}{
									{
										types.StringElement{
											Content: "require 'sinatra'",
										},
									},
									{},
									{
										types.StringElement{
											Content: "get '/hi' do",
										},
									},
									{
										types.StringElement{
											Content: "  \"Hello World!\"",
										},
									},
									{
										types.StringElement{
											Content: "end",
										},
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})
			})
		})
	})

	Context("with custom substitutions", func() {

		// testing custom substitutions on listing blocks only, as
		// other verbatim blocks (fenced, literal, source, passthrough)
		// share the same implementation

		source := `:github-url: https://github.com
			
[subs="$SUBS"]
----
a link to https://example.com[] <1>
and <more text> on the +
*next* lines with a link to {github-url}[]

* not a list item
----

<1> a callout
`

		It("should apply the default substitution", func() {
			s := strings.ReplaceAll(source, "[subs=\"$SUBS\"]", "")
			expected := `<div class="listingblock">
<div class="content">
<pre>a link to https://example.com[] <b class="conum">(1)</b>
and &lt;more text&gt; on the +
*next* lines with a link to {github-url}[]

* not a list item</pre>
</div>
</div>
<div class="colist arabic">
<ol>
<li>
<p>a callout</p>
</li>
</ol>
</div>
`
			Expect(RenderHTML(s)).To(MatchHTML(expected))
		})

		It("should apply the 'normal' substitution", func() {
			s := strings.ReplaceAll(source, "$SUBS", "normal")
			expected := `<div class="listingblock">
<div class="content">
<pre>a link to <a href="https://example.com" class="bare">https://example.com</a> &lt;1&gt;
and &lt;more text&gt; on the<br>
<strong>next</strong> lines with a link to <a href="https://github.com" class="bare">https://github.com</a>

* not a list item</pre>
</div>
</div>
<div class="colist arabic">
<ol>
<li>
<p>a callout</p>
</li>
</ol>
</div>
`
			Expect(RenderHTML(s)).To(MatchHTML(expected))
		})

		It("should apply the 'quotes' substitution", func() {
			s := strings.ReplaceAll(source, "$SUBS", "quotes")
			expected := `<div class="listingblock">
<div class="content">
<pre>a link to https://example.com[] <1>
and <more text> on the +
<strong>next</strong> lines with a link to {github-url}[]

* not a list item</pre>
</div>
</div>
<div class="colist arabic">
<ol>
<li>
<p>a callout</p>
</li>
</ol>
</div>
`
			Expect(RenderHTML(s)).To(MatchHTML(expected))
		})

		It("should apply the 'macros' substitution", func() {
			s := strings.ReplaceAll(source, "$SUBS", "macros")
			expected := `<div class="listingblock">
<div class="content">
<pre>a link to <a href="https://example.com" class="bare">https://example.com</a> <1>
and <more text> on the +
*next* lines with a link to {github-url}[]

* not a list item</pre>
</div>
</div>
<div class="colist arabic">
<ol>
<li>
<p>a callout</p>
</li>
</ol>
</div>
`
			Expect(RenderHTML(s)).To(MatchHTML(expected))
		})

		It("should apply the 'attributes' substitution", func() {
			s := strings.ReplaceAll(source, "$SUBS", "attributes")
			expected := `<div class="listingblock">
<div class="content">
<pre>a link to https://example.com[] <1>
and <more text> on the +
*next* lines with a link to https://github.com[]

* not a list item</pre>
</div>
</div>
<div class="colist arabic">
<ol>
<li>
<p>a callout</p>
</li>
</ol>
</div>
`
			Expect(RenderHTML(s)).To(MatchHTML(expected))
		})

		It("should apply the 'attributes,macros' substitution", func() {
			s := strings.ReplaceAll(source, "$SUBS", "attributes,macros")
			expected := `<div class="listingblock">
<div class="content">
<pre>a link to <a href="https://example.com" class="bare">https://example.com</a> <1>
and <more text> on the +
*next* lines with a link to <a href="https://github.com" class="bare">https://github.com</a>

* not a list item</pre>
</div>
</div>
<div class="colist arabic">
<ol>
<li>
<p>a callout</p>
</li>
</ol>
</div>
`
			Expect(RenderHTML(s)).To(MatchHTML(expected))
		})

		It("should apply the 'specialchars' substitution", func() {
			s := strings.ReplaceAll(source, "$SUBS", "specialchars")
			expected := `<div class="listingblock">
<div class="content">
<pre>a link to https://example.com[] &lt;1&gt;
and &lt;more text&gt; on the +
*next* lines with a link to {github-url}[]

* not a list item</pre>
</div>
</div>
<div class="colist arabic">
<ol>
<li>
<p>a callout</p>
</li>
</ol>
</div>
`
			Expect(RenderHTML(s)).To(MatchHTML(expected))
		})

		It("should apply the 'replacements' substitution", func() {
			s := strings.ReplaceAll(source, "$SUBS", "replacements")
			expected := `<div class="listingblock">
<div class="content">
<pre>a link to https://example.com[] <1>
and <more text> on the +
*next* lines with a link to {github-url}[]

* not a list item</pre>
</div>
</div>
<div class="colist arabic">
<ol>
<li>
<p>a callout</p>
</li>
</ol>
</div>
`
			Expect(RenderHTML(s)).To(MatchHTML(expected))
		})

		It("should apply the 'post_replacements' substitution", func() {
			s := strings.ReplaceAll(source, "$SUBS", "post_replacements")
			expected := `<div class="listingblock">
<div class="content">
<pre>a link to https://example.com[] <1>
and <more text> on the<br>
*next* lines with a link to {github-url}[]

* not a list item</pre>
</div>
</div>
<div class="colist arabic">
<ol>
<li>
<p>a callout</p>
</li>
</ol>
</div>
`
			Expect(RenderHTML(s)).To(MatchHTML(expected))
		})

		It("should apply the 'quotes,macros' substitution", func() {
			s := strings.ReplaceAll(source, "$SUBS", "quotes,macros")
			expected := `<div class="listingblock">
<div class="content">
<pre>a link to <a href="https://example.com" class="bare">https://example.com</a> <1>
and <more text> on the +
<strong>next</strong> lines with a link to {github-url}[]

* not a list item</pre>
</div>
</div>
<div class="colist arabic">
<ol>
<li>
<p>a callout</p>
</li>
</ol>
</div>
`
			Expect(RenderHTML(s)).To(MatchHTML(expected))
		})

		It("should apply the 'macros,quotes' substitution", func() {
			s := strings.ReplaceAll(source, "$SUBS", "macros,quotes")
			expected := `<div class="listingblock">
<div class="content">
<pre>a link to <a href="https://example.com" class="bare">https://example.com</a> <1>
and <more text> on the +
<strong>next</strong> lines with a link to {github-url}[]

* not a list item</pre>
</div>
</div>
<div class="colist arabic">
<ol>
<li>
<p>a callout</p>
</li>
</ol>
</div>
`
			Expect(RenderHTML(s)).To(MatchHTML(expected))
		})

		It("should apply the 'none' substitution", func() {
			s := strings.ReplaceAll(source, "$SUBS", "none")
			expected := `<div class="listingblock">
<div class="content">
<pre>a link to https://example.com[] <1>
and <more text> on the +
*next* lines with a link to {github-url}[]

* not a list item</pre>
</div>
</div>
<div class="colist arabic">
<ol>
<li>
<p>a callout</p>
</li>
</ol>
</div>
`
			Expect(RenderHTML(s)).To(MatchHTML(expected))
		})
	})

})
