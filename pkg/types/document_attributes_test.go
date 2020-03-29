package types_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("document attributes", func() {

	Context("regular attributes", func() {

		It("normal value", func() {
			// given
			attributes := types.DocumentAttributes{}
			// when
			attributes.Add("foo", "bar")
			// then
			Expect("bar").To(Equal(attributes["foo"]))
		})
	})
})

// var _ = Describe("document headers", func() {

// 	XContext("authors", func() {

// 		It("single author", func() {
// 			title := []interface{}{
// 				types.StringElement{
// 					Content: "title",
// 				},
// 			}
// 			authors := []types.DocumentAuthor{
// 				{
// 					Email:    "jdoe@example.com",
// 					FullName: "john foo doe",
// 				},
// 			}
// 			// when
// 			s, err := types.NewDocumentHeader(title, authors, nil)
// 			// then
// 			Expect(err).NotTo(HaveOccurred())
// 			Expect(s.Attributes).To(HaveKeyWithValue(types.AttrAuthors, authors)) // unchanged authors
// 			Expect(s.Attributes).To(HaveKeyWithValue("firstname", "john"))
// 			Expect(s.Attributes).To(HaveKeyWithValue("middlename", "foo"))
// 			Expect(s.Attributes).To(HaveKeyWithValue("lastname", "doe"))
// 			Expect(s.Attributes).To(HaveKeyWithValue("author", "john foo doe"))
// 			Expect(s.Attributes).To(HaveKeyWithValue("authorinitials", "jfd"))
// 			Expect(s.Attributes).To(HaveKeyWithValue("email", "jdoe@example.com"))
// 		})

// 		It("single author with extra spaces", func() {
// 			title := []interface{}{
// 				types.StringElement{
// 					Content: "title",
// 				},
// 			}
// 			// when
// 			s, err := types.NewDocumentHeader(title, []types.DocumentAuthor{
// 				{
// 					Email:    "jdoe@example.com",
// 					FullName: "john  foo   doe",
// 				},
// 			}, nil)
// 			// then
// 			Expect(err).NotTo(HaveOccurred())
// 			Expect(s.Attributes).To(HaveKeyWithValue(types.AttrAuthors, []types.DocumentAuthor{
// 				{
// 					Email:    "jdoe@example.com",
// 					FullName: "john foo doe", // spaces trimmed
// 				},
// 			}))
// 			Expect(s.Attributes).To(HaveKeyWithValue("firstname", "john"))
// 			Expect(s.Attributes).To(HaveKeyWithValue("middlename", "foo"))
// 			Expect(s.Attributes).To(HaveKeyWithValue("lastname", "doe"))
// 			Expect(s.Attributes).To(HaveKeyWithValue("author", "john foo doe"))
// 			Expect(s.Attributes).To(HaveKeyWithValue("authorinitials", "jfd"))
// 			Expect(s.Attributes).To(HaveKeyWithValue("email", "jdoe@example.com"))
// 		})

// 		It("single author with underscore", func() {
// 			title := []interface{}{
// 				types.StringElement{
// 					Content: "title",
// 				},
// 			}
// 			// when
// 			s, err := types.NewDocumentHeader(title, []types.DocumentAuthor{
// 				{
// 					Email:    "jane@example.com",
// 					FullName: "Jane the_Doe",
// 				},
// 			}, nil)
// 			// then
// 			Expect(err).NotTo(HaveOccurred())
// 			Expect(s.Attributes).To(HaveKeyWithValue(types.AttrAuthors, []types.DocumentAuthor{
// 				{
// 					Email:    "jane@example.com",
// 					FullName: "Jane the Doe", // underscore replaced with a space
// 				},
// 			}))
// 			Expect(s.Attributes).To(HaveKeyWithValue("firstname", "Jane"))
// 			Expect(s.Attributes).To(HaveKeyWithValue("lastname", "the_Doe"))
// 			Expect(s.Attributes).To(HaveKeyWithValue("author", "Jane the_Doe"))
// 			Expect(s.Attributes).To(HaveKeyWithValue("authorinitials", "Jt"))
// 			Expect(s.Attributes).To(HaveKeyWithValue("email", "jane@example.com"))
// 		})

// 		It("multiple authors", func() {
// 			title := []interface{}{
// 				types.StringElement{
// 					Content: "title",
// 				},
// 			}
// 			authors := []types.DocumentAuthor{
// 				{
// 					Email:    "johndoe@example.com",
// 					FullName: "john foo doe",
// 				},
// 				{
// 					Email:    "janedoe@example.com",
// 					FullName: "jane doe",
// 				},
// 			}
// 			// when
// 			s, err := types.NewDocumentHeader(title, authors, nil)
// 			// then
// 			Expect(err).NotTo(HaveOccurred())
// 			Expect(s.Attributes).To(HaveKeyWithValue(types.AttrAuthors, authors))
// 			Expect(s.Attributes).To(HaveKeyWithValue("firstname", "john"))
// 			Expect(s.Attributes).To(HaveKeyWithValue("middlename", "foo"))
// 			Expect(s.Attributes).To(HaveKeyWithValue("lastname", "doe"))
// 			Expect(s.Attributes).To(HaveKeyWithValue("author", "john foo doe"))
// 			Expect(s.Attributes).To(HaveKeyWithValue("authorinitials", "jfd"))
// 			Expect(s.Attributes).To(HaveKeyWithValue("email", "johndoe@example.com"))
// 			Expect(s.Attributes).To(HaveKeyWithValue("firstname_2", "jane"))
// 			Expect(s.Attributes).To(HaveKeyWithValue("lastname_2", "doe"))
// 			Expect(s.Attributes).To(HaveKeyWithValue("author_2", "jane doe"))
// 			Expect(s.Attributes).To(HaveKeyWithValue("authorinitials_2", "jd"))
// 			Expect(s.Attributes).To(HaveKeyWithValue("email_2", "janedoe@example.com"))
// 			Expect(s.Attributes).NotTo(HaveKey(types.AttrRevision))

// 		})
// 	})

// 	Context("revisions", func() {

// 		It("revision", func() {
// 			title := []interface{}{
// 				types.StringElement{
// 					Content: "title",
// 				},
// 			}
// 			revision := types.DocumentRevision{
// 				Revdate:   "March 28 2020",
// 				Revnumber: "v1.0",
// 				Revremark: "testing",
// 			}
// 			// when
// 			s, err := types.NewDocumentHeader(title, nil, revision)
// 			// then
// 			Expect(err).NotTo(HaveOccurred())
// 			Expect(s.Attributes).NotTo(HaveKey(types.AttrAuthors))
// 			Expect(s.Attributes).To(HaveKeyWithValue(types.AttrRevision, revision))
// 			Expect(s.Attributes).To(HaveKeyWithValue("revdate", revision.Revdate))
// 			Expect(s.Attributes).To(HaveKeyWithValue("revnumber", revision.Revnumber))
// 			Expect(s.Attributes).To(HaveKeyWithValue("revremark", revision.Revremark))
// 		})
// 	})

// })

var _ = DescribeTable("document attribute overrides",
	func(key string, expectedValue string, expectedFound bool) {
		// given
		attributes := types.DocumentAttributesWithOverrides{
			Content: map[string]interface{}{
				"normal":   "ok",
				"override": "ok, too",
			},
			Overrides: map[string]string{
				"foo":      "cheesecake",
				"!bar":     "",
				"baz":      "",
				"override": "overridden",
			},
		}
		// when
		value, found := attributes.GetAsString(key)
		// then
		Expect(found).To(Equal(expectedFound))
		Expect(value).To(Equal(expectedValue))
	},
	Entry("normal", "normal", "ok", true),
	Entry("override", "override", "overridden", true), // entry is overridden
	Entry("foo", "foo", "cheesecake", true),
	Entry("!bar", "bar", "", false), // entry is reset
	Entry("baz", "baz", "", true),   // entry exists but its value is empty
)

var _ = DescribeTable("document attribute overrides with default",
	func(key string, expectedValue string) {
		// given
		attributes := types.DocumentAttributesWithOverrides{
			Content: map[string]interface{}{
				"normal":   "ok",
				"override": "ok, too",
			},
			Overrides: map[string]string{
				"foo":      "cheesecake",
				"!bar":     "",
				"baz":      "",
				"override": "overridden",
			},
		}
		// when
		value := attributes.GetAsStringWithDefault(key, "default")
		// then
		Expect(value).To(Equal(expectedValue))
	},
	Entry("normal", "normal", "ok"),
	Entry("override", "override", "overridden"), // entry is overridden
	Entry("foo", "foo", "cheesecake"),
	Entry("!bar", "bar", "default"), // entry is reset, default is returned
	Entry("baz", "baz", ""),         // entry exists but its value is empty
)
