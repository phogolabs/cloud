package cloud

import (
	"context"

	cloudevents "github.com/cloudevents/sdk-go"
	cloudtypes "github.com/cloudevents/sdk-go/pkg/cloudevents/types"
	log "github.com/phogolabs/log"
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

const (
	// ApplicationXML is XML content type
	ApplicationXML = cloudevents.ApplicationXML
	// ApplicationJSON is JSON content type
	ApplicationJSON = cloudevents.ApplicationJSON
	// ApplicationCloudEventsJSON is Cloud Event JSON
	ApplicationCloudEventsJSON = cloudevents.ApplicationCloudEventsJSON
	// ApplicationCloudEventsBatchJSON is Cloud Event Batch JSON
	ApplicationCloudEventsBatchJSON = cloudevents.ApplicationCloudEventsBatchJSON
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
	logger := log.GetContext(ctx)

	logger = logger.WithFields(log.Map{
		"event_id":      event.ID(),
		"event_type":    event.Type(),
		"event_source":  event.Source(),
		"event_handler": "event_type",
	})

	if handler, ok := (*kv)[event.Type()]; ok {
		logger.Info("event handling start")

		if err := handler.Handle(ctx, event); err != nil {
			logger.WithError(err).Error("event handling fail")
			return err
		}

		logger.Info("event handling success")
		return nil
	}

	logger.Info("event handling not started")
	return nil
}

var _ EventHandler = &CompositeEventSourceHandler{}

// CompositeEventSourceHandler represents a composite event handler for given source
type CompositeEventSourceHandler map[string]EventHandler

// Handle handles cloud events
func (kv *CompositeEventSourceHandler) Handle(ctx context.Context, event *Event) error {
	logger := log.GetContext(ctx)

	logger = logger.WithFields(log.Map{
		"event_id":      event.ID(),
		"event_type":    event.Type(),
		"event_source":  event.Source(),
		"event_handler": "event_source",
	})

	if handler, ok := (*kv)[event.Source()]; ok {
		logger.Info("event handling start")

		if err := handler.Handle(ctx, event); err != nil {
			logger.WithError(err).Error("event handling fail")
			return err
		}

		logger.Info("event handling success")
		return nil
	}

	logger.Info("event handling not started")
	return nil
}

var _ EventHandler = &CompositeEventHandler{}

// CompositeEventHandler represent a multiplex event handler
type CompositeEventHandler []EventHandler

// Handle handles cloud events
func (items *CompositeEventHandler) Handle(ctx context.Context, event *Event) error {
	logger := log.GetContext(ctx)

	logger = logger.WithFields(log.Map{
		"event_id":      event.ID(),
		"event_type":    event.Type(),
		"event_source":  event.Source(),
		"event_handler": "event_composite",
	})

	for _, handler := range *items {
		logger.Info("event handling start")

		if err := handler.Handle(ctx, event); err != nil {
			logger.WithError(err).Error("event handling fail")
			return err
		}

		logger.Info("event handling success")
	}

	logger.Info("event handling not started")
	return nil
}

//go:generate counterfeiter -fake-name EventSender -o ./fake/event_sender.go . EventSender

// EventSender sends the event
type EventSender interface {
	// Send sends the event
	Send(ctx context.Context, event *Event) error
}

var _ EventHandler = &EventDispatcher{}

// EventDispatcher dispatches the event immediately
type EventDispatcher struct {
	Sender EventSender
}

// Handle handles cloud events
func (h *EventDispatcher) Handle(ctx context.Context, event *Event) error {
	logger := log.GetContext(ctx)

	logger = logger.WithFields(log.Map{
		"event_id":      event.ID(),
		"event_type":    event.Type(),
		"event_source":  event.Source(),
		"event_handler": "event_dispatch",
	})

	logger.Info("event dispatching start")

	if err := h.Sender.Send(ctx, event); err != nil {
		logger.WithError(err).Error("event dispatching fail")
		return err
	}

	logger.Info("event dispatching success")
	return nil
}
