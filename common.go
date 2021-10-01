package cloud

import (
	v2 "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/binding/format"
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
	// NewEvent returns a new Event, an optional version can be passed to change the
	// default spec version from 1.0 to the provided version.
	NewEvent = v2.NewEvent

	// NewClient creates a new client
	NewClient = v2.NewClient

	// NewHTTP creates a new http protocol
	NewHTTP = v2.NewHTTP

	// NewHTTPReceiveHandler creates a new HTTP handler
	NewHTTPReceiveHandler = v2.NewHTTPReceiveHandler

	// AddFormat adds a new Format. It can be retrieved by Lookup(f.MediaType())
	AddFormat = format.Add
)

var (
	TextPlain                       = v2.TextPlain
	ApplicationXML                  = v2.ApplicationXML
	ApplicationJSON                 = v2.ApplicationJSON
	ApplicationCloudEventsJSON      = v2.ApplicationCloudEventsJSON
	ApplicationCloudEventsBatchJSON = v2.ApplicationCloudEventsBatchJSON
)

// Dictionary represents a dictionary
type Dictionary = map[string]interface{}
