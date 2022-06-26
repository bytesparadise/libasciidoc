package sgml

import (
	"bytes"
	"strconv"
	"strings"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	"github.com/bytesparadise/libasciidoc/pkg/types"
	"github.com/pkg/errors"
)

func (r *sgmlRenderer) renderDocumentDetails(ctx *renderer.Context) (string, error) {
	if !ctx.Attributes.Has(types.AttrAuthors) {
		return "", nil
	}
	authors, err := r.renderDocumentAuthorsDetails(ctx)
	if err != nil {
		return "", errors.Wrap(err, "error while rendering the document details")
	}
	documentDetailsBuff := &bytes.Buffer{}
	revLabel, _, err := ctx.Attributes.GetAsString(types.AttrVersionLabel)
	if err != nil {
		return "", errors.Wrap(err, "error while rendering the document details")
	}
	revNumber, _, err := ctx.Attributes.GetAsString("revnumber")
	if err != nil {
		return "", errors.Wrap(err, "error while rendering the document details")
	}
	revDate, _, err := ctx.Attributes.GetAsString("revdate")
	if err != nil {
		return "", errors.Wrap(err, "error while rendering the document details")
	}
	revRemark, _, err := ctx.Attributes.GetAsString("revremark")
	if err != nil {
		return "", errors.Wrap(err, "error while rendering the document details")
	}
	tmpl, err := r.documentDetails()
	if err != nil {
		return "", errors.Wrap(err, "unable to load document details template")
	}
	if err = tmpl.Execute(documentDetailsBuff, struct {
		Authors   string
		RevLabel  string
		RevNumber string
		RevDate   string
		RevRemark string
	}{
		Authors:   authors,
		RevLabel:  revLabel,
		RevNumber: revNumber,
		RevDate:   revDate,
		RevRemark: revRemark,
	}); err != nil {
		return "", errors.Wrap(err, "error while rendering the document details")
	}
	return documentDetailsBuff.String(), nil
}

func (r *sgmlRenderer) renderDocumentAuthorsDetails(ctx *renderer.Context) (string, error) { // TODO: use  `types.DocumentAuthor` attribute in context
	authorsDetailsBuff := &strings.Builder{}
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
		if author, ok, err := ctx.Attributes.GetAsString(authorKey); err != nil {
			return "", errors.Wrap(err, "error while rendering the document authors")
		} else if ok {
			if i > 1 {
				authorsDetailsBuff.WriteString("\n")
			}
			email, _, err := ctx.Attributes.GetAsString(emailKey)
			if err != nil {
				return "", errors.Wrap(err, "error while rendering the document authors")
			}
			tmpl, err := r.documentAuthorDetails()
			if err != nil {
				return "", errors.Wrap(err, "unable to load document authors template")
			}
			if err := tmpl.Execute(authorsDetailsBuff, struct {
				Index string
				Name  string
				Email string
			}{
				Index: index,
				Name:  author,
				Email: email,
			}); err != nil {
				return "", errors.Wrap(err, "error while rendering the document authors")
			}
			// if there were authors before, need to insert a `\n`
			i++
		} else {
			break
		}
	}
	return authorsDetailsBuff.String(), nil
}
