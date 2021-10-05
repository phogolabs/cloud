package cloud

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/phogolabs/log"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
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
	// router represents the HTTP handler that routes the requests to the
	// mounted receiver
	router chi.Router
	// protocol represents the HTTP protocol
	protocol *HTTPProtocol
}

// NewWebhook creates a new event handler
func NewWebhook() *Webhook {
	protocol, _ := NewHTTP()
	router := chi.NewRouter()

	// create the webhook for the give protocol
	webhook := &Webhook{
		router:   router,
		protocol: protocol,
	}

	// register the middlewares
	webhook.router.Use(webhook.logger)

	return webhook
}

// Mount mounts the receiver to a given path
func (h *Webhook) Mount(pattern string, receiver EventReceiver) {
	handler, _ := NewHTTPReceiveHandler(
		context.Background(), h.protocol, h.wrap(receiver))

	// Route represents a given route
	type Route interface {
		Use(chi.Router)
	}

	h.router.Group(func(r chi.Router) {
		if route, ok := receiver.(Route); ok {
			// use the route if needed
			route.Use(r)
		}
		// mount the handler to the router
		r.Mount(pattern, handler)
	})
}

func (h *Webhook) logger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		if logger, err := zap.NewProduction(); err == nil {
			logger = logger.Named("cloud")
			// owerwire the context
			ctx = WithLogger(ctx, logger.Sugar())
			// overwrite the request
			r = r.WithContext(ctx)
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func (h *Webhook) wrap(receiver EventReceiver) EventReceiver {
	fn := func(ctx context.Context, event Event) Result {
		// enrich the logger
		logger := log.GetContext(ctx).WithFields(log.Map{
			"incoming_event_id":                event.ID(),
			"incoming_event_type":              event.Type(),
			"incoming_event_source":            event.Source(),
			"incoming_event_subject":           event.Subject(),
			"incoming_event_data_schema":       event.DataSchema(),
			"incoming_event_data_content_type": event.DataContentType(),
		})

		// overwrite the context
		ctx = log.SetContext(ctx, logger)
		// overwrite the context
		ctx = h.metadata(ctx, event)

		// execute the receiver
		if err := receiver.Receive(ctx, event); err != nil {
			logger.WithError(err).Error("receiver failure")
			return err
		}

		return nil
	}

	return EventReceiverFunc(fn)
}

func (h *Webhook) metadata(ctx context.Context, event Event) context.Context {
	header := metadata.New(make(map[string]string))

	for k, v := range event.Extensions() {
		var (
			name  = fmt.Sprintf("x-%v", k)
			value = fmt.Sprintf("%v", v)
		)

		// append the pair
		header.Append(name, value)
	}

	return metadata.NewIncomingContext(ctx, header)
}
