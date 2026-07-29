package cache

import (
	"context"
	"os"
)

// NewStringKeyspace creates a keyspace that stores string values in the given cluster.
//
// The type parameter K specifies the key type, which can either be a
// named struct type or a basic type (string, int, etc).
func NewStringKeyspace[K any](cluster *Cluster, cfg KeyspaceConfig) (_ *StringKeyspace[K]) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/cache/basic.go#L16-L25
	doPanic("encore apps must be run using the encore command")
	return
}

// StringKeyspace represents a set of cache keys that hold string values.
type StringKeyspace[K any] struct {
	_ int // for godoc to show unexported fields
}

// Get gets the value stored at key.
// If the key does not exist, it returns an error matching Miss.
//
// See https://redis.io/commands/get/ for more information.
func (*StringKeyspace[K]) Get(ctx context.Context, key K) (_ string, _ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/cache/basic.go#L36-L38
	doPanic("encore apps must be run using the encore command")
	return
}

// MultiGet gets the values stored at multiple keys.
// For each key, the result contains an Err field indicating success or failure.
// If Err is nil, Value contains the cached value.
// If Err matches Miss, the key was not found.
//
// See https://redis.io/commands/mget/ for more information.
func (*StringKeyspace[K]) MultiGet(ctx context.Context, keys ...K) (_ []Result[string], _ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/cache/basic.go#L46-L48
	doPanic("encore apps must be run using the encore command")
	return
}

// MultiSet updates the values stored at multiple keys.
// The keyspace's expiry is applied to all keys.
//
// Use KV to construct the key-value pairs.
//
// See https://redis.io/commands/mset/ for more information.
func (*StringKeyspace[K]) MultiSet(ctx context.Context, entries ...KeyValue[K, string]) (_ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/cache/basic.go#L56-L58
	doPanic("encore apps must be run using the encore command")
	return
}

// Set updates the value stored at key to val.
//
// See https://redis.io/commands/set/ for more information.
func (*StringKeyspace[K]) Set(ctx context.Context, key K, val string) (_ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/cache/basic.go#L63-L65
	doPanic("encore apps must be run using the encore command")
	return
}

// SetIfNotExists sets the value stored at key to val, but only if the key does not exist beforehand.
// If the key already exists, it reports an error matching KeyExists.
//
// See https://redis.io/commands/setnx/ for more information.
func (*StringKeyspace[K]) SetIfNotExists(ctx context.Context, key K, val string) (_ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/cache/basic.go#L71-L73
	doPanic("encore apps must be run using the encore command")
	return
}

// Replace replaces the existing value stored at key to val.
// If the key does not already exist, it reports an error matching Miss.
//
// See https://redis.io/commands/set/ for more information.
func (*StringKeyspace[K]) Replace(ctx context.Context, key K, val string) (_ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/cache/basic.go#L79-L81
	doPanic("encore apps must be run using the encore command")
	return
}

// GetAndSet updates the value of key to val and returns the previously stored value.
// If the key does not already exist, it sets it and returns "", nil.
//
// See https://redis.io/commands/getset/ for more information.
func (*StringKeyspace[K]) GetAndSet(ctx context.Context, key K, val string) (oldVal string, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/cache/basic.go#L87-L89
	doPanic("encore apps must be run using the encore command")
	return
}

// GetAndDelete deletes the key and returns the previously stored value.
// If the key does not already exist, it does nothing and returns "", nil.
//
// See https://redis.io/commands/getdel/ for more information.
func (*StringKeyspace[K]) GetAndDelete(ctx context.Context, key K) (oldVal string, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/cache/basic.go#L95-L97
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
func (*StringKeyspace[K]) Delete(ctx context.Context, keys ...K) (deleted int, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/cache/basic.go#L106-L108
	doPanic("encore apps must be run using the encore command")
	return
}

// With returns a reference to the same keyspace but with customized write options.
// The primary use case is for overriding the expiration time for certain cache operations.
//
// It is intended to be used with method chaining:
//
//	myKeyspace.With(cache.ExpireIn(3 * time.Second)).Set(...)
func (*StringKeyspace[K]) With(opts ...WriteOption) (_ *StringKeyspace[K]) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/cache/basic.go#L116-L118
	doPanic("encore apps must be run using the encore command")
	return
}

// Append appends to the string with the given key.
//
// If the key does not exist it is first created and set as the empty string,
// causing Append to behave like Set.
//
// It returns the new string length.
//
// See https://redis.io/commands/append/ for more information.
func (*StringKeyspace[K]) Append(ctx context.Context, key K, val string) (newLen int64, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/cache/basic.go#L128-L142
	doPanic("encore apps must be run using the encore command")
	return
}

// GetRange returns a substring of the string value stored in key.
//
// The from and to values are zero-based indices, but unlike Go slicing
// the 'to' value is inclusive.
//
// Negative values can be used in order to provide an offset starting
// from the end of the string, so -1 means the last character
// and -len(str) the first character, and so forth.
//
// If the string does not exist it returns the empty string.
//
// See https://redis.io/commands/setrange/ for more information.
func (*StringKeyspace[K]) GetRange(ctx context.Context, key K, from, to int64) (val string, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/cache/basic.go#L156-L168
	doPanic("encore apps must be run using the encore command")
	return
}

// SetRange overwrites part of the string stored at key, starting at
// the zero-based offset and for the entire length of val, extending
// the string if necessary to make room for val.
//
// If the offset is larger than the current string length stored at key,
// the string is first padded with zero-bytes to make offset fit.
//
// Non-existing keys are considered as empty strings.
//
// See https://redis.io/commands/setrange/ for more information.
func (*StringKeyspace[K]) SetRange(ctx context.Context, key K, offset int64, val string) (newLen int64, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/cache/basic.go#L180-L194
	doPanic("encore apps must be run using the encore command")
	return
}

// Len reports the length of the string value stored at key.
//
// Non-existing keys are considered as empty strings.
//
// See https://redis.io/commands/strlen/ for more information.
func (*StringKeyspace[K]) Len(ctx context.Context, key K) (length int64, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/cache/basic.go#L201-L213
	doPanic("encore apps must be run using the encore command")
	return
}

// NewIntKeyspace creates a keyspace that stores int64 values in the given cluster.
//
// The type parameter K specifies the key type, which can either be a
// named struct type or a basic type (string, int, etc).
func NewIntKeyspace[K any](cluster *Cluster, cfg KeyspaceConfig) (_ *IntKeyspace[K]) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/cache/basic.go#L219-L228
	doPanic("encore apps must be run using the encore command")
	return
}

// IntKeyspace is a cache keyspace that stores int64 values.
type IntKeyspace[K any] struct {
	_ int // for godoc to show unexported fields
}

// With returns a reference to the same keyspace but with customized write options.
// The primary use case is for overriding the expiration time for certain cache operations.
//
// It is intended to be used with method chaining:
//
//	myKeyspace.With(cache.ExpireIn(3 * time.Second)).Set(...)
func (*IntKeyspace[K]) With(opts ...WriteOption) (_ *IntKeyspace[K]) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/cache/basic.go#L241-L243
	doPanic("encore apps must be run using the encore command")
	return
}

// Get gets the value stored at key.
// If the key does not exist, it returns an error matching Miss.
//
// See https://redis.io/commands/get/ for more information.
func (*IntKeyspace[K]) Get(ctx context.Context, key K) (_ int64, _ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/cache/basic.go#L249-L251
	doPanic("encore apps must be run using the encore command")
	return
}

// MultiGet gets the values stored at multiple keys.
// For each key, the result contains an Err field indicating success or failure.
// If Err is nil, Value contains the cached value.
// If Err matches Miss, the key was not found.
//
// See https://redis.io/commands/mget/ for more information.
func (*IntKeyspace[K]) MultiGet(ctx context.Context, keys ...K) (_ []Result[int64], _ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/cache/basic.go#L259-L261
	doPanic("encore apps must be run using the encore command")
	return
}

// MultiSet updates the values stored at multiple keys.
// The keyspace's expiry is applied to all keys.
//
// Use KV to construct the key-value pairs.
//
// See https://redis.io/commands/mset/ for more information.
func (*IntKeyspace[K]) MultiSet(ctx context.Context, entries ...KeyValue[K, int64]) (_ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/cache/basic.go#L269-L271
	doPanic("encore apps must be run using the encore command")
	return
}

// Set updates the value stored at key to val.
//
// See https://redis.io/commands/set/ for more information.
func (*IntKeyspace[K]) Set(ctx context.Context, key K, val int64) (_ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/cache/basic.go#L276-L278
	doPanic("encore apps must be run using the encore command")
	return
}

// SetIfNotExists sets the value stored at key to val, but only if the key does not exist beforehand.
// If the key already exists, it reports an error matching KeyExists.
//
// See https://redis.io/commands/setnx/ for more information.
func (*IntKeyspace[K]) SetIfNotExists(ctx context.Context, key K, val int64) (_ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/cache/basic.go#L284-L286
	doPanic("encore apps must be run using the encore command")
	return
}

// Replace replaces the existing value stored at key to val.
// If the key does not already exist, it reports an error matching Miss.
//
// See https://redis.io/commands/set/ for more information.
func (*IntKeyspace[K]) Replace(ctx context.Context, key K, val int64) (_ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/cache/basic.go#L292-L294
	doPanic("encore apps must be run using the encore command")
	return
}

// GetAndSet updates the value of key to val and returns the previously stored value.
// If the key does not already exist, it sets it and returns 0, nil.
//
// See https://redis.io/commands/getset/ for more information.
func (*IntKeyspace[K]) GetAndSet(ctx context.Context, key K, val int64) (oldVal int64, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/cache/basic.go#L300-L302
	doPanic("encore apps must be run using the encore command")
	return
}

// GetAndDelete deletes the key and returns the previously stored value.
// If the key does not already exist, it does nothing and returns 0, nil.
//
// See https://redis.io/commands/getdel/ for more information.
func (*IntKeyspace[K]) GetAndDelete(ctx context.Context, key K) (oldVal int64, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/cache/basic.go#L308-L310
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
func (*IntKeyspace[K]) Delete(ctx context.Context, keys ...K) (deleted int, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/cache/basic.go#L319-L321
	doPanic("encore apps must be run using the encore command")
	return
}

// Increment increments the number stored in key by delta,
// and returns the new value.
//
// If the key does not exist it is first created with a value of 0
// before incrementing.
//
// Negative values can be used to decrease the value,
// but typically you want to use the Decrement method for that.
//
// See https://redis.io/commands/incrby/ for more information.
func (*IntKeyspace[K]) Increment(ctx context.Context, key K, delta int64) (newVal int64, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/cache/basic.go#L333-L347
	doPanic("encore apps must be run using the encore command")
	return
}

// Decrement decrements the number stored in key by delta,
// and returns the new value.
//
// If the key does not exist it is first created with a value of 0
// before decrementing.
//
// Negative values can be used to increase the value,
// but typically you want to use the Increment method for that.
//
// See https://redis.io/commands/decrby/ for more information.
func (*IntKeyspace[K]) Decrement(ctx context.Context, key K, delta int64) (newVal int64, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/cache/basic.go#L359-L374
	doPanic("encore apps must be run using the encore command")
	return
}

// NewFloatKeyspace creates a keyspace that stores float64 values in the given cluster.
//
// The type parameter K specifies the key type, which can either be a
// named struct type or a basic type (string, int, etc).
func NewFloatKeyspace[K any](cluster *Cluster, cfg KeyspaceConfig) (_ *FloatKeyspace[K]) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/cache/basic.go#L380-L389
	doPanic("encore apps must be run using the encore command")
	return
}

// FloatKeyspace is a cache keyspace that stores float64 values.
type FloatKeyspace[K any] struct {
	_ int // for godoc to show unexported fields
}

// With returns a reference to the same keyspace but with customized write options.
// The primary use case is for overriding the expiration time for certain cache operations.
//
// It is intended to be used with method chaining:
//
//	myKeyspace.With(cache.ExpireIn(3 * time.Second)).Set(...)
func (*FloatKeyspace[K]) With(opts ...WriteOption) (_ *FloatKeyspace[K]) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/cache/basic.go#L402-L404
	doPanic("encore apps must be run using the encore command")
	return
}

// Get gets the value stored at key.
// If the key does not exist, it returns an error matching Miss.
//
// See https://redis.io/commands/get/ for more information.
func (*FloatKeyspace[K]) Get(ctx context.Context, key K) (_ float64, _ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/cache/basic.go#L410-L412
	doPanic("encore apps must be run using the encore command")
	return
}

// MultiGet gets the values stored at multiple keys.
// For each key, the result contains an Err field indicating success or failure.
// If Err is nil, Value contains the cached value.
// If Err matches Miss, the key was not found.
//
// See https://redis.io/commands/mget/ for more information.
func (*FloatKeyspace[K]) MultiGet(ctx context.Context, keys ...K) (_ []Result[float64], _ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/cache/basic.go#L420-L422
	doPanic("encore apps must be run using the encore command")
	return
}

// MultiSet updates the values stored at multiple keys.
// The keyspace's expiry is applied to all keys.
//
// Use KV to construct the key-value pairs.
//
// See https://redis.io/commands/mset/ for more information.
func (*FloatKeyspace[K]) MultiSet(ctx context.Context, entries ...KeyValue[K, float64]) (_ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/cache/basic.go#L430-L432
	doPanic("encore apps must be run using the encore command")
	return
}

// Set updates the value stored at key to val.
//
// See https://redis.io/commands/set/ for more information.
func (*FloatKeyspace[K]) Set(ctx context.Context, key K, val float64) (_ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/cache/basic.go#L437-L439
	doPanic("encore apps must be run using the encore command")
	return
}

// SetIfNotExists sets the value stored at key to val, but only if the key does not exist beforehand.
// If the key already exists, it reports an error matching KeyExists.
//
// See https://redis.io/commands/setnx/ for more information.
func (*FloatKeyspace[K]) SetIfNotExists(ctx context.Context, key K, val float64) (_ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/cache/basic.go#L445-L447
	doPanic("encore apps must be run using the encore command")
	return
}

// Replace replaces the existing value stored at key to val.
// If the key does not already exist, it reports an error matching Miss.
//
// See https://redis.io/commands/set/ for more information.
func (*FloatKeyspace[K]) Replace(ctx context.Context, key K, val float64) (_ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/cache/basic.go#L453-L455
	doPanic("encore apps must be run using the encore command")
	return
}

// GetAndSet updates the value of key to val and returns the previously stored value.
// If the key does not already exist, it sets it and returns 0, nil.
//
// See https://redis.io/commands/getset/ for more information.
func (*FloatKeyspace[K]) GetAndSet(ctx context.Context, key K, val float64) (oldVal float64, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/cache/basic.go#L461-L463
	doPanic("encore apps must be run using the encore command")
	return
}

// GetAndDelete deletes the key and returns the previously stored value.
// If the key does not already exist, it does nothing and returns 0, nil.
//
// See https://redis.io/commands/getdel/ for more information.
func (*FloatKeyspace[K]) GetAndDelete(ctx context.Context, key K) (oldVal float64, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/cache/basic.go#L469-L471
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
func (*FloatKeyspace[K]) Delete(ctx context.Context, keys ...K) (deleted int, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/cache/basic.go#L480-L482
	doPanic("encore apps must be run using the encore command")
	return
}

// Increment increments the number stored in key by delta,
// and returns the new value.
//
// If the key does not exist it is first created with a value of 0
// before incrementing.
//
// Negative values can be used to decrease the value,
// but typically you want to use the Decrement method for that.
//
// See https://redis.io/commands/incrbyfloat/ for more information.
func (*FloatKeyspace[K]) Increment(ctx context.Context, key K, delta float64) (newVal float64, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/cache/basic.go#L494-L508
	doPanic("encore apps must be run using the encore command")
	return
}

// Decrement decrements the number stored in key by delta,
// and returns the new value.
//
// If the key does not exist it is first created with a value of 0
// before decrementing.
//
// Negative values can be used to increase the value,
// but typically you want to use the Increment method for that.
//
// See https://redis.io/commands/incrbyfloat/ for more information.
func (*FloatKeyspace[K]) Decrement(ctx context.Context, key K, delta float64) (newVal float64, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.57.13/runtimes/go/storage/cache/basic.go#L520-L534
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
