package gcp

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// Decoder represents the Pub/Sub decoder
func Decoder(next http.Handler) http.Handler {
	// Message represents the message
	type Message struct {
		Attributes map[string]string `json:"attributes" validate:"required"`
		Data       []byte            `json:"data"`
	}

	// Payload represents the request payload
	type Payload struct {
		Message *Message `json:"message" validate:"required"`
	}

	fn := func(w http.ResponseWriter, r *http.Request) {
		payload := &Payload{}

		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// close the body
		r.Body.Close()

		if err := validator.New().StructCtx(r.Context(), payload); err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		// prepare the headers
		for key, value := range payload.Message.Attributes {
			r.Header.Set(key, value)
		}
		// prepare the body
		r.Body = io.NopCloser(bytes.NewBuffer(payload.Message.Data))

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
