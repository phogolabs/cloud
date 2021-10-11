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

type PayloadType byte

const (
	// PayloadTypeNone represents the No Op decoder
	PayloadTypeNone PayloadType = iota
	// PayloadTypeQueue represents the Queue decoder
	PayloadTypeQueue
	// PayloadTypeStorage represents the Storage decoder
	PayloadTypeStorage
)

// Decoder adapts the request as a different format
func Decoder(kind PayloadType) func(http.Handler) http.Handler {
	handler := func(next http.Handler) http.Handler {
		switch kind {
		case PayloadTypeQueue:
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
		case PayloadTypeStorage:
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

				next.ServeHTTP(w, r)
			}

			return http.HandlerFunc(fn)
		default:
			return next
		}
	}

	return handler
}
