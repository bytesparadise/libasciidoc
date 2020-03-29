package html5_test

import (
	"time"

	"github.com/bytesparadise/libasciidoc/pkg/configuration"
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("document header", func() {

	Context("header with inline elements in title", func() {

		It("header with quoted text", func() {
			source := `= The _Document_ *Title*`
			expected := `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<!--[if IE]><meta http-equiv="X-UA-Compatible" content="IE=edge"><![endif]-->
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<meta name="generator" content="libasciidoc">
<link type="text/css" rel="stylesheet" href="/path/to/style.css">
<title>The Document Title</title>
</head>
<body class="article">
<div id="header">
<h1>The <em>Document</em> <strong>Title</strong></h1>
</div>
<div id="content">

</div>
<div id="footer">
<div id="footer-text">
Last updated {{.LastUpdated}}
</div>
</div>
</body>
</html>`
			now := time.Now()
			Expect(RenderHTML5Body(source, configuration.WithHeaderFooter(true),
				configuration.WithCSS("/path/to/style.css"),
				configuration.WithLastUpdated(now))).
				To(MatchHTML5Template(expected, now))
		})
	})

})
