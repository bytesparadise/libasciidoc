package renderer

import (
	"fmt"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/types"
)

// ProcessDocumentHeader includes the authors and revision in the document attributes
func ProcessDocumentHeader(ctx *Context) {
	if authors, ok := ctx.Document.Authors(); ok {
		for i, author := range authors {
			var part1, part2, part3, email string
			parts := strings.Split(author.FullName, " ")
			if len(parts) > 0 {
				part1 = types.Apply(parts[0],
					func(s string) string {
						return strings.TrimSpace(s)
					},
					func(s string) string {
						return strings.Replace(s, "_", " ", -1)
					},
				)
			}
			if len(parts) > 1 {
				part2 = types.Apply(parts[1],
					func(s string) string {
						return strings.TrimSpace(s)
					},
					func(s string) string {
						return strings.Replace(s, "_", " ", -1)
					},
				)
			}
			if len(parts) > 2 {
				part3 = types.Apply(strings.Join(parts[2:], " "),
					func(s string) string {
						return strings.TrimSpace(s)
					},
					func(s string) string {
						return strings.Replace(s, "_", " ", -1)
					},
				)
			}
			if author.Email != "" {
				email = strings.TrimSpace(author.Email)
			}
			if part2 != "" && part3 != "" {
				ctx.Document.Attributes.AddNonEmpty(key("firstname", i), strings.TrimSpace(part1))
				ctx.Document.Attributes.AddNonEmpty(key("middlename", i), strings.TrimSpace(part2))
				ctx.Document.Attributes.AddNonEmpty(key("lastname", i), strings.TrimSpace(part3))
				ctx.Document.Attributes.AddNonEmpty(key("author", i), strings.Join([]string{part1, part2, part3}, " "))
				ctx.Document.Attributes.AddNonEmpty(key("authorinitials", i), strings.Join([]string{initial(part1), initial(part2), initial(part3)}, ""))
			} else if part2 != "" {
				ctx.Document.Attributes.AddNonEmpty(key("firstname", i), strings.TrimSpace(part1))
				ctx.Document.Attributes.AddNonEmpty(key("lastname", i), strings.TrimSpace(part2))
				ctx.Document.Attributes.AddNonEmpty(key("author", i), strings.Join([]string{part1, part2}, " "))
				ctx.Document.Attributes.AddNonEmpty(key("authorinitials", i), strings.Join([]string{initial(part1), initial(part2)}, ""))
			} else {
				ctx.Document.Attributes.AddNonEmpty(key("firstname", i), strings.TrimSpace(part1))
				ctx.Document.Attributes.AddNonEmpty(key("author", i), strings.TrimSpace(part1))
				ctx.Document.Attributes.AddNonEmpty(key("authorinitials", i), initial(part1))
			}
			ctx.Document.Attributes.AddNonEmpty(key("email", i), email)
		}
	}
	if revision, ok := ctx.Document.Revision(); ok {
		ctx.Document.Attributes.AddNonEmpty("revnumber", revision.Revnumber)
		ctx.Document.Attributes.AddNonEmpty("revdate", revision.Revdate)
		ctx.Document.Attributes.AddNonEmpty("revremark", revision.Revremark)
	}
}

func key(k string, i int) string {
	if i == 0 {
		return k
	}
	return fmt.Sprintf("%s_%d", k, i+1)
}

func initial(s string) string {
	if len(s) > 0 {
		return s[0:1]
	}
	return ""
}
