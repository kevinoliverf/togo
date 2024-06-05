package main

import (
	"context"
	"log"

	"github.com/kozloz/togo/internal/errors"
	"github.com/kozloz/togo/internal/genproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// Validate the protobuf request
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
