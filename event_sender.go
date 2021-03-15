package cloud

import "context"

// EventSender sends cloud events
type EventSender interface {
	// Send will transmit the given event over the client's configured transport.
	Send(ctx context.Context, event Event) Result
}

var _ EventSender = &NopEventSender{}

// NopEventSender represent no-operation event sender
type NopEventSender struct{}

// Send sends the event
func (n *NopEventSender) Send(ctx context.Context, args Event) Result {
	return nil
}
