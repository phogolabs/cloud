package gcp

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/googleapis/google-cloudevents-go/cloud/pubsub/v1"
	"github.com/googleapis/google-cloudevents-go/cloud/storage/v1"
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
				if err := MessageWriteRequest(r.Context(), message, r); err != nil {
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

				data := &storage.StorageObjectData{}
				// decode the data
				if err := json.Unmarshal(payload.Message.Data, data); err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}

				event, err := NewStorageEvent(payload.Message.Attributes["eventType"], data)
				if err != nil {
					http.Error(w, err.Error(), http.StatusUnprocessableEntity)
					return
				}

				message := MessageFromEvent(event)
				// write the context
				if err := MessageWriteRequest(r.Context(), message, r); err != nil {
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
