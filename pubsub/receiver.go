package pubsub

import (
	"context"
	"encoding/base64"
	"strings"

	pubsub "cloud.google.com/go/pubsub"
	cloudevents "github.com/cloudevents/sdk-go"
	pubsubevent "github.com/cloudevents/sdk-go/pkg/cloudevents/transport/pubsub"
	ptypes "github.com/golang/protobuf/ptypes"
	empty "github.com/golang/protobuf/ptypes/empty"
	log "github.com/phogolabs/log"
	plex "github.com/phogolabs/plex"
	v1 "google.golang.org/genproto/googleapis/pubsub/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type (
	// Event represents the canonical representation of a CloudEvent.
	Event = cloudevents.Event

	// Message that is published by publishers and consumed by subscribers. The
	// message must contain either a non-empty data field or at least one attribute.
	// Note that client libraries represent this object differently
	// depending on the language. See the corresponding
	// <a href="https://cloud.google.com/pubsub/docs/reference/libraries">client
	// library documentation</a> for more information. See
	// <a href="https://cloud.google.com/pubsub/quotas">Quotas and limits</a>
	// for more information about message limits.
	Message = v1.PubsubMessage
)

// EventHandler handles cloud events
type EventHandler interface {
	// Handle handles cloud events
	Handle(ctx context.Context, event *Event) error
}

// Receiver receives a pub sub message
type Receiver struct {
	// Codec represents the code
	Codec *pubsubevent.Codec
	// Handler handles the event
	Handler EventHandler
}

// Mount mounts the receiver to the server
func (r *Receiver) Mount(server *plex.Server) {
	// register grpc
	server.Socket().Register(RegisterReceiverServer, r)
	// register http
	server.Mux().Register(RegisterReceiverHandlerServer, r)
}

// Receive receives the pubsub message
func (r *Receiver) Receive(ctx context.Context, payload *ReceivedMessage) (*empty.Empty, error) {
	var (
		none   = &empty.Empty{}
		logger = log.GetContext(ctx)
	)

	if payload.Message == nil {
		err := status.Error(codes.InvalidArgument, "event body not provided")
		logger.WithError(err).Info("event validation fail")
		return none, err
	}

	if r.Codec == nil {
		r.Codec = &pubsubevent.Codec{
			Encoding: pubsubevent.StructuredV1,
		}
	}

	var (
		data   = payload.Message.Data
		buffer []byte
	)

	if _, err := base64.StdEncoding.Decode(buffer, data); err == nil {
		data = buffer
	}

	event, err := r.Codec.Decode(ctx,
		&pubsubevent.Message{
			Attributes: payload.Message.Attributes,
			Data:       data,
		},
	)

	if err != nil {
		err = status.Error(codes.InvalidArgument, err.Error())
		logger.WithError(err).Info("event decoding fail")
		return none, err
	}

	logger = logger.WithFields(log.Map{
		"event_id":       event.ID(),
		"event_type":     event.Type(),
		"event_source":   event.Source(),
		"event_receiver": "pubsub",
	})

	if err := event.Validate(); err != nil {
		err = status.Error(codes.InvalidArgument, err.Error())
		logger.WithError(err).Info("event validation fail")
		return none, err
	}

	ctx = log.SetContext(ctx, logger)
	ctx = r.context(ctx, payload)

	logger.Info("handling event start")

	if err := r.Handler.Handle(ctx, event); err != nil {
		if _, ok := status.FromError(err); !ok {
			err = status.Error(codes.Internal, err.Error())
		}

		logger.WithError(err).Info("event handling fail")
		return none, err
	}

	logger.Info("handling event success")
	return none, nil
}

func (r *Receiver) context(ctx context.Context, payload *ReceivedMessage) context.Context {
	var (
		msg = r.message(payload)
		tx  = TransportContext{
			ID:          msg.ID,
			PublishTime: msg.PublishTime,
			Topic:       TopicFromContext(ctx),
		}
	)

	if parts := strings.Split(payload.Subscription, "/"); len(parts) == 4 {
		tx.Project = parts[1]
		tx.Subscription = parts[3]
	}

	if name, ok := grpc.Method(ctx); ok {
		tx.Method = name
	}

	return ContextWithTransport(ctx, tx)
}

func (r *Receiver) message(payload *ReceivedMessage) *pubsub.Message {
	message := &pubsub.Message{
		ID:         payload.Message.MessageId,
		Attributes: payload.Message.Attributes,
		Data:       payload.Message.Data,
	}

	if timestamp, err := ptypes.Timestamp(payload.Message.PublishTime); err == nil {
		message.PublishTime = timestamp
	}

	return message
}
