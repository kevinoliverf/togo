package eventhandler

import "context"

// Command executes operations after an event message is serialized
type Command interface {
	Execute(ctx context.Context)
	String() string
}

// CommandDeserializer serializes a message string into a Command object
type CommandDeserializer interface {
	Deserialize(ctx context.Context, command string) (Command, error)
}
