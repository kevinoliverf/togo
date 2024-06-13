package main

import (
	"github.com/kozloz/togo"
	"github.com/kozloz/togo/internal/genproto"
)

func TaskToGenProto(task *togo.Task) *genproto.Task {
	protoTask := &genproto.Task{
		ID:     task.ID,
		UserID: task.UserID,
		Name:   task.Name,
	}
	return protoTask
}
