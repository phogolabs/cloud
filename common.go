package cloud

import (
	"context"
	"fmt"
	"reflect"

	v2 "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/binding/format"
	"github.com/cloudevents/sdk-go/v2/event/datacodec"
	"github.com/cloudevents/sdk-go/v2/event/datacodec/json"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type (
	// Event represents the canonical representation of a CloudEvent.
	Event = v2.Event

	// Client represents a client
	Client = v2.Client

	// Result represents the result
	Result = v2.Result
)

var (
	// Decode looks up and invokes the decoder registered for the given content
	// type. An error is returned if no decoder is registered for the given
	// content type.
	JSONDecode = json.Decode
	// Encode looks up and invokes the encoder registered for the given content
	// type. An error is returned if no encoder is registered for the given
	// content type.
	JSONEncode = json.Encode
)

// JSONPBDecode takes `in` as []byte.
// If Event sent the payload as base64, Decoder assumes that `in` is the
// decoded base64 byte array.
func JSONPBDecode(ctx context.Context, in []byte, out interface{}) error {
	if in == nil {
		return nil
	}
	if out == nil {
		return fmt.Errorf("out is nil")
	}

	msg, ok := out.(proto.Message)
	if !ok {
		return fmt.Errorf("out is not proto.Message")
	}

	if err := protojson.Unmarshal(in, msg); err != nil {
		return fmt.Errorf("[jsonpb] found bytes \"%s\", but failed to unmarshal: %s", string(in), err.Error())
	}
	return nil
}

// Encode attempts to json.Marshal `in` into bytes. Encode will inspect `in`
// and returns `in` unmodified if it is detected that `in` is already a []byte;
// Or json.Marshal errors.
func JSONPBEncode(ctx context.Context, in interface{}) ([]byte, error) {
	if in == nil {
		return nil, nil
	}

	it := reflect.TypeOf(in)
	switch it.Kind() {
	case reflect.Slice:
		if it.Elem().Kind() == reflect.Uint8 {

			if b, ok := in.([]byte); ok && len(b) > 0 {
				// check to see if it is a pre-encoded byte string.
				if b[0] == byte('"') || b[0] == byte('{') || b[0] == byte('[') {
					return b, nil
				}
			}

		}
	}

	msg, ok := in.(proto.Message)
	if !ok {
		return nil, fmt.Errorf("out is not proto.Message")
	}

	return protojson.Marshal(msg)
}

var (
	// NewEvent returns a new Event, an optional version can be passed to change the
	// default spec version from 1.0 to the provided version.
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
	ApplicationProtobuf             = "application/protobuf"
)

// EventSender sends cloud events
type EventSender interface {
	// Send will transmit the given event over the client's configured transport.
	Send(ctx context.Context, event Event) Result
}
