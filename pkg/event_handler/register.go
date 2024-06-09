package eventhandler

// EventRegistry is responsible for registering command object serializers
type EventRegistry struct {
	handlers map[string]CommandDeserializer
}

// NewEventRegistry creates a new EventRegistry instance
func NewEventRegistry() *EventRegistry {
	return &EventRegistry{
		handlers: make(map[string]CommandDeserializer),
	}
}

// RegisterHandler registers a command serializer for a specific command type
func (r *EventRegistry) RegisterHandler(eventType string, serializer CommandDeserializer) {
	r.handlers[eventType] = serializer
}
