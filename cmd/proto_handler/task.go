package main

import (
	"context"
	"log"

	"github.com/kozloz/togo/internal/errors"
	"github.com/kozloz/togo/internal/genproto"
	"github.com/kozloz/togo/internal/tasks"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

type TaskHandler struct {
	genproto.UnimplementedTaskServiceServer
	op *tasks.Operation
}

func Validate(ctx context.Context, createTaskReq *genproto.CreateTaskRequest) error {
	log.Printf("Validating request: %v", createTaskReq)
	if createTaskReq.GetUserID() == 0 {
		log.Printf("Error: Invalid UserID provided")
		_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "400"))
		return errors.InvalidUserID
	}
	if createTaskReq.GetName() == "" {
		log.Printf("Error: Empty task name provided")
		_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "400"))
		return errors.InvalidTaskName
	}
	return nil
}

func (t *TaskHandler) CreateTask(ctx context.Context, createTaskReq *genproto.CreateTaskRequest) (*genproto.CreateTaskResponse, error) {
	log.Println(proto.Marshal(createTaskReq))

	createTaskResp := &genproto.CreateTaskResponse{}
	// Validate the request
	err := Validate(ctx, createTaskReq)
	if err != nil {
		resErr := CustomErrorToProto(errors.InvalidUserID)
		createTaskResp.Error = &resErr
		return createTaskResp, nil
	}

	// Create the task via operation class
	task, err := t.op.Create(createTaskReq.UserID, createTaskReq.Name)
	if err != nil {
		resErr := CustomErrorToProto(err)
		createTaskResp.Error = &resErr
		return createTaskResp, nil
	}

	createTaskResp.Task = task
	resErr := CustomErrorToProto(errors.Success)
	createTaskResp.Error = &resErr
	log.Printf("Task created: %v", task)
	return createTaskResp, nil

}
