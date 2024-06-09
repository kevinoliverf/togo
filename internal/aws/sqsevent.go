package aws

import (
	"fmt"

	eventhandler "github.com/kozloz/togo/pkg/event_handler"
)

type SQSEvent struct {
	eventhandler.Event
	ID      string
	GroupID string
	Body    string
}

func (s *SQSEvent) GetID() string {
	return s.ID
}
func (s *SQSEvent) GetType() string {
	return s.GroupID
}

func (s *SQSEvent) GetBody() string {
	return s.Body
}

func (s *SQSEvent) String() string {
	return fmt.Sprintf("ID: %s, GroupID: %s, Body: %s\n", s.ID, s.GroupID, s.Body)
}
