package cloud

import (
	"context"
	"time"

	"github.com/phogolabs/log"
	"github.com/phogolabs/schema"
)

//go:generate counterfeiter -o ./fake/event_args.go . EventArgs

// EventArgs represents the event arguments
type EventArgs interface {
	Type() string
	Source() string
	Subject() string
	DataSchema() string
	Extensions() Dictionary
	DataContentType() string
}

//go:generate counterfeiter -o ./fake/event_args_converter.go . EventArgsConverter

// EventArgsConverter converts an event arguments
type EventArgsConverter interface {
	// inherits from EventArgs
	EventArgs
	// DataAs converts the event arguments to the given type
	DataAs(interface{}) error
}

//go:generate counterfeiter -o ./fake/event_args_dispatcher.go . EventArgsDispatcher

// EventArgsDispatcher dispatches the events
type EventArgsDispatcher interface {
	// Dispatch dispatches the event args
	Dispatch(context.Context, EventArgs) error
}

var _ EventArgsDispatcher = &EventDispatcher{}

// EventDispatcher represents an event dispatcher
type EventDispatcher struct {
	// EventSender sends the actual event
	EventSender EventSender
}

// Dispatch dispatches a given event
func (d *EventDispatcher) Dispatch(ctx context.Context, args EventArgs) error {
	logger := log.GetContext(ctx)

	logger.Infof("create an outbound event")
	// create the event
	event, err := NewEventArgs(args)
	if err != nil {
		logger.WithError(err).Errorf("create an outbound event failure")
		return err
	}

	// enrich the logger
	logger = logger.WithFields(log.Map{
		"outgoing_event_id":                event.ID(),
		"outgoing_event_type":              event.Type(),
		"outgoing_event_source":            event.Source(),
		"outgoing_event_subject":           event.Subject(),
		"outgoing_event_data_schema":       event.DataSchema(),
		"outgoing_event_data_content_type": event.DataContentType(),
	})

	logger.Info("send an outbound event")
	// send the outbound event to another pub subs
	if err := d.EventSender.Send(ctx, *event); err != nil {
		logger.WithError(err).Error("send an outbound event")
		return err
	}

	return nil
}

// NewEventArgs creates a new event args
func NewEventArgs(args EventArgs) (*Event, error) {
	event := NewEvent()
	event.SetType(args.Type())
	event.SetID(schema.NewUUID().String())
	event.SetSubject(args.Subject())
	event.SetSource(args.Source())
	event.SetDataSchema(args.DataSchema())
	event.SetTime(time.Now())

	// add any extensions at the end
	for k, v := range args.Extensions() {
		event.SetExtension(k, v)
	}

	if err := event.SetData(args.DataContentType(), args); err != nil {
		return nil, err
	}

	return &event, nil
}
