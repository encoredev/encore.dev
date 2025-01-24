package et

import (
	"os"

	"encore.dev/config"
)

// SetCfg changes the value of cfg to newValue within the current test and any subtests.
// Other tests running will not be affected.
//
// It does not support setting slices and panics if given a config value that is a slice.
func SetCfg[T any](cfg config.Value[T], newValue T) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.46.1/runtimes/go/et/config.go#L15-L25
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
