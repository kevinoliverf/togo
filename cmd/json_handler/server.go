package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/kozloz/togo/internal/genproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitializeServer(port string, protoEndpoint string) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux(runtime.WithForwardResponseOption(httpResponseModifier))
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := genproto.RegisterTaskServiceHandlerFromEndpoint(ctx, mux, protoEndpoint, opts)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	http.ListenAndServe(":"+port, mux)
}
