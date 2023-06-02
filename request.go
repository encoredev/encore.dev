package encore

import (
	"reflect"
	"time"
)

// APIDesc describes the API endpoint being called.
type APIDesc struct {
	// RequestType specifies the type of the request payload,
	// or nil if the endpoint has no request payload or is Raw.
	RequestType reflect.Type

	// ResponseType specifies the type of the response payload,
	// or nil if the endpoint has no response payload or is Raw.
	ResponseType reflect.Type

	// Raw specifies whether the endpoint is a Raw endpoint.
	Raw bool
}

// Request provides metadata about how and why the currently running code was started.
//
// The current request can be returned by calling CurrentRequest()
type Request struct {
	Type    RequestType // What caused this request to start
	Started time.Time   // What time the trigger occurred

	// Trace contains the trace information for the current request.
	// It is nil if the request is not traced.
	Trace *TraceData

	// APICall specific parameters.
	// These will be empty for operations with a type not APICall
	API        *APIDesc   // Metadata about the API endpoint being called
	Service    string     // Which service is processing this request
	Endpoint   string     // Which API endpoint is being called
	Path       string     // What was the path made to the API server
	PathParams PathParams // If there are path parameters, what are they?

	// PubSubMessage specific parameters.
	// Message contains information about the PubSub message,
	Message *MessageData

	// Payload is the decoded request payload or Pub/Sub message payload,
	// or nil if the API endpoint has no request payload or the endpoint is raw.
	Payload any

	// CronIdempotencyKey contains a unique id for a particular cron job execution
	// if this request was triggered by a Cron Job.
	//
	// It can be used to uniquely identify a particular Cron Job execution event,
	// and also serves as a way to distinguish between Cron Job-triggered requests
	// and other requests.
	//
	// If the request was not triggered by a Cron Job the value is the empty string.
	CronIdempotencyKey string
}

// TraceData describes the trace information for a request.
type TraceData struct {
	TraceID          string
	SpanID           string
	ParentTraceID    string // empty if no parent trace
	ParentSpanID     string // empty if no parent span
	ExtCorrelationID string // empty if no correlation id
}

// MessageData describes the request data for a Pub/Sub message.
type MessageData struct {
	// Service is the name of the service with the subscription.
	Service string

	// Topic is the name of the topic the message was published to.
	Topic string

	// Subscription is the name of the subscription the message was received on.
	Subscription string

	// ID is the unique ID of the message assigned by the messaging service.
	// It is the same value returned by topic.Publish() and is the same
	// across delivery attempts.
	ID string

	// Published is the time the message was first published.
	Published time.Time

	// DeliveryAttempt is a counter for how many times the messages
	// has been attempted to be delivered.
	DeliveryAttempt int
}

// RequestType describes how the currently running code was triggered
type RequestType string

const (
	None          RequestType = "none"           // There was no external trigger which caused this code to run. Most likely it was triggered by a package level init function.
	APICall       RequestType = "api-call"       // The code was triggered via an API call to a service
	PubSubMessage RequestType = "pubsub-message" // The code was triggered by a PubSub subscriber
)

// PathParams contains the path parameters parsed from the request path.
// The ordering of the parameters in the path will be maintained from the URL.
type PathParams []PathParam

// PathParam represents a parsed path parameter.
type PathParam struct {
	Name  string // the name of the path parameter, without leading ':' or '*'.
	Value string // the parsed path parameter value.
}

// Get returns the value of the path parameter with the given name.
// If no such parameter exists it reports "".
func (PathParams) Get(name string) (_ string) {
	// Encore will provide an implementation to this function at runtime, we do not expose
	// the implementation in the API contract as it is an implementation detail, which may change
	// between releases.
	//
	// The current implementation of this function can be found here:
	//    https://github.com/encoredev/encore/blob/v1.20.0/runtime/request.go#L118-L126
	doPanic("encore apps must be run using the encore command")
	return
}
