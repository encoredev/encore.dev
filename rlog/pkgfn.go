package rlog

// Debug logs a debug-level message.
// The variadic key-value pairs are treated as they are in With.
func Debug(msg string, keysAndValues ...any) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.13.0/runtime/rlog/pkgfn.go#L10-L12
	doPanic("encore apps must be run using the encore command")
	return
}

// Info logs an info-level message.
// The variadic key-value pairs are treated as they are in With.
func Info(msg string, keysAndValues ...any) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.13.0/runtime/rlog/pkgfn.go#L16-L18
	doPanic("encore apps must be run using the encore command")
	return
}

// Warn logs a warn-level message.
// The variadic key-value pairs are treated as they are in With.
func Warn(msg string, keysAndValues ...any) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.13.0/runtime/rlog/pkgfn.go#L22-L24
	doPanic("encore apps must be run using the encore command")
	return
}

// Error logs an error-level message.
// The variadic key-value pairs are treated as they are in With.
func Error(msg string, keysAndValues ...any) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.13.0/runtime/rlog/pkgfn.go#L28-L30
	doPanic("encore apps must be run using the encore command")
	return
}

// With adds a variadic number of fields to the logging context.
// The keysAndValues must be pairs of string keys and arbitrary data.
func With(keysAndValues ...any) (_ Ctx) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.13.0/runtime/rlog/pkgfn.go#L34-L36
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
