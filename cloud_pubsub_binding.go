package cloud

import (
	"encoding/json"
	"fmt"

	"github.com/AlekSi/pointer"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/cloudevents/sdk-go/v2/types"
)

type PubsubFormat struct{}

func (PubsubFormat) MediaType() string {
	return event.ApplicationJSON
}

// This method is wrong, but I don't need marshalling, just unmarshalling.
func (PubsubFormat) Marshal(e *event.Event) ([]byte, error) {
	return json.Marshal(e)
}

func (PubsubFormat) Unmarshal(data []byte, e *event.Event) error {
	payload := &PubsubEvent{}

	if err := json.Unmarshal(data, payload); err != nil {
		return err
	}

	timestamp, err := types.ParseTimestamp(payload.Message.Attributes["ce-time"])
	if err != nil {
		return err
	}

	source := types.ParseURIRef(payload.Message.Attributes["ce-source"])
	if source == nil {
		return fmt.Errorf("ce-source is not provided")
	}

	e.DataEncoded = payload.Message.Data
	e.Context = &event.EventContextV1{
		ID:              payload.Message.Attributes["ce-id"],
		Subject:         pointer.ToStringOrNil(payload.Message.Attributes["ce-subject"]),
		DataContentType: pointer.ToStringOrNil(payload.Message.Attributes["ce-datacontenttype"]),
		Type:            payload.Message.Attributes["ce-type"],
		Source:          *source,
		Time:            timestamp,
	}

	return e.Validate()
}
