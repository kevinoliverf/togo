package main

import (
	"context"
	"flag"

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
	client, _ := aws.NewClient(*queueURL)
	processor := eventhandler.NewEventProcessor(client, 1, 10)
	processor.RegisterHandler("Test", events.NewCreateTaskCmdSerializer(*addr))
	processor.Start(context.TODO())

}
