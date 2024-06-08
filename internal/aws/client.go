package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	eventhandler "github.com/kozloz/togo/pkg/event_handler"
)

type Client struct {
	eventhandler.EventSource
	session  *session.Session
	queueURL string
	client   *sqs.SQS
}

func NewClient(queueURL string) (*Client, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-2")},
	)
	if err != nil {
		return nil, err
	}

	// Create a SQS service client.

	return &Client{
		session:  sess,
		client:   sqs.New(sess),
		queueURL: queueURL,
	}, nil
}

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
