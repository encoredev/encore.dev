package rlog

// Debug logs a debug-level message.
// The variadic key-value pairs are treated as they are in With.
func Debug(msg string, keysAndValues ...interface{}) {
	panic("not implemented")
}

// Info logs an info-level message.
// The variadic key-value pairs are treated as they are in With.
func Info(msg string, keysAndValues ...interface{}) {
	panic("not implemented")
}

// Error logs an error-level message.
// The variadic key-value pairs are treated as they are in With.
func Error(msg string, keysAndValues ...interface{}) {
	panic("not implemented")
}

// Ctx holds additional logging context for use with the Infoc and family
// of logging functions.
type Ctx struct{}

// With adds a variadic number of fields to the logging context.
// The keysAndValues must be pairs of string keys and arbitrary data.
func With(keysAndValues ...interface{}) Ctx {
	panic("not implemented")
}

// With creates a new logging context that inherits the context
// from the original ctx and adds additional context on top.
// The original ctx is not affected.
func (ctx Ctx) With(keysAndValues ...interface{}) Ctx {
	panic("not implemented")
}

// Debugc is like Debug but takes additional logging context.
// The variadic key-value pairs are treated as they are in With.
func Debugc(ctx Ctx, msg string, keysAndValues ...interface{}) {
	panic("not implemented")
}

// Infoc is like Debug but takes additional logging context.
// The variadic key-value pairs are treated as they are in With.
func Infoc(ctx Ctx, msg string, keysAndValues ...interface{}) {
	panic("not implemented")
}

// Errorc is like Error but takes additional logging context.
// The variadic key-value pairs are treated as they are in With.
func Errorc(ctx Ctx, msg string, keysAndValues ...interface{}) {
	panic("not implemented")
}
