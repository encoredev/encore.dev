package cache

import (
	"context"
)

// NewSetKeyspace creates a keyspace that stores unordered sets in the given cluster.
//
// The type parameter K specifies the key type, which can either be a
// named struct type or a basic type (string, int, etc).
//
// The type parameter V specifies the value type, which is the type
// of the elements in each set. It must be a basic type (string, int, int64, or float64).
func NewSetKeyspace[K any, V BasicType](cluster *Cluster, cfg KeyspaceConfig) (_ *SetKeyspace[K, V]) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.37.0/runtimes/go/storage/cache/set.go#L17-L24
	doPanic("encore apps must be run using the encore command")
	return
}

// SetKeyspace represents a set of cache keys,
// each containing an unordered set of values of type V.
type SetKeyspace[K any, V BasicType] struct {
	_ int // for godoc to show unexported fields
}

// With returns a reference to the same keyspace but with customized write options.
// The primary use case is for overriding the expiration time for certain cache operations.
//
// It is intended to be used with method chaining:
//
//	myKeyspace.With(cache.ExpireIn(3 * time.Second)).Set(...)
func (*SetKeyspace[K, V]) With(opts ...WriteOption) (_ *SetKeyspace[K, V]) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.37.0/runtimes/go/storage/cache/set.go#L38-L40
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
func (*SetKeyspace[K, V]) Delete(ctx context.Context, keys ...K) (deleted int, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.37.0/runtimes/go/storage/cache/set.go#L49-L51
	doPanic("encore apps must be run using the encore command")
	return
}

// Add adds one or more values to the set stored at key.
// If the key does not already exist, it is first created as an empty set.
//
// It reports the number of values that were added to the set,
// not including values already present beforehand.
//
// See https://redis.io/commands/sadd/ for more information.
func (*SetKeyspace[K, V]) Add(ctx context.Context, key K, values ...V) (added int, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.37.0/runtimes/go/storage/cache/set.go#L60-L75
	doPanic("encore apps must be run using the encore command")
	return
}

// Remove removes one or more values from the set stored at key.
//
// If a value is not present in the set is it ignored.
//
// Remove reports the number of values that were removed from the set.
// If the key does not already exist, it is a no-op and reports 0, nil.
//
// See https://redis.io/commands/srem/ for more information.
func (*SetKeyspace[K, V]) Remove(ctx context.Context, key K, values ...V) (removed int, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.37.0/runtimes/go/storage/cache/set.go#L85-L100
	doPanic("encore apps must be run using the encore command")
	return
}

// PopOne removes a random element from the set stored at key and returns it.
//
// If the set is empty it reports an error matching Miss.
//
// See https://redis.io/commands/spop/ for more information.
func (*SetKeyspace[K, V]) PopOne(ctx context.Context, key K) (val V, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.37.0/runtimes/go/storage/cache/set.go#L107-L124
	doPanic("encore apps must be run using the encore command")
	return
}

// Pop removes up to 'count' random elements (bounded by the set's size)
// from the set stored at key and returns them.
//
// If the set is empty it returns an empty slice and no error.
//
// See https://redis.io/commands/spop/ for more information.
func (*SetKeyspace[K, V]) Pop(ctx context.Context, key K, count int) (values []V, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.37.0/runtimes/go/storage/cache/set.go#L132-L149
	doPanic("encore apps must be run using the encore command")
	return
}

// Contains reports whether the set stored at key contains the given value.
//
// If the key does not exist it reports false, nil.
//
// See https://redis.io/commands/sismember/ for more information.
func (*SetKeyspace[K, V]) Contains(ctx context.Context, key K, val V) (contains bool, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.37.0/runtimes/go/storage/cache/set.go#L156-L167
	doPanic("encore apps must be run using the encore command")
	return
}

// Len reports the number of elements in the set stored at key.
//
// If the key does not exist it reports 0, nil.
//
// See https://redis.io/commands/slen/ for more information.
func (*SetKeyspace[K, V]) Len(ctx context.Context, key K) (length int64, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.37.0/runtimes/go/storage/cache/set.go#L174-L185
	doPanic("encore apps must be run using the encore command")
	return
}

// Items returns the elements in the set stored at key.
//
// If the key does not exist it returns an empty slice and no error.
//
// See https://redis.io/commands/smembers/ for more information.
func (*SetKeyspace[K, V]) Items(ctx context.Context, key K) (values []V, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.37.0/runtimes/go/storage/cache/set.go#L192-L204
	doPanic("encore apps must be run using the encore command")
	return
}

// ItemsMap is identical to Items except it returns the values as a map.
//
// If the key does not exist it returns an empty (but non-nil) map and no error.
//
// See https://redis.io/commands/smembers/ for more information.
func (*SetKeyspace[K, V]) ItemsMap(ctx context.Context, key K) (values map[V]struct{}, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.37.0/runtimes/go/storage/cache/set.go#L211-L224
	doPanic("encore apps must be run using the encore command")
	return
}

// Diff computes the set difference, between the first set and all the consecutive sets.
//
// Set difference means the values present in the first set that are not present
// in any of the other sets.
//
// Keys that do not exist are considered as empty sets.
//
// At least one key must be provided: if no keys are provided an error is reported.
//
// See https://redis.io/commands/sdiff/ for more information.
func (*SetKeyspace[K, V]) Diff(ctx context.Context, keys ...K) (_ []V, _ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.37.0/runtimes/go/storage/cache/set.go#L242-L249
	doPanic("encore apps must be run using the encore command")
	return
}

// DiffMap is identical to Diff except it returns the values as a map.
//
// See https://redis.io/commands/sdiff/ for more information.
func (*SetKeyspace[K, V]) DiffMap(ctx context.Context, keys ...K) (_ map[V]struct{}, _ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.37.0/runtimes/go/storage/cache/set.go#L254-L261
	doPanic("encore apps must be run using the encore command")
	return
}

// DiffStore computes the set difference between keys (like Diff) and stores the result in destination.
//
// It reports the size of the resulting set.
//
// See https://redis.io/commands/sdiffstore/ for more information.
func (*SetKeyspace[K, V]) DiffStore(ctx context.Context, destination K, keys ...K) (size int64, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.37.0/runtimes/go/storage/cache/set.go#L282-L301
	doPanic("encore apps must be run using the encore command")
	return
}

// Intersect computes the set intersection between the sets stored at the given keys.
//
// Set intersection means the values common to all the provided sets.
//
// Keys that do not exist are considered to be empty sets.
// As a result, if any key is missing the final result is the empty set.
//
// At least one key must be provided: if no keys are provided an error is reported.
//
// See https://redis.io/commands/sinter/ for more information.
func (*SetKeyspace[K, V]) Intersect(ctx context.Context, keys ...K) (_ []V, _ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.37.0/runtimes/go/storage/cache/set.go#L313-L320
	doPanic("encore apps must be run using the encore command")
	return
}

// IntersectMap is identical to Intersect except it returns the values as a map.
//
// See https://redis.io/commands/sinter/ for more information.
func (*SetKeyspace[K, V]) IntersectMap(ctx context.Context, keys ...K) (_ map[V]struct{}, _ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.37.0/runtimes/go/storage/cache/set.go#L325-L332
	doPanic("encore apps must be run using the encore command")
	return
}

// IntersectStore computes the set intersection between keys (like Intersect) and stores the result in destination.
//
// It reports the size of the resulting set.
//
// See https://redis.io/commands/sinterstore/ for more information.
func (*SetKeyspace[K, V]) IntersectStore(ctx context.Context, destination K, keys ...K) (size int64, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.37.0/runtimes/go/storage/cache/set.go#L354-L371
	doPanic("encore apps must be run using the encore command")
	return
}

// Union computes the set union between the sets stored at the given keys.
//
// Set union means the values present in at least one of the provided sets.
//
// Keys that do not exist are considered to be empty sets.
//
// At least one key must be provided: if no keys are provided an error is reported.
//
// See https://redis.io/commands/sunion/ for more information.
func (*SetKeyspace[K, V]) Union(ctx context.Context, keys ...K) (_ []V, _ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.37.0/runtimes/go/storage/cache/set.go#L382-L389
	doPanic("encore apps must be run using the encore command")
	return
}

// UnionMap is identical to Union except it returns the values as a map.
//
// See https://redis.io/commands/sunion/ for more information.
func (*SetKeyspace[K, V]) UnionMap(ctx context.Context, keys ...K) (_ map[V]struct{}, _ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.37.0/runtimes/go/storage/cache/set.go#L394-L401
	doPanic("encore apps must be run using the encore command")
	return
}

// UnionStore computes the set union between sets (like Union) and stores the result in destination.
//
// It reports the size of the resulting set.
//
// See https://redis.io/commands/sunionstore/ for more information.
func (*SetKeyspace[K, V]) UnionStore(ctx context.Context, destination K, keys ...K) (size int64, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.37.0/runtimes/go/storage/cache/set.go#L424-L440
	doPanic("encore apps must be run using the encore command")
	return
}

// SampleOne returns a random member from the set stored at key.
//
// If the key does not exist it reports an error matching Miss.
//
// See https://redis.io/commands/srandmember/ for more information.
func (*SetKeyspace[K, V]) SampleOne(ctx context.Context, key K) (val V, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.37.0/runtimes/go/storage/cache/set.go#L447-L461
	doPanic("encore apps must be run using the encore command")
	return
}

// Sample returns up to 'count' distinct random elements from the set stored at key.
// The same element is never returned multiple times.
//
// If the key does not exist it returns an empty slice and no error.
//
// See https://redis.io/commands/srandmember/ for more information.
func (*SetKeyspace[K, V]) Sample(ctx context.Context, key K, count int) (values []V, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.37.0/runtimes/go/storage/cache/set.go#L469-L490
	doPanic("encore apps must be run using the encore command")
	return
}

// SampleWithReplacement returns count random elements from the set stored at key.
// The same element may be returned multiple times.
//
// If the key does not exist it returns an empty slice and no error.
//
// See https://redis.io/commands/srandmember/ for more information.
func (*SetKeyspace[K, V]) SampleWithReplacement(ctx context.Context, key K, count int) (values []V, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.37.0/runtimes/go/storage/cache/set.go#L498-L519
	doPanic("encore apps must be run using the encore command")
	return
}

// Move atomically moves the given value from the set stored at src
// to the set stored at dst.
//
// If the element does not exist in src it reports false, nil.
//
// If the element already exists in dst it is still removed from src
// and Move still reports true, nil.
func (*SetKeyspace[K, V]) Move(ctx context.Context, src, dst K, val V) (moved bool, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.37.0/runtimes/go/storage/cache/set.go#L528-L541
	doPanic("encore apps must be run using the encore command")
	return
}
