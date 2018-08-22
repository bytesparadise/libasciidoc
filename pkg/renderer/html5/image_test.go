package html5_test

import . "github.com/onsi/ginkgo"

var _ = Describe("images", func() {

	Context("block images", func() {

		It("block image alone", func() {

			actualContent := "image::foo.png[]"
			expectedResult := `<div class="imageblock">
<div class="content">
<img src="foo.png" alt="foo">
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("block image with alt", func() {

			actualContent := "image::foo.png[foo image]"
			expectedResult := `<div class="imageblock">
<div class="content">
<img src="foo.png" alt="foo image">
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("block image with alt and dimensions", func() {

			actualContent := "image::foo.png[foo image, 600, 400]"
			expectedResult := `<div class="imageblock">
<div class="content">
<img src="foo.png" alt="foo image" width="600" height="400">
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("block image with title, alt and dimensions", func() {
			actualContent := `[#img-foobar]
.A title to foobar
[link=http://foo.bar]
image::images/foo.png[the foo.png image,600,400]`
			expectedResult := `<div id="img-foobar" class="imageblock">
<div class="content">
<a class="image" href="http://foo.bar"><img src="images/foo.png" alt="the foo.png image" width="600" height="400"></a>
</div>
<div class="title">Figure 1. A title to foobar</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("block image with role above", func() {
			actualContent := `.mytitle
[#myid]
[.myrole]
image::foo.png[foo image, 600, 400]`
			expectedResult := `<div id="myid" class="imageblock myrole">
<div class="content">
<img src="foo.png" alt="foo image" width="600" height="400">
</div>
<div class="title">Figure 1. mytitle</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("block image with id, title and role inline", func() {
			actualContent := `image::foo.png[foo image, 600, 400,id = myid, title= mytitle, role=myrole]`
			expectedResult := `<div id="myid" class="imageblock myrole">
<div class="content">
<img src="foo.png" alt="foo image" width="600" height="400">
</div>
<div class="title">Figure 1. mytitle</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

	})

	Context("inline images", func() {

		Context("valid inline Images", func() {

			It("inline image alone", func() {
				actualContent := "image:foo.png[]"
				expectedResult := `<div class="paragraph">
<p><span class="image"><img src="foo.png" alt="foo"></span></p>
</div>`
				verify(GinkgoT(), expectedResult, actualContent)
			})

			It("inline image with id, title and role", func() {
				actualContent := "image:foo.png[id=myid, title=mytitle, role=myrole]"
				expectedResult := `<div class="paragraph">
<p><span class="image myrole"><img src="foo.png" alt="foo" title="mytitle"></span></p>
</div>`
				verify(GinkgoT(), expectedResult, actualContent)
			})

			It("inline image with alt", func() {
				actualContent := "image:foo.png[foo image]"
				expectedResult := `<div class="paragraph">
<p><span class="image"><img src="foo.png" alt="foo image"></span></p>
</div>`
				verify(GinkgoT(), expectedResult, actualContent)
			})

			It("inline image with alt and dimensions", func() {
				actualContent := "image:foo.png[foo image, 600, 400]"
				expectedResult := `<div class="paragraph">
<p><span class="image"><img src="foo.png" alt="foo image" width="600" height="400"></span></p>
</div>`
				verify(GinkgoT(), expectedResult, actualContent)
			})

			It("paragraph with inline image with alt and dimensions", func() {
				actualContent := "a foo image:foo.png[foo image, 600, 400] bar"
				expectedResult := `<div class="paragraph">
<p>a foo <span class="image"><img src="foo.png" alt="foo image" width="600" height="400"></span> bar</p>
</div>`
				verify(GinkgoT(), expectedResult, actualContent)
			})
		})

		Context("invalid Inline Images", func() {

			It("paragraph with block image with alt and dimensions", func() {
				actualContent := "a foo image::foo.png[foo image, 600, 400] bar"
				expectedResult := `<div class="paragraph">
<p>a foo image::foo.png[foo image, 600, 400] bar</p>
</div>`
				verify(GinkgoT(), expectedResult, actualContent)
			})
		})
	})

	Context("imagesdir", func() {

		It("block image with relative location", func() {

			actualContent := `:imagesdir: ./assets
image::foo.png[]`
			expectedResult := `<div class="imageblock">
<div class="content">
<img src="./assets/foo.png" alt="foo">
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("2 block images with relative locations and imagedir changed in-between", func() {

			actualContent := `:imagesdir: ./assets1
image::foo.png[]

:imagesdir: ./assets2
image::bar.png[]`
			expectedResult := `<div class="imageblock">
<div class="content">
<img src="./assets1/foo.png" alt="foo">
</div>
</div>
<div class="imageblock">
<div class="content">
<img src="./assets2/bar.png" alt="bar">
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("block image with absolute URL", func() {

			actualContent := `:imagesdir: ./assets
image::https://example.com/foo.png[]`
			expectedResult := `<div class="imageblock">
<div class="content">
<img src="https://example.com/foo.png" alt="foo">
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("block image with absolute filepath", func() {

			actualContent := `:imagesdir: ./assets
image::/bar/foo.png[]`
			expectedResult := `<div class="imageblock">
<div class="content">
<img src="/bar/foo.png" alt="foo">
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})

		It("block image with absolute file scheme and path", func() {

			actualContent := `:imagesdir: ./assets
image::file:///bar/foo.png[]`
			expectedResult := `<div class="imageblock">
<div class="content">
<img src="file:///bar/foo.png" alt="foo">
</div>
</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})
})
