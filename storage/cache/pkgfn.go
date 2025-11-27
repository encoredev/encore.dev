package cache

// NewCluster declares a new cache cluster.
//
// See https://encore.dev/docs/develop/caching for more information.
func NewCluster(name string, cfg ClusterConfig) (_ *Cluster) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.52.1/runtimes/go/storage/cache/pkgfn.go#L8-L14
	doPanic("encore apps must be run using the encore command")
	return
}
