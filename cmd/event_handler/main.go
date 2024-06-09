package main

import (
	"context"
	"flag"
	"log"

	"github.com/kozloz/togo/internal/aws"
	"github.com/kozloz/togo/internal/events"
	eventhandler "github.com/kozloz/togo/pkg/event_handler"
)

var (
	addr     = flag.String("addr", "localhost:8081", "the address to connect to")
	queueURL = flag.String("queue", "", "the address to connect to")
)

func main() {
	flag.Parse()

	// Create AWS SQS Client
	client, err := aws.NewClient(*queueURL)
	if err != nil {
		log.Fatalf("Could not create aws client: %v", err)
	}

	// Create Event Processor
	processor := eventhandler.NewEventProcessor(client, 1, 10)

	// Register Create Task Command Serializer
	processor.RegisterHandler(events.CreateTaskCmdType, events.NewCreateTaskCmdSerializer(*addr))

	// Start the processor
	processor.Start(context.TODO())

}
