package cloud

import (
	"context"

	"github.com/cloudevents/sdk-go/protocol/pubsub/v2"
)

type PubsubConfig struct {
	ProjectID string
	TopicID   string
}

// NewPubsub creates a new pub-sub client
func NewClientPubsub(config *PubsubConfig) (Client, error) {
	protocol, err := pubsub.New(context.Background(),
		pubsub.WithProjectID(config.ProjectID),
		pubsub.WithTopicID(config.TopicID),
	)

	if err != nil {
		return nil, err
	}

	return NewClient(protocol)
}
