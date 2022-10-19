package et

import (
	"encore.dev/pubsub"
)

// Topic returns a TopicHelper for the given topic.
func Topic[T any](topic *pubsub.Topic[T]) TopicHelpers[T] {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/49a2d858ee8ab00336b162540061e232e9d3f70e/runtime/et/pubsub.go#L8-L10
	panic("encore apps must be run using the encore command")
}

// TopicHelpers provides functions for interacting with the backing topic implementation
// during unit tests. It is designed to help test code that uses the pubsub.Topic
//
// Note all functions on this TopicHelpers are scoped to the current test
// and will only impact and observe state from the current test
type TopicHelpers[T any] interface {
	// PublishedMessages returns a slice of all messages published during this test on this topic.
	PublishedMessages() []T
}
