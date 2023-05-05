package pubsub

import (
	"context"
)

// Topic presents a flow of events of type T from any number of publishers to
// any number of subscribers.
//
// Each subscription will receive a copy of each message published to the topic.
//
// See NewTopic for more information on how to declare a Topic.
type Topic[T any] struct {
	_ int // for godoc to show unexported fields
}

// TopicMeta contains metadata about a topic.
// The fields should not be modified by the caller.
// Additional fields may be added in the future.
type TopicMeta struct {
	// Name is the name of the topic, as provided in the constructor to NewTopic.
	Name string
	// Config is the topic's configuration.
	Config TopicConfig
}

// Meta returns metadata about the topic.
func (*Topic[T]) Meta() (_ TopicMeta) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/38dbb67953561748253891e3e50181bdf095d6e5/runtime/pubsub/topic.go#L82-L87
	doPanic("encore apps must be run using the encore command")
	return
}

// Publish will publish a message to the topic and returns a unique message ID for the message.
//
// This function will not return until the message has been successfully accepted by the topic.
//
// If an error is returned, it is probable that the message failed to be published, however it is possible
// that the message could still be received by subscriptions to the topic.
func (*Topic[T]) Publish(ctx context.Context, msg T) (id string, err error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/38dbb67953561748253891e3e50181bdf095d6e5/runtime/pubsub/topic.go#L95-L151
	doPanic("encore apps must be run using the encore command")
	return
}
