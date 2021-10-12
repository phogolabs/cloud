package gcp

import (
	"encoding/json"
	"time"

	"github.com/cloudevents/sdk-go/v2/binding"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/cloudevents/sdk-go/v2/protocol/http"
	"github.com/googleapis/google-cloudevents-go/cloud/storage/v1"
	"github.com/phogolabs/schema"
)

var (
	// NewEvent returns a new Event, an optional version can be passed to change the
	// default spec version from 1.0 to the provided version.
	NewEvent = event.New

	// NewMessage returns a binding.Message with header and data.
	// The returned binding.Message *cannot* be read several times. In order to read it more times, buffer it using binding/buffering methods
	NewMessage = http.NewMessage

	// MessageFromEvent converts an event to a message.
	MessageFromEvent = binding.ToMessage

	// MessageWriteRequest write the message to a request.
	MessageWriteRequest = http.WriteRequest
)

type (
	// Event represents the canonical representation of a CloudEvent.
	Event = event.Event
)

// NewStorageEvent creates a new storage event
func NewStorageEvent(name string, data *storage.StorageObjectData) (*Event, error) {
	deref := func(v *string) string {
		if v != nil {
			return *v
		}

		return ""
	}

	var (
		id      = schema.NewUUID().String()
		kind    = "https://www.googleapis.com/" + deref(data.Kind)
		source  = deref(data.SelfLink)
		subject = deref(data.Bucket) + "/" + deref(data.Name)
	)

	event := NewEvent()
	event.SetID(id)
	event.SetTime(time.Now())
	event.SetSubject(subject)
	event.SetSource(source)
	event.SetDataSchema(kind)
	event.SetDataContentType("application/json")

	switch name {
	case "OBJECT_FINALIZE":
		event.SetType("google.storage.object.finalized")
	case "OBJECT_ARCHIVE":
		event.SetType("google.storage.object.archived")
	case "OBJECT_DELETE":
		event.SetType("google.storage.object.deleted")
	case "OBJECT_METADATA_UPDATE":
		event.SetType("google.storage.object.metadata.updated")
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// copy the actual data
	event.DataEncoded = body

	// copy the metadata as well
	for k, v := range data.Metadata {
		if err := event.Context.SetExtension(k, v); err != nil {
			continue
		}
	}

	// validate the event
	if err := event.Validate(); err != nil {
		return nil, err
	}

	return &event, nil
}
