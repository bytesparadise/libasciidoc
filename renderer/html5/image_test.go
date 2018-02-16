package html5_test

import . "github.com/onsi/ginkgo"

var _ = Describe("Images", func() {
	Context("Block Images", func() {

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

		It("block image with alt and dimensions", func() {
			actualContent := "[#img-foobar]\n.A title to foobar\n[link=http://foo.bar]\nimage::images/foo.png[the foo.png image,600,400]"
			expectedResult := `<div id="img-foobar" class="imageblock">
			<div class="content">
			<a class="image" href="http://foo.bar"><img src="images/foo.png" alt="the foo.png image" width="600" height="400"></a>
			</div>
			<div class="doctitle">A title to foobar</div>
			</div>`
			verify(GinkgoT(), expectedResult, actualContent)
		})
	})

	Context("Inline Images", func() {
		Context("Valid Inline Images", func() {

			It("inline image alone", func() {
				actualContent := "image:foo.png[]"
				expectedResult := `<div class="paragraph">
				<p><span class="image"><img src="foo.png" alt="foo"></span></p>
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

		Context("Invalid Inline Images", func() {

			It("paragraph with block image with alt and dimensions", func() {
				actualContent := "a foo image::foo.png[foo image, 600, 400] bar"
				expectedResult := `<div class="paragraph">
			<p>a foo image::foo.png[foo image, 600, 400] bar</p>
			</div>`
				verify(GinkgoT(), expectedResult, actualContent)
			})
		})
	})
})
