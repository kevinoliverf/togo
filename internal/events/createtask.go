package events

import (
	"context"
	"encoding/base64"
	"log"

	"github.com/kozloz/togo/internal/genproto"
	eventhandler "github.com/kozloz/togo/pkg/event_handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
)

const (
	CreateTaskCmdType = "TogoCreateTask"
)

// CreateTaskCommand is a Command to create a todo task
type CreateTaskCommand struct {
	eventhandler.Command
	ProtoEncoded []byte
	conn         genproto.TaskServiceClient
}

func (e *CreateTaskCommand) Execute(ctx context.Context) {
	req := genproto.CreateTaskRequest{}

	// Decode the base64 encoded proto
	protobuf, err := base64.StdEncoding.DecodeString(string(e.ProtoEncoded))
	if err != nil {
		log.Fatalf("could not decode: %v", err)
	}

	// Unmarshal the proto
	err = proto.Unmarshal(protobuf, &req)
	if err != nil {
		log.Fatalf("could not unmarshal: %v", err)
	}

	// Call the gRPC service
	r, err := e.conn.CreateTask(ctx, &req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Response: %v", r)
}

func (e *CreateTaskCommand) String() string {
	return string(e.ProtoEncoded)
}

// CreateTaskCmdDeserializer is a CommandSerializer for CreateTaskCommand
type CreateTaskCmdDeserializer struct {
	eventhandler.CommandDeserializer
	conn genproto.TaskServiceClient
}

// NewCreateTaskCmdSerializer creates a new CreateTaskCmdSerializer
func NewCreateTaskCmdSerializer(addr string) *CreateTaskCmdDeserializer {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not create client: %v", err)
		return nil
	}
	c := genproto.NewTaskServiceClient(conn)
	return &CreateTaskCmdDeserializer{
		conn: c,
	}
}

// Serialize deserializes a string into a CreateTaskCommand and includes the gRPC client
func (s *CreateTaskCmdDeserializer) Deserialize(ctx context.Context, command string) (eventhandler.Command, error) {
	return &CreateTaskCommand{
		ProtoEncoded: []byte(command),
		conn:         s.conn,
	}, nil
}
