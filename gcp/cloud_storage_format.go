package gcp

import (
	"encoding/json"

	"github.com/AlekSi/pointer"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/cloudevents/sdk-go/v2/types"
	"google.golang.org/protobuf/encoding/protojson"
)

// StorageFormat represents a storage event format
type StorageFormat struct{}

// MediaType returns the media type
func (StorageFormat) MediaType() string {
	return "application/grpc+json;message=phogolabs.cloud.gcp.StorageObject"
}

// This method is wrong, but I don't need marshalling, just unmarshalling.
func (StorageFormat) Marshal(e *event.Event) ([]byte, error) {
	return json.Marshal(e)
}

// Unmarshal unmarshals the event for given payload
func (StorageFormat) Unmarshal(data []byte, e *event.Event) error {
	payload := &PubsubEvent{}

	decoder := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}

	if err := decoder.Unmarshal(data, payload); err != nil {
		return err
	}

	timestamp, err := types.ParseTimestamp(payload.Message.Attributes["eventTime"])
	if err != nil {
		return err
	}

	bucketID := payload.Message.Attributes["bucketId"]
	objectID := payload.Message.Attributes["objectId"]

	e.DataEncoded = payload.Message.Data
	e.Context = &event.EventContextV1{
		ID:              payload.Message.Attributes["objectGeneration"],
		Subject:         pointer.ToStringOrNil(objectID),
		DataContentType: pointer.ToStringOrNil("application/protobuf"),
		Type:            payload.Message.Attributes["eventType"],
		Source:          *types.ParseURIRef("https://" + bucketID),
		Time:            timestamp,
	}

	return e.Validate()
}
