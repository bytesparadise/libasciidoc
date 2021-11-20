package parser

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("new substitutions", func() {

	Context("listing block", func() {

		It("default subs", func() {
			b := &types.DelimitedBlock{
				Kind: types.Listing,
			}
			Expect(newSubstitutions(b)).To(Equal(substitutions{
				{
					entrypoint: "VerbatimGroup",
					rules: map[substitutionKind]bool{
						SpecialCharacters: true,
						Callouts:          true,
					},
				},
			}))
		})
		Context("incremental subs", func() {

			It("-callouts", func() {
				b := &types.DelimitedBlock{
					Kind: types.Listing,
					Attributes: types.Attributes{
						types.AttrSubstitutions: "-callouts",
					},
				}
				Expect(newSubstitutions(b)).To(Equal(substitutions{
					{
						entrypoint: "VerbatimGroup",
						rules: map[substitutionKind]bool{
							SpecialCharacters: true,
						},
					},
				}))
			})

			It("-specialchars", func() {
				b := &types.DelimitedBlock{
					Kind: types.Listing,
					Attributes: types.Attributes{
						types.AttrSubstitutions: "-specialchars",
					},
				}
				Expect(newSubstitutions(b)).To(Equal(substitutions{
					{
						entrypoint: "VerbatimGroup",
						rules: map[substitutionKind]bool{
							Callouts: true,
						},
					},
				}))
			})
		})

		Context("absolute subs", func() {

			It("macros", func() {
				b := &types.DelimitedBlock{
					Kind: types.Listing,
					Attributes: types.Attributes{
						types.AttrSubstitutions: "macros",
					},
				}
				Expect(newSubstitutions(b)).To(Equal(substitutions{
					{
						entrypoint: "MacrosGroup",
						rules: map[substitutionKind]bool{
							Macros: true,
						},
					},
				}))
			})
		})

		Context("mixing subs", func() {

			It("macros,+quotes,-quotes", func() {
				b := &types.DelimitedBlock{
					Kind: types.Listing,
					Attributes: types.Attributes{
						types.AttrSubstitutions: "macros,+quotes,-quotes",
					},
				}
				Expect(newSubstitutions(b)).To(Equal(substitutions{
					{
						entrypoint: "MacrosGroup",
						rules: map[substitutionKind]bool{
							Macros: true,
						},
					},
					{
						entrypoint: "QuotesGroup",
						rules:      map[substitutionKind]bool{}, // TODO: remove substitution entry when `rules` is empty
					},
				}))
			})
		})
	})
})
