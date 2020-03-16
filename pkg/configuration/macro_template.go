package configuration

import "io"

// MacroTemplate an interface of template for user macro.
type MacroTemplate interface {
	Execute(wr io.Writer, data interface{}) error
}
