package main

import (
	"context"
	"log"
	"strings"

	"github.com/bufbuild/protovalidate-go"
	"github.com/kozloz/togo/internal/errors"
	"github.com/kozloz/togo/internal/genproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var v *protovalidate.Validator

func init() {
	var err error
	v, err = protovalidate.New()
	if err != nil {
		log.Fatalf("failed to initialize protovalidate.Validator: %v", err)
	}
}

// Validate the protobuf request
func Validate(ctx context.Context, createTaskReq *genproto.CreateTaskRequest) error {
	log.Printf("Validating request: %v", createTaskReq)
	err := v.Validate(createTaskReq)
	if err != nil {
		cleanedErr := strings.ReplaceAll(err.Error(), "\n", " ")
		log.Printf("Error: request validation failed: %s", cleanedErr)

		_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "400"))
		return errors.InternalError
	}
	return nil
}
