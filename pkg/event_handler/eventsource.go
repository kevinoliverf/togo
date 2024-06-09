package eventhandler

import "context"

// EventSource is an interface that defines the methods that an event source(e.g. queue) should implement
type EventSource interface {
	Fetch(ctx context.Context) ([]Event, error)
	Push(ctx context.Context, event Event) error
}
