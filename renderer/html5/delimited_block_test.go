package html5_test

import "testing"

func TestRenderDelimitedSourceBlock(t *testing.T) {
	t.Run("source with multiple lines", func(t *testing.T) {
		// given
		content := "```\nsome source code\n\nhere\n```"
		expected := `<div class="listingblock">
<div class="content">
<pre class="highlight"><code>some source code

here</code></pre>
</div>
</div>`
		verify(t, expected, content)
	})
}
