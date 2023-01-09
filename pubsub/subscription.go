package pubsub

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
// in kebab-case (lowercase alphanumerics and hyphen seperated). The subscription name must start with a letter
// and end with either a letter or number. It cannot be longer than 63 characters.
//
// Once created and deployed never change the subscription name, or the topic name otherwise messages will be lost which
// could be in flight.
//
// Example:
//
//	import "encore.dev/pubsub"
//
//	type MyEvent struct {
//	  Foo string
//	}
//
//	var MyTopic = pubsub.NewTopic[*MyEvent]("my-topic", pubsub.TopicConfig{
//	  DeliveryGuarantee: pubsub.AtLeastOnce,
//	})
//
//	var Subscription = pubsub.NewSubscription(MyTopic, "my-subscription", pubsub.SubscriptionConfig[*MyEvent]{
//	  Handler:     HandleEvent,
//	  RetryPolicy: &pubsub.RetryPolicy { MaxRetries: 10 },
//	})
//
//	func HandleEvent(ctx context.Context, event *MyEvent) error {
//	  rlog.Info("received foo")
//	  return nil
//	}
func NewSubscription[T any](topic *Topic[T], name string, subscriptionCfg SubscriptionConfig[T]) (_ *Subscription[T]) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/4d212a6471c0a6f5e7df1114b5238c8084d76c07/runtime/pubsub/subscription.go#L56-L211
	doPanic("encore apps must be run using the encore command")
	return
}
