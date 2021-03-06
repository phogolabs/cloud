package gcp

import (
	"context"

	"github.com/cloudevents/sdk-go/protocol/pubsub/v2"
	v2 "github.com/cloudevents/sdk-go/v2"
)

type (
	// PubsubOption is the function signature required to be considered an pubsub.Option.
	PubsubOption = pubsub.Option

	// PubsubProtocol is the pubsub.Protocol
	PubsubProtocol = pubsub.Protocol
)

var (
	// PubsubWithProjectID sets the project ID for pubsub transport.
	PubsubWithProjectID = pubsub.WithProjectID

	// PubsubWithTopicID sets the topic ID for pubsub transport.
	PubsubWithTopicID = pubsub.WithTopicID
)

// NewPubsub creates a new pub-sub client
func NewClientPubsub(opts ...PubsubOption) (v2.Client, error) {
	protocol, err := pubsub.New(context.Background(), opts...)
	if err != nil {
		return nil, err
	}

	return v2.NewClient(protocol)
}
