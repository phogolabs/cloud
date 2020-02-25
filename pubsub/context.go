package pubsub

import (
	cloudcontext "github.com/cloudevents/sdk-go/pkg/cloudevents/context"
	pubsubcontext "github.com/cloudevents/sdk-go/pkg/cloudevents/transport/pubsub/context"
)

type (
	// TransportContext allows a Receiver to understand the context of a request.
	TransportContext = pubsubcontext.TransportContext
)

var (
	// ContextWithTopic returns back a new context with the given topic. Topic is intended to be transport dependent.
	// For pubsub transport, `topic` should be a Pub/Sub Topic ID.
	ContextWithTopic = cloudcontext.WithTopic

	// TopicFromContext looks in the given context and returns `topic` as a string if found and valid, otherwise "".
	TopicFromContext = cloudcontext.TopicFrom

	// ContextWithTransport return a context with the given TransportContext into the provided context object.
	ContextWithTransport = pubsubcontext.WithTransportContext

	// TransportFromContext pulls a TransportContext out of a context. Always
	// returns a non-nil object.
	TransportFromContext = pubsubcontext.TransportContextFrom
)
