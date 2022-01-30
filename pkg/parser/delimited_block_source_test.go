package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" // nolint:golint
	. "github.com/onsi/gomega" // nolint:golintt
)

var _ = Describe("source blocks", func() {

	Context("in final documents", func() {

		It("with source attribute only", func() {
			source := `[source]
----
require 'sinatra'

get '/hi' do
  "Hello World!"
end
----`
			expected := &types.Document{
				Elements: []interface{}{
					&types.DelimitedBlock{
						Kind: types.Listing,
						Attributes: types.Attributes{
							types.AttrStyle: types.Source,
						},
						Elements: []interface{}{
							&types.StringElement{
								Content: `require 'sinatra'

get '/hi' do
  "Hello World!"
end`,
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
			expected := &types.Document{
				Elements: []interface{}{
					&types.DelimitedBlock{
						Kind: types.Listing,
						Attributes: types.Attributes{
							types.AttrStyle:    types.Source,
							types.AttrLanguage: "ruby",
							types.AttrTitle:    "Source block title",
						},
						Elements: []interface{}{
							&types.StringElement{
								Content: `require 'sinatra'

get '/hi' do
  "Hello World!"
end`,
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("with id, title, source, languages and linenums attributes", func() {
			source := `[#id-for-source-block]
[source,ruby,linenums]
.app.rb
----
require 'sinatra'

get '/hi' do
  "Hello World!"
end
----`
			expected := &types.Document{
				Elements: []interface{}{
					&types.DelimitedBlock{
						Kind: types.Listing,
						Attributes: types.Attributes{
							types.AttrStyle:    types.Source,
							types.AttrLanguage: "ruby",
							types.AttrID:       "id-for-source-block",
							types.AttrTitle:    "app.rb",
							types.AttrLineNums: true,
						},
						Elements: []interface{}{
							&types.StringElement{
								Content: `require 'sinatra'

get '/hi' do
  "Hello World!"
end`,
							},
						},
					},
				},
				ElementReferences: types.ElementReferences{
					"id-for-source-block": "app.rb",
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("with callout and admonition block afterwards", func() {
			source := `[source]
----
const cookies = "cookies" <1>
----
<1> a constant

[NOTE]
====
a note
====`

			expected := &types.Document{
				Elements: []interface{}{
					&types.DelimitedBlock{
						Kind: types.Listing,
						Attributes: types.Attributes{
							types.AttrStyle: types.Source,
						},
						Elements: []interface{}{
							&types.StringElement{
								Content: `const cookies = "cookies" `,
							},
							&types.Callout{
								Ref: 1,
							},
						},
					},
					&types.List{
						Kind: types.CalloutListKind,
						Elements: []types.ListElement{
							&types.CalloutListElement{
								Ref: 1,
								Elements: []interface{}{
									&types.Paragraph{
										Elements: []interface{}{
											&types.StringElement{
												Content: "a constant",
											},
										},
									},
								},
							},
						},
					},
					&types.DelimitedBlock{
						Kind: types.Example,
						Attributes: types.Attributes{
							types.AttrStyle: types.Note,
						},
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "a note",
									},
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
