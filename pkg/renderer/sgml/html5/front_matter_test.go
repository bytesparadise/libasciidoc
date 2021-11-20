package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("documents with front-matter", func() {

	It("should render with front-matter", func() {
		source := `---
description: User Manual
---

{description}
		`
		expected := `<div class="paragraph">
<p>User Manual</p>
</div>
`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

})
