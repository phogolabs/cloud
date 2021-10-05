package gcp

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/googleapis/google-cloudevents-go/cloud/pubsub/v1"
	"github.com/phogolabs/schema"
)

// PubsubDecoder represents the Pubsub decoder
func PubsubDecoder(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		payload := &pubsub.MessagePublishedData{}

		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// close the body
		r.Body.Close()

		header := http.Header{}
		// prepare the headers
		for key, value := range payload.Message.Attributes {
			header.Set(key, value)
		}

		message := NewMessage(header, io.NopCloser(
			bytes.NewBuffer(payload.Message.Data),
		))
		// write the message into the request
		if err := WriteRequest(r.Context(), message, r); err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

// DecodeStorage represents the Storage decoder
func StorageDecoder(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		payload := &pubsub.MessagePublishedData{}

		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// close the body
		r.Body.Close()

		var (
			eventType = payload.Message.Attributes["eventType"]
			bucketID  = payload.Message.Attributes["bucketId"]
			objectID  = payload.Message.Attributes["objectId"]
		)

		event := NewEvent()
		event.SetID(schema.NewUUID().String())
		event.SetTime(time.Now())
		event.SetSubject(bucketID + "/" + objectID)
		event.SetSource("https://" + bucketID)
		event.SetDataSchema("google.storage.object")
		event.SetDataContentType("application/json")

		switch eventType {
		case "OBJECT_FINALIZE":
			event.SetType("google.storage.object.finalize")
		case "OBJECT_ARCHIVE":
			event.SetType("google.storage.object.archive")
		case "OBJECT_DELETE":
			event.SetType("google.storage.object.delete")
		case "OBJECT_METADATA_UPDATE":
			event.SetType("google.storage.object.metadataUpdate")
		}

		// copy the actual data
		event.DataEncoded = payload.Message.Data

		if err := event.Validate(); err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		message := ToMessage(&event)
		// write the context
		if err := WriteRequest(r.Context(), message, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	return http.HandlerFunc(fn)
}
