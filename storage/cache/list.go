package cache

import (
	"context"
)

// BasicType is a constraint for basic types that
// can be used as element types in Redis lists and sets.
type BasicType interface {
	string | int | int64 | float64
}

// NewListKeyspace creates a keyspace that stores ordered lists in the given cluster.
//
// The type parameter K specifies the key type, which can either be a
// named struct type or a basic type (string, int, etc).
//
// The type parameter V specifies the value type, which is the type
// of the elements in each list. It must be a basic type (string, int, int64, or float64).
func NewListKeyspace[K any, V BasicType](cluster *Cluster, cfg KeyspaceConfig) *ListKeyspace[K, V] {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.11.0/runtime/storage/cache/list.go#L25-L32
	panic("encore apps must be run using the encore command")
}

// ListKeyspace represents a set of cache keys,
// each containing an ordered list of values of type V.
type ListKeyspace[K any, V BasicType] struct {
	_ int // for godoc to show unexported fields
}

// With returns a reference to the same keyspace but with customized write options.
// The primary use case is for overriding the expiration time for certain cache operations.
//
// It is intended to be used with method chaining:
//
//	myKeyspace.With(cache.ExpireIn(3 * time.Second)).Set(...)
func (*ListKeyspace[K, V]) With(opts ...WriteOption) *ListKeyspace[K, V] {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.11.0/runtime/storage/cache/list.go#L45-L47
	panic("encore apps must be run using the encore command")
}

// Delete deletes the specified keys.
//
// If a key does not exist it is ignored.
//
// It reports the number of keys that were deleted.
//
// See https://redis.io/commands/del/ for more information.
func (*ListKeyspace[K, V]) Delete(ctx context.Context, keys ...K) (deleted int, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.11.0/runtime/storage/cache/list.go#L56-L58
	panic("encore apps must be run using the encore command")
}

// PushLeft pushes one or more values at the head of the list stored at key.
// If the key does not already exist, it is first created as an empty list.
//
// If multiple values are given, they are inserted one after another,
// starting with the leftmost value. For instance,
//
//	PushLeft(ctx, "mylist", "a", "b", "c")
//
// will result in a list containing "c" as its first element,
// "b" as its second element, and "a" as its third element.
//
// See https://redis.io/commands/lpush/ for more information.
func (*ListKeyspace[K, V]) PushLeft(ctx context.Context, key K, values ...V) (newLen int64, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.11.0/runtime/storage/cache/list.go#L70-L85
	panic("encore apps must be run using the encore command")
}

// PushRight pushes one or more values at the tail of the list stored at key.
// If the key does not already exist, it is first created as an empty list.
//
// If multiple values are given, they are inserted one after another,
// starting with the leftmost value. For instance,
//
//	PushRight(ctx, "mylist", "a", "b", "c")
//
// will result in a list containing "c" as its last element,
// "b" as its second to last element, and "a" as its third-to-last element.
//
// See https://redis.io/commands/rpush/ for more information.
func (*ListKeyspace[K, V]) PushRight(ctx context.Context, key K, values ...V) (newLen int64, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.11.0/runtime/storage/cache/list.go#L97-L112
	panic("encore apps must be run using the encore command")
}

// PopLeft pops a single element off the head of the list stored at key and returns it.
// If the key does not exist, it returns an error matching Miss.
//
// See https://redis.io/commands/lpop/ for more information.
func (*ListKeyspace[K, V]) PopLeft(ctx context.Context, key K) (val V, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.11.0/runtime/storage/cache/list.go#L118-L135
	panic("encore apps must be run using the encore command")
}

// PopRight pops a single element off the tail of the list stored at key and returns it.
// If the key does not exist, it returns an error matching Miss.
//
// See https://redis.io/commands/rpop/ for more information.
func (*ListKeyspace[K, V]) PopRight(ctx context.Context, key K) (val V, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.11.0/runtime/storage/cache/list.go#L141-L157
	panic("encore apps must be run using the encore command")
}

// Len reports the length of the list stored at key.
//
// Non-existing keys are considered as empty lists.
//
// See https://redis.io/commands/llen/ for more information.
func (*ListKeyspace[K, V]) Len(ctx context.Context, key K) (length int64, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.11.0/runtime/storage/cache/list.go#L164-L175
	panic("encore apps must be run using the encore command")
}

// Trim trims the list stored at key to only contain the elements between the indices
// start and stop (inclusive). Both start and stop are zero-based indices.
//
// Negative indices can be used to indicate offsets from the end of the list,
// where -1 is the last element of the list, -2 the penultimate element, and so on.
//
// Out of range indices are valid and are treated as if they specify the start or end of the list,
// respectively. If start > stop the end result is an empty list.
//
// See https://redis.io/commands/ltrim/ for more information.
func (*ListKeyspace[K, V]) Trim(ctx context.Context, key K, start, stop int64) error {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.11.0/runtime/storage/cache/list.go#L187-L198
	panic("encore apps must be run using the encore command")
}

// Set updates the list element with the given idx to val.
//
// Negative indices can be used to indicate offsets from the end of the list,
// where -1 is the last element of the list, -2 the penultimate element, and so on.
//
// An error is returned for out of bounds indices.
//
// See https://redis.io/commands/lset/ for more information.
func (*ListKeyspace[K, V]) Set(ctx context.Context, key K, idx int64, val V) (err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.11.0/runtime/storage/cache/list.go#L208-L220
	panic("encore apps must be run using the encore command")
}

// Get returns the value of list element with the given idx.
//
// Negative indices can be used to indicate offsets from the end of the list,
// where -1 is the last element of the list, -2 the penultimate element, and so on.
//
// For out of bounds indices or the key not existing in the cache, an error matching Miss is returned.
//
// See https://redis.io/commands/lget/ for more information.
func (*ListKeyspace[K, V]) Get(ctx context.Context, key K, idx int64) (val V, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.11.0/runtime/storage/cache/list.go#L230-L244
	panic("encore apps must be run using the encore command")
}

// Items returns all the elements in the list stored at key.
//
// If the key does not exist it returns an empty list.
//
// See https://redis.io/commands/lrange/ for more information.
func (*ListKeyspace[K, V]) Items(ctx context.Context, key K) ([]V, error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.11.0/runtime/storage/cache/list.go#L251-L253
	panic("encore apps must be run using the encore command")
}

// GetRange returns all the elements in the list stored at key between start and stop.
//
// Negative indices can be used to indicate offsets from the end of the list,
// where -1 is the last element of the list, -2 the penultimate element, and so on.
//
// If the key does not exist it returns an empty list.
//
// See https://redis.io/commands/lrange/ for more information.
func (*ListKeyspace[K, V]) GetRange(ctx context.Context, key K, start, stop int64) ([]V, error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.11.0/runtime/storage/cache/list.go#L263-L265
	panic("encore apps must be run using the encore command")
}

// InsertBefore inserts newVal into the list stored at key, at the position just before needle.
//
// If the list stored at key does not contain needle the value is not inserted,
// and an error matching Miss is returned.
//
// It returns the new list length.
//
// See https://redis.io/commands/linsert/ for more information.
func (*ListKeyspace[K, V]) InsertBefore(ctx context.Context, key K, needle, newVal V) (newLen int64, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.11.0/runtime/storage/cache/list.go#L297-L313
	panic("encore apps must be run using the encore command")
}

// InsertAfter inserts newVal into the list stored at key, at the position just after needle.
//
// It reports the new list length.
//
// If the list stored at key does not contain needle the value is not inserted,
// and an error matching Miss is returned.
//
// See https://redis.io/commands/linsert/ for more information.
func (*ListKeyspace[K, V]) InsertAfter(ctx context.Context, key K, needle, newVal V) (newLen int64, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.11.0/runtime/storage/cache/list.go#L323-L339
	panic("encore apps must be run using the encore command")
}

// RemoveAll removes all values equal to needle in the list stored at key.
//
// It reports the number of elements removed.
//
// If the list does not contain needle, or the list does not exist, it reports 0, nil.
//
// See https://redis.io/commands/lrem/ for more information.
func (*ListKeyspace[K, V]) RemoveAll(ctx context.Context, key K, needle V) (removed int64, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.11.0/runtime/storage/cache/list.go#L348-L361
	panic("encore apps must be run using the encore command")
}

// RemoveFirst removes the first 'count' values equal to needle in the list stored at key.
//
// It reports the number of elements removed.
//
// If the list does not contain needle, or the list does not exist, it reports 0, nil.
//
// See https://redis.io/commands/lrem/ for more information.
func (*ListKeyspace[K, V]) RemoveFirst(ctx context.Context, key K, count int64, needle V) (removed int64, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.11.0/runtime/storage/cache/list.go#L370-L390
	panic("encore apps must be run using the encore command")
}

// RemoveLast removes the last 'count' values equal to needle in the list stored at key.
//
// It reports the number of elements removed.
//
// If the list does not contain needle, or the list does not exist, it reports 0, nil.
//
// See https://redis.io/commands/lrem/ for more information.
func (*ListKeyspace[K, V]) RemoveLast(ctx context.Context, key K, count int64, needle V) (removed int64, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.11.0/runtime/storage/cache/list.go#L399-L419
	panic("encore apps must be run using the encore command")
}

type ListPos string

const (
	Left  ListPos = "LEFT"
	Right ListPos = "RIGHT"
)

// Move atomically moves an element from the list stored at src to the list stored at dst.
//
// The value moved can be either the head (fromPos == Left) or tail (fromPos == Right) of the list at src.
// Similarly, the value can be placed either at the head (toPos == Left) or tail (toPos == Right) of the list at dst.
//
// If src does not exist it reports an error matching Miss.
//
// If src and dst are the same list, the value is atomically rotated from one end to the other when fromPos != toPos,
// or if fromPos == toPos nothing happens.
func (*ListKeyspace[K, V]) Move(ctx context.Context, src, dst K, fromPos, toPos ListPos) (moved V, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.11.0/runtime/storage/cache/list.go#L437-L455
	panic("encore apps must be run using the encore command")
}
