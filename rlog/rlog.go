// Package rlog provides a simple logging interface which is integrated with Encore's
// inbuilt distributed tracing.
//
// For more information about logging inside Encore applications see https://encore.dev/docs/observability/logging.
package rlog

// Ctx holds additional logging context for use with the Infoc and family
// of logging functions.
type Ctx struct {
	_ int // for godoc to show unexported fields
}

// Debug logs a debug-level message, merging the context from ctx
// with the additional context provided as key-value pairs.
// The variadic key-value pairs are treated as they are in With.
func (Ctx) Debug(msg string, keysAndValues ...any) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.46.1/runtimes/go/rlog/rlog.go#L79-L83
	doPanic("encore apps must be run using the encore command")
	return
}

// Info logs an info-level message, merging the context from ctx
// with the additional context provided as key-value pairs.
// The variadic key-value pairs are treated as they are in With.
func (Ctx) Info(msg string, keysAndValues ...any) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.46.1/runtimes/go/rlog/rlog.go#L88-L92
	doPanic("encore apps must be run using the encore command")
	return
}

// Warn logs a warn-level message, merging the context from ctx
// with the additional context provided as key-value pairs.
// The variadic key-value pairs are treated as they are in With.
func (Ctx) Warn(msg string, keysAndValues ...any) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.46.1/runtimes/go/rlog/rlog.go#L97-L101
	doPanic("encore apps must be run using the encore command")
	return
}

// Error logs an error-level message, merging the context from ctx
// with the additional context provided as key-value pairs.
// The variadic key-value pairs are treated as they are in With.
func (Ctx) Error(msg string, keysAndValues ...any) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.46.1/runtimes/go/rlog/rlog.go#L106-L110
	doPanic("encore apps must be run using the encore command")
	return
}

// With creates a new logging context that inherits the context
// from the original ctx and adds additional context on top.
// The original ctx is not affected.
func (Ctx) With(keysAndValues ...any) (_ Ctx) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.46.1/runtimes/go/rlog/rlog.go#L115-L129
	doPanic("encore apps must be run using the encore command")
	return
}
