package cloud

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/phogolabs/plex"
)

// EventReceiver represents a receiver
type EventReceiver interface {
	Receive(ctx context.Context, eventArgs Event) Result
}

// EventReceiverFunc represents a receiver func
type EventReceiverFunc func(context.Context, Event) Result

// Receive receives the event
func (fn EventReceiverFunc) Receive(ctx context.Context, eventArgs Event) Result {
	return fn(ctx, eventArgs)
}

// EventHandler handles the events received from pubsub topic
type Webhook struct {
	routes  map[string]EventReceiver
	formats map[string]string
}

// NewWebhook creates a new event handler
func NewWebhook() *Webhook {
	return &Webhook{
		routes:  make(map[string]EventReceiver),
		formats: make(map[string]string),
	}
}

// Route registers receiver for given topic
func (h *Webhook) Route(topic string, receiver EventReceiver) {
	h.routes[topic] = receiver

	// Format marshals and unmarshals structured events to bytes.
	type Format interface {
		// MediaType identifies the format
		MediaType() string
	}

	if formatter, ok := receiver.(Format); ok {
		h.formats[topic] = formatter.MediaType()
	}
}

// Mount mounts the event handler to a given router
func (h *Webhook) Mount(r *plex.Server) {
	ctx := context.Background()

	protocol, err := NewHTTP()
	if err != nil {
		panic(err)
	}

	handler, err := NewHTTPReceiveHandler(ctx, protocol, h.receive)
	if err != nil {
		panic(err)
	}

	router := r.Proxy.Router()

	router.Route("/internal/topics/{topic}", func(route chi.Router) {
		route.Use(h.format)
		route.Mount("/", handler)
	})
}

func (h *Webhook) format(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		topic := chi.URLParam(r, "topic")

		if mediaType, ok := h.formats[topic]; ok {
			r.Header.Set("Content-Type", mediaType)
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func (h *Webhook) receive(ctx context.Context, eventArgs Event) Result {
	topic := chi.URLParamFromCtx(ctx, "topic")

	if handler, ok := h.routes[topic]; ok {
		return handler.Receive(ctx, eventArgs)
	}

	return nil
}
