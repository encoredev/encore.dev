package pubsub

import "context"

// TopicPerms is the type constraint for all permission-declaring
// interfaces that can be used with TopicRef.
type TopicPerms[T any] interface {
	Meta() TopicMeta
}

// Publisher is the interface for publishing messages to a topic.
// It can be used in conjunction with [TopicRef] to declare
// a reference that can publish messages to the topic.
//
// For example:
//
//	var MyTopic = pubsub.NewTopic[Msg](...)
//	var ref = pubsub.TopicRef[pubsub.Publisher[Msg]](MyTopic)
//
// The ref object can then be used to publish messages and can be
// passed around freely within the service, without being subject
// to Encore's static analysis restrictions that apply to MyTopic.
type Publisher[T any] interface {
	// Publish publishes a message to the topic.
	Publish(ctx context.Context, msg T) (id string, err error)

	// Meta returns metadata about the topic.
	Meta() TopicMeta
}

// TopicRef returns an interface reference to a topic,
// that can be freely passed around within a service
// without being subject to Encore's typical static analysis
// restrictions that normally apply to *Topic objects.
//
// This works because using TopicRef effectively declares
// which operations you want to be able to perform since the
// type argument P must be a permission-declaring interface (TopicPerms[T]).
//
// The returned reference is scoped down to those permissions.
//
// For example:
//
//	var MyTopic = pubsub.NewTopic[Msg](...)
//	var ref = pubsub.TopicRef[pubsub.Publisher[Msg]](MyTopic)
//	// ref.Publish(...) can now be used to publish messages to MyTopic.
func TopicRef[P TopicPerms[T], T any](topic *Topic[T]) (_ P) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.41.9/runtimes/go/pubsub/refs.go#L47-L49
	doPanic("encore apps must be run using the encore command")
	return
}
