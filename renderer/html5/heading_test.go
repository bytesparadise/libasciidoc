package html5_test

import "testing"

func TestRenderHeading(t *testing.T) {
	t.Run("heading level 1", func(t *testing.T) {
		// given
		content := "= a title"
		expected := `<div id="header">
<h1>a title</h1>
</div>`
		verify(t, expected, content)
	})
	t.Run("heading with just bold content", func(t *testing.T) {
		// given
		content := `==  *2 spaces and bold content*`
		expected := `<h2 id="__strong_2_spaces_and_bold_content_strong"><strong>2 spaces and bold content</strong></h2>`
		verify(t, expected, content)
	})
	t.Run("heading with nested bold content", func(t *testing.T) {
		// given
		content := `== a section title, with *bold content*`
		expected := `<h2 id="_a_section_title_with_strong_bold_content_strong">a section title, with <strong>bold content</strong></h2>`
		verify(t, expected, content)
	})
	t.Run("heading with custom ID", func(t *testing.T) {
		// given
		content := `[#custom_id]
== a section title, with *bold content*`
		expected := `<h2 id="custom_id">a section title, with <strong>bold content</strong></h2>`
		verify(t, expected, content)
	})
}
