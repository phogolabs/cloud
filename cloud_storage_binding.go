package cloud

import (
	"bytes"
	"encoding/json"

	"github.com/AlekSi/pointer"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/cloudevents/sdk-go/v2/types"
	"github.com/golang/protobuf/jsonpb"
)

// StorageFormat represents a storage event format
type StorageFormat struct{}

func (StorageFormat) MediaType() string {
	return event.ApplicationJSON
}

// This method is wrong, but I don't need marshalling, just unmarshalling.
func (StorageFormat) Marshal(e *event.Event) ([]byte, error) {
	return json.Marshal(e)
}

// Unmarshal unmarshals the event for given payload
func (StorageFormat) Unmarshal(data []byte, e *event.Event) error {
	payload := &PubsubEvent{}

	if err := jsonpb.Unmarshal(bytes.NewBuffer(data), payload); err != nil {
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
		DataContentType: pointer.ToStringOrNil(event.ApplicationJSON),
		Type:            payload.Message.Attributes["eventType"],
		Source:          *types.ParseURIRef("https://" + bucketID),
		Time:            timestamp,
	}

	return e.Validate()
}
