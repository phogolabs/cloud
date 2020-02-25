package cloud

import (
	cloudevents "github.com/cloudevents/sdk-go"
	cloudcontext "github.com/cloudevents/sdk-go/pkg/cloudevents/context"
)

// WithTarget returns back a new context with the given target. Target is intended to be transport dependent.
// For http transport, `target` should be a full URL and will be injected into the outbound http request.
// func WithTarget(ctx context.Context, target string) context.Context {

// TargetFrom looks in the given context and returns `target` as a parsed url if found and valid, otherwise nil.
// func TargetFrom(ctx context.Context) *url.URL {

var (
	// ContextWithTarget returns back a new context with the given target. Target is intended to be transport dependent.
	// For http transport, `target` should be a full URL and will be injected into the outbound http request.
	ContextWithTarget = cloudevents.ContextWithTarget

	// TargetFromContext looks in the given context and returns `target` as a parsed url if found and valid, otherwise nil.
	TargetFromContext = cloudevents.TargetFromContext

	// ContextWithEncoding returns back a new context with the given encoding. Encoding is intended to be transport dependent.
	// For http transport, `encoding` should be one of [binary, structured] and will be used to override the outbound
	// codec encoding setting. If the transport does not understand the encoding, it will be ignored.
	ContextWithEncoding = cloudevents.ContextWithEncoding

	// EncodingFromContext looks in the given context and returns `target` as a parsed url if found and valid, otherwise nil.
	EncodingFromContext = cloudevents.EncodingFromContext

	// ContextWithTopic returns back a new context with the given topic. Topic is intended to be transport dependent.
	// For pubsub transport, `topic` should be a Pub/Sub Topic ID.
	ContextWithTopic = cloudcontext.WithTopic

	// TopicFromContext looks in the given context and returns `topic` as a string if found and valid, otherwise "".
	TopicFromContext = cloudcontext.TopicFrom
)
