package gcp

import (
	"encoding/json"

	"github.com/cloudevents/sdk-go/v2/binding/format"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/cloudevents/sdk-go/v2/types"
	"google.golang.org/protobuf/encoding/protojson"
)

func init() {
	format.Add(&PubsubFormat{})
}

// PubsubFormat represents a pub-sub push format
type PubsubFormat struct{}

// MediaType returns the media type
func (PubsubFormat) MediaType() string {
	return "application/google.pubsub.message+json"
}

// This method is wrong, but I don't need marshalling, just unmarshalling.
func (PubsubFormat) Marshal(e *event.Event) ([]byte, error) {
	return json.Marshal(e)
}

// Unmarshal unmarshals the event for a given data
func (PubsubFormat) Unmarshal(data []byte, e *event.Event) error {
	payload := &PubsubEvent{}

	decoder := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}

	if err := decoder.Unmarshal(data, payload); err != nil {
		return err
	}

	timestamp, err := types.ParseTimestamp(payload.Message.Attributes["ce-time"])
	if err != nil {
		return err
	}

	e.Context = &event.EventContextV1{}
	e.DataEncoded = payload.Message.Data

	if err := e.Context.SetID(payload.Message.Attributes["ce-id"]); err != nil {
		return err
	}

	if err := e.Context.SetType(payload.Message.Attributes["ce-type"]); err != nil {
		return err
	}

	if err := e.Context.SetSubject(payload.Message.Attributes["ce-subject"]); err != nil {
		return err
	}

	if err := e.Context.SetSource(payload.Message.Attributes["ce-source"]); err != nil {
		return err
	}

	if err := e.Context.SetTime(timestamp.Time); err != nil {
		return err
	}

	if err := e.Context.SetDataSchema(payload.Message.Attributes["ce-dataschema"]); err != nil {
		return err
	}

	if err := e.Context.SetDataContentType(payload.Message.Attributes["ce-datacontenttype"]); err != nil {
		return err
	}

	return nil
}
