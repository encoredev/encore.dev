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

// NewTopic is used to declare a Topic. Encore will use static
// analysis to identify Topics and automatically provision them
// for you.
//
// A call to NewTopic can only be made when declaring a package level variable. Any
// calls to this function made outside a package level variable declaration will result
// in a compiler error.
//
// The topic name must be unique within an Encore application. Topic names must be defined
// in kebab-case (lowercase alphanumerics and hyphen seperated). The topic name must start with a letter
// and end with either a letter or number. It cannot be longer than 63 characters. Once created and deployed never
// change the topic name. When refactoring the topic name must stay the same.
// This allows for messages already on the topic to continue to be received after the refactored
// code is deployed.
//
// Example:
//
//     import "encore.dev/pubsub"
//
//     type MyEvent struct {
//       Foo string
//     }
//
//     var MyTopic = pubsub.NewTopic[*MyEvent]("my-topic", pubsub.TopicConfig{
//       DeliveryGuarantee: pubsub.AtLeastOnce,
//     })
//
//    //encore:api public
//    func DoFoo(ctx context.Context) error {
//      msgID, err := MyTopic.Publish(ctx, &MyEvent{Foo: "bar"})
//      if err != nil { return err }
//      rlog.Info("foo published", "message_id", msgID)
//      return nil
//    }
func NewTopic[T any](name string, cfg TopicConfig) *Topic[T] {
	panic("encore apps must be run using the encore command")
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
