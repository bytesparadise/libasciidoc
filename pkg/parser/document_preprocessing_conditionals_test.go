package parser_test

import (
	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2" // nolint:golint
	. "github.com/onsi/gomega"    // nolint:golintt
)

var _ = Describe("conditional inclusions", func() {

	Context("in preparsed documents", func() {

		Context("ifdef", func() {

			It("should not include content when var is not defined", func() {
				source := `intro content

ifdef::cookie[]
cookie content
endif::[]

closing content`
				expected := `intro content


closing content`
				Expect(PreparseDocument(source)).To(Equal(expected))
			})

			It("should include content when var is defined", func() {
				source := `:cookie:
intro content

ifdef::cookie[]
cookie content
endif::[]

closing content`
				expected := `:cookie:
intro content

cookie content

closing content`
				Expect(PreparseDocument(source)).To(Equal(expected))
			})

			It("should include all content when vars are defined", func() {
				source := `:cookie:
:chocolate:

intro content

ifdef::cookie[]
cookie content (1)

ifdef::cookie[]
chocolate content

endif::[]
cookie content (2)
endif::[]

closing content`
				expected := `:cookie:
:chocolate:

intro content

cookie content (1)

chocolate content

cookie content (2)

closing content`
				Expect(PreparseDocument(source)).To(Equal(expected))
			})

			It("should not include sub content when child var is not defined", func() {
				source := `:cookie:

intro content

ifdef::cookie[]
cookie content (1)

ifdef::chocolate[]
chocolate content

endif::[]
cookie content (2)
endif::[]

closing content`
				expected := `:cookie:

intro content

cookie content (1)

cookie content (2)

closing content`
				Expect(PreparseDocument(source)).To(Equal(expected))
			})

			Context("single-line", func() {

				It("with attribute not defined", func() {
					source := `* some content
ifdef::cookie[* conditional content]
* more content`
					expected := `* some content
* more content`

					Expect(PreparseDocument(source)).To(Equal(expected))
				})

				It("with attribute defined", func() {
					source := `:cookie:

* some content
ifdef::cookie[* conditional content]
* more content`
					expected := `:cookie:

* some content
* conditional content
* more content`
					Expect(PreparseDocument(source)).To(Equal(expected))
				})

				It("with backend attribute", func() {
					source := `ifdef::basebackend-html[* to pass through +++<b>HTML</b>+++ directly, surround the text with triple plus]`
					expected := `* to pass through +++<b>HTML</b>+++ directly, surround the text with triple plus`
					Expect(PreparseDocument(source)).To(Equal(expected))
				})

			})
		})

		Context("ifndef", func() {

			It("should include content when var is not defined", func() {
				source := `intro content

ifndef::cookie[]
cookie content
endif::[]

closing content`
				expected := `intro content

cookie content

closing content`
				Expect(PreparseDocument(source)).To(Equal(expected))
			})

			It("should not include content when var is defined", func() {
				source := `:cookie:
intro content

ifndef::cookie[]
cookie content
endif::[]

closing content`
				expected := `:cookie:
intro content


closing content`
				Expect(PreparseDocument(source)).To(Equal(expected))
			})

			It("should include no content when vars are defined", func() {
				source := `:cookie:
:chocolate:

intro content

ifndef::cookie[]
cookie content (1)

ifndef::cookie[]
chocolate content

endif::[]
cookie content (2)
endif::[]

closing content`
				expected := `:cookie:
:chocolate:

intro content


closing content`
				Expect(PreparseDocument(source)).To(Equal(expected))
			})

			It("should include sub content when parent var is not defined", func() {
				source := `:chocolate:

intro content

ifndef::cookie[]
cookie content (1)

ifndef::chocolate[]
chocolate content

endif::[]
cookie content (2)
endif::[]

closing content`
				expected := `:chocolate:

intro content

cookie content (1)

cookie content (2)

closing content`
				Expect(PreparseDocument(source)).To(Equal(expected))
			})

			Context("single-line", func() {

				It("with attribute not defined", func() {
					source := `* some content
ifndef::cookie[* conditional content]
* more content`
					expected := `* some content
* conditional content
* more content`

					Expect(PreparseDocument(source)).To(Equal(expected))
				})

				It("with attribute defined", func() {
					source := `:cookie:

* some content
ifndef::cookie[* conditional content]
* more content`
					expected := `:cookie:

* some content
* more content`
					Expect(PreparseDocument(source)).To(Equal(expected))
				})
			})
		})

		Context("ifeval", func() {

			It("with basic comparison", func() {
				source := `intro content

ifeval::[2 > 1]
conditional content
endif::[]

closing content`
				expected := `intro content

conditional content

closing content`
				Expect(PreparseDocument(source)).To(Equal(expected))
			})

			It("with string equality on attribute", func() {
				source := `intro content

ifeval::["{backend}" == "html5"]
conditional content
endif::[]

closing content`
				expected := `intro content

conditional content

closing content`
				Expect(PreparseDocument(source, configuration.WithBackEnd("html5"))).To(Equal(expected))
			})

			It("with string inequality on attribute", func() {
				source := `intro content

ifeval::["{backend}" == "html5"]
conditional content
endif::[]

closing content`
				expected := `intro content


closing content`
				Expect(PreparseDocument(source, configuration.WithBackEnd("pdf"))).To(Equal(expected))
			})

			It("with num equality on attribute", func() {
				source := `intro content

ifeval::[{sectnumlevels} == 3]
conditional content
endif::[]

closing content`
				expected := `intro content

conditional content

closing content`
				Expect(PreparseDocument(source, configuration.WithAttribute("sectnumlevels", 3))).To(Equal(expected))
			})

			It("with num inequality on attribute", func() {
				source := `intro content

ifeval::[{sectnumlevels} == 3]
conditional content
endif::[]

closing content`
				expected := `intro content


closing content`
				Expect(PreparseDocument(source, configuration.WithAttribute("sectnumlevels", "2"))).To(Equal(expected))
			})
		})

		Context("endif with attribute name", func() {

			It("should support attribute name in endif directive", func() {
				source := `:cookie:

intro content

ifdef::cookie[]
cookie content (1)

ifdef::chocolate[]
chocolate content

endif::chocolate[]
cookie content (2)
endif::cookie[]

closing content`
				expected := `:cookie:

intro content

cookie content (1)

cookie content (2)

closing content`
				Expect(PreparseDocument(source)).To(Equal(expected))
			})
		})
	})
})
