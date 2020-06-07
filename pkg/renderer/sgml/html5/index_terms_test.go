package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("index terms", func() {

	It("index term in existing paragraph line", func() {
		source := `a paragraph with an ((index)) term.`
		expected := `<div class="paragraph">
<p>a paragraph with an index term.</p>
</div>`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("index term in separate paragraph line", func() {
		source := `((foo_bar_baz _italic_))
a paragraph with an index term.`
		expected := `<div class="paragraph">
<p>foo_bar_baz <em>italic</em>
a paragraph with an index term.</p>
</div>`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})
})

var _ = Describe("concealed index terms", func() {

	It("concealed index term in existing paragraph line", func() {
		source := `a paragraph with an index term (((index, term, here))).`
		expected := `<div class="paragraph">
<p>a paragraph with an index term .</p>
</div>`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("concealed index term in single paragraph line", func() {
		source := `(((index, term)))
a paragraph with an index term.`
		expected := `<div class="paragraph">
<p>a paragraph with an index term.</p>
</div>`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("concealed index term in single paragraph line", func() {
		source := `(((index, term)))
a paragraph with an index term.`
		expected := `<div class="paragraph">
<p>a paragraph with an index term.</p>
</div>`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})

	It("concealed index term in labeled list term", func() {
		source := `((NNG_OPT_SUB_SUBSCRIBE))(((subscribe)))::

This option registers a topic that the subscriber is interested in.
The option is write-only, and takes an array of bytes, of arbitrary size.
Each incoming message is checked against the list of subscribed topics.
If the body begins with the entire set of bytes in the topic, then the
message is accepted.  If no topic matches, then the message is
discarded.
`
		expected := `<div class="dlist">
<dl>
<dt class="hdlist1">NNG_OPT_SUB_SUBSCRIBE</dt>
<dd>
<p>This option registers a topic that the subscriber is interested in.
The option is write-only, and takes an array of bytes, of arbitrary size.
Each incoming message is checked against the list of subscribed topics.
If the body begins with the entire set of bytes in the topic, then the
message is accepted.  If no topic matches, then the message is
discarded.</p>
</dd>
</dl>
</div>`
		Expect(RenderHTML(source)).To(MatchHTML(expected))
	})
})
