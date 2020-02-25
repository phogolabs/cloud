package cloud

import (
	"context"

	cloudevents "github.com/cloudevents/sdk-go"
	cloudtypes "github.com/cloudevents/sdk-go/pkg/cloudevents/types"
)

type (
	// Event represents the canonical representation of a CloudEvent.
	Event = cloudevents.Event

	// EventContext is conical interface for a CloudEvents Context.
	EventContext = cloudevents.EventContext

	// EventContextV1 represents the non-data attributes of a CloudEvents v1.0
	// event.
	EventContextV1 = cloudevents.EventContextV1

	// EventContextV01 holds standard metadata about an event. See
	// https://github.com/cloudevents/spec/blob/v0.1/spec.md#context-attributes for
	// details on these fields.
	EventContextV01 = cloudevents.EventContextV01

	// EventContextV02 represents the non-data attributes of a CloudEvents v0.2
	// event.
	EventContextV02 = cloudevents.EventContextV02

	// EventContextV03 represents the non-data attributes of a CloudEvents v0.3
	// event.
	EventContextV03 = cloudevents.EventContextV03

	// Timestamp wraps time.Time to normalize the time layout to RFC3339. It is
	// intended to enforce compliance with the CloudEvents spec for their
	// definition of Timestamp. Custom marshal methods are implemented to ensure
	// the outbound Timestamp is a string in the RFC3339 layout.
	Timestamp = cloudevents.Timestamp

	// URIRef is a wrapper to url.URL. It is intended to enforce compliance with
	// the CloudEvents spec for their definition of URI-Reference. Custom
	// marshal methods are implemented to ensure the outbound URIRef object is
	// is a flat string.
	URIRef = cloudtypes.URIRef
)

var (
	// NewEvent returns a new Event, an optional version can be passed to change the
	// default spec version from 0.2 to the provided version.
	NewEvent = cloudevents.NewEvent

	// ParseTimestamp attempts to parse the given time assuming RFC3339 layout
	ParseTimestamp = cloudevents.ParseTimestamp

	// ParseURIRef attempts to parse the given string as a URI-Reference.
	ParseURIRef = cloudevents.ParseURIRef
)

//go:generate counterfeiter -fake-name EventHandler -o ./fake/event_handler.go . EventHandler

// EventHandler handles cloud events
type EventHandler interface {
	// Handle handles cloud events
	Handle(ctx context.Context, event *Event) error
}

var _ EventHandler = &CompositeEventTypeHandler{}

// CompositeEventTypeHandler represents a composite event handler for given type
type CompositeEventTypeHandler map[string]EventHandler

// Handle handles cloud events
func (kv *CompositeEventTypeHandler) Handle(ctx context.Context, event *Event) error {
	if handler, ok := (*kv)[event.Type()]; ok {
		return handler.Handle(ctx, event)
	}

	return nil
}

var _ EventHandler = &CompositeEventSourceHandler{}

// CompositeEventSourceHandler represents a composite event handler for given source
type CompositeEventSourceHandler map[string]EventHandler

// Handle handles cloud events
func (kv *CompositeEventSourceHandler) Handle(ctx context.Context, event *Event) error {
	if handler, ok := (*kv)[event.Source()]; ok {
		return handler.Handle(ctx, event)
	}

	return nil
}

var _ EventHandler = &CompositeEventHandler{}

// CompositeEventHandler represent a multiplex event handler
type CompositeEventHandler []EventHandler

// Handle handles cloud events
func (items *CompositeEventHandler) Handle(ctx context.Context, event *Event) error {
	for _, handler := range *items {
		if err := handler.Handle(ctx, event); err != nil {
			return err
		}
	}

	return nil
}

//go:generate counterfeiter -fake-name EventSender -o ./fake/event_sender.go . EventSender

// EventSender sends the event
type EventSender interface {
	// Send sends the event
	Send(ctx context.Context, event *Event) error
}
