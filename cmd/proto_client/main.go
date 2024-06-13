package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/kozloz/togo/internal/genproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "task"
	defaultId   = 1
)

var (
	addr   = flag.String("addr", "localhost:8081", "the address to connect to")
	userid = flag.Int("id", defaultId, "user id")
	name   = flag.String("name", defaultName, "name")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := genproto.NewTaskServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CreateTask(ctx, &genproto.CreateTaskRequest{UserID: int64(*userid), Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Response: %v", r)
}
