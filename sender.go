package cloud

import "context"

var _ EventSender = &NopEventSender{}

// NopEventSender represent no-operation event sender
type NopEventSender struct{}

func (n *NopEventSender) Send(ctx context.Context, args Event) Result {
	return nil
}
