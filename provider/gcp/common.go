package gcp

import (
	"github.com/cloudevents/sdk-go/v2/binding"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/cloudevents/sdk-go/v2/protocol/http"
)

var NewEvent = event.New
var NewMessage = http.NewMessage
var ToMessage = binding.ToMessage
var WriteRequest = http.WriteRequest
