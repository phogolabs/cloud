package json

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	cloudevents "github.com/cloudevents/sdk-go"
	uuid "github.com/google/uuid"
	"github.com/phogolabs/log"
)

type (
	// Event represents the canonical representation of a CloudEvent.
	Event = cloudevents.Event

	// RawMessage is a raw encoded JSON value.
	// It implements Marshaler and Unmarshaler and can
	// be used to delay JSON decoding or precompute a JSON encoding.
	RawMessage = json.RawMessage
)

var (
	// NewEvent returns a new Event, an optional version can be passed to change the
	// default spec version from 0.2 to the provided version.
	NewEvent = cloudevents.NewEvent

	// Marshal returns m as the JSON encoding of m.
	Marshal = json.Marshal
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
func (r *Receiver) Receive(ctx context.Context, message RawMessage) error {
	logger := log.GetContext(ctx)

	event := NewEvent()
	event.SetID(uuid.New().String())
	event.SetType(r.Config.EventName)
	event.SetSource(r.Config.EventSource)
	event.SetDataContentType(cloudevents.ApplicationJSON)
	event.SetData(message)
	event.SetTime(time.Now())

	logger = logger.WithFields(log.Map{
		"event_id":       event.ID(),
		"event_type":     event.Type(),
		"event_source":   event.Source(),
		"event_receiver": "json",
	})

	if len(message) > 0 {
		if subject := r.Config.EventSubject; subject != "" {
			kv := make(map[string]interface{})

			if err := json.Unmarshal(message, &kv); err != nil {
				logger.WithError(err).Error("event decoding fail")
				return err
			}

			if subject, ok := kv[r.Config.EventSubject]; ok {
				event.SetSubject(fmt.Sprintf("%v", subject))
			}
		}
	}

	logger.Info("handling event start")

	if err := r.Handler.Handle(ctx, &event); err != nil {
		logger.WithError(err).Info("event handling fail")
		return err
	}

	logger.Info("handling event success")
	return nil
}
