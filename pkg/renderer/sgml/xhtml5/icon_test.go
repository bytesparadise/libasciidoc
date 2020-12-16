package xhtml5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("icons", func() {

	Context("inline icons", func() {

		Context("icon text", func() {

			It("inline icon alone", func() {
				source := "icon:caution[]"
				expected := `<div class="paragraph">
<p><span class="icon">[caution]</span></p>
</div>
`
				Expect(RenderXHTML(source)).To(MatchHTML(expected))
			})

			It("inline icon title", func() {
				source := `icon:caution[title="title"]`
				expected := `<div class="paragraph">
<p><span class="icon">[caution]</span></p>
</div>
`
				Expect(RenderXHTML(source)).To(MatchHTML(expected))
			})

			It("inline icon alt", func() {
				source := `icon:caution[alt="Alternate"]`
				expected := `<div class="paragraph">
<p><span class="icon">[Alternate]</span></p>
</div>
`
				Expect(RenderXHTML(source)).To(MatchHTML(expected))
			})

			It("inline icon link", func() {
				source := `
icon:amazon[link="https://amazon.com"]`
				expected := `<div class="paragraph">
<p><span class="icon"><a class="image" href="https://amazon.com">[amazon]</a></span></p>
</div>
`
				Expect(RenderXHTML(source)).To(MatchHTML(expected))
			})

			It("inline icon link with target", func() {
				source := `
icon:amazon[link="https://amazon.com",window="new front"]`
				expected := `<div class="paragraph">
<p><span class="icon"><a class="image" href="https://amazon.com" target="new front">[amazon]</a></span></p>
</div>
`
				Expect(RenderXHTML(source)).To(MatchHTML(expected))
			})

		})

		Context("icon fonts", func() {

			It("inline icon alone", func() {
				source := `:icons: font

icon:caution[]`
				expected := `<div class="paragraph">
<p><span class="icon"><i class="fa fa-caution"></i></span></p>
</div>
`
				Expect(RenderXHTML(source)).To(MatchHTML(expected))
			})

			It("inline icon title", func() {
				source := `:icons: font

icon:caution[title="title"]`
				expected := `<div class="paragraph">
<p><span class="icon"><i class="fa fa-caution" title="title"></i></span></p>
</div>
`
				Expect(RenderXHTML(source)).To(MatchHTML(expected))
			})

			It("inline icon alt", func() {
				source := `:icons: font

icon:caution[alt="Alternate"]`
				expected := `<div class="paragraph">
<p><span class="icon"><i class="fa fa-caution"></i></span></p>
</div>
`
				Expect(RenderXHTML(source)).To(MatchHTML(expected))
			})

			It("inline icon size", func() {
				source := `:icons: font

icon:tip[fw]`
				expected := `<div class="paragraph">
<p><span class="icon"><i class="fa fa-tip fa-fw"></i></span></p>
</div>
`
				Expect(RenderXHTML(source)).To(MatchHTML(expected))
			})
			It("inline icon attributes", func() {
				source := `:icons: font

icon:tip[fw,rotate=90,flip=horizontal]`
				expected := `<div class="paragraph">
<p><span class="icon"><i class="fa fa-tip fa-fw fa-rotate-90 fa-flip-horizontal"></i></span></p>
</div>
`
				Expect(RenderXHTML(source)).To(MatchHTML(expected))
			})
			It("inline icon width, height", func() {
				source := `:icons: font

icon:warning[width=20px,height=30px]`
				expected := `<div class="paragraph">
<p><span class="icon"><i class="fa fa-warning"></i></span></p>
</div>
`
				Expect(RenderXHTML(source)).To(MatchHTML(expected))
			})

			It("inline icon link", func() {
				source := `:icons: font

icon:amazon[link="https://amazon.com"]`
				expected := `<div class="paragraph">
<p><span class="icon"><a class="image" href="https://amazon.com"><i class="fa fa-amazon"></i></a></span></p>
</div>
`
				Expect(RenderXHTML(source)).To(MatchHTML(expected))
			})

			It("inline icon link with target", func() {
				source := `:icons: font

icon:amazon[link="https://amazon.com",window="new front"]`
				expected := `<div class="paragraph">
<p><span class="icon"><a class="image" href="https://amazon.com" target="new front"><i class="fa fa-amazon"></i></a></span></p>
</div>
`
				Expect(RenderXHTML(source)).To(MatchHTML(expected))
			})

		})

		Context("icon images", func() {

			It("inline icon alone", func() {
				source := `:icons:

icon:caution[]`
				expected := `<div class="paragraph">
<p><span class="icon"><img src="images/icons/caution.png" alt="caution"/></span></p>
</div>
`
				Expect(RenderXHTML(source)).To(MatchHTML(expected))
			})

			It("inline icon title", func() {
				source := `:icons: image

icon:caution[title="title"]`
				expected := `<div class="paragraph">
<p><span class="icon"><img src="images/icons/caution.png" alt="caution" title="title"/></span></p>
</div>
`
				Expect(RenderXHTML(source)).To(MatchHTML(expected))
			})

			It("inline icon size", func() {
				source := `:icons:

icon:tip[fw]`
				expected := `<div class="paragraph">
<p><span class="icon"><img src="images/icons/tip.png" alt="tip" class="fa-fw"/></span></p>
</div>
`
				Expect(RenderXHTML(source)).To(MatchHTML(expected))
			})
			It("inline icon attributes", func() {
				source := `:icons:

icon:tip[fw,rotate=90,flip=horizontal]`
				expected := `<div class="paragraph">
<p><span class="icon"><img src="images/icons/tip.png" alt="tip" class="fa-fw fa-rotate-90 fa-flip-horizontal"/></span></p>
</div>
`
				Expect(RenderXHTML(source)).To(MatchHTML(expected))
			})

			It("inline icon link", func() {
				source := `:icons:

icon:amazon[link="https://amazon.com"]`
				expected := `<div class="paragraph">
<p><span class="icon"><a class="image" href="https://amazon.com"><img src="images/icons/amazon.png" alt="amazon"/></a></span></p>
</div>
`
				Expect(RenderXHTML(source)).To(MatchHTML(expected))
			})

			It("inline icon link with target", func() {
				source := `:icons:

icon:amazon[link="https://amazon.com",window="new front"]`
				expected := `<div class="paragraph">
<p><span class="icon"><a class="image" href="https://amazon.com" target="new front"><img src="images/icons/amazon.png" alt="amazon"/></a></span></p>
</div>
`
				Expect(RenderXHTML(source)).To(MatchHTML(expected))
			})

			It("inline icon directory", func() {
				source := `:icons:
:iconsdir: assets/ico

icon:caution[]`
				expected := `<div class="paragraph">
<p><span class="icon"><img src="assets/ico/caution.png" alt="caution"/></span></p>
</div>
`
				Expect(RenderXHTML(source)).To(MatchHTML(expected))
			})

			It("svg icons", func() {
				source := `:icons:
:icontype: svg

icon:caution[]`
				expected := `<div class="paragraph">
<p><span class="icon"><img src="images/icons/caution.svg" alt="caution"/></span></p>
</div>
`
				Expect(RenderXHTML(source)).To(MatchHTML(expected))
			})
		})

		Context("icon contexts", func() {

			It("icon in a title", func() {
				source := `:icons:
:icontype: svg

== icon:caution[alt="!"] Choke Hazard`
				expected := `<div class="sect1">
<h2 id="_choke_hazard"><span class="icon"><img src="images/icons/caution.svg" alt="!"/></span> Choke Hazard</h2>
<div class="sectionbody">
</div>
</div>
`
				Expect(RenderXHTML(source)).To(MatchHTML(expected))
			})

			It("icon in a definition", func() {
				source := `:icons:

what:: icon:question[]`
				expected := `<div class="dlist">
<dl>
<dt class="hdlist1">what</dt>
<dd>
<p><span class="icon"><img src="images/icons/question.png" alt="question"/></span></p>
</dd>
</dl>
</div>
`
				Expect(RenderXHTML(source)).To(MatchHTML(expected))
			})

			It("icon in a term", func() {
				source := `:icons:

icon:tip[]:: tip of the day`
				expected := `<div class="dlist">
<dl>
<dt class="hdlist1"><span class="icon"><img src="images/icons/tip.png" alt="tip"/></span></dt>
<dd>
<p>tip of the day</p>
</dd>
</dl>
</div>
`
				Expect(RenderXHTML(source)).To(MatchHTML(expected))
			})

			It("icon in quoted text", func() {
				source := `:icons: font

here [.strikeout]##we go icon:stop[]##`
				expected := `<div class="paragraph">
<p>here <span class="strikeout">we go <span class="icon"><i class="fa fa-stop"></i></span></span></p>
</div>
`
				Expect(RenderXHTML(source)).To(MatchHTML(expected))
			})

			It("icon in italic text", func() {
				source := `:icons: font

here _we go icon:stop[]_`
				expected := `<div class="paragraph">
<p>here <em>we go <span class="icon"><i class="fa fa-stop"></i></span></em></p>
</div>
`
				Expect(RenderXHTML(source)).To(MatchHTML(expected))
			})
			It("icon in bold text", func() {
				source := `:icons: font

here *we go icon:stop[]*`
				expected := `<div class="paragraph">
<p>here <strong>we go <span class="icon"><i class="fa fa-stop"></i></span></strong></p>
</div>
`
				Expect(RenderXHTML(source)).To(MatchHTML(expected))
			})

			It("icon in monospace text", func() {
				source := "here `we go icon:stop[]`"
				expected := `<div class="paragraph">
<p>here <code>we go <span class="icon">[stop]</span></code></p>
</div>
`
				Expect(RenderXHTML(source)).To(MatchHTML(expected))
			})

		})
	})
})
