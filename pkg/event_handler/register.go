package eventhandler

// EventRegistry is responsible for registering event handlers and dispatching events
type EventRegistry struct {
	handlers map[string]CommandSerializer
}

// NewEventRegistry creates a new EventRegistry instance
func NewEventRegistry() *EventRegistry {
	return &EventRegistry{
		handlers: make(map[string]CommandSerializer),
	}
}

// RegisterHandler registers an event handler for a specific event type
func (r *EventRegistry) RegisterHandler(eventType string, serializer CommandSerializer) {
	r.handlers[eventType] = serializer
}
