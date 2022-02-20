package parser_test

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" // nolint:golint
	. "github.com/onsi/gomega" // nolint:golintt
	log "github.com/sirupsen/logrus"
)

var _ = Describe("file inclusions", func() {

	Context("in preparsed documents", func() {

		It("without file inclusion", func() {
			source := `= Title
:toc: 
:author: Xavier
:!author:

cookie
chocolate`
			expected := source // unchanged
			Expect(PreparseDocument(source)).To(Equal(expected))
		})

		It("should include adoc file without leveloffset from local file", func() {
			source := "include::../../test/includes/chapter-a.adoc[]"
			expected := `= Chapter A

content`
			Expect(PreparseDocument(source)).To(Equal(expected))
		})

		It("should include adoc file with leveloffset", func() {
			source := "include::../../test/includes/chapter-a.adoc[leveloffset=+1]"
			expected := `== Chapter A

content`
			Expect(PreparseDocument(source)).To(Equal(expected))
		})

		It("should include file with attribute in path", func() {
			source := `:includedir: ../../test/includes

include::{includedir}/chapter-a.adoc[]`
			expected := `:includedir: ../../test/includes

= Chapter A

content`
			Expect(PreparseDocument(source)).To(Equal(expected))
		})

		It("should not further process with non-asciidoc files", func() {
			source := `:includedir: ../../test/includes

include::{includedir}/include.foo[]`
			expected := `:includedir: ../../test/includes

*some strong content*

include::hello_world.go.txt[]
`
			Expect(PreparseDocument(source)).To(Equal(expected))
		})

		It("should include grandchild content without offset", func() {
			source := `include::../../test/includes/grandchild-include.adoc[]`
			expected := `== grandchild title

first line of grandchild

last line of grandchild`
			Expect(PreparseDocument(source)).To(Equal(expected))
		})

		It("should include grandchild content with relative offset", func() {
			source := `include::../../test/includes/grandchild-include.adoc[leveloffset=+1]`
			expected := `=== grandchild title

first line of grandchild

last line of grandchild`
			Expect(PreparseDocument(source)).To(Equal(expected))
		})

		It("should include grandchild content with absolute offset", func() {
			source := `include::../../test/includes/grandchild-include.adoc[leveloffset=0]`
			expected := `= grandchild title

first line of grandchild

last line of grandchild`
			Expect(PreparseDocument(source)).To(Equal(expected))
		})

		It("should include child and grandchild content with relative level offset", func() {
			source := `include::../../test/includes/parent-include-relative-offset.adoc[leveloffset=+1]`
			expected := `== parent title

first line of parent

child preamble

==== child section 1

first line of child

===== grandchild title

first line of grandchild

last line of grandchild

===== child section 2

last line of child

last line of parent`
			Expect(PreparseDocument(source)).To(Equal(expected))
		})

		It("should include child and grandchild content with relative then absolute level offset", func() {
			source := `include::../../test/includes/parent-include-absolute-offset.adoc[leveloffset=+1]`
			expected := `== parent title

first line of parent

child preamble

==== child section 1

first line of child

===== grandchild title

first line of grandchild

last line of grandchild

===== child section 2

last line of child

last line of parent`
			Expect(PreparseDocument(source)).To(Equal(expected))
		})

		It("should include adoc with attributes file within main content", func() {
			source := `include::../../test/includes/attributes.adoc[]`
			expected := `:author: Xavier
:leveloffset: +1

some content`
			Expect(PreparseDocument(source)).To(Equal(expected))
		})

		Context("within delimited blocks", func() {

			It("should include adoc with attributes file within listing block", func() {
				// the `leveloffset=+1` attribute does not have effect on sections within a delimited block
				source := `----
include::../../test/includes/attributes.adoc[]
----`
				expected := `----
:author: Xavier
:leveloffset: +1

some content
----`
				Expect(PreparseDocument(source)).To(Equal(expected))
			})

			It("should include adoc file within comment block", func() {
				// the `leveloffset=+1` attribute does not have effect on sections within a delimited block
				source := `////
include::../../test/includes/parent-include.adoc[leveloffset=+1]
////`
				expected := `////
:leveloffset: +1

= parent title

first line of parent

= child title

first line of child

== grandchild title

first line of grandchild

last line of grandchild

last line of child

last line of parent <1>

:leveloffset!:
////`
				Expect(PreparseDocument(source)).To(Equal(expected))
			})

			It("should include adoc file within example block", func() {
				// the `leveloffset=+1` attribute does not have effect on sections within a delimited block
				source := `====
include::../../test/includes/parent-include.adoc[leveloffset=+1]
====`
				expected := `====
:leveloffset: +1

= parent title

first line of parent

= child title

first line of child

== grandchild title

first line of grandchild

last line of grandchild

last line of child

last line of parent <1>

:leveloffset!:
====`
				Expect(PreparseDocument(source)).To(Equal(expected))
			})

			It("should include adoc file within fenced block", func() {
				// the `leveloffset=+1` attribute does not have effect on sections within a delimited block
				source := "```\n" +
					"include::../../test/includes/parent-include.adoc[leveloffset=+1]\n" +
					"```\n" +
					"<1> a callout"
				expected := "```\n" +
					`:leveloffset: +1

= parent title

first line of parent

= child title

first line of child

== grandchild title

first line of grandchild

last line of grandchild

last line of child

last line of parent <1>

:leveloffset!:` +
					"\n```\n" +
					`<1> a callout`
				Expect(PreparseDocument(source)).To(Equal(expected))
			})

			It("should include adoc file within listing block", func() {
				// the `leveloffset=+1` attribute does not have effect on sections within a delimited block
				source := `----
include::../../test/includes/parent-include.adoc[leveloffset=+1]
----`
				expected := `----
:leveloffset: +1

= parent title

first line of parent

= child title

first line of child

== grandchild title

first line of grandchild

last line of grandchild

last line of child

last line of parent <1>

:leveloffset!:
----`
				Expect(PreparseDocument(source)).To(Equal(expected))
			})

			It("should include adoc file within literal block", func() {
				// the `leveloffset=+1` attribute does not have effect on sections within a delimited block
				source := `....
include::../../test/includes/parent-include.adoc[leveloffset=+1]
....`
				expected := `....
:leveloffset: +1

= parent title

first line of parent

= child title

first line of child

== grandchild title

first line of grandchild

last line of grandchild

last line of child

last line of parent <1>

:leveloffset!:
....`
				Expect(PreparseDocument(source)).To(Equal(expected))
			})

			It("should include adoc file within passthrough block", func() {
				// the `leveloffset=+1` attribute does not have effect on sections within a delimited block
				source := `++++
include::../../test/includes/parent-include.adoc[leveloffset=+1]
++++`
				expected := `++++
:leveloffset: +1

= parent title

first line of parent

= child title

first line of child

== grandchild title

first line of grandchild

last line of grandchild

last line of child

last line of parent <1>

:leveloffset!:
++++`
				Expect(PreparseDocument(source)).To(Equal(expected))
			})

			It("should include adoc file within quote block", func() {
				// the `leveloffset=+1` attribute does not have effect on sections within a delimited block
				source := `____
include::../../test/includes/parent-include.adoc[leveloffset=+1]
____`
				expected := `____
:leveloffset: +1

= parent title

first line of parent

= child title

first line of child

== grandchild title

first line of grandchild

last line of grandchild

last line of child

last line of parent <1>

:leveloffset!:
____`
				Expect(PreparseDocument(source)).To(Equal(expected))
			})

			It("should include adoc file within sidebar block", func() {
				// the `leveloffset=+1` attribute does not have effect on sections within a delimited block
				source := `****
include::../../test/includes/parent-include.adoc[leveloffset=+1]
****`
				expected := `****
:leveloffset: +1

= parent title

first line of parent

= child title

first line of child

== grandchild title

first line of grandchild

last line of grandchild

last line of child

last line of parent <1>

:leveloffset!:
****`
				Expect(PreparseDocument(source)).To(Equal(expected))
			})
		})

		Context("within tables", func() {

			It("default", func() {
				source := "include::../../test/includes/table_parent.adoc[]"
				expected := `|===
| Header A | Header B

| Column A.1 | Column A.2
| Column B.1 | Column B.2
|===`
				Expect(PreparseDocument(source)).To(Equal(expected))
			})
		})

		Context("with line ranges", func() {

			Context("unquoted", func() {

				It("with single unquoted line", func() {
					source := `include::../../test/includes/chapter-a.adoc[lines=1]`
					expected := `= Chapter A`
					Expect(PreparseDocument(source)).To(Equal(expected))
				})

				It("with multiple unquoted lines", func() {
					source := `include::../../test/includes/chapter-a.adoc[lines=1..3]`
					expected := `= Chapter A

content`
					Expect(PreparseDocument(source)).To(Equal(expected))
				})

				It("with multiple unquoted ranges (becoming authors)", func() {
					source := `include::../../test/includes/chapter-a.adoc[lines=1;3..4;6..-1]` // paragraph becomes the author since the in-between blank line is stripped out
					expected := `= Chapter A
content`
					Expect(PreparseDocument(source)).To(Equal(expected))
				})

				It("with invalid unquoted range - case 1", func() {
					source := `include::../../test/includes/chapter-a.adoc[lines=1;3..4;6..foo]` // not a number
					_, err := PreparseDocument(source)
					Expect(err).To(MatchError(`Unresolved directive in test.adoc - include::../../test/includes/chapter-a.adoc[lines=1;3..4;6..foo]: 1:11 (10): no match found, expected: "-" or [0-9]`))
				})

				It("with invalid unquoted range - case 2", func() {
					source := `include::../../test/includes/chapter-a.adoc[lines=1,3..4,6..-1]` // using commas instead of semi-colons
					expected := `= Chapter A`
					Expect(PreparseDocument(source)).To(Equal(expected))
				})
			})

			Context("quoted", func() {

				It("with single line", func() {
					source := `include::../../test/includes/chapter-a.adoc[lines="1"]`
					expected := `= Chapter A`
					Expect(PreparseDocument(source)).To(Equal(expected))
				})

				It("with multiple lines", func() {
					source := `include::../../test/includes/chapter-a.adoc[lines="1..2"]`
					expected := `= Chapter A
`
					Expect(PreparseDocument(source)).To(Equal(expected))
				})

				It("with multiple ranges with colons (becoming authors)", func() {
					// here, the `content` paragraph gets attached to the header and becomes the author
					source := `include::../../test/includes/chapter-a.adoc[lines="1,3..4,6..-1"]`
					expected := `= Chapter A
content`
					Expect(PreparseDocument(source)).To(Equal(expected))
				})

				It("with multiple ranges with semicolons (becoming authors)", func() {
					// here, the `content` paragraph gets attached to the header and becomes the author
					source := `include::../../test/includes/chapter-a.adoc[lines="1;3..4;6..-1"]`
					expected := `= Chapter A
content`
					Expect(PreparseDocument(source)).To(Equal(expected))
				})

				It("with invalid range - case 1", func() {
					source := `include::../../test/includes/chapter-a.adoc[lines="1,3..4,6..foo"]` // not a number
					_, err := PreparseDocument(source)
					Expect(err).To(MatchError(`Unresolved directive in test.adoc - include::../../test/includes/chapter-a.adoc[lines="1,3..4,6..foo"]: 1:11 (10): no match found, expected: "-" or [0-9]`))
				})

				It("with ignored tags", func() {
					// include using a line range a file having tags
					source := `include::../../test/includes/tag-include.adoc[lines=3]`
					expected := `== Section 1`
					Expect(PreparseDocument(source)).To(Equal(expected))
				})
			})
		})

		Context("with tag ranges", func() {

			It("with single tag", func() {
				source := `include::../../test/includes/tag-include.adoc[tag=section]`
				expected := `== Section 1`
				Expect(PreparseDocument(source)).To(Equal(expected))
			})

			It("with surrounding tag", func() {
				source := `include::../../test/includes/tag-include.adoc[tag=doc]`
				expected := `== Section 1

content
`
				Expect(PreparseDocument(source)).To(Equal(expected))
			})

			It("with unclosed tag", func() {
				// setup logger to write in a buffer so we can check the output
				source := `include::../../test/includes/tag-include-unclosed.adoc[tag=unclosed]`
				expected := `
content


end`
				Expect(PreparseDocument(source)).To(Equal(expected))
			})

			It("with unknown tag", func() {
				source := `include::../../test/includes/tag-include.adoc[tag=unknown]`
				// when/then
				_, err := PreparseDocument(source)
				// verify error in logs
				Expect(err).To(MatchError("Unresolved directive in test.adoc - include::../../test/includes/tag-include.adoc[tag=unknown]: tag 'unknown' not found in file to include"))
			})

			It("with no tag", func() {
				source := `include::../../test/includes/tag-include.adoc[]`
				expected := `== Section 1

content


end`
				Expect(PreparseDocument(source)).To(Equal(expected))
			})

			Context("permutations", func() {

				It("all lines", func() {
					source := `include::../../test/includes/tag-include.adoc[tag=**]` // includes all content except lines with tags
					expected := `== Section 1

content


end`
					Expect(PreparseDocument(source)).To(Equal(expected))
				})

				It("all tagged regions", func() {
					source := `include::../../test/includes/tag-include.adoc[tag=*]` // includes all regions (ie, skip last `end` line)
					expected := `== Section 1

content
`
					Expect(PreparseDocument(source)).To(Equal(expected))
				})

				It("all the lines outside and inside of tagged regions", func() {
					source := `include::../../test/includes/tag-include.adoc[tag=**;*]` // includes all regions
					expected := `== Section 1

content


end`
					Expect(PreparseDocument(source)).To(Equal(expected))
				})

				It("regions tagged doc, but not nested regions tagged content", func() {
					source := `include::../../test/includes/tag-include.adoc[tag=doc;!content]` // includes all `doc` but `content` nested region
					expected := `== Section 1
`
					Expect(PreparseDocument(source)).To(Equal(expected))
				})

				It("all tagged regions, but excludes any regions tagged content", func() {
					source := `include::../../test/includes/tag-include.adoc[tag=*;!content]` // includes all regions but `content` (and `end`)
					expected := `== Section 1
`
					Expect(PreparseDocument(source)).To(Equal(expected))
				})

				It("all tagged regions, but excludes any regions tagged content", func() {
					source := `include::../../test/includes/tag-include.adoc[tag=**;!content]` // includes all lines but `content`
					expected := `== Section 1


end`
					Expect(PreparseDocument(source)).To(Equal(expected))
				})

				It("**;!* — selects only the regions of the document outside of tags", func() {
					source := `include::../../test/includes/tag-include.adoc[tag=**;!*]` // excludes all regions
					expected := `
end`
					Expect(PreparseDocument(source)).To(Equal(expected))
				})
			})

		})

		Context("with missing file to include", func() {

			var wd string
			BeforeEach(func() {
				var err error
				wd, err = os.Getwd()
				Expect(err).NotTo(HaveOccurred())
			})

			It("should fail if directory does not exist in standalone block", func() {
				source := `include::{unknown}/unknown.adoc[leveloffset=+1]`
				_, err := PreparseDocument(source)
				GinkgoT().Log(err.Error())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring(fmt.Sprintf("Unresolved directive in test.adoc - include::{unknown}/unknown.adoc[leveloffset=+1]: chdir %s", filepath.Join(wd, "{unknown}"))))
			})

			It("should fail if file is missing in standalone block", func() {
				source := `include::unknown.adoc[leveloffset=+1]`
				_, err := PreparseDocument(source)
				GinkgoT().Log(err.Error())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring(fmt.Sprintf("Unresolved directive in test.adoc - include::unknown.adoc[leveloffset=+1]: open %s", filepath.Join(wd, "unknown.adoc"))))
			})

			It("should fail if file with attribute in path is not resolved in standalone block", func() {
				source := `include::{includedir}/unknown.adoc[leveloffset=+1]`
				_, err := PreparseDocument(source)
				GinkgoT().Log(err.Error())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring(fmt.Sprintf("Unresolved directive in test.adoc - include::{includedir}/unknown.adoc[leveloffset=+1]: chdir %s", filepath.Join(wd, "{includedir}"))))
			})

			It("should fail if file is missing in delimited block", func() {
				source := `----
include::../../test/includes/unknown.adoc[leveloffset=+1]
----`
				_, err := PreparseDocument(source)
				GinkgoT().Log(err.Error())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring(fmt.Sprintf("Unresolved directive in test.adoc - include::../../test/includes/unknown.adoc[leveloffset=+1]: open %s", filepath.Join(wd, "../../test/includes/unknown.adoc"))))
			})

			It("should fail if file with attribute in path is not resolved in delimited block", func() {
				source := `----
include::{includedir}/unknown.adoc[leveloffset=+1]
----`
				_, err := PreparseDocument(source)
				GinkgoT().Log(err.Error())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring(fmt.Sprintf("Unresolved directive in test.adoc - include::{includedir}/unknown.adoc[leveloffset=+1]: chdir %s", filepath.Join(wd, "{includedir}"))))
			})
		})

		Context("with inclusion with attribute in path", func() {

			It("should resolve path with attribute in standalone block from local file", func() {
				source := `:includedir: ../../test/includes

include::{includedir}/grandchild-include.adoc[]`
				expected := `:includedir: ../../test/includes

== grandchild title

first line of grandchild

last line of grandchild`
				Expect(PreparseDocument(source)).To(Equal(expected))
			})

			It("should resolve path with attribute in delimited block", func() {
				source := `:includedir: ../../test/includes

----
include::{includedir}/grandchild-include.adoc[]
----`
				expected := `:includedir: ../../test/includes

----
== grandchild title

first line of grandchild

last line of grandchild
----`
				Expect(PreparseDocument(source)).To(Equal(expected))
			})
		})

		Context("inclusion of non-asciidoc file", func() {

			It("include go file without any range in listing block", func() {

				source := `----
include::../../test/includes/hello_world.go.txt[] 
----`
				expected := `----
package includes

import "fmt"

func helloworld() {
	fmt.Println("hello, world!")
}
----`
				Expect(PreparseDocument(source)).To(Equal(expected))
			})

			It("include go file with a simple range in listing block", func() {

				source := `----
include::../../test/includes/hello_world.go.txt[lines=1] 
----`
				expected := `----
package includes
----`
				Expect(PreparseDocument(source)).To(Equal(expected))
			})
		})

	})

	Context("in final documents", func() {

		It("should include adoc file without leveloffset from local file", func() {
			logs, reset := ConfigureLogger(log.WarnLevel)
			defer reset()
			source := "include::../../test/includes/chapter-a.adoc[]"
			expected := &types.Document{
				Elements: []interface{}{
					&types.DocumentHeader{
						Title: []interface{}{
							&types.StringElement{
								Content: "Chapter A",
							},
						},
					},
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "content",
							},
						},
					},
				},
			}
			result, err := ParseDocument(source)
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(MatchDocument(expected))
			// verify no error/warning in logs
			Expect(logs).ToNot(ContainAnyMessageWithLevels(log.ErrorLevel, log.WarnLevel))
		})

		It("should include adoc file with leveloffset", func() {
			logs, reset := ConfigureLogger(log.WarnLevel)
			defer reset()
			source := "include::../../test/includes/chapter-a.adoc[leveloffset=+1]"
			title := []interface{}{
				&types.StringElement{
					Content: "Chapter A",
				},
			}
			expected := &types.Document{
				Elements: []interface{}{
					&types.Section{
						Attributes: types.Attributes{
							types.AttrID: "_chapter_a",
						},
						Level: 1, // offset by +1
						Title: title,
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "content",
									},
								},
							},
						},
					},
				},
				ElementReferences: types.ElementReferences{
					"_chapter_a": title,
				},
				TableOfContents: &types.TableOfContents{
					MaxDepth: 2,
					Sections: []*types.ToCSection{
						{
							ID:    "_chapter_a",
							Level: 1,
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
			// verify no error/warning in logs
			Expect(logs).ToNot(ContainAnyMessageWithLevels(log.ErrorLevel, log.WarnLevel))
		})

		It("should include file with attribute in path", func() {
			source := `:includedir: ../../test/includes

include::{includedir}/chapter-a.adoc[]`
			expected := &types.Document{
				Elements: []interface{}{
					&types.DocumentHeader{
						Title: []interface{}{
							&types.StringElement{
								Content: "Chapter A",
							},
						},
						Elements: []interface{}{
							&types.AttributeDeclaration{
								Name:  "includedir",
								Value: "../../test/includes",
							},
						},
					},
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "content",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should not further process with non-asciidoc files", func() {
			source := `:includedir: ../../test/includes

include::{includedir}/include.foo[]`
			expected := &types.Document{
				Elements: []interface{}{
					&types.DocumentHeader{
						Elements: []interface{}{
							&types.AttributeDeclaration{
								Name:  "includedir",
								Value: "../../test/includes",
							},
						},
					},
					&types.Paragraph{
						Elements: []interface{}{
							&types.QuotedText{
								Kind: types.SingleQuoteBold,
								Elements: []interface{}{
									&types.StringElement{
										Content: "some strong content",
									},
								},
							},
						},
					},
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: `include::hello_world.go.txt[]`,
							},
						},
					},
				},
			}
			// Expect(ParseDocument(source, WithFilename("foo.bar"))).To(MatchDocumentFragments(expected)) // parent doc may not need to be a '.adoc'
			result, err := ParseDocument(source)
			Expect(err).NotTo(HaveOccurred())
			Expect(result).To(MatchDocument(expected)) // parent doc may not need to be a '.adoc'
		})

		It("should include grandchild content without offset", func() {
			source := `include::../../test/includes/grandchild-include.adoc[]`
			title := []interface{}{
				&types.StringElement{
					Content: "grandchild title",
				},
			}
			expected := &types.Document{
				Elements: []interface{}{
					&types.Section{
						Attributes: types.Attributes{
							types.AttrID: "_grandchild_title",
						},
						Level: 1,
						Title: title,
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "first line of grandchild",
									},
								},
							},
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "last line of grandchild",
									},
								},
							},
						},
					},
				},
				ElementReferences: types.ElementReferences{
					"_grandchild_title": title,
				},
				TableOfContents: &types.TableOfContents{
					MaxDepth: 2,
					Sections: []*types.ToCSection{
						{
							ID:    "_grandchild_title",
							Level: 1,
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should include grandchild content with relative offset", func() {
			source := `include::../../test/includes/grandchild-include.adoc[leveloffset=+1]`
			title := []interface{}{
				&types.StringElement{
					Content: "grandchild title",
				},
			}
			expected := &types.Document{
				Elements: []interface{}{
					&types.Section{
						Attributes: types.Attributes{
							types.AttrID: "_grandchild_title",
						},
						Level: 2,
						Title: title,
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "first line of grandchild",
									},
								},
							},
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "last line of grandchild",
									},
								},
							},
						},
					},
				},
				ElementReferences: types.ElementReferences{
					"_grandchild_title": title,
				},
				TableOfContents: &types.TableOfContents{
					MaxDepth: 2,
					Sections: []*types.ToCSection{
						{
							ID:    "_grandchild_title",
							Level: 2,
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should include grandchild content with absolute offset", func() {
			source := `include::../../test/includes/grandchild-include.adoc[leveloffset=0]`
			title := []interface{}{
				&types.StringElement{
					Content: "grandchild title",
				},
			}
			expected := &types.Document{
				Elements: []interface{}{
					&types.DocumentHeader{
						Title: title,
					},
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "first line of grandchild",
							},
						},
					},
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "last line of grandchild",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should include child and grandchild content with relative level offset", func() {
			source := `include::../../test/includes/parent-include-relative-offset.adoc[leveloffset=+1]`
			parentTitle := []interface{}{
				&types.StringElement{
					Content: "parent title",
				},
			}
			childSection1Title := []interface{}{
				&types.StringElement{
					Content: "child section 1",
				},
			}
			childSection2Title := []interface{}{
				&types.StringElement{
					Content: "child section 2",
				},
			}
			grandchildTitle := []interface{}{
				&types.StringElement{
					Content: "grandchild title",
				},
			}
			expected := &types.Document{
				Elements: []interface{}{
					&types.Section{
						Attributes: types.Attributes{
							types.AttrID: "_parent_title",
						},
						Level: 1, // here the level is changed from `0` to `1` since `root` doc has a `leveloffset=+1` during its inclusion
						Title: parentTitle,
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "first line of parent",
									},
								},
							},
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "child preamble",
									},
								},
							},
							&types.Section{
								Attributes: types.Attributes{
									types.AttrID: "_child_section_1",
								},
								Level: 3, // here the level is changed from `1` to `3` since both `root` and `parent` docs have a `leveloffset=+1` during their inclusion
								Title: childSection1Title,
								Elements: []interface{}{
									&types.Paragraph{
										Elements: []interface{}{
											&types.StringElement{
												Content: "first line of child",
											},
										},
									},
									&types.Section{
										Attributes: types.Attributes{
											types.AttrID: "_grandchild_title",
										},
										Level: 4, // here the level is changed from `1` to `4` since both `root`, `parent` and `child` docs have a `leveloffset=+1` during their inclusion
										Title: grandchildTitle,
										Elements: []interface{}{
											&types.Paragraph{
												Elements: []interface{}{
													&types.StringElement{
														Content: "first line of grandchild",
													},
												},
											},
											&types.Paragraph{
												Elements: []interface{}{
													&types.StringElement{
														Content: "last line of grandchild",
													},
												},
											},
										},
									},
									&types.Section{
										Attributes: types.Attributes{
											types.AttrID: "_child_section_2",
										},
										Level: 4, // here the level is changed from `2` to `4` since both `root` and `parent` docs have a `leveloffset=+1` during their inclusion
										Title: childSection2Title,
										Elements: []interface{}{
											&types.Paragraph{
												Elements: []interface{}{
													&types.StringElement{
														Content: "last line of child",
													},
												},
											},
											&types.Paragraph{
												Elements: []interface{}{
													&types.StringElement{
														Content: "last line of parent",
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
				ElementReferences: types.ElementReferences{
					"_parent_title":     parentTitle,
					"_child_section_1":  childSection1Title,
					"_child_section_2":  childSection2Title,
					"_grandchild_title": grandchildTitle,
				},
				TableOfContents: &types.TableOfContents{
					MaxDepth: 2,
					Sections: []*types.ToCSection{
						{
							ID:    "_parent_title",
							Level: 1,
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should include child and grandchild content with relative then absolute level offset", func() {
			source := `include::../../test/includes/parent-include-absolute-offset.adoc[leveloffset=+1]`
			parentTitle := []interface{}{
				&types.StringElement{
					Content: "parent title",
				},
			}
			childSection1Title := []interface{}{
				&types.StringElement{
					Content: "child section 1",
				},
			}
			childSection2Title := []interface{}{
				&types.StringElement{
					Content: "child section 2",
				},
			}
			grandchildTitle := []interface{}{
				&types.StringElement{
					Content: "grandchild title",
				},
			}
			expected := &types.Document{
				Elements: []interface{}{
					&types.Section{
						Attributes: types.Attributes{
							types.AttrID: "_parent_title",
						},
						Level: 1, // here the level is changed from `0` to `1` since `root` doc has a `leveloffset=+1` during its inclusion
						Title: parentTitle,
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "first line of parent",
									},
								},
							},
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "child preamble",
									},
								},
							},
							&types.Section{
								Attributes: types.Attributes{
									types.AttrID: "_child_section_1",
								},
								Level: 3, // here level is forced to "absolute 3"
								Title: childSection1Title,
								Elements: []interface{}{
									&types.Paragraph{
										Elements: []interface{}{
											&types.StringElement{
												Content: "first line of child",
											},
										},
									},
									&types.Section{
										Attributes: types.Attributes{
											types.AttrID: "_grandchild_title",
										},
										Level: 4, // here the level is set to `4` because it the first section was moved from `1` to `3` so we use the same offset here
										Title: grandchildTitle,
										Elements: []interface{}{
											&types.Paragraph{
												Elements: []interface{}{
													&types.StringElement{
														Content: "first line of grandchild",
													},
												},
											},
											&types.Paragraph{
												Elements: []interface{}{
													&types.StringElement{
														Content: "last line of grandchild",
													},
												},
											},
										},
									},
									&types.Section{
										Attributes: types.Attributes{
											types.AttrID: "_child_section_2",
										},
										Level: 4, // here the level is changed from `2` to `4` since both `root` and `parent` docs have a `leveloffset=+1` during their inclusion
										Title: childSection2Title,
										Elements: []interface{}{
											&types.Paragraph{
												Elements: []interface{}{
													&types.StringElement{
														Content: "last line of child",
													},
												},
											},
											&types.Paragraph{
												Elements: []interface{}{
													&types.StringElement{
														Content: "last line of parent",
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
				ElementReferences: types.ElementReferences{
					"_parent_title":     parentTitle,
					"_child_section_1":  childSection1Title,
					"_child_section_2":  childSection2Title,
					"_grandchild_title": grandchildTitle,
				},
				TableOfContents: &types.TableOfContents{
					MaxDepth: 2,
					Sections: []*types.ToCSection{
						{
							ID:    "_parent_title",
							Level: 1,
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		It("should include adoc with attributes file within main content", func() {
			source := `include::../../test/includes/attributes.adoc[]`
			expected := &types.Document{
				Elements: []interface{}{
					&types.DocumentHeader{
						Elements: []interface{}{
							&types.AttributeDeclaration{
								Name:  "author",
								Value: "Xavier",
							},
							&types.AttributeDeclaration{
								Name:  "leveloffset",
								Value: "+1",
							},
						},
					},
					&types.Paragraph{
						Elements: []interface{}{
							&types.StringElement{
								Content: "some content",
							},
						},
					},
				},
			}
			Expect(ParseDocument(source)).To(MatchDocument(expected))
		})

		Context("with delimited blocks", func() {

			It("should include adoc with attributes file within commment block", func() {
				source := `////
include::../../test/includes/attributes.adoc[]
////`
				expected := &types.Document{}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("should include adoc with attributes file within example block", func() {
				source := `====
include::../../test/includes/attributes.adoc[]
====`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Example,
							Elements: []interface{}{
								&types.AttributeDeclaration{
									Name:  "author",
									Value: "Xavier",
								},
								&types.AttributeDeclaration{
									Name:  "leveloffset",
									Value: "+1",
								},
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											Content: `some content`,
										},
									},
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("should include adoc with attributes file within listing block", func() {
				source := `----
include::../../test/includes/attributes.adoc[]
----`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								&types.StringElement{
									Content: `:author: Xavier
:leveloffset: +1

some content`,
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("should include adoc file within fenced block", func() {
				source := "```\n" +
					"include::../../test/includes/parent-include.adoc[]\n" +
					"```\n" +
					"<1> a callout"
				// include the doc without parsing the elements (besides the file inclusions)
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Fenced,
							Elements: []interface{}{
								&types.StringElement{
									Content: `:leveloffset: +1

= parent title

first line of parent

= child title

first line of child

== grandchild title

first line of grandchild

last line of grandchild

last line of child

last line of parent `,
								},
								&types.Callout{
									Ref: 1,
								},
								&types.StringElement{
									Content: "\n\n:leveloffset!:",
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
													Content: "a callout",
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

			It("should include adoc file within quote block", func() {
				source := `____
include::../../test/includes/parent-include.adoc[]
____`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Quote,
							Elements: []interface{}{
								&types.AttributeDeclaration{
									Name:  "leveloffset",
									Value: string("+1"),
								},
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											Content: "= parent title",
										},
									},
								},
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											Content: "first line of parent",
										},
									},
								},
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											Content: "= child title",
										},
									},
								},
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											Content: "first line of child",
										},
									},
								},
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											Content: "== grandchild title",
										},
									},
								},
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											Content: "first line of grandchild",
										},
									},
								},
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											Content: "last line of grandchild",
										},
									},
								},
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											Content: "last line of child",
										},
									},
								},
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											Content: "last line of parent ",
										},
										&types.SpecialCharacter{
											Name: "<",
										},
										&types.StringElement{
											Content: "1",
										},
										&types.SpecialCharacter{
											Name: ">",
										},
									},
								},
								&types.AttributeReset{
									Name: "leveloffset",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

		})

		Context("with line ranges", func() {

			Context("unquoted", func() {

				It("with single unquoted line", func() {
					source := `include::../../test/includes/chapter-a.adoc[lines=1]`
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: []interface{}{
									&types.StringElement{
										Content: "Chapter A",
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("with multiple unquoted lines", func() {
					source := `include::../../test/includes/chapter-a.adoc[lines=1..3]`
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: []interface{}{
									&types.StringElement{
										Content: "Chapter A",
									},
								},
							},
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "content",
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("with multiple unquoted ranges (becoming authors)", func() {
					source := `include::../../test/includes/chapter-a.adoc[lines=1;3..4;6..-1]` // paragraph becomes the author since the in-between blank line is stripped out
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: []interface{}{
									&types.StringElement{
										Content: "Chapter A",
									},
								},
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name: types.AttrAuthors,
										Value: types.DocumentAuthors{
											{
												DocumentAuthorFullName: &types.DocumentAuthorFullName{
													FirstName: "content",
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

				It("with invalid unquoted range - case 1", func() {
					source := `include::../../test/includes/chapter-a.adoc[lines=1;3..4;6..foo]` // not a number
					_, err := ParseDocument(source)
					GinkgoT().Logf(err.Error())
					Expect(err).To(MatchError(`Unresolved directive in test.adoc - include::../../test/includes/chapter-a.adoc[lines=1;3..4;6..foo]: 1:11 (10): no match found, expected: "-" or [0-9]`))
				})

				It("with invalid unquoted range - case 2", func() {
					source := `include::../../test/includes/chapter-a.adoc[lines=1,3..4,6..-1]` // using commas instead of semi-colons
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: []interface{}{
									&types.StringElement{
										Content: "Chapter A",
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})
			})

			Context("quoted", func() {

				It("with single line", func() {
					logs, reset := ConfigureLogger(log.WarnLevel)
					defer reset()
					source := `include::../../test/includes/chapter-a.adoc[lines="1"]`
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: []interface{}{
									&types.StringElement{
										Content: "Chapter A",
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
					// verify no error/warning in logs
					Expect(logs).ToNot(ContainAnyMessageWithLevels(log.ErrorLevel, log.WarnLevel))
				})

				It("with multiple lines", func() {
					source := `include::../../test/includes/chapter-a.adoc[lines="1..2"]`
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: []interface{}{
									&types.StringElement{
										Content: "Chapter A",
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("with multiple ranges with colons (becoming authors)", func() {
					// here, the `content` paragraph gets attached to the header and becomes the author
					source := `include::../../test/includes/chapter-a.adoc[lines="1,3..4,6..-1"]`
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: []interface{}{
									&types.StringElement{
										Content: "Chapter A",
									},
								},
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name: types.AttrAuthors,
										Value: types.DocumentAuthors{
											{
												DocumentAuthorFullName: &types.DocumentAuthorFullName{
													FirstName: "content",
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

				It("with multiple ranges with semicolons (becoming authors)", func() {
					// here, the `content` paragraph gets attached to the header and becomes the author
					source := `include::../../test/includes/chapter-a.adoc[lines="1;3..4;6..-1"]`
					expected := &types.Document{
						Elements: []interface{}{
							&types.DocumentHeader{
								Title: []interface{}{
									&types.StringElement{
										Content: "Chapter A",
									},
								},
								Elements: []interface{}{
									&types.AttributeDeclaration{
										Name: types.AttrAuthors,
										Value: types.DocumentAuthors{
											{
												DocumentAuthorFullName: &types.DocumentAuthorFullName{
													FirstName: "content",
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

				It("with invalid range - case 1", func() {
					source := `include::../../test/includes/chapter-a.adoc[lines="1,3..4,6..foo"]` // not a number
					_, err := ParseDocument(source)
					GinkgoT().Logf(err.Error())
					Expect(err).To(MatchError(`Unresolved directive in test.adoc - include::../../test/includes/chapter-a.adoc[lines="1,3..4,6..foo"]: 1:11 (10): no match found, expected: "-" or [0-9]`))
				})

				It("with ignored tags", func() {
					// include using a line range a file having tags
					source := `include::../../test/includes/tag-include.adoc[lines=3]`
					title := []interface{}{
						&types.StringElement{
							Content: "Section 1",
						},
					}
					expected := &types.Document{
						Elements: []interface{}{
							&types.Section{
								Attributes: types.Attributes{
									types.AttrID: "_section_1",
								},
								Level: 1,
								Title: title,
							},
						},
						ElementReferences: types.ElementReferences{
							"_section_1": title,
						},
						TableOfContents: &types.TableOfContents{
							MaxDepth: 2,
							Sections: []*types.ToCSection{
								{
									ID:    "_section_1",
									Level: 1,
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})
			})
		})

		Context("with tag ranges", func() {

			It("with single tag", func() {
				logs, reset := ConfigureLogger(log.WarnLevel)
				defer reset()
				source := `include::../../test/includes/tag-include.adoc[tag=section]`
				title := []interface{}{
					&types.StringElement{
						Content: "Section 1",
					},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_section_1",
							},
							Level: 1,
							Title: title,
						},
					},
					ElementReferences: types.ElementReferences{
						"_section_1": title,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_section_1",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
				// verify no error/warning in logs
				Expect(logs).ToNot(ContainAnyMessageWithLevels(log.ErrorLevel, log.WarnLevel))
			})

			It("with surrounding tag", func() {
				logs, reset := ConfigureLogger(log.WarnLevel)
				defer reset()
				source := `include::../../test/includes/tag-include.adoc[tag=doc]`
				title := []interface{}{
					&types.StringElement{
						Content: "Section 1",
					},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_section_1",
							},
							Level: 1,
							Title: title,
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											Content: "content",
										},
									},
								},
							},
						},
					},
					ElementReferences: types.ElementReferences{
						"_section_1": title,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_section_1",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
				// verify no error/warning in logs
				Expect(logs).ToNot(ContainAnyMessageWithLevels(log.ErrorLevel, log.WarnLevel))
			})

			It("with unclosed tag", func() {
				// setup logger to write in a buffer so we can check the output
				logs, reset := ConfigureLogger(log.WarnLevel)
				defer reset()
				source := `include::../../test/includes/tag-include-unclosed.adoc[tag=unclosed]`
				expected := &types.Document{
					Elements: []interface{}{
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "content",
								},
							},
						},
						&types.Paragraph{
							Elements: []interface{}{
								&types.StringElement{
									Content: "end",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
				// verify error in logs
				Expect(logs).To(ContainJSONLog(log.WarnLevel, "detected unclosed tag 'unclosed' starting at line 6 of include file: ../../test/includes/tag-include-unclosed.adoc"))
			})

			It("with unknown tag", func() {
				// given
				source := `include::../../test/includes/tag-include.adoc[tag=unknown]`
				// when/then
				_, err := ParseDocument(source)
				// verify error
				Expect(err).To(MatchError("Unresolved directive in test.adoc - include::../../test/includes/tag-include.adoc[tag=unknown]: tag 'unknown' not found in file to include"))
			})

			It("with no tag", func() {
				source := `include::../../test/includes/tag-include.adoc[]`
				title := []interface{}{
					&types.StringElement{
						Content: "Section 1",
					},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_section_1",
							},
							Level: 1,
							Title: title,
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											Content: "content",
										},
									},
								},
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											Content: "end",
										},
									},
								},
							},
						},
					},
					ElementReferences: types.ElementReferences{
						"_section_1": title,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_section_1",
								Level: 1,
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			Context("permutations", func() {

				It("all lines", func() {
					source := `include::../../test/includes/tag-include.adoc[tag=**]` // includes all content except lines with tags
					title := []interface{}{
						&types.StringElement{
							Content: "Section 1",
						},
					}
					expected := &types.Document{
						Elements: []interface{}{
							&types.Section{
								Attributes: types.Attributes{
									types.AttrID: "_section_1",
								},
								Level: 1,
								Title: title,
								Elements: []interface{}{
									&types.Paragraph{
										Elements: []interface{}{
											&types.StringElement{
												Content: "content",
											},
										},
									},
									&types.Paragraph{
										Elements: []interface{}{
											&types.StringElement{
												Content: "end",
											},
										},
									},
								},
							},
						},
						ElementReferences: types.ElementReferences{
							"_section_1": title,
						},
						TableOfContents: &types.TableOfContents{
							MaxDepth: 2,
							Sections: []*types.ToCSection{
								{
									ID:    "_section_1",
									Level: 1,
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("all tagged regions", func() {
					source := `include::../../test/includes/tag-include.adoc[tag=*]` // includes all regions
					title := []interface{}{
						&types.StringElement{
							Content: "Section 1",
						},
					}
					expected := &types.Document{
						Elements: []interface{}{
							&types.Section{
								Attributes: types.Attributes{
									types.AttrID: "_section_1",
								},
								Level: 1,
								Title: title,
								Elements: []interface{}{
									&types.Paragraph{
										Elements: []interface{}{
											&types.StringElement{
												Content: "content",
											},
										},
									},
								},
							},
						},
						ElementReferences: types.ElementReferences{
							"_section_1": title,
						},
						TableOfContents: &types.TableOfContents{
							MaxDepth: 2,
							Sections: []*types.ToCSection{
								{
									ID:    "_section_1",
									Level: 1,
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("all the lines outside and inside of tagged regions", func() {
					source := `include::../../test/includes/tag-include.adoc[tag=**;*]` // includes all regions
					title := []interface{}{
						&types.StringElement{
							Content: "Section 1",
						},
					}
					expected := &types.Document{
						Elements: []interface{}{
							&types.Section{
								Attributes: types.Attributes{
									types.AttrID: "_section_1",
								},
								Level: 1,
								Title: title,
								Elements: []interface{}{
									&types.Paragraph{
										Elements: []interface{}{
											&types.StringElement{
												Content: "content",
											},
										},
									},
									&types.Paragraph{
										Elements: []interface{}{
											&types.StringElement{
												Content: "end",
											},
										},
									},
								},
							},
						},
						ElementReferences: types.ElementReferences{
							"_section_1": title,
						},
						TableOfContents: &types.TableOfContents{
							MaxDepth: 2,
							Sections: []*types.ToCSection{
								{
									ID:    "_section_1",
									Level: 1,
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("regions tagged doc, but not nested regions tagged content", func() {
					source := `include::../../test/includes/tag-include.adoc[tag=doc;!content]` // includes all `doc` but `content`
					title := []interface{}{
						&types.StringElement{
							Content: "Section 1",
						},
					}
					expected := &types.Document{
						Elements: []interface{}{
							&types.Section{
								Attributes: types.Attributes{
									types.AttrID: "_section_1",
								},
								Level: 1,
								Title: title,
							},
						},
						ElementReferences: types.ElementReferences{
							"_section_1": title,
						},
						TableOfContents: &types.TableOfContents{
							MaxDepth: 2,
							Sections: []*types.ToCSection{
								{
									ID:    "_section_1",
									Level: 1,
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("all tagged regions, but excludes any regions tagged content", func() {
					source := `include::../../test/includes/tag-include.adoc[tag=*;!content]` // includes all but `content`
					title := []interface{}{
						&types.StringElement{
							Content: "Section 1",
						},
					}
					expected := &types.Document{
						Elements: []interface{}{
							&types.Section{
								Attributes: types.Attributes{
									types.AttrID: "_section_1",
								},
								Level: 1,
								Title: title,
							},
						},
						ElementReferences: types.ElementReferences{
							"_section_1": title,
						},
						TableOfContents: &types.TableOfContents{
							MaxDepth: 2,
							Sections: []*types.ToCSection{
								{
									ID:    "_section_1",
									Level: 1,
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("all tagged regions, but excludes any regions tagged content", func() {
					source := `include::../../test/includes/tag-include.adoc[tag=**;!content]` // includes all lines but `content`
					title := []interface{}{
						&types.StringElement{
							Content: "Section 1",
						},
					}
					expected := &types.Document{
						Elements: []interface{}{
							&types.Section{
								Attributes: types.Attributes{
									types.AttrID: "_section_1",
								},
								Level: 1,
								Title: title,
								Elements: []interface{}{
									&types.Paragraph{
										Elements: []interface{}{
											&types.StringElement{
												Content: "end",
											},
										},
									},
								},
							},
						},
						ElementReferences: types.ElementReferences{
							"_section_1": title,
						},
						TableOfContents: &types.TableOfContents{
							MaxDepth: 2,
							Sections: []*types.ToCSection{
								{
									ID:    "_section_1",
									Level: 1,
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})

				It("**;!* — selects only the regions of the document outside of tags", func() {
					source := `include::../../test/includes/tag-include.adoc[tag=**;!*]` // excludes all regions
					expected := &types.Document{
						Elements: []interface{}{
							&types.Paragraph{
								Elements: []interface{}{
									&types.StringElement{
										Content: "end",
									},
								},
							},
						},
					}
					Expect(ParseDocument(source)).To(MatchDocument(expected))
				})
			})
		})

		Context("with missing file to include", func() {

			var wd string
			BeforeEach(func() {
				var err error
				wd, err = os.Getwd()
				Expect(err).NotTo(HaveOccurred())
			})

			It("should fail if directory does not exist in standalone block", func() {
				source := `include::{unknown}/unknown.adoc[leveloffset=+1]`
				_, err := ParseDocument(source)
				GinkgoT().Log(err.Error())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring(fmt.Sprintf("Unresolved directive in test.adoc - include::{unknown}/unknown.adoc[leveloffset=+1]: chdir %s", filepath.Join(wd, "{unknown}"))))
			})

			It("should fail if file is missing in standalone block", func() {
				source := `include::unknown.adoc[leveloffset=+1]`
				_, err := ParseDocument(source)
				GinkgoT().Log(err.Error())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring(fmt.Sprintf("Unresolved directive in test.adoc - include::unknown.adoc[leveloffset=+1]: open %s", filepath.Join(wd, "unknown.adoc"))))
			})

			It("should fail if file with attribute in path is not resolved in standalone block", func() {
				source := `include::{includedir}/unknown.adoc[leveloffset=+1]`
				_, err := ParseDocument(source)
				GinkgoT().Log(err.Error())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring(fmt.Sprintf("Unresolved directive in test.adoc - include::{includedir}/unknown.adoc[leveloffset=+1]: chdir %s", filepath.Join(wd, "{includedir}"))))
			})

			It("should fail if file is missing in delimited block", func() {
				source := `----
include::../../test/includes/unknown.adoc[leveloffset=+1]
----`
				_, err := ParseDocument(source)
				GinkgoT().Log(err.Error())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring(fmt.Sprintf("Unresolved directive in test.adoc - include::../../test/includes/unknown.adoc[leveloffset=+1]: open %s", filepath.Join(wd, "../../test/includes/unknown.adoc"))))
			})

			It("should fail if file with attribute in path is not resolved in delimited block", func() {
				source := `----
include::{includedir}/unknown.adoc[leveloffset=+1]
----`
				_, err := ParseDocument(source)
				GinkgoT().Log(err.Error())
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring(fmt.Sprintf("Unresolved directive in test.adoc - include::{includedir}/unknown.adoc[leveloffset=+1]: chdir %s", filepath.Join(wd, "{includedir}"))))
			})
		})

		Context("with inclusion with attribute in path", func() {

			It("should resolve path with attribute in standalone block from local file", func() {
				source := `:includedir: ../../test/includes
			
include::{includedir}/grandchild-include.adoc[]`
				title := []interface{}{
					&types.StringElement{
						Content: "grandchild title",
					},
				}
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Elements: []interface{}{
								&types.AttributeDeclaration{
									Name:  "includedir",
									Value: "../../test/includes",
								},
							},
						},
						&types.Section{
							Attributes: types.Attributes{
								types.AttrID: "_grandchild_title",
							},
							Level: 1,
							Title: title,
							Elements: []interface{}{
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											Content: "first line of grandchild",
										},
									},
								},
								&types.Paragraph{
									Elements: []interface{}{
										&types.StringElement{
											Content: "last line of grandchild",
										},
									},
								},
							},
						},
					},
					ElementReferences: types.ElementReferences{
						"_grandchild_title": title,
					},
					TableOfContents: &types.TableOfContents{
						MaxDepth: 2,
						Sections: []*types.ToCSection{
							{
								ID:    "_grandchild_title",
								Level: 1,
							},
						},
					},
				}
				// Expect(ParseDocument(source, WithFilename("foo.adoc"))).To(MatchDocument(expected))
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("should resolve path with attribute in delimited block", func() {
				source := `:includedir: ../../test/includes

----
include::{includedir}/grandchild-include.adoc[]
----`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DocumentHeader{
							Elements: []interface{}{
								&types.AttributeDeclaration{
									Name:  "includedir",
									Value: "../../test/includes",
								},
							},
						},
						&types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								&types.StringElement{
									Content: "== grandchild title\n\nfirst line of grandchild\n\nlast line of grandchild",
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})
		})

		Context("inclusion of non-asciidoc file", func() {

			It("include go file without any range in listing block", func() {

				source := `----
include::../../test/includes/hello_world.go.txt[] 
----`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								&types.StringElement{
									Content: `package includes

import "fmt"

func helloworld() {
	fmt.Println("hello, world!")
}`,
								},
							},
						},
					},
				}
				Expect(ParseDocument(source)).To(MatchDocument(expected))
			})

			It("include go file with a simple range in listing block", func() {

				source := `----
include::../../test/includes/hello_world.go.txt[lines=1] 
----`
				expected := &types.Document{
					Elements: []interface{}{
						&types.DelimitedBlock{
							Kind: types.Listing,
							Elements: []interface{}{
								&types.StringElement{
									Content: `package includes`,
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
