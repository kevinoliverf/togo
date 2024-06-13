package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
	grpcmux := runtime.NewServeMux(runtime.WithForwardResponseOption(httpResponseModifier))
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := genproto.RegisterTaskServiceHandlerFromEndpoint(ctx, grpcmux, protoEndpoint, opts)
	if err != nil {
		log.Fatal(err)
		return
	}

	router := mux.NewRouter()
	router.Use(loggingMiddleware)

	// Match v1 api and strip before calling the proto handler
	router.PathPrefix("/v1").Handler(http.StripPrefix("/v1",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			grpcmux.ServeHTTP(w, r)
		})))

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	http.ListenAndServe(":"+port, router)
}
