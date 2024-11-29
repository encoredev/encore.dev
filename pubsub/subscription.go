package pubsub

import (
	"context"
)

// Subscription represents a subscription to a Topic.
type Subscription[T any] struct {
	_ int // for godoc to show unexported fields
}

// NewSubscription is used to declare a Subscription to a topic. The passed in handler will be called
// for each message published to the topic.
//
// A call to NewSubscription can only be made when declaring a package level variable. Any
// calls to this function made outside a package level variable declaration will result
// in a compiler error.
//
// The subscription name must be unique for that topic. Subscription names must be defined
// in kebab-case (lowercase alphanumerics and hyphen separated). The subscription name must start with a letter
// and end with either a letter or number. It cannot be longer than 63 characters.
//
// Once created and deployed never change the subscription name, or the topic name otherwise messages will be lost which
// could be in flight.
//
// Example:
//
//		import "encore.dev/pubsub"
//
//		type MyEvent struct {
//		  Foo string
//		}
//
//		var MyTopic = pubsub.NewTopic[*MyEvent]("my-topic", pubsub.TopicConfig{
//		  DeliveryGuarantee: pubsub.AtLeastOnce,
//		})
//
//		var Subscription = pubsub.NewSubscription(MyTopic, "my-subscription", pubsub.SubscriptionConfig[*MyEvent]{
//		  Handler:     HandleEvent,
//		  RetryPolicy: &pubsub.RetryPolicy{MaxRetries: 10},
//	      MaxConcurrency: 5,
//		})
//
//		func HandleEvent(ctx context.Context, event *MyEvent) error {
//		  rlog.Info("received foo")
//		  return nil
//		}
func NewSubscription[T any](topic *Topic[T], name string, cfg SubscriptionConfig[T]) (_ *Subscription[T]) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.44.6/runtimes/go/pubsub/subscription.go#L63-L247
	doPanic("encore apps must be run using the encore command")
	return
}

// SubscriptionMeta contains metadata about a subscription.
// The fields should not be modified by the caller.
// Additional fields may be added in the future.
type SubscriptionMeta[T any] struct {
	// Name is the name of the subscription, as provided in the constructor to NewSubscription.
	Name string

	// Config is the subscriptions's configuration.
	Config SubscriptionConfig[T]

	// Topic provides metadata about the topic it subscribes to.
	Topic TopicMeta
}

// Meta returns metadata about the topic.
func (*Subscription[T]) Meta() (_ SubscriptionMeta[T]) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.44.6/runtimes/go/pubsub/subscription.go#L264-L270
	doPanic("encore apps must be run using the encore command")
	return
}

// Config returns the subscription's configuration.
// It must not be modified by the caller.
func (*Subscription[T]) Config() (_ SubscriptionConfig[T]) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.44.6/runtimes/go/pubsub/subscription.go#L274-L276
	doPanic("encore apps must be run using the encore command")
	return
}

// MethodHandler is used to define a subscription Handler that references a service struct method.
//
// Example Usage:
//
//	//encore:service
//	type Service struct {}
//
//	func (s *Service) Method(ctx context.Context, msg *Event) error { /* ... */ }
//
//	var _ = pubsub.NewSubscription(Topic, "subscription-name", pubsub.SubscriptionConfig[*Event]{
//		Handler: pubsub.MethodHandler((*MyService).MyMethod),
//		// ...
//	})
func MethodHandler[T, SvcStruct any](handler func(s SvcStruct, ctx context.Context, msg T) error) (_ func(ctx context.Context, msg T) error) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.44.6/runtimes/go/pubsub/subscription.go#L320-L327
	doPanic("encore apps must be run using the encore command")
	return
}
