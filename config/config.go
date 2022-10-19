// Package config provides a simple way to access configuration values for a
// service using the Load function.
//
// # By default configuration is pulled at build time from CUE files in each service directory
//
// For more information about configuration see https://encore.dev/docs/develop/config.
package config

// Load returns the fully loaded configuration for this service.
//
// Note: This function can only be called from within services.
func Load[T any]() T {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/49a2d858ee8ab00336b162540061e232e9d3f70e/runtime/config/config.go#L32-L52
	panic("encore apps must be run using the encore command")
}
