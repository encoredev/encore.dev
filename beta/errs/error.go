package errs

import (
	"net/http"
)

// An Error is an error that provides structured information
// about the error. It includes an error code, a message, and
// optionally additional structured details about the error.
//
// Internally it captures an underlying error for printing
// and for use with errors.Is/As and call stack information.
//
// To provide accurate stack information, users are expected
// to convert non-Error errors into *Error as close to the
// root cause as possible. This is made simple with Wrap.
type Error struct {
	Code    ErrCode    `json:"code"`
	Message string     `json:"message"`
	Details ErrDetails `json:"details"`
}

// New creates a new Error without wrapping another underlying error.
// If code == OK it returns nil.
func New(code ErrCode, msg string, details ErrDetails) error {
	panic("encore apps must be run using the encore command")
}

// Wrap wraps the err, adding additional error information.
// If err is nil it returns nil.
//
// If err is already an *Error its code, message, and details
// are copied over to the new error.
//
// The fields are used to update the corresponding field of the error:
// Passing in an ErrCode updates the Code field.
// Passing in a string adds context to the error message.
// Passing in a type that implements ErrDetails updates the Details field,
// and passing in untyped nil sets Details to nil.
//
// Passing in another type causes Wrap to panic.
func Wrap(err error, fields ...interface{}) error {
	panic("encore apps must be run using the encore command")
}

// Code reports the error code from an error.
// If err is nil it reports OK.
// Otherwise if err is not an *Error it reports Unknown.
func Code(err error) ErrCode {
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
