package cloud

import "github.com/cloudevents/sdk-go/v2/context"

// WithLogger returns a new context with the logger injected into the given context.
var WithLogger = context.WithLogger
