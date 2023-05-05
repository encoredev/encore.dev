// Package cache provides the ability to define distributed Redis cache clusters
// and functionality to use them in a fully type-safe manner.
//
// For more information see https://encore.dev/docs/develop/caching
package cache

import (
	"errors"
	"time"
)

// ClusterConfig represents the configuration of cache clusters.
type ClusterConfig struct {
	// EvictionPolicy decides how the cache evicts existing keys
	// to make room for new data when the max memory limit is reached.
	//
	// If not specified the cache defaults to AllKeysLRU.
	EvictionPolicy EvictionPolicy
}

// An EvictionPolicy describes how the cache evicts keys to make room for new data
// when the maximum memory limit is reached.
//
// See https://redis.io/docs/manual/eviction/#eviction-policies for more information.
type EvictionPolicy string

// The eviction policies Encore supports.
// See https://redis.io/docs/manual/eviction/#eviction-policies for more information.
const (
	// AllKeysLRU keeps most recently used keys and removes least recently used (LRU) keys.
	// This is a good default choice for most cache use cases if you're not sure.
	AllKeysLRU EvictionPolicy = "allkeys-lru"

	// AllKeysLFU keeps frequently used keys and removes least frequently used (LFU) keys.
	AllKeysLFU EvictionPolicy = "allkeys-lfu"

	// AllKeysRandom randomly removes keys as needed.
	AllKeysRandom EvictionPolicy = "allkeys-random"

	// VolatileLRU removes least recently used keys with an expiration set.
	// It behaves like NoEviction if there are no keys to evict with an expiration set.
	VolatileLRU EvictionPolicy = "volatile-lru"

	// VolatileLFU removes least frequently used keys with an expiration set.
	// It behaves like NoEviction if there are no keys to evict with an expiration set.
	VolatileLFU EvictionPolicy = "volatile-lfu"

	// VolatileTTL removes keys with an expiration set and the shortest time to live (TTL).
	// It behaves like NoEviction if there are no keys to evict with an expiration set.
	VolatileTTL EvictionPolicy = "volatile-ttl"

	// VolatileRandom randomly removes keys with an expiration set.
	// It behaves like NoEviction if there are no keys to evict with an expiration set.
	VolatileRandom EvictionPolicy = "volatile-random"

	// NoEviction does not evict any keys, and instead returns an error to the client
	// when the max memory limit is reached.
	NoEviction EvictionPolicy = "noeviction"
)

type constStr string

// Cluster represents a Redis cache cluster.
type Cluster struct {
	_ int // for godoc to show unexported fields
}

// KeyspaceConfig specifies the configuration options for a cache keyspace.
type KeyspaceConfig struct {
	// KeyPattern is a string literal representing the
	// cache key pattern for this keyspace.
	KeyPattern constStr

	// DefaultExpiry specifies the default key expiry for cache items
	// in this keyspace.
	//
	// When set, all write operations set (for new keys)
	// or update (for existing keys) the expiration time.
	//
	// When updating the expiration time Encore always
	// performs the combined operation atomically.
	//
	// If nil, cache items have no expiry date by default.
	//
	// The default behavior can be overridden by passing in
	// an ExpiryFunc or KeepTTL as a WriteOption to a specific operation.
	DefaultExpiry ExpiryFunc
}

// An OpError describes the operation that failed.
type OpError struct {
	Operation string
	RawKey    string
	Err       error
}

func (*OpError) Error() (_ string) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.18.0/runtime/storage/cache/cache.go#L120-L122
	doPanic("encore apps must be run using the encore command")
	return
}

func (*OpError) Unwrap() (_ error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.18.0/runtime/storage/cache/cache.go#L124-L126
	doPanic("encore apps must be run using the encore command")
	return
}

// Miss is the error value reported when a key is missing from the cache.
// It must be checked against with errors.Is.
var Miss = errors.New("cache miss")

// KeyExists is the error reported when a key already exists
// and the requested operation is specified to only apply to
// keys that do not already exist.
// It must be checked against with errors.Is.
var KeyExists = errors.New("key already exists")

// An WriteOption customizes the behavior of a single cache write operation.
type WriteOption interface {
	writeOption() // ensure only our package can implement
}

// ExpiryFunc is a function that reports when a key should expire
// given the current time. It can be used as a WriteOption to customize
// the expiration for that particular operation.
type ExpiryFunc func(now time.Time) time.Time

// option implements WriteOption.
func (ExpiryFunc) writeOption() {}

// ExpireIn returns an ExpiryFunc that expires keys after a constant duration.
func ExpireIn(dur time.Duration) (_ ExpiryFunc) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.18.0/runtime/storage/cache/cache.go#L166-L168
	doPanic("encore apps must be run using the encore command")
	return
}

// ExpireDailyAt returns an ExpiryFunc that expires keys daily at the given time of day in loc.
// ExpireDailyAt panics if loc is nil.
func ExpireDailyAt(hour, minute, second int, loc *time.Location) (_ ExpiryFunc) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.18.0/runtime/storage/cache/cache.go#L172-L182
	doPanic("encore apps must be run using the encore command")
	return
}

// expiryTime is a type for time constants that are also WriteOptions.
type expiryTime time.Time

func (expiryTime) writeOption() {}

var (
	// NeverExpire is a WriteOption indicating the key should never expire.
	NeverExpire = expiryTime(neverExpire)

	// KeepTTL is a WriteOption indicating the key's TTL should be kept the same.
	KeepTTL = expiryTime(keepTTL)
)

var (
	neverExpire = time.Unix(0, 1)
	keepTTL     = time.Unix(0, 2)
)
