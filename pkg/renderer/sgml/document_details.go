package sgml

import (
	"bytes"
	"strconv"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderDocumentDetails(ctx *renderer.Context) (*sanitized, error) {
	if ctx.Attributes.Has(types.AttrAuthors) {
		authors, err := r.renderDocumentAuthorsDetails(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "error while rendering the document details")
		}
		documentDetailsBuff := &bytes.Buffer{}
		revNumber, _ := ctx.Attributes.GetAsString("revnumber")
		revDate, _ := ctx.Attributes.GetAsString("revdate")
		revRemark, _ := ctx.Attributes.GetAsString("revremark")
		err = r.documentDetails.Execute(documentDetailsBuff, struct {
			Authors   sanitized
			RevNumber string
			RevDate   string
			RevRemark string
		}{
			Authors:   *authors,
			RevNumber: revNumber,
			RevDate:   revDate,
			RevRemark: revRemark,
		})
		if err != nil {
			return nil, errors.Wrap(err, "error while rendering the document details")
		}
		documentDetails := sanitized(documentDetailsBuff.String()) //nolint: gosec
		return &documentDetails, nil
	}
	return nil, nil
}

func (r *sgmlRenderer) renderDocumentAuthorsDetails(ctx *renderer.Context) (*sanitized, error) { // TODO: use  `types.DocumentAuthor` attribute in context
	authorsDetailsBuff := &bytes.Buffer{}
	i := 1
	for {
		var authorKey string
		var emailKey string
		var index string
		if i == 1 {
			authorKey = "author"
			emailKey = "email"
			index = ""
		} else {
			index = strconv.Itoa(i)
			authorKey = "author_" + index
			emailKey = "email_" + index
		}
		// having at least one author is the minimal requirement for document details
		if author, ok := ctx.Attributes.GetAsString(authorKey); ok {
			authorDetailsBuff := &bytes.Buffer{}
			email, _ := ctx.Attributes.GetAsString(emailKey)
			err := r.documentAuthorDetails.Execute(authorDetailsBuff, struct {
				Index string
				Name  string
				Email string
			}{
				Index: index,
				Name:  author,
				Email: email,
			})
			if err != nil {
				return nil, errors.Wrap(err, "error while rendering the document author")
			}
			// if there were authors before, need to insert a `\n`
			if i > 1 {
				authorsDetailsBuff.WriteString("\n")
			}
			authorsDetailsBuff.Write(authorDetailsBuff.Bytes())
			i++
		} else {
			break
		}
	}
	result := sanitized(authorsDetailsBuff.String()) //nolint: gosec
	return &result, nil
}
