package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("images", func() {

	Context("block images", func() {

		It("alone", func() {

			source := "image::foo.png[]"
			expected := `<div class="imageblock">
<div class="content">
<img src="foo.png" alt="foo">
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with alt", func() {

			source := "image::foo.png[foo image]"
			expected := `<div class="imageblock">
<div class="content">
<img src="foo.png" alt="foo image">
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with alt and dimensions", func() {

			source := "image::foo.png[foo image, 600, 400]"
			expected := `<div class="imageblock">
<div class="content">
<img src="foo.png" alt="foo image" width="600" height="400">
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with alt and dimensions, float, align", func() {

			source := "image::foo.png[foo image, 600, 400,float=left,align=center]"
			expected := `<div class="imageblock left text-center">
<div class="content">
<img src="foo.png" alt="foo image" width="600" height="400">
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with custom caption", func() {
			// TODO: split on multiple lines for readability
			source := ".Image Title\nimage::foo.png[foo image, 600, 400,caption=\"Bar A. \"]"
			expected := `<div class="imageblock">
<div class="content">
<img src="foo.png" alt="foo image" width="600" height="400">
</div>
<div class="title">Bar A. Image Title</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with custom global figure-caption", func() {

			source := ":figure-caption: Picture\n" +
				".Image Title\nimage::foo.png[foo image, 600, 400]"
			expected := `<div class="imageblock">
<div class="content">
<img src="foo.png" alt="foo image" width="600" height="400">
</div>
<div class="title">Picture 1. Image Title</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with suppressed caption", func() {

			source := ":figure-caption!:\n" +
				".Image Title\nimage::foo.png[foo image, 600, 400]"
			expected := `<div class="imageblock">
<div class="content">
<img src="foo.png" alt="foo image" width="600" height="400">
</div>
<div class="title">Image Title</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with alt and dimensions and multiple roles", func() {

			source := `[.role1.role2]
image::foo.png[foo image, 600, 400]`
			expected := `<div class="imageblock role1 role2">
<div class="content">
<img src="foo.png" alt="foo image" width="600" height="400">
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with title, alt and dimensions", func() {
			source := `[#img-foobar]
.A title to foobar
[link=http://foo.bar]
image::images/foo.png[the foo.png image,600,400]`
			expected := `<div id="img-foobar" class="imageblock">
<div class="content">
<a class="image" href="http://foo.bar"><img src="images/foo.png" alt="the foo.png image" width="600" height="400"></a>
</div>
<div class="title">Figure 1. A title to foobar</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with role above", func() {
			source := `.mytitle
[#myid]
[.myrole]
image::foo.png[foo image, 600, 400]`
			expected := `<div id="myid" class="imageblock myrole">
<div class="content">
<img src="foo.png" alt="foo image" width="600" height="400">
</div>
<div class="title">Figure 1. mytitle</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with id, title and role inline", func() {
			source := `image::foo.png[foo image, 600, 400,id = myid, title= mytitle, role=myrole]`
			expected := `<div id="myid" class="imageblock myrole">
<div class="content">
<img src="foo.png" alt="foo image" width="600" height="400">
</div>
<div class="title">Figure 1. mytitle</div>
</div>
`
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
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with special characters", func() {
			source := `image::http://example.com/foo.png?a=1&b=2[]`
			expected := `<div class="imageblock">
<div class="content">
<img src="http://example.com/foo.png?a=1&b=2" alt="foo">
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

	})

	Context("inline images", func() {

		Context("valid inline images", func() {

			It("alone", func() {
				source := "image:app.png[]"
				expected := `<div class="paragraph">
<p><span class="image"><img src="app.png" alt="app"></span></p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("with id, title and role", func() {
				source := "image:foo.png[id=myid, title=mytitle, role=myrole]"
				expected := `<div class="paragraph">
<p><span class="image myrole"><img src="foo.png" alt="foo" title="mytitle"></span></p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("with alt", func() {
				source := "image:foo.png[foo image]"
				expected := `<div class="paragraph">
<p><span class="image"><img src="foo.png" alt="foo image"></span></p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("with alt and dimensions", func() {
				source := "image:foo.png[foo image, 600, 400]"
				expected := `<div class="paragraph">
<p><span class="image"><img src="foo.png" alt="foo image" width="600" height="400"></span></p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("with link", func() {
				source := "image:foo.png[foo image, link=http://foo.bar]"
				expected := `<div class="paragraph">
<p><span class="image"><a class="image" href="http://foo.bar"><img src="foo.png" alt="foo image"></a></span></p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("paragraph with inline image with alt and dimensions", func() {
				source := "a foo image:foo.png[foo image, 600, 400] bar"
				expected := `<div class="paragraph">
<p>a foo <span class="image"><img src="foo.png" alt="foo image" width="600" height="400"></span> bar</p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("with special characters", func() {
				source := `image:http://example.com/foo.png?a=1&b=2[]`
				expected := `<div class="paragraph">
<p><span class="image"><img src="http://example.com/foo.png?a=1&b=2" alt="foo"></span></p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})
		})

		Context("invalid inlines images", func() {

			It("paragraph with block image with alt and dimensions", func() {
				source := "a foo image::foo.png[foo image, 600, 400] bar"
				expected := `<div class="paragraph">
<p>a foo image::foo.png[foo image, 600, 400] bar</p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})
		})
	})

	Context("imagesdir", func() {

		It("with relative location", func() {

			source := `:imagesdir: ./assets
image::foo.png[]`
			expected := `<div class="imageblock">
<div class="content">
<img src="./assets/foo.png" alt="foo">
</div>
</div>
`
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
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with absolute URL", func() {

			source := `:imagesdir: ./assets
image::https://example.com/foo.png[]`
			expected := `<div class="imageblock">
<div class="content">
<img src="https://example.com/foo.png" alt="foo">
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with absolute filepath", func() {

			source := `:imagesdir: ./assets
image::/bar/foo.png[]`
			expected := `<div class="imageblock">
<div class="content">
<img src="/bar/foo.png" alt="foo">
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})

		It("with absolute file scheme and path", func() {

			source := `:imagesdir: ./assets
image::file:///bar/foo.png[]`
			expected := `<div class="imageblock">
<div class="content">
<img src="file:///bar/foo.png" alt="foo">
</div>
</div>
`
			Expect(RenderHTML(source)).To(MatchHTML(expected))
		})
	})
})
