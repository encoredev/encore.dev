package cache

import "context"

// NewStructKeyspace creates a keyspace that stores structs in the given cluster.
//
// The type parameter K specifies the key type, which can either be a
// named struct type or a basic type (string, int, etc).
//
// The value parameter V specifies the named struct type that should be stored.
func NewStructKeyspace[K, V any](cluster *Cluster, cfg KeyspaceConfig) (_ *StructKeyspace[K, V]) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.17.0/runtime/storage/cache/struct.go#L11-L27
	doPanic("encore apps must be run using the encore command")
	return
}

// StructKeyspace represents a set of cache keys that hold struct values.
type StructKeyspace[K, V any] struct {
	_ int // for godoc to show unexported fields
}

// With returns a reference to the same keyspace but with customized write options.
// The primary use case is for overriding the expiration time for certain cache operations.
//
// It is intended to be used with method chaining:
//
//	myKeyspace.With(cache.ExpireIn(3 * time.Second)).Set(...)
func (*StructKeyspace[K, V]) With(opts ...WriteOption) (_ *StructKeyspace[K, V]) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.17.0/runtime/storage/cache/struct.go#L39-L41
	doPanic("encore apps must be run using the encore command")
	return
}

// Get gets the value stored at key.
// If the key does not exist, it returns an error matching Miss.
//
// See https://redis.io/commands/get/ for more information.
func (*StructKeyspace[K, V]) Get(ctx context.Context, key K) (_ V, _ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.17.0/runtime/storage/cache/struct.go#L47-L49
	doPanic("encore apps must be run using the encore command")
	return
}

// Set updates the value stored at key to val.
//
// See https://redis.io/commands/set/ for more information.
func (*StructKeyspace[K, V]) Set(ctx context.Context, key K, val V) (_ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.17.0/runtime/storage/cache/struct.go#L54-L56
	doPanic("encore apps must be run using the encore command")
	return
}

// SetIfNotExists sets the value stored at key to val, but only if the key does not exist beforehand.
// If the key already exists, it reports an error matching KeyExists.
//
// See https://redis.io/commands/setnx/ for more information.
func (*StructKeyspace[K, V]) SetIfNotExists(ctx context.Context, key K, val V) (_ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.17.0/runtime/storage/cache/struct.go#L62-L64
	doPanic("encore apps must be run using the encore command")
	return
}

// Replace replaces the existing value stored at key to val.
// If the key does not already exist, it reports an error matching Miss.
//
// See https://redis.io/commands/set/ for more information.
func (*StructKeyspace[K, V]) Replace(ctx context.Context, key K, val V) (_ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.17.0/runtime/storage/cache/struct.go#L70-L72
	doPanic("encore apps must be run using the encore command")
	return
}

// GetAndSet updates the value of key to val and returns the previously stored value.
// If the key does not already exist, it sets it and returns 0, nil.
//
// See https://redis.io/commands/getset/ for more information.
func (*StructKeyspace[K, V]) GetAndSet(ctx context.Context, key K, val V) (oldVal V, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.17.0/runtime/storage/cache/struct.go#L78-L80
	doPanic("encore apps must be run using the encore command")
	return
}

// GetAndDelete deletes the key and returns the previously stored value.
// If the key does not already exist, it does nothing and returns the zero value of V and nil.
//
// See https://redis.io/commands/getdel/ for more information.
func (*StructKeyspace[K, V]) GetAndDelete(ctx context.Context, key K) (oldVal V, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.17.0/runtime/storage/cache/struct.go#L86-L88
	doPanic("encore apps must be run using the encore command")
	return
}

// Delete deletes the specified keys.
//
// If a key does not exist it is ignored.
//
// It reports the number of keys that were deleted.
//
// See https://redis.io/commands/del/ for more information.
func (*StructKeyspace[K, V]) Delete(ctx context.Context, keys ...K) (deleted int, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.17.0/runtime/storage/cache/struct.go#L97-L99
	doPanic("encore apps must be run using the encore command")
	return
}
