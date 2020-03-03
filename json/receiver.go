package json

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"

	cloudevents "github.com/cloudevents/sdk-go"
	uuid "github.com/google/uuid"
	log "github.com/phogolabs/log"
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
	// Regexp is the parameter regexp
	Regexp = regexp.MustCompile("{[a-z_]+}")

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

// EventReceiverConfig represent the receiver's configuration
type EventReceiverConfig struct {
	EventName    string
	EventSubject string
	EventSource  string
}

// EventReceiver receives a pub sub message
type EventReceiver struct {
	Config  *EventReceiverConfig
	Handler EventHandler
}

// Receive receives the pubsub message
func (r *EventReceiver) Receive(ctx context.Context, message RawMessage) error {
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
		"event_provider": "providers/json",
	})

	if len(message) > 0 {
		subject, err := r.subject(message)

		if err != nil {
			logger.WithError(err).Error("event subject getting failed")
			return err
		}

		event.SetSubject(subject)
	}

	logger.Info("handling event start")

	if err := r.Handler.Handle(ctx, &event); err != nil {
		logger.WithError(err).Info("event handling fail")
		return err
	}

	logger.Info("handling event success")
	return nil
}

func (r *EventReceiver) subject(message RawMessage) (string, error) {
	subject := r.Config.EventSubject

	if subject == "" {
		return subject, nil
	}

	var (
		data = make(map[string]interface{})
		keys = make(map[string]string)
	)

	if err := json.Unmarshal(message, &data); err != nil {
		return "", err
	}

	params := Regexp.FindAllString(subject, -1)

	if len(params) == 0 {
		params = append(params, subject)
	}

	for _, param := range params {
		key := param
		key = strings.TrimPrefix(key, "{")
		key = strings.TrimSuffix(key, "}")

		if value, ok := data[key]; ok {
			keys[param] = fmt.Sprintf("%v", value)
		}
	}

	for key, value := range keys {
		subject = strings.Replace(subject, key, value, -1)
	}

	return subject, nil
}
