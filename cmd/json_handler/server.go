package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kozloz/togo/internal/store/test"
	"github.com/kozloz/togo/internal/tasks"
	"github.com/kozloz/togo/internal/users"
)

func InitializeServer() {
	// Setup storage
	store := &test.Store{}
	userOp := users.NewOperation(store)
	taskOp := tasks.NewOperation(store, userOp)
	// Create the server
	taskHandler := TaskHandler{
		op: taskOp,
	}
	router := mux.NewRouter()

	// RESTful API for creating a task. A user is required to make a task.
	router.HandleFunc("/users/{id}/tasks", taskHandler.CreateTask).Methods(http.MethodPost)

	// Start the server
	log.Fatal(http.ListenAndServe(":8000", router))
}
