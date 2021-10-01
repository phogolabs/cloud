package cloud

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/phogolabs/log"
	"github.com/phogolabs/plex"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type HandlerFunc = func(http.Handler) http.Handler

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
	routes     map[string]EventReceiver
	middleware map[string]HandlerFunc
}

// NewWebhook creates a new event handler
func NewWebhook() *Webhook {
	return &Webhook{
		routes:     make(map[string]EventReceiver),
		middleware: make(map[string]HandlerFunc),
	}
}

// Route registers receiver for given topic
func (h *Webhook) Route(topic string, receiver EventReceiver) {
	// Interceptor represents an http interceptor
	type Interceptor interface {
		// Middleware represents a middleware
		Intercept(http.Handler) http.Handler
	}

	h.routes[topic] = receiver

	if endpoint, ok := receiver.(Interceptor); ok {
		h.middleware[topic] = endpoint.Intercept
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
		route.Use(h.logger)
		route.Use(h.trailer)
		route.Mount("/", handler)
	})
}

func (h *Webhook) logger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if logger, err := zap.NewProduction(); err == nil {
			r = r.WithContext(
				WithLogger(r.Context(), logger.Named("cloud").Sugar()),
			)
		}
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func (h *Webhook) trailer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		topic := chi.URLParam(r, "topic")

		if middleare, ok := h.middleware[topic]; ok {
			next = middleare(next)
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func (h *Webhook) receive(ctx context.Context, eventArgs Event) Result {
	topic := chi.URLParamFromCtx(ctx, "topic")

	// enrich the logger
	logger := log.GetContext(ctx).WithFields(log.Map{
		"incoming_event_id":                eventArgs.ID(),
		"incoming_event_type":              eventArgs.Type(),
		"incoming_event_source":            eventArgs.Source(),
		"incoming_event_source_topic":      topic,
		"incoming_event_subject":           eventArgs.Subject(),
		"incoming_event_data_schema":       eventArgs.DataSchema(),
		"incoming_event_data_content_type": eventArgs.DataContentType(),
	})

	// find the handler for given topic
	handler, ok := h.routes[topic]
	if !ok {
		err := status.Errorf(codes.NotFound, "receiver %s not found", topic)
		// log the error
		logger.WithError(err).Error("receiver does not exist")
		// stop eht execution
		return err
	}

	// overwrite the logger
	ctx = log.SetContext(ctx, logger)
	// enrich the context with extensions
	ctx = h.metadata(ctx, eventArgs)

	// execute the receiver
	if err := handler.Receive(ctx, eventArgs); err != nil {
		logger.WithError(err).Error("receiver failure")
		// stop eht execution
		return err
	}

	return nil
}

func (h *Webhook) metadata(ctx context.Context, eventArgs Event) context.Context {
	kv := metadata.New(make(map[string]string))

	for k, v := range eventArgs.Extensions() {
		var (
			key   = fmt.Sprintf("x-%v", k)
			value = fmt.Sprintf("%v", v)
		)

		// append the pair
		kv.Append(key, value)
	}

	return metadata.NewIncomingContext(ctx, kv)
}
