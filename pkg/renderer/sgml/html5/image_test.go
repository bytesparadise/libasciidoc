package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("images", func() {

	Context("block images", func() {

		It("block image alone", func() {

			source := "image::foo.png[]"
			expected := `<div class="imageblock">
<div class="content">
<img src="foo.png" alt="foo">
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("block image with alt", func() {

			source := "image::foo.png[foo image]"
			expected := `<div class="imageblock">
<div class="content">
<img src="foo.png" alt="foo image">
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("block image with alt and dimensions", func() {

			source := "image::foo.png[foo image, 600, 400]"
			expected := `<div class="imageblock">
<div class="content">
<img src="foo.png" alt="foo image" width="600" height="400">
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("block image with title, alt and dimensions", func() {
			source := `[#img-foobar]
.A title to foobar
[link=http://foo.bar]
image::images/foo.png[the foo.png image,600,400]`
			expected := `<div id="img-foobar" class="imageblock">
<div class="content">
<a class="image" href="http://foo.bar"><img src="images/foo.png" alt="the foo.png image" width="600" height="400"></a>
</div>
<div class="title">Figure 1. A title to foobar</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("block image with role above", func() {
			source := `.mytitle
[#myid]
[.myrole]
image::foo.png[foo image, 600, 400]`
			expected := `<div id="myid" class="imageblock myrole">
<div class="content">
<img src="foo.png" alt="foo image" width="600" height="400">
</div>
<div class="title">Figure 1. mytitle</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("block image with id, title and role inline", func() {
			source := `image::foo.png[foo image, 600, 400,id = myid, title= mytitle, role=myrole]`
			expected := `<div id="myid" class="imageblock myrole">
<div class="content">
<img src="foo.png" alt="foo image" width="600" height="400">
</div>
<div class="title">Figure 1. mytitle</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("2 block images", func() {
			source := `image::app.png[]
image::appa.png[]`
			expected := `<div class="imageblock">
<div class="content">
<img src="app.png" alt="app">
</div>
</div>
<div class="imageblock">
<div class="content">
<img src="appa.png" alt="appa">
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

	})

	Context("inline images", func() {

		Context("valid inline Images", func() {

			It("inline image alone", func() {
				source := "image:app.png[]"
				expected := `<div class="paragraph">
<p><span class="image"><img src="app.png" alt="app"></span></p>
</div>`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("inline image with id, title and role", func() {
				source := "image:foo.png[id=myid, title=mytitle, role=myrole]"
				expected := `<div class="paragraph">
<p><span class="image myrole"><img src="foo.png" alt="foo" title="mytitle"></span></p>
</div>`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("inline image with alt", func() {
				source := "image:foo.png[foo image]"
				expected := `<div class="paragraph">
<p><span class="image"><img src="foo.png" alt="foo image"></span></p>
</div>`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("inline image with alt and dimensions", func() {
				source := "image:foo.png[foo image, 600, 400]"
				expected := `<div class="paragraph">
<p><span class="image"><img src="foo.png" alt="foo image" width="600" height="400"></span></p>
</div>`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("paragraph with inline image with alt and dimensions", func() {
				source := "a foo image:foo.png[foo image, 600, 400] bar"
				expected := `<div class="paragraph">
<p>a foo <span class="image"><img src="foo.png" alt="foo image" width="600" height="400"></span> bar</p>
</div>`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})
		})

		Context("invalid Inline Images", func() {

			It("paragraph with block image with alt and dimensions", func() {
				source := "a foo image::foo.png[foo image, 600, 400] bar"
				expected := `<div class="paragraph">
<p>a foo image::foo.png[foo image, 600, 400] bar</p>
</div>`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})
		})
	})

	Context("imagesdir", func() {

		It("block image with relative location", func() {

			source := `:imagesdir: ./assets
image::foo.png[]`
			expected := `<div class="imageblock">
<div class="content">
<img src="./assets/foo.png" alt="foo">
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("2 block images with relative locations and imagesdir changed in-between", func() {

			source := `:imagesdir: ./assets1
image::foo.png[]

:imagesdir: ./assets2
image::bar.png[]`
			expected := `<div class="imageblock">
<div class="content">
<img src="./assets1/foo.png" alt="foo">
</div>
</div>
<div class="imageblock">
<div class="content">
<img src="./assets2/bar.png" alt="bar">
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("block image with absolute URL", func() {

			source := `:imagesdir: ./assets
image::https://example.com/foo.png[]`
			expected := `<div class="imageblock">
<div class="content">
<img src="https://example.com/foo.png" alt="foo">
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("block image with absolute filepath", func() {

			source := `:imagesdir: ./assets
image::/bar/foo.png[]`
			expected := `<div class="imageblock">
<div class="content">
<img src="/bar/foo.png" alt="foo">
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("block image with absolute file scheme and path", func() {

			source := `:imagesdir: ./assets
image::file:///bar/foo.png[]`
			expected := `<div class="imageblock">
<div class="content">
<img src="file:///bar/foo.png" alt="foo">
</div>
</div>`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})
})
