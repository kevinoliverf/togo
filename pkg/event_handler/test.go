package eventhandler

import (
	"context"
	"log"
)

type TestEvent struct {
	ID   string
	Type string
	Body string
}

func (t *TestEvent) GetID() string {
	return t.ID
}

func (t *TestEvent) GetType() string {
	return t.Type
}

func (t *TestEvent) GetBody() string {
	return t.Body
}

type TestFetcher struct {
	EventSource
}

func (t *TestFetcher) Fetch(ctx context.Context) ([]Event, error) {
	log.Println("Fetching test event")
	return []Event{
		&TestEvent{
			Type: "Test",
			Body: "test body",
		},
	}, nil
}

func (t *TestFetcher) Push(ctx context.Context, event Event) error {
	log.Println("Pushing test event")
	return nil
}

type TestCommand struct {
	command string
}

func (t *TestCommand) Execute(ctx context.Context) {
	log.Println("Executing test command with command:", t.command)
}

func (t *TestCommand) String() string {
	return t.command
}

type TestCommandDeserializer struct {
}

func (t *TestCommandDeserializer) Deserialize(ctx context.Context, command string) (Command, error) {
	return &TestCommand{command}, nil
}
