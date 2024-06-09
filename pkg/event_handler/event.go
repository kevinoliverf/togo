package eventhandler

// Event is a message that is pushed and fetched from an EventSource
type Event interface {
	GetID() string
	GetType() string
	GetBody() string
}
