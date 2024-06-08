package eventhandler

import "context"

type Command interface {
	Execute(ctx context.Context)
	String() string
}

type CommandSerializer interface {
	Serialize(ctx context.Context, command string) (Command, error)
}
