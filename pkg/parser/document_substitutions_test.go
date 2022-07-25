package parser

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("new substitutions",
	func(actual string, expected *substitutions) {
		b := &types.DelimitedBlock{
			Attributes: types.Attributes{
				types.AttrSubstitutions: actual,
			},
		}
		Expect(newSubstitutions(b)).To(Equal(expected))
	},
	// default/built-in subs groups
	Entry("normal", "normal",
		&substitutions{
			sequence: []string{
				InlinePassthroughs,
				AttributeRefs,
				SpecialCharacters,
				Quotes,
				Replacements,
				Macros,
				PostReplacements,
			},
		}),
	Entry("none", "none",
		&substitutions{
			sequence: []string{},
		}),
	Entry("verbatim", "verbatim",
		&substitutions{
			sequence: []string{
				Callouts,
				SpecialCharacters,
			},
		}),
	Entry("attributes", "attributes",
		&substitutions{
			sequence: []string{
				AttributeRefs,
			},
		}),
	Entry("macros", "macros",
		&substitutions{
			sequence: []string{
				Macros,
			},
		}),
	Entry("quotes", "quotes",
		&substitutions{
			sequence: []string{
				Quotes,
			},
		}),
	Entry("replacements", "replacements",
		&substitutions{
			sequence: []string{
				Replacements,
			},
		}),
	Entry("post_replacements", "post_replacements",
		&substitutions{
			sequence: []string{
				PostReplacements,
			},
		}),
	Entry("callouts", "callouts",
		&substitutions{
			sequence: []string{
				Callouts,
			},
		}),
	Entry("specialchars", "specialchars",
		&substitutions{
			sequence: []string{
				SpecialCharacters,
			},
		}),

	// custom subs
	Entry("verbatim,+attributes", "verbatim,+attributes", // append
		&substitutions{
			sequence: []string{
				Callouts,
				SpecialCharacters,
				AttributeRefs,
			},
		}),
	Entry("verbatim,attributes+", "verbatim,attributes+", // prepend
		&substitutions{
			sequence: []string{
				AttributeRefs,
				Callouts,
				SpecialCharacters,
			},
		}),
	Entry("verbatim,-callouts", "verbatim,-callouts", // remove
		&substitutions{
			sequence: []string{
				SpecialCharacters,
			},
		}),
	Entry("verbatim,-callouts,attributes+,+replacements", "verbatim,-callouts,attributes+,+replacements", // remove
		&substitutions{
			sequence: []string{
				AttributeRefs,
				SpecialCharacters,
				Replacements,
			},
		}),
)

var _ = DescribeTable("split substitutions",
	func(actual string, expectedPhase1, expectedPhase2 *substitutions) {
		b := &types.DelimitedBlock{
			Attributes: types.Attributes{
				types.AttrSubstitutions: actual,
			},
		}
		s, err := newSubstitutions(b)
		Expect(err).NotTo(HaveOccurred())
		phase1, phase2 := s.split()
		Expect(phase1).To(Equal(expectedPhase1))
		Expect(phase2).To(Equal(expectedPhase2))
	},
	// default/built-in subs groups
	Entry("normal", "normal",
		&substitutions{
			sequence: []string{
				InlinePassthroughs,
				AttributeRefs,
				SpecialCharacters,
				Quotes,
				Replacements,
				Macros,
				PostReplacements,
			},
		},
		&substitutions{
			sequence: []string{
				SpecialCharacters,
				Quotes,
				Replacements,
				Macros,
				PostReplacements,
			},
		}),
	Entry("attributes,macros", "attributes,macros",
		&substitutions{
			sequence: []string{
				AttributeRefs,
				Macros,
			},
		},
		&substitutions{
			sequence: []string{
				Macros,
			},
		}),
	Entry("macros,attributes", "macros,attributes",
		&substitutions{
			sequence: []string{
				Macros,
				AttributeRefs,
			},
		},
		nil,
	),
)
