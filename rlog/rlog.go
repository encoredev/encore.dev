// Package rlog provides a simple logging interface which is integrated with Encore's
// inbuilt distributed tracing.
//
// For more information about logging inside Encore applications see https://encore.dev/docs/observability/logging.
package rlog

// Debug logs a debug-level message.
// The variadic key-value pairs are treated as they are in With.
func Debug(msg string, keysAndValues ...interface{}) {
	panic("encore apps must be run using the encore command")
}

// Info logs an info-level message.
// The variadic key-value pairs are treated as they are in With.
func Info(msg string, keysAndValues ...interface{}) {
	panic("encore apps must be run using the encore command")
}

// Error logs an error-level message.
// The variadic key-value pairs are treated as they are in With.
func Error(msg string, keysAndValues ...interface{}) {
	panic("encore apps must be run using the encore command")
}

// Ctx holds additional logging context for use with the Infoc and family
// of logging functions.
type Ctx struct{}

// With adds a variadic number of fields to the logging context.
// The keysAndValues must be pairs of string keys and arbitrary data.
func With(keysAndValues ...interface{}) Ctx {
	panic("encore apps must be run using the encore command")
}

// With creates a new logging context that inherits the context
// from the original ctx and adds additional context on top.
// The original ctx is not affected.
func (ctx Ctx) With(keysAndValues ...interface{}) Ctx {
	panic("encore apps must be run using the encore command")
}

// Debug logs a debug-level message, merging the context from ctx
// with the additional context provided as key-value pairs.
// The variadic key-value pairs are treated as they are in With.
func (ctx Ctx) Debug(msg string, keysAndValues ...interface{}) {
	panic("encore apps must be run using the encore command")
}

// Info logs an info-level message, merging the context from ctx
// with the additional context provided as key-value pairs.
// The variadic key-value pairs are treated as they are in With.
func (ctx Ctx) Info(msg string, keysAndValues ...interface{}) {
	panic("encore apps must be run using the encore command")
}

// Error logs an error-level message, merging the context from ctx
// with the additional context provided as key-value pairs.
// The variadic key-value pairs are treated as they are in With.
func (ctx Ctx) Error(msg string, keysAndValues ...interface{}) {
	panic("encore apps must be run using the encore command")
}
