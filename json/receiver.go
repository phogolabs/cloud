package json

import (
	"context"
	"fmt"
	"time"

	cloudevents "github.com/cloudevents/sdk-go"
	uuid "github.com/google/uuid"
)

type (
	// Event represents the canonical representation of a CloudEvent.
	Event = cloudevents.Event

	// Payload is a key value map
	Payload = map[string]interface{}
)

var (
	// NewEvent returns a new Event, an optional version can be passed to change the
	// default spec version from 0.2 to the provided version.
	NewEvent = cloudevents.NewEvent
)

// EventHandler handles cloud events
type EventHandler interface {
	// Handle handles cloud events
	Handle(ctx context.Context, event *Event) error
}

// ReceiverConfig represent the receiver's configuration
type ReceiverConfig struct {
	EventName    string
	EventSubject string
	EventSource  string
}

// Receiver receives a pub sub message
type Receiver struct {
	Config  *ReceiverConfig
	Handler EventHandler
}

// Receive receives the pubsub message
func (r *Receiver) Receive(ctx context.Context, payload Payload) error {
	event := NewEvent()
	event.SetID(uuid.New().String())
	event.SetType(r.Config.EventName)
	event.SetSource(r.Config.EventSource)
	event.SetDataContentType(cloudevents.ApplicationJSON)
	event.SetData(payload)
	event.SetTime(time.Now())

	if subject, ok := payload[r.Config.EventSubject]; ok {
		event.SetSubject(fmt.Sprintf("%v", subject))
	}

	return r.Handler.Handle(ctx, &event)
}
