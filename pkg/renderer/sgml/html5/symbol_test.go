package html5_test

import (
	. "github.com/bytesparadise/libasciidoc/testsupport"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("symbols", func() {

	Context("in final documents", func() {

		Context("m-dashes", func() {

			It("should detect between word characters", func() {
				source := "some text--idea apart--continues here"
				expected := `<div class="paragraph">
<p>some text&#8212;&#8203;idea apart&#8212;&#8203;continues here</p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("should detect between word character and line boundary", func() {
				source := "some text--idea apart--"
				expected := `<div class="paragraph">
<p>some text&#8212;&#8203;idea apart&#8212;&#8203;</p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("should detect between spaces", func() {
				source := "some text -- idea apart -- continues here"
				expected := `<div class="paragraph">
<p>some text&#8201;&#8212;&#8201;idea apart&#8201;&#8212;&#8201;continues here</p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("should detect between space and line boudary", func() {
				source := "some text -- idea apart --"
				expected := `<div class="paragraph">
<p>some text&#8201;&#8212;&#8201;idea apart&#8201;&#8212;&#8201;</p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			Context("invalid", func() {

				It("should not detect when missing spaces", func() {
					source := "some text --idea apart-- continues here" // `--idea` and `apart--` are missing spaces between characters and dashes
					expected := `<div class="paragraph">
<p>some text --idea apart-- continues here</p>
</div>
`
					Expect(RenderHTML(source)).To(MatchHTML(expected))
				})
			})
		})

		Context("single right arrows", func() {

			It("should detect between spaces", func() {
				source := "go -> here"
				expected := `<div class="paragraph">
<p>go &#8594; here</p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("should detect between character and space", func() {
				source := "go-> here"
				expected := `<div class="paragraph">
<p>go&#8594; here</p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("should detect between space and character", func() {
				source := "go ->here"
				expected := `<div class="paragraph">
<p>go &#8594;here</p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("should detect between characters", func() {
				source := "go->here"
				expected := `<div class="paragraph">
<p>go&#8594;here</p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})
		})

		Context("single left arrows", func() {

			It("should detect between spaces", func() {
				source := "go <- here"
				expected := `<div class="paragraph">
<p>go &#8592; here</p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("should detect between character and space", func() {
				source := "go<- here"
				expected := `<div class="paragraph">
<p>go&#8592; here</p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("should detect between space and character", func() {
				source := "go <-here"
				expected := `<div class="paragraph">
<p>go &#8592;here</p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("should detect between characters", func() {
				source := "go<-here"
				expected := `<div class="paragraph">
<p>go&#8592;here</p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})
		})

		Context("double right arrows", func() {

			It("should detect between spaces", func() {
				source := "go => here"
				expected := `<div class="paragraph">
<p>go &#8658; here</p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("should detect between character and space", func() {
				source := "go=> here"
				expected := `<div class="paragraph">
<p>go&#8658; here</p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("should detect between space and character", func() {
				source := "go =>here"
				expected := `<div class="paragraph">
<p>go &#8658;here</p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("should detect between characters", func() {
				source := "go=>here"
				expected := `<div class="paragraph">
<p>go&#8658;here</p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})
		})

		Context("double left arrows", func() {

			It("should detect between spaces", func() {
				source := "go <= here"
				expected := `<div class="paragraph">
<p>go &#8656; here</p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("should detect between character and space", func() {
				source := "go<= here"
				expected := `<div class="paragraph">
<p>go&#8656; here</p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("should detect between space and character", func() {
				source := "go <=here"
				expected := `<div class="paragraph">
<p>go &#8656;here</p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})

			It("should detect between characters", func() {
				source := "go<=here"
				expected := `<div class="paragraph">
<p>go&#8656;here</p>
</div>
`
				Expect(RenderHTML(source)).To(MatchHTML(expected))
			})
		})
	})

})
