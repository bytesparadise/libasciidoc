package html5_test

import (
	"reflect"

	"github.com/bytesparadise/libasciidoc/pkg/renderer/sgml/html5"
	. "github.com/onsi/ginkgo" //nolint golint
	. "github.com/onsi/gomega" //nolint golint
)

var _ = Describe("fields", func() {

	It("template fields are not empty", func() {
		tmp := html5.Templates() // sgml.Templates
		typ := reflect.TypeOf(tmp)
		val := reflect.ValueOf(tmp)

		for i := 0; i < typ.NumField(); i++ {
			fn := typ.Field(i).Name
			fv := val.FieldByName(fn)

			s, ok := fv.Interface().(string)
			Expect(ok).To(BeTrue())
			Expect(s).NotTo(BeEmpty())
		}
	})
})
