package parser

import (
	"github.com/bytesparadise/libasciidoc/pkg/types"
	log "github.com/sirupsen/logrus"

	. "github.com/onsi/ginkgo/extensions/table" //nolint golint
	. "github.com/onsi/gomega"                  //nolint golint
)

var _ = DescribeTable("compute valid custom substitutions",

	func(block types.BlockWithAttributes, expected substitutions) {
		// given
		log.Debugf("processing '%v'", block)
		// when
		result, err := newSubstitutions(block)
		// then
		Expect(err).NotTo(HaveOccurred())
		Expect(result).To(Equal(expected))
	},

	Entry(`listing block with default substitutions`,
		&types.DelimitedBlock{
			Kind: types.Listing,
		},
		substitutions{
			// default
			&substitution{
				entrypoint: VerbatimGroup,
				rules: map[substitutionKind]bool{
					SpecialCharacters: true,
					Callouts:          true,
				},
			},
		},
	),
	Entry(`example block with default substitutions`,
		&types.DelimitedBlock{
			Kind: types.Example,
		},
		substitutions{
			// default
			&substitution{
				entrypoint: NormalGroup,
				rules: map[substitutionKind]bool{
					InlinePassthroughs: true,
					SpecialCharacters:  true,
					Attributes:         true,
					Quotes:             true,
					Replacements:       true,
					Macros:             true,
					PostReplacements:   true,
				},
			},
		},
	),
	Entry(`listing block with custom 'attributes,quotes' substitutions`,
		&types.DelimitedBlock{
			Kind: types.Listing,
			Attributes: types.Attributes{
				types.AttrSubstitutions: `attributes,quotes`,
			},
		},
		substitutions{
			&substitution{
				entrypoint: AttributesGroup,
				rules: map[substitutionKind]bool{
					InlinePassthroughs: true,
					Attributes:         true,
				},
			},
			&substitution{
				entrypoint: QuotesGroup,
				rules: map[substitutionKind]bool{
					Quotes: true,
				},
			},
		},
	),
	// incremental substitutions
	Entry(`listing block with incremental '+attributes' substitutions `,
		&types.DelimitedBlock{
			Kind: types.Listing,
			Attributes: types.Attributes{
				types.AttrSubstitutions: `+attributes`,
			},
		},
		substitutions{
			// default
			&substitution{
				entrypoint: VerbatimGroup,
				rules: map[substitutionKind]bool{
					SpecialCharacters: true,
					Callouts:          true,
				},
			},
			// appended
			&substitution{
				entrypoint: AttributesGroup,
				rules: map[substitutionKind]bool{
					InlinePassthroughs: true,
					Attributes:         true,
				},
			},
		},
	),
	Entry(`listing block with incremental 'attributes+' substitutions `,
		&types.DelimitedBlock{
			Kind: types.Listing,
			Attributes: types.Attributes{
				types.AttrSubstitutions: `attributes+`,
			},
		},
		substitutions{
			// prepended
			&substitution{
				entrypoint: AttributesGroup,
				rules: map[substitutionKind]bool{
					InlinePassthroughs: true,
					Attributes:         true,
				},
			},
			// default
			&substitution{
				entrypoint: VerbatimGroup,
				rules: map[substitutionKind]bool{
					SpecialCharacters: true,
					Callouts:          true,
				},
			},
		},
	),
	Entry(`listing block with incremental 'attributes+,-specialchars' substitutions `,
		&types.DelimitedBlock{
			Kind: types.Listing,
			Attributes: types.Attributes{
				types.AttrSubstitutions: `attributes+,-specialchars`,
			},
		},
		substitutions{
			// prepended
			&substitution{
				entrypoint: AttributesGroup,
				rules: map[substitutionKind]bool{
					InlinePassthroughs: true,
					Attributes:         true,
				},
			},
			// default
			&substitution{
				entrypoint: VerbatimGroup,
				rules: map[substitutionKind]bool{
					// SpecialCharacters: true, // removed
					Callouts: true,
				},
			},
		},
	),
)
