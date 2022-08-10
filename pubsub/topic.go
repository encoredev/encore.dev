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

// Publish will publish a message to the topic and returns a unique message ID for the message.
//
// This function will not return until the message has been successfully accepted by the topic.
//
// If an error is returned, it is probable that the message failed to be published, however it is possible
// that the message could still be received by subscriptions to the topic.
func (*Topic[T]) Publish(ctx context.Context, msg T) (id string, err error) {
	panic("encore apps must be run using the encore command")
}
