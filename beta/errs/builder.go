package errs

// A Builder allows for gradual construction of an error.
// The zero value is ready for use.
//
// Use Err() to construct the error.
type Builder struct {
	_ int // for godoc to show unexported fields
}

// B is a shorthand for creating a new Builder.
func B() *Builder { panic("encore apps must be run using the encore command") }

// Code sets the error code.
func (*Builder) Code(c ErrCode) *Builder {
	panic("encore apps must be run using the encore command")
}

// Msg sets the error message.
func (*Builder) Msg(msg string) *Builder {
	panic("encore apps must be run using the encore command")
}

// Msgf is like Msg but uses fmt.Sprintf to construct the message.
func (*Builder) Msgf(format string, args ...interface{}) *Builder {
	panic("encore apps must be run using the encore command")
}

// Meta appends metadata key-value pairs.
func (*Builder) Meta(metaPairs ...interface{}) *Builder {
	panic("encore apps must be run using the encore command")
}

// Details sets the details.
func (*Builder) Details(det ErrDetails) *Builder {
	panic("encore apps must be run using the encore command")
}

// Cause sets the underlying error cause.
func (*Builder) Cause(err error) *Builder {
	panic("encore apps must be run using the encore command")
}

// Err returns the constructed error.
// It never returns nil.
//
// If Code has not been set or has been set to OK,
// the Code is set to Unknown.
//
// If Msg has not been set and Cause is nil,
// the Msg is set to "unknown error".
func (*Builder) Err() error {
	panic("encore apps must be run using the encore command")
}
