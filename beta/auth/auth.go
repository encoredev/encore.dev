// Package auth provides the APIs to get information about the authenticated users.
//
// For more information about how authentication works with Encore applications see https://encore.dev/docs/develop/auth.
package auth

import (
	"context"
	"os"
)

// UID is a unique identifier representing a user (a user id).
type UID string

// WithContext returns a new context that sets the auth information for outgoing API calls.
// It does not affect the auth information for the current request.
//
// Passing in an empty string as the uid results in unsetting the auth information,
// causing future API calls to behave as if there was no authenticated user.
//
// If the application's auth handler returns custom auth data, two additional
// requirements exist. First, the data parameter passed to WithContext must be of
// the same type as the auth handler returns. Second, if the uid argument is not
// the empty string then data may not be nil. If these requirements are not met,
// API calls made with these options will not be made and will immediately return
// a client-side error.
func WithContext(ctx context.Context, uid UID, data interface{}) (_ context.Context) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/38dbb67953561748253891e3e50181bdf095d6e5/runtime/beta/auth/auth.go#L63-L67
	doPanic("encore apps must be run using the encore command")
	return
}

// doPanic is a wrapper around panic to prevent static analysis tools
// from thinking Encore APIs unconditionally panic.,
func doPanic(v any) {
	if os.Getenv("ENCORERUNTIME_NOPANIC") == "" {
		panic(v)
	}
}
