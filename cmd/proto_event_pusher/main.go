package main

import (
	"context"
	"flag"
	"log"
	"strconv"

	"github.com/kozloz/togo/internal/aws"
	"github.com/kozloz/togo/internal/events"
	"github.com/kozloz/togo/internal/genproto"
	"google.golang.org/protobuf/proto"
)

const (
	defaultName = "task"
	defaultId   = 1
)

var (
	userid   = flag.Int("id", defaultId, "user id")
	name     = flag.String("name", defaultName, "name")
	queueURL = flag.String("queue", "", "the address to connect to")
)

func main() {
	flag.Parse()
	// Create AWS SQS client
	client, err := aws.NewClient(*queueURL)
	if err != nil {
		log.Fatalf("Could not create aws client: %v", err)
	}

	// Create the proto request
	req := genproto.CreateTaskRequest{UserID: int64(*userid), Name: *name}
	reqBytes, err := proto.Marshal(&req)
	if err != nil {
		log.Fatalf("Could not marshal request: %v", err)
	}

	// Push event to SQS
	event := aws.SQSEvent{
		ID:      strconv.Itoa(*userid),
		GroupID: events.CreateTaskCmdType,
		Body:    string(reqBytes),
	}
	log.Printf("Event to push: %v", event)
	err = client.Push(context.TODO(), &event)
	if err != nil {
		log.Fatalf("Could not push event: %v", err)
	}

}
