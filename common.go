package cloud

import (
	"context"

	v2 "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/binding/format"
	"github.com/cloudevents/sdk-go/v2/event/datacodec"
)

type (
	// Event represents the canonical representation of a CloudEvent.
	Event = v2.Event

	// Client represents a client
	Client = v2.Client

	// Result represents the result
	Result = v2.Result
)

// NewEvent returns a new Event, an optional version can be passed to change the
// default spec version from 1.0 to the provided version.
var (
	NewEvent = v2.NewEvent

	// NewClient creates a new client
	NewClient = v2.NewClient

	// NewClientObserved creates an observable client
	NewClientObserved = v2.NewClientObserved

	// NewHTTP creates a new http protocol
	NewHTTP = v2.NewHTTP

	// NewHTTPReceiveHandler creates a new HTTP handler
	NewHTTPReceiveHandler = v2.NewHTTPReceiveHandler

	// Add a new Format. It can be retrieved by Lookup(f.MediaType())
	AddFormat = format.Add

	// AddDecoder registers a decoder for a given content type. The codecs will use
	// these to decode the data payload from a cloudevent.Event object.
	AddDecoder = datacodec.AddDecoder

	// AddEncoder registers an encoder for a given content type. The codecs will
	// use these to encode the data payload for a cloudevent.Event object.
	AddEncoder = datacodec.AddEncoder
)

var (
	ApplicationXML                  = v2.ApplicationXML
	ApplicationJSON                 = v2.ApplicationJSON
	TextPlain                       = v2.TextPlain
	ApplicationCloudEventsJSON      = v2.ApplicationCloudEventsJSON
	ApplicationCloudEventsBatchJSON = v2.ApplicationCloudEventsBatchJSON
)

// EventSender sends cloud events
type EventSender interface {
	// Send will transmit the given event over the client's configured transport.
	Send(ctx context.Context, event Event) Result
}
