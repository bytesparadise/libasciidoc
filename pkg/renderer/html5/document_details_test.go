package html5_test

import (
	"time"

	"github.com/bytesparadise/libasciidoc/pkg/renderer"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("document Header", func() {

	Context("header with attributes", func() {

		It("header with author and revision", func() {
			actualContent := `= The Dangerous and Thrilling Documentation Chronicles
Kismet Rainbow Chameleon <kismet@asciidoctor.org>
v1.0, June 19, 2017: First incarnation`
			// top-level section is not rendered per-say,
			// but the section will be used to set the HTML page's <title> element
			expectedResult := `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<!--[if IE]><meta http-equiv="X-UA-Compatible" content="IE=edge"><![endif]-->
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<meta name="generator" content="libasciidoc">
<title>The Dangerous and Thrilling Documentation Chronicles</title>
<body class="article">
<div id="header">
<h1>The Dangerous and Thrilling Documentation Chronicles</h1>
<div class="details">
<span id="author" class="author">Kismet Rainbow Chameleon</span><br>
<span id="email" class="email"><a href="mailto:kismet@asciidoctor.org">kismet@asciidoctor.org</a></span><br>
<span id="revnumber">version 1.0,</span>
<span id="revdate">June 19, 2017</span>
<br><span id="revremark">First incarnation</span>
</div>
</div>
<div id="content">

</div>
<div id="footer">
<div id="footer-text">
Version 1.0<br>
Last updated {{.LastUpdated}}
</div>
</div>
</body>
</html>`
			verify(GinkgoT(), expectedResult, actualContent, renderer.IncludeHeaderFooter(true), renderer.LastUpdated(time.Now()))
		})

		It("header with 2 authors and no revision", func() {
			actualContent := `= The Dangerous and Thrilling Documentation Chronicles
Kismet Rainbow Chameleon <kismet@asciidoctor.org>; Lazarus het_Draeke <lazarus@asciidoctor.org>`
			// top-level section is not rendered per-say,
			// but the section will be used to set the HTML page's <title> element
			expectedResult := `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<!--[if IE]><meta http-equiv="X-UA-Compatible" content="IE=edge"><![endif]-->
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<meta name="generator" content="libasciidoc">
<title>The Dangerous and Thrilling Documentation Chronicles</title>
<body class="article">
<div id="header">
<h1>The Dangerous and Thrilling Documentation Chronicles</h1>
<div class="details">
<span id="author" class="author">Kismet Rainbow Chameleon</span><br>
<span id="email" class="email"><a href="mailto:kismet@asciidoctor.org">kismet@asciidoctor.org</a></span><br>
<span id="author2" class="author">Lazarus het Draeke</span><br>
<span id="email2" class="email"><a href="mailto:lazarus@asciidoctor.org">lazarus@asciidoctor.org</a></span><br>
</div>
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
			verify(GinkgoT(), expectedResult, actualContent, renderer.IncludeHeaderFooter(true), renderer.LastUpdated(time.Now()))

		})
	})
})
