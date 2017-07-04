package html5_test

import "testing"

func TestRenderImageBlocks(t *testing.T) {
	t.Run("image alone", func(t *testing.T) {
		// given
		content := "image::foo.png[]"
		expected := `<div class="imageblock">
<div class="content">
<img src="foo.png" alt="foo">
</div>
</div>`
		verify(t, expected, content)
	})
	t.Run("image with alt", func(t *testing.T) {
		// given
		content := "image::foo.png[foo image]"
		expected := `<div class="imageblock">
<div class="content">
<img src="foo.png" alt="foo image">
</div>
</div>`
		verify(t, expected, content)
	})
	t.Run("image with alt and dimensions", func(t *testing.T) {
		// given
		content := "image::foo.png[foo image, 600, 400]"
		expected := `<div class="imageblock">
<div class="content">
<img src="foo.png" alt="foo image" width="600" height="400">
</div>
</div>`
		verify(t, expected, content)
	})
	t.Run("image with alt and dimensions", func(t *testing.T) {
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
		verify(t, expected, content)
	})
}
