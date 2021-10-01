package gcp

import (
	"context"

	"github.com/cloudevents/sdk-go/protocol/pubsub/v2"
	v2 "github.com/cloudevents/sdk-go/v2"
)

type (
	// Option is the function signature required to be considered an pubsub.Option.
	Option = pubsub.Option

	// Protocol is the pubsub.Protocol
	Protocol = pubsub.Protocol
)

var (
	// WithProjectID sets the project ID for pubsub transport.
	WithProjectID = pubsub.WithProjectID

	// WithTopicID sets the topic ID for pubsub transport.
	WithTopicID = pubsub.WithTopicID
)

// NewClient creates a new pub-sub client
func NewClient(opts ...Option) (v2.Client, error) {
	protocol, err := pubsub.New(context.Background(), opts...)
	if err != nil {
		return nil, err
	}

	return v2.NewClient(protocol)
}
