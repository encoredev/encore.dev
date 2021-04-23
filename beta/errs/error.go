// Package errs provides structured error handling for Encore applications.
package errs

import (
	"net/http"
)

// An Error is an error that provides structured information
// about the error. It includes an error code, a message,
// optionally additional structured details about the error
// and arbitrary key-value metadata.
//
// The Details field is returned to external clients.
// The Meta field is only exposed to internal calls within Encore.
//
// Internally it captures an underlying error for printing
// and for use with errors.Is/As and call stack information.
//
// To provide accurate stack information, users are expected
// to convert non-Error errors into *Error as close to the
// root cause as possible. This is made simple with Wrap.
type Error struct {
	// Code is the error code to return.
	Code ErrCode `json:"code"`
	// Message is a descriptive message of the error.
	Message string `json:"message"`
	// Details are user-defined additional details.
	Details ErrDetails `json:"details"`
	// Meta are arbitrary key-value pairs for use within
	// the Encore application. They are not exposed to external clients.
	Meta Metadata `json:"-"`
}

// Metadata represents structured key-value pairs, for attaching arbitrary
// metadata to errors. The metadata is propagated between internal services
// but is not exposed to external clients.
type Metadata map[string]interface{}

// Wrap wraps the err, adding additional error information.
// If err is nil it returns nil.
//
// If err is already an *Error its code, message, and details
// are copied over to the new error.
func Wrap(err error, msg string, metaPairs ...interface{}) error {
	panic("encore apps must be run using the encore command")
}

// WrapCode is like Wrap but also sets the error code.
// If code is OK it reports nil.
func WrapCode(err error, code ErrCode, msg string, metaPairs ...interface{}) error {
	panic("encore apps must be run using the encore command")
}

// Convert converts an error to an *Error.
// If the error is already an *Error it returns it unmodified.
// If err is nil it returns nil.
func Convert(err error) error {
	panic("encore apps must be run using the encore command")
}

// Code reports the error code from an error.
// If err is nil it reports OK.
// Otherwise if err is not an *Error it reports Unknown.
func Code(err error) ErrCode {
	panic("encore apps must be run using the encore command")
}

// Meta reports the metadata included in the error.
// If err is nil or the error lacks metadata it reports nil.
func Meta(err error) Metadata {
	panic("encore apps must be run using the encore command")
}

// Details reports the error details included in the error.
// If err is nil or the error lacks details it reports nil.
func Details(err error) ErrDetails {
	panic("encore apps must be run using the encore command")
}

// Error reports the error code and message.
func (e *Error) Error() string {
	panic("encore apps must be run using the encore command")
}

// ErrorMessage reports the error message, joining this
// error's message with the messages from any underlying errors.
func (e *Error) ErrorMessage() string {
	panic("encore apps must be run using the encore command")
}

// Unwrap returns the underlying error, if any.
func (e *Error) Unwrap() error {
	panic("encore apps must be run using the encore command")
}

// HTTPError writes structured error information to w using JSON encoding.
// The status code is computed with HTTPStatus.
//
// If err is nil it writes:
// 	{"code": "ok", "message": "", "details": null}
func HTTPError(w http.ResponseWriter, err error) {
	panic("encore apps must be run using the encore command")
}
