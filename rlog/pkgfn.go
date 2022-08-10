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

// With adds a variadic number of fields to the logging context.
// The keysAndValues must be pairs of string keys and arbitrary data.
func With(keysAndValues ...interface{}) Ctx {
	panic("encore apps must be run using the encore command")
}
