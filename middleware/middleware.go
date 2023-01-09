// Package middleware provides middleware functionality for defining
// generic processing across multiple API endpoints, typically for
// cross-cutting concerns like validation, caching, or error monitoring.
//
// For documentation on how to use middleware in Encore see https://encore.dev/docs/develop/middleware.
package middleware

import (
	"context"

	encore "encore.dev"
)

// Signature is the signature a middleware should have.
// The implementation must return a Response which then
// becomes the API's response.
//
// Middleware processing forms a chain, where the first
// defined middleware is called first. The provided next
// function is used to invoke the next middleware in the
// chain, or the actual API handler if it's the last middleware.
//
// The implementation should call next at most once.
// If the middleware encounters an error it is permissible to
// not call next and directly return a Response with Err set
// to a non-nil value.
//
// The type exists solely for documentation purposes;
// it is otherwise not used.
type Signature func(req Request, next Next) Response

// Next is the function to invoke the next middleware in the chain,
// or the actual API handler if the middleware is the last one.
type Next func(Request) Response

// A Request provides information about the incoming request that
// the middleware is processing.
type Request struct {
	_ int // for godoc to show unexported fields
}

// WithContext returns a new Request with the context set to ctx.
func (*Request) WithContext(ctx context.Context) (_ Request) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/4d212a6471c0a6f5e7df1114b5238c8084d76c07/runtime/middleware/middleware.go#L44-L48
	doPanic("encore apps must be run using the encore command")
	return
}

// Context reports the request's context.
func (*Request) Context() (_ context.Context) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/4d212a6471c0a6f5e7df1114b5238c8084d76c07/runtime/middleware/middleware.go#L51-L53
	doPanic("encore apps must be run using the encore command")
	return
}

// Data returns information about the request.
func (*Request) Data() (_ *encore.Request) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/4d212a6471c0a6f5e7df1114b5238c8084d76c07/runtime/middleware/middleware.go#L56-L58
	doPanic("encore apps must be run using the encore command")
	return
}

// Response represents the API handler's response.
// It can be modified by each middleware to customize
// the API response.
//
// The Payload's type must not be changed: it is not permissible
// to return a payload with a different type than the API returns.
//
// If Err is non-nil it becomes serialized as the response.
//
// For Raw endpoints neither Payload nor Err are used.
type Response struct {
	// Payload is the API's response payload.
	// It is nil if the API handler returned an error or if
	// the API is a raw endpoint.
	//
	// The payload can be mutated by middleware but the type must not be changed;
	// it must always be of the same type as the API handler's return type.
	// Otherwise, an error is returned.
	Payload any

	// Err is the error returned from the API handler or another middleware.
	// It is written as output for non-raw endpoints.
	Err error

	// HTTPStatus is the HTTP status code the response is written with.
	//
	// If zero is returned from middleware, Encore will choose an appropriate status code based
	// status code depending on the type of error being returned (or 200 for success).
	//
	// If a non-zero value is returned from a middleware, Encore will use that status code
	// regardless of what Err is set to.
	//
	// For raw handlers middleware cannot modify this as it has already
	// been written to the network.
	HTTPStatus int
}

// NewRequest constructs a new Request that returns the given context and request data.
// It is primarily used for testing middleware.
func NewRequest(ctx context.Context, data *encore.Request) (_ Request) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/4d212a6471c0a6f5e7df1114b5238c8084d76c07/runtime/middleware/middleware.go#L99-L104
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
