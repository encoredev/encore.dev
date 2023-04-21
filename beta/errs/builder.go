package errs

// A Builder allows for gradual construction of an error.
// The zero value is ready for use.
//
// Use Err() to construct the error.
type Builder struct {
	_ int // for godoc to show unexported fields
}

// B is a shorthand for creating a new Builder.
func B() (_ *Builder) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//
	//	https://github.com/encoredev/encore/blob/v1.16.1/runtime/beta/errs/builder.go#L25-L25
	doPanic("encore apps must be run using the encore command")
	return
}

// Code sets the error code.
func (*Builder) Code(c ErrCode) (_ *Builder) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.16.1/runtime/beta/errs/builder.go#L28-L32
	doPanic("encore apps must be run using the encore command")
	return
}

// Msg sets the error message.
func (*Builder) Msg(msg string) (_ *Builder) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.16.1/runtime/beta/errs/builder.go#L35-L38
	doPanic("encore apps must be run using the encore command")
	return
}

// Msgf is like Msg but uses fmt.Sprintf to construct the message.
func (*Builder) Msgf(format string, args ...interface{}) (_ *Builder) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.16.1/runtime/beta/errs/builder.go#L41-L44
	doPanic("encore apps must be run using the encore command")
	return
}

// Meta appends metadata key-value pairs.
func (*Builder) Meta(metaPairs ...interface{}) (_ *Builder) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.16.1/runtime/beta/errs/builder.go#L47-L50
	doPanic("encore apps must be run using the encore command")
	return
}

// Details sets the details.
func (*Builder) Details(det ErrDetails) (_ *Builder) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.16.1/runtime/beta/errs/builder.go#L53-L57
	doPanic("encore apps must be run using the encore command")
	return
}

// Cause sets the underlying error cause.
func (*Builder) Cause(err error) (_ *Builder) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.16.1/runtime/beta/errs/builder.go#L60-L71
	doPanic("encore apps must be run using the encore command")
	return
}

// Err returns the constructed error.
// It never returns nil.
//
// If Code has not been set or has been set to OK,
// the Code is set to Unknown.
//
// If Msg has not been set and Cause is nil,
// the Msg is set to "unknown error".
func (*Builder) Err() (_ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.16.1/runtime/beta/errs/builder.go#L81-L109
	doPanic("encore apps must be run using the encore command")
	return
}

// doPanic is a wrapper around panic to prevent static analysis tools
// from thinking Encore APIs unconditionally panic.,
func doPanic(v any) {
	if true {
		panic(v)
	}
}
