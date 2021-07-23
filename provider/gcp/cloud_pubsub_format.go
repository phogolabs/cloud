package gcp

import (
	"encoding/json"

	"github.com/cloudevents/sdk-go/v2/binding/format"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/cloudevents/sdk-go/v2/types"
	"github.com/phogolabs/log"
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
		log.WithError(err).Error("unmarshal payload failure")
		return err
	}

	timestamp, err := types.ParseTimestamp(payload.Message.Attributes["ce-time"])
	if err != nil {
		log.WithError(err).Error("parse ce-time failure")
		return err
	}

	e.Context = &event.EventContextV1{}
	e.DataEncoded = payload.Message.Data

	if err := e.Context.SetID(payload.Message.Attributes["ce-id"]); err != nil {
		log.WithError(err).Error("set ce-id failure")
		return err
	}

	if err := e.Context.SetType(payload.Message.Attributes["ce-type"]); err != nil {
		log.WithError(err).Error("set ce-type failure")
		return err
	}

	if err := e.Context.SetSubject(payload.Message.Attributes["ce-subject"]); err != nil {
		log.WithError(err).Error("set ce-subject failure")
		return err
	}

	if err := e.Context.SetSource(payload.Message.Attributes["ce-source"]); err != nil {
		log.WithError(err).Error("set ce-source failure")
		return err
	}

	if err := e.Context.SetTime(timestamp.Time); err != nil {
		log.WithError(err).Error("set ce-time failure")
		return err
	}

	if err := e.Context.SetDataSchema(payload.Message.Attributes["ce-dataschema"]); err != nil {
		log.WithError(err).Error("set ce-dataschema failure")
		return err
	}

	if err := e.Context.SetDataContentType(payload.Message.Attributes["Content-Type"]); err != nil {
		log.WithError(err).Error("set content-type failure")
		return err
	}

	return nil
}
