package objects

// NewBucket declares a new object storage bucket.
//
// See https://encore.dev/docs/primitives/object-storage for more information.
func NewBucket(name string, cfg BucketConfig) (_ *Bucket) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.44.6/runtimes/go/storage/objects/objects.go#L8-L10
	doPanic("encore apps must be run using the encore command")
	return
}

// constStr is a string that can only be provided as a constant.
type constStr string

// Named returns a database object connected to the database with the given name.
//
// The name must be a string literal constant, to facilitate static analysis.
func Named(name constStr) (_ *Bucket) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.44.6/runtimes/go/storage/objects/objects.go#L20-L22
	doPanic("encore apps must be run using the encore command")
	return
}
