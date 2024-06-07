package test

import (
	"log"
	"time"

	"github.com/kozloz/togo/internal/genproto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Store is a test implementation of the app's storage functions. Primarily used for unit tests.
type Store struct {
}

var user1 *genproto.User = &genproto.User{
	ID:         1,
	DailyLimit: 2,
}

var user2 *genproto.User = &genproto.User{
	ID:         2,
	DailyLimit: 2,
	DailyCounter: &genproto.DailyCounter{
		DailyCount:  3,
		LastUpdated: timestamppb.New(time.Now().Add((-24 * time.Hour))),
	},
}

func (s *Store) GetUser(userID int64) (*genproto.User, error) {
	if userID == 1 {
		return user1, nil
	}
	if userID == 2 {
		return user2, nil
	}

	return nil, nil
}

func (s *Store) CreateTask(userID int64, task string) (*genproto.Task, error) {
	log.Println(userID, task)
	return &genproto.Task{
		ID:     1,
		UserID: userID,
		Name:   task,
	}, nil
}

func (s *Store) CreateUser(userID int64) (*genproto.User, error) {
	return &genproto.User{
		ID:         userID,
		DailyLimit: 5,
	}, nil
}

func (s *Store) UpdateUser(user *genproto.User) (*genproto.User, error) {
	log.Printf("Saving user %v", user)
	return user, nil
}
