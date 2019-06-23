package html5_test

import (
	"time"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("document header", func() {

	Context("header with inline elements in title", func() {

		It("header with quoted text", func() {
			source := `= The _Dangerous_ and *Thrilling* Documentation Chronicles`
			expected := `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<!--[if IE]><meta http-equiv="X-UA-Compatible" content="IE=edge"><![endif]-->
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<meta name="generator" content="libasciidoc">
<title>The Dangerous and Thrilling Documentation Chronicles</title>
</head>
<body class="article">
<div id="header">
<h1>The <em>Dangerous</em> and <strong>Thrilling</strong> Documentation Chronicles</h1>
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
			verify(expected, source, renderer.IncludeHeaderFooter(true), renderer.LastUpdated(time.Now()))
		})
	})

})
