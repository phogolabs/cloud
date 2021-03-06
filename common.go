package cloud

import (
	v2 "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/binding/format"
)

type (
	// Event represents the canonical representation of a CloudEvent.
	Event = v2.Event
)

// NewEvent returns a new Event, an optional version can be passed to change the
// default spec version from 1.0 to the provided version.
var NewEvent = v2.NewEvent

// NewHTTP creates a new http protocol
var NewHTTP = v2.NewHTTP

// NewHTTPReceiveHandler creates a new HTTP handler
var NewHTTPReceiveHandler = v2.NewHTTPReceiveHandler

// Add a new Format. It can be retrieved by Lookup(f.MediaType())
var AddFormat = format.Add

var (
	ApplicationXML                  = v2.ApplicationXML
	ApplicationJSON                 = v2.ApplicationJSON
	TextPlain                       = v2.TextPlain
	ApplicationCloudEventsJSON      = v2.ApplicationCloudEventsJSON
	ApplicationCloudEventsBatchJSON = v2.ApplicationCloudEventsBatchJSON
)
