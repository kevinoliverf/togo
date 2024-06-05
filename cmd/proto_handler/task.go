package main

import (
	"context"
	"log"

	"github.com/kozloz/togo/internal/errors"
	"github.com/kozloz/togo/internal/genproto"
	"github.com/kozloz/togo/internal/tasks"
	"google.golang.org/protobuf/proto"
)

type TaskHandler struct {
	genproto.UnimplementedTaskServiceServer
	op *tasks.Operation
}

func (t *TaskHandler) CreateTask(ctx context.Context, createTaskReq *genproto.CreateTaskRequest) (*genproto.CreateTaskResponse, error) {
	log.Println(proto.Marshal(createTaskReq))
	// Create the task via operation class
	createTaskResp := &genproto.CreateTaskResponse{}
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
