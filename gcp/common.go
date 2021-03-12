package gcp

import (
	v2 "github.com/cloudevents/sdk-go/v2"
)

type (
	// Client represents a client
	Client = v2.Client
)

// NewClient creates a new client
var NewClient = v2.NewClient
