package eventhandler

import "context"

type EventSource interface {
	Fetch(ctx context.Context) ([]Event, error)
	Push(ctx context.Context, event Event) error
}
