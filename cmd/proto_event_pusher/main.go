package main

import (
	"context"
	"flag"
	"log"
	"strconv"

	"github.com/kozloz/togo/internal/aws"
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
	req := genproto.CreateTaskRequest{UserID: int64(*userid), Name: *name}
	reqBytes, _ := proto.Marshal(&req)
	log.Printf("Request: %v", reqBytes)
	event := aws.SQSEvent{
		ID:      strconv.Itoa(*userid),
		GroupID: "Test",
		Body:    string(reqBytes),
	}
	log.Printf("Event: %v\n", event)
	client, _ := aws.NewClient(*queueURL)
	client.Push(context.TODO(), &event)

}
