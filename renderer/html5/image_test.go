package html5_test

import . "github.com/onsi/ginkgo"

var _ = Describe("render block images", func() {
	It("image alone", func() {

		content := "image::foo.png[]"
		expected := `<div class="imageblock">
<div class="content">
<img src="foo.png" alt="foo">
</div>
</div>`
		verify(GinkgoT(), expected, content)
	})
	It("image with alt", func() {

		content := "image::foo.png[foo image]"
		expected := `<div class="imageblock">
<div class="content">
<img src="foo.png" alt="foo image">
</div>
</div>`
		verify(GinkgoT(), expected, content)
	})
	It("image with alt and dimensions", func() {

		content := "image::foo.png[foo image, 600, 400]"
		expected := `<div class="imageblock">
<div class="content">
<img src="foo.png" alt="foo image" width="600" height="400">
</div>
</div>`
		verify(GinkgoT(), expected, content)
	})
	It("image with alt and dimensions", func() {
		content := `[#img-foobar]
.A title to foobar
[link=http://foo.bar]
image::images/foo.png[the foo.png image,600,400]`
		expected := `<div id="img-foobar" class="imageblock">
<div class="content">
<a class="image" href="http://foo.bar"><img src="images/foo.png" alt="the foo.png image" width="600" height="400"></a>
</div>
<div class="title">A title to foobar</div>
</div>`
		verify(GinkgoT(), expected, content)
	})
})
