// Package config provides a simple way to access configuration values for a
// service using the Load function.
//
// # By default configuration is pulled at build time from CUE files in each service directory
//
// For more information about configuration see https://encore.dev/docs/develop/config.
package config

import "os"

// Load returns the fully loaded configuration for this service.
//
// The configuration is loaded from the CUE files in the service directory and
// will be validated by Encore at compile time, which ensures this function will
// return a valid configuration at runtime.
//
// Encore will generate a `encore.gen.cue` file in the service directory which
// will contain generated CUE matching the configuration type T.
//
// Note: This function can only be called from within services and cannot be
// referenced from other services.
func Load[T any]() (_ T) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.44.6/runtimes/go/config/pkgfn.go#L32-L54
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
