package main

import (
	"log"
	"net"
	"time"

	"github.com/kozloz/togo/internal/genproto"
	"github.com/kozloz/togo/internal/store/mysql"
	"github.com/kozloz/togo/internal/tasks"
	"github.com/kozloz/togo/internal/users"
	"google.golang.org/grpc"
)

func InitializeServer(port string) {
	// Setup storage
	// Hard code db connection parameters for now
	store, err := mysql.NewStore("togo", "db", "3306", "togo", "togo")
	for i := 0; i < 2; i++ {
		if err == nil {
			time.Sleep(3 * time.Second)
			break
		}
		store, err = mysql.NewStore("togo", "db", "3306", "togo", "togo")
	}
	// Initialize operation classes
	userOp := users.NewOperation(store)
	taskOp := tasks.NewOperation(store, userOp)

	// Create the server
	taskHandler := TaskHandler{
		op: taskOp,
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	genproto.RegisterTaskServiceServer(s, &taskHandler)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
