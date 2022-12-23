package et

import (
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
	//    https://github.com/encoredev/encore/blob/v1.12.0/runtime/et/config.go#L13-L23
	panic("encore apps must be run using the encore command")
}
