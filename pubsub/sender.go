package pubsub

import (
	"context"

	pubsub "cloud.google.com/go/pubsub"
	pubsubevent "github.com/cloudevents/sdk-go/pkg/cloudevents/transport/pubsub"
	option "google.golang.org/api/option"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type (
	// Client is a Google Pub/Sub client scoped to a single project.
	//
	// Clients should be reused rather than being created as needed.
	// A Client may be shared by multiple goroutines.
	Client = pubsub.Client

	// ClientOption is an option for a Google API client.
	ClientOption = option.ClientOption

	// Topic is a reference to a PubSub topic.
	//
	// The methods of Topic are safe for use by multiple goroutines.
	Topic = pubsub.Topic
)

var (
	// NewClient creates a new PubSub client.
	NewClient = pubsub.NewClient
)

// Sender sends the event to pub sub topic
type Sender struct {
	// TopicID destination
	TopicID string
	// Codec represents the code
	Codec *pubsubevent.Codec
	// Client that sends the message
	Client *pubsub.Client
}

// Send sends the event
func (s *Sender) Send(ctx context.Context, event *Event) error {
	if err := event.Validate(); err != nil {
		err = status.Error(codes.InvalidArgument, err.Error())
		return err
	}

	if s.Codec == nil {
		s.Codec = &pubsubevent.Codec{
			Encoding: pubsubevent.StructuredV1,
		}
	}

	payload, err := s.Codec.Encode(ctx, *event)
	if err != nil {
		err = status.Error(codes.Internal, err.Error())
		return err
	}

	name := TopicFromContext(ctx)

	if name == "" {
		name = s.TopicID
	}

	topic := s.Client.Topic(name)

	if message, ok := payload.(*pubsubevent.Message); ok {
		response := topic.Publish(ctx, &pubsub.Message{
			Attributes: message.Attributes,
			Data:       message.Data,
		})

		if _, err := response.Get(ctx); err != nil {
			err = status.Error(codes.Internal, err.Error())
			return err
		}
	}

	return nil
}