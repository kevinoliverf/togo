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

type CreateTaskCommand struct {
	eventhandler.Command
	ProtoEncoded []byte
	conn         genproto.TaskServiceClient
}

func (e *CreateTaskCommand) Execute(ctx context.Context) {
	// Contact the server and print out its response.
	req := genproto.CreateTaskRequest{}
	protobuf, err := base64.StdEncoding.DecodeString(string(e.ProtoEncoded))
	if err != nil {
		log.Fatalf("could not decode: %v", err)
	}

	err = proto.Unmarshal(protobuf, &req)
	if err != nil {
		log.Fatalf("could not unmarshal: %v", err)
	}

	r, err := e.conn.CreateTask(ctx, &req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Response: %v", r)
}

func (e *CreateTaskCommand) String() string {
	return string(e.ProtoEncoded)
}

type CreateTaskCmdSerializer struct {
	eventhandler.CommandSerializer
	conn genproto.TaskServiceClient
}

func NewCreateTaskCmdSerializer(addr string) *CreateTaskCmdSerializer {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := genproto.NewTaskServiceClient(conn)
	return &CreateTaskCmdSerializer{
		conn: c,
	}
}

func (s *CreateTaskCmdSerializer) Serialize(ctx context.Context, command string) (eventhandler.Command, error) {
	return &CreateTaskCommand{
		ProtoEncoded: []byte(command),
		conn:         s.conn,
	}, nil
}
